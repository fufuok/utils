package utils

import (
	"net/url"
	"testing"
)

func TestReplaceHost(t *testing.T) {
	u, _ := url.Parse("postgres://user:pass@host.com:5432/path?k=v#f")
	for _, v := range []struct {
		a   string
		b   string
		out string
	}{
		{"使用时应传入 (*url.URL).Host", "Fufu\n 中　文", "Fufu\n 中　文"},
		{"ff.cn:123", u.Host, "host.com:123"},
		{"", "", ""},
		{"a.cn:77", "b.cn:88", "b.cn:77"},
		{"a.cn", "b.cn:88", "b.cn"},
		{"a.cn:77", "b.cn", "b.cn:77"},
		{"a.cn", "b.cn", "b.cn"},
		{"", "b.cn", "b.cn"},
		{"a.cn", "", ""},
		{"[fe77::1]:77", "[fe88::1]:88", "[fe88::1]:77"},
		{"[fe77::1]", "[fe88::1]:88", "[fe88::1]"},
		{"[fe77::1]:77", "[fe88::1]", "[fe88::1]:77"},
		{"[fe77::1]", "[fe88::1]", "[fe88::1]"},
	} {
		AssertEqual(t, v.out, ReplaceHost(v.a, v.b))
	}
}
