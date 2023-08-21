// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Data] class.
var DataClass = _DataClass{objc.GetClass("NSData")}

type _DataClass struct {
	objc.Class
}

// An interface definition for the [Data] class.
type IData interface {
	objc.IObject
	WriteToFileOptionsError(path string, writeOptionsMask DataWritingOptions, errorPtr IError) bool
	SubdataWithRange(range_ Range) []byte
	RangeOfDataOptionsRange(dataToFind []byte, mask DataSearchOptions, searchRange Range) Range
	IsEqualToData(other []byte) bool
	Base64EncodedStringWithOptions(options DataBase64EncodingOptions) string
	Base64EncodedDataWithOptions(options DataBase64EncodingOptions) []byte
	GetBytesLength(buffer unsafe.Pointer, length uint)
	WriteToURLAtomically(url IURL, atomically bool) bool
	EnumerateByteRangesUsingBlock(block func(bytes unsafe.Pointer, byteRange Range, stop *bool))
	Description() string
	Length() uint
	Bytes() unsafe.Pointer
}

// A static byte buffer in memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata?language=objc
type Data struct {
	objc.Object
}

func DataFrom(ptr unsafe.Pointer) Data {
	return Data{
		Object: objc.ObjectFrom(ptr),
	}
}

func (d_ Data) InitWithContentsOfFileOptionsError(path string, readOptionsMask DataReadingOptions, errorPtr IError) Data {
	rv := objc.Call[Data](d_, objc.Sel("initWithContentsOfFile:options:error:"), path, readOptionsMask, objc.Ptr(errorPtr))
	return rv
}

// Initializes a data object with the content of the file at a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1411145-initwithcontentsoffile?language=objc
func NewDataWithContentsOfFileOptionsError(path string, readOptionsMask DataReadingOptions, errorPtr IError) Data {
	instance := DataClass.Alloc().InitWithContentsOfFileOptionsError(path, readOptionsMask, errorPtr)
	instance.Autorelease()
	return instance
}

func (d_ Data) CompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) Data {
	rv := objc.Call[Data](d_, objc.Sel("compressedDataUsingAlgorithm:error:"), algorithm, objc.Ptr(error))
	return rv
}

// Returns a new data object by compressing the data object’s bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/3174960-compresseddatausingalgorithm?language=objc
func Data_CompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) Data {
	instance := DataClass.Alloc().CompressedDataUsingAlgorithmError(algorithm, error)
	instance.Autorelease()
	return instance
}

func (dc _DataClass) DataWithContentsOfURL(url IURL) Data {
	rv := objc.Call[Data](dc, objc.Sel("dataWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Creates a data object containing the data from the location specified by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547245-datawithcontentsofurl?language=objc
func Data_DataWithContentsOfURL(url IURL) Data {
	return DataClass.DataWithContentsOfURL(url)
}

func (d_ Data) InitWithData(data []byte) Data {
	rv := objc.Call[Data](d_, objc.Sel("initWithData:"), data)
	return rv
}

// Initializes a data object with the contents of another data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1417055-initwithdata?language=objc
func NewDataWithData(data []byte) Data {
	instance := DataClass.Alloc().InitWithData(data)
	instance.Autorelease()
	return instance
}

func (d_ Data) InitWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) Data {
	rv := objc.Call[Data](d_, objc.Sel("initWithBytesNoCopy:length:freeWhenDone:"), bytes, length, b)
	return rv
}

// Initializes a newly allocated data object by adding the given number of bytes from the given buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1416020-initwithbytesnocopy?language=objc
func NewDataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) Data {
	instance := DataClass.Alloc().InitWithBytesNoCopyLengthFreeWhenDone(bytes, length, b)
	instance.Autorelease()
	return instance
}

func (dc _DataClass) DataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) Data {
	rv := objc.Call[Data](dc, objc.Sel("dataWithBytesNoCopy:length:freeWhenDone:"), bytes, length, b)
	return rv
}

// Creates a data object that holds a given number of bytes from a given buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547240-datawithbytesnocopy?language=objc
func Data_DataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) Data {
	return DataClass.DataWithBytesNoCopyLengthFreeWhenDone(bytes, length, b)
}

