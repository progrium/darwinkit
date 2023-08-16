// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LinguisticTagger] class.
var LinguisticTaggerClass = _LinguisticTaggerClass{objc.GetClass("NSLinguisticTagger")}

type _LinguisticTaggerClass struct {
	objc.Class
}

// An interface definition for the [LinguisticTagger] class.
type ILinguisticTagger interface {
	objc.IObject
}

// Analyze natural language text to tag part of speech and lexical class, identify names, perform lemmatization, and determine the language and script. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictagger?language=objc
type LinguisticTagger struct {
	objc.Object
}

func LinguisticTaggerFrom(ptr unsafe.Pointer) LinguisticTagger {
	return LinguisticTagger{
		Object: objc.ObjectFrom(ptr),
	}
}

func (lc _LinguisticTaggerClass) Alloc() LinguisticTagger {
	rv := objc.Call[LinguisticTagger](lc, objc.Sel("alloc"))
	return rv
}

func LinguisticTagger_Alloc() LinguisticTagger {
	return LinguisticTaggerClass.Alloc()
}

func (lc _LinguisticTaggerClass) New() LinguisticTagger {
	rv := objc.Call[LinguisticTagger](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLinguisticTagger() LinguisticTagger {
	return LinguisticTaggerClass.New()
}

func (l_ LinguisticTagger) Init() LinguisticTagger {
	rv := objc.Call[LinguisticTagger](l_, objc.Sel("init"))
	return rv
}
