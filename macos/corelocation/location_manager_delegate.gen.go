// AUTO-GENERATED CODE, DO NOT MODIFY

package corelocation

import (
	"github.com/progrium/macdriver/objc"
)

// The methods that you use to receive events from an associated location-manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate?language=objc
type PLocationManagerDelegate interface {
	// optional
	LocationManagerDidResumeLocationUpdates(manager LocationManager)
	HasLocationManagerDidResumeLocationUpdates() bool

	// optional
	LocationManagerShouldDisplayHeadingCalibration(manager LocationManager) bool
	HasLocationManagerShouldDisplayHeadingCalibration() bool

	// optional
	LocationManagerDidChangeAuthorization(manager LocationManager)
	HasLocationManagerDidChangeAuthorization() bool

	// optional
	LocationManagerDidPauseLocationUpdates(manager LocationManager)
	HasLocationManagerDidPauseLocationUpdates() bool

	// optional
	LocationManagerDidStartMonitoringForRegion(manager LocationManager, region Region)
	HasLocationManagerDidStartMonitoringForRegion() bool
}

// A delegate implementation builder for the [PLocationManagerDelegate] protocol.
type LocationManagerDelegate struct {
	_LocationManagerDidResumeLocationUpdates        func(manager LocationManager)
	_LocationManagerShouldDisplayHeadingCalibration func(manager LocationManager) bool
	_LocationManagerDidChangeAuthorization          func(manager LocationManager)
	_LocationManagerDidPauseLocationUpdates         func(manager LocationManager)
	_LocationManagerDidStartMonitoringForRegion     func(manager LocationManager, region Region)
}

func (di *LocationManagerDelegate) HasLocationManagerDidResumeLocationUpdates() bool {
	return di._LocationManagerDidResumeLocationUpdates != nil
}

// Tells the delegate that the delivery of location updates has resumed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate/1621512-locationmanagerdidresumelocation?language=objc
func (di *LocationManagerDelegate) SetLocationManagerDidResumeLocationUpdates(f func(manager LocationManager)) {
	di._LocationManagerDidResumeLocationUpdates = f
}

// Tells the delegate that the delivery of location updates has resumed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate/1621512-locationmanagerdidresumelocation?language=objc
func (di *LocationManagerDelegate) LocationManagerDidResumeLocationUpdates(manager LocationManager) {
	di._LocationManagerDidResumeLocationUpdates(manager)
}
func (di *LocationManagerDelegate) HasLocationManagerShouldDisplayHeadingCalibration() bool {
	return di._LocationManagerShouldDisplayHeadingCalibration != nil
}

// Asks the delegate whether the heading calibration alert should be displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate/1621457-locationmanagershoulddisplayhead?language=objc
func (di *LocationManagerDelegate) SetLocationManagerShouldDisplayHeadingCalibration(f func(manager LocationManager) bool) {
	di._LocationManagerShouldDisplayHeadingCalibration = f
}

// Asks the delegate whether the heading calibration alert should be displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate/1621457-locationmanagershoulddisplayhead?language=objc
func (di *LocationManagerDelegate) LocationManagerShouldDisplayHeadingCalibration(manager LocationManager) bool {
	return di._LocationManagerShouldDisplayHeadingCalibration(manager)
}
func (di *LocationManagerDelegate) HasLocationManagerDidChangeAuthorization() bool {
	return di._LocationManagerDidChangeAuthorization != nil
}

// Tells the delegate when the app creates the location manager and when the authorization status changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate/3563956-locationmanagerdidchangeauthoriz?language=objc
func (di *LocationManagerDelegate) SetLocationManagerDidChangeAuthorization(f func(manager LocationManager)) {
	di._LocationManagerDidChangeAuthorization = f
}

// Tells the delegate when the app creates the location manager and when the authorization status changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate/3563956-locationmanagerdidchangeauthoriz?language=objc
func (di *LocationManagerDelegate) LocationManagerDidChangeAuthorization(manager LocationManager) {
	di._LocationManagerDidChangeAuthorization(manager)
}
func (di *LocationManagerDelegate) HasLocationManagerDidPauseLocationUpdates() bool {
	return di._LocationManagerDidPauseLocationUpdates != nil
}

// Tells the delegate that location updates were paused. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate/1621553-locationmanagerdidpauselocationu?language=objc
func (di *LocationManagerDelegate) SetLocationManagerDidPauseLocationUpdates(f func(manager LocationManager)) {
	di._LocationManagerDidPauseLocationUpdates = f
}

