// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Drawer] class.
var DrawerClass = _DrawerClass{objc.GetClass("NSDrawer")}

type _DrawerClass struct {
	objc.Class
}

// An interface definition for the [Drawer] class.
type IDrawer interface {
	IResponder
}

// A user interface element that contains and displays text, scroll, and browser views, in addition to other view subclasses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer?language=objc
type Drawer struct {
	Responder
}

func DrawerFrom(ptr unsafe.Pointer) Drawer {
	return Drawer{
		Responder: ResponderFrom(ptr),
	}
}

func (dc _DrawerClass) Alloc() Drawer {
	rv := objc.Call[Drawer](dc, objc.Sel("alloc"))
	return rv
}

func Drawer_Alloc() Drawer {
	return DrawerClass.Alloc()
}

func (dc _DrawerClass) New() Drawer {
	rv := objc.Call[Drawer](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDrawer() Drawer {
	return DrawerClass.New()
}

func (d_ Drawer) Init() Drawer {
	rv := objc.Call[Drawer](d_, objc.Sel("init"))
	return rv
}
