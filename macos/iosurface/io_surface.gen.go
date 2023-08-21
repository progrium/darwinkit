// AUTO-GENERATED CODE, DO NOT MODIFY

package iosurface

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [IOSurface] class.
var IOSurfaceClass = _IOSurfaceClass{objc.GetClass("IOSurface")}

type _IOSurfaceClass struct {
	objc.Class
}

// An interface definition for the [IOSurface] class.
type IIOSurface interface {
	objc.IObject
	AllAttachments() map[string]objc.Object
	ElementHeightOfPlaneAtIndex(planeIndex uint) int
	LockWithOptionsSeed(options LockOptions, seed *uint32) int
	SetPurgeableOldState(newState PurgeabilityState, oldState *PurgeabilityState) int
	HeightOfPlaneAtIndex(planeIndex uint) int
	RemoveAttachmentForKey(key string)
	IncrementUseCount()
	SetAttachmentForKey(anObject objc.IObject, key string)
	SetAllAttachments(dict map[string]objc.IObject)
	RemoveAllAttachments()
	DecrementUseCount()
	UnlockWithOptionsSeed(options LockOptions, seed *uint32) int
	BaseAddressOfPlaneAtIndex(planeIndex uint) unsafe.Pointer
	BytesPerRowOfPlaneAtIndex(planeIndex uint) int
	AttachmentForKey(key string) objc.Object
	ElementWidthOfPlaneAtIndex(planeIndex uint) int
	WidthOfPlaneAtIndex(planeIndex uint) int
	BytesPerElementOfPlaneAtIndex(planeIndex uint) int
	IsInUse() bool
	Width() int
	LocalUseCount() int32
	PlaneCount() uint
	ElementHeight() int
	ElementWidth() int
	Height() int
	BytesPerRow() int
	BaseAddress() unsafe.Pointer
	PixelFormat() uint
	AllowsPixelSizeCasting() bool
	BytesPerElement() int
	AllocationSize() int
	Seed() uint32
}

// Data type representing an IOSurface opaque object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface?language=objc
type IOSurface struct {
	objc.Object
}

func IOSurfaceFrom(ptr unsafe.Pointer) IOSurface {
	return IOSurface{
		Object: objc.ObjectFrom(ptr),
	}
}

