package coreaudiotypes

// A structure that represents a timestamp value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiotimestamp?language=objc
type TimeStamp struct {
	MSampleTime    float64
	MHostTime      uint64
	MRateScalar    float64
	MWordClockTime uint64
	MSMPTETime     SMPTETime
	MFlags         uint32
	MReserved      uint32
}

// A structure that represents a continuous range of values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiovaluerange?language=objc
type ValueRange struct {
	MMinimum float64
	MMaximum float64
}

// A structure that describes a channel of audio data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiochanneldescription?language=objc
type ChannelDescription struct {
	MChannelLabel uint32
	MChannelFlags uint32
	MCoordinates  [3]float32
}

// Represents a value from the [audiotoolbox/audio_format_property_identifier/kaudioformatproperty_formatlist] property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audioformatlistitem?language=objc
type FormatListItem struct {
	MASBD             StreamBasicDescription
	MChannelLayoutTag uint32
	Pad_cgo_0         [4]byte
}

// A structure that stores buffers to use in translation operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiovaluetranslation?language=objc
type ValueTranslation struct {
	MInputData      uintptr
	MInputDataSize  uint32
	MOutputData     uintptr
	MOutputDataSize uint32
	Pad_cgo_0       [4]byte
}

// A format specification for an audio stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiostreambasicdescription?language=objc
type StreamBasicDescription struct {
	MSampleRate       float64
	MFormatID         uint32
	MFormatFlags      uint32
	MBytesPerPacket   uint32
	MFramesPerPacket  uint32
	MBytesPerFrame    uint32
	MChannelsPerFrame uint32
	MBitsPerChannel   uint32
	MReserved         uint32
}

// A value that describes a packet in a buffer of audio data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiostreampacketdescription?language=objc
type StreamPacketDescription struct {
	MStartOffset            int64
	MVariableFramesInPacket uint32
	MDataByteSize           uint32
}

// A structure that holds a buffer of audio data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiobuffer?language=objc
type Buffer struct {
	MNumberChannels uint32
	MDataByteSize   uint32
	MData           uintptr
}

// A structure that describes an audio codec. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audioclassdescription?language=objc
type ClassDescription struct {
	MType         uint32
	MSubType      uint32
	MManufacturer uint32
}

// A structure that stores a variable-length array of audio buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiobufferlist?language=objc
type BufferList struct {
	MNumberBuffers uint32
	MBuffers       [1]Buffer
}

// A structure that specifies a channel layout in a file or in hardware. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiochannellayout?language=objc
type ChannelLayout struct {
	MChannelLayoutTag          uint32
	MChannelBitmap             uint32
	MNumberChannelDescriptions uint32
	MChannelDescriptions       [1]ChannelDescription
}

// A structure that defines an SMPTE time value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/smptetime?language=objc
type SMPTETime struct {
	MSubframes       int16
	MSubframeDivisor int16
	MCounter         uint32
	MType            uint32
	MFlags           uint32
	MHours           int16
	MMinutes         int16
	MSeconds         int16
	MFrames          int16
}
