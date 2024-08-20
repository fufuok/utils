package utils

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

// GetDomain 检查并返回清除前后空白的域名
func GetDomain(name string) string {
	name = strings.TrimSpace(name)
	if CheckDomain(name) != nil {
		return ""
	}
	return name
}

// CheckDomain returns an error if the host name is not valid.
// See https://tools.ietf.org/html/rfc1034#section-3.5 and
// https://tools.ietf.org/html/rfc1123#section-2.
// Ref: chmike/domain
func CheckDomain(name string) error {
	if len(name) == 0 {
		return errors.New("domain name is empty")
	}
	if name[len(name)-1] == '.' {
		if len(name) > 254 {
			return fmt.Errorf("domain name length is %d, can't exceed 254 with a trailing dot", len(name))
		}
		name = name[:len(name)-1] // drop valid trailing dot
		if len(name) == 0 {
			return errors.New("domain name is a single dot")
		}
	} else if len(name) > 253 {
		return fmt.Errorf("domain name length is %d, can't exceed 253 without a trailing dot", len(name))
	}
	var l int
	for i := 0; i < len(name); i++ {
		b := name[i]
		if b == '.' {
			// check domain labels validity
			switch {
			case i == l:
				return fmt.Errorf("domain has an empty label at offset %d", l)
			case i-l > 63:
				return fmt.Errorf("domain byte length of label '%s' is %d, can't exceed 63", name[l:i], i-l)
			case name[l] == '-':
				return fmt.Errorf("domain label '%s' at offset %d begins with a hyphen", name[l:i], l)
			case name[i-1] == '-':
				return fmt.Errorf("domain label '%s' at offset %d ends with a hyphen", name[l:i], l)
			}
			l = i + 1
			continue
		}
		// test label character validity, note: tests are ordered by decreasing validity frequency
		if !(b >= 'a' && b <= 'z' || b >= '0' && b <= '9' || b == '-' || b >= 'A' && b <= 'Z') {
			// show the printable unicode character starting at byte offset i
			c, _ := utf8.DecodeRuneInString(name[i:])
			if c == utf8.RuneError {
				return fmt.Errorf("domain has invalid rune at offset %d", i)
			}
			return fmt.Errorf("domain has invalid character '%c' at offset %d", c, i)
		}
	}
	// check top level domain validity
	switch {
	case len(name)-l > 63:
		return fmt.Errorf("domain's top level domain '%s' has byte length %d, can't exceed 63", name[l:], len(name)-l)
	case name[l] == '-':
		return fmt.Errorf("domain's top level domain '%s' at offset %d begin with a hyphen", name[l:], l)
	case name[len(name)-1] == '-':
		return fmt.Errorf("domain's top level domain '%s' at offset %d ends with a hyphen", name[l:], l)
	case name[l] >= '0' && name[l] <= '9':
		return fmt.Errorf("domain's top level domain '%s' at offset %d begins with a digit", name[l:], l)
	}
	return nil
}
