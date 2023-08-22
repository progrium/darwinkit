// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

// An animal identifier string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnanimalidentifier?language=objc
type AnimalIdentifier string

const (
	AnimalIdentifierCat AnimalIdentifier = "Cat"
	AnimalIdentifierDog AnimalIdentifier = "Dog"
)

// A type alias for expressing rectangle aspect ratios in Vision. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnaspectratio?language=objc
type AspectRatio float64

// The barcode symbologies that the framework detects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnbarcodesymbology?language=objc
type BarcodeSymbology string

const (
	BarcodeSymbologyAztec                   BarcodeSymbology = "VNBarcodeSymbologyAztec"
	BarcodeSymbologyCodabar                 BarcodeSymbology = "VNBarcodeSymbologyCodabar"
	BarcodeSymbologyCode128                 BarcodeSymbology = "VNBarcodeSymbologyCode128"
	BarcodeSymbologyCode39                  BarcodeSymbology = "VNBarcodeSymbologyCode39"
	BarcodeSymbologyCode39Checksum          BarcodeSymbology = "VNBarcodeSymbologyCode39Checksum"
	BarcodeSymbologyCode39FullASCII         BarcodeSymbology = "VNBarcodeSymbologyCode39FullASCII"
	BarcodeSymbologyCode39FullASCIIChecksum BarcodeSymbology = "VNBarcodeSymbologyCode39FullASCIIChecksum"
	BarcodeSymbologyCode93                  BarcodeSymbology = "VNBarcodeSymbologyCode93"
	BarcodeSymbologyCode93i                 BarcodeSymbology = "VNBarcodeSymbologyCode93i"
	BarcodeSymbologyDataMatrix              BarcodeSymbology = "VNBarcodeSymbologyDataMatrix"
	BarcodeSymbologyEAN13                   BarcodeSymbology = "VNBarcodeSymbologyEAN13"
	BarcodeSymbologyEAN8                    BarcodeSymbology = "VNBarcodeSymbologyEAN8"
	BarcodeSymbologyGS1DataBar              BarcodeSymbology = "VNBarcodeSymbologyGS1DataBar"
	BarcodeSymbologyGS1DataBarExpanded      BarcodeSymbology = "VNBarcodeSymbologyGS1DataBarExpanded"
	BarcodeSymbologyGS1DataBarLimited       BarcodeSymbology = "VNBarcodeSymbologyGS1DataBarLimited"
	BarcodeSymbologyI2of5                   BarcodeSymbology = "VNBarcodeSymbologyI2of5"
	BarcodeSymbologyI2of5Checksum           BarcodeSymbology = "VNBarcodeSymbologyI2of5Checksum"
	BarcodeSymbologyITF14                   BarcodeSymbology = "VNBarcodeSymbologyITF14"
	BarcodeSymbologyMicroPDF417             BarcodeSymbology = "VNBarcodeSymbologyMicroPDF417"
	BarcodeSymbologyMicroQR                 BarcodeSymbology = "VNBarcodeSymbologyMicroQR"
	BarcodeSymbologyPDF417                  BarcodeSymbology = "VNBarcodeSymbologyPDF417"
	BarcodeSymbologyQR                      BarcodeSymbology = "VNBarcodeSymbologyQR"
	BarcodeSymbologyUPCE                    BarcodeSymbology = "VNBarcodeSymbologyUPCE"
)

// Constants that the define the chirality, or handedness, of a pose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnchirality?language=objc
type Chirality int

const (
	ChiralityLeft    Chirality = -1
	ChiralityRight   Chirality = 1
	ChiralityUnknown Chirality = 0
)

// A type alias for the confidence value of an observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnconfidence?language=objc
type Confidence float64

// A typealias for expressing tolerance angles in Vision. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndegrees?language=objc
type Degrees float64

// An enumeration of the type of element in feature print data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnelementtype?language=objc
type ElementType uint

const (
	ElementTypeDouble  ElementType = 2
	ElementTypeFloat   ElementType = 1
	ElementTypeUnknown ElementType = 0
)

// Constants that identify errors from the framework. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnerrorcode?language=objc
type ErrorCode int

const (
	ErrorDataUnavailable     ErrorCode = 17
	ErrorIOError             ErrorCode = 6
	ErrorInternalError       ErrorCode = 9
	ErrorInvalidArgument     ErrorCode = 14
	ErrorInvalidFormat       ErrorCode = 2
	ErrorInvalidImage        ErrorCode = 13
	ErrorInvalidModel        ErrorCode = 15
	ErrorInvalidOperation    ErrorCode = 12
	ErrorInvalidOption       ErrorCode = 5
	ErrorMissingOption       ErrorCode = 7
	ErrorNotImplemented      ErrorCode = 8
	ErrorOK                  ErrorCode = 0
	ErrorOperationFailed     ErrorCode = 3
	ErrorOutOfBoundsError    ErrorCode = 4
	ErrorOutOfMemory         ErrorCode = 10
	ErrorRequestCancelled    ErrorCode = 1
	ErrorTimeStampNotFound   ErrorCode = 18
	ErrorUnknownError        ErrorCode = 11
	ErrorUnsupportedRequest  ErrorCode = 19
	ErrorUnsupportedRevision ErrorCode = 16
)

