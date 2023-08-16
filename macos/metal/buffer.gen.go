// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A resource that stores data in a format defined by your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer?language=objc
type PBuffer interface {
	// optional
	DidModifyRange(range_ foundation.Range)
	HasDidModifyRange() bool

	// optional
	NewTextureWithDescriptorOffsetBytesPerRow(descriptor TextureDescriptor, offset uint, bytesPerRow uint) PTexture
	HasNewTextureWithDescriptorOffsetBytesPerRow() bool

	// optional
	AddDebugMarkerRange(marker string, range_ foundation.Range)
	HasAddDebugMarkerRange() bool

	// optional
	Contents() unsafe.Pointer
	HasContents() bool

	// optional
	NewRemoteBufferViewForDevice(device DeviceWrapper) PBuffer
	HasNewRemoteBufferViewForDevice() bool

	// optional
	RemoveAllDebugMarkers()
	HasRemoveAllDebugMarkers() bool

	// optional
	RemoteStorageBuffer() PBuffer
	HasRemoteStorageBuffer() bool

	// optional
	Length() uint
	HasLength() bool
}

// A concrete type wrapper for the [PBuffer] protocol.
type BufferWrapper struct {
	objc.Object
}

func (b_ BufferWrapper) HasDidModifyRange() bool {
	return b_.RespondsToSelector(objc.Sel("didModifyRange:"))
}

// Informs the GPU that the CPU has modified a section of the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/1516121-didmodifyrange?language=objc
func (b_ BufferWrapper) DidModifyRange(range_ foundation.Range) {
	objc.Call[objc.Void](b_, objc.Sel("didModifyRange:"), range_)
}

func (b_ BufferWrapper) HasNewTextureWithDescriptorOffsetBytesPerRow() bool {
	return b_.RespondsToSelector(objc.Sel("newTextureWithDescriptor:offset:bytesPerRow:"))
}

// Creates a texture that shares its storage with the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/1613852-newtexturewithdescriptor?language=objc
func (b_ BufferWrapper) NewTextureWithDescriptorOffsetBytesPerRow(descriptor ITextureDescriptor, offset uint, bytesPerRow uint) TextureWrapper {
	rv := objc.Call[TextureWrapper](b_, objc.Sel("newTextureWithDescriptor:offset:bytesPerRow:"), objc.Ptr(descriptor), offset, bytesPerRow)
	rv.Autorelease()
	return rv
}

func (b_ BufferWrapper) HasAddDebugMarkerRange() bool {
	return b_.RespondsToSelector(objc.Sel("addDebugMarker:range:"))
}

// Adds a debug marker string to a specific buffer range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/1779576-adddebugmarker?language=objc
func (b_ BufferWrapper) AddDebugMarkerRange(marker string, range_ foundation.Range) {
	objc.Call[objc.Void](b_, objc.Sel("addDebugMarker:range:"), marker, range_)
}

func (b_ BufferWrapper) HasContents() bool {
	return b_.RespondsToSelector(objc.Sel("contents"))
}

// Gets the system address of the buffer’s storage allocation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/1515716-contents?language=objc
func (b_ BufferWrapper) Contents() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](b_, objc.Sel("contents"))
	return rv
}

func (b_ BufferWrapper) HasNewRemoteBufferViewForDevice() bool {
	return b_.RespondsToSelector(objc.Sel("newRemoteBufferViewForDevice:"))
}

// Creates a remote view of the buffer for another GPU in the same peer group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/2967415-newremotebufferviewfordevice?language=objc
func (b_ BufferWrapper) NewRemoteBufferViewForDevice(device PDevice) BufferWrapper {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[BufferWrapper](b_, objc.Sel("newRemoteBufferViewForDevice:"), po0)
	rv.Autorelease()
	return rv
}

func (b_ BufferWrapper) HasRemoveAllDebugMarkers() bool {
	return b_.RespondsToSelector(objc.Sel("removeAllDebugMarkers"))
}

// Removes all debug marker strings from the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/1779577-removealldebugmarkers?language=objc
func (b_ BufferWrapper) RemoveAllDebugMarkers() {
	objc.Call[objc.Void](b_, objc.Sel("removeAllDebugMarkers"))
}

func (b_ BufferWrapper) HasRemoteStorageBuffer() bool {
	return b_.RespondsToSelector(objc.Sel("remoteStorageBuffer"))
}

// The buffer on another GPU that the buffer was created from, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/2967416-remotestoragebuffer?language=objc
func (b_ BufferWrapper) RemoteStorageBuffer() BufferWrapper {
	rv := objc.Call[BufferWrapper](b_, objc.Sel("remoteStorageBuffer"))
	return rv
}

func (b_ BufferWrapper) HasLength() bool {
	return b_.RespondsToSelector(objc.Sel("length"))
}

// The logical size of the buffer, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/1515373-length?language=objc
func (b_ BufferWrapper) Length() uint {
	rv := objc.Call[uint](b_, objc.Sel("length"))
	return rv
}