// Tells the delegate that location updates were paused. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate/1621553-locationmanagerdidpauselocationu?language=objc
func (di *LocationManagerDelegate) LocationManagerDidPauseLocationUpdates(manager LocationManager) {
	di._LocationManagerDidPauseLocationUpdates(manager)
}
func (di *LocationManagerDelegate) HasLocationManagerDidStartMonitoringForRegion() bool {
	return di._LocationManagerDidStartMonitoringForRegion != nil
}

// Tells the delegate that a new region is being monitored. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate/1423842-locationmanager?language=objc
func (di *LocationManagerDelegate) SetLocationManagerDidStartMonitoringForRegion(f func(manager LocationManager, region Region)) {
	di._LocationManagerDidStartMonitoringForRegion = f
}

// Tells the delegate that a new region is being monitored. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate/1423842-locationmanager?language=objc
func (di *LocationManagerDelegate) LocationManagerDidStartMonitoringForRegion(manager LocationManager, region Region) {
	di._LocationManagerDidStartMonitoringForRegion(manager, region)
}

// A concrete type wrapper for the [PLocationManagerDelegate] protocol.
type LocationManagerDelegateWrapper struct {
	objc.Object
}

func (l_ LocationManagerDelegateWrapper) HasLocationManagerDidResumeLocationUpdates() bool {
	return l_.RespondsToSelector(objc.Sel("locationManagerDidResumeLocationUpdates:"))
}

// Tells the delegate that the delivery of location updates has resumed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate/1621512-locationmanagerdidresumelocation?language=objc
func (l_ LocationManagerDelegateWrapper) LocationManagerDidResumeLocationUpdates(manager ILocationManager) {
	objc.Call[objc.Void](l_, objc.Sel("locationManagerDidResumeLocationUpdates:"), objc.Ptr(manager))
}

func (l_ LocationManagerDelegateWrapper) HasLocationManagerShouldDisplayHeadingCalibration() bool {
	return l_.RespondsToSelector(objc.Sel("locationManagerShouldDisplayHeadingCalibration:"))
}

// Asks the delegate whether the heading calibration alert should be displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate/1621457-locationmanagershoulddisplayhead?language=objc
func (l_ LocationManagerDelegateWrapper) LocationManagerShouldDisplayHeadingCalibration(manager ILocationManager) bool {
	rv := objc.Call[bool](l_, objc.Sel("locationManagerShouldDisplayHeadingCalibration:"), objc.Ptr(manager))
	return rv
}

func (l_ LocationManagerDelegateWrapper) HasLocationManagerDidChangeAuthorization() bool {
	return l_.RespondsToSelector(objc.Sel("locationManagerDidChangeAuthorization:"))
}

// Tells the delegate when the app creates the location manager and when the authorization status changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate/3563956-locationmanagerdidchangeauthoriz?language=objc
func (l_ LocationManagerDelegateWrapper) LocationManagerDidChangeAuthorization(manager ILocationManager) {
	objc.Call[objc.Void](l_, objc.Sel("locationManagerDidChangeAuthorization:"), objc.Ptr(manager))
}

func (l_ LocationManagerDelegateWrapper) HasLocationManagerDidPauseLocationUpdates() bool {
	return l_.RespondsToSelector(objc.Sel("locationManagerDidPauseLocationUpdates:"))
}

// Tells the delegate that location updates were paused. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate/1621553-locationmanagerdidpauselocationu?language=objc
func (l_ LocationManagerDelegateWrapper) LocationManagerDidPauseLocationUpdates(manager ILocationManager) {
	objc.Call[objc.Void](l_, objc.Sel("locationManagerDidPauseLocationUpdates:"), objc.Ptr(manager))
}

func (l_ LocationManagerDelegateWrapper) HasLocationManagerDidStartMonitoringForRegion() bool {
	return l_.RespondsToSelector(objc.Sel("locationManager:didStartMonitoringForRegion:"))
}

// Tells the delegate that a new region is being monitored. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate/1423842-locationmanager?language=objc
func (l_ LocationManagerDelegateWrapper) LocationManagerDidStartMonitoringForRegion(manager ILocationManager, region IRegion) {
	objc.Call[objc.Void](l_, objc.Sel("locationManager:didStartMonitoringForRegion:"), objc.Ptr(manager), objc.Ptr(region))
}
