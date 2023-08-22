// AUTO-GENERATED CODE, DO NOT MODIFY

package coremediaio

import (
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExtensionProvider] class.
var ExtensionProviderClass = _ExtensionProviderClass{objc.GetClass("CMIOExtensionProvider")}

type _ExtensionProviderClass struct {
	objc.Class
}

// An interface definition for the [ExtensionProvider] class.
type IExtensionProvider interface {
	objc.IObject
	NotifyPropertiesChanged(propertyStates map[ExtensionProperty]IExtensionPropertyState)
	AddDeviceError(device IExtensionDevice, outError foundation.IError) bool
	RemoveDeviceError(device IExtensionDevice, outError foundation.IError) bool
	Source() ExtensionProviderSourceWrapper
	ConnectedClients() []ExtensionClient
	ClientQueue() dispatch.Queue
	Devices() []ExtensionDevice
}

// An object that manages device connections for a provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovider?language=objc
type ExtensionProvider struct {
	objc.Object
}

func ExtensionProviderFrom(ptr unsafe.Pointer) ExtensionProvider {
	return ExtensionProvider{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _ExtensionProviderClass) ProviderWithSourceClientQueue(source PExtensionProviderSource, clientQueue dispatch.Queue) ExtensionProvider {
	po0 := objc.WrapAsProtocol("CMIOExtensionProviderSource", source)
	rv := objc.Call[ExtensionProvider](ec, objc.Sel("providerWithSource:clientQueue:"), po0, clientQueue)
	return rv
}

// Returns a new extension provider with the specified source and dispatch queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovider/3915912-providerwithsource?language=objc
func ExtensionProvider_ProviderWithSourceClientQueue(source PExtensionProviderSource, clientQueue dispatch.Queue) ExtensionProvider {
	return ExtensionProviderClass.ProviderWithSourceClientQueue(source, clientQueue)
}

func (e_ ExtensionProvider) InitWithSourceClientQueue(source PExtensionProviderSource, clientQueue dispatch.Queue) ExtensionProvider {
	po0 := objc.WrapAsProtocol("CMIOExtensionProviderSource", source)
	rv := objc.Call[ExtensionProvider](e_, objc.Sel("initWithSource:clientQueue:"), po0, clientQueue)
	return rv
}

// Creates an extension provider with the specified source and dispatch queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovider/3915910-initwithsource?language=objc
func NewExtensionProviderWithSourceClientQueue(source PExtensionProviderSource, clientQueue dispatch.Queue) ExtensionProvider {
	instance := ExtensionProviderClass.Alloc().InitWithSourceClientQueue(source, clientQueue)
	instance.Autorelease()
	return instance
}

func (ec _ExtensionProviderClass) Alloc() ExtensionProvider {
	rv := objc.Call[ExtensionProvider](ec, objc.Sel("alloc"))
	return rv
}

func ExtensionProvider_Alloc() ExtensionProvider {
	return ExtensionProviderClass.Alloc()
}

func (ec _ExtensionProviderClass) New() ExtensionProvider {
	rv := objc.Call[ExtensionProvider](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionProvider() ExtensionProvider {
	return ExtensionProviderClass.New()
}

func (e_ ExtensionProvider) Init() ExtensionProvider {
	rv := objc.Call[ExtensionProvider](e_, objc.Sel("init"))
	return rv
}

// Notifies connected clients of device property changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovider/3915911-notifypropertieschanged?language=objc
func (e_ ExtensionProvider) NotifyPropertiesChanged(propertyStates map[ExtensionProperty]IExtensionPropertyState) {
	objc.Call[objc.Void](e_, objc.Sel("notifyPropertiesChanged:"), propertyStates)
}

// Starts the system extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovider/3915915-startservicewithprovider?language=objc
func (ec _ExtensionProviderClass) StartServiceWithProvider(provider IExtensionProvider) {
	objc.Call[objc.Void](ec, objc.Sel("startServiceWithProvider:"), objc.Ptr(provider))
}

// Starts the system extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovider/3915915-startservicewithprovider?language=objc
func ExtensionProvider_StartServiceWithProvider(provider IExtensionProvider) {
	ExtensionProviderClass.StartServiceWithProvider(provider)
}

// Adds a device to a provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovider/3915906-adddevice?language=objc
func (e_ ExtensionProvider) AddDeviceError(device IExtensionDevice, outError foundation.IError) bool {
	rv := objc.Call[bool](e_, objc.Sel("addDevice:error:"), objc.Ptr(device), objc.Ptr(outError))
	return rv
}

// Removes a device from a provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovider/3915913-removedevice?language=objc
func (e_ ExtensionProvider) RemoveDeviceError(device IExtensionDevice, outError foundation.IError) bool {
	rv := objc.Call[bool](e_, objc.Sel("removeDevice:error:"), objc.Ptr(device), objc.Ptr(outError))
	return rv
}

// The source for the provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovider/3915914-source?language=objc
func (e_ ExtensionProvider) Source() ExtensionProviderSourceWrapper {
	rv := objc.Call[ExtensionProviderSourceWrapper](e_, objc.Sel("source"))
	return rv
}

// An array of connected clients. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovider/3915908-connectedclients?language=objc
func (e_ ExtensionProvider) ConnectedClients() []ExtensionClient {
	rv := objc.Call[[]ExtensionClient](e_, objc.Sel("connectedClients"))
	return rv
}

// The dispatch queue on which the system performs client operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovider/3915907-clientqueue?language=objc
func (e_ ExtensionProvider) ClientQueue() dispatch.Queue {
	rv := objc.Call[dispatch.Queue](e_, objc.Sel("clientQueue"))
	return rv
}

// An array of connected devices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionprovider/3915909-devices?language=objc
func (e_ ExtensionProvider) Devices() []ExtensionDevice {
	rv := objc.Call[[]ExtensionDevice](e_, objc.Sel("devices"))
	return rv
}
