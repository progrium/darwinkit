// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLSession] class.
var URLSessionClass = _URLSessionClass{objc.GetClass("NSURLSession")}

type _URLSessionClass struct {
	objc.Class
}

// An interface definition for the [URLSession] class.
type IURLSession interface {
	objc.IObject
	ResetWithCompletionHandler(completionHandler func())
	UploadTaskWithRequestFromFile(request IURLRequest, fileURL IURL) URLSessionUploadTask
	FlushWithCompletionHandler(completionHandler func())
	WebSocketTaskWithRequest(request IURLRequest) URLSessionWebSocketTask
	GetTasksWithCompletionHandler(completionHandler func(dataTasks []URLSessionDataTask, uploadTasks []URLSessionUploadTask, downloadTasks []URLSessionDownloadTask))
	WebSocketTaskWithURL(url IURL) URLSessionWebSocketTask
	StreamTaskWithHostNamePort(hostname string, port int) URLSessionStreamTask
	DownloadTaskWithResumeData(resumeData []byte) URLSessionDownloadTask
	DownloadTaskWithRequest(request IURLRequest) URLSessionDownloadTask
	UploadTaskWithStreamedRequest(request IURLRequest) URLSessionUploadTask
	InvalidateAndCancel()
	DataTaskWithRequest(request IURLRequest) URLSessionDataTask
	DownloadTaskWithURL(url IURL) URLSessionDownloadTask
	GetAllTasksWithCompletionHandler(completionHandler func(tasks []URLSessionTask))
	DataTaskWithURL(url IURL) URLSessionDataTask
	FinishTasksAndInvalidate()
	Delegate() URLSessionDelegateWrapper
	Configuration() URLSessionConfiguration
	SessionDescription() string
	SetSessionDescription(value string)
	DelegateQueue() OperationQueue
}

// An object that coordinates a group of related, network data transfer tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession?language=objc
type URLSession struct {
	objc.Object
}

