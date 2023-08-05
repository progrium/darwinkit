// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/uniformtypeidentifiers"
	"github.com/progrium/macdriver/objc"
)

var WorkspaceClass = _WorkspaceClass{objc.GetClass("NSWorkspace")}

type _WorkspaceClass struct {
	objc.Class
}

type IWorkspace interface {
	objc.IObject
	OpenURL(url foundation.IURL) bool
	HideOtherApplications()
	ActivateFileViewerSelectingURLs(fileURLs []foundation.IURL)
	SelectFileInFileViewerRootedAtPath(fullPath string, rootFullPath string) bool
	URLForApplicationWithBundleIdentifier(bundleIdentifier string) foundation.URL
	URLForApplicationToOpenURL(url foundation.IURL) foundation.URL
	GetFileSystemInfoForPathIsRemovableIsWritableIsUnmountableDescriptionType(fullPath string, removableFlag *bool, writableFlag *bool, unmountableFlag *bool, description *foundation.String, fileSystemType *foundation.String) bool
	IsFilePackageAtPath(fullPath string) bool
	IconForFile(fullPath string) Image
	IconForFiles(fullPaths []string) Image
	IconForContentType(contentType uniformtypeidentifiers.IType) Image
	SetIconForFileOptions(image IImage, fullPath string, options WorkspaceIconCreationOptions) bool
	UnmountAndEjectDeviceAtPath(path string) bool
	UnmountAndEjectDeviceAtURLError(url foundation.IURL, error *foundation.Error) bool
	DesktopImageURLForScreen(screen IScreen) foundation.URL
	SetDesktopImageURLForScreenOptionsError(url foundation.IURL, screen IScreen, options map[WorkspaceDesktopImageOptionKey]objc.IObject, error *foundation.Error) bool
	DesktopImageOptionsForScreen(screen IScreen) map[WorkspaceDesktopImageOptionKey]objc.Object
	ShowSearchResultsForQueryString(queryString string) bool
	NoteFileSystemChanged_(path string)
	ExtendPowerOffBy(requested int) int
	SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler(applicationURL foundation.IURL, contentType uniformtypeidentifiers.IType, completionHandler func(error foundation.Error))
	SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler func(error foundation.Error))
	SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler func(error foundation.Error))
	SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler(applicationURL foundation.IURL, urlScheme string, completionHandler func(error foundation.Error))
	URLForApplicationToOpenContentType(contentType uniformtypeidentifiers.IType) foundation.URL
	URLsForApplicationsToOpenContentType(contentType uniformtypeidentifiers.IType) []foundation.URL
	URLsForApplicationsToOpenURL(url foundation.IURL) []foundation.URL
	URLsForApplicationsWithBundleIdentifier(bundleIdentifier string) []foundation.URL
	NotificationCenter() foundation.NotificationCenter
	FrontmostApplication() RunningApplication
	RunningApplications() []RunningApplication
	MenuBarOwningApplication() RunningApplication
	FileLabels() []string
	FileLabelColors() []Color
	AccessibilityDisplayShouldDifferentiateWithoutColor() bool
	AccessibilityDisplayShouldIncreaseContrast() bool
	AccessibilityDisplayShouldReduceTransparency() bool
	AccessibilityDisplayShouldInvertColors() bool
	AccessibilityDisplayShouldReduceMotion() bool
	IsSwitchControlEnabled() bool
	IsVoiceOverEnabled() bool
}

type Workspace struct {
	objc.Object
}

func MakeWorkspace(ptr unsafe.Pointer) Workspace {
	return Workspace{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WorkspaceClass) Alloc() Workspace {
	rv := objc.CallMethod[Workspace](wc, objc.GetSelector("alloc"))
	return rv
}

func Workspace_Alloc() Workspace {
	return WorkspaceClass.Alloc()
}

func (wc _WorkspaceClass) New() Workspace {
	rv := objc.CallMethod[Workspace](wc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewWorkspace() Workspace {
	return WorkspaceClass.New()
}

func Workspace_New() Workspace {
	return WorkspaceClass.New()
}

func (w_ Workspace) Init() Workspace {
	rv := objc.CallMethod[Workspace](w_, objc.GetSelector("init"))
	return rv
}

func Workspace_Init() Workspace {
	return WorkspaceClass.Alloc().Init()
}

func (w_ Workspace) OpenURL(url foundation.IURL) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("openURL:"), objc.ExtractPtr(url))
	return rv
}

func (w_ Workspace) HideOtherApplications() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("hideOtherApplications"))
}

