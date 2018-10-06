package kagemusha

import (
	"reflect"
	"syscall"
	"unsafe"
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

