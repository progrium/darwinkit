// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IPathCellDelegate interface {
	ImplementsPathCellWillDisplayOpenPanel() bool
	// optional
	PathCellWillDisplayOpenPanel(pathCell PathCell, openPanel OpenPanel)
	ImplementsPathCellWillPopUpMenu() bool
	// optional
	PathCellWillPopUpMenu(pathCell PathCell, menu Menu)
}

type PathCellDelegate struct {
	_PathCellWillDisplayOpenPanel func(pathCell PathCell, openPanel OpenPanel)
	_PathCellWillPopUpMenu        func(pathCell PathCell, menu Menu)
}

func (di *PathCellDelegate) ImplementsPathCellWillDisplayOpenPanel() bool {
	return di._PathCellWillDisplayOpenPanel != nil
}

func (di *PathCellDelegate) SetPathCellWillDisplayOpenPanel(f func(pathCell PathCell, openPanel OpenPanel)) {
	di._PathCellWillDisplayOpenPanel = f
}

func (di *PathCellDelegate) PathCellWillDisplayOpenPanel(pathCell PathCell, openPanel OpenPanel) {
	di._PathCellWillDisplayOpenPanel(pathCell, openPanel)
}
func (di *PathCellDelegate) ImplementsPathCellWillPopUpMenu() bool {
	return di._PathCellWillPopUpMenu != nil
}

func (di *PathCellDelegate) SetPathCellWillPopUpMenu(f func(pathCell PathCell, menu Menu)) {
	di._PathCellWillPopUpMenu = f
}

func (di *PathCellDelegate) PathCellWillPopUpMenu(pathCell PathCell, menu Menu) {
	di._PathCellWillPopUpMenu(pathCell, menu)
}

type PathCellDelegateWrapper struct {
	objc.Object
}

func (p_ PathCellDelegateWrapper) ImplementsPathCellWillDisplayOpenPanel() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathCell:willDisplayOpenPanel:"))
}

func (p_ PathCellDelegateWrapper) PathCellWillDisplayOpenPanel(pathCell IPathCell, openPanel IOpenPanel) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("pathCell:willDisplayOpenPanel:"), objc.ExtractPtr(pathCell), objc.ExtractPtr(openPanel))
}

func (p_ PathCellDelegateWrapper) ImplementsPathCellWillPopUpMenu() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathCell:willPopUpMenu:"))
}

func (p_ PathCellDelegateWrapper) PathCellWillPopUpMenu(pathCell IPathCell, menu IMenu) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("pathCell:willPopUpMenu:"), objc.ExtractPtr(pathCell), objc.ExtractPtr(menu))
}