// The supported optical flow accuracy levels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequestcomputationaccuracy?language=objc
type GenerateOpticalFlowRequestComputationAccuracy uint

const (
	GenerateOpticalFlowRequestComputationAccuracyHigh     GenerateOpticalFlowRequestComputationAccuracy = 2
	GenerateOpticalFlowRequestComputationAccuracyLow      GenerateOpticalFlowRequestComputationAccuracy = 0
	GenerateOpticalFlowRequestComputationAccuracyMedium   GenerateOpticalFlowRequestComputationAccuracy = 1
	GenerateOpticalFlowRequestComputationAccuracyVeryHigh GenerateOpticalFlowRequestComputationAccuracy = 3
)

// Constants that define the levels of quality for a person segmentation request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersonsegmentationrequestqualitylevel?language=objc
type GeneratePersonSegmentationRequestQualityLevel uint

const (
	GeneratePersonSegmentationRequestQualityLevelAccurate GeneratePersonSegmentationRequestQualityLevel = 0
	GeneratePersonSegmentationRequestQualityLevelBalanced GeneratePersonSegmentationRequestQualityLevel = 1
	GeneratePersonSegmentationRequestQualityLevelFast     GeneratePersonSegmentationRequestQualityLevel = 2
)

// The supported joint names for the body pose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservationjointname?language=objc
type HumanBodyPoseObservationJointName RecognizedPointKey

const (
	HumanBodyPoseObservationJointNameLeftAnkle     HumanBodyPoseObservationJointName = "left_foot_joint"
	HumanBodyPoseObservationJointNameLeftEar       HumanBodyPoseObservationJointName = "left_ear_joint"
	HumanBodyPoseObservationJointNameLeftElbow     HumanBodyPoseObservationJointName = "left_forearm_joint"
	HumanBodyPoseObservationJointNameLeftEye       HumanBodyPoseObservationJointName = "left_eye_joint"
	HumanBodyPoseObservationJointNameLeftHip       HumanBodyPoseObservationJointName = "left_upLeg_joint"
	HumanBodyPoseObservationJointNameLeftKnee      HumanBodyPoseObservationJointName = "left_leg_joint"
	HumanBodyPoseObservationJointNameLeftShoulder  HumanBodyPoseObservationJointName = "left_shoulder_1_joint"
	HumanBodyPoseObservationJointNameLeftWrist     HumanBodyPoseObservationJointName = "left_hand_joint"
	HumanBodyPoseObservationJointNameNeck          HumanBodyPoseObservationJointName = "neck_1_joint"
	HumanBodyPoseObservationJointNameNose          HumanBodyPoseObservationJointName = "head_joint"
	HumanBodyPoseObservationJointNameRightAnkle    HumanBodyPoseObservationJointName = "right_foot_joint"
	HumanBodyPoseObservationJointNameRightEar      HumanBodyPoseObservationJointName = "right_ear_joint"
	HumanBodyPoseObservationJointNameRightElbow    HumanBodyPoseObservationJointName = "right_forearm_joint"
	HumanBodyPoseObservationJointNameRightEye      HumanBodyPoseObservationJointName = "right_eye_joint"
	HumanBodyPoseObservationJointNameRightHip      HumanBodyPoseObservationJointName = "right_upLeg_joint"
	HumanBodyPoseObservationJointNameRightKnee     HumanBodyPoseObservationJointName = "right_leg_joint"
	HumanBodyPoseObservationJointNameRightShoulder HumanBodyPoseObservationJointName = "right_shoulder_1_joint"
	HumanBodyPoseObservationJointNameRightWrist    HumanBodyPoseObservationJointName = "right_hand_joint"
	HumanBodyPoseObservationJointNameRoot          HumanBodyPoseObservationJointName = "root"
)

// The supported joint group names for the body pose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservationjointsgroupname?language=objc
type HumanBodyPoseObservationJointsGroupName RecognizedPointGroupKey

