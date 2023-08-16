// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods implemented by plug-ins that allow an app’s Dock tile to be customized while the app is not running. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktileplugin?language=objc
type PDockTilePlugIn interface {
	// optional
	SetDockTile(dockTile DockTile)
	HasSetDockTile() bool

	// optional
	DockMenu() IMenu
	HasDockMenu() bool
}

// A concrete type wrapper for the [PDockTilePlugIn] protocol.
type DockTilePlugInWrapper struct {
	objc.Object
}

func (d_ DockTilePlugInWrapper) HasSetDockTile() bool {
	return d_.RespondsToSelector(objc.Sel("setDockTile:"))
}

// Invoked when the plug-in is first loaded and when the application is removed from the Dock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktileplugin/1534120-setdocktile?language=objc
func (d_ DockTilePlugInWrapper) SetDockTile(dockTile IDockTile) {
	objc.Call[objc.Void](d_, objc.Sel("setDockTile:"), objc.Ptr(dockTile))
}

func (d_ DockTilePlugInWrapper) HasDockMenu() bool {
	return d_.RespondsToSelector(objc.Sel("dockMenu"))
}

// Invoked when the user causes the application's dock menu to be shown. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktileplugin/1527547-dockmenu?language=objc
func (d_ DockTilePlugInWrapper) DockMenu() Menu {
	rv := objc.Call[Menu](d_, objc.Sel("dockMenu"))
	return rv
}
