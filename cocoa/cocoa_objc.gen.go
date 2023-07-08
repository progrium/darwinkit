package cocoa

import (
	core "github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"unsafe"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -lobjc -framework Foundation -framework AppKit
#define __OBJC2__ 1
#include <objc/message.h>
#include <stdlib.h>

#include <Foundation/Foundation.h>
#include <AppKit/AppKit.h>

bool cocoa_convertObjCBool(BOOL b) {
	if (b) { return true; }
	return false;
}


void* NSBundle_type_Alloc() {
	return [NSBundle
		alloc];
}
void* NSBundle_type_BundleWithURL(void* url) {
	return [NSBundle
		bundleWithURL: url];
}
void* NSBundle_type_BundleWithPath(void* path) {
	return [NSBundle
		bundleWithPath: path];
}
void* NSBundle_type_BundleWithIdentifier(void* identifier) {
	return [NSBundle
		bundleWithIdentifier: identifier];
}
void* NSBundle_type_URLForResourceWithExtensionSubdirectoryInBundleWithURL(void* name, void* ext, void* subpath, void* bundleURL) {
	return [NSBundle
		URLForResource: name
		withExtension: ext
		subdirectory: subpath
		inBundleWithURL: bundleURL];
}
void* NSBundle_type_URLsForResourcesWithExtensionSubdirectoryInBundleWithURL(void* ext, void* subpath, void* bundleURL) {
	return [NSBundle
		URLsForResourcesWithExtension: ext
		subdirectory: subpath
		inBundleWithURL: bundleURL];
}
void* NSBundle_type_PathForResourceOfTypeInDirectory(void* name, void* ext, void* bundlePath) {
	return [NSBundle
		pathForResource: name
		ofType: ext
		inDirectory: bundlePath];
}
void* NSBundle_type_PathsForResourcesOfTypeInDirectory(void* ext, void* bundlePath) {
	return [NSBundle
		pathsForResourcesOfType: ext
		inDirectory: bundlePath];
}
void* NSBundle_type_PreferredLocalizationsFromArray(void* localizationsArray) {
	return [NSBundle
		preferredLocalizationsFromArray: localizationsArray];
}
void* NSBundle_type_PreferredLocalizationsFromArrayForPreferences(void* localizationsArray, void* preferencesArray) {
	return [NSBundle
		preferredLocalizationsFromArray: localizationsArray
		forPreferences: preferencesArray];
}
void* NSBundle_type_MainBundle() {
	return [NSBundle
		mainBundle];
}
void* NSBundle_type_AllFrameworks() {
	return [NSBundle
		allFrameworks];
}
void* NSBundle_type_AllBundles() {
	return [NSBundle
		allBundles];
}
void* NSSound_type_Alloc() {
	return [NSSound
		alloc];
}
BOOL NSSound_type_CanInitWithPasteboard(void* pasteboard) {
	return [NSSound
		canInitWithPasteboard: pasteboard];
}
void* NSSound_type_SoundUnfilteredTypes() {
	return [NSSound
		soundUnfilteredTypes];
}
void* NSApplication_type_Alloc() {
	return [NSApplication
		alloc];
}
void NSApplication_type_DetachDrawingThreadToTargetWithObject(void* selector, void* target, void* argument) {
	[NSApplication
		detachDrawingThread: selector
		toTarget: target
		withObject: argument];
}
void* NSApplication_type_SharedApplication() {
	return [NSApplication
		sharedApplication];
}
void* NSControl_type_Alloc() {
	return [NSControl
		alloc];
}
void* NSButton_type_Alloc() {
	return [NSButton
		alloc];
}
void* NSButton_type_CheckboxWithTitleTargetAction(void* title, void* target, void* action) {
	return [NSButton
		checkboxWithTitle: title
		target: target
		action: action];
}
void* NSButton_type_ButtonWithImageTargetAction(void* image, void* target, void* action) {
	return [NSButton
		buttonWithImage: image
		target: target
		action: action];
}
void* NSButton_type_RadioButtonWithTitleTargetAction(void* title, void* target, void* action) {
	return [NSButton
		radioButtonWithTitle: title
		target: target
		action: action];
}
void* NSButton_type_ButtonWithTitleImageTargetAction(void* title, void* image, void* target, void* action) {
	return [NSButton
		buttonWithTitle: title
		image: image
		target: target
		action: action];
}
void* NSButton_type_ButtonWithTitleTargetAction(void* title, void* target, void* action) {
	return [NSButton
		buttonWithTitle: title
		target: target
		action: action];
}
void* NSEvent_type_Alloc() {
	return [NSEvent
		alloc];
}
void* NSEvent_type_EventWithEventRef(void* eventRef) {
	return [NSEvent
		eventWithEventRef: eventRef];
}
void NSEvent_type_StopPeriodicEvents() {
	[NSEvent
		stopPeriodicEvents];
}
void NSEvent_type_RemoveMonitor(void* eventMonitor) {
	[NSEvent
		removeMonitor: eventMonitor];
}
unsigned long NSEvent_type_PressedMouseButtons() {
	return [NSEvent
		pressedMouseButtons];
}
NSPoint NSEvent_type_MouseLocation() {
	return [NSEvent
		mouseLocation];
}
BOOL NSEvent_type_MouseCoalescingEnabled() {
	return [NSEvent
		mouseCoalescingEnabled];
}
void NSEvent_type_SetMouseCoalescingEnabled(BOOL value) {
	[NSEvent
		setMouseCoalescingEnabled: value];
}
BOOL NSEvent_type_SwipeTrackingFromScrollEventsEnabled() {
	return [NSEvent
		swipeTrackingFromScrollEventsEnabled];
}
void* NSFont_type_Alloc() {
	return [NSFont
		alloc];
}
void* NSFont_type_FontWithNameSize(void* fontName, double fontSize) {
	return [NSFont
		fontWithName: fontName
		size: fontSize];
}
void* NSFont_type_UserFontOfSize(double fontSize) {
	return [NSFont
		userFontOfSize: fontSize];
}
void* NSFont_type_UserFixedPitchFontOfSize(double fontSize) {
	return [NSFont
		userFixedPitchFontOfSize: fontSize];
}
void* NSFont_type_SystemFontOfSize(double fontSize) {
	return [NSFont
		systemFontOfSize: fontSize];
}
void* NSFont_type_BoldSystemFontOfSize(double fontSize) {
	return [NSFont
		boldSystemFontOfSize: fontSize];
}
void* NSFont_type_LabelFontOfSize(double fontSize) {
	return [NSFont
		labelFontOfSize: fontSize];
}
void* NSFont_type_MessageFontOfSize(double fontSize) {
	return [NSFont
		messageFontOfSize: fontSize];
}
void* NSFont_type_MenuBarFontOfSize(double fontSize) {
	return [NSFont
		menuBarFontOfSize: fontSize];
}
void* NSFont_type_MenuFontOfSize(double fontSize) {
	return [NSFont
		menuFontOfSize: fontSize];
}
void* NSFont_type_ControlContentFontOfSize(double fontSize) {
	return [NSFont
		controlContentFontOfSize: fontSize];
}
void* NSFont_type_TitleBarFontOfSize(double fontSize) {
	return [NSFont
		titleBarFontOfSize: fontSize];
}
void* NSFont_type_PaletteFontOfSize(double fontSize) {
	return [NSFont
		paletteFontOfSize: fontSize];
}
void* NSFont_type_ToolTipsFontOfSize(double fontSize) {
	return [NSFont
		toolTipsFontOfSize: fontSize];
}
void NSFont_type_SetUserFont(void* font) {
	[NSFont
		setUserFont: font];
}
void NSFont_type_SetUserFixedPitchFont(void* font) {
	[NSFont
		setUserFixedPitchFont: font];
}
double NSFont_type_SystemFontSize() {
	return [NSFont
		systemFontSize];
}
double NSFont_type_SmallSystemFontSize() {
	return [NSFont
		smallSystemFontSize];
}
double NSFont_type_LabelFontSize() {
	return [NSFont
		labelFontSize];
}
void* NSImage_type_Alloc() {
	return [NSImage
		alloc];
}
void* NSImage_type_ImageWithSystemSymbolNameAccessibilityDescription(void* symbolName, void* description) {
	return [NSImage
		imageWithSystemSymbolName: symbolName
		accessibilityDescription: description];
}
BOOL NSImage_type_CanInitWithPasteboard(void* pasteboard) {
	return [NSImage
		canInitWithPasteboard: pasteboard];
}
void* NSImage_type_ImageTypes() {
	return [NSImage
		imageTypes];
}
void* NSImage_type_ImageUnfilteredTypes() {
	return [NSImage
		imageUnfilteredTypes];
}
void* NSImageView_type_Alloc() {
	return [NSImageView
		alloc];
}
void* NSImageView_type_ImageViewWithImage(void* image) {
	return [NSImageView
		imageViewWithImage: image];
}
void* NSNib_type_Alloc() {
	return [NSNib
		alloc];
}
void* NSPasteboard_type_Alloc() {
	return [NSPasteboard
		alloc];
}
void* NSPasteboard_type_PasteboardByFilteringFile(void* filename) {
	return [NSPasteboard
		pasteboardByFilteringFile: filename];
}
void* NSPasteboard_type_PasteboardByFilteringTypesInPasteboard(void* pboard) {
	return [NSPasteboard
		pasteboardByFilteringTypesInPasteboard: pboard];
}
void* NSPasteboard_type_PasteboardWithUniqueName() {
	return [NSPasteboard
		pasteboardWithUniqueName];
}
void* NSPasteboard_type_GeneralPasteboard() {
	return [NSPasteboard
		generalPasteboard];
}
void* NSLayoutManager_type_Alloc() {
	return [NSLayoutManager
		alloc];
}
void* NSMenu_type_Alloc() {
	return [NSMenu
		alloc];
}
BOOL NSMenu_type_MenuBarVisible() {
	return [NSMenu
		menuBarVisible];
}
void NSMenu_type_SetMenuBarVisible(BOOL visible) {
	[NSMenu
		setMenuBarVisible: visible];
}
void NSMenu_type_PopUpContextMenuWithEventForView(void* menu, void* event, void* view) {
	[NSMenu
		popUpContextMenu: menu
		withEvent: event
		forView: view];
}
void NSMenu_type_PopUpContextMenuWithEventForViewWithFont(void* menu, void* event, void* view, void* font) {
	[NSMenu
		popUpContextMenu: menu
		withEvent: event
		forView: view
		withFont: font];
}
void* NSPopover_type_Alloc() {
	return [NSPopover
		alloc];
}
void* NSMenuItem_type_Alloc() {
	return [NSMenuItem
		alloc];
}
void* NSMenuItem_type_SeparatorItem() {
	return [NSMenuItem
		separatorItem];
}
BOOL NSMenuItem_type_UsesUserKeyEquivalents() {
	return [NSMenuItem
		usesUserKeyEquivalents];
}
void NSMenuItem_type_SetUsesUserKeyEquivalents(BOOL value) {
	[NSMenuItem
		setUsesUserKeyEquivalents: value];
}
void* NSRunningApplication_type_Alloc() {
	return [NSRunningApplication
		alloc];
}
void* NSRunningApplication_type_RunningApplicationsWithBundleIdentifier(void* bundleIdentifier) {
	return [NSRunningApplication
		runningApplicationsWithBundleIdentifier: bundleIdentifier];
}
void NSRunningApplication_type_TerminateAutomaticallyTerminableApplications() {
	[NSRunningApplication
		terminateAutomaticallyTerminableApplications];
}
void* NSRunningApplication_type_CurrentApplication() {
	return [NSRunningApplication
		currentApplication];
}
void* NSScreen_type_Alloc() {
	return [NSScreen
		alloc];
}
void* NSScreen_type_MainScreen() {
	return [NSScreen
		mainScreen];
}
void* NSScreen_type_DeepestScreen() {
	return [NSScreen
		deepestScreen];
}
void* NSScreen_type_Screens() {
	return [NSScreen
		screens];
}
BOOL NSScreen_type_ScreensHaveSeparateSpaces() {
	return [NSScreen
		screensHaveSeparateSpaces];
}
void* NSStatusBar_type_Alloc() {
	return [NSStatusBar
		alloc];
}
void* NSStatusBar_type_SystemStatusBar() {
	return [NSStatusBar
		systemStatusBar];
}
void* NSStatusBarButton_type_Alloc() {
	return [NSStatusBarButton
		alloc];
}
void* NSStatusItem_type_Alloc() {
	return [NSStatusItem
		alloc];
}
void* NSText_type_Alloc() {
	return [NSText
		alloc];
}
void* NSTextField_type_Alloc() {
	return [NSTextField
		alloc];
}
void* NSTextField_type_LabelWithAttributedString(void* attributedStringValue) {
	return [NSTextField
		labelWithAttributedString: attributedStringValue];
}
void* NSTextField_type_LabelWithString(void* stringValue) {
	return [NSTextField
		labelWithString: stringValue];
}
void* NSTextField_type_TextFieldWithString(void* stringValue) {
	return [NSTextField
		textFieldWithString: stringValue];
}
void* NSTextField_type_WrappingLabelWithString(void* stringValue) {
	return [NSTextField
		wrappingLabelWithString: stringValue];
}
void* NSTextContainer_type_Alloc() {
	return [NSTextContainer
		alloc];
}
void* NSViewController_type_Alloc() {
	return [NSViewController
		alloc];
}
void* NSVisualEffectView_type_Alloc() {
	return [NSVisualEffectView
		alloc];
}
void* NSWindow_type_Alloc() {
	return [NSWindow
		alloc];
}
void* NSWindow_type_WindowWithContentViewController(void* contentViewController) {
	return [NSWindow
		windowWithContentViewController: contentViewController];
}
NSRect NSWindow_type_ContentRectForFrameRectStyleMask(NSRect fRect, unsigned long style) {
	return [NSWindow
		contentRectForFrameRect: fRect
		styleMask: style];
}
NSRect NSWindow_type_FrameRectForContentRectStyleMask(NSRect cRect, unsigned long style) {
	return [NSWindow
		frameRectForContentRect: cRect
		styleMask: style];
}
double NSWindow_type_MinFrameWidthWithTitleStyleMask(void* title, unsigned long style) {
	return [NSWindow
		minFrameWidthWithTitle: title
		styleMask: style];
}
long NSWindow_type_WindowNumberAtPointBelowWindowWithWindowNumber(NSPoint point, long windowNumber) {
	return [NSWindow
		windowNumberAtPoint: point
		belowWindowWithWindowNumber: windowNumber];
}
BOOL NSWindow_type_AllowsAutomaticWindowTabbing() {
	return [NSWindow
		allowsAutomaticWindowTabbing];
}
void NSWindow_type_SetAllowsAutomaticWindowTabbing(BOOL value) {
	[NSWindow
		setAllowsAutomaticWindowTabbing: value];
}
void* NSWorkspace_type_Alloc() {
	return [NSWorkspace
		alloc];
}
void* NSWorkspace_type_SharedWorkspace() {
	return [NSWorkspace
		sharedWorkspace];
}
void* NSColor_type_Alloc() {
	return [NSColor
		alloc];
}
void* NSColor_type_ColorFromPasteboard(void* pasteBoard) {
	return [NSColor
		colorFromPasteboard: pasteBoard];
}
void* NSColor_type_ColorWithRedGreenBlueAlpha(double red, double green, double blue, double alpha) {
	return [NSColor
		colorWithRed: red
		green: green
		blue: blue
		alpha: alpha];
}
BOOL NSColor_type_IgnoresAlpha() {
	return [NSColor
		ignoresAlpha];
}
void NSColor_type_SetIgnoresAlpha(BOOL value) {
	[NSColor
		setIgnoresAlpha: value];
}
void* NSColor_type_SystemCyanColor() {
	return [NSColor
		systemCyanColor];
}
void* NSColor_type_SystemMintColor() {
	return [NSColor
		systemMintColor];
}
void* NSColor_type_ClearColor() {
	return [NSColor
		clearColor];
}
void* NSTextView_type_Alloc() {
	return [NSTextView
		alloc];
}
void NSTextView_type_RegisterForServices() {
	[NSTextView
		registerForServices];
}
void* NSTextView_type_FieldEditor() {
	return [NSTextView
		fieldEditor];
}
BOOL NSTextView_type_StronglyReferencesTextStorage() {
	return [NSTextView
		stronglyReferencesTextStorage];
}
void* NSView_type_Alloc() {
	return [NSView
		alloc];
}
BOOL NSView_type_RequiresConstraintBasedLayout() {
	return [NSView
		requiresConstraintBasedLayout];
}
void* NSView_type_FocusView() {
	return [NSView
		focusView];
}
void* NSView_type_DefaultMenu() {
	return [NSView
		defaultMenu];
}
BOOL NSView_type_CompatibleWithResponsiveScrolling() {
	return [NSView
		compatibleWithResponsiveScrolling];
}


void* NSBundle_inst_URLForAuxiliaryExecutable(void *id, void* executableName) {
	return [(NSBundle*)id
		URLForAuxiliaryExecutable: executableName];
}

void* NSBundle_inst_URLForResourceWithExtension(void *id, void* name, void* ext) {
	return [(NSBundle*)id
		URLForResource: name
		withExtension: ext];
}

void* NSBundle_inst_URLForResourceWithExtensionSubdirectory(void *id, void* name, void* ext, void* subpath) {
	return [(NSBundle*)id
		URLForResource: name
		withExtension: ext
		subdirectory: subpath];
}

void* NSBundle_inst_URLForResourceWithExtensionSubdirectoryLocalization(void *id, void* name, void* ext, void* subpath, void* localizationName) {
	return [(NSBundle*)id
		URLForResource: name
		withExtension: ext
		subdirectory: subpath
		localization: localizationName];
}

void* NSBundle_inst_URLsForResourcesWithExtensionSubdirectory(void *id, void* ext, void* subpath) {
	return [(NSBundle*)id
		URLsForResourcesWithExtension: ext
		subdirectory: subpath];
}

void* NSBundle_inst_URLsForResourcesWithExtensionSubdirectoryLocalization(void *id, void* ext, void* subpath, void* localizationName) {
	return [(NSBundle*)id
		URLsForResourcesWithExtension: ext
		subdirectory: subpath
		localization: localizationName];
}

void* NSBundle_inst_InitWithPath(void *id, void* path) {
	return [(NSBundle*)id
		initWithPath: path];
}

void* NSBundle_inst_InitWithURL(void *id, void* url) {
	return [(NSBundle*)id
		initWithURL: url];
}

BOOL NSBundle_inst_Load(void *id) {
	return [(NSBundle*)id
		load];
}

void* NSBundle_inst_LoadNibNamedOwnerOptions(void *id, void* name, void* owner, void* options) {
	return [(NSBundle*)id
		loadNibNamed: name
		owner: owner
		options: options];
}

void* NSBundle_inst_LocalizedAttributedStringForKeyValueTable(void *id, void* key, void* value, void* tableName) {
	return [(NSBundle*)id
		localizedAttributedStringForKey: key
		value: value
		table: tableName];
}

void* NSBundle_inst_LocalizedStringForKeyValueTable(void *id, void* key, void* value, void* tableName) {
	return [(NSBundle*)id
		localizedStringForKey: key
		value: value
		table: tableName];
}

void* NSBundle_inst_ObjectForInfoDictionaryKey(void *id, void* key) {
	return [(NSBundle*)id
		objectForInfoDictionaryKey: key];
}

void* NSBundle_inst_PathForAuxiliaryExecutable(void *id, void* executableName) {
	return [(NSBundle*)id
		pathForAuxiliaryExecutable: executableName];
}

void* NSBundle_inst_PathForResourceOfType(void *id, void* name, void* ext) {
	return [(NSBundle*)id
		pathForResource: name
		ofType: ext];
}

void* NSBundle_inst_PathForResourceOfTypeInDirectory(void *id, void* name, void* ext, void* subpath) {
	return [(NSBundle*)id
		pathForResource: name
		ofType: ext
		inDirectory: subpath];
}

void* NSBundle_inst_PathForResourceOfTypeInDirectoryForLocalization(void *id, void* name, void* ext, void* subpath, void* localizationName) {
	return [(NSBundle*)id
		pathForResource: name
		ofType: ext
		inDirectory: subpath
		forLocalization: localizationName];
}

void* NSBundle_inst_PathsForResourcesOfTypeInDirectory(void *id, void* ext, void* subpath) {
	return [(NSBundle*)id
		pathsForResourcesOfType: ext
		inDirectory: subpath];
}

void* NSBundle_inst_PathsForResourcesOfTypeInDirectoryForLocalization(void *id, void* ext, void* subpath, void* localizationName) {
	return [(NSBundle*)id
		pathsForResourcesOfType: ext
		inDirectory: subpath
		forLocalization: localizationName];
}

BOOL NSBundle_inst_Unload(void *id) {
	return [(NSBundle*)id
		unload];
}

void* NSBundle_inst_Init(void *id) {
	return [(NSBundle*)id
		init];
}

void* NSBundle_inst_ResourceURL(void *id) {
	return [(NSBundle*)id
		resourceURL];
}

void* NSBundle_inst_ExecutableURL(void *id) {
	return [(NSBundle*)id
		executableURL];
}

void* NSBundle_inst_PrivateFrameworksURL(void *id) {
	return [(NSBundle*)id
		privateFrameworksURL];
}

void* NSBundle_inst_SharedFrameworksURL(void *id) {
	return [(NSBundle*)id
		sharedFrameworksURL];
}

void* NSBundle_inst_BuiltInPlugInsURL(void *id) {
	return [(NSBundle*)id
		builtInPlugInsURL];
}

void* NSBundle_inst_SharedSupportURL(void *id) {
	return [(NSBundle*)id
		sharedSupportURL];
}

void* NSBundle_inst_AppStoreReceiptURL(void *id) {
	return [(NSBundle*)id
		appStoreReceiptURL];
}

void* NSBundle_inst_ResourcePath(void *id) {
	return [(NSBundle*)id
		resourcePath];
}

void* NSBundle_inst_ExecutablePath(void *id) {
	return [(NSBundle*)id
		executablePath];
}

void* NSBundle_inst_PrivateFrameworksPath(void *id) {
	return [(NSBundle*)id
		privateFrameworksPath];
}

void* NSBundle_inst_SharedFrameworksPath(void *id) {
	return [(NSBundle*)id
		sharedFrameworksPath];
}

void* NSBundle_inst_BuiltInPlugInsPath(void *id) {
	return [(NSBundle*)id
		builtInPlugInsPath];
}

void* NSBundle_inst_SharedSupportPath(void *id) {
	return [(NSBundle*)id
		sharedSupportPath];
}

void* NSBundle_inst_BundleURL(void *id) {
	return [(NSBundle*)id
		bundleURL];
}

void* NSBundle_inst_BundlePath(void *id) {
	return [(NSBundle*)id
		bundlePath];
}

void* NSBundle_inst_BundleIdentifier(void *id) {
	return [(NSBundle*)id
		bundleIdentifier];
}

void* NSBundle_inst_InfoDictionary(void *id) {
	return [(NSBundle*)id
		infoDictionary];
}

void* NSBundle_inst_Localizations(void *id) {
	return [(NSBundle*)id
		localizations];
}

void* NSBundle_inst_PreferredLocalizations(void *id) {
	return [(NSBundle*)id
		preferredLocalizations];
}

void* NSBundle_inst_DevelopmentLocalization(void *id) {
	return [(NSBundle*)id
		developmentLocalization];
}

void* NSBundle_inst_LocalizedInfoDictionary(void *id) {
	return [(NSBundle*)id
		localizedInfoDictionary];
}

void* NSBundle_inst_ExecutableArchitectures(void *id) {
	return [(NSBundle*)id
		executableArchitectures];
}

BOOL NSBundle_inst_IsLoaded(void *id) {
	return [(NSBundle*)id
		isLoaded];
}

void* NSSound_inst_InitWithContentsOfFileByReference(void *id, void* path, BOOL byRef) {
	return [(NSSound*)id
		initWithContentsOfFile: path
		byReference: byRef];
}

void* NSSound_inst_InitWithContentsOfURLByReference(void *id, void* url, BOOL byRef) {
	return [(NSSound*)id
		initWithContentsOfURL: url
		byReference: byRef];
}

void* NSSound_inst_InitWithData(void *id, void* data) {
	return [(NSSound*)id
		initWithData: data];
}

void* NSSound_inst_InitWithPasteboard(void *id, void* pasteboard) {
	return [(NSSound*)id
		initWithPasteboard: pasteboard];
}

BOOL NSSound_inst_Pause(void *id) {
	return [(NSSound*)id
		pause];
}

BOOL NSSound_inst_Play(void *id) {
	return [(NSSound*)id
		play];
}

BOOL NSSound_inst_Resume(void *id) {
	return [(NSSound*)id
		resume];
}

BOOL NSSound_inst_Stop(void *id) {
	return [(NSSound*)id
		stop];
}

void NSSound_inst_WriteToPasteboard(void *id, void* pasteboard) {
	[(NSSound*)id
		writeToPasteboard: pasteboard];
}

void* NSSound_inst_Init(void *id) {
	return [(NSSound*)id
		init];
}

void* NSSound_inst_Delegate(void *id) {
	return [(NSSound*)id
		delegate];
}

void NSSound_inst_SetDelegate(void *id, void* value) {
	[(NSSound*)id
		setDelegate: value];
}

BOOL NSSound_inst_Loops(void *id) {
	return [(NSSound*)id
		loops];
}

void NSSound_inst_SetLoops(void *id, BOOL value) {
	[(NSSound*)id
		setLoops: value];
}

BOOL NSSound_inst_IsPlaying(void *id) {
	return [(NSSound*)id
		isPlaying];
}

void NSApplication_inst_ActivateContextHelpMode(void *id, void* sender) {
	[(NSApplication*)id
		activateContextHelpMode: sender];
}

void NSApplication_inst_ActivateIgnoringOtherApps(void *id, BOOL flag) {
	[(NSApplication*)id
		activateIgnoringOtherApps: flag];
}

long NSApplication_inst_ActivationPolicy(void *id) {
	return [(NSApplication*)id
		activationPolicy];
}

void NSApplication_inst_CancelUserAttentionRequest(void *id, long request) {
	[(NSApplication*)id
		cancelUserAttentionRequest: request];
}

void NSApplication_inst_Deactivate(void *id) {
	[(NSApplication*)id
		deactivate];
}

void NSApplication_inst_DisableRelaunchOnLogin(void *id) {
	[(NSApplication*)id
		disableRelaunchOnLogin];
}

void NSApplication_inst_EnableRelaunchOnLogin(void *id) {
	[(NSApplication*)id
		enableRelaunchOnLogin];
}

void NSApplication_inst_FinishLaunching(void *id) {
	[(NSApplication*)id
		finishLaunching];
}

void NSApplication_inst_HideOtherApplications(void *id, void* sender) {
	[(NSApplication*)id
		hideOtherApplications: sender];
}

void NSApplication_inst_PostEventAtStart(void *id, void* event, BOOL flag) {
	[(NSApplication*)id
		postEvent: event
		atStart: flag];
}

void NSApplication_inst_RegisterForRemoteNotifications(void *id) {
	[(NSApplication*)id
		registerForRemoteNotifications];
}

void NSApplication_inst_RegisterUserInterfaceItemSearchHandler(void *id, void* handler) {
	[(NSApplication*)id
		registerUserInterfaceItemSearchHandler: handler];
}

void NSApplication_inst_ReplyToApplicationShouldTerminate(void *id, BOOL shouldTerminate) {
	[(NSApplication*)id
		replyToApplicationShouldTerminate: shouldTerminate];
}

void NSApplication_inst_Run(void *id) {
	[(NSApplication*)id
		run];
}

BOOL NSApplication_inst_SendActionToFrom(void *id, void* action, void* target, void* sender) {
	return [(NSApplication*)id
		sendAction: action
		to: target
		from: sender];
}

void NSApplication_inst_SendEvent(void *id, void* event) {
	[(NSApplication*)id
		sendEvent: event];
}

BOOL NSApplication_inst_SetActivationPolicy(void *id, long activationPolicy) {
	return [(NSApplication*)id
		setActivationPolicy: activationPolicy];
}

void NSApplication_inst_ShowHelp(void *id, void* sender) {
	[(NSApplication*)id
		showHelp: sender];
}

void NSApplication_inst_Stop(void *id, void* sender) {
	[(NSApplication*)id
		stop: sender];
}

void* NSApplication_inst_TargetForAction(void *id, void* action) {
	return [(NSApplication*)id
		targetForAction: action];
}

void* NSApplication_inst_TargetForActionToFrom(void *id, void* action, void* target, void* sender) {
	return [(NSApplication*)id
		targetForAction: action
		to: target
		from: sender];
}

void NSApplication_inst_Terminate(void *id, void* sender) {
	[(NSApplication*)id
		terminate: sender];
}

void NSApplication_inst_ToggleTouchBarCustomizationPalette(void *id, void* sender) {
	[(NSApplication*)id
		toggleTouchBarCustomizationPalette: sender];
}

BOOL NSApplication_inst_TryToPerformWith(void *id, void* action, void* object) {
	return [(NSApplication*)id
		tryToPerform: action
		with: object];
}

void NSApplication_inst_UnhideAllApplications(void *id, void* sender) {
	[(NSApplication*)id
		unhideAllApplications: sender];
}

void NSApplication_inst_UnregisterForRemoteNotifications(void *id) {
	[(NSApplication*)id
		unregisterForRemoteNotifications];
}

void NSApplication_inst_UnregisterUserInterfaceItemSearchHandler(void *id, void* handler) {
	[(NSApplication*)id
		unregisterUserInterfaceItemSearchHandler: handler];
}

void* NSApplication_inst_Init(void *id) {
	return [(NSApplication*)id
		init];
}

void* NSApplication_inst_Delegate(void *id) {
	return [(NSApplication*)id
		delegate];
}

void NSApplication_inst_SetDelegate(void *id, void* value) {
	[(NSApplication*)id
		setDelegate: value];
}

void* NSApplication_inst_CurrentEvent(void *id) {
	return [(NSApplication*)id
		currentEvent];
}

BOOL NSApplication_inst_IsRunning(void *id) {
	return [(NSApplication*)id
		isRunning];
}

BOOL NSApplication_inst_IsActive(void *id) {
	return [(NSApplication*)id
		isActive];
}

BOOL NSApplication_inst_IsRegisteredForRemoteNotifications(void *id) {
	return [(NSApplication*)id
		isRegisteredForRemoteNotifications];
}

void* NSApplication_inst_ApplicationIconImage(void *id) {
	return [(NSApplication*)id
		applicationIconImage];
}

void NSApplication_inst_SetApplicationIconImage(void *id, void* value) {
	[(NSApplication*)id
		setApplicationIconImage: value];
}

void* NSApplication_inst_HelpMenu(void *id) {
	return [(NSApplication*)id
		helpMenu];
}

void NSApplication_inst_SetHelpMenu(void *id, void* value) {
	[(NSApplication*)id
		setHelpMenu: value];
}

void* NSApplication_inst_ServicesProvider(void *id) {
	return [(NSApplication*)id
		servicesProvider];
}

void NSApplication_inst_SetServicesProvider(void *id, void* value) {
	[(NSApplication*)id
		setServicesProvider: value];
}

BOOL NSApplication_inst_IsFullKeyboardAccessEnabled(void *id) {
	return [(NSApplication*)id
		isFullKeyboardAccessEnabled];
}

void* NSApplication_inst_OrderedDocuments(void *id) {
	return [(NSApplication*)id
		orderedDocuments];
}

void* NSApplication_inst_OrderedWindows(void *id) {
	return [(NSApplication*)id
		orderedWindows];
}

void* NSApplication_inst_MainMenu(void *id) {
	return [(NSApplication*)id
		mainMenu];
}

void NSApplication_inst_SetMainMenu(void *id, void* value) {
	[(NSApplication*)id
		setMainMenu: value];
}

BOOL NSControl_inst_AbortEditing(void *id) {
	return [(NSControl*)id
		abortEditing];
}

void* NSControl_inst_CurrentEditor(void *id) {
	return [(NSControl*)id
		currentEditor];
}

void NSControl_inst_DrawWithExpansionFrameInView(void *id, NSRect contentFrame, void* view) {
	[(NSControl*)id
		drawWithExpansionFrame: contentFrame
		inView: view];
}

void NSControl_inst_EditWithFrameEditorDelegateEvent(void *id, NSRect rect, void* textObj, void* delegate, void* event) {
	[(NSControl*)id
		editWithFrame: rect
		editor: textObj
		delegate: delegate
		event: event];
}

void NSControl_inst_EndEditing(void *id, void* textObj) {
	[(NSControl*)id
		endEditing: textObj];
}

NSRect NSControl_inst_ExpansionFrameWithFrame(void *id, NSRect contentFrame) {
	return [(NSControl*)id
		expansionFrameWithFrame: contentFrame];
}

void* NSControl_inst_InitWithFrame(void *id, NSRect frameRect) {
	return [(NSControl*)id
		initWithFrame: frameRect];
}

void NSControl_inst_MouseDown(void *id, void* event) {
	[(NSControl*)id
		mouseDown: event];
}

void NSControl_inst_PerformClick(void *id, void* sender) {
	[(NSControl*)id
		performClick: sender];
}

void NSControl_inst_SelectWithFrameEditorDelegateStartLength(void *id, NSRect rect, void* textObj, void* delegate, long selStart, long selLength) {
	[(NSControl*)id
		selectWithFrame: rect
		editor: textObj
		delegate: delegate
		start: selStart
		length: selLength];
}

BOOL NSControl_inst_SendActionTo(void *id, void* action, void* target) {
	return [(NSControl*)id
		sendAction: action
		to: target];
}

NSSize NSControl_inst_SizeThatFits(void *id, NSSize size) {
	return [(NSControl*)id
		sizeThatFits: size];
}

void NSControl_inst_SizeToFit(void *id) {
	[(NSControl*)id
		sizeToFit];
}

void NSControl_inst_TakeDoubleValueFrom(void *id, void* sender) {
	[(NSControl*)id
		takeDoubleValueFrom: sender];
}

void NSControl_inst_TakeFloatValueFrom(void *id, void* sender) {
	[(NSControl*)id
		takeFloatValueFrom: sender];
}

void NSControl_inst_TakeIntValueFrom(void *id, void* sender) {
	[(NSControl*)id
		takeIntValueFrom: sender];
}

void NSControl_inst_TakeIntegerValueFrom(void *id, void* sender) {
	[(NSControl*)id
		takeIntegerValueFrom: sender];
}

void NSControl_inst_TakeObjectValueFrom(void *id, void* sender) {
	[(NSControl*)id
		takeObjectValueFrom: sender];
}

void NSControl_inst_TakeStringValueFrom(void *id, void* sender) {
	[(NSControl*)id
		takeStringValueFrom: sender];
}

void NSControl_inst_ValidateEditing(void *id) {
	[(NSControl*)id
		validateEditing];
}

void* NSControl_inst_Init(void *id) {
	return [(NSControl*)id
		init];
}

BOOL NSControl_inst_IsEnabled(void *id) {
	return [(NSControl*)id
		isEnabled];
}

void NSControl_inst_SetEnabled(void *id, BOOL value) {
	[(NSControl*)id
		setEnabled: value];
}

int NSControl_inst_IntValue(void *id) {
	return [(NSControl*)id
		intValue];
}

void NSControl_inst_SetIntValue(void *id, int value) {
	[(NSControl*)id
		setIntValue: value];
}

long NSControl_inst_IntegerValue(void *id) {
	return [(NSControl*)id
		integerValue];
}

void NSControl_inst_SetIntegerValue(void *id, long value) {
	[(NSControl*)id
		setIntegerValue: value];
}

void* NSControl_inst_ObjectValue(void *id) {
	return [(NSControl*)id
		objectValue];
}

void NSControl_inst_SetObjectValue(void *id, void* value) {
	[(NSControl*)id
		setObjectValue: value];
}

void* NSControl_inst_StringValue(void *id) {
	return [(NSControl*)id
		stringValue];
}

void NSControl_inst_SetStringValue(void *id, void* value) {
	[(NSControl*)id
		setStringValue: value];
}

void* NSControl_inst_AttributedStringValue(void *id) {
	return [(NSControl*)id
		attributedStringValue];
}

void NSControl_inst_SetAttributedStringValue(void *id, void* value) {
	[(NSControl*)id
		setAttributedStringValue: value];
}

void* NSControl_inst_Font(void *id) {
	return [(NSControl*)id
		font];
}

void NSControl_inst_SetFont(void *id, void* value) {
	[(NSControl*)id
		setFont: value];
}

BOOL NSControl_inst_UsesSingleLineMode(void *id) {
	return [(NSControl*)id
		usesSingleLineMode];
}

void NSControl_inst_SetUsesSingleLineMode(void *id, BOOL value) {
	[(NSControl*)id
		setUsesSingleLineMode: value];
}

BOOL NSControl_inst_AllowsExpansionToolTips(void *id) {
	return [(NSControl*)id
		allowsExpansionToolTips];
}

void NSControl_inst_SetAllowsExpansionToolTips(void *id, BOOL value) {
	[(NSControl*)id
		setAllowsExpansionToolTips: value];
}

BOOL NSControl_inst_IsHighlighted(void *id) {
	return [(NSControl*)id
		isHighlighted];
}

void NSControl_inst_SetHighlighted(void *id, BOOL value) {
	[(NSControl*)id
		setHighlighted: value];
}

void* NSControl_inst_Action(void *id) {
	return [(NSControl*)id
		action];
}

void NSControl_inst_SetAction(void *id, void* value) {
	[(NSControl*)id
		setAction: value];
}

void* NSControl_inst_Target(void *id) {
	return [(NSControl*)id
		target];
}

void NSControl_inst_SetTarget(void *id, void* value) {
	[(NSControl*)id
		setTarget: value];
}

BOOL NSControl_inst_IsContinuous(void *id) {
	return [(NSControl*)id
		isContinuous];
}

void NSControl_inst_SetContinuous(void *id, BOOL value) {
	[(NSControl*)id
		setContinuous: value];
}

long NSControl_inst_Tag(void *id) {
	return [(NSControl*)id
		tag];
}

void NSControl_inst_SetTag(void *id, long value) {
	[(NSControl*)id
		setTag: value];
}

BOOL NSControl_inst_RefusesFirstResponder(void *id) {
	return [(NSControl*)id
		refusesFirstResponder];
}

void NSControl_inst_SetRefusesFirstResponder(void *id, BOOL value) {
	[(NSControl*)id
		setRefusesFirstResponder: value];
}

BOOL NSControl_inst_IgnoresMultiClick(void *id) {
	return [(NSControl*)id
		ignoresMultiClick];
}

void NSControl_inst_SetIgnoresMultiClick(void *id, BOOL value) {
	[(NSControl*)id
		setIgnoresMultiClick: value];
}

void NSButton_inst_CompressWithPrioritizedCompressionOptions(void *id, void* prioritizedOptions) {
	[(NSButton*)id
		compressWithPrioritizedCompressionOptions: prioritizedOptions];
}

void NSButton_inst_Highlight(void *id, BOOL flag) {
	[(NSButton*)id
		highlight: flag];
}

NSSize NSButton_inst_MinimumSizeWithPrioritizedCompressionOptions(void *id, void* prioritizedOptions) {
	return [(NSButton*)id
		minimumSizeWithPrioritizedCompressionOptions: prioritizedOptions];
}

BOOL NSButton_inst_PerformKeyEquivalent(void *id, void* key) {
	return [(NSButton*)id
		performKeyEquivalent: key];
}

void NSButton_inst_SetNextState(void *id) {
	[(NSButton*)id
		setNextState];
}

void* NSButton_inst_Init(void *id) {
	return [(NSButton*)id
		init];
}

void* NSButton_inst_ContentTintColor(void *id) {
	return [(NSButton*)id
		contentTintColor];
}

void NSButton_inst_SetContentTintColor(void *id, void* value) {
	[(NSButton*)id
		setContentTintColor: value];
}

BOOL NSButton_inst_HasDestructiveAction(void *id) {
	return [(NSButton*)id
		hasDestructiveAction];
}

void NSButton_inst_SetHasDestructiveAction(void *id, BOOL value) {
	[(NSButton*)id
		setHasDestructiveAction: value];
}

void* NSButton_inst_AlternateTitle(void *id) {
	return [(NSButton*)id
		alternateTitle];
}

void NSButton_inst_SetAlternateTitle(void *id, void* value) {
	[(NSButton*)id
		setAlternateTitle: value];
}

void* NSButton_inst_AttributedTitle(void *id) {
	return [(NSButton*)id
		attributedTitle];
}

void NSButton_inst_SetAttributedTitle(void *id, void* value) {
	[(NSButton*)id
		setAttributedTitle: value];
}

void* NSButton_inst_AttributedAlternateTitle(void *id) {
	return [(NSButton*)id
		attributedAlternateTitle];
}

void NSButton_inst_SetAttributedAlternateTitle(void *id, void* value) {
	[(NSButton*)id
		setAttributedAlternateTitle: value];
}

void* NSButton_inst_Title(void *id) {
	return [(NSButton*)id
		title];
}

void NSButton_inst_SetTitle(void *id, void* value) {
	[(NSButton*)id
		setTitle: value];
}

void* NSButton_inst_Sound(void *id) {
	return [(NSButton*)id
		sound];
}

void NSButton_inst_SetSound(void *id, void* value) {
	[(NSButton*)id
		setSound: value];
}

BOOL NSButton_inst_IsSpringLoaded(void *id) {
	return [(NSButton*)id
		isSpringLoaded];
}

void NSButton_inst_SetSpringLoaded(void *id, BOOL value) {
	[(NSButton*)id
		setSpringLoaded: value];
}

long NSButton_inst_MaxAcceleratorLevel(void *id) {
	return [(NSButton*)id
		maxAcceleratorLevel];
}

void NSButton_inst_SetMaxAcceleratorLevel(void *id, long value) {
	[(NSButton*)id
		setMaxAcceleratorLevel: value];
}

void* NSButton_inst_Image(void *id) {
	return [(NSButton*)id
		image];
}

void NSButton_inst_SetImage(void *id, void* value) {
	[(NSButton*)id
		setImage: value];
}

void* NSButton_inst_AlternateImage(void *id) {
	return [(NSButton*)id
		alternateImage];
}

void NSButton_inst_SetAlternateImage(void *id, void* value) {
	[(NSButton*)id
		setAlternateImage: value];
}

BOOL NSButton_inst_IsBordered(void *id) {
	return [(NSButton*)id
		isBordered];
}

void NSButton_inst_SetBordered(void *id, BOOL value) {
	[(NSButton*)id
		setBordered: value];
}

BOOL NSButton_inst_IsTransparent(void *id) {
	return [(NSButton*)id
		isTransparent];
}

void NSButton_inst_SetTransparent(void *id, BOOL value) {
	[(NSButton*)id
		setTransparent: value];
}

void* NSButton_inst_BezelColor(void *id) {
	return [(NSButton*)id
		bezelColor];
}

void NSButton_inst_SetBezelColor(void *id, void* value) {
	[(NSButton*)id
		setBezelColor: value];
}

BOOL NSButton_inst_ShowsBorderOnlyWhileMouseInside(void *id) {
	return [(NSButton*)id
		showsBorderOnlyWhileMouseInside];
}

void NSButton_inst_SetShowsBorderOnlyWhileMouseInside(void *id, BOOL value) {
	[(NSButton*)id
		setShowsBorderOnlyWhileMouseInside: value];
}

BOOL NSButton_inst_ImageHugsTitle(void *id) {
	return [(NSButton*)id
		imageHugsTitle];
}

void NSButton_inst_SetImageHugsTitle(void *id, BOOL value) {
	[(NSButton*)id
		setImageHugsTitle: value];
}

BOOL NSButton_inst_AllowsMixedState(void *id) {
	return [(NSButton*)id
		allowsMixedState];
}

void NSButton_inst_SetAllowsMixedState(void *id, BOOL value) {
	[(NSButton*)id
		setAllowsMixedState: value];
}

long NSButton_inst_State(void *id) {
	return [(NSButton*)id
		state];
}

void NSButton_inst_SetState(void *id, long value) {
	[(NSButton*)id
		setState: value];
}

void* NSButton_inst_KeyEquivalent(void *id) {
	return [(NSButton*)id
		keyEquivalent];
}

void NSButton_inst_SetKeyEquivalent(void *id, void* value) {
	[(NSButton*)id
		setKeyEquivalent: value];
}

void* NSEvent_inst_Init(void *id) {
	return [(NSEvent*)id
		init];
}

NSPoint NSEvent_inst_LocationInWindow(void *id) {
	return [(NSEvent*)id
		locationInWindow];
}

void* NSEvent_inst_Window(void *id) {
	return [(NSEvent*)id
		window];
}

long NSEvent_inst_WindowNumber(void *id) {
	return [(NSEvent*)id
		windowNumber];
}

void* NSEvent_inst_EventRef(void *id) {
	return [(NSEvent*)id
		eventRef];
}

void* NSEvent_inst_Characters(void *id) {
	return [(NSEvent*)id
		characters];
}

void* NSEvent_inst_CharactersIgnoringModifiers(void *id) {
	return [(NSEvent*)id
		charactersIgnoringModifiers];
}

BOOL NSEvent_inst_IsARepeat(void *id) {
	return [(NSEvent*)id
		isARepeat];
}

long NSEvent_inst_ButtonNumber(void *id) {
	return [(NSEvent*)id
		buttonNumber];
}

long NSEvent_inst_ClickCount(void *id) {
	return [(NSEvent*)id
		clickCount];
}

long NSEvent_inst_EventNumber(void *id) {
	return [(NSEvent*)id
		eventNumber];
}

long NSEvent_inst_TrackingNumber(void *id) {
	return [(NSEvent*)id
		trackingNumber];
}

void* NSEvent_inst_UserData(void *id) {
	return [(NSEvent*)id
		userData];
}

long NSEvent_inst_Data1(void *id) {
	return [(NSEvent*)id
		data1];
}

long NSEvent_inst_Data2(void *id) {
	return [(NSEvent*)id
		data2];
}

double NSEvent_inst_DeltaX(void *id) {
	return [(NSEvent*)id
		deltaX];
}

double NSEvent_inst_DeltaY(void *id) {
	return [(NSEvent*)id
		deltaY];
}

double NSEvent_inst_DeltaZ(void *id) {
	return [(NSEvent*)id
		deltaZ];
}

long NSEvent_inst_Stage(void *id) {
	return [(NSEvent*)id
		stage];
}

double NSEvent_inst_StageTransition(void *id) {
	return [(NSEvent*)id
		stageTransition];
}

unsigned long NSEvent_inst_CapabilityMask(void *id) {
	return [(NSEvent*)id
		capabilityMask];
}

unsigned long NSEvent_inst_DeviceID(void *id) {
	return [(NSEvent*)id
		deviceID];
}

BOOL NSEvent_inst_IsEnteringProximity(void *id) {
	return [(NSEvent*)id
		isEnteringProximity];
}

unsigned long NSEvent_inst_PointingDeviceID(void *id) {
	return [(NSEvent*)id
		pointingDeviceID];
}

unsigned long NSEvent_inst_PointingDeviceSerialNumber(void *id) {
	return [(NSEvent*)id
		pointingDeviceSerialNumber];
}

unsigned long NSEvent_inst_SystemTabletID(void *id) {
	return [(NSEvent*)id
		systemTabletID];
}

unsigned long NSEvent_inst_TabletID(void *id) {
	return [(NSEvent*)id
		tabletID];
}

unsigned long NSEvent_inst_VendorID(void *id) {
	return [(NSEvent*)id
		vendorID];
}

unsigned long NSEvent_inst_VendorPointingDeviceType(void *id) {
	return [(NSEvent*)id
		vendorPointingDeviceType];
}

long NSEvent_inst_AbsoluteX(void *id) {
	return [(NSEvent*)id
		absoluteX];
}

long NSEvent_inst_AbsoluteY(void *id) {
	return [(NSEvent*)id
		absoluteY];
}

long NSEvent_inst_AbsoluteZ(void *id) {
	return [(NSEvent*)id
		absoluteZ];
}

NSPoint NSEvent_inst_Tilt(void *id) {
	return [(NSEvent*)id
		tilt];
}

void* NSEvent_inst_VendorDefined(void *id) {
	return [(NSEvent*)id
		vendorDefined];
}

double NSEvent_inst_Magnification(void *id) {
	return [(NSEvent*)id
		magnification];
}

BOOL NSEvent_inst_HasPreciseScrollingDeltas(void *id) {
	return [(NSEvent*)id
		hasPreciseScrollingDeltas];
}

double NSEvent_inst_ScrollingDeltaX(void *id) {
	return [(NSEvent*)id
		scrollingDeltaX];
}

double NSEvent_inst_ScrollingDeltaY(void *id) {
	return [(NSEvent*)id
		scrollingDeltaY];
}

BOOL NSEvent_inst_IsDirectionInvertedFromDevice(void *id) {
	return [(NSEvent*)id
		isDirectionInvertedFromDevice];
}

void* NSFont_inst_FontWithSize(void *id, double fontSize) {
	return [(NSFont*)id
		fontWithSize: fontSize];
}

void NSFont_inst_Set(void *id) {
	[(NSFont*)id
		set];
}

void* NSFont_inst_Init(void *id) {
	return [(NSFont*)id
		init];
}

double NSFont_inst_PointSize(void *id) {
	return [(NSFont*)id
		pointSize];
}

BOOL NSFont_inst_IsFixedPitch(void *id) {
	return [(NSFont*)id
		isFixedPitch];
}

unsigned long NSFont_inst_MostCompatibleStringEncoding(void *id) {
	return [(NSFont*)id
		mostCompatibleStringEncoding];
}

unsigned long NSFont_inst_NumberOfGlyphs(void *id) {
	return [(NSFont*)id
		numberOfGlyphs];
}

void* NSFont_inst_DisplayName(void *id) {
	return [(NSFont*)id
		displayName];
}

void* NSFont_inst_FamilyName(void *id) {
	return [(NSFont*)id
		familyName];
}

void* NSFont_inst_FontName(void *id) {
	return [(NSFont*)id
		fontName];
}

BOOL NSFont_inst_IsVertical(void *id) {
	return [(NSFont*)id
		isVertical];
}

void* NSFont_inst_VerticalFont(void *id) {
	return [(NSFont*)id
		verticalFont];
}

void NSImage_inst_AddRepresentations(void *id, void* imageReps) {
	[(NSImage*)id
		addRepresentations: imageReps];
}

void NSImage_inst_CancelIncrementalLoad(void *id) {
	[(NSImage*)id
		cancelIncrementalLoad];
}

void NSImage_inst_DrawInRect(void *id, NSRect rect) {
	[(NSImage*)id
		drawInRect: rect];
}

void* NSImage_inst_InitByReferencingFile(void *id, void* fileName) {
	return [(NSImage*)id
		initByReferencingFile: fileName];
}

void* NSImage_inst_InitByReferencingURL(void *id, void* url) {
	return [(NSImage*)id
		initByReferencingURL: url];
}

void* NSImage_inst_InitWithContentsOfFile(void *id, void* fileName) {
	return [(NSImage*)id
		initWithContentsOfFile: fileName];
}

void* NSImage_inst_InitWithContentsOfURL(void *id, void* url) {
	return [(NSImage*)id
		initWithContentsOfURL: url];
}

void* NSImage_inst_InitWithData(void *id, void* data) {
	return [(NSImage*)id
		initWithData: data];
}

void* NSImage_inst_InitWithDataIgnoringOrientation(void *id, void* data) {
	return [(NSImage*)id
		initWithDataIgnoringOrientation: data];
}

void* NSImage_inst_InitWithPasteboard(void *id, void* pasteboard) {
	return [(NSImage*)id
		initWithPasteboard: pasteboard];
}

void* NSImage_inst_InitWithSize(void *id, NSSize size) {
	return [(NSImage*)id
		initWithSize: size];
}

BOOL NSImage_inst_IsTemplate(void *id) {
	return [(NSImage*)id
		isTemplate];
}

void* NSImage_inst_LayerContentsForContentsScale(void *id, double layerContentsScale) {
	return [(NSImage*)id
		layerContentsForContentsScale: layerContentsScale];
}

void NSImage_inst_LockFocus(void *id) {
	[(NSImage*)id
		lockFocus];
}

void NSImage_inst_LockFocusFlipped(void *id, BOOL flipped) {
	[(NSImage*)id
		lockFocusFlipped: flipped];
}

void NSImage_inst_Recache(void *id) {
	[(NSImage*)id
		recache];
}

double NSImage_inst_RecommendedLayerContentsScale(void *id, double preferredContentsScale) {
	return [(NSImage*)id
		recommendedLayerContentsScale: preferredContentsScale];
}

void NSImage_inst_UnlockFocus(void *id) {
	[(NSImage*)id
		unlockFocus];
}

void* NSImage_inst_Init(void *id) {
	return [(NSImage*)id
		init];
}

void* NSImage_inst_Delegate(void *id) {
	return [(NSImage*)id
		delegate];
}

void NSImage_inst_SetDelegate(void *id, void* value) {
	[(NSImage*)id
		setDelegate: value];
}

NSSize NSImage_inst_Size(void *id) {
	return [(NSImage*)id
		size];
}

void NSImage_inst_SetSize(void *id, NSSize value) {
	[(NSImage*)id
		setSize: value];
}

void NSImage_inst_SetTemplate(void *id, BOOL value) {
	[(NSImage*)id
		setTemplate: value];
}

void* NSImage_inst_Representations(void *id) {
	return [(NSImage*)id
		representations];
}

BOOL NSImage_inst_PrefersColorMatch(void *id) {
	return [(NSImage*)id
		prefersColorMatch];
}

void NSImage_inst_SetPrefersColorMatch(void *id, BOOL value) {
	[(NSImage*)id
		setPrefersColorMatch: value];
}

BOOL NSImage_inst_UsesEPSOnResolutionMismatch(void *id) {
	return [(NSImage*)id
		usesEPSOnResolutionMismatch];
}

void NSImage_inst_SetUsesEPSOnResolutionMismatch(void *id, BOOL value) {
	[(NSImage*)id
		setUsesEPSOnResolutionMismatch: value];
}

BOOL NSImage_inst_MatchesOnMultipleResolution(void *id) {
	return [(NSImage*)id
		matchesOnMultipleResolution];
}

void NSImage_inst_SetMatchesOnMultipleResolution(void *id, BOOL value) {
	[(NSImage*)id
		setMatchesOnMultipleResolution: value];
}

BOOL NSImage_inst_IsValid(void *id) {
	return [(NSImage*)id
		isValid];
}

void* NSImage_inst_BackgroundColor(void *id) {
	return [(NSImage*)id
		backgroundColor];
}

void NSImage_inst_SetBackgroundColor(void *id, void* value) {
	[(NSImage*)id
		setBackgroundColor: value];
}

NSRect NSImage_inst_AlignmentRect(void *id) {
	return [(NSImage*)id
		alignmentRect];
}

void NSImage_inst_SetAlignmentRect(void *id, NSRect value) {
	[(NSImage*)id
		setAlignmentRect: value];
}

void* NSImage_inst_TIFFRepresentation(void *id) {
	return [(NSImage*)id
		TIFFRepresentation];
}

void* NSImage_inst_AccessibilityDescription(void *id) {
	return [(NSImage*)id
		accessibilityDescription];
}

void NSImage_inst_SetAccessibilityDescription(void *id, void* value) {
	[(NSImage*)id
		setAccessibilityDescription: value];
}

BOOL NSImage_inst_MatchesOnlyOnBestFittingAxis(void *id) {
	return [(NSImage*)id
		matchesOnlyOnBestFittingAxis];
}

void NSImage_inst_SetMatchesOnlyOnBestFittingAxis(void *id, BOOL value) {
	[(NSImage*)id
		setMatchesOnlyOnBestFittingAxis: value];
}

void* NSImageView_inst_Init(void *id) {
	return [(NSImageView*)id
		init];
}

void* NSImageView_inst_Image(void *id) {
	return [(NSImageView*)id
		image];
}

void NSImageView_inst_SetImage(void *id, void* value) {
	[(NSImageView*)id
		setImage: value];
}

BOOL NSImageView_inst_Animates(void *id) {
	return [(NSImageView*)id
		animates];
}

void NSImageView_inst_SetAnimates(void *id, BOOL value) {
	[(NSImageView*)id
		setAnimates: value];
}

BOOL NSImageView_inst_IsEditable(void *id) {
	return [(NSImageView*)id
		isEditable];
}

void NSImageView_inst_SetEditable(void *id, BOOL value) {
	[(NSImageView*)id
		setEditable: value];
}

BOOL NSImageView_inst_AllowsCutCopyPaste(void *id) {
	return [(NSImageView*)id
		allowsCutCopyPaste];
}

void NSImageView_inst_SetAllowsCutCopyPaste(void *id, BOOL value) {
	[(NSImageView*)id
		setAllowsCutCopyPaste: value];
}

void* NSImageView_inst_ContentTintColor(void *id) {
	return [(NSImageView*)id
		contentTintColor];
}

void NSImageView_inst_SetContentTintColor(void *id, void* value) {
	[(NSImageView*)id
		setContentTintColor: value];
}

void* NSNib_inst_InitWithNibDataBundle(void *id, void* nibData, void* bundle) {
	return [(NSNib*)id
		initWithNibData: nibData
		bundle: bundle];
}

BOOL NSNib_inst_InstantiateWithOwnerTopLevelObjects(void *id, void* owner, void* topLevelObjects) {
	return [(NSNib*)id
		instantiateWithOwner: owner
		topLevelObjects: topLevelObjects];
}

void* NSNib_inst_Init(void *id) {
	return [(NSNib*)id
		init];
}

long NSPasteboard_inst_AddTypesOwner(void *id, void* newTypes, void* newOwner) {
	return [(NSPasteboard*)id
		addTypes: newTypes
		owner: newOwner];
}

BOOL NSPasteboard_inst_CanReadItemWithDataConformingToTypes(void *id, void* types) {
	return [(NSPasteboard*)id
		canReadItemWithDataConformingToTypes: types];
}

BOOL NSPasteboard_inst_CanReadObjectForClassesOptions(void *id, void* classArray, void* options) {
	return [(NSPasteboard*)id
		canReadObjectForClasses: classArray
		options: options];
}

long NSPasteboard_inst_ClearContents(void *id) {
	return [(NSPasteboard*)id
		clearContents];
}

long NSPasteboard_inst_DeclareTypesOwner(void *id, void* newTypes, void* newOwner) {
	return [(NSPasteboard*)id
		declareTypes: newTypes
		owner: newOwner];
}

void* NSPasteboard_inst_ReadObjectsForClassesOptions(void *id, void* classArray, void* options) {
	return [(NSPasteboard*)id
		readObjectsForClasses: classArray
		options: options];
}

void NSPasteboard_inst_ReleaseGlobally(void *id) {
	[(NSPasteboard*)id
		releaseGlobally];
}

BOOL NSPasteboard_inst_WriteFileContents(void *id, void* filename) {
	return [(NSPasteboard*)id
		writeFileContents: filename];
}

BOOL NSPasteboard_inst_WriteObjects(void *id, void* objects) {
	return [(NSPasteboard*)id
		writeObjects: objects];
}

void* NSPasteboard_inst_Init(void *id) {
	return [(NSPasteboard*)id
		init];
}

void* NSPasteboard_inst_PasteboardItems(void *id) {
	return [(NSPasteboard*)id
		pasteboardItems];
}

void* NSPasteboard_inst_Types(void *id) {
	return [(NSPasteboard*)id
		types];
}

long NSPasteboard_inst_ChangeCount(void *id) {
	return [(NSPasteboard*)id
		changeCount];
}

void NSLayoutManager_inst_AddTextContainer(void *id, void* container) {
	[(NSLayoutManager*)id
		addTextContainer: container];
}

NSSize NSLayoutManager_inst_AttachmentSizeForGlyphAtIndex(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		attachmentSizeForGlyphAtIndex: glyphIndex];
}

unsigned long NSLayoutManager_inst_CharacterIndexForGlyphAtIndex(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		characterIndexForGlyphAtIndex: glyphIndex];
}

double NSLayoutManager_inst_DefaultBaselineOffsetForFont(void *id, void* theFont) {
	return [(NSLayoutManager*)id
		defaultBaselineOffsetForFont: theFont];
}

double NSLayoutManager_inst_DefaultLineHeightForFont(void *id, void* theFont) {
	return [(NSLayoutManager*)id
		defaultLineHeightForFont: theFont];
}

BOOL NSLayoutManager_inst_DrawsOutsideLineFragmentForGlyphAtIndex(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		drawsOutsideLineFragmentForGlyphAtIndex: glyphIndex];
}

void NSLayoutManager_inst_EnsureLayoutForBoundingRectInTextContainer(void *id, NSRect bounds, void* container) {
	[(NSLayoutManager*)id
		ensureLayoutForBoundingRect: bounds
		inTextContainer: container];
}

void NSLayoutManager_inst_EnsureLayoutForTextContainer(void *id, void* container) {
	[(NSLayoutManager*)id
		ensureLayoutForTextContainer: container];
}

unsigned long NSLayoutManager_inst_FirstUnlaidCharacterIndex(void *id) {
	return [(NSLayoutManager*)id
		firstUnlaidCharacterIndex];
}

unsigned long NSLayoutManager_inst_FirstUnlaidGlyphIndex(void *id) {
	return [(NSLayoutManager*)id
		firstUnlaidGlyphIndex];
}

unsigned long NSLayoutManager_inst_GlyphIndexForCharacterAtIndex(void *id, unsigned long charIndex) {
	return [(NSLayoutManager*)id
		glyphIndexForCharacterAtIndex: charIndex];
}

void* NSLayoutManager_inst_Init(void *id) {
	return [(NSLayoutManager*)id
		init];
}

void NSLayoutManager_inst_InsertTextContainerAtIndex(void *id, void* container, unsigned long index) {
	[(NSLayoutManager*)id
		insertTextContainer: container
		atIndex: index];
}

BOOL NSLayoutManager_inst_IsValidGlyphIndex(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		isValidGlyphIndex: glyphIndex];
}

BOOL NSLayoutManager_inst_LayoutManagerOwnsFirstResponderInWindow(void *id, void* window) {
	return [(NSLayoutManager*)id
		layoutManagerOwnsFirstResponderInWindow: window];
}

BOOL NSLayoutManager_inst_NotShownAttributeForGlyphAtIndex(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		notShownAttributeForGlyphAtIndex: glyphIndex];
}

void NSLayoutManager_inst_RemoveTextContainerAtIndex(void *id, unsigned long index) {
	[(NSLayoutManager*)id
		removeTextContainerAtIndex: index];
}

void NSLayoutManager_inst_SetDrawsOutsideLineFragmentForGlyphAtIndex(void *id, BOOL flag, unsigned long glyphIndex) {
	[(NSLayoutManager*)id
		setDrawsOutsideLineFragment: flag
		forGlyphAtIndex: glyphIndex];
}

void NSLayoutManager_inst_SetExtraLineFragmentRectUsedRectTextContainer(void *id, NSRect fragmentRect, NSRect usedRect, void* container) {
	[(NSLayoutManager*)id
		setExtraLineFragmentRect: fragmentRect
		usedRect: usedRect
		textContainer: container];
}

void NSLayoutManager_inst_SetNotShownAttributeForGlyphAtIndex(void *id, BOOL flag, unsigned long glyphIndex) {
	[(NSLayoutManager*)id
		setNotShownAttribute: flag
		forGlyphAtIndex: glyphIndex];
}

void NSLayoutManager_inst_TextContainerChangedGeometry(void *id, void* container) {
	[(NSLayoutManager*)id
		textContainerChangedGeometry: container];
}

void NSLayoutManager_inst_TextContainerChangedTextView(void *id, void* container) {
	[(NSLayoutManager*)id
		textContainerChangedTextView: container];
}

NSRect NSLayoutManager_inst_UsedRectForTextContainer(void *id, void* container) {
	return [(NSLayoutManager*)id
		usedRectForTextContainer: container];
}

void* NSLayoutManager_inst_Delegate(void *id) {
	return [(NSLayoutManager*)id
		delegate];
}

void NSLayoutManager_inst_SetDelegate(void *id, void* value) {
	[(NSLayoutManager*)id
		setDelegate: value];
}

BOOL NSLayoutManager_inst_AllowsNonContiguousLayout(void *id) {
	return [(NSLayoutManager*)id
		allowsNonContiguousLayout];
}

void NSLayoutManager_inst_SetAllowsNonContiguousLayout(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setAllowsNonContiguousLayout: value];
}

BOOL NSLayoutManager_inst_HasNonContiguousLayout(void *id) {
	return [(NSLayoutManager*)id
		hasNonContiguousLayout];
}

BOOL NSLayoutManager_inst_ShowsInvisibleCharacters(void *id) {
	return [(NSLayoutManager*)id
		showsInvisibleCharacters];
}

void NSLayoutManager_inst_SetShowsInvisibleCharacters(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setShowsInvisibleCharacters: value];
}

BOOL NSLayoutManager_inst_ShowsControlCharacters(void *id) {
	return [(NSLayoutManager*)id
		showsControlCharacters];
}

void NSLayoutManager_inst_SetShowsControlCharacters(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setShowsControlCharacters: value];
}

BOOL NSLayoutManager_inst_UsesFontLeading(void *id) {
	return [(NSLayoutManager*)id
		usesFontLeading];
}

void NSLayoutManager_inst_SetUsesFontLeading(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setUsesFontLeading: value];
}

BOOL NSLayoutManager_inst_BackgroundLayoutEnabled(void *id) {
	return [(NSLayoutManager*)id
		backgroundLayoutEnabled];
}

void NSLayoutManager_inst_SetBackgroundLayoutEnabled(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setBackgroundLayoutEnabled: value];
}

BOOL NSLayoutManager_inst_LimitsLayoutForSuspiciousContents(void *id) {
	return [(NSLayoutManager*)id
		limitsLayoutForSuspiciousContents];
}

void NSLayoutManager_inst_SetLimitsLayoutForSuspiciousContents(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setLimitsLayoutForSuspiciousContents: value];
}

BOOL NSLayoutManager_inst_UsesDefaultHyphenation(void *id) {
	return [(NSLayoutManager*)id
		usesDefaultHyphenation];
}

void NSLayoutManager_inst_SetUsesDefaultHyphenation(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setUsesDefaultHyphenation: value];
}

void* NSLayoutManager_inst_TextContainers(void *id) {
	return [(NSLayoutManager*)id
		textContainers];
}

unsigned long NSLayoutManager_inst_NumberOfGlyphs(void *id) {
	return [(NSLayoutManager*)id
		numberOfGlyphs];
}

NSRect NSLayoutManager_inst_ExtraLineFragmentRect(void *id) {
	return [(NSLayoutManager*)id
		extraLineFragmentRect];
}

void* NSLayoutManager_inst_ExtraLineFragmentTextContainer(void *id) {
	return [(NSLayoutManager*)id
		extraLineFragmentTextContainer];
}

NSRect NSLayoutManager_inst_ExtraLineFragmentUsedRect(void *id) {
	return [(NSLayoutManager*)id
		extraLineFragmentUsedRect];
}

void* NSLayoutManager_inst_FirstTextView(void *id) {
	return [(NSLayoutManager*)id
		firstTextView];
}

void* NSLayoutManager_inst_TextViewForBeginningOfSelection(void *id) {
	return [(NSLayoutManager*)id
		textViewForBeginningOfSelection];
}

void NSMenu_inst_AddItem(void *id, void* newItem) {
	[(NSMenu*)id
		addItem: newItem];
}

void* NSMenu_inst_AddItemWithTitleActionKeyEquivalent(void *id, void* string, void* selector, void* charCode) {
	return [(NSMenu*)id
		addItemWithTitle: string
		action: selector
		keyEquivalent: charCode];
}

void NSMenu_inst_CancelTracking(void *id) {
	[(NSMenu*)id
		cancelTracking];
}

void NSMenu_inst_CancelTrackingWithoutAnimation(void *id) {
	[(NSMenu*)id
		cancelTrackingWithoutAnimation];
}

long NSMenu_inst_IndexOfItem(void *id, void* item) {
	return [(NSMenu*)id
		indexOfItem: item];
}

long NSMenu_inst_IndexOfItemWithRepresentedObject(void *id, void* object) {
	return [(NSMenu*)id
		indexOfItemWithRepresentedObject: object];
}

long NSMenu_inst_IndexOfItemWithSubmenu(void *id, void* submenu) {
	return [(NSMenu*)id
		indexOfItemWithSubmenu: submenu];
}

long NSMenu_inst_IndexOfItemWithTag(void *id, long tag) {
	return [(NSMenu*)id
		indexOfItemWithTag: tag];
}

long NSMenu_inst_IndexOfItemWithTargetAndAction(void *id, void* target, void* actionSelector) {
	return [(NSMenu*)id
		indexOfItemWithTarget: target
		andAction: actionSelector];
}

long NSMenu_inst_IndexOfItemWithTitle(void *id, void* title) {
	return [(NSMenu*)id
		indexOfItemWithTitle: title];
}

void* NSMenu_inst_InitWithTitle(void *id, void* title) {
	return [(NSMenu*)id
		initWithTitle: title];
}

void NSMenu_inst_InsertItemAtIndex(void *id, void* newItem, long index) {
	[(NSMenu*)id
		insertItem: newItem
		atIndex: index];
}

void* NSMenu_inst_InsertItemWithTitleActionKeyEquivalentAtIndex(void *id, void* string, void* selector, void* charCode, long index) {
	return [(NSMenu*)id
		insertItemWithTitle: string
		action: selector
		keyEquivalent: charCode
		atIndex: index];
}

void* NSMenu_inst_ItemAtIndex(void *id, long index) {
	return [(NSMenu*)id
		itemAtIndex: index];
}

void NSMenu_inst_ItemChanged(void *id, void* item) {
	[(NSMenu*)id
		itemChanged: item];
}

void* NSMenu_inst_ItemWithTag(void *id, long tag) {
	return [(NSMenu*)id
		itemWithTag: tag];
}

void* NSMenu_inst_ItemWithTitle(void *id, void* title) {
	return [(NSMenu*)id
		itemWithTitle: title];
}

void NSMenu_inst_PerformActionForItemAtIndex(void *id, long index) {
	[(NSMenu*)id
		performActionForItemAtIndex: index];
}

BOOL NSMenu_inst_PerformKeyEquivalent(void *id, void* event) {
	return [(NSMenu*)id
		performKeyEquivalent: event];
}

BOOL NSMenu_inst_PopUpMenuPositioningItemAtLocationInView(void *id, void* item, NSPoint location, void* view) {
	return [(NSMenu*)id
		popUpMenuPositioningItem: item
		atLocation: location
		inView: view];
}

void NSMenu_inst_RemoveAllItems(void *id) {
	[(NSMenu*)id
		removeAllItems];
}

void NSMenu_inst_RemoveItem(void *id, void* item) {
	[(NSMenu*)id
		removeItem: item];
}

void NSMenu_inst_RemoveItemAtIndex(void *id, long index) {
	[(NSMenu*)id
		removeItemAtIndex: index];
}

void NSMenu_inst_SetSubmenuForItem(void *id, void* menu, void* item) {
	[(NSMenu*)id
		setSubmenu: menu
		forItem: item];
}

void NSMenu_inst_SubmenuAction(void *id, void* sender) {
	[(NSMenu*)id
		submenuAction: sender];
}

void NSMenu_inst_Update(void *id) {
	[(NSMenu*)id
		update];
}

void* NSMenu_inst_Init(void *id) {
	return [(NSMenu*)id
		init];
}

double NSMenu_inst_MenuBarHeight(void *id) {
	return [(NSMenu*)id
		menuBarHeight];
}

long NSMenu_inst_NumberOfItems(void *id) {
	return [(NSMenu*)id
		numberOfItems];
}

void* NSMenu_inst_ItemArray(void *id) {
	return [(NSMenu*)id
		itemArray];
}

void NSMenu_inst_SetItemArray(void *id, void* value) {
	[(NSMenu*)id
		setItemArray: value];
}

void* NSMenu_inst_Supermenu(void *id) {
	return [(NSMenu*)id
		supermenu];
}

void NSMenu_inst_SetSupermenu(void *id, void* value) {
	[(NSMenu*)id
		setSupermenu: value];
}

BOOL NSMenu_inst_AutoenablesItems(void *id) {
	return [(NSMenu*)id
		autoenablesItems];
}

void NSMenu_inst_SetAutoenablesItems(void *id, BOOL value) {
	[(NSMenu*)id
		setAutoenablesItems: value];
}

void* NSMenu_inst_Font(void *id) {
	return [(NSMenu*)id
		font];
}

void NSMenu_inst_SetFont(void *id, void* value) {
	[(NSMenu*)id
		setFont: value];
}

void* NSMenu_inst_Title(void *id) {
	return [(NSMenu*)id
		title];
}

void NSMenu_inst_SetTitle(void *id, void* value) {
	[(NSMenu*)id
		setTitle: value];
}

double NSMenu_inst_MinimumWidth(void *id) {
	return [(NSMenu*)id
		minimumWidth];
}

void NSMenu_inst_SetMinimumWidth(void *id, double value) {
	[(NSMenu*)id
		setMinimumWidth: value];
}

NSSize NSMenu_inst_Size(void *id) {
	return [(NSMenu*)id
		size];
}

BOOL NSMenu_inst_AllowsContextMenuPlugIns(void *id) {
	return [(NSMenu*)id
		allowsContextMenuPlugIns];
}

void NSMenu_inst_SetAllowsContextMenuPlugIns(void *id, BOOL value) {
	[(NSMenu*)id
		setAllowsContextMenuPlugIns: value];
}

BOOL NSMenu_inst_ShowsStateColumn(void *id) {
	return [(NSMenu*)id
		showsStateColumn];
}

void NSMenu_inst_SetShowsStateColumn(void *id, BOOL value) {
	[(NSMenu*)id
		setShowsStateColumn: value];
}

void* NSMenu_inst_HighlightedItem(void *id) {
	return [(NSMenu*)id
		highlightedItem];
}

void* NSMenu_inst_Delegate(void *id) {
	return [(NSMenu*)id
		delegate];
}

void NSMenu_inst_SetDelegate(void *id, void* value) {
	[(NSMenu*)id
		setDelegate: value];
}

void NSPopover_inst_Close(void *id) {
	[(NSPopover*)id
		close];
}

void* NSPopover_inst_Init(void *id) {
	return [(NSPopover*)id
		init];
}

void NSPopover_inst_PerformClose(void *id, void* sender) {
	[(NSPopover*)id
		performClose: sender];
}

long NSPopover_inst_Behavior(void *id) {
	return [(NSPopover*)id
		behavior];
}

void NSPopover_inst_SetBehavior(void *id, long value) {
	[(NSPopover*)id
		setBehavior: value];
}

NSRect NSPopover_inst_PositioningRect(void *id) {
	return [(NSPopover*)id
		positioningRect];
}

void NSPopover_inst_SetPositioningRect(void *id, NSRect value) {
	[(NSPopover*)id
		setPositioningRect: value];
}

BOOL NSPopover_inst_Animates(void *id) {
	return [(NSPopover*)id
		animates];
}

void NSPopover_inst_SetAnimates(void *id, BOOL value) {
	[(NSPopover*)id
		setAnimates: value];
}

NSSize NSPopover_inst_ContentSize(void *id) {
	return [(NSPopover*)id
		contentSize];
}

void NSPopover_inst_SetContentSize(void *id, NSSize value) {
	[(NSPopover*)id
		setContentSize: value];
}

BOOL NSPopover_inst_IsShown(void *id) {
	return [(NSPopover*)id
		isShown];
}

BOOL NSPopover_inst_IsDetached(void *id) {
	return [(NSPopover*)id
		isDetached];
}

void* NSMenuItem_inst_InitWithTitleActionKeyEquivalent(void *id, void* string, void* selector, void* charCode) {
	return [(NSMenuItem*)id
		initWithTitle: string
		action: selector
		keyEquivalent: charCode];
}

void* NSMenuItem_inst_Init(void *id) {
	return [(NSMenuItem*)id
		init];
}

BOOL NSMenuItem_inst_IsEnabled(void *id) {
	return [(NSMenuItem*)id
		isEnabled];
}

void NSMenuItem_inst_SetEnabled(void *id, BOOL value) {
	[(NSMenuItem*)id
		setEnabled: value];
}

BOOL NSMenuItem_inst_IsHidden(void *id) {
	return [(NSMenuItem*)id
		isHidden];
}

void NSMenuItem_inst_SetHidden(void *id, BOOL value) {
	[(NSMenuItem*)id
		setHidden: value];
}

BOOL NSMenuItem_inst_IsHiddenOrHasHiddenAncestor(void *id) {
	return [(NSMenuItem*)id
		isHiddenOrHasHiddenAncestor];
}

void* NSMenuItem_inst_Target(void *id) {
	return [(NSMenuItem*)id
		target];
}

void NSMenuItem_inst_SetTarget(void *id, void* value) {
	[(NSMenuItem*)id
		setTarget: value];
}

void* NSMenuItem_inst_Action(void *id) {
	return [(NSMenuItem*)id
		action];
}

void NSMenuItem_inst_SetAction(void *id, void* value) {
	[(NSMenuItem*)id
		setAction: value];
}

void* NSMenuItem_inst_Title(void *id) {
	return [(NSMenuItem*)id
		title];
}

void NSMenuItem_inst_SetTitle(void *id, void* value) {
	[(NSMenuItem*)id
		setTitle: value];
}

void* NSMenuItem_inst_AttributedTitle(void *id) {
	return [(NSMenuItem*)id
		attributedTitle];
}

void NSMenuItem_inst_SetAttributedTitle(void *id, void* value) {
	[(NSMenuItem*)id
		setAttributedTitle: value];
}

long NSMenuItem_inst_Tag(void *id) {
	return [(NSMenuItem*)id
		tag];
}

void NSMenuItem_inst_SetTag(void *id, long value) {
	[(NSMenuItem*)id
		setTag: value];
}

long NSMenuItem_inst_State(void *id) {
	return [(NSMenuItem*)id
		state];
}

void NSMenuItem_inst_SetState(void *id, long value) {
	[(NSMenuItem*)id
		setState: value];
}

void* NSMenuItem_inst_Image(void *id) {
	return [(NSMenuItem*)id
		image];
}

void NSMenuItem_inst_SetImage(void *id, void* value) {
	[(NSMenuItem*)id
		setImage: value];
}

void* NSMenuItem_inst_OnStateImage(void *id) {
	return [(NSMenuItem*)id
		onStateImage];
}

void NSMenuItem_inst_SetOnStateImage(void *id, void* value) {
	[(NSMenuItem*)id
		setOnStateImage: value];
}

void* NSMenuItem_inst_OffStateImage(void *id) {
	return [(NSMenuItem*)id
		offStateImage];
}

void NSMenuItem_inst_SetOffStateImage(void *id, void* value) {
	[(NSMenuItem*)id
		setOffStateImage: value];
}

void* NSMenuItem_inst_MixedStateImage(void *id) {
	return [(NSMenuItem*)id
		mixedStateImage];
}

void NSMenuItem_inst_SetMixedStateImage(void *id, void* value) {
	[(NSMenuItem*)id
		setMixedStateImage: value];
}

void* NSMenuItem_inst_Submenu(void *id) {
	return [(NSMenuItem*)id
		submenu];
}

void NSMenuItem_inst_SetSubmenu(void *id, void* value) {
	[(NSMenuItem*)id
		setSubmenu: value];
}

BOOL NSMenuItem_inst_HasSubmenu(void *id) {
	return [(NSMenuItem*)id
		hasSubmenu];
}

void* NSMenuItem_inst_ParentItem(void *id) {
	return [(NSMenuItem*)id
		parentItem];
}

BOOL NSMenuItem_inst_IsSeparatorItem(void *id) {
	return [(NSMenuItem*)id
		isSeparatorItem];
}

void* NSMenuItem_inst_Menu(void *id) {
	return [(NSMenuItem*)id
		menu];
}

void NSMenuItem_inst_SetMenu(void *id, void* value) {
	[(NSMenuItem*)id
		setMenu: value];
}

void* NSMenuItem_inst_KeyEquivalent(void *id) {
	return [(NSMenuItem*)id
		keyEquivalent];
}

void NSMenuItem_inst_SetKeyEquivalent(void *id, void* value) {
	[(NSMenuItem*)id
		setKeyEquivalent: value];
}

void* NSMenuItem_inst_UserKeyEquivalent(void *id) {
	return [(NSMenuItem*)id
		userKeyEquivalent];
}

BOOL NSMenuItem_inst_IsAlternate(void *id) {
	return [(NSMenuItem*)id
		isAlternate];
}

void NSMenuItem_inst_SetAlternate(void *id, BOOL value) {
	[(NSMenuItem*)id
		setAlternate: value];
}

long NSMenuItem_inst_IndentationLevel(void *id) {
	return [(NSMenuItem*)id
		indentationLevel];
}

void NSMenuItem_inst_SetIndentationLevel(void *id, long value) {
	[(NSMenuItem*)id
		setIndentationLevel: value];
}

void* NSMenuItem_inst_ToolTip(void *id) {
	return [(NSMenuItem*)id
		toolTip];
}

void NSMenuItem_inst_SetToolTip(void *id, void* value) {
	[(NSMenuItem*)id
		setToolTip: value];
}

void* NSMenuItem_inst_RepresentedObject(void *id) {
	return [(NSMenuItem*)id
		representedObject];
}

void NSMenuItem_inst_SetRepresentedObject(void *id, void* value) {
	[(NSMenuItem*)id
		setRepresentedObject: value];
}

void* NSMenuItem_inst_View(void *id) {
	return [(NSMenuItem*)id
		view];
}

void NSMenuItem_inst_SetView(void *id, void* value) {
	[(NSMenuItem*)id
		setView: value];
}

BOOL NSMenuItem_inst_IsHighlighted(void *id) {
	return [(NSMenuItem*)id
		isHighlighted];
}

BOOL NSMenuItem_inst_AllowsAutomaticKeyEquivalentLocalization(void *id) {
	return [(NSMenuItem*)id
		allowsAutomaticKeyEquivalentLocalization];
}

void NSMenuItem_inst_SetAllowsAutomaticKeyEquivalentLocalization(void *id, BOOL value) {
	[(NSMenuItem*)id
		setAllowsAutomaticKeyEquivalentLocalization: value];
}

BOOL NSMenuItem_inst_AllowsAutomaticKeyEquivalentMirroring(void *id) {
	return [(NSMenuItem*)id
		allowsAutomaticKeyEquivalentMirroring];
}

void NSMenuItem_inst_SetAllowsAutomaticKeyEquivalentMirroring(void *id, BOOL value) {
	[(NSMenuItem*)id
		setAllowsAutomaticKeyEquivalentMirroring: value];
}

BOOL NSMenuItem_inst_AllowsKeyEquivalentWhenHidden(void *id) {
	return [(NSMenuItem*)id
		allowsKeyEquivalentWhenHidden];
}

void NSMenuItem_inst_SetAllowsKeyEquivalentWhenHidden(void *id, BOOL value) {
	[(NSMenuItem*)id
		setAllowsKeyEquivalentWhenHidden: value];
}

BOOL NSRunningApplication_inst_ForceTerminate(void *id) {
	return [(NSRunningApplication*)id
		forceTerminate];
}

BOOL NSRunningApplication_inst_Hide(void *id) {
	return [(NSRunningApplication*)id
		hide];
}

BOOL NSRunningApplication_inst_Terminate(void *id) {
	return [(NSRunningApplication*)id
		terminate];
}

BOOL NSRunningApplication_inst_Unhide(void *id) {
	return [(NSRunningApplication*)id
		unhide];
}

void* NSRunningApplication_inst_Init(void *id) {
	return [(NSRunningApplication*)id
		init];
}

BOOL NSRunningApplication_inst_IsActive(void *id) {
	return [(NSRunningApplication*)id
		isActive];
}

long NSRunningApplication_inst_ActivationPolicy(void *id) {
	return [(NSRunningApplication*)id
		activationPolicy];
}

BOOL NSRunningApplication_inst_IsHidden(void *id) {
	return [(NSRunningApplication*)id
		isHidden];
}

void* NSRunningApplication_inst_LocalizedName(void *id) {
	return [(NSRunningApplication*)id
		localizedName];
}

void* NSRunningApplication_inst_Icon(void *id) {
	return [(NSRunningApplication*)id
		icon];
}

void* NSRunningApplication_inst_BundleIdentifier(void *id) {
	return [(NSRunningApplication*)id
		bundleIdentifier];
}

void* NSRunningApplication_inst_BundleURL(void *id) {
	return [(NSRunningApplication*)id
		bundleURL];
}

long NSRunningApplication_inst_ExecutableArchitecture(void *id) {
	return [(NSRunningApplication*)id
		executableArchitecture];
}

void* NSRunningApplication_inst_ExecutableURL(void *id) {
	return [(NSRunningApplication*)id
		executableURL];
}

BOOL NSRunningApplication_inst_IsFinishedLaunching(void *id) {
	return [(NSRunningApplication*)id
		isFinishedLaunching];
}

BOOL NSRunningApplication_inst_OwnsMenuBar(void *id) {
	return [(NSRunningApplication*)id
		ownsMenuBar];
}

BOOL NSRunningApplication_inst_IsTerminated(void *id) {
	return [(NSRunningApplication*)id
		isTerminated];
}

NSRect NSScreen_inst_ConvertRectFromBacking(void *id, NSRect rect) {
	return [(NSScreen*)id
		convertRectFromBacking: rect];
}

NSRect NSScreen_inst_ConvertRectToBacking(void *id, NSRect rect) {
	return [(NSScreen*)id
		convertRectToBacking: rect];
}

void* NSScreen_inst_Init(void *id) {
	return [(NSScreen*)id
		init];
}

NSRect NSScreen_inst_Frame(void *id) {
	return [(NSScreen*)id
		frame];
}

void* NSScreen_inst_DeviceDescription(void *id) {
	return [(NSScreen*)id
		deviceDescription];
}

NSRect NSScreen_inst_VisibleFrame(void *id) {
	return [(NSScreen*)id
		visibleFrame];
}

double NSScreen_inst_BackingScaleFactor(void *id) {
	return [(NSScreen*)id
		backingScaleFactor];
}

double NSScreen_inst_MaximumPotentialExtendedDynamicRangeColorComponentValue(void *id) {
	return [(NSScreen*)id
		maximumPotentialExtendedDynamicRangeColorComponentValue];
}

double NSScreen_inst_MaximumExtendedDynamicRangeColorComponentValue(void *id) {
	return [(NSScreen*)id
		maximumExtendedDynamicRangeColorComponentValue];
}

double NSScreen_inst_MaximumReferenceExtendedDynamicRangeColorComponentValue(void *id) {
	return [(NSScreen*)id
		maximumReferenceExtendedDynamicRangeColorComponentValue];
}

void* NSScreen_inst_LocalizedName(void *id) {
	return [(NSScreen*)id
		localizedName];
}

long NSScreen_inst_MaximumFramesPerSecond(void *id) {
	return [(NSScreen*)id
		maximumFramesPerSecond];
}

void NSStatusBar_inst_RemoveStatusItem(void *id, void* item) {
	[(NSStatusBar*)id
		removeStatusItem: item];
}

void* NSStatusBar_inst_StatusItemWithLength(void *id, double length) {
	return [(NSStatusBar*)id
		statusItemWithLength: length];
}

void* NSStatusBar_inst_Init(void *id) {
	return [(NSStatusBar*)id
		init];
}

BOOL NSStatusBar_inst_IsVertical(void *id) {
	return [(NSStatusBar*)id
		isVertical];
}

double NSStatusBar_inst_Thickness(void *id) {
	return [(NSStatusBar*)id
		thickness];
}

void* NSStatusBarButton_inst_Init(void *id) {
	return [(NSStatusBarButton*)id
		init];
}

BOOL NSStatusBarButton_inst_AppearsDisabled(void *id) {
	return [(NSStatusBarButton*)id
		appearsDisabled];
}

void NSStatusBarButton_inst_SetAppearsDisabled(void *id, BOOL value) {
	[(NSStatusBarButton*)id
		setAppearsDisabled: value];
}

void* NSStatusItem_inst_Init(void *id) {
	return [(NSStatusItem*)id
		init];
}

void* NSStatusItem_inst_StatusBar(void *id) {
	return [(NSStatusItem*)id
		statusBar];
}

void* NSStatusItem_inst_Button(void *id) {
	return [(NSStatusItem*)id
		button];
}

void* NSStatusItem_inst_Menu(void *id) {
	return [(NSStatusItem*)id
		menu];
}

void NSStatusItem_inst_SetMenu(void *id, void* value) {
	[(NSStatusItem*)id
		setMenu: value];
}

BOOL NSStatusItem_inst_IsVisible(void *id) {
	return [(NSStatusItem*)id
		isVisible];
}

void NSStatusItem_inst_SetVisible(void *id, BOOL value) {
	[(NSStatusItem*)id
		setVisible: value];
}

double NSStatusItem_inst_Length(void *id) {
	return [(NSStatusItem*)id
		length];
}

void NSStatusItem_inst_SetLength(void *id, double value) {
	[(NSStatusItem*)id
		setLength: value];
}

void NSText_inst_AlignCenter(void *id, void* sender) {
	[(NSText*)id
		alignCenter: sender];
}

void NSText_inst_AlignLeft(void *id, void* sender) {
	[(NSText*)id
		alignLeft: sender];
}

void NSText_inst_AlignRight(void *id, void* sender) {
	[(NSText*)id
		alignRight: sender];
}

void NSText_inst_ChangeFont(void *id, void* sender) {
	[(NSText*)id
		changeFont: sender];
}

void NSText_inst_CheckSpelling(void *id, void* sender) {
	[(NSText*)id
		checkSpelling: sender];
}

void NSText_inst_Copy(void *id, void* sender) {
	[(NSText*)id
		copy: sender];
}

void NSText_inst_CopyFont(void *id, void* sender) {
	[(NSText*)id
		copyFont: sender];
}

void NSText_inst_CopyRuler(void *id, void* sender) {
	[(NSText*)id
		copyRuler: sender];
}

void NSText_inst_Cut(void *id, void* sender) {
	[(NSText*)id
		cut: sender];
}

void NSText_inst_Delete(void *id, void* sender) {
	[(NSText*)id
		delete: sender];
}

void* NSText_inst_InitWithFrame(void *id, NSRect frameRect) {
	return [(NSText*)id
		initWithFrame: frameRect];
}

void NSText_inst_Paste(void *id, void* sender) {
	[(NSText*)id
		paste: sender];
}

void NSText_inst_PasteFont(void *id, void* sender) {
	[(NSText*)id
		pasteFont: sender];
}

void NSText_inst_PasteRuler(void *id, void* sender) {
	[(NSText*)id
		pasteRuler: sender];
}

BOOL NSText_inst_ReadRTFDFromFile(void *id, void* path) {
	return [(NSText*)id
		readRTFDFromFile: path];
}

void NSText_inst_SelectAll(void *id, void* sender) {
	[(NSText*)id
		selectAll: sender];
}

void NSText_inst_ShowGuessPanel(void *id, void* sender) {
	[(NSText*)id
		showGuessPanel: sender];
}

void NSText_inst_SizeToFit(void *id) {
	[(NSText*)id
		sizeToFit];
}

void NSText_inst_Subscript(void *id, void* sender) {
	[(NSText*)id
		subscript: sender];
}

void NSText_inst_Superscript(void *id, void* sender) {
	[(NSText*)id
		superscript: sender];
}

void NSText_inst_ToggleRuler(void *id, void* sender) {
	[(NSText*)id
		toggleRuler: sender];
}

void NSText_inst_Underline(void *id, void* sender) {
	[(NSText*)id
		underline: sender];
}

void NSText_inst_Unscript(void *id, void* sender) {
	[(NSText*)id
		unscript: sender];
}

BOOL NSText_inst_WriteRTFDToFileAtomically(void *id, void* path, BOOL flag) {
	return [(NSText*)id
		writeRTFDToFile: path
		atomically: flag];
}

void* NSText_inst_Init(void *id) {
	return [(NSText*)id
		init];
}

void* NSText_inst_String(void *id) {
	return [(NSText*)id
		string];
}

void NSText_inst_SetString(void *id, void* value) {
	[(NSText*)id
		setString: value];
}

void* NSText_inst_BackgroundColor(void *id) {
	return [(NSText*)id
		backgroundColor];
}

void NSText_inst_SetBackgroundColor(void *id, void* value) {
	[(NSText*)id
		setBackgroundColor: value];
}

BOOL NSText_inst_DrawsBackground(void *id) {
	return [(NSText*)id
		drawsBackground];
}

void NSText_inst_SetDrawsBackground(void *id, BOOL value) {
	[(NSText*)id
		setDrawsBackground: value];
}

BOOL NSText_inst_IsEditable(void *id) {
	return [(NSText*)id
		isEditable];
}

void NSText_inst_SetEditable(void *id, BOOL value) {
	[(NSText*)id
		setEditable: value];
}

BOOL NSText_inst_IsSelectable(void *id) {
	return [(NSText*)id
		isSelectable];
}

void NSText_inst_SetSelectable(void *id, BOOL value) {
	[(NSText*)id
		setSelectable: value];
}

BOOL NSText_inst_IsFieldEditor(void *id) {
	return [(NSText*)id
		isFieldEditor];
}

void NSText_inst_SetFieldEditor(void *id, BOOL value) {
	[(NSText*)id
		setFieldEditor: value];
}

BOOL NSText_inst_IsRichText(void *id) {
	return [(NSText*)id
		isRichText];
}

void NSText_inst_SetRichText(void *id, BOOL value) {
	[(NSText*)id
		setRichText: value];
}

BOOL NSText_inst_ImportsGraphics(void *id) {
	return [(NSText*)id
		importsGraphics];
}

void NSText_inst_SetImportsGraphics(void *id, BOOL value) {
	[(NSText*)id
		setImportsGraphics: value];
}

BOOL NSText_inst_UsesFontPanel(void *id) {
	return [(NSText*)id
		usesFontPanel];
}

void NSText_inst_SetUsesFontPanel(void *id, BOOL value) {
	[(NSText*)id
		setUsesFontPanel: value];
}

BOOL NSText_inst_IsRulerVisible(void *id) {
	return [(NSText*)id
		isRulerVisible];
}

void* NSText_inst_Font(void *id) {
	return [(NSText*)id
		font];
}

void NSText_inst_SetFont(void *id, void* value) {
	[(NSText*)id
		setFont: value];
}

void* NSText_inst_TextColor(void *id) {
	return [(NSText*)id
		textColor];
}

void NSText_inst_SetTextColor(void *id, void* value) {
	[(NSText*)id
		setTextColor: value];
}

NSSize NSText_inst_MaxSize(void *id) {
	return [(NSText*)id
		maxSize];
}

void NSText_inst_SetMaxSize(void *id, NSSize value) {
	[(NSText*)id
		setMaxSize: value];
}

NSSize NSText_inst_MinSize(void *id) {
	return [(NSText*)id
		minSize];
}

void NSText_inst_SetMinSize(void *id, NSSize value) {
	[(NSText*)id
		setMinSize: value];
}

BOOL NSText_inst_IsVerticallyResizable(void *id) {
	return [(NSText*)id
		isVerticallyResizable];
}

void NSText_inst_SetVerticallyResizable(void *id, BOOL value) {
	[(NSText*)id
		setVerticallyResizable: value];
}

BOOL NSText_inst_IsHorizontallyResizable(void *id) {
	return [(NSText*)id
		isHorizontallyResizable];
}

void NSText_inst_SetHorizontallyResizable(void *id, BOOL value) {
	[(NSText*)id
		setHorizontallyResizable: value];
}

void* NSText_inst_Delegate(void *id) {
	return [(NSText*)id
		delegate];
}

void NSText_inst_SetDelegate(void *id, void* value) {
	[(NSText*)id
		setDelegate: value];
}

void NSTextField_inst_SelectText(void *id, void* sender) {
	[(NSTextField*)id
		selectText: sender];
}

BOOL NSTextField_inst_TextShouldBeginEditing(void *id, void* textObject) {
	return [(NSTextField*)id
		textShouldBeginEditing: textObject];
}

BOOL NSTextField_inst_TextShouldEndEditing(void *id, void* textObject) {
	return [(NSTextField*)id
		textShouldEndEditing: textObject];
}

void* NSTextField_inst_Init(void *id) {
	return [(NSTextField*)id
		init];
}

BOOL NSTextField_inst_IsSelectable(void *id) {
	return [(NSTextField*)id
		isSelectable];
}

void NSTextField_inst_SetSelectable(void *id, BOOL value) {
	[(NSTextField*)id
		setSelectable: value];
}

BOOL NSTextField_inst_IsEditable(void *id) {
	return [(NSTextField*)id
		isEditable];
}

void NSTextField_inst_SetEditable(void *id, BOOL value) {
	[(NSTextField*)id
		setEditable: value];
}

BOOL NSTextField_inst_AllowsEditingTextAttributes(void *id) {
	return [(NSTextField*)id
		allowsEditingTextAttributes];
}

void NSTextField_inst_SetAllowsEditingTextAttributes(void *id, BOOL value) {
	[(NSTextField*)id
		setAllowsEditingTextAttributes: value];
}

BOOL NSTextField_inst_ImportsGraphics(void *id) {
	return [(NSTextField*)id
		importsGraphics];
}

void NSTextField_inst_SetImportsGraphics(void *id, BOOL value) {
	[(NSTextField*)id
		setImportsGraphics: value];
}

void* NSTextField_inst_PlaceholderString(void *id) {
	return [(NSTextField*)id
		placeholderString];
}

void NSTextField_inst_SetPlaceholderString(void *id, void* value) {
	[(NSTextField*)id
		setPlaceholderString: value];
}

void* NSTextField_inst_PlaceholderAttributedString(void *id) {
	return [(NSTextField*)id
		placeholderAttributedString];
}

void NSTextField_inst_SetPlaceholderAttributedString(void *id, void* value) {
	[(NSTextField*)id
		setPlaceholderAttributedString: value];
}

BOOL NSTextField_inst_AllowsDefaultTighteningForTruncation(void *id) {
	return [(NSTextField*)id
		allowsDefaultTighteningForTruncation];
}

void NSTextField_inst_SetAllowsDefaultTighteningForTruncation(void *id, BOOL value) {
	[(NSTextField*)id
		setAllowsDefaultTighteningForTruncation: value];
}

long NSTextField_inst_MaximumNumberOfLines(void *id) {
	return [(NSTextField*)id
		maximumNumberOfLines];
}

void NSTextField_inst_SetMaximumNumberOfLines(void *id, long value) {
	[(NSTextField*)id
		setMaximumNumberOfLines: value];
}

double NSTextField_inst_PreferredMaxLayoutWidth(void *id) {
	return [(NSTextField*)id
		preferredMaxLayoutWidth];
}

void NSTextField_inst_SetPreferredMaxLayoutWidth(void *id, double value) {
	[(NSTextField*)id
		setPreferredMaxLayoutWidth: value];
}

void* NSTextField_inst_TextColor(void *id) {
	return [(NSTextField*)id
		textColor];
}

void NSTextField_inst_SetTextColor(void *id, void* value) {
	[(NSTextField*)id
		setTextColor: value];
}

void* NSTextField_inst_BackgroundColor(void *id) {
	return [(NSTextField*)id
		backgroundColor];
}

void NSTextField_inst_SetBackgroundColor(void *id, void* value) {
	[(NSTextField*)id
		setBackgroundColor: value];
}

BOOL NSTextField_inst_DrawsBackground(void *id) {
	return [(NSTextField*)id
		drawsBackground];
}

void NSTextField_inst_SetDrawsBackground(void *id, BOOL value) {
	[(NSTextField*)id
		setDrawsBackground: value];
}

BOOL NSTextField_inst_IsBezeled(void *id) {
	return [(NSTextField*)id
		isBezeled];
}

void NSTextField_inst_SetBezeled(void *id, BOOL value) {
	[(NSTextField*)id
		setBezeled: value];
}

BOOL NSTextField_inst_IsBordered(void *id) {
	return [(NSTextField*)id
		isBordered];
}

void NSTextField_inst_SetBordered(void *id, BOOL value) {
	[(NSTextField*)id
		setBordered: value];
}

BOOL NSTextField_inst_AcceptsFirstResponder(void *id) {
	return [(NSTextField*)id
		acceptsFirstResponder];
}

BOOL NSTextField_inst_AllowsCharacterPickerTouchBarItem(void *id) {
	return [(NSTextField*)id
		allowsCharacterPickerTouchBarItem];
}

void NSTextField_inst_SetAllowsCharacterPickerTouchBarItem(void *id, BOOL value) {
	[(NSTextField*)id
		setAllowsCharacterPickerTouchBarItem: value];
}

BOOL NSTextField_inst_IsAutomaticTextCompletionEnabled(void *id) {
	return [(NSTextField*)id
		isAutomaticTextCompletionEnabled];
}

void NSTextField_inst_SetAutomaticTextCompletionEnabled(void *id, BOOL value) {
	[(NSTextField*)id
		setAutomaticTextCompletionEnabled: value];
}

void* NSTextField_inst_Delegate(void *id) {
	return [(NSTextField*)id
		delegate];
}

void NSTextField_inst_SetDelegate(void *id, void* value) {
	[(NSTextField*)id
		setDelegate: value];
}

void* NSTextContainer_inst_InitWithSize(void *id, NSSize size) {
	return [(NSTextContainer*)id
		initWithSize: size];
}

void NSTextContainer_inst_ReplaceLayoutManager(void *id, void* newLayoutManager) {
	[(NSTextContainer*)id
		replaceLayoutManager: newLayoutManager];
}

void* NSTextContainer_inst_Init(void *id) {
	return [(NSTextContainer*)id
		init];
}

void* NSTextContainer_inst_LayoutManager(void *id) {
	return [(NSTextContainer*)id
		layoutManager];
}

void NSTextContainer_inst_SetLayoutManager(void *id, void* value) {
	[(NSTextContainer*)id
		setLayoutManager: value];
}

void* NSTextContainer_inst_TextView(void *id) {
	return [(NSTextContainer*)id
		textView];
}

void NSTextContainer_inst_SetTextView(void *id, void* value) {
	[(NSTextContainer*)id
		setTextView: value];
}

NSSize NSTextContainer_inst_Size(void *id) {
	return [(NSTextContainer*)id
		size];
}

void NSTextContainer_inst_SetSize(void *id, NSSize value) {
	[(NSTextContainer*)id
		setSize: value];
}

void* NSTextContainer_inst_ExclusionPaths(void *id) {
	return [(NSTextContainer*)id
		exclusionPaths];
}

void NSTextContainer_inst_SetExclusionPaths(void *id, void* value) {
	[(NSTextContainer*)id
		setExclusionPaths: value];
}

BOOL NSTextContainer_inst_WidthTracksTextView(void *id) {
	return [(NSTextContainer*)id
		widthTracksTextView];
}

void NSTextContainer_inst_SetWidthTracksTextView(void *id, BOOL value) {
	[(NSTextContainer*)id
		setWidthTracksTextView: value];
}

BOOL NSTextContainer_inst_HeightTracksTextView(void *id) {
	return [(NSTextContainer*)id
		heightTracksTextView];
}

void NSTextContainer_inst_SetHeightTracksTextView(void *id, BOOL value) {
	[(NSTextContainer*)id
		setHeightTracksTextView: value];
}

unsigned long NSTextContainer_inst_MaximumNumberOfLines(void *id) {
	return [(NSTextContainer*)id
		maximumNumberOfLines];
}

void NSTextContainer_inst_SetMaximumNumberOfLines(void *id, unsigned long value) {
	[(NSTextContainer*)id
		setMaximumNumberOfLines: value];
}

double NSTextContainer_inst_LineFragmentPadding(void *id) {
	return [(NSTextContainer*)id
		lineFragmentPadding];
}

void NSTextContainer_inst_SetLineFragmentPadding(void *id, double value) {
	[(NSTextContainer*)id
		setLineFragmentPadding: value];
}

BOOL NSTextContainer_inst_IsSimpleRectangularTextContainer(void *id) {
	return [(NSTextContainer*)id
		isSimpleRectangularTextContainer];
}

void NSViewController_inst_AddChildViewController(void *id, void* childViewController) {
	[(NSViewController*)id
		addChildViewController: childViewController];
}

BOOL NSViewController_inst_CommitEditing(void *id) {
	return [(NSViewController*)id
		commitEditing];
}

void NSViewController_inst_CommitEditingWithDelegateDidCommitSelectorContextInfo(void *id, void* delegate, void* didCommitSelector, void* contextInfo) {
	[(NSViewController*)id
		commitEditingWithDelegate: delegate
		didCommitSelector: didCommitSelector
		contextInfo: contextInfo];
}

void NSViewController_inst_DiscardEditing(void *id) {
	[(NSViewController*)id
		discardEditing];
}

void NSViewController_inst_DismissController(void *id, void* sender) {
	[(NSViewController*)id
		dismissController: sender];
}

void NSViewController_inst_DismissViewController(void *id, void* viewController) {
	[(NSViewController*)id
		dismissViewController: viewController];
}

void NSViewController_inst_InsertChildViewControllerAtIndex(void *id, void* childViewController, long index) {
	[(NSViewController*)id
		insertChildViewController: childViewController
		atIndex: index];
}

void NSViewController_inst_LoadView(void *id) {
	[(NSViewController*)id
		loadView];
}

void NSViewController_inst_PreferredContentSizeDidChangeForViewController(void *id, void* viewController) {
	[(NSViewController*)id
		preferredContentSizeDidChangeForViewController: viewController];
}

void NSViewController_inst_PresentViewControllerAnimator(void *id, void* viewController, void* animator) {
	[(NSViewController*)id
		presentViewController: viewController
		animator: animator];
}

void NSViewController_inst_PresentViewControllerAsModalWindow(void *id, void* viewController) {
	[(NSViewController*)id
		presentViewControllerAsModalWindow: viewController];
}

void NSViewController_inst_PresentViewControllerAsSheet(void *id, void* viewController) {
	[(NSViewController*)id
		presentViewControllerAsSheet: viewController];
}

void NSViewController_inst_RemoveChildViewControllerAtIndex(void *id, long index) {
	[(NSViewController*)id
		removeChildViewControllerAtIndex: index];
}

void NSViewController_inst_RemoveFromParentViewController(void *id) {
	[(NSViewController*)id
		removeFromParentViewController];
}

void NSViewController_inst_UpdateViewConstraints(void *id) {
	[(NSViewController*)id
		updateViewConstraints];
}

void NSViewController_inst_ViewDidAppear(void *id) {
	[(NSViewController*)id
		viewDidAppear];
}

void NSViewController_inst_ViewDidDisappear(void *id) {
	[(NSViewController*)id
		viewDidDisappear];
}

void NSViewController_inst_ViewDidLayout(void *id) {
	[(NSViewController*)id
		viewDidLayout];
}

void NSViewController_inst_ViewDidLoad(void *id) {
	[(NSViewController*)id
		viewDidLoad];
}

void NSViewController_inst_ViewWillAppear(void *id) {
	[(NSViewController*)id
		viewWillAppear];
}

void NSViewController_inst_ViewWillDisappear(void *id) {
	[(NSViewController*)id
		viewWillDisappear];
}

void NSViewController_inst_ViewWillLayout(void *id) {
	[(NSViewController*)id
		viewWillLayout];
}

void NSViewController_inst_ViewWillTransitionToSize(void *id, NSSize newSize) {
	[(NSViewController*)id
		viewWillTransitionToSize: newSize];
}

void* NSViewController_inst_Init(void *id) {
	return [(NSViewController*)id
		init];
}

void* NSViewController_inst_RepresentedObject(void *id) {
	return [(NSViewController*)id
		representedObject];
}

void NSViewController_inst_SetRepresentedObject(void *id, void* value) {
	[(NSViewController*)id
		setRepresentedObject: value];
}

void* NSViewController_inst_NibBundle(void *id) {
	return [(NSViewController*)id
		nibBundle];
}

void* NSViewController_inst_View(void *id) {
	return [(NSViewController*)id
		view];
}

void NSViewController_inst_SetView(void *id, void* value) {
	[(NSViewController*)id
		setView: value];
}

void* NSViewController_inst_Title(void *id) {
	return [(NSViewController*)id
		title];
}

void NSViewController_inst_SetTitle(void *id, void* value) {
	[(NSViewController*)id
		setTitle: value];
}

BOOL NSViewController_inst_IsViewLoaded(void *id) {
	return [(NSViewController*)id
		isViewLoaded];
}

NSSize NSViewController_inst_PreferredContentSize(void *id) {
	return [(NSViewController*)id
		preferredContentSize];
}

void NSViewController_inst_SetPreferredContentSize(void *id, NSSize value) {
	[(NSViewController*)id
		setPreferredContentSize: value];
}

void* NSViewController_inst_ChildViewControllers(void *id) {
	return [(NSViewController*)id
		childViewControllers];
}

void NSViewController_inst_SetChildViewControllers(void *id, void* value) {
	[(NSViewController*)id
		setChildViewControllers: value];
}

void* NSViewController_inst_ParentViewController(void *id) {
	return [(NSViewController*)id
		parentViewController];
}

void* NSViewController_inst_PresentedViewControllers(void *id) {
	return [(NSViewController*)id
		presentedViewControllers];
}

void* NSViewController_inst_PresentingViewController(void *id) {
	return [(NSViewController*)id
		presentingViewController];
}

NSPoint NSViewController_inst_PreferredScreenOrigin(void *id) {
	return [(NSViewController*)id
		preferredScreenOrigin];
}

void NSViewController_inst_SetPreferredScreenOrigin(void *id, NSPoint value) {
	[(NSViewController*)id
		setPreferredScreenOrigin: value];
}

NSSize NSViewController_inst_PreferredMaximumSize(void *id) {
	return [(NSViewController*)id
		preferredMaximumSize];
}

NSSize NSViewController_inst_PreferredMinimumSize(void *id) {
	return [(NSViewController*)id
		preferredMinimumSize];
}

void* NSViewController_inst_SourceItemView(void *id) {
	return [(NSViewController*)id
		sourceItemView];
}

void NSViewController_inst_SetSourceItemView(void *id, void* value) {
	[(NSViewController*)id
		setSourceItemView: value];
}

void NSVisualEffectView_inst_ViewDidMoveToWindow(void *id) {
	[(NSVisualEffectView*)id
		viewDidMoveToWindow];
}

void NSVisualEffectView_inst_ViewWillMoveToWindow(void *id, void* newWindow) {
	[(NSVisualEffectView*)id
		viewWillMoveToWindow: newWindow];
}

void* NSVisualEffectView_inst_Init(void *id) {
	return [(NSVisualEffectView*)id
		init];
}

BOOL NSVisualEffectView_inst_IsEmphasized(void *id) {
	return [(NSVisualEffectView*)id
		isEmphasized];
}

void NSVisualEffectView_inst_SetEmphasized(void *id, BOOL value) {
	[(NSVisualEffectView*)id
		setEmphasized: value];
}

void* NSVisualEffectView_inst_MaskImage(void *id) {
	return [(NSVisualEffectView*)id
		maskImage];
}

void NSVisualEffectView_inst_SetMaskImage(void *id, void* value) {
	[(NSVisualEffectView*)id
		setMaskImage: value];
}

void NSWindow_inst_AddChildWindowOrdered(void *id, void* childWin, unsigned long place) {
	[(NSWindow*)id
		addChildWindow: childWin
		ordered: place];
}

void NSWindow_inst_AddTabbedWindowOrdered(void *id, void* window, unsigned long ordered) {
	[(NSWindow*)id
		addTabbedWindow: window
		ordered: ordered];
}

void NSWindow_inst_BecomeKeyWindow(void *id) {
	[(NSWindow*)id
		becomeKeyWindow];
}

void NSWindow_inst_BecomeMainWindow(void *id) {
	[(NSWindow*)id
		becomeMainWindow];
}

NSPoint NSWindow_inst_CascadeTopLeftFromPoint(void *id, NSPoint topLeftPoint) {
	return [(NSWindow*)id
		cascadeTopLeftFromPoint: topLeftPoint];
}

void NSWindow_inst_Center(void *id) {
	[(NSWindow*)id
		center];
}

void NSWindow_inst_Close(void *id) {
	[(NSWindow*)id
		close];
}

NSRect NSWindow_inst_ConstrainFrameRectToScreen(void *id, NSRect frameRect, void* screen) {
	return [(NSWindow*)id
		constrainFrameRect: frameRect
		toScreen: screen];
}

NSRect NSWindow_inst_ContentRectForFrameRect(void *id, NSRect frameRect) {
	return [(NSWindow*)id
		contentRectForFrameRect: frameRect];
}

NSPoint NSWindow_inst_ConvertPointFromBacking(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointFromBacking: point];
}

NSPoint NSWindow_inst_ConvertPointFromScreen(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointFromScreen: point];
}

NSPoint NSWindow_inst_ConvertPointToBacking(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointToBacking: point];
}

NSPoint NSWindow_inst_ConvertPointToScreen(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointToScreen: point];
}

NSRect NSWindow_inst_ConvertRectFromBacking(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectFromBacking: rect];
}

NSRect NSWindow_inst_ConvertRectFromScreen(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectFromScreen: rect];
}

NSRect NSWindow_inst_ConvertRectToBacking(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectToBacking: rect];
}

NSRect NSWindow_inst_ConvertRectToScreen(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectToScreen: rect];
}

void* NSWindow_inst_DataWithEPSInsideRect(void *id, NSRect rect) {
	return [(NSWindow*)id
		dataWithEPSInsideRect: rect];
}

void* NSWindow_inst_DataWithPDFInsideRect(void *id, NSRect rect) {
	return [(NSWindow*)id
		dataWithPDFInsideRect: rect];
}

void NSWindow_inst_Deminiaturize(void *id, void* sender) {
	[(NSWindow*)id
		deminiaturize: sender];
}

void NSWindow_inst_DisableCursorRects(void *id) {
	[(NSWindow*)id
		disableCursorRects];
}

void NSWindow_inst_DisableKeyEquivalentForDefaultButtonCell(void *id) {
	[(NSWindow*)id
		disableKeyEquivalentForDefaultButtonCell];
}

void NSWindow_inst_DisableScreenUpdatesUntilFlush(void *id) {
	[(NSWindow*)id
		disableScreenUpdatesUntilFlush];
}

void NSWindow_inst_DisableSnapshotRestoration(void *id) {
	[(NSWindow*)id
		disableSnapshotRestoration];
}

void NSWindow_inst_DiscardCursorRects(void *id) {
	[(NSWindow*)id
		discardCursorRects];
}

void NSWindow_inst_Display(void *id) {
	[(NSWindow*)id
		display];
}

void NSWindow_inst_DisplayIfNeeded(void *id) {
	[(NSWindow*)id
		displayIfNeeded];
}

void NSWindow_inst_DragImageAtOffsetEventPasteboardSourceSlideBack(void *id, void* image, NSPoint baseLocation, NSSize initialOffset, void* event, void* pboard, void* sourceObj, BOOL slideFlag) {
	[(NSWindow*)id
		dragImage: image
		at: baseLocation
		offset: initialOffset
		event: event
		pasteboard: pboard
		source: sourceObj
		slideBack: slideFlag];
}

void NSWindow_inst_EnableCursorRects(void *id) {
	[(NSWindow*)id
		enableCursorRects];
}

void NSWindow_inst_EnableKeyEquivalentForDefaultButtonCell(void *id) {
	[(NSWindow*)id
		enableKeyEquivalentForDefaultButtonCell];
}

void NSWindow_inst_EnableSnapshotRestoration(void *id) {
	[(NSWindow*)id
		enableSnapshotRestoration];
}

void NSWindow_inst_EndEditingFor(void *id, void* object) {
	[(NSWindow*)id
		endEditingFor: object];
}

void NSWindow_inst_EndSheet(void *id, void* sheetWindow) {
	[(NSWindow*)id
		endSheet: sheetWindow];
}

void* NSWindow_inst_FieldEditorForObject(void *id, BOOL createFlag, void* object) {
	return [(NSWindow*)id
		fieldEditor: createFlag
		forObject: object];
}

NSRect NSWindow_inst_FrameRectForContentRect(void *id, NSRect contentRect) {
	return [(NSWindow*)id
		frameRectForContentRect: contentRect];
}

void* NSWindow_inst_InitWithContentRectStyleMaskBackingDefer(void *id, NSRect contentRect, unsigned long style, unsigned long backingStoreType, BOOL flag) {
	return [(NSWindow*)id
		initWithContentRect: contentRect
		styleMask: style
		backing: backingStoreType
		defer: flag];
}

void* NSWindow_inst_InitWithContentRectStyleMaskBackingDeferScreen(void *id, NSRect contentRect, unsigned long style, unsigned long backingStoreType, BOOL flag, void* screen) {
	return [(NSWindow*)id
		initWithContentRect: contentRect
		styleMask: style
		backing: backingStoreType
		defer: flag
		screen: screen];
}

void NSWindow_inst_InvalidateCursorRectsForView(void *id, void* view) {
	[(NSWindow*)id
		invalidateCursorRectsForView: view];
}

void NSWindow_inst_InvalidateShadow(void *id) {
	[(NSWindow*)id
		invalidateShadow];
}

void NSWindow_inst_LayoutIfNeeded(void *id) {
	[(NSWindow*)id
		layoutIfNeeded];
}

void NSWindow_inst_MakeKeyAndOrderFront(void *id, void* sender) {
	[(NSWindow*)id
		makeKeyAndOrderFront: sender];
}

void NSWindow_inst_MakeKeyWindow(void *id) {
	[(NSWindow*)id
		makeKeyWindow];
}

void NSWindow_inst_MakeMainWindow(void *id) {
	[(NSWindow*)id
		makeMainWindow];
}

void NSWindow_inst_MergeAllWindows(void *id, void* sender) {
	[(NSWindow*)id
		mergeAllWindows: sender];
}

void NSWindow_inst_Miniaturize(void *id, void* sender) {
	[(NSWindow*)id
		miniaturize: sender];
}

void NSWindow_inst_MoveTabToNewWindow(void *id, void* sender) {
	[(NSWindow*)id
		moveTabToNewWindow: sender];
}

void NSWindow_inst_OrderBack(void *id, void* sender) {
	[(NSWindow*)id
		orderBack: sender];
}

void NSWindow_inst_OrderFront(void *id, void* sender) {
	[(NSWindow*)id
		orderFront: sender];
}

void NSWindow_inst_OrderFrontRegardless(void *id) {
	[(NSWindow*)id
		orderFrontRegardless];
}

void NSWindow_inst_OrderOut(void *id, void* sender) {
	[(NSWindow*)id
		orderOut: sender];
}

void NSWindow_inst_OrderWindowRelativeTo(void *id, unsigned long place, long otherWin) {
	[(NSWindow*)id
		orderWindow: place
		relativeTo: otherWin];
}

void NSWindow_inst_PerformClose(void *id, void* sender) {
	[(NSWindow*)id
		performClose: sender];
}

void NSWindow_inst_PerformMiniaturize(void *id, void* sender) {
	[(NSWindow*)id
		performMiniaturize: sender];
}

void NSWindow_inst_PerformWindowDragWithEvent(void *id, void* event) {
	[(NSWindow*)id
		performWindowDragWithEvent: event];
}

void NSWindow_inst_PerformZoom(void *id, void* sender) {
	[(NSWindow*)id
		performZoom: sender];
}

void NSWindow_inst_PostEventAtStart(void *id, void* event, BOOL flag) {
	[(NSWindow*)id
		postEvent: event
		atStart: flag];
}

void NSWindow_inst_Print(void *id, void* sender) {
	[(NSWindow*)id
		print: sender];
}

void NSWindow_inst_RecalculateKeyViewLoop(void *id) {
	[(NSWindow*)id
		recalculateKeyViewLoop];
}

void NSWindow_inst_RegisterForDraggedTypes(void *id, void* newTypes) {
	[(NSWindow*)id
		registerForDraggedTypes: newTypes];
}

void NSWindow_inst_RemoveChildWindow(void *id, void* childWin) {
	[(NSWindow*)id
		removeChildWindow: childWin];
}

void NSWindow_inst_RemoveTitlebarAccessoryViewControllerAtIndex(void *id, long index) {
	[(NSWindow*)id
		removeTitlebarAccessoryViewControllerAtIndex: index];
}

void NSWindow_inst_ResetCursorRects(void *id) {
	[(NSWindow*)id
		resetCursorRects];
}

void NSWindow_inst_ResignKeyWindow(void *id) {
	[(NSWindow*)id
		resignKeyWindow];
}

void NSWindow_inst_ResignMainWindow(void *id) {
	[(NSWindow*)id
		resignMainWindow];
}

void NSWindow_inst_RunToolbarCustomizationPalette(void *id, void* sender) {
	[(NSWindow*)id
		runToolbarCustomizationPalette: sender];
}

void NSWindow_inst_SelectKeyViewFollowingView(void *id, void* view) {
	[(NSWindow*)id
		selectKeyViewFollowingView: view];
}

void NSWindow_inst_SelectKeyViewPrecedingView(void *id, void* view) {
	[(NSWindow*)id
		selectKeyViewPrecedingView: view];
}

void NSWindow_inst_SelectNextKeyView(void *id, void* sender) {
	[(NSWindow*)id
		selectNextKeyView: sender];
}

void NSWindow_inst_SelectNextTab(void *id, void* sender) {
	[(NSWindow*)id
		selectNextTab: sender];
}

void NSWindow_inst_SelectPreviousKeyView(void *id, void* sender) {
	[(NSWindow*)id
		selectPreviousKeyView: sender];
}

void NSWindow_inst_SelectPreviousTab(void *id, void* sender) {
	[(NSWindow*)id
		selectPreviousTab: sender];
}

void NSWindow_inst_SendEvent(void *id, void* event) {
	[(NSWindow*)id
		sendEvent: event];
}

void NSWindow_inst_SetContentSize(void *id, NSSize size) {
	[(NSWindow*)id
		setContentSize: size];
}

void NSWindow_inst_SetDynamicDepthLimit(void *id, BOOL flag) {
	[(NSWindow*)id
		setDynamicDepthLimit: flag];
}

void NSWindow_inst_SetFrameDisplay(void *id, NSRect frameRect, BOOL flag) {
	[(NSWindow*)id
		setFrame: frameRect
		display: flag];
}

void NSWindow_inst_SetFrameDisplayAnimate(void *id, NSRect frameRect, BOOL displayFlag, BOOL animateFlag) {
	[(NSWindow*)id
		setFrame: frameRect
		display: displayFlag
		animate: animateFlag];
}

void NSWindow_inst_SetFrameOrigin(void *id, NSPoint point) {
	[(NSWindow*)id
		setFrameOrigin: point];
}

void NSWindow_inst_SetFrameTopLeftPoint(void *id, NSPoint point) {
	[(NSWindow*)id
		setFrameTopLeftPoint: point];
}

void NSWindow_inst_SetIsMiniaturized(void *id, BOOL flag) {
	[(NSWindow*)id
		setIsMiniaturized: flag];
}

void NSWindow_inst_SetIsVisible(void *id, BOOL flag) {
	[(NSWindow*)id
		setIsVisible: flag];
}

void NSWindow_inst_SetIsZoomed(void *id, BOOL flag) {
	[(NSWindow*)id
		setIsZoomed: flag];
}

void NSWindow_inst_SetTitleWithRepresentedFilename(void *id, void* filename) {
	[(NSWindow*)id
		setTitleWithRepresentedFilename: filename];
}

void NSWindow_inst_ToggleFullScreen(void *id, void* sender) {
	[(NSWindow*)id
		toggleFullScreen: sender];
}

void NSWindow_inst_ToggleTabBar(void *id, void* sender) {
	[(NSWindow*)id
		toggleTabBar: sender];
}

void NSWindow_inst_ToggleTabOverview(void *id, void* sender) {
	[(NSWindow*)id
		toggleTabOverview: sender];
}

void NSWindow_inst_ToggleToolbarShown(void *id, void* sender) {
	[(NSWindow*)id
		toggleToolbarShown: sender];
}

BOOL NSWindow_inst_TryToPerformWith(void *id, void* action, void* object) {
	return [(NSWindow*)id
		tryToPerform: action
		with: object];
}

void NSWindow_inst_UnregisterDraggedTypes(void *id) {
	[(NSWindow*)id
		unregisterDraggedTypes];
}

void NSWindow_inst_Update(void *id) {
	[(NSWindow*)id
		update];
}

void NSWindow_inst_UpdateConstraintsIfNeeded(void *id) {
	[(NSWindow*)id
		updateConstraintsIfNeeded];
}

void NSWindow_inst_VisualizeConstraints(void *id, void* constraints) {
	[(NSWindow*)id
		visualizeConstraints: constraints];
}

void NSWindow_inst_Zoom(void *id, void* sender) {
	[(NSWindow*)id
		zoom: sender];
}

void* NSWindow_inst_Init(void *id) {
	return [(NSWindow*)id
		init];
}

void* NSWindow_inst_Delegate(void *id) {
	return [(NSWindow*)id
		delegate];
}

void NSWindow_inst_SetDelegate(void *id, void* value) {
	[(NSWindow*)id
		setDelegate: value];
}

void* NSWindow_inst_ContentViewController(void *id) {
	return [(NSWindow*)id
		contentViewController];
}

void NSWindow_inst_SetContentViewController(void *id, void* value) {
	[(NSWindow*)id
		setContentViewController: value];
}

void* NSWindow_inst_ContentView(void *id) {
	return [(NSWindow*)id
		contentView];
}

void NSWindow_inst_SetContentView(void *id, void* value) {
	[(NSWindow*)id
		setContentView: value];
}

unsigned long NSWindow_inst_StyleMask(void *id) {
	return [(NSWindow*)id
		styleMask];
}

void NSWindow_inst_SetStyleMask(void *id, unsigned long value) {
	[(NSWindow*)id
		setStyleMask: value];
}

BOOL NSWindow_inst_WorksWhenModal(void *id) {
	return [(NSWindow*)id
		worksWhenModal];
}

double NSWindow_inst_AlphaValue(void *id) {
	return [(NSWindow*)id
		alphaValue];
}

void NSWindow_inst_SetAlphaValue(void *id, double value) {
	[(NSWindow*)id
		setAlphaValue: value];
}

void* NSWindow_inst_BackgroundColor(void *id) {
	return [(NSWindow*)id
		backgroundColor];
}

void NSWindow_inst_SetBackgroundColor(void *id, void* value) {
	[(NSWindow*)id
		setBackgroundColor: value];
}

BOOL NSWindow_inst_CanHide(void *id) {
	return [(NSWindow*)id
		canHide];
}

void NSWindow_inst_SetCanHide(void *id, BOOL value) {
	[(NSWindow*)id
		setCanHide: value];
}

BOOL NSWindow_inst_IsOnActiveSpace(void *id) {
	return [(NSWindow*)id
		isOnActiveSpace];
}

BOOL NSWindow_inst_HidesOnDeactivate(void *id) {
	return [(NSWindow*)id
		hidesOnDeactivate];
}

void NSWindow_inst_SetHidesOnDeactivate(void *id, BOOL value) {
	[(NSWindow*)id
		setHidesOnDeactivate: value];
}

unsigned long NSWindow_inst_CollectionBehavior(void *id) {
	return [(NSWindow*)id
		collectionBehavior];
}

void NSWindow_inst_SetCollectionBehavior(void *id, unsigned long value) {
	[(NSWindow*)id
		setCollectionBehavior: value];
}

BOOL NSWindow_inst_IsOpaque(void *id) {
	return [(NSWindow*)id
		isOpaque];
}

void NSWindow_inst_SetOpaque(void *id, BOOL value) {
	[(NSWindow*)id
		setOpaque: value];
}

BOOL NSWindow_inst_HasShadow(void *id) {
	return [(NSWindow*)id
		hasShadow];
}

void NSWindow_inst_SetHasShadow(void *id, BOOL value) {
	[(NSWindow*)id
		setHasShadow: value];
}

BOOL NSWindow_inst_PreventsApplicationTerminationWhenModal(void *id) {
	return [(NSWindow*)id
		preventsApplicationTerminationWhenModal];
}

void NSWindow_inst_SetPreventsApplicationTerminationWhenModal(void *id, BOOL value) {
	[(NSWindow*)id
		setPreventsApplicationTerminationWhenModal: value];
}

BOOL NSWindow_inst_HasDynamicDepthLimit(void *id) {
	return [(NSWindow*)id
		hasDynamicDepthLimit];
}

long NSWindow_inst_WindowNumber(void *id) {
	return [(NSWindow*)id
		windowNumber];
}

void* NSWindow_inst_DeviceDescription(void *id) {
	return [(NSWindow*)id
		deviceDescription];
}

BOOL NSWindow_inst_CanBecomeVisibleWithoutLogin(void *id) {
	return [(NSWindow*)id
		canBecomeVisibleWithoutLogin];
}

void NSWindow_inst_SetCanBecomeVisibleWithoutLogin(void *id, BOOL value) {
	[(NSWindow*)id
		setCanBecomeVisibleWithoutLogin: value];
}

unsigned long NSWindow_inst_BackingType(void *id) {
	return [(NSWindow*)id
		backingType];
}

void NSWindow_inst_SetBackingType(void *id, unsigned long value) {
	[(NSWindow*)id
		setBackingType: value];
}

void* NSWindow_inst_AttachedSheet(void *id) {
	return [(NSWindow*)id
		attachedSheet];
}

BOOL NSWindow_inst_IsSheet(void *id) {
	return [(NSWindow*)id
		isSheet];
}

void* NSWindow_inst_SheetParent(void *id) {
	return [(NSWindow*)id
		sheetParent];
}

void* NSWindow_inst_Sheets(void *id) {
	return [(NSWindow*)id
		sheets];
}

NSRect NSWindow_inst_Frame(void *id) {
	return [(NSWindow*)id
		frame];
}

NSSize NSWindow_inst_AspectRatio(void *id) {
	return [(NSWindow*)id
		aspectRatio];
}

void NSWindow_inst_SetAspectRatio(void *id, NSSize value) {
	[(NSWindow*)id
		setAspectRatio: value];
}

NSSize NSWindow_inst_MinSize(void *id) {
	return [(NSWindow*)id
		minSize];
}

void NSWindow_inst_SetMinSize(void *id, NSSize value) {
	[(NSWindow*)id
		setMinSize: value];
}

NSSize NSWindow_inst_MaxSize(void *id) {
	return [(NSWindow*)id
		maxSize];
}

void NSWindow_inst_SetMaxSize(void *id, NSSize value) {
	[(NSWindow*)id
		setMaxSize: value];
}

BOOL NSWindow_inst_IsZoomed(void *id) {
	return [(NSWindow*)id
		isZoomed];
}

NSSize NSWindow_inst_ResizeIncrements(void *id) {
	return [(NSWindow*)id
		resizeIncrements];
}

void NSWindow_inst_SetResizeIncrements(void *id, NSSize value) {
	[(NSWindow*)id
		setResizeIncrements: value];
}

BOOL NSWindow_inst_PreservesContentDuringLiveResize(void *id) {
	return [(NSWindow*)id
		preservesContentDuringLiveResize];
}

void NSWindow_inst_SetPreservesContentDuringLiveResize(void *id, BOOL value) {
	[(NSWindow*)id
		setPreservesContentDuringLiveResize: value];
}

BOOL NSWindow_inst_InLiveResize(void *id) {
	return [(NSWindow*)id
		inLiveResize];
}

NSSize NSWindow_inst_ContentAspectRatio(void *id) {
	return [(NSWindow*)id
		contentAspectRatio];
}

void NSWindow_inst_SetContentAspectRatio(void *id, NSSize value) {
	[(NSWindow*)id
		setContentAspectRatio: value];
}

NSSize NSWindow_inst_ContentMinSize(void *id) {
	return [(NSWindow*)id
		contentMinSize];
}

void NSWindow_inst_SetContentMinSize(void *id, NSSize value) {
	[(NSWindow*)id
		setContentMinSize: value];
}

NSSize NSWindow_inst_ContentMaxSize(void *id) {
	return [(NSWindow*)id
		contentMaxSize];
}

void NSWindow_inst_SetContentMaxSize(void *id, NSSize value) {
	[(NSWindow*)id
		setContentMaxSize: value];
}

NSSize NSWindow_inst_ContentResizeIncrements(void *id) {
	return [(NSWindow*)id
		contentResizeIncrements];
}

void NSWindow_inst_SetContentResizeIncrements(void *id, NSSize value) {
	[(NSWindow*)id
		setContentResizeIncrements: value];
}

void* NSWindow_inst_ContentLayoutGuide(void *id) {
	return [(NSWindow*)id
		contentLayoutGuide];
}

NSRect NSWindow_inst_ContentLayoutRect(void *id) {
	return [(NSWindow*)id
		contentLayoutRect];
}

NSSize NSWindow_inst_MaxFullScreenContentSize(void *id) {
	return [(NSWindow*)id
		maxFullScreenContentSize];
}

void NSWindow_inst_SetMaxFullScreenContentSize(void *id, NSSize value) {
	[(NSWindow*)id
		setMaxFullScreenContentSize: value];
}

NSSize NSWindow_inst_MinFullScreenContentSize(void *id) {
	return [(NSWindow*)id
		minFullScreenContentSize];
}

void NSWindow_inst_SetMinFullScreenContentSize(void *id, NSSize value) {
	[(NSWindow*)id
		setMinFullScreenContentSize: value];
}

long NSWindow_inst_Level(void *id) {
	return [(NSWindow*)id
		level];
}

void NSWindow_inst_SetLevel(void *id, long value) {
	[(NSWindow*)id
		setLevel: value];
}

BOOL NSWindow_inst_IsVisible(void *id) {
	return [(NSWindow*)id
		isVisible];
}

BOOL NSWindow_inst_IsKeyWindow(void *id) {
	return [(NSWindow*)id
		isKeyWindow];
}

BOOL NSWindow_inst_CanBecomeKeyWindow(void *id) {
	return [(NSWindow*)id
		canBecomeKeyWindow];
}

BOOL NSWindow_inst_IsMainWindow(void *id) {
	return [(NSWindow*)id
		isMainWindow];
}

BOOL NSWindow_inst_CanBecomeMainWindow(void *id) {
	return [(NSWindow*)id
		canBecomeMainWindow];
}

void* NSWindow_inst_ChildWindows(void *id) {
	return [(NSWindow*)id
		childWindows];
}

void* NSWindow_inst_ParentWindow(void *id) {
	return [(NSWindow*)id
		parentWindow];
}

void NSWindow_inst_SetParentWindow(void *id, void* value) {
	[(NSWindow*)id
		setParentWindow: value];
}

BOOL NSWindow_inst_IsExcludedFromWindowsMenu(void *id) {
	return [(NSWindow*)id
		isExcludedFromWindowsMenu];
}

void NSWindow_inst_SetExcludedFromWindowsMenu(void *id, BOOL value) {
	[(NSWindow*)id
		setExcludedFromWindowsMenu: value];
}

BOOL NSWindow_inst_AreCursorRectsEnabled(void *id) {
	return [(NSWindow*)id
		areCursorRectsEnabled];
}

BOOL NSWindow_inst_ShowsToolbarButton(void *id) {
	return [(NSWindow*)id
		showsToolbarButton];
}

void NSWindow_inst_SetShowsToolbarButton(void *id, BOOL value) {
	[(NSWindow*)id
		setShowsToolbarButton: value];
}

BOOL NSWindow_inst_TitlebarAppearsTransparent(void *id) {
	return [(NSWindow*)id
		titlebarAppearsTransparent];
}

void NSWindow_inst_SetTitlebarAppearsTransparent(void *id, BOOL value) {
	[(NSWindow*)id
		setTitlebarAppearsTransparent: value];
}

void* NSWindow_inst_TitlebarAccessoryViewControllers(void *id) {
	return [(NSWindow*)id
		titlebarAccessoryViewControllers];
}

void NSWindow_inst_SetTitlebarAccessoryViewControllers(void *id, void* value) {
	[(NSWindow*)id
		setTitlebarAccessoryViewControllers: value];
}

void* NSWindow_inst_TabbedWindows(void *id) {
	return [(NSWindow*)id
		tabbedWindows];
}

BOOL NSWindow_inst_AllowsToolTipsWhenApplicationIsInactive(void *id) {
	return [(NSWindow*)id
		allowsToolTipsWhenApplicationIsInactive];
}

void NSWindow_inst_SetAllowsToolTipsWhenApplicationIsInactive(void *id, BOOL value) {
	[(NSWindow*)id
		setAllowsToolTipsWhenApplicationIsInactive: value];
}

void* NSWindow_inst_CurrentEvent(void *id) {
	return [(NSWindow*)id
		currentEvent];
}

void* NSWindow_inst_InitialFirstResponder(void *id) {
	return [(NSWindow*)id
		initialFirstResponder];
}

void NSWindow_inst_SetInitialFirstResponder(void *id, void* value) {
	[(NSWindow*)id
		setInitialFirstResponder: value];
}

BOOL NSWindow_inst_AutorecalculatesKeyViewLoop(void *id) {
	return [(NSWindow*)id
		autorecalculatesKeyViewLoop];
}

void NSWindow_inst_SetAutorecalculatesKeyViewLoop(void *id, BOOL value) {
	[(NSWindow*)id
		setAutorecalculatesKeyViewLoop: value];
}

BOOL NSWindow_inst_AcceptsMouseMovedEvents(void *id) {
	return [(NSWindow*)id
		acceptsMouseMovedEvents];
}

void NSWindow_inst_SetAcceptsMouseMovedEvents(void *id, BOOL value) {
	[(NSWindow*)id
		setAcceptsMouseMovedEvents: value];
}

BOOL NSWindow_inst_IgnoresMouseEvents(void *id) {
	return [(NSWindow*)id
		ignoresMouseEvents];
}

void NSWindow_inst_SetIgnoresMouseEvents(void *id, BOOL value) {
	[(NSWindow*)id
		setIgnoresMouseEvents: value];
}

NSPoint NSWindow_inst_MouseLocationOutsideOfEventStream(void *id) {
	return [(NSWindow*)id
		mouseLocationOutsideOfEventStream];
}

BOOL NSWindow_inst_IsRestorable(void *id) {
	return [(NSWindow*)id
		isRestorable];
}

void NSWindow_inst_SetRestorable(void *id, BOOL value) {
	[(NSWindow*)id
		setRestorable: value];
}

BOOL NSWindow_inst_ViewsNeedDisplay(void *id) {
	return [(NSWindow*)id
		viewsNeedDisplay];
}

void NSWindow_inst_SetViewsNeedDisplay(void *id, BOOL value) {
	[(NSWindow*)id
		setViewsNeedDisplay: value];
}

BOOL NSWindow_inst_AllowsConcurrentViewDrawing(void *id) {
	return [(NSWindow*)id
		allowsConcurrentViewDrawing];
}

void NSWindow_inst_SetAllowsConcurrentViewDrawing(void *id, BOOL value) {
	[(NSWindow*)id
		setAllowsConcurrentViewDrawing: value];
}

BOOL NSWindow_inst_IsDocumentEdited(void *id) {
	return [(NSWindow*)id
		isDocumentEdited];
}

void NSWindow_inst_SetDocumentEdited(void *id, BOOL value) {
	[(NSWindow*)id
		setDocumentEdited: value];
}

double NSWindow_inst_BackingScaleFactor(void *id) {
	return [(NSWindow*)id
		backingScaleFactor];
}

void* NSWindow_inst_Title(void *id) {
	return [(NSWindow*)id
		title];
}

void NSWindow_inst_SetTitle(void *id, void* value) {
	[(NSWindow*)id
		setTitle: value];
}

void* NSWindow_inst_Subtitle(void *id) {
	return [(NSWindow*)id
		subtitle];
}

void NSWindow_inst_SetSubtitle(void *id, void* value) {
	[(NSWindow*)id
		setSubtitle: value];
}

long NSWindow_inst_TitleVisibility(void *id) {
	return [(NSWindow*)id
		titleVisibility];
}

void NSWindow_inst_SetTitleVisibility(void *id, long value) {
	[(NSWindow*)id
		setTitleVisibility: value];
}

void* NSWindow_inst_RepresentedFilename(void *id) {
	return [(NSWindow*)id
		representedFilename];
}

void NSWindow_inst_SetRepresentedFilename(void *id, void* value) {
	[(NSWindow*)id
		setRepresentedFilename: value];
}

void* NSWindow_inst_RepresentedURL(void *id) {
	return [(NSWindow*)id
		representedURL];
}

void NSWindow_inst_SetRepresentedURL(void *id, void* value) {
	[(NSWindow*)id
		setRepresentedURL: value];
}

void* NSWindow_inst_Screen(void *id) {
	return [(NSWindow*)id
		screen];
}

void* NSWindow_inst_DeepestScreen(void *id) {
	return [(NSWindow*)id
		deepestScreen];
}

BOOL NSWindow_inst_DisplaysWhenScreenProfileChanges(void *id) {
	return [(NSWindow*)id
		displaysWhenScreenProfileChanges];
}

void NSWindow_inst_SetDisplaysWhenScreenProfileChanges(void *id, BOOL value) {
	[(NSWindow*)id
		setDisplaysWhenScreenProfileChanges: value];
}

BOOL NSWindow_inst_IsMovableByWindowBackground(void *id) {
	return [(NSWindow*)id
		isMovableByWindowBackground];
}

void NSWindow_inst_SetMovableByWindowBackground(void *id, BOOL value) {
	[(NSWindow*)id
		setMovableByWindowBackground: value];
}

BOOL NSWindow_inst_IsMovable(void *id) {
	return [(NSWindow*)id
		isMovable];
}

void NSWindow_inst_SetMovable(void *id, BOOL value) {
	[(NSWindow*)id
		setMovable: value];
}

BOOL NSWindow_inst_IsReleasedWhenClosed(void *id) {
	return [(NSWindow*)id
		isReleasedWhenClosed];
}

void NSWindow_inst_SetReleasedWhenClosed(void *id, BOOL value) {
	[(NSWindow*)id
		setReleasedWhenClosed: value];
}

BOOL NSWindow_inst_IsMiniaturized(void *id) {
	return [(NSWindow*)id
		isMiniaturized];
}

void* NSWindow_inst_MiniwindowImage(void *id) {
	return [(NSWindow*)id
		miniwindowImage];
}

void NSWindow_inst_SetMiniwindowImage(void *id, void* value) {
	[(NSWindow*)id
		setMiniwindowImage: value];
}

void* NSWindow_inst_MiniwindowTitle(void *id) {
	return [(NSWindow*)id
		miniwindowTitle];
}

void NSWindow_inst_SetMiniwindowTitle(void *id, void* value) {
	[(NSWindow*)id
		setMiniwindowTitle: value];
}

BOOL NSWindow_inst_HasCloseBox(void *id) {
	return [(NSWindow*)id
		hasCloseBox];
}

BOOL NSWindow_inst_HasTitleBar(void *id) {
	return [(NSWindow*)id
		hasTitleBar];
}

BOOL NSWindow_inst_IsModalPanel(void *id) {
	return [(NSWindow*)id
		isModalPanel];
}

BOOL NSWindow_inst_IsFloatingPanel(void *id) {
	return [(NSWindow*)id
		isFloatingPanel];
}

BOOL NSWindow_inst_IsZoomable(void *id) {
	return [(NSWindow*)id
		isZoomable];
}

BOOL NSWindow_inst_IsResizable(void *id) {
	return [(NSWindow*)id
		isResizable];
}

BOOL NSWindow_inst_IsMiniaturizable(void *id) {
	return [(NSWindow*)id
		isMiniaturizable];
}

long NSWindow_inst_OrderedIndex(void *id) {
	return [(NSWindow*)id
		orderedIndex];
}

void NSWindow_inst_SetOrderedIndex(void *id, long value) {
	[(NSWindow*)id
		setOrderedIndex: value];
}

void* NSWorkspace_inst_URLForApplicationToOpenURL(void *id, void* url) {
	return [(NSWorkspace*)id
		URLForApplicationToOpenURL: url];
}

void* NSWorkspace_inst_URLForApplicationWithBundleIdentifier(void *id, void* bundleIdentifier) {
	return [(NSWorkspace*)id
		URLForApplicationWithBundleIdentifier: bundleIdentifier];
}

void* NSWorkspace_inst_URLsForApplicationsToOpenURL(void *id, void* url) {
	return [(NSWorkspace*)id
		URLsForApplicationsToOpenURL: url];
}

void* NSWorkspace_inst_URLsForApplicationsWithBundleIdentifier(void *id, void* bundleIdentifier) {
	return [(NSWorkspace*)id
		URLsForApplicationsWithBundleIdentifier: bundleIdentifier];
}

void NSWorkspace_inst_ActivateFileViewerSelectingURLs(void *id, void* fileURLs) {
	[(NSWorkspace*)id
		activateFileViewerSelectingURLs: fileURLs];
}

void* NSWorkspace_inst_DesktopImageOptionsForScreen(void *id, void* screen) {
	return [(NSWorkspace*)id
		desktopImageOptionsForScreen: screen];
}

void* NSWorkspace_inst_DesktopImageURLForScreen(void *id, void* screen) {
	return [(NSWorkspace*)id
		desktopImageURLForScreen: screen];
}

long NSWorkspace_inst_ExtendPowerOffBy(void *id, long requested) {
	return [(NSWorkspace*)id
		extendPowerOffBy: requested];
}

void NSWorkspace_inst_HideOtherApplications(void *id) {
	[(NSWorkspace*)id
		hideOtherApplications];
}

void* NSWorkspace_inst_IconForFile(void *id, void* fullPath) {
	return [(NSWorkspace*)id
		iconForFile: fullPath];
}

void* NSWorkspace_inst_IconForFiles(void *id, void* fullPaths) {
	return [(NSWorkspace*)id
		iconForFiles: fullPaths];
}

BOOL NSWorkspace_inst_IsFilePackageAtPath(void *id, void* fullPath) {
	return [(NSWorkspace*)id
		isFilePackageAtPath: fullPath];
}

void NSWorkspace_inst_NoteFileSystemChanged(void *id, void* path) {
	[(NSWorkspace*)id
		noteFileSystemChanged: path];
}

BOOL NSWorkspace_inst_OpenURL(void *id, void* url) {
	return [(NSWorkspace*)id
		openURL: url];
}

BOOL NSWorkspace_inst_SelectFileInFileViewerRootedAtPath(void *id, void* fullPath, void* rootFullPath) {
	return [(NSWorkspace*)id
		selectFile: fullPath
		inFileViewerRootedAtPath: rootFullPath];
}

BOOL NSWorkspace_inst_ShowSearchResultsForQueryString(void *id, void* queryString) {
	return [(NSWorkspace*)id
		showSearchResultsForQueryString: queryString];
}

BOOL NSWorkspace_inst_UnmountAndEjectDeviceAtPath(void *id, void* path) {
	return [(NSWorkspace*)id
		unmountAndEjectDeviceAtPath: path];
}

void* NSWorkspace_inst_Init(void *id) {
	return [(NSWorkspace*)id
		init];
}

void* NSWorkspace_inst_FrontmostApplication(void *id) {
	return [(NSWorkspace*)id
		frontmostApplication];
}

void* NSWorkspace_inst_RunningApplications(void *id) {
	return [(NSWorkspace*)id
		runningApplications];
}

void* NSWorkspace_inst_MenuBarOwningApplication(void *id) {
	return [(NSWorkspace*)id
		menuBarOwningApplication];
}

void* NSWorkspace_inst_FileLabels(void *id) {
	return [(NSWorkspace*)id
		fileLabels];
}

void* NSWorkspace_inst_FileLabelColors(void *id) {
	return [(NSWorkspace*)id
		fileLabelColors];
}

BOOL NSWorkspace_inst_AccessibilityDisplayShouldDifferentiateWithoutColor(void *id) {
	return [(NSWorkspace*)id
		accessibilityDisplayShouldDifferentiateWithoutColor];
}

BOOL NSWorkspace_inst_AccessibilityDisplayShouldIncreaseContrast(void *id) {
	return [(NSWorkspace*)id
		accessibilityDisplayShouldIncreaseContrast];
}

BOOL NSWorkspace_inst_AccessibilityDisplayShouldReduceTransparency(void *id) {
	return [(NSWorkspace*)id
		accessibilityDisplayShouldReduceTransparency];
}

BOOL NSWorkspace_inst_AccessibilityDisplayShouldInvertColors(void *id) {
	return [(NSWorkspace*)id
		accessibilityDisplayShouldInvertColors];
}

BOOL NSWorkspace_inst_AccessibilityDisplayShouldReduceMotion(void *id) {
	return [(NSWorkspace*)id
		accessibilityDisplayShouldReduceMotion];
}

BOOL NSWorkspace_inst_IsSwitchControlEnabled(void *id) {
	return [(NSWorkspace*)id
		isSwitchControlEnabled];
}

BOOL NSWorkspace_inst_IsVoiceOverEnabled(void *id) {
	return [(NSWorkspace*)id
		isVoiceOverEnabled];
}

void* NSColor_inst_BlendedColorWithFractionOfColor(void *id, double fraction, void* color) {
	return [(NSColor*)id
		blendedColorWithFraction: fraction
		ofColor: color];
}

void* NSColor_inst_ColorWithAlphaComponent(void *id, double alpha) {
	return [(NSColor*)id
		colorWithAlphaComponent: alpha];
}

void NSColor_inst_DrawSwatchInRect(void *id, NSRect rect) {
	[(NSColor*)id
		drawSwatchInRect: rect];
}

void* NSColor_inst_HighlightWithLevel(void *id, double val) {
	return [(NSColor*)id
		highlightWithLevel: val];
}

void NSColor_inst_Set(void *id) {
	[(NSColor*)id
		set];
}

void NSColor_inst_SetFill(void *id) {
	[(NSColor*)id
		setFill];
}

void NSColor_inst_SetStroke(void *id) {
	[(NSColor*)id
		setStroke];
}

void* NSColor_inst_ShadowWithLevel(void *id, double val) {
	return [(NSColor*)id
		shadowWithLevel: val];
}

void NSColor_inst_WriteToPasteboard(void *id, void* pasteBoard) {
	[(NSColor*)id
		writeToPasteboard: pasteBoard];
}

void* NSColor_inst_Init(void *id) {
	return [(NSColor*)id
		init];
}

long NSColor_inst_NumberOfComponents(void *id) {
	return [(NSColor*)id
		numberOfComponents];
}

double NSColor_inst_AlphaComponent(void *id) {
	return [(NSColor*)id
		alphaComponent];
}

double NSColor_inst_WhiteComponent(void *id) {
	return [(NSColor*)id
		whiteComponent];
}

double NSColor_inst_RedComponent(void *id) {
	return [(NSColor*)id
		redComponent];
}

double NSColor_inst_GreenComponent(void *id) {
	return [(NSColor*)id
		greenComponent];
}

double NSColor_inst_BlueComponent(void *id) {
	return [(NSColor*)id
		blueComponent];
}

double NSColor_inst_CyanComponent(void *id) {
	return [(NSColor*)id
		cyanComponent];
}

double NSColor_inst_MagentaComponent(void *id) {
	return [(NSColor*)id
		magentaComponent];
}

double NSColor_inst_YellowComponent(void *id) {
	return [(NSColor*)id
		yellowComponent];
}

double NSColor_inst_BlackComponent(void *id) {
	return [(NSColor*)id
		blackComponent];
}

double NSColor_inst_HueComponent(void *id) {
	return [(NSColor*)id
		hueComponent];
}

double NSColor_inst_SaturationComponent(void *id) {
	return [(NSColor*)id
		saturationComponent];
}

double NSColor_inst_BrightnessComponent(void *id) {
	return [(NSColor*)id
		brightnessComponent];
}

void* NSColor_inst_LocalizedCatalogNameComponent(void *id) {
	return [(NSColor*)id
		localizedCatalogNameComponent];
}

void* NSColor_inst_LocalizedColorNameComponent(void *id) {
	return [(NSColor*)id
		localizedColorNameComponent];
}

void NSTextView_inst_AlignJustified(void *id, void* sender) {
	[(NSTextView*)id
		alignJustified: sender];
}

void NSTextView_inst_BreakUndoCoalescing(void *id) {
	[(NSTextView*)id
		breakUndoCoalescing];
}

void NSTextView_inst_ChangeAttributes(void *id, void* sender) {
	[(NSTextView*)id
		changeAttributes: sender];
}

void NSTextView_inst_ChangeColor(void *id, void* sender) {
	[(NSTextView*)id
		changeColor: sender];
}

void NSTextView_inst_ChangeDocumentBackgroundColor(void *id, void* sender) {
	[(NSTextView*)id
		changeDocumentBackgroundColor: sender];
}

void NSTextView_inst_ChangeLayoutOrientation(void *id, void* sender) {
	[(NSTextView*)id
		changeLayoutOrientation: sender];
}

unsigned long NSTextView_inst_CharacterIndexForInsertionAtPoint(void *id, NSPoint point) {
	return [(NSTextView*)id
		characterIndexForInsertionAtPoint: point];
}

void NSTextView_inst_CheckTextInDocument(void *id, void* sender) {
	[(NSTextView*)id
		checkTextInDocument: sender];
}

void NSTextView_inst_CheckTextInSelection(void *id, void* sender) {
	[(NSTextView*)id
		checkTextInSelection: sender];
}

void NSTextView_inst_CleanUpAfterDragOperation(void *id) {
	[(NSTextView*)id
		cleanUpAfterDragOperation];
}

void NSTextView_inst_ClickedOnLinkAtIndex(void *id, void* link, unsigned long charIndex) {
	[(NSTextView*)id
		clickedOnLink: link
		atIndex: charIndex];
}

void NSTextView_inst_Complete(void *id, void* sender) {
	[(NSTextView*)id
		complete: sender];
}

void NSTextView_inst_DidChangeText(void *id) {
	[(NSTextView*)id
		didChangeText];
}

BOOL NSTextView_inst_DragSelectionWithEventOffsetSlideBack(void *id, void* event, NSSize mouseOffset, BOOL slideBack) {
	return [(NSTextView*)id
		dragSelectionWithEvent: event
		offset: mouseOffset
		slideBack: slideBack];
}

void NSTextView_inst_DrawInsertionPointInRectColorTurnedOn(void *id, NSRect rect, void* color, BOOL flag) {
	[(NSTextView*)id
		drawInsertionPointInRect: rect
		color: color
		turnedOn: flag];
}

void NSTextView_inst_DrawViewBackgroundInRect(void *id, NSRect rect) {
	[(NSTextView*)id
		drawViewBackgroundInRect: rect];
}

void* NSTextView_inst_InitWithFrame(void *id, NSRect frameRect) {
	return [(NSTextView*)id
		initWithFrame: frameRect];
}

void* NSTextView_inst_InitWithFrameTextContainer(void *id, NSRect frameRect, void* container) {
	return [(NSTextView*)id
		initWithFrame: frameRect
		textContainer: container];
}

void NSTextView_inst_InvalidateTextContainerOrigin(void *id) {
	[(NSTextView*)id
		invalidateTextContainerOrigin];
}

void NSTextView_inst_LoosenKerning(void *id, void* sender) {
	[(NSTextView*)id
		loosenKerning: sender];
}

void NSTextView_inst_LowerBaseline(void *id, void* sender) {
	[(NSTextView*)id
		lowerBaseline: sender];
}

void NSTextView_inst_OrderFrontLinkPanel(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontLinkPanel: sender];
}

void NSTextView_inst_OrderFrontListPanel(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontListPanel: sender];
}

void NSTextView_inst_OrderFrontSharingServicePicker(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontSharingServicePicker: sender];
}

void NSTextView_inst_OrderFrontSpacingPanel(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontSpacingPanel: sender];
}

void NSTextView_inst_OrderFrontSubstitutionsPanel(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontSubstitutionsPanel: sender];
}

void NSTextView_inst_OrderFrontTablePanel(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontTablePanel: sender];
}

void NSTextView_inst_Outline(void *id, void* sender) {
	[(NSTextView*)id
		outline: sender];
}

void NSTextView_inst_PasteAsPlainText(void *id, void* sender) {
	[(NSTextView*)id
		pasteAsPlainText: sender];
}

void NSTextView_inst_PasteAsRichText(void *id, void* sender) {
	[(NSTextView*)id
		pasteAsRichText: sender];
}

void NSTextView_inst_PerformFindPanelAction(void *id, void* sender) {
	[(NSTextView*)id
		performFindPanelAction: sender];
}

void* NSTextView_inst_QuickLookPreviewableItemsInRanges(void *id, void* ranges) {
	return [(NSTextView*)id
		quickLookPreviewableItemsInRanges: ranges];
}

void NSTextView_inst_RaiseBaseline(void *id, void* sender) {
	[(NSTextView*)id
		raiseBaseline: sender];
}

BOOL NSTextView_inst_ReadSelectionFromPasteboard(void *id, void* pboard) {
	return [(NSTextView*)id
		readSelectionFromPasteboard: pboard];
}

void NSTextView_inst_ReplaceTextContainer(void *id, void* newContainer) {
	[(NSTextView*)id
		replaceTextContainer: newContainer];
}

void NSTextView_inst_SetConstrainedFrameSize(void *id, NSSize desiredSize) {
	[(NSTextView*)id
		setConstrainedFrameSize: desiredSize];
}

void NSTextView_inst_SetNeedsDisplayInRectAvoidAdditionalLayout(void *id, NSRect rect, BOOL flag) {
	[(NSTextView*)id
		setNeedsDisplayInRect: rect
		avoidAdditionalLayout: flag];
}

BOOL NSTextView_inst_ShouldChangeTextInRangesReplacementStrings(void *id, void* affectedRanges, void* replacementStrings) {
	return [(NSTextView*)id
		shouldChangeTextInRanges: affectedRanges
		replacementStrings: replacementStrings];
}

void NSTextView_inst_StartSpeaking(void *id, void* sender) {
	[(NSTextView*)id
		startSpeaking: sender];
}

void NSTextView_inst_StopSpeaking(void *id, void* sender) {
	[(NSTextView*)id
		stopSpeaking: sender];
}

void NSTextView_inst_TightenKerning(void *id, void* sender) {
	[(NSTextView*)id
		tightenKerning: sender];
}

void NSTextView_inst_ToggleAutomaticDashSubstitution(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticDashSubstitution: sender];
}

void NSTextView_inst_ToggleAutomaticDataDetection(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticDataDetection: sender];
}

void NSTextView_inst_ToggleAutomaticLinkDetection(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticLinkDetection: sender];
}

void NSTextView_inst_ToggleAutomaticQuoteSubstitution(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticQuoteSubstitution: sender];
}

void NSTextView_inst_ToggleAutomaticSpellingCorrection(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticSpellingCorrection: sender];
}

void NSTextView_inst_ToggleAutomaticTextCompletion(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticTextCompletion: sender];
}

void NSTextView_inst_ToggleAutomaticTextReplacement(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticTextReplacement: sender];
}

void NSTextView_inst_ToggleContinuousSpellChecking(void *id, void* sender) {
	[(NSTextView*)id
		toggleContinuousSpellChecking: sender];
}

void NSTextView_inst_ToggleGrammarChecking(void *id, void* sender) {
	[(NSTextView*)id
		toggleGrammarChecking: sender];
}

void NSTextView_inst_ToggleQuickLookPreviewPanel(void *id, void* sender) {
	[(NSTextView*)id
		toggleQuickLookPreviewPanel: sender];
}

void NSTextView_inst_ToggleSmartInsertDelete(void *id, void* sender) {
	[(NSTextView*)id
		toggleSmartInsertDelete: sender];
}

void NSTextView_inst_TurnOffKerning(void *id, void* sender) {
	[(NSTextView*)id
		turnOffKerning: sender];
}

void NSTextView_inst_TurnOffLigatures(void *id, void* sender) {
	[(NSTextView*)id
		turnOffLigatures: sender];
}

void NSTextView_inst_UpdateCandidates(void *id) {
	[(NSTextView*)id
		updateCandidates];
}

void NSTextView_inst_UpdateDragTypeRegistration(void *id) {
	[(NSTextView*)id
		updateDragTypeRegistration];
}

void NSTextView_inst_UpdateFontPanel(void *id) {
	[(NSTextView*)id
		updateFontPanel];
}

void NSTextView_inst_UpdateInsertionPointStateAndRestartTimer(void *id, BOOL restartFlag) {
	[(NSTextView*)id
		updateInsertionPointStateAndRestartTimer: restartFlag];
}

void NSTextView_inst_UpdateQuickLookPreviewPanel(void *id) {
	[(NSTextView*)id
		updateQuickLookPreviewPanel];
}

void NSTextView_inst_UpdateRuler(void *id) {
	[(NSTextView*)id
		updateRuler];
}

void NSTextView_inst_UpdateTextTouchBarItems(void *id) {
	[(NSTextView*)id
		updateTextTouchBarItems];
}

void NSTextView_inst_UpdateTouchBarItemIdentifiers(void *id) {
	[(NSTextView*)id
		updateTouchBarItemIdentifiers];
}

void NSTextView_inst_UseAllLigatures(void *id, void* sender) {
	[(NSTextView*)id
		useAllLigatures: sender];
}

void NSTextView_inst_UseStandardKerning(void *id, void* sender) {
	[(NSTextView*)id
		useStandardKerning: sender];
}

void NSTextView_inst_UseStandardLigatures(void *id, void* sender) {
	[(NSTextView*)id
		useStandardLigatures: sender];
}

BOOL NSTextView_inst_WriteSelectionToPasteboardTypes(void *id, void* pboard, void* types) {
	return [(NSTextView*)id
		writeSelectionToPasteboard: pboard
		types: types];
}

void* NSTextView_inst_Init(void *id) {
	return [(NSTextView*)id
		init];
}

void* NSTextView_inst_Delegate(void *id) {
	return [(NSTextView*)id
		delegate];
}

void NSTextView_inst_SetDelegate(void *id, void* value) {
	[(NSTextView*)id
		setDelegate: value];
}

void* NSTextView_inst_TextContainer(void *id) {
	return [(NSTextView*)id
		textContainer];
}

void NSTextView_inst_SetTextContainer(void *id, void* value) {
	[(NSTextView*)id
		setTextContainer: value];
}

NSSize NSTextView_inst_TextContainerInset(void *id) {
	return [(NSTextView*)id
		textContainerInset];
}

void NSTextView_inst_SetTextContainerInset(void *id, NSSize value) {
	[(NSTextView*)id
		setTextContainerInset: value];
}

NSPoint NSTextView_inst_TextContainerOrigin(void *id) {
	return [(NSTextView*)id
		textContainerOrigin];
}

void* NSTextView_inst_LayoutManager(void *id) {
	return [(NSTextView*)id
		layoutManager];
}

void* NSTextView_inst_BackgroundColor(void *id) {
	return [(NSTextView*)id
		backgroundColor];
}

void NSTextView_inst_SetBackgroundColor(void *id, void* value) {
	[(NSTextView*)id
		setBackgroundColor: value];
}

BOOL NSTextView_inst_DrawsBackground(void *id) {
	return [(NSTextView*)id
		drawsBackground];
}

void NSTextView_inst_SetDrawsBackground(void *id, BOOL value) {
	[(NSTextView*)id
		setDrawsBackground: value];
}

BOOL NSTextView_inst_AllowsDocumentBackgroundColorChange(void *id) {
	return [(NSTextView*)id
		allowsDocumentBackgroundColorChange];
}

void NSTextView_inst_SetAllowsDocumentBackgroundColorChange(void *id, BOOL value) {
	[(NSTextView*)id
		setAllowsDocumentBackgroundColorChange: value];
}

BOOL NSTextView_inst_ShouldDrawInsertionPoint(void *id) {
	return [(NSTextView*)id
		shouldDrawInsertionPoint];
}

void* NSTextView_inst_AllowedInputSourceLocales(void *id) {
	return [(NSTextView*)id
		allowedInputSourceLocales];
}

void NSTextView_inst_SetAllowedInputSourceLocales(void *id, void* value) {
	[(NSTextView*)id
		setAllowedInputSourceLocales: value];
}

BOOL NSTextView_inst_AllowsUndo(void *id) {
	return [(NSTextView*)id
		allowsUndo];
}

void NSTextView_inst_SetAllowsUndo(void *id, BOOL value) {
	[(NSTextView*)id
		setAllowsUndo: value];
}

BOOL NSTextView_inst_IsEditable(void *id) {
	return [(NSTextView*)id
		isEditable];
}

void NSTextView_inst_SetEditable(void *id, BOOL value) {
	[(NSTextView*)id
		setEditable: value];
}

BOOL NSTextView_inst_IsSelectable(void *id) {
	return [(NSTextView*)id
		isSelectable];
}

void NSTextView_inst_SetSelectable(void *id, BOOL value) {
	[(NSTextView*)id
		setSelectable: value];
}

BOOL NSTextView_inst_IsFieldEditor(void *id) {
	return [(NSTextView*)id
		isFieldEditor];
}

void NSTextView_inst_SetFieldEditor(void *id, BOOL value) {
	[(NSTextView*)id
		setFieldEditor: value];
}

BOOL NSTextView_inst_IsRichText(void *id) {
	return [(NSTextView*)id
		isRichText];
}

void NSTextView_inst_SetRichText(void *id, BOOL value) {
	[(NSTextView*)id
		setRichText: value];
}

BOOL NSTextView_inst_ImportsGraphics(void *id) {
	return [(NSTextView*)id
		importsGraphics];
}

void NSTextView_inst_SetImportsGraphics(void *id, BOOL value) {
	[(NSTextView*)id
		setImportsGraphics: value];
}

BOOL NSTextView_inst_AllowsImageEditing(void *id) {
	return [(NSTextView*)id
		allowsImageEditing];
}

void NSTextView_inst_SetAllowsImageEditing(void *id, BOOL value) {
	[(NSTextView*)id
		setAllowsImageEditing: value];
}

BOOL NSTextView_inst_IsAutomaticQuoteSubstitutionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticQuoteSubstitutionEnabled];
}

void NSTextView_inst_SetAutomaticQuoteSubstitutionEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticQuoteSubstitutionEnabled: value];
}

BOOL NSTextView_inst_IsAutomaticLinkDetectionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticLinkDetectionEnabled];
}

void NSTextView_inst_SetAutomaticLinkDetectionEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticLinkDetectionEnabled: value];
}

BOOL NSTextView_inst_DisplaysLinkToolTips(void *id) {
	return [(NSTextView*)id
		displaysLinkToolTips];
}

void NSTextView_inst_SetDisplaysLinkToolTips(void *id, BOOL value) {
	[(NSTextView*)id
		setDisplaysLinkToolTips: value];
}

BOOL NSTextView_inst_IsAutomaticTextCompletionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticTextCompletionEnabled];
}

void NSTextView_inst_SetAutomaticTextCompletionEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticTextCompletionEnabled: value];
}

BOOL NSTextView_inst_UsesAdaptiveColorMappingForDarkAppearance(void *id) {
	return [(NSTextView*)id
		usesAdaptiveColorMappingForDarkAppearance];
}

void NSTextView_inst_SetUsesAdaptiveColorMappingForDarkAppearance(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesAdaptiveColorMappingForDarkAppearance: value];
}

BOOL NSTextView_inst_UsesRolloverButtonForSelection(void *id) {
	return [(NSTextView*)id
		usesRolloverButtonForSelection];
}

void NSTextView_inst_SetUsesRolloverButtonForSelection(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesRolloverButtonForSelection: value];
}

BOOL NSTextView_inst_UsesRuler(void *id) {
	return [(NSTextView*)id
		usesRuler];
}

void NSTextView_inst_SetUsesRuler(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesRuler: value];
}

BOOL NSTextView_inst_IsRulerVisible(void *id) {
	return [(NSTextView*)id
		isRulerVisible];
}

void NSTextView_inst_SetRulerVisible(void *id, BOOL value) {
	[(NSTextView*)id
		setRulerVisible: value];
}

BOOL NSTextView_inst_UsesInspectorBar(void *id) {
	return [(NSTextView*)id
		usesInspectorBar];
}

void NSTextView_inst_SetUsesInspectorBar(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesInspectorBar: value];
}

void* NSTextView_inst_SelectedRanges(void *id) {
	return [(NSTextView*)id
		selectedRanges];
}

void NSTextView_inst_SetSelectedRanges(void *id, void* value) {
	[(NSTextView*)id
		setSelectedRanges: value];
}

void* NSTextView_inst_InsertionPointColor(void *id) {
	return [(NSTextView*)id
		insertionPointColor];
}

void NSTextView_inst_SetInsertionPointColor(void *id, void* value) {
	[(NSTextView*)id
		setInsertionPointColor: value];
}

void* NSTextView_inst_SelectedTextAttributes(void *id) {
	return [(NSTextView*)id
		selectedTextAttributes];
}

void NSTextView_inst_SetSelectedTextAttributes(void *id, void* value) {
	[(NSTextView*)id
		setSelectedTextAttributes: value];
}

void* NSTextView_inst_MarkedTextAttributes(void *id) {
	return [(NSTextView*)id
		markedTextAttributes];
}

void NSTextView_inst_SetMarkedTextAttributes(void *id, void* value) {
	[(NSTextView*)id
		setMarkedTextAttributes: value];
}

void* NSTextView_inst_LinkTextAttributes(void *id) {
	return [(NSTextView*)id
		linkTextAttributes];
}

void NSTextView_inst_SetLinkTextAttributes(void *id, void* value) {
	[(NSTextView*)id
		setLinkTextAttributes: value];
}

void* NSTextView_inst_ReadablePasteboardTypes(void *id) {
	return [(NSTextView*)id
		readablePasteboardTypes];
}

void* NSTextView_inst_WritablePasteboardTypes(void *id) {
	return [(NSTextView*)id
		writablePasteboardTypes];
}

void* NSTextView_inst_TypingAttributes(void *id) {
	return [(NSTextView*)id
		typingAttributes];
}

void NSTextView_inst_SetTypingAttributes(void *id, void* value) {
	[(NSTextView*)id
		setTypingAttributes: value];
}

BOOL NSTextView_inst_IsCoalescingUndo(void *id) {
	return [(NSTextView*)id
		isCoalescingUndo];
}

void* NSTextView_inst_AcceptableDragTypes(void *id) {
	return [(NSTextView*)id
		acceptableDragTypes];
}

void* NSTextView_inst_RangesForUserCharacterAttributeChange(void *id) {
	return [(NSTextView*)id
		rangesForUserCharacterAttributeChange];
}

void* NSTextView_inst_RangesForUserParagraphAttributeChange(void *id) {
	return [(NSTextView*)id
		rangesForUserParagraphAttributeChange];
}

void* NSTextView_inst_RangesForUserTextChange(void *id) {
	return [(NSTextView*)id
		rangesForUserTextChange];
}

BOOL NSTextView_inst_SmartInsertDeleteEnabled(void *id) {
	return [(NSTextView*)id
		smartInsertDeleteEnabled];
}

void NSTextView_inst_SetSmartInsertDeleteEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setSmartInsertDeleteEnabled: value];
}

BOOL NSTextView_inst_IsContinuousSpellCheckingEnabled(void *id) {
	return [(NSTextView*)id
		isContinuousSpellCheckingEnabled];
}

void NSTextView_inst_SetContinuousSpellCheckingEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setContinuousSpellCheckingEnabled: value];
}

long NSTextView_inst_SpellCheckerDocumentTag(void *id) {
	return [(NSTextView*)id
		spellCheckerDocumentTag];
}

BOOL NSTextView_inst_IsGrammarCheckingEnabled(void *id) {
	return [(NSTextView*)id
		isGrammarCheckingEnabled];
}

void NSTextView_inst_SetGrammarCheckingEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setGrammarCheckingEnabled: value];
}

BOOL NSTextView_inst_AcceptsGlyphInfo(void *id) {
	return [(NSTextView*)id
		acceptsGlyphInfo];
}

void NSTextView_inst_SetAcceptsGlyphInfo(void *id, BOOL value) {
	[(NSTextView*)id
		setAcceptsGlyphInfo: value];
}

BOOL NSTextView_inst_UsesFontPanel(void *id) {
	return [(NSTextView*)id
		usesFontPanel];
}

void NSTextView_inst_SetUsesFontPanel(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesFontPanel: value];
}

BOOL NSTextView_inst_UsesFindPanel(void *id) {
	return [(NSTextView*)id
		usesFindPanel];
}

void NSTextView_inst_SetUsesFindPanel(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesFindPanel: value];
}

BOOL NSTextView_inst_IsAutomaticDashSubstitutionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticDashSubstitutionEnabled];
}

void NSTextView_inst_SetAutomaticDashSubstitutionEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticDashSubstitutionEnabled: value];
}

BOOL NSTextView_inst_IsAutomaticDataDetectionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticDataDetectionEnabled];
}

void NSTextView_inst_SetAutomaticDataDetectionEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticDataDetectionEnabled: value];
}

BOOL NSTextView_inst_IsAutomaticSpellingCorrectionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticSpellingCorrectionEnabled];
}

void NSTextView_inst_SetAutomaticSpellingCorrectionEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticSpellingCorrectionEnabled: value];
}

BOOL NSTextView_inst_IsAutomaticTextReplacementEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticTextReplacementEnabled];
}

void NSTextView_inst_SetAutomaticTextReplacementEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticTextReplacementEnabled: value];
}

BOOL NSTextView_inst_UsesFindBar(void *id) {
	return [(NSTextView*)id
		usesFindBar];
}

void NSTextView_inst_SetUsesFindBar(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesFindBar: value];
}

BOOL NSTextView_inst_IsIncrementalSearchingEnabled(void *id) {
	return [(NSTextView*)id
		isIncrementalSearchingEnabled];
}

void NSTextView_inst_SetIncrementalSearchingEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setIncrementalSearchingEnabled: value];
}

BOOL NSTextView_inst_AllowsCharacterPickerTouchBarItem(void *id) {
	return [(NSTextView*)id
		allowsCharacterPickerTouchBarItem];
}

void NSTextView_inst_SetAllowsCharacterPickerTouchBarItem(void *id, BOOL value) {
	[(NSTextView*)id
		setAllowsCharacterPickerTouchBarItem: value];
}

void* NSTextView_inst_Font(void *id) {
	return [(NSTextView*)id
		font];
}

void NSTextView_inst_SetFont(void *id, void* value) {
	[(NSTextView*)id
		setFont: value];
}

BOOL NSView_inst_AcceptsFirstMouse(void *id, void* event) {
	return [(NSView*)id
		acceptsFirstMouse: event];
}

void NSView_inst_AddConstraints(void *id, void* constraints) {
	[(NSView*)id
		addConstraints: constraints];
}

void NSView_inst_AddSubview(void *id, void* view) {
	[(NSView*)id
		addSubview: view];
}

void NSView_inst_AddSubviewPositionedRelativeTo(void *id, void* view, unsigned long place, void* otherView) {
	[(NSView*)id
		addSubview: view
		positioned: place
		relativeTo: otherView];
}

NSRect NSView_inst_AdjustScroll(void *id, NSRect newVisible) {
	return [(NSView*)id
		adjustScroll: newVisible];
}

NSRect NSView_inst_AlignmentRectForFrame(void *id, NSRect frame) {
	return [(NSView*)id
		alignmentRectForFrame: frame];
}

void* NSView_inst_AncestorSharedWithView(void *id, void* view) {
	return [(NSView*)id
		ancestorSharedWithView: view];
}

BOOL NSView_inst_Autoscroll(void *id, void* event) {
	return [(NSView*)id
		autoscroll: event];
}

void NSView_inst_BeginDocument(void *id) {
	[(NSView*)id
		beginDocument];
}

void NSView_inst_BeginPageInRectAtPlacement(void *id, NSRect rect, NSPoint location) {
	[(NSView*)id
		beginPageInRect: rect
		atPlacement: location];
}

NSRect NSView_inst_CenterScanRect(void *id, NSRect rect) {
	return [(NSView*)id
		centerScanRect: rect];
}

NSPoint NSView_inst_ConvertPointFromView(void *id, NSPoint point, void* view) {
	return [(NSView*)id
		convertPoint: point
		fromView: view];
}

NSPoint NSView_inst_ConvertPointToView(void *id, NSPoint point, void* view) {
	return [(NSView*)id
		convertPoint: point
		toView: view];
}

NSPoint NSView_inst_ConvertPointFromBacking(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointFromBacking: point];
}

NSPoint NSView_inst_ConvertPointFromLayer(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointFromLayer: point];
}

NSPoint NSView_inst_ConvertPointToBacking(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointToBacking: point];
}

NSPoint NSView_inst_ConvertPointToLayer(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointToLayer: point];
}

NSRect NSView_inst_ConvertRectFromView(void *id, NSRect rect, void* view) {
	return [(NSView*)id
		convertRect: rect
		fromView: view];
}

NSRect NSView_inst_ConvertRectToView(void *id, NSRect rect, void* view) {
	return [(NSView*)id
		convertRect: rect
		toView: view];
}

NSRect NSView_inst_ConvertRectFromBacking(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectFromBacking: rect];
}

NSRect NSView_inst_ConvertRectFromLayer(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectFromLayer: rect];
}

NSRect NSView_inst_ConvertRectToBacking(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectToBacking: rect];
}

NSRect NSView_inst_ConvertRectToLayer(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectToLayer: rect];
}

NSSize NSView_inst_ConvertSizeFromView(void *id, NSSize size, void* view) {
	return [(NSView*)id
		convertSize: size
		fromView: view];
}

NSSize NSView_inst_ConvertSizeToView(void *id, NSSize size, void* view) {
	return [(NSView*)id
		convertSize: size
		toView: view];
}

NSSize NSView_inst_ConvertSizeFromBacking(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeFromBacking: size];
}

NSSize NSView_inst_ConvertSizeFromLayer(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeFromLayer: size];
}

NSSize NSView_inst_ConvertSizeToBacking(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeToBacking: size];
}

NSSize NSView_inst_ConvertSizeToLayer(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeToLayer: size];
}

void* NSView_inst_DataWithEPSInsideRect(void *id, NSRect rect) {
	return [(NSView*)id
		dataWithEPSInsideRect: rect];
}

void* NSView_inst_DataWithPDFInsideRect(void *id, NSRect rect) {
	return [(NSView*)id
		dataWithPDFInsideRect: rect];
}

void NSView_inst_DidAddSubview(void *id, void* subview) {
	[(NSView*)id
		didAddSubview: subview];
}

void NSView_inst_DidCloseMenuWithEvent(void *id, void* menu, void* event) {
	[(NSView*)id
		didCloseMenu: menu
		withEvent: event];
}

void NSView_inst_DiscardCursorRects(void *id) {
	[(NSView*)id
		discardCursorRects];
}

void NSView_inst_Display(void *id) {
	[(NSView*)id
		display];
}

void NSView_inst_DisplayIfNeeded(void *id) {
	[(NSView*)id
		displayIfNeeded];
}

void NSView_inst_DisplayIfNeededIgnoringOpacity(void *id) {
	[(NSView*)id
		displayIfNeededIgnoringOpacity];
}

void NSView_inst_DisplayIfNeededInRect(void *id, NSRect rect) {
	[(NSView*)id
		displayIfNeededInRect: rect];
}

void NSView_inst_DisplayIfNeededInRectIgnoringOpacity(void *id, NSRect rect) {
	[(NSView*)id
		displayIfNeededInRectIgnoringOpacity: rect];
}

void NSView_inst_DisplayRect(void *id, NSRect rect) {
	[(NSView*)id
		displayRect: rect];
}

void NSView_inst_DisplayRectIgnoringOpacity(void *id, NSRect rect) {
	[(NSView*)id
		displayRectIgnoringOpacity: rect];
}

void NSView_inst_DrawFocusRingMask(void *id) {
	[(NSView*)id
		drawFocusRingMask];
}

void NSView_inst_DrawPageBorderWithSize(void *id, NSSize borderSize) {
	[(NSView*)id
		drawPageBorderWithSize: borderSize];
}

void NSView_inst_DrawRect(void *id, NSRect dirtyRect) {
	[(NSView*)id
		drawRect: dirtyRect];
}

void NSView_inst_EndDocument(void *id) {
	[(NSView*)id
		endDocument];
}

void NSView_inst_EndPage(void *id) {
	[(NSView*)id
		endPage];
}

BOOL NSView_inst_EnterFullScreenModeWithOptions(void *id, void* screen, void* options) {
	return [(NSView*)id
		enterFullScreenMode: screen
		withOptions: options];
}

void NSView_inst_ExerciseAmbiguityInLayout(void *id) {
	[(NSView*)id
		exerciseAmbiguityInLayout];
}

void NSView_inst_ExitFullScreenModeWithOptions(void *id, void* options) {
	[(NSView*)id
		exitFullScreenModeWithOptions: options];
}

NSRect NSView_inst_FrameForAlignmentRect(void *id, NSRect alignmentRect) {
	return [(NSView*)id
		frameForAlignmentRect: alignmentRect];
}

void* NSView_inst_HitTest(void *id, NSPoint point) {
	return [(NSView*)id
		hitTest: point];
}

void* NSView_inst_InitWithFrame(void *id, NSRect frameRect) {
	return [(NSView*)id
		initWithFrame: frameRect];
}

void NSView_inst_InvalidateIntrinsicContentSize(void *id) {
	[(NSView*)id
		invalidateIntrinsicContentSize];
}

BOOL NSView_inst_IsDescendantOf(void *id, void* view) {
	return [(NSView*)id
		isDescendantOf: view];
}

void NSView_inst_Layout(void *id) {
	[(NSView*)id
		layout];
}

void NSView_inst_LayoutSubtreeIfNeeded(void *id) {
	[(NSView*)id
		layoutSubtreeIfNeeded];
}

NSPoint NSView_inst_LocationOfPrintRect(void *id, NSRect rect) {
	return [(NSView*)id
		locationOfPrintRect: rect];
}

void* NSView_inst_MakeBackingLayer(void *id) {
	return [(NSView*)id
		makeBackingLayer];
}

void* NSView_inst_MenuForEvent(void *id, void* event) {
	return [(NSView*)id
		menuForEvent: event];
}

BOOL NSView_inst_MouseInRect(void *id, NSPoint point, NSRect rect) {
	return [(NSView*)id
		mouse: point
		inRect: rect];
}

BOOL NSView_inst_NeedsToDrawRect(void *id, NSRect rect) {
	return [(NSView*)id
		needsToDrawRect: rect];
}

void NSView_inst_NoteFocusRingMaskChanged(void *id) {
	[(NSView*)id
		noteFocusRingMaskChanged];
}

BOOL NSView_inst_PerformKeyEquivalent(void *id, void* event) {
	return [(NSView*)id
		performKeyEquivalent: event];
}

void NSView_inst_PrepareContentInRect(void *id, NSRect rect) {
	[(NSView*)id
		prepareContentInRect: rect];
}

void NSView_inst_PrepareForReuse(void *id) {
	[(NSView*)id
		prepareForReuse];
}

void NSView_inst_Print(void *id, void* sender) {
	[(NSView*)id
		print: sender];
}

NSRect NSView_inst_RectForPage(void *id, long page) {
	return [(NSView*)id
		rectForPage: page];
}

NSRect NSView_inst_RectForSmartMagnificationAtPointInRect(void *id, NSPoint location, NSRect visibleRect) {
	return [(NSView*)id
		rectForSmartMagnificationAtPoint: location
		inRect: visibleRect];
}

void NSView_inst_RegisterForDraggedTypes(void *id, void* newTypes) {
	[(NSView*)id
		registerForDraggedTypes: newTypes];
}

void NSView_inst_RemoveAllToolTips(void *id) {
	[(NSView*)id
		removeAllToolTips];
}

void NSView_inst_RemoveConstraints(void *id, void* constraints) {
	[(NSView*)id
		removeConstraints: constraints];
}

void NSView_inst_RemoveFromSuperview(void *id) {
	[(NSView*)id
		removeFromSuperview];
}

void NSView_inst_RemoveFromSuperviewWithoutNeedingDisplay(void *id) {
	[(NSView*)id
		removeFromSuperviewWithoutNeedingDisplay];
}

void NSView_inst_ReplaceSubviewWith(void *id, void* oldView, void* newView) {
	[(NSView*)id
		replaceSubview: oldView
		with: newView];
}

void NSView_inst_ResetCursorRects(void *id) {
	[(NSView*)id
		resetCursorRects];
}

void NSView_inst_ResizeSubviewsWithOldSize(void *id, NSSize oldSize) {
	[(NSView*)id
		resizeSubviewsWithOldSize: oldSize];
}

void NSView_inst_ResizeWithOldSuperviewSize(void *id, NSSize oldSize) {
	[(NSView*)id
		resizeWithOldSuperviewSize: oldSize];
}

void NSView_inst_RotateByAngle(void *id, double angle) {
	[(NSView*)id
		rotateByAngle: angle];
}

void NSView_inst_ScaleUnitSquareToSize(void *id, NSSize newUnitSize) {
	[(NSView*)id
		scaleUnitSquareToSize: newUnitSize];
}

void NSView_inst_ScrollPoint(void *id, NSPoint point) {
	[(NSView*)id
		scrollPoint: point];
}

BOOL NSView_inst_ScrollRectToVisible(void *id, NSRect rect) {
	return [(NSView*)id
		scrollRectToVisible: rect];
}

void NSView_inst_SetBoundsOrigin(void *id, NSPoint newOrigin) {
	[(NSView*)id
		setBoundsOrigin: newOrigin];
}

void NSView_inst_SetBoundsSize(void *id, NSSize newSize) {
	[(NSView*)id
		setBoundsSize: newSize];
}

void NSView_inst_SetFrameOrigin(void *id, NSPoint newOrigin) {
	[(NSView*)id
		setFrameOrigin: newOrigin];
}

void NSView_inst_SetFrameSize(void *id, NSSize newSize) {
	[(NSView*)id
		setFrameSize: newSize];
}

void NSView_inst_SetKeyboardFocusRingNeedsDisplayInRect(void *id, NSRect rect) {
	[(NSView*)id
		setKeyboardFocusRingNeedsDisplayInRect: rect];
}

void NSView_inst_SetNeedsDisplayInRect(void *id, NSRect invalidRect) {
	[(NSView*)id
		setNeedsDisplayInRect: invalidRect];
}

BOOL NSView_inst_ShouldDelayWindowOrderingForEvent(void *id, void* event) {
	return [(NSView*)id
		shouldDelayWindowOrderingForEvent: event];
}

void NSView_inst_ShowDefinitionForAttributedStringAtPoint(void *id, void* attrString, NSPoint textBaselineOrigin) {
	[(NSView*)id
		showDefinitionForAttributedString: attrString
		atPoint: textBaselineOrigin];
}

void NSView_inst_TranslateOriginToPoint(void *id, NSPoint translation) {
	[(NSView*)id
		translateOriginToPoint: translation];
}

void NSView_inst_TranslateRectsNeedingDisplayInRectBy(void *id, NSRect clipRect, NSSize delta) {
	[(NSView*)id
		translateRectsNeedingDisplayInRect: clipRect
		by: delta];
}

void NSView_inst_UnregisterDraggedTypes(void *id) {
	[(NSView*)id
		unregisterDraggedTypes];
}

void NSView_inst_UpdateConstraints(void *id) {
	[(NSView*)id
		updateConstraints];
}

void NSView_inst_UpdateConstraintsForSubtreeIfNeeded(void *id) {
	[(NSView*)id
		updateConstraintsForSubtreeIfNeeded];
}

void NSView_inst_UpdateLayer(void *id) {
	[(NSView*)id
		updateLayer];
}

void NSView_inst_UpdateTrackingAreas(void *id) {
	[(NSView*)id
		updateTrackingAreas];
}

void NSView_inst_ViewDidChangeBackingProperties(void *id) {
	[(NSView*)id
		viewDidChangeBackingProperties];
}

void NSView_inst_ViewDidChangeEffectiveAppearance(void *id) {
	[(NSView*)id
		viewDidChangeEffectiveAppearance];
}

void NSView_inst_ViewDidEndLiveResize(void *id) {
	[(NSView*)id
		viewDidEndLiveResize];
}

void NSView_inst_ViewDidHide(void *id) {
	[(NSView*)id
		viewDidHide];
}

void NSView_inst_ViewDidMoveToSuperview(void *id) {
	[(NSView*)id
		viewDidMoveToSuperview];
}

void NSView_inst_ViewDidMoveToWindow(void *id) {
	[(NSView*)id
		viewDidMoveToWindow];
}

void NSView_inst_ViewDidUnhide(void *id) {
	[(NSView*)id
		viewDidUnhide];
}

void NSView_inst_ViewWillDraw(void *id) {
	[(NSView*)id
		viewWillDraw];
}

void NSView_inst_ViewWillMoveToSuperview(void *id, void* newSuperview) {
	[(NSView*)id
		viewWillMoveToSuperview: newSuperview];
}

void NSView_inst_ViewWillMoveToWindow(void *id, void* newWindow) {
	[(NSView*)id
		viewWillMoveToWindow: newWindow];
}

void NSView_inst_ViewWillStartLiveResize(void *id) {
	[(NSView*)id
		viewWillStartLiveResize];
}

void* NSView_inst_ViewWithTag(void *id, long tag) {
	return [(NSView*)id
		viewWithTag: tag];
}

void NSView_inst_WillOpenMenuWithEvent(void *id, void* menu, void* event) {
	[(NSView*)id
		willOpenMenu: menu
		withEvent: event];
}

void NSView_inst_WillRemoveSubview(void *id, void* subview) {
	[(NSView*)id
		willRemoveSubview: subview];
}

void NSView_inst_WriteEPSInsideRectToPasteboard(void *id, NSRect rect, void* pasteboard) {
	[(NSView*)id
		writeEPSInsideRect: rect
		toPasteboard: pasteboard];
}

void NSView_inst_WritePDFInsideRectToPasteboard(void *id, NSRect rect, void* pasteboard) {
	[(NSView*)id
		writePDFInsideRect: rect
		toPasteboard: pasteboard];
}

void* NSView_inst_Init(void *id) {
	return [(NSView*)id
		init];
}

void* NSView_inst_Superview(void *id) {
	return [(NSView*)id
		superview];
}

void* NSView_inst_Subviews(void *id) {
	return [(NSView*)id
		subviews];
}

void NSView_inst_SetSubviews(void *id, void* value) {
	[(NSView*)id
		setSubviews: value];
}

void* NSView_inst_Window(void *id) {
	return [(NSView*)id
		window];
}

void* NSView_inst_OpaqueAncestor(void *id) {
	return [(NSView*)id
		opaqueAncestor];
}

void* NSView_inst_EnclosingMenuItem(void *id) {
	return [(NSView*)id
		enclosingMenuItem];
}

NSRect NSView_inst_Frame(void *id) {
	return [(NSView*)id
		frame];
}

void NSView_inst_SetFrame(void *id, NSRect value) {
	[(NSView*)id
		setFrame: value];
}

double NSView_inst_FrameRotation(void *id) {
	return [(NSView*)id
		frameRotation];
}

void NSView_inst_SetFrameRotation(void *id, double value) {
	[(NSView*)id
		setFrameRotation: value];
}

NSRect NSView_inst_Bounds(void *id) {
	return [(NSView*)id
		bounds];
}

void NSView_inst_SetBounds(void *id, NSRect value) {
	[(NSView*)id
		setBounds: value];
}

double NSView_inst_BoundsRotation(void *id) {
	return [(NSView*)id
		boundsRotation];
}

void NSView_inst_SetBoundsRotation(void *id, double value) {
	[(NSView*)id
		setBoundsRotation: value];
}

BOOL NSView_inst_WantsLayer(void *id) {
	return [(NSView*)id
		wantsLayer];
}

void NSView_inst_SetWantsLayer(void *id, BOOL value) {
	[(NSView*)id
		setWantsLayer: value];
}

BOOL NSView_inst_WantsUpdateLayer(void *id) {
	return [(NSView*)id
		wantsUpdateLayer];
}

void* NSView_inst_Layer(void *id) {
	return [(NSView*)id
		layer];
}

void NSView_inst_SetLayer(void *id, void* value) {
	[(NSView*)id
		setLayer: value];
}

BOOL NSView_inst_CanDrawSubviewsIntoLayer(void *id) {
	return [(NSView*)id
		canDrawSubviewsIntoLayer];
}

void NSView_inst_SetCanDrawSubviewsIntoLayer(void *id, BOOL value) {
	[(NSView*)id
		setCanDrawSubviewsIntoLayer: value];
}

BOOL NSView_inst_LayerUsesCoreImageFilters(void *id) {
	return [(NSView*)id
		layerUsesCoreImageFilters];
}

void NSView_inst_SetLayerUsesCoreImageFilters(void *id, BOOL value) {
	[(NSView*)id
		setLayerUsesCoreImageFilters: value];
}

double NSView_inst_AlphaValue(void *id) {
	return [(NSView*)id
		alphaValue];
}

void NSView_inst_SetAlphaValue(void *id, double value) {
	[(NSView*)id
		setAlphaValue: value];
}

double NSView_inst_FrameCenterRotation(void *id) {
	return [(NSView*)id
		frameCenterRotation];
}

void NSView_inst_SetFrameCenterRotation(void *id, double value) {
	[(NSView*)id
		setFrameCenterRotation: value];
}

void* NSView_inst_BackgroundFilters(void *id) {
	return [(NSView*)id
		backgroundFilters];
}

void NSView_inst_SetBackgroundFilters(void *id, void* value) {
	[(NSView*)id
		setBackgroundFilters: value];
}

void* NSView_inst_ContentFilters(void *id) {
	return [(NSView*)id
		contentFilters];
}

void NSView_inst_SetContentFilters(void *id, void* value) {
	[(NSView*)id
		setContentFilters: value];
}

BOOL NSView_inst_CanDrawConcurrently(void *id) {
	return [(NSView*)id
		canDrawConcurrently];
}

void NSView_inst_SetCanDrawConcurrently(void *id, BOOL value) {
	[(NSView*)id
		setCanDrawConcurrently: value];
}

NSRect NSView_inst_VisibleRect(void *id) {
	return [(NSView*)id
		visibleRect];
}

BOOL NSView_inst_WantsDefaultClipping(void *id) {
	return [(NSView*)id
		wantsDefaultClipping];
}

void* NSView_inst_PrintJobTitle(void *id) {
	return [(NSView*)id
		printJobTitle];
}

void* NSView_inst_PageHeader(void *id) {
	return [(NSView*)id
		pageHeader];
}

void* NSView_inst_PageFooter(void *id) {
	return [(NSView*)id
		pageFooter];
}

double NSView_inst_HeightAdjustLimit(void *id) {
	return [(NSView*)id
		heightAdjustLimit];
}

double NSView_inst_WidthAdjustLimit(void *id) {
	return [(NSView*)id
		widthAdjustLimit];
}

BOOL NSView_inst_NeedsDisplay(void *id) {
	return [(NSView*)id
		needsDisplay];
}

void NSView_inst_SetNeedsDisplay(void *id, BOOL value) {
	[(NSView*)id
		setNeedsDisplay: value];
}

BOOL NSView_inst_IsOpaque(void *id) {
	return [(NSView*)id
		isOpaque];
}

BOOL NSView_inst_IsFlipped(void *id) {
	return [(NSView*)id
		isFlipped];
}

BOOL NSView_inst_IsRotatedFromBase(void *id) {
	return [(NSView*)id
		isRotatedFromBase];
}

BOOL NSView_inst_IsRotatedOrScaledFromBase(void *id) {
	return [(NSView*)id
		isRotatedOrScaledFromBase];
}

BOOL NSView_inst_AutoresizesSubviews(void *id) {
	return [(NSView*)id
		autoresizesSubviews];
}

void NSView_inst_SetAutoresizesSubviews(void *id, BOOL value) {
	[(NSView*)id
		setAutoresizesSubviews: value];
}

void* NSView_inst_Constraints(void *id) {
	return [(NSView*)id
		constraints];
}

void* NSView_inst_LayoutGuides(void *id) {
	return [(NSView*)id
		layoutGuides];
}

NSSize NSView_inst_FittingSize(void *id) {
	return [(NSView*)id
		fittingSize];
}

NSSize NSView_inst_IntrinsicContentSize(void *id) {
	return [(NSView*)id
		intrinsicContentSize];
}

double NSView_inst_BaselineOffsetFromBottom(void *id) {
	return [(NSView*)id
		baselineOffsetFromBottom];
}

double NSView_inst_FirstBaselineOffsetFromTop(void *id) {
	return [(NSView*)id
		firstBaselineOffsetFromTop];
}

double NSView_inst_LastBaselineOffsetFromBottom(void *id) {
	return [(NSView*)id
		lastBaselineOffsetFromBottom];
}

BOOL NSView_inst_NeedsLayout(void *id) {
	return [(NSView*)id
		needsLayout];
}

void NSView_inst_SetNeedsLayout(void *id, BOOL value) {
	[(NSView*)id
		setNeedsLayout: value];
}

BOOL NSView_inst_NeedsUpdateConstraints(void *id) {
	return [(NSView*)id
		needsUpdateConstraints];
}

void NSView_inst_SetNeedsUpdateConstraints(void *id, BOOL value) {
	[(NSView*)id
		setNeedsUpdateConstraints: value];
}

BOOL NSView_inst_TranslatesAutoresizingMaskIntoConstraints(void *id) {
	return [(NSView*)id
		translatesAutoresizingMaskIntoConstraints];
}

void NSView_inst_SetTranslatesAutoresizingMaskIntoConstraints(void *id, BOOL value) {
	[(NSView*)id
		setTranslatesAutoresizingMaskIntoConstraints: value];
}

BOOL NSView_inst_HasAmbiguousLayout(void *id) {
	return [(NSView*)id
		hasAmbiguousLayout];
}

NSRect NSView_inst_FocusRingMaskBounds(void *id) {
	return [(NSView*)id
		focusRingMaskBounds];
}

BOOL NSView_inst_AllowsVibrancy(void *id) {
	return [(NSView*)id
		allowsVibrancy];
}

BOOL NSView_inst_IsInFullScreenMode(void *id) {
	return [(NSView*)id
		isInFullScreenMode];
}

BOOL NSView_inst_IsHidden(void *id) {
	return [(NSView*)id
		isHidden];
}

void NSView_inst_SetHidden(void *id, BOOL value) {
	[(NSView*)id
		setHidden: value];
}

BOOL NSView_inst_IsHiddenOrHasHiddenAncestor(void *id) {
	return [(NSView*)id
		isHiddenOrHasHiddenAncestor];
}

BOOL NSView_inst_InLiveResize(void *id) {
	return [(NSView*)id
		inLiveResize];
}

BOOL NSView_inst_PreservesContentDuringLiveResize(void *id) {
	return [(NSView*)id
		preservesContentDuringLiveResize];
}

NSRect NSView_inst_RectPreservedDuringLiveResize(void *id) {
	return [(NSView*)id
		rectPreservedDuringLiveResize];
}

void* NSView_inst_GestureRecognizers(void *id) {
	return [(NSView*)id
		gestureRecognizers];
}

void NSView_inst_SetGestureRecognizers(void *id, void* value) {
	[(NSView*)id
		setGestureRecognizers: value];
}

BOOL NSView_inst_MouseDownCanMoveWindow(void *id) {
	return [(NSView*)id
		mouseDownCanMoveWindow];
}

BOOL NSView_inst_WantsRestingTouches(void *id) {
	return [(NSView*)id
		wantsRestingTouches];
}

void NSView_inst_SetWantsRestingTouches(void *id, BOOL value) {
	[(NSView*)id
		setWantsRestingTouches: value];
}

BOOL NSView_inst_CanBecomeKeyView(void *id) {
	return [(NSView*)id
		canBecomeKeyView];
}

BOOL NSView_inst_NeedsPanelToBecomeKey(void *id) {
	return [(NSView*)id
		needsPanelToBecomeKey];
}

void* NSView_inst_NextKeyView(void *id) {
	return [(NSView*)id
		nextKeyView];
}

void NSView_inst_SetNextKeyView(void *id, void* value) {
	[(NSView*)id
		setNextKeyView: value];
}

void* NSView_inst_NextValidKeyView(void *id) {
	return [(NSView*)id
		nextValidKeyView];
}

void* NSView_inst_PreviousKeyView(void *id) {
	return [(NSView*)id
		previousKeyView];
}

void* NSView_inst_PreviousValidKeyView(void *id) {
	return [(NSView*)id
		previousValidKeyView];
}

NSRect NSView_inst_PreparedContentRect(void *id) {
	return [(NSView*)id
		preparedContentRect];
}

void NSView_inst_SetPreparedContentRect(void *id, NSRect value) {
	[(NSView*)id
		setPreparedContentRect: value];
}

void* NSView_inst_RegisteredDraggedTypes(void *id) {
	return [(NSView*)id
		registeredDraggedTypes];
}

BOOL NSView_inst_PostsFrameChangedNotifications(void *id) {
	return [(NSView*)id
		postsFrameChangedNotifications];
}

void NSView_inst_SetPostsFrameChangedNotifications(void *id, BOOL value) {
	[(NSView*)id
		setPostsFrameChangedNotifications: value];
}

BOOL NSView_inst_PostsBoundsChangedNotifications(void *id) {
	return [(NSView*)id
		postsBoundsChangedNotifications];
}

void NSView_inst_SetPostsBoundsChangedNotifications(void *id, BOOL value) {
	[(NSView*)id
		setPostsBoundsChangedNotifications: value];
}

long NSView_inst_Tag(void *id) {
	return [(NSView*)id
		tag];
}

void* NSView_inst_ToolTip(void *id) {
	return [(NSView*)id
		toolTip];
}

void NSView_inst_SetToolTip(void *id, void* value) {
	[(NSView*)id
		setToolTip: value];
}

void* NSView_inst_TrackingAreas(void *id) {
	return [(NSView*)id
		trackingAreas];
}

BOOL NSView_inst_IsDrawingFindIndicator(void *id) {
	return [(NSView*)id
		isDrawingFindIndicator];
}

BOOL NSView_inst_IsHorizontalContentSizeConstraintActive(void *id) {
	return [(NSView*)id
		isHorizontalContentSizeConstraintActive];
}

void NSView_inst_SetHorizontalContentSizeConstraintActive(void *id, BOOL value) {
	[(NSView*)id
		setHorizontalContentSizeConstraintActive: value];
}

BOOL NSView_inst_IsVerticalContentSizeConstraintActive(void *id) {
	return [(NSView*)id
		isVerticalContentSizeConstraintActive];
}

void NSView_inst_SetVerticalContentSizeConstraintActive(void *id, BOOL value) {
	[(NSView*)id
		setVerticalContentSizeConstraintActive: value];
}

void* NSView_inst_BackgroundColor(void *id) {
	return [(NSView*)id
		backgroundColor];
}

void NSView_inst_SetBackgroundColor(void *id, void* value) {
	[(NSView*)id
		setBackgroundColor: value];
}


BOOL cocoa_objc_bool_true = YES;
BOOL cocoa_objc_bool_false = NO;

*/
import "C"

func convertObjCBoolToGo(b C.BOOL) bool {
	// NOTE: the prefix here is used to namespace these since the linker will
	// otherwise report a "duplicate symbol" because the C functions have the
	// same name.
	return bool(C.cocoa_convertObjCBool(b))
}

func convertToObjCBool(b bool) C.BOOL {
	if b {
		return C.cocoa_objc_bool_true
	}
	return C.cocoa_objc_bool_false
}

// NSBundle_Alloc is undocumented.
func NSBundle_Alloc() NSBundle {
	ret := C.NSBundle_type_Alloc()

	return NSBundle_FromPointer(ret)
}

// NSBundle_BundleWithURL returns an NSBundle object that corresponds to the specified file URL.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1494992-bundlewithurl?language=objc for details.
func NSBundle_BundleWithURL(url core.NSURLRef) NSBundle {
	ret := C.NSBundle_type_BundleWithURL(
		objc.RefPointer(url),
	)

	return NSBundle_FromPointer(ret)
}

// NSBundle_BundleWithPath returns an NSBundle object that corresponds to the specified directory.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1495012-bundlewithpath?language=objc for details.
func NSBundle_BundleWithPath(path core.NSStringRef) NSBundle {
	ret := C.NSBundle_type_BundleWithPath(
		objc.RefPointer(path),
	)

	return NSBundle_FromPointer(ret)
}

// NSBundle_BundleWithIdentifier returns the NSBundle instance that has the specified bundle identifier.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1411929-bundlewithidentifier?language=objc for details.
func NSBundle_BundleWithIdentifier(identifier core.NSStringRef) NSBundle {
	ret := C.NSBundle_type_BundleWithIdentifier(
		objc.RefPointer(identifier),
	)

	return NSBundle_FromPointer(ret)
}

// NSBundle_URLForResourceWithExtensionSubdirectoryInBundleWithURL creates and returns a file URL for the resource with the specified name and extension in the specified bundle.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1416361-urlforresource?language=objc for details.
func NSBundle_URLForResourceWithExtensionSubdirectoryInBundleWithURL(name core.NSStringRef, ext core.NSStringRef, subpath core.NSStringRef, bundleURL core.NSURLRef) core.NSURL {
	ret := C.NSBundle_type_URLForResourceWithExtensionSubdirectoryInBundleWithURL(
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(bundleURL),
	)

	return core.NSURL_FromPointer(ret)
}

// NSBundle_URLsForResourcesWithExtensionSubdirectoryInBundleWithURL returns an array containing the file URLs for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, within the specified bundle.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1409807-urlsforresourceswithextension?language=objc for details.
func NSBundle_URLsForResourcesWithExtensionSubdirectoryInBundleWithURL(ext core.NSStringRef, subpath core.NSStringRef, bundleURL core.NSURLRef) core.NSArray {
	ret := C.NSBundle_type_URLsForResourcesWithExtensionSubdirectoryInBundleWithURL(
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(bundleURL),
	)

	return core.NSArray_FromPointer(ret)
}

// NSBundle_PathForResourceOfTypeInDirectory returns the full pathname for the resource file identified by the specified name and extension and residing in a given bundle directory.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1409523-pathforresource?language=objc for details.
func NSBundle_PathForResourceOfTypeInDirectory(name core.NSStringRef, ext core.NSStringRef, bundlePath core.NSStringRef) core.NSString {
	ret := C.NSBundle_type_PathForResourceOfTypeInDirectory(
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(bundlePath),
	)

	return core.NSString_FromPointer(ret)
}

// NSBundle_PathsForResourcesOfTypeInDirectory returns an array containing the pathnames for all bundle resources having the specified extension and residing in the bundle directory at the specified path.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1415876-pathsforresourcesoftype?language=objc for details.
func NSBundle_PathsForResourcesOfTypeInDirectory(ext core.NSStringRef, bundlePath core.NSStringRef) core.NSArray {
	ret := C.NSBundle_type_PathsForResourcesOfTypeInDirectory(
		objc.RefPointer(ext),
		objc.RefPointer(bundlePath),
	)

	return core.NSArray_FromPointer(ret)
}

// NSBundle_PreferredLocalizationsFromArray returns one or more localizations from the specified list that a bundle object would use to locate resources for the current user.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1417249-preferredlocalizationsfromarray?language=objc for details.
func NSBundle_PreferredLocalizationsFromArray(localizationsArray core.NSArrayRef) core.NSArray {
	ret := C.NSBundle_type_PreferredLocalizationsFromArray(
		objc.RefPointer(localizationsArray),
	)

	return core.NSArray_FromPointer(ret)
}

// NSBundle_PreferredLocalizationsFromArrayForPreferences returns locale identifiers for which a bundle would provide localized content, given a specified list of candidates for a user's language preferences.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1409418-preferredlocalizationsfromarray?language=objc for details.
func NSBundle_PreferredLocalizationsFromArrayForPreferences(localizationsArray core.NSArrayRef, preferencesArray core.NSArrayRef) core.NSArray {
	ret := C.NSBundle_type_PreferredLocalizationsFromArrayForPreferences(
		objc.RefPointer(localizationsArray),
		objc.RefPointer(preferencesArray),
	)

	return core.NSArray_FromPointer(ret)
}

// NSBundle_MainBundle returns the bundle object that contains the current executable.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1410786-mainbundle?language=objc for details.
func NSBundle_MainBundle() NSBundle {
	ret := C.NSBundle_type_MainBundle()

	return NSBundle_FromPointer(ret)
}

// NSBundle_AllFrameworks returns an array of all of the application’s bundles that represent frameworks.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1408056-allframeworks?language=objc for details.
func NSBundle_AllFrameworks() core.NSArray {
	ret := C.NSBundle_type_AllFrameworks()

	return core.NSArray_FromPointer(ret)
}

// NSBundle_AllBundles returns an array of all the application’s non-framework bundles.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1413705-allbundles?language=objc for details.
func NSBundle_AllBundles() core.NSArray {
	ret := C.NSBundle_type_AllBundles()

	return core.NSArray_FromPointer(ret)
}

// NSSound_Alloc is undocumented.
func NSSound_Alloc() NSSound {
	ret := C.NSSound_type_Alloc()

	return NSSound_FromPointer(ret)
}

// NSSound_CanInitWithPasteboard indicates whether the receiver can create an instance of itself from the data in a pasteboard.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477276-caninitwithpasteboard?language=objc for details.
func NSSound_CanInitWithPasteboard(pasteboard NSPasteboardRef) bool {
	ret := C.NSSound_type_CanInitWithPasteboard(
		objc.RefPointer(pasteboard),
	)

	return convertObjCBoolToGo(ret)
}

// NSSound_SoundUnfilteredTypes provides the file types the NSSound class understands.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477290-soundunfilteredtypes?language=objc for details.
func NSSound_SoundUnfilteredTypes() core.NSArray {
	ret := C.NSSound_type_SoundUnfilteredTypes()

	return core.NSArray_FromPointer(ret)
}

// NSApplication_Alloc is undocumented.
func NSApplication_Alloc() NSApplication {
	ret := C.NSApplication_type_Alloc()

	return NSApplication_FromPointer(ret)
}

// NSApplication_DetachDrawingThreadToTargetWithObject creates and executes a new thread based on the specified target and selector.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428374-detachdrawingthread?language=objc for details.
func NSApplication_DetachDrawingThreadToTargetWithObject(selector objc.Selector, target objc.Ref, argument objc.Ref) {
	C.NSApplication_type_DetachDrawingThreadToTargetWithObject(
		selector.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(argument),
	)

	return
}

// NSApplication_SharedApplication returns the application instance, creating it if it doesn’t exist yet.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428360-sharedapplication?language=objc for details.
func NSApplication_SharedApplication() NSApplication {
	ret := C.NSApplication_type_SharedApplication()

	return NSApplication_FromPointer(ret)
}

// NSControl_Alloc is undocumented.
func NSControl_Alloc() NSControl {
	ret := C.NSControl_type_Alloc()

	return NSControl_FromPointer(ret)
}

// NSButton_Alloc is undocumented.
func NSButton_Alloc() NSButton {
	ret := C.NSButton_type_Alloc()

	return NSButton_FromPointer(ret)
}

// NSButton_CheckboxWithTitleTargetAction creates a standard checkbox with the title you specify.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1644525-checkboxwithtitle?language=objc for details.
func NSButton_CheckboxWithTitleTargetAction(title core.NSStringRef, target objc.Ref, action objc.Selector) NSButton {
	ret := C.NSButton_type_CheckboxWithTitleTargetAction(
		objc.RefPointer(title),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)

	return NSButton_FromPointer(ret)
}

// NSButton_ButtonWithImageTargetAction creates a standard push button with the image you specify.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1644659-buttonwithimage?language=objc for details.
func NSButton_ButtonWithImageTargetAction(image NSImageRef, target objc.Ref, action objc.Selector) NSButton {
	ret := C.NSButton_type_ButtonWithImageTargetAction(
		objc.RefPointer(image),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)

	return NSButton_FromPointer(ret)
}

// NSButton_RadioButtonWithTitleTargetAction creates a standard radio button with the title you specify.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1644340-radiobuttonwithtitle?language=objc for details.
func NSButton_RadioButtonWithTitleTargetAction(title core.NSStringRef, target objc.Ref, action objc.Selector) NSButton {
	ret := C.NSButton_type_RadioButtonWithTitleTargetAction(
		objc.RefPointer(title),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)

	return NSButton_FromPointer(ret)
}

// NSButton_ButtonWithTitleImageTargetAction creates a standard push button with a title and image.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1644719-buttonwithtitle?language=objc for details.
func NSButton_ButtonWithTitleImageTargetAction(title core.NSStringRef, image NSImageRef, target objc.Ref, action objc.Selector) NSButton {
	ret := C.NSButton_type_ButtonWithTitleImageTargetAction(
		objc.RefPointer(title),
		objc.RefPointer(image),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)

	return NSButton_FromPointer(ret)
}

// NSButton_ButtonWithTitleTargetAction creates a standard push button with the title you specify.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1644256-buttonwithtitle?language=objc for details.
func NSButton_ButtonWithTitleTargetAction(title core.NSStringRef, target objc.Ref, action objc.Selector) NSButton {
	ret := C.NSButton_type_ButtonWithTitleTargetAction(
		objc.RefPointer(title),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)

	return NSButton_FromPointer(ret)
}

// NSEvent_Alloc is undocumented.
func NSEvent_Alloc() NSEvent {
	ret := C.NSEvent_type_Alloc()

	return NSEvent_FromPointer(ret)
}

// NSEvent_EventWithEventRef creates an event object that is based on a Carbon type of event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1528021-eventwitheventref?language=objc for details.
func NSEvent_EventWithEventRef(eventRef unsafe.Pointer) NSEvent {
	ret := C.NSEvent_type_EventWithEventRef(
		eventRef,
	)

	return NSEvent_FromPointer(ret)
}

// NSEvent_StopPeriodicEvents stops generating periodic events for the current thread and discards any periodic events remaining in the queue.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1533746-stopperiodicevents?language=objc for details.
func NSEvent_StopPeriodicEvents() {
	C.NSEvent_type_StopPeriodicEvents()

	return
}

// NSEvent_RemoveMonitor remove the specified event monitor.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1533709-removemonitor?language=objc for details.
func NSEvent_RemoveMonitor(eventMonitor objc.Ref) {
	C.NSEvent_type_RemoveMonitor(
		objc.RefPointer(eventMonitor),
	)

	return
}

// NSEvent_PressedMouseButtons returns the indices of the currently depressed mouse buttons.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1527943-pressedmousebuttons?language=objc for details.
func NSEvent_PressedMouseButtons() core.NSUInteger {
	ret := C.NSEvent_type_PressedMouseButtons()

	return core.NSUInteger(ret)
}

// NSEvent_MouseLocation reports the current mouse position in screen coordinates.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1533380-mouselocation?language=objc for details.
func NSEvent_MouseLocation() core.NSPoint {
	ret := C.NSEvent_type_MouseLocation()

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// NSEvent_MouseCoalescingEnabled is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsevent/2870068-mousecoalescingenabled?language=objc for details.
func NSEvent_MouseCoalescingEnabled() bool {
	ret := C.NSEvent_type_MouseCoalescingEnabled()

	return convertObjCBoolToGo(ret)
}

// NSEvent_SetMouseCoalescingEnabled is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsevent/2870068-mousecoalescingenabled?language=objc for details.
func NSEvent_SetMouseCoalescingEnabled(value bool) {
	C.NSEvent_type_SetMouseCoalescingEnabled(
		convertToObjCBool(value),
	)

	return
}

// NSEvent_SwipeTrackingFromScrollEventsEnabled is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsevent/2870067-swipetrackingfromscrolleventsena?language=objc for details.
func NSEvent_SwipeTrackingFromScrollEventsEnabled() bool {
	ret := C.NSEvent_type_SwipeTrackingFromScrollEventsEnabled()

	return convertObjCBoolToGo(ret)
}

// NSFont_Alloc is undocumented.
func NSFont_Alloc() NSFont {
	ret := C.NSFont_type_Alloc()

	return NSFont_FromPointer(ret)
}

// NSFont_FontWithNameSize creates a font object for the specified font name and font size.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1525977-fontwithname?language=objc for details.
func NSFont_FontWithNameSize(fontName core.NSStringRef, fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_FontWithNameSize(
		objc.RefPointer(fontName),
		C.double(fontSize),
	)

	return NSFont_FromPointer(ret)
}

// NSFont_UserFontOfSize returns the font used by default for documents and other text under the user’s control (that is, text whose font the user can normally change), in the specified size.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1524559-userfontofsize?language=objc for details.
func NSFont_UserFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_UserFontOfSize(
		C.double(fontSize),
	)

	return NSFont_FromPointer(ret)
}

// NSFont_UserFixedPitchFontOfSize returns the font used by default for documents and other text under the user’s control (that is, text whose font the user can normally change), when that font should be fixed-pitch, in the specified size.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1531381-userfixedpitchfontofsize?language=objc for details.
func NSFont_UserFixedPitchFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_UserFixedPitchFontOfSize(
		C.double(fontSize),
	)

	return NSFont_FromPointer(ret)
}

// NSFont_SystemFontOfSize returns the standard system font with the specified size.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1530094-systemfontofsize?language=objc for details.
func NSFont_SystemFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_SystemFontOfSize(
		C.double(fontSize),
	)

	return NSFont_FromPointer(ret)
}

// NSFont_BoldSystemFontOfSize returns the standard system font in boldface type with the specified size.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1533549-boldsystemfontofsize?language=objc for details.
func NSFont_BoldSystemFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_BoldSystemFontOfSize(
		C.double(fontSize),
	)

	return NSFont_FromPointer(ret)
}

// NSFont_LabelFontOfSize returns the font used for standard interface labels in the specified size.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1528213-labelfontofsize?language=objc for details.
func NSFont_LabelFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_LabelFontOfSize(
		C.double(fontSize),
	)

	return NSFont_FromPointer(ret)
}

// NSFont_MessageFontOfSize returns the font used for standard interface items, such as button labels, menu items, and so on, in the specified size.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1525777-messagefontofsize?language=objc for details.
func NSFont_MessageFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_MessageFontOfSize(
		C.double(fontSize),
	)

	return NSFont_FromPointer(ret)
}

// NSFont_MenuBarFontOfSize returns the font used for menu bar items, in the specified size.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1534194-menubarfontofsize?language=objc for details.
func NSFont_MenuBarFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_MenuBarFontOfSize(
		C.double(fontSize),
	)

	return NSFont_FromPointer(ret)
}

// NSFont_MenuFontOfSize returns the font used for menu items, in the specified size.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1533068-menufontofsize?language=objc for details.
func NSFont_MenuFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_MenuFontOfSize(
		C.double(fontSize),
	)

	return NSFont_FromPointer(ret)
}

// NSFont_ControlContentFontOfSize returns the font used for the content of controls in the specified size.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1527070-controlcontentfontofsize?language=objc for details.
func NSFont_ControlContentFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_ControlContentFontOfSize(
		C.double(fontSize),
	)

	return NSFont_FromPointer(ret)
}

// NSFont_TitleBarFontOfSize returns the font used for window title bars, in the specified size.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1530200-titlebarfontofsize?language=objc for details.
func NSFont_TitleBarFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_TitleBarFontOfSize(
		C.double(fontSize),
	)

	return NSFont_FromPointer(ret)
}

// NSFont_PaletteFontOfSize returns the font used for palette window title bars, in the specified size.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1535462-palettefontofsize?language=objc for details.
func NSFont_PaletteFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_PaletteFontOfSize(
		C.double(fontSize),
	)

	return NSFont_FromPointer(ret)
}

// NSFont_ToolTipsFontOfSize returns the font used for tool tips labels, in the specified size.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1527704-tooltipsfontofsize?language=objc for details.
func NSFont_ToolTipsFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_ToolTipsFontOfSize(
		C.double(fontSize),
	)

	return NSFont_FromPointer(ret)
}

// NSFont_SetUserFont sets the font used by default for documents and other text under the user’s control to the specified font.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1526068-setuserfont?language=objc for details.
func NSFont_SetUserFont(font NSFontRef) {
	C.NSFont_type_SetUserFont(
		objc.RefPointer(font),
	)

	return
}

// NSFont_SetUserFixedPitchFont sets the font used by default for documents and other text under the user’s control, when that font should be fixed-pitch, to the specified font.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1529050-setuserfixedpitchfont?language=objc for details.
func NSFont_SetUserFixedPitchFont(font NSFontRef) {
	C.NSFont_type_SetUserFixedPitchFont(
		objc.RefPointer(font),
	)

	return
}

// NSFont_SystemFontSize returns the size of the standard system font.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1531931-systemfontsize?language=objc for details.
func NSFont_SystemFontSize() core.CGFloat {
	ret := C.NSFont_type_SystemFontSize()

	return core.CGFloat(ret)
}

// NSFont_SmallSystemFontSize returns the size of the standard small system font.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1535612-smallsystemfontsize?language=objc for details.
func NSFont_SmallSystemFontSize() core.CGFloat {
	ret := C.NSFont_type_SmallSystemFontSize()

	return core.CGFloat(ret)
}

// NSFont_LabelFontSize returns the size of the standard label font.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1534629-labelfontsize?language=objc for details.
func NSFont_LabelFontSize() core.CGFloat {
	ret := C.NSFont_type_LabelFontSize()

	return core.CGFloat(ret)
}

// NSImage_Alloc is undocumented.
func NSImage_Alloc() NSImage {
	ret := C.NSImage_type_Alloc()

	return NSImage_FromPointer(ret)
}

// NSImage_ImageWithSystemSymbolNameAccessibilityDescription creates a symbol image with the system symbol name and accessibility description that you specify.
//
// See https://developer.apple.com/documentation/appkit/nsimage/3622472-imagewithsystemsymbolname?language=objc for details.
func NSImage_ImageWithSystemSymbolNameAccessibilityDescription(symbolName core.NSStringRef, description core.NSStringRef) NSImage {
	ret := C.NSImage_type_ImageWithSystemSymbolNameAccessibilityDescription(
		objc.RefPointer(symbolName),
		objc.RefPointer(description),
	)

	return NSImage_FromPointer(ret)
}

// NSImage_CanInitWithPasteboard tests whether the image can create an instance of itself using pasteboard data.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1520039-caninitwithpasteboard?language=objc for details.
func NSImage_CanInitWithPasteboard(pasteboard NSPasteboardRef) bool {
	ret := C.NSImage_type_CanInitWithPasteboard(
		objc.RefPointer(pasteboard),
	)

	return convertObjCBoolToGo(ret)
}

// NSImage_ImageTypes returns an array of UTI strings identifying the image types supported by the registered image representation objects, either directly or through a user-installed filter service.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519988-imagetypes?language=objc for details.
func NSImage_ImageTypes() core.NSArray {
	ret := C.NSImage_type_ImageTypes()

	return core.NSArray_FromPointer(ret)
}

// NSImage_ImageUnfilteredTypes returns an array of UTI strings identifying the image types supported directly by the registered image representation objects.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519899-imageunfilteredtypes?language=objc for details.
func NSImage_ImageUnfilteredTypes() core.NSArray {
	ret := C.NSImage_type_ImageUnfilteredTypes()

	return core.NSArray_FromPointer(ret)
}

// NSImageView_Alloc is undocumented.
func NSImageView_Alloc() NSImageView {
	ret := C.NSImageView_type_Alloc()

	return NSImageView_FromPointer(ret)
}

// NSImageView_ImageViewWithImage is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsimageview/1644708-imageviewwithimage?language=objc for details.
func NSImageView_ImageViewWithImage(image NSImageRef) NSImageView {
	ret := C.NSImageView_type_ImageViewWithImage(
		objc.RefPointer(image),
	)

	return NSImageView_FromPointer(ret)
}

// NSNib_Alloc is undocumented.
func NSNib_Alloc() NSNib {
	ret := C.NSNib_type_Alloc()

	return NSNib_FromPointer(ret)
}

// NSPasteboard_Alloc is undocumented.
func NSPasteboard_Alloc() NSPasteboard {
	ret := C.NSPasteboard_type_Alloc()

	return NSPasteboard_FromPointer(ret)
}

// NSPasteboard_PasteboardByFilteringFile creates a new pasteboard object that supplies the specified file in as many types as possible based on the available filter services.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1532744-pasteboardbyfilteringfile?language=objc for details.
func NSPasteboard_PasteboardByFilteringFile(filename core.NSStringRef) NSPasteboard {
	ret := C.NSPasteboard_type_PasteboardByFilteringFile(
		objc.RefPointer(filename),
	)

	return NSPasteboard_FromPointer(ret)
}

// NSPasteboard_PasteboardByFilteringTypesInPasteboard creates a new pasteboard object that supplies the specified pasteboard data in as many types as possible based on the available filter services.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1530088-pasteboardbyfilteringtypesinpast?language=objc for details.
func NSPasteboard_PasteboardByFilteringTypesInPasteboard(pboard NSPasteboardRef) NSPasteboard {
	ret := C.NSPasteboard_type_PasteboardByFilteringTypesInPasteboard(
		objc.RefPointer(pboard),
	)

	return NSPasteboard_FromPointer(ret)
}

// NSPasteboard_PasteboardWithUniqueName creates and returns a new pasteboard with a name that is guaranteed to be unique with respect to other pasteboards in the system.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1528936-pasteboardwithuniquename?language=objc for details.
func NSPasteboard_PasteboardWithUniqueName() NSPasteboard {
	ret := C.NSPasteboard_type_PasteboardWithUniqueName()

	return NSPasteboard_FromPointer(ret)
}

// NSPasteboard_GeneralPasteboard returns the shared pasteboard object to use for general content.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1530091-generalpasteboard?language=objc for details.
func NSPasteboard_GeneralPasteboard() NSPasteboard {
	ret := C.NSPasteboard_type_GeneralPasteboard()

	return NSPasteboard_FromPointer(ret)
}

// NSLayoutManager_Alloc is undocumented.
func NSLayoutManager_Alloc() NSLayoutManager {
	ret := C.NSLayoutManager_type_Alloc()

	return NSLayoutManager_FromPointer(ret)
}

// NSMenu_Alloc is undocumented.
func NSMenu_Alloc() NSMenu {
	ret := C.NSMenu_type_Alloc()

	return NSMenu_FromPointer(ret)
}

// NSMenu_MenuBarVisible returns a Boolean value that indicates whether the menu bar is visible.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518236-menubarvisible?language=objc for details.
func NSMenu_MenuBarVisible() bool {
	ret := C.NSMenu_type_MenuBarVisible()

	return convertObjCBoolToGo(ret)
}

// NSMenu_SetMenuBarVisible sets whether the menu bar is visible and selectable by the user.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518200-setmenubarvisible?language=objc for details.
func NSMenu_SetMenuBarVisible(visible bool) {
	C.NSMenu_type_SetMenuBarVisible(
		convertToObjCBool(visible),
	)

	return
}

// NSMenu_PopUpContextMenuWithEventForView displays a contextual menu over a view for an event.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518170-popupcontextmenu?language=objc for details.
func NSMenu_PopUpContextMenuWithEventForView(menu NSMenuRef, event NSEventRef, view NSViewRef) {
	C.NSMenu_type_PopUpContextMenuWithEventForView(
		objc.RefPointer(menu),
		objc.RefPointer(event),
		objc.RefPointer(view),
	)

	return
}

// NSMenu_PopUpContextMenuWithEventForViewWithFont displays a contextual menu over a view for an event using a specified font.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518165-popupcontextmenu?language=objc for details.
func NSMenu_PopUpContextMenuWithEventForViewWithFont(menu NSMenuRef, event NSEventRef, view NSViewRef, font NSFontRef) {
	C.NSMenu_type_PopUpContextMenuWithEventForViewWithFont(
		objc.RefPointer(menu),
		objc.RefPointer(event),
		objc.RefPointer(view),
		objc.RefPointer(font),
	)

	return
}

// NSPopover_Alloc is undocumented.
func NSPopover_Alloc() NSPopover {
	ret := C.NSPopover_type_Alloc()

	return NSPopover_FromPointer(ret)
}

// NSMenuItem_Alloc is undocumented.
func NSMenuItem_Alloc() NSMenuItem {
	ret := C.NSMenuItem_type_Alloc()

	return NSMenuItem_FromPointer(ret)
}

// NSMenuItem_SeparatorItem returns a menu item that is used to separate logical groups of menu commands.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514838-separatoritem?language=objc for details.
func NSMenuItem_SeparatorItem() NSMenuItem {
	ret := C.NSMenuItem_type_SeparatorItem()

	return NSMenuItem_FromPointer(ret)
}

// NSMenuItem_UsesUserKeyEquivalents returns a Boolean value that indicates whether menu items conform to user preferences for key equivalents.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514811-usesuserkeyequivalents?language=objc for details.
func NSMenuItem_UsesUserKeyEquivalents() bool {
	ret := C.NSMenuItem_type_UsesUserKeyEquivalents()

	return convertObjCBoolToGo(ret)
}

// NSMenuItem_SetUsesUserKeyEquivalents returns a Boolean value that indicates whether menu items conform to user preferences for key equivalents.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514811-usesuserkeyequivalents?language=objc for details.
func NSMenuItem_SetUsesUserKeyEquivalents(value bool) {
	C.NSMenuItem_type_SetUsesUserKeyEquivalents(
		convertToObjCBool(value),
	)

	return
}

// NSRunningApplication_Alloc is undocumented.
func NSRunningApplication_Alloc() NSRunningApplication {
	ret := C.NSRunningApplication_type_Alloc()

	return NSRunningApplication_FromPointer(ret)
}

// NSRunningApplication_RunningApplicationsWithBundleIdentifier returns an array of currently running applications with the specified bundle identifier.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1530798-runningapplicationswithbundleide?language=objc for details.
func NSRunningApplication_RunningApplicationsWithBundleIdentifier(bundleIdentifier core.NSStringRef) core.NSArray {
	ret := C.NSRunningApplication_type_RunningApplicationsWithBundleIdentifier(
		objc.RefPointer(bundleIdentifier),
	)

	return core.NSArray_FromPointer(ret)
}

// NSRunningApplication_TerminateAutomaticallyTerminableApplications terminates invisibly running applications as if triggered by system memory pressure.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1529538-terminateautomaticallyterminable?language=objc for details.
func NSRunningApplication_TerminateAutomaticallyTerminableApplications() {
	C.NSRunningApplication_type_TerminateAutomaticallyTerminableApplications()

	return
}

// NSRunningApplication_CurrentApplication returns an NSRunningApplication representing this application.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1533604-currentapplication?language=objc for details.
func NSRunningApplication_CurrentApplication() NSRunningApplication {
	ret := C.NSRunningApplication_type_CurrentApplication()

	return NSRunningApplication_FromPointer(ret)
}

// NSScreen_Alloc is undocumented.
func NSScreen_Alloc() NSScreen {
	ret := C.NSScreen_type_Alloc()

	return NSScreen_FromPointer(ret)
}

// NSScreen_MainScreen returns the screen object containing the window with the keyboard focus.
//
// See https://developer.apple.com/documentation/appkit/nsscreen/1388371-mainscreen?language=objc for details.
func NSScreen_MainScreen() NSScreen {
	ret := C.NSScreen_type_MainScreen()

	return NSScreen_FromPointer(ret)
}

// NSScreen_DeepestScreen returns a screen object representing the screen that can best represent color.
//
// See https://developer.apple.com/documentation/appkit/nsscreen/1388374-deepestscreen?language=objc for details.
func NSScreen_DeepestScreen() NSScreen {
	ret := C.NSScreen_type_DeepestScreen()

	return NSScreen_FromPointer(ret)
}

// NSScreen_Screens returns an array of screen objects representing all of the screens available on the system.
//
// See https://developer.apple.com/documentation/appkit/nsscreen/1388393-screens?language=objc for details.
func NSScreen_Screens() core.NSArray {
	ret := C.NSScreen_type_Screens()

	return core.NSArray_FromPointer(ret)
}

// NSScreen_ScreensHaveSeparateSpaces returns a Boolean value indicating whether each screen can have its own set of spaces.
//
// See https://developer.apple.com/documentation/appkit/nsscreen/1388365-screenshaveseparatespaces?language=objc for details.
func NSScreen_ScreensHaveSeparateSpaces() bool {
	ret := C.NSScreen_type_ScreensHaveSeparateSpaces()

	return convertObjCBoolToGo(ret)
}

// NSStatusBar_Alloc is undocumented.
func NSStatusBar_Alloc() NSStatusBar {
	ret := C.NSStatusBar_type_Alloc()

	return NSStatusBar_FromPointer(ret)
}

// NSStatusBar_SystemStatusBar returns the system-wide status bar located in the menu bar.
//
// See https://developer.apple.com/documentation/appkit/nsstatusbar/1530619-systemstatusbar?language=objc for details.
func NSStatusBar_SystemStatusBar() NSStatusBar {
	ret := C.NSStatusBar_type_SystemStatusBar()

	return NSStatusBar_FromPointer(ret)
}

// NSStatusBarButton_Alloc is undocumented.
func NSStatusBarButton_Alloc() NSStatusBarButton {
	ret := C.NSStatusBarButton_type_Alloc()

	return NSStatusBarButton_FromPointer(ret)
}

// NSStatusItem_Alloc is undocumented.
func NSStatusItem_Alloc() NSStatusItem {
	ret := C.NSStatusItem_type_Alloc()

	return NSStatusItem_FromPointer(ret)
}

// NSText_Alloc is undocumented.
func NSText_Alloc() NSText {
	ret := C.NSText_type_Alloc()

	return NSText_FromPointer(ret)
}

// NSTextField_Alloc is undocumented.
func NSTextField_Alloc() NSTextField {
	ret := C.NSTextField_type_Alloc()

	return NSTextField_FromPointer(ret)
}

// NSTextField_LabelWithAttributedString creates a text field for use as a static label that displays styled text, doesn’t wrap, and doesn’t have selectable text.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1644658-labelwithattributedstring?language=objc for details.
func NSTextField_LabelWithAttributedString(attributedStringValue core.NSAttributedStringRef) NSTextField {
	ret := C.NSTextField_type_LabelWithAttributedString(
		objc.RefPointer(attributedStringValue),
	)

	return NSTextField_FromPointer(ret)
}

// NSTextField_LabelWithString initializes a text field for use as a static label that uses the system default font, doesn’t wrap, and doesn’t have selectable text.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1644377-labelwithstring?language=objc for details.
func NSTextField_LabelWithString(stringValue core.NSStringRef) NSTextField {
	ret := C.NSTextField_type_LabelWithString(
		objc.RefPointer(stringValue),
	)

	return NSTextField_FromPointer(ret)
}

// NSTextField_TextFieldWithString initializes a single-line editable text field for user input using the system default font and standard visual appearance.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1644706-textfieldwithstring?language=objc for details.
func NSTextField_TextFieldWithString(stringValue core.NSStringRef) NSTextField {
	ret := C.NSTextField_type_TextFieldWithString(
		objc.RefPointer(stringValue),
	)

	return NSTextField_FromPointer(ret)
}

// NSTextField_WrappingLabelWithString initializes a text field for use as a multiline static label with selectable text that uses the system default font.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1644543-wrappinglabelwithstring?language=objc for details.
func NSTextField_WrappingLabelWithString(stringValue core.NSStringRef) NSTextField {
	ret := C.NSTextField_type_WrappingLabelWithString(
		objc.RefPointer(stringValue),
	)

	return NSTextField_FromPointer(ret)
}

// NSTextContainer_Alloc is undocumented.
func NSTextContainer_Alloc() NSTextContainer {
	ret := C.NSTextContainer_type_Alloc()

	return NSTextContainer_FromPointer(ret)
}

// NSViewController_Alloc is undocumented.
func NSViewController_Alloc() NSViewController {
	ret := C.NSViewController_type_Alloc()

	return NSViewController_FromPointer(ret)
}

// NSVisualEffectView_Alloc is undocumented.
func NSVisualEffectView_Alloc() NSVisualEffectView {
	ret := C.NSVisualEffectView_type_Alloc()

	return NSVisualEffectView_FromPointer(ret)
}

// NSWindow_Alloc is undocumented.
func NSWindow_Alloc() NSWindow {
	ret := C.NSWindow_type_Alloc()

	return NSWindow_FromPointer(ret)
}

// NSWindow_WindowWithContentViewController creates a titled window that contains the specified content view controller.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419551-windowwithcontentviewcontroller?language=objc for details.
func NSWindow_WindowWithContentViewController(contentViewController NSViewControllerRef) NSWindow {
	ret := C.NSWindow_type_WindowWithContentViewController(
		objc.RefPointer(contentViewController),
	)

	return NSWindow_FromPointer(ret)
}

// NSWindow_ContentRectForFrameRectStyleMask returns the content rectangle used by a window with a given frame rectangle and window style.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419586-contentrectforframerect?language=objc for details.
func NSWindow_ContentRectForFrameRectStyleMask(fRect core.NSRect, style core.NSUInteger) core.NSRect {
	ret := C.NSWindow_type_ContentRectForFrameRectStyleMask(
		*(*C.NSRect)(unsafe.Pointer(&fRect)),
		C.ulong(style),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// NSWindow_FrameRectForContentRectStyleMask returns the frame rectangle used by a window with a given content rectangle and window style.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419372-framerectforcontentrect?language=objc for details.
func NSWindow_FrameRectForContentRectStyleMask(cRect core.NSRect, style core.NSUInteger) core.NSRect {
	ret := C.NSWindow_type_FrameRectForContentRectStyleMask(
		*(*C.NSRect)(unsafe.Pointer(&cRect)),
		C.ulong(style),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// NSWindow_MinFrameWidthWithTitleStyleMask returns the minimum width a window’s frame rectangle must have for it to display a title, with a given window style.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419294-minframewidthwithtitle?language=objc for details.
func NSWindow_MinFrameWidthWithTitleStyleMask(title core.NSStringRef, style core.NSUInteger) core.CGFloat {
	ret := C.NSWindow_type_MinFrameWidthWithTitleStyleMask(
		objc.RefPointer(title),
		C.ulong(style),
	)

	return core.CGFloat(ret)
}

// NSWindow_WindowNumberAtPointBelowWindowWithWindowNumber returns the number of the frontmost window that would be hit by a mouse-down at the specified screen location.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419210-windownumberatpoint?language=objc for details.
func NSWindow_WindowNumberAtPointBelowWindowWithWindowNumber(point core.NSPoint, windowNumber core.NSInteger) core.NSInteger {
	ret := C.NSWindow_type_WindowNumberAtPointBelowWindowWithWindowNumber(
		*(*C.NSPoint)(unsafe.Pointer(&point)),
		C.long(windowNumber),
	)

	return core.NSInteger(ret)
}

// NSWindow_AllowsAutomaticWindowTabbing returns a Boolean value that indicates whether the app can automatically organize windows into tabs.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1646657-allowsautomaticwindowtabbing?language=objc for details.
func NSWindow_AllowsAutomaticWindowTabbing() bool {
	ret := C.NSWindow_type_AllowsAutomaticWindowTabbing()

	return convertObjCBoolToGo(ret)
}

// NSWindow_SetAllowsAutomaticWindowTabbing returns a Boolean value that indicates whether the app can automatically organize windows into tabs.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1646657-allowsautomaticwindowtabbing?language=objc for details.
func NSWindow_SetAllowsAutomaticWindowTabbing(value bool) {
	C.NSWindow_type_SetAllowsAutomaticWindowTabbing(
		convertToObjCBool(value),
	)

	return
}

// NSWorkspace_Alloc is undocumented.
func NSWorkspace_Alloc() NSWorkspace {
	ret := C.NSWorkspace_type_Alloc()

	return NSWorkspace_FromPointer(ret)
}

// NSWorkspace_SharedWorkspace returns the shared workspace object.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1530344-sharedworkspace?language=objc for details.
func NSWorkspace_SharedWorkspace() NSWorkspace {
	ret := C.NSWorkspace_type_SharedWorkspace()

	return NSWorkspace_FromPointer(ret)
}

// NSColor_Alloc is undocumented.
func NSColor_Alloc() NSColor {
	ret := C.NSColor_type_Alloc()

	return NSColor_FromPointer(ret)
}

// NSColor_ColorFromPasteboard creates a color object from color data currently on the pasteboard.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1535057-colorfrompasteboard?language=objc for details.
func NSColor_ColorFromPasteboard(pasteBoard NSPasteboardRef) NSColor {
	ret := C.NSColor_type_ColorFromPasteboard(
		objc.RefPointer(pasteBoard),
	)

	return NSColor_FromPointer(ret)
}

// NSColor_ColorWithRedGreenBlueAlpha is undocumented.
func NSColor_ColorWithRedGreenBlueAlpha(red core.CGFloat, green core.CGFloat, blue core.CGFloat, alpha core.CGFloat) NSColor {
	ret := C.NSColor_type_ColorWithRedGreenBlueAlpha(
		C.double(red),
		C.double(green),
		C.double(blue),
		C.double(alpha),
	)

	return NSColor_FromPointer(ret)
}

// NSColor_IgnoresAlpha returns a Boolean value that indicates whether the app supports alpha.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1533565-ignoresalpha?language=objc for details.
func NSColor_IgnoresAlpha() bool {
	ret := C.NSColor_type_IgnoresAlpha()

	return convertObjCBoolToGo(ret)
}

// NSColor_SetIgnoresAlpha returns a Boolean value that indicates whether the app supports alpha.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1533565-ignoresalpha?language=objc for details.
func NSColor_SetIgnoresAlpha(value bool) {
	C.NSColor_type_SetIgnoresAlpha(
		convertToObjCBool(value),
	)

	return
}

// NSColor_SystemCyanColor is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nscolor/3816005-systemcyancolor?language=objc for details.
func NSColor_SystemCyanColor() NSColor {
	ret := C.NSColor_type_SystemCyanColor()

	return NSColor_FromPointer(ret)
}

// NSColor_SystemMintColor is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nscolor/3816006-systemmintcolor?language=objc for details.
func NSColor_SystemMintColor() NSColor {
	ret := C.NSColor_type_SystemMintColor()

	return NSColor_FromPointer(ret)
}

// NSColor_ClearColor returns a color object whose grayscale and alpha values are both 0.0.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1527217-clearcolor?language=objc for details.
func NSColor_ClearColor() NSColor {
	ret := C.NSColor_type_ClearColor()

	return NSColor_FromPointer(ret)
}

// NSTextView_Alloc is undocumented.
func NSTextView_Alloc() NSTextView {
	ret := C.NSTextView_type_Alloc()

	return NSTextView_FromPointer(ret)
}

// NSTextView_RegisterForServices registers send and return types for the Services facility.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449507-registerforservices?language=objc for details.
func NSTextView_RegisterForServices() {
	C.NSTextView_type_RegisterForServices()

	return
}

// NSTextView_FieldEditor is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nstextview/2990525-fieldeditor?language=objc for details.
func NSTextView_FieldEditor() NSTextView {
	ret := C.NSTextView_type_FieldEditor()

	return NSTextView_FromPointer(ret)
}

// NSTextView_StronglyReferencesTextStorage is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nstextview/2269433-stronglyreferencestextstorage?language=objc for details.
func NSTextView_StronglyReferencesTextStorage() bool {
	ret := C.NSTextView_type_StronglyReferencesTextStorage()

	return convertObjCBoolToGo(ret)
}

// NSView_Alloc is undocumented.
func NSView_Alloc() NSView {
	ret := C.NSView_type_Alloc()

	return NSView_FromPointer(ret)
}

// NSView_RequiresConstraintBasedLayout returns a Boolean value indicating whether the view depends on the constraint-based layout system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526926-requiresconstraintbasedlayout?language=objc for details.
func NSView_RequiresConstraintBasedLayout() bool {
	ret := C.NSView_type_RequiresConstraintBasedLayout()

	return convertObjCBoolToGo(ret)
}

// NSView_FocusView returns the currently focused NSView object, or nil if there is none.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483662-focusview?language=objc for details.
func NSView_FocusView() NSView {
	ret := C.NSView_type_FocusView()

	return NSView_FromPointer(ret)
}

// NSView_DefaultMenu overridden by subclasses to return the default pop-up menu for instances of the receiving class.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483417-defaultmenu?language=objc for details.
func NSView_DefaultMenu() NSMenu {
	ret := C.NSView_type_DefaultMenu()

	return NSMenu_FromPointer(ret)
}

// NSView_CompatibleWithResponsiveScrolling is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsview/2870005-compatiblewithresponsivescrollin?language=objc for details.
func NSView_CompatibleWithResponsiveScrolling() bool {
	ret := C.NSView_type_CompatibleWithResponsiveScrolling()

	return convertObjCBoolToGo(ret)
}

type NSBundleRef interface {
	Pointer() uintptr
	Init() NSBundle
}

type gen_NSBundle struct {
	objc.Object
}

func NSBundle_FromPointer(ptr unsafe.Pointer) NSBundle {
	return NSBundle{gen_NSBundle{
		objc.Object_FromPointer(ptr),
	}}
}

func NSBundle_FromRef(ref objc.Ref) NSBundle {
	return NSBundle_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// URLForAuxiliaryExecutable returns the file URL of the executable with the specified name in the receiver’s bundle.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1411412-urlforauxiliaryexecutable?language=objc for details.
func (x gen_NSBundle) URLForAuxiliaryExecutable(
	executableName core.NSStringRef,
) core.NSURL {
	ret := C.NSBundle_inst_URLForAuxiliaryExecutable(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(executableName),
	)

	return core.NSURL_FromPointer(ret)
}

// URLForResourceWithExtension returns the file URL for the resource identified by the specified name and file extension.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1411540-urlforresource?language=objc for details.
func (x gen_NSBundle) URLForResourceWithExtension(
	name core.NSStringRef,
	ext core.NSStringRef,
) core.NSURL {
	ret := C.NSBundle_inst_URLForResourceWithExtension(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
	)

	return core.NSURL_FromPointer(ret)
}

// URLForResourceWithExtensionSubdirectory returns the file URL for the resource file identified by the specified name and extension and residing in a given bundle directory.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1416712-urlforresource?language=objc for details.
func (x gen_NSBundle) URLForResourceWithExtensionSubdirectory(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
) core.NSURL {
	ret := C.NSBundle_inst_URLForResourceWithExtensionSubdirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)

	return core.NSURL_FromPointer(ret)
}

// URLForResourceWithExtensionSubdirectoryLocalization returns the file URL for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1417378-urlforresource?language=objc for details.
func (x gen_NSBundle) URLForResourceWithExtensionSubdirectoryLocalization(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) core.NSURL {
	ret := C.NSBundle_inst_URLForResourceWithExtensionSubdirectoryLocalization(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)

	return core.NSURL_FromPointer(ret)
}

// URLsForResourcesWithExtensionSubdirectory returns an array of file URLs for all resources identified by the specified file extension and located in the specified bundle subdirectory.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1407424-urlsforresourceswithextension?language=objc for details.
func (x gen_NSBundle) URLsForResourcesWithExtensionSubdirectory(
	ext core.NSStringRef,
	subpath core.NSStringRef,
) core.NSArray {
	ret := C.NSBundle_inst_URLsForResourcesWithExtensionSubdirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)

	return core.NSArray_FromPointer(ret)
}

// URLsForResourcesWithExtensionSubdirectoryLocalization returns an array containing the file URLs for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1414688-urlsforresourceswithextension?language=objc for details.
func (x gen_NSBundle) URLsForResourcesWithExtensionSubdirectoryLocalization(
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) core.NSArray {
	ret := C.NSBundle_inst_URLsForResourcesWithExtensionSubdirectoryLocalization(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)

	return core.NSArray_FromPointer(ret)
}

// InitWithPath returns an NSBundle object initialized to correspond to the specified directory.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1412741-initwithpath?language=objc for details.
func (x gen_NSBundle) InitWithPath(
	path core.NSStringRef,
) NSBundle {
	ret := C.NSBundle_inst_InitWithPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)

	return NSBundle_FromPointer(ret)
}

// InitWithURL returns an NSBundle object initialized to correspond to the specified file URL.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1409352-initwithurl?language=objc for details.
func (x gen_NSBundle) InitWithURL(
	url core.NSURLRef,
) NSBundle {
	ret := C.NSBundle_inst_InitWithURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)

	return NSBundle_FromPointer(ret)
}

// Load dynamically loads the bundle’s executable code into a running program, if the code has not already been loaded.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1415927-load?language=objc for details.
func (x gen_NSBundle) Load() bool {
	ret := C.NSBundle_inst_Load(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// LoadNibNamedOwnerOptions unarchives the contents of a nib file located in the receiver's bundle.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1618147-loadnibnamed?language=objc for details.
func (x gen_NSBundle) LoadNibNamedOwnerOptions(
	name core.NSStringRef,
	owner objc.Ref,
	options core.NSDictionaryRef,
) core.NSArray {
	ret := C.NSBundle_inst_LoadNibNamedOwnerOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(owner),
		objc.RefPointer(options),
	)

	return core.NSArray_FromPointer(ret)
}

// LocalizedAttributedStringForKeyValueTable is undocumented.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/3746904-localizedattributedstringforkey?language=objc for details.
func (x gen_NSBundle) LocalizedAttributedStringForKeyValueTable(
	key core.NSStringRef,
	value core.NSStringRef,
	tableName core.NSStringRef,
) core.NSAttributedString {
	ret := C.NSBundle_inst_LocalizedAttributedStringForKeyValueTable(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
		objc.RefPointer(value),
		objc.RefPointer(tableName),
	)

	return core.NSAttributedString_FromPointer(ret)
}

// LocalizedStringForKeyValueTable returns a localized version of the string designated by the specified key and residing in the specified table.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1417694-localizedstringforkey?language=objc for details.
func (x gen_NSBundle) LocalizedStringForKeyValueTable(
	key core.NSStringRef,
	value core.NSStringRef,
	tableName core.NSStringRef,
) core.NSString {
	ret := C.NSBundle_inst_LocalizedStringForKeyValueTable(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
		objc.RefPointer(value),
		objc.RefPointer(tableName),
	)

	return core.NSString_FromPointer(ret)
}

// ObjectForInfoDictionaryKey returns the value associated with the specified key in the receiver's information property list.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1408696-objectforinfodictionarykey?language=objc for details.
func (x gen_NSBundle) ObjectForInfoDictionaryKey(
	key core.NSStringRef,
) objc.Object {
	ret := C.NSBundle_inst_ObjectForInfoDictionaryKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)

	return objc.Object_FromPointer(ret)
}

// PathForAuxiliaryExecutable returns the full pathname of the executable with the specified name in the receiver’s bundle.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1415214-pathforauxiliaryexecutable?language=objc for details.
func (x gen_NSBundle) PathForAuxiliaryExecutable(
	executableName core.NSStringRef,
) core.NSString {
	ret := C.NSBundle_inst_PathForAuxiliaryExecutable(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(executableName),
	)

	return core.NSString_FromPointer(ret)
}

// PathForResourceOfType returns the full pathname for the resource identified by the specified name and file extension.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1410989-pathforresource?language=objc for details.
func (x gen_NSBundle) PathForResourceOfType(
	name core.NSStringRef,
	ext core.NSStringRef,
) core.NSString {
	ret := C.NSBundle_inst_PathForResourceOfType(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
	)

	return core.NSString_FromPointer(ret)
}

// PathForResourceOfTypeInDirectory returns the full pathname for the resource identified by the specified name and file extension and located in the specified bundle subdirectory.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1409670-pathforresource?language=objc for details.
func (x gen_NSBundle) PathForResourceOfTypeInDirectory(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
) core.NSString {
	ret := C.NSBundle_inst_PathForResourceOfTypeInDirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)

	return core.NSString_FromPointer(ret)
}

// PathForResourceOfTypeInDirectoryForLocalization returns the full pathname for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1413471-pathforresource?language=objc for details.
func (x gen_NSBundle) PathForResourceOfTypeInDirectoryForLocalization(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) core.NSString {
	ret := C.NSBundle_inst_PathForResourceOfTypeInDirectoryForLocalization(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)

	return core.NSString_FromPointer(ret)
}

// PathsForResourcesOfTypeInDirectory returns an array containing the pathnames for all bundle resources having the specified filename extension and residing in the resource subdirectory.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1413058-pathsforresourcesoftype?language=objc for details.
func (x gen_NSBundle) PathsForResourcesOfTypeInDirectory(
	ext core.NSStringRef,
	subpath core.NSStringRef,
) core.NSArray {
	ret := C.NSBundle_inst_PathsForResourcesOfTypeInDirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)

	return core.NSArray_FromPointer(ret)
}

// PathsForResourcesOfTypeInDirectoryForLocalization returns an array containing the file for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1416940-pathsforresourcesoftype?language=objc for details.
func (x gen_NSBundle) PathsForResourcesOfTypeInDirectoryForLocalization(
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) core.NSArray {
	ret := C.NSBundle_inst_PathsForResourcesOfTypeInDirectoryForLocalization(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)

	return core.NSArray_FromPointer(ret)
}

// Unload unloads the code associated with the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1412388-unload?language=objc for details.
func (x gen_NSBundle) Unload() bool {
	ret := C.NSBundle_inst_Unload(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Init is undocumented.
func (x gen_NSBundle) Init() NSBundle {
	ret := C.NSBundle_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSBundle_FromPointer(ret)
}

// ResourceURL returns the file URL of the bundle’s subdirectory containing resource files.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1414821-resourceurl?language=objc for details.
func (x gen_NSBundle) ResourceURL() core.NSURL {
	ret := C.NSBundle_inst_ResourceURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_FromPointer(ret)
}

// ExecutableURL returns the file URL of the receiver's executable file.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1410470-executableurl?language=objc for details.
func (x gen_NSBundle) ExecutableURL() core.NSURL {
	ret := C.NSBundle_inst_ExecutableURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_FromPointer(ret)
}

// PrivateFrameworksURL returns the file URL of the bundle’s subdirectory containing private frameworks.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1417617-privateframeworksurl?language=objc for details.
func (x gen_NSBundle) PrivateFrameworksURL() core.NSURL {
	ret := C.NSBundle_inst_PrivateFrameworksURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_FromPointer(ret)
}

// SharedFrameworksURL returns the file URL of the receiver's subdirectory containing shared frameworks.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1411774-sharedframeworksurl?language=objc for details.
func (x gen_NSBundle) SharedFrameworksURL() core.NSURL {
	ret := C.NSBundle_inst_SharedFrameworksURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_FromPointer(ret)
}

// BuiltInPlugInsURL returns the file URL of the receiver's subdirectory containing plug-ins.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1409603-builtinpluginsurl?language=objc for details.
func (x gen_NSBundle) BuiltInPlugInsURL() core.NSURL {
	ret := C.NSBundle_inst_BuiltInPlugInsURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_FromPointer(ret)
}

// SharedSupportURL returns the file URL of the bundle’s subdirectory containing shared support files.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1416823-sharedsupporturl?language=objc for details.
func (x gen_NSBundle) SharedSupportURL() core.NSURL {
	ret := C.NSBundle_inst_SharedSupportURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_FromPointer(ret)
}

// AppStoreReceiptURL returns the file URL for the bundle’s App Store receipt.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1407276-appstorereceipturl?language=objc for details.
func (x gen_NSBundle) AppStoreReceiptURL() core.NSURL {
	ret := C.NSBundle_inst_AppStoreReceiptURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_FromPointer(ret)
}

// ResourcePath returns the full pathname of the bundle’s subdirectory containing resources.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1417723-resourcepath?language=objc for details.
func (x gen_NSBundle) ResourcePath() core.NSString {
	ret := C.NSBundle_inst_ResourcePath(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// ExecutablePath returns the full pathname of the receiver's executable file.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1409078-executablepath?language=objc for details.
func (x gen_NSBundle) ExecutablePath() core.NSString {
	ret := C.NSBundle_inst_ExecutablePath(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// PrivateFrameworksPath returns the full pathname of the bundle’s subdirectory containing private frameworks.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1415562-privateframeworkspath?language=objc for details.
func (x gen_NSBundle) PrivateFrameworksPath() core.NSString {
	ret := C.NSBundle_inst_PrivateFrameworksPath(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SharedFrameworksPath returns the full pathname of the bundle’s subdirectory containing shared frameworks.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1417226-sharedframeworkspath?language=objc for details.
func (x gen_NSBundle) SharedFrameworksPath() core.NSString {
	ret := C.NSBundle_inst_SharedFrameworksPath(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// BuiltInPlugInsPath returns the full pathname of the receiver's subdirectory containing plug-ins.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1408900-builtinpluginspath?language=objc for details.
func (x gen_NSBundle) BuiltInPlugInsPath() core.NSString {
	ret := C.NSBundle_inst_BuiltInPlugInsPath(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SharedSupportPath returns the full pathname of the bundle’s subdirectory containing shared support files.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1411609-sharedsupportpath?language=objc for details.
func (x gen_NSBundle) SharedSupportPath() core.NSString {
	ret := C.NSBundle_inst_SharedSupportPath(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// BundleURL returns the full URL of the receiver’s bundle directory.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1415654-bundleurl?language=objc for details.
func (x gen_NSBundle) BundleURL() core.NSURL {
	ret := C.NSBundle_inst_BundleURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_FromPointer(ret)
}

// BundlePath returns the full pathname of the receiver’s bundle directory.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1407973-bundlepath?language=objc for details.
func (x gen_NSBundle) BundlePath() core.NSString {
	ret := C.NSBundle_inst_BundlePath(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// BundleIdentifier returns the receiver’s bundle identifier.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1418023-bundleidentifier?language=objc for details.
func (x gen_NSBundle) BundleIdentifier() core.NSString {
	ret := C.NSBundle_inst_BundleIdentifier(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// InfoDictionary returns a dictionary, constructed from the bundle’s Info.plist file, that contains information about the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1413477-infodictionary?language=objc for details.
func (x gen_NSBundle) InfoDictionary() core.NSDictionary {
	ret := C.NSBundle_inst_InfoDictionary(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_FromPointer(ret)
}

// Localizations returns a list of all the localizations contained in the bundle.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1417415-localizations?language=objc for details.
func (x gen_NSBundle) Localizations() core.NSArray {
	ret := C.NSBundle_inst_Localizations(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// PreferredLocalizations an ordered list of preferred localizations contained in the bundle.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1413220-preferredlocalizations?language=objc for details.
func (x gen_NSBundle) PreferredLocalizations() core.NSArray {
	ret := C.NSBundle_inst_PreferredLocalizations(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// DevelopmentLocalization returns the localization for the development language.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1417526-developmentlocalization?language=objc for details.
func (x gen_NSBundle) DevelopmentLocalization() core.NSString {
	ret := C.NSBundle_inst_DevelopmentLocalization(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// LocalizedInfoDictionary returns a dictionary with the keys from the bundle’s localized property list.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1407645-localizedinfodictionary?language=objc for details.
func (x gen_NSBundle) LocalizedInfoDictionary() core.NSDictionary {
	ret := C.NSBundle_inst_LocalizedInfoDictionary(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_FromPointer(ret)
}

// ExecutableArchitectures an array of numbers indicating the architecture types supported by the bundle’s executable.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1415499-executablearchitectures?language=objc for details.
func (x gen_NSBundle) ExecutableArchitectures() core.NSArray {
	ret := C.NSBundle_inst_ExecutableArchitectures(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// IsLoaded returns the load status of a bundle.
//
// See https://developer.apple.com/documentation/foundation/nsbundle/1413594-loaded?language=objc for details.
func (x gen_NSBundle) IsLoaded() bool {
	ret := C.NSBundle_inst_IsLoaded(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

type NSSoundRef interface {
	Pointer() uintptr
	Init() NSSound
}

type gen_NSSound struct {
	objc.Object
}

func NSSound_FromPointer(ptr unsafe.Pointer) NSSound {
	return NSSound{gen_NSSound{
		objc.Object_FromPointer(ptr),
	}}
}

func NSSound_FromRef(ref objc.Ref) NSSound {
	return NSSound_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// InitWithContentsOfFileByReference initializes the receiver with the audio data located at a given filepath.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477274-initwithcontentsoffile?language=objc for details.
func (x gen_NSSound) InitWithContentsOfFileByReference(
	path core.NSStringRef,
	byRef bool,
) NSSound {
	ret := C.NSSound_inst_InitWithContentsOfFileByReference(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(byRef),
	)

	return NSSound_FromPointer(ret)
}

// InitWithContentsOfURLByReference initializes the receiver with the audio data located at a given URL.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477288-initwithcontentsofurl?language=objc for details.
func (x gen_NSSound) InitWithContentsOfURLByReference(
	url core.NSURLRef,
	byRef bool,
) NSSound {
	ret := C.NSSound_inst_InitWithContentsOfURLByReference(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
		convertToObjCBool(byRef),
	)

	return NSSound_FromPointer(ret)
}

// InitWithData initializes the receiver with a given audio data.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477292-initwithdata?language=objc for details.
func (x gen_NSSound) InitWithData(
	data core.NSDataRef,
) NSSound {
	ret := C.NSSound_inst_InitWithData(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
	)

	return NSSound_FromPointer(ret)
}

// InitWithPasteboard initializes the receiver with data from a pasteboard. The pasteboard should contain a type returned by NSSound. NSSound expects the data to have a proper magic number, sound header, and data for the formats it supports.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477294-initwithpasteboard?language=objc for details.
func (x gen_NSSound) InitWithPasteboard(
	pasteboard NSPasteboardRef,
) NSSound {
	ret := C.NSSound_inst_InitWithPasteboard(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pasteboard),
	)

	return NSSound_FromPointer(ret)
}

// Pause pauses audio playback.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477307-pause?language=objc for details.
func (x gen_NSSound) Pause() bool {
	ret := C.NSSound_inst_Pause(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Play initiates audio playback.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477322-play?language=objc for details.
func (x gen_NSSound) Play() bool {
	ret := C.NSSound_inst_Play(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Resume resumes audio playback.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477336-resume?language=objc for details.
func (x gen_NSSound) Resume() bool {
	ret := C.NSSound_inst_Resume(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Stop concludes audio playback.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477282-stop?language=objc for details.
func (x gen_NSSound) Stop() bool {
	ret := C.NSSound_inst_Stop(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// WriteToPasteboard writes the receiver’s data to a pasteboard.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477338-writetopasteboard?language=objc for details.
func (x gen_NSSound) WriteToPasteboard(
	pasteboard NSPasteboardRef,
) {
	C.NSSound_inst_WriteToPasteboard(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pasteboard),
	)

	return
}

// Init is undocumented.
func (x gen_NSSound) Init() NSSound {
	ret := C.NSSound_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSSound_FromPointer(ret)
}

// Delegate returns the sound’s delegate.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477300-delegate?language=objc for details.
func (x gen_NSSound) Delegate() objc.Object {
	ret := C.NSSound_inst_Delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetDelegate returns the sound’s delegate.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477300-delegate?language=objc for details.
func (x gen_NSSound) SetDelegate(
	value objc.Ref,
) {
	C.NSSound_inst_SetDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Loops returns a Boolean that indicates whether the sound restarts playback when it reaches the end of its content.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477311-loops?language=objc for details.
func (x gen_NSSound) Loops() bool {
	ret := C.NSSound_inst_Loops(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetLoops returns a Boolean that indicates whether the sound restarts playback when it reaches the end of its content.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477311-loops?language=objc for details.
func (x gen_NSSound) SetLoops(
	value bool,
) {
	C.NSSound_inst_SetLoops(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsPlaying returns a Boolean that indicates whether the sound is playing its audio data.
//
// See https://developer.apple.com/documentation/appkit/nssound/1477302-playing?language=objc for details.
func (x gen_NSSound) IsPlaying() bool {
	ret := C.NSSound_inst_IsPlaying(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

type NSApplicationRef interface {
	Pointer() uintptr
	Init() NSApplication
}

type gen_NSApplication struct {
	objc.Object
}

func NSApplication_FromPointer(ptr unsafe.Pointer) NSApplication {
	return NSApplication{gen_NSApplication{
		objc.Object_FromPointer(ptr),
	}}
}

func NSApplication_FromRef(ref objc.Ref) NSApplication {
	return NSApplication_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// ActivateContextHelpMode places the receiver in context-sensitive help mode.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1500925-activatecontexthelpmode?language=objc for details.
func (x gen_NSApplication) ActivateContextHelpMode(
	sender objc.Ref,
) {
	C.NSApplication_inst_ActivateContextHelpMode(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ActivateIgnoringOtherApps makes the receiver the active app.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428468-activateignoringotherapps?language=objc for details.
func (x gen_NSApplication) ActivateIgnoringOtherApps(
	flag bool,
) {
	C.NSApplication_inst_ActivateIgnoringOtherApps(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)

	return
}

// ActivationPolicy returns the app’s activation policy.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428703-activationpolicy?language=objc for details.
func (x gen_NSApplication) ActivationPolicy() core.NSInteger {
	ret := C.NSApplication_inst_ActivationPolicy(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// CancelUserAttentionRequest cancels a previous user attention request.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428683-canceluserattentionrequest?language=objc for details.
func (x gen_NSApplication) CancelUserAttentionRequest(
	request core.NSInteger,
) {
	C.NSApplication_inst_CancelUserAttentionRequest(
		unsafe.Pointer(x.Pointer()),
		C.long(request),
	)

	return
}

// Deactivate deactivates the receiver.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428428-deactivate?language=objc for details.
func (x gen_NSApplication) Deactivate() {
	C.NSApplication_inst_Deactivate(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DisableRelaunchOnLogin disables relaunching the app on login.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428376-disablerelaunchonlogin?language=objc for details.
func (x gen_NSApplication) DisableRelaunchOnLogin() {
	C.NSApplication_inst_DisableRelaunchOnLogin(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// EnableRelaunchOnLogin enables relaunching the app on login.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428453-enablerelaunchonlogin?language=objc for details.
func (x gen_NSApplication) EnableRelaunchOnLogin() {
	C.NSApplication_inst_EnableRelaunchOnLogin(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// FinishLaunching activates the app, opens any files specified by the NSOpen user default, and unhighlights the app’s icon.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428771-finishlaunching?language=objc for details.
func (x gen_NSApplication) FinishLaunching() {
	C.NSApplication_inst_FinishLaunching(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// HideOtherApplications hides all apps, except the receiver.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428746-hideotherapplications?language=objc for details.
func (x gen_NSApplication) HideOtherApplications(
	sender objc.Ref,
) {
	C.NSApplication_inst_HideOtherApplications(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// PostEventAtStart adds a given event to the receiver’s event queue.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428710-postevent?language=objc for details.
func (x gen_NSApplication) PostEventAtStart(
	event NSEventRef,
	flag bool,
) {
	C.NSApplication_inst_PostEventAtStart(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
		convertToObjCBool(flag),
	)

	return
}

// RegisterForRemoteNotifications register for notifications sent by Apple Push Notification service (APNs).
//
// See https://developer.apple.com/documentation/appkit/nsapplication/2967172-registerforremotenotifications?language=objc for details.
func (x gen_NSApplication) RegisterForRemoteNotifications() {
	C.NSApplication_inst_RegisterForRemoteNotifications(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// RegisterUserInterfaceItemSearchHandler register an object that provides help data to your app.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1420818-registeruserinterfaceitemsearchh?language=objc for details.
func (x gen_NSApplication) RegisterUserInterfaceItemSearchHandler(
	handler objc.Ref,
) {
	C.NSApplication_inst_RegisterUserInterfaceItemSearchHandler(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(handler),
	)

	return
}

// ReplyToApplicationShouldTerminate responds to NSTerminateLater once the app knows whether it can terminate.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428594-replytoapplicationshouldterminat?language=objc for details.
func (x gen_NSApplication) ReplyToApplicationShouldTerminate(
	shouldTerminate bool,
) {
	C.NSApplication_inst_ReplyToApplicationShouldTerminate(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(shouldTerminate),
	)

	return
}

// Run starts the main event loop.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428631-run?language=objc for details.
func (x gen_NSApplication) Run() {
	C.NSApplication_inst_Run(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// SendActionToFrom sends the given action message to the given target.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428509-sendaction?language=objc for details.
func (x gen_NSApplication) SendActionToFrom(
	action objc.Selector,
	target objc.Ref,
	sender objc.Ref,
) bool {
	ret := C.NSApplication_inst_SendActionToFrom(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(sender),
	)

	return convertObjCBoolToGo(ret)
}

// SendEvent dispatches an event to other objects.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428359-sendevent?language=objc for details.
func (x gen_NSApplication) SendEvent(
	event NSEventRef,
) {
	C.NSApplication_inst_SendEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return
}

// SetActivationPolicy attempts to modify the app’s activation policy.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428621-setactivationpolicy?language=objc for details.
func (x gen_NSApplication) SetActivationPolicy(
	activationPolicy core.NSInteger,
) bool {
	ret := C.NSApplication_inst_SetActivationPolicy(
		unsafe.Pointer(x.Pointer()),
		C.long(activationPolicy),
	)

	return convertObjCBoolToGo(ret)
}

// ShowHelp if your project is properly registered, and the necessary keys have been set in the property list, this method launches Help Viewer and displays the first page of your app’s help book.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1500910-showhelp?language=objc for details.
func (x gen_NSApplication) ShowHelp(
	sender objc.Ref,
) {
	C.NSApplication_inst_ShowHelp(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// Stop stops the main event loop.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428473-stop?language=objc for details.
func (x gen_NSApplication) Stop(
	sender objc.Ref,
) {
	C.NSApplication_inst_Stop(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// TargetForAction returns the object that receives the action message specified by the given selector.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428449-targetforaction?language=objc for details.
func (x gen_NSApplication) TargetForAction(
	action objc.Selector,
) objc.Object {
	ret := C.NSApplication_inst_TargetForAction(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
	)

	return objc.Object_FromPointer(ret)
}

// TargetForActionToFrom searches for an object that can receive the message specified by the given selector.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428658-targetforaction?language=objc for details.
func (x gen_NSApplication) TargetForActionToFrom(
	action objc.Selector,
	target objc.Ref,
	sender objc.Ref,
) objc.Object {
	ret := C.NSApplication_inst_TargetForActionToFrom(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(sender),
	)

	return objc.Object_FromPointer(ret)
}

// Terminate terminates the receiver.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428417-terminate?language=objc for details.
func (x gen_NSApplication) Terminate(
	sender objc.Ref,
) {
	C.NSApplication_inst_Terminate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleTouchBarCustomizationPalette show or hides the interface for customizing the Touch Bar.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/2646920-toggletouchbarcustomizationpalet?language=objc for details.
func (x gen_NSApplication) ToggleTouchBarCustomizationPalette(
	sender objc.Ref,
) {
	C.NSApplication_inst_ToggleTouchBarCustomizationPalette(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// TryToPerformWith dispatches an action message to the specified target.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428366-trytoperform?language=objc for details.
func (x gen_NSApplication) TryToPerformWith(
	action objc.Selector,
	object objc.Ref,
) bool {
	ret := C.NSApplication_inst_TryToPerformWith(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(object),
	)

	return convertObjCBoolToGo(ret)
}

// UnhideAllApplications unhides all apps, including the receiver.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428737-unhideallapplications?language=objc for details.
func (x gen_NSApplication) UnhideAllApplications(
	sender objc.Ref,
) {
	C.NSApplication_inst_UnhideAllApplications(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// UnregisterForRemoteNotifications unregister for notifications received from Apple Push Notification service.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428747-unregisterforremotenotifications?language=objc for details.
func (x gen_NSApplication) UnregisterForRemoteNotifications() {
	C.NSApplication_inst_UnregisterForRemoteNotifications(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// UnregisterUserInterfaceItemSearchHandler unregister an object that provides help data to your app.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1420820-unregisteruserinterfaceitemsearc?language=objc for details.
func (x gen_NSApplication) UnregisterUserInterfaceItemSearchHandler(
	handler objc.Ref,
) {
	C.NSApplication_inst_UnregisterUserInterfaceItemSearchHandler(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(handler),
	)

	return
}

// Init is undocumented.
func (x gen_NSApplication) Init() NSApplication {
	ret := C.NSApplication_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSApplication_FromPointer(ret)
}

// Delegate returns the app delegate object.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428705-delegate?language=objc for details.
func (x gen_NSApplication) Delegate() objc.Object {
	ret := C.NSApplication_inst_Delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetDelegate returns the app delegate object.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428705-delegate?language=objc for details.
func (x gen_NSApplication) SetDelegate(
	value objc.Ref,
) {
	C.NSApplication_inst_SetDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// CurrentEvent returns the last event object that the app retrieved from the event queue.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428668-currentevent?language=objc for details.
func (x gen_NSApplication) CurrentEvent() NSEvent {
	ret := C.NSApplication_inst_CurrentEvent(
		unsafe.Pointer(x.Pointer()),
	)

	return NSEvent_FromPointer(ret)
}

// IsRunning returns a Boolean value indicating whether the main event loop is running.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428759-running?language=objc for details.
func (x gen_NSApplication) IsRunning() bool {
	ret := C.NSApplication_inst_IsRunning(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsActive returns a Boolean value indicating whether this is the active app.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428493-active?language=objc for details.
func (x gen_NSApplication) IsActive() bool {
	ret := C.NSApplication_inst_IsActive(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsRegisteredForRemoteNotifications returns a Boolean value indicating whether the app is registered with Apple Push Notification service (APNs).
//
// See https://developer.apple.com/documentation/appkit/nsapplication/2967173-registeredforremotenotifications?language=objc for details.
func (x gen_NSApplication) IsRegisteredForRemoteNotifications() bool {
	ret := C.NSApplication_inst_IsRegisteredForRemoteNotifications(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// ApplicationIconImage returns the image used for the app’s icon.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428744-applicationiconimage?language=objc for details.
func (x gen_NSApplication) ApplicationIconImage() NSImage {
	ret := C.NSApplication_inst_ApplicationIconImage(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_FromPointer(ret)
}

// SetApplicationIconImage returns the image used for the app’s icon.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428744-applicationiconimage?language=objc for details.
func (x gen_NSApplication) SetApplicationIconImage(
	value NSImageRef,
) {
	C.NSApplication_inst_SetApplicationIconImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// HelpMenu returns the help menu used by the app.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428644-helpmenu?language=objc for details.
func (x gen_NSApplication) HelpMenu() NSMenu {
	ret := C.NSApplication_inst_HelpMenu(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenu_FromPointer(ret)
}

// SetHelpMenu returns the help menu used by the app.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428644-helpmenu?language=objc for details.
func (x gen_NSApplication) SetHelpMenu(
	value NSMenuRef,
) {
	C.NSApplication_inst_SetHelpMenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// ServicesProvider returns the object that provides the services the current app advertises in the Services menu of other apps.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428467-servicesprovider?language=objc for details.
func (x gen_NSApplication) ServicesProvider() objc.Object {
	ret := C.NSApplication_inst_ServicesProvider(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetServicesProvider returns the object that provides the services the current app advertises in the Services menu of other apps.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428467-servicesprovider?language=objc for details.
func (x gen_NSApplication) SetServicesProvider(
	value objc.Ref,
) {
	C.NSApplication_inst_SetServicesProvider(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// IsFullKeyboardAccessEnabled returns a Boolean value indicating whether Full Keyboard Access is enabled in the Keyboard preference pane.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428469-fullkeyboardaccessenabled?language=objc for details.
func (x gen_NSApplication) IsFullKeyboardAccessEnabled() bool {
	ret := C.NSApplication_inst_IsFullKeyboardAccessEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// OrderedDocuments an array of document objects arranged according to the front-to-back ordering of their associated windows.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1494283-ordereddocuments?language=objc for details.
func (x gen_NSApplication) OrderedDocuments() core.NSArray {
	ret := C.NSApplication_inst_OrderedDocuments(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// OrderedWindows an array of window objects arranged according to their front-to-back ordering on the screen.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1494287-orderedwindows?language=objc for details.
func (x gen_NSApplication) OrderedWindows() core.NSArray {
	ret := C.NSApplication_inst_OrderedWindows(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// MainMenu returns the app’s main menu bar.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428634-mainmenu?language=objc for details.
func (x gen_NSApplication) MainMenu() NSMenu {
	ret := C.NSApplication_inst_MainMenu(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenu_FromPointer(ret)
}

// SetMainMenu returns the app’s main menu bar.
//
// See https://developer.apple.com/documentation/appkit/nsapplication/1428634-mainmenu?language=objc for details.
func (x gen_NSApplication) SetMainMenu(
	value NSMenuRef,
) {
	C.NSApplication_inst_SetMainMenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

type NSControlRef interface {
	Pointer() uintptr
	Init() NSControl
}

type gen_NSControl struct {
	NSView
}

func NSControl_FromPointer(ptr unsafe.Pointer) NSControl {
	return NSControl{gen_NSControl{
		NSView_FromPointer(ptr),
	}}
}

func NSControl_FromRef(ref objc.Ref) NSControl {
	return NSControl_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// AbortEditing terminates the current editing operation and discards any edited text.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428867-abortediting?language=objc for details.
func (x gen_NSControl) AbortEditing() bool {
	ret := C.NSControl_inst_AbortEditing(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// CurrentEditor returns the current field editor for the control.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428980-currenteditor?language=objc for details.
func (x gen_NSControl) CurrentEditor() NSText {
	ret := C.NSControl_inst_CurrentEditor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSText_FromPointer(ret)
}

// DrawWithExpansionFrameInView performs custom expansion tool tip drawing.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428895-drawwithexpansionframe?language=objc for details.
func (x gen_NSControl) DrawWithExpansionFrameInView(
	contentFrame core.NSRect,
	view NSViewRef,
) {
	C.NSControl_inst_DrawWithExpansionFrameInView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentFrame)),
		objc.RefPointer(view),
	)

	return
}

// EditWithFrameEditorDelegateEvent begins editing of the receiver’s text using the specified field editor.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428919-editwithframe?language=objc for details.
func (x gen_NSControl) EditWithFrameEditorDelegateEvent(
	rect core.NSRect,
	textObj NSTextRef,
	delegate objc.Ref,
	event NSEventRef,
) {
	C.NSControl_inst_EditWithFrameEditorDelegateEvent(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(textObj),
		objc.RefPointer(delegate),
		objc.RefPointer(event),
	)

	return
}

// EndEditing ends the editing of text in the receiver using the specified field editor.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428936-endediting?language=objc for details.
func (x gen_NSControl) EndEditing(
	textObj NSTextRef,
) {
	C.NSControl_inst_EndEditing(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(textObj),
	)

	return
}

// ExpansionFrameWithFrame returns the frame in which a tool tip can be displayed, if needed.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428932-expansionframewithframe?language=objc for details.
func (x gen_NSControl) ExpansionFrameWithFrame(
	contentFrame core.NSRect,
) core.NSRect {
	ret := C.NSControl_inst_ExpansionFrameWithFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentFrame)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// InitWithFrame initializes a control with the specified frame rectangle.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc for details.
func (x gen_NSControl) InitWithFrame(
	frameRect core.NSRect,
) NSControl {
	ret := C.NSControl_inst_InitWithFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)

	return NSControl_FromPointer(ret)
}

// MouseDown informs the receiver that the user has pressed the left mouse button.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428918-mousedown?language=objc for details.
func (x gen_NSControl) MouseDown(
	event NSEventRef,
) {
	C.NSControl_inst_MouseDown(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return
}

// PerformClick simulates a single mouse click on the receiver.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428974-performclick?language=objc for details.
func (x gen_NSControl) PerformClick(
	sender objc.Ref,
) {
	C.NSControl_inst_PerformClick(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// SelectWithFrameEditorDelegateStartLength selects the specified text range in the receiver's field editor.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428968-selectwithframe?language=objc for details.
func (x gen_NSControl) SelectWithFrameEditorDelegateStartLength(
	rect core.NSRect,
	textObj NSTextRef,
	delegate objc.Ref,
	selStart core.NSInteger,
	selLength core.NSInteger,
) {
	C.NSControl_inst_SelectWithFrameEditorDelegateStartLength(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(textObj),
		objc.RefPointer(delegate),
		C.long(selStart),
		C.long(selLength),
	)

	return
}

// SendActionTo causes the specified action to be sent to the target.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428851-sendaction?language=objc for details.
func (x gen_NSControl) SendActionTo(
	action objc.Selector,
	target objc.Ref,
) bool {
	ret := C.NSControl_inst_SendActionTo(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(target),
	)

	return convertObjCBoolToGo(ret)
}

// SizeThatFits asks the control to calculate and return the size that best fits the specified size.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428902-sizethatfits?language=objc for details.
func (x gen_NSControl) SizeThatFits(
	size core.NSSize,
) core.NSSize {
	ret := C.NSControl_inst_SizeThatFits(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SizeToFit resizes the receiver’s frame so that it’s the minimum size needed to contain its cell.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428877-sizetofit?language=objc for details.
func (x gen_NSControl) SizeToFit() {
	C.NSControl_inst_SizeToFit(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// TakeDoubleValueFrom sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428958-takedoublevaluefrom?language=objc for details.
func (x gen_NSControl) TakeDoubleValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_TakeDoubleValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// TakeFloatValueFrom sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428938-takefloatvaluefrom?language=objc for details.
func (x gen_NSControl) TakeFloatValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_TakeFloatValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// TakeIntValueFrom sets the value of the receiver’s cell to an integer value obtained from the specified object.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428859-takeintvaluefrom?language=objc for details.
func (x gen_NSControl) TakeIntValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_TakeIntValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// TakeIntegerValueFrom sets the value of the receiver’s cell to an NSInteger value obtained from the specified object.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428875-takeintegervaluefrom?language=objc for details.
func (x gen_NSControl) TakeIntegerValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_TakeIntegerValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// TakeObjectValueFrom sets the value of the receiver’s cell to the object value obtained from the specified object.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428853-takeobjectvaluefrom?language=objc for details.
func (x gen_NSControl) TakeObjectValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_TakeObjectValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// TakeStringValueFrom sets the value of the receiver’s cell to the string value obtained from the specified object.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428912-takestringvaluefrom?language=objc for details.
func (x gen_NSControl) TakeStringValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_TakeStringValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ValidateEditing validates changes to any user-typed text.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428855-validateediting?language=objc for details.
func (x gen_NSControl) ValidateEditing() {
	C.NSControl_inst_ValidateEditing(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Init is undocumented.
func (x gen_NSControl) Init() NSControl {
	ret := C.NSControl_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSControl_FromPointer(ret)
}

// IsEnabled returns a Boolean value that indicates whether the receiver reacts to mouse events.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428970-enabled?language=objc for details.
func (x gen_NSControl) IsEnabled() bool {
	ret := C.NSControl_inst_IsEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetEnabled returns a Boolean value that indicates whether the receiver reacts to mouse events.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428970-enabled?language=objc for details.
func (x gen_NSControl) SetEnabled(
	value bool,
) {
	C.NSControl_inst_SetEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IntValue returns the value of the receiver’s cell as an integer.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428939-intvalue?language=objc for details.
func (x gen_NSControl) IntValue() int32 {
	ret := C.NSControl_inst_IntValue(
		unsafe.Pointer(x.Pointer()),
	)

	return int32(ret)
}

// SetIntValue returns the value of the receiver’s cell as an integer.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428939-intvalue?language=objc for details.
func (x gen_NSControl) SetIntValue(
	value int32,
) {
	C.NSControl_inst_SetIntValue(
		unsafe.Pointer(x.Pointer()),
		C.int(value),
	)

	return
}

// IntegerValue returns the value of the receiver’s cell as an NSInteger value.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428969-integervalue?language=objc for details.
func (x gen_NSControl) IntegerValue() core.NSInteger {
	ret := C.NSControl_inst_IntegerValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// SetIntegerValue returns the value of the receiver’s cell as an NSInteger value.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428969-integervalue?language=objc for details.
func (x gen_NSControl) SetIntegerValue(
	value core.NSInteger,
) {
	C.NSControl_inst_SetIntegerValue(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return
}

// ObjectValue returns the value of the receiver’s cell as an Objective-C object.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428849-objectvalue?language=objc for details.
func (x gen_NSControl) ObjectValue() objc.Object {
	ret := C.NSControl_inst_ObjectValue(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetObjectValue returns the value of the receiver’s cell as an Objective-C object.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428849-objectvalue?language=objc for details.
func (x gen_NSControl) SetObjectValue(
	value objc.Ref,
) {
	C.NSControl_inst_SetObjectValue(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// StringValue returns the value of the receiver’s cell as an NSString object.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428950-stringvalue?language=objc for details.
func (x gen_NSControl) StringValue() core.NSString {
	ret := C.NSControl_inst_StringValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetStringValue returns the value of the receiver’s cell as an NSString object.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428950-stringvalue?language=objc for details.
func (x gen_NSControl) SetStringValue(
	value core.NSStringRef,
) {
	C.NSControl_inst_SetStringValue(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// AttributedStringValue returns the value of the receiver’s cell as an attributed string.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428916-attributedstringvalue?language=objc for details.
func (x gen_NSControl) AttributedStringValue() core.NSAttributedString {
	ret := C.NSControl_inst_AttributedStringValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSAttributedString_FromPointer(ret)
}

// SetAttributedStringValue returns the value of the receiver’s cell as an attributed string.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428916-attributedstringvalue?language=objc for details.
func (x gen_NSControl) SetAttributedStringValue(
	value core.NSAttributedStringRef,
) {
	C.NSControl_inst_SetAttributedStringValue(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Font returns the font used to draw text in the receiver’s cell.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428914-font?language=objc for details.
func (x gen_NSControl) Font() NSFont {
	ret := C.NSControl_inst_Font(
		unsafe.Pointer(x.Pointer()),
	)

	return NSFont_FromPointer(ret)
}

// SetFont returns the font used to draw text in the receiver’s cell.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428914-font?language=objc for details.
func (x gen_NSControl) SetFont(
	value NSFontRef,
) {
	C.NSControl_inst_SetFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// UsesSingleLineMode returns a Boolean value that indicates whether the text in the control’s cell uses single line mode.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428929-usessinglelinemode?language=objc for details.
func (x gen_NSControl) UsesSingleLineMode() bool {
	ret := C.NSControl_inst_UsesSingleLineMode(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetUsesSingleLineMode returns a Boolean value that indicates whether the text in the control’s cell uses single line mode.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428929-usessinglelinemode?language=objc for details.
func (x gen_NSControl) SetUsesSingleLineMode(
	value bool,
) {
	C.NSControl_inst_SetUsesSingleLineMode(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AllowsExpansionToolTips returns a Boolean value that indicates whether expansion tool tips are shown when the control is hovered over.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428962-allowsexpansiontooltips?language=objc for details.
func (x gen_NSControl) AllowsExpansionToolTips() bool {
	ret := C.NSControl_inst_AllowsExpansionToolTips(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsExpansionToolTips returns a Boolean value that indicates whether expansion tool tips are shown when the control is hovered over.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428962-allowsexpansiontooltips?language=objc for details.
func (x gen_NSControl) SetAllowsExpansionToolTips(
	value bool,
) {
	C.NSControl_inst_SetAllowsExpansionToolTips(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsHighlighted returns a Boolean value that indicates whether the cell is highlighted.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428927-highlighted?language=objc for details.
func (x gen_NSControl) IsHighlighted() bool {
	ret := C.NSControl_inst_IsHighlighted(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetHighlighted returns a Boolean value that indicates whether the cell is highlighted.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428927-highlighted?language=objc for details.
func (x gen_NSControl) SetHighlighted(
	value bool,
) {
	C.NSControl_inst_SetHighlighted(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// Action returns the default action-message selector associated with the control.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428956-action?language=objc for details.
func (x gen_NSControl) Action() objc.Selector {
	ret := C.NSControl_inst_Action(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.SelectorAt(ret)
}

// SetAction returns the default action-message selector associated with the control.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428956-action?language=objc for details.
func (x gen_NSControl) SetAction(
	value objc.Selector,
) {
	C.NSControl_inst_SetAction(
		unsafe.Pointer(x.Pointer()),
		value.SelectorAddress(),
	)

	return
}

// Target returns the target object that receives action messages from the cell.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428885-target?language=objc for details.
func (x gen_NSControl) Target() objc.Object {
	ret := C.NSControl_inst_Target(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetTarget returns the target object that receives action messages from the cell.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428885-target?language=objc for details.
func (x gen_NSControl) SetTarget(
	value objc.Ref,
) {
	C.NSControl_inst_SetTarget(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// IsContinuous returns a Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428952-continuous?language=objc for details.
func (x gen_NSControl) IsContinuous() bool {
	ret := C.NSControl_inst_IsContinuous(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetContinuous returns a Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428952-continuous?language=objc for details.
func (x gen_NSControl) SetContinuous(
	value bool,
) {
	C.NSControl_inst_SetContinuous(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// Tag returns the tag identifying the receiver (not the tag of the receiver’s cell).
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428910-tag?language=objc for details.
func (x gen_NSControl) Tag() core.NSInteger {
	ret := C.NSControl_inst_Tag(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// SetTag returns the tag identifying the receiver (not the tag of the receiver’s cell).
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428910-tag?language=objc for details.
func (x gen_NSControl) SetTag(
	value core.NSInteger,
) {
	C.NSControl_inst_SetTag(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return
}

// RefusesFirstResponder returns a Boolean value indicating whether the receiver refuses the first responder role.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428976-refusesfirstresponder?language=objc for details.
func (x gen_NSControl) RefusesFirstResponder() bool {
	ret := C.NSControl_inst_RefusesFirstResponder(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetRefusesFirstResponder returns a Boolean value indicating whether the receiver refuses the first responder role.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428976-refusesfirstresponder?language=objc for details.
func (x gen_NSControl) SetRefusesFirstResponder(
	value bool,
) {
	C.NSControl_inst_SetRefusesFirstResponder(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IgnoresMultiClick returns a Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428863-ignoresmulticlick?language=objc for details.
func (x gen_NSControl) IgnoresMultiClick() bool {
	ret := C.NSControl_inst_IgnoresMultiClick(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetIgnoresMultiClick returns a Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession.
//
// See https://developer.apple.com/documentation/appkit/nscontrol/1428863-ignoresmulticlick?language=objc for details.
func (x gen_NSControl) SetIgnoresMultiClick(
	value bool,
) {
	C.NSControl_inst_SetIgnoresMultiClick(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

type NSButtonRef interface {
	Pointer() uintptr
	Init() NSButton
}

type gen_NSButton struct {
	NSControl
}

func NSButton_FromPointer(ptr unsafe.Pointer) NSButton {
	return NSButton{gen_NSButton{
		NSControl_FromPointer(ptr),
	}}
}

func NSButton_FromRef(ref objc.Ref) NSButton {
	return NSButton_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// CompressWithPrioritizedCompressionOptions sets the priority compression options for this button.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/2952060-compresswithprioritizedcompressi?language=objc for details.
func (x gen_NSButton) CompressWithPrioritizedCompressionOptions(
	prioritizedOptions core.NSArrayRef,
) {
	C.NSButton_inst_CompressWithPrioritizedCompressionOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(prioritizedOptions),
	)

	return
}

// Highlight highlights (or unhighlights) the button.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1534156-highlight?language=objc for details.
func (x gen_NSButton) Highlight(
	flag bool,
) {
	C.NSButton_inst_Highlight(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)

	return
}

// MinimumSizeWithPrioritizedCompressionOptions returns the minimum size of the button by using the compression options.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/2952059-minimumsizewithprioritizedcompre?language=objc for details.
func (x gen_NSButton) MinimumSizeWithPrioritizedCompressionOptions(
	prioritizedOptions core.NSArrayRef,
) core.NSSize {
	ret := C.NSButton_inst_MinimumSizeWithPrioritizedCompressionOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(prioritizedOptions),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// PerformKeyEquivalent checks the button's key equivalent against the specified event and, if they match, simulates the button being clicked.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1524423-performkeyequivalent?language=objc for details.
func (x gen_NSButton) PerformKeyEquivalent(
	key NSEventRef,
) bool {
	ret := C.NSButton_inst_PerformKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)

	return convertObjCBoolToGo(ret)
}

// SetNextState sets the button to its next state.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1530594-setnextstate?language=objc for details.
func (x gen_NSButton) SetNextState() {
	C.NSButton_inst_SetNextState(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Init is undocumented.
func (x gen_NSButton) Init() NSButton {
	ret := C.NSButton_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSButton_FromPointer(ret)
}

// ContentTintColor returns a tint color to use for the template image and text content.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/3000781-contenttintcolor?language=objc for details.
func (x gen_NSButton) ContentTintColor() NSColor {
	ret := C.NSButton_inst_ContentTintColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_FromPointer(ret)
}

// SetContentTintColor returns a tint color to use for the template image and text content.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/3000781-contenttintcolor?language=objc for details.
func (x gen_NSButton) SetContentTintColor(
	value NSColorRef,
) {
	C.NSButton_inst_SetContentTintColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// HasDestructiveAction returns a Boolean value that defines whether a button’s action has a destructive effect.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/3622469-hasdestructiveaction?language=objc for details.
func (x gen_NSButton) HasDestructiveAction() bool {
	ret := C.NSButton_inst_HasDestructiveAction(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetHasDestructiveAction returns a Boolean value that defines whether a button’s action has a destructive effect.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/3622469-hasdestructiveaction?language=objc for details.
func (x gen_NSButton) SetHasDestructiveAction(
	value bool,
) {
	C.NSButton_inst_SetHasDestructiveAction(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AlternateTitle returns the title that the button displays when the button is in an on state.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1529588-alternatetitle?language=objc for details.
func (x gen_NSButton) AlternateTitle() core.NSString {
	ret := C.NSButton_inst_AlternateTitle(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetAlternateTitle returns the title that the button displays when the button is in an on state.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1529588-alternatetitle?language=objc for details.
func (x gen_NSButton) SetAlternateTitle(
	value core.NSStringRef,
) {
	C.NSButton_inst_SetAlternateTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// AttributedTitle returns the title that the button displays in an off state, as an attributed string.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1524640-attributedtitle?language=objc for details.
func (x gen_NSButton) AttributedTitle() core.NSAttributedString {
	ret := C.NSButton_inst_AttributedTitle(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSAttributedString_FromPointer(ret)
}

// SetAttributedTitle returns the title that the button displays in an off state, as an attributed string.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1524640-attributedtitle?language=objc for details.
func (x gen_NSButton) SetAttributedTitle(
	value core.NSAttributedStringRef,
) {
	C.NSButton_inst_SetAttributedTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// AttributedAlternateTitle returns the title that the button displays as an attributed string when the button is in an on state.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1526723-attributedalternatetitle?language=objc for details.
func (x gen_NSButton) AttributedAlternateTitle() core.NSAttributedString {
	ret := C.NSButton_inst_AttributedAlternateTitle(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSAttributedString_FromPointer(ret)
}

// SetAttributedAlternateTitle returns the title that the button displays as an attributed string when the button is in an on state.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1526723-attributedalternatetitle?language=objc for details.
func (x gen_NSButton) SetAttributedAlternateTitle(
	value core.NSAttributedStringRef,
) {
	C.NSButton_inst_SetAttributedAlternateTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Title returns the title displayed on the button when it’s in an off state.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1524430-title?language=objc for details.
func (x gen_NSButton) Title() core.NSString {
	ret := C.NSButton_inst_Title(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetTitle returns the title displayed on the button when it’s in an off state.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1524430-title?language=objc for details.
func (x gen_NSButton) SetTitle(
	value core.NSStringRef,
) {
	C.NSButton_inst_SetTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Sound returns the sound that plays when the user clicks the button.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1530910-sound?language=objc for details.
func (x gen_NSButton) Sound() NSSound {
	ret := C.NSButton_inst_Sound(
		unsafe.Pointer(x.Pointer()),
	)

	return NSSound_FromPointer(ret)
}

// SetSound returns the sound that plays when the user clicks the button.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1530910-sound?language=objc for details.
func (x gen_NSButton) SetSound(
	value NSSoundRef,
) {
	C.NSButton_inst_SetSound(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// IsSpringLoaded returns a Boolean value that indicates whether spring loading is enabled for the button.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1532300-springloaded?language=objc for details.
func (x gen_NSButton) IsSpringLoaded() bool {
	ret := C.NSButton_inst_IsSpringLoaded(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetSpringLoaded returns a Boolean value that indicates whether spring loading is enabled for the button.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1532300-springloaded?language=objc for details.
func (x gen_NSButton) SetSpringLoaded(
	value bool,
) {
	C.NSButton_inst_SetSpringLoaded(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// MaxAcceleratorLevel an integer value indicating the maximum pressure level for a button of type NSMultiLevelAcceleratorButton.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1534413-maxacceleratorlevel?language=objc for details.
func (x gen_NSButton) MaxAcceleratorLevel() core.NSInteger {
	ret := C.NSButton_inst_MaxAcceleratorLevel(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// SetMaxAcceleratorLevel an integer value indicating the maximum pressure level for a button of type NSMultiLevelAcceleratorButton.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1534413-maxacceleratorlevel?language=objc for details.
func (x gen_NSButton) SetMaxAcceleratorLevel(
	value core.NSInteger,
) {
	C.NSButton_inst_SetMaxAcceleratorLevel(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return
}

// Image returns the image that appears on the button when it’s in an off state, or nil if there is no such image.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1534221-image?language=objc for details.
func (x gen_NSButton) Image() NSImage {
	ret := C.NSButton_inst_Image(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_FromPointer(ret)
}

// SetImage returns the image that appears on the button when it’s in an off state, or nil if there is no such image.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1534221-image?language=objc for details.
func (x gen_NSButton) SetImage(
	value NSImageRef,
) {
	C.NSButton_inst_SetImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// AlternateImage an alternate image that appears on the button when the button is in an on state.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1533935-alternateimage?language=objc for details.
func (x gen_NSButton) AlternateImage() NSImage {
	ret := C.NSButton_inst_AlternateImage(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_FromPointer(ret)
}

// SetAlternateImage an alternate image that appears on the button when the button is in an on state.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1533935-alternateimage?language=objc for details.
func (x gen_NSButton) SetAlternateImage(
	value NSImageRef,
) {
	C.NSButton_inst_SetAlternateImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// IsBordered returns a Boolean value that determines whether the button has a border.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1525565-bordered?language=objc for details.
func (x gen_NSButton) IsBordered() bool {
	ret := C.NSButton_inst_IsBordered(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetBordered returns a Boolean value that determines whether the button has a border.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1525565-bordered?language=objc for details.
func (x gen_NSButton) SetBordered(
	value bool,
) {
	C.NSButton_inst_SetBordered(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsTransparent returns a Boolean value that indicates whether the button is transparent.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1529659-transparent?language=objc for details.
func (x gen_NSButton) IsTransparent() bool {
	ret := C.NSButton_inst_IsTransparent(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetTransparent returns a Boolean value that indicates whether the button is transparent.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1529659-transparent?language=objc for details.
func (x gen_NSButton) SetTransparent(
	value bool,
) {
	C.NSButton_inst_SetTransparent(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// BezelColor returns the color of the button's bezel, in appearances that support it.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/2561000-bezelcolor?language=objc for details.
func (x gen_NSButton) BezelColor() NSColor {
	ret := C.NSButton_inst_BezelColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_FromPointer(ret)
}

// SetBezelColor returns the color of the button's bezel, in appearances that support it.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/2561000-bezelcolor?language=objc for details.
func (x gen_NSButton) SetBezelColor(
	value NSColorRef,
) {
	C.NSButton_inst_SetBezelColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// ShowsBorderOnlyWhileMouseInside returns a Boolean value that determines whether the button displays its border only when the pointer is over it.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1532248-showsborderonlywhilemouseinside?language=objc for details.
func (x gen_NSButton) ShowsBorderOnlyWhileMouseInside() bool {
	ret := C.NSButton_inst_ShowsBorderOnlyWhileMouseInside(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetShowsBorderOnlyWhileMouseInside returns a Boolean value that determines whether the button displays its border only when the pointer is over it.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1532248-showsborderonlywhilemouseinside?language=objc for details.
func (x gen_NSButton) SetShowsBorderOnlyWhileMouseInside(
	value bool,
) {
	C.NSButton_inst_SetShowsBorderOnlyWhileMouseInside(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// ImageHugsTitle returns a Boolean value that determines how the button’s image and title are positioned together within the button bezel.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/2092414-imagehugstitle?language=objc for details.
func (x gen_NSButton) ImageHugsTitle() bool {
	ret := C.NSButton_inst_ImageHugsTitle(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetImageHugsTitle returns a Boolean value that determines how the button’s image and title are positioned together within the button bezel.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/2092414-imagehugstitle?language=objc for details.
func (x gen_NSButton) SetImageHugsTitle(
	value bool,
) {
	C.NSButton_inst_SetImageHugsTitle(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AllowsMixedState returns a Boolean value that indicates whether the button allows a mixed state.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1528670-allowsmixedstate?language=objc for details.
func (x gen_NSButton) AllowsMixedState() bool {
	ret := C.NSButton_inst_AllowsMixedState(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsMixedState returns a Boolean value that indicates whether the button allows a mixed state.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1528670-allowsmixedstate?language=objc for details.
func (x gen_NSButton) SetAllowsMixedState(
	value bool,
) {
	C.NSButton_inst_SetAllowsMixedState(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// State returns the button’s state.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1528907-state?language=objc for details.
func (x gen_NSButton) State() core.NSInteger {
	ret := C.NSButton_inst_State(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// SetState returns the button’s state.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1528907-state?language=objc for details.
func (x gen_NSButton) SetState(
	value core.NSInteger,
) {
	C.NSButton_inst_SetState(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return
}

// KeyEquivalent returns the key-equivalent character of the button.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1525368-keyequivalent?language=objc for details.
func (x gen_NSButton) KeyEquivalent() core.NSString {
	ret := C.NSButton_inst_KeyEquivalent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetKeyEquivalent returns the key-equivalent character of the button.
//
// See https://developer.apple.com/documentation/appkit/nsbutton/1525368-keyequivalent?language=objc for details.
func (x gen_NSButton) SetKeyEquivalent(
	value core.NSStringRef,
) {
	C.NSButton_inst_SetKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

type NSEventRef interface {
	Pointer() uintptr
	Init() NSEvent
}

type gen_NSEvent struct {
	objc.Object
}

func NSEvent_FromPointer(ptr unsafe.Pointer) NSEvent {
	return NSEvent{gen_NSEvent{
		objc.Object_FromPointer(ptr),
	}}
}

func NSEvent_FromRef(ref objc.Ref) NSEvent {
	return NSEvent_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init is undocumented.
func (x gen_NSEvent) Init() NSEvent {
	ret := C.NSEvent_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSEvent_FromPointer(ret)
}

// LocationInWindow returns the receiver’s location in the base coordinate system of the associated window.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1529068-locationinwindow?language=objc for details.
func (x gen_NSEvent) LocationInWindow() core.NSPoint {
	ret := C.NSEvent_inst_LocationInWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// Window returns the window object associated with the event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1530808-window?language=objc for details.
func (x gen_NSEvent) Window() NSWindow {
	ret := C.NSEvent_inst_Window(
		unsafe.Pointer(x.Pointer()),
	)

	return NSWindow_FromPointer(ret)
}

// WindowNumber returns the identifier for the window device associated with the event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1531361-windownumber?language=objc for details.
func (x gen_NSEvent) WindowNumber() core.NSInteger {
	ret := C.NSEvent_inst_WindowNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// EventRef an opaque Carbon type associated with this event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1525143-eventref?language=objc for details.
func (x gen_NSEvent) EventRef() unsafe.Pointer {
	ret := C.NSEvent_inst_EventRef(
		unsafe.Pointer(x.Pointer()),
	)

	return ret
}

// Characters returns the characters associated with a key-up or key-down event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1534183-characters?language=objc for details.
func (x gen_NSEvent) Characters() core.NSString {
	ret := C.NSEvent_inst_Characters(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// CharactersIgnoringModifiers returns the characters generated by a key event as if no modifier key (except for Shift) applies.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1524605-charactersignoringmodifiers?language=objc for details.
func (x gen_NSEvent) CharactersIgnoringModifiers() core.NSString {
	ret := C.NSEvent_inst_CharactersIgnoringModifiers(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// IsARepeat returns a Boolean value that indicates whether the key event is a repeat.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1528049-arepeat?language=objc for details.
func (x gen_NSEvent) IsARepeat() bool {
	ret := C.NSEvent_inst_IsARepeat(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// ButtonNumber returns the button number for a mouse event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1527828-buttonnumber?language=objc for details.
func (x gen_NSEvent) ButtonNumber() core.NSInteger {
	ret := C.NSEvent_inst_ButtonNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// ClickCount returns the number of mouse clicks associated with a mouse-down or mouse-up event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1528200-clickcount?language=objc for details.
func (x gen_NSEvent) ClickCount() core.NSInteger {
	ret := C.NSEvent_inst_ClickCount(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// EventNumber returns the counter value of the latest mouse or tracking-rectangle event object; every system-generated mouse and tracking-rectangle event increments this counter.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1535220-eventnumber?language=objc for details.
func (x gen_NSEvent) EventNumber() core.NSInteger {
	ret := C.NSEvent_inst_EventNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// TrackingNumber returns the identifier of a mouse-tracking event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1533974-trackingnumber?language=objc for details.
func (x gen_NSEvent) TrackingNumber() core.NSInteger {
	ret := C.NSEvent_inst_TrackingNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// UserData returns the data associated with a mouse-tracking event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1526810-userdata?language=objc for details.
func (x gen_NSEvent) UserData() unsafe.Pointer {
	ret := C.NSEvent_inst_UserData(
		unsafe.Pointer(x.Pointer()),
	)

	return ret
}

// Data1 additional data associated with this event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1528289-data1?language=objc for details.
func (x gen_NSEvent) Data1() core.NSInteger {
	ret := C.NSEvent_inst_Data1(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// Data2 additional data associated with this event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1528647-data2?language=objc for details.
func (x gen_NSEvent) Data2() core.NSInteger {
	ret := C.NSEvent_inst_Data2(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// DeltaX returns the x-coordinate change for mouse-move, mouse-drag, and swipe events.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1534871-deltax?language=objc for details.
func (x gen_NSEvent) DeltaX() core.CGFloat {
	ret := C.NSEvent_inst_DeltaX(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// DeltaY returns the y-coordinate change for mouse-move, mouse-drag, and swipe events.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1534158-deltay?language=objc for details.
func (x gen_NSEvent) DeltaY() core.CGFloat {
	ret := C.NSEvent_inst_DeltaY(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// DeltaZ returns the z-coordinate change for a scroll wheel, mouse-move, or mouse-drag event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1531528-deltaz?language=objc for details.
func (x gen_NSEvent) DeltaZ() core.CGFloat {
	ret := C.NSEvent_inst_DeltaZ(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// Stage returns a value of 0, 1, or 2, indicating the stage of a gesture event of type NSEventTypePressure.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1527242-stage?language=objc for details.
func (x gen_NSEvent) Stage() core.NSInteger {
	ret := C.NSEvent_inst_Stage(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// StageTransition returns the transition value for the stage of a pressure gesture event of type NSEventTypePressure.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1526739-stagetransition?language=objc for details.
func (x gen_NSEvent) StageTransition() core.CGFloat {
	ret := C.NSEvent_inst_StageTransition(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// CapabilityMask returns a mask whose set bits indicate the capabilities of the tablet device that generated this event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1534648-capabilitymask?language=objc for details.
func (x gen_NSEvent) CapabilityMask() core.NSUInteger {
	ret := C.NSEvent_inst_CapabilityMask(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// DeviceID returns a special identifier that is used to match tablet-pointer and tablet-proximity events.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1530014-deviceid?language=objc for details.
func (x gen_NSEvent) DeviceID() core.NSUInteger {
	ret := C.NSEvent_inst_DeviceID(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// IsEnteringProximity returns a Boolean value that indicates whether a pointing device is entering or leaving the proximity of its tablet.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1531702-enteringproximity?language=objc for details.
func (x gen_NSEvent) IsEnteringProximity() bool {
	ret := C.NSEvent_inst_IsEnteringProximity(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// PointingDeviceID returns the index of the pointing device currently in proximity with the tablet.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1528818-pointingdeviceid?language=objc for details.
func (x gen_NSEvent) PointingDeviceID() core.NSUInteger {
	ret := C.NSEvent_inst_PointingDeviceID(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// PointingDeviceSerialNumber returns the vendor-assigned serial number of a pointing device.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1533420-pointingdeviceserialnumber?language=objc for details.
func (x gen_NSEvent) PointingDeviceSerialNumber() core.NSUInteger {
	ret := C.NSEvent_inst_PointingDeviceSerialNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// SystemTabletID returns the index of the tablet device connected to the system.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1528299-systemtabletid?language=objc for details.
func (x gen_NSEvent) SystemTabletID() core.NSUInteger {
	ret := C.NSEvent_inst_SystemTabletID(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// TabletID returns the USB model identifier of the tablet device associated with this event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1527003-tabletid?language=objc for details.
func (x gen_NSEvent) TabletID() core.NSUInteger {
	ret := C.NSEvent_inst_TabletID(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// VendorID returns the vendor identifier of the tablet associated with the event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1525177-vendorid?language=objc for details.
func (x gen_NSEvent) VendorID() core.NSUInteger {
	ret := C.NSEvent_inst_VendorID(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// VendorPointingDeviceType returns a coded bit field whose set bits indicate the type of pointing device (within a vendor selection) associated with the event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1527736-vendorpointingdevicetype?language=objc for details.
func (x gen_NSEvent) VendorPointingDeviceType() core.NSUInteger {
	ret := C.NSEvent_inst_VendorPointingDeviceType(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// AbsoluteX returns the absolute x coordinate of a pointing device on its tablet at full tablet resolution.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1530617-absolutex?language=objc for details.
func (x gen_NSEvent) AbsoluteX() core.NSInteger {
	ret := C.NSEvent_inst_AbsoluteX(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// AbsoluteY returns the absolute y coordinate of a pointing device on its tablet at full tablet resolution.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1528904-absolutey?language=objc for details.
func (x gen_NSEvent) AbsoluteY() core.NSInteger {
	ret := C.NSEvent_inst_AbsoluteY(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// AbsoluteZ returns the absolute z coordinate of pointing device on its tablet at full tablet resolution.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1532154-absolutez?language=objc for details.
func (x gen_NSEvent) AbsoluteZ() core.NSInteger {
	ret := C.NSEvent_inst_AbsoluteZ(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// Tilt returns the scaled tilt values of the pointing device that generated this event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1534226-tilt?language=objc for details.
func (x gen_NSEvent) Tilt() core.NSPoint {
	ret := C.NSEvent_inst_Tilt(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// VendorDefined an array of three vendor-defined NSNumber objects associated with a pointing-type event.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1530551-vendordefined?language=objc for details.
func (x gen_NSEvent) VendorDefined() objc.Object {
	ret := C.NSEvent_inst_VendorDefined(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// Magnification returns the change in magnification.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1531642-magnification?language=objc for details.
func (x gen_NSEvent) Magnification() core.CGFloat {
	ret := C.NSEvent_inst_Magnification(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// HasPreciseScrollingDeltas returns a Boolean value that indicates whether precise scrolling deltas are available.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1525758-hasprecisescrollingdeltas?language=objc for details.
func (x gen_NSEvent) HasPreciseScrollingDeltas() bool {
	ret := C.NSEvent_inst_HasPreciseScrollingDeltas(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// ScrollingDeltaX returns the scroll wheel’s horizontal delta.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1524505-scrollingdeltax?language=objc for details.
func (x gen_NSEvent) ScrollingDeltaX() core.CGFloat {
	ret := C.NSEvent_inst_ScrollingDeltaX(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// ScrollingDeltaY returns the scroll wheel’s vertical delta.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1535387-scrollingdeltay?language=objc for details.
func (x gen_NSEvent) ScrollingDeltaY() core.CGFloat {
	ret := C.NSEvent_inst_ScrollingDeltaY(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// IsDirectionInvertedFromDevice returns a Boolean value that indicates whether the user has changed the device inversion.
//
// See https://developer.apple.com/documentation/appkit/nsevent/1525151-directioninvertedfromdevice?language=objc for details.
func (x gen_NSEvent) IsDirectionInvertedFromDevice() bool {
	ret := C.NSEvent_inst_IsDirectionInvertedFromDevice(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

type NSFontRef interface {
	Pointer() uintptr
	Init() NSFont
}

type gen_NSFont struct {
	objc.Object
}

func NSFont_FromPointer(ptr unsafe.Pointer) NSFont {
	return NSFont{gen_NSFont{
		objc.Object_FromPointer(ptr),
	}}
}

func NSFont_FromRef(ref objc.Ref) NSFont {
	return NSFont_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// FontWithSize is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsfont/3667454-fontwithsize?language=objc for details.
func (x gen_NSFont) FontWithSize(
	fontSize core.CGFloat,
) NSFont {
	ret := C.NSFont_inst_FontWithSize(
		unsafe.Pointer(x.Pointer()),
		C.double(fontSize),
	)

	return NSFont_FromPointer(ret)
}

// Set sets this font as the font for the current graphics context.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1531373-set?language=objc for details.
func (x gen_NSFont) Set() {
	C.NSFont_inst_Set(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Init is undocumented.
func (x gen_NSFont) Init() NSFont {
	ret := C.NSFont_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSFont_FromPointer(ret)
}

// PointSize returns the point size of the font.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1524511-pointsize?language=objc for details.
func (x gen_NSFont) PointSize() core.CGFloat {
	ret := C.NSFont_inst_PointSize(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// IsFixedPitch returns a Boolean value indicating whether all glyphs in the font have the same advancement.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1529210-fixedpitch?language=objc for details.
func (x gen_NSFont) IsFixedPitch() bool {
	ret := C.NSFont_inst_IsFixedPitch(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// MostCompatibleStringEncoding returns the string encoding that works best with the font.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1527635-mostcompatiblestringencoding?language=objc for details.
func (x gen_NSFont) MostCompatibleStringEncoding() core.NSStringEncoding {
	ret := C.NSFont_inst_MostCompatibleStringEncoding(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSStringEncoding(ret)
}

// NumberOfGlyphs returns the number of glyphs in the font.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1533968-numberofglyphs?language=objc for details.
func (x gen_NSFont) NumberOfGlyphs() core.NSUInteger {
	ret := C.NSFont_inst_NumberOfGlyphs(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// DisplayName returns the name of the font, including family and face names, to use when displaying the font information to the user.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1531660-displayname?language=objc for details.
func (x gen_NSFont) DisplayName() core.NSString {
	ret := C.NSFont_inst_DisplayName(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// FamilyName returns the family name of the font—for example, “Times” or “Helvetica.”
//
// See https://developer.apple.com/documentation/appkit/nsfont/1529585-familyname?language=objc for details.
func (x gen_NSFont) FamilyName() core.NSString {
	ret := C.NSFont_inst_FamilyName(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// FontName returns the full name of the font, as used in PostScript language code—for example, “Times-Roman” or “Helvetica-Oblique.”
//
// See https://developer.apple.com/documentation/appkit/nsfont/1526183-fontname?language=objc for details.
func (x gen_NSFont) FontName() core.NSString {
	ret := C.NSFont_inst_FontName(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// IsVertical returns a Boolean value indicating whether the font is a vertical font.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1534644-vertical?language=objc for details.
func (x gen_NSFont) IsVertical() bool {
	ret := C.NSFont_inst_IsVertical(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// VerticalFont returns a vertical version of the font.
//
// See https://developer.apple.com/documentation/appkit/nsfont/1535152-verticalfont?language=objc for details.
func (x gen_NSFont) VerticalFont() NSFont {
	ret := C.NSFont_inst_VerticalFont(
		unsafe.Pointer(x.Pointer()),
	)

	return NSFont_FromPointer(ret)
}

type NSImageRef interface {
	Pointer() uintptr
	Init() NSImage
}

type gen_NSImage struct {
	objc.Object
}

func NSImage_FromPointer(ptr unsafe.Pointer) NSImage {
	return NSImage{gen_NSImage{
		objc.Object_FromPointer(ptr),
	}}
}

func NSImage_FromRef(ref objc.Ref) NSImage {
	return NSImage_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// AddRepresentations adds an array of image representation objects to the image.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519964-addrepresentations?language=objc for details.
func (x gen_NSImage) AddRepresentations(
	imageReps core.NSArrayRef,
) {
	C.NSImage_inst_AddRepresentations(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(imageReps),
	)

	return
}

// CancelIncrementalLoad cancels the current download operation, if any, for an incrementally loaded image.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1520041-cancelincrementalload?language=objc for details.
func (x gen_NSImage) CancelIncrementalLoad() {
	C.NSImage_inst_CancelIncrementalLoad(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DrawInRect draws the image in the specified rectangle.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519863-drawinrect?language=objc for details.
func (x gen_NSImage) DrawInRect(
	rect core.NSRect,
) {
	C.NSImage_inst_DrawInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return
}

// InitByReferencingFile initializes and returns an image object using the specified file.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519955-initbyreferencingfile?language=objc for details.
func (x gen_NSImage) InitByReferencingFile(
	fileName core.NSStringRef,
) NSImage {
	ret := C.NSImage_inst_InitByReferencingFile(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fileName),
	)

	return NSImage_FromPointer(ret)
}

// InitByReferencingURL initializes and returns an image object using the specified URL.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519990-initbyreferencingurl?language=objc for details.
func (x gen_NSImage) InitByReferencingURL(
	url core.NSURLRef,
) NSImage {
	ret := C.NSImage_inst_InitByReferencingURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)

	return NSImage_FromPointer(ret)
}

// InitWithContentsOfFile initializes and returns an image object with the contents of the specified file.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519918-initwithcontentsoffile?language=objc for details.
func (x gen_NSImage) InitWithContentsOfFile(
	fileName core.NSStringRef,
) NSImage {
	ret := C.NSImage_inst_InitWithContentsOfFile(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fileName),
	)

	return NSImage_FromPointer(ret)
}

// InitWithContentsOfURL initializes and returns an image object with the contents of the specified URL.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519907-initwithcontentsofurl?language=objc for details.
func (x gen_NSImage) InitWithContentsOfURL(
	url core.NSURLRef,
) NSImage {
	ret := C.NSImage_inst_InitWithContentsOfURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)

	return NSImage_FromPointer(ret)
}

// InitWithData initializes and returns an image object using the provided image data.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519941-initwithdata?language=objc for details.
func (x gen_NSImage) InitWithData(
	data core.NSDataRef,
) NSImage {
	ret := C.NSImage_inst_InitWithData(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
	)

	return NSImage_FromPointer(ret)
}

// InitWithDataIgnoringOrientation initializes and returns an image object using the provided image data and ignoring the EXIF orientation tags.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519915-initwithdataignoringorientation?language=objc for details.
func (x gen_NSImage) InitWithDataIgnoringOrientation(
	data core.NSDataRef,
) NSImage {
	ret := C.NSImage_inst_InitWithDataIgnoringOrientation(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
	)

	return NSImage_FromPointer(ret)
}

// InitWithPasteboard initializes and returns an image object with data from the specified pasteboard.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519952-initwithpasteboard?language=objc for details.
func (x gen_NSImage) InitWithPasteboard(
	pasteboard NSPasteboardRef,
) NSImage {
	ret := C.NSImage_inst_InitWithPasteboard(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pasteboard),
	)

	return NSImage_FromPointer(ret)
}

// InitWithSize initializes and returns an image object with the specified dimensions.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1520033-initwithsize?language=objc for details.
func (x gen_NSImage) InitWithSize(
	size core.NSSize,
) NSImage {
	ret := C.NSImage_inst_InitWithSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return NSImage_FromPointer(ret)
}

// IsTemplate returns a Boolean value that indicates whether the image is a template image.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1807274-istemplate?language=objc for details.
func (x gen_NSImage) IsTemplate() bool {
	ret := C.NSImage_inst_IsTemplate(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// LayerContentsForContentsScale returns an object that may be used as the contents of a layer.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519851-layercontentsforcontentsscale?language=objc for details.
func (x gen_NSImage) LayerContentsForContentsScale(
	layerContentsScale core.CGFloat,
) objc.Object {
	ret := C.NSImage_inst_LayerContentsForContentsScale(
		unsafe.Pointer(x.Pointer()),
		C.double(layerContentsScale),
	)

	return objc.Object_FromPointer(ret)
}

// LockFocus prepares the image to receive drawing commands.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519891-lockfocus?language=objc for details.
func (x gen_NSImage) LockFocus() {
	C.NSImage_inst_LockFocus(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// LockFocusFlipped prepares the image to receive drawing commands using the specified flipped state.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519914-lockfocusflipped?language=objc for details.
func (x gen_NSImage) LockFocusFlipped(
	flipped bool,
) {
	C.NSImage_inst_LockFocusFlipped(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flipped),
	)

	return
}

// Recache invalidates and frees offscreen caches of all image representations.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519890-recache?language=objc for details.
func (x gen_NSImage) Recache() {
	C.NSImage_inst_Recache(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// RecommendedLayerContentsScale returns the recommended layer contents scale for this image.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519878-recommendedlayercontentsscale?language=objc for details.
func (x gen_NSImage) RecommendedLayerContentsScale(
	preferredContentsScale core.CGFloat,
) core.CGFloat {
	ret := C.NSImage_inst_RecommendedLayerContentsScale(
		unsafe.Pointer(x.Pointer()),
		C.double(preferredContentsScale),
	)

	return core.CGFloat(ret)
}

// UnlockFocus removes the focus from the image.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519853-unlockfocus?language=objc for details.
func (x gen_NSImage) UnlockFocus() {
	C.NSImage_inst_UnlockFocus(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Init is undocumented.
func (x gen_NSImage) Init() NSImage {
	ret := C.NSImage_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_FromPointer(ret)
}

// Delegate returns the image’s delegate object.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519926-delegate?language=objc for details.
func (x gen_NSImage) Delegate() objc.Object {
	ret := C.NSImage_inst_Delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetDelegate returns the image’s delegate object.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519926-delegate?language=objc for details.
func (x gen_NSImage) SetDelegate(
	value objc.Ref,
) {
	C.NSImage_inst_SetDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Size returns the size of the image.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519987-size?language=objc for details.
func (x gen_NSImage) Size() core.NSSize {
	ret := C.NSImage_inst_Size(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetSize returns the size of the image.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519987-size?language=objc for details.
func (x gen_NSImage) SetSize(
	value core.NSSize,
) {
	C.NSImage_inst_SetSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// SetTemplate returns a Boolean value that determines whether the image represents a template image.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1520017-template?language=objc for details.
func (x gen_NSImage) SetTemplate(
	value bool,
) {
	C.NSImage_inst_SetTemplate(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// Representations an array containing all of the image object’s image representations.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519858-representations?language=objc for details.
func (x gen_NSImage) Representations() core.NSArray {
	ret := C.NSImage_inst_Representations(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// PrefersColorMatch returns a Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1520010-preferscolormatch?language=objc for details.
func (x gen_NSImage) PrefersColorMatch() bool {
	ret := C.NSImage_inst_PrefersColorMatch(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetPrefersColorMatch returns a Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1520010-preferscolormatch?language=objc for details.
func (x gen_NSImage) SetPrefersColorMatch(
	value bool,
) {
	C.NSImage_inst_SetPrefersColorMatch(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// UsesEPSOnResolutionMismatch returns a Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519868-usesepsonresolutionmismatch?language=objc for details.
func (x gen_NSImage) UsesEPSOnResolutionMismatch() bool {
	ret := C.NSImage_inst_UsesEPSOnResolutionMismatch(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetUsesEPSOnResolutionMismatch returns a Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519868-usesepsonresolutionmismatch?language=objc for details.
func (x gen_NSImage) SetUsesEPSOnResolutionMismatch(
	value bool,
) {
	C.NSImage_inst_SetUsesEPSOnResolutionMismatch(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// MatchesOnMultipleResolution returns a Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519963-matchesonmultipleresolution?language=objc for details.
func (x gen_NSImage) MatchesOnMultipleResolution() bool {
	ret := C.NSImage_inst_MatchesOnMultipleResolution(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetMatchesOnMultipleResolution returns a Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519963-matchesonmultipleresolution?language=objc for details.
func (x gen_NSImage) SetMatchesOnMultipleResolution(
	value bool,
) {
	C.NSImage_inst_SetMatchesOnMultipleResolution(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsValid returns a Boolean value that indicates whether it is possible to draw an image representation.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519991-valid?language=objc for details.
func (x gen_NSImage) IsValid() bool {
	ret := C.NSImage_inst_IsValid(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// BackgroundColor returns the background color for the image.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1520059-backgroundcolor?language=objc for details.
func (x gen_NSImage) BackgroundColor() NSColor {
	ret := C.NSImage_inst_BackgroundColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_FromPointer(ret)
}

// SetBackgroundColor returns the background color for the image.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1520059-backgroundcolor?language=objc for details.
func (x gen_NSImage) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSImage_inst_SetBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// AlignmentRect returns a rectangle that you can use to position the image during layout.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519905-alignmentrect?language=objc for details.
func (x gen_NSImage) AlignmentRect() core.NSRect {
	ret := C.NSImage_inst_AlignmentRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// SetAlignmentRect returns a rectangle that you can use to position the image during layout.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519905-alignmentrect?language=objc for details.
func (x gen_NSImage) SetAlignmentRect(
	value core.NSRect,
) {
	C.NSImage_inst_SetAlignmentRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return
}

// TIFFRepresentation returns a data object containing TIFF data for all of the image representations in the image.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519841-tiffrepresentation?language=objc for details.
func (x gen_NSImage) TIFFRepresentation() core.NSData {
	ret := C.NSImage_inst_TIFFRepresentation(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSData_FromPointer(ret)
}

// AccessibilityDescription returns the image’s accessibility description.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519943-accessibilitydescription?language=objc for details.
func (x gen_NSImage) AccessibilityDescription() core.NSString {
	ret := C.NSImage_inst_AccessibilityDescription(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetAccessibilityDescription returns the image’s accessibility description.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519943-accessibilitydescription?language=objc for details.
func (x gen_NSImage) SetAccessibilityDescription(
	value core.NSStringRef,
) {
	C.NSImage_inst_SetAccessibilityDescription(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// MatchesOnlyOnBestFittingAxis returns a Boolean value that indicates whether the image matches only on the best fitting axis.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519848-matchesonlyonbestfittingaxis?language=objc for details.
func (x gen_NSImage) MatchesOnlyOnBestFittingAxis() bool {
	ret := C.NSImage_inst_MatchesOnlyOnBestFittingAxis(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetMatchesOnlyOnBestFittingAxis returns a Boolean value that indicates whether the image matches only on the best fitting axis.
//
// See https://developer.apple.com/documentation/appkit/nsimage/1519848-matchesonlyonbestfittingaxis?language=objc for details.
func (x gen_NSImage) SetMatchesOnlyOnBestFittingAxis(
	value bool,
) {
	C.NSImage_inst_SetMatchesOnlyOnBestFittingAxis(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

type NSImageViewRef interface {
	Pointer() uintptr
	Init() NSImageView
}

type gen_NSImageView struct {
	NSControl
}

func NSImageView_FromPointer(ptr unsafe.Pointer) NSImageView {
	return NSImageView{gen_NSImageView{
		NSControl_FromPointer(ptr),
	}}
}

func NSImageView_FromRef(ref objc.Ref) NSImageView {
	return NSImageView_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init is undocumented.
func (x gen_NSImageView) Init() NSImageView {
	ret := C.NSImageView_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImageView_FromPointer(ret)
}

// Image returns the image displayed by the image view.
//
// See https://developer.apple.com/documentation/appkit/nsimageview/1404952-image?language=objc for details.
func (x gen_NSImageView) Image() NSImage {
	ret := C.NSImageView_inst_Image(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_FromPointer(ret)
}

// SetImage returns the image displayed by the image view.
//
// See https://developer.apple.com/documentation/appkit/nsimageview/1404952-image?language=objc for details.
func (x gen_NSImageView) SetImage(
	value NSImageRef,
) {
	C.NSImageView_inst_SetImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Animates returns a Boolean value indicating whether the image view automatically plays animated images.
//
// See https://developer.apple.com/documentation/appkit/nsimageview/1404950-animates?language=objc for details.
func (x gen_NSImageView) Animates() bool {
	ret := C.NSImageView_inst_Animates(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAnimates returns a Boolean value indicating whether the image view automatically plays animated images.
//
// See https://developer.apple.com/documentation/appkit/nsimageview/1404950-animates?language=objc for details.
func (x gen_NSImageView) SetAnimates(
	value bool,
) {
	C.NSImageView_inst_SetAnimates(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsEditable returns a Boolean value indicating whether the user can drag a new image into the image view.
//
// See https://developer.apple.com/documentation/appkit/nsimageview/1404954-editable?language=objc for details.
func (x gen_NSImageView) IsEditable() bool {
	ret := C.NSImageView_inst_IsEditable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetEditable returns a Boolean value indicating whether the user can drag a new image into the image view.
//
// See https://developer.apple.com/documentation/appkit/nsimageview/1404954-editable?language=objc for details.
func (x gen_NSImageView) SetEditable(
	value bool,
) {
	C.NSImageView_inst_SetEditable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AllowsCutCopyPaste returns a Boolean value indicating whether the image view lets the user cut, copy, and paste the image contents.
//
// See https://developer.apple.com/documentation/appkit/nsimageview/1404961-allowscutcopypaste?language=objc for details.
func (x gen_NSImageView) AllowsCutCopyPaste() bool {
	ret := C.NSImageView_inst_AllowsCutCopyPaste(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsCutCopyPaste returns a Boolean value indicating whether the image view lets the user cut, copy, and paste the image contents.
//
// See https://developer.apple.com/documentation/appkit/nsimageview/1404961-allowscutcopypaste?language=objc for details.
func (x gen_NSImageView) SetAllowsCutCopyPaste(
	value bool,
) {
	C.NSImageView_inst_SetAllowsCutCopyPaste(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// ContentTintColor is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsimageview/3000783-contenttintcolor?language=objc for details.
func (x gen_NSImageView) ContentTintColor() NSColor {
	ret := C.NSImageView_inst_ContentTintColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_FromPointer(ret)
}

// SetContentTintColor is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsimageview/3000783-contenttintcolor?language=objc for details.
func (x gen_NSImageView) SetContentTintColor(
	value NSColorRef,
) {
	C.NSImageView_inst_SetContentTintColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

type NSNibRef interface {
	Pointer() uintptr
	Init() NSNib
}

type gen_NSNib struct {
	objc.Object
}

func NSNib_FromPointer(ptr unsafe.Pointer) NSNib {
	return NSNib{gen_NSNib{
		objc.Object_FromPointer(ptr),
	}}
}

func NSNib_FromRef(ref objc.Ref) NSNib {
	return NSNib_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// InitWithNibDataBundle initializes an instance with nib data and specified bundle for locating resources.
//
// See https://developer.apple.com/documentation/appkit/nsnib/1535865-initwithnibdata?language=objc for details.
func (x gen_NSNib) InitWithNibDataBundle(
	nibData core.NSDataRef,
	bundle NSBundleRef,
) NSNib {
	ret := C.NSNib_inst_InitWithNibDataBundle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(nibData),
		objc.RefPointer(bundle),
	)

	return NSNib_FromPointer(ret)
}

// InstantiateWithOwnerTopLevelObjects instantiates objects in the nib file with the specified owner.
//
// See https://developer.apple.com/documentation/appkit/nsnib/1527173-instantiatewithowner?language=objc for details.
func (x gen_NSNib) InstantiateWithOwnerTopLevelObjects(
	owner objc.Ref,
	topLevelObjects core.NSArrayRef,
) bool {
	ret := C.NSNib_inst_InstantiateWithOwnerTopLevelObjects(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(owner),
		objc.RefPointer(topLevelObjects),
	)

	return convertObjCBoolToGo(ret)
}

// Init is undocumented.
func (x gen_NSNib) Init() NSNib {
	ret := C.NSNib_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSNib_FromPointer(ret)
}

type NSPasteboardRef interface {
	Pointer() uintptr
	Init() NSPasteboard
}

type gen_NSPasteboard struct {
	objc.Object
}

func NSPasteboard_FromPointer(ptr unsafe.Pointer) NSPasteboard {
	return NSPasteboard{gen_NSPasteboard{
		objc.Object_FromPointer(ptr),
	}}
}

func NSPasteboard_FromRef(ref objc.Ref) NSPasteboard {
	return NSPasteboard_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// AddTypesOwner adds promises for the specified types to the first pasteboard item.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1533580-addtypes?language=objc for details.
func (x gen_NSPasteboard) AddTypesOwner(
	newTypes core.NSArrayRef,
	newOwner objc.Ref,
) core.NSInteger {
	ret := C.NSPasteboard_inst_AddTypesOwner(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newTypes),
		objc.RefPointer(newOwner),
	)

	return core.NSInteger(ret)
}

// CanReadItemWithDataConformingToTypes returns a Boolean value that indicates whether the receiver contains any items that conform to the specified UTIs.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1533576-canreaditemwithdataconformingtot?language=objc for details.
func (x gen_NSPasteboard) CanReadItemWithDataConformingToTypes(
	types core.NSArrayRef,
) bool {
	ret := C.NSPasteboard_inst_CanReadItemWithDataConformingToTypes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(types),
	)

	return convertObjCBoolToGo(ret)
}

// CanReadObjectForClassesOptions returns a Boolean value that indicates whether the receiver contains any items that can be represented as an instance of any class in a given array.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1533360-canreadobjectforclasses?language=objc for details.
func (x gen_NSPasteboard) CanReadObjectForClassesOptions(
	classArray core.NSArrayRef,
	options core.NSDictionaryRef,
) bool {
	ret := C.NSPasteboard_inst_CanReadObjectForClassesOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(classArray),
		objc.RefPointer(options),
	)

	return convertObjCBoolToGo(ret)
}

// ClearContents clears the existing contents of the pasteboard.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1533599-clearcontents?language=objc for details.
func (x gen_NSPasteboard) ClearContents() core.NSInteger {
	ret := C.NSPasteboard_inst_ClearContents(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// DeclareTypesOwner prepares the receiver for a change in its contents by declaring the new types of data it will contain and a new owner.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1533561-declaretypes?language=objc for details.
func (x gen_NSPasteboard) DeclareTypesOwner(
	newTypes core.NSArrayRef,
	newOwner objc.Ref,
) core.NSInteger {
	ret := C.NSPasteboard_inst_DeclareTypesOwner(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newTypes),
		objc.RefPointer(newOwner),
	)

	return core.NSInteger(ret)
}

// ReadObjectsForClassesOptions reads from the receiver objects that best match the specified array of classes.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1524454-readobjectsforclasses?language=objc for details.
func (x gen_NSPasteboard) ReadObjectsForClassesOptions(
	classArray core.NSArrayRef,
	options core.NSDictionaryRef,
) core.NSArray {
	ret := C.NSPasteboard_inst_ReadObjectsForClassesOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(classArray),
		objc.RefPointer(options),
	)

	return core.NSArray_FromPointer(ret)
}

// ReleaseGlobally releases the receiver’s resources in the pasteboard server.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1527044-releaseglobally?language=objc for details.
func (x gen_NSPasteboard) ReleaseGlobally() {
	C.NSPasteboard_inst_ReleaseGlobally(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// WriteFileContents writes the contents of the specified file to the pasteboard.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1531224-writefilecontents?language=objc for details.
func (x gen_NSPasteboard) WriteFileContents(
	filename core.NSStringRef,
) bool {
	ret := C.NSPasteboard_inst_WriteFileContents(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(filename),
	)

	return convertObjCBoolToGo(ret)
}

// WriteObjects writes an array of objects to the receiver.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1525945-writeobjects?language=objc for details.
func (x gen_NSPasteboard) WriteObjects(
	objects core.NSArrayRef,
) bool {
	ret := C.NSPasteboard_inst_WriteObjects(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(objects),
	)

	return convertObjCBoolToGo(ret)
}

// Init is undocumented.
func (x gen_NSPasteboard) Init() NSPasteboard {
	ret := C.NSPasteboard_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSPasteboard_FromPointer(ret)
}

// PasteboardItems an array that contains all the items held by the pasteboard.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1529995-pasteboarditems?language=objc for details.
func (x gen_NSPasteboard) PasteboardItems() core.NSArray {
	ret := C.NSPasteboard_inst_PasteboardItems(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// Types an array of the receiver’s supported data types.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1529599-types?language=objc for details.
func (x gen_NSPasteboard) Types() core.NSArray {
	ret := C.NSPasteboard_inst_Types(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// ChangeCount returns the receiver’s change count.
//
// See https://developer.apple.com/documentation/appkit/nspasteboard/1533544-changecount?language=objc for details.
func (x gen_NSPasteboard) ChangeCount() core.NSInteger {
	ret := C.NSPasteboard_inst_ChangeCount(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

type NSLayoutManagerRef interface {
	Pointer() uintptr
	Init() NSLayoutManager
}

type gen_NSLayoutManager struct {
	objc.Object
}

func NSLayoutManager_FromPointer(ptr unsafe.Pointer) NSLayoutManager {
	return NSLayoutManager{gen_NSLayoutManager{
		objc.Object_FromPointer(ptr),
	}}
}

func NSLayoutManager_FromRef(ref objc.Ref) NSLayoutManager {
	return NSLayoutManager_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// AddTextContainer appends the specified text container to the series of text containers where the layout manager arranges text.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1402946-addtextcontainer?language=objc for details.
func (x gen_NSLayoutManager) AddTextContainer(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_AddTextContainer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)

	return
}

// AttachmentSizeForGlyphAtIndex returns the size of the attachment glyph at the specified index.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403099-attachmentsizeforglyphatindex?language=objc for details.
func (x gen_NSLayoutManager) AttachmentSizeForGlyphAtIndex(
	glyphIndex core.NSUInteger,
) core.NSSize {
	ret := C.NSLayoutManager_inst_AttachmentSizeForGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// CharacterIndexForGlyphAtIndex returns the index in the text storage for the first character of the specified glyph.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1402944-characterindexforglyphatindex?language=objc for details.
func (x gen_NSLayoutManager) CharacterIndexForGlyphAtIndex(
	glyphIndex core.NSUInteger,
) core.NSUInteger {
	ret := C.NSLayoutManager_inst_CharacterIndexForGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)

	return core.NSUInteger(ret)
}

// DefaultBaselineOffsetForFont returns the default baseline offset that the layout manager's typesetter uses for the specified font.
//
// See https://developer.apple.com/documentation/appkit/nslayoutmanager/1403058-defaultbaselineoffsetforfont?language=objc for details.
func (x gen_NSLayoutManager) DefaultBaselineOffsetForFont(
	theFont NSFontRef,
) core.CGFloat {
	ret := C.NSLayoutManager_inst_DefaultBaselineOffsetForFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(theFont),
	)

	return core.CGFloat(ret)
}

// DefaultLineHeightForFont returns the default line height for a line of text that uses a specified font.
//
// See https://developer.apple.com/documentation/appkit/nslayoutmanager/1403007-defaultlineheightforfont?language=objc for details.
func (x gen_NSLayoutManager) DefaultLineHeightForFont(
	theFont NSFontRef,
) core.CGFloat {
	ret := C.NSLayoutManager_inst_DefaultLineHeightForFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(theFont),
	)

	return core.CGFloat(ret)
}

// DrawsOutsideLineFragmentForGlyphAtIndex indicates whether the glyph draws outside its line fragment rectangle.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403003-drawsoutsidelinefragmentforglyph?language=objc for details.
func (x gen_NSLayoutManager) DrawsOutsideLineFragmentForGlyphAtIndex(
	glyphIndex core.NSUInteger,
) bool {
	ret := C.NSLayoutManager_inst_DrawsOutsideLineFragmentForGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)

	return convertObjCBoolToGo(ret)
}

// EnsureLayoutForBoundingRectInTextContainer forces the layout manager to perform layout for the specified area in the specified text container if it hasn’t already.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1402962-ensurelayoutforboundingrect?language=objc for details.
func (x gen_NSLayoutManager) EnsureLayoutForBoundingRectInTextContainer(
	bounds core.NSRect,
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_EnsureLayoutForBoundingRectInTextContainer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&bounds)),
		objc.RefPointer(container),
	)

	return
}

// EnsureLayoutForTextContainer forces the layout manager to perform layout for the specified text container if it hasn’t already.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1402967-ensurelayoutfortextcontainer?language=objc for details.
func (x gen_NSLayoutManager) EnsureLayoutForTextContainer(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_EnsureLayoutForTextContainer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)

	return
}

// FirstUnlaidCharacterIndex returns the index for the first character in the layout manager that isn’t in the layout.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403067-firstunlaidcharacterindex?language=objc for details.
func (x gen_NSLayoutManager) FirstUnlaidCharacterIndex() core.NSUInteger {
	ret := C.NSLayoutManager_inst_FirstUnlaidCharacterIndex(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// FirstUnlaidGlyphIndex returns the index for the first glyph in the layout manager that isn’t in the layout.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403245-firstunlaidglyphindex?language=objc for details.
func (x gen_NSLayoutManager) FirstUnlaidGlyphIndex() core.NSUInteger {
	ret := C.NSLayoutManager_inst_FirstUnlaidGlyphIndex(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// GlyphIndexForCharacterAtIndex returns the index of the first glyph of the character at the specified index.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403001-glyphindexforcharacteratindex?language=objc for details.
func (x gen_NSLayoutManager) GlyphIndexForCharacterAtIndex(
	charIndex core.NSUInteger,
) core.NSUInteger {
	ret := C.NSLayoutManager_inst_GlyphIndexForCharacterAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(charIndex),
	)

	return core.NSUInteger(ret)
}

// Init initializes a newly created layout manager object.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1402975-init?language=objc for details.
func (x gen_NSLayoutManager) Init() NSLayoutManager {
	ret := C.NSLayoutManager_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSLayoutManager_FromPointer(ret)
}

// InsertTextContainerAtIndex inserts a text container at the specified index in the list of text containers.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403010-inserttextcontainer?language=objc for details.
func (x gen_NSLayoutManager) InsertTextContainerAtIndex(
	container NSTextContainerRef,
	index core.NSUInteger,
) {
	C.NSLayoutManager_inst_InsertTextContainerAtIndex(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
		C.ulong(index),
	)

	return
}

// IsValidGlyphIndex indicates whether the specified index refers to a valid glyph.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1402950-isvalidglyphindex?language=objc for details.
func (x gen_NSLayoutManager) IsValidGlyphIndex(
	glyphIndex core.NSUInteger,
) bool {
	ret := C.NSLayoutManager_inst_IsValidGlyphIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)

	return convertObjCBoolToGo(ret)
}

// LayoutManagerOwnsFirstResponderInWindow indicates whether the first responder in the specified window is a text view for the layout manager.
//
// See https://developer.apple.com/documentation/appkit/nslayoutmanager/1403026-layoutmanagerownsfirstresponderi?language=objc for details.
func (x gen_NSLayoutManager) LayoutManagerOwnsFirstResponderInWindow(
	window NSWindowRef,
) bool {
	ret := C.NSLayoutManager_inst_LayoutManagerOwnsFirstResponderInWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(window),
	)

	return convertObjCBoolToGo(ret)
}

// NotShownAttributeForGlyphAtIndex indicates whether the glyph at the specified index has a visible representation.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1402931-notshownattributeforglyphatindex?language=objc for details.
func (x gen_NSLayoutManager) NotShownAttributeForGlyphAtIndex(
	glyphIndex core.NSUInteger,
) bool {
	ret := C.NSLayoutManager_inst_NotShownAttributeForGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)

	return convertObjCBoolToGo(ret)
}

// RemoveTextContainerAtIndex removes the text container at the specified index and invalidates the layout as necessary.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403017-removetextcontaineratindex?language=objc for details.
func (x gen_NSLayoutManager) RemoveTextContainerAtIndex(
	index core.NSUInteger,
) {
	C.NSLayoutManager_inst_RemoveTextContainerAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(index),
	)

	return
}

// SetDrawsOutsideLineFragmentForGlyphAtIndex indicates whether the specified glyph exceeds the bounds of the line fragment for its layout.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1402964-setdrawsoutsidelinefragment?language=objc for details.
func (x gen_NSLayoutManager) SetDrawsOutsideLineFragmentForGlyphAtIndex(
	flag bool,
	glyphIndex core.NSUInteger,
) {
	C.NSLayoutManager_inst_SetDrawsOutsideLineFragmentForGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
		C.ulong(glyphIndex),
	)

	return
}

// SetExtraLineFragmentRectUsedRectTextContainer sets the bounds and container for the extra line fragment.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403071-setextralinefragmentrect?language=objc for details.
func (x gen_NSLayoutManager) SetExtraLineFragmentRectUsedRectTextContainer(
	fragmentRect core.NSRect,
	usedRect core.NSRect,
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_SetExtraLineFragmentRectUsedRectTextContainer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&fragmentRect)),
		*(*C.NSRect)(unsafe.Pointer(&usedRect)),
		objc.RefPointer(container),
	)

	return
}

// SetNotShownAttributeForGlyphAtIndex sets the visibility of the glyph at the specified index.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403078-setnotshownattribute?language=objc for details.
func (x gen_NSLayoutManager) SetNotShownAttributeForGlyphAtIndex(
	flag bool,
	glyphIndex core.NSUInteger,
) {
	C.NSLayoutManager_inst_SetNotShownAttributeForGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
		C.ulong(glyphIndex),
	)

	return
}

// TextContainerChangedGeometry invalidates the layout information, and possibly glyphs, for the specified text container and all subsequent text container objects.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403091-textcontainerchangedgeometry?language=objc for details.
func (x gen_NSLayoutManager) TextContainerChangedGeometry(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_TextContainerChangedGeometry(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)

	return
}

// TextContainerChangedTextView updates the information necessary to manage text view objects for the specified text container.
//
// See https://developer.apple.com/documentation/appkit/nslayoutmanager/1403229-textcontainerchangedtextview?language=objc for details.
func (x gen_NSLayoutManager) TextContainerChangedTextView(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_TextContainerChangedTextView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)

	return
}

// UsedRectForTextContainer returns the bounding rectangle for the glyphs in the specified text container.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1402980-usedrectfortextcontainer?language=objc for details.
func (x gen_NSLayoutManager) UsedRectForTextContainer(
	container NSTextContainerRef,
) core.NSRect {
	ret := C.NSLayoutManager_inst_UsedRectForTextContainer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// Delegate returns the layout manager’s delegate.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1402920-delegate?language=objc for details.
func (x gen_NSLayoutManager) Delegate() objc.Object {
	ret := C.NSLayoutManager_inst_Delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetDelegate returns the layout manager’s delegate.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1402920-delegate?language=objc for details.
func (x gen_NSLayoutManager) SetDelegate(
	value objc.Ref,
) {
	C.NSLayoutManager_inst_SetDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// AllowsNonContiguousLayout returns a Boolean value that indicates whether the layout manager allows noncontiguous layout.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403197-allowsnoncontiguouslayout?language=objc for details.
func (x gen_NSLayoutManager) AllowsNonContiguousLayout() bool {
	ret := C.NSLayoutManager_inst_AllowsNonContiguousLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsNonContiguousLayout returns a Boolean value that indicates whether the layout manager allows noncontiguous layout.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403197-allowsnoncontiguouslayout?language=objc for details.
func (x gen_NSLayoutManager) SetAllowsNonContiguousLayout(
	value bool,
) {
	C.NSLayoutManager_inst_SetAllowsNonContiguousLayout(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// HasNonContiguousLayout returns a Boolean value that indicates whether the layout manager currently has any areas of noncontiguous layout.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403207-hasnoncontiguouslayout?language=objc for details.
func (x gen_NSLayoutManager) HasNonContiguousLayout() bool {
	ret := C.NSLayoutManager_inst_HasNonContiguousLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// ShowsInvisibleCharacters returns a Boolean value that indicates whether to substitute visible glyphs for whitespace and other typically invisible characters.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403254-showsinvisiblecharacters?language=objc for details.
func (x gen_NSLayoutManager) ShowsInvisibleCharacters() bool {
	ret := C.NSLayoutManager_inst_ShowsInvisibleCharacters(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetShowsInvisibleCharacters returns a Boolean value that indicates whether to substitute visible glyphs for whitespace and other typically invisible characters.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403254-showsinvisiblecharacters?language=objc for details.
func (x gen_NSLayoutManager) SetShowsInvisibleCharacters(
	value bool,
) {
	C.NSLayoutManager_inst_SetShowsInvisibleCharacters(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// ShowsControlCharacters returns a Boolean value that indicates whether the layout manager substitutes visible glyphs for control characters in the layout.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1402912-showscontrolcharacters?language=objc for details.
func (x gen_NSLayoutManager) ShowsControlCharacters() bool {
	ret := C.NSLayoutManager_inst_ShowsControlCharacters(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetShowsControlCharacters returns a Boolean value that indicates whether the layout manager substitutes visible glyphs for control characters in the layout.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1402912-showscontrolcharacters?language=objc for details.
func (x gen_NSLayoutManager) SetShowsControlCharacters(
	value bool,
) {
	C.NSLayoutManager_inst_SetShowsControlCharacters(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// UsesFontLeading returns a Boolean value that indicates whether the layout manager uses the leading of the font.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403156-usesfontleading?language=objc for details.
func (x gen_NSLayoutManager) UsesFontLeading() bool {
	ret := C.NSLayoutManager_inst_UsesFontLeading(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetUsesFontLeading returns a Boolean value that indicates whether the layout manager uses the leading of the font.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403156-usesfontleading?language=objc for details.
func (x gen_NSLayoutManager) SetUsesFontLeading(
	value bool,
) {
	C.NSLayoutManager_inst_SetUsesFontLeading(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// BackgroundLayoutEnabled returns a Boolean value that indicates whether the layout manager generates glyphs and lays them out when the app's run loop is idle.
//
// See https://developer.apple.com/documentation/appkit/nslayoutmanager/1402952-backgroundlayoutenabled?language=objc for details.
func (x gen_NSLayoutManager) BackgroundLayoutEnabled() bool {
	ret := C.NSLayoutManager_inst_BackgroundLayoutEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetBackgroundLayoutEnabled returns a Boolean value that indicates whether the layout manager generates glyphs and lays them out when the app's run loop is idle.
//
// See https://developer.apple.com/documentation/appkit/nslayoutmanager/1402952-backgroundlayoutenabled?language=objc for details.
func (x gen_NSLayoutManager) SetBackgroundLayoutEnabled(
	value bool,
) {
	C.NSLayoutManager_inst_SetBackgroundLayoutEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// LimitsLayoutForSuspiciousContents returns a Boolean value that indicates whether the layout manager avoids laying out unusually long or suspicious input.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/3021179-limitslayoutforsuspiciouscontent?language=objc for details.
func (x gen_NSLayoutManager) LimitsLayoutForSuspiciousContents() bool {
	ret := C.NSLayoutManager_inst_LimitsLayoutForSuspiciousContents(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetLimitsLayoutForSuspiciousContents returns a Boolean value that indicates whether the layout manager avoids laying out unusually long or suspicious input.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/3021179-limitslayoutforsuspiciouscontent?language=objc for details.
func (x gen_NSLayoutManager) SetLimitsLayoutForSuspiciousContents(
	value bool,
) {
	C.NSLayoutManager_inst_SetLimitsLayoutForSuspiciousContents(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// UsesDefaultHyphenation returns a Boolean value that indicates whether the layout manager uses the default hyphenation rules to wrap lines.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/3180380-usesdefaulthyphenation?language=objc for details.
func (x gen_NSLayoutManager) UsesDefaultHyphenation() bool {
	ret := C.NSLayoutManager_inst_UsesDefaultHyphenation(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetUsesDefaultHyphenation returns a Boolean value that indicates whether the layout manager uses the default hyphenation rules to wrap lines.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/3180380-usesdefaulthyphenation?language=objc for details.
func (x gen_NSLayoutManager) SetUsesDefaultHyphenation(
	value bool,
) {
	C.NSLayoutManager_inst_SetUsesDefaultHyphenation(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// TextContainers returns the current text containers of the layout manager.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403144-textcontainers?language=objc for details.
func (x gen_NSLayoutManager) TextContainers() core.NSArray {
	ret := C.NSLayoutManager_inst_TextContainers(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// NumberOfGlyphs returns the number of glyphs in the layout manager.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1402937-numberofglyphs?language=objc for details.
func (x gen_NSLayoutManager) NumberOfGlyphs() core.NSUInteger {
	ret := C.NSLayoutManager_inst_NumberOfGlyphs(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// ExtraLineFragmentRect returns the rectangle for the extra line fragment at the end of a document.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403175-extralinefragmentrect?language=objc for details.
func (x gen_NSLayoutManager) ExtraLineFragmentRect() core.NSRect {
	ret := C.NSLayoutManager_inst_ExtraLineFragmentRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// ExtraLineFragmentTextContainer returns the text container for the extra line fragment rectangle.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1403165-extralinefragmenttextcontainer?language=objc for details.
func (x gen_NSLayoutManager) ExtraLineFragmentTextContainer() NSTextContainer {
	ret := C.NSLayoutManager_inst_ExtraLineFragmentTextContainer(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextContainer_FromPointer(ret)
}

// ExtraLineFragmentUsedRect returns the rectangle that encloses the insertion point in the extra line fragment rectangle.
//
// See https://developer.apple.com/documentation/uikit/nslayoutmanager/1402988-extralinefragmentusedrect?language=objc for details.
func (x gen_NSLayoutManager) ExtraLineFragmentUsedRect() core.NSRect {
	ret := C.NSLayoutManager_inst_ExtraLineFragmentUsedRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// FirstTextView returns the first text view in the layout manager’s series of text views.
//
// See https://developer.apple.com/documentation/appkit/nslayoutmanager/1402995-firsttextview?language=objc for details.
func (x gen_NSLayoutManager) FirstTextView() NSTextView {
	ret := C.NSLayoutManager_inst_FirstTextView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextView_FromPointer(ret)
}

// TextViewForBeginningOfSelection returns the text view that contains the first glyph in the selection.
//
// See https://developer.apple.com/documentation/appkit/nslayoutmanager/1403089-textviewforbeginningofselection?language=objc for details.
func (x gen_NSLayoutManager) TextViewForBeginningOfSelection() NSTextView {
	ret := C.NSLayoutManager_inst_TextViewForBeginningOfSelection(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextView_FromPointer(ret)
}

type NSMenuRef interface {
	Pointer() uintptr
	Init() NSMenu
}

type gen_NSMenu struct {
	objc.Object
}

func NSMenu_FromPointer(ptr unsafe.Pointer) NSMenu {
	return NSMenu{gen_NSMenu{
		objc.Object_FromPointer(ptr),
	}}
}

func NSMenu_FromRef(ref objc.Ref) NSMenu {
	return NSMenu_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// AddItem adds a menu item to the end of the menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518176-additem?language=objc for details.
func (x gen_NSMenu) AddItem(
	newItem NSMenuItemRef,
) {
	C.NSMenu_inst_AddItem(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newItem),
	)

	return
}

// AddItemWithTitleActionKeyEquivalent creates a new menu item and adds it to the end of the menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518181-additemwithtitle?language=objc for details.
func (x gen_NSMenu) AddItemWithTitleActionKeyEquivalent(
	string core.NSStringRef,
	selector objc.Selector,
	charCode core.NSStringRef,
) NSMenuItem {
	ret := C.NSMenu_inst_AddItemWithTitleActionKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		selector.SelectorAddress(),
		objc.RefPointer(charCode),
	)

	return NSMenuItem_FromPointer(ret)
}

// CancelTracking dismisses the menu and ends all menu tracking.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518150-canceltracking?language=objc for details.
func (x gen_NSMenu) CancelTracking() {
	C.NSMenu_inst_CancelTracking(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// CancelTrackingWithoutAnimation dismisses the menu and ends all menu tracking without displaying the associated animation.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518244-canceltrackingwithoutanimation?language=objc for details.
func (x gen_NSMenu) CancelTrackingWithoutAnimation() {
	C.NSMenu_inst_CancelTrackingWithoutAnimation(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// IndexOfItem returns the index identifying the location of a specified menu item in the menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518178-indexofitem?language=objc for details.
func (x gen_NSMenu) IndexOfItem(
	item NSMenuItemRef,
) core.NSInteger {
	ret := C.NSMenu_inst_IndexOfItem(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
	)

	return core.NSInteger(ret)
}

// IndexOfItemWithRepresentedObject returns the index of the first menu item in the menu that has a given represented object.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518175-indexofitemwithrepresentedobject?language=objc for details.
func (x gen_NSMenu) IndexOfItemWithRepresentedObject(
	object objc.Ref,
) core.NSInteger {
	ret := C.NSMenu_inst_IndexOfItemWithRepresentedObject(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)

	return core.NSInteger(ret)
}

// IndexOfItemWithSubmenu returns the index of the menu item in the menu with the given submenu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518216-indexofitemwithsubmenu?language=objc for details.
func (x gen_NSMenu) IndexOfItemWithSubmenu(
	submenu NSMenuRef,
) core.NSInteger {
	ret := C.NSMenu_inst_IndexOfItemWithSubmenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(submenu),
	)

	return core.NSInteger(ret)
}

// IndexOfItemWithTag returns the index of the first menu item in the menu identified by a tag.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518164-indexofitemwithtag?language=objc for details.
func (x gen_NSMenu) IndexOfItemWithTag(
	tag core.NSInteger,
) core.NSInteger {
	ret := C.NSMenu_inst_IndexOfItemWithTag(
		unsafe.Pointer(x.Pointer()),
		C.long(tag),
	)

	return core.NSInteger(ret)
}

// IndexOfItemWithTargetAndAction returns the index of the first menu item in the menu that has a specified action and target.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518153-indexofitemwithtarget?language=objc for details.
func (x gen_NSMenu) IndexOfItemWithTargetAndAction(
	target objc.Ref,
	actionSelector objc.Selector,
) core.NSInteger {
	ret := C.NSMenu_inst_IndexOfItemWithTargetAndAction(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(target),
		actionSelector.SelectorAddress(),
	)

	return core.NSInteger(ret)
}

// IndexOfItemWithTitle returns the index of the first menu item in the menu that has a specified title.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518237-indexofitemwithtitle?language=objc for details.
func (x gen_NSMenu) IndexOfItemWithTitle(
	title core.NSStringRef,
) core.NSInteger {
	ret := C.NSMenu_inst_IndexOfItemWithTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(title),
	)

	return core.NSInteger(ret)
}

// InitWithTitle initializes and returns a menu having the specified title and with autoenabling of menu items turned on.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518144-initwithtitle?language=objc for details.
func (x gen_NSMenu) InitWithTitle(
	title core.NSStringRef,
) NSMenu {
	ret := C.NSMenu_inst_InitWithTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(title),
	)

	return NSMenu_FromPointer(ret)
}

// InsertItemAtIndex inserts a menu item into the menu at a specific location.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518201-insertitem?language=objc for details.
func (x gen_NSMenu) InsertItemAtIndex(
	newItem NSMenuItemRef,
	index core.NSInteger,
) {
	C.NSMenu_inst_InsertItemAtIndex(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newItem),
		C.long(index),
	)

	return
}

// InsertItemWithTitleActionKeyEquivalentAtIndex creates and adds a menu item at a specified location in the menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518146-insertitemwithtitle?language=objc for details.
func (x gen_NSMenu) InsertItemWithTitleActionKeyEquivalentAtIndex(
	string core.NSStringRef,
	selector objc.Selector,
	charCode core.NSStringRef,
	index core.NSInteger,
) NSMenuItem {
	ret := C.NSMenu_inst_InsertItemWithTitleActionKeyEquivalentAtIndex(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		selector.SelectorAddress(),
		objc.RefPointer(charCode),
		C.long(index),
	)

	return NSMenuItem_FromPointer(ret)
}

// ItemAtIndex returns the menu item at a specific location of the menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518218-itematindex?language=objc for details.
func (x gen_NSMenu) ItemAtIndex(
	index core.NSInteger,
) NSMenuItem {
	ret := C.NSMenu_inst_ItemAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)

	return NSMenuItem_FromPointer(ret)
}

// ItemChanged invoked when a menu item is modified visually (for example, its title changes).
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518154-itemchanged?language=objc for details.
func (x gen_NSMenu) ItemChanged(
	item NSMenuItemRef,
) {
	C.NSMenu_inst_ItemChanged(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
	)

	return
}

// ItemWithTag returns the first menu item in the menu with the specified tag.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518223-itemwithtag?language=objc for details.
func (x gen_NSMenu) ItemWithTag(
	tag core.NSInteger,
) NSMenuItem {
	ret := C.NSMenu_inst_ItemWithTag(
		unsafe.Pointer(x.Pointer()),
		C.long(tag),
	)

	return NSMenuItem_FromPointer(ret)
}

// ItemWithTitle returns the first menu item in the menu with a specified title.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518248-itemwithtitle?language=objc for details.
func (x gen_NSMenu) ItemWithTitle(
	title core.NSStringRef,
) NSMenuItem {
	ret := C.NSMenu_inst_ItemWithTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(title),
	)

	return NSMenuItem_FromPointer(ret)
}

// PerformActionForItemAtIndex causes the application to send the action message of a specified menu item to its target.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518210-performactionforitematindex?language=objc for details.
func (x gen_NSMenu) PerformActionForItemAtIndex(
	index core.NSInteger,
) {
	C.NSMenu_inst_PerformActionForItemAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)

	return
}

// PerformKeyEquivalent performs the action for the menu item that corresponds to the given key equivalent.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518198-performkeyequivalent?language=objc for details.
func (x gen_NSMenu) PerformKeyEquivalent(
	event NSEventRef,
) bool {
	ret := C.NSMenu_inst_PerformKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return convertObjCBoolToGo(ret)
}

// PopUpMenuPositioningItemAtLocationInView pops up the menu at the specified location.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518212-popupmenupositioningitem?language=objc for details.
func (x gen_NSMenu) PopUpMenuPositioningItemAtLocationInView(
	item NSMenuItemRef,
	location core.NSPoint,
	view NSViewRef,
) bool {
	ret := C.NSMenu_inst_PopUpMenuPositioningItemAtLocationInView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
		*(*C.NSPoint)(unsafe.Pointer(&location)),
		objc.RefPointer(view),
	)

	return convertObjCBoolToGo(ret)
}

// RemoveAllItems removes all the menu items in the menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518234-removeallitems?language=objc for details.
func (x gen_NSMenu) RemoveAllItems() {
	C.NSMenu_inst_RemoveAllItems(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// RemoveItem removes a menu item from the menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518257-removeitem?language=objc for details.
func (x gen_NSMenu) RemoveItem(
	item NSMenuItemRef,
) {
	C.NSMenu_inst_RemoveItem(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
	)

	return
}

// RemoveItemAtIndex removes the menu item at a specified location in the menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518207-removeitematindex?language=objc for details.
func (x gen_NSMenu) RemoveItemAtIndex(
	index core.NSInteger,
) {
	C.NSMenu_inst_RemoveItemAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)

	return
}

// SetSubmenuForItem assigns a menu to be a submenu of the menu controlled by a given menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518194-setsubmenu?language=objc for details.
func (x gen_NSMenu) SetSubmenuForItem(
	menu NSMenuRef,
	item NSMenuItemRef,
) {
	C.NSMenu_inst_SetSubmenuForItem(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(menu),
		objc.RefPointer(item),
	)

	return
}

// SubmenuAction returns the action method assigned to menu items that open submenus.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518179-submenuaction?language=objc for details.
func (x gen_NSMenu) SubmenuAction(
	sender objc.Ref,
) {
	C.NSMenu_inst_SubmenuAction(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// Update enables or disables the menu items of the menu based on the NSMenuValidation informal protocol and sizes the menu to fit its current menu items if necessary.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518249-update?language=objc for details.
func (x gen_NSMenu) Update() {
	C.NSMenu_inst_Update(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Init is undocumented.
func (x gen_NSMenu) Init() NSMenu {
	ret := C.NSMenu_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenu_FromPointer(ret)
}

// MenuBarHeight returns the menu bar height for the main menu in pixels.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518141-menubarheight?language=objc for details.
func (x gen_NSMenu) MenuBarHeight() core.CGFloat {
	ret := C.NSMenu_inst_MenuBarHeight(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// NumberOfItems returns the number of menu items in the menu, including separator items.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518202-numberofitems?language=objc for details.
func (x gen_NSMenu) NumberOfItems() core.NSInteger {
	ret := C.NSMenu_inst_NumberOfItems(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// ItemArray an array containing the menu items in the menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518186-itemarray?language=objc for details.
func (x gen_NSMenu) ItemArray() core.NSArray {
	ret := C.NSMenu_inst_ItemArray(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// SetItemArray an array containing the menu items in the menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518186-itemarray?language=objc for details.
func (x gen_NSMenu) SetItemArray(
	value core.NSArrayRef,
) {
	C.NSMenu_inst_SetItemArray(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Supermenu returns the parent menu that contains the menu as a submenu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518204-supermenu?language=objc for details.
func (x gen_NSMenu) Supermenu() NSMenu {
	ret := C.NSMenu_inst_Supermenu(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenu_FromPointer(ret)
}

// SetSupermenu returns the parent menu that contains the menu as a submenu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518204-supermenu?language=objc for details.
func (x gen_NSMenu) SetSupermenu(
	value NSMenuRef,
) {
	C.NSMenu_inst_SetSupermenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// AutoenablesItems indicates whether the menu automatically enables and disables its menu items.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518227-autoenablesitems?language=objc for details.
func (x gen_NSMenu) AutoenablesItems() bool {
	ret := C.NSMenu_inst_AutoenablesItems(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAutoenablesItems indicates whether the menu automatically enables and disables its menu items.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518227-autoenablesitems?language=objc for details.
func (x gen_NSMenu) SetAutoenablesItems(
	value bool,
) {
	C.NSMenu_inst_SetAutoenablesItems(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// Font returns the font of the menu and its submenus.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518230-font?language=objc for details.
func (x gen_NSMenu) Font() NSFont {
	ret := C.NSMenu_inst_Font(
		unsafe.Pointer(x.Pointer()),
	)

	return NSFont_FromPointer(ret)
}

// SetFont returns the font of the menu and its submenus.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518230-font?language=objc for details.
func (x gen_NSMenu) SetFont(
	value NSFontRef,
) {
	C.NSMenu_inst_SetFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Title returns the title of the menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518192-title?language=objc for details.
func (x gen_NSMenu) Title() core.NSString {
	ret := C.NSMenu_inst_Title(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetTitle returns the title of the menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518192-title?language=objc for details.
func (x gen_NSMenu) SetTitle(
	value core.NSStringRef,
) {
	C.NSMenu_inst_SetTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// MinimumWidth returns the minimum width of the menu in screen coordinates.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518221-minimumwidth?language=objc for details.
func (x gen_NSMenu) MinimumWidth() core.CGFloat {
	ret := C.NSMenu_inst_MinimumWidth(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// SetMinimumWidth returns the minimum width of the menu in screen coordinates.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518221-minimumwidth?language=objc for details.
func (x gen_NSMenu) SetMinimumWidth(
	value core.CGFloat,
) {
	C.NSMenu_inst_SetMinimumWidth(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// Size returns the size of the menu in screen coordinates
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518185-size?language=objc for details.
func (x gen_NSMenu) Size() core.NSSize {
	ret := C.NSMenu_inst_Size(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// AllowsContextMenuPlugIns indicates whether the pop-up menu allows appending of contextual menu plug-in items.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518220-allowscontextmenuplugins?language=objc for details.
func (x gen_NSMenu) AllowsContextMenuPlugIns() bool {
	ret := C.NSMenu_inst_AllowsContextMenuPlugIns(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsContextMenuPlugIns indicates whether the pop-up menu allows appending of contextual menu plug-in items.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518220-allowscontextmenuplugins?language=objc for details.
func (x gen_NSMenu) SetAllowsContextMenuPlugIns(
	value bool,
) {
	C.NSMenu_inst_SetAllowsContextMenuPlugIns(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// ShowsStateColumn indicates whether the menu displays the state column.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518253-showsstatecolumn?language=objc for details.
func (x gen_NSMenu) ShowsStateColumn() bool {
	ret := C.NSMenu_inst_ShowsStateColumn(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetShowsStateColumn indicates whether the menu displays the state column.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518253-showsstatecolumn?language=objc for details.
func (x gen_NSMenu) SetShowsStateColumn(
	value bool,
) {
	C.NSMenu_inst_SetShowsStateColumn(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// HighlightedItem indicates the currently highlighted item in the menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518222-highlighteditem?language=objc for details.
func (x gen_NSMenu) HighlightedItem() NSMenuItem {
	ret := C.NSMenu_inst_HighlightedItem(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenuItem_FromPointer(ret)
}

// Delegate returns the delegate of the menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518169-delegate?language=objc for details.
func (x gen_NSMenu) Delegate() objc.Object {
	ret := C.NSMenu_inst_Delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetDelegate returns the delegate of the menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenu/1518169-delegate?language=objc for details.
func (x gen_NSMenu) SetDelegate(
	value objc.Ref,
) {
	C.NSMenu_inst_SetDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

type NSPopoverRef interface {
	Pointer() uintptr
	Init() NSPopover
}

type gen_NSPopover struct {
	objc.Object
}

func NSPopover_FromPointer(ptr unsafe.Pointer) NSPopover {
	return NSPopover{gen_NSPopover{
		objc.Object_FromPointer(ptr),
	}}
}

func NSPopover_FromRef(ref objc.Ref) NSPopover {
	return NSPopover_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Close forces the popover to close without consulting its delegate.
//
// See https://developer.apple.com/documentation/appkit/nspopover/1526823-close?language=objc for details.
func (x gen_NSPopover) Close() {
	C.NSPopover_inst_Close(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Init is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nspopover/1526851-init?language=objc for details.
func (x gen_NSPopover) Init() NSPopover {
	ret := C.NSPopover_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSPopover_FromPointer(ret)
}

// PerformClose attempts to close the popover.
//
// See https://developer.apple.com/documentation/appkit/nspopover/1534290-performclose?language=objc for details.
func (x gen_NSPopover) PerformClose(
	sender objc.Ref,
) {
	C.NSPopover_inst_PerformClose(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// Behavior specifies the behavior of the popover.
//
// See https://developer.apple.com/documentation/appkit/nspopover/1533539-behavior?language=objc for details.
func (x gen_NSPopover) Behavior() core.NSInteger {
	ret := C.NSPopover_inst_Behavior(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// SetBehavior specifies the behavior of the popover.
//
// See https://developer.apple.com/documentation/appkit/nspopover/1533539-behavior?language=objc for details.
func (x gen_NSPopover) SetBehavior(
	value core.NSInteger,
) {
	C.NSPopover_inst_SetBehavior(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return
}

// PositioningRect returns the rectangle within the positioning view relative to which the popover should be positioned.
//
// See https://developer.apple.com/documentation/appkit/nspopover/1526090-positioningrect?language=objc for details.
func (x gen_NSPopover) PositioningRect() core.NSRect {
	ret := C.NSPopover_inst_PositioningRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// SetPositioningRect returns the rectangle within the positioning view relative to which the popover should be positioned.
//
// See https://developer.apple.com/documentation/appkit/nspopover/1526090-positioningrect?language=objc for details.
func (x gen_NSPopover) SetPositioningRect(
	value core.NSRect,
) {
	C.NSPopover_inst_SetPositioningRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return
}

// Animates specifies if the popover is to be animated.
//
// See https://developer.apple.com/documentation/appkit/nspopover/1526527-animates?language=objc for details.
func (x gen_NSPopover) Animates() bool {
	ret := C.NSPopover_inst_Animates(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAnimates specifies if the popover is to be animated.
//
// See https://developer.apple.com/documentation/appkit/nspopover/1526527-animates?language=objc for details.
func (x gen_NSPopover) SetAnimates(
	value bool,
) {
	C.NSPopover_inst_SetAnimates(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// ContentSize returns the content size of the popover.
//
// See https://developer.apple.com/documentation/appkit/nspopover/1524677-contentsize?language=objc for details.
func (x gen_NSPopover) ContentSize() core.NSSize {
	ret := C.NSPopover_inst_ContentSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetContentSize returns the content size of the popover.
//
// See https://developer.apple.com/documentation/appkit/nspopover/1524677-contentsize?language=objc for details.
func (x gen_NSPopover) SetContentSize(
	value core.NSSize,
) {
	C.NSPopover_inst_SetContentSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// IsShown returns the display state of the popover.
//
// See https://developer.apple.com/documentation/appkit/nspopover/1535120-shown?language=objc for details.
func (x gen_NSPopover) IsShown() bool {
	ret := C.NSPopover_inst_IsShown(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsDetached returns a Boolean value that indicates whether the window created by a popover's detachment is automatically created.
//
// See https://developer.apple.com/documentation/appkit/nspopover/1534278-detached?language=objc for details.
func (x gen_NSPopover) IsDetached() bool {
	ret := C.NSPopover_inst_IsDetached(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

type NSMenuItemRef interface {
	Pointer() uintptr
	Init() NSMenuItem
}

type gen_NSMenuItem struct {
	objc.Object
}

func NSMenuItem_FromPointer(ptr unsafe.Pointer) NSMenuItem {
	return NSMenuItem{gen_NSMenuItem{
		objc.Object_FromPointer(ptr),
	}}
}

func NSMenuItem_FromRef(ref objc.Ref) NSMenuItem {
	return NSMenuItem_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// InitWithTitleActionKeyEquivalent returns an initialized instance of NSMenuItem.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514858-initwithtitle?language=objc for details.
func (x gen_NSMenuItem) InitWithTitleActionKeyEquivalent(
	string core.NSStringRef,
	selector objc.Selector,
	charCode core.NSStringRef,
) NSMenuItem {
	ret := C.NSMenuItem_inst_InitWithTitleActionKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		selector.SelectorAddress(),
		objc.RefPointer(charCode),
	)

	return NSMenuItem_FromPointer(ret)
}

// Init is undocumented.
func (x gen_NSMenuItem) Init() NSMenuItem {
	ret := C.NSMenuItem_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenuItem_FromPointer(ret)
}

// IsEnabled returns a Boolean value that indicates whether the menu item is enabled.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514863-enabled?language=objc for details.
func (x gen_NSMenuItem) IsEnabled() bool {
	ret := C.NSMenuItem_inst_IsEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetEnabled returns a Boolean value that indicates whether the menu item is enabled.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514863-enabled?language=objc for details.
func (x gen_NSMenuItem) SetEnabled(
	value bool,
) {
	C.NSMenuItem_inst_SetEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsHidden returns a Boolean value that indicates whether the menu item is hidden.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514846-hidden?language=objc for details.
func (x gen_NSMenuItem) IsHidden() bool {
	ret := C.NSMenuItem_inst_IsHidden(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetHidden returns a Boolean value that indicates whether the menu item is hidden.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514846-hidden?language=objc for details.
func (x gen_NSMenuItem) SetHidden(
	value bool,
) {
	C.NSMenuItem_inst_SetHidden(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsHiddenOrHasHiddenAncestor returns a Boolean value that indicates whether the menu item or any of its superitems is hidden.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514832-hiddenorhashiddenancestor?language=objc for details.
func (x gen_NSMenuItem) IsHiddenOrHasHiddenAncestor() bool {
	ret := C.NSMenuItem_inst_IsHiddenOrHasHiddenAncestor(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Target returns the menu item's target.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514843-target?language=objc for details.
func (x gen_NSMenuItem) Target() objc.Object {
	ret := C.NSMenuItem_inst_Target(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetTarget returns the menu item's target.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514843-target?language=objc for details.
func (x gen_NSMenuItem) SetTarget(
	value objc.Ref,
) {
	C.NSMenuItem_inst_SetTarget(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Action returns the menu item's action-method selector.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514825-action?language=objc for details.
func (x gen_NSMenuItem) Action() objc.Selector {
	ret := C.NSMenuItem_inst_Action(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.SelectorAt(ret)
}

// SetAction returns the menu item's action-method selector.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514825-action?language=objc for details.
func (x gen_NSMenuItem) SetAction(
	value objc.Selector,
) {
	C.NSMenuItem_inst_SetAction(
		unsafe.Pointer(x.Pointer()),
		value.SelectorAddress(),
	)

	return
}

// Title returns the menu item's title.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514805-title?language=objc for details.
func (x gen_NSMenuItem) Title() core.NSString {
	ret := C.NSMenuItem_inst_Title(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetTitle returns the menu item's title.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514805-title?language=objc for details.
func (x gen_NSMenuItem) SetTitle(
	value core.NSStringRef,
) {
	C.NSMenuItem_inst_SetTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// AttributedTitle returns a custom string for a menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514860-attributedtitle?language=objc for details.
func (x gen_NSMenuItem) AttributedTitle() core.NSAttributedString {
	ret := C.NSMenuItem_inst_AttributedTitle(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSAttributedString_FromPointer(ret)
}

// SetAttributedTitle returns a custom string for a menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514860-attributedtitle?language=objc for details.
func (x gen_NSMenuItem) SetAttributedTitle(
	value core.NSAttributedStringRef,
) {
	C.NSMenuItem_inst_SetAttributedTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Tag returns the menu item's tag.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514840-tag?language=objc for details.
func (x gen_NSMenuItem) Tag() core.NSInteger {
	ret := C.NSMenuItem_inst_Tag(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// SetTag returns the menu item's tag.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514840-tag?language=objc for details.
func (x gen_NSMenuItem) SetTag(
	value core.NSInteger,
) {
	C.NSMenuItem_inst_SetTag(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return
}

// State returns the state of the menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514804-state?language=objc for details.
func (x gen_NSMenuItem) State() core.NSInteger {
	ret := C.NSMenuItem_inst_State(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// SetState returns the state of the menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514804-state?language=objc for details.
func (x gen_NSMenuItem) SetState(
	value core.NSInteger,
) {
	C.NSMenuItem_inst_SetState(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return
}

// Image returns the menu item’s image.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514819-image?language=objc for details.
func (x gen_NSMenuItem) Image() NSImage {
	ret := C.NSMenuItem_inst_Image(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_FromPointer(ret)
}

// SetImage returns the menu item’s image.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514819-image?language=objc for details.
func (x gen_NSMenuItem) SetImage(
	value NSImageRef,
) {
	C.NSMenuItem_inst_SetImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// OnStateImage returns the image of the menu item that indicates an “on” state.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514861-onstateimage?language=objc for details.
func (x gen_NSMenuItem) OnStateImage() NSImage {
	ret := C.NSMenuItem_inst_OnStateImage(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_FromPointer(ret)
}

// SetOnStateImage returns the image of the menu item that indicates an “on” state.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514861-onstateimage?language=objc for details.
func (x gen_NSMenuItem) SetOnStateImage(
	value NSImageRef,
) {
	C.NSMenuItem_inst_SetOnStateImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// OffStateImage returns the image of the menu item that indicates an “off” state.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514821-offstateimage?language=objc for details.
func (x gen_NSMenuItem) OffStateImage() NSImage {
	ret := C.NSMenuItem_inst_OffStateImage(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_FromPointer(ret)
}

// SetOffStateImage returns the image of the menu item that indicates an “off” state.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514821-offstateimage?language=objc for details.
func (x gen_NSMenuItem) SetOffStateImage(
	value NSImageRef,
) {
	C.NSMenuItem_inst_SetOffStateImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// MixedStateImage returns the image of the menu item that indicates a “mixed” state, that is, a state neither “on” nor “off.”
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514827-mixedstateimage?language=objc for details.
func (x gen_NSMenuItem) MixedStateImage() NSImage {
	ret := C.NSMenuItem_inst_MixedStateImage(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_FromPointer(ret)
}

// SetMixedStateImage returns the image of the menu item that indicates a “mixed” state, that is, a state neither “on” nor “off.”
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514827-mixedstateimage?language=objc for details.
func (x gen_NSMenuItem) SetMixedStateImage(
	value NSImageRef,
) {
	C.NSMenuItem_inst_SetMixedStateImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Submenu returns the submenu of the menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514845-submenu?language=objc for details.
func (x gen_NSMenuItem) Submenu() NSMenu {
	ret := C.NSMenuItem_inst_Submenu(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenu_FromPointer(ret)
}

// SetSubmenu returns the submenu of the menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514845-submenu?language=objc for details.
func (x gen_NSMenuItem) SetSubmenu(
	value NSMenuRef,
) {
	C.NSMenuItem_inst_SetSubmenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// HasSubmenu returns a Boolean value that indicates whether the menu item has a submenu.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514817-hassubmenu?language=objc for details.
func (x gen_NSMenuItem) HasSubmenu() bool {
	ret := C.NSMenuItem_inst_HasSubmenu(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// ParentItem returns the menu item whose submenu contains the receiver.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514813-parentitem?language=objc for details.
func (x gen_NSMenuItem) ParentItem() NSMenuItem {
	ret := C.NSMenuItem_inst_ParentItem(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenuItem_FromPointer(ret)
}

// IsSeparatorItem returns a menu item that is used to separate logical groups of menu commands.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514837-separatoritem?language=objc for details.
func (x gen_NSMenuItem) IsSeparatorItem() bool {
	ret := C.NSMenuItem_inst_IsSeparatorItem(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Menu returns the menu item’s menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514830-menu?language=objc for details.
func (x gen_NSMenuItem) Menu() NSMenu {
	ret := C.NSMenuItem_inst_Menu(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenu_FromPointer(ret)
}

// SetMenu returns the menu item’s menu.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514830-menu?language=objc for details.
func (x gen_NSMenuItem) SetMenu(
	value NSMenuRef,
) {
	C.NSMenuItem_inst_SetMenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// KeyEquivalent returns the menu item’s unmodified key equivalent.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514842-keyequivalent?language=objc for details.
func (x gen_NSMenuItem) KeyEquivalent() core.NSString {
	ret := C.NSMenuItem_inst_KeyEquivalent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetKeyEquivalent returns the menu item’s unmodified key equivalent.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514842-keyequivalent?language=objc for details.
func (x gen_NSMenuItem) SetKeyEquivalent(
	value core.NSStringRef,
) {
	C.NSMenuItem_inst_SetKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// UserKeyEquivalent returns the user-assigned key equivalent for the menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514850-userkeyequivalent?language=objc for details.
func (x gen_NSMenuItem) UserKeyEquivalent() core.NSString {
	ret := C.NSMenuItem_inst_UserKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// IsAlternate returns a Boolean value that marks the menu item as an alternate to the previous menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514823-alternate?language=objc for details.
func (x gen_NSMenuItem) IsAlternate() bool {
	ret := C.NSMenuItem_inst_IsAlternate(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAlternate returns a Boolean value that marks the menu item as an alternate to the previous menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514823-alternate?language=objc for details.
func (x gen_NSMenuItem) SetAlternate(
	value bool,
) {
	C.NSMenuItem_inst_SetAlternate(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IndentationLevel returns the menu item indentation level for the menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514809-indentationlevel?language=objc for details.
func (x gen_NSMenuItem) IndentationLevel() core.NSInteger {
	ret := C.NSMenuItem_inst_IndentationLevel(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// SetIndentationLevel returns the menu item indentation level for the menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514809-indentationlevel?language=objc for details.
func (x gen_NSMenuItem) SetIndentationLevel(
	value core.NSInteger,
) {
	C.NSMenuItem_inst_SetIndentationLevel(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return
}

// ToolTip returns a help tag for the menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514848-tooltip?language=objc for details.
func (x gen_NSMenuItem) ToolTip() core.NSString {
	ret := C.NSMenuItem_inst_ToolTip(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetToolTip returns a help tag for the menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514848-tooltip?language=objc for details.
func (x gen_NSMenuItem) SetToolTip(
	value core.NSStringRef,
) {
	C.NSMenuItem_inst_SetToolTip(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// RepresentedObject returns the object represented by the menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514834-representedobject?language=objc for details.
func (x gen_NSMenuItem) RepresentedObject() objc.Object {
	ret := C.NSMenuItem_inst_RepresentedObject(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetRepresentedObject returns the object represented by the menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514834-representedobject?language=objc for details.
func (x gen_NSMenuItem) SetRepresentedObject(
	value objc.Ref,
) {
	C.NSMenuItem_inst_SetRepresentedObject(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// View returns the content view for the menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514835-view?language=objc for details.
func (x gen_NSMenuItem) View() NSView {
	ret := C.NSMenuItem_inst_View(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_FromPointer(ret)
}

// SetView returns the content view for the menu item.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514835-view?language=objc for details.
func (x gen_NSMenuItem) SetView(
	value NSViewRef,
) {
	C.NSMenuItem_inst_SetView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// IsHighlighted returns a Boolean value that indicates whether the menu item should be drawn highlighted.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/1514856-highlighted?language=objc for details.
func (x gen_NSMenuItem) IsHighlighted() bool {
	ret := C.NSMenuItem_inst_IsHighlighted(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AllowsAutomaticKeyEquivalentLocalization is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/3787554-allowsautomatickeyequivalentloca?language=objc for details.
func (x gen_NSMenuItem) AllowsAutomaticKeyEquivalentLocalization() bool {
	ret := C.NSMenuItem_inst_AllowsAutomaticKeyEquivalentLocalization(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsAutomaticKeyEquivalentLocalization is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/3787554-allowsautomatickeyequivalentloca?language=objc for details.
func (x gen_NSMenuItem) SetAllowsAutomaticKeyEquivalentLocalization(
	value bool,
) {
	C.NSMenuItem_inst_SetAllowsAutomaticKeyEquivalentLocalization(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AllowsAutomaticKeyEquivalentMirroring is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/3787555-allowsautomatickeyequivalentmirr?language=objc for details.
func (x gen_NSMenuItem) AllowsAutomaticKeyEquivalentMirroring() bool {
	ret := C.NSMenuItem_inst_AllowsAutomaticKeyEquivalentMirroring(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsAutomaticKeyEquivalentMirroring is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/3787555-allowsautomatickeyequivalentmirr?language=objc for details.
func (x gen_NSMenuItem) SetAllowsAutomaticKeyEquivalentMirroring(
	value bool,
) {
	C.NSMenuItem_inst_SetAllowsAutomaticKeyEquivalentMirroring(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AllowsKeyEquivalentWhenHidden is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/2880316-allowskeyequivalentwhenhidden?language=objc for details.
func (x gen_NSMenuItem) AllowsKeyEquivalentWhenHidden() bool {
	ret := C.NSMenuItem_inst_AllowsKeyEquivalentWhenHidden(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsKeyEquivalentWhenHidden is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsmenuitem/2880316-allowskeyequivalentwhenhidden?language=objc for details.
func (x gen_NSMenuItem) SetAllowsKeyEquivalentWhenHidden(
	value bool,
) {
	C.NSMenuItem_inst_SetAllowsKeyEquivalentWhenHidden(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

type NSRunningApplicationRef interface {
	Pointer() uintptr
	Init() NSRunningApplication
}

type gen_NSRunningApplication struct {
	objc.Object
}

func NSRunningApplication_FromPointer(ptr unsafe.Pointer) NSRunningApplication {
	return NSRunningApplication{gen_NSRunningApplication{
		objc.Object_FromPointer(ptr),
	}}
}

func NSRunningApplication_FromRef(ref objc.Ref) NSRunningApplication {
	return NSRunningApplication_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// ForceTerminate attempts to force the receiver to quit.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1530370-forceterminate?language=objc for details.
func (x gen_NSRunningApplication) ForceTerminate() bool {
	ret := C.NSRunningApplication_inst_ForceTerminate(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Hide attempts to hide or the application.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1526608-hide?language=objc for details.
func (x gen_NSRunningApplication) Hide() bool {
	ret := C.NSRunningApplication_inst_Hide(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Terminate attempts to quit the receiver normally.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1528922-terminate?language=objc for details.
func (x gen_NSRunningApplication) Terminate() bool {
	ret := C.NSRunningApplication_inst_Terminate(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Unhide attempts to unhide or the application.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1534676-unhide?language=objc for details.
func (x gen_NSRunningApplication) Unhide() bool {
	ret := C.NSRunningApplication_inst_Unhide(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Init is undocumented.
func (x gen_NSRunningApplication) Init() NSRunningApplication {
	ret := C.NSRunningApplication_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSRunningApplication_FromPointer(ret)
}

// IsActive indicates whether the application is currently frontmost.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1528778-active?language=objc for details.
func (x gen_NSRunningApplication) IsActive() bool {
	ret := C.NSRunningApplication_inst_IsActive(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// ActivationPolicy indicates the activation policy of the application.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1533103-activationpolicy?language=objc for details.
func (x gen_NSRunningApplication) ActivationPolicy() core.NSInteger {
	ret := C.NSRunningApplication_inst_ActivationPolicy(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// IsHidden indicates whether the application is currently hidden.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1525949-hidden?language=objc for details.
func (x gen_NSRunningApplication) IsHidden() bool {
	ret := C.NSRunningApplication_inst_IsHidden(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// LocalizedName indicates the localized name of the application.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1526751-localizedname?language=objc for details.
func (x gen_NSRunningApplication) LocalizedName() core.NSString {
	ret := C.NSRunningApplication_inst_LocalizedName(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// Icon returns the icon for the receiver’s application.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1529885-icon?language=objc for details.
func (x gen_NSRunningApplication) Icon() NSImage {
	ret := C.NSRunningApplication_inst_Icon(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_FromPointer(ret)
}

// BundleIdentifier indicates the CFBundleIdentifier of the application.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1529140-bundleidentifier?language=objc for details.
func (x gen_NSRunningApplication) BundleIdentifier() core.NSString {
	ret := C.NSRunningApplication_inst_BundleIdentifier(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// BundleURL indicates the URL to the application's bundle.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1535500-bundleurl?language=objc for details.
func (x gen_NSRunningApplication) BundleURL() core.NSURL {
	ret := C.NSRunningApplication_inst_BundleURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_FromPointer(ret)
}

// ExecutableArchitecture indicates the executing processor architecture for the application.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1524287-executablearchitecture?language=objc for details.
func (x gen_NSRunningApplication) ExecutableArchitecture() core.NSInteger {
	ret := C.NSRunningApplication_inst_ExecutableArchitecture(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// ExecutableURL indicates the URL to the application's executable.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1531062-executableurl?language=objc for details.
func (x gen_NSRunningApplication) ExecutableURL() core.NSURL {
	ret := C.NSRunningApplication_inst_ExecutableURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_FromPointer(ret)
}

// IsFinishedLaunching indicates whether the receiver’s process has finished launching,
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1532002-finishedlaunching?language=objc for details.
func (x gen_NSRunningApplication) IsFinishedLaunching() bool {
	ret := C.NSRunningApplication_inst_IsFinishedLaunching(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// OwnsMenuBar returns whether the application owns the current menu bar.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1525915-ownsmenubar?language=objc for details.
func (x gen_NSRunningApplication) OwnsMenuBar() bool {
	ret := C.NSRunningApplication_inst_OwnsMenuBar(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsTerminated indicates that the receiver’s application has terminated.
//
// See https://developer.apple.com/documentation/appkit/nsrunningapplication/1532239-terminated?language=objc for details.
func (x gen_NSRunningApplication) IsTerminated() bool {
	ret := C.NSRunningApplication_inst_IsTerminated(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

type NSScreenRef interface {
	Pointer() uintptr
	Init() NSScreen
}

type gen_NSScreen struct {
	objc.Object
}

func NSScreen_FromPointer(ptr unsafe.Pointer) NSScreen {
	return NSScreen{gen_NSScreen{
		objc.Object_FromPointer(ptr),
	}}
}

func NSScreen_FromRef(ref objc.Ref) NSScreen {
	return NSScreen_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// ConvertRectFromBacking converts the rectangle from the device pixel aligned coordinates system of a screen.
//
// See https://developer.apple.com/documentation/appkit/nsscreen/1388364-convertrectfrombacking?language=objc for details.
func (x gen_NSScreen) ConvertRectFromBacking(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSScreen_inst_ConvertRectFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// ConvertRectToBacking converts the rectangle to the device pixel aligned coordinates system of a screen.
//
// See https://developer.apple.com/documentation/appkit/nsscreen/1388389-convertrecttobacking?language=objc for details.
func (x gen_NSScreen) ConvertRectToBacking(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSScreen_inst_ConvertRectToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// Init is undocumented.
func (x gen_NSScreen) Init() NSScreen {
	ret := C.NSScreen_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSScreen_FromPointer(ret)
}

// Frame returns the dimensions and location of the screen.
//
// See https://developer.apple.com/documentation/appkit/nsscreen/1388387-frame?language=objc for details.
func (x gen_NSScreen) Frame() core.NSRect {
	ret := C.NSScreen_inst_Frame(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// DeviceDescription returns the device dictionary for the screen.
//
// See https://developer.apple.com/documentation/appkit/nsscreen/1388360-devicedescription?language=objc for details.
func (x gen_NSScreen) DeviceDescription() core.NSDictionary {
	ret := C.NSScreen_inst_DeviceDescription(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_FromPointer(ret)
}

// VisibleFrame returns the current location and dimensions of the visible screen.
//
// See https://developer.apple.com/documentation/appkit/nsscreen/1388369-visibleframe?language=objc for details.
func (x gen_NSScreen) VisibleFrame() core.NSRect {
	ret := C.NSScreen_inst_VisibleFrame(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// BackingScaleFactor returns the backing store pixel scale factor for the screen.
//
// See https://developer.apple.com/documentation/appkit/nsscreen/1388385-backingscalefactor?language=objc for details.
func (x gen_NSScreen) BackingScaleFactor() core.CGFloat {
	ret := C.NSScreen_inst_BackingScaleFactor(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// MaximumPotentialExtendedDynamicRangeColorComponentValue returns the maximum possible color component value for the screen when it's in extended dynamic range (EDR) mode.
//
// See https://developer.apple.com/documentation/appkit/nsscreen/3180381-maximumpotentialextendeddynamicr?language=objc for details.
func (x gen_NSScreen) MaximumPotentialExtendedDynamicRangeColorComponentValue() core.CGFloat {
	ret := C.NSScreen_inst_MaximumPotentialExtendedDynamicRangeColorComponentValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// MaximumExtendedDynamicRangeColorComponentValue returns the current maximum color component value for the screen.
//
// See https://developer.apple.com/documentation/appkit/nsscreen/1388362-maximumextendeddynamicrangecolor?language=objc for details.
func (x gen_NSScreen) MaximumExtendedDynamicRangeColorComponentValue() core.CGFloat {
	ret := C.NSScreen_inst_MaximumExtendedDynamicRangeColorComponentValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// MaximumReferenceExtendedDynamicRangeColorComponentValue returns the current maximum color component value for reference rendering to the screen.
//
// See https://developer.apple.com/documentation/appkit/nsscreen/3180382-maximumreferenceextendeddynamicr?language=objc for details.
func (x gen_NSScreen) MaximumReferenceExtendedDynamicRangeColorComponentValue() core.CGFloat {
	ret := C.NSScreen_inst_MaximumReferenceExtendedDynamicRangeColorComponentValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// LocalizedName is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsscreen/3228043-localizedname?language=objc for details.
func (x gen_NSScreen) LocalizedName() core.NSString {
	ret := C.NSScreen_inst_LocalizedName(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// MaximumFramesPerSecond is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsscreen/3824745-maximumframespersecond?language=objc for details.
func (x gen_NSScreen) MaximumFramesPerSecond() core.NSInteger {
	ret := C.NSScreen_inst_MaximumFramesPerSecond(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

type NSStatusBarRef interface {
	Pointer() uintptr
	Init() NSStatusBar
}

type gen_NSStatusBar struct {
	objc.Object
}

func NSStatusBar_FromPointer(ptr unsafe.Pointer) NSStatusBar {
	return NSStatusBar{gen_NSStatusBar{
		objc.Object_FromPointer(ptr),
	}}
}

func NSStatusBar_FromRef(ref objc.Ref) NSStatusBar {
	return NSStatusBar_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// RemoveStatusItem removes the specified status item from the receiver.
//
// See https://developer.apple.com/documentation/appkit/nsstatusbar/1530377-removestatusitem?language=objc for details.
func (x gen_NSStatusBar) RemoveStatusItem(
	item NSStatusItemRef,
) {
	C.NSStatusBar_inst_RemoveStatusItem(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
	)

	return
}

// StatusItemWithLength returns a newly created status item that has been allotted a specified space within the status bar.
//
// See https://developer.apple.com/documentation/appkit/nsstatusbar/1532895-statusitemwithlength?language=objc for details.
func (x gen_NSStatusBar) StatusItemWithLength(
	length core.CGFloat,
) NSStatusItem {
	ret := C.NSStatusBar_inst_StatusItemWithLength(
		unsafe.Pointer(x.Pointer()),
		C.double(length),
	)

	return NSStatusItem_FromPointer(ret)
}

// Init is undocumented.
func (x gen_NSStatusBar) Init() NSStatusBar {
	ret := C.NSStatusBar_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSStatusBar_FromPointer(ret)
}

// IsVertical returns a Boolean value indicating whether the status bar has a vertical orientation.
//
// See https://developer.apple.com/documentation/appkit/nsstatusbar/1530580-vertical?language=objc for details.
func (x gen_NSStatusBar) IsVertical() bool {
	ret := C.NSStatusBar_inst_IsVertical(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Thickness returns the thickness of the status bar, in pixels.
//
// See https://developer.apple.com/documentation/appkit/nsstatusbar/1534591-thickness?language=objc for details.
func (x gen_NSStatusBar) Thickness() core.CGFloat {
	ret := C.NSStatusBar_inst_Thickness(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

type NSStatusBarButtonRef interface {
	Pointer() uintptr
	Init() NSStatusBarButton
}

type gen_NSStatusBarButton struct {
	NSButton
}

func NSStatusBarButton_FromPointer(ptr unsafe.Pointer) NSStatusBarButton {
	return NSStatusBarButton{gen_NSStatusBarButton{
		NSButton_FromPointer(ptr),
	}}
}

func NSStatusBarButton_FromRef(ref objc.Ref) NSStatusBarButton {
	return NSStatusBarButton_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init is undocumented.
func (x gen_NSStatusBarButton) Init() NSStatusBarButton {
	ret := C.NSStatusBarButton_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSStatusBarButton_FromPointer(ret)
}

// AppearsDisabled is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsstatusbarbutton/1409292-appearsdisabled?language=objc for details.
func (x gen_NSStatusBarButton) AppearsDisabled() bool {
	ret := C.NSStatusBarButton_inst_AppearsDisabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAppearsDisabled is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsstatusbarbutton/1409292-appearsdisabled?language=objc for details.
func (x gen_NSStatusBarButton) SetAppearsDisabled(
	value bool,
) {
	C.NSStatusBarButton_inst_SetAppearsDisabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

type NSStatusItemRef interface {
	Pointer() uintptr
	Init() NSStatusItem
}

type gen_NSStatusItem struct {
	objc.Object
}

func NSStatusItem_FromPointer(ptr unsafe.Pointer) NSStatusItem {
	return NSStatusItem{gen_NSStatusItem{
		objc.Object_FromPointer(ptr),
	}}
}

func NSStatusItem_FromRef(ref objc.Ref) NSStatusItem {
	return NSStatusItem_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init is undocumented.
func (x gen_NSStatusItem) Init() NSStatusItem {
	ret := C.NSStatusItem_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSStatusItem_FromPointer(ret)
}

// StatusBar returns the status bar that displays the status item.
//
// See https://developer.apple.com/documentation/appkit/nsstatusitem/1525951-statusbar?language=objc for details.
func (x gen_NSStatusItem) StatusBar() NSStatusBar {
	ret := C.NSStatusItem_inst_StatusBar(
		unsafe.Pointer(x.Pointer()),
	)

	return NSStatusBar_FromPointer(ret)
}

// Button returns the button displayed in the status bar.
//
// See https://developer.apple.com/documentation/appkit/nsstatusitem/1535056-button?language=objc for details.
func (x gen_NSStatusItem) Button() NSStatusBarButton {
	ret := C.NSStatusItem_inst_Button(
		unsafe.Pointer(x.Pointer()),
	)

	return NSStatusBarButton_FromPointer(ret)
}

// Menu returns the pull-down menu displayed when the user clicks the status item.
//
// See https://developer.apple.com/documentation/appkit/nsstatusitem/1535918-menu?language=objc for details.
func (x gen_NSStatusItem) Menu() NSMenu {
	ret := C.NSStatusItem_inst_Menu(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenu_FromPointer(ret)
}

// SetMenu returns the pull-down menu displayed when the user clicks the status item.
//
// See https://developer.apple.com/documentation/appkit/nsstatusitem/1535918-menu?language=objc for details.
func (x gen_NSStatusItem) SetMenu(
	value NSMenuRef,
) {
	C.NSStatusItem_inst_SetMenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// IsVisible returns a Boolean value indicating if the menu bar currently displays the status item.
//
// See https://developer.apple.com/documentation/appkit/nsstatusitem/1644025-visible?language=objc for details.
func (x gen_NSStatusItem) IsVisible() bool {
	ret := C.NSStatusItem_inst_IsVisible(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetVisible returns a Boolean value indicating if the menu bar currently displays the status item.
//
// See https://developer.apple.com/documentation/appkit/nsstatusitem/1644025-visible?language=objc for details.
func (x gen_NSStatusItem) SetVisible(
	value bool,
) {
	C.NSStatusItem_inst_SetVisible(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// Length returns the amount of space in the status bar that should be allocated to the status item.
//
// See https://developer.apple.com/documentation/appkit/nsstatusitem/1529402-length?language=objc for details.
func (x gen_NSStatusItem) Length() core.CGFloat {
	ret := C.NSStatusItem_inst_Length(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// SetLength returns the amount of space in the status bar that should be allocated to the status item.
//
// See https://developer.apple.com/documentation/appkit/nsstatusitem/1529402-length?language=objc for details.
func (x gen_NSStatusItem) SetLength(
	value core.CGFloat,
) {
	C.NSStatusItem_inst_SetLength(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

type NSTextRef interface {
	Pointer() uintptr
	Init() NSText
}

type gen_NSText struct {
	NSView
}

func NSText_FromPointer(ptr unsafe.Pointer) NSText {
	return NSText{gen_NSText{
		NSView_FromPointer(ptr),
	}}
}

func NSText_FromRef(ref objc.Ref) NSText {
	return NSText_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// AlignCenter this action method applies center alignment to selected paragraphs (or all text if the receiver is a plain text object).
//
// See https://developer.apple.com/documentation/appkit/nstext/1535569-aligncenter?language=objc for details.
func (x gen_NSText) AlignCenter(
	sender objc.Ref,
) {
	C.NSText_inst_AlignCenter(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// AlignLeft this action method applies left alignment to selected paragraphs (or all text if the receiver is a plain text object).
//
// See https://developer.apple.com/documentation/appkit/nstext/1535705-alignleft?language=objc for details.
func (x gen_NSText) AlignLeft(
	sender objc.Ref,
) {
	C.NSText_inst_AlignLeft(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// AlignRight this action method applies right alignment to selected paragraphs (or all text if the receiver is a plain text object).
//
// See https://developer.apple.com/documentation/appkit/nstext/1526477-alignright?language=objc for details.
func (x gen_NSText) AlignRight(
	sender objc.Ref,
) {
	C.NSText_inst_AlignRight(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ChangeFont this action method changes the font of the selection for a rich text object, or of all text for a plain text object.
//
// See https://developer.apple.com/documentation/appkit/nstext/1531459-changefont?language=objc for details.
func (x gen_NSText) ChangeFont(
	sender objc.Ref,
) {
	C.NSText_inst_ChangeFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// CheckSpelling this action method searches for a misspelled word in the receiver’s text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1534926-checkspelling?language=objc for details.
func (x gen_NSText) CheckSpelling(
	sender objc.Ref,
) {
	C.NSText_inst_CheckSpelling(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// Copy this action method copies the selected text onto the general pasteboard, in as many formats as the receiver supports.
//
// See https://developer.apple.com/documentation/appkit/nstext/1525497-copy?language=objc for details.
func (x gen_NSText) Copy(
	sender objc.Ref,
) {
	C.NSText_inst_Copy(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// CopyFont this action method copies the font information for the first character of the selection (or for the insertion point) onto the font pasteboard, as NSFontPboardType.
//
// See https://developer.apple.com/documentation/appkit/nstext/1531255-copyfont?language=objc for details.
func (x gen_NSText) CopyFont(
	sender objc.Ref,
) {
	C.NSText_inst_CopyFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// CopyRuler this action method copies the paragraph style information for first selected paragraph onto the ruler pasteboard, as NSRulerPboardType, and expands the selection to paragraph boundaries.
//
// See https://developer.apple.com/documentation/appkit/nstext/1533303-copyruler?language=objc for details.
func (x gen_NSText) CopyRuler(
	sender objc.Ref,
) {
	C.NSText_inst_CopyRuler(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// Cut this action method deletes the selected text and places it onto the general pasteboard, in as many formats as the receiver supports.
//
// See https://developer.apple.com/documentation/appkit/nstext/1524858-cut?language=objc for details.
func (x gen_NSText) Cut(
	sender objc.Ref,
) {
	C.NSText_inst_Cut(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// Delete this action method deletes the selected text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1524660-delete?language=objc for details.
func (x gen_NSText) Delete(
	sender objc.Ref,
) {
	C.NSText_inst_Delete(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// InitWithFrame is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nstext/1525191-initwithframe?language=objc for details.
func (x gen_NSText) InitWithFrame(
	frameRect core.NSRect,
) NSText {
	ret := C.NSText_inst_InitWithFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)

	return NSText_FromPointer(ret)
}

// Paste this action method pastes text from the general pasteboard at the insertion point or over the selection.
//
// See https://developer.apple.com/documentation/appkit/nstext/1527209-paste?language=objc for details.
func (x gen_NSText) Paste(
	sender objc.Ref,
) {
	C.NSText_inst_Paste(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// PasteFont this action method pastes font information from the font pasteboard onto the selected text or insertion point of a rich text object, or over all text of a plain text object.
//
// See https://developer.apple.com/documentation/appkit/nstext/1531099-pastefont?language=objc for details.
func (x gen_NSText) PasteFont(
	sender objc.Ref,
) {
	C.NSText_inst_PasteFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// PasteRuler this action method pastes paragraph style information from the ruler pasteboard onto the selected paragraphs of a rich text object.
//
// See https://developer.apple.com/documentation/appkit/nstext/1530420-pasteruler?language=objc for details.
func (x gen_NSText) PasteRuler(
	sender objc.Ref,
) {
	C.NSText_inst_PasteRuler(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ReadRTFDFromFile attempts to read the RTFD file at path, returning YES if successful and NO if not.
//
// See https://developer.apple.com/documentation/appkit/nstext/1532564-readrtfdfromfile?language=objc for details.
func (x gen_NSText) ReadRTFDFromFile(
	path core.NSStringRef,
) bool {
	ret := C.NSText_inst_ReadRTFDFromFile(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)

	return convertObjCBoolToGo(ret)
}

// SelectAll this action method selects all of the receiver’s text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1527642-selectall?language=objc for details.
func (x gen_NSText) SelectAll(
	sender objc.Ref,
) {
	C.NSText_inst_SelectAll(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ShowGuessPanel this action method opens the Spelling panel, allowing the user to make a correction during spell checking.
//
// See https://developer.apple.com/documentation/appkit/nstext/1533456-showguesspanel?language=objc for details.
func (x gen_NSText) ShowGuessPanel(
	sender objc.Ref,
) {
	C.NSText_inst_ShowGuessPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// SizeToFit resizes the receiver to fit its text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1533991-sizetofit?language=objc for details.
func (x gen_NSText) SizeToFit() {
	C.NSText_inst_SizeToFit(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Subscript this action method applies a subscript attribute to selected text (or all text if the receiver is a plain text object), lowering its baseline offset by a predefined amount.
//
// See https://developer.apple.com/documentation/appkit/nstext/1535373-subscript?language=objc for details.
func (x gen_NSText) Subscript(
	sender objc.Ref,
) {
	C.NSText_inst_Subscript(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// Superscript this action method applies a superscript attribute to selected text (or all text if the receiver is a plain text object), raising its baseline offset by a predefined amount.
//
// See https://developer.apple.com/documentation/appkit/nstext/1525743-superscript?language=objc for details.
func (x gen_NSText) Superscript(
	sender objc.Ref,
) {
	C.NSText_inst_Superscript(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleRuler this action method shows or hides the ruler, if the receiver is enclosed in a scroll view.
//
// See https://developer.apple.com/documentation/appkit/nstext/1535773-toggleruler?language=objc for details.
func (x gen_NSText) ToggleRuler(
	sender objc.Ref,
) {
	C.NSText_inst_ToggleRuler(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// Underline adds the underline attribute to the selected text attributes if absent; removes the attribute if present.
//
// See https://developer.apple.com/documentation/appkit/nstext/1534203-underline?language=objc for details.
func (x gen_NSText) Underline(
	sender objc.Ref,
) {
	C.NSText_inst_Underline(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// Unscript this action method removes any superscripting or subscripting from selected text (or all text if the receiver is a plain text object).
//
// See https://developer.apple.com/documentation/appkit/nstext/1527542-unscript?language=objc for details.
func (x gen_NSText) Unscript(
	sender objc.Ref,
) {
	C.NSText_inst_Unscript(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// WriteRTFDToFileAtomically writes the receiver’s text as RTF with attachments to a file or directory at path.
//
// See https://developer.apple.com/documentation/appkit/nstext/1527085-writertfdtofile?language=objc for details.
func (x gen_NSText) WriteRTFDToFileAtomically(
	path core.NSStringRef,
	flag bool,
) bool {
	ret := C.NSText_inst_WriteRTFDToFileAtomically(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(flag),
	)

	return convertObjCBoolToGo(ret)
}

// Init is undocumented.
func (x gen_NSText) Init() NSText {
	ret := C.NSText_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSText_FromPointer(ret)
}

// String returns the characters of the receiver’s text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1528601-string?language=objc for details.
func (x gen_NSText) String() core.NSString {
	ret := C.NSText_inst_String(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetString returns the characters of the receiver’s text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1528601-string?language=objc for details.
func (x gen_NSText) SetString(
	value core.NSStringRef,
) {
	C.NSText_inst_SetString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// BackgroundColor returns the receiver’s background color to a given color.
//
// See https://developer.apple.com/documentation/appkit/nstext/1535324-backgroundcolor?language=objc for details.
func (x gen_NSText) BackgroundColor() NSColor {
	ret := C.NSText_inst_BackgroundColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_FromPointer(ret)
}

// SetBackgroundColor returns the receiver’s background color to a given color.
//
// See https://developer.apple.com/documentation/appkit/nstext/1535324-backgroundcolor?language=objc for details.
func (x gen_NSText) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSText_inst_SetBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// DrawsBackground returns a Boolean that controls whether the receiver draws its background.
//
// See https://developer.apple.com/documentation/appkit/nstext/1531772-drawsbackground?language=objc for details.
func (x gen_NSText) DrawsBackground() bool {
	ret := C.NSText_inst_DrawsBackground(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetDrawsBackground returns a Boolean that controls whether the receiver draws its background.
//
// See https://developer.apple.com/documentation/appkit/nstext/1531772-drawsbackground?language=objc for details.
func (x gen_NSText) SetDrawsBackground(
	value bool,
) {
	C.NSText_inst_SetDrawsBackground(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsEditable returns a Boolean that controls whether the receiver allows the user to edit its text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1529876-editable?language=objc for details.
func (x gen_NSText) IsEditable() bool {
	ret := C.NSText_inst_IsEditable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetEditable returns a Boolean that controls whether the receiver allows the user to edit its text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1529876-editable?language=objc for details.
func (x gen_NSText) SetEditable(
	value bool,
) {
	C.NSText_inst_SetEditable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsSelectable returns a Boolean that controls whether the receiver allows the user to select its text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1535368-selectable?language=objc for details.
func (x gen_NSText) IsSelectable() bool {
	ret := C.NSText_inst_IsSelectable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetSelectable returns a Boolean that controls whether the receiver allows the user to select its text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1535368-selectable?language=objc for details.
func (x gen_NSText) SetSelectable(
	value bool,
) {
	C.NSText_inst_SetSelectable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsFieldEditor returns a Boolean that controls whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder.
//
// See https://developer.apple.com/documentation/appkit/nstext/1533080-fieldeditor?language=objc for details.
func (x gen_NSText) IsFieldEditor() bool {
	ret := C.NSText_inst_IsFieldEditor(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetFieldEditor returns a Boolean that controls whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder.
//
// See https://developer.apple.com/documentation/appkit/nstext/1533080-fieldeditor?language=objc for details.
func (x gen_NSText) SetFieldEditor(
	value bool,
) {
	C.NSText_inst_SetFieldEditor(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsRichText returns a Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1531003-richtext?language=objc for details.
func (x gen_NSText) IsRichText() bool {
	ret := C.NSText_inst_IsRichText(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetRichText returns a Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1531003-richtext?language=objc for details.
func (x gen_NSText) SetRichText(
	value bool,
) {
	C.NSText_inst_SetRichText(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// ImportsGraphics returns a Boolean that controls whether the receiver allows the user to import files by dragging.
//
// See https://developer.apple.com/documentation/appkit/nstext/1531887-importsgraphics?language=objc for details.
func (x gen_NSText) ImportsGraphics() bool {
	ret := C.NSText_inst_ImportsGraphics(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetImportsGraphics returns a Boolean that controls whether the receiver allows the user to import files by dragging.
//
// See https://developer.apple.com/documentation/appkit/nstext/1531887-importsgraphics?language=objc for details.
func (x gen_NSText) SetImportsGraphics(
	value bool,
) {
	C.NSText_inst_SetImportsGraphics(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// UsesFontPanel returns a Boolean that controls whether the receiver uses the Font panel and Font menu.
//
// See https://developer.apple.com/documentation/appkit/nstext/1527431-usesfontpanel?language=objc for details.
func (x gen_NSText) UsesFontPanel() bool {
	ret := C.NSText_inst_UsesFontPanel(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetUsesFontPanel returns a Boolean that controls whether the receiver uses the Font panel and Font menu.
//
// See https://developer.apple.com/documentation/appkit/nstext/1527431-usesfontpanel?language=objc for details.
func (x gen_NSText) SetUsesFontPanel(
	value bool,
) {
	C.NSText_inst_SetUsesFontPanel(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsRulerVisible returns a Boolean value that indicates whether the receiver’s enclosing scroll view shows its ruler.
//
// See https://developer.apple.com/documentation/appkit/nstext/1533732-rulervisible?language=objc for details.
func (x gen_NSText) IsRulerVisible() bool {
	ret := C.NSText_inst_IsRulerVisible(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Font returns the font of all the receiver’s text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1534646-font?language=objc for details.
func (x gen_NSText) Font() NSFont {
	ret := C.NSText_inst_Font(
		unsafe.Pointer(x.Pointer()),
	)

	return NSFont_FromPointer(ret)
}

// SetFont returns the font of all the receiver’s text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1534646-font?language=objc for details.
func (x gen_NSText) SetFont(
	value NSFontRef,
) {
	C.NSText_inst_SetFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// TextColor returns the text color of all characters in the receiver.
//
// See https://developer.apple.com/documentation/appkit/nstext/1534875-textcolor?language=objc for details.
func (x gen_NSText) TextColor() NSColor {
	ret := C.NSText_inst_TextColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_FromPointer(ret)
}

// SetTextColor returns the text color of all characters in the receiver.
//
// See https://developer.apple.com/documentation/appkit/nstext/1534875-textcolor?language=objc for details.
func (x gen_NSText) SetTextColor(
	value NSColorRef,
) {
	C.NSText_inst_SetTextColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// MaxSize returns the receiver’s maximum size.
//
// See https://developer.apple.com/documentation/appkit/nstext/1535900-maxsize?language=objc for details.
func (x gen_NSText) MaxSize() core.NSSize {
	ret := C.NSText_inst_MaxSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetMaxSize returns the receiver’s maximum size.
//
// See https://developer.apple.com/documentation/appkit/nstext/1535900-maxsize?language=objc for details.
func (x gen_NSText) SetMaxSize(
	value core.NSSize,
) {
	C.NSText_inst_SetMaxSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// MinSize returns the receiver’s minimum size.
//
// See https://developer.apple.com/documentation/appkit/nstext/1526222-minsize?language=objc for details.
func (x gen_NSText) MinSize() core.NSSize {
	ret := C.NSText_inst_MinSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetMinSize returns the receiver’s minimum size.
//
// See https://developer.apple.com/documentation/appkit/nstext/1526222-minsize?language=objc for details.
func (x gen_NSText) SetMinSize(
	value core.NSSize,
) {
	C.NSText_inst_SetMinSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// IsVerticallyResizable returns a Boolean that controls whether the receiver changes its height to fit the height of its text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1535082-verticallyresizable?language=objc for details.
func (x gen_NSText) IsVerticallyResizable() bool {
	ret := C.NSText_inst_IsVerticallyResizable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetVerticallyResizable returns a Boolean that controls whether the receiver changes its height to fit the height of its text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1535082-verticallyresizable?language=objc for details.
func (x gen_NSText) SetVerticallyResizable(
	value bool,
) {
	C.NSText_inst_SetVerticallyResizable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsHorizontallyResizable returns a Boolean that controls whether the receiver changes its width to fit the width of its text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1527489-horizontallyresizable?language=objc for details.
func (x gen_NSText) IsHorizontallyResizable() bool {
	ret := C.NSText_inst_IsHorizontallyResizable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetHorizontallyResizable returns a Boolean that controls whether the receiver changes its width to fit the width of its text.
//
// See https://developer.apple.com/documentation/appkit/nstext/1527489-horizontallyresizable?language=objc for details.
func (x gen_NSText) SetHorizontallyResizable(
	value bool,
) {
	C.NSText_inst_SetHorizontallyResizable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// Delegate returns the receiver’s delegate.
//
// See https://developer.apple.com/documentation/appkit/nstext/1529480-delegate?language=objc for details.
func (x gen_NSText) Delegate() objc.Object {
	ret := C.NSText_inst_Delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetDelegate returns the receiver’s delegate.
//
// See https://developer.apple.com/documentation/appkit/nstext/1529480-delegate?language=objc for details.
func (x gen_NSText) SetDelegate(
	value objc.Ref,
) {
	C.NSText_inst_SetDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

type NSTextFieldRef interface {
	Pointer() uintptr
	Init() NSTextField
}

type gen_NSTextField struct {
	NSControl
}

func NSTextField_FromPointer(ptr unsafe.Pointer) NSTextField {
	return NSTextField{gen_NSTextField{
		NSControl_FromPointer(ptr),
	}}
}

func NSTextField_FromRef(ref objc.Ref) NSTextField {
	return NSTextField_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// SelectText ends editing in the text field and, if it’s selectable, selects the entire text content.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399430-selecttext?language=objc for details.
func (x gen_NSTextField) SelectText(
	sender objc.Ref,
) {
	C.NSTextField_inst_SelectText(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// TextShouldBeginEditing requests permission to begin editing a text object.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399399-textshouldbeginediting?language=objc for details.
func (x gen_NSTextField) TextShouldBeginEditing(
	textObject NSTextRef,
) bool {
	ret := C.NSTextField_inst_TextShouldBeginEditing(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(textObject),
	)

	return convertObjCBoolToGo(ret)
}

// TextShouldEndEditing performs validation on the text field’s new value.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399434-textshouldendediting?language=objc for details.
func (x gen_NSTextField) TextShouldEndEditing(
	textObject NSTextRef,
) bool {
	ret := C.NSTextField_inst_TextShouldEndEditing(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(textObject),
	)

	return convertObjCBoolToGo(ret)
}

// Init is undocumented.
func (x gen_NSTextField) Init() NSTextField {
	ret := C.NSTextField_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextField_FromPointer(ret)
}

// IsSelectable returns a Boolean value that determines whether the user can select the content of the text field.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399422-selectable?language=objc for details.
func (x gen_NSTextField) IsSelectable() bool {
	ret := C.NSTextField_inst_IsSelectable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetSelectable returns a Boolean value that determines whether the user can select the content of the text field.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399422-selectable?language=objc for details.
func (x gen_NSTextField) SetSelectable(
	value bool,
) {
	C.NSTextField_inst_SetSelectable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsEditable returns a Boolean value that controls whether the user can edit the value in the text field.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399407-editable?language=objc for details.
func (x gen_NSTextField) IsEditable() bool {
	ret := C.NSTextField_inst_IsEditable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetEditable returns a Boolean value that controls whether the user can edit the value in the text field.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399407-editable?language=objc for details.
func (x gen_NSTextField) SetEditable(
	value bool,
) {
	C.NSTextField_inst_SetEditable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AllowsEditingTextAttributes returns a Boolean value that controls whether the user can change font attributes of the text field’s string.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399401-allowseditingtextattributes?language=objc for details.
func (x gen_NSTextField) AllowsEditingTextAttributes() bool {
	ret := C.NSTextField_inst_AllowsEditingTextAttributes(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsEditingTextAttributes returns a Boolean value that controls whether the user can change font attributes of the text field’s string.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399401-allowseditingtextattributes?language=objc for details.
func (x gen_NSTextField) SetAllowsEditingTextAttributes(
	value bool,
) {
	C.NSTextField_inst_SetAllowsEditingTextAttributes(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// ImportsGraphics returns a Boolean value that controls whether the user can drag image files into the text field.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399428-importsgraphics?language=objc for details.
func (x gen_NSTextField) ImportsGraphics() bool {
	ret := C.NSTextField_inst_ImportsGraphics(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetImportsGraphics returns a Boolean value that controls whether the user can drag image files into the text field.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399428-importsgraphics?language=objc for details.
func (x gen_NSTextField) SetImportsGraphics(
	value bool,
) {
	C.NSTextField_inst_SetImportsGraphics(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// PlaceholderString returns the string the text field displays when empty to help the user understand the text field’s purpose.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399391-placeholderstring?language=objc for details.
func (x gen_NSTextField) PlaceholderString() core.NSString {
	ret := C.NSTextField_inst_PlaceholderString(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetPlaceholderString returns the string the text field displays when empty to help the user understand the text field’s purpose.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399391-placeholderstring?language=objc for details.
func (x gen_NSTextField) SetPlaceholderString(
	value core.NSStringRef,
) {
	C.NSTextField_inst_SetPlaceholderString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// PlaceholderAttributedString returns the attributed string the text field displays when empty to help the user understand the text field’s purpose.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399387-placeholderattributedstring?language=objc for details.
func (x gen_NSTextField) PlaceholderAttributedString() core.NSAttributedString {
	ret := C.NSTextField_inst_PlaceholderAttributedString(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSAttributedString_FromPointer(ret)
}

// SetPlaceholderAttributedString returns the attributed string the text field displays when empty to help the user understand the text field’s purpose.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399387-placeholderattributedstring?language=objc for details.
func (x gen_NSTextField) SetPlaceholderAttributedString(
	value core.NSAttributedStringRef,
) {
	C.NSTextField_inst_SetPlaceholderAttributedString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// AllowsDefaultTighteningForTruncation returns a Boolean value that controls whether single-line text fields tighten intercharacter spacing before truncating the text.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399405-allowsdefaulttighteningfortrunca?language=objc for details.
func (x gen_NSTextField) AllowsDefaultTighteningForTruncation() bool {
	ret := C.NSTextField_inst_AllowsDefaultTighteningForTruncation(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsDefaultTighteningForTruncation returns a Boolean value that controls whether single-line text fields tighten intercharacter spacing before truncating the text.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399405-allowsdefaulttighteningfortrunca?language=objc for details.
func (x gen_NSTextField) SetAllowsDefaultTighteningForTruncation(
	value bool,
) {
	C.NSTextField_inst_SetAllowsDefaultTighteningForTruncation(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// MaximumNumberOfLines returns the maximum number of lines a wrapping text field displays before clipping or truncating the text.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399424-maximumnumberoflines?language=objc for details.
func (x gen_NSTextField) MaximumNumberOfLines() core.NSInteger {
	ret := C.NSTextField_inst_MaximumNumberOfLines(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// SetMaximumNumberOfLines returns the maximum number of lines a wrapping text field displays before clipping or truncating the text.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399424-maximumnumberoflines?language=objc for details.
func (x gen_NSTextField) SetMaximumNumberOfLines(
	value core.NSInteger,
) {
	C.NSTextField_inst_SetMaximumNumberOfLines(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return
}

// PreferredMaxLayoutWidth returns the maximum width of the text field’s intrinsic content size.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399395-preferredmaxlayoutwidth?language=objc for details.
func (x gen_NSTextField) PreferredMaxLayoutWidth() core.CGFloat {
	ret := C.NSTextField_inst_PreferredMaxLayoutWidth(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// SetPreferredMaxLayoutWidth returns the maximum width of the text field’s intrinsic content size.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399395-preferredmaxlayoutwidth?language=objc for details.
func (x gen_NSTextField) SetPreferredMaxLayoutWidth(
	value core.CGFloat,
) {
	C.NSTextField_inst_SetPreferredMaxLayoutWidth(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// TextColor returns the color of the text field’s content.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399409-textcolor?language=objc for details.
func (x gen_NSTextField) TextColor() NSColor {
	ret := C.NSTextField_inst_TextColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_FromPointer(ret)
}

// SetTextColor returns the color of the text field’s content.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399409-textcolor?language=objc for details.
func (x gen_NSTextField) SetTextColor(
	value NSColorRef,
) {
	C.NSTextField_inst_SetTextColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// BackgroundColor returns the color of the background the text field’s cell draws behind the text.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399389-backgroundcolor?language=objc for details.
func (x gen_NSTextField) BackgroundColor() NSColor {
	ret := C.NSTextField_inst_BackgroundColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_FromPointer(ret)
}

// SetBackgroundColor returns the color of the background the text field’s cell draws behind the text.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399389-backgroundcolor?language=objc for details.
func (x gen_NSTextField) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSTextField_inst_SetBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// DrawsBackground returns a Boolean value that controls whether the text field’s cell draws a background color behind the text.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399416-drawsbackground?language=objc for details.
func (x gen_NSTextField) DrawsBackground() bool {
	ret := C.NSTextField_inst_DrawsBackground(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetDrawsBackground returns a Boolean value that controls whether the text field’s cell draws a background color behind the text.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399416-drawsbackground?language=objc for details.
func (x gen_NSTextField) SetDrawsBackground(
	value bool,
) {
	C.NSTextField_inst_SetDrawsBackground(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsBezeled returns a Boolean value that controls whether the text field draws a bezeled background around its contents.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399435-bezeled?language=objc for details.
func (x gen_NSTextField) IsBezeled() bool {
	ret := C.NSTextField_inst_IsBezeled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetBezeled returns a Boolean value that controls whether the text field draws a bezeled background around its contents.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399435-bezeled?language=objc for details.
func (x gen_NSTextField) SetBezeled(
	value bool,
) {
	C.NSTextField_inst_SetBezeled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsBordered returns a Boolean value that controls whether the text field draws a solid black border around its contents.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399403-bordered?language=objc for details.
func (x gen_NSTextField) IsBordered() bool {
	ret := C.NSTextField_inst_IsBordered(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetBordered returns a Boolean value that controls whether the text field draws a solid black border around its contents.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399403-bordered?language=objc for details.
func (x gen_NSTextField) SetBordered(
	value bool,
) {
	C.NSTextField_inst_SetBordered(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AcceptsFirstResponder returns a Boolean value that indicates whether the text field is editable and accepts first responder status.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399393-acceptsfirstresponder?language=objc for details.
func (x gen_NSTextField) AcceptsFirstResponder() bool {
	ret := C.NSTextField_inst_AcceptsFirstResponder(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AllowsCharacterPickerTouchBarItem returns a Boolean value that controls whether the Touch Bar displays the character picker item for rich text fields.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/2539553-allowscharacterpickertouchbarite?language=objc for details.
func (x gen_NSTextField) AllowsCharacterPickerTouchBarItem() bool {
	ret := C.NSTextField_inst_AllowsCharacterPickerTouchBarItem(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsCharacterPickerTouchBarItem returns a Boolean value that controls whether the Touch Bar displays the character picker item for rich text fields.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/2539553-allowscharacterpickertouchbarite?language=objc for details.
func (x gen_NSTextField) SetAllowsCharacterPickerTouchBarItem(
	value bool,
) {
	C.NSTextField_inst_SetAllowsCharacterPickerTouchBarItem(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsAutomaticTextCompletionEnabled returns a Boolean value that indicates whether the text field automatically completes text as the user types.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/2539554-automatictextcompletionenabled?language=objc for details.
func (x gen_NSTextField) IsAutomaticTextCompletionEnabled() bool {
	ret := C.NSTextField_inst_IsAutomaticTextCompletionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAutomaticTextCompletionEnabled returns a Boolean value that indicates whether the text field automatically completes text as the user types.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/2539554-automatictextcompletionenabled?language=objc for details.
func (x gen_NSTextField) SetAutomaticTextCompletionEnabled(
	value bool,
) {
	C.NSTextField_inst_SetAutomaticTextCompletionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// Delegate returns the text field’s delegate.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399437-delegate?language=objc for details.
func (x gen_NSTextField) Delegate() objc.Object {
	ret := C.NSTextField_inst_Delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetDelegate returns the text field’s delegate.
//
// See https://developer.apple.com/documentation/appkit/nstextfield/1399437-delegate?language=objc for details.
func (x gen_NSTextField) SetDelegate(
	value objc.Ref,
) {
	C.NSTextField_inst_SetDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

type NSTextContainerRef interface {
	Pointer() uintptr
	Init() NSTextContainer
}

type gen_NSTextContainer struct {
	objc.Object
}

func NSTextContainer_FromPointer(ptr unsafe.Pointer) NSTextContainer {
	return NSTextContainer{gen_NSTextContainer{
		objc.Object_FromPointer(ptr),
	}}
}

func NSTextContainer_FromRef(ref objc.Ref) NSTextContainer {
	return NSTextContainer_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// InitWithSize initializes a text container with a specified bounding rectangle.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444529-initwithsize?language=objc for details.
func (x gen_NSTextContainer) InitWithSize(
	size core.NSSize,
) NSTextContainer {
	ret := C.NSTextContainer_inst_InitWithSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return NSTextContainer_FromPointer(ret)
}

// ReplaceLayoutManager replaces the layout manager for the group of text system objects that contains the text container.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444545-replacelayoutmanager?language=objc for details.
func (x gen_NSTextContainer) ReplaceLayoutManager(
	newLayoutManager NSLayoutManagerRef,
) {
	C.NSTextContainer_inst_ReplaceLayoutManager(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newLayoutManager),
	)

	return
}

// Init is undocumented.
func (x gen_NSTextContainer) Init() NSTextContainer {
	ret := C.NSTextContainer_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextContainer_FromPointer(ret)
}

// LayoutManager returns the text container’s layout manager.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444517-layoutmanager?language=objc for details.
func (x gen_NSTextContainer) LayoutManager() NSLayoutManager {
	ret := C.NSTextContainer_inst_LayoutManager(
		unsafe.Pointer(x.Pointer()),
	)

	return NSLayoutManager_FromPointer(ret)
}

// SetLayoutManager returns the text container’s layout manager.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444517-layoutmanager?language=objc for details.
func (x gen_NSTextContainer) SetLayoutManager(
	value NSLayoutManagerRef,
) {
	C.NSTextContainer_inst_SetLayoutManager(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// TextView returns the text container’s text view.
//
// See https://developer.apple.com/documentation/appkit/nstextcontainer/1444537-textview?language=objc for details.
func (x gen_NSTextContainer) TextView() NSTextView {
	ret := C.NSTextContainer_inst_TextView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextView_FromPointer(ret)
}

// SetTextView returns the text container’s text view.
//
// See https://developer.apple.com/documentation/appkit/nstextcontainer/1444537-textview?language=objc for details.
func (x gen_NSTextContainer) SetTextView(
	value NSTextViewRef,
) {
	C.NSTextContainer_inst_SetTextView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Size returns the size of the text container’s bounding rectangle.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444553-size?language=objc for details.
func (x gen_NSTextContainer) Size() core.NSSize {
	ret := C.NSTextContainer_inst_Size(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetSize returns the size of the text container’s bounding rectangle.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444553-size?language=objc for details.
func (x gen_NSTextContainer) SetSize(
	value core.NSSize,
) {
	C.NSTextContainer_inst_SetSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// ExclusionPaths an array of path objects that represents the regions where text doesn’t display in the text container.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444569-exclusionpaths?language=objc for details.
func (x gen_NSTextContainer) ExclusionPaths() core.NSArray {
	ret := C.NSTextContainer_inst_ExclusionPaths(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// SetExclusionPaths an array of path objects that represents the regions where text doesn’t display in the text container.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444569-exclusionpaths?language=objc for details.
func (x gen_NSTextContainer) SetExclusionPaths(
	value core.NSArrayRef,
) {
	C.NSTextContainer_inst_SetExclusionPaths(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// WidthTracksTextView returns a Boolean that controls whether the text container adjusts the width of its bounding rectangle when its text view resizes.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444563-widthtrackstextview?language=objc for details.
func (x gen_NSTextContainer) WidthTracksTextView() bool {
	ret := C.NSTextContainer_inst_WidthTracksTextView(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetWidthTracksTextView returns a Boolean that controls whether the text container adjusts the width of its bounding rectangle when its text view resizes.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444563-widthtrackstextview?language=objc for details.
func (x gen_NSTextContainer) SetWidthTracksTextView(
	value bool,
) {
	C.NSTextContainer_inst_SetWidthTracksTextView(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// HeightTracksTextView returns a Boolean that controls whether the text container adjusts the height of its bounding rectangle when its text view resizes.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444559-heighttrackstextview?language=objc for details.
func (x gen_NSTextContainer) HeightTracksTextView() bool {
	ret := C.NSTextContainer_inst_HeightTracksTextView(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetHeightTracksTextView returns a Boolean that controls whether the text container adjusts the height of its bounding rectangle when its text view resizes.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444559-heighttrackstextview?language=objc for details.
func (x gen_NSTextContainer) SetHeightTracksTextView(
	value bool,
) {
	C.NSTextContainer_inst_SetHeightTracksTextView(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// MaximumNumberOfLines returns the maximum number of lines that the text container can store.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444531-maximumnumberoflines?language=objc for details.
func (x gen_NSTextContainer) MaximumNumberOfLines() core.NSUInteger {
	ret := C.NSTextContainer_inst_MaximumNumberOfLines(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// SetMaximumNumberOfLines returns the maximum number of lines that the text container can store.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444531-maximumnumberoflines?language=objc for details.
func (x gen_NSTextContainer) SetMaximumNumberOfLines(
	value core.NSUInteger,
) {
	C.NSTextContainer_inst_SetMaximumNumberOfLines(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)

	return
}

// LineFragmentPadding returns the value for the text inset within line fragment rectangles.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444527-linefragmentpadding?language=objc for details.
func (x gen_NSTextContainer) LineFragmentPadding() core.CGFloat {
	ret := C.NSTextContainer_inst_LineFragmentPadding(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// SetLineFragmentPadding returns the value for the text inset within line fragment rectangles.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444527-linefragmentpadding?language=objc for details.
func (x gen_NSTextContainer) SetLineFragmentPadding(
	value core.CGFloat,
) {
	C.NSTextContainer_inst_SetLineFragmentPadding(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// IsSimpleRectangularTextContainer returns a Boolean that indicates whether the text container’s region is a rectangle with no holes or gaps, and whose edges are parallel to the text view's coordinate system axes.
//
// See https://developer.apple.com/documentation/uikit/nstextcontainer/1444525-simplerectangulartextcontainer?language=objc for details.
func (x gen_NSTextContainer) IsSimpleRectangularTextContainer() bool {
	ret := C.NSTextContainer_inst_IsSimpleRectangularTextContainer(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

type NSViewControllerRef interface {
	Pointer() uintptr
	Init() NSViewController
}

type gen_NSViewController struct {
	objc.Object
}

func NSViewController_FromPointer(ptr unsafe.Pointer) NSViewController {
	return NSViewController{gen_NSViewController{
		objc.Object_FromPointer(ptr),
	}}
}

func NSViewController_FromRef(ref objc.Ref) NSViewController {
	return NSViewController_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// AddChildViewController returns a convenience method for adding a child view controller at the end of the childViewControllers array.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434501-addchildviewcontroller?language=objc for details.
func (x gen_NSViewController) AddChildViewController(
	childViewController NSViewControllerRef,
) {
	C.NSViewController_inst_AddChildViewController(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(childViewController),
	)

	return
}

// CommitEditing returns whether the receiver was able to commit any pending edits.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434485-commitediting?language=objc for details.
func (x gen_NSViewController) CommitEditing() bool {
	ret := C.NSViewController_inst_CommitEditing(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// CommitEditingWithDelegateDidCommitSelectorContextInfo attempt to commit any currently edited results of the receiver.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434464-commiteditingwithdelegate?language=objc for details.
func (x gen_NSViewController) CommitEditingWithDelegateDidCommitSelectorContextInfo(
	delegate objc.Ref,
	didCommitSelector objc.Selector,
	contextInfo unsafe.Pointer,
) {
	C.NSViewController_inst_CommitEditingWithDelegateDidCommitSelectorContextInfo(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(delegate),
		didCommitSelector.SelectorAddress(),
		contextInfo,
	)

	return
}

// DiscardEditing causes the receiver to discard any changes, restoring the previous values.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434487-discardediting?language=objc for details.
func (x gen_NSViewController) DiscardEditing() {
	C.NSViewController_inst_DiscardEditing(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DismissController is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434447-dismisscontroller?language=objc for details.
func (x gen_NSViewController) DismissController(
	sender objc.Ref,
) {
	C.NSViewController_inst_DismissController(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// DismissViewController dismisses a presented view controller, using the same animator that presented it.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434413-dismissviewcontroller?language=objc for details.
func (x gen_NSViewController) DismissViewController(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_DismissViewController(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)

	return
}

// InsertChildViewControllerAtIndex inserts a specified child view controller into the childViewControllers array at a specified position.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434437-insertchildviewcontroller?language=objc for details.
func (x gen_NSViewController) InsertChildViewControllerAtIndex(
	childViewController NSViewControllerRef,
	index core.NSInteger,
) {
	C.NSViewController_inst_InsertChildViewControllerAtIndex(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(childViewController),
		C.long(index),
	)

	return
}

// LoadView instantiates a view from a nib file and sets the value of the view property.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434405-loadview?language=objc for details.
func (x gen_NSViewController) LoadView() {
	C.NSViewController_inst_LoadView(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// PreferredContentSizeDidChangeForViewController called when there is a change in value of the preferredContentSize property of a child view controller or a presented view controller.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434434-preferredcontentsizedidchangefor?language=objc for details.
func (x gen_NSViewController) PreferredContentSizeDidChangeForViewController(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_PreferredContentSizeDidChangeForViewController(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)

	return
}

// PresentViewControllerAnimator presents another view controller using a specified, custom animator for presentation and dismissal.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434431-presentviewcontroller?language=objc for details.
func (x gen_NSViewController) PresentViewControllerAnimator(
	viewController NSViewControllerRef,
	animator objc.Ref,
) {
	C.NSViewController_inst_PresentViewControllerAnimator(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
		objc.RefPointer(animator),
	)

	return
}

// PresentViewControllerAsModalWindow presents another view controller as a modal window, also known as an alert.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434462-presentviewcontrollerasmodalwind?language=objc for details.
func (x gen_NSViewController) PresentViewControllerAsModalWindow(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_PresentViewControllerAsModalWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)

	return
}

// PresentViewControllerAsSheet presents another view controller as a sheet.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434489-presentviewcontrollerassheet?language=objc for details.
func (x gen_NSViewController) PresentViewControllerAsSheet(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_PresentViewControllerAsSheet(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)

	return
}

// RemoveChildViewControllerAtIndex removes a specified child controller from the view controller.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434404-removechildviewcontrolleratindex?language=objc for details.
func (x gen_NSViewController) RemoveChildViewControllerAtIndex(
	index core.NSInteger,
) {
	C.NSViewController_inst_RemoveChildViewControllerAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)

	return
}

// RemoveFromParentViewController removes the called view controller from its parent view controller.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434466-removefromparentviewcontroller?language=objc for details.
func (x gen_NSViewController) RemoveFromParentViewController() {
	C.NSViewController_inst_RemoveFromParentViewController(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// UpdateViewConstraints called during Auto Layout constraint updating to enable the view controller to mediate the process.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434400-updateviewconstraints?language=objc for details.
func (x gen_NSViewController) UpdateViewConstraints() {
	C.NSViewController_inst_UpdateViewConstraints(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewDidAppear called when the view controller’s view is fully transitioned onto the screen.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434455-viewdidappear?language=objc for details.
func (x gen_NSViewController) ViewDidAppear() {
	C.NSViewController_inst_ViewDidAppear(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewDidDisappear called after the view controller’s view is removed from the view hierarchy in a window.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434416-viewdiddisappear?language=objc for details.
func (x gen_NSViewController) ViewDidDisappear() {
	C.NSViewController_inst_ViewDidDisappear(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewDidLayout called immediately after the layout method of the view controller's view is called.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434451-viewdidlayout?language=objc for details.
func (x gen_NSViewController) ViewDidLayout() {
	C.NSViewController_inst_ViewDidLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewDidLoad called after the view controller’s view has been loaded into memory.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434476-viewdidload?language=objc for details.
func (x gen_NSViewController) ViewDidLoad() {
	C.NSViewController_inst_ViewDidLoad(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewWillAppear called after the view controller’s view has been loaded into memory is about to be added to the view hierarchy in the window.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434415-viewwillappear?language=objc for details.
func (x gen_NSViewController) ViewWillAppear() {
	C.NSViewController_inst_ViewWillAppear(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewWillDisappear called when the view controller’s view is about to be removed from the view hierarchy in the window.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434483-viewwilldisappear?language=objc for details.
func (x gen_NSViewController) ViewWillDisappear() {
	C.NSViewController_inst_ViewWillDisappear(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewWillLayout called just before the layout method of the view controller's view is called.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434495-viewwilllayout?language=objc for details.
func (x gen_NSViewController) ViewWillLayout() {
	C.NSViewController_inst_ViewWillLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewWillTransitionToSize for a view controller that is part of an app extension, called when its view is about to be resized.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434443-viewwilltransitiontosize?language=objc for details.
func (x gen_NSViewController) ViewWillTransitionToSize(
	newSize core.NSSize,
) {
	C.NSViewController_inst_ViewWillTransitionToSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&newSize)),
	)

	return
}

// Init is undocumented.
func (x gen_NSViewController) Init() NSViewController {
	ret := C.NSViewController_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSViewController_FromPointer(ret)
}

// RepresentedObject returns the object whose value is presented in the receiver’s primary view.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434453-representedobject?language=objc for details.
func (x gen_NSViewController) RepresentedObject() objc.Object {
	ret := C.NSViewController_inst_RepresentedObject(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetRepresentedObject returns the object whose value is presented in the receiver’s primary view.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434453-representedobject?language=objc for details.
func (x gen_NSViewController) SetRepresentedObject(
	value objc.Ref,
) {
	C.NSViewController_inst_SetRepresentedObject(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// NibBundle returns the nib bundle to be loaded to instantiate the receiver’s primary view.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434433-nibbundle?language=objc for details.
func (x gen_NSViewController) NibBundle() NSBundle {
	ret := C.NSViewController_inst_NibBundle(
		unsafe.Pointer(x.Pointer()),
	)

	return NSBundle_FromPointer(ret)
}

// View returns the view controller’s primary view.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434401-view?language=objc for details.
func (x gen_NSViewController) View() NSView {
	ret := C.NSViewController_inst_View(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_FromPointer(ret)
}

// SetView returns the view controller’s primary view.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434401-view?language=objc for details.
func (x gen_NSViewController) SetView(
	value NSViewRef,
) {
	C.NSViewController_inst_SetView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Title returns the localized title of the receiver’s primary view.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434426-title?language=objc for details.
func (x gen_NSViewController) Title() core.NSString {
	ret := C.NSViewController_inst_Title(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetTitle returns the localized title of the receiver’s primary view.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434426-title?language=objc for details.
func (x gen_NSViewController) SetTitle(
	value core.NSStringRef,
) {
	C.NSViewController_inst_SetTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// IsViewLoaded returns a Boolean value indicating whether the view controller’s view is loaded into memory.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434435-viewloaded?language=objc for details.
func (x gen_NSViewController) IsViewLoaded() bool {
	ret := C.NSViewController_inst_IsViewLoaded(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// PreferredContentSize returns the desired size of the view controller’s view, in screen units.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434409-preferredcontentsize?language=objc for details.
func (x gen_NSViewController) PreferredContentSize() core.NSSize {
	ret := C.NSViewController_inst_PreferredContentSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetPreferredContentSize returns the desired size of the view controller’s view, in screen units.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434409-preferredcontentsize?language=objc for details.
func (x gen_NSViewController) SetPreferredContentSize(
	value core.NSSize,
) {
	C.NSViewController_inst_SetPreferredContentSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// ChildViewControllers an array of view controllers that are hierarchical children of the view controller.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434432-childviewcontrollers?language=objc for details.
func (x gen_NSViewController) ChildViewControllers() core.NSArray {
	ret := C.NSViewController_inst_ChildViewControllers(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// SetChildViewControllers an array of view controllers that are hierarchical children of the view controller.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434432-childviewcontrollers?language=objc for details.
func (x gen_NSViewController) SetChildViewControllers(
	value core.NSArrayRef,
) {
	C.NSViewController_inst_SetChildViewControllers(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// ParentViewController returns the immediate ancestor view controller of the view controller.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434491-parentviewcontroller?language=objc for details.
func (x gen_NSViewController) ParentViewController() NSViewController {
	ret := C.NSViewController_inst_ParentViewController(
		unsafe.Pointer(x.Pointer()),
	)

	return NSViewController_FromPointer(ret)
}

// PresentedViewControllers returns the view controllers, if any, that are currently presented by the view controller.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434497-presentedviewcontrollers?language=objc for details.
func (x gen_NSViewController) PresentedViewControllers() core.NSArray {
	ret := C.NSViewController_inst_PresentedViewControllers(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// PresentingViewController returns the view controller that presented the view controller or that presented its farthest ancestor view controller.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434439-presentingviewcontroller?language=objc for details.
func (x gen_NSViewController) PresentingViewController() NSViewController {
	ret := C.NSViewController_inst_PresentingViewController(
		unsafe.Pointer(x.Pointer()),
	)

	return NSViewController_FromPointer(ret)
}

// PreferredScreenOrigin for a view controller that is part of an app extension, the preferred screen origin.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434468-preferredscreenorigin?language=objc for details.
func (x gen_NSViewController) PreferredScreenOrigin() core.NSPoint {
	ret := C.NSViewController_inst_PreferredScreenOrigin(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// SetPreferredScreenOrigin for a view controller that is part of an app extension, the preferred screen origin.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434468-preferredscreenorigin?language=objc for details.
func (x gen_NSViewController) SetPreferredScreenOrigin(
	value core.NSPoint,
) {
	C.NSViewController_inst_SetPreferredScreenOrigin(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&value)),
	)

	return
}

// PreferredMaximumSize for a view controller that is part of an app extension, the largest allowable size for the app extension’s primary view, in screen units.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434403-preferredmaximumsize?language=objc for details.
func (x gen_NSViewController) PreferredMaximumSize() core.NSSize {
	ret := C.NSViewController_inst_PreferredMaximumSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// PreferredMinimumSize for a view controller that is part of an app extension, the smallest allowable size for the app extension’s primary view, in screen units.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434418-preferredminimumsize?language=objc for details.
func (x gen_NSViewController) PreferredMinimumSize() core.NSSize {
	ret := C.NSViewController_inst_PreferredMinimumSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SourceItemView is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434479-sourceitemview?language=objc for details.
func (x gen_NSViewController) SourceItemView() NSView {
	ret := C.NSViewController_inst_SourceItemView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_FromPointer(ret)
}

// SetSourceItemView is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsviewcontroller/1434479-sourceitemview?language=objc for details.
func (x gen_NSViewController) SetSourceItemView(
	value NSViewRef,
) {
	C.NSViewController_inst_SetSourceItemView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

type NSVisualEffectViewRef interface {
	Pointer() uintptr
	Init() NSVisualEffectView
}

type gen_NSVisualEffectView struct {
	NSView
}

func NSVisualEffectView_FromPointer(ptr unsafe.Pointer) NSVisualEffectView {
	return NSVisualEffectView{gen_NSVisualEffectView{
		NSView_FromPointer(ptr),
	}}
}

func NSVisualEffectView_FromRef(ref objc.Ref) NSVisualEffectView {
	return NSVisualEffectView_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// ViewDidMoveToWindow notifies the view that it moved to a new window.
//
// See https://developer.apple.com/documentation/appkit/nsvisualeffectview/1534300-viewdidmovetowindow?language=objc for details.
func (x gen_NSVisualEffectView) ViewDidMoveToWindow() {
	C.NSVisualEffectView_inst_ViewDidMoveToWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewWillMoveToWindow notifies the view immediately before it moves to a new window (which may be nil).
//
// See https://developer.apple.com/documentation/appkit/nsvisualeffectview/1534276-viewwillmovetowindow?language=objc for details.
func (x gen_NSVisualEffectView) ViewWillMoveToWindow(
	newWindow NSWindowRef,
) {
	C.NSVisualEffectView_inst_ViewWillMoveToWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newWindow),
	)

	return
}

// Init is undocumented.
func (x gen_NSVisualEffectView) Init() NSVisualEffectView {
	ret := C.NSVisualEffectView_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSVisualEffectView_FromPointer(ret)
}

// IsEmphasized returns a Boolean value indicating whether to emphasize the look of the material.
//
// See https://developer.apple.com/documentation/appkit/nsvisualeffectview/1644721-emphasized?language=objc for details.
func (x gen_NSVisualEffectView) IsEmphasized() bool {
	ret := C.NSVisualEffectView_inst_IsEmphasized(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetEmphasized returns a Boolean value indicating whether to emphasize the look of the material.
//
// See https://developer.apple.com/documentation/appkit/nsvisualeffectview/1644721-emphasized?language=objc for details.
func (x gen_NSVisualEffectView) SetEmphasized(
	value bool,
) {
	C.NSVisualEffectView_inst_SetEmphasized(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// MaskImage an image whose alpha channel masks the visual effect view's material.
//
// See https://developer.apple.com/documentation/appkit/nsvisualeffectview/1535318-maskimage?language=objc for details.
func (x gen_NSVisualEffectView) MaskImage() NSImage {
	ret := C.NSVisualEffectView_inst_MaskImage(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_FromPointer(ret)
}

// SetMaskImage an image whose alpha channel masks the visual effect view's material.
//
// See https://developer.apple.com/documentation/appkit/nsvisualeffectview/1535318-maskimage?language=objc for details.
func (x gen_NSVisualEffectView) SetMaskImage(
	value NSImageRef,
) {
	C.NSVisualEffectView_inst_SetMaskImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

type NSWindowRef interface {
	Pointer() uintptr
	Init() NSWindow
}

type gen_NSWindow struct {
	objc.Object
}

func NSWindow_FromPointer(ptr unsafe.Pointer) NSWindow {
	return NSWindow{gen_NSWindow{
		objc.Object_FromPointer(ptr),
	}}
}

func NSWindow_FromRef(ref objc.Ref) NSWindow {
	return NSWindow_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// AddChildWindowOrdered adds a given window as a child window of the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419152-addchildwindow?language=objc for details.
func (x gen_NSWindow) AddChildWindowOrdered(
	childWin NSWindowRef,
	place core.NSUInteger,
) {
	C.NSWindow_inst_AddChildWindowOrdered(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(childWin),
		C.ulong(place),
	)

	return
}

// AddTabbedWindowOrdered adds the provided window as a new tab in a tabbed window using the specified ordering instruction.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1855947-addtabbedwindow?language=objc for details.
func (x gen_NSWindow) AddTabbedWindowOrdered(
	window NSWindowRef,
	ordered core.NSUInteger,
) {
	C.NSWindow_inst_AddTabbedWindowOrdered(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(window),
		C.ulong(ordered),
	)

	return
}

// BecomeKeyWindow informs the window that it has become the key window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419338-becomekeywindow?language=objc for details.
func (x gen_NSWindow) BecomeKeyWindow() {
	C.NSWindow_inst_BecomeKeyWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// BecomeMainWindow informs the window that it has become the main window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419084-becomemainwindow?language=objc for details.
func (x gen_NSWindow) BecomeMainWindow() {
	C.NSWindow_inst_BecomeMainWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// CascadeTopLeftFromPoint positions the window’s top-left to a given point.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419392-cascadetopleftfrompoint?language=objc for details.
func (x gen_NSWindow) CascadeTopLeftFromPoint(
	topLeftPoint core.NSPoint,
) core.NSPoint {
	ret := C.NSWindow_inst_CascadeTopLeftFromPoint(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&topLeftPoint)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// Center sets the window’s location to the center of the screen.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419090-center?language=objc for details.
func (x gen_NSWindow) Center() {
	C.NSWindow_inst_Center(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Close removes the window from the screen.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419662-close?language=objc for details.
func (x gen_NSWindow) Close() {
	C.NSWindow_inst_Close(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ConstrainFrameRectToScreen modifies and returns a frame rectangle so that its top edge lies on a specific screen.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419779-constrainframerect?language=objc for details.
func (x gen_NSWindow) ConstrainFrameRectToScreen(
	frameRect core.NSRect,
	screen NSScreenRef,
) core.NSRect {
	ret := C.NSWindow_inst_ConstrainFrameRectToScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
		objc.RefPointer(screen),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// ContentRectForFrameRect returns the window’s content rectangle with a given frame rectangle.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419108-contentrectforframerect?language=objc for details.
func (x gen_NSWindow) ContentRectForFrameRect(
	frameRect core.NSRect,
) core.NSRect {
	ret := C.NSWindow_inst_ContentRectForFrameRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// ConvertPointFromBacking converts a point from its pixel-aligned backing store coordinate system to the window’s coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nswindow/2967179-convertpointfrombacking?language=objc for details.
func (x gen_NSWindow) ConvertPointFromBacking(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSWindow_inst_ConvertPointFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// ConvertPointFromScreen converts a point from the screen coordinate system to the window’s coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nswindow/2967180-convertpointfromscreen?language=objc for details.
func (x gen_NSWindow) ConvertPointFromScreen(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSWindow_inst_ConvertPointFromScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// ConvertPointToBacking converts a point from the window’s coordinate system to its pixel-aligned backing store coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nswindow/2967181-convertpointtobacking?language=objc for details.
func (x gen_NSWindow) ConvertPointToBacking(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSWindow_inst_ConvertPointToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// ConvertPointToScreen converts a point to the screen coordinate system from the window’s coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nswindow/2967182-convertpointtoscreen?language=objc for details.
func (x gen_NSWindow) ConvertPointToScreen(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSWindow_inst_ConvertPointToScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// ConvertRectFromBacking converts a rectangle from its pixel-aligned backing store coordinate system to the window’s coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419273-convertrectfrombacking?language=objc for details.
func (x gen_NSWindow) ConvertRectFromBacking(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSWindow_inst_ConvertRectFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// ConvertRectFromScreen converts a rectangle from the screen coordinate system to the window’s coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419603-convertrectfromscreen?language=objc for details.
func (x gen_NSWindow) ConvertRectFromScreen(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSWindow_inst_ConvertRectFromScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// ConvertRectToBacking converts a rectangle from the window’s coordinate system to its pixel-aligned backing store coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419260-convertrecttobacking?language=objc for details.
func (x gen_NSWindow) ConvertRectToBacking(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSWindow_inst_ConvertRectToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// ConvertRectToScreen converts a rectangle to the screen coordinate system from the window’s coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419286-convertrecttoscreen?language=objc for details.
func (x gen_NSWindow) ConvertRectToScreen(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSWindow_inst_ConvertRectToScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// DataWithEPSInsideRect returns EPS data that draws the region of the window within a given rectangle.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419128-datawithepsinsiderect?language=objc for details.
func (x gen_NSWindow) DataWithEPSInsideRect(
	rect core.NSRect,
) core.NSData {
	ret := C.NSWindow_inst_DataWithEPSInsideRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return core.NSData_FromPointer(ret)
}

// DataWithPDFInsideRect returns PDF data that draws the region of the window within a given rectangle.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419418-datawithpdfinsiderect?language=objc for details.
func (x gen_NSWindow) DataWithPDFInsideRect(
	rect core.NSRect,
) core.NSData {
	ret := C.NSWindow_inst_DataWithPDFInsideRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return core.NSData_FromPointer(ret)
}

// Deminiaturize de-minimizes the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419334-deminiaturize?language=objc for details.
func (x gen_NSWindow) Deminiaturize(
	sender objc.Ref,
) {
	C.NSWindow_inst_Deminiaturize(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// DisableCursorRects disables all cursor rectangle management within the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419639-disablecursorrects?language=objc for details.
func (x gen_NSWindow) DisableCursorRects() {
	C.NSWindow_inst_DisableCursorRects(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DisableKeyEquivalentForDefaultButtonCell disables the default button cell’s key equivalent, so it doesn’t perform a click when the user presses Return (or Enter).
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419242-disablekeyequivalentfordefaultbu?language=objc for details.
func (x gen_NSWindow) DisableKeyEquivalentForDefaultButtonCell() {
	C.NSWindow_inst_DisableKeyEquivalentForDefaultButtonCell(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DisableScreenUpdatesUntilFlush disables the window’s screen updates until the window is flushed.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419483-disablescreenupdatesuntilflush?language=objc for details.
func (x gen_NSWindow) DisableScreenUpdatesUntilFlush() {
	C.NSWindow_inst_DisableScreenUpdatesUntilFlush(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DisableSnapshotRestoration disables snapshot restoration.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1526239-disablesnapshotrestoration?language=objc for details.
func (x gen_NSWindow) DisableSnapshotRestoration() {
	C.NSWindow_inst_DisableSnapshotRestoration(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DiscardCursorRects invalidates all cursor rectangles in the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419269-discardcursorrects?language=objc for details.
func (x gen_NSWindow) DiscardCursorRects() {
	C.NSWindow_inst_DiscardCursorRects(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Display passes a display message down the window’s view hierarchy, thus redrawing all views within the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419358-display?language=objc for details.
func (x gen_NSWindow) Display() {
	C.NSWindow_inst_Display(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DisplayIfNeeded passes a display message down the window’s view hierarchy, thus redrawing all views that need displaying.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419096-displayifneeded?language=objc for details.
func (x gen_NSWindow) DisplayIfNeeded() {
	C.NSWindow_inst_DisplayIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DragImageAtOffsetEventPasteboardSourceSlideBack begins a dragging session.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419224-dragimage?language=objc for details.
func (x gen_NSWindow) DragImageAtOffsetEventPasteboardSourceSlideBack(
	image NSImageRef,
	baseLocation core.NSPoint,
	initialOffset core.NSSize,
	event NSEventRef,
	pboard NSPasteboardRef,
	sourceObj objc.Ref,
	slideFlag bool,
) {
	C.NSWindow_inst_DragImageAtOffsetEventPasteboardSourceSlideBack(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(image),
		*(*C.NSPoint)(unsafe.Pointer(&baseLocation)),
		*(*C.NSSize)(unsafe.Pointer(&initialOffset)),
		objc.RefPointer(event),
		objc.RefPointer(pboard),
		objc.RefPointer(sourceObj),
		convertToObjCBool(slideFlag),
	)

	return
}

// EnableCursorRects reenables cursor rectangle management within the window after a disableCursorRects message.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419202-enablecursorrects?language=objc for details.
func (x gen_NSWindow) EnableCursorRects() {
	C.NSWindow_inst_EnableCursorRects(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// EnableKeyEquivalentForDefaultButtonCell reenables the default button cell’s key equivalent, so it performs a click when the user presses Return (or Enter).
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419276-enablekeyequivalentfordefaultbut?language=objc for details.
func (x gen_NSWindow) EnableKeyEquivalentForDefaultButtonCell() {
	C.NSWindow_inst_EnableKeyEquivalentForDefaultButtonCell(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// EnableSnapshotRestoration enables snapshot restoration.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1525288-enablesnapshotrestoration?language=objc for details.
func (x gen_NSWindow) EnableSnapshotRestoration() {
	C.NSWindow_inst_EnableSnapshotRestoration(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// EndEditingFor forces the field editor to give up its first responder status and prepares it for its next assignment.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419469-endeditingfor?language=objc for details.
func (x gen_NSWindow) EndEditingFor(
	object objc.Ref,
) {
	C.NSWindow_inst_EndEditingFor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)

	return
}

// EndSheet ends a document-modal session and dismisses the specified sheet.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419318-endsheet?language=objc for details.
func (x gen_NSWindow) EndSheet(
	sheetWindow NSWindowRef,
) {
	C.NSWindow_inst_EndSheet(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sheetWindow),
	)

	return
}

// FieldEditorForObject returns the window’s field editor, creating it if requested.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419647-fieldeditor?language=objc for details.
func (x gen_NSWindow) FieldEditorForObject(
	createFlag bool,
	object objc.Ref,
) NSText {
	ret := C.NSWindow_inst_FieldEditorForObject(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(createFlag),
		objc.RefPointer(object),
	)

	return NSText_FromPointer(ret)
}

// FrameRectForContentRect returns the window’s frame rectangle with a given content rectangle.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419134-framerectforcontentrect?language=objc for details.
func (x gen_NSWindow) FrameRectForContentRect(
	contentRect core.NSRect,
) core.NSRect {
	ret := C.NSWindow_inst_FrameRectForContentRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentRect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// InitWithContentRectStyleMaskBackingDefer initializes the window with the specified values.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419477-initwithcontentrect?language=objc for details.
func (x gen_NSWindow) InitWithContentRectStyleMaskBackingDefer(
	contentRect core.NSRect,
	style core.NSUInteger,
	backingStoreType core.NSUInteger,
	flag bool,
) NSWindow {
	ret := C.NSWindow_inst_InitWithContentRectStyleMaskBackingDefer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentRect)),
		C.ulong(style),
		C.ulong(backingStoreType),
		convertToObjCBool(flag),
	)

	return NSWindow_FromPointer(ret)
}

// InitWithContentRectStyleMaskBackingDeferScreen initializes an allocated window with the specified values.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419755-initwithcontentrect?language=objc for details.
func (x gen_NSWindow) InitWithContentRectStyleMaskBackingDeferScreen(
	contentRect core.NSRect,
	style core.NSUInteger,
	backingStoreType core.NSUInteger,
	flag bool,
	screen NSScreenRef,
) NSWindow {
	ret := C.NSWindow_inst_InitWithContentRectStyleMaskBackingDeferScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentRect)),
		C.ulong(style),
		C.ulong(backingStoreType),
		convertToObjCBool(flag),
		objc.RefPointer(screen),
	)

	return NSWindow_FromPointer(ret)
}

// InvalidateCursorRectsForView marks as invalid the cursor rectangles of a given view object in the window, so they’ll be set up again when the window becomes key.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419601-invalidatecursorrectsforview?language=objc for details.
func (x gen_NSWindow) InvalidateCursorRectsForView(
	view NSViewRef,
) {
	C.NSWindow_inst_InvalidateCursorRectsForView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)

	return
}

// InvalidateShadow invalidates the window shadow so that it is recomputed based on the current window shape.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419529-invalidateshadow?language=objc for details.
func (x gen_NSWindow) InvalidateShadow() {
	C.NSWindow_inst_InvalidateShadow(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// LayoutIfNeeded updates the layout of views in the window based on the current views and constraints.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1526910-layoutifneeded?language=objc for details.
func (x gen_NSWindow) LayoutIfNeeded() {
	C.NSWindow_inst_LayoutIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// MakeKeyAndOrderFront moves the window to the front of the screen list, within its level, and makes it the key window; that is, it shows the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419208-makekeyandorderfront?language=objc for details.
func (x gen_NSWindow) MakeKeyAndOrderFront(
	sender objc.Ref,
) {
	C.NSWindow_inst_MakeKeyAndOrderFront(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// MakeKeyWindow makes the window the key window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419368-makekeywindow?language=objc for details.
func (x gen_NSWindow) MakeKeyWindow() {
	C.NSWindow_inst_MakeKeyWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// MakeMainWindow makes the window the main window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419271-makemainwindow?language=objc for details.
func (x gen_NSWindow) MakeMainWindow() {
	C.NSWindow_inst_MakeMainWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// MergeAllWindows merges all open windows into a single tabbed window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1644639-mergeallwindows?language=objc for details.
func (x gen_NSWindow) MergeAllWindows(
	sender objc.Ref,
) {
	C.NSWindow_inst_MergeAllWindows(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// Miniaturize removes the window from the screen list and displays the minimized window in the Dock.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419426-miniaturize?language=objc for details.
func (x gen_NSWindow) Miniaturize(
	sender objc.Ref,
) {
	C.NSWindow_inst_Miniaturize(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// MoveTabToNewWindow moves the tab to a new containing window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1644410-movetabtonewwindow?language=objc for details.
func (x gen_NSWindow) MoveTabToNewWindow(
	sender objc.Ref,
) {
	C.NSWindow_inst_MoveTabToNewWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// OrderBack moves the window to the back of its level in the screen list, without changing either the key window or the main window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419204-orderback?language=objc for details.
func (x gen_NSWindow) OrderBack(
	sender objc.Ref,
) {
	C.NSWindow_inst_OrderBack(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// OrderFront moves the window to the front of its level in the screen list, without changing either the key window or the main window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419495-orderfront?language=objc for details.
func (x gen_NSWindow) OrderFront(
	sender objc.Ref,
) {
	C.NSWindow_inst_OrderFront(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// OrderFrontRegardless moves the window to the front of its level, even if its application isn’t active, without changing either the key window or the main window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419444-orderfrontregardless?language=objc for details.
func (x gen_NSWindow) OrderFrontRegardless() {
	C.NSWindow_inst_OrderFrontRegardless(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// OrderOut removes the window from the screen list, which hides the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419660-orderout?language=objc for details.
func (x gen_NSWindow) OrderOut(
	sender objc.Ref,
) {
	C.NSWindow_inst_OrderOut(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// OrderWindowRelativeTo repositions the window’s window device in the window server’s screen list.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419672-orderwindow?language=objc for details.
func (x gen_NSWindow) OrderWindowRelativeTo(
	place core.NSUInteger,
	otherWin core.NSInteger,
) {
	C.NSWindow_inst_OrderWindowRelativeTo(
		unsafe.Pointer(x.Pointer()),
		C.ulong(place),
		C.long(otherWin),
	)

	return
}

// PerformClose simulates the user clicking the close button by momentarily highlighting the button and then closing the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419288-performclose?language=objc for details.
func (x gen_NSWindow) PerformClose(
	sender objc.Ref,
) {
	C.NSWindow_inst_PerformClose(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// PerformMiniaturize simulates the user clicking the minimize button by momentarily highlighting the button, then minimizing the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419749-performminiaturize?language=objc for details.
func (x gen_NSWindow) PerformMiniaturize(
	sender objc.Ref,
) {
	C.NSWindow_inst_PerformMiniaturize(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// PerformWindowDragWithEvent starts a window drag based on the specified mouse-down event.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419386-performwindowdragwithevent?language=objc for details.
func (x gen_NSWindow) PerformWindowDragWithEvent(
	event NSEventRef,
) {
	C.NSWindow_inst_PerformWindowDragWithEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return
}

// PerformZoom this action method simulates the user clicking the zoom box by momentarily highlighting the button and then zooming the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419450-performzoom?language=objc for details.
func (x gen_NSWindow) PerformZoom(
	sender objc.Ref,
) {
	C.NSWindow_inst_PerformZoom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// PostEventAtStart forwards the message to the global application object.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419376-postevent?language=objc for details.
func (x gen_NSWindow) PostEventAtStart(
	event NSEventRef,
	flag bool,
) {
	C.NSWindow_inst_PostEventAtStart(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
		convertToObjCBool(flag),
	)

	return
}

// Print runs the Print panel, and if the user chooses an option other than canceling, prints the window (its frame view and all subviews).
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419767-print?language=objc for details.
func (x gen_NSWindow) Print(
	sender objc.Ref,
) {
	C.NSWindow_inst_Print(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// RecalculateKeyViewLoop marks the key view loop as “dirty” and in need of recalculation.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419350-recalculatekeyviewloop?language=objc for details.
func (x gen_NSWindow) RecalculateKeyViewLoop() {
	C.NSWindow_inst_RecalculateKeyViewLoop(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// RegisterForDraggedTypes registers a set of pasteboard types that the window accepts as the destination of an image-dragging session.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419140-registerfordraggedtypes?language=objc for details.
func (x gen_NSWindow) RegisterForDraggedTypes(
	newTypes core.NSArrayRef,
) {
	C.NSWindow_inst_RegisterForDraggedTypes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newTypes),
	)

	return
}

// RemoveChildWindow detaches a given child window from the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419063-removechildwindow?language=objc for details.
func (x gen_NSWindow) RemoveChildWindow(
	childWin NSWindowRef,
) {
	C.NSWindow_inst_RemoveChildWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(childWin),
	)

	return
}

// RemoveTitlebarAccessoryViewControllerAtIndex removes the view controller at the specified index from the window’s array of title bar accessory view controllers.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419643-removetitlebaraccessoryviewcontr?language=objc for details.
func (x gen_NSWindow) RemoveTitlebarAccessoryViewControllerAtIndex(
	index core.NSInteger,
) {
	C.NSWindow_inst_RemoveTitlebarAccessoryViewControllerAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)

	return
}

// ResetCursorRects clears the window’s cursor rectangles and the cursor rectangles of the NSView objects in its view hierarchy.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419464-resetcursorrects?language=objc for details.
func (x gen_NSWindow) ResetCursorRects() {
	C.NSWindow_inst_ResetCursorRects(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ResignKeyWindow resigns the window’s key window status.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419047-resignkeywindow?language=objc for details.
func (x gen_NSWindow) ResignKeyWindow() {
	C.NSWindow_inst_ResignKeyWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ResignMainWindow resigns the window’s main window status.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419212-resignmainwindow?language=objc for details.
func (x gen_NSWindow) ResignMainWindow() {
	C.NSWindow_inst_ResignMainWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// RunToolbarCustomizationPalette presents the toolbar customization user interface.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419284-runtoolbarcustomizationpalette?language=objc for details.
func (x gen_NSWindow) RunToolbarCustomizationPalette(
	sender objc.Ref,
) {
	C.NSWindow_inst_RunToolbarCustomizationPalette(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// SelectKeyViewFollowingView gives key view status to the view that follows the given view.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419633-selectkeyviewfollowingview?language=objc for details.
func (x gen_NSWindow) SelectKeyViewFollowingView(
	view NSViewRef,
) {
	C.NSWindow_inst_SelectKeyViewFollowingView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)

	return
}

// SelectKeyViewPrecedingView gives key view status to the view that precedes the given view.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419757-selectkeyviewprecedingview?language=objc for details.
func (x gen_NSWindow) SelectKeyViewPrecedingView(
	view NSViewRef,
) {
	C.NSWindow_inst_SelectKeyViewPrecedingView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)

	return
}

// SelectNextKeyView searches for a candidate next key view and, if it finds one, tries to make it the first responder.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419715-selectnextkeyview?language=objc for details.
func (x gen_NSWindow) SelectNextKeyView(
	sender objc.Ref,
) {
	C.NSWindow_inst_SelectNextKeyView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// SelectNextTab selects the next tab in the tab group in the trailing direction.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1644693-selectnexttab?language=objc for details.
func (x gen_NSWindow) SelectNextTab(
	sender objc.Ref,
) {
	C.NSWindow_inst_SelectNextTab(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// SelectPreviousKeyView searches for a candidate previous key view and, if it finds one, tries to make it the first responder.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419110-selectpreviouskeyview?language=objc for details.
func (x gen_NSWindow) SelectPreviousKeyView(
	sender objc.Ref,
) {
	C.NSWindow_inst_SelectPreviousKeyView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// SelectPreviousTab selects the previous tab in the tab group in the leading direction.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1644555-selectprevioustab?language=objc for details.
func (x gen_NSWindow) SelectPreviousTab(
	sender objc.Ref,
) {
	C.NSWindow_inst_SelectPreviousTab(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// SendEvent this action method dispatches mouse and keyboard events the global application object sends to the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419228-sendevent?language=objc for details.
func (x gen_NSWindow) SendEvent(
	event NSEventRef,
) {
	C.NSWindow_inst_SendEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return
}

// SetContentSize sets the size of the window’s content view to a given size, which is expressed in the window’s base coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419100-setcontentsize?language=objc for details.
func (x gen_NSWindow) SetContentSize(
	size core.NSSize,
) {
	C.NSWindow_inst_SetContentSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return
}

// SetDynamicDepthLimit sets a Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419473-setdynamicdepthlimit?language=objc for details.
func (x gen_NSWindow) SetDynamicDepthLimit(
	flag bool,
) {
	C.NSWindow_inst_SetDynamicDepthLimit(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)

	return
}

// SetFrameDisplay sets the origin and size of the window’s frame rectangle according to a given frame rectangle, thereby setting its position and size onscreen.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419753-setframe?language=objc for details.
func (x gen_NSWindow) SetFrameDisplay(
	frameRect core.NSRect,
	flag bool,
) {
	C.NSWindow_inst_SetFrameDisplay(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
		convertToObjCBool(flag),
	)

	return
}

// SetFrameDisplayAnimate sets the origin and size of the window’s frame rectangle, with optional animation, according to a given frame rectangle, thereby setting its position and size onscreen.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419519-setframe?language=objc for details.
func (x gen_NSWindow) SetFrameDisplayAnimate(
	frameRect core.NSRect,
	displayFlag bool,
	animateFlag bool,
) {
	C.NSWindow_inst_SetFrameDisplayAnimate(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
		convertToObjCBool(displayFlag),
		convertToObjCBool(animateFlag),
	)

	return
}

// SetFrameOrigin positions the bottom-left corner of the window’s frame rectangle at a given point in screen coordinates.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419690-setframeorigin?language=objc for details.
func (x gen_NSWindow) SetFrameOrigin(
	point core.NSPoint,
) {
	C.NSWindow_inst_SetFrameOrigin(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return
}

// SetFrameTopLeftPoint positions the top-left corner of the window’s frame rectangle at a given point in screen coordinates.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419658-setframetopleftpoint?language=objc for details.
func (x gen_NSWindow) SetFrameTopLeftPoint(
	point core.NSPoint,
) {
	C.NSWindow_inst_SetFrameTopLeftPoint(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return
}

// SetIsMiniaturized sets the window’s miniaturized state to the value you specify.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1449566-setisminiaturized?language=objc for details.
func (x gen_NSWindow) SetIsMiniaturized(
	flag bool,
) {
	C.NSWindow_inst_SetIsMiniaturized(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)

	return
}

// SetIsVisible sets the window’s visible state to the value you specify.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1449570-setisvisible?language=objc for details.
func (x gen_NSWindow) SetIsVisible(
	flag bool,
) {
	C.NSWindow_inst_SetIsVisible(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)

	return
}

// SetIsZoomed sets the window’s zoomed state to the value you specify.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1449589-setiszoomed?language=objc for details.
func (x gen_NSWindow) SetIsZoomed(
	flag bool,
) {
	C.NSWindow_inst_SetIsZoomed(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)

	return
}

// SetTitleWithRepresentedFilename sets a given path as the window’s title, formatting it as a file-system path, and records this path as the window’s associated file.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419192-settitlewithrepresentedfilename?language=objc for details.
func (x gen_NSWindow) SetTitleWithRepresentedFilename(
	filename core.NSStringRef,
) {
	C.NSWindow_inst_SetTitleWithRepresentedFilename(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(filename),
	)

	return
}

// ToggleFullScreen takes the window into or out of fullscreen mode,
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419527-togglefullscreen?language=objc for details.
func (x gen_NSWindow) ToggleFullScreen(
	sender objc.Ref,
) {
	C.NSWindow_inst_ToggleFullScreen(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleTabBar shows or hides the tab bar.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1644517-toggletabbar?language=objc for details.
func (x gen_NSWindow) ToggleTabBar(
	sender objc.Ref,
) {
	C.NSWindow_inst_ToggleTabBar(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleTabOverview shows or hides the tab overview.
//
// See https://developer.apple.com/documentation/appkit/nswindow/2870175-toggletaboverview?language=objc for details.
func (x gen_NSWindow) ToggleTabOverview(
	sender objc.Ref,
) {
	C.NSWindow_inst_ToggleTabOverview(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleToolbarShown toggles the visibility of the window’s toolbar.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419554-toggletoolbarshown?language=objc for details.
func (x gen_NSWindow) ToggleToolbarShown(
	sender objc.Ref,
) {
	C.NSWindow_inst_ToggleToolbarShown(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// TryToPerformWith dispatches action messages with a given argument.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419428-trytoperform?language=objc for details.
func (x gen_NSWindow) TryToPerformWith(
	action objc.Selector,
	object objc.Ref,
) bool {
	ret := C.NSWindow_inst_TryToPerformWith(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(object),
	)

	return convertObjCBoolToGo(ret)
}

// UnregisterDraggedTypes unregisters the window as a possible destination for dragging operations.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419456-unregisterdraggedtypes?language=objc for details.
func (x gen_NSWindow) UnregisterDraggedTypes() {
	C.NSWindow_inst_UnregisterDraggedTypes(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Update updates the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419577-update?language=objc for details.
func (x gen_NSWindow) Update() {
	C.NSWindow_inst_Update(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// UpdateConstraintsIfNeeded updates the constraints based on changes to views in the window since the last layout.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1526915-updateconstraintsifneeded?language=objc for details.
func (x gen_NSWindow) UpdateConstraintsIfNeeded() {
	C.NSWindow_inst_UpdateConstraintsIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// VisualizeConstraints displays a visual representation of the supplied constraints in the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1526997-visualizeconstraints?language=objc for details.
func (x gen_NSWindow) VisualizeConstraints(
	constraints core.NSArrayRef,
) {
	C.NSWindow_inst_VisualizeConstraints(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(constraints),
	)

	return
}

// Zoom toggles the size and location of the window between its standard state (which the application provides as the best size to display the window’s data) and its user state (a new size and location the user may have set by moving or resizing the window).
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419513-zoom?language=objc for details.
func (x gen_NSWindow) Zoom(
	sender objc.Ref,
) {
	C.NSWindow_inst_Zoom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// Init is undocumented.
func (x gen_NSWindow) Init() NSWindow {
	ret := C.NSWindow_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSWindow_FromPointer(ret)
}

// Delegate returns the window’s delegate.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419060-delegate?language=objc for details.
func (x gen_NSWindow) Delegate() objc.Object {
	ret := C.NSWindow_inst_Delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetDelegate returns the window’s delegate.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419060-delegate?language=objc for details.
func (x gen_NSWindow) SetDelegate(
	value objc.Ref,
) {
	C.NSWindow_inst_SetDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// ContentViewController returns the main content view controller for the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419615-contentviewcontroller?language=objc for details.
func (x gen_NSWindow) ContentViewController() NSViewController {
	ret := C.NSWindow_inst_ContentViewController(
		unsafe.Pointer(x.Pointer()),
	)

	return NSViewController_FromPointer(ret)
}

// SetContentViewController returns the main content view controller for the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419615-contentviewcontroller?language=objc for details.
func (x gen_NSWindow) SetContentViewController(
	value NSViewControllerRef,
) {
	C.NSWindow_inst_SetContentViewController(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// ContentView returns the window’s content view, the highest accessible view object in the window’s view hierarchy.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419160-contentview?language=objc for details.
func (x gen_NSWindow) ContentView() NSView {
	ret := C.NSWindow_inst_ContentView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_FromPointer(ret)
}

// SetContentView returns the window’s content view, the highest accessible view object in the window’s view hierarchy.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419160-contentview?language=objc for details.
func (x gen_NSWindow) SetContentView(
	value NSViewRef,
) {
	C.NSWindow_inst_SetContentView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// StyleMask flags that describe the window’s current style, such as if it’s resizable or in full-screen mode.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419078-stylemask?language=objc for details.
func (x gen_NSWindow) StyleMask() core.NSUInteger {
	ret := C.NSWindow_inst_StyleMask(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// SetStyleMask flags that describe the window’s current style, such as if it’s resizable or in full-screen mode.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419078-stylemask?language=objc for details.
func (x gen_NSWindow) SetStyleMask(
	value core.NSUInteger,
) {
	C.NSWindow_inst_SetStyleMask(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)

	return
}

// WorksWhenModal returns a Boolean value that indicates whether the window is able to receive keyboard and mouse events even when some other window is being run modally.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419220-workswhenmodal?language=objc for details.
func (x gen_NSWindow) WorksWhenModal() bool {
	ret := C.NSWindow_inst_WorksWhenModal(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AlphaValue returns the window’s alpha value.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419186-alphavalue?language=objc for details.
func (x gen_NSWindow) AlphaValue() core.CGFloat {
	ret := C.NSWindow_inst_AlphaValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// SetAlphaValue returns the window’s alpha value.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419186-alphavalue?language=objc for details.
func (x gen_NSWindow) SetAlphaValue(
	value core.CGFloat,
) {
	C.NSWindow_inst_SetAlphaValue(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// BackgroundColor returns the color of the window’s background.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419751-backgroundcolor?language=objc for details.
func (x gen_NSWindow) BackgroundColor() NSColor {
	ret := C.NSWindow_inst_BackgroundColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_FromPointer(ret)
}

// SetBackgroundColor returns the color of the window’s background.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419751-backgroundcolor?language=objc for details.
func (x gen_NSWindow) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSWindow_inst_SetBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// CanHide returns a Boolean value that indicates whether the window can hide when its application becomes hidden.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419725-canhide?language=objc for details.
func (x gen_NSWindow) CanHide() bool {
	ret := C.NSWindow_inst_CanHide(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetCanHide returns a Boolean value that indicates whether the window can hide when its application becomes hidden.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419725-canhide?language=objc for details.
func (x gen_NSWindow) SetCanHide(
	value bool,
) {
	C.NSWindow_inst_SetCanHide(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsOnActiveSpace returns a Boolean value that indicates whether the window is on the currently active space.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419707-onactivespace?language=objc for details.
func (x gen_NSWindow) IsOnActiveSpace() bool {
	ret := C.NSWindow_inst_IsOnActiveSpace(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// HidesOnDeactivate returns a Boolean value that indicates whether the window is removed from the screen when its application becomes inactive.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419777-hidesondeactivate?language=objc for details.
func (x gen_NSWindow) HidesOnDeactivate() bool {
	ret := C.NSWindow_inst_HidesOnDeactivate(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetHidesOnDeactivate returns a Boolean value that indicates whether the window is removed from the screen when its application becomes inactive.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419777-hidesondeactivate?language=objc for details.
func (x gen_NSWindow) SetHidesOnDeactivate(
	value bool,
) {
	C.NSWindow_inst_SetHidesOnDeactivate(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// CollectionBehavior returns a value that identifies the window’s behavior in window collections.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419471-collectionbehavior?language=objc for details.
func (x gen_NSWindow) CollectionBehavior() core.NSUInteger {
	ret := C.NSWindow_inst_CollectionBehavior(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// SetCollectionBehavior returns a value that identifies the window’s behavior in window collections.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419471-collectionbehavior?language=objc for details.
func (x gen_NSWindow) SetCollectionBehavior(
	value core.NSUInteger,
) {
	C.NSWindow_inst_SetCollectionBehavior(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)

	return
}

// IsOpaque returns a Boolean value that indicates whether the window is opaque.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419086-opaque?language=objc for details.
func (x gen_NSWindow) IsOpaque() bool {
	ret := C.NSWindow_inst_IsOpaque(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetOpaque returns a Boolean value that indicates whether the window is opaque.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419086-opaque?language=objc for details.
func (x gen_NSWindow) SetOpaque(
	value bool,
) {
	C.NSWindow_inst_SetOpaque(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// HasShadow returns a Boolean value that indicates whether the window has a shadow.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419234-hasshadow?language=objc for details.
func (x gen_NSWindow) HasShadow() bool {
	ret := C.NSWindow_inst_HasShadow(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetHasShadow returns a Boolean value that indicates whether the window has a shadow.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419234-hasshadow?language=objc for details.
func (x gen_NSWindow) SetHasShadow(
	value bool,
) {
	C.NSWindow_inst_SetHasShadow(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// PreventsApplicationTerminationWhenModal returns a Boolean value that indicates whether the window prevents application termination when modal.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419743-preventsapplicationterminationwh?language=objc for details.
func (x gen_NSWindow) PreventsApplicationTerminationWhenModal() bool {
	ret := C.NSWindow_inst_PreventsApplicationTerminationWhenModal(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetPreventsApplicationTerminationWhenModal returns a Boolean value that indicates whether the window prevents application termination when modal.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419743-preventsapplicationterminationwh?language=objc for details.
func (x gen_NSWindow) SetPreventsApplicationTerminationWhenModal(
	value bool,
) {
	C.NSWindow_inst_SetPreventsApplicationTerminationWhenModal(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// HasDynamicDepthLimit returns a Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419330-hasdynamicdepthlimit?language=objc for details.
func (x gen_NSWindow) HasDynamicDepthLimit() bool {
	ret := C.NSWindow_inst_HasDynamicDepthLimit(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// WindowNumber returns the window number of the window’s window device.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419068-windownumber?language=objc for details.
func (x gen_NSWindow) WindowNumber() core.NSInteger {
	ret := C.NSWindow_inst_WindowNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// DeviceDescription returns a dictionary containing information about the window’s resolution, such as color, depth, and so on.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419741-devicedescription?language=objc for details.
func (x gen_NSWindow) DeviceDescription() core.NSDictionary {
	ret := C.NSWindow_inst_DeviceDescription(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_FromPointer(ret)
}

// CanBecomeVisibleWithoutLogin returns a Boolean value that indicates whether the window can be displayed at the login window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419179-canbecomevisiblewithoutlogin?language=objc for details.
func (x gen_NSWindow) CanBecomeVisibleWithoutLogin() bool {
	ret := C.NSWindow_inst_CanBecomeVisibleWithoutLogin(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetCanBecomeVisibleWithoutLogin returns a Boolean value that indicates whether the window can be displayed at the login window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419179-canbecomevisiblewithoutlogin?language=objc for details.
func (x gen_NSWindow) SetCanBecomeVisibleWithoutLogin(
	value bool,
) {
	C.NSWindow_inst_SetCanBecomeVisibleWithoutLogin(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// BackingType returns the window’s backing store type.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419599-backingtype?language=objc for details.
func (x gen_NSWindow) BackingType() core.NSUInteger {
	ret := C.NSWindow_inst_BackingType(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)
}

// SetBackingType returns the window’s backing store type.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419599-backingtype?language=objc for details.
func (x gen_NSWindow) SetBackingType(
	value core.NSUInteger,
) {
	C.NSWindow_inst_SetBackingType(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)

	return
}

// AttachedSheet returns the sheet attached to the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419467-attachedsheet?language=objc for details.
func (x gen_NSWindow) AttachedSheet() NSWindow {
	ret := C.NSWindow_inst_AttachedSheet(
		unsafe.Pointer(x.Pointer()),
	)

	return NSWindow_FromPointer(ret)
}

// IsSheet returns a Boolean value that indicates whether the window has ever run as a modal sheet.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419364-sheet?language=objc for details.
func (x gen_NSWindow) IsSheet() bool {
	ret := C.NSWindow_inst_IsSheet(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SheetParent returns the window to which the sheet is attached.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419052-sheetparent?language=objc for details.
func (x gen_NSWindow) SheetParent() NSWindow {
	ret := C.NSWindow_inst_SheetParent(
		unsafe.Pointer(x.Pointer()),
	)

	return NSWindow_FromPointer(ret)
}

// Sheets an array of the sheets currently attached to the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419765-sheets?language=objc for details.
func (x gen_NSWindow) Sheets() core.NSArray {
	ret := C.NSWindow_inst_Sheets(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// Frame returns the window’s frame rectangle in screen coordinates, including the title bar.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419697-frame?language=objc for details.
func (x gen_NSWindow) Frame() core.NSRect {
	ret := C.NSWindow_inst_Frame(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// AspectRatio returns the window’s aspect ratio, which constrains the size of its frame rectangle to integral multiples of this ratio when the user resizes it.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419507-aspectratio?language=objc for details.
func (x gen_NSWindow) AspectRatio() core.NSSize {
	ret := C.NSWindow_inst_AspectRatio(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetAspectRatio returns the window’s aspect ratio, which constrains the size of its frame rectangle to integral multiples of this ratio when the user resizes it.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419507-aspectratio?language=objc for details.
func (x gen_NSWindow) SetAspectRatio(
	value core.NSSize,
) {
	C.NSWindow_inst_SetAspectRatio(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// MinSize returns the minimum size to which the window’s frame (including its title bar) can be sized.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419206-minsize?language=objc for details.
func (x gen_NSWindow) MinSize() core.NSSize {
	ret := C.NSWindow_inst_MinSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetMinSize returns the minimum size to which the window’s frame (including its title bar) can be sized.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419206-minsize?language=objc for details.
func (x gen_NSWindow) SetMinSize(
	value core.NSSize,
) {
	C.NSWindow_inst_SetMinSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// MaxSize returns the maximum size to which the window’s frame (including its title bar) can be sized.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419595-maxsize?language=objc for details.
func (x gen_NSWindow) MaxSize() core.NSSize {
	ret := C.NSWindow_inst_MaxSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetMaxSize returns the maximum size to which the window’s frame (including its title bar) can be sized.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419595-maxsize?language=objc for details.
func (x gen_NSWindow) SetMaxSize(
	value core.NSSize,
) {
	C.NSWindow_inst_SetMaxSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// IsZoomed returns a Boolean value that indicates whether the window is in a zoomed state.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419398-zoomed?language=objc for details.
func (x gen_NSWindow) IsZoomed() bool {
	ret := C.NSWindow_inst_IsZoomed(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// ResizeIncrements returns the window’s resizing increments.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419390-resizeincrements?language=objc for details.
func (x gen_NSWindow) ResizeIncrements() core.NSSize {
	ret := C.NSWindow_inst_ResizeIncrements(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetResizeIncrements returns the window’s resizing increments.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419390-resizeincrements?language=objc for details.
func (x gen_NSWindow) SetResizeIncrements(
	value core.NSSize,
) {
	C.NSWindow_inst_SetResizeIncrements(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// PreservesContentDuringLiveResize returns a Boolean value that indicates whether the window tries to optimize user-initiated resize operations by preserving the content of views that have not changed.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419588-preservescontentduringliveresize?language=objc for details.
func (x gen_NSWindow) PreservesContentDuringLiveResize() bool {
	ret := C.NSWindow_inst_PreservesContentDuringLiveResize(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetPreservesContentDuringLiveResize returns a Boolean value that indicates whether the window tries to optimize user-initiated resize operations by preserving the content of views that have not changed.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419588-preservescontentduringliveresize?language=objc for details.
func (x gen_NSWindow) SetPreservesContentDuringLiveResize(
	value bool,
) {
	C.NSWindow_inst_SetPreservesContentDuringLiveResize(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// InLiveResize returns a Boolean value that indicates whether the window is being resized by the user.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419378-inliveresize?language=objc for details.
func (x gen_NSWindow) InLiveResize() bool {
	ret := C.NSWindow_inst_InLiveResize(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// ContentAspectRatio returns the window’s content aspect ratio.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419148-contentaspectratio?language=objc for details.
func (x gen_NSWindow) ContentAspectRatio() core.NSSize {
	ret := C.NSWindow_inst_ContentAspectRatio(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetContentAspectRatio returns the window’s content aspect ratio.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419148-contentaspectratio?language=objc for details.
func (x gen_NSWindow) SetContentAspectRatio(
	value core.NSSize,
) {
	C.NSWindow_inst_SetContentAspectRatio(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// ContentMinSize returns the minimum size of the window’s content view in the window’s base coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419670-contentminsize?language=objc for details.
func (x gen_NSWindow) ContentMinSize() core.NSSize {
	ret := C.NSWindow_inst_ContentMinSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetContentMinSize returns the minimum size of the window’s content view in the window’s base coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419670-contentminsize?language=objc for details.
func (x gen_NSWindow) SetContentMinSize(
	value core.NSSize,
) {
	C.NSWindow_inst_SetContentMinSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// ContentMaxSize returns the maximum size of the window’s content view in the window’s base coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419154-contentmaxsize?language=objc for details.
func (x gen_NSWindow) ContentMaxSize() core.NSSize {
	ret := C.NSWindow_inst_ContentMaxSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetContentMaxSize returns the maximum size of the window’s content view in the window’s base coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419154-contentmaxsize?language=objc for details.
func (x gen_NSWindow) SetContentMaxSize(
	value core.NSSize,
) {
	C.NSWindow_inst_SetContentMaxSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// ContentResizeIncrements returns the window’s content-view resizing increments.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419649-contentresizeincrements?language=objc for details.
func (x gen_NSWindow) ContentResizeIncrements() core.NSSize {
	ret := C.NSWindow_inst_ContentResizeIncrements(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetContentResizeIncrements returns the window’s content-view resizing increments.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419649-contentresizeincrements?language=objc for details.
func (x gen_NSWindow) SetContentResizeIncrements(
	value core.NSSize,
) {
	C.NSWindow_inst_SetContentResizeIncrements(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// ContentLayoutGuide returns a value used by Auto Layout constraints to automatically bind to the value of contentLayoutRect.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419094-contentlayoutguide?language=objc for details.
func (x gen_NSWindow) ContentLayoutGuide() objc.Object {
	ret := C.NSWindow_inst_ContentLayoutGuide(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// ContentLayoutRect returns the area inside the window that is for non-obscured content, in window coordinates.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419124-contentlayoutrect?language=objc for details.
func (x gen_NSWindow) ContentLayoutRect() core.NSRect {
	ret := C.NSWindow_inst_ContentLayoutRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// MaxFullScreenContentSize returns a maximum size that is used to determine if a window can fit when it is in full screen in a tile.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419438-maxfullscreencontentsize?language=objc for details.
func (x gen_NSWindow) MaxFullScreenContentSize() core.NSSize {
	ret := C.NSWindow_inst_MaxFullScreenContentSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetMaxFullScreenContentSize returns a maximum size that is used to determine if a window can fit when it is in full screen in a tile.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419438-maxfullscreencontentsize?language=objc for details.
func (x gen_NSWindow) SetMaxFullScreenContentSize(
	value core.NSSize,
) {
	C.NSWindow_inst_SetMaxFullScreenContentSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// MinFullScreenContentSize returns a minimum size that is used to determine if a window can fit when it is in full screen in a tile.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419627-minfullscreencontentsize?language=objc for details.
func (x gen_NSWindow) MinFullScreenContentSize() core.NSSize {
	ret := C.NSWindow_inst_MinFullScreenContentSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetMinFullScreenContentSize returns a minimum size that is used to determine if a window can fit when it is in full screen in a tile.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419627-minfullscreencontentsize?language=objc for details.
func (x gen_NSWindow) SetMinFullScreenContentSize(
	value core.NSSize,
) {
	C.NSWindow_inst_SetMinFullScreenContentSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// Level returns the window level of the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419511-level?language=objc for details.
func (x gen_NSWindow) Level() core.NSInteger {
	ret := C.NSWindow_inst_Level(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// SetLevel returns the window level of the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419511-level?language=objc for details.
func (x gen_NSWindow) SetLevel(
	value core.NSInteger,
) {
	C.NSWindow_inst_SetLevel(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return
}

// IsVisible returns a Boolean value that indicates whether the window is visible onscreen (even when it’s obscured by other windows).
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419132-visible?language=objc for details.
func (x gen_NSWindow) IsVisible() bool {
	ret := C.NSWindow_inst_IsVisible(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsKeyWindow returns a Boolean value that indicates whether the window is the key window for the application.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419735-keywindow?language=objc for details.
func (x gen_NSWindow) IsKeyWindow() bool {
	ret := C.NSWindow_inst_IsKeyWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// CanBecomeKeyWindow returns a Boolean value that indicates whether the window can become the key window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419543-canbecomekeywindow?language=objc for details.
func (x gen_NSWindow) CanBecomeKeyWindow() bool {
	ret := C.NSWindow_inst_CanBecomeKeyWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsMainWindow returns a Boolean value that indicates whether the window is the application’s main window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419130-mainwindow?language=objc for details.
func (x gen_NSWindow) IsMainWindow() bool {
	ret := C.NSWindow_inst_IsMainWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// CanBecomeMainWindow returns a Boolean value that indicates whether the window can become the application’s main window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419162-canbecomemainwindow?language=objc for details.
func (x gen_NSWindow) CanBecomeMainWindow() bool {
	ret := C.NSWindow_inst_CanBecomeMainWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// ChildWindows an array of the window’s attached child windows.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419236-childwindows?language=objc for details.
func (x gen_NSWindow) ChildWindows() core.NSArray {
	ret := C.NSWindow_inst_ChildWindows(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// ParentWindow returns the parent window to which the window is attached as a child.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419695-parentwindow?language=objc for details.
func (x gen_NSWindow) ParentWindow() NSWindow {
	ret := C.NSWindow_inst_ParentWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return NSWindow_FromPointer(ret)
}

// SetParentWindow returns the parent window to which the window is attached as a child.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419695-parentwindow?language=objc for details.
func (x gen_NSWindow) SetParentWindow(
	value NSWindowRef,
) {
	C.NSWindow_inst_SetParentWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// IsExcludedFromWindowsMenu returns a Boolean value that indicates whether the window is excluded from the application’s Windows menu.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419175-excludedfromwindowsmenu?language=objc for details.
func (x gen_NSWindow) IsExcludedFromWindowsMenu() bool {
	ret := C.NSWindow_inst_IsExcludedFromWindowsMenu(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetExcludedFromWindowsMenu returns a Boolean value that indicates whether the window is excluded from the application’s Windows menu.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419175-excludedfromwindowsmenu?language=objc for details.
func (x gen_NSWindow) SetExcludedFromWindowsMenu(
	value bool,
) {
	C.NSWindow_inst_SetExcludedFromWindowsMenu(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AreCursorRectsEnabled returns a Boolean value that indicates whether the window’s cursor rectangles are enabled.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419668-arecursorrectsenabled?language=objc for details.
func (x gen_NSWindow) AreCursorRectsEnabled() bool {
	ret := C.NSWindow_inst_AreCursorRectsEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// ShowsToolbarButton returns a Boolean value that indicates whether the toolbar control button is currently displayed.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419196-showstoolbarbutton?language=objc for details.
func (x gen_NSWindow) ShowsToolbarButton() bool {
	ret := C.NSWindow_inst_ShowsToolbarButton(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetShowsToolbarButton returns a Boolean value that indicates whether the toolbar control button is currently displayed.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419196-showstoolbarbutton?language=objc for details.
func (x gen_NSWindow) SetShowsToolbarButton(
	value bool,
) {
	C.NSWindow_inst_SetShowsToolbarButton(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// TitlebarAppearsTransparent returns a Boolean value that indicates whether the title bar draws its background.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419167-titlebarappearstransparent?language=objc for details.
func (x gen_NSWindow) TitlebarAppearsTransparent() bool {
	ret := C.NSWindow_inst_TitlebarAppearsTransparent(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetTitlebarAppearsTransparent returns a Boolean value that indicates whether the title bar draws its background.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419167-titlebarappearstransparent?language=objc for details.
func (x gen_NSWindow) SetTitlebarAppearsTransparent(
	value bool,
) {
	C.NSWindow_inst_SetTitlebarAppearsTransparent(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// TitlebarAccessoryViewControllers an array of title bar accessory view controllers that are currently added to the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419547-titlebaraccessoryviewcontrollers?language=objc for details.
func (x gen_NSWindow) TitlebarAccessoryViewControllers() core.NSArray {
	ret := C.NSWindow_inst_TitlebarAccessoryViewControllers(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// SetTitlebarAccessoryViewControllers an array of title bar accessory view controllers that are currently added to the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419547-titlebaraccessoryviewcontrollers?language=objc for details.
func (x gen_NSWindow) SetTitlebarAccessoryViewControllers(
	value core.NSArrayRef,
) {
	C.NSWindow_inst_SetTitlebarAccessoryViewControllers(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// TabbedWindows an array of windows that display as tabs.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1792044-tabbedwindows?language=objc for details.
func (x gen_NSWindow) TabbedWindows() core.NSArray {
	ret := C.NSWindow_inst_TabbedWindows(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// AllowsToolTipsWhenApplicationIsInactive returns a Boolean value that indicates whether the window can display tooltips even when the application is in the background.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419138-allowstooltipswhenapplicationisi?language=objc for details.
func (x gen_NSWindow) AllowsToolTipsWhenApplicationIsInactive() bool {
	ret := C.NSWindow_inst_AllowsToolTipsWhenApplicationIsInactive(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsToolTipsWhenApplicationIsInactive returns a Boolean value that indicates whether the window can display tooltips even when the application is in the background.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419138-allowstooltipswhenapplicationisi?language=objc for details.
func (x gen_NSWindow) SetAllowsToolTipsWhenApplicationIsInactive(
	value bool,
) {
	C.NSWindow_inst_SetAllowsToolTipsWhenApplicationIsInactive(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// CurrentEvent returns the event currently being processed by the application.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419298-currentevent?language=objc for details.
func (x gen_NSWindow) CurrentEvent() NSEvent {
	ret := C.NSWindow_inst_CurrentEvent(
		unsafe.Pointer(x.Pointer()),
	)

	return NSEvent_FromPointer(ret)
}

// InitialFirstResponder returns the view that’s made first responder (also called the key view) the first time the window is placed onscreen.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419479-initialfirstresponder?language=objc for details.
func (x gen_NSWindow) InitialFirstResponder() NSView {
	ret := C.NSWindow_inst_InitialFirstResponder(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_FromPointer(ret)
}

// SetInitialFirstResponder returns the view that’s made first responder (also called the key view) the first time the window is placed onscreen.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419479-initialfirstresponder?language=objc for details.
func (x gen_NSWindow) SetInitialFirstResponder(
	value NSViewRef,
) {
	C.NSWindow_inst_SetInitialFirstResponder(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// AutorecalculatesKeyViewLoop returns a Boolean value that indicates whether the window automatically recalculates the key view loop when views are added.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419214-autorecalculateskeyviewloop?language=objc for details.
func (x gen_NSWindow) AutorecalculatesKeyViewLoop() bool {
	ret := C.NSWindow_inst_AutorecalculatesKeyViewLoop(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAutorecalculatesKeyViewLoop returns a Boolean value that indicates whether the window automatically recalculates the key view loop when views are added.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419214-autorecalculateskeyviewloop?language=objc for details.
func (x gen_NSWindow) SetAutorecalculatesKeyViewLoop(
	value bool,
) {
	C.NSWindow_inst_SetAutorecalculatesKeyViewLoop(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AcceptsMouseMovedEvents returns a Boolean value that indicates whether the window accepts mouse-moved events.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419340-acceptsmousemovedevents?language=objc for details.
func (x gen_NSWindow) AcceptsMouseMovedEvents() bool {
	ret := C.NSWindow_inst_AcceptsMouseMovedEvents(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAcceptsMouseMovedEvents returns a Boolean value that indicates whether the window accepts mouse-moved events.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419340-acceptsmousemovedevents?language=objc for details.
func (x gen_NSWindow) SetAcceptsMouseMovedEvents(
	value bool,
) {
	C.NSWindow_inst_SetAcceptsMouseMovedEvents(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IgnoresMouseEvents returns a Boolean value that indicates whether the window is transparent to mouse events.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419354-ignoresmouseevents?language=objc for details.
func (x gen_NSWindow) IgnoresMouseEvents() bool {
	ret := C.NSWindow_inst_IgnoresMouseEvents(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetIgnoresMouseEvents returns a Boolean value that indicates whether the window is transparent to mouse events.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419354-ignoresmouseevents?language=objc for details.
func (x gen_NSWindow) SetIgnoresMouseEvents(
	value bool,
) {
	C.NSWindow_inst_SetIgnoresMouseEvents(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// MouseLocationOutsideOfEventStream returns the current location of the pointer reckoned in the window’s base coordinate system, regardless of the current event being handled or of any events pending.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419280-mouselocationoutsideofeventstrea?language=objc for details.
func (x gen_NSWindow) MouseLocationOutsideOfEventStream() core.NSPoint {
	ret := C.NSWindow_inst_MouseLocationOutsideOfEventStream(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// IsRestorable returns a Boolean value indicating whether the window configuration is preserved between application launches.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1526255-restorable?language=objc for details.
func (x gen_NSWindow) IsRestorable() bool {
	ret := C.NSWindow_inst_IsRestorable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetRestorable returns a Boolean value indicating whether the window configuration is preserved between application launches.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1526255-restorable?language=objc for details.
func (x gen_NSWindow) SetRestorable(
	value bool,
) {
	C.NSWindow_inst_SetRestorable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// ViewsNeedDisplay returns a Boolean value that indicates whether any of the window’s views need to be displayed.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419609-viewsneeddisplay?language=objc for details.
func (x gen_NSWindow) ViewsNeedDisplay() bool {
	ret := C.NSWindow_inst_ViewsNeedDisplay(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetViewsNeedDisplay returns a Boolean value that indicates whether any of the window’s views need to be displayed.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419609-viewsneeddisplay?language=objc for details.
func (x gen_NSWindow) SetViewsNeedDisplay(
	value bool,
) {
	C.NSWindow_inst_SetViewsNeedDisplay(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AllowsConcurrentViewDrawing returns a Boolean value that indicates whether the window allows multithreaded view drawing.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419300-allowsconcurrentviewdrawing?language=objc for details.
func (x gen_NSWindow) AllowsConcurrentViewDrawing() bool {
	ret := C.NSWindow_inst_AllowsConcurrentViewDrawing(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsConcurrentViewDrawing returns a Boolean value that indicates whether the window allows multithreaded view drawing.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419300-allowsconcurrentviewdrawing?language=objc for details.
func (x gen_NSWindow) SetAllowsConcurrentViewDrawing(
	value bool,
) {
	C.NSWindow_inst_SetAllowsConcurrentViewDrawing(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsDocumentEdited returns a Boolean value that indicates whether the window’s document has been edited.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419311-documentedited?language=objc for details.
func (x gen_NSWindow) IsDocumentEdited() bool {
	ret := C.NSWindow_inst_IsDocumentEdited(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetDocumentEdited returns a Boolean value that indicates whether the window’s document has been edited.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419311-documentedited?language=objc for details.
func (x gen_NSWindow) SetDocumentEdited(
	value bool,
) {
	C.NSWindow_inst_SetDocumentEdited(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// BackingScaleFactor returns the backing scale factor.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419459-backingscalefactor?language=objc for details.
func (x gen_NSWindow) BackingScaleFactor() core.CGFloat {
	ret := C.NSWindow_inst_BackingScaleFactor(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// Title returns the string that appears in the title bar of the window or the path to the represented file.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419404-title?language=objc for details.
func (x gen_NSWindow) Title() core.NSString {
	ret := C.NSWindow_inst_Title(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetTitle returns the string that appears in the title bar of the window or the path to the represented file.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419404-title?language=objc for details.
func (x gen_NSWindow) SetTitle(
	value core.NSStringRef,
) {
	C.NSWindow_inst_SetTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Subtitle returns a secondary line of text that appears in the title bar of the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/3608198-subtitle?language=objc for details.
func (x gen_NSWindow) Subtitle() core.NSString {
	ret := C.NSWindow_inst_Subtitle(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetSubtitle returns a secondary line of text that appears in the title bar of the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/3608198-subtitle?language=objc for details.
func (x gen_NSWindow) SetSubtitle(
	value core.NSStringRef,
) {
	C.NSWindow_inst_SetSubtitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// TitleVisibility returns a value that indicates the visibility of the window’s title and title bar buttons.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419635-titlevisibility?language=objc for details.
func (x gen_NSWindow) TitleVisibility() core.NSInteger {
	ret := C.NSWindow_inst_TitleVisibility(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// SetTitleVisibility returns a value that indicates the visibility of the window’s title and title bar buttons.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419635-titlevisibility?language=objc for details.
func (x gen_NSWindow) SetTitleVisibility(
	value core.NSInteger,
) {
	C.NSWindow_inst_SetTitleVisibility(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return
}

// RepresentedFilename returns the path to the file of the window’s represented file.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419631-representedfilename?language=objc for details.
func (x gen_NSWindow) RepresentedFilename() core.NSString {
	ret := C.NSWindow_inst_RepresentedFilename(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetRepresentedFilename returns the path to the file of the window’s represented file.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419631-representedfilename?language=objc for details.
func (x gen_NSWindow) SetRepresentedFilename(
	value core.NSStringRef,
) {
	C.NSWindow_inst_SetRepresentedFilename(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// RepresentedURL returns the URL of the file the window represents.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419066-representedurl?language=objc for details.
func (x gen_NSWindow) RepresentedURL() core.NSURL {
	ret := C.NSWindow_inst_RepresentedURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_FromPointer(ret)
}

// SetRepresentedURL returns the URL of the file the window represents.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419066-representedurl?language=objc for details.
func (x gen_NSWindow) SetRepresentedURL(
	value core.NSURLRef,
) {
	C.NSWindow_inst_SetRepresentedURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Screen returns the screen the window is on.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419232-screen?language=objc for details.
func (x gen_NSWindow) Screen() NSScreen {
	ret := C.NSWindow_inst_Screen(
		unsafe.Pointer(x.Pointer()),
	)

	return NSScreen_FromPointer(ret)
}

// DeepestScreen returns the deepest screen the window is on (it may be split over several screens).
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419080-deepestscreen?language=objc for details.
func (x gen_NSWindow) DeepestScreen() NSScreen {
	ret := C.NSWindow_inst_DeepestScreen(
		unsafe.Pointer(x.Pointer()),
	)

	return NSScreen_FromPointer(ret)
}

// DisplaysWhenScreenProfileChanges returns a Boolean value that indicates whether the window context should be updated when the screen profile changes or when the window moves to a different screen.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419430-displayswhenscreenprofilechanges?language=objc for details.
func (x gen_NSWindow) DisplaysWhenScreenProfileChanges() bool {
	ret := C.NSWindow_inst_DisplaysWhenScreenProfileChanges(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetDisplaysWhenScreenProfileChanges returns a Boolean value that indicates whether the window context should be updated when the screen profile changes or when the window moves to a different screen.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419430-displayswhenscreenprofilechanges?language=objc for details.
func (x gen_NSWindow) SetDisplaysWhenScreenProfileChanges(
	value bool,
) {
	C.NSWindow_inst_SetDisplaysWhenScreenProfileChanges(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsMovableByWindowBackground returns a Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419072-movablebywindowbackground?language=objc for details.
func (x gen_NSWindow) IsMovableByWindowBackground() bool {
	ret := C.NSWindow_inst_IsMovableByWindowBackground(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetMovableByWindowBackground returns a Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419072-movablebywindowbackground?language=objc for details.
func (x gen_NSWindow) SetMovableByWindowBackground(
	value bool,
) {
	C.NSWindow_inst_SetMovableByWindowBackground(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsMovable returns a Boolean value that indicates whether the window can be dragged by clicking in its title bar or background.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419579-movable?language=objc for details.
func (x gen_NSWindow) IsMovable() bool {
	ret := C.NSWindow_inst_IsMovable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetMovable returns a Boolean value that indicates whether the window can be dragged by clicking in its title bar or background.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419579-movable?language=objc for details.
func (x gen_NSWindow) SetMovable(
	value bool,
) {
	C.NSWindow_inst_SetMovable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsReleasedWhenClosed returns a Boolean value that indicates whether the window is released when it receives the close message.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419062-releasedwhenclosed?language=objc for details.
func (x gen_NSWindow) IsReleasedWhenClosed() bool {
	ret := C.NSWindow_inst_IsReleasedWhenClosed(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetReleasedWhenClosed returns a Boolean value that indicates whether the window is released when it receives the close message.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419062-releasedwhenclosed?language=objc for details.
func (x gen_NSWindow) SetReleasedWhenClosed(
	value bool,
) {
	C.NSWindow_inst_SetReleasedWhenClosed(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsMiniaturized returns a Boolean value that indicates whether the window is minimized.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419699-miniaturized?language=objc for details.
func (x gen_NSWindow) IsMiniaturized() bool {
	ret := C.NSWindow_inst_IsMiniaturized(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// MiniwindowImage returns the custom miniaturized window image of the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419185-miniwindowimage?language=objc for details.
func (x gen_NSWindow) MiniwindowImage() NSImage {
	ret := C.NSWindow_inst_MiniwindowImage(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_FromPointer(ret)
}

// SetMiniwindowImage returns the custom miniaturized window image of the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419185-miniwindowimage?language=objc for details.
func (x gen_NSWindow) SetMiniwindowImage(
	value NSImageRef,
) {
	C.NSWindow_inst_SetMiniwindowImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// MiniwindowTitle returns the title displayed in the window’s minimized window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419571-miniwindowtitle?language=objc for details.
func (x gen_NSWindow) MiniwindowTitle() core.NSString {
	ret := C.NSWindow_inst_MiniwindowTitle(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetMiniwindowTitle returns the title displayed in the window’s minimized window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1419571-miniwindowtitle?language=objc for details.
func (x gen_NSWindow) SetMiniwindowTitle(
	value core.NSStringRef,
) {
	C.NSWindow_inst_SetMiniwindowTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// HasCloseBox returns a Boolean value that indicates if the window has a close box.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1449574-hasclosebox?language=objc for details.
func (x gen_NSWindow) HasCloseBox() bool {
	ret := C.NSWindow_inst_HasCloseBox(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// HasTitleBar returns a Boolean value that indicates if the window has a title bar.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1449568-hastitlebar?language=objc for details.
func (x gen_NSWindow) HasTitleBar() bool {
	ret := C.NSWindow_inst_HasTitleBar(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsModalPanel returns a Boolean value that indicates whether the window is a modal panel.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1449576-modalpanel?language=objc for details.
func (x gen_NSWindow) IsModalPanel() bool {
	ret := C.NSWindow_inst_IsModalPanel(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsFloatingPanel returns a Boolean value that indicates whether the window is a floating panel.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1449579-floatingpanel?language=objc for details.
func (x gen_NSWindow) IsFloatingPanel() bool {
	ret := C.NSWindow_inst_IsFloatingPanel(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsZoomable returns a Boolean value that indicates whether the window allows zooming.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1449587-zoomable?language=objc for details.
func (x gen_NSWindow) IsZoomable() bool {
	ret := C.NSWindow_inst_IsZoomable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsResizable returns a Boolean value that indicates if the user can resize the window.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1449572-resizable?language=objc for details.
func (x gen_NSWindow) IsResizable() bool {
	ret := C.NSWindow_inst_IsResizable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsMiniaturizable returns a Boolean value that indicates whether the window can minimize.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1449583-miniaturizable?language=objc for details.
func (x gen_NSWindow) IsMiniaturizable() bool {
	ret := C.NSWindow_inst_IsMiniaturizable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// OrderedIndex returns the zero-based position of the window, based on its order from front to back among all visible application windows.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1449577-orderedindex?language=objc for details.
func (x gen_NSWindow) OrderedIndex() core.NSInteger {
	ret := C.NSWindow_inst_OrderedIndex(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// SetOrderedIndex returns the zero-based position of the window, based on its order from front to back among all visible application windows.
//
// See https://developer.apple.com/documentation/appkit/nswindow/1449577-orderedindex?language=objc for details.
func (x gen_NSWindow) SetOrderedIndex(
	value core.NSInteger,
) {
	C.NSWindow_inst_SetOrderedIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return
}

type NSWorkspaceRef interface {
	Pointer() uintptr
	Init() NSWorkspace
}

type gen_NSWorkspace struct {
	objc.Object
}

func NSWorkspace_FromPointer(ptr unsafe.Pointer) NSWorkspace {
	return NSWorkspace{gen_NSWorkspace{
		objc.Object_FromPointer(ptr),
	}}
}

func NSWorkspace_FromRef(ref objc.Ref) NSWorkspace {
	return NSWorkspace_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// URLForApplicationToOpenURL returns the URL to the default app that would be opened.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1533391-urlforapplicationtoopenurl?language=objc for details.
func (x gen_NSWorkspace) URLForApplicationToOpenURL(
	url core.NSURLRef,
) core.NSURL {
	ret := C.NSWorkspace_inst_URLForApplicationToOpenURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)

	return core.NSURL_FromPointer(ret)
}

// URLForApplicationWithBundleIdentifier returns the URL for the app with the specified identifier.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1534053-urlforapplicationwithbundleident?language=objc for details.
func (x gen_NSWorkspace) URLForApplicationWithBundleIdentifier(
	bundleIdentifier core.NSStringRef,
) core.NSURL {
	ret := C.NSWorkspace_inst_URLForApplicationWithBundleIdentifier(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(bundleIdentifier),
	)

	return core.NSURL_FromPointer(ret)
}

// URLsForApplicationsToOpenURL is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/3753000-urlsforapplicationstoopenurl?language=objc for details.
func (x gen_NSWorkspace) URLsForApplicationsToOpenURL(
	url core.NSURLRef,
) core.NSArray {
	ret := C.NSWorkspace_inst_URLsForApplicationsToOpenURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)

	return core.NSArray_FromPointer(ret)
}

// URLsForApplicationsWithBundleIdentifier is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/3753001-urlsforapplicationswithbundleide?language=objc for details.
func (x gen_NSWorkspace) URLsForApplicationsWithBundleIdentifier(
	bundleIdentifier core.NSStringRef,
) core.NSArray {
	ret := C.NSWorkspace_inst_URLsForApplicationsWithBundleIdentifier(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(bundleIdentifier),
	)

	return core.NSArray_FromPointer(ret)
}

// ActivateFileViewerSelectingURLs activates the Finder, and opens one or more windows selecting the specified files.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1524549-activatefileviewerselectingurls?language=objc for details.
func (x gen_NSWorkspace) ActivateFileViewerSelectingURLs(
	fileURLs core.NSArrayRef,
) {
	C.NSWorkspace_inst_ActivateFileViewerSelectingURLs(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fileURLs),
	)

	return
}

// DesktopImageOptionsForScreen returns the desktop image options for the given screen.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1530855-desktopimageoptionsforscreen?language=objc for details.
func (x gen_NSWorkspace) DesktopImageOptionsForScreen(
	screen NSScreenRef,
) core.NSDictionary {
	ret := C.NSWorkspace_inst_DesktopImageOptionsForScreen(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(screen),
	)

	return core.NSDictionary_FromPointer(ret)
}

// DesktopImageURLForScreen returns the URL for the desktop image for the given screen.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1530635-desktopimageurlforscreen?language=objc for details.
func (x gen_NSWorkspace) DesktopImageURLForScreen(
	screen NSScreenRef,
) core.NSURL {
	ret := C.NSWorkspace_inst_DesktopImageURLForScreen(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(screen),
	)

	return core.NSURL_FromPointer(ret)
}

// ExtendPowerOffBy requests the system wait for the specified amount of time before turning off the power or logging out the user.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1533106-extendpoweroffby?language=objc for details.
func (x gen_NSWorkspace) ExtendPowerOffBy(
	requested core.NSInteger,
) core.NSInteger {
	ret := C.NSWorkspace_inst_ExtendPowerOffBy(
		unsafe.Pointer(x.Pointer()),
		C.long(requested),
	)

	return core.NSInteger(ret)
}

// HideOtherApplications hides all applications other than the sender.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1530417-hideotherapplications?language=objc for details.
func (x gen_NSWorkspace) HideOtherApplications() {
	C.NSWorkspace_inst_HideOtherApplications(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// IconForFile returns an image containing the icon for the specified file.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1528158-iconforfile?language=objc for details.
func (x gen_NSWorkspace) IconForFile(
	fullPath core.NSStringRef,
) NSImage {
	ret := C.NSWorkspace_inst_IconForFile(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fullPath),
	)

	return NSImage_FromPointer(ret)
}

// IconForFiles returns an image containing the icon for the specified files.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1525487-iconforfiles?language=objc for details.
func (x gen_NSWorkspace) IconForFiles(
	fullPaths core.NSArrayRef,
) NSImage {
	ret := C.NSWorkspace_inst_IconForFiles(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fullPaths),
	)

	return NSImage_FromPointer(ret)
}

// IsFilePackageAtPath determines whether the specified path is a file package.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1529991-isfilepackageatpath?language=objc for details.
func (x gen_NSWorkspace) IsFilePackageAtPath(
	fullPath core.NSStringRef,
) bool {
	ret := C.NSWorkspace_inst_IsFilePackageAtPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fullPath),
	)

	return convertObjCBoolToGo(ret)
}

// NoteFileSystemChanged informs the workspace object that the file system changed at the specified path.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1525376-notefilesystemchanged?language=objc for details.
func (x gen_NSWorkspace) NoteFileSystemChanged(
	path core.NSStringRef,
) {
	C.NSWorkspace_inst_NoteFileSystemChanged(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)

	return
}

// OpenURL opens the location at the specified URL.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1533463-openurl?language=objc for details.
func (x gen_NSWorkspace) OpenURL(
	url core.NSURLRef,
) bool {
	ret := C.NSWorkspace_inst_OpenURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)

	return convertObjCBoolToGo(ret)
}

// SelectFileInFileViewerRootedAtPath selects the file at the specified path.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1524399-selectfile?language=objc for details.
func (x gen_NSWorkspace) SelectFileInFileViewerRootedAtPath(
	fullPath core.NSStringRef,
	rootFullPath core.NSStringRef,
) bool {
	ret := C.NSWorkspace_inst_SelectFileInFileViewerRootedAtPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fullPath),
		objc.RefPointer(rootFullPath),
	)

	return convertObjCBoolToGo(ret)
}

// ShowSearchResultsForQueryString displays a Spotlight search results window in Finder for the specified query string.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1532131-showsearchresultsforquerystring?language=objc for details.
func (x gen_NSWorkspace) ShowSearchResultsForQueryString(
	queryString core.NSStringRef,
) bool {
	ret := C.NSWorkspace_inst_ShowSearchResultsForQueryString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(queryString),
	)

	return convertObjCBoolToGo(ret)
}

// UnmountAndEjectDeviceAtPath unmounts and ejects the device at the specified path.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1527741-unmountandejectdeviceatpath?language=objc for details.
func (x gen_NSWorkspace) UnmountAndEjectDeviceAtPath(
	path core.NSStringRef,
) bool {
	ret := C.NSWorkspace_inst_UnmountAndEjectDeviceAtPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)

	return convertObjCBoolToGo(ret)
}

// Init is undocumented.
func (x gen_NSWorkspace) Init() NSWorkspace {
	ret := C.NSWorkspace_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSWorkspace_FromPointer(ret)
}

// FrontmostApplication returns the frontmost app, which is the app that receives key events.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1532097-frontmostapplication?language=objc for details.
func (x gen_NSWorkspace) FrontmostApplication() NSRunningApplication {
	ret := C.NSWorkspace_inst_FrontmostApplication(
		unsafe.Pointer(x.Pointer()),
	)

	return NSRunningApplication_FromPointer(ret)
}

// RunningApplications returns an array of running apps.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1534059-runningapplications?language=objc for details.
func (x gen_NSWorkspace) RunningApplications() core.NSArray {
	ret := C.NSWorkspace_inst_RunningApplications(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// MenuBarOwningApplication returns the app that owns the currently displayed menu bar.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1525848-menubarowningapplication?language=objc for details.
func (x gen_NSWorkspace) MenuBarOwningApplication() NSRunningApplication {
	ret := C.NSWorkspace_inst_MenuBarOwningApplication(
		unsafe.Pointer(x.Pointer()),
	)

	return NSRunningApplication_FromPointer(ret)
}

// FileLabels returns the array of file labels, returned as strings.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1533953-filelabels?language=objc for details.
func (x gen_NSWorkspace) FileLabels() core.NSArray {
	ret := C.NSWorkspace_inst_FileLabels(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// FileLabelColors returns the array of colors for the file labels.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1527553-filelabelcolors?language=objc for details.
func (x gen_NSWorkspace) FileLabelColors() core.NSArray {
	ret := C.NSWorkspace_inst_FileLabelColors(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// AccessibilityDisplayShouldDifferentiateWithoutColor returns a Boolean value that indicates whether the app avoids conveying information through color alone.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1524656-accessibilitydisplayshoulddiffer?language=objc for details.
func (x gen_NSWorkspace) AccessibilityDisplayShouldDifferentiateWithoutColor() bool {
	ret := C.NSWorkspace_inst_AccessibilityDisplayShouldDifferentiateWithoutColor(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AccessibilityDisplayShouldIncreaseContrast returns a Boolean value that indicates whether the app presents a high-contrast user interface.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1526290-accessibilitydisplayshouldincrea?language=objc for details.
func (x gen_NSWorkspace) AccessibilityDisplayShouldIncreaseContrast() bool {
	ret := C.NSWorkspace_inst_AccessibilityDisplayShouldIncreaseContrast(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AccessibilityDisplayShouldReduceTransparency returns a Boolean value that indicates whether the app avoids using semitransparent backgrounds.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1533006-accessibilitydisplayshouldreduce?language=objc for details.
func (x gen_NSWorkspace) AccessibilityDisplayShouldReduceTransparency() bool {
	ret := C.NSWorkspace_inst_AccessibilityDisplayShouldReduceTransparency(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AccessibilityDisplayShouldInvertColors returns a Boolean value that indicates whether the accessibility option to invert colors is in an enabled state.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1644068-accessibilitydisplayshouldinvert?language=objc for details.
func (x gen_NSWorkspace) AccessibilityDisplayShouldInvertColors() bool {
	ret := C.NSWorkspace_inst_AccessibilityDisplayShouldInvertColors(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AccessibilityDisplayShouldReduceMotion returns a Boolean value that indicates whether the accessibility option to reduce motion is in an enabled state.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/1644069-accessibilitydisplayshouldreduce?language=objc for details.
func (x gen_NSWorkspace) AccessibilityDisplayShouldReduceMotion() bool {
	ret := C.NSWorkspace_inst_AccessibilityDisplayShouldReduceMotion(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsSwitchControlEnabled returns a Boolean value that indicates whether Switch Control is currently running.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/2880322-switchcontrolenabled?language=objc for details.
func (x gen_NSWorkspace) IsSwitchControlEnabled() bool {
	ret := C.NSWorkspace_inst_IsSwitchControlEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsVoiceOverEnabled returns a Boolean value that indicates whether VoiceOver is currently running.
//
// See https://developer.apple.com/documentation/appkit/nsworkspace/2880317-voiceoverenabled?language=objc for details.
func (x gen_NSWorkspace) IsVoiceOverEnabled() bool {
	ret := C.NSWorkspace_inst_IsVoiceOverEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

type NSColorRef interface {
	Pointer() uintptr
	Init() NSColor
}

type gen_NSColor struct {
	objc.Object
}

func NSColor_FromPointer(ptr unsafe.Pointer) NSColor {
	return NSColor{gen_NSColor{
		objc.Object_FromPointer(ptr),
	}}
}

func NSColor_FromRef(ref objc.Ref) NSColor {
	return NSColor_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// BlendedColorWithFractionOfColor creates a new color object whose component values are a weighted sum of the current color object and the specified color object's.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1524689-blendedcolorwithfraction?language=objc for details.
func (x gen_NSColor) BlendedColorWithFractionOfColor(
	fraction core.CGFloat,
	color NSColorRef,
) NSColor {
	ret := C.NSColor_inst_BlendedColorWithFractionOfColor(
		unsafe.Pointer(x.Pointer()),
		C.double(fraction),
		objc.RefPointer(color),
	)

	return NSColor_FromPointer(ret)
}

// ColorWithAlphaComponent creates a new color object that has the same color space and component values as the current color object, but the specified alpha component.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1526906-colorwithalphacomponent?language=objc for details.
func (x gen_NSColor) ColorWithAlphaComponent(
	alpha core.CGFloat,
) NSColor {
	ret := C.NSColor_inst_ColorWithAlphaComponent(
		unsafe.Pointer(x.Pointer()),
		C.double(alpha),
	)

	return NSColor_FromPointer(ret)
}

// DrawSwatchInRect draws the current color in the specified rectangle.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1531770-drawswatchinrect?language=objc for details.
func (x gen_NSColor) DrawSwatchInRect(
	rect core.NSRect,
) {
	C.NSColor_inst_DrawSwatchInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return
}

// HighlightWithLevel creates a new color object that represents a blend between the current color and the highlight color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1533061-highlightwithlevel?language=objc for details.
func (x gen_NSColor) HighlightWithLevel(
	val core.CGFloat,
) NSColor {
	ret := C.NSColor_inst_HighlightWithLevel(
		unsafe.Pointer(x.Pointer()),
		C.double(val),
	)

	return NSColor_FromPointer(ret)
}

// Set sets the color of subsequent drawing to the color that the color object represents.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1527089-set?language=objc for details.
func (x gen_NSColor) Set() {
	C.NSColor_inst_Set(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// SetFill sets the fill color of subsequent drawing to the color object’s color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1524755-setfill?language=objc for details.
func (x gen_NSColor) SetFill() {
	C.NSColor_inst_SetFill(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// SetStroke sets the stroke color of subsequent drawing to the color object’s color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1531019-setstroke?language=objc for details.
func (x gen_NSColor) SetStroke() {
	C.NSColor_inst_SetStroke(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ShadowWithLevel creates a new color object that represents a blend between the current color and the shadow color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1528523-shadowwithlevel?language=objc for details.
func (x gen_NSColor) ShadowWithLevel(
	val core.CGFloat,
) NSColor {
	ret := C.NSColor_inst_ShadowWithLevel(
		unsafe.Pointer(x.Pointer()),
		C.double(val),
	)

	return NSColor_FromPointer(ret)
}

// WriteToPasteboard writes the color object’s data to the specified pasteboard.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1532199-writetopasteboard?language=objc for details.
func (x gen_NSColor) WriteToPasteboard(
	pasteBoard NSPasteboardRef,
) {
	C.NSColor_inst_WriteToPasteboard(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pasteBoard),
	)

	return
}

// Init is undocumented.
func (x gen_NSColor) Init() NSColor {
	ret := C.NSColor_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_FromPointer(ret)
}

// NumberOfComponents returns the number of components in the color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1531308-numberofcomponents?language=objc for details.
func (x gen_NSColor) NumberOfComponents() core.NSInteger {
	ret := C.NSColor_inst_NumberOfComponents(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// AlphaComponent returns the alpha (opacity) component value of the color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1532504-alphacomponent?language=objc for details.
func (x gen_NSColor) AlphaComponent() core.CGFloat {
	ret := C.NSColor_inst_AlphaComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// WhiteComponent returns the white component value of the color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1534051-whitecomponent?language=objc for details.
func (x gen_NSColor) WhiteComponent() core.CGFloat {
	ret := C.NSColor_inst_WhiteComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// RedComponent returns the red component value of the color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1530483-redcomponent?language=objc for details.
func (x gen_NSColor) RedComponent() core.CGFloat {
	ret := C.NSColor_inst_RedComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// GreenComponent returns the green component value of the color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1525935-greencomponent?language=objc for details.
func (x gen_NSColor) GreenComponent() core.CGFloat {
	ret := C.NSColor_inst_GreenComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// BlueComponent returns the blue component value of the color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1534229-bluecomponent?language=objc for details.
func (x gen_NSColor) BlueComponent() core.CGFloat {
	ret := C.NSColor_inst_BlueComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// CyanComponent returns the cyan component value of the color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1528234-cyancomponent?language=objc for details.
func (x gen_NSColor) CyanComponent() core.CGFloat {
	ret := C.NSColor_inst_CyanComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// MagentaComponent returns the magenta component value of the color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1535560-magentacomponent?language=objc for details.
func (x gen_NSColor) MagentaComponent() core.CGFloat {
	ret := C.NSColor_inst_MagentaComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// YellowComponent returns the yellow component value of the color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1531965-yellowcomponent?language=objc for details.
func (x gen_NSColor) YellowComponent() core.CGFloat {
	ret := C.NSColor_inst_YellowComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// BlackComponent returns the black component value of the color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1526883-blackcomponent?language=objc for details.
func (x gen_NSColor) BlackComponent() core.CGFloat {
	ret := C.NSColor_inst_BlackComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// HueComponent returns the hue component value of the color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1531780-huecomponent?language=objc for details.
func (x gen_NSColor) HueComponent() core.CGFloat {
	ret := C.NSColor_inst_HueComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// SaturationComponent returns the saturation component value of the color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1526326-saturationcomponent?language=objc for details.
func (x gen_NSColor) SaturationComponent() core.CGFloat {
	ret := C.NSColor_inst_SaturationComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// BrightnessComponent returns the brightness component value of the color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1529355-brightnesscomponent?language=objc for details.
func (x gen_NSColor) BrightnessComponent() core.CGFloat {
	ret := C.NSColor_inst_BrightnessComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// LocalizedCatalogNameComponent returns the localized version of the catalog name containing the color.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1535351-localizedcatalognamecomponent?language=objc for details.
func (x gen_NSColor) LocalizedCatalogNameComponent() core.NSString {
	ret := C.NSColor_inst_LocalizedCatalogNameComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// LocalizedColorNameComponent returns the localized version of the color name.
//
// See https://developer.apple.com/documentation/appkit/nscolor/1527286-localizedcolornamecomponent?language=objc for details.
func (x gen_NSColor) LocalizedColorNameComponent() core.NSString {
	ret := C.NSColor_inst_LocalizedColorNameComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

type NSTextViewRef interface {
	Pointer() uintptr
	Init() NSTextView
}

type gen_NSTextView struct {
	NSText
}

func NSTextView_FromPointer(ptr unsafe.Pointer) NSTextView {
	return NSTextView{gen_NSTextView{
		NSText_FromPointer(ptr),
	}}
}

func NSTextView_FromRef(ref objc.Ref) NSTextView {
	return NSTextView_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// AlignJustified applies full justification to selected paragraphs (or all text, if the receiver is a plain text object).
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449515-alignjustified?language=objc for details.
func (x gen_NSTextView) AlignJustified(
	sender objc.Ref,
) {
	C.NSTextView_inst_AlignJustified(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// BreakUndoCoalescing informs the receiver that it should begin coalescing successive typing operations in a new undo grouping.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449384-breakundocoalescing?language=objc for details.
func (x gen_NSTextView) BreakUndoCoalescing() {
	C.NSTextView_inst_BreakUndoCoalescing(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ChangeAttributes changes the attributes of the current selection.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449216-changeattributes?language=objc for details.
func (x gen_NSTextView) ChangeAttributes(
	sender objc.Ref,
) {
	C.NSTextView_inst_ChangeAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ChangeColor sets the color of the selected text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449282-changecolor?language=objc for details.
func (x gen_NSTextView) ChangeColor(
	sender objc.Ref,
) {
	C.NSTextView_inst_ChangeColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ChangeDocumentBackgroundColor an action method used to set the background color.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449475-changedocumentbackgroundcolor?language=objc for details.
func (x gen_NSTextView) ChangeDocumentBackgroundColor(
	sender objc.Ref,
) {
	C.NSTextView_inst_ChangeDocumentBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ChangeLayoutOrientation an action method that sets the layout orientation of the text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449286-changelayoutorientation?language=objc for details.
func (x gen_NSTextView) ChangeLayoutOrientation(
	sender objc.Ref,
) {
	C.NSTextView_inst_ChangeLayoutOrientation(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// CharacterIndexForInsertionAtPoint returns a character index appropriate for placing a zero-length selection for an insertion point associated with the mouse at the given point.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449505-characterindexforinsertionatpoin?language=objc for details.
func (x gen_NSTextView) CharacterIndexForInsertionAtPoint(
	point core.NSPoint,
) core.NSUInteger {
	ret := C.NSTextView_inst_CharacterIndexForInsertionAtPoint(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return core.NSUInteger(ret)
}

// CheckTextInDocument performs the default text checking on the entire document.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449440-checktextindocument?language=objc for details.
func (x gen_NSTextView) CheckTextInDocument(
	sender objc.Ref,
) {
	C.NSTextView_inst_CheckTextInDocument(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// CheckTextInSelection performs the default text checking on the current selection.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449382-checktextinselection?language=objc for details.
func (x gen_NSTextView) CheckTextInSelection(
	sender objc.Ref,
) {
	C.NSTextView_inst_CheckTextInSelection(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// CleanUpAfterDragOperation releases the drag information still existing after the dragging session has completed.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449202-cleanupafterdragoperation?language=objc for details.
func (x gen_NSTextView) CleanUpAfterDragOperation() {
	C.NSTextView_inst_CleanUpAfterDragOperation(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ClickedOnLinkAtIndex causes the text view to act as if the user clicked on some text with the given link as the value of a link attribute associated with the text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449497-clickedonlink?language=objc for details.
func (x gen_NSTextView) ClickedOnLinkAtIndex(
	link objc.Ref,
	charIndex core.NSUInteger,
) {
	C.NSTextView_inst_ClickedOnLinkAtIndex(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(link),
		C.ulong(charIndex),
	)

	return
}

// Complete invokes completion in a text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449359-complete?language=objc for details.
func (x gen_NSTextView) Complete(
	sender objc.Ref,
) {
	C.NSTextView_inst_Complete(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// DidChangeText sends out necessary notifications when a text change completes.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449296-didchangetext?language=objc for details.
func (x gen_NSTextView) DidChangeText() {
	C.NSTextView_inst_DidChangeText(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DragSelectionWithEventOffsetSlideBack begins dragging the current selected text range.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449413-dragselectionwithevent?language=objc for details.
func (x gen_NSTextView) DragSelectionWithEventOffsetSlideBack(
	event NSEventRef,
	mouseOffset core.NSSize,
	slideBack bool,
) bool {
	ret := C.NSTextView_inst_DragSelectionWithEventOffsetSlideBack(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
		*(*C.NSSize)(unsafe.Pointer(&mouseOffset)),
		convertToObjCBool(slideBack),
	)

	return convertObjCBoolToGo(ret)
}

// DrawInsertionPointInRectColorTurnedOn draws or erases the insertion point.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449232-drawinsertionpointinrect?language=objc for details.
func (x gen_NSTextView) DrawInsertionPointInRectColorTurnedOn(
	rect core.NSRect,
	color NSColorRef,
	flag bool,
) {
	C.NSTextView_inst_DrawInsertionPointInRectColorTurnedOn(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(color),
		convertToObjCBool(flag),
	)

	return
}

// DrawViewBackgroundInRect draws the background of the text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449135-drawviewbackgroundinrect?language=objc for details.
func (x gen_NSTextView) DrawViewBackgroundInRect(
	rect core.NSRect,
) {
	C.NSTextView_inst_DrawViewBackgroundInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return
}

// InitWithFrame initializes a text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449262-initwithframe?language=objc for details.
func (x gen_NSTextView) InitWithFrame(
	frameRect core.NSRect,
) NSTextView {
	ret := C.NSTextView_inst_InitWithFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)

	return NSTextView_FromPointer(ret)
}

// InitWithFrameTextContainer initializes a text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449347-initwithframe?language=objc for details.
func (x gen_NSTextView) InitWithFrameTextContainer(
	frameRect core.NSRect,
	container NSTextContainerRef,
) NSTextView {
	ret := C.NSTextView_inst_InitWithFrameTextContainer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
		objc.RefPointer(container),
	)

	return NSTextView_FromPointer(ret)
}

// InvalidateTextContainerOrigin invalidates the calculated origin of the text container.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449546-invalidatetextcontainerorigin?language=objc for details.
func (x gen_NSTextView) InvalidateTextContainerOrigin() {
	C.NSTextView_inst_InvalidateTextContainerOrigin(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// LoosenKerning increases the space between glyphs in the receiver’s selection, or in all text if the receiver is a plain text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449183-loosenkerning?language=objc for details.
func (x gen_NSTextView) LoosenKerning(
	sender objc.Ref,
) {
	C.NSTextView_inst_LoosenKerning(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// LowerBaseline lowers the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449289-lowerbaseline?language=objc for details.
func (x gen_NSTextView) LowerBaseline(
	sender objc.Ref,
) {
	C.NSTextView_inst_LowerBaseline(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// OrderFrontLinkPanel brings forward a panel allowing the user to manipulate links in the text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449238-orderfrontlinkpanel?language=objc for details.
func (x gen_NSTextView) OrderFrontLinkPanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_OrderFrontLinkPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// OrderFrontListPanel brings forward a panel allowing the user to manipulate text lists in the text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449349-orderfrontlistpanel?language=objc for details.
func (x gen_NSTextView) OrderFrontListPanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_OrderFrontListPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// OrderFrontSharingServicePicker creates and displays a new instance of the sharing service picker.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449150-orderfrontsharingservicepicker?language=objc for details.
func (x gen_NSTextView) OrderFrontSharingServicePicker(
	sender objc.Ref,
) {
	C.NSTextView_inst_OrderFrontSharingServicePicker(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// OrderFrontSpacingPanel brings forward a panel allowing the user to manipulate text line heights, interline spacing, and paragraph spacing, in the text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449438-orderfrontspacingpanel?language=objc for details.
func (x gen_NSTextView) OrderFrontSpacingPanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_OrderFrontSpacingPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// OrderFrontSubstitutionsPanel brings forward a panel allowing the user to specify string substitutions in the text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449327-orderfrontsubstitutionspanel?language=objc for details.
func (x gen_NSTextView) OrderFrontSubstitutionsPanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_OrderFrontSubstitutionsPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// OrderFrontTablePanel brings forward a panel allowing the user to manipulate text tables in the text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449442-orderfronttablepanel?language=objc for details.
func (x gen_NSTextView) OrderFrontTablePanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_OrderFrontTablePanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// Outline adds the outline attribute to the selected text attributes if absent; removes the attribute if present.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449386-outline?language=objc for details.
func (x gen_NSTextView) Outline(
	sender objc.Ref,
) {
	C.NSTextView_inst_Outline(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// PasteAsPlainText inserts the contents of the pasteboard into the receiver’s text as plain text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449250-pasteasplaintext?language=objc for details.
func (x gen_NSTextView) PasteAsPlainText(
	sender objc.Ref,
) {
	C.NSTextView_inst_PasteAsPlainText(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// PasteAsRichText this action method inserts the contents of the pasteboard into the receiver’s text as rich text, maintaining its attributes.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449395-pasteasrichtext?language=objc for details.
func (x gen_NSTextView) PasteAsRichText(
	sender objc.Ref,
) {
	C.NSTextView_inst_PasteAsRichText(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// PerformFindPanelAction performs a find panel action specified by the sender's tag.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449525-performfindpanelaction?language=objc for details.
func (x gen_NSTextView) PerformFindPanelAction(
	sender objc.Ref,
) {
	C.NSTextView_inst_PerformFindPanelAction(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// QuickLookPreviewableItemsInRanges returns an array of URLs for items that can be displayed by QuickLook in the specified ranges.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449426-quicklookpreviewableitemsinrange?language=objc for details.
func (x gen_NSTextView) QuickLookPreviewableItemsInRanges(
	ranges core.NSArrayRef,
) core.NSArray {
	ret := C.NSTextView_inst_QuickLookPreviewableItemsInRanges(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ranges),
	)

	return core.NSArray_FromPointer(ret)
}

// RaiseBaseline raises the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449198-raisebaseline?language=objc for details.
func (x gen_NSTextView) RaiseBaseline(
	sender objc.Ref,
) {
	C.NSTextView_inst_RaiseBaseline(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ReadSelectionFromPasteboard reads the text view’s preferred type of data from the specified pasteboard.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449469-readselectionfrompasteboard?language=objc for details.
func (x gen_NSTextView) ReadSelectionFromPasteboard(
	pboard NSPasteboardRef,
) bool {
	ret := C.NSTextView_inst_ReadSelectionFromPasteboard(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pboard),
	)

	return convertObjCBoolToGo(ret)
}

// ReplaceTextContainer replaces the text container for the group of text system objects containing the receiver, keeping the association between the receiver and its layout manager intact.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449479-replacetextcontainer?language=objc for details.
func (x gen_NSTextView) ReplaceTextContainer(
	newContainer NSTextContainerRef,
) {
	C.NSTextView_inst_ReplaceTextContainer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newContainer),
	)

	return
}

// SetConstrainedFrameSize attempts to set the frame size as if by user action.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449230-setconstrainedframesize?language=objc for details.
func (x gen_NSTextView) SetConstrainedFrameSize(
	desiredSize core.NSSize,
) {
	C.NSTextView_inst_SetConstrainedFrameSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&desiredSize)),
	)

	return
}

// SetNeedsDisplayInRectAvoidAdditionalLayout marks the receiver as requiring display.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449279-setneedsdisplayinrect?language=objc for details.
func (x gen_NSTextView) SetNeedsDisplayInRectAvoidAdditionalLayout(
	rect core.NSRect,
	flag bool,
) {
	C.NSTextView_inst_SetNeedsDisplayInRectAvoidAdditionalLayout(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		convertToObjCBool(flag),
	)

	return
}

// ShouldChangeTextInRangesReplacementStrings initiates a series of delegate messages (and general notifications) to determine whether modifications can be made to the characters and attributes of the receiver’s text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449311-shouldchangetextinranges?language=objc for details.
func (x gen_NSTextView) ShouldChangeTextInRangesReplacementStrings(
	affectedRanges core.NSArrayRef,
	replacementStrings core.NSArrayRef,
) bool {
	ret := C.NSTextView_inst_ShouldChangeTextInRangesReplacementStrings(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(affectedRanges),
		objc.RefPointer(replacementStrings),
	)

	return convertObjCBoolToGo(ret)
}

// StartSpeaking speaks the selected text, or all text if no selection.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449519-startspeaking?language=objc for details.
func (x gen_NSTextView) StartSpeaking(
	sender objc.Ref,
) {
	C.NSTextView_inst_StartSpeaking(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// StopSpeaking stops the speaking of text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449172-stopspeaking?language=objc for details.
func (x gen_NSTextView) StopSpeaking(
	sender objc.Ref,
) {
	C.NSTextView_inst_StopSpeaking(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// TightenKerning decreases the space between glyphs in the receiver’s selection, or for all glyphs if the receiver is a plain text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449137-tightenkerning?language=objc for details.
func (x gen_NSTextView) TightenKerning(
	sender objc.Ref,
) {
	C.NSTextView_inst_TightenKerning(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleAutomaticDashSubstitution toggles the state of the automatic dash substitution.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449305-toggleautomaticdashsubstitution?language=objc for details.
func (x gen_NSTextView) ToggleAutomaticDashSubstitution(
	sender objc.Ref,
) {
	C.NSTextView_inst_ToggleAutomaticDashSubstitution(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleAutomaticDataDetection toggles the state of the automatic data detection.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449499-toggleautomaticdatadetection?language=objc for details.
func (x gen_NSTextView) ToggleAutomaticDataDetection(
	sender objc.Ref,
) {
	C.NSTextView_inst_ToggleAutomaticDataDetection(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleAutomaticLinkDetection changes the state of automatic link detection from enabled to disabled and vice versa.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449353-toggleautomaticlinkdetection?language=objc for details.
func (x gen_NSTextView) ToggleAutomaticLinkDetection(
	sender objc.Ref,
) {
	C.NSTextView_inst_ToggleAutomaticLinkDetection(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleAutomaticQuoteSubstitution changes the state of automatic quotation mark substitution from enabled to disabled and vice versa.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449444-toggleautomaticquotesubstitution?language=objc for details.
func (x gen_NSTextView) ToggleAutomaticQuoteSubstitution(
	sender objc.Ref,
) {
	C.NSTextView_inst_ToggleAutomaticQuoteSubstitution(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleAutomaticSpellingCorrection toggles the state of the automatic spelling correction.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449178-toggleautomaticspellingcorrectio?language=objc for details.
func (x gen_NSTextView) ToggleAutomaticSpellingCorrection(
	sender objc.Ref,
) {
	C.NSTextView_inst_ToggleAutomaticSpellingCorrection(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleAutomaticTextCompletion is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nstextview/2544841-toggleautomatictextcompletion?language=objc for details.
func (x gen_NSTextView) ToggleAutomaticTextCompletion(
	sender objc.Ref,
) {
	C.NSTextView_inst_ToggleAutomaticTextCompletion(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleAutomaticTextReplacement toggles the state of the automatic text replacement.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449200-toggleautomatictextreplacement?language=objc for details.
func (x gen_NSTextView) ToggleAutomaticTextReplacement(
	sender objc.Ref,
) {
	C.NSTextView_inst_ToggleAutomaticTextReplacement(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleContinuousSpellChecking toggles whether continuous spell checking is enabled for the receiver.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449471-togglecontinuousspellchecking?language=objc for details.
func (x gen_NSTextView) ToggleContinuousSpellChecking(
	sender objc.Ref,
) {
	C.NSTextView_inst_ToggleContinuousSpellChecking(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleGrammarChecking changes the state of grammar checking from enabled to disabled and vice versa.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449393-togglegrammarchecking?language=objc for details.
func (x gen_NSTextView) ToggleGrammarChecking(
	sender objc.Ref,
) {
	C.NSTextView_inst_ToggleGrammarChecking(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleQuickLookPreviewPanel an action message that toggles the visibility state of the Quick Look preview panel.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449415-togglequicklookpreviewpanel?language=objc for details.
func (x gen_NSTextView) ToggleQuickLookPreviewPanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_ToggleQuickLookPreviewPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ToggleSmartInsertDelete changes the state of smart insert and delete from enabled to disabled and vice versa.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449273-togglesmartinsertdelete?language=objc for details.
func (x gen_NSTextView) ToggleSmartInsertDelete(
	sender objc.Ref,
) {
	C.NSTextView_inst_ToggleSmartInsertDelete(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// TurnOffKerning sets the receiver to use nominal glyph spacing for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449464-turnoffkerning?language=objc for details.
func (x gen_NSTextView) TurnOffKerning(
	sender objc.Ref,
) {
	C.NSTextView_inst_TurnOffKerning(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// TurnOffLigatures sets the receiver to use only required ligatures when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449436-turnoffligatures?language=objc for details.
func (x gen_NSTextView) TurnOffLigatures(
	sender objc.Ref,
) {
	C.NSTextView_inst_TurnOffLigatures(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// UpdateCandidates is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nstextview/2544833-updatecandidates?language=objc for details.
func (x gen_NSTextView) UpdateCandidates() {
	C.NSTextView_inst_UpdateCandidates(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// UpdateDragTypeRegistration updates the acceptable drag types of all text views associated with the receiver's layout manager.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449181-updatedragtyperegistration?language=objc for details.
func (x gen_NSTextView) UpdateDragTypeRegistration() {
	C.NSTextView_inst_UpdateDragTypeRegistration(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// UpdateFontPanel updates the Font panel to contain the font attributes of the selection.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449401-updatefontpanel?language=objc for details.
func (x gen_NSTextView) UpdateFontPanel() {
	C.NSTextView_inst_UpdateFontPanel(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// UpdateInsertionPointStateAndRestartTimer updates the insertion point’s location and optionally restarts the blinking cursor timer.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449268-updateinsertionpointstateandrest?language=objc for details.
func (x gen_NSTextView) UpdateInsertionPointStateAndRestartTimer(
	restartFlag bool,
) {
	C.NSTextView_inst_UpdateInsertionPointStateAndRestartTimer(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(restartFlag),
	)

	return
}

// UpdateQuickLookPreviewPanel notifies the QuickLook panel that an update may be required.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449409-updatequicklookpreviewpanel?language=objc for details.
func (x gen_NSTextView) UpdateQuickLookPreviewPanel() {
	C.NSTextView_inst_UpdateQuickLookPreviewPanel(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// UpdateRuler updates the ruler view in the receiver’s enclosing scroll view to reflect the selection’s paragraph and marker attributes.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449323-updateruler?language=objc for details.
func (x gen_NSTextView) UpdateRuler() {
	C.NSTextView_inst_UpdateRuler(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// UpdateTextTouchBarItems is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nstextview/2544676-updatetexttouchbaritems?language=objc for details.
func (x gen_NSTextView) UpdateTextTouchBarItems() {
	C.NSTextView_inst_UpdateTextTouchBarItems(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// UpdateTouchBarItemIdentifiers is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nstextview/2544834-updatetouchbaritemidentifiers?language=objc for details.
func (x gen_NSTextView) UpdateTouchBarItemIdentifiers() {
	C.NSTextView_inst_UpdateTouchBarItemIdentifiers(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// UseAllLigatures sets the receiver to use all ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449213-useallligatures?language=objc for details.
func (x gen_NSTextView) UseAllLigatures(
	sender objc.Ref,
) {
	C.NSTextView_inst_UseAllLigatures(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// UseStandardKerning set the receiver to use pair kerning data for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449491-usestandardkerning?language=objc for details.
func (x gen_NSTextView) UseStandardKerning(
	sender objc.Ref,
) {
	C.NSTextView_inst_UseStandardKerning(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// UseStandardLigatures sets the receiver to use the standard ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449144-usestandardligatures?language=objc for details.
func (x gen_NSTextView) UseStandardLigatures(
	sender objc.Ref,
) {
	C.NSTextView_inst_UseStandardLigatures(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// WriteSelectionToPasteboardTypes writes the current selection to the specified pasteboard under each given type.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449277-writeselectiontopasteboard?language=objc for details.
func (x gen_NSTextView) WriteSelectionToPasteboardTypes(
	pboard NSPasteboardRef,
	types core.NSArrayRef,
) bool {
	ret := C.NSTextView_inst_WriteSelectionToPasteboardTypes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pboard),
		objc.RefPointer(types),
	)

	return convertObjCBoolToGo(ret)
}

// Init is undocumented.
func (x gen_NSTextView) Init() NSTextView {
	ret := C.NSTextView_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextView_FromPointer(ret)
}

// Delegate returns the delegate for all text views sharing the receiver’s layout manager.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449521-delegate?language=objc for details.
func (x gen_NSTextView) Delegate() objc.Object {
	ret := C.NSTextView_inst_Delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetDelegate returns the delegate for all text views sharing the receiver’s layout manager.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449521-delegate?language=objc for details.
func (x gen_NSTextView) SetDelegate(
	value objc.Ref,
) {
	C.NSTextView_inst_SetDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// TextContainer returns the receiver’s text container.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449364-textcontainer?language=objc for details.
func (x gen_NSTextView) TextContainer() NSTextContainer {
	ret := C.NSTextView_inst_TextContainer(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextContainer_FromPointer(ret)
}

// SetTextContainer returns the receiver’s text container.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449364-textcontainer?language=objc for details.
func (x gen_NSTextView) SetTextContainer(
	value NSTextContainerRef,
) {
	C.NSTextView_inst_SetTextContainer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// TextContainerInset returns the empty space the receiver leaves around its associated text container.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449168-textcontainerinset?language=objc for details.
func (x gen_NSTextView) TextContainerInset() core.NSSize {
	ret := C.NSTextView_inst_TextContainerInset(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// SetTextContainerInset returns the empty space the receiver leaves around its associated text container.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449168-textcontainerinset?language=objc for details.
func (x gen_NSTextView) SetTextContainerInset(
	value core.NSSize,
) {
	C.NSTextView_inst_SetTextContainerInset(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// TextContainerOrigin returns the origin of the receiver’s text container.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449477-textcontainerorigin?language=objc for details.
func (x gen_NSTextView) TextContainerOrigin() core.NSPoint {
	ret := C.NSTextView_inst_TextContainerOrigin(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// LayoutManager returns the layout manager that lays out text for the receiver’s text container.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449148-layoutmanager?language=objc for details.
func (x gen_NSTextView) LayoutManager() NSLayoutManager {
	ret := C.NSTextView_inst_LayoutManager(
		unsafe.Pointer(x.Pointer()),
	)

	return NSLayoutManager_FromPointer(ret)
}

// BackgroundColor returns the receiver’s background color.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449501-backgroundcolor?language=objc for details.
func (x gen_NSTextView) BackgroundColor() NSColor {
	ret := C.NSTextView_inst_BackgroundColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_FromPointer(ret)
}

// SetBackgroundColor returns the receiver’s background color.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449501-backgroundcolor?language=objc for details.
func (x gen_NSTextView) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSTextView_inst_SetBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// DrawsBackground returns a Boolean value that indicates whether the receiver draws its background.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449530-drawsbackground?language=objc for details.
func (x gen_NSTextView) DrawsBackground() bool {
	ret := C.NSTextView_inst_DrawsBackground(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetDrawsBackground returns a Boolean value that indicates whether the receiver draws its background.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449530-drawsbackground?language=objc for details.
func (x gen_NSTextView) SetDrawsBackground(
	value bool,
) {
	C.NSTextView_inst_SetDrawsBackground(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AllowsDocumentBackgroundColorChange returns a Boolean value that indicates whether the receiver allows its background color to change.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449397-allowsdocumentbackgroundcolorcha?language=objc for details.
func (x gen_NSTextView) AllowsDocumentBackgroundColorChange() bool {
	ret := C.NSTextView_inst_AllowsDocumentBackgroundColorChange(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsDocumentBackgroundColorChange returns a Boolean value that indicates whether the receiver allows its background color to change.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449397-allowsdocumentbackgroundcolorcha?language=objc for details.
func (x gen_NSTextView) SetAllowsDocumentBackgroundColorChange(
	value bool,
) {
	C.NSTextView_inst_SetAllowsDocumentBackgroundColorChange(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// ShouldDrawInsertionPoint returns a Boolean value that determines whether the receiver should draw its insertion point.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449152-shoulddrawinsertionpoint?language=objc for details.
func (x gen_NSTextView) ShouldDrawInsertionPoint() bool {
	ret := C.NSTextView_inst_ShouldDrawInsertionPoint(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AllowedInputSourceLocales an array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449370-allowedinputsourcelocales?language=objc for details.
func (x gen_NSTextView) AllowedInputSourceLocales() core.NSArray {
	ret := C.NSTextView_inst_AllowedInputSourceLocales(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// SetAllowedInputSourceLocales an array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449370-allowedinputsourcelocales?language=objc for details.
func (x gen_NSTextView) SetAllowedInputSourceLocales(
	value core.NSArrayRef,
) {
	C.NSTextView_inst_SetAllowedInputSourceLocales(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// AllowsUndo returns a Boolean value that indicates whether the receiver allows undo.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449450-allowsundo?language=objc for details.
func (x gen_NSTextView) AllowsUndo() bool {
	ret := C.NSTextView_inst_AllowsUndo(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsUndo returns a Boolean value that indicates whether the receiver allows undo.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449450-allowsundo?language=objc for details.
func (x gen_NSTextView) SetAllowsUndo(
	value bool,
) {
	C.NSTextView_inst_SetAllowsUndo(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsEditable returns a Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449345-editable?language=objc for details.
func (x gen_NSTextView) IsEditable() bool {
	ret := C.NSTextView_inst_IsEditable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetEditable returns a Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449345-editable?language=objc for details.
func (x gen_NSTextView) SetEditable(
	value bool,
) {
	C.NSTextView_inst_SetEditable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsSelectable returns a Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449297-selectable?language=objc for details.
func (x gen_NSTextView) IsSelectable() bool {
	ret := C.NSTextView_inst_IsSelectable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetSelectable returns a Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449297-selectable?language=objc for details.
func (x gen_NSTextView) SetSelectable(
	value bool,
) {
	C.NSTextView_inst_SetSelectable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsFieldEditor returns a Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449156-fieldeditor?language=objc for details.
func (x gen_NSTextView) IsFieldEditor() bool {
	ret := C.NSTextView_inst_IsFieldEditor(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetFieldEditor returns a Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449156-fieldeditor?language=objc for details.
func (x gen_NSTextView) SetFieldEditor(
	value bool,
) {
	C.NSTextView_inst_SetFieldEditor(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsRichText returns a Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449538-richtext?language=objc for details.
func (x gen_NSTextView) IsRichText() bool {
	ret := C.NSTextView_inst_IsRichText(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetRichText returns a Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449538-richtext?language=objc for details.
func (x gen_NSTextView) SetRichText(
	value bool,
) {
	C.NSTextView_inst_SetRichText(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// ImportsGraphics returns a Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to import files by dragging.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449266-importsgraphics?language=objc for details.
func (x gen_NSTextView) ImportsGraphics() bool {
	ret := C.NSTextView_inst_ImportsGraphics(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetImportsGraphics returns a Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to import files by dragging.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449266-importsgraphics?language=objc for details.
func (x gen_NSTextView) SetImportsGraphics(
	value bool,
) {
	C.NSTextView_inst_SetImportsGraphics(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AllowsImageEditing indicates whether image attachments should permit editing of their images.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449425-allowsimageediting?language=objc for details.
func (x gen_NSTextView) AllowsImageEditing() bool {
	ret := C.NSTextView_inst_AllowsImageEditing(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsImageEditing indicates whether image attachments should permit editing of their images.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449425-allowsimageediting?language=objc for details.
func (x gen_NSTextView) SetAllowsImageEditing(
	value bool,
) {
	C.NSTextView_inst_SetAllowsImageEditing(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsAutomaticQuoteSubstitutionEnabled returns a Boolean value that enables and disables automatic quotation mark substitution.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449258-automaticquotesubstitutionenable?language=objc for details.
func (x gen_NSTextView) IsAutomaticQuoteSubstitutionEnabled() bool {
	ret := C.NSTextView_inst_IsAutomaticQuoteSubstitutionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAutomaticQuoteSubstitutionEnabled returns a Boolean value that enables and disables automatic quotation mark substitution.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449258-automaticquotesubstitutionenable?language=objc for details.
func (x gen_NSTextView) SetAutomaticQuoteSubstitutionEnabled(
	value bool,
) {
	C.NSTextView_inst_SetAutomaticQuoteSubstitutionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsAutomaticLinkDetectionEnabled returns a Boolean value that enables or disables automatic link detection.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449170-automaticlinkdetectionenabled?language=objc for details.
func (x gen_NSTextView) IsAutomaticLinkDetectionEnabled() bool {
	ret := C.NSTextView_inst_IsAutomaticLinkDetectionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAutomaticLinkDetectionEnabled returns a Boolean value that enables or disables automatic link detection.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449170-automaticlinkdetectionenabled?language=objc for details.
func (x gen_NSTextView) SetAutomaticLinkDetectionEnabled(
	value bool,
) {
	C.NSTextView_inst_SetAutomaticLinkDetectionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// DisplaysLinkToolTips returns a Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449204-displayslinktooltips?language=objc for details.
func (x gen_NSTextView) DisplaysLinkToolTips() bool {
	ret := C.NSTextView_inst_DisplaysLinkToolTips(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetDisplaysLinkToolTips returns a Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449204-displayslinktooltips?language=objc for details.
func (x gen_NSTextView) SetDisplaysLinkToolTips(
	value bool,
) {
	C.NSTextView_inst_SetDisplaysLinkToolTips(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsAutomaticTextCompletionEnabled returns a Boolean vlaue that inidcates whether the text view supplies autocompletion suggestions as the user types.
//
// See https://developer.apple.com/documentation/appkit/nstextview/2544655-automatictextcompletionenabled?language=objc for details.
func (x gen_NSTextView) IsAutomaticTextCompletionEnabled() bool {
	ret := C.NSTextView_inst_IsAutomaticTextCompletionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAutomaticTextCompletionEnabled returns a Boolean vlaue that inidcates whether the text view supplies autocompletion suggestions as the user types.
//
// See https://developer.apple.com/documentation/appkit/nstextview/2544655-automatictextcompletionenabled?language=objc for details.
func (x gen_NSTextView) SetAutomaticTextCompletionEnabled(
	value bool,
) {
	C.NSTextView_inst_SetAutomaticTextCompletionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// UsesAdaptiveColorMappingForDarkAppearance returns a Boolean vlaues that indicates whether the framework should use adaptive color mapping for dark appearance.
//
// See https://developer.apple.com/documentation/appkit/nstextview/3237223-usesadaptivecolormappingfordarka?language=objc for details.
func (x gen_NSTextView) UsesAdaptiveColorMappingForDarkAppearance() bool {
	ret := C.NSTextView_inst_UsesAdaptiveColorMappingForDarkAppearance(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetUsesAdaptiveColorMappingForDarkAppearance returns a Boolean vlaues that indicates whether the framework should use adaptive color mapping for dark appearance.
//
// See https://developer.apple.com/documentation/appkit/nstextview/3237223-usesadaptivecolormappingfordarka?language=objc for details.
func (x gen_NSTextView) SetUsesAdaptiveColorMappingForDarkAppearance(
	value bool,
) {
	C.NSTextView_inst_SetUsesAdaptiveColorMappingForDarkAppearance(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// UsesRolloverButtonForSelection is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449357-usesrolloverbuttonforselection?language=objc for details.
func (x gen_NSTextView) UsesRolloverButtonForSelection() bool {
	ret := C.NSTextView_inst_UsesRolloverButtonForSelection(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetUsesRolloverButtonForSelection is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449357-usesrolloverbuttonforselection?language=objc for details.
func (x gen_NSTextView) SetUsesRolloverButtonForSelection(
	value bool,
) {
	C.NSTextView_inst_SetUsesRolloverButtonForSelection(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// UsesRuler returns a Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449218-usesruler?language=objc for details.
func (x gen_NSTextView) UsesRuler() bool {
	ret := C.NSTextView_inst_UsesRuler(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetUsesRuler returns a Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449218-usesruler?language=objc for details.
func (x gen_NSTextView) SetUsesRuler(
	value bool,
) {
	C.NSTextView_inst_SetUsesRuler(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsRulerVisible returns a Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449406-rulervisible?language=objc for details.
func (x gen_NSTextView) IsRulerVisible() bool {
	ret := C.NSTextView_inst_IsRulerVisible(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetRulerVisible returns a Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449406-rulervisible?language=objc for details.
func (x gen_NSTextView) SetRulerVisible(
	value bool,
) {
	C.NSTextView_inst_SetRulerVisible(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// UsesInspectorBar returns a Boolean value that indicates whether this text view uses the inspector bar.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449407-usesinspectorbar?language=objc for details.
func (x gen_NSTextView) UsesInspectorBar() bool {
	ret := C.NSTextView_inst_UsesInspectorBar(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetUsesInspectorBar returns a Boolean value that indicates whether this text view uses the inspector bar.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449407-usesinspectorbar?language=objc for details.
func (x gen_NSTextView) SetUsesInspectorBar(
	value bool,
) {
	C.NSTextView_inst_SetUsesInspectorBar(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// SelectedRanges an array containing the ranges of characters selected in the receiver’s layout manager.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449129-selectedranges?language=objc for details.
func (x gen_NSTextView) SelectedRanges() core.NSArray {
	ret := C.NSTextView_inst_SelectedRanges(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// SetSelectedRanges an array containing the ranges of characters selected in the receiver’s layout manager.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449129-selectedranges?language=objc for details.
func (x gen_NSTextView) SetSelectedRanges(
	value core.NSArrayRef,
) {
	C.NSTextView_inst_SetSelectedRanges(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// InsertionPointColor returns the color of the insertion point.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449309-insertionpointcolor?language=objc for details.
func (x gen_NSTextView) InsertionPointColor() NSColor {
	ret := C.NSTextView_inst_InsertionPointColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_FromPointer(ret)
}

// SetInsertionPointColor returns the color of the insertion point.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449309-insertionpointcolor?language=objc for details.
func (x gen_NSTextView) SetInsertionPointColor(
	value NSColorRef,
) {
	C.NSTextView_inst_SetInsertionPointColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// SelectedTextAttributes returns the attributes used to indicate the selection.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449270-selectedtextattributes?language=objc for details.
func (x gen_NSTextView) SelectedTextAttributes() core.NSDictionary {
	ret := C.NSTextView_inst_SelectedTextAttributes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_FromPointer(ret)
}

// SetSelectedTextAttributes returns the attributes used to indicate the selection.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449270-selectedtextattributes?language=objc for details.
func (x gen_NSTextView) SetSelectedTextAttributes(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_SetSelectedTextAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// MarkedTextAttributes returns the attributes used to draw marked text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449179-markedtextattributes?language=objc for details.
func (x gen_NSTextView) MarkedTextAttributes() core.NSDictionary {
	ret := C.NSTextView_inst_MarkedTextAttributes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_FromPointer(ret)
}

// SetMarkedTextAttributes returns the attributes used to draw marked text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449179-markedtextattributes?language=objc for details.
func (x gen_NSTextView) SetMarkedTextAttributes(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_SetMarkedTextAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// LinkTextAttributes returns the attributes used to draw the onscreen presentation of link text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449452-linktextattributes?language=objc for details.
func (x gen_NSTextView) LinkTextAttributes() core.NSDictionary {
	ret := C.NSTextView_inst_LinkTextAttributes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_FromPointer(ret)
}

// SetLinkTextAttributes returns the attributes used to draw the onscreen presentation of link text.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449452-linktextattributes?language=objc for details.
func (x gen_NSTextView) SetLinkTextAttributes(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_SetLinkTextAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// ReadablePasteboardTypes returns the types this text view can read immediately from the pasteboard.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449361-readablepasteboardtypes?language=objc for details.
func (x gen_NSTextView) ReadablePasteboardTypes() core.NSArray {
	ret := C.NSTextView_inst_ReadablePasteboardTypes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// WritablePasteboardTypes returns the pasteboard types that can be provided from the current selection.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449222-writablepasteboardtypes?language=objc for details.
func (x gen_NSTextView) WritablePasteboardTypes() core.NSArray {
	ret := C.NSTextView_inst_WritablePasteboardTypes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// TypingAttributes returns the receiver’s typing attributes.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449487-typingattributes?language=objc for details.
func (x gen_NSTextView) TypingAttributes() core.NSDictionary {
	ret := C.NSTextView_inst_TypingAttributes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_FromPointer(ret)
}

// SetTypingAttributes returns the receiver’s typing attributes.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449487-typingattributes?language=objc for details.
func (x gen_NSTextView) SetTypingAttributes(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_SetTypingAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// IsCoalescingUndo returns a Boolean value that indicates whether undo coalescing is in progress.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449368-coalescingundo?language=objc for details.
func (x gen_NSTextView) IsCoalescingUndo() bool {
	ret := C.NSTextView_inst_IsCoalescingUndo(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AcceptableDragTypes returns the data types that the receiver accepts as the destination view of a dragging operation.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449234-acceptabledragtypes?language=objc for details.
func (x gen_NSTextView) AcceptableDragTypes() core.NSArray {
	ret := C.NSTextView_inst_AcceptableDragTypes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// RangesForUserCharacterAttributeChange an array containing the ranges of characters affected by an action method that changes character (not paragraph) attributes.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449503-rangesforusercharacterattributec?language=objc for details.
func (x gen_NSTextView) RangesForUserCharacterAttributeChange() core.NSArray {
	ret := C.NSTextView_inst_RangesForUserCharacterAttributeChange(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// RangesForUserParagraphAttributeChange an array containing the ranges of characters affected by a method that changes paragraph (not character) attributes.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449161-rangesforuserparagraphattributec?language=objc for details.
func (x gen_NSTextView) RangesForUserParagraphAttributeChange() core.NSArray {
	ret := C.NSTextView_inst_RangesForUserParagraphAttributeChange(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// RangesForUserTextChange an array containing the ranges of characters affected by a method that changes characters (as opposed to attributes).
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449434-rangesforusertextchange?language=objc for details.
func (x gen_NSTextView) RangesForUserTextChange() core.NSArray {
	ret := C.NSTextView_inst_RangesForUserTextChange(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// SmartInsertDeleteEnabled returns a Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449236-smartinsertdeleteenabled?language=objc for details.
func (x gen_NSTextView) SmartInsertDeleteEnabled() bool {
	ret := C.NSTextView_inst_SmartInsertDeleteEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetSmartInsertDeleteEnabled returns a Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449236-smartinsertdeleteenabled?language=objc for details.
func (x gen_NSTextView) SetSmartInsertDeleteEnabled(
	value bool,
) {
	C.NSTextView_inst_SetSmartInsertDeleteEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsContinuousSpellCheckingEnabled returns a Boolean value that indicates whether the receiver has continuous spell checking enabled.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449430-continuousspellcheckingenabled?language=objc for details.
func (x gen_NSTextView) IsContinuousSpellCheckingEnabled() bool {
	ret := C.NSTextView_inst_IsContinuousSpellCheckingEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetContinuousSpellCheckingEnabled returns a Boolean value that indicates whether the receiver has continuous spell checking enabled.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449430-continuousspellcheckingenabled?language=objc for details.
func (x gen_NSTextView) SetContinuousSpellCheckingEnabled(
	value bool,
) {
	C.NSTextView_inst_SetContinuousSpellCheckingEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// SpellCheckerDocumentTag returns a tag identifying the text view's text as a document for the spell checker server.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449513-spellcheckerdocumenttag?language=objc for details.
func (x gen_NSTextView) SpellCheckerDocumentTag() core.NSInteger {
	ret := C.NSTextView_inst_SpellCheckerDocumentTag(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// IsGrammarCheckingEnabled enables and disables grammar checking.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449166-grammarcheckingenabled?language=objc for details.
func (x gen_NSTextView) IsGrammarCheckingEnabled() bool {
	ret := C.NSTextView_inst_IsGrammarCheckingEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetGrammarCheckingEnabled enables and disables grammar checking.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449166-grammarcheckingenabled?language=objc for details.
func (x gen_NSTextView) SetGrammarCheckingEnabled(
	value bool,
) {
	C.NSTextView_inst_SetGrammarCheckingEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AcceptsGlyphInfo returns a Boolean value that indicates whether the receiver accepts the glyph info attribute.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449163-acceptsglyphinfo?language=objc for details.
func (x gen_NSTextView) AcceptsGlyphInfo() bool {
	ret := C.NSTextView_inst_AcceptsGlyphInfo(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAcceptsGlyphInfo returns a Boolean value that indicates whether the receiver accepts the glyph info attribute.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449163-acceptsglyphinfo?language=objc for details.
func (x gen_NSTextView) SetAcceptsGlyphInfo(
	value bool,
) {
	C.NSTextView_inst_SetAcceptsGlyphInfo(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// UsesFontPanel returns a Boolean value that controls whether the text views sharing the receiver’s layout manager use the Font panel and Font menu.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449534-usesfontpanel?language=objc for details.
func (x gen_NSTextView) UsesFontPanel() bool {
	ret := C.NSTextView_inst_UsesFontPanel(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetUsesFontPanel returns a Boolean value that controls whether the text views sharing the receiver’s layout manager use the Font panel and Font menu.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449534-usesfontpanel?language=objc for details.
func (x gen_NSTextView) SetUsesFontPanel(
	value bool,
) {
	C.NSTextView_inst_SetUsesFontPanel(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// UsesFindPanel returns a Boolean value that indicates whether the receiver allows for a find panel.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449293-usesfindpanel?language=objc for details.
func (x gen_NSTextView) UsesFindPanel() bool {
	ret := C.NSTextView_inst_UsesFindPanel(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetUsesFindPanel returns a Boolean value that indicates whether the receiver allows for a find panel.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449293-usesfindpanel?language=objc for details.
func (x gen_NSTextView) SetUsesFindPanel(
	value bool,
) {
	C.NSTextView_inst_SetUsesFindPanel(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsAutomaticDashSubstitutionEnabled returns a Boolean value that indicates whether automatic dash substitution is enabled.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449403-automaticdashsubstitutionenabled?language=objc for details.
func (x gen_NSTextView) IsAutomaticDashSubstitutionEnabled() bool {
	ret := C.NSTextView_inst_IsAutomaticDashSubstitutionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAutomaticDashSubstitutionEnabled returns a Boolean value that indicates whether automatic dash substitution is enabled.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449403-automaticdashsubstitutionenabled?language=objc for details.
func (x gen_NSTextView) SetAutomaticDashSubstitutionEnabled(
	value bool,
) {
	C.NSTextView_inst_SetAutomaticDashSubstitutionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsAutomaticDataDetectionEnabled returns a Boolean value that indicates whether automatic data detection is enabled.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449192-automaticdatadetectionenabled?language=objc for details.
func (x gen_NSTextView) IsAutomaticDataDetectionEnabled() bool {
	ret := C.NSTextView_inst_IsAutomaticDataDetectionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAutomaticDataDetectionEnabled returns a Boolean value that indicates whether automatic data detection is enabled.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449192-automaticdatadetectionenabled?language=objc for details.
func (x gen_NSTextView) SetAutomaticDataDetectionEnabled(
	value bool,
) {
	C.NSTextView_inst_SetAutomaticDataDetectionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsAutomaticSpellingCorrectionEnabled returns a Boolean value that indicates whether automatic spelling correction is enabled.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449254-automaticspellingcorrectionenabl?language=objc for details.
func (x gen_NSTextView) IsAutomaticSpellingCorrectionEnabled() bool {
	ret := C.NSTextView_inst_IsAutomaticSpellingCorrectionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAutomaticSpellingCorrectionEnabled returns a Boolean value that indicates whether automatic spelling correction is enabled.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449254-automaticspellingcorrectionenabl?language=objc for details.
func (x gen_NSTextView) SetAutomaticSpellingCorrectionEnabled(
	value bool,
) {
	C.NSTextView_inst_SetAutomaticSpellingCorrectionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsAutomaticTextReplacementEnabled returns a Boolean value that indicates whether automatic text replacement is enabled.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449210-automatictextreplacementenabled?language=objc for details.
func (x gen_NSTextView) IsAutomaticTextReplacementEnabled() bool {
	ret := C.NSTextView_inst_IsAutomaticTextReplacementEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAutomaticTextReplacementEnabled returns a Boolean value that indicates whether automatic text replacement is enabled.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449210-automatictextreplacementenabled?language=objc for details.
func (x gen_NSTextView) SetAutomaticTextReplacementEnabled(
	value bool,
) {
	C.NSTextView_inst_SetAutomaticTextReplacementEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// UsesFindBar returns a Boolean value that indicates whether to use the find bar for this text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449456-usesfindbar?language=objc for details.
func (x gen_NSTextView) UsesFindBar() bool {
	ret := C.NSTextView_inst_UsesFindBar(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetUsesFindBar returns a Boolean value that indicates whether to use the find bar for this text view.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449456-usesfindbar?language=objc for details.
func (x gen_NSTextView) SetUsesFindBar(
	value bool,
) {
	C.NSTextView_inst_SetUsesFindBar(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsIncrementalSearchingEnabled returns a Boolean value that indicates whether incremental searching is enabled.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449458-incrementalsearchingenabled?language=objc for details.
func (x gen_NSTextView) IsIncrementalSearchingEnabled() bool {
	ret := C.NSTextView_inst_IsIncrementalSearchingEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetIncrementalSearchingEnabled returns a Boolean value that indicates whether incremental searching is enabled.
//
// See https://developer.apple.com/documentation/appkit/nstextview/1449458-incrementalsearchingenabled?language=objc for details.
func (x gen_NSTextView) SetIncrementalSearchingEnabled(
	value bool,
) {
	C.NSTextView_inst_SetIncrementalSearchingEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AllowsCharacterPickerTouchBarItem is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nstextview/2544680-allowscharacterpickertouchbarite?language=objc for details.
func (x gen_NSTextView) AllowsCharacterPickerTouchBarItem() bool {
	ret := C.NSTextView_inst_AllowsCharacterPickerTouchBarItem(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsCharacterPickerTouchBarItem is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nstextview/2544680-allowscharacterpickertouchbarite?language=objc for details.
func (x gen_NSTextView) SetAllowsCharacterPickerTouchBarItem(
	value bool,
) {
	C.NSTextView_inst_SetAllowsCharacterPickerTouchBarItem(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// Font is undocumented.
func (x gen_NSTextView) Font() NSFont {
	ret := C.NSTextView_inst_Font(
		unsafe.Pointer(x.Pointer()),
	)

	return NSFont_FromPointer(ret)
}

// SetFont is undocumented.
func (x gen_NSTextView) SetFont(
	value NSFontRef,
) {
	C.NSTextView_inst_SetFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

type NSViewRef interface {
	Pointer() uintptr
	Init() NSView
}

type gen_NSView struct {
	objc.Object
}

func NSView_FromPointer(ptr unsafe.Pointer) NSView {
	return NSView{gen_NSView{
		objc.Object_FromPointer(ptr),
	}}
}

func NSView_FromRef(ref objc.Ref) NSView {
	return NSView_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// AcceptsFirstMouse overridden by subclasses to return YES if the view should be sent a mouseDown: message for an initial mouse-down event, NO if not.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483410-acceptsfirstmouse?language=objc for details.
func (x gen_NSView) AcceptsFirstMouse(
	event NSEventRef,
) bool {
	ret := C.NSView_inst_AcceptsFirstMouse(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return convertObjCBoolToGo(ret)
}

// AddConstraints adds multiple constraints on the layout of the receiving view or its subviews.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526931-addconstraints?language=objc for details.
func (x gen_NSView) AddConstraints(
	constraints core.NSArrayRef,
) {
	C.NSView_inst_AddConstraints(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(constraints),
	)

	return
}

// AddSubview adds a view to the view’s subviews so it’s displayed above its siblings.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483783-addsubview?language=objc for details.
func (x gen_NSView) AddSubview(
	view NSViewRef,
) {
	C.NSView_inst_AddSubview(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)

	return
}

// AddSubviewPositionedRelativeTo inserts a view among the view’s subviews so it’s displayed immediately above or below another view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483640-addsubview?language=objc for details.
func (x gen_NSView) AddSubviewPositionedRelativeTo(
	view NSViewRef,
	place core.NSUInteger,
	otherView NSViewRef,
) {
	C.NSView_inst_AddSubviewPositionedRelativeTo(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
		C.ulong(place),
		objc.RefPointer(otherView),
	)

	return
}

// AdjustScroll overridden by subclasses to modify a given rectangle, returning the altered rectangle.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483616-adjustscroll?language=objc for details.
func (x gen_NSView) AdjustScroll(
	newVisible core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_AdjustScroll(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&newVisible)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// AlignmentRectForFrame returns the view’s alignment rectangle for a given frame.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526905-alignmentrectforframe?language=objc for details.
func (x gen_NSView) AlignmentRectForFrame(
	frame core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_AlignmentRectForFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frame)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// AncestorSharedWithView returns the closest ancestor shared by the view and another specified view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483353-ancestorsharedwithview?language=objc for details.
func (x gen_NSView) AncestorSharedWithView(
	view NSViewRef,
) NSView {
	ret := C.NSView_inst_AncestorSharedWithView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)

	return NSView_FromPointer(ret)
}

// Autoscroll scrolls the view’s closest ancestor NSClipView object proportionally to the distance of an event that occurs outside of it.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483471-autoscroll?language=objc for details.
func (x gen_NSView) Autoscroll(
	event NSEventRef,
) bool {
	ret := C.NSView_inst_Autoscroll(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return convertObjCBoolToGo(ret)
}

// BeginDocument invoked at the beginning of the printing session, this method sets up the current graphics context.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483423-begindocument?language=objc for details.
func (x gen_NSView) BeginDocument() {
	C.NSView_inst_BeginDocument(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// BeginPageInRectAtPlacement called at the beginning of each page, this method sets up the coordinate system so that a region inside the view’s bounds is translated to a specified location.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483438-beginpageinrect?language=objc for details.
func (x gen_NSView) BeginPageInRectAtPlacement(
	rect core.NSRect,
	location core.NSPoint,
) {
	C.NSView_inst_BeginPageInRectAtPlacement(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		*(*C.NSPoint)(unsafe.Pointer(&location)),
	)

	return
}

// CenterScanRect converts the corners of a specified rectangle to lie on the center of device pixels, which is useful in compensating for rendering overscanning when the coordinate system has been scaled.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483725-centerscanrect?language=objc for details.
func (x gen_NSView) CenterScanRect(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_CenterScanRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// ConvertPointFromView converts a point from the coordinate system of a given view to that of the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483269-convertpoint?language=objc for details.
func (x gen_NSView) ConvertPointFromView(
	point core.NSPoint,
	view NSViewRef,
) core.NSPoint {
	ret := C.NSView_inst_ConvertPointFromView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
		objc.RefPointer(view),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// ConvertPointToView converts a point from the view’s coordinate system to that of a given view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483406-convertpoint?language=objc for details.
func (x gen_NSView) ConvertPointToView(
	point core.NSPoint,
	view NSViewRef,
) core.NSPoint {
	ret := C.NSView_inst_ConvertPointToView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
		objc.RefPointer(view),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// ConvertPointFromBacking converts a point from its pixel aligned backing store coordinate system to the view’s interior coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483456-convertpointfrombacking?language=objc for details.
func (x gen_NSView) ConvertPointFromBacking(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSView_inst_ConvertPointFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// ConvertPointFromLayer convert the point from the layer's interior coordinate system to the view’s interior coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483554-convertpointfromlayer?language=objc for details.
func (x gen_NSView) ConvertPointFromLayer(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSView_inst_ConvertPointFromLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// ConvertPointToBacking converts a point from the view’s interior coordinate system to its pixel aligned backing store coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483803-convertpointtobacking?language=objc for details.
func (x gen_NSView) ConvertPointToBacking(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSView_inst_ConvertPointToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// ConvertPointToLayer convert the size from the view’s interior coordinate system to the layer's interior coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483315-convertpointtolayer?language=objc for details.
func (x gen_NSView) ConvertPointToLayer(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSView_inst_ConvertPointToLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// ConvertRectFromView converts a rectangle from the coordinate system of another view to that of the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483785-convertrect?language=objc for details.
func (x gen_NSView) ConvertRectFromView(
	rect core.NSRect,
	view NSViewRef,
) core.NSRect {
	ret := C.NSView_inst_ConvertRectFromView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(view),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// ConvertRectToView converts a rectangle from the view’s coordinate system to that of another view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483217-convertrect?language=objc for details.
func (x gen_NSView) ConvertRectToView(
	rect core.NSRect,
	view NSViewRef,
) core.NSRect {
	ret := C.NSView_inst_ConvertRectToView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(view),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// ConvertRectFromBacking converts a rectangle from its pixel aligned backing store coordinate system to the view’s interior coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483819-convertrectfrombacking?language=objc for details.
func (x gen_NSView) ConvertRectFromBacking(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_ConvertRectFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// ConvertRectFromLayer convert the rectangle from the layer's interior coordinate system to the view’s interior coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483404-convertrectfromlayer?language=objc for details.
func (x gen_NSView) ConvertRectFromLayer(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_ConvertRectFromLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// ConvertRectToBacking converts a rectangle from the view’s interior coordinate system to its pixel aligned backing store coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483648-convertrecttobacking?language=objc for details.
func (x gen_NSView) ConvertRectToBacking(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_ConvertRectToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// ConvertRectToLayer convert the size from the view’s interior coordinate system to the layer's interior coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483776-convertrecttolayer?language=objc for details.
func (x gen_NSView) ConvertRectToLayer(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_ConvertRectToLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// ConvertSizeFromView converts a size from another view’s coordinate system to that of the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483307-convertsize?language=objc for details.
func (x gen_NSView) ConvertSizeFromView(
	size core.NSSize,
	view NSViewRef,
) core.NSSize {
	ret := C.NSView_inst_ConvertSizeFromView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
		objc.RefPointer(view),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// ConvertSizeToView converts a size from the view’s coordinate system to that of another view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483744-convertsize?language=objc for details.
func (x gen_NSView) ConvertSizeToView(
	size core.NSSize,
	view NSViewRef,
) core.NSSize {
	ret := C.NSView_inst_ConvertSizeToView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
		objc.RefPointer(view),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// ConvertSizeFromBacking converts a size from its pixel aligned backing store coordinate system to the view’s interior coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483319-convertsizefrombacking?language=objc for details.
func (x gen_NSView) ConvertSizeFromBacking(
	size core.NSSize,
) core.NSSize {
	ret := C.NSView_inst_ConvertSizeFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// ConvertSizeFromLayer convert the size from the layer's interior coordinate system to the view’s interior coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483479-convertsizefromlayer?language=objc for details.
func (x gen_NSView) ConvertSizeFromLayer(
	size core.NSSize,
) core.NSSize {
	ret := C.NSView_inst_ConvertSizeFromLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// ConvertSizeToBacking converts a size from the view’s interior coordinate system to its pixel aligned backing store coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483227-convertsizetobacking?language=objc for details.
func (x gen_NSView) ConvertSizeToBacking(
	size core.NSSize,
) core.NSSize {
	ret := C.NSView_inst_ConvertSizeToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// ConvertSizeToLayer convert the size from the view’s interior coordinate system to the layer's interior coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483701-convertsizetolayer?language=objc for details.
func (x gen_NSView) ConvertSizeToLayer(
	size core.NSSize,
) core.NSSize {
	ret := C.NSView_inst_ConvertSizeToLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// DataWithEPSInsideRect returns EPS data that draws the region of the view within a specified rectangle.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483735-datawithepsinsiderect?language=objc for details.
func (x gen_NSView) DataWithEPSInsideRect(
	rect core.NSRect,
) core.NSData {
	ret := C.NSView_inst_DataWithEPSInsideRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return core.NSData_FromPointer(ret)
}

// DataWithPDFInsideRect returns PDF data that draws the region of the view within a specified rectangle.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483797-datawithpdfinsiderect?language=objc for details.
func (x gen_NSView) DataWithPDFInsideRect(
	rect core.NSRect,
) core.NSData {
	ret := C.NSView_inst_DataWithPDFInsideRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return core.NSData_FromPointer(ret)
}

// DidAddSubview overridden by subclasses to perform additional actions when subviews are added to the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483454-didaddsubview?language=objc for details.
func (x gen_NSView) DidAddSubview(
	subview NSViewRef,
) {
	C.NSView_inst_DidAddSubview(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(subview),
	)

	return
}

// DidCloseMenuWithEvent called after a contextual menu that was displayed from the receiving view has been closed.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483770-didclosemenu?language=objc for details.
func (x gen_NSView) DidCloseMenuWithEvent(
	menu NSMenuRef,
	event NSEventRef,
) {
	C.NSView_inst_DidCloseMenuWithEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(menu),
		objc.RefPointer(event),
	)

	return
}

// DiscardCursorRects invalidates all cursor rectangles set up using addCursorRect:cursor:.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483733-discardcursorrects?language=objc for details.
func (x gen_NSView) DiscardCursorRects() {
	C.NSView_inst_DiscardCursorRects(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Display displays the view and all its subviews if possible, invoking each of the NSView methods lockFocus, drawRect:, and unlockFocus as necessary.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483487-display?language=objc for details.
func (x gen_NSView) Display() {
	C.NSView_inst_Display(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DisplayIfNeeded displays the view and all its subviews if any part of the view has been marked as needing display.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483566-displayifneeded?language=objc for details.
func (x gen_NSView) DisplayIfNeeded() {
	C.NSView_inst_DisplayIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DisplayIfNeededIgnoringOpacity acts as displayIfNeeded, except that this method doesn’t back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483526-displayifneededignoringopacity?language=objc for details.
func (x gen_NSView) DisplayIfNeededIgnoringOpacity() {
	C.NSView_inst_DisplayIfNeededIgnoringOpacity(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DisplayIfNeededInRect acts as displayIfNeeded, confining drawing to a specified region of the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483813-displayifneededinrect?language=objc for details.
func (x gen_NSView) DisplayIfNeededInRect(
	rect core.NSRect,
) {
	C.NSView_inst_DisplayIfNeededInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return
}

// DisplayIfNeededInRectIgnoringOpacity acts as displayIfNeeded, but confining drawing to aRect and not backing up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483481-displayifneededinrectignoringopa?language=objc for details.
func (x gen_NSView) DisplayIfNeededInRectIgnoringOpacity(
	rect core.NSRect,
) {
	C.NSView_inst_DisplayIfNeededInRectIgnoringOpacity(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return
}

// DisplayRect acts as display, but confining drawing to a rectangular region of the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483518-displayrect?language=objc for details.
func (x gen_NSView) DisplayRect(
	rect core.NSRect,
) {
	C.NSView_inst_DisplayRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return
}

// DisplayRectIgnoringOpacity displays the view but confines drawing to a specified region and does not back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483699-displayrectignoringopacity?language=objc for details.
func (x gen_NSView) DisplayRectIgnoringOpacity(
	rect core.NSRect,
) {
	C.NSView_inst_DisplayRectIgnoringOpacity(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return
}

// DrawFocusRingMask draws the focus ring mask for the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483335-drawfocusringmask?language=objc for details.
func (x gen_NSView) DrawFocusRingMask() {
	C.NSView_inst_DrawFocusRingMask(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DrawPageBorderWithSize allows applications that use the AppKit pagination facility to draw additional marks on each logical page.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483292-drawpageborderwithsize?language=objc for details.
func (x gen_NSView) DrawPageBorderWithSize(
	borderSize core.NSSize,
) {
	C.NSView_inst_DrawPageBorderWithSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&borderSize)),
	)

	return
}

// DrawRect overridden by subclasses to draw the view’s image within the specified rectangle.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483686-drawrect?language=objc for details.
func (x gen_NSView) DrawRect(
	dirtyRect core.NSRect,
) {
	C.NSView_inst_DrawRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&dirtyRect)),
	)

	return
}

// EndDocument this method is invoked at the end of the printing session.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483610-enddocument?language=objc for details.
func (x gen_NSView) EndDocument() {
	C.NSView_inst_EndDocument(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// EndPage writes the end of a conforming page.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483549-endpage?language=objc for details.
func (x gen_NSView) EndPage() {
	C.NSView_inst_EndPage(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// EnterFullScreenModeWithOptions sets the view to full screen mode.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483780-enterfullscreenmode?language=objc for details.
func (x gen_NSView) EnterFullScreenModeWithOptions(
	screen NSScreenRef,
	options core.NSDictionaryRef,
) bool {
	ret := C.NSView_inst_EnterFullScreenModeWithOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(screen),
		objc.RefPointer(options),
	)

	return convertObjCBoolToGo(ret)
}

// ExerciseAmbiguityInLayout randomly changes the frame of a view with an ambiguous layout between the different valid values.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526934-exerciseambiguityinlayout?language=objc for details.
func (x gen_NSView) ExerciseAmbiguityInLayout() {
	C.NSView_inst_ExerciseAmbiguityInLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ExitFullScreenModeWithOptions instructs the view to exit full screen mode.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483256-exitfullscreenmodewithoptions?language=objc for details.
func (x gen_NSView) ExitFullScreenModeWithOptions(
	options core.NSDictionaryRef,
) {
	C.NSView_inst_ExitFullScreenModeWithOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(options),
	)

	return
}

// FrameForAlignmentRect returns the view’s frame for a given alignment rectangle.
//
// See https://developer.apple.com/documentation/appkit/nsview/1525584-frameforalignmentrect?language=objc for details.
func (x gen_NSView) FrameForAlignmentRect(
	alignmentRect core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_FrameForAlignmentRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&alignmentRect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// HitTest returns the farthest descendant of the view in the view hierarchy (including itself) that contains a specified point, or nil if that point lies completely outside the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483364-hittest?language=objc for details.
func (x gen_NSView) HitTest(
	point core.NSPoint,
) NSView {
	ret := C.NSView_inst_HitTest(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return NSView_FromPointer(ret)
}

// InitWithFrame initializes and returns a newly allocated NSView object with a specified frame rectangle.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc for details.
func (x gen_NSView) InitWithFrame(
	frameRect core.NSRect,
) NSView {
	ret := C.NSView_inst_InitWithFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)

	return NSView_FromPointer(ret)
}

// InvalidateIntrinsicContentSize invalidates the view’s intrinsic content size.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526864-invalidateintrinsiccontentsize?language=objc for details.
func (x gen_NSView) InvalidateIntrinsicContentSize() {
	C.NSView_inst_InvalidateIntrinsicContentSize(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// IsDescendantOf returns YES if the view is a subview of a given view or if it’s identical to that view; otherwise, it returns NO.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483219-isdescendantof?language=objc for details.
func (x gen_NSView) IsDescendantOf(
	view NSViewRef,
) bool {
	ret := C.NSView_inst_IsDescendantOf(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)

	return convertObjCBoolToGo(ret)
}

// Layout perform layout in concert with the constraint-based layout system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526146-layout?language=objc for details.
func (x gen_NSView) Layout() {
	C.NSView_inst_Layout(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// LayoutSubtreeIfNeeded updates the layout of the receiving view and its subviews based on the current views and constraints.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526871-layoutsubtreeifneeded?language=objc for details.
func (x gen_NSView) LayoutSubtreeIfNeeded() {
	C.NSView_inst_LayoutSubtreeIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// LocationOfPrintRect invoked by print: to determine the location of the region of the view being printed on the physical page.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483223-locationofprintrect?language=objc for details.
func (x gen_NSView) LocationOfPrintRect(
	rect core.NSRect,
) core.NSPoint {
	ret := C.NSView_inst_LocationOfPrintRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))
}

// MakeBackingLayer creates the view’s backing layer.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483687-makebackinglayer?language=objc for details.
func (x gen_NSView) MakeBackingLayer() core.CALayer {
	ret := C.NSView_inst_MakeBackingLayer(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CALayer_FromPointer(ret)
}

// MenuForEvent overridden by subclasses to return a context-sensitive pop-up menu for a given mouse-down event.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483231-menuforevent?language=objc for details.
func (x gen_NSView) MenuForEvent(
	event NSEventRef,
) NSMenu {
	ret := C.NSView_inst_MenuForEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return NSMenu_FromPointer(ret)
}

// MouseInRect returns whether a region of the view contains a specified point, accounting for whether the view is flipped or not.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483237-mouse?language=objc for details.
func (x gen_NSView) MouseInRect(
	point core.NSPoint,
	rect core.NSRect,
) bool {
	ret := C.NSView_inst_MouseInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return convertObjCBoolToGo(ret)
}

// NeedsToDrawRect returns a Boolean value indicating whether the specified rectangle intersects any part of the area that the view is being asked to draw.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483570-needstodrawrect?language=objc for details.
func (x gen_NSView) NeedsToDrawRect(
	rect core.NSRect,
) bool {
	ret := C.NSView_inst_NeedsToDrawRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return convertObjCBoolToGo(ret)
}

// NoteFocusRingMaskChanged invoked to notify the view that the focus ring mask requires updating.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483809-notefocusringmaskchanged?language=objc for details.
func (x gen_NSView) NoteFocusRingMaskChanged() {
	C.NSView_inst_NoteFocusRingMaskChanged(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// PerformKeyEquivalent implemented by subclasses to respond to key equivalents (also known as keyboard shortcuts).
//
// See https://developer.apple.com/documentation/appkit/nsview/1483664-performkeyequivalent?language=objc for details.
func (x gen_NSView) PerformKeyEquivalent(
	event NSEventRef,
) bool {
	ret := C.NSView_inst_PerformKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return convertObjCBoolToGo(ret)
}

// PrepareContentInRect prepares the overdraw region for drawing.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483427-preparecontentinrect?language=objc for details.
func (x gen_NSView) PrepareContentInRect(
	rect core.NSRect,
) {
	C.NSView_inst_PrepareContentInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return
}

// PrepareForReuse restores the view to an initial state so that it can be reused.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483626-prepareforreuse?language=objc for details.
func (x gen_NSView) PrepareForReuse() {
	C.NSView_inst_PrepareForReuse(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Print this action method opens the Print panel, and if the user chooses an option other than canceling, prints the view and all its subviews to the device specified in the Print panel.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483705-print?language=objc for details.
func (x gen_NSView) Print(
	sender objc.Ref,
) {
	C.NSView_inst_Print(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// RectForPage implemented by subclasses to determine the portion of the view to be printed for the specified page number.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483252-rectforpage?language=objc for details.
func (x gen_NSView) RectForPage(
	page core.NSInteger,
) core.NSRect {
	ret := C.NSView_inst_RectForPage(
		unsafe.Pointer(x.Pointer()),
		C.long(page),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// RectForSmartMagnificationAtPointInRect returns the appropriate rectangle to use when magnifying around the specified point.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483305-rectforsmartmagnificationatpoint?language=objc for details.
func (x gen_NSView) RectForSmartMagnificationAtPointInRect(
	location core.NSPoint,
	visibleRect core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_RectForSmartMagnificationAtPointInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&location)),
		*(*C.NSRect)(unsafe.Pointer(&visibleRect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// RegisterForDraggedTypes registers the pasteboard types that the view will accept as the destination of an image-dragging session.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483578-registerfordraggedtypes?language=objc for details.
func (x gen_NSView) RegisterForDraggedTypes(
	newTypes core.NSArrayRef,
) {
	C.NSView_inst_RegisterForDraggedTypes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newTypes),
	)

	return
}

// RemoveAllToolTips removes all tooltips assigned to the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483801-removealltooltips?language=objc for details.
func (x gen_NSView) RemoveAllToolTips() {
	C.NSView_inst_RemoveAllToolTips(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// RemoveConstraints removes the specified constraints from the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526932-removeconstraints?language=objc for details.
func (x gen_NSView) RemoveConstraints(
	constraints core.NSArrayRef,
) {
	C.NSView_inst_RemoveConstraints(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(constraints),
	)

	return
}

// RemoveFromSuperview unlinks the view from its superview and its window, removes it from the responder chain, and invalidates its cursor rectangles.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483265-removefromsuperview?language=objc for details.
func (x gen_NSView) RemoveFromSuperview() {
	C.NSView_inst_RemoveFromSuperview(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// RemoveFromSuperviewWithoutNeedingDisplay unlinks the view from its superview and its window and removes it from the responder chain, but does not invalidate its cursor rectangles to cause redrawing.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483644-removefromsuperviewwithoutneedin?language=objc for details.
func (x gen_NSView) RemoveFromSuperviewWithoutNeedingDisplay() {
	C.NSView_inst_RemoveFromSuperviewWithoutNeedingDisplay(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ReplaceSubviewWith replaces one of the view’s subviews with another view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483632-replacesubview?language=objc for details.
func (x gen_NSView) ReplaceSubviewWith(
	oldView NSViewRef,
	newView NSViewRef,
) {
	C.NSView_inst_ReplaceSubviewWith(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(oldView),
		objc.RefPointer(newView),
	)

	return
}

// ResetCursorRects overridden by subclasses to define their default cursor rectangles.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483448-resetcursorrects?language=objc for details.
func (x gen_NSView) ResetCursorRects() {
	C.NSView_inst_ResetCursorRects(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ResizeSubviewsWithOldSize informs the view’s subviews that the view’s bounds rectangle size has changed.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483495-resizesubviewswitholdsize?language=objc for details.
func (x gen_NSView) ResizeSubviewsWithOldSize(
	oldSize core.NSSize,
) {
	C.NSView_inst_ResizeSubviewsWithOldSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&oldSize)),
	)

	return
}

// ResizeWithOldSuperviewSize informs the view that the bounds size of its superview has changed.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483697-resizewitholdsuperviewsize?language=objc for details.
func (x gen_NSView) ResizeWithOldSuperviewSize(
	oldSize core.NSSize,
) {
	C.NSView_inst_ResizeWithOldSuperviewSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&oldSize)),
	)

	return
}

// RotateByAngle rotates the view’s bounds rectangle by a specified degree value around the origin of the coordinate system, (0.0, 0.0).
//
// See https://developer.apple.com/documentation/appkit/nsview/1483444-rotatebyangle?language=objc for details.
func (x gen_NSView) RotateByAngle(
	angle core.CGFloat,
) {
	C.NSView_inst_RotateByAngle(
		unsafe.Pointer(x.Pointer()),
		C.double(angle),
	)

	return
}

// ScaleUnitSquareToSize scales the view’s coordinate system so that the unit square scales to the specified dimensions.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483721-scaleunitsquaretosize?language=objc for details.
func (x gen_NSView) ScaleUnitSquareToSize(
	newUnitSize core.NSSize,
) {
	C.NSView_inst_ScaleUnitSquareToSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&newUnitSize)),
	)

	return
}

// ScrollPoint scrolls the view’s closest ancestor NSClipView object so a point in the view lies at the origin of the clip view's bounds rectangle.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483394-scrollpoint?language=objc for details.
func (x gen_NSView) ScrollPoint(
	point core.NSPoint,
) {
	C.NSView_inst_ScrollPoint(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return
}

// ScrollRectToVisible scrolls the view’s closest ancestor NSClipView object the minimum distance needed so a specified region of the view becomes visible in the clip view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483811-scrollrecttovisible?language=objc for details.
func (x gen_NSView) ScrollRectToVisible(
	rect core.NSRect,
) bool {
	ret := C.NSView_inst_ScrollRectToVisible(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return convertObjCBoolToGo(ret)
}

// SetBoundsOrigin sets the origin of the view’s bounds rectangle to a specified point.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483345-setboundsorigin?language=objc for details.
func (x gen_NSView) SetBoundsOrigin(
	newOrigin core.NSPoint,
) {
	C.NSView_inst_SetBoundsOrigin(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&newOrigin)),
	)

	return
}

// SetBoundsSize sets the size of the view’s bounds rectangle to specified dimensions, inversely scaling its coordinate system relative to its frame rectangle.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483399-setboundssize?language=objc for details.
func (x gen_NSView) SetBoundsSize(
	newSize core.NSSize,
) {
	C.NSView_inst_SetBoundsSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&newSize)),
	)

	return
}

// SetFrameOrigin sets the origin of the view’s frame rectangle to the specified point, effectively repositioning it within its superview.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483283-setframeorigin?language=objc for details.
func (x gen_NSView) SetFrameOrigin(
	newOrigin core.NSPoint,
) {
	C.NSView_inst_SetFrameOrigin(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&newOrigin)),
	)

	return
}

// SetFrameSize sets the size of the view’s frame rectangle to the specified dimensions, resizing it within its superview without affecting its coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483530-setframesize?language=objc for details.
func (x gen_NSView) SetFrameSize(
	newSize core.NSSize,
) {
	C.NSView_inst_SetFrameSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&newSize)),
	)

	return
}

// SetKeyboardFocusRingNeedsDisplayInRect invalidates the area around the focus ring.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483556-setkeyboardfocusringneedsdisplay?language=objc for details.
func (x gen_NSView) SetKeyboardFocusRingNeedsDisplayInRect(
	rect core.NSRect,
) {
	C.NSView_inst_SetKeyboardFocusRingNeedsDisplayInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return
}

// SetNeedsDisplayInRect marks the region of the view within the specified rectangle as needing display, increasing the view’s existing invalid region to include it.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483475-setneedsdisplayinrect?language=objc for details.
func (x gen_NSView) SetNeedsDisplayInRect(
	invalidRect core.NSRect,
) {
	C.NSView_inst_SetNeedsDisplayInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&invalidRect)),
	)

	return
}

// ShouldDelayWindowOrderingForEvent allows the user to drag objects from the view without activating the app or moving the window of the view forward, possibly obscuring the destination.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483244-shoulddelaywindoworderingforeven?language=objc for details.
func (x gen_NSView) ShouldDelayWindowOrderingForEvent(
	event NSEventRef,
) bool {
	ret := C.NSView_inst_ShouldDelayWindowOrderingForEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return convertObjCBoolToGo(ret)
}

// ShowDefinitionForAttributedStringAtPoint shows a window displaying the definition of the attributed string at the specified point.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483747-showdefinitionforattributedstrin?language=objc for details.
func (x gen_NSView) ShowDefinitionForAttributedStringAtPoint(
	attrString core.NSAttributedStringRef,
	textBaselineOrigin core.NSPoint,
) {
	C.NSView_inst_ShowDefinitionForAttributedStringAtPoint(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(attrString),
		*(*C.NSPoint)(unsafe.Pointer(&textBaselineOrigin)),
	)

	return
}

// TranslateOriginToPoint translates the view’s coordinate system so that its origin moves to a new location.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483385-translateorigintopoint?language=objc for details.
func (x gen_NSView) TranslateOriginToPoint(
	translation core.NSPoint,
) {
	C.NSView_inst_TranslateOriginToPoint(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&translation)),
	)

	return
}

// TranslateRectsNeedingDisplayInRectBy translates the display rectangles by the specified delta.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483731-translaterectsneedingdisplayinre?language=objc for details.
func (x gen_NSView) TranslateRectsNeedingDisplayInRectBy(
	clipRect core.NSRect,
	delta core.NSSize,
) {
	C.NSView_inst_TranslateRectsNeedingDisplayInRectBy(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&clipRect)),
		*(*C.NSSize)(unsafe.Pointer(&delta)),
	)

	return
}

// UnregisterDraggedTypes unregisters the view as a possible destination in a dragging session.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483602-unregisterdraggedtypes?language=objc for details.
func (x gen_NSView) UnregisterDraggedTypes() {
	C.NSView_inst_UnregisterDraggedTypes(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// UpdateConstraints update constraints for the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526891-updateconstraints?language=objc for details.
func (x gen_NSView) UpdateConstraints() {
	C.NSView_inst_UpdateConstraints(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// UpdateConstraintsForSubtreeIfNeeded updates the constraints for the receiving view and its subviews.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526939-updateconstraintsforsubtreeifnee?language=objc for details.
func (x gen_NSView) UpdateConstraintsForSubtreeIfNeeded() {
	C.NSView_inst_UpdateConstraintsForSubtreeIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// UpdateLayer updates the view’s content by modifying its underlying layer.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483580-updatelayer?language=objc for details.
func (x gen_NSView) UpdateLayer() {
	C.NSView_inst_UpdateLayer(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// UpdateTrackingAreas invoked automatically when the view’s geometry changes such that its tracking areas need to be recalculated.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483719-updatetrackingareas?language=objc for details.
func (x gen_NSView) UpdateTrackingAreas() {
	C.NSView_inst_UpdateTrackingAreas(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewDidChangeBackingProperties responds when the view’s backing store properties change.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483742-viewdidchangebackingproperties?language=objc for details.
func (x gen_NSView) ViewDidChangeBackingProperties() {
	C.NSView_inst_ViewDidChangeBackingProperties(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewDidChangeEffectiveAppearance is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsview/2977088-viewdidchangeeffectiveappearance?language=objc for details.
func (x gen_NSView) ViewDidChangeEffectiveAppearance() {
	C.NSView_inst_ViewDidChangeEffectiveAppearance(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewDidEndLiveResize informs the view of the end of a live resize—the user has finished resizing the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483543-viewdidendliveresize?language=objc for details.
func (x gen_NSView) ViewDidEndLiveResize() {
	C.NSView_inst_ViewDidEndLiveResize(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewDidHide invoked when the view is hidden, either directly, or in response to an ancestor being hidden.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483596-viewdidhide?language=objc for details.
func (x gen_NSView) ViewDidHide() {
	C.NSView_inst_ViewDidHide(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewDidMoveToSuperview informs the view that its superview has changed (possibly to nil).
//
// See https://developer.apple.com/documentation/appkit/nsview/1483568-viewdidmovetosuperview?language=objc for details.
func (x gen_NSView) ViewDidMoveToSuperview() {
	C.NSView_inst_ViewDidMoveToSuperview(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewDidMoveToWindow informs the view that it has been added to a new view hierarchy.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483329-viewdidmovetowindow?language=objc for details.
func (x gen_NSView) ViewDidMoveToWindow() {
	C.NSView_inst_ViewDidMoveToWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewDidUnhide invoked when the view is unhidden, either directly, or in response to an ancestor being unhidden
//
// See https://developer.apple.com/documentation/appkit/nsview/1483275-viewdidunhide?language=objc for details.
func (x gen_NSView) ViewDidUnhide() {
	C.NSView_inst_ViewDidUnhide(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewWillDraw informs the view that it’s required to draw content.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483351-viewwilldraw?language=objc for details.
func (x gen_NSView) ViewWillDraw() {
	C.NSView_inst_ViewWillDraw(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewWillMoveToSuperview informs the view that its superview is about to change to the specified superview (which may be nil).
//
// See https://developer.apple.com/documentation/appkit/nsview/1483545-viewwillmovetosuperview?language=objc for details.
func (x gen_NSView) ViewWillMoveToSuperview(
	newSuperview NSViewRef,
) {
	C.NSView_inst_ViewWillMoveToSuperview(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newSuperview),
	)

	return
}

// ViewWillMoveToWindow informs the view that it’s being added to the view hierarchy of the specified window object (which may be nil).
//
// See https://developer.apple.com/documentation/appkit/nsview/1483415-viewwillmovetowindow?language=objc for details.
func (x gen_NSView) ViewWillMoveToWindow(
	newWindow NSWindowRef,
) {
	C.NSView_inst_ViewWillMoveToWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newWindow),
	)

	return
}

// ViewWillStartLiveResize informs the view of the start of a live resize—the user has started resizing the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483620-viewwillstartliveresize?language=objc for details.
func (x gen_NSView) ViewWillStartLiveResize() {
	C.NSView_inst_ViewWillStartLiveResize(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ViewWithTag returns the view’s nearest descendant (including itself) with a specific tag, or nil if no subview has that tag.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483294-viewwithtag?language=objc for details.
func (x gen_NSView) ViewWithTag(
	tag core.NSInteger,
) NSView {
	ret := C.NSView_inst_ViewWithTag(
		unsafe.Pointer(x.Pointer()),
		C.long(tag),
	)

	return NSView_FromPointer(ret)
}

// WillOpenMenuWithEvent called just before a contextual menu for a view is opened on screen.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483429-willopenmenu?language=objc for details.
func (x gen_NSView) WillOpenMenuWithEvent(
	menu NSMenuRef,
	event NSEventRef,
) {
	C.NSView_inst_WillOpenMenuWithEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(menu),
		objc.RefPointer(event),
	)

	return
}

// WillRemoveSubview overridden by subclasses to perform additional actions before subviews are removed from the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483624-willremovesubview?language=objc for details.
func (x gen_NSView) WillRemoveSubview(
	subview NSViewRef,
) {
	C.NSView_inst_WillRemoveSubview(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(subview),
	)

	return
}

// WriteEPSInsideRectToPasteboard writes EPS data that draws the region of the view within a specified rectangle onto a pasteboard.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483235-writeepsinsiderect?language=objc for details.
func (x gen_NSView) WriteEPSInsideRectToPasteboard(
	rect core.NSRect,
	pasteboard NSPasteboardRef,
) {
	C.NSView_inst_WriteEPSInsideRectToPasteboard(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(pasteboard),
	)

	return
}

// WritePDFInsideRectToPasteboard writes PDF data that draws the region of the view within a specified rectangle onto a pasteboard.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483499-writepdfinsiderect?language=objc for details.
func (x gen_NSView) WritePDFInsideRectToPasteboard(
	rect core.NSRect,
	pasteboard NSPasteboardRef,
) {
	C.NSView_inst_WritePDFInsideRectToPasteboard(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(pasteboard),
	)

	return
}

// Init is undocumented.
func (x gen_NSView) Init() NSView {
	ret := C.NSView_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_FromPointer(ret)
}

// Superview returns the view that is the parent of the current view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483737-superview?language=objc for details.
func (x gen_NSView) Superview() NSView {
	ret := C.NSView_inst_Superview(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_FromPointer(ret)
}

// Subviews returns the array of views embedded in the current view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483539-subviews?language=objc for details.
func (x gen_NSView) Subviews() core.NSArray {
	ret := C.NSView_inst_Subviews(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// SetSubviews returns the array of views embedded in the current view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483539-subviews?language=objc for details.
func (x gen_NSView) SetSubviews(
	value core.NSArrayRef,
) {
	C.NSView_inst_SetSubviews(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Window returns the view’s window object, if it is installed in a window.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483301-window?language=objc for details.
func (x gen_NSView) Window() NSWindow {
	ret := C.NSView_inst_Window(
		unsafe.Pointer(x.Pointer()),
	)

	return NSWindow_FromPointer(ret)
}

// OpaqueAncestor returns the view’s closest opaque ancestor, which might be the view itself.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483383-opaqueancestor?language=objc for details.
func (x gen_NSView) OpaqueAncestor() NSView {
	ret := C.NSView_inst_OpaqueAncestor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_FromPointer(ret)
}

// EnclosingMenuItem returns the menu item containing the view or any of its superviews in the view hierarchy.
//
// See https://developer.apple.com/documentation/appkit/nsview/1514865-enclosingmenuitem?language=objc for details.
func (x gen_NSView) EnclosingMenuItem() NSMenuItem {
	ret := C.NSView_inst_EnclosingMenuItem(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenuItem_FromPointer(ret)
}

// Frame returns the view’s frame rectangle, which defines its position and size in its superview’s coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483713-frame?language=objc for details.
func (x gen_NSView) Frame() core.NSRect {
	ret := C.NSView_inst_Frame(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// SetFrame returns the view’s frame rectangle, which defines its position and size in its superview’s coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483713-frame?language=objc for details.
func (x gen_NSView) SetFrame(
	value core.NSRect,
) {
	C.NSView_inst_SetFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return
}

// FrameRotation returns the angle of rotation, measured in degrees, applied to the view’s frame rectangle relative to its superview’s coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483412-framerotation?language=objc for details.
func (x gen_NSView) FrameRotation() core.CGFloat {
	ret := C.NSView_inst_FrameRotation(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// SetFrameRotation returns the angle of rotation, measured in degrees, applied to the view’s frame rectangle relative to its superview’s coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483412-framerotation?language=objc for details.
func (x gen_NSView) SetFrameRotation(
	value core.CGFloat,
) {
	C.NSView_inst_SetFrameRotation(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// Bounds returns the view’s bounds rectangle, which expresses its location and size in its own coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483817-bounds?language=objc for details.
func (x gen_NSView) Bounds() core.NSRect {
	ret := C.NSView_inst_Bounds(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// SetBounds returns the view’s bounds rectangle, which expresses its location and size in its own coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483817-bounds?language=objc for details.
func (x gen_NSView) SetBounds(
	value core.NSRect,
) {
	C.NSView_inst_SetBounds(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return
}

// BoundsRotation returns the angle of rotation, measured in degrees, applied to the view’s bounds rectangle relative to its frame rectangle.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483746-boundsrotation?language=objc for details.
func (x gen_NSView) BoundsRotation() core.CGFloat {
	ret := C.NSView_inst_BoundsRotation(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// SetBoundsRotation returns the angle of rotation, measured in degrees, applied to the view’s bounds rectangle relative to its frame rectangle.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483746-boundsrotation?language=objc for details.
func (x gen_NSView) SetBoundsRotation(
	value core.CGFloat,
) {
	C.NSView_inst_SetBoundsRotation(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// WantsLayer returns a Boolean value indicating whether the view uses a layer as its backing store.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483695-wantslayer?language=objc for details.
func (x gen_NSView) WantsLayer() bool {
	ret := C.NSView_inst_WantsLayer(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetWantsLayer returns a Boolean value indicating whether the view uses a layer as its backing store.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483695-wantslayer?language=objc for details.
func (x gen_NSView) SetWantsLayer(
	value bool,
) {
	C.NSView_inst_SetWantsLayer(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// WantsUpdateLayer returns a Boolean value indicating which drawing path the view takes when updating its contents.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483461-wantsupdatelayer?language=objc for details.
func (x gen_NSView) WantsUpdateLayer() bool {
	ret := C.NSView_inst_WantsUpdateLayer(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Layer returns the Core Animation layer that the view uses as its backing store.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483298-layer?language=objc for details.
func (x gen_NSView) Layer() core.CALayer {
	ret := C.NSView_inst_Layer(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CALayer_FromPointer(ret)
}

// SetLayer returns the Core Animation layer that the view uses as its backing store.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483298-layer?language=objc for details.
func (x gen_NSView) SetLayer(
	value core.CALayerRef,
) {
	C.NSView_inst_SetLayer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// CanDrawSubviewsIntoLayer returns a Boolean value indicating whether the view incorporates content from its subviews into its own layer.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483347-candrawsubviewsintolayer?language=objc for details.
func (x gen_NSView) CanDrawSubviewsIntoLayer() bool {
	ret := C.NSView_inst_CanDrawSubviewsIntoLayer(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetCanDrawSubviewsIntoLayer returns a Boolean value indicating whether the view incorporates content from its subviews into its own layer.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483347-candrawsubviewsintolayer?language=objc for details.
func (x gen_NSView) SetCanDrawSubviewsIntoLayer(
	value bool,
) {
	C.NSView_inst_SetCanDrawSubviewsIntoLayer(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// LayerUsesCoreImageFilters returns a Boolean value indicating whether the view’s layer uses Core Image filters and needs in-process rendering.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483576-layerusescoreimagefilters?language=objc for details.
func (x gen_NSView) LayerUsesCoreImageFilters() bool {
	ret := C.NSView_inst_LayerUsesCoreImageFilters(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetLayerUsesCoreImageFilters returns a Boolean value indicating whether the view’s layer uses Core Image filters and needs in-process rendering.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483576-layerusescoreimagefilters?language=objc for details.
func (x gen_NSView) SetLayerUsesCoreImageFilters(
	value bool,
) {
	C.NSView_inst_SetLayerUsesCoreImageFilters(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AlphaValue returns the opacity of the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483560-alphavalue?language=objc for details.
func (x gen_NSView) AlphaValue() core.CGFloat {
	ret := C.NSView_inst_AlphaValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// SetAlphaValue returns the opacity of the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483560-alphavalue?language=objc for details.
func (x gen_NSView) SetAlphaValue(
	value core.CGFloat,
) {
	C.NSView_inst_SetAlphaValue(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// FrameCenterRotation returns the rotation angle of the view around the center of its layer.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483367-framecenterrotation?language=objc for details.
func (x gen_NSView) FrameCenterRotation() core.CGFloat {
	ret := C.NSView_inst_FrameCenterRotation(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// SetFrameCenterRotation returns the rotation angle of the view around the center of its layer.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483367-framecenterrotation?language=objc for details.
func (x gen_NSView) SetFrameCenterRotation(
	value core.CGFloat,
) {
	C.NSView_inst_SetFrameCenterRotation(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// BackgroundFilters an array of Core Image filters to apply to the view’s background.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483689-backgroundfilters?language=objc for details.
func (x gen_NSView) BackgroundFilters() core.NSArray {
	ret := C.NSView_inst_BackgroundFilters(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// SetBackgroundFilters an array of Core Image filters to apply to the view’s background.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483689-backgroundfilters?language=objc for details.
func (x gen_NSView) SetBackgroundFilters(
	value core.NSArrayRef,
) {
	C.NSView_inst_SetBackgroundFilters(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// ContentFilters an array of Core Image filters to apply to the contents of the view and its sublayers.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483703-contentfilters?language=objc for details.
func (x gen_NSView) ContentFilters() core.NSArray {
	ret := C.NSView_inst_ContentFilters(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// SetContentFilters an array of Core Image filters to apply to the contents of the view and its sublayers.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483703-contentfilters?language=objc for details.
func (x gen_NSView) SetContentFilters(
	value core.NSArrayRef,
) {
	C.NSView_inst_SetContentFilters(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// CanDrawConcurrently returns a Boolean value indicating whether the view can draw its contents on a background thread.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483425-candrawconcurrently?language=objc for details.
func (x gen_NSView) CanDrawConcurrently() bool {
	ret := C.NSView_inst_CanDrawConcurrently(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetCanDrawConcurrently returns a Boolean value indicating whether the view can draw its contents on a background thread.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483425-candrawconcurrently?language=objc for details.
func (x gen_NSView) SetCanDrawConcurrently(
	value bool,
) {
	C.NSView_inst_SetCanDrawConcurrently(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// VisibleRect returns the portion of the view that is not clipped by its superviews.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483446-visiblerect?language=objc for details.
func (x gen_NSView) VisibleRect() core.NSRect {
	ret := C.NSView_inst_VisibleRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// WantsDefaultClipping returns a Boolean value indicating whether AppKit’s default clipping behavior is in effect.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483365-wantsdefaultclipping?language=objc for details.
func (x gen_NSView) WantsDefaultClipping() bool {
	ret := C.NSView_inst_WantsDefaultClipping(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// PrintJobTitle returns the view’s print job title.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483753-printjobtitle?language=objc for details.
func (x gen_NSView) PrintJobTitle() core.NSString {
	ret := C.NSView_inst_PrintJobTitle(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// PageHeader returns a default header string that includes the print job title and date.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483674-pageheader?language=objc for details.
func (x gen_NSView) PageHeader() core.NSAttributedString {
	ret := C.NSView_inst_PageHeader(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSAttributedString_FromPointer(ret)
}

// PageFooter returns a default footer string that includes the current page number and page count.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483355-pagefooter?language=objc for details.
func (x gen_NSView) PageFooter() core.NSAttributedString {
	ret := C.NSView_inst_PageFooter(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSAttributedString_FromPointer(ret)
}

// HeightAdjustLimit returns the fraction of the page that can be pushed onto the next page during automatic pagination to prevent items such as lines of text from being divided across pages.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483691-heightadjustlimit?language=objc for details.
func (x gen_NSView) HeightAdjustLimit() core.CGFloat {
	ret := C.NSView_inst_HeightAdjustLimit(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// WidthAdjustLimit returns the fraction of the page that can be pushed onto the next page during automatic pagination to prevent items such as small images or text columns from being divided across pages.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483392-widthadjustlimit?language=objc for details.
func (x gen_NSView) WidthAdjustLimit() core.CGFloat {
	ret := C.NSView_inst_WidthAdjustLimit(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// NeedsDisplay returns a Boolean value that determines whether the view needs to be redrawn before being displayed.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483360-needsdisplay?language=objc for details.
func (x gen_NSView) NeedsDisplay() bool {
	ret := C.NSView_inst_NeedsDisplay(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetNeedsDisplay returns a Boolean value that determines whether the view needs to be redrawn before being displayed.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483360-needsdisplay?language=objc for details.
func (x gen_NSView) SetNeedsDisplay(
	value bool,
) {
	C.NSView_inst_SetNeedsDisplay(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsOpaque returns a Boolean value indicating whether the view fills its frame rectangle with opaque content.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483558-opaque?language=objc for details.
func (x gen_NSView) IsOpaque() bool {
	ret := C.NSView_inst_IsOpaque(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsFlipped returns a Boolean value indicating whether the view uses a flipped coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483532-flipped?language=objc for details.
func (x gen_NSView) IsFlipped() bool {
	ret := C.NSView_inst_IsFlipped(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsRotatedFromBase returns a Boolean value indicating whether the view or any of its ancestors has ever had a rotation factor applied to its frame or bounds.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483709-rotatedfrombase?language=objc for details.
func (x gen_NSView) IsRotatedFromBase() bool {
	ret := C.NSView_inst_IsRotatedFromBase(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsRotatedOrScaledFromBase returns a Boolean value indicating whether the view or any of its ancestors has ever had a rotation factor applied to its frame or bounds, or has been scaled from the window’s base coordinate system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483390-rotatedorscaledfrombase?language=objc for details.
func (x gen_NSView) IsRotatedOrScaledFromBase() bool {
	ret := C.NSView_inst_IsRotatedOrScaledFromBase(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AutoresizesSubviews returns a Boolean value indicating whether the view applies the autoresizing behavior to its subviews when its frame size changes.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483358-autoresizessubviews?language=objc for details.
func (x gen_NSView) AutoresizesSubviews() bool {
	ret := C.NSView_inst_AutoresizesSubviews(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAutoresizesSubviews returns a Boolean value indicating whether the view applies the autoresizing behavior to its subviews when its frame size changes.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483358-autoresizessubviews?language=objc for details.
func (x gen_NSView) SetAutoresizesSubviews(
	value bool,
) {
	C.NSView_inst_SetAutoresizesSubviews(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// Constraints returns the constraints held by the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526917-constraints?language=objc for details.
func (x gen_NSView) Constraints() core.NSArray {
	ret := C.NSView_inst_Constraints(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// LayoutGuides returns the array of layout guide objects owned by this view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1534395-layoutguides?language=objc for details.
func (x gen_NSView) LayoutGuides() core.NSArray {
	ret := C.NSView_inst_LayoutGuides(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// FittingSize returns the minimum size of the view that satisfies the constraints it holds.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526904-fittingsize?language=objc for details.
func (x gen_NSView) FittingSize() core.NSSize {
	ret := C.NSView_inst_FittingSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// IntrinsicContentSize returns the natural size for the receiving view, considering only properties of the view itself.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526996-intrinsiccontentsize?language=objc for details.
func (x gen_NSView) IntrinsicContentSize() core.NSSize {
	ret := C.NSView_inst_IntrinsicContentSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))
}

// BaselineOffsetFromBottom returns the distance (in points) between the bottom of the view’s alignment rectangle and its baseline.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526949-baselineoffsetfrombottom?language=objc for details.
func (x gen_NSView) BaselineOffsetFromBottom() core.CGFloat {
	ret := C.NSView_inst_BaselineOffsetFromBottom(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// FirstBaselineOffsetFromTop returns the distance (in points) between the top of the view’s alignment rectangle and its topmost baseline.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526963-firstbaselineoffsetfromtop?language=objc for details.
func (x gen_NSView) FirstBaselineOffsetFromTop() core.CGFloat {
	ret := C.NSView_inst_FirstBaselineOffsetFromTop(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// LastBaselineOffsetFromBottom returns the distance (in points) between the bottom of the view’s alignment rectangle and its bottommost baseline.
//
// See https://developer.apple.com/documentation/appkit/nsview/1525942-lastbaselineoffsetfrombottom?language=objc for details.
func (x gen_NSView) LastBaselineOffsetFromBottom() core.CGFloat {
	ret := C.NSView_inst_LastBaselineOffsetFromBottom(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// NeedsLayout returns a Boolean value indicating whether the view needs a layout pass before it can be drawn.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526912-needslayout?language=objc for details.
func (x gen_NSView) NeedsLayout() bool {
	ret := C.NSView_inst_NeedsLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetNeedsLayout returns a Boolean value indicating whether the view needs a layout pass before it can be drawn.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526912-needslayout?language=objc for details.
func (x gen_NSView) SetNeedsLayout(
	value bool,
) {
	C.NSView_inst_SetNeedsLayout(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// NeedsUpdateConstraints returns a Boolean value indicating whether the view’s constraints need to be updated.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526856-needsupdateconstraints?language=objc for details.
func (x gen_NSView) NeedsUpdateConstraints() bool {
	ret := C.NSView_inst_NeedsUpdateConstraints(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetNeedsUpdateConstraints returns a Boolean value indicating whether the view’s constraints need to be updated.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526856-needsupdateconstraints?language=objc for details.
func (x gen_NSView) SetNeedsUpdateConstraints(
	value bool,
) {
	C.NSView_inst_SetNeedsUpdateConstraints(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// TranslatesAutoresizingMaskIntoConstraints returns a Boolean value indicating whether the view’s autoresizing mask is translated into constraints for the constraint-based layout system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526961-translatesautoresizingmaskintoco?language=objc for details.
func (x gen_NSView) TranslatesAutoresizingMaskIntoConstraints() bool {
	ret := C.NSView_inst_TranslatesAutoresizingMaskIntoConstraints(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetTranslatesAutoresizingMaskIntoConstraints returns a Boolean value indicating whether the view’s autoresizing mask is translated into constraints for the constraint-based layout system.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526961-translatesautoresizingmaskintoco?language=objc for details.
func (x gen_NSView) SetTranslatesAutoresizingMaskIntoConstraints(
	value bool,
) {
	C.NSView_inst_SetTranslatesAutoresizingMaskIntoConstraints(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// HasAmbiguousLayout returns a Boolean value indicating whether the constraints impacting the layout of the view incompletely specify the location of the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1526907-hasambiguouslayout?language=objc for details.
func (x gen_NSView) HasAmbiguousLayout() bool {
	ret := C.NSView_inst_HasAmbiguousLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// FocusRingMaskBounds returns the focus ring mask bounds, specified in the view’s coordinate space.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483287-focusringmaskbounds?language=objc for details.
func (x gen_NSView) FocusRingMaskBounds() core.NSRect {
	ret := C.NSView_inst_FocusRingMaskBounds(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// AllowsVibrancy returns a Boolean value indicating whether the view ensures it is vibrant on top of other content.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483793-allowsvibrancy?language=objc for details.
func (x gen_NSView) AllowsVibrancy() bool {
	ret := C.NSView_inst_AllowsVibrancy(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsInFullScreenMode returns a Boolean value indicating whether the view is in full screen mode.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483337-infullscreenmode?language=objc for details.
func (x gen_NSView) IsInFullScreenMode() bool {
	ret := C.NSView_inst_IsInFullScreenMode(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsHidden returns a Boolean value indicating whether the view is hidden.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483369-hidden?language=objc for details.
func (x gen_NSView) IsHidden() bool {
	ret := C.NSView_inst_IsHidden(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetHidden returns a Boolean value indicating whether the view is hidden.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483369-hidden?language=objc for details.
func (x gen_NSView) SetHidden(
	value bool,
) {
	C.NSView_inst_SetHidden(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsHiddenOrHasHiddenAncestor returns a Boolean value indicating whether the view is hidden from sight because it, or one of its ancestors, is marked as hidden.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483473-hiddenorhashiddenancestor?language=objc for details.
func (x gen_NSView) IsHiddenOrHasHiddenAncestor() bool {
	ret := C.NSView_inst_IsHiddenOrHasHiddenAncestor(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// InLiveResize returns a Boolean value indicating whether the view is being rendered as part of a live resizing operation.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483267-inliveresize?language=objc for details.
func (x gen_NSView) InLiveResize() bool {
	ret := C.NSView_inst_InLiveResize(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// PreservesContentDuringLiveResize returns a Boolean value indicating whether the view optimizes live-resize operations by preserving content that has not moved.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483795-preservescontentduringliveresize?language=objc for details.
func (x gen_NSView) PreservesContentDuringLiveResize() bool {
	ret := C.NSView_inst_PreservesContentDuringLiveResize(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// RectPreservedDuringLiveResize returns the rectangle identifying the portion of your view that did not change during a live resize operation.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483528-rectpreservedduringliveresize?language=objc for details.
func (x gen_NSView) RectPreservedDuringLiveResize() core.NSRect {
	ret := C.NSView_inst_RectPreservedDuringLiveResize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// GestureRecognizers returns the gesture recognize objects currently attached to the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483658-gesturerecognizers?language=objc for details.
func (x gen_NSView) GestureRecognizers() core.NSArray {
	ret := C.NSView_inst_GestureRecognizers(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// SetGestureRecognizers returns the gesture recognize objects currently attached to the view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483658-gesturerecognizers?language=objc for details.
func (x gen_NSView) SetGestureRecognizers(
	value core.NSArrayRef,
) {
	C.NSView_inst_SetGestureRecognizers(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// MouseDownCanMoveWindow returns a Boolean value indicating whether the view can pass mouse down events through to its superviews.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483666-mousedowncanmovewindow?language=objc for details.
func (x gen_NSView) MouseDownCanMoveWindow() bool {
	ret := C.NSView_inst_MouseDownCanMoveWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// WantsRestingTouches returns a Boolean value indicating whether the view wants resting touches.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483594-wantsrestingtouches?language=objc for details.
func (x gen_NSView) WantsRestingTouches() bool {
	ret := C.NSView_inst_WantsRestingTouches(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetWantsRestingTouches returns a Boolean value indicating whether the view wants resting touches.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483594-wantsrestingtouches?language=objc for details.
func (x gen_NSView) SetWantsRestingTouches(
	value bool,
) {
	C.NSView_inst_SetWantsRestingTouches(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// CanBecomeKeyView returns a Boolean value indicating whether the view can become key view.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483759-canbecomekeyview?language=objc for details.
func (x gen_NSView) CanBecomeKeyView() bool {
	ret := C.NSView_inst_CanBecomeKeyView(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// NeedsPanelToBecomeKey returns a Boolean value indicating whether the view needs its panel to become the key window before it can handle keyboard input and navigation.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483512-needspaneltobecomekey?language=objc for details.
func (x gen_NSView) NeedsPanelToBecomeKey() bool {
	ret := C.NSView_inst_NeedsPanelToBecomeKey(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// NextKeyView returns the view object that follows the current view in the key view loop.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483465-nextkeyview?language=objc for details.
func (x gen_NSView) NextKeyView() NSView {
	ret := C.NSView_inst_NextKeyView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_FromPointer(ret)
}

// SetNextKeyView returns the view object that follows the current view in the key view loop.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483465-nextkeyview?language=objc for details.
func (x gen_NSView) SetNextKeyView(
	value NSViewRef,
) {
	C.NSView_inst_SetNextKeyView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// NextValidKeyView returns the closest view object in the key view loop that follows the current view in the key view loop and accepts first responder status.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483572-nextvalidkeyview?language=objc for details.
func (x gen_NSView) NextValidKeyView() NSView {
	ret := C.NSView_inst_NextValidKeyView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_FromPointer(ret)
}

// PreviousKeyView returns the view object preceding the current view in the key view loop.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483646-previouskeyview?language=objc for details.
func (x gen_NSView) PreviousKeyView() NSView {
	ret := C.NSView_inst_PreviousKeyView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_FromPointer(ret)
}

// PreviousValidKeyView returns the closest view object in the key view loop that precedes the current view and accepts first responder status.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483371-previousvalidkeyview?language=objc for details.
func (x gen_NSView) PreviousValidKeyView() NSView {
	ret := C.NSView_inst_PreviousValidKeyView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_FromPointer(ret)
}

// PreparedContentRect returns the portion of the view that has been rendered and is available for responsive scrolling.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483215-preparedcontentrect?language=objc for details.
func (x gen_NSView) PreparedContentRect() core.NSRect {
	ret := C.NSView_inst_PreparedContentRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))
}

// SetPreparedContentRect returns the portion of the view that has been rendered and is available for responsive scrolling.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483215-preparedcontentrect?language=objc for details.
func (x gen_NSView) SetPreparedContentRect(
	value core.NSRect,
) {
	C.NSView_inst_SetPreparedContentRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return
}

// RegisteredDraggedTypes returns the array of pasteboard drag types that the view can accept.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483564-registereddraggedtypes?language=objc for details.
func (x gen_NSView) RegisteredDraggedTypes() core.NSArray {
	ret := C.NSView_inst_RegisteredDraggedTypes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// PostsFrameChangedNotifications returns a Boolean value indicating whether the view posts notifications when its frame rectangle changes.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483524-postsframechangednotifications?language=objc for details.
func (x gen_NSView) PostsFrameChangedNotifications() bool {
	ret := C.NSView_inst_PostsFrameChangedNotifications(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetPostsFrameChangedNotifications returns a Boolean value indicating whether the view posts notifications when its frame rectangle changes.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483524-postsframechangednotifications?language=objc for details.
func (x gen_NSView) SetPostsFrameChangedNotifications(
	value bool,
) {
	C.NSView_inst_SetPostsFrameChangedNotifications(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// PostsBoundsChangedNotifications returns a Boolean value indicating whether the view posts notifications when its bounds rectangle changes.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483239-postsboundschangednotifications?language=objc for details.
func (x gen_NSView) PostsBoundsChangedNotifications() bool {
	ret := C.NSView_inst_PostsBoundsChangedNotifications(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetPostsBoundsChangedNotifications returns a Boolean value indicating whether the view posts notifications when its bounds rectangle changes.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483239-postsboundschangednotifications?language=objc for details.
func (x gen_NSView) SetPostsBoundsChangedNotifications(
	value bool,
) {
	C.NSView_inst_SetPostsBoundsChangedNotifications(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// Tag returns the view’s tag, which is an integer that you use to identify the view within your app.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483248-tag?language=objc for details.
func (x gen_NSView) Tag() core.NSInteger {
	ret := C.NSView_inst_Tag(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)
}

// ToolTip returns the text for the view’s tooltip.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483541-tooltip?language=objc for details.
func (x gen_NSView) ToolTip() core.NSString {
	ret := C.NSView_inst_ToolTip(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetToolTip returns the text for the view’s tooltip.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483541-tooltip?language=objc for details.
func (x gen_NSView) SetToolTip(
	value core.NSStringRef,
) {
	C.NSView_inst_SetToolTip(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// TrackingAreas an array of the view’s tracking areas.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483333-trackingareas?language=objc for details.
func (x gen_NSView) TrackingAreas() core.NSArray {
	ret := C.NSView_inst_TrackingAreas(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

// IsDrawingFindIndicator returns a Boolean value indicating whether the view or one of its ancestors is being drawn for a find indicator.
//
// See https://developer.apple.com/documentation/appkit/nsview/1483317-drawingfindindicator?language=objc for details.
func (x gen_NSView) IsDrawingFindIndicator() bool {
	ret := C.NSView_inst_IsDrawingFindIndicator(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsHorizontalContentSizeConstraintActive is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsview/3353053-horizontalcontentsizeconstrainta?language=objc for details.
func (x gen_NSView) IsHorizontalContentSizeConstraintActive() bool {
	ret := C.NSView_inst_IsHorizontalContentSizeConstraintActive(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetHorizontalContentSizeConstraintActive is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsview/3353053-horizontalcontentsizeconstrainta?language=objc for details.
func (x gen_NSView) SetHorizontalContentSizeConstraintActive(
	value bool,
) {
	C.NSView_inst_SetHorizontalContentSizeConstraintActive(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsVerticalContentSizeConstraintActive is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsview/3353054-verticalcontentsizeconstraintact?language=objc for details.
func (x gen_NSView) IsVerticalContentSizeConstraintActive() bool {
	ret := C.NSView_inst_IsVerticalContentSizeConstraintActive(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetVerticalContentSizeConstraintActive is undocumented.
//
// See https://developer.apple.com/documentation/appkit/nsview/3353054-verticalcontentsizeconstraintact?language=objc for details.
func (x gen_NSView) SetVerticalContentSizeConstraintActive(
	value bool,
) {
	C.NSView_inst_SetVerticalContentSizeConstraintActive(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// BackgroundColor is undocumented.
func (x gen_NSView) BackgroundColor() NSColor {
	ret := C.NSView_inst_BackgroundColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_FromPointer(ret)
}

// SetBackgroundColor is undocumented.
func (x gen_NSView) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSView_inst_SetBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}
