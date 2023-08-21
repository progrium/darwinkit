// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PanGestureRecognizer] class.
var PanGestureRecognizerClass = _PanGestureRecognizerClass{objc.GetClass("NSPanGestureRecognizer")}

type _PanGestureRecognizerClass struct {
	objc.Class
}

// An interface definition for the [PanGestureRecognizer] class.
type IPanGestureRecognizer interface {
	IGestureRecognizer
	SetTranslationInView(translation foundation.Point, view IView)
	VelocityInView(view IView) foundation.Point
	TranslationInView(view IView) foundation.Point
	ButtonMask() uint
	SetButtonMask(value uint)
	NumberOfTouchesRequired() int
	SetNumberOfTouchesRequired(value int)
}

// A continuous gesture recognizer for panning gestures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspangesturerecognizer?language=objc
type PanGestureRecognizer struct {
	GestureRecognizer
}

func PanGestureRecognizerFrom(ptr unsafe.Pointer) PanGestureRecognizer {
	return PanGestureRecognizer{
		GestureRecognizer: GestureRecognizerFrom(ptr),
	}
}

func (pc _PanGestureRecognizerClass) Alloc() PanGestureRecognizer {
	rv := objc.Call[PanGestureRecognizer](pc, objc.Sel("alloc"))
	return rv
}

func PanGestureRecognizer_Alloc() PanGestureRecognizer {
	return PanGestureRecognizerClass.Alloc()
}

func (pc _PanGestureRecognizerClass) New() PanGestureRecognizer {
	rv := objc.Call[PanGestureRecognizer](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPanGestureRecognizer() PanGestureRecognizer {
	return PanGestureRecognizerClass.New()
}

func (p_ PanGestureRecognizer) Init() PanGestureRecognizer {
	rv := objc.Call[PanGestureRecognizer](p_, objc.Sel("init"))
	return rv
}

func (p_ PanGestureRecognizer) InitWithTargetAction(target objc.IObject, action objc.Selector) PanGestureRecognizer {
	rv := objc.Call[PanGestureRecognizer](p_, objc.Sel("initWithTarget:action:"), target, action)
	return rv
}

// Initializes the gesture recognizer with the specified target and action information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1535012-initwithtarget?language=objc
func NewPanGestureRecognizerWithTargetAction(target objc.IObject, action objc.Selector) PanGestureRecognizer {
	instance := PanGestureRecognizerClass.Alloc().InitWithTargetAction(target, action)
	instance.Autorelease()
	return instance
}

// Changes the current translation value of the gesture recognizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspangesturerecognizer/1515533-settranslation?language=objc
func (p_ PanGestureRecognizer) SetTranslationInView(translation foundation.Point, view IView) {
	objc.Call[objc.Void](p_, objc.Sel("setTranslation:inView:"), translation, objc.Ptr(view))
}

// The velocity of the pan, measured in points per second. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspangesturerecognizer/1515532-velocityinview?language=objc
func (p_ PanGestureRecognizer) VelocityInView(view IView) foundation.Point {
	rv := objc.Call[foundation.Point](p_, objc.Sel("velocityInView:"), objc.Ptr(view))
	return rv
}

// The distance traveled by the mouse during the gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspangesturerecognizer/1515531-translationinview?language=objc
func (p_ PanGestureRecognizer) TranslationInView(view IView) foundation.Point {
	rv := objc.Call[foundation.Point](p_, objc.Sel("translationInView:"), objc.Ptr(view))
	return rv
}

// A bit mask of the button (or buttons) required to recognize this gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspangesturerecognizer/1515529-buttonmask?language=objc
func (p_ PanGestureRecognizer) ButtonMask() uint {
	rv := objc.Call[uint](p_, objc.Sel("buttonMask"))
	return rv
}

// A bit mask of the button (or buttons) required to recognize this gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspangesturerecognizer/1515529-buttonmask?language=objc
func (p_ PanGestureRecognizer) SetButtonMask(value uint) {
	objc.Call[objc.Void](p_, objc.Sel("setButtonMask:"), value)
}

// The number of necessary touches on a Touch Bar for the gesture recognizer to match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspangesturerecognizer/2544781-numberoftouchesrequired?language=objc
func (p_ PanGestureRecognizer) NumberOfTouchesRequired() int {
	rv := objc.Call[int](p_, objc.Sel("numberOfTouchesRequired"))
	return rv
}

// The number of necessary touches on a Touch Bar for the gesture recognizer to match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspangesturerecognizer/2544781-numberoftouchesrequired?language=objc
func (p_ PanGestureRecognizer) SetNumberOfTouchesRequired(value int) {
	objc.Call[objc.Void](p_, objc.Sel("setNumberOfTouchesRequired:"), value)
}
