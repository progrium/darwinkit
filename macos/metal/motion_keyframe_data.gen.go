// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MotionKeyframeData] class.
var MotionKeyframeDataClass = _MotionKeyframeDataClass{objc.GetClass("MTLMotionKeyframeData")}

type _MotionKeyframeDataClass struct {
	objc.Class
}

// An interface definition for the [MotionKeyframeData] class.
type IMotionKeyframeData interface {
	objc.IObject
	Buffer() BufferWrapper
	SetBuffer(value PBuffer)
	SetBufferObject(valueObject objc.IObject)
	Offset() uint
	SetOffset(value uint)
}

// Geometry data for a specific keyframe to use in a moving object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmotionkeyframedata?language=objc
type MotionKeyframeData struct {
	objc.Object
}

func MotionKeyframeDataFrom(ptr unsafe.Pointer) MotionKeyframeData {
	return MotionKeyframeData{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MotionKeyframeDataClass) Data() MotionKeyframeData {
	rv := objc.Call[MotionKeyframeData](mc, objc.Sel("data"))
	return rv
}

// Creates a new keyframe object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmotionkeyframedata/3750507-data?language=objc
func MotionKeyframeData_Data() MotionKeyframeData {
	return MotionKeyframeDataClass.Data()
}

func (mc _MotionKeyframeDataClass) Alloc() MotionKeyframeData {
	rv := objc.Call[MotionKeyframeData](mc, objc.Sel("alloc"))
	return rv
}

func MotionKeyframeData_Alloc() MotionKeyframeData {
	return MotionKeyframeDataClass.Alloc()
}

func (mc _MotionKeyframeDataClass) New() MotionKeyframeData {
	rv := objc.Call[MotionKeyframeData](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMotionKeyframeData() MotionKeyframeData {
	return MotionKeyframeDataClass.New()
}

func (m_ MotionKeyframeData) Init() MotionKeyframeData {
	rv := objc.Call[MotionKeyframeData](m_, objc.Sel("init"))
	return rv
}

// The buffer that holds the geometry data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmotionkeyframedata/3750506-buffer?language=objc
func (m_ MotionKeyframeData) Buffer() BufferWrapper {
	rv := objc.Call[BufferWrapper](m_, objc.Sel("buffer"))
	return rv
}

// The buffer that holds the geometry data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmotionkeyframedata/3750506-buffer?language=objc
func (m_ MotionKeyframeData) SetBuffer(value PBuffer) {
	po0 := objc.WrapAsProtocol("MTLBuffer", value)
	objc.Call[objc.Void](m_, objc.Sel("setBuffer:"), po0)
}

// The buffer that holds the geometry data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmotionkeyframedata/3750506-buffer?language=objc
func (m_ MotionKeyframeData) SetBufferObject(valueObject objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setBuffer:"), objc.Ptr(valueObject))
}

// The offset, in bytes, to the keyframe data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmotionkeyframedata/3750508-offset?language=objc
func (m_ MotionKeyframeData) Offset() uint {
	rv := objc.Call[uint](m_, objc.Sel("offset"))
	return rv
}

// The offset, in bytes, to the keyframe data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmotionkeyframedata/3750508-offset?language=objc
func (m_ MotionKeyframeData) SetOffset(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setOffset:"), value)
}
