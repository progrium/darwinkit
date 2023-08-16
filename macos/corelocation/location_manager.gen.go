// AUTO-GENERATED CODE, DO NOT MODIFY

package corelocation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LocationManager] class.
var LocationManagerClass = _LocationManagerClass{objc.GetClass("CLLocationManager")}

type _LocationManagerClass struct {
	objc.Class
}

// An interface definition for the [LocationManager] class.
type ILocationManager interface {
	objc.IObject
	StopUpdatingLocation()
	RequestAlwaysAuthorization()
	StartMonitoringVisits()
	StopMonitoringSignificantLocationChanges()
	RequestLocation()
	StartMonitoringSignificantLocationChanges()
	DismissHeadingCalibrationDisplay()
	StartUpdatingHeading()
	RequestWhenInUseAuthorization()
	StopMonitoringVisits()
	StartUpdatingLocation()
	RequestTemporaryFullAccuracyAuthorizationWithPurposeKey(purposeKey string)
	Location() Location
	ActivityType() ActivityType
	SetActivityType(value ActivityType)
	IsAuthorizedForWidgetUpdates() bool
	AuthorizationStatus() AuthorizationStatus
	Delegate() LocationManagerDelegateWrapper
	SetDelegate(value PLocationManagerDelegate)
	SetDelegateObject(valueObject objc.IObject)
	DistanceFilter() LocationDistance
	SetDistanceFilter(value LocationDistance)
	PausesLocationUpdatesAutomatically() bool
	SetPausesLocationUpdatesAutomatically(value bool)
	AccuracyAuthorization() AccuracyAuthorization
	DesiredAccuracy() LocationAccuracy
	SetDesiredAccuracy(value LocationAccuracy)
	HeadingFilter() LocationDegrees
	SetHeadingFilter(value LocationDegrees)
	MaximumRegionMonitoringDistance() LocationDistance
	Heading() Heading
	AllowsBackgroundLocationUpdates() bool
	SetAllowsBackgroundLocationUpdates(value bool)
	MonitoredRegions() foundation.Set
	HeadingOrientation() DeviceOrientation
	SetHeadingOrientation(value DeviceOrientation)
}

// The object that you use to start and stop the delivery of location-related events to your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager?language=objc
type LocationManager struct {
	objc.Object
}

