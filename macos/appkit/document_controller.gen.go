// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DocumentController] class.
var DocumentControllerClass = _DocumentControllerClass{objc.GetClass("NSDocumentController")}

type _DocumentControllerClass struct {
	objc.Class
}

// An interface definition for the [DocumentController] class.
type IDocumentController interface {
	objc.IObject
	ClearRecentDocuments(sender objc.IObject) objc.Object
	OpenDocument(sender objc.IObject) objc.Object
	DocumentForURL(url foundation.IURL) Document
	WillPresentError(error foundation.IError) foundation.Error
	ReviewUnsavedDocumentsWithAlertTitleCancellableDelegateDidReviewAllSelectorContextInfo(title string, cancellable bool, delegate objc.IObject, didReviewAllSelector objc.Selector, contextInfo unsafe.Pointer)
	RemoveDocument(document IDocument)
	BeginOpenPanelWithCompletionHandler(completionHandler func(arg0 []foundation.URL))
	DocumentForWindow(window IWindow) Document
	AddDocument(document IDocument)
	DisplayNameForType(typeName string) string
	SaveAllDocuments(sender objc.IObject) objc.Object
	ValidateUserInterfaceItem(item PValidatedUserInterfaceItem) bool
	ValidateUserInterfaceItemObject(itemObject objc.IObject) bool
	TypeForContentsOfURLError(url foundation.IURL, outError foundation.IError) string
	RunModalOpenPanelForTypes(openPanel IOpenPanel, types []string) int
	NoteNewRecentDocument(document IDocument)
	URLsFromRunningOpenPanel() []foundation.URL
	NoteNewRecentDocumentURL(url foundation.IURL)
	DocumentClassForType(typeName string) objc.Class
	OpenUntitledDocumentAndDisplayError(displayDocument bool, outError foundation.IError) Document
	DuplicateDocumentWithContentsOfURLCopyingDisplayNameError(url foundation.IURL, duplicateByCopying bool, displayNameOrNil string, outError foundation.IError) Document
	NewDocument(sender objc.IObject) objc.Object
	CloseAllDocumentsWithDelegateDidCloseAllSelectorContextInfo(delegate objc.IObject, didCloseAllSelector objc.Selector, contextInfo unsafe.Pointer)
	MakeDocumentForURLWithContentsOfURLOfTypeError(urlOrNil foundation.IURL, contentsURL foundation.IURL, typeName string, outError foundation.IError) Document
	PresentError(error foundation.IError) bool
	StandardShareMenuItem() MenuItem
	BeginOpenPanelForTypesCompletionHandler(openPanel IOpenPanel, inTypes []string, completionHandler func(result int))
	CurrentDocument() Document
	MaximumRecentDocumentCount() uint
	DefaultType() string
	DocumentClassNames() []string
	RecentDocumentURLs() []foundation.URL
	AutosavingDelay() foundation.TimeInterval
	SetAutosavingDelay(value foundation.TimeInterval)
	AllowsAutomaticShareMenu() bool
	Documents() []Document
	HasEditedDocuments() bool
	CurrentDirectory() string
}

// An object that manages an app’s documents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller?language=objc
type DocumentController struct {
	objc.Object
}

func DocumentControllerFrom(ptr unsafe.Pointer) DocumentController {
	return DocumentController{
		Object: objc.ObjectFrom(ptr),
	}
}

func (d_ DocumentController) Init() DocumentController {
	rv := objc.Call[DocumentController](d_, objc.Sel("init"))
	return rv
}

func (dc _DocumentControllerClass) Alloc() DocumentController {
	rv := objc.Call[DocumentController](dc, objc.Sel("alloc"))
	return rv
}

func DocumentController_Alloc() DocumentController {
	return DocumentControllerClass.Alloc()
}

func (dc _DocumentControllerClass) New() DocumentController {
	rv := objc.Call[DocumentController](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDocumentController() DocumentController {
	return DocumentControllerClass.New()
}

// Empties the recent documents list for the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514933-clearrecentdocuments?language=objc
func (d_ DocumentController) ClearRecentDocuments(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("clearRecentDocuments:"), sender)
	return rv
}

// An action method called by the Open menu command, it runs the modal Open panel and, based on the selected filenames, creates one or more NSDocument objects from the contents of the files. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1515005-opendocument?language=objc
func (d_ DocumentController) OpenDocument(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("openDocument:"), sender)
	return rv
}

