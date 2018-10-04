package kagemusha

import (
	"testing"
)

func return1() int {
	return 1
}

func TestReplaceFunction(t *testing.T) {
	Mock(return1).Return(2)
	if return1() != 2 {
		t.Fatal("function was not replaced")
	}
}
