// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Archiver] class.
var ArchiverClass = _ArchiverClass{objc.GetClass("NSArchiver")}

type _ArchiverClass struct {
	objc.Class
}

// An interface definition for the [Archiver] class.
type IArchiver interface {
	ICoder
}

// A coder that stores an object's data to an archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarchiver?language=objc
type Archiver struct {
	Coder
}

func ArchiverFrom(ptr unsafe.Pointer) Archiver {
	return Archiver{
		Coder: CoderFrom(ptr),
	}
}

func (ac _ArchiverClass) Alloc() Archiver {
	rv := objc.Call[Archiver](ac, objc.Sel("alloc"))
	return rv
}

func Archiver_Alloc() Archiver {
	return ArchiverClass.Alloc()
}

func (ac _ArchiverClass) New() Archiver {
	rv := objc.Call[Archiver](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewArchiver() Archiver {
	return ArchiverClass.New()
}

func (a_ Archiver) Init() Archiver {
	rv := objc.Call[Archiver](a_, objc.Sel("init"))
	return rv
}
