package avfaudio

// Priming information for audio conversion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverterprimeinfo?language=objc
type AudioConverterPrimeInfo struct {
	LeadingFrames  uint32
	TrailingFrames uint32
}

// A structure that represents a point in 3D space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudio3dpoint?language=objc
type Audio3DPoint struct {
	X float32
	Y float32
	Z float32
}

// A structure that represents two orthogonal vectors that describe the orientation of the listener in 3D space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudio3dvectororientation?language=objc
type Audio3DVectorOrientation struct {
	Forward Audio3DPoint
	Up      Audio3DPoint
}

// A structure that represents the angular orientation of the listener in 3D space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudio3dangularorientation?language=objc
type Audio3DAngularOrientation struct {
	Yaw   float32
	Pitch float32
	Roll  float32
}

// A specific time range within a music track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avbeatrange?language=objc
type BeatRange struct {
	Start  float64
	Length float64
}
