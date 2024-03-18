// Copyright 2022 Changkun Ou <changkun.de>. All rights reserved.
// Use of this source code is governed by a GPLv3 license that
// can be found in the LICENSE file.

package sched_test

import (
	"sync/atomic"
	"testing"

	"github.com/fufuok/utils/sched"
)

func TestSched(t *testing.T) {
	s := sched.New()
	s.Add(10)
	sum := uint32(0)
	for i := 0; i < 10; i++ {
		s.RunWithArgs(func(n ...interface{}) {
			atomic.AddUint32(&sum, uint32(n[0].(int)))
			atomic.AddUint32(&sum, -uint32(n[1].(int)))
		}, i, i)
	}
	s.Wait()
	if sum != 0 {
		t.Fatalf("wrong sum, expect: %d, want %d", 0, sum)
	}

	s.Add(10)
	sum = uint32(0)
	for i := 0; i < 10; i++ {
		ii := uint32(i)
		s.RunWithArgs(func(_ ...interface{}) {
			atomic.AddUint32(&sum, ii)
		})
	}
	s.Wait()
	if sum != 45 {
		t.Fatalf("wrong sum, expect: %d, want %d", 45, sum)
	}

	if s.Running() != 0 {
		t.Fatalf("wrong counter inside the pool")
	}

	s.Add(10)
	sum = uint32(0)
	for i := 0; i < 10; i++ {
		ii := uint32(i)
		s.Run(func() {
			atomic.AddUint32(&sum, ii)
		})
	}
	s.Wait()
	if sum != 45 {
		t.Fatalf("wrong sum, expect: %d, want %d", 45, sum)
	}

	if s.Running() != 0 {
		t.Fatalf("wrong counter inside the pool")
	}
	s.Release()

	s = sched.New(sched.Workers(4), sched.Queues(2))
	s.Add(10)
	sum = uint32(0)
	for i := 0; i < 10; i++ {
		ii := uint32(i)
		s.Run(func() {
			atomic.AddUint32(&sum, ii)
		})
	}
	s.WaitAndRelease()
	if sum != 45 {
		t.Fatalf("wrong sum, expect: %d, want %d", 45, sum)
	}

	if s.IsRunning() {
		t.Fatalf("wrong counter inside the pool")
	}
}

var f = func() {
	p := 0
	for i := 0; i < 0; i++ {
		p += i
	}
}

func BenchmarkSched(b *testing.B) {
	l := sched.New()
	b.Run("no-args", func(b *testing.B) {
		l.Add(b.N)
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			l.Run(f)
		}
		l.Wait()
	})
	b.Run("with-args", func(b *testing.B) {
		l.Add(b.N)
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			l.RunWithArgs(func(x ...interface{}) {
				_ = x[0]
			}, 42)
		}
		l.Wait()
	})
	l.Release()
}
