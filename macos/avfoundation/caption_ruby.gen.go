// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptionRuby] class.
var CaptionRubyClass = _CaptionRubyClass{objc.GetClass("AVCaptionRuby")}

type _CaptionRubyClass struct {
	objc.Class
}

// An interface definition for the [CaptionRuby] class.
type ICaptionRuby interface {
	objc.IObject
	Alignment() CaptionRubyAlignment
	Position() CaptionRubyPosition
	Text() string
}

// An object that presents ruby characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionruby?language=objc
type CaptionRuby struct {
	objc.Object
}

func CaptionRubyFrom(ptr unsafe.Pointer) CaptionRuby {
	return CaptionRuby{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ CaptionRuby) InitWithText(text string) CaptionRuby {
	rv := objc.Call[CaptionRuby](c_, objc.Sel("initWithText:"), text)
	return rv
}

// Creates ruby text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionruby/3752870-initwithtext?language=objc
func NewCaptionRubyWithText(text string) CaptionRuby {
	instance := CaptionRubyClass.Alloc().InitWithText(text)
	instance.Autorelease()
	return instance
}

func (cc _CaptionRubyClass) Alloc() CaptionRuby {
	rv := objc.Call[CaptionRuby](cc, objc.Sel("alloc"))
	return rv
}

func CaptionRuby_Alloc() CaptionRuby {
	return CaptionRubyClass.Alloc()
}

func (cc _CaptionRubyClass) New() CaptionRuby {
	rv := objc.Call[CaptionRuby](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptionRuby() CaptionRuby {
	return CaptionRubyClass.New()
}

func (c_ CaptionRuby) Init() CaptionRuby {
	rv := objc.Call[CaptionRuby](c_, objc.Sel("init"))
	return rv
}

// The ruby text alignment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionruby/3752869-alignment?language=objc
func (c_ CaptionRuby) Alignment() CaptionRubyAlignment {
	rv := objc.Call[CaptionRubyAlignment](c_, objc.Sel("alignment"))
	return rv
}

// The ruby text position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionruby/3752872-position?language=objc
func (c_ CaptionRuby) Position() CaptionRubyPosition {
	rv := objc.Call[CaptionRubyPosition](c_, objc.Sel("position"))
	return rv
}

// The ruby text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionruby/3752873-text?language=objc
func (c_ CaptionRuby) Text() string {
	rv := objc.Call[string](c_, objc.Sel("text"))
	return rv
}
