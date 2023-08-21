// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RecordID] class.
var RecordIDClass = _RecordIDClass{objc.GetClass("CKRecordID")}

type _RecordIDClass struct {
	objc.Class
}

// An interface definition for the [RecordID] class.
type IRecordID interface {
	objc.IObject
	ZoneID() RecordZoneID
	RecordName() string
}

// An object that uniquely identifies a record in a database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordid?language=objc
type RecordID struct {
	objc.Object
}

func RecordIDFrom(ptr unsafe.Pointer) RecordID {
	return RecordID{
		Object: objc.ObjectFrom(ptr),
	}
}

func (r_ RecordID) InitWithRecordName(recordName string) RecordID {
	rv := objc.Call[RecordID](r_, objc.Sel("initWithRecordName:"), recordName)
	return rv
}

// Creates a new record ID with the specified name in the default zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordid/1500975-initwithrecordname?language=objc
func NewRecordIDWithRecordName(recordName string) RecordID {
	instance := RecordIDClass.Alloc().InitWithRecordName(recordName)
	instance.Autorelease()
	return instance
}

func (rc _RecordIDClass) Alloc() RecordID {
	rv := objc.Call[RecordID](rc, objc.Sel("alloc"))
	return rv
}

func RecordID_Alloc() RecordID {
	return RecordIDClass.Alloc()
}

func (rc _RecordIDClass) New() RecordID {
	rv := objc.Call[RecordID](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRecordID() RecordID {
	return RecordIDClass.New()
}

func (r_ RecordID) Init() RecordID {
	rv := objc.Call[RecordID](r_, objc.Sel("init"))
	return rv
}

// The ID of the zone that contains the record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordid/1500969-zoneid?language=objc
func (r_ RecordID) ZoneID() RecordZoneID {
	rv := objc.Call[RecordZoneID](r_, objc.Sel("zoneID"))
	return rv
}

// The unique name of the record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordid/1500973-recordname?language=objc
func (r_ RecordID) RecordName() string {
	rv := objc.Call[string](r_, objc.Sel("recordName"))
	return rv
}
