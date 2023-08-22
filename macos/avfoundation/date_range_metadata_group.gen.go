// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DateRangeMetadataGroup] class.
var DateRangeMetadataGroupClass = _DateRangeMetadataGroupClass{objc.GetClass("AVDateRangeMetadataGroup")}

type _DateRangeMetadataGroupClass struct {
	objc.Class
}

// An interface definition for the [DateRangeMetadataGroup] class.
type IDateRangeMetadataGroup interface {
	IMetadataGroup
	StartDate() foundation.Date
	EndDate() foundation.Date
}

// A collection of metadata items that are valid for use within a specific date range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdaterangemetadatagroup?language=objc
type DateRangeMetadataGroup struct {
	MetadataGroup
}

func DateRangeMetadataGroupFrom(ptr unsafe.Pointer) DateRangeMetadataGroup {
	return DateRangeMetadataGroup{
		MetadataGroup: MetadataGroupFrom(ptr),
	}
}

func (d_ DateRangeMetadataGroup) InitWithItemsStartDateEndDate(items []IMetadataItem, startDate foundation.IDate, endDate foundation.IDate) DateRangeMetadataGroup {
	rv := objc.Call[DateRangeMetadataGroup](d_, objc.Sel("initWithItems:startDate:endDate:"), items, objc.Ptr(startDate), objc.Ptr(endDate))
	return rv
}

// Initializes an instance of AVDateRangeMetadataGroup with a collection of metadata items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdaterangemetadatagroup/1389614-initwithitems?language=objc
func NewDateRangeMetadataGroupWithItemsStartDateEndDate(items []IMetadataItem, startDate foundation.IDate, endDate foundation.IDate) DateRangeMetadataGroup {
	instance := DateRangeMetadataGroupClass.Alloc().InitWithItemsStartDateEndDate(items, startDate, endDate)
	instance.Autorelease()
	return instance
}

func (dc _DateRangeMetadataGroupClass) Alloc() DateRangeMetadataGroup {
	rv := objc.Call[DateRangeMetadataGroup](dc, objc.Sel("alloc"))
	return rv
}

func DateRangeMetadataGroup_Alloc() DateRangeMetadataGroup {
	return DateRangeMetadataGroupClass.Alloc()
}

func (dc _DateRangeMetadataGroupClass) New() DateRangeMetadataGroup {
	rv := objc.Call[DateRangeMetadataGroup](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDateRangeMetadataGroup() DateRangeMetadataGroup {
	return DateRangeMetadataGroupClass.New()
}

func (d_ DateRangeMetadataGroup) Init() DateRangeMetadataGroup {
	rv := objc.Call[DateRangeMetadataGroup](d_, objc.Sel("init"))
	return rv
}

// The start date for the metadata date range group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdaterangemetadatagroup/1386420-startdate?language=objc
func (d_ DateRangeMetadataGroup) StartDate() foundation.Date {
	rv := objc.Call[foundation.Date](d_, objc.Sel("startDate"))
	return rv
}

// The end date for the metadata date range group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdaterangemetadatagroup/1386255-enddate?language=objc
func (d_ DateRangeMetadataGroup) EndDate() foundation.Date {
	rv := objc.Call[foundation.Date](d_, objc.Sel("endDate"))
	return rv
}
