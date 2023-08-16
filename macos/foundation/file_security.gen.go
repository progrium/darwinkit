// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FileSecurity] class.
var FileSecurityClass = _FileSecurityClass{objc.GetClass("NSFileSecurity")}

type _FileSecurityClass struct {
	objc.Class
}

// An interface definition for the [FileSecurity] class.
type IFileSecurity interface {
	objc.IObject
}

// A stub class that encapsulates security information about a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilesecurity?language=objc
type FileSecurity struct {
	objc.Object
}

func FileSecurityFrom(ptr unsafe.Pointer) FileSecurity {
	return FileSecurity{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FileSecurityClass) Alloc() FileSecurity {
	rv := objc.Call[FileSecurity](fc, objc.Sel("alloc"))
	return rv
}

func FileSecurity_Alloc() FileSecurity {
	return FileSecurityClass.Alloc()
}

func (fc _FileSecurityClass) New() FileSecurity {
	rv := objc.Call[FileSecurity](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFileSecurity() FileSecurity {
	return FileSecurityClass.New()
}

func (f_ FileSecurity) Init() FileSecurity {
	rv := objc.Call[FileSecurity](f_, objc.Sel("init"))
	return rv
}
