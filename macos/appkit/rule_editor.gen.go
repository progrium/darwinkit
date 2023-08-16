// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RuleEditor] class.
var RuleEditorClass = _RuleEditorClass{objc.GetClass("NSRuleEditor")}

type _RuleEditorClass struct {
	objc.Class
}

// An interface definition for the [RuleEditor] class.
type IRuleEditor interface {
	IControl
	CriteriaForRow(row int) []objc.Object
	AddRow(sender objc.IObject)
	RowTypeForRow(rowIndex int) RuleEditorRowType
	RowForDisplayValue(displayValue objc.IObject) int
	SetCriteriaAndDisplayValuesForRowAtIndex(criteria []objc.IObject, values []objc.IObject, rowIndex int)
	ParentRowForRow(rowIndex int) int
	ReloadCriteria()
	RemoveRowsAtIndexesIncludeSubrows(rowIndexes foundation.IIndexSet, includeSubrows bool)
	SubrowIndexesForRow(rowIndex int) foundation.IndexSet
	InsertRowAtIndexWithTypeAsSubrowOfRowAnimate(rowIndex int, rowType RuleEditorRowType, parentRow int, shouldAnimate bool)
	PredicateForRow(row int) foundation.Predicate
	SelectRowIndexesByExtendingSelection(indexes foundation.IIndexSet, extend bool)
	DisplayValuesForRow(row int) []objc.Object
	ReloadPredicate()
	RemoveRowAtIndex(rowIndex int)
	FormattingStringsFilename() string
	SetFormattingStringsFilename(value string)
	CanRemoveAllRows() bool
	SetCanRemoveAllRows(value bool)
	RowHeight() float64
	SetRowHeight(value float64)
	IsEditable() bool
	SetEditable(value bool)
	RowClass() objc.Class
	SetRowClass(value objc.IClass)
	Delegate() RuleEditorDelegateWrapper
	SetDelegate(value PRuleEditorDelegate)
	SetDelegateObject(valueObject objc.IObject)
	Predicate() foundation.Predicate
	SubrowsKeyPath() string
	SetSubrowsKeyPath(value string)
	FormattingDictionary() map[string]string
	SetFormattingDictionary(value map[string]string)
	RowTypeKeyPath() string
	SetRowTypeKeyPath(value string)
	DisplayValuesKeyPath() string
	SetDisplayValuesKeyPath(value string)
	CriteriaKeyPath() string
	SetCriteriaKeyPath(value string)
	NumberOfRows() int
	SelectedRowIndexes() foundation.IndexSet
	NestingMode() RuleEditorNestingMode
	SetNestingMode(value RuleEditorNestingMode)
}

// An interface for configuring a rule-based list of options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor?language=objc
type RuleEditor struct {
	Control
}

func RuleEditorFrom(ptr unsafe.Pointer) RuleEditor {
	return RuleEditor{
		Control: ControlFrom(ptr),
	}
}

func (rc _RuleEditorClass) Alloc() RuleEditor {
	rv := objc.Call[RuleEditor](rc, objc.Sel("alloc"))
	return rv
}

func RuleEditor_Alloc() RuleEditor {
	return RuleEditorClass.Alloc()
}

