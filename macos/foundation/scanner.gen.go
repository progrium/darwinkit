// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Scanner] class.
var ScannerClass = _ScannerClass{objc.GetClass("NSScanner")}

type _ScannerClass struct {
	objc.Class
}

// An interface definition for the [Scanner] class.
type IScanner interface {
	objc.IObject
	ScanInteger(result *int) bool
	ScanCharactersFromSetIntoString(set ICharacterSet, result string) bool
	ScanUnsignedLongLong(result *int64) bool
	ScanDouble(result *float64) bool
	ScanDecimal(dcm *Decimal) bool
	ScanLongLong(result *int64) bool
	ScanFloat(result *float64) bool
	ScanUpToCharactersFromSetIntoString(set ICharacterSet, result string) bool
	ScanHexLongLong(result *int64) bool
	ScanHexInt(result *int) bool
	ScanHexDouble(result *float64) bool
	ScanStringIntoString(string_ string, result string) bool
	ScanHexFloat(result *float64) bool
	ScanUpToStringIntoString(string_ string, result string) bool
	ScanInt(result *int) bool
	CharactersToBeSkipped() CharacterSet
	SetCharactersToBeSkipped(value ICharacterSet)
	ScanLocation() uint
	SetScanLocation(value uint)
	Locale() objc.Object
	SetLocale(value objc.IObject)
	IsAtEnd() bool
	String() string
	CaseSensitive() bool
	SetCaseSensitive(value bool)
}

// A string parser that scans for substrings or characters in a character set, and for numeric values from decimal, hexadecimal, and floating-point representations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner?language=objc
type Scanner struct {
	objc.Object
}

