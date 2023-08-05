// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var DocumentClass = _DocumentClass{objc.GetClass("NSDocument")}

type _DocumentClass struct {
	objc.Class
}

type IDocument interface {
	objc.IObject
	ReadFromURLOfTypeError(url foundation.IURL, typeName string, outError *foundation.Error) bool
	ReadFromFileWrapperOfTypeError(fileWrapper foundation.IFileWrapper, typeName string, outError *foundation.Error) bool
	ReadFromDataOfTypeError(data []byte, typeName string, outError *foundation.Error) bool
	CanAsynchronouslyWriteToURLOfTypeForSaveOperation(url foundation.IURL, typeName string, saveOperation SaveOperationType) bool
	UnblockUserInteraction()
	WriteToURLOfTypeError(url foundation.IURL, typeName string, outError *foundation.Error) bool
	WriteSafelyToURLOfTypeForSaveOperationError(url foundation.IURL, typeName string, saveOperation SaveOperationType, outError *foundation.Error) bool
	FileWrapperOfTypeError(typeName string, outError *foundation.Error) foundation.FileWrapper
	DataOfTypeError(typeName string, outError *foundation.Error) []byte
	WriteToURLOfTypeForSaveOperationOriginalContentsURLError(url foundation.IURL, typeName string, saveOperation SaveOperationType, absoluteOriginalContentsURL foundation.IURL, outError *foundation.Error) bool
	SaveToURLOfTypeForSaveOperationDelegateDidSaveSelectorContextInfo(url foundation.IURL, typeName string, saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer)
	SaveToURLOfTypeForSaveOperationCompletionHandler(url foundation.IURL, typeName string, saveOperation SaveOperationType, completionHandler func(errorOrNil foundation.Error))
	FileAttributesToWriteToURLOfTypeForSaveOperationOriginalContentsURLError(url foundation.IURL, typeName string, saveOperation SaveOperationType, absoluteOriginalContentsURL foundation.IURL, outError *foundation.Error) map[string]objc.Object
	WritableTypesForSaveOperation(saveOperation SaveOperationType) []string
	FileNameExtensionForTypeSaveOperation(typeName string, saveOperation SaveOperationType) string
	MakeWindowControllers()
	AddWindowController(windowController IWindowController)
	RemoveWindowController(windowController IWindowController)
	WindowControllerDidLoadNib(windowController IWindowController)
	WindowControllerWillLoadNib(windowController IWindowController)
	ShouldCloseWindowControllerDelegateShouldCloseSelectorContextInfo(windowController IWindowController, delegate objc.IObject, shouldCloseSelector objc.Selector, contextInfo unsafe.Pointer)
	ShowWindows()
	SetWindow(window IWindow)
	SetDisplayName(displayNameOrNil string)
	DefaultDraftName() string
	EncodeRestorableStateWithCoderBackgroundQueue(coder foundation.ICoder, queue foundation.IOperationQueue)
	CheckAutosavingSafetyAndReturnError(outError *foundation.Error) bool
	ScheduleAutosaving()
	AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo(delegate objc.IObject, didAutosaveSelector objc.Selector, contextInfo unsafe.Pointer)
	AutosaveWithImplicitCancellabilityCompletionHandler(autosavingIsImplicitlyCancellable bool, completionHandler func(errorOrNil foundation.Error))
	BrowseDocumentVersions(sender objc.IObject)
	StopBrowsingVersionsWithCompletionHandler(completionHandler func())
	MoveDocumentToUbiquityContainer(sender objc.IObject)
	UpdateChangeCountWithTokenForSaveOperation(changeCountToken objc.IObject, saveOperation SaveOperationType)
	UpdateChangeCount(change DocumentChangeType)
	ChangeCountTokenForSaveOperation(saveOperation SaveOperationType) objc.Object
	EncodeRestorableStateWithCoder(coder foundation.ICoder)
	RestoreStateWithCoder(coder foundation.ICoder)
	InvalidateRestorableState()
	RestoreDocumentWindowWithIdentifierStateCompletionHandler(identifier UserInterfaceItemIdentifier, state foundation.ICoder, completionHandler func(param1 Window, param2 foundation.Error))
	RunModalSavePanelForSaveOperationDelegateDidSaveSelectorContextInfo(saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer)
	PrepareSavePanel(savePanel ISavePanel) bool
	UpdateUserActivityState(activity foundation.IUserActivity)
	ValidateUserInterfaceItem(item IValidatedUserInterfaceItem) bool
	ValidateUserInterfaceItem0(item objc.IObject) bool
	PerformSynchronousFileAccessUsingBlock(block func())
	PerformAsynchronousFileAccessUsingBlock(block func(param1 func()))
	PerformActivityWithSynchronousWaitingUsingBlock(waitSynchronously bool, block func(param1 func()))
	ContinueActivityUsingBlock(block func())
	ContinueAsynchronousWorkOnMainThreadUsingBlock(block func())
	PrintDocument(sender objc.IObject)
	RunPageLayout(sender objc.IObject)
	RevertDocumentToSaved(sender objc.IObject)
	SaveDocument(sender objc.IObject)
	SaveDocumentAs(sender objc.IObject)
	SaveDocumentTo(sender objc.IObject)
	SaveDocumentWithDelegateDidSaveSelectorContextInfo(delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer)
	CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo(delegate objc.IObject, shouldCloseSelector objc.Selector, contextInfo unsafe.Pointer)
	Close()
	RevertToContentsOfURLOfTypeError(url foundation.IURL, typeName string, outError *foundation.Error) bool
	DuplicateAndReturnError(outError *foundation.Error) Document
	DuplicateDocument(sender objc.IObject)
	DuplicateDocumentWithDelegateDidDuplicateSelectorContextInfo(delegate objc.IObject, didDuplicateSelector objc.Selector, contextInfo unsafe.Pointer)
	RenameDocument(sender objc.IObject)
	MoveDocument(sender objc.IObject)
	MoveDocumentWithCompletionHandler(completionHandler func(didMove bool))
	MoveToURLCompletionHandler(url foundation.IURL, completionHandler func(param1 foundation.Error))
	LockDocument(sender objc.IObject)
	UnlockDocument(sender objc.IObject)
	LockDocumentWithCompletionHandler(completionHandler func(didLock bool))
	LockWithCompletionHandler(completionHandler func(param1 foundation.Error))
	UnlockDocumentWithCompletionHandler(completionHandler func(didUnlock bool))
	UnlockWithCompletionHandler(completionHandler func(param1 foundation.Error))
	PreparePageLayout(pageLayout IPageLayout) bool
	RunModalPageLayoutWithPrintInfoDelegateDidRunSelectorContextInfo(printInfo IPrintInfo, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer)
	RunModalPrintOperationDelegateDidRunSelectorContextInfo(printOperation IPrintOperation, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer)
	ShouldChangePrintInfo(newPrintInfo IPrintInfo) bool
	PrintDocumentWithSettingsShowPrintPanelDelegateDidPrintSelectorContextInfo(printSettings map[PrintInfoAttributeKey]objc.IObject, showPrintPanel bool, delegate objc.IObject, didPrintSelector objc.Selector, contextInfo unsafe.Pointer)
	PrintOperationWithSettingsError(printSettings map[PrintInfoAttributeKey]objc.IObject, outError *foundation.Error) PrintOperation
	SaveDocumentToPDF(sender objc.IObject)
	PrepareSharingServicePicker(sharingServicePicker ISharingServicePicker)
	ShareDocumentWithSharingServiceCompletionHandler(sharingService ISharingService, completionHandler func(success bool))
	HandleCloseScriptCommand(command foundation.ICloseCommand) objc.Object
	HandlePrintScriptCommand(command foundation.IScriptCommand) objc.Object
	HandleSaveScriptCommand(command foundation.IScriptCommand) objc.Object
	PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error foundation.IError, window IWindow, delegate objc.IObject, didPresentSelector objc.Selector, contextInfo unsafe.Pointer)
	PresentError(error foundation.IError) bool
	WillPresentError(error foundation.IError) foundation.Error
	WillNotPresentError(error foundation.IError)
	FileURL() foundation.URL
	SetFileURL(value foundation.IURL)
	IsEntireFileLoaded() bool
	FileModificationDate() foundation.Date
	SetFileModificationDate(value foundation.IDate)
	KeepBackupFile() bool
	IsDraft() bool
	SetDraft(value bool)
	FileType() string
	SetFileType(value string)
	IsDocumentEdited() bool
	IsInViewingMode() bool
	WindowControllers() []WindowController
	WindowNibName() NibName
	WindowForSheet() Window
	DisplayName() string
	AutosavedContentsFileURL() foundation.URL
	SetAutosavedContentsFileURL(value foundation.IURL)
	AutosavingFileType() string
	AutosavingIsImplicitlyCancellable() bool
	HasUnautosavedChanges() bool
	BackupFileURL() foundation.URL
	IsBrowsingVersions() bool
	UndoManager() foundation.UndoManager
	SetUndoManager(value foundation.IUndoManager)
	HasUndoManager() bool
	SetHasUndoManager(value bool)
	ShouldRunSavePanelWithAccessoryView() bool
	FileTypeFromLastRunSavePanel() string
	FileNameExtensionWasHiddenInLastRunSavePanel() bool
	UserActivity() foundation.UserActivity
	SetUserActivity(value foundation.IUserActivity)
	IsLocked() bool
	PrintInfo() PrintInfo
	SetPrintInfo(value IPrintInfo)
	PDFPrintOperation() PrintOperation
	AllowsDocumentSharing() bool
	ObjectSpecifier() foundation.ScriptObjectSpecifier
	LastComponentOfFileName() string
	SetLastComponentOfFileName(value string)
}

