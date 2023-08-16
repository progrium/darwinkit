// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// A depth and stencil state object that specifies the depth and stencil configuration and operations used in a render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencilstate?language=objc
type PDepthStencilState interface {
	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	Label() string
	HasLabel() bool
}

// A concrete type wrapper for the [PDepthStencilState] protocol.
type DepthStencilStateWrapper struct {
	objc.Object
}

func (d_ DepthStencilStateWrapper) HasDevice() bool {
	return d_.RespondsToSelector(objc.Sel("device"))
}

// The device from which this state object was created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencilstate/1462453-device?language=objc
func (d_ DepthStencilStateWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](d_, objc.Sel("device"))
	return rv
}

func (d_ DepthStencilStateWrapper) HasLabel() bool {
	return d_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies this object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencilstate/1462472-label?language=objc
func (d_ DepthStencilStateWrapper) Label() string {
	rv := objc.Call[string](d_, objc.Sel("label"))
	return rv
}
