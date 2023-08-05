// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var InputStreamClass = _InputStreamClass{objc.GetClass("NSInputStream")}

type _InputStreamClass struct {
	objc.Class
}

type IInputStream interface {
	IStream
	HasBytesAvailable() bool
}

type InputStream struct {
	Stream
}

func MakeInputStream(ptr unsafe.Pointer) InputStream {
	return InputStream{
		Stream: MakeStream(ptr),
	}
}

func (ic _InputStreamClass) InputStreamWithData(data []byte) InputStream {
	rv := objc.CallMethod[InputStream](ic, objc.GetSelector("inputStreamWithData:"), data)
	return rv
}

func InputStream_InputStreamWithData(data []byte) InputStream {
	return InputStreamClass.InputStreamWithData(data)
}

func (ic _InputStreamClass) InputStreamWithFileAtPath(path string) InputStream {
	rv := objc.CallMethod[InputStream](ic, objc.GetSelector("inputStreamWithFileAtPath:"), path)
	return rv
}

func InputStream_InputStreamWithFileAtPath(path string) InputStream {
	return InputStreamClass.InputStreamWithFileAtPath(path)
}

func (ic _InputStreamClass) InputStreamWithURL(url IURL) InputStream {
	rv := objc.CallMethod[InputStream](ic, objc.GetSelector("inputStreamWithURL:"), objc.ExtractPtr(url))
	return rv
}

func InputStream_InputStreamWithURL(url IURL) InputStream {
	return InputStreamClass.InputStreamWithURL(url)
}

func (i_ InputStream) InitWithData(data []byte) InputStream {
	rv := objc.CallMethod[InputStream](i_, objc.GetSelector("initWithData:"), data)
	return rv
}

func InputStream_InitWithData(data []byte) InputStream {
	return InputStreamClass.Alloc().InitWithData(data)
}

func (i_ InputStream) InitWithFileAtPath(path string) InputStream {
	rv := objc.CallMethod[InputStream](i_, objc.GetSelector("initWithFileAtPath:"), path)
	return rv
}

func InputStream_InitWithFileAtPath(path string) InputStream {
	return InputStreamClass.Alloc().InitWithFileAtPath(path)
}

func (i_ InputStream) InitWithURL(url IURL) InputStream {
	rv := objc.CallMethod[InputStream](i_, objc.GetSelector("initWithURL:"), objc.ExtractPtr(url))
	return rv
}

func InputStream_InitWithURL(url IURL) InputStream {
	return InputStreamClass.Alloc().InitWithURL(url)
}

func (ic _InputStreamClass) Alloc() InputStream {
	rv := objc.CallMethod[InputStream](ic, objc.GetSelector("alloc"))
	return rv
}

func InputStream_Alloc() InputStream {
	return InputStreamClass.Alloc()
}

func (ic _InputStreamClass) New() InputStream {
	rv := objc.CallMethod[InputStream](ic, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewInputStream() InputStream {
	return InputStreamClass.New()
}

func InputStream_New() InputStream {
	return InputStreamClass.New()
}

func (i_ InputStream) Init() InputStream {
	rv := objc.CallMethod[InputStream](i_, objc.GetSelector("init"))
	return rv
}

func InputStream_Init() InputStream {
	return InputStreamClass.Alloc().Init()
}

func (i_ InputStream) HasBytesAvailable() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("hasBytesAvailable"))
	return rv
}
