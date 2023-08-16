// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Renderer] class.
var RendererClass = _RendererClass{objc.GetClass("CARenderer")}

type _RendererClass struct {
	objc.Class
}

// An interface definition for the [Renderer] class.
type IRenderer interface {
	objc.IObject
	NextFrameTime() corefoundation.TimeInterval
	SetDestination(tex metal.PTexture)
	SetDestinationObject(texObject objc.IObject)
	Render()
	UpdateBounds() coregraphics.Rect
	EndFrame()
	AddUpdateRect(r coregraphics.Rect)
	BeginFrameAtTimeTimeStamp(t corefoundation.TimeInterval, ts *corevideo.TimeStamp)
	Layer() Layer
	SetLayer(value ILayer)
	Bounds() coregraphics.Rect
	SetBounds(value coregraphics.Rect)
}

// A layer that allows an application to render a layer tree into a Core OpenGL context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer?language=objc
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

// Returns the time at which the next update should happen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/1519592-nextframetime?language=objc
func (r_ Renderer) NextFrameTime() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](r_, objc.Sel("nextFrameTime"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/2998892-setdestination?language=objc
func (r_ Renderer) SetDestination(tex metal.PTexture) {
	po0 := objc.WrapAsProtocol("MTLTexture", tex)
	objc.Call[objc.Void](r_, objc.Sel("setDestination:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/2998892-setdestination?language=objc
func (r_ Renderer) SetDestinationObject(texObject objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("setDestination:"), objc.Ptr(texObject))
}

// Render the update region of the current frame to the target context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/1519582-render?language=objc
func (r_ Renderer) Render() {
	objc.Call[objc.Void](r_, objc.Sel("render"))
}

// Creates a layer renderer from a Metal texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/2869759-rendererwithmtltexture?language=objc
func (rc _RendererClass) RendererWithMTLTextureOptions(tex metal.PTexture, dict foundation.Dictionary) Renderer {
	po0 := objc.WrapAsProtocol("MTLTexture", tex)
	rv := objc.Call[Renderer](rc, objc.Sel("rendererWithMTLTexture:options:"), po0, dict)
	return rv
}

// Creates a layer renderer from a Metal texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/2869759-rendererwithmtltexture?language=objc
func Renderer_RendererWithMTLTextureOptions(tex metal.PTexture, dict foundation.Dictionary) Renderer {
	return RendererClass.RendererWithMTLTextureOptions(tex, dict)
}

// Returns the bounds of the update region that contains all pixels that will be rendered by the current frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/1519594-updatebounds?language=objc
func (r_ Renderer) UpdateBounds() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](r_, objc.Sel("updateBounds"))
	return rv
}

// Release any data associated with the current frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/1519593-endframe?language=objc
func (r_ Renderer) EndFrame() {
	objc.Call[objc.Void](r_, objc.Sel("endFrame"))
}

// Adds the rectangle to the update region of the current frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/1519585-addupdaterect?language=objc
func (r_ Renderer) AddUpdateRect(r coregraphics.Rect) {
	objc.Call[objc.Void](r_, objc.Sel("addUpdateRect:"), r)
}

// Begin rendering a frame at the specified time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/1519595-beginframeattime?language=objc
func (r_ Renderer) BeginFrameAtTimeTimeStamp(t corefoundation.TimeInterval, ts *corevideo.TimeStamp) {
	objc.Call[objc.Void](r_, objc.Sel("beginFrameAtTime:timeStamp:"), t, ts)
}

// The root layer of the layer-tree the receiver should render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/1519583-layer?language=objc
func (r_ Renderer) Layer() Layer {
	rv := objc.Call[Layer](r_, objc.Sel("layer"))
	return rv
}

// The root layer of the layer-tree the receiver should render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/1519583-layer?language=objc
func (r_ Renderer) SetLayer(value ILayer) {
	objc.Call[objc.Void](r_, objc.Sel("setLayer:"), objc.Ptr(value))
}

// The bounds of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/1519591-bounds?language=objc
func (r_ Renderer) Bounds() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](r_, objc.Sel("bounds"))
	return rv
}

// The bounds of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/1519591-bounds?language=objc
func (r_ Renderer) SetBounds(value coregraphics.Rect) {
	objc.Call[objc.Void](r_, objc.Sel("setBounds:"), value)
}
