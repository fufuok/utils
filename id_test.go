package utils_test

import (
	"testing"

	"github.com/fufuok/utils"
)

func TestID(t *testing.T) {
	a := utils.ID()
	b := utils.ID()
	utils.AssertEqual(t, b, a+1)
}
