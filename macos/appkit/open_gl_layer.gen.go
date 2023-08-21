// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/quartzcore"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OpenGLLayer] class.
var OpenGLLayerClass = _OpenGLLayerClass{objc.GetClass("NSOpenGLLayer")}

type _OpenGLLayerClass struct {
	objc.Class
}

// An interface definition for the [OpenGLLayer] class.
type IOpenGLLayer interface {
	quartzcore.IOpenGLLayer
}

// A subclass of CAOpenGLLayer that is suitable for rendering OpenGL into layers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopengllayer?language=objc
type OpenGLLayer struct {
	quartzcore.OpenGLLayer
}

func OpenGLLayerFrom(ptr unsafe.Pointer) OpenGLLayer {
	return OpenGLLayer{
		OpenGLLayer: quartzcore.OpenGLLayerFrom(ptr),
	}
}

func (oc _OpenGLLayerClass) Alloc() OpenGLLayer {
	rv := objc.Call[OpenGLLayer](oc, objc.Sel("alloc"))
	return rv
}

func OpenGLLayer_Alloc() OpenGLLayer {
	return OpenGLLayerClass.Alloc()
}

func (oc _OpenGLLayerClass) New() OpenGLLayer {
	rv := objc.Call[OpenGLLayer](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOpenGLLayer() OpenGLLayer {
	return OpenGLLayerClass.New()
}

func (o_ OpenGLLayer) Init() OpenGLLayer {
	rv := objc.Call[OpenGLLayer](o_, objc.Sel("init"))
	return rv
}

func (oc _OpenGLLayerClass) Layer() OpenGLLayer {
	rv := objc.Call[OpenGLLayer](oc, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func OpenGLLayer_Layer() OpenGLLayer {
	return OpenGLLayerClass.Layer()
}

func (o_ OpenGLLayer) InitWithLayer(layer objc.IObject) OpenGLLayer {
	rv := objc.Call[OpenGLLayer](o_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func NewOpenGLLayerWithLayer(layer objc.IObject) OpenGLLayer {
	instance := OpenGLLayerClass.Alloc().InitWithLayer(layer)
	instance.Autorelease()
	return instance
}

func (o_ OpenGLLayer) ModelLayer() OpenGLLayer {
	rv := objc.Call[OpenGLLayer](o_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func OpenGLLayer_ModelLayer() OpenGLLayer {
	instance := OpenGLLayerClass.Alloc().ModelLayer()
	instance.Autorelease()
	return instance
}

func (o_ OpenGLLayer) PresentationLayer() OpenGLLayer {
	rv := objc.Call[OpenGLLayer](o_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func OpenGLLayer_PresentationLayer() OpenGLLayer {
	instance := OpenGLLayerClass.Alloc().PresentationLayer()
	instance.Autorelease()
	return instance
}
