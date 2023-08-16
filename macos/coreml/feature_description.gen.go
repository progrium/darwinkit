// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FeatureDescription] class.
var FeatureDescriptionClass = _FeatureDescriptionClass{objc.GetClass("MLFeatureDescription")}

type _FeatureDescriptionClass struct {
	objc.Class
}

// An interface definition for the [FeatureDescription] class.
type IFeatureDescription interface {
	objc.IObject
	IsAllowedValue(value IFeatureValue) bool
	DictionaryConstraint() DictionaryConstraint
	MultiArrayConstraint() MultiArrayConstraint
	Name() string
	SequenceConstraint() SequenceConstraint
	ImageConstraint() ImageConstraint
	IsOptional() bool
	Type() FeatureType
}

// The name, type, and constraints of an input or output feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription?language=objc
type FeatureDescription struct {
	objc.Object
}

func FeatureDescriptionFrom(ptr unsafe.Pointer) FeatureDescription {
	return FeatureDescription{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FeatureDescriptionClass) Alloc() FeatureDescription {
	rv := objc.Call[FeatureDescription](fc, objc.Sel("alloc"))
	return rv
}

func FeatureDescription_Alloc() FeatureDescription {
	return FeatureDescriptionClass.Alloc()
}

func (fc _FeatureDescriptionClass) New() FeatureDescription {
	rv := objc.Call[FeatureDescription](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFeatureDescription() FeatureDescription {
	return FeatureDescriptionClass.New()
}

func (f_ FeatureDescription) Init() FeatureDescription {
	rv := objc.Call[FeatureDescription](f_, objc.Sel("init"))
	return rv
}

// Checks whether the model will accept an input feature value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/2879363-isallowedvalue?language=objc
func (f_ FeatureDescription) IsAllowedValue(value IFeatureValue) bool {
	rv := objc.Call[bool](f_, objc.Sel("isAllowedValue:"), objc.Ptr(value))
	return rv
}

// The constraint for a dictionary feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/2921263-dictionaryconstraint?language=objc
func (f_ FeatureDescription) DictionaryConstraint() DictionaryConstraint {
	rv := objc.Call[DictionaryConstraint](f_, objc.Sel("dictionaryConstraint"))
	return rv
}

// The constraints on a multidimensional array feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/2921264-multiarrayconstraint?language=objc
func (f_ FeatureDescription) MultiArrayConstraint() MultiArrayConstraint {
	rv := objc.Call[MultiArrayConstraint](f_, objc.Sel("multiArrayConstraint"))
	return rv
}

// The name of this feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/2879359-name?language=objc
func (f_ FeatureDescription) Name() string {
	rv := objc.Call[string](f_, objc.Sel("name"))
	return rv
}

// The constraints for a sequence feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/2962858-sequenceconstraint?language=objc
func (f_ FeatureDescription) SequenceConstraint() SequenceConstraint {
	rv := objc.Call[SequenceConstraint](f_, objc.Sel("sequenceConstraint"))
	return rv
}

// The size and format constraints for an image feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/2873067-imageconstraint?language=objc
func (f_ FeatureDescription) ImageConstraint() ImageConstraint {
	rv := objc.Call[ImageConstraint](f_, objc.Sel("imageConstraint"))
	return rv
}

// A Boolean value that indicates whether this feature is optional. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/2879384-optional?language=objc
func (f_ FeatureDescription) IsOptional() bool {
	rv := objc.Call[bool](f_, objc.Sel("isOptional"))
	return rv
}

// The type of this feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/2879374-type?language=objc
func (f_ FeatureDescription) Type() FeatureType {
	rv := objc.Call[FeatureType](f_, objc.Sel("type"))
	return rv
}
