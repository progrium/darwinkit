package gameplaykit

// The definition of an axis-aligned rectangular bounding volume addressed by the tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkbox?language=objc
type Box struct {
	BoxMin    [3]float32
	Pad_cgo_0 [4]byte
	BoxMax    [3]float32
	Pad_cgo_1 [4]byte
}

// The definition of an axis-aligned rectangle addressed by the tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkquad?language=objc
type Quad struct {
	QuadMin [2]float32
	QuadMax [2]float32
}

// The definition of a triangle in the mesh, available with the [gameplaykit/gkmeshgraph/triangleatindex] method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gktriangle?language=objc
type Triangle struct {
	Points    [3][3]float32
	Pad_cgo_0 [12]byte
}
