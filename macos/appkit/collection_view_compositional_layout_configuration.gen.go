// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CollectionViewCompositionalLayoutConfigurationClass = _CollectionViewCompositionalLayoutConfigurationClass{objc.GetClass("NSCollectionViewCompositionalLayoutConfiguration")}

type _CollectionViewCompositionalLayoutConfigurationClass struct {
	objc.Class
}

type ICollectionViewCompositionalLayoutConfiguration interface {
	objc.IObject
	ScrollDirection() CollectionViewScrollDirection
	SetScrollDirection(value CollectionViewScrollDirection)
	InterSectionSpacing() float64
	SetInterSectionSpacing(value float64)
	BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem)
}

type CollectionViewCompositionalLayoutConfiguration struct {
	objc.Object
}

func MakeCollectionViewCompositionalLayoutConfiguration(ptr unsafe.Pointer) CollectionViewCompositionalLayoutConfiguration {
	return CollectionViewCompositionalLayoutConfiguration{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionViewCompositionalLayoutConfigurationClass) Alloc() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.CallMethod[CollectionViewCompositionalLayoutConfiguration](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionViewCompositionalLayoutConfiguration_Alloc() CollectionViewCompositionalLayoutConfiguration {
	return CollectionViewCompositionalLayoutConfigurationClass.Alloc()
}

func (cc _CollectionViewCompositionalLayoutConfigurationClass) New() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.CallMethod[CollectionViewCompositionalLayoutConfiguration](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewCompositionalLayoutConfiguration() CollectionViewCompositionalLayoutConfiguration {
	return CollectionViewCompositionalLayoutConfigurationClass.New()
}

func CollectionViewCompositionalLayoutConfiguration_New() CollectionViewCompositionalLayoutConfiguration {
	return CollectionViewCompositionalLayoutConfigurationClass.New()
}

func (c_ CollectionViewCompositionalLayoutConfiguration) Init() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.CallMethod[CollectionViewCompositionalLayoutConfiguration](c_, objc.GetSelector("init"))
	return rv
}

func CollectionViewCompositionalLayoutConfiguration_Init() CollectionViewCompositionalLayoutConfiguration {
	return CollectionViewCompositionalLayoutConfigurationClass.Alloc().Init()
}

func (c_ CollectionViewCompositionalLayoutConfiguration) ScrollDirection() CollectionViewScrollDirection {
	rv := objc.CallMethod[CollectionViewScrollDirection](c_, objc.GetSelector("scrollDirection"))
	return rv
}

func (c_ CollectionViewCompositionalLayoutConfiguration) SetScrollDirection(value CollectionViewScrollDirection) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setScrollDirection:"), value)
}

func (c_ CollectionViewCompositionalLayoutConfiguration) InterSectionSpacing() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("interSectionSpacing"))
	return rv
}

func (c_ CollectionViewCompositionalLayoutConfiguration) SetInterSectionSpacing(value float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setInterSectionSpacing:"), value)
}

func (c_ CollectionViewCompositionalLayoutConfiguration) BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[[]CollectionLayoutBoundarySupplementaryItem](c_, objc.GetSelector("boundarySupplementaryItems"))
	// TODO: convert slice items...
	return rv
}

func (c_ CollectionViewCompositionalLayoutConfiguration) SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setBoundarySupplementaryItems:"), value)
}
