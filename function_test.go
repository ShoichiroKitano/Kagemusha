package kagemusha

import (
	"testing"
	. "reflect"
)

func noArgs() int {
	return 1
}

func TestNoArgsFunctionCreation(t *testing.T) {
	newFunction := Function{ValueOf(noArgs)}.Stub(2)
	actual := newFunction.value.Call([]Value{})
	expect := []Value{ValueOf(2)}
	if DeepEqual(actual, expect) {
		t.Fatal("fail create new function")
	}
}
