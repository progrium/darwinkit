// AUTO-GENERATED CODE, DO NOT MODIFY

package coremidi

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CIResponder] class.
var CIResponderClass = _CIResponderClass{objc.GetClass("MIDICIResponder")}

type _CIResponderClass struct {
	objc.Class
}

// An interface definition for the [CIResponder] class.
type ICIResponder interface {
	objc.IObject
	NotifyProfileOnChannelIsEnabled(aProfile ICIProfile, channel ChannelNumber, enabledState bool) bool
	SendProfileOnChannelProfileData(aProfile ICIProfile, channel ChannelNumber, profileSpecificData []byte) bool
	Start() bool
	Stop()
	Initiators() []CIInitiatiorMUID
	DeviceInfo() CIDeviceInfo
	ProfileDelegate() CIProfileResponderDelegateWrapper
}

// An object that responds to MIDI-CI inquiries from an initiator on behalf of a MIDI client, and handles profile and property exchange operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder?language=objc
type CIResponder struct {
	objc.Object
}

func CIResponderFrom(ptr unsafe.Pointer) CIResponder {
	return CIResponder{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ CIResponder) InitWithDeviceInfoProfileDelegateProfileStatesSupportProperties(deviceInfo ICIDeviceInfo, delegate PCIProfileResponderDelegate, profileList *foundation.Array, propertiesSupported bool) CIResponder {
	po1 := objc.WrapAsProtocol("MIDICIProfileResponderDelegate", delegate)
	rv := objc.Call[CIResponder](c_, objc.Sel("initWithDeviceInfo:profileDelegate:profileStates:supportProperties:"), objc.Ptr(deviceInfo), po1, profileList, propertiesSupported)
	return rv
}

// Creates a new responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/3580332-initwithdeviceinfo?language=objc
func NewCIResponderWithDeviceInfoProfileDelegateProfileStatesSupportProperties(deviceInfo ICIDeviceInfo, delegate PCIProfileResponderDelegate, profileList *foundation.Array, propertiesSupported bool) CIResponder {
	instance := CIResponderClass.Alloc().InitWithDeviceInfoProfileDelegateProfileStatesSupportProperties(deviceInfo, delegate, profileList, propertiesSupported)
	instance.Autorelease()
	return instance
}

func (cc _CIResponderClass) Alloc() CIResponder {
	rv := objc.Call[CIResponder](cc, objc.Sel("alloc"))
	return rv
}

func CIResponder_Alloc() CIResponder {
	return CIResponderClass.Alloc()
}

func (cc _CIResponderClass) New() CIResponder {
	rv := objc.Call[CIResponder](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCIResponder() CIResponder {
	return CIResponderClass.New()
}

func (c_ CIResponder) Init() CIResponder {
	rv := objc.Call[CIResponder](c_, objc.Sel("init"))
	return rv
}

// Enables or disables a profile and notifies all connected initiators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/3553257-notifyprofile?language=objc
func (c_ CIResponder) NotifyProfileOnChannelIsEnabled(aProfile ICIProfile, channel ChannelNumber, enabledState bool) bool {
	rv := objc.Call[bool](c_, objc.Sel("notifyProfile:onChannel:isEnabled:"), objc.Ptr(aProfile), channel, enabledState)
	return rv
}

// Sends profile-specific data to all connected initiators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/3553259-sendprofile?language=objc
func (c_ CIResponder) SendProfileOnChannelProfileData(aProfile ICIProfile, channel ChannelNumber, profileSpecificData []byte) bool {
	rv := objc.Call[bool](c_, objc.Sel("sendProfile:onChannel:profileData:"), objc.Ptr(aProfile), channel, profileSpecificData)
	return rv
}

// Starts receiving initiator requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/3553260-start?language=objc
func (c_ CIResponder) Start() bool {
	rv := objc.Call[bool](c_, objc.Sel("start"))
	return rv
}

// Stops receiving initiator requests and disconnects all connected initiators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/3553261-stop?language=objc
func (c_ CIResponder) Stop() {
	objc.Call[objc.Void](c_, objc.Sel("stop"))
}

// An array of initiators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/3580333-initiators?language=objc
func (c_ CIResponder) Initiators() []CIInitiatiorMUID {
	rv := objc.Call[[]CIInitiatiorMUID](c_, objc.Sel("initiators"))
	return rv
}

// The MIDI-CI device’s information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/3553255-deviceinfo?language=objc
func (c_ CIResponder) DeviceInfo() CIDeviceInfo {
	rv := objc.Call[CIDeviceInfo](c_, objc.Sel("deviceInfo"))
	return rv
}

// The profile delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/3580334-profiledelegate?language=objc
func (c_ CIResponder) ProfileDelegate() CIProfileResponderDelegateWrapper {
	rv := objc.Call[CIProfileResponderDelegateWrapper](c_, objc.Sel("profileDelegate"))
	return rv
}
