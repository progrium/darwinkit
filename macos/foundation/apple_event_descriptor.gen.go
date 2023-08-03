// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var AppleEventDescriptorClass = _AppleEventDescriptorClass{objc.GetClass("NSAppleEventDescriptor")}

type _AppleEventDescriptorClass struct {
	objc.Class
}

type IAppleEventDescriptor interface {
	objc.IObject
	DescriptorAtIndex(index int) AppleEventDescriptor
	InsertDescriptorAtIndex(descriptor IAppleEventDescriptor, index int)
	RemoveDescriptorAtIndex(index int)
	SendEventWithOptionsTimeoutError(sendOptions AppleEventSendOptions, timeoutInSeconds TimeInterval, error *Error) AppleEventDescriptor
	BooleanValue() bool
	Data() []byte
	Int32Value() int32
	NumberOfItems() int
	StringValue() string
	DateValue() Date
	DoubleValue() float64
	FileURLValue() URL
	IsRecordDescriptor() bool
}

type AppleEventDescriptor struct {
	objc.Object
}

func MakeAppleEventDescriptor(ptr unsafe.Pointer) AppleEventDescriptor {
	return AppleEventDescriptor{
		Object: objc.MakeObject(ptr),
	}
}

func (a_ AppleEventDescriptor) InitListDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](a_, objc.GetSelector("initListDescriptor"))
	return rv
}

func AppleEventDescriptor_InitListDescriptor() AppleEventDescriptor {
	return AppleEventDescriptorClass.Alloc().InitListDescriptor()
}

func (a_ AppleEventDescriptor) InitRecordDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](a_, objc.GetSelector("initRecordDescriptor"))
	return rv
}

func AppleEventDescriptor_InitRecordDescriptor() AppleEventDescriptor {
	return AppleEventDescriptorClass.Alloc().InitRecordDescriptor()
}

func (ac _AppleEventDescriptorClass) Alloc() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("alloc"))
	return rv
}

func AppleEventDescriptor_Alloc() AppleEventDescriptor {
	return AppleEventDescriptorClass.Alloc()
}

func (ac _AppleEventDescriptorClass) New() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewAppleEventDescriptor() AppleEventDescriptor {
	return AppleEventDescriptorClass.New()
}

func AppleEventDescriptor_New() AppleEventDescriptor {
	return AppleEventDescriptorClass.New()
}

func (a_ AppleEventDescriptor) Init() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](a_, objc.GetSelector("init"))
	return rv
}

func AppleEventDescriptor_Init() AppleEventDescriptor {
	return AppleEventDescriptorClass.Alloc().Init()
}

func (ac _AppleEventDescriptorClass) DescriptorWithBoolean(boolean bool) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithBoolean:"), boolean)
	return rv
}

func AppleEventDescriptor_DescriptorWithBoolean(boolean bool) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithBoolean(boolean)
}

func (ac _AppleEventDescriptorClass) DescriptorWithInt32(signedInt int32) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithInt32:"), signedInt)
	return rv
}

func AppleEventDescriptor_DescriptorWithInt32(signedInt int32) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithInt32(signedInt)
}

func (ac _AppleEventDescriptorClass) DescriptorWithString(string_ string) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithString:"), string_)
	return rv
}

func AppleEventDescriptor_DescriptorWithString(string_ string) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithString(string_)
}

func (ac _AppleEventDescriptorClass) ListDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("listDescriptor"))
	return rv
}

func AppleEventDescriptor_ListDescriptor() AppleEventDescriptor {
	return AppleEventDescriptorClass.ListDescriptor()
}

func (ac _AppleEventDescriptorClass) NullDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("nullDescriptor"))
	return rv
}

func AppleEventDescriptor_NullDescriptor() AppleEventDescriptor {
	return AppleEventDescriptorClass.NullDescriptor()
}

func (ac _AppleEventDescriptorClass) RecordDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("recordDescriptor"))
	return rv
}

func AppleEventDescriptor_RecordDescriptor() AppleEventDescriptor {
	return AppleEventDescriptorClass.RecordDescriptor()
}

