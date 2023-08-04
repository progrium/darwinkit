// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var RuleEditorClass = _RuleEditorClass{objc.GetClass("NSRuleEditor")}

type _RuleEditorClass struct {
	objc.Class
}

type IRuleEditor interface {
	IControl
	ReloadCriteria()
	SetCriteriaAndDisplayValuesForRowAtIndex(criteria []objc.IObject, values []objc.IObject, rowIndex int)
	CriteriaForRow(row int) []objc.Object
	DisplayValuesForRow(row int) []objc.Object
	ParentRowForRow(rowIndex int) int
	RowForDisplayValue(displayValue objc.IObject) int
	RowTypeForRow(rowIndex int) RuleEditorRowType
	SubrowIndexesForRow(rowIndex int) foundation.IndexSet
	SelectRowIndexesByExtendingSelection(indexes foundation.IIndexSet, extend bool)
	AddRow(sender objc.IObject)
	InsertRowAtIndexWithTypeAsSubrowOfRowAnimate(rowIndex int, rowType RuleEditorRowType, parentRow int, shouldAnimate bool)
	RemoveRowAtIndex(rowIndex int)
	RemoveRowsAtIndexesIncludeSubrows(rowIndexes foundation.IIndexSet, includeSubrows bool)
	ReloadPredicate()
	PredicateForRow(row int) foundation.Predicate
	Delegate() RuleEditorDelegateWrapper
	SetDelegate(value IRuleEditorDelegate)
	SetDelegate0(value objc.IObject)
	IsEditable() bool
	SetEditable(value bool)
	NestingMode() RuleEditorNestingMode
	SetNestingMode(value RuleEditorNestingMode)
	CanRemoveAllRows() bool
	SetCanRemoveAllRows(value bool)
	RowHeight() float64
	SetRowHeight(value float64)
	FormattingDictionary() map[string]string
	SetFormattingDictionary(value map[string]string)
	FormattingStringsFilename() string
	SetFormattingStringsFilename(value string)
	NumberOfRows() int
	SelectedRowIndexes() foundation.IndexSet
	Predicate() foundation.Predicate
	RowClass() objc.Class
	SetRowClass(value objc.IClass)
	RowTypeKeyPath() string
	SetRowTypeKeyPath(value string)
	SubrowsKeyPath() string
	SetSubrowsKeyPath(value string)
	CriteriaKeyPath() string
	SetCriteriaKeyPath(value string)
	DisplayValuesKeyPath() string
	SetDisplayValuesKeyPath(value string)
}

type RuleEditor struct {
	Control
}

func MakeRuleEditor(ptr unsafe.Pointer) RuleEditor {
	return RuleEditor{
		Control: MakeControl(ptr),
	}
}

func (r_ RuleEditor) InitWithFrame(frameRect foundation.Rect) RuleEditor {
	rv := objc.CallMethod[RuleEditor](r_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func RuleEditor_InitWithFrame(frameRect foundation.Rect) RuleEditor {
	return RuleEditorClass.Alloc().InitWithFrame(frameRect)
}

func (r_ RuleEditor) Init() RuleEditor {
	rv := objc.CallMethod[RuleEditor](r_, objc.GetSelector("init"))
	return rv
}

func RuleEditor_Init() RuleEditor {
	return RuleEditorClass.Alloc().Init()
}

func (rc _RuleEditorClass) Alloc() RuleEditor {
	rv := objc.CallMethod[RuleEditor](rc, objc.GetSelector("alloc"))
	return rv
}

func RuleEditor_Alloc() RuleEditor {
	return RuleEditorClass.Alloc()
}

func (rc _RuleEditorClass) New() RuleEditor {
	rv := objc.CallMethod[RuleEditor](rc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewRuleEditor() RuleEditor {
	return RuleEditorClass.New()
}

func RuleEditor_New() RuleEditor {
	return RuleEditorClass.New()
}

func (r_ RuleEditor) ReloadCriteria() {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("reloadCriteria"))
}

func (r_ RuleEditor) SetCriteriaAndDisplayValuesForRowAtIndex(criteria []objc.IObject, values []objc.IObject, rowIndex int) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setCriteria:andDisplayValues:forRowAtIndex:"), criteria, values, rowIndex)
}

func (r_ RuleEditor) CriteriaForRow(row int) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](r_, objc.GetSelector("criteriaForRow:"), row)
	return rv
}

func (r_ RuleEditor) DisplayValuesForRow(row int) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](r_, objc.GetSelector("displayValuesForRow:"), row)
	return rv
}

func (r_ RuleEditor) ParentRowForRow(rowIndex int) int {
	rv := objc.CallMethod[int](r_, objc.GetSelector("parentRowForRow:"), rowIndex)
	return rv
}

func (r_ RuleEditor) RowForDisplayValue(displayValue objc.IObject) int {
	rv := objc.CallMethod[int](r_, objc.GetSelector("rowForDisplayValue:"), objc.ExtractPtr(displayValue))
	return rv
}

func (r_ RuleEditor) RowTypeForRow(rowIndex int) RuleEditorRowType {
	rv := objc.CallMethod[RuleEditorRowType](r_, objc.GetSelector("rowTypeForRow:"), rowIndex)
	return rv
}

func (r_ RuleEditor) SubrowIndexesForRow(rowIndex int) foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](r_, objc.GetSelector("subrowIndexesForRow:"), rowIndex)
	return rv
}

