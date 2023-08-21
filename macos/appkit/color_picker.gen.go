// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ColorPicker] class.
var ColorPickerClass = _ColorPickerClass{objc.GetClass("NSColorPicker")}

type _ColorPickerClass struct {
	objc.Class
}

// An interface definition for the [ColorPicker] class.
type IColorPicker interface {
	objc.IObject
	AttachColorList(colorList IColorList)
	InsertNewButtonImageIn(newButtonImage IImage, buttonCell IButtonCell)
	SetMode(mode ColorPanelMode)
	DetachColorList(colorList IColorList)
	ViewSizeChanged(sender objc.IObject)
	ColorPanel() ColorPanel
	ProvideNewButtonImage() Image
	ButtonToolTip() string
	MinContentSize() foundation.Size
}

// An abstract superclass that implements the default color picking protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker?language=objc
type ColorPicker struct {
	objc.Object
}

func ColorPickerFrom(ptr unsafe.Pointer) ColorPicker {
	return ColorPicker{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ ColorPicker) InitWithPickerMaskColorPanel(mask uint, owningColorPanel IColorPanel) ColorPicker {
	rv := objc.Call[ColorPicker](c_, objc.Sel("initWithPickerMask:colorPanel:"), mask, objc.Ptr(owningColorPanel))
	return rv
}

// Initializes the color picker with the specified color panel and color picker mode mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/1492397-initwithpickermask?language=objc
func NewColorPickerWithPickerMaskColorPanel(mask uint, owningColorPanel IColorPanel) ColorPicker {
	instance := ColorPickerClass.Alloc().InitWithPickerMaskColorPanel(mask, owningColorPanel)
	instance.Autorelease()
	return instance
}

func (cc _ColorPickerClass) Alloc() ColorPicker {
	rv := objc.Call[ColorPicker](cc, objc.Sel("alloc"))
	return rv
}

func ColorPicker_Alloc() ColorPicker {
	return ColorPickerClass.Alloc()
}

func (cc _ColorPickerClass) New() ColorPicker {
	rv := objc.Call[ColorPicker](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewColorPicker() ColorPicker {
	return ColorPickerClass.New()
}

func (c_ ColorPicker) Init() ColorPicker {
	rv := objc.Call[ColorPicker](c_, objc.Sel("init"))
	return rv
}

// Overriden to attach a color list to a color picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/1492395-attachcolorlist?language=objc
func (c_ ColorPicker) AttachColorList(colorList IColorList) {
	objc.Call[objc.Void](c_, objc.Sel("attachColorList:"), objc.Ptr(colorList))
}

// Sets the image used for the specified button cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/1492402-insertnewbuttonimage?language=objc
func (c_ ColorPicker) InsertNewButtonImageIn(newButtonImage IImage, buttonCell IButtonCell) {
	objc.Call[objc.Void](c_, objc.Sel("insertNewButtonImage:in:"), objc.Ptr(newButtonImage), objc.Ptr(buttonCell))
}

// Overriden to set the color picker’s mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/1492394-setmode?language=objc
func (c_ ColorPicker) SetMode(mode ColorPanelMode) {
	objc.Call[objc.Void](c_, objc.Sel("setMode:"), mode)
}

// Overriden to detach a color list from a color picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/1492388-detachcolorlist?language=objc
func (c_ ColorPicker) DetachColorList(colorList IColorList) {
	objc.Call[objc.Void](c_, objc.Sel("detachColorList:"), objc.Ptr(colorList))
}

// Overriden to respond to a size change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/1492389-viewsizechanged?language=objc
func (c_ ColorPicker) ViewSizeChanged(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("viewSizeChanged:"), sender)
}

// The color panel instance that owns the color picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/1492396-colorpanel?language=objc
func (c_ ColorPicker) ColorPanel() ColorPanel {
	rv := objc.Call[ColorPanel](c_, objc.Sel("colorPanel"))
	return rv
}

// The button image used by the color picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/1492393-providenewbuttonimage?language=objc
func (c_ ColorPicker) ProvideNewButtonImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("provideNewButtonImage"))
	return rv
}

// The tool tip that is shown when the mouse cursor is over the color picker’s button image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/1492404-buttontooltip?language=objc
func (c_ ColorPicker) ButtonToolTip() string {
	rv := objc.Call[string](c_, objc.Sel("buttonToolTip"))
	return rv
}

// The minimum content size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/1492391-mincontentsize?language=objc
func (c_ ColorPicker) MinContentSize() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("minContentSize"))
	return rv
}
