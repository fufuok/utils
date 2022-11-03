//go:build go1.18
// +build go1.18

package xsync

import (
	"hash/maphash"
	"strconv"
	"testing"
)

func TestHashSeedString(t *testing.T) {
	const numEntries = 1000
	c := NewIntegerMapOf[uint64, int]()
	seed := maphash.MakeSeed()
	for i := 0; i < numEntries; i++ {
		if _, ok := c.LoadOrStore(HashSeedString(seed, strconv.Itoa(i)), i); ok {
			t.Fatal("value was not expected")
		}
	}
	if c.Size() != numEntries {
		t.Fatalf("expect count of 10000, but got: %d", c.Size())
	}
}

func TestHashSeedUint64(t *testing.T) {
	const numEntries = 1000
	c := NewIntegerMapOf[uint64, int]()
	seed := maphash.MakeSeed()
	for i := 0; i < numEntries; i++ {
		if _, ok := c.LoadOrStore(HashSeedUint64(seed, uint64(i)), i); ok {
			t.Fatal("value was not expected")
		}
	}
	if c.Size() != numEntries {
		t.Fatalf("expect count of 10000, but got: %d", c.Size())
	}
}
