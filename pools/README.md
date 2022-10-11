### 常用的池

`[]byte` 更多功能的字节切片池化见: [github.com/fufuok/bytespool](https://github.com/fufuok/bytespool)

```go
package bytespool // import "github.com/fufuok/utils/pools/bytespool"

func Append(buf []byte, elems ...byte) []byte
func AppendString(buf []byte, elems string) []byte
func Get(size int) []byte
func Make(size int) []byte
func New(size int) []byte
func Put(buf []byte)
func Release(buf []byte) bool
func SetMaxSize(size int) bool
type CapacityPools struct{ ... }

package bufferpool // import "github.com/fufuok/utils/pools/bufferpool"

func Get() *bytes.Buffer
func New(bs []byte) *bytes.Buffer
func NewByte(c byte) *bytes.Buffer
func NewRune(r rune) *bytes.Buffer
func NewString(s string) *bytes.Buffer
func Put(buf *bytes.Buffer)
func Release(buf *bytes.Buffer) bool
func SetMaxSize(size int) bool

package readerpool // import "github.com/fufuok/utils/pools/readerpool"

func New(b []byte) *bytes.Reader
func Release(r *bytes.Reader)

package timerpool // import "github.com/fufuok/utils/pools/timerpool"

func New(d time.Duration) *time.Timer
func Release(t *time.Timer)

package tickerpool // import "github.com/fufuok/utils/pools/tickerpool"

func New(d time.Duration) *time.Ticker
func Release(t *time.Ticker)
```







*ff*