type Document struct {
	objc.Object
}

func MakeDocument(ptr unsafe.Pointer) Document {
	return Document{
		Object: objc.MakeObject(ptr),
	}
}

func (d_ Document) Init() Document {
	rv := objc.CallMethod[Document](d_, objc.GetSelector("init"))
	return rv
}

func Document_Init() Document {
	return DocumentClass.Alloc().Init()
}

func (d_ Document) InitWithContentsOfURLOfTypeError(url foundation.IURL, typeName string, outError *foundation.Error) Document {
	rv := objc.CallMethod[Document](d_, objc.GetSelector("initWithContentsOfURL:ofType:error:"), objc.ExtractPtr(url), typeName, unsafe.Pointer(outError))
	return rv
}

func Document_InitWithContentsOfURLOfTypeError(url foundation.IURL, typeName string, outError *foundation.Error) Document {
	return DocumentClass.Alloc().InitWithContentsOfURLOfTypeError(url, typeName, outError)
}

func (d_ Document) InitForURLWithContentsOfURLOfTypeError(urlOrNil foundation.IURL, contentsURL foundation.IURL, typeName string, outError *foundation.Error) Document {
	rv := objc.CallMethod[Document](d_, objc.GetSelector("initForURL:withContentsOfURL:ofType:error:"), objc.ExtractPtr(urlOrNil), objc.ExtractPtr(contentsURL), typeName, unsafe.Pointer(outError))
	return rv
}

