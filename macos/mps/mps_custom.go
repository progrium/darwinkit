package mps

import "unsafe"

// https://developer.apple.com/documentation/metalperformanceshaders/mpscopyallocator?language=objc
// todo: should be a function signature
type CopyAllocator unsafe.Pointer
