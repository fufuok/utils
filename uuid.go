package utils

import (
	"encoding/hex"

	"github.com/fufuok/utils/base58"
)

// 随机 UUID
func UUIDString() string {
	return B2S(EncodeUUID(UUID()))
}

// 随机 UUID, 无短横线
func UUIDSimple() string {
	return hex.EncodeToString(UUID())
}

// 随机 UUID, 短版, base58
func UUIDShort() string {
	return base58.Encode(UUID())
}

// 随机 UUID, RFC4122, Version 4
func UUID() []byte {
	id := RandBytes(16)
	id[6] = (id[6] & 0x0f) | 0x40 // Version 4
	id[8] = (id[8] & 0x3f) | 0x80 // Variant is 10

	return id
}

// 编码 UUID
func EncodeUUID(id []byte) []byte {
	src := make([]byte, 16)
	copy(src, id)
	dst := make([]byte, 36)
	hex.Encode(dst, src[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], src[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], src[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], src[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], src[10:])

	return dst
}
