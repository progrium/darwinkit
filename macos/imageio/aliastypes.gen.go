// AUTO-GENERATED CODE, DO NOT MODIFY

package imageio

import (
	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/macos/coregraphics"
)

// The block to execute when enumerating the tags of a metadata object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/imageio/cgimagemetadatatagblock?language=objc
type ImageMetadataTagBlock = func(path corefoundation.StringRef, tag ImageMetadataTagRef) bool

// The block to execute for each frame of an image animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/imageio/cgimagesourceanimationblock?language=objc
type ImageSourceAnimationBlock = func(index uint, image coregraphics.ImageRef, stop *bool)
