// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MagnificationGestureRecognizer] class.
var MagnificationGestureRecognizerClass = _MagnificationGestureRecognizerClass{objc.GetClass("NSMagnificationGestureRecognizer")}

type _MagnificationGestureRecognizerClass struct {
	objc.Class
}

// An interface definition for the [MagnificationGestureRecognizer] class.
type IMagnificationGestureRecognizer interface {
	IGestureRecognizer
	Magnification() float64
	SetMagnification(value float64)
}

// A continuous gesture recognizer that tracks a pinch gesture that magnifies content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmagnificationgesturerecognizer?language=objc
type MagnificationGestureRecognizer struct {
	GestureRecognizer
}

func MagnificationGestureRecognizerFrom(ptr unsafe.Pointer) MagnificationGestureRecognizer {
	return MagnificationGestureRecognizer{
		GestureRecognizer: GestureRecognizerFrom(ptr),
	}
}

func (mc _MagnificationGestureRecognizerClass) Alloc() MagnificationGestureRecognizer {
	rv := objc.Call[MagnificationGestureRecognizer](mc, objc.Sel("alloc"))
	return rv
}

func MagnificationGestureRecognizer_Alloc() MagnificationGestureRecognizer {
	return MagnificationGestureRecognizerClass.Alloc()
}

func (mc _MagnificationGestureRecognizerClass) New() MagnificationGestureRecognizer {
	rv := objc.Call[MagnificationGestureRecognizer](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMagnificationGestureRecognizer() MagnificationGestureRecognizer {
	return MagnificationGestureRecognizerClass.New()
}

func (m_ MagnificationGestureRecognizer) Init() MagnificationGestureRecognizer {
	rv := objc.Call[MagnificationGestureRecognizer](m_, objc.Sel("init"))
	return rv
}

func (m_ MagnificationGestureRecognizer) InitWithTargetAction(target objc.IObject, action objc.Selector) MagnificationGestureRecognizer {
	rv := objc.Call[MagnificationGestureRecognizer](m_, objc.Sel("initWithTarget:action:"), target, action)
	return rv
}

// Initializes the gesture recognizer with the specified target and action information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1535012-initwithtarget?language=objc
func NewMagnificationGestureRecognizerWithTargetAction(target objc.IObject, action objc.Selector) MagnificationGestureRecognizer {
	instance := MagnificationGestureRecognizerClass.Alloc().InitWithTargetAction(target, action)
	instance.Autorelease()
	return instance
}

// The amount of magnification to apply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmagnificationgesturerecognizer/1510428-magnification?language=objc
func (m_ MagnificationGestureRecognizer) Magnification() float64 {
	rv := objc.Call[float64](m_, objc.Sel("magnification"))
	return rv
}

// The amount of magnification to apply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmagnificationgesturerecognizer/1510428-magnification?language=objc
func (m_ MagnificationGestureRecognizer) SetMagnification(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setMagnification:"), value)
}
