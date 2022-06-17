package utils

// NoCopy may be added to structs which must not be copied
// after the first use.
//
// See https://github.com/golang/go/issues/8005#issuecomment-190753527 for details.
// and also: https://stackoverflow.com/questions/52494458/nocopy-minimal-example
//
// Note that it must not be embedded, due to the Lock and Unlock methods.
type NoCopy struct{} //nolint:unused

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*NoCopy) Lock()   {} //nolint:unused
func (*NoCopy) Unlock() {} //nolint:unused
