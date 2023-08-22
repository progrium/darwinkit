// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CompositionTrackFormatDescriptionReplacement] class.
var CompositionTrackFormatDescriptionReplacementClass = _CompositionTrackFormatDescriptionReplacementClass{objc.GetClass("AVCompositionTrackFormatDescriptionReplacement")}

type _CompositionTrackFormatDescriptionReplacementClass struct {
	objc.Class
}

// An interface definition for the [CompositionTrackFormatDescriptionReplacement] class.
type ICompositionTrackFormatDescriptionReplacement interface {
	objc.IObject
	OriginalFormatDescription() coremedia.FormatDescriptionRef
	ReplacementFormatDescription() coremedia.FormatDescriptionRef
}

// An object that represents a format description and its replacement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrackformatdescriptionreplacement?language=objc
type CompositionTrackFormatDescriptionReplacement struct {
	objc.Object
}

func CompositionTrackFormatDescriptionReplacementFrom(ptr unsafe.Pointer) CompositionTrackFormatDescriptionReplacement {
	return CompositionTrackFormatDescriptionReplacement{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CompositionTrackFormatDescriptionReplacementClass) Alloc() CompositionTrackFormatDescriptionReplacement {
	rv := objc.Call[CompositionTrackFormatDescriptionReplacement](cc, objc.Sel("alloc"))
	return rv
}

func CompositionTrackFormatDescriptionReplacement_Alloc() CompositionTrackFormatDescriptionReplacement {
	return CompositionTrackFormatDescriptionReplacementClass.Alloc()
}

func (cc _CompositionTrackFormatDescriptionReplacementClass) New() CompositionTrackFormatDescriptionReplacement {
	rv := objc.Call[CompositionTrackFormatDescriptionReplacement](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCompositionTrackFormatDescriptionReplacement() CompositionTrackFormatDescriptionReplacement {
	return CompositionTrackFormatDescriptionReplacementClass.New()
}

func (c_ CompositionTrackFormatDescriptionReplacement) Init() CompositionTrackFormatDescriptionReplacement {
	rv := objc.Call[CompositionTrackFormatDescriptionReplacement](c_, objc.Sel("init"))
	return rv
}

// The format description to replace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrackformatdescriptionreplacement/3180002-originalformatdescription?language=objc
func (c_ CompositionTrackFormatDescriptionReplacement) OriginalFormatDescription() coremedia.FormatDescriptionRef {
	rv := objc.Call[coremedia.FormatDescriptionRef](c_, objc.Sel("originalFormatDescription"))
	return rv
}

// The replacement format description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrackformatdescriptionreplacement/3180003-replacementformatdescription?language=objc
func (c_ CompositionTrackFormatDescriptionReplacement) ReplacementFormatDescription() coremedia.FormatDescriptionRef {
	rv := objc.Call[coremedia.FormatDescriptionRef](c_, objc.Sel("replacementFormatDescription"))
	return rv
}
