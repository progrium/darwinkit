// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableData] class.
var MutableDataClass = _MutableDataClass{objc.GetClass("NSMutableData")}

type _MutableDataClass struct {
	objc.Class
}

// An interface definition for the [MutableData] class.
type IMutableData interface {
	IData
	AppendBytesLength(bytes unsafe.Pointer, length uint)
	CompressUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) bool
	ReplaceBytesInRangeWithBytes(range_ Range, bytes unsafe.Pointer)
	SetData(data []byte)
	AppendData(other []byte)
	DecompressUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) bool
	IncreaseLengthBy(extraLength uint)
	ResetBytesInRange(range_ Range)
	MutableBytes() unsafe.Pointer
	SetLength(value uint)
}

// An object representing a dynamic byte buffer in memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata?language=objc
type MutableData struct {
	Data
}

func MutableDataFrom(ptr unsafe.Pointer) MutableData {
	return MutableData{
		Data: DataFrom(ptr),
	}
}

func (m_ MutableData) InitWithLength(length uint) MutableData {
	rv := objc.Call[MutableData](m_, objc.Sel("initWithLength:"), length)
	return rv
}

// Initializes and returns a mutable data object containing a given number of zeroed bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1413159-initwithlength?language=objc
func NewMutableDataWithLength(length uint) MutableData {
	instance := MutableDataClass.Alloc().InitWithLength(length)
	instance.Autorelease()
	return instance
}

func (mc _MutableDataClass) DataWithLength(length uint) MutableData {
	rv := objc.Call[MutableData](mc, objc.Sel("dataWithLength:"), length)
	return rv
}

// Creates and returns an mutable data object containing a given number of zeroed bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1547233-datawithlength?language=objc
func MutableData_DataWithLength(length uint) MutableData {
	return MutableDataClass.DataWithLength(length)
}

func (m_ MutableData) InitWithCapacity(capacity uint) MutableData {
	rv := objc.Call[MutableData](m_, objc.Sel("initWithCapacity:"), capacity)
	return rv
}

// Returns an initialized mutable data object capable of holding the specified number of bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1413350-initwithcapacity?language=objc
func NewMutableDataWithCapacity(capacity uint) MutableData {
	instance := MutableDataClass.Alloc().InitWithCapacity(capacity)
	instance.Autorelease()
	return instance
}

func (mc _MutableDataClass) DataWithCapacity(aNumItems uint) MutableData {
	rv := objc.Call[MutableData](mc, objc.Sel("dataWithCapacity:"), aNumItems)
	return rv
}

// Creates and returns a mutable data object capable of holding the specified number of bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1547236-datawithcapacity?language=objc
func MutableData_DataWithCapacity(aNumItems uint) MutableData {
	return MutableDataClass.DataWithCapacity(aNumItems)
}

func (mc _MutableDataClass) Alloc() MutableData {
	rv := objc.Call[MutableData](mc, objc.Sel("alloc"))
	return rv
}

func MutableData_Alloc() MutableData {
	return MutableDataClass.Alloc()
}

func (mc _MutableDataClass) New() MutableData {
	rv := objc.Call[MutableData](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableData() MutableData {
	return MutableDataClass.New()
}

func (m_ MutableData) Init() MutableData {
	rv := objc.Call[MutableData](m_, objc.Sel("init"))
	return rv
}

func (m_ MutableData) InitWithContentsOfFileOptionsError(path string, readOptionsMask DataReadingOptions, errorPtr IError) MutableData {
	rv := objc.Call[MutableData](m_, objc.Sel("initWithContentsOfFile:options:error:"), path, readOptionsMask, objc.Ptr(errorPtr))
	return rv
}

// Initializes a data object with the content of the file at a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1411145-initwithcontentsoffile?language=objc
func NewMutableDataWithContentsOfFileOptionsError(path string, readOptionsMask DataReadingOptions, errorPtr IError) MutableData {
	instance := MutableDataClass.Alloc().InitWithContentsOfFileOptionsError(path, readOptionsMask, errorPtr)
	instance.Autorelease()
	return instance
}

func (m_ MutableData) CompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) MutableData {
	rv := objc.Call[MutableData](m_, objc.Sel("compressedDataUsingAlgorithm:error:"), algorithm, objc.Ptr(error))
	return rv
}

// Returns a new data object by compressing the data object’s bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/3174960-compresseddatausingalgorithm?language=objc
func MutableData_CompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) MutableData {
	instance := MutableDataClass.Alloc().CompressedDataUsingAlgorithmError(algorithm, error)
	instance.Autorelease()
	return instance
}

