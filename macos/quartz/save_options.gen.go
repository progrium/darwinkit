// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SaveOptions] class.
var SaveOptionsClass = _SaveOptionsClass{objc.GetClass("IKSaveOptions")}

type _SaveOptionsClass struct {
	objc.Class
}

// An interface definition for the [SaveOptions] class.
type ISaveOptions interface {
	objc.IObject
	AddSaveOptionsToView(view appkit.IView)
	AddSaveOptionsAccessoryViewToSavePanel(savePanel appkit.ISavePanel)
	ImageUTType() string
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
	ImageProperties() foundation.Dictionary
	UserSelection() foundation.Dictionary
	RememberLastSetting() bool
	SetRememberLastSetting(value bool)
}

// The IKSaveOptions class initializes, adds, and manages user interface options for saving image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions?language=objc
type SaveOptions struct {
	objc.Object
}

func SaveOptionsFrom(ptr unsafe.Pointer) SaveOptions {
	return SaveOptions{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SaveOptions) InitWithImagePropertiesImageUTType(imageProperties foundation.Dictionary, imageUTType string) SaveOptions {
	rv := objc.Call[SaveOptions](s_, objc.Sel("initWithImageProperties:imageUTType:"), imageProperties, imageUTType)
	return rv
}

// Initializes a save options accessory pane for the provided image properties and uniform type identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/1503412-initwithimageproperties?language=objc
func NewSaveOptionsWithImagePropertiesImageUTType(imageProperties foundation.Dictionary, imageUTType string) SaveOptions {
	instance := SaveOptionsClass.Alloc().InitWithImagePropertiesImageUTType(imageProperties, imageUTType)
	instance.Autorelease()
	return instance
}

func (sc _SaveOptionsClass) Alloc() SaveOptions {
	rv := objc.Call[SaveOptions](sc, objc.Sel("alloc"))
	return rv
}

func SaveOptions_Alloc() SaveOptions {
	return SaveOptionsClass.Alloc()
}

func (sc _SaveOptionsClass) New() SaveOptions {
	rv := objc.Call[SaveOptions](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSaveOptions() SaveOptions {
	return SaveOptionsClass.New()
}

func (s_ SaveOptions) Init() SaveOptions {
	rv := objc.Call[SaveOptions](s_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/1504945-addsaveoptionstoview?language=objc
func (s_ SaveOptions) AddSaveOptionsToView(view appkit.IView) {
	objc.Call[objc.Void](s_, objc.Sel("addSaveOptionsToView:"), objc.Ptr(view))
}

// Adds IKSaveOptions accessory view to a NSSavePanel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/1503458-addsaveoptionsaccessoryviewtosav?language=objc
func (s_ SaveOptions) AddSaveOptionsAccessoryViewToSavePanel(savePanel appkit.ISavePanel) {
	objc.Call[objc.Void](s_, objc.Sel("addSaveOptionsAccessoryViewToSavePanel:"), objc.Ptr(savePanel))
}

// Returns the uniform type identifier that reflects the user’s selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/1504388-imageuttype?language=objc
func (s_ SaveOptions) ImageUTType() string {
	rv := objc.Call[string](s_, objc.Sel("imageUTType"))
	return rv
}

// Specifies the delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/1503653-delegate?language=objc
func (s_ SaveOptions) Delegate() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("delegate"))
	return rv
}

// Specifies the delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/1503653-delegate?language=objc
func (s_ SaveOptions) SetDelegate(value objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), value)
}

// Returns a dictionary of updated image properties that reflects the user’s selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/1505299-imageproperties?language=objc
func (s_ SaveOptions) ImageProperties() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](s_, objc.Sel("imageProperties"))
	return rv
}

// Returns a dictionary that contains the save options selected by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/1504791-userselection?language=objc
func (s_ SaveOptions) UserSelection() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](s_, objc.Sel("userSelection"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/3738462-rememberlastsetting?language=objc
func (s_ SaveOptions) RememberLastSetting() bool {
	rv := objc.Call[bool](s_, objc.Sel("rememberLastSetting"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/3738462-rememberlastsetting?language=objc
func (s_ SaveOptions) SetRememberLastSetting(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setRememberLastSetting:"), value)
}
