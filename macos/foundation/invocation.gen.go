// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Invocation] class.
var InvocationClass = _InvocationClass{objc.GetClass("NSInvocation")}

type _InvocationClass struct {
	objc.Class
}

// An interface definition for the [Invocation] class.
type IInvocation interface {
	objc.IObject
	GetArgumentAtIndex(argumentLocation unsafe.Pointer, idx int)
	SetReturnValue(retLoc unsafe.Pointer)
	RetainArguments()
	SetArgumentAtIndex(argumentLocation unsafe.Pointer, idx int)
	GetReturnValue(retLoc unsafe.Pointer)
	Invoke()
	InvokeWithTarget(target objc.IObject)
	Target() objc.Object
	SetTarget(value objc.IObject)
	Selector() objc.Selector
	SetSelector(value objc.Selector)
	ArgumentsRetained() bool
	MethodSignature() MethodSignature
}

// An Objective-C message rendered as an object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation?language=objc
type Invocation struct {
	objc.Object
}

func InvocationFrom(ptr unsafe.Pointer) Invocation {
	return Invocation{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _InvocationClass) Alloc() Invocation {
	rv := objc.Call[Invocation](ic, objc.Sel("alloc"))
	return rv
}

func Invocation_Alloc() Invocation {
	return InvocationClass.Alloc()
}

func (ic _InvocationClass) New() Invocation {
	rv := objc.Call[Invocation](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewInvocation() Invocation {
	return InvocationClass.New()
}

func (i_ Invocation) Init() Invocation {
	rv := objc.Call[Invocation](i_, objc.Sel("init"))
	return rv
}

// Returns by indirection the receiver's argument at a specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation/1437830-getargument?language=objc
func (i_ Invocation) GetArgumentAtIndex(argumentLocation unsafe.Pointer, idx int) {
	objc.Call[objc.Void](i_, objc.Sel("getArgument:atIndex:"), argumentLocation, idx)
}

// Sets the receiver’s return value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation/1437848-setreturnvalue?language=objc
func (i_ Invocation) SetReturnValue(retLoc unsafe.Pointer) {
	objc.Call[objc.Void](i_, objc.Sel("setReturnValue:"), retLoc)
}

// If the receiver hasn’t already done so, retains the target and all object arguments of the receiver and copies all of its C-string arguments and blocks. If a returnvalue has been set, this is also retained or copied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation/1437838-retainarguments?language=objc
func (i_ Invocation) RetainArguments() {
	objc.Call[objc.Void](i_, objc.Sel("retainArguments"))
}

// Sets an argument of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation/1437834-setargument?language=objc
func (i_ Invocation) SetArgumentAtIndex(argumentLocation unsafe.Pointer, idx int) {
	objc.Call[objc.Void](i_, objc.Sel("setArgument:atIndex:"), argumentLocation, idx)
}

// Returns an NSInvocation object able to construct messages using a given method signature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation/1437844-invocationwithmethodsignature?language=objc
func (ic _InvocationClass) InvocationWithMethodSignature(sig IMethodSignature) Invocation {
	rv := objc.Call[Invocation](ic, objc.Sel("invocationWithMethodSignature:"), objc.Ptr(sig))
	return rv
}

// Returns an NSInvocation object able to construct messages using a given method signature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation/1437844-invocationwithmethodsignature?language=objc
func Invocation_InvocationWithMethodSignature(sig IMethodSignature) Invocation {
	return InvocationClass.InvocationWithMethodSignature(sig)
}

// Gets the receiver's return value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation/1437832-getreturnvalue?language=objc
func (i_ Invocation) GetReturnValue(retLoc unsafe.Pointer) {
	objc.Call[objc.Void](i_, objc.Sel("getReturnValue:"), retLoc)
}

// Sends the receiver’s message (with arguments) to its target and sets the return value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation/1437850-invoke?language=objc
func (i_ Invocation) Invoke() {
	objc.Call[objc.Void](i_, objc.Sel("invoke"))
}

// Sets the receiver’s target, sends the receiver’s message (with arguments) to that target, and sets the return value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation/1437854-invokewithtarget?language=objc
func (i_ Invocation) InvokeWithTarget(target objc.IObject) {
	objc.Call[objc.Void](i_, objc.Sel("invokeWithTarget:"), target)
}

// The receiver’s target, or nil if the receiver has no target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation/1437852-target?language=objc
func (i_ Invocation) Target() objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("target"))
	return rv
}

// The receiver’s target, or nil if the receiver has no target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation/1437852-target?language=objc
func (i_ Invocation) SetTarget(value objc.IObject) {
	objc.Call[objc.Void](i_, objc.Sel("setTarget:"), value)
}

// The receiver’s selector, or 0 if it hasn’t been set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation/1437836-selector?language=objc
func (i_ Invocation) Selector() objc.Selector {
	rv := objc.Call[objc.Selector](i_, objc.Sel("selector"))
	return rv
}

// The receiver’s selector, or 0 if it hasn’t been set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation/1437836-selector?language=objc
func (i_ Invocation) SetSelector(value objc.Selector) {
	objc.Call[objc.Void](i_, objc.Sel("setSelector:"), value)
}

// YES if the receiver has retained its arguments, NO otherwise. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation/1437842-argumentsretained?language=objc
func (i_ Invocation) ArgumentsRetained() bool {
	rv := objc.Call[bool](i_, objc.Sel("argumentsRetained"))
	return rv
}

// The receiver’s method signature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocation/1437846-methodsignature?language=objc
func (i_ Invocation) MethodSignature() MethodSignature {
	rv := objc.Call[MethodSignature](i_, objc.Sel("methodSignature"))
	return rv
}
