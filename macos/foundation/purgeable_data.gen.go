// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PurgeableData] class.
var PurgeableDataClass = _PurgeableDataClass{objc.GetClass("NSPurgeableData")}

type _PurgeableDataClass struct {
	objc.Class
}

// An interface definition for the [PurgeableData] class.
type IPurgeableData interface {
	IMutableData
}

// A mutable data object containing bytes that can be discarded when they're no longer needed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspurgeabledata?language=objc
type PurgeableData struct {
	MutableData
}

func PurgeableDataFrom(ptr unsafe.Pointer) PurgeableData {
	return PurgeableData{
		MutableData: MutableDataFrom(ptr),
	}
}

func (pc _PurgeableDataClass) Alloc() PurgeableData {
	rv := objc.Call[PurgeableData](pc, objc.Sel("alloc"))
	return rv
}

func PurgeableData_Alloc() PurgeableData {
	return PurgeableDataClass.Alloc()
}

func (pc _PurgeableDataClass) New() PurgeableData {
	rv := objc.Call[PurgeableData](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPurgeableData() PurgeableData {
	return PurgeableDataClass.New()
}

func (p_ PurgeableData) Init() PurgeableData {
	rv := objc.Call[PurgeableData](p_, objc.Sel("init"))
	return rv
}

func (p_ PurgeableData) InitWithLength(length uint) PurgeableData {
	rv := objc.Call[PurgeableData](p_, objc.Sel("initWithLength:"), length)
	return rv
}

// Initializes and returns a mutable data object containing a given number of zeroed bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1413159-initwithlength?language=objc
func NewPurgeableDataWithLength(length uint) PurgeableData {
	instance := PurgeableDataClass.Alloc().InitWithLength(length)
	instance.Autorelease()
	return instance
}

func (pc _PurgeableDataClass) DataWithLength(length uint) PurgeableData {
	rv := objc.Call[PurgeableData](pc, objc.Sel("dataWithLength:"), length)
	return rv
}

// Creates and returns an mutable data object containing a given number of zeroed bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1547233-datawithlength?language=objc
func PurgeableData_DataWithLength(length uint) PurgeableData {
	return PurgeableDataClass.DataWithLength(length)
}

func (p_ PurgeableData) InitWithCapacity(capacity uint) PurgeableData {
	rv := objc.Call[PurgeableData](p_, objc.Sel("initWithCapacity:"), capacity)
	return rv
}

// Returns an initialized mutable data object capable of holding the specified number of bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1413350-initwithcapacity?language=objc
func NewPurgeableDataWithCapacity(capacity uint) PurgeableData {
	instance := PurgeableDataClass.Alloc().InitWithCapacity(capacity)
	instance.Autorelease()
	return instance
}

func (pc _PurgeableDataClass) DataWithCapacity(aNumItems uint) PurgeableData {
	rv := objc.Call[PurgeableData](pc, objc.Sel("dataWithCapacity:"), aNumItems)
	return rv
}

// Creates and returns a mutable data object capable of holding the specified number of bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1547236-datawithcapacity?language=objc
func PurgeableData_DataWithCapacity(aNumItems uint) PurgeableData {
	return PurgeableDataClass.DataWithCapacity(aNumItems)
}

func (p_ PurgeableData) InitWithContentsOfFileOptionsError(path string, readOptionsMask DataReadingOptions, errorPtr IError) PurgeableData {
	rv := objc.Call[PurgeableData](p_, objc.Sel("initWithContentsOfFile:options:error:"), path, readOptionsMask, objc.Ptr(errorPtr))
	return rv
}

// Initializes a data object with the content of the file at a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1411145-initwithcontentsoffile?language=objc
func NewPurgeableDataWithContentsOfFileOptionsError(path string, readOptionsMask DataReadingOptions, errorPtr IError) PurgeableData {
	instance := PurgeableDataClass.Alloc().InitWithContentsOfFileOptionsError(path, readOptionsMask, errorPtr)
	instance.Autorelease()
	return instance
}

func (p_ PurgeableData) CompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) PurgeableData {
	rv := objc.Call[PurgeableData](p_, objc.Sel("compressedDataUsingAlgorithm:error:"), algorithm, objc.Ptr(error))
	return rv
}

// Returns a new data object by compressing the data object’s bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/3174960-compresseddatausingalgorithm?language=objc
func PurgeableData_CompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) PurgeableData {
	instance := PurgeableDataClass.Alloc().CompressedDataUsingAlgorithmError(algorithm, error)
	instance.Autorelease()
	return instance
}

func (pc _PurgeableDataClass) DataWithContentsOfURL(url IURL) PurgeableData {
	rv := objc.Call[PurgeableData](pc, objc.Sel("dataWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Creates a data object containing the data from the location specified by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547245-datawithcontentsofurl?language=objc
func PurgeableData_DataWithContentsOfURL(url IURL) PurgeableData {
	return PurgeableDataClass.DataWithContentsOfURL(url)
}

func (p_ PurgeableData) InitWithData(data []byte) PurgeableData {
	rv := objc.Call[PurgeableData](p_, objc.Sel("initWithData:"), data)
	return rv
}

// Initializes a data object with the contents of another data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1417055-initwithdata?language=objc
func NewPurgeableDataWithData(data []byte) PurgeableData {
	instance := PurgeableDataClass.Alloc().InitWithData(data)
	instance.Autorelease()
	return instance
}

func (p_ PurgeableData) InitWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) PurgeableData {
	rv := objc.Call[PurgeableData](p_, objc.Sel("initWithBytesNoCopy:length:freeWhenDone:"), bytes, length, b)
	return rv
}

// Initializes a newly allocated data object by adding the given number of bytes from the given buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1416020-initwithbytesnocopy?language=objc
func NewPurgeableDataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) PurgeableData {
	instance := PurgeableDataClass.Alloc().InitWithBytesNoCopyLengthFreeWhenDone(bytes, length, b)
	instance.Autorelease()
	return instance
}

func (pc _PurgeableDataClass) DataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) PurgeableData {
	rv := objc.Call[PurgeableData](pc, objc.Sel("dataWithBytesNoCopy:length:freeWhenDone:"), bytes, length, b)
	return rv
}

