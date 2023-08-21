// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Timer] class.
var TimerClass = _TimerClass{objc.GetClass("NSTimer")}

type _TimerClass struct {
	objc.Class
}

// An interface definition for the [Timer] class.
type ITimer interface {
	objc.IObject
	Fire()
	Invalidate()
	TimeInterval() TimeInterval
	FireDate() Date
	SetFireDate(value IDate)
	Tolerance() TimeInterval
	SetTolerance(value TimeInterval)
	UserInfo() objc.Object
	IsValid() bool
}

// A timer that fires after a certain time interval has elapsed, sending a specified message to a target object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer?language=objc
type Timer struct {
	objc.Object
}

func TimerFrom(ptr unsafe.Pointer) Timer {
	return Timer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ Timer) InitWithFireDateIntervalRepeatsBlock(date IDate, interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	rv := objc.Call[Timer](t_, objc.Sel("initWithFireDate:interval:repeats:block:"), objc.Ptr(date), interval, repeats, block)
	return rv
}

// Initializes a timer for the specified date and time interval with the specified block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/2091887-initwithfiredate?language=objc
func NewTimerWithFireDateIntervalRepeatsBlock(date IDate, interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	instance := TimerClass.Alloc().InitWithFireDateIntervalRepeatsBlock(date, interval, repeats, block)
	instance.Autorelease()
	return instance
}

func (tc _TimerClass) Alloc() Timer {
	rv := objc.Call[Timer](tc, objc.Sel("alloc"))
	return rv
}

func Timer_Alloc() Timer {
	return TimerClass.Alloc()
}

func (tc _TimerClass) New() Timer {
	rv := objc.Call[Timer](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTimer() Timer {
	return TimerClass.New()
}

func (t_ Timer) Init() Timer {
	rv := objc.Call[Timer](t_, objc.Sel("init"))
	return rv
}

// Creates a timer and schedules it on the current run loop in the default mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/2091889-scheduledtimerwithtimeinterval?language=objc
func (tc _TimerClass) ScheduledTimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	rv := objc.Call[Timer](tc, objc.Sel("scheduledTimerWithTimeInterval:repeats:block:"), interval, repeats, block)
	return rv
}

// Creates a timer and schedules it on the current run loop in the default mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/2091889-scheduledtimerwithtimeinterval?language=objc
func Timer_ScheduledTimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	return TimerClass.ScheduledTimerWithTimeIntervalRepeatsBlock(interval, repeats, block)
}

// Causes the timer's message to be sent to its target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1414035-fire?language=objc
func (t_ Timer) Fire() {
	objc.Call[objc.Void](t_, objc.Sel("fire"))
}

// Initializes a timer object with the specified time interval and block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/2091888-timerwithtimeinterval?language=objc
func (tc _TimerClass) TimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	rv := objc.Call[Timer](tc, objc.Sel("timerWithTimeInterval:repeats:block:"), interval, repeats, block)
	return rv
}

// Initializes a timer object with the specified time interval and block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/2091888-timerwithtimeinterval?language=objc
func Timer_TimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	return TimerClass.TimerWithTimeIntervalRepeatsBlock(interval, repeats, block)
}

// Stops the timer from ever firing again and requests its removal from its run loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1415405-invalidate?language=objc
func (t_ Timer) Invalidate() {
	objc.Call[objc.Void](t_, objc.Sel("invalidate"))
}

// The timer’s time interval, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1409024-timeinterval?language=objc
func (t_ Timer) TimeInterval() TimeInterval {
	rv := objc.Call[TimeInterval](t_, objc.Sel("timeInterval"))
	return rv
}

// The date at which the timer will fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1407353-firedate?language=objc
func (t_ Timer) FireDate() Date {
	rv := objc.Call[Date](t_, objc.Sel("fireDate"))
	return rv
}

// The date at which the timer will fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1407353-firedate?language=objc
func (t_ Timer) SetFireDate(value IDate) {
	objc.Call[objc.Void](t_, objc.Sel("setFireDate:"), objc.Ptr(value))
}

// The amount of time after the scheduled fire date that the timer may fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1415085-tolerance?language=objc
func (t_ Timer) Tolerance() TimeInterval {
	rv := objc.Call[TimeInterval](t_, objc.Sel("tolerance"))
	return rv
}

// The amount of time after the scheduled fire date that the timer may fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1415085-tolerance?language=objc
func (t_ Timer) SetTolerance(value TimeInterval) {
	objc.Call[objc.Void](t_, objc.Sel("setTolerance:"), value)
}

// The receiver's userInfo object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1408911-userinfo?language=objc
func (t_ Timer) UserInfo() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("userInfo"))
	return rv
}

// A Boolean value that indicates whether the timer is currently valid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimer/1408249-valid?language=objc
func (t_ Timer) IsValid() bool {
	rv := objc.Call[bool](t_, objc.Sel("isValid"))
	return rv
}
