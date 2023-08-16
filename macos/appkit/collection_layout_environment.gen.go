// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol used to provide information about the layout’s container and environment traits, such as size classes and display scale factor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutenvironment?language=objc
type PCollectionLayoutEnvironment interface {
	// optional
	Container() PCollectionLayoutContainer
	HasContainer() bool
}

// A concrete type wrapper for the [PCollectionLayoutEnvironment] protocol.
type CollectionLayoutEnvironmentWrapper struct {
	objc.Object
}

func (c_ CollectionLayoutEnvironmentWrapper) HasContainer() bool {
	return c_.RespondsToSelector(objc.Sel("container"))
}

// Information about the layout's container, such as its size and content insets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutenvironment/3199073-container?language=objc
func (c_ CollectionLayoutEnvironmentWrapper) Container() CollectionLayoutContainerWrapper {
	rv := objc.Call[CollectionLayoutContainerWrapper](c_, objc.Sel("container"))
	return rv
}
