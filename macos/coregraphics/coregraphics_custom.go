package coregraphics

// An affine transformation matrix for use in drawing 2D graphics. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cgaffinetransform?language=objc
type AffineTransform struct {
	M11 float64
	M12 float64
	M21 float64
	M22 float64
	TX  float64
	TY  float64
}