func (w_ Workspace) ActivateFileViewerSelectingURLs(fileURLs []foundation.IURL) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("activateFileViewerSelectingURLs:"), fileURLs)
}

func (w_ Workspace) SelectFileInFileViewerRootedAtPath(fullPath string, rootFullPath string) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("selectFile:inFileViewerRootedAtPath:"), fullPath, rootFullPath)
	return rv
}

func (w_ Workspace) URLForApplicationWithBundleIdentifier(bundleIdentifier string) foundation.URL {
	rv := objc.CallMethod[foundation.URL](w_, objc.GetSelector("URLForApplicationWithBundleIdentifier:"), bundleIdentifier)
	return rv
}

func (w_ Workspace) URLForApplicationToOpenURL(url foundation.IURL) foundation.URL {
	rv := objc.CallMethod[foundation.URL](w_, objc.GetSelector("URLForApplicationToOpenURL:"), objc.ExtractPtr(url))
	return rv
}

func (w_ Workspace) GetFileSystemInfoForPathIsRemovableIsWritableIsUnmountableDescriptionType(fullPath string, removableFlag *bool, writableFlag *bool, unmountableFlag *bool, description *foundation.String, fileSystemType *foundation.String) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("getFileSystemInfoForPath:isRemovable:isWritable:isUnmountable:description:type:"), fullPath, removableFlag, writableFlag, unmountableFlag, description, fileSystemType)
	return rv
}

func (w_ Workspace) IsFilePackageAtPath(fullPath string) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isFilePackageAtPath:"), fullPath)
	return rv
}

func (w_ Workspace) IconForFile(fullPath string) Image {
	rv := objc.CallMethod[Image](w_, objc.GetSelector("iconForFile:"), fullPath)
	return rv
}

func (w_ Workspace) IconForFiles(fullPaths []string) Image {
	rv := objc.CallMethod[Image](w_, objc.GetSelector("iconForFiles:"), fullPaths)
	return rv
}

func (w_ Workspace) IconForContentType(contentType uniformtypeidentifiers.IType) Image {
	rv := objc.CallMethod[Image](w_, objc.GetSelector("iconForContentType:"), objc.ExtractPtr(contentType))
	return rv
}

func (w_ Workspace) SetIconForFileOptions(image IImage, fullPath string, options WorkspaceIconCreationOptions) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("setIcon:forFile:options:"), objc.ExtractPtr(image), fullPath, options)
	return rv
}

func (w_ Workspace) UnmountAndEjectDeviceAtPath(path string) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("unmountAndEjectDeviceAtPath:"), path)
	return rv
}

func (w_ Workspace) UnmountAndEjectDeviceAtURLError(url foundation.IURL, error *foundation.Error) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("unmountAndEjectDeviceAtURL:error:"), objc.ExtractPtr(url), unsafe.Pointer(error))
	return rv
}

func (w_ Workspace) DesktopImageURLForScreen(screen IScreen) foundation.URL {
	rv := objc.CallMethod[foundation.URL](w_, objc.GetSelector("desktopImageURLForScreen:"), objc.ExtractPtr(screen))
	return rv
}

func (w_ Workspace) SetDesktopImageURLForScreenOptionsError(url foundation.IURL, screen IScreen, options map[WorkspaceDesktopImageOptionKey]objc.IObject, error *foundation.Error) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("setDesktopImageURL:forScreen:options:error:"), objc.ExtractPtr(url), objc.ExtractPtr(screen), options, unsafe.Pointer(error))
	return rv
}

func (w_ Workspace) DesktopImageOptionsForScreen(screen IScreen) map[WorkspaceDesktopImageOptionKey]objc.Object {
	rv := objc.CallMethod[map[WorkspaceDesktopImageOptionKey]objc.Object](w_, objc.GetSelector("desktopImageOptionsForScreen:"), objc.ExtractPtr(screen))
	return rv
}

func (w_ Workspace) ShowSearchResultsForQueryString(queryString string) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("showSearchResultsForQueryString:"), queryString)
	return rv
}

func (w_ Workspace) NoteFileSystemChanged_(path string) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("noteFileSystemChanged:"), path)
}

func (w_ Workspace) ExtendPowerOffBy(requested int) int {
	rv := objc.CallMethod[int](w_, objc.GetSelector("extendPowerOffBy:"), requested)
	return rv
}

