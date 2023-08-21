// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [String] class.
var StringClass = _StringClass{objc.GetClass("NSString")}

type _StringClass struct {
	objc.Class
}

// An interface definition for the [String] class.
type IString interface {
	objc.IObject
	IsEqualToString(aString string) bool
	HasPrefix(str string) bool
	MaximumLengthOfBytesUsingEncoding(enc StringEncoding) uint
	LocalizedStandardCompare(string_ string) ComparisonResult
	LocalizedStandardContainsString(str string) bool
	ComponentsSeparatedByCharactersInSet(separator ICharacterSet) []string
	RangeOfComposedCharacterSequenceAtIndex(index uint) Range
	LocalizedCaseInsensitiveContainsString(str string) bool
	CapitalizedStringWithLocale(locale ILocale) string
	SubstringFromIndex(from uint) string
	CommonPrefixWithStringOptions(str string, mask StringCompareOptions) string
	StringByReplacingCharactersInRangeWithString(range_ Range, replacement string) string
	LocalizedCaseInsensitiveCompare(string_ string) ComparisonResult
	StringByAppendingString(aString string) string
	RangeOfCharacterFromSetOptions(searchSet ICharacterSet, mask StringCompareOptions) Range
	SubstringWithRange(range_ Range) string
	StringByAddingPercentEncodingWithAllowedCharacters(allowedCharacters ICharacterSet) string
	VariantFittingPresentationWidth(width int) string
	LowercaseStringWithLocale(locale ILocale) string
	LineRangeForRange(range_ Range) Range
	SubstringToIndex(to uint) string
	CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes(outputName string, flag bool, outputArray []string, filterTypes []string) uint
	ParagraphRangeForRange(range_ Range) Range
	StringByAppendingFormat(format string, args ...any) string
	DataUsingEncoding(encoding StringEncoding) []byte
	SizeWithAttributes(attrs map[AttributedStringKey]objc.IObject) coregraphics.Size
	StringByReplacingOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options StringCompareOptions, searchRange Range) string
	StringByPaddingToLengthWithStringStartingAtIndex(newLength uint, padString string, padIndex uint) string
	PropertyList() objc.Object
	LengthOfBytesUsingEncoding(enc StringEncoding) uint
	StringByTrimmingCharactersInSet(set ICharacterSet) string
	RangeOfStringOptions(searchString string, mask StringCompareOptions) Range
	StringByAppendingPathExtension(str string) string
	GetLineStartEndContentsEndForRange(startPtr *uint, lineEndPtr *uint, contentsEndPtr *uint, range_ Range)
	GetFileSystemRepresentationMaxLength(cname *uint8, max uint) bool
	ContainsString(str string) bool
	CaseInsensitiveCompare(string_ string) ComparisonResult
	PropertyListFromStringsFileFormat() Dictionary
	ComponentsSeparatedByString(separator string) []string
	EnumerateLinesUsingBlock(block func(line string, stop *bool))
	LocalizedStandardRangeOfString(str string) Range
	GetBytesMaxLengthUsedLengthEncodingOptionsRangeRemainingRange(buffer unsafe.Pointer, maxBufferCount uint, usedBufferCount *uint, encoding StringEncoding, options StringEncodingConversionOptions, range_ Range, leftover RangePointer) bool
	UppercaseStringWithLocale(locale ILocale) string
	StringByAppendingPathExtensionForType(contentType objc.IObject) string
	CanBeConvertedToEncoding(encoding StringEncoding) bool
	CStringUsingEncoding(encoding StringEncoding) *uint8
	EnumerateSubstringsInRangeOptionsUsingBlock(range_ Range, opts StringEnumerationOptions, block func(substring string, substringRange Range, enclosingRange Range, stop *bool))
	StringByFoldingWithOptionsLocale(options StringCompareOptions, locale ILocale) string
	LocalizedCompare(string_ string) ComparisonResult
	StringsByAppendingPaths(paths []string) []string
	StringByAppendingPathComponentConformingToType(partialName string, contentType objc.IObject) string
	Compare(string_ string) ComparisonResult
	CharacterAtIndex(index uint) Unichar
	HasSuffix(str string) bool
	StringByApplyingTransformReverse(transform StringTransform, reverse bool) string
	GetParagraphStartEndContentsEndForRange(startPtr *uint, parEndPtr *uint, contentsEndPtr *uint, range_ Range)
	IntegerValue() int
	StringByDeletingPathExtension() string
	IntValue() int
	UppercaseString() string
	UTF8String() *uint8
	PrecomposedStringWithCanonicalMapping() string
	PathComponents() []string
	StringByExpandingTildeInPath() string
	IsAbsolutePath() bool
	CapitalizedString() string
	LongLongValue() int64
	PrecomposedStringWithCompatibilityMapping() string
	LocalizedLowercaseString() string
	LocalizedUppercaseString() string
	Description() string
	DoubleValue() float64
	SmallestEncoding() StringEncoding
	Length() uint
	StringByDeletingLastPathComponent() string
	Hash() uint
	DecomposedStringWithCanonicalMapping() string
	PathExtension() string
	StringByResolvingSymlinksInPath() string
	StringByAbbreviatingWithTildeInPath() string
	BoolValue() bool
	FastestEncoding() StringEncoding
	FloatValue() float64
	LastPathComponent() string
	DecomposedStringWithCompatibilityMapping() string
	LowercaseString() string
	FileSystemRepresentation() *uint8
	StringByStandardizingPath() string
	StringByRemovingPercentEncoding() string
	LocalizedCapitalizedString() string
}

// A static, plain-text Unicode string object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring?language=objc
type String struct {
	objc.Object
}