func Document_InitForURLWithContentsOfURLOfTypeError(urlOrNil foundation.IURL, contentsURL foundation.IURL, typeName string, outError *foundation.Error) Document {
	return DocumentClass.Alloc().InitForURLWithContentsOfURLOfTypeError(urlOrNil, contentsURL, typeName, outError)
}

func (d_ Document) InitWithTypeError(typeName string, outError *foundation.Error) Document {
	rv := objc.CallMethod[Document](d_, objc.GetSelector("initWithType:error:"), typeName, unsafe.Pointer(outError))
	return rv
}

func Document_InitWithTypeError(typeName string, outError *foundation.Error) Document {
	return DocumentClass.Alloc().InitWithTypeError(typeName, outError)
}

func (dc _DocumentClass) Alloc() Document {
	rv := objc.CallMethod[Document](dc, objc.GetSelector("alloc"))
	return rv
}

func Document_Alloc() Document {
	return DocumentClass.Alloc()
}

func (dc _DocumentClass) New() Document {
	rv := objc.CallMethod[Document](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDocument() Document {
	return DocumentClass.New()
}

func Document_New() Document {
	return DocumentClass.New()
}

func (dc _DocumentClass) CanConcurrentlyReadDocumentsOfType(typeName string) bool {
	rv := objc.CallMethod[bool](dc, objc.GetSelector("canConcurrentlyReadDocumentsOfType:"), typeName)
	return rv
}

func Document_CanConcurrentlyReadDocumentsOfType(typeName string) bool {
	return DocumentClass.CanConcurrentlyReadDocumentsOfType(typeName)
}

func (d_ Document) ReadFromURLOfTypeError(url foundation.IURL, typeName string, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("readFromURL:ofType:error:"), objc.ExtractPtr(url), typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) ReadFromFileWrapperOfTypeError(fileWrapper foundation.IFileWrapper, typeName string, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("readFromFileWrapper:ofType:error:"), objc.ExtractPtr(fileWrapper), typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) ReadFromDataOfTypeError(data []byte, typeName string, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("readFromData:ofType:error:"), data, typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) CanAsynchronouslyWriteToURLOfTypeForSaveOperation(url foundation.IURL, typeName string, saveOperation SaveOperationType) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("canAsynchronouslyWriteToURL:ofType:forSaveOperation:"), objc.ExtractPtr(url), typeName, saveOperation)
	return rv
}

func (d_ Document) UnblockUserInteraction() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("unblockUserInteraction"))
}

