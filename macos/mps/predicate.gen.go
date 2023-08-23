// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Predicate] class.
var PredicateClass = _PredicateClass{objc.GetClass("MPSPredicate")}

type _PredicateClass struct {
	objc.Class
}

// An interface definition for the [Predicate] class.
type IPredicate interface {
	objc.IObject
	PredicateOffset() uint
	PredicateBuffer() metal.BufferWrapper
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate?language=objc
type Predicate struct {
	objc.Object
}

func PredicateFrom(ptr unsafe.Pointer) Predicate {
	return Predicate{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ Predicate) InitWithDevice(device metal.PDevice) Predicate {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[Predicate](p_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/3114035-initwithdevice?language=objc
func NewPredicateWithDevice(device metal.PDevice) Predicate {
	instance := PredicateClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (pc _PredicateClass) PredicateWithBufferOffset(buffer metal.PBuffer, offset uint) Predicate {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	rv := objc.Call[Predicate](pc, objc.Sel("predicateWithBuffer:offset:"), po0, offset)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/3114038-predicatewithbuffer?language=objc
func Predicate_PredicateWithBufferOffset(buffer metal.PBuffer, offset uint) Predicate {
	return PredicateClass.PredicateWithBufferOffset(buffer, offset)
}

func (p_ Predicate) InitWithBufferOffset(buffer metal.PBuffer, offset uint) Predicate {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	rv := objc.Call[Predicate](p_, objc.Sel("initWithBuffer:offset:"), po0, offset)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/3114034-initwithbuffer?language=objc
func NewPredicateWithBufferOffset(buffer metal.PBuffer, offset uint) Predicate {
	instance := PredicateClass.Alloc().InitWithBufferOffset(buffer, offset)
	instance.Autorelease()
	return instance
}

func (pc _PredicateClass) Alloc() Predicate {
	rv := objc.Call[Predicate](pc, objc.Sel("alloc"))
	return rv
}

func Predicate_Alloc() Predicate {
	return PredicateClass.Alloc()
}

func (pc _PredicateClass) New() Predicate {
	rv := objc.Call[Predicate](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPredicate() Predicate {
	return PredicateClass.New()
}

func (p_ Predicate) Init() Predicate {
	rv := objc.Call[Predicate](p_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/3114037-predicateoffset?language=objc
func (p_ Predicate) PredicateOffset() uint {
	rv := objc.Call[uint](p_, objc.Sel("predicateOffset"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/3114036-predicatebuffer?language=objc
func (p_ Predicate) PredicateBuffer() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](p_, objc.Sel("predicateBuffer"))
	return rv
}
