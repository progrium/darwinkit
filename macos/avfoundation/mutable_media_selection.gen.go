// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableMediaSelection] class.
var MutableMediaSelectionClass = _MutableMediaSelectionClass{objc.GetClass("AVMutableMediaSelection")}

type _MutableMediaSelectionClass struct {
	objc.Class
}

// An interface definition for the [MutableMediaSelection] class.
type IMutableMediaSelection interface {
	IMediaSelection
	SelectMediaOptionInMediaSelectionGroup(mediaSelectionOption IMediaSelectionOption, mediaSelectionGroup IMediaSelectionGroup)
}

// A mutable object that represents a complete rendition of media selection options on an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemediaselection?language=objc
type MutableMediaSelection struct {
	MediaSelection
}

func MutableMediaSelectionFrom(ptr unsafe.Pointer) MutableMediaSelection {
	return MutableMediaSelection{
		MediaSelection: MediaSelectionFrom(ptr),
	}
}

func (mc _MutableMediaSelectionClass) Alloc() MutableMediaSelection {
	rv := objc.Call[MutableMediaSelection](mc, objc.Sel("alloc"))
	return rv
}

func MutableMediaSelection_Alloc() MutableMediaSelection {
	return MutableMediaSelectionClass.Alloc()
}

func (mc _MutableMediaSelectionClass) New() MutableMediaSelection {
	rv := objc.Call[MutableMediaSelection](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableMediaSelection() MutableMediaSelection {
	return MutableMediaSelectionClass.New()
}

func (m_ MutableMediaSelection) Init() MutableMediaSelection {
	rv := objc.Call[MutableMediaSelection](m_, objc.Sel("init"))
	return rv
}

// Selects the media option in the specified media selection group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemediaselection/1386768-selectmediaoption?language=objc
func (m_ MutableMediaSelection) SelectMediaOptionInMediaSelectionGroup(mediaSelectionOption IMediaSelectionOption, mediaSelectionGroup IMediaSelectionGroup) {
	objc.Call[objc.Void](m_, objc.Sel("selectMediaOption:inMediaSelectionGroup:"), objc.Ptr(mediaSelectionOption), objc.Ptr(mediaSelectionGroup))
}
