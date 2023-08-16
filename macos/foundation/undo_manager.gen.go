// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UndoManager] class.
var UndoManagerClass = _UndoManagerClass{objc.GetClass("NSUndoManager")}

type _UndoManagerClass struct {
	objc.Class
}

// An interface definition for the [UndoManager] class.
type IUndoManager interface {
	objc.IObject
	SetActionName(actionName string)
	UndoNestedGroup()
	RegisterUndoWithTargetHandler(target objc.IObject, undoHandler func(target objc.Object))
	RedoMenuTitleForUndoActionName(actionName string) string
	BeginUndoGrouping()
	EnableUndoRegistration()
	RemoveAllActionsWithTarget(target objc.IObject)
	EndUndoGrouping()
	SetActionIsDiscardable(discardable bool)
	UndoMenuTitleForUndoActionName(actionName string) string
	Undo()
	DisableUndoRegistration()
	RemoveAllActions()
	PrepareWithInvocationTarget(target objc.IObject) objc.Object
	Redo()
	LevelsOfUndo() uint
	SetLevelsOfUndo(value uint)
	GroupsByEvent() bool
	SetGroupsByEvent(value bool)
	CanUndo() bool
	RedoActionIsDiscardable() bool
	RedoActionName() string
	UndoMenuItemTitle() string
	IsUndoRegistrationEnabled() bool
	UndoActionName() string
	CanRedo() bool
	RedoMenuItemTitle() string
	RunLoopModes() []RunLoopMode
	SetRunLoopModes(value []RunLoopMode)
	IsUndoing() bool
	UndoActionIsDiscardable() bool
	GroupingLevel() int
	IsRedoing() bool
}

// A general-purpose recorder of operations that enables undo and redo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager?language=objc
type UndoManager struct {
	objc.Object
}

func UndoManagerFrom(ptr unsafe.Pointer) UndoManager {
	return UndoManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _UndoManagerClass) Alloc() UndoManager {
	rv := objc.Call[UndoManager](uc, objc.Sel("alloc"))
	return rv
}

func UndoManager_Alloc() UndoManager {
	return UndoManagerClass.Alloc()
}