// Creates a data object that holds a given number of bytes from a given buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547240-datawithbytesnocopy?language=objc
func PurgeableData_DataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) PurgeableData {
	return PurgeableDataClass.DataWithBytesNoCopyLengthFreeWhenDone(bytes, length, b)
}

func (pc _PurgeableDataClass) DataWithData(data []byte) PurgeableData {
	rv := objc.Call[PurgeableData](pc, objc.Sel("dataWithData:"), data)
	return rv
}

// Creates a data object containing the contents of another data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547230-datawithdata?language=objc
func PurgeableData_DataWithData(data []byte) PurgeableData {
	return PurgeableDataClass.DataWithData(data)
}

func (pc _PurgeableDataClass) DataWithContentsOfFileOptionsError(path string, readOptionsMask DataReadingOptions, errorPtr IError) PurgeableData {
	rv := objc.Call[PurgeableData](pc, objc.Sel("dataWithContentsOfFile:options:error:"), path, readOptionsMask, objc.Ptr(errorPtr))
	return rv
}

// Creates a data object by reading every byte from the file at a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547244-datawithcontentsoffile?language=objc
func PurgeableData_DataWithContentsOfFileOptionsError(path string, readOptionsMask DataReadingOptions, errorPtr IError) PurgeableData {
	return PurgeableDataClass.DataWithContentsOfFileOptionsError(path, readOptionsMask, errorPtr)
}

func (pc _PurgeableDataClass) Data() PurgeableData {
	rv := objc.Call[PurgeableData](pc, objc.Sel("data"))
	return rv
}

// Creates an empty data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547234-data?language=objc
func PurgeableData_Data() PurgeableData {
	return PurgeableDataClass.Data()
}

func (p_ PurgeableData) InitWithBytesLength(bytes unsafe.Pointer, length uint) PurgeableData {
	rv := objc.Call[PurgeableData](p_, objc.Sel("initWithBytes:length:"), bytes, length)
	return rv
}

// Initializes a data object filled with a given number of bytes copied from a given buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1412793-initwithbytes?language=objc
func NewPurgeableDataWithBytesLength(bytes unsafe.Pointer, length uint) PurgeableData {
	instance := PurgeableDataClass.Alloc().InitWithBytesLength(bytes, length)
	instance.Autorelease()
	return instance
}

func (p_ PurgeableData) InitWithContentsOfURL(url IURL) PurgeableData {
	rv := objc.Call[PurgeableData](p_, objc.Sel("initWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Initializes a data object with the data from the location specified by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1413892-initwithcontentsofurl?language=objc
func NewPurgeableDataWithContentsOfURL(url IURL) PurgeableData {
	instance := PurgeableDataClass.Alloc().InitWithContentsOfURL(url)
	instance.Autorelease()
	return instance
}

func (p_ PurgeableData) DecompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) PurgeableData {
	rv := objc.Call[PurgeableData](p_, objc.Sel("decompressedDataUsingAlgorithm:error:"), algorithm, objc.Ptr(error))
	return rv
}

// Returns a new data object by decompressing data object’s bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/3174961-decompresseddatausingalgorithm?language=objc
func PurgeableData_DecompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) PurgeableData {
	instance := PurgeableDataClass.Alloc().DecompressedDataUsingAlgorithmError(algorithm, error)
	instance.Autorelease()
	return instance
}

func (p_ PurgeableData) InitWithBase64EncodedDataOptions(base64Data []byte, options DataBase64DecodingOptions) PurgeableData {
	rv := objc.Call[PurgeableData](p_, objc.Sel("initWithBase64EncodedData:options:"), base64Data, options)
	return rv
}

// Initializes a data object with the given Base64 encoded data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1417833-initwithbase64encodeddata?language=objc
func NewPurgeableDataWithBase64EncodedDataOptions(base64Data []byte, options DataBase64DecodingOptions) PurgeableData {
	instance := PurgeableDataClass.Alloc().InitWithBase64EncodedDataOptions(base64Data, options)
	instance.Autorelease()
	return instance
}

func (p_ PurgeableData) InitWithBase64EncodedStringOptions(base64String string, options DataBase64DecodingOptions) PurgeableData {
	rv := objc.Call[PurgeableData](p_, objc.Sel("initWithBase64EncodedString:options:"), base64String, options)
	return rv
}

// Initializes a data object with the given Base64 encoded string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1410081-initwithbase64encodedstring?language=objc
func NewPurgeableDataWithBase64EncodedStringOptions(base64String string, options DataBase64DecodingOptions) PurgeableData {
	instance := PurgeableDataClass.Alloc().InitWithBase64EncodedStringOptions(base64String, options)
	instance.Autorelease()
	return instance
}

func (pc _PurgeableDataClass) DataWithBytesLength(bytes unsafe.Pointer, length uint) PurgeableData {
	rv := objc.Call[PurgeableData](pc, objc.Sel("dataWithBytes:length:"), bytes, length)
	return rv
}

// Creates a data object containing a given number of bytes copied from a given buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547231-datawithbytes?language=objc
func PurgeableData_DataWithBytesLength(bytes unsafe.Pointer, length uint) PurgeableData {
	return PurgeableDataClass.DataWithBytesLength(bytes, length)
}
