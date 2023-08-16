// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FileAccessIntent] class.
var FileAccessIntentClass = _FileAccessIntentClass{objc.GetClass("NSFileAccessIntent")}

type _FileAccessIntentClass struct {
	objc.Class
}

// An interface definition for the [FileAccessIntent] class.
type IFileAccessIntent interface {
	objc.IObject
	URL() URL
}

// The details of a coordinated-read or coordinated-write operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileaccessintent?language=objc
type FileAccessIntent struct {
	objc.Object
}

func FileAccessIntentFrom(ptr unsafe.Pointer) FileAccessIntent {
	return FileAccessIntent{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FileAccessIntentClass) WritingIntentWithURLOptions(url IURL, options FileCoordinatorWritingOptions) FileAccessIntent {
	rv := objc.Call[FileAccessIntent](fc, objc.Sel("writingIntentWithURL:options:"), objc.Ptr(url), options)
	return rv
}

// Returns a file access intent object for writing to the given URL with the provided options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileaccessintent/1413014-writingintentwithurl?language=objc
func FileAccessIntent_WritingIntentWithURLOptions(url IURL, options FileCoordinatorWritingOptions) FileAccessIntent {
	return FileAccessIntentClass.WritingIntentWithURLOptions(url, options)
}

func (fc _FileAccessIntentClass) ReadingIntentWithURLOptions(url IURL, options FileCoordinatorReadingOptions) FileAccessIntent {
	rv := objc.Call[FileAccessIntent](fc, objc.Sel("readingIntentWithURL:options:"), objc.Ptr(url), options)
	return rv
}

// Returns a file access intent object for reading the given URL with the provided options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileaccessintent/1417829-readingintentwithurl?language=objc
func FileAccessIntent_ReadingIntentWithURLOptions(url IURL, options FileCoordinatorReadingOptions) FileAccessIntent {
	return FileAccessIntentClass.ReadingIntentWithURLOptions(url, options)
}

func (fc _FileAccessIntentClass) Alloc() FileAccessIntent {
	rv := objc.Call[FileAccessIntent](fc, objc.Sel("alloc"))
	return rv
}

func FileAccessIntent_Alloc() FileAccessIntent {
	return FileAccessIntentClass.Alloc()
}

func (fc _FileAccessIntentClass) New() FileAccessIntent {
	rv := objc.Call[FileAccessIntent](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFileAccessIntent() FileAccessIntent {
	return FileAccessIntentClass.New()
}

func (f_ FileAccessIntent) Init() FileAccessIntent {
	rv := objc.Call[FileAccessIntent](f_, objc.Sel("init"))
	return rv
}

// The current URL for the item managed by the file access intent instance. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileaccessintent/1411459-url?language=objc
func (f_ FileAccessIntent) URL() URL {
	rv := objc.Call[URL](f_, objc.Sel("URL"))
	return rv
}