// Returns, for a given URL, the open document whose file or file package is located by the URL, or nil if there is no such open document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514939-documentforurl?language=objc
func (d_ DocumentController) DocumentForURL(url foundation.IURL) Document {
	rv := objc.Call[Document](d_, objc.Sel("documentForURL:"), objc.Ptr(url))
	return rv
}

// Indicates an error condition and provides the opportunity to return the same or a different error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514994-willpresenterror?language=objc
func (d_ DocumentController) WillPresentError(error foundation.IError) foundation.Error {
	rv := objc.Call[foundation.Error](d_, objc.Sel("willPresentError:"), objc.Ptr(error))
	return rv
}

// Displays an alert asking if the user wants to review unsaved documents, quit regardless of unsaved documents, or cancel the save operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514940-reviewunsaveddocumentswithalertt?language=objc
func (d_ DocumentController) ReviewUnsavedDocumentsWithAlertTitleCancellableDelegateDidReviewAllSelectorContextInfo(title string, cancellable bool, delegate objc.IObject, didReviewAllSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](d_, objc.Sel("reviewUnsavedDocumentsWithAlertTitle:cancellable:delegate:didReviewAllSelector:contextInfo:"), title, cancellable, delegate, didReviewAllSelector, contextInfo)
}

// Removes the given document from the list of open documents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514984-removedocument?language=objc
func (d_ DocumentController) RemoveDocument(document IDocument) {
	objc.Call[objc.Void](d_, objc.Sel("removeDocument:"), objc.Ptr(document))
}

// Presents an Open dialog and delivers the results to a completion handler as an array of URLs for the chosen files (or nil). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1515001-beginopenpanelwithcompletionhand?language=objc
func (d_ DocumentController) BeginOpenPanelWithCompletionHandler(completionHandler func(arg0 []foundation.URL)) {
	objc.Call[objc.Void](d_, objc.Sel("beginOpenPanelWithCompletionHandler:"), completionHandler)
}

// Returns the document object whose window controller owns a specified window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514970-documentforwindow?language=objc
func (d_ DocumentController) DocumentForWindow(window IWindow) Document {
	rv := objc.Call[Document](d_, objc.Sel("documentForWindow:"), objc.Ptr(window))
	return rv
}

// Adds the given document to the list of open documents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1515013-adddocument?language=objc
func (d_ DocumentController) AddDocument(document IDocument) {
	objc.Call[objc.Void](d_, objc.Sel("addDocument:"), objc.Ptr(document))
}

// Returns the descriptive name for the specified document type, which is used in the File Format pop-up menu of the Save As dialog. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514937-displaynamefortype?language=objc
func (d_ DocumentController) DisplayNameForType(typeName string) string {
	rv := objc.Call[string](d_, objc.Sel("displayNameForType:"), typeName)
	return rv
}

// As the action method called by the Save All command, saves all open documents of the application that need to be saved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514959-savealldocuments?language=objc
func (d_ DocumentController) SaveAllDocuments(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("saveAllDocuments:"), sender)
	return rv
}

// Returns a Boolean value that indicates whether a given user interface item should be enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514943-validateuserinterfaceitem?language=objc
func (d_ DocumentController) ValidateUserInterfaceItem(item PValidatedUserInterfaceItem) bool {
	po0 := objc.WrapAsProtocol("NSValidatedUserInterfaceItem", item)
	rv := objc.Call[bool](d_, objc.Sel("validateUserInterfaceItem:"), po0)
	return rv
}

// Returns a Boolean value that indicates whether a given user interface item should be enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514943-validateuserinterfaceitem?language=objc
func (d_ DocumentController) ValidateUserInterfaceItemObject(itemObject objc.IObject) bool {
	rv := objc.Call[bool](d_, objc.Sel("validateUserInterfaceItem:"), objc.Ptr(itemObject))
	return rv
}

// Returns, for a specified URL, the document type identifier to use when opening the document at that location, if successful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514946-typeforcontentsofurl?language=objc
func (d_ DocumentController) TypeForContentsOfURLError(url foundation.IURL, outError foundation.IError) string {
	rv := objc.Call[string](d_, objc.Sel("typeForContentsOfURL:error:"), objc.Ptr(url), objc.Ptr(outError))
	return rv
}

