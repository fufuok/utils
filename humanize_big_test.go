package utils

import (
	"math/big"
	"testing"
)

// Ref: dustin/go-humanize
func TestHumanBigByteParsing(t *testing.T) {
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
		{"1,005.03 Mbps", 1005030000},
		// Large testing, breaks when too much larger than
		// this.
		{"12.5 EB", uint64(12.5 * float64(EByte))},
		{"12.5 E", uint64(12.5 * float64(EByte))},
		{"12.5 Ebps", uint64(12.5 * float64(EByte))},
		{"12.5 EiB", uint64(12.5 * float64(EiByte))},
	}

	for _, p := range tests {
		got, err := ParseHumanBigBytes(p.in)
		if err != nil {
			t.Errorf("Couldn't parse %v: %v", p.in, err)
		} else {
			if got.Uint64() != p.exp {
				t.Errorf("Expected %v for %v, got %v",
					p.exp, p.in, got)
			}
		}
	}
}

func TestHumanBigByteErrors(t *testing.T) {
	got, err := ParseHumanBigBytes("84 JB")
	if err == nil {
		t.Errorf("Expected error, got %v", got)
	}
	got, err = ParseHumanBigBytes("")
	if err == nil {
		t.Errorf("Expected error parsing nothing")
	}
}

func bbyte(in uint64) string {
	return HumanBigBytes((&big.Int{}).SetUint64(in))
}

func bibyte(in uint64) string {
	return HumanBigIBytes((&big.Int{}).SetUint64(in))
}

func TestHumanBigBytes(t *testing.T) {
	for _, v := range []struct {
		title    string
		actual   string
		expected string
	}{
		{"bytes(0)", bbyte(0), "0 B"},
		{"bytes(1)", bbyte(1), "1 B"},
		{"bytes(803)", bbyte(803), "803 B"},
		{"bytes(999)", bbyte(999), "999 B"},

		{"bytes(1024)", bbyte(1023), "1.0 KB"},
		{"bytes(1024)", bbyte(1024), "1.0 KB"},
		{"bytes(1MB - 1)", bbyte(MByte - Byte), "1000 KB"},

		{"bytes(1MB)", bbyte(1024 * 1024), "1.0 MB"},
		{"bytes(1GB - 1K)", bbyte(GByte - KByte), "1000 MB"},

		{"bytes(1GB)", bbyte(GByte), "1.0 GB"},
		{"bytes(1TB - 1M)", bbyte(TByte - MByte), "1000 GB"},

		{"bytes(1TB)", bbyte(TByte), "1.0 TB"},
		{"bytes(1PB - 1T)", bbyte(PByte - TByte), "999 TB"},

		{"bytes(1PB)", bbyte(PByte), "1.0 PB"},
		{"bytes(1PB - 1T)", bbyte(EByte - PByte), "999 PB"},

		{"bytes(1EB)", bbyte(EByte), "1.0 EB"},
		// Overflows.
		// {"bytes(1EB - 1P)", Bytes((KByte*EByte)-PByte), "1023EB"},

		{"bytes(0)", bibyte(0), "0 B"},
		{"bytes(1)", bibyte(1), "1 B"},
		{"bytes(803)", bibyte(803), "803 B"},
		{"bytes(1023)", bibyte(1023), "1023 B"},

		{"bytes(1024)", bibyte(1024), "1.0 KiB"},
		{"bytes(1MB - 1)", bibyte(MiByte - IByte), "1024 KiB"},

		{"bytes(1MB)", bibyte(1024 * 1024), "1.0 MiB"},
		{"bytes(1GB - 1K)", bibyte(GiByte - KiByte), "1024 MiB"},

		{"bytes(1GB)", bibyte(GiByte), "1.0 GiB"},
		{"bytes(1TB - 1M)", bibyte(TiByte - MiByte), "1024 GiB"},

		{"bytes(1TB)", bibyte(TiByte), "1.0 TiB"},
		{"bytes(1PB - 1T)", bibyte(PiByte - TiByte), "1023 TiB"},

		{"bytes(1PB)", bibyte(PiByte), "1.0 PiB"},
		{"bytes(1PB - 1T)", bibyte(EiByte - PiByte), "1023 PiB"},

		{"bytes(1EiB)", bibyte(EiByte), "1.0 EiB"},
		// Overflows.
		// {"bytes(1EB - 1P)", bibyte((KIByte*EIByte)-PiByte), "1023EB"},

		{"bytes(5.5GiB)", bibyte(5.5 * GiByte), "5.5 GiB"},
		{"bytes(5.5GB)", bbyte(5.5 * GByte), "5.5 GB"},

		{"HumanBigKbps(1023)", HumanBigKbps((&big.Int{}).SetUint64(1023)), "1.0 Kbps"},
		{"HumanBigKbps(1024)", HumanBigKbps((&big.Int{}).SetUint64(1024)), "1.0 Kbps"},
		{"HumanBigKbps(5.5GB)", HumanBigKbps((&big.Int{}).SetUint64(5.5 * GByte)), "5.5 Gbps"},
	} {
		AssertEqual(t, v.expected, v.actual, v.title)
	}
}

