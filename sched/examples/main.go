package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/fufuok/utils/sched"
)

func main() {
	count := int64(0)
	bus := sched.New() // 默认并发数: runtime.NumCPU()
	for i := 0; i < 30; i++ {
		bus.Add(1)
		bus.RunWithArgs(func(n ...interface{}) {
			atomic.AddInt64(&count, int64(n[0].(int)))
		}, i)
	}
	bus.Wait()
	fmt.Println("count:", atomic.LoadInt64(&count)) // count: 435

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
