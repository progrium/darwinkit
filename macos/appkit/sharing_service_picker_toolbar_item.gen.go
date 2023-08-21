// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SharingServicePickerToolbarItem] class.
var SharingServicePickerToolbarItemClass = _SharingServicePickerToolbarItemClass{objc.GetClass("NSSharingServicePickerToolbarItem")}

type _SharingServicePickerToolbarItemClass struct {
	objc.Class
}

// An interface definition for the [SharingServicePickerToolbarItem] class.
type ISharingServicePickerToolbarItem interface {
	IToolbarItem
	Delegate() SharingServicePickerToolbarItemDelegateWrapper
	SetDelegate(value PSharingServicePickerToolbarItemDelegate)
	SetDelegateObject(valueObject objc.IObject)
}

// A toolbar item that displays the macOS share sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertoolbaritem?language=objc
type SharingServicePickerToolbarItem struct {
	ToolbarItem
}

func SharingServicePickerToolbarItemFrom(ptr unsafe.Pointer) SharingServicePickerToolbarItem {
	return SharingServicePickerToolbarItem{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}

func (sc _SharingServicePickerToolbarItemClass) Alloc() SharingServicePickerToolbarItem {
	rv := objc.Call[SharingServicePickerToolbarItem](sc, objc.Sel("alloc"))
	return rv
}

func SharingServicePickerToolbarItem_Alloc() SharingServicePickerToolbarItem {
	return SharingServicePickerToolbarItemClass.Alloc()
}

func (sc _SharingServicePickerToolbarItemClass) New() SharingServicePickerToolbarItem {
	rv := objc.Call[SharingServicePickerToolbarItem](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSharingServicePickerToolbarItem() SharingServicePickerToolbarItem {
	return SharingServicePickerToolbarItemClass.New()
}

func (s_ SharingServicePickerToolbarItem) Init() SharingServicePickerToolbarItem {
	rv := objc.Call[SharingServicePickerToolbarItem](s_, objc.Sel("init"))
	return rv
}

func (s_ SharingServicePickerToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) SharingServicePickerToolbarItem {
	rv := objc.Call[SharingServicePickerToolbarItem](s_, objc.Sel("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

// Creates a toolbar item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1534084-initwithitemidentifier?language=objc
func NewSharingServicePickerToolbarItemWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) SharingServicePickerToolbarItem {
	instance := SharingServicePickerToolbarItemClass.Alloc().InitWithItemIdentifier(itemIdentifier)
	instance.Autorelease()
	return instance
}

// The custom object from your app that provides the items to share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertoolbaritem/3365981-delegate?language=objc
func (s_ SharingServicePickerToolbarItem) Delegate() SharingServicePickerToolbarItemDelegateWrapper {
	rv := objc.Call[SharingServicePickerToolbarItemDelegateWrapper](s_, objc.Sel("delegate"))
	return rv
}

// The custom object from your app that provides the items to share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertoolbaritem/3365981-delegate?language=objc
func (s_ SharingServicePickerToolbarItem) SetDelegate(value PSharingServicePickerToolbarItemDelegate) {
	po0 := objc.WrapAsProtocol("NSSharingServicePickerToolbarItemDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), po0)
}

// The custom object from your app that provides the items to share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertoolbaritem/3365981-delegate?language=objc
func (s_ SharingServicePickerToolbarItem) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}
