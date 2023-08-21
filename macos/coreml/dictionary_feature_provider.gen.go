// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DictionaryFeatureProvider] class.
var DictionaryFeatureProviderClass = _DictionaryFeatureProviderClass{objc.GetClass("MLDictionaryFeatureProvider")}

type _DictionaryFeatureProviderClass struct {
	objc.Class
}

// An interface definition for the [DictionaryFeatureProvider] class.
type IDictionaryFeatureProvider interface {
	objc.IObject
	ObjectForKeyedSubscript(featureName string) FeatureValue
	Dictionary() map[string]FeatureValue
}

// A convenience wrapper for the given dictionary of data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mldictionaryfeatureprovider?language=objc
type DictionaryFeatureProvider struct {
	objc.Object
}

func DictionaryFeatureProviderFrom(ptr unsafe.Pointer) DictionaryFeatureProvider {
	return DictionaryFeatureProvider{
		Object: objc.ObjectFrom(ptr),
	}
}

func (d_ DictionaryFeatureProvider) InitWithDictionaryError(dictionary map[string]objc.IObject, error foundation.IError) DictionaryFeatureProvider {
	rv := objc.Call[DictionaryFeatureProvider](d_, objc.Sel("initWithDictionary:error:"), dictionary, objc.Ptr(error))
	return rv
}

// Creates the feature provider based on a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mldictionaryfeatureprovider/2879366-initwithdictionary?language=objc
func NewDictionaryFeatureProviderWithDictionaryError(dictionary map[string]objc.IObject, error foundation.IError) DictionaryFeatureProvider {
	instance := DictionaryFeatureProviderClass.Alloc().InitWithDictionaryError(dictionary, error)
	instance.Autorelease()
	return instance
}

func (dc _DictionaryFeatureProviderClass) Alloc() DictionaryFeatureProvider {
	rv := objc.Call[DictionaryFeatureProvider](dc, objc.Sel("alloc"))
	return rv
}

func DictionaryFeatureProvider_Alloc() DictionaryFeatureProvider {
	return DictionaryFeatureProviderClass.Alloc()
}

func (dc _DictionaryFeatureProviderClass) New() DictionaryFeatureProvider {
	rv := objc.Call[DictionaryFeatureProvider](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDictionaryFeatureProvider() DictionaryFeatureProvider {
	return DictionaryFeatureProviderClass.New()
}

func (d_ DictionaryFeatureProvider) Init() DictionaryFeatureProvider {
	rv := objc.Call[DictionaryFeatureProvider](d_, objc.Sel("init"))
	return rv
}

// Subscript interface for the feature provider to pass through to the dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mldictionaryfeatureprovider/2881954-objectforkeyedsubscript?language=objc
func (d_ DictionaryFeatureProvider) ObjectForKeyedSubscript(featureName string) FeatureValue {
	rv := objc.Call[FeatureValue](d_, objc.Sel("objectForKeyedSubscript:"), featureName)
	return rv
}

// The backing dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mldictionaryfeatureprovider/2879354-dictionary?language=objc
func (d_ DictionaryFeatureProvider) Dictionary() map[string]FeatureValue {
	rv := objc.Call[map[string]FeatureValue](d_, objc.Sel("dictionary"))
	return rv
}
