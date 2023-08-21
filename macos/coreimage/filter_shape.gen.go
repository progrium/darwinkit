// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FilterShape] class.
var FilterShapeClass = _FilterShapeClass{objc.GetClass("CIFilterShape")}

type _FilterShapeClass struct {
	objc.Class
}

// An interface definition for the [FilterShape] class.
type IFilterShape interface {
	objc.IObject
	IntersectWith(s2 IFilterShape) FilterShape
	UnionWithRect(r coregraphics.Rect) FilterShape
	UnionWith(s2 IFilterShape) FilterShape
	InsetByXY(dx int, dy int) FilterShape
	IntersectWithRect(r coregraphics.Rect) FilterShape
	TransformByInterior(m coregraphics.AffineTransform, flag bool) FilterShape
	Extent() coregraphics.Rect
}

// A description of the bounding shape of a filter and the domain of definition for a filter operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltershape?language=objc
type FilterShape struct {
	objc.Object
}

func FilterShapeFrom(ptr unsafe.Pointer) FilterShape {
	return FilterShape{
		Object: objc.ObjectFrom(ptr),
	}
}

func (f_ FilterShape) InitWithRect(r coregraphics.Rect) FilterShape {
	rv := objc.Call[FilterShape](f_, objc.Sel("initWithRect:"), r)
	return rv
}

// Initializes a filter shape object with a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltershape/1437921-initwithrect?language=objc
func NewFilterShapeWithRect(r coregraphics.Rect) FilterShape {
	instance := FilterShapeClass.Alloc().InitWithRect(r)
	instance.Autorelease()
	return instance
}

func (fc _FilterShapeClass) ShapeWithRect(r coregraphics.Rect) FilterShape {
	rv := objc.Call[FilterShape](fc, objc.Sel("shapeWithRect:"), r)
	return rv
}

// Creates a filter shape object and initializes it with a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltershape/1562074-shapewithrect?language=objc
func FilterShape_ShapeWithRect(r coregraphics.Rect) FilterShape {
	return FilterShapeClass.ShapeWithRect(r)
}

func (fc _FilterShapeClass) Alloc() FilterShape {
	rv := objc.Call[FilterShape](fc, objc.Sel("alloc"))
	return rv
}

func FilterShape_Alloc() FilterShape {
	return FilterShapeClass.Alloc()
}

func (fc _FilterShapeClass) New() FilterShape {
	rv := objc.Call[FilterShape](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFilterShape() FilterShape {
	return FilterShapeClass.New()
}

func (f_ FilterShape) Init() FilterShape {
	rv := objc.Call[FilterShape](f_, objc.Sel("init"))
	return rv
}

// Creates a filter shape object that represents the intersection of the current filter shape and the specified filter shape object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltershape/1437881-intersectwith?language=objc
func (f_ FilterShape) IntersectWith(s2 IFilterShape) FilterShape {
	rv := objc.Call[FilterShape](f_, objc.Sel("intersectWith:"), objc.Ptr(s2))
	return rv
}

// Creates a filter shape that results from the union of the current filter shape and a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltershape/1437601-unionwithrect?language=objc
func (f_ FilterShape) UnionWithRect(r coregraphics.Rect) FilterShape {
	rv := objc.Call[FilterShape](f_, objc.Sel("unionWithRect:"), r)
	return rv
}

// Creates a filter shape that results from the union of the current filter shape and another filter shape object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltershape/1438227-unionwith?language=objc
func (f_ FilterShape) UnionWith(s2 IFilterShape) FilterShape {
	rv := objc.Call[FilterShape](f_, objc.Sel("unionWith:"), objc.Ptr(s2))
	return rv
}

// Modifies a filter shape object so that it is inset by the specified x and y values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltershape/1437987-insetbyx?language=objc
func (f_ FilterShape) InsetByXY(dx int, dy int) FilterShape {
	rv := objc.Call[FilterShape](f_, objc.Sel("insetByX:Y:"), dx, dy)
	return rv
}

// Creates a filter shape that represents the intersection of the current filter shape and a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltershape/1437806-intersectwithrect?language=objc
func (f_ FilterShape) IntersectWithRect(r coregraphics.Rect) FilterShape {
	rv := objc.Call[FilterShape](f_, objc.Sel("intersectWithRect:"), r)
	return rv
}

// Creates a filter shape that results from applying a transform to the current filter shape. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltershape/1437808-transformby?language=objc
func (f_ FilterShape) TransformByInterior(m coregraphics.AffineTransform, flag bool) FilterShape {
	rv := objc.Call[FilterShape](f_, objc.Sel("transformBy:interior:"), m, flag)
	return rv
}

// The extent of the filter shape. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltershape/1438022-extent?language=objc
func (f_ FilterShape) Extent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](f_, objc.Sel("extent"))
	return rv
}
