// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetWriterInput] class.
var AssetWriterInputClass = _AssetWriterInputClass{objc.GetClass("AVAssetWriterInput")}

type _AssetWriterInputClass struct {
	objc.Class
}

// An interface definition for the [AssetWriterInput] class.
type IAssetWriterInput interface {
	objc.IObject
	AddTrackAssociationWithTrackOfInputType(input IAssetWriterInput, trackAssociationType string)
	MarkCurrentPassAsFinished()
	RespondToEachPassDescriptionOnQueueUsingBlock(queue dispatch.Queue, block func())
	RequestMediaDataWhenReadyOnQueueUsingBlock(queue dispatch.Queue, block func())
	AppendSampleBuffer(sampleBuffer coremedia.SampleBufferRef) bool
	CanAddTrackAssociationWithTrackOfInputType(input IAssetWriterInput, trackAssociationType string) bool
	MarkAsFinished()
	CurrentPassDescription() AssetWriterInputPassDescription
	IsReadyForMoreMediaData() bool
	ExtendedLanguageTag() string
	SetExtendedLanguageTag(value string)
	PerformsMultiPassEncodingIfSupported() bool
	SetPerformsMultiPassEncodingIfSupported(value bool)
	OutputSettings() map[string]objc.Object
	PreferredMediaChunkAlignment() int
	SetPreferredMediaChunkAlignment(value int)
	PreferredVolume() float64
	SetPreferredVolume(value float64)
	Metadata() []MetadataItem
	SetMetadata(value []IMetadataItem)
	ExpectsMediaDataInRealTime() bool
	SetExpectsMediaDataInRealTime(value bool)
	MediaType() MediaType
	SampleReferenceBaseURL() foundation.URL
	SetSampleReferenceBaseURL(value foundation.IURL)
	MarksOutputTrackAsEnabled() bool
	SetMarksOutputTrackAsEnabled(value bool)
	NaturalSize() coregraphics.Size
	SetNaturalSize(value coregraphics.Size)
	CanPerformMultiplePasses() bool
	MediaTimeScale() coremedia.TimeScale
	SetMediaTimeScale(value coremedia.TimeScale)
	PreferredMediaChunkDuration() coremedia.Time
	SetPreferredMediaChunkDuration(value coremedia.Time)
	Transform() coregraphics.AffineTransform
	SetTransform(value coregraphics.AffineTransform)
	LanguageCode() string
	SetLanguageCode(value string)
	SourceFormatHint() coremedia.FormatDescriptionRef
	MediaDataLocation() AssetWriterInputMediaDataLocation
	SetMediaDataLocation(value AssetWriterInputMediaDataLocation)
}

// An object that appends media samples to a track in an asset writer’s output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput?language=objc
type AssetWriterInput struct {
	objc.Object
}

