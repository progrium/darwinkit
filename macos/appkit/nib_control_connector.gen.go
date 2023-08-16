// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NibControlConnector] class.
var NibControlConnectorClass = _NibControlConnectorClass{objc.GetClass("NSNibControlConnector")}

type _NibControlConnectorClass struct {
	objc.Class
}

// An interface definition for the [NibControlConnector] class.
type INibControlConnector interface {
	INibConnector
}

// A control connection between two Interface Builder objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnibcontrolconnector?language=objc
type NibControlConnector struct {
	NibConnector
}

func NibControlConnectorFrom(ptr unsafe.Pointer) NibControlConnector {
	return NibControlConnector{
		NibConnector: NibConnectorFrom(ptr),
	}
}

func (nc _NibControlConnectorClass) Alloc() NibControlConnector {
	rv := objc.Call[NibControlConnector](nc, objc.Sel("alloc"))
	return rv
}

func NibControlConnector_Alloc() NibControlConnector {
	return NibControlConnectorClass.Alloc()
}

func (nc _NibControlConnectorClass) New() NibControlConnector {
	rv := objc.Call[NibControlConnector](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNibControlConnector() NibControlConnector {
	return NibControlConnectorClass.New()
}

func (n_ NibControlConnector) Init() NibControlConnector {
	rv := objc.Call[NibControlConnector](n_, objc.Sel("init"))
	return rv
}
