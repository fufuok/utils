package utils

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/fufuok/utils/assert"
)

func TestSafeGo(t *testing.T) {
	var (
		err   interface{}
		trace []byte
		ctx   = context.Background()
	)
	rcb := func(e interface{}, s []byte) {
		err = e
		trace = s
	}
	SafeGo(testFn2, rcb)
	time.Sleep(5 * time.Millisecond)
	assert.Equal(t, "fn1", err)
	assert.Equal(t, true, bytes.Contains(trace, []byte("panic")))
	SafeGoWithContext(ctx, testFn3, rcb)
	time.Sleep(5 * time.Millisecond)
	assert.Equal(t, "fn1", err)
	assert.Equal(t, true, bytes.Contains(trace, []byte("panic")))
}

var (
	testFn1 = func() {
		panic("fn1")
	}
	testFn2 = func() {
		testFn1()
	}
	testFn3 = func(ctx context.Context) {
		testFn1()
	}
)
