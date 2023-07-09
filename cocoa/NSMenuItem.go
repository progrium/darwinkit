package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

// NSMenuItem is a command item in an app menu.
// https://developer.apple.com/documentation/appkit/nsmenuitem?language=objc
type NSMenuItem struct {
	gen_NSMenuItem
}

// NSMenuItem_Init returns an initialized instance of NSMenuItem.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514858-initwithtitle?language=objc
func NSMenuItem_Init(itemName string, action objc.Selector, keyEquivalent string) NSMenuItem {
	return NSMenuItem_Alloc().InitWithTitleActionKeyEquivalent(
		core.String(itemName), action, core.String(keyEquivalent))
}

// NSMenuItem_New returns an initialized instance of NSMenuItem.
func NSMenuItem_New() NSMenuItem {
	return NSMenuItem_Alloc().Init()
}

// NSMenuItem_Separator returns a menu item that is used to separate logical groups of menu commands.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514838-separatoritem?language=objc
func NSMenuItem_Separator() NSMenuItem {
	return NSMenuItem_SeparatorItem()
}

// Hidden returns a boolean value that indicates whether the menu item is hidden.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514846-hidden?language=objc
func (i NSMenuItem) Hidden() bool {
	return i.IsHidden()
}

// Enabled returns a boolean value that indicates whether the menu item is enabled.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514863-enabled?language=objc
func (i NSMenuItem) Enabled() bool {
	return i.IsEnabled()
}

// Title returns the menu item's title.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514805-title?language=objc
func (i NSMenuItem) Title() string {
	return i.gen_NSMenuItem.Title().String()
}

// SetTitle sets the menu item's title.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514805-title?language=objc
func (i NSMenuItem) SetTitle(s string) {
	i.gen_NSMenuItem.SetTitle(core.String(s))
}

// SetAttributedTitle sets a custom string for the menu item's title.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514860-attributedtitle?language=objc
func (i NSMenuItem) SetAttributedTitle(s string) {
	i.gen_NSMenuItem.SetAttributedTitle(core.NSAttributedString_FromString(s))
}

// ToolTip returns a help tag for the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514848-tooltip?language=objc
func (i NSMenuItem) ToolTip() string {
	return i.gen_NSMenuItem.ToolTip().String()
}

// SetToolTip sets a help tag for the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514848-tooltip?language=objc
func (i NSMenuItem) SetToolTip(s string) {
	i.gen_NSMenuItem.SetToolTip(core.String(s))
}

// Target returns the menu item's target.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514843-target?language=objc
func (i NSMenuItem) Target() objc.Object {
	return i.gen_NSMenuItem.Target()
}

// Action returns the menu item's action-method selector.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514825-action?language=objc
func (i NSMenuItem) Action() objc.Selector {
	return objc.Sel(i.Get("action").String())
}

// SetState sets the state of the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514804-state?language=objc
func (i NSMenuItem) SetState(state int) {
	i.gen_NSMenuItem.SetState(core.NSInteger(state))
}

// State returns the state of the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514804-state?language=objc
func (i NSMenuItem) State() int {
	return int(i.gen_NSMenuItem.State())
}
