// AUTO-GENERATED CODE, DO NOT MODIFY

package corespotlight

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/uti"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SearchableItemAttributeSet] class.
var SearchableItemAttributeSetClass = _SearchableItemAttributeSetClass{objc.GetClass("CSSearchableItemAttributeSet")}

type _SearchableItemAttributeSetClass struct {
	objc.Class
}

// An interface definition for the [SearchableItemAttributeSet] class.
type ISearchableItemAttributeSet interface {
	objc.IObject
	ValueForCustomKey(key ICustomAttributeKey) objc.Object
	SetValueForCustomKey(value objc.IObject, key ICustomAttributeKey)
	GPSMapDatum() string
	SetGPSMapDatum(value string)
	AuthorAddresses() []string
	SetAuthorAddresses(value []string)
	Path() string
	SetPath(value string)
	KeySignature() string
	SetKeySignature(value string)
	Publishers() []string
	SetPublishers(value []string)
	IsGeneralMIDISequence() foundation.Number
	SetGeneralMIDISequence(value foundation.INumber)
	Headline() string
	SetHeadline(value string)
	ContainerOrder() foundation.Number
	SetContainerOrder(value foundation.INumber)
	Latitude() foundation.Number
	SetLatitude(value foundation.INumber)
	PixelWidth() foundation.Number
	SetPixelWidth(value foundation.INumber)
	ImportantDates() []foundation.Date
	SetImportantDates(value []foundation.IDate)
	PhoneNumbers() []string
	SetPhoneNumbers(value []string)
	HiddenAdditionalRecipients() []Person
	SetHiddenAdditionalRecipients(value []IPerson)
	Creator() string
	SetCreator(value string)
	Version() string
	SetVersion(value string)
	OriginalFormat() string
	SetOriginalFormat(value string)
	PageHeight() foundation.Number
	SetPageHeight(value foundation.INumber)
	MusicalInstrumentCategory() string
	SetMusicalInstrumentCategory(value string)
	GPSDifferental() foundation.Number
	SetGPSDifferental(value foundation.INumber)
	FontNames() []string
	SetFontNames(value []string)
	ThumbnailURL() foundation.URL
	SetThumbnailURL(value foundation.IURL)
	Languages() []string
	SetLanguages(value []string)
	ProfileName() string
	SetProfileName(value string)
	Audiences() []string
	SetAudiences(value []string)
	GPSDOP() foundation.Number
	SetGPSDOP(value foundation.INumber)
	TextContent() string
	SetTextContent(value string)
	CompletionDate() foundation.Date
	SetCompletionDate(value foundation.IDate)
	Theme() string
	SetTheme(value string)
	PixelHeight() foundation.Number
	SetPixelHeight(value foundation.INumber)
	HasAlphaChannel() foundation.Number
	SetHasAlphaChannel(value foundation.INumber)
	ResolutionWidthDPI() foundation.Number
	SetResolutionWidthDPI(value foundation.INumber)
	Aperture() foundation.Number
	SetAperture(value foundation.INumber)
	FullyFormattedAddress() string
	SetFullyFormattedAddress(value string)
	ContentURL() foundation.URL
	SetContentURL(value foundation.IURL)
	AudioEncodingApplication() string
	SetAudioEncodingApplication(value string)
	EXIFVersion() string
	SetEXIFVersion(value string)
	RankingHint() foundation.Number
	SetRankingHint(value foundation.INumber)
	LastUsedDate() foundation.Date
	SetLastUsedDate(value foundation.IDate)
	ContentModificationDate() foundation.Date
	SetContentModificationDate(value foundation.IDate)
	EncodingApplications() []string
	SetEncodingApplications(value []string)
	GPSMeasureMode() string
	SetGPSMeasureMode(value string)
	Role() string
	SetRole(value string)
	SupportsNavigation() foundation.Number
	SetSupportsNavigation(value foundation.INumber)
	StateOrProvince() string
	SetStateOrProvince(value string)
	AdditionalRecipients() []Person
	SetAdditionalRecipients(value []IPerson)
	AlternateNames() []string
	SetAlternateNames(value []string)
	ProviderFileTypeIdentifiers() []string
	SetProviderFileTypeIdentifiers(value []string)
	AudioChannelCount() foundation.Number
	SetAudioChannelCount(value foundation.INumber)
	Rating() foundation.Number
	SetRating(value foundation.INumber)
	IsStreamable() foundation.Number
	SetStreamable(value foundation.INumber)
	AudioTrackNumber() foundation.Number
	SetAudioTrackNumber(value foundation.INumber)
	ContentTypeTree() []string
	SetContentTypeTree(value []string)
	MusicalGenre() string
	SetMusicalGenre(value string)
	GPSStatus() string
	SetGPSStatus(value string)
	Genre() string
	SetGenre(value string)
	SecurityMethod() string
	SetSecurityMethod(value string)
	GPSDestBearing() foundation.Number
	SetGPSDestBearing(value foundation.INumber)
	AcquisitionMake() string
	SetAcquisitionMake(value string)
	IsUserCurated() foundation.Number
	SetUserCurated(value foundation.INumber)
	Codecs() []string
	SetCodecs(value []string)
	DeliveryType() foundation.Number
	SetDeliveryType(value foundation.INumber)
	ImageDirection() foundation.Number
	SetImageDirection(value foundation.INumber)
	Artist() string
	SetArtist(value string)
	City() string
	SetCity(value string)
	ExposureTimeString() string
	SetExposureTimeString(value string)
	ProviderInPlaceFileTypeIdentifiers() []string
	SetProviderInPlaceFileTypeIdentifiers(value []string)
	ContainerTitle() string
	SetContainerTitle(value string)
	WhiteBalance() foundation.Number
	SetWhiteBalance(value foundation.INumber)
	Producer() string
	SetProducer(value string)
	Contributors() []string
	SetContributors(value []string)
	Comment() string
	SetComment(value string)
	GPSTrack() foundation.Number
	SetGPSTrack(value foundation.INumber)
	ContentDescription() string
	SetContentDescription(value string)
	Organizations() []string
	SetOrganizations(value []string)
	Authors() []Person
	SetAuthors(value []IPerson)
	Performers() []string
	SetPerformers(value []string)
	Speed() foundation.Number
	SetSpeed(value foundation.INumber)
	Instructions() string
	SetInstructions(value string)
	PageCount() foundation.Number
	SetPageCount(value foundation.INumber)
	GPSDestLatitude() foundation.Number
	SetGPSDestLatitude(value foundation.INumber)
	FNumber() foundation.Number
	SetFNumber(value foundation.INumber)
	Director() string
	SetDirector(value string)
	Information() string
	SetInformation(value string)
	GPSDateStamp() foundation.Date
	SetGPSDateStamp(value foundation.IDate)
	ColorSpace() string
	SetColorSpace(value string)
	TotalBitRate() foundation.Number
	SetTotalBitRate(value foundation.INumber)
	AccountHandles() []string
	SetAccountHandles(value []string)
	AcquisitionModel() string
	SetAcquisitionModel(value string)
	MailboxIdentifiers() []string
	SetMailboxIdentifiers(value []string)
	ISOSpeed() foundation.Number
	SetISOSpeed(value foundation.INumber)
	MediaTypes() []string
	SetMediaTypes(value []string)
	ContentCreationDate() foundation.Date
	SetContentCreationDate(value foundation.IDate)
	PostalCode() string
	SetPostalCode(value string)
	TimeSignature() string
	SetTimeSignature(value string)
	DarkThumbnailURL() foundation.URL
	SetDarkThumbnailURL(value foundation.IURL)
	DomainIdentifier() string
	SetDomainIdentifier(value string)
	IsFlashOn() foundation.Number
	SetFlashOn(value foundation.INumber)
	ContentRating() foundation.Number
	SetContentRating(value foundation.INumber)
	Altitude() foundation.Number
	SetAltitude(value foundation.INumber)
	ContainerDisplayName() string
	SetContainerDisplayName(value string)
	ProviderDataTypeIdentifiers() []string
	SetProviderDataTypeIdentifiers(value []string)
	Projects() []string
	SetProjects(value []string)
	NamedLocation() string
	SetNamedLocation(value string)
	DownloadedDate() foundation.Date
	SetDownloadedDate(value foundation.IDate)
	RecipientNames() []string
	SetRecipientNames(value []string)
	RecipientAddresses() []string
	SetRecipientAddresses(value []string)
	Longitude() foundation.Number
	SetLongitude(value foundation.INumber)
	EmailAddresses() []string
	SetEmailAddresses(value []string)
	InstantMessageAddresses() []string
	SetInstantMessageAddresses(value []string)
	PlayCount() foundation.Number
	SetPlayCount(value foundation.INumber)
	AllDay() foundation.Number
	SetAllDay(value foundation.INumber)
	PixelCount() foundation.Number
	SetPixelCount(value foundation.INumber)
	Lyricist() string
	SetLyricist(value string)
	GPSProcessingMethod() string
	SetGPSProcessingMethod(value string)
	ExposureTime() foundation.Number
	SetExposureTime(value foundation.INumber)
	URL() foundation.URL
	SetURL(value foundation.IURL)
	RatingDescription() string
	SetRatingDescription(value string)
	Kind() string
	SetKind(value string)
	ContentType() string
	SetContentType(value string)
	DueDate() foundation.Date
	SetDueDate(value foundation.IDate)
	MusicalInstrumentName() string
	SetMusicalInstrumentName(value string)
	Rights() string
	SetRights(value string)
	AuthorEmailAddresses() []string
	SetAuthorEmailAddresses(value []string)
	AudioSampleRate() foundation.Number
	SetAudioSampleRate(value foundation.INumber)
	MetadataModificationDate() foundation.Date
	SetMetadataModificationDate(value foundation.IDate)
	PrimaryRecipients() []Person
	SetPrimaryRecipients(value []IPerson)
	AudioBitRate() foundation.Number
	SetAudioBitRate(value foundation.INumber)
	IsFocalLength35mm() foundation.Number
	SetFocalLength35mm(value foundation.INumber)
	EXIFGPSVersion() string
	SetEXIFGPSVersion(value string)
	MaxAperture() foundation.Number
	SetMaxAperture(value foundation.INumber)
	IsLikelyJunk() foundation.Number
	SetLikelyJunk(value foundation.INumber)
	Coverage() []string
	SetCoverage(value []string)
	StartDate() foundation.Date
	SetStartDate(value foundation.IDate)
	FileSize() foundation.Number
	SetFileSize(value foundation.INumber)
	IsUserCreated() foundation.Number
	SetUserCreated(value foundation.INumber)
	Title() string
	SetTitle(value string)
	ExposureMode() foundation.Number
	SetExposureMode(value foundation.INumber)
	ThumbnailData() []byte
	SetThumbnailData(value []byte)
	RelatedUniqueIdentifier() string
	SetRelatedUniqueIdentifier(value string)
	Participants() []string
	SetParticipants(value []string)
	AccountIdentifier() string
	SetAccountIdentifier(value string)
	Album() string
	SetAlbum(value string)
	OriginalSource() string
	SetOriginalSource(value string)
	Editors() []string
	SetEditors(value []string)
	Subject() string
	SetSubject(value string)
	Timestamp() foundation.Date
	SetTimestamp(value foundation.IDate)
	RecordingDate() foundation.Date
	SetRecordingDate(value foundation.IDate)
	ResolutionHeightDPI() foundation.Number
	SetResolutionHeightDPI(value foundation.INumber)
	ContactKeywords() []string
	SetContactKeywords(value []string)
	LayerNames() []string
	SetLayerNames(value []string)
	ExposureProgram() string
	SetExposureProgram(value string)
	AddedDate() foundation.Date
	SetAddedDate(value foundation.IDate)
	ContainerIdentifier() string
	SetContainerIdentifier(value string)
	GPSDestDistance() foundation.Number
	SetGPSDestDistance(value foundation.INumber)
	CameraOwner() string
	SetCameraOwner(value string)
	Copyright() string
	SetCopyright(value string)
	MeteringMode() string
	SetMeteringMode(value string)
	Duration() foundation.Number
	SetDuration(value foundation.INumber)
	EndDate() foundation.Date
	SetEndDate(value foundation.IDate)
	SubThoroughfare() string
	SetSubThoroughfare(value string)
	RecipientEmailAddresses() []string
	SetRecipientEmailAddresses(value []string)
	FocalLength() foundation.Number
	SetFocalLength(value foundation.INumber)
	GPSDestLongitude() foundation.Number
	SetGPSDestLongitude(value foundation.INumber)
	VideoBitRate() foundation.Number
	SetVideoBitRate(value foundation.INumber)
	SupportsPhoneCall() foundation.Number
	SetSupportsPhoneCall(value foundation.INumber)
	LensModel() string
	SetLensModel(value string)
	EmailHeaders() map[string][]objc.Object
	SetEmailHeaders(value map[string][]objc.IObject)
	DisplayName() string
	SetDisplayName(value string)
	Composer() string
	SetComposer(value string)
	HTMLContentData() []byte
	SetHTMLContentData(value []byte)
	Identifier() string
	SetIdentifier(value string)
	BitsPerSample() foundation.Number
	SetBitsPerSample(value foundation.INumber)
	Keywords() []string
	SetKeywords(value []string)
	Country() string
	SetCountry(value string)
	AuthorNames() []string
	SetAuthorNames(value []string)
	IsLocal() foundation.Number
	SetLocal(value foundation.INumber)
	Tempo() foundation.Number
	SetTempo(value foundation.INumber)
	GPSAreaInformation() string
	SetGPSAreaInformation(value string)
	ContentSources() []string
	SetContentSources(value []string)
	IsRedEyeOn() foundation.Number
	SetRedEyeOn(value foundation.INumber)
	WeakRelatedUniqueIdentifier() string
	SetWeakRelatedUniqueIdentifier(value string)
	Thoroughfare() string
	SetThoroughfare(value string)
	Orientation() foundation.Number
	SetOrientation(value foundation.INumber)
	PageWidth() foundation.Number
	SetPageWidth(value foundation.INumber)
	IsUserOwned() foundation.Number
	SetUserOwned(value foundation.INumber)
}

