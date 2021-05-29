package objc

/*
#cgo LDFLAGS: -lobjc
#define __OBJC2__ 1
#include <objc/message.h>

// + (NSInvocation *)invocationWithMethodSignature:(NSMethodSignature *)sig;
id (* send_NSInvocation_invocationWithMethodSignature)(Class cls, SEL _cmd, id sig) = (id(*)(Class,SEL,id))objc_msgSend;
// + (NSMethodSignature *)signatureWithObjCTypes:(const char *)types;
// - (NSMethodSignature *)methodSignatureForSelector:(SEL)aSelector;
id (* send_NSObject_methodSignatureForSelector)(id self, SEL _cmd, SEL aSelector) = (id(*)(id,SEL,SEL))objc_msgSend;
void (* send_NSInvocation_setTarget)(id self, SEL _cmd, id target) = (void(*)(id,SEL,id))objc_msgSend;
void (* send_NSInvocation_setSelector)(id self, SEL _cmd, SEL sel) = (void(*)(id,SEL,SEL))objc_msgSend;

// - (void)setArgument:(void *)argumentLocation atIndex:(NSInteger)idx;
// void (* send_setArgument)(id self, SEL _cmd, void *arg, NSInteger idx) = (void(*)(id,SEL,void*,NSInteger))objc_msgSend;
// - (void)invoke;
void (* send_invoke)(id self, SEL _cmd) = (void(*)(id,SEL))objc_msgSend;
void (* send_invokeWithTarget)(id self, SEL _cmd,id) = (void(*)(id,SEL,id))objc_msgSend;
// - (void)getReturnValue:(void *)retLoc;
void (* send_getReturnValue)(id self, SEL _cmd, void *retLoc) = (void(*)(id,SEL,void*))objc_msgSend;

void *cgoMethodSignatureForSelector(void *target, char *selName) {
	return send_NSObject_methodSignatureForSelector(target, sel_registerName("methodSignatureForSelector:"), sel_registerName(selName));
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
	// send_invokeWithTarget(invocation, sel_registerName("invoke"), target);
	send_getReturnValue(invocation, sel_registerName("getReturnValue:"), retLoc);
}
*/
import "C"
import "unsafe"

func MethodSignatureForSelector(target uintptr, selName string) uintptr {
	return uintptr(C.cgoMethodSignatureForSelector(unsafe.Pointer(target), C.CString(selName)))
}

func NewInvocation(target uintptr, selName string) uintptr {
	return uintptr(C.nsInvocation(unsafe.Pointer(target), C.CString(selName)))
}

// TODO variant for calls with no return?
// or just split into two calls?
func invoke(call uintptr, dest unsafe.Pointer) {
	C.invoke(unsafe.Pointer(call), dest)
}

// + (NSInvocation *)invocationWithMethodSignature:(NSMethodSignature *)sig;
// + (NSMethodSignature *)signatureWithObjCTypes:(const char *)types;
// - (NSMethodSignature *)methodSignatureForSelector:(SEL)aSelector;
