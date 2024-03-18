package assert

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"text/tabwriter"
)

func True(tb testing.TB, value bool, msgAndArgs ...interface{}) {
	if tb != nil {
		tb.Helper()
	}
	if value {
		return
	}
	result := "Should be true"
	assertLog(tb, nil, value, "True", result, msgAndArgs...)
}

func False(tb testing.TB, value bool, msgAndArgs ...interface{}) {
	if tb != nil {
		tb.Helper()
	}
	if !value {
		return
	}
	result := "Should be false"
	assertLog(tb, nil, value, "False", result, msgAndArgs...)
}

func NotNil(tb testing.TB, value interface{}, msgAndArgs ...interface{}) {
	if tb != nil {
		tb.Helper()
	}
	if !IsNil(value) {
		return
	}
	result := "Expected value not to be nil."
	assertLog(tb, nil, value, "NotNil", result, msgAndArgs...)
}

func Nil(tb testing.TB, value interface{}, msgAndArgs ...interface{}) {
	if tb != nil {
		tb.Helper()
	}
	if IsNil(value) {
		return
	}
	result := fmt.Sprintf("Expected nil, but got: %#v", value)
	assertLog(tb, nil, value, "Nil", result, msgAndArgs...)
}

func NotEmpty(tb testing.TB, value interface{}, msgAndArgs ...interface{}) {
	if tb != nil {
		tb.Helper()
	}
	if !IsEmpty(value) {
		return
	}
	result := fmt.Sprintf("Should NOT be empty, but was %v", value)
	assertLog(tb, nil, value, "NotEmpty", result, msgAndArgs...)
}

func Empty(tb testing.TB, value interface{}, msgAndArgs ...interface{}) {
	if tb != nil {
		tb.Helper()
	}
	if IsEmpty(value) {
		return
	}
	result := fmt.Sprintf("Should be empty, but was %v", value)
	assertLog(tb, nil, value, "Empty", result, msgAndArgs...)
}

func Contains(tb testing.TB, value string, msgAndArgs ...string) {
	if tb != nil {
		tb.Helper()
	}
	for _, s := range msgAndArgs {
		if ok := strings.Contains(s, value); !ok {
			result := fmt.Sprintf("%s does not contain %s", s, value)
			assertLog(tb, nil, nil, "Contains", result, msgAndArgs)
		}
	}
}

func NotEqual(tb testing.TB, expected, actual interface{}, msgAndArgs ...interface{}) {
	if tb != nil {
		tb.Helper()
	}
	if !DeepEqual(expected, actual) {
		return
	}
	result := fmt.Sprintf("Should not be: %#v\n", actual)
	assertLog(tb, expected, actual, "NotEqual", result, msgAndArgs...)
}

// Equal checks if values are equal
// Ref: gofiber/utils
func Equal(tb testing.TB, expected, actual interface{}, msgAndArgs ...interface{}) {
	if tb != nil {
		tb.Helper()
	}
	if DeepEqual(expected, actual) {
		return
	}
	assertLog(tb, expected, actual, "Equal", "", msgAndArgs...)
}

func assertLog(tb testing.TB, a, b interface{}, testType, result string, msgAndArgs ...interface{}) {
	if tb != nil {
		testType = fmt.Sprintf("%s(%s)", tb.Name(), testType)
	}

	_, file, line, _ := runtime.Caller(2)

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 5, ' ', 0)
	_, _ = fmt.Fprintf(w, "\nTest:\t%s", testType)
	_, _ = fmt.Fprintf(w, "\nTrace:\t%s:%d", filepath.Base(file), line)
	message := messageFromMsgAndArgs(msgAndArgs...)
	if message != "" {
		_, _ = fmt.Fprintf(w, "\nDescription:\t%s", message)
	}
	if result != "" {
		_, _ = fmt.Fprintf(w, "\nResult:\t%s", result)
	} else {
		aType := "<nil>"
		bType := "<nil>"
		if a != nil {
			aType = reflect.TypeOf(a).String()
		}
		if b != nil {
			bType = reflect.TypeOf(b).String()
		}
		_, _ = fmt.Fprintf(w, "\nResult:\tNot equal")
		_, _ = fmt.Fprintf(w, "\nExpected:\t%v\t(%s)", a, aType)
		_, _ = fmt.Fprintf(w, "\nActual:\t%v\t(%s)", b, bType)
	}

	msg := ""
	if err := w.Flush(); err != nil {
		msg = err.Error()
	} else {
		msg = buf.String()
	}

	if tb != nil {
		tb.Fatal(msg)
	} else {
		log.Fatal(msg)
	}
}

// Ref: stretchr/testify
func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return runeSubString(msgAsStr, 300, "..")
		}
		return runeSubString(fmt.Sprintf("%+v", msg), 300, "..")
	}
	if len(msgAndArgs) > 1 {
		return runeSubString(fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...), 300, "..")
	}
	return ""
}

// Panics 断言 panic
func Panics(t *testing.T, title string, f func()) {
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

// DeepEqual Ref: stretchr/testify
func DeepEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}

// IsNil 判断对象(pointer, channel, func, interface, map, slice)是否为 nil
// nil 是一个 Type 类型的变量, Type 类型是基于 int 的类型
// var 若变量本身是指针, 占用 8 字节, 指向类型内部结构体并置 0, 仅定义了变量本身, 此时为 nil
//
//	指针是非复合类型, 赋值 nil 时, 将 8 字节置 0, 即没有指向任何值的指针 0x0
//	map, channel: var 时仅定义了指针, 需要 make 初始化内部结构后才能使用, make 后非 nil
//
// var 若变量非指针, 如 struct, int, 非 nil
// slice:
//
//	type slice struct, 占用 24 字节, 1 指针(array unsafe.Pointer) 2 个整型字段(len, cap int)
//	var 定义后即可使用, 置 0 并分配, 此时 array 指针为 0 即没有实际数据时为 nil
//
// interface:
//
//	type iface struct(interface 类型), type eface struct(空接口), 占用 16 字节
//	判断 data 指针为 0 即为 nil, 初始化后即非 0
//
// Ref: stretchr/testify
func IsNil(o interface{}) bool {
	if o == nil {
		return true
	}

	value := reflect.ValueOf(o)
	kind := value.Kind()
	isNilableKind := containsKind(
		[]reflect.Kind{
			reflect.Chan, reflect.Func,
			reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice, reflect.UnsafePointer,
		},
		kind)

	if isNilableKind && value.IsNil() {
		return true
	}

	return false
}

// containsKind checks if a specified kind in the slice of kinds.
// Ref: stretchr/testify
func containsKind(kinds []reflect.Kind, kind reflect.Kind) bool {
	for i := 0; i < len(kinds); i++ {
		if kind == kinds[i] {
			return true
		}
	}

	return false
}

// IsEmpty gets whether the specified object is considered empty or not.
// Ref: stretchr/testify
func IsEmpty(o interface{}) bool {
	// get nil case out of the way
	if o == nil {
		return true
	}

	v := reflect.ValueOf(o)
	switch v.Kind() {
	// collection types are empty when they have no element
	case reflect.Chan, reflect.Map, reflect.Slice:
		return v.Len() == 0
	// pointers are empty if nil or if the value they point to is empty
	case reflect.Ptr:
		if v.IsNil() {
			return true
		}
		deref := v.Elem().Interface()
		return IsEmpty(deref)
	// for all other types, compare against the zero value
	// array types are empty when they match their zero-initialized state
	default:
		zero := reflect.Zero(v.Type())
		return reflect.DeepEqual(o, zero.Interface())
	}
}
