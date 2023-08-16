// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BackgroundActivityScheduler] class.
var BackgroundActivitySchedulerClass = _BackgroundActivitySchedulerClass{objc.GetClass("NSBackgroundActivityScheduler")}

type _BackgroundActivitySchedulerClass struct {
	objc.Class
}

// An interface definition for the [BackgroundActivityScheduler] class.
type IBackgroundActivityScheduler interface {
	objc.IObject
	ScheduleWithBlock(block func(completionHandler BackgroundActivityCompletionHandler))
	Invalidate()
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	Tolerance() TimeInterval
	SetTolerance(value TimeInterval)
	ShouldDefer() bool
	Repeats() bool
	SetRepeats(value bool)
	Interval() TimeInterval
	SetInterval(value TimeInterval)
	Identifier() string
}

// A task scheduler suitable for low priority operations that can run in the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivityscheduler?language=objc
type BackgroundActivityScheduler struct {
	objc.Object
}

func BackgroundActivitySchedulerFrom(ptr unsafe.Pointer) BackgroundActivityScheduler {
	return BackgroundActivityScheduler{
		Object: objc.ObjectFrom(ptr),
	}
}

func (b_ BackgroundActivityScheduler) InitWithIdentifier(identifier string) BackgroundActivityScheduler {
	rv := objc.Call[BackgroundActivityScheduler](b_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Initializes a background activity scheduler object with a specified unique identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivityscheduler/1407482-initwithidentifier?language=objc
func BackgroundActivityScheduler_InitWithIdentifier(identifier string) BackgroundActivityScheduler {
	return BackgroundActivitySchedulerClass.Alloc().InitWithIdentifier(identifier)
}

func (bc _BackgroundActivitySchedulerClass) Alloc() BackgroundActivityScheduler {
	rv := objc.Call[BackgroundActivityScheduler](bc, objc.Sel("alloc"))
	return rv
}

func BackgroundActivityScheduler_Alloc() BackgroundActivityScheduler {
	return BackgroundActivitySchedulerClass.Alloc()
}

func (bc _BackgroundActivitySchedulerClass) New() BackgroundActivityScheduler {
	rv := objc.Call[BackgroundActivityScheduler](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBackgroundActivityScheduler() BackgroundActivityScheduler {
	return BackgroundActivitySchedulerClass.New()
}

func (b_ BackgroundActivityScheduler) Init() BackgroundActivityScheduler {
	rv := objc.Call[BackgroundActivityScheduler](b_, objc.Sel("init"))
	return rv
}

// Begins scheduling the background activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivityscheduler/1412813-schedulewithblock?language=objc
func (b_ BackgroundActivityScheduler) ScheduleWithBlock(block func(completionHandler BackgroundActivityCompletionHandler)) {
	objc.Call[objc.Void](b_, objc.Sel("scheduleWithBlock:"), block)
}

// Prevents the background activity from being scheduled again. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivityscheduler/1408878-invalidate?language=objc
func (b_ BackgroundActivityScheduler) Invalidate() {
	objc.Call[objc.Void](b_, objc.Sel("invalidate"))
}

// A value of type NSQualityOfService, which controls how aggressively the system schedules the activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivityscheduler/1412688-qualityofservice?language=objc
func (b_ BackgroundActivityScheduler) QualityOfService() QualityOfService {
	rv := objc.Call[QualityOfService](b_, objc.Sel("qualityOfService"))
	return rv
}

// A value of type NSQualityOfService, which controls how aggressively the system schedules the activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivityscheduler/1412688-qualityofservice?language=objc
func (b_ BackgroundActivityScheduler) SetQualityOfService(value QualityOfService) {
	objc.Call[objc.Void](b_, objc.Sel("setQualityOfService:"), value)
}

// A value of type NSTimeInterval, which specifies a range of time during which the background activity may occur. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivityscheduler/1408138-tolerance?language=objc
func (b_ BackgroundActivityScheduler) Tolerance() TimeInterval {
	rv := objc.Call[TimeInterval](b_, objc.Sel("tolerance"))
	return rv
}

// A value of type NSTimeInterval, which specifies a range of time during which the background activity may occur. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivityscheduler/1408138-tolerance?language=objc
func (b_ BackgroundActivityScheduler) SetTolerance(value TimeInterval) {
	objc.Call[objc.Void](b_, objc.Sel("setTolerance:"), value)
}

// A Boolean value indicating whether your app should stop performing background activity and resume at a more optimal time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivityscheduler/1412167-shoulddefer?language=objc
func (b_ BackgroundActivityScheduler) ShouldDefer() bool {
	rv := objc.Call[bool](b_, objc.Sel("shouldDefer"))
	return rv
}

// A Boolean value indicating whether the activity should be rescheduled after it completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivityscheduler/1409853-repeats?language=objc
func (b_ BackgroundActivityScheduler) Repeats() bool {
	rv := objc.Call[bool](b_, objc.Sel("repeats"))
	return rv
}

// A Boolean value indicating whether the activity should be rescheduled after it completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivityscheduler/1409853-repeats?language=objc
func (b_ BackgroundActivityScheduler) SetRepeats(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setRepeats:"), value)
}

// An integer providing a suggested interval between scheduling and invoking the activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivityscheduler/1408819-interval?language=objc
func (b_ BackgroundActivityScheduler) Interval() TimeInterval {
	rv := objc.Call[TimeInterval](b_, objc.Sel("interval"))
	return rv
}

// An integer providing a suggested interval between scheduling and invoking the activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivityscheduler/1408819-interval?language=objc
func (b_ BackgroundActivityScheduler) SetInterval(value TimeInterval) {
	objc.Call[objc.Void](b_, objc.Sel("setInterval:"), value)
}

// A unique reverse DNS notation string, such as com.example.MyApp.updatecheck, that identifies the activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivityscheduler/1412285-identifier?language=objc
func (b_ BackgroundActivityScheduler) Identifier() string {
	rv := objc.Call[string](b_, objc.Sel("identifier"))
	return rv
}
