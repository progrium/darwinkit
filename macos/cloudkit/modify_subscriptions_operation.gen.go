// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ModifySubscriptionsOperation] class.
var ModifySubscriptionsOperationClass = _ModifySubscriptionsOperationClass{objc.GetClass("CKModifySubscriptionsOperation")}

type _ModifySubscriptionsOperationClass struct {
	objc.Class
}

// An interface definition for the [ModifySubscriptionsOperation] class.
type IModifySubscriptionsOperation interface {
	IDatabaseOperation
	PerSubscriptionSaveBlock() func(subscriptionID SubscriptionID, subscription Subscription, error foundation.Error)
	SetPerSubscriptionSaveBlock(value func(subscriptionID SubscriptionID, subscription Subscription, error foundation.Error))
	SubscriptionsToSave() []Subscription
	SetSubscriptionsToSave(value []ISubscription)
	PerSubscriptionDeleteBlock() func(subscriptionID SubscriptionID, error foundation.Error)
	SetPerSubscriptionDeleteBlock(value func(subscriptionID SubscriptionID, error foundation.Error))
	ModifySubscriptionsCompletionBlock() func(savedSubscriptions []Subscription, deletedSubscriptionIDs []SubscriptionID, operationError foundation.Error)
	SetModifySubscriptionsCompletionBlock(value func(savedSubscriptions []Subscription, deletedSubscriptionIDs []SubscriptionID, operationError foundation.Error))
	SubscriptionIDsToDelete() []SubscriptionID
	SetSubscriptionIDsToDelete(value []SubscriptionID)
}

// An operation for modifying one or more subscriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation?language=objc
type ModifySubscriptionsOperation struct {
	DatabaseOperation
}

func ModifySubscriptionsOperationFrom(ptr unsafe.Pointer) ModifySubscriptionsOperation {
	return ModifySubscriptionsOperation{
		DatabaseOperation: DatabaseOperationFrom(ptr),
	}
}

func (m_ ModifySubscriptionsOperation) InitWithSubscriptionsToSaveSubscriptionIDsToDelete(subscriptionsToSave []ISubscription, subscriptionIDsToDelete []SubscriptionID) ModifySubscriptionsOperation {
	rv := objc.Call[ModifySubscriptionsOperation](m_, objc.Sel("initWithSubscriptionsToSave:subscriptionIDsToDelete:"), subscriptionsToSave, subscriptionIDsToDelete)
	return rv
}

// Creates an operation for saving and deleting the specified subscriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/1515015-initwithsubscriptionstosave?language=objc
func NewModifySubscriptionsOperationWithSubscriptionsToSaveSubscriptionIDsToDelete(subscriptionsToSave []ISubscription, subscriptionIDsToDelete []SubscriptionID) ModifySubscriptionsOperation {
	instance := ModifySubscriptionsOperationClass.Alloc().InitWithSubscriptionsToSaveSubscriptionIDsToDelete(subscriptionsToSave, subscriptionIDsToDelete)
	instance.Autorelease()
	return instance
}

func (m_ ModifySubscriptionsOperation) Init() ModifySubscriptionsOperation {
	rv := objc.Call[ModifySubscriptionsOperation](m_, objc.Sel("init"))
	return rv
}

func (mc _ModifySubscriptionsOperationClass) Alloc() ModifySubscriptionsOperation {
	rv := objc.Call[ModifySubscriptionsOperation](mc, objc.Sel("alloc"))
	return rv
}

func ModifySubscriptionsOperation_Alloc() ModifySubscriptionsOperation {
	return ModifySubscriptionsOperationClass.Alloc()
}

func (mc _ModifySubscriptionsOperationClass) New() ModifySubscriptionsOperation {
	rv := objc.Call[ModifySubscriptionsOperation](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewModifySubscriptionsOperation() ModifySubscriptionsOperation {
	return ModifySubscriptionsOperationClass.New()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/3793707-persubscriptionsaveblock?language=objc
func (m_ ModifySubscriptionsOperation) PerSubscriptionSaveBlock() func(subscriptionID SubscriptionID, subscription Subscription, error foundation.Error) {
	rv := objc.Call[func(subscriptionID SubscriptionID, subscription Subscription, error foundation.Error)](m_, objc.Sel("perSubscriptionSaveBlock"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/3793707-persubscriptionsaveblock?language=objc
func (m_ ModifySubscriptionsOperation) SetPerSubscriptionSaveBlock(value func(subscriptionID SubscriptionID, subscription Subscription, error foundation.Error)) {
	objc.Call[objc.Void](m_, objc.Sel("setPerSubscriptionSaveBlock:"), value)
}

// The subscriptions to save to the database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/1515135-subscriptionstosave?language=objc
func (m_ ModifySubscriptionsOperation) SubscriptionsToSave() []Subscription {
	rv := objc.Call[[]Subscription](m_, objc.Sel("subscriptionsToSave"))
	return rv
}

// The subscriptions to save to the database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/1515135-subscriptionstosave?language=objc
func (m_ ModifySubscriptionsOperation) SetSubscriptionsToSave(value []ISubscription) {
	objc.Call[objc.Void](m_, objc.Sel("setSubscriptionsToSave:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/3793706-persubscriptiondeleteblock?language=objc
func (m_ ModifySubscriptionsOperation) PerSubscriptionDeleteBlock() func(subscriptionID SubscriptionID, error foundation.Error) {
	rv := objc.Call[func(subscriptionID SubscriptionID, error foundation.Error)](m_, objc.Sel("perSubscriptionDeleteBlock"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/3793706-persubscriptiondeleteblock?language=objc
func (m_ ModifySubscriptionsOperation) SetPerSubscriptionDeleteBlock(value func(subscriptionID SubscriptionID, error foundation.Error)) {
	objc.Call[objc.Void](m_, objc.Sel("setPerSubscriptionDeleteBlock:"), value)
}

// The block to execute after the operation modifies the subscriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/1515288-modifysubscriptionscompletionblo?language=objc
func (m_ ModifySubscriptionsOperation) ModifySubscriptionsCompletionBlock() func(savedSubscriptions []Subscription, deletedSubscriptionIDs []SubscriptionID, operationError foundation.Error) {
	rv := objc.Call[func(savedSubscriptions []Subscription, deletedSubscriptionIDs []SubscriptionID, operationError foundation.Error)](m_, objc.Sel("modifySubscriptionsCompletionBlock"))
	return rv
}

// The block to execute after the operation modifies the subscriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/1515288-modifysubscriptionscompletionblo?language=objc
func (m_ ModifySubscriptionsOperation) SetModifySubscriptionsCompletionBlock(value func(savedSubscriptions []Subscription, deletedSubscriptionIDs []SubscriptionID, operationError foundation.Error)) {
	objc.Call[objc.Void](m_, objc.Sel("setModifySubscriptionsCompletionBlock:"), value)
}

// The IDs of the subscriptions that you want to delete. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/1514892-subscriptionidstodelete?language=objc
func (m_ ModifySubscriptionsOperation) SubscriptionIDsToDelete() []SubscriptionID {
	rv := objc.Call[[]SubscriptionID](m_, objc.Sel("subscriptionIDsToDelete"))
	return rv
}

// The IDs of the subscriptions that you want to delete. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/1514892-subscriptionidstodelete?language=objc
func (m_ ModifySubscriptionsOperation) SetSubscriptionIDsToDelete(value []SubscriptionID) {
	objc.Call[objc.Void](m_, objc.Sel("setSubscriptionIDsToDelete:"), value)
}
