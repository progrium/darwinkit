// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// An allocation of memory that is accessible to a GPU. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresource?language=objc
type PResource interface {
	// optional
	IsAliasable() bool
	HasIsAliasable() bool

	// optional
	MakeAliasable()
	HasMakeAliasable() bool

	// optional
	SetPurgeableState(state PurgeableState) PurgeableState
	HasSetPurgeableState() bool

	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	AllocatedSize() uint
	HasAllocatedSize() bool

	// optional
	Heap() PHeap
	HasHeap() bool

	// optional
	CpuCacheMode() CPUCacheMode
	HasCpuCacheMode() bool

	// optional
	StorageMode() StorageMode
	HasStorageMode() bool

	// optional
	ResourceOptions() ResourceOptions
	HasResourceOptions() bool

	// optional
	HeapOffset() uint
	HasHeapOffset() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	HazardTrackingMode() HazardTrackingMode
	HasHazardTrackingMode() bool
}

// A concrete type wrapper for the [PResource] protocol.
type ResourceWrapper struct {
	objc.Object
}

func (r_ ResourceWrapper) HasIsAliasable() bool {
	return r_.RespondsToSelector(objc.Sel("isAliasable"))
}

// A Boolean value that indicates whether future heap resource allocations may alias against the resource’s memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresource/1771702-isaliasable?language=objc
func (r_ ResourceWrapper) IsAliasable() bool {
	rv := objc.Call[bool](r_, objc.Sel("isAliasable"))
	return rv
}

func (r_ ResourceWrapper) HasMakeAliasable() bool {
	return r_.RespondsToSelector(objc.Sel("makeAliasable"))
}

// Allows future heap resource allocations to alias against the resource’s memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresource/1771705-makealiasable?language=objc
func (r_ ResourceWrapper) MakeAliasable() {
	objc.Call[objc.Void](r_, objc.Sel("makeAliasable"))
}

func (r_ ResourceWrapper) HasSetPurgeableState() bool {
	return r_.RespondsToSelector(objc.Sel("setPurgeableState:"))
}

// Specifies or queries the resource’s purgeable state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresource/1515898-setpurgeablestate?language=objc
func (r_ ResourceWrapper) SetPurgeableState(state PurgeableState) PurgeableState {
	rv := objc.Call[PurgeableState](r_, objc.Sel("setPurgeableState:"), state)
	return rv
}

func (r_ ResourceWrapper) HasDevice() bool {
	return r_.RespondsToSelector(objc.Sel("device"))
}

// The device object that created the resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresource/1515442-device?language=objc
func (r_ ResourceWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](r_, objc.Sel("device"))
	return rv
}

func (r_ ResourceWrapper) HasAllocatedSize() bool {
	return r_.RespondsToSelector(objc.Sel("allocatedSize"))
}

// The size of the resource, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresource/2915287-allocatedsize?language=objc
func (r_ ResourceWrapper) AllocatedSize() uint {
	rv := objc.Call[uint](r_, objc.Sel("allocatedSize"))
	return rv
}

func (r_ ResourceWrapper) HasHeap() bool {
	return r_.RespondsToSelector(objc.Sel("heap"))
}

// The heap on which the resource is allocated, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresource/1682333-heap?language=objc
func (r_ ResourceWrapper) Heap() HeapWrapper {
	rv := objc.Call[HeapWrapper](r_, objc.Sel("heap"))
	return rv
}

func (r_ ResourceWrapper) HasCpuCacheMode() bool {
	return r_.RespondsToSelector(objc.Sel("cpuCacheMode"))
}

// The CPU cache mode that defines the CPU mapping of the resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresource/1516127-cpucachemode?language=objc
func (r_ ResourceWrapper) CpuCacheMode() CPUCacheMode {
	rv := objc.Call[CPUCacheMode](r_, objc.Sel("cpuCacheMode"))
	return rv
}

func (r_ ResourceWrapper) HasStorageMode() bool {
	return r_.RespondsToSelector(objc.Sel("storageMode"))
}

// The location and access permissions of the resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresource/1515477-storagemode?language=objc
func (r_ ResourceWrapper) StorageMode() StorageMode {
	rv := objc.Call[StorageMode](r_, objc.Sel("storageMode"))
	return rv
}

func (r_ ResourceWrapper) HasResourceOptions() bool {
	return r_.RespondsToSelector(objc.Sel("resourceOptions"))
}

// The storage mode, CPU cache mode, and hazard tracking mode of the resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresource/3131694-resourceoptions?language=objc
func (r_ ResourceWrapper) ResourceOptions() ResourceOptions {
	rv := objc.Call[ResourceOptions](r_, objc.Sel("resourceOptions"))
	return rv
}

func (r_ ResourceWrapper) HasHeapOffset() bool {
	return r_.RespondsToSelector(objc.Sel("heapOffset"))
}

// The distance, in bytes, from the beginning of the heap to the first byte of the resource, if you allocated the resource on a heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresource/3043406-heapoffset?language=objc
func (r_ ResourceWrapper) HeapOffset() uint {
	rv := objc.Call[uint](r_, objc.Sel("heapOffset"))
	return rv
}

func (r_ ResourceWrapper) HasSetLabel() bool {
	return r_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that identifies the resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresource/1515814-label?language=objc
func (r_ ResourceWrapper) SetLabel(value string) {
	objc.Call[objc.Void](r_, objc.Sel("setLabel:"), value)
}

func (r_ ResourceWrapper) HasLabel() bool {
	return r_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresource/1515814-label?language=objc
func (r_ ResourceWrapper) Label() string {
	rv := objc.Call[string](r_, objc.Sel("label"))
	return rv
}

func (r_ ResourceWrapper) HasHazardTrackingMode() bool {
	return r_.RespondsToSelector(objc.Sel("hazardTrackingMode"))
}

// A mode that determines whether Metal tracks and synchronizes resource access. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresource/3131693-hazardtrackingmode?language=objc
func (r_ ResourceWrapper) HazardTrackingMode() HazardTrackingMode {
	rv := objc.Call[HazardTrackingMode](r_, objc.Sel("hazardTrackingMode"))
	return rv
}
