// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that provides basic behavior for a color picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickingdefault?language=objc
type PColorPickingDefault interface {
	// optional
	AlphaControlAddedOrRemoved(sender objc.Object)
	HasAlphaControlAddedOrRemoved() bool

	// optional
	InitWithPickerMaskColorPanel(mask uint, owningColorPanel ColorPanel) objc.IObject
	HasInitWithPickerMaskColorPanel() bool

	// optional
	AttachColorList(colorList ColorList)
	HasAttachColorList() bool

	// optional
	InsertNewButtonImageIn(newButtonImage Image, buttonCell ButtonCell)
	HasInsertNewButtonImageIn() bool

	// optional
	SetMode(mode ColorPanelMode)
	HasSetMode() bool

	// optional
	DetachColorList(colorList ColorList)
	HasDetachColorList() bool

	// optional
	ProvideNewButtonImage() IImage
	HasProvideNewButtonImage() bool

	// optional
	ButtonToolTip() string
	HasButtonToolTip() bool

	// optional
	MinContentSize() foundation.Size
	HasMinContentSize() bool

	// optional
	ViewSizeChanged(sender objc.Object)
	HasViewSizeChanged() bool
}

// A concrete type wrapper for the [PColorPickingDefault] protocol.
type ColorPickingDefaultWrapper struct {
	objc.Object
}

func (c_ ColorPickingDefaultWrapper) HasAlphaControlAddedOrRemoved() bool {
	return c_.RespondsToSelector(objc.Sel("alphaControlAddedOrRemoved:"))
}

// Sent when the color panel's opacity controls have been hidden or displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickingdefault/1535478-alphacontroladdedorremoved?language=objc
func (c_ ColorPickingDefaultWrapper) AlphaControlAddedOrRemoved(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("alphaControlAddedOrRemoved:"), sender)
}

func (c_ ColorPickingDefaultWrapper) HasInitWithPickerMaskColorPanel() bool {
	return c_.RespondsToSelector(objc.Sel("initWithPickerMask:colorPanel:"))
}

// Initializes the receiver with a given color panel and its mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickingdefault/1528432-initwithpickermask?language=objc
func (c_ ColorPickingDefaultWrapper) InitWithPickerMaskColorPanel(mask uint, owningColorPanel IColorPanel) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("initWithPickerMask:colorPanel:"), mask, objc.Ptr(owningColorPanel))
	return rv
}

func (c_ ColorPickingDefaultWrapper) HasAttachColorList() bool {
	return c_.RespondsToSelector(objc.Sel("attachColorList:"))
}

// Tells the receiver to attach the given color list, if it isn’t already displaying the list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickingdefault/1531650-attachcolorlist?language=objc
func (c_ ColorPickingDefaultWrapper) AttachColorList(colorList IColorList) {
	objc.Call[objc.Void](c_, objc.Sel("attachColorList:"), objc.Ptr(colorList))
}

func (c_ ColorPickingDefaultWrapper) HasInsertNewButtonImageIn() bool {
	return c_.RespondsToSelector(objc.Sel("insertNewButtonImage:in:"))
}

// Sets the image of a given button cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickingdefault/1525078-insertnewbuttonimage?language=objc
func (c_ ColorPickingDefaultWrapper) InsertNewButtonImageIn(newButtonImage IImage, buttonCell IButtonCell) {
	objc.Call[objc.Void](c_, objc.Sel("insertNewButtonImage:in:"), objc.Ptr(newButtonImage), objc.Ptr(buttonCell))
}

func (c_ ColorPickingDefaultWrapper) HasSetMode() bool {
	return c_.RespondsToSelector(objc.Sel("setMode:"))
}

// Specifies the receiver’s mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickingdefault/1525088-setmode?language=objc
func (c_ ColorPickingDefaultWrapper) SetMode(mode ColorPanelMode) {
	objc.Call[objc.Void](c_, objc.Sel("setMode:"), mode)
}

func (c_ ColorPickingDefaultWrapper) HasDetachColorList() bool {
	return c_.RespondsToSelector(objc.Sel("detachColorList:"))
}

// Tells the receiver to detach the given color list, unless the receiver isn’t displaying the list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickingdefault/1532761-detachcolorlist?language=objc
func (c_ ColorPickingDefaultWrapper) DetachColorList(colorList IColorList) {
	objc.Call[objc.Void](c_, objc.Sel("detachColorList:"), objc.Ptr(colorList))
}

func (c_ ColorPickingDefaultWrapper) HasProvideNewButtonImage() bool {
	return c_.RespondsToSelector(objc.Sel("provideNewButtonImage"))
}

// Provides the image of the button used to select the receiver in the color panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickingdefault/1534615-providenewbuttonimage?language=objc
func (c_ ColorPickingDefaultWrapper) ProvideNewButtonImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("provideNewButtonImage"))
	return rv
}

func (c_ ColorPickingDefaultWrapper) HasButtonToolTip() bool {
	return c_.RespondsToSelector(objc.Sel("buttonToolTip"))
}

// Provides the toolbar button help tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickingdefault/1535160-buttontooltip?language=objc
func (c_ ColorPickingDefaultWrapper) ButtonToolTip() string {
	rv := objc.Call[string](c_, objc.Sel("buttonToolTip"))
	return rv
}

func (c_ ColorPickingDefaultWrapper) HasMinContentSize() bool {
	return c_.RespondsToSelector(objc.Sel("minContentSize"))
}

// Indicates the receiver’s minimum content size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickingdefault/1535437-mincontentsize?language=objc
func (c_ ColorPickingDefaultWrapper) MinContentSize() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("minContentSize"))
	return rv
}

func (c_ ColorPickingDefaultWrapper) HasViewSizeChanged() bool {
	return c_.RespondsToSelector(objc.Sel("viewSizeChanged:"))
}

// Tells the recever when the color panel's view size changes in a way that might affect the color picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickingdefault/1535866-viewsizechanged?language=objc
func (c_ ColorPickingDefaultWrapper) ViewSizeChanged(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("viewSizeChanged:"), sender)
}
