// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

var ScrollLayerClass = _ScrollLayerClass{objc.GetClass("CAScrollLayer")}

type _ScrollLayerClass struct {
	objc.Class
}

type IScrollLayer interface {
	ILayer
	ScrollToPoint(p coregraphics.Point)
	ScrollToRect(r coregraphics.Rect)
	ScrollMode() ScrollLayerScrollMode
	SetScrollMode(value ScrollLayerScrollMode)
}

type ScrollLayer struct {
	Layer
}

func MakeScrollLayer(ptr unsafe.Pointer) ScrollLayer {
	return ScrollLayer{
		Layer: MakeLayer(ptr),
	}
}

func (sc _ScrollLayerClass) Layer() ScrollLayer {
	rv := objc.CallMethod[ScrollLayer](sc, objc.GetSelector("layer"))
	return rv
}

func ScrollLayer_Layer() ScrollLayer {
	return ScrollLayerClass.Layer()
}

func (s_ ScrollLayer) Init() ScrollLayer {
	rv := objc.CallMethod[ScrollLayer](s_, objc.GetSelector("init"))
	return rv
}

func ScrollLayer_Init() ScrollLayer {
	return ScrollLayerClass.Alloc().Init()
}

func (s_ ScrollLayer) InitWithLayer(layer objc.IObject) ScrollLayer {
	rv := objc.CallMethod[ScrollLayer](s_, objc.GetSelector("initWithLayer:"), objc.ExtractPtr(layer))
	return rv
}

func ScrollLayer_InitWithLayer(layer objc.IObject) ScrollLayer {
	return ScrollLayerClass.Alloc().InitWithLayer(layer)
}

func (s_ ScrollLayer) PresentationLayer() ScrollLayer {
	rv := objc.CallMethod[ScrollLayer](s_, objc.GetSelector("presentationLayer"))
	return rv
}

func ScrollLayer_PresentationLayer() ScrollLayer {
	return ScrollLayerClass.Alloc().PresentationLayer()
}

func (s_ ScrollLayer) ModelLayer() ScrollLayer {
	rv := objc.CallMethod[ScrollLayer](s_, objc.GetSelector("modelLayer"))
	return rv
}

func ScrollLayer_ModelLayer() ScrollLayer {
	return ScrollLayerClass.Alloc().ModelLayer()
}

func (sc _ScrollLayerClass) Alloc() ScrollLayer {
	rv := objc.CallMethod[ScrollLayer](sc, objc.GetSelector("alloc"))
	return rv
}

func ScrollLayer_Alloc() ScrollLayer {
	return ScrollLayerClass.Alloc()
}

func (sc _ScrollLayerClass) New() ScrollLayer {
	rv := objc.CallMethod[ScrollLayer](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewScrollLayer() ScrollLayer {
	return ScrollLayerClass.New()
}

func ScrollLayer_New() ScrollLayer {
	return ScrollLayerClass.New()
}

func (s_ ScrollLayer) ScrollToPoint(p coregraphics.Point) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("scrollToPoint:"), p)
}

func (s_ ScrollLayer) ScrollToRect(r coregraphics.Rect) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("scrollToRect:"), r)
}

func (s_ ScrollLayer) ScrollMode() ScrollLayerScrollMode {
	rv := objc.CallMethod[ScrollLayerScrollMode](s_, objc.GetSelector("scrollMode"))
	return rv
}

func (s_ ScrollLayer) SetScrollMode(value ScrollLayerScrollMode) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setScrollMode:"), value)
}