func (mc _MutableDataClass) DataWithContentsOfURL(url IURL) MutableData {
	rv := objc.Call[MutableData](mc, objc.Sel("dataWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Creates a data object containing the data from the location specified by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547245-datawithcontentsofurl?language=objc
func MutableData_DataWithContentsOfURL(url IURL) MutableData {
	return MutableDataClass.DataWithContentsOfURL(url)
}

func (m_ MutableData) InitWithData(data []byte) MutableData {
	rv := objc.Call[MutableData](m_, objc.Sel("initWithData:"), data)
	return rv
}

// Initializes a data object with the contents of another data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1417055-initwithdata?language=objc
func NewMutableDataWithData(data []byte) MutableData {
	instance := MutableDataClass.Alloc().InitWithData(data)
	instance.Autorelease()
	return instance
}

func (m_ MutableData) InitWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) MutableData {
	rv := objc.Call[MutableData](m_, objc.Sel("initWithBytesNoCopy:length:freeWhenDone:"), bytes, length, b)
	return rv
}

// Initializes a newly allocated data object by adding the given number of bytes from the given buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1416020-initwithbytesnocopy?language=objc
func NewMutableDataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) MutableData {
	instance := MutableDataClass.Alloc().InitWithBytesNoCopyLengthFreeWhenDone(bytes, length, b)
	instance.Autorelease()
	return instance
}

func (mc _MutableDataClass) DataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) MutableData {
	rv := objc.Call[MutableData](mc, objc.Sel("dataWithBytesNoCopy:length:freeWhenDone:"), bytes, length, b)
	return rv
}

// Creates a data object that holds a given number of bytes from a given buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547240-datawithbytesnocopy?language=objc
func MutableData_DataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) MutableData {
	return MutableDataClass.DataWithBytesNoCopyLengthFreeWhenDone(bytes, length, b)
}

func (mc _MutableDataClass) DataWithData(data []byte) MutableData {
	rv := objc.Call[MutableData](mc, objc.Sel("dataWithData:"), data)
	return rv
}

// Creates a data object containing the contents of another data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547230-datawithdata?language=objc
func MutableData_DataWithData(data []byte) MutableData {
	return MutableDataClass.DataWithData(data)
}

func (mc _MutableDataClass) DataWithContentsOfFileOptionsError(path string, readOptionsMask DataReadingOptions, errorPtr IError) MutableData {
	rv := objc.Call[MutableData](mc, objc.Sel("dataWithContentsOfFile:options:error:"), path, readOptionsMask, objc.Ptr(errorPtr))
	return rv
}

// Creates a data object by reading every byte from the file at a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547244-datawithcontentsoffile?language=objc
func MutableData_DataWithContentsOfFileOptionsError(path string, readOptionsMask DataReadingOptions, errorPtr IError) MutableData {
	return MutableDataClass.DataWithContentsOfFileOptionsError(path, readOptionsMask, errorPtr)
}

func (mc _MutableDataClass) Data() MutableData {
	rv := objc.Call[MutableData](mc, objc.Sel("data"))
	return rv
}

// Creates an empty data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547234-data?language=objc
func MutableData_Data() MutableData {
	return MutableDataClass.Data()
}

func (m_ MutableData) InitWithBytesLength(bytes unsafe.Pointer, length uint) MutableData {
	rv := objc.Call[MutableData](m_, objc.Sel("initWithBytes:length:"), bytes, length)
	return rv
}

// Initializes a data object filled with a given number of bytes copied from a given buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1412793-initwithbytes?language=objc
func NewMutableDataWithBytesLength(bytes unsafe.Pointer, length uint) MutableData {
	instance := MutableDataClass.Alloc().InitWithBytesLength(bytes, length)
	instance.Autorelease()
	return instance
}

