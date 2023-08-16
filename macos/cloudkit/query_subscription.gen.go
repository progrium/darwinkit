// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [QuerySubscription] class.
var QuerySubscriptionClass = _QuerySubscriptionClass{objc.GetClass("CKQuerySubscription")}

type _QuerySubscriptionClass struct {
	objc.Class
}

// An interface definition for the [QuerySubscription] class.
type IQuerySubscription interface {
	ISubscription
	QuerySubscriptionOptions() QuerySubscriptionOptions
	ZoneID() RecordZoneID
	SetZoneID(value IRecordZoneID)
	Predicate() foundation.Predicate
	RecordType() RecordType
}

// A subscription that generates push notifications when CloudKit modifies records that match a predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerysubscription?language=objc
type QuerySubscription struct {
	Subscription
}

func QuerySubscriptionFrom(ptr unsafe.Pointer) QuerySubscription {
	return QuerySubscription{
		Subscription: SubscriptionFrom(ptr),
	}
}

func (qc _QuerySubscriptionClass) Alloc() QuerySubscription {
	rv := objc.Call[QuerySubscription](qc, objc.Sel("alloc"))
	return rv
}

func QuerySubscription_Alloc() QuerySubscription {
	return QuerySubscriptionClass.Alloc()
}

func (qc _QuerySubscriptionClass) New() QuerySubscription {
	rv := objc.Call[QuerySubscription](qc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewQuerySubscription() QuerySubscription {
	return QuerySubscriptionClass.New()
}

func (q_ QuerySubscription) Init() QuerySubscription {
	rv := objc.Call[QuerySubscription](q_, objc.Sel("init"))
	return rv
}

// Options that define the behavior of the subscription. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerysubscription/1640414-querysubscriptionoptions?language=objc
func (q_ QuerySubscription) QuerySubscriptionOptions() QuerySubscriptionOptions {
	rv := objc.Call[QuerySubscriptionOptions](q_, objc.Sel("querySubscriptionOptions"))
	return rv
}

// The ID of the record zone that the subscription queries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerysubscription/1640390-zoneid?language=objc
func (q_ QuerySubscription) ZoneID() RecordZoneID {
	rv := objc.Call[RecordZoneID](q_, objc.Sel("zoneID"))
	return rv
}

// The ID of the record zone that the subscription queries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerysubscription/1640390-zoneid?language=objc
func (q_ QuerySubscription) SetZoneID(value IRecordZoneID) {
	objc.Call[objc.Void](q_, objc.Sel("setZoneID:"), objc.Ptr(value))
}

// The matching criteria to apply to records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerysubscription/1640485-predicate?language=objc
func (q_ QuerySubscription) Predicate() foundation.Predicate {
	rv := objc.Call[foundation.Predicate](q_, objc.Sel("predicate"))
	return rv
}

// The type of record that the subscription queries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerysubscription/1640393-recordtype?language=objc
func (q_ QuerySubscription) RecordType() RecordType {
	rv := objc.Call[RecordType](q_, objc.Sel("recordType"))
	return rv
}
