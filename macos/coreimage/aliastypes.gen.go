// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
)

// The signature for a block that computes the region of interest (ROI) for a given area of destination image pixels. Core Image calls this block when applying the kernel. You specify this block when using the applyWithExtent:roiCallback:arguments: method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikernelroicallback?language=objc
type KernelROICallback = func(index int, destRect coregraphics.Rect) coregraphics.Rect
