// AUTO-GENERATED CODE, DO NOT MODIFY

package coremediaio

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExtensionStreamProperties] class.
var ExtensionStreamPropertiesClass = _ExtensionStreamPropertiesClass{objc.GetClass("CMIOExtensionStreamProperties")}

type _ExtensionStreamPropertiesClass struct {
	objc.Class
}

// An interface definition for the [ExtensionStreamProperties] class.
type IExtensionStreamProperties interface {
	objc.IObject
	SetPropertyStateForProperty(propertyState IExtensionPropertyState, property ExtensionProperty)
	MaxFrameDuration() foundation.Dictionary
	SetMaxFrameDuration(value foundation.Dictionary)
	SinkBufferUnderrunCount() foundation.Number
	SetSinkBufferUnderrunCount(value foundation.INumber)
	SinkBuffersRequiredForStartup() foundation.Number
	SetSinkBuffersRequiredForStartup(value foundation.INumber)
	ActiveFormatIndex() foundation.Number
	SetActiveFormatIndex(value foundation.INumber)
	FrameDuration() foundation.Dictionary
	SetFrameDuration(value foundation.Dictionary)
	SinkEndOfData() foundation.Number
	SetSinkEndOfData(value foundation.INumber)
	SinkBufferQueueSize() foundation.Number
	SetSinkBufferQueueSize(value foundation.INumber)
	PropertiesDictionary() map[ExtensionProperty]ExtensionPropertyState
	SetPropertiesDictionary(value map[ExtensionProperty]IExtensionPropertyState)
}

// An object that describes the properties of an extension stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties?language=objc
type ExtensionStreamProperties struct {
	objc.Object
}

func ExtensionStreamPropertiesFrom(ptr unsafe.Pointer) ExtensionStreamProperties {
	return ExtensionStreamProperties{
		Object: objc.ObjectFrom(ptr),
	}
}

func (e_ ExtensionStreamProperties) InitWithDictionary(propertiesDictionary map[ExtensionProperty]IExtensionPropertyState) ExtensionStreamProperties {
	rv := objc.Call[ExtensionStreamProperties](e_, objc.Sel("initWithDictionary:"), propertiesDictionary)
	return rv
}

// Creates a properties object that provides the specified properties and default states. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915960-initwithdictionary?language=objc
func NewExtensionStreamPropertiesWithDictionary(propertiesDictionary map[ExtensionProperty]IExtensionPropertyState) ExtensionStreamProperties {
	instance := ExtensionStreamPropertiesClass.Alloc().InitWithDictionary(propertiesDictionary)
	instance.Autorelease()
	return instance
}

func (ec _ExtensionStreamPropertiesClass) StreamPropertiesWithDictionary(propertiesDictionary map[ExtensionProperty]IExtensionPropertyState) ExtensionStreamProperties {
	rv := objc.Call[ExtensionStreamProperties](ec, objc.Sel("streamPropertiesWithDictionary:"), propertiesDictionary)
	return rv
}

// Returns a new properties object that provides the specified properties and default states. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915968-streampropertieswithdictionary?language=objc
func ExtensionStreamProperties_StreamPropertiesWithDictionary(propertiesDictionary map[ExtensionProperty]IExtensionPropertyState) ExtensionStreamProperties {
	return ExtensionStreamPropertiesClass.StreamPropertiesWithDictionary(propertiesDictionary)
}

func (ec _ExtensionStreamPropertiesClass) Alloc() ExtensionStreamProperties {
	rv := objc.Call[ExtensionStreamProperties](ec, objc.Sel("alloc"))
	return rv
}

func ExtensionStreamProperties_Alloc() ExtensionStreamProperties {
	return ExtensionStreamPropertiesClass.Alloc()
}

func (ec _ExtensionStreamPropertiesClass) New() ExtensionStreamProperties {
	rv := objc.Call[ExtensionStreamProperties](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionStreamProperties() ExtensionStreamProperties {
	return ExtensionStreamPropertiesClass.New()
}

func (e_ ExtensionStreamProperties) Init() ExtensionStreamProperties {
	rv := objc.Call[ExtensionStreamProperties](e_, objc.Sel("init"))
	return rv
}

// Sets the state of the specified property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915963-setpropertystate?language=objc
func (e_ ExtensionStreamProperties) SetPropertyStateForProperty(propertyState IExtensionPropertyState, property ExtensionProperty) {
	objc.Call[objc.Void](e_, objc.Sel("setPropertyState:forProperty:"), objc.Ptr(propertyState), property)
}

// The maximum duration of a frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915961-maxframeduration?language=objc
func (e_ ExtensionStreamProperties) MaxFrameDuration() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](e_, objc.Sel("maxFrameDuration"))
	return rv
}

