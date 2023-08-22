// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerInterstitialEventController] class.
var PlayerInterstitialEventControllerClass = _PlayerInterstitialEventControllerClass{objc.GetClass("AVPlayerInterstitialEventController")}

type _PlayerInterstitialEventControllerClass struct {
	objc.Class
}

// An interface definition for the [PlayerInterstitialEventController] class.
type IPlayerInterstitialEventController interface {
	IPlayerInterstitialEventMonitor
	CancelCurrentEventWithResumptionOffset(resumptionOffset coremedia.Time)
	SetEvents(value []IPlayerInterstitialEvent)
}

// An object that schedules interstitial events for items played by the primary player. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventcontroller?language=objc
type PlayerInterstitialEventController struct {
	PlayerInterstitialEventMonitor
}

func PlayerInterstitialEventControllerFrom(ptr unsafe.Pointer) PlayerInterstitialEventController {
	return PlayerInterstitialEventController{
		PlayerInterstitialEventMonitor: PlayerInterstitialEventMonitorFrom(ptr),
	}
}

func (p_ PlayerInterstitialEventController) InitWithPrimaryPlayer(primaryPlayer IPlayer) PlayerInterstitialEventController {
	rv := objc.Call[PlayerInterstitialEventController](p_, objc.Sel("initWithPrimaryPlayer:"), objc.Ptr(primaryPlayer))
	return rv
}

// Creates an event controller with a player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventcontroller/3726126-initwithprimaryplayer?language=objc
func NewPlayerInterstitialEventControllerWithPrimaryPlayer(primaryPlayer IPlayer) PlayerInterstitialEventController {
	instance := PlayerInterstitialEventControllerClass.Alloc().InitWithPrimaryPlayer(primaryPlayer)
	instance.Autorelease()
	return instance
}

func (pc _PlayerInterstitialEventControllerClass) InterstitialEventControllerWithPrimaryPlayer(primaryPlayer IPlayer) PlayerInterstitialEventController {
	rv := objc.Call[PlayerInterstitialEventController](pc, objc.Sel("interstitialEventControllerWithPrimaryPlayer:"), objc.Ptr(primaryPlayer))
	return rv
}

// A convenience initializer that creates an event controller with a player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventcontroller/3726127-interstitialeventcontrollerwithp?language=objc
func PlayerInterstitialEventController_InterstitialEventControllerWithPrimaryPlayer(primaryPlayer IPlayer) PlayerInterstitialEventController {
	return PlayerInterstitialEventControllerClass.InterstitialEventControllerWithPrimaryPlayer(primaryPlayer)
}

func (pc _PlayerInterstitialEventControllerClass) Alloc() PlayerInterstitialEventController {
	rv := objc.Call[PlayerInterstitialEventController](pc, objc.Sel("alloc"))
	return rv
}

func PlayerInterstitialEventController_Alloc() PlayerInterstitialEventController {
	return PlayerInterstitialEventControllerClass.Alloc()
}

func (pc _PlayerInterstitialEventControllerClass) New() PlayerInterstitialEventController {
	rv := objc.Call[PlayerInterstitialEventController](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerInterstitialEventController() PlayerInterstitialEventController {
	return PlayerInterstitialEventControllerClass.New()
}

func (p_ PlayerInterstitialEventController) Init() PlayerInterstitialEventController {
	rv := objc.Call[PlayerInterstitialEventController](p_, objc.Sel("init"))
	return rv
}

func (pc _PlayerInterstitialEventControllerClass) InterstitialEventMonitorWithPrimaryPlayer(primaryPlayer IPlayer) PlayerInterstitialEventController {
	rv := objc.Call[PlayerInterstitialEventController](pc, objc.Sel("interstitialEventMonitorWithPrimaryPlayer:"), objc.Ptr(primaryPlayer))
	return rv
}

// A convenience initializer that creates an observer with a player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/3800566-interstitialeventmonitorwithprim?language=objc
func PlayerInterstitialEventController_InterstitialEventMonitorWithPrimaryPlayer(primaryPlayer IPlayer) PlayerInterstitialEventController {
	return PlayerInterstitialEventControllerClass.InterstitialEventMonitorWithPrimaryPlayer(primaryPlayer)
}

// Cancels the playback of all currently playing and scheduled interstitial events, and resumes playback of primary content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventcontroller/3726124-cancelcurrenteventwithresumption?language=objc
func (p_ PlayerInterstitialEventController) CancelCurrentEventWithResumptionOffset(resumptionOffset coremedia.Time) {
	objc.Call[objc.Void](p_, objc.Sel("cancelCurrentEventWithResumptionOffset:"), resumptionOffset)
}

// The current schedule of interstitial events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventcontroller/3726125-events?language=objc
func (p_ PlayerInterstitialEventController) SetEvents(value []IPlayerInterstitialEvent) {
	objc.Call[objc.Void](p_, objc.Sel("setEvents:"), value)
}