func (dc _DataClass) DataWithData(data []byte) Data {
	rv := objc.Call[Data](dc, objc.Sel("dataWithData:"), data)
	return rv
}

// Creates a data object containing the contents of another data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547230-datawithdata?language=objc
func Data_DataWithData(data []byte) Data {
	return DataClass.DataWithData(data)
}

func (dc _DataClass) DataWithContentsOfFileOptionsError(path string, readOptionsMask DataReadingOptions, errorPtr IError) Data {
	rv := objc.Call[Data](dc, objc.Sel("dataWithContentsOfFile:options:error:"), path, readOptionsMask, objc.Ptr(errorPtr))
	return rv
}

// Creates a data object by reading every byte from the file at a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547244-datawithcontentsoffile?language=objc
func Data_DataWithContentsOfFileOptionsError(path string, readOptionsMask DataReadingOptions, errorPtr IError) Data {
	return DataClass.DataWithContentsOfFileOptionsError(path, readOptionsMask, errorPtr)
}

func (dc _DataClass) Data() Data {
	rv := objc.Call[Data](dc, objc.Sel("data"))
	return rv
}

// Creates an empty data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547234-data?language=objc
func Data_Data() Data {
	return DataClass.Data()
}

func (d_ Data) InitWithBytesLength(bytes unsafe.Pointer, length uint) Data {
	rv := objc.Call[Data](d_, objc.Sel("initWithBytes:length:"), bytes, length)
	return rv
}

// Initializes a data object filled with a given number of bytes copied from a given buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1412793-initwithbytes?language=objc
func NewDataWithBytesLength(bytes unsafe.Pointer, length uint) Data {
	instance := DataClass.Alloc().InitWithBytesLength(bytes, length)
	instance.Autorelease()
	return instance
}

