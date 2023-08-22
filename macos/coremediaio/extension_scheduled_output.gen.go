// AUTO-GENERATED CODE, DO NOT MODIFY

package coremediaio

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExtensionScheduledOutput] class.
var ExtensionScheduledOutputClass = _ExtensionScheduledOutputClass{objc.GetClass("CMIOExtensionScheduledOutput")}

type _ExtensionScheduledOutputClass struct {
	objc.Class
}

// An interface definition for the [ExtensionScheduledOutput] class.
type IExtensionScheduledOutput interface {
	objc.IObject
	HostTimeInNanoseconds() uint64
	SequenceNumber() uint64
}

// An object that represents scheduled output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionscheduledoutput?language=objc
type ExtensionScheduledOutput struct {
	objc.Object
}

func ExtensionScheduledOutputFrom(ptr unsafe.Pointer) ExtensionScheduledOutput {
	return ExtensionScheduledOutput{
		Object: objc.ObjectFrom(ptr),
	}
}

func (e_ ExtensionScheduledOutput) InitWithSequenceNumberHostTimeInNanoseconds(sequenceNumber uint64, hostTimeInNanoseconds uint64) ExtensionScheduledOutput {
	rv := objc.Call[ExtensionScheduledOutput](e_, objc.Sel("initWithSequenceNumber:hostTimeInNanoseconds:"), sequenceNumber, hostTimeInNanoseconds)
	return rv
}

// Creates a scheduled output object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionscheduledoutput/3915885-initwithsequencenumber?language=objc
func NewExtensionScheduledOutputWithSequenceNumberHostTimeInNanoseconds(sequenceNumber uint64, hostTimeInNanoseconds uint64) ExtensionScheduledOutput {
	instance := ExtensionScheduledOutputClass.Alloc().InitWithSequenceNumberHostTimeInNanoseconds(sequenceNumber, hostTimeInNanoseconds)
	instance.Autorelease()
	return instance
}

func (ec _ExtensionScheduledOutputClass) ScheduledOutputWithSequenceNumberHostTimeInNanoseconds(sequenceNumber uint64, hostTimeInNanoseconds uint64) ExtensionScheduledOutput {
	rv := objc.Call[ExtensionScheduledOutput](ec, objc.Sel("scheduledOutputWithSequenceNumber:hostTimeInNanoseconds:"), sequenceNumber, hostTimeInNanoseconds)
	return rv
}

// Returns a new scheduled output object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionscheduledoutput/3915886-scheduledoutputwithsequencenumbe?language=objc
func ExtensionScheduledOutput_ScheduledOutputWithSequenceNumberHostTimeInNanoseconds(sequenceNumber uint64, hostTimeInNanoseconds uint64) ExtensionScheduledOutput {
	return ExtensionScheduledOutputClass.ScheduledOutputWithSequenceNumberHostTimeInNanoseconds(sequenceNumber, hostTimeInNanoseconds)
}

func (ec _ExtensionScheduledOutputClass) Alloc() ExtensionScheduledOutput {
	rv := objc.Call[ExtensionScheduledOutput](ec, objc.Sel("alloc"))
	return rv
}

func ExtensionScheduledOutput_Alloc() ExtensionScheduledOutput {
	return ExtensionScheduledOutputClass.Alloc()
}

func (ec _ExtensionScheduledOutputClass) New() ExtensionScheduledOutput {
	rv := objc.Call[ExtensionScheduledOutput](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionScheduledOutput() ExtensionScheduledOutput {
	return ExtensionScheduledOutputClass.New()
}

func (e_ ExtensionScheduledOutput) Init() ExtensionScheduledOutput {
	rv := objc.Call[ExtensionScheduledOutput](e_, objc.Sel("init"))
	return rv
}

// The host time in nanoseconds when the buffer was output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionscheduledoutput/3915884-hosttimeinnanoseconds?language=objc
func (e_ ExtensionScheduledOutput) HostTimeInNanoseconds() uint64 {
	rv := objc.Call[uint64](e_, objc.Sel("hostTimeInNanoseconds"))
	return rv
}

// The buffer sequence number that was output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionscheduledoutput/3915887-sequencenumber?language=objc
func (e_ ExtensionScheduledOutput) SequenceNumber() uint64 {
	rv := objc.Call[uint64](e_, objc.Sel("sequenceNumber"))
	return rv
}
