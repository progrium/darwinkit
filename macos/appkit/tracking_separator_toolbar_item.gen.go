// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TrackingSeparatorToolbarItem] class.
var TrackingSeparatorToolbarItemClass = _TrackingSeparatorToolbarItemClass{objc.GetClass("NSTrackingSeparatorToolbarItem")}

type _TrackingSeparatorToolbarItemClass struct {
	objc.Class
}

// An interface definition for the [TrackingSeparatorToolbarItem] class.
type ITrackingSeparatorToolbarItem interface {
	IToolbarItem
	SplitView() SplitView
	SetSplitView(value ISplitView)
	DividerIndex() int
	SetDividerIndex(value int)
}

// A toolbar separator that aligns with the vertical split view in the same window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingseparatortoolbaritem?language=objc
type TrackingSeparatorToolbarItem struct {
	ToolbarItem
}

func TrackingSeparatorToolbarItemFrom(ptr unsafe.Pointer) TrackingSeparatorToolbarItem {
	return TrackingSeparatorToolbarItem{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}

func (tc _TrackingSeparatorToolbarItemClass) TrackingSeparatorToolbarItemWithIdentifierSplitViewDividerIndex(identifier ToolbarItemIdentifier, splitView ISplitView, dividerIndex int) TrackingSeparatorToolbarItem {
	rv := objc.Call[TrackingSeparatorToolbarItem](tc, objc.Sel("trackingSeparatorToolbarItemWithIdentifier:splitView:dividerIndex:"), identifier, objc.Ptr(splitView), dividerIndex)
	return rv
}

// Creates a new tracking separator toolbar item and configures it to align with the divider of the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingseparatortoolbaritem/3626828-trackingseparatortoolbaritemwith?language=objc
func TrackingSeparatorToolbarItem_TrackingSeparatorToolbarItemWithIdentifierSplitViewDividerIndex(identifier ToolbarItemIdentifier, splitView ISplitView, dividerIndex int) TrackingSeparatorToolbarItem {
	return TrackingSeparatorToolbarItemClass.TrackingSeparatorToolbarItemWithIdentifierSplitViewDividerIndex(identifier, splitView, dividerIndex)
}

func (tc _TrackingSeparatorToolbarItemClass) Alloc() TrackingSeparatorToolbarItem {
	rv := objc.Call[TrackingSeparatorToolbarItem](tc, objc.Sel("alloc"))
	return rv
}

func TrackingSeparatorToolbarItem_Alloc() TrackingSeparatorToolbarItem {
	return TrackingSeparatorToolbarItemClass.Alloc()
}

func (tc _TrackingSeparatorToolbarItemClass) New() TrackingSeparatorToolbarItem {
	rv := objc.Call[TrackingSeparatorToolbarItem](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTrackingSeparatorToolbarItem() TrackingSeparatorToolbarItem {
	return TrackingSeparatorToolbarItemClass.New()
}

func (t_ TrackingSeparatorToolbarItem) Init() TrackingSeparatorToolbarItem {
	rv := objc.Call[TrackingSeparatorToolbarItem](t_, objc.Sel("init"))
	return rv
}

func (t_ TrackingSeparatorToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) TrackingSeparatorToolbarItem {
	rv := objc.Call[TrackingSeparatorToolbarItem](t_, objc.Sel("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

// Creates a toolbar item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1534084-initwithitemidentifier?language=objc
func TrackingSeparatorToolbarItem_InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) TrackingSeparatorToolbarItem {
	return TrackingSeparatorToolbarItemClass.Alloc().InitWithItemIdentifier(itemIdentifier)
}

// The vertical split view to align with the toolbar separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingseparatortoolbaritem/3626827-splitview?language=objc
func (t_ TrackingSeparatorToolbarItem) SplitView() SplitView {
	rv := objc.Call[SplitView](t_, objc.Sel("splitView"))
	return rv
}

// The vertical split view to align with the toolbar separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingseparatortoolbaritem/3626827-splitview?language=objc
func (t_ TrackingSeparatorToolbarItem) SetSplitView(value ISplitView) {
	objc.Call[objc.Void](t_, objc.Sel("setSplitView:"), objc.Ptr(value))
}

// The index of the split view divider to align with the tracking separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingseparatortoolbaritem/3626826-dividerindex?language=objc
func (t_ TrackingSeparatorToolbarItem) DividerIndex() int {
	rv := objc.Call[int](t_, objc.Sel("dividerIndex"))
	return rv
}

// The index of the split view divider to align with the tracking separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingseparatortoolbaritem/3626826-dividerindex?language=objc
func (t_ TrackingSeparatorToolbarItem) SetDividerIndex(value int) {
	objc.Call[objc.Void](t_, objc.Sel("setDividerIndex:"), value)
}
