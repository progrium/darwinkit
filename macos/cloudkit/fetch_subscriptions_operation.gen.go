// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchSubscriptionsOperation] class.
var FetchSubscriptionsOperationClass = _FetchSubscriptionsOperationClass{objc.GetClass("CKFetchSubscriptionsOperation")}

type _FetchSubscriptionsOperationClass struct {
	objc.Class
}

// An interface definition for the [FetchSubscriptionsOperation] class.
type IFetchSubscriptionsOperation interface {
	IDatabaseOperation
	FetchSubscriptionCompletionBlock() func(subscriptionsBySubscriptionID map[SubscriptionID]Subscription, operationError foundation.Error)
	SetFetchSubscriptionCompletionBlock(value func(subscriptionsBySubscriptionID map[SubscriptionID]Subscription, operationError foundation.Error))
	PerSubscriptionCompletionBlock() func(subscriptionID SubscriptionID, subscription Subscription, error foundation.Error)
	SetPerSubscriptionCompletionBlock(value func(subscriptionID SubscriptionID, subscription Subscription, error foundation.Error))
	SubscriptionIDs() []SubscriptionID
	SetSubscriptionIDs(value []SubscriptionID)
}

// An operation for fetching subscriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation?language=objc
type FetchSubscriptionsOperation struct {
	DatabaseOperation
}

func FetchSubscriptionsOperationFrom(ptr unsafe.Pointer) FetchSubscriptionsOperation {
	return FetchSubscriptionsOperation{
		DatabaseOperation: DatabaseOperationFrom(ptr),
	}
}

func (fc _FetchSubscriptionsOperationClass) FetchAllSubscriptionsOperation() FetchSubscriptionsOperation {
	rv := objc.Call[FetchSubscriptionsOperation](fc, objc.Sel("fetchAllSubscriptionsOperation"))
	return rv
}

// Returns an operation that fetches all of the user’s subscriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/1515282-fetchallsubscriptionsoperation?language=objc
func FetchSubscriptionsOperation_FetchAllSubscriptionsOperation() FetchSubscriptionsOperation {
	return FetchSubscriptionsOperationClass.FetchAllSubscriptionsOperation()
}

func (f_ FetchSubscriptionsOperation) InitWithSubscriptionIDs(subscriptionIDs []SubscriptionID) FetchSubscriptionsOperation {
	rv := objc.Call[FetchSubscriptionsOperation](f_, objc.Sel("initWithSubscriptionIDs:"), subscriptionIDs)
	return rv
}

// Creates an operation for fetching the specified subscriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/1515157-initwithsubscriptionids?language=objc
func NewFetchSubscriptionsOperationWithSubscriptionIDs(subscriptionIDs []SubscriptionID) FetchSubscriptionsOperation {
	instance := FetchSubscriptionsOperationClass.Alloc().InitWithSubscriptionIDs(subscriptionIDs)
	instance.Autorelease()
	return instance
}

func (f_ FetchSubscriptionsOperation) Init() FetchSubscriptionsOperation {
	rv := objc.Call[FetchSubscriptionsOperation](f_, objc.Sel("init"))
	return rv
}

func (fc _FetchSubscriptionsOperationClass) Alloc() FetchSubscriptionsOperation {
	rv := objc.Call[FetchSubscriptionsOperation](fc, objc.Sel("alloc"))
	return rv
}

func FetchSubscriptionsOperation_Alloc() FetchSubscriptionsOperation {
	return FetchSubscriptionsOperationClass.Alloc()
}

func (fc _FetchSubscriptionsOperationClass) New() FetchSubscriptionsOperation {
	rv := objc.Call[FetchSubscriptionsOperation](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchSubscriptionsOperation() FetchSubscriptionsOperation {
	return FetchSubscriptionsOperationClass.New()
}

// The block to execute after the operation fetches the subscriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/1515261-fetchsubscriptioncompletionblock?language=objc
func (f_ FetchSubscriptionsOperation) FetchSubscriptionCompletionBlock() func(subscriptionsBySubscriptionID map[SubscriptionID]Subscription, operationError foundation.Error) {
	rv := objc.Call[func(subscriptionsBySubscriptionID map[SubscriptionID]Subscription, operationError foundation.Error)](f_, objc.Sel("fetchSubscriptionCompletionBlock"))
	return rv
}

// The block to execute after the operation fetches the subscriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/1515261-fetchsubscriptioncompletionblock?language=objc
func (f_ FetchSubscriptionsOperation) SetFetchSubscriptionCompletionBlock(value func(subscriptionsBySubscriptionID map[SubscriptionID]Subscription, operationError foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("setFetchSubscriptionCompletionBlock:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/3793701-persubscriptioncompletionblock?language=objc
func (f_ FetchSubscriptionsOperation) PerSubscriptionCompletionBlock() func(subscriptionID SubscriptionID, subscription Subscription, error foundation.Error) {
	rv := objc.Call[func(subscriptionID SubscriptionID, subscription Subscription, error foundation.Error)](f_, objc.Sel("perSubscriptionCompletionBlock"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/3793701-persubscriptioncompletionblock?language=objc
func (f_ FetchSubscriptionsOperation) SetPerSubscriptionCompletionBlock(value func(subscriptionID SubscriptionID, subscription Subscription, error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("setPerSubscriptionCompletionBlock:"), value)
}

// The IDs of the subscriptions to fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/1515011-subscriptionids?language=objc
func (f_ FetchSubscriptionsOperation) SubscriptionIDs() []SubscriptionID {
	rv := objc.Call[[]SubscriptionID](f_, objc.Sel("subscriptionIDs"))
	return rv
}

// The IDs of the subscriptions to fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/1515011-subscriptionids?language=objc
func (f_ FetchSubscriptionsOperation) SetSubscriptionIDs(value []SubscriptionID) {
	objc.Call[objc.Void](f_, objc.Sel("setSubscriptionIDs:"), value)
}
