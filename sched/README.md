### 并发任务调度库

见: [sched](sched)

或: https://github.com/fufuok/sched

简洁, 高效, 并发限制, 复用 goroutine

```go
package sched // import "github.com/fufuok/utils/sched"

type Option func(w *Pool)
    func Queues(limit int) Option
    func Workers(limit int) Option
type Pool struct{ ... }
    func New(opts ...Option) *Pool
```

## 使用示例

```go
package main

import (
	"fmt"
	"time"

	"github.com/fufuok/utils/sched"
	"github.com/fufuok/utils/xsync"
)

func main() {
	count := xsync.NewCounter()
	bus := sched.New() // 默认并发数: runtime.NumCPU()
	for i := 0; i < 30; i++ {
		bus.Add(1)
		bus.RunWithArgs(func(n ...interface{}) {
			count.Add(int64(n[0].(int)))
		}, i)
	}
	bus.Wait()
	fmt.Println("count:", count.Value()) // count: 435

	// 继续下一批任务
	bus.Add(1)
	bus.Run(func() {
		fmt.Println("is running:", bus.IsRunning(), bus.Running()) // is running: true 1
	})
	bus.Wait()
	bus.Release()

	// 指定并发数, 指定队列缓冲数
	bus = sched.New(sched.Workers(2), sched.Queues(1))
	bus.Add(5)
	for i := 0; i < 5; i++ {
		bus.Run(func() {
			fmt.Println(time.Now())
			time.Sleep(time.Second)
		})
	}
	bus.WaitAndRelease()
	fmt.Println("is running:", bus.IsRunning()) // is running: false
}
```







*ff*