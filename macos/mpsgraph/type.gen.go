// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Type] class.
var TypeClass = _TypeClass{objc.GetClass("MPSGraphType")}

type _TypeClass struct {
	objc.Class
}

// An interface definition for the [Type] class.
type IType interface {
	objc.IObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtype?language=objc
type Type struct {
	objc.Object
}

func TypeFrom(ptr unsafe.Pointer) Type {
	return Type{
		Object: objc.ObjectFrom(ptr),
	}
}

func (tc _TypeClass) Alloc() Type {
	rv := objc.Call[Type](tc, objc.Sel("alloc"))
	return rv
}

func Type_Alloc() Type {
	return TypeClass.Alloc()
}

func (tc _TypeClass) New() Type {
	rv := objc.Call[Type](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewType() Type {
	return TypeClass.New()
}

func (t_ Type) Init() Type {
	rv := objc.Call[Type](t_, objc.Sel("init"))
	return rv
}
