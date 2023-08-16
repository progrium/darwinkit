// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An interface that represents a collection of values for either a model's input or its output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeatureprovider?language=objc
type PFeatureProvider interface {
	// optional
	FeatureValueForName(featureName string) IFeatureValue
	HasFeatureValueForName() bool

	// optional
	FeatureNames() foundation.ISet
	HasFeatureNames() bool
}

// A concrete type wrapper for the [PFeatureProvider] protocol.
type FeatureProviderWrapper struct {
	objc.Object
}

func (f_ FeatureProviderWrapper) HasFeatureValueForName() bool {
	return f_.RespondsToSelector(objc.Sel("featureValueForName:"))
}

// Accesses the feature value given the feature's name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeatureprovider/2879185-featurevalueforname?language=objc
func (f_ FeatureProviderWrapper) FeatureValueForName(featureName string) FeatureValue {
	rv := objc.Call[FeatureValue](f_, objc.Sel("featureValueForName:"), featureName)
	return rv
}

func (f_ FeatureProviderWrapper) HasFeatureNames() bool {
	return f_.RespondsToSelector(objc.Sel("featureNames"))
}

// The set of valid feature names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeatureprovider/2879184-featurenames?language=objc
func (f_ FeatureProviderWrapper) FeatureNames() foundation.Set {
	rv := objc.Call[foundation.Set](f_, objc.Sel("featureNames"))
	return rv
}
