// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchWebAuthTokenOperation] class.
var FetchWebAuthTokenOperationClass = _FetchWebAuthTokenOperationClass{objc.GetClass("CKFetchWebAuthTokenOperation")}

type _FetchWebAuthTokenOperationClass struct {
	objc.Class
}

// An interface definition for the [FetchWebAuthTokenOperation] class.
type IFetchWebAuthTokenOperation interface {
	IDatabaseOperation
	FetchWebAuthTokenCompletionBlock() func(webAuthToken string, operationError foundation.Error)
	SetFetchWebAuthTokenCompletionBlock(value func(webAuthToken string, operationError foundation.Error))
	APIToken() string
	SetAPIToken(value string)
}

// An operation that creates an authentication token for use with CloudKit web services. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchwebauthtokenoperation?language=objc
type FetchWebAuthTokenOperation struct {
	DatabaseOperation
}

func FetchWebAuthTokenOperationFrom(ptr unsafe.Pointer) FetchWebAuthTokenOperation {
	return FetchWebAuthTokenOperation{
		DatabaseOperation: DatabaseOperationFrom(ptr),
	}
}

func (f_ FetchWebAuthTokenOperation) InitWithAPIToken(APIToken string) FetchWebAuthTokenOperation {
	rv := objc.Call[FetchWebAuthTokenOperation](f_, objc.Sel("initWithAPIToken:"), APIToken)
	return rv
}

// Creates a fetch operation for the specified API token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchwebauthtokenoperation/1515266-initwithapitoken?language=objc
func FetchWebAuthTokenOperation_InitWithAPIToken(APIToken string) FetchWebAuthTokenOperation {
	return FetchWebAuthTokenOperationClass.Alloc().InitWithAPIToken(APIToken)
}

func (f_ FetchWebAuthTokenOperation) Init() FetchWebAuthTokenOperation {
	rv := objc.Call[FetchWebAuthTokenOperation](f_, objc.Sel("init"))
	return rv
}

func (fc _FetchWebAuthTokenOperationClass) Alloc() FetchWebAuthTokenOperation {
	rv := objc.Call[FetchWebAuthTokenOperation](fc, objc.Sel("alloc"))
	return rv
}

func FetchWebAuthTokenOperation_Alloc() FetchWebAuthTokenOperation {
	return FetchWebAuthTokenOperationClass.Alloc()
}

func (fc _FetchWebAuthTokenOperationClass) New() FetchWebAuthTokenOperation {
	rv := objc.Call[FetchWebAuthTokenOperation](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchWebAuthTokenOperation() FetchWebAuthTokenOperation {
	return FetchWebAuthTokenOperationClass.New()
}

// The block to execute when the operation finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchwebauthtokenoperation/1514980-fetchwebauthtokencompletionblock?language=objc
func (f_ FetchWebAuthTokenOperation) FetchWebAuthTokenCompletionBlock() func(webAuthToken string, operationError foundation.Error) {
	rv := objc.Call[func(webAuthToken string, operationError foundation.Error)](f_, objc.Sel("fetchWebAuthTokenCompletionBlock"))
	return rv
}

// The block to execute when the operation finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchwebauthtokenoperation/1514980-fetchwebauthtokencompletionblock?language=objc
func (f_ FetchWebAuthTokenOperation) SetFetchWebAuthTokenCompletionBlock(value func(webAuthToken string, operationError foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("setFetchWebAuthTokenCompletionBlock:"), value)
}

// The API token that allows access to an app’s container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchwebauthtokenoperation/1515095-apitoken?language=objc
func (f_ FetchWebAuthTokenOperation) APIToken() string {
	rv := objc.Call[string](f_, objc.Sel("APIToken"))
	return rv
}

// The API token that allows access to an app’s container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchwebauthtokenoperation/1515095-apitoken?language=objc
func (f_ FetchWebAuthTokenOperation) SetAPIToken(value string) {
	objc.Call[objc.Void](f_, objc.Sel("setAPIToken:"), value)
}
