// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FileProviderRequest] class.
var FileProviderRequestClass = _FileProviderRequestClass{objc.GetClass("NSFileProviderRequest")}

type _FileProviderRequestClass struct {
	objc.Class
}

// An interface definition for the [FileProviderRequest] class.
type IFileProviderRequest interface {
	objc.IObject
	IsFileViewerRequest() bool
	IsSystemRequest() bool
	DomainVersion() FileProviderDomainVersion
	RequestingExecutable() foundation.URL
}

// An object that provides information about the application requesting data from the File Provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderrequest?language=objc
type FileProviderRequest struct {
	objc.Object
}

func FileProviderRequestFrom(ptr unsafe.Pointer) FileProviderRequest {
	return FileProviderRequest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FileProviderRequestClass) Alloc() FileProviderRequest {
	rv := objc.Call[FileProviderRequest](fc, objc.Sel("alloc"))
	return rv
}

func FileProviderRequest_Alloc() FileProviderRequest {
	return FileProviderRequestClass.Alloc()
}

func (fc _FileProviderRequestClass) New() FileProviderRequest {
	rv := objc.Call[FileProviderRequest](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFileProviderRequest() FileProviderRequest {
	return FileProviderRequestClass.New()
}

func (f_ FileProviderRequest) Init() FileProviderRequest {
	rv := objc.Call[FileProviderRequest](f_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the request came from Finder or related system file browsers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderrequest/3656553-isfileviewerrequest?language=objc
func (f_ FileProviderRequest) IsFileViewerRequest() bool {
	rv := objc.Call[bool](f_, objc.Sel("isFileViewerRequest"))
	return rv
}

// A Boolean value that indicates whether the request came from a system process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderrequest/3656554-issystemrequest?language=objc
func (f_ FileProviderRequest) IsSystemRequest() bool {
	rv := objc.Call[bool](f_, objc.Sel("isSystemRequest"))
	return rv
}

// The version of the domain for the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderrequest/3727822-domainversion?language=objc
func (f_ FileProviderRequest) DomainVersion() FileProviderDomainVersion {
	rv := objc.Call[FileProviderDomainVersion](f_, objc.Sel("domainVersion"))
	return rv
}

// The URL of the requesting executable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderrequest/3142339-requestingexecutable?language=objc
func (f_ FileProviderRequest) RequestingExecutable() foundation.URL {
	rv := objc.Call[foundation.URL](f_, objc.Sel("requestingExecutable"))
	return rv
}
