package xbinary

import (
	"unsafe"
)

// IsLittleEndian 判断系统是否为小端序
func IsLittleEndian() bool {
	i := uint16(1)
	return (*(*[2]byte)(unsafe.Pointer(&i)))[0] == 1
}

// SwapEndianUint32 大小端交换
func SwapEndianUint32(val uint32) uint32 {
	return (val&0xff000000)>>24 | (val&0x00ff0000)>>8 | (val&0x0000ff00)<<8 | (val&0x000000ff)<<24
}
