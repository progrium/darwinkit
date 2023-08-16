// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchRecordZoneChangesConfiguration] class.
var FetchRecordZoneChangesConfigurationClass = _FetchRecordZoneChangesConfigurationClass{objc.GetClass("CKFetchRecordZoneChangesConfiguration")}

type _FetchRecordZoneChangesConfigurationClass struct {
	objc.Class
}

// An interface definition for the [FetchRecordZoneChangesConfiguration] class.
type IFetchRecordZoneChangesConfiguration interface {
	objc.IObject
	PreviousServerChangeToken() ServerChangeToken
	SetPreviousServerChangeToken(value IServerChangeToken)
	ResultsLimit() uint
	SetResultsLimit(value uint)
	DesiredKeys() []RecordFieldKey
	SetDesiredKeys(value []RecordFieldKey)
}

// A configuration object that describes the information to fetch from a record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesconfiguration?language=objc
type FetchRecordZoneChangesConfiguration struct {
	objc.Object
}

func FetchRecordZoneChangesConfigurationFrom(ptr unsafe.Pointer) FetchRecordZoneChangesConfiguration {
	return FetchRecordZoneChangesConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FetchRecordZoneChangesConfigurationClass) Alloc() FetchRecordZoneChangesConfiguration {
	rv := objc.Call[FetchRecordZoneChangesConfiguration](fc, objc.Sel("alloc"))
	return rv
}

func FetchRecordZoneChangesConfiguration_Alloc() FetchRecordZoneChangesConfiguration {
	return FetchRecordZoneChangesConfigurationClass.Alloc()
}

func (fc _FetchRecordZoneChangesConfigurationClass) New() FetchRecordZoneChangesConfiguration {
	rv := objc.Call[FetchRecordZoneChangesConfiguration](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchRecordZoneChangesConfiguration() FetchRecordZoneChangesConfiguration {
	return FetchRecordZoneChangesConfigurationClass.New()
}

func (f_ FetchRecordZoneChangesConfiguration) Init() FetchRecordZoneChangesConfiguration {
	rv := objc.Call[FetchRecordZoneChangesConfiguration](f_, objc.Sel("init"))
	return rv
}

// The server change token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesconfiguration/2980662-previousserverchangetoken?language=objc
func (f_ FetchRecordZoneChangesConfiguration) PreviousServerChangeToken() ServerChangeToken {
	rv := objc.Call[ServerChangeToken](f_, objc.Sel("previousServerChangeToken"))
	return rv
}

// The server change token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesconfiguration/2980662-previousserverchangetoken?language=objc
func (f_ FetchRecordZoneChangesConfiguration) SetPreviousServerChangeToken(value IServerChangeToken) {
	objc.Call[objc.Void](f_, objc.Sel("setPreviousServerChangeToken:"), objc.Ptr(value))
}

// The maximum number of records that CloudKit retrieves when fetching zone changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesconfiguration/2980663-resultslimit?language=objc
func (f_ FetchRecordZoneChangesConfiguration) ResultsLimit() uint {
	rv := objc.Call[uint](f_, objc.Sel("resultsLimit"))
	return rv
}

// The maximum number of records that CloudKit retrieves when fetching zone changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesconfiguration/2980663-resultslimit?language=objc
func (f_ FetchRecordZoneChangesConfiguration) SetResultsLimit(value uint) {
	objc.Call[objc.Void](f_, objc.Sel("setResultsLimit:"), value)
}

// An array of the record keys to retrieve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesconfiguration/2980661-desiredkeys?language=objc
func (f_ FetchRecordZoneChangesConfiguration) DesiredKeys() []RecordFieldKey {
	rv := objc.Call[[]RecordFieldKey](f_, objc.Sel("desiredKeys"))
	return rv
}

// An array of the record keys to retrieve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesconfiguration/2980661-desiredkeys?language=objc
func (f_ FetchRecordZoneChangesConfiguration) SetDesiredKeys(value []RecordFieldKey) {
	objc.Call[objc.Void](f_, objc.Sel("setDesiredKeys:"), value)
}