func (d_ Document) WriteToURLOfTypeError(url foundation.IURL, typeName string, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("writeToURL:ofType:error:"), objc.ExtractPtr(url), typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) WriteSafelyToURLOfTypeForSaveOperationError(url foundation.IURL, typeName string, saveOperation SaveOperationType, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("writeSafelyToURL:ofType:forSaveOperation:error:"), objc.ExtractPtr(url), typeName, saveOperation, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) FileWrapperOfTypeError(typeName string, outError *foundation.Error) foundation.FileWrapper {
	rv := objc.CallMethod[foundation.FileWrapper](d_, objc.GetSelector("fileWrapperOfType:error:"), typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) DataOfTypeError(typeName string, outError *foundation.Error) []byte {
	rv := objc.CallMethod[[]byte](d_, objc.GetSelector("dataOfType:error:"), typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) WriteToURLOfTypeForSaveOperationOriginalContentsURLError(url foundation.IURL, typeName string, saveOperation SaveOperationType, absoluteOriginalContentsURL foundation.IURL, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("writeToURL:ofType:forSaveOperation:originalContentsURL:error:"), objc.ExtractPtr(url), typeName, saveOperation, objc.ExtractPtr(absoluteOriginalContentsURL), unsafe.Pointer(outError))
	return rv
}

func (d_ Document) SaveToURLOfTypeForSaveOperationDelegateDidSaveSelectorContextInfo(url foundation.IURL, typeName string, saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("saveToURL:ofType:forSaveOperation:delegate:didSaveSelector:contextInfo:"), objc.ExtractPtr(url), typeName, saveOperation, objc.ExtractPtr(delegate), didSaveSelector, contextInfo)
}

func (d_ Document) SaveToURLOfTypeForSaveOperationCompletionHandler(url foundation.IURL, typeName string, saveOperation SaveOperationType, completionHandler func(errorOrNil foundation.Error)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("saveToURL:ofType:forSaveOperation:completionHandler:"), objc.ExtractPtr(url), typeName, saveOperation, completionHandler)
}

func (d_ Document) FileAttributesToWriteToURLOfTypeForSaveOperationOriginalContentsURLError(url foundation.IURL, typeName string, saveOperation SaveOperationType, absoluteOriginalContentsURL foundation.IURL, outError *foundation.Error) map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](d_, objc.GetSelector("fileAttributesToWriteToURL:ofType:forSaveOperation:originalContentsURL:error:"), objc.ExtractPtr(url), typeName, saveOperation, objc.ExtractPtr(absoluteOriginalContentsURL), unsafe.Pointer(outError))
	return rv
}

func (dc _DocumentClass) IsNativeType(type_ string) bool {
	rv := objc.CallMethod[bool](dc, objc.GetSelector("isNativeType:"), type_)
	return rv
}

func Document_IsNativeType(type_ string) bool {
	return DocumentClass.IsNativeType(type_)
}

func (d_ Document) WritableTypesForSaveOperation(saveOperation SaveOperationType) []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("writableTypesForSaveOperation:"), saveOperation)
	return rv
}

func (d_ Document) FileNameExtensionForTypeSaveOperation(typeName string, saveOperation SaveOperationType) string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("fileNameExtensionForType:saveOperation:"), typeName, saveOperation)
	return rv
}

func (d_ Document) MakeWindowControllers() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("makeWindowControllers"))
}

func (d_ Document) AddWindowController(windowController IWindowController) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("addWindowController:"), objc.ExtractPtr(windowController))
}

func (d_ Document) RemoveWindowController(windowController IWindowController) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("removeWindowController:"), objc.ExtractPtr(windowController))
}

func (d_ Document) WindowControllerDidLoadNib(windowController IWindowController) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("windowControllerDidLoadNib:"), objc.ExtractPtr(windowController))
}

func (d_ Document) WindowControllerWillLoadNib(windowController IWindowController) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("windowControllerWillLoadNib:"), objc.ExtractPtr(windowController))
}

func (d_ Document) ShouldCloseWindowControllerDelegateShouldCloseSelectorContextInfo(windowController IWindowController, delegate objc.IObject, shouldCloseSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("shouldCloseWindowController:delegate:shouldCloseSelector:contextInfo:"), objc.ExtractPtr(windowController), objc.ExtractPtr(delegate), shouldCloseSelector, contextInfo)
}

func (d_ Document) ShowWindows() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("showWindows"))
}

func (d_ Document) SetWindow(window IWindow) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setWindow:"), objc.ExtractPtr(window))
}

func (d_ Document) SetDisplayName(displayNameOrNil string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDisplayName:"), displayNameOrNil)
}

func (d_ Document) DefaultDraftName() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("defaultDraftName"))
	return rv
}

func (d_ Document) EncodeRestorableStateWithCoderBackgroundQueue(coder foundation.ICoder, queue foundation.IOperationQueue) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("encodeRestorableStateWithCoder:backgroundQueue:"), objc.ExtractPtr(coder), objc.ExtractPtr(queue))
}

func (d_ Document) CheckAutosavingSafetyAndReturnError(outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("checkAutosavingSafetyAndReturnError:"), unsafe.Pointer(outError))
	return rv
}

