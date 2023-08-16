// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that delegates of a URL connection implement to receive status about and provide feedback to the connection object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate?language=objc
type PURLConnectionDelegate interface {
	// optional
	ConnectionShouldUseCredentialStorage(connection URLConnection) bool
	HasConnectionShouldUseCredentialStorage() bool

	// optional
	ConnectionDidFailWithError(connection URLConnection, error Error)
	HasConnectionDidFailWithError() bool
}

// A delegate implementation builder for the [PURLConnectionDelegate] protocol.
type URLConnectionDelegate struct {
	_ConnectionShouldUseCredentialStorage func(connection URLConnection) bool
	_ConnectionDidFailWithError           func(connection URLConnection, error Error)
}

func (di *URLConnectionDelegate) HasConnectionShouldUseCredentialStorage() bool {
	return di._ConnectionShouldUseCredentialStorage != nil
}

// Sent to determine whether the URL loader should use the credential storage for authenticating the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate/1414890-connectionshouldusecredentialsto?language=objc
func (di *URLConnectionDelegate) SetConnectionShouldUseCredentialStorage(f func(connection URLConnection) bool) {
	di._ConnectionShouldUseCredentialStorage = f
}

// Sent to determine whether the URL loader should use the credential storage for authenticating the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate/1414890-connectionshouldusecredentialsto?language=objc
func (di *URLConnectionDelegate) ConnectionShouldUseCredentialStorage(connection URLConnection) bool {
	return di._ConnectionShouldUseCredentialStorage(connection)
}
func (di *URLConnectionDelegate) HasConnectionDidFailWithError() bool {
	return di._ConnectionDidFailWithError != nil
}

// Sent when a connection fails to load its request successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate/1418443-connection?language=objc
func (di *URLConnectionDelegate) SetConnectionDidFailWithError(f func(connection URLConnection, error Error)) {
	di._ConnectionDidFailWithError = f
}

// Sent when a connection fails to load its request successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate/1418443-connection?language=objc
func (di *URLConnectionDelegate) ConnectionDidFailWithError(connection URLConnection, error Error) {
	di._ConnectionDidFailWithError(connection, error)
}

// A concrete type wrapper for the [PURLConnectionDelegate] protocol.
type URLConnectionDelegateWrapper struct {
	objc.Object
}

func (u_ URLConnectionDelegateWrapper) HasConnectionShouldUseCredentialStorage() bool {
	return u_.RespondsToSelector(objc.Sel("connectionShouldUseCredentialStorage:"))
}

// Sent to determine whether the URL loader should use the credential storage for authenticating the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate/1414890-connectionshouldusecredentialsto?language=objc
func (u_ URLConnectionDelegateWrapper) ConnectionShouldUseCredentialStorage(connection IURLConnection) bool {
	rv := objc.Call[bool](u_, objc.Sel("connectionShouldUseCredentialStorage:"), objc.Ptr(connection))
	return rv
}

func (u_ URLConnectionDelegateWrapper) HasConnectionDidFailWithError() bool {
	return u_.RespondsToSelector(objc.Sel("connection:didFailWithError:"))
}

// Sent when a connection fails to load its request successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate/1418443-connection?language=objc
func (u_ URLConnectionDelegateWrapper) ConnectionDidFailWithError(connection IURLConnection, error IError) {
	objc.Call[objc.Void](u_, objc.Sel("connection:didFailWithError:"), objc.Ptr(connection), objc.Ptr(error))
}
