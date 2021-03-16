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

func TestReplaceFunction(t *testing.T) {
	mock := Mock(return1)
	mock.Return(2)
	defer mock.Unmock()
	if return1() != 2 {
		t.Fatal("function was not replaced")
	}
}

func TestUnmockFunction(t *testing.T) {
	mock := Mock(return2)
	mock.Return(3)
	mock.Unmock()
	if return2() != 2 {
		t.Fatal("function was not unmock")
	}
}

type A struct {}

func (a *A) Method() int {
	return 1
}

func TestMockObject(t *testing.T) {
	a1 := new(A)
	a2 := new(A)
	Allow(a1, (*A).Method, 2)
	Allow(a2, (*A).Method, 3)
	a1.Method()
	a2.Method()
	if a1.Method() != 2 {
		t.Fatal("a1 method was not mocked")
	}
	if a2.Method() != 3 {
		t.Fatal("a2 method was not mocked")
	}
}

