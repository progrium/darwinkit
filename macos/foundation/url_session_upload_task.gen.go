// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLSessionUploadTask] class.
var URLSessionUploadTaskClass = _URLSessionUploadTaskClass{objc.GetClass("NSURLSessionUploadTask")}

type _URLSessionUploadTaskClass struct {
	objc.Class
}

// An interface definition for the [URLSessionUploadTask] class.
type IURLSessionUploadTask interface {
	IURLSessionDataTask
}

// A URL session task that uploads data to the network in a request body. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionuploadtask?language=objc
type URLSessionUploadTask struct {
	URLSessionDataTask
}

func URLSessionUploadTaskFrom(ptr unsafe.Pointer) URLSessionUploadTask {
	return URLSessionUploadTask{
		URLSessionDataTask: URLSessionDataTaskFrom(ptr),
	}
}

func (uc _URLSessionUploadTaskClass) Alloc() URLSessionUploadTask {
	rv := objc.Call[URLSessionUploadTask](uc, objc.Sel("alloc"))
	return rv
}

func URLSessionUploadTask_Alloc() URLSessionUploadTask {
	return URLSessionUploadTaskClass.Alloc()
}

func (uc _URLSessionUploadTaskClass) New() URLSessionUploadTask {
	rv := objc.Call[URLSessionUploadTask](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLSessionUploadTask() URLSessionUploadTask {
	return URLSessionUploadTaskClass.New()
}

func (u_ URLSessionUploadTask) Init() URLSessionUploadTask {
	rv := objc.Call[URLSessionUploadTask](u_, objc.Sel("init"))
	return rv
}
