// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableCharacterSet] class.
var MutableCharacterSetClass = _MutableCharacterSetClass{objc.GetClass("NSMutableCharacterSet")}

type _MutableCharacterSetClass struct {
	objc.Class
}

// An interface definition for the [MutableCharacterSet] class.
type IMutableCharacterSet interface {
	ICharacterSet
	AddCharactersInRange(aRange Range)
	Invert()
	FormUnionWithCharacterSet(otherSet ICharacterSet)
	AddCharactersInString(aString string)
	RemoveCharactersInString(aString string)
	FormIntersectionWithCharacterSet(otherSet ICharacterSet)
	RemoveCharactersInRange(aRange Range)
}

// An object representing a mutable set of Unicode character values for use in search operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablecharacterset?language=objc
type MutableCharacterSet struct {
	CharacterSet
}

func MutableCharacterSetFrom(ptr unsafe.Pointer) MutableCharacterSet {
	return MutableCharacterSet{
		CharacterSet: CharacterSetFrom(ptr),
	}
}

func (mc _MutableCharacterSetClass) Alloc() MutableCharacterSet {
	rv := objc.Call[MutableCharacterSet](mc, objc.Sel("alloc"))
	return rv
}

func MutableCharacterSet_Alloc() MutableCharacterSet {
	return MutableCharacterSetClass.Alloc()
}

func (mc _MutableCharacterSetClass) New() MutableCharacterSet {
	rv := objc.Call[MutableCharacterSet](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableCharacterSet() MutableCharacterSet {
	return MutableCharacterSetClass.New()
}

func (m_ MutableCharacterSet) Init() MutableCharacterSet {
	rv := objc.Call[MutableCharacterSet](m_, objc.Sel("init"))
	return rv
}

// Adds to the receiver the characters whose Unicode values are in a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablecharacterset/1412225-addcharactersinrange?language=objc
func (m_ MutableCharacterSet) AddCharactersInRange(aRange Range) {
	objc.Call[objc.Void](m_, objc.Sel("addCharactersInRange:"), aRange)
}

// Replaces all the characters in the receiver with all the characters it didn’t previously contain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablecharacterset/1410977-invert?language=objc
func (m_ MutableCharacterSet) Invert() {
	objc.Call[objc.Void](m_, objc.Sel("invert"))
}

// Modifies the receiver so it contains all characters that exist in either the receiver or another set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablecharacterset/1408380-formunionwithcharacterset?language=objc
func (m_ MutableCharacterSet) FormUnionWithCharacterSet(otherSet ICharacterSet) {
	objc.Call[objc.Void](m_, objc.Sel("formUnionWithCharacterSet:"), objc.Ptr(otherSet))
}

// Adds to the receiver the characters in a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablecharacterset/1413999-addcharactersinstring?language=objc
func (m_ MutableCharacterSet) AddCharactersInString(aString string) {
	objc.Call[objc.Void](m_, objc.Sel("addCharactersInString:"), aString)
}

// Removes from the receiver the characters in a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablecharacterset/1414812-removecharactersinstring?language=objc
func (m_ MutableCharacterSet) RemoveCharactersInString(aString string) {
	objc.Call[objc.Void](m_, objc.Sel("removeCharactersInString:"), aString)
}

// Modifies the receiver so it contains only characters that exist in both the receiver and another set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablecharacterset/1412512-formintersectionwithcharacterset?language=objc
func (m_ MutableCharacterSet) FormIntersectionWithCharacterSet(otherSet ICharacterSet) {
	objc.Call[objc.Void](m_, objc.Sel("formIntersectionWithCharacterSet:"), objc.Ptr(otherSet))
}

// Removes from the receiver the characters whose Unicode values are in a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablecharacterset/1416987-removecharactersinrange?language=objc
func (m_ MutableCharacterSet) RemoveCharactersInRange(aRange Range) {
	objc.Call[objc.Void](m_, objc.Sel("removeCharactersInRange:"), aRange)
}