const (
	HumanBodyPoseObservationJointsGroupNameAll      HumanBodyPoseObservationJointsGroupName = "VNIPOAll"
	HumanBodyPoseObservationJointsGroupNameFace     HumanBodyPoseObservationJointsGroupName = "VNBLKFACE"
	HumanBodyPoseObservationJointsGroupNameLeftArm  HumanBodyPoseObservationJointsGroupName = "VNBLKLARM"
	HumanBodyPoseObservationJointsGroupNameLeftLeg  HumanBodyPoseObservationJointsGroupName = "VNBLKLLEG"
	HumanBodyPoseObservationJointsGroupNameRightArm HumanBodyPoseObservationJointsGroupName = "VNBLKRARM"
	HumanBodyPoseObservationJointsGroupNameRightLeg HumanBodyPoseObservationJointsGroupName = "VNBLKRLEG"
	HumanBodyPoseObservationJointsGroupNameTorso    HumanBodyPoseObservationJointsGroupName = "VNBLKTORSO"
)

// The supported joint names for the hand pose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanhandposeobservationjointname?language=objc
type HumanHandPoseObservationJointName RecognizedPointKey

const (
	HumanHandPoseObservationJointNameIndexDIP  HumanHandPoseObservationJointName = "VNHLKIDIP"
	HumanHandPoseObservationJointNameIndexMCP  HumanHandPoseObservationJointName = "VNHLKIMCP"
	HumanHandPoseObservationJointNameIndexPIP  HumanHandPoseObservationJointName = "VNHLKIPIP"
	HumanHandPoseObservationJointNameIndexTip  HumanHandPoseObservationJointName = "VNHLKITIP"
	HumanHandPoseObservationJointNameLittleDIP HumanHandPoseObservationJointName = "VNHLKPDIP"
	HumanHandPoseObservationJointNameLittleMCP HumanHandPoseObservationJointName = "VNHLKPMCP"
	HumanHandPoseObservationJointNameLittlePIP HumanHandPoseObservationJointName = "VNHLKPPIP"
	HumanHandPoseObservationJointNameLittleTip HumanHandPoseObservationJointName = "VNHLKPTIP"
	HumanHandPoseObservationJointNameMiddleDIP HumanHandPoseObservationJointName = "VNHLKMDIP"
	HumanHandPoseObservationJointNameMiddleMCP HumanHandPoseObservationJointName = "VNHLKMMCP"
	HumanHandPoseObservationJointNameMiddlePIP HumanHandPoseObservationJointName = "VNHLKMPIP"
	HumanHandPoseObservationJointNameMiddleTip HumanHandPoseObservationJointName = "VNHLKMTIP"
	HumanHandPoseObservationJointNameRingDIP   HumanHandPoseObservationJointName = "VNHLKRDIP"
	HumanHandPoseObservationJointNameRingMCP   HumanHandPoseObservationJointName = "VNHLKRMCP"
	HumanHandPoseObservationJointNameRingPIP   HumanHandPoseObservationJointName = "VNHLKRPIP"
	HumanHandPoseObservationJointNameRingTip   HumanHandPoseObservationJointName = "VNHLKRTIP"
	HumanHandPoseObservationJointNameThumbCMC  HumanHandPoseObservationJointName = "VNHLKTCMC"
	HumanHandPoseObservationJointNameThumbIP   HumanHandPoseObservationJointName = "VNHLKTIP"
	HumanHandPoseObservationJointNameThumbMP   HumanHandPoseObservationJointName = "VNHLKTMP"
	HumanHandPoseObservationJointNameThumbTip  HumanHandPoseObservationJointName = "VNHLKTTIP"
	HumanHandPoseObservationJointNameWrist     HumanHandPoseObservationJointName = "VNHLKWRI"
)

// The supported joint group names for the hand pose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanhandposeobservationjointsgroupname?language=objc
type HumanHandPoseObservationJointsGroupName RecognizedPointGroupKey

const (
	HumanHandPoseObservationJointsGroupNameAll          HumanHandPoseObservationJointsGroupName = "VNIPOAll"
	HumanHandPoseObservationJointsGroupNameIndexFinger  HumanHandPoseObservationJointsGroupName = "VNHLRKI"
	HumanHandPoseObservationJointsGroupNameLittleFinger HumanHandPoseObservationJointsGroupName = "VNHLRKP"
	HumanHandPoseObservationJointsGroupNameMiddleFinger HumanHandPoseObservationJointsGroupName = "VNHLRKM"
	HumanHandPoseObservationJointsGroupNameRingFinger   HumanHandPoseObservationJointsGroupName = "VNHLRKR"
	HumanHandPoseObservationJointsGroupNameThumb        HumanHandPoseObservationJointsGroupName = "VNHLRKT"
)

// Options that define how Vision crops and scales an input-image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagecropandscaleoption?language=objc
type ImageCropAndScaleOption uint

const (
	ImageCropAndScaleOptionCenterCrop ImageCropAndScaleOption = 0
	ImageCropAndScaleOptionScaleFill  ImageCropAndScaleOption = 2
	ImageCropAndScaleOptionScaleFit   ImageCropAndScaleOption = 1
)

