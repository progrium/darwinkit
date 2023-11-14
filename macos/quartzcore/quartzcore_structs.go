package quartzcore

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caframeraterange?language=objc
type FrameRateRange struct {
	Minimum   float32
	Maximum   float32
	Preferred float32
}

// The standard transform matrix used throughout Core Animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransform3d?language=objc
type Transform3D struct {
	M11 float64
	M12 float64
	M13 float64
	M14 float64
	M21 float64
	M22 float64
	M23 float64
	M24 float64
	M31 float64
	M32 float64
	M33 float64
	M34 float64
	M41 float64
	M42 float64
	M43 float64
	M44 float64
}
