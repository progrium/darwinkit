// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MediaDataStorage] class.
var MediaDataStorageClass = _MediaDataStorageClass{objc.GetClass("AVMediaDataStorage")}

type _MediaDataStorageClass struct {
	objc.Class
}

// An interface definition for the [MediaDataStorage] class.
type IMediaDataStorage interface {
	objc.IObject
	URL() foundation.URL
}

// An object that represents the media sample data storage file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediadatastorage?language=objc
type MediaDataStorage struct {
	objc.Object
}

func MediaDataStorageFrom(ptr unsafe.Pointer) MediaDataStorage {
	return MediaDataStorage{
		Object: objc.ObjectFrom(ptr),
	}
}

func (m_ MediaDataStorage) InitWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) MediaDataStorage {
	rv := objc.Call[MediaDataStorage](m_, objc.Sel("initWithURL:options:"), objc.Ptr(URL), options)
	return rv
}

// Creates a media data storage object associated with a file URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediadatastorage/1386008-initwithurl?language=objc
func NewMediaDataStorageWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) MediaDataStorage {
	instance := MediaDataStorageClass.Alloc().InitWithURLOptions(URL, options)
	instance.Autorelease()
	return instance
}

func (mc _MediaDataStorageClass) Alloc() MediaDataStorage {
	rv := objc.Call[MediaDataStorage](mc, objc.Sel("alloc"))
	return rv
}

func MediaDataStorage_Alloc() MediaDataStorage {
	return MediaDataStorageClass.Alloc()
}

func (mc _MediaDataStorageClass) New() MediaDataStorage {
	rv := objc.Call[MediaDataStorage](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMediaDataStorage() MediaDataStorage {
	return MediaDataStorageClass.New()
}

func (m_ MediaDataStorage) Init() MediaDataStorage {
	rv := objc.Call[MediaDataStorage](m_, objc.Sel("init"))
	return rv
}

// Returns the URL used to initialize the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediadatastorage/1385809-url?language=objc
func (m_ MediaDataStorage) URL() foundation.URL {
	rv := objc.Call[foundation.URL](m_, objc.Sel("URL"))
	return rv
}
