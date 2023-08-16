// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BinaryArchiveDescriptor] class.
var BinaryArchiveDescriptorClass = _BinaryArchiveDescriptorClass{objc.GetClass("MTLBinaryArchiveDescriptor")}

type _BinaryArchiveDescriptorClass struct {
	objc.Class
}

// An interface definition for the [BinaryArchiveDescriptor] class.
type IBinaryArchiveDescriptor interface {
	objc.IObject
	Url() foundation.URL
	SetUrl(value foundation.IURL)
}

// A description of a binary shader archive that you want to create. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchivedescriptor?language=objc
type BinaryArchiveDescriptor struct {
	objc.Object
}

func BinaryArchiveDescriptorFrom(ptr unsafe.Pointer) BinaryArchiveDescriptor {
	return BinaryArchiveDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (bc _BinaryArchiveDescriptorClass) Alloc() BinaryArchiveDescriptor {
	rv := objc.Call[BinaryArchiveDescriptor](bc, objc.Sel("alloc"))
	return rv
}

func BinaryArchiveDescriptor_Alloc() BinaryArchiveDescriptor {
	return BinaryArchiveDescriptorClass.Alloc()
}

func (bc _BinaryArchiveDescriptorClass) New() BinaryArchiveDescriptor {
	rv := objc.Call[BinaryArchiveDescriptor](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBinaryArchiveDescriptor() BinaryArchiveDescriptor {
	return BinaryArchiveDescriptorClass.New()
}

func (b_ BinaryArchiveDescriptor) Init() BinaryArchiveDescriptor {
	rv := objc.Call[BinaryArchiveDescriptor](b_, objc.Sel("init"))
	return rv
}

// A URL to a Metal binary archive file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchivedescriptor/3553930-url?language=objc
func (b_ BinaryArchiveDescriptor) Url() foundation.URL {
	rv := objc.Call[foundation.URL](b_, objc.Sel("url"))
	return rv
}

// A URL to a Metal binary archive file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchivedescriptor/3553930-url?language=objc
func (b_ BinaryArchiveDescriptor) SetUrl(value foundation.IURL) {
	objc.Call[objc.Void](b_, objc.Sel("setUrl:"), objc.Ptr(value))
}
