// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/quartzcore"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SynchronizedLayer] class.
var SynchronizedLayerClass = _SynchronizedLayerClass{objc.GetClass("AVSynchronizedLayer")}

type _SynchronizedLayerClass struct {
	objc.Class
}

// An interface definition for the [SynchronizedLayer] class.
type ISynchronizedLayer interface {
	quartzcore.ILayer
	PlayerItem() PlayerItem
	SetPlayerItem(value IPlayerItem)
}

// A Core Animation layer that derives its timing from a player item so that you can synchronize layer animations with media playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsynchronizedlayer?language=objc
type SynchronizedLayer struct {
	quartzcore.Layer
}

func SynchronizedLayerFrom(ptr unsafe.Pointer) SynchronizedLayer {
	return SynchronizedLayer{
		Layer: quartzcore.LayerFrom(ptr),
	}
}

func (sc _SynchronizedLayerClass) Alloc() SynchronizedLayer {
	rv := objc.Call[SynchronizedLayer](sc, objc.Sel("alloc"))
	return rv
}

func SynchronizedLayer_Alloc() SynchronizedLayer {
	return SynchronizedLayerClass.Alloc()
}

func (sc _SynchronizedLayerClass) New() SynchronizedLayer {
	rv := objc.Call[SynchronizedLayer](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSynchronizedLayer() SynchronizedLayer {
	return SynchronizedLayerClass.New()
}

func (s_ SynchronizedLayer) Init() SynchronizedLayer {
	rv := objc.Call[SynchronizedLayer](s_, objc.Sel("init"))
	return rv
}

func (sc _SynchronizedLayerClass) Layer() SynchronizedLayer {
	rv := objc.Call[SynchronizedLayer](sc, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func SynchronizedLayer_Layer() SynchronizedLayer {
	return SynchronizedLayerClass.Layer()
}

func (s_ SynchronizedLayer) InitWithLayer(layer objc.IObject) SynchronizedLayer {
	rv := objc.Call[SynchronizedLayer](s_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func NewSynchronizedLayerWithLayer(layer objc.IObject) SynchronizedLayer {
	instance := SynchronizedLayerClass.Alloc().InitWithLayer(layer)
	instance.Autorelease()
	return instance
}

func (s_ SynchronizedLayer) ModelLayer() SynchronizedLayer {
	rv := objc.Call[SynchronizedLayer](s_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func SynchronizedLayer_ModelLayer() SynchronizedLayer {
	instance := SynchronizedLayerClass.Alloc().ModelLayer()
	instance.Autorelease()
	return instance
}

func (s_ SynchronizedLayer) PresentationLayer() SynchronizedLayer {
	rv := objc.Call[SynchronizedLayer](s_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func SynchronizedLayer_PresentationLayer() SynchronizedLayer {
	instance := SynchronizedLayerClass.Alloc().PresentationLayer()
	instance.Autorelease()
	return instance
}

// Creates a new synchronized layer with timing synchronized with a given player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsynchronizedlayer/1388781-synchronizedlayerwithplayeritem?language=objc
func (sc _SynchronizedLayerClass) SynchronizedLayerWithPlayerItem(playerItem IPlayerItem) SynchronizedLayer {
	rv := objc.Call[SynchronizedLayer](sc, objc.Sel("synchronizedLayerWithPlayerItem:"), objc.Ptr(playerItem))
	return rv
}

// Creates a new synchronized layer with timing synchronized with a given player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsynchronizedlayer/1388781-synchronizedlayerwithplayeritem?language=objc
func SynchronizedLayer_SynchronizedLayerWithPlayerItem(playerItem IPlayerItem) SynchronizedLayer {
	return SynchronizedLayerClass.SynchronizedLayerWithPlayerItem(playerItem)
}

// The player item to which the timing of the layer is synchronized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsynchronizedlayer/1385679-playeritem?language=objc
func (s_ SynchronizedLayer) PlayerItem() PlayerItem {
	rv := objc.Call[PlayerItem](s_, objc.Sel("playerItem"))
	return rv
}

// The player item to which the timing of the layer is synchronized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsynchronizedlayer/1385679-playeritem?language=objc
func (s_ SynchronizedLayer) SetPlayerItem(value IPlayerItem) {
	objc.Call[objc.Void](s_, objc.Sel("setPlayerItem:"), objc.Ptr(value))
}
