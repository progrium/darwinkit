// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ToolbarItemGroupClass = _ToolbarItemGroupClass{objc.GetClass("NSToolbarItemGroup")}

type _ToolbarItemGroupClass struct {
	objc.Class
}

type IToolbarItemGroup interface {
	IToolbarItem
	IsSelectedAtIndex(index int) bool
	SetSelectedAtIndex(selected bool, index int)
	Subitems() []ToolbarItem
	SetSubitems(value []IToolbarItem)
	SelectedIndex() int
	SetSelectedIndex(value int)
	ControlRepresentation() ToolbarItemGroupControlRepresentation
	SetControlRepresentation(value ToolbarItemGroupControlRepresentation)
	SelectionMode() ToolbarItemGroupSelectionMode
	SetSelectionMode(value ToolbarItemGroupSelectionMode)
}

type ToolbarItemGroup struct {
	ToolbarItem
}

func MakeToolbarItemGroup(ptr unsafe.Pointer) ToolbarItemGroup {
	return ToolbarItemGroup{
		ToolbarItem: MakeToolbarItem(ptr),
	}
}

func (tc _ToolbarItemGroupClass) GroupWithItemIdentifierImagesSelectionModeLabelsTargetAction(itemIdentifier ToolbarItemIdentifier, images []IImage, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.Selector) ToolbarItemGroup {
	rv := objc.CallMethod[ToolbarItemGroup](tc, objc.GetSelector("groupWithItemIdentifier:images:selectionMode:labels:target:action:"), itemIdentifier, images, selectionMode, labels, objc.ExtractPtr(target), action)
	return rv
}

func ToolbarItemGroup_GroupWithItemIdentifierImagesSelectionModeLabelsTargetAction(itemIdentifier ToolbarItemIdentifier, images []IImage, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.Selector) ToolbarItemGroup {
	return ToolbarItemGroupClass.GroupWithItemIdentifierImagesSelectionModeLabelsTargetAction(itemIdentifier, images, selectionMode, labels, target, action)
}

func (tc _ToolbarItemGroupClass) GroupWithItemIdentifierTitlesSelectionModeLabelsTargetAction(itemIdentifier ToolbarItemIdentifier, titles []string, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.Selector) ToolbarItemGroup {
	rv := objc.CallMethod[ToolbarItemGroup](tc, objc.GetSelector("groupWithItemIdentifier:titles:selectionMode:labels:target:action:"), itemIdentifier, titles, selectionMode, labels, objc.ExtractPtr(target), action)
	return rv
}

func ToolbarItemGroup_GroupWithItemIdentifierTitlesSelectionModeLabelsTargetAction(itemIdentifier ToolbarItemIdentifier, titles []string, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.Selector) ToolbarItemGroup {
	return ToolbarItemGroupClass.GroupWithItemIdentifierTitlesSelectionModeLabelsTargetAction(itemIdentifier, titles, selectionMode, labels, target, action)
}

func (t_ ToolbarItemGroup) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) ToolbarItemGroup {
	rv := objc.CallMethod[ToolbarItemGroup](t_, objc.GetSelector("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

func ToolbarItemGroup_InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) ToolbarItemGroup {
	return ToolbarItemGroupClass.Alloc().InitWithItemIdentifier(itemIdentifier)
}

func (tc _ToolbarItemGroupClass) Alloc() ToolbarItemGroup {
	rv := objc.CallMethod[ToolbarItemGroup](tc, objc.GetSelector("alloc"))
	return rv
}

func ToolbarItemGroup_Alloc() ToolbarItemGroup {
	return ToolbarItemGroupClass.Alloc()
}

func (tc _ToolbarItemGroupClass) New() ToolbarItemGroup {
	rv := objc.CallMethod[ToolbarItemGroup](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewToolbarItemGroup() ToolbarItemGroup {
	return ToolbarItemGroupClass.New()
}

func ToolbarItemGroup_New() ToolbarItemGroup {
	return ToolbarItemGroupClass.New()
}

func (t_ ToolbarItemGroup) Init() ToolbarItemGroup {
	rv := objc.CallMethod[ToolbarItemGroup](t_, objc.GetSelector("init"))
	return rv
}

func ToolbarItemGroup_Init() ToolbarItemGroup {
	return ToolbarItemGroupClass.Alloc().Init()
}

func (t_ ToolbarItemGroup) IsSelectedAtIndex(index int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isSelectedAtIndex:"), index)
	return rv
}

func (t_ ToolbarItemGroup) SetSelectedAtIndex(selected bool, index int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelected:atIndex:"), selected, index)
}

func (t_ ToolbarItemGroup) Subitems() []ToolbarItem {
	rv := objc.CallMethod[[]ToolbarItem](t_, objc.GetSelector("subitems"))
	// TODO: convert slice items...
	return rv
}

func (t_ ToolbarItemGroup) SetSubitems(value []IToolbarItem) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSubitems:"), value)
}

func (t_ ToolbarItemGroup) SelectedIndex() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("selectedIndex"))
	return rv
}

func (t_ ToolbarItemGroup) SetSelectedIndex(value int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectedIndex:"), value)
}

func (t_ ToolbarItemGroup) ControlRepresentation() ToolbarItemGroupControlRepresentation {
	rv := objc.CallMethod[ToolbarItemGroupControlRepresentation](t_, objc.GetSelector("controlRepresentation"))
	return rv
}

func (t_ ToolbarItemGroup) SetControlRepresentation(value ToolbarItemGroupControlRepresentation) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setControlRepresentation:"), value)
}

func (t_ ToolbarItemGroup) SelectionMode() ToolbarItemGroupSelectionMode {
	rv := objc.CallMethod[ToolbarItemGroupSelectionMode](t_, objc.GetSelector("selectionMode"))
	return rv
}

func (t_ ToolbarItemGroup) SetSelectionMode(value ToolbarItemGroupSelectionMode) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectionMode:"), value)
}
