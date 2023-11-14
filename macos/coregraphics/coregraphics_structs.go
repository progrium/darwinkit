package coregraphics

// A structure that holds a version and two callback functions for drawing a custom pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpatterncallbacks?language=objc
type PatternCallbacks struct {
	Version     uint32
	DrawPattern uintptr
	ReleaseInfo uintptr
}

// A structure that contains pointers to callback functions that manage the copying of data for a data consumer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataconsumercallbacks?language=objc
type DataConsumerCallbacks struct {
	PutBytes        uintptr
	ReleaseConsumer uintptr
}

// A reference to frame update’s metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystreamupdateref?language=objc
type DisplayStreamUpdateRef uintptr

// A profile that specifies how to interpret a color value for display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspaceref?language=objc
type ColorSpaceRef uintptr

// An abstraction for data-reading tasks that eliminates the need to manage a raw memory buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataproviderref?language=objc
type DataProviderRef uintptr

// An opaque data type used to convert PostScript data to PDF data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpsconverterref?language=objc
type PSConverterRef uintptr

// A document that contains PDF (Portable Document Format) drawing information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocumentref?language=objc
type PDFDocumentRef uintptr

// A general facility for defining and using callback functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfunctionref?language=objc
type FunctionRef uintptr

// An object that describes how to convert between color spaces for use by other system services. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorconversioninforef?language=objc
type ColorConversionInfoRef uintptr

// A structure that contains a point in a two-dimensional coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpoint?language=objc
type Point struct {
	X float64
	Y float64
}

// A Quartz 2D drawing environment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextref?language=objc
type ContextRef uintptr

// Defines an opaque type that represents a low-level hardware event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventref?language=objc
type EventRef uintptr

// A mutable graphics path: a mathematical description of shapes or lines to be drawn in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgmutablepathref?language=objc
type MutablePathRef uintptr //*_Ctype_struct_CGPath

// A definition for a smooth transition between colors for drawing radial and axial gradient fills. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggradientref?language=objc
type GradientRef uintptr

// Defines a structure containing pointers to client-defined callback functions that manage the sending of data for a sequential-access data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataprovidersequentialcallbacks?language=objc
type DataProviderSequentialCallbacks struct {
	Version     uint32
	GetBytes    uintptr
	SkipForward uintptr
	Rewind      uintptr
	ReleaseInfo uintptr
}

// A set of character glyphs and layout information for drawing text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfontref?language=objc
type FontRef uintptr

// A structure that contains width and height values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgsize?language=objc
type Size struct {
	Width  float64
	Height float64
}

// A reference to a display stream object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystreamref?language=objc
type DisplayStreamRef uintptr

// A set of components that define a color, with a color space specifying how to interpret them. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorref?language=objc
type ColorRef uintptr

// A definition for a smooth transition between colors, controlled by a custom function you provide, for drawing radial and axial gradient fills. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgshadingref?language=objc
type ShadingRef uintptr

// A bitmap image or image mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimageref?language=objc
type ImageRef uintptr

// An abstraction for data-writing tasks that eliminates the need to manage a raw memory buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataconsumerref?language=objc
type DataConsumerRef uintptr

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolordataformat?language=objc
type ColorDataFormat struct {
	Version            uint32
	Colorspace_info    uintptr
	Bitmap_info        uint32
	Bits_per_component uint64
	Bytes_per_row      uint64
	Intent             uint32
	Decode             *float64
}

// A 2D pattern to be used for drawing graphics paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpatternref?language=objc
type PatternRef uintptr

// The distance, in pixel units, that an onscreen region moves. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgscreenupdatemovedelta?language=objc
type ScreenUpdateMoveDelta struct {
	DX int32
	DY int32
}

// Defines pointers to client-defined callback functions that manage the sending of data for a direct-access data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataproviderdirectcallbacks?language=objc
type DataProviderDirectCallbacks struct {
	Version            uint32
	GetBytePointer     uintptr
	ReleaseBytePointer uintptr
	GetBytesAtPosition uintptr
	ReleaseInfo        uintptr
}

// A reference to a display mode object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymoderef?language=objc
type DisplayModeRef uintptr

// A type that represents a page in a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfpageref?language=objc
type PDFPageRef uintptr

// A structure that contains a two-dimensional vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgvector?language=objc
type Vector struct {
	Dx float64
	Dy float64
}

// An offscreen context for reusing content drawn with Core Graphics. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cglayerref?language=objc
type LayerRef uintptr

// Defines an opaque type that represents the source of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsourceref?language=objc
type EventSourceRef uintptr

// A structure that contains the location and dimensions of a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrect?language=objc
type Rect struct {
	Origin Point
	Size   Size
}

// A structure that contains callbacks needed by a CGFunctionRef object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfunctioncallbacks?language=objc
type FunctionCallbacks struct {
	Version     uint32
	Evaluate    uintptr
	ReleaseInfo uintptr
}

// A data structure that provides information about a path element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathelement?language=objc
type PathElement struct {
	Type   uint32
	Points *Point
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdevicecolor?language=objc
type DeviceColor struct {
	Red   float32
	Green float32
	Blue  float32
}

// Defines the structure used to report information about event taps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventtapinformation?language=objc
type EventTapInformation struct {
	EventTapID         uint32
	TapPoint           uint32
	Options            uint32
	EventsOfInterest   uint64
	TappingProcess     int32
	ProcessBeingTapped int32
	Enabled            bool
	MinUsecLatency     float32
	AvgUsecLatency     float32
	MaxUsecLatency     float32
}

// A structure for holding the callbacks provided when you create a PostScript converter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpsconvertercallbacks?language=objc
type PSConverterCallbacks struct {
	Version       uint32
	BeginDocument uintptr
	EndDocument   uintptr
	BeginPage     uintptr
	EndPage       uintptr
	NoteProgress  uintptr
	NoteMessage   uintptr
	ReleaseInfo   uintptr
}
