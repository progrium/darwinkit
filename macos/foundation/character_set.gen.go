// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CharacterSet] class.
var CharacterSetClass = _CharacterSetClass{objc.GetClass("NSCharacterSet")}

type _CharacterSetClass struct {
	objc.Class
}

// An interface definition for the [CharacterSet] class.
type ICharacterSet interface {
	objc.IObject
	HasMemberInPlane(thePlane uint8) bool
	CharacterIsMember(aCharacter Unichar) bool
	LongCharacterIsMember(theLongChar objc.IObject) bool
	IsSupersetOfSet(theOtherSet ICharacterSet) bool
	BitmapRepresentation() []byte
	InvertedSet() CharacterSet
}

// An object representing a fixed set of Unicode character values for use in search operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset?language=objc
type CharacterSet struct {
	objc.Object
}

func CharacterSetFrom(ptr unsafe.Pointer) CharacterSet {
	return CharacterSet{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CharacterSetClass) Alloc() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("alloc"))
	return rv
}

func CharacterSet_Alloc() CharacterSet {
	return CharacterSetClass.Alloc()
}

func (cc _CharacterSetClass) New() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCharacterSet() CharacterSet {
	return CharacterSetClass.New()
}

func (c_ CharacterSet) Init() CharacterSet {
	rv := objc.Call[CharacterSet](c_, objc.Sel("init"))
	return rv
}

// Returns a Boolean value that indicates whether the receiver has at least one member in a given character plane. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1412406-hasmemberinplane?language=objc
func (c_ CharacterSet) HasMemberInPlane(thePlane uint8) bool {
	rv := objc.Call[bool](c_, objc.Sel("hasMemberInPlane:"), thePlane)
	return rv
}

// Returns a Boolean value that indicates whether a given character is in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1407659-characterismember?language=objc
func (c_ CharacterSet) CharacterIsMember(aCharacter Unichar) bool {
	rv := objc.Call[bool](c_, objc.Sel("characterIsMember:"), aCharacter)
	return rv
}

// Returns a character set read from the bitmap representation stored in the file a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1418269-charactersetwithcontentsoffile?language=objc
func (cc _CharacterSetClass) CharacterSetWithContentsOfFile(fName string) CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("characterSetWithContentsOfFile:"), fName)
	return rv
}

// Returns a character set read from the bitmap representation stored in the file a given path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1418269-charactersetwithcontentsoffile?language=objc
func CharacterSet_CharacterSetWithContentsOfFile(fName string) CharacterSet {
	return CharacterSetClass.CharacterSetWithContentsOfFile(fName)
}

// Returns a Boolean value that indicates whether a given long character is a member of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1409757-longcharacterismember?language=objc
func (c_ CharacterSet) LongCharacterIsMember(theLongChar objc.IObject) bool {
	rv := objc.Call[bool](c_, objc.Sel("longCharacterIsMember:"), objc.Ptr(theLongChar))
	return rv
}

// Returns a character set containing characters with Unicode values in a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1414398-charactersetwithrange?language=objc
func (cc _CharacterSetClass) CharacterSetWithRange(aRange Range) CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("characterSetWithRange:"), aRange)
	return rv
}

// Returns a character set containing characters with Unicode values in a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1414398-charactersetwithrange?language=objc
func CharacterSet_CharacterSetWithRange(aRange Range) CharacterSet {
	return CharacterSetClass.CharacterSetWithRange(aRange)
}

// Returns a character set containing characters determined by a given bitmap representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1415042-charactersetwithbitmaprepresenta?language=objc
func (cc _CharacterSetClass) CharacterSetWithBitmapRepresentation(data []byte) CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("characterSetWithBitmapRepresentation:"), data)
	return rv
}

// Returns a character set containing characters determined by a given bitmap representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1415042-charactersetwithbitmaprepresenta?language=objc
func CharacterSet_CharacterSetWithBitmapRepresentation(data []byte) CharacterSet {
	return CharacterSetClass.CharacterSetWithBitmapRepresentation(data)
}