func (d_ Document) ScheduleAutosaving() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("scheduleAutosaving"))
}

func (d_ Document) AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo(delegate objc.IObject, didAutosaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("autosaveDocumentWithDelegate:didAutosaveSelector:contextInfo:"), objc.ExtractPtr(delegate), didAutosaveSelector, contextInfo)
}

func (d_ Document) AutosaveWithImplicitCancellabilityCompletionHandler(autosavingIsImplicitlyCancellable bool, completionHandler func(errorOrNil foundation.Error)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("autosaveWithImplicitCancellability:completionHandler:"), autosavingIsImplicitlyCancellable, completionHandler)
}

func (d_ Document) BrowseDocumentVersions(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("browseDocumentVersions:"), objc.ExtractPtr(sender))
}

func (d_ Document) StopBrowsingVersionsWithCompletionHandler(completionHandler func()) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("stopBrowsingVersionsWithCompletionHandler:"), completionHandler)
}

func (d_ Document) MoveDocumentToUbiquityContainer(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("moveDocumentToUbiquityContainer:"), objc.ExtractPtr(sender))
}

func (d_ Document) UpdateChangeCountWithTokenForSaveOperation(changeCountToken objc.IObject, saveOperation SaveOperationType) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("updateChangeCountWithToken:forSaveOperation:"), objc.ExtractPtr(changeCountToken), saveOperation)
}

func (d_ Document) UpdateChangeCount(change DocumentChangeType) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("updateChangeCount:"), change)
}

func (d_ Document) ChangeCountTokenForSaveOperation(saveOperation SaveOperationType) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("changeCountTokenForSaveOperation:"), saveOperation)
	return rv
}

func (dc _DocumentClass) AllowedClassesForRestorableStateKeyPath(keyPath string) []objc.Class {
	rv := objc.CallMethod[[]objc.Class](dc, objc.GetSelector("allowedClassesForRestorableStateKeyPath:"), keyPath)
	return rv
}

func Document_AllowedClassesForRestorableStateKeyPath(keyPath string) []objc.Class {
	return DocumentClass.AllowedClassesForRestorableStateKeyPath(keyPath)
}

func (d_ Document) EncodeRestorableStateWithCoder(coder foundation.ICoder) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("encodeRestorableStateWithCoder:"), objc.ExtractPtr(coder))
}

func (d_ Document) RestoreStateWithCoder(coder foundation.ICoder) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("restoreStateWithCoder:"), objc.ExtractPtr(coder))
}

func (d_ Document) InvalidateRestorableState() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("invalidateRestorableState"))
}

func (d_ Document) RestoreDocumentWindowWithIdentifierStateCompletionHandler(identifier UserInterfaceItemIdentifier, state foundation.ICoder, completionHandler func(param1 Window, param2 foundation.Error)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("restoreDocumentWindowWithIdentifier:state:completionHandler:"), identifier, objc.ExtractPtr(state), completionHandler)
}

func (d_ Document) RunModalSavePanelForSaveOperationDelegateDidSaveSelectorContextInfo(saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("runModalSavePanelForSaveOperation:delegate:didSaveSelector:contextInfo:"), saveOperation, objc.ExtractPtr(delegate), didSaveSelector, contextInfo)
}

func (d_ Document) PrepareSavePanel(savePanel ISavePanel) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("prepareSavePanel:"), objc.ExtractPtr(savePanel))
	return rv
}

func (d_ Document) UpdateUserActivityState(activity foundation.IUserActivity) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("updateUserActivityState:"), objc.ExtractPtr(activity))
}

func (d_ Document) ValidateUserInterfaceItem(item IValidatedUserInterfaceItem) bool {
	po := objc.WrapAsProtocol("NSValidatedUserInterfaceItem", item)
	rv := objc.CallMethod[bool](d_, objc.GetSelector("validateUserInterfaceItem:"), po)
	return rv
}

func (d_ Document) ValidateUserInterfaceItem0(item objc.IObject) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("validateUserInterfaceItem:"), objc.ExtractPtr(item))
	return rv
}

func (d_ Document) PerformSynchronousFileAccessUsingBlock(block func()) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("performSynchronousFileAccessUsingBlock:"), block)
}

func (d_ Document) PerformAsynchronousFileAccessUsingBlock(block func(param1 func())) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("performAsynchronousFileAccessUsingBlock:"), block)
}

func (d_ Document) PerformActivityWithSynchronousWaitingUsingBlock(waitSynchronously bool, block func(param1 func())) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("performActivityWithSynchronousWaiting:usingBlock:"), waitSynchronously, block)
}

func (d_ Document) ContinueActivityUsingBlock(block func()) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("continueActivityUsingBlock:"), block)
}