func TestHumanVeryBigBytes(t *testing.T) {
	b, _ := (&big.Int{}).SetString("15347691069326346944512", 10)
	s := HumanBigBytes(b)
	if s != "15 ZB" {
		t.Errorf("Expected 15 ZB, got %v", s)
	}
	s = HumanBigIBytes(b)
	if s != "13 ZiB" {
		t.Errorf("Expected 13 ZiB, got %v", s)
	}

	b, _ = (&big.Int{}).SetString("15716035654990179271180288", 10)
	s = HumanBigBytes(b)
	if s != "16 YB" {
		t.Errorf("Expected 16 YB, got %v", s)
	}
	s = HumanBigIBytes(b)
	if s != "13 YiB" {
		t.Errorf("Expected 13 YiB, got %v", s)
	}
}

func TestHumanVeryVeryBigBytes(t *testing.T) {
	b, _ := (&big.Int{}).SetString("16093220510709943573688614912", 10)
	s := HumanBigKbps(b)
	if s != "16093 Ybps" {
		t.Errorf("Expected 16093 Ybps, got %v", s)
	}
	s = HumanBigBytes(b)
	if s != "16093 YB" {
		t.Errorf("Expected 16093 YB, got %v", s)
	}
	s = HumanBigIBytes(b)
	if s != "13312 YiB" {
		t.Errorf("Expected 13312 YiB, got %v", s)
	}
}

func TestHumanParseVeryBig(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"16 ZB", "16000000000000000000000"},
		{"16 ZiB", "18889465931478580854784"},
		{"16.5 ZB", "16500000000000000000000"},
		{"16.5 ZiB", "19479761741837286506496"},
		{"16 Z", "16000000000000000000000"},
		{"16 Zi", "18889465931478580854784"},
		{"16.5 Z", "16500000000000000000000"},
		{"16.5 Zi", "19479761741837286506496"},

		{"16 YB", "16000000000000000000000000"},
		{"16 YiB", "19342813113834066795298816"},
		{"16.5 YB", "16500000000000000000000000"},
		{"16.5 YiB", "19947276023641381382651904"},
		{"16 Y", "16000000000000000000000000"},
		{"16 Yi", "19342813113834066795298816"},
		{"16.5 Y", "16500000000000000000000000"},
		{"16.5 Yi", "19947276023641381382651904"},

		{"16.5 Ybps", "16500000000000000000000000"},
	}

	for _, test := range tests {
		x, err := ParseHumanBigBytes(test.in)
		if err != nil {
			t.Errorf("Error parsing %q: %v", test.in, err)
			continue
		}

		if x.String() != test.out {
			t.Errorf("Expected %q for %q, got %v", test.out, test.in, x)
		}
	}
}

func BenchmarkParseHumanBigBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseHumanBigBytes("16.5 Z")
	}
}

func BenchmarkHumanBigBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bibyte(16.5 * GByte)
	}
}
