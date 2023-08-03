// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var CollectionViewFlowLayoutClass = _CollectionViewFlowLayoutClass{objc.GetClass("NSCollectionViewFlowLayout")}

type _CollectionViewFlowLayoutClass struct {
	objc.Class
}

type ICollectionViewFlowLayout interface {
	ICollectionViewLayout
	CollapseSectionAtIndex(sectionIndex uint)
	ExpandSectionAtIndex(sectionIndex uint)
	SectionAtIndexIsCollapsed(sectionIndex uint) bool
	ScrollDirection() CollectionViewScrollDirection
	SetScrollDirection(value CollectionViewScrollDirection)
	MinimumLineSpacing() float64
	SetMinimumLineSpacing(value float64)
	MinimumInteritemSpacing() float64
	SetMinimumInteritemSpacing(value float64)
	EstimatedItemSize() foundation.Size
	SetEstimatedItemSize(value foundation.Size)
	ItemSize() foundation.Size
	SetItemSize(value foundation.Size)
	SectionInset() foundation.EdgeInsets
	SetSectionInset(value foundation.EdgeInsets)
	HeaderReferenceSize() foundation.Size
	SetHeaderReferenceSize(value foundation.Size)
	FooterReferenceSize() foundation.Size
	SetFooterReferenceSize(value foundation.Size)
	SectionFootersPinToVisibleBounds() bool
	SetSectionFootersPinToVisibleBounds(value bool)
	SectionHeadersPinToVisibleBounds() bool
	SetSectionHeadersPinToVisibleBounds(value bool)
}

type CollectionViewFlowLayout struct {
	CollectionViewLayout
}

func MakeCollectionViewFlowLayout(ptr unsafe.Pointer) CollectionViewFlowLayout {
	return CollectionViewFlowLayout{
		CollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func (cc _CollectionViewFlowLayoutClass) Alloc() CollectionViewFlowLayout {
	rv := objc.CallMethod[CollectionViewFlowLayout](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionViewFlowLayout_Alloc() CollectionViewFlowLayout {
	return CollectionViewFlowLayoutClass.Alloc()
}

func (cc _CollectionViewFlowLayoutClass) New() CollectionViewFlowLayout {
	rv := objc.CallMethod[CollectionViewFlowLayout](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewFlowLayout() CollectionViewFlowLayout {
	return CollectionViewFlowLayoutClass.New()
}

func CollectionViewFlowLayout_New() CollectionViewFlowLayout {
	return CollectionViewFlowLayoutClass.New()
}

func (c_ CollectionViewFlowLayout) Init() CollectionViewFlowLayout {
	rv := objc.CallMethod[CollectionViewFlowLayout](c_, objc.GetSelector("init"))
	return rv
}

func CollectionViewFlowLayout_Init() CollectionViewFlowLayout {
	return CollectionViewFlowLayoutClass.Alloc().Init()
}

func (c_ CollectionViewFlowLayout) CollapseSectionAtIndex(sectionIndex uint) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collapseSectionAtIndex:"), sectionIndex)
}

func (c_ CollectionViewFlowLayout) ExpandSectionAtIndex(sectionIndex uint) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("expandSectionAtIndex:"), sectionIndex)
}

func (c_ CollectionViewFlowLayout) SectionAtIndexIsCollapsed(sectionIndex uint) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("sectionAtIndexIsCollapsed:"), sectionIndex)
	return rv
}

func (c_ CollectionViewFlowLayout) ScrollDirection() CollectionViewScrollDirection {
	rv := objc.CallMethod[CollectionViewScrollDirection](c_, objc.GetSelector("scrollDirection"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetScrollDirection(value CollectionViewScrollDirection) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setScrollDirection:"), value)
}

func (c_ CollectionViewFlowLayout) MinimumLineSpacing() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("minimumLineSpacing"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetMinimumLineSpacing(value float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMinimumLineSpacing:"), value)
}

func (c_ CollectionViewFlowLayout) MinimumInteritemSpacing() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("minimumInteritemSpacing"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetMinimumInteritemSpacing(value float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMinimumInteritemSpacing:"), value)
}

func (c_ CollectionViewFlowLayout) EstimatedItemSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("estimatedItemSize"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetEstimatedItemSize(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setEstimatedItemSize:"), value)
}

func (c_ CollectionViewFlowLayout) ItemSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("itemSize"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetItemSize(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setItemSize:"), value)
}

func (c_ CollectionViewFlowLayout) SectionInset() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](c_, objc.GetSelector("sectionInset"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetSectionInset(value foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setSectionInset:"), value)
}

func (c_ CollectionViewFlowLayout) HeaderReferenceSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("headerReferenceSize"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetHeaderReferenceSize(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setHeaderReferenceSize:"), value)
}

func (c_ CollectionViewFlowLayout) FooterReferenceSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("footerReferenceSize"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetFooterReferenceSize(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setFooterReferenceSize:"), value)
}

func (c_ CollectionViewFlowLayout) SectionFootersPinToVisibleBounds() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("sectionFootersPinToVisibleBounds"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetSectionFootersPinToVisibleBounds(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setSectionFootersPinToVisibleBounds:"), value)
}

func (c_ CollectionViewFlowLayout) SectionHeadersPinToVisibleBounds() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("sectionHeadersPinToVisibleBounds"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetSectionHeadersPinToVisibleBounds(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setSectionHeadersPinToVisibleBounds:"), value)
}
