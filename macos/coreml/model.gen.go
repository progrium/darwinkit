// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Model] class.
var ModelClass = _ModelClass{objc.GetClass("MLModel")}

type _ModelClass struct {
	objc.Class
}

// An interface definition for the [Model] class.
type IModel interface {
	objc.IObject
	ParameterValueForKeyError(key IParameterKey, error foundation.IError) objc.Object
	PredictionsFromBatchError(inputBatch PBatchProvider, error foundation.IError) BatchProviderWrapper
	PredictionsFromBatchObjectError(inputBatchObject objc.IObject, error foundation.IError) BatchProviderWrapper
	PredictionFromFeaturesError(input PFeatureProvider, error foundation.IError) FeatureProviderWrapper
	PredictionFromFeaturesObjectError(inputObject objc.IObject, error foundation.IError) FeatureProviderWrapper
	Configuration() ModelConfiguration
	ModelDescription() ModelDescription
}

// An encapsulation of all the details of your machine learning model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel?language=objc
type Model struct {
	objc.Object
}

func ModelFrom(ptr unsafe.Pointer) Model {
	return Model{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _ModelClass) ModelWithContentsOfURLError(url foundation.IURL, error foundation.IError) Model {
	rv := objc.Call[Model](mc, objc.Sel("modelWithContentsOfURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Creates a Core ML model instance from a compiled model file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/2880279-modelwithcontentsofurl?language=objc
func Model_ModelWithContentsOfURLError(url foundation.IURL, error foundation.IError) Model {
	return ModelClass.ModelWithContentsOfURLError(url, error)
}

func (mc _ModelClass) Alloc() Model {
	rv := objc.Call[Model](mc, objc.Sel("alloc"))
	return rv
}

func Model_Alloc() Model {
	return ModelClass.Alloc()
}

func (mc _ModelClass) New() Model {
	rv := objc.Call[Model](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewModel() Model {
	return ModelClass.New()
}

func (m_ Model) Init() Model {
	rv := objc.Call[Model](m_, objc.Sel("init"))
	return rv
}

// Returns a model parameter value for a key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/3362526-parametervalueforkey?language=objc
func (m_ Model) ParameterValueForKeyError(key IParameterKey, error foundation.IError) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("parameterValueForKey:error:"), objc.Ptr(key), objc.Ptr(error))
	return rv
}

// Creates a Core ML model instance asynchronously from a compiled model file, a custom configuration, and a completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/3600218-loadcontentsofurl?language=objc
func (mc _ModelClass) LoadContentsOfURLConfigurationCompletionHandler(url foundation.IURL, configuration IModelConfiguration, handler func(model Model, error foundation.Error)) {
	objc.Call[objc.Void](mc, objc.Sel("loadContentsOfURL:configuration:completionHandler:"), objc.Ptr(url), objc.Ptr(configuration), handler)
}

// Creates a Core ML model instance asynchronously from a compiled model file, a custom configuration, and a completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/3600218-loadcontentsofurl?language=objc
func Model_LoadContentsOfURLConfigurationCompletionHandler(url foundation.IURL, configuration IModelConfiguration, handler func(model Model, error foundation.Error)) {
	ModelClass.LoadContentsOfURLConfigurationCompletionHandler(url, configuration, handler)
}

// Generates predictions for each input feature provider within the batch provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/3088750-predictionsfrombatch?language=objc
func (m_ Model) PredictionsFromBatchError(inputBatch PBatchProvider, error foundation.IError) BatchProviderWrapper {
	po0 := objc.WrapAsProtocol("MLBatchProvider", inputBatch)
	rv := objc.Call[BatchProviderWrapper](m_, objc.Sel("predictionsFromBatch:error:"), po0, objc.Ptr(error))
	return rv
}

// Generates predictions for each input feature provider within the batch provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/3088750-predictionsfrombatch?language=objc
func (m_ Model) PredictionsFromBatchObjectError(inputBatchObject objc.IObject, error foundation.IError) BatchProviderWrapper {
	rv := objc.Call[BatchProviderWrapper](m_, objc.Sel("predictionsFromBatch:error:"), objc.Ptr(inputBatchObject), objc.Ptr(error))
	return rv
}

// Generates a prediction from the feature values within the input feature provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/2880280-predictionfromfeatures?language=objc
func (m_ Model) PredictionFromFeaturesError(input PFeatureProvider, error foundation.IError) FeatureProviderWrapper {
	po0 := objc.WrapAsProtocol("MLFeatureProvider", input)
	rv := objc.Call[FeatureProviderWrapper](m_, objc.Sel("predictionFromFeatures:error:"), po0, objc.Ptr(error))
	return rv
}

// Generates a prediction from the feature values within the input feature provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/2880280-predictionfromfeatures?language=objc
func (m_ Model) PredictionFromFeaturesObjectError(inputObject objc.IObject, error foundation.IError) FeatureProviderWrapper {
	rv := objc.Call[FeatureProviderWrapper](m_, objc.Sel("predictionFromFeatures:error:"), objc.Ptr(inputObject), objc.Ptr(error))
	return rv
}

// The configuration of the model set during initialization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/3022228-configuration?language=objc
func (m_ Model) Configuration() ModelConfiguration {
	rv := objc.Call[ModelConfiguration](m_, objc.Sel("configuration"))
	return rv
}

// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/2879179-modeldescription?language=objc
func (m_ Model) ModelDescription() ModelDescription {
	rv := objc.Call[ModelDescription](m_, objc.Sel("modelDescription"))
	return rv
}
