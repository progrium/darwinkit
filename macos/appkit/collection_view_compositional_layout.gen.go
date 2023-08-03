// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CollectionViewCompositionalLayoutClass = _CollectionViewCompositionalLayoutClass{objc.GetClass("NSCollectionViewCompositionalLayout")}

type _CollectionViewCompositionalLayoutClass struct {
	objc.Class
}

type ICollectionViewCompositionalLayout interface {
	ICollectionViewLayout
	Configuration() CollectionViewCompositionalLayoutConfiguration
	SetConfiguration(value ICollectionViewCompositionalLayoutConfiguration)
}

type CollectionViewCompositionalLayout struct {
	CollectionViewLayout
}

func MakeCollectionViewCompositionalLayout(ptr unsafe.Pointer) CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayout{
		CollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func (c_ CollectionViewCompositionalLayout) InitWithSection(section ICollectionLayoutSection) CollectionViewCompositionalLayout {
	rv := objc.CallMethod[CollectionViewCompositionalLayout](c_, objc.GetSelector("initWithSection:"), objc.ExtractPtr(section))
	return rv
}

func CollectionViewCompositionalLayout_InitWithSection(section ICollectionLayoutSection) CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayoutClass.Alloc().InitWithSection(section)
}

func (c_ CollectionViewCompositionalLayout) InitWithSectionConfiguration(section ICollectionLayoutSection, configuration ICollectionViewCompositionalLayoutConfiguration) CollectionViewCompositionalLayout {
	rv := objc.CallMethod[CollectionViewCompositionalLayout](c_, objc.GetSelector("initWithSection:configuration:"), objc.ExtractPtr(section), objc.ExtractPtr(configuration))
	return rv
}

func CollectionViewCompositionalLayout_InitWithSectionConfiguration(section ICollectionLayoutSection, configuration ICollectionViewCompositionalLayoutConfiguration) CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayoutClass.Alloc().InitWithSectionConfiguration(section, configuration)
}

func (c_ CollectionViewCompositionalLayout) InitWithSectionProvider(sectionProvider func(section int, param2 CollectionLayoutEnvironmentWrapper) CollectionLayoutSection) CollectionViewCompositionalLayout {
	rv := objc.CallMethod[CollectionViewCompositionalLayout](c_, objc.GetSelector("initWithSectionProvider:"), sectionProvider)
	return rv
}

func CollectionViewCompositionalLayout_InitWithSectionProvider(sectionProvider func(section int, param2 CollectionLayoutEnvironmentWrapper) CollectionLayoutSection) CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayoutClass.Alloc().InitWithSectionProvider(sectionProvider)
}

func (c_ CollectionViewCompositionalLayout) InitWithSectionProviderConfiguration(sectionProvider func(section int, param2 CollectionLayoutEnvironmentWrapper) CollectionLayoutSection, configuration ICollectionViewCompositionalLayoutConfiguration) CollectionViewCompositionalLayout {
	rv := objc.CallMethod[CollectionViewCompositionalLayout](c_, objc.GetSelector("initWithSectionProvider:configuration:"), sectionProvider, objc.ExtractPtr(configuration))
	return rv
}

func CollectionViewCompositionalLayout_InitWithSectionProviderConfiguration(sectionProvider func(section int, param2 CollectionLayoutEnvironmentWrapper) CollectionLayoutSection, configuration ICollectionViewCompositionalLayoutConfiguration) CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayoutClass.Alloc().InitWithSectionProviderConfiguration(sectionProvider, configuration)
}

func (cc _CollectionViewCompositionalLayoutClass) Alloc() CollectionViewCompositionalLayout {
	rv := objc.CallMethod[CollectionViewCompositionalLayout](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionViewCompositionalLayout_Alloc() CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayoutClass.Alloc()
}

func (cc _CollectionViewCompositionalLayoutClass) New() CollectionViewCompositionalLayout {
	rv := objc.CallMethod[CollectionViewCompositionalLayout](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewCompositionalLayout() CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayoutClass.New()
}

func CollectionViewCompositionalLayout_New() CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayoutClass.New()
}

func (c_ CollectionViewCompositionalLayout) Init() CollectionViewCompositionalLayout {
	rv := objc.CallMethod[CollectionViewCompositionalLayout](c_, objc.GetSelector("init"))
	return rv
}

func CollectionViewCompositionalLayout_Init() CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayoutClass.Alloc().Init()
}

func (c_ CollectionViewCompositionalLayout) Configuration() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.CallMethod[CollectionViewCompositionalLayoutConfiguration](c_, objc.GetSelector("configuration"))
	return rv
}

func (c_ CollectionViewCompositionalLayout) SetConfiguration(value ICollectionViewCompositionalLayoutConfiguration) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setConfiguration:"), objc.ExtractPtr(value))
}
