package utils

// Pad 填充字符串到指定长度, 同 Python3: 'str'.center()
func Pad(s, pad string, n int) string {
	need := n - len(s)
	if need <= 0 {
		return s
	}

	bs := make([]byte, n)
	half := need - need>>1

	// 一半填充到左边
	// 奇数时左边比右边多: 1
	copyPad(bs, pad, half)

	m := copy(bs[half:], s)
	m += half

	// 一半填充到右边
	copyPad(bs[m:], pad, n-m)

	return B2S(bs)
}

// LeftPad 从左填充字符串到指定长度
func LeftPad(s, pad string, n int) string {
	need := n - len(s)
	if need <= 0 {
		return s
	}

	bs := make([]byte, n)
	copyPad(bs, pad, need)
	copy(bs[need:], s)

	return B2S(bs)
}

// RightPad 从右填充字符串到指定长度
func RightPad(s, pad string, n int) string {
	need := n - len(s)
	if need <= 0 {
		return s
	}

	bs := make([]byte, n)
	l := copy(bs, s)
	copyPad(bs[l:], pad, need)

	return B2S(bs)
}

// PadBytes 填充到指定长度
func PadBytes(s, pad []byte, n int) []byte {
	need := n - len(s)
	if need <= 0 {
		return s
	}

	bs := make([]byte, n)
	half := need - need>>1

	// 一半填充到左边
	// 奇数时左边比右边多: 1
	copyPadBytes(bs, pad, half)

	m := copy(bs[half:], s)
	m += half

	// 一半填充到右边
	copyPadBytes(bs[m:], pad, n-m)

	return bs
}

// LeftPadBytes 从左填充到指定长度
func LeftPadBytes(b, pad []byte, n int) []byte {
	need := n - len(b)
	if need <= 0 {
		return b
	}

	bs := make([]byte, n)
	copyPadBytes(bs, pad, need)
	copy(bs[need:], b)

	return bs
}

// RightPadBytes 从右填充到指定长度
func RightPadBytes(b, pad []byte, n int) []byte {
	need := n - len(b)
	if need <= 0 {
		return b
	}

	bs := make([]byte, n)
	l := copy(bs, b)
	copyPadBytes(bs[l:], pad, need)

	return bs
}

// 重复填充到指定长度
func copyPad(bs []byte, pad string, need int) {
	if need <= 0 {
		return
	}
	n := copy(bs[:need], pad)
	for n < need {
		copy(bs[n:need], bs[:n])
		n = n << 1
	}
}

func copyPadBytes(bs, pad []byte, need int) {
	if need <= 0 {
		return
	}
	n := copy(bs[:need], pad)
	for n < need {
		copy(bs[n:need], bs[:n])
		n = n << 1
	}
}
