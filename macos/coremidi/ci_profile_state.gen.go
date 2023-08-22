// AUTO-GENERATED CODE, DO NOT MODIFY

package coremidi

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CIProfileState] class.
var CIProfileStateClass = _CIProfileStateClass{objc.GetClass("MIDICIProfileState")}

type _CIProfileStateClass struct {
	objc.Class
}

// An interface definition for the [CIProfileState] class.
type ICIProfileState interface {
	objc.IObject
	MidiChannel() ChannelNumber
	EnabledProfiles() []CIProfile
	DisabledProfiles() []CIProfile
}

// An object that provides the enabled and disabled profiles for a MIDI channel or port on a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofilestate?language=objc
type CIProfileState struct {
	objc.Object
}

func CIProfileStateFrom(ptr unsafe.Pointer) CIProfileState {
	return CIProfileState{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ CIProfileState) InitWithChannelEnabledProfilesDisabledProfiles(midiChannelNum ChannelNumber, enabled []ICIProfile, disabled []ICIProfile) CIProfileState {
	rv := objc.Call[CIProfileState](c_, objc.Sel("initWithChannel:enabledProfiles:disabledProfiles:"), midiChannelNum, enabled, disabled)
	return rv
}

// Creates a new profile state object for the specified MIDI channel and profiles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofilestate/3553245-initwithchannel?language=objc
func NewCIProfileStateWithChannelEnabledProfilesDisabledProfiles(midiChannelNum ChannelNumber, enabled []ICIProfile, disabled []ICIProfile) CIProfileState {
	instance := CIProfileStateClass.Alloc().InitWithChannelEnabledProfilesDisabledProfiles(midiChannelNum, enabled, disabled)
	instance.Autorelease()
	return instance
}

func (cc _CIProfileStateClass) Alloc() CIProfileState {
	rv := objc.Call[CIProfileState](cc, objc.Sel("alloc"))
	return rv
}

func CIProfileState_Alloc() CIProfileState {
	return CIProfileStateClass.Alloc()
}

func (cc _CIProfileStateClass) New() CIProfileState {
	rv := objc.Call[CIProfileState](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCIProfileState() CIProfileState {
	return CIProfileStateClass.New()
}

func (c_ CIProfileState) Init() CIProfileState {
	rv := objc.Call[CIProfileState](c_, objc.Sel("init"))
	return rv
}

// The MIDI channel to which this state applies. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofilestate/3553246-midichannel?language=objc
func (c_ CIProfileState) MidiChannel() ChannelNumber {
	rv := objc.Call[ChannelNumber](c_, objc.Sel("midiChannel"))
	return rv
}

// The object’s enabled profiles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofilestate/2977110-enabledprofiles?language=objc
func (c_ CIProfileState) EnabledProfiles() []CIProfile {
	rv := objc.Call[[]CIProfile](c_, objc.Sel("enabledProfiles"))
	return rv
}

// The object’s disabled profiles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofilestate/2977109-disabledprofiles?language=objc
func (c_ CIProfileState) DisabledProfiles() []CIProfile {
	rv := objc.Call[[]CIProfile](c_, objc.Sel("disabledProfiles"))
	return rv
}
