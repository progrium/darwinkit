package kernel

import "unsafe"

type Vector_int2 = [2]int32
type Vector_int4 = [4]int32

type Vector_float2 = [2]float32
type Vector_float3 = [3]float32
type Vector_float4 = [4]float32

type Vector_double2 = [2]float64
type Vector_double3 = [3]float64
type Vector_double4 = [4]float64

type Vector_uchar16 = [16]uint8

// not sure these will work, placeholder for now
type Matrix_float2x2 unsafe.Pointer
type Matrix_float3x3 unsafe.Pointer
type Matrix_float4x4 unsafe.Pointer
type Matrix_float4x3 unsafe.Pointer
type Matrix_double4x4 unsafe.Pointer

// https://developer.apple.com/documentation/kernel/kernel_data_types?language=objc
// https://github.com/apple/darwin-xnu/blob/main/bsd/sys/_types.h
type Pid = int32
type Gid = uint32
type Uid = uint32
