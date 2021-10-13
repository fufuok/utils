package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"unsafe"
)

// CallPath 运行时路径, 编译目录
// 假如: mklink E:\tmp\linkapp.exe D:\Fufu\Test\abc\app.exe
// 执行: E:\tmp\linkapp.exe
// CallPath: E:\Go\src\github.com\fufuok\utils\tmp\osext
func CallPath() string {
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		return filepath.Clean(filepath.Dir(filename))
	}

	return RunPath()
}

// RunPath 实际程序所在目录
// RunPath: E:\tmp
func RunPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		// 调用时工作目录
		dir, _ = os.Getwd()
		return dir
	}

	return dir
}

// Executable 当前执行程序绝对路径
// true 时返回解析符号链接后的绝对路径
// Excutable: E:\tmp\linkapp.exe
// Excutable(true): D:\Fufu\Test\abc\app.exe
func Executable(evalSymlinks ...bool) string {
	exe, _ := os.Executable()
	if len(evalSymlinks) > 0 && evalSymlinks[0] {
		exe, _ = filepath.EvalSymlinks(exe)
	}

	return filepath.Clean(exe)
}

// ExecutableDir 当前执行程序所在目录
// true 时返回解析符号链接后的目录
// ExcutableDir: E:\tmp
// ExcutableDir(true): D:\Fufu\Test\abc
func ExecutableDir(evalSymlinks ...bool) string {
	return filepath.Dir(Executable(evalSymlinks...))
}

// FastRand 随机数
//go:linkname FastRand runtime.fastrand
func FastRand() uint32

// CPUTicks CPU 时钟周期, 更高精度 (云服务器做伪随机数种子时慎用)
//go:linkname CPUTicks runtime.cputicks
func CPUTicks() int64

// NanoTime 返回当前时间 (以纳秒为单位)
//go:linkname NanoTime runtime.nanotime
func NanoTime() int64

//go:noescape
//go:linkname memhash runtime.memhash
func memhash(p unsafe.Pointer, h, s uintptr) uintptr
