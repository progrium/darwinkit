// AUTO-GENERATED CODE, DO NOT MODIFY

package coremidi

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines the methods to respond to MIDI-CI responder life-cycle events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofileresponderdelegate?language=objc
type PCIProfileResponderDelegate interface {
	// optional
	WillSetProfileOnChannelEnabled(aProfile CIProfile, channel ChannelNumber, shouldEnable bool) bool
	HasWillSetProfileOnChannelEnabled() bool

	// optional
	InitiatorDisconnected(initiatorMUID CIInitiatiorMUID)
	HasInitiatorDisconnected() bool

	// optional
	HandleDataForProfileOnChannelData(aProfile CIProfile, channel ChannelNumber, inData []byte)
	HasHandleDataForProfileOnChannelData() bool

	// optional
	ConnectInitiatorWithDeviceInfo(initiatorMUID CIInitiatiorMUID, deviceInfo CIDeviceInfo) bool
	HasConnectInitiatorWithDeviceInfo() bool
}

// A delegate implementation builder for the [PCIProfileResponderDelegate] protocol.
type CIProfileResponderDelegate struct {
	_WillSetProfileOnChannelEnabled    func(aProfile CIProfile, channel ChannelNumber, shouldEnable bool) bool
	_InitiatorDisconnected             func(initiatorMUID CIInitiatiorMUID)
	_HandleDataForProfileOnChannelData func(aProfile CIProfile, channel ChannelNumber, inData []byte)
	_ConnectInitiatorWithDeviceInfo    func(initiatorMUID CIInitiatiorMUID, deviceInfo CIDeviceInfo) bool
}

func (di *CIProfileResponderDelegate) HasWillSetProfileOnChannelEnabled() bool {
	return di._WillSetProfileOnChannelEnabled != nil
}

// Provides an opportunity to perform an action before the system sets the profile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofileresponderdelegate/3580331-willsetprofile?language=objc
func (di *CIProfileResponderDelegate) SetWillSetProfileOnChannelEnabled(f func(aProfile CIProfile, channel ChannelNumber, shouldEnable bool) bool) {
	di._WillSetProfileOnChannelEnabled = f
}

// Provides an opportunity to perform an action before the system sets the profile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofileresponderdelegate/3580331-willsetprofile?language=objc
func (di *CIProfileResponderDelegate) WillSetProfileOnChannelEnabled(aProfile CIProfile, channel ChannelNumber, shouldEnable bool) bool {
	return di._WillSetProfileOnChannelEnabled(aProfile, channel, shouldEnable)
}
func (di *CIProfileResponderDelegate) HasInitiatorDisconnected() bool {
	return di._InitiatorDisconnected != nil
}

// Provides an opportunity to perform an action after the system disconnects the initiator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofileresponderdelegate/3580330-initiatordisconnected?language=objc
func (di *CIProfileResponderDelegate) SetInitiatorDisconnected(f func(initiatorMUID CIInitiatiorMUID)) {
	di._InitiatorDisconnected = f
}

// Provides an opportunity to perform an action after the system disconnects the initiator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofileresponderdelegate/3580330-initiatordisconnected?language=objc
func (di *CIProfileResponderDelegate) InitiatorDisconnected(initiatorMUID CIInitiatiorMUID) {
	di._InitiatorDisconnected(initiatorMUID)
}
func (di *CIProfileResponderDelegate) HasHandleDataForProfileOnChannelData() bool {
	return di._HandleDataForProfileOnChannelData != nil
}

// Processes MIDI data for a profile and channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofileresponderdelegate/3580329-handledataforprofile?language=objc
func (di *CIProfileResponderDelegate) SetHandleDataForProfileOnChannelData(f func(aProfile CIProfile, channel ChannelNumber, inData []byte)) {
	di._HandleDataForProfileOnChannelData = f
}

