// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

// The set of processing-unit configurations the model can use to make predictions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlcomputeunits?language=objc
type ComputeUnits int

const (
	ComputeUnitsAll       ComputeUnits = 2
	ComputeUnitsCPUAndGPU ComputeUnits = 1
	ComputeUnitsCPUOnly   ComputeUnits = 0
)

// The possible types for feature values, input features, and output features. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturetype?language=objc
type FeatureType int

const (
	FeatureTypeDictionary FeatureType = 6
	FeatureTypeDouble     FeatureType = 2
	FeatureTypeImage      FeatureType = 4
	FeatureTypeInt64      FeatureType = 1
	FeatureTypeInvalid    FeatureType = 0
	FeatureTypeMultiArray FeatureType = 5
	FeatureTypeSequence   FeatureType = 7
	FeatureTypeString     FeatureType = 3
)

// The initializer options you use to crop and scale an image when creating an image feature value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalueimageoption?language=objc
type FeatureValueImageOption string

const (
	FeatureValueImageOptionCropAndScale FeatureValueImageOption = "MLFeatureValueImageOptionCropAndScale"
	FeatureValueImageOptionCropRect     FeatureValueImageOption = "MLFeatureValueImageOptionCropRect"
)

// The modes that determine how the model defines a feature's image size constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesizeconstrainttype?language=objc
type ImageSizeConstraintType int

const (
	ImageSizeConstraintTypeEnumerated  ImageSizeConstraintType = 2
	ImageSizeConstraintTypeRange       ImageSizeConstraintType = 3
	ImageSizeConstraintTypeUnspecified ImageSizeConstraintType = 0
)

// The categories of model-specific errors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelerror?language=objc
type ModelError int

const (
	ModelErrorCustomLayer             ModelError = 4
	ModelErrorCustomModel             ModelError = 5
	ModelErrorFeatureType             ModelError = 1
	ModelErrorGeneric                 ModelError = 0
	ModelErrorIO                      ModelError = 3
	ModelErrorModelCollection         ModelError = 10
	ModelErrorModelDecryption         ModelError = 9
	ModelErrorModelDecryptionKeyFetch ModelError = 8
	ModelErrorParameters              ModelError = 7
	ModelErrorUpdate                  ModelError = 6
)

// The set of keys the model uses to store values in its metadata dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelmetadatakey?language=objc
type ModelMetadataKey string

const (
	ModelAuthorKey         ModelMetadataKey = "MLModelAuthorKey"
	ModelCreatorDefinedKey ModelMetadataKey = "MLModelCreatorDefinedKey"
	ModelDescriptionKey    ModelMetadataKey = "MLModelDescriptionKey"
	ModelLicenseKey        ModelMetadataKey = "MLModelLicenseKey"
	ModelVersionStringKey  ModelMetadataKey = "MLModelVersionStringKey"
)

// Constants that define the underlying element types a multiarray can store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarraydatatype?language=objc
type MultiArrayDataType int

const (
	MultiArrayDataTypeDouble  MultiArrayDataType = 65600
	MultiArrayDataTypeFloat   MultiArrayDataType = 65568
	MultiArrayDataTypeFloat16 MultiArrayDataType = 65552
	MultiArrayDataTypeFloat32 MultiArrayDataType = 65568
	MultiArrayDataTypeFloat64 MultiArrayDataType = 65600
	MultiArrayDataTypeInt32   MultiArrayDataType = 131104
)

// The possible types of shape constraints. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayshapeconstrainttype?language=objc
type MultiArrayShapeConstraintType int

const (
	MultiArrayShapeConstraintTypeEnumerated  MultiArrayShapeConstraintType = 2
	MultiArrayShapeConstraintTypeRange       MultiArrayShapeConstraintType = 3
	MultiArrayShapeConstraintTypeUnspecified MultiArrayShapeConstraintType = 1
)

// The state of a machine learning task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mltaskstate?language=objc
type TaskState int

const (
	TaskStateCancelling TaskState = 3
	TaskStateCompleted  TaskState = 4
	TaskStateFailed     TaskState = 5
	TaskStateRunning    TaskState = 2
	TaskStateSuspended  TaskState = 1
)

// A type of event during a model update task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdateprogressevent?language=objc
type UpdateProgressEvent int

const (
	UpdateProgressEventEpochEnd      UpdateProgressEvent = 2
	UpdateProgressEventMiniBatchEnd  UpdateProgressEvent = 4
	UpdateProgressEventTrainingBegin UpdateProgressEvent = 1
)
