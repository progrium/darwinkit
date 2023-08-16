// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLSessionDownloadTask] class.
var URLSessionDownloadTaskClass = _URLSessionDownloadTaskClass{objc.GetClass("NSURLSessionDownloadTask")}

type _URLSessionDownloadTaskClass struct {
	objc.Class
}

// An interface definition for the [URLSessionDownloadTask] class.
type IURLSessionDownloadTask interface {
	IURLSessionTask
	CancelByProducingResumeData(completionHandler func(resumeData []byte))
}

// A URL session task that stores downloaded data to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondownloadtask?language=objc
type URLSessionDownloadTask struct {
	URLSessionTask
}

func URLSessionDownloadTaskFrom(ptr unsafe.Pointer) URLSessionDownloadTask {
	return URLSessionDownloadTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}

func (uc _URLSessionDownloadTaskClass) Alloc() URLSessionDownloadTask {
	rv := objc.Call[URLSessionDownloadTask](uc, objc.Sel("alloc"))
	return rv
}

func URLSessionDownloadTask_Alloc() URLSessionDownloadTask {
	return URLSessionDownloadTaskClass.Alloc()
}

func (uc _URLSessionDownloadTaskClass) New() URLSessionDownloadTask {
	rv := objc.Call[URLSessionDownloadTask](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLSessionDownloadTask() URLSessionDownloadTask {
	return URLSessionDownloadTaskClass.New()
}

func (u_ URLSessionDownloadTask) Init() URLSessionDownloadTask {
	rv := objc.Call[URLSessionDownloadTask](u_, objc.Sel("init"))
	return rv
}

// Cancels a download and calls a callback with resume data for later use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondownloadtask/1411634-cancelbyproducingresumedata?language=objc
func (u_ URLSessionDownloadTask) CancelByProducingResumeData(completionHandler func(resumeData []byte)) {
	objc.Call[objc.Void](u_, objc.Sel("cancelByProducingResumeData:"), completionHandler)
}
