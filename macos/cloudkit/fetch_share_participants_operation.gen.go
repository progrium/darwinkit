// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchShareParticipantsOperation] class.
var FetchShareParticipantsOperationClass = _FetchShareParticipantsOperationClass{objc.GetClass("CKFetchShareParticipantsOperation")}

type _FetchShareParticipantsOperationClass struct {
	objc.Class
}

// An interface definition for the [FetchShareParticipantsOperation] class.
type IFetchShareParticipantsOperation interface {
	IOperation
	UserIdentityLookupInfos() []UserIdentityLookupInfo
	SetUserIdentityLookupInfos(value []IUserIdentityLookupInfo)
	FetchShareParticipantsCompletionBlock() func(operationError foundation.Error)
	SetFetchShareParticipantsCompletionBlock(value func(operationError foundation.Error))
	PerShareParticipantCompletionBlock() func(lookupInfo UserIdentityLookupInfo, participant ShareParticipant, error foundation.Error)
	SetPerShareParticipantCompletionBlock(value func(lookupInfo UserIdentityLookupInfo, participant ShareParticipant, error foundation.Error))
}

// An operation that converts user identities into share participants. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation?language=objc
type FetchShareParticipantsOperation struct {
	Operation
}

func FetchShareParticipantsOperationFrom(ptr unsafe.Pointer) FetchShareParticipantsOperation {
	return FetchShareParticipantsOperation{
		Operation: OperationFrom(ptr),
	}
}

func (f_ FetchShareParticipantsOperation) InitWithUserIdentityLookupInfos(userIdentityLookupInfos []IUserIdentityLookupInfo) FetchShareParticipantsOperation {
	rv := objc.Call[FetchShareParticipantsOperation](f_, objc.Sel("initWithUserIdentityLookupInfos:"), userIdentityLookupInfos)
	return rv
}

// Creates an operation for generating share participants from the specified user data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/1640471-initwithuseridentitylookupinfos?language=objc
func FetchShareParticipantsOperation_InitWithUserIdentityLookupInfos(userIdentityLookupInfos []IUserIdentityLookupInfo) FetchShareParticipantsOperation {
	return FetchShareParticipantsOperationClass.Alloc().InitWithUserIdentityLookupInfos(userIdentityLookupInfos)
}

func (f_ FetchShareParticipantsOperation) Init() FetchShareParticipantsOperation {
	rv := objc.Call[FetchShareParticipantsOperation](f_, objc.Sel("init"))
	return rv
}

func (fc _FetchShareParticipantsOperationClass) Alloc() FetchShareParticipantsOperation {
	rv := objc.Call[FetchShareParticipantsOperation](fc, objc.Sel("alloc"))
	return rv
}

func FetchShareParticipantsOperation_Alloc() FetchShareParticipantsOperation {
	return FetchShareParticipantsOperationClass.Alloc()
}

func (fc _FetchShareParticipantsOperationClass) New() FetchShareParticipantsOperation {
	rv := objc.Call[FetchShareParticipantsOperation](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchShareParticipantsOperation() FetchShareParticipantsOperation {
	return FetchShareParticipantsOperationClass.New()
}

// The user data for the participants. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/1640380-useridentitylookupinfos?language=objc
func (f_ FetchShareParticipantsOperation) UserIdentityLookupInfos() []UserIdentityLookupInfo {
	rv := objc.Call[[]UserIdentityLookupInfo](f_, objc.Sel("userIdentityLookupInfos"))
	return rv
}

// The user data for the participants. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/1640380-useridentitylookupinfos?language=objc
func (f_ FetchShareParticipantsOperation) SetUserIdentityLookupInfos(value []IUserIdentityLookupInfo) {
	objc.Call[objc.Void](f_, objc.Sel("setUserIdentityLookupInfos:"), value)
}

// The block to execute when the operation finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/1640529-fetchshareparticipantscompletion?language=objc
func (f_ FetchShareParticipantsOperation) FetchShareParticipantsCompletionBlock() func(operationError foundation.Error) {
	rv := objc.Call[func(operationError foundation.Error)](f_, objc.Sel("fetchShareParticipantsCompletionBlock"))
	return rv
}

// The block to execute when the operation finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/1640529-fetchshareparticipantscompletion?language=objc
func (f_ FetchShareParticipantsOperation) SetFetchShareParticipantsCompletionBlock(value func(operationError foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("setFetchShareParticipantsCompletionBlock:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/3793700-pershareparticipantcompletionblo?language=objc
func (f_ FetchShareParticipantsOperation) PerShareParticipantCompletionBlock() func(lookupInfo UserIdentityLookupInfo, participant ShareParticipant, error foundation.Error) {
	rv := objc.Call[func(lookupInfo UserIdentityLookupInfo, participant ShareParticipant, error foundation.Error)](f_, objc.Sel("perShareParticipantCompletionBlock"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/3793700-pershareparticipantcompletionblo?language=objc
func (f_ FetchShareParticipantsOperation) SetPerShareParticipantCompletionBlock(value func(lookupInfo UserIdentityLookupInfo, participant ShareParticipant, error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("setPerShareParticipantCompletionBlock:"), value)
}
