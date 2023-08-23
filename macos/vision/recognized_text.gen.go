// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RecognizedText] class.
var RecognizedTextClass = _RecognizedTextClass{objc.GetClass("VNRecognizedText")}

type _RecognizedTextClass struct {
	objc.Class
}

// An interface definition for the [RecognizedText] class.
type IRecognizedText interface {
	objc.IObject
	BoundingBoxForRangeError(range_ foundation.Range, error foundation.IError) RectangleObservation
	String() string
	Confidence() Confidence
}

// Text recognized in an image through a text recognition request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedtext?language=objc
type RecognizedText struct {
	objc.Object
}

func RecognizedTextFrom(ptr unsafe.Pointer) RecognizedText {
	return RecognizedText{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RecognizedTextClass) Alloc() RecognizedText {
	rv := objc.Call[RecognizedText](rc, objc.Sel("alloc"))
	return rv
}

func RecognizedText_Alloc() RecognizedText {
	return RecognizedTextClass.Alloc()
}

func (rc _RecognizedTextClass) New() RecognizedText {
	rv := objc.Call[RecognizedText](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRecognizedText() RecognizedText {
	return RecognizedTextClass.New()
}

func (r_ RecognizedText) Init() RecognizedText {
	rv := objc.Call[RecognizedText](r_, objc.Sel("init"))
	return rv
}

// Calculates the bounding box around the characters in the range of a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedtext/3152634-boundingboxforrange?language=objc
func (r_ RecognizedText) BoundingBoxForRangeError(range_ foundation.Range, error foundation.IError) RectangleObservation {
	rv := objc.Call[RectangleObservation](r_, objc.Sel("boundingBoxForRange:error:"), range_, objc.Ptr(error))
	return rv
}

// The top candidate for recognized text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedtext/3152636-string?language=objc
func (r_ RecognizedText) String() string {
	rv := objc.Call[string](r_, objc.Sel("string"))
	return rv
}

// A normalized confidence score for the text recognition result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedtext/3152635-confidence?language=objc
func (r_ RecognizedText) Confidence() Confidence {
	rv := objc.Call[Confidence](r_, objc.Sel("confidence"))
	return rv
}