// The set of properties to display for a searchable item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset?language=objc
type SearchableItemAttributeSet struct {
	objc.Object
}

func SearchableItemAttributeSetFrom(ptr unsafe.Pointer) SearchableItemAttributeSet {
	return SearchableItemAttributeSet{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SearchableItemAttributeSet) InitWithContentType(contentType uti.IType) SearchableItemAttributeSet {
	rv := objc.Call[SearchableItemAttributeSet](s_, objc.Sel("initWithContentType:"), objc.Ptr(contentType))
	return rv
}

// Creates an attribute set for the specified content type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/3573909-initwithcontenttype?language=objc
func SearchableItemAttributeSet_InitWithContentType(contentType uti.IType) SearchableItemAttributeSet {
	return SearchableItemAttributeSetClass.Alloc().InitWithContentType(contentType)
}

func (sc _SearchableItemAttributeSetClass) Alloc() SearchableItemAttributeSet {
	rv := objc.Call[SearchableItemAttributeSet](sc, objc.Sel("alloc"))
	return rv
}

func SearchableItemAttributeSet_Alloc() SearchableItemAttributeSet {
	return SearchableItemAttributeSetClass.Alloc()
}

func (sc _SearchableItemAttributeSetClass) New() SearchableItemAttributeSet {
	rv := objc.Call[SearchableItemAttributeSet](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSearchableItemAttributeSet() SearchableItemAttributeSet {
	return SearchableItemAttributeSetClass.New()
}

func (s_ SearchableItemAttributeSet) Init() SearchableItemAttributeSet {
	rv := objc.Call[SearchableItemAttributeSet](s_, objc.Sel("init"))
	return rv
}

// Returns the value associated with the specified custom attribute key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616407-valueforcustomkey?language=objc
func (s_ SearchableItemAttributeSet) ValueForCustomKey(key ICustomAttributeKey) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("valueForCustomKey:"), objc.Ptr(key))
	return rv
}

// Sets the value for a custom attribute key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616395-setvalue?language=objc
func (s_ SearchableItemAttributeSet) SetValueForCustomKey(value objc.IObject, key ICustomAttributeKey) {
	objc.Call[objc.Void](s_, objc.Sel("setValue:forCustomKey:"), value, objc.Ptr(key))
}

// The geodetic data that the GPS receiver uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620583-gpsmapdatum?language=objc
func (s_ SearchableItemAttributeSet) GPSMapDatum() string {
	rv := objc.Call[string](s_, objc.Sel("GPSMapDatum"))
	return rv
}

// The geodetic data that the GPS receiver uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620583-gpsmapdatum?language=objc
func (s_ SearchableItemAttributeSet) SetGPSMapDatum(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setGPSMapDatum:"), value)
}

// An array of addresses associated with the author of the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621671-authoraddresses?language=objc
func (s_ SearchableItemAttributeSet) AuthorAddresses() []string {
	rv := objc.Call[[]string](s_, objc.Sel("authorAddresses"))
	return rv
}

// An array of addresses associated with the author of the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621671-authoraddresses?language=objc
func (s_ SearchableItemAttributeSet) SetAuthorAddresses(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setAuthorAddresses:"), value)
}

// The complete path to the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621559-path?language=objc
func (s_ SearchableItemAttributeSet) Path() string {
	rv := objc.Call[string](s_, objc.Sel("path"))
	return rv
}

// The complete path to the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621559-path?language=objc
func (s_ SearchableItemAttributeSet) SetPath(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setPath:"), value)
}

// The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616020-keysignature?language=objc
func (s_ SearchableItemAttributeSet) KeySignature() string {
	rv := objc.Call[string](s_, objc.Sel("keySignature"))
	return rv
}

// The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616020-keysignature?language=objc
func (s_ SearchableItemAttributeSet) SetKeySignature(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setKeySignature:"), value)
}

// A list of people, organizations, services, or other entities responsible for making the media available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616003-publishers?language=objc
func (s_ SearchableItemAttributeSet) Publishers() []string {
	rv := objc.Call[[]string](s_, objc.Sel("publishers"))
	return rv
}

// A list of people, organizations, services, or other entities responsible for making the media available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616003-publishers?language=objc
func (s_ SearchableItemAttributeSet) SetPublishers(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setPublishers:"), value)
}

// A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616081-generalmidisequence?language=objc
func (s_ SearchableItemAttributeSet) IsGeneralMIDISequence() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("isGeneralMIDISequence"))
	return rv
}

// A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616081-generalmidisequence?language=objc
func (s_ SearchableItemAttributeSet) SetGeneralMIDISequence(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setGeneralMIDISequence:"), objc.Ptr(value))
}

// A publishable string that provides a synopsis of the contents of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620585-headline?language=objc
func (s_ SearchableItemAttributeSet) Headline() string {
	rv := objc.Call[string](s_, objc.Sel("headline"))
	return rv
}

// A publishable string that provides a synopsis of the contents of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620585-headline?language=objc
func (s_ SearchableItemAttributeSet) SetHeadline(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setHeadline:"), value)
}

// The order of the item within the container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621586-containerorder?language=objc
func (s_ SearchableItemAttributeSet) ContainerOrder() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("containerOrder"))
	return rv
}

// The order of the item within the container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621586-containerorder?language=objc
func (s_ SearchableItemAttributeSet) SetContainerOrder(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setContainerOrder:"), objc.Ptr(value))
}

// The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620586-latitude?language=objc
func (s_ SearchableItemAttributeSet) Latitude() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("latitude"))
	return rv
}

// The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620586-latitude?language=objc
func (s_ SearchableItemAttributeSet) SetLatitude(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setLatitude:"), objc.Ptr(value))
}

// The width of the item, such as image or video frame width, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621527-pixelwidth?language=objc
func (s_ SearchableItemAttributeSet) PixelWidth() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("pixelWidth"))
	return rv
}

// The width of the item, such as image or video frame width, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621527-pixelwidth?language=objc
func (s_ SearchableItemAttributeSet) SetPixelWidth(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setPixelWidth:"), objc.Ptr(value))
}

// An array of important dates associated with the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616639-importantdates?language=objc
func (s_ SearchableItemAttributeSet) ImportantDates() []foundation.Date {
	rv := objc.Call[[]foundation.Date](s_, objc.Sel("importantDates"))
	return rv
}

// An array of important dates associated with the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616639-importantdates?language=objc
func (s_ SearchableItemAttributeSet) SetImportantDates(value []foundation.IDate) {
	objc.Call[objc.Void](s_, objc.Sel("setImportantDates:"), value)
}

// An array of phone numbers associated with the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621650-phonenumbers?language=objc
func (s_ SearchableItemAttributeSet) PhoneNumbers() []string {
	rv := objc.Call[[]string](s_, objc.Sel("phoneNumbers"))
	return rv
}

// An array of phone numbers associated with the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621650-phonenumbers?language=objc
func (s_ SearchableItemAttributeSet) SetPhoneNumbers(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setPhoneNumbers:"), value)
}

// An array of CSPerson objects representing the content of the Bcc: field in an email message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621645-hiddenadditionalrecipients?language=objc
func (s_ SearchableItemAttributeSet) HiddenAdditionalRecipients() []Person {
	rv := objc.Call[[]Person](s_, objc.Sel("hiddenAdditionalRecipients"))
	return rv
}

// An array of CSPerson objects representing the content of the Bcc: field in an email message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621645-hiddenadditionalrecipients?language=objc
func (s_ SearchableItemAttributeSet) SetHiddenAdditionalRecipients(value []IPerson) {
	objc.Call[objc.Void](s_, objc.Sel("setHiddenAdditionalRecipients:"), value)
}

// The name of the app that created the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621677-creator?language=objc
func (s_ SearchableItemAttributeSet) Creator() string {
	rv := objc.Call[string](s_, objc.Sel("creator"))
	return rv
}

// The name of the app that created the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621677-creator?language=objc
func (s_ SearchableItemAttributeSet) SetCreator(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setCreator:"), value)
}

// A version string associated with the file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616093-version?language=objc
func (s_ SearchableItemAttributeSet) Version() string {
	rv := objc.Call[string](s_, objc.Sel("version"))
	return rv
}

// A version string associated with the file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616093-version?language=objc
func (s_ SearchableItemAttributeSet) SetVersion(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setVersion:"), value)
}

// The original format of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616097-originalformat?language=objc
func (s_ SearchableItemAttributeSet) OriginalFormat() string {
	rv := objc.Call[string](s_, objc.Sel("originalFormat"))
	return rv
}

// The original format of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616097-originalformat?language=objc
func (s_ SearchableItemAttributeSet) SetOriginalFormat(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setOriginalFormat:"), value)
}

// The height of the document page, in points (72 points per inch). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621673-pageheight?language=objc
func (s_ SearchableItemAttributeSet) PageHeight() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("pageHeight"))
	return rv
}

// The height of the document page, in points (72 points per inch). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621673-pageheight?language=objc
func (s_ SearchableItemAttributeSet) SetPageHeight(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setPageHeight:"), objc.Ptr(value))
}

// The category of the instrument associated with the audio file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616099-musicalinstrumentcategory?language=objc
func (s_ SearchableItemAttributeSet) MusicalInstrumentCategory() string {
	rv := objc.Call[string](s_, objc.Sel("musicalInstrumentCategory"))
	return rv
}

