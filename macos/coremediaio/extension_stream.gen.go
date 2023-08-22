// AUTO-GENERATED CODE, DO NOT MODIFY

package coremediaio

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExtensionStream] class.
var ExtensionStreamClass = _ExtensionStreamClass{objc.GetClass("CMIOExtensionStream")}

type _ExtensionStreamClass struct {
	objc.Class
}

// An interface definition for the [ExtensionStream] class.
type IExtensionStream interface {
	objc.IObject
	NotifyPropertiesChanged(propertyStates map[ExtensionProperty]IExtensionPropertyState)
	NotifyScheduledOutputChanged(scheduledOutput IExtensionScheduledOutput)
	SendSampleBufferDiscontinuityHostTimeInNanoseconds(sampleBuffer coremedia.SampleBufferRef, discontinuity ExtensionStreamDiscontinuityFlags, hostTimeInNanoseconds uint64)
	ConsumeSampleBufferFromClientCompletionHandler(client IExtensionClient, completionHandler func(sampleBuffer coremedia.SampleBufferRef, sampleBufferSequenceNumber uint64, discontinuity ExtensionStreamDiscontinuityFlags, hasMoreSampleBuffers bool, error foundation.Error))
	StreamID() foundation.UUID
	Direction() ExtensionStreamDirection
	Source() ExtensionStreamSourceWrapper
	LocalizedName() string
	ClockType() ExtensionStreamClockType
	CustomClockConfiguration() ExtensionStreamCustomClockConfiguration
	StreamingClients() []ExtensionClient
}

// An object that represents a stream of media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstream?language=objc
type ExtensionStream struct {
	objc.Object
}

func ExtensionStreamFrom(ptr unsafe.Pointer) ExtensionStream {
	return ExtensionStream{
		Object: objc.ObjectFrom(ptr),
	}
}

func (e_ ExtensionStream) InitWithLocalizedNameStreamIDDirectionClockTypeSource(localizedName string, streamID foundation.IUUID, direction ExtensionStreamDirection, clockType ExtensionStreamClockType, source PExtensionStreamSource) ExtensionStream {
	po4 := objc.WrapAsProtocol("CMIOExtensionStreamSource", source)
	rv := objc.Call[ExtensionStream](e_, objc.Sel("initWithLocalizedName:streamID:direction:clockType:source:"), localizedName, objc.Ptr(streamID), direction, clockType, po4)
	return rv
}

// Creates a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstream/3915934-initwithlocalizedname?language=objc
func NewExtensionStreamWithLocalizedNameStreamIDDirectionClockTypeSource(localizedName string, streamID foundation.IUUID, direction ExtensionStreamDirection, clockType ExtensionStreamClockType, source PExtensionStreamSource) ExtensionStream {
	instance := ExtensionStreamClass.Alloc().InitWithLocalizedNameStreamIDDirectionClockTypeSource(localizedName, streamID, direction, clockType, source)
	instance.Autorelease()
	return instance
}

func (ec _ExtensionStreamClass) StreamWithLocalizedNameStreamIDDirectionClockTypeSource(localizedName string, streamID foundation.IUUID, direction ExtensionStreamDirection, clockType ExtensionStreamClockType, source PExtensionStreamSource) ExtensionStream {
	po4 := objc.WrapAsProtocol("CMIOExtensionStreamSource", source)
	rv := objc.Call[ExtensionStream](ec, objc.Sel("streamWithLocalizedName:streamID:direction:clockType:source:"), localizedName, objc.Ptr(streamID), direction, clockType, po4)
	return rv
}

// Returns a new stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstream/3915942-streamwithlocalizedname?language=objc
func ExtensionStream_StreamWithLocalizedNameStreamIDDirectionClockTypeSource(localizedName string, streamID foundation.IUUID, direction ExtensionStreamDirection, clockType ExtensionStreamClockType, source PExtensionStreamSource) ExtensionStream {
	return ExtensionStreamClass.StreamWithLocalizedNameStreamIDDirectionClockTypeSource(localizedName, streamID, direction, clockType, source)
}

func (ec _ExtensionStreamClass) Alloc() ExtensionStream {
	rv := objc.Call[ExtensionStream](ec, objc.Sel("alloc"))
	return rv
}

