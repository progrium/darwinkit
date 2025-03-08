package healthkit

import (
	"github.com/progrium/darwinkit/objc"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/corelocation"
)

// HealthStore provides an interface for accessing and using health data
type HealthStore struct {
	objc.Object
}

// HealthStoreClass is the class instance for HealthStore
var HealthStoreClass = objc.GetClass("HKHealthStore")

// NewHealthStore creates a new health store instance
func NewHealthStore() HealthStore {
	return HealthStore{objc.Call[objc.Object](HealthStoreClass, objc.Sel("alloc")).Send(objc.Sel("init"))}
}

// IsHealthDataAvailable returns whether HealthKit is available on this device
func IsHealthDataAvailable() bool {
	return objc.Call[bool](HealthStoreClass, objc.Sel("isHealthDataAvailable"))
}

// RequestAuthorizationToShareTypes requests permission to share health data
func (h HealthStore) RequestAuthorizationToShareTypes(typesToShare foundation.Set, typesToRead foundation.Set, completion foundation.CompletionHandler) {
	h.Send(objc.Sel("requestAuthorizationToShareTypes:typesToRead:completion:"), typesToShare, typesToRead, completion)
}

// QuantityType represents a type of numerical sample
type QuantityType struct {
	objc.Object
}

// QuantityTypeClass is the class instance for QuantityType
var QuantityTypeClass = objc.GetClass("HKQuantityType")

// QuantityTypeIdentifierHeartRate returns the identifier for heart rate data
func QuantityTypeIdentifierHeartRate() foundation.String {
	return objc.Call[foundation.String](QuantityTypeClass, objc.Sel("quantityTypeIdentifierHeartRate"))
}

// QuantityTypeIdentifierStepCount returns the identifier for step count data
func QuantityTypeIdentifierStepCount() foundation.String {
	return objc.Call[foundation.String](QuantityTypeClass, objc.Sel("quantityTypeIdentifierStepCount"))
}

// QuantityTypeIdentifierActiveEnergyBurned returns the identifier for active energy burned data
func QuantityTypeIdentifierActiveEnergyBurned() foundation.String {
	return objc.Call[foundation.String](QuantityTypeClass, objc.Sel("quantityTypeIdentifierActiveEnergyBurned"))
}

// QuantitySample represents a single measurement of a quantity type
type QuantitySample struct {
	objc.Object
}

// QuantitySampleClass is the class instance for QuantitySample
var QuantitySampleClass = objc.GetClass("HKQuantitySample")

// NewQuantitySampleWithTypeQuantityStartEndMetadata creates a new quantity sample
func NewQuantitySampleWithTypeQuantityStartEndMetadata(quantityType QuantityType, quantity Quantity, startDate foundation.Date, endDate foundation.Date, metadata foundation.Dictionary) QuantitySample {
	return QuantitySample{objc.Call[objc.Object](QuantitySampleClass, objc.Sel("quantitySampleWithType:quantity:startDate:endDate:metadata:"), 
		quantityType, quantity, startDate, endDate, metadata)}
}

// Quantity represents a numerical value with a unit
type Quantity struct {
	objc.Object
}

// QuantityClass is the class instance for Quantity
var QuantityClass = objc.GetClass("HKQuantity")

// NewQuantityWithUnitDoubleValue creates a new quantity with the specified unit and value
func NewQuantityWithUnitDoubleValue(unit Unit, value float64) Quantity {
	return Quantity{objc.Call[objc.Object](QuantityClass, objc.Sel("quantityWithUnit:doubleValue:"), unit, value)}
}

// DoubleValueForUnit returns the quantity's value in the specified unit
func (q Quantity) DoubleValueForUnit(unit Unit) float64 {
	return objc.Call[float64](q, objc.Sel("doubleValueForUnit:"), unit)
}

// Unit represents a unit of measurement
type Unit struct {
	objc.Object
}

// UnitClass is the class instance for Unit
var UnitClass = objc.GetClass("HKUnit")

// Count returns the count unit
func Count() Unit {
	return Unit{objc.Call[objc.Object](UnitClass, objc.Sel("count"))}
}

// MeterUnitWithMetricPrefix returns a meter unit with the specified prefix
func MeterUnitWithMetricPrefix(prefix int) Unit {
	return Unit{objc.Call[objc.Object](UnitClass, objc.Sel("meterUnitWithMetricPrefix:"), prefix)}
}

// NewUnitFromString creates a unit from a string representation
func NewUnitFromString(unitString foundation.String) Unit {
	return Unit{objc.Call[objc.Object](UnitClass, objc.Sel("unitFromString:"), unitString)}
}

// WorkoutType represents a type of workout
type WorkoutType struct {
	objc.Object
}

// WorkoutTypeClass is the class instance for WorkoutType
var WorkoutTypeClass = objc.GetClass("HKWorkoutType")

// Workout represents a workout session
type Workout struct {
	objc.Object
}

// WorkoutClass is the class instance for Workout
var WorkoutClass = objc.GetClass("HKWorkout")

// NewWorkoutWithActivityTypeStartEndEnergyBurnedDistance totalEnergyBurnedMetadata creates a new workout
func NewWorkoutWithActivityTypeStartEndEnergyBurnedDistanceTotalEnergyBurnedMetadata(
	activityType int,
	startDate foundation.Date,
	endDate foundation.Date,
	energyBurned Quantity,
	distance Quantity,
	totalEnergyBurned Quantity,
	metadata foundation.Dictionary,
) Workout {
	return Workout{objc.Call[objc.Object](WorkoutClass, objc.Sel("workoutWithActivityType:startDate:endDate:energyBurned:distance:totalEnergyBurned:metadata:"), 
		activityType, startDate, endDate, energyBurned, distance, totalEnergyBurned, metadata)}
}

// Various activity types for workouts
const (
	WorkoutActivityTypeRunning = 8
	WorkoutActivityTypeWalking = 7
	WorkoutActivityTypeCycling = 13
	WorkoutActivityTypeSwimming = 26
	WorkoutActivityTypeYoga = 35
	WorkoutActivityTypeFunctionalStrengthTraining = 37
	WorkoutActivityTypeTraditionalStrengthTraining = 3
	WorkoutActivityTypeHiit = 63
)

// Various metric prefixes
const (
	MetricPrefixNone  = 0
	MetricPrefixPico  = 1
	MetricPrefixNano  = 2
	MetricPrefixMicro = 3
	MetricPrefixMilli = 4
	MetricPrefixCenti = 5
	MetricPrefixDeci  = 6
	MetricPrefixDeca  = 7
	MetricPrefixHecto = 8
	MetricPrefixKilo  = 9
	MetricPrefixMega  = 10
	MetricPrefixGiga  = 11
	MetricPrefixTera  = 12
)