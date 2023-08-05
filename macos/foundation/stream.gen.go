// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var StreamClass = _StreamClass{objc.GetClass("NSStream")}

type _StreamClass struct {
	objc.Class
}

type IStream interface {
	objc.IObject
	PropertyForKey(key StreamPropertyKey) objc.Object
	SetPropertyForKey(property objc.IObject, key StreamPropertyKey) bool
	Open()
	Close()
	ScheduleInRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode)
	RemoveFromRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode)
	Delegate() StreamDelegateWrapper
	SetDelegate(value IStreamDelegate)
	SetDelegate0(value objc.IObject)
	StreamStatus() StreamStatus
	StreamError() Error
}

type Stream struct {
	objc.Object
}

func MakeStream(ptr unsafe.Pointer) Stream {
	return Stream{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _StreamClass) Alloc() Stream {
	rv := objc.CallMethod[Stream](sc, objc.GetSelector("alloc"))
	return rv
}

func Stream_Alloc() Stream {
	return StreamClass.Alloc()
}

func (sc _StreamClass) New() Stream {
	rv := objc.CallMethod[Stream](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewStream() Stream {
	return StreamClass.New()
}

func Stream_New() Stream {
	return StreamClass.New()
}

func (s_ Stream) Init() Stream {
	rv := objc.CallMethod[Stream](s_, objc.GetSelector("init"))
	return rv
}

func Stream_Init() Stream {
	return StreamClass.Alloc().Init()
}

func (s_ Stream) PropertyForKey(key StreamPropertyKey) objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.GetSelector("propertyForKey:"), key)
	return rv
}

func (s_ Stream) SetPropertyForKey(property objc.IObject, key StreamPropertyKey) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("setProperty:forKey:"), objc.ExtractPtr(property), key)
	return rv
}

func (s_ Stream) Open() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("open"))
}

func (s_ Stream) Close() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("close"))
}

func (s_ Stream) ScheduleInRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("scheduleInRunLoop:forMode:"), objc.ExtractPtr(aRunLoop), mode)
}

func (s_ Stream) RemoveFromRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("removeFromRunLoop:forMode:"), objc.ExtractPtr(aRunLoop), mode)
}

func (sc _StreamClass) GetBoundStreamsWithBufferSizeInputStreamOutputStream(bufferSize uint, inputStream *InputStream, outputStream *OutputStream) {
	objc.CallMethod[objc.Void](sc, objc.GetSelector("getBoundStreamsWithBufferSize:inputStream:outputStream:"), bufferSize, unsafe.Pointer(inputStream), unsafe.Pointer(outputStream))
}

func Stream_GetBoundStreamsWithBufferSizeInputStreamOutputStream(bufferSize uint, inputStream *InputStream, outputStream *OutputStream) {
	StreamClass.GetBoundStreamsWithBufferSizeInputStreamOutputStream(bufferSize, inputStream, outputStream)
}

func (s_ Stream) Delegate() StreamDelegateWrapper {
	rv := objc.CallMethod[StreamDelegateWrapper](s_, objc.GetSelector("delegate"))
	return rv
}

func (s_ Stream) SetDelegate(value IStreamDelegate) {
	po := objc.WrapAsProtocol("NSStreamDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), po)
}

func (s_ Stream) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (s_ Stream) StreamStatus() StreamStatus {
	rv := objc.CallMethod[StreamStatus](s_, objc.GetSelector("streamStatus"))
	return rv
}

func (s_ Stream) StreamError() Error {
	rv := objc.CallMethod[Error](s_, objc.GetSelector("streamError"))
	return rv
}
