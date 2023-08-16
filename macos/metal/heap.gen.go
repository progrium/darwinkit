// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// A memory pool from which you can suballocate resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap?language=objc
type PHeap interface {
	// optional
	NewTextureWithDescriptor(descriptor TextureDescriptor) PTexture
	HasNewTextureWithDescriptor() bool

	// optional
	NewBufferWithLengthOptions(length uint, options ResourceOptions) PBuffer
	HasNewBufferWithLengthOptions() bool

	// optional
	MaxAvailableSizeWithAlignment(alignment uint) uint
	HasMaxAvailableSizeWithAlignment() bool

	// optional
	SetPurgeableState(state PurgeableState) PurgeableState
	HasSetPurgeableState() bool

	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	CpuCacheMode() CPUCacheMode
	HasCpuCacheMode() bool

	// optional
	StorageMode() StorageMode
	HasStorageMode() bool

	// optional
	UsedSize() uint
	HasUsedSize() bool

	// optional
	ResourceOptions() ResourceOptions
	HasResourceOptions() bool

	// optional
	Type() HeapType
	HasType() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	Size() uint
	HasSize() bool

	// optional
	CurrentAllocatedSize() uint
	HasCurrentAllocatedSize() bool

	// optional
	HazardTrackingMode() HazardTrackingMode
	HasHazardTrackingMode() bool
}

// A concrete type wrapper for the [PHeap] protocol.
type HeapWrapper struct {
	objc.Object
}

func (h_ HeapWrapper) HasNewTextureWithDescriptor() bool {
	return h_.RespondsToSelector(objc.Sel("newTextureWithDescriptor:"))
}

// Creates a texture on the heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap/1649574-newtexturewithdescriptor?language=objc
func (h_ HeapWrapper) NewTextureWithDescriptor(descriptor ITextureDescriptor) TextureWrapper {
	rv := objc.Call[TextureWrapper](h_, objc.Sel("newTextureWithDescriptor:"), objc.Ptr(descriptor))
	rv.Autorelease()
	return rv
}

func (h_ HeapWrapper) HasNewBufferWithLengthOptions() bool {
	return h_.RespondsToSelector(objc.Sel("newBufferWithLength:options:"))
}

// Creates a buffer on the heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap/1649571-newbufferwithlength?language=objc
func (h_ HeapWrapper) NewBufferWithLengthOptions(length uint, options ResourceOptions) BufferWrapper {
	rv := objc.Call[BufferWrapper](h_, objc.Sel("newBufferWithLength:options:"), length, options)
	rv.Autorelease()
	return rv
}

func (h_ HeapWrapper) HasMaxAvailableSizeWithAlignment() bool {
	return h_.RespondsToSelector(objc.Sel("maxAvailableSizeWithAlignment:"))
}

// The maximum size of a resource, in bytes, that can be currently allocated from the heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap/1771284-maxavailablesizewithalignment?language=objc
func (h_ HeapWrapper) MaxAvailableSizeWithAlignment(alignment uint) uint {
	rv := objc.Call[uint](h_, objc.Sel("maxAvailableSizeWithAlignment:"), alignment)
	return rv
}

func (h_ HeapWrapper) HasSetPurgeableState() bool {
	return h_.RespondsToSelector(objc.Sel("setPurgeableState:"))
}

// Sets the purgeable state of the heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap/1771281-setpurgeablestate?language=objc
func (h_ HeapWrapper) SetPurgeableState(state PurgeableState) PurgeableState {
	rv := objc.Call[PurgeableState](h_, objc.Sel("setPurgeableState:"), state)
	return rv
}

func (h_ HeapWrapper) HasDevice() bool {
	return h_.RespondsToSelector(objc.Sel("device"))
}

// The device object that created the heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap/1771285-device?language=objc
func (h_ HeapWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](h_, objc.Sel("device"))
	return rv
}

func (h_ HeapWrapper) HasCpuCacheMode() bool {
	return h_.RespondsToSelector(objc.Sel("cpuCacheMode"))
}

// The heap's CPU cache mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap/1771280-cpucachemode?language=objc
func (h_ HeapWrapper) CpuCacheMode() CPUCacheMode {
	rv := objc.Call[CPUCacheMode](h_, objc.Sel("cpuCacheMode"))
	return rv
}

func (h_ HeapWrapper) HasStorageMode() bool {
	return h_.RespondsToSelector(objc.Sel("storageMode"))
}

// The heap's storage mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap/1771282-storagemode?language=objc
func (h_ HeapWrapper) StorageMode() StorageMode {
	rv := objc.Call[StorageMode](h_, objc.Sel("storageMode"))
	return rv
}

func (h_ HeapWrapper) HasUsedSize() bool {
	return h_.RespondsToSelector(objc.Sel("usedSize"))
}

// The size of all resources currently in the heap, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap/2097557-usedsize?language=objc
func (h_ HeapWrapper) UsedSize() uint {
	rv := objc.Call[uint](h_, objc.Sel("usedSize"))
	return rv
}

func (h_ HeapWrapper) HasResourceOptions() bool {
	return h_.RespondsToSelector(objc.Sel("resourceOptions"))
}

// The options for resources created by the heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap/3131685-resourceoptions?language=objc
func (h_ HeapWrapper) ResourceOptions() ResourceOptions {
	rv := objc.Call[ResourceOptions](h_, objc.Sel("resourceOptions"))
	return rv
}

func (h_ HeapWrapper) HasType() bool {
	return h_.RespondsToSelector(objc.Sel("type"))
}

// The heap's type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap/3043386-type?language=objc
func (h_ HeapWrapper) Type() HeapType {
	rv := objc.Call[HeapType](h_, objc.Sel("type"))
	return rv
}

func (h_ HeapWrapper) HasSetLabel() bool {
	return h_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that identifies the heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap/1771279-label?language=objc
func (h_ HeapWrapper) SetLabel(value string) {
	objc.Call[objc.Void](h_, objc.Sel("setLabel:"), value)
}

func (h_ HeapWrapper) HasLabel() bool {
	return h_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap/1771279-label?language=objc
func (h_ HeapWrapper) Label() string {
	rv := objc.Call[string](h_, objc.Sel("label"))
	return rv
}

func (h_ HeapWrapper) HasSize() bool {
	return h_.RespondsToSelector(objc.Sel("size"))
}

// The total size of the heap, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap/1649569-size?language=objc
func (h_ HeapWrapper) Size() uint {
	rv := objc.Call[uint](h_, objc.Sel("size"))
	return rv
}

func (h_ HeapWrapper) HasCurrentAllocatedSize() bool {
	return h_.RespondsToSelector(objc.Sel("currentAllocatedSize"))
}

// The size, in bytes, of the current heap allocation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap/2915348-currentallocatedsize?language=objc
func (h_ HeapWrapper) CurrentAllocatedSize() uint {
	rv := objc.Call[uint](h_, objc.Sel("currentAllocatedSize"))
	return rv
}

func (h_ HeapWrapper) HasHazardTrackingMode() bool {
	return h_.RespondsToSelector(objc.Sel("hazardTrackingMode"))
}

// The heap's hazard tracking mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheap/3131684-hazardtrackingmode?language=objc
func (h_ HeapWrapper) HazardTrackingMode() HazardTrackingMode {
	rv := objc.Call[HazardTrackingMode](h_, objc.Sel("hazardTrackingMode"))
	return rv
}
