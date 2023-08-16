// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersistentHistoryTransaction] class.
var PersistentHistoryTransactionClass = _PersistentHistoryTransactionClass{objc.GetClass("NSPersistentHistoryTransaction")}

type _PersistentHistoryTransactionClass struct {
	objc.Class
}

// An interface definition for the [PersistentHistoryTransaction] class.
type IPersistentHistoryTransaction interface {
	objc.IObject
	ObjectIDNotification() foundation.Notification
	Token() PersistentHistoryToken
	BundleID() string
	TransactionNumber() int64
	StoreID() string
	Author() string
	ProcessID() string
	Changes() []PersistentHistoryChange
	Timestamp() foundation.Date
	ContextName() string
}

// A set of changes in the persistent history based on a context save or batch operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction?language=objc
type PersistentHistoryTransaction struct {
	objc.Object
}

func PersistentHistoryTransactionFrom(ptr unsafe.Pointer) PersistentHistoryTransaction {
	return PersistentHistoryTransaction{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PersistentHistoryTransactionClass) Alloc() PersistentHistoryTransaction {
	rv := objc.Call[PersistentHistoryTransaction](pc, objc.Sel("alloc"))
	return rv
}

func PersistentHistoryTransaction_Alloc() PersistentHistoryTransaction {
	return PersistentHistoryTransactionClass.Alloc()
}

func (pc _PersistentHistoryTransactionClass) New() PersistentHistoryTransaction {
	rv := objc.Call[PersistentHistoryTransaction](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentHistoryTransaction() PersistentHistoryTransaction {
	return PersistentHistoryTransactionClass.New()
}

func (p_ PersistentHistoryTransaction) Init() PersistentHistoryTransaction {
	rv := objc.Call[PersistentHistoryTransaction](p_, objc.Sel("init"))
	return rv
}

// Obtains a notification for use in merging the transaction’s changes into a managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/2894946-objectidnotification?language=objc
func (p_ PersistentHistoryTransaction) ObjectIDNotification() foundation.Notification {
	rv := objc.Call[foundation.Notification](p_, objc.Sel("objectIDNotification"))
	return rv
}

// Requests an entity description using the provided context for the managed object type affected by the transaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/3240595-entitydescriptionwithcontext?language=objc
func (pc _PersistentHistoryTransactionClass) EntityDescriptionWithContext(context IManagedObjectContext) EntityDescription {
	rv := objc.Call[EntityDescription](pc, objc.Sel("entityDescriptionWithContext:"), objc.Ptr(context))
	return rv
}

// Requests an entity description using the provided context for the managed object type affected by the transaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/3240595-entitydescriptionwithcontext?language=objc
func PersistentHistoryTransaction_EntityDescriptionWithContext(context IManagedObjectContext) EntityDescription {
	return PersistentHistoryTransactionClass.EntityDescriptionWithContext(context)
}

// A fetch request that has the persistent history transaction as the entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/3240596-fetchrequest?language=objc
func (pc _PersistentHistoryTransactionClass) FetchRequest() FetchRequest {
	rv := objc.Call[FetchRequest](pc, objc.Sel("fetchRequest"))
	return rv
}

// A fetch request that has the persistent history transaction as the entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/3240596-fetchrequest?language=objc
func PersistentHistoryTransaction_FetchRequest() FetchRequest {
	return PersistentHistoryTransactionClass.FetchRequest()
}

// The token that represents this transaction in the persistent history. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/2894938-token?language=objc
func (p_ PersistentHistoryTransaction) Token() PersistentHistoryToken {
	rv := objc.Call[PersistentHistoryToken](p_, objc.Sel("token"))
	return rv
}

// The originating bundle’s identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/2894943-bundleid?language=objc
func (p_ PersistentHistoryTransaction) BundleID() string {
	rv := objc.Call[string](p_, objc.Sel("bundleID"))
	return rv
}

// The transaction’s numeric identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/2894945-transactionnumber?language=objc
func (p_ PersistentHistoryTransaction) TransactionNumber() int64 {
	rv := objc.Call[int64](p_, objc.Sel("transactionNumber"))
	return rv
}

// The originating store’s identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/2894941-storeid?language=objc
func (p_ PersistentHistoryTransaction) StoreID() string {
	rv := objc.Call[string](p_, objc.Sel("storeID"))
	return rv
}

// A granular description of the context that made the persistent history change, if available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/2894947-author?language=objc
func (p_ PersistentHistoryTransaction) Author() string {
	rv := objc.Call[string](p_, objc.Sel("author"))
	return rv
}

// The originating process’s identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/2894939-processid?language=objc
func (p_ PersistentHistoryTransaction) ProcessID() string {
	rv := objc.Call[string](p_, objc.Sel("processID"))
	return rv
}

// The entity description of the persistent history transaction entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/3240594-entitydescription?language=objc
func (pc _PersistentHistoryTransactionClass) EntityDescription() EntityDescription {
	rv := objc.Call[EntityDescription](pc, objc.Sel("entityDescription"))
	return rv
}

// The entity description of the persistent history transaction entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/3240594-entitydescription?language=objc
func PersistentHistoryTransaction_EntityDescription() EntityDescription {
	return PersistentHistoryTransactionClass.EntityDescription()
}

// The array of persistent history changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/2894948-changes?language=objc
func (p_ PersistentHistoryTransaction) Changes() []PersistentHistoryChange {
	rv := objc.Call[[]PersistentHistoryChange](p_, objc.Sel("changes"))
	return rv
}

// The date of the persistent history change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/2894940-timestamp?language=objc
func (p_ PersistentHistoryTransaction) Timestamp() foundation.Date {
	rv := objc.Call[foundation.Date](p_, objc.Sel("timestamp"))
	return rv
}

// The originating context’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/2894944-contextname?language=objc
func (p_ PersistentHistoryTransaction) ContextName() string {
	rv := objc.Call[string](p_, objc.Sel("contextName"))
	return rv
}
