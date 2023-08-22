// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableDateRangeMetadataGroup] class.
var MutableDateRangeMetadataGroupClass = _MutableDateRangeMetadataGroupClass{objc.GetClass("AVMutableDateRangeMetadataGroup")}

type _MutableDateRangeMetadataGroupClass struct {
	objc.Class
}

// An interface definition for the [MutableDateRangeMetadataGroup] class.
type IMutableDateRangeMetadataGroup interface {
	IDateRangeMetadataGroup
	SetItems(value []IMetadataItem)
	SetStartDate(value foundation.IDate)
	SetEndDate(value foundation.IDate)
}

// A mutable collection of metadata items that are valid for use within a specific range of dates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutabledaterangemetadatagroup?language=objc
type MutableDateRangeMetadataGroup struct {
	DateRangeMetadataGroup
}

func MutableDateRangeMetadataGroupFrom(ptr unsafe.Pointer) MutableDateRangeMetadataGroup {
	return MutableDateRangeMetadataGroup{
		DateRangeMetadataGroup: DateRangeMetadataGroupFrom(ptr),
	}
}

func (mc _MutableDateRangeMetadataGroupClass) Alloc() MutableDateRangeMetadataGroup {
	rv := objc.Call[MutableDateRangeMetadataGroup](mc, objc.Sel("alloc"))
	return rv
}

func MutableDateRangeMetadataGroup_Alloc() MutableDateRangeMetadataGroup {
	return MutableDateRangeMetadataGroupClass.Alloc()
}

func (mc _MutableDateRangeMetadataGroupClass) New() MutableDateRangeMetadataGroup {
	rv := objc.Call[MutableDateRangeMetadataGroup](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableDateRangeMetadataGroup() MutableDateRangeMetadataGroup {
	return MutableDateRangeMetadataGroupClass.New()
}

func (m_ MutableDateRangeMetadataGroup) Init() MutableDateRangeMetadataGroup {
	rv := objc.Call[MutableDateRangeMetadataGroup](m_, objc.Sel("init"))
	return rv
}

func (m_ MutableDateRangeMetadataGroup) InitWithItemsStartDateEndDate(items []IMetadataItem, startDate foundation.IDate, endDate foundation.IDate) MutableDateRangeMetadataGroup {
	rv := objc.Call[MutableDateRangeMetadataGroup](m_, objc.Sel("initWithItems:startDate:endDate:"), items, objc.Ptr(startDate), objc.Ptr(endDate))
	return rv
}

// Initializes an instance of AVDateRangeMetadataGroup with a collection of metadata items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdaterangemetadatagroup/1389614-initwithitems?language=objc
func NewMutableDateRangeMetadataGroupWithItemsStartDateEndDate(items []IMetadataItem, startDate foundation.IDate, endDate foundation.IDate) MutableDateRangeMetadataGroup {
	instance := MutableDateRangeMetadataGroupClass.Alloc().InitWithItemsStartDateEndDate(items, startDate, endDate)
	instance.Autorelease()
	return instance
}

// An array of associated metadata items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutabledaterangemetadatagroup/1388262-items?language=objc
func (m_ MutableDateRangeMetadataGroup) SetItems(value []IMetadataItem) {
	objc.Call[objc.Void](m_, objc.Sel("setItems:"), value)
}

// The start date for the metadata date range group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutabledaterangemetadatagroup/1390555-startdate?language=objc
func (m_ MutableDateRangeMetadataGroup) SetStartDate(value foundation.IDate) {
	objc.Call[objc.Void](m_, objc.Sel("setStartDate:"), objc.Ptr(value))
}

// The end date for the metadata date range group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutabledaterangemetadatagroup/1387651-enddate?language=objc
func (m_ MutableDateRangeMetadataGroup) SetEndDate(value foundation.IDate) {
	objc.Call[objc.Void](m_, objc.Sel("setEndDate:"), objc.Ptr(value))
}
