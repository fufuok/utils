package utils

import (
	"encoding/json"
	"strconv"
	"sync/atomic"
)

const (
	False = 0
	True  = 1
)

// Bool is an atomic type-safe wrapper for bool values
type Bool uint32

func NewBool(v bool) Bool {
	var b Bool
	b.Store(v)
	return b
}

func NewTrue() Bool {
	var b Bool
	b.StoreTrue()
	return b
}

func NewFalse() Bool {
	var b Bool
	b.StoreFalse()
	return b
}

func (a *Bool) Store(v bool) {
	atomic.StoreUint32(a.ptr(), b2i(v))
}

func (a *Bool) StoreTrue() {
	atomic.StoreUint32(a.ptr(), True)
}

func (a *Bool) StoreFalse() {
	atomic.StoreUint32(a.ptr(), False)
}

func (a *Bool) Load() bool {
	return atomic.LoadUint32(a.ptr()) == True
}

func (a *Bool) Swap(v bool) (old bool) {
	return atomic.SwapUint32(a.ptr(), b2i(v)) == True
}

func (a *Bool) CAS(old, v bool) bool {
	return atomic.CompareAndSwapUint32(a.ptr(), b2i(old), b2i(v))
}

// Toggle atomically negates the Boolean and returns the previous value
func (a *Bool) Toggle() (old bool) {
	for {
		old := a.Load()
		if a.CAS(old, !old) {
			return old
		}
	}
}

func (a *Bool) String() string {
	return strconv.FormatBool(a.Load())
}

func (a *Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Load())
}

func (a *Bool) UnmarshalJSON(b []byte) error {
	var v bool
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	a.Store(v)
	return nil
}

func (a *Bool) ptr() *uint32 {
	return (*uint32)(a)
}

func b2i(v bool) uint32 {
	if v {
		return True
	}
	return False
}
