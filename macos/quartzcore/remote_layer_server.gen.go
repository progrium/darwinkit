// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RemoteLayerServer] class.
var RemoteLayerServerClass = _RemoteLayerServerClass{objc.GetClass("CARemoteLayerServer")}

type _RemoteLayerServerClass struct {
	objc.Class
}

// An interface definition for the [RemoteLayerServer] class.
type IRemoteLayerServer interface {
	objc.IObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caremotelayerserver?language=objc
type RemoteLayerServer struct {
	objc.Object
}

func RemoteLayerServerFrom(ptr unsafe.Pointer) RemoteLayerServer {
	return RemoteLayerServer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RemoteLayerServerClass) Alloc() RemoteLayerServer {
	rv := objc.Call[RemoteLayerServer](rc, objc.Sel("alloc"))
	return rv
}

func RemoteLayerServer_Alloc() RemoteLayerServer {
	return RemoteLayerServerClass.Alloc()
}

func (rc _RemoteLayerServerClass) New() RemoteLayerServer {
	rv := objc.Call[RemoteLayerServer](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRemoteLayerServer() RemoteLayerServer {
	return RemoteLayerServerClass.New()
}

func (r_ RemoteLayerServer) Init() RemoteLayerServer {
	rv := objc.Call[RemoteLayerServer](r_, objc.Sel("init"))
	return rv
}

// Returns the (singleton) instance of the shared remote layer server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caremotelayerserver/1521954-sharedserver?language=objc
func (rc _RemoteLayerServerClass) SharedServer() RemoteLayerServer {
	rv := objc.Call[RemoteLayerServer](rc, objc.Sel("sharedServer"))
	return rv
}

// Returns the (singleton) instance of the shared remote layer server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caremotelayerserver/1521954-sharedserver?language=objc
func RemoteLayerServer_SharedServer() RemoteLayerServer {
	return RemoteLayerServerClass.SharedServer()
}
