package kagemusha

import (
	"reflect"
	"strings"
	"syscall"
	"unsafe"
	"runtime"
)

type ValuePointer struct {
	_   uintptr
	pointer unsafe.Pointer
}

func toPointer(v reflect.Value) uintptr {
	pointer := (*ValuePointer)(unsafe.Pointer(&v)).pointer
	return (uintptr)(pointer)
}

type Kagemusha struct {
	mocked Function
	original []byte
}

func Mock(mocked interface{}) *Kagemusha {
	return &Kagemusha{mocked: Function{mocked}}
}

func (self *Kagemusha) Unmock() {
	originalFuncAddress := reflect.ValueOf(self.mocked.value).Pointer()
	original := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: originalFuncAddress,
		Len: len(self.original),
		Cap: len(self.original),
	}))
	pageStart := originalFuncAddress & ^(uintptr(syscall.Getpagesize() - 1))
	page := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: pageStart,
		Len:  syscall.Getpagesize(),
		Cap:  syscall.Getpagesize(),
	}))
	syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
	copy(original, self.original)
	syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_EXEC)
}

func (self *Kagemusha) Return(returnValue interface{}) {
	stub := self.mocked.Stub(returnValue).value
	stubAddress := toPointer(reflect.ValueOf(stub))
	callStub := []byte{
		0x48, 0xBA,
	  byte(stubAddress >> 0),
	  byte(stubAddress >> 8),
	  byte(stubAddress >> 16),
	  byte(stubAddress >> 24),
	  byte(stubAddress >> 32),
	  byte(stubAddress >> 40),
	  byte(stubAddress >> 48),
	  byte(stubAddress >> 56),
	  0xFF, 0x22,
	}
	originalFuncAddress := reflect.ValueOf(self.mocked.value).Pointer()
	original := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: originalFuncAddress,
		Len: len(callStub),
		Cap: len(callStub),
	}))
	self.original = make([]byte, len(callStub))
	copy(self.original, original)
	pageStart := originalFuncAddress & ^(uintptr(syscall.Getpagesize() - 1))
	page := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: pageStart,
		Len:  syscall.Getpagesize(),
		Cap:  syscall.Getpagesize(),
	}))
	syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
	copy(original, callStub)
	syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_EXEC)
}

var m map[uintptr]reflect.Value = map[uintptr]reflect.Value{}

func Allow(obj interface{}, method interface{}, ret interface{}) {
	methodValue := reflect.ValueOf(method)
	slice := strings.Split(runtime.FuncForPC(methodValue.Pointer()).Name(), ".")
	methodName := slice[len(slice)-1]
	o := reflect.ValueOf(obj)
	if len(m) == 0 {
		table := func(args []reflect.Value) []reflect.Value {
			return []reflect.Value{reflect.ValueOf(2)}
		}
		newMethod := reflect.MakeFunc(o.MethodByName(methodName).Type(), table)
		stubAddress := toPointer(newMethod)
		callStub := []byte{
			0x48, 0xBA,
			byte(stubAddress >> 0),
			byte(stubAddress >> 8),
			byte(stubAddress >> 16),
			byte(stubAddress >> 24),
			byte(stubAddress >> 32),
			byte(stubAddress >> 40),
			byte(stubAddress >> 48),
			byte(stubAddress >> 56),
			0xFF, 0x22,
		}
		originalFuncAddress := methodValue.Pointer()
		original := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
			Data: originalFuncAddress,
			Len: len(callStub),
			Cap: len(callStub),
		}))
		pageStart := originalFuncAddress & ^(uintptr(syscall.Getpagesize() - 1))
		page := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
			Data: pageStart,
			Len:  syscall.Getpagesize(),
			Cap:  syscall.Getpagesize(),
		}))
		syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
		copy(original, callStub)
		syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_EXEC)
	}
	m[toPointer(o)] = o.MethodByName(methodName)
}