func (d_ Data) InitWithContentsOfURL(url IURL) Data {
	rv := objc.Call[Data](d_, objc.Sel("initWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Initializes a data object with the data from the location specified by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1413892-initwithcontentsofurl?language=objc
func NewDataWithContentsOfURL(url IURL) Data {
	instance := DataClass.Alloc().InitWithContentsOfURL(url)
	instance.Autorelease()
	return instance
}

func (d_ Data) DecompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) Data {
	rv := objc.Call[Data](d_, objc.Sel("decompressedDataUsingAlgorithm:error:"), algorithm, objc.Ptr(error))
	return rv
}

// Returns a new data object by decompressing data object’s bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/3174961-decompresseddatausingalgorithm?language=objc
func Data_DecompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error IError) Data {
	instance := DataClass.Alloc().DecompressedDataUsingAlgorithmError(algorithm, error)
	instance.Autorelease()
	return instance
}

func (d_ Data) InitWithBase64EncodedDataOptions(base64Data []byte, options DataBase64DecodingOptions) Data {
	rv := objc.Call[Data](d_, objc.Sel("initWithBase64EncodedData:options:"), base64Data, options)
	return rv
}

// Initializes a data object with the given Base64 encoded data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1417833-initwithbase64encodeddata?language=objc
func NewDataWithBase64EncodedDataOptions(base64Data []byte, options DataBase64DecodingOptions) Data {
	instance := DataClass.Alloc().InitWithBase64EncodedDataOptions(base64Data, options)
	instance.Autorelease()
	return instance
}

func (d_ Data) InitWithBase64EncodedStringOptions(base64String string, options DataBase64DecodingOptions) Data {
	rv := objc.Call[Data](d_, objc.Sel("initWithBase64EncodedString:options:"), base64String, options)
	return rv
}

// Initializes a data object with the given Base64 encoded string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1410081-initwithbase64encodedstring?language=objc
func NewDataWithBase64EncodedStringOptions(base64String string, options DataBase64DecodingOptions) Data {
	instance := DataClass.Alloc().InitWithBase64EncodedStringOptions(base64String, options)
	instance.Autorelease()
	return instance
}

func (dc _DataClass) DataWithBytesLength(bytes unsafe.Pointer, length uint) Data {
	rv := objc.Call[Data](dc, objc.Sel("dataWithBytes:length:"), bytes, length)
	return rv
}

// Creates a data object containing a given number of bytes copied from a given buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1547231-datawithbytes?language=objc
func Data_DataWithBytesLength(bytes unsafe.Pointer, length uint) Data {
	return DataClass.DataWithBytesLength(bytes, length)
}

func (dc _DataClass) Alloc() Data {
	rv := objc.Call[Data](dc, objc.Sel("alloc"))
	return rv
}

func Data_Alloc() Data {
	return DataClass.Alloc()
}

func (dc _DataClass) New() Data {
	rv := objc.Call[Data](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewData() Data {
	return DataClass.New()
}

func (d_ Data) Init() Data {
	rv := objc.Call[Data](d_, objc.Sel("init"))
	return rv
}

// Writes the data object’s bytes to the file specified by a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1414800-writetofile?language=objc
func (d_ Data) WriteToFileOptionsError(path string, writeOptionsMask DataWritingOptions, errorPtr IError) bool {
	rv := objc.Call[bool](d_, objc.Sel("writeToFile:options:error:"), path, writeOptionsMask, objc.Ptr(errorPtr))
	return rv
}

// Returns a new data object containing the data object's bytes that fall within the limits specified by a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1410128-subdatawithrange?language=objc
func (d_ Data) SubdataWithRange(range_ Range) []byte {
	rv := objc.Call[[]byte](d_, objc.Sel("subdataWithRange:"), range_)
	return rv
}

// Finds and returns the range of the first occurrence of the given data, within the given range, subject to given options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1410391-rangeofdata?language=objc
func (d_ Data) RangeOfDataOptionsRange(dataToFind []byte, mask DataSearchOptions, searchRange Range) Range {
	rv := objc.Call[Range](d_, objc.Sel("rangeOfData:options:range:"), dataToFind, mask, searchRange)
	return rv
}

// Returns a Boolean value indicating whether this data object is the same as another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1409330-isequaltodata?language=objc
func (d_ Data) IsEqualToData(other []byte) bool {
	rv := objc.Call[bool](d_, objc.Sel("isEqualToData:"), other)
	return rv
}

// Creates a Base64 encoded string from the string using the given options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1413546-base64encodedstringwithoptions?language=objc
func (d_ Data) Base64EncodedStringWithOptions(options DataBase64EncodingOptions) string {
	rv := objc.Call[string](d_, objc.Sel("base64EncodedStringWithOptions:"), options)
	return rv
}

// Creates a Base64, UTF-8 encoded data object from the string using the given options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1412739-base64encodeddatawithoptions?language=objc
func (d_ Data) Base64EncodedDataWithOptions(options DataBase64EncodingOptions) []byte {
	rv := objc.Call[[]byte](d_, objc.Sel("base64EncodedDataWithOptions:"), options)
	return rv
}

// Copies a number of bytes from the start of the data object into a given buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1411450-getbytes?language=objc
func (d_ Data) GetBytesLength(buffer unsafe.Pointer, length uint) {
	objc.Call[objc.Void](d_, objc.Sel("getBytes:length:"), buffer, length)
}

// Writes the data object's bytes to the location specified by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1415134-writetourl?language=objc
func (d_ Data) WriteToURLAtomically(url IURL, atomically bool) bool {
	rv := objc.Call[bool](d_, objc.Sel("writeToURL:atomically:"), objc.Ptr(url), atomically)
	return rv
}

// Enumerates each range of bytes in the data object using a block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1408400-enumeratebyterangesusingblock?language=objc
func (d_ Data) EnumerateByteRangesUsingBlock(block func(bytes unsafe.Pointer, byteRange Range, stop *bool)) {
	objc.Call[objc.Void](d_, objc.Sel("enumerateByteRangesUsingBlock:"), block)
}

// A string that contains a hexadecimal representation of the data object’s contents in a property list format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1412579-description?language=objc
func (d_ Data) Description() string {
	rv := objc.Call[string](d_, objc.Sel("description"))
	return rv
}

// The number of bytes contained by the data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1416769-length?language=objc
func (d_ Data) Length() uint {
	rv := objc.Call[uint](d_, objc.Sel("length"))
	return rv
}

// A pointer to the data object's contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/1410616-bytes?language=objc
func (d_ Data) Bytes() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](d_, objc.Sel("bytes"))
	return rv
}
