// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NibOutletConnector] class.
var NibOutletConnectorClass = _NibOutletConnectorClass{objc.GetClass("NSNibOutletConnector")}

type _NibOutletConnectorClass struct {
	objc.Class
}

// An interface definition for the [NibOutletConnector] class.
type INibOutletConnector interface {
	INibConnector
}

// An outlet connection between Interface Builder objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsniboutletconnector?language=objc
type NibOutletConnector struct {
	NibConnector
}

func NibOutletConnectorFrom(ptr unsafe.Pointer) NibOutletConnector {
	return NibOutletConnector{
		NibConnector: NibConnectorFrom(ptr),
	}
}

func (nc _NibOutletConnectorClass) Alloc() NibOutletConnector {
	rv := objc.Call[NibOutletConnector](nc, objc.Sel("alloc"))
	return rv
}

func NibOutletConnector_Alloc() NibOutletConnector {
	return NibOutletConnectorClass.Alloc()
}

func (nc _NibOutletConnectorClass) New() NibOutletConnector {
	rv := objc.Call[NibOutletConnector](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNibOutletConnector() NibOutletConnector {
	return NibOutletConnectorClass.New()
}

func (n_ NibOutletConnector) Init() NibOutletConnector {
	rv := objc.Call[NibOutletConnector](n_, objc.Sel("init"))
	return rv
}
