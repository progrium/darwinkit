// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptionRegion] class.
var CaptionRegionClass = _CaptionRegionClass{objc.GetClass("AVCaptionRegion")}

type _CaptionRegionClass struct {
	objc.Class
}

// An interface definition for the [CaptionRegion] class.
type ICaptionRegion interface {
	objc.IObject
	Origin() CaptionPoint
	Scroll() CaptionRegionScroll
	WritingMode() CaptionRegionWritingMode
	Identifier() string
	Size() CaptionSize
	DisplayAlignment() CaptionRegionDisplayAlignment
}

// An object that represents the region in which the system presents a caption. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion?language=objc
type CaptionRegion struct {
	objc.Object
}

func CaptionRegionFrom(ptr unsafe.Pointer) CaptionRegion {
	return CaptionRegion{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptionRegionClass) Alloc() CaptionRegion {
	rv := objc.Call[CaptionRegion](cc, objc.Sel("alloc"))
	return rv
}

func CaptionRegion_Alloc() CaptionRegion {
	return CaptionRegionClass.Alloc()
}

func (cc _CaptionRegionClass) New() CaptionRegion {
	rv := objc.Call[CaptionRegion](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptionRegion() CaptionRegion {
	return CaptionRegionClass.New()
}

func (c_ CaptionRegion) Init() CaptionRegion {
	rv := objc.Call[CaptionRegion](c_, objc.Sel("init"))
	return rv
}

// The bottom caption region for SubRip Text (SRT) format captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3857629-subriptextbottomregion?language=objc
func (cc _CaptionRegionClass) SubRipTextBottomRegion() CaptionRegion {
	rv := objc.Call[CaptionRegion](cc, objc.Sel("subRipTextBottomRegion"))
	return rv
}

// The bottom caption region for SubRip Text (SRT) format captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3857629-subriptextbottomregion?language=objc
func CaptionRegion_SubRipTextBottomRegion() CaptionRegion {
	return CaptionRegionClass.SubRipTextBottomRegion()
}

// The right region for iTT format captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3857625-appleittrightregion?language=objc
func (cc _CaptionRegionClass) AppleITTRightRegion() CaptionRegion {
	rv := objc.Call[CaptionRegion](cc, objc.Sel("appleITTRightRegion"))
	return rv
}

// The right region for iTT format captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3857625-appleittrightregion?language=objc
func CaptionRegion_AppleITTRightRegion() CaptionRegion {
	return CaptionRegionClass.AppleITTRightRegion()
}

// The region’s top-left position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3857627-origin?language=objc
func (c_ CaptionRegion) Origin() CaptionPoint {
	rv := objc.Call[CaptionPoint](c_, objc.Sel("origin"))
	return rv
}

// The scroll mode of the region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3752855-scroll?language=objc
func (c_ CaptionRegion) Scroll() CaptionRegionScroll {
	rv := objc.Call[CaptionRegionScroll](c_, objc.Sel("scroll"))
	return rv
}

// The top region for iTT format captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3857626-appleitttopregion?language=objc
func (cc _CaptionRegionClass) AppleITTTopRegion() CaptionRegion {
	rv := objc.Call[CaptionRegion](cc, objc.Sel("appleITTTopRegion"))
	return rv
}

// The top region for iTT format captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3857626-appleitttopregion?language=objc
func CaptionRegion_AppleITTTopRegion() CaptionRegion {
	return CaptionRegionClass.AppleITTTopRegion()
}

// The block and inline progression direction of the region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3752857-writingmode?language=objc
func (c_ CaptionRegion) WritingMode() CaptionRegionWritingMode {
	rv := objc.Call[CaptionRegionWritingMode](c_, objc.Sel("writingMode"))
	return rv
}

// The left region for iTT format captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3857624-appleittleftregion?language=objc
func (cc _CaptionRegionClass) AppleITTLeftRegion() CaptionRegion {
	rv := objc.Call[CaptionRegion](cc, objc.Sel("appleITTLeftRegion"))
	return rv
}

// The left region for iTT format captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3857624-appleittleftregion?language=objc
func CaptionRegion_AppleITTLeftRegion() CaptionRegion {
	return CaptionRegionClass.AppleITTLeftRegion()
}

// The bottom region for iTT format captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3857623-appleittbottomregion?language=objc
func (cc _CaptionRegionClass) AppleITTBottomRegion() CaptionRegion {
	rv := objc.Call[CaptionRegion](cc, objc.Sel("appleITTBottomRegion"))
	return rv
}

// The bottom region for iTT format captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3857623-appleittbottomregion?language=objc
func CaptionRegion_AppleITTBottomRegion() CaptionRegion {
	return CaptionRegionClass.AppleITTBottomRegion()
}

// A string that identifies the region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3752853-identifier?language=objc
func (c_ CaptionRegion) Identifier() string {
	rv := objc.Call[string](c_, objc.Sel("identifier"))
	return rv
}

// The height and width of the region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3857628-size?language=objc
func (c_ CaptionRegion) Size() CaptionSize {
	rv := objc.Call[CaptionSize](c_, objc.Sel("size"))
	return rv
}

// The alignment of lines for the region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregion/3752850-displayalignment?language=objc
func (c_ CaptionRegion) DisplayAlignment() CaptionRegionDisplayAlignment {
	rv := objc.Call[CaptionRegionDisplayAlignment](c_, objc.Sel("displayAlignment"))
	return rv
}
