// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Vector] class.
var VectorClass = _VectorClass{objc.GetClass("VNVector")}

type _VectorClass struct {
	objc.Class
}

// An interface definition for the [Vector] class.
type IVector interface {
	objc.IObject
	X() float64
	Y() float64
	R() float64
	Length() float64
	Theta() float64
	SquaredLength() float64
}

// An immutable, two-dimensional vector represented by its x-axis and y-axis projections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector?language=objc
type Vector struct {
	objc.Object
}

func VectorFrom(ptr unsafe.Pointer) Vector {
	return Vector{
		Object: objc.ObjectFrom(ptr),
	}
}

func (v_ Vector) InitWithRTheta(r float64, theta float64) Vector {
	rv := objc.Call[Vector](v_, objc.Sel("initWithR:theta:"), r, theta)
	return rv
}

// Creates a new vector in polar coordinate space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548339-initwithr?language=objc
func NewVectorWithRTheta(r float64, theta float64) Vector {
	instance := VectorClass.Alloc().InitWithRTheta(r, theta)
	instance.Autorelease()
	return instance
}

func (v_ Vector) InitWithXComponentYComponent(x float64, y float64) Vector {
	rv := objc.Call[Vector](v_, objc.Sel("initWithXComponent:yComponent:"), x, y)
	return rv
}

// Creates a new vector in Cartesian coordinate space, based on its x-axis and y-axis projections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548341-initwithxcomponent?language=objc
func NewVectorWithXComponentYComponent(x float64, y float64) Vector {
	instance := VectorClass.Alloc().InitWithXComponentYComponent(x, y)
	instance.Autorelease()
	return instance
}

func (v_ Vector) InitWithVectorHeadTail(head IPoint, tail IPoint) Vector {
	rv := objc.Call[Vector](v_, objc.Sel("initWithVectorHead:tail:"), objc.Ptr(head), objc.Ptr(tail))
	return rv
}

// Creates a new vector in Cartesian coordinate space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548340-initwithvectorhead?language=objc
func NewVectorWithVectorHeadTail(head IPoint, tail IPoint) Vector {
	instance := VectorClass.Alloc().InitWithVectorHeadTail(head, tail)
	instance.Autorelease()
	return instance
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

// Creates a new vector by subtracting the first vector from the second vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548349-vectorbysubtractingvector?language=objc
func (vc _VectorClass) VectorBySubtractingVectorFromVector(v1 IVector, v2 IVector) Vector {
	rv := objc.Call[Vector](vc, objc.Sel("vectorBySubtractingVector:fromVector:"), objc.Ptr(v1), objc.Ptr(v2))
	return rv
}

// Creates a new vector by subtracting the first vector from the second vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548349-vectorbysubtractingvector?language=objc
func Vector_VectorBySubtractingVectorFromVector(v1 IVector, v2 IVector) Vector {
	return VectorClass.VectorBySubtractingVectorFromVector(v1, v2)
}

// Caclulates the dot product of two vectors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548338-dotproductofvector?language=objc
func (vc _VectorClass) DotProductOfVectorVector(v1 IVector, v2 IVector) float64 {
	rv := objc.Call[float64](vc, objc.Sel("dotProductOfVector:vector:"), objc.Ptr(v1), objc.Ptr(v2))
	return rv
}

// Caclulates the dot product of two vectors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548338-dotproductofvector?language=objc
func Vector_DotProductOfVectorVector(v1 IVector, v2 IVector) float64 {
	return VectorClass.DotProductOfVectorVector(v1, v2)
}

// Calculates a vector that’s normalized by preserving its direction, so that the vector length equals 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548346-unitvectorforvector?language=objc
func (vc _VectorClass) UnitVectorForVector(vector IVector) Vector {
	rv := objc.Call[Vector](vc, objc.Sel("unitVectorForVector:"), objc.Ptr(vector))
	return rv
}

// Calculates a vector that’s normalized by preserving its direction, so that the vector length equals 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548346-unitvectorforvector?language=objc
func Vector_UnitVectorForVector(vector IVector) Vector {
	return VectorClass.UnitVectorForVector(vector)
}

// Creates a new vector by adding the specified vectors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548347-vectorbyaddingvector?language=objc
func (vc _VectorClass) VectorByAddingVectorToVector(v1 IVector, v2 IVector) Vector {
	rv := objc.Call[Vector](vc, objc.Sel("vectorByAddingVector:toVector:"), objc.Ptr(v1), objc.Ptr(v2))
	return rv
}

// Creates a new vector by adding the specified vectors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548347-vectorbyaddingvector?language=objc
func Vector_VectorByAddingVectorToVector(v1 IVector, v2 IVector) Vector {
	return VectorClass.VectorByAddingVectorToVector(v1, v2)
}

// Creates a new vector by multiplying the specified vector’s x-axis and y-axis projections by the scalar value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548348-vectorbymultiplyingvector?language=objc
func (vc _VectorClass) VectorByMultiplyingVectorByScalar(vector IVector, scalar float64) Vector {
	rv := objc.Call[Vector](vc, objc.Sel("vectorByMultiplyingVector:byScalar:"), objc.Ptr(vector), scalar)
	return rv
}

// Creates a new vector by multiplying the specified vector’s x-axis and y-axis projections by the scalar value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548348-vectorbymultiplyingvector?language=objc
func Vector_VectorByMultiplyingVectorByScalar(vector IVector, scalar float64) Vector {
	return VectorClass.VectorByMultiplyingVectorByScalar(vector, scalar)
}

// A signed projection that indicates the vector’s direction on the x-axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548350-x?language=objc
func (v_ Vector) X() float64 {
	rv := objc.Call[float64](v_, objc.Sel("x"))
	return rv
}

// A signed projection that indicates the vector’s direction on the y-axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548351-y?language=objc
func (v_ Vector) Y() float64 {
	rv := objc.Call[float64](v_, objc.Sel("y"))
	return rv
}

// A vector object with zero length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548352-zerovector?language=objc
func (vc _VectorClass) ZeroVector() Vector {
	rv := objc.Call[Vector](vc, objc.Sel("zeroVector"))
	return rv
}

// A vector object with zero length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548352-zerovector?language=objc
func Vector_ZeroVector() Vector {
	return VectorClass.ZeroVector()
}

// The radius, absolute value, or length of the vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548343-r?language=objc
func (v_ Vector) R() float64 {
	rv := objc.Call[float64](v_, objc.Sel("r"))
	return rv
}

// The length, or absolute value, of the vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548342-length?language=objc
func (v_ Vector) Length() float64 {
	rv := objc.Call[float64](v_, objc.Sel("length"))
	return rv
}

// The angle between the vector direction and the positive direction of the x-axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548345-theta?language=objc
func (v_ Vector) Theta() float64 {
	rv := objc.Call[float64](v_, objc.Sel("theta"))
	return rv
}

// The squared length of the vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/3548344-squaredlength?language=objc
func (v_ Vector) SquaredLength() float64 {
	rv := objc.Call[float64](v_, objc.Sel("squaredLength"))
	return rv
}
