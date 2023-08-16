// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines methods that URL session instances call on their delegates to handle task-level events specific to data and upload tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondatadelegate?language=objc
type PURLSessionDataDelegate interface {
	// optional
	URLSessionDataTaskDidReceiveData(session URLSession, dataTask URLSessionDataTask, data []byte)
	HasURLSessionDataTaskDidReceiveData() bool
}

// A delegate implementation builder for the [PURLSessionDataDelegate] protocol.
type URLSessionDataDelegate struct {
	_URLSessionDataTaskDidReceiveData func(session URLSession, dataTask URLSessionDataTask, data []byte)
}

func (di *URLSessionDataDelegate) HasURLSessionDataTaskDidReceiveData() bool {
	return di._URLSessionDataTaskDidReceiveData != nil
}

// Tells the delegate that the data task has received some of the expected data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondatadelegate/1411528-urlsession?language=objc
func (di *URLSessionDataDelegate) SetURLSessionDataTaskDidReceiveData(f func(session URLSession, dataTask URLSessionDataTask, data []byte)) {
	di._URLSessionDataTaskDidReceiveData = f
}

// Tells the delegate that the data task has received some of the expected data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondatadelegate/1411528-urlsession?language=objc
func (di *URLSessionDataDelegate) URLSessionDataTaskDidReceiveData(session URLSession, dataTask URLSessionDataTask, data []byte) {
	di._URLSessionDataTaskDidReceiveData(session, dataTask, data)
}

// A concrete type wrapper for the [PURLSessionDataDelegate] protocol.
type URLSessionDataDelegateWrapper struct {
	objc.Object
}

func (u_ URLSessionDataDelegateWrapper) HasURLSessionDataTaskDidReceiveData() bool {
	return u_.RespondsToSelector(objc.Sel("URLSession:dataTask:didReceiveData:"))
}

// Tells the delegate that the data task has received some of the expected data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondatadelegate/1411528-urlsession?language=objc
func (u_ URLSessionDataDelegateWrapper) URLSessionDataTaskDidReceiveData(session IURLSession, dataTask IURLSessionDataTask, data []byte) {
	objc.Call[objc.Void](u_, objc.Sel("URLSession:dataTask:didReceiveData:"), objc.Ptr(session), objc.Ptr(dataTask), data)
}
