// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OpenPanelParameters] class.
var OpenPanelParametersClass = _OpenPanelParametersClass{objc.GetClass("WKOpenPanelParameters")}

type _OpenPanelParametersClass struct {
	objc.Class
}

// An interface definition for the [OpenPanelParameters] class.
type IOpenPanelParameters interface {
	objc.IObject
	AllowsMultipleSelection() bool
	AllowsDirectories() bool
}

// The configuration details of a file upload control in your web content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkopenpanelparameters?language=objc
type OpenPanelParameters struct {
	objc.Object
}

func OpenPanelParametersFrom(ptr unsafe.Pointer) OpenPanelParameters {
	return OpenPanelParameters{
		Object: objc.ObjectFrom(ptr),
	}
}

func (oc _OpenPanelParametersClass) Alloc() OpenPanelParameters {
	rv := objc.Call[OpenPanelParameters](oc, objc.Sel("alloc"))
	return rv
}

func OpenPanelParameters_Alloc() OpenPanelParameters {
	return OpenPanelParametersClass.Alloc()
}

func (oc _OpenPanelParametersClass) New() OpenPanelParameters {
	rv := objc.Call[OpenPanelParameters](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOpenPanelParameters() OpenPanelParameters {
	return OpenPanelParametersClass.New()
}

func (o_ OpenPanelParameters) Init() OpenPanelParameters {
	rv := objc.Call[OpenPanelParameters](o_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the file upload control supports multiple files. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkopenpanelparameters/1639524-allowsmultipleselection?language=objc
func (o_ OpenPanelParameters) AllowsMultipleSelection() bool {
	rv := objc.Call[bool](o_, objc.Sel("allowsMultipleSelection"))
	return rv
}

// A Boolean value that indicates whether the file upload control supports the selection of directories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkopenpanelparameters/2937920-allowsdirectories?language=objc
func (o_ OpenPanelParameters) AllowsDirectories() bool {
	rv := objc.Call[bool](o_, objc.Sel("allowsDirectories"))
	return rv
}
