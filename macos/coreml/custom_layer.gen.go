// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// An interface that defines the behavior of a custom layer in your neural network model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlcustomlayer?language=objc
type PCustomLayer interface {
	// optional
	InitWithParameterDictionaryError(parameters map[string]objc.Object, error foundation.Error) objc.IObject
	HasInitWithParameterDictionaryError() bool

	// optional
	SetWeightDataError(weights [][]byte, error foundation.Error) bool
	HasSetWeightDataError() bool

	// optional
	EncodeToCommandBufferInputsOutputsError(commandBuffer metal.CommandBufferWrapper, inputs []metal.TextureWrapper, outputs []metal.TextureWrapper, error foundation.Error) bool
	HasEncodeToCommandBufferInputsOutputsError() bool

	// optional
	EvaluateOnCPUWithInputsOutputsError(inputs []MultiArray, outputs []MultiArray, error foundation.Error) bool
	HasEvaluateOnCPUWithInputsOutputsError() bool

	// optional
	OutputShapesForInputShapesError(inputShapes [][]foundation.Number, error foundation.Error) [][]foundation.INumber
	HasOutputShapesForInputShapesError() bool
}

// A concrete type wrapper for the [PCustomLayer] protocol.
type CustomLayerWrapper struct {
	objc.Object
}

func (c_ CustomLayerWrapper) HasInitWithParameterDictionaryError() bool {
	return c_.RespondsToSelector(objc.Sel("initWithParameterDictionary:error:"))
}

// Initializes the custom layer implementation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlcustomlayer/2935523-initwithparameterdictionary?language=objc
func (c_ CustomLayerWrapper) InitWithParameterDictionaryError(parameters map[string]objc.IObject, error foundation.IError) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("initWithParameterDictionary:error:"), parameters, objc.Ptr(error))
	return rv
}

func (c_ CustomLayerWrapper) HasSetWeightDataError() bool {
	return c_.RespondsToSelector(objc.Sel("setWeightData:error:"))
}

// Assigns the weights for the connections within the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlcustomlayer/2936860-setweightdata?language=objc
func (c_ CustomLayerWrapper) SetWeightDataError(weights [][]byte, error foundation.IError) bool {
	rv := objc.Call[bool](c_, objc.Sel("setWeightData:error:"), weights, objc.Ptr(error))
	return rv
}

func (c_ CustomLayerWrapper) HasEncodeToCommandBufferInputsOutputsError() bool {
	return c_.RespondsToSelector(objc.Sel("encodeToCommandBuffer:inputs:outputs:error:"))
}

// Encodes GPU commands to evaluate the custom layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlcustomlayer/2936859-encodetocommandbuffer?language=objc
func (c_ CustomLayerWrapper) EncodeToCommandBufferInputsOutputsError(commandBuffer metal.PCommandBuffer, inputs []metal.PTexture, outputs []metal.PTexture, error foundation.IError) bool {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[bool](c_, objc.Sel("encodeToCommandBuffer:inputs:outputs:error:"), po0, inputs, outputs, objc.Ptr(error))
	return rv
}

func (c_ CustomLayerWrapper) HasEvaluateOnCPUWithInputsOutputsError() bool {
	return c_.RespondsToSelector(objc.Sel("evaluateOnCPUWithInputs:outputs:error:"))
}

// Evaluates the custom layer with the given inputs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlcustomlayer/2935520-evaluateoncpuwithinputs?language=objc
func (c_ CustomLayerWrapper) EvaluateOnCPUWithInputsOutputsError(inputs []IMultiArray, outputs []IMultiArray, error foundation.IError) bool {
	rv := objc.Call[bool](c_, objc.Sel("evaluateOnCPUWithInputs:outputs:error:"), inputs, outputs, objc.Ptr(error))
	return rv
}

func (c_ CustomLayerWrapper) HasOutputShapesForInputShapesError() bool {
	return c_.RespondsToSelector(objc.Sel("outputShapesForInputShapes:error:"))
}

// Calculates the shapes of the output of this layer for the given input shapes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlcustomlayer/2935525-outputshapesforinputshapes?language=objc
func (c_ CustomLayerWrapper) OutputShapesForInputShapesError(inputShapes [][]foundation.INumber, error foundation.IError) [][]foundation.Number {
	rv := objc.Call[[][]foundation.Number](c_, objc.Sel("outputShapesForInputShapes:error:"), inputShapes, objc.Ptr(error))
	return rv
}
