// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that a Cloud-sharing toolbar item uses to get validation of an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingvalidation?language=objc
type PCloudSharingValidation interface {
	// optional
	CloudShareForUserInterfaceItem(item ValidatedUserInterfaceItemWrapper) objc.IObject
	HasCloudShareForUserInterfaceItem() bool
}

// A concrete type wrapper for the [PCloudSharingValidation] protocol.
type CloudSharingValidationWrapper struct {
	objc.Object
}

func (c_ CloudSharingValidationWrapper) HasCloudShareForUserInterfaceItem() bool {
	return c_.RespondsToSelector(objc.Sel("cloudShareForUserInterfaceItem:"))
}

// Returns the Cloud share object that corresponds to the specified item, if one exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingvalidation/2315049-cloudshareforuserinterfaceitem?language=objc
func (c_ CloudSharingValidationWrapper) CloudShareForUserInterfaceItem(item PValidatedUserInterfaceItem) objc.Object {
	po0 := objc.WrapAsProtocol("NSValidatedUserInterfaceItem", item)
	rv := objc.Call[objc.Object](c_, objc.Sel("cloudShareForUserInterfaceItem:"), po0)
	return rv
}
