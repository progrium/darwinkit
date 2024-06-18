package audiotoolbox

import (
	"github.com/progrium/darwinkit/macos/coreaudiotypes"
	"github.com/progrium/darwinkit/macos/corefoundation"
	"github.com/progrium/darwinkit/macos/coremidi"
)

// A data structure for describing SMPTE (Society of Motion Picture and Television Engineers) time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiofile_smpte_time?language=objc
type File_SMPTE_Time struct {
	MHours                int8
	MMinutes              uint8
	MSeconds              uint8
	MFrames               uint8
	MSubFrameSampleOffset uint32
}

// Used for registering an input callback function with an audio unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/aurendercallbackstruct?language=objc
type RenderCallbackStruct struct {
	InputProc       uintptr
	InputProcRefCon uintptr
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitotherplugindesc?language=objc
type UnitOtherPluginDesc struct {
	Format uint32
	// TODO: Plugin _Ctype_struct_AudioClassDescription
}

// The common header for a render event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/aurendereventheader?language=objc
type RenderEventHeader struct {
	Next            *[296]byte
	EventSampleTime int64
	EventType       uint8
	Reserved        uint8
	Pad_cgo_0       [2]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitpresetmas_settings?language=objc
type UnitPresetMAS_Settings struct {
	ManufacturerID   uint32
	EffectID         uint32
	VariantID        uint32
	SettingsVersion  uint32
	NumberOfSettings uint32
	Settings         [1]UnitPresetMAS_SettingData
}

// A structure specifying the number of leading and trailing empty frames to be inserted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiocodecprimeinfo?language=objc
type CodecPrimeInfo struct {
	LeadingFrames  uint32
	TrailingFrames uint32
}

// Describes an audio unit preset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/aupresetevent?language=objc
type PresetEvent struct {
	Scope   uint32
	Element uint32
	Preset  uintptr
}

// An adjustable audio unit attribute such as volume, pitch, or filter cutoff frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitparameter?language=objc
type UnitParameter struct {
	// TODO: MAudioUnit   *_Ctype_struct_ComponentInstanceRecord
	MParameterID uint32
	MScope       uint32
	MElement     uint32
	Pad_cgo_0    [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auhostidentifier?language=objc
type HostIdentifier struct {
	HostName    corefoundation.StringRef // *_Ctype_struct___CFString
	HostVersion NumVersion
	Pad_cgo_0   [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafpackettableheader?language=objc
type FPacketTableHeader struct {
	MNumberPackets      int64
	MNumberValidFrames  int64
	MPrimingFrames      int32
	MRemainderFrames    int32
	MPacketDescriptions [1]uint8
}

// A list of the audio file regions in a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiofileregionlist?language=objc
type FileRegionList struct {
	TimeType       uint32
	MNumberRegions uint32
	MRegions       [1]FileRegion
}

// A specifier for the [audiotoolbox/audio_format_property_identifier/kaudioformatproperty_formatlist] property, including the codec to use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/extendedaudioformatinfo?language=objc
type ExtendedAudioFormatInfo struct {
	MASBD             coreaudiotypes.StreamBasicDescription // _Ctype_struct_AudioStreamBasicDescription
	MMagicCookie      uintptr
	MMagicCookieSize  uint32
	MClassDescription coreaudiotypes.ClassDescription // _Ctype_struct_AudioClassDescription
}

// A timestamp for scheduled starting of an I/O audio unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiooutputunitstartattimeparams?language=objc
type OutputUnitStartAtTimeParams struct {
	MTimestamp coreaudiotypes.TimeStamp //_Ctype_struct_AudioTimeStamp
	MFlags     uint32
	Pad_cgo_0  [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audistanceattenuationdata?language=objc
type DistanceAttenuationData struct {
	InNumberOfPairs uint32
	// TODO: Pairs           [1]_Ctype_struct___0
}

// The time- and transport-related callback functions for an audio unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/hostcallbackinfo?language=objc
type HostCallbackInfo struct {
	HostUserData            uintptr
	BeatAndTempoProc        uintptr
	MusicalTimeLocationProc uintptr
	TransportStateProc      uintptr
	TransportStateProc2     uintptr
}

// Describes a note-on event with extended parameters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/extendednoteonevent?language=objc
type ExtendedNoteOnEvent struct {
	InstrumentID   uint32
	GroupID        uint32
	Duration       float32
	ExtendedParams MusicDeviceNoteParams
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiopacketrolldistancetranslation?language=objc
type PacketRollDistanceTranslation struct {
	MPacket       int64
	MRollDistance int64
}

// Describes a MIDI system-exclusive (SysEx) message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/midirawdata?language=objc
type RawData struct {
	Length    uint32
	Data      [1]uint8
	Pad_cgo_0 [3]byte
}

// An audio unit source-to-destination connection specification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitconnection?language=objc
type UnitConnection struct {
	// TODO: SourceAudioUnit    *_Ctype_struct_ComponentInstanceRecord
	SourceOutputNumber uint32
	DestInputNumber    uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafstringid?language=objc
type FStringID struct {
	MStringID uint32
	Pad_cgo_0 [8]byte
}

// Describes audio left/right balance and front/back fade values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiobalancefade?language=objc
type BalanceFade struct {
	MLeftRightBalance float32
	MBackFrontFade    float32
	MType             uint32
	MChannelLayout    *coreaudiotypes.ChannelLayout //*_Ctype_struct_AudioChannelLayout
}

// The callback function and custom data for providing input-to-output sample mapping for an audio unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auinputsamplesinoutputcallbackstruct?language=objc
type InputSamplesInOutputCallbackStruct struct {
	InputToOutputCallback uintptr
	UserData              uintptr
}

// Identifying information for an audio component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiocomponentdescription?language=objc
type ComponentDescription struct {
	ComponentType         uint32
	ComponentSubType      uint32
	ComponentManufacturer uint32
	ComponentFlags        uint32
	ComponentFlagsMask    uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafaudiodescription?language=objc
type FAudioDescription struct {
	MSampleRate       float64
	MFormatID         uint32
	MFormatFlags      uint32
	MBytesPerPacket   uint32
	MFramesPerPacket  uint32
	MChannelsPerFrame uint32
	MBitsPerChannel   uint32
}

// A data structure used by the [audiotoolbox/audio_file_properties/kaudiofilepropertypackettoframe] and [audiotoolbox/audio_file_properties/kaudiofilepropertyframetopacket] properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audioframepackettranslation?language=objc
type FramePacketTranslation struct {
	MFrame               int64
	MPacket              int64
	MFrameOffsetInPacket uint32
	Pad_cgo_0            [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafumidchunk?language=objc
type FUMIDChunk struct {
	MBytes [64]uint8
}

// The callback function and custom data for an audio unit that provides MIDI output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/aumidioutputcallbackstruct?language=objc
type MIDIOutputCallbackStruct struct {
	MidiOutputCallback uintptr
	UserData           uintptr
}

// A specifier for the constant[audiotoolbox/audio_file_global_info_propertie/kaudiofileglobalinfo_availablestreamdescriptionsforformat]. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiofiletypeandformatid?language=objc
type FileTypeAndFormatID struct {
	MFileType uint32
	MFormatID uint32
}

// The audio input and output channel capabilities for an audio unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auchannelinfo?language=objc
type ChannelInfo struct {
	InChannels  int16
	OutChannels int16
}

// A structure holding magic cookie information needed by some codecs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiocodecmagiccookieinfo?language=objc
type CodecMagicCookieInfo struct {
	MMagicCookieSize uint32
	MMagicCookie     uintptr
}

// Describes a music track tempo in beats-per-minute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/extendedtempoevent?language=objc
type ExtendedTempoEvent struct {
	Bpm float64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafstrings?language=objc
type FStrings struct {
	MNumEntries uint32
	MStringsIDs [1]FStringID
}

// A string representation of a parameter’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitparameterstringfromvalue?language=objc
type UnitParameterStringFromValue struct {
	InParamID uint32
	InValue   *float32
	OutString corefoundation.StringRef // *_Ctype_struct___CFString
}

// A data structure used by the [audiotoolbox/audio_file_properties/kaudiofilepropertybytetopacket] and [audiotoolbox/audio_file_properties/kaudiofilepropertypackettobyte] properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiobytepackettranslation?language=objc
type BytePacketTranslation struct {
	MByte               int64
	MPacket             int64
	MByteOffsetInPacket uint32
	MFlags              uint32
}

// A component instance, or object, is an audio unit or audio codec. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiocomponentinstance?language=objc
// type ComponentInstance *_Ctype_struct_ComponentInstanceRecord

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/caffileheader?language=objc
type FFileHeader struct {
	MFileType    uint32
	MFileVersion uint16
	MFileFlags   uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafinfostrings?language=objc
type FInfoStrings struct {
	MNumEntries uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafoverviewsample?language=objc
type FOverviewSample struct {
	MMinValue int16
	MMaxValue int16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitparameterinfo?language=objc
type UnitParameterInfo struct {
	Name         [52]int8
	UnitName     corefoundation.StringRef // *_Ctype_struct___CFString
	ClumpID      uint32
	CfNameString corefoundation.StringRef // *_Ctype_struct___CFString
	Unit         uint32
	MinValue     float32
	MaxValue     float32
	DefaultValue float32
	Flags        uint32
	Pad_cgo_0    [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/caclocktime?language=objc
type ClockTime struct {
	Format   uint32
	Reserved uint32
	Time     [24]byte
}

// A structure that contains thread context information for a real-time rendering operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitrendercontext?language=objc
type UnitRenderContext struct {
	// TODO: Workgroup *_Ctype_struct_OS_os_workgroup
	Reserved [6]uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audioindependentpackettranslation?language=objc
type IndependentPacketTranslation struct {
	MPacket                       int64
	MIndependentlyDecodablePacket int64
}

// An event recording the changing of a parameter at a particular host time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/aurecordedparameterevent?language=objc
type RecordedParameterEvent struct {
	HostTime  uint64
	Address   uint64
	Value     float32
	Pad_cgo_0 [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/caf_smpte_time?language=objc
type F_SMPTE_Time struct {
	MHours                int8
	MMinutes              int8
	MSeconds              int8
	MFrames               int8
	MSubFrameSampleOffset uint32
}

// An audio unit’s audio level at a particular frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitfrequencyresponsebin?language=objc
type UnitFrequencyResponseBin struct {
	MFrequency float64
	MMagnitude float64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafregionchunk?language=objc
type FRegionChunk struct {
	TimeType       uint32
	MNumberRegions uint32
	MRegions       [1]FRegion
}

// A list of markers associated with an audio file, including their SMPTE time type, the number of markers, and the markers themselves. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiofilemarkerlist?language=objc
type FileMarkerList struct {
	TimeType       uint32
	MNumberMarkers uint32
	MMarkers       [1]FileMarker
}

// Audio clipping that has occurred in a mixer unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitmeterclipping?language=objc
type UnitMeterClipping struct {
	PeakValueSinceLastCall float32
	SawInfinity            uint8
	SawNotANumber          uint8
	Pad_cgo_0              [2]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafmarker?language=objc
type FMarker struct {
	MType      uint32
	Pad_cgo_0  [8]byte
	MMarkerID  uint32
	MSMPTETime F_SMPTE_Time
	MChannel   uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiofilefdftableextended?language=objc
type FileFDFTableExtended struct {
	MComponentStorage   uintptr
	MReadBytesFDF       uintptr
	MWriteBytesFDF      uintptr
	MReadPacketsFDF     uintptr
	MWritePacketsFDF    uintptr
	MGetPropertyInfoFDF uintptr
	MGetPropertyFDF     uintptr
	MSetPropertyFDF     uintptr
	MCountUserDataFDF   uintptr
	MGetUserDataSizeFDF uintptr
	MGetUserDataFDF     uintptr
	MSetUserDataFDF     uintptr
	MReadPacketDataFDF  uintptr
}

// The name and number of custom Cocoa views for an audio unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitcocoaviewinfo?language=objc
type UnitCocoaViewInfo struct {
	MCocoaAUViewBundleLocation corefoundation.URLRef       // *_Ctype_struct___CFURL
	MCocoaAUViewClass          [1]corefoundation.StringRef // *_Ctype_struct___CFString
}

// An audio file region specifies a segment of audio data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiofileregion?language=objc
type FileRegion struct {
	MRegionID      uint32
	MName          corefoundation.StringRef // *_Ctype_struct___CFString
	MFlags         uint32
	MNumberMarkers uint32
	MMarkers       [1]FileMarker
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiocomponentplugininterface?language=objc
type ComponentPlugInInterface struct {
	Open     uintptr
	Close    uintptr
	Lookup   uintptr
	Reserved uintptr
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/catempomapentry?language=objc
type TempoMapEntry struct {
	Beats    float64
	TempoBPM float64
}

// The name and version of an audio unit’s host application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auhostversionidentifier?language=objc
type HostVersionIdentifier struct {
	HostName    corefoundation.StringRef // *_Ctype_struct___CFString
	HostVersion uint32
	Pad_cgo_0   [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/noteparamscontrolvalue?language=objc
type NoteParamsControlValue struct {
	MID    uint32
	MValue float32
}

// Defines an audio queue buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audioqueuebuffer?language=objc
type QueueBuffer struct {
	MAudioDataBytesCapacity    uint32
	MAudioData                 uintptr
	MAudioDataByteSize         uint32
	MUserData                  uintptr
	MPacketDescriptionCapacity uint32
	MPacketDescriptions        *coreaudiotypes.StreamPacketDescription //*_Ctype_struct_AudioStreamPacketDescription
	MPacketDescriptionCount    uint32
	Pad_cgo_0                  [4]byte
}

// Specifies the current level metering information for one channel of an audio queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audioqueuelevelmeterstate?language=objc
type QueueLevelMeterState struct {
	MAveragePower float32
	MPeakPower    float32
}

// A callback used to provide input to an audio unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/aunoderendercallback?language=objc
type NodeRenderCallback struct {
	DestNode        int32
	DestInputNumber uint32
	Cback           RenderCallbackStruct
}

// An audio unit parameter whose value can change in response to a change in its parent metaparameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audependentparameter?language=objc
type DependentParameter struct {
	MScope       uint32
	MParameterID uint32
}

// Describes a MIDI note. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/midinotemessage?language=objc
type NoteMessage struct {
	Channel         uint8
	Note            uint8
	Velocity        uint8
	ReleaseVelocity uint8
	Duration        float32
}

// Contains information about the number of valid frames in a file and where they begin and end. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiofilepackettableinfo?language=objc
type FilePacketTableInfo struct {
	MNumberValidFrames int64
	MPrimingFrames     int32
	MRemainderFrames   int32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audioqueuechannelassignment?language=objc
type QueueChannelAssignment struct {
	MDeviceUID     corefoundation.StringRef // *_Ctype_struct___CFString
	MChannelNumber uint32
	Pad_cgo_0      [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auparameterautomationevent?language=objc
type ParameterAutomationEvent struct {
	HostTime  uint64
	Address   uint64
	Value     float32
	EventType uint32
	Reserved  uint64
}

// Allows an audio unit host application to tell an audio unit to use a specified buffer for its input callback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitexternalbuffer?language=objc
type UnitExternalBuffer struct {
	Buffer    *uint8
	Size      uint32
	Pad_cgo_0 [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiooutputunitmidicallbacks?language=objc
type OutputUnitMIDICallbacks struct {
	UserData      uintptr
	MIDIEventProc uintptr
	MIDISysExProc uintptr
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auparametermidimapping?language=objc
type ParameterMIDIMapping struct {
	MScope       uint32
	MElement     uint32
	MParameterID uint32
	MFlags       uint32
	MSubRangeMin float32
	MSubRangeMax float32
	MStatus      uint8
	MData1       uint8
	Reserved1    uint8
	Reserved2    uint8
	Reserved3    uint32
}

// A structure that describes a scheduled MIDI event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/aumidievent?language=objc
type MIDIEvent struct {
	Next            *[296]byte
	EventSampleTime int64
	EventType       uint8
	Reserved        uint8
	Length          uint16
	Cable           uint8
	Data            [3]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafregion?language=objc
type FRegion struct {
	MRegionID      uint32
	MFlags         uint32
	MNumberMarkers uint32
	MMarkers       [1]FMarker
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/musicdevicenoteparams?language=objc
type MusicDeviceNoteParams struct {
	ArgCount  uint32
	MPitch    float32
	MVelocity float32
	MControls [1]NoteParamsControlValue
}

// A connection between two node objects in an audio processing graph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitnodeconnection?language=objc
type UnitNodeConnection struct {
	SourceNode         int32
	SourceOutputNumber uint32
	DestNode           int32
	DestInputNumber    uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafmarkerchunk?language=objc
type FMarkerChunk struct {
	TimeType       uint32
	MNumberMarkers uint32
	MMarkers       [1]FMarker
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafpositionpeak?language=objc
type FPositionPeak struct {
	MValue    float32
	Pad_cgo_0 [8]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/scheduledaudiofileregion?language=objc
type ScheduledAudioFileRegion struct {
	MTimeStamp              coreaudiotypes.TimeStamp //_Ctype_struct_AudioTimeStamp
	MCompletionProc         uintptr
	MCompletionProcUserData uintptr
	// TODO: MAudioFile              *_Ctype_struct_OpaqueAudioFileID
	MLoopCount    uint32
	MStartFrame   int64
	MFramesToPlay uint32
	Pad_cgo_0     [4]byte
}

// Annotates a position in an audio file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiofilemarker?language=objc
type FileMarker struct {
	MFramePosition float64
	MName          corefoundation.StringRef // *_Ctype_struct___CFString
	MMarkerID      int32
	MSMPTETime     File_SMPTE_Time
	MType          uint32
	MReserved      uint16
	MChannel       uint16
	Pad_cgo_0      [4]byte
}

// Describes a user-defined event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/musiceventuserdata?language=objc
type MusicEventUserData struct {
	Length    uint32
	Data      [1]uint8
	Pad_cgo_0 [3]byte
}

// A key-value pair that declares an attribute or behavior for an audio unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitproperty?language=objc
type UnitProperty struct {
	// TODO: MAudioUnit  *_Ctype_struct_ComponentInstanceRecord
	MPropertyID uint32
	MScope      uint32
	MElement    uint32
	Pad_cgo_0   [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/aunumversion?language=objc
type NumVersion struct {
	NonRelRev      uint8
	Stage          uint8
	MinorAndBugRev uint8
	MajorRev       uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiofilefdftable?language=objc
type FileFDFTable struct {
	MComponentStorage   uintptr
	MReadBytesFDF       uintptr
	MWriteBytesFDF      uintptr
	MReadPacketsFDF     uintptr
	MWritePacketsFDF    uintptr
	MGetPropertyInfoFDF uintptr
	MGetPropertyFDF     uintptr
	MSetPropertyFDF     uintptr
	MCountUserDataFDF   uintptr
	MGetUserDataSizeFDF uintptr
	MGetUserDataFDF     uintptr
	MSetUserDataFDF     uintptr
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafinstrumentchunk?language=objc
type FInstrumentChunk struct {
	MBaseNote         float32
	MMIDILowNote      uint8
	MMIDIHighNote     uint8
	MMIDILowVelocity  uint8
	MMIDIHighVelocity uint8
	MdBGain           float32
	MStartRegionID    uint32
	MSustainRegionID  uint32
	MReleaseRegionID  uint32
	MInstrumentID     uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitevent?language=objc
type UnitEvent struct {
	MEventType uint32
	Pad_cgo_0  [4]byte
	MArgument  [24]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/caf_uuid_chunkheader?language=objc
type F_UUID_ChunkHeader struct {
	MHeader FChunkHeader
	MUUID   [16]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/scheduledaudioslice?language=objc
type ScheduledAudioSlice struct {
	MTimeStamp              coreaudiotypes.TimeStamp // _Ctype_struct_AudioTimeStamp
	MCompletionProc         uintptr
	MCompletionProcUserData uintptr
	MFlags                  uint32
	MReserved               uint32
	MReserved2              uintptr
	MNumberFrames           uint32
	MBufferList             *coreaudiotypes.BufferList // *_Ctype_struct_AudioBufferList
}

// Specifies an audio queue parameter and associated value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audioqueueparameterevent?language=objc
type QueueParameterEvent struct {
	MID    uint32
	MValue float32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafoverviewchunk?language=objc
type FOverviewChunk struct {
	MEditCount             uint32
	MNumFramesPerOVWSample uint32
	MData                  [1]FOverviewSample
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafdatachunk?language=objc
type FDataChunk struct {
	MEditCount uint32
	MData      [1]uint8
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafchunkheader?language=objc
type FChunkHeader struct {
	MChunkType uint32
	Pad_cgo_0  [8]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/ausamplerinstrumentdata?language=objc
type SamplerInstrumentData struct {
	FileURL        corefoundation.URLRef // *_Ctype_struct___CFURL
	InstrumentType uint8
	BankMSB        uint8
	BankLSB        uint8
	PresetID       uint8
	Pad_cgo_0      [4]byte
}

// Supports control of the looping behavior of a music track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/musictrackloopinfo?language=objc
type MusicTrackLoopInfo struct {
	LoopDuration  float64
	NumberOfLoops int32
	Pad_cgo_0     [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cabarbeattime?language=objc
type BarBeatTime struct {
	Bar            int32
	Beat           uint16
	Subbeat        uint16
	SubbeatDivisor uint16
	Reserved       uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/extendedcontrolevent?language=objc
type ExtendedControlEvent struct {
	GroupID   uint32
	ControlID uint32
	Value     float32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiopacketdependencyinfotranslation?language=objc
type PacketDependencyInfoTranslation struct {
	MPacket                   int64
	MIsIndependentlyDecodable uint32
	MNumberPrerollPackets     uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cafpeakchunk?language=objc
type FPeakChunk struct {
	MEditCount uint32
	MPeaks     [1]FPositionPeak
}

// A scheduled change to an audio unit parameter’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitparameterevent?language=objc
type UnitParameterEvent struct {
	Scope       uint32
	Element     uint32
	Parameter   uint32
	EventType   uint32
	EventValues [16]byte
}

// The suggested update rate and history duration for parameters which have the kAudioUnitParameterFlag_PlotHistory flag set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitparameterhistoryinfo?language=objc
type UnitParameterHistoryInfo struct {
	UpdatesPerSecond         float32
	HistoryDurationInSeconds float32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitpresetmas_settingdata?language=objc
type UnitPresetMAS_SettingData struct {
	IsStockSetting uint32
	SettingID      uint32
	DataLen        uint32
	Data           [1]uint8
	Pad_cgo_0      [3]byte
}

// Audio panning information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiopanninginfo?language=objc
type PanningInfo struct {
	MPanningMode      uint32
	MCoordinateFlags  uint32
	MCoordinates      [3]float32
	MGainScale        float32
	MOutputChannelMap *coreaudiotypes.ChannelLayout //*_Ctype_struct_AudioChannelLayout
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitparametervaluetranslation?language=objc
type UnitParameterValueTranslation struct {
	OtherDesc    UnitOtherPluginDesc
	OtherParamID uint32
	OtherValue   float32
	AuParamID    uint32
	AuValue      float32
}

// Describes the interaction between two node objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/aunodeinteraction?language=objc
type NodeInteraction struct {
	NodeInteractionType uint32
	Pad_cgo_0           [4]byte
	NodeInteraction     [24]byte
}

// A parameter's value based on a string representation of the value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitparametervaluefromstring?language=objc
type UnitParameterValueFromString struct {
	InParamID uint32
	InString  corefoundation.StringRef // *_Ctype_struct___CFString
	OutValue  float32
	Pad_cgo_0 [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/ausamplerbankpresetdata?language=objc
type SamplerBankPresetData struct {
	BankURL   corefoundation.URLRef // *_Ctype_struct___CFURL
	BankMSB   uint8
	BankLSB   uint8
	PresetID  uint8
	Reserved  uint8
	Pad_cgo_0 [4]byte
}

// Specifies priming information for an audio converter, used as a value for the [audiotoolbox/audio_converter_properties/kaudioconverterprimeinfo] property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audioconverterprimeinfo?language=objc
type ConverterPrimeInfo struct {
	LeadingFrames  uint32
	TrailingFrames uint32
}

// A structure that describes a scheduled parameter event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auparameterevent?language=objc
type ParameterEvent struct {
	Next                     *[296]byte
	EventSampleTime          int64
	EventType                uint8
	Reserved                 [3]uint8
	RampDurationSampleFrames uint32
	ParameterAddress         uint64
	Value                    float32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/cametertrackentry?language=objc
type MeterTrackEntry struct {
	Beats      float64
	MeterNumer uint16
	MeterDenom uint16
	Pad_cgo_0  [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/midieventlist?language=objc
type EventList struct {
	Protocol   uint32
	NumPackets uint32
	Packet     [1]coremidi.EventPacket //[1]_Ctype_struct_MIDIEventPacket
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/mixerdistanceparams?language=objc
type MixerDistanceParams struct {
	MReferenceDistance float32
	MMaxDistance       float32
	MMaxAttenuation    float32
}

// Describes a MIDI channel message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/midichannelmessage?language=objc
type ChannelMessage struct {
	Status   uint8
	Data1    uint8
	Data2    uint8
	Reserved uint8
}

// Used to set factory presets for an audio unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/aupreset?language=objc
type Preset struct {
	PresetNumber int32
	PresetName   corefoundation.StringRef // *_Ctype_struct___CFString
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiopacketrangebytecounttranslation?language=objc
type PacketRangeByteCountTranslation struct {
	MPacket              int64
	MPacketCount         int64
	MByteCountUpperBound int64
}

// A specifier for the [audiotoolbox/audio_format_property_identifier/kaudioformatproperty_formatlist] property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audioformatinfo?language=objc
type FormatInfo struct {
	MASBD            coreaudiotypes.StreamBasicDescription //_Ctype_struct_AudioStreamBasicDescription
	MMagicCookie     uintptr
	MMagicCookieSize uint32
	Pad_cgo_0        [4]byte
}

// Describes a MIDI metaevent such as lyric text, time signature, and so on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/midimetaevent?language=objc
type MetaEvent struct {
	MetaEventType uint8
	Unused1       uint8
	Unused2       uint8
	Unused3       uint8
	DataLength    uint32
	Data          [1]uint8
	Pad_cgo_0     [3]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitmidicontrolmapping?language=objc
type UnitMIDIControlMapping struct {
	MidiNRPN    uint16
	MidiControl uint8
	Scope       uint8
	Element     uint32
	Parameter   uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/musicdevicestdnoteparams?language=objc
type MusicDeviceStdNoteParams struct {
	ArgCount  uint32
	MPitch    float32
	MVelocity float32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitparametervaluename?language=objc
type UnitParameterValueName struct {
	InParamID uint32
	InValue   *float32
	OutName   corefoundation.StringRef // *_Ctype_struct___CFString
}

// A short version of the name for an audio unit parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/audiounitparameternameinfo?language=objc
type UnitParameterNameInfo struct {
	InID            uint32
	InDesiredLength int32
	OutName         corefoundation.StringRef // *_Ctype_struct___CFString
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/aumidieventlist?language=objc
type MIDIEventList struct {
	Next            *[296]byte
	EventSampleTime int64
	EventType       uint8
	Reserved        uint8
	Cable           uint8
	Pad_cgo_0       [277]byte
}

// TODO (unable to generate):
// CAFAudioFormatListItem
