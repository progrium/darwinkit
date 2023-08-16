// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A container for pipeline state descriptors and their associated compiled shader code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive?language=objc
type PBinaryArchive interface {
	// optional
	AddFunctionWithDescriptorLibraryError(descriptor FunctionDescriptor, library LibraryWrapper, error foundation.Error) bool
	HasAddFunctionWithDescriptorLibraryError() bool

	// optional
	AddComputePipelineFunctionsWithDescriptorError(descriptor ComputePipelineDescriptor, error foundation.Error) bool
	HasAddComputePipelineFunctionsWithDescriptorError() bool

	// optional
	AddRenderPipelineFunctionsWithDescriptorError(descriptor RenderPipelineDescriptor, error foundation.Error) bool
	HasAddRenderPipelineFunctionsWithDescriptorError() bool

	// optional
	AddTileRenderPipelineFunctionsWithDescriptorError(descriptor TileRenderPipelineDescriptor, error foundation.Error) bool
	HasAddTileRenderPipelineFunctionsWithDescriptorError() bool

	// optional
	SerializeToURLError(url foundation.URL, error foundation.Error) bool
	HasSerializeToURLError() bool

	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool
}

// A concrete type wrapper for the [PBinaryArchive] protocol.
type BinaryArchiveWrapper struct {
	objc.Object
}

func (b_ BinaryArchiveWrapper) HasAddFunctionWithDescriptorLibraryError() bool {
	return b_.RespondsToSelector(objc.Sel("addFunctionWithDescriptor:library:error:"))
}

// Adds a description of a function to the archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3750523-addfunctionwithdescriptor?language=objc
func (b_ BinaryArchiveWrapper) AddFunctionWithDescriptorLibraryError(descriptor IFunctionDescriptor, library PLibrary, error foundation.IError) bool {
	po1 := objc.WrapAsProtocol("MTLLibrary", library)
	rv := objc.Call[bool](b_, objc.Sel("addFunctionWithDescriptor:library:error:"), objc.Ptr(descriptor), po1, objc.Ptr(error))
	return rv
}

func (b_ BinaryArchiveWrapper) HasAddComputePipelineFunctionsWithDescriptorError() bool {
	return b_.RespondsToSelector(objc.Sel("addComputePipelineFunctionsWithDescriptor:error:"))
}

// Adds a description of a compute pipeline to the archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3553924-addcomputepipelinefunctionswithd?language=objc
func (b_ BinaryArchiveWrapper) AddComputePipelineFunctionsWithDescriptorError(descriptor IComputePipelineDescriptor, error foundation.IError) bool {
	rv := objc.Call[bool](b_, objc.Sel("addComputePipelineFunctionsWithDescriptor:error:"), objc.Ptr(descriptor), objc.Ptr(error))
	return rv
}

func (b_ BinaryArchiveWrapper) HasAddRenderPipelineFunctionsWithDescriptorError() bool {
	return b_.RespondsToSelector(objc.Sel("addRenderPipelineFunctionsWithDescriptor:error:"))
}

// Adds a description of a render pipeline to the archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3553925-addrenderpipelinefunctionswithde?language=objc
func (b_ BinaryArchiveWrapper) AddRenderPipelineFunctionsWithDescriptorError(descriptor IRenderPipelineDescriptor, error foundation.IError) bool {
	rv := objc.Call[bool](b_, objc.Sel("addRenderPipelineFunctionsWithDescriptor:error:"), objc.Ptr(descriptor), objc.Ptr(error))
	return rv
}

func (b_ BinaryArchiveWrapper) HasAddTileRenderPipelineFunctionsWithDescriptorError() bool {
	return b_.RespondsToSelector(objc.Sel("addTileRenderPipelineFunctionsWithDescriptor:error:"))
}

// Adds a description of a tile renderer pipeline to the archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3564419-addtilerenderpipelinefunctionswi?language=objc
func (b_ BinaryArchiveWrapper) AddTileRenderPipelineFunctionsWithDescriptorError(descriptor ITileRenderPipelineDescriptor, error foundation.IError) bool {
	rv := objc.Call[bool](b_, objc.Sel("addTileRenderPipelineFunctionsWithDescriptor:error:"), objc.Ptr(descriptor), objc.Ptr(error))
	return rv
}

func (b_ BinaryArchiveWrapper) HasSerializeToURLError() bool {
	return b_.RespondsToSelector(objc.Sel("serializeToURL:error:"))
}

// Writes the contents of the archive to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3553928-serializetourl?language=objc
func (b_ BinaryArchiveWrapper) SerializeToURLError(url foundation.IURL, error foundation.IError) bool {
	rv := objc.Call[bool](b_, objc.Sel("serializeToURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

func (b_ BinaryArchiveWrapper) HasDevice() bool {
	return b_.RespondsToSelector(objc.Sel("device"))
}

// The Metal device object that created the binary archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3553926-device?language=objc
func (b_ BinaryArchiveWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](b_, objc.Sel("device"))
	return rv
}

func (b_ BinaryArchiveWrapper) HasSetLabel() bool {
	return b_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that identifies the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3553927-label?language=objc
func (b_ BinaryArchiveWrapper) SetLabel(value string) {
	objc.Call[objc.Void](b_, objc.Sel("setLabel:"), value)
}

func (b_ BinaryArchiveWrapper) HasLabel() bool {
	return b_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3553927-label?language=objc
func (b_ BinaryArchiveWrapper) Label() string {
	rv := objc.Call[string](b_, objc.Sel("label"))
	return rv
}