func (w_ Workspace) SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler(applicationURL foundation.IURL, contentType uniformtypeidentifiers.IType, completionHandler func(error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDefaultApplicationAtURL:toOpenContentType:completionHandler:"), objc.ExtractPtr(applicationURL), objc.ExtractPtr(contentType), completionHandler)
}

func (w_ Workspace) SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler func(error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDefaultApplicationAtURL:toOpenContentTypeOfFileAtURL:completionHandler:"), objc.ExtractPtr(applicationURL), objc.ExtractPtr(url), completionHandler)
}

func (w_ Workspace) SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler func(error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDefaultApplicationAtURL:toOpenFileAtURL:completionHandler:"), objc.ExtractPtr(applicationURL), objc.ExtractPtr(url), completionHandler)
}

func (w_ Workspace) SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler(applicationURL foundation.IURL, urlScheme string, completionHandler func(error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDefaultApplicationAtURL:toOpenURLsWithScheme:completionHandler:"), objc.ExtractPtr(applicationURL), urlScheme, completionHandler)
}

func (w_ Workspace) URLForApplicationToOpenContentType(contentType uniformtypeidentifiers.IType) foundation.URL {
	rv := objc.CallMethod[foundation.URL](w_, objc.GetSelector("URLForApplicationToOpenContentType:"), objc.ExtractPtr(contentType))
	return rv
}

func (w_ Workspace) URLsForApplicationsToOpenContentType(contentType uniformtypeidentifiers.IType) []foundation.URL {
	rv := objc.CallMethod[[]foundation.URL](w_, objc.GetSelector("URLsForApplicationsToOpenContentType:"), objc.ExtractPtr(contentType))
	return rv
}

func (w_ Workspace) URLsForApplicationsToOpenURL(url foundation.IURL) []foundation.URL {
	rv := objc.CallMethod[[]foundation.URL](w_, objc.GetSelector("URLsForApplicationsToOpenURL:"), objc.ExtractPtr(url))
	return rv
}

func (w_ Workspace) URLsForApplicationsWithBundleIdentifier(bundleIdentifier string) []foundation.URL {
	rv := objc.CallMethod[[]foundation.URL](w_, objc.GetSelector("URLsForApplicationsWithBundleIdentifier:"), bundleIdentifier)
	return rv
}

func (wc _WorkspaceClass) SharedWorkspace() Workspace {
	rv := objc.CallMethod[Workspace](wc, objc.GetSelector("sharedWorkspace"))
	return rv
}

func Workspace_SharedWorkspace() Workspace {
	return WorkspaceClass.SharedWorkspace()
}

func (w_ Workspace) NotificationCenter() foundation.NotificationCenter {
	rv := objc.CallMethod[foundation.NotificationCenter](w_, objc.GetSelector("notificationCenter"))
	return rv
}

func (w_ Workspace) FrontmostApplication() RunningApplication {
	rv := objc.CallMethod[RunningApplication](w_, objc.GetSelector("frontmostApplication"))
	return rv
}

func (w_ Workspace) RunningApplications() []RunningApplication {
	rv := objc.CallMethod[[]RunningApplication](w_, objc.GetSelector("runningApplications"))
	return rv
}

func (w_ Workspace) MenuBarOwningApplication() RunningApplication {
	rv := objc.CallMethod[RunningApplication](w_, objc.GetSelector("menuBarOwningApplication"))
	return rv
}

func (w_ Workspace) FileLabels() []string {
	rv := objc.CallMethod[[]string](w_, objc.GetSelector("fileLabels"))
	return rv
}

func (w_ Workspace) FileLabelColors() []Color {
	rv := objc.CallMethod[[]Color](w_, objc.GetSelector("fileLabelColors"))
	return rv
}

func (w_ Workspace) AccessibilityDisplayShouldDifferentiateWithoutColor() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("accessibilityDisplayShouldDifferentiateWithoutColor"))
	return rv
}

func (w_ Workspace) AccessibilityDisplayShouldIncreaseContrast() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("accessibilityDisplayShouldIncreaseContrast"))
	return rv
}

func (w_ Workspace) AccessibilityDisplayShouldReduceTransparency() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("accessibilityDisplayShouldReduceTransparency"))
	return rv
}

func (w_ Workspace) AccessibilityDisplayShouldInvertColors() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("accessibilityDisplayShouldInvertColors"))
	return rv
}

func (w_ Workspace) AccessibilityDisplayShouldReduceMotion() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("accessibilityDisplayShouldReduceMotion"))
	return rv
}

func (w_ Workspace) IsSwitchControlEnabled() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isSwitchControlEnabled"))
	return rv
}

func (w_ Workspace) IsVoiceOverEnabled() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isVoiceOverEnabled"))
	return rv
}
