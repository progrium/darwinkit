// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GarbageCollector] class.
var GarbageCollectorClass = _GarbageCollectorClass{objc.GetClass("NSGarbageCollector")}

type _GarbageCollectorClass struct {
	objc.Class
}

// An interface definition for the [GarbageCollector] class.
type IGarbageCollector interface {
	objc.IObject
}

// A convenient interface to the garbage collection system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsgarbagecollector?language=objc
type GarbageCollector struct {
	objc.Object
}

func GarbageCollectorFrom(ptr unsafe.Pointer) GarbageCollector {
	return GarbageCollector{
		Object: objc.ObjectFrom(ptr),
	}
}

func (gc _GarbageCollectorClass) Alloc() GarbageCollector {
	rv := objc.Call[GarbageCollector](gc, objc.Sel("alloc"))
	return rv
}

func GarbageCollector_Alloc() GarbageCollector {
	return GarbageCollectorClass.Alloc()
}

func (gc _GarbageCollectorClass) New() GarbageCollector {
	rv := objc.Call[GarbageCollector](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGarbageCollector() GarbageCollector {
	return GarbageCollectorClass.New()
}

func (g_ GarbageCollector) Init() GarbageCollector {
	rv := objc.Call[GarbageCollector](g_, objc.Sel("init"))
	return rv
}
