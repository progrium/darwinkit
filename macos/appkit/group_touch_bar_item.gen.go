// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GroupTouchBarItem] class.
var GroupTouchBarItemClass = _GroupTouchBarItemClass{objc.GetClass("NSGroupTouchBarItem")}

type _GroupTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [GroupTouchBarItem] class.
type IGroupTouchBarItem interface {
	ITouchBarItem
	SetCustomizationLabel(value string)
	PreferredItemWidth() float64
	SetPreferredItemWidth(value float64)
	EffectiveCompressionOptions() UserInterfaceCompressionOptions
	PrioritizedCompressionOptions() []UserInterfaceCompressionOptions
	SetPrioritizedCompressionOptions(value []IUserInterfaceCompressionOptions)
	GroupUserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetGroupUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	GroupTouchBar() TouchBar
	SetGroupTouchBar(value ITouchBar)
	PrefersEqualWidths() bool
	SetPrefersEqualWidths(value bool)
}

// A bar item that provides a bar to contain other items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem?language=objc
type GroupTouchBarItem struct {
	TouchBarItem
}

func GroupTouchBarItemFrom(ptr unsafe.Pointer) GroupTouchBarItem {
	return GroupTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

func (gc _GroupTouchBarItemClass) GroupItemWithIdentifierItems(identifier TouchBarItemIdentifier, items []ITouchBarItem) GroupTouchBarItem {
	rv := objc.Call[GroupTouchBarItem](gc, objc.Sel("groupItemWithIdentifier:items:"), identifier, items)
	return rv
}

// Initializes and returns a group item whose bar is constructed from the supplied items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/2544686-groupitemwithidentifier?language=objc
func GroupTouchBarItem_GroupItemWithIdentifierItems(identifier TouchBarItemIdentifier, items []ITouchBarItem) GroupTouchBarItem {
	return GroupTouchBarItemClass.GroupItemWithIdentifierItems(identifier, items)
}

func (gc _GroupTouchBarItemClass) AlertStyleGroupItemWithIdentifier(identifier TouchBarItemIdentifier) GroupTouchBarItem {
	rv := objc.Call[GroupTouchBarItem](gc, objc.Sel("alertStyleGroupItemWithIdentifier:"), identifier)
	return rv
}

// Initializes and returns a group item configured to match system alerts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/2909967-alertstylegroupitemwithidentifie?language=objc
func GroupTouchBarItem_AlertStyleGroupItemWithIdentifier(identifier TouchBarItemIdentifier) GroupTouchBarItem {
	return GroupTouchBarItemClass.AlertStyleGroupItemWithIdentifier(identifier)
}

func (gc _GroupTouchBarItemClass) Alloc() GroupTouchBarItem {
	rv := objc.Call[GroupTouchBarItem](gc, objc.Sel("alloc"))
	return rv
}

func GroupTouchBarItem_Alloc() GroupTouchBarItem {
	return GroupTouchBarItemClass.Alloc()
}

func (gc _GroupTouchBarItemClass) New() GroupTouchBarItem {
	rv := objc.Call[GroupTouchBarItem](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGroupTouchBarItem() GroupTouchBarItem {
	return GroupTouchBarItemClass.New()
}

func (g_ GroupTouchBarItem) Init() GroupTouchBarItem {
	rv := objc.Call[GroupTouchBarItem](g_, objc.Sel("init"))
	return rv
}

func (g_ GroupTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) GroupTouchBarItem {
	rv := objc.Call[GroupTouchBarItem](g_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Creates a new item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544755-initwithidentifier?language=objc
func NewGroupTouchBarItemWithIdentifier(identifier TouchBarItemIdentifier) GroupTouchBarItem {
	instance := GroupTouchBarItemClass.Alloc().InitWithIdentifier(identifier)
	instance.Autorelease()
	return instance
}

// The user-visible string identifying this item during bar customization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/2544877-customizationlabel?language=objc
func (g_ GroupTouchBarItem) SetCustomizationLabel(value string) {
	objc.Call[objc.Void](g_, objc.Sel("setCustomizationLabel:"), value)
}

// The preferred width for items in the group when prefersEqualWidths is YES. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/2909975-preferreditemwidth?language=objc
func (g_ GroupTouchBarItem) PreferredItemWidth() float64 {
	rv := objc.Call[float64](g_, objc.Sel("preferredItemWidth"))
	return rv
}

// The preferred width for items in the group when prefersEqualWidths is YES. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/2909975-preferreditemwidth?language=objc
func (g_ GroupTouchBarItem) SetPreferredItemWidth(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setPreferredItemWidth:"), value)
}

// The compression options that are currently active on the group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/2909985-effectivecompressionoptions?language=objc
func (g_ GroupTouchBarItem) EffectiveCompressionOptions() UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](g_, objc.Sel("effectiveCompressionOptions"))
	return rv
}

// The allowed compression options, in the order they should be applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/2909968-prioritizedcompressionoptions?language=objc
func (g_ GroupTouchBarItem) PrioritizedCompressionOptions() []UserInterfaceCompressionOptions {
	rv := objc.Call[[]UserInterfaceCompressionOptions](g_, objc.Sel("prioritizedCompressionOptions"))
	return rv
}

// The allowed compression options, in the order they should be applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/2909968-prioritizedcompressionoptions?language=objc
func (g_ GroupTouchBarItem) SetPrioritizedCompressionOptions(value []IUserInterfaceCompressionOptions) {
	objc.Call[objc.Void](g_, objc.Sel("setPrioritizedCompressionOptions:"), value)
}

// The user interface direction that controls the layout order of the items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/2876344-groupuserinterfacelayoutdirectio?language=objc
func (g_ GroupTouchBarItem) GroupUserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Call[UserInterfaceLayoutDirection](g_, objc.Sel("groupUserInterfaceLayoutDirection"))
	return rv
}

// The user interface direction that controls the layout order of the items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/2876344-groupuserinterfacelayoutdirectio?language=objc
func (g_ GroupTouchBarItem) SetGroupUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Call[objc.Void](g_, objc.Sel("setGroupUserInterfaceLayoutDirection:"), value)
}

// A bar that holds this group's items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/2544717-grouptouchbar?language=objc
func (g_ GroupTouchBarItem) GroupTouchBar() TouchBar {
	rv := objc.Call[TouchBar](g_, objc.Sel("groupTouchBar"))
	return rv
}

// A bar that holds this group's items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/2544717-grouptouchbar?language=objc
func (g_ GroupTouchBarItem) SetGroupTouchBar(value ITouchBar) {
	objc.Call[objc.Void](g_, objc.Sel("setGroupTouchBar:"), objc.Ptr(value))
}

// A Boolean value that specifies that items should have equal widths when possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/2909972-prefersequalwidths?language=objc
func (g_ GroupTouchBarItem) PrefersEqualWidths() bool {
	rv := objc.Call[bool](g_, objc.Sel("prefersEqualWidths"))
	return rv
}

// A Boolean value that specifies that items should have equal widths when possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/2909972-prefersequalwidths?language=objc
func (g_ GroupTouchBarItem) SetPrefersEqualWidths(value bool) {
	objc.Call[objc.Void](g_, objc.Sel("setPrefersEqualWidths:"), value)
}
