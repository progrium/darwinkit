// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var DateClass = _DateClass{objc.GetClass("NSDate")}

type _DateClass struct {
	objc.Class
}

type IDate interface {
	objc.IObject
	IsEqualToDate(otherDate IDate) bool
	EarlierDate(anotherDate IDate) Date
	LaterDate(anotherDate IDate) Date
	Compare(other IDate) ComparisonResult
	TimeIntervalSinceDate(anotherDate IDate) TimeInterval
	DescriptionWithLocale(locale objc.IObject) string
	TimeIntervalSinceNow() TimeInterval
	TimeIntervalSinceReferenceDate() TimeInterval
	TimeIntervalSince1970() TimeInterval
	Description() string
}

type Date struct {
	objc.Object
}

func MakeDate(ptr unsafe.Pointer) Date {
	return Date{
		Object: objc.MakeObject(ptr),
	}
}

func (dc _DateClass) Date() Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("date"))
	return rv
}

func Date_Date() Date {
	return DateClass.Date()
}

func (dc _DateClass) DateWithTimeIntervalSinceNow(secs TimeInterval) Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("dateWithTimeIntervalSinceNow:"), secs)
	return rv
}

func Date_DateWithTimeIntervalSinceNow(secs TimeInterval) Date {
	return DateClass.DateWithTimeIntervalSinceNow(secs)
}

func (dc _DateClass) DateWithTimeIntervalSinceDate(secsToBeAdded TimeInterval, date IDate) Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("dateWithTimeInterval:sinceDate:"), secsToBeAdded, objc.ExtractPtr(date))
	return rv
}

func Date_DateWithTimeIntervalSinceDate(secsToBeAdded TimeInterval, date IDate) Date {
	return DateClass.DateWithTimeIntervalSinceDate(secsToBeAdded, date)
}

func (dc _DateClass) DateWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("dateWithTimeIntervalSinceReferenceDate:"), ti)
	return rv
}

func Date_DateWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	return DateClass.DateWithTimeIntervalSinceReferenceDate(ti)
}

func (dc _DateClass) DateWithTimeIntervalSince1970(secs TimeInterval) Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("dateWithTimeIntervalSince1970:"), secs)
	return rv
}

func Date_DateWithTimeIntervalSince1970(secs TimeInterval) Date {
	return DateClass.DateWithTimeIntervalSince1970(secs)
}

func (d_ Date) Init() Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("init"))
	return rv
}

func Date_Init() Date {
	return DateClass.Alloc().Init()
}

func (d_ Date) InitWithTimeIntervalSinceNow(secs TimeInterval) Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("initWithTimeIntervalSinceNow:"), secs)
	return rv
}

func Date_InitWithTimeIntervalSinceNow(secs TimeInterval) Date {
	return DateClass.Alloc().InitWithTimeIntervalSinceNow(secs)
}

func (d_ Date) InitWithTimeIntervalSinceDate(secsToBeAdded TimeInterval, date IDate) Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("initWithTimeInterval:sinceDate:"), secsToBeAdded, objc.ExtractPtr(date))
	return rv
}

func Date_InitWithTimeIntervalSinceDate(secsToBeAdded TimeInterval, date IDate) Date {
	return DateClass.Alloc().InitWithTimeIntervalSinceDate(secsToBeAdded, date)
}

func (d_ Date) InitWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("initWithTimeIntervalSinceReferenceDate:"), ti)
	return rv
}

func Date_InitWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	return DateClass.Alloc().InitWithTimeIntervalSinceReferenceDate(ti)
}

func (d_ Date) InitWithTimeIntervalSince1970(secs TimeInterval) Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("initWithTimeIntervalSince1970:"), secs)
	return rv
}

func Date_InitWithTimeIntervalSince1970(secs TimeInterval) Date {
	return DateClass.Alloc().InitWithTimeIntervalSince1970(secs)
}

func (d_ Date) DateByAddingTimeInterval(ti TimeInterval) Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("dateByAddingTimeInterval:"), ti)
	return rv
}

func Date_DateByAddingTimeInterval(ti TimeInterval) Date {
	return DateClass.Alloc().DateByAddingTimeInterval(ti)
}

func (dc _DateClass) Alloc() Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("alloc"))
	return rv
}

func Date_Alloc() Date {
	return DateClass.Alloc()
}

func (dc _DateClass) New() Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDate() Date {
	return DateClass.New()
}

func Date_New() Date {
	return DateClass.New()
}

func (d_ Date) IsEqualToDate(otherDate IDate) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isEqualToDate:"), objc.ExtractPtr(otherDate))
	return rv
}

func (d_ Date) EarlierDate(anotherDate IDate) Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("earlierDate:"), objc.ExtractPtr(anotherDate))
	return rv
}

func (d_ Date) LaterDate(anotherDate IDate) Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("laterDate:"), objc.ExtractPtr(anotherDate))
	return rv
}

func (d_ Date) Compare(other IDate) ComparisonResult {
	rv := objc.CallMethod[ComparisonResult](d_, objc.GetSelector("compare:"), objc.ExtractPtr(other))
	return rv
}

func (d_ Date) TimeIntervalSinceDate(anotherDate IDate) TimeInterval {
	rv := objc.CallMethod[TimeInterval](d_, objc.GetSelector("timeIntervalSinceDate:"), objc.ExtractPtr(anotherDate))
	return rv
}

func (d_ Date) DescriptionWithLocale(locale objc.IObject) string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("descriptionWithLocale:"), objc.ExtractPtr(locale))
	return rv
}

func (dc _DateClass) DistantFuture() Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("distantFuture"))
	return rv
}

func Date_DistantFuture() Date {
	return DateClass.DistantFuture()
}

func (dc _DateClass) DistantPast() Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("distantPast"))
	return rv
}

func Date_DistantPast() Date {
	return DateClass.DistantPast()
}

func (dc _DateClass) Now() Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("now"))
	return rv
}

func Date_Now() Date {
	return DateClass.Now()
}

func (d_ Date) TimeIntervalSinceNow() TimeInterval {
	rv := objc.CallMethod[TimeInterval](d_, objc.GetSelector("timeIntervalSinceNow"))
	return rv
}

func (d_ Date) TimeIntervalSinceReferenceDate() TimeInterval {
	rv := objc.CallMethod[TimeInterval](d_, objc.GetSelector("timeIntervalSinceReferenceDate"))
	return rv
}

func (d_ Date) TimeIntervalSince1970() TimeInterval {
	rv := objc.CallMethod[TimeInterval](d_, objc.GetSelector("timeIntervalSince1970"))
	return rv
}

func (d_ Date) Description() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("description"))
	return rv
}
