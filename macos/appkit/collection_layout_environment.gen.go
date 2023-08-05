// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type ICollectionLayoutEnvironment interface {
	ImplementsContainer() bool
	// optional
	Container() ICollectionLayoutContainer
}

type CollectionLayoutEnvironmentWrapper struct {
	objc.Object
}

func (c_ CollectionLayoutEnvironmentWrapper) ImplementsContainer() bool {
	return c_.RespondsToSelector(objc.GetSelector("container"))
}

func (c_ CollectionLayoutEnvironmentWrapper) Container() CollectionLayoutContainerWrapper {
	rv := objc.CallMethod[CollectionLayoutContainerWrapper](c_, objc.GetSelector("container"))
	return rv
}
