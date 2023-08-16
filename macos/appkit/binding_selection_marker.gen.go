// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BindingSelectionMarker] class.
var BindingSelectionMarkerClass = _BindingSelectionMarkerClass{objc.GetClass("NSBindingSelectionMarker")}

type _BindingSelectionMarkerClass struct {
	objc.Class
}

// An interface definition for the [BindingSelectionMarker] class.
type IBindingSelectionMarker interface {
	objc.IObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbindingselectionmarker?language=objc
type BindingSelectionMarker struct {
	objc.Object
}

func BindingSelectionMarkerFrom(ptr unsafe.Pointer) BindingSelectionMarker {
	return BindingSelectionMarker{
		Object: objc.ObjectFrom(ptr),
	}
}

func (bc _BindingSelectionMarkerClass) Alloc() BindingSelectionMarker {
	rv := objc.Call[BindingSelectionMarker](bc, objc.Sel("alloc"))
	return rv
}

func BindingSelectionMarker_Alloc() BindingSelectionMarker {
	return BindingSelectionMarkerClass.Alloc()
}

func (bc _BindingSelectionMarkerClass) New() BindingSelectionMarker {
	rv := objc.Call[BindingSelectionMarker](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBindingSelectionMarker() BindingSelectionMarker {
	return BindingSelectionMarkerClass.New()
}

func (b_ BindingSelectionMarker) Init() BindingSelectionMarker {
	rv := objc.Call[BindingSelectionMarker](b_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbindingselectionmarker/3088801-defaultplaceholderformarker?language=objc
func (bc _BindingSelectionMarkerClass) DefaultPlaceholderForMarkerOnClassWithBinding(marker IBindingSelectionMarker, objectClass objc.IClass, binding BindingName) objc.Object {
	rv := objc.Call[objc.Object](bc, objc.Sel("defaultPlaceholderForMarker:onClass:withBinding:"), objc.Ptr(marker), objc.Ptr(objectClass), binding)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbindingselectionmarker/3088801-defaultplaceholderformarker?language=objc
func BindingSelectionMarker_DefaultPlaceholderForMarkerOnClassWithBinding(marker IBindingSelectionMarker, objectClass objc.IClass, binding BindingName) objc.Object {
	return BindingSelectionMarkerClass.DefaultPlaceholderForMarkerOnClassWithBinding(marker, objectClass, binding)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbindingselectionmarker/3088802-setdefaultplaceholder?language=objc
func (bc _BindingSelectionMarkerClass) SetDefaultPlaceholderForMarkerOnClassWithBinding(placeholder objc.IObject, marker IBindingSelectionMarker, objectClass objc.IClass, binding BindingName) {
	objc.Call[objc.Void](bc, objc.Sel("setDefaultPlaceholder:forMarker:onClass:withBinding:"), placeholder, objc.Ptr(marker), objc.Ptr(objectClass), binding)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbindingselectionmarker/3088802-setdefaultplaceholder?language=objc
func BindingSelectionMarker_SetDefaultPlaceholderForMarkerOnClassWithBinding(placeholder objc.IObject, marker IBindingSelectionMarker, objectClass objc.IClass, binding BindingName) {
	BindingSelectionMarkerClass.SetDefaultPlaceholderForMarkerOnClassWithBinding(placeholder, marker, objectClass, binding)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbindingselectionmarker/3022483-notapplicableselectionmarker?language=objc
func (bc _BindingSelectionMarkerClass) NotApplicableSelectionMarker() BindingSelectionMarker {
	rv := objc.Call[BindingSelectionMarker](bc, objc.Sel("notApplicableSelectionMarker"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbindingselectionmarker/3022483-notapplicableselectionmarker?language=objc
func BindingSelectionMarker_NotApplicableSelectionMarker() BindingSelectionMarker {
	return BindingSelectionMarkerClass.NotApplicableSelectionMarker()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbindingselectionmarker/3022482-noselectionmarker?language=objc
func (bc _BindingSelectionMarkerClass) NoSelectionMarker() BindingSelectionMarker {
	rv := objc.Call[BindingSelectionMarker](bc, objc.Sel("noSelectionMarker"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbindingselectionmarker/3022482-noselectionmarker?language=objc
func BindingSelectionMarker_NoSelectionMarker() BindingSelectionMarker {
	return BindingSelectionMarkerClass.NoSelectionMarker()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbindingselectionmarker/3022481-multiplevaluesselectionmarker?language=objc
func (bc _BindingSelectionMarkerClass) MultipleValuesSelectionMarker() BindingSelectionMarker {
	rv := objc.Call[BindingSelectionMarker](bc, objc.Sel("multipleValuesSelectionMarker"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbindingselectionmarker/3022481-multiplevaluesselectionmarker?language=objc
func BindingSelectionMarker_MultipleValuesSelectionMarker() BindingSelectionMarker {
	return BindingSelectionMarkerClass.MultipleValuesSelectionMarker()
}
