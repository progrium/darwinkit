// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OutputStream] class.
var OutputStreamClass = _OutputStreamClass{objc.GetClass("NSOutputStream")}

type _OutputStreamClass struct {
	objc.Class
}

// An interface definition for the [OutputStream] class.
type IOutputStream interface {
	IStream
	WriteMaxLength(buffer *uint8, len uint) int
	HasSpaceAvailable() bool
}

// A stream that provides write-only stream functionality. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoutputstream?language=objc
type OutputStream struct {
	Stream
}

func OutputStreamFrom(ptr unsafe.Pointer) OutputStream {
	return OutputStream{
		Stream: StreamFrom(ptr),
	}
}

func (o_ OutputStream) InitWithURLAppend(url IURL, shouldAppend bool) OutputStream {
	rv := objc.Call[OutputStream](o_, objc.Sel("initWithURL:append:"), objc.Ptr(url), shouldAppend)
	return rv
}

// Returns an initialized output stream for writing to a specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoutputstream/1414446-initwithurl?language=objc
func NewOutputStreamWithURLAppend(url IURL, shouldAppend bool) OutputStream {
	instance := OutputStreamClass.Alloc().InitWithURLAppend(url, shouldAppend)
	instance.Autorelease()
	return instance
}

func (o_ OutputStream) InitToBufferCapacity(buffer *uint8, capacity uint) OutputStream {
	rv := objc.Call[OutputStream](o_, objc.Sel("initToBuffer:capacity:"), buffer, capacity)
	return rv
}

// Returns an initialized output stream that can write to a provided buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoutputstream/1410805-inittobuffer?language=objc
func NewOutputStreamToBufferCapacity(buffer *uint8, capacity uint) OutputStream {
	instance := OutputStreamClass.Alloc().InitToBufferCapacity(buffer, capacity)
	instance.Autorelease()
	return instance
}

func (oc _OutputStreamClass) OutputStreamWithURLAppend(url IURL, shouldAppend bool) OutputStream {
	rv := objc.Call[OutputStream](oc, objc.Sel("outputStreamWithURL:append:"), objc.Ptr(url), shouldAppend)
	return rv
}

// Creates and returns an initialized output stream for writing to a specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoutputstream/1564840-outputstreamwithurl?language=objc
func OutputStream_OutputStreamWithURLAppend(url IURL, shouldAppend bool) OutputStream {
	return OutputStreamClass.OutputStreamWithURLAppend(url, shouldAppend)
}

func (oc _OutputStreamClass) OutputStreamToBufferCapacity(buffer *uint8, capacity uint) OutputStream {
	rv := objc.Call[OutputStream](oc, objc.Sel("outputStreamToBuffer:capacity:"), buffer, capacity)
	return rv
}

// Creates and returns an initialized output stream that can write to a provided buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoutputstream/1564837-outputstreamtobuffer?language=objc
func OutputStream_OutputStreamToBufferCapacity(buffer *uint8, capacity uint) OutputStream {
	return OutputStreamClass.OutputStreamToBufferCapacity(buffer, capacity)
}

func (o_ OutputStream) InitToFileAtPathAppend(path string, shouldAppend bool) OutputStream {
	rv := objc.Call[OutputStream](o_, objc.Sel("initToFileAtPath:append:"), path, shouldAppend)
	return rv
}

// Returns an initialized output stream for writing to a specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoutputstream/1416367-inittofileatpath?language=objc
func NewOutputStreamToFileAtPathAppend(path string, shouldAppend bool) OutputStream {
	instance := OutputStreamClass.Alloc().InitToFileAtPathAppend(path, shouldAppend)
	instance.Autorelease()
	return instance
}

func (oc _OutputStreamClass) OutputStreamToMemory() OutputStream {
	rv := objc.Call[OutputStream](oc, objc.Sel("outputStreamToMemory"))
	return rv
}

// Creates and returns an initialized output stream that will write stream data to memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoutputstream/1411948-outputstreamtomemory?language=objc
func OutputStream_OutputStreamToMemory() OutputStream {
	return OutputStreamClass.OutputStreamToMemory()
}

func (oc _OutputStreamClass) OutputStreamToFileAtPathAppend(path string, shouldAppend bool) OutputStream {
	rv := objc.Call[OutputStream](oc, objc.Sel("outputStreamToFileAtPath:append:"), path, shouldAppend)
	return rv
}

// Creates and returns an initialized output stream for writing to a specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoutputstream/1564841-outputstreamtofileatpath?language=objc
func OutputStream_OutputStreamToFileAtPathAppend(path string, shouldAppend bool) OutputStream {
	return OutputStreamClass.OutputStreamToFileAtPathAppend(path, shouldAppend)
}

func (o_ OutputStream) InitToMemory() OutputStream {
	rv := objc.Call[OutputStream](o_, objc.Sel("initToMemory"))
	return rv
}

// Returns an initialized output stream that will write to memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoutputstream/1409909-inittomemory?language=objc
func NewOutputStreamToMemory() OutputStream {
	instance := OutputStreamClass.Alloc().InitToMemory()
	instance.Autorelease()
	return instance
}

func (oc _OutputStreamClass) Alloc() OutputStream {
	rv := objc.Call[OutputStream](oc, objc.Sel("alloc"))
	return rv
}

func OutputStream_Alloc() OutputStream {
	return OutputStreamClass.Alloc()
}

func (oc _OutputStreamClass) New() OutputStream {
	rv := objc.Call[OutputStream](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOutputStream() OutputStream {
	return OutputStreamClass.New()
}

func (o_ OutputStream) Init() OutputStream {
	rv := objc.Call[OutputStream](o_, objc.Sel("init"))
	return rv
}

// Writes the contents of a provided data buffer to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoutputstream/1410720-write?language=objc
func (o_ OutputStream) WriteMaxLength(buffer *uint8, len uint) int {
	rv := objc.Call[int](o_, objc.Sel("write:maxLength:"), buffer, len)
	return rv
}

// A boolean value that indicates whether the receiver can be written to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoutputstream/1411335-hasspaceavailable?language=objc
func (o_ OutputStream) HasSpaceAvailable() bool {
	rv := objc.Call[bool](o_, objc.Sel("hasSpaceAvailable"))
	return rv
}
