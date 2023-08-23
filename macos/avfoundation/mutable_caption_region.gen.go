// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableCaptionRegion] class.
var MutableCaptionRegionClass = _MutableCaptionRegionClass{objc.GetClass("AVMutableCaptionRegion")}

type _MutableCaptionRegionClass struct {
	objc.Class
}

// An interface definition for the [MutableCaptionRegion] class.
type IMutableCaptionRegion interface {
	ICaptionRegion
	SetOrigin(value CaptionPoint)
	SetScroll(value CaptionRegionScroll)
	SetWritingMode(value CaptionRegionWritingMode)
	SetSize(value CaptionSize)
	SetDisplayAlignment(value CaptionRegionDisplayAlignment)
}

// A mutable caption region subclass that you use to create new caption regions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaptionregion?language=objc
type MutableCaptionRegion struct {
	CaptionRegion
}

func MutableCaptionRegionFrom(ptr unsafe.Pointer) MutableCaptionRegion {
	return MutableCaptionRegion{
		CaptionRegion: CaptionRegionFrom(ptr),
	}
}

func (m_ MutableCaptionRegion) InitWithIdentifier(identifier string) MutableCaptionRegion {
	rv := objc.Call[MutableCaptionRegion](m_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Creates a caption region that has an identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaptionregion/3752923-initwithidentifier?language=objc
func NewMutableCaptionRegionWithIdentifier(identifier string) MutableCaptionRegion {
	instance := MutableCaptionRegionClass.Alloc().InitWithIdentifier(identifier)
	instance.Autorelease()
	return instance
}

func (m_ MutableCaptionRegion) Init() MutableCaptionRegion {
	rv := objc.Call[MutableCaptionRegion](m_, objc.Sel("init"))
	return rv
}

func (mc _MutableCaptionRegionClass) Alloc() MutableCaptionRegion {
	rv := objc.Call[MutableCaptionRegion](mc, objc.Sel("alloc"))
	return rv
}

func MutableCaptionRegion_Alloc() MutableCaptionRegion {
	return MutableCaptionRegionClass.Alloc()
}

func (mc _MutableCaptionRegionClass) New() MutableCaptionRegion {
	rv := objc.Call[MutableCaptionRegion](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableCaptionRegion() MutableCaptionRegion {
	return MutableCaptionRegionClass.New()
}

// The region’s top-left position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaptionregion/3857638-origin?language=objc
func (m_ MutableCaptionRegion) SetOrigin(value CaptionPoint) {
	objc.Call[objc.Void](m_, objc.Sel("setOrigin:"), value)
}

// The scroll mode of the region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaptionregion/3752925-scroll?language=objc
func (m_ MutableCaptionRegion) SetScroll(value CaptionRegionScroll) {
	objc.Call[objc.Void](m_, objc.Sel("setScroll:"), value)
}

// The block and inline progression direction of the region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaptionregion/3752926-writingmode?language=objc
func (m_ MutableCaptionRegion) SetWritingMode(value CaptionRegionWritingMode) {
	objc.Call[objc.Void](m_, objc.Sel("setWritingMode:"), value)
}

// The height and width of the region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaptionregion/3857639-size?language=objc
func (m_ MutableCaptionRegion) SetSize(value CaptionSize) {
	objc.Call[objc.Void](m_, objc.Sel("setSize:"), value)
}

// The alignment of lines for the region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaptionregion/3752919-displayalignment?language=objc
func (m_ MutableCaptionRegion) SetDisplayAlignment(value CaptionRegionDisplayAlignment) {
	objc.Call[objc.Void](m_, objc.Sel("setDisplayAlignment:"), value)
}
