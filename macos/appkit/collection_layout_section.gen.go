// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CollectionLayoutSectionClass = _CollectionLayoutSectionClass{objc.GetClass("NSCollectionLayoutSection")}

type _CollectionLayoutSectionClass struct {
	objc.Class
}

type ICollectionLayoutSection interface {
	objc.IObject
	OrthogonalScrollingBehavior() CollectionLayoutSectionOrthogonalScrollingBehavior
	SetOrthogonalScrollingBehavior(value CollectionLayoutSectionOrthogonalScrollingBehavior)
	InterGroupSpacing() float64
	SetInterGroupSpacing(value float64)
	ContentInsets() DirectionalEdgeInsets
	SetContentInsets(value DirectionalEdgeInsets)
	BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem)
	DecorationItems() []CollectionLayoutDecorationItem
	SetDecorationItems(value []ICollectionLayoutDecorationItem)
	SupplementariesFollowContentInsets() bool
	SetSupplementariesFollowContentInsets(value bool)
}

type CollectionLayoutSection struct {
	objc.Object
}

func MakeCollectionLayoutSection(ptr unsafe.Pointer) CollectionLayoutSection {
	return CollectionLayoutSection{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutSectionClass) SectionWithGroup(group ICollectionLayoutGroup) CollectionLayoutSection {
	rv := objc.CallMethod[CollectionLayoutSection](cc, objc.GetSelector("sectionWithGroup:"), objc.ExtractPtr(group))
	return rv
}

func CollectionLayoutSection_SectionWithGroup(group ICollectionLayoutGroup) CollectionLayoutSection {
	return CollectionLayoutSectionClass.SectionWithGroup(group)
}

func (cc _CollectionLayoutSectionClass) Alloc() CollectionLayoutSection {
	rv := objc.CallMethod[CollectionLayoutSection](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionLayoutSection_Alloc() CollectionLayoutSection {
	return CollectionLayoutSectionClass.Alloc()
}

func (cc _CollectionLayoutSectionClass) New() CollectionLayoutSection {
	rv := objc.CallMethod[CollectionLayoutSection](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSection() CollectionLayoutSection {
	return CollectionLayoutSectionClass.New()
}

func CollectionLayoutSection_New() CollectionLayoutSection {
	return CollectionLayoutSectionClass.New()
}

func (c_ CollectionLayoutSection) Init() CollectionLayoutSection {
	rv := objc.CallMethod[CollectionLayoutSection](c_, objc.GetSelector("init"))
	return rv
}

func CollectionLayoutSection_Init() CollectionLayoutSection {
	return CollectionLayoutSectionClass.Alloc().Init()
}

func (c_ CollectionLayoutSection) OrthogonalScrollingBehavior() CollectionLayoutSectionOrthogonalScrollingBehavior {
	rv := objc.CallMethod[CollectionLayoutSectionOrthogonalScrollingBehavior](c_, objc.GetSelector("orthogonalScrollingBehavior"))
	return rv
}

func (c_ CollectionLayoutSection) SetOrthogonalScrollingBehavior(value CollectionLayoutSectionOrthogonalScrollingBehavior) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setOrthogonalScrollingBehavior:"), value)
}

func (c_ CollectionLayoutSection) InterGroupSpacing() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("interGroupSpacing"))
	return rv
}

func (c_ CollectionLayoutSection) SetInterGroupSpacing(value float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setInterGroupSpacing:"), value)
}

func (c_ CollectionLayoutSection) ContentInsets() DirectionalEdgeInsets {
	rv := objc.CallMethod[DirectionalEdgeInsets](c_, objc.GetSelector("contentInsets"))
	return rv
}

func (c_ CollectionLayoutSection) SetContentInsets(value DirectionalEdgeInsets) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setContentInsets:"), value)
}

func (c_ CollectionLayoutSection) BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[[]CollectionLayoutBoundarySupplementaryItem](c_, objc.GetSelector("boundarySupplementaryItems"))
	// TODO: convert slice items...
	return rv
}

func (c_ CollectionLayoutSection) SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setBoundarySupplementaryItems:"), value)
}

func (c_ CollectionLayoutSection) DecorationItems() []CollectionLayoutDecorationItem {
	rv := objc.CallMethod[[]CollectionLayoutDecorationItem](c_, objc.GetSelector("decorationItems"))
	// TODO: convert slice items...
	return rv
}

func (c_ CollectionLayoutSection) SetDecorationItems(value []ICollectionLayoutDecorationItem) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setDecorationItems:"), value)
}

func (c_ CollectionLayoutSection) SupplementariesFollowContentInsets() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("supplementariesFollowContentInsets"))
	return rv
}

func (c_ CollectionLayoutSection) SetSupplementariesFollowContentInsets(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setSupplementariesFollowContentInsets:"), value)
}
