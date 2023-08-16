// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [InputStream] class.
var InputStreamClass = _InputStreamClass{objc.GetClass("NSInputStream")}

type _InputStreamClass struct {
	objc.Class
}

// An interface definition for the [InputStream] class.
type IInputStream interface {
	IStream
	GetBufferLength(buffer *uint8, len *uint) bool
	ReadMaxLength(buffer *uint8, len uint) int
	HasBytesAvailable() bool
}

// A stream that provides read-only stream functionality. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinputstream?language=objc
type InputStream struct {
	Stream
}

func InputStreamFrom(ptr unsafe.Pointer) InputStream {
	return InputStream{
		Stream: StreamFrom(ptr),
	}
}

func (ic _InputStreamClass) InputStreamWithFileAtPath(path string) InputStream {
	rv := objc.Call[InputStream](ic, objc.Sel("inputStreamWithFileAtPath:"), path)
	return rv
}

// Creates and returns an initialized NSInputStream object that reads data from the file at a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinputstream/1564839-inputstreamwithfileatpath?language=objc
func InputStream_InputStreamWithFileAtPath(path string) InputStream {
	return InputStreamClass.InputStreamWithFileAtPath(path)
}

func (i_ InputStream) InitWithData(data []byte) InputStream {
	rv := objc.Call[InputStream](i_, objc.Sel("initWithData:"), data)
	return rv
}

// Initializes and returns an NSInputStream object for reading from a given NSData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinputstream/1412470-initwithdata?language=objc
func InputStream_InitWithData(data []byte) InputStream {
	return InputStreamClass.Alloc().InitWithData(data)
}

func (i_ InputStream) InitWithURL(url IURL) InputStream {
	rv := objc.Call[InputStream](i_, objc.Sel("initWithURL:"), objc.Ptr(url))
	return rv
}

// Initializes and returns an NSInputStream object that reads data from the file at a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinputstream/1417891-initwithurl?language=objc
func InputStream_InitWithURL(url IURL) InputStream {
	return InputStreamClass.Alloc().InitWithURL(url)
}

func (i_ InputStream) InitWithFileAtPath(path string) InputStream {
	rv := objc.Call[InputStream](i_, objc.Sel("initWithFileAtPath:"), path)
	return rv
}

// Initializes and returns an NSInputStream object that reads data from the file at a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinputstream/1408976-initwithfileatpath?language=objc
func InputStream_InitWithFileAtPath(path string) InputStream {
	return InputStreamClass.Alloc().InitWithFileAtPath(path)
}

func (ic _InputStreamClass) InputStreamWithData(data []byte) InputStream {
	rv := objc.Call[InputStream](ic, objc.Sel("inputStreamWithData:"), data)
	return rv
}

// Creates and returns an initialized NSInputStream object for reading from a given NSData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinputstream/1564842-inputstreamwithdata?language=objc
func InputStream_InputStreamWithData(data []byte) InputStream {
	return InputStreamClass.InputStreamWithData(data)
}

func (ic _InputStreamClass) InputStreamWithURL(url IURL) InputStream {
	rv := objc.Call[InputStream](ic, objc.Sel("inputStreamWithURL:"), objc.Ptr(url))
	return rv
}

// Creates and returns an initialized NSInputStream object that reads data from the file at a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinputstream/1564838-inputstreamwithurl?language=objc
func InputStream_InputStreamWithURL(url IURL) InputStream {
	return InputStreamClass.InputStreamWithURL(url)
}

func (ic _InputStreamClass) Alloc() InputStream {
	rv := objc.Call[InputStream](ic, objc.Sel("alloc"))
	return rv
}

func InputStream_Alloc() InputStream {
	return InputStreamClass.Alloc()
}

func (ic _InputStreamClass) New() InputStream {
	rv := objc.Call[InputStream](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewInputStream() InputStream {
	return InputStreamClass.New()
}

func (i_ InputStream) Init() InputStream {
	rv := objc.Call[InputStream](i_, objc.Sel("init"))
	return rv
}

// Returns by reference a pointer to a read buffer and, by reference, the number of bytes available, and returns a Boolean value that indicates whether the buffer is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinputstream/1416811-getbuffer?language=objc
func (i_ InputStream) GetBufferLength(buffer *uint8, len *uint) bool {
	rv := objc.Call[bool](i_, objc.Sel("getBuffer:length:"), buffer, len)
	return rv
}

// Reads up to a given number of bytes into a given buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinputstream/1411544-read?language=objc
func (i_ InputStream) ReadMaxLength(buffer *uint8, len uint) int {
	rv := objc.Call[int](i_, objc.Sel("read:maxLength:"), buffer, len)
	return rv
}

// A Boolean value that indicates whether the receiver has bytes available to read. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinputstream/1409410-hasbytesavailable?language=objc
func (i_ InputStream) HasBytesAvailable() bool {
	rv := objc.Call[bool](i_, objc.Sel("hasBytesAvailable"))
	return rv
}
