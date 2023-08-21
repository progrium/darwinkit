// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FunctionStitchingFunctionNode] class.
var FunctionStitchingFunctionNodeClass = _FunctionStitchingFunctionNodeClass{objc.GetClass("MTLFunctionStitchingFunctionNode")}

type _FunctionStitchingFunctionNodeClass struct {
	objc.Class
}

// An interface definition for the [FunctionStitchingFunctionNode] class.
type IFunctionStitchingFunctionNode interface {
	objc.IObject
	Arguments() []FunctionStitchingNodeWrapper
	SetArguments(value []PFunctionStitchingNode)
	Name() string
	SetName(value string)
	ControlDependencies() []FunctionStitchingFunctionNode
	SetControlDependencies(value []IFunctionStitchingFunctionNode)
}

// A call graph node that describes a function call and its inputs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingfunctionnode?language=objc
type FunctionStitchingFunctionNode struct {
	objc.Object
}

func FunctionStitchingFunctionNodeFrom(ptr unsafe.Pointer) FunctionStitchingFunctionNode {
	return FunctionStitchingFunctionNode{
		Object: objc.ObjectFrom(ptr),
	}
}

func (f_ FunctionStitchingFunctionNode) InitWithNameArgumentsControlDependencies(name string, arguments []PFunctionStitchingNode, controlDependencies []IFunctionStitchingFunctionNode) FunctionStitchingFunctionNode {
	rv := objc.Call[FunctionStitchingFunctionNode](f_, objc.Sel("initWithName:arguments:controlDependencies:"), name, arguments, controlDependencies)
	return rv
}

// Creates a new function node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingfunctionnode/3750536-initwithname?language=objc
func NewFunctionStitchingFunctionNodeWithNameArgumentsControlDependencies(name string, arguments []PFunctionStitchingNode, controlDependencies []IFunctionStitchingFunctionNode) FunctionStitchingFunctionNode {
	instance := FunctionStitchingFunctionNodeClass.Alloc().InitWithNameArgumentsControlDependencies(name, arguments, controlDependencies)
	instance.Autorelease()
	return instance
}

func (fc _FunctionStitchingFunctionNodeClass) Alloc() FunctionStitchingFunctionNode {
	rv := objc.Call[FunctionStitchingFunctionNode](fc, objc.Sel("alloc"))
	return rv
}

func FunctionStitchingFunctionNode_Alloc() FunctionStitchingFunctionNode {
	return FunctionStitchingFunctionNodeClass.Alloc()
}

func (fc _FunctionStitchingFunctionNodeClass) New() FunctionStitchingFunctionNode {
	rv := objc.Call[FunctionStitchingFunctionNode](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFunctionStitchingFunctionNode() FunctionStitchingFunctionNode {
	return FunctionStitchingFunctionNodeClass.New()
}

func (f_ FunctionStitchingFunctionNode) Init() FunctionStitchingFunctionNode {
	rv := objc.Call[FunctionStitchingFunctionNode](f_, objc.Sel("init"))
	return rv
}

// An ordered list of the nodes that provide the function’s arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingfunctionnode/3750534-arguments?language=objc
func (f_ FunctionStitchingFunctionNode) Arguments() []FunctionStitchingNodeWrapper {
	rv := objc.Call[[]FunctionStitchingNodeWrapper](f_, objc.Sel("arguments"))
	return rv
}

// An ordered list of the nodes that provide the function’s arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingfunctionnode/3750534-arguments?language=objc
func (f_ FunctionStitchingFunctionNode) SetArguments(value []PFunctionStitchingNode) {
	objc.Call[objc.Void](f_, objc.Sel("setArguments:"), value)
}

// The name of the function to call. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingfunctionnode/3750537-name?language=objc
func (f_ FunctionStitchingFunctionNode) Name() string {
	rv := objc.Call[string](f_, objc.Sel("name"))
	return rv
}

// The name of the function to call. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingfunctionnode/3750537-name?language=objc
func (f_ FunctionStitchingFunctionNode) SetName(value string) {
	objc.Call[objc.Void](f_, objc.Sel("setName:"), value)
}

// The list of nodes that must execute before executing the node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingfunctionnode/3750535-controldependencies?language=objc
func (f_ FunctionStitchingFunctionNode) ControlDependencies() []FunctionStitchingFunctionNode {
	rv := objc.Call[[]FunctionStitchingFunctionNode](f_, objc.Sel("controlDependencies"))
	return rv
}

// The list of nodes that must execute before executing the node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingfunctionnode/3750535-controldependencies?language=objc
func (f_ FunctionStitchingFunctionNode) SetControlDependencies(value []IFunctionStitchingFunctionNode) {
	objc.Call[objc.Void](f_, objc.Sel("setControlDependencies:"), value)
}