// The category of the instrument associated with the audio file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616099-musicalinstrumentcategory?language=objc
func (s_ SearchableItemAttributeSet) SetMusicalInstrumentCategory(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setMusicalInstrumentCategory:"), value)
}

// The differential correction applied to the GPS receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620592-gpsdifferental?language=objc
func (s_ SearchableItemAttributeSet) GPSDifferental() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("GPSDifferental"))
	return rv
}

// The differential correction applied to the GPS receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620592-gpsdifferental?language=objc
func (s_ SearchableItemAttributeSet) SetGPSDifferental(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setGPSDifferental:"), objc.Ptr(value))
}

// An array of font names the document uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621566-fontnames?language=objc
func (s_ SearchableItemAttributeSet) FontNames() []string {
	rv := objc.Call[[]string](s_, objc.Sel("fontNames"))
	return rv
}

// An array of font names the document uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621566-fontnames?language=objc
func (s_ SearchableItemAttributeSet) SetFontNames(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setFontNames:"), value)
}

// The local file URL of the thumbnail image for the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621560-thumbnailurl?language=objc
func (s_ SearchableItemAttributeSet) ThumbnailURL() foundation.URL {
	rv := objc.Call[foundation.URL](s_, objc.Sel("thumbnailURL"))
	return rv
}

// The local file URL of the thumbnail image for the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621560-thumbnailurl?language=objc
func (s_ SearchableItemAttributeSet) SetThumbnailURL(value foundation.IURL) {
	objc.Call[objc.Void](s_, objc.Sel("setThumbnailURL:"), objc.Ptr(value))
}

// A list of the included languages for the intellectual content of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616068-languages?language=objc
func (s_ SearchableItemAttributeSet) Languages() []string {
	rv := objc.Call[[]string](s_, objc.Sel("languages"))
	return rv
}

// A list of the included languages for the intellectual content of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616068-languages?language=objc
func (s_ SearchableItemAttributeSet) SetLanguages(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setLanguages:"), value)
}

// The name of the color profile the camera used for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621531-profilename?language=objc
func (s_ SearchableItemAttributeSet) ProfileName() string {
	rv := objc.Call[string](s_, objc.Sel("profileName"))
	return rv
}

// The name of the color profile the camera used for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621531-profilename?language=objc
func (s_ SearchableItemAttributeSet) SetProfileName(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setProfileName:"), value)
}

// A class of entity for which the item is intended or useful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621571-audiences?language=objc
func (s_ SearchableItemAttributeSet) Audiences() []string {
	rv := objc.Call[[]string](s_, objc.Sel("audiences"))
	return rv
}

// A class of entity for which the item is intended or useful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621571-audiences?language=objc
func (s_ SearchableItemAttributeSet) SetAudiences(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setAudiences:"), value)
}

// The GPS dilution of precision value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620593-gpsdop?language=objc
func (s_ SearchableItemAttributeSet) GPSDOP() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("GPSDOP"))
	return rv
}

// The GPS dilution of precision value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620593-gpsdop?language=objc
func (s_ SearchableItemAttributeSet) SetGPSDOP(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setGPSDOP:"), objc.Ptr(value))
}

// The textual content of the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621674-textcontent?language=objc
func (s_ SearchableItemAttributeSet) TextContent() string {
	rv := objc.Call[string](s_, objc.Sel("textContent"))
	return rv
}

// The textual content of the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621674-textcontent?language=objc
func (s_ SearchableItemAttributeSet) SetTextContent(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setTextContent:"), value)
}

// The date on which the item was completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616637-completiondate?language=objc
func (s_ SearchableItemAttributeSet) CompletionDate() foundation.Date {
	rv := objc.Call[foundation.Date](s_, objc.Sel("completionDate"))
	return rv
}

// The date on which the item was completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616637-completiondate?language=objc
func (s_ SearchableItemAttributeSet) SetCompletionDate(value foundation.IDate) {
	objc.Call[objc.Void](s_, objc.Sel("setCompletionDate:"), objc.Ptr(value))
}

// The theme of the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621574-theme?language=objc
func (s_ SearchableItemAttributeSet) Theme() string {
	rv := objc.Call[string](s_, objc.Sel("theme"))
	return rv
}

// The theme of the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621574-theme?language=objc
func (s_ SearchableItemAttributeSet) SetTheme(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setTheme:"), value)
}

// The height of the item, such as image or video frame height, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621535-pixelheight?language=objc
func (s_ SearchableItemAttributeSet) PixelHeight() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("pixelHeight"))
	return rv
}

// The height of the item, such as image or video frame height, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621535-pixelheight?language=objc
func (s_ SearchableItemAttributeSet) SetPixelHeight(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setPixelHeight:"), objc.Ptr(value))
}

// Indicates if the image file has an alpha channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621518-hasalphachannel?language=objc
func (s_ SearchableItemAttributeSet) HasAlphaChannel() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("hasAlphaChannel"))
	return rv
}

// Indicates if the image file has an alpha channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621518-hasalphachannel?language=objc
func (s_ SearchableItemAttributeSet) SetHasAlphaChannel(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setHasAlphaChannel:"), objc.Ptr(value))
}

// The resolution width of the image, in DPI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621544-resolutionwidthdpi?language=objc
func (s_ SearchableItemAttributeSet) ResolutionWidthDPI() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("resolutionWidthDPI"))
	return rv
}

// The resolution width of the image, in DPI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621544-resolutionwidthdpi?language=objc
func (s_ SearchableItemAttributeSet) SetResolutionWidthDPI(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setResolutionWidthDPI:"), objc.Ptr(value))
}

// The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621528-aperture?language=objc
func (s_ SearchableItemAttributeSet) Aperture() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("aperture"))
	return rv
}

// The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621528-aperture?language=objc
func (s_ SearchableItemAttributeSet) SetAperture(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setAperture:"), objc.Ptr(value))
}

// The fully formatted address of the item, received from MapKit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1649301-fullyformattedaddress?language=objc
func (s_ SearchableItemAttributeSet) FullyFormattedAddress() string {
	rv := objc.Call[string](s_, objc.Sel("fullyFormattedAddress"))
	return rv
}

// The fully formatted address of the item, received from MapKit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1649301-fullyformattedaddress?language=objc
func (s_ SearchableItemAttributeSet) SetFullyFormattedAddress(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setFullyFormattedAddress:"), value)
}

// The file URL of the content to index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621651-contenturl?language=objc
func (s_ SearchableItemAttributeSet) ContentURL() foundation.URL {
	rv := objc.Call[foundation.URL](s_, objc.Sel("contentURL"))
	return rv
}

// The file URL of the content to index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621651-contenturl?language=objc
func (s_ SearchableItemAttributeSet) SetContentURL(value foundation.IURL) {
	objc.Call[objc.Void](s_, objc.Sel("setContentURL:"), objc.Ptr(value))
}

// The name of the application that encoded the data the audio file contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616039-audioencodingapplication?language=objc
func (s_ SearchableItemAttributeSet) AudioEncodingApplication() string {
	rv := objc.Call[string](s_, objc.Sel("audioEncodingApplication"))
	return rv
}

// The name of the application that encoded the data the audio file contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616039-audioencodingapplication?language=objc
func (s_ SearchableItemAttributeSet) SetAudioEncodingApplication(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setAudioEncodingApplication:"), value)
}

// The version of the EXIF header that was used to generate the metadata for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621550-exifversion?language=objc
func (s_ SearchableItemAttributeSet) EXIFVersion() string {
	rv := objc.Call[string](s_, objc.Sel("EXIFVersion"))
	return rv
}

// The version of the EXIF header that was used to generate the metadata for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621550-exifversion?language=objc
func (s_ SearchableItemAttributeSet) SetEXIFVersion(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setEXIFVersion:"), value)
}

// A number that indicates the relative importance of the item among other items from the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/2887563-rankinghint?language=objc
func (s_ SearchableItemAttributeSet) RankingHint() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("rankingHint"))
	return rv
}

// A number that indicates the relative importance of the item among other items from the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/2887563-rankinghint?language=objc
func (s_ SearchableItemAttributeSet) SetRankingHint(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setRankingHint:"), objc.Ptr(value))
}

// The date on which the file was last used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616018-lastuseddate?language=objc
func (s_ SearchableItemAttributeSet) LastUsedDate() foundation.Date {
	rv := objc.Call[foundation.Date](s_, objc.Sel("lastUsedDate"))
	return rv
}

// The date on which the file was last used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616018-lastuseddate?language=objc
func (s_ SearchableItemAttributeSet) SetLastUsedDate(value foundation.IDate) {
	objc.Call[objc.Void](s_, objc.Sel("setLastUsedDate:"), objc.Ptr(value))
}

// The date on which the contents of the file was last modified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616085-contentmodificationdate?language=objc
func (s_ SearchableItemAttributeSet) ContentModificationDate() foundation.Date {
	rv := objc.Call[foundation.Date](s_, objc.Sel("contentModificationDate"))
	return rv
}

// The date on which the contents of the file was last modified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616085-contentmodificationdate?language=objc
func (s_ SearchableItemAttributeSet) SetContentModificationDate(value foundation.IDate) {
	objc.Call[objc.Void](s_, objc.Sel("setContentModificationDate:"), objc.Ptr(value))
}

// The name of the apps that converted the original content into a PDF stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621659-encodingapplications?language=objc
func (s_ SearchableItemAttributeSet) EncodingApplications() []string {
	rv := objc.Call[[]string](s_, objc.Sel("encodingApplications"))
	return rv
}

// The name of the apps that converted the original content into a PDF stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621659-encodingapplications?language=objc
func (s_ SearchableItemAttributeSet) SetEncodingApplications(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setEncodingApplications:"), value)
}

// The measurement precision mode in use by the GPS receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620573-gpsmeasuremode?language=objc
func (s_ SearchableItemAttributeSet) GPSMeasureMode() string {
	rv := objc.Call[string](s_, objc.Sel("GPSMeasureMode"))
	return rv
}

// The measurement precision mode in use by the GPS receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620573-gpsmeasuremode?language=objc
func (s_ SearchableItemAttributeSet) SetGPSMeasureMode(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setGPSMeasureMode:"), value)
}

// Indicates the role of the content creator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616105-role?language=objc
func (s_ SearchableItemAttributeSet) Role() string {
	rv := objc.Call[string](s_, objc.Sel("role"))
	return rv
}

// Indicates the role of the content creator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616105-role?language=objc
func (s_ SearchableItemAttributeSet) SetRole(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setRole:"), value)
}

// A value that indicates whether the item contains information sufficient to provide navigation to the location it represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621564-supportsnavigation?language=objc
func (s_ SearchableItemAttributeSet) SupportsNavigation() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("supportsNavigation"))
	return rv
}

// A value that indicates whether the item contains information sufficient to provide navigation to the location it represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621564-supportsnavigation?language=objc
func (s_ SearchableItemAttributeSet) SetSupportsNavigation(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setSupportsNavigation:"), objc.Ptr(value))
}

// The province or state of origin according to guidelines the provider establishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620570-stateorprovince?language=objc
func (s_ SearchableItemAttributeSet) StateOrProvince() string {
	rv := objc.Call[string](s_, objc.Sel("stateOrProvince"))
	return rv
}

// The province or state of origin according to guidelines the provider establishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620570-stateorprovince?language=objc
func (s_ SearchableItemAttributeSet) SetStateOrProvince(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setStateOrProvince:"), value)
}

// An array of CSPerson objects representing the content of the Cc: field in an email message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621664-additionalrecipients?language=objc
func (s_ SearchableItemAttributeSet) AdditionalRecipients() []Person {
	rv := objc.Call[[]Person](s_, objc.Sel("additionalRecipients"))
	return rv
}

