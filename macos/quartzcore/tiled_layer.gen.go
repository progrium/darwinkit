// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

var TiledLayerClass = _TiledLayerClass{objc.GetClass("CATiledLayer")}

type _TiledLayerClass struct {
	objc.Class
}

type ITiledLayer interface {
	ILayer
	TileSize() coregraphics.Size
	SetTileSize(value coregraphics.Size)
}

type TiledLayer struct {
	Layer
}

func MakeTiledLayer(ptr unsafe.Pointer) TiledLayer {
	return TiledLayer{
		Layer: MakeLayer(ptr),
	}
}

func (tc _TiledLayerClass) Layer() TiledLayer {
	rv := objc.CallMethod[TiledLayer](tc, objc.GetSelector("layer"))
	return rv
}

func TiledLayer_Layer() TiledLayer {
	return TiledLayerClass.Layer()
}

func (t_ TiledLayer) Init() TiledLayer {
	rv := objc.CallMethod[TiledLayer](t_, objc.GetSelector("init"))
	return rv
}

func TiledLayer_Init() TiledLayer {
	return TiledLayerClass.Alloc().Init()
}

func (t_ TiledLayer) InitWithLayer(layer objc.IObject) TiledLayer {
	rv := objc.CallMethod[TiledLayer](t_, objc.GetSelector("initWithLayer:"), objc.ExtractPtr(layer))
	return rv
}

func TiledLayer_InitWithLayer(layer objc.IObject) TiledLayer {
	return TiledLayerClass.Alloc().InitWithLayer(layer)
}

func (t_ TiledLayer) PresentationLayer() TiledLayer {
	rv := objc.CallMethod[TiledLayer](t_, objc.GetSelector("presentationLayer"))
	return rv
}

func TiledLayer_PresentationLayer() TiledLayer {
	return TiledLayerClass.Alloc().PresentationLayer()
}

func (t_ TiledLayer) ModelLayer() TiledLayer {
	rv := objc.CallMethod[TiledLayer](t_, objc.GetSelector("modelLayer"))
	return rv
}

func TiledLayer_ModelLayer() TiledLayer {
	return TiledLayerClass.Alloc().ModelLayer()
}

func (tc _TiledLayerClass) Alloc() TiledLayer {
	rv := objc.CallMethod[TiledLayer](tc, objc.GetSelector("alloc"))
	return rv
}

func TiledLayer_Alloc() TiledLayer {
	return TiledLayerClass.Alloc()
}

func (tc _TiledLayerClass) New() TiledLayer {
	rv := objc.CallMethod[TiledLayer](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTiledLayer() TiledLayer {
	return TiledLayerClass.New()
}

func TiledLayer_New() TiledLayer {
	return TiledLayerClass.New()
}

func (t_ TiledLayer) TileSize() coregraphics.Size {
	rv := objc.CallMethod[coregraphics.Size](t_, objc.GetSelector("tileSize"))
	return rv
}

func (t_ TiledLayer) SetTileSize(value coregraphics.Size) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTileSize:"), value)
}
