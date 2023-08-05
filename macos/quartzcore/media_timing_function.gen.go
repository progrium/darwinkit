// AUTO-GENERATED CODE, DO NOT MODIFY
package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var MediaTimingFunctionClass = _MediaTimingFunctionClass{objc.GetClass("CAMediaTimingFunction")}

type _MediaTimingFunctionClass struct {
	objc.Class
}

type IMediaTimingFunction interface {
	objc.IObject
}

type MediaTimingFunction struct {
	objc.Object
}

func MakeMediaTimingFunction(ptr unsafe.Pointer) MediaTimingFunction {
	return MediaTimingFunction{
		Object: objc.MakeObject(ptr),
	}
}

func (mc _MediaTimingFunctionClass) Alloc() MediaTimingFunction {
	rv := objc.CallMethod[MediaTimingFunction](mc, objc.GetSelector("alloc"))
	return rv
}

func MediaTimingFunction_Alloc() MediaTimingFunction {
	return MediaTimingFunctionClass.Alloc()
}

func (mc _MediaTimingFunctionClass) New() MediaTimingFunction {
	rv := objc.CallMethod[MediaTimingFunction](mc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewMediaTimingFunction() MediaTimingFunction {
	return MediaTimingFunctionClass.New()
}

func MediaTimingFunction_New() MediaTimingFunction {
	return MediaTimingFunctionClass.New()
}

func (m_ MediaTimingFunction) Init() MediaTimingFunction {
	rv := objc.CallMethod[MediaTimingFunction](m_, objc.GetSelector("init"))
	return rv
}

func MediaTimingFunction_Init() MediaTimingFunction {
	return MediaTimingFunctionClass.Alloc().Init()
}
