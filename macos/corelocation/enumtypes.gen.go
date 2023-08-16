// AUTO-GENERATED CODE, DO NOT MODIFY

package corelocation

import "math"

// Constants indicating the level of location accuracy the app has authorization to use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/claccuracyauthorization?language=objc
type AccuracyAuthorization int

const (
	AccuracyAuthorizationFullAccuracy    AccuracyAuthorization = 0
	AccuracyAuthorizationReducedAccuracy AccuracyAuthorization = 1
)

// Constants that indicate the type of activity associated with location updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clactivitytype?language=objc
type ActivityType int

const (
	ActivityTypeAirborne             ActivityType = 5
	ActivityTypeAutomotiveNavigation ActivityType = 2
	ActivityTypeFitness              ActivityType = 3
	ActivityTypeOther                ActivityType = 1
	ActivityTypeOtherNavigation      ActivityType = 4
)

// Constants indicating the app's authorization to use location services. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clauthorizationstatus?language=objc
type AuthorizationStatus int

const (
	kCLAuthorizationStatusAuthorized       AuthorizationStatus = 3
	kCLAuthorizationStatusAuthorizedAlways AuthorizationStatus = 3
	kCLAuthorizationStatusDenied           AuthorizationStatus = 2
	kCLAuthorizationStatusNotDetermined    AuthorizationStatus = 0
	kCLAuthorizationStatusRestricted       AuthorizationStatus = 1
)

// The most significant value in a beacon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeaconmajorvalue?language=objc
type BeaconMajorValue uint16

// The least significant value in a beacon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeaconminorvalue?language=objc
type BeaconMinorValue uint16

// Constants indicating the physical orientation of the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cldeviceorientation?language=objc
type DeviceOrientation int

const (
	DeviceOrientationFaceDown           DeviceOrientation = 6
	DeviceOrientationFaceUp             DeviceOrientation = 5
	DeviceOrientationLandscapeLeft      DeviceOrientation = 3
	DeviceOrientationLandscapeRight     DeviceOrientation = 4
	DeviceOrientationPortrait           DeviceOrientation = 1
	DeviceOrientationPortraitUpsideDown DeviceOrientation = 2
	DeviceOrientationUnknown            DeviceOrientation = 0
)

// Error codes returned by the location manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clerror?language=objc
type Error int

const (
	kCLErrorDeferredAccuracyTooLow          Error = 13
	kCLErrorDeferredCanceled                Error = 15
	kCLErrorDeferredDistanceFiltered        Error = 14
	kCLErrorDeferredFailed                  Error = 11
	kCLErrorDeferredNotUpdatingLocation     Error = 12
	kCLErrorDenied                          Error = 1
	kCLErrorGeocodeCanceled                 Error = 10
	kCLErrorGeocodeFoundNoResult            Error = 8
	kCLErrorGeocodeFoundPartialResult       Error = 9
	kCLErrorHeadingFailure                  Error = 3
	kCLErrorLocationUnknown                 Error = 0
	kCLErrorNetwork                         Error = 2
	kCLErrorPromptDeclined                  Error = 18
	kCLErrorRangingFailure                  Error = 17
	kCLErrorRangingUnavailable              Error = 16
	kCLErrorRegionMonitoringDenied          Error = 4
	kCLErrorRegionMonitoringFailure         Error = 5
	kCLErrorRegionMonitoringResponseDelayed Error = 7
	kCLErrorRegionMonitoringSetupDelayed    Error = 6
)

// A type used to report magnetic differences reported by the onboard hardware. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clheadingcomponentvalue?language=objc
type HeadingComponentValue float64

// The accuracy of a geographical coordinate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationaccuracy?language=objc
type LocationAccuracy float64

const (
	kCLLocationAccuracyBest              LocationAccuracy = -1.000000
	kCLLocationAccuracyBestForNavigation LocationAccuracy = -2.000000
	kCLLocationAccuracyHundredMeters     LocationAccuracy = 100.000000
	kCLLocationAccuracyKilometer         LocationAccuracy = 1000.000000
	kCLLocationAccuracyNearestTenMeters  LocationAccuracy = 10.000000
	kCLLocationAccuracyReduced           LocationAccuracy = 6380000.000000
	kCLLocationAccuracyThreeKilometers   LocationAccuracy = 3000.000000
)

// A latitude or longitude value specified in degrees. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationdegrees?language=objc
type LocationDegrees float64

const (
	kCLHeadingFilterNone LocationDegrees = -1.000000
)

// An azimuth that is measured in degrees relative to true north. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationdirection?language=objc
type LocationDirection float64

// The accuracy of a compass heading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationdirectionaccuracy?language=objc
type LocationDirectionAccuracy float64

// A distance in meters from an existing location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationdistance?language=objc
type LocationDistance float64

const (
	kCLDistanceFilterNone LocationDistance = -1.000000
)

var LocationDistanceMax = LocationDistance(math.Inf(1))

// The velocity (measured in meters per second) at which the device is moving. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationspeed?language=objc
type LocationSpeed float64

// The accuracy of a speed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationspeedaccuracy?language=objc
type LocationSpeedAccuracy float64

// Constants that reflect the relative distance to a beacon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clproximity?language=objc
type Proximity int

const (
	ProximityFar       Proximity = 3
	ProximityImmediate Proximity = 1
	ProximityNear      Proximity = 2
	ProximityUnknown   Proximity = 0
)

// Constants that reflect the relationship of the current location to the region boundaries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clregionstate?language=objc
type RegionState int

const (
	RegionStateInside  RegionState = 1
	RegionStateOutside RegionState = 2
	RegionStateUnknown RegionState = 0
)
