// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionViewFlowLayout] class.
var CollectionViewFlowLayoutClass = _CollectionViewFlowLayoutClass{objc.GetClass("NSCollectionViewFlowLayout")}

type _CollectionViewFlowLayoutClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewFlowLayout] class.
type ICollectionViewFlowLayout interface {
	ICollectionViewLayout
	ExpandSectionAtIndex(sectionIndex uint)
	SectionAtIndexIsCollapsed(sectionIndex uint) bool
	CollapseSectionAtIndex(sectionIndex uint)
	ItemSize() foundation.Size
	SetItemSize(value foundation.Size)
	HeaderReferenceSize() foundation.Size
	SetHeaderReferenceSize(value foundation.Size)
	MinimumInteritemSpacing() float64
	SetMinimumInteritemSpacing(value float64)
	FooterReferenceSize() foundation.Size
	SetFooterReferenceSize(value foundation.Size)
	SectionInset() foundation.EdgeInsets
	SetSectionInset(value foundation.EdgeInsets)
	ScrollDirection() CollectionViewScrollDirection
	SetScrollDirection(value CollectionViewScrollDirection)
	EstimatedItemSize() foundation.Size
	SetEstimatedItemSize(value foundation.Size)
	SectionFootersPinToVisibleBounds() bool
	SetSectionFootersPinToVisibleBounds(value bool)
	MinimumLineSpacing() float64
	SetMinimumLineSpacing(value float64)
	SectionHeadersPinToVisibleBounds() bool
	SetSectionHeadersPinToVisibleBounds(value bool)
}

// A layout that organizes items into a flexible and configurable arrangement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout?language=objc
type CollectionViewFlowLayout struct {
	CollectionViewLayout
}

func CollectionViewFlowLayoutFrom(ptr unsafe.Pointer) CollectionViewFlowLayout {
	return CollectionViewFlowLayout{
		CollectionViewLayout: CollectionViewLayoutFrom(ptr),
	}
}

func (cc _CollectionViewFlowLayoutClass) Alloc() CollectionViewFlowLayout {
	rv := objc.Call[CollectionViewFlowLayout](cc, objc.Sel("alloc"))
	return rv
}

func CollectionViewFlowLayout_Alloc() CollectionViewFlowLayout {
	return CollectionViewFlowLayoutClass.Alloc()
}