// Presents a modal Open dialog and limits selection to specific file types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514960-runmodalopenpanel?language=objc
func (d_ DocumentController) RunModalOpenPanelForTypes(openPanel IOpenPanel, types []string) int {
	rv := objc.Call[int](d_, objc.Sel("runModalOpenPanel:forTypes:"), objc.Ptr(openPanel), types)
	return rv
}

// Adds or replaces an Open Recent menu item corresponding to the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1515009-notenewrecentdocument?language=objc
func (d_ DocumentController) NoteNewRecentDocument(document IDocument) {
	objc.Call[objc.Void](d_, objc.Sel("noteNewRecentDocument:"), objc.Ptr(document))
}

// An array of URLs corresponding to the files selected in a running open panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514972-urlsfromrunningopenpanel?language=objc
func (d_ DocumentController) URLsFromRunningOpenPanel() []foundation.URL {
	rv := objc.Call[[]foundation.URL](d_, objc.Sel("URLsFromRunningOpenPanel"))
	return rv
}

// Adds or replaces an Open Recent menu item corresponding to the data located by the URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514967-notenewrecentdocumenturl?language=objc
func (d_ DocumentController) NoteNewRecentDocumentURL(url foundation.IURL) {
	objc.Call[objc.Void](d_, objc.Sel("noteNewRecentDocumentURL:"), objc.Ptr(url))
}

// Returns the NSDocument subclass associated with a given document type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514947-documentclassfortype?language=objc
func (d_ DocumentController) DocumentClassForType(typeName string) objc.Class {
	rv := objc.Call[objc.Class](d_, objc.Sel("documentClassForType:"), typeName)
	return rv
}

// Creates a new untitled document, presents its user interface if displayDocument is YES, and returns the document if successful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1515014-openuntitleddocumentanddisplay?language=objc
func (d_ DocumentController) OpenUntitledDocumentAndDisplayError(displayDocument bool, outError foundation.IError) Document {
	rv := objc.Call[Document](d_, objc.Sel("openUntitledDocumentAndDisplay:error:"), displayDocument, objc.Ptr(outError))
	return rv
}

// Creates a new document by reading the contents for the document from another URL, presents its user interface, and returns the document if successful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514982-duplicatedocumentwithcontentsofu?language=objc
func (d_ DocumentController) DuplicateDocumentWithContentsOfURLCopyingDisplayNameError(url foundation.IURL, duplicateByCopying bool, displayNameOrNil string, outError foundation.IError) Document {
	rv := objc.Call[Document](d_, objc.Sel("duplicateDocumentWithContentsOfURL:copying:displayName:error:"), objc.Ptr(url), duplicateByCopying, displayNameOrNil, objc.Ptr(outError))
	return rv
}

// An action method called by the New menu command, this method creates a new NSDocument object and adds it to the list of such objects managed by the document controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514997-newdocument?language=objc
func (d_ DocumentController) NewDocument(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("newDocument:"), sender)
	rv.Autorelease()
	return rv
}

// Iterates through all the open documents and tries to close them one by one using the specified delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514995-closealldocumentswithdelegate?language=objc
func (d_ DocumentController) CloseAllDocumentsWithDelegateDidCloseAllSelectorContextInfo(delegate objc.IObject, didCloseAllSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](d_, objc.Sel("closeAllDocumentsWithDelegate:didCloseAllSelector:contextInfo:"), delegate, didCloseAllSelector, contextInfo)
}

// Instantiates a document located by a URL, of a specified type, but by reading the contents for the document from another URL, and returns it if successful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514930-makedocumentforurl?language=objc
func (d_ DocumentController) MakeDocumentForURLWithContentsOfURLOfTypeError(urlOrNil foundation.IURL, contentsURL foundation.IURL, typeName string, outError foundation.IError) Document {
	rv := objc.Call[Document](d_, objc.Sel("makeDocumentForURL:withContentsOfURL:ofType:error:"), objc.Ptr(urlOrNil), objc.Ptr(contentsURL), typeName, objc.Ptr(outError))
	return rv
}