func (uc _UndoManagerClass) New() UndoManager {
	rv := objc.Call[UndoManager](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUndoManager() UndoManager {
	return UndoManagerClass.New()
}

func (u_ UndoManager) Init() UndoManager {
	rv := objc.Call[UndoManager](u_, objc.Sel("init"))
	return rv
}

// Sets the name of the action associated with the Undo or Redo command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1412915-setactionname?language=objc
func (u_ UndoManager) SetActionName(actionName string) {
	objc.Call[objc.Void](u_, objc.Sel("setActionName:"), actionName)
}

// Performs the undo operations in the last undo group (whether top-level or nested), recording the operations on the redo stack as a single group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1410826-undonestedgroup?language=objc
func (u_ UndoManager) UndoNestedGroup() {
	objc.Call[objc.Void](u_, objc.Sel("undoNestedGroup"))
}

// Records a single undo operation for a given target so that when an undo is performed, it executes the specified block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1437863-registerundowithtarget?language=objc
func (u_ UndoManager) RegisterUndoWithTargetHandler(target objc.IObject, undoHandler func(target objc.Object)) {
	objc.Call[objc.Void](u_, objc.Sel("registerUndoWithTarget:handler:"), target, undoHandler)
}

// Returns the complete, localized title of the Redo menu command for the action identified by the given name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1407438-redomenutitleforundoactionname?language=objc
func (u_ UndoManager) RedoMenuTitleForUndoActionName(actionName string) string {
	rv := objc.Call[string](u_, objc.Sel("redoMenuTitleForUndoActionName:"), actionName)
	return rv
}

// Marks the beginning of an undo group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1409894-beginundogrouping?language=objc
func (u_ UndoManager) BeginUndoGrouping() {
	objc.Call[objc.Void](u_, objc.Sel("beginUndoGrouping"))
}

// Enables the recording of undo operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1408957-enableundoregistration?language=objc
func (u_ UndoManager) EnableUndoRegistration() {
	objc.Call[objc.Void](u_, objc.Sel("enableUndoRegistration"))
}

// Clears the undo and redo stacks of all operations involving the specified target as the recipient of the undo message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1409896-removeallactionswithtarget?language=objc
func (u_ UndoManager) RemoveAllActionsWithTarget(target objc.IObject) {
	objc.Call[objc.Void](u_, objc.Sel("removeAllActionsWithTarget:"), target)
}

// Marks the end of an undo group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1416490-endundogrouping?language=objc
func (u_ UndoManager) EndUndoGrouping() {
	objc.Call[objc.Void](u_, objc.Sel("endUndoGrouping"))
}

// Sets whether the next undo or redo action is discardable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1412159-setactionisdiscardable?language=objc
func (u_ UndoManager) SetActionIsDiscardable(discardable bool) {
	objc.Call[objc.Void](u_, objc.Sel("setActionIsDiscardable:"), discardable)
}

// Returns the complete, localized title of the Undo menu command for the action identified by the given name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1413122-undomenutitleforundoactionname?language=objc
func (u_ UndoManager) UndoMenuTitleForUndoActionName(actionName string) string {
	rv := objc.Call[string](u_, objc.Sel("undoMenuTitleForUndoActionName:"), actionName)
	return rv
}

// Closes the top-level undo group if necessary and invokes undoNestedGroup. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1412189-undo?language=objc
func (u_ UndoManager) Undo() {
	objc.Call[objc.Void](u_, objc.Sel("undo"))
}

// Disables the recording of undo operations, whether by registerUndoWithTarget:handler: or by invocation-based undo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1412239-disableundoregistration?language=objc
func (u_ UndoManager) DisableUndoRegistration() {
	objc.Call[objc.Void](u_, objc.Sel("disableUndoRegistration"))
}

// Clears the undo and redo stacks and re-enables the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1407442-removeallactions?language=objc
func (u_ UndoManager) RemoveAllActions() {
	objc.Call[objc.Void](u_, objc.Sel("removeAllActions"))
}

// Prepares the undo manager for invocation-based undo with the given target as the subject of the next undo operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1409564-preparewithinvocationtarget?language=objc
func (u_ UndoManager) PrepareWithInvocationTarget(target objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](u_, objc.Sel("prepareWithInvocationTarget:"), target)
	return rv
}

// Performs the operations in the last group on the redo stack, if there are any, recording them on the undo stack as a single group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1417030-redo?language=objc
func (u_ UndoManager) Redo() {
	objc.Call[objc.Void](u_, objc.Sel("redo"))
}

// The maximum number of top-level undo groups the receiver holds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1409753-levelsofundo?language=objc
func (u_ UndoManager) LevelsOfUndo() uint {
	rv := objc.Call[uint](u_, objc.Sel("levelsOfUndo"))
	return rv
}

// The maximum number of top-level undo groups the receiver holds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1409753-levelsofundo?language=objc
func (u_ UndoManager) SetLevelsOfUndo(value uint) {
	objc.Call[objc.Void](u_, objc.Sel("setLevelsOfUndo:"), value)
}

// A Boolean value that indicates whether the receiver automatically creates undo groups around each pass of the run loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1417407-groupsbyevent?language=objc
func (u_ UndoManager) GroupsByEvent() bool {
	rv := objc.Call[bool](u_, objc.Sel("groupsByEvent"))
	return rv
}

// A Boolean value that indicates whether the receiver automatically creates undo groups around each pass of the run loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1417407-groupsbyevent?language=objc
func (u_ UndoManager) SetGroupsByEvent(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setGroupsByEvent:"), value)
}

// A Boolean value that indicates whether the receiver has any actions to undo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1412394-canundo?language=objc
func (u_ UndoManager) CanUndo() bool {
	rv := objc.Call[bool](u_, objc.Sel("canUndo"))
	return rv
}

// Boolean value that indicates whether the next redo action is discardable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1413488-redoactionisdiscardable?language=objc
func (u_ UndoManager) RedoActionIsDiscardable() bool {
	rv := objc.Call[bool](u_, objc.Sel("redoActionIsDiscardable"))
	return rv
}

// The name identifying the redo action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1417487-redoactionname?language=objc
func (u_ UndoManager) RedoActionName() string {
	rv := objc.Call[string](u_, objc.Sel("redoActionName"))
	return rv
}

// The complete title of the Undo menu command, for example, “Undo Paste.” [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1412953-undomenuitemtitle?language=objc
func (u_ UndoManager) UndoMenuItemTitle() string {
	rv := objc.Call[string](u_, objc.Sel("undoMenuItemTitle"))
	return rv
}

// A Boolean value that indicates whether the recording of undo operations is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1415464-undoregistrationenabled?language=objc
func (u_ UndoManager) IsUndoRegistrationEnabled() bool {
	rv := objc.Call[bool](u_, objc.Sel("isUndoRegistrationEnabled"))
	return rv
}

// The name identifying the undo action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1415127-undoactionname?language=objc
func (u_ UndoManager) UndoActionName() string {
	rv := objc.Call[string](u_, objc.Sel("undoActionName"))
	return rv
}

// A Boolean value that indicates whether the receiver has any actions to redo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1415212-canredo?language=objc
func (u_ UndoManager) CanRedo() bool {
	rv := objc.Call[bool](u_, objc.Sel("canRedo"))
	return rv
}

// The complete title of the Redo menu command, for example, “Redo Paste.” [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1409938-redomenuitemtitle?language=objc
func (u_ UndoManager) RedoMenuItemTitle() string {
	rv := objc.Call[string](u_, objc.Sel("redoMenuItemTitle"))
	return rv
}

// The modes governing the types of input handled during a cycle of the run loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1409504-runloopmodes?language=objc
func (u_ UndoManager) RunLoopModes() []RunLoopMode {
	rv := objc.Call[[]RunLoopMode](u_, objc.Sel("runLoopModes"))
	return rv
}

// The modes governing the types of input handled during a cycle of the run loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1409504-runloopmodes?language=objc
func (u_ UndoManager) SetRunLoopModes(value []RunLoopMode) {
	objc.Call[objc.Void](u_, objc.Sel("setRunLoopModes:"), value)
}

// Returns a Boolean value that indicates whether the receiver is in the process of performing its undo or undoNestedGroup method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1407345-undoing?language=objc
func (u_ UndoManager) IsUndoing() bool {
	rv := objc.Call[bool](u_, objc.Sel("isUndoing"))
	return rv
}

// Boolean value that indicates whether the next undo action is discardable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1415261-undoactionisdiscardable?language=objc
func (u_ UndoManager) UndoActionIsDiscardable() bool {
	rv := objc.Call[bool](u_, objc.Sel("undoActionIsDiscardable"))
	return rv
}

// The number of nested undo groups (or redo groups, if Redo was invoked last) in the current event loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1409006-groupinglevel?language=objc
func (u_ UndoManager) GroupingLevel() int {
	rv := objc.Call[int](u_, objc.Sel("groupingLevel"))
	return rv
}

// Returns a Boolean value that indicates whether the receiver is in the process of performing its redo method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanager/1411654-redoing?language=objc
func (u_ UndoManager) IsRedoing() bool {
	rv := objc.Call[bool](u_, objc.Sel("isRedoing"))
	return rv
}