func (d_ Document) ContinueAsynchronousWorkOnMainThreadUsingBlock(block func()) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("continueAsynchronousWorkOnMainThreadUsingBlock:"), block)
}

func (d_ Document) PrintDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("printDocument:"), objc.ExtractPtr(sender))
}

func (d_ Document) RunPageLayout(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("runPageLayout:"), objc.ExtractPtr(sender))
}

func (d_ Document) RevertDocumentToSaved(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("revertDocumentToSaved:"), objc.ExtractPtr(sender))
}

func (d_ Document) SaveDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("saveDocument:"), objc.ExtractPtr(sender))
}

func (d_ Document) SaveDocumentAs(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("saveDocumentAs:"), objc.ExtractPtr(sender))
}

func (d_ Document) SaveDocumentTo(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("saveDocumentTo:"), objc.ExtractPtr(sender))
}

func (d_ Document) SaveDocumentWithDelegateDidSaveSelectorContextInfo(delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("saveDocumentWithDelegate:didSaveSelector:contextInfo:"), objc.ExtractPtr(delegate), didSaveSelector, contextInfo)
}

func (d_ Document) CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo(delegate objc.IObject, shouldCloseSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("canCloseDocumentWithDelegate:shouldCloseSelector:contextInfo:"), objc.ExtractPtr(delegate), shouldCloseSelector, contextInfo)
}

func (d_ Document) Close() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("close"))
}

func (d_ Document) RevertToContentsOfURLOfTypeError(url foundation.IURL, typeName string, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("revertToContentsOfURL:ofType:error:"), objc.ExtractPtr(url), typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) DuplicateAndReturnError(outError *foundation.Error) Document {
	rv := objc.CallMethod[Document](d_, objc.GetSelector("duplicateAndReturnError:"), unsafe.Pointer(outError))
	return rv
}

func (d_ Document) DuplicateDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("duplicateDocument:"), objc.ExtractPtr(sender))
}

func (d_ Document) DuplicateDocumentWithDelegateDidDuplicateSelectorContextInfo(delegate objc.IObject, didDuplicateSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("duplicateDocumentWithDelegate:didDuplicateSelector:contextInfo:"), objc.ExtractPtr(delegate), didDuplicateSelector, contextInfo)
}

func (d_ Document) RenameDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("renameDocument:"), objc.ExtractPtr(sender))
}

func (d_ Document) MoveDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("moveDocument:"), objc.ExtractPtr(sender))
}

func (d_ Document) MoveDocumentWithCompletionHandler(completionHandler func(didMove bool)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("moveDocumentWithCompletionHandler:"), completionHandler)
}

func (d_ Document) MoveToURLCompletionHandler(url foundation.IURL, completionHandler func(param1 foundation.Error)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("moveToURL:completionHandler:"), objc.ExtractPtr(url), completionHandler)
}

func (d_ Document) LockDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("lockDocument:"), objc.ExtractPtr(sender))
}

func (d_ Document) UnlockDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("unlockDocument:"), objc.ExtractPtr(sender))
}

func (d_ Document) LockDocumentWithCompletionHandler(completionHandler func(didLock bool)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("lockDocumentWithCompletionHandler:"), completionHandler)
}

func (d_ Document) LockWithCompletionHandler(completionHandler func(param1 foundation.Error)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("lockWithCompletionHandler:"), completionHandler)
}

func (d_ Document) UnlockDocumentWithCompletionHandler(completionHandler func(didUnlock bool)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("unlockDocumentWithCompletionHandler:"), completionHandler)
}

func (d_ Document) UnlockWithCompletionHandler(completionHandler func(param1 foundation.Error)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("unlockWithCompletionHandler:"), completionHandler)
}

func (d_ Document) PreparePageLayout(pageLayout IPageLayout) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("preparePageLayout:"), objc.ExtractPtr(pageLayout))
	return rv
}

func (d_ Document) RunModalPageLayoutWithPrintInfoDelegateDidRunSelectorContextInfo(printInfo IPrintInfo, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("runModalPageLayoutWithPrintInfo:delegate:didRunSelector:contextInfo:"), objc.ExtractPtr(printInfo), objc.ExtractPtr(delegate), didRunSelector, contextInfo)
}

func (d_ Document) RunModalPrintOperationDelegateDidRunSelectorContextInfo(printOperation IPrintOperation, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("runModalPrintOperation:delegate:didRunSelector:contextInfo:"), objc.ExtractPtr(printOperation), objc.ExtractPtr(delegate), didRunSelector, contextInfo)
}

func (d_ Document) ShouldChangePrintInfo(newPrintInfo IPrintInfo) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("shouldChangePrintInfo:"), objc.ExtractPtr(newPrintInfo))
	return rv
}

