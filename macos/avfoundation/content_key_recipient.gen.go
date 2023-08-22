// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol for requiring decryption keys for media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrecipient?language=objc
type PContentKeyRecipient interface {
	// optional
	ContentKeySessionDidProvideContentKey(contentKeySession ContentKeySession, contentKey ContentKey)
	HasContentKeySessionDidProvideContentKey() bool

	// optional
	MayRequireContentKeysForMediaDataProcessing() bool
	HasMayRequireContentKeysForMediaDataProcessing() bool
}

// A concrete type wrapper for the [PContentKeyRecipient] protocol.
type ContentKeyRecipientWrapper struct {
	objc.Object
}

func (c_ ContentKeyRecipientWrapper) HasContentKeySessionDidProvideContentKey() bool {
	return c_.RespondsToSelector(objc.Sel("contentKeySession:didProvideContentKey:"))
}

// Tells the recipient that a content key is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrecipient/3726101-contentkeysession?language=objc
func (c_ ContentKeyRecipientWrapper) ContentKeySessionDidProvideContentKey(contentKeySession IContentKeySession, contentKey IContentKey) {
	objc.Call[objc.Void](c_, objc.Sel("contentKeySession:didProvideContentKey:"), objc.Ptr(contentKeySession), objc.Ptr(contentKey))
}

func (c_ ContentKeyRecipientWrapper) HasMayRequireContentKeysForMediaDataProcessing() bool {
	return c_.RespondsToSelector(objc.Sel("mayRequireContentKeysForMediaDataProcessing"))
}

// A Boolean value that indicates whether the recipient requires decryption keys for media data to enable processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrecipient/2799194-mayrequirecontentkeysformediadat?language=objc
func (c_ ContentKeyRecipientWrapper) MayRequireContentKeysForMediaDataProcessing() bool {
	rv := objc.Call[bool](c_, objc.Sel("mayRequireContentKeysForMediaDataProcessing"))
	return rv
}
