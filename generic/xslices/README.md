<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# xslices

```go
import "github.com/fufuok/utils/generic/xslices"
```

## Index

- [func Deduplication[E comparable](s []E) []E](<#func-deduplication>)
- [func Filter[E any, S ~[]E](s S, pred func(E) bool) S](<#func-filter>)
- [func Merge[E any](s []E, ss ...[]E) []E](<#func-merge>)


## func Deduplication

```go
func Deduplication[E comparable](s []E) []E
```

Deduplication removes repeatable elements from s.

## func Filter

```go
func Filter[E any, S ~[]E](s S, pred func(E) bool) S
```

Filter removes any elements from s for which pred\(element\) is false.

## func Merge

```go
func Merge[E any](s []E, ss ...[]E) []E
```

Merge 浅拷贝合并多个切片, 不影响原切片



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
