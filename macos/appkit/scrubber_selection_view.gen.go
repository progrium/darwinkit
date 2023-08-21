// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScrubberSelectionView] class.
var ScrubberSelectionViewClass = _ScrubberSelectionViewClass{objc.GetClass("NSScrubberSelectionView")}

type _ScrubberSelectionViewClass struct {
	objc.Class
}

// An interface definition for the [ScrubberSelectionView] class.
type IScrubberSelectionView interface {
	IScrubberArrangedView
}

// An abstract base class for specifying the appearance of a highlighted or selected item in a scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberselectionview?language=objc
type ScrubberSelectionView struct {
	ScrubberArrangedView
}

func ScrubberSelectionViewFrom(ptr unsafe.Pointer) ScrubberSelectionView {
	return ScrubberSelectionView{
		ScrubberArrangedView: ScrubberArrangedViewFrom(ptr),
	}
}

func (sc _ScrubberSelectionViewClass) Alloc() ScrubberSelectionView {
	rv := objc.Call[ScrubberSelectionView](sc, objc.Sel("alloc"))
	return rv
}

func ScrubberSelectionView_Alloc() ScrubberSelectionView {
	return ScrubberSelectionViewClass.Alloc()
}

func (sc _ScrubberSelectionViewClass) New() ScrubberSelectionView {
	rv := objc.Call[ScrubberSelectionView](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScrubberSelectionView() ScrubberSelectionView {
	return ScrubberSelectionViewClass.New()
}

func (s_ ScrubberSelectionView) Init() ScrubberSelectionView {
	rv := objc.Call[ScrubberSelectionView](s_, objc.Sel("init"))
	return rv
}

func (s_ ScrubberSelectionView) InitWithFrame(frameRect foundation.Rect) ScrubberSelectionView {
	rv := objc.Call[ScrubberSelectionView](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewScrubberSelectionViewWithFrame(frameRect foundation.Rect) ScrubberSelectionView {
	instance := ScrubberSelectionViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}
