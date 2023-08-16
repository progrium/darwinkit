// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"github.com/progrium/macdriver/objc"
)

// An interface that represents a collection of feature providers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlbatchprovider?language=objc
type PBatchProvider interface {
	// optional
	FeaturesAtIndex(index int) PFeatureProvider
	HasFeaturesAtIndex() bool

	// optional
	Count() int
	HasCount() bool
}

// A concrete type wrapper for the [PBatchProvider] protocol.
type BatchProviderWrapper struct {
	objc.Object
}

func (b_ BatchProviderWrapper) HasFeaturesAtIndex() bool {
	return b_.RespondsToSelector(objc.Sel("featuresAtIndex:"))
}

// Returns the feature provider at the given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlbatchprovider/2994294-featuresatindex?language=objc
func (b_ BatchProviderWrapper) FeaturesAtIndex(index int) FeatureProviderWrapper {
	rv := objc.Call[FeatureProviderWrapper](b_, objc.Sel("featuresAtIndex:"), index)
	return rv
}

func (b_ BatchProviderWrapper) HasCount() bool {
	return b_.RespondsToSelector(objc.Sel("count"))
}

// The number of feature providers in this batch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlbatchprovider/2994293-count?language=objc
func (b_ BatchProviderWrapper) Count() int {
	rv := objc.Call[int](b_, objc.Sel("count"))
	return rv
}
