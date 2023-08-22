// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [KeyedUnarchiver] class.
var KeyedUnarchiverClass = _KeyedUnarchiverClass{objc.GetClass("MPSKeyedUnarchiver")}

type _KeyedUnarchiverClass struct {
	objc.Class
}

// An interface definition for the [KeyedUnarchiver] class.
type IKeyedUnarchiver interface {
	foundation.IKeyedUnarchiver
	MpsMTLDevice() metal.DeviceWrapper
}

// A keyed archiver that supports Metal Performance Shaders kernel decoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver?language=objc
type KeyedUnarchiver struct {
	foundation.KeyedUnarchiver
}

func KeyedUnarchiverFrom(ptr unsafe.Pointer) KeyedUnarchiver {
	return KeyedUnarchiver{
		KeyedUnarchiver: foundation.KeyedUnarchiverFrom(ptr),
	}
}

func (k_ KeyedUnarchiver) InitForReadingFromDataDeviceError(data []byte, device metal.PDevice, error foundation.IError) KeyedUnarchiver {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[KeyedUnarchiver](k_, objc.Sel("initForReadingFromData:device:error:"), data, po1, objc.Ptr(error))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2966644-initforreadingfromdata?language=objc
func NewKeyedUnarchiverForReadingFromDataDeviceError(data []byte, device metal.PDevice, error foundation.IError) KeyedUnarchiver {
	instance := KeyedUnarchiverClass.Alloc().InitForReadingFromDataDeviceError(data, device, error)
	instance.Autorelease()
	return instance
}

func (kc _KeyedUnarchiverClass) Alloc() KeyedUnarchiver {
	rv := objc.Call[KeyedUnarchiver](kc, objc.Sel("alloc"))
	return rv
}

func KeyedUnarchiver_Alloc() KeyedUnarchiver {
	return KeyedUnarchiverClass.Alloc()
}

func (kc _KeyedUnarchiverClass) New() KeyedUnarchiver {
	rv := objc.Call[KeyedUnarchiver](kc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewKeyedUnarchiver() KeyedUnarchiver {
	return KeyedUnarchiverClass.New()
}

func (k_ KeyedUnarchiver) Init() KeyedUnarchiver {
	rv := objc.Call[KeyedUnarchiver](k_, objc.Sel("init"))
	return rv
}

func (k_ KeyedUnarchiver) InitForReadingFromDataError(data []byte, error foundation.IError) KeyedUnarchiver {
	rv := objc.Call[KeyedUnarchiver](k_, objc.Sel("initForReadingFromData:error:"), data, objc.Ptr(error))
	return rv
}

// Initializes an archiver to decode data from the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/2962883-initforreadingfromdata?language=objc
func NewKeyedUnarchiverForReadingFromDataError(data []byte, error foundation.IError) KeyedUnarchiver {
	instance := KeyedUnarchiverClass.Alloc().InitForReadingFromDataError(data, error)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2976453-unarchivedobjectofclass?language=objc
func (kc _KeyedUnarchiverClass) UnarchivedObjectOfClassFromDataDeviceError(cls objc.IClass, data []byte, device metal.PDevice, error foundation.IError) objc.Object {
	po2 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[objc.Object](kc, objc.Sel("unarchivedObjectOfClass:fromData:device:error:"), objc.Ptr(cls), data, po2, objc.Ptr(error))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2976453-unarchivedobjectofclass?language=objc
func KeyedUnarchiver_UnarchivedObjectOfClassFromDataDeviceError(cls objc.IClass, data []byte, device metal.PDevice, error foundation.IError) objc.Object {
	return KeyedUnarchiverClass.UnarchivedObjectOfClassFromDataDeviceError(cls, data, device, error)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2976453-unarchivedobjectofclass?language=objc
func (kc _KeyedUnarchiverClass) UnarchivedObjectOfClassFromDataDeviceObjectError(cls objc.IClass, data []byte, deviceObject objc.IObject, error foundation.IError) objc.Object {
	rv := objc.Call[objc.Object](kc, objc.Sel("unarchivedObjectOfClass:fromData:device:error:"), objc.Ptr(cls), data, objc.Ptr(deviceObject), objc.Ptr(error))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2976453-unarchivedobjectofclass?language=objc
func KeyedUnarchiver_UnarchivedObjectOfClassFromDataDeviceObjectError(cls objc.IClass, data []byte, deviceObject objc.IObject, error foundation.IError) objc.Object {
	return KeyedUnarchiverClass.UnarchivedObjectOfClassFromDataDeviceObjectError(cls, data, deviceObject, error)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2976454-unarchivedobjectofclasses?language=objc
func (kc _KeyedUnarchiverClass) UnarchivedObjectOfClassesFromDataDeviceError(classes foundation.ISet, data []byte, device metal.PDevice, error foundation.IError) objc.Object {
	po2 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[objc.Object](kc, objc.Sel("unarchivedObjectOfClasses:fromData:device:error:"), objc.Ptr(classes), data, po2, objc.Ptr(error))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2976454-unarchivedobjectofclasses?language=objc
func KeyedUnarchiver_UnarchivedObjectOfClassesFromDataDeviceError(classes foundation.ISet, data []byte, device metal.PDevice, error foundation.IError) objc.Object {
	return KeyedUnarchiverClass.UnarchivedObjectOfClassesFromDataDeviceError(classes, data, device, error)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2976454-unarchivedobjectofclasses?language=objc
func (kc _KeyedUnarchiverClass) UnarchivedObjectOfClassesFromDataDeviceObjectError(classes foundation.ISet, data []byte, deviceObject objc.IObject, error foundation.IError) objc.Object {
	rv := objc.Call[objc.Object](kc, objc.Sel("unarchivedObjectOfClasses:fromData:device:error:"), objc.Ptr(classes), data, objc.Ptr(deviceObject), objc.Ptr(error))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2976454-unarchivedobjectofclasses?language=objc
func KeyedUnarchiver_UnarchivedObjectOfClassesFromDataDeviceObjectError(classes foundation.ISet, data []byte, deviceObject objc.IObject, error foundation.IError) objc.Object {
	return KeyedUnarchiverClass.UnarchivedObjectOfClassesFromDataDeviceObjectError(classes, data, deviceObject, error)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951880-mpsmtldevice?language=objc
func (k_ KeyedUnarchiver) MpsMTLDevice() metal.DeviceWrapper {
	rv := objc.Call[metal.DeviceWrapper](k_, objc.Sel("mpsMTLDevice"))
	return rv
}