func (d_ Document) PrintDocumentWithSettingsShowPrintPanelDelegateDidPrintSelectorContextInfo(printSettings map[PrintInfoAttributeKey]objc.IObject, showPrintPanel bool, delegate objc.IObject, didPrintSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("printDocumentWithSettings:showPrintPanel:delegate:didPrintSelector:contextInfo:"), printSettings, showPrintPanel, objc.ExtractPtr(delegate), didPrintSelector, contextInfo)
}

func (d_ Document) PrintOperationWithSettingsError(printSettings map[PrintInfoAttributeKey]objc.IObject, outError *foundation.Error) PrintOperation {
	rv := objc.CallMethod[PrintOperation](d_, objc.GetSelector("printOperationWithSettings:error:"), printSettings, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) SaveDocumentToPDF(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("saveDocumentToPDF:"), objc.ExtractPtr(sender))
}

func (d_ Document) PrepareSharingServicePicker(sharingServicePicker ISharingServicePicker) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("prepareSharingServicePicker:"), objc.ExtractPtr(sharingServicePicker))
}

func (d_ Document) ShareDocumentWithSharingServiceCompletionHandler(sharingService ISharingService, completionHandler func(success bool)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("shareDocumentWithSharingService:completionHandler:"), objc.ExtractPtr(sharingService), completionHandler)
}

func (d_ Document) HandleCloseScriptCommand(command foundation.ICloseCommand) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("handleCloseScriptCommand:"), objc.ExtractPtr(command))
	return rv
}

func (d_ Document) HandlePrintScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("handlePrintScriptCommand:"), objc.ExtractPtr(command))
	return rv
}

func (d_ Document) HandleSaveScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("handleSaveScriptCommand:"), objc.ExtractPtr(command))
	return rv
}

func (d_ Document) PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error foundation.IError, window IWindow, delegate objc.IObject, didPresentSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), objc.ExtractPtr(error), objc.ExtractPtr(window), objc.ExtractPtr(delegate), didPresentSelector, contextInfo)
}

func (d_ Document) PresentError(error foundation.IError) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("presentError:"), objc.ExtractPtr(error))
	return rv
}

func (d_ Document) WillPresentError(error foundation.IError) foundation.Error {
	rv := objc.CallMethod[foundation.Error](d_, objc.GetSelector("willPresentError:"), objc.ExtractPtr(error))
	return rv
}

func (d_ Document) WillNotPresentError(error foundation.IError) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("willNotPresentError:"), objc.ExtractPtr(error))
}

func (d_ Document) FileURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](d_, objc.GetSelector("fileURL"))
	return rv
}

func (d_ Document) SetFileURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setFileURL:"), objc.ExtractPtr(value))
}

func (d_ Document) IsEntireFileLoaded() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isEntireFileLoaded"))
	return rv
}

func (d_ Document) FileModificationDate() foundation.Date {
	rv := objc.CallMethod[foundation.Date](d_, objc.GetSelector("fileModificationDate"))
	return rv
}

func (d_ Document) SetFileModificationDate(value foundation.IDate) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setFileModificationDate:"), objc.ExtractPtr(value))
}

func (d_ Document) KeepBackupFile() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("keepBackupFile"))
	return rv
}

func (d_ Document) IsDraft() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isDraft"))
	return rv
}

func (d_ Document) SetDraft(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDraft:"), value)
}

func (d_ Document) FileType() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("fileType"))
	return rv
}

func (d_ Document) SetFileType(value string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setFileType:"), value)
}

func (d_ Document) IsDocumentEdited() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isDocumentEdited"))
	return rv
}

func (d_ Document) IsInViewingMode() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isInViewingMode"))
	return rv
}

func (dc _DocumentClass) ReadableTypes() []string {
	rv := objc.CallMethod[[]string](dc, objc.GetSelector("readableTypes"))
	return rv
}

func Document_ReadableTypes() []string {
	return DocumentClass.ReadableTypes()
}

func (dc _DocumentClass) WritableTypes() []string {
	rv := objc.CallMethod[[]string](dc, objc.GetSelector("writableTypes"))
	return rv
}

func Document_WritableTypes() []string {
	return DocumentClass.WritableTypes()
}

func (d_ Document) WindowControllers() []WindowController {
	rv := objc.CallMethod[[]WindowController](d_, objc.GetSelector("windowControllers"))
	return rv
}

func (d_ Document) WindowNibName() NibName {
	rv := objc.CallMethod[NibName](d_, objc.GetSelector("windowNibName"))
	return rv
}

func (d_ Document) WindowForSheet() Window {
	rv := objc.CallMethod[Window](d_, objc.GetSelector("windowForSheet"))
	return rv
}

