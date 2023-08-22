// AUTO-GENERATED CODE, DO NOT MODIFY

package coremediaio

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A protocol for objects that act as stream sources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamsource?language=objc
type PExtensionStreamSource interface {
	// optional
	AuthorizedToStartStreamForClient(client ExtensionClient) bool
	HasAuthorizedToStartStreamForClient() bool

	// optional
	StopStreamAndReturnError(outError foundation.Error) bool
	HasStopStreamAndReturnError() bool

	// optional
	StreamPropertiesForPropertiesError(properties foundation.Set, outError foundation.Error) IExtensionStreamProperties
	HasStreamPropertiesForPropertiesError() bool

	// optional
	SetStreamPropertiesError(streamProperties ExtensionStreamProperties, outError foundation.Error) bool
	HasSetStreamPropertiesError() bool

	// optional
	StartStreamAndReturnError(outError foundation.Error) bool
	HasStartStreamAndReturnError() bool

	// optional
	Formats() []IExtensionStreamFormat
	HasFormats() bool

	// optional
	AvailableProperties() foundation.ISet
	HasAvailableProperties() bool
}

// A concrete type wrapper for the [PExtensionStreamSource] protocol.
type ExtensionStreamSourceWrapper struct {
	objc.Object
}

func (e_ ExtensionStreamSourceWrapper) HasAuthorizedToStartStreamForClient() bool {
	return e_.RespondsToSelector(objc.Sel("authorizedToStartStreamForClient:"))
}

// Determines whether to authorize an app to use this stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamsource/3915970-authorizedtostartstreamforclient?language=objc
func (e_ ExtensionStreamSourceWrapper) AuthorizedToStartStreamForClient(client IExtensionClient) bool {
	rv := objc.Call[bool](e_, objc.Sel("authorizedToStartStreamForClient:"), objc.Ptr(client))
	return rv
}

func (e_ ExtensionStreamSourceWrapper) HasStopStreamAndReturnError() bool {
	return e_.RespondsToSelector(objc.Sel("stopStreamAndReturnError:"))
}

// Stops the stream of media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamsource/3915975-stopstreamandreturnerror?language=objc
func (e_ ExtensionStreamSourceWrapper) StopStreamAndReturnError(outError foundation.IError) bool {
	rv := objc.Call[bool](e_, objc.Sel("stopStreamAndReturnError:"), objc.Ptr(outError))
	return rv
}

func (e_ ExtensionStreamSourceWrapper) HasStreamPropertiesForPropertiesError() bool {
	return e_.RespondsToSelector(objc.Sel("streamPropertiesForProperties:error:"))
}

// Gets the states of specified properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamsource/3915976-streampropertiesforproperties?language=objc
func (e_ ExtensionStreamSourceWrapper) StreamPropertiesForPropertiesError(properties foundation.ISet, outError foundation.IError) ExtensionStreamProperties {
	rv := objc.Call[ExtensionStreamProperties](e_, objc.Sel("streamPropertiesForProperties:error:"), objc.Ptr(properties), objc.Ptr(outError))
	return rv
}

func (e_ ExtensionStreamSourceWrapper) HasSetStreamPropertiesError() bool {
	return e_.RespondsToSelector(objc.Sel("setStreamProperties:error:"))
}

// Sets the property state of a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamsource/3915973-setstreamproperties?language=objc
func (e_ ExtensionStreamSourceWrapper) SetStreamPropertiesError(streamProperties IExtensionStreamProperties, outError foundation.IError) bool {
	rv := objc.Call[bool](e_, objc.Sel("setStreamProperties:error:"), objc.Ptr(streamProperties), objc.Ptr(outError))
	return rv
}

func (e_ ExtensionStreamSourceWrapper) HasStartStreamAndReturnError() bool {
	return e_.RespondsToSelector(objc.Sel("startStreamAndReturnError:"))
}

// Starts the stream of media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamsource/3915974-startstreamandreturnerror?language=objc
func (e_ ExtensionStreamSourceWrapper) StartStreamAndReturnError(outError foundation.IError) bool {
	rv := objc.Call[bool](e_, objc.Sel("startStreamAndReturnError:"), objc.Ptr(outError))
	return rv
}

func (e_ ExtensionStreamSourceWrapper) HasFormats() bool {
	return e_.RespondsToSelector(objc.Sel("formats"))
}

// An array of formats that a stream supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamsource/3915972-formats?language=objc
func (e_ ExtensionStreamSourceWrapper) Formats() []ExtensionStreamFormat {
	rv := objc.Call[[]ExtensionStreamFormat](e_, objc.Sel("formats"))
	return rv
}

func (e_ ExtensionStreamSourceWrapper) HasAvailableProperties() bool {
	return e_.RespondsToSelector(objc.Sel("availableProperties"))
}

// A set of properties available for the stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamsource/3915971-availableproperties?language=objc
func (e_ ExtensionStreamSourceWrapper) AvailableProperties() foundation.Set {
	rv := objc.Call[foundation.Set](e_, objc.Sel("availableProperties"))
	return rv
}
