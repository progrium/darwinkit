// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OpenGLPixelFormat] class.
var OpenGLPixelFormatClass = _OpenGLPixelFormatClass{objc.GetClass("NSOpenGLPixelFormat")}

type _OpenGLPixelFormatClass struct {
	objc.Class
}

// An interface definition for the [OpenGLPixelFormat] class.
type IOpenGLPixelFormat interface {
	objc.IObject
}

// An object that specifies the types of buffers and other attributes of the OpenGL context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenglpixelformat?language=objc
type OpenGLPixelFormat struct {
	objc.Object
}

func OpenGLPixelFormatFrom(ptr unsafe.Pointer) OpenGLPixelFormat {
	return OpenGLPixelFormat{
		Object: objc.ObjectFrom(ptr),
	}
}

func (oc _OpenGLPixelFormatClass) Alloc() OpenGLPixelFormat {
	rv := objc.Call[OpenGLPixelFormat](oc, objc.Sel("alloc"))
	return rv
}

func OpenGLPixelFormat_Alloc() OpenGLPixelFormat {
	return OpenGLPixelFormatClass.Alloc()
}

func (oc _OpenGLPixelFormatClass) New() OpenGLPixelFormat {
	rv := objc.Call[OpenGLPixelFormat](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOpenGLPixelFormat() OpenGLPixelFormat {
	return OpenGLPixelFormatClass.New()
}

func (o_ OpenGLPixelFormat) Init() OpenGLPixelFormat {
	rv := objc.Call[OpenGLPixelFormat](o_, objc.Sel("init"))
	return rv
}
