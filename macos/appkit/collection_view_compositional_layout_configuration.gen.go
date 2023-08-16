// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionViewCompositionalLayoutConfiguration] class.
var CollectionViewCompositionalLayoutConfigurationClass = _CollectionViewCompositionalLayoutConfigurationClass{objc.GetClass("NSCollectionViewCompositionalLayoutConfiguration")}

type _CollectionViewCompositionalLayoutConfigurationClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewCompositionalLayoutConfiguration] class.
type ICollectionViewCompositionalLayoutConfiguration interface {
	objc.IObject
	ScrollDirection() CollectionViewScrollDirection
	SetScrollDirection(value CollectionViewScrollDirection)
	InterSectionSpacing() float64
	SetInterSectionSpacing(value float64)
	BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem)
}

// An object that defines scroll direction, section spacing, and headers or footers for the layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayoutconfiguration?language=objc
type CollectionViewCompositionalLayoutConfiguration struct {
	objc.Object
}

func CollectionViewCompositionalLayoutConfigurationFrom(ptr unsafe.Pointer) CollectionViewCompositionalLayoutConfiguration {
	return CollectionViewCompositionalLayoutConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CollectionViewCompositionalLayoutConfigurationClass) Alloc() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.Call[CollectionViewCompositionalLayoutConfiguration](cc, objc.Sel("alloc"))
	return rv
}

func CollectionViewCompositionalLayoutConfiguration_Alloc() CollectionViewCompositionalLayoutConfiguration {
	return CollectionViewCompositionalLayoutConfigurationClass.Alloc()
}

func (cc _CollectionViewCompositionalLayoutConfigurationClass) New() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.Call[CollectionViewCompositionalLayoutConfiguration](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewCompositionalLayoutConfiguration() CollectionViewCompositionalLayoutConfiguration {
	return CollectionViewCompositionalLayoutConfigurationClass.New()
}

func (c_ CollectionViewCompositionalLayoutConfiguration) Init() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.Call[CollectionViewCompositionalLayoutConfiguration](c_, objc.Sel("init"))
	return rv
}

// The axis that the content in the collection view layout scrolls along. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayoutconfiguration/3281813-scrolldirection?language=objc
func (c_ CollectionViewCompositionalLayoutConfiguration) ScrollDirection() CollectionViewScrollDirection {
	rv := objc.Call[CollectionViewScrollDirection](c_, objc.Sel("scrollDirection"))
	return rv
}

// The axis that the content in the collection view layout scrolls along. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayoutconfiguration/3281813-scrolldirection?language=objc
func (c_ CollectionViewCompositionalLayoutConfiguration) SetScrollDirection(value CollectionViewScrollDirection) {
	objc.Call[objc.Void](c_, objc.Sel("setScrollDirection:"), value)
}

// The amount of space between the sections in the layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayoutconfiguration/3281812-intersectionspacing?language=objc
func (c_ CollectionViewCompositionalLayoutConfiguration) InterSectionSpacing() float64 {
	rv := objc.Call[float64](c_, objc.Sel("interSectionSpacing"))
	return rv
}

// The amount of space between the sections in the layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayoutconfiguration/3281812-intersectionspacing?language=objc
func (c_ CollectionViewCompositionalLayoutConfiguration) SetInterSectionSpacing(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setInterSectionSpacing:"), value)
}

// An array of the supplementary items that are associated with the boundary edges of the entire layout, such as global headers and footers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayoutconfiguration/3281811-boundarysupplementaryitems?language=objc
func (c_ CollectionViewCompositionalLayoutConfiguration) BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Call[[]CollectionLayoutBoundarySupplementaryItem](c_, objc.Sel("boundarySupplementaryItems"))
	return rv
}

// An array of the supplementary items that are associated with the boundary edges of the entire layout, such as global headers and footers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayoutconfiguration/3281811-boundarysupplementaryitems?language=objc
func (c_ CollectionViewCompositionalLayoutConfiguration) SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem) {
	objc.Call[objc.Void](c_, objc.Sel("setBoundarySupplementaryItems:"), value)
}
