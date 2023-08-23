// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"github.com/progrium/macdriver/objc"
)

// The QCPlugInInputImageSource protocol eliminates the need to use explicit image types for the image input ports on your custom patch. Not only does using the protocol avoid restrictions of a specific image type, but it avoids impedance mismatches, and provides better performance by deferring pixel computation until it is needed. When you need to access the pixels in an image, you simply convert the image to a representation (texture or buffer) using one of the methods defined by the QCPlugInInputImageSource protocol. Use a texture representation when you want to use input images on the GPU. Use a buffer representation when you want to use input images on the CPU. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qcplugininputimagesource?language=objc
type PPlugInInputImageSource interface {
}

// A concrete type wrapper for the [PPlugInInputImageSource] protocol.
type PlugInInputImageSourceWrapper struct {
	objc.Object
}
