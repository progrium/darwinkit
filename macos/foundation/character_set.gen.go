// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CharacterSetClass = _CharacterSetClass{objc.GetClass("NSCharacterSet")}

type _CharacterSetClass struct {
	objc.Class
}

type ICharacterSet interface {
	objc.IObject
	CharacterIsMember(aCharacter unichar) bool
	IsSupersetOfSet(theOtherSet ICharacterSet) bool
	LongCharacterIsMember(theLongChar uint32) bool
	BitmapRepresentation() []byte
	InvertedSet() CharacterSet
}

type CharacterSet struct {
	objc.Object
}

func MakeCharacterSet(ptr unsafe.Pointer) CharacterSet {
	return CharacterSet{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CharacterSetClass) Alloc() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("alloc"))
	return rv
}

func CharacterSet_Alloc() CharacterSet {
	return CharacterSetClass.Alloc()
}

func (cc _CharacterSetClass) New() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCharacterSet() CharacterSet {
	return CharacterSetClass.New()
}

func CharacterSet_New() CharacterSet {
	return CharacterSetClass.New()
}

func (c_ CharacterSet) Init() CharacterSet {
	rv := objc.CallMethod[CharacterSet](c_, objc.GetSelector("init"))
	return rv
}

func CharacterSet_Init() CharacterSet {
	return CharacterSetClass.Alloc().Init()
}

func (cc _CharacterSetClass) CharacterSetWithCharactersInString(aString string) CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("characterSetWithCharactersInString:"), aString)
	return rv
}

func CharacterSet_CharacterSetWithCharactersInString(aString string) CharacterSet {
	return CharacterSetClass.CharacterSetWithCharactersInString(aString)
}

func (cc _CharacterSetClass) CharacterSetWithRange(aRange Range) CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("characterSetWithRange:"), aRange)
	return rv
}

func CharacterSet_CharacterSetWithRange(aRange Range) CharacterSet {
	return CharacterSetClass.CharacterSetWithRange(aRange)
}

func (cc _CharacterSetClass) CharacterSetWithBitmapRepresentation(data []byte) CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("characterSetWithBitmapRepresentation:"), data)
	return rv
}

func CharacterSet_CharacterSetWithBitmapRepresentation(data []byte) CharacterSet {
	return CharacterSetClass.CharacterSetWithBitmapRepresentation(data)
}

func (cc _CharacterSetClass) CharacterSetWithContentsOfFile(fName string) CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("characterSetWithContentsOfFile:"), fName)
	return rv
}

func CharacterSet_CharacterSetWithContentsOfFile(fName string) CharacterSet {
	return CharacterSetClass.CharacterSetWithContentsOfFile(fName)
}

func (c_ CharacterSet) CharacterIsMember(aCharacter unichar) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("characterIsMember:"), aCharacter)
	return rv
}

func (c_ CharacterSet) IsSupersetOfSet(theOtherSet ICharacterSet) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isSupersetOfSet:"), objc.ExtractPtr(theOtherSet))
	return rv
}

func (c_ CharacterSet) LongCharacterIsMember(theLongChar uint32) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("longCharacterIsMember:"), theLongChar)
	return rv
}

func (cc _CharacterSetClass) AlphanumericCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("alphanumericCharacterSet"))
	return rv
}

func CharacterSet_AlphanumericCharacterSet() CharacterSet {
	return CharacterSetClass.AlphanumericCharacterSet()
}

func (cc _CharacterSetClass) CapitalizedLetterCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("capitalizedLetterCharacterSet"))
	return rv
}

func CharacterSet_CapitalizedLetterCharacterSet() CharacterSet {
	return CharacterSetClass.CapitalizedLetterCharacterSet()
}

func (cc _CharacterSetClass) ControlCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("controlCharacterSet"))
	return rv
}

func CharacterSet_ControlCharacterSet() CharacterSet {
	return CharacterSetClass.ControlCharacterSet()
}

func (cc _CharacterSetClass) DecimalDigitCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("decimalDigitCharacterSet"))
	return rv
}

func CharacterSet_DecimalDigitCharacterSet() CharacterSet {
	return CharacterSetClass.DecimalDigitCharacterSet()
}

func (cc _CharacterSetClass) DecomposableCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("decomposableCharacterSet"))
	return rv
}

func CharacterSet_DecomposableCharacterSet() CharacterSet {
	return CharacterSetClass.DecomposableCharacterSet()
}

