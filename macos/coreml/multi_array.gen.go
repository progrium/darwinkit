// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MultiArray] class.
var MultiArrayClass = _MultiArrayClass{objc.GetClass("MLMultiArray")}

type _MultiArrayClass struct {
	objc.Class
}

// An interface definition for the [MultiArray] class.
type IMultiArray interface {
	objc.IObject
	ObjectAtIndexedSubscript(idx int) foundation.Number
	ObjectForKeyedSubscript(key []foundation.INumber) foundation.Number
	SetObjectAtIndexedSubscript(obj foundation.INumber, idx int)
	Shape() []foundation.Number
	PixelBuffer() corevideo.PixelBufferRef
	Count() int
	DataType() MultiArrayDataType
	Strides() []foundation.Number
}

// A machine learning collection type that stores numeric values in an array with multiple dimensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarray?language=objc
type MultiArray struct {
	objc.Object
}

func MultiArrayFrom(ptr unsafe.Pointer) MultiArray {
	return MultiArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MultiArrayClass) MultiArrayByConcatenatingMultiArraysAlongAxisDataType(multiArrays []IMultiArray, axis int, dataType MultiArrayDataType) MultiArray {
	rv := objc.Call[MultiArray](mc, objc.Sel("multiArrayByConcatenatingMultiArrays:alongAxis:dataType:"), multiArrays, axis, dataType)
	return rv
}

// Merges an array of multiarrays into one multiarray along an axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarray/3563973-multiarraybyconcatenatingmultiar?language=objc
func MultiArray_MultiArrayByConcatenatingMultiArraysAlongAxisDataType(multiArrays []IMultiArray, axis int, dataType MultiArrayDataType) MultiArray {
	return MultiArrayClass.MultiArrayByConcatenatingMultiArraysAlongAxisDataType(multiArrays, axis, dataType)
}

func (m_ MultiArray) InitWithShapeDataTypeError(shape []foundation.INumber, dataType MultiArrayDataType, error foundation.IError) MultiArray {
	rv := objc.Call[MultiArray](m_, objc.Sel("initWithShape:dataType:error:"), shape, dataType, objc.Ptr(error))
	return rv
}

// Creates a multidimensional array with a shape and type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarray/2879232-initwithshape?language=objc
func NewMultiArrayWithShapeDataTypeError(shape []foundation.INumber, dataType MultiArrayDataType, error foundation.IError) MultiArray {
	instance := MultiArrayClass.Alloc().InitWithShapeDataTypeError(shape, dataType, error)
	instance.Autorelease()
	return instance
}

func (m_ MultiArray) InitWithDataPointerShapeDataTypeStridesDeallocatorError(dataPointer unsafe.Pointer, shape []foundation.INumber, dataType MultiArrayDataType, strides []foundation.INumber, deallocator func(bytes unsafe.Pointer), error foundation.IError) MultiArray {
	rv := objc.Call[MultiArray](m_, objc.Sel("initWithDataPointer:shape:dataType:strides:deallocator:error:"), dataPointer, shape, dataType, strides, deallocator, objc.Ptr(error))
	return rv
}

// Creates a multiarray from a data pointer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarray/2881219-initwithdatapointer?language=objc
func NewMultiArrayWithDataPointerShapeDataTypeStridesDeallocatorError(dataPointer unsafe.Pointer, shape []foundation.INumber, dataType MultiArrayDataType, strides []foundation.INumber, deallocator func(bytes unsafe.Pointer), error foundation.IError) MultiArray {
	instance := MultiArrayClass.Alloc().InitWithDataPointerShapeDataTypeStridesDeallocatorError(dataPointer, shape, dataType, strides, deallocator, error)
	instance.Autorelease()
	return instance
}

func (m_ MultiArray) InitWithPixelBufferShape(pixelBuffer corevideo.PixelBufferRef, shape []foundation.INumber) MultiArray {
	rv := objc.Call[MultiArray](m_, objc.Sel("initWithPixelBuffer:shape:"), pixelBuffer, shape)
	return rv
}

// Creates a multiarray sharing the surface of a pixel buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarray/3882834-initwithpixelbuffer?language=objc
func NewMultiArrayWithPixelBufferShape(pixelBuffer corevideo.PixelBufferRef, shape []foundation.INumber) MultiArray {
	instance := MultiArrayClass.Alloc().InitWithPixelBufferShape(pixelBuffer, shape)
	instance.Autorelease()
	return instance
}

func (mc _MultiArrayClass) Alloc() MultiArray {
	rv := objc.Call[MultiArray](mc, objc.Sel("alloc"))
	return rv
}

func MultiArray_Alloc() MultiArray {
	return MultiArrayClass.Alloc()
}

func (mc _MultiArrayClass) New() MultiArray {
	rv := objc.Call[MultiArray](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMultiArray() MultiArray {
	return MultiArrayClass.New()
}

func (m_ MultiArray) Init() MultiArray {
	rv := objc.Call[MultiArray](m_, objc.Sel("init"))
	return rv
}

// Accesses the multiarray by using a linear offset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarray/2879228-objectatindexedsubscript?language=objc
func (m_ MultiArray) ObjectAtIndexedSubscript(idx int) foundation.Number {
	rv := objc.Call[foundation.Number](m_, objc.Sel("objectAtIndexedSubscript:"), idx)
	return rv
}

// Accesses the multiarray by using a number array that has an element for each dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarray/2879231-objectforkeyedsubscript?language=objc
func (m_ MultiArray) ObjectForKeyedSubscript(key []foundation.INumber) foundation.Number {
	rv := objc.Call[foundation.Number](m_, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}

// Assigns a number to the multiarray’s element at the location that the linear offset defines. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarray/2879226-setobject?language=objc
func (m_ MultiArray) SetObjectAtIndexedSubscript(obj foundation.INumber, idx int) {
	objc.Call[objc.Void](m_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(obj), idx)
}

// The multiarray’s multidimensional shape as a number array in which each element’s value is the size of the corresponding dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarray/2879229-shape?language=objc
func (m_ MultiArray) Shape() []foundation.Number {
	rv := objc.Call[[]foundation.Number](m_, objc.Sel("shape"))
	return rv
}

// A reference to the multiarray’s underlying pixel buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarray/3882835-pixelbuffer?language=objc
func (m_ MultiArray) PixelBuffer() corevideo.PixelBufferRef {
	rv := objc.Call[corevideo.PixelBufferRef](m_, objc.Sel("pixelBuffer"))
	return rv
}

// The total number of elements in the multiarray. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarray/2879233-count?language=objc
func (m_ MultiArray) Count() int {
	rv := objc.Call[int](m_, objc.Sel("count"))
	return rv
}

// The underlying type of the multiarray. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarray/2879219-datatype?language=objc
func (m_ MultiArray) DataType() MultiArrayDataType {
	rv := objc.Call[MultiArrayDataType](m_, objc.Sel("dataType"))
	return rv
}

// A number array in which each element is the number of memory locations that span the length of the corresponding dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarray/2879222-strides?language=objc
func (m_ MultiArray) Strides() []foundation.Number {
	rv := objc.Call[[]foundation.Number](m_, objc.Sel("strides"))
	return rv
}
