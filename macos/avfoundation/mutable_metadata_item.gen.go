// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableMetadataItem] class.
var MutableMetadataItemClass = _MutableMetadataItemClass{objc.GetClass("AVMutableMetadataItem")}

type _MutableMetadataItemClass struct {
	objc.Class
}

// An interface definition for the [MutableMetadataItem] class.
type IMutableMetadataItem interface {
	IMetadataItem
	SetKeySpace(value MetadataKeySpace)
	SetExtendedLanguageTag(value string)
	SetKey(value objc.IObject)
	SetValue(value objc.IObject)
	SetExtraAttributes(value map[MetadataExtraAttributeKey]objc.IObject)
	SetLocale(value foundation.ILocale)
	SetTime(value coremedia.Time)
	SetDataType(value string)
	SetStartDate(value foundation.IDate)
	SetDuration(value coremedia.Time)
	SetIdentifier(value MetadataIdentifier)
}

// A mutable metadata item for an audiovisual asset or for one of its tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemetadataitem?language=objc
type MutableMetadataItem struct {
	MetadataItem
}

func MutableMetadataItemFrom(ptr unsafe.Pointer) MutableMetadataItem {
	return MutableMetadataItem{
		MetadataItem: MetadataItemFrom(ptr),
	}
}

func (mc _MutableMetadataItemClass) Alloc() MutableMetadataItem {
	rv := objc.Call[MutableMetadataItem](mc, objc.Sel("alloc"))
	return rv
}

func MutableMetadataItem_Alloc() MutableMetadataItem {
	return MutableMetadataItemClass.Alloc()
}

func (mc _MutableMetadataItemClass) New() MutableMetadataItem {
	rv := objc.Call[MutableMetadataItem](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableMetadataItem() MutableMetadataItem {
	return MutableMetadataItemClass.New()
}

func (m_ MutableMetadataItem) Init() MutableMetadataItem {
	rv := objc.Call[MutableMetadataItem](m_, objc.Sel("init"))
	return rv
}

// Returns a new mutable metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemetadataitem/1426379-metadataitem?language=objc
func (mc _MutableMetadataItemClass) MetadataItem() MutableMetadataItem {
	rv := objc.Call[MutableMetadataItem](mc, objc.Sel("metadataItem"))
	return rv
}

// Returns a new mutable metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemetadataitem/1426379-metadataitem?language=objc
func MutableMetadataItem_MetadataItem() MutableMetadataItem {
	return MutableMetadataItemClass.MetadataItem()
}

// The key space of the metadata item’s key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemetadataitem/1385655-keyspace?language=objc
func (m_ MutableMetadataItem) SetKeySpace(value MetadataKeySpace) {
	objc.Call[objc.Void](m_, objc.Sel("setKeySpace:"), value)
}

// The IETF BCP 47 (RFC 4646) language identifier of the metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemetadataitem/1386664-extendedlanguagetag?language=objc
func (m_ MutableMetadataItem) SetExtendedLanguageTag(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setExtendedLanguageTag:"), value)
}

// The key for a mutable metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemetadataitem/1386776-key?language=objc
func (m_ MutableMetadataItem) SetKey(value objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setKey:"), value)
}

// The value for the mutable metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemetadataitem/1388296-value?language=objc
func (m_ MutableMetadataItem) SetValue(value objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setValue:"), value)
}

// A dictionary of additional attributes for a metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemetadataitem/1390397-extraattributes?language=objc
func (m_ MutableMetadataItem) SetExtraAttributes(value map[MetadataExtraAttributeKey]objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setExtraAttributes:"), value)
}

// The locale for a mutable metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemetadataitem/1389292-locale?language=objc
func (m_ MutableMetadataItem) SetLocale(value foundation.ILocale) {
	objc.Call[objc.Void](m_, objc.Sel("setLocale:"), objc.Ptr(value))
}

// The timestamp for a mutable metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemetadataitem/1389990-time?language=objc
func (m_ MutableMetadataItem) SetTime(value coremedia.Time) {
	objc.Call[objc.Void](m_, objc.Sel("setTime:"), value)
}

// The data type of the metadata item’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemetadataitem/1389471-datatype?language=objc
func (m_ MutableMetadataItem) SetDataType(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setDataType:"), value)
}

// The start date of the timed metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemetadataitem/1389966-startdate?language=objc
func (m_ MutableMetadataItem) SetStartDate(value foundation.IDate) {
	objc.Call[objc.Void](m_, objc.Sel("setStartDate:"), objc.Ptr(value))
}

// The duration of a mutable metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemetadataitem/1389980-duration?language=objc
func (m_ MutableMetadataItem) SetDuration(value coremedia.Time) {
	objc.Call[objc.Void](m_, objc.Sel("setDuration:"), value)
}

// Indicates the identifier of the metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemetadataitem/1386688-identifier?language=objc
func (m_ MutableMetadataItem) SetIdentifier(value MetadataIdentifier) {
	objc.Call[objc.Void](m_, objc.Sel("setIdentifier:"), value)
}
