// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Vector] class.
var VectorClass = _VectorClass{objc.GetClass("CIVector")}

type _VectorClass struct {
	objc.Class
}

// An interface definition for the [Vector] class.
type IVector interface {
	objc.IObject
	ValueAtIndex(index uint) float64
	X() float64
	SetX(value float64)
	Y() float64
	SetY(value float64)
	CGRectValue() coregraphics.Rect
	Count() uint
	CGPointValue() coregraphics.Point
	CGAffineTransformValue() coregraphics.AffineTransform
	Z() float64
	SetZ(value float64)
	StringRepresentation() string
	W() float64
	SetW(value float64)
}

// A container for coordinate values, direction vectors, matrices, and other non-scalar values, typically used in Core Image for filter parameters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector?language=objc
type Vector struct {
	objc.Object
}

func VectorFrom(ptr unsafe.Pointer) Vector {
	return Vector{
		Object: objc.ObjectFrom(ptr),
	}
}

func (v_ Vector) InitWithCGRect(r coregraphics.Rect) Vector {
	rv := objc.Call[Vector](v_, objc.Sel("initWithCGRect:"), r)
	return rv
}

// Initializes a vector that is initialized with values provided by a CGRect structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1437644-initwithcgrect?language=objc
func NewVectorWithCGRect(r coregraphics.Rect) Vector {
	instance := VectorClass.Alloc().InitWithCGRect(r)
	instance.Autorelease()
	return instance
}

func (v_ Vector) InitWithValuesCount(values *float64, count uint) Vector {
	rv := objc.Call[Vector](v_, objc.Sel("initWithValues:count:"), values, count)
	return rv
}

// Initializes a vector with the provided values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1437849-initwithvalues?language=objc
func NewVectorWithValuesCount(values *float64, count uint) Vector {
	instance := VectorClass.Alloc().InitWithValuesCount(values, count)
	instance.Autorelease()
	return instance
}

func (vc _VectorClass) VectorWithCGAffineTransform(t coregraphics.AffineTransform) Vector {
	rv := objc.Call[Vector](vc, objc.Sel("vectorWithCGAffineTransform:"), t)
	return rv
}

// Creates and returns a vector that is initialized with values provided by a CGAffineTransform structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1564090-vectorwithcgaffinetransform?language=objc
func Vector_VectorWithCGAffineTransform(t coregraphics.AffineTransform) Vector {
	return VectorClass.VectorWithCGAffineTransform(t)
}

func (v_ Vector) InitWithString(representation string) Vector {
	rv := objc.Call[Vector](v_, objc.Sel("initWithString:"), representation)
	return rv
}

// Initializes a vector with values provided in a string representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1437938-initwithstring?language=objc
func NewVectorWithString(representation string) Vector {
	instance := VectorClass.Alloc().InitWithString(representation)
	instance.Autorelease()
	return instance
}

func (v_ Vector) InitWithXY(x float64, y float64) Vector {
	rv := objc.Call[Vector](v_, objc.Sel("initWithX:Y:"), x, y)
	return rv
}

// Initializes the first two positions of a vector with the provided values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1437865-initwithx?language=objc
func NewVectorWithXY(x float64, y float64) Vector {
	instance := VectorClass.Alloc().InitWithXY(x, y)
	instance.Autorelease()
	return instance
}

func (vc _VectorClass) VectorWithX(x float64) Vector {
	rv := objc.Call[Vector](vc, objc.Sel("vectorWithX:"), x)
	return rv
}

// Creates and returns a vector that is initialized with one value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1564092-vectorwithx?language=objc
func Vector_VectorWithX(x float64) Vector {
	return VectorClass.VectorWithX(x)
}

func (v_ Vector) InitWithCGAffineTransform(r coregraphics.AffineTransform) Vector {
	rv := objc.Call[Vector](v_, objc.Sel("initWithCGAffineTransform:"), r)
	return rv
}

// Initializes a vector that is initialized with values provided by a CGAffineTransform structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1438102-initwithcgaffinetransform?language=objc
func NewVectorWithCGAffineTransform(r coregraphics.AffineTransform) Vector {
	instance := VectorClass.Alloc().InitWithCGAffineTransform(r)
	instance.Autorelease()
	return instance
}

func (vc _VectorClass) VectorWithCGRect(r coregraphics.Rect) Vector {
	rv := objc.Call[Vector](vc, objc.Sel("vectorWithCGRect:"), r)
	return rv
}

// Creates and returns a vector that is initialized with values provided by a CGRect structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1564085-vectorwithcgrect?language=objc
func Vector_VectorWithCGRect(r coregraphics.Rect) Vector {
	return VectorClass.VectorWithCGRect(r)
}

func (v_ Vector) InitWithCGPoint(p coregraphics.Point) Vector {
	rv := objc.Call[Vector](v_, objc.Sel("initWithCGPoint:"), p)
	return rv
}

