// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var UndoManagerClass = _UndoManagerClass{objc.GetClass("NSUndoManager")}

type _UndoManagerClass struct {
	objc.Class
}

type IUndoManager interface {
	objc.IObject
	RegisterUndoWithTargetHandler(target objc.IObject, undoHandler func(target objc.Object))
	RegisterUndoWithTargetSelectorObject(target objc.IObject, selector objc.Selector, anObject objc.IObject)
	PrepareWithInvocationTarget(target objc.IObject) objc.Object
	Undo()
	UndoNestedGroup()
	Redo()
	BeginUndoGrouping()
	EndUndoGrouping()
	DisableUndoRegistration()
	EnableUndoRegistration()
	RemoveAllActions()
	RemoveAllActionsWithTarget(target objc.IObject)
	SetActionName(actionName string)
	UndoMenuTitleForUndoActionName(actionName string) string
	RedoMenuTitleForUndoActionName(actionName string) string
	SetActionIsDiscardable(discardable bool)
	CanUndo() bool
	CanRedo() bool
	LevelsOfUndo() uint
	SetLevelsOfUndo(value uint)
	GroupsByEvent() bool
	SetGroupsByEvent(value bool)
	GroupingLevel() int
	IsUndoRegistrationEnabled() bool
	IsUndoing() bool
	IsRedoing() bool
	UndoActionName() string
	RedoActionName() string
	UndoMenuItemTitle() string
	RedoMenuItemTitle() string
	RunLoopModes() []RunLoopMode
	SetRunLoopModes(value []RunLoopMode)
	UndoActionIsDiscardable() bool
	RedoActionIsDiscardable() bool
}

type UndoManager struct {
	objc.Object
}

func MakeUndoManager(ptr unsafe.Pointer) UndoManager {
	return UndoManager{
		Object: objc.MakeObject(ptr),
	}
}

func (uc _UndoManagerClass) Alloc() UndoManager {
	rv := objc.CallMethod[UndoManager](uc, objc.GetSelector("alloc"))
	return rv
}

func UndoManager_Alloc() UndoManager {
	return UndoManagerClass.Alloc()
}

func (uc _UndoManagerClass) New() UndoManager {
	rv := objc.CallMethod[UndoManager](uc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewUndoManager() UndoManager {
	return UndoManagerClass.New()
}

func UndoManager_New() UndoManager {
	return UndoManagerClass.New()
}

func (u_ UndoManager) Init() UndoManager {
	rv := objc.CallMethod[UndoManager](u_, objc.GetSelector("init"))
	return rv
}

func UndoManager_Init() UndoManager {
	return UndoManagerClass.Alloc().Init()
}

func (u_ UndoManager) RegisterUndoWithTargetHandler(target objc.IObject, undoHandler func(target objc.Object)) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("registerUndoWithTarget:handler:"), objc.ExtractPtr(target), undoHandler)
}

func (u_ UndoManager) RegisterUndoWithTargetSelectorObject(target objc.IObject, selector objc.Selector, anObject objc.IObject) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("registerUndoWithTarget:selector:object:"), objc.ExtractPtr(target), selector, objc.ExtractPtr(anObject))
}

func (u_ UndoManager) PrepareWithInvocationTarget(target objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](u_, objc.GetSelector("prepareWithInvocationTarget:"), objc.ExtractPtr(target))
	return rv
}

func (u_ UndoManager) Undo() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("undo"))
}

func (u_ UndoManager) UndoNestedGroup() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("undoNestedGroup"))
}

func (u_ UndoManager) Redo() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("redo"))
}

func (u_ UndoManager) BeginUndoGrouping() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("beginUndoGrouping"))
}

func (u_ UndoManager) EndUndoGrouping() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("endUndoGrouping"))
}

func (u_ UndoManager) DisableUndoRegistration() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("disableUndoRegistration"))
}

func (u_ UndoManager) EnableUndoRegistration() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("enableUndoRegistration"))
}

func (u_ UndoManager) RemoveAllActions() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeAllActions"))
}

func (u_ UndoManager) RemoveAllActionsWithTarget(target objc.IObject) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeAllActionsWithTarget:"), objc.ExtractPtr(target))
}

func (u_ UndoManager) SetActionName(actionName string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setActionName:"), actionName)
}

func (u_ UndoManager) UndoMenuTitleForUndoActionName(actionName string) string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("undoMenuTitleForUndoActionName:"), actionName)
	return rv
}

func (u_ UndoManager) RedoMenuTitleForUndoActionName(actionName string) string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("redoMenuTitleForUndoActionName:"), actionName)
	return rv
}

func (u_ UndoManager) SetActionIsDiscardable(discardable bool) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setActionIsDiscardable:"), discardable)
}

func (u_ UndoManager) CanUndo() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("canUndo"))
	return rv
}

func (u_ UndoManager) CanRedo() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("canRedo"))
	return rv
}

func (u_ UndoManager) LevelsOfUndo() uint {
	rv := objc.CallMethod[uint](u_, objc.GetSelector("levelsOfUndo"))
	return rv
}

func (u_ UndoManager) SetLevelsOfUndo(value uint) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setLevelsOfUndo:"), value)
}

func (u_ UndoManager) GroupsByEvent() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("groupsByEvent"))
	return rv
}

func (u_ UndoManager) SetGroupsByEvent(value bool) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setGroupsByEvent:"), value)
}

func (u_ UndoManager) GroupingLevel() int {
	rv := objc.CallMethod[int](u_, objc.GetSelector("groupingLevel"))
	return rv
}

func (u_ UndoManager) IsUndoRegistrationEnabled() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("isUndoRegistrationEnabled"))
	return rv
}

func (u_ UndoManager) IsUndoing() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("isUndoing"))
	return rv
}

func (u_ UndoManager) IsRedoing() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("isRedoing"))
	return rv
}

func (u_ UndoManager) UndoActionName() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("undoActionName"))
	return rv
}

func (u_ UndoManager) RedoActionName() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("redoActionName"))
	return rv
}

func (u_ UndoManager) UndoMenuItemTitle() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("undoMenuItemTitle"))
	return rv
}

func (u_ UndoManager) RedoMenuItemTitle() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("redoMenuItemTitle"))
	return rv
}

func (u_ UndoManager) RunLoopModes() []RunLoopMode {
	rv := objc.CallMethod[[]RunLoopMode](u_, objc.GetSelector("runLoopModes"))
	return rv
}

func (u_ UndoManager) SetRunLoopModes(value []RunLoopMode) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setRunLoopModes:"), value)
}

func (u_ UndoManager) UndoActionIsDiscardable() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("undoActionIsDiscardable"))
	return rv
}

func (u_ UndoManager) RedoActionIsDiscardable() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("redoActionIsDiscardable"))
	return rv
}
