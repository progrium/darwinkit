// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Nib] class.
var NibClass = _NibClass{objc.GetClass("NSNib")}

type _NibClass struct {
	objc.Class
}

// An interface definition for the [Nib] class.
type INib interface {
	objc.IObject
	InstantiateWithOwnerTopLevelObjects(owner objc.IObject, topLevelObjects []objc.IObject) bool
}

// An object wrapper, or container, for an Interface Builder nib file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnib?language=objc
type Nib struct {
	objc.Object
}

func NibFrom(ptr unsafe.Pointer) Nib {
	return Nib{
		Object: objc.ObjectFrom(ptr),
	}
}

func (n_ Nib) InitWithNibNamedBundle(nibName NibName, bundle foundation.IBundle) Nib {
	rv := objc.Call[Nib](n_, objc.Sel("initWithNibNamed:bundle:"), nibName, objc.Ptr(bundle))
	return rv
}

// Returns an NSNib object initialized to the nib file in the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnib/1533932-initwithnibnamed?language=objc
func Nib_InitWithNibNamedBundle(nibName NibName, bundle foundation.IBundle) Nib {
	return NibClass.Alloc().InitWithNibNamedBundle(nibName, bundle)
}

func (n_ Nib) InitWithNibDataBundle(nibData []byte, bundle foundation.IBundle) Nib {
	rv := objc.Call[Nib](n_, objc.Sel("initWithNibData:bundle:"), nibData, objc.Ptr(bundle))
	return rv
}

// Initializes an instance with nib data and specified bundle for locating resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnib/1535865-initwithnibdata?language=objc
func Nib_InitWithNibDataBundle(nibData []byte, bundle foundation.IBundle) Nib {
	return NibClass.Alloc().InitWithNibDataBundle(nibData, bundle)
}

func (nc _NibClass) Alloc() Nib {
	rv := objc.Call[Nib](nc, objc.Sel("alloc"))
	return rv
}

func Nib_Alloc() Nib {
	return NibClass.Alloc()
}

func (nc _NibClass) New() Nib {
	rv := objc.Call[Nib](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNib() Nib {
	return NibClass.New()
}

func (n_ Nib) Init() Nib {
	rv := objc.Call[Nib](n_, objc.Sel("init"))
	return rv
}

// Instantiates objects in the nib file with the specified owner. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnib/1527173-instantiatewithowner?language=objc
func (n_ Nib) InstantiateWithOwnerTopLevelObjects(owner objc.IObject, topLevelObjects []objc.IObject) bool {
	rv := objc.Call[bool](n_, objc.Sel("instantiateWithOwner:topLevelObjects:"), owner, topLevelObjects)
	return rv
}
