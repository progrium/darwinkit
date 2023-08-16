// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines methods that URL session instances call on their delegates to handle task-level events specific to stream tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionstreamdelegate?language=objc
type PURLSessionStreamDelegate interface {
	// optional
	URLSessionWriteClosedForStreamTask(session URLSession, streamTask URLSessionStreamTask)
	HasURLSessionWriteClosedForStreamTask() bool
}

// A delegate implementation builder for the [PURLSessionStreamDelegate] protocol.
type URLSessionStreamDelegate struct {
	_URLSessionWriteClosedForStreamTask func(session URLSession, streamTask URLSessionStreamTask)
}

func (di *URLSessionStreamDelegate) HasURLSessionWriteClosedForStreamTask() bool {
	return di._URLSessionWriteClosedForStreamTask != nil
}

// Tells the delegate that the write side of the underlying socket has been closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionstreamdelegate/1411507-urlsession?language=objc
func (di *URLSessionStreamDelegate) SetURLSessionWriteClosedForStreamTask(f func(session URLSession, streamTask URLSessionStreamTask)) {
	di._URLSessionWriteClosedForStreamTask = f
}

// Tells the delegate that the write side of the underlying socket has been closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionstreamdelegate/1411507-urlsession?language=objc
func (di *URLSessionStreamDelegate) URLSessionWriteClosedForStreamTask(session URLSession, streamTask URLSessionStreamTask) {
	di._URLSessionWriteClosedForStreamTask(session, streamTask)
}

// A concrete type wrapper for the [PURLSessionStreamDelegate] protocol.
type URLSessionStreamDelegateWrapper struct {
	objc.Object
}

func (u_ URLSessionStreamDelegateWrapper) HasURLSessionWriteClosedForStreamTask() bool {
	return u_.RespondsToSelector(objc.Sel("URLSession:writeClosedForStreamTask:"))
}

// Tells the delegate that the write side of the underlying socket has been closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionstreamdelegate/1411507-urlsession?language=objc
func (u_ URLSessionStreamDelegateWrapper) URLSessionWriteClosedForStreamTask(session IURLSession, streamTask IURLSessionStreamTask) {
	objc.Call[objc.Void](u_, objc.Sel("URLSession:writeClosedForStreamTask:"), objc.Ptr(session), objc.Ptr(streamTask))
}
