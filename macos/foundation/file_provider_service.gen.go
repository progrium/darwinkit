// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FileProviderService] class.
var FileProviderServiceClass = _FileProviderServiceClass{objc.GetClass("NSFileProviderService")}

type _FileProviderServiceClass struct {
	objc.Class
}

// An interface definition for the [FileProviderService] class.
type IFileProviderService interface {
	objc.IObject
	GetFileProviderConnectionWithCompletionHandler(completionHandler func(connection XPCConnection, error Error))
	Name() FileProviderServiceName
}

// A service that provides a custom communication channel between your app and a File Provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileproviderservice?language=objc
type FileProviderService struct {
	objc.Object
}

func FileProviderServiceFrom(ptr unsafe.Pointer) FileProviderService {
	return FileProviderService{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FileProviderServiceClass) Alloc() FileProviderService {
	rv := objc.Call[FileProviderService](fc, objc.Sel("alloc"))
	return rv
}

func FileProviderService_Alloc() FileProviderService {
	return FileProviderServiceClass.Alloc()
}

func (fc _FileProviderServiceClass) New() FileProviderService {
	rv := objc.Call[FileProviderService](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFileProviderService() FileProviderService {
	return FileProviderServiceClass.New()
}

func (f_ FileProviderService) Init() FileProviderService {
	rv := objc.Call[FileProviderService](f_, objc.Sel("init"))
	return rv
}

// Asynchronously returns the service’s connection object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileproviderservice/2915257-getfileproviderconnectionwithcom?language=objc
func (f_ FileProviderService) GetFileProviderConnectionWithCompletionHandler(completionHandler func(connection XPCConnection, error Error)) {
	objc.Call[objc.Void](f_, objc.Sel("getFileProviderConnectionWithCompletionHandler:"), completionHandler)
}

// The File Provider service’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileproviderservice/2915258-name?language=objc
func (f_ FileProviderService) Name() FileProviderServiceName {
	rv := objc.Call[FileProviderServiceName](f_, objc.Sel("name"))
	return rv
}
