// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MorphologyCustomPronoun] class.
var MorphologyCustomPronounClass = _MorphologyCustomPronounClass{objc.GetClass("NSMorphologyCustomPronoun")}

type _MorphologyCustomPronounClass struct {
	objc.Class
}

// An interface definition for the [MorphologyCustomPronoun] class.
type IMorphologyCustomPronoun interface {
	objc.IObject
}

// A custom pronoun behavior for use in a specific langauge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmorphologycustompronoun?language=objc
type MorphologyCustomPronoun struct {
	objc.Object
}

func MorphologyCustomPronounFrom(ptr unsafe.Pointer) MorphologyCustomPronoun {
	return MorphologyCustomPronoun{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MorphologyCustomPronounClass) Alloc() MorphologyCustomPronoun {
	rv := objc.Call[MorphologyCustomPronoun](mc, objc.Sel("alloc"))
	return rv
}

func MorphologyCustomPronoun_Alloc() MorphologyCustomPronoun {
	return MorphologyCustomPronounClass.Alloc()
}

func (mc _MorphologyCustomPronounClass) New() MorphologyCustomPronoun {
	rv := objc.Call[MorphologyCustomPronoun](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMorphologyCustomPronoun() MorphologyCustomPronoun {
	return MorphologyCustomPronounClass.New()
}

func (m_ MorphologyCustomPronoun) Init() MorphologyCustomPronoun {
	rv := objc.Call[MorphologyCustomPronoun](m_, objc.Sel("init"))
	return rv
}
