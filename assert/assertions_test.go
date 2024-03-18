package assert

import (
	"bytes"
	"errors"
	"os"
	"testing"
	"time"
	"unsafe"
)

func TestAssertTrue(t *testing.T) {
	t.Parallel()
	True(nil, true)
	True(t, true)
	False(nil, false)
	False(t, false)
}

func TestAssertEqual(t *testing.T) {
	t.Parallel()
	Equal(nil, []string{}, []string{})
	Equal(t, []byte{}, []byte{})
	Equal(t, []byte{0, 1, 2}, []byte{0, 1, 2})
	Equal(t, map[int]int{1: 2}, map[int]int{1: 2})
	Equal(t, []string{}, []string{})
	Equal(t, &struct{}{}, &struct{}{})

	var bs []byte
	NotEqual(t, bs, []byte{})
	NotEqual(t, []byte(nil), []byte{})
	NotEqual(t, []byte{0, 2, 1}, []byte{0, 1, 2})
	NotEqual(t, map[int]int{2: 2}, map[int]int{1: 2})
	NotEqual(t, []string{"f"}, []string{}, "desc: %s", "abc")
	NotEqual(t, nil, []string{})
	NotEqual(t, func() {}, func() {})
	NotEqual(t, 'f', "f")
	NotEqual(t, 10, uint(10))
}

func TestAssertPanics(t *testing.T) {
	t.Parallel()
	var a []int
	Panics(nil, "should panic when index out of range", func() {
		_ = a[1]
	})
	Panics(t, "should panic when index out of range", func() {
		_ = a[1]
	})
}

func TestIsNil(t *testing.T) {
	t.Parallel()
	var map1 map[int]bool
	Nil(t, map1)
	map2 := make(map[int]bool)
	NotNil(t, map2)

	var ch1 chan int
	Nil(t, ch1)
	ch2 := make(chan struct{})
	NotNil(t, ch2)

	var slice1 []string
	Nil(t, slice1)
	slice2 := slice1
	Nil(t, slice2)
	slice1 = append(slice1, "")
	NotNil(t, slice1)
	Nil(t, slice2)

	slice3 := make([]int, 0)
	NotNil(t, slice3)
	slice4 := slice3
	NotNil(t, slice4)
	slice3 = nil
	Nil(t, slice3)
	NotNil(t, slice4)

	var iface1 interface{}
	Nil(t, iface1)
	iface1 = nil
	Nil(t, iface1)
	iface1 = map1
	Nil(t, iface1)
	iface1 = map2
	NotNil(t, iface1)
	var iface2 interface{} = (*int)(nil)
	Nil(t, iface2)

	var eface1 error
	Nil(t, eface1)
	eface2 := new(error)
	NotNil(t, eface2)
	var iface3 interface{} = eface1
	Nil(t, iface3)

	var ptr *int
	Nil(t, ptr)

	var iface4 interface{} = ptr
	Nil(t, iface4)
	// Equal(t, false, iface4 == nil) // go1.16.4

	var fun func(int) error
	Nil(t, fun)

	struct0 := new(bytes.Buffer)
	NotNil(t, struct0)
	var struct1 *bytes.Buffer
	Nil(t, struct1)
	var struct2 bytes.Buffer
	NotNil(t, struct2)
	struct3 := &bytes.Buffer{}
	NotNil(t, struct3)
	var struct4 *struct{}
	Nil(t, struct4)
	var struct5 struct{}
	NotNil(t, struct5)
	struct6 := struct{}{}
	NotNil(t, struct6)

	var s string
	NotNil(t, s)
	var iface5 interface{} = s
	NotNil(t, iface5)

	var n unsafe.Pointer = nil
	Nil(t, n)

	// var nil1 = (*int)(unsafe.Pointer(uintptr(0x0)))
	// Equal(t, true, IsNil(nil1))
}

// Ref: stretchr/testify
func Test_IsEmpty(t *testing.T) {
	chWithValue := make(chan struct{}, 1)
	chWithValue <- struct{}{}

	True(t, IsEmpty(struct{}{}))
	True(t, IsEmpty(""))
	True(t, IsEmpty(nil))
	True(t, IsEmpty([]string{}))
	True(t, IsEmpty(0))
	True(t, IsEmpty(int32(0)))
	True(t, IsEmpty(int64(0)))
	True(t, IsEmpty(false))
	True(t, IsEmpty(map[string]string{}))
	True(t, IsEmpty(new(time.Time)))
	True(t, IsEmpty(time.Time{}))
	True(t, IsEmpty(make(chan struct{})))
	True(t, IsEmpty([1]int{}))

	False(t, IsEmpty("something"))
	False(t, IsEmpty(errors.New("something")))
	False(t, IsEmpty([]string{"something"}))
	False(t, IsEmpty(1))
	False(t, IsEmpty(true))
	False(t, IsEmpty(map[string]string{"Hello": "World"}))
	False(t, IsEmpty(chWithValue))
	False(t, IsEmpty([1]int{42}))
}

// Ref: stretchr/testify
func TestEmpty(t *testing.T) {
	chWithValue := make(chan struct{}, 1)
	chWithValue <- struct{}{}
	var tiP *time.Time
	var tiNP time.Time
	var s *string
	var f *os.File
	sP := &s
	x := 1
	xP := &x

	type TString string
	type TStruct struct {
		x int
	}

	Empty(t, "", "Empty string is empty")
	Empty(t, nil, "Nil is empty")
	Empty(t, []string{}, "Empty string array is empty")
	Empty(t, 0, "Zero int value is empty")
	Empty(t, false, "False value is empty")
	Empty(t, make(chan struct{}), "Channel without values is empty")
	Empty(t, s, "Nil string pointer is empty")
	Empty(t, f, "Nil os.File pointer is empty")
	Empty(t, tiP, "Nil time.Time pointer is empty")
	Empty(t, tiNP, "time.Time is empty")
	Empty(t, TStruct{}, "struct with zero values is empty")
	Empty(t, TString(""), "empty aliased string is empty")
	Empty(t, sP, "ptr to nil value is empty")
	Empty(t, [1]int{}, "array is state")

	NotEmpty(t, "something", "Non Empty string is not empty")
	NotEmpty(t, errors.New("something"), "Non nil object is not empty")
	NotEmpty(t, []string{"something"}, "Non empty string array is not empty")
	NotEmpty(t, 1, "Non-zero int value is not empty")
	NotEmpty(t, true, "True value is not empty")
	NotEmpty(t, chWithValue, "Channel with values is not empty")
	NotEmpty(t, TStruct{x: 1}, "struct with initialized values is empty")
	NotEmpty(t, TString("abc"), "non-empty aliased string is empty")
	NotEmpty(t, xP, "ptr to non-nil value is not empty")
	NotEmpty(t, [1]int{42}, "array is not state")
}

func TestAssertContains(t *testing.T) {
	t.Parallel()
	Contains(t, "ab", "ab", "abc", "aabc", "babca")
}
