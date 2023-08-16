// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CounterSampleBufferDescriptor] class.
var CounterSampleBufferDescriptorClass = _CounterSampleBufferDescriptorClass{objc.GetClass("MTLCounterSampleBufferDescriptor")}

type _CounterSampleBufferDescriptorClass struct {
	objc.Class
}

// An interface definition for the [CounterSampleBufferDescriptor] class.
type ICounterSampleBufferDescriptor interface {
	objc.IObject
	SampleCount() uint
	SetSampleCount(value uint)
	CounterSet() CounterSetWrapper
	SetCounterSet(value PCounterSet)
	SetCounterSetObject(valueObject objc.IObject)
	StorageMode() StorageMode
	SetStorageMode(value StorageMode)
	Label() string
	SetLabel(value string)
}

// A group of properties that configures the counter sample buffers you create with it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebufferdescriptor?language=objc
type CounterSampleBufferDescriptor struct {
	objc.Object
}

func CounterSampleBufferDescriptorFrom(ptr unsafe.Pointer) CounterSampleBufferDescriptor {
	return CounterSampleBufferDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CounterSampleBufferDescriptorClass) Alloc() CounterSampleBufferDescriptor {
	rv := objc.Call[CounterSampleBufferDescriptor](cc, objc.Sel("alloc"))
	return rv
}

func CounterSampleBufferDescriptor_Alloc() CounterSampleBufferDescriptor {
	return CounterSampleBufferDescriptorClass.Alloc()
}

func (cc _CounterSampleBufferDescriptorClass) New() CounterSampleBufferDescriptor {
	rv := objc.Call[CounterSampleBufferDescriptor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCounterSampleBufferDescriptor() CounterSampleBufferDescriptor {
	return CounterSampleBufferDescriptorClass.New()
}

func (c_ CounterSampleBufferDescriptor) Init() CounterSampleBufferDescriptor {
	rv := objc.Call[CounterSampleBufferDescriptor](c_, objc.Sel("init"))
	return rv
}

// The number of instances of a counter set’s data that a counter sample buffer can store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebufferdescriptor/3081732-samplecount?language=objc
func (c_ CounterSampleBufferDescriptor) SampleCount() uint {
	rv := objc.Call[uint](c_, objc.Sel("sampleCount"))
	return rv
}

// The number of instances of a counter set’s data that a counter sample buffer can store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebufferdescriptor/3081732-samplecount?language=objc
func (c_ CounterSampleBufferDescriptor) SetSampleCount(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setSampleCount:"), value)
}

// A GPU device’s counter set instance that you want to sample. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebufferdescriptor/3081730-counterset?language=objc
func (c_ CounterSampleBufferDescriptor) CounterSet() CounterSetWrapper {
	rv := objc.Call[CounterSetWrapper](c_, objc.Sel("counterSet"))
	return rv
}

// A GPU device’s counter set instance that you want to sample. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebufferdescriptor/3081730-counterset?language=objc
func (c_ CounterSampleBufferDescriptor) SetCounterSet(value PCounterSet) {
	po0 := objc.WrapAsProtocol("MTLCounterSet", value)
	objc.Call[objc.Void](c_, objc.Sel("setCounterSet:"), po0)
}

// A GPU device’s counter set instance that you want to sample. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebufferdescriptor/3081730-counterset?language=objc
func (c_ CounterSampleBufferDescriptor) SetCounterSetObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setCounterSet:"), objc.Ptr(valueObject))
}

// The memory storage mode for the counter sample buffers you create with the descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebufferdescriptor/3081733-storagemode?language=objc
func (c_ CounterSampleBufferDescriptor) StorageMode() StorageMode {
	rv := objc.Call[StorageMode](c_, objc.Sel("storageMode"))
	return rv
}

// The memory storage mode for the counter sample buffers you create with the descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebufferdescriptor/3081733-storagemode?language=objc
func (c_ CounterSampleBufferDescriptor) SetStorageMode(value StorageMode) {
	objc.Call[objc.Void](c_, objc.Sel("setStorageMode:"), value)
}

// The name for the counter sample buffer you create with the descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebufferdescriptor/3081731-label?language=objc
func (c_ CounterSampleBufferDescriptor) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}

// The name for the counter sample buffer you create with the descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebufferdescriptor/3081731-label?language=objc
func (c_ CounterSampleBufferDescriptor) SetLabel(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setLabel:"), value)
}
