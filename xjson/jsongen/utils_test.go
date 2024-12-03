package jsongen

import (
	"testing"

	"github.com/fufuok/utils"
	"github.com/fufuok/utils/pools/bytespool"
)

func TestAppendJSONString(t *testing.T) {
	bs := AppendJSONString([]byte("[1,2,"), "A\fB")
	bs = append(bs, "]"...)
	if string(bs) != `[1,2,"A\fB"]` {
		t.Fatalf(`[1,2,"A\fB"] != %s`, bs)
	}
	bs = AppendJSONString(nil, "A\nB")
	if string(bs) != `"A\nB"` {
		t.Fatalf(`"A\nB" != %s`, bs)
	}
	s := "A\rB"
	dst := bytespool.Make(len(s) + 2)
	ss := utils.B2S(AppendJSONString(dst, s))
	if ss != `"A\rB"` {
		t.Fatalf(`"A\rB" != %s`, ss)
	}
}
