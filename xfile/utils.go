package xfile

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
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

// ResetDir 清除并重建空目录
func ResetDir(dirPath string) error {
	if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
		if err := os.RemoveAll(dirPath); err != nil {
			return err
		}
	}
	return os.MkdirAll(dirPath, 0755)
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

// CopyDir 目录拷贝
func CopyDir(srcDir, dstDir string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dstDir, relPath)
		if info.IsDir() {
			return os.MkdirAll(dstPath, os.ModePerm)
		}
		return CopyFile(path, dstPath)
	})
}

// CopyFile 文件拷贝
func CopyFile(srcFile, dstFile string) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer func() {
		_ = src.Close()
	}()

	dst, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer func() {
		_ = dst.Close()
	}()

	_, err = io.Copy(dst, src)
	return err
}

// ZipDir 目录打包为 zip 文件
func ZipDir(srcDir, zipFilePath string) error {
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer func() {
		_ = zipFile.Close()
	}()

	zipWriter := zip.NewWriter(zipFile)
	defer func() {
		_ = zipWriter.Close()
	}()

	err = filepath.Walk(srcDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to access file: %w", err)
		}

		// 获取文件在 zip 中的相对路径
		relPath, err := filepath.Rel(srcDir, filePath)
		if err != nil {
			return fmt.Errorf("failed to get relative path for file: %w", err)
		}

		if relPath == "." {
			return nil
		}

		// 创建 zip 文件中的文件或目录
		if info.IsDir() {
			// 添加目录条目
			if _, err := zipWriter.Create(relPath + "/"); err != nil {
				return fmt.Errorf("failed to create directory entry for %s: %w", relPath, err)
			}
			return nil
		}

		// 添加文件条目
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer func() {
			_ = file.Close()
		}()

		w, err := zipWriter.Create(relPath)
		if err != nil {
			return fmt.Errorf("failed to create file entry for %s: %w", relPath, err)
		}
		if _, err := io.Copy(w, file); err != nil {
			return fmt.Errorf("failed to copy file: %w", err)
		}
		return nil
	})
	return err
}

// UnzipDir 解压 zip 到目录
func UnzipDir(zipFile, dstDir string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return fmt.Errorf("failed to open the zip file: %w", err)
	}
	defer func() {
		_ = reader.Close()
	}()

	for _, file := range reader.File {
		if strings.Contains(file.Name, "..") {
			return fmt.Errorf("illegal file path in zip: %v", file.Name)
		}

		fullPath := filepath.Join(dstDir, file.Name)

		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(fullPath, file.Mode()); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
			continue
		}

		if err := UnzipFile(file, fullPath); err != nil {
			return err
		}
	}
	return nil
}

// UnzipFile 解压单个文件
func UnzipFile(zipFile *zip.File, dstFile string) error {
	src, err := zipFile.Open()
	if err != nil {
		return fmt.Errorf("failed to open zip file: %w", err)
	}
	defer func() {
		_ = src.Close()
	}()

	dst, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zipFile.Mode())
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		_ = dst.Close()
	}()

	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	return nil
}
