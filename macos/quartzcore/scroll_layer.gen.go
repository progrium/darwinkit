// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScrollLayer] class.
var ScrollLayerClass = _ScrollLayerClass{objc.GetClass("CAScrollLayer")}

type _ScrollLayerClass struct {
	objc.Class
}

// An interface definition for the [ScrollLayer] class.
type IScrollLayer interface {
	ILayer
	ScrollToPoint(p coregraphics.Point)
	ScrollToRect(r coregraphics.Rect)
	ScrollMode() ScrollLayerScrollMode
	SetScrollMode(value ScrollLayerScrollMode)
}

// A layer that displays scrollable content larger than its own bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cascrolllayer?language=objc
type ScrollLayer struct {
	Layer
}

func ScrollLayerFrom(ptr unsafe.Pointer) ScrollLayer {
	return ScrollLayer{
		Layer: LayerFrom(ptr),
	}
}

func (sc _ScrollLayerClass) Alloc() ScrollLayer {
	rv := objc.Call[ScrollLayer](sc, objc.Sel("alloc"))
	return rv
}

func ScrollLayer_Alloc() ScrollLayer {
	return ScrollLayerClass.Alloc()
}

func (sc _ScrollLayerClass) New() ScrollLayer {
	rv := objc.Call[ScrollLayer](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScrollLayer() ScrollLayer {
	return ScrollLayerClass.New()
}

func (s_ ScrollLayer) Init() ScrollLayer {
	rv := objc.Call[ScrollLayer](s_, objc.Sel("init"))
	return rv
}

func (sc _ScrollLayerClass) Layer() ScrollLayer {
	rv := objc.Call[ScrollLayer](sc, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func ScrollLayer_Layer() ScrollLayer {
	return ScrollLayerClass.Layer()
}

func (s_ ScrollLayer) InitWithLayer(layer objc.IObject) ScrollLayer {
	rv := objc.Call[ScrollLayer](s_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func ScrollLayer_InitWithLayer(layer objc.IObject) ScrollLayer {
	return ScrollLayerClass.Alloc().InitWithLayer(layer)
}

func (s_ ScrollLayer) ModelLayer() ScrollLayer {
	rv := objc.Call[ScrollLayer](s_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func ScrollLayer_ModelLayer() ScrollLayer {
	return ScrollLayerClass.Alloc().ModelLayer()
}

func (s_ ScrollLayer) PresentationLayer() ScrollLayer {
	rv := objc.Call[ScrollLayer](s_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func ScrollLayer_PresentationLayer() ScrollLayer {
	return ScrollLayerClass.Alloc().PresentationLayer()
}

// Changes the origin of the receiver to the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cascrolllayer/1522021-scrolltopoint?language=objc
func (s_ ScrollLayer) ScrollToPoint(p coregraphics.Point) {
	objc.Call[objc.Void](s_, objc.Sel("scrollToPoint:"), p)
}

// Scroll the contents of the receiver to ensure that the rectangle is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cascrolllayer/1522167-scrolltorect?language=objc
func (s_ ScrollLayer) ScrollToRect(r coregraphics.Rect) {
	objc.Call[objc.Void](s_, objc.Sel("scrollToRect:"), r)
}

// Defines the axes in which the layer may be scrolled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cascrolllayer/1522111-scrollmode?language=objc
func (s_ ScrollLayer) ScrollMode() ScrollLayerScrollMode {
	rv := objc.Call[ScrollLayerScrollMode](s_, objc.Sel("scrollMode"))
	return rv
}

// Defines the axes in which the layer may be scrolled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cascrolllayer/1522111-scrollmode?language=objc
func (s_ ScrollLayer) SetScrollMode(value ScrollLayerScrollMode) {
	objc.Call[objc.Void](s_, objc.Sel("setScrollMode:"), value)
}
