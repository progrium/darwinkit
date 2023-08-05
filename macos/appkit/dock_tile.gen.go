// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var DockTileClass = _DockTileClass{objc.GetClass("NSDockTile")}

type _DockTileClass struct {
	objc.Class
}

type IDockTile interface {
	objc.IObject
	Display()
	ContentView() View
	SetContentView(value IView)
	Size() foundation.Size
	Owner() objc.Object
	ShowsApplicationBadge() bool
	SetShowsApplicationBadge(value bool)
	BadgeLabel() string
	SetBadgeLabel(value string)
}

type DockTile struct {
	objc.Object
}

func MakeDockTile(ptr unsafe.Pointer) DockTile {
	return DockTile{
		Object: objc.MakeObject(ptr),
	}
}

func (dc _DockTileClass) Alloc() DockTile {
	rv := objc.CallMethod[DockTile](dc, objc.GetSelector("alloc"))
	return rv
}

func DockTile_Alloc() DockTile {
	return DockTileClass.Alloc()
}

func (dc _DockTileClass) New() DockTile {
	rv := objc.CallMethod[DockTile](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDockTile() DockTile {
	return DockTileClass.New()
}

func DockTile_New() DockTile {
	return DockTileClass.New()
}

func (d_ DockTile) Init() DockTile {
	rv := objc.CallMethod[DockTile](d_, objc.GetSelector("init"))
	return rv
}

func DockTile_Init() DockTile {
	return DockTileClass.Alloc().Init()
}

func (d_ DockTile) Display() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("display"))
}

func (d_ DockTile) ContentView() View {
	rv := objc.CallMethod[View](d_, objc.GetSelector("contentView"))
	return rv
}

func (d_ DockTile) SetContentView(value IView) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setContentView:"), objc.ExtractPtr(value))
}

func (d_ DockTile) Size() foundation.Size {
	rv := objc.CallMethod[foundation.Size](d_, objc.GetSelector("size"))
	return rv
}

func (d_ DockTile) Owner() objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("owner"))
	return rv
}

func (d_ DockTile) ShowsApplicationBadge() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("showsApplicationBadge"))
	return rv
}

func (d_ DockTile) SetShowsApplicationBadge(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setShowsApplicationBadge:"), value)
}

func (d_ DockTile) BadgeLabel() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("badgeLabel"))
	return rv
}

func (d_ DockTile) SetBadgeLabel(value string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setBadgeLabel:"), value)
}
