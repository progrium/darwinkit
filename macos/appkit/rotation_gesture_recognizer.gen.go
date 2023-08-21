// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RotationGestureRecognizer] class.
var RotationGestureRecognizerClass = _RotationGestureRecognizerClass{objc.GetClass("NSRotationGestureRecognizer")}

type _RotationGestureRecognizerClass struct {
	objc.Class
}

// An interface definition for the [RotationGestureRecognizer] class.
type IRotationGestureRecognizer interface {
	IGestureRecognizer
	Rotation() float64
	SetRotation(value float64)
	RotationInDegrees() float64
	SetRotationInDegrees(value float64)
}

// A continuous gesture recognizer that tracks two trackpad touches moving opposite each other in a circular motion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrotationgesturerecognizer?language=objc
type RotationGestureRecognizer struct {
	GestureRecognizer
}

func RotationGestureRecognizerFrom(ptr unsafe.Pointer) RotationGestureRecognizer {
	return RotationGestureRecognizer{
		GestureRecognizer: GestureRecognizerFrom(ptr),
	}
}

func (rc _RotationGestureRecognizerClass) Alloc() RotationGestureRecognizer {
	rv := objc.Call[RotationGestureRecognizer](rc, objc.Sel("alloc"))
	return rv
}

func RotationGestureRecognizer_Alloc() RotationGestureRecognizer {
	return RotationGestureRecognizerClass.Alloc()
}

func (rc _RotationGestureRecognizerClass) New() RotationGestureRecognizer {
	rv := objc.Call[RotationGestureRecognizer](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRotationGestureRecognizer() RotationGestureRecognizer {
	return RotationGestureRecognizerClass.New()
}

func (r_ RotationGestureRecognizer) Init() RotationGestureRecognizer {
	rv := objc.Call[RotationGestureRecognizer](r_, objc.Sel("init"))
	return rv
}

func (r_ RotationGestureRecognizer) InitWithTargetAction(target objc.IObject, action objc.Selector) RotationGestureRecognizer {
	rv := objc.Call[RotationGestureRecognizer](r_, objc.Sel("initWithTarget:action:"), target, action)
	return rv
}

// Initializes the gesture recognizer with the specified target and action information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1535012-initwithtarget?language=objc
func NewRotationGestureRecognizerWithTargetAction(target objc.IObject, action objc.Selector) RotationGestureRecognizer {
	instance := RotationGestureRecognizerClass.Alloc().InitWithTargetAction(target, action)
	instance.Autorelease()
	return instance
}

// The rotation of the gesture in radians. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrotationgesturerecognizer/1527087-rotation?language=objc
func (r_ RotationGestureRecognizer) Rotation() float64 {
	rv := objc.Call[float64](r_, objc.Sel("rotation"))
	return rv
}

// The rotation of the gesture in radians. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrotationgesturerecognizer/1527087-rotation?language=objc
func (r_ RotationGestureRecognizer) SetRotation(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setRotation:"), value)
}

// The rotation of the gesture in degrees. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrotationgesturerecognizer/1535523-rotationindegrees?language=objc
func (r_ RotationGestureRecognizer) RotationInDegrees() float64 {
	rv := objc.Call[float64](r_, objc.Sel("rotationInDegrees"))
	return rv
}

// The rotation of the gesture in degrees. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrotationgesturerecognizer/1535523-rotationindegrees?language=objc
func (r_ RotationGestureRecognizer) SetRotationInDegrees(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setRotationInDegrees:"), value)
}
