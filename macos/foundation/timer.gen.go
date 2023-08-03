// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var TimerClass = _TimerClass{objc.GetClass("NSTimer")}

type _TimerClass struct {
	objc.Class
}

type ITimer interface {
	objc.IObject
	Fire()
	Invalidate()
	IsValid() bool
	FireDate() Date
	SetFireDate(value IDate)
	TimeInterval() TimeInterval
	UserInfo() objc.Object
	Tolerance() TimeInterval
	SetTolerance(value TimeInterval)
}

type Timer struct {
	objc.Object
}

func MakeTimer(ptr unsafe.Pointer) Timer {
	return Timer{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ Timer) InitWithFireDateIntervalRepeatsBlock(date IDate, interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	rv := objc.CallMethod[Timer](t_, objc.GetSelector("initWithFireDate:interval:repeats:block:"), objc.ExtractPtr(date), interval, repeats, block)
	return rv
}

func Timer_InitWithFireDateIntervalRepeatsBlock(date IDate, interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	return TimerClass.Alloc().InitWithFireDateIntervalRepeatsBlock(date, interval, repeats, block)
}

func (t_ Timer) InitWithFireDateIntervalTargetSelectorUserInfoRepeats(date IDate, ti TimeInterval, t objc.IObject, s objc.Selector, ui objc.IObject, rep bool) Timer {
	rv := objc.CallMethod[Timer](t_, objc.GetSelector("initWithFireDate:interval:target:selector:userInfo:repeats:"), objc.ExtractPtr(date), ti, objc.ExtractPtr(t), s, objc.ExtractPtr(ui), rep)
	return rv
}

func Timer_InitWithFireDateIntervalTargetSelectorUserInfoRepeats(date IDate, ti TimeInterval, t objc.IObject, s objc.Selector, ui objc.IObject, rep bool) Timer {
	return TimerClass.Alloc().InitWithFireDateIntervalTargetSelectorUserInfoRepeats(date, ti, t, s, ui, rep)
}

func (tc _TimerClass) Alloc() Timer {
	rv := objc.CallMethod[Timer](tc, objc.GetSelector("alloc"))
	return rv
}

func Timer_Alloc() Timer {
	return TimerClass.Alloc()
}

func (tc _TimerClass) New() Timer {
	rv := objc.CallMethod[Timer](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTimer() Timer {
	return TimerClass.New()
}

func Timer_New() Timer {
	return TimerClass.New()
}

func (t_ Timer) Init() Timer {
	rv := objc.CallMethod[Timer](t_, objc.GetSelector("init"))
	return rv
}

func Timer_Init() Timer {
	return TimerClass.Alloc().Init()
}

func (tc _TimerClass) ScheduledTimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	rv := objc.CallMethod[Timer](tc, objc.GetSelector("scheduledTimerWithTimeInterval:repeats:block:"), interval, repeats, block)
	return rv
}

func Timer_ScheduledTimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	return TimerClass.ScheduledTimerWithTimeIntervalRepeatsBlock(interval, repeats, block)
}

func (tc _TimerClass) ScheduledTimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objc.IObject, aSelector objc.Selector, userInfo objc.IObject, yesOrNo bool) Timer {
	rv := objc.CallMethod[Timer](tc, objc.GetSelector("scheduledTimerWithTimeInterval:target:selector:userInfo:repeats:"), ti, objc.ExtractPtr(aTarget), aSelector, objc.ExtractPtr(userInfo), yesOrNo)
	return rv
}

func Timer_ScheduledTimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objc.IObject, aSelector objc.Selector, userInfo objc.IObject, yesOrNo bool) Timer {
	return TimerClass.ScheduledTimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti, aTarget, aSelector, userInfo, yesOrNo)
}

func (tc _TimerClass) TimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	rv := objc.CallMethod[Timer](tc, objc.GetSelector("timerWithTimeInterval:repeats:block:"), interval, repeats, block)
	return rv
}

func Timer_TimerWithTimeIntervalRepeatsBlock(interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	return TimerClass.TimerWithTimeIntervalRepeatsBlock(interval, repeats, block)
}

func (tc _TimerClass) TimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objc.IObject, aSelector objc.Selector, userInfo objc.IObject, yesOrNo bool) Timer {
	rv := objc.CallMethod[Timer](tc, objc.GetSelector("timerWithTimeInterval:target:selector:userInfo:repeats:"), ti, objc.ExtractPtr(aTarget), aSelector, objc.ExtractPtr(userInfo), yesOrNo)
	return rv
}

func Timer_TimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti TimeInterval, aTarget objc.IObject, aSelector objc.Selector, userInfo objc.IObject, yesOrNo bool) Timer {
	return TimerClass.TimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti, aTarget, aSelector, userInfo, yesOrNo)
}

func (t_ Timer) Fire() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("fire"))
}

func (t_ Timer) Invalidate() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("invalidate"))
}

func (t_ Timer) IsValid() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isValid"))
	return rv
}

func (t_ Timer) FireDate() Date {
	rv := objc.CallMethod[Date](t_, objc.GetSelector("fireDate"))
	return rv
}

func (t_ Timer) SetFireDate(value IDate) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setFireDate:"), objc.ExtractPtr(value))
}

func (t_ Timer) TimeInterval() TimeInterval {
	rv := objc.CallMethod[TimeInterval](t_, objc.GetSelector("timeInterval"))
	return rv
}

func (t_ Timer) UserInfo() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("userInfo"))
	return rv
}

func (t_ Timer) Tolerance() TimeInterval {
	rv := objc.CallMethod[TimeInterval](t_, objc.GetSelector("tolerance"))
	return rv
}

func (t_ Timer) SetTolerance(value TimeInterval) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTolerance:"), value)
}
