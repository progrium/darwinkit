// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MediaSelectionOption] class.
var MediaSelectionOptionClass = _MediaSelectionOptionClass{objc.GetClass("AVMediaSelectionOption")}

type _MediaSelectionOptionClass struct {
	objc.Class
}

// An interface definition for the [MediaSelectionOption] class.
type IMediaSelectionOption interface {
	objc.IObject
	DisplayNameWithLocale(locale foundation.ILocale) string
	AssociatedMediaSelectionOptionInMediaSelectionGroup(mediaSelectionGroup IMediaSelectionGroup) MediaSelectionOption
	HasMediaCharacteristic(mediaCharacteristic MediaCharacteristic) bool
	PropertyList() objc.Object
	MetadataForFormat(format string) []MetadataItem
	ExtendedLanguageTag() string
	MediaSubTypes() []foundation.Number
	MediaType() MediaType
	Locale() foundation.Locale
	AvailableMetadataFormats() []string
	DisplayName() string
	IsPlayable() bool
	CommonMetadata() []MetadataItem
}

// An object that represents a specific option for the presentation of media within a group of options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption?language=objc
type MediaSelectionOption struct {
	objc.Object
}

func MediaSelectionOptionFrom(ptr unsafe.Pointer) MediaSelectionOption {
	return MediaSelectionOption{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MediaSelectionOptionClass) Alloc() MediaSelectionOption {
	rv := objc.Call[MediaSelectionOption](mc, objc.Sel("alloc"))
	return rv
}

func MediaSelectionOption_Alloc() MediaSelectionOption {
	return MediaSelectionOptionClass.Alloc()
}

func (mc _MediaSelectionOptionClass) New() MediaSelectionOption {
	rv := objc.Call[MediaSelectionOption](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMediaSelectionOption() MediaSelectionOption {
	return MediaSelectionOptionClass.New()
}

func (m_ MediaSelectionOption) Init() MediaSelectionOption {
	rv := objc.Call[MediaSelectionOption](m_, objc.Sel("init"))
	return rv
}

// Returns a string suitable for display using the specified locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/1388021-displaynamewithlocale?language=objc
func (m_ MediaSelectionOption) DisplayNameWithLocale(locale foundation.ILocale) string {
	rv := objc.Call[string](m_, objc.Sel("displayNameWithLocale:"), objc.Ptr(locale))
	return rv
}

// Returns a media selection option associated with the receiver in a given group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/1388232-associatedmediaselectionoptionin?language=objc
func (m_ MediaSelectionOption) AssociatedMediaSelectionOptionInMediaSelectionGroup(mediaSelectionGroup IMediaSelectionGroup) MediaSelectionOption {
	rv := objc.Call[MediaSelectionOption](m_, objc.Sel("associatedMediaSelectionOptionInMediaSelectionGroup:"), objc.Ptr(mediaSelectionGroup))
	return rv
}

// Returns a Boolean value that indicates whether the receiver has media with the given media characteristic. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/1388531-hasmediacharacteristic?language=objc
func (m_ MediaSelectionOption) HasMediaCharacteristic(mediaCharacteristic MediaCharacteristic) bool {
	rv := objc.Call[bool](m_, objc.Sel("hasMediaCharacteristic:"), mediaCharacteristic)
	return rv
}

// Returns a serializable property list that’s sufficient to identify the option within its group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/1386310-propertylist?language=objc
func (m_ MediaSelectionOption) PropertyList() objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("propertyList"))
	return rv
}

// Returns an array of metadata items—one for each metadata item in the container of a given format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/1386666-metadataforformat?language=objc
func (m_ MediaSelectionOption) MetadataForFormat(format string) []MetadataItem {
	rv := objc.Call[[]MetadataItem](m_, objc.Sel("metadataForFormat:"), format)
	return rv
}

// The IETF BCP 47 language tag associated with the option [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/1387619-extendedlanguagetag?language=objc
func (m_ MediaSelectionOption) ExtendedLanguageTag() string {
	rv := objc.Call[string](m_, objc.Sel("extendedLanguageTag"))
	return rv
}

// The media sub-types of the media data associated with the option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/1385587-mediasubtypes?language=objc
func (m_ MediaSelectionOption) MediaSubTypes() []foundation.Number {
	rv := objc.Call[[]foundation.Number](m_, objc.Sel("mediaSubTypes"))
	return rv
}

// The media type of the media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/1386322-mediatype?language=objc
func (m_ MediaSelectionOption) MediaType() MediaType {
	rv := objc.Call[MediaType](m_, objc.Sel("mediaType"))
	return rv
}

// The locale for which the media option was authored. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/1388436-locale?language=objc
func (m_ MediaSelectionOption) Locale() foundation.Locale {
	rv := objc.Call[foundation.Locale](m_, objc.Sel("locale"))
	return rv
}

// The metadata formats that contain metadata associated with the option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/1389504-availablemetadataformats?language=objc
func (m_ MediaSelectionOption) AvailableMetadataFormats() []string {
	rv := objc.Call[[]string](m_, objc.Sel("availableMetadataFormats"))
	return rv
}

// A string suitable for display using the current system locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/1388485-displayname?language=objc
func (m_ MediaSelectionOption) DisplayName() string {
	rv := objc.Call[string](m_, objc.Sel("displayName"))
	return rv
}

// A Boolean value that indicates whether the media selection option is playable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/1390917-playable?language=objc
func (m_ MediaSelectionOption) IsPlayable() bool {
	rv := objc.Call[bool](m_, objc.Sel("isPlayable"))
	return rv
}

// An array of metadata items for each common metadata key for which a value is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/1387859-commonmetadata?language=objc
func (m_ MediaSelectionOption) CommonMetadata() []MetadataItem {
	rv := objc.Call[[]MetadataItem](m_, objc.Sel("commonMetadata"))
	return rv
}
