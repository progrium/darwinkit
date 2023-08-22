// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetadataItem] class.
var MetadataItemClass = _MetadataItemClass{objc.GetClass("AVMetadataItem")}

type _MetadataItemClass struct {
	objc.Class
}

// An interface definition for the [MetadataItem] class.
type IMetadataItem interface {
	objc.IObject
	StatusOfValueForKeyError(key string, outError foundation.IError) KeyValueStatus
	LoadValuesAsynchronouslyForKeysCompletionHandler(keys []string, handler func())
	KeySpace() MetadataKeySpace
	ExtendedLanguageTag() string
	DataValue() []byte
	Key() objc.Object
	Value() objc.Object
	StringValue() string
	ExtraAttributes() map[MetadataExtraAttributeKey]objc.Object
	NumberValue() foundation.Number
	Locale() foundation.Locale
	DateValue() foundation.Date
	Time() coremedia.Time
	DataType() string
	StartDate() foundation.Date
	Duration() coremedia.Time
	CommonKey() MetadataKey
	Identifier() MetadataIdentifier
}

// A metadata item for an audiovisual asset or one of its tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem?language=objc
type MetadataItem struct {
	objc.Object
}

func MetadataItemFrom(ptr unsafe.Pointer) MetadataItem {
	return MetadataItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MetadataItemClass) Alloc() MetadataItem {
	rv := objc.Call[MetadataItem](mc, objc.Sel("alloc"))
	return rv
}

func MetadataItem_Alloc() MetadataItem {
	return MetadataItemClass.Alloc()
}

func (mc _MetadataItemClass) New() MetadataItem {
	rv := objc.Call[MetadataItem](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataItem() MetadataItem {
	return MetadataItemClass.New()
}

func (m_ MetadataItem) Init() MetadataItem {
	rv := objc.Call[MetadataItem](m_, objc.Sel("init"))
	return rv
}

// Reports whether the value for a given key is immediately available without blocking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1388523-statusofvalueforkey?language=objc
func (m_ MetadataItem) StatusOfValueForKeyError(key string, outError foundation.IError) KeyValueStatus {
	rv := objc.Call[KeyValueStatus](m_, objc.Sel("statusOfValueForKey:error:"), key, objc.Ptr(outError))
	return rv
}

// Returns a metadata key space for the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1390663-keyspaceforidentifier?language=objc
func (mc _MetadataItemClass) KeySpaceForIdentifier(identifier MetadataIdentifier) MetadataKeySpace {
	rv := objc.Call[MetadataKeySpace](mc, objc.Sel("keySpaceForIdentifier:"), identifier)
	return rv
}

// Returns a metadata key space for the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1390663-keyspaceforidentifier?language=objc
func MetadataItem_KeySpaceForIdentifier(identifier MetadataIdentifier) MetadataKeySpace {
	return MetadataItemClass.KeySpaceForIdentifier(identifier)
}

// Returns a metadata identifier for the specified key and key space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1387869-identifierforkey?language=objc
func (mc _MetadataItemClass) IdentifierForKeyKeySpace(key objc.IObject, keySpace MetadataKeySpace) MetadataIdentifier {
	rv := objc.Call[MetadataIdentifier](mc, objc.Sel("identifierForKey:keySpace:"), key, keySpace)
	return rv
}

// Returns a metadata identifier for the specified key and key space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1387869-identifierforkey?language=objc
func MetadataItem_IdentifierForKeyKeySpace(key objc.IObject, keySpace MetadataKeySpace) MetadataIdentifier {
	return MetadataItemClass.IdentifierForKeyKeySpace(key, keySpace)
}

// Returns a metadata key for the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1385613-keyforidentifier?language=objc
func (mc _MetadataItemClass) KeyForIdentifier(identifier MetadataIdentifier) objc.Object {
	rv := objc.Call[objc.Object](mc, objc.Sel("keyForIdentifier:"), identifier)
	return rv
}

// Returns a metadata key for the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1385613-keyforidentifier?language=objc
func MetadataItem_KeyForIdentifier(identifier MetadataIdentifier) objc.Object {
	return MetadataItemClass.KeyForIdentifier(identifier)
}

// Tells the object to load the values of any of the specified keys that aren’t already loaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1387102-loadvaluesasynchronouslyforkeys?language=objc
func (m_ MetadataItem) LoadValuesAsynchronouslyForKeysCompletionHandler(keys []string, handler func()) {
	objc.Call[objc.Void](m_, objc.Sel("loadValuesAsynchronouslyForKeys:completionHandler:"), keys, handler)
}

// Creates a metadata item whose value loads on an on-demand basis only. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1387745-metadataitemwithpropertiesofmeta?language=objc
func (mc _MetadataItemClass) MetadataItemWithPropertiesOfMetadataItemValueLoadingHandler(metadataItem IMetadataItem, handler func(valueRequest MetadataItemValueRequest)) MetadataItem {
	rv := objc.Call[MetadataItem](mc, objc.Sel("metadataItemWithPropertiesOfMetadataItem:valueLoadingHandler:"), objc.Ptr(metadataItem), handler)
	return rv
}

// Creates a metadata item whose value loads on an on-demand basis only. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1387745-metadataitemwithpropertiesofmeta?language=objc
func MetadataItem_MetadataItemWithPropertiesOfMetadataItemValueLoadingHandler(metadataItem IMetadataItem, handler func(valueRequest MetadataItemValueRequest)) MetadataItem {
	return MetadataItemClass.MetadataItemWithPropertiesOfMetadataItemValueLoadingHandler(metadataItem, handler)
}

// Returns metadata items that match a specified locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1389374-metadataitemsfromarray?language=objc
func (mc _MetadataItemClass) MetadataItemsFromArrayWithLocale(metadataItems []IMetadataItem, locale foundation.ILocale) []MetadataItem {
	rv := objc.Call[[]MetadataItem](mc, objc.Sel("metadataItemsFromArray:withLocale:"), metadataItems, objc.Ptr(locale))
	return rv
}

// Returns metadata items that match a specified locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1389374-metadataitemsfromarray?language=objc
func MetadataItem_MetadataItemsFromArrayWithLocale(metadataItems []IMetadataItem, locale foundation.ILocale) []MetadataItem {
	return MetadataItemClass.MetadataItemsFromArrayWithLocale(metadataItems, locale)
}

// The key space for the metadata item’s key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1385757-keyspace?language=objc
func (m_ MetadataItem) KeySpace() MetadataKeySpace {
	rv := objc.Call[MetadataKeySpace](m_, objc.Sel("keySpace"))
	return rv
}

// The IETF BCP 47 (RFC 4646) language identifier of the metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1387068-extendedlanguagetag?language=objc
func (m_ MetadataItem) ExtendedLanguageTag() string {
	rv := objc.Call[string](m_, objc.Sel("extendedLanguageTag"))
	return rv
}

// The value of the metadata item as a data value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1387641-datavalue?language=objc
func (m_ MetadataItem) DataValue() []byte {
	rv := objc.Call[[]byte](m_, objc.Sel("dataValue"))
	return rv
}

// The key of the metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1387843-key?language=objc
func (m_ MetadataItem) Key() objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("key"))
	return rv
}

