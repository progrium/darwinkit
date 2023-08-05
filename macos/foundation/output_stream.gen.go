// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var OutputStreamClass = _OutputStreamClass{objc.GetClass("NSOutputStream")}

type _OutputStreamClass struct {
	objc.Class
}

type IOutputStream interface {
	IStream
	HasSpaceAvailable() bool
}

type OutputStream struct {
	Stream
}

func MakeOutputStream(ptr unsafe.Pointer) OutputStream {
	return OutputStream{
		Stream: MakeStream(ptr),
	}
}

func (oc _OutputStreamClass) OutputStreamToMemory() OutputStream {
	rv := objc.CallMethod[OutputStream](oc, objc.GetSelector("outputStreamToMemory"))
	return rv
}

func OutputStream_OutputStreamToMemory() OutputStream {
	return OutputStreamClass.OutputStreamToMemory()
}

func (oc _OutputStreamClass) OutputStreamToFileAtPathAppend(path string, shouldAppend bool) OutputStream {
	rv := objc.CallMethod[OutputStream](oc, objc.GetSelector("outputStreamToFileAtPath:append:"), path, shouldAppend)
	return rv
}

func OutputStream_OutputStreamToFileAtPathAppend(path string, shouldAppend bool) OutputStream {
	return OutputStreamClass.OutputStreamToFileAtPathAppend(path, shouldAppend)
}

func (oc _OutputStreamClass) OutputStreamWithURLAppend(url IURL, shouldAppend bool) OutputStream {
	rv := objc.CallMethod[OutputStream](oc, objc.GetSelector("outputStreamWithURL:append:"), objc.ExtractPtr(url), shouldAppend)
	return rv
}

func OutputStream_OutputStreamWithURLAppend(url IURL, shouldAppend bool) OutputStream {
	return OutputStreamClass.OutputStreamWithURLAppend(url, shouldAppend)
}

func (o_ OutputStream) InitToMemory() OutputStream {
	rv := objc.CallMethod[OutputStream](o_, objc.GetSelector("initToMemory"))
	return rv
}

func OutputStream_InitToMemory() OutputStream {
	return OutputStreamClass.Alloc().InitToMemory()
}

func (o_ OutputStream) InitToFileAtPathAppend(path string, shouldAppend bool) OutputStream {
	rv := objc.CallMethod[OutputStream](o_, objc.GetSelector("initToFileAtPath:append:"), path, shouldAppend)
	return rv
}

func OutputStream_InitToFileAtPathAppend(path string, shouldAppend bool) OutputStream {
	return OutputStreamClass.Alloc().InitToFileAtPathAppend(path, shouldAppend)
}

func (o_ OutputStream) InitWithURLAppend(url IURL, shouldAppend bool) OutputStream {
	rv := objc.CallMethod[OutputStream](o_, objc.GetSelector("initWithURL:append:"), objc.ExtractPtr(url), shouldAppend)
	return rv
}

func OutputStream_InitWithURLAppend(url IURL, shouldAppend bool) OutputStream {
	return OutputStreamClass.Alloc().InitWithURLAppend(url, shouldAppend)
}

func (oc _OutputStreamClass) Alloc() OutputStream {
	rv := objc.CallMethod[OutputStream](oc, objc.GetSelector("alloc"))
	return rv
}

func OutputStream_Alloc() OutputStream {
	return OutputStreamClass.Alloc()
}

func (oc _OutputStreamClass) New() OutputStream {
	rv := objc.CallMethod[OutputStream](oc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewOutputStream() OutputStream {
	return OutputStreamClass.New()
}

func OutputStream_New() OutputStream {
	return OutputStreamClass.New()
}

func (o_ OutputStream) Init() OutputStream {
	rv := objc.CallMethod[OutputStream](o_, objc.GetSelector("init"))
	return rv
}

func OutputStream_Init() OutputStream {
	return OutputStreamClass.Alloc().Init()
}

func (o_ OutputStream) HasSpaceAvailable() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("hasSpaceAvailable"))
	return rv
}
