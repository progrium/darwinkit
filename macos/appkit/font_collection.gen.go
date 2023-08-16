// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FontCollection] class.
var FontCollectionClass = _FontCollectionClass{objc.GetClass("NSFontCollection")}

type _FontCollectionClass struct {
	objc.Class
}

// An interface definition for the [FontCollection] class.
type IFontCollection interface {
	objc.IObject
	MatchingDescriptorsForFamily(family string) []FontDescriptor
	MatchingDescriptorsWithOptions(options map[FontCollectionMatchingOptionKey]foundation.INumber) []FontDescriptor
	ExclusionDescriptors() []FontDescriptor
	QueryDescriptors() []FontDescriptor
	MatchingDescriptors() []FontDescriptor
}

// A font collection, which is a group of font descriptors taken together as a single object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection?language=objc
type FontCollection struct {
	objc.Object
}

func FontCollectionFrom(ptr unsafe.Pointer) FontCollection {
	return FontCollection{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FontCollectionClass) Alloc() FontCollection {
	rv := objc.Call[FontCollection](fc, objc.Sel("alloc"))
	return rv
}

func FontCollection_Alloc() FontCollection {
	return FontCollectionClass.Alloc()
}

func (fc _FontCollectionClass) New() FontCollection {
	rv := objc.Call[FontCollection](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFontCollection() FontCollection {
	return FontCollectionClass.New()
}

func (f_ FontCollection) Init() FontCollection {
	rv := objc.Call[FontCollection](f_, objc.Sel("init"))
	return rv
}

// Creates a named font collection object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497514-fontcollectionwithname?language=objc
func (fc _FontCollectionClass) FontCollectionWithName(name FontCollectionName) FontCollection {
	rv := objc.Call[FontCollection](fc, objc.Sel("fontCollectionWithName:"), name)
	return rv
}

// Creates a named font collection object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497514-fontcollectionwithname?language=objc
func FontCollection_FontCollectionWithName(name FontCollectionName) FontCollection {
	return FontCollectionClass.FontCollectionWithName(name)
}

// Returns an array of font descriptors matching the logical descriptors for the given font family. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497496-matchingdescriptorsforfamily?language=objc
func (f_ FontCollection) MatchingDescriptorsForFamily(family string) []FontDescriptor {
	rv := objc.Call[[]FontDescriptor](f_, objc.Sel("matchingDescriptorsForFamily:"), family)
	return rv
}

// Renames the font collection with the specified name and visibility to the second name specified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497521-renamefontcollectionwithname?language=objc
func (fc _FontCollectionClass) RenameFontCollectionWithNameVisibilityToNameError(oldName FontCollectionName, visibility FontCollectionVisibility, newName FontCollectionName, outError foundation.IError) bool {
	rv := objc.Call[bool](fc, objc.Sel("renameFontCollectionWithName:visibility:toName:error:"), oldName, visibility, newName, objc.Ptr(outError))
	return rv
}

// Renames the font collection with the specified name and visibility to the second name specified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497521-renamefontcollectionwithname?language=objc
func FontCollection_RenameFontCollectionWithNameVisibilityToNameError(oldName FontCollectionName, visibility FontCollectionVisibility, newName FontCollectionName, outError foundation.IError) bool {
	return FontCollectionClass.RenameFontCollectionWithNameVisibilityToNameError(oldName, visibility, newName, outError)
}

// Make the given font collection visible by giving it a name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497512-showfontcollection?language=objc
func (fc _FontCollectionClass) ShowFontCollectionWithNameVisibilityError(collection IFontCollection, name FontCollectionName, visibility FontCollectionVisibility, error foundation.IError) bool {
	rv := objc.Call[bool](fc, objc.Sel("showFontCollection:withName:visibility:error:"), objc.Ptr(collection), name, visibility, objc.Ptr(error))
	return rv
}

// Make the given font collection visible by giving it a name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497512-showfontcollection?language=objc
func FontCollection_ShowFontCollectionWithNameVisibilityError(collection IFontCollection, name FontCollectionName, visibility FontCollectionVisibility, error foundation.IError) bool {
	return FontCollectionClass.ShowFontCollectionWithNameVisibilityError(collection, name, visibility, error)
}

// Remove from view the named font collection with the specified visibility. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497459-hidefontcollectionwithname?language=objc
func (fc _FontCollectionClass) HideFontCollectionWithNameVisibilityError(name FontCollectionName, visibility FontCollectionVisibility, error foundation.IError) bool {
	rv := objc.Call[bool](fc, objc.Sel("hideFontCollectionWithName:visibility:error:"), name, visibility, objc.Ptr(error))
	return rv
}

// Remove from view the named font collection with the specified visibility. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497459-hidefontcollectionwithname?language=objc
func FontCollection_HideFontCollectionWithNameVisibilityError(name FontCollectionName, visibility FontCollectionVisibility, error foundation.IError) bool {
	return FontCollectionClass.HideFontCollectionWithNameVisibilityError(name, visibility, error)
}

// Returns a font collection matching the given descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497467-fontcollectionwithdescriptors?language=objc
func (fc _FontCollectionClass) FontCollectionWithDescriptors(queryDescriptors []IFontDescriptor) FontCollection {
	rv := objc.Call[FontCollection](fc, objc.Sel("fontCollectionWithDescriptors:"), queryDescriptors)
	return rv
}

// Returns a font collection matching the given descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497467-fontcollectionwithdescriptors?language=objc
func FontCollection_FontCollectionWithDescriptors(queryDescriptors []IFontDescriptor) FontCollection {
	return FontCollectionClass.FontCollectionWithDescriptors(queryDescriptors)
}

// Returns a collection of fonts matching the given locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497481-fontcollectionwithlocale?language=objc
func (fc _FontCollectionClass) FontCollectionWithLocale(locale foundation.ILocale) FontCollection {
	rv := objc.Call[FontCollection](fc, objc.Sel("fontCollectionWithLocale:"), objc.Ptr(locale))
	return rv
}

// Returns a collection of fonts matching the given locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497481-fontcollectionwithlocale?language=objc
func FontCollection_FontCollectionWithLocale(locale foundation.ILocale) FontCollection {
	return FontCollectionClass.FontCollectionWithLocale(locale)
}

// Returns an array of font descriptors matching the logical descriptors with the given options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497510-matchingdescriptorswithoptions?language=objc
func (f_ FontCollection) MatchingDescriptorsWithOptions(options map[FontCollectionMatchingOptionKey]foundation.INumber) []FontDescriptor {
	rv := objc.Call[[]FontDescriptor](f_, objc.Sel("matchingDescriptorsWithOptions:"), options)
	return rv
}

// A list of query font descriptors whose matching results are excluded from the list of matching descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497456-exclusiondescriptors?language=objc
func (f_ FontCollection) ExclusionDescriptors() []FontDescriptor {
	rv := objc.Call[[]FontDescriptor](f_, objc.Sel("exclusionDescriptors"))
	return rv
}

// The font collection that matches all registered fonts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497473-fontcollectionwithallavailablede?language=objc
func (fc _FontCollectionClass) FontCollectionWithAllAvailableDescriptors() FontCollection {
	rv := objc.Call[FontCollection](fc, objc.Sel("fontCollectionWithAllAvailableDescriptors"))
	return rv
}

// The font collection that matches all registered fonts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497473-fontcollectionwithallavailablede?language=objc
func FontCollection_FontCollectionWithAllAvailableDescriptors() FontCollection {
	return FontCollectionClass.FontCollectionWithAllAvailableDescriptors()
}

// Returns all named collections visible to this process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497477-allfontcollectionnames?language=objc
func (fc _FontCollectionClass) AllFontCollectionNames() []FontCollectionName {
	rv := objc.Call[[]FontCollectionName](fc, objc.Sel("allFontCollectionNames"))
	return rv
}

// Returns all named collections visible to this process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497477-allfontcollectionnames?language=objc
func FontCollection_AllFontCollectionNames() []FontCollectionName {
	return FontCollectionClass.AllFontCollectionNames()
}

// An array of font descriptors whose matching results produce the collection’s matching descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497441-querydescriptors?language=objc
func (f_ FontCollection) QueryDescriptors() []FontDescriptor {
	rv := objc.Call[[]FontDescriptor](f_, objc.Sel("queryDescriptors"))
	return rv
}

// An array of font descriptors matching the logical descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/1497523-matchingdescriptors?language=objc
func (f_ FontCollection) MatchingDescriptors() []FontDescriptor {
	rv := objc.Call[[]FontDescriptor](f_, objc.Sel("matchingDescriptors"))
	return rv
}
