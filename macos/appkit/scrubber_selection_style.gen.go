// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScrubberSelectionStyle] class.
var ScrubberSelectionStyleClass = _ScrubberSelectionStyleClass{objc.GetClass("NSScrubberSelectionStyle")}

type _ScrubberSelectionStyleClass struct {
	objc.Class
}

// An interface definition for the [ScrubberSelectionStyle] class.
type IScrubberSelectionStyle interface {
	objc.IObject
	MakeSelectionView() ScrubberSelectionView
}

// An abstract class that provides decorative accessory views for selected and highlighted items within a scrubber control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberselectionstyle?language=objc
type ScrubberSelectionStyle struct {
	objc.Object
}

func ScrubberSelectionStyleFrom(ptr unsafe.Pointer) ScrubberSelectionStyle {
	return ScrubberSelectionStyle{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ ScrubberSelectionStyle) Init() ScrubberSelectionStyle {
	rv := objc.Call[ScrubberSelectionStyle](s_, objc.Sel("init"))
	return rv
}

func (sc _ScrubberSelectionStyleClass) Alloc() ScrubberSelectionStyle {
	rv := objc.Call[ScrubberSelectionStyle](sc, objc.Sel("alloc"))
	return rv
}

func ScrubberSelectionStyle_Alloc() ScrubberSelectionStyle {
	return ScrubberSelectionStyleClass.Alloc()
}

func (sc _ScrubberSelectionStyleClass) New() ScrubberSelectionStyle {
	rv := objc.Call[ScrubberSelectionStyle](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScrubberSelectionStyle() ScrubberSelectionStyle {
	return ScrubberSelectionStyleClass.New()
}

// Provides an opportunity to create a customized scrubber selection style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberselectionstyle/2588257-makeselectionview?language=objc
func (s_ ScrubberSelectionStyle) MakeSelectionView() ScrubberSelectionView {
	rv := objc.Call[ScrubberSelectionView](s_, objc.Sel("makeSelectionView"))
	return rv
}

// A built-in selection style that draws a rounded rectangle as the background of the scrubber item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberselectionstyle/2588261-roundedbackgroundstyle?language=objc
func (sc _ScrubberSelectionStyleClass) RoundedBackgroundStyle() ScrubberSelectionStyle {
	rv := objc.Call[ScrubberSelectionStyle](sc, objc.Sel("roundedBackgroundStyle"))
	return rv
}

// A built-in selection style that draws a rounded rectangle as the background of the scrubber item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberselectionstyle/2588261-roundedbackgroundstyle?language=objc
func ScrubberSelectionStyle_RoundedBackgroundStyle() ScrubberSelectionStyle {
	return ScrubberSelectionStyleClass.RoundedBackgroundStyle()
}

// A built-in selection style that draws the outline of the scrubber item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberselectionstyle/2588266-outlineoverlaystyle?language=objc
func (sc _ScrubberSelectionStyleClass) OutlineOverlayStyle() ScrubberSelectionStyle {
	rv := objc.Call[ScrubberSelectionStyle](sc, objc.Sel("outlineOverlayStyle"))
	return rv
}

// A built-in selection style that draws the outline of the scrubber item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberselectionstyle/2588266-outlineoverlaystyle?language=objc
func ScrubberSelectionStyle_OutlineOverlayStyle() ScrubberSelectionStyle {
	return ScrubberSelectionStyleClass.OutlineOverlayStyle()
}
