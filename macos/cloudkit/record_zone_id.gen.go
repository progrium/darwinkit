// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RecordZoneID] class.
var RecordZoneIDClass = _RecordZoneIDClass{objc.GetClass("CKRecordZoneID")}

type _RecordZoneIDClass struct {
	objc.Class
}

// An interface definition for the [RecordZoneID] class.
type IRecordZoneID interface {
	objc.IObject
	ZoneName() string
	OwnerName() string
}

// An object that uniquely identifies a record zone in a database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzoneid?language=objc
type RecordZoneID struct {
	objc.Object
}

func RecordZoneIDFrom(ptr unsafe.Pointer) RecordZoneID {
	return RecordZoneID{
		Object: objc.ObjectFrom(ptr),
	}
}

func (r_ RecordZoneID) InitWithZoneNameOwnerName(zoneName string, ownerName string) RecordZoneID {
	rv := objc.Call[RecordZoneID](r_, objc.Sel("initWithZoneName:ownerName:"), zoneName, ownerName)
	return rv
}

// Creates a record zone ID with the specified name and owner. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzoneid/1508089-initwithzonename?language=objc
func RecordZoneID_InitWithZoneNameOwnerName(zoneName string, ownerName string) RecordZoneID {
	return RecordZoneIDClass.Alloc().InitWithZoneNameOwnerName(zoneName, ownerName)
}

func (rc _RecordZoneIDClass) Alloc() RecordZoneID {
	rv := objc.Call[RecordZoneID](rc, objc.Sel("alloc"))
	return rv
}

func RecordZoneID_Alloc() RecordZoneID {
	return RecordZoneIDClass.Alloc()
}

func (rc _RecordZoneIDClass) New() RecordZoneID {
	rv := objc.Call[RecordZoneID](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRecordZoneID() RecordZoneID {
	return RecordZoneIDClass.New()
}

func (r_ RecordZoneID) Init() RecordZoneID {
	rv := objc.Call[RecordZoneID](r_, objc.Sel("init"))
	return rv
}

// The unique name of the record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzoneid/1508094-zonename?language=objc
func (r_ RecordZoneID) ZoneName() string {
	rv := objc.Call[string](r_, objc.Sel("zoneName"))
	return rv
}

// The ID of the user who owns the record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzoneid/1508096-ownername?language=objc
func (r_ RecordZoneID) OwnerName() string {
	rv := objc.Call[string](r_, objc.Sel("ownerName"))
	return rv
}
