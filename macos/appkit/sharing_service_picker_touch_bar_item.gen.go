// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SharingServicePickerTouchBarItem] class.
var SharingServicePickerTouchBarItemClass = _SharingServicePickerTouchBarItemClass{objc.GetClass("NSSharingServicePickerTouchBarItem")}

type _SharingServicePickerTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [SharingServicePickerTouchBarItem] class.
type ISharingServicePickerTouchBarItem interface {
	ITouchBarItem
	ButtonImage() Image
	SetButtonImage(value IImage)
	ButtonTitle() string
	SetButtonTitle(value string)
	Delegate() SharingServicePickerTouchBarItemDelegateWrapper
	SetDelegate(value PSharingServicePickerTouchBarItemDelegate)
	SetDelegateObject(valueObject objc.IObject)
	IsEnabled() bool
	SetEnabled(value bool)
}

// A bar item that, along with its delegate, provides a list of objects eligible for sharing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem?language=objc
type SharingServicePickerTouchBarItem struct {
	TouchBarItem
}

func SharingServicePickerTouchBarItemFrom(ptr unsafe.Pointer) SharingServicePickerTouchBarItem {
	return SharingServicePickerTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

func (sc _SharingServicePickerTouchBarItemClass) Alloc() SharingServicePickerTouchBarItem {
	rv := objc.Call[SharingServicePickerTouchBarItem](sc, objc.Sel("alloc"))
	return rv
}

func SharingServicePickerTouchBarItem_Alloc() SharingServicePickerTouchBarItem {
	return SharingServicePickerTouchBarItemClass.Alloc()
}

func (sc _SharingServicePickerTouchBarItemClass) New() SharingServicePickerTouchBarItem {
	rv := objc.Call[SharingServicePickerTouchBarItem](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSharingServicePickerTouchBarItem() SharingServicePickerTouchBarItem {
	return SharingServicePickerTouchBarItemClass.New()
}

func (s_ SharingServicePickerTouchBarItem) Init() SharingServicePickerTouchBarItem {
	rv := objc.Call[SharingServicePickerTouchBarItem](s_, objc.Sel("init"))
	return rv
}

func (s_ SharingServicePickerTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) SharingServicePickerTouchBarItem {
	rv := objc.Call[SharingServicePickerTouchBarItem](s_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Creates a new item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544755-initwithidentifier?language=objc
func SharingServicePickerTouchBarItem_InitWithIdentifier(identifier TouchBarItemIdentifier) SharingServicePickerTouchBarItem {
	return SharingServicePickerTouchBarItemClass.Alloc().InitWithIdentifier(identifier)
}

// The image displayed in the sharing service picker item button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/2646965-buttonimage?language=objc
func (s_ SharingServicePickerTouchBarItem) ButtonImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("buttonImage"))
	return rv
}

// The image displayed in the sharing service picker item button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/2646965-buttonimage?language=objc
func (s_ SharingServicePickerTouchBarItem) SetButtonImage(value IImage) {
	objc.Call[objc.Void](s_, objc.Sel("setButtonImage:"), objc.Ptr(value))
}

// The text displayed in the sharing service picker item button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/2646971-buttontitle?language=objc
func (s_ SharingServicePickerTouchBarItem) ButtonTitle() string {
	rv := objc.Call[string](s_, objc.Sel("buttonTitle"))
	return rv
}

// The text displayed in the sharing service picker item button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/2646971-buttontitle?language=objc
func (s_ SharingServicePickerTouchBarItem) SetButtonTitle(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setButtonTitle:"), value)
}

// The object that acts as the delegate of the sharing service picker bar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/2539408-delegate?language=objc
func (s_ SharingServicePickerTouchBarItem) Delegate() SharingServicePickerTouchBarItemDelegateWrapper {
	rv := objc.Call[SharingServicePickerTouchBarItemDelegateWrapper](s_, objc.Sel("delegate"))
	return rv
}

// The object that acts as the delegate of the sharing service picker bar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/2539408-delegate?language=objc
func (s_ SharingServicePickerTouchBarItem) SetDelegate(value PSharingServicePickerTouchBarItemDelegate) {
	po0 := objc.WrapAsProtocol("NSSharingServicePickerTouchBarItemDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), po0)
}

// The object that acts as the delegate of the sharing service picker bar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/2539408-delegate?language=objc
func (s_ SharingServicePickerTouchBarItem) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// A Boolean value that specifies whether the sharing service picker item is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/2560996-enabled?language=objc
func (s_ SharingServicePickerTouchBarItem) IsEnabled() bool {
	rv := objc.Call[bool](s_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value that specifies whether the sharing service picker item is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/2560996-enabled?language=objc
func (s_ SharingServicePickerTouchBarItem) SetEnabled(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setEnabled:"), value)
}
