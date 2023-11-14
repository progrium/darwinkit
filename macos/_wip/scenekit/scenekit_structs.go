package scenekit

// A representation of a three-component vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/scenekit/scnvector3?language=objc
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

// A representation of a 4 x 4 matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/scenekit/scnmatrix4?language=objc
type Matrix4 struct {
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

// A representation of a four-component vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/scenekit/scnvector4?language=objc
type Vector4 struct {
	X float64
	Y float64
	Z float64
	W float64
}
