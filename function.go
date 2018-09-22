package kagemusha

import (
	. "reflect"
)

type Function struct {
	value interface{}
}

func (f Function) Stub(returnValue interface{}) Function {
	ret := ValueOf(returnValue)
	fun := func(args []Value) []Value {
		return []Value{ret}
	}
	newFunc := MakeFunc(ValueOf(f.value).Type(), fun)
	return Function{newFunc.Interface()}
}
