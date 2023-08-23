// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MediaSelectionGroup] class.
var MediaSelectionGroupClass = _MediaSelectionGroupClass{objc.GetClass("AVMediaSelectionGroup")}

type _MediaSelectionGroupClass struct {
	objc.Class
}

// An interface definition for the [MediaSelectionGroup] class.
type IMediaSelectionGroup interface {
	objc.IObject
	MediaSelectionOptionWithPropertyList(plist objc.IObject) MediaSelectionOption
	AllowsEmptySelection() bool
	Options() []MediaSelectionOption
	DefaultOption() MediaSelectionOption
}

// An object that represents a collection of mutually exclusive options for the presentation of media within an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectiongroup?language=objc
type MediaSelectionGroup struct {
	objc.Object
}

func MediaSelectionGroupFrom(ptr unsafe.Pointer) MediaSelectionGroup {
	return MediaSelectionGroup{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MediaSelectionGroupClass) Alloc() MediaSelectionGroup {
	rv := objc.Call[MediaSelectionGroup](mc, objc.Sel("alloc"))
	return rv
}

func MediaSelectionGroup_Alloc() MediaSelectionGroup {
	return MediaSelectionGroupClass.Alloc()
}

func (mc _MediaSelectionGroupClass) New() MediaSelectionGroup {
	rv := objc.Call[MediaSelectionGroup](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMediaSelectionGroup() MediaSelectionGroup {
	return MediaSelectionGroupClass.New()
}

func (m_ MediaSelectionGroup) Init() MediaSelectionGroup {
	rv := objc.Call[MediaSelectionGroup](m_, objc.Sel("init"))
	return rv
}

// Returns an array containing the media selection options from a given array that are playable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectiongroup/1387351-playablemediaselectionoptionsfro?language=objc
func (mc _MediaSelectionGroupClass) PlayableMediaSelectionOptionsFromArray(mediaSelectionOptions []IMediaSelectionOption) []MediaSelectionOption {
	rv := objc.Call[[]MediaSelectionOption](mc, objc.Sel("playableMediaSelectionOptionsFromArray:"), mediaSelectionOptions)
	return rv
}

// Returns an array containing the media selection options from a given array that are playable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectiongroup/1387351-playablemediaselectionoptionsfro?language=objc
func MediaSelectionGroup_PlayableMediaSelectionOptionsFromArray(mediaSelectionOptions []IMediaSelectionOption) []MediaSelectionOption {
	return MediaSelectionGroupClass.PlayableMediaSelectionOptionsFromArray(mediaSelectionOptions)
}

// Returns the media selection options that match the given property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectiongroup/1389968-mediaselectionoptionwithproperty?language=objc
func (m_ MediaSelectionGroup) MediaSelectionOptionWithPropertyList(plist objc.IObject) MediaSelectionOption {
	rv := objc.Call[MediaSelectionOption](m_, objc.Sel("mediaSelectionOptionWithPropertyList:"), plist)
	return rv
}

// Returns an array containing the media selection options from a given array that match the specified locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectiongroup/1387494-mediaselectionoptionsfromarray?language=objc
func (mc _MediaSelectionGroupClass) MediaSelectionOptionsFromArrayWithLocale(mediaSelectionOptions []IMediaSelectionOption, locale foundation.ILocale) []MediaSelectionOption {
	rv := objc.Call[[]MediaSelectionOption](mc, objc.Sel("mediaSelectionOptionsFromArray:withLocale:"), mediaSelectionOptions, objc.Ptr(locale))
	return rv
}

// Returns an array containing the media selection options from a given array that match the specified locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectiongroup/1387494-mediaselectionoptionsfromarray?language=objc
func MediaSelectionGroup_MediaSelectionOptionsFromArrayWithLocale(mediaSelectionOptions []IMediaSelectionOption, locale foundation.ILocale) []MediaSelectionOption {
	return MediaSelectionGroupClass.MediaSelectionOptionsFromArrayWithLocale(mediaSelectionOptions, locale)
}

// A Boolean value that indicates whether it’s possible to present none of the options in the group when an associated player item is played. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectiongroup/1388483-allowsemptyselection?language=objc
func (m_ MediaSelectionGroup) AllowsEmptySelection() bool {
	rv := objc.Call[bool](m_, objc.Sel("allowsEmptySelection"))
	return rv
}

// A collection of mutually exclusive media selection options [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectiongroup/1388351-options?language=objc
func (m_ MediaSelectionGroup) Options() []MediaSelectionOption {
	rv := objc.Call[[]MediaSelectionOption](m_, objc.Sel("options"))
	return rv
}

// The default option in the group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectiongroup/1388440-defaultoption?language=objc
func (m_ MediaSelectionGroup) DefaultOption() MediaSelectionOption {
	rv := objc.Call[MediaSelectionOption](m_, objc.Sel("defaultOption"))
	return rv
}
