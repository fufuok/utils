package utils

import (
	"testing"
)

func TestUUIDString(t *testing.T) {
	m := make(map[string]bool)
	for i := 0; i < 10000; i++ {
		id := UUIDString()
		if m[id] {
			t.Error("duplicated UUID:", id)
		}
		m[id] = true
		AssertEqual(t, uint8(4), UUID()[6]>>4)
		AssertEqual(t, uint8(0x80), UUID()[8]&0xc0)
	}
}
