package variadic

import "unsafe"

/*
#cgo LDFLAGS: -lobjc
#import <objc/message.h>
void * addr_msgSend = &objc_msgSend;
void * addr_msgSendSuper = &objc_msgSendSuper;
#if defined(__aarch64__)
void * addr_msgSend_stret = NULL;
void * addr_msgSendSuper_stret = NULL;
#else
void * addr_msgSend_stret = &objc_msgSend_stret;
void * addr_msgSendSuper_stret = &objc_msgSendSuper_stret;
#endif
*/
import "C"

type Function int

const (
	F_msgSend Function = iota
	F_msgSendSuper
	F_msgSend_stret
	F_msgSendSuper_stret
)

func (f Function) String() string {
	switch f {
	case F_msgSend:
		return "objc_msgSend"
	case F_msgSendSuper:
		return "objc_msgSendSuper"
	case F_msgSend_stret:
		return "objc_msgSend_stret"
	case F_msgSendSuper_stret:
		return "objc_msgSendSuper_stret"
	}
	panic("unknown function")
}

func (f Function) IsSuper() bool {
	switch f {
	case F_msgSendSuper, F_msgSendSuper_stret:
		return true
	default:
		return false
	}
}

func (f Function) IsStRet() bool {
	switch f {
	case F_msgSend_stret, F_msgSendSuper_stret:
		return true
	default:
		return false
	}
}

func (f Function) StRet() Function {
	switch f {
	case F_msgSend:
		return F_msgSend_stret
	case F_msgSendSuper:
		return F_msgSendSuper_stret
	default:
		return f
	}
}

func (f Function) Addr() unsafe.Pointer {
	switch f {
	case F_msgSend:
		return C.addr_msgSend
	case F_msgSendSuper:
		return C.addr_msgSendSuper
	case F_msgSend_stret:
		return C.addr_msgSend_stret
	case F_msgSendSuper_stret:
		return C.addr_msgSendSuper_stret
	}
	panic("unknown function")
}

func (f Function) NewCall() *FunctionCall {
	return NewFunctionCallAddr(f.Addr())
}
