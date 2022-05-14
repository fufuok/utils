package utils

import (
	"encoding/json"
	"strconv"
	"sync/atomic"
)

// A Bool is an atomic boolean value.
// The zero value is false.
type Bool struct {
	_ NoCopy
	_ NoCmp
	v uint32
}

func NewBool(val bool) *Bool {
	x := &Bool{}
	x.Store(val)
	return x
}

func NewTrue() *Bool {
	x := &Bool{}
	x.StoreTrue()
	return x
}

func NewFalse() *Bool {
	x := &Bool{}
	x.StoreFalse()
	return x
}

// Load atomically loads and returns the value stored in x.
func (x *Bool) Load() bool {
	return atomic.LoadUint32(&x.v) != 0
}

// Store atomically stores val into x.
func (x *Bool) Store(val bool) {
	atomic.StoreUint32(&x.v, b32(val))
}

func (x *Bool) StoreTrue() {
	x.Store(true)
}

func (x *Bool) StoreFalse() {
	x.Store(false)
}

// Swap atomically stores new into x and returns the previous value.
func (x *Bool) Swap(new bool) (old bool) {
	return atomic.SwapUint32(&x.v, b32(new)) != 0
}

// CompareAndSwap executes the compare-and-swap operation for the boolean value x.
func (x *Bool) CompareAndSwap(old, new bool) (swapped bool) {
	return atomic.CompareAndSwapUint32(&x.v, b32(old), b32(new))
}

func (x *Bool) CAS(old, new bool) bool {
	return x.CompareAndSwap(old, new)
}

// Toggle atomically negates the Boolean and returns the previous value
func (x *Bool) Toggle() (old bool) {
	for {
		old := x.Load()
		if x.CompareAndSwap(old, !old) {
			return old
		}
	}
}

func (x *Bool) String() string {
	return strconv.FormatBool(x.Load())
}

func (x *Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Load())
}

func (x *Bool) UnmarshalJSON(b []byte) error {
	var v bool
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	x.Store(v)
	return nil
}

// b32 returns a uint32 0 or 1 representing b.
func b32(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}
