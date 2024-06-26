// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A protocol used to provide information about the size and content insets of a layout’s container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutcontainer?language=objc
type PCollectionLayoutContainer interface {
	// optional
	EffectiveContentInsets() DirectionalEdgeInsets
	HasEffectiveContentInsets() bool

	// optional
	ContentInsets() DirectionalEdgeInsets
	HasContentInsets() bool

	// optional
	ContentSize() foundation.Size
	HasContentSize() bool

	// optional
	EffectiveContentSize() foundation.Size
	HasEffectiveContentSize() bool
}

// ensure impl type implements protocol interface
var _ PCollectionLayoutContainer = (*CollectionLayoutContainerObject)(nil)

// A concrete type for the [PCollectionLayoutContainer] protocol.
type CollectionLayoutContainerObject struct {
	objc.Object
}

func (c_ CollectionLayoutContainerObject) HasEffectiveContentInsets() bool {
	return c_.RespondsToSelector(objc.Sel("effectiveContentInsets"))
}

// The amount of space added around the content of the container to adjust its final size after item content insets are applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutcontainer/3199048-effectivecontentinsets?language=objc
func (c_ CollectionLayoutContainerObject) EffectiveContentInsets() DirectionalEdgeInsets {
	rv := objc.Call[DirectionalEdgeInsets](c_, objc.Sel("effectiveContentInsets"))
	return rv
}

func (c_ CollectionLayoutContainerObject) HasContentInsets() bool {
	return c_.RespondsToSelector(objc.Sel("contentInsets"))
}

// The amount of space added around the content of the container to adjust its final size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutcontainer/3199046-contentinsets?language=objc
func (c_ CollectionLayoutContainerObject) ContentInsets() DirectionalEdgeInsets {
	rv := objc.Call[DirectionalEdgeInsets](c_, objc.Sel("contentInsets"))
	return rv
}

func (c_ CollectionLayoutContainerObject) HasContentSize() bool {
	return c_.RespondsToSelector(objc.Sel("contentSize"))
}

// The size of the container before content insets are applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutcontainer/3199047-contentsize?language=objc
func (c_ CollectionLayoutContainerObject) ContentSize() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("contentSize"))
	return rv
}

func (c_ CollectionLayoutContainerObject) HasEffectiveContentSize() bool {
	return c_.RespondsToSelector(objc.Sel("effectiveContentSize"))
}

// The size of the container after content insets are applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutcontainer/3199049-effectivecontentsize?language=objc
func (c_ CollectionLayoutContainerObject) EffectiveContentSize() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("effectiveContentSize"))
	return rv
}
