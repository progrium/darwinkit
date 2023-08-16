// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TransformLayer] class.
var TransformLayerClass = _TransformLayerClass{objc.GetClass("CATransformLayer")}

type _TransformLayerClass struct {
	objc.Class
}

// An interface definition for the [TransformLayer] class.
type ITransformLayer interface {
	ILayer
}

// Objects used to create true 3D layer hierarchies, rather than the flattened hierarchy rendering model used by other CALayer classes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransformlayer?language=objc
type TransformLayer struct {
	Layer
}

func TransformLayerFrom(ptr unsafe.Pointer) TransformLayer {
	return TransformLayer{
		Layer: LayerFrom(ptr),
	}
}

func (tc _TransformLayerClass) Alloc() TransformLayer {
	rv := objc.Call[TransformLayer](tc, objc.Sel("alloc"))
	return rv
}

func TransformLayer_Alloc() TransformLayer {
	return TransformLayerClass.Alloc()
}

func (tc _TransformLayerClass) New() TransformLayer {
	rv := objc.Call[TransformLayer](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTransformLayer() TransformLayer {
	return TransformLayerClass.New()
}

func (t_ TransformLayer) Init() TransformLayer {
	rv := objc.Call[TransformLayer](t_, objc.Sel("init"))
	return rv
}

func (tc _TransformLayerClass) Layer() TransformLayer {
	rv := objc.Call[TransformLayer](tc, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func TransformLayer_Layer() TransformLayer {
	return TransformLayerClass.Layer()
}

func (t_ TransformLayer) InitWithLayer(layer objc.IObject) TransformLayer {
	rv := objc.Call[TransformLayer](t_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func TransformLayer_InitWithLayer(layer objc.IObject) TransformLayer {
	return TransformLayerClass.Alloc().InitWithLayer(layer)
}

func (t_ TransformLayer) ModelLayer() TransformLayer {
	rv := objc.Call[TransformLayer](t_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func TransformLayer_ModelLayer() TransformLayer {
	return TransformLayerClass.Alloc().ModelLayer()
}

func (t_ TransformLayer) PresentationLayer() TransformLayer {
	rv := objc.Call[TransformLayer](t_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func TransformLayer_PresentationLayer() TransformLayer {
	return TransformLayerClass.Alloc().PresentationLayer()
}
