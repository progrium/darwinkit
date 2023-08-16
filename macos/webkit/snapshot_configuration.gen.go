// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SnapshotConfiguration] class.
var SnapshotConfigurationClass = _SnapshotConfigurationClass{objc.GetClass("WKSnapshotConfiguration")}

type _SnapshotConfigurationClass struct {
	objc.Class
}

// An interface definition for the [SnapshotConfiguration] class.
type ISnapshotConfiguration interface {
	objc.IObject
	AfterScreenUpdates() bool
	SetAfterScreenUpdates(value bool)
	Rect() coregraphics.Rect
	SetRect(value coregraphics.Rect)
	SnapshotWidth() foundation.Number
	SetSnapshotWidth(value foundation.INumber)
}

// The configuration data to use when generating an image from a web view’s contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wksnapshotconfiguration?language=objc
type SnapshotConfiguration struct {
	objc.Object
}

func SnapshotConfigurationFrom(ptr unsafe.Pointer) SnapshotConfiguration {
	return SnapshotConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SnapshotConfigurationClass) Alloc() SnapshotConfiguration {
	rv := objc.Call[SnapshotConfiguration](sc, objc.Sel("alloc"))
	return rv
}

func SnapshotConfiguration_Alloc() SnapshotConfiguration {
	return SnapshotConfigurationClass.Alloc()
}

func (sc _SnapshotConfigurationClass) New() SnapshotConfiguration {
	rv := objc.Call[SnapshotConfiguration](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSnapshotConfiguration() SnapshotConfiguration {
	return SnapshotConfigurationClass.New()
}

func (s_ SnapshotConfiguration) Init() SnapshotConfiguration {
	rv := objc.Call[SnapshotConfiguration](s_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether to take the snapshot after incorporating any pending screen updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wksnapshotconfiguration/3172739-afterscreenupdates?language=objc
func (s_ SnapshotConfiguration) AfterScreenUpdates() bool {
	rv := objc.Call[bool](s_, objc.Sel("afterScreenUpdates"))
	return rv
}

// A Boolean value that indicates whether to take the snapshot after incorporating any pending screen updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wksnapshotconfiguration/3172739-afterscreenupdates?language=objc
func (s_ SnapshotConfiguration) SetAfterScreenUpdates(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setAfterScreenUpdates:"), value)
}

// The portion of your web view to capture, specified as a rectangle in the view’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wksnapshotconfiguration/2873250-rect?language=objc
func (s_ SnapshotConfiguration) Rect() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](s_, objc.Sel("rect"))
	return rv
}

// The portion of your web view to capture, specified as a rectangle in the view’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wksnapshotconfiguration/2873250-rect?language=objc
func (s_ SnapshotConfiguration) SetRect(value coregraphics.Rect) {
	objc.Call[objc.Void](s_, objc.Sel("setRect:"), value)
}

// The width of the captured image, in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wksnapshotconfiguration/2873249-snapshotwidth?language=objc
func (s_ SnapshotConfiguration) SnapshotWidth() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("snapshotWidth"))
	return rv
}

// The width of the captured image, in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wksnapshotconfiguration/2873249-snapshotwidth?language=objc
func (s_ SnapshotConfiguration) SetSnapshotWidth(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setSnapshotWidth:"), objc.Ptr(value))
}
