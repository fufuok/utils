// Ref: dustin/go-humanize
package utils

import (
	"testing"
)

func TestParseHumanBytes(t *testing.T) {
	tests := []struct {
		in  string
		exp uint64
	}{
		{"42", 42},
		{"42MB", 42000000},
		{"42MiB", 44040192},
		{"42mb", 42000000},
		{"42mib", 44040192},
		{"42MIB", 44040192},
		{"42 MB", 42000000},
		{"42 MiB", 44040192},
		{"42 mb", 42000000},
		{"42 mib", 44040192},
		{"42 MIB", 44040192},
		{"42.5MB", 42500000},
		{"42.5MiB", 44564480},
		{"42.5 MB", 42500000},
		{"42.5 MiB", 44564480},
		// No need to say B
		{"42M", 42000000},
		{"42Mi", 44040192},
		{"42m", 42000000},
		{"42mi", 44040192},
		{"42MI", 44040192},
		{"42 M", 42000000},
		{"42 Mi", 44040192},
		{"42 m", 42000000},
		{"42 mi", 44040192},
		{"42 MI", 44040192},
		{"42.5M", 42500000},
		{"42.5Mi", 44564480},
		{"42.5 M", 42500000},
		{"42.5 Mi", 44564480},
		{"1,005.03 MB", 1005030000},
		// Large testing, breaks when too much larger than this.
		{"12.5 EB", uint64(12.5 * float64(EByte))},
		{"12.5 E", uint64(12.5 * float64(EByte))},
		{"12.5 EiB", uint64(12.5 * float64(EiByte))},
	}

	for _, p := range tests {
		got, err := ParseHumanBytes(p.in)
		if err != nil {
			t.Errorf("Couldn't parse %v: %v", p.in, err)
		}
		if got != p.exp {
			t.Errorf("Expected %v for %v, got %v",
				p.exp, p.in, got)
		}
	}
}

func TestParseHumanBytesErrors(t *testing.T) {
	got, err := ParseHumanBytes("84 JB")
	if err == nil {
		t.Errorf("Expected error, got %v", got)
	}
	got, err = ParseHumanBytes("")
	if err == nil {
		t.Errorf("Expected error parsing nothing")
	}
	got, err = ParseHumanBytes("16 EiB")
	if err == nil {
		t.Errorf("Expected error, got %v", got)
	}
}

func TestHumanBytes(t *testing.T) {
	for _, v := range []struct {
		title    string
		actual   string
		expected string
	}{
		{"HumanBytes(0)", HumanBytes(0), "0 B"},
		{"HumanBytes(1)", HumanBytes(1), "1 B"},
		{"HumanBytes(803)", HumanBytes(803), "803 B"},
		{"HumanBytes(999)", HumanBytes(999), "999 B"},

		{"HumanBytes(1024)", HumanBytes(1024), "1.0 kB"},
		{"HumanBytes(9999)", HumanBytes(9999), "10 kB"},
		{"HumanBytes(1MB - 1)", HumanBytes(MByte - Byte), "1000 kB"},

		{"HumanBytes(1MB)", HumanBytes(1024 * 1024), "1.0 MB"},
		{"HumanBytes(1GB - 1K)", HumanBytes(GByte - KByte), "1000 MB"},

		{"HumanBytes(5.5GB)", HumanBytes(5.5 * GByte), "5.5 GB"},
		{"HumanBytes(1GB)", HumanBytes(GByte), "1.0 GB"},
		{"HumanBytes(1TB - 1M)", HumanBytes(TByte - MByte), "1000 GB"},
		{"HumanBytes(10MB)", HumanBytes(9999 * 1000), "10 MB"},

		{"HumanBytes(1TB)", HumanBytes(TByte), "1.0 TB"},
		{"HumanBytes(1PB - 1T)", HumanBytes(PByte - TByte), "999 TB"},

		{"HumanBytes(1PB)", HumanBytes(PByte), "1.0 PB"},
		{"HumanBytes(1PB - 1T)", HumanBytes(EByte - PByte), "999 PB"},

		{"HumanBytes(1EB)", HumanBytes(EByte), "1.0 EB"},

		{"HumanIBytes(0)", HumanIBytes(0), "0 B"},
		{"HumanIBytes(1)", HumanIBytes(1), "1 B"},
		{"HumanIBytes(803)", HumanIBytes(803), "803 B"},
		{"HumanIBytes(1023)", HumanIBytes(1023), "1023 B"},

		{"HumanIBytes(1024)", HumanIBytes(1024), "1.0 KiB"},
		{"HumanIBytes(1MB - 1)", HumanIBytes(MiByte - IByte), "1024 KiB"},

		{"HumanIBytes(1MB)", HumanIBytes(1024 * 1024), "1.0 MiB"},
		{"HumanIBytes(1GB - 1K)", HumanIBytes(GiByte - KiByte), "1024 MiB"},

		{"HumanIBytes(5.5GiB)", HumanIBytes(5.5 * GiByte), "5.5 GiB"},
		{"HumanIBytes(1GB)", HumanIBytes(GiByte), "1.0 GiB"},
		{"HumanIBytes(1TB - 1M)", HumanIBytes(TiByte - MiByte), "1024 GiB"},

		{"HumanIBytes(1TB)", HumanIBytes(TiByte), "1.0 TiB"},
		{"HumanIBytes(1PB - 1T)", HumanIBytes(PiByte - TiByte), "1023 TiB"},

		{"HumanIBytes(1PB)", HumanIBytes(PiByte), "1.0 PiB"},
		{"HumanIBytes(1PB - 1T)", HumanIBytes(EiByte - PiByte), "1023 PiB"},

		{"HumanIBytes(1EiB)", HumanIBytes(EiByte), "1.0 EiB"},
	} {
		AssertEqual(t, v.expected, v.actual, v.title)
	}
}

func BenchmarkParseBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseHumanBytes("16.5 GB")
	}
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HumanBytes(16.5 * GByte)
	}
}