func (r_ RuleEditor) SelectRowIndexesByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("selectRowIndexes:byExtendingSelection:"), objc.ExtractPtr(indexes), extend)
}

func (r_ RuleEditor) AddRow(sender objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("addRow:"), objc.ExtractPtr(sender))
}

func (r_ RuleEditor) InsertRowAtIndexWithTypeAsSubrowOfRowAnimate(rowIndex int, rowType RuleEditorRowType, parentRow int, shouldAnimate bool) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("insertRowAtIndex:withType:asSubrowOfRow:animate:"), rowIndex, rowType, parentRow, shouldAnimate)
}

func (r_ RuleEditor) RemoveRowAtIndex(rowIndex int) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("removeRowAtIndex:"), rowIndex)
}

func (r_ RuleEditor) RemoveRowsAtIndexesIncludeSubrows(rowIndexes foundation.IIndexSet, includeSubrows bool) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("removeRowsAtIndexes:includeSubrows:"), objc.ExtractPtr(rowIndexes), includeSubrows)
}

func (r_ RuleEditor) ReloadPredicate() {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("reloadPredicate"))
}

func (r_ RuleEditor) PredicateForRow(row int) foundation.Predicate {
	rv := objc.CallMethod[foundation.Predicate](r_, objc.GetSelector("predicateForRow:"), row)
	return rv
}

func (r_ RuleEditor) Delegate() RuleEditorDelegateWrapper {
	rv := objc.CallMethod[RuleEditorDelegateWrapper](r_, objc.GetSelector("delegate"))
	return rv
}

func (r_ RuleEditor) SetDelegate(value IRuleEditorDelegate) {
	po := objc.WrapAsProtocol("NSRuleEditorDelegate", value)
	objc.SetAssociatedObject(r_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setDelegate:"), po)
}

func (r_ RuleEditor) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (r_ RuleEditor) IsEditable() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("isEditable"))
	return rv
}

func (r_ RuleEditor) SetEditable(value bool) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setEditable:"), value)
}

func (r_ RuleEditor) NestingMode() RuleEditorNestingMode {
	rv := objc.CallMethod[RuleEditorNestingMode](r_, objc.GetSelector("nestingMode"))
	return rv
}

func (r_ RuleEditor) SetNestingMode(value RuleEditorNestingMode) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setNestingMode:"), value)
}

func (r_ RuleEditor) CanRemoveAllRows() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("canRemoveAllRows"))
	return rv
}

func (r_ RuleEditor) SetCanRemoveAllRows(value bool) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setCanRemoveAllRows:"), value)
}

func (r_ RuleEditor) RowHeight() float64 {
	rv := objc.CallMethod[float64](r_, objc.GetSelector("rowHeight"))
	return rv
}

func (r_ RuleEditor) SetRowHeight(value float64) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setRowHeight:"), value)
}

func (r_ RuleEditor) FormattingDictionary() map[string]string {
	rv := objc.CallMethod[map[string]string](r_, objc.GetSelector("formattingDictionary"))
	return rv
}

func (r_ RuleEditor) SetFormattingDictionary(value map[string]string) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setFormattingDictionary:"), value)
}

func (r_ RuleEditor) FormattingStringsFilename() string {
	rv := objc.CallMethod[string](r_, objc.GetSelector("formattingStringsFilename"))
	return rv
}

func (r_ RuleEditor) SetFormattingStringsFilename(value string) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setFormattingStringsFilename:"), value)
}

func (r_ RuleEditor) NumberOfRows() int {
	rv := objc.CallMethod[int](r_, objc.GetSelector("numberOfRows"))
	return rv
}

func (r_ RuleEditor) SelectedRowIndexes() foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](r_, objc.GetSelector("selectedRowIndexes"))
	return rv
}

func (r_ RuleEditor) Predicate() foundation.Predicate {
	rv := objc.CallMethod[foundation.Predicate](r_, objc.GetSelector("predicate"))
	return rv
}

func (r_ RuleEditor) RowClass() objc.Class {
	rv := objc.CallMethod[objc.Class](r_, objc.GetSelector("rowClass"))
	return rv
}

func (r_ RuleEditor) SetRowClass(value objc.IClass) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setRowClass:"), objc.ExtractPtr(value))
}

func (r_ RuleEditor) RowTypeKeyPath() string {
	rv := objc.CallMethod[string](r_, objc.GetSelector("rowTypeKeyPath"))
	return rv
}

func (r_ RuleEditor) SetRowTypeKeyPath(value string) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setRowTypeKeyPath:"), value)
}

func (r_ RuleEditor) SubrowsKeyPath() string {
	rv := objc.CallMethod[string](r_, objc.GetSelector("subrowsKeyPath"))
	return rv
}

func (r_ RuleEditor) SetSubrowsKeyPath(value string) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setSubrowsKeyPath:"), value)
}

func (r_ RuleEditor) CriteriaKeyPath() string {
	rv := objc.CallMethod[string](r_, objc.GetSelector("criteriaKeyPath"))
	return rv
}

func (r_ RuleEditor) SetCriteriaKeyPath(value string) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setCriteriaKeyPath:"), value)
}

func (r_ RuleEditor) DisplayValuesKeyPath() string {
	rv := objc.CallMethod[string](r_, objc.GetSelector("displayValuesKeyPath"))
	return rv
}

func (r_ RuleEditor) SetDisplayValuesKeyPath(value string) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setDisplayValuesKeyPath:"), value)
}
