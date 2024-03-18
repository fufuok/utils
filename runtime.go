package utils

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	_ "unsafe"
)

const (
	// PtrSize 4 on 32-bit systems, 8 on 64-bit.
	PtrSize = 4 << (^uintptr(0) >> 63)
)

var StackTraceBufferSize = 4 << 10

// RecoveryCallback 自定义恢复信息回调
type RecoveryCallback func(err interface{}, trace []byte)

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

// Recover 从 panic 中恢复并记录堆栈信息
func Recover(cb ...RecoveryCallback) {
	if err := recover(); err != nil {
		buf := make([]byte, StackTraceBufferSize)
		buf = buf[:runtime.Stack(buf, false)]
		if len(cb) > 0 && cb[0] != nil {
			cb[0](err, buf)
			return
		}
		log.Printf("Recovery: %v\n--- Traceback:\n%v\n", err, B2S(buf))
	}
}

// SafeGo 带 Recover 的 goroutine 运行
func SafeGo(fn func(), cb ...RecoveryCallback) {
	go func() {
		defer Recover(cb...)
		fn()
	}()
}

// SafeGoWithContext 带 Recover 的 goroutine 运行
func SafeGoWithContext(ctx context.Context, fn func(ctx context.Context), cb ...RecoveryCallback) {
	go func() {
		defer Recover(cb...)
		fn(ctx)
	}()
}

// SafeGoCommonFunc 带 Recover 的 goroutine 运行
func SafeGoCommonFunc(args interface{}, fn func(args interface{}), cb ...RecoveryCallback) {
	go func() {
		defer Recover(cb...)
		fn(args)
	}()
}

// WaitSignal 等待系统信号
// 默认捕获退出类信息
func WaitSignal(sig ...os.Signal) os.Signal {
	if len(sig) == 0 {
		sig = []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, sig...)
	return <-ch
}

// FastRand 随机数
//
//go:linkname FastRand runtime.fastrand
func FastRand() uint32

// FastRandn 等同于 FastRand() % n, 但更快
// See https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
//
//go:linkname FastRandn runtime.fastrandn
func FastRandn(n uint32) uint32

// CPUTicks CPU 时钟周期, 更高精度 (云服务器做伪随机数种子时慎用)
//
//go:linkname CPUTicks runtime.cputicks
func CPUTicks() int64

// NanoTime 返回当前时间 (以纳秒为单位)
//
//go:linkname NanoTime runtime.nanotime
func NanoTime() int64