// Returns a character set containing the characters in a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1414061-charactersetwithcharactersinstri?language=objc
func (cc _CharacterSetClass) CharacterSetWithCharactersInString(aString string) CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("characterSetWithCharactersInString:"), aString)
	return rv
}

// Returns a character set containing the characters in a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1414061-charactersetwithcharactersinstri?language=objc
func CharacterSet_CharacterSetWithCharactersInString(aString string) CharacterSet {
	return CharacterSetClass.CharacterSetWithCharactersInString(aString)
}

// Returns a Boolean value that indicates whether the receiver is a superset of another given character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1415606-issupersetofset?language=objc
func (c_ CharacterSet) IsSupersetOfSet(theOtherSet ICharacterSet) bool {
	rv := objc.Call[bool](c_, objc.Sel("isSupersetOfSet:"), objc.Ptr(theOtherSet))
	return rv
}

// A character set containing the newline characters (U+000A ~ U+000D, U+0085, U+2028, and U+2029). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416730-newlinecharacterset?language=objc
func (cc _CharacterSetClass) NewlineCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("newlineCharacterSet"))
	rv.Autorelease()
	return rv
}

// A character set containing the newline characters (U+000A ~ U+000D, U+0085, U+2028, and U+2029). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416730-newlinecharacterset?language=objc
func CharacterSet_NewlineCharacterSet() CharacterSet {
	return CharacterSetClass.NewlineCharacterSet()
}

// An NSData object encoding the receiver in binary format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1417719-bitmaprepresentation?language=objc
func (c_ CharacterSet) BitmapRepresentation() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("bitmapRepresentation"))
	return rv
}

// A character set containing the characters in Unicode General Categories L*, M*, and N*. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1407466-alphanumericcharacterset?language=objc
func (cc _CharacterSetClass) AlphanumericCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("alphanumericCharacterSet"))
	return rv
}

// A character set containing the characters in Unicode General Categories L*, M*, and N*. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1407466-alphanumericcharacterset?language=objc
func CharacterSet_AlphanumericCharacterSet() CharacterSet {
	return CharacterSetClass.AlphanumericCharacterSet()
}

// A character set containing the characters in Unicode General Category Lt. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1414409-capitalizedlettercharacterset?language=objc
func (cc _CharacterSetClass) CapitalizedLetterCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("capitalizedLetterCharacterSet"))
	return rv
}

// A character set containing the characters in Unicode General Category Lt. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1414409-capitalizedlettercharacterset?language=objc
func CharacterSet_CapitalizedLetterCharacterSet() CharacterSet {
	return CharacterSetClass.CapitalizedLetterCharacterSet()
}

// A character set containing the characters in Unicode General Category Ll. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1417123-lowercaselettercharacterset?language=objc
func (cc _CharacterSetClass) LowercaseLetterCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("lowercaseLetterCharacterSet"))
	return rv
}

// A character set containing the characters in Unicode General Category Ll. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1417123-lowercaselettercharacterset?language=objc
func CharacterSet_LowercaseLetterCharacterSet() CharacterSet {
	return CharacterSetClass.LowercaseLetterCharacterSet()
}

// Returns the character set for characters allowed in a query URL component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416698-urlqueryallowedcharacterset?language=objc
func (cc _CharacterSetClass) URLQueryAllowedCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("URLQueryAllowedCharacterSet"))
	return rv
}

// Returns the character set for characters allowed in a query URL component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416698-urlqueryallowedcharacterset?language=objc
func CharacterSet_URLQueryAllowedCharacterSet() CharacterSet {
	return CharacterSetClass.URLQueryAllowedCharacterSet()
}

// A character set containing the characters in Unicode General Category L* & M*. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1408569-lettercharacterset?language=objc
func (cc _CharacterSetClass) LetterCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("letterCharacterSet"))
	return rv
}

// A character set containing the characters in Unicode General Category L* & M*. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1408569-lettercharacterset?language=objc
func CharacterSet_LetterCharacterSet() CharacterSet {
	return CharacterSetClass.LetterCharacterSet()
}

