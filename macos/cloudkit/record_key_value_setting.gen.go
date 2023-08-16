// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol for managing the key-value pairs of a CloudKit record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordkeyvaluesetting?language=objc
type PRecordKeyValueSetting interface {
	// optional
	ObjectForKey(key RecordFieldKey) PRecordValue
	HasObjectForKey() bool

	// optional
	ObjectForKeyedSubscript(key RecordFieldKey) PRecordValue
	HasObjectForKeyedSubscript() bool

	// optional
	SetObjectForKeyedSubscript(object RecordValueWrapper, key RecordFieldKey)
	HasSetObjectForKeyedSubscript() bool

	// optional
	AllKeys() []RecordFieldKey
	HasAllKeys() bool

	// optional
	ChangedKeys() []RecordFieldKey
	HasChangedKeys() bool
}

// A concrete type wrapper for the [PRecordKeyValueSetting] protocol.
type RecordKeyValueSettingWrapper struct {
	objc.Object
}

func (r_ RecordKeyValueSettingWrapper) HasObjectForKey() bool {
	return r_.RespondsToSelector(objc.Sel("objectForKey:"))
}

// Returns the object that the record stores for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordkeyvaluesetting/2976185-objectforkey?language=objc
func (r_ RecordKeyValueSettingWrapper) ObjectForKey(key RecordFieldKey) RecordValueWrapper {
	rv := objc.Call[RecordValueWrapper](r_, objc.Sel("objectForKey:"), key)
	return rv
}

func (r_ RecordKeyValueSettingWrapper) HasObjectForKeyedSubscript() bool {
	return r_.RespondsToSelector(objc.Sel("objectForKeyedSubscript:"))
}

// Returns the object that the record stores for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordkeyvaluesetting/2976186-objectforkeyedsubscript?language=objc
func (r_ RecordKeyValueSettingWrapper) ObjectForKeyedSubscript(key RecordFieldKey) RecordValueWrapper {
	rv := objc.Call[RecordValueWrapper](r_, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}

func (r_ RecordKeyValueSettingWrapper) HasSetObjectForKeyedSubscript() bool {
	return r_.RespondsToSelector(objc.Sel("setObject:forKeyedSubscript:"))
}

// Stores an object in the record using the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordkeyvaluesetting/2976188-setobject?language=objc
func (r_ RecordKeyValueSettingWrapper) SetObjectForKeyedSubscript(object PRecordValue, key RecordFieldKey) {
	po0 := objc.WrapAsProtocol("CKRecordValue", object)
	objc.Call[objc.Void](r_, objc.Sel("setObject:forKeyedSubscript:"), po0, key)
}

func (r_ RecordKeyValueSettingWrapper) HasAllKeys() bool {
	return r_.RespondsToSelector(objc.Sel("allKeys"))
}

// Returns an array of the record’s keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordkeyvaluesetting/2976183-allkeys?language=objc
func (r_ RecordKeyValueSettingWrapper) AllKeys() []RecordFieldKey {
	rv := objc.Call[[]RecordFieldKey](r_, objc.Sel("allKeys"))
	return rv
}

func (r_ RecordKeyValueSettingWrapper) HasChangedKeys() bool {
	return r_.RespondsToSelector(objc.Sel("changedKeys"))
}

// Returns an array of keys with recent changes to their values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordkeyvaluesetting/2976184-changedkeys?language=objc
func (r_ RecordKeyValueSettingWrapper) ChangedKeys() []RecordFieldKey {
	rv := objc.Call[[]RecordFieldKey](r_, objc.Sel("changedKeys"))
	return rv
}
