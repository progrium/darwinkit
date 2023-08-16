// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [HeapDescriptor] class.
var HeapDescriptorClass = _HeapDescriptorClass{objc.GetClass("MTLHeapDescriptor")}

type _HeapDescriptorClass struct {
	objc.Class
}

// An interface definition for the [HeapDescriptor] class.
type IHeapDescriptor interface {
	objc.IObject
	CpuCacheMode() CPUCacheMode
	SetCpuCacheMode(value CPUCacheMode)
	StorageMode() StorageMode
	SetStorageMode(value StorageMode)
	ResourceOptions() ResourceOptions
	SetResourceOptions(value ResourceOptions)
	Type() HeapType
	SetType(value HeapType)
	Size() uint
	SetSize(value uint)
	HazardTrackingMode() HazardTrackingMode
	SetHazardTrackingMode(value HazardTrackingMode)
}

// A configuration that customizes the behavior for a Metal memory heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheapdescriptor?language=objc
type HeapDescriptor struct {
	objc.Object
}

func HeapDescriptorFrom(ptr unsafe.Pointer) HeapDescriptor {
	return HeapDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (hc _HeapDescriptorClass) Alloc() HeapDescriptor {
	rv := objc.Call[HeapDescriptor](hc, objc.Sel("alloc"))
	return rv
}

func HeapDescriptor_Alloc() HeapDescriptor {
	return HeapDescriptorClass.Alloc()
}

func (hc _HeapDescriptorClass) New() HeapDescriptor {
	rv := objc.Call[HeapDescriptor](hc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewHeapDescriptor() HeapDescriptor {
	return HeapDescriptorClass.New()
}

func (h_ HeapDescriptor) Init() HeapDescriptor {
	rv := objc.Call[HeapDescriptor](h_, objc.Sel("init"))
	return rv
}

// The CPU cache behavior for any resources you allocate from the heaps you create with this descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheapdescriptor/1649573-cpucachemode?language=objc
func (h_ HeapDescriptor) CpuCacheMode() CPUCacheMode {
	rv := objc.Call[CPUCacheMode](h_, objc.Sel("cpuCacheMode"))
	return rv
}

// The CPU cache behavior for any resources you allocate from the heaps you create with this descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheapdescriptor/1649573-cpucachemode?language=objc
func (h_ HeapDescriptor) SetCpuCacheMode(value CPUCacheMode) {
	objc.Call[objc.Void](h_, objc.Sel("setCpuCacheMode:"), value)
}

// The storage mode for the heaps you create with this descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheapdescriptor/1649567-storagemode?language=objc
func (h_ HeapDescriptor) StorageMode() StorageMode {
	rv := objc.Call[StorageMode](h_, objc.Sel("storageMode"))
	return rv
}

// The storage mode for the heaps you create with this descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheapdescriptor/1649567-storagemode?language=objc
func (h_ HeapDescriptor) SetStorageMode(value StorageMode) {
	objc.Call[objc.Void](h_, objc.Sel("setStorageMode:"), value)
}

// The combined behavior for any resources you allocate from the heaps you create with this descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheapdescriptor/3131687-resourceoptions?language=objc
func (h_ HeapDescriptor) ResourceOptions() ResourceOptions {
	rv := objc.Call[ResourceOptions](h_, objc.Sel("resourceOptions"))
	return rv
}

// The combined behavior for any resources you allocate from the heaps you create with this descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheapdescriptor/3131687-resourceoptions?language=objc
func (h_ HeapDescriptor) SetResourceOptions(value ResourceOptions) {
	objc.Call[objc.Void](h_, objc.Sel("setResourceOptions:"), value)
}

// The memory placement strategy for any resources you allocate from the heaps you create with this descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheapdescriptor/3043389-type?language=objc
func (h_ HeapDescriptor) Type() HeapType {
	rv := objc.Call[HeapType](h_, objc.Sel("type"))
	return rv
}

// The memory placement strategy for any resources you allocate from the heaps you create with this descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheapdescriptor/3043389-type?language=objc
func (h_ HeapDescriptor) SetType(value HeapType) {
	objc.Call[objc.Void](h_, objc.Sel("setType:"), value)
}

// The total amount of memory, in bytes, for the heaps you create with this descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheapdescriptor/1649568-size?language=objc
func (h_ HeapDescriptor) Size() uint {
	rv := objc.Call[uint](h_, objc.Sel("size"))
	return rv
}

// The total amount of memory, in bytes, for the heaps you create with this descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheapdescriptor/1649568-size?language=objc
func (h_ HeapDescriptor) SetSize(value uint) {
	objc.Call[objc.Void](h_, objc.Sel("setSize:"), value)
}

// The hazard tracking behavior for any resources you allocate from the heaps you create with this descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheapdescriptor/3131686-hazardtrackingmode?language=objc
func (h_ HeapDescriptor) HazardTrackingMode() HazardTrackingMode {
	rv := objc.Call[HazardTrackingMode](h_, objc.Sel("hazardTrackingMode"))
	return rv
}

// The hazard tracking behavior for any resources you allocate from the heaps you create with this descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheapdescriptor/3131686-hazardtrackingmode?language=objc
func (h_ HeapDescriptor) SetHazardTrackingMode(value HazardTrackingMode) {
	objc.Call[objc.Void](h_, objc.Sel("setHazardTrackingMode:"), value)
}
