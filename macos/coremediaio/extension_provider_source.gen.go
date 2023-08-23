// AUTO-GENERATED CODE, DO NOT MODIFY

package coremediaio

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A protocol for objects that act as provider sources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovidersource?language=objc
type PExtensionProviderSource interface {
	// optional
	SetProviderPropertiesError(providerProperties ExtensionProviderProperties, outError foundation.Error) bool
	HasSetProviderPropertiesError() bool

	// optional
	ProviderPropertiesForPropertiesError(properties foundation.Set, outError foundation.Error) IExtensionProviderProperties
	HasProviderPropertiesForPropertiesError() bool

	// optional
	ConnectClientError(client ExtensionClient, outError foundation.Error) bool
	HasConnectClientError() bool

	// optional
	DisconnectClient(client ExtensionClient)
	HasDisconnectClient() bool

	// optional
	AvailableProperties() foundation.ISet
	HasAvailableProperties() bool
}

// A concrete type wrapper for the [PExtensionProviderSource] protocol.
type ExtensionProviderSourceWrapper struct {
	objc.Object
}

func (e_ ExtensionProviderSourceWrapper) HasSetProviderPropertiesError() bool {
	return e_.RespondsToSelector(objc.Sel("setProviderProperties:error:"))
}

// Set the state of provider properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovidersource/3915928-setproviderproperties?language=objc
func (e_ ExtensionProviderSourceWrapper) SetProviderPropertiesError(providerProperties IExtensionProviderProperties, outError foundation.IError) bool {
	rv := objc.Call[bool](e_, objc.Sel("setProviderProperties:error:"), objc.Ptr(providerProperties), objc.Ptr(outError))
	return rv
}

func (e_ ExtensionProviderSourceWrapper) HasProviderPropertiesForPropertiesError() bool {
	return e_.RespondsToSelector(objc.Sel("providerPropertiesForProperties:error:"))
}

// Gets the state of provider properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovidersource/3915927-providerpropertiesforproperties?language=objc
func (e_ ExtensionProviderSourceWrapper) ProviderPropertiesForPropertiesError(properties foundation.ISet, outError foundation.IError) ExtensionProviderProperties {
	rv := objc.Call[ExtensionProviderProperties](e_, objc.Sel("providerPropertiesForProperties:error:"), objc.Ptr(properties), objc.Ptr(outError))
	return rv
}

func (e_ ExtensionProviderSourceWrapper) HasConnectClientError() bool {
	return e_.RespondsToSelector(objc.Sel("connectClient:error:"))
}

// Connects a client to a source’s provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovidersource/3915925-connectclient?language=objc
func (e_ ExtensionProviderSourceWrapper) ConnectClientError(client IExtensionClient, outError foundation.IError) bool {
	rv := objc.Call[bool](e_, objc.Sel("connectClient:error:"), objc.Ptr(client), objc.Ptr(outError))
	return rv
}

func (e_ ExtensionProviderSourceWrapper) HasDisconnectClient() bool {
	return e_.RespondsToSelector(objc.Sel("disconnectClient:"))
}

// Disconnects a client from a source’s provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovidersource/3915926-disconnectclient?language=objc
func (e_ ExtensionProviderSourceWrapper) DisconnectClient(client IExtensionClient) {
	objc.Call[objc.Void](e_, objc.Sel("disconnectClient:"), objc.Ptr(client))
}

func (e_ ExtensionProviderSourceWrapper) HasAvailableProperties() bool {
	return e_.RespondsToSelector(objc.Sel("availableProperties"))
}

// A set of available properties for a provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovidersource/3915924-availableproperties?language=objc
func (e_ ExtensionProviderSourceWrapper) AvailableProperties() foundation.Set {
	rv := objc.Call[foundation.Set](e_, objc.Sel("availableProperties"))
	return rv
}
