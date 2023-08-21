// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScrubberArrangedView] class.
var ScrubberArrangedViewClass = _ScrubberArrangedViewClass{objc.GetClass("NSScrubberArrangedView")}

type _ScrubberArrangedViewClass struct {
	objc.Class
}

// An interface definition for the [ScrubberArrangedView] class.
type IScrubberArrangedView interface {
	IView
	ApplyLayoutAttributes(layoutAttributes IScrubberLayoutAttributes)
	IsHighlighted() bool
	SetHighlighted(value bool)
	IsSelected() bool
	SetSelected(value bool)
}

// An abstract base class for the views whose layout is managed by a scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberarrangedview?language=objc
type ScrubberArrangedView struct {
	View
}

func ScrubberArrangedViewFrom(ptr unsafe.Pointer) ScrubberArrangedView {
	return ScrubberArrangedView{
		View: ViewFrom(ptr),
	}
}

func (sc _ScrubberArrangedViewClass) Alloc() ScrubberArrangedView {
	rv := objc.Call[ScrubberArrangedView](sc, objc.Sel("alloc"))
	return rv
}

func ScrubberArrangedView_Alloc() ScrubberArrangedView {
	return ScrubberArrangedViewClass.Alloc()
}

func (sc _ScrubberArrangedViewClass) New() ScrubberArrangedView {
	rv := objc.Call[ScrubberArrangedView](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScrubberArrangedView() ScrubberArrangedView {
	return ScrubberArrangedViewClass.New()
}

func (s_ ScrubberArrangedView) Init() ScrubberArrangedView {
	rv := objc.Call[ScrubberArrangedView](s_, objc.Sel("init"))
	return rv
}

func (s_ ScrubberArrangedView) InitWithFrame(frameRect foundation.Rect) ScrubberArrangedView {
	rv := objc.Call[ScrubberArrangedView](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewScrubberArrangedViewWithFrame(frameRect foundation.Rect) ScrubberArrangedView {
	instance := ScrubberArrangedViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Updates the layout of the arranged view to respect the provided layout attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberarrangedview/2588268-applylayoutattributes?language=objc
func (s_ ScrubberArrangedView) ApplyLayoutAttributes(layoutAttributes IScrubberLayoutAttributes) {
	objc.Call[objc.Void](s_, objc.Sel("applyLayoutAttributes:"), objc.Ptr(layoutAttributes))
}

// A Boolean value that specifies whether the view is currently highlighted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberarrangedview/2588255-highlighted?language=objc
func (s_ ScrubberArrangedView) IsHighlighted() bool {
	rv := objc.Call[bool](s_, objc.Sel("isHighlighted"))
	return rv
}

// A Boolean value that specifies whether the view is currently highlighted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberarrangedview/2588255-highlighted?language=objc
func (s_ ScrubberArrangedView) SetHighlighted(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setHighlighted:"), value)
}

// A Boolean value that specifies whether the current view is selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberarrangedview/2588256-selected?language=objc
func (s_ ScrubberArrangedView) IsSelected() bool {
	rv := objc.Call[bool](s_, objc.Sel("isSelected"))
	return rv
}

// A Boolean value that specifies whether the current view is selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberarrangedview/2588256-selected?language=objc
func (s_ ScrubberArrangedView) SetSelected(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setSelected:"), value)
}
