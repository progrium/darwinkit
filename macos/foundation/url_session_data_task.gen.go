// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLSessionDataTask] class.
var URLSessionDataTaskClass = _URLSessionDataTaskClass{objc.GetClass("NSURLSessionDataTask")}

type _URLSessionDataTaskClass struct {
	objc.Class
}

// An interface definition for the [URLSessionDataTask] class.
type IURLSessionDataTask interface {
	IURLSessionTask
}

// A URL session task that returns downloaded data directly to the app in memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondatatask?language=objc
type URLSessionDataTask struct {
	URLSessionTask
}

func URLSessionDataTaskFrom(ptr unsafe.Pointer) URLSessionDataTask {
	return URLSessionDataTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}

func (uc _URLSessionDataTaskClass) Alloc() URLSessionDataTask {
	rv := objc.Call[URLSessionDataTask](uc, objc.Sel("alloc"))
	return rv
}

func URLSessionDataTask_Alloc() URLSessionDataTask {
	return URLSessionDataTaskClass.Alloc()
}

func (uc _URLSessionDataTaskClass) New() URLSessionDataTask {
	rv := objc.Call[URLSessionDataTask](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLSessionDataTask() URLSessionDataTask {
	return URLSessionDataTaskClass.New()
}

func (u_ URLSessionDataTask) Init() URLSessionDataTask {
	rv := objc.Call[URLSessionDataTask](u_, objc.Sel("init"))
	return rv
}