func (rc _RuleEditorClass) New() RuleEditor {
	rv := objc.Call[RuleEditor](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRuleEditor() RuleEditor {
	return RuleEditorClass.New()
}

func (r_ RuleEditor) Init() RuleEditor {
	rv := objc.Call[RuleEditor](r_, objc.Sel("init"))
	return rv
}

func (r_ RuleEditor) InitWithFrame(frameRect foundation.Rect) RuleEditor {
	rv := objc.Call[RuleEditor](r_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func RuleEditor_InitWithFrame(frameRect foundation.Rect) RuleEditor {
	return RuleEditorClass.Alloc().InitWithFrame(frameRect)
}

// Returns the currently chosen items for a given row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1524906-criteriaforrow?language=objc
func (r_ RuleEditor) CriteriaForRow(row int) []objc.Object {
	rv := objc.Call[[]objc.Object](r_, objc.Sel("criteriaForRow:"), row)
	return rv
}

// Adds a row to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1530067-addrow?language=objc
func (r_ RuleEditor) AddRow(sender objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("addRow:"), sender)
}

// Returns the type of a given row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1530356-rowtypeforrow?language=objc
func (r_ RuleEditor) RowTypeForRow(rowIndex int) RuleEditorRowType {
	rv := objc.Call[RuleEditorRowType](r_, objc.Sel("rowTypeForRow:"), rowIndex)
	return rv
}

// Returns the index of the row containing a given value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1524590-rowfordisplayvalue?language=objc
func (r_ RuleEditor) RowForDisplayValue(displayValue objc.IObject) int {
	rv := objc.Call[int](r_, objc.Sel("rowForDisplayValue:"), displayValue)
	return rv
}

// Modifies the row at a given index to contain the given items and values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1534317-setcriteria?language=objc
func (r_ RuleEditor) SetCriteriaAndDisplayValuesForRowAtIndex(criteria []objc.IObject, values []objc.IObject, rowIndex int) {
	objc.Call[objc.Void](r_, objc.Sel("setCriteria:andDisplayValues:forRowAtIndex:"), criteria, values, rowIndex)
}

// Returns the index of the parent of a given row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1534763-parentrowforrow?language=objc
func (r_ RuleEditor) ParentRowForRow(rowIndex int) int {
	rv := objc.Call[int](r_, objc.Sel("parentRowForRow:"), rowIndex)
	return rv
}

// Instructs the receiver to refetch criteria from its delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1528188-reloadcriteria?language=objc
func (r_ RuleEditor) ReloadCriteria() {
	objc.Call[objc.Void](r_, objc.Sel("reloadCriteria"))
}

// Removes the rows at given indexes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1525947-removerowsatindexes?language=objc
func (r_ RuleEditor) RemoveRowsAtIndexesIncludeSubrows(rowIndexes foundation.IIndexSet, includeSubrows bool) {
	objc.Call[objc.Void](r_, objc.Sel("removeRowsAtIndexes:includeSubrows:"), objc.Ptr(rowIndexes), includeSubrows)
}

// Returns the immediate subrows of a given row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1527898-subrowindexesforrow?language=objc
func (r_ RuleEditor) SubrowIndexesForRow(rowIndex int) foundation.IndexSet {
	rv := objc.Call[foundation.IndexSet](r_, objc.Sel("subrowIndexesForRow:"), rowIndex)
	return rv
}

// Adds a new row of a given type at a given location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1528842-insertrowatindex?language=objc
func (r_ RuleEditor) InsertRowAtIndexWithTypeAsSubrowOfRowAnimate(rowIndex int, rowType RuleEditorRowType, parentRow int, shouldAnimate bool) {
	objc.Call[objc.Void](r_, objc.Sel("insertRowAtIndex:withType:asSubrowOfRow:animate:"), rowIndex, rowType, parentRow, shouldAnimate)
}

// Returns the predicate for a given row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1531272-predicateforrow?language=objc
func (r_ RuleEditor) PredicateForRow(row int) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](r_, objc.Sel("predicateForRow:"), row)
	return rv
}

// Sets in the receiver the indexes of rows that are selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1534559-selectrowindexes?language=objc
func (r_ RuleEditor) SelectRowIndexesByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	objc.Call[objc.Void](r_, objc.Sel("selectRowIndexes:byExtendingSelection:"), objc.Ptr(indexes), extend)
}

// Returns the chosen values for a given row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1529645-displayvaluesforrow?language=objc
func (r_ RuleEditor) DisplayValuesForRow(row int) []objc.Object {
	rv := objc.Call[[]objc.Object](r_, objc.Sel("displayValuesForRow:"), row)
	return rv
}

// Instructs the receiver to regenerate its predicate by invoking the corresponding delegate method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1529215-reloadpredicate?language=objc
func (r_ RuleEditor) ReloadPredicate() {
	objc.Call[objc.Void](r_, objc.Sel("reloadPredicate"))
}

// Removes the row at a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1534953-removerowatindex?language=objc
func (r_ RuleEditor) RemoveRowAtIndex(rowIndex int) {
	objc.Call[objc.Void](r_, objc.Sel("removeRowAtIndex:"), rowIndex)
}

// The name of the rule editor’s strings file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1533323-formattingstringsfilename?language=objc
func (r_ RuleEditor) FormattingStringsFilename() string {
	rv := objc.Call[string](r_, objc.Sel("formattingStringsFilename"))
	return rv
}

// The name of the rule editor’s strings file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1533323-formattingstringsfilename?language=objc
func (r_ RuleEditor) SetFormattingStringsFilename(value string) {
	objc.Call[objc.Void](r_, objc.Sel("setFormattingStringsFilename:"), value)
}

// A Boolean value that indicates whether all the rows can be removed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1535531-canremoveallrows?language=objc
func (r_ RuleEditor) CanRemoveAllRows() bool {
	rv := objc.Call[bool](r_, objc.Sel("canRemoveAllRows"))
	return rv
}

// A Boolean value that indicates whether all the rows can be removed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1535531-canremoveallrows?language=objc
func (r_ RuleEditor) SetCanRemoveAllRows(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setCanRemoveAllRows:"), value)
}

// The rule editor’s row height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1531071-rowheight?language=objc
func (r_ RuleEditor) RowHeight() float64 {
	rv := objc.Call[float64](r_, objc.Sel("rowHeight"))
	return rv
}

// The rule editor’s row height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1531071-rowheight?language=objc
func (r_ RuleEditor) SetRowHeight(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setRowHeight:"), value)
}

// A Boolean value that determines whether the rule editor is editable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1530425-editable?language=objc
func (r_ RuleEditor) IsEditable() bool {
	rv := objc.Call[bool](r_, objc.Sel("isEditable"))
	return rv
}