// An array of CSPerson objects representing the content of the Cc: field in an email message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621664-additionalrecipients?language=objc
func (s_ SearchableItemAttributeSet) SetAdditionalRecipients(value []IPerson) {
	objc.Call[objc.Void](s_, objc.Sel("setAdditionalRecipients:"), value)
}

// An array of localized strings that represent alternate display names for the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621662-alternatenames?language=objc
func (s_ SearchableItemAttributeSet) AlternateNames() []string {
	rv := objc.Call[[]string](s_, objc.Sel("alternateNames"))
	return rv
}

// An array of localized strings that represent alternate display names for the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621662-alternatenames?language=objc
func (s_ SearchableItemAttributeSet) SetAlternateNames(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setAlternateNames:"), value)
}

// An array of identifiers that corresponds to file representations the delegate provides. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/2867861-providerfiletypeidentifiers?language=objc
func (s_ SearchableItemAttributeSet) ProviderFileTypeIdentifiers() []string {
	rv := objc.Call[[]string](s_, objc.Sel("providerFileTypeIdentifiers"))
	return rv
}

// An array of identifiers that corresponds to file representations the delegate provides. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/2867861-providerfiletypeidentifiers?language=objc
func (s_ SearchableItemAttributeSet) SetProviderFileTypeIdentifiers(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setProviderFileTypeIdentifiers:"), value)
}

// The number of channels in the audio data that the file contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616072-audiochannelcount?language=objc
func (s_ SearchableItemAttributeSet) AudioChannelCount() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("audioChannelCount"))
	return rv
}

// The number of channels in the audio data that the file contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616072-audiochannelcount?language=objc
func (s_ SearchableItemAttributeSet) SetAudioChannelCount(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setAudioChannelCount:"), objc.Ptr(value))
}

// The user-supplied rating of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616011-rating?language=objc
func (s_ SearchableItemAttributeSet) Rating() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("rating"))
	return rv
}

// The user-supplied rating of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616011-rating?language=objc
func (s_ SearchableItemAttributeSet) SetRating(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setRating:"), objc.Ptr(value))
}

// A value that indicates if the content is prepared for streaming. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616054-streamable?language=objc
func (s_ SearchableItemAttributeSet) IsStreamable() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("isStreamable"))
	return rv
}

// A value that indicates if the content is prepared for streaming. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616054-streamable?language=objc
func (s_ SearchableItemAttributeSet) SetStreamable(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setStreamable:"), objc.Ptr(value))
}

// The track number of a song or audio composition when part of an album. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616095-audiotracknumber?language=objc
func (s_ SearchableItemAttributeSet) AudioTrackNumber() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("audioTrackNumber"))
	return rv
}

// The track number of a song or audio composition when part of an album. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616095-audiotracknumber?language=objc
func (s_ SearchableItemAttributeSet) SetAudioTrackNumber(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setAudioTrackNumber:"), objc.Ptr(value))
}

// An attribute type that identifies a custom hierarchy of types to describe the attributes of your item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621660-contenttypetree?language=objc
func (s_ SearchableItemAttributeSet) ContentTypeTree() []string {
	rv := objc.Call[[]string](s_, objc.Sel("contentTypeTree"))
	return rv
}

// An attribute type that identifies a custom hierarchy of types to describe the attributes of your item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621660-contenttypetree?language=objc
func (s_ SearchableItemAttributeSet) SetContentTypeTree(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setContentTypeTree:"), value)
}

// The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616070-musicalgenre?language=objc
func (s_ SearchableItemAttributeSet) MusicalGenre() string {
	rv := objc.Call[string](s_, objc.Sel("musicalGenre"))
	return rv
}

// The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616070-musicalgenre?language=objc
func (s_ SearchableItemAttributeSet) SetMusicalGenre(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setMusicalGenre:"), value)
}

// The status of the GPS receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620589-gpsstatus?language=objc
func (s_ SearchableItemAttributeSet) GPSStatus() string {
	rv := objc.Call[string](s_, objc.Sel("GPSStatus"))
	return rv
}

// The status of the GPS receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620589-gpsstatus?language=objc
func (s_ SearchableItemAttributeSet) SetGPSStatus(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setGPSStatus:"), value)
}

// The genre of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616031-genre?language=objc
func (s_ SearchableItemAttributeSet) Genre() string {
	rv := objc.Call[string](s_, objc.Sel("genre"))
	return rv
}

// The genre of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616031-genre?language=objc
func (s_ SearchableItemAttributeSet) SetGenre(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setGenre:"), value)
}

// The security method (a type of encryption) that protects the document file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621658-securitymethod?language=objc
func (s_ SearchableItemAttributeSet) SecurityMethod() string {
	rv := objc.Call[string](s_, objc.Sel("securityMethod"))
	return rv
}

// The security method (a type of encryption) that protects the document file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621658-securitymethod?language=objc
func (s_ SearchableItemAttributeSet) SetSecurityMethod(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setSecurityMethod:"), value)
}

// The bearing to the destination point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620577-gpsdestbearing?language=objc
func (s_ SearchableItemAttributeSet) GPSDestBearing() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("GPSDestBearing"))
	return rv
}

// The bearing to the destination point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620577-gpsdestbearing?language=objc
func (s_ SearchableItemAttributeSet) SetGPSDestBearing(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setGPSDestBearing:"), objc.Ptr(value))
}

// The manufacturer of the device that captured the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621542-acquisitionmake?language=objc
func (s_ SearchableItemAttributeSet) AcquisitionMake() string {
	rv := objc.Call[string](s_, objc.Sel("acquisitionMake"))
	return rv
}

// The manufacturer of the device that captured the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621542-acquisitionmake?language=objc
func (s_ SearchableItemAttributeSet) SetAcquisitionMake(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setAcquisitionMake:"), value)
}

// A value that indicates the user selected the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/2887562-usercurated?language=objc
func (s_ SearchableItemAttributeSet) IsUserCurated() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("isUserCurated"))
	return rv
}

// A value that indicates the user selected the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/2887562-usercurated?language=objc
func (s_ SearchableItemAttributeSet) SetUserCurated(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setUserCurated:"), objc.Ptr(value))
}

// The codecs used to encode/decode the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616047-codecs?language=objc
func (s_ SearchableItemAttributeSet) Codecs() []string {
	rv := objc.Call[[]string](s_, objc.Sel("codecs"))
	return rv
}

// The codecs used to encode/decode the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616047-codecs?language=objc
func (s_ SearchableItemAttributeSet) SetCodecs(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setCodecs:"), value)
}

// The delivery type of the file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616062-deliverytype?language=objc
func (s_ SearchableItemAttributeSet) DeliveryType() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("deliveryType"))
	return rv
}

// The delivery type of the file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616062-deliverytype?language=objc
func (s_ SearchableItemAttributeSet) SetDeliveryType(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setDeliveryType:"), objc.Ptr(value))
}

// The direction of the item's image in degrees from true north. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620576-imagedirection?language=objc
func (s_ SearchableItemAttributeSet) ImageDirection() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("imageDirection"))
	return rv
}

// The direction of the item's image in degrees from true north. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620576-imagedirection?language=objc
func (s_ SearchableItemAttributeSet) SetImageDirection(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setImageDirection:"), objc.Ptr(value))
}

// The artist associated with the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616077-artist?language=objc
func (s_ SearchableItemAttributeSet) Artist() string {
	rv := objc.Call[string](s_, objc.Sel("artist"))
	return rv
}

// The artist associated with the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616077-artist?language=objc
func (s_ SearchableItemAttributeSet) SetArtist(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setArtist:"), value)
}

// The city of the item’s origin according to guidelines that the provider establishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620594-city?language=objc
func (s_ SearchableItemAttributeSet) City() string {
	rv := objc.Call[string](s_, objc.Sel("city"))
	return rv
}

// The city of the item’s origin according to guidelines that the provider establishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620594-city?language=objc
func (s_ SearchableItemAttributeSet) SetCity(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setCity:"), value)
}

// The time that the lens was open during exposure, in a string, such as "1/250 seconds". [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621521-exposuretimestring?language=objc
func (s_ SearchableItemAttributeSet) ExposureTimeString() string {
	rv := objc.Call[string](s_, objc.Sel("exposureTimeString"))
	return rv
}

// The time that the lens was open during exposure, in a string, such as "1/250 seconds". [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621521-exposuretimestring?language=objc
func (s_ SearchableItemAttributeSet) SetExposureTimeString(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setExposureTimeString:"), value)
}

// An array of identifiers that corresponds to in-place file representations the delegate provides. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/2908990-providerinplacefiletypeidentifie?language=objc
func (s_ SearchableItemAttributeSet) ProviderInPlaceFileTypeIdentifiers() []string {
	rv := objc.Call[[]string](s_, objc.Sel("providerInPlaceFileTypeIdentifiers"))
	return rv
}

// An array of identifiers that corresponds to in-place file representations the delegate provides. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/2908990-providerinplacefiletypeidentifie?language=objc
func (s_ SearchableItemAttributeSet) SetProviderInPlaceFileTypeIdentifiers(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setProviderInPlaceFileTypeIdentifiers:"), value)
}

// The title of the container to which the item belongs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621661-containertitle?language=objc
func (s_ SearchableItemAttributeSet) ContainerTitle() string {
	rv := objc.Call[string](s_, objc.Sel("containerTitle"))
	return rv
}

// The title of the container to which the item belongs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621661-containertitle?language=objc
func (s_ SearchableItemAttributeSet) SetContainerTitle(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setContainerTitle:"), value)
}

// The white balance setting when the camera captured the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621539-whitebalance?language=objc
func (s_ SearchableItemAttributeSet) WhiteBalance() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("whiteBalance"))
	return rv
}

// The white balance setting when the camera captured the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621539-whitebalance?language=objc
func (s_ SearchableItemAttributeSet) SetWhiteBalance(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setWhiteBalance:"), objc.Ptr(value))
}

// The producer of the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616092-producer?language=objc
func (s_ SearchableItemAttributeSet) Producer() string {
	rv := objc.Call[string](s_, objc.Sel("producer"))
	return rv
}

// The producer of the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616092-producer?language=objc
func (s_ SearchableItemAttributeSet) SetProducer(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setProducer:"), value)
}

// A list of people, organizations, or services that made contributions to the media content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616060-contributors?language=objc
func (s_ SearchableItemAttributeSet) Contributors() []string {
	rv := objc.Call[[]string](s_, objc.Sel("contributors"))
	return rv
}

// A list of people, organizations, or services that made contributions to the media content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616060-contributors?language=objc
func (s_ SearchableItemAttributeSet) SetContributors(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setContributors:"), value)
}

// A comment related to the media file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616083-comment?language=objc
func (s_ SearchableItemAttributeSet) Comment() string {
	rv := objc.Call[string](s_, objc.Sel("comment"))
	return rv
}

// A comment related to the media file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616083-comment?language=objc
func (s_ SearchableItemAttributeSet) SetComment(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setComment:"), value)
}

// The direction of travel of the item in degrees from true north. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620588-gpstrack?language=objc
func (s_ SearchableItemAttributeSet) GPSTrack() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("GPSTrack"))
	return rv
}

// The direction of travel of the item in degrees from true north. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620588-gpstrack?language=objc
func (s_ SearchableItemAttributeSet) SetGPSTrack(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setGPSTrack:"), objc.Ptr(value))
}

// A description of the item’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621584-contentdescription?language=objc
func (s_ SearchableItemAttributeSet) ContentDescription() string {
	rv := objc.Call[string](s_, objc.Sel("contentDescription"))
	return rv
}

// A description of the item’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621584-contentdescription?language=objc
func (s_ SearchableItemAttributeSet) SetContentDescription(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setContentDescription:"), value)
}

// A list of companies or organizations that created the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616026-organizations?language=objc
func (s_ SearchableItemAttributeSet) Organizations() []string {
	rv := objc.Call[[]string](s_, objc.Sel("organizations"))
	return rv
}