func StringFrom(ptr unsafe.Pointer) String {
	return String{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ String) InitWithCharactersNoCopyLengthDeallocator(chars *Unichar, len uint, deallocator func(arg0 *Unichar, arg1 uint)) String {
	rv := objc.Call[String](s_, objc.Sel("initWithCharactersNoCopy:length:deallocator:"), chars, len, deallocator)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/3547180-initwithcharactersnocopy?language=objc
func NewStringWithCharactersNoCopyLengthDeallocator(chars *Unichar, len uint, deallocator func(arg0 *Unichar, arg1 uint)) String {
	instance := StringClass.Alloc().InitWithCharactersNoCopyLengthDeallocator(chars, len, deallocator)
	instance.Autorelease()
	return instance
}

func (sc _StringClass) StringWithString(string_ string) String {
	rv := objc.Call[String](sc, objc.Sel("stringWithString:"), string_)
	return rv
}

// Returns a string created by copying the characters from another given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497372-stringwithstring?language=objc
func String_StringWithString(string_ string) String {
	return StringClass.StringWithString(string_)
}

func (s_ String) InitWithFormatLocale(format string, locale objc.IObject, args ...any) String {
	rv := objc.Call[String](s_, objc.Sel("initWithFormat:locale:"), append([]any{format, locale}, args...)...)
	return rv
}

// Returns an NSString object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497317-initwithformat?language=objc
func NewStringWithFormatLocale(format string, locale objc.IObject, args ...any) String {
	instance := StringClass.Alloc().InitWithFormatLocale(format, locale, args...)
	instance.Autorelease()
	return instance
}

func (s_ String) InitWithDataEncoding(data []byte, encoding StringEncoding) String {
	rv := objc.Call[String](s_, objc.Sel("initWithData:encoding:"), data, encoding)
	return rv
}

// Returns an NSString object initialized by converting given data into UTF-16 code units using a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416374-initwithdata?language=objc
func NewStringWithDataEncoding(data []byte, encoding StringEncoding) String {
	instance := StringClass.Alloc().InitWithDataEncoding(data, encoding)
	instance.Autorelease()
	return instance
}

func (sc _StringClass) StringWithFormat(format string, args ...any) String {
	rv := objc.Call[String](sc, objc.Sel("stringWithFormat:"), append([]any{format}, args...)...)
	return rv
}

// Returns a string created by using a given format string as a template into which the remaining argument values are substituted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497275-stringwithformat?language=objc
func String_StringWithFormat(format string, args ...any) String {
	return StringClass.StringWithFormat(format, args...)
}

func (s_ String) InitWithBytesNoCopyLengthEncodingDeallocator(bytes unsafe.Pointer, len uint, encoding StringEncoding, deallocator func(arg0 unsafe.Pointer, arg1 uint)) String {
	rv := objc.Call[String](s_, objc.Sel("initWithBytesNoCopy:length:encoding:deallocator:"), bytes, len, encoding, deallocator)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/3547179-initwithbytesnocopy?language=objc
func NewStringWithBytesNoCopyLengthEncodingDeallocator(bytes unsafe.Pointer, len uint, encoding StringEncoding, deallocator func(arg0 unsafe.Pointer, arg1 uint)) String {
	instance := StringClass.Alloc().InitWithBytesNoCopyLengthEncodingDeallocator(bytes, len, encoding, deallocator)
	instance.Autorelease()
	return instance
}

func (s_ String) InitWithString(aString string) String {
	rv := objc.Call[String](s_, objc.Sel("initWithString:"), aString)
	return rv
}

// Returns an NSString object initialized by copying the characters from another given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1411293-initwithstring?language=objc
func NewStringWithString(aString string) String {
	instance := StringClass.Alloc().InitWithString(aString)
	instance.Autorelease()
	return instance
}

func (s_ String) InitWithBytesLengthEncoding(bytes unsafe.Pointer, len uint, encoding StringEncoding) String {
	rv := objc.Call[String](s_, objc.Sel("initWithBytes:length:encoding:"), bytes, len, encoding)
	return rv
}

// Returns an initialized NSString object containing a given number of bytes from a given buffer of bytes interpreted in a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1407339-initwithbytes?language=objc
func NewStringWithBytesLengthEncoding(bytes unsafe.Pointer, len uint, encoding StringEncoding) String {
	instance := StringClass.Alloc().InitWithBytesLengthEncoding(bytes, len, encoding)
	instance.Autorelease()
	return instance
}

func (sc _StringClass) StringWithCharactersLength(characters *Unichar, length uint) String {
	rv := objc.Call[String](sc, objc.Sel("stringWithCharacters:length:"), characters, length)
	return rv
}

// Returns a string containing a given number of characters taken from a given C array of UTF-16 code units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497248-stringwithcharacters?language=objc
func String_StringWithCharactersLength(characters *Unichar, length uint) String {
	return StringClass.StringWithCharactersLength(characters, length)
}

func (s_ String) InitWithCharactersLength(characters *Unichar, length uint) String {
	rv := objc.Call[String](s_, objc.Sel("initWithCharacters:length:"), characters, length)
	return rv
}

// Returns an initialized NSString object that contains a given number of characters from a given C array of UTF-16 code units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1410997-initwithcharacters?language=objc
func NewStringWithCharactersLength(characters *Unichar, length uint) String {
	instance := StringClass.Alloc().InitWithCharactersLength(characters, length)
	instance.Autorelease()
	return instance
}

func (s_ String) InitWithUTF8String(nullTerminatedCString *uint8) String {
	rv := objc.Call[String](s_, objc.Sel("initWithUTF8String:"), nullTerminatedCString)
	return rv
}

// Returns an NSString object initialized by copying the characters from a given C array of UTF8-encoded bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1412128-initwithutf8string?language=objc
func NewStringWithUTF8String(nullTerminatedCString *uint8) String {
	instance := StringClass.Alloc().InitWithUTF8String(nullTerminatedCString)
	instance.Autorelease()
	return instance
}

func (sc _StringClass) LocalizedStringWithFormat(format string, args ...any) String {
	rv := objc.Call[String](sc, objc.Sel("localizedStringWithFormat:"), append([]any{format}, args...)...)
	return rv
}

// Returns a string created by using a given format string as a template into which the remaining argument values are substituted according to the current locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497301-localizedstringwithformat?language=objc
func String_LocalizedStringWithFormat(format string, args ...any) String {
	return StringClass.LocalizedStringWithFormat(format, args...)
}

func (s_ String) Init() String {
	rv := objc.Call[String](s_, objc.Sel("init"))
	return rv
}

func (sc _StringClass) StringWithUTF8String(nullTerminatedCString *uint8) String {
	rv := objc.Call[String](sc, objc.Sel("stringWithUTF8String:"), nullTerminatedCString)
	return rv
}

// Returns a string created by copying the data from a given C array of UTF8-encoded bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497379-stringwithutf8string?language=objc
func String_StringWithUTF8String(nullTerminatedCString *uint8) String {
	return StringClass.StringWithUTF8String(nullTerminatedCString)
}

func (sc _StringClass) StringWithContentsOfURLEncodingError(url IURL, enc StringEncoding, error IError) String {
	rv := objc.Call[String](sc, objc.Sel("stringWithContentsOfURL:encoding:error:"), objc.Ptr(url), enc, objc.Ptr(error))
	return rv
}

// Returns a string created by reading data from a given URL interpreted using a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497360-stringwithcontentsofurl?language=objc
func String_StringWithContentsOfURLEncodingError(url IURL, enc StringEncoding, error IError) String {
	return StringClass.StringWithContentsOfURLEncodingError(url, enc, error)
}

func (sc _StringClass) StringWithContentsOfFileEncodingError(path string, enc StringEncoding, error IError) String {
	rv := objc.Call[String](sc, objc.Sel("stringWithContentsOfFile:encoding:error:"), path, enc, objc.Ptr(error))
	return rv
}

// Returns a string created by reading data from the file at a given path interpreted using a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497327-stringwithcontentsoffile?language=objc
func String_StringWithContentsOfFileEncodingError(path string, enc StringEncoding, error IError) String {
	return StringClass.StringWithContentsOfFileEncodingError(path, enc, error)
}

func (sc _StringClass) Alloc() String {
	rv := objc.Call[String](sc, objc.Sel("alloc"))
	return rv
}

func String_Alloc() String {
	return StringClass.Alloc()
}

func (sc _StringClass) New() String {
	rv := objc.Call[String](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewString() String {
	return StringClass.New()
}

// Returns a Boolean value that indicates whether a given string is equal to the receiver using a literal Unicode-based comparison. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1407803-isequaltostring?language=objc
func (s_ String) IsEqualToString(aString string) bool {
	rv := objc.Call[bool](s_, objc.Sel("isEqualToString:"), aString)
	return rv
}

// Returns a Boolean value that indicates whether a given string matches the beginning characters of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1410309-hasprefix?language=objc
func (s_ String) HasPrefix(str string) bool {
	rv := objc.Call[bool](s_, objc.Sel("hasPrefix:"), str)
	return rv
}

// Returns the maximum number of bytes needed to store the receiver in a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1411611-maximumlengthofbytesusingencodin?language=objc
func (s_ String) MaximumLengthOfBytesUsingEncoding(enc StringEncoding) uint {
	rv := objc.Call[uint](s_, objc.Sel("maximumLengthOfBytesUsingEncoding:"), enc)
	return rv
}

// Compares strings as sorted by the Finder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1409742-localizedstandardcompare?language=objc
func (s_ String) LocalizedStandardCompare(string_ string) ComparisonResult {
	rv := objc.Call[ComparisonResult](s_, objc.Sel("localizedStandardCompare:"), string_)
	return rv
}

// Returns a Boolean value indicating whether the string contains a given string by performing a case and diacritic insensitive, locale-aware search. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416328-localizedstandardcontainsstring?language=objc
func (s_ String) LocalizedStandardContainsString(str string) bool {
	rv := objc.Call[bool](s_, objc.Sel("localizedStandardContainsString:"), str)
	return rv
}

// Returns an array containing substrings from the receiver that have been divided by characters in a given set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1410120-componentsseparatedbycharactersi?language=objc
func (s_ String) ComponentsSeparatedByCharactersInSet(separator ICharacterSet) []string {
	rv := objc.Call[[]string](s_, objc.Sel("componentsSeparatedByCharactersInSet:"), objc.Ptr(separator))
	return rv
}

// Returns the range in the receiver of the composed character sequence located at a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416036-rangeofcomposedcharactersequence?language=objc
func (s_ String) RangeOfComposedCharacterSequenceAtIndex(index uint) Range {
	rv := objc.Call[Range](s_, objc.Sel("rangeOfComposedCharacterSequenceAtIndex:"), index)
	return rv
}

// Returns a Boolean value indicating whether the string contains a given string by performing a case-insensitive, locale-aware search. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1412098-localizedcaseinsensitivecontains?language=objc
func (s_ String) LocalizedCaseInsensitiveContainsString(str string) bool {
	rv := objc.Call[bool](s_, objc.Sel("localizedCaseInsensitiveContainsString:"), str)
	return rv
}

// Returns a capitalized representation of the receiver using the specified locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1414023-capitalizedstringwithlocale?language=objc
func (s_ String) CapitalizedStringWithLocale(locale ILocale) string {
	rv := objc.Call[string](s_, objc.Sel("capitalizedStringWithLocale:"), objc.Ptr(locale))
	return rv
}

// Returns a new string containing the characters of the receiver from the one at a given index to the end. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1414368-substringfromindex?language=objc
func (s_ String) SubstringFromIndex(from uint) string {
	rv := objc.Call[string](s_, objc.Sel("substringFromIndex:"), from)
	return rv
}

// Returns a string containing characters the receiver and a given string have in common, starting from the beginning of each up to the first characters that aren’t equivalent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1408169-commonprefixwithstring?language=objc
func (s_ String) CommonPrefixWithStringOptions(str string, mask StringCompareOptions) string {
	rv := objc.Call[string](s_, objc.Sel("commonPrefixWithString:options:"), str, mask)
	return rv
}

// Returns a new string in which the characters in a specified range of the receiver are replaced by a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1410029-stringbyreplacingcharactersinran?language=objc
func (s_ String) StringByReplacingCharactersInRangeWithString(range_ Range, replacement string) string {
	rv := objc.Call[string](s_, objc.Sel("stringByReplacingCharactersInRange:withString:"), range_, replacement)
	return rv
}

// Compares the string with a given string using a case-insensitive, localized, comparison. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1417333-localizedcaseinsensitivecompare?language=objc
func (s_ String) LocalizedCaseInsensitiveCompare(string_ string) ComparisonResult {
	rv := objc.Call[ComparisonResult](s_, objc.Sel("localizedCaseInsensitiveCompare:"), string_)
	return rv
}

// Returns a new string made by appending a given string to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1412307-stringbyappendingstring?language=objc
func (s_ String) StringByAppendingString(aString string) string {
	rv := objc.Call[string](s_, objc.Sel("stringByAppendingString:"), aString)
	return rv
}

// Finds and returns the range in the string of the first character, using given options, from a given character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416898-rangeofcharacterfromset?language=objc
func (s_ String) RangeOfCharacterFromSetOptions(searchSet ICharacterSet, mask StringCompareOptions) Range {
	rv := objc.Call[Range](s_, objc.Sel("rangeOfCharacterFromSet:options:"), objc.Ptr(searchSet), mask)
	return rv
}

// Returns a string object containing the characters of the receiver that lie within a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1418469-substringwithrange?language=objc
func (s_ String) SubstringWithRange(range_ Range) string {
	rv := objc.Call[string](s_, objc.Sel("substringWithRange:"), range_)
	return rv
}

// Returns a new string made from the receiver by replacing all characters not in the specified set with percent-encoded characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1411946-stringbyaddingpercentencodingwit?language=objc
func (s_ String) StringByAddingPercentEncodingWithAllowedCharacters(allowedCharacters ICharacterSet) string {
	rv := objc.Call[string](s_, objc.Sel("stringByAddingPercentEncodingWithAllowedCharacters:"), objc.Ptr(allowedCharacters))
	return rv
}

// Returns a localized string intended for display in a notification alert. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1649585-localizedusernotificationstringf?language=objc
func (sc _StringClass) LocalizedUserNotificationStringForKeyArguments(key string, arguments []objc.IObject) string {
	rv := objc.Call[string](sc, objc.Sel("localizedUserNotificationStringForKey:arguments:"), key, arguments)
	return rv
}

// Returns a localized string intended for display in a notification alert. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1649585-localizedusernotificationstringf?language=objc
func String_LocalizedUserNotificationStringForKeyArguments(key string, arguments []objc.IObject) string {
	return StringClass.LocalizedUserNotificationStringForKeyArguments(key, arguments)
}

// Returns a string variation suitable for the specified presentation width. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1413104-variantfittingpresentationwidth?language=objc
func (s_ String) VariantFittingPresentationWidth(width int) string {
	rv := objc.Call[string](s_, objc.Sel("variantFittingPresentationWidth:"), width)
	return rv
}

// Returns a version of the string with all letters converted to lowercase, taking into account the specified locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1417298-lowercasestringwithlocale?language=objc
func (s_ String) LowercaseStringWithLocale(locale ILocale) string {
	rv := objc.Call[string](s_, objc.Sel("lowercaseStringWithLocale:"), objc.Ptr(locale))
	return rv
}

// Returns the range of characters representing the line or lines containing a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1407736-linerangeforrange?language=objc
func (s_ String) LineRangeForRange(range_ Range) Range {
	rv := objc.Call[Range](s_, objc.Sel("lineRangeForRange:"), range_)
	return rv
}

// Returns a new string containing the characters of the receiver up to, but not including, the one at a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1408017-substringtoindex?language=objc
func (s_ String) SubstringToIndex(to uint) string {
	rv := objc.Call[string](s_, objc.Sel("substringToIndex:"), to)
	return rv
}

// Interprets the receiver as a path in the file system and attempts to perform filename completion, returning a numeric value that indicates whether a match was possible, and by reference the longest path that matches the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1411841-completepathintostring?language=objc
func (s_ String) CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes(outputName string, flag bool, outputArray []string, filterTypes []string) uint {
	rv := objc.Call[uint](s_, objc.Sel("completePathIntoString:caseSensitive:matchesIntoArray:filterTypes:"), outputName, flag, outputArray, filterTypes)
	return rv
}

// Returns the range of characters representing the paragraph or paragraphs containing a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1408548-paragraphrangeforrange?language=objc
func (s_ String) ParagraphRangeForRange(range_ Range) Range {
	rv := objc.Call[Range](s_, objc.Sel("paragraphRangeForRange:"), range_)
	return rv
}

// Returns a string built from the strings in a given array by concatenating them with a path separator between each pair. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1417198-pathwithcomponents?language=objc
func (sc _StringClass) PathWithComponents(components []string) string {
	rv := objc.Call[string](sc, objc.Sel("pathWithComponents:"), components)
	return rv
}

// Returns a string built from the strings in a given array by concatenating them with a path separator between each pair. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1417198-pathwithcomponents?language=objc
func String_PathWithComponents(components []string) string {
	return StringClass.PathWithComponents(components)
}

// Returns a string made by appending to the receiver a string constructed from a given format string and the following arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497272-stringbyappendingformat?language=objc
func (s_ String) StringByAppendingFormat(format string, args ...any) string {
	rv := objc.Call[string](s_, objc.Sel("stringByAppendingFormat:"), append([]any{format}, args...)...)
	return rv
}

// Returns an NSData object containing a representation of the receiver encoded using a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416696-datausingencoding?language=objc
func (s_ String) DataUsingEncoding(encoding StringEncoding) []byte {
	rv := objc.Call[[]byte](s_, objc.Sel("dataUsingEncoding:"), encoding)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/2990401-deferredlocalizedintentsstringwi?language=objc
func (sc _StringClass) DeferredLocalizedIntentsStringWithFormat(format string, args ...any) string {
	rv := objc.Call[string](sc, objc.Sel("deferredLocalizedIntentsStringWithFormat:"), append([]any{format}, args...)...)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/2990401-deferredlocalizedintentsstringwi?language=objc
func String_DeferredLocalizedIntentsStringWithFormat(format string, args ...any) string {
	return StringClass.DeferredLocalizedIntentsStringWithFormat(format, args...)
}

// Returns the bounding box size the receiver occupies when drawn with the given attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1531844-sizewithattributes?language=objc
func (s_ String) SizeWithAttributes(attrs map[AttributedStringKey]objc.IObject) coregraphics.Size {
	rv := objc.Call[coregraphics.Size](s_, objc.Sel("sizeWithAttributes:"), attrs)
	return rv
}

// Returns a new string in which all occurrences of a target string in a specified range of the receiver are replaced by another given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416484-stringbyreplacingoccurrencesofst?language=objc
func (s_ String) StringByReplacingOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options StringCompareOptions, searchRange Range) string {
	rv := objc.Call[string](s_, objc.Sel("stringByReplacingOccurrencesOfString:withString:options:range:"), target, replacement, options, searchRange)
	return rv
}

// Returns the string encoding for the given data as detected by attempting to create a string according to the specified encoding options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1413576-stringencodingfordata?language=objc
func (sc _StringClass) StringEncodingForDataEncodingOptionsConvertedStringUsedLossyConversion(data []byte, opts map[StringEncodingDetectionOptionsKey]objc.IObject, string_ string, usedLossyConversion *bool) StringEncoding {
	rv := objc.Call[StringEncoding](sc, objc.Sel("stringEncodingForData:encodingOptions:convertedString:usedLossyConversion:"), data, opts, string_, usedLossyConversion)
	return rv
}

// Returns the string encoding for the given data as detected by attempting to create a string according to the specified encoding options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1413576-stringencodingfordata?language=objc
func String_StringEncodingForDataEncodingOptionsConvertedStringUsedLossyConversion(data []byte, opts map[StringEncodingDetectionOptionsKey]objc.IObject, string_ string, usedLossyConversion *bool) StringEncoding {
	return StringClass.StringEncodingForDataEncodingOptionsConvertedStringUsedLossyConversion(data, opts, string_, usedLossyConversion)
}

// Returns a new string formed from the receiver by either removing characters from the end, or by appending as many occurrences as necessary of a given pad string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416395-stringbypaddingtolength?language=objc
func (s_ String) StringByPaddingToLengthWithStringStartingAtIndex(newLength uint, padString string, padIndex uint) string {
	rv := objc.Call[string](s_, objc.Sel("stringByPaddingToLength:withString:startingAtIndex:"), newLength, padString, padIndex)
	return rv
}

// Parses the receiver as a text representation of a property list, returning an NSString, NSData, NSArray, or NSDictionary object, according to the topmost element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1413115-propertylist?language=objc
func (s_ String) PropertyList() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("propertyList"))
	return rv
}

// Returns the number of bytes required to store the receiver in a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1410710-lengthofbytesusingencoding?language=objc
func (s_ String) LengthOfBytesUsingEncoding(enc StringEncoding) uint {
	rv := objc.Call[uint](s_, objc.Sel("lengthOfBytesUsingEncoding:"), enc)
	return rv
}

// Returns a new string made by removing from both ends of the receiver characters contained in a given character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1415462-stringbytrimmingcharactersinset?language=objc
func (s_ String) StringByTrimmingCharactersInSet(set ICharacterSet) string {
	rv := objc.Call[string](s_, objc.Sel("stringByTrimmingCharactersInSet:"), objc.Ptr(set))
	return rv
}

// Finds and returns the range of the first occurrence of a given string within the string, subject to given options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416849-rangeofstring?language=objc
func (s_ String) RangeOfStringOptions(searchString string, mask StringCompareOptions) Range {
	rv := objc.Call[Range](s_, objc.Sel("rangeOfString:options:"), searchString, mask)
	return rv
}

// Returns a new string made by appending to the receiver an extension separator followed by a given extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1412501-stringbyappendingpathextension?language=objc
func (s_ String) StringByAppendingPathExtension(str string) string {
	rv := objc.Call[string](s_, objc.Sel("stringByAppendingPathExtension:"), str)
	return rv
}

// Returns by reference the beginning of the first line and the end of the last line touched by the given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1415111-getlinestart?language=objc
func (s_ String) GetLineStartEndContentsEndForRange(startPtr *uint, lineEndPtr *uint, contentsEndPtr *uint, range_ Range) {
	objc.Call[objc.Void](s_, objc.Sel("getLineStart:end:contentsEnd:forRange:"), startPtr, lineEndPtr, contentsEndPtr, range_)
}

// Interprets the receiver as a system-independent path and fills a buffer with a C-string in a format and encoding suitable for use with file-system calls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1410269-getfilesystemrepresentation?language=objc
func (s_ String) GetFileSystemRepresentationMaxLength(cname *uint8, max uint) bool {
	rv := objc.Call[bool](s_, objc.Sel("getFileSystemRepresentation:maxLength:"), cname, max)
	return rv
}

// Returns a Boolean value indicating whether the string contains a given string by performing a case-sensitive, locale-unaware search. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1414563-containsstring?language=objc
func (s_ String) ContainsString(str string) bool {
	rv := objc.Call[bool](s_, objc.Sel("containsString:"), str)
	return rv
}

// Returns the result of invoking compare: with NSCaseInsensitiveSearch as the only option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1414769-caseinsensitivecompare?language=objc
func (s_ String) CaseInsensitiveCompare(string_ string) ComparisonResult {
	rv := objc.Call[ComparisonResult](s_, objc.Sel("caseInsensitiveCompare:"), string_)
	return rv
}

// Returns a dictionary object initialized with the keys and values found in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1407697-propertylistfromstringsfileforma?language=objc
func (s_ String) PropertyListFromStringsFileFormat() Dictionary {
	rv := objc.Call[Dictionary](s_, objc.Sel("propertyListFromStringsFileFormat"))
	return rv
}

// Returns an array containing substrings from the receiver that have been divided by a given separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1413214-componentsseparatedbystring?language=objc
func (s_ String) ComponentsSeparatedByString(separator string) []string {
	rv := objc.Call[[]string](s_, objc.Sel("componentsSeparatedByString:"), separator)
	return rv
}

// Enumerates all the lines in the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1408459-enumeratelinesusingblock?language=objc
func (s_ String) EnumerateLinesUsingBlock(block func(line string, stop *bool)) {
	objc.Call[objc.Void](s_, objc.Sel("enumerateLinesUsingBlock:"), block)
}

// Finds and returns the range of the first occurrence of a given string within the string by performing a case and diacritic insensitive, locale-aware search. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1413574-localizedstandardrangeofstring?language=objc
func (s_ String) LocalizedStandardRangeOfString(str string) Range {
	rv := objc.Call[Range](s_, objc.Sel("localizedStandardRangeOfString:"), str)
	return rv
}

// Gets a given range of characters as bytes in a specified encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1413453-getbytes?language=objc
func (s_ String) GetBytesMaxLengthUsedLengthEncodingOptionsRangeRemainingRange(buffer unsafe.Pointer, maxBufferCount uint, usedBufferCount *uint, encoding StringEncoding, options StringEncodingConversionOptions, range_ Range, leftover RangePointer) bool {
	rv := objc.Call[bool](s_, objc.Sel("getBytes:maxLength:usedLength:encoding:options:range:remainingRange:"), buffer, maxBufferCount, usedBufferCount, encoding, options, range_, leftover)
	return rv
}

// Returns a version of the string with all letters converted to uppercase, taking into account the specified locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1413316-uppercasestringwithlocale?language=objc
func (s_ String) UppercaseStringWithLocale(locale ILocale) string {
	rv := objc.Call[string](s_, objc.Sel("uppercaseStringWithLocale:"), objc.Ptr(locale))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/3564809-stringbyappendingpathextensionfo?language=objc
func (s_ String) StringByAppendingPathExtensionForType(contentType objc.IObject) string {
	rv := objc.Call[string](s_, objc.Sel("stringByAppendingPathExtensionForType:"), objc.Ptr(contentType))
	return rv
}

// Returns a Boolean value that indicates whether the receiver can be converted to a given encoding without loss of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1409496-canbeconvertedtoencoding?language=objc
func (s_ String) CanBeConvertedToEncoding(encoding StringEncoding) bool {
	rv := objc.Call[bool](s_, objc.Sel("canBeConvertedToEncoding:"), encoding)
	return rv
}

// Returns a representation of the string as a C string using a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1408489-cstringusingencoding?language=objc
func (s_ String) CStringUsingEncoding(encoding StringEncoding) *uint8 {
	rv := objc.Call[*uint8](s_, objc.Sel("cStringUsingEncoding:"), encoding)
	return rv
}

// Enumerates the substrings of the specified type in the specified range of the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416774-enumeratesubstringsinrange?language=objc
func (s_ String) EnumerateSubstringsInRangeOptionsUsingBlock(range_ Range, opts StringEnumerationOptions, block func(substring string, substringRange Range, enclosingRange Range, stop *bool)) {
	objc.Call[objc.Void](s_, objc.Sel("enumerateSubstringsInRange:options:usingBlock:"), range_, opts, block)
}

// Creates a string suitable for comparison by removing the specified character distinctions from a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1413779-stringbyfoldingwithoptions?language=objc
func (s_ String) StringByFoldingWithOptionsLocale(options StringCompareOptions, locale ILocale) string {
	rv := objc.Call[string](s_, objc.Sel("stringByFoldingWithOptions:locale:"), options, objc.Ptr(locale))
	return rv
}

// Compares the string and a given string using a localized comparison. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416999-localizedcompare?language=objc
func (s_ String) LocalizedCompare(string_ string) ComparisonResult {
	rv := objc.Call[ComparisonResult](s_, objc.Sel("localizedCompare:"), string_)
	return rv
}

// Returns an array of strings made by separately appending to the receiver each string in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1415100-stringsbyappendingpaths?language=objc
func (s_ String) StringsByAppendingPaths(paths []string) []string {
	rv := objc.Call[[]string](s_, objc.Sel("stringsByAppendingPaths:"), paths)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/3564808-stringbyappendingpathcomponent?language=objc
func (s_ String) StringByAppendingPathComponentConformingToType(partialName string, contentType objc.IObject) string {
	rv := objc.Call[string](s_, objc.Sel("stringByAppendingPathComponent:conformingToType:"), partialName, objc.Ptr(contentType))
	return rv
}

// Returns the result of invoking compare: with no options and the receiver’s full extent as the range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1414082-compare?language=objc
func (s_ String) Compare(string_ string) ComparisonResult {
	rv := objc.Call[ComparisonResult](s_, objc.Sel("compare:"), string_)
	return rv
}

// Returns a human-readable string giving the name of a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1408318-localizednameofstringencoding?language=objc
func (sc _StringClass) LocalizedNameOfStringEncoding(encoding StringEncoding) string {
	rv := objc.Call[string](sc, objc.Sel("localizedNameOfStringEncoding:"), encoding)
	return rv
}

// Returns a human-readable string giving the name of a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1408318-localizednameofstringencoding?language=objc
func String_LocalizedNameOfStringEncoding(encoding StringEncoding) string {
	return StringClass.LocalizedNameOfStringEncoding(encoding)
}

// Returns the character at a given UTF-16 code unit index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1414645-characteratindex?language=objc
func (s_ String) CharacterAtIndex(index uint) Unichar {
	rv := objc.Call[Unichar](s_, objc.Sel("characterAtIndex:"), index)
	return rv
}

// Returns a Boolean value that indicates whether a given string matches the ending characters of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416529-hassuffix?language=objc
func (s_ String) HasSuffix(str string) bool {
	rv := objc.Call[bool](s_, objc.Sel("hasSuffix:"), str)
	return rv
}

// Returns a new string by applying a specified transform to the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1407787-stringbyapplyingtransform?language=objc
func (s_ String) StringByApplyingTransformReverse(transform StringTransform, reverse bool) string {
	rv := objc.Call[string](s_, objc.Sel("stringByApplyingTransform:reverse:"), transform, reverse)
	return rv
}

// Returns by reference the beginning of the first paragraph and the end of the last paragraph touched by the given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416407-getparagraphstart?language=objc
func (s_ String) GetParagraphStartEndContentsEndForRange(startPtr *uint, parEndPtr *uint, contentsEndPtr *uint, range_ Range) {
	objc.Call[objc.Void](s_, objc.Sel("getParagraphStart:end:contentsEnd:forRange:"), startPtr, parEndPtr, contentsEndPtr, range_)
}

// The NSInteger value of the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1410267-integervalue?language=objc
func (s_ String) IntegerValue() int {
	rv := objc.Call[int](s_, objc.Sel("integerValue"))
	return rv
}

// A new string made by deleting the extension (if any, and only the last) from the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1418214-stringbydeletingpathextension?language=objc
func (s_ String) StringByDeletingPathExtension() string {
	rv := objc.Call[string](s_, objc.Sel("stringByDeletingPathExtension"))
	return rv
}

// The integer value of the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1414988-intvalue?language=objc
func (s_ String) IntValue() int {
	rv := objc.Call[int](s_, objc.Sel("intValue"))
	return rv
}

// An uppercase representation of the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1409855-uppercasestring?language=objc
func (s_ String) UppercaseString() string {
	rv := objc.Call[string](s_, objc.Sel("uppercaseString"))
	return rv
}

// A null-terminated UTF8 representation of the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1411189-utf8string?language=objc
func (s_ String) UTF8String() *uint8 {
	rv := objc.Call[*uint8](s_, objc.Sel("UTF8String"))
	return rv
}

// A string made by normalizing the string’s contents using the Unicode Normalization Form C. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1412645-precomposedstringwithcanonicalma?language=objc
func (s_ String) PrecomposedStringWithCanonicalMapping() string {
	rv := objc.Call[string](s_, objc.Sel("precomposedStringWithCanonicalMapping"))
	return rv
}

// The file-system path components of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1414489-pathcomponents?language=objc
func (s_ String) PathComponents() []string {
	rv := objc.Call[[]string](s_, objc.Sel("pathComponents"))
	return rv
}

// A new string made by expanding the initial component of the receiver to its full path value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1407716-stringbyexpandingtildeinpath?language=objc
func (s_ String) StringByExpandingTildeInPath() string {
	rv := objc.Call[string](s_, objc.Sel("stringByExpandingTildeInPath"))
	return rv
}

// Returns the C-string encoding assumed for any method accepting a C string as an argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1410091-defaultcstringencoding?language=objc
func (sc _StringClass) DefaultCStringEncoding() StringEncoding {
	rv := objc.Call[StringEncoding](sc, objc.Sel("defaultCStringEncoding"))
	return rv
}

// Returns the C-string encoding assumed for any method accepting a C string as an argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1410091-defaultcstringencoding?language=objc
func String_DefaultCStringEncoding() StringEncoding {
	return StringClass.DefaultCStringEncoding()
}

// Returns a zero-terminated list of the encodings string objects support in the application’s environment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1417579-availablestringencodings?language=objc
func (sc _StringClass) AvailableStringEncodings() *StringEncoding {
	rv := objc.Call[*StringEncoding](sc, objc.Sel("availableStringEncodings"))
	return rv
}

// Returns a zero-terminated list of the encodings string objects support in the application’s environment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1417579-availablestringencodings?language=objc
func String_AvailableStringEncodings() *StringEncoding {
	return StringClass.AvailableStringEncodings()
}

// A Boolean value that indicates whether the receiver represents an absolute path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1409068-absolutepath?language=objc
func (s_ String) IsAbsolutePath() bool {
	rv := objc.Call[bool](s_, objc.Sel("isAbsolutePath"))
	return rv
}

// A capitalized representation of the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416784-capitalizedstring?language=objc
func (s_ String) CapitalizedString() string {
	rv := objc.Call[string](s_, objc.Sel("capitalizedString"))
	return rv
}

// The long long value of the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1417731-longlongvalue?language=objc
func (s_ String) LongLongValue() int64 {
	rv := objc.Call[int64](s_, objc.Sel("longLongValue"))
	return rv
}

// A string made by normalizing the receiver’s contents using the Unicode Normalization Form KC. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1412625-precomposedstringwithcompatibili?language=objc
func (s_ String) PrecomposedStringWithCompatibilityMapping() string {
	rv := objc.Call[string](s_, objc.Sel("precomposedStringWithCompatibilityMapping"))
	return rv
}

// Returns a version of the string with all letters converted to lowercase, taking into account the current locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1414125-localizedlowercasestring?language=objc
func (s_ String) LocalizedLowercaseString() string {
	rv := objc.Call[string](s_, objc.Sel("localizedLowercaseString"))
	return rv
}

// Returns a version of the string with all letters converted to uppercase, taking into account the current locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1413331-localizeduppercasestring?language=objc
func (s_ String) LocalizedUppercaseString() string {
	rv := objc.Call[string](s_, objc.Sel("localizedUppercaseString"))
	return rv
}

// This NSString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1410889-description?language=objc
func (s_ String) Description() string {
	rv := objc.Call[string](s_, objc.Sel("description"))
	return rv
}

// The floating-point value of the string as a double. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1414031-doublevalue?language=objc
func (s_ String) DoubleValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("doubleValue"))
	return rv
}

// The smallest encoding to which the receiver can be converted without loss of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1418037-smallestencoding?language=objc
func (s_ String) SmallestEncoding() StringEncoding {
	rv := objc.Call[StringEncoding](s_, objc.Sel("smallestEncoding"))
	return rv
}

// The number of UTF-16 code units in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1414212-length?language=objc
func (s_ String) Length() uint {
	rv := objc.Call[uint](s_, objc.Sel("length"))
	return rv
}

// A new string made by deleting the last path component from the receiver, along with any final path separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1411141-stringbydeletinglastpathcomponen?language=objc
func (s_ String) StringByDeletingLastPathComponent() string {
	rv := objc.Call[string](s_, objc.Sel("stringByDeletingLastPathComponent"))
	return rv
}

// An unsigned integer that can be used as a hash table address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1417245-hash?language=objc
func (s_ String) Hash() uint {
	rv := objc.Call[uint](s_, objc.Sel("hash"))
	return rv
}

// A string made by normalizing the string’s contents using the Unicode Normalization Form D. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1409474-decomposedstringwithcanonicalmap?language=objc
func (s_ String) DecomposedStringWithCanonicalMapping() string {
	rv := objc.Call[string](s_, objc.Sel("decomposedStringWithCanonicalMapping"))
	return rv
}

// The path extension, if any, of the string as interpreted as a path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1407801-pathextension?language=objc
func (s_ String) PathExtension() string {
	rv := objc.Call[string](s_, objc.Sel("pathExtension"))
	return rv
}

// A new string made from the receiver by resolving all symbolic links and standardizing path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1417783-stringbyresolvingsymlinksinpath?language=objc
func (s_ String) StringByResolvingSymlinksInPath() string {
	rv := objc.Call[string](s_, objc.Sel("stringByResolvingSymlinksInPath"))
	return rv
}

// A new string that replaces the current home directory portion of the current path with a tilde (~) character. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1407943-stringbyabbreviatingwithtildeinp?language=objc
func (s_ String) StringByAbbreviatingWithTildeInPath() string {
	rv := objc.Call[string](s_, objc.Sel("stringByAbbreviatingWithTildeInPath"))
	return rv
}

// The Boolean value of the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1409420-boolvalue?language=objc
func (s_ String) BoolValue() bool {
	rv := objc.Call[bool](s_, objc.Sel("boolValue"))
	return rv
}

// The fastest encoding to which the receiver may be converted without loss of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1409567-fastestencoding?language=objc
func (s_ String) FastestEncoding() StringEncoding {
	rv := objc.Call[StringEncoding](s_, objc.Sel("fastestEncoding"))
	return rv
}

// The floating-point value of the string as a float. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1412321-floatvalue?language=objc
func (s_ String) FloatValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("floatValue"))
	return rv
}

// The last path component of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416528-lastpathcomponent?language=objc
func (s_ String) LastPathComponent() string {
	rv := objc.Call[string](s_, objc.Sel("lastPathComponent"))
	return rv
}

// A string made by normalizing the receiver’s contents using the Unicode Normalization Form KD. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1415417-decomposedstringwithcompatibilit?language=objc
func (s_ String) DecomposedStringWithCompatibilityMapping() string {
	rv := objc.Call[string](s_, objc.Sel("decomposedStringWithCompatibilityMapping"))
	return rv
}

// A lowercase representation of the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1408467-lowercasestring?language=objc
func (s_ String) LowercaseString() string {
	rv := objc.Call[string](s_, objc.Sel("lowercaseString"))
	return rv
}

// A file system-specific representation of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1414559-filesystemrepresentation?language=objc
func (s_ String) FileSystemRepresentation() *uint8 {
	rv := objc.Call[*uint8](s_, objc.Sel("fileSystemRepresentation"))
	return rv
}

// A new string made by removing extraneous path components from the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1407194-stringbystandardizingpath?language=objc
func (s_ String) StringByStandardizingPath() string {
	rv := objc.Call[string](s_, objc.Sel("stringByStandardizingPath"))
	return rv
}

// Returns a new string made from the receiver by replacing all percent encoded sequences with the matching UTF-8 characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1409569-stringbyremovingpercentencoding?language=objc
func (s_ String) StringByRemovingPercentEncoding() string {
	rv := objc.Call[string](s_, objc.Sel("stringByRemovingPercentEncoding"))
	return rv
}

// Returns a capitalized representation of the receiver using the current locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1414885-localizedcapitalizedstring?language=objc
func (s_ String) LocalizedCapitalizedString() string {
	rv := objc.Call[string](s_, objc.Sel("localizedCapitalizedString"))
	return rv
}
