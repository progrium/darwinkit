// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Slideshow] class.
var SlideshowClass = _SlideshowClass{objc.GetClass("IKSlideshow")}

type _SlideshowClass struct {
	objc.Class
}

// An interface definition for the [Slideshow] class.
type ISlideshow interface {
	objc.IObject
	IndexOfCurrentSlideshowItem() uint
	RunSlideshowWithDataSourceInModeOptions(dataSource PSlideshowDataSource, slideshowMode string, slideshowOptions foundation.Dictionary)
	RunSlideshowWithDataSourceObjectInModeOptions(dataSourceObject objc.IObject, slideshowMode string, slideshowOptions foundation.Dictionary)
	StopSlideshow(sender objc.IObject)
	ReloadSlideshowItemAtIndex(index uint)
	ReloadData()
	AutoPlayDelay() foundation.TimeInterval
	SetAutoPlayDelay(value foundation.TimeInterval)
}

// The IKSlideshow class encapsulates a data source and options for a slideshow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow?language=objc
type Slideshow struct {
	objc.Object
}

func SlideshowFrom(ptr unsafe.Pointer) Slideshow {
	return Slideshow{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SlideshowClass) Alloc() Slideshow {
	rv := objc.Call[Slideshow](sc, objc.Sel("alloc"))
	return rv
}

func Slideshow_Alloc() Slideshow {
	return SlideshowClass.Alloc()
}

func (sc _SlideshowClass) New() Slideshow {
	rv := objc.Call[Slideshow](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSlideshow() Slideshow {
	return SlideshowClass.New()
}

func (s_ Slideshow) Init() Slideshow {
	rv := objc.Call[Slideshow](s_, objc.Sel("init"))
	return rv
}

// Exports a slideshow item to the application that has the provided bundle identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/1503513-exportslideshowitem?language=objc
func (sc _SlideshowClass) ExportSlideshowItemToApplication(item objc.IObject, applicationBundleIdentifier string) {
	objc.Call[objc.Void](sc, objc.Sel("exportSlideshowItem:toApplication:"), item, applicationBundleIdentifier)
}

// Exports a slideshow item to the application that has the provided bundle identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/1503513-exportslideshowitem?language=objc
func Slideshow_ExportSlideshowItemToApplication(item objc.IObject, applicationBundleIdentifier string) {
	SlideshowClass.ExportSlideshowItemToApplication(item, applicationBundleIdentifier)
}

// Returns the index of the current slideshow item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/1503700-indexofcurrentslideshowitem?language=objc
func (s_ Slideshow) IndexOfCurrentSlideshowItem() uint {
	rv := objc.Call[uint](s_, objc.Sel("indexOfCurrentSlideshowItem"))
	return rv
}

// Runs a slideshow that contains the specified kind of items, provided from a data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/1504036-runslideshowwithdatasource?language=objc
func (s_ Slideshow) RunSlideshowWithDataSourceInModeOptions(dataSource PSlideshowDataSource, slideshowMode string, slideshowOptions foundation.Dictionary) {
	po0 := objc.WrapAsProtocol("IKSlideshowDataSource", dataSource)
	objc.Call[objc.Void](s_, objc.Sel("runSlideshowWithDataSource:inMode:options:"), po0, slideshowMode, slideshowOptions)
}

// Runs a slideshow that contains the specified kind of items, provided from a data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/1504036-runslideshowwithdatasource?language=objc
func (s_ Slideshow) RunSlideshowWithDataSourceObjectInModeOptions(dataSourceObject objc.IObject, slideshowMode string, slideshowOptions foundation.Dictionary) {
	objc.Call[objc.Void](s_, objc.Sel("runSlideshowWithDataSource:inMode:options:"), objc.Ptr(dataSourceObject), slideshowMode, slideshowOptions)
}

// Stops a slideshow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/1503801-stopslideshow?language=objc
func (s_ Slideshow) StopSlideshow(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("stopSlideshow:"), sender)
}

// Finds out whether the slideshow can export its contents to an application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/1504783-canexporttoapplication?language=objc
func (sc _SlideshowClass) CanExportToApplication(applicationBundleIdentifier string) bool {
	rv := objc.Call[bool](sc, objc.Sel("canExportToApplication:"), applicationBundleIdentifier)
	return rv
}

// Finds out whether the slideshow can export its contents to an application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/1504783-canexporttoapplication?language=objc
func Slideshow_CanExportToApplication(applicationBundleIdentifier string) bool {
	return SlideshowClass.CanExportToApplication(applicationBundleIdentifier)
}

// Reloads the data for a slideshow, starting at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/1504713-reloadslideshowitematindex?language=objc
func (s_ Slideshow) ReloadSlideshowItemAtIndex(index uint) {
	objc.Call[objc.Void](s_, objc.Sel("reloadSlideshowItemAtIndex:"), index)
}

// Returns a shared instance of a slideshow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/1504883-sharedslideshow?language=objc
func (sc _SlideshowClass) SharedSlideshow() Slideshow {
	rv := objc.Call[Slideshow](sc, objc.Sel("sharedSlideshow"))
	return rv
}

// Returns a shared instance of a slideshow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/1504883-sharedslideshow?language=objc
func Slideshow_SharedSlideshow() Slideshow {
	return SlideshowClass.SharedSlideshow()
}

// Reloads the data for a slideshow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/1504298-reloaddata?language=objc
func (s_ Slideshow) ReloadData() {
	objc.Call[objc.Void](s_, objc.Sel("reloadData"))
}

// Controls the interval of time before a slideshow starts to play automatically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/1504967-autoplaydelay?language=objc
func (s_ Slideshow) AutoPlayDelay() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](s_, objc.Sel("autoPlayDelay"))
	return rv
}

// Controls the interval of time before a slideshow starts to play automatically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/1504967-autoplaydelay?language=objc
func (s_ Slideshow) SetAutoPlayDelay(value foundation.TimeInterval) {
	objc.Call[objc.Void](s_, objc.Sel("setAutoPlayDelay:"), value)
}