// A Boolean value that determines whether the rule editor is editable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1530425-editable?language=objc
func (r_ RuleEditor) SetEditable(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setEditable:"), value)
}

// The class used to create a new row in the “rows” binding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1535414-rowclass?language=objc
func (r_ RuleEditor) RowClass() objc.Class {
	rv := objc.Call[objc.Class](r_, objc.Sel("rowClass"))
	return rv
}

// The class used to create a new row in the “rows” binding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1535414-rowclass?language=objc
func (r_ RuleEditor) SetRowClass(value objc.IClass) {
	objc.Call[objc.Void](r_, objc.Sel("setRowClass:"), objc.Ptr(value))
}

// The rule editor’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1528017-delegate?language=objc
func (r_ RuleEditor) Delegate() RuleEditorDelegateWrapper {
	rv := objc.Call[RuleEditorDelegateWrapper](r_, objc.Sel("delegate"))
	return rv
}

// The rule editor’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1528017-delegate?language=objc
func (r_ RuleEditor) SetDelegate(value PRuleEditorDelegate) {
	po0 := objc.WrapAsProtocol("NSRuleEditorDelegate", value)
	objc.SetAssociatedObject(r_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](r_, objc.Sel("setDelegate:"), po0)
}

// The rule editor’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1528017-delegate?language=objc
func (r_ RuleEditor) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The rule editor’s predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1525681-predicate?language=objc
func (r_ RuleEditor) Predicate() foundation.Predicate {
	rv := objc.Call[foundation.Predicate](r_, objc.Sel("predicate"))
	return rv
}

// The key path for the subrows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1535242-subrowskeypath?language=objc
func (r_ RuleEditor) SubrowsKeyPath() string {
	rv := objc.Call[string](r_, objc.Sel("subrowsKeyPath"))
	return rv
}

// The key path for the subrows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1535242-subrowskeypath?language=objc
func (r_ RuleEditor) SetSubrowsKeyPath(value string) {
	objc.Call[objc.Void](r_, objc.Sel("setSubrowsKeyPath:"), value)
}

// The formatting dictionary for the rule editor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1526737-formattingdictionary?language=objc
func (r_ RuleEditor) FormattingDictionary() map[string]string {
	rv := objc.Call[map[string]string](r_, objc.Sel("formattingDictionary"))
	return rv
}

// The formatting dictionary for the rule editor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1526737-formattingdictionary?language=objc
func (r_ RuleEditor) SetFormattingDictionary(value map[string]string) {
	objc.Call[objc.Void](r_, objc.Sel("setFormattingDictionary:"), value)
}

// The key path for the row type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1527020-rowtypekeypath?language=objc
func (r_ RuleEditor) RowTypeKeyPath() string {
	rv := objc.Call[string](r_, objc.Sel("rowTypeKeyPath"))
	return rv
}

// The key path for the row type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1527020-rowtypekeypath?language=objc
func (r_ RuleEditor) SetRowTypeKeyPath(value string) {
	objc.Call[objc.Void](r_, objc.Sel("setRowTypeKeyPath:"), value)
}

// The display values key path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1535572-displayvalueskeypath?language=objc
func (r_ RuleEditor) DisplayValuesKeyPath() string {
	rv := objc.Call[string](r_, objc.Sel("displayValuesKeyPath"))
	return rv
}

// The display values key path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1535572-displayvalueskeypath?language=objc
func (r_ RuleEditor) SetDisplayValuesKeyPath(value string) {
	objc.Call[objc.Void](r_, objc.Sel("setDisplayValuesKeyPath:"), value)
}

// The criteria key path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1524761-criteriakeypath?language=objc
func (r_ RuleEditor) CriteriaKeyPath() string {
	rv := objc.Call[string](r_, objc.Sel("criteriaKeyPath"))
	return rv
}

// The criteria key path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1524761-criteriakeypath?language=objc
func (r_ RuleEditor) SetCriteriaKeyPath(value string) {
	objc.Call[objc.Void](r_, objc.Sel("setCriteriaKeyPath:"), value)
}

// The number of rows in the rule editor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1529330-numberofrows?language=objc
func (r_ RuleEditor) NumberOfRows() int {
	rv := objc.Call[int](r_, objc.Sel("numberOfRows"))
	return rv
}

// The indexes of the rule editor’s selected rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1529841-selectedrowindexes?language=objc
func (r_ RuleEditor) SelectedRowIndexes() foundation.IndexSet {
	rv := objc.Call[foundation.IndexSet](r_, objc.Sel("selectedRowIndexes"))
	return rv
}

// The rule editor’s nesting mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1533955-nestingmode?language=objc
func (r_ RuleEditor) NestingMode() RuleEditorNestingMode {
	rv := objc.Call[RuleEditorNestingMode](r_, objc.Sel("nestingMode"))
	return rv
}

// The rule editor’s nesting mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/1533955-nestingmode?language=objc
func (r_ RuleEditor) SetNestingMode(value RuleEditorNestingMode) {
	objc.Call[objc.Void](r_, objc.Sel("setNestingMode:"), value)
}
