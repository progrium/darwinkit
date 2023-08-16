// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OpenGLContext] class.
var OpenGLContextClass = _OpenGLContextClass{objc.GetClass("NSOpenGLContext")}

type _OpenGLContextClass struct {
	objc.Class
}

// An interface definition for the [OpenGLContext] class.
type IOpenGLContext interface {
	objc.IObject
}

// An object that represents an OpenGL graphics context, into which all OpenGL calls are rendered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenglcontext?language=objc
type OpenGLContext struct {
	objc.Object
}

func OpenGLContextFrom(ptr unsafe.Pointer) OpenGLContext {
	return OpenGLContext{
		Object: objc.ObjectFrom(ptr),
	}
}

func (oc _OpenGLContextClass) Alloc() OpenGLContext {
	rv := objc.Call[OpenGLContext](oc, objc.Sel("alloc"))
	return rv
}

func OpenGLContext_Alloc() OpenGLContext {
	return OpenGLContextClass.Alloc()
}

func (oc _OpenGLContextClass) New() OpenGLContext {
	rv := objc.Call[OpenGLContext](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOpenGLContext() OpenGLContext {
	return OpenGLContextClass.New()
}

func (o_ OpenGLContext) Init() OpenGLContext {
	rv := objc.Call[OpenGLContext](o_, objc.Sel("init"))
	return rv
}
