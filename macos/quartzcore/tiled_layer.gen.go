// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TiledLayer] class.
var TiledLayerClass = _TiledLayerClass{objc.GetClass("CATiledLayer")}

type _TiledLayerClass struct {
	objc.Class
}

// An interface definition for the [TiledLayer] class.
type ITiledLayer interface {
	ILayer
	TileSize() coregraphics.Size
	SetTileSize(value coregraphics.Size)
	LevelsOfDetail() uint
	SetLevelsOfDetail(value uint)
	LevelsOfDetailBias() uint
	SetLevelsOfDetailBias(value uint)
}

// A layer that provides a way to asynchronously provide tiles of the layer's content, potentially cached at multiple levels of detail. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catiledlayer?language=objc
type TiledLayer struct {
	Layer
}

func TiledLayerFrom(ptr unsafe.Pointer) TiledLayer {
	return TiledLayer{
		Layer: LayerFrom(ptr),
	}
}

func (tc _TiledLayerClass) Alloc() TiledLayer {
	rv := objc.Call[TiledLayer](tc, objc.Sel("alloc"))
	return rv
}

func TiledLayer_Alloc() TiledLayer {
	return TiledLayerClass.Alloc()
}

func (tc _TiledLayerClass) New() TiledLayer {
	rv := objc.Call[TiledLayer](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTiledLayer() TiledLayer {
	return TiledLayerClass.New()
}

func (t_ TiledLayer) Init() TiledLayer {
	rv := objc.Call[TiledLayer](t_, objc.Sel("init"))
	return rv
}

func (tc _TiledLayerClass) Layer() TiledLayer {
	rv := objc.Call[TiledLayer](tc, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func TiledLayer_Layer() TiledLayer {
	return TiledLayerClass.Layer()
}

func (t_ TiledLayer) InitWithLayer(layer objc.IObject) TiledLayer {
	rv := objc.Call[TiledLayer](t_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func NewTiledLayerWithLayer(layer objc.IObject) TiledLayer {
	instance := TiledLayerClass.Alloc().InitWithLayer(layer)
	instance.Autorelease()
	return instance
}

func (t_ TiledLayer) ModelLayer() TiledLayer {
	rv := objc.Call[TiledLayer](t_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func TiledLayer_ModelLayer() TiledLayer {
	instance := TiledLayerClass.Alloc().ModelLayer()
	instance.Autorelease()
	return instance
}

func (t_ TiledLayer) PresentationLayer() TiledLayer {
	rv := objc.Call[TiledLayer](t_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func TiledLayer_PresentationLayer() TiledLayer {
	instance := TiledLayerClass.Alloc().PresentationLayer()
	instance.Autorelease()
	return instance
}

// The time, in seconds, that newly added images take to "fade-in" to the rendered representation of the tiled layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catiledlayer/1522145-fadeduration?language=objc
func (tc _TiledLayerClass) FadeDuration() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](tc, objc.Sel("fadeDuration"))
	return rv
}

// The time, in seconds, that newly added images take to "fade-in" to the rendered representation of the tiled layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catiledlayer/1522145-fadeduration?language=objc
func TiledLayer_FadeDuration() corefoundation.TimeInterval {
	return TiledLayerClass.FadeDuration()
}

// The maximum size of each tile used to create the layer's content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catiledlayer/1522114-tilesize?language=objc
func (t_ TiledLayer) TileSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](t_, objc.Sel("tileSize"))
	return rv
}

// The maximum size of each tile used to create the layer's content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catiledlayer/1522114-tilesize?language=objc
func (t_ TiledLayer) SetTileSize(value coregraphics.Size) {
	objc.Call[objc.Void](t_, objc.Sel("setTileSize:"), value)
}

// The number of levels of detail maintained by this layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catiledlayer/1522244-levelsofdetail?language=objc
func (t_ TiledLayer) LevelsOfDetail() uint {
	rv := objc.Call[uint](t_, objc.Sel("levelsOfDetail"))
	return rv
}

// The number of levels of detail maintained by this layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catiledlayer/1522244-levelsofdetail?language=objc
func (t_ TiledLayer) SetLevelsOfDetail(value uint) {
	objc.Call[objc.Void](t_, objc.Sel("setLevelsOfDetail:"), value)
}

// The number of magnified levels of detail for this layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catiledlayer/1522099-levelsofdetailbias?language=objc
func (t_ TiledLayer) LevelsOfDetailBias() uint {
	rv := objc.Call[uint](t_, objc.Sel("levelsOfDetailBias"))
	return rv
}

// The number of magnified levels of detail for this layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catiledlayer/1522099-levelsofdetailbias?language=objc
func (t_ TiledLayer) SetLevelsOfDetailBias(value uint) {
	objc.Call[objc.Void](t_, objc.Sel("setLevelsOfDetailBias:"), value)
}
