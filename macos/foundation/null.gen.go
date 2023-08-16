// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Null] class.
var NullClass = _NullClass{objc.GetClass("NSNull")}

type _NullClass struct {
	objc.Class
}

// An interface definition for the [Null] class.
type INull interface {
	objc.IObject
}

// A singleton object used to represent null values in collection objects that don’t allow nil values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnull?language=objc
type Null struct {
	objc.Object
}

func NullFrom(ptr unsafe.Pointer) Null {
	return Null{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NullClass) Alloc() Null {
	rv := objc.Call[Null](nc, objc.Sel("alloc"))
	return rv
}

func Null_Alloc() Null {
	return NullClass.Alloc()
}

func (nc _NullClass) New() Null {
	rv := objc.Call[Null](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNull() Null {
	return NullClass.New()
}

func (n_ Null) Init() Null {
	rv := objc.Call[Null](n_, objc.Sel("init"))
	return rv
}

// Returns the singleton instance of NSNull. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnull/1520557-null?language=objc
func (nc _NullClass) Null() Null {
	rv := objc.Call[Null](nc, objc.Sel("null"))
	return rv
}

// Returns the singleton instance of NSNull. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnull/1520557-null?language=objc
func Null_Null() Null {
	return NullClass.Null()
}
