// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FilterGenerator] class.
var FilterGeneratorClass = _FilterGeneratorClass{objc.GetClass("CIFilterGenerator")}

type _FilterGeneratorClass struct {
	objc.Class
}

// An interface definition for the [FilterGenerator] class.
type IFilterGenerator interface {
	objc.IObject
	RegisterFilterName(name string)
	Filter() Filter
	ExportKeyFromObjectWithName(key string, targetObject objc.IObject, exportedKeyName string)
	ConnectObjectWithKeyToObjectWithKey(sourceObject objc.IObject, sourceKey string, targetObject objc.IObject, targetKey string)
	InitWithContentsOfURL(aURL foundation.IURL) objc.Object
	RemoveExportedKey(exportedKeyName string)
	DisconnectObjectWithKeyToObjectWithKey(sourceObject objc.IObject, sourceKey string, targetObject objc.IObject, targetKey string)
	WriteToURLAtomically(aURL foundation.IURL, flag bool) bool
	SetAttributesForExportedKey(attributes foundation.Dictionary, key string)
	ExportedKeys() foundation.Dictionary
	ClassAttributes() foundation.Dictionary
	SetClassAttributes(value foundation.Dictionary)
}

// An object that creates and configures chains of individual image filters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator?language=objc
type FilterGenerator struct {
	objc.Object
}

func FilterGeneratorFrom(ptr unsafe.Pointer) FilterGenerator {
	return FilterGenerator{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FilterGeneratorClass) Alloc() FilterGenerator {
	rv := objc.Call[FilterGenerator](fc, objc.Sel("alloc"))
	return rv
}

func FilterGenerator_Alloc() FilterGenerator {
	return FilterGeneratorClass.Alloc()
}

func (fc _FilterGeneratorClass) New() FilterGenerator {
	rv := objc.Call[FilterGenerator](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFilterGenerator() FilterGenerator {
	return FilterGeneratorClass.New()
}

func (f_ FilterGenerator) Init() FilterGenerator {
	rv := objc.Call[FilterGenerator](f_, objc.Sel("init"))
	return rv
}

// Registers the name associated with a filter chain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1437891-registerfiltername?language=objc
func (f_ FilterGenerator) RegisterFilterName(name string) {
	objc.Call[objc.Void](f_, objc.Sel("registerFilterName:"), name)
}

// Creates a filter object based on the filter chain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1438044-filter?language=objc
func (f_ FilterGenerator) Filter() Filter {
	rv := objc.Call[Filter](f_, objc.Sel("filter"))
	return rv
}

// Exports an input or output key of an object in the filter chain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1438155-exportkey?language=objc
func (f_ FilterGenerator) ExportKeyFromObjectWithName(key string, targetObject objc.IObject, exportedKeyName string) {
	objc.Call[objc.Void](f_, objc.Sel("exportKey:fromObject:withName:"), key, targetObject, exportedKeyName)
}

// Creates and returns an empty filter generator object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1525954-filtergenerator?language=objc
func (fc _FilterGeneratorClass) FilterGenerator() FilterGenerator {
	rv := objc.Call[FilterGenerator](fc, objc.Sel("filterGenerator"))
	return rv
}

// Creates and returns an empty filter generator object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1525954-filtergenerator?language=objc
func FilterGenerator_FilterGenerator() FilterGenerator {
	return FilterGeneratorClass.FilterGenerator()
}

// Adds an object to the filter chain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1438159-connectobject?language=objc
func (f_ FilterGenerator) ConnectObjectWithKeyToObjectWithKey(sourceObject objc.IObject, sourceKey string, targetObject objc.IObject, targetKey string) {
	objc.Call[objc.Void](f_, objc.Sel("connectObject:withKey:toObject:withKey:"), sourceObject, sourceKey, targetObject, targetKey)
}

// Initializes a filter generator object with the contents of a filter generator file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1437742-initwithcontentsofurl?language=objc
func (f_ FilterGenerator) InitWithContentsOfURL(aURL foundation.IURL) objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("initWithContentsOfURL:"), objc.Ptr(aURL))
	return rv
}

// Removes a key that was previously exported. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1438191-removeexportedkey?language=objc
func (f_ FilterGenerator) RemoveExportedKey(exportedKeyName string) {
	objc.Call[objc.Void](f_, objc.Sel("removeExportedKey:"), exportedKeyName)
}

// Removes the connection between two objects in the filter chain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1438075-disconnectobject?language=objc
func (f_ FilterGenerator) DisconnectObjectWithKeyToObjectWithKey(sourceObject objc.IObject, sourceKey string, targetObject objc.IObject, targetKey string) {
	objc.Call[objc.Void](f_, objc.Sel("disconnectObject:withKey:toObject:withKey:"), sourceObject, sourceKey, targetObject, targetKey)
}

// Archives a filter generator object to a filter generator file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1438179-writetourl?language=objc
func (f_ FilterGenerator) WriteToURLAtomically(aURL foundation.IURL, flag bool) bool {
	rv := objc.Call[bool](f_, objc.Sel("writeToURL:atomically:"), objc.Ptr(aURL), flag)
	return rv
}

// Sets a dictionary of attributes for an exported key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1438069-setattributes?language=objc
func (f_ FilterGenerator) SetAttributesForExportedKey(attributes foundation.Dictionary, key string) {
	objc.Call[objc.Void](f_, objc.Sel("setAttributes:forExportedKey:"), attributes, key)
}

// Creates and returns a filter generator object and initializes it with the contents of a filter generator file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1525950-filtergeneratorwithcontentsofurl?language=objc
func (fc _FilterGeneratorClass) FilterGeneratorWithContentsOfURL(aURL foundation.IURL) FilterGenerator {
	rv := objc.Call[FilterGenerator](fc, objc.Sel("filterGeneratorWithContentsOfURL:"), objc.Ptr(aURL))
	return rv
}

// Creates and returns a filter generator object and initializes it with the contents of a filter generator file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1525950-filtergeneratorwithcontentsofurl?language=objc
func FilterGenerator_FilterGeneratorWithContentsOfURL(aURL foundation.IURL) FilterGenerator {
	return FilterGeneratorClass.FilterGeneratorWithContentsOfURL(aURL)
}

// Returns an array of the exported keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1437955-exportedkeys?language=objc
func (f_ FilterGenerator) ExportedKeys() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](f_, objc.Sel("exportedKeys"))
	return rv
}

// The class attributes associated with the filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1437855-classattributes?language=objc
func (f_ FilterGenerator) ClassAttributes() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](f_, objc.Sel("classAttributes"))
	return rv
}

// The class attributes associated with the filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifiltergenerator/1437855-classattributes?language=objc
func (f_ FilterGenerator) SetClassAttributes(value foundation.Dictionary) {
	objc.Call[objc.Void](f_, objc.Sel("setClassAttributes:"), value)
}