func (cc _CollectionViewFlowLayoutClass) New() CollectionViewFlowLayout {
	rv := objc.Call[CollectionViewFlowLayout](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewFlowLayout() CollectionViewFlowLayout {
	return CollectionViewFlowLayoutClass.New()
}

func (c_ CollectionViewFlowLayout) Init() CollectionViewFlowLayout {
	rv := objc.Call[CollectionViewFlowLayout](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1644651-expandsectionatindex?language=objc
func (c_ CollectionViewFlowLayout) ExpandSectionAtIndex(sectionIndex uint) {
	objc.Call[objc.Void](c_, objc.Sel("expandSectionAtIndex:"), sectionIndex)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1644596-sectionatindexiscollapsed?language=objc
func (c_ CollectionViewFlowLayout) SectionAtIndexIsCollapsed(sectionIndex uint) bool {
	rv := objc.Call[bool](c_, objc.Sel("sectionAtIndexIsCollapsed:"), sectionIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1644723-collapsesectionatindex?language=objc
func (c_ CollectionViewFlowLayout) CollapseSectionAtIndex(sectionIndex uint) {
	objc.Call[objc.Void](c_, objc.Sel("collapseSectionAtIndex:"), sectionIndex)
}

// The default size to use for items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402863-itemsize?language=objc
func (c_ CollectionViewFlowLayout) ItemSize() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("itemSize"))
	return rv
}

// The default size to use for items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402863-itemsize?language=objc
func (c_ CollectionViewFlowLayout) SetItemSize(value foundation.Size) {
	objc.Call[objc.Void](c_, objc.Sel("setItemSize:"), value)
}

// The default size to use for section headers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402893-headerreferencesize?language=objc
func (c_ CollectionViewFlowLayout) HeaderReferenceSize() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("headerReferenceSize"))
	return rv
}

// The default size to use for section headers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402893-headerreferencesize?language=objc
func (c_ CollectionViewFlowLayout) SetHeaderReferenceSize(value foundation.Size) {
	objc.Call[objc.Void](c_, objc.Sel("setHeaderReferenceSize:"), value)
}

// The minimum spacing (in points) to use between items in the same row or column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402872-minimuminteritemspacing?language=objc
func (c_ CollectionViewFlowLayout) MinimumInteritemSpacing() float64 {
	rv := objc.Call[float64](c_, objc.Sel("minimumInteritemSpacing"))
	return rv
}

// The minimum spacing (in points) to use between items in the same row or column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402872-minimuminteritemspacing?language=objc
func (c_ CollectionViewFlowLayout) SetMinimumInteritemSpacing(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setMinimumInteritemSpacing:"), value)
}

// The default size to use for section footers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402892-footerreferencesize?language=objc
func (c_ CollectionViewFlowLayout) FooterReferenceSize() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("footerReferenceSize"))
	return rv
}

// The default size to use for section footers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402892-footerreferencesize?language=objc
func (c_ CollectionViewFlowLayout) SetFooterReferenceSize(value foundation.Size) {
	objc.Call[objc.Void](c_, objc.Sel("setFooterReferenceSize:"), value)
}

// The margins used to lay out content in a section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402862-sectioninset?language=objc
func (c_ CollectionViewFlowLayout) SectionInset() foundation.EdgeInsets {
	rv := objc.Call[foundation.EdgeInsets](c_, objc.Sel("sectionInset"))
	return rv
}

// The margins used to lay out content in a section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402862-sectioninset?language=objc
func (c_ CollectionViewFlowLayout) SetSectionInset(value foundation.EdgeInsets) {
	objc.Call[objc.Void](c_, objc.Sel("setSectionInset:"), value)
}

// The scroll direction of the layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402869-scrolldirection?language=objc
func (c_ CollectionViewFlowLayout) ScrollDirection() CollectionViewScrollDirection {
	rv := objc.Call[CollectionViewScrollDirection](c_, objc.Sel("scrollDirection"))
	return rv
}

// The scroll direction of the layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402869-scrolldirection?language=objc
func (c_ CollectionViewFlowLayout) SetScrollDirection(value CollectionViewScrollDirection) {
	objc.Call[objc.Void](c_, objc.Sel("setScrollDirection:"), value)
}

// The estimated size of items in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402896-estimateditemsize?language=objc
func (c_ CollectionViewFlowLayout) EstimatedItemSize() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("estimatedItemSize"))
	return rv
}

// The estimated size of items in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402896-estimateditemsize?language=objc
func (c_ CollectionViewFlowLayout) SetEstimatedItemSize(value foundation.Size) {
	objc.Call[objc.Void](c_, objc.Sel("setEstimatedItemSize:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1644671-sectionfooterspintovisiblebounds?language=objc
func (c_ CollectionViewFlowLayout) SectionFootersPinToVisibleBounds() bool {
	rv := objc.Call[bool](c_, objc.Sel("sectionFootersPinToVisibleBounds"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1644671-sectionfooterspintovisiblebounds?language=objc
func (c_ CollectionViewFlowLayout) SetSectionFootersPinToVisibleBounds(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setSectionFootersPinToVisibleBounds:"), value)
}

// The minimum spacing (in points) to use between rows or columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402898-minimumlinespacing?language=objc
func (c_ CollectionViewFlowLayout) MinimumLineSpacing() float64 {
	rv := objc.Call[float64](c_, objc.Sel("minimumLineSpacing"))
	return rv
}

// The minimum spacing (in points) to use between rows or columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1402898-minimumlinespacing?language=objc
func (c_ CollectionViewFlowLayout) SetMinimumLineSpacing(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setMinimumLineSpacing:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1644640-sectionheaderspintovisiblebounds?language=objc
func (c_ CollectionViewFlowLayout) SectionHeadersPinToVisibleBounds() bool {
	rv := objc.Call[bool](c_, objc.Sel("sectionHeadersPinToVisibleBounds"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/1644640-sectionheaderspintovisiblebounds?language=objc
func (c_ CollectionViewFlowLayout) SetSectionHeadersPinToVisibleBounds(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setSectionHeadersPinToVisibleBounds:"), value)
}