func AssetWriterInputFrom(ptr unsafe.Pointer) AssetWriterInput {
	return AssetWriterInput{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetWriterInputClass) AssetWriterInputWithMediaTypeOutputSettingsSourceFormatHint(mediaType MediaType, outputSettings map[string]objc.IObject, sourceFormatHint coremedia.FormatDescriptionRef) AssetWriterInput {
	rv := objc.Call[AssetWriterInput](ac, objc.Sel("assetWriterInputWithMediaType:outputSettings:sourceFormatHint:"), mediaType, outputSettings, sourceFormatHint)
	return rv
}

// Returns a new input that appends sample buffers of the specified type and format hint to the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1449091-assetwriterinputwithmediatype?language=objc
func AssetWriterInput_AssetWriterInputWithMediaTypeOutputSettingsSourceFormatHint(mediaType MediaType, outputSettings map[string]objc.IObject, sourceFormatHint coremedia.FormatDescriptionRef) AssetWriterInput {
	return AssetWriterInputClass.AssetWriterInputWithMediaTypeOutputSettingsSourceFormatHint(mediaType, outputSettings, sourceFormatHint)
}

func (a_ AssetWriterInput) InitWithMediaTypeOutputSettingsSourceFormatHint(mediaType MediaType, outputSettings map[string]objc.IObject, sourceFormatHint coremedia.FormatDescriptionRef) AssetWriterInput {
	rv := objc.Call[AssetWriterInput](a_, objc.Sel("initWithMediaType:outputSettings:sourceFormatHint:"), mediaType, outputSettings, sourceFormatHint)
	return rv
}

// Creates an input that appends sample buffers of the specified type and format hint to the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1389994-initwithmediatype?language=objc
func NewAssetWriterInputWithMediaTypeOutputSettingsSourceFormatHint(mediaType MediaType, outputSettings map[string]objc.IObject, sourceFormatHint coremedia.FormatDescriptionRef) AssetWriterInput {
	instance := AssetWriterInputClass.Alloc().InitWithMediaTypeOutputSettingsSourceFormatHint(mediaType, outputSettings, sourceFormatHint)
	instance.Autorelease()
	return instance
}

func (ac _AssetWriterInputClass) Alloc() AssetWriterInput {
	rv := objc.Call[AssetWriterInput](ac, objc.Sel("alloc"))
	return rv
}

func AssetWriterInput_Alloc() AssetWriterInput {
	return AssetWriterInputClass.Alloc()
}

func (ac _AssetWriterInputClass) New() AssetWriterInput {
	rv := objc.Call[AssetWriterInput](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetWriterInput() AssetWriterInput {
	return AssetWriterInputClass.New()
}

func (a_ AssetWriterInput) Init() AssetWriterInput {
	rv := objc.Call[AssetWriterInput](a_, objc.Sel("init"))
	return rv
}

// Adds an association between input tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1388347-addtrackassociationwithtrackofin?language=objc
func (a_ AssetWriterInput) AddTrackAssociationWithTrackOfInputType(input IAssetWriterInput, trackAssociationType string) {
	objc.Call[objc.Void](a_, objc.Sel("addTrackAssociationWithTrackOfInput:type:"), objc.Ptr(input), trackAssociationType)
}

// Tells the input to analyze the appended media to determine whether it can improve the results by reencoding certain segments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1389652-markcurrentpassasfinished?language=objc
func (a_ AssetWriterInput) MarkCurrentPassAsFinished() {
	objc.Call[objc.Void](a_, objc.Sel("markCurrentPassAsFinished"))
}

// Tells the input to invoke a callback whenever it begins a new pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1388489-respondtoeachpassdescriptiononqu?language=objc
func (a_ AssetWriterInput) RespondToEachPassDescriptionOnQueueUsingBlock(queue dispatch.Queue, block func()) {
	objc.Call[objc.Void](a_, objc.Sel("respondToEachPassDescriptionOnQueue:usingBlock:"), queue, block)
}

// Tells the input to request media data, at its convenience, to write to the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1387508-requestmediadatawhenreadyonqueue?language=objc
func (a_ AssetWriterInput) RequestMediaDataWhenReadyOnQueueUsingBlock(queue dispatch.Queue, block func()) {
	objc.Call[objc.Void](a_, objc.Sel("requestMediaDataWhenReadyOnQueue:usingBlock:"), queue, block)
}

// Appends a sample buffer to an input to write to the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1389566-appendsamplebuffer?language=objc
func (a_ AssetWriterInput) AppendSampleBuffer(sampleBuffer coremedia.SampleBufferRef) bool {
	rv := objc.Call[bool](a_, objc.Sel("appendSampleBuffer:"), sampleBuffer)
	return rv
}

// Determines whether it’s valid to associate another input’s track with this input’s track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1388292-canaddtrackassociationwithtracko?language=objc
func (a_ AssetWriterInput) CanAddTrackAssociationWithTrackOfInputType(input IAssetWriterInput, trackAssociationType string) bool {
	rv := objc.Call[bool](a_, objc.Sel("canAddTrackAssociationWithTrackOfInput:type:"), objc.Ptr(input), trackAssociationType)
	return rv
}

// Marks the input as finished to indicate that you’re done appending samples to it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1390122-markasfinished?language=objc
func (a_ AssetWriterInput) MarkAsFinished() {
	objc.Call[objc.Void](a_, objc.Sel("markAsFinished"))
}

// An object that describes the requirements for the current pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1390627-currentpassdescription?language=objc
func (a_ AssetWriterInput) CurrentPassDescription() AssetWriterInputPassDescription {
	rv := objc.Call[AssetWriterInputPassDescription](a_, objc.Sel("currentPassDescription"))
	return rv
}

// A Boolean value that indicates whether the input is ready to accept media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1389084-readyformoremediadata?language=objc
func (a_ AssetWriterInput) IsReadyForMoreMediaData() bool {
	rv := objc.Call[bool](a_, objc.Sel("isReadyForMoreMediaData"))
	return rv
}

// The extended language for the input’s track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1390768-extendedlanguagetag?language=objc
func (a_ AssetWriterInput) ExtendedLanguageTag() string {
	rv := objc.Call[string](a_, objc.Sel("extendedLanguageTag"))
	return rv
}

// The extended language for the input’s track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1390768-extendedlanguagetag?language=objc
func (a_ AssetWriterInput) SetExtendedLanguageTag(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setExtendedLanguageTag:"), value)
}

// A Boolean value that indicates whether the input attempts to encode the source media data using multiple passes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1386570-performsmultipassencodingifsuppo?language=objc
func (a_ AssetWriterInput) PerformsMultiPassEncodingIfSupported() bool {
	rv := objc.Call[bool](a_, objc.Sel("performsMultiPassEncodingIfSupported"))
	return rv
}

// A Boolean value that indicates whether the input attempts to encode the source media data using multiple passes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1386570-performsmultipassencodingifsuppo?language=objc
func (a_ AssetWriterInput) SetPerformsMultiPassEncodingIfSupported(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setPerformsMultiPassEncodingIfSupported:"), value)
}

// The settings to use for encoding media data you append to the output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1388406-outputsettings?language=objc
func (a_ AssetWriterInput) OutputSettings() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](a_, objc.Sel("outputSettings"))
	return rv
}

// The boundary, in bytes, for aligning media chunks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1388163-preferredmediachunkalignment?language=objc
func (a_ AssetWriterInput) PreferredMediaChunkAlignment() int {
	rv := objc.Call[int](a_, objc.Sel("preferredMediaChunkAlignment"))
	return rv
}

// The boundary, in bytes, for aligning media chunks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1388163-preferredmediachunkalignment?language=objc
func (a_ AssetWriterInput) SetPreferredMediaChunkAlignment(value int) {
	objc.Call[objc.Void](a_, objc.Sel("setPreferredMediaChunkAlignment:"), value)
}

// The volume to prefer for playback of the output’s audio data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1389949-preferredvolume?language=objc
func (a_ AssetWriterInput) PreferredVolume() float64 {
	rv := objc.Call[float64](a_, objc.Sel("preferredVolume"))
	return rv
}

// The volume to prefer for playback of the output’s audio data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1389949-preferredvolume?language=objc
func (a_ AssetWriterInput) SetPreferredVolume(value float64) {
	objc.Call[objc.Void](a_, objc.Sel("setPreferredVolume:"), value)
}

// The track-level metadata to write to the output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1386328-metadata?language=objc
func (a_ AssetWriterInput) Metadata() []MetadataItem {
	rv := objc.Call[[]MetadataItem](a_, objc.Sel("metadata"))
	return rv
}

// The track-level metadata to write to the output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1386328-metadata?language=objc
func (a_ AssetWriterInput) SetMetadata(value []IMetadataItem) {
	objc.Call[objc.Void](a_, objc.Sel("setMetadata:"), value)
}

// A Boolean value that indicates whether the input tailors its processing for real-time sources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1387827-expectsmediadatainrealtime?language=objc
func (a_ AssetWriterInput) ExpectsMediaDataInRealTime() bool {
	rv := objc.Call[bool](a_, objc.Sel("expectsMediaDataInRealTime"))
	return rv
}

// A Boolean value that indicates whether the input tailors its processing for real-time sources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1387827-expectsmediadatainrealtime?language=objc
func (a_ AssetWriterInput) SetExpectsMediaDataInRealTime(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setExpectsMediaDataInRealTime:"), value)
}

// The media type of the samples that the input accepts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1385565-mediatype?language=objc
func (a_ AssetWriterInput) MediaType() MediaType {
	rv := objc.Call[MediaType](a_, objc.Sel("mediaType"))
	return rv
}

// The base URL sample references are relative to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1386316-samplereferencebaseurl?language=objc
func (a_ AssetWriterInput) SampleReferenceBaseURL() foundation.URL {
	rv := objc.Call[foundation.URL](a_, objc.Sel("sampleReferenceBaseURL"))
	return rv
}

// The base URL sample references are relative to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1386316-samplereferencebaseurl?language=objc
func (a_ AssetWriterInput) SetSampleReferenceBaseURL(value foundation.IURL) {
	objc.Call[objc.Void](a_, objc.Sel("setSampleReferenceBaseURL:"), objc.Ptr(value))
}

// A Boolean value that indicates whether to enable a track in the output for playback and processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1386764-marksoutputtrackasenabled?language=objc
func (a_ AssetWriterInput) MarksOutputTrackAsEnabled() bool {
	rv := objc.Call[bool](a_, objc.Sel("marksOutputTrackAsEnabled"))
	return rv
}

// A Boolean value that indicates whether to enable a track in the output for playback and processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1386764-marksoutputtrackasenabled?language=objc
func (a_ AssetWriterInput) SetMarksOutputTrackAsEnabled(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setMarksOutputTrackAsEnabled:"), value)
}

// The natural display dimensions of the output’s visual media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1387437-naturalsize?language=objc
func (a_ AssetWriterInput) NaturalSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](a_, objc.Sel("naturalSize"))
	return rv
}

