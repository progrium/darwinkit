// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableFontCollection] class.
var MutableFontCollectionClass = _MutableFontCollectionClass{objc.GetClass("NSMutableFontCollection")}

type _MutableFontCollectionClass struct {
	objc.Class
}

// An interface definition for the [MutableFontCollection] class.
type IMutableFontCollection interface {
	IFontCollection
	AddQueryForDescriptors(descriptors []IFontDescriptor)
	RemoveQueryForDescriptors(descriptors []IFontDescriptor)
	SetExclusionDescriptors(value []IFontDescriptor)
	SetQueryDescriptors(value []IFontDescriptor)
}

// A mutable collection of font descriptors taken together as a single object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutablefontcollection?language=objc
type MutableFontCollection struct {
	FontCollection
}

func MutableFontCollectionFrom(ptr unsafe.Pointer) MutableFontCollection {
	return MutableFontCollection{
		FontCollection: FontCollectionFrom(ptr),
	}
}

func (mc _MutableFontCollectionClass) Alloc() MutableFontCollection {
	rv := objc.Call[MutableFontCollection](mc, objc.Sel("alloc"))
	return rv
}

func MutableFontCollection_Alloc() MutableFontCollection {
	return MutableFontCollectionClass.Alloc()
}

func (mc _MutableFontCollectionClass) New() MutableFontCollection {
	rv := objc.Call[MutableFontCollection](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableFontCollection() MutableFontCollection {
	return MutableFontCollectionClass.New()
}

func (m_ MutableFontCollection) Init() MutableFontCollection {
	rv := objc.Call[MutableFontCollection](m_, objc.Sel("init"))
	return rv
}

// Edits the query and exclusion arrays by adding the specified font descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutablefontcollection/1497446-addqueryfordescriptors?language=objc
func (m_ MutableFontCollection) AddQueryForDescriptors(descriptors []IFontDescriptor) {
	objc.Call[objc.Void](m_, objc.Sel("addQueryForDescriptors:"), descriptors)
}

// Edits the query and exclusion arrays by removing the specified font descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutablefontcollection/1497471-removequeryfordescriptors?language=objc
func (m_ MutableFontCollection) RemoveQueryForDescriptors(descriptors []IFontDescriptor) {
	objc.Call[objc.Void](m_, objc.Sel("removeQueryForDescriptors:"), descriptors)
}

// The font descriptors to exclude from query results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutablefontcollection/1497442-exclusiondescriptors?language=objc
func (m_ MutableFontCollection) SetExclusionDescriptors(value []IFontDescriptor) {
	objc.Call[objc.Void](m_, objc.Sel("setExclusionDescriptors:"), value)
}

// The font descriptors to include in query results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutablefontcollection/1497457-querydescriptors?language=objc
func (m_ MutableFontCollection) SetQueryDescriptors(value []IFontDescriptor) {
	objc.Call[objc.Void](m_, objc.Sel("setQueryDescriptors:"), value)
}
