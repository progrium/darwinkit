// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines methods that URL session instances call on their delegates to handle task-level events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontaskdelegate?language=objc
type PURLSessionTaskDelegate interface {
	// optional
	URLSessionTaskDidFinishCollectingMetrics(session URLSession, task URLSessionTask, metrics URLSessionTaskMetrics)
	HasURLSessionTaskDidFinishCollectingMetrics() bool
}

// A delegate implementation builder for the [PURLSessionTaskDelegate] protocol.
type URLSessionTaskDelegate struct {
	_URLSessionTaskDidFinishCollectingMetrics func(session URLSession, task URLSessionTask, metrics URLSessionTaskMetrics)
}

func (di *URLSessionTaskDelegate) HasURLSessionTaskDidFinishCollectingMetrics() bool {
	return di._URLSessionTaskDidFinishCollectingMetrics != nil
}

// Tells the delegate that the session finished collecting metrics for the task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontaskdelegate/1643148-urlsession?language=objc
func (di *URLSessionTaskDelegate) SetURLSessionTaskDidFinishCollectingMetrics(f func(session URLSession, task URLSessionTask, metrics URLSessionTaskMetrics)) {
	di._URLSessionTaskDidFinishCollectingMetrics = f
}

// Tells the delegate that the session finished collecting metrics for the task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontaskdelegate/1643148-urlsession?language=objc
func (di *URLSessionTaskDelegate) URLSessionTaskDidFinishCollectingMetrics(session URLSession, task URLSessionTask, metrics URLSessionTaskMetrics) {
	di._URLSessionTaskDidFinishCollectingMetrics(session, task, metrics)
}

// A concrete type wrapper for the [PURLSessionTaskDelegate] protocol.
type URLSessionTaskDelegateWrapper struct {
	objc.Object
}

func (u_ URLSessionTaskDelegateWrapper) HasURLSessionTaskDidFinishCollectingMetrics() bool {
	return u_.RespondsToSelector(objc.Sel("URLSession:task:didFinishCollectingMetrics:"))
}

// Tells the delegate that the session finished collecting metrics for the task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontaskdelegate/1643148-urlsession?language=objc
func (u_ URLSessionTaskDelegateWrapper) URLSessionTaskDidFinishCollectingMetrics(session IURLSession, task IURLSessionTask, metrics IURLSessionTaskMetrics) {
	objc.Call[objc.Void](u_, objc.Sel("URLSession:task:didFinishCollectingMetrics:"), objc.Ptr(session), objc.Ptr(task), objc.Ptr(metrics))
}
