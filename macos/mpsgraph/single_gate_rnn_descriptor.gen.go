// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SingleGateRNNDescriptor] class.
var SingleGateRNNDescriptorClass = _SingleGateRNNDescriptorClass{objc.GetClass("MPSGraphSingleGateRNNDescriptor")}

type _SingleGateRNNDescriptorClass struct {
	objc.Class
}

// An interface definition for the [SingleGateRNNDescriptor] class.
type ISingleGateRNNDescriptor interface {
	objc.IObject
	Activation() RNNActivation
	SetActivation(value RNNActivation)
	Training() bool
	SetTraining(value bool)
	Bidirectional() bool
	SetBidirectional(value bool)
	Reverse() bool
	SetReverse(value bool)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor?language=objc
type SingleGateRNNDescriptor struct {
	objc.Object
}

func SingleGateRNNDescriptorFrom(ptr unsafe.Pointer) SingleGateRNNDescriptor {
	return SingleGateRNNDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SingleGateRNNDescriptorClass) Descriptor() SingleGateRNNDescriptor {
	rv := objc.Call[SingleGateRNNDescriptor](sc, objc.Sel("descriptor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/3919774-descriptor?language=objc
func SingleGateRNNDescriptor_Descriptor() SingleGateRNNDescriptor {
	return SingleGateRNNDescriptorClass.Descriptor()
}

func (sc _SingleGateRNNDescriptorClass) Alloc() SingleGateRNNDescriptor {
	rv := objc.Call[SingleGateRNNDescriptor](sc, objc.Sel("alloc"))
	return rv
}

func SingleGateRNNDescriptor_Alloc() SingleGateRNNDescriptor {
	return SingleGateRNNDescriptorClass.Alloc()
}

func (sc _SingleGateRNNDescriptorClass) New() SingleGateRNNDescriptor {
	rv := objc.Call[SingleGateRNNDescriptor](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSingleGateRNNDescriptor() SingleGateRNNDescriptor {
	return SingleGateRNNDescriptorClass.New()
}

func (s_ SingleGateRNNDescriptor) Init() SingleGateRNNDescriptor {
	rv := objc.Call[SingleGateRNNDescriptor](s_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/3919772-activation?language=objc
func (s_ SingleGateRNNDescriptor) Activation() RNNActivation {
	rv := objc.Call[RNNActivation](s_, objc.Sel("activation"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/3919772-activation?language=objc
func (s_ SingleGateRNNDescriptor) SetActivation(value RNNActivation) {
	objc.Call[objc.Void](s_, objc.Sel("setActivation:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/3919776-training?language=objc
func (s_ SingleGateRNNDescriptor) Training() bool {
	rv := objc.Call[bool](s_, objc.Sel("training"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/3919776-training?language=objc
func (s_ SingleGateRNNDescriptor) SetTraining(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setTraining:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/3919773-bidirectional?language=objc
func (s_ SingleGateRNNDescriptor) Bidirectional() bool {
	rv := objc.Call[bool](s_, objc.Sel("bidirectional"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/3919773-bidirectional?language=objc
func (s_ SingleGateRNNDescriptor) SetBidirectional(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setBidirectional:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/3919775-reverse?language=objc
func (s_ SingleGateRNNDescriptor) Reverse() bool {
	rv := objc.Call[bool](s_, objc.Sel("reverse"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/3919775-reverse?language=objc
func (s_ SingleGateRNNDescriptor) SetReverse(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setReverse:"), value)
}
