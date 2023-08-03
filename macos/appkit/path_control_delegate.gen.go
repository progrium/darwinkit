// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IPathControlDelegate interface {
	ImplementsPathControlShouldDragPathComponentCellWithPasteboard() bool
	// optional
	PathControlShouldDragPathComponentCellWithPasteboard(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool
	ImplementsPathControlValidateDrop() bool
	// optional
	PathControlValidateDrop(pathControl PathControl, info DraggingInfoWrapper) DragOperation
	ImplementsPathControlAcceptDrop() bool
	// optional
	PathControlAcceptDrop(pathControl PathControl, info DraggingInfoWrapper) bool
	ImplementsPathControlWillDisplayOpenPanel() bool
	// optional
	PathControlWillDisplayOpenPanel(pathControl PathControl, openPanel OpenPanel)
	ImplementsPathControlWillPopUpMenu() bool
	// optional
	PathControlWillPopUpMenu(pathControl PathControl, menu Menu)
	ImplementsPathControlShouldDragItemWithPasteboard() bool
	// optional
	PathControlShouldDragItemWithPasteboard(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool
}

type PathControlDelegate struct {
	_PathControlShouldDragPathComponentCellWithPasteboard func(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool
	_PathControlValidateDrop                              func(pathControl PathControl, info DraggingInfoWrapper) DragOperation
	_PathControlAcceptDrop                                func(pathControl PathControl, info DraggingInfoWrapper) bool
	_PathControlWillDisplayOpenPanel                      func(pathControl PathControl, openPanel OpenPanel)
	_PathControlWillPopUpMenu                             func(pathControl PathControl, menu Menu)
	_PathControlShouldDragItemWithPasteboard              func(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool
}

func (di *PathControlDelegate) ImplementsPathControlShouldDragPathComponentCellWithPasteboard() bool {
	return di._PathControlShouldDragPathComponentCellWithPasteboard != nil
}

func (di *PathControlDelegate) SetPathControlShouldDragPathComponentCellWithPasteboard(f func(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool) {
	di._PathControlShouldDragPathComponentCellWithPasteboard = f
}

func (di *PathControlDelegate) PathControlShouldDragPathComponentCellWithPasteboard(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool {
	return di._PathControlShouldDragPathComponentCellWithPasteboard(pathControl, pathComponentCell, pasteboard)
}
func (di *PathControlDelegate) ImplementsPathControlValidateDrop() bool {
	return di._PathControlValidateDrop != nil
}

func (di *PathControlDelegate) SetPathControlValidateDrop(f func(pathControl PathControl, info DraggingInfoWrapper) DragOperation) {
	di._PathControlValidateDrop = f
}

func (di *PathControlDelegate) PathControlValidateDrop(pathControl PathControl, info DraggingInfoWrapper) DragOperation {
	return di._PathControlValidateDrop(pathControl, info)
}
func (di *PathControlDelegate) ImplementsPathControlAcceptDrop() bool {
	return di._PathControlAcceptDrop != nil
}

func (di *PathControlDelegate) SetPathControlAcceptDrop(f func(pathControl PathControl, info DraggingInfoWrapper) bool) {
	di._PathControlAcceptDrop = f
}

func (di *PathControlDelegate) PathControlAcceptDrop(pathControl PathControl, info DraggingInfoWrapper) bool {
	return di._PathControlAcceptDrop(pathControl, info)
}
func (di *PathControlDelegate) ImplementsPathControlWillDisplayOpenPanel() bool {
	return di._PathControlWillDisplayOpenPanel != nil
}

func (di *PathControlDelegate) SetPathControlWillDisplayOpenPanel(f func(pathControl PathControl, openPanel OpenPanel)) {
	di._PathControlWillDisplayOpenPanel = f
}

func (di *PathControlDelegate) PathControlWillDisplayOpenPanel(pathControl PathControl, openPanel OpenPanel) {
	di._PathControlWillDisplayOpenPanel(pathControl, openPanel)
}
func (di *PathControlDelegate) ImplementsPathControlWillPopUpMenu() bool {
	return di._PathControlWillPopUpMenu != nil
}

func (di *PathControlDelegate) SetPathControlWillPopUpMenu(f func(pathControl PathControl, menu Menu)) {
	di._PathControlWillPopUpMenu = f
}

func (di *PathControlDelegate) PathControlWillPopUpMenu(pathControl PathControl, menu Menu) {
	di._PathControlWillPopUpMenu(pathControl, menu)
}
func (di *PathControlDelegate) ImplementsPathControlShouldDragItemWithPasteboard() bool {
	return di._PathControlShouldDragItemWithPasteboard != nil
}

func (di *PathControlDelegate) SetPathControlShouldDragItemWithPasteboard(f func(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool) {
	di._PathControlShouldDragItemWithPasteboard = f
}

func (di *PathControlDelegate) PathControlShouldDragItemWithPasteboard(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool {
	return di._PathControlShouldDragItemWithPasteboard(pathControl, pathItem, pasteboard)
}

type PathControlDelegateWrapper struct {
	objc.Object
}

func (p_ PathControlDelegateWrapper) ImplementsPathControlShouldDragPathComponentCellWithPasteboard() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:shouldDragPathComponentCell:withPasteboard:"))
}

func (p_ PathControlDelegateWrapper) PathControlShouldDragPathComponentCellWithPasteboard(pathControl IPathControl, pathComponentCell IPathComponentCell, pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("pathControl:shouldDragPathComponentCell:withPasteboard:"), objc.ExtractPtr(pathControl), objc.ExtractPtr(pathComponentCell), objc.ExtractPtr(pasteboard))
	return rv
}

func (p_ PathControlDelegateWrapper) ImplementsPathControlValidateDrop() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:validateDrop:"))
}

func (p_ PathControlDelegateWrapper) PathControlValidateDrop(pathControl IPathControl, info IDraggingInfo) DragOperation {
	po := objc.WrapAsProtocol("NSDraggingInfo", info)
	rv := objc.CallMethod[DragOperation](p_, objc.GetSelector("pathControl:validateDrop:"), objc.ExtractPtr(pathControl), po)
	return rv
}

func (p_ PathControlDelegateWrapper) ImplementsPathControlAcceptDrop() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:acceptDrop:"))
}

func (p_ PathControlDelegateWrapper) PathControlAcceptDrop(pathControl IPathControl, info IDraggingInfo) bool {
	po := objc.WrapAsProtocol("NSDraggingInfo", info)
	rv := objc.CallMethod[bool](p_, objc.GetSelector("pathControl:acceptDrop:"), objc.ExtractPtr(pathControl), po)
	return rv
}

func (p_ PathControlDelegateWrapper) ImplementsPathControlWillDisplayOpenPanel() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:willDisplayOpenPanel:"))
}

func (p_ PathControlDelegateWrapper) PathControlWillDisplayOpenPanel(pathControl IPathControl, openPanel IOpenPanel) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("pathControl:willDisplayOpenPanel:"), objc.ExtractPtr(pathControl), objc.ExtractPtr(openPanel))
}

func (p_ PathControlDelegateWrapper) ImplementsPathControlWillPopUpMenu() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:willPopUpMenu:"))
}

func (p_ PathControlDelegateWrapper) PathControlWillPopUpMenu(pathControl IPathControl, menu IMenu) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("pathControl:willPopUpMenu:"), objc.ExtractPtr(pathControl), objc.ExtractPtr(menu))
}

func (p_ PathControlDelegateWrapper) ImplementsPathControlShouldDragItemWithPasteboard() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:shouldDragItem:withPasteboard:"))
}

func (p_ PathControlDelegateWrapper) PathControlShouldDragItemWithPasteboard(pathControl IPathControl, pathItem IPathControlItem, pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("pathControl:shouldDragItem:withPasteboard:"), objc.ExtractPtr(pathControl), objc.ExtractPtr(pathItem), objc.ExtractPtr(pasteboard))
	return rv
}
