// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlugIn] class.
var PlugInClass = _PlugInClass{objc.GetClass("CIPlugIn")}

type _PlugInClass struct {
	objc.Class
}

// An interface definition for the [PlugIn] class.
type IPlugIn interface {
	objc.IObject
}

// The mechanism for loading image units in macOS. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciplugin?language=objc
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

// Scans directories for files that have the .plugin extension and then loads only those filters that are marked by the image unit as non-executable filters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciplugin/1437599-loadnonexecutableplugins?language=objc
func (pc _PlugInClass) LoadNonExecutablePlugIns() {
	objc.Call[objc.Void](pc, objc.Sel("loadNonExecutablePlugIns"))
}

// Scans directories for files that have the .plugin extension and then loads only those filters that are marked by the image unit as non-executable filters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciplugin/1437599-loadnonexecutableplugins?language=objc
func PlugIn_LoadNonExecutablePlugIns() {
	PlugInClass.LoadNonExecutablePlugIns()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciplugin/3180431-loadnonexecutableplugin?language=objc
func (pc _PlugInClass) LoadNonExecutablePlugIn(url foundation.IURL) {
	objc.Call[objc.Void](pc, objc.Sel("loadNonExecutablePlugIn:"), objc.Ptr(url))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciplugin/3180431-loadnonexecutableplugin?language=objc
func PlugIn_LoadNonExecutablePlugIn(url foundation.IURL) {
	PlugInClass.LoadNonExecutablePlugIn(url)
}
