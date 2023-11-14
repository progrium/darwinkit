package foundation

// Allows successive elements of a hash table to be returned each time this structure is passed to [foundation/nsnexthashenumeratoritem]. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashenumerator?language=objc
type HashEnumerator struct {
	X_pi uint64
	X_si uint64
	X_bs uintptr
}

// The function pointers used to configure behavior of NSMapTable with respect to key elements within a map table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptablekeycallbacks?language=objc
type MapTableKeyCallBacks struct {
	Hash          uintptr
	IsEqual       uintptr
	Retain        uintptr
	Release       uintptr
	Describe      uintptr
	NotAKeyMarker uintptr
}

// A structure representing a base-10 number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimal?language=objc
type Decimal struct {
	Pad_cgo_0  [4]byte
	X_mantissa [8]uint16
}

// Opaque type containing an endian-independent float value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsswappedfloat?language=objc
type SwappedFloat struct {
	V uint32
}

// Allows successive elements of a map table to be returned each time this structure is passed to [foundation/nsnextmapenumeratorpair]. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmapenumerator?language=objc
type MapEnumerator struct {
	X_pi uint64
	X_si uint64
	X_bs uintptr
}

// A description of the distance between the edges of two rectangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsedgeinsets?language=objc
type EdgeInsets struct {
	Top    float64
	Left   float64
	Bottom float64
	Right  float64
}

// Opaque structure containing endian-independent double value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsswappeddouble?language=objc
type SwappedDouble struct {
	V uint64
}

// A structure that defines the three-by-three matrix that performs an affine transform between two coordinate systems. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransformstruct?language=objc
type AffineTransformStruct struct {
	M11 float64
	M12 float64
	M21 float64
	M22 float64
	TX  float64
	TY  float64
}

// A structure that contains version information about the currently executing operating system, including major, minor, and patch version numbers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperatingsystemversion?language=objc
type OperatingSystemVersion struct {
	MajorVersion int64
	MinorVersion int64
	PatchVersion int64
}

// Defines a structure that contains the function pointers used to configure behavior of NSHashTable with respect to elements within a hash table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtablecallbacks?language=objc
type HashTableCallBacks struct {
	Hash     uintptr
	IsEqual  uintptr
	Retain   uintptr
	Release  uintptr
	Describe uintptr
}

// A structure used to describe a portion of a series, such as characters in a string or objects in an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrange?language=objc
type Range struct {
	Location uint64
	Length   uint64
}

// The function pointers used to configure behavior of NSMapTable with respect to value elements within a map table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptablevaluecallbacks?language=objc
type MapTableValueCallBacks struct {
	Retain   uintptr
	Release  uintptr
	Describe uintptr
}

// TODO (unable to generate):
// NSFastEnumerationState
