// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Reference] class.
var ReferenceClass = _ReferenceClass{objc.GetClass("CKReference")}

type _ReferenceClass struct {
	objc.Class
}

// An interface definition for the [Reference] class.
type IReference interface {
	objc.IObject
	ReferenceAction() ReferenceAction
	RecordID() RecordID
}

// A relationship between two records in a record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckreference?language=objc
type Reference struct {
	objc.Object
}

func ReferenceFrom(ptr unsafe.Pointer) Reference {
	return Reference{
		Object: objc.ObjectFrom(ptr),
	}
}

func (r_ Reference) InitWithRecordIDAction(recordID IRecordID, action ReferenceAction) Reference {
	rv := objc.Call[Reference](r_, objc.Sel("initWithRecordID:action:"), objc.Ptr(recordID), action)
	return rv
}

// Creates a reference object that points to the record with the specified ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckreference/1515280-initwithrecordid?language=objc
func Reference_InitWithRecordIDAction(recordID IRecordID, action ReferenceAction) Reference {
	return ReferenceClass.Alloc().InitWithRecordIDAction(recordID, action)
}

func (r_ Reference) InitWithRecordAction(record IRecord, action ReferenceAction) Reference {
	rv := objc.Call[Reference](r_, objc.Sel("initWithRecord:action:"), objc.Ptr(record), action)
	return rv
}

// Creates a reference object that points to the specified record object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckreference/1515312-initwithrecord?language=objc
func Reference_InitWithRecordAction(record IRecord, action ReferenceAction) Reference {
	return ReferenceClass.Alloc().InitWithRecordAction(record, action)
}

func (rc _ReferenceClass) Alloc() Reference {
	rv := objc.Call[Reference](rc, objc.Sel("alloc"))
	return rv
}

func Reference_Alloc() Reference {
	return ReferenceClass.Alloc()
}

func (rc _ReferenceClass) New() Reference {
	rv := objc.Call[Reference](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewReference() Reference {
	return ReferenceClass.New()
}

func (r_ Reference) Init() Reference {
	rv := objc.Call[Reference](r_, objc.Sel("init"))
	return rv
}

// The ownership behavior for the records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckreference/1514900-referenceaction?language=objc
func (r_ Reference) ReferenceAction() ReferenceAction {
	rv := objc.Call[ReferenceAction](r_, objc.Sel("referenceAction"))
	return rv
}

// The ID of the referenced record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckreference/1514956-recordid?language=objc
func (r_ Reference) RecordID() RecordID {
	rv := objc.Call[RecordID](r_, objc.Sel("recordID"))
	return rv
}
