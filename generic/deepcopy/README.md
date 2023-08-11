deepCopy
========

*forked from smallnest/deepcopy*

[![License](https://img.shields.io/:license-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![GoDoc](https://godoc.org/github.com/smallnest/deepcopy?status.png)](http://godoc.org/github.com/smallnest/deepcopy)  ![build status](https://github.com/smallnest/deepcopy/actions/workflows/go.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/smallnest/deepcopy)](https://goreportcard.com/report/github.com/smallnest/deepcopy)


DeepCopy makes deep copies of things: unexported field values are not copied.

- Support Embed type
- Support Pointer deep copy
- Support Map
- Support Slice
- Support Interface
- Support Channel;

Forked from [mohae/deepcopy](https://github.com/mohae/deepcopy) and add generic support.

## Usage
```go
   cpy := deepcopy.Copy[T](orig)
```
 