func (m_ MutableData) InitWithContentsOfURL(url IURL) MutableData {
	rv := objc.Call[MutableData](m_, objc.Sel("initWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Initializes a data object with the data from the location specified by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1413892-initwithcontentsofurl?language=objc
func NewMutableDataWithContentsOfURL(url IURL) MutableData {
	instance := MutableDataClass.Alloc().InitWithContentsOfURL(url)
	instance.Autorelease()
	return instance
}

func (m_ MutableData) DecompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) MutableData {
	rv := objc.Call[MutableData](m_, objc.Sel("decompressedDataUsingAlgorithm:error:"), algorithm, objc.Ptr(error))
	return rv
}

// Returns a new data object by decompressing data object’s bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/3174961-decompresseddatausingalgorithm?language=objc
func MutableData_DecompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) MutableData {
	instance := MutableDataClass.Alloc().DecompressedDataUsingAlgorithmError(algorithm, error)
	instance.Autorelease()
	return instance
}

func (m_ MutableData) InitWithBase64EncodedDataOptions(base64Data []byte, options DataBase64DecodingOptions) MutableData {
	rv := objc.Call[MutableData](m_, objc.Sel("initWithBase64EncodedData:options:"), base64Data, options)
	return rv
}

// Initializes a data object with the given Base64 encoded data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1417833-initwithbase64encodeddata?language=objc
func NewMutableDataWithBase64EncodedDataOptions(base64Data []byte, options DataBase64DecodingOptions) MutableData {
	instance := MutableDataClass.Alloc().InitWithBase64EncodedDataOptions(base64Data, options)
	instance.Autorelease()
	return instance
}

func (m_ MutableData) InitWithBase64EncodedStringOptions(base64String string, options DataBase64DecodingOptions) MutableData {
	rv := objc.Call[MutableData](m_, objc.Sel("initWithBase64EncodedString:options:"), base64String, options)
	return rv
}

// Initializes a data object with the given Base64 encoded string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1410081-initwithbase64encodedstring?language=objc
func NewMutableDataWithBase64EncodedStringOptions(base64String string, options DataBase64DecodingOptions) MutableData {
	instance := MutableDataClass.Alloc().InitWithBase64EncodedStringOptions(base64String, options)
	instance.Autorelease()
	return instance
}

func (mc _MutableDataClass) DataWithBytesLength(bytes unsafe.Pointer, length uint) MutableData {
	rv := objc.Call[MutableData](mc, objc.Sel("dataWithBytes:length:"), bytes, length)
	return rv
}

// Creates a data object containing a given number of bytes copied from a given buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547231-datawithbytes?language=objc
func MutableData_DataWithBytesLength(bytes unsafe.Pointer, length uint) MutableData {
	return MutableDataClass.DataWithBytesLength(bytes, length)
}

// Appends to the receiver a given number of bytes from a given buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1407704-appendbytes?language=objc
func (m_ MutableData) AppendBytesLength(bytes unsafe.Pointer, length uint) {
	objc.Call[objc.Void](m_, objc.Sel("appendBytes:length:"), bytes, length)
}

// Compresses the data object’s bytes using an algorithm that you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/3174967-compressusingalgorithm?language=objc
func (m_ MutableData) CompressUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("compressUsingAlgorithm:error:"), algorithm, objc.Ptr(error))
	return rv
}

// Replaces with a given set of bytes a given range within the contents of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1414281-replacebytesinrange?language=objc
func (m_ MutableData) ReplaceBytesInRangeWithBytes(range_ Range, bytes unsafe.Pointer) {
	objc.Call[objc.Void](m_, objc.Sel("replaceBytesInRange:withBytes:"), range_, bytes)
}

// Replaces the entire contents of the receiver with the contents of another data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1417012-setdata?language=objc
func (m_ MutableData) SetData(data []byte) {
	objc.Call[objc.Void](m_, objc.Sel("setData:"), data)
}

// Appends the content of another data object to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1410724-appenddata?language=objc
func (m_ MutableData) AppendData(other []byte) {
	objc.Call[objc.Void](m_, objc.Sel("appendData:"), other)
}

// Decompresses the data object’s bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/3174968-decompressusingalgorithm?language=objc
func (m_ MutableData) DecompressUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("decompressUsingAlgorithm:error:"), algorithm, objc.Ptr(error))
	return rv
}

// Increases the length of the receiver by a given number of bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1416186-increaselengthby?language=objc
func (m_ MutableData) IncreaseLengthBy(extraLength uint) {
	objc.Call[objc.Void](m_, objc.Sel("increaseLengthBy:"), extraLength)
}

// Replaces with zeroes the contents of the receiver in a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1415526-resetbytesinrange?language=objc
func (m_ MutableData) ResetBytesInRange(range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("resetBytesInRange:"), range_)
}

// A pointer to the data contained by the mutable data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1410770-mutablebytes?language=objc
func (m_ MutableData) MutableBytes() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](m_, objc.Sel("mutableBytes"))
	return rv
}

// The number of bytes contained in the mutable data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/1413333-length?language=objc
func (m_ MutableData) SetLength(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setLength:"), value)
}
