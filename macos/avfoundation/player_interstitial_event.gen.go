// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [PlayerInterstitialEvent] class.
var PlayerInterstitialEventClass = _PlayerInterstitialEventClass{objc.GetClass("AVPlayerInterstitialEvent")}

type _PlayerInterstitialEventClass struct {
	objc.Class
}

// An interface definition for the [PlayerInterstitialEvent] class.
type IPlayerInterstitialEvent interface {
	objc.IObject
	TemplateItems() []PlayerItem
	Time() coremedia.Time
	Date() foundation.Date
	PrimaryItem() PlayerItem
	Identifier() string
	ResumptionOffset() coremedia.Time
	PlayoutLimit() coremedia.Time
	Restrictions() PlayerInterstitialEventRestrictions
	UserDefinedAttributes() foundation.Dictionary
}

// An object that provides instructions for how a player presents interstitial content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent?language=objc
type PlayerInterstitialEvent struct {
	objc.Object
}

func PlayerInterstitialEventFrom(ptr unsafe.Pointer) PlayerInterstitialEvent {
	return PlayerInterstitialEvent{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PlayerInterstitialEventClass) Alloc() PlayerInterstitialEvent {
	rv := objc.Call[PlayerInterstitialEvent](pc, objc.Sel("alloc"))
	return rv
}

func (pc _PlayerInterstitialEventClass) New() PlayerInterstitialEvent {
	rv := objc.Call[PlayerInterstitialEvent](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerInterstitialEvent() PlayerInterstitialEvent {
	return PlayerInterstitialEventClass.New()
}

func (p_ PlayerInterstitialEvent) Init() PlayerInterstitialEvent {
	rv := objc.Call[PlayerInterstitialEvent](p_, objc.Sel("init"))
	return rv
}

// An array of player item configurations to use as templates for player items that play interstitial content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/3726121-templateitems?language=objc
func (p_ PlayerInterstitialEvent) TemplateItems() []PlayerItem {
	rv := objc.Call[[]PlayerItem](p_, objc.Sel("templateItems"))
	return rv
}

// A time within the timeline of the primary content that playback of interstitial content begins. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/3726122-time?language=objc
func (p_ PlayerInterstitialEvent) Time() coremedia.Time {
	rv := objc.Call[coremedia.Time](p_, objc.Sel("time"))
	return rv
}

// A date within the date range of the primary content that playback of interstitial content begins. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/3726112-date?language=objc
func (p_ PlayerInterstitialEvent) Date() foundation.Date {
	rv := objc.Call[foundation.Date](p_, objc.Sel("date"))
	return rv
}

// The player item that represents the primary content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/3726118-primaryitem?language=objc
func (p_ PlayerInterstitialEvent) PrimaryItem() PlayerItem {
	rv := objc.Call[PlayerItem](p_, objc.Sel("primaryItem"))
	return rv
}

// An identifier for the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/3746585-identifier?language=objc
func (p_ PlayerInterstitialEvent) Identifier() string {
	rv := objc.Call[string](p_, objc.Sel("identifier"))
	return rv
}

// A time offset at which playback of primary content resumes after interstitial content finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/3726120-resumptionoffset?language=objc
func (p_ PlayerInterstitialEvent) ResumptionOffset() coremedia.Time {
	rv := objc.Call[coremedia.Time](p_, objc.Sel("resumptionOffset"))
	return rv
}

// The time offset at which playback of the interstitial ends. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/3746588-playoutlimit?language=objc
func (p_ PlayerInterstitialEvent) PlayoutLimit() coremedia.Time {
	rv := objc.Call[coremedia.Time](p_, objc.Sel("playoutLimit"))
	return rv
}

// The restrictions the event imposes on the playback of interstitial content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/3726119-restrictions?language=objc
func (p_ PlayerInterstitialEvent) Restrictions() PlayerInterstitialEventRestrictions {
	rv := objc.Call[PlayerInterstitialEventRestrictions](p_, objc.Sel("restrictions"))
	return rv
}

// Attributes of the event that the vendor or app defines. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/3820994-userdefinedattributes?language=objc
func (p_ PlayerInterstitialEvent) UserDefinedAttributes() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](p_, objc.Sel("userDefinedAttributes"))
	return rv
}
