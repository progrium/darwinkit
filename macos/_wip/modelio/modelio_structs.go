package modelio 

// The minimal volume containing an object, used by the [modelio/mdlobject/boundingboxattime] method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/modelio/mdlaxisalignedboundingbox?language=objc
type AxisAlignedBoundingBox struct {
	MaxBounds	[3]float32
	Pad_cgo_0	[4]byte
	MinBounds	[3]float32
	Pad_cgo_1	[4]byte
}


// The corner voxel indices defining a solid rectangular volume of voxels. Used by the [modelio/mdlvoxelarray/voxelindexextent] property and [modelio/mdlvoxelarray/voxelswithinextent] method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/modelio/mdlvoxelindexextent?language=objc
type VoxelIndexExtent struct {
	MinimumExtent	[4]int32
	MaximumExtent	[4]int32
}
