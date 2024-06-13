package avfoundation

// A structure that describes the independent decodability of audio samples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursoraudiodependencyinfo?language=objc
type SampleCursorAudioDependencyInfo struct {
	AudioSampleIsIndependentlyDecodable bool
	AudioSamplePacketRefreshCount       int64
}

// A structure that defines edge processing region widths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avedgewidths?language=objc
type EdgeWidths struct {
	Left   float64
	Top    float64
	Right  float64
	Bottom float64
}

// A value that provides information about a chunk of media samples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursorchunkinfo?language=objc
type SampleCursorChunkInfo struct {
	ChunkSampleCount                  int64
	ChunkHasUniformSampleSizes        bool
	ChunkHasUniformSampleDurations    bool
	ChunkHasUniformFormatDescriptions bool
	Pad_cgo_0                         [5]byte
}

// A structure that defines the origin point for a caption. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionpoint?language=objc
type CaptionPoint struct {
	X CaptionDimension
	Y CaptionDimension
}

// A value for describing dependencies between a media sample and other media samples in the same sample sequence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursordependencyinfo?language=objc
type SampleCursorDependencyInfo struct {
	SampleIndicatesWhetherItHasDependentSamples bool
	SampleHasDependentSamples                   bool
	SampleIndicatesWhetherItDependsOnOthers     bool
	SampleDependsOnOthers                       bool
	SampleIndicatesWhetherItHasRedundantCoding  bool
	SampleHasRedundantCoding                    bool
}

// A structure that defines the height and width of a caption. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionsize?language=objc
type CaptionSize struct {
	Width  CaptionDimension
	Height CaptionDimension
}

// A structure that defines a pixel aspect ratio for a rendering context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avpixelaspectratio?language=objc
type PixelAspectRatio struct {
	HorizontalSpacing int64
	VerticalSpacing   int64
}

// A structure that describes the attributes of media samples to consider when resynchronizing a decoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursorsyncinfo?language=objc
type SampleCursorSyncInfo struct {
	SampleIsFullSync    bool
	SampleIsPartialSync bool
	SampleIsDroppable   bool
}

// A structure that defines a caption dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptiondimension?language=objc
type CaptionDimension struct {
	Value float64
	Units uint64
}

// A structure that indicates the offset and length of storage for a media sample or its chunk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursorstoragerange?language=objc
type SampleCursorStorageRange struct {
	Offset int64
	Length int64
}
