// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PolygonBuffer] class.
var PolygonBufferClass = _PolygonBufferClass{objc.GetClass("MPSPolygonBuffer")}

type _PolygonBufferClass struct {
	objc.Class
}

// An interface definition for the [PolygonBuffer] class.
type IPolygonBuffer interface {
	objc.IObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer?language=objc
type PolygonBuffer struct {
	objc.Object
}

func PolygonBufferFrom(ptr unsafe.Pointer) PolygonBuffer {
	return PolygonBuffer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PolygonBufferClass) Alloc() PolygonBuffer {
	rv := objc.Call[PolygonBuffer](pc, objc.Sel("alloc"))
	return rv
}

func PolygonBuffer_Alloc() PolygonBuffer {
	return PolygonBufferClass.Alloc()
}

func (pc _PolygonBufferClass) New() PolygonBuffer {
	rv := objc.Call[PolygonBuffer](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPolygonBuffer() PolygonBuffer {
	return PolygonBufferClass.New()
}

func (p_ PolygonBuffer) Init() PolygonBuffer {
	rv := objc.Call[PolygonBuffer](p_, objc.Sel("init"))
	return rv
}
