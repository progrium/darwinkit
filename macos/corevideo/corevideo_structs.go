package corevideo

// A reference to a pixel buffer pool object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvpixelbufferpoolref?language=objc
type PixelBufferPoolRef uintptr

// A structure for describing YCbCr biplanar buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvplanarpixelbufferinfo_ycbcrbiplanar?language=objc
type PlanarPixelBufferInfo_YCbCrBiPlanar struct {
	ComponentInfoY    PlanarComponentInfo
	ComponentInfoCbCr PlanarComponentInfo
}

// A structure for holding an SMPTE time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvsmptetime?language=objc
type SMPTETime struct {
	Subframes       int16
	SubframeDivisor int16
	Counter         uint32
	Type            uint32
	Flags           uint32
	Hours           int16
	Minutes         int16
	Seconds         int16
	Frames          int16
}

// A structure for describing planar buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvplanarpixelbufferinfo?language=objc
type PlanarPixelBufferInfo struct {
	ComponentInfo [1]PlanarComponentInfo
}

// A reference to a Core Video buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvbufferref?language=objc
type BufferRef uintptr //*_Ctype_struct___CVBuffer

// A structure for holding information that describes a custom extended pixel fill algorithm. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvfillextendedpixelscallbackdata?language=objc
type FillExtendedPixelsCallBackData struct {
	Version      int64
	FillCallBack uintptr
	RefCon       uintptr
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvopengltexturecacheref?language=objc
type OpenGLTextureCacheRef uintptr

// A structure for defining a display timestamp. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvtimestamp?language=objc
type TimeStamp struct {
	Version            uint32
	VideoTimeScale     int32
	VideoTime          int64
	HostTime           uint64
	RateScalar         float64
	VideoRefreshPeriod int64
	SmpteTime          SMPTETime
	Flags              uint64
	Reserved           uint64
}

// A reference to a Core Video Metal texture cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvmetaltexturecacheref?language=objc
type MetalTextureCacheRef uintptr

// A structure for describing YCbCr planar buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvplanarpixelbufferinfo_ycbcrplanar?language=objc
type PlanarPixelBufferInfo_YCbCrPlanar struct {
	ComponentInfoY  PlanarComponentInfo
	ComponentInfoCb PlanarComponentInfo
	ComponentInfoCr PlanarComponentInfo
}

// A reference to a display link object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvdisplaylinkref?language=objc
type DisplayLinkRef uintptr

// A structure for reporting Core Video time values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvtime?language=objc
type Time struct {
	TimeValue int64
	TimeScale int32
	Flags     int32
}

// A reference to an OpenGL buffer pool object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvopenglbufferpoolref?language=objc
type OpenGLBufferPoolRef uintptr

// A structure for describing planar components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvplanarcomponentinfo?language=objc
type PlanarComponentInfo struct {
	Offset   int32
	RowBytes uint32
}
