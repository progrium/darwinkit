// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DatabaseOperation] class.
var DatabaseOperationClass = _DatabaseOperationClass{objc.GetClass("CKDatabaseOperation")}

type _DatabaseOperationClass struct {
	objc.Class
}

// An interface definition for the [DatabaseOperation] class.
type IDatabaseOperation interface {
	IOperation
	Database() Database
	SetDatabase(value IDatabase)
}

// The abstract base class for operations that act upon databases in CloudKit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabaseoperation?language=objc
type DatabaseOperation struct {
	Operation
}

func DatabaseOperationFrom(ptr unsafe.Pointer) DatabaseOperation {
	return DatabaseOperation{
		Operation: OperationFrom(ptr),
	}
}

func (dc _DatabaseOperationClass) Alloc() DatabaseOperation {
	rv := objc.Call[DatabaseOperation](dc, objc.Sel("alloc"))
	return rv
}

func DatabaseOperation_Alloc() DatabaseOperation {
	return DatabaseOperationClass.Alloc()
}

func (dc _DatabaseOperationClass) New() DatabaseOperation {
	rv := objc.Call[DatabaseOperation](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDatabaseOperation() DatabaseOperation {
	return DatabaseOperationClass.New()
}

func (d_ DatabaseOperation) Init() DatabaseOperation {
	rv := objc.Call[DatabaseOperation](d_, objc.Sel("init"))
	return rv
}

// The database that the operation uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabaseoperation/1515274-database?language=objc
func (d_ DatabaseOperation) Database() Database {
	rv := objc.Call[Database](d_, objc.Sel("database"))
	return rv
}

// The database that the operation uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabaseoperation/1515274-database?language=objc
func (d_ DatabaseOperation) SetDatabase(value IDatabase) {
	objc.Call[objc.Void](d_, objc.Sel("setDatabase:"), objc.Ptr(value))
}
