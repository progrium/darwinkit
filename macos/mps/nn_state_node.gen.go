// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNStateNode] class.
var NNStateNodeClass = _NNStateNodeClass{objc.GetClass("MPSNNStateNode")}

type _NNStateNodeClass struct {
	objc.Class
}

// An interface definition for the [NNStateNode] class.
type INNStateNode interface {
	objc.IObject
	Handle() HandleWrapper
	SetHandle(value PHandle)
	SetHandleObject(valueObject objc.IObject)
	ExportFromGraph() bool
	SetExportFromGraph(value bool)
	SynchronizeResource() bool
	SetSynchronizeResource(value bool)
}

// A placeholder node denoting the position in the graph of a state object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode?language=objc
type NNStateNode struct {
	objc.Object
}

func NNStateNodeFrom(ptr unsafe.Pointer) NNStateNode {
	return NNStateNode{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NNStateNodeClass) Alloc() NNStateNode {
	rv := objc.Call[NNStateNode](nc, objc.Sel("alloc"))
	return rv
}

func NNStateNode_Alloc() NNStateNode {
	return NNStateNodeClass.Alloc()
}

func (nc _NNStateNodeClass) New() NNStateNode {
	rv := objc.Call[NNStateNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNStateNode() NNStateNode {
	return NNStateNodeClass.New()
}

func (n_ NNStateNode) Init() NNStateNode {
	rv := objc.Call[NNStateNode](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/2866426-handle?language=objc
func (n_ NNStateNode) Handle() HandleWrapper {
	rv := objc.Call[HandleWrapper](n_, objc.Sel("handle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/2866426-handle?language=objc
func (n_ NNStateNode) SetHandle(value PHandle) {
	po0 := objc.WrapAsProtocol("MPSHandle", value)
	objc.Call[objc.Void](n_, objc.Sel("setHandle:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/2866426-handle?language=objc
func (n_ NNStateNode) SetHandleObject(valueObject objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setHandle:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/2942640-exportfromgraph?language=objc
func (n_ NNStateNode) ExportFromGraph() bool {
	rv := objc.Call[bool](n_, objc.Sel("exportFromGraph"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/2942640-exportfromgraph?language=objc
func (n_ NNStateNode) SetExportFromGraph(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setExportFromGraph:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/2942639-synchronizeresource?language=objc
func (n_ NNStateNode) SynchronizeResource() bool {
	rv := objc.Call[bool](n_, objc.Sel("synchronizeResource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/2942639-synchronizeresource?language=objc
func (n_ NNStateNode) SetSynchronizeResource(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setSynchronizeResource:"), value)
}
