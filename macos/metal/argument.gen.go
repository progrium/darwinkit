// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Argument] class.
var ArgumentClass = _ArgumentClass{objc.GetClass("MTLArgument")}

type _ArgumentClass struct {
	objc.Class
}

// An interface definition for the [Argument] class.
type IArgument interface {
	objc.IObject
}

// Information about an argument of a graphics or compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargument?language=objc
type Argument struct {
	objc.Object
}

func ArgumentFrom(ptr unsafe.Pointer) Argument {
	return Argument{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _ArgumentClass) Alloc() Argument {
	rv := objc.Call[Argument](ac, objc.Sel("alloc"))
	return rv
}

func Argument_Alloc() Argument {
	return ArgumentClass.Alloc()
}

func (ac _ArgumentClass) New() Argument {
	rv := objc.Call[Argument](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewArgument() Argument {
	return ArgumentClass.New()
}

func (a_ Argument) Init() Argument {
	rv := objc.Call[Argument](a_, objc.Sel("init"))
	return rv
}
