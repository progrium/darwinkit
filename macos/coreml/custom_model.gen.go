// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An interface that defines the behavior of a custom model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlcustommodel?language=objc
type PCustomModel interface {
	// optional
	InitWithModelDescriptionParameterDictionaryError(modelDescription ModelDescription, parameters map[string]objc.Object, error foundation.Error) objc.IObject
	HasInitWithModelDescriptionParameterDictionaryError() bool

	// optional
	PredictionsFromBatchOptionsError(inputBatch BatchProviderWrapper, options PredictionOptions, error foundation.Error) PBatchProvider
	HasPredictionsFromBatchOptionsError() bool

	// optional
	PredictionFromFeaturesOptionsError(input FeatureProviderWrapper, options PredictionOptions, error foundation.Error) PFeatureProvider
	HasPredictionFromFeaturesOptionsError() bool
}

// A concrete type wrapper for the [PCustomModel] protocol.
type CustomModelWrapper struct {
	objc.Object
}

func (c_ CustomModelWrapper) HasInitWithModelDescriptionParameterDictionaryError() bool {
	return c_.RespondsToSelector(objc.Sel("initWithModelDescription:parameterDictionary:error:"))
}

// Creates a custom model with the given description and parameters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlcustommodel/2994296-initwithmodeldescription?language=objc
func (c_ CustomModelWrapper) InitWithModelDescriptionParameterDictionaryError(modelDescription IModelDescription, parameters map[string]objc.IObject, error foundation.IError) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("initWithModelDescription:parameterDictionary:error:"), objc.Ptr(modelDescription), parameters, objc.Ptr(error))
	return rv
}

func (c_ CustomModelWrapper) HasPredictionsFromBatchOptionsError() bool {
	return c_.RespondsToSelector(objc.Sel("predictionsFromBatch:options:error:"))
}

// Predicts output values from the given batch of input features. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlcustommodel/2994298-predictionsfrombatch?language=objc
func (c_ CustomModelWrapper) PredictionsFromBatchOptionsError(inputBatch PBatchProvider, options IPredictionOptions, error foundation.IError) BatchProviderWrapper {
	po0 := objc.WrapAsProtocol("MLBatchProvider", inputBatch)
	rv := objc.Call[BatchProviderWrapper](c_, objc.Sel("predictionsFromBatch:options:error:"), po0, objc.Ptr(options), objc.Ptr(error))
	return rv
}

func (c_ CustomModelWrapper) HasPredictionFromFeaturesOptionsError() bool {
	return c_.RespondsToSelector(objc.Sel("predictionFromFeatures:options:error:"))
}

// Predicts output values from the given input features. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlcustommodel/2994297-predictionfromfeatures?language=objc
func (c_ CustomModelWrapper) PredictionFromFeaturesOptionsError(input PFeatureProvider, options IPredictionOptions, error foundation.IError) FeatureProviderWrapper {
	po0 := objc.WrapAsProtocol("MLFeatureProvider", input)
	rv := objc.Call[FeatureProviderWrapper](c_, objc.Sel("predictionFromFeatures:options:error:"), po0, objc.Ptr(options), objc.Ptr(error))
	return rv
}