// A list of companies or organizations that created the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616026-organizations?language=objc
func (s_ SearchableItemAttributeSet) SetOrganizations(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setOrganizations:"), value)
}

// An array of CSPerson objects representing the content of the From: field in an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621608-authors?language=objc
func (s_ SearchableItemAttributeSet) Authors() []Person {
	rv := objc.Call[[]Person](s_, objc.Sel("authors"))
	return rv
}

// An array of CSPerson objects representing the content of the From: field in an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621608-authors?language=objc
func (s_ SearchableItemAttributeSet) SetAuthors(value []IPerson) {
	objc.Call[objc.Void](s_, objc.Sel("setAuthors:"), value)
}

// A list of performers in the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616043-performers?language=objc
func (s_ SearchableItemAttributeSet) Performers() []string {
	rv := objc.Call[[]string](s_, objc.Sel("performers"))
	return rv
}

// A list of performers in the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616043-performers?language=objc
func (s_ SearchableItemAttributeSet) SetPerformers(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setPerformers:"), value)
}

// The speed of the item, in kilometers per hour. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620591-speed?language=objc
func (s_ SearchableItemAttributeSet) Speed() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("speed"))
	return rv
}

// The speed of the item, in kilometers per hour. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620591-speed?language=objc
func (s_ SearchableItemAttributeSet) SetSpeed(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setSpeed:"), objc.Ptr(value))
}

// Instructions that concern the use of the item, such as an embargo or warning. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620590-instructions?language=objc
func (s_ SearchableItemAttributeSet) Instructions() string {
	rv := objc.Call[string](s_, objc.Sel("instructions"))
	return rv
}

// Instructions that concern the use of the item, such as an embargo or warning. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620590-instructions?language=objc
func (s_ SearchableItemAttributeSet) SetInstructions(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setInstructions:"), value)
}

// The number of pages in the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621627-pagecount?language=objc
func (s_ SearchableItemAttributeSet) PageCount() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("pageCount"))
	return rv
}

// The number of pages in the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621627-pagecount?language=objc
func (s_ SearchableItemAttributeSet) SetPageCount(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setPageCount:"), objc.Ptr(value))
}

// The latitude of the destination point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620572-gpsdestlatitude?language=objc
func (s_ SearchableItemAttributeSet) GPSDestLatitude() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("GPSDestLatitude"))
	return rv
}

// The latitude of the destination point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620572-gpsdestlatitude?language=objc
func (s_ SearchableItemAttributeSet) SetGPSDestLatitude(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setGPSDestLatitude:"), objc.Ptr(value))
}

// The focal length of the lens, divided by the diameter of the aperture when the camera captured the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621549-fnumber?language=objc
func (s_ SearchableItemAttributeSet) FNumber() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("fNumber"))
	return rv
}

// The focal length of the lens, divided by the diameter of the aperture when the camera captured the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621549-fnumber?language=objc
func (s_ SearchableItemAttributeSet) SetFNumber(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setFNumber:"), objc.Ptr(value))
}

// The name of the director of the media (for example, a movie director). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616009-director?language=objc
func (s_ SearchableItemAttributeSet) Director() string {
	rv := objc.Call[string](s_, objc.Sel("director"))
	return rv
}

// The name of the director of the media (for example, a movie director). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616009-director?language=objc
func (s_ SearchableItemAttributeSet) SetDirector(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setDirector:"), value)
}

// Information about the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616101-information?language=objc
func (s_ SearchableItemAttributeSet) Information() string {
	rv := objc.Call[string](s_, objc.Sel("information"))
	return rv
}

// Information about the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616101-information?language=objc
func (s_ SearchableItemAttributeSet) SetInformation(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setInformation:"), value)
}

// The date and time related to the GPS value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620587-gpsdatestamp?language=objc
func (s_ SearchableItemAttributeSet) GPSDateStamp() foundation.Date {
	rv := objc.Call[foundation.Date](s_, objc.Sel("GPSDateStamp"))
	return rv
}

// The date and time related to the GPS value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620587-gpsdatestamp?language=objc
func (s_ SearchableItemAttributeSet) SetGPSDateStamp(value foundation.IDate) {
	objc.Call[objc.Void](s_, objc.Sel("setGPSDateStamp:"), objc.Ptr(value))
}

// The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621541-colorspace?language=objc
func (s_ SearchableItemAttributeSet) ColorSpace() string {
	rv := objc.Call[string](s_, objc.Sel("colorSpace"))
	return rv
}

// The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621541-colorspace?language=objc
func (s_ SearchableItemAttributeSet) SetColorSpace(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setColorSpace:"), value)
}

// The total bit rate of the media, combining audio and video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616058-totalbitrate?language=objc
func (s_ SearchableItemAttributeSet) TotalBitRate() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("totalBitRate"))
	return rv
}

// The total bit rate of the media, combining audio and video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616058-totalbitrate?language=objc
func (s_ SearchableItemAttributeSet) SetTotalBitRate(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setTotalBitRate:"), objc.Ptr(value))
}

// An array of the canonical handles for the account with which the message is associated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621567-accounthandles?language=objc
func (s_ SearchableItemAttributeSet) AccountHandles() []string {
	rv := objc.Call[[]string](s_, objc.Sel("accountHandles"))
	return rv
}

// An array of the canonical handles for the account with which the message is associated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621567-accounthandles?language=objc
func (s_ SearchableItemAttributeSet) SetAccountHandles(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setAccountHandles:"), value)
}

// The model of the device that captured the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621523-acquisitionmodel?language=objc
func (s_ SearchableItemAttributeSet) AcquisitionModel() string {
	rv := objc.Call[string](s_, objc.Sel("acquisitionModel"))
	return rv
}

// The model of the device that captured the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621523-acquisitionmodel?language=objc
func (s_ SearchableItemAttributeSet) SetAcquisitionModel(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setAcquisitionModel:"), value)
}

// An array of mailbox identifiers associated with the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621654-mailboxidentifiers?language=objc
func (s_ SearchableItemAttributeSet) MailboxIdentifiers() []string {
	rv := objc.Call[[]string](s_, objc.Sel("mailboxIdentifiers"))
	return rv
}

// An array of mailbox identifiers associated with the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621654-mailboxidentifiers?language=objc
func (s_ SearchableItemAttributeSet) SetMailboxIdentifiers(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setMailboxIdentifiers:"), value)
}

// The ISO speed setting at the time the camera captured the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621548-isospeed?language=objc
func (s_ SearchableItemAttributeSet) ISOSpeed() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("ISOSpeed"))
	return rv
}

// The ISO speed setting at the time the camera captured the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621548-isospeed?language=objc
func (s_ SearchableItemAttributeSet) SetISOSpeed(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setISOSpeed:"), objc.Ptr(value))
}

// The media types present in the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616064-mediatypes?language=objc
func (s_ SearchableItemAttributeSet) MediaTypes() []string {
	rv := objc.Call[[]string](s_, objc.Sel("mediaTypes"))
	return rv
}

// The media types present in the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616064-mediatypes?language=objc
func (s_ SearchableItemAttributeSet) SetMediaTypes(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setMediaTypes:"), value)
}

// The creation date of an edited or optimized version of the song or composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616001-contentcreationdate?language=objc
func (s_ SearchableItemAttributeSet) ContentCreationDate() foundation.Date {
	rv := objc.Call[foundation.Date](s_, objc.Sel("contentCreationDate"))
	return rv
}

// The creation date of an edited or optimized version of the song or composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616001-contentcreationdate?language=objc
func (s_ SearchableItemAttributeSet) SetContentCreationDate(value foundation.IDate) {
	objc.Call[objc.Void](s_, objc.Sel("setContentCreationDate:"), objc.Ptr(value))
}

// The postal code for the item according to guidelines the provider establishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1649284-postalcode?language=objc
func (s_ SearchableItemAttributeSet) PostalCode() string {
	rv := objc.Call[string](s_, objc.Sel("postalCode"))
	return rv
}

// The postal code for the item according to guidelines the provider establishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1649284-postalcode?language=objc
func (s_ SearchableItemAttributeSet) SetPostalCode(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setPostalCode:"), value)
}

// The time signature of the musical composition that the audio or MIDI file contains, in a string, such as "4/4" or "7/8". [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616079-timesignature?language=objc
func (s_ SearchableItemAttributeSet) TimeSignature() string {
	rv := objc.Call[string](s_, objc.Sel("timeSignature"))
	return rv
}

// The time signature of the musical composition that the audio or MIDI file contains, in a string, such as "4/4" or "7/8". [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616079-timesignature?language=objc
func (s_ SearchableItemAttributeSet) SetTimeSignature(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setTimeSignature:"), value)
}

// The local file URL of the thumbnail image for the item when Dark Mode is active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/3752013-darkthumbnailurl?language=objc
func (s_ SearchableItemAttributeSet) DarkThumbnailURL() foundation.URL {
	rv := objc.Call[foundation.URL](s_, objc.Sel("darkThumbnailURL"))
	return rv
}

// The local file URL of the thumbnail image for the item when Dark Mode is active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/3752013-darkthumbnailurl?language=objc
func (s_ SearchableItemAttributeSet) SetDarkThumbnailURL(value foundation.IURL) {
	objc.Call[objc.Void](s_, objc.Sel("setDarkThumbnailURL:"), objc.Ptr(value))
}

// An identifier that represents the domain or owner of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1649285-domainidentifier?language=objc
func (s_ SearchableItemAttributeSet) DomainIdentifier() string {
	rv := objc.Call[string](s_, objc.Sel("domainIdentifier"))
	return rv
}

// An identifier that represents the domain or owner of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1649285-domainidentifier?language=objc
func (s_ SearchableItemAttributeSet) SetDomainIdentifier(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setDomainIdentifier:"), value)
}

// A value that indicates if the camera used a flash to capture the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621525-flashon?language=objc
func (s_ SearchableItemAttributeSet) IsFlashOn() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("isFlashOn"))
	return rv
}

// A value that indicates if the camera used a flash to capture the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621525-flashon?language=objc
func (s_ SearchableItemAttributeSet) SetFlashOn(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setFlashOn:"), objc.Ptr(value))
}

// A value that indicates if the media contains explicit content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616027-contentrating?language=objc
func (s_ SearchableItemAttributeSet) ContentRating() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("contentRating"))
	return rv
}

// A value that indicates if the media contains explicit content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616027-contentrating?language=objc
func (s_ SearchableItemAttributeSet) SetContentRating(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setContentRating:"), objc.Ptr(value))
}

// The altitude of the item in meters above sea level, expressed using the WGS84 datum. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620579-altitude?language=objc
func (s_ SearchableItemAttributeSet) Altitude() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("altitude"))
	return rv
}

// The altitude of the item in meters above sea level, expressed using the WGS84 datum. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620579-altitude?language=objc
func (s_ SearchableItemAttributeSet) SetAltitude(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setAltitude:"), objc.Ptr(value))
}

// A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621622-containerdisplayname?language=objc
func (s_ SearchableItemAttributeSet) ContainerDisplayName() string {
	rv := objc.Call[string](s_, objc.Sel("containerDisplayName"))
	return rv
}

// A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621622-containerdisplayname?language=objc
func (s_ SearchableItemAttributeSet) SetContainerDisplayName(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setContainerDisplayName:"), value)
}

// An array of identifiers that corresponds to data representations the delegate provides. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/2867859-providerdatatypeidentifiers?language=objc
func (s_ SearchableItemAttributeSet) ProviderDataTypeIdentifiers() []string {
	rv := objc.Call[[]string](s_, objc.Sel("providerDataTypeIdentifiers"))
	return rv
}

