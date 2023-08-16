// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BlitPassDescriptor] class.
var BlitPassDescriptorClass = _BlitPassDescriptorClass{objc.GetClass("MTLBlitPassDescriptor")}

type _BlitPassDescriptorClass struct {
	objc.Class
}

// An interface definition for the [BlitPassDescriptor] class.
type IBlitPassDescriptor interface {
	objc.IObject
	SampleBufferAttachments() BlitPassSampleBufferAttachmentDescriptorArray
}

// A configuration you create to customize a blit command encoder, which affects the runtime behavior of the blit pass you encode with it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpassdescriptor?language=objc
type BlitPassDescriptor struct {
	objc.Object
}

func BlitPassDescriptorFrom(ptr unsafe.Pointer) BlitPassDescriptor {
	return BlitPassDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (bc _BlitPassDescriptorClass) Alloc() BlitPassDescriptor {
	rv := objc.Call[BlitPassDescriptor](bc, objc.Sel("alloc"))
	return rv
}

func BlitPassDescriptor_Alloc() BlitPassDescriptor {
	return BlitPassDescriptorClass.Alloc()
}

func (bc _BlitPassDescriptorClass) New() BlitPassDescriptor {
	rv := objc.Call[BlitPassDescriptor](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBlitPassDescriptor() BlitPassDescriptor {
	return BlitPassDescriptorClass.New()
}

func (b_ BlitPassDescriptor) Init() BlitPassDescriptor {
	rv := objc.Call[BlitPassDescriptor](b_, objc.Sel("init"))
	return rv
}

// Creates a new blit pass descriptor with a default configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpassdescriptor/3564421-blitpassdescriptor?language=objc
func (bc _BlitPassDescriptorClass) BlitPassDescriptor() BlitPassDescriptor {
	rv := objc.Call[BlitPassDescriptor](bc, objc.Sel("blitPassDescriptor"))
	return rv
}

// Creates a new blit pass descriptor with a default configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpassdescriptor/3564421-blitpassdescriptor?language=objc
func BlitPassDescriptor_BlitPassDescriptor() BlitPassDescriptor {
	return BlitPassDescriptorClass.BlitPassDescriptor()
}

// An array of counter sample buffer attachments that you configure for a blit pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpassdescriptor/3564422-samplebufferattachments?language=objc
func (b_ BlitPassDescriptor) SampleBufferAttachments() BlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[BlitPassSampleBufferAttachmentDescriptorArray](b_, objc.Sel("sampleBufferAttachments"))
	return rv
}