// Initializes a vector that is initialized with values provided by a CGPoint structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1438133-initwithcgpoint?language=objc
func NewVectorWithCGPoint(p coregraphics.Point) Vector {
	instance := VectorClass.Alloc().InitWithCGPoint(p)
	instance.Autorelease()
	return instance
}

func (vc _VectorClass) VectorWithValuesCount(values *float64, count uint) Vector {
	rv := objc.Call[Vector](vc, objc.Sel("vectorWithValues:count:"), values, count)
	return rv
}

// Creates and returns a vector that is initialized with the specified values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1564088-vectorwithvalues?language=objc
func Vector_VectorWithValuesCount(values *float64, count uint) Vector {
	return VectorClass.VectorWithValuesCount(values, count)
}

func (vc _VectorClass) VectorWithString(representation string) Vector {
	rv := objc.Call[Vector](vc, objc.Sel("vectorWithString:"), representation)
	return rv
}

// Creates and returns a vector that is initialized with values provided in a string representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1564093-vectorwithstring?language=objc
func Vector_VectorWithString(representation string) Vector {
	return VectorClass.VectorWithString(representation)
}

func (vc _VectorClass) VectorWithCGPoint(p coregraphics.Point) Vector {
	rv := objc.Call[Vector](vc, objc.Sel("vectorWithCGPoint:"), p)
	return rv
}

// Creates and returns a vector that is initialized with values provided by a CGPoint structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1564086-vectorwithcgpoint?language=objc
func Vector_VectorWithCGPoint(p coregraphics.Point) Vector {
	return VectorClass.VectorWithCGPoint(p)
}

func (vc _VectorClass) Alloc() Vector {
	rv := objc.Call[Vector](vc, objc.Sel("alloc"))
	return rv
}

func Vector_Alloc() Vector {
	return VectorClass.Alloc()
}

func (vc _VectorClass) New() Vector {
	rv := objc.Call[Vector](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVector() Vector {
	return VectorClass.New()
}

func (v_ Vector) Init() Vector {
	rv := objc.Call[Vector](v_, objc.Sel("init"))
	return rv
}

// Returns a value from a specific position in the vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1438207-valueatindex?language=objc
func (v_ Vector) ValueAtIndex(index uint) float64 {
	rv := objc.Call[float64](v_, objc.Sel("valueAtIndex:"), index)
	return rv
}

// The value located in the first position in the vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1437738-x?language=objc
func (v_ Vector) X() float64 {
	rv := objc.Call[float64](v_, objc.Sel("X"))
	return rv
}

// The value located in the first position in the vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1437738-x?language=objc
func (v_ Vector) SetX(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setX:"), value)
}

// The value located in the second position in the vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1437843-y?language=objc
func (v_ Vector) Y() float64 {
	rv := objc.Call[float64](v_, objc.Sel("Y"))
	return rv
}

// The value located in the second position in the vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1437843-y?language=objc
func (v_ Vector) SetY(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setY:"), value)
}

// The values in the vector as a Core Graphics rectangle structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1438108-cgrectvalue?language=objc
func (v_ Vector) CGRectValue() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](v_, objc.Sel("CGRectValue"))
	return rv
}

// The number of items in the vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1438197-count?language=objc
func (v_ Vector) Count() uint {
	rv := objc.Call[uint](v_, objc.Sel("count"))
	return rv
}

// The values in the vector as a Core Graphics point structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1437672-cgpointvalue?language=objc
func (v_ Vector) CGPointValue() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](v_, objc.Sel("CGPointValue"))
	return rv
}

// The values in the vector represented as an affine transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1438249-cgaffinetransformvalue?language=objc
func (v_ Vector) CGAffineTransformValue() coregraphics.AffineTransform {
	rv := objc.Call[coregraphics.AffineTransform](v_, objc.Sel("CGAffineTransformValue"))
	return rv
}

// The value located in the third position in the vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1437627-z?language=objc
func (v_ Vector) Z() float64 {
	rv := objc.Call[float64](v_, objc.Sel("Z"))
	return rv
}

// The value located in the third position in the vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1437627-z?language=objc
func (v_ Vector) SetZ(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setZ:"), value)
}

// The string representation of the vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1437752-stringrepresentation?language=objc
func (v_ Vector) StringRepresentation() string {
	rv := objc.Call[string](v_, objc.Sel("stringRepresentation"))
	return rv
}

// The value located in the fourth position in the vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1438058-w?language=objc
func (v_ Vector) W() float64 {
	rv := objc.Call[float64](v_, objc.Sel("W"))
	return rv
}

// The value located in the fourth position in the vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civector/1438058-w?language=objc
func (v_ Vector) SetW(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setW:"), value)
}
