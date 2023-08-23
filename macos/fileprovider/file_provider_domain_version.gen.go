// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FileProviderDomainVersion] class.
var FileProviderDomainVersionClass = _FileProviderDomainVersionClass{objc.GetClass("NSFileProviderDomainVersion")}

type _FileProviderDomainVersionClass struct {
	objc.Class
}

// An interface definition for the [FileProviderDomainVersion] class.
type IFileProviderDomainVersion interface {
	objc.IObject
	Next() FileProviderDomainVersion
	Compare(otherVersion IFileProviderDomainVersion) foundation.ComparisonResult
}

// An opaque object that identifies a specific version of a domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomainversion?language=objc
type FileProviderDomainVersion struct {
	objc.Object
}

func FileProviderDomainVersionFrom(ptr unsafe.Pointer) FileProviderDomainVersion {
	return FileProviderDomainVersion{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FileProviderDomainVersionClass) Alloc() FileProviderDomainVersion {
	rv := objc.Call[FileProviderDomainVersion](fc, objc.Sel("alloc"))
	return rv
}

func FileProviderDomainVersion_Alloc() FileProviderDomainVersion {
	return FileProviderDomainVersionClass.Alloc()
}

func (fc _FileProviderDomainVersionClass) New() FileProviderDomainVersion {
	rv := objc.Call[FileProviderDomainVersion](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFileProviderDomainVersion() FileProviderDomainVersion {
	return FileProviderDomainVersionClass.New()
}

func (f_ FileProviderDomainVersion) Init() FileProviderDomainVersion {
	rv := objc.Call[FileProviderDomainVersion](f_, objc.Sel("init"))
	return rv
}

// Creates a new version that supersedes the current version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomainversion/3727804-next?language=objc
func (f_ FileProviderDomainVersion) Next() FileProviderDomainVersion {
	rv := objc.Call[FileProviderDomainVersion](f_, objc.Sel("next"))
	return rv
}

// Compares another domain version with this one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomainversion/3727803-compare?language=objc
func (f_ FileProviderDomainVersion) Compare(otherVersion IFileProviderDomainVersion) foundation.ComparisonResult {
	rv := objc.Call[foundation.ComparisonResult](f_, objc.Sel("compare:"), objc.Ptr(otherVersion))
	return rv
}
