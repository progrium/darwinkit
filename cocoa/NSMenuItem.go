package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

// NSMenuItem is a command item in an app menu.
// https://developer.apple.com/documentation/appkit/nsmenuitem?language=objc
type NSMenuItem struct {
	objc.Object
}

var nsMenuItem = objc.Get("NSMenuItem")

// NSMenuItem_Init returns an initialized instance of NSMenuItem.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514858-initwithtitle?language=objc
func NSMenuItem_Init(itemName string, action objc.Selector, keyEquivalent string) NSMenuItem {
	return NSMenuItem{nsMenuItem.Alloc().Send("initWithTitle:action:keyEquivalent:",
		core.String(itemName), action, core.String(keyEquivalent))}
}

// NSMenuItem_New returns an initialized instance of NSMenuItem.
func NSMenuItem_New() NSMenuItem {
	return NSMenuItem{nsMenuItem.Alloc().Init()}
}

// NSMenuItem_Separator returns a menu item that is used to separate logical groups of menu commands.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514838-separatoritem?language=objc
func NSMenuItem_Separator() NSMenuItem {
	return NSMenuItem{nsMenuItem.Get("separatorItem")}
}

// SetSubmenu sets the submenu of the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514845-submenu?language=objc
func (i NSMenuItem) SetSubmenu(submenu NSMenu) {
	i.Set("submenu:", submenu)
}

// Submenu returns the submenu of the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514845-submenu?language=objc
func (i NSMenuItem) Submenu() NSMenu {
	return NSMenu{i.Get("submenu")}
}

// Hidden returns a boolean value that indicates whether the menu item is hidden.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514846-hidden?language=objc
func (i NSMenuItem) Hidden() bool {
	return i.Get("hidden").Bool()
}

// SetHidden sets a boolean value that indicates whether the menu item is hidden.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514846-hidden?language=objc
func (i NSMenuItem) SetHidden(b bool) {
	i.Set("hidden:", b)
}

// Enabled returns a boolean value that indicates whether the menu item is enabled.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514863-enabled?language=objc
func (i NSMenuItem) Enabled() bool {
	return i.Send("isEnabled").Bool()
}

// SetEnabled sets a boolean value that indicates whether the menu item is enabled.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514863-enabled?language=objc
func (i NSMenuItem) SetEnabled(b bool) {
	i.Set("enabled:", b)
}

// Title returns the menu item's title.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514805-title?language=objc
func (i NSMenuItem) Title() string {
	return i.Get("title").String()
}

// SetTitle sets the menu item's title.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514805-title?language=objc
func (i NSMenuItem) SetTitle(s string) {
	i.Set("title:", core.String(s))
}

// Image returns the menu item’s image.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514819-image?language=objc
func (i NSMenuItem) Image() NSImage {
	return NSImage{i.Get("image")}
}

// SetImage sets the menu item’s image.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514819-image?language=objc
func (i NSMenuItem) SetImage(obj NSImage) {
	i.Set("image:", obj)
}

// ToolTip returns a help tag for the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514848-tooltip?language=objc
func (i NSMenuItem) ToolTip() string {
	return i.Get("toolTip").String()
}

// SetToolTip sets a help tag for the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514848-tooltip?language=objc
func (i NSMenuItem) SetToolTip(s string) {
	i.Set("toolTip:", core.String(s))
}

// Target returns the menu item's target.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514843-target?language=objc
func (i NSMenuItem) Target() objc.Object {
	return i.Get("target")
}

// SetTarget sets the menu item's target.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514843-target?language=objc
func (i NSMenuItem) SetTarget(obj objc.Object) {
	i.Set("target:", obj)
}

// Action returns the menu item's action-method selector.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514825-action?language=objc
func (i NSMenuItem) Action() objc.Selector {
	return objc.Sel(i.Get("action").String())
}

// SetAction sets the menu item's action-method selector.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514825-action?language=objc
func (i NSMenuItem) SetAction(sel objc.Selector) {
	i.Set("action:", sel)
}

// SetState sets the state of the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514804-state?language=objc
func (i NSMenuItem) SetState(state int) {
	i.Set("state:", state)
}

// State returns the state of the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514804-state?language=objc
func (i NSMenuItem) State() int64 {
	return i.Get("state").Int()
}
