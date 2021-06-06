package utils

import (
	"os"
)

// IsExist 文件或目录是否存在
func IsExist(s string) bool {
	_, err := os.Stat(s)
	return err == nil || os.IsExist(err)
}

// IsFile 文件是否存在
func IsFile(s string) bool {
	info, err := os.Stat(s)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// IsDir 目录是否存在
func IsDir(s string) bool {
	info, err := os.Stat(s)
	if err != nil {
		return false
	}
	return info.IsDir()
}
