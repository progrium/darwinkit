// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var OutlineViewClass = _OutlineViewClass{objc.GetClass("NSOutlineView")}

type _OutlineViewClass struct {
	objc.Class
}

type IOutlineView interface {
	ITableView
	IsExpandable(item objc.IObject) bool
	IsItemExpanded(item objc.IObject) bool
	ExpandItem(item objc.IObject)
	ExpandItemExpandChildren(item objc.IObject, expandChildren bool)
	CollapseItem(item objc.IObject)
	CollapseItemCollapseChildren(item objc.IObject, collapseChildren bool)
	ReloadItem(item objc.IObject)
	ReloadItemReloadChildren(item objc.IObject, reloadChildren bool)
	ItemAtRow(row int) objc.Object
	RowForItem(item objc.IObject) int
	LevelForItem(item objc.IObject) int
	LevelForRow(row int) int
	SetDropItemDropChildIndex(item objc.IObject, index int)
	ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool
	ParentForItem(item objc.IObject) objc.Object
	ChildIndexForItem(item objc.IObject) int
	ChildOfItem(index int, item objc.IObject) objc.Object
	NumberOfChildrenOfItem(item objc.IObject) int
	FrameOfOutlineCellAtRow(row int) foundation.Rect
	InsertItemsAtIndexesInParentWithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions)
	MoveItemAtIndexInParentToIndexInParent(fromIndex int, oldParent objc.IObject, toIndex int, newParent objc.IObject)
	RemoveItemsAtIndexesInParentWithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions)
	StronglyReferencesItems() bool
	SetStronglyReferencesItems(value bool)
	OutlineTableColumn() TableColumn
	SetOutlineTableColumn(value ITableColumn)
	AutoresizesOutlineColumn() bool
	SetAutoresizesOutlineColumn(value bool)
	IndentationPerLevel() float64
	SetIndentationPerLevel(value float64)
	IndentationMarkerFollowsCell() bool
	SetIndentationMarkerFollowsCell(value bool)
	AutosaveExpandedItems() bool
	SetAutosaveExpandedItems(value bool)
}

type OutlineView struct {
	TableView
}

func MakeOutlineView(ptr unsafe.Pointer) OutlineView {
	return OutlineView{
		TableView: MakeTableView(ptr),
	}
}

func (o_ OutlineView) InitWithFrame(frameRect foundation.Rect) OutlineView {
	rv := objc.CallMethod[OutlineView](o_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func OutlineView_InitWithFrame(frameRect foundation.Rect) OutlineView {
	return OutlineViewClass.Alloc().InitWithFrame(frameRect)
}

func (o_ OutlineView) Init() OutlineView {
	rv := objc.CallMethod[OutlineView](o_, objc.GetSelector("init"))
	return rv
}

func OutlineView_Init() OutlineView {
	return OutlineViewClass.Alloc().Init()
}

func (oc _OutlineViewClass) Alloc() OutlineView {
	rv := objc.CallMethod[OutlineView](oc, objc.GetSelector("alloc"))
	return rv
}

func OutlineView_Alloc() OutlineView {
	return OutlineViewClass.Alloc()
}

func (oc _OutlineViewClass) New() OutlineView {
	rv := objc.CallMethod[OutlineView](oc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewOutlineView() OutlineView {
	return OutlineViewClass.New()
}

func OutlineView_New() OutlineView {
	return OutlineViewClass.New()
}

func (o_ OutlineView) IsExpandable(item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("isExpandable:"), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineView) IsItemExpanded(item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("isItemExpanded:"), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineView) ExpandItem(item objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("expandItem:"), objc.ExtractPtr(item))
}

func (o_ OutlineView) ExpandItemExpandChildren(item objc.IObject, expandChildren bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("expandItem:expandChildren:"), objc.ExtractPtr(item), expandChildren)
}

func (o_ OutlineView) CollapseItem(item objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("collapseItem:"), objc.ExtractPtr(item))
}

func (o_ OutlineView) CollapseItemCollapseChildren(item objc.IObject, collapseChildren bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("collapseItem:collapseChildren:"), objc.ExtractPtr(item), collapseChildren)
}

func (o_ OutlineView) ReloadItem(item objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("reloadItem:"), objc.ExtractPtr(item))
}