// The maximum duration of a frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915961-maxframeduration?language=objc
func (e_ ExtensionStreamProperties) SetMaxFrameDuration(value foundation.Dictionary) {
	objc.Call[objc.Void](e_, objc.Sel("setMaxFrameDuration:"), value)
}

// The buffer underrun count. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915965-sinkbufferunderruncount?language=objc
func (e_ ExtensionStreamProperties) SinkBufferUnderrunCount() foundation.Number {
	rv := objc.Call[foundation.Number](e_, objc.Sel("sinkBufferUnderrunCount"))
	return rv
}

// The buffer underrun count. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915965-sinkbufferunderruncount?language=objc
func (e_ ExtensionStreamProperties) SetSinkBufferUnderrunCount(value foundation.INumber) {
	objc.Call[objc.Void](e_, objc.Sel("setSinkBufferUnderrunCount:"), objc.Ptr(value))
}

// The number of buffers the stream requires for startup. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915966-sinkbuffersrequiredforstartup?language=objc
func (e_ ExtensionStreamProperties) SinkBuffersRequiredForStartup() foundation.Number {
	rv := objc.Call[foundation.Number](e_, objc.Sel("sinkBuffersRequiredForStartup"))
	return rv
}

// The number of buffers the stream requires for startup. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915966-sinkbuffersrequiredforstartup?language=objc
func (e_ ExtensionStreamProperties) SetSinkBuffersRequiredForStartup(value foundation.INumber) {
	objc.Call[objc.Void](e_, objc.Sel("setSinkBuffersRequiredForStartup:"), objc.Ptr(value))
}

// The index of the active format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915958-activeformatindex?language=objc
func (e_ ExtensionStreamProperties) ActiveFormatIndex() foundation.Number {
	rv := objc.Call[foundation.Number](e_, objc.Sel("activeFormatIndex"))
	return rv
}

// The index of the active format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915958-activeformatindex?language=objc
func (e_ ExtensionStreamProperties) SetActiveFormatIndex(value foundation.INumber) {
	objc.Call[objc.Void](e_, objc.Sel("setActiveFormatIndex:"), objc.Ptr(value))
}

// A dictionary representation of a frame duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915959-frameduration?language=objc
func (e_ ExtensionStreamProperties) FrameDuration() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](e_, objc.Sel("frameDuration"))
	return rv
}

// A dictionary representation of a frame duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915959-frameduration?language=objc
func (e_ ExtensionStreamProperties) SetFrameDuration(value foundation.Dictionary) {
	objc.Call[objc.Void](e_, objc.Sel("setFrameDuration:"), value)
}

// A value that indicates whether the stream has reached its end. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915967-sinkendofdata?language=objc
func (e_ ExtensionStreamProperties) SinkEndOfData() foundation.Number {
	rv := objc.Call[foundation.Number](e_, objc.Sel("sinkEndOfData"))
	return rv
}

// A value that indicates whether the stream has reached its end. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915967-sinkendofdata?language=objc
func (e_ ExtensionStreamProperties) SetSinkEndOfData(value foundation.INumber) {
	objc.Call[objc.Void](e_, objc.Sel("setSinkEndOfData:"), objc.Ptr(value))
}

// The buffer queue size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915964-sinkbufferqueuesize?language=objc
func (e_ ExtensionStreamProperties) SinkBufferQueueSize() foundation.Number {
	rv := objc.Call[foundation.Number](e_, objc.Sel("sinkBufferQueueSize"))
	return rv
}

// The buffer queue size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915964-sinkbufferqueuesize?language=objc
func (e_ ExtensionStreamProperties) SetSinkBufferQueueSize(value foundation.INumber) {
	objc.Call[objc.Void](e_, objc.Sel("setSinkBufferQueueSize:"), objc.Ptr(value))
}

// A dictionary representation of the property state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915962-propertiesdictionary?language=objc
func (e_ ExtensionStreamProperties) PropertiesDictionary() map[ExtensionProperty]ExtensionPropertyState {
	rv := objc.Call[map[ExtensionProperty]ExtensionPropertyState](e_, objc.Sel("propertiesDictionary"))
	return rv
}

// A dictionary representation of the property state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamproperties/3915962-propertiesdictionary?language=objc
func (e_ ExtensionStreamProperties) SetPropertiesDictionary(value map[ExtensionProperty]IExtensionPropertyState) {
	objc.Call[objc.Void](e_, objc.Sel("setPropertiesDictionary:"), value)
}