func (d_ Document) DisplayName() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("displayName"))
	return rv
}

func (dc _DocumentClass) AutosavesInPlace() bool {
	rv := objc.CallMethod[bool](dc, objc.GetSelector("autosavesInPlace"))
	return rv
}

func Document_AutosavesInPlace() bool {
	return DocumentClass.AutosavesInPlace()
}

func (dc _DocumentClass) AutosavesDrafts() bool {
	rv := objc.CallMethod[bool](dc, objc.GetSelector("autosavesDrafts"))
	return rv
}

func Document_AutosavesDrafts() bool {
	return DocumentClass.AutosavesDrafts()
}

func (dc _DocumentClass) PreservesVersions() bool {
	rv := objc.CallMethod[bool](dc, objc.GetSelector("preservesVersions"))
	return rv
}

func Document_PreservesVersions() bool {
	return DocumentClass.PreservesVersions()
}

func (d_ Document) AutosavedContentsFileURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](d_, objc.GetSelector("autosavedContentsFileURL"))
	return rv
}

func (d_ Document) SetAutosavedContentsFileURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setAutosavedContentsFileURL:"), objc.ExtractPtr(value))
}

func (d_ Document) AutosavingFileType() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("autosavingFileType"))
	return rv
}

func (d_ Document) AutosavingIsImplicitlyCancellable() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("autosavingIsImplicitlyCancellable"))
	return rv
}

func (d_ Document) HasUnautosavedChanges() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("hasUnautosavedChanges"))
	return rv
}

func (d_ Document) BackupFileURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](d_, objc.GetSelector("backupFileURL"))
	return rv
}

func (d_ Document) IsBrowsingVersions() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isBrowsingVersions"))
	return rv
}

func (dc _DocumentClass) UsesUbiquitousStorage() bool {
	rv := objc.CallMethod[bool](dc, objc.GetSelector("usesUbiquitousStorage"))
	return rv
}

func Document_UsesUbiquitousStorage() bool {
	return DocumentClass.UsesUbiquitousStorage()
}

func (d_ Document) UndoManager() foundation.UndoManager {
	rv := objc.CallMethod[foundation.UndoManager](d_, objc.GetSelector("undoManager"))
	return rv
}

func (d_ Document) SetUndoManager(value foundation.IUndoManager) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setUndoManager:"), objc.ExtractPtr(value))
}

func (d_ Document) HasUndoManager() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("hasUndoManager"))
	return rv
}

func (d_ Document) SetHasUndoManager(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setHasUndoManager:"), value)
}

func (dc _DocumentClass) RestorableStateKeyPaths() []string {
	rv := objc.CallMethod[[]string](dc, objc.GetSelector("restorableStateKeyPaths"))
	return rv
}

func Document_RestorableStateKeyPaths() []string {
	return DocumentClass.RestorableStateKeyPaths()
}

func (d_ Document) ShouldRunSavePanelWithAccessoryView() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("shouldRunSavePanelWithAccessoryView"))
	return rv
}

func (d_ Document) FileTypeFromLastRunSavePanel() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("fileTypeFromLastRunSavePanel"))
	return rv
}

func (d_ Document) FileNameExtensionWasHiddenInLastRunSavePanel() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("fileNameExtensionWasHiddenInLastRunSavePanel"))
	return rv
}

func (d_ Document) UserActivity() foundation.UserActivity {
	rv := objc.CallMethod[foundation.UserActivity](d_, objc.GetSelector("userActivity"))
	return rv
}

func (d_ Document) SetUserActivity(value foundation.IUserActivity) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setUserActivity:"), objc.ExtractPtr(value))
}

func (d_ Document) IsLocked() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isLocked"))
	return rv
}

func (d_ Document) PrintInfo() PrintInfo {
	rv := objc.CallMethod[PrintInfo](d_, objc.GetSelector("printInfo"))
	return rv
}

func (d_ Document) SetPrintInfo(value IPrintInfo) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setPrintInfo:"), objc.ExtractPtr(value))
}

func (d_ Document) PDFPrintOperation() PrintOperation {
	rv := objc.CallMethod[PrintOperation](d_, objc.GetSelector("PDFPrintOperation"))
	return rv
}

func (d_ Document) AllowsDocumentSharing() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("allowsDocumentSharing"))
	return rv
}

func (d_ Document) ObjectSpecifier() foundation.ScriptObjectSpecifier {
	rv := objc.CallMethod[foundation.ScriptObjectSpecifier](d_, objc.GetSelector("objectSpecifier"))
	return rv
}

func (d_ Document) LastComponentOfFileName() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("lastComponentOfFileName"))
	return rv
}

func (d_ Document) SetLastComponentOfFileName(value string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setLastComponentOfFileName:"), value)
}
