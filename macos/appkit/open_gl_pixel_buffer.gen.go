// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OpenGLPixelBuffer] class.
var OpenGLPixelBufferClass = _OpenGLPixelBufferClass{objc.GetClass("NSOpenGLPixelBuffer")}

type _OpenGLPixelBufferClass struct {
	objc.Class
}

// An interface definition for the [OpenGLPixelBuffer] class.
type IOpenGLPixelBuffer interface {
	objc.IObject
}

// An object that provides access to accelerated offscreen rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenglpixelbuffer?language=objc
type OpenGLPixelBuffer struct {
	objc.Object
}

func OpenGLPixelBufferFrom(ptr unsafe.Pointer) OpenGLPixelBuffer {
	return OpenGLPixelBuffer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (oc _OpenGLPixelBufferClass) Alloc() OpenGLPixelBuffer {
	rv := objc.Call[OpenGLPixelBuffer](oc, objc.Sel("alloc"))
	return rv
}

func OpenGLPixelBuffer_Alloc() OpenGLPixelBuffer {
	return OpenGLPixelBufferClass.Alloc()
}

func (oc _OpenGLPixelBufferClass) New() OpenGLPixelBuffer {
	rv := objc.Call[OpenGLPixelBuffer](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOpenGLPixelBuffer() OpenGLPixelBuffer {
	return OpenGLPixelBufferClass.New()
}

func (o_ OpenGLPixelBuffer) Init() OpenGLPixelBuffer {
	rv := objc.Call[OpenGLPixelBuffer](o_, objc.Sel("init"))
	return rv
}
