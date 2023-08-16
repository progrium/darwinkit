// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DatabaseSubscription] class.
var DatabaseSubscriptionClass = _DatabaseSubscriptionClass{objc.GetClass("CKDatabaseSubscription")}

type _DatabaseSubscriptionClass struct {
	objc.Class
}

// An interface definition for the [DatabaseSubscription] class.
type IDatabaseSubscription interface {
	ISubscription
	RecordType() RecordType
	SetRecordType(value RecordType)
}

// A subscription that generates push notifications when CloudKit modifies records in a database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabasesubscription?language=objc
type DatabaseSubscription struct {
	Subscription
}

func DatabaseSubscriptionFrom(ptr unsafe.Pointer) DatabaseSubscription {
	return DatabaseSubscription{
		Subscription: SubscriptionFrom(ptr),
	}
}

func (dc _DatabaseSubscriptionClass) New() DatabaseSubscription {
	rv := objc.Call[DatabaseSubscription](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDatabaseSubscription() DatabaseSubscription {
	return DatabaseSubscriptionClass.New()
}

func (d_ DatabaseSubscription) InitWithSubscriptionID(subscriptionID SubscriptionID) DatabaseSubscription {
	rv := objc.Call[DatabaseSubscription](d_, objc.Sel("initWithSubscriptionID:"), subscriptionID)
	return rv
}

// Creates a named subscription for all records in a database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabasesubscription/1640530-initwithsubscriptionid?language=objc
func DatabaseSubscription_InitWithSubscriptionID(subscriptionID SubscriptionID) DatabaseSubscription {
	return DatabaseSubscriptionClass.Alloc().InitWithSubscriptionID(subscriptionID)
}

func (dc _DatabaseSubscriptionClass) Alloc() DatabaseSubscription {
	rv := objc.Call[DatabaseSubscription](dc, objc.Sel("alloc"))
	return rv
}

func DatabaseSubscription_Alloc() DatabaseSubscription {
	return DatabaseSubscriptionClass.Alloc()
}

func (d_ DatabaseSubscription) Init() DatabaseSubscription {
	rv := objc.Call[DatabaseSubscription](d_, objc.Sel("init"))
	return rv
}

// The type of record that the subscription queries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabasesubscription/1640418-recordtype?language=objc
func (d_ DatabaseSubscription) RecordType() RecordType {
	rv := objc.Call[RecordType](d_, objc.Sel("recordType"))
	return rv
}

// The type of record that the subscription queries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabasesubscription/1640418-recordtype?language=objc
func (d_ DatabaseSubscription) SetRecordType(value RecordType) {
	objc.Call[objc.Void](d_, objc.Sel("setRecordType:"), value)
}
