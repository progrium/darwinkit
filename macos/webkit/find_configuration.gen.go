// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var FindConfigurationClass = _FindConfigurationClass{objc.GetClass("WKFindConfiguration")}

type _FindConfigurationClass struct {
	objc.Class
}

type IFindConfiguration interface {
	objc.IObject
	Backwards() bool
	SetBackwards(value bool)
	CaseSensitive() bool
	SetCaseSensitive(value bool)
	Wraps() bool
	SetWraps(value bool)
}

type FindConfiguration struct {
	objc.Object
}

func MakeFindConfiguration(ptr unsafe.Pointer) FindConfiguration {
	return FindConfiguration{
		Object: objc.MakeObject(ptr),
	}
}

func (fc _FindConfigurationClass) Alloc() FindConfiguration {
	rv := objc.CallMethod[FindConfiguration](fc, objc.GetSelector("alloc"))
	return rv
}

func FindConfiguration_Alloc() FindConfiguration {
	return FindConfigurationClass.Alloc()
}

func (fc _FindConfigurationClass) New() FindConfiguration {
	rv := objc.CallMethod[FindConfiguration](fc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewFindConfiguration() FindConfiguration {
	return FindConfigurationClass.New()
}

func FindConfiguration_New() FindConfiguration {
	return FindConfigurationClass.New()
}

func (f_ FindConfiguration) Init() FindConfiguration {
	rv := objc.CallMethod[FindConfiguration](f_, objc.GetSelector("init"))
	return rv
}

func FindConfiguration_Init() FindConfiguration {
	return FindConfigurationClass.Alloc().Init()
}

func (f_ FindConfiguration) Backwards() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("backwards"))
	return rv
}

func (f_ FindConfiguration) SetBackwards(value bool) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setBackwards:"), value)
}

func (f_ FindConfiguration) CaseSensitive() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("caseSensitive"))
	return rv
}

func (f_ FindConfiguration) SetCaseSensitive(value bool) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setCaseSensitive:"), value)
}

func (f_ FindConfiguration) Wraps() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("wraps"))
	return rv
}

func (f_ FindConfiguration) SetWraps(value bool) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setWraps:"), value)
}
