// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MappingModel] class.
var MappingModelClass = _MappingModelClass{objc.GetClass("NSMappingModel")}

type _MappingModelClass struct {
	objc.Class
}

// An interface definition for the [MappingModel] class.
type IMappingModel interface {
	objc.IObject
	EntityMappingsByName() map[string]EntityMapping
	EntityMappings() []EntityMapping
	SetEntityMappings(value []IEntityMapping)
}

// A model instance that specifies how to map a model from a source to a destination managed object model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmappingmodel?language=objc
type MappingModel struct {
	objc.Object
}

func MappingModelFrom(ptr unsafe.Pointer) MappingModel {
	return MappingModel{
		Object: objc.ObjectFrom(ptr),
	}
}

func (m_ MappingModel) InitWithContentsOfURL(url foundation.IURL) MappingModel {
	rv := objc.Call[MappingModel](m_, objc.Sel("initWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Returns a mapping model initialized from a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmappingmodel/1506304-initwithcontentsofurl?language=objc
func MappingModel_InitWithContentsOfURL(url foundation.IURL) MappingModel {
	return MappingModelClass.Alloc().InitWithContentsOfURL(url)
}

func (mc _MappingModelClass) Alloc() MappingModel {
	rv := objc.Call[MappingModel](mc, objc.Sel("alloc"))
	return rv
}

func MappingModel_Alloc() MappingModel {
	return MappingModelClass.Alloc()
}

func (mc _MappingModelClass) New() MappingModel {
	rv := objc.Call[MappingModel](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMappingModel() MappingModel {
	return MappingModelClass.New()
}

func (m_ MappingModel) Init() MappingModel {
	rv := objc.Call[MappingModel](m_, objc.Sel("init"))
	return rv
}

// Returns the mapping model that will translate data from the source to the destination model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmappingmodel/1506930-mappingmodelfrombundles?language=objc
func (mc _MappingModelClass) MappingModelFromBundlesForSourceModelDestinationModel(bundles []foundation.IBundle, sourceModel IManagedObjectModel, destinationModel IManagedObjectModel) MappingModel {
	rv := objc.Call[MappingModel](mc, objc.Sel("mappingModelFromBundles:forSourceModel:destinationModel:"), bundles, objc.Ptr(sourceModel), objc.Ptr(destinationModel))
	return rv
}

// Returns the mapping model that will translate data from the source to the destination model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmappingmodel/1506930-mappingmodelfrombundles?language=objc
func MappingModel_MappingModelFromBundlesForSourceModelDestinationModel(bundles []foundation.IBundle, sourceModel IManagedObjectModel, destinationModel IManagedObjectModel) MappingModel {
	return MappingModelClass.MappingModelFromBundlesForSourceModelDestinationModel(bundles, sourceModel, destinationModel)
}

// Returns a newly created mapping model that will migrate data from the source to the destination model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmappingmodel/1506468-inferredmappingmodelforsourcemod?language=objc
func (mc _MappingModelClass) InferredMappingModelForSourceModelDestinationModelError(sourceModel IManagedObjectModel, destinationModel IManagedObjectModel, error foundation.IError) MappingModel {
	rv := objc.Call[MappingModel](mc, objc.Sel("inferredMappingModelForSourceModel:destinationModel:error:"), objc.Ptr(sourceModel), objc.Ptr(destinationModel), objc.Ptr(error))
	return rv
}

// Returns a newly created mapping model that will migrate data from the source to the destination model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmappingmodel/1506468-inferredmappingmodelforsourcemod?language=objc
func MappingModel_InferredMappingModelForSourceModelDestinationModelError(sourceModel IManagedObjectModel, destinationModel IManagedObjectModel, error foundation.IError) MappingModel {
	return MappingModelClass.InferredMappingModelForSourceModelDestinationModelError(sourceModel, destinationModel, error)
}

// The entity mappings for the mapping model, keyed by name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmappingmodel/1506179-entitymappingsbyname?language=objc
func (m_ MappingModel) EntityMappingsByName() map[string]EntityMapping {
	rv := objc.Call[map[string]EntityMapping](m_, objc.Sel("entityMappingsByName"))
	return rv
}

// The entity mappings for the mapping model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmappingmodel/1506196-entitymappings?language=objc
func (m_ MappingModel) EntityMappings() []EntityMapping {
	rv := objc.Call[[]EntityMapping](m_, objc.Sel("entityMappings"))
	return rv
}

// The entity mappings for the mapping model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmappingmodel/1506196-entitymappings?language=objc
func (m_ MappingModel) SetEntityMappings(value []IEntityMapping) {
	objc.Call[objc.Void](m_, objc.Sel("setEntityMappings:"), value)
}
