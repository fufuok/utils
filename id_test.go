package utils_test

import (
	"testing"

	"github.com/fufuok/utils"
	"github.com/fufuok/utils/assert"
)

func TestID(t *testing.T) {
	a := utils.ID()
	b := utils.ID()
	assert.Equal(t, b, a+1)
}
