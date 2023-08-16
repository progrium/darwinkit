// AUTO-GENERATED CODE, DO NOT MODIFY

package imageio

// Constants that indicate the result of animating an image sequence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/imageio/cgimageanimationstatus?language=objc
type ImageAnimationStatus uint

const (
	KImageAnimationStatus_AllocationFailure    ImageAnimationStatus = 22144
	KImageAnimationStatus_CorruptInputImage    ImageAnimationStatus = 22141
	KImageAnimationStatus_IncompleteInputImage ImageAnimationStatus = 22143
	KImageAnimationStatus_ParameterError       ImageAnimationStatus = 22140
	KImageAnimationStatus_UnsupportedFormat    ImageAnimationStatus = 22142
)

// Constants for errors that occur when getting or setting metadata information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/imageio/cgimagemetadataerrors?language=objc
type ImageMetadataErrors int32

const (
	KImageMetadataErrorBadArgument          ImageMetadataErrors = 2
	KImageMetadataErrorConflictingArguments ImageMetadataErrors = 3
	KImageMetadataErrorPrefixConflict       ImageMetadataErrors = 4
	KImageMetadataErrorUnknown              ImageMetadataErrors = 0
	KImageMetadataErrorUnsupportedFormat    ImageMetadataErrors = 1
)

// Constants that indicate the XMP type for a metadata tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/imageio/cgimagemetadatatype?language=objc
type ImageMetadataType int32

const (
	KImageMetadataTypeAlternateArray ImageMetadataType = 4
	KImageMetadataTypeAlternateText  ImageMetadataType = 5
	KImageMetadataTypeArrayOrdered   ImageMetadataType = 3
	KImageMetadataTypeArrayUnordered ImageMetadataType = 2
	KImageMetadataTypeDefault        ImageMetadataType = 0
	KImageMetadataTypeInvalid        ImageMetadataType = -1
	KImageMetadataTypeString         ImageMetadataType = 1
	KImageMetadataTypeStructure      ImageMetadataType = 6
)

// A value describing the intended display orientation for an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/imageio/cgimagepropertyorientation?language=objc
type ImagePropertyOrientation uint32

const (
	KImagePropertyOrientationDown          ImagePropertyOrientation = 3
	KImagePropertyOrientationDownMirrored  ImagePropertyOrientation = 4
	KImagePropertyOrientationLeft          ImagePropertyOrientation = 8
	KImagePropertyOrientationLeftMirrored  ImagePropertyOrientation = 5
	KImagePropertyOrientationRight         ImagePropertyOrientation = 6
	KImagePropertyOrientationRightMirrored ImagePropertyOrientation = 7
	KImagePropertyOrientationUp            ImagePropertyOrientation = 1
	KImagePropertyOrientationUpMirrored    ImagePropertyOrientation = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/imageio/cgimagepropertytgacompression?language=objc
type ImagePropertyTGACompression uint32

const (
	KImageTGACompressionNone ImagePropertyTGACompression = 0
	KImageTGACompressionRLE  ImagePropertyTGACompression = 1
)

// The set of status values for images and image sources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/imageio/cgimagesourcestatus?language=objc
type ImageSourceStatus int32

const (
	KImageStatusComplete      ImageSourceStatus = 0
	KImageStatusIncomplete    ImageSourceStatus = -1
	KImageStatusInvalidData   ImageSourceStatus = -4
	KImageStatusReadingHeader ImageSourceStatus = -2
	KImageStatusUnexpectedEOF ImageSourceStatus = -5
	KImageStatusUnknownType   ImageSourceStatus = -3
)
