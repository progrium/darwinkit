// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScrubberItemView] class.
var ScrubberItemViewClass = _ScrubberItemViewClass{objc.GetClass("NSScrubberItemView")}

type _ScrubberItemViewClass struct {
	objc.Class
}

// An interface definition for the [ScrubberItemView] class.
type IScrubberItemView interface {
	IScrubberArrangedView
}

// An item at a specific index position in the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberitemview?language=objc
type ScrubberItemView struct {
	ScrubberArrangedView
}

func ScrubberItemViewFrom(ptr unsafe.Pointer) ScrubberItemView {
	return ScrubberItemView{
		ScrubberArrangedView: ScrubberArrangedViewFrom(ptr),
	}
}

func (sc _ScrubberItemViewClass) Alloc() ScrubberItemView {
	rv := objc.Call[ScrubberItemView](sc, objc.Sel("alloc"))
	return rv
}

func ScrubberItemView_Alloc() ScrubberItemView {
	return ScrubberItemViewClass.Alloc()
}

func (sc _ScrubberItemViewClass) New() ScrubberItemView {
	rv := objc.Call[ScrubberItemView](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScrubberItemView() ScrubberItemView {
	return ScrubberItemViewClass.New()
}

func (s_ ScrubberItemView) Init() ScrubberItemView {
	rv := objc.Call[ScrubberItemView](s_, objc.Sel("init"))
	return rv
}

func (s_ ScrubberItemView) InitWithFrame(frameRect foundation.Rect) ScrubberItemView {
	rv := objc.Call[ScrubberItemView](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func ScrubberItemView_InitWithFrame(frameRect foundation.Rect) ScrubberItemView {
	return ScrubberItemViewClass.Alloc().InitWithFrame(frameRect)
}
