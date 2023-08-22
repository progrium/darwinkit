// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/quartzcore"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerLayer] class.
var PlayerLayerClass = _PlayerLayerClass{objc.GetClass("AVPlayerLayer")}

type _PlayerLayerClass struct {
	objc.Class
}

// An interface definition for the [PlayerLayer] class.
type IPlayerLayer interface {
	quartzcore.ILayer
	VideoGravity() LayerVideoGravity
	SetVideoGravity(value LayerVideoGravity)
	PixelBufferAttributes() map[string]objc.Object
	SetPixelBufferAttributes(value map[string]objc.IObject)
	Player() Player
	SetPlayer(value IPlayer)
	IsReadyForDisplay() bool
	VideoRect() coregraphics.Rect
}

// An object that presents the visual contents of a player object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlayer?language=objc
type PlayerLayer struct {
	quartzcore.Layer
}

func PlayerLayerFrom(ptr unsafe.Pointer) PlayerLayer {
	return PlayerLayer{
		Layer: quartzcore.LayerFrom(ptr),
	}
}

func (pc _PlayerLayerClass) Alloc() PlayerLayer {
	rv := objc.Call[PlayerLayer](pc, objc.Sel("alloc"))
	return rv
}

func PlayerLayer_Alloc() PlayerLayer {
	return PlayerLayerClass.Alloc()
}

func (pc _PlayerLayerClass) New() PlayerLayer {
	rv := objc.Call[PlayerLayer](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerLayer() PlayerLayer {
	return PlayerLayerClass.New()
}

func (p_ PlayerLayer) Init() PlayerLayer {
	rv := objc.Call[PlayerLayer](p_, objc.Sel("init"))
	return rv
}

func (pc _PlayerLayerClass) Layer() PlayerLayer {
	rv := objc.Call[PlayerLayer](pc, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func PlayerLayer_Layer() PlayerLayer {
	return PlayerLayerClass.Layer()
}

func (p_ PlayerLayer) InitWithLayer(layer objc.IObject) PlayerLayer {
	rv := objc.Call[PlayerLayer](p_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func NewPlayerLayerWithLayer(layer objc.IObject) PlayerLayer {
	instance := PlayerLayerClass.Alloc().InitWithLayer(layer)
	instance.Autorelease()
	return instance
}

func (p_ PlayerLayer) ModelLayer() PlayerLayer {
	rv := objc.Call[PlayerLayer](p_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func PlayerLayer_ModelLayer() PlayerLayer {
	instance := PlayerLayerClass.Alloc().ModelLayer()
	instance.Autorelease()
	return instance
}

func (p_ PlayerLayer) PresentationLayer() PlayerLayer {
	rv := objc.Call[PlayerLayer](p_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func PlayerLayer_PresentationLayer() PlayerLayer {
	instance := PlayerLayerClass.Alloc().PresentationLayer()
	instance.Autorelease()
	return instance
}

// Creates a layer object to present the visual contents of a player’s current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlayer/1389308-playerlayerwithplayer?language=objc
func (pc _PlayerLayerClass) PlayerLayerWithPlayer(player IPlayer) PlayerLayer {
	rv := objc.Call[PlayerLayer](pc, objc.Sel("playerLayerWithPlayer:"), objc.Ptr(player))
	return rv
}

// Creates a layer object to present the visual contents of a player’s current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlayer/1389308-playerlayerwithplayer?language=objc
func PlayerLayer_PlayerLayerWithPlayer(player IPlayer) PlayerLayer {
	return PlayerLayerClass.PlayerLayerWithPlayer(player)
}

// A value that specifies how the layer displays the player’s visual content within the layer’s bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlayer/1388915-videogravity?language=objc
func (p_ PlayerLayer) VideoGravity() LayerVideoGravity {
	rv := objc.Call[LayerVideoGravity](p_, objc.Sel("videoGravity"))
	return rv
}

// A value that specifies how the layer displays the player’s visual content within the layer’s bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlayer/1388915-videogravity?language=objc
func (p_ PlayerLayer) SetVideoGravity(value LayerVideoGravity) {
	objc.Call[objc.Void](p_, objc.Sel("setVideoGravity:"), value)
}

// The attributes of the visual output that displays in the player layer during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlayer/1390055-pixelbufferattributes?language=objc
func (p_ PlayerLayer) PixelBufferAttributes() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](p_, objc.Sel("pixelBufferAttributes"))
	return rv
}

// The attributes of the visual output that displays in the player layer during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlayer/1390055-pixelbufferattributes?language=objc
func (p_ PlayerLayer) SetPixelBufferAttributes(value map[string]objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("setPixelBufferAttributes:"), value)
}

// The player whose visual content the layer displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlayer/1390434-player?language=objc
func (p_ PlayerLayer) Player() Player {
	rv := objc.Call[Player](p_, objc.Sel("player"))
	return rv
}

// The player whose visual content the layer displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlayer/1390434-player?language=objc
func (p_ PlayerLayer) SetPlayer(value IPlayer) {
	objc.Call[objc.Void](p_, objc.Sel("setPlayer:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the first video frame of the player’s current item is ready for display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlayer/1389748-readyfordisplay?language=objc
func (p_ PlayerLayer) IsReadyForDisplay() bool {
	rv := objc.Call[bool](p_, objc.Sel("isReadyForDisplay"))
	return rv
}

// The current size and position of the video image that displays within the layer’s bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlayer/1385745-videorect?language=objc
func (p_ PlayerLayer) VideoRect() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](p_, objc.Sel("videoRect"))
	return rv
}
