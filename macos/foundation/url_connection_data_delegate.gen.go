// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that most delegates of a URL connection implement to receive data associated with the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondatadelegate?language=objc
type PURLConnectionDataDelegate interface {
	// optional
	ConnectionDidReceiveData(connection URLConnection, data []byte)
	HasConnectionDidReceiveData() bool

	// optional
	ConnectionDidFinishLoading(connection URLConnection)
	HasConnectionDidFinishLoading() bool
}

// A delegate implementation builder for the [PURLConnectionDataDelegate] protocol.
type URLConnectionDataDelegate struct {
	_ConnectionDidReceiveData   func(connection URLConnection, data []byte)
	_ConnectionDidFinishLoading func(connection URLConnection)
}

func (di *URLConnectionDataDelegate) HasConnectionDidReceiveData() bool {
	return di._ConnectionDidReceiveData != nil
}

// Sent as a connection loads data incrementally. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondatadelegate/1414090-connection?language=objc
func (di *URLConnectionDataDelegate) SetConnectionDidReceiveData(f func(connection URLConnection, data []byte)) {
	di._ConnectionDidReceiveData = f
}

// Sent as a connection loads data incrementally. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondatadelegate/1414090-connection?language=objc
func (di *URLConnectionDataDelegate) ConnectionDidReceiveData(connection URLConnection, data []byte) {
	di._ConnectionDidReceiveData(connection, data)
}
func (di *URLConnectionDataDelegate) HasConnectionDidFinishLoading() bool {
	return di._ConnectionDidFinishLoading != nil
}

// Sent when a connection has finished loading successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondatadelegate/1416409-connectiondidfinishloading?language=objc
func (di *URLConnectionDataDelegate) SetConnectionDidFinishLoading(f func(connection URLConnection)) {
	di._ConnectionDidFinishLoading = f
}

// Sent when a connection has finished loading successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondatadelegate/1416409-connectiondidfinishloading?language=objc
func (di *URLConnectionDataDelegate) ConnectionDidFinishLoading(connection URLConnection) {
	di._ConnectionDidFinishLoading(connection)
}

// A concrete type wrapper for the [PURLConnectionDataDelegate] protocol.
type URLConnectionDataDelegateWrapper struct {
	objc.Object
}

func (u_ URLConnectionDataDelegateWrapper) HasConnectionDidReceiveData() bool {
	return u_.RespondsToSelector(objc.Sel("connection:didReceiveData:"))
}

// Sent as a connection loads data incrementally. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondatadelegate/1414090-connection?language=objc
func (u_ URLConnectionDataDelegateWrapper) ConnectionDidReceiveData(connection IURLConnection, data []byte) {
	objc.Call[objc.Void](u_, objc.Sel("connection:didReceiveData:"), objc.Ptr(connection), data)
}

func (u_ URLConnectionDataDelegateWrapper) HasConnectionDidFinishLoading() bool {
	return u_.RespondsToSelector(objc.Sel("connectionDidFinishLoading:"))
}

// Sent when a connection has finished loading successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondatadelegate/1416409-connectiondidfinishloading?language=objc
func (u_ URLConnectionDataDelegateWrapper) ConnectionDidFinishLoading(connection IURLConnection) {
	objc.Call[objc.Void](u_, objc.Sel("connectionDidFinishLoading:"), objc.Ptr(connection))
}
