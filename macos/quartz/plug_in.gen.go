// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlugIn] class.
var PlugInClass = _PlugInClass{objc.GetClass("QCPlugIn")}

type _PlugInClass struct {
	objc.Class
}

// An interface definition for the [PlugIn] class.
type IPlugIn interface {
	objc.IObject
}

// The QCPlugIn class provides the base class to subclass for writing custom  Quartz Composer patches. You implement a custom patch by subclassing QCPlugIn, overriding the appropriate methods, packaging the code as an NSBundle object, and installing the bundle in the appropriate location. A bundle can contain more than one subclass  of QCPlugIn, allowing you to provide a suite of custom patches in one bundle. [devLink-1733122] provides detailed instructions on how to create and package a custom patch. QCPlugIn Class Reference supplements the information in the programming guide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qcplugin?language=objc
type PlugIn struct {
	objc.Object
}

func PlugInFrom(ptr unsafe.Pointer) PlugIn {
	return PlugIn{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PlugInClass) Alloc() PlugIn {
	rv := objc.Call[PlugIn](pc, objc.Sel("alloc"))
	return rv
}

func PlugIn_Alloc() PlugIn {
	return PlugInClass.Alloc()
}

func (pc _PlugInClass) New() PlugIn {
	rv := objc.Call[PlugIn](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlugIn() PlugIn {
	return PlugInClass.New()
}

func (p_ PlugIn) Init() PlugIn {
	rv := objc.Call[PlugIn](p_, objc.Sel("init"))
	return rv
}
