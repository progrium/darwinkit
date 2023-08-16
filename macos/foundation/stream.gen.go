// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Stream] class.
var StreamClass = _StreamClass{objc.GetClass("NSStream")}

type _StreamClass struct {
	objc.Class
}

// An interface definition for the [Stream] class.
type IStream interface {
	objc.IObject
	ScheduleInRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode)
	RemoveFromRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode)
	Open()
	Close()
	PropertyForKey(key StreamPropertyKey) objc.Object
	SetPropertyForKey(property objc.IObject, key StreamPropertyKey) bool
	StreamError() Error
	Delegate() StreamDelegateWrapper
	SetDelegate(value PStreamDelegate)
	SetDelegateObject(valueObject objc.IObject)
	StreamStatus() StreamStatus
}

// An abstract class representing a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstream?language=objc
type Stream struct {
	objc.Object
}

func StreamFrom(ptr unsafe.Pointer) Stream {
	return Stream{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _StreamClass) Alloc() Stream {
	rv := objc.Call[Stream](sc, objc.Sel("alloc"))
	return rv
}

func Stream_Alloc() Stream {
	return StreamClass.Alloc()
}

func (sc _StreamClass) New() Stream {
	rv := objc.Call[Stream](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStream() Stream {
	return StreamClass.New()
}

func (s_ Stream) Init() Stream {
	rv := objc.Call[Stream](s_, objc.Sel("init"))
	return rv
}

// Schedules the receiver on a given run loop in a given mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstream/1417370-scheduleinrunloop?language=objc
func (s_ Stream) ScheduleInRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.Call[objc.Void](s_, objc.Sel("scheduleInRunLoop:forMode:"), objc.Ptr(aRunLoop), mode)
}

// Removes the receiver from a given run loop running in a given mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstream/1411285-removefromrunloop?language=objc
func (s_ Stream) RemoveFromRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.Call[objc.Void](s_, objc.Sel("removeFromRunLoop:forMode:"), objc.Ptr(aRunLoop), mode)
}

// Opens the receiving stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstream/1411963-open?language=objc
func (s_ Stream) Open() {
	objc.Call[objc.Void](s_, objc.Sel("open"))
}

// Closes the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstream/1410399-close?language=objc
func (s_ Stream) Close() {
	objc.Call[objc.Void](s_, objc.Sel("close"))
}

// Returns the receiver’s property for a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstream/1410226-propertyforkey?language=objc
func (s_ Stream) PropertyForKey(key StreamPropertyKey) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("propertyForKey:"), key)
	return rv
}

// Attempts to set the value of a given property of the receiver and returns a Boolean value that indicates whether the value is accepted by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstream/1412045-setproperty?language=objc
func (s_ Stream) SetPropertyForKey(property objc.IObject, key StreamPropertyKey) bool {
	rv := objc.Call[bool](s_, objc.Sel("setProperty:forKey:"), property, key)
	return rv
}

// Creates and returns by reference a bound pair of input and output streams. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstream/1412683-getboundstreamswithbuffersize?language=objc
func (sc _StreamClass) GetBoundStreamsWithBufferSizeInputStreamOutputStream(bufferSize uint, inputStream IInputStream, outputStream IOutputStream) {
	objc.Call[objc.Void](sc, objc.Sel("getBoundStreamsWithBufferSize:inputStream:outputStream:"), bufferSize, objc.Ptr(inputStream), objc.Ptr(outputStream))
}

// Creates and returns by reference a bound pair of input and output streams. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstream/1412683-getboundstreamswithbuffersize?language=objc
func Stream_GetBoundStreamsWithBufferSizeInputStreamOutputStream(bufferSize uint, inputStream IInputStream, outputStream IOutputStream) {
	StreamClass.GetBoundStreamsWithBufferSizeInputStreamOutputStream(bufferSize, inputStream, outputStream)
}

// Returns an NSError object representing the stream error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstream/1416359-streamerror?language=objc
func (s_ Stream) StreamError() Error {
	rv := objc.Call[Error](s_, objc.Sel("streamError"))
	return rv
}

// Sets the receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstream/1418423-delegate?language=objc
func (s_ Stream) Delegate() StreamDelegateWrapper {
	rv := objc.Call[StreamDelegateWrapper](s_, objc.Sel("delegate"))
	return rv
}

// Sets the receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstream/1418423-delegate?language=objc
func (s_ Stream) SetDelegate(value PStreamDelegate) {
	po0 := objc.WrapAsProtocol("NSStreamDelegate", value)
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), po0)
}

// Sets the receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstream/1418423-delegate?language=objc
func (s_ Stream) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// Returns the receiver’s status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstream/1413038-streamstatus?language=objc
func (s_ Stream) StreamStatus() StreamStatus {
	rv := objc.Call[StreamStatus](s_, objc.Sel("streamStatus"))
	return rv
}
