// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Composition] class.
var CompositionClass = _CompositionClass{objc.GetClass("QCComposition")}

type _CompositionClass struct {
	objc.Class
}

// An interface definition for the [Composition] class.
type IComposition interface {
	objc.IObject
}

// The QCComposition class represents a Quartz Composer composition that either: [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qccomposition?language=objc
type Composition struct {
	objc.Object
}

func CompositionFrom(ptr unsafe.Pointer) Composition {
	return Composition{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CompositionClass) Alloc() Composition {
	rv := objc.Call[Composition](cc, objc.Sel("alloc"))
	return rv
}

func Composition_Alloc() Composition {
	return CompositionClass.Alloc()
}

func (cc _CompositionClass) New() Composition {
	rv := objc.Call[Composition](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewComposition() Composition {
	return CompositionClass.New()
}

func (c_ Composition) Init() Composition {
	rv := objc.Call[Composition](c_, objc.Sel("init"))
	return rv
}
