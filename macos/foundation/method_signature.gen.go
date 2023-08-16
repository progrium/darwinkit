// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MethodSignature] class.
var MethodSignatureClass = _MethodSignatureClass{objc.GetClass("NSMethodSignature")}

type _MethodSignatureClass struct {
	objc.Class
}

// An interface definition for the [MethodSignature] class.
type IMethodSignature interface {
	objc.IObject
	IsOneway() bool
	GetArgumentTypeAtIndex(idx uint) *uint8
	MethodReturnLength() uint
	FrameLength() uint
	NumberOfArguments() uint
	MethodReturnType() *uint8
}

// A record of the type information for the return value and parameters of a method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmethodsignature?language=objc
type MethodSignature struct {
	objc.Object
}

func MethodSignatureFrom(ptr unsafe.Pointer) MethodSignature {
	return MethodSignature{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MethodSignatureClass) Alloc() MethodSignature {
	rv := objc.Call[MethodSignature](mc, objc.Sel("alloc"))
	return rv
}

func MethodSignature_Alloc() MethodSignature {
	return MethodSignatureClass.Alloc()
}

func (mc _MethodSignatureClass) New() MethodSignature {
	rv := objc.Call[MethodSignature](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMethodSignature() MethodSignature {
	return MethodSignatureClass.New()
}

func (m_ MethodSignature) Init() MethodSignature {
	rv := objc.Call[MethodSignature](m_, objc.Sel("init"))
	return rv
}

// Whether the receiver is asynchronous when invoked through distributed objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmethodsignature/1519664-isoneway?language=objc
func (m_ MethodSignature) IsOneway() bool {
	rv := objc.Call[bool](m_, objc.Sel("isOneway"))
	return rv
}

// Returns an NSMethodSignature object for the given Objective-C method type string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmethodsignature/1519670-signaturewithobjctypes?language=objc
func (mc _MethodSignatureClass) SignatureWithObjCTypes(types *uint8) MethodSignature {
	rv := objc.Call[MethodSignature](mc, objc.Sel("signatureWithObjCTypes:"), types)
	return rv
}

// Returns an NSMethodSignature object for the given Objective-C method type string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmethodsignature/1519670-signaturewithobjctypes?language=objc
func MethodSignature_SignatureWithObjCTypes(types *uint8) MethodSignature {
	return MethodSignatureClass.SignatureWithObjCTypes(types)
}

// Returns the type encoding for the argument at a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmethodsignature/1519660-getargumenttypeatindex?language=objc
func (m_ MethodSignature) GetArgumentTypeAtIndex(idx uint) *uint8 {
	rv := objc.Call[*uint8](m_, objc.Sel("getArgumentTypeAtIndex:"), idx)
	return rv
}

// The number of bytes required for the return value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmethodsignature/1519666-methodreturnlength?language=objc
func (m_ MethodSignature) MethodReturnLength() uint {
	rv := objc.Call[uint](m_, objc.Sel("methodReturnLength"))
	return rv
}

// The number of bytes that the arguments, taken together, occupy on the stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmethodsignature/1519658-framelength?language=objc
func (m_ MethodSignature) FrameLength() uint {
	rv := objc.Call[uint](m_, objc.Sel("frameLength"))
	return rv
}

// The number of arguments recorded in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmethodsignature/1519662-numberofarguments?language=objc
func (m_ MethodSignature) NumberOfArguments() uint {
	rv := objc.Call[uint](m_, objc.Sel("numberOfArguments"))
	return rv
}

// A C string encoding the return type of the method in Objective-C type encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmethodsignature/1519667-methodreturntype?language=objc
func (m_ MethodSignature) MethodReturnType() *uint8 {
	rv := objc.Call[*uint8](m_, objc.Sel("methodReturnType"))
	return rv
}
