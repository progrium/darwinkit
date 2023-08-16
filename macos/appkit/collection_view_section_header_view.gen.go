// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines a button to control the collapse of a collection view’s section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewsectionheaderview?language=objc
type PCollectionViewSectionHeaderView interface {
	// optional
	SetSectionCollapseButton(value Button)
	HasSetSectionCollapseButton() bool

	// optional
	SectionCollapseButton() IButton
	HasSectionCollapseButton() bool
}

// A concrete type wrapper for the [PCollectionViewSectionHeaderView] protocol.
type CollectionViewSectionHeaderViewWrapper struct {
	objc.Object
}

func (c_ CollectionViewSectionHeaderViewWrapper) HasSetSectionCollapseButton() bool {
	return c_.RespondsToSelector(objc.Sel("setSectionCollapseButton:"))
}

// A control that lets users collapse and open a collection view section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewsectionheaderview/1644266-sectioncollapsebutton?language=objc
func (c_ CollectionViewSectionHeaderViewWrapper) SetSectionCollapseButton(value IButton) {
	objc.Call[objc.Void](c_, objc.Sel("setSectionCollapseButton:"), objc.Ptr(value))
}

func (c_ CollectionViewSectionHeaderViewWrapper) HasSectionCollapseButton() bool {
	return c_.RespondsToSelector(objc.Sel("sectionCollapseButton"))
}

// A control that lets users collapse and open a collection view section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewsectionheaderview/1644266-sectioncollapsebutton?language=objc
func (c_ CollectionViewSectionHeaderViewWrapper) SectionCollapseButton() Button {
	rv := objc.Call[Button](c_, objc.Sel("sectionCollapseButton"))
	return rv
}