func (o_ OutlineView) ReloadItemReloadChildren(item objc.IObject, reloadChildren bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("reloadItem:reloadChildren:"), objc.ExtractPtr(item), reloadChildren)
}

func (o_ OutlineView) ItemAtRow(row int) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("itemAtRow:"), row)
	return rv
}

func (o_ OutlineView) RowForItem(item objc.IObject) int {
	rv := objc.CallMethod[int](o_, objc.GetSelector("rowForItem:"), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineView) LevelForItem(item objc.IObject) int {
	rv := objc.CallMethod[int](o_, objc.GetSelector("levelForItem:"), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineView) LevelForRow(row int) int {
	rv := objc.CallMethod[int](o_, objc.GetSelector("levelForRow:"), row)
	return rv
}

func (o_ OutlineView) SetDropItemDropChildIndex(item objc.IObject, index int) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setDropItem:dropChildIndex:"), objc.ExtractPtr(item), index)
}

func (o_ OutlineView) ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("shouldCollapseAutoExpandedItemsForDeposited:"), deposited)
	return rv
}

func (o_ OutlineView) ParentForItem(item objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("parentForItem:"), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineView) ChildIndexForItem(item objc.IObject) int {
	rv := objc.CallMethod[int](o_, objc.GetSelector("childIndexForItem:"), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineView) ChildOfItem(index int, item objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("child:ofItem:"), index, objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineView) NumberOfChildrenOfItem(item objc.IObject) int {
	rv := objc.CallMethod[int](o_, objc.GetSelector("numberOfChildrenOfItem:"), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineView) FrameOfOutlineCellAtRow(row int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](o_, objc.GetSelector("frameOfOutlineCellAtRow:"), row)
	return rv
}

func (o_ OutlineView) InsertItemsAtIndexesInParentWithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("insertItemsAtIndexes:inParent:withAnimation:"), objc.ExtractPtr(indexes), objc.ExtractPtr(parent), animationOptions)
}

func (o_ OutlineView) MoveItemAtIndexInParentToIndexInParent(fromIndex int, oldParent objc.IObject, toIndex int, newParent objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("moveItemAtIndex:inParent:toIndex:inParent:"), fromIndex, objc.ExtractPtr(oldParent), toIndex, objc.ExtractPtr(newParent))
}

func (o_ OutlineView) RemoveItemsAtIndexesInParentWithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("removeItemsAtIndexes:inParent:withAnimation:"), objc.ExtractPtr(indexes), objc.ExtractPtr(parent), animationOptions)
}

func (o_ OutlineView) StronglyReferencesItems() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("stronglyReferencesItems"))
	return rv
}

func (o_ OutlineView) SetStronglyReferencesItems(value bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setStronglyReferencesItems:"), value)
}

func (o_ OutlineView) OutlineTableColumn() TableColumn {
	rv := objc.CallMethod[TableColumn](o_, objc.GetSelector("outlineTableColumn"))
	return rv
}

func (o_ OutlineView) SetOutlineTableColumn(value ITableColumn) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setOutlineTableColumn:"), objc.ExtractPtr(value))
}

func (o_ OutlineView) AutoresizesOutlineColumn() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("autoresizesOutlineColumn"))
	return rv
}

func (o_ OutlineView) SetAutoresizesOutlineColumn(value bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setAutoresizesOutlineColumn:"), value)
}

func (o_ OutlineView) IndentationPerLevel() float64 {
	rv := objc.CallMethod[float64](o_, objc.GetSelector("indentationPerLevel"))
	return rv
}

func (o_ OutlineView) SetIndentationPerLevel(value float64) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setIndentationPerLevel:"), value)
}

func (o_ OutlineView) IndentationMarkerFollowsCell() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("indentationMarkerFollowsCell"))
	return rv
}

func (o_ OutlineView) SetIndentationMarkerFollowsCell(value bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setIndentationMarkerFollowsCell:"), value)
}

func (o_ OutlineView) AutosaveExpandedItems() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("autosaveExpandedItems"))
	return rv
}

func (o_ OutlineView) SetAutosaveExpandedItems(value bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setAutosaveExpandedItems:"), value)
}
