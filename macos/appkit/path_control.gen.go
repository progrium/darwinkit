// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PathControl] class.
var PathControlClass = _PathControlClass{objc.GetClass("NSPathControl")}

type _PathControlClass struct {
	objc.Class
}

// An interface definition for the [PathControl] class.
type IPathControl interface {
	IControl
	SetDraggingSourceOperationMaskForLocal(mask DragOperation, isLocal bool)
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	IsEditable() bool
	SetEditable(value bool)
	ClickedPathItem() PathControlItem
	Delegate() PathControlDelegateWrapper
	SetDelegate(value PPathControlDelegate)
	SetDelegateObject(valueObject objc.IObject)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.IAttributedString)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	PathStyle() PathStyle
	SetPathStyle(value PathStyle)
	AllowedTypes() []string
	SetAllowedTypes(value []string)
	URL() foundation.URL
	SetURL(value foundation.IURL)
	PathItems() []PathControlItem
	SetPathItems(value []IPathControlItem)
	PlaceholderString() string
	SetPlaceholderString(value string)
}

// A display of a file system path or virtual path information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol?language=objc
type PathControl struct {
	Control
}

func PathControlFrom(ptr unsafe.Pointer) PathControl {
	return PathControl{
		Control: ControlFrom(ptr),
	}
}

func (pc _PathControlClass) Alloc() PathControl {
	rv := objc.Call[PathControl](pc, objc.Sel("alloc"))
	return rv
}

func PathControl_Alloc() PathControl {
	return PathControlClass.Alloc()
}

func (pc _PathControlClass) New() PathControl {
	rv := objc.Call[PathControl](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPathControl() PathControl {
	return PathControlClass.New()
}

func (p_ PathControl) Init() PathControl {
	rv := objc.Call[PathControl](p_, objc.Sel("init"))
	return rv
}

func (p_ PathControl) InitWithFrame(frameRect foundation.Rect) PathControl {
	rv := objc.Call[PathControl](p_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func PathControl_InitWithFrame(frameRect foundation.Rect) PathControl {
	return PathControlClass.Alloc().InitWithFrame(frameRect)
}

// Configures the default value returned from draggingSourceOperationMaskForLocal:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1526909-setdraggingsourceoperationmask?language=objc
func (p_ PathControl) SetDraggingSourceOperationMaskForLocal(mask DragOperation, isLocal bool) {
	objc.Call[objc.Void](p_, objc.Sel("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}

// The receiver’s double-click action method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1534088-doubleaction?language=objc
func (p_ PathControl) DoubleAction() objc.Selector {
	rv := objc.Call[objc.Selector](p_, objc.Sel("doubleAction"))
	return rv
}

// The receiver’s double-click action method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1534088-doubleaction?language=objc
func (p_ PathControl) SetDoubleAction(value objc.Selector) {
	objc.Call[objc.Void](p_, objc.Sel("setDoubleAction:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1535833-editable?language=objc
func (p_ PathControl) IsEditable() bool {
	rv := objc.Call[bool](p_, objc.Sel("isEditable"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1535833-editable?language=objc
func (p_ PathControl) SetEditable(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setEditable:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1535047-clickedpathitem?language=objc
func (p_ PathControl) ClickedPathItem() PathControlItem {
	rv := objc.Call[PathControlItem](p_, objc.Sel("clickedPathItem"))
	return rv
}

// The receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1526753-delegate?language=objc
func (p_ PathControl) Delegate() PathControlDelegateWrapper {
	rv := objc.Call[PathControlDelegateWrapper](p_, objc.Sel("delegate"))
	return rv
}

// The receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1526753-delegate?language=objc
func (p_ PathControl) SetDelegate(value PPathControlDelegate) {
	po0 := objc.WrapAsProtocol("NSPathControlDelegate", value)
	objc.SetAssociatedObject(p_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), po0)
}

// The receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1526753-delegate?language=objc
func (p_ PathControl) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1531486-placeholderattributedstring?language=objc
func (p_ PathControl) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](p_, objc.Sel("placeholderAttributedString"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1531486-placeholderattributedstring?language=objc
func (p_ PathControl) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.Call[objc.Void](p_, objc.Sel("setPlaceholderAttributedString:"), objc.Ptr(value))
}

// The receiver’s background color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1534164-backgroundcolor?language=objc
func (p_ PathControl) BackgroundColor() Color {
	rv := objc.Call[Color](p_, objc.Sel("backgroundColor"))
	return rv
}

// The receiver’s background color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1534164-backgroundcolor?language=objc
func (p_ PathControl) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](p_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// The receiver’s path style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1532330-pathstyle?language=objc
func (p_ PathControl) PathStyle() PathStyle {
	rv := objc.Call[PathStyle](p_, objc.Sel("pathStyle"))
	return rv
}

// The receiver’s path style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1532330-pathstyle?language=objc
func (p_ PathControl) SetPathStyle(value PathStyle) {
	objc.Call[objc.Void](p_, objc.Sel("setPathStyle:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1527415-allowedtypes?language=objc
func (p_ PathControl) AllowedTypes() []string {
	rv := objc.Call[[]string](p_, objc.Sel("allowedTypes"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1527415-allowedtypes?language=objc
func (p_ PathControl) SetAllowedTypes(value []string) {
	objc.Call[objc.Void](p_, objc.Sel("setAllowedTypes:"), value)
}

// The path value displayed by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1527205-url?language=objc
func (p_ PathControl) URL() foundation.URL {
	rv := objc.Call[foundation.URL](p_, objc.Sel("URL"))
	return rv
}

// The path value displayed by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1527205-url?language=objc
func (p_ PathControl) SetURL(value foundation.IURL) {
	objc.Call[objc.Void](p_, objc.Sel("setURL:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1528208-pathitems?language=objc
func (p_ PathControl) PathItems() []PathControlItem {
	rv := objc.Call[[]PathControlItem](p_, objc.Sel("pathItems"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1528208-pathitems?language=objc
func (p_ PathControl) SetPathItems(value []IPathControlItem) {
	objc.Call[objc.Void](p_, objc.Sel("setPathItems:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1531787-placeholderstring?language=objc
func (p_ PathControl) PlaceholderString() string {
	rv := objc.Call[string](p_, objc.Sel("placeholderString"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/1531787-placeholderstring?language=objc
func (p_ PathControl) SetPlaceholderString(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setPlaceholderString:"), value)
}
