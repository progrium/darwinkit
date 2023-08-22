// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CompositionRepository] class.
var CompositionRepositoryClass = _CompositionRepositoryClass{objc.GetClass("QCCompositionRepository")}

type _CompositionRepositoryClass struct {
	objc.Class
}

// An interface definition for the [CompositionRepository] class.
type ICompositionRepository interface {
	objc.IObject
}

// The QCCompositionRepository class represents a system-wide centralized repository of built-in and installed Quartz Composer compositions (/Library/Compositions and ~/Library/Compositions). The QCCompositionRepository class cannot be subclassed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qccompositionrepository?language=objc
type CompositionRepository struct {
	objc.Object
}

func CompositionRepositoryFrom(ptr unsafe.Pointer) CompositionRepository {
	return CompositionRepository{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CompositionRepositoryClass) Alloc() CompositionRepository {
	rv := objc.Call[CompositionRepository](cc, objc.Sel("alloc"))
	return rv
}

func CompositionRepository_Alloc() CompositionRepository {
	return CompositionRepositoryClass.Alloc()
}

func (cc _CompositionRepositoryClass) New() CompositionRepository {
	rv := objc.Call[CompositionRepository](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCompositionRepository() CompositionRepository {
	return CompositionRepositoryClass.New()
}

func (c_ CompositionRepository) Init() CompositionRepository {
	rv := objc.Call[CompositionRepository](c_, objc.Sel("init"))
	return rv
}
