// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLSessionTaskMetrics] class.
var URLSessionTaskMetricsClass = _URLSessionTaskMetricsClass{objc.GetClass("NSURLSessionTaskMetrics")}

type _URLSessionTaskMetricsClass struct {
	objc.Class
}

// An interface definition for the [URLSessionTaskMetrics] class.
type IURLSessionTaskMetrics interface {
	objc.IObject
	RedirectCount() uint
	TransactionMetrics() []URLSessionTaskTransactionMetrics
	TaskInterval() DateInterval
}

// An object encapsulating the metrics for a session task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontaskmetrics?language=objc
type URLSessionTaskMetrics struct {
	objc.Object
}

func URLSessionTaskMetricsFrom(ptr unsafe.Pointer) URLSessionTaskMetrics {
	return URLSessionTaskMetrics{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _URLSessionTaskMetricsClass) Alloc() URLSessionTaskMetrics {
	rv := objc.Call[URLSessionTaskMetrics](uc, objc.Sel("alloc"))
	return rv
}

func URLSessionTaskMetrics_Alloc() URLSessionTaskMetrics {
	return URLSessionTaskMetricsClass.Alloc()
}

func (uc _URLSessionTaskMetricsClass) New() URLSessionTaskMetrics {
	rv := objc.Call[URLSessionTaskMetrics](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLSessionTaskMetrics() URLSessionTaskMetrics {
	return URLSessionTaskMetricsClass.New()
}

func (u_ URLSessionTaskMetrics) Init() URLSessionTaskMetrics {
	rv := objc.Call[URLSessionTaskMetrics](u_, objc.Sel("init"))
	return rv
}

// The number of redirects that occurred during the execution of the task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontaskmetrics/1643136-redirectcount?language=objc
func (u_ URLSessionTaskMetrics) RedirectCount() uint {
	rv := objc.Call[uint](u_, objc.Sel("redirectCount"))
	return rv
}

// An array of metrics for each individual request-response transaction made during the execution of the task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontaskmetrics/1642789-transactionmetrics?language=objc
func (u_ URLSessionTaskMetrics) TransactionMetrics() []URLSessionTaskTransactionMetrics {
	rv := objc.Call[[]URLSessionTaskTransactionMetrics](u_, objc.Sel("transactionMetrics"))
	return rv
}

// The time interval between when a task is instantiated and when the task is completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontaskmetrics/1643169-taskinterval?language=objc
func (u_ URLSessionTaskMetrics) TaskInterval() DateInterval {
	rv := objc.Call[DateInterval](u_, objc.Sel("taskInterval"))
	return rv
}
