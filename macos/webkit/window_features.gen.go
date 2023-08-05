// AUTO-GENERATED CODE, DO NOT MODIFY
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var WindowFeaturesClass = _WindowFeaturesClass{objc.GetClass("WKWindowFeatures")}

type _WindowFeaturesClass struct {
	objc.Class
}

type IWindowFeatures interface {
	objc.IObject
	AllowsResizing() foundation.Number
	Height() foundation.Number
	Width() foundation.Number
	X() foundation.Number
	Y() foundation.Number
	MenuBarVisibility() foundation.Number
	StatusBarVisibility() foundation.Number
	ToolbarsVisibility() foundation.Number
}

type WindowFeatures struct {
	objc.Object
}

func MakeWindowFeatures(ptr unsafe.Pointer) WindowFeatures {
	return WindowFeatures{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WindowFeaturesClass) Alloc() WindowFeatures {
	rv := objc.CallMethod[WindowFeatures](wc, objc.GetSelector("alloc"))
	return rv
}

func WindowFeatures_Alloc() WindowFeatures {
	return WindowFeaturesClass.Alloc()
}

func (wc _WindowFeaturesClass) New() WindowFeatures {
	rv := objc.CallMethod[WindowFeatures](wc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewWindowFeatures() WindowFeatures {
	return WindowFeaturesClass.New()
}

func WindowFeatures_New() WindowFeatures {
	return WindowFeaturesClass.New()
}

func (w_ WindowFeatures) Init() WindowFeatures {
	rv := objc.CallMethod[WindowFeatures](w_, objc.GetSelector("init"))
	return rv
}

func WindowFeatures_Init() WindowFeatures {
	return WindowFeaturesClass.Alloc().Init()
}

func (w_ WindowFeatures) AllowsResizing() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.GetSelector("allowsResizing"))
	return rv
}

func (w_ WindowFeatures) Height() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.GetSelector("height"))
	return rv
}

func (w_ WindowFeatures) Width() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.GetSelector("width"))
	return rv
}

func (w_ WindowFeatures) X() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.GetSelector("x"))
	return rv
}

func (w_ WindowFeatures) Y() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.GetSelector("y"))
	return rv
}

func (w_ WindowFeatures) MenuBarVisibility() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.GetSelector("menuBarVisibility"))
	return rv
}

func (w_ WindowFeatures) StatusBarVisibility() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.GetSelector("statusBarVisibility"))
	return rv
}

func (w_ WindowFeatures) ToolbarsVisibility() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.GetSelector("toolbarsVisibility"))
	return rv
}