// A character set containing individual Unicode characters that can also be represented as composed character sequences (such as for letters with accents), by the definition of “standard decomposition” in version 3.2 of the Unicode character encoding standard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416868-decomposablecharacterset?language=objc
func (cc _CharacterSetClass) DecomposableCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("decomposableCharacterSet"))
	return rv
}

// A character set containing individual Unicode characters that can also be represented as composed character sequences (such as for letters with accents), by the definition of “standard decomposition” in version 3.2 of the Unicode character encoding standard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416868-decomposablecharacterset?language=objc
func CharacterSet_DecomposableCharacterSet() CharacterSet {
	return CharacterSetClass.DecomposableCharacterSet()
}

// A character set containing the characters in the category of Decimal Numbers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1408239-decimaldigitcharacterset?language=objc
func (cc _CharacterSetClass) DecimalDigitCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("decimalDigitCharacterSet"))
	return rv
}

// A character set containing the characters in the category of Decimal Numbers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1408239-decimaldigitcharacterset?language=objc
func CharacterSet_DecimalDigitCharacterSet() CharacterSet {
	return CharacterSetClass.DecimalDigitCharacterSet()
}

// Returns the character set for characters allowed in a fragment URL component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1412537-urlfragmentallowedcharacterset?language=objc
func (cc _CharacterSetClass) URLFragmentAllowedCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("URLFragmentAllowedCharacterSet"))
	return rv
}

// Returns the character set for characters allowed in a fragment URL component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1412537-urlfragmentallowedcharacterset?language=objc
func CharacterSet_URLFragmentAllowedCharacterSet() CharacterSet {
	return CharacterSetClass.URLFragmentAllowedCharacterSet()
}

// A character set containing the characters in Unicode General Category P*. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1411415-punctuationcharacterset?language=objc
func (cc _CharacterSetClass) PunctuationCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("punctuationCharacterSet"))
	return rv
}

// A character set containing the characters in Unicode General Category P*. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1411415-punctuationcharacterset?language=objc
func CharacterSet_PunctuationCharacterSet() CharacterSet {
	return CharacterSetClass.PunctuationCharacterSet()
}

// A character set containing the characters in Unicode General Category Zs and CHARACTER TABULATION (U+0009). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416393-whitespacecharacterset?language=objc
func (cc _CharacterSetClass) WhitespaceCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("whitespaceCharacterSet"))
	return rv
}

// A character set containing the characters in Unicode General Category Zs and CHARACTER TABULATION (U+0009). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416393-whitespacecharacterset?language=objc
func CharacterSet_WhitespaceCharacterSet() CharacterSet {
	return CharacterSetClass.WhitespaceCharacterSet()
}

// Returns the character set for characters allowed in a host URL subcomponent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416426-urlhostallowedcharacterset?language=objc
func (cc _CharacterSetClass) URLHostAllowedCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("URLHostAllowedCharacterSet"))
	return rv
}

// Returns the character set for characters allowed in a host URL subcomponent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416426-urlhostallowedcharacterset?language=objc
func CharacterSet_URLHostAllowedCharacterSet() CharacterSet {
	return CharacterSetClass.URLHostAllowedCharacterSet()
}

// A character set containing the characters in Unicode General Category Cc and Cf. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416371-controlcharacterset?language=objc
func (cc _CharacterSetClass) ControlCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("controlCharacterSet"))
	return rv
}

// A character set containing the characters in Unicode General Category Cc and Cf. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416371-controlcharacterset?language=objc
func CharacterSet_ControlCharacterSet() CharacterSet {
	return CharacterSetClass.ControlCharacterSet()
}

// Returns the character set for characters allowed in a path URL component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416804-urlpathallowedcharacterset?language=objc
func (cc _CharacterSetClass) URLPathAllowedCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("URLPathAllowedCharacterSet"))
	return rv
}

