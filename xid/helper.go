package xid

// NewString New().String()
func NewString() string {
	return New().String()
}

// NewBytes New().Bytes()
func NewBytes() []byte {
	return New().Bytes()
}
