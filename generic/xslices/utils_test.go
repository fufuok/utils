//go:build go1.18
// +build go1.18

package xslices

import (
	"testing"

	"github.com/fufuok/utils/assert"
)

func TestToString(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	assert.Equal(t, "12345", ToString(xs, ""))
	assert.Equal(t, "1.2.3.4.5", ToString(xs, "."))
	assert.Equal(t, "1--2--3--4--5", ToString(xs, "--"))
	assert.Equal(t, "", ToString([]int{}, "."))
	assert.Equal(t, "0", ToString([]float64{0.0}, "."))
	assert.Equal(t, "0,1.2", ToString([]float64{0.0, 1.2}, ","))
}

func TestAverage(t *testing.T) {
	xs := []uint32{1, 2, 3, 4}
	assert.Equal(t, 2.5, Average(xs))
	assert.Equal(t, float64(0), Average([]float32{0.0}))
	assert.Equal(t, 0.6, Average([]float64{0.0, 1.2}))

	fs := []float64{0.0, 0.5, 0.5}
	assert.Equal(t, 0.33, Average(fs, 2))
	fs = []float64{0.0, 1, 1}
	assert.Equal(t, 0.667, Average(fs, 3))
}

func TestChunk(t *testing.T) {
	cases := []struct {
		ss   []int
		size int
		want [][]int
	}{
		{nil, 0, [][]int{}},
		{nil, 1, [][]int{}},
		{[]int{}, 1, [][]int{}},
		{[]int{1}, 1, [][]int{{1}}},
		{[]int{1}, 2, [][]int{{1}}},
		{[]int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}},
		{[]int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
	}
	for _, c := range cases {
		got := Chunk(c.ss, c.size)
		assert.Equal(t, c.want, got)
	}

	assert.Equal(t, [][]string{{"a", "b"}, {"c"}}, Chunk([]string{"a", "b", "c"}, 2))

	ss := []int{1, 2, 3}
	ssbak := []int{1, 2, 3}
	res := Chunk(ss, 2)
	assert.Equal(t, ssbak, ss)
	assert.Equal(t, 2, len(res))
	assert.Equal(t, 3, res[1][0])

	// 分组后元素的追加, 不影响原切片数据
	res[0] = append(res[0], 8)
	assert.Equal(t, []int{1, 2, 8}, res[0])
	assert.Equal(t, ssbak, ss)

	// 修改切片数据, 会影响原切片数据 (此类场景建议深拷贝原数据后传入函数)
	res[1][0] = 9
	assert.Equal(t, 9, res[1][0])
	assert.Equal(t, []int{1, 2, 9}, ss)
}