// Returns the character set for characters allowed in a path URL component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416804-urlpathallowedcharacterset?language=objc
func CharacterSet_URLPathAllowedCharacterSet() CharacterSet {
	return CharacterSetClass.URLPathAllowedCharacterSet()
}

// Returns the character set for characters allowed in a user URL subcomponent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1411851-urluserallowedcharacterset?language=objc
func (cc _CharacterSetClass) URLUserAllowedCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("URLUserAllowedCharacterSet"))
	return rv
}

// Returns the character set for characters allowed in a user URL subcomponent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1411851-urluserallowedcharacterset?language=objc
func CharacterSet_URLUserAllowedCharacterSet() CharacterSet {
	return CharacterSetClass.URLUserAllowedCharacterSet()
}

// A character set containing values in the category of Non-Characters or that have not yet been defined in version 3.2 of the Unicode standard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416321-illegalcharacterset?language=objc
func (cc _CharacterSetClass) IllegalCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("illegalCharacterSet"))
	return rv
}

// A character set containing values in the category of Non-Characters or that have not yet been defined in version 3.2 of the Unicode standard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1416321-illegalcharacterset?language=objc
func CharacterSet_IllegalCharacterSet() CharacterSet {
	return CharacterSetClass.IllegalCharacterSet()
}

// A character set containing characters in Unicode General Category Z*, U+000A ~ U+000D, and U+0085. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1413732-whitespaceandnewlinecharacterset?language=objc
func (cc _CharacterSetClass) WhitespaceAndNewlineCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("whitespaceAndNewlineCharacterSet"))
	return rv
}

// A character set containing characters in Unicode General Category Z*, U+000A ~ U+000D, and U+0085. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1413732-whitespaceandnewlinecharacterset?language=objc
func CharacterSet_WhitespaceAndNewlineCharacterSet() CharacterSet {
	return CharacterSetClass.WhitespaceAndNewlineCharacterSet()
}

// A character set containing the characters in Unicode General Category S*. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1410965-symbolcharacterset?language=objc
func (cc _CharacterSetClass) SymbolCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("symbolCharacterSet"))
	return rv
}

// A character set containing the characters in Unicode General Category S*. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1410965-symbolcharacterset?language=objc
func CharacterSet_SymbolCharacterSet() CharacterSet {
	return CharacterSetClass.SymbolCharacterSet()
}

// A character set containing the characters in Unicode General Category Lu and Lt. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1417569-uppercaselettercharacterset?language=objc
func (cc _CharacterSetClass) UppercaseLetterCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("uppercaseLetterCharacterSet"))
	return rv
}

// A character set containing the characters in Unicode General Category Lu and Lt. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1417569-uppercaselettercharacterset?language=objc
func CharacterSet_UppercaseLetterCharacterSet() CharacterSet {
	return CharacterSetClass.UppercaseLetterCharacterSet()
}

// Returns the character set for characters allowed in a password URL subcomponent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1417313-urlpasswordallowedcharacterset?language=objc
func (cc _CharacterSetClass) URLPasswordAllowedCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("URLPasswordAllowedCharacterSet"))
	return rv
}

// Returns the character set for characters allowed in a password URL subcomponent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1417313-urlpasswordallowedcharacterset?language=objc
func CharacterSet_URLPasswordAllowedCharacterSet() CharacterSet {
	return CharacterSetClass.URLPasswordAllowedCharacterSet()
}

// A character set containing only characters that don’t exist in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1414025-invertedset?language=objc
func (c_ CharacterSet) InvertedSet() CharacterSet {
	rv := objc.Call[CharacterSet](c_, objc.Sel("invertedSet"))
	return rv
}

// A character set containing the characters in Unicode General Category M*. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1407836-nonbasecharacterset?language=objc
func (cc _CharacterSetClass) NonBaseCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](cc, objc.Sel("nonBaseCharacterSet"))
	return rv
}

// A character set containing the characters in Unicode General Category M*. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/1407836-nonbasecharacterset?language=objc
func CharacterSet_NonBaseCharacterSet() CharacterSet {
	return CharacterSetClass.NonBaseCharacterSet()
}