// An option key passed into VNImageRequestHandler creations or requests that take an auxiliary image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimageoption?language=objc
type ImageOption string

const (
	ImageOptionCIContext        ImageOption = "VNImageOptionCIContext"
	ImageOptionCameraIntrinsics ImageOption = "VNImageOptionCameraIntrinsics"
	ImageOptionProperties       ImageOption = "VNImageOptionProperties"
)

// The data type for all recognized point group keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpointgroupkey?language=objc
type RecognizedPointGroupKey string

const (
	BodyLandmarkRegionKeyFace     RecognizedPointGroupKey = "VNBLKFACE"
	BodyLandmarkRegionKeyLeftArm  RecognizedPointGroupKey = "VNBLKLARM"
	BodyLandmarkRegionKeyLeftLeg  RecognizedPointGroupKey = "VNBLKLLEG"
	BodyLandmarkRegionKeyRightArm RecognizedPointGroupKey = "VNBLKRARM"
	BodyLandmarkRegionKeyRightLeg RecognizedPointGroupKey = "VNBLKRLEG"
	BodyLandmarkRegionKeyTorso    RecognizedPointGroupKey = "VNBLKTORSO"
	RecognizedPointGroupKeyAll    RecognizedPointGroupKey = "VNIPOAll"
)

// The data type for all recognized point keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpointkey?language=objc
type RecognizedPointKey string

const (
	BodyLandmarkKeyLeftAnkle     RecognizedPointKey = "left_foot_joint"
	BodyLandmarkKeyLeftEar       RecognizedPointKey = "left_ear_joint"
	BodyLandmarkKeyLeftElbow     RecognizedPointKey = "left_forearm_joint"
	BodyLandmarkKeyLeftEye       RecognizedPointKey = "left_eye_joint"
	BodyLandmarkKeyLeftHip       RecognizedPointKey = "left_upLeg_joint"
	BodyLandmarkKeyLeftKnee      RecognizedPointKey = "left_leg_joint"
	BodyLandmarkKeyLeftShoulder  RecognizedPointKey = "left_shoulder_1_joint"
	BodyLandmarkKeyLeftWrist     RecognizedPointKey = "left_hand_joint"
	BodyLandmarkKeyNeck          RecognizedPointKey = "neck_1_joint"
	BodyLandmarkKeyNose          RecognizedPointKey = "head_joint"
	BodyLandmarkKeyRightAnkle    RecognizedPointKey = "right_foot_joint"
	BodyLandmarkKeyRightEar      RecognizedPointKey = "right_ear_joint"
	BodyLandmarkKeyRightElbow    RecognizedPointKey = "right_forearm_joint"
	BodyLandmarkKeyRightEye      RecognizedPointKey = "right_eye_joint"
	BodyLandmarkKeyRightHip      RecognizedPointKey = "right_upLeg_joint"
	BodyLandmarkKeyRightKnee     RecognizedPointKey = "right_leg_joint"
	BodyLandmarkKeyRightShoulder RecognizedPointKey = "right_shoulder_1_joint"
	BodyLandmarkKeyRightWrist    RecognizedPointKey = "right_hand_joint"
	BodyLandmarkKeyRoot          RecognizedPointKey = "root"
)

// An enumeration of face landmarks in a constellation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequestfacelandmarksconstellation?language=objc
type RequestFaceLandmarksConstellation uint

const (
	RequestFaceLandmarksConstellation65Points   RequestFaceLandmarksConstellation = 1
	RequestFaceLandmarksConstellation76Points   RequestFaceLandmarksConstellation = 2
	RequestFaceLandmarksConstellationNotDefined RequestFaceLandmarksConstellation = 0
)

// Constants that identify the performance and accuracy of the text recognition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequesttextrecognitionlevel?language=objc
type RequestTextRecognitionLevel int

const (
	RequestTextRecognitionLevelAccurate RequestTextRecognitionLevel = 0
	RequestTextRecognitionLevelFast     RequestTextRecognitionLevel = 1
)

// An enumeration of tracking priorities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequesttrackinglevel?language=objc
type RequestTrackingLevel uint

const (
	RequestTrackingLevelAccurate RequestTrackingLevel = 0
	RequestTrackingLevelFast     RequestTrackingLevel = 1
)

// Options to pass to the video processor when adding requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessingoption?language=objc
type VideoProcessingOption string

const (
	VideoProcessingOptionFrameCadence VideoProcessingOption = "VNVideoProcessingOptionFrameCadence"
	VideoProcessingOptionTimeInterval VideoProcessingOption = "VNVideoProcessingOptionTimeInterval"
)
