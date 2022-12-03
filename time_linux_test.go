//go:build !windows
// +build !windows

package utils

import (
	"testing"
	"time"
)

// 忽略 Github Windows 环境
func TestWaitNextSecond(t *testing.T) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			go func() {
				WaitNextSecond()
				t.Log(time.Now().Format("15:04:05.000"))
				s := time.Now().Format("05.000")
				if s[4] != '0' {
					t.Errorf("expect: `0`, got: `%s`", string(s[4]))
				}
			}()
		}
		WaitNextSecond()
		t.Log(time.Now().Format("15:04:05.000"))
	}
}
