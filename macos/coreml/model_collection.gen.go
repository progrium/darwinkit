// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ModelCollection] class.
var ModelCollectionClass = _ModelCollectionClass{objc.GetClass("MLModelCollection")}

type _ModelCollectionClass struct {
	objc.Class
}

// An interface definition for the [ModelCollection] class.
type IModelCollection interface {
	objc.IObject
}

// A set of Core ML models from a model deployment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelcollection?language=objc
type ModelCollection struct {
	objc.Object
}

func ModelCollectionFrom(ptr unsafe.Pointer) ModelCollection {
	return ModelCollection{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _ModelCollectionClass) Alloc() ModelCollection {
	rv := objc.Call[ModelCollection](mc, objc.Sel("alloc"))
	return rv
}

func ModelCollection_Alloc() ModelCollection {
	return ModelCollectionClass.Alloc()
}

func (mc _ModelCollectionClass) New() ModelCollection {
	rv := objc.Call[ModelCollection](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewModelCollection() ModelCollection {
	return ModelCollectionClass.New()
}

func (m_ ModelCollection) Init() ModelCollection {
	rv := objc.Call[ModelCollection](m_, objc.Sel("init"))
	return rv
}
