// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ResourceStatePassDescriptor] class.
var ResourceStatePassDescriptorClass = _ResourceStatePassDescriptorClass{objc.GetClass("MTLResourceStatePassDescriptor")}

type _ResourceStatePassDescriptorClass struct {
	objc.Class
}

// An interface definition for the [ResourceStatePassDescriptor] class.
type IResourceStatePassDescriptor interface {
	objc.IObject
	SampleBufferAttachments() ResourceStatePassSampleBufferAttachmentDescriptorArray
}

// A configuration for a resource state pass, used to create a resource state command encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepassdescriptor?language=objc
type ResourceStatePassDescriptor struct {
	objc.Object
}

func ResourceStatePassDescriptorFrom(ptr unsafe.Pointer) ResourceStatePassDescriptor {
	return ResourceStatePassDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _ResourceStatePassDescriptorClass) Alloc() ResourceStatePassDescriptor {
	rv := objc.Call[ResourceStatePassDescriptor](rc, objc.Sel("alloc"))
	return rv
}

func ResourceStatePassDescriptor_Alloc() ResourceStatePassDescriptor {
	return ResourceStatePassDescriptorClass.Alloc()
}

func (rc _ResourceStatePassDescriptorClass) New() ResourceStatePassDescriptor {
	rv := objc.Call[ResourceStatePassDescriptor](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewResourceStatePassDescriptor() ResourceStatePassDescriptor {
	return ResourceStatePassDescriptorClass.New()
}

func (r_ ResourceStatePassDescriptor) Init() ResourceStatePassDescriptor {
	rv := objc.Call[ResourceStatePassDescriptor](r_, objc.Sel("init"))
	return rv
}

// Creates a new resource state pass descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepassdescriptor/3566563-resourcestatepassdescriptor?language=objc
func (rc _ResourceStatePassDescriptorClass) ResourceStatePassDescriptor() ResourceStatePassDescriptor {
	rv := objc.Call[ResourceStatePassDescriptor](rc, objc.Sel("resourceStatePassDescriptor"))
	return rv
}

// Creates a new resource state pass descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepassdescriptor/3566563-resourcestatepassdescriptor?language=objc
func ResourceStatePassDescriptor_ResourceStatePassDescriptor() ResourceStatePassDescriptor {
	return ResourceStatePassDescriptorClass.ResourceStatePassDescriptor()
}

// The array of sample buffers that the resource state pass can access. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepassdescriptor/3566564-samplebufferattachments?language=objc
func (r_ ResourceStatePassDescriptor) SampleBufferAttachments() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[ResourceStatePassSampleBufferAttachmentDescriptorArray](r_, objc.Sel("sampleBufferAttachments"))
	return rv
}
