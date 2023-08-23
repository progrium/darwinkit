// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [QuartzFilter] class.
var QuartzFilterClass = _QuartzFilterClass{objc.GetClass("QuartzFilter")}

type _QuartzFilterClass struct {
	objc.Class
}

// An interface definition for the [QuartzFilter] class.
type IQuartzFilter interface {
	objc.IObject
	ApplyToContext(aContext coregraphics.ContextRef) bool
	RemoveFromContext(aContext coregraphics.ContextRef)
	LocalizedName() string
	Properties() foundation.Dictionary
	Url() foundation.URL
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfilter?language=objc
type QuartzFilter struct {
	objc.Object
}

func QuartzFilterFrom(ptr unsafe.Pointer) QuartzFilter {
	return QuartzFilter{
		Object: objc.ObjectFrom(ptr),
	}
}

func (qc _QuartzFilterClass) Alloc() QuartzFilter {
	rv := objc.Call[QuartzFilter](qc, objc.Sel("alloc"))
	return rv
}

func QuartzFilter_Alloc() QuartzFilter {
	return QuartzFilterClass.Alloc()
}

func (qc _QuartzFilterClass) New() QuartzFilter {
	rv := objc.Call[QuartzFilter](qc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewQuartzFilter() QuartzFilter {
	return QuartzFilterClass.New()
}

func (q_ QuartzFilter) Init() QuartzFilter {
	rv := objc.Call[QuartzFilter](q_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfilter/1433684-applytocontext?language=objc
func (q_ QuartzFilter) ApplyToContext(aContext coregraphics.ContextRef) bool {
	rv := objc.Call[bool](q_, objc.Sel("applyToContext:"), aContext)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfilter/1433679-removefromcontext?language=objc
func (q_ QuartzFilter) RemoveFromContext(aContext coregraphics.ContextRef) {
	objc.Call[objc.Void](q_, objc.Sel("removeFromContext:"), aContext)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfilter/1433671-quartzfilterwithurl?language=objc
func (qc _QuartzFilterClass) QuartzFilterWithURL(aURL foundation.IURL) QuartzFilter {
	rv := objc.Call[QuartzFilter](qc, objc.Sel("quartzFilterWithURL:"), objc.Ptr(aURL))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfilter/1433671-quartzfilterwithurl?language=objc
func QuartzFilter_QuartzFilterWithURL(aURL foundation.IURL) QuartzFilter {
	return QuartzFilterClass.QuartzFilterWithURL(aURL)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfilter/1433675-quartzfilterwithoutputintents?language=objc
func (qc _QuartzFilterClass) QuartzFilterWithOutputIntents(outputIntents []objc.IObject) QuartzFilter {
	rv := objc.Call[QuartzFilter](qc, objc.Sel("quartzFilterWithOutputIntents:"), outputIntents)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfilter/1433675-quartzfilterwithoutputintents?language=objc
func QuartzFilter_QuartzFilterWithOutputIntents(outputIntents []objc.IObject) QuartzFilter {
	return QuartzFilterClass.QuartzFilterWithOutputIntents(outputIntents)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfilter/1433673-quartzfilterwithproperties?language=objc
func (qc _QuartzFilterClass) QuartzFilterWithProperties(properties foundation.Dictionary) QuartzFilter {
	rv := objc.Call[QuartzFilter](qc, objc.Sel("quartzFilterWithProperties:"), properties)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfilter/1433673-quartzfilterwithproperties?language=objc
func QuartzFilter_QuartzFilterWithProperties(properties foundation.Dictionary) QuartzFilter {
	return QuartzFilterClass.QuartzFilterWithProperties(properties)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfilter/1433682-localizedname?language=objc
func (q_ QuartzFilter) LocalizedName() string {
	rv := objc.Call[string](q_, objc.Sel("localizedName"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfilter/1433686-properties?language=objc
func (q_ QuartzFilter) Properties() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](q_, objc.Sel("properties"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfilter/1433677-url?language=objc
func (q_ QuartzFilter) Url() foundation.URL {
	rv := objc.Call[foundation.URL](q_, objc.Sel("url"))
	return rv
}
