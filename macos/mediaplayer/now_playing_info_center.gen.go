// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NowPlayingInfoCenter] class.
var NowPlayingInfoCenterClass = _NowPlayingInfoCenterClass{objc.GetClass("MPNowPlayingInfoCenter")}

type _NowPlayingInfoCenterClass struct {
	objc.Class
}

// An interface definition for the [NowPlayingInfoCenter] class.
type INowPlayingInfoCenter interface {
	objc.IObject
	PlaybackState() NowPlayingPlaybackState
	SetPlaybackState(value NowPlayingPlaybackState)
	NowPlayingInfo() map[string]objc.Object
	SetNowPlayingInfo(value map[string]objc.IObject)
}

// An object for setting the Now Playing information for media that your app plays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfocenter?language=objc
type NowPlayingInfoCenter struct {
	objc.Object
}

func NowPlayingInfoCenterFrom(ptr unsafe.Pointer) NowPlayingInfoCenter {
	return NowPlayingInfoCenter{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NowPlayingInfoCenterClass) Alloc() NowPlayingInfoCenter {
	rv := objc.Call[NowPlayingInfoCenter](nc, objc.Sel("alloc"))
	return rv
}

func NowPlayingInfoCenter_Alloc() NowPlayingInfoCenter {
	return NowPlayingInfoCenterClass.Alloc()
}

func (nc _NowPlayingInfoCenterClass) New() NowPlayingInfoCenter {
	rv := objc.Call[NowPlayingInfoCenter](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNowPlayingInfoCenter() NowPlayingInfoCenter {
	return NowPlayingInfoCenterClass.New()
}

func (n_ NowPlayingInfoCenter) Init() NowPlayingInfoCenter {
	rv := objc.Call[NowPlayingInfoCenter](n_, objc.Sel("init"))
	return rv
}

// Returns the singleton Now Playing info center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfocenter/1615899-defaultcenter?language=objc
func (nc _NowPlayingInfoCenterClass) DefaultCenter() NowPlayingInfoCenter {
	rv := objc.Call[NowPlayingInfoCenter](nc, objc.Sel("defaultCenter"))
	return rv
}

// Returns the singleton Now Playing info center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfocenter/1615899-defaultcenter?language=objc
func NowPlayingInfoCenter_DefaultCenter() NowPlayingInfoCenter {
	return NowPlayingInfoCenterClass.DefaultCenter()
}

// The current playback state of the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfocenter/2588243-playbackstate?language=objc
func (n_ NowPlayingInfoCenter) PlaybackState() NowPlayingPlaybackState {
	rv := objc.Call[NowPlayingPlaybackState](n_, objc.Sel("playbackState"))
	return rv
}

// The current playback state of the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfocenter/2588243-playbackstate?language=objc
func (n_ NowPlayingInfoCenter) SetPlaybackState(value NowPlayingPlaybackState) {
	objc.Call[objc.Void](n_, objc.Sel("setPlaybackState:"), value)
}

// The current Now Playing information for the default Now Playing info center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfocenter/1615903-nowplayinginfo?language=objc
func (n_ NowPlayingInfoCenter) NowPlayingInfo() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](n_, objc.Sel("nowPlayingInfo"))
	return rv
}

// The current Now Playing information for the default Now Playing info center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfocenter/1615903-nowplayinginfo?language=objc
func (n_ NowPlayingInfoCenter) SetNowPlayingInfo(value map[string]objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setNowPlayingInfo:"), value)
}
