// AUTO-GENERATED CODE, DO NOT MODIFY

package sysconfig

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
)

// Type of the callback function used when the preferences have been updated or applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scpreferencescallback?language=objc
type PreferencesCallBack = func(prefs PreferencesRef, notificationType PreferencesNotification, info unsafe.Pointer)

// The type of callback function used when a status event is delivered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scnetworkconnectioncallback?language=objc
type etworkConnectionCallBack = func(connection EtworkConnectionRef, status EtworkConnectionStatus, info unsafe.Pointer)

// Type of callback function used when the reachability of a network address or name changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scnetworkreachabilitycallback?language=objc
type etworkReachabilityCallBack = func(target EtworkReachabilityRef, flags EtworkReachabilityFlags, info unsafe.Pointer)

// Callback used when notification of changes made to the dynamic store is delivered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/scdynamicstorecallback?language=objc
type DynamicStoreCallBack = func(store DynamicStoreRef, changedKeys corefoundation.ArrayRef, info unsafe.Pointer)
