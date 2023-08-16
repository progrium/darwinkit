// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// A protocol used to provide information about the size and content insets of a layout’s container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutcontainer?language=objc
type PCollectionLayoutContainer interface {
	// optional
	ContentInsets() DirectionalEdgeInsets
	HasContentInsets() bool

	// optional
	ContentSize() coregraphics.Size
	HasContentSize() bool

	// optional
	EffectiveContentInsets() DirectionalEdgeInsets
	HasEffectiveContentInsets() bool

	// optional
	EffectiveContentSize() coregraphics.Size
	HasEffectiveContentSize() bool
}

// A concrete type wrapper for the [PCollectionLayoutContainer] protocol.
type CollectionLayoutContainerWrapper struct {
	objc.Object
}

func (c_ CollectionLayoutContainerWrapper) HasContentInsets() bool {
	return c_.RespondsToSelector(objc.Sel("contentInsets"))
}

// The amount of space added around the content of the container to adjust its final size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutcontainer/3199046-contentinsets?language=objc
func (c_ CollectionLayoutContainerWrapper) ContentInsets() DirectionalEdgeInsets {
	rv := objc.Call[DirectionalEdgeInsets](c_, objc.Sel("contentInsets"))
	return rv
}

func (c_ CollectionLayoutContainerWrapper) HasContentSize() bool {
	return c_.RespondsToSelector(objc.Sel("contentSize"))
}

// The size of the container before content insets are applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutcontainer/3199047-contentsize?language=objc
func (c_ CollectionLayoutContainerWrapper) ContentSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](c_, objc.Sel("contentSize"))
	return rv
}

func (c_ CollectionLayoutContainerWrapper) HasEffectiveContentInsets() bool {
	return c_.RespondsToSelector(objc.Sel("effectiveContentInsets"))
}

// The amount of space added around the content of the container to adjust its final size after item content insets are applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutcontainer/3199048-effectivecontentinsets?language=objc
func (c_ CollectionLayoutContainerWrapper) EffectiveContentInsets() DirectionalEdgeInsets {
	rv := objc.Call[DirectionalEdgeInsets](c_, objc.Sel("effectiveContentInsets"))
	return rv
}

func (c_ CollectionLayoutContainerWrapper) HasEffectiveContentSize() bool {
	return c_.RespondsToSelector(objc.Sel("effectiveContentSize"))
}

// The size of the container after content insets are applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutcontainer/3199049-effectivecontentsize?language=objc
func (c_ CollectionLayoutContainerWrapper) EffectiveContentSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](c_, objc.Sel("effectiveContentSize"))
	return rv
}
