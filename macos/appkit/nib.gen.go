// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var NibClass = _NibClass{objc.GetClass("NSNib")}

type _NibClass struct {
	objc.Class
}

type INib interface {
	objc.IObject
	InstantiateWithOwnerTopLevelObjects(owner objc.IObject, topLevelObjects *foundation.Array) bool
}

type Nib struct {
	objc.Object
}

func MakeNib(ptr unsafe.Pointer) Nib {
	return Nib{
		Object: objc.MakeObject(ptr),
	}
}

func (n_ Nib) InitWithNibNamedBundle(nibName NibName, bundle foundation.IBundle) Nib {
	rv := objc.CallMethod[Nib](n_, objc.GetSelector("initWithNibNamed:bundle:"), nibName, objc.ExtractPtr(bundle))
	return rv
}

func Nib_InitWithNibNamedBundle(nibName NibName, bundle foundation.IBundle) Nib {
	return NibClass.Alloc().InitWithNibNamedBundle(nibName, bundle)
}

func (n_ Nib) InitWithNibDataBundle(nibData []byte, bundle foundation.IBundle) Nib {
	rv := objc.CallMethod[Nib](n_, objc.GetSelector("initWithNibData:bundle:"), nibData, objc.ExtractPtr(bundle))
	return rv
}

func Nib_InitWithNibDataBundle(nibData []byte, bundle foundation.IBundle) Nib {
	return NibClass.Alloc().InitWithNibDataBundle(nibData, bundle)
}

func (nc _NibClass) Alloc() Nib {
	rv := objc.CallMethod[Nib](nc, objc.GetSelector("alloc"))
	return rv
}

func Nib_Alloc() Nib {
	return NibClass.Alloc()
}

func (nc _NibClass) New() Nib {
	rv := objc.CallMethod[Nib](nc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewNib() Nib {
	return NibClass.New()
}

func Nib_New() Nib {
	return NibClass.New()
}

func (n_ Nib) Init() Nib {
	rv := objc.CallMethod[Nib](n_, objc.GetSelector("init"))
	return rv
}

func Nib_Init() Nib {
	return NibClass.Alloc().Init()
}

func (n_ Nib) InstantiateWithOwnerTopLevelObjects(owner objc.IObject, topLevelObjects *foundation.Array) bool {
	rv := objc.CallMethod[bool](n_, objc.GetSelector("instantiateWithOwner:topLevelObjects:"), objc.ExtractPtr(owner), topLevelObjects)
	return rv
}
