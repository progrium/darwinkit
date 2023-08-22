// AUTO-GENERATED CODE, DO NOT MODIFY

package coreaudio

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coreaudiotypes"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiostreampropertylistenerproc?language=objc
type StreamPropertyListenerProc = func(inStream StreamID, inChannel uint32, inPropertyID DevicePropertyID, inClientData unsafe.Pointer) uint

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiohardwarepropertylistenerproc?language=objc
type HardwarePropertyListenerProc = func(inPropertyID HardwarePropertyID, inClientData unsafe.Pointer) uint

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audioobjectpropertylistenerproc?language=objc
type ObjectPropertyListenerProc = func(inObjectID ObjectID, inNumberAddresses uint32, inAddresses *ObjectPropertyAddress, inClientData unsafe.Pointer) uint

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiodriverplugindevicepropertychangedproc?language=objc
type DriverPlugInDevicePropertyChangedProc = func(inDevice DeviceID, inChannel uint32, isInput bool, inPropertyID DevicePropertyID) uint

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiodeviceioblock?language=objc
type DeviceIOBlock = func(inNow *coreaudiotypes.TimeStamp, inInputData *coreaudiotypes.BufferList, inInputTime *coreaudiotypes.TimeStamp, outOutputData *coreaudiotypes.BufferList, inOutputTime *coreaudiotypes.TimeStamp)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiodriverpluginstreampropertychangedproc?language=objc
type DriverPlugInStreamPropertyChangedProc = func(inDevice DeviceID, inIOAudioStream unsafe.Pointer, inChannel uint32, inPropertyID DevicePropertyID) uint

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiodeviceioproc?language=objc
type DeviceIOProc = func(inDevice ObjectID, inNow *coreaudiotypes.TimeStamp, inInputData *coreaudiotypes.BufferList, inInputTime *coreaudiotypes.TimeStamp, outOutputData *coreaudiotypes.BufferList, inOutputTime *coreaudiotypes.TimeStamp, inClientData unsafe.Pointer) uint

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audioobjectpropertylistenerblock?language=objc
type ObjectPropertyListenerBlock = func(inNumberAddresses uint32, inAddresses *ObjectPropertyAddress)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/audiodevicepropertylistenerproc?language=objc
type DevicePropertyListenerProc = func(inDevice DeviceID, inChannel uint32, isInput bool, inPropertyID DevicePropertyID, inClientData unsafe.Pointer) uint
