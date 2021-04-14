package utils

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// 运行时路径
func CallPath() string {
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		return path.Dir(filename)
	}

	return RunPath()
}

// 实际程序所在路径
func RunPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		// 调用时工作路径
		dir, _ = os.Getwd()
		return dir
	}

	return strings.Replace(dir, "\\", "/", -1)
}
