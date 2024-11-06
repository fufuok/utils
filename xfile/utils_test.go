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

	ret, err = ReadLines("utils_nonono_test.go")
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(ret))
}

func TestHeadlines(t *testing.T) {
	ret, err := HeadLines("utils_test.go", -1)
	assert.Nil(t, err)
	assert.Equal(t, "package xfile", strings.TrimSpace(ret[0]))

	ret, err = HeadLines("utils_test.go", 3)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(ret))
	assert.Equal(t, "", strings.TrimSpace(ret[1]))
	assert.True(t, strings.Contains(ret[2], "import ("))

	ret, err = HeadLines("utils_nonono_test.go", -1)
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(ret))
}

func TestReadLinesOffsetN(t *testing.T) {
	ret, err := ReadLinesOffsetN("utils_test.go", 2, 1)
	assert.Nil(t, err)
	assert.True(t, strings.Contains(ret[0], "import ("))
}

func TestTaillines(t *testing.T) {
	ret, err := TailLines("utils_test.go", -1)
	assert.Nil(t, err)
	assert.Equal(t, "package xfile", strings.TrimSpace(ret[0]))

	ret, err = TailLines("utils_test.go", 2)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(ret))
	assert.Equal(t, "}", strings.TrimSpace(ret[0]))
	assert.Equal(t, "", strings.TrimSpace(ret[1]))

	ret, err = TailLines("utils_test.go", 2, true)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(ret))
	assert.True(t, strings.Contains(ret[0], "assert"))
	assert.Equal(t, "}", strings.TrimSpace(ret[1]))

	ret, err = TailLines("utils_nonono_test.go", 1)
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(ret))
}

func TestModTime(t *testing.T) {
	mtime := ModTime("utils_test.go")
	assert.NotEqual(t, time.Time{}, mtime)
	mtime = ModTime("not_exist__file.go")
	assert.Equal(t, time.Time{}, mtime)
}
