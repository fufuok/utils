### 高效的 JSON 字符串操作库集

1. **只有 `1` 次内存分配的 JSON 字符串生成器**

见: [jsongen](xjson/jsongen)

或: [https://github.com/fufuok/jsongen](https://github.com/fufuok/jsongen)

```go
package jsongen // import "github.com/fufuok/utils/xjson/jsongen"

Package jsongen forked from darjun/json-gen

type Array []Value
    func NewArray() *Array
type Map struct{ ... }
    func NewMap() *Map
type QuotedValue string
type RawBytes []byte
type RawString string
type UnquotedValue string
type Value interface{ ... }
    func EscapeString(s string) Value
```

2. **超高效的 JSON 字符串解析和字段搜索**

来自: `tidwall/gjson`

见: [gjson](xjson/gjson)

```go
package gjson // import "github.com/fufuok/utils/xjson/gjson"

Package gjson provides searching for json strings.

var DisableModifiers = false
func AddModifier(name string, fn func(json, arg string) string)
func AppendJSONString(dst []byte, s string) []byte
func ForEachLine(json string, iterator func(line Result) bool)
func ModifierExists(name string, fn func(json, arg string) string) bool
func Valid(json string) bool
func ValidBytes(json []byte) bool
type Result struct{ ... }
    func Get(json, path string) Result
    func GetBytes(json []byte, path string) Result
    func GetMany(json string, path ...string) []Result
    func GetManyBytes(json []byte, path ...string) []Result
    func Parse(json string) Result
    func ParseBytes(json []byte) Result
    func (t Result) Array() []Result
    func (t Result) Bool() bool
    func (t Result) Exists() bool
    func (t Result) Float() float64
    func (t Result) ForEach(iterator func(key, value Result) bool)
    func (t Result) Get(path string) Result
    func (t Result) Int() int64
    func (t Result) IsArray() bool
    func (t Result) IsBool() bool
    func (t Result) IsObject() bool
    func (t Result) Less(token Result, caseSensitive bool) bool
    func (t Result) Map() map[string]Result
    func (t Result) Path(json string) string
    func (t Result) Paths(json string) []string
    func (t Result) String() string
    func (t Result) Time() time.Time
    func (t Result) Uint() uint64
    func (t Result) Value() interface{}
type Type int
    const Null Type = iota ...
```

3. **JSON 字符串字段修改和删除**

来自: `tidwall/sjson`

见: [sjson](xjson/sjson)

```go
package sjson // import "github.com/fufuok/utils/xjson/sjson"

Package sjson provides setting json values.

func Delete(json, path string) (string, error)
func DeleteBytes(json []byte, path string) ([]byte, error)
func Set(json, path string, value interface{}) (string, error)
func SetBytes(json []byte, path string, value interface{}) ([]byte, error)
func SetBytesOptions(json []byte, path string, value interface{}, opts *Options) ([]byte, error)
func SetOptions(json, path string, value interface{}, opts *Options) (string, error)
func SetRaw(json, path, value string) (string, error)
func SetRawBytes(json []byte, path string, value []byte) ([]byte, error)
func SetRawBytesOptions(json []byte, path string, value []byte, opts *Options) ([]byte, error)
func SetRawOptions(json, path, value string, opts *Options) (string, error)
type Options struct{ ... }
```

4. **JSON 字符串格式化和校验**

来自: `tidwall/pretty`

见: [pretty](xjson/pretty)

```go
package pretty // import "github.com/fufuok/utils/xjson/pretty"

var DefaultOptions = &Options{ ... }
func Color(src []byte, style *Style) []byte
func Pretty(json []byte) []byte
func PrettyOptions(json []byte, opts *Options) []byte
func Spec(src []byte) []byte
func SpecInPlace(src []byte) []byte
func Ugly(json []byte) []byte
func UglyInPlace(json []byte) []byte
type Options struct{ ... }
type Style struct{ ... }
    var TerminalStyle *Style
```

5. 字符串模式匹配(`*?`通配符搜索)

来自: `tidwall/match`

见: [match](xjson/match)

```go
package match // import "github.com/fufuok/utils/xjson/match"

Package match provides a simple pattern matcher with unicode support.

func Allowable(pattern string) (min, max string)
func IsPattern(str string) bool
func Match(str, pattern string) bool
func MatchLimit(str, pattern string, maxcomp int) (matched, stopped bool)
```







*ff*