// The value of the metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1390537-value?language=objc
func (m_ MetadataItem) Value() objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("value"))
	return rv
}

// The value of the metadata item as a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1390846-stringvalue?language=objc
func (m_ MetadataItem) StringValue() string {
	rv := objc.Call[string](m_, objc.Sel("stringValue"))
	return rv
}

// A dictionary of additional attributes for a metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1389570-extraattributes?language=objc
func (m_ MetadataItem) ExtraAttributes() map[MetadataExtraAttributeKey]objc.Object {
	rv := objc.Call[map[MetadataExtraAttributeKey]objc.Object](m_, objc.Sel("extraAttributes"))
	return rv
}

// The value of the metadata item as a number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1390681-numbervalue?language=objc
func (m_ MetadataItem) NumberValue() foundation.Number {
	rv := objc.Call[foundation.Number](m_, objc.Sel("numberValue"))
	return rv
}

// The locale of the metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1387114-locale?language=objc
func (m_ MetadataItem) Locale() foundation.Locale {
	rv := objc.Call[foundation.Locale](m_, objc.Sel("locale"))
	return rv
}

// The value of the metadata item as a date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1385563-datevalue?language=objc
func (m_ MetadataItem) DateValue() foundation.Date {
	rv := objc.Call[foundation.Date](m_, objc.Sel("dateValue"))
	return rv
}

// The timestamp of the metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1388612-time?language=objc
func (m_ MetadataItem) Time() coremedia.Time {
	rv := objc.Call[coremedia.Time](m_, objc.Sel("time"))
	return rv
}

// The data type of the metadata item’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1386856-datatype?language=objc
func (m_ MetadataItem) DataType() string {
	rv := objc.Call[string](m_, objc.Sel("dataType"))
	return rv
}

// The start date of the timed metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1388535-startdate?language=objc
func (m_ MetadataItem) StartDate() foundation.Date {
	rv := objc.Call[foundation.Date](m_, objc.Sel("startDate"))
	return rv
}

// The duration of the metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1386610-duration?language=objc
func (m_ MetadataItem) Duration() coremedia.Time {
	rv := objc.Call[coremedia.Time](m_, objc.Sel("duration"))
	return rv
}

// The common key of the metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1389864-commonkey?language=objc
func (m_ MetadataItem) CommonKey() MetadataKey {
	rv := objc.Call[MetadataKey](m_, objc.Sel("commonKey"))
	return rv
}

// An identifier for a metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/1386968-identifier?language=objc
func (m_ MetadataItem) Identifier() MetadataIdentifier {
	rv := objc.Call[MetadataIdentifier](m_, objc.Sel("identifier"))
	return rv
}