func (a_ AppleEventDescriptor) DescriptorAtIndex(index int) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](a_, objc.GetSelector("descriptorAtIndex:"), index)
	return rv
}

func (a_ AppleEventDescriptor) InsertDescriptorAtIndex(descriptor IAppleEventDescriptor, index int) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("insertDescriptor:atIndex:"), objc.ExtractPtr(descriptor), index)
}

func (a_ AppleEventDescriptor) RemoveDescriptorAtIndex(index int) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("removeDescriptorAtIndex:"), index)
}

func (ac _AppleEventDescriptorClass) DescriptorWithApplicationURL(applicationURL IURL) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithApplicationURL:"), objc.ExtractPtr(applicationURL))
	return rv
}

func AppleEventDescriptor_DescriptorWithApplicationURL(applicationURL IURL) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithApplicationURL(applicationURL)
}

func (ac _AppleEventDescriptorClass) DescriptorWithBundleIdentifier(bundleIdentifier string) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithBundleIdentifier:"), bundleIdentifier)
	return rv
}

func AppleEventDescriptor_DescriptorWithBundleIdentifier(bundleIdentifier string) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithBundleIdentifier(bundleIdentifier)
}

func (ac _AppleEventDescriptorClass) DescriptorWithDate(date IDate) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithDate:"), objc.ExtractPtr(date))
	return rv
}

func AppleEventDescriptor_DescriptorWithDate(date IDate) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithDate(date)
}

func (ac _AppleEventDescriptorClass) DescriptorWithDouble(doubleValue float64) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithDouble:"), doubleValue)
	return rv
}

func AppleEventDescriptor_DescriptorWithDouble(doubleValue float64) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithDouble(doubleValue)
}

func (ac _AppleEventDescriptorClass) DescriptorWithFileURL(fileURL IURL) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithFileURL:"), objc.ExtractPtr(fileURL))
	return rv
}

func AppleEventDescriptor_DescriptorWithFileURL(fileURL IURL) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithFileURL(fileURL)
}

func (a_ AppleEventDescriptor) SendEventWithOptionsTimeoutError(sendOptions AppleEventSendOptions, timeoutInSeconds TimeInterval, error *Error) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](a_, objc.GetSelector("sendEventWithOptions:timeout:error:"), sendOptions, timeoutInSeconds, unsafe.Pointer(error))
	return rv
}

func (ac _AppleEventDescriptorClass) CurrentProcessDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("currentProcessDescriptor"))
	return rv
}

func AppleEventDescriptor_CurrentProcessDescriptor() AppleEventDescriptor {
	return AppleEventDescriptorClass.CurrentProcessDescriptor()
}

func (a_ AppleEventDescriptor) BooleanValue() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("booleanValue"))
	return rv
}

func (a_ AppleEventDescriptor) Data() []byte {
	rv := objc.CallMethod[[]byte](a_, objc.GetSelector("data"))
	return rv
}

func (a_ AppleEventDescriptor) Int32Value() int32 {
	rv := objc.CallMethod[int32](a_, objc.GetSelector("int32Value"))
	return rv
}

func (a_ AppleEventDescriptor) NumberOfItems() int {
	rv := objc.CallMethod[int](a_, objc.GetSelector("numberOfItems"))
	return rv
}

func (a_ AppleEventDescriptor) StringValue() string {
	rv := objc.CallMethod[string](a_, objc.GetSelector("stringValue"))
	return rv
}

func (a_ AppleEventDescriptor) DateValue() Date {
	rv := objc.CallMethod[Date](a_, objc.GetSelector("dateValue"))
	return rv
}

func (a_ AppleEventDescriptor) DoubleValue() float64 {
	rv := objc.CallMethod[float64](a_, objc.GetSelector("doubleValue"))
	return rv
}

func (a_ AppleEventDescriptor) FileURLValue() URL {
	rv := objc.CallMethod[URL](a_, objc.GetSelector("fileURLValue"))
	return rv
}

func (a_ AppleEventDescriptor) IsRecordDescriptor() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("isRecordDescriptor"))
	return rv
}
