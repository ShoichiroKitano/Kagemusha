package kagemusha

import (
	"testing"
)

func return1() int {
	return 1
}

func return2() int {
	return 2
}

//func TestReplaceFunction(t *testing.T) {
//	Mock(return1).Return(2)
//	if return1() != 2 {
//		t.Fatal("function was not replaced")
//	}
//}

func TestUnmockFunction(t *testing.T) {
	mock := Mock(return2)
	mock.Return(3)
	mock.Unmock()
	if return2() != 2 {
		t.Fatal("function was not unmock")
	}
}

