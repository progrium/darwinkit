// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coreml"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CoreMLModel] class.
var CoreMLModelClass = _CoreMLModelClass{objc.GetClass("VNCoreMLModel")}

type _CoreMLModelClass struct {
	objc.Class
}

// An interface definition for the [CoreMLModel] class.
type ICoreMLModel interface {
	objc.IObject
	FeatureProvider() coreml.FeatureProviderWrapper
	SetFeatureProvider(value coreml.PFeatureProvider)
	SetFeatureProviderObject(valueObject objc.IObject)
	InputImageFeatureName() string
	SetInputImageFeatureName(value string)
}

// A container for the model to use with Vision requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlmodel?language=objc
type CoreMLModel struct {
	objc.Object
}

func CoreMLModelFrom(ptr unsafe.Pointer) CoreMLModel {
	return CoreMLModel{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CoreMLModelClass) ModelForMLModelError(model coreml.IModel, error foundation.IError) CoreMLModel {
	rv := objc.Call[CoreMLModel](cc, objc.Sel("modelForMLModel:error:"), objc.Ptr(model), objc.Ptr(error))
	return rv
}

// Creates a model container to be used with VNCoreMLRequest. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlmodel/2890148-modelformlmodel?language=objc
func CoreMLModel_ModelForMLModelError(model coreml.IModel, error foundation.IError) CoreMLModel {
	return CoreMLModelClass.ModelForMLModelError(model, error)
}

func (cc _CoreMLModelClass) Alloc() CoreMLModel {
	rv := objc.Call[CoreMLModel](cc, objc.Sel("alloc"))
	return rv
}

func CoreMLModel_Alloc() CoreMLModel {
	return CoreMLModelClass.Alloc()
}

func (cc _CoreMLModelClass) New() CoreMLModel {
	rv := objc.Call[CoreMLModel](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCoreMLModel() CoreMLModel {
	return CoreMLModelClass.New()
}

func (c_ CoreMLModel) Init() CoreMLModel {
	rv := objc.Call[CoreMLModel](c_, objc.Sel("init"))
	return rv
}

// An optional object to support inputs outside Vision. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlmodel/3131933-featureprovider?language=objc
func (c_ CoreMLModel) FeatureProvider() coreml.FeatureProviderWrapper {
	rv := objc.Call[coreml.FeatureProviderWrapper](c_, objc.Sel("featureProvider"))
	return rv
}

// An optional object to support inputs outside Vision. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlmodel/3131933-featureprovider?language=objc
func (c_ CoreMLModel) SetFeatureProvider(value coreml.PFeatureProvider) {
	po0 := objc.WrapAsProtocol("MLFeatureProvider", value)
	objc.Call[objc.Void](c_, objc.Sel("setFeatureProvider:"), po0)
}

// An optional object to support inputs outside Vision. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlmodel/3131933-featureprovider?language=objc
func (c_ CoreMLModel) SetFeatureProviderObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setFeatureProvider:"), objc.Ptr(valueObject))
}

// The name of the MLFeatureValue that Vision sets from the request handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlmodel/3131934-inputimagefeaturename?language=objc
func (c_ CoreMLModel) InputImageFeatureName() string {
	rv := objc.Call[string](c_, objc.Sel("inputImageFeatureName"))
	return rv
}

// The name of the MLFeatureValue that Vision sets from the request handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlmodel/3131934-inputimagefeaturename?language=objc
func (c_ CoreMLModel) SetInputImageFeatureName(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImageFeatureName:"), value)
}
