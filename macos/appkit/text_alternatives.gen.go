// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextAlternatives] class.
var TextAlternativesClass = _TextAlternativesClass{objc.GetClass("NSTextAlternatives")}

type _TextAlternativesClass struct {
	objc.Class
}

// An interface definition for the [TextAlternatives] class.
type ITextAlternatives interface {
	objc.IObject
	NoteSelectedAlternativeString(alternativeString string)
	PrimaryString() string
	AlternativeStrings() []string
}

// A list of alternative strings for a piece of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextalternatives?language=objc
type TextAlternatives struct {
	objc.Object
}

func TextAlternativesFrom(ptr unsafe.Pointer) TextAlternatives {
	return TextAlternatives{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextAlternatives) InitWithPrimaryStringAlternativeStrings(primaryString string, alternativeStrings []string) TextAlternatives {
	rv := objc.Call[TextAlternatives](t_, objc.Sel("initWithPrimaryString:alternativeStrings:"), primaryString, alternativeStrings)
	return rv
}

// Initializes an NSTextAlternatives instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextalternatives/1529445-initwithprimarystring?language=objc
func NewTextAlternativesWithPrimaryStringAlternativeStrings(primaryString string, alternativeStrings []string) TextAlternatives {
	instance := TextAlternativesClass.Alloc().InitWithPrimaryStringAlternativeStrings(primaryString, alternativeStrings)
	instance.Autorelease()
	return instance
}

func (tc _TextAlternativesClass) Alloc() TextAlternatives {
	rv := objc.Call[TextAlternatives](tc, objc.Sel("alloc"))
	return rv
}

func TextAlternatives_Alloc() TextAlternatives {
	return TextAlternativesClass.Alloc()
}

func (tc _TextAlternativesClass) New() TextAlternatives {
	rv := objc.Call[TextAlternatives](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextAlternatives() TextAlternatives {
	return TextAlternativesClass.New()
}

func (t_ TextAlternatives) Init() TextAlternatives {
	rv := objc.Call[TextAlternatives](t_, objc.Sel("init"))
	return rv
}

// Sent to the NSTextAlternatives object by the text view when the user chooses one of the alternative strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextalternatives/1525721-noteselectedalternativestring?language=objc
func (t_ TextAlternatives) NoteSelectedAlternativeString(alternativeString string) {
	objc.Call[objc.Void](t_, objc.Sel("noteSelectedAlternativeString:"), alternativeString)
}

// The text that was initially chosen as the input string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextalternatives/1526166-primarystring?language=objc
func (t_ TextAlternatives) PrimaryString() string {
	rv := objc.Call[string](t_, objc.Sel("primaryString"))
	return rv
}

// An array of alternative possible interpretations that the user might select. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextalternatives/1527585-alternativestrings?language=objc
func (t_ TextAlternatives) AlternativeStrings() []string {
	rv := objc.Call[[]string](t_, objc.Sel("alternativeStrings"))
	return rv
}
