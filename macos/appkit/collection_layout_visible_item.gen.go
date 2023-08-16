// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An item that’s currently visible within the bounds of a section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutvisibleitem?language=objc
type PCollectionLayoutVisibleItem interface {
	// optional
	SetHidden(value bool)
	HasSetHidden() bool

	// optional
	IsHidden() bool
	HasIsHidden() bool

	// optional
	Name() string
	HasName() bool

	// optional
	Bounds() coregraphics.Rect
	HasBounds() bool

	// optional
	RepresentedElementKind() string
	HasRepresentedElementKind() bool

	// optional
	RepresentedElementCategory() CollectionElementCategory
	HasRepresentedElementCategory() bool

	// optional
	SetAlpha(value float64)
	HasSetAlpha() bool

	// optional
	Alpha() float64
	HasAlpha() bool

	// optional
	IndexPath() foundation.IIndexPath
	HasIndexPath() bool

	// optional
	Frame() coregraphics.Rect
	HasFrame() bool

	// optional
	SetZIndex(value int)
	HasSetZIndex() bool

	// optional
	ZIndex() int
	HasZIndex() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PCollectionLayoutVisibleItem] protocol.
type CollectionLayoutVisibleItemWrapper struct {
	objc.Object
}

func (c_ CollectionLayoutVisibleItemWrapper) HasSetHidden() bool {
	return c_.RespondsToSelector(objc.Sel("setHidden:"))
}

// A Boolean value that determines whether the item is hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutvisibleitem/3199121-hidden?language=objc
func (c_ CollectionLayoutVisibleItemWrapper) SetHidden(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setHidden:"), value)
}

func (c_ CollectionLayoutVisibleItemWrapper) HasIsHidden() bool {
	return c_.RespondsToSelector(objc.Sel("isHidden"))
}

// A Boolean value that determines whether the item is hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutvisibleitem/3199121-hidden?language=objc
func (c_ CollectionLayoutVisibleItemWrapper) IsHidden() bool {
	rv := objc.Call[bool](c_, objc.Sel("isHidden"))
	return rv
}

func (c_ CollectionLayoutVisibleItemWrapper) HasName() bool {
	return c_.RespondsToSelector(objc.Sel("name"))
}

// The name of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutvisibleitem/3199122-name?language=objc
func (c_ CollectionLayoutVisibleItemWrapper) Name() string {
	rv := objc.Call[string](c_, objc.Sel("name"))
	return rv
}

func (c_ CollectionLayoutVisibleItemWrapper) HasBounds() bool {
	return c_.RespondsToSelector(objc.Sel("bounds"))
}

// The bounds rectangle, which describes the item's location and size in its own coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutvisibleitem/3199117-bounds?language=objc
func (c_ CollectionLayoutVisibleItemWrapper) Bounds() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](c_, objc.Sel("bounds"))
	return rv
}

func (c_ CollectionLayoutVisibleItemWrapper) HasRepresentedElementKind() bool {
	return c_.RespondsToSelector(objc.Sel("representedElementKind"))
}

// A string that identifies the type of item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutvisibleitem/3199124-representedelementkind?language=objc
func (c_ CollectionLayoutVisibleItemWrapper) RepresentedElementKind() string {
	rv := objc.Call[string](c_, objc.Sel("representedElementKind"))
	return rv
}

func (c_ CollectionLayoutVisibleItemWrapper) HasRepresentedElementCategory() bool {
	return c_.RespondsToSelector(objc.Sel("representedElementCategory"))
}

// A category that identifies the item, such as decoration or supplementary view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutvisibleitem/3199123-representedelementcategory?language=objc
func (c_ CollectionLayoutVisibleItemWrapper) RepresentedElementCategory() CollectionElementCategory {
	rv := objc.Call[CollectionElementCategory](c_, objc.Sel("representedElementCategory"))
	return rv
}

func (c_ CollectionLayoutVisibleItemWrapper) HasSetAlpha() bool {
	return c_.RespondsToSelector(objc.Sel("setAlpha:"))
}

// The transparency of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutvisibleitem/3199116-alpha?language=objc
func (c_ CollectionLayoutVisibleItemWrapper) SetAlpha(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setAlpha:"), value)
}

func (c_ CollectionLayoutVisibleItemWrapper) HasAlpha() bool {
	return c_.RespondsToSelector(objc.Sel("alpha"))
}

// The transparency of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutvisibleitem/3199116-alpha?language=objc
func (c_ CollectionLayoutVisibleItemWrapper) Alpha() float64 {
	rv := objc.Call[float64](c_, objc.Sel("alpha"))
	return rv
}

func (c_ CollectionLayoutVisibleItemWrapper) HasIndexPath() bool {
	return c_.RespondsToSelector(objc.Sel("indexPath"))
}

// The index path of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutvisibleitem/3199120-indexpath?language=objc
func (c_ CollectionLayoutVisibleItemWrapper) IndexPath() foundation.IndexPath {
	rv := objc.Call[foundation.IndexPath](c_, objc.Sel("indexPath"))
	return rv
}

func (c_ CollectionLayoutVisibleItemWrapper) HasFrame() bool {
	return c_.RespondsToSelector(objc.Sel("frame"))
}

// The frame rectangle, which describes the item's location and size in its section's coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutvisibleitem/3199119-frame?language=objc
func (c_ CollectionLayoutVisibleItemWrapper) Frame() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](c_, objc.Sel("frame"))
	return rv
}

func (c_ CollectionLayoutVisibleItemWrapper) HasSetZIndex() bool {
	return c_.RespondsToSelector(objc.Sel("setZIndex:"))
}

// The vertical stacking order of the item in relation to other items in the section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutvisibleitem/3199127-zindex?language=objc
func (c_ CollectionLayoutVisibleItemWrapper) SetZIndex(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setZIndex:"), value)
}

func (c_ CollectionLayoutVisibleItemWrapper) HasZIndex() bool {
	return c_.RespondsToSelector(objc.Sel("zIndex"))
}

// The vertical stacking order of the item in relation to other items in the section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutvisibleitem/3199127-zindex?language=objc
func (c_ CollectionLayoutVisibleItemWrapper) ZIndex() int {
	rv := objc.Call[int](c_, objc.Sel("zIndex"))
	return rv
}

func (c_ CollectionLayoutVisibleItemWrapper) HasSetCenter() bool {
	return c_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The center point of the item's frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutvisibleitem/3199118-center?language=objc
func (c_ CollectionLayoutVisibleItemWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](c_, objc.Sel("setCenter:"), value)
}

func (c_ CollectionLayoutVisibleItemWrapper) HasCenter() bool {
	return c_.RespondsToSelector(objc.Sel("center"))
}

// The center point of the item's frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutvisibleitem/3199118-center?language=objc
func (c_ CollectionLayoutVisibleItemWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("center"))
	return rv
}