func (i_ IOSurface) InitWithProperties(properties map[PropertyKey]objc.IObject) IOSurface {
	rv := objc.Call[IOSurface](i_, objc.Sel("initWithProperties:"), properties)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092523-initwithproperties?language=objc
func NewIOSurfaceWithProperties(properties map[PropertyKey]objc.IObject) IOSurface {
	instance := IOSurfaceClass.Alloc().InitWithProperties(properties)
	instance.Autorelease()
	return instance
}

func (ic _IOSurfaceClass) Alloc() IOSurface {
	rv := objc.Call[IOSurface](ic, objc.Sel("alloc"))
	return rv
}

func IOSurface_Alloc() IOSurface {
	return IOSurfaceClass.Alloc()
}

func (ic _IOSurfaceClass) New() IOSurface {
	rv := objc.Call[IOSurface](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewIOSurface() IOSurface {
	return IOSurfaceClass.New()
}

func (i_ IOSurface) Init() IOSurface {
	rv := objc.Call[IOSurface](i_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092487-allattachments?language=objc
func (i_ IOSurface) AllAttachments() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](i_, objc.Sel("allAttachments"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092515-elementheightofplaneatindex?language=objc
func (i_ IOSurface) ElementHeightOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Call[int](i_, objc.Sel("elementHeightOfPlaneAtIndex:"), planeIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092522-lockwithoptions?language=objc
func (i_ IOSurface) LockWithOptionsSeed(options LockOptions, seed *uint32) int {
	rv := objc.Call[int](i_, objc.Sel("lockWithOptions:seed:"), options, seed)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2870085-setpurgeable?language=objc
func (i_ IOSurface) SetPurgeableOldState(newState PurgeabilityState, oldState *PurgeabilityState) int {
	rv := objc.Call[int](i_, objc.Sel("setPurgeable:oldState:"), newState, oldState)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092493-heightofplaneatindex?language=objc
func (i_ IOSurface) HeightOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Call[int](i_, objc.Sel("heightOfPlaneAtIndex:"), planeIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092541-removeattachmentforkey?language=objc
func (i_ IOSurface) RemoveAttachmentForKey(key string) {
	objc.Call[objc.Void](i_, objc.Sel("removeAttachmentForKey:"), key)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092506-incrementusecount?language=objc
func (i_ IOSurface) IncrementUseCount() {
	objc.Call[objc.Void](i_, objc.Sel("incrementUseCount"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092540-setattachment?language=objc
func (i_ IOSurface) SetAttachmentForKey(anObject objc.IObject, key string) {
	objc.Call[objc.Void](i_, objc.Sel("setAttachment:forKey:"), anObject, key)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092532-setallattachments?language=objc
func (i_ IOSurface) SetAllAttachments(dict map[string]objc.IObject) {
	objc.Call[objc.Void](i_, objc.Sel("setAllAttachments:"), dict)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092490-removeallattachments?language=objc
func (i_ IOSurface) RemoveAllAttachments() {
	objc.Call[objc.Void](i_, objc.Sel("removeAllAttachments"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092538-decrementusecount?language=objc
func (i_ IOSurface) DecrementUseCount() {
	objc.Call[objc.Void](i_, objc.Sel("decrementUseCount"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092530-unlockwithoptions?language=objc
func (i_ IOSurface) UnlockWithOptionsSeed(options LockOptions, seed *uint32) int {
	rv := objc.Call[int](i_, objc.Sel("unlockWithOptions:seed:"), options, seed)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092494-baseaddressofplaneatindex?language=objc
func (i_ IOSurface) BaseAddressOfPlaneAtIndex(planeIndex uint) unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](i_, objc.Sel("baseAddressOfPlaneAtIndex:"), planeIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092497-bytesperrowofplaneatindex?language=objc
func (i_ IOSurface) BytesPerRowOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Call[int](i_, objc.Sel("bytesPerRowOfPlaneAtIndex:"), planeIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092535-attachmentforkey?language=objc
func (i_ IOSurface) AttachmentForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("attachmentForKey:"), key)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092498-elementwidthofplaneatindex?language=objc
func (i_ IOSurface) ElementWidthOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Call[int](i_, objc.Sel("elementWidthOfPlaneAtIndex:"), planeIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092520-widthofplaneatindex?language=objc
func (i_ IOSurface) WidthOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Call[int](i_, objc.Sel("widthOfPlaneAtIndex:"), planeIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092511-bytesperelementofplaneatindex?language=objc
func (i_ IOSurface) BytesPerElementOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Call[int](i_, objc.Sel("bytesPerElementOfPlaneAtIndex:"), planeIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092504-inuse?language=objc
func (i_ IOSurface) IsInUse() bool {
	rv := objc.Call[bool](i_, objc.Sel("isInUse"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092516-width?language=objc
func (i_ IOSurface) Width() int {
	rv := objc.Call[int](i_, objc.Sel("width"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092529-localusecount?language=objc
func (i_ IOSurface) LocalUseCount() int32 {
	rv := objc.Call[int32](i_, objc.Sel("localUseCount"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092508-planecount?language=objc
func (i_ IOSurface) PlaneCount() uint {
	rv := objc.Call[uint](i_, objc.Sel("planeCount"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092539-elementheight?language=objc
func (i_ IOSurface) ElementHeight() int {
	rv := objc.Call[int](i_, objc.Sel("elementHeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092514-elementwidth?language=objc
func (i_ IOSurface) ElementWidth() int {
	rv := objc.Call[int](i_, objc.Sel("elementWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092501-height?language=objc
func (i_ IOSurface) Height() int {
	rv := objc.Call[int](i_, objc.Sel("height"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092502-bytesperrow?language=objc
func (i_ IOSurface) BytesPerRow() int {
	rv := objc.Call[int](i_, objc.Sel("bytesPerRow"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092525-baseaddress?language=objc
func (i_ IOSurface) BaseAddress() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](i_, objc.Sel("baseAddress"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092537-pixelformat?language=objc
func (i_ IOSurface) PixelFormat() uint {
	rv := objc.Call[uint](i_, objc.Sel("pixelFormat"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092536-allowspixelsizecasting?language=objc
func (i_ IOSurface) AllowsPixelSizeCasting() bool {
	rv := objc.Call[bool](i_, objc.Sel("allowsPixelSizeCasting"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092503-bytesperelement?language=objc
func (i_ IOSurface) BytesPerElement() int {
	rv := objc.Call[int](i_, objc.Sel("bytesPerElement"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092534-allocationsize?language=objc
func (i_ IOSurface) AllocationSize() int {
	rv := objc.Call[int](i_, objc.Sel("allocationSize"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/2092533-seed?language=objc
func (i_ IOSurface) Seed() uint32 {
	rv := objc.Call[uint32](i_, objc.Sel("seed"))
	return rv
}
