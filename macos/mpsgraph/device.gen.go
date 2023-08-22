// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Device] class.
var DeviceClass = _DeviceClass{objc.GetClass("MPSGraphDevice")}

type _DeviceClass struct {
	objc.Class
}

// An interface definition for the [Device] class.
type IDevice interface {
	objc.IObject
	MetalDevice() metal.DeviceWrapper
	Type() DeviceType
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdevice?language=objc
type Device struct {
	objc.Object
}

func DeviceFrom(ptr unsafe.Pointer) Device {
	return Device{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DeviceClass) DeviceWithMTLDevice(metalDevice metal.PDevice) Device {
	po0 := objc.WrapAsProtocol("MTLDevice", metalDevice)
	rv := objc.Call[Device](dc, objc.Sel("deviceWithMTLDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdevice/3564649-devicewithmtldevice?language=objc
func Device_DeviceWithMTLDevice(metalDevice metal.PDevice) Device {
	return DeviceClass.DeviceWithMTLDevice(metalDevice)
}

func (dc _DeviceClass) Alloc() Device {
	rv := objc.Call[Device](dc, objc.Sel("alloc"))
	return rv
}

func Device_Alloc() Device {
	return DeviceClass.Alloc()
}

func (dc _DeviceClass) New() Device {
	rv := objc.Call[Device](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDevice() Device {
	return DeviceClass.New()
}

func (d_ Device) Init() Device {
	rv := objc.Call[Device](d_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdevice/3564650-metaldevice?language=objc
func (d_ Device) MetalDevice() metal.DeviceWrapper {
	rv := objc.Call[metal.DeviceWrapper](d_, objc.Sel("metalDevice"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdevice/3564651-type?language=objc
func (d_ Device) Type() DeviceType {
	rv := objc.Call[DeviceType](d_, objc.Sel("type"))
	return rv
}
