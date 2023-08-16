// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FunctionStitchingAttributeAlwaysInline] class.
var FunctionStitchingAttributeAlwaysInlineClass = _FunctionStitchingAttributeAlwaysInlineClass{objc.GetClass("MTLFunctionStitchingAttributeAlwaysInline")}

type _FunctionStitchingAttributeAlwaysInlineClass struct {
	objc.Class
}

// An interface definition for the [FunctionStitchingAttributeAlwaysInline] class.
type IFunctionStitchingAttributeAlwaysInline interface {
	objc.IObject
}

// An attribute to specify that Metal needs to inline all of the function calls when generating the stitched function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingattributealwaysinline?language=objc
type FunctionStitchingAttributeAlwaysInline struct {
	objc.Object
}

func FunctionStitchingAttributeAlwaysInlineFrom(ptr unsafe.Pointer) FunctionStitchingAttributeAlwaysInline {
	return FunctionStitchingAttributeAlwaysInline{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FunctionStitchingAttributeAlwaysInlineClass) Alloc() FunctionStitchingAttributeAlwaysInline {
	rv := objc.Call[FunctionStitchingAttributeAlwaysInline](fc, objc.Sel("alloc"))
	return rv
}

func FunctionStitchingAttributeAlwaysInline_Alloc() FunctionStitchingAttributeAlwaysInline {
	return FunctionStitchingAttributeAlwaysInlineClass.Alloc()
}

func (fc _FunctionStitchingAttributeAlwaysInlineClass) New() FunctionStitchingAttributeAlwaysInline {
	rv := objc.Call[FunctionStitchingAttributeAlwaysInline](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFunctionStitchingAttributeAlwaysInline() FunctionStitchingAttributeAlwaysInline {
	return FunctionStitchingAttributeAlwaysInlineClass.New()
}

func (f_ FunctionStitchingAttributeAlwaysInline) Init() FunctionStitchingAttributeAlwaysInline {
	rv := objc.Call[FunctionStitchingAttributeAlwaysInline](f_, objc.Sel("init"))
	return rv
}
