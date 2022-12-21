package assert

import (
	"bytes"
	"testing"
)

func TestAssertEqual(t *testing.T) {
	t.Parallel()
	Equal(nil, []string{}, []string{})
	Equal(t, []string{}, []string{})
	NotEqual(t, []string{"f"}, []string{}, "desc: %s", "abc")
	NotEqual(t, nil, []string{})
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
	var eface2 = new(error)
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

	var struct0 = new(bytes.Buffer)
	NotNil(t, struct0)
	var struct1 *bytes.Buffer
	Nil(t, struct1)
	var struct2 bytes.Buffer
	NotNil(t, struct2)
	var struct3 = &bytes.Buffer{}
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

	// var nil1 = (*int)(unsafe.Pointer(uintptr(0x0)))
	// Equal(t, true, IsNil(nil1))
}
