// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ModelCollectionEntry] class.
var ModelCollectionEntryClass = _ModelCollectionEntryClass{objc.GetClass("MLModelCollectionEntry")}

type _ModelCollectionEntryClass struct {
	objc.Class
}

// An interface definition for the [ModelCollectionEntry] class.
type IModelCollectionEntry interface {
	objc.IObject
}

// A model and its identifier within a model collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelcollectionentry?language=objc
type ModelCollectionEntry struct {
	objc.Object
}

func ModelCollectionEntryFrom(ptr unsafe.Pointer) ModelCollectionEntry {
	return ModelCollectionEntry{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _ModelCollectionEntryClass) Alloc() ModelCollectionEntry {
	rv := objc.Call[ModelCollectionEntry](mc, objc.Sel("alloc"))
	return rv
}

func ModelCollectionEntry_Alloc() ModelCollectionEntry {
	return ModelCollectionEntryClass.Alloc()
}

func (mc _ModelCollectionEntryClass) New() ModelCollectionEntry {
	rv := objc.Call[ModelCollectionEntry](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewModelCollectionEntry() ModelCollectionEntry {
	return ModelCollectionEntryClass.New()
}

func (m_ ModelCollectionEntry) Init() ModelCollectionEntry {
	rv := objc.Call[ModelCollectionEntry](m_, objc.Sel("init"))
	return rv
}
