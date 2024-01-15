package xfile

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/fufuok/utils"
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

// ReadFile reads contents from a file
func ReadFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return utils.B2S(content), nil
}

// ReadLines reads contents from a file and splits them by new lines.
// A convenience wrapper to ReadLinesOffsetN(filename, 0, -1).
func ReadLines(filename string) ([]string, error) {
	return ReadLinesOffsetN(filename, 0, -1)
}

// ReadLinesOffsetN reads contents from file and splits them by new line.
// The offset tells at which line number to start.
// The count determines the number of lines to read (starting from offset):
// n >= 0: at most n lines
// n < 0: whole file
// Ref: gopsutil
func ReadLinesOffsetN(filename string, offset uint, n int) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{""}, err
	}
	defer func() {
		_ = f.Close()
	}()

	var ret []string
	r := bufio.NewReader(f)
	for i := 0; i < n+int(offset) || n < 0; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF && len(line) > 0 {
				ret = append(ret, strings.Trim(line, "\n"))
			}
			break
		}
		if i < int(offset) {
			continue
		}
		ret = append(ret, strings.Trim(line, "\n"))
	}
	return ret, nil
}

// ModTime 文件最后修改时间
func ModTime(filename string) time.Time {
	info, err := os.Stat(filename)
	if err != nil {
		return time.Time{}
	}
	return info.ModTime()
}
