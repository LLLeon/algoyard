package bit

import (
	"testing"
)

func TestDuplicateIntI(t *testing.T) {
	l := []int{1, 1, 2, 2, 3, 4, 5, 5}

	res := DuplicateIntI(l)
	t.Log(res)
}

func TestDuplicateIntII(t *testing.T) {
	l := []int{1, 1, 2, 2, 3, 4, 5, 5}
	res := DuplicateIntII(l)
	t.Log(res)
}
