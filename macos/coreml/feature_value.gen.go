// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/imageio"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FeatureValue] class.
var FeatureValueClass = _FeatureValueClass{objc.GetClass("MLFeatureValue")}

type _FeatureValueClass struct {
	objc.Class
}

// An interface definition for the [FeatureValue] class.
type IFeatureValue interface {
	objc.IObject
	IsEqualToFeatureValue(value IFeatureValue) bool
	Int64Value() int64
	StringValue() string
	MultiArrayValue() MultiArray
	DoubleValue() float64
	SequenceValue() Sequence
	ImageBufferValue() corevideo.PixelBufferRef
	IsUndefined() bool
	Type() FeatureType
	DictionaryValue() foundation.Dictionary
}

// A generic wrapper around an underlying value and the value’s type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue?language=objc
type FeatureValue struct {
	objc.Object
}

func FeatureValueFrom(ptr unsafe.Pointer) FeatureValue {
	return FeatureValue{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FeatureValueClass) FeatureValueWithString(value string) FeatureValue {
	rv := objc.Call[FeatureValue](fc, objc.Sel("featureValueWithString:"), value)
	return rv
}

// Creates a feature value that contains a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879343-featurevaluewithstring?language=objc
func FeatureValue_FeatureValueWithString(value string) FeatureValue {
	return FeatureValueClass.FeatureValueWithString(value)
}

func (fc _FeatureValueClass) FeatureValueWithSequence(sequence ISequence) FeatureValue {
	rv := objc.Call[FeatureValue](fc, objc.Sel("featureValueWithSequence:"), objc.Ptr(sequence))
	return rv
}

// Creates a feature value that contains a sequence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2962860-featurevaluewithsequence?language=objc
func FeatureValue_FeatureValueWithSequence(sequence ISequence) FeatureValue {
	return FeatureValueClass.FeatureValueWithSequence(sequence)
}

func (fc _FeatureValueClass) FeatureValueWithDictionaryError(value foundation.Dictionary, error foundation.IError) FeatureValue {
	rv := objc.Call[FeatureValue](fc, objc.Sel("featureValueWithDictionary:error:"), value, objc.Ptr(error))
	return rv
}

// Creates a feature value that contains a dictionary of numbers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879393-featurevaluewithdictionary?language=objc
func FeatureValue_FeatureValueWithDictionaryError(value foundation.Dictionary, error foundation.IError) FeatureValue {
	return FeatureValueClass.FeatureValueWithDictionaryError(value, error)
}

func (fc _FeatureValueClass) FeatureValueWithMultiArray(value IMultiArray) FeatureValue {
	rv := objc.Call[FeatureValue](fc, objc.Sel("featureValueWithMultiArray:"), objc.Ptr(value))
	return rv
}

// Creates a feature value that contains a multidimensional array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879356-featurevaluewithmultiarray?language=objc
func FeatureValue_FeatureValueWithMultiArray(value IMultiArray) FeatureValue {
	return FeatureValueClass.FeatureValueWithMultiArray(value)
}

func (fc _FeatureValueClass) FeatureValueWithCGImageOrientationConstraintOptionsError(cgImage coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, constraint IImageConstraint, options map[FeatureValueImageOption]objc.IObject, error foundation.IError) FeatureValue {
	rv := objc.Call[FeatureValue](fc, objc.Sel("featureValueWithCGImage:orientation:constraint:options:error:"), cgImage, orientation, objc.Ptr(constraint), options, objc.Ptr(error))
	return rv
}

// Creates a feature value that contains an image defined by a core graphics image, an orientation, and a constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/3362522-featurevaluewithcgimage?language=objc
func FeatureValue_FeatureValueWithCGImageOrientationConstraintOptionsError(cgImage coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, constraint IImageConstraint, options map[FeatureValueImageOption]objc.IObject, error foundation.IError) FeatureValue {
	return FeatureValueClass.FeatureValueWithCGImageOrientationConstraintOptionsError(cgImage, orientation, constraint, options, error)
}

func (fc _FeatureValueClass) FeatureValueWithImageAtURLConstraintOptionsError(url foundation.IURL, constraint IImageConstraint, options map[FeatureValueImageOption]objc.IObject, error foundation.IError) FeatureValue {
	rv := objc.Call[FeatureValue](fc, objc.Sel("featureValueWithImageAtURL:constraint:options:error:"), objc.Ptr(url), objc.Ptr(constraint), options, objc.Ptr(error))
	return rv
}

// Creates a feature value that contains an image defined by an image URL and a constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/3200162-featurevaluewithimageaturl?language=objc
func FeatureValue_FeatureValueWithImageAtURLConstraintOptionsError(url foundation.IURL, constraint IImageConstraint, options map[FeatureValueImageOption]objc.IObject, error foundation.IError) FeatureValue {
	return FeatureValueClass.FeatureValueWithImageAtURLConstraintOptionsError(url, constraint, options, error)
}

func (fc _FeatureValueClass) FeatureValueWithInt64(value int64) FeatureValue {
	rv := objc.Call[FeatureValue](fc, objc.Sel("featureValueWithInt64:"), value)
	return rv
}

// Creates a feature value that contains an integer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879396-featurevaluewithint64?language=objc
func FeatureValue_FeatureValueWithInt64(value int64) FeatureValue {
	return FeatureValueClass.FeatureValueWithInt64(value)
}

func (fc _FeatureValueClass) FeatureValueWithPixelBuffer(value corevideo.PixelBufferRef) FeatureValue {
	rv := objc.Call[FeatureValue](fc, objc.Sel("featureValueWithPixelBuffer:"), value)
	return rv
}

// Creates a feature value that contains an image from a pixel buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879346-featurevaluewithpixelbuffer?language=objc
func FeatureValue_FeatureValueWithPixelBuffer(value corevideo.PixelBufferRef) FeatureValue {
	return FeatureValueClass.FeatureValueWithPixelBuffer(value)
}

func (fc _FeatureValueClass) FeatureValueWithDouble(value float64) FeatureValue {
	rv := objc.Call[FeatureValue](fc, objc.Sel("featureValueWithDouble:"), value)
	return rv
}

// Creates a feature value that contains a double. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879398-featurevaluewithdouble?language=objc
func FeatureValue_FeatureValueWithDouble(value float64) FeatureValue {
	return FeatureValueClass.FeatureValueWithDouble(value)
}

func (fc _FeatureValueClass) UndefinedFeatureValueWithType(type_ FeatureType) FeatureValue {
	rv := objc.Call[FeatureValue](fc, objc.Sel("undefinedFeatureValueWithType:"), type_)
	return rv
}

// Creates a feature value with a type that represents an undefined or missing value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879362-undefinedfeaturevaluewithtype?language=objc
func FeatureValue_UndefinedFeatureValueWithType(type_ FeatureType) FeatureValue {
	return FeatureValueClass.UndefinedFeatureValueWithType(type_)
}

func (fc _FeatureValueClass) Alloc() FeatureValue {
	rv := objc.Call[FeatureValue](fc, objc.Sel("alloc"))
	return rv
}

func FeatureValue_Alloc() FeatureValue {
	return FeatureValueClass.Alloc()
}

func (fc _FeatureValueClass) New() FeatureValue {
	rv := objc.Call[FeatureValue](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFeatureValue() FeatureValue {
	return FeatureValueClass.New()
}

func (f_ FeatureValue) Init() FeatureValue {
	rv := objc.Call[FeatureValue](f_, objc.Sel("init"))
	return rv
}

// Returns a Boolean value that indicates whether a feature value is equal to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879399-isequaltofeaturevalue?language=objc
func (f_ FeatureValue) IsEqualToFeatureValue(value IFeatureValue) bool {
	rv := objc.Call[bool](f_, objc.Sel("isEqualToFeatureValue:"), objc.Ptr(value))
	return rv
}

// The underlying integer of the feature value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879348-int64value?language=objc
func (f_ FeatureValue) Int64Value() int64 {
	rv := objc.Call[int64](f_, objc.Sel("int64Value"))
	return rv
}

// The underlying string of the feature value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879349-stringvalue?language=objc
func (f_ FeatureValue) StringValue() string {
	rv := objc.Call[string](f_, objc.Sel("stringValue"))
	return rv
}

// The underlying multiarray of the feature value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879377-multiarrayvalue?language=objc
func (f_ FeatureValue) MultiArrayValue() MultiArray {
	rv := objc.Call[MultiArray](f_, objc.Sel("multiArrayValue"))
	return rv
}

// The underlying double of the feature value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879375-doublevalue?language=objc
func (f_ FeatureValue) DoubleValue() float64 {
	rv := objc.Call[float64](f_, objc.Sel("doubleValue"))
	return rv
}

// The underlying sequence of the feature value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2962861-sequencevalue?language=objc
func (f_ FeatureValue) SequenceValue() Sequence {
	rv := objc.Call[Sequence](f_, objc.Sel("sequenceValue"))
	return rv
}

// The underlying image of the feature value as a pixel buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879400-imagebuffervalue?language=objc
func (f_ FeatureValue) ImageBufferValue() corevideo.PixelBufferRef {
	rv := objc.Call[corevideo.PixelBufferRef](f_, objc.Sel("imageBufferValue"))
	return rv
}

// A Boolean value that indicates whether the feature value is undefined or missing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879392-undefined?language=objc
func (f_ FeatureValue) IsUndefined() bool {
	rv := objc.Call[bool](f_, objc.Sel("isUndefined"))
	return rv
}

// The type of the feature value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879368-type?language=objc
func (f_ FeatureValue) Type() FeatureType {
	rv := objc.Call[FeatureType](f_, objc.Sel("type"))
	return rv
}

// The underlying dictionary of the feature value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/2879387-dictionaryvalue?language=objc
func (f_ FeatureValue) DictionaryValue() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](f_, objc.Sel("dictionaryValue"))
	return rv
}
