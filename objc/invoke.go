package objc

/*
#cgo LDFLAGS: -lobjc
#define __OBJC2__ 1
#include <objc/message.h>
#include <stdlib.h>

// + (NSInvocation *)invocationWithMethodSignature:(NSMethodSignature *)sig;
id (* send_NSInvocation_invocationWithMethodSignature)(Class cls, SEL _cmd, id sig) = (id(*)(Class,SEL,id))objc_msgSend;

// - (NSMethodSignature *)methodSignatureForSelector:(SEL)aSelector;
id (* send_NSObject_methodSignatureForSelector)(id self, SEL _cmd, SEL aSelector) = (id(*)(id,SEL,SEL))objc_msgSend;

void (* send_NSInvocation_setTarget)(id self, SEL _cmd, id target) = (void(*)(id,SEL,id))objc_msgSend;
void (* send_NSInvocation_setSelector)(id self, SEL _cmd, SEL sel) = (void(*)(id,SEL,SEL))objc_msgSend;

// - (void)setArgument:(void *)argumentLocation atIndex:(NSInteger)idx;
void (* send_setArgument)(id self, SEL _cmd, void *arg, int idx) = (void(*)(id,SEL,void*,int))objc_msgSend;
void setArgument(void *invocation, void *arg, int idx) {
	send_setArgument(invocation, sel_registerName("setArgument:atIndex:"), arg, idx);
}

// - (void)invoke;
void (* send_invoke)(id self, SEL _cmd) = (void(*)(id,SEL))objc_msgSend;

// - (void)getReturnValue:(void *)retLoc;
void (* send_getReturnValue)(id self, SEL _cmd, void *retLoc) = (void(*)(id,SEL,void*))objc_msgSend;

void *methodSignatureForSelector(void *target, char *selName) {
	return send_NSObject_methodSignatureForSelector(target, sel_registerName("methodSignatureForSelector:"), sel_registerName(selName));
}

int (* send_numberOfArguments)(id self, SEL _cmd) = (int(*)(id,SEL))objc_msgSend;

int numberOfArguments(void *target) {
	return send_numberOfArguments(target, sel_registerName("numberOfArguments"));
}

const char * (* send_methodReturnType)(id self, SEL _cmd) = (const char * (*)(id,SEL))objc_msgSend;

const char * methodReturnType(void *target) {
	return send_methodReturnType(target, sel_registerName("methodReturnType"));
}

const char * (* send_getArgumentTypeAtIndex)(id self, SEL _cmd, int i) = (const char * (*)(id,SEL,int))objc_msgSend;

const char * getArgumentTypeAtIndex(void *target, int idx) {
	return send_getArgumentTypeAtIndex(target, sel_registerName("getArgumentTypeAtIndex:"), idx);
}

void *nsInvocation(void *target, char *selName) {
	Class cls = objc_getClass("NSInvocation");
	SEL sel = sel_registerName(selName);
	id sig = send_NSObject_methodSignatureForSelector(target, sel_registerName("methodSignatureForSelector:"), sel);
	id inv = send_NSInvocation_invocationWithMethodSignature(cls, sel_registerName("invocationWithMethodSignature:"), sig);
	send_NSInvocation_setTarget(inv, sel_registerName("setTarget:"), target);
	send_NSInvocation_setSelector(inv, sel_registerName("setSelector:"), sel);
	return inv;
}

void invoke(void *invocation, void *retLoc) {
	send_invoke(invocation, sel_registerName("invoke"));
	if (retLoc != NULL) {
		send_getReturnValue(invocation, sel_registerName("getReturnValue:"), retLoc);
	}
}
*/
import "C"
import "unsafe"

func numberOfArguments(id uintptr) int {
	return int(C.numberOfArguments(unsafe.Pointer(id)))
}

func methodReturnType(id uintptr) string {
	return C.GoString(C.methodReturnType(unsafe.Pointer(id)))
}

func getArgumentTypeAtIndex(id uintptr, idx int) string {
	return C.GoString(C.getArgumentTypeAtIndex(unsafe.Pointer(id), C.int(idx)))
}

func methodSignatureForSelector(target uintptr, selName string) uintptr {
	cSelName := C.CString(selName)
	sig := uintptr(C.methodSignatureForSelector(unsafe.Pointer(target), cSelName))
	C.free(unsafe.Pointer(cSelName))
	return sig
}

func newInvocation(target uintptr, selName string) uintptr {
	return uintptr(C.nsInvocation(unsafe.Pointer(target), C.CString(selName)))
}

func setArgumentAtIndex(call uintptr, arg unsafe.Pointer, index int) {
	C.setArgument(unsafe.Pointer(call), arg, C.int(index))
}

// TODO variant for calls with no return?
// or just split into two calls?
func invoke(call uintptr, dest unsafe.Pointer) {
	C.invoke(unsafe.Pointer(call), dest)
}

func cstring(ptr uintptr) string {
	return C.GoString((*C.char)(unsafe.Pointer(ptr)))
}

func send(retDest unsafe.Pointer, target uintptr, selName string, args ...unsafe.Pointer) {
	inv := newInvocation(target, selName)
	for i, arg := range args {
		setArgumentAtIndex(inv, arg, i+2)
	}
	invoke(inv, retDest)
}

func wrapArg()
