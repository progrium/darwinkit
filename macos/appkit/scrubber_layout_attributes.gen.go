// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScrubberLayoutAttributes] class.
var ScrubberLayoutAttributesClass = _ScrubberLayoutAttributesClass{objc.GetClass("NSScrubberLayoutAttributes")}

type _ScrubberLayoutAttributesClass struct {
	objc.Class
}

// An interface definition for the [ScrubberLayoutAttributes] class.
type IScrubberLayoutAttributes interface {
	objc.IObject
	Alpha() float64
	SetAlpha(value float64)
	Frame() foundation.Rect
	SetFrame(value foundation.Rect)
	ItemIndex() int
	SetItemIndex(value int)
}

// The layout of a scrubber item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes?language=objc
type ScrubberLayoutAttributes struct {
	objc.Object
}

func ScrubberLayoutAttributesFrom(ptr unsafe.Pointer) ScrubberLayoutAttributes {
	return ScrubberLayoutAttributes{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _ScrubberLayoutAttributesClass) LayoutAttributesForItemAtIndex(index int) ScrubberLayoutAttributes {
	rv := objc.Call[ScrubberLayoutAttributes](sc, objc.Sel("layoutAttributesForItemAtIndex:"), index)
	return rv
}

// Creates a new layout attributes object for the specified scrubber item index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/2544645-layoutattributesforitematindex?language=objc
func ScrubberLayoutAttributes_LayoutAttributesForItemAtIndex(index int) ScrubberLayoutAttributes {
	return ScrubberLayoutAttributesClass.LayoutAttributesForItemAtIndex(index)
}

func (sc _ScrubberLayoutAttributesClass) Alloc() ScrubberLayoutAttributes {
	rv := objc.Call[ScrubberLayoutAttributes](sc, objc.Sel("alloc"))
	return rv
}

func ScrubberLayoutAttributes_Alloc() ScrubberLayoutAttributes {
	return ScrubberLayoutAttributesClass.Alloc()
}

func (sc _ScrubberLayoutAttributesClass) New() ScrubberLayoutAttributes {
	rv := objc.Call[ScrubberLayoutAttributes](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScrubberLayoutAttributes() ScrubberLayoutAttributes {
	return ScrubberLayoutAttributesClass.New()
}

func (s_ ScrubberLayoutAttributes) Init() ScrubberLayoutAttributes {
	rv := objc.Call[ScrubberLayoutAttributes](s_, objc.Sel("init"))
	return rv
}

// The item's alpha value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/2544648-alpha?language=objc
func (s_ ScrubberLayoutAttributes) Alpha() float64 {
	rv := objc.Call[float64](s_, objc.Sel("alpha"))
	return rv
}

// The item's alpha value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/2544648-alpha?language=objc
func (s_ ScrubberLayoutAttributes) SetAlpha(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setAlpha:"), value)
}

// The frame of the scrubber item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/2544623-frame?language=objc
func (s_ ScrubberLayoutAttributes) Frame() foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("frame"))
	return rv
}

// The frame of the scrubber item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/2544623-frame?language=objc
func (s_ ScrubberLayoutAttributes) SetFrame(value foundation.Rect) {
	objc.Call[objc.Void](s_, objc.Sel("setFrame:"), value)
}

// The index of the scrubber item that is represented by the item's layout attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/2544635-itemindex?language=objc
func (s_ ScrubberLayoutAttributes) ItemIndex() int {
	rv := objc.Call[int](s_, objc.Sel("itemIndex"))
	return rv
}

// The index of the scrubber item that is represented by the item's layout attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/2544635-itemindex?language=objc
func (s_ ScrubberLayoutAttributes) SetItemIndex(value int) {
	objc.Call[objc.Void](s_, objc.Sel("setItemIndex:"), value)
}
