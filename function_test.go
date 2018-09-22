package kagemusha

import (
	"testing"
	. "reflect"
)

func noArgs() int {
	return 1
}

func TestNoArgsFunctionCreation(t *testing.T) {
	newFunction := Function{noArgs}.Stub(2)
	actual, _ := ValueOf(newFunction.value).Call([]Value{})[0].Interface().(int)
	if actual != 2 {
		t.Fatal("fail create new function")
	}
}