func URLSessionFrom(ptr unsafe.Pointer) URLSession {
	return URLSession{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _URLSessionClass) Alloc() URLSession {
	rv := objc.Call[URLSession](uc, objc.Sel("alloc"))
	return rv
}

func URLSession_Alloc() URLSession {
	return URLSessionClass.Alloc()
}

func (uc _URLSessionClass) New() URLSession {
	rv := objc.Call[URLSession](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLSession() URLSession {
	return URLSessionClass.New()
}

func (u_ URLSession) Init() URLSession {
	rv := objc.Call[URLSession](u_, objc.Sel("init"))
	return rv
}

// Empties all cookies, caches and credential stores, removes disk files, flushes in-progress downloads to disk, and ensures that future requests occur on a new socket. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411479-resetwithcompletionhandler?language=objc
func (u_ URLSession) ResetWithCompletionHandler(completionHandler func()) {
	objc.Call[objc.Void](u_, objc.Sel("resetWithCompletionHandler:"), completionHandler)
}

// Creates a task that performs an HTTP request for uploading the specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411550-uploadtaskwithrequest?language=objc
func (u_ URLSession) UploadTaskWithRequestFromFile(request IURLRequest, fileURL IURL) URLSessionUploadTask {
	rv := objc.Call[URLSessionUploadTask](u_, objc.Sel("uploadTaskWithRequest:fromFile:"), objc.Ptr(request), objc.Ptr(fileURL))
	return rv
}

// Flushes cookies and credentials to disk, clears transient caches, and ensures that future requests occur on a new TCP connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411622-flushwithcompletionhandler?language=objc
func (u_ URLSession) FlushWithCompletionHandler(completionHandler func()) {
	objc.Call[objc.Void](u_, objc.Sel("flushWithCompletionHandler:"), completionHandler)
}

// Creates a WebSocket task for the provided URL request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/3235750-websockettaskwithrequest?language=objc
func (u_ URLSession) WebSocketTaskWithRequest(request IURLRequest) URLSessionWebSocketTask {
	rv := objc.Call[URLSessionWebSocketTask](u_, objc.Sel("webSocketTaskWithRequest:"), objc.Ptr(request))
	return rv
}

// Asynchronously calls a completion callback with all data, upload, and download tasks in a session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411578-gettaskswithcompletionhandler?language=objc
func (u_ URLSession) GetTasksWithCompletionHandler(completionHandler func(dataTasks []URLSessionDataTask, uploadTasks []URLSessionUploadTask, downloadTasks []URLSessionDownloadTask)) {
	objc.Call[objc.Void](u_, objc.Sel("getTasksWithCompletionHandler:"), completionHandler)
}

// Creates a WebSocket task for the provided URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/3181171-websockettaskwithurl?language=objc
func (u_ URLSession) WebSocketTaskWithURL(url IURL) URLSessionWebSocketTask {
	rv := objc.Call[URLSessionWebSocketTask](u_, objc.Sel("webSocketTaskWithURL:"), objc.Ptr(url))
	return rv
}

// Creates a task that establishes a bidirectional TCP/IP connection to a specified hostname and port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411587-streamtaskwithhostname?language=objc
func (u_ URLSession) StreamTaskWithHostNamePort(hostname string, port int) URLSessionStreamTask {
	rv := objc.Call[URLSessionStreamTask](u_, objc.Sel("streamTaskWithHostName:port:"), hostname, port)
	return rv
}

// Creates a download task to resume a previously canceled or failed download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1409226-downloadtaskwithresumedata?language=objc
func (u_ URLSession) DownloadTaskWithResumeData(resumeData []byte) URLSessionDownloadTask {
	rv := objc.Call[URLSessionDownloadTask](u_, objc.Sel("downloadTaskWithResumeData:"), resumeData)
	return rv
}

// Creates a download task that retrieves the contents of a URL based on the specified URL request object and saves the results to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411481-downloadtaskwithrequest?language=objc
func (u_ URLSession) DownloadTaskWithRequest(request IURLRequest) URLSessionDownloadTask {
	rv := objc.Call[URLSessionDownloadTask](u_, objc.Sel("downloadTaskWithRequest:"), objc.Ptr(request))
	return rv
}

// Creates a task that performs an HTTP request for uploading data based on the specified URL request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1410934-uploadtaskwithstreamedrequest?language=objc
func (u_ URLSession) UploadTaskWithStreamedRequest(request IURLRequest) URLSessionUploadTask {
	rv := objc.Call[URLSessionUploadTask](u_, objc.Sel("uploadTaskWithStreamedRequest:"), objc.Ptr(request))
	return rv
}

// Cancels all outstanding tasks and then invalidates the session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411538-invalidateandcancel?language=objc
func (u_ URLSession) InvalidateAndCancel() {
	objc.Call[objc.Void](u_, objc.Sel("invalidateAndCancel"))
}

// Creates a task that retrieves the contents of a URL based on the specified URL request object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1410592-datataskwithrequest?language=objc
func (u_ URLSession) DataTaskWithRequest(request IURLRequest) URLSessionDataTask {
	rv := objc.Call[URLSessionDataTask](u_, objc.Sel("dataTaskWithRequest:"), objc.Ptr(request))
	return rv
}

// Creates a session with the specified session configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411474-sessionwithconfiguration?language=objc
func (uc _URLSessionClass) SessionWithConfiguration(configuration IURLSessionConfiguration) URLSession {
	rv := objc.Call[URLSession](uc, objc.Sel("sessionWithConfiguration:"), objc.Ptr(configuration))
	return rv
}

// Creates a session with the specified session configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411474-sessionwithconfiguration?language=objc
func URLSession_SessionWithConfiguration(configuration IURLSessionConfiguration) URLSession {
	return URLSessionClass.SessionWithConfiguration(configuration)
}

// Creates a download task that retrieves the contents of the specified URL and saves the results to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411482-downloadtaskwithurl?language=objc
func (u_ URLSession) DownloadTaskWithURL(url IURL) URLSessionDownloadTask {
	rv := objc.Call[URLSessionDownloadTask](u_, objc.Sel("downloadTaskWithURL:"), objc.Ptr(url))
	return rv
}

// Asynchronously calls a completion callback with all tasks in a session [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411618-getalltaskswithcompletionhandler?language=objc
func (u_ URLSession) GetAllTasksWithCompletionHandler(completionHandler func(tasks []URLSessionTask)) {
	objc.Call[objc.Void](u_, objc.Sel("getAllTasksWithCompletionHandler:"), completionHandler)
}

// Creates a task that retrieves the contents of the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411554-datataskwithurl?language=objc
func (u_ URLSession) DataTaskWithURL(url IURL) URLSessionDataTask {
	rv := objc.Call[URLSessionDataTask](u_, objc.Sel("dataTaskWithURL:"), objc.Ptr(url))
	return rv
}

// Invalidates the session, allowing any outstanding tasks to finish. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1407428-finishtasksandinvalidate?language=objc
func (u_ URLSession) FinishTasksAndInvalidate() {
	objc.Call[objc.Void](u_, objc.Sel("finishTasksAndInvalidate"))
}

// The delegate assigned when this object was created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411530-delegate?language=objc
func (u_ URLSession) Delegate() URLSessionDelegateWrapper {
	rv := objc.Call[URLSessionDelegateWrapper](u_, objc.Sel("delegate"))
	return rv
}

// A copy of the configuration object for this session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411477-configuration?language=objc
func (u_ URLSession) Configuration() URLSessionConfiguration {
	rv := objc.Call[URLSessionConfiguration](u_, objc.Sel("configuration"))
	return rv
}

// An app-defined descriptive label for the session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1408277-sessiondescription?language=objc
func (u_ URLSession) SessionDescription() string {
	rv := objc.Call[string](u_, objc.Sel("sessionDescription"))
	return rv
}

// An app-defined descriptive label for the session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1408277-sessiondescription?language=objc
func (u_ URLSession) SetSessionDescription(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setSessionDescription:"), value)
}

// The operation queue provided when this object was created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411571-delegatequeue?language=objc
func (u_ URLSession) DelegateQueue() OperationQueue {
	rv := objc.Call[OperationQueue](u_, objc.Sel("delegateQueue"))
	return rv
}

// The shared singleton session object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1409000-sharedsession?language=objc
func (uc _URLSessionClass) SharedSession() URLSession {
	rv := objc.Call[URLSession](uc, objc.Sel("sharedSession"))
	return rv
}

// The shared singleton session object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1409000-sharedsession?language=objc
func URLSession_SharedSession() URLSession {
	return URLSessionClass.SharedSession()
}