func ExtensionStream_Alloc() ExtensionStream {
	return ExtensionStreamClass.Alloc()
}

func (ec _ExtensionStreamClass) New() ExtensionStream {
	rv := objc.Call[ExtensionStream](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionStream() ExtensionStream {
	return ExtensionStreamClass.New()
}

func (e_ ExtensionStream) Init() ExtensionStream {
	rv := objc.Call[ExtensionStream](e_, objc.Sel("init"))
	return rv
}

// Notifies clients about stream property changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstream/3915937-notifypropertieschanged?language=objc
func (e_ ExtensionStream) NotifyPropertiesChanged(propertyStates map[ExtensionProperty]IExtensionPropertyState) {
	objc.Call[objc.Void](e_, objc.Sel("notifyPropertiesChanged:"), propertyStates)
}

// Notifies clients when a particular buffer is output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstream/3915938-notifyscheduledoutputchanged?language=objc
func (e_ ExtensionStream) NotifyScheduledOutputChanged(scheduledOutput IExtensionScheduledOutput) {
	objc.Call[objc.Void](e_, objc.Sel("notifyScheduledOutputChanged:"), objc.Ptr(scheduledOutput))
}

// Sends a media sample to stream client. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstream/3915939-sendsamplebuffer?language=objc
func (e_ ExtensionStream) SendSampleBufferDiscontinuityHostTimeInNanoseconds(sampleBuffer coremedia.SampleBufferRef, discontinuity ExtensionStreamDiscontinuityFlags, hostTimeInNanoseconds uint64) {
	objc.Call[objc.Void](e_, objc.Sel("sendSampleBuffer:discontinuity:hostTimeInNanoseconds:"), sampleBuffer, discontinuity, hostTimeInNanoseconds)
}

// Consumes a sample buffer from a client. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstream/3915931-consumesamplebufferfromclient?language=objc
func (e_ ExtensionStream) ConsumeSampleBufferFromClientCompletionHandler(client IExtensionClient, completionHandler func(sampleBuffer coremedia.SampleBufferRef, sampleBufferSequenceNumber uint64, discontinuity ExtensionStreamDiscontinuityFlags, hasMoreSampleBuffers bool, error foundation.Error)) {
	objc.Call[objc.Void](e_, objc.Sel("consumeSampleBufferFromClient:completionHandler:"), objc.Ptr(client), completionHandler)
}

// A universally unique identifier for the stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstream/3915941-streamid?language=objc
func (e_ ExtensionStream) StreamID() foundation.UUID {
	rv := objc.Call[foundation.UUID](e_, objc.Sel("streamID"))
	return rv
}

// The data-flow direction of the stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstream/3915933-direction?language=objc
func (e_ ExtensionStream) Direction() ExtensionStreamDirection {
	rv := objc.Call[ExtensionStreamDirection](e_, objc.Sel("direction"))
	return rv
}

// The source object for the stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstream/3915940-source?language=objc
func (e_ ExtensionStream) Source() ExtensionStreamSourceWrapper {
	rv := objc.Call[ExtensionStreamSourceWrapper](e_, objc.Sel("source"))
	return rv
}

// A localized name for the stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstream/3915936-localizedname?language=objc
func (e_ ExtensionStream) LocalizedName() string {
	rv := objc.Call[string](e_, objc.Sel("localizedName"))
	return rv
}

// A clock type for the stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstream/3915930-clocktype?language=objc
func (e_ ExtensionStream) ClockType() ExtensionStreamClockType {
	rv := objc.Call[ExtensionStreamClockType](e_, objc.Sel("clockType"))
	return rv
}

// An optional custom clock configuration for a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstream/3915932-customclockconfiguration?language=objc
func (e_ ExtensionStream) CustomClockConfiguration() ExtensionStreamCustomClockConfiguration {
	rv := objc.Call[ExtensionStreamCustomClockConfiguration](e_, objc.Sel("customClockConfiguration"))
	return rv
}

// An array of clients of the stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstream/3915944-streamingclients?language=objc
func (e_ ExtensionStream) StreamingClients() []ExtensionClient {
	rv := objc.Call[[]ExtensionClient](e_, objc.Sel("streamingClients"))
	return rv
}