// Presents an error alert to the user as a modal panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514977-presenterror?language=objc
func (d_ DocumentController) PresentError(error foundation.IError) bool {
	rv := objc.Call[bool](d_, objc.Sel("presentError:"), objc.Ptr(error))
	return rv
}

// Returns a menu item that your app uses for sharing the current document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/2902324-standardsharemenuitem?language=objc
func (d_ DocumentController) StandardShareMenuItem() MenuItem {
	rv := objc.Call[MenuItem](d_, objc.Sel("standardShareMenuItem"))
	return rv
}

// Presents a nonmodal Open dialog that displays files you can open from a list of UTIs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514969-beginopenpanel?language=objc
func (d_ DocumentController) BeginOpenPanelForTypesCompletionHandler(openPanel IOpenPanel, inTypes []string, completionHandler func(result int)) {
	objc.Call[objc.Void](d_, objc.Sel("beginOpenPanel:forTypes:completionHandler:"), objc.Ptr(openPanel), inTypes, completionHandler)
}

// The document object associated with the main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514957-currentdocument?language=objc
func (d_ DocumentController) CurrentDocument() Document {
	rv := objc.Call[Document](d_, objc.Sel("currentDocument"))
	return rv
}

// The maximum number of items that may be presented in the standard Open Recent menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514991-maximumrecentdocumentcount?language=objc
func (d_ DocumentController) MaximumRecentDocumentCount() uint {
	rv := objc.Call[uint](d_, objc.Sel("maximumRecentDocumentCount"))
	return rv
}

// Returns the name of the document type that should be used when creating new documents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514986-defaulttype?language=objc
func (d_ DocumentController) DefaultType() string {
	rv := objc.Call[string](d_, objc.Sel("defaultType"))
	return rv
}

// An array of strings representing the custom document classes supported by this app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514988-documentclassnames?language=objc
func (d_ DocumentController) DocumentClassNames() []string {
	rv := objc.Call[[]string](d_, objc.Sel("documentClassNames"))
	return rv
}

// The list of recent-document URLs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514976-recentdocumenturls?language=objc
func (d_ DocumentController) RecentDocumentURLs() []foundation.URL {
	rv := objc.Call[[]foundation.URL](d_, objc.Sel("recentDocumentURLs"))
	return rv
}

// The time interval (in seconds) for periodic autosaving. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514953-autosavingdelay?language=objc
func (d_ DocumentController) AutosavingDelay() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](d_, objc.Sel("autosavingDelay"))
	return rv
}

// The time interval (in seconds) for periodic autosaving. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514953-autosavingdelay?language=objc
func (d_ DocumentController) SetAutosavingDelay(value foundation.TimeInterval) {
	objc.Call[objc.Void](d_, objc.Sel("setAutosavingDelay:"), value)
}

// A Boolean value that the system uses to insert a Share menu in the File menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/2902308-allowsautomaticsharemenu?language=objc
func (d_ DocumentController) AllowsAutomaticShareMenu() bool {
	rv := objc.Call[bool](d_, objc.Sel("allowsAutomaticShareMenu"))
	return rv
}

// The document objects managed by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1515003-documents?language=objc
func (d_ DocumentController) Documents() []Document {
	rv := objc.Call[[]Document](d_, objc.Sel("documents"))
	return rv
}

// A Boolean value indicating whether the receiver has any documents with unsaved changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514964-hasediteddocuments?language=objc
func (d_ DocumentController) HasEditedDocuments() bool {
	rv := objc.Call[bool](d_, objc.Sel("hasEditedDocuments"))
	return rv
}

// Returns the shared NSDocumentController instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514981-shareddocumentcontroller?language=objc
func (dc _DocumentControllerClass) SharedDocumentController() DocumentController {
	rv := objc.Call[DocumentController](dc, objc.Sel("sharedDocumentController"))
	return rv
}

// Returns the shared NSDocumentController instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514981-shareddocumentcontroller?language=objc
func DocumentController_SharedDocumentController() DocumentController {
	return DocumentControllerClass.SharedDocumentController()
}

// The directory path to be used as the starting point in the Open panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/1514942-currentdirectory?language=objc
func (d_ DocumentController) CurrentDirectory() string {
	rv := objc.Call[string](d_, objc.Sel("currentDirectory"))
	return rv
}
