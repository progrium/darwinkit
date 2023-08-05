// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ClassDescriptionClass = _ClassDescriptionClass{objc.GetClass("NSClassDescription")}

type _ClassDescriptionClass struct {
	objc.Class
}

type IClassDescription interface {
	objc.IObject
	InverseForRelationshipKey(relationshipKey string) string
	AttributeKeys() []string
	ToManyRelationshipKeys() []string
	ToOneRelationshipKeys() []string
}

type ClassDescription struct {
	objc.Object
}

func MakeClassDescription(ptr unsafe.Pointer) ClassDescription {
	return ClassDescription{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _ClassDescriptionClass) Alloc() ClassDescription {
	rv := objc.CallMethod[ClassDescription](cc, objc.GetSelector("alloc"))
	return rv
}

func ClassDescription_Alloc() ClassDescription {
	return ClassDescriptionClass.Alloc()
}

func (cc _ClassDescriptionClass) New() ClassDescription {
	rv := objc.CallMethod[ClassDescription](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewClassDescription() ClassDescription {
	return ClassDescriptionClass.New()
}

func ClassDescription_New() ClassDescription {
	return ClassDescriptionClass.New()
}

func (c_ ClassDescription) Init() ClassDescription {
	rv := objc.CallMethod[ClassDescription](c_, objc.GetSelector("init"))
	return rv
}

func ClassDescription_Init() ClassDescription {
	return ClassDescriptionClass.Alloc().Init()
}

func (cc _ClassDescriptionClass) ClassDescriptionForClass(aClass objc.IClass) ClassDescription {
	rv := objc.CallMethod[ClassDescription](cc, objc.GetSelector("classDescriptionForClass:"), objc.ExtractPtr(aClass))
	return rv
}

func ClassDescription_ClassDescriptionForClass(aClass objc.IClass) ClassDescription {
	return ClassDescriptionClass.ClassDescriptionForClass(aClass)
}

func (cc _ClassDescriptionClass) InvalidateClassDescriptionCache() {
	objc.CallMethod[objc.Void](cc, objc.GetSelector("invalidateClassDescriptionCache"))
}

func ClassDescription_InvalidateClassDescriptionCache() {
	ClassDescriptionClass.InvalidateClassDescriptionCache()
}

func (cc _ClassDescriptionClass) RegisterClassDescriptionForClass(description IClassDescription, aClass objc.IClass) {
	objc.CallMethod[objc.Void](cc, objc.GetSelector("registerClassDescription:forClass:"), objc.ExtractPtr(description), objc.ExtractPtr(aClass))
}

func ClassDescription_RegisterClassDescriptionForClass(description IClassDescription, aClass objc.IClass) {
	ClassDescriptionClass.RegisterClassDescriptionForClass(description, aClass)
}

func (c_ ClassDescription) InverseForRelationshipKey(relationshipKey string) string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("inverseForRelationshipKey:"), relationshipKey)
	return rv
}

func (c_ ClassDescription) AttributeKeys() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("attributeKeys"))
	return rv
}

func (c_ ClassDescription) ToManyRelationshipKeys() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("toManyRelationshipKeys"))
	return rv
}

func (c_ ClassDescription) ToOneRelationshipKeys() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("toOneRelationshipKeys"))
	return rv
}
