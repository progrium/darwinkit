// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CollectionLayoutSpacing] class.
var CollectionLayoutSpacingClass = _CollectionLayoutSpacingClass{objc.GetClass("NSCollectionLayoutSpacing")}

type _CollectionLayoutSpacingClass struct {
	objc.Class
}

// An interface definition for the [CollectionLayoutSpacing] class.
type ICollectionLayoutSpacing interface {
	objc.IObject
	IsFlexibleSpacing() bool
	Spacing() float64
	IsFixedSpacing() bool
}

// An object that defines the space between or around items in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutspacing?language=objc
type CollectionLayoutSpacing struct {
	objc.Object
}

func CollectionLayoutSpacingFrom(ptr unsafe.Pointer) CollectionLayoutSpacing {
	return CollectionLayoutSpacing{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CollectionLayoutSpacingClass) FlexibleSpacing(flexibleSpacing float64) CollectionLayoutSpacing {
	rv := objc.Call[CollectionLayoutSpacing](cc, objc.Sel("flexibleSpacing:"), flexibleSpacing)
	return rv
}

// Creates a space equivalent to or greater than the specified number of points, depending on the available space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutspacing/3199104-flexiblespacing?language=objc
func CollectionLayoutSpacing_FlexibleSpacing(flexibleSpacing float64) CollectionLayoutSpacing {
	return CollectionLayoutSpacingClass.FlexibleSpacing(flexibleSpacing)
}

func (cc _CollectionLayoutSpacingClass) FixedSpacing(fixedSpacing float64) CollectionLayoutSpacing {
	rv := objc.Call[CollectionLayoutSpacing](cc, objc.Sel("fixedSpacing:"), fixedSpacing)
	return rv
}

// Creates a space equivalent to the specified number of points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutspacing/3199103-fixedspacing?language=objc
func CollectionLayoutSpacing_FixedSpacing(fixedSpacing float64) CollectionLayoutSpacing {
	return CollectionLayoutSpacingClass.FixedSpacing(fixedSpacing)
}

func (cc _CollectionLayoutSpacingClass) Alloc() CollectionLayoutSpacing {
	rv := objc.Call[CollectionLayoutSpacing](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CollectionLayoutSpacingClass) New() CollectionLayoutSpacing {
	rv := objc.Call[CollectionLayoutSpacing](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSpacing() CollectionLayoutSpacing {
	return CollectionLayoutSpacingClass.New()
}

func (c_ CollectionLayoutSpacing) Init() CollectionLayoutSpacing {
	rv := objc.Call[CollectionLayoutSpacing](c_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the space is flexible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutspacing/3199106-isflexiblespacing?language=objc
func (c_ CollectionLayoutSpacing) IsFlexibleSpacing() bool {
	rv := objc.Call[bool](c_, objc.Sel("isFlexibleSpacing"))
	return rv
}

// The floating-point value of the space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutspacing/3199107-spacing?language=objc
func (c_ CollectionLayoutSpacing) Spacing() float64 {
	rv := objc.Call[float64](c_, objc.Sel("spacing"))
	return rv
}

// A Boolean value that indicates whether the space is fixed to a specific number of points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutspacing/3199105-isfixedspacing?language=objc
func (c_ CollectionLayoutSpacing) IsFixedSpacing() bool {
	rv := objc.Call[bool](c_, objc.Sel("isFixedSpacing"))
	return rv
}