func ScannerFrom(ptr unsafe.Pointer) Scanner {
	return Scanner{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ Scanner) InitWithString(string_ string) Scanner {
	rv := objc.Call[Scanner](s_, objc.Sel("initWithString:"), string_)
	return rv
}

// Returns an NSScanner object initialized to scan a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1417131-initwithstring?language=objc
func Scanner_InitWithString(string_ string) Scanner {
	return ScannerClass.Alloc().InitWithString(string_)
}

func (sc _ScannerClass) ScannerWithString(string_ string) Scanner {
	rv := objc.Call[Scanner](sc, objc.Sel("scannerWithString:"), string_)
	return rv
}

// Returns an NSScanner object that scans a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1420483-scannerwithstring?language=objc
func Scanner_ScannerWithString(string_ string) Scanner {
	return ScannerClass.ScannerWithString(string_)
}

func (sc _ScannerClass) Alloc() Scanner {
	rv := objc.Call[Scanner](sc, objc.Sel("alloc"))
	return rv
}

func Scanner_Alloc() Scanner {
	return ScannerClass.Alloc()
}

func (sc _ScannerClass) New() Scanner {
	rv := objc.Call[Scanner](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScanner() Scanner {
	return ScannerClass.New()
}

func (s_ Scanner) Init() Scanner {
	rv := objc.Call[Scanner](s_, objc.Sel("init"))
	return rv
}

// Scans for an NSInteger value from a decimal representation, returning a found value by reference [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1411082-scaninteger?language=objc
func (s_ Scanner) ScanInteger(result *int) bool {
	rv := objc.Call[bool](s_, objc.Sel("scanInteger:"), result)
	return rv
}

// Scans the string as long as characters from a given character set are encountered, accumulating characters into a string that’s returned by reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1417575-scancharactersfromset?language=objc
func (s_ Scanner) ScanCharactersFromSetIntoString(set ICharacterSet, result string) bool {
	rv := objc.Call[bool](s_, objc.Sel("scanCharactersFromSet:intoString:"), objc.Ptr(set), result)
	return rv
}

// Scans for an unsigned long long value from a decimal representation, returning a found value by reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1408559-scanunsignedlonglong?language=objc
func (s_ Scanner) ScanUnsignedLongLong(result *int64) bool {
	rv := objc.Call[bool](s_, objc.Sel("scanUnsignedLongLong:"), result)
	return rv
}

// Scans for a double value, returning a found value by reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1415922-scandouble?language=objc
func (s_ Scanner) ScanDouble(result *float64) bool {
	rv := objc.Call[bool](s_, objc.Sel("scanDouble:"), result)
	return rv
}

// Scans for an NSDecimal value, returning a found value by reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1408025-scandecimal?language=objc
func (s_ Scanner) ScanDecimal(dcm *Decimal) bool {
	rv := objc.Call[bool](s_, objc.Sel("scanDecimal:"), dcm)
	return rv
}

// Returns an NSScanner object that scans a given string according to the user’s default locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1415577-localizedscannerwithstring?language=objc
func (sc _ScannerClass) LocalizedScannerWithString(string_ string) objc.Object {
	rv := objc.Call[objc.Object](sc, objc.Sel("localizedScannerWithString:"), string_)
	return rv
}

// Returns an NSScanner object that scans a given string according to the user’s default locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1415577-localizedscannerwithstring?language=objc
func Scanner_LocalizedScannerWithString(string_ string) objc.Object {
	return ScannerClass.LocalizedScannerWithString(string_)
}

// Scans for a long long value from a decimal representation, returning a found value by reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1412629-scanlonglong?language=objc
func (s_ Scanner) ScanLongLong(result *int64) bool {
	rv := objc.Call[bool](s_, objc.Sel("scanLongLong:"), result)
	return rv
}

// Scans for a float value, returning a found value by reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1411677-scanfloat?language=objc
func (s_ Scanner) ScanFloat(result *float64) bool {
	rv := objc.Call[bool](s_, objc.Sel("scanFloat:"), result)
	return rv
}

// Scans the string until a character from a given character set is encountered, accumulating characters into a string that’s returned by reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1413598-scanuptocharactersfromset?language=objc
func (s_ Scanner) ScanUpToCharactersFromSetIntoString(set ICharacterSet, result string) bool {
	rv := objc.Call[bool](s_, objc.Sel("scanUpToCharactersFromSet:intoString:"), objc.Ptr(set), result)
	return rv
}

// Scans for a long long value from a hexadecimal representation, returning a found value by reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1412759-scanhexlonglong?language=objc
func (s_ Scanner) ScanHexLongLong(result *int64) bool {
	rv := objc.Call[bool](s_, objc.Sel("scanHexLongLong:"), result)
	return rv
}

// Scans for an unsigned value from a hexadecimal representation, returning a found value by reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1413857-scanhexint?language=objc
func (s_ Scanner) ScanHexInt(result *int) bool {
	rv := objc.Call[bool](s_, objc.Sel("scanHexInt:"), result)
	return rv
}

// Scans for a double value from a hexadecimal representation, returning a found value by reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1411780-scanhexdouble?language=objc
func (s_ Scanner) ScanHexDouble(result *float64) bool {
	rv := objc.Call[bool](s_, objc.Sel("scanHexDouble:"), result)
	return rv
}

// Scans a given string, returning an equivalent string object by reference if a match is found. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1407591-scanstring?language=objc
func (s_ Scanner) ScanStringIntoString(string_ string, result string) bool {
	rv := objc.Call[bool](s_, objc.Sel("scanString:intoString:"), string_, result)
	return rv
}

// Scans for a double value from a hexadecimal representation, returning a found value by reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1411236-scanhexfloat?language=objc
func (s_ Scanner) ScanHexFloat(result *float64) bool {
	rv := objc.Call[bool](s_, objc.Sel("scanHexFloat:"), result)
	return rv
}

// Scans the string until a given string is encountered, accumulating characters into a string that’s returned by reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1411971-scanuptostring?language=objc
func (s_ Scanner) ScanUpToStringIntoString(string_ string, result string) bool {
	rv := objc.Call[bool](s_, objc.Sel("scanUpToString:intoString:"), string_, result)
	return rv
}

// Scans for an int value from a decimal representation, returning a found value by reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1410914-scanint?language=objc
func (s_ Scanner) ScanInt(result *int) bool {
	rv := objc.Call[bool](s_, objc.Sel("scanInt:"), result)
	return rv
}

// Character set containing the characters the scanner ignores when looking for a scannable element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1410204-characterstobeskipped?language=objc
func (s_ Scanner) CharactersToBeSkipped() CharacterSet {
	rv := objc.Call[CharacterSet](s_, objc.Sel("charactersToBeSkipped"))
	return rv
}

// Character set containing the characters the scanner ignores when looking for a scannable element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1410204-characterstobeskipped?language=objc
func (s_ Scanner) SetCharactersToBeSkipped(value ICharacterSet) {
	objc.Call[objc.Void](s_, objc.Sel("setCharactersToBeSkipped:"), objc.Ptr(value))
}

// The character position at which the receiver will begin its next scanning operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1413294-scanlocation?language=objc
func (s_ Scanner) ScanLocation() uint {
	rv := objc.Call[uint](s_, objc.Sel("scanLocation"))
	return rv
}

// The character position at which the receiver will begin its next scanning operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1413294-scanlocation?language=objc
func (s_ Scanner) SetScanLocation(value uint) {
	objc.Call[objc.Void](s_, objc.Sel("setScanLocation:"), value)
}

// The locale to use when scanning. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1409531-locale?language=objc
func (s_ Scanner) Locale() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("locale"))
	return rv
}

// The locale to use when scanning. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1409531-locale?language=objc
func (s_ Scanner) SetLocale(value objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setLocale:"), value)
}

// Flag that indicates whether the receiver has exhausted all significant characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1412801-atend?language=objc
func (s_ Scanner) IsAtEnd() bool {
	rv := objc.Call[bool](s_, objc.Sel("isAtEnd"))
	return rv
}

// The string the scanner will scan. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1418109-string?language=objc
func (s_ Scanner) String() string {
	rv := objc.Call[string](s_, objc.Sel("string"))
	return rv
}

// Flag that indicates whether the receiver distinguishes case in the characters it scans. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1409488-casesensitive?language=objc
func (s_ Scanner) CaseSensitive() bool {
	rv := objc.Call[bool](s_, objc.Sel("caseSensitive"))
	return rv
}

// Flag that indicates whether the receiver distinguishes case in the characters it scans. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscanner/1409488-casesensitive?language=objc
func (s_ Scanner) SetCaseSensitive(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setCaseSensitive:"), value)
}