// An array of identifiers that corresponds to data representations the delegate provides. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/2867859-providerdatatypeidentifiers?language=objc
func (s_ SearchableItemAttributeSet) SetProviderDataTypeIdentifiers(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setProviderDataTypeIdentifiers:"), value)
}

// A list of projects of which this file is a part. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616066-projects?language=objc
func (s_ SearchableItemAttributeSet) Projects() []string {
	rv := objc.Call[[]string](s_, objc.Sel("projects"))
	return rv
}

// A list of projects of which this file is a part. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616066-projects?language=objc
func (s_ SearchableItemAttributeSet) SetProjects(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setProjects:"), value)
}

// The name of the location or point of interest associated with the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620584-namedlocation?language=objc
func (s_ SearchableItemAttributeSet) NamedLocation() string {
	rv := objc.Call[string](s_, objc.Sel("namedLocation"))
	return rv
}

// The name of the location or point of interest associated with the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620584-namedlocation?language=objc
func (s_ SearchableItemAttributeSet) SetNamedLocation(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setNamedLocation:"), value)
}

// The most recent date on which the file was downloaded or received. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616035-downloadeddate?language=objc
func (s_ SearchableItemAttributeSet) DownloadedDate() foundation.Date {
	rv := objc.Call[foundation.Date](s_, objc.Sel("downloadedDate"))
	return rv
}

// The most recent date on which the file was downloaded or received. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616035-downloadeddate?language=objc
func (s_ SearchableItemAttributeSet) SetDownloadedDate(value foundation.IDate) {
	objc.Call[objc.Void](s_, objc.Sel("setDownloadedDate:"), objc.Ptr(value))
}

// An array of names representing the recipients of this message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621667-recipientnames?language=objc
func (s_ SearchableItemAttributeSet) RecipientNames() []string {
	rv := objc.Call[[]string](s_, objc.Sel("recipientNames"))
	return rv
}

// An array of names representing the recipients of this message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621667-recipientnames?language=objc
func (s_ SearchableItemAttributeSet) SetRecipientNames(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setRecipientNames:"), value)
}

// An array of addresses associated with the recipients of the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621648-recipientaddresses?language=objc
func (s_ SearchableItemAttributeSet) RecipientAddresses() []string {
	rv := objc.Call[[]string](s_, objc.Sel("recipientAddresses"))
	return rv
}

// An array of addresses associated with the recipients of the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621648-recipientaddresses?language=objc
func (s_ SearchableItemAttributeSet) SetRecipientAddresses(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setRecipientAddresses:"), value)
}

// The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620571-longitude?language=objc
func (s_ SearchableItemAttributeSet) Longitude() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("longitude"))
	return rv
}

// The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620571-longitude?language=objc
func (s_ SearchableItemAttributeSet) SetLongitude(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setLongitude:"), objc.Ptr(value))
}

// An array of email addresses associated with the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621572-emailaddresses?language=objc
func (s_ SearchableItemAttributeSet) EmailAddresses() []string {
	rv := objc.Call[[]string](s_, objc.Sel("emailAddresses"))
	return rv
}

// An array of email addresses associated with the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621572-emailaddresses?language=objc
func (s_ SearchableItemAttributeSet) SetEmailAddresses(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setEmailAddresses:"), value)
}

// An array of instant message addresses for the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621575-instantmessageaddresses?language=objc
func (s_ SearchableItemAttributeSet) InstantMessageAddresses() []string {
	rv := objc.Call[[]string](s_, objc.Sel("instantMessageAddresses"))
	return rv
}

// An array of instant message addresses for the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621575-instantmessageaddresses?language=objc
func (s_ SearchableItemAttributeSet) SetInstantMessageAddresses(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setInstantMessageAddresses:"), value)
}

// A user-supplied play count for the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616041-playcount?language=objc
func (s_ SearchableItemAttributeSet) PlayCount() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("playCount"))
	return rv
}

// A user-supplied play count for the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616041-playcount?language=objc
func (s_ SearchableItemAttributeSet) SetPlayCount(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setPlayCount:"), objc.Ptr(value))
}

// A value that indicates if the event covers an entire day. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616636-allday?language=objc
func (s_ SearchableItemAttributeSet) AllDay() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("allDay"))
	return rv
}

// A value that indicates if the event covers an entire day. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616636-allday?language=objc
func (s_ SearchableItemAttributeSet) SetAllDay(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setAllDay:"), objc.Ptr(value))
}

// The total number of pixels in the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621517-pixelcount?language=objc
func (s_ SearchableItemAttributeSet) PixelCount() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("pixelCount"))
	return rv
}

// The total number of pixels in the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621517-pixelcount?language=objc
func (s_ SearchableItemAttributeSet) SetPixelCount(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setPixelCount:"), objc.Ptr(value))
}

// The lyricist or text writer for the song or audio composition that the file contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616052-lyricist?language=objc
func (s_ SearchableItemAttributeSet) Lyricist() string {
	rv := objc.Call[string](s_, objc.Sel("lyricist"))
	return rv
}

// The lyricist or text writer for the song or audio composition that the file contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616052-lyricist?language=objc
func (s_ SearchableItemAttributeSet) SetLyricist(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setLyricist:"), value)
}

// The location finding method that the GPS receiver uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620575-gpsprocessingmethod?language=objc
func (s_ SearchableItemAttributeSet) GPSProcessingMethod() string {
	rv := objc.Call[string](s_, objc.Sel("GPSProcessingMethod"))
	return rv
}

// The location finding method that the GPS receiver uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620575-gpsprocessingmethod?language=objc
func (s_ SearchableItemAttributeSet) SetGPSProcessingMethod(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setGPSProcessingMethod:"), value)
}

// The time that the lens was open during exposure, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621532-exposuretime?language=objc
func (s_ SearchableItemAttributeSet) ExposureTime() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("exposureTime"))
	return rv
}

// The time that the lens was open during exposure, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621532-exposuretime?language=objc
func (s_ SearchableItemAttributeSet) SetExposureTime(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setExposureTime:"), objc.Ptr(value))
}

// The URL associated with the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616087-url?language=objc
func (s_ SearchableItemAttributeSet) URL() foundation.URL {
	rv := objc.Call[foundation.URL](s_, objc.Sel("URL"))
	return rv
}

// The URL associated with the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616087-url?language=objc
func (s_ SearchableItemAttributeSet) SetURL(value foundation.IURL) {
	objc.Call[objc.Void](s_, objc.Sel("setURL:"), objc.Ptr(value))
}

// A description of the rating. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616037-ratingdescription?language=objc
func (s_ SearchableItemAttributeSet) RatingDescription() string {
	rv := objc.Call[string](s_, objc.Sel("ratingDescription"))
	return rv
}

// A description of the rating. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616037-ratingdescription?language=objc
func (s_ SearchableItemAttributeSet) SetRatingDescription(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setRatingDescription:"), value)
}

// A description of the kind of document the item represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621642-kind?language=objc
func (s_ SearchableItemAttributeSet) Kind() string {
	rv := objc.Call[string](s_, objc.Sel("kind"))
	return rv
}

// A description of the kind of document the item represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621642-kind?language=objc
func (s_ SearchableItemAttributeSet) SetKind(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setKind:"), value)
}

// The uniform type identifier (UTI) of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621561-contenttype?language=objc
func (s_ SearchableItemAttributeSet) ContentType() string {
	rv := objc.Call[string](s_, objc.Sel("contentType"))
	return rv
}

// The uniform type identifier (UTI) of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621561-contenttype?language=objc
func (s_ SearchableItemAttributeSet) SetContentType(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setContentType:"), value)
}

// The date on which the item is due. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616641-duedate?language=objc
func (s_ SearchableItemAttributeSet) DueDate() foundation.Date {
	rv := objc.Call[foundation.Date](s_, objc.Sel("dueDate"))
	return rv
}

// The date on which the item is due. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616641-duedate?language=objc
func (s_ SearchableItemAttributeSet) SetDueDate(value foundation.IDate) {
	objc.Call[objc.Void](s_, objc.Sel("setDueDate:"), objc.Ptr(value))
}

// The name of an instrument within the context of an instrument category. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616033-musicalinstrumentname?language=objc
func (s_ SearchableItemAttributeSet) MusicalInstrumentName() string {
	rv := objc.Call[string](s_, objc.Sel("musicalInstrumentName"))
	return rv
}

// The name of an instrument within the context of an instrument category. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616033-musicalinstrumentname?language=objc
func (s_ SearchableItemAttributeSet) SetMusicalInstrumentName(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setMusicalInstrumentName:"), value)
}

// A link to information about the rights held in and over the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616076-rights?language=objc
func (s_ SearchableItemAttributeSet) Rights() string {
	rv := objc.Call[string](s_, objc.Sel("rights"))
	return rv
}

// A link to information about the rights held in and over the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616076-rights?language=objc
func (s_ SearchableItemAttributeSet) SetRights(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setRights:"), value)
}

// An array of email addresses associated with the author of the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621625-authoremailaddresses?language=objc
func (s_ SearchableItemAttributeSet) AuthorEmailAddresses() []string {
	rv := objc.Call[[]string](s_, objc.Sel("authorEmailAddresses"))
	return rv
}

// An array of email addresses associated with the author of the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621625-authoremailaddresses?language=objc
func (s_ SearchableItemAttributeSet) SetAuthorEmailAddresses(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setAuthorEmailAddresses:"), value)
}

// The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616008-audiosamplerate?language=objc
func (s_ SearchableItemAttributeSet) AudioSampleRate() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("audioSampleRate"))
	return rv
}

// The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616008-audiosamplerate?language=objc
func (s_ SearchableItemAttributeSet) SetAudioSampleRate(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setAudioSampleRate:"), objc.Ptr(value))
}

// The date on which the last metadata attribute was changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621657-metadatamodificationdate?language=objc
func (s_ SearchableItemAttributeSet) MetadataModificationDate() foundation.Date {
	rv := objc.Call[foundation.Date](s_, objc.Sel("metadataModificationDate"))
	return rv
}

// The date on which the last metadata attribute was changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621657-metadatamodificationdate?language=objc
func (s_ SearchableItemAttributeSet) SetMetadataModificationDate(value foundation.IDate) {
	objc.Call[objc.Void](s_, objc.Sel("setMetadataModificationDate:"), objc.Ptr(value))
}

// An array of CSPerson objects representing the content of the To: field in an email message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621675-primaryrecipients?language=objc
func (s_ SearchableItemAttributeSet) PrimaryRecipients() []Person {
	rv := objc.Call[[]Person](s_, objc.Sel("primaryRecipients"))
	return rv
}

// An array of CSPerson objects representing the content of the To: field in an email message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621675-primaryrecipients?language=objc
func (s_ SearchableItemAttributeSet) SetPrimaryRecipients(value []IPerson) {
	objc.Call[objc.Void](s_, objc.Sel("setPrimaryRecipients:"), value)
}

// The audio bit rate of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616103-audiobitrate?language=objc
func (s_ SearchableItemAttributeSet) AudioBitRate() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("audioBitRate"))
	return rv
}

// The audio bit rate of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616103-audiobitrate?language=objc
func (s_ SearchableItemAttributeSet) SetAudioBitRate(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setAudioBitRate:"), objc.Ptr(value))
}

// A value that indicates if the focal length is 35mm. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621546-focallength35mm?language=objc
func (s_ SearchableItemAttributeSet) IsFocalLength35mm() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("isFocalLength35mm"))
	return rv
}

// A value that indicates if the focal length is 35mm. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621546-focallength35mm?language=objc
func (s_ SearchableItemAttributeSet) SetFocalLength35mm(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setFocalLength35mm:"), objc.Ptr(value))
}

