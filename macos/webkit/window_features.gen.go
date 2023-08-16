// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WindowFeatures] class.
var WindowFeaturesClass = _WindowFeaturesClass{objc.GetClass("WKWindowFeatures")}

type _WindowFeaturesClass struct {
	objc.Class
}

// An interface definition for the [WindowFeatures] class.
type IWindowFeatures interface {
	objc.IObject
	AllowsResizing() foundation.Number
	Width() foundation.Number
	X() foundation.Number
	SetX(value foundation.INumber)
	Y() foundation.Number
	SetY(value foundation.INumber)
	StatusBarVisibility() foundation.Number
	Height() foundation.Number
	ToolbarsVisibility() foundation.Number
	MenuBarVisibility() foundation.Number
}

// Display-related attributes that a webpage requests for its window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures?language=objc
type WindowFeatures struct {
	objc.Object
}

func WindowFeaturesFrom(ptr unsafe.Pointer) WindowFeatures {
	return WindowFeatures{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WindowFeaturesClass) Alloc() WindowFeatures {
	rv := objc.Call[WindowFeatures](wc, objc.Sel("alloc"))
	return rv
}

func WindowFeatures_Alloc() WindowFeatures {
	return WindowFeaturesClass.Alloc()
}

func (wc _WindowFeaturesClass) New() WindowFeatures {
	rv := objc.Call[WindowFeatures](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWindowFeatures() WindowFeatures {
	return WindowFeaturesClass.New()
}

func (w_ WindowFeatures) Init() WindowFeatures {
	rv := objc.Call[WindowFeatures](w_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether to make the containing window window resizable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/1536871-allowsresizing?language=objc
func (w_ WindowFeatures) AllowsResizing() foundation.Number {
	rv := objc.Call[foundation.Number](w_, objc.Sel("allowsResizing"))
	return rv
}

// The requested width of the containing window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/1537562-width?language=objc
func (w_ WindowFeatures) Width() foundation.Number {
	rv := objc.Call[foundation.Number](w_, objc.Sel("width"))
	return rv
}

// The requested x-coordinate of the containing window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/1537705-x?language=objc
func (w_ WindowFeatures) X() foundation.Number {
	rv := objc.Call[foundation.Number](w_, objc.Sel("x"))
	return rv
}

// The requested x-coordinate of the containing window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/1537705-x?language=objc
func (w_ WindowFeatures) SetX(value foundation.INumber) {
	objc.Call[objc.Void](w_, objc.Sel("setX:"), objc.Ptr(value))
}

// The requested y-coordinate of the containing window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/1537052-y?language=objc
func (w_ WindowFeatures) Y() foundation.Number {
	rv := objc.Call[foundation.Number](w_, objc.Sel("y"))
	return rv
}

// The requested y-coordinate of the containing window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/1537052-y?language=objc
func (w_ WindowFeatures) SetY(value foundation.INumber) {
	objc.Call[objc.Void](w_, objc.Sel("setY:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the webpage requested a visible status bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/1536638-statusbarvisibility?language=objc
func (w_ WindowFeatures) StatusBarVisibility() foundation.Number {
	rv := objc.Call[foundation.Number](w_, objc.Sel("statusBarVisibility"))
	return rv
}

// The requested height of the containing window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/1536826-height?language=objc
func (w_ WindowFeatures) Height() foundation.Number {
	rv := objc.Call[foundation.Number](w_, objc.Sel("height"))
	return rv
}

// A Boolean value that indicates whether the webpage requested a visible toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/1536218-toolbarsvisibility?language=objc
func (w_ WindowFeatures) ToolbarsVisibility() foundation.Number {
	rv := objc.Call[foundation.Number](w_, objc.Sel("toolbarsVisibility"))
	return rv
}

// A Boolean value that indicates whether the webpage requests a visible menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/1538001-menubarvisibility?language=objc
func (w_ WindowFeatures) MenuBarVisibility() foundation.Number {
	rv := objc.Call[foundation.Number](w_, objc.Sel("menuBarVisibility"))
	return rv
}