func LocationManagerFrom(ptr unsafe.Pointer) LocationManager {
	return LocationManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (lc _LocationManagerClass) Alloc() LocationManager {
	rv := objc.Call[LocationManager](lc, objc.Sel("alloc"))
	return rv
}

func LocationManager_Alloc() LocationManager {
	return LocationManagerClass.Alloc()
}

func (lc _LocationManagerClass) New() LocationManager {
	rv := objc.Call[LocationManager](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLocationManager() LocationManager {
	return LocationManagerClass.New()
}

func (l_ LocationManager) Init() LocationManager {
	rv := objc.Call[LocationManager](l_, objc.Sel("init"))
	return rv
}

// Returns a Boolean value indicating whether the significant-change location service is available on the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423677-significantlocationchangemonitor?language=objc
func (lc _LocationManagerClass) SignificantLocationChangeMonitoringAvailable() bool {
	rv := objc.Call[bool](lc, objc.Sel("significantLocationChangeMonitoringAvailable"))
	return rv
}

// Returns a Boolean value indicating whether the significant-change location service is available on the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423677-significantlocationchangemonitor?language=objc
func LocationManager_SignificantLocationChangeMonitoringAvailable() bool {
	return LocationManagerClass.SignificantLocationChangeMonitoringAvailable()
}

// Stops the generation of location updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423695-stopupdatinglocation?language=objc
func (l_ LocationManager) StopUpdatingLocation() {
	objc.Call[objc.Void](l_, objc.Sel("stopUpdatingLocation"))
}

// Requests the user’s permission to use location services regardless of whether the app is in use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620551-requestalwaysauthorization?language=objc
func (l_ LocationManager) RequestAlwaysAuthorization() {
	objc.Call[objc.Void](l_, objc.Sel("requestAlwaysAuthorization"))
}

// Starts the delivery of visit-related events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1618692-startmonitoringvisits?language=objc
func (l_ LocationManager) StartMonitoringVisits() {
	objc.Call[objc.Void](l_, objc.Sel("startMonitoringVisits"))
}

// Stops the delivery of location events based on significant location changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423679-stopmonitoringsignificantlocatio?language=objc
func (l_ LocationManager) StopMonitoringSignificantLocationChanges() {
	objc.Call[objc.Void](l_, objc.Sel("stopMonitoringSignificantLocationChanges"))
}

// Requests the one-time delivery of the user’s current location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620548-requestlocation?language=objc
func (l_ LocationManager) RequestLocation() {
	objc.Call[objc.Void](l_, objc.Sel("requestLocation"))
}

// Starts the generation of updates based on significant location changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423531-startmonitoringsignificantlocati?language=objc
func (l_ LocationManager) StartMonitoringSignificantLocationChanges() {
	objc.Call[objc.Void](l_, objc.Sel("startMonitoringSignificantLocationChanges"))
}

// Returns a Boolean value indicating whether the device supports region monitoring using the specified class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423654-ismonitoringavailableforclass?language=objc
func (lc _LocationManagerClass) IsMonitoringAvailableForClass(regionClass objc.IClass) bool {
	rv := objc.Call[bool](lc, objc.Sel("isMonitoringAvailableForClass:"), objc.Ptr(regionClass))
	return rv
}

// Returns a Boolean value indicating whether the device supports region monitoring using the specified class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423654-ismonitoringavailableforclass?language=objc
func LocationManager_IsMonitoringAvailableForClass(regionClass objc.IClass) bool {
	return LocationManagerClass.IsMonitoringAvailableForClass(regionClass)
}

// Dismisses the heading calibration view from the screen immediately. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620563-dismissheadingcalibrationdisplay?language=objc
func (l_ LocationManager) DismissHeadingCalibrationDisplay() {
	objc.Call[objc.Void](l_, objc.Sel("dismissHeadingCalibrationDisplay"))
}

// Returns a Boolean value indicating whether the device supports ranging of beacons that use the iBeacon protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620549-israngingavailable?language=objc
func (lc _LocationManagerClass) IsRangingAvailable() bool {
	rv := objc.Call[bool](lc, objc.Sel("isRangingAvailable"))
	return rv
}

// Returns a Boolean value indicating whether the device supports ranging of beacons that use the iBeacon protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620549-israngingavailable?language=objc
func LocationManager_IsRangingAvailable() bool {
	return LocationManagerClass.IsRangingAvailable()
}

// Starts the generation of updates that report the user’s current heading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620558-startupdatingheading?language=objc
func (l_ LocationManager) StartUpdatingHeading() {
	objc.Call[objc.Void](l_, objc.Sel("startUpdatingHeading"))
}

// Requests the user’s permission to use location services while the app is in use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620562-requestwheninuseauthorization?language=objc
func (l_ LocationManager) RequestWhenInUseAuthorization() {
	objc.Call[objc.Void](l_, objc.Sel("requestWhenInUseAuthorization"))
}

// Stops the delivery of visit-related events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1618693-stopmonitoringvisits?language=objc
func (l_ LocationManager) StopMonitoringVisits() {
	objc.Call[objc.Void](l_, objc.Sel("stopMonitoringVisits"))
}

// Starts the generation of updates that report the user’s current location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423750-startupdatinglocation?language=objc
func (l_ LocationManager) StartUpdatingLocation() {
	objc.Call[objc.Void](l_, objc.Sel("startUpdatingLocation"))
}

// Requests permission to temporarily use location services with full accuracy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/3600216-requesttemporaryfullaccuracyauth?language=objc
func (l_ LocationManager) RequestTemporaryFullAccuracyAuthorizationWithPurposeKey(purposeKey string) {
	objc.Call[objc.Void](l_, objc.Sel("requestTemporaryFullAccuracyAuthorizationWithPurposeKey:"), purposeKey)
}

// The most recently retrieved user location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423687-location?language=objc
func (l_ LocationManager) Location() Location {
	rv := objc.Call[Location](l_, objc.Sel("location"))
	return rv
}

// The type of activity the app expects the user to typically perform while in the app’s location session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620567-activitytype?language=objc
func (l_ LocationManager) ActivityType() ActivityType {
	rv := objc.Call[ActivityType](l_, objc.Sel("activityType"))
	return rv
}

// The type of activity the app expects the user to typically perform while in the app’s location session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620567-activitytype?language=objc
func (l_ LocationManager) SetActivityType(value ActivityType) {
	objc.Call[objc.Void](l_, objc.Sel("setActivityType:"), value)
}

// A Boolean value that indicates whether a widget is eligible to receive location updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/3675588-authorizedforwidgetupdates?language=objc
func (l_ LocationManager) IsAuthorizedForWidgetUpdates() bool {
	rv := objc.Call[bool](l_, objc.Sel("isAuthorizedForWidgetUpdates"))
	return rv
}

// The current authorization status for the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/3563952-authorizationstatus?language=objc
func (l_ LocationManager) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Call[AuthorizationStatus](l_, objc.Sel("authorizationStatus"))
	return rv
}

// The delegate object to receive update events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423792-delegate?language=objc
func (l_ LocationManager) Delegate() LocationManagerDelegateWrapper {
	rv := objc.Call[LocationManagerDelegateWrapper](l_, objc.Sel("delegate"))
	return rv
}

// The delegate object to receive update events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423792-delegate?language=objc
func (l_ LocationManager) SetDelegate(value PLocationManagerDelegate) {
	po0 := objc.WrapAsProtocol("CLLocationManagerDelegate", value)
	objc.SetAssociatedObject(l_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](l_, objc.Sel("setDelegate:"), po0)
}

// The delegate object to receive update events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423792-delegate?language=objc
func (l_ LocationManager) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The minimum distance in meters the device must move horizontally before an update event is generated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423500-distancefilter?language=objc
func (l_ LocationManager) DistanceFilter() LocationDistance {
	rv := objc.Call[LocationDistance](l_, objc.Sel("distanceFilter"))
	return rv
}

// The minimum distance in meters the device must move horizontally before an update event is generated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423500-distancefilter?language=objc
func (l_ LocationManager) SetDistanceFilter(value LocationDistance) {
	objc.Call[objc.Void](l_, objc.Sel("setDistanceFilter:"), value)
}

// A Boolean value that indicates whether the location-manager object may pause location updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620553-pauseslocationupdatesautomatical?language=objc
func (l_ LocationManager) PausesLocationUpdatesAutomatically() bool {
	rv := objc.Call[bool](l_, objc.Sel("pausesLocationUpdatesAutomatically"))
	return rv
}

// A Boolean value that indicates whether the location-manager object may pause location updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620553-pauseslocationupdatesautomatical?language=objc
func (l_ LocationManager) SetPausesLocationUpdatesAutomatically(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setPausesLocationUpdatesAutomatically:"), value)
}

// A value that indicates the level of location accuracy the app has permission to use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/3600215-accuracyauthorization?language=objc
func (l_ LocationManager) AccuracyAuthorization() AccuracyAuthorization {
	rv := objc.Call[AccuracyAuthorization](l_, objc.Sel("accuracyAuthorization"))
	return rv
}

// The accuracy of the location data that your app wants to receive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423836-desiredaccuracy?language=objc
func (l_ LocationManager) DesiredAccuracy() LocationAccuracy {
	rv := objc.Call[LocationAccuracy](l_, objc.Sel("desiredAccuracy"))
	return rv
}

// The accuracy of the location data that your app wants to receive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423836-desiredaccuracy?language=objc
func (l_ LocationManager) SetDesiredAccuracy(value LocationAccuracy) {
	objc.Call[objc.Void](l_, objc.Sel("setDesiredAccuracy:"), value)
}

// The minimum angular change in degrees required to generate new heading events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620550-headingfilter?language=objc
func (l_ LocationManager) HeadingFilter() LocationDegrees {
	rv := objc.Call[LocationDegrees](l_, objc.Sel("headingFilter"))
	return rv
}

// The minimum angular change in degrees required to generate new heading events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620550-headingfilter?language=objc
func (l_ LocationManager) SetHeadingFilter(value LocationDegrees) {
	objc.Call[objc.Void](l_, objc.Sel("setHeadingFilter:"), value)
}

// The largest boundary distance that can be assigned to a region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423740-maximumregionmonitoringdistance?language=objc
func (l_ LocationManager) MaximumRegionMonitoringDistance() LocationDistance {
	rv := objc.Call[LocationDistance](l_, objc.Sel("maximumRegionMonitoringDistance"))
	return rv
}

// The most recently reported heading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620555-heading?language=objc
func (l_ LocationManager) Heading() Heading {
	rv := objc.Call[Heading](l_, objc.Sel("heading"))
	return rv
}

// A Boolean value that indicates whether the app receives location updates when running in the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620568-allowsbackgroundlocationupdates?language=objc
func (l_ LocationManager) AllowsBackgroundLocationUpdates() bool {
	rv := objc.Call[bool](l_, objc.Sel("allowsBackgroundLocationUpdates"))
	return rv
}

// A Boolean value that indicates whether the app receives location updates when running in the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620568-allowsbackgroundlocationupdates?language=objc
func (l_ LocationManager) SetAllowsBackgroundLocationUpdates(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setAllowsBackgroundLocationUpdates:"), value)
}

// The set of shared regions monitored by all location-manager objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1423790-monitoredregions?language=objc
func (l_ LocationManager) MonitoredRegions() foundation.Set {
	rv := objc.Call[foundation.Set](l_, objc.Sel("monitoredRegions"))
	return rv
}

// The device orientation to use when computing heading values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620556-headingorientation?language=objc
func (l_ LocationManager) HeadingOrientation() DeviceOrientation {
	rv := objc.Call[DeviceOrientation](l_, objc.Sel("headingOrientation"))
	return rv
}

// The device orientation to use when computing heading values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/1620556-headingorientation?language=objc
func (l_ LocationManager) SetHeadingOrientation(value DeviceOrientation) {
	objc.Call[objc.Void](l_, objc.Sel("setHeadingOrientation:"), value)
}
