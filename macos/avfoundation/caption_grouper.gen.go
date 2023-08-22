// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptionGrouper] class.
var CaptionGrouperClass = _CaptionGrouperClass{objc.GetClass("AVCaptionGrouper")}

type _CaptionGrouperClass struct {
	objc.Class
}

// An interface definition for the [CaptionGrouper] class.
type ICaptionGrouper interface {
	objc.IObject
	AddCaption(input ICaption)
	FlushAddedCaptionsIntoGroupsUpToTime(upToTime coremedia.Time) []CaptionGroup
}

// An object that analyzes the temporal overlaps of caption objects to create caption groups for each span of concurrent captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptiongrouper?language=objc
type CaptionGrouper struct {
	objc.Object
}

func CaptionGrouperFrom(ptr unsafe.Pointer) CaptionGrouper {
	return CaptionGrouper{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptionGrouperClass) Alloc() CaptionGrouper {
	rv := objc.Call[CaptionGrouper](cc, objc.Sel("alloc"))
	return rv
}

func CaptionGrouper_Alloc() CaptionGrouper {
	return CaptionGrouperClass.Alloc()
}

func (cc _CaptionGrouperClass) New() CaptionGrouper {
	rv := objc.Call[CaptionGrouper](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptionGrouper() CaptionGrouper {
	return CaptionGrouperClass.New()
}

func (c_ CaptionGrouper) Init() CaptionGrouper {
	rv := objc.Call[CaptionGrouper](c_, objc.Sel("init"))
	return rv
}

// Adds a caption to the pending group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptiongrouper/3752965-addcaption?language=objc
func (c_ CaptionGrouper) AddCaption(input ICaption) {
	objc.Call[objc.Void](c_, objc.Sel("addCaption:"), objc.Ptr(input))
}

// Creates caption groups for the captions you enqueue up to the time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptiongrouper/3752966-flushaddedcaptionsintogroupsupto?language=objc
func (c_ CaptionGrouper) FlushAddedCaptionsIntoGroupsUpToTime(upToTime coremedia.Time) []CaptionGroup {
	rv := objc.Call[[]CaptionGroup](c_, objc.Sel("flushAddedCaptionsIntoGroupsUpToTime:"), upToTime)
	return rv
}
