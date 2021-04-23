//go:build gojson
// +build gojson

package json

import (
	gojson "encoding/json"
)

var (
	Marshal       = gojson.Marshal
	Unmarshal     = gojson.Unmarshal
	MarshalIndent = gojson.MarshalIndent
	NewDecoder    = gojson.NewDecoder
	NewEncoder    = gojson.NewEncoder
)