func (cc _CharacterSetClass) IllegalCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("illegalCharacterSet"))
	return rv
}

func CharacterSet_IllegalCharacterSet() CharacterSet {
	return CharacterSetClass.IllegalCharacterSet()
}

func (cc _CharacterSetClass) LetterCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("letterCharacterSet"))
	return rv
}

func CharacterSet_LetterCharacterSet() CharacterSet {
	return CharacterSetClass.LetterCharacterSet()
}

func (cc _CharacterSetClass) LowercaseLetterCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("lowercaseLetterCharacterSet"))
	return rv
}

func CharacterSet_LowercaseLetterCharacterSet() CharacterSet {
	return CharacterSetClass.LowercaseLetterCharacterSet()
}

func (cc _CharacterSetClass) NewlineCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("newlineCharacterSet"))
	rv.Autorelease()
	return rv
}

func CharacterSet_NewlineCharacterSet() CharacterSet {
	return CharacterSetClass.NewlineCharacterSet()
}

func (cc _CharacterSetClass) NonBaseCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("nonBaseCharacterSet"))
	return rv
}

func CharacterSet_NonBaseCharacterSet() CharacterSet {
	return CharacterSetClass.NonBaseCharacterSet()
}

func (cc _CharacterSetClass) PunctuationCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("punctuationCharacterSet"))
	return rv
}

func CharacterSet_PunctuationCharacterSet() CharacterSet {
	return CharacterSetClass.PunctuationCharacterSet()
}

func (cc _CharacterSetClass) SymbolCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("symbolCharacterSet"))
	return rv
}

func CharacterSet_SymbolCharacterSet() CharacterSet {
	return CharacterSetClass.SymbolCharacterSet()
}

func (cc _CharacterSetClass) UppercaseLetterCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("uppercaseLetterCharacterSet"))
	return rv
}

func CharacterSet_UppercaseLetterCharacterSet() CharacterSet {
	return CharacterSetClass.UppercaseLetterCharacterSet()
}

func (cc _CharacterSetClass) WhitespaceAndNewlineCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("whitespaceAndNewlineCharacterSet"))
	return rv
}

func CharacterSet_WhitespaceAndNewlineCharacterSet() CharacterSet {
	return CharacterSetClass.WhitespaceAndNewlineCharacterSet()
}

func (cc _CharacterSetClass) WhitespaceCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("whitespaceCharacterSet"))
	return rv
}

func CharacterSet_WhitespaceCharacterSet() CharacterSet {
	return CharacterSetClass.WhitespaceCharacterSet()
}

func (cc _CharacterSetClass) URLFragmentAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("URLFragmentAllowedCharacterSet"))
	return rv
}

func CharacterSet_URLFragmentAllowedCharacterSet() CharacterSet {
	return CharacterSetClass.URLFragmentAllowedCharacterSet()
}

func (cc _CharacterSetClass) URLHostAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("URLHostAllowedCharacterSet"))
	return rv
}

func CharacterSet_URLHostAllowedCharacterSet() CharacterSet {
	return CharacterSetClass.URLHostAllowedCharacterSet()
}

func (cc _CharacterSetClass) URLPasswordAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("URLPasswordAllowedCharacterSet"))
	return rv
}

func CharacterSet_URLPasswordAllowedCharacterSet() CharacterSet {
	return CharacterSetClass.URLPasswordAllowedCharacterSet()
}

func (cc _CharacterSetClass) URLPathAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("URLPathAllowedCharacterSet"))
	return rv
}

func CharacterSet_URLPathAllowedCharacterSet() CharacterSet {
	return CharacterSetClass.URLPathAllowedCharacterSet()
}

func (cc _CharacterSetClass) URLQueryAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("URLQueryAllowedCharacterSet"))
	return rv
}

func CharacterSet_URLQueryAllowedCharacterSet() CharacterSet {
	return CharacterSetClass.URLQueryAllowedCharacterSet()
}

func (cc _CharacterSetClass) URLUserAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.GetSelector("URLUserAllowedCharacterSet"))
	return rv
}

func CharacterSet_URLUserAllowedCharacterSet() CharacterSet {
	return CharacterSetClass.URLUserAllowedCharacterSet()
}

func (c_ CharacterSet) BitmapRepresentation() []byte {
	rv := objc.CallMethod[[]byte](c_, objc.GetSelector("bitmapRepresentation"))
	return rv
}

func (c_ CharacterSet) InvertedSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](c_, objc.GetSelector("invertedSet"))
	return rv
}
