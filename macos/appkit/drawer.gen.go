// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var DrawerClass = _DrawerClass{objc.GetClass("NSDrawer")}

type _DrawerClass struct {
	objc.Class
}

type IDrawer interface {
	IResponder
}

type Drawer struct {
	Responder
}

func MakeDrawer(ptr unsafe.Pointer) Drawer {
	return Drawer{
		Responder: MakeResponder(ptr),
	}
}

func (d_ Drawer) InitWithContentSizePreferredEdge(contentSize foundation.Size, edge foundation.RectEdge) Drawer {
	rv := objc.CallMethod[Drawer](d_, objc.GetSelector("initWithContentSize:preferredEdge:"), contentSize, edge)
	return rv
}

func Drawer_InitWithContentSizePreferredEdge(contentSize foundation.Size, edge foundation.RectEdge) Drawer {
	return DrawerClass.Alloc().InitWithContentSizePreferredEdge(contentSize, edge)
}

func (d_ Drawer) Init() Drawer {
	rv := objc.CallMethod[Drawer](d_, objc.GetSelector("init"))
	return rv
}

func Drawer_Init() Drawer {
	return DrawerClass.Alloc().Init()
}

func (dc _DrawerClass) Alloc() Drawer {
	rv := objc.CallMethod[Drawer](dc, objc.GetSelector("alloc"))
	return rv
}

func Drawer_Alloc() Drawer {
	return DrawerClass.Alloc()
}

func (dc _DrawerClass) New() Drawer {
	rv := objc.CallMethod[Drawer](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDrawer() Drawer {
	return DrawerClass.New()
}

func Drawer_New() Drawer {
	return DrawerClass.New()
}
