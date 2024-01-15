package xfile

import (
	"strings"
	"testing"
	"time"

	"github.com/fufuok/utils/assert"
)

func TestReadlines(t *testing.T) {
	ret, err := ReadLines("utils_test.go")
	assert.Nil(t, err)
	assert.Equal(t, "package xfile", strings.TrimSpace(ret[0]))
}

func TestReadLinesOffsetN(t *testing.T) {
	ret, err := ReadLinesOffsetN("utils_test.go", 2, 1)
	assert.Nil(t, err)
	assert.True(t, strings.Contains(ret[0], "import ("))
}

func TestModTime(t *testing.T) {
	mtime := ModTime("utils_test.go")
	assert.NotEqual(t, time.Time{}, mtime)
	mtime = ModTime("not_exist__file.go")
	assert.Equal(t, time.Time{}, mtime)
}
