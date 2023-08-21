// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableString] class.
var MutableStringClass = _MutableStringClass{objc.GetClass("NSMutableString")}

type _MutableStringClass struct {
	objc.Class
}

// An interface definition for the [MutableString] class.
type IMutableString interface {
	IString
	AppendFormat(format string, args ...any)
	DeleteCharactersInRange(range_ Range)
	ApplyTransformReverseRangeUpdatedRange(transform StringTransform, reverse bool, range_ Range, resultingRange RangePointer) bool
	InitWithCapacity(capacity uint) MutableString
	InsertStringAtIndex(aString string, loc uint)
	AppendString(aString string)
	SetString(aString string)
	ReplaceCharactersInRangeWithString(range_ Range, aString string)
	ReplaceOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options StringCompareOptions, searchRange Range) uint
}

// A dynamic plain-text Unicode string object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablestring?language=objc
type MutableString struct {
	String
}

func MutableStringFrom(ptr unsafe.Pointer) MutableString {
	return MutableString{
		String: StringFrom(ptr),
	}
}

func (mc _MutableStringClass) Alloc() MutableString {
	rv := objc.Call[MutableString](mc, objc.Sel("alloc"))
	return rv
}

func MutableString_Alloc() MutableString {
	return MutableStringClass.Alloc()
}

