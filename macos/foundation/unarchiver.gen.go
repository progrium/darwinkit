// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Unarchiver] class.
var UnarchiverClass = _UnarchiverClass{objc.GetClass("NSUnarchiver")}

type _UnarchiverClass struct {
	objc.Class
}

// An interface definition for the [Unarchiver] class.
type IUnarchiver interface {
	ICoder
}

// A decoder that restores data from an archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunarchiver?language=objc
type Unarchiver struct {
	Coder
}

func UnarchiverFrom(ptr unsafe.Pointer) Unarchiver {
	return Unarchiver{
		Coder: CoderFrom(ptr),
	}
}

func (uc _UnarchiverClass) Alloc() Unarchiver {
	rv := objc.Call[Unarchiver](uc, objc.Sel("alloc"))
	return rv
}

func Unarchiver_Alloc() Unarchiver {
	return UnarchiverClass.Alloc()
}

func (uc _UnarchiverClass) New() Unarchiver {
	rv := objc.Call[Unarchiver](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnarchiver() Unarchiver {
	return UnarchiverClass.New()
}

func (u_ Unarchiver) Init() Unarchiver {
	rv := objc.Call[Unarchiver](u_, objc.Sel("init"))
	return rv
}
