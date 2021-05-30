package utils

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"text/tabwriter"
)

// AssertEqual checks if values are equal
// Ref: gofiber/utils
func AssertEqual(t testing.TB, expected, actual interface{}, description ...string) {
	if reflect.DeepEqual(expected, actual) {
		return
	}

	var aType = "<nil>"
	var bType = "<nil>"

	if expected != nil {
		aType = reflect.TypeOf(expected).String()
	}
	if actual != nil {
		bType = reflect.TypeOf(actual).String()
	}

	testName := "AssertEqual"
	if t != nil {
		testName = t.Name()
	}

	_, file, line, _ := runtime.Caller(1)

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 5, ' ', 0)
	_, _ = fmt.Fprintf(w, "\nTest:\t%s", testName)
	_, _ = fmt.Fprintf(w, "\nTrace:\t%s:%d", filepath.Base(file), line)
	if len(description) > 0 {
		_, _ = fmt.Fprintf(w, "\nDescription:\t%s", description[0])
	}
	_, _ = fmt.Fprintf(w, "\nExpect:\t%v\t(%s)", expected, aType)
	_, _ = fmt.Fprintf(w, "\nResult:\t%v\t(%s)", actual, bType)

	result := ""
	if err := w.Flush(); err != nil {
		result = err.Error()
	} else {
		result = buf.String()
	}

	if t != nil {
		t.Fatal(result)
	} else {
		log.Fatal(result)
	}
}

// 断言 panic
func AssertPanics(t *testing.T, title string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			if t != nil {
				t.Fatalf("%s: didn't panic as expected", title)
			} else {
				log.Fatalf("%s: didn't panic as expected", title)
			}
		}
	}()

	f()
}

// 判断对象(pointer, channel, func, interface, map, slice)是否为 nil
// nil 是一个 Type 类型的变量, Type 类型是基于 int 的类型
// var 若变量本身是指针, 占用 8 字节, 指向类型内部结构体并置 0, 仅定义了变量本身, 此时为 nil
//     指针是非复合类型, 赋值 nil 时, 将 8 字节置 0, 即没有指向任何值的指针 0x0
//     map, channel: var 时仅定义了指针, 需要 make 初始化内部结构后才能使用, make 后非 nil
// var 若变量非指针, 如 struct, int, 非 nil
// slice:
//   type slice struct, 占用 24 字节, 1 指针(array unsafe.Pointer) 2 个整型字段(len, cap int)
//   var 定义后即可使用, 置 0 并分配, 此时 array 指针为 0 即没有实际数据时为 nil
// interface:
//   type iface struct(interface 类型), type eface struct(空接口), 占用 16 字节
//   判断 data 指针为 0 即为 nil, 初始化后即非 0
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}

	defer func() {
		recover()
	}()

	return reflect.ValueOf(i).IsNil()
}
