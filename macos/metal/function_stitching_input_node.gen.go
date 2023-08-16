// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FunctionStitchingInputNode] class.
var FunctionStitchingInputNodeClass = _FunctionStitchingInputNodeClass{objc.GetClass("MTLFunctionStitchingInputNode")}

type _FunctionStitchingInputNodeClass struct {
	objc.Class
}

// An interface definition for the [FunctionStitchingInputNode] class.
type IFunctionStitchingInputNode interface {
	objc.IObject
	ArgumentIndex() uint
	SetArgumentIndex(value uint)
}

// A call graph node that describes an input to the call graph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinginputnode?language=objc
type FunctionStitchingInputNode struct {
	objc.Object
}

func FunctionStitchingInputNodeFrom(ptr unsafe.Pointer) FunctionStitchingInputNode {
	return FunctionStitchingInputNode{
		Object: objc.ObjectFrom(ptr),
	}
}

func (f_ FunctionStitchingInputNode) InitWithArgumentIndex(argument uint) FunctionStitchingInputNode {
	rv := objc.Call[FunctionStitchingInputNode](f_, objc.Sel("initWithArgumentIndex:"), argument)
	return rv
}

// Creates a new input node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinginputnode/3750546-initwithargumentindex?language=objc
func FunctionStitchingInputNode_InitWithArgumentIndex(argument uint) FunctionStitchingInputNode {
	return FunctionStitchingInputNodeClass.Alloc().InitWithArgumentIndex(argument)
}

func (fc _FunctionStitchingInputNodeClass) Alloc() FunctionStitchingInputNode {
	rv := objc.Call[FunctionStitchingInputNode](fc, objc.Sel("alloc"))
	return rv
}

func FunctionStitchingInputNode_Alloc() FunctionStitchingInputNode {
	return FunctionStitchingInputNodeClass.Alloc()
}

func (fc _FunctionStitchingInputNodeClass) New() FunctionStitchingInputNode {
	rv := objc.Call[FunctionStitchingInputNode](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFunctionStitchingInputNode() FunctionStitchingInputNode {
	return FunctionStitchingInputNodeClass.New()
}

func (f_ FunctionStitchingInputNode) Init() FunctionStitchingInputNode {
	rv := objc.Call[FunctionStitchingInputNode](f_, objc.Sel("init"))
	return rv
}

// The index in the command’s buffer argument table that declares which data to read for this input node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinginputnode/3750545-argumentindex?language=objc
func (f_ FunctionStitchingInputNode) ArgumentIndex() uint {
	rv := objc.Call[uint](f_, objc.Sel("argumentIndex"))
	return rv
}

// The index in the command’s buffer argument table that declares which data to read for this input node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinginputnode/3750545-argumentindex?language=objc
func (f_ FunctionStitchingInputNode) SetArgumentIndex(value uint) {
	objc.Call[objc.Void](f_, objc.Sel("setArgumentIndex:"), value)
}
