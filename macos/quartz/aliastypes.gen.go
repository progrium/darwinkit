// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qcplugintexturereleasecallback?language=objc
type PlugInTextureReleaseCallback = func(cgl_ctx objc.Object, name objc.Object, context unsafe.Pointer)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qcpluginbufferreleasecallback?language=objc
type PlugInBufferReleaseCallback = func(address unsafe.Pointer, context unsafe.Pointer)
