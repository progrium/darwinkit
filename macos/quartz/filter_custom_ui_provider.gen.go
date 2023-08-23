// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The IKFilterCustomUIProvider protocol is an addition to the CIFilter class that defines a method for providing a view for a filter. This protocol is implemented by any filter that provides its own user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfiltercustomuiprovider?language=objc
type PFilterCustomUIProvider interface {
	// optional
	ProvideViewForUIConfigurationExcludedKeys(inUIConfiguration foundation.Dictionary, inKeys []objc.Object) IFilterUIView
	HasProvideViewForUIConfigurationExcludedKeys() bool
}

// A concrete type wrapper for the [PFilterCustomUIProvider] protocol.
type FilterCustomUIProviderWrapper struct {
	objc.Object
}

func (f_ FilterCustomUIProviderWrapper) HasProvideViewForUIConfigurationExcludedKeys() bool {
	return f_.RespondsToSelector(objc.Sel("provideViewForUIConfiguration:excludedKeys:"))
}

// Provides a custom view for a filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfiltercustomuiprovider/1427525-provideviewforuiconfiguration?language=objc
func (f_ FilterCustomUIProviderWrapper) ProvideViewForUIConfigurationExcludedKeys(inUIConfiguration foundation.Dictionary, inKeys []objc.IObject) FilterUIView {
	rv := objc.Call[FilterUIView](f_, objc.Sel("provideViewForUIConfiguration:excludedKeys:"), inUIConfiguration, inKeys)
	return rv
}
