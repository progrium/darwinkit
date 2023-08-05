// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type ICollectionViewElement interface {
	IUserInterfaceItemIdentification
	ImplementsPrepareForReuse() bool
	// optional
	PrepareForReuse()
	ImplementsPreferredLayoutAttributesFittingAttributes() bool
	// optional
	PreferredLayoutAttributesFittingAttributes(layoutAttributes CollectionViewLayoutAttributes) ICollectionViewLayoutAttributes
	ImplementsApplyLayoutAttributes() bool
	// optional
	ApplyLayoutAttributes(layoutAttributes CollectionViewLayoutAttributes)
	ImplementsWillTransitionFromLayoutToLayout() bool
	// optional
	WillTransitionFromLayoutToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
	ImplementsDidTransitionFromLayoutToLayout() bool
	// optional
	DidTransitionFromLayoutToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
}

type CollectionViewElementWrapper struct {
	UserInterfaceItemIdentificationWrapper
}

func (c_ CollectionViewElementWrapper) ImplementsPrepareForReuse() bool {
	return c_.RespondsToSelector(objc.GetSelector("prepareForReuse"))
}

func (c_ CollectionViewElementWrapper) PrepareForReuse() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("prepareForReuse"))
}

func (c_ CollectionViewElementWrapper) ImplementsPreferredLayoutAttributesFittingAttributes() bool {
	return c_.RespondsToSelector(objc.GetSelector("preferredLayoutAttributesFittingAttributes:"))
}

func (c_ CollectionViewElementWrapper) PreferredLayoutAttributesFittingAttributes(layoutAttributes ICollectionViewLayoutAttributes) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("preferredLayoutAttributesFittingAttributes:"), objc.ExtractPtr(layoutAttributes))
	return rv
}

func (c_ CollectionViewElementWrapper) ImplementsApplyLayoutAttributes() bool {
	return c_.RespondsToSelector(objc.GetSelector("applyLayoutAttributes:"))
}

func (c_ CollectionViewElementWrapper) ApplyLayoutAttributes(layoutAttributes ICollectionViewLayoutAttributes) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("applyLayoutAttributes:"), objc.ExtractPtr(layoutAttributes))
}

func (c_ CollectionViewElementWrapper) ImplementsWillTransitionFromLayoutToLayout() bool {
	return c_.RespondsToSelector(objc.GetSelector("willTransitionFromLayout:toLayout:"))
}

func (c_ CollectionViewElementWrapper) WillTransitionFromLayoutToLayout(oldLayout ICollectionViewLayout, newLayout ICollectionViewLayout) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("willTransitionFromLayout:toLayout:"), objc.ExtractPtr(oldLayout), objc.ExtractPtr(newLayout))
}

func (c_ CollectionViewElementWrapper) ImplementsDidTransitionFromLayoutToLayout() bool {
	return c_.RespondsToSelector(objc.GetSelector("didTransitionFromLayout:toLayout:"))
}

func (c_ CollectionViewElementWrapper) DidTransitionFromLayoutToLayout(oldLayout ICollectionViewLayout, newLayout ICollectionViewLayout) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("didTransitionFromLayout:toLayout:"), objc.ExtractPtr(oldLayout), objc.ExtractPtr(newLayout))
}
