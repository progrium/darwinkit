// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that you use to manage the content in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewelement?language=objc
type PCollectionViewElement interface {
	// optional
	WillTransitionFromLayoutToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
	HasWillTransitionFromLayoutToLayout() bool

	// optional
	ApplyLayoutAttributes(layoutAttributes CollectionViewLayoutAttributes)
	HasApplyLayoutAttributes() bool

	// optional
	DidTransitionFromLayoutToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
	HasDidTransitionFromLayoutToLayout() bool

	// optional
	PrepareForReuse()
	HasPrepareForReuse() bool

	// optional
	PreferredLayoutAttributesFittingAttributes(layoutAttributes CollectionViewLayoutAttributes) ICollectionViewLayoutAttributes
	HasPreferredLayoutAttributesFittingAttributes() bool
}

// A concrete type wrapper for the [PCollectionViewElement] protocol.
type CollectionViewElementWrapper struct {
	objc.Object
}

func (c_ CollectionViewElementWrapper) HasWillTransitionFromLayoutToLayout() bool {
	return c_.RespondsToSelector(objc.Sel("willTransitionFromLayout:toLayout:"))
}

// Tells the element that the layout object of the collection view is about to change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewelement/1528165-willtransitionfromlayout?language=objc
func (c_ CollectionViewElementWrapper) WillTransitionFromLayoutToLayout(oldLayout ICollectionViewLayout, newLayout ICollectionViewLayout) {
	objc.Call[objc.Void](c_, objc.Sel("willTransitionFromLayout:toLayout:"), objc.Ptr(oldLayout), objc.Ptr(newLayout))
}

func (c_ CollectionViewElementWrapper) HasApplyLayoutAttributes() bool {
	return c_.RespondsToSelector(objc.Sel("applyLayoutAttributes:"))
}

// Applies the specified layout attributes to the element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewelement/1528294-applylayoutattributes?language=objc
func (c_ CollectionViewElementWrapper) ApplyLayoutAttributes(layoutAttributes ICollectionViewLayoutAttributes) {
	objc.Call[objc.Void](c_, objc.Sel("applyLayoutAttributes:"), objc.Ptr(layoutAttributes))
}

func (c_ CollectionViewElementWrapper) HasDidTransitionFromLayoutToLayout() bool {
	return c_.RespondsToSelector(objc.Sel("didTransitionFromLayout:toLayout:"))
}

// Tells the element that the layout object of the collection view changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewelement/1525851-didtransitionfromlayout?language=objc
func (c_ CollectionViewElementWrapper) DidTransitionFromLayoutToLayout(oldLayout ICollectionViewLayout, newLayout ICollectionViewLayout) {
	objc.Call[objc.Void](c_, objc.Sel("didTransitionFromLayout:toLayout:"), objc.Ptr(oldLayout), objc.Ptr(newLayout))
}

func (c_ CollectionViewElementWrapper) HasPrepareForReuse() bool {
	return c_.RespondsToSelector(objc.Sel("prepareForReuse"))
}

// Performs any necessary cleanup to prepare the element for use again. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewelement/1528248-prepareforreuse?language=objc
func (c_ CollectionViewElementWrapper) PrepareForReuse() {
	objc.Call[objc.Void](c_, objc.Sel("prepareForReuse"))
}

func (c_ CollectionViewElementWrapper) HasPreferredLayoutAttributesFittingAttributes() bool {
	return c_.RespondsToSelector(objc.Sel("preferredLayoutAttributesFittingAttributes:"))
}

// Asks your element if it wants to modify any layout attributes before they are applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewelement/1528259-preferredlayoutattributesfitting?language=objc
func (c_ CollectionViewElementWrapper) PreferredLayoutAttributesFittingAttributes(layoutAttributes ICollectionViewLayoutAttributes) CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](c_, objc.Sel("preferredLayoutAttributesFittingAttributes:"), objc.Ptr(layoutAttributes))
	return rv
}