func (mc _MutableStringClass) New() MutableString {
	rv := objc.Call[MutableString](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableString() MutableString {
	return MutableStringClass.New()
}

func (m_ MutableString) Init() MutableString {
	rv := objc.Call[MutableString](m_, objc.Sel("init"))
	return rv
}

func (m_ MutableString) InitWithCharactersNoCopyLengthDeallocator(chars *Unichar, len uint, deallocator func(arg0 *Unichar, arg1 uint)) MutableString {
	rv := objc.Call[MutableString](m_, objc.Sel("initWithCharactersNoCopy:length:deallocator:"), chars, len, deallocator)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/3547180-initwithcharactersnocopy?language=objc
func NewMutableStringWithCharactersNoCopyLengthDeallocator(chars *Unichar, len uint, deallocator func(arg0 *Unichar, arg1 uint)) MutableString {
	instance := MutableStringClass.Alloc().InitWithCharactersNoCopyLengthDeallocator(chars, len, deallocator)
	instance.Autorelease()
	return instance
}

func (mc _MutableStringClass) StringWithString(string_ string) MutableString {
	rv := objc.Call[MutableString](mc, objc.Sel("stringWithString:"), string_)
	return rv
}

// Returns a string created by copying the characters from another given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497372-stringwithstring?language=objc
func MutableString_StringWithString(string_ string) MutableString {
	return MutableStringClass.StringWithString(string_)
}

func (m_ MutableString) InitWithFormatLocale(format string, locale objc.IObject, args ...any) MutableString {
	rv := objc.Call[MutableString](m_, objc.Sel("initWithFormat:locale:"), append([]any{format, locale}, args...)...)
	return rv
}

// Returns an NSString object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497317-initwithformat?language=objc
func NewMutableStringWithFormatLocale(format string, locale objc.IObject, args ...any) MutableString {
	instance := MutableStringClass.Alloc().InitWithFormatLocale(format, locale, args...)
	instance.Autorelease()
	return instance
}

func (m_ MutableString) InitWithDataEncoding(data []byte, encoding StringEncoding) MutableString {
	rv := objc.Call[MutableString](m_, objc.Sel("initWithData:encoding:"), data, encoding)
	return rv
}

// Returns an NSString object initialized by converting given data into UTF-16 code units using a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416374-initwithdata?language=objc
func NewMutableStringWithDataEncoding(data []byte, encoding StringEncoding) MutableString {
	instance := MutableStringClass.Alloc().InitWithDataEncoding(data, encoding)
	instance.Autorelease()
	return instance
}

func (mc _MutableStringClass) StringWithFormat(format string, args ...any) MutableString {
	rv := objc.Call[MutableString](mc, objc.Sel("stringWithFormat:"), append([]any{format}, args...)...)
	return rv
}

// Returns a string created by using a given format string as a template into which the remaining argument values are substituted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497275-stringwithformat?language=objc
func MutableString_StringWithFormat(format string, args ...any) MutableString {
	return MutableStringClass.StringWithFormat(format, args...)
}

func (m_ MutableString) InitWithBytesNoCopyLengthEncodingDeallocator(bytes unsafe.Pointer, len uint, encoding StringEncoding, deallocator func(arg0 unsafe.Pointer, arg1 uint)) MutableString {
	rv := objc.Call[MutableString](m_, objc.Sel("initWithBytesNoCopy:length:encoding:deallocator:"), bytes, len, encoding, deallocator)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/3547179-initwithbytesnocopy?language=objc
func NewMutableStringWithBytesNoCopyLengthEncodingDeallocator(bytes unsafe.Pointer, len uint, encoding StringEncoding, deallocator func(arg0 unsafe.Pointer, arg1 uint)) MutableString {
	instance := MutableStringClass.Alloc().InitWithBytesNoCopyLengthEncodingDeallocator(bytes, len, encoding, deallocator)
	instance.Autorelease()
	return instance
}

func (m_ MutableString) InitWithString(aString string) MutableString {
	rv := objc.Call[MutableString](m_, objc.Sel("initWithString:"), aString)
	return rv
}

// Returns an NSString object initialized by copying the characters from another given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1411293-initwithstring?language=objc
func NewMutableStringWithString(aString string) MutableString {
	instance := MutableStringClass.Alloc().InitWithString(aString)
	instance.Autorelease()
	return instance
}

func (m_ MutableString) InitWithBytesLengthEncoding(bytes unsafe.Pointer, len uint, encoding StringEncoding) MutableString {
	rv := objc.Call[MutableString](m_, objc.Sel("initWithBytes:length:encoding:"), bytes, len, encoding)
	return rv
}

// Returns an initialized NSString object containing a given number of bytes from a given buffer of bytes interpreted in a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1407339-initwithbytes?language=objc
func NewMutableStringWithBytesLengthEncoding(bytes unsafe.Pointer, len uint, encoding StringEncoding) MutableString {
	instance := MutableStringClass.Alloc().InitWithBytesLengthEncoding(bytes, len, encoding)
	instance.Autorelease()
	return instance
}

func (mc _MutableStringClass) StringWithCharactersLength(characters *Unichar, length uint) MutableString {
	rv := objc.Call[MutableString](mc, objc.Sel("stringWithCharacters:length:"), characters, length)
	return rv
}

// Returns a string containing a given number of characters taken from a given C array of UTF-16 code units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497248-stringwithcharacters?language=objc
func MutableString_StringWithCharactersLength(characters *Unichar, length uint) MutableString {
	return MutableStringClass.StringWithCharactersLength(characters, length)
}

func (m_ MutableString) InitWithCharactersLength(characters *Unichar, length uint) MutableString {
	rv := objc.Call[MutableString](m_, objc.Sel("initWithCharacters:length:"), characters, length)
	return rv
}

// Returns an initialized NSString object that contains a given number of characters from a given C array of UTF-16 code units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1410997-initwithcharacters?language=objc
func NewMutableStringWithCharactersLength(characters *Unichar, length uint) MutableString {
	instance := MutableStringClass.Alloc().InitWithCharactersLength(characters, length)
	instance.Autorelease()
	return instance
}

func (m_ MutableString) InitWithUTF8String(nullTerminatedCString *uint8) MutableString {
	rv := objc.Call[MutableString](m_, objc.Sel("initWithUTF8String:"), nullTerminatedCString)
	return rv
}

// Returns an NSString object initialized by copying the characters from a given C array of UTF8-encoded bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1412128-initwithutf8string?language=objc
func NewMutableStringWithUTF8String(nullTerminatedCString *uint8) MutableString {
	instance := MutableStringClass.Alloc().InitWithUTF8String(nullTerminatedCString)
	instance.Autorelease()
	return instance
}

func (mc _MutableStringClass) LocalizedStringWithFormat(format string, args ...any) MutableString {
	rv := objc.Call[MutableString](mc, objc.Sel("localizedStringWithFormat:"), append([]any{format}, args...)...)
	return rv
}

// Returns a string created by using a given format string as a template into which the remaining argument values are substituted according to the current locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497301-localizedstringwithformat?language=objc
func MutableString_LocalizedStringWithFormat(format string, args ...any) MutableString {
	return MutableStringClass.LocalizedStringWithFormat(format, args...)
}

func (mc _MutableStringClass) StringWithUTF8String(nullTerminatedCString *uint8) MutableString {
	rv := objc.Call[MutableString](mc, objc.Sel("stringWithUTF8String:"), nullTerminatedCString)
	return rv
}

// Returns a string created by copying the data from a given C array of UTF8-encoded bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497379-stringwithutf8string?language=objc
func MutableString_StringWithUTF8String(nullTerminatedCString *uint8) MutableString {
	return MutableStringClass.StringWithUTF8String(nullTerminatedCString)
}

func (mc _MutableStringClass) StringWithContentsOfURLEncodingError(url IURL, enc StringEncoding, error IError) MutableString {
	rv := objc.Call[MutableString](mc, objc.Sel("stringWithContentsOfURL:encoding:error:"), objc.Ptr(url), enc, objc.Ptr(error))
	return rv
}

// Returns a string created by reading data from a given URL interpreted using a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497360-stringwithcontentsofurl?language=objc
func MutableString_StringWithContentsOfURLEncodingError(url IURL, enc StringEncoding, error IError) MutableString {
	return MutableStringClass.StringWithContentsOfURLEncodingError(url, enc, error)
}

func (mc _MutableStringClass) StringWithContentsOfFileEncodingError(path string, enc StringEncoding, error IError) MutableString {
	rv := objc.Call[MutableString](mc, objc.Sel("stringWithContentsOfFile:encoding:error:"), path, enc, objc.Ptr(error))
	return rv
}

// Returns a string created by reading data from the file at a given path interpreted using a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497327-stringwithcontentsoffile?language=objc
func MutableString_StringWithContentsOfFileEncodingError(path string, enc StringEncoding, error IError) MutableString {
	return MutableStringClass.StringWithContentsOfFileEncodingError(path, enc, error)
}

// Returns an empty NSMutableString object with initial storage for a given number of characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablestring/1497396-stringwithcapacity?language=objc
func (mc _MutableStringClass) StringWithCapacity(capacity uint) MutableString {
	rv := objc.Call[MutableString](mc, objc.Sel("stringWithCapacity:"), capacity)
	return rv
}

// Returns an empty NSMutableString object with initial storage for a given number of characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablestring/1497396-stringwithcapacity?language=objc
func MutableString_StringWithCapacity(capacity uint) MutableString {
	return MutableStringClass.StringWithCapacity(capacity)
}

// Adds a constructed string to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablestring/1497308-appendformat?language=objc
func (m_ MutableString) AppendFormat(format string, args ...any) {
	objc.Call[objc.Void](m_, objc.Sel("appendFormat:"), append([]any{format}, args...)...)
}

// Removes from the receiver the characters in a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablestring/1415003-deletecharactersinrange?language=objc
func (m_ MutableString) DeleteCharactersInRange(range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("deleteCharactersInRange:"), range_)
}

// Transliterates the receiver by applying a specified ICU string transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablestring/1415742-applytransform?language=objc
func (m_ MutableString) ApplyTransformReverseRangeUpdatedRange(transform StringTransform, reverse bool, range_ Range, resultingRange RangePointer) bool {
	rv := objc.Call[bool](m_, objc.Sel("applyTransform:reverse:range:updatedRange:"), transform, reverse, range_, resultingRange)
	return rv
}

// Returns an NSMutableString object initialized with initial storage for a given number of characters, [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablestring/1416610-initwithcapacity?language=objc
func (m_ MutableString) InitWithCapacity(capacity uint) MutableString {
	rv := objc.Call[MutableString](m_, objc.Sel("initWithCapacity:"), capacity)
	return rv
}

// Inserts into the receiver the characters of a given string at a given location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablestring/1410999-insertstring?language=objc
func (m_ MutableString) InsertStringAtIndex(aString string, loc uint) {
	objc.Call[objc.Void](m_, objc.Sel("insertString:atIndex:"), aString, loc)
}

// Adds to the end of the receiver the characters of a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablestring/1417883-appendstring?language=objc
func (m_ MutableString) AppendString(aString string) {
	objc.Call[objc.Void](m_, objc.Sel("appendString:"), aString)
}

// Replaces the characters of the receiver with those in a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablestring/1409483-setstring?language=objc
func (m_ MutableString) SetString(aString string) {
	objc.Call[objc.Void](m_, objc.Sel("setString:"), aString)
}

// Replaces the characters from aRange with those in aString. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablestring/1416524-replacecharactersinrange?language=objc
func (m_ MutableString) ReplaceCharactersInRangeWithString(range_ Range, aString string) {
	objc.Call[objc.Void](m_, objc.Sel("replaceCharactersInRange:withString:"), range_, aString)
}

// Replaces all occurrences of a given string in a given range with another given string, returning the number of replacements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablestring/1412453-replaceoccurrencesofstring?language=objc
func (m_ MutableString) ReplaceOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options StringCompareOptions, searchRange Range) uint {
	rv := objc.Call[uint](m_, objc.Sel("replaceOccurrencesOfString:withString:options:range:"), target, replacement, options, searchRange)
	return rv
}
