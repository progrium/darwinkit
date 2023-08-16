// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScrubberProportionalLayout] class.
var ScrubberProportionalLayoutClass = _ScrubberProportionalLayoutClass{objc.GetClass("NSScrubberProportionalLayout")}

type _ScrubberProportionalLayoutClass struct {
	objc.Class
}

// An interface definition for the [ScrubberProportionalLayout] class.
type IScrubberProportionalLayout interface {
	IScrubberLayout
	NumberOfVisibleItems() int
	SetNumberOfVisibleItems(value int)
}

// A concrete layout object that sizes each item to some fraction of the scrubber's visible size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberproportionallayout?language=objc
type ScrubberProportionalLayout struct {
	ScrubberLayout
}

func ScrubberProportionalLayoutFrom(ptr unsafe.Pointer) ScrubberProportionalLayout {
	return ScrubberProportionalLayout{
		ScrubberLayout: ScrubberLayoutFrom(ptr),
	}
}

func (s_ ScrubberProportionalLayout) InitWithNumberOfVisibleItems(numberOfVisibleItems int) ScrubberProportionalLayout {
	rv := objc.Call[ScrubberProportionalLayout](s_, objc.Sel("initWithNumberOfVisibleItems:"), numberOfVisibleItems)
	return rv
}

// Initializes and returns a newly allocated proportional layout, configured to display the given number of items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberproportionallayout/2544646-initwithnumberofvisibleitems?language=objc
func ScrubberProportionalLayout_InitWithNumberOfVisibleItems(numberOfVisibleItems int) ScrubberProportionalLayout {
	return ScrubberProportionalLayoutClass.Alloc().InitWithNumberOfVisibleItems(numberOfVisibleItems)
}

func (sc _ScrubberProportionalLayoutClass) Alloc() ScrubberProportionalLayout {
	rv := objc.Call[ScrubberProportionalLayout](sc, objc.Sel("alloc"))
	return rv
}

func ScrubberProportionalLayout_Alloc() ScrubberProportionalLayout {
	return ScrubberProportionalLayoutClass.Alloc()
}

func (sc _ScrubberProportionalLayoutClass) New() ScrubberProportionalLayout {
	rv := objc.Call[ScrubberProportionalLayout](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScrubberProportionalLayout() ScrubberProportionalLayout {
	return ScrubberProportionalLayoutClass.New()
}

func (s_ ScrubberProportionalLayout) Init() ScrubberProportionalLayout {
	rv := objc.Call[ScrubberProportionalLayout](s_, objc.Sel("init"))
	return rv
}

// The number of items visible in the scrubber at once. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberproportionallayout/2544651-numberofvisibleitems?language=objc
func (s_ ScrubberProportionalLayout) NumberOfVisibleItems() int {
	rv := objc.Call[int](s_, objc.Sel("numberOfVisibleItems"))
	return rv
}

// The number of items visible in the scrubber at once. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberproportionallayout/2544651-numberofvisibleitems?language=objc
func (s_ ScrubberProportionalLayout) SetNumberOfVisibleItems(value int) {
	objc.Call[objc.Void](s_, objc.Sel("setNumberOfVisibleItems:"), value)
}
