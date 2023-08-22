// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Renderer] class.
var RendererClass = _RendererClass{objc.GetClass("QCRenderer")}

type _RendererClass struct {
	objc.Class
}

// An interface definition for the [Renderer] class.
type IRenderer interface {
	objc.IObject
}

// A QCRenderer class is designed for low-level rendering of Quartz Composer compositions. This is the class to use if you want to be in charge of rendering a composition to a specific OpenGL context—either using the NSOpenGLContext class or a CGLContextObj object. QCRenderer also allows you to load, play, and control a composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qcrenderer?language=objc
type Renderer struct {
	objc.Object
}

func RendererFrom(ptr unsafe.Pointer) Renderer {
	return Renderer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RendererClass) Alloc() Renderer {
	rv := objc.Call[Renderer](rc, objc.Sel("alloc"))
	return rv
}

func Renderer_Alloc() Renderer {
	return RendererClass.Alloc()
}

func (rc _RendererClass) New() Renderer {
	rv := objc.Call[Renderer](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRenderer() Renderer {
	return RendererClass.New()
}

func (r_ Renderer) Init() Renderer {
	rv := objc.Call[Renderer](r_, objc.Sel("init"))
	return rv
}
