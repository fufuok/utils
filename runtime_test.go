package utils

import (
	"bytes"
	"testing"
	"time"
)

func TestSafeGo(t *testing.T) {
	var (
		err   interface{}
		trace []byte
	)
	rcb := func(e interface{}, s []byte) {
		err = e
		trace = s
	}
	SafeGo(testFn2, rcb)
	time.Sleep(5 * time.Millisecond)
	AssertEqual(t, "fn1", err)
	AssertEqual(t, true, bytes.Contains(trace, []byte("panic")))
}

var (
	testFn1 = func() {
		panic("fn1")
	}
	testFn2 = func() {
		testFn1()
	}
)
