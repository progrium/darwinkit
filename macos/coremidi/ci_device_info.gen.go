// AUTO-GENERATED CODE, DO NOT MODIFY

package coremidi

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CIDeviceInfo] class.
var CIDeviceInfoClass = _CIDeviceInfoClass{objc.GetClass("MIDICIDeviceInfo")}

type _CIDeviceInfoClass struct {
	objc.Class
}

// An interface definition for the [CIDeviceInfo] class.
type ICIDeviceInfo interface {
	objc.IObject
	ManufacturerID() []byte
	RevisionLevel() []byte
	Family() []byte
	MidiDestination() EndpointRef
	ModelNumber() []byte
}

// An object that provides basic information about a MIDI-CI device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo?language=objc
type CIDeviceInfo struct {
	objc.Object
}

func CIDeviceInfoFrom(ptr unsafe.Pointer) CIDeviceInfo {
	return CIDeviceInfo{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ CIDeviceInfo) InitWithDestinationManufacturerFamilyModelRevision(midiDestination EntityRef, manufacturer []byte, family []byte, modelNumber []byte, revisionLevel []byte) CIDeviceInfo {
	rv := objc.Call[CIDeviceInfo](c_, objc.Sel("initWithDestination:manufacturer:family:model:revision:"), midiDestination, manufacturer, family, modelNumber, revisionLevel)
	return rv
}

// Creates a new device information instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/3580317-initwithdestination?language=objc
func NewCIDeviceInfoWithDestinationManufacturerFamilyModelRevision(midiDestination EntityRef, manufacturer []byte, family []byte, modelNumber []byte, revisionLevel []byte) CIDeviceInfo {
	instance := CIDeviceInfoClass.Alloc().InitWithDestinationManufacturerFamilyModelRevision(midiDestination, manufacturer, family, modelNumber, revisionLevel)
	instance.Autorelease()
	return instance
}

func (cc _CIDeviceInfoClass) Alloc() CIDeviceInfo {
	rv := objc.Call[CIDeviceInfo](cc, objc.Sel("alloc"))
	return rv
}

func CIDeviceInfo_Alloc() CIDeviceInfo {
	return CIDeviceInfoClass.Alloc()
}

func (cc _CIDeviceInfoClass) New() CIDeviceInfo {
	rv := objc.Call[CIDeviceInfo](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCIDeviceInfo() CIDeviceInfo {
	return CIDeviceInfoClass.New()
}

func (c_ CIDeviceInfo) Init() CIDeviceInfo {
	rv := objc.Call[CIDeviceInfo](c_, objc.Sel("init"))
	return rv
}

// The MIDI System Exclusive (SysEx) ID of the device manufacturer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/3553237-manufacturerid?language=objc
func (c_ CIDeviceInfo) ManufacturerID() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("manufacturerID"))
	return rv
}

// The revision number of the device model number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/3553239-revisionlevel?language=objc
func (c_ CIDeviceInfo) RevisionLevel() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("revisionLevel"))
	return rv
}

// The family to which the device belongs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/3553235-family?language=objc
func (c_ CIDeviceInfo) Family() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("family"))
	return rv
}

// The MIDI destination the device’s MIDI entity uses for capability inquiries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/3580318-mididestination?language=objc
func (c_ CIDeviceInfo) MidiDestination() EndpointRef {
	rv := objc.Call[EndpointRef](c_, objc.Sel("midiDestination"))
	return rv
}

// The model number of the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/3553238-modelnumber?language=objc
func (c_ CIDeviceInfo) ModelNumber() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("modelNumber"))
	return rv
}
