package kagemusha

import (
	. "reflect"
)

type Function struct {
	value Value
}

func (f Function) Stub(returnValue interface{}) Function {
	ret := ValueOf(returnValue)
	fun := func(args []Value) []Value {
		return []Value{ret}
	}
	newFunc := MakeFunc(f.value.Type(), fun)
	return Function{newFunc}
}