// Processes MIDI data for a profile and channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofileresponderdelegate/3580329-handledataforprofile?language=objc
func (di *CIProfileResponderDelegate) HandleDataForProfileOnChannelData(aProfile CIProfile, channel ChannelNumber, inData []byte) {
	di._HandleDataForProfileOnChannelData(aProfile, channel, inData)
}
func (di *CIProfileResponderDelegate) HasConnectInitiatorWithDeviceInfo() bool {
	return di._ConnectInitiatorWithDeviceInfo != nil
}

// Enables a MIDI-CI initiator to create a session or reject the connection attempt. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofileresponderdelegate/3580328-connectinitiator?language=objc
func (di *CIProfileResponderDelegate) SetConnectInitiatorWithDeviceInfo(f func(initiatorMUID CIInitiatiorMUID, deviceInfo CIDeviceInfo) bool) {
	di._ConnectInitiatorWithDeviceInfo = f
}

// Enables a MIDI-CI initiator to create a session or reject the connection attempt. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofileresponderdelegate/3580328-connectinitiator?language=objc
func (di *CIProfileResponderDelegate) ConnectInitiatorWithDeviceInfo(initiatorMUID CIInitiatiorMUID, deviceInfo CIDeviceInfo) bool {
	return di._ConnectInitiatorWithDeviceInfo(initiatorMUID, deviceInfo)
}

// A concrete type wrapper for the [PCIProfileResponderDelegate] protocol.
type CIProfileResponderDelegateWrapper struct {
	objc.Object
}

func (c_ CIProfileResponderDelegateWrapper) HasWillSetProfileOnChannelEnabled() bool {
	return c_.RespondsToSelector(objc.Sel("willSetProfile:onChannel:enabled:"))
}

// Provides an opportunity to perform an action before the system sets the profile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofileresponderdelegate/3580331-willsetprofile?language=objc
func (c_ CIProfileResponderDelegateWrapper) WillSetProfileOnChannelEnabled(aProfile ICIProfile, channel ChannelNumber, shouldEnable bool) bool {
	rv := objc.Call[bool](c_, objc.Sel("willSetProfile:onChannel:enabled:"), objc.Ptr(aProfile), channel, shouldEnable)
	return rv
}

func (c_ CIProfileResponderDelegateWrapper) HasInitiatorDisconnected() bool {
	return c_.RespondsToSelector(objc.Sel("initiatorDisconnected:"))
}

// Provides an opportunity to perform an action after the system disconnects the initiator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofileresponderdelegate/3580330-initiatordisconnected?language=objc
func (c_ CIProfileResponderDelegateWrapper) InitiatorDisconnected(initiatorMUID CIInitiatiorMUID) {
	objc.Call[objc.Void](c_, objc.Sel("initiatorDisconnected:"), initiatorMUID)
}

func (c_ CIProfileResponderDelegateWrapper) HasHandleDataForProfileOnChannelData() bool {
	return c_.RespondsToSelector(objc.Sel("handleDataForProfile:onChannel:data:"))
}

// Processes MIDI data for a profile and channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofileresponderdelegate/3580329-handledataforprofile?language=objc
func (c_ CIProfileResponderDelegateWrapper) HandleDataForProfileOnChannelData(aProfile ICIProfile, channel ChannelNumber, inData []byte) {
	objc.Call[objc.Void](c_, objc.Sel("handleDataForProfile:onChannel:data:"), objc.Ptr(aProfile), channel, inData)
}

func (c_ CIProfileResponderDelegateWrapper) HasConnectInitiatorWithDeviceInfo() bool {
	return c_.RespondsToSelector(objc.Sel("connectInitiator:withDeviceInfo:"))
}

// Enables a MIDI-CI initiator to create a session or reject the connection attempt. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofileresponderdelegate/3580328-connectinitiator?language=objc
func (c_ CIProfileResponderDelegateWrapper) ConnectInitiatorWithDeviceInfo(initiatorMUID CIInitiatiorMUID, deviceInfo ICIDeviceInfo) bool {
	rv := objc.Call[bool](c_, objc.Sel("connectInitiator:withDeviceInfo:"), initiatorMUID, objc.Ptr(deviceInfo))
	return rv
}
