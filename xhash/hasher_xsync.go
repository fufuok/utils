//go:build go1.18
// +build go1.18

package xhash

/*
From https://github.com/puzpuzpuz/xsync

MIT License

Copyright (c) 2021 Andrey Pechkurov

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

import (
	"reflect"
	"unsafe"
)

// MakeHasher creates a fast hash function for the given comparable type.
// The only limitation is that the type should not contain interfaces inside
// based on runtime.typehash.
func MakeHasher[T comparable]() func(T) uint64 {
	var zero T
	seed := MakeSeed()
	if reflect.TypeOf(&zero).Elem().Kind() == reflect.Interface {
		return func(value T) uint64 {
			iValue := any(value)
			i := (*iface)(unsafe.Pointer(&iValue))
			return runtimeTypehash64(i.typ, i.word, seed)
		}
	}
	var iZero any = zero
	i := (*iface)(unsafe.Pointer(&iZero))
	return func(value T) uint64 {
		return runtimeTypehash64(i.typ, unsafe.Pointer(&value), seed)
	}
}

// MakeSeed creates a random seed.
func MakeSeed() uint64 {
	var s1 uint32
	for {
		s1 = runtimeFastrand()
		// We use seed 0 to indicate an uninitialized seed/hash,
		// so keep trying until we get a non-zero seed.
		if s1 != 0 {
			break
		}
	}
	s2 := runtimeFastrand()
	return uint64(s1)<<32 | uint64(s2)
}

// how interface is represented in memory
type iface struct {
	typ  uintptr
	word unsafe.Pointer
}

// same as runtimeTypehash, but always returns a uint64
// see: maphash.rthash function for details
func runtimeTypehash64(t uintptr, p unsafe.Pointer, seed uint64) uint64 {
	if unsafe.Sizeof(uintptr(0)) == 8 {
		return uint64(runtimeTypehash(t, p, uintptr(seed)))
	}

	lo := runtimeTypehash(t, p, uintptr(seed))
	hi := runtimeTypehash(t, p, uintptr(seed>>32))
	return uint64(hi)<<32 | uint64(lo)
}

//go:noescape
//go:linkname runtimeTypehash runtime.typehash
func runtimeTypehash(t uintptr, p unsafe.Pointer, h uintptr) uintptr

//go:noescape
//go:linkname runtimeFastrand runtime.fastrand
func runtimeFastrand() uint32