// The version of GPS Info IFD header that was used to generate the metadata for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621543-exifgpsversion?language=objc
func (s_ SearchableItemAttributeSet) EXIFGPSVersion() string {
	rv := objc.Call[string](s_, objc.Sel("EXIFGPSVersion"))
	return rv
}

// The version of GPS Info IFD header that was used to generate the metadata for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621543-exifgpsversion?language=objc
func (s_ SearchableItemAttributeSet) SetEXIFGPSVersion(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setEXIFGPSVersion:"), value)
}

// The smallest F number of the lens. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621545-maxaperture?language=objc
func (s_ SearchableItemAttributeSet) MaxAperture() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("maxAperture"))
	return rv
}

// The smallest F number of the lens. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621545-maxaperture?language=objc
func (s_ SearchableItemAttributeSet) SetMaxAperture(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setMaxAperture:"), objc.Ptr(value))
}

// A value that indicates if the message is likely to be considered junk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621646-likelyjunk?language=objc
func (s_ SearchableItemAttributeSet) IsLikelyJunk() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("isLikelyJunk"))
	return rv
}

// A value that indicates if the message is likely to be considered junk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621646-likelyjunk?language=objc
func (s_ SearchableItemAttributeSet) SetLikelyJunk(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setLikelyJunk:"), objc.Ptr(value))
}

// A list of descriptors that specify the extent or scope of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616006-coverage?language=objc
func (s_ SearchableItemAttributeSet) Coverage() []string {
	rv := objc.Call[[]string](s_, objc.Sel("coverage"))
	return rv
}

// A list of descriptors that specify the extent or scope of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616006-coverage?language=objc
func (s_ SearchableItemAttributeSet) SetCoverage(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setCoverage:"), value)
}

// The start date for the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616640-startdate?language=objc
func (s_ SearchableItemAttributeSet) StartDate() foundation.Date {
	rv := objc.Call[foundation.Date](s_, objc.Sel("startDate"))
	return rv
}

// The start date for the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616640-startdate?language=objc
func (s_ SearchableItemAttributeSet) SetStartDate(value foundation.IDate) {
	objc.Call[objc.Void](s_, objc.Sel("setStartDate:"), objc.Ptr(value))
}

// The size of the document file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621565-filesize?language=objc
func (s_ SearchableItemAttributeSet) FileSize() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("fileSize"))
	return rv
}

// The size of the document file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621565-filesize?language=objc
func (s_ SearchableItemAttributeSet) SetFileSize(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setFileSize:"), objc.Ptr(value))
}

// A value that indicates the user created the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/2887561-usercreated?language=objc
func (s_ SearchableItemAttributeSet) IsUserCreated() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("isUserCreated"))
	return rv
}

// A value that indicates the user created the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/2887561-usercreated?language=objc
func (s_ SearchableItemAttributeSet) SetUserCreated(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setUserCreated:"), objc.Ptr(value))
}

// The title of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621615-title?language=objc
func (s_ SearchableItemAttributeSet) Title() string {
	rv := objc.Call[string](s_, objc.Sel("title"))
	return rv
}

// The title of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621615-title?language=objc
func (s_ SearchableItemAttributeSet) SetTitle(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setTitle:"), value)
}

// The mode the camera used for the exposure of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621540-exposuremode?language=objc
func (s_ SearchableItemAttributeSet) ExposureMode() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("exposureMode"))
	return rv
}

// The mode the camera used for the exposure of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621540-exposuremode?language=objc
func (s_ SearchableItemAttributeSet) SetExposureMode(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setExposureMode:"), objc.Ptr(value))
}

// Image data that represents the thumbnail of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621582-thumbnaildata?language=objc
func (s_ SearchableItemAttributeSet) ThumbnailData() []byte {
	rv := objc.Call[[]byte](s_, objc.Sel("thumbnailData"))
	return rv
}

// Image data that represents the thumbnail of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621582-thumbnaildata?language=objc
func (s_ SearchableItemAttributeSet) SetThumbnailData(value []byte) {
	objc.Call[objc.Void](s_, objc.Sel("setThumbnailData:"), value)
}

// The unique identifier for the item to which the activity is related. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621569-relateduniqueidentifier?language=objc
func (s_ SearchableItemAttributeSet) RelatedUniqueIdentifier() string {
	rv := objc.Call[string](s_, objc.Sel("relatedUniqueIdentifier"))
	return rv
}

// The unique identifier for the item to which the activity is related. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621569-relateduniqueidentifier?language=objc
func (s_ SearchableItemAttributeSet) SetRelatedUniqueIdentifier(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setRelatedUniqueIdentifier:"), value)
}

// A list of people who are visible in an image or movie or written about in a document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616045-participants?language=objc
func (s_ SearchableItemAttributeSet) Participants() []string {
	rv := objc.Call[[]string](s_, objc.Sel("participants"))
	return rv
}

// A list of people who are visible in an image or movie or written about in a document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616045-participants?language=objc
func (s_ SearchableItemAttributeSet) SetParticipants(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setParticipants:"), value)
}

// The unique identifier for the account with which the message is associated, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621573-accountidentifier?language=objc
func (s_ SearchableItemAttributeSet) AccountIdentifier() string {
	rv := objc.Call[string](s_, objc.Sel("accountIdentifier"))
	return rv
}

// The unique identifier for the account with which the message is associated, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621573-accountidentifier?language=objc
func (s_ SearchableItemAttributeSet) SetAccountIdentifier(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setAccountIdentifier:"), value)
}

// The title for a collection of audio media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616016-album?language=objc
func (s_ SearchableItemAttributeSet) Album() string {
	rv := objc.Call[string](s_, objc.Sel("album"))
	return rv
}

// The title for a collection of audio media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616016-album?language=objc
func (s_ SearchableItemAttributeSet) SetAlbum(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setAlbum:"), value)
}

// The original source of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616022-originalsource?language=objc
func (s_ SearchableItemAttributeSet) OriginalSource() string {
	rv := objc.Call[string](s_, objc.Sel("originalSource"))
	return rv
}

// The original source of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616022-originalsource?language=objc
func (s_ SearchableItemAttributeSet) SetOriginalSource(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setOriginalSource:"), value)
}

// A list of editors who have worked on the file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616088-editors?language=objc
func (s_ SearchableItemAttributeSet) Editors() []string {
	rv := objc.Call[[]string](s_, objc.Sel("editors"))
	return rv
}

// A list of editors who have worked on the file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616088-editors?language=objc
func (s_ SearchableItemAttributeSet) SetEditors(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setEditors:"), value)
}

// The subject of the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621598-subject?language=objc
func (s_ SearchableItemAttributeSet) Subject() string {
	rv := objc.Call[string](s_, objc.Sel("subject"))
	return rv
}

// The subject of the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621598-subject?language=objc
func (s_ SearchableItemAttributeSet) SetSubject(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setSubject:"), value)
}

// The timestamp on the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620574-timestamp?language=objc
func (s_ SearchableItemAttributeSet) Timestamp() foundation.Date {
	rv := objc.Call[foundation.Date](s_, objc.Sel("timestamp"))
	return rv
}

// The timestamp on the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620574-timestamp?language=objc
func (s_ SearchableItemAttributeSet) SetTimestamp(value foundation.IDate) {
	objc.Call[objc.Void](s_, objc.Sel("setTimestamp:"), objc.Ptr(value))
}

// The recording date of the song or composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616051-recordingdate?language=objc
func (s_ SearchableItemAttributeSet) RecordingDate() foundation.Date {
	rv := objc.Call[foundation.Date](s_, objc.Sel("recordingDate"))
	return rv
}

// The recording date of the song or composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616051-recordingdate?language=objc
func (s_ SearchableItemAttributeSet) SetRecordingDate(value foundation.IDate) {
	objc.Call[objc.Void](s_, objc.Sel("setRecordingDate:"), objc.Ptr(value))
}

// The resolution height of the image, in DPI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621526-resolutionheightdpi?language=objc
func (s_ SearchableItemAttributeSet) ResolutionHeightDPI() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("resolutionHeightDPI"))
	return rv
}

// The resolution height of the image, in DPI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621526-resolutionheightdpi?language=objc
func (s_ SearchableItemAttributeSet) SetResolutionHeightDPI(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setResolutionHeightDPI:"), objc.Ptr(value))
}

// A list of contacts who are associated with the content in some way, not including the author. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616014-contactkeywords?language=objc
func (s_ SearchableItemAttributeSet) ContactKeywords() []string {
	rv := objc.Call[[]string](s_, objc.Sel("contactKeywords"))
	return rv
}

// A list of contacts who are associated with the content in some way, not including the author. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616014-contactkeywords?language=objc
func (s_ SearchableItemAttributeSet) SetContactKeywords(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setContactKeywords:"), value)
}

// An array that contains the names of the various layers in the file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621519-layernames?language=objc
func (s_ SearchableItemAttributeSet) LayerNames() []string {
	rv := objc.Call[[]string](s_, objc.Sel("layerNames"))
	return rv
}

// An array that contains the names of the various layers in the file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621519-layernames?language=objc
func (s_ SearchableItemAttributeSet) SetLayerNames(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setLayerNames:"), value)
}

// The class of the program the camera used to set exposure when capturing the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621547-exposureprogram?language=objc
func (s_ SearchableItemAttributeSet) ExposureProgram() string {
	rv := objc.Call[string](s_, objc.Sel("exposureProgram"))
	return rv
}

// The class of the program the camera used to set exposure when capturing the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621547-exposureprogram?language=objc
func (s_ SearchableItemAttributeSet) SetExposureProgram(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setExposureProgram:"), value)
}

// The date on which the item was moved into its current location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616029-addeddate?language=objc
func (s_ SearchableItemAttributeSet) AddedDate() foundation.Date {
	rv := objc.Call[foundation.Date](s_, objc.Sel("addedDate"))
	return rv
}

// The date on which the item was moved into its current location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616029-addeddate?language=objc
func (s_ SearchableItemAttributeSet) SetAddedDate(value foundation.IDate) {
	objc.Call[objc.Void](s_, objc.Sel("setAddedDate:"), objc.Ptr(value))
}

// The identifier of the container to which the item belongs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621594-containeridentifier?language=objc
func (s_ SearchableItemAttributeSet) ContainerIdentifier() string {
	rv := objc.Call[string](s_, objc.Sel("containerIdentifier"))
	return rv
}

// The identifier of the container to which the item belongs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621594-containeridentifier?language=objc
func (s_ SearchableItemAttributeSet) SetContainerIdentifier(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setContainerIdentifier:"), value)
}

// The distance to the destination point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620582-gpsdestdistance?language=objc
func (s_ SearchableItemAttributeSet) GPSDestDistance() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("GPSDestDistance"))
	return rv
}

// The distance to the destination point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620582-gpsdestdistance?language=objc
func (s_ SearchableItemAttributeSet) SetGPSDestDistance(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setGPSDestDistance:"), objc.Ptr(value))
}

// The owner of the camera that captured the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621522-cameraowner?language=objc
func (s_ SearchableItemAttributeSet) CameraOwner() string {
	rv := objc.Call[string](s_, objc.Sel("cameraOwner"))
	return rv
}

// The owner of the camera that captured the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621522-cameraowner?language=objc
func (s_ SearchableItemAttributeSet) SetCameraOwner(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setCameraOwner:"), value)
}

// The copyright date of the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616012-copyright?language=objc
func (s_ SearchableItemAttributeSet) Copyright() string {
	rv := objc.Call[string](s_, objc.Sel("copyright"))
	return rv
}

// The copyright date of the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616012-copyright?language=objc
func (s_ SearchableItemAttributeSet) SetCopyright(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setCopyright:"), value)
}

// The metering mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621538-meteringmode?language=objc
func (s_ SearchableItemAttributeSet) MeteringMode() string {
	rv := objc.Call[string](s_, objc.Sel("meteringMode"))
	return rv
}

