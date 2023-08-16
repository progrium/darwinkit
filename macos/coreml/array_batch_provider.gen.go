// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ArrayBatchProvider] class.
var ArrayBatchProviderClass = _ArrayBatchProviderClass{objc.GetClass("MLArrayBatchProvider")}

type _ArrayBatchProviderClass struct {
	objc.Class
}

// An interface definition for the [ArrayBatchProvider] class.
type IArrayBatchProvider interface {
	objc.IObject
	Array() []FeatureProviderWrapper
}

// A convenience wrapper for batches of feature providers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlarraybatchprovider?language=objc
type ArrayBatchProvider struct {
	objc.Object
}

func ArrayBatchProviderFrom(ptr unsafe.Pointer) ArrayBatchProvider {
	return ArrayBatchProvider{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ ArrayBatchProvider) InitWithDictionaryError(dictionary map[string][]objc.IObject, error foundation.IError) ArrayBatchProvider {
	rv := objc.Call[ArrayBatchProvider](a_, objc.Sel("initWithDictionary:error:"), dictionary, objc.Ptr(error))
	return rv
}

// Creates a batch provider based on feature names and their associated arrays of data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlarraybatchprovider/2962853-initwithdictionary?language=objc
func ArrayBatchProvider_InitWithDictionaryError(dictionary map[string][]objc.IObject, error foundation.IError) ArrayBatchProvider {
	return ArrayBatchProviderClass.Alloc().InitWithDictionaryError(dictionary, error)
}

func (a_ ArrayBatchProvider) InitWithFeatureProviderArray(array []PFeatureProvider) ArrayBatchProvider {
	rv := objc.Call[ArrayBatchProvider](a_, objc.Sel("initWithFeatureProviderArray:"), array)
	return rv
}

// Creates the batch provider based on the array of feature providers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlarraybatchprovider/2962854-initwithfeatureproviderarray?language=objc
func ArrayBatchProvider_InitWithFeatureProviderArray(array []PFeatureProvider) ArrayBatchProvider {
	return ArrayBatchProviderClass.Alloc().InitWithFeatureProviderArray(array)
}

func (ac _ArrayBatchProviderClass) Alloc() ArrayBatchProvider {
	rv := objc.Call[ArrayBatchProvider](ac, objc.Sel("alloc"))
	return rv
}

func ArrayBatchProvider_Alloc() ArrayBatchProvider {
	return ArrayBatchProviderClass.Alloc()
}

func (ac _ArrayBatchProviderClass) New() ArrayBatchProvider {
	rv := objc.Call[ArrayBatchProvider](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewArrayBatchProvider() ArrayBatchProvider {
	return ArrayBatchProviderClass.New()
}

func (a_ ArrayBatchProvider) Init() ArrayBatchProvider {
	rv := objc.Call[ArrayBatchProvider](a_, objc.Sel("init"))
	return rv
}

// The array of feature providers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlarraybatchprovider/2962852-array?language=objc
func (a_ ArrayBatchProvider) Array() []FeatureProviderWrapper {
	rv := objc.Call[[]FeatureProviderWrapper](a_, objc.Sel("array"))
	return rv
}
