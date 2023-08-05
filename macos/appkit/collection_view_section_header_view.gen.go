// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type ICollectionViewSectionHeaderView interface {
	ICollectionViewElement
	ImplementsSetSectionCollapseButton() bool
	// optional
	SetSectionCollapseButton(value Button)
	ImplementsSectionCollapseButton() bool
	// optional
	SectionCollapseButton() IButton
}

type CollectionViewSectionHeaderViewWrapper struct {
	CollectionViewElementWrapper
}

func (c_ CollectionViewSectionHeaderViewWrapper) ImplementsSetSectionCollapseButton() bool {
	return c_.RespondsToSelector(objc.GetSelector("setSectionCollapseButton:"))
}

func (c_ CollectionViewSectionHeaderViewWrapper) SetSectionCollapseButton(value IButton) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setSectionCollapseButton:"), objc.ExtractPtr(value))
}

func (c_ CollectionViewSectionHeaderViewWrapper) ImplementsSectionCollapseButton() bool {
	return c_.RespondsToSelector(objc.GetSelector("sectionCollapseButton"))
}

func (c_ CollectionViewSectionHeaderViewWrapper) SectionCollapseButton() Button {
	rv := objc.CallMethod[Button](c_, objc.GetSelector("sectionCollapseButton"))
	return rv
}
