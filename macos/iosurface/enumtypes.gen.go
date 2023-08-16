// AUTO-GENERATED CODE, DO NOT MODIFY

package iosurface

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurfacecomponentname?language=objc
type ComponentName int32

const (
	KComponentNameAlpha      ComponentName = 1
	KComponentNameBlue       ComponentName = 4
	KComponentNameChromaBlue ComponentName = 7
	KComponentNameChromaRed  ComponentName = 6
	KComponentNameGreen      ComponentName = 3
	KComponentNameLuma       ComponentName = 5
	KComponentNameRed        ComponentName = 2
	KComponentNameUnknown    ComponentName = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurfacecomponentrange?language=objc
type ComponentRange int32

const (
	KComponentRangeFullRange  ComponentRange = 1
	KComponentRangeUnknown    ComponentRange = 0
	KComponentRangeVideoRange ComponentRange = 2
	KComponentRangeWideRange  ComponentRange = 3
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurfacecomponenttype?language=objc
type ComponentType int32

const (
	KComponentTypeFloat            ComponentType = 3
	KComponentTypeSignedInteger    ComponentType = 2
	KComponentTypeSignedNormalized ComponentType = 4
	KComponentTypeUnknown          ComponentType = 0
	KComponentTypeUnsignedInteger  ComponentType = 1
)

// An IOSurface identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurfaceid?language=objc
type ID uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurfacelockoptions?language=objc
type LockOptions uint32

const (
	KLockAvoidSync LockOptions = 2
	KLockReadOnly  LockOptions = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurfacepropertykey?language=objc
type PropertyKey string

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurfacepurgeabilitystate?language=objc
type PurgeabilityState uint32

const (
	KPurgeableEmpty       PurgeabilityState = 2
	KPurgeableKeepCurrent PurgeabilityState = 3
	KPurgeableNonVolatile PurgeabilityState = 0
	KPurgeableVolatile    PurgeabilityState = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurfacesubsampling?language=objc
type Subsampling int32

const (
	KSubsampling411     Subsampling = 4
	KSubsampling420     Subsampling = 3
	KSubsampling422     Subsampling = 2
	KSubsamplingNone    Subsampling = 1
	KSubsamplingUnknown Subsampling = 0
)