// The natural display dimensions of the output’s visual media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1387437-naturalsize?language=objc
func (a_ AssetWriterInput) SetNaturalSize(value coregraphics.Size) {
	objc.Call[objc.Void](a_, objc.Sel("setNaturalSize:"), value)
}

// A Boolean value that indicates whether the input may perform multiple passes over appended media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1388284-canperformmultiplepasses?language=objc
func (a_ AssetWriterInput) CanPerformMultiplePasses() bool {
	rv := objc.Call[bool](a_, objc.Sel("canPerformMultiplePasses"))
	return rv
}

// The time scale of the track in the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1386902-mediatimescale?language=objc
func (a_ AssetWriterInput) MediaTimeScale() coremedia.TimeScale {
	rv := objc.Call[coremedia.TimeScale](a_, objc.Sel("mediaTimeScale"))
	return rv
}

// The time scale of the track in the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1386902-mediatimescale?language=objc
func (a_ AssetWriterInput) SetMediaTimeScale(value coremedia.TimeScale) {
	objc.Call[objc.Void](a_, objc.Sel("setMediaTimeScale:"), value)
}

// The duration to use for each chunk of sample data in the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1390463-preferredmediachunkduration?language=objc
func (a_ AssetWriterInput) PreferredMediaChunkDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("preferredMediaChunkDuration"))
	return rv
}