// The metering mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621538-meteringmode?language=objc
func (s_ SearchableItemAttributeSet) SetMeteringMode(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setMeteringMode:"), value)
}

// The duration (if appropriate) of the content of the file, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616056-duration?language=objc
func (s_ SearchableItemAttributeSet) Duration() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("duration"))
	return rv
}

// The duration (if appropriate) of the content of the file, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616056-duration?language=objc
func (s_ SearchableItemAttributeSet) SetDuration(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setDuration:"), objc.Ptr(value))
}

// The end date for the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616638-enddate?language=objc
func (s_ SearchableItemAttributeSet) EndDate() foundation.Date {
	rv := objc.Call[foundation.Date](s_, objc.Sel("endDate"))
	return rv
}

// The end date for the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616638-enddate?language=objc
func (s_ SearchableItemAttributeSet) SetEndDate(value foundation.IDate) {
	objc.Call[objc.Void](s_, objc.Sel("setEndDate:"), objc.Ptr(value))
}

// The sublocation, such as a street number, for the item according to guidelines the provider establishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1649290-subthoroughfare?language=objc
func (s_ SearchableItemAttributeSet) SubThoroughfare() string {
	rv := objc.Call[string](s_, objc.Sel("subThoroughfare"))
	return rv
}

// The sublocation, such as a street number, for the item according to guidelines the provider establishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1649290-subthoroughfare?language=objc
func (s_ SearchableItemAttributeSet) SetSubThoroughfare(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setSubThoroughfare:"), value)
}

// An array of email addresses associated with the recipient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621580-recipientemailaddresses?language=objc
func (s_ SearchableItemAttributeSet) RecipientEmailAddresses() []string {
	rv := objc.Call[[]string](s_, objc.Sel("recipientEmailAddresses"))
	return rv
}

// An array of email addresses associated with the recipient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621580-recipientemailaddresses?language=objc
func (s_ SearchableItemAttributeSet) SetRecipientEmailAddresses(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setRecipientEmailAddresses:"), value)
}

// The actual focal length of the lens, in millimeters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621533-focallength?language=objc
func (s_ SearchableItemAttributeSet) FocalLength() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("focalLength"))
	return rv
}

// The actual focal length of the lens, in millimeters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621533-focallength?language=objc
func (s_ SearchableItemAttributeSet) SetFocalLength(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setFocalLength:"), objc.Ptr(value))
}

// The longitude of the destination point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620580-gpsdestlongitude?language=objc
func (s_ SearchableItemAttributeSet) GPSDestLongitude() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("GPSDestLongitude"))
	return rv
}

// The longitude of the destination point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620580-gpsdestlongitude?language=objc
func (s_ SearchableItemAttributeSet) SetGPSDestLongitude(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setGPSDestLongitude:"), objc.Ptr(value))
}

// The video bit rate of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616090-videobitrate?language=objc
func (s_ SearchableItemAttributeSet) VideoBitRate() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("videoBitRate"))
	return rv
}

// The video bit rate of the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616090-videobitrate?language=objc
func (s_ SearchableItemAttributeSet) SetVideoBitRate(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setVideoBitRate:"), objc.Ptr(value))
}

// A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621653-supportsphonecall?language=objc
func (s_ SearchableItemAttributeSet) SupportsPhoneCall() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("supportsPhoneCall"))
	return rv
}

// A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621653-supportsphonecall?language=objc
func (s_ SearchableItemAttributeSet) SetSupportsPhoneCall(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setSupportsPhoneCall:"), objc.Ptr(value))
}

// The model of the lens that captured the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621537-lensmodel?language=objc
func (s_ SearchableItemAttributeSet) LensModel() string {
	rv := objc.Call[string](s_, objc.Sel("lensModel"))
	return rv
}

// The model of the lens that captured the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621537-lensmodel?language=objc
func (s_ SearchableItemAttributeSet) SetLensModel(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setLensModel:"), value)
}

// A dictionary that contains all the headers of the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621568-emailheaders?language=objc
func (s_ SearchableItemAttributeSet) EmailHeaders() map[string][]objc.Object {
	rv := objc.Call[map[string][]objc.Object](s_, objc.Sel("emailHeaders"))
	return rv
}

// A dictionary that contains all the headers of the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621568-emailheaders?language=objc
func (s_ SearchableItemAttributeSet) SetEmailHeaders(value map[string][]objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setEmailHeaders:"), value)
}

// A localized string that contains the name of the item, suitable to display in the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621676-displayname?language=objc
func (s_ SearchableItemAttributeSet) DisplayName() string {
	rv := objc.Call[string](s_, objc.Sel("displayName"))
	return rv
}

// A localized string that contains the name of the item, suitable to display in the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621676-displayname?language=objc
func (s_ SearchableItemAttributeSet) SetDisplayName(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setDisplayName:"), value)
}

// The composer of the song or audio composition that the audio file contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616074-composer?language=objc
func (s_ SearchableItemAttributeSet) Composer() string {
	rv := objc.Call[string](s_, objc.Sel("composer"))
	return rv
}

// The composer of the song or audio composition that the audio file contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616074-composer?language=objc
func (s_ SearchableItemAttributeSet) SetComposer(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setComposer:"), value)
}

// The HTML content of the document encoded as an NSData object representing a UTF-8 encoded string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621669-htmlcontentdata?language=objc
func (s_ SearchableItemAttributeSet) HTMLContentData() []byte {
	rv := objc.Call[[]byte](s_, objc.Sel("HTMLContentData"))
	return rv
}

// The HTML content of the document encoded as an NSData object representing a UTF-8 encoded string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621669-htmlcontentdata?language=objc
func (s_ SearchableItemAttributeSet) SetHTMLContentData(value []byte) {
	objc.Call[objc.Void](s_, objc.Sel("setHTMLContentData:"), value)
}

// A formal identifier that references the document the item represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621577-identifier?language=objc
func (s_ SearchableItemAttributeSet) Identifier() string {
	rv := objc.Call[string](s_, objc.Sel("identifier"))
	return rv
}

// A formal identifier that references the document the item represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621577-identifier?language=objc
func (s_ SearchableItemAttributeSet) SetIdentifier(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setIdentifier:"), value)
}

// The number of bits per sample. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621530-bitspersample?language=objc
func (s_ SearchableItemAttributeSet) BitsPerSample() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("bitsPerSample"))
	return rv
}

// The number of bits per sample. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621530-bitspersample?language=objc
func (s_ SearchableItemAttributeSet) SetBitsPerSample(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setBitsPerSample:"), objc.Ptr(value))
}

// An array of keywords associated with the item, such as work, birthday, important, and so on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621652-keywords?language=objc
func (s_ SearchableItemAttributeSet) Keywords() []string {
	rv := objc.Call[[]string](s_, objc.Sel("keywords"))
	return rv
}

// An array of keywords associated with the item, such as work, birthday, important, and so on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621652-keywords?language=objc
func (s_ SearchableItemAttributeSet) SetKeywords(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setKeywords:"), value)
}

// The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620578-country?language=objc
func (s_ SearchableItemAttributeSet) Country() string {
	rv := objc.Call[string](s_, objc.Sel("country"))
	return rv
}

// The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620578-country?language=objc
func (s_ SearchableItemAttributeSet) SetCountry(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setCountry:"), value)
}

// An array of names representing the authors who have worked on the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621620-authornames?language=objc
func (s_ SearchableItemAttributeSet) AuthorNames() []string {
	rv := objc.Call[[]string](s_, objc.Sel("authorNames"))
	return rv
}

// An array of names representing the authors who have worked on the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621620-authornames?language=objc
func (s_ SearchableItemAttributeSet) SetAuthorNames(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setAuthorNames:"), value)
}

// A value that indicates if the media is local. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616049-local?language=objc
func (s_ SearchableItemAttributeSet) IsLocal() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("isLocal"))
	return rv
}

// A value that indicates if the media is local. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616049-local?language=objc
func (s_ SearchableItemAttributeSet) SetLocal(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setLocal:"), objc.Ptr(value))
}

// The tempo of the music that the audio file contains, in beats per minute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616005-tempo?language=objc
func (s_ SearchableItemAttributeSet) Tempo() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("tempo"))
	return rv
}

// The tempo of the music that the audio file contains, in beats per minute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616005-tempo?language=objc
func (s_ SearchableItemAttributeSet) SetTempo(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setTempo:"), objc.Ptr(value))
}

// Information about the GPS area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620581-gpsareainformation?language=objc
func (s_ SearchableItemAttributeSet) GPSAreaInformation() string {
	rv := objc.Call[string](s_, objc.Sel("GPSAreaInformation"))
	return rv
}

// Information about the GPS area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1620581-gpsareainformation?language=objc
func (s_ SearchableItemAttributeSet) SetGPSAreaInformation(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setGPSAreaInformation:"), value)
}

// An array of sources from which the media was obtained. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616024-contentsources?language=objc
func (s_ SearchableItemAttributeSet) ContentSources() []string {
	rv := objc.Call[[]string](s_, objc.Sel("contentSources"))
	return rv
}

// An array of sources from which the media was obtained. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1616024-contentsources?language=objc
func (s_ SearchableItemAttributeSet) SetContentSources(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setContentSources:"), value)
}

// A value that indicates if the camera used red-eye reduction when capturing the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621520-redeyeon?language=objc
func (s_ SearchableItemAttributeSet) IsRedEyeOn() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("isRedEyeOn"))
	return rv
}

// A value that indicates if the camera used red-eye reduction when capturing the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621520-redeyeon?language=objc
func (s_ SearchableItemAttributeSet) SetRedEyeOn(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setRedEyeOn:"), objc.Ptr(value))
}

// The unique identifier for the item to which the activity is related, but not linked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1649297-weakrelateduniqueidentifier?language=objc
func (s_ SearchableItemAttributeSet) WeakRelatedUniqueIdentifier() string {
	rv := objc.Call[string](s_, objc.Sel("weakRelatedUniqueIdentifier"))
	return rv
}

// The unique identifier for the item to which the activity is related, but not linked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1649297-weakrelateduniqueidentifier?language=objc
func (s_ SearchableItemAttributeSet) SetWeakRelatedUniqueIdentifier(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setWeakRelatedUniqueIdentifier:"), value)
}

// The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1649310-thoroughfare?language=objc
func (s_ SearchableItemAttributeSet) Thoroughfare() string {
	rv := objc.Call[string](s_, objc.Sel("thoroughfare"))
	return rv
}

// The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1649310-thoroughfare?language=objc
func (s_ SearchableItemAttributeSet) SetThoroughfare(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setThoroughfare:"), value)
}

// The orientation of the data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621524-orientation?language=objc
func (s_ SearchableItemAttributeSet) Orientation() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("orientation"))
	return rv
}

// The orientation of the data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621524-orientation?language=objc
func (s_ SearchableItemAttributeSet) SetOrientation(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setOrientation:"), objc.Ptr(value))
}

// The width of the document page, in points (72 points per inch). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621670-pagewidth?language=objc
func (s_ SearchableItemAttributeSet) PageWidth() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("pageWidth"))
	return rv
}

// The width of the document page, in points (72 points per inch). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621670-pagewidth?language=objc
func (s_ SearchableItemAttributeSet) SetPageWidth(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setPageWidth:"), objc.Ptr(value))
}

// A value that indicates the user purchased or owns the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/2887564-userowned?language=objc
func (s_ SearchableItemAttributeSet) IsUserOwned() foundation.Number {
	rv := objc.Call[foundation.Number](s_, objc.Sel("isUserOwned"))
	return rv
}

// A value that indicates the user purchased or owns the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/2887564-userowned?language=objc
func (s_ SearchableItemAttributeSet) SetUserOwned(value foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setUserOwned:"), objc.Ptr(value))
}
