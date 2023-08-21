// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionViewCompositionalLayout] class.
var CollectionViewCompositionalLayoutClass = _CollectionViewCompositionalLayoutClass{objc.GetClass("NSCollectionViewCompositionalLayout")}

type _CollectionViewCompositionalLayoutClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewCompositionalLayout] class.
type ICollectionViewCompositionalLayout interface {
	ICollectionViewLayout
	Configuration() CollectionViewCompositionalLayoutConfiguration
	SetConfiguration(value ICollectionViewCompositionalLayoutConfiguration)
}

// A layout object that lets you combine items in highly adaptive and flexible visual arrangements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayout?language=objc
type CollectionViewCompositionalLayout struct {
	CollectionViewLayout
}

func CollectionViewCompositionalLayoutFrom(ptr unsafe.Pointer) CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayout{
		CollectionViewLayout: CollectionViewLayoutFrom(ptr),
	}
}

func (c_ CollectionViewCompositionalLayout) InitWithSection(section ICollectionLayoutSection) CollectionViewCompositionalLayout {
	rv := objc.Call[CollectionViewCompositionalLayout](c_, objc.Sel("initWithSection:"), objc.Ptr(section))
	return rv
}

// Creates a compositional layout object with a single section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayout/3281806-initwithsection?language=objc
func NewCollectionViewCompositionalLayoutWithSection(section ICollectionLayoutSection) CollectionViewCompositionalLayout {
	instance := CollectionViewCompositionalLayoutClass.Alloc().InitWithSection(section)
	instance.Autorelease()
	return instance
}

func (c_ CollectionViewCompositionalLayout) InitWithSectionProvider(sectionProvider CollectionViewCompositionalLayoutSectionProvider) CollectionViewCompositionalLayout {
	rv := objc.Call[CollectionViewCompositionalLayout](c_, objc.Sel("initWithSectionProvider:"), sectionProvider)
	return rv
}

// Creates a compositional layout object with a section provider to supply the layout's sections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayout/3281808-initwithsectionprovider?language=objc
func NewCollectionViewCompositionalLayoutWithSectionProvider(sectionProvider CollectionViewCompositionalLayoutSectionProvider) CollectionViewCompositionalLayout {
	instance := CollectionViewCompositionalLayoutClass.Alloc().InitWithSectionProvider(sectionProvider)
	instance.Autorelease()
	return instance
}

func (cc _CollectionViewCompositionalLayoutClass) Alloc() CollectionViewCompositionalLayout {
	rv := objc.Call[CollectionViewCompositionalLayout](cc, objc.Sel("alloc"))
	return rv
}

func CollectionViewCompositionalLayout_Alloc() CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayoutClass.Alloc()
}

func (cc _CollectionViewCompositionalLayoutClass) New() CollectionViewCompositionalLayout {
	rv := objc.Call[CollectionViewCompositionalLayout](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewCompositionalLayout() CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayoutClass.New()
}

func (c_ CollectionViewCompositionalLayout) Init() CollectionViewCompositionalLayout {
	rv := objc.Call[CollectionViewCompositionalLayout](c_, objc.Sel("init"))
	return rv
}

// The layout's configuration, such as its scroll direction and section spacing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayout/3281805-configuration?language=objc
func (c_ CollectionViewCompositionalLayout) Configuration() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.Call[CollectionViewCompositionalLayoutConfiguration](c_, objc.Sel("configuration"))
	return rv
}

// The layout's configuration, such as its scroll direction and section spacing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayout/3281805-configuration?language=objc
func (c_ CollectionViewCompositionalLayout) SetConfiguration(value ICollectionViewCompositionalLayoutConfiguration) {
	objc.Call[objc.Void](c_, objc.Sel("setConfiguration:"), objc.Ptr(value))
}