// The duration to use for each chunk of sample data in the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1390463-preferredmediachunkduration?language=objc
func (a_ AssetWriterInput) SetPreferredMediaChunkDuration(value coremedia.Time) {
	objc.Call[objc.Void](a_, objc.Sel("setPreferredMediaChunkDuration:"), value)
}

// The transform to use for display of the output’s visual media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1390183-transform?language=objc
func (a_ AssetWriterInput) Transform() coregraphics.AffineTransform {
	rv := objc.Call[coregraphics.AffineTransform](a_, objc.Sel("transform"))
	return rv
}

// The transform to use for display of the output’s visual media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1390183-transform?language=objc
func (a_ AssetWriterInput) SetTransform(value coregraphics.AffineTransform) {
	objc.Call[objc.Void](a_, objc.Sel("setTransform:"), value)
}

// The language code of the input’s track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1388507-languagecode?language=objc
func (a_ AssetWriterInput) LanguageCode() string {
	rv := objc.Call[string](a_, objc.Sel("languageCode"))
	return rv
}

// The language code of the input’s track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1388507-languagecode?language=objc
func (a_ AssetWriterInput) SetLanguageCode(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setLanguageCode:"), value)
}

// A hint about the format of the sample buffers to append to the input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/1387647-sourceformathint?language=objc
func (a_ AssetWriterInput) SourceFormatHint() coremedia.FormatDescriptionRef {
	rv := objc.Call[coremedia.FormatDescriptionRef](a_, objc.Sel("sourceFormatHint"))
	return rv
}

// Specifies how the input lays out and interleaves media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/2867633-mediadatalocation?language=objc
func (a_ AssetWriterInput) MediaDataLocation() AssetWriterInputMediaDataLocation {
	rv := objc.Call[AssetWriterInputMediaDataLocation](a_, objc.Sel("mediaDataLocation"))
	return rv
}

// Specifies how the input lays out and interleaves media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/2867633-mediadatalocation?language=objc
func (a_ AssetWriterInput) SetMediaDataLocation(value AssetWriterInputMediaDataLocation) {
	objc.Call[objc.Void](a_, objc.Sel("setMediaDataLocation:"), value)
}
