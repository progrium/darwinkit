// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScrubberFlowLayout] class.
var ScrubberFlowLayoutClass = _ScrubberFlowLayoutClass{objc.GetClass("NSScrubberFlowLayout")}

type _ScrubberFlowLayoutClass struct {
	objc.Class
}

// An interface definition for the [ScrubberFlowLayout] class.
type IScrubberFlowLayout interface {
	IScrubberLayout
	InvalidateLayoutForItemsAtIndexes(invalidItemIndexes foundation.IIndexSet)
	ItemSize() foundation.Size
	SetItemSize(value foundation.Size)
	ItemSpacing() float64
	SetItemSpacing(value float64)
}

// A concrete layout object that arranges items end-to-end in a linear strip. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberflowlayout?language=objc
type ScrubberFlowLayout struct {
	ScrubberLayout
}

func ScrubberFlowLayoutFrom(ptr unsafe.Pointer) ScrubberFlowLayout {
	return ScrubberFlowLayout{
		ScrubberLayout: ScrubberLayoutFrom(ptr),
	}
}

func (sc _ScrubberFlowLayoutClass) Alloc() ScrubberFlowLayout {
	rv := objc.Call[ScrubberFlowLayout](sc, objc.Sel("alloc"))
	return rv
}

func ScrubberFlowLayout_Alloc() ScrubberFlowLayout {
	return ScrubberFlowLayoutClass.Alloc()
}

func (sc _ScrubberFlowLayoutClass) New() ScrubberFlowLayout {
	rv := objc.Call[ScrubberFlowLayout](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScrubberFlowLayout() ScrubberFlowLayout {
	return ScrubberFlowLayoutClass.New()
}

func (s_ ScrubberFlowLayout) Init() ScrubberFlowLayout {
	rv := objc.Call[ScrubberFlowLayout](s_, objc.Sel("init"))
	return rv
}

// Informs the scrubber that it should perform a new layout pass for the items at the specified indexes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberflowlayout/2544642-invalidatelayoutforitemsatindexe?language=objc
func (s_ ScrubberFlowLayout) InvalidateLayoutForItemsAtIndexes(invalidItemIndexes foundation.IIndexSet) {
	objc.Call[objc.Void](s_, objc.Sel("invalidateLayoutForItemsAtIndexes:"), objc.Ptr(invalidItemIndexes))
}

// The frame size for each item in the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberflowlayout/2544627-itemsize?language=objc
func (s_ ScrubberFlowLayout) ItemSize() foundation.Size {
	rv := objc.Call[foundation.Size](s_, objc.Sel("itemSize"))
	return rv
}

// The frame size for each item in the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberflowlayout/2544627-itemsize?language=objc
func (s_ ScrubberFlowLayout) SetItemSize(value foundation.Size) {
	objc.Call[objc.Void](s_, objc.Sel("setItemSize:"), value)
}

// The horizontal spacing between items, specified in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberflowlayout/2544649-itemspacing?language=objc
func (s_ ScrubberFlowLayout) ItemSpacing() float64 {
	rv := objc.Call[float64](s_, objc.Sel("itemSpacing"))
	return rv
}

// The horizontal spacing between items, specified in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberflowlayout/2544649-itemspacing?language=objc
func (s_ ScrubberFlowLayout) SetItemSpacing(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setItemSpacing:"), value)
}
