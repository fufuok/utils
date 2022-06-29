package utils

import (
	"bytes"
	"testing"
)

// Ref: gofiber/utils
func TestAssertEqual(t *testing.T) {
	t.Parallel()
	AssertEqual(nil, []string{}, []string{})
	AssertEqual(t, []string{}, []string{})
}

func TestAssertPanics(t *testing.T) {
	t.Parallel()
	var a []int
	AssertPanics(nil, "should panic when index out of range", func() {
		_ = a[1]
	})
	AssertPanics(t, "should panic when index out of range", func() {
		_ = a[1]
	})
}

func TestIsNil(t *testing.T) {
	t.Parallel()
	var map1 map[int]bool
	AssertEqual(t, true, IsNil(map1))
	map2 := make(map[int]bool)
	AssertEqual(t, false, IsNil(map2))

	var ch1 chan int
	AssertEqual(t, true, IsNil(ch1))
	ch2 := make(chan struct{})
	AssertEqual(t, false, IsNil(ch2))

	var slice1 []string
	AssertEqual(t, true, IsNil(slice1))
	slice2 := slice1
	AssertEqual(t, true, IsNil(slice2))
	slice1 = append(slice1, "")
	AssertEqual(t, false, IsNil(slice1))
	AssertEqual(t, true, IsNil(slice2))

	slice3 := make([]int, 0)
	AssertEqual(t, false, IsNil(slice3))
	slice4 := slice3
	AssertEqual(t, false, IsNil(slice4))
	slice3 = nil
	AssertEqual(t, true, IsNil(slice3))
	AssertEqual(t, false, IsNil(slice4))

	var iface1 interface{}
	AssertEqual(t, true, IsNil(iface1))
	iface1 = nil
	AssertEqual(t, true, IsNil(iface1))
	iface1 = map1
	AssertEqual(t, true, IsNil(iface1))
	iface1 = map2
	AssertEqual(t, false, IsNil(iface1))
	var iface2 interface{} = (*int)(nil)
	AssertEqual(t, true, IsNil(iface2))

	var eface1 error
	AssertEqual(t, true, IsNil(eface1))
	var eface2 = new(error)
	AssertEqual(t, false, IsNil(eface2))
	var iface3 interface{} = eface1
	AssertEqual(t, true, IsNil(iface3))

	var ptr *int
	AssertEqual(t, true, IsNil(ptr))

	var iface4 interface{} = ptr
	AssertEqual(t, true, IsNil(iface4))
	// AssertEqual(t, false, iface4 == nil) // go1.16.4

	var fun func(int) error
	AssertEqual(t, true, IsNil(fun))

	var struct0 = new(bytes.Buffer)
	AssertEqual(t, false, IsNil(struct0))
	var struct1 *bytes.Buffer
	AssertEqual(t, true, IsNil(struct1))
	var struct2 bytes.Buffer
	AssertEqual(t, false, IsNil(struct2))
	var struct3 = &bytes.Buffer{}
	AssertEqual(t, false, IsNil(struct3))
	var struct4 *struct{}
	AssertEqual(t, true, IsNil(struct4))
	var struct5 struct{}
	AssertEqual(t, false, IsNil(struct5))
	struct6 := struct{}{}
	AssertEqual(t, false, IsNil(struct6))

	var s string
	AssertEqual(t, false, IsNil(s))
	var iface5 interface{} = s
	AssertEqual(t, false, IsNil(iface5))

	// var nil1 = (*int)(unsafe.Pointer(uintptr(0x0)))
	// AssertEqual(t, true, IsNil(nil1))
}
