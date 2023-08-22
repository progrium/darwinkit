// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"github.com/progrium/macdriver/objc"
)

// The QCPlugInOuputImageProvider protocol eliminates the need to use explicit image types for the image output ports on a custom patch. The methods in this protocol are called by the Quartz Composer engine when the output image is needed. If your custom patch has an image output port, you need to implement the appropriate methods for rendering image data and to supply information about the rendering destination and the image bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qcpluginoutputimageprovider?language=objc
type PPlugInOutputImageProvider interface {
}

// A concrete type wrapper for the [PPlugInOutputImageProvider] protocol.
type PlugInOutputImageProviderWrapper struct {
	objc.Object
}
