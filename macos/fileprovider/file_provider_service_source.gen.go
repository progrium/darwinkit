// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A service that provides a custom communication channel between the host app and the File Provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderservicesource?language=objc
type PFileProviderServiceSource interface {
	// optional
	MakeListenerEndpointAndReturnError(error foundation.Error) foundation.IXPCListenerEndpoint
	HasMakeListenerEndpointAndReturnError() bool

	// optional
	ServiceName() foundation.FileProviderServiceName
	HasServiceName() bool
}

// A concrete type wrapper for the [PFileProviderServiceSource] protocol.
type FileProviderServiceSourceWrapper struct {
	objc.Object
}

func (f_ FileProviderServiceSourceWrapper) HasMakeListenerEndpointAndReturnError() bool {
	return f_.RespondsToSelector(objc.Sel("makeListenerEndpointAndReturnError:"))
}

// Returns an endpoint object that lets the host app communicate with the File Provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderservicesource/2915876-makelistenerendpointandreturnerr?language=objc
func (f_ FileProviderServiceSourceWrapper) MakeListenerEndpointAndReturnError(error foundation.IError) foundation.XPCListenerEndpoint {
	rv := objc.Call[foundation.XPCListenerEndpoint](f_, objc.Sel("makeListenerEndpointAndReturnError:"), objc.Ptr(error))
	return rv
}

func (f_ FileProviderServiceSourceWrapper) HasServiceName() bool {
	return f_.RespondsToSelector(objc.Sel("serviceName"))
}

// A name that uniquely identifies the service (reverse domain name notation is recommended). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderservicesource/2915879-servicename?language=objc
func (f_ FileProviderServiceSourceWrapper) ServiceName() foundation.FileProviderServiceName {
	rv := objc.Call[foundation.FileProviderServiceName](f_, objc.Sel("serviceName"))
	return rv
}
