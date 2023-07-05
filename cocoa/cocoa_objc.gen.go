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


void* NSBundle_type_alloc() {
	return [NSBundle
		alloc];
}
void* NSBundle_type_bundleWithURL(void* url) {
	return [NSBundle
		bundleWithURL: url];
}
void* NSBundle_type_bundleWithPath(void* path) {
	return [NSBundle
		bundleWithPath: path];
}
void* NSBundle_type_bundleWithIdentifier(void* identifier) {
	return [NSBundle
		bundleWithIdentifier: identifier];
}
void* NSBundle_type_URLForResource_withExtension_subdirectory_inBundleWithURL(void* name, void* ext, void* subpath, void* bundleURL) {
	return [NSBundle
		URLForResource: name
		withExtension: ext
		subdirectory: subpath
		inBundleWithURL: bundleURL];
}
void* NSBundle_type_URLsForResourcesWithExtension_subdirectory_inBundleWithURL(void* ext, void* subpath, void* bundleURL) {
	return [NSBundle
		URLsForResourcesWithExtension: ext
		subdirectory: subpath
		inBundleWithURL: bundleURL];
}
void* NSBundle_type_pathForResource_ofType_inDirectory(void* name, void* ext, void* bundlePath) {
	return [NSBundle
		pathForResource: name
		ofType: ext
		inDirectory: bundlePath];
}
void* NSBundle_type_pathsForResourcesOfType_inDirectory(void* ext, void* bundlePath) {
	return [NSBundle
		pathsForResourcesOfType: ext
		inDirectory: bundlePath];
}
void* NSBundle_type_preferredLocalizationsFromArray(void* localizationsArray) {
	return [NSBundle
		preferredLocalizationsFromArray: localizationsArray];
}
void* NSBundle_type_preferredLocalizationsFromArray_forPreferences(void* localizationsArray, void* preferencesArray) {
	return [NSBundle
		preferredLocalizationsFromArray: localizationsArray
		forPreferences: preferencesArray];
}
void* NSBundle_type_mainBundle() {
	return [NSBundle
		mainBundle];
}
void* NSBundle_type_allFrameworks() {
	return [NSBundle
		allFrameworks];
}
void* NSBundle_type_allBundles() {
	return [NSBundle
		allBundles];
}
void* NSSound_type_alloc() {
	return [NSSound
		alloc];
}
BOOL NSSound_type_canInitWithPasteboard(void* pasteboard) {
	return [NSSound
		canInitWithPasteboard: pasteboard];
}
void* NSSound_type_soundUnfilteredTypes() {
	return [NSSound
		soundUnfilteredTypes];
}
void* NSApplication_type_alloc() {
	return [NSApplication
		alloc];
}
void NSApplication_type_detachDrawingThread_toTarget_withObject(void* selector, void* target, void* argument) {
	[NSApplication
		detachDrawingThread: selector
		toTarget: target
		withObject: argument];
}
void* NSApplication_type_sharedApplication() {
	return [NSApplication
		sharedApplication];
}
void* NSControl_type_alloc() {
	return [NSControl
		alloc];
}
void* NSButton_type_alloc() {
	return [NSButton
		alloc];
}
void* NSButton_type_checkboxWithTitle_target_action(void* title, void* target, void* action) {
	return [NSButton
		checkboxWithTitle: title
		target: target
		action: action];
}
void* NSButton_type_buttonWithImage_target_action(void* image, void* target, void* action) {
	return [NSButton
		buttonWithImage: image
		target: target
		action: action];
}
void* NSButton_type_radioButtonWithTitle_target_action(void* title, void* target, void* action) {
	return [NSButton
		radioButtonWithTitle: title
		target: target
		action: action];
}
void* NSButton_type_buttonWithTitle_image_target_action(void* title, void* image, void* target, void* action) {
	return [NSButton
		buttonWithTitle: title
		image: image
		target: target
		action: action];
}
void* NSButton_type_buttonWithTitle_target_action(void* title, void* target, void* action) {
	return [NSButton
		buttonWithTitle: title
		target: target
		action: action];
}
void* NSEvent_type_alloc() {
	return [NSEvent
		alloc];
}
void* NSEvent_type_eventWithEventRef(void* eventRef) {
	return [NSEvent
		eventWithEventRef: eventRef];
}
void NSEvent_type_stopPeriodicEvents() {
	[NSEvent
		stopPeriodicEvents];
}
void NSEvent_type_removeMonitor(void* eventMonitor) {
	[NSEvent
		removeMonitor: eventMonitor];
}
unsigned long NSEvent_type_pressedMouseButtons() {
	return [NSEvent
		pressedMouseButtons];
}
NSPoint NSEvent_type_mouseLocation() {
	return [NSEvent
		mouseLocation];
}
BOOL NSEvent_type_mouseCoalescingEnabled() {
	return [NSEvent
		mouseCoalescingEnabled];
}
void NSEvent_type_setMouseCoalescingEnabled(BOOL value) {
	[NSEvent
		setMouseCoalescingEnabled: value];
}
BOOL NSEvent_type_swipeTrackingFromScrollEventsEnabled() {
	return [NSEvent
		swipeTrackingFromScrollEventsEnabled];
}
void* NSFont_type_alloc() {
	return [NSFont
		alloc];
}
void* NSFont_type_fontWithName_size(void* fontName, double fontSize) {
	return [NSFont
		fontWithName: fontName
		size: fontSize];
}
void* NSFont_type_userFontOfSize(double fontSize) {
	return [NSFont
		userFontOfSize: fontSize];
}
void* NSFont_type_userFixedPitchFontOfSize(double fontSize) {
	return [NSFont
		userFixedPitchFontOfSize: fontSize];
}
void* NSFont_type_systemFontOfSize(double fontSize) {
	return [NSFont
		systemFontOfSize: fontSize];
}
void* NSFont_type_boldSystemFontOfSize(double fontSize) {
	return [NSFont
		boldSystemFontOfSize: fontSize];
}
void* NSFont_type_labelFontOfSize(double fontSize) {
	return [NSFont
		labelFontOfSize: fontSize];
}
void* NSFont_type_messageFontOfSize(double fontSize) {
	return [NSFont
		messageFontOfSize: fontSize];
}
void* NSFont_type_menuBarFontOfSize(double fontSize) {
	return [NSFont
		menuBarFontOfSize: fontSize];
}
void* NSFont_type_menuFontOfSize(double fontSize) {
	return [NSFont
		menuFontOfSize: fontSize];
}
void* NSFont_type_controlContentFontOfSize(double fontSize) {
	return [NSFont
		controlContentFontOfSize: fontSize];
}
void* NSFont_type_titleBarFontOfSize(double fontSize) {
	return [NSFont
		titleBarFontOfSize: fontSize];
}
void* NSFont_type_paletteFontOfSize(double fontSize) {
	return [NSFont
		paletteFontOfSize: fontSize];
}
void* NSFont_type_toolTipsFontOfSize(double fontSize) {
	return [NSFont
		toolTipsFontOfSize: fontSize];
}
void NSFont_type_setUserFont(void* font) {
	[NSFont
		setUserFont: font];
}
void NSFont_type_setUserFixedPitchFont(void* font) {
	[NSFont
		setUserFixedPitchFont: font];
}
double NSFont_type_systemFontSize() {
	return [NSFont
		systemFontSize];
}
double NSFont_type_smallSystemFontSize() {
	return [NSFont
		smallSystemFontSize];
}
double NSFont_type_labelFontSize() {
	return [NSFont
		labelFontSize];
}
void* NSImage_type_alloc() {
	return [NSImage
		alloc];
}
void* NSImage_type_imageWithSystemSymbolName_accessibilityDescription(void* symbolName, void* description) {
	return [NSImage
		imageWithSystemSymbolName: symbolName
		accessibilityDescription: description];
}
BOOL NSImage_type_canInitWithPasteboard(void* pasteboard) {
	return [NSImage
		canInitWithPasteboard: pasteboard];
}
void* NSImage_type_imageTypes() {
	return [NSImage
		imageTypes];
}
void* NSImage_type_imageUnfilteredTypes() {
	return [NSImage
		imageUnfilteredTypes];
}
void* NSImageView_type_alloc() {
	return [NSImageView
		alloc];
}
void* NSImageView_type_imageViewWithImage(void* image) {
	return [NSImageView
		imageViewWithImage: image];
}
void* NSNib_type_alloc() {
	return [NSNib
		alloc];
}
void* NSPasteboard_type_alloc() {
	return [NSPasteboard
		alloc];
}
void* NSPasteboard_type_pasteboardByFilteringFile(void* filename) {
	return [NSPasteboard
		pasteboardByFilteringFile: filename];
}
void* NSPasteboard_type_pasteboardByFilteringTypesInPasteboard(void* pboard) {
	return [NSPasteboard
		pasteboardByFilteringTypesInPasteboard: pboard];
}
void* NSPasteboard_type_pasteboardWithUniqueName() {
	return [NSPasteboard
		pasteboardWithUniqueName];
}
void* NSPasteboard_type_generalPasteboard() {
	return [NSPasteboard
		generalPasteboard];
}
void* NSLayoutManager_type_alloc() {
	return [NSLayoutManager
		alloc];
}
void* NSMenu_type_alloc() {
	return [NSMenu
		alloc];
}
BOOL NSMenu_type_menuBarVisible() {
	return [NSMenu
		menuBarVisible];
}
void NSMenu_type_setMenuBarVisible(BOOL visible) {
	[NSMenu
		setMenuBarVisible: visible];
}
void NSMenu_type_popUpContextMenu_withEvent_forView(void* menu, void* event, void* view) {
	[NSMenu
		popUpContextMenu: menu
		withEvent: event
		forView: view];
}
void NSMenu_type_popUpContextMenu_withEvent_forView_withFont(void* menu, void* event, void* view, void* font) {
	[NSMenu
		popUpContextMenu: menu
		withEvent: event
		forView: view
		withFont: font];
}
void* NSPopover_type_alloc() {
	return [NSPopover
		alloc];
}
void* NSMenuItem_type_alloc() {
	return [NSMenuItem
		alloc];
}
void* NSMenuItem_type_separatorItem() {
	return [NSMenuItem
		separatorItem];
}
BOOL NSMenuItem_type_usesUserKeyEquivalents() {
	return [NSMenuItem
		usesUserKeyEquivalents];
}
void NSMenuItem_type_setUsesUserKeyEquivalents(BOOL value) {
	[NSMenuItem
		setUsesUserKeyEquivalents: value];
}
void* NSRunningApplication_type_alloc() {
	return [NSRunningApplication
		alloc];
}
void* NSRunningApplication_type_runningApplicationsWithBundleIdentifier(void* bundleIdentifier) {
	return [NSRunningApplication
		runningApplicationsWithBundleIdentifier: bundleIdentifier];
}
void NSRunningApplication_type_terminateAutomaticallyTerminableApplications() {
	[NSRunningApplication
		terminateAutomaticallyTerminableApplications];
}
void* NSRunningApplication_type_currentApplication() {
	return [NSRunningApplication
		currentApplication];
}
void* NSScreen_type_alloc() {
	return [NSScreen
		alloc];
}
void* NSScreen_type_mainScreen() {
	return [NSScreen
		mainScreen];
}
void* NSScreen_type_deepestScreen() {
	return [NSScreen
		deepestScreen];
}
void* NSScreen_type_screens() {
	return [NSScreen
		screens];
}
BOOL NSScreen_type_screensHaveSeparateSpaces() {
	return [NSScreen
		screensHaveSeparateSpaces];
}
void* NSStatusBar_type_alloc() {
	return [NSStatusBar
		alloc];
}
void* NSStatusBar_type_systemStatusBar() {
	return [NSStatusBar
		systemStatusBar];
}
void* NSStatusBarButton_type_alloc() {
	return [NSStatusBarButton
		alloc];
}
void* NSStatusItem_type_alloc() {
	return [NSStatusItem
		alloc];
}
void* NSText_type_alloc() {
	return [NSText
		alloc];
}
void* NSTextField_type_alloc() {
	return [NSTextField
		alloc];
}
void* NSTextField_type_labelWithAttributedString(void* attributedStringValue) {
	return [NSTextField
		labelWithAttributedString: attributedStringValue];
}
void* NSTextField_type_labelWithString(void* stringValue) {
	return [NSTextField
		labelWithString: stringValue];
}
void* NSTextField_type_textFieldWithString(void* stringValue) {
	return [NSTextField
		textFieldWithString: stringValue];
}
void* NSTextField_type_wrappingLabelWithString(void* stringValue) {
	return [NSTextField
		wrappingLabelWithString: stringValue];
}
void* NSTextContainer_type_alloc() {
	return [NSTextContainer
		alloc];
}
void* NSViewController_type_alloc() {
	return [NSViewController
		alloc];
}
void* NSVisualEffectView_type_alloc() {
	return [NSVisualEffectView
		alloc];
}
void* NSWindow_type_alloc() {
	return [NSWindow
		alloc];
}
void* NSWindow_type_windowWithContentViewController(void* contentViewController) {
	return [NSWindow
		windowWithContentViewController: contentViewController];
}
NSRect NSWindow_type_contentRectForFrameRect_styleMask(NSRect fRect, unsigned long style) {
	return [NSWindow
		contentRectForFrameRect: fRect
		styleMask: style];
}
NSRect NSWindow_type_frameRectForContentRect_styleMask(NSRect cRect, unsigned long style) {
	return [NSWindow
		frameRectForContentRect: cRect
		styleMask: style];
}
double NSWindow_type_minFrameWidthWithTitle_styleMask(void* title, unsigned long style) {
	return [NSWindow
		minFrameWidthWithTitle: title
		styleMask: style];
}
long NSWindow_type_windowNumberAtPoint_belowWindowWithWindowNumber(NSPoint point, long windowNumber) {
	return [NSWindow
		windowNumberAtPoint: point
		belowWindowWithWindowNumber: windowNumber];
}
BOOL NSWindow_type_allowsAutomaticWindowTabbing() {
	return [NSWindow
		allowsAutomaticWindowTabbing];
}
void NSWindow_type_setAllowsAutomaticWindowTabbing(BOOL value) {
	[NSWindow
		setAllowsAutomaticWindowTabbing: value];
}
void* NSWorkspace_type_alloc() {
	return [NSWorkspace
		alloc];
}
void* NSWorkspace_type_sharedWorkspace() {
	return [NSWorkspace
		sharedWorkspace];
}
void* NSColor_type_alloc() {
	return [NSColor
		alloc];
}
void* NSColor_type_colorFromPasteboard(void* pasteBoard) {
	return [NSColor
		colorFromPasteboard: pasteBoard];
}
void* NSColor_type_colorWithRed_green_blue_alpha(double red, double green, double blue, double alpha) {
	return [NSColor
		colorWithRed: red
		green: green
		blue: blue
		alpha: alpha];
}
BOOL NSColor_type_ignoresAlpha() {
	return [NSColor
		ignoresAlpha];
}
void NSColor_type_setIgnoresAlpha(BOOL value) {
	[NSColor
		setIgnoresAlpha: value];
}
void* NSColor_type_systemCyanColor() {
	return [NSColor
		systemCyanColor];
}
void* NSColor_type_systemMintColor() {
	return [NSColor
		systemMintColor];
}
void* NSColor_type_clearColor() {
	return [NSColor
		clearColor];
}
void* NSTextView_type_alloc() {
	return [NSTextView
		alloc];
}
void NSTextView_type_registerForServices() {
	[NSTextView
		registerForServices];
}
void* NSTextView_type_fieldEditor() {
	return [NSTextView
		fieldEditor];
}
BOOL NSTextView_type_stronglyReferencesTextStorage() {
	return [NSTextView
		stronglyReferencesTextStorage];
}
void* NSView_type_alloc() {
	return [NSView
		alloc];
}
BOOL NSView_type_requiresConstraintBasedLayout() {
	return [NSView
		requiresConstraintBasedLayout];
}
void* NSView_type_focusView() {
	return [NSView
		focusView];
}
void* NSView_type_defaultMenu() {
	return [NSView
		defaultMenu];
}
BOOL NSView_type_compatibleWithResponsiveScrolling() {
	return [NSView
		compatibleWithResponsiveScrolling];
}


void* NSBundle_inst_URLForAuxiliaryExecutable(void *id, void* executableName) {
	return [(NSBundle*)id
		URLForAuxiliaryExecutable: executableName];
}

void* NSBundle_inst_URLForResource_withExtension(void *id, void* name, void* ext) {
	return [(NSBundle*)id
		URLForResource: name
		withExtension: ext];
}

void* NSBundle_inst_URLForResource_withExtension_subdirectory(void *id, void* name, void* ext, void* subpath) {
	return [(NSBundle*)id
		URLForResource: name
		withExtension: ext
		subdirectory: subpath];
}

void* NSBundle_inst_URLForResource_withExtension_subdirectory_localization(void *id, void* name, void* ext, void* subpath, void* localizationName) {
	return [(NSBundle*)id
		URLForResource: name
		withExtension: ext
		subdirectory: subpath
		localization: localizationName];
}

void* NSBundle_inst_URLsForResourcesWithExtension_subdirectory(void *id, void* ext, void* subpath) {
	return [(NSBundle*)id
		URLsForResourcesWithExtension: ext
		subdirectory: subpath];
}

void* NSBundle_inst_URLsForResourcesWithExtension_subdirectory_localization(void *id, void* ext, void* subpath, void* localizationName) {
	return [(NSBundle*)id
		URLsForResourcesWithExtension: ext
		subdirectory: subpath
		localization: localizationName];
}

void* NSBundle_inst_initWithPath(void *id, void* path) {
	return [(NSBundle*)id
		initWithPath: path];
}

void* NSBundle_inst_initWithURL(void *id, void* url) {
	return [(NSBundle*)id
		initWithURL: url];
}

BOOL NSBundle_inst_load(void *id) {
	return [(NSBundle*)id
		load];
}

void* NSBundle_inst_loadNibNamed_owner_options(void *id, void* name, void* owner, void* options) {
	return [(NSBundle*)id
		loadNibNamed: name
		owner: owner
		options: options];
}

void* NSBundle_inst_localizedAttributedStringForKey_value_table(void *id, void* key, void* value, void* tableName) {
	return [(NSBundle*)id
		localizedAttributedStringForKey: key
		value: value
		table: tableName];
}

void* NSBundle_inst_localizedStringForKey_value_table(void *id, void* key, void* value, void* tableName) {
	return [(NSBundle*)id
		localizedStringForKey: key
		value: value
		table: tableName];
}

void* NSBundle_inst_objectForInfoDictionaryKey(void *id, void* key) {
	return [(NSBundle*)id
		objectForInfoDictionaryKey: key];
}

void* NSBundle_inst_pathForAuxiliaryExecutable(void *id, void* executableName) {
	return [(NSBundle*)id
		pathForAuxiliaryExecutable: executableName];
}

void* NSBundle_inst_pathForResource_ofType(void *id, void* name, void* ext) {
	return [(NSBundle*)id
		pathForResource: name
		ofType: ext];
}

void* NSBundle_inst_pathForResource_ofType_inDirectory(void *id, void* name, void* ext, void* subpath) {
	return [(NSBundle*)id
		pathForResource: name
		ofType: ext
		inDirectory: subpath];
}

void* NSBundle_inst_pathForResource_ofType_inDirectory_forLocalization(void *id, void* name, void* ext, void* subpath, void* localizationName) {
	return [(NSBundle*)id
		pathForResource: name
		ofType: ext
		inDirectory: subpath
		forLocalization: localizationName];
}

void* NSBundle_inst_pathsForResourcesOfType_inDirectory(void *id, void* ext, void* subpath) {
	return [(NSBundle*)id
		pathsForResourcesOfType: ext
		inDirectory: subpath];
}

void* NSBundle_inst_pathsForResourcesOfType_inDirectory_forLocalization(void *id, void* ext, void* subpath, void* localizationName) {
	return [(NSBundle*)id
		pathsForResourcesOfType: ext
		inDirectory: subpath
		forLocalization: localizationName];
}

BOOL NSBundle_inst_unload(void *id) {
	return [(NSBundle*)id
		unload];
}

void* NSBundle_inst_init(void *id) {
	return [(NSBundle*)id
		init];
}

void* NSBundle_inst_resourceURL(void *id) {
	return [(NSBundle*)id
		resourceURL];
}

void* NSBundle_inst_executableURL(void *id) {
	return [(NSBundle*)id
		executableURL];
}

void* NSBundle_inst_privateFrameworksURL(void *id) {
	return [(NSBundle*)id
		privateFrameworksURL];
}

void* NSBundle_inst_sharedFrameworksURL(void *id) {
	return [(NSBundle*)id
		sharedFrameworksURL];
}

void* NSBundle_inst_builtInPlugInsURL(void *id) {
	return [(NSBundle*)id
		builtInPlugInsURL];
}

void* NSBundle_inst_sharedSupportURL(void *id) {
	return [(NSBundle*)id
		sharedSupportURL];
}

void* NSBundle_inst_appStoreReceiptURL(void *id) {
	return [(NSBundle*)id
		appStoreReceiptURL];
}

void* NSBundle_inst_resourcePath(void *id) {
	return [(NSBundle*)id
		resourcePath];
}

void* NSBundle_inst_executablePath(void *id) {
	return [(NSBundle*)id
		executablePath];
}

void* NSBundle_inst_privateFrameworksPath(void *id) {
	return [(NSBundle*)id
		privateFrameworksPath];
}

void* NSBundle_inst_sharedFrameworksPath(void *id) {
	return [(NSBundle*)id
		sharedFrameworksPath];
}

void* NSBundle_inst_builtInPlugInsPath(void *id) {
	return [(NSBundle*)id
		builtInPlugInsPath];
}

void* NSBundle_inst_sharedSupportPath(void *id) {
	return [(NSBundle*)id
		sharedSupportPath];
}

void* NSBundle_inst_bundleURL(void *id) {
	return [(NSBundle*)id
		bundleURL];
}

void* NSBundle_inst_bundlePath(void *id) {
	return [(NSBundle*)id
		bundlePath];
}

void* NSBundle_inst_bundleIdentifier(void *id) {
	return [(NSBundle*)id
		bundleIdentifier];
}

void* NSBundle_inst_infoDictionary(void *id) {
	return [(NSBundle*)id
		infoDictionary];
}

void* NSBundle_inst_localizations(void *id) {
	return [(NSBundle*)id
		localizations];
}

void* NSBundle_inst_preferredLocalizations(void *id) {
	return [(NSBundle*)id
		preferredLocalizations];
}

void* NSBundle_inst_developmentLocalization(void *id) {
	return [(NSBundle*)id
		developmentLocalization];
}

void* NSBundle_inst_localizedInfoDictionary(void *id) {
	return [(NSBundle*)id
		localizedInfoDictionary];
}

void* NSBundle_inst_executableArchitectures(void *id) {
	return [(NSBundle*)id
		executableArchitectures];
}

BOOL NSBundle_inst_isLoaded(void *id) {
	return [(NSBundle*)id
		isLoaded];
}

void* NSSound_inst_initWithContentsOfFile_byReference(void *id, void* path, BOOL byRef) {
	return [(NSSound*)id
		initWithContentsOfFile: path
		byReference: byRef];
}

void* NSSound_inst_initWithContentsOfURL_byReference(void *id, void* url, BOOL byRef) {
	return [(NSSound*)id
		initWithContentsOfURL: url
		byReference: byRef];
}

void* NSSound_inst_initWithData(void *id, void* data) {
	return [(NSSound*)id
		initWithData: data];
}

void* NSSound_inst_initWithPasteboard(void *id, void* pasteboard) {
	return [(NSSound*)id
		initWithPasteboard: pasteboard];
}

BOOL NSSound_inst_pause(void *id) {
	return [(NSSound*)id
		pause];
}

BOOL NSSound_inst_play(void *id) {
	return [(NSSound*)id
		play];
}

BOOL NSSound_inst_resume(void *id) {
	return [(NSSound*)id
		resume];
}

BOOL NSSound_inst_stop(void *id) {
	return [(NSSound*)id
		stop];
}

void NSSound_inst_writeToPasteboard(void *id, void* pasteboard) {
	[(NSSound*)id
		writeToPasteboard: pasteboard];
}

void* NSSound_inst_init(void *id) {
	return [(NSSound*)id
		init];
}

void* NSSound_inst_delegate(void *id) {
	return [(NSSound*)id
		delegate];
}

void NSSound_inst_setDelegate(void *id, void* value) {
	[(NSSound*)id
		setDelegate: value];
}

BOOL NSSound_inst_loops(void *id) {
	return [(NSSound*)id
		loops];
}

void NSSound_inst_setLoops(void *id, BOOL value) {
	[(NSSound*)id
		setLoops: value];
}

BOOL NSSound_inst_isPlaying(void *id) {
	return [(NSSound*)id
		isPlaying];
}

void NSApplication_inst_activateContextHelpMode(void *id, void* sender) {
	[(NSApplication*)id
		activateContextHelpMode: sender];
}

void NSApplication_inst_activateIgnoringOtherApps(void *id, BOOL flag) {
	[(NSApplication*)id
		activateIgnoringOtherApps: flag];
}

long NSApplication_inst_activationPolicy(void *id) {
	return [(NSApplication*)id
		activationPolicy];
}

void NSApplication_inst_cancelUserAttentionRequest(void *id, long request) {
	[(NSApplication*)id
		cancelUserAttentionRequest: request];
}

void NSApplication_inst_deactivate(void *id) {
	[(NSApplication*)id
		deactivate];
}

void NSApplication_inst_disableRelaunchOnLogin(void *id) {
	[(NSApplication*)id
		disableRelaunchOnLogin];
}

void NSApplication_inst_enableRelaunchOnLogin(void *id) {
	[(NSApplication*)id
		enableRelaunchOnLogin];
}

void NSApplication_inst_finishLaunching(void *id) {
	[(NSApplication*)id
		finishLaunching];
}

void NSApplication_inst_hideOtherApplications(void *id, void* sender) {
	[(NSApplication*)id
		hideOtherApplications: sender];
}

void NSApplication_inst_postEvent_atStart(void *id, void* event, BOOL flag) {
	[(NSApplication*)id
		postEvent: event
		atStart: flag];
}

void NSApplication_inst_registerForRemoteNotifications(void *id) {
	[(NSApplication*)id
		registerForRemoteNotifications];
}

void NSApplication_inst_registerUserInterfaceItemSearchHandler(void *id, void* handler) {
	[(NSApplication*)id
		registerUserInterfaceItemSearchHandler: handler];
}

void NSApplication_inst_replyToApplicationShouldTerminate(void *id, BOOL shouldTerminate) {
	[(NSApplication*)id
		replyToApplicationShouldTerminate: shouldTerminate];
}

void NSApplication_inst_run(void *id) {
	[(NSApplication*)id
		run];
}

BOOL NSApplication_inst_sendAction_to_from(void *id, void* action, void* target, void* sender) {
	return [(NSApplication*)id
		sendAction: action
		to: target
		from: sender];
}

void NSApplication_inst_sendEvent(void *id, void* event) {
	[(NSApplication*)id
		sendEvent: event];
}

BOOL NSApplication_inst_setActivationPolicy(void *id, long activationPolicy) {
	return [(NSApplication*)id
		setActivationPolicy: activationPolicy];
}

void NSApplication_inst_showHelp(void *id, void* sender) {
	[(NSApplication*)id
		showHelp: sender];
}

void NSApplication_inst_stop(void *id, void* sender) {
	[(NSApplication*)id
		stop: sender];
}

void* NSApplication_inst_targetForAction(void *id, void* action) {
	return [(NSApplication*)id
		targetForAction: action];
}

void* NSApplication_inst_targetForAction_to_from(void *id, void* action, void* target, void* sender) {
	return [(NSApplication*)id
		targetForAction: action
		to: target
		from: sender];
}

void NSApplication_inst_terminate(void *id, void* sender) {
	[(NSApplication*)id
		terminate: sender];
}

void NSApplication_inst_toggleTouchBarCustomizationPalette(void *id, void* sender) {
	[(NSApplication*)id
		toggleTouchBarCustomizationPalette: sender];
}

BOOL NSApplication_inst_tryToPerform_with(void *id, void* action, void* object) {
	return [(NSApplication*)id
		tryToPerform: action
		with: object];
}

void NSApplication_inst_unhideAllApplications(void *id, void* sender) {
	[(NSApplication*)id
		unhideAllApplications: sender];
}

void NSApplication_inst_unregisterForRemoteNotifications(void *id) {
	[(NSApplication*)id
		unregisterForRemoteNotifications];
}

void NSApplication_inst_unregisterUserInterfaceItemSearchHandler(void *id, void* handler) {
	[(NSApplication*)id
		unregisterUserInterfaceItemSearchHandler: handler];
}

void* NSApplication_inst_init(void *id) {
	return [(NSApplication*)id
		init];
}

void* NSApplication_inst_delegate(void *id) {
	return [(NSApplication*)id
		delegate];
}

void NSApplication_inst_setDelegate(void *id, void* value) {
	[(NSApplication*)id
		setDelegate: value];
}

void* NSApplication_inst_currentEvent(void *id) {
	return [(NSApplication*)id
		currentEvent];
}

BOOL NSApplication_inst_isRunning(void *id) {
	return [(NSApplication*)id
		isRunning];
}

BOOL NSApplication_inst_isActive(void *id) {
	return [(NSApplication*)id
		isActive];
}

BOOL NSApplication_inst_isRegisteredForRemoteNotifications(void *id) {
	return [(NSApplication*)id
		isRegisteredForRemoteNotifications];
}

void* NSApplication_inst_applicationIconImage(void *id) {
	return [(NSApplication*)id
		applicationIconImage];
}

void NSApplication_inst_setApplicationIconImage(void *id, void* value) {
	[(NSApplication*)id
		setApplicationIconImage: value];
}

void* NSApplication_inst_helpMenu(void *id) {
	return [(NSApplication*)id
		helpMenu];
}

void NSApplication_inst_setHelpMenu(void *id, void* value) {
	[(NSApplication*)id
		setHelpMenu: value];
}

void* NSApplication_inst_servicesProvider(void *id) {
	return [(NSApplication*)id
		servicesProvider];
}

void NSApplication_inst_setServicesProvider(void *id, void* value) {
	[(NSApplication*)id
		setServicesProvider: value];
}

BOOL NSApplication_inst_isFullKeyboardAccessEnabled(void *id) {
	return [(NSApplication*)id
		isFullKeyboardAccessEnabled];
}

void* NSApplication_inst_orderedDocuments(void *id) {
	return [(NSApplication*)id
		orderedDocuments];
}

void* NSApplication_inst_orderedWindows(void *id) {
	return [(NSApplication*)id
		orderedWindows];
}

void* NSApplication_inst_mainMenu(void *id) {
	return [(NSApplication*)id
		mainMenu];
}

void NSApplication_inst_setMainMenu(void *id, void* value) {
	[(NSApplication*)id
		setMainMenu: value];
}

BOOL NSControl_inst_abortEditing(void *id) {
	return [(NSControl*)id
		abortEditing];
}

void* NSControl_inst_currentEditor(void *id) {
	return [(NSControl*)id
		currentEditor];
}

void NSControl_inst_drawWithExpansionFrame_inView(void *id, NSRect contentFrame, void* view) {
	[(NSControl*)id
		drawWithExpansionFrame: contentFrame
		inView: view];
}

void NSControl_inst_editWithFrame_editor_delegate_event(void *id, NSRect rect, void* textObj, void* delegate, void* event) {
	[(NSControl*)id
		editWithFrame: rect
		editor: textObj
		delegate: delegate
		event: event];
}

void NSControl_inst_endEditing(void *id, void* textObj) {
	[(NSControl*)id
		endEditing: textObj];
}

NSRect NSControl_inst_expansionFrameWithFrame(void *id, NSRect contentFrame) {
	return [(NSControl*)id
		expansionFrameWithFrame: contentFrame];
}

void* NSControl_inst_initWithFrame(void *id, NSRect frameRect) {
	return [(NSControl*)id
		initWithFrame: frameRect];
}

void NSControl_inst_mouseDown(void *id, void* event) {
	[(NSControl*)id
		mouseDown: event];
}

void NSControl_inst_performClick(void *id, void* sender) {
	[(NSControl*)id
		performClick: sender];
}

void NSControl_inst_selectWithFrame_editor_delegate_start_length(void *id, NSRect rect, void* textObj, void* delegate, long selStart, long selLength) {
	[(NSControl*)id
		selectWithFrame: rect
		editor: textObj
		delegate: delegate
		start: selStart
		length: selLength];
}

BOOL NSControl_inst_sendAction_to(void *id, void* action, void* target) {
	return [(NSControl*)id
		sendAction: action
		to: target];
}

NSSize NSControl_inst_sizeThatFits(void *id, NSSize size) {
	return [(NSControl*)id
		sizeThatFits: size];
}

void NSControl_inst_sizeToFit(void *id) {
	[(NSControl*)id
		sizeToFit];
}

void NSControl_inst_takeDoubleValueFrom(void *id, void* sender) {
	[(NSControl*)id
		takeDoubleValueFrom: sender];
}

void NSControl_inst_takeFloatValueFrom(void *id, void* sender) {
	[(NSControl*)id
		takeFloatValueFrom: sender];
}

void NSControl_inst_takeIntValueFrom(void *id, void* sender) {
	[(NSControl*)id
		takeIntValueFrom: sender];
}

void NSControl_inst_takeIntegerValueFrom(void *id, void* sender) {
	[(NSControl*)id
		takeIntegerValueFrom: sender];
}

void NSControl_inst_takeObjectValueFrom(void *id, void* sender) {
	[(NSControl*)id
		takeObjectValueFrom: sender];
}

void NSControl_inst_takeStringValueFrom(void *id, void* sender) {
	[(NSControl*)id
		takeStringValueFrom: sender];
}

void NSControl_inst_validateEditing(void *id) {
	[(NSControl*)id
		validateEditing];
}

void* NSControl_inst_init(void *id) {
	return [(NSControl*)id
		init];
}

BOOL NSControl_inst_isEnabled(void *id) {
	return [(NSControl*)id
		isEnabled];
}

void NSControl_inst_setEnabled(void *id, BOOL value) {
	[(NSControl*)id
		setEnabled: value];
}

int NSControl_inst_intValue(void *id) {
	return [(NSControl*)id
		intValue];
}

void NSControl_inst_setIntValue(void *id, int value) {
	[(NSControl*)id
		setIntValue: value];
}

long NSControl_inst_integerValue(void *id) {
	return [(NSControl*)id
		integerValue];
}

void NSControl_inst_setIntegerValue(void *id, long value) {
	[(NSControl*)id
		setIntegerValue: value];
}

void* NSControl_inst_objectValue(void *id) {
	return [(NSControl*)id
		objectValue];
}

void NSControl_inst_setObjectValue(void *id, void* value) {
	[(NSControl*)id
		setObjectValue: value];
}

void* NSControl_inst_stringValue(void *id) {
	return [(NSControl*)id
		stringValue];
}

void NSControl_inst_setStringValue(void *id, void* value) {
	[(NSControl*)id
		setStringValue: value];
}

void* NSControl_inst_attributedStringValue(void *id) {
	return [(NSControl*)id
		attributedStringValue];
}

void NSControl_inst_setAttributedStringValue(void *id, void* value) {
	[(NSControl*)id
		setAttributedStringValue: value];
}

void* NSControl_inst_font(void *id) {
	return [(NSControl*)id
		font];
}

void NSControl_inst_setFont(void *id, void* value) {
	[(NSControl*)id
		setFont: value];
}

BOOL NSControl_inst_usesSingleLineMode(void *id) {
	return [(NSControl*)id
		usesSingleLineMode];
}

void NSControl_inst_setUsesSingleLineMode(void *id, BOOL value) {
	[(NSControl*)id
		setUsesSingleLineMode: value];
}

BOOL NSControl_inst_allowsExpansionToolTips(void *id) {
	return [(NSControl*)id
		allowsExpansionToolTips];
}

void NSControl_inst_setAllowsExpansionToolTips(void *id, BOOL value) {
	[(NSControl*)id
		setAllowsExpansionToolTips: value];
}

BOOL NSControl_inst_isHighlighted(void *id) {
	return [(NSControl*)id
		isHighlighted];
}

void NSControl_inst_setHighlighted(void *id, BOOL value) {
	[(NSControl*)id
		setHighlighted: value];
}

void* NSControl_inst_action(void *id) {
	return [(NSControl*)id
		action];
}

void NSControl_inst_setAction(void *id, void* value) {
	[(NSControl*)id
		setAction: value];
}

void* NSControl_inst_target(void *id) {
	return [(NSControl*)id
		target];
}

void NSControl_inst_setTarget(void *id, void* value) {
	[(NSControl*)id
		setTarget: value];
}

BOOL NSControl_inst_isContinuous(void *id) {
	return [(NSControl*)id
		isContinuous];
}

void NSControl_inst_setContinuous(void *id, BOOL value) {
	[(NSControl*)id
		setContinuous: value];
}

long NSControl_inst_tag(void *id) {
	return [(NSControl*)id
		tag];
}

void NSControl_inst_setTag(void *id, long value) {
	[(NSControl*)id
		setTag: value];
}

BOOL NSControl_inst_refusesFirstResponder(void *id) {
	return [(NSControl*)id
		refusesFirstResponder];
}

void NSControl_inst_setRefusesFirstResponder(void *id, BOOL value) {
	[(NSControl*)id
		setRefusesFirstResponder: value];
}

BOOL NSControl_inst_ignoresMultiClick(void *id) {
	return [(NSControl*)id
		ignoresMultiClick];
}

void NSControl_inst_setIgnoresMultiClick(void *id, BOOL value) {
	[(NSControl*)id
		setIgnoresMultiClick: value];
}

void NSButton_inst_compressWithPrioritizedCompressionOptions(void *id, void* prioritizedOptions) {
	[(NSButton*)id
		compressWithPrioritizedCompressionOptions: prioritizedOptions];
}

void NSButton_inst_highlight(void *id, BOOL flag) {
	[(NSButton*)id
		highlight: flag];
}

NSSize NSButton_inst_minimumSizeWithPrioritizedCompressionOptions(void *id, void* prioritizedOptions) {
	return [(NSButton*)id
		minimumSizeWithPrioritizedCompressionOptions: prioritizedOptions];
}

BOOL NSButton_inst_performKeyEquivalent(void *id, void* key) {
	return [(NSButton*)id
		performKeyEquivalent: key];
}

void NSButton_inst_setNextState(void *id) {
	[(NSButton*)id
		setNextState];
}

void* NSButton_inst_init(void *id) {
	return [(NSButton*)id
		init];
}

void* NSButton_inst_contentTintColor(void *id) {
	return [(NSButton*)id
		contentTintColor];
}

void NSButton_inst_setContentTintColor(void *id, void* value) {
	[(NSButton*)id
		setContentTintColor: value];
}

BOOL NSButton_inst_hasDestructiveAction(void *id) {
	return [(NSButton*)id
		hasDestructiveAction];
}

void NSButton_inst_setHasDestructiveAction(void *id, BOOL value) {
	[(NSButton*)id
		setHasDestructiveAction: value];
}

void* NSButton_inst_alternateTitle(void *id) {
	return [(NSButton*)id
		alternateTitle];
}

void NSButton_inst_setAlternateTitle(void *id, void* value) {
	[(NSButton*)id
		setAlternateTitle: value];
}

void* NSButton_inst_attributedTitle(void *id) {
	return [(NSButton*)id
		attributedTitle];
}

void NSButton_inst_setAttributedTitle(void *id, void* value) {
	[(NSButton*)id
		setAttributedTitle: value];
}

void* NSButton_inst_attributedAlternateTitle(void *id) {
	return [(NSButton*)id
		attributedAlternateTitle];
}

void NSButton_inst_setAttributedAlternateTitle(void *id, void* value) {
	[(NSButton*)id
		setAttributedAlternateTitle: value];
}

void* NSButton_inst_title(void *id) {
	return [(NSButton*)id
		title];
}

void NSButton_inst_setTitle(void *id, void* value) {
	[(NSButton*)id
		setTitle: value];
}

void* NSButton_inst_sound(void *id) {
	return [(NSButton*)id
		sound];
}

void NSButton_inst_setSound(void *id, void* value) {
	[(NSButton*)id
		setSound: value];
}

BOOL NSButton_inst_isSpringLoaded(void *id) {
	return [(NSButton*)id
		isSpringLoaded];
}

void NSButton_inst_setSpringLoaded(void *id, BOOL value) {
	[(NSButton*)id
		setSpringLoaded: value];
}

long NSButton_inst_maxAcceleratorLevel(void *id) {
	return [(NSButton*)id
		maxAcceleratorLevel];
}

void NSButton_inst_setMaxAcceleratorLevel(void *id, long value) {
	[(NSButton*)id
		setMaxAcceleratorLevel: value];
}

void* NSButton_inst_image(void *id) {
	return [(NSButton*)id
		image];
}

void NSButton_inst_setImage(void *id, void* value) {
	[(NSButton*)id
		setImage: value];
}

void* NSButton_inst_alternateImage(void *id) {
	return [(NSButton*)id
		alternateImage];
}

void NSButton_inst_setAlternateImage(void *id, void* value) {
	[(NSButton*)id
		setAlternateImage: value];
}

BOOL NSButton_inst_isBordered(void *id) {
	return [(NSButton*)id
		isBordered];
}

void NSButton_inst_setBordered(void *id, BOOL value) {
	[(NSButton*)id
		setBordered: value];
}

BOOL NSButton_inst_isTransparent(void *id) {
	return [(NSButton*)id
		isTransparent];
}

void NSButton_inst_setTransparent(void *id, BOOL value) {
	[(NSButton*)id
		setTransparent: value];
}

void* NSButton_inst_bezelColor(void *id) {
	return [(NSButton*)id
		bezelColor];
}

void NSButton_inst_setBezelColor(void *id, void* value) {
	[(NSButton*)id
		setBezelColor: value];
}

BOOL NSButton_inst_showsBorderOnlyWhileMouseInside(void *id) {
	return [(NSButton*)id
		showsBorderOnlyWhileMouseInside];
}

void NSButton_inst_setShowsBorderOnlyWhileMouseInside(void *id, BOOL value) {
	[(NSButton*)id
		setShowsBorderOnlyWhileMouseInside: value];
}

BOOL NSButton_inst_imageHugsTitle(void *id) {
	return [(NSButton*)id
		imageHugsTitle];
}

void NSButton_inst_setImageHugsTitle(void *id, BOOL value) {
	[(NSButton*)id
		setImageHugsTitle: value];
}

BOOL NSButton_inst_allowsMixedState(void *id) {
	return [(NSButton*)id
		allowsMixedState];
}

void NSButton_inst_setAllowsMixedState(void *id, BOOL value) {
	[(NSButton*)id
		setAllowsMixedState: value];
}

long NSButton_inst_state(void *id) {
	return [(NSButton*)id
		state];
}

void NSButton_inst_setState(void *id, long value) {
	[(NSButton*)id
		setState: value];
}

void* NSButton_inst_keyEquivalent(void *id) {
	return [(NSButton*)id
		keyEquivalent];
}

void NSButton_inst_setKeyEquivalent(void *id, void* value) {
	[(NSButton*)id
		setKeyEquivalent: value];
}

void* NSEvent_inst_init(void *id) {
	return [(NSEvent*)id
		init];
}

NSPoint NSEvent_inst_locationInWindow(void *id) {
	return [(NSEvent*)id
		locationInWindow];
}

void* NSEvent_inst_window(void *id) {
	return [(NSEvent*)id
		window];
}

long NSEvent_inst_windowNumber(void *id) {
	return [(NSEvent*)id
		windowNumber];
}

void* NSEvent_inst_eventRef(void *id) {
	return [(NSEvent*)id
		eventRef];
}

void* NSEvent_inst_characters(void *id) {
	return [(NSEvent*)id
		characters];
}

void* NSEvent_inst_charactersIgnoringModifiers(void *id) {
	return [(NSEvent*)id
		charactersIgnoringModifiers];
}

BOOL NSEvent_inst_isARepeat(void *id) {
	return [(NSEvent*)id
		isARepeat];
}

long NSEvent_inst_buttonNumber(void *id) {
	return [(NSEvent*)id
		buttonNumber];
}

long NSEvent_inst_clickCount(void *id) {
	return [(NSEvent*)id
		clickCount];
}

long NSEvent_inst_eventNumber(void *id) {
	return [(NSEvent*)id
		eventNumber];
}

long NSEvent_inst_trackingNumber(void *id) {
	return [(NSEvent*)id
		trackingNumber];
}

void* NSEvent_inst_userData(void *id) {
	return [(NSEvent*)id
		userData];
}

long NSEvent_inst_data1(void *id) {
	return [(NSEvent*)id
		data1];
}

long NSEvent_inst_data2(void *id) {
	return [(NSEvent*)id
		data2];
}

double NSEvent_inst_deltaX(void *id) {
	return [(NSEvent*)id
		deltaX];
}

double NSEvent_inst_deltaY(void *id) {
	return [(NSEvent*)id
		deltaY];
}

double NSEvent_inst_deltaZ(void *id) {
	return [(NSEvent*)id
		deltaZ];
}

long NSEvent_inst_stage(void *id) {
	return [(NSEvent*)id
		stage];
}

double NSEvent_inst_stageTransition(void *id) {
	return [(NSEvent*)id
		stageTransition];
}

unsigned long NSEvent_inst_capabilityMask(void *id) {
	return [(NSEvent*)id
		capabilityMask];
}

unsigned long NSEvent_inst_deviceID(void *id) {
	return [(NSEvent*)id
		deviceID];
}

BOOL NSEvent_inst_isEnteringProximity(void *id) {
	return [(NSEvent*)id
		isEnteringProximity];
}

unsigned long NSEvent_inst_pointingDeviceID(void *id) {
	return [(NSEvent*)id
		pointingDeviceID];
}

unsigned long NSEvent_inst_pointingDeviceSerialNumber(void *id) {
	return [(NSEvent*)id
		pointingDeviceSerialNumber];
}

unsigned long NSEvent_inst_systemTabletID(void *id) {
	return [(NSEvent*)id
		systemTabletID];
}

unsigned long NSEvent_inst_tabletID(void *id) {
	return [(NSEvent*)id
		tabletID];
}

unsigned long NSEvent_inst_vendorID(void *id) {
	return [(NSEvent*)id
		vendorID];
}

unsigned long NSEvent_inst_vendorPointingDeviceType(void *id) {
	return [(NSEvent*)id
		vendorPointingDeviceType];
}

long NSEvent_inst_absoluteX(void *id) {
	return [(NSEvent*)id
		absoluteX];
}

long NSEvent_inst_absoluteY(void *id) {
	return [(NSEvent*)id
		absoluteY];
}

long NSEvent_inst_absoluteZ(void *id) {
	return [(NSEvent*)id
		absoluteZ];
}

NSPoint NSEvent_inst_tilt(void *id) {
	return [(NSEvent*)id
		tilt];
}

void* NSEvent_inst_vendorDefined(void *id) {
	return [(NSEvent*)id
		vendorDefined];
}

double NSEvent_inst_magnification(void *id) {
	return [(NSEvent*)id
		magnification];
}

BOOL NSEvent_inst_hasPreciseScrollingDeltas(void *id) {
	return [(NSEvent*)id
		hasPreciseScrollingDeltas];
}

double NSEvent_inst_scrollingDeltaX(void *id) {
	return [(NSEvent*)id
		scrollingDeltaX];
}

double NSEvent_inst_scrollingDeltaY(void *id) {
	return [(NSEvent*)id
		scrollingDeltaY];
}

BOOL NSEvent_inst_isDirectionInvertedFromDevice(void *id) {
	return [(NSEvent*)id
		isDirectionInvertedFromDevice];
}

void* NSFont_inst_fontWithSize(void *id, double fontSize) {
	return [(NSFont*)id
		fontWithSize: fontSize];
}

void NSFont_inst_set(void *id) {
	[(NSFont*)id
		set];
}

void* NSFont_inst_init(void *id) {
	return [(NSFont*)id
		init];
}

double NSFont_inst_pointSize(void *id) {
	return [(NSFont*)id
		pointSize];
}

BOOL NSFont_inst_isFixedPitch(void *id) {
	return [(NSFont*)id
		isFixedPitch];
}

unsigned long NSFont_inst_mostCompatibleStringEncoding(void *id) {
	return [(NSFont*)id
		mostCompatibleStringEncoding];
}

unsigned long NSFont_inst_numberOfGlyphs(void *id) {
	return [(NSFont*)id
		numberOfGlyphs];
}

void* NSFont_inst_displayName(void *id) {
	return [(NSFont*)id
		displayName];
}

void* NSFont_inst_familyName(void *id) {
	return [(NSFont*)id
		familyName];
}

void* NSFont_inst_fontName(void *id) {
	return [(NSFont*)id
		fontName];
}

BOOL NSFont_inst_isVertical(void *id) {
	return [(NSFont*)id
		isVertical];
}

void* NSFont_inst_verticalFont(void *id) {
	return [(NSFont*)id
		verticalFont];
}

void NSImage_inst_addRepresentations(void *id, void* imageReps) {
	[(NSImage*)id
		addRepresentations: imageReps];
}

void NSImage_inst_cancelIncrementalLoad(void *id) {
	[(NSImage*)id
		cancelIncrementalLoad];
}

void NSImage_inst_drawInRect(void *id, NSRect rect) {
	[(NSImage*)id
		drawInRect: rect];
}

void* NSImage_inst_initByReferencingFile(void *id, void* fileName) {
	return [(NSImage*)id
		initByReferencingFile: fileName];
}

void* NSImage_inst_initByReferencingURL(void *id, void* url) {
	return [(NSImage*)id
		initByReferencingURL: url];
}

void* NSImage_inst_initWithContentsOfFile(void *id, void* fileName) {
	return [(NSImage*)id
		initWithContentsOfFile: fileName];
}

void* NSImage_inst_initWithContentsOfURL(void *id, void* url) {
	return [(NSImage*)id
		initWithContentsOfURL: url];
}

void* NSImage_inst_initWithData(void *id, void* data) {
	return [(NSImage*)id
		initWithData: data];
}

void* NSImage_inst_initWithDataIgnoringOrientation(void *id, void* data) {
	return [(NSImage*)id
		initWithDataIgnoringOrientation: data];
}

void* NSImage_inst_initWithPasteboard(void *id, void* pasteboard) {
	return [(NSImage*)id
		initWithPasteboard: pasteboard];
}

void* NSImage_inst_initWithSize(void *id, NSSize size) {
	return [(NSImage*)id
		initWithSize: size];
}

BOOL NSImage_inst_isTemplate(void *id) {
	return [(NSImage*)id
		isTemplate];
}

void* NSImage_inst_layerContentsForContentsScale(void *id, double layerContentsScale) {
	return [(NSImage*)id
		layerContentsForContentsScale: layerContentsScale];
}

void NSImage_inst_lockFocus(void *id) {
	[(NSImage*)id
		lockFocus];
}

void NSImage_inst_lockFocusFlipped(void *id, BOOL flipped) {
	[(NSImage*)id
		lockFocusFlipped: flipped];
}

void NSImage_inst_recache(void *id) {
	[(NSImage*)id
		recache];
}

double NSImage_inst_recommendedLayerContentsScale(void *id, double preferredContentsScale) {
	return [(NSImage*)id
		recommendedLayerContentsScale: preferredContentsScale];
}

void NSImage_inst_unlockFocus(void *id) {
	[(NSImage*)id
		unlockFocus];
}

void* NSImage_inst_init(void *id) {
	return [(NSImage*)id
		init];
}

void* NSImage_inst_delegate(void *id) {
	return [(NSImage*)id
		delegate];
}

void NSImage_inst_setDelegate(void *id, void* value) {
	[(NSImage*)id
		setDelegate: value];
}

NSSize NSImage_inst_size(void *id) {
	return [(NSImage*)id
		size];
}

void NSImage_inst_setSize(void *id, NSSize value) {
	[(NSImage*)id
		setSize: value];
}

void NSImage_inst_setTemplate(void *id, BOOL value) {
	[(NSImage*)id
		setTemplate: value];
}

void* NSImage_inst_representations(void *id) {
	return [(NSImage*)id
		representations];
}

BOOL NSImage_inst_prefersColorMatch(void *id) {
	return [(NSImage*)id
		prefersColorMatch];
}

void NSImage_inst_setPrefersColorMatch(void *id, BOOL value) {
	[(NSImage*)id
		setPrefersColorMatch: value];
}

BOOL NSImage_inst_usesEPSOnResolutionMismatch(void *id) {
	return [(NSImage*)id
		usesEPSOnResolutionMismatch];
}

void NSImage_inst_setUsesEPSOnResolutionMismatch(void *id, BOOL value) {
	[(NSImage*)id
		setUsesEPSOnResolutionMismatch: value];
}

BOOL NSImage_inst_matchesOnMultipleResolution(void *id) {
	return [(NSImage*)id
		matchesOnMultipleResolution];
}

void NSImage_inst_setMatchesOnMultipleResolution(void *id, BOOL value) {
	[(NSImage*)id
		setMatchesOnMultipleResolution: value];
}

BOOL NSImage_inst_isValid(void *id) {
	return [(NSImage*)id
		isValid];
}

void* NSImage_inst_backgroundColor(void *id) {
	return [(NSImage*)id
		backgroundColor];
}

void NSImage_inst_setBackgroundColor(void *id, void* value) {
	[(NSImage*)id
		setBackgroundColor: value];
}

NSRect NSImage_inst_alignmentRect(void *id) {
	return [(NSImage*)id
		alignmentRect];
}

void NSImage_inst_setAlignmentRect(void *id, NSRect value) {
	[(NSImage*)id
		setAlignmentRect: value];
}

void* NSImage_inst_TIFFRepresentation(void *id) {
	return [(NSImage*)id
		TIFFRepresentation];
}

void* NSImage_inst_accessibilityDescription(void *id) {
	return [(NSImage*)id
		accessibilityDescription];
}

void NSImage_inst_setAccessibilityDescription(void *id, void* value) {
	[(NSImage*)id
		setAccessibilityDescription: value];
}

BOOL NSImage_inst_matchesOnlyOnBestFittingAxis(void *id) {
	return [(NSImage*)id
		matchesOnlyOnBestFittingAxis];
}

void NSImage_inst_setMatchesOnlyOnBestFittingAxis(void *id, BOOL value) {
	[(NSImage*)id
		setMatchesOnlyOnBestFittingAxis: value];
}

void* NSImageView_inst_init(void *id) {
	return [(NSImageView*)id
		init];
}

void* NSImageView_inst_image(void *id) {
	return [(NSImageView*)id
		image];
}

void NSImageView_inst_setImage(void *id, void* value) {
	[(NSImageView*)id
		setImage: value];
}

BOOL NSImageView_inst_animates(void *id) {
	return [(NSImageView*)id
		animates];
}

void NSImageView_inst_setAnimates(void *id, BOOL value) {
	[(NSImageView*)id
		setAnimates: value];
}

BOOL NSImageView_inst_isEditable(void *id) {
	return [(NSImageView*)id
		isEditable];
}

void NSImageView_inst_setEditable(void *id, BOOL value) {
	[(NSImageView*)id
		setEditable: value];
}

BOOL NSImageView_inst_allowsCutCopyPaste(void *id) {
	return [(NSImageView*)id
		allowsCutCopyPaste];
}

void NSImageView_inst_setAllowsCutCopyPaste(void *id, BOOL value) {
	[(NSImageView*)id
		setAllowsCutCopyPaste: value];
}

void* NSImageView_inst_contentTintColor(void *id) {
	return [(NSImageView*)id
		contentTintColor];
}

void NSImageView_inst_setContentTintColor(void *id, void* value) {
	[(NSImageView*)id
		setContentTintColor: value];
}

void* NSNib_inst_initWithNibData_bundle(void *id, void* nibData, void* bundle) {
	return [(NSNib*)id
		initWithNibData: nibData
		bundle: bundle];
}

BOOL NSNib_inst_instantiateWithOwner_topLevelObjects(void *id, void* owner, void* topLevelObjects) {
	return [(NSNib*)id
		instantiateWithOwner: owner
		topLevelObjects: topLevelObjects];
}

void* NSNib_inst_init(void *id) {
	return [(NSNib*)id
		init];
}

long NSPasteboard_inst_addTypes_owner(void *id, void* newTypes, void* newOwner) {
	return [(NSPasteboard*)id
		addTypes: newTypes
		owner: newOwner];
}

BOOL NSPasteboard_inst_canReadItemWithDataConformingToTypes(void *id, void* types) {
	return [(NSPasteboard*)id
		canReadItemWithDataConformingToTypes: types];
}

BOOL NSPasteboard_inst_canReadObjectForClasses_options(void *id, void* classArray, void* options) {
	return [(NSPasteboard*)id
		canReadObjectForClasses: classArray
		options: options];
}

long NSPasteboard_inst_clearContents(void *id) {
	return [(NSPasteboard*)id
		clearContents];
}

long NSPasteboard_inst_declareTypes_owner(void *id, void* newTypes, void* newOwner) {
	return [(NSPasteboard*)id
		declareTypes: newTypes
		owner: newOwner];
}

void* NSPasteboard_inst_readObjectsForClasses_options(void *id, void* classArray, void* options) {
	return [(NSPasteboard*)id
		readObjectsForClasses: classArray
		options: options];
}

void NSPasteboard_inst_releaseGlobally(void *id) {
	[(NSPasteboard*)id
		releaseGlobally];
}

BOOL NSPasteboard_inst_writeFileContents(void *id, void* filename) {
	return [(NSPasteboard*)id
		writeFileContents: filename];
}

BOOL NSPasteboard_inst_writeObjects(void *id, void* objects) {
	return [(NSPasteboard*)id
		writeObjects: objects];
}

void* NSPasteboard_inst_init(void *id) {
	return [(NSPasteboard*)id
		init];
}

void* NSPasteboard_inst_pasteboardItems(void *id) {
	return [(NSPasteboard*)id
		pasteboardItems];
}

void* NSPasteboard_inst_types(void *id) {
	return [(NSPasteboard*)id
		types];
}

long NSPasteboard_inst_changeCount(void *id) {
	return [(NSPasteboard*)id
		changeCount];
}

void NSLayoutManager_inst_addTextContainer(void *id, void* container) {
	[(NSLayoutManager*)id
		addTextContainer: container];
}

NSSize NSLayoutManager_inst_attachmentSizeForGlyphAtIndex(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		attachmentSizeForGlyphAtIndex: glyphIndex];
}

unsigned long NSLayoutManager_inst_characterIndexForGlyphAtIndex(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		characterIndexForGlyphAtIndex: glyphIndex];
}

double NSLayoutManager_inst_defaultBaselineOffsetForFont(void *id, void* theFont) {
	return [(NSLayoutManager*)id
		defaultBaselineOffsetForFont: theFont];
}

double NSLayoutManager_inst_defaultLineHeightForFont(void *id, void* theFont) {
	return [(NSLayoutManager*)id
		defaultLineHeightForFont: theFont];
}

BOOL NSLayoutManager_inst_drawsOutsideLineFragmentForGlyphAtIndex(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		drawsOutsideLineFragmentForGlyphAtIndex: glyphIndex];
}

void NSLayoutManager_inst_ensureLayoutForBoundingRect_inTextContainer(void *id, NSRect bounds, void* container) {
	[(NSLayoutManager*)id
		ensureLayoutForBoundingRect: bounds
		inTextContainer: container];
}

void NSLayoutManager_inst_ensureLayoutForTextContainer(void *id, void* container) {
	[(NSLayoutManager*)id
		ensureLayoutForTextContainer: container];
}

unsigned long NSLayoutManager_inst_firstUnlaidCharacterIndex(void *id) {
	return [(NSLayoutManager*)id
		firstUnlaidCharacterIndex];
}

unsigned long NSLayoutManager_inst_firstUnlaidGlyphIndex(void *id) {
	return [(NSLayoutManager*)id
		firstUnlaidGlyphIndex];
}

unsigned long NSLayoutManager_inst_glyphIndexForCharacterAtIndex(void *id, unsigned long charIndex) {
	return [(NSLayoutManager*)id
		glyphIndexForCharacterAtIndex: charIndex];
}

void* NSLayoutManager_inst_init(void *id) {
	return [(NSLayoutManager*)id
		init];
}

void NSLayoutManager_inst_insertTextContainer_atIndex(void *id, void* container, unsigned long index) {
	[(NSLayoutManager*)id
		insertTextContainer: container
		atIndex: index];
}

BOOL NSLayoutManager_inst_isValidGlyphIndex(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		isValidGlyphIndex: glyphIndex];
}

BOOL NSLayoutManager_inst_layoutManagerOwnsFirstResponderInWindow(void *id, void* window) {
	return [(NSLayoutManager*)id
		layoutManagerOwnsFirstResponderInWindow: window];
}

BOOL NSLayoutManager_inst_notShownAttributeForGlyphAtIndex(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		notShownAttributeForGlyphAtIndex: glyphIndex];
}

void NSLayoutManager_inst_removeTextContainerAtIndex(void *id, unsigned long index) {
	[(NSLayoutManager*)id
		removeTextContainerAtIndex: index];
}

void NSLayoutManager_inst_setDrawsOutsideLineFragment_forGlyphAtIndex(void *id, BOOL flag, unsigned long glyphIndex) {
	[(NSLayoutManager*)id
		setDrawsOutsideLineFragment: flag
		forGlyphAtIndex: glyphIndex];
}

void NSLayoutManager_inst_setExtraLineFragmentRect_usedRect_textContainer(void *id, NSRect fragmentRect, NSRect usedRect, void* container) {
	[(NSLayoutManager*)id
		setExtraLineFragmentRect: fragmentRect
		usedRect: usedRect
		textContainer: container];
}

void NSLayoutManager_inst_setNotShownAttribute_forGlyphAtIndex(void *id, BOOL flag, unsigned long glyphIndex) {
	[(NSLayoutManager*)id
		setNotShownAttribute: flag
		forGlyphAtIndex: glyphIndex];
}

void NSLayoutManager_inst_textContainerChangedGeometry(void *id, void* container) {
	[(NSLayoutManager*)id
		textContainerChangedGeometry: container];
}

void NSLayoutManager_inst_textContainerChangedTextView(void *id, void* container) {
	[(NSLayoutManager*)id
		textContainerChangedTextView: container];
}

NSRect NSLayoutManager_inst_usedRectForTextContainer(void *id, void* container) {
	return [(NSLayoutManager*)id
		usedRectForTextContainer: container];
}

void* NSLayoutManager_inst_delegate(void *id) {
	return [(NSLayoutManager*)id
		delegate];
}

void NSLayoutManager_inst_setDelegate(void *id, void* value) {
	[(NSLayoutManager*)id
		setDelegate: value];
}

BOOL NSLayoutManager_inst_allowsNonContiguousLayout(void *id) {
	return [(NSLayoutManager*)id
		allowsNonContiguousLayout];
}

void NSLayoutManager_inst_setAllowsNonContiguousLayout(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setAllowsNonContiguousLayout: value];
}

BOOL NSLayoutManager_inst_hasNonContiguousLayout(void *id) {
	return [(NSLayoutManager*)id
		hasNonContiguousLayout];
}

BOOL NSLayoutManager_inst_showsInvisibleCharacters(void *id) {
	return [(NSLayoutManager*)id
		showsInvisibleCharacters];
}

void NSLayoutManager_inst_setShowsInvisibleCharacters(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setShowsInvisibleCharacters: value];
}

BOOL NSLayoutManager_inst_showsControlCharacters(void *id) {
	return [(NSLayoutManager*)id
		showsControlCharacters];
}

void NSLayoutManager_inst_setShowsControlCharacters(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setShowsControlCharacters: value];
}

BOOL NSLayoutManager_inst_usesFontLeading(void *id) {
	return [(NSLayoutManager*)id
		usesFontLeading];
}

void NSLayoutManager_inst_setUsesFontLeading(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setUsesFontLeading: value];
}

BOOL NSLayoutManager_inst_backgroundLayoutEnabled(void *id) {
	return [(NSLayoutManager*)id
		backgroundLayoutEnabled];
}

void NSLayoutManager_inst_setBackgroundLayoutEnabled(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setBackgroundLayoutEnabled: value];
}

BOOL NSLayoutManager_inst_limitsLayoutForSuspiciousContents(void *id) {
	return [(NSLayoutManager*)id
		limitsLayoutForSuspiciousContents];
}

void NSLayoutManager_inst_setLimitsLayoutForSuspiciousContents(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setLimitsLayoutForSuspiciousContents: value];
}

BOOL NSLayoutManager_inst_usesDefaultHyphenation(void *id) {
	return [(NSLayoutManager*)id
		usesDefaultHyphenation];
}

void NSLayoutManager_inst_setUsesDefaultHyphenation(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setUsesDefaultHyphenation: value];
}

void* NSLayoutManager_inst_textContainers(void *id) {
	return [(NSLayoutManager*)id
		textContainers];
}

unsigned long NSLayoutManager_inst_numberOfGlyphs(void *id) {
	return [(NSLayoutManager*)id
		numberOfGlyphs];
}

NSRect NSLayoutManager_inst_extraLineFragmentRect(void *id) {
	return [(NSLayoutManager*)id
		extraLineFragmentRect];
}

void* NSLayoutManager_inst_extraLineFragmentTextContainer(void *id) {
	return [(NSLayoutManager*)id
		extraLineFragmentTextContainer];
}

NSRect NSLayoutManager_inst_extraLineFragmentUsedRect(void *id) {
	return [(NSLayoutManager*)id
		extraLineFragmentUsedRect];
}

void* NSLayoutManager_inst_firstTextView(void *id) {
	return [(NSLayoutManager*)id
		firstTextView];
}

void* NSLayoutManager_inst_textViewForBeginningOfSelection(void *id) {
	return [(NSLayoutManager*)id
		textViewForBeginningOfSelection];
}

void NSMenu_inst_addItem(void *id, void* newItem) {
	[(NSMenu*)id
		addItem: newItem];
}

void* NSMenu_inst_addItemWithTitle_action_keyEquivalent(void *id, void* string, void* selector, void* charCode) {
	return [(NSMenu*)id
		addItemWithTitle: string
		action: selector
		keyEquivalent: charCode];
}

void NSMenu_inst_cancelTracking(void *id) {
	[(NSMenu*)id
		cancelTracking];
}

void NSMenu_inst_cancelTrackingWithoutAnimation(void *id) {
	[(NSMenu*)id
		cancelTrackingWithoutAnimation];
}

long NSMenu_inst_indexOfItem(void *id, void* item) {
	return [(NSMenu*)id
		indexOfItem: item];
}

long NSMenu_inst_indexOfItemWithRepresentedObject(void *id, void* object) {
	return [(NSMenu*)id
		indexOfItemWithRepresentedObject: object];
}

long NSMenu_inst_indexOfItemWithSubmenu(void *id, void* submenu) {
	return [(NSMenu*)id
		indexOfItemWithSubmenu: submenu];
}

long NSMenu_inst_indexOfItemWithTag(void *id, long tag) {
	return [(NSMenu*)id
		indexOfItemWithTag: tag];
}

long NSMenu_inst_indexOfItemWithTarget_andAction(void *id, void* target, void* actionSelector) {
	return [(NSMenu*)id
		indexOfItemWithTarget: target
		andAction: actionSelector];
}

long NSMenu_inst_indexOfItemWithTitle(void *id, void* title) {
	return [(NSMenu*)id
		indexOfItemWithTitle: title];
}

void* NSMenu_inst_initWithTitle(void *id, void* title) {
	return [(NSMenu*)id
		initWithTitle: title];
}

void NSMenu_inst_insertItem_atIndex(void *id, void* newItem, long index) {
	[(NSMenu*)id
		insertItem: newItem
		atIndex: index];
}

void* NSMenu_inst_insertItemWithTitle_action_keyEquivalent_atIndex(void *id, void* string, void* selector, void* charCode, long index) {
	return [(NSMenu*)id
		insertItemWithTitle: string
		action: selector
		keyEquivalent: charCode
		atIndex: index];
}

void* NSMenu_inst_itemAtIndex(void *id, long index) {
	return [(NSMenu*)id
		itemAtIndex: index];
}

void NSMenu_inst_itemChanged(void *id, void* item) {
	[(NSMenu*)id
		itemChanged: item];
}

void* NSMenu_inst_itemWithTag(void *id, long tag) {
	return [(NSMenu*)id
		itemWithTag: tag];
}

void* NSMenu_inst_itemWithTitle(void *id, void* title) {
	return [(NSMenu*)id
		itemWithTitle: title];
}

void NSMenu_inst_performActionForItemAtIndex(void *id, long index) {
	[(NSMenu*)id
		performActionForItemAtIndex: index];
}

BOOL NSMenu_inst_performKeyEquivalent(void *id, void* event) {
	return [(NSMenu*)id
		performKeyEquivalent: event];
}

BOOL NSMenu_inst_popUpMenuPositioningItem_atLocation_inView(void *id, void* item, NSPoint location, void* view) {
	return [(NSMenu*)id
		popUpMenuPositioningItem: item
		atLocation: location
		inView: view];
}

void NSMenu_inst_removeAllItems(void *id) {
	[(NSMenu*)id
		removeAllItems];
}

void NSMenu_inst_removeItem(void *id, void* item) {
	[(NSMenu*)id
		removeItem: item];
}

void NSMenu_inst_removeItemAtIndex(void *id, long index) {
	[(NSMenu*)id
		removeItemAtIndex: index];
}

void NSMenu_inst_setSubmenu_forItem(void *id, void* menu, void* item) {
	[(NSMenu*)id
		setSubmenu: menu
		forItem: item];
}

void NSMenu_inst_submenuAction(void *id, void* sender) {
	[(NSMenu*)id
		submenuAction: sender];
}

void NSMenu_inst_update(void *id) {
	[(NSMenu*)id
		update];
}

void* NSMenu_inst_init(void *id) {
	return [(NSMenu*)id
		init];
}

double NSMenu_inst_menuBarHeight(void *id) {
	return [(NSMenu*)id
		menuBarHeight];
}

long NSMenu_inst_numberOfItems(void *id) {
	return [(NSMenu*)id
		numberOfItems];
}

void* NSMenu_inst_itemArray(void *id) {
	return [(NSMenu*)id
		itemArray];
}

void NSMenu_inst_setItemArray(void *id, void* value) {
	[(NSMenu*)id
		setItemArray: value];
}

void* NSMenu_inst_supermenu(void *id) {
	return [(NSMenu*)id
		supermenu];
}

void NSMenu_inst_setSupermenu(void *id, void* value) {
	[(NSMenu*)id
		setSupermenu: value];
}

BOOL NSMenu_inst_autoenablesItems(void *id) {
	return [(NSMenu*)id
		autoenablesItems];
}

void NSMenu_inst_setAutoenablesItems(void *id, BOOL value) {
	[(NSMenu*)id
		setAutoenablesItems: value];
}

void* NSMenu_inst_font(void *id) {
	return [(NSMenu*)id
		font];
}

void NSMenu_inst_setFont(void *id, void* value) {
	[(NSMenu*)id
		setFont: value];
}

void* NSMenu_inst_title(void *id) {
	return [(NSMenu*)id
		title];
}

void NSMenu_inst_setTitle(void *id, void* value) {
	[(NSMenu*)id
		setTitle: value];
}

double NSMenu_inst_minimumWidth(void *id) {
	return [(NSMenu*)id
		minimumWidth];
}

void NSMenu_inst_setMinimumWidth(void *id, double value) {
	[(NSMenu*)id
		setMinimumWidth: value];
}

NSSize NSMenu_inst_size(void *id) {
	return [(NSMenu*)id
		size];
}

BOOL NSMenu_inst_allowsContextMenuPlugIns(void *id) {
	return [(NSMenu*)id
		allowsContextMenuPlugIns];
}

void NSMenu_inst_setAllowsContextMenuPlugIns(void *id, BOOL value) {
	[(NSMenu*)id
		setAllowsContextMenuPlugIns: value];
}

BOOL NSMenu_inst_showsStateColumn(void *id) {
	return [(NSMenu*)id
		showsStateColumn];
}

void NSMenu_inst_setShowsStateColumn(void *id, BOOL value) {
	[(NSMenu*)id
		setShowsStateColumn: value];
}

void* NSMenu_inst_highlightedItem(void *id) {
	return [(NSMenu*)id
		highlightedItem];
}

void* NSMenu_inst_delegate(void *id) {
	return [(NSMenu*)id
		delegate];
}

void NSMenu_inst_setDelegate(void *id, void* value) {
	[(NSMenu*)id
		setDelegate: value];
}

void NSPopover_inst_close(void *id) {
	[(NSPopover*)id
		close];
}

void* NSPopover_inst_init(void *id) {
	return [(NSPopover*)id
		init];
}

void NSPopover_inst_performClose(void *id, void* sender) {
	[(NSPopover*)id
		performClose: sender];
}

long NSPopover_inst_behavior(void *id) {
	return [(NSPopover*)id
		behavior];
}

void NSPopover_inst_setBehavior(void *id, long value) {
	[(NSPopover*)id
		setBehavior: value];
}

NSRect NSPopover_inst_positioningRect(void *id) {
	return [(NSPopover*)id
		positioningRect];
}

void NSPopover_inst_setPositioningRect(void *id, NSRect value) {
	[(NSPopover*)id
		setPositioningRect: value];
}

BOOL NSPopover_inst_animates(void *id) {
	return [(NSPopover*)id
		animates];
}

void NSPopover_inst_setAnimates(void *id, BOOL value) {
	[(NSPopover*)id
		setAnimates: value];
}

NSSize NSPopover_inst_contentSize(void *id) {
	return [(NSPopover*)id
		contentSize];
}

void NSPopover_inst_setContentSize(void *id, NSSize value) {
	[(NSPopover*)id
		setContentSize: value];
}

BOOL NSPopover_inst_isShown(void *id) {
	return [(NSPopover*)id
		isShown];
}

BOOL NSPopover_inst_isDetached(void *id) {
	return [(NSPopover*)id
		isDetached];
}

void* NSMenuItem_inst_initWithTitle_action_keyEquivalent(void *id, void* string, void* selector, void* charCode) {
	return [(NSMenuItem*)id
		initWithTitle: string
		action: selector
		keyEquivalent: charCode];
}

void* NSMenuItem_inst_init(void *id) {
	return [(NSMenuItem*)id
		init];
}

BOOL NSMenuItem_inst_isEnabled(void *id) {
	return [(NSMenuItem*)id
		isEnabled];
}

void NSMenuItem_inst_setEnabled(void *id, BOOL value) {
	[(NSMenuItem*)id
		setEnabled: value];
}

BOOL NSMenuItem_inst_isHidden(void *id) {
	return [(NSMenuItem*)id
		isHidden];
}

void NSMenuItem_inst_setHidden(void *id, BOOL value) {
	[(NSMenuItem*)id
		setHidden: value];
}

BOOL NSMenuItem_inst_isHiddenOrHasHiddenAncestor(void *id) {
	return [(NSMenuItem*)id
		isHiddenOrHasHiddenAncestor];
}

void* NSMenuItem_inst_target(void *id) {
	return [(NSMenuItem*)id
		target];
}

void NSMenuItem_inst_setTarget(void *id, void* value) {
	[(NSMenuItem*)id
		setTarget: value];
}

void* NSMenuItem_inst_action(void *id) {
	return [(NSMenuItem*)id
		action];
}

void NSMenuItem_inst_setAction(void *id, void* value) {
	[(NSMenuItem*)id
		setAction: value];
}

void* NSMenuItem_inst_title(void *id) {
	return [(NSMenuItem*)id
		title];
}

void NSMenuItem_inst_setTitle(void *id, void* value) {
	[(NSMenuItem*)id
		setTitle: value];
}

void* NSMenuItem_inst_attributedTitle(void *id) {
	return [(NSMenuItem*)id
		attributedTitle];
}

void NSMenuItem_inst_setAttributedTitle(void *id, void* value) {
	[(NSMenuItem*)id
		setAttributedTitle: value];
}

long NSMenuItem_inst_tag(void *id) {
	return [(NSMenuItem*)id
		tag];
}

void NSMenuItem_inst_setTag(void *id, long value) {
	[(NSMenuItem*)id
		setTag: value];
}

long NSMenuItem_inst_state(void *id) {
	return [(NSMenuItem*)id
		state];
}

void NSMenuItem_inst_setState(void *id, long value) {
	[(NSMenuItem*)id
		setState: value];
}

void* NSMenuItem_inst_image(void *id) {
	return [(NSMenuItem*)id
		image];
}

void NSMenuItem_inst_setImage(void *id, void* value) {
	[(NSMenuItem*)id
		setImage: value];
}

void* NSMenuItem_inst_onStateImage(void *id) {
	return [(NSMenuItem*)id
		onStateImage];
}

void NSMenuItem_inst_setOnStateImage(void *id, void* value) {
	[(NSMenuItem*)id
		setOnStateImage: value];
}

void* NSMenuItem_inst_offStateImage(void *id) {
	return [(NSMenuItem*)id
		offStateImage];
}

void NSMenuItem_inst_setOffStateImage(void *id, void* value) {
	[(NSMenuItem*)id
		setOffStateImage: value];
}

void* NSMenuItem_inst_mixedStateImage(void *id) {
	return [(NSMenuItem*)id
		mixedStateImage];
}

void NSMenuItem_inst_setMixedStateImage(void *id, void* value) {
	[(NSMenuItem*)id
		setMixedStateImage: value];
}

void* NSMenuItem_inst_submenu(void *id) {
	return [(NSMenuItem*)id
		submenu];
}

void NSMenuItem_inst_setSubmenu(void *id, void* value) {
	[(NSMenuItem*)id
		setSubmenu: value];
}

BOOL NSMenuItem_inst_hasSubmenu(void *id) {
	return [(NSMenuItem*)id
		hasSubmenu];
}

void* NSMenuItem_inst_parentItem(void *id) {
	return [(NSMenuItem*)id
		parentItem];
}

BOOL NSMenuItem_inst_isSeparatorItem(void *id) {
	return [(NSMenuItem*)id
		isSeparatorItem];
}

void* NSMenuItem_inst_menu(void *id) {
	return [(NSMenuItem*)id
		menu];
}

void NSMenuItem_inst_setMenu(void *id, void* value) {
	[(NSMenuItem*)id
		setMenu: value];
}

void* NSMenuItem_inst_keyEquivalent(void *id) {
	return [(NSMenuItem*)id
		keyEquivalent];
}

void NSMenuItem_inst_setKeyEquivalent(void *id, void* value) {
	[(NSMenuItem*)id
		setKeyEquivalent: value];
}

void* NSMenuItem_inst_userKeyEquivalent(void *id) {
	return [(NSMenuItem*)id
		userKeyEquivalent];
}

BOOL NSMenuItem_inst_isAlternate(void *id) {
	return [(NSMenuItem*)id
		isAlternate];
}

void NSMenuItem_inst_setAlternate(void *id, BOOL value) {
	[(NSMenuItem*)id
		setAlternate: value];
}

long NSMenuItem_inst_indentationLevel(void *id) {
	return [(NSMenuItem*)id
		indentationLevel];
}

void NSMenuItem_inst_setIndentationLevel(void *id, long value) {
	[(NSMenuItem*)id
		setIndentationLevel: value];
}

void* NSMenuItem_inst_toolTip(void *id) {
	return [(NSMenuItem*)id
		toolTip];
}

void NSMenuItem_inst_setToolTip(void *id, void* value) {
	[(NSMenuItem*)id
		setToolTip: value];
}

void* NSMenuItem_inst_representedObject(void *id) {
	return [(NSMenuItem*)id
		representedObject];
}

void NSMenuItem_inst_setRepresentedObject(void *id, void* value) {
	[(NSMenuItem*)id
		setRepresentedObject: value];
}

void* NSMenuItem_inst_view(void *id) {
	return [(NSMenuItem*)id
		view];
}

void NSMenuItem_inst_setView(void *id, void* value) {
	[(NSMenuItem*)id
		setView: value];
}

BOOL NSMenuItem_inst_isHighlighted(void *id) {
	return [(NSMenuItem*)id
		isHighlighted];
}

BOOL NSMenuItem_inst_allowsAutomaticKeyEquivalentLocalization(void *id) {
	return [(NSMenuItem*)id
		allowsAutomaticKeyEquivalentLocalization];
}

void NSMenuItem_inst_setAllowsAutomaticKeyEquivalentLocalization(void *id, BOOL value) {
	[(NSMenuItem*)id
		setAllowsAutomaticKeyEquivalentLocalization: value];
}

BOOL NSMenuItem_inst_allowsAutomaticKeyEquivalentMirroring(void *id) {
	return [(NSMenuItem*)id
		allowsAutomaticKeyEquivalentMirroring];
}

void NSMenuItem_inst_setAllowsAutomaticKeyEquivalentMirroring(void *id, BOOL value) {
	[(NSMenuItem*)id
		setAllowsAutomaticKeyEquivalentMirroring: value];
}

BOOL NSMenuItem_inst_allowsKeyEquivalentWhenHidden(void *id) {
	return [(NSMenuItem*)id
		allowsKeyEquivalentWhenHidden];
}

void NSMenuItem_inst_setAllowsKeyEquivalentWhenHidden(void *id, BOOL value) {
	[(NSMenuItem*)id
		setAllowsKeyEquivalentWhenHidden: value];
}

BOOL NSRunningApplication_inst_forceTerminate(void *id) {
	return [(NSRunningApplication*)id
		forceTerminate];
}

BOOL NSRunningApplication_inst_hide(void *id) {
	return [(NSRunningApplication*)id
		hide];
}

BOOL NSRunningApplication_inst_terminate(void *id) {
	return [(NSRunningApplication*)id
		terminate];
}

BOOL NSRunningApplication_inst_unhide(void *id) {
	return [(NSRunningApplication*)id
		unhide];
}

void* NSRunningApplication_inst_init(void *id) {
	return [(NSRunningApplication*)id
		init];
}

BOOL NSRunningApplication_inst_isActive(void *id) {
	return [(NSRunningApplication*)id
		isActive];
}

long NSRunningApplication_inst_activationPolicy(void *id) {
	return [(NSRunningApplication*)id
		activationPolicy];
}

BOOL NSRunningApplication_inst_isHidden(void *id) {
	return [(NSRunningApplication*)id
		isHidden];
}

void* NSRunningApplication_inst_localizedName(void *id) {
	return [(NSRunningApplication*)id
		localizedName];
}

void* NSRunningApplication_inst_icon(void *id) {
	return [(NSRunningApplication*)id
		icon];
}

void* NSRunningApplication_inst_bundleIdentifier(void *id) {
	return [(NSRunningApplication*)id
		bundleIdentifier];
}

void* NSRunningApplication_inst_bundleURL(void *id) {
	return [(NSRunningApplication*)id
		bundleURL];
}

long NSRunningApplication_inst_executableArchitecture(void *id) {
	return [(NSRunningApplication*)id
		executableArchitecture];
}

void* NSRunningApplication_inst_executableURL(void *id) {
	return [(NSRunningApplication*)id
		executableURL];
}

BOOL NSRunningApplication_inst_isFinishedLaunching(void *id) {
	return [(NSRunningApplication*)id
		isFinishedLaunching];
}

BOOL NSRunningApplication_inst_ownsMenuBar(void *id) {
	return [(NSRunningApplication*)id
		ownsMenuBar];
}

BOOL NSRunningApplication_inst_isTerminated(void *id) {
	return [(NSRunningApplication*)id
		isTerminated];
}

NSRect NSScreen_inst_convertRectFromBacking(void *id, NSRect rect) {
	return [(NSScreen*)id
		convertRectFromBacking: rect];
}

NSRect NSScreen_inst_convertRectToBacking(void *id, NSRect rect) {
	return [(NSScreen*)id
		convertRectToBacking: rect];
}

void* NSScreen_inst_init(void *id) {
	return [(NSScreen*)id
		init];
}

NSRect NSScreen_inst_frame(void *id) {
	return [(NSScreen*)id
		frame];
}

void* NSScreen_inst_deviceDescription(void *id) {
	return [(NSScreen*)id
		deviceDescription];
}

NSRect NSScreen_inst_visibleFrame(void *id) {
	return [(NSScreen*)id
		visibleFrame];
}

double NSScreen_inst_backingScaleFactor(void *id) {
	return [(NSScreen*)id
		backingScaleFactor];
}

double NSScreen_inst_maximumPotentialExtendedDynamicRangeColorComponentValue(void *id) {
	return [(NSScreen*)id
		maximumPotentialExtendedDynamicRangeColorComponentValue];
}

double NSScreen_inst_maximumExtendedDynamicRangeColorComponentValue(void *id) {
	return [(NSScreen*)id
		maximumExtendedDynamicRangeColorComponentValue];
}

double NSScreen_inst_maximumReferenceExtendedDynamicRangeColorComponentValue(void *id) {
	return [(NSScreen*)id
		maximumReferenceExtendedDynamicRangeColorComponentValue];
}

void* NSScreen_inst_localizedName(void *id) {
	return [(NSScreen*)id
		localizedName];
}

long NSScreen_inst_maximumFramesPerSecond(void *id) {
	return [(NSScreen*)id
		maximumFramesPerSecond];
}

void NSStatusBar_inst_removeStatusItem(void *id, void* item) {
	[(NSStatusBar*)id
		removeStatusItem: item];
}

void* NSStatusBar_inst_statusItemWithLength(void *id, double length) {
	return [(NSStatusBar*)id
		statusItemWithLength: length];
}

void* NSStatusBar_inst_init(void *id) {
	return [(NSStatusBar*)id
		init];
}

BOOL NSStatusBar_inst_isVertical(void *id) {
	return [(NSStatusBar*)id
		isVertical];
}

double NSStatusBar_inst_thickness(void *id) {
	return [(NSStatusBar*)id
		thickness];
}

void* NSStatusBarButton_inst_init(void *id) {
	return [(NSStatusBarButton*)id
		init];
}

BOOL NSStatusBarButton_inst_appearsDisabled(void *id) {
	return [(NSStatusBarButton*)id
		appearsDisabled];
}

void NSStatusBarButton_inst_setAppearsDisabled(void *id, BOOL value) {
	[(NSStatusBarButton*)id
		setAppearsDisabled: value];
}

void* NSStatusItem_inst_init(void *id) {
	return [(NSStatusItem*)id
		init];
}

void* NSStatusItem_inst_statusBar(void *id) {
	return [(NSStatusItem*)id
		statusBar];
}

void* NSStatusItem_inst_button(void *id) {
	return [(NSStatusItem*)id
		button];
}

void* NSStatusItem_inst_menu(void *id) {
	return [(NSStatusItem*)id
		menu];
}

void NSStatusItem_inst_setMenu(void *id, void* value) {
	[(NSStatusItem*)id
		setMenu: value];
}

BOOL NSStatusItem_inst_isVisible(void *id) {
	return [(NSStatusItem*)id
		isVisible];
}

void NSStatusItem_inst_setVisible(void *id, BOOL value) {
	[(NSStatusItem*)id
		setVisible: value];
}

double NSStatusItem_inst_length(void *id) {
	return [(NSStatusItem*)id
		length];
}

void NSStatusItem_inst_setLength(void *id, double value) {
	[(NSStatusItem*)id
		setLength: value];
}

void NSText_inst_alignCenter(void *id, void* sender) {
	[(NSText*)id
		alignCenter: sender];
}

void NSText_inst_alignLeft(void *id, void* sender) {
	[(NSText*)id
		alignLeft: sender];
}

void NSText_inst_alignRight(void *id, void* sender) {
	[(NSText*)id
		alignRight: sender];
}

void NSText_inst_changeFont(void *id, void* sender) {
	[(NSText*)id
		changeFont: sender];
}

void NSText_inst_checkSpelling(void *id, void* sender) {
	[(NSText*)id
		checkSpelling: sender];
}

void NSText_inst_copy(void *id, void* sender) {
	[(NSText*)id
		copy: sender];
}

void NSText_inst_copyFont(void *id, void* sender) {
	[(NSText*)id
		copyFont: sender];
}

void NSText_inst_copyRuler(void *id, void* sender) {
	[(NSText*)id
		copyRuler: sender];
}

void NSText_inst_cut(void *id, void* sender) {
	[(NSText*)id
		cut: sender];
}

void NSText_inst_delete(void *id, void* sender) {
	[(NSText*)id
		delete: sender];
}

void* NSText_inst_initWithFrame(void *id, NSRect frameRect) {
	return [(NSText*)id
		initWithFrame: frameRect];
}

void NSText_inst_paste(void *id, void* sender) {
	[(NSText*)id
		paste: sender];
}

void NSText_inst_pasteFont(void *id, void* sender) {
	[(NSText*)id
		pasteFont: sender];
}

void NSText_inst_pasteRuler(void *id, void* sender) {
	[(NSText*)id
		pasteRuler: sender];
}

BOOL NSText_inst_readRTFDFromFile(void *id, void* path) {
	return [(NSText*)id
		readRTFDFromFile: path];
}

void NSText_inst_selectAll(void *id, void* sender) {
	[(NSText*)id
		selectAll: sender];
}

void NSText_inst_showGuessPanel(void *id, void* sender) {
	[(NSText*)id
		showGuessPanel: sender];
}

void NSText_inst_sizeToFit(void *id) {
	[(NSText*)id
		sizeToFit];
}

void NSText_inst_subscript(void *id, void* sender) {
	[(NSText*)id
		subscript: sender];
}

void NSText_inst_superscript(void *id, void* sender) {
	[(NSText*)id
		superscript: sender];
}

void NSText_inst_toggleRuler(void *id, void* sender) {
	[(NSText*)id
		toggleRuler: sender];
}

void NSText_inst_underline(void *id, void* sender) {
	[(NSText*)id
		underline: sender];
}

void NSText_inst_unscript(void *id, void* sender) {
	[(NSText*)id
		unscript: sender];
}

BOOL NSText_inst_writeRTFDToFile_atomically(void *id, void* path, BOOL flag) {
	return [(NSText*)id
		writeRTFDToFile: path
		atomically: flag];
}

void* NSText_inst_init(void *id) {
	return [(NSText*)id
		init];
}

void* NSText_inst_string(void *id) {
	return [(NSText*)id
		string];
}

void NSText_inst_setString(void *id, void* value) {
	[(NSText*)id
		setString: value];
}

void* NSText_inst_backgroundColor(void *id) {
	return [(NSText*)id
		backgroundColor];
}

void NSText_inst_setBackgroundColor(void *id, void* value) {
	[(NSText*)id
		setBackgroundColor: value];
}

BOOL NSText_inst_drawsBackground(void *id) {
	return [(NSText*)id
		drawsBackground];
}

void NSText_inst_setDrawsBackground(void *id, BOOL value) {
	[(NSText*)id
		setDrawsBackground: value];
}

BOOL NSText_inst_isEditable(void *id) {
	return [(NSText*)id
		isEditable];
}

void NSText_inst_setEditable(void *id, BOOL value) {
	[(NSText*)id
		setEditable: value];
}

BOOL NSText_inst_isSelectable(void *id) {
	return [(NSText*)id
		isSelectable];
}

void NSText_inst_setSelectable(void *id, BOOL value) {
	[(NSText*)id
		setSelectable: value];
}

BOOL NSText_inst_isFieldEditor(void *id) {
	return [(NSText*)id
		isFieldEditor];
}

void NSText_inst_setFieldEditor(void *id, BOOL value) {
	[(NSText*)id
		setFieldEditor: value];
}

BOOL NSText_inst_isRichText(void *id) {
	return [(NSText*)id
		isRichText];
}

void NSText_inst_setRichText(void *id, BOOL value) {
	[(NSText*)id
		setRichText: value];
}

BOOL NSText_inst_importsGraphics(void *id) {
	return [(NSText*)id
		importsGraphics];
}

void NSText_inst_setImportsGraphics(void *id, BOOL value) {
	[(NSText*)id
		setImportsGraphics: value];
}

BOOL NSText_inst_usesFontPanel(void *id) {
	return [(NSText*)id
		usesFontPanel];
}

void NSText_inst_setUsesFontPanel(void *id, BOOL value) {
	[(NSText*)id
		setUsesFontPanel: value];
}

BOOL NSText_inst_isRulerVisible(void *id) {
	return [(NSText*)id
		isRulerVisible];
}

void* NSText_inst_font(void *id) {
	return [(NSText*)id
		font];
}

void NSText_inst_setFont(void *id, void* value) {
	[(NSText*)id
		setFont: value];
}

void* NSText_inst_textColor(void *id) {
	return [(NSText*)id
		textColor];
}

void NSText_inst_setTextColor(void *id, void* value) {
	[(NSText*)id
		setTextColor: value];
}

NSSize NSText_inst_maxSize(void *id) {
	return [(NSText*)id
		maxSize];
}

void NSText_inst_setMaxSize(void *id, NSSize value) {
	[(NSText*)id
		setMaxSize: value];
}

NSSize NSText_inst_minSize(void *id) {
	return [(NSText*)id
		minSize];
}

void NSText_inst_setMinSize(void *id, NSSize value) {
	[(NSText*)id
		setMinSize: value];
}

BOOL NSText_inst_isVerticallyResizable(void *id) {
	return [(NSText*)id
		isVerticallyResizable];
}

void NSText_inst_setVerticallyResizable(void *id, BOOL value) {
	[(NSText*)id
		setVerticallyResizable: value];
}

BOOL NSText_inst_isHorizontallyResizable(void *id) {
	return [(NSText*)id
		isHorizontallyResizable];
}

void NSText_inst_setHorizontallyResizable(void *id, BOOL value) {
	[(NSText*)id
		setHorizontallyResizable: value];
}

void* NSText_inst_delegate(void *id) {
	return [(NSText*)id
		delegate];
}

void NSText_inst_setDelegate(void *id, void* value) {
	[(NSText*)id
		setDelegate: value];
}

void NSTextField_inst_selectText(void *id, void* sender) {
	[(NSTextField*)id
		selectText: sender];
}

BOOL NSTextField_inst_textShouldBeginEditing(void *id, void* textObject) {
	return [(NSTextField*)id
		textShouldBeginEditing: textObject];
}

BOOL NSTextField_inst_textShouldEndEditing(void *id, void* textObject) {
	return [(NSTextField*)id
		textShouldEndEditing: textObject];
}

void* NSTextField_inst_init(void *id) {
	return [(NSTextField*)id
		init];
}

BOOL NSTextField_inst_isSelectable(void *id) {
	return [(NSTextField*)id
		isSelectable];
}

void NSTextField_inst_setSelectable(void *id, BOOL value) {
	[(NSTextField*)id
		setSelectable: value];
}

BOOL NSTextField_inst_isEditable(void *id) {
	return [(NSTextField*)id
		isEditable];
}

void NSTextField_inst_setEditable(void *id, BOOL value) {
	[(NSTextField*)id
		setEditable: value];
}

BOOL NSTextField_inst_allowsEditingTextAttributes(void *id) {
	return [(NSTextField*)id
		allowsEditingTextAttributes];
}

void NSTextField_inst_setAllowsEditingTextAttributes(void *id, BOOL value) {
	[(NSTextField*)id
		setAllowsEditingTextAttributes: value];
}

BOOL NSTextField_inst_importsGraphics(void *id) {
	return [(NSTextField*)id
		importsGraphics];
}

void NSTextField_inst_setImportsGraphics(void *id, BOOL value) {
	[(NSTextField*)id
		setImportsGraphics: value];
}

void* NSTextField_inst_placeholderString(void *id) {
	return [(NSTextField*)id
		placeholderString];
}

void NSTextField_inst_setPlaceholderString(void *id, void* value) {
	[(NSTextField*)id
		setPlaceholderString: value];
}

void* NSTextField_inst_placeholderAttributedString(void *id) {
	return [(NSTextField*)id
		placeholderAttributedString];
}

void NSTextField_inst_setPlaceholderAttributedString(void *id, void* value) {
	[(NSTextField*)id
		setPlaceholderAttributedString: value];
}

BOOL NSTextField_inst_allowsDefaultTighteningForTruncation(void *id) {
	return [(NSTextField*)id
		allowsDefaultTighteningForTruncation];
}

void NSTextField_inst_setAllowsDefaultTighteningForTruncation(void *id, BOOL value) {
	[(NSTextField*)id
		setAllowsDefaultTighteningForTruncation: value];
}

long NSTextField_inst_maximumNumberOfLines(void *id) {
	return [(NSTextField*)id
		maximumNumberOfLines];
}

void NSTextField_inst_setMaximumNumberOfLines(void *id, long value) {
	[(NSTextField*)id
		setMaximumNumberOfLines: value];
}

double NSTextField_inst_preferredMaxLayoutWidth(void *id) {
	return [(NSTextField*)id
		preferredMaxLayoutWidth];
}

void NSTextField_inst_setPreferredMaxLayoutWidth(void *id, double value) {
	[(NSTextField*)id
		setPreferredMaxLayoutWidth: value];
}

void* NSTextField_inst_textColor(void *id) {
	return [(NSTextField*)id
		textColor];
}

void NSTextField_inst_setTextColor(void *id, void* value) {
	[(NSTextField*)id
		setTextColor: value];
}

void* NSTextField_inst_backgroundColor(void *id) {
	return [(NSTextField*)id
		backgroundColor];
}

void NSTextField_inst_setBackgroundColor(void *id, void* value) {
	[(NSTextField*)id
		setBackgroundColor: value];
}

BOOL NSTextField_inst_drawsBackground(void *id) {
	return [(NSTextField*)id
		drawsBackground];
}

void NSTextField_inst_setDrawsBackground(void *id, BOOL value) {
	[(NSTextField*)id
		setDrawsBackground: value];
}

BOOL NSTextField_inst_isBezeled(void *id) {
	return [(NSTextField*)id
		isBezeled];
}

void NSTextField_inst_setBezeled(void *id, BOOL value) {
	[(NSTextField*)id
		setBezeled: value];
}

BOOL NSTextField_inst_isBordered(void *id) {
	return [(NSTextField*)id
		isBordered];
}

void NSTextField_inst_setBordered(void *id, BOOL value) {
	[(NSTextField*)id
		setBordered: value];
}

BOOL NSTextField_inst_acceptsFirstResponder(void *id) {
	return [(NSTextField*)id
		acceptsFirstResponder];
}

BOOL NSTextField_inst_allowsCharacterPickerTouchBarItem(void *id) {
	return [(NSTextField*)id
		allowsCharacterPickerTouchBarItem];
}

void NSTextField_inst_setAllowsCharacterPickerTouchBarItem(void *id, BOOL value) {
	[(NSTextField*)id
		setAllowsCharacterPickerTouchBarItem: value];
}

BOOL NSTextField_inst_isAutomaticTextCompletionEnabled(void *id) {
	return [(NSTextField*)id
		isAutomaticTextCompletionEnabled];
}

void NSTextField_inst_setAutomaticTextCompletionEnabled(void *id, BOOL value) {
	[(NSTextField*)id
		setAutomaticTextCompletionEnabled: value];
}

void* NSTextField_inst_delegate(void *id) {
	return [(NSTextField*)id
		delegate];
}

void NSTextField_inst_setDelegate(void *id, void* value) {
	[(NSTextField*)id
		setDelegate: value];
}

void* NSTextContainer_inst_initWithSize(void *id, NSSize size) {
	return [(NSTextContainer*)id
		initWithSize: size];
}

void NSTextContainer_inst_replaceLayoutManager(void *id, void* newLayoutManager) {
	[(NSTextContainer*)id
		replaceLayoutManager: newLayoutManager];
}

void* NSTextContainer_inst_init(void *id) {
	return [(NSTextContainer*)id
		init];
}

void* NSTextContainer_inst_layoutManager(void *id) {
	return [(NSTextContainer*)id
		layoutManager];
}

void NSTextContainer_inst_setLayoutManager(void *id, void* value) {
	[(NSTextContainer*)id
		setLayoutManager: value];
}

void* NSTextContainer_inst_textView(void *id) {
	return [(NSTextContainer*)id
		textView];
}

void NSTextContainer_inst_setTextView(void *id, void* value) {
	[(NSTextContainer*)id
		setTextView: value];
}

NSSize NSTextContainer_inst_size(void *id) {
	return [(NSTextContainer*)id
		size];
}

void NSTextContainer_inst_setSize(void *id, NSSize value) {
	[(NSTextContainer*)id
		setSize: value];
}

void* NSTextContainer_inst_exclusionPaths(void *id) {
	return [(NSTextContainer*)id
		exclusionPaths];
}

void NSTextContainer_inst_setExclusionPaths(void *id, void* value) {
	[(NSTextContainer*)id
		setExclusionPaths: value];
}

BOOL NSTextContainer_inst_widthTracksTextView(void *id) {
	return [(NSTextContainer*)id
		widthTracksTextView];
}

void NSTextContainer_inst_setWidthTracksTextView(void *id, BOOL value) {
	[(NSTextContainer*)id
		setWidthTracksTextView: value];
}

BOOL NSTextContainer_inst_heightTracksTextView(void *id) {
	return [(NSTextContainer*)id
		heightTracksTextView];
}

void NSTextContainer_inst_setHeightTracksTextView(void *id, BOOL value) {
	[(NSTextContainer*)id
		setHeightTracksTextView: value];
}

unsigned long NSTextContainer_inst_maximumNumberOfLines(void *id) {
	return [(NSTextContainer*)id
		maximumNumberOfLines];
}

void NSTextContainer_inst_setMaximumNumberOfLines(void *id, unsigned long value) {
	[(NSTextContainer*)id
		setMaximumNumberOfLines: value];
}

double NSTextContainer_inst_lineFragmentPadding(void *id) {
	return [(NSTextContainer*)id
		lineFragmentPadding];
}

void NSTextContainer_inst_setLineFragmentPadding(void *id, double value) {
	[(NSTextContainer*)id
		setLineFragmentPadding: value];
}

BOOL NSTextContainer_inst_isSimpleRectangularTextContainer(void *id) {
	return [(NSTextContainer*)id
		isSimpleRectangularTextContainer];
}

void NSViewController_inst_addChildViewController(void *id, void* childViewController) {
	[(NSViewController*)id
		addChildViewController: childViewController];
}

BOOL NSViewController_inst_commitEditing(void *id) {
	return [(NSViewController*)id
		commitEditing];
}

void NSViewController_inst_commitEditingWithDelegate_didCommitSelector_contextInfo(void *id, void* delegate, void* didCommitSelector, void* contextInfo) {
	[(NSViewController*)id
		commitEditingWithDelegate: delegate
		didCommitSelector: didCommitSelector
		contextInfo: contextInfo];
}

void NSViewController_inst_discardEditing(void *id) {
	[(NSViewController*)id
		discardEditing];
}

void NSViewController_inst_dismissController(void *id, void* sender) {
	[(NSViewController*)id
		dismissController: sender];
}

void NSViewController_inst_dismissViewController(void *id, void* viewController) {
	[(NSViewController*)id
		dismissViewController: viewController];
}

void NSViewController_inst_insertChildViewController_atIndex(void *id, void* childViewController, long index) {
	[(NSViewController*)id
		insertChildViewController: childViewController
		atIndex: index];
}

void NSViewController_inst_loadView(void *id) {
	[(NSViewController*)id
		loadView];
}

void NSViewController_inst_preferredContentSizeDidChangeForViewController(void *id, void* viewController) {
	[(NSViewController*)id
		preferredContentSizeDidChangeForViewController: viewController];
}

void NSViewController_inst_presentViewController_animator(void *id, void* viewController, void* animator) {
	[(NSViewController*)id
		presentViewController: viewController
		animator: animator];
}

void NSViewController_inst_presentViewControllerAsModalWindow(void *id, void* viewController) {
	[(NSViewController*)id
		presentViewControllerAsModalWindow: viewController];
}

void NSViewController_inst_presentViewControllerAsSheet(void *id, void* viewController) {
	[(NSViewController*)id
		presentViewControllerAsSheet: viewController];
}

void NSViewController_inst_removeChildViewControllerAtIndex(void *id, long index) {
	[(NSViewController*)id
		removeChildViewControllerAtIndex: index];
}

void NSViewController_inst_removeFromParentViewController(void *id) {
	[(NSViewController*)id
		removeFromParentViewController];
}

void NSViewController_inst_updateViewConstraints(void *id) {
	[(NSViewController*)id
		updateViewConstraints];
}

void NSViewController_inst_viewDidAppear(void *id) {
	[(NSViewController*)id
		viewDidAppear];
}

void NSViewController_inst_viewDidDisappear(void *id) {
	[(NSViewController*)id
		viewDidDisappear];
}

void NSViewController_inst_viewDidLayout(void *id) {
	[(NSViewController*)id
		viewDidLayout];
}

void NSViewController_inst_viewDidLoad(void *id) {
	[(NSViewController*)id
		viewDidLoad];
}

void NSViewController_inst_viewWillAppear(void *id) {
	[(NSViewController*)id
		viewWillAppear];
}

void NSViewController_inst_viewWillDisappear(void *id) {
	[(NSViewController*)id
		viewWillDisappear];
}

void NSViewController_inst_viewWillLayout(void *id) {
	[(NSViewController*)id
		viewWillLayout];
}

void NSViewController_inst_viewWillTransitionToSize(void *id, NSSize newSize) {
	[(NSViewController*)id
		viewWillTransitionToSize: newSize];
}

void* NSViewController_inst_init(void *id) {
	return [(NSViewController*)id
		init];
}

void* NSViewController_inst_representedObject(void *id) {
	return [(NSViewController*)id
		representedObject];
}

void NSViewController_inst_setRepresentedObject(void *id, void* value) {
	[(NSViewController*)id
		setRepresentedObject: value];
}

void* NSViewController_inst_nibBundle(void *id) {
	return [(NSViewController*)id
		nibBundle];
}

void* NSViewController_inst_view(void *id) {
	return [(NSViewController*)id
		view];
}

void NSViewController_inst_setView(void *id, void* value) {
	[(NSViewController*)id
		setView: value];
}

void* NSViewController_inst_title(void *id) {
	return [(NSViewController*)id
		title];
}

void NSViewController_inst_setTitle(void *id, void* value) {
	[(NSViewController*)id
		setTitle: value];
}

BOOL NSViewController_inst_isViewLoaded(void *id) {
	return [(NSViewController*)id
		isViewLoaded];
}

NSSize NSViewController_inst_preferredContentSize(void *id) {
	return [(NSViewController*)id
		preferredContentSize];
}

void NSViewController_inst_setPreferredContentSize(void *id, NSSize value) {
	[(NSViewController*)id
		setPreferredContentSize: value];
}

void* NSViewController_inst_childViewControllers(void *id) {
	return [(NSViewController*)id
		childViewControllers];
}

void NSViewController_inst_setChildViewControllers(void *id, void* value) {
	[(NSViewController*)id
		setChildViewControllers: value];
}

void* NSViewController_inst_parentViewController(void *id) {
	return [(NSViewController*)id
		parentViewController];
}

void* NSViewController_inst_presentedViewControllers(void *id) {
	return [(NSViewController*)id
		presentedViewControllers];
}

void* NSViewController_inst_presentingViewController(void *id) {
	return [(NSViewController*)id
		presentingViewController];
}

NSPoint NSViewController_inst_preferredScreenOrigin(void *id) {
	return [(NSViewController*)id
		preferredScreenOrigin];
}

void NSViewController_inst_setPreferredScreenOrigin(void *id, NSPoint value) {
	[(NSViewController*)id
		setPreferredScreenOrigin: value];
}

NSSize NSViewController_inst_preferredMaximumSize(void *id) {
	return [(NSViewController*)id
		preferredMaximumSize];
}

NSSize NSViewController_inst_preferredMinimumSize(void *id) {
	return [(NSViewController*)id
		preferredMinimumSize];
}

void* NSViewController_inst_sourceItemView(void *id) {
	return [(NSViewController*)id
		sourceItemView];
}

void NSViewController_inst_setSourceItemView(void *id, void* value) {
	[(NSViewController*)id
		setSourceItemView: value];
}

void NSVisualEffectView_inst_viewDidMoveToWindow(void *id) {
	[(NSVisualEffectView*)id
		viewDidMoveToWindow];
}

void NSVisualEffectView_inst_viewWillMoveToWindow(void *id, void* newWindow) {
	[(NSVisualEffectView*)id
		viewWillMoveToWindow: newWindow];
}

void* NSVisualEffectView_inst_init(void *id) {
	return [(NSVisualEffectView*)id
		init];
}

BOOL NSVisualEffectView_inst_isEmphasized(void *id) {
	return [(NSVisualEffectView*)id
		isEmphasized];
}

void NSVisualEffectView_inst_setEmphasized(void *id, BOOL value) {
	[(NSVisualEffectView*)id
		setEmphasized: value];
}

void* NSVisualEffectView_inst_maskImage(void *id) {
	return [(NSVisualEffectView*)id
		maskImage];
}

void NSVisualEffectView_inst_setMaskImage(void *id, void* value) {
	[(NSVisualEffectView*)id
		setMaskImage: value];
}

void NSWindow_inst_addChildWindow_ordered(void *id, void* childWin, unsigned long place) {
	[(NSWindow*)id
		addChildWindow: childWin
		ordered: place];
}

void NSWindow_inst_addTabbedWindow_ordered(void *id, void* window, unsigned long ordered) {
	[(NSWindow*)id
		addTabbedWindow: window
		ordered: ordered];
}

void NSWindow_inst_becomeKeyWindow(void *id) {
	[(NSWindow*)id
		becomeKeyWindow];
}

void NSWindow_inst_becomeMainWindow(void *id) {
	[(NSWindow*)id
		becomeMainWindow];
}

NSPoint NSWindow_inst_cascadeTopLeftFromPoint(void *id, NSPoint topLeftPoint) {
	return [(NSWindow*)id
		cascadeTopLeftFromPoint: topLeftPoint];
}

void NSWindow_inst_center(void *id) {
	[(NSWindow*)id
		center];
}

void NSWindow_inst_close(void *id) {
	[(NSWindow*)id
		close];
}

NSRect NSWindow_inst_constrainFrameRect_toScreen(void *id, NSRect frameRect, void* screen) {
	return [(NSWindow*)id
		constrainFrameRect: frameRect
		toScreen: screen];
}

NSRect NSWindow_inst_contentRectForFrameRect(void *id, NSRect frameRect) {
	return [(NSWindow*)id
		contentRectForFrameRect: frameRect];
}

NSPoint NSWindow_inst_convertPointFromBacking(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointFromBacking: point];
}

NSPoint NSWindow_inst_convertPointFromScreen(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointFromScreen: point];
}

NSPoint NSWindow_inst_convertPointToBacking(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointToBacking: point];
}

NSPoint NSWindow_inst_convertPointToScreen(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointToScreen: point];
}

NSRect NSWindow_inst_convertRectFromBacking(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectFromBacking: rect];
}

NSRect NSWindow_inst_convertRectFromScreen(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectFromScreen: rect];
}

NSRect NSWindow_inst_convertRectToBacking(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectToBacking: rect];
}

NSRect NSWindow_inst_convertRectToScreen(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectToScreen: rect];
}

void* NSWindow_inst_dataWithEPSInsideRect(void *id, NSRect rect) {
	return [(NSWindow*)id
		dataWithEPSInsideRect: rect];
}

void* NSWindow_inst_dataWithPDFInsideRect(void *id, NSRect rect) {
	return [(NSWindow*)id
		dataWithPDFInsideRect: rect];
}

void NSWindow_inst_deminiaturize(void *id, void* sender) {
	[(NSWindow*)id
		deminiaturize: sender];
}

void NSWindow_inst_disableCursorRects(void *id) {
	[(NSWindow*)id
		disableCursorRects];
}

void NSWindow_inst_disableKeyEquivalentForDefaultButtonCell(void *id) {
	[(NSWindow*)id
		disableKeyEquivalentForDefaultButtonCell];
}

void NSWindow_inst_disableScreenUpdatesUntilFlush(void *id) {
	[(NSWindow*)id
		disableScreenUpdatesUntilFlush];
}

void NSWindow_inst_disableSnapshotRestoration(void *id) {
	[(NSWindow*)id
		disableSnapshotRestoration];
}

void NSWindow_inst_discardCursorRects(void *id) {
	[(NSWindow*)id
		discardCursorRects];
}

void NSWindow_inst_display(void *id) {
	[(NSWindow*)id
		display];
}

void NSWindow_inst_displayIfNeeded(void *id) {
	[(NSWindow*)id
		displayIfNeeded];
}

void NSWindow_inst_dragImage_at_offset_event_pasteboard_source_slideBack(void *id, void* image, NSPoint baseLocation, NSSize initialOffset, void* event, void* pboard, void* sourceObj, BOOL slideFlag) {
	[(NSWindow*)id
		dragImage: image
		at: baseLocation
		offset: initialOffset
		event: event
		pasteboard: pboard
		source: sourceObj
		slideBack: slideFlag];
}

void NSWindow_inst_enableCursorRects(void *id) {
	[(NSWindow*)id
		enableCursorRects];
}

void NSWindow_inst_enableKeyEquivalentForDefaultButtonCell(void *id) {
	[(NSWindow*)id
		enableKeyEquivalentForDefaultButtonCell];
}

void NSWindow_inst_enableSnapshotRestoration(void *id) {
	[(NSWindow*)id
		enableSnapshotRestoration];
}

void NSWindow_inst_endEditingFor(void *id, void* object) {
	[(NSWindow*)id
		endEditingFor: object];
}

void NSWindow_inst_endSheet(void *id, void* sheetWindow) {
	[(NSWindow*)id
		endSheet: sheetWindow];
}

void* NSWindow_inst_fieldEditor_forObject(void *id, BOOL createFlag, void* object) {
	return [(NSWindow*)id
		fieldEditor: createFlag
		forObject: object];
}

NSRect NSWindow_inst_frameRectForContentRect(void *id, NSRect contentRect) {
	return [(NSWindow*)id
		frameRectForContentRect: contentRect];
}

void* NSWindow_inst_initWithContentRect_styleMask_backing_defer(void *id, NSRect contentRect, unsigned long style, unsigned long backingStoreType, BOOL flag) {
	return [(NSWindow*)id
		initWithContentRect: contentRect
		styleMask: style
		backing: backingStoreType
		defer: flag];
}

void* NSWindow_inst_initWithContentRect_styleMask_backing_defer_screen(void *id, NSRect contentRect, unsigned long style, unsigned long backingStoreType, BOOL flag, void* screen) {
	return [(NSWindow*)id
		initWithContentRect: contentRect
		styleMask: style
		backing: backingStoreType
		defer: flag
		screen: screen];
}

void NSWindow_inst_invalidateCursorRectsForView(void *id, void* view) {
	[(NSWindow*)id
		invalidateCursorRectsForView: view];
}

void NSWindow_inst_invalidateShadow(void *id) {
	[(NSWindow*)id
		invalidateShadow];
}

void NSWindow_inst_layoutIfNeeded(void *id) {
	[(NSWindow*)id
		layoutIfNeeded];
}

void NSWindow_inst_makeKeyAndOrderFront(void *id, void* sender) {
	[(NSWindow*)id
		makeKeyAndOrderFront: sender];
}

void NSWindow_inst_makeKeyWindow(void *id) {
	[(NSWindow*)id
		makeKeyWindow];
}

void NSWindow_inst_makeMainWindow(void *id) {
	[(NSWindow*)id
		makeMainWindow];
}

void NSWindow_inst_mergeAllWindows(void *id, void* sender) {
	[(NSWindow*)id
		mergeAllWindows: sender];
}

void NSWindow_inst_miniaturize(void *id, void* sender) {
	[(NSWindow*)id
		miniaturize: sender];
}

void NSWindow_inst_moveTabToNewWindow(void *id, void* sender) {
	[(NSWindow*)id
		moveTabToNewWindow: sender];
}

void NSWindow_inst_orderBack(void *id, void* sender) {
	[(NSWindow*)id
		orderBack: sender];
}

void NSWindow_inst_orderFront(void *id, void* sender) {
	[(NSWindow*)id
		orderFront: sender];
}

void NSWindow_inst_orderFrontRegardless(void *id) {
	[(NSWindow*)id
		orderFrontRegardless];
}

void NSWindow_inst_orderOut(void *id, void* sender) {
	[(NSWindow*)id
		orderOut: sender];
}

void NSWindow_inst_orderWindow_relativeTo(void *id, unsigned long place, long otherWin) {
	[(NSWindow*)id
		orderWindow: place
		relativeTo: otherWin];
}

void NSWindow_inst_performClose(void *id, void* sender) {
	[(NSWindow*)id
		performClose: sender];
}

void NSWindow_inst_performMiniaturize(void *id, void* sender) {
	[(NSWindow*)id
		performMiniaturize: sender];
}

void NSWindow_inst_performWindowDragWithEvent(void *id, void* event) {
	[(NSWindow*)id
		performWindowDragWithEvent: event];
}

void NSWindow_inst_performZoom(void *id, void* sender) {
	[(NSWindow*)id
		performZoom: sender];
}

void NSWindow_inst_postEvent_atStart(void *id, void* event, BOOL flag) {
	[(NSWindow*)id
		postEvent: event
		atStart: flag];
}

void NSWindow_inst_print(void *id, void* sender) {
	[(NSWindow*)id
		print: sender];
}

void NSWindow_inst_recalculateKeyViewLoop(void *id) {
	[(NSWindow*)id
		recalculateKeyViewLoop];
}

void NSWindow_inst_registerForDraggedTypes(void *id, void* newTypes) {
	[(NSWindow*)id
		registerForDraggedTypes: newTypes];
}

void NSWindow_inst_removeChildWindow(void *id, void* childWin) {
	[(NSWindow*)id
		removeChildWindow: childWin];
}

void NSWindow_inst_removeTitlebarAccessoryViewControllerAtIndex(void *id, long index) {
	[(NSWindow*)id
		removeTitlebarAccessoryViewControllerAtIndex: index];
}

void NSWindow_inst_resetCursorRects(void *id) {
	[(NSWindow*)id
		resetCursorRects];
}

void NSWindow_inst_resignKeyWindow(void *id) {
	[(NSWindow*)id
		resignKeyWindow];
}

void NSWindow_inst_resignMainWindow(void *id) {
	[(NSWindow*)id
		resignMainWindow];
}

void NSWindow_inst_runToolbarCustomizationPalette(void *id, void* sender) {
	[(NSWindow*)id
		runToolbarCustomizationPalette: sender];
}

void NSWindow_inst_selectKeyViewFollowingView(void *id, void* view) {
	[(NSWindow*)id
		selectKeyViewFollowingView: view];
}

void NSWindow_inst_selectKeyViewPrecedingView(void *id, void* view) {
	[(NSWindow*)id
		selectKeyViewPrecedingView: view];
}

void NSWindow_inst_selectNextKeyView(void *id, void* sender) {
	[(NSWindow*)id
		selectNextKeyView: sender];
}

void NSWindow_inst_selectNextTab(void *id, void* sender) {
	[(NSWindow*)id
		selectNextTab: sender];
}

void NSWindow_inst_selectPreviousKeyView(void *id, void* sender) {
	[(NSWindow*)id
		selectPreviousKeyView: sender];
}

void NSWindow_inst_selectPreviousTab(void *id, void* sender) {
	[(NSWindow*)id
		selectPreviousTab: sender];
}

void NSWindow_inst_sendEvent(void *id, void* event) {
	[(NSWindow*)id
		sendEvent: event];
}

void NSWindow_inst_setContentSize(void *id, NSSize size) {
	[(NSWindow*)id
		setContentSize: size];
}

void NSWindow_inst_setDynamicDepthLimit(void *id, BOOL flag) {
	[(NSWindow*)id
		setDynamicDepthLimit: flag];
}

void NSWindow_inst_setFrame_display(void *id, NSRect frameRect, BOOL flag) {
	[(NSWindow*)id
		setFrame: frameRect
		display: flag];
}

void NSWindow_inst_setFrame_display_animate(void *id, NSRect frameRect, BOOL displayFlag, BOOL animateFlag) {
	[(NSWindow*)id
		setFrame: frameRect
		display: displayFlag
		animate: animateFlag];
}

void NSWindow_inst_setFrameOrigin(void *id, NSPoint point) {
	[(NSWindow*)id
		setFrameOrigin: point];
}

void NSWindow_inst_setFrameTopLeftPoint(void *id, NSPoint point) {
	[(NSWindow*)id
		setFrameTopLeftPoint: point];
}

void NSWindow_inst_setIsMiniaturized(void *id, BOOL flag) {
	[(NSWindow*)id
		setIsMiniaturized: flag];
}

void NSWindow_inst_setIsVisible(void *id, BOOL flag) {
	[(NSWindow*)id
		setIsVisible: flag];
}

void NSWindow_inst_setIsZoomed(void *id, BOOL flag) {
	[(NSWindow*)id
		setIsZoomed: flag];
}

void NSWindow_inst_setTitleWithRepresentedFilename(void *id, void* filename) {
	[(NSWindow*)id
		setTitleWithRepresentedFilename: filename];
}

void NSWindow_inst_toggleFullScreen(void *id, void* sender) {
	[(NSWindow*)id
		toggleFullScreen: sender];
}

void NSWindow_inst_toggleTabBar(void *id, void* sender) {
	[(NSWindow*)id
		toggleTabBar: sender];
}

void NSWindow_inst_toggleTabOverview(void *id, void* sender) {
	[(NSWindow*)id
		toggleTabOverview: sender];
}

void NSWindow_inst_toggleToolbarShown(void *id, void* sender) {
	[(NSWindow*)id
		toggleToolbarShown: sender];
}

BOOL NSWindow_inst_tryToPerform_with(void *id, void* action, void* object) {
	return [(NSWindow*)id
		tryToPerform: action
		with: object];
}

void NSWindow_inst_unregisterDraggedTypes(void *id) {
	[(NSWindow*)id
		unregisterDraggedTypes];
}

void NSWindow_inst_update(void *id) {
	[(NSWindow*)id
		update];
}

void NSWindow_inst_updateConstraintsIfNeeded(void *id) {
	[(NSWindow*)id
		updateConstraintsIfNeeded];
}

void NSWindow_inst_visualizeConstraints(void *id, void* constraints) {
	[(NSWindow*)id
		visualizeConstraints: constraints];
}

void NSWindow_inst_zoom(void *id, void* sender) {
	[(NSWindow*)id
		zoom: sender];
}

void* NSWindow_inst_init(void *id) {
	return [(NSWindow*)id
		init];
}

void* NSWindow_inst_delegate(void *id) {
	return [(NSWindow*)id
		delegate];
}

void NSWindow_inst_setDelegate(void *id, void* value) {
	[(NSWindow*)id
		setDelegate: value];
}

void* NSWindow_inst_contentViewController(void *id) {
	return [(NSWindow*)id
		contentViewController];
}

void NSWindow_inst_setContentViewController(void *id, void* value) {
	[(NSWindow*)id
		setContentViewController: value];
}

void* NSWindow_inst_contentView(void *id) {
	return [(NSWindow*)id
		contentView];
}

void NSWindow_inst_setContentView(void *id, void* value) {
	[(NSWindow*)id
		setContentView: value];
}

unsigned long NSWindow_inst_styleMask(void *id) {
	return [(NSWindow*)id
		styleMask];
}

void NSWindow_inst_setStyleMask(void *id, unsigned long value) {
	[(NSWindow*)id
		setStyleMask: value];
}

BOOL NSWindow_inst_worksWhenModal(void *id) {
	return [(NSWindow*)id
		worksWhenModal];
}

double NSWindow_inst_alphaValue(void *id) {
	return [(NSWindow*)id
		alphaValue];
}

void NSWindow_inst_setAlphaValue(void *id, double value) {
	[(NSWindow*)id
		setAlphaValue: value];
}

void* NSWindow_inst_backgroundColor(void *id) {
	return [(NSWindow*)id
		backgroundColor];
}

void NSWindow_inst_setBackgroundColor(void *id, void* value) {
	[(NSWindow*)id
		setBackgroundColor: value];
}

BOOL NSWindow_inst_canHide(void *id) {
	return [(NSWindow*)id
		canHide];
}

void NSWindow_inst_setCanHide(void *id, BOOL value) {
	[(NSWindow*)id
		setCanHide: value];
}

BOOL NSWindow_inst_isOnActiveSpace(void *id) {
	return [(NSWindow*)id
		isOnActiveSpace];
}

BOOL NSWindow_inst_hidesOnDeactivate(void *id) {
	return [(NSWindow*)id
		hidesOnDeactivate];
}

void NSWindow_inst_setHidesOnDeactivate(void *id, BOOL value) {
	[(NSWindow*)id
		setHidesOnDeactivate: value];
}

unsigned long NSWindow_inst_collectionBehavior(void *id) {
	return [(NSWindow*)id
		collectionBehavior];
}

void NSWindow_inst_setCollectionBehavior(void *id, unsigned long value) {
	[(NSWindow*)id
		setCollectionBehavior: value];
}

BOOL NSWindow_inst_isOpaque(void *id) {
	return [(NSWindow*)id
		isOpaque];
}

void NSWindow_inst_setOpaque(void *id, BOOL value) {
	[(NSWindow*)id
		setOpaque: value];
}

BOOL NSWindow_inst_hasShadow(void *id) {
	return [(NSWindow*)id
		hasShadow];
}

void NSWindow_inst_setHasShadow(void *id, BOOL value) {
	[(NSWindow*)id
		setHasShadow: value];
}

BOOL NSWindow_inst_preventsApplicationTerminationWhenModal(void *id) {
	return [(NSWindow*)id
		preventsApplicationTerminationWhenModal];
}

void NSWindow_inst_setPreventsApplicationTerminationWhenModal(void *id, BOOL value) {
	[(NSWindow*)id
		setPreventsApplicationTerminationWhenModal: value];
}

BOOL NSWindow_inst_hasDynamicDepthLimit(void *id) {
	return [(NSWindow*)id
		hasDynamicDepthLimit];
}

long NSWindow_inst_windowNumber(void *id) {
	return [(NSWindow*)id
		windowNumber];
}

void* NSWindow_inst_deviceDescription(void *id) {
	return [(NSWindow*)id
		deviceDescription];
}

BOOL NSWindow_inst_canBecomeVisibleWithoutLogin(void *id) {
	return [(NSWindow*)id
		canBecomeVisibleWithoutLogin];
}

void NSWindow_inst_setCanBecomeVisibleWithoutLogin(void *id, BOOL value) {
	[(NSWindow*)id
		setCanBecomeVisibleWithoutLogin: value];
}

unsigned long NSWindow_inst_backingType(void *id) {
	return [(NSWindow*)id
		backingType];
}

void NSWindow_inst_setBackingType(void *id, unsigned long value) {
	[(NSWindow*)id
		setBackingType: value];
}

void* NSWindow_inst_attachedSheet(void *id) {
	return [(NSWindow*)id
		attachedSheet];
}

BOOL NSWindow_inst_isSheet(void *id) {
	return [(NSWindow*)id
		isSheet];
}

void* NSWindow_inst_sheetParent(void *id) {
	return [(NSWindow*)id
		sheetParent];
}

void* NSWindow_inst_sheets(void *id) {
	return [(NSWindow*)id
		sheets];
}

NSRect NSWindow_inst_frame(void *id) {
	return [(NSWindow*)id
		frame];
}

NSSize NSWindow_inst_aspectRatio(void *id) {
	return [(NSWindow*)id
		aspectRatio];
}

void NSWindow_inst_setAspectRatio(void *id, NSSize value) {
	[(NSWindow*)id
		setAspectRatio: value];
}

NSSize NSWindow_inst_minSize(void *id) {
	return [(NSWindow*)id
		minSize];
}

void NSWindow_inst_setMinSize(void *id, NSSize value) {
	[(NSWindow*)id
		setMinSize: value];
}

NSSize NSWindow_inst_maxSize(void *id) {
	return [(NSWindow*)id
		maxSize];
}

void NSWindow_inst_setMaxSize(void *id, NSSize value) {
	[(NSWindow*)id
		setMaxSize: value];
}

BOOL NSWindow_inst_isZoomed(void *id) {
	return [(NSWindow*)id
		isZoomed];
}

NSSize NSWindow_inst_resizeIncrements(void *id) {
	return [(NSWindow*)id
		resizeIncrements];
}

void NSWindow_inst_setResizeIncrements(void *id, NSSize value) {
	[(NSWindow*)id
		setResizeIncrements: value];
}

BOOL NSWindow_inst_preservesContentDuringLiveResize(void *id) {
	return [(NSWindow*)id
		preservesContentDuringLiveResize];
}

void NSWindow_inst_setPreservesContentDuringLiveResize(void *id, BOOL value) {
	[(NSWindow*)id
		setPreservesContentDuringLiveResize: value];
}

BOOL NSWindow_inst_inLiveResize(void *id) {
	return [(NSWindow*)id
		inLiveResize];
}

NSSize NSWindow_inst_contentAspectRatio(void *id) {
	return [(NSWindow*)id
		contentAspectRatio];
}

void NSWindow_inst_setContentAspectRatio(void *id, NSSize value) {
	[(NSWindow*)id
		setContentAspectRatio: value];
}

NSSize NSWindow_inst_contentMinSize(void *id) {
	return [(NSWindow*)id
		contentMinSize];
}

void NSWindow_inst_setContentMinSize(void *id, NSSize value) {
	[(NSWindow*)id
		setContentMinSize: value];
}

NSSize NSWindow_inst_contentMaxSize(void *id) {
	return [(NSWindow*)id
		contentMaxSize];
}

void NSWindow_inst_setContentMaxSize(void *id, NSSize value) {
	[(NSWindow*)id
		setContentMaxSize: value];
}

NSSize NSWindow_inst_contentResizeIncrements(void *id) {
	return [(NSWindow*)id
		contentResizeIncrements];
}

void NSWindow_inst_setContentResizeIncrements(void *id, NSSize value) {
	[(NSWindow*)id
		setContentResizeIncrements: value];
}

void* NSWindow_inst_contentLayoutGuide(void *id) {
	return [(NSWindow*)id
		contentLayoutGuide];
}

NSRect NSWindow_inst_contentLayoutRect(void *id) {
	return [(NSWindow*)id
		contentLayoutRect];
}

NSSize NSWindow_inst_maxFullScreenContentSize(void *id) {
	return [(NSWindow*)id
		maxFullScreenContentSize];
}

void NSWindow_inst_setMaxFullScreenContentSize(void *id, NSSize value) {
	[(NSWindow*)id
		setMaxFullScreenContentSize: value];
}

NSSize NSWindow_inst_minFullScreenContentSize(void *id) {
	return [(NSWindow*)id
		minFullScreenContentSize];
}

void NSWindow_inst_setMinFullScreenContentSize(void *id, NSSize value) {
	[(NSWindow*)id
		setMinFullScreenContentSize: value];
}

long NSWindow_inst_level(void *id) {
	return [(NSWindow*)id
		level];
}

void NSWindow_inst_setLevel(void *id, long value) {
	[(NSWindow*)id
		setLevel: value];
}

BOOL NSWindow_inst_isVisible(void *id) {
	return [(NSWindow*)id
		isVisible];
}

BOOL NSWindow_inst_isKeyWindow(void *id) {
	return [(NSWindow*)id
		isKeyWindow];
}

BOOL NSWindow_inst_canBecomeKeyWindow(void *id) {
	return [(NSWindow*)id
		canBecomeKeyWindow];
}

BOOL NSWindow_inst_isMainWindow(void *id) {
	return [(NSWindow*)id
		isMainWindow];
}

BOOL NSWindow_inst_canBecomeMainWindow(void *id) {
	return [(NSWindow*)id
		canBecomeMainWindow];
}

void* NSWindow_inst_childWindows(void *id) {
	return [(NSWindow*)id
		childWindows];
}

void* NSWindow_inst_parentWindow(void *id) {
	return [(NSWindow*)id
		parentWindow];
}

void NSWindow_inst_setParentWindow(void *id, void* value) {
	[(NSWindow*)id
		setParentWindow: value];
}

BOOL NSWindow_inst_isExcludedFromWindowsMenu(void *id) {
	return [(NSWindow*)id
		isExcludedFromWindowsMenu];
}

void NSWindow_inst_setExcludedFromWindowsMenu(void *id, BOOL value) {
	[(NSWindow*)id
		setExcludedFromWindowsMenu: value];
}

BOOL NSWindow_inst_areCursorRectsEnabled(void *id) {
	return [(NSWindow*)id
		areCursorRectsEnabled];
}

BOOL NSWindow_inst_showsToolbarButton(void *id) {
	return [(NSWindow*)id
		showsToolbarButton];
}

void NSWindow_inst_setShowsToolbarButton(void *id, BOOL value) {
	[(NSWindow*)id
		setShowsToolbarButton: value];
}

BOOL NSWindow_inst_titlebarAppearsTransparent(void *id) {
	return [(NSWindow*)id
		titlebarAppearsTransparent];
}

void NSWindow_inst_setTitlebarAppearsTransparent(void *id, BOOL value) {
	[(NSWindow*)id
		setTitlebarAppearsTransparent: value];
}

void* NSWindow_inst_titlebarAccessoryViewControllers(void *id) {
	return [(NSWindow*)id
		titlebarAccessoryViewControllers];
}

void NSWindow_inst_setTitlebarAccessoryViewControllers(void *id, void* value) {
	[(NSWindow*)id
		setTitlebarAccessoryViewControllers: value];
}

void* NSWindow_inst_tabbedWindows(void *id) {
	return [(NSWindow*)id
		tabbedWindows];
}

BOOL NSWindow_inst_allowsToolTipsWhenApplicationIsInactive(void *id) {
	return [(NSWindow*)id
		allowsToolTipsWhenApplicationIsInactive];
}

void NSWindow_inst_setAllowsToolTipsWhenApplicationIsInactive(void *id, BOOL value) {
	[(NSWindow*)id
		setAllowsToolTipsWhenApplicationIsInactive: value];
}

void* NSWindow_inst_currentEvent(void *id) {
	return [(NSWindow*)id
		currentEvent];
}

void* NSWindow_inst_initialFirstResponder(void *id) {
	return [(NSWindow*)id
		initialFirstResponder];
}

void NSWindow_inst_setInitialFirstResponder(void *id, void* value) {
	[(NSWindow*)id
		setInitialFirstResponder: value];
}

BOOL NSWindow_inst_autorecalculatesKeyViewLoop(void *id) {
	return [(NSWindow*)id
		autorecalculatesKeyViewLoop];
}

void NSWindow_inst_setAutorecalculatesKeyViewLoop(void *id, BOOL value) {
	[(NSWindow*)id
		setAutorecalculatesKeyViewLoop: value];
}

BOOL NSWindow_inst_acceptsMouseMovedEvents(void *id) {
	return [(NSWindow*)id
		acceptsMouseMovedEvents];
}

void NSWindow_inst_setAcceptsMouseMovedEvents(void *id, BOOL value) {
	[(NSWindow*)id
		setAcceptsMouseMovedEvents: value];
}

BOOL NSWindow_inst_ignoresMouseEvents(void *id) {
	return [(NSWindow*)id
		ignoresMouseEvents];
}

void NSWindow_inst_setIgnoresMouseEvents(void *id, BOOL value) {
	[(NSWindow*)id
		setIgnoresMouseEvents: value];
}

NSPoint NSWindow_inst_mouseLocationOutsideOfEventStream(void *id) {
	return [(NSWindow*)id
		mouseLocationOutsideOfEventStream];
}

BOOL NSWindow_inst_isRestorable(void *id) {
	return [(NSWindow*)id
		isRestorable];
}

void NSWindow_inst_setRestorable(void *id, BOOL value) {
	[(NSWindow*)id
		setRestorable: value];
}

BOOL NSWindow_inst_viewsNeedDisplay(void *id) {
	return [(NSWindow*)id
		viewsNeedDisplay];
}

void NSWindow_inst_setViewsNeedDisplay(void *id, BOOL value) {
	[(NSWindow*)id
		setViewsNeedDisplay: value];
}

BOOL NSWindow_inst_allowsConcurrentViewDrawing(void *id) {
	return [(NSWindow*)id
		allowsConcurrentViewDrawing];
}

void NSWindow_inst_setAllowsConcurrentViewDrawing(void *id, BOOL value) {
	[(NSWindow*)id
		setAllowsConcurrentViewDrawing: value];
}

BOOL NSWindow_inst_isDocumentEdited(void *id) {
	return [(NSWindow*)id
		isDocumentEdited];
}

void NSWindow_inst_setDocumentEdited(void *id, BOOL value) {
	[(NSWindow*)id
		setDocumentEdited: value];
}

double NSWindow_inst_backingScaleFactor(void *id) {
	return [(NSWindow*)id
		backingScaleFactor];
}

void* NSWindow_inst_title(void *id) {
	return [(NSWindow*)id
		title];
}

void NSWindow_inst_setTitle(void *id, void* value) {
	[(NSWindow*)id
		setTitle: value];
}

void* NSWindow_inst_subtitle(void *id) {
	return [(NSWindow*)id
		subtitle];
}

void NSWindow_inst_setSubtitle(void *id, void* value) {
	[(NSWindow*)id
		setSubtitle: value];
}

long NSWindow_inst_titleVisibility(void *id) {
	return [(NSWindow*)id
		titleVisibility];
}

void NSWindow_inst_setTitleVisibility(void *id, long value) {
	[(NSWindow*)id
		setTitleVisibility: value];
}

void* NSWindow_inst_representedFilename(void *id) {
	return [(NSWindow*)id
		representedFilename];
}

void NSWindow_inst_setRepresentedFilename(void *id, void* value) {
	[(NSWindow*)id
		setRepresentedFilename: value];
}

void* NSWindow_inst_representedURL(void *id) {
	return [(NSWindow*)id
		representedURL];
}

void NSWindow_inst_setRepresentedURL(void *id, void* value) {
	[(NSWindow*)id
		setRepresentedURL: value];
}

void* NSWindow_inst_screen(void *id) {
	return [(NSWindow*)id
		screen];
}

void* NSWindow_inst_deepestScreen(void *id) {
	return [(NSWindow*)id
		deepestScreen];
}

BOOL NSWindow_inst_displaysWhenScreenProfileChanges(void *id) {
	return [(NSWindow*)id
		displaysWhenScreenProfileChanges];
}

void NSWindow_inst_setDisplaysWhenScreenProfileChanges(void *id, BOOL value) {
	[(NSWindow*)id
		setDisplaysWhenScreenProfileChanges: value];
}

BOOL NSWindow_inst_isMovableByWindowBackground(void *id) {
	return [(NSWindow*)id
		isMovableByWindowBackground];
}

void NSWindow_inst_setMovableByWindowBackground(void *id, BOOL value) {
	[(NSWindow*)id
		setMovableByWindowBackground: value];
}

BOOL NSWindow_inst_isMovable(void *id) {
	return [(NSWindow*)id
		isMovable];
}

void NSWindow_inst_setMovable(void *id, BOOL value) {
	[(NSWindow*)id
		setMovable: value];
}

BOOL NSWindow_inst_isReleasedWhenClosed(void *id) {
	return [(NSWindow*)id
		isReleasedWhenClosed];
}

void NSWindow_inst_setReleasedWhenClosed(void *id, BOOL value) {
	[(NSWindow*)id
		setReleasedWhenClosed: value];
}

BOOL NSWindow_inst_isMiniaturized(void *id) {
	return [(NSWindow*)id
		isMiniaturized];
}

void* NSWindow_inst_miniwindowImage(void *id) {
	return [(NSWindow*)id
		miniwindowImage];
}

void NSWindow_inst_setMiniwindowImage(void *id, void* value) {
	[(NSWindow*)id
		setMiniwindowImage: value];
}

void* NSWindow_inst_miniwindowTitle(void *id) {
	return [(NSWindow*)id
		miniwindowTitle];
}

void NSWindow_inst_setMiniwindowTitle(void *id, void* value) {
	[(NSWindow*)id
		setMiniwindowTitle: value];
}

BOOL NSWindow_inst_hasCloseBox(void *id) {
	return [(NSWindow*)id
		hasCloseBox];
}

BOOL NSWindow_inst_hasTitleBar(void *id) {
	return [(NSWindow*)id
		hasTitleBar];
}

BOOL NSWindow_inst_isModalPanel(void *id) {
	return [(NSWindow*)id
		isModalPanel];
}

BOOL NSWindow_inst_isFloatingPanel(void *id) {
	return [(NSWindow*)id
		isFloatingPanel];
}

BOOL NSWindow_inst_isZoomable(void *id) {
	return [(NSWindow*)id
		isZoomable];
}

BOOL NSWindow_inst_isResizable(void *id) {
	return [(NSWindow*)id
		isResizable];
}

BOOL NSWindow_inst_isMiniaturizable(void *id) {
	return [(NSWindow*)id
		isMiniaturizable];
}

long NSWindow_inst_orderedIndex(void *id) {
	return [(NSWindow*)id
		orderedIndex];
}

void NSWindow_inst_setOrderedIndex(void *id, long value) {
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

void NSWorkspace_inst_activateFileViewerSelectingURLs(void *id, void* fileURLs) {
	[(NSWorkspace*)id
		activateFileViewerSelectingURLs: fileURLs];
}

void* NSWorkspace_inst_desktopImageOptionsForScreen(void *id, void* screen) {
	return [(NSWorkspace*)id
		desktopImageOptionsForScreen: screen];
}

void* NSWorkspace_inst_desktopImageURLForScreen(void *id, void* screen) {
	return [(NSWorkspace*)id
		desktopImageURLForScreen: screen];
}

long NSWorkspace_inst_extendPowerOffBy(void *id, long requested) {
	return [(NSWorkspace*)id
		extendPowerOffBy: requested];
}

void NSWorkspace_inst_hideOtherApplications(void *id) {
	[(NSWorkspace*)id
		hideOtherApplications];
}

void* NSWorkspace_inst_iconForFile(void *id, void* fullPath) {
	return [(NSWorkspace*)id
		iconForFile: fullPath];
}

void* NSWorkspace_inst_iconForFiles(void *id, void* fullPaths) {
	return [(NSWorkspace*)id
		iconForFiles: fullPaths];
}

BOOL NSWorkspace_inst_isFilePackageAtPath(void *id, void* fullPath) {
	return [(NSWorkspace*)id
		isFilePackageAtPath: fullPath];
}

void NSWorkspace_inst_noteFileSystemChanged(void *id, void* path) {
	[(NSWorkspace*)id
		noteFileSystemChanged: path];
}

BOOL NSWorkspace_inst_openURL(void *id, void* url) {
	return [(NSWorkspace*)id
		openURL: url];
}

BOOL NSWorkspace_inst_selectFile_inFileViewerRootedAtPath(void *id, void* fullPath, void* rootFullPath) {
	return [(NSWorkspace*)id
		selectFile: fullPath
		inFileViewerRootedAtPath: rootFullPath];
}

BOOL NSWorkspace_inst_showSearchResultsForQueryString(void *id, void* queryString) {
	return [(NSWorkspace*)id
		showSearchResultsForQueryString: queryString];
}

BOOL NSWorkspace_inst_unmountAndEjectDeviceAtPath(void *id, void* path) {
	return [(NSWorkspace*)id
		unmountAndEjectDeviceAtPath: path];
}

void* NSWorkspace_inst_init(void *id) {
	return [(NSWorkspace*)id
		init];
}

void* NSWorkspace_inst_frontmostApplication(void *id) {
	return [(NSWorkspace*)id
		frontmostApplication];
}

void* NSWorkspace_inst_runningApplications(void *id) {
	return [(NSWorkspace*)id
		runningApplications];
}

void* NSWorkspace_inst_menuBarOwningApplication(void *id) {
	return [(NSWorkspace*)id
		menuBarOwningApplication];
}

void* NSWorkspace_inst_fileLabels(void *id) {
	return [(NSWorkspace*)id
		fileLabels];
}

void* NSWorkspace_inst_fileLabelColors(void *id) {
	return [(NSWorkspace*)id
		fileLabelColors];
}

BOOL NSWorkspace_inst_accessibilityDisplayShouldDifferentiateWithoutColor(void *id) {
	return [(NSWorkspace*)id
		accessibilityDisplayShouldDifferentiateWithoutColor];
}

BOOL NSWorkspace_inst_accessibilityDisplayShouldIncreaseContrast(void *id) {
	return [(NSWorkspace*)id
		accessibilityDisplayShouldIncreaseContrast];
}

BOOL NSWorkspace_inst_accessibilityDisplayShouldReduceTransparency(void *id) {
	return [(NSWorkspace*)id
		accessibilityDisplayShouldReduceTransparency];
}

BOOL NSWorkspace_inst_accessibilityDisplayShouldInvertColors(void *id) {
	return [(NSWorkspace*)id
		accessibilityDisplayShouldInvertColors];
}

BOOL NSWorkspace_inst_accessibilityDisplayShouldReduceMotion(void *id) {
	return [(NSWorkspace*)id
		accessibilityDisplayShouldReduceMotion];
}

BOOL NSWorkspace_inst_isSwitchControlEnabled(void *id) {
	return [(NSWorkspace*)id
		isSwitchControlEnabled];
}

BOOL NSWorkspace_inst_isVoiceOverEnabled(void *id) {
	return [(NSWorkspace*)id
		isVoiceOverEnabled];
}

void* NSColor_inst_blendedColorWithFraction_ofColor(void *id, double fraction, void* color) {
	return [(NSColor*)id
		blendedColorWithFraction: fraction
		ofColor: color];
}

void* NSColor_inst_colorWithAlphaComponent(void *id, double alpha) {
	return [(NSColor*)id
		colorWithAlphaComponent: alpha];
}

void NSColor_inst_drawSwatchInRect(void *id, NSRect rect) {
	[(NSColor*)id
		drawSwatchInRect: rect];
}

void* NSColor_inst_highlightWithLevel(void *id, double val) {
	return [(NSColor*)id
		highlightWithLevel: val];
}

void NSColor_inst_set(void *id) {
	[(NSColor*)id
		set];
}

void NSColor_inst_setFill(void *id) {
	[(NSColor*)id
		setFill];
}

void NSColor_inst_setStroke(void *id) {
	[(NSColor*)id
		setStroke];
}

void* NSColor_inst_shadowWithLevel(void *id, double val) {
	return [(NSColor*)id
		shadowWithLevel: val];
}

void NSColor_inst_writeToPasteboard(void *id, void* pasteBoard) {
	[(NSColor*)id
		writeToPasteboard: pasteBoard];
}

void* NSColor_inst_init(void *id) {
	return [(NSColor*)id
		init];
}

long NSColor_inst_numberOfComponents(void *id) {
	return [(NSColor*)id
		numberOfComponents];
}

double NSColor_inst_alphaComponent(void *id) {
	return [(NSColor*)id
		alphaComponent];
}

double NSColor_inst_whiteComponent(void *id) {
	return [(NSColor*)id
		whiteComponent];
}

double NSColor_inst_redComponent(void *id) {
	return [(NSColor*)id
		redComponent];
}

double NSColor_inst_greenComponent(void *id) {
	return [(NSColor*)id
		greenComponent];
}

double NSColor_inst_blueComponent(void *id) {
	return [(NSColor*)id
		blueComponent];
}

double NSColor_inst_cyanComponent(void *id) {
	return [(NSColor*)id
		cyanComponent];
}

double NSColor_inst_magentaComponent(void *id) {
	return [(NSColor*)id
		magentaComponent];
}

double NSColor_inst_yellowComponent(void *id) {
	return [(NSColor*)id
		yellowComponent];
}

double NSColor_inst_blackComponent(void *id) {
	return [(NSColor*)id
		blackComponent];
}

double NSColor_inst_hueComponent(void *id) {
	return [(NSColor*)id
		hueComponent];
}

double NSColor_inst_saturationComponent(void *id) {
	return [(NSColor*)id
		saturationComponent];
}

double NSColor_inst_brightnessComponent(void *id) {
	return [(NSColor*)id
		brightnessComponent];
}

void* NSColor_inst_localizedCatalogNameComponent(void *id) {
	return [(NSColor*)id
		localizedCatalogNameComponent];
}

void* NSColor_inst_localizedColorNameComponent(void *id) {
	return [(NSColor*)id
		localizedColorNameComponent];
}

void NSTextView_inst_alignJustified(void *id, void* sender) {
	[(NSTextView*)id
		alignJustified: sender];
}

void NSTextView_inst_breakUndoCoalescing(void *id) {
	[(NSTextView*)id
		breakUndoCoalescing];
}

void NSTextView_inst_changeAttributes(void *id, void* sender) {
	[(NSTextView*)id
		changeAttributes: sender];
}

void NSTextView_inst_changeColor(void *id, void* sender) {
	[(NSTextView*)id
		changeColor: sender];
}

void NSTextView_inst_changeDocumentBackgroundColor(void *id, void* sender) {
	[(NSTextView*)id
		changeDocumentBackgroundColor: sender];
}

void NSTextView_inst_changeLayoutOrientation(void *id, void* sender) {
	[(NSTextView*)id
		changeLayoutOrientation: sender];
}

unsigned long NSTextView_inst_characterIndexForInsertionAtPoint(void *id, NSPoint point) {
	return [(NSTextView*)id
		characterIndexForInsertionAtPoint: point];
}

void NSTextView_inst_checkTextInDocument(void *id, void* sender) {
	[(NSTextView*)id
		checkTextInDocument: sender];
}

void NSTextView_inst_checkTextInSelection(void *id, void* sender) {
	[(NSTextView*)id
		checkTextInSelection: sender];
}

void NSTextView_inst_cleanUpAfterDragOperation(void *id) {
	[(NSTextView*)id
		cleanUpAfterDragOperation];
}

void NSTextView_inst_clickedOnLink_atIndex(void *id, void* link, unsigned long charIndex) {
	[(NSTextView*)id
		clickedOnLink: link
		atIndex: charIndex];
}

void NSTextView_inst_complete(void *id, void* sender) {
	[(NSTextView*)id
		complete: sender];
}

void NSTextView_inst_didChangeText(void *id) {
	[(NSTextView*)id
		didChangeText];
}

BOOL NSTextView_inst_dragSelectionWithEvent_offset_slideBack(void *id, void* event, NSSize mouseOffset, BOOL slideBack) {
	return [(NSTextView*)id
		dragSelectionWithEvent: event
		offset: mouseOffset
		slideBack: slideBack];
}

void NSTextView_inst_drawInsertionPointInRect_color_turnedOn(void *id, NSRect rect, void* color, BOOL flag) {
	[(NSTextView*)id
		drawInsertionPointInRect: rect
		color: color
		turnedOn: flag];
}

void NSTextView_inst_drawViewBackgroundInRect(void *id, NSRect rect) {
	[(NSTextView*)id
		drawViewBackgroundInRect: rect];
}

void* NSTextView_inst_initWithFrame(void *id, NSRect frameRect) {
	return [(NSTextView*)id
		initWithFrame: frameRect];
}

void* NSTextView_inst_initWithFrame_textContainer(void *id, NSRect frameRect, void* container) {
	return [(NSTextView*)id
		initWithFrame: frameRect
		textContainer: container];
}

void NSTextView_inst_invalidateTextContainerOrigin(void *id) {
	[(NSTextView*)id
		invalidateTextContainerOrigin];
}

void NSTextView_inst_loosenKerning(void *id, void* sender) {
	[(NSTextView*)id
		loosenKerning: sender];
}

void NSTextView_inst_lowerBaseline(void *id, void* sender) {
	[(NSTextView*)id
		lowerBaseline: sender];
}

void NSTextView_inst_orderFrontLinkPanel(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontLinkPanel: sender];
}

void NSTextView_inst_orderFrontListPanel(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontListPanel: sender];
}

void NSTextView_inst_orderFrontSharingServicePicker(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontSharingServicePicker: sender];
}

void NSTextView_inst_orderFrontSpacingPanel(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontSpacingPanel: sender];
}

void NSTextView_inst_orderFrontSubstitutionsPanel(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontSubstitutionsPanel: sender];
}

void NSTextView_inst_orderFrontTablePanel(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontTablePanel: sender];
}

void NSTextView_inst_outline(void *id, void* sender) {
	[(NSTextView*)id
		outline: sender];
}

void NSTextView_inst_pasteAsPlainText(void *id, void* sender) {
	[(NSTextView*)id
		pasteAsPlainText: sender];
}

void NSTextView_inst_pasteAsRichText(void *id, void* sender) {
	[(NSTextView*)id
		pasteAsRichText: sender];
}

void NSTextView_inst_performFindPanelAction(void *id, void* sender) {
	[(NSTextView*)id
		performFindPanelAction: sender];
}

void* NSTextView_inst_quickLookPreviewableItemsInRanges(void *id, void* ranges) {
	return [(NSTextView*)id
		quickLookPreviewableItemsInRanges: ranges];
}

void NSTextView_inst_raiseBaseline(void *id, void* sender) {
	[(NSTextView*)id
		raiseBaseline: sender];
}

BOOL NSTextView_inst_readSelectionFromPasteboard(void *id, void* pboard) {
	return [(NSTextView*)id
		readSelectionFromPasteboard: pboard];
}

void NSTextView_inst_replaceTextContainer(void *id, void* newContainer) {
	[(NSTextView*)id
		replaceTextContainer: newContainer];
}

void NSTextView_inst_setConstrainedFrameSize(void *id, NSSize desiredSize) {
	[(NSTextView*)id
		setConstrainedFrameSize: desiredSize];
}

void NSTextView_inst_setNeedsDisplayInRect_avoidAdditionalLayout(void *id, NSRect rect, BOOL flag) {
	[(NSTextView*)id
		setNeedsDisplayInRect: rect
		avoidAdditionalLayout: flag];
}

BOOL NSTextView_inst_shouldChangeTextInRanges_replacementStrings(void *id, void* affectedRanges, void* replacementStrings) {
	return [(NSTextView*)id
		shouldChangeTextInRanges: affectedRanges
		replacementStrings: replacementStrings];
}

void NSTextView_inst_startSpeaking(void *id, void* sender) {
	[(NSTextView*)id
		startSpeaking: sender];
}

void NSTextView_inst_stopSpeaking(void *id, void* sender) {
	[(NSTextView*)id
		stopSpeaking: sender];
}

void NSTextView_inst_tightenKerning(void *id, void* sender) {
	[(NSTextView*)id
		tightenKerning: sender];
}

void NSTextView_inst_toggleAutomaticDashSubstitution(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticDashSubstitution: sender];
}

void NSTextView_inst_toggleAutomaticDataDetection(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticDataDetection: sender];
}

void NSTextView_inst_toggleAutomaticLinkDetection(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticLinkDetection: sender];
}

void NSTextView_inst_toggleAutomaticQuoteSubstitution(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticQuoteSubstitution: sender];
}

void NSTextView_inst_toggleAutomaticSpellingCorrection(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticSpellingCorrection: sender];
}

void NSTextView_inst_toggleAutomaticTextCompletion(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticTextCompletion: sender];
}

void NSTextView_inst_toggleAutomaticTextReplacement(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticTextReplacement: sender];
}

void NSTextView_inst_toggleContinuousSpellChecking(void *id, void* sender) {
	[(NSTextView*)id
		toggleContinuousSpellChecking: sender];
}

void NSTextView_inst_toggleGrammarChecking(void *id, void* sender) {
	[(NSTextView*)id
		toggleGrammarChecking: sender];
}

void NSTextView_inst_toggleQuickLookPreviewPanel(void *id, void* sender) {
	[(NSTextView*)id
		toggleQuickLookPreviewPanel: sender];
}

void NSTextView_inst_toggleSmartInsertDelete(void *id, void* sender) {
	[(NSTextView*)id
		toggleSmartInsertDelete: sender];
}

void NSTextView_inst_turnOffKerning(void *id, void* sender) {
	[(NSTextView*)id
		turnOffKerning: sender];
}

void NSTextView_inst_turnOffLigatures(void *id, void* sender) {
	[(NSTextView*)id
		turnOffLigatures: sender];
}

void NSTextView_inst_updateCandidates(void *id) {
	[(NSTextView*)id
		updateCandidates];
}

void NSTextView_inst_updateDragTypeRegistration(void *id) {
	[(NSTextView*)id
		updateDragTypeRegistration];
}

void NSTextView_inst_updateFontPanel(void *id) {
	[(NSTextView*)id
		updateFontPanel];
}

void NSTextView_inst_updateInsertionPointStateAndRestartTimer(void *id, BOOL restartFlag) {
	[(NSTextView*)id
		updateInsertionPointStateAndRestartTimer: restartFlag];
}

void NSTextView_inst_updateQuickLookPreviewPanel(void *id) {
	[(NSTextView*)id
		updateQuickLookPreviewPanel];
}

void NSTextView_inst_updateRuler(void *id) {
	[(NSTextView*)id
		updateRuler];
}

void NSTextView_inst_updateTextTouchBarItems(void *id) {
	[(NSTextView*)id
		updateTextTouchBarItems];
}

void NSTextView_inst_updateTouchBarItemIdentifiers(void *id) {
	[(NSTextView*)id
		updateTouchBarItemIdentifiers];
}

void NSTextView_inst_useAllLigatures(void *id, void* sender) {
	[(NSTextView*)id
		useAllLigatures: sender];
}

void NSTextView_inst_useStandardKerning(void *id, void* sender) {
	[(NSTextView*)id
		useStandardKerning: sender];
}

void NSTextView_inst_useStandardLigatures(void *id, void* sender) {
	[(NSTextView*)id
		useStandardLigatures: sender];
}

BOOL NSTextView_inst_writeSelectionToPasteboard_types(void *id, void* pboard, void* types) {
	return [(NSTextView*)id
		writeSelectionToPasteboard: pboard
		types: types];
}

void* NSTextView_inst_init(void *id) {
	return [(NSTextView*)id
		init];
}

void* NSTextView_inst_delegate(void *id) {
	return [(NSTextView*)id
		delegate];
}

void NSTextView_inst_setDelegate(void *id, void* value) {
	[(NSTextView*)id
		setDelegate: value];
}

void* NSTextView_inst_textContainer(void *id) {
	return [(NSTextView*)id
		textContainer];
}

void NSTextView_inst_setTextContainer(void *id, void* value) {
	[(NSTextView*)id
		setTextContainer: value];
}

NSSize NSTextView_inst_textContainerInset(void *id) {
	return [(NSTextView*)id
		textContainerInset];
}

void NSTextView_inst_setTextContainerInset(void *id, NSSize value) {
	[(NSTextView*)id
		setTextContainerInset: value];
}

NSPoint NSTextView_inst_textContainerOrigin(void *id) {
	return [(NSTextView*)id
		textContainerOrigin];
}

void* NSTextView_inst_layoutManager(void *id) {
	return [(NSTextView*)id
		layoutManager];
}

void* NSTextView_inst_backgroundColor(void *id) {
	return [(NSTextView*)id
		backgroundColor];
}

void NSTextView_inst_setBackgroundColor(void *id, void* value) {
	[(NSTextView*)id
		setBackgroundColor: value];
}

BOOL NSTextView_inst_drawsBackground(void *id) {
	return [(NSTextView*)id
		drawsBackground];
}

void NSTextView_inst_setDrawsBackground(void *id, BOOL value) {
	[(NSTextView*)id
		setDrawsBackground: value];
}

BOOL NSTextView_inst_allowsDocumentBackgroundColorChange(void *id) {
	return [(NSTextView*)id
		allowsDocumentBackgroundColorChange];
}

void NSTextView_inst_setAllowsDocumentBackgroundColorChange(void *id, BOOL value) {
	[(NSTextView*)id
		setAllowsDocumentBackgroundColorChange: value];
}

BOOL NSTextView_inst_shouldDrawInsertionPoint(void *id) {
	return [(NSTextView*)id
		shouldDrawInsertionPoint];
}

void* NSTextView_inst_allowedInputSourceLocales(void *id) {
	return [(NSTextView*)id
		allowedInputSourceLocales];
}

void NSTextView_inst_setAllowedInputSourceLocales(void *id, void* value) {
	[(NSTextView*)id
		setAllowedInputSourceLocales: value];
}

BOOL NSTextView_inst_allowsUndo(void *id) {
	return [(NSTextView*)id
		allowsUndo];
}

void NSTextView_inst_setAllowsUndo(void *id, BOOL value) {
	[(NSTextView*)id
		setAllowsUndo: value];
}

BOOL NSTextView_inst_isEditable(void *id) {
	return [(NSTextView*)id
		isEditable];
}

void NSTextView_inst_setEditable(void *id, BOOL value) {
	[(NSTextView*)id
		setEditable: value];
}

BOOL NSTextView_inst_isSelectable(void *id) {
	return [(NSTextView*)id
		isSelectable];
}

void NSTextView_inst_setSelectable(void *id, BOOL value) {
	[(NSTextView*)id
		setSelectable: value];
}

BOOL NSTextView_inst_isFieldEditor(void *id) {
	return [(NSTextView*)id
		isFieldEditor];
}

void NSTextView_inst_setFieldEditor(void *id, BOOL value) {
	[(NSTextView*)id
		setFieldEditor: value];
}

BOOL NSTextView_inst_isRichText(void *id) {
	return [(NSTextView*)id
		isRichText];
}

void NSTextView_inst_setRichText(void *id, BOOL value) {
	[(NSTextView*)id
		setRichText: value];
}

BOOL NSTextView_inst_importsGraphics(void *id) {
	return [(NSTextView*)id
		importsGraphics];
}

void NSTextView_inst_setImportsGraphics(void *id, BOOL value) {
	[(NSTextView*)id
		setImportsGraphics: value];
}

BOOL NSTextView_inst_allowsImageEditing(void *id) {
	return [(NSTextView*)id
		allowsImageEditing];
}

void NSTextView_inst_setAllowsImageEditing(void *id, BOOL value) {
	[(NSTextView*)id
		setAllowsImageEditing: value];
}

BOOL NSTextView_inst_isAutomaticQuoteSubstitutionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticQuoteSubstitutionEnabled];
}

void NSTextView_inst_setAutomaticQuoteSubstitutionEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticQuoteSubstitutionEnabled: value];
}

BOOL NSTextView_inst_isAutomaticLinkDetectionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticLinkDetectionEnabled];
}

void NSTextView_inst_setAutomaticLinkDetectionEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticLinkDetectionEnabled: value];
}

BOOL NSTextView_inst_displaysLinkToolTips(void *id) {
	return [(NSTextView*)id
		displaysLinkToolTips];
}

void NSTextView_inst_setDisplaysLinkToolTips(void *id, BOOL value) {
	[(NSTextView*)id
		setDisplaysLinkToolTips: value];
}

BOOL NSTextView_inst_isAutomaticTextCompletionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticTextCompletionEnabled];
}

void NSTextView_inst_setAutomaticTextCompletionEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticTextCompletionEnabled: value];
}

BOOL NSTextView_inst_usesAdaptiveColorMappingForDarkAppearance(void *id) {
	return [(NSTextView*)id
		usesAdaptiveColorMappingForDarkAppearance];
}

void NSTextView_inst_setUsesAdaptiveColorMappingForDarkAppearance(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesAdaptiveColorMappingForDarkAppearance: value];
}

BOOL NSTextView_inst_usesRolloverButtonForSelection(void *id) {
	return [(NSTextView*)id
		usesRolloverButtonForSelection];
}

void NSTextView_inst_setUsesRolloverButtonForSelection(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesRolloverButtonForSelection: value];
}

BOOL NSTextView_inst_usesRuler(void *id) {
	return [(NSTextView*)id
		usesRuler];
}

void NSTextView_inst_setUsesRuler(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesRuler: value];
}

BOOL NSTextView_inst_isRulerVisible(void *id) {
	return [(NSTextView*)id
		isRulerVisible];
}

void NSTextView_inst_setRulerVisible(void *id, BOOL value) {
	[(NSTextView*)id
		setRulerVisible: value];
}

BOOL NSTextView_inst_usesInspectorBar(void *id) {
	return [(NSTextView*)id
		usesInspectorBar];
}

void NSTextView_inst_setUsesInspectorBar(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesInspectorBar: value];
}

void* NSTextView_inst_selectedRanges(void *id) {
	return [(NSTextView*)id
		selectedRanges];
}

void NSTextView_inst_setSelectedRanges(void *id, void* value) {
	[(NSTextView*)id
		setSelectedRanges: value];
}

void* NSTextView_inst_insertionPointColor(void *id) {
	return [(NSTextView*)id
		insertionPointColor];
}

void NSTextView_inst_setInsertionPointColor(void *id, void* value) {
	[(NSTextView*)id
		setInsertionPointColor: value];
}

void* NSTextView_inst_selectedTextAttributes(void *id) {
	return [(NSTextView*)id
		selectedTextAttributes];
}

void NSTextView_inst_setSelectedTextAttributes(void *id, void* value) {
	[(NSTextView*)id
		setSelectedTextAttributes: value];
}

void* NSTextView_inst_markedTextAttributes(void *id) {
	return [(NSTextView*)id
		markedTextAttributes];
}

void NSTextView_inst_setMarkedTextAttributes(void *id, void* value) {
	[(NSTextView*)id
		setMarkedTextAttributes: value];
}

void* NSTextView_inst_linkTextAttributes(void *id) {
	return [(NSTextView*)id
		linkTextAttributes];
}

void NSTextView_inst_setLinkTextAttributes(void *id, void* value) {
	[(NSTextView*)id
		setLinkTextAttributes: value];
}

void* NSTextView_inst_readablePasteboardTypes(void *id) {
	return [(NSTextView*)id
		readablePasteboardTypes];
}

void* NSTextView_inst_writablePasteboardTypes(void *id) {
	return [(NSTextView*)id
		writablePasteboardTypes];
}

void* NSTextView_inst_typingAttributes(void *id) {
	return [(NSTextView*)id
		typingAttributes];
}

void NSTextView_inst_setTypingAttributes(void *id, void* value) {
	[(NSTextView*)id
		setTypingAttributes: value];
}

BOOL NSTextView_inst_isCoalescingUndo(void *id) {
	return [(NSTextView*)id
		isCoalescingUndo];
}

void* NSTextView_inst_acceptableDragTypes(void *id) {
	return [(NSTextView*)id
		acceptableDragTypes];
}

void* NSTextView_inst_rangesForUserCharacterAttributeChange(void *id) {
	return [(NSTextView*)id
		rangesForUserCharacterAttributeChange];
}

void* NSTextView_inst_rangesForUserParagraphAttributeChange(void *id) {
	return [(NSTextView*)id
		rangesForUserParagraphAttributeChange];
}

void* NSTextView_inst_rangesForUserTextChange(void *id) {
	return [(NSTextView*)id
		rangesForUserTextChange];
}

BOOL NSTextView_inst_smartInsertDeleteEnabled(void *id) {
	return [(NSTextView*)id
		smartInsertDeleteEnabled];
}

void NSTextView_inst_setSmartInsertDeleteEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setSmartInsertDeleteEnabled: value];
}

BOOL NSTextView_inst_isContinuousSpellCheckingEnabled(void *id) {
	return [(NSTextView*)id
		isContinuousSpellCheckingEnabled];
}

void NSTextView_inst_setContinuousSpellCheckingEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setContinuousSpellCheckingEnabled: value];
}

long NSTextView_inst_spellCheckerDocumentTag(void *id) {
	return [(NSTextView*)id
		spellCheckerDocumentTag];
}

BOOL NSTextView_inst_isGrammarCheckingEnabled(void *id) {
	return [(NSTextView*)id
		isGrammarCheckingEnabled];
}

void NSTextView_inst_setGrammarCheckingEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setGrammarCheckingEnabled: value];
}

BOOL NSTextView_inst_acceptsGlyphInfo(void *id) {
	return [(NSTextView*)id
		acceptsGlyphInfo];
}

void NSTextView_inst_setAcceptsGlyphInfo(void *id, BOOL value) {
	[(NSTextView*)id
		setAcceptsGlyphInfo: value];
}

BOOL NSTextView_inst_usesFontPanel(void *id) {
	return [(NSTextView*)id
		usesFontPanel];
}

void NSTextView_inst_setUsesFontPanel(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesFontPanel: value];
}

BOOL NSTextView_inst_usesFindPanel(void *id) {
	return [(NSTextView*)id
		usesFindPanel];
}

void NSTextView_inst_setUsesFindPanel(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesFindPanel: value];
}

BOOL NSTextView_inst_isAutomaticDashSubstitutionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticDashSubstitutionEnabled];
}

void NSTextView_inst_setAutomaticDashSubstitutionEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticDashSubstitutionEnabled: value];
}

BOOL NSTextView_inst_isAutomaticDataDetectionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticDataDetectionEnabled];
}

void NSTextView_inst_setAutomaticDataDetectionEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticDataDetectionEnabled: value];
}

BOOL NSTextView_inst_isAutomaticSpellingCorrectionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticSpellingCorrectionEnabled];
}

void NSTextView_inst_setAutomaticSpellingCorrectionEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticSpellingCorrectionEnabled: value];
}

BOOL NSTextView_inst_isAutomaticTextReplacementEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticTextReplacementEnabled];
}

void NSTextView_inst_setAutomaticTextReplacementEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticTextReplacementEnabled: value];
}

BOOL NSTextView_inst_usesFindBar(void *id) {
	return [(NSTextView*)id
		usesFindBar];
}

void NSTextView_inst_setUsesFindBar(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesFindBar: value];
}

BOOL NSTextView_inst_isIncrementalSearchingEnabled(void *id) {
	return [(NSTextView*)id
		isIncrementalSearchingEnabled];
}

void NSTextView_inst_setIncrementalSearchingEnabled(void *id, BOOL value) {
	[(NSTextView*)id
		setIncrementalSearchingEnabled: value];
}

BOOL NSTextView_inst_allowsCharacterPickerTouchBarItem(void *id) {
	return [(NSTextView*)id
		allowsCharacterPickerTouchBarItem];
}

void NSTextView_inst_setAllowsCharacterPickerTouchBarItem(void *id, BOOL value) {
	[(NSTextView*)id
		setAllowsCharacterPickerTouchBarItem: value];
}

void* NSTextView_inst_font(void *id) {
	return [(NSTextView*)id
		font];
}

void NSTextView_inst_setFont(void *id, void* value) {
	[(NSTextView*)id
		setFont: value];
}

BOOL NSView_inst_acceptsFirstMouse(void *id, void* event) {
	return [(NSView*)id
		acceptsFirstMouse: event];
}

void NSView_inst_addConstraints(void *id, void* constraints) {
	[(NSView*)id
		addConstraints: constraints];
}

void NSView_inst_addSubview(void *id, void* view) {
	[(NSView*)id
		addSubview: view];
}

void NSView_inst_addSubview_positioned_relativeTo(void *id, void* view, unsigned long place, void* otherView) {
	[(NSView*)id
		addSubview: view
		positioned: place
		relativeTo: otherView];
}

NSRect NSView_inst_adjustScroll(void *id, NSRect newVisible) {
	return [(NSView*)id
		adjustScroll: newVisible];
}

NSRect NSView_inst_alignmentRectForFrame(void *id, NSRect frame) {
	return [(NSView*)id
		alignmentRectForFrame: frame];
}

void* NSView_inst_ancestorSharedWithView(void *id, void* view) {
	return [(NSView*)id
		ancestorSharedWithView: view];
}

BOOL NSView_inst_autoscroll(void *id, void* event) {
	return [(NSView*)id
		autoscroll: event];
}

void NSView_inst_beginDocument(void *id) {
	[(NSView*)id
		beginDocument];
}

void NSView_inst_beginPageInRect_atPlacement(void *id, NSRect rect, NSPoint location) {
	[(NSView*)id
		beginPageInRect: rect
		atPlacement: location];
}

NSRect NSView_inst_centerScanRect(void *id, NSRect rect) {
	return [(NSView*)id
		centerScanRect: rect];
}

NSPoint NSView_inst_convertPoint_fromView(void *id, NSPoint point, void* view) {
	return [(NSView*)id
		convertPoint: point
		fromView: view];
}

NSPoint NSView_inst_convertPoint_toView(void *id, NSPoint point, void* view) {
	return [(NSView*)id
		convertPoint: point
		toView: view];
}

NSPoint NSView_inst_convertPointFromBacking(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointFromBacking: point];
}

NSPoint NSView_inst_convertPointFromLayer(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointFromLayer: point];
}

NSPoint NSView_inst_convertPointToBacking(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointToBacking: point];
}

NSPoint NSView_inst_convertPointToLayer(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointToLayer: point];
}

NSRect NSView_inst_convertRect_fromView(void *id, NSRect rect, void* view) {
	return [(NSView*)id
		convertRect: rect
		fromView: view];
}

NSRect NSView_inst_convertRect_toView(void *id, NSRect rect, void* view) {
	return [(NSView*)id
		convertRect: rect
		toView: view];
}

NSRect NSView_inst_convertRectFromBacking(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectFromBacking: rect];
}

NSRect NSView_inst_convertRectFromLayer(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectFromLayer: rect];
}

NSRect NSView_inst_convertRectToBacking(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectToBacking: rect];
}

NSRect NSView_inst_convertRectToLayer(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectToLayer: rect];
}

NSSize NSView_inst_convertSize_fromView(void *id, NSSize size, void* view) {
	return [(NSView*)id
		convertSize: size
		fromView: view];
}

NSSize NSView_inst_convertSize_toView(void *id, NSSize size, void* view) {
	return [(NSView*)id
		convertSize: size
		toView: view];
}

NSSize NSView_inst_convertSizeFromBacking(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeFromBacking: size];
}

NSSize NSView_inst_convertSizeFromLayer(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeFromLayer: size];
}

NSSize NSView_inst_convertSizeToBacking(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeToBacking: size];
}

NSSize NSView_inst_convertSizeToLayer(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeToLayer: size];
}

void* NSView_inst_dataWithEPSInsideRect(void *id, NSRect rect) {
	return [(NSView*)id
		dataWithEPSInsideRect: rect];
}

void* NSView_inst_dataWithPDFInsideRect(void *id, NSRect rect) {
	return [(NSView*)id
		dataWithPDFInsideRect: rect];
}

void NSView_inst_didAddSubview(void *id, void* subview) {
	[(NSView*)id
		didAddSubview: subview];
}

void NSView_inst_didCloseMenu_withEvent(void *id, void* menu, void* event) {
	[(NSView*)id
		didCloseMenu: menu
		withEvent: event];
}

void NSView_inst_discardCursorRects(void *id) {
	[(NSView*)id
		discardCursorRects];
}

void NSView_inst_display(void *id) {
	[(NSView*)id
		display];
}

void NSView_inst_displayIfNeeded(void *id) {
	[(NSView*)id
		displayIfNeeded];
}

void NSView_inst_displayIfNeededIgnoringOpacity(void *id) {
	[(NSView*)id
		displayIfNeededIgnoringOpacity];
}

void NSView_inst_displayIfNeededInRect(void *id, NSRect rect) {
	[(NSView*)id
		displayIfNeededInRect: rect];
}

void NSView_inst_displayIfNeededInRectIgnoringOpacity(void *id, NSRect rect) {
	[(NSView*)id
		displayIfNeededInRectIgnoringOpacity: rect];
}

void NSView_inst_displayRect(void *id, NSRect rect) {
	[(NSView*)id
		displayRect: rect];
}

void NSView_inst_displayRectIgnoringOpacity(void *id, NSRect rect) {
	[(NSView*)id
		displayRectIgnoringOpacity: rect];
}

void NSView_inst_drawFocusRingMask(void *id) {
	[(NSView*)id
		drawFocusRingMask];
}

void NSView_inst_drawPageBorderWithSize(void *id, NSSize borderSize) {
	[(NSView*)id
		drawPageBorderWithSize: borderSize];
}

void NSView_inst_drawRect(void *id, NSRect dirtyRect) {
	[(NSView*)id
		drawRect: dirtyRect];
}

void NSView_inst_endDocument(void *id) {
	[(NSView*)id
		endDocument];
}

void NSView_inst_endPage(void *id) {
	[(NSView*)id
		endPage];
}

BOOL NSView_inst_enterFullScreenMode_withOptions(void *id, void* screen, void* options) {
	return [(NSView*)id
		enterFullScreenMode: screen
		withOptions: options];
}

void NSView_inst_exerciseAmbiguityInLayout(void *id) {
	[(NSView*)id
		exerciseAmbiguityInLayout];
}

void NSView_inst_exitFullScreenModeWithOptions(void *id, void* options) {
	[(NSView*)id
		exitFullScreenModeWithOptions: options];
}

NSRect NSView_inst_frameForAlignmentRect(void *id, NSRect alignmentRect) {
	return [(NSView*)id
		frameForAlignmentRect: alignmentRect];
}

void* NSView_inst_hitTest(void *id, NSPoint point) {
	return [(NSView*)id
		hitTest: point];
}

void* NSView_inst_initWithFrame(void *id, NSRect frameRect) {
	return [(NSView*)id
		initWithFrame: frameRect];
}

void NSView_inst_invalidateIntrinsicContentSize(void *id) {
	[(NSView*)id
		invalidateIntrinsicContentSize];
}

BOOL NSView_inst_isDescendantOf(void *id, void* view) {
	return [(NSView*)id
		isDescendantOf: view];
}

void NSView_inst_layout(void *id) {
	[(NSView*)id
		layout];
}

void NSView_inst_layoutSubtreeIfNeeded(void *id) {
	[(NSView*)id
		layoutSubtreeIfNeeded];
}

NSPoint NSView_inst_locationOfPrintRect(void *id, NSRect rect) {
	return [(NSView*)id
		locationOfPrintRect: rect];
}

void* NSView_inst_makeBackingLayer(void *id) {
	return [(NSView*)id
		makeBackingLayer];
}

void* NSView_inst_menuForEvent(void *id, void* event) {
	return [(NSView*)id
		menuForEvent: event];
}

BOOL NSView_inst_mouse_inRect(void *id, NSPoint point, NSRect rect) {
	return [(NSView*)id
		mouse: point
		inRect: rect];
}

BOOL NSView_inst_needsToDrawRect(void *id, NSRect rect) {
	return [(NSView*)id
		needsToDrawRect: rect];
}

void NSView_inst_noteFocusRingMaskChanged(void *id) {
	[(NSView*)id
		noteFocusRingMaskChanged];
}

BOOL NSView_inst_performKeyEquivalent(void *id, void* event) {
	return [(NSView*)id
		performKeyEquivalent: event];
}

void NSView_inst_prepareContentInRect(void *id, NSRect rect) {
	[(NSView*)id
		prepareContentInRect: rect];
}

void NSView_inst_prepareForReuse(void *id) {
	[(NSView*)id
		prepareForReuse];
}

void NSView_inst_print(void *id, void* sender) {
	[(NSView*)id
		print: sender];
}

NSRect NSView_inst_rectForPage(void *id, long page) {
	return [(NSView*)id
		rectForPage: page];
}

NSRect NSView_inst_rectForSmartMagnificationAtPoint_inRect(void *id, NSPoint location, NSRect visibleRect) {
	return [(NSView*)id
		rectForSmartMagnificationAtPoint: location
		inRect: visibleRect];
}

void NSView_inst_registerForDraggedTypes(void *id, void* newTypes) {
	[(NSView*)id
		registerForDraggedTypes: newTypes];
}

void NSView_inst_removeAllToolTips(void *id) {
	[(NSView*)id
		removeAllToolTips];
}

void NSView_inst_removeConstraints(void *id, void* constraints) {
	[(NSView*)id
		removeConstraints: constraints];
}

void NSView_inst_removeFromSuperview(void *id) {
	[(NSView*)id
		removeFromSuperview];
}

void NSView_inst_removeFromSuperviewWithoutNeedingDisplay(void *id) {
	[(NSView*)id
		removeFromSuperviewWithoutNeedingDisplay];
}

void NSView_inst_replaceSubview_with(void *id, void* oldView, void* newView) {
	[(NSView*)id
		replaceSubview: oldView
		with: newView];
}

void NSView_inst_resetCursorRects(void *id) {
	[(NSView*)id
		resetCursorRects];
}

void NSView_inst_resizeSubviewsWithOldSize(void *id, NSSize oldSize) {
	[(NSView*)id
		resizeSubviewsWithOldSize: oldSize];
}

void NSView_inst_resizeWithOldSuperviewSize(void *id, NSSize oldSize) {
	[(NSView*)id
		resizeWithOldSuperviewSize: oldSize];
}

void NSView_inst_rotateByAngle(void *id, double angle) {
	[(NSView*)id
		rotateByAngle: angle];
}

void NSView_inst_scaleUnitSquareToSize(void *id, NSSize newUnitSize) {
	[(NSView*)id
		scaleUnitSquareToSize: newUnitSize];
}

void NSView_inst_scrollPoint(void *id, NSPoint point) {
	[(NSView*)id
		scrollPoint: point];
}

BOOL NSView_inst_scrollRectToVisible(void *id, NSRect rect) {
	return [(NSView*)id
		scrollRectToVisible: rect];
}

void NSView_inst_setBoundsOrigin(void *id, NSPoint newOrigin) {
	[(NSView*)id
		setBoundsOrigin: newOrigin];
}

void NSView_inst_setBoundsSize(void *id, NSSize newSize) {
	[(NSView*)id
		setBoundsSize: newSize];
}

void NSView_inst_setFrameOrigin(void *id, NSPoint newOrigin) {
	[(NSView*)id
		setFrameOrigin: newOrigin];
}

void NSView_inst_setFrameSize(void *id, NSSize newSize) {
	[(NSView*)id
		setFrameSize: newSize];
}

void NSView_inst_setKeyboardFocusRingNeedsDisplayInRect(void *id, NSRect rect) {
	[(NSView*)id
		setKeyboardFocusRingNeedsDisplayInRect: rect];
}

void NSView_inst_setNeedsDisplayInRect(void *id, NSRect invalidRect) {
	[(NSView*)id
		setNeedsDisplayInRect: invalidRect];
}

BOOL NSView_inst_shouldDelayWindowOrderingForEvent(void *id, void* event) {
	return [(NSView*)id
		shouldDelayWindowOrderingForEvent: event];
}

void NSView_inst_showDefinitionForAttributedString_atPoint(void *id, void* attrString, NSPoint textBaselineOrigin) {
	[(NSView*)id
		showDefinitionForAttributedString: attrString
		atPoint: textBaselineOrigin];
}

void NSView_inst_translateOriginToPoint(void *id, NSPoint translation) {
	[(NSView*)id
		translateOriginToPoint: translation];
}

void NSView_inst_translateRectsNeedingDisplayInRect_by(void *id, NSRect clipRect, NSSize delta) {
	[(NSView*)id
		translateRectsNeedingDisplayInRect: clipRect
		by: delta];
}

void NSView_inst_unregisterDraggedTypes(void *id) {
	[(NSView*)id
		unregisterDraggedTypes];
}

void NSView_inst_updateConstraints(void *id) {
	[(NSView*)id
		updateConstraints];
}

void NSView_inst_updateConstraintsForSubtreeIfNeeded(void *id) {
	[(NSView*)id
		updateConstraintsForSubtreeIfNeeded];
}

void NSView_inst_updateLayer(void *id) {
	[(NSView*)id
		updateLayer];
}

void NSView_inst_updateTrackingAreas(void *id) {
	[(NSView*)id
		updateTrackingAreas];
}

void NSView_inst_viewDidChangeBackingProperties(void *id) {
	[(NSView*)id
		viewDidChangeBackingProperties];
}

void NSView_inst_viewDidChangeEffectiveAppearance(void *id) {
	[(NSView*)id
		viewDidChangeEffectiveAppearance];
}

void NSView_inst_viewDidEndLiveResize(void *id) {
	[(NSView*)id
		viewDidEndLiveResize];
}

void NSView_inst_viewDidHide(void *id) {
	[(NSView*)id
		viewDidHide];
}

void NSView_inst_viewDidMoveToSuperview(void *id) {
	[(NSView*)id
		viewDidMoveToSuperview];
}

void NSView_inst_viewDidMoveToWindow(void *id) {
	[(NSView*)id
		viewDidMoveToWindow];
}

void NSView_inst_viewDidUnhide(void *id) {
	[(NSView*)id
		viewDidUnhide];
}

void NSView_inst_viewWillDraw(void *id) {
	[(NSView*)id
		viewWillDraw];
}

void NSView_inst_viewWillMoveToSuperview(void *id, void* newSuperview) {
	[(NSView*)id
		viewWillMoveToSuperview: newSuperview];
}

void NSView_inst_viewWillMoveToWindow(void *id, void* newWindow) {
	[(NSView*)id
		viewWillMoveToWindow: newWindow];
}

void NSView_inst_viewWillStartLiveResize(void *id) {
	[(NSView*)id
		viewWillStartLiveResize];
}

void* NSView_inst_viewWithTag(void *id, long tag) {
	return [(NSView*)id
		viewWithTag: tag];
}

void NSView_inst_willOpenMenu_withEvent(void *id, void* menu, void* event) {
	[(NSView*)id
		willOpenMenu: menu
		withEvent: event];
}

void NSView_inst_willRemoveSubview(void *id, void* subview) {
	[(NSView*)id
		willRemoveSubview: subview];
}

void NSView_inst_writeEPSInsideRect_toPasteboard(void *id, NSRect rect, void* pasteboard) {
	[(NSView*)id
		writeEPSInsideRect: rect
		toPasteboard: pasteboard];
}

void NSView_inst_writePDFInsideRect_toPasteboard(void *id, NSRect rect, void* pasteboard) {
	[(NSView*)id
		writePDFInsideRect: rect
		toPasteboard: pasteboard];
}

void* NSView_inst_init(void *id) {
	return [(NSView*)id
		init];
}

void* NSView_inst_superview(void *id) {
	return [(NSView*)id
		superview];
}

void* NSView_inst_subviews(void *id) {
	return [(NSView*)id
		subviews];
}

void NSView_inst_setSubviews(void *id, void* value) {
	[(NSView*)id
		setSubviews: value];
}

void* NSView_inst_window(void *id) {
	return [(NSView*)id
		window];
}

void* NSView_inst_opaqueAncestor(void *id) {
	return [(NSView*)id
		opaqueAncestor];
}

void* NSView_inst_enclosingMenuItem(void *id) {
	return [(NSView*)id
		enclosingMenuItem];
}

NSRect NSView_inst_frame(void *id) {
	return [(NSView*)id
		frame];
}

void NSView_inst_setFrame(void *id, NSRect value) {
	[(NSView*)id
		setFrame: value];
}

double NSView_inst_frameRotation(void *id) {
	return [(NSView*)id
		frameRotation];
}

void NSView_inst_setFrameRotation(void *id, double value) {
	[(NSView*)id
		setFrameRotation: value];
}

NSRect NSView_inst_bounds(void *id) {
	return [(NSView*)id
		bounds];
}

void NSView_inst_setBounds(void *id, NSRect value) {
	[(NSView*)id
		setBounds: value];
}

double NSView_inst_boundsRotation(void *id) {
	return [(NSView*)id
		boundsRotation];
}

void NSView_inst_setBoundsRotation(void *id, double value) {
	[(NSView*)id
		setBoundsRotation: value];
}

BOOL NSView_inst_wantsLayer(void *id) {
	return [(NSView*)id
		wantsLayer];
}

void NSView_inst_setWantsLayer(void *id, BOOL value) {
	[(NSView*)id
		setWantsLayer: value];
}

BOOL NSView_inst_wantsUpdateLayer(void *id) {
	return [(NSView*)id
		wantsUpdateLayer];
}

void* NSView_inst_layer(void *id) {
	return [(NSView*)id
		layer];
}

void NSView_inst_setLayer(void *id, void* value) {
	[(NSView*)id
		setLayer: value];
}

BOOL NSView_inst_canDrawSubviewsIntoLayer(void *id) {
	return [(NSView*)id
		canDrawSubviewsIntoLayer];
}

void NSView_inst_setCanDrawSubviewsIntoLayer(void *id, BOOL value) {
	[(NSView*)id
		setCanDrawSubviewsIntoLayer: value];
}

BOOL NSView_inst_layerUsesCoreImageFilters(void *id) {
	return [(NSView*)id
		layerUsesCoreImageFilters];
}

void NSView_inst_setLayerUsesCoreImageFilters(void *id, BOOL value) {
	[(NSView*)id
		setLayerUsesCoreImageFilters: value];
}

double NSView_inst_alphaValue(void *id) {
	return [(NSView*)id
		alphaValue];
}

void NSView_inst_setAlphaValue(void *id, double value) {
	[(NSView*)id
		setAlphaValue: value];
}

double NSView_inst_frameCenterRotation(void *id) {
	return [(NSView*)id
		frameCenterRotation];
}

void NSView_inst_setFrameCenterRotation(void *id, double value) {
	[(NSView*)id
		setFrameCenterRotation: value];
}

void* NSView_inst_backgroundFilters(void *id) {
	return [(NSView*)id
		backgroundFilters];
}

void NSView_inst_setBackgroundFilters(void *id, void* value) {
	[(NSView*)id
		setBackgroundFilters: value];
}

void* NSView_inst_contentFilters(void *id) {
	return [(NSView*)id
		contentFilters];
}

void NSView_inst_setContentFilters(void *id, void* value) {
	[(NSView*)id
		setContentFilters: value];
}

BOOL NSView_inst_canDrawConcurrently(void *id) {
	return [(NSView*)id
		canDrawConcurrently];
}

void NSView_inst_setCanDrawConcurrently(void *id, BOOL value) {
	[(NSView*)id
		setCanDrawConcurrently: value];
}

NSRect NSView_inst_visibleRect(void *id) {
	return [(NSView*)id
		visibleRect];
}

BOOL NSView_inst_wantsDefaultClipping(void *id) {
	return [(NSView*)id
		wantsDefaultClipping];
}

void* NSView_inst_printJobTitle(void *id) {
	return [(NSView*)id
		printJobTitle];
}

void* NSView_inst_pageHeader(void *id) {
	return [(NSView*)id
		pageHeader];
}

void* NSView_inst_pageFooter(void *id) {
	return [(NSView*)id
		pageFooter];
}

double NSView_inst_heightAdjustLimit(void *id) {
	return [(NSView*)id
		heightAdjustLimit];
}

double NSView_inst_widthAdjustLimit(void *id) {
	return [(NSView*)id
		widthAdjustLimit];
}

BOOL NSView_inst_needsDisplay(void *id) {
	return [(NSView*)id
		needsDisplay];
}

void NSView_inst_setNeedsDisplay(void *id, BOOL value) {
	[(NSView*)id
		setNeedsDisplay: value];
}

BOOL NSView_inst_isOpaque(void *id) {
	return [(NSView*)id
		isOpaque];
}

BOOL NSView_inst_isFlipped(void *id) {
	return [(NSView*)id
		isFlipped];
}

BOOL NSView_inst_isRotatedFromBase(void *id) {
	return [(NSView*)id
		isRotatedFromBase];
}

BOOL NSView_inst_isRotatedOrScaledFromBase(void *id) {
	return [(NSView*)id
		isRotatedOrScaledFromBase];
}

BOOL NSView_inst_autoresizesSubviews(void *id) {
	return [(NSView*)id
		autoresizesSubviews];
}

void NSView_inst_setAutoresizesSubviews(void *id, BOOL value) {
	[(NSView*)id
		setAutoresizesSubviews: value];
}

void* NSView_inst_constraints(void *id) {
	return [(NSView*)id
		constraints];
}

void* NSView_inst_layoutGuides(void *id) {
	return [(NSView*)id
		layoutGuides];
}

NSSize NSView_inst_fittingSize(void *id) {
	return [(NSView*)id
		fittingSize];
}

NSSize NSView_inst_intrinsicContentSize(void *id) {
	return [(NSView*)id
		intrinsicContentSize];
}

double NSView_inst_baselineOffsetFromBottom(void *id) {
	return [(NSView*)id
		baselineOffsetFromBottom];
}

double NSView_inst_firstBaselineOffsetFromTop(void *id) {
	return [(NSView*)id
		firstBaselineOffsetFromTop];
}

double NSView_inst_lastBaselineOffsetFromBottom(void *id) {
	return [(NSView*)id
		lastBaselineOffsetFromBottom];
}

BOOL NSView_inst_needsLayout(void *id) {
	return [(NSView*)id
		needsLayout];
}

void NSView_inst_setNeedsLayout(void *id, BOOL value) {
	[(NSView*)id
		setNeedsLayout: value];
}

BOOL NSView_inst_needsUpdateConstraints(void *id) {
	return [(NSView*)id
		needsUpdateConstraints];
}

void NSView_inst_setNeedsUpdateConstraints(void *id, BOOL value) {
	[(NSView*)id
		setNeedsUpdateConstraints: value];
}

BOOL NSView_inst_translatesAutoresizingMaskIntoConstraints(void *id) {
	return [(NSView*)id
		translatesAutoresizingMaskIntoConstraints];
}

void NSView_inst_setTranslatesAutoresizingMaskIntoConstraints(void *id, BOOL value) {
	[(NSView*)id
		setTranslatesAutoresizingMaskIntoConstraints: value];
}

BOOL NSView_inst_hasAmbiguousLayout(void *id) {
	return [(NSView*)id
		hasAmbiguousLayout];
}

NSRect NSView_inst_focusRingMaskBounds(void *id) {
	return [(NSView*)id
		focusRingMaskBounds];
}

BOOL NSView_inst_allowsVibrancy(void *id) {
	return [(NSView*)id
		allowsVibrancy];
}

BOOL NSView_inst_isInFullScreenMode(void *id) {
	return [(NSView*)id
		isInFullScreenMode];
}

BOOL NSView_inst_isHidden(void *id) {
	return [(NSView*)id
		isHidden];
}

void NSView_inst_setHidden(void *id, BOOL value) {
	[(NSView*)id
		setHidden: value];
}

BOOL NSView_inst_isHiddenOrHasHiddenAncestor(void *id) {
	return [(NSView*)id
		isHiddenOrHasHiddenAncestor];
}

BOOL NSView_inst_inLiveResize(void *id) {
	return [(NSView*)id
		inLiveResize];
}

BOOL NSView_inst_preservesContentDuringLiveResize(void *id) {
	return [(NSView*)id
		preservesContentDuringLiveResize];
}

NSRect NSView_inst_rectPreservedDuringLiveResize(void *id) {
	return [(NSView*)id
		rectPreservedDuringLiveResize];
}

void* NSView_inst_gestureRecognizers(void *id) {
	return [(NSView*)id
		gestureRecognizers];
}

void NSView_inst_setGestureRecognizers(void *id, void* value) {
	[(NSView*)id
		setGestureRecognizers: value];
}

BOOL NSView_inst_mouseDownCanMoveWindow(void *id) {
	return [(NSView*)id
		mouseDownCanMoveWindow];
}

BOOL NSView_inst_wantsRestingTouches(void *id) {
	return [(NSView*)id
		wantsRestingTouches];
}

void NSView_inst_setWantsRestingTouches(void *id, BOOL value) {
	[(NSView*)id
		setWantsRestingTouches: value];
}

BOOL NSView_inst_canBecomeKeyView(void *id) {
	return [(NSView*)id
		canBecomeKeyView];
}

BOOL NSView_inst_needsPanelToBecomeKey(void *id) {
	return [(NSView*)id
		needsPanelToBecomeKey];
}

void* NSView_inst_nextKeyView(void *id) {
	return [(NSView*)id
		nextKeyView];
}

void NSView_inst_setNextKeyView(void *id, void* value) {
	[(NSView*)id
		setNextKeyView: value];
}

void* NSView_inst_nextValidKeyView(void *id) {
	return [(NSView*)id
		nextValidKeyView];
}

void* NSView_inst_previousKeyView(void *id) {
	return [(NSView*)id
		previousKeyView];
}

void* NSView_inst_previousValidKeyView(void *id) {
	return [(NSView*)id
		previousValidKeyView];
}

NSRect NSView_inst_preparedContentRect(void *id) {
	return [(NSView*)id
		preparedContentRect];
}

void NSView_inst_setPreparedContentRect(void *id, NSRect value) {
	[(NSView*)id
		setPreparedContentRect: value];
}

void* NSView_inst_registeredDraggedTypes(void *id) {
	return [(NSView*)id
		registeredDraggedTypes];
}

BOOL NSView_inst_postsFrameChangedNotifications(void *id) {
	return [(NSView*)id
		postsFrameChangedNotifications];
}

void NSView_inst_setPostsFrameChangedNotifications(void *id, BOOL value) {
	[(NSView*)id
		setPostsFrameChangedNotifications: value];
}

BOOL NSView_inst_postsBoundsChangedNotifications(void *id) {
	return [(NSView*)id
		postsBoundsChangedNotifications];
}

void NSView_inst_setPostsBoundsChangedNotifications(void *id, BOOL value) {
	[(NSView*)id
		setPostsBoundsChangedNotifications: value];
}

long NSView_inst_tag(void *id) {
	return [(NSView*)id
		tag];
}

void* NSView_inst_toolTip(void *id) {
	return [(NSView*)id
		toolTip];
}

void NSView_inst_setToolTip(void *id, void* value) {
	[(NSView*)id
		setToolTip: value];
}

void* NSView_inst_trackingAreas(void *id) {
	return [(NSView*)id
		trackingAreas];
}

BOOL NSView_inst_isDrawingFindIndicator(void *id) {
	return [(NSView*)id
		isDrawingFindIndicator];
}

BOOL NSView_inst_isHorizontalContentSizeConstraintActive(void *id) {
	return [(NSView*)id
		isHorizontalContentSizeConstraintActive];
}

void NSView_inst_setHorizontalContentSizeConstraintActive(void *id, BOOL value) {
	[(NSView*)id
		setHorizontalContentSizeConstraintActive: value];
}

BOOL NSView_inst_isVerticalContentSizeConstraintActive(void *id) {
	return [(NSView*)id
		isVerticalContentSizeConstraintActive];
}

void NSView_inst_setVerticalContentSizeConstraintActive(void *id, BOOL value) {
	[(NSView*)id
		setVerticalContentSizeConstraintActive: value];
}

void* NSView_inst_backgroundColor(void *id) {
	return [(NSView*)id
		backgroundColor];
}

void NSView_inst_setBackgroundColor(void *id, void* value) {
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

// NSBundle_alloc
func NSBundle_alloc() NSBundle {
	ret := C.NSBundle_type_alloc()

	return NSBundle_fromPointer(ret)

}

// NSBundle_bundleWithURL Returns an NSBundle object that corresponds to the specified file URL.
// https://developer.apple.com/documentation/foundation/nsbundle/1494992-bundlewithurl?language=objc
func NSBundle_bundleWithURL(url core.NSURLRef) NSBundle {
	ret := C.NSBundle_type_bundleWithURL(
		objc.RefPointer(url),
	)

	return NSBundle_fromPointer(ret)

}

// NSBundle_bundleWithPath Returns an NSBundle object that corresponds to the specified directory.
// https://developer.apple.com/documentation/foundation/nsbundle/1495012-bundlewithpath?language=objc
func NSBundle_bundleWithPath(path core.NSStringRef) NSBundle {
	ret := C.NSBundle_type_bundleWithPath(
		objc.RefPointer(path),
	)

	return NSBundle_fromPointer(ret)

}

// NSBundle_bundleWithIdentifier Returns the NSBundle instance that has the specified bundle identifier.
// https://developer.apple.com/documentation/foundation/nsbundle/1411929-bundlewithidentifier?language=objc
func NSBundle_bundleWithIdentifier(identifier core.NSStringRef) NSBundle {
	ret := C.NSBundle_type_bundleWithIdentifier(
		objc.RefPointer(identifier),
	)

	return NSBundle_fromPointer(ret)

}

// NSBundle_URLForResource_withExtension_subdirectory_inBundleWithURL Creates and returns a file URL for the resource with the specified name and extension in the specified bundle.
// https://developer.apple.com/documentation/foundation/nsbundle/1416361-urlforresource?language=objc
func NSBundle_URLForResource_withExtension_subdirectory_inBundleWithURL(name core.NSStringRef, ext core.NSStringRef, subpath core.NSStringRef, bundleURL core.NSURLRef) core.NSURL {
	ret := C.NSBundle_type_URLForResource_withExtension_subdirectory_inBundleWithURL(
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(bundleURL),
	)

	return core.NSURL_fromPointer(ret)

}

// NSBundle_URLsForResourcesWithExtension_subdirectory_inBundleWithURL Returns an array containing the file URLs for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, within the specified bundle.
// https://developer.apple.com/documentation/foundation/nsbundle/1409807-urlsforresourceswithextension?language=objc
func NSBundle_URLsForResourcesWithExtension_subdirectory_inBundleWithURL(ext core.NSStringRef, subpath core.NSStringRef, bundleURL core.NSURLRef) core.NSArray {
	ret := C.NSBundle_type_URLsForResourcesWithExtension_subdirectory_inBundleWithURL(
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(bundleURL),
	)

	return core.NSArray_fromPointer(ret)

}

// NSBundle_pathForResource_ofType_inDirectory Returns the full pathname for the resource file identified by the specified name and extension and residing in a given bundle directory.
// https://developer.apple.com/documentation/foundation/nsbundle/1409523-pathforresource?language=objc
func NSBundle_pathForResource_ofType_inDirectory(name core.NSStringRef, ext core.NSStringRef, bundlePath core.NSStringRef) core.NSString {
	ret := C.NSBundle_type_pathForResource_ofType_inDirectory(
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(bundlePath),
	)

	return core.NSString_fromPointer(ret)

}

// NSBundle_pathsForResourcesOfType_inDirectory Returns an array containing the pathnames for all bundle resources having the specified extension and residing in the bundle directory at the specified path.
// https://developer.apple.com/documentation/foundation/nsbundle/1415876-pathsforresourcesoftype?language=objc
func NSBundle_pathsForResourcesOfType_inDirectory(ext core.NSStringRef, bundlePath core.NSStringRef) core.NSArray {
	ret := C.NSBundle_type_pathsForResourcesOfType_inDirectory(
		objc.RefPointer(ext),
		objc.RefPointer(bundlePath),
	)

	return core.NSArray_fromPointer(ret)

}

// NSBundle_preferredLocalizationsFromArray Returns one or more localizations from the specified list that a bundle object would use to locate resources for the current user.
// https://developer.apple.com/documentation/foundation/nsbundle/1417249-preferredlocalizationsfromarray?language=objc
func NSBundle_preferredLocalizationsFromArray(localizationsArray core.NSArrayRef) core.NSArray {
	ret := C.NSBundle_type_preferredLocalizationsFromArray(
		objc.RefPointer(localizationsArray),
	)

	return core.NSArray_fromPointer(ret)

}

// NSBundle_preferredLocalizationsFromArray_forPreferences Returns locale identifiers for which a bundle would provide localized content, given a specified list of candidates for a user's language preferences.
// https://developer.apple.com/documentation/foundation/nsbundle/1409418-preferredlocalizationsfromarray?language=objc
func NSBundle_preferredLocalizationsFromArray_forPreferences(localizationsArray core.NSArrayRef, preferencesArray core.NSArrayRef) core.NSArray {
	ret := C.NSBundle_type_preferredLocalizationsFromArray_forPreferences(
		objc.RefPointer(localizationsArray),
		objc.RefPointer(preferencesArray),
	)

	return core.NSArray_fromPointer(ret)

}

// NSBundle_mainBundle Returns the bundle object that contains the current executable.
// https://developer.apple.com/documentation/foundation/nsbundle/1410786-mainbundle?language=objc
func NSBundle_mainBundle() NSBundle {
	ret := C.NSBundle_type_mainBundle()

	return NSBundle_fromPointer(ret)

}

// NSBundle_allFrameworks Returns an array of all of the application’s bundles that represent frameworks.
// https://developer.apple.com/documentation/foundation/nsbundle/1408056-allframeworks?language=objc
func NSBundle_allFrameworks() core.NSArray {
	ret := C.NSBundle_type_allFrameworks()

	return core.NSArray_fromPointer(ret)

}

// NSBundle_allBundles Returns an array of all the application’s non-framework bundles.
// https://developer.apple.com/documentation/foundation/nsbundle/1413705-allbundles?language=objc
func NSBundle_allBundles() core.NSArray {
	ret := C.NSBundle_type_allBundles()

	return core.NSArray_fromPointer(ret)

}

// NSSound_alloc
func NSSound_alloc() NSSound {
	ret := C.NSSound_type_alloc()

	return NSSound_fromPointer(ret)

}

// NSSound_canInitWithPasteboard Indicates whether the receiver can create an instance of itself from the data in a pasteboard.
// https://developer.apple.com/documentation/appkit/nssound/1477276-caninitwithpasteboard?language=objc
func NSSound_canInitWithPasteboard(pasteboard NSPasteboardRef) bool {
	ret := C.NSSound_type_canInitWithPasteboard(
		objc.RefPointer(pasteboard),
	)

	return convertObjCBoolToGo(ret)

}

// NSSound_soundUnfilteredTypes Provides the file types the NSSound class understands.
// https://developer.apple.com/documentation/appkit/nssound/1477290-soundunfilteredtypes?language=objc
func NSSound_soundUnfilteredTypes() core.NSArray {
	ret := C.NSSound_type_soundUnfilteredTypes()

	return core.NSArray_fromPointer(ret)

}

// NSApplication_alloc
func NSApplication_alloc() NSApplication {
	ret := C.NSApplication_type_alloc()

	return NSApplication_fromPointer(ret)

}

// NSApplication_detachDrawingThread_toTarget_withObject Creates and executes a new thread based on the specified target and selector.
// https://developer.apple.com/documentation/appkit/nsapplication/1428374-detachdrawingthread?language=objc
func NSApplication_detachDrawingThread_toTarget_withObject(selector objc.Selector, target objc.Ref, argument objc.Ref) {
	C.NSApplication_type_detachDrawingThread_toTarget_withObject(
		selector.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(argument),
	)

	return

}

// NSApplication_sharedApplication Returns the application instance, creating it if it doesn’t exist yet.
// https://developer.apple.com/documentation/appkit/nsapplication/1428360-sharedapplication?language=objc
func NSApplication_sharedApplication() NSApplication {
	ret := C.NSApplication_type_sharedApplication()

	return NSApplication_fromPointer(ret)

}

// NSControl_alloc
func NSControl_alloc() NSControl {
	ret := C.NSControl_type_alloc()

	return NSControl_fromPointer(ret)

}

// NSButton_alloc
func NSButton_alloc() NSButton {
	ret := C.NSButton_type_alloc()

	return NSButton_fromPointer(ret)

}

// NSButton_checkboxWithTitle_target_action Creates a standard checkbox with the title you specify.
// https://developer.apple.com/documentation/appkit/nsbutton/1644525-checkboxwithtitle?language=objc
func NSButton_checkboxWithTitle_target_action(title core.NSStringRef, target objc.Ref, action objc.Selector) NSButton {
	ret := C.NSButton_type_checkboxWithTitle_target_action(
		objc.RefPointer(title),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)

	return NSButton_fromPointer(ret)

}

// NSButton_buttonWithImage_target_action Creates a standard push button with the image you specify.
// https://developer.apple.com/documentation/appkit/nsbutton/1644659-buttonwithimage?language=objc
func NSButton_buttonWithImage_target_action(image NSImageRef, target objc.Ref, action objc.Selector) NSButton {
	ret := C.NSButton_type_buttonWithImage_target_action(
		objc.RefPointer(image),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)

	return NSButton_fromPointer(ret)

}

// NSButton_radioButtonWithTitle_target_action Creates a standard radio button with the title you specify.
// https://developer.apple.com/documentation/appkit/nsbutton/1644340-radiobuttonwithtitle?language=objc
func NSButton_radioButtonWithTitle_target_action(title core.NSStringRef, target objc.Ref, action objc.Selector) NSButton {
	ret := C.NSButton_type_radioButtonWithTitle_target_action(
		objc.RefPointer(title),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)

	return NSButton_fromPointer(ret)

}

// NSButton_buttonWithTitle_image_target_action Creates a standard push button with a title and image.
// https://developer.apple.com/documentation/appkit/nsbutton/1644719-buttonwithtitle?language=objc
func NSButton_buttonWithTitle_image_target_action(title core.NSStringRef, image NSImageRef, target objc.Ref, action objc.Selector) NSButton {
	ret := C.NSButton_type_buttonWithTitle_image_target_action(
		objc.RefPointer(title),
		objc.RefPointer(image),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)

	return NSButton_fromPointer(ret)

}

// NSButton_buttonWithTitle_target_action Creates a standard push button with the title you specify.
// https://developer.apple.com/documentation/appkit/nsbutton/1644256-buttonwithtitle?language=objc
func NSButton_buttonWithTitle_target_action(title core.NSStringRef, target objc.Ref, action objc.Selector) NSButton {
	ret := C.NSButton_type_buttonWithTitle_target_action(
		objc.RefPointer(title),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)

	return NSButton_fromPointer(ret)

}

// NSEvent_alloc
func NSEvent_alloc() NSEvent {
	ret := C.NSEvent_type_alloc()

	return NSEvent_fromPointer(ret)

}

// NSEvent_eventWithEventRef Creates an event object that is based on a Carbon type of event.
// https://developer.apple.com/documentation/appkit/nsevent/1528021-eventwitheventref?language=objc
func NSEvent_eventWithEventRef(eventRef unsafe.Pointer) NSEvent {
	ret := C.NSEvent_type_eventWithEventRef(
		eventRef,
	)

	return NSEvent_fromPointer(ret)

}

// NSEvent_stopPeriodicEvents Stops generating periodic events for the current thread and discards any periodic events remaining in the queue.
// https://developer.apple.com/documentation/appkit/nsevent/1533746-stopperiodicevents?language=objc
func NSEvent_stopPeriodicEvents() {
	C.NSEvent_type_stopPeriodicEvents()

	return

}

// NSEvent_removeMonitor Remove the specified event monitor.
// https://developer.apple.com/documentation/appkit/nsevent/1533709-removemonitor?language=objc
func NSEvent_removeMonitor(eventMonitor objc.Ref) {
	C.NSEvent_type_removeMonitor(
		objc.RefPointer(eventMonitor),
	)

	return

}

// NSEvent_pressedMouseButtons Returns the indices of the currently depressed mouse buttons.
// https://developer.apple.com/documentation/appkit/nsevent/1527943-pressedmousebuttons?language=objc
func NSEvent_pressedMouseButtons() core.NSUInteger {
	ret := C.NSEvent_type_pressedMouseButtons()

	return core.NSUInteger(ret)

}

// NSEvent_mouseLocation Reports the current mouse position in screen coordinates.
// https://developer.apple.com/documentation/appkit/nsevent/1533380-mouselocation?language=objc
func NSEvent_mouseLocation() core.NSPoint {
	ret := C.NSEvent_type_mouseLocation()

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// NSEvent_mouseCoalescingEnabled
// https://developer.apple.com/documentation/appkit/nsevent/2870068-mousecoalescingenabled?language=objc
func NSEvent_mouseCoalescingEnabled() bool {
	ret := C.NSEvent_type_mouseCoalescingEnabled()

	return convertObjCBoolToGo(ret)

}

// NSEvent_setMouseCoalescingEnabled
// https://developer.apple.com/documentation/appkit/nsevent/2870068-mousecoalescingenabled?language=objc
func NSEvent_setMouseCoalescingEnabled(value bool) {
	C.NSEvent_type_setMouseCoalescingEnabled(
		convertToObjCBool(value),
	)

	return

}

// NSEvent_swipeTrackingFromScrollEventsEnabled
// https://developer.apple.com/documentation/appkit/nsevent/2870067-swipetrackingfromscrolleventsena?language=objc
func NSEvent_swipeTrackingFromScrollEventsEnabled() bool {
	ret := C.NSEvent_type_swipeTrackingFromScrollEventsEnabled()

	return convertObjCBoolToGo(ret)

}

// NSFont_alloc
func NSFont_alloc() NSFont {
	ret := C.NSFont_type_alloc()

	return NSFont_fromPointer(ret)

}

// NSFont_fontWithName_size Creates a font object for the specified font name and font size.
// https://developer.apple.com/documentation/appkit/nsfont/1525977-fontwithname?language=objc
func NSFont_fontWithName_size(fontName core.NSStringRef, fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_fontWithName_size(
		objc.RefPointer(fontName),
		C.double(fontSize),
	)

	return NSFont_fromPointer(ret)

}

// NSFont_userFontOfSize Returns the font used by default for documents and other text under the user’s control (that is, text whose font the user can normally change), in the specified size.
// https://developer.apple.com/documentation/appkit/nsfont/1524559-userfontofsize?language=objc
func NSFont_userFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_userFontOfSize(
		C.double(fontSize),
	)

	return NSFont_fromPointer(ret)

}

// NSFont_userFixedPitchFontOfSize Returns the font used by default for documents and other text under the user’s control (that is, text whose font the user can normally change), when that font should be fixed-pitch, in the specified size.
// https://developer.apple.com/documentation/appkit/nsfont/1531381-userfixedpitchfontofsize?language=objc
func NSFont_userFixedPitchFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_userFixedPitchFontOfSize(
		C.double(fontSize),
	)

	return NSFont_fromPointer(ret)

}

// NSFont_systemFontOfSize Returns the standard system font with the specified size.
// https://developer.apple.com/documentation/appkit/nsfont/1530094-systemfontofsize?language=objc
func NSFont_systemFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_systemFontOfSize(
		C.double(fontSize),
	)

	return NSFont_fromPointer(ret)

}

// NSFont_boldSystemFontOfSize Returns the standard system font in boldface type with the specified size.
// https://developer.apple.com/documentation/appkit/nsfont/1533549-boldsystemfontofsize?language=objc
func NSFont_boldSystemFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_boldSystemFontOfSize(
		C.double(fontSize),
	)

	return NSFont_fromPointer(ret)

}

// NSFont_labelFontOfSize Returns the font used for standard interface labels in the specified size.
// https://developer.apple.com/documentation/appkit/nsfont/1528213-labelfontofsize?language=objc
func NSFont_labelFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_labelFontOfSize(
		C.double(fontSize),
	)

	return NSFont_fromPointer(ret)

}

// NSFont_messageFontOfSize Returns the font used for standard interface items, such as button labels, menu items, and so on, in the specified size.
// https://developer.apple.com/documentation/appkit/nsfont/1525777-messagefontofsize?language=objc
func NSFont_messageFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_messageFontOfSize(
		C.double(fontSize),
	)

	return NSFont_fromPointer(ret)

}

// NSFont_menuBarFontOfSize Returns the font used for menu bar items, in the specified size.
// https://developer.apple.com/documentation/appkit/nsfont/1534194-menubarfontofsize?language=objc
func NSFont_menuBarFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_menuBarFontOfSize(
		C.double(fontSize),
	)

	return NSFont_fromPointer(ret)

}

// NSFont_menuFontOfSize Returns the font used for menu items, in the specified size.
// https://developer.apple.com/documentation/appkit/nsfont/1533068-menufontofsize?language=objc
func NSFont_menuFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_menuFontOfSize(
		C.double(fontSize),
	)

	return NSFont_fromPointer(ret)

}

// NSFont_controlContentFontOfSize Returns the font used for the content of controls in the specified size.
// https://developer.apple.com/documentation/appkit/nsfont/1527070-controlcontentfontofsize?language=objc
func NSFont_controlContentFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_controlContentFontOfSize(
		C.double(fontSize),
	)

	return NSFont_fromPointer(ret)

}

// NSFont_titleBarFontOfSize Returns the font used for window title bars, in the specified size.
// https://developer.apple.com/documentation/appkit/nsfont/1530200-titlebarfontofsize?language=objc
func NSFont_titleBarFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_titleBarFontOfSize(
		C.double(fontSize),
	)

	return NSFont_fromPointer(ret)

}

// NSFont_paletteFontOfSize Returns the font used for palette window title bars, in the specified size.
// https://developer.apple.com/documentation/appkit/nsfont/1535462-palettefontofsize?language=objc
func NSFont_paletteFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_paletteFontOfSize(
		C.double(fontSize),
	)

	return NSFont_fromPointer(ret)

}

// NSFont_toolTipsFontOfSize Returns the font used for tool tips labels, in the specified size.
// https://developer.apple.com/documentation/appkit/nsfont/1527704-tooltipsfontofsize?language=objc
func NSFont_toolTipsFontOfSize(fontSize core.CGFloat) NSFont {
	ret := C.NSFont_type_toolTipsFontOfSize(
		C.double(fontSize),
	)

	return NSFont_fromPointer(ret)

}

// NSFont_setUserFont Sets the font used by default for documents and other text under the user’s control to the specified font.
// https://developer.apple.com/documentation/appkit/nsfont/1526068-setuserfont?language=objc
func NSFont_setUserFont(font NSFontRef) {
	C.NSFont_type_setUserFont(
		objc.RefPointer(font),
	)

	return

}

// NSFont_setUserFixedPitchFont Sets the font used by default for documents and other text under the user’s control, when that font should be fixed-pitch, to the specified font.
// https://developer.apple.com/documentation/appkit/nsfont/1529050-setuserfixedpitchfont?language=objc
func NSFont_setUserFixedPitchFont(font NSFontRef) {
	C.NSFont_type_setUserFixedPitchFont(
		objc.RefPointer(font),
	)

	return

}

// NSFont_systemFontSize Returns the size of the standard system font.
// https://developer.apple.com/documentation/appkit/nsfont/1531931-systemfontsize?language=objc
func NSFont_systemFontSize() core.CGFloat {
	ret := C.NSFont_type_systemFontSize()

	return core.CGFloat(ret)

}

// NSFont_smallSystemFontSize Returns the size of the standard small system font.
// https://developer.apple.com/documentation/appkit/nsfont/1535612-smallsystemfontsize?language=objc
func NSFont_smallSystemFontSize() core.CGFloat {
	ret := C.NSFont_type_smallSystemFontSize()

	return core.CGFloat(ret)

}

// NSFont_labelFontSize Returns the size of the standard label font.
// https://developer.apple.com/documentation/appkit/nsfont/1534629-labelfontsize?language=objc
func NSFont_labelFontSize() core.CGFloat {
	ret := C.NSFont_type_labelFontSize()

	return core.CGFloat(ret)

}

// NSImage_alloc
func NSImage_alloc() NSImage {
	ret := C.NSImage_type_alloc()

	return NSImage_fromPointer(ret)

}

// NSImage_imageWithSystemSymbolName_accessibilityDescription Creates a symbol image with the system symbol name and accessibility description that you specify.
// https://developer.apple.com/documentation/appkit/nsimage/3622472-imagewithsystemsymbolname?language=objc
func NSImage_imageWithSystemSymbolName_accessibilityDescription(symbolName core.NSStringRef, description core.NSStringRef) NSImage {
	ret := C.NSImage_type_imageWithSystemSymbolName_accessibilityDescription(
		objc.RefPointer(symbolName),
		objc.RefPointer(description),
	)

	return NSImage_fromPointer(ret)

}

// NSImage_canInitWithPasteboard Tests whether the image can create an instance of itself using pasteboard data.
// https://developer.apple.com/documentation/appkit/nsimage/1520039-caninitwithpasteboard?language=objc
func NSImage_canInitWithPasteboard(pasteboard NSPasteboardRef) bool {
	ret := C.NSImage_type_canInitWithPasteboard(
		objc.RefPointer(pasteboard),
	)

	return convertObjCBoolToGo(ret)

}

// NSImage_imageTypes Returns an array of UTI strings identifying the image types supported by the registered image representation objects, either directly or through a user-installed filter service.
// https://developer.apple.com/documentation/appkit/nsimage/1519988-imagetypes?language=objc
func NSImage_imageTypes() core.NSArray {
	ret := C.NSImage_type_imageTypes()

	return core.NSArray_fromPointer(ret)

}

// NSImage_imageUnfilteredTypes Returns an array of UTI strings identifying the image types supported directly by the registered image representation objects.
// https://developer.apple.com/documentation/appkit/nsimage/1519899-imageunfilteredtypes?language=objc
func NSImage_imageUnfilteredTypes() core.NSArray {
	ret := C.NSImage_type_imageUnfilteredTypes()

	return core.NSArray_fromPointer(ret)

}

// NSImageView_alloc
func NSImageView_alloc() NSImageView {
	ret := C.NSImageView_type_alloc()

	return NSImageView_fromPointer(ret)

}

// NSImageView_imageViewWithImage
// https://developer.apple.com/documentation/appkit/nsimageview/1644708-imageviewwithimage?language=objc
func NSImageView_imageViewWithImage(image NSImageRef) NSImageView {
	ret := C.NSImageView_type_imageViewWithImage(
		objc.RefPointer(image),
	)

	return NSImageView_fromPointer(ret)

}

// NSNib_alloc
func NSNib_alloc() NSNib {
	ret := C.NSNib_type_alloc()

	return NSNib_fromPointer(ret)

}

// NSPasteboard_alloc
func NSPasteboard_alloc() NSPasteboard {
	ret := C.NSPasteboard_type_alloc()

	return NSPasteboard_fromPointer(ret)

}

// NSPasteboard_pasteboardByFilteringFile Creates a new pasteboard object that supplies the specified file in as many types as possible based on the available filter services.
// https://developer.apple.com/documentation/appkit/nspasteboard/1532744-pasteboardbyfilteringfile?language=objc
func NSPasteboard_pasteboardByFilteringFile(filename core.NSStringRef) NSPasteboard {
	ret := C.NSPasteboard_type_pasteboardByFilteringFile(
		objc.RefPointer(filename),
	)

	return NSPasteboard_fromPointer(ret)

}

// NSPasteboard_pasteboardByFilteringTypesInPasteboard Creates a new pasteboard object that supplies the specified pasteboard data in as many types as possible based on the available filter services.
// https://developer.apple.com/documentation/appkit/nspasteboard/1530088-pasteboardbyfilteringtypesinpast?language=objc
func NSPasteboard_pasteboardByFilteringTypesInPasteboard(pboard NSPasteboardRef) NSPasteboard {
	ret := C.NSPasteboard_type_pasteboardByFilteringTypesInPasteboard(
		objc.RefPointer(pboard),
	)

	return NSPasteboard_fromPointer(ret)

}

// NSPasteboard_pasteboardWithUniqueName Creates and returns a new pasteboard with a name that is guaranteed to be unique with respect to other pasteboards in the system.
// https://developer.apple.com/documentation/appkit/nspasteboard/1528936-pasteboardwithuniquename?language=objc
func NSPasteboard_pasteboardWithUniqueName() NSPasteboard {
	ret := C.NSPasteboard_type_pasteboardWithUniqueName()

	return NSPasteboard_fromPointer(ret)

}

// NSPasteboard_generalPasteboard The shared pasteboard object to use for general content.
// https://developer.apple.com/documentation/appkit/nspasteboard/1530091-generalpasteboard?language=objc
func NSPasteboard_generalPasteboard() NSPasteboard {
	ret := C.NSPasteboard_type_generalPasteboard()

	return NSPasteboard_fromPointer(ret)

}

// NSLayoutManager_alloc
func NSLayoutManager_alloc() NSLayoutManager {
	ret := C.NSLayoutManager_type_alloc()

	return NSLayoutManager_fromPointer(ret)

}

// NSMenu_alloc
func NSMenu_alloc() NSMenu {
	ret := C.NSMenu_type_alloc()

	return NSMenu_fromPointer(ret)

}

// NSMenu_menuBarVisible Returns a Boolean value that indicates whether the menu bar is visible.
// https://developer.apple.com/documentation/appkit/nsmenu/1518236-menubarvisible?language=objc
func NSMenu_menuBarVisible() bool {
	ret := C.NSMenu_type_menuBarVisible()

	return convertObjCBoolToGo(ret)

}

// NSMenu_setMenuBarVisible Sets whether the menu bar is visible and selectable by the user.
// https://developer.apple.com/documentation/appkit/nsmenu/1518200-setmenubarvisible?language=objc
func NSMenu_setMenuBarVisible(visible bool) {
	C.NSMenu_type_setMenuBarVisible(
		convertToObjCBool(visible),
	)

	return

}

// NSMenu_popUpContextMenu_withEvent_forView Displays a contextual menu over a view for an event.
// https://developer.apple.com/documentation/appkit/nsmenu/1518170-popupcontextmenu?language=objc
func NSMenu_popUpContextMenu_withEvent_forView(menu NSMenuRef, event NSEventRef, view NSViewRef) {
	C.NSMenu_type_popUpContextMenu_withEvent_forView(
		objc.RefPointer(menu),
		objc.RefPointer(event),
		objc.RefPointer(view),
	)

	return

}

// NSMenu_popUpContextMenu_withEvent_forView_withFont Displays a contextual menu over a view for an event using a specified font.
// https://developer.apple.com/documentation/appkit/nsmenu/1518165-popupcontextmenu?language=objc
func NSMenu_popUpContextMenu_withEvent_forView_withFont(menu NSMenuRef, event NSEventRef, view NSViewRef, font NSFontRef) {
	C.NSMenu_type_popUpContextMenu_withEvent_forView_withFont(
		objc.RefPointer(menu),
		objc.RefPointer(event),
		objc.RefPointer(view),
		objc.RefPointer(font),
	)

	return

}

// NSPopover_alloc
func NSPopover_alloc() NSPopover {
	ret := C.NSPopover_type_alloc()

	return NSPopover_fromPointer(ret)

}

// NSMenuItem_alloc
func NSMenuItem_alloc() NSMenuItem {
	ret := C.NSMenuItem_type_alloc()

	return NSMenuItem_fromPointer(ret)

}

// NSMenuItem_separatorItem Returns a menu item that is used to separate logical groups of menu commands.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514838-separatoritem?language=objc
func NSMenuItem_separatorItem() NSMenuItem {
	ret := C.NSMenuItem_type_separatorItem()

	return NSMenuItem_fromPointer(ret)

}

// NSMenuItem_usesUserKeyEquivalents Returns a Boolean value that indicates whether menu items conform to user preferences for key equivalents.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514811-usesuserkeyequivalents?language=objc
func NSMenuItem_usesUserKeyEquivalents() bool {
	ret := C.NSMenuItem_type_usesUserKeyEquivalents()

	return convertObjCBoolToGo(ret)

}

// NSMenuItem_setUsesUserKeyEquivalents Returns a Boolean value that indicates whether menu items conform to user preferences for key equivalents.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514811-usesuserkeyequivalents?language=objc
func NSMenuItem_setUsesUserKeyEquivalents(value bool) {
	C.NSMenuItem_type_setUsesUserKeyEquivalents(
		convertToObjCBool(value),
	)

	return

}

// NSRunningApplication_alloc
func NSRunningApplication_alloc() NSRunningApplication {
	ret := C.NSRunningApplication_type_alloc()

	return NSRunningApplication_fromPointer(ret)

}

// NSRunningApplication_runningApplicationsWithBundleIdentifier Returns an array of currently running applications with the specified bundle identifier.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1530798-runningapplicationswithbundleide?language=objc
func NSRunningApplication_runningApplicationsWithBundleIdentifier(bundleIdentifier core.NSStringRef) core.NSArray {
	ret := C.NSRunningApplication_type_runningApplicationsWithBundleIdentifier(
		objc.RefPointer(bundleIdentifier),
	)

	return core.NSArray_fromPointer(ret)

}

// NSRunningApplication_terminateAutomaticallyTerminableApplications Terminates invisibly running applications as if triggered by system memory pressure.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1529538-terminateautomaticallyterminable?language=objc
func NSRunningApplication_terminateAutomaticallyTerminableApplications() {
	C.NSRunningApplication_type_terminateAutomaticallyTerminableApplications()

	return

}

// NSRunningApplication_currentApplication Returns an NSRunningApplication representing this application.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1533604-currentapplication?language=objc
func NSRunningApplication_currentApplication() NSRunningApplication {
	ret := C.NSRunningApplication_type_currentApplication()

	return NSRunningApplication_fromPointer(ret)

}

// NSScreen_alloc
func NSScreen_alloc() NSScreen {
	ret := C.NSScreen_type_alloc()

	return NSScreen_fromPointer(ret)

}

// NSScreen_mainScreen Returns the screen object containing the window with the keyboard focus.
// https://developer.apple.com/documentation/appkit/nsscreen/1388371-mainscreen?language=objc
func NSScreen_mainScreen() NSScreen {
	ret := C.NSScreen_type_mainScreen()

	return NSScreen_fromPointer(ret)

}

// NSScreen_deepestScreen Returns a screen object representing the screen that can best represent color.
// https://developer.apple.com/documentation/appkit/nsscreen/1388374-deepestscreen?language=objc
func NSScreen_deepestScreen() NSScreen {
	ret := C.NSScreen_type_deepestScreen()

	return NSScreen_fromPointer(ret)

}

// NSScreen_screens Returns an array of screen objects representing all of the screens available on the system.
// https://developer.apple.com/documentation/appkit/nsscreen/1388393-screens?language=objc
func NSScreen_screens() core.NSArray {
	ret := C.NSScreen_type_screens()

	return core.NSArray_fromPointer(ret)

}

// NSScreen_screensHaveSeparateSpaces Returns a Boolean value indicating whether each screen can have its own set of spaces.
// https://developer.apple.com/documentation/appkit/nsscreen/1388365-screenshaveseparatespaces?language=objc
func NSScreen_screensHaveSeparateSpaces() bool {
	ret := C.NSScreen_type_screensHaveSeparateSpaces()

	return convertObjCBoolToGo(ret)

}

// NSStatusBar_alloc
func NSStatusBar_alloc() NSStatusBar {
	ret := C.NSStatusBar_type_alloc()

	return NSStatusBar_fromPointer(ret)

}

// NSStatusBar_systemStatusBar Returns the system-wide status bar located in the menu bar.
// https://developer.apple.com/documentation/appkit/nsstatusbar/1530619-systemstatusbar?language=objc
func NSStatusBar_systemStatusBar() NSStatusBar {
	ret := C.NSStatusBar_type_systemStatusBar()

	return NSStatusBar_fromPointer(ret)

}

// NSStatusBarButton_alloc
func NSStatusBarButton_alloc() NSStatusBarButton {
	ret := C.NSStatusBarButton_type_alloc()

	return NSStatusBarButton_fromPointer(ret)

}

// NSStatusItem_alloc
func NSStatusItem_alloc() NSStatusItem {
	ret := C.NSStatusItem_type_alloc()

	return NSStatusItem_fromPointer(ret)

}

// NSText_alloc
func NSText_alloc() NSText {
	ret := C.NSText_type_alloc()

	return NSText_fromPointer(ret)

}

// NSTextField_alloc
func NSTextField_alloc() NSTextField {
	ret := C.NSTextField_type_alloc()

	return NSTextField_fromPointer(ret)

}

// NSTextField_labelWithAttributedString Creates a text field for use as a static label that displays styled text, doesn’t wrap, and doesn’t have selectable text.
// https://developer.apple.com/documentation/appkit/nstextfield/1644658-labelwithattributedstring?language=objc
func NSTextField_labelWithAttributedString(attributedStringValue core.NSAttributedStringRef) NSTextField {
	ret := C.NSTextField_type_labelWithAttributedString(
		objc.RefPointer(attributedStringValue),
	)

	return NSTextField_fromPointer(ret)

}

// NSTextField_labelWithString Initializes a text field for use as a static label that uses the system default font, doesn’t wrap, and doesn’t have selectable text.
// https://developer.apple.com/documentation/appkit/nstextfield/1644377-labelwithstring?language=objc
func NSTextField_labelWithString(stringValue core.NSStringRef) NSTextField {
	ret := C.NSTextField_type_labelWithString(
		objc.RefPointer(stringValue),
	)

	return NSTextField_fromPointer(ret)

}

// NSTextField_textFieldWithString Initializes a single-line editable text field for user input using the system default font and standard visual appearance.
// https://developer.apple.com/documentation/appkit/nstextfield/1644706-textfieldwithstring?language=objc
func NSTextField_textFieldWithString(stringValue core.NSStringRef) NSTextField {
	ret := C.NSTextField_type_textFieldWithString(
		objc.RefPointer(stringValue),
	)

	return NSTextField_fromPointer(ret)

}

// NSTextField_wrappingLabelWithString Initializes a text field for use as a multiline static label with selectable text that uses the system default font.
// https://developer.apple.com/documentation/appkit/nstextfield/1644543-wrappinglabelwithstring?language=objc
func NSTextField_wrappingLabelWithString(stringValue core.NSStringRef) NSTextField {
	ret := C.NSTextField_type_wrappingLabelWithString(
		objc.RefPointer(stringValue),
	)

	return NSTextField_fromPointer(ret)

}

// NSTextContainer_alloc
func NSTextContainer_alloc() NSTextContainer {
	ret := C.NSTextContainer_type_alloc()

	return NSTextContainer_fromPointer(ret)

}

// NSViewController_alloc
func NSViewController_alloc() NSViewController {
	ret := C.NSViewController_type_alloc()

	return NSViewController_fromPointer(ret)

}

// NSVisualEffectView_alloc
func NSVisualEffectView_alloc() NSVisualEffectView {
	ret := C.NSVisualEffectView_type_alloc()

	return NSVisualEffectView_fromPointer(ret)

}

// NSWindow_alloc
func NSWindow_alloc() NSWindow {
	ret := C.NSWindow_type_alloc()

	return NSWindow_fromPointer(ret)

}

// NSWindow_windowWithContentViewController Creates a titled window that contains the specified content view controller.
// https://developer.apple.com/documentation/appkit/nswindow/1419551-windowwithcontentviewcontroller?language=objc
func NSWindow_windowWithContentViewController(contentViewController NSViewControllerRef) NSWindow {
	ret := C.NSWindow_type_windowWithContentViewController(
		objc.RefPointer(contentViewController),
	)

	return NSWindow_fromPointer(ret)

}

// NSWindow_contentRectForFrameRect_styleMask Returns the content rectangle used by a window with a given frame rectangle and window style.
// https://developer.apple.com/documentation/appkit/nswindow/1419586-contentrectforframerect?language=objc
func NSWindow_contentRectForFrameRect_styleMask(fRect core.NSRect, style core.NSUInteger) core.NSRect {
	ret := C.NSWindow_type_contentRectForFrameRect_styleMask(
		*(*C.NSRect)(unsafe.Pointer(&fRect)),
		C.ulong(style),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// NSWindow_frameRectForContentRect_styleMask Returns the frame rectangle used by a window with a given content rectangle and window style.
// https://developer.apple.com/documentation/appkit/nswindow/1419372-framerectforcontentrect?language=objc
func NSWindow_frameRectForContentRect_styleMask(cRect core.NSRect, style core.NSUInteger) core.NSRect {
	ret := C.NSWindow_type_frameRectForContentRect_styleMask(
		*(*C.NSRect)(unsafe.Pointer(&cRect)),
		C.ulong(style),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// NSWindow_minFrameWidthWithTitle_styleMask Returns the minimum width a window’s frame rectangle must have for it to display a title, with a given window style.
// https://developer.apple.com/documentation/appkit/nswindow/1419294-minframewidthwithtitle?language=objc
func NSWindow_minFrameWidthWithTitle_styleMask(title core.NSStringRef, style core.NSUInteger) core.CGFloat {
	ret := C.NSWindow_type_minFrameWidthWithTitle_styleMask(
		objc.RefPointer(title),
		C.ulong(style),
	)

	return core.CGFloat(ret)

}

// NSWindow_windowNumberAtPoint_belowWindowWithWindowNumber Returns the number of the frontmost window that would be hit by a mouse-down at the specified screen location.
// https://developer.apple.com/documentation/appkit/nswindow/1419210-windownumberatpoint?language=objc
func NSWindow_windowNumberAtPoint_belowWindowWithWindowNumber(point core.NSPoint, windowNumber core.NSInteger) core.NSInteger {
	ret := C.NSWindow_type_windowNumberAtPoint_belowWindowWithWindowNumber(
		*(*C.NSPoint)(unsafe.Pointer(&point)),
		C.long(windowNumber),
	)

	return core.NSInteger(ret)

}

// NSWindow_allowsAutomaticWindowTabbing A Boolean value that indicates whether the app can automatically organize windows into tabs.
// https://developer.apple.com/documentation/appkit/nswindow/1646657-allowsautomaticwindowtabbing?language=objc
func NSWindow_allowsAutomaticWindowTabbing() bool {
	ret := C.NSWindow_type_allowsAutomaticWindowTabbing()

	return convertObjCBoolToGo(ret)

}

// NSWindow_setAllowsAutomaticWindowTabbing A Boolean value that indicates whether the app can automatically organize windows into tabs.
// https://developer.apple.com/documentation/appkit/nswindow/1646657-allowsautomaticwindowtabbing?language=objc
func NSWindow_setAllowsAutomaticWindowTabbing(value bool) {
	C.NSWindow_type_setAllowsAutomaticWindowTabbing(
		convertToObjCBool(value),
	)

	return

}

// NSWorkspace_alloc
func NSWorkspace_alloc() NSWorkspace {
	ret := C.NSWorkspace_type_alloc()

	return NSWorkspace_fromPointer(ret)

}

// NSWorkspace_sharedWorkspace The shared workspace object.
// https://developer.apple.com/documentation/appkit/nsworkspace/1530344-sharedworkspace?language=objc
func NSWorkspace_sharedWorkspace() NSWorkspace {
	ret := C.NSWorkspace_type_sharedWorkspace()

	return NSWorkspace_fromPointer(ret)

}

// NSColor_alloc
func NSColor_alloc() NSColor {
	ret := C.NSColor_type_alloc()

	return NSColor_fromPointer(ret)

}

// NSColor_colorFromPasteboard Creates a color object from color data currently on the pasteboard.
// https://developer.apple.com/documentation/appkit/nscolor/1535057-colorfrompasteboard?language=objc
func NSColor_colorFromPasteboard(pasteBoard NSPasteboardRef) NSColor {
	ret := C.NSColor_type_colorFromPasteboard(
		objc.RefPointer(pasteBoard),
	)

	return NSColor_fromPointer(ret)

}

// NSColor_colorWithRed_green_blue_alpha
func NSColor_colorWithRed_green_blue_alpha(red core.CGFloat, green core.CGFloat, blue core.CGFloat, alpha core.CGFloat) NSColor {
	ret := C.NSColor_type_colorWithRed_green_blue_alpha(
		C.double(red),
		C.double(green),
		C.double(blue),
		C.double(alpha),
	)

	return NSColor_fromPointer(ret)

}

// NSColor_ignoresAlpha A Boolean value that indicates whether the app supports alpha.
// https://developer.apple.com/documentation/appkit/nscolor/1533565-ignoresalpha?language=objc
func NSColor_ignoresAlpha() bool {
	ret := C.NSColor_type_ignoresAlpha()

	return convertObjCBoolToGo(ret)

}

// NSColor_setIgnoresAlpha A Boolean value that indicates whether the app supports alpha.
// https://developer.apple.com/documentation/appkit/nscolor/1533565-ignoresalpha?language=objc
func NSColor_setIgnoresAlpha(value bool) {
	C.NSColor_type_setIgnoresAlpha(
		convertToObjCBool(value),
	)

	return

}

// NSColor_systemCyanColor
// https://developer.apple.com/documentation/appkit/nscolor/3816005-systemcyancolor?language=objc
func NSColor_systemCyanColor() NSColor {
	ret := C.NSColor_type_systemCyanColor()

	return NSColor_fromPointer(ret)

}

// NSColor_systemMintColor
// https://developer.apple.com/documentation/appkit/nscolor/3816006-systemmintcolor?language=objc
func NSColor_systemMintColor() NSColor {
	ret := C.NSColor_type_systemMintColor()

	return NSColor_fromPointer(ret)

}

// NSColor_clearColor Returns a color object whose grayscale and alpha values are both 0.0.
// https://developer.apple.com/documentation/appkit/nscolor/1527217-clearcolor?language=objc
func NSColor_clearColor() NSColor {
	ret := C.NSColor_type_clearColor()

	return NSColor_fromPointer(ret)

}

// NSTextView_alloc
func NSTextView_alloc() NSTextView {
	ret := C.NSTextView_type_alloc()

	return NSTextView_fromPointer(ret)

}

// NSTextView_registerForServices Registers send and return types for the Services facility.
// https://developer.apple.com/documentation/appkit/nstextview/1449507-registerforservices?language=objc
func NSTextView_registerForServices() {
	C.NSTextView_type_registerForServices()

	return

}

// NSTextView_fieldEditor
// https://developer.apple.com/documentation/appkit/nstextview/2990525-fieldeditor?language=objc
func NSTextView_fieldEditor() NSTextView {
	ret := C.NSTextView_type_fieldEditor()

	return NSTextView_fromPointer(ret)

}

// NSTextView_stronglyReferencesTextStorage
// https://developer.apple.com/documentation/appkit/nstextview/2269433-stronglyreferencestextstorage?language=objc
func NSTextView_stronglyReferencesTextStorage() bool {
	ret := C.NSTextView_type_stronglyReferencesTextStorage()

	return convertObjCBoolToGo(ret)

}

// NSView_alloc
func NSView_alloc() NSView {
	ret := C.NSView_type_alloc()

	return NSView_fromPointer(ret)

}

// NSView_requiresConstraintBasedLayout Returns a Boolean value indicating whether the view depends on the constraint-based layout system.
// https://developer.apple.com/documentation/appkit/nsview/1526926-requiresconstraintbasedlayout?language=objc
func NSView_requiresConstraintBasedLayout() bool {
	ret := C.NSView_type_requiresConstraintBasedLayout()

	return convertObjCBoolToGo(ret)

}

// NSView_focusView Returns the currently focused NSView object, or nil if there is none.
// https://developer.apple.com/documentation/appkit/nsview/1483662-focusview?language=objc
func NSView_focusView() NSView {
	ret := C.NSView_type_focusView()

	return NSView_fromPointer(ret)

}

// NSView_defaultMenu Overridden by subclasses to return the default pop-up menu for instances of the receiving class.
// https://developer.apple.com/documentation/appkit/nsview/1483417-defaultmenu?language=objc
func NSView_defaultMenu() NSMenu {
	ret := C.NSView_type_defaultMenu()

	return NSMenu_fromPointer(ret)

}

// NSView_compatibleWithResponsiveScrolling
// https://developer.apple.com/documentation/appkit/nsview/2870005-compatiblewithresponsivescrollin?language=objc
func NSView_compatibleWithResponsiveScrolling() bool {
	ret := C.NSView_type_compatibleWithResponsiveScrolling()

	return convertObjCBoolToGo(ret)

}

type NSBundleRef interface {
	Pointer() uintptr
	Init_asNSBundle() NSBundle
}

type gen_NSBundle struct {
	objc.Object
}

func NSBundle_fromPointer(ptr unsafe.Pointer) NSBundle {
	return NSBundle{gen_NSBundle{
		objc.Object_fromPointer(ptr),
	}}
}

func NSBundle_fromRef(ref objc.Ref) NSBundle {
	return NSBundle_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// URLForAuxiliaryExecutable Returns the file URL of the executable with the specified name in the receiver’s bundle.
// https://developer.apple.com/documentation/foundation/nsbundle/1411412-urlforauxiliaryexecutable?language=objc
func (x gen_NSBundle) URLForAuxiliaryExecutable(
	executableName core.NSStringRef,
) core.NSURL {
	ret := C.NSBundle_inst_URLForAuxiliaryExecutable(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(executableName),
	)

	return core.NSURL_fromPointer(ret)

}

// URLForResource_withExtension Returns the file URL for the resource identified by the specified name and file extension.
// https://developer.apple.com/documentation/foundation/nsbundle/1411540-urlforresource?language=objc
func (x gen_NSBundle) URLForResource_withExtension(
	name core.NSStringRef,
	ext core.NSStringRef,
) core.NSURL {
	ret := C.NSBundle_inst_URLForResource_withExtension(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
	)

	return core.NSURL_fromPointer(ret)

}

// URLForResource_withExtension_subdirectory Returns the file URL for the resource file identified by the specified name and extension and residing in a given bundle directory.
// https://developer.apple.com/documentation/foundation/nsbundle/1416712-urlforresource?language=objc
func (x gen_NSBundle) URLForResource_withExtension_subdirectory(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
) core.NSURL {
	ret := C.NSBundle_inst_URLForResource_withExtension_subdirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)

	return core.NSURL_fromPointer(ret)

}

// URLForResource_withExtension_subdirectory_localization Returns the file URL for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization.
// https://developer.apple.com/documentation/foundation/nsbundle/1417378-urlforresource?language=objc
func (x gen_NSBundle) URLForResource_withExtension_subdirectory_localization(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) core.NSURL {
	ret := C.NSBundle_inst_URLForResource_withExtension_subdirectory_localization(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)

	return core.NSURL_fromPointer(ret)

}

// URLsForResourcesWithExtension_subdirectory Returns an array of file URLs for all resources identified by the specified file extension and located in the specified bundle subdirectory.
// https://developer.apple.com/documentation/foundation/nsbundle/1407424-urlsforresourceswithextension?language=objc
func (x gen_NSBundle) URLsForResourcesWithExtension_subdirectory(
	ext core.NSStringRef,
	subpath core.NSStringRef,
) core.NSArray {
	ret := C.NSBundle_inst_URLsForResourcesWithExtension_subdirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)

	return core.NSArray_fromPointer(ret)

}

// URLsForResourcesWithExtension_subdirectory_localization Returns an array containing the file URLs for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization.
// https://developer.apple.com/documentation/foundation/nsbundle/1414688-urlsforresourceswithextension?language=objc
func (x gen_NSBundle) URLsForResourcesWithExtension_subdirectory_localization(
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) core.NSArray {
	ret := C.NSBundle_inst_URLsForResourcesWithExtension_subdirectory_localization(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)

	return core.NSArray_fromPointer(ret)

}

// InitWithPath_asNSBundle Returns an NSBundle object initialized to correspond to the specified directory.
// https://developer.apple.com/documentation/foundation/nsbundle/1412741-initwithpath?language=objc
func (x gen_NSBundle) InitWithPath_asNSBundle(
	path core.NSStringRef,
) NSBundle {
	ret := C.NSBundle_inst_initWithPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)

	return NSBundle_fromPointer(ret)

}

// InitWithURL_asNSBundle Returns an NSBundle object initialized to correspond to the specified file URL.
// https://developer.apple.com/documentation/foundation/nsbundle/1409352-initwithurl?language=objc
func (x gen_NSBundle) InitWithURL_asNSBundle(
	url core.NSURLRef,
) NSBundle {
	ret := C.NSBundle_inst_initWithURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)

	return NSBundle_fromPointer(ret)

}

// Load Dynamically loads the bundle’s executable code into a running program, if the code has not already been loaded.
// https://developer.apple.com/documentation/foundation/nsbundle/1415927-load?language=objc
func (x gen_NSBundle) Load() bool {
	ret := C.NSBundle_inst_load(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// LoadNibNamed_owner_options Unarchives the contents of a nib file located in the receiver's bundle.
// https://developer.apple.com/documentation/foundation/nsbundle/1618147-loadnibnamed?language=objc
func (x gen_NSBundle) LoadNibNamed_owner_options(
	name core.NSStringRef,
	owner objc.Ref,
	options core.NSDictionaryRef,
) core.NSArray {
	ret := C.NSBundle_inst_loadNibNamed_owner_options(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(owner),
		objc.RefPointer(options),
	)

	return core.NSArray_fromPointer(ret)

}

// LocalizedAttributedStringForKey_value_table
// https://developer.apple.com/documentation/foundation/nsbundle/3746904-localizedattributedstringforkey?language=objc
func (x gen_NSBundle) LocalizedAttributedStringForKey_value_table(
	key core.NSStringRef,
	value core.NSStringRef,
	tableName core.NSStringRef,
) core.NSAttributedString {
	ret := C.NSBundle_inst_localizedAttributedStringForKey_value_table(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
		objc.RefPointer(value),
		objc.RefPointer(tableName),
	)

	return core.NSAttributedString_fromPointer(ret)

}

// LocalizedStringForKey_value_table Returns a localized version of the string designated by the specified key and residing in the specified table.
// https://developer.apple.com/documentation/foundation/nsbundle/1417694-localizedstringforkey?language=objc
func (x gen_NSBundle) LocalizedStringForKey_value_table(
	key core.NSStringRef,
	value core.NSStringRef,
	tableName core.NSStringRef,
) core.NSString {
	ret := C.NSBundle_inst_localizedStringForKey_value_table(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
		objc.RefPointer(value),
		objc.RefPointer(tableName),
	)

	return core.NSString_fromPointer(ret)

}

// ObjectForInfoDictionaryKey Returns the value associated with the specified key in the receiver's information property list.
// https://developer.apple.com/documentation/foundation/nsbundle/1408696-objectforinfodictionarykey?language=objc
func (x gen_NSBundle) ObjectForInfoDictionaryKey(
	key core.NSStringRef,
) objc.Object {
	ret := C.NSBundle_inst_objectForInfoDictionaryKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)

	return objc.Object_fromPointer(ret)

}

// PathForAuxiliaryExecutable Returns the full pathname of the executable with the specified name in the receiver’s bundle.
// https://developer.apple.com/documentation/foundation/nsbundle/1415214-pathforauxiliaryexecutable?language=objc
func (x gen_NSBundle) PathForAuxiliaryExecutable(
	executableName core.NSStringRef,
) core.NSString {
	ret := C.NSBundle_inst_pathForAuxiliaryExecutable(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(executableName),
	)

	return core.NSString_fromPointer(ret)

}

// PathForResource_ofType Returns the full pathname for the resource identified by the specified name and file extension.
// https://developer.apple.com/documentation/foundation/nsbundle/1410989-pathforresource?language=objc
func (x gen_NSBundle) PathForResource_ofType(
	name core.NSStringRef,
	ext core.NSStringRef,
) core.NSString {
	ret := C.NSBundle_inst_pathForResource_ofType(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
	)

	return core.NSString_fromPointer(ret)

}

// PathForResource_ofType_inDirectory Returns the full pathname for the resource identified by the specified name and file extension and located in the specified bundle subdirectory.
// https://developer.apple.com/documentation/foundation/nsbundle/1409670-pathforresource?language=objc
func (x gen_NSBundle) PathForResource_ofType_inDirectory(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
) core.NSString {
	ret := C.NSBundle_inst_pathForResource_ofType_inDirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)

	return core.NSString_fromPointer(ret)

}

// PathForResource_ofType_inDirectory_forLocalization Returns the full pathname for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization.
// https://developer.apple.com/documentation/foundation/nsbundle/1413471-pathforresource?language=objc
func (x gen_NSBundle) PathForResource_ofType_inDirectory_forLocalization(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) core.NSString {
	ret := C.NSBundle_inst_pathForResource_ofType_inDirectory_forLocalization(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)

	return core.NSString_fromPointer(ret)

}

// PathsForResourcesOfType_inDirectory Returns an array containing the pathnames for all bundle resources having the specified filename extension and residing in the resource subdirectory.
// https://developer.apple.com/documentation/foundation/nsbundle/1413058-pathsforresourcesoftype?language=objc
func (x gen_NSBundle) PathsForResourcesOfType_inDirectory(
	ext core.NSStringRef,
	subpath core.NSStringRef,
) core.NSArray {
	ret := C.NSBundle_inst_pathsForResourcesOfType_inDirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)

	return core.NSArray_fromPointer(ret)

}

// PathsForResourcesOfType_inDirectory_forLocalization Returns an array containing the file for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization.
// https://developer.apple.com/documentation/foundation/nsbundle/1416940-pathsforresourcesoftype?language=objc
func (x gen_NSBundle) PathsForResourcesOfType_inDirectory_forLocalization(
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) core.NSArray {
	ret := C.NSBundle_inst_pathsForResourcesOfType_inDirectory_forLocalization(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)

	return core.NSArray_fromPointer(ret)

}

// Unload Unloads the code associated with the receiver.
// https://developer.apple.com/documentation/foundation/nsbundle/1412388-unload?language=objc
func (x gen_NSBundle) Unload() bool {
	ret := C.NSBundle_inst_unload(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Init_asNSBundle
func (x gen_NSBundle) Init_asNSBundle() NSBundle {
	ret := C.NSBundle_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSBundle_fromPointer(ret)

}

// ResourceURL The file URL of the bundle’s subdirectory containing resource files.
// https://developer.apple.com/documentation/foundation/nsbundle/1414821-resourceurl?language=objc
func (x gen_NSBundle) ResourceURL() core.NSURL {
	ret := C.NSBundle_inst_resourceURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_fromPointer(ret)

}

// ExecutableURL The file URL of the receiver's executable file.
// https://developer.apple.com/documentation/foundation/nsbundle/1410470-executableurl?language=objc
func (x gen_NSBundle) ExecutableURL() core.NSURL {
	ret := C.NSBundle_inst_executableURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_fromPointer(ret)

}

// PrivateFrameworksURL The file URL of the bundle’s subdirectory containing private frameworks.
// https://developer.apple.com/documentation/foundation/nsbundle/1417617-privateframeworksurl?language=objc
func (x gen_NSBundle) PrivateFrameworksURL() core.NSURL {
	ret := C.NSBundle_inst_privateFrameworksURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_fromPointer(ret)

}

// SharedFrameworksURL The file URL of the receiver's subdirectory containing shared frameworks.
// https://developer.apple.com/documentation/foundation/nsbundle/1411774-sharedframeworksurl?language=objc
func (x gen_NSBundle) SharedFrameworksURL() core.NSURL {
	ret := C.NSBundle_inst_sharedFrameworksURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_fromPointer(ret)

}

// BuiltInPlugInsURL The file URL of the receiver's subdirectory containing plug-ins.
// https://developer.apple.com/documentation/foundation/nsbundle/1409603-builtinpluginsurl?language=objc
func (x gen_NSBundle) BuiltInPlugInsURL() core.NSURL {
	ret := C.NSBundle_inst_builtInPlugInsURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_fromPointer(ret)

}

// SharedSupportURL The file URL of the bundle’s subdirectory containing shared support files.
// https://developer.apple.com/documentation/foundation/nsbundle/1416823-sharedsupporturl?language=objc
func (x gen_NSBundle) SharedSupportURL() core.NSURL {
	ret := C.NSBundle_inst_sharedSupportURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_fromPointer(ret)

}

// AppStoreReceiptURL The file URL for the bundle’s App Store receipt.
// https://developer.apple.com/documentation/foundation/nsbundle/1407276-appstorereceipturl?language=objc
func (x gen_NSBundle) AppStoreReceiptURL() core.NSURL {
	ret := C.NSBundle_inst_appStoreReceiptURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_fromPointer(ret)

}

// ResourcePath The full pathname of the bundle’s subdirectory containing resources.
// https://developer.apple.com/documentation/foundation/nsbundle/1417723-resourcepath?language=objc
func (x gen_NSBundle) ResourcePath() core.NSString {
	ret := C.NSBundle_inst_resourcePath(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// ExecutablePath The full pathname of the receiver's executable file.
// https://developer.apple.com/documentation/foundation/nsbundle/1409078-executablepath?language=objc
func (x gen_NSBundle) ExecutablePath() core.NSString {
	ret := C.NSBundle_inst_executablePath(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// PrivateFrameworksPath The full pathname of the bundle’s subdirectory containing private frameworks.
// https://developer.apple.com/documentation/foundation/nsbundle/1415562-privateframeworkspath?language=objc
func (x gen_NSBundle) PrivateFrameworksPath() core.NSString {
	ret := C.NSBundle_inst_privateFrameworksPath(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SharedFrameworksPath The full pathname of the bundle’s subdirectory containing shared frameworks.
// https://developer.apple.com/documentation/foundation/nsbundle/1417226-sharedframeworkspath?language=objc
func (x gen_NSBundle) SharedFrameworksPath() core.NSString {
	ret := C.NSBundle_inst_sharedFrameworksPath(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// BuiltInPlugInsPath The full pathname of the receiver's subdirectory containing plug-ins.
// https://developer.apple.com/documentation/foundation/nsbundle/1408900-builtinpluginspath?language=objc
func (x gen_NSBundle) BuiltInPlugInsPath() core.NSString {
	ret := C.NSBundle_inst_builtInPlugInsPath(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SharedSupportPath The full pathname of the bundle’s subdirectory containing shared support files.
// https://developer.apple.com/documentation/foundation/nsbundle/1411609-sharedsupportpath?language=objc
func (x gen_NSBundle) SharedSupportPath() core.NSString {
	ret := C.NSBundle_inst_sharedSupportPath(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// BundleURL The full URL of the receiver’s bundle directory.
// https://developer.apple.com/documentation/foundation/nsbundle/1415654-bundleurl?language=objc
func (x gen_NSBundle) BundleURL() core.NSURL {
	ret := C.NSBundle_inst_bundleURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_fromPointer(ret)

}

// BundlePath The full pathname of the receiver’s bundle directory.
// https://developer.apple.com/documentation/foundation/nsbundle/1407973-bundlepath?language=objc
func (x gen_NSBundle) BundlePath() core.NSString {
	ret := C.NSBundle_inst_bundlePath(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// BundleIdentifier The receiver’s bundle identifier.
// https://developer.apple.com/documentation/foundation/nsbundle/1418023-bundleidentifier?language=objc
func (x gen_NSBundle) BundleIdentifier() core.NSString {
	ret := C.NSBundle_inst_bundleIdentifier(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// InfoDictionary A dictionary, constructed from the bundle’s Info.plist file, that contains information about the receiver.
// https://developer.apple.com/documentation/foundation/nsbundle/1413477-infodictionary?language=objc
func (x gen_NSBundle) InfoDictionary() core.NSDictionary {
	ret := C.NSBundle_inst_infoDictionary(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_fromPointer(ret)

}

// Localizations A list of all the localizations contained in the bundle.
// https://developer.apple.com/documentation/foundation/nsbundle/1417415-localizations?language=objc
func (x gen_NSBundle) Localizations() core.NSArray {
	ret := C.NSBundle_inst_localizations(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// PreferredLocalizations An ordered list of preferred localizations contained in the bundle.
// https://developer.apple.com/documentation/foundation/nsbundle/1413220-preferredlocalizations?language=objc
func (x gen_NSBundle) PreferredLocalizations() core.NSArray {
	ret := C.NSBundle_inst_preferredLocalizations(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// DevelopmentLocalization The localization for the development language.
// https://developer.apple.com/documentation/foundation/nsbundle/1417526-developmentlocalization?language=objc
func (x gen_NSBundle) DevelopmentLocalization() core.NSString {
	ret := C.NSBundle_inst_developmentLocalization(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// LocalizedInfoDictionary A dictionary with the keys from the bundle’s localized property list.
// https://developer.apple.com/documentation/foundation/nsbundle/1407645-localizedinfodictionary?language=objc
func (x gen_NSBundle) LocalizedInfoDictionary() core.NSDictionary {
	ret := C.NSBundle_inst_localizedInfoDictionary(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_fromPointer(ret)

}

// ExecutableArchitectures An array of numbers indicating the architecture types supported by the bundle’s executable.
// https://developer.apple.com/documentation/foundation/nsbundle/1415499-executablearchitectures?language=objc
func (x gen_NSBundle) ExecutableArchitectures() core.NSArray {
	ret := C.NSBundle_inst_executableArchitectures(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// IsLoaded The load status of a bundle.
// https://developer.apple.com/documentation/foundation/nsbundle/1413594-loaded?language=objc
func (x gen_NSBundle) IsLoaded() bool {
	ret := C.NSBundle_inst_isLoaded(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

type NSSoundRef interface {
	Pointer() uintptr
	Init_asNSSound() NSSound
}

type gen_NSSound struct {
	objc.Object
}

func NSSound_fromPointer(ptr unsafe.Pointer) NSSound {
	return NSSound{gen_NSSound{
		objc.Object_fromPointer(ptr),
	}}
}

func NSSound_fromRef(ref objc.Ref) NSSound {
	return NSSound_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// InitWithContentsOfFile_byReference_asNSSound Initializes the receiver with the audio data located at a given filepath.
// https://developer.apple.com/documentation/appkit/nssound/1477274-initwithcontentsoffile?language=objc
func (x gen_NSSound) InitWithContentsOfFile_byReference_asNSSound(
	path core.NSStringRef,
	byRef bool,
) NSSound {
	ret := C.NSSound_inst_initWithContentsOfFile_byReference(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(byRef),
	)

	return NSSound_fromPointer(ret)

}

// InitWithContentsOfURL_byReference_asNSSound Initializes the receiver with the audio data located at a given URL.
// https://developer.apple.com/documentation/appkit/nssound/1477288-initwithcontentsofurl?language=objc
func (x gen_NSSound) InitWithContentsOfURL_byReference_asNSSound(
	url core.NSURLRef,
	byRef bool,
) NSSound {
	ret := C.NSSound_inst_initWithContentsOfURL_byReference(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
		convertToObjCBool(byRef),
	)

	return NSSound_fromPointer(ret)

}

// InitWithData_asNSSound Initializes the receiver with a given audio data.
// https://developer.apple.com/documentation/appkit/nssound/1477292-initwithdata?language=objc
func (x gen_NSSound) InitWithData_asNSSound(
	data core.NSDataRef,
) NSSound {
	ret := C.NSSound_inst_initWithData(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
	)

	return NSSound_fromPointer(ret)

}

// InitWithPasteboard_asNSSound Initializes the receiver with data from a pasteboard. The pasteboard should contain a type returned by NSSound. NSSound expects the data to have a proper magic number, sound header, and data for the formats it supports.
// https://developer.apple.com/documentation/appkit/nssound/1477294-initwithpasteboard?language=objc
func (x gen_NSSound) InitWithPasteboard_asNSSound(
	pasteboard NSPasteboardRef,
) NSSound {
	ret := C.NSSound_inst_initWithPasteboard(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pasteboard),
	)

	return NSSound_fromPointer(ret)

}

// Pause Pauses audio playback.
// https://developer.apple.com/documentation/appkit/nssound/1477307-pause?language=objc
func (x gen_NSSound) Pause() bool {
	ret := C.NSSound_inst_pause(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Play Initiates audio playback.
// https://developer.apple.com/documentation/appkit/nssound/1477322-play?language=objc
func (x gen_NSSound) Play() bool {
	ret := C.NSSound_inst_play(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Resume Resumes audio playback.
// https://developer.apple.com/documentation/appkit/nssound/1477336-resume?language=objc
func (x gen_NSSound) Resume() bool {
	ret := C.NSSound_inst_resume(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Stop Concludes audio playback.
// https://developer.apple.com/documentation/appkit/nssound/1477282-stop?language=objc
func (x gen_NSSound) Stop() bool {
	ret := C.NSSound_inst_stop(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// WriteToPasteboard Writes the receiver’s data to a pasteboard.
// https://developer.apple.com/documentation/appkit/nssound/1477338-writetopasteboard?language=objc
func (x gen_NSSound) WriteToPasteboard(
	pasteboard NSPasteboardRef,
) {
	C.NSSound_inst_writeToPasteboard(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pasteboard),
	)

	return

}

// Init_asNSSound
func (x gen_NSSound) Init_asNSSound() NSSound {
	ret := C.NSSound_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSSound_fromPointer(ret)

}

// Delegate The sound’s delegate.
// https://developer.apple.com/documentation/appkit/nssound/1477300-delegate?language=objc
func (x gen_NSSound) Delegate() objc.Object {
	ret := C.NSSound_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetDelegate The sound’s delegate.
// https://developer.apple.com/documentation/appkit/nssound/1477300-delegate?language=objc
func (x gen_NSSound) SetDelegate(
	value objc.Ref,
) {
	C.NSSound_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Loops A Boolean that indicates whether the sound restarts playback when it reaches the end of its content.
// https://developer.apple.com/documentation/appkit/nssound/1477311-loops?language=objc
func (x gen_NSSound) Loops() bool {
	ret := C.NSSound_inst_loops(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetLoops A Boolean that indicates whether the sound restarts playback when it reaches the end of its content.
// https://developer.apple.com/documentation/appkit/nssound/1477311-loops?language=objc
func (x gen_NSSound) SetLoops(
	value bool,
) {
	C.NSSound_inst_setLoops(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsPlaying A Boolean that indicates whether the sound is playing its audio data.
// https://developer.apple.com/documentation/appkit/nssound/1477302-playing?language=objc
func (x gen_NSSound) IsPlaying() bool {
	ret := C.NSSound_inst_isPlaying(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

type NSApplicationRef interface {
	Pointer() uintptr
	Init_asNSApplication() NSApplication
}

type gen_NSApplication struct {
	objc.Object
}

func NSApplication_fromPointer(ptr unsafe.Pointer) NSApplication {
	return NSApplication{gen_NSApplication{
		objc.Object_fromPointer(ptr),
	}}
}

func NSApplication_fromRef(ref objc.Ref) NSApplication {
	return NSApplication_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// ActivateContextHelpMode Places the receiver in context-sensitive help mode.
// https://developer.apple.com/documentation/appkit/nsapplication/1500925-activatecontexthelpmode?language=objc
func (x gen_NSApplication) ActivateContextHelpMode(
	sender objc.Ref,
) {
	C.NSApplication_inst_activateContextHelpMode(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ActivateIgnoringOtherApps Makes the receiver the active app.
// https://developer.apple.com/documentation/appkit/nsapplication/1428468-activateignoringotherapps?language=objc
func (x gen_NSApplication) ActivateIgnoringOtherApps(
	flag bool,
) {
	C.NSApplication_inst_activateIgnoringOtherApps(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)

	return

}

// ActivationPolicy Returns the app’s activation policy.
// https://developer.apple.com/documentation/appkit/nsapplication/1428703-activationpolicy?language=objc
func (x gen_NSApplication) ActivationPolicy() core.NSInteger {
	ret := C.NSApplication_inst_activationPolicy(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// CancelUserAttentionRequest Cancels a previous user attention request.
// https://developer.apple.com/documentation/appkit/nsapplication/1428683-canceluserattentionrequest?language=objc
func (x gen_NSApplication) CancelUserAttentionRequest(
	request core.NSInteger,
) {
	C.NSApplication_inst_cancelUserAttentionRequest(
		unsafe.Pointer(x.Pointer()),
		C.long(request),
	)

	return

}

// Deactivate Deactivates the receiver.
// https://developer.apple.com/documentation/appkit/nsapplication/1428428-deactivate?language=objc
func (x gen_NSApplication) Deactivate() {
	C.NSApplication_inst_deactivate(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// DisableRelaunchOnLogin Disables relaunching the app on login.
// https://developer.apple.com/documentation/appkit/nsapplication/1428376-disablerelaunchonlogin?language=objc
func (x gen_NSApplication) DisableRelaunchOnLogin() {
	C.NSApplication_inst_disableRelaunchOnLogin(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// EnableRelaunchOnLogin Enables relaunching the app on login.
// https://developer.apple.com/documentation/appkit/nsapplication/1428453-enablerelaunchonlogin?language=objc
func (x gen_NSApplication) EnableRelaunchOnLogin() {
	C.NSApplication_inst_enableRelaunchOnLogin(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// FinishLaunching Activates the app, opens any files specified by the NSOpen user default, and unhighlights the app’s icon.
// https://developer.apple.com/documentation/appkit/nsapplication/1428771-finishlaunching?language=objc
func (x gen_NSApplication) FinishLaunching() {
	C.NSApplication_inst_finishLaunching(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// HideOtherApplications Hides all apps, except the receiver.
// https://developer.apple.com/documentation/appkit/nsapplication/1428746-hideotherapplications?language=objc
func (x gen_NSApplication) HideOtherApplications(
	sender objc.Ref,
) {
	C.NSApplication_inst_hideOtherApplications(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// PostEvent_atStart Adds a given event to the receiver’s event queue.
// https://developer.apple.com/documentation/appkit/nsapplication/1428710-postevent?language=objc
func (x gen_NSApplication) PostEvent_atStart(
	event NSEventRef,
	flag bool,
) {
	C.NSApplication_inst_postEvent_atStart(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
		convertToObjCBool(flag),
	)

	return

}

// RegisterForRemoteNotifications Register for notifications sent by Apple Push Notification service (APNs).
// https://developer.apple.com/documentation/appkit/nsapplication/2967172-registerforremotenotifications?language=objc
func (x gen_NSApplication) RegisterForRemoteNotifications() {
	C.NSApplication_inst_registerForRemoteNotifications(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// RegisterUserInterfaceItemSearchHandler Register an object that provides help data to your app.
// https://developer.apple.com/documentation/appkit/nsapplication/1420818-registeruserinterfaceitemsearchh?language=objc
func (x gen_NSApplication) RegisterUserInterfaceItemSearchHandler(
	handler objc.Ref,
) {
	C.NSApplication_inst_registerUserInterfaceItemSearchHandler(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(handler),
	)

	return

}

// ReplyToApplicationShouldTerminate Responds to NSTerminateLater once the app knows whether it can terminate.
// https://developer.apple.com/documentation/appkit/nsapplication/1428594-replytoapplicationshouldterminat?language=objc
func (x gen_NSApplication) ReplyToApplicationShouldTerminate(
	shouldTerminate bool,
) {
	C.NSApplication_inst_replyToApplicationShouldTerminate(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(shouldTerminate),
	)

	return

}

// Run Starts the main event loop.
// https://developer.apple.com/documentation/appkit/nsapplication/1428631-run?language=objc
func (x gen_NSApplication) Run() {
	C.NSApplication_inst_run(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// SendAction_to_from Sends the given action message to the given target.
// https://developer.apple.com/documentation/appkit/nsapplication/1428509-sendaction?language=objc
func (x gen_NSApplication) SendAction_to_from(
	action objc.Selector,
	target objc.Ref,
	sender objc.Ref,
) bool {
	ret := C.NSApplication_inst_sendAction_to_from(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(sender),
	)

	return convertObjCBoolToGo(ret)

}

// SendEvent Dispatches an event to other objects.
// https://developer.apple.com/documentation/appkit/nsapplication/1428359-sendevent?language=objc
func (x gen_NSApplication) SendEvent(
	event NSEventRef,
) {
	C.NSApplication_inst_sendEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return

}

// SetActivationPolicy Attempts to modify the app’s activation policy.
// https://developer.apple.com/documentation/appkit/nsapplication/1428621-setactivationpolicy?language=objc
func (x gen_NSApplication) SetActivationPolicy(
	activationPolicy core.NSInteger,
) bool {
	ret := C.NSApplication_inst_setActivationPolicy(
		unsafe.Pointer(x.Pointer()),
		C.long(activationPolicy),
	)

	return convertObjCBoolToGo(ret)

}

// ShowHelp If your project is properly registered, and the necessary keys have been set in the property list, this method launches Help Viewer and displays the first page of your app’s help book.
// https://developer.apple.com/documentation/appkit/nsapplication/1500910-showhelp?language=objc
func (x gen_NSApplication) ShowHelp(
	sender objc.Ref,
) {
	C.NSApplication_inst_showHelp(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// Stop Stops the main event loop.
// https://developer.apple.com/documentation/appkit/nsapplication/1428473-stop?language=objc
func (x gen_NSApplication) Stop(
	sender objc.Ref,
) {
	C.NSApplication_inst_stop(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// TargetForAction Returns the object that receives the action message specified by the given selector.
// https://developer.apple.com/documentation/appkit/nsapplication/1428449-targetforaction?language=objc
func (x gen_NSApplication) TargetForAction(
	action objc.Selector,
) objc.Object {
	ret := C.NSApplication_inst_targetForAction(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
	)

	return objc.Object_fromPointer(ret)

}

// TargetForAction_to_from Searches for an object that can receive the message specified by the given selector.
// https://developer.apple.com/documentation/appkit/nsapplication/1428658-targetforaction?language=objc
func (x gen_NSApplication) TargetForAction_to_from(
	action objc.Selector,
	target objc.Ref,
	sender objc.Ref,
) objc.Object {
	ret := C.NSApplication_inst_targetForAction_to_from(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(sender),
	)

	return objc.Object_fromPointer(ret)

}

// Terminate Terminates the receiver.
// https://developer.apple.com/documentation/appkit/nsapplication/1428417-terminate?language=objc
func (x gen_NSApplication) Terminate(
	sender objc.Ref,
) {
	C.NSApplication_inst_terminate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleTouchBarCustomizationPalette Show or hides the interface for customizing the Touch Bar.
// https://developer.apple.com/documentation/appkit/nsapplication/2646920-toggletouchbarcustomizationpalet?language=objc
func (x gen_NSApplication) ToggleTouchBarCustomizationPalette(
	sender objc.Ref,
) {
	C.NSApplication_inst_toggleTouchBarCustomizationPalette(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// TryToPerform_with Dispatches an action message to the specified target.
// https://developer.apple.com/documentation/appkit/nsapplication/1428366-trytoperform?language=objc
func (x gen_NSApplication) TryToPerform_with(
	action objc.Selector,
	object objc.Ref,
) bool {
	ret := C.NSApplication_inst_tryToPerform_with(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(object),
	)

	return convertObjCBoolToGo(ret)

}

// UnhideAllApplications Unhides all apps, including the receiver.
// https://developer.apple.com/documentation/appkit/nsapplication/1428737-unhideallapplications?language=objc
func (x gen_NSApplication) UnhideAllApplications(
	sender objc.Ref,
) {
	C.NSApplication_inst_unhideAllApplications(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// UnregisterForRemoteNotifications Unregister for notifications received from Apple Push Notification service.
// https://developer.apple.com/documentation/appkit/nsapplication/1428747-unregisterforremotenotifications?language=objc
func (x gen_NSApplication) UnregisterForRemoteNotifications() {
	C.NSApplication_inst_unregisterForRemoteNotifications(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// UnregisterUserInterfaceItemSearchHandler Unregister an object that provides help data to your app.
// https://developer.apple.com/documentation/appkit/nsapplication/1420820-unregisteruserinterfaceitemsearc?language=objc
func (x gen_NSApplication) UnregisterUserInterfaceItemSearchHandler(
	handler objc.Ref,
) {
	C.NSApplication_inst_unregisterUserInterfaceItemSearchHandler(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(handler),
	)

	return

}

// Init_asNSApplication
func (x gen_NSApplication) Init_asNSApplication() NSApplication {
	ret := C.NSApplication_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSApplication_fromPointer(ret)

}

// Delegate The app delegate object.
// https://developer.apple.com/documentation/appkit/nsapplication/1428705-delegate?language=objc
func (x gen_NSApplication) Delegate() objc.Object {
	ret := C.NSApplication_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetDelegate The app delegate object.
// https://developer.apple.com/documentation/appkit/nsapplication/1428705-delegate?language=objc
func (x gen_NSApplication) SetDelegate(
	value objc.Ref,
) {
	C.NSApplication_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// CurrentEvent The last event object that the app retrieved from the event queue.
// https://developer.apple.com/documentation/appkit/nsapplication/1428668-currentevent?language=objc
func (x gen_NSApplication) CurrentEvent() NSEvent {
	ret := C.NSApplication_inst_currentEvent(
		unsafe.Pointer(x.Pointer()),
	)

	return NSEvent_fromPointer(ret)

}

// IsRunning A Boolean value indicating whether the main event loop is running.
// https://developer.apple.com/documentation/appkit/nsapplication/1428759-running?language=objc
func (x gen_NSApplication) IsRunning() bool {
	ret := C.NSApplication_inst_isRunning(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsActive A Boolean value indicating whether this is the active app.
// https://developer.apple.com/documentation/appkit/nsapplication/1428493-active?language=objc
func (x gen_NSApplication) IsActive() bool {
	ret := C.NSApplication_inst_isActive(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsRegisteredForRemoteNotifications A Boolean value indicating whether the app is registered with Apple Push Notification service (APNs).
// https://developer.apple.com/documentation/appkit/nsapplication/2967173-registeredforremotenotifications?language=objc
func (x gen_NSApplication) IsRegisteredForRemoteNotifications() bool {
	ret := C.NSApplication_inst_isRegisteredForRemoteNotifications(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// ApplicationIconImage The image used for the app’s icon.
// https://developer.apple.com/documentation/appkit/nsapplication/1428744-applicationiconimage?language=objc
func (x gen_NSApplication) ApplicationIconImage() NSImage {
	ret := C.NSApplication_inst_applicationIconImage(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_fromPointer(ret)

}

// SetApplicationIconImage The image used for the app’s icon.
// https://developer.apple.com/documentation/appkit/nsapplication/1428744-applicationiconimage?language=objc
func (x gen_NSApplication) SetApplicationIconImage(
	value NSImageRef,
) {
	C.NSApplication_inst_setApplicationIconImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// HelpMenu The help menu used by the app.
// https://developer.apple.com/documentation/appkit/nsapplication/1428644-helpmenu?language=objc
func (x gen_NSApplication) HelpMenu() NSMenu {
	ret := C.NSApplication_inst_helpMenu(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenu_fromPointer(ret)

}

// SetHelpMenu The help menu used by the app.
// https://developer.apple.com/documentation/appkit/nsapplication/1428644-helpmenu?language=objc
func (x gen_NSApplication) SetHelpMenu(
	value NSMenuRef,
) {
	C.NSApplication_inst_setHelpMenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// ServicesProvider The object that provides the services the current app advertises in the Services menu of other apps.
// https://developer.apple.com/documentation/appkit/nsapplication/1428467-servicesprovider?language=objc
func (x gen_NSApplication) ServicesProvider() objc.Object {
	ret := C.NSApplication_inst_servicesProvider(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetServicesProvider The object that provides the services the current app advertises in the Services menu of other apps.
// https://developer.apple.com/documentation/appkit/nsapplication/1428467-servicesprovider?language=objc
func (x gen_NSApplication) SetServicesProvider(
	value objc.Ref,
) {
	C.NSApplication_inst_setServicesProvider(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// IsFullKeyboardAccessEnabled A Boolean value indicating whether Full Keyboard Access is enabled in the Keyboard preference pane.
// https://developer.apple.com/documentation/appkit/nsapplication/1428469-fullkeyboardaccessenabled?language=objc
func (x gen_NSApplication) IsFullKeyboardAccessEnabled() bool {
	ret := C.NSApplication_inst_isFullKeyboardAccessEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// OrderedDocuments An array of document objects arranged according to the front-to-back ordering of their associated windows.
// https://developer.apple.com/documentation/appkit/nsapplication/1494283-ordereddocuments?language=objc
func (x gen_NSApplication) OrderedDocuments() core.NSArray {
	ret := C.NSApplication_inst_orderedDocuments(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// OrderedWindows An array of window objects arranged according to their front-to-back ordering on the screen.
// https://developer.apple.com/documentation/appkit/nsapplication/1494287-orderedwindows?language=objc
func (x gen_NSApplication) OrderedWindows() core.NSArray {
	ret := C.NSApplication_inst_orderedWindows(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// MainMenu The app’s main menu bar.
// https://developer.apple.com/documentation/appkit/nsapplication/1428634-mainmenu?language=objc
func (x gen_NSApplication) MainMenu() NSMenu {
	ret := C.NSApplication_inst_mainMenu(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenu_fromPointer(ret)

}

// SetMainMenu The app’s main menu bar.
// https://developer.apple.com/documentation/appkit/nsapplication/1428634-mainmenu?language=objc
func (x gen_NSApplication) SetMainMenu(
	value NSMenuRef,
) {
	C.NSApplication_inst_setMainMenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

type NSControlRef interface {
	Pointer() uintptr
	Init_asNSControl() NSControl
}

type gen_NSControl struct {
	NSView
}

func NSControl_fromPointer(ptr unsafe.Pointer) NSControl {
	return NSControl{gen_NSControl{
		NSView_fromPointer(ptr),
	}}
}

func NSControl_fromRef(ref objc.Ref) NSControl {
	return NSControl_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// AbortEditing Terminates the current editing operation and discards any edited text.
// https://developer.apple.com/documentation/appkit/nscontrol/1428867-abortediting?language=objc
func (x gen_NSControl) AbortEditing() bool {
	ret := C.NSControl_inst_abortEditing(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// CurrentEditor Returns the current field editor for the control.
// https://developer.apple.com/documentation/appkit/nscontrol/1428980-currenteditor?language=objc
func (x gen_NSControl) CurrentEditor() NSText {
	ret := C.NSControl_inst_currentEditor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSText_fromPointer(ret)

}

// DrawWithExpansionFrame_inView Performs custom expansion tool tip drawing.
// https://developer.apple.com/documentation/appkit/nscontrol/1428895-drawwithexpansionframe?language=objc
func (x gen_NSControl) DrawWithExpansionFrame_inView(
	contentFrame core.NSRect,
	view NSViewRef,
) {
	C.NSControl_inst_drawWithExpansionFrame_inView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentFrame)),
		objc.RefPointer(view),
	)

	return

}

// EditWithFrame_editor_delegate_event Begins editing of the receiver’s text using the specified field editor.
// https://developer.apple.com/documentation/appkit/nscontrol/1428919-editwithframe?language=objc
func (x gen_NSControl) EditWithFrame_editor_delegate_event(
	rect core.NSRect,
	textObj NSTextRef,
	delegate objc.Ref,
	event NSEventRef,
) {
	C.NSControl_inst_editWithFrame_editor_delegate_event(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(textObj),
		objc.RefPointer(delegate),
		objc.RefPointer(event),
	)

	return

}

// EndEditing Ends the editing of text in the receiver using the specified field editor.
// https://developer.apple.com/documentation/appkit/nscontrol/1428936-endediting?language=objc
func (x gen_NSControl) EndEditing(
	textObj NSTextRef,
) {
	C.NSControl_inst_endEditing(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(textObj),
	)

	return

}

// ExpansionFrameWithFrame The frame in which a tool tip can be displayed, if needed.
// https://developer.apple.com/documentation/appkit/nscontrol/1428932-expansionframewithframe?language=objc
func (x gen_NSControl) ExpansionFrameWithFrame(
	contentFrame core.NSRect,
) core.NSRect {
	ret := C.NSControl_inst_expansionFrameWithFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentFrame)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// InitWithFrame_asNSControl Initializes a control with the specified frame rectangle.
// https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func (x gen_NSControl) InitWithFrame_asNSControl(
	frameRect core.NSRect,
) NSControl {
	ret := C.NSControl_inst_initWithFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)

	return NSControl_fromPointer(ret)

}

// MouseDown Informs the receiver that the user has pressed the left mouse button.
// https://developer.apple.com/documentation/appkit/nscontrol/1428918-mousedown?language=objc
func (x gen_NSControl) MouseDown(
	event NSEventRef,
) {
	C.NSControl_inst_mouseDown(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return

}

// PerformClick Simulates a single mouse click on the receiver.
// https://developer.apple.com/documentation/appkit/nscontrol/1428974-performclick?language=objc
func (x gen_NSControl) PerformClick(
	sender objc.Ref,
) {
	C.NSControl_inst_performClick(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// SelectWithFrame_editor_delegate_start_length Selects the specified text range in the receiver's field editor.
// https://developer.apple.com/documentation/appkit/nscontrol/1428968-selectwithframe?language=objc
func (x gen_NSControl) SelectWithFrame_editor_delegate_start_length(
	rect core.NSRect,
	textObj NSTextRef,
	delegate objc.Ref,
	selStart core.NSInteger,
	selLength core.NSInteger,
) {
	C.NSControl_inst_selectWithFrame_editor_delegate_start_length(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(textObj),
		objc.RefPointer(delegate),
		C.long(selStart),
		C.long(selLength),
	)

	return

}

// SendAction_to Causes the specified action to be sent to the target.
// https://developer.apple.com/documentation/appkit/nscontrol/1428851-sendaction?language=objc
func (x gen_NSControl) SendAction_to(
	action objc.Selector,
	target objc.Ref,
) bool {
	ret := C.NSControl_inst_sendAction_to(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(target),
	)

	return convertObjCBoolToGo(ret)

}

// SizeThatFits Asks the control to calculate and return the size that best fits the specified size.
// https://developer.apple.com/documentation/appkit/nscontrol/1428902-sizethatfits?language=objc
func (x gen_NSControl) SizeThatFits(
	size core.NSSize,
) core.NSSize {
	ret := C.NSControl_inst_sizeThatFits(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SizeToFit Resizes the receiver’s frame so that it’s the minimum size needed to contain its cell.
// https://developer.apple.com/documentation/appkit/nscontrol/1428877-sizetofit?language=objc
func (x gen_NSControl) SizeToFit() {
	C.NSControl_inst_sizeToFit(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// TakeDoubleValueFrom Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object.
// https://developer.apple.com/documentation/appkit/nscontrol/1428958-takedoublevaluefrom?language=objc
func (x gen_NSControl) TakeDoubleValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_takeDoubleValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// TakeFloatValueFrom Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object.
// https://developer.apple.com/documentation/appkit/nscontrol/1428938-takefloatvaluefrom?language=objc
func (x gen_NSControl) TakeFloatValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_takeFloatValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// TakeIntValueFrom Sets the value of the receiver’s cell to an integer value obtained from the specified object.
// https://developer.apple.com/documentation/appkit/nscontrol/1428859-takeintvaluefrom?language=objc
func (x gen_NSControl) TakeIntValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_takeIntValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// TakeIntegerValueFrom Sets the value of the receiver’s cell to an NSInteger value obtained from the specified object.
// https://developer.apple.com/documentation/appkit/nscontrol/1428875-takeintegervaluefrom?language=objc
func (x gen_NSControl) TakeIntegerValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_takeIntegerValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// TakeObjectValueFrom Sets the value of the receiver’s cell to the object value obtained from the specified object.
// https://developer.apple.com/documentation/appkit/nscontrol/1428853-takeobjectvaluefrom?language=objc
func (x gen_NSControl) TakeObjectValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_takeObjectValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// TakeStringValueFrom Sets the value of the receiver’s cell to the string value obtained from the specified object.
// https://developer.apple.com/documentation/appkit/nscontrol/1428912-takestringvaluefrom?language=objc
func (x gen_NSControl) TakeStringValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_takeStringValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ValidateEditing Validates changes to any user-typed text.
// https://developer.apple.com/documentation/appkit/nscontrol/1428855-validateediting?language=objc
func (x gen_NSControl) ValidateEditing() {
	C.NSControl_inst_validateEditing(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Init_asNSControl
func (x gen_NSControl) Init_asNSControl() NSControl {
	ret := C.NSControl_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSControl_fromPointer(ret)

}

// IsEnabled A Boolean value that indicates whether the receiver reacts to mouse events.
// https://developer.apple.com/documentation/appkit/nscontrol/1428970-enabled?language=objc
func (x gen_NSControl) IsEnabled() bool {
	ret := C.NSControl_inst_isEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetEnabled A Boolean value that indicates whether the receiver reacts to mouse events.
// https://developer.apple.com/documentation/appkit/nscontrol/1428970-enabled?language=objc
func (x gen_NSControl) SetEnabled(
	value bool,
) {
	C.NSControl_inst_setEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IntValue The value of the receiver’s cell as an integer.
// https://developer.apple.com/documentation/appkit/nscontrol/1428939-intvalue?language=objc
func (x gen_NSControl) IntValue() int32 {
	ret := C.NSControl_inst_intValue(
		unsafe.Pointer(x.Pointer()),
	)

	return int32(ret)

}

// SetIntValue The value of the receiver’s cell as an integer.
// https://developer.apple.com/documentation/appkit/nscontrol/1428939-intvalue?language=objc
func (x gen_NSControl) SetIntValue(
	value int32,
) {
	C.NSControl_inst_setIntValue(
		unsafe.Pointer(x.Pointer()),
		C.int(value),
	)

	return

}

// IntegerValue The value of the receiver’s cell as an NSInteger value.
// https://developer.apple.com/documentation/appkit/nscontrol/1428969-integervalue?language=objc
func (x gen_NSControl) IntegerValue() core.NSInteger {
	ret := C.NSControl_inst_integerValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// SetIntegerValue The value of the receiver’s cell as an NSInteger value.
// https://developer.apple.com/documentation/appkit/nscontrol/1428969-integervalue?language=objc
func (x gen_NSControl) SetIntegerValue(
	value core.NSInteger,
) {
	C.NSControl_inst_setIntegerValue(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return

}

// ObjectValue The value of the receiver’s cell as an Objective-C object.
// https://developer.apple.com/documentation/appkit/nscontrol/1428849-objectvalue?language=objc
func (x gen_NSControl) ObjectValue() objc.Object {
	ret := C.NSControl_inst_objectValue(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetObjectValue The value of the receiver’s cell as an Objective-C object.
// https://developer.apple.com/documentation/appkit/nscontrol/1428849-objectvalue?language=objc
func (x gen_NSControl) SetObjectValue(
	value objc.Ref,
) {
	C.NSControl_inst_setObjectValue(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// StringValue The value of the receiver’s cell as an NSString object.
// https://developer.apple.com/documentation/appkit/nscontrol/1428950-stringvalue?language=objc
func (x gen_NSControl) StringValue() core.NSString {
	ret := C.NSControl_inst_stringValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetStringValue The value of the receiver’s cell as an NSString object.
// https://developer.apple.com/documentation/appkit/nscontrol/1428950-stringvalue?language=objc
func (x gen_NSControl) SetStringValue(
	value core.NSStringRef,
) {
	C.NSControl_inst_setStringValue(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// AttributedStringValue The value of the receiver’s cell as an attributed string.
// https://developer.apple.com/documentation/appkit/nscontrol/1428916-attributedstringvalue?language=objc
func (x gen_NSControl) AttributedStringValue() core.NSAttributedString {
	ret := C.NSControl_inst_attributedStringValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSAttributedString_fromPointer(ret)

}

// SetAttributedStringValue The value of the receiver’s cell as an attributed string.
// https://developer.apple.com/documentation/appkit/nscontrol/1428916-attributedstringvalue?language=objc
func (x gen_NSControl) SetAttributedStringValue(
	value core.NSAttributedStringRef,
) {
	C.NSControl_inst_setAttributedStringValue(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Font The font used to draw text in the receiver’s cell.
// https://developer.apple.com/documentation/appkit/nscontrol/1428914-font?language=objc
func (x gen_NSControl) Font() NSFont {
	ret := C.NSControl_inst_font(
		unsafe.Pointer(x.Pointer()),
	)

	return NSFont_fromPointer(ret)

}

// SetFont The font used to draw text in the receiver’s cell.
// https://developer.apple.com/documentation/appkit/nscontrol/1428914-font?language=objc
func (x gen_NSControl) SetFont(
	value NSFontRef,
) {
	C.NSControl_inst_setFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// UsesSingleLineMode A Boolean value that indicates whether the text in the control’s cell uses single line mode.
// https://developer.apple.com/documentation/appkit/nscontrol/1428929-usessinglelinemode?language=objc
func (x gen_NSControl) UsesSingleLineMode() bool {
	ret := C.NSControl_inst_usesSingleLineMode(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetUsesSingleLineMode A Boolean value that indicates whether the text in the control’s cell uses single line mode.
// https://developer.apple.com/documentation/appkit/nscontrol/1428929-usessinglelinemode?language=objc
func (x gen_NSControl) SetUsesSingleLineMode(
	value bool,
) {
	C.NSControl_inst_setUsesSingleLineMode(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AllowsExpansionToolTips A Boolean value that indicates whether expansion tool tips are shown when the control is hovered over.
// https://developer.apple.com/documentation/appkit/nscontrol/1428962-allowsexpansiontooltips?language=objc
func (x gen_NSControl) AllowsExpansionToolTips() bool {
	ret := C.NSControl_inst_allowsExpansionToolTips(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsExpansionToolTips A Boolean value that indicates whether expansion tool tips are shown when the control is hovered over.
// https://developer.apple.com/documentation/appkit/nscontrol/1428962-allowsexpansiontooltips?language=objc
func (x gen_NSControl) SetAllowsExpansionToolTips(
	value bool,
) {
	C.NSControl_inst_setAllowsExpansionToolTips(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsHighlighted A Boolean value that indicates whether the cell is highlighted.
// https://developer.apple.com/documentation/appkit/nscontrol/1428927-highlighted?language=objc
func (x gen_NSControl) IsHighlighted() bool {
	ret := C.NSControl_inst_isHighlighted(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetHighlighted A Boolean value that indicates whether the cell is highlighted.
// https://developer.apple.com/documentation/appkit/nscontrol/1428927-highlighted?language=objc
func (x gen_NSControl) SetHighlighted(
	value bool,
) {
	C.NSControl_inst_setHighlighted(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// Action The default action-message selector associated with the control.
// https://developer.apple.com/documentation/appkit/nscontrol/1428956-action?language=objc
func (x gen_NSControl) Action() objc.Selector {
	ret := C.NSControl_inst_action(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.SelectorAt(ret)

}

// SetAction The default action-message selector associated with the control.
// https://developer.apple.com/documentation/appkit/nscontrol/1428956-action?language=objc
func (x gen_NSControl) SetAction(
	value objc.Selector,
) {
	C.NSControl_inst_setAction(
		unsafe.Pointer(x.Pointer()),
		value.SelectorAddress(),
	)

	return

}

// Target The target object that receives action messages from the cell.
// https://developer.apple.com/documentation/appkit/nscontrol/1428885-target?language=objc
func (x gen_NSControl) Target() objc.Object {
	ret := C.NSControl_inst_target(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetTarget The target object that receives action messages from the cell.
// https://developer.apple.com/documentation/appkit/nscontrol/1428885-target?language=objc
func (x gen_NSControl) SetTarget(
	value objc.Ref,
) {
	C.NSControl_inst_setTarget(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// IsContinuous A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
// https://developer.apple.com/documentation/appkit/nscontrol/1428952-continuous?language=objc
func (x gen_NSControl) IsContinuous() bool {
	ret := C.NSControl_inst_isContinuous(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetContinuous A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
// https://developer.apple.com/documentation/appkit/nscontrol/1428952-continuous?language=objc
func (x gen_NSControl) SetContinuous(
	value bool,
) {
	C.NSControl_inst_setContinuous(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// Tag The tag identifying the receiver (not the tag of the receiver’s cell).
// https://developer.apple.com/documentation/appkit/nscontrol/1428910-tag?language=objc
func (x gen_NSControl) Tag() core.NSInteger {
	ret := C.NSControl_inst_tag(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// SetTag The tag identifying the receiver (not the tag of the receiver’s cell).
// https://developer.apple.com/documentation/appkit/nscontrol/1428910-tag?language=objc
func (x gen_NSControl) SetTag(
	value core.NSInteger,
) {
	C.NSControl_inst_setTag(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return

}

// RefusesFirstResponder A Boolean value indicating whether the receiver refuses the first responder role.
// https://developer.apple.com/documentation/appkit/nscontrol/1428976-refusesfirstresponder?language=objc
func (x gen_NSControl) RefusesFirstResponder() bool {
	ret := C.NSControl_inst_refusesFirstResponder(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetRefusesFirstResponder A Boolean value indicating whether the receiver refuses the first responder role.
// https://developer.apple.com/documentation/appkit/nscontrol/1428976-refusesfirstresponder?language=objc
func (x gen_NSControl) SetRefusesFirstResponder(
	value bool,
) {
	C.NSControl_inst_setRefusesFirstResponder(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IgnoresMultiClick A Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession.
// https://developer.apple.com/documentation/appkit/nscontrol/1428863-ignoresmulticlick?language=objc
func (x gen_NSControl) IgnoresMultiClick() bool {
	ret := C.NSControl_inst_ignoresMultiClick(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetIgnoresMultiClick A Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession.
// https://developer.apple.com/documentation/appkit/nscontrol/1428863-ignoresmulticlick?language=objc
func (x gen_NSControl) SetIgnoresMultiClick(
	value bool,
) {
	C.NSControl_inst_setIgnoresMultiClick(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

type NSButtonRef interface {
	Pointer() uintptr
	Init_asNSButton() NSButton
}

type gen_NSButton struct {
	NSControl
}

func NSButton_fromPointer(ptr unsafe.Pointer) NSButton {
	return NSButton{gen_NSButton{
		NSControl_fromPointer(ptr),
	}}
}

func NSButton_fromRef(ref objc.Ref) NSButton {
	return NSButton_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// CompressWithPrioritizedCompressionOptions Sets the priority compression options for this button.
// https://developer.apple.com/documentation/appkit/nsbutton/2952060-compresswithprioritizedcompressi?language=objc
func (x gen_NSButton) CompressWithPrioritizedCompressionOptions(
	prioritizedOptions core.NSArrayRef,
) {
	C.NSButton_inst_compressWithPrioritizedCompressionOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(prioritizedOptions),
	)

	return

}

// Highlight Highlights (or unhighlights) the button.
// https://developer.apple.com/documentation/appkit/nsbutton/1534156-highlight?language=objc
func (x gen_NSButton) Highlight(
	flag bool,
) {
	C.NSButton_inst_highlight(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)

	return

}

// MinimumSizeWithPrioritizedCompressionOptions Returns the minimum size of the button by using the compression options.
// https://developer.apple.com/documentation/appkit/nsbutton/2952059-minimumsizewithprioritizedcompre?language=objc
func (x gen_NSButton) MinimumSizeWithPrioritizedCompressionOptions(
	prioritizedOptions core.NSArrayRef,
) core.NSSize {
	ret := C.NSButton_inst_minimumSizeWithPrioritizedCompressionOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(prioritizedOptions),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// PerformKeyEquivalent Checks the button's key equivalent against the specified event and, if they match, simulates the button being clicked.
// https://developer.apple.com/documentation/appkit/nsbutton/1524423-performkeyequivalent?language=objc
func (x gen_NSButton) PerformKeyEquivalent(
	key NSEventRef,
) bool {
	ret := C.NSButton_inst_performKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)

	return convertObjCBoolToGo(ret)

}

// SetNextState Sets the button to its next state.
// https://developer.apple.com/documentation/appkit/nsbutton/1530594-setnextstate?language=objc
func (x gen_NSButton) SetNextState() {
	C.NSButton_inst_setNextState(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Init_asNSButton
func (x gen_NSButton) Init_asNSButton() NSButton {
	ret := C.NSButton_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSButton_fromPointer(ret)

}

// ContentTintColor A tint color to use for the template image and text content.
// https://developer.apple.com/documentation/appkit/nsbutton/3000781-contenttintcolor?language=objc
func (x gen_NSButton) ContentTintColor() NSColor {
	ret := C.NSButton_inst_contentTintColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_fromPointer(ret)

}

// SetContentTintColor A tint color to use for the template image and text content.
// https://developer.apple.com/documentation/appkit/nsbutton/3000781-contenttintcolor?language=objc
func (x gen_NSButton) SetContentTintColor(
	value NSColorRef,
) {
	C.NSButton_inst_setContentTintColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// HasDestructiveAction A Boolean value that defines whether a button’s action has a destructive effect.
// https://developer.apple.com/documentation/appkit/nsbutton/3622469-hasdestructiveaction?language=objc
func (x gen_NSButton) HasDestructiveAction() bool {
	ret := C.NSButton_inst_hasDestructiveAction(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetHasDestructiveAction A Boolean value that defines whether a button’s action has a destructive effect.
// https://developer.apple.com/documentation/appkit/nsbutton/3622469-hasdestructiveaction?language=objc
func (x gen_NSButton) SetHasDestructiveAction(
	value bool,
) {
	C.NSButton_inst_setHasDestructiveAction(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AlternateTitle The title that the button displays when the button is in an on state.
// https://developer.apple.com/documentation/appkit/nsbutton/1529588-alternatetitle?language=objc
func (x gen_NSButton) AlternateTitle() core.NSString {
	ret := C.NSButton_inst_alternateTitle(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetAlternateTitle The title that the button displays when the button is in an on state.
// https://developer.apple.com/documentation/appkit/nsbutton/1529588-alternatetitle?language=objc
func (x gen_NSButton) SetAlternateTitle(
	value core.NSStringRef,
) {
	C.NSButton_inst_setAlternateTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// AttributedTitle The title that the button displays in an off state, as an attributed string.
// https://developer.apple.com/documentation/appkit/nsbutton/1524640-attributedtitle?language=objc
func (x gen_NSButton) AttributedTitle() core.NSAttributedString {
	ret := C.NSButton_inst_attributedTitle(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSAttributedString_fromPointer(ret)

}

// SetAttributedTitle The title that the button displays in an off state, as an attributed string.
// https://developer.apple.com/documentation/appkit/nsbutton/1524640-attributedtitle?language=objc
func (x gen_NSButton) SetAttributedTitle(
	value core.NSAttributedStringRef,
) {
	C.NSButton_inst_setAttributedTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// AttributedAlternateTitle The title that the button displays as an attributed string when the button is in an on state.
// https://developer.apple.com/documentation/appkit/nsbutton/1526723-attributedalternatetitle?language=objc
func (x gen_NSButton) AttributedAlternateTitle() core.NSAttributedString {
	ret := C.NSButton_inst_attributedAlternateTitle(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSAttributedString_fromPointer(ret)

}

// SetAttributedAlternateTitle The title that the button displays as an attributed string when the button is in an on state.
// https://developer.apple.com/documentation/appkit/nsbutton/1526723-attributedalternatetitle?language=objc
func (x gen_NSButton) SetAttributedAlternateTitle(
	value core.NSAttributedStringRef,
) {
	C.NSButton_inst_setAttributedAlternateTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Title The title displayed on the button when it’s in an off state.
// https://developer.apple.com/documentation/appkit/nsbutton/1524430-title?language=objc
func (x gen_NSButton) Title() core.NSString {
	ret := C.NSButton_inst_title(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetTitle The title displayed on the button when it’s in an off state.
// https://developer.apple.com/documentation/appkit/nsbutton/1524430-title?language=objc
func (x gen_NSButton) SetTitle(
	value core.NSStringRef,
) {
	C.NSButton_inst_setTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Sound The sound that plays when the user clicks the button.
// https://developer.apple.com/documentation/appkit/nsbutton/1530910-sound?language=objc
func (x gen_NSButton) Sound() NSSound {
	ret := C.NSButton_inst_sound(
		unsafe.Pointer(x.Pointer()),
	)

	return NSSound_fromPointer(ret)

}

// SetSound The sound that plays when the user clicks the button.
// https://developer.apple.com/documentation/appkit/nsbutton/1530910-sound?language=objc
func (x gen_NSButton) SetSound(
	value NSSoundRef,
) {
	C.NSButton_inst_setSound(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// IsSpringLoaded A Boolean value that indicates whether spring loading is enabled for the button.
// https://developer.apple.com/documentation/appkit/nsbutton/1532300-springloaded?language=objc
func (x gen_NSButton) IsSpringLoaded() bool {
	ret := C.NSButton_inst_isSpringLoaded(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetSpringLoaded A Boolean value that indicates whether spring loading is enabled for the button.
// https://developer.apple.com/documentation/appkit/nsbutton/1532300-springloaded?language=objc
func (x gen_NSButton) SetSpringLoaded(
	value bool,
) {
	C.NSButton_inst_setSpringLoaded(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// MaxAcceleratorLevel An integer value indicating the maximum pressure level for a button of type NSMultiLevelAcceleratorButton.
// https://developer.apple.com/documentation/appkit/nsbutton/1534413-maxacceleratorlevel?language=objc
func (x gen_NSButton) MaxAcceleratorLevel() core.NSInteger {
	ret := C.NSButton_inst_maxAcceleratorLevel(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// SetMaxAcceleratorLevel An integer value indicating the maximum pressure level for a button of type NSMultiLevelAcceleratorButton.
// https://developer.apple.com/documentation/appkit/nsbutton/1534413-maxacceleratorlevel?language=objc
func (x gen_NSButton) SetMaxAcceleratorLevel(
	value core.NSInteger,
) {
	C.NSButton_inst_setMaxAcceleratorLevel(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return

}

// Image The image that appears on the button when it’s in an off state, or nil if there is no such image.
// https://developer.apple.com/documentation/appkit/nsbutton/1534221-image?language=objc
func (x gen_NSButton) Image() NSImage {
	ret := C.NSButton_inst_image(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_fromPointer(ret)

}

// SetImage The image that appears on the button when it’s in an off state, or nil if there is no such image.
// https://developer.apple.com/documentation/appkit/nsbutton/1534221-image?language=objc
func (x gen_NSButton) SetImage(
	value NSImageRef,
) {
	C.NSButton_inst_setImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// AlternateImage An alternate image that appears on the button when the button is in an on state.
// https://developer.apple.com/documentation/appkit/nsbutton/1533935-alternateimage?language=objc
func (x gen_NSButton) AlternateImage() NSImage {
	ret := C.NSButton_inst_alternateImage(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_fromPointer(ret)

}

// SetAlternateImage An alternate image that appears on the button when the button is in an on state.
// https://developer.apple.com/documentation/appkit/nsbutton/1533935-alternateimage?language=objc
func (x gen_NSButton) SetAlternateImage(
	value NSImageRef,
) {
	C.NSButton_inst_setAlternateImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// IsBordered A Boolean value that determines whether the button has a border.
// https://developer.apple.com/documentation/appkit/nsbutton/1525565-bordered?language=objc
func (x gen_NSButton) IsBordered() bool {
	ret := C.NSButton_inst_isBordered(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetBordered A Boolean value that determines whether the button has a border.
// https://developer.apple.com/documentation/appkit/nsbutton/1525565-bordered?language=objc
func (x gen_NSButton) SetBordered(
	value bool,
) {
	C.NSButton_inst_setBordered(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsTransparent A Boolean value that indicates whether the button is transparent.
// https://developer.apple.com/documentation/appkit/nsbutton/1529659-transparent?language=objc
func (x gen_NSButton) IsTransparent() bool {
	ret := C.NSButton_inst_isTransparent(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetTransparent A Boolean value that indicates whether the button is transparent.
// https://developer.apple.com/documentation/appkit/nsbutton/1529659-transparent?language=objc
func (x gen_NSButton) SetTransparent(
	value bool,
) {
	C.NSButton_inst_setTransparent(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// BezelColor The color of the button's bezel, in appearances that support it.
// https://developer.apple.com/documentation/appkit/nsbutton/2561000-bezelcolor?language=objc
func (x gen_NSButton) BezelColor() NSColor {
	ret := C.NSButton_inst_bezelColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_fromPointer(ret)

}

// SetBezelColor The color of the button's bezel, in appearances that support it.
// https://developer.apple.com/documentation/appkit/nsbutton/2561000-bezelcolor?language=objc
func (x gen_NSButton) SetBezelColor(
	value NSColorRef,
) {
	C.NSButton_inst_setBezelColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// ShowsBorderOnlyWhileMouseInside A Boolean value that determines whether the button displays its border only when the pointer is over it.
// https://developer.apple.com/documentation/appkit/nsbutton/1532248-showsborderonlywhilemouseinside?language=objc
func (x gen_NSButton) ShowsBorderOnlyWhileMouseInside() bool {
	ret := C.NSButton_inst_showsBorderOnlyWhileMouseInside(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetShowsBorderOnlyWhileMouseInside A Boolean value that determines whether the button displays its border only when the pointer is over it.
// https://developer.apple.com/documentation/appkit/nsbutton/1532248-showsborderonlywhilemouseinside?language=objc
func (x gen_NSButton) SetShowsBorderOnlyWhileMouseInside(
	value bool,
) {
	C.NSButton_inst_setShowsBorderOnlyWhileMouseInside(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// ImageHugsTitle A Boolean value that determines how the button’s image and title are positioned together within the button bezel.
// https://developer.apple.com/documentation/appkit/nsbutton/2092414-imagehugstitle?language=objc
func (x gen_NSButton) ImageHugsTitle() bool {
	ret := C.NSButton_inst_imageHugsTitle(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetImageHugsTitle A Boolean value that determines how the button’s image and title are positioned together within the button bezel.
// https://developer.apple.com/documentation/appkit/nsbutton/2092414-imagehugstitle?language=objc
func (x gen_NSButton) SetImageHugsTitle(
	value bool,
) {
	C.NSButton_inst_setImageHugsTitle(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AllowsMixedState A Boolean value that indicates whether the button allows a mixed state.
// https://developer.apple.com/documentation/appkit/nsbutton/1528670-allowsmixedstate?language=objc
func (x gen_NSButton) AllowsMixedState() bool {
	ret := C.NSButton_inst_allowsMixedState(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsMixedState A Boolean value that indicates whether the button allows a mixed state.
// https://developer.apple.com/documentation/appkit/nsbutton/1528670-allowsmixedstate?language=objc
func (x gen_NSButton) SetAllowsMixedState(
	value bool,
) {
	C.NSButton_inst_setAllowsMixedState(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// State The button’s state.
// https://developer.apple.com/documentation/appkit/nsbutton/1528907-state?language=objc
func (x gen_NSButton) State() core.NSInteger {
	ret := C.NSButton_inst_state(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// SetState The button’s state.
// https://developer.apple.com/documentation/appkit/nsbutton/1528907-state?language=objc
func (x gen_NSButton) SetState(
	value core.NSInteger,
) {
	C.NSButton_inst_setState(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return

}

// KeyEquivalent The key-equivalent character of the button.
// https://developer.apple.com/documentation/appkit/nsbutton/1525368-keyequivalent?language=objc
func (x gen_NSButton) KeyEquivalent() core.NSString {
	ret := C.NSButton_inst_keyEquivalent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetKeyEquivalent The key-equivalent character of the button.
// https://developer.apple.com/documentation/appkit/nsbutton/1525368-keyequivalent?language=objc
func (x gen_NSButton) SetKeyEquivalent(
	value core.NSStringRef,
) {
	C.NSButton_inst_setKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

type NSEventRef interface {
	Pointer() uintptr
	Init_asNSEvent() NSEvent
}

type gen_NSEvent struct {
	objc.Object
}

func NSEvent_fromPointer(ptr unsafe.Pointer) NSEvent {
	return NSEvent{gen_NSEvent{
		objc.Object_fromPointer(ptr),
	}}
}

func NSEvent_fromRef(ref objc.Ref) NSEvent {
	return NSEvent_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init_asNSEvent
func (x gen_NSEvent) Init_asNSEvent() NSEvent {
	ret := C.NSEvent_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSEvent_fromPointer(ret)

}

// LocationInWindow The receiver’s location in the base coordinate system of the associated window.
// https://developer.apple.com/documentation/appkit/nsevent/1529068-locationinwindow?language=objc
func (x gen_NSEvent) LocationInWindow() core.NSPoint {
	ret := C.NSEvent_inst_locationInWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// Window The window object associated with the event.
// https://developer.apple.com/documentation/appkit/nsevent/1530808-window?language=objc
func (x gen_NSEvent) Window() NSWindow {
	ret := C.NSEvent_inst_window(
		unsafe.Pointer(x.Pointer()),
	)

	return NSWindow_fromPointer(ret)

}

// WindowNumber The identifier for the window device associated with the event.
// https://developer.apple.com/documentation/appkit/nsevent/1531361-windownumber?language=objc
func (x gen_NSEvent) WindowNumber() core.NSInteger {
	ret := C.NSEvent_inst_windowNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// EventRef An opaque Carbon type associated with this event.
// https://developer.apple.com/documentation/appkit/nsevent/1525143-eventref?language=objc
func (x gen_NSEvent) EventRef() unsafe.Pointer {
	ret := C.NSEvent_inst_eventRef(
		unsafe.Pointer(x.Pointer()),
	)

	return ret

}

// Characters The characters associated with a key-up or key-down event.
// https://developer.apple.com/documentation/appkit/nsevent/1534183-characters?language=objc
func (x gen_NSEvent) Characters() core.NSString {
	ret := C.NSEvent_inst_characters(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// CharactersIgnoringModifiers The characters generated by a key event as if no modifier key (except for Shift) applies.
// https://developer.apple.com/documentation/appkit/nsevent/1524605-charactersignoringmodifiers?language=objc
func (x gen_NSEvent) CharactersIgnoringModifiers() core.NSString {
	ret := C.NSEvent_inst_charactersIgnoringModifiers(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// IsARepeat A Boolean value that indicates whether the key event is a repeat.
// https://developer.apple.com/documentation/appkit/nsevent/1528049-arepeat?language=objc
func (x gen_NSEvent) IsARepeat() bool {
	ret := C.NSEvent_inst_isARepeat(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// ButtonNumber The button number for a mouse event.
// https://developer.apple.com/documentation/appkit/nsevent/1527828-buttonnumber?language=objc
func (x gen_NSEvent) ButtonNumber() core.NSInteger {
	ret := C.NSEvent_inst_buttonNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// ClickCount The number of mouse clicks associated with a mouse-down or mouse-up event.
// https://developer.apple.com/documentation/appkit/nsevent/1528200-clickcount?language=objc
func (x gen_NSEvent) ClickCount() core.NSInteger {
	ret := C.NSEvent_inst_clickCount(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// EventNumber The counter value of the latest mouse or tracking-rectangle event object; every system-generated mouse and tracking-rectangle event increments this counter.
// https://developer.apple.com/documentation/appkit/nsevent/1535220-eventnumber?language=objc
func (x gen_NSEvent) EventNumber() core.NSInteger {
	ret := C.NSEvent_inst_eventNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// TrackingNumber The identifier of a mouse-tracking event.
// https://developer.apple.com/documentation/appkit/nsevent/1533974-trackingnumber?language=objc
func (x gen_NSEvent) TrackingNumber() core.NSInteger {
	ret := C.NSEvent_inst_trackingNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// UserData The data associated with a mouse-tracking event.
// https://developer.apple.com/documentation/appkit/nsevent/1526810-userdata?language=objc
func (x gen_NSEvent) UserData() unsafe.Pointer {
	ret := C.NSEvent_inst_userData(
		unsafe.Pointer(x.Pointer()),
	)

	return ret

}

// Data1 Additional data associated with this event.
// https://developer.apple.com/documentation/appkit/nsevent/1528289-data1?language=objc
func (x gen_NSEvent) Data1() core.NSInteger {
	ret := C.NSEvent_inst_data1(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// Data2 Additional data associated with this event.
// https://developer.apple.com/documentation/appkit/nsevent/1528647-data2?language=objc
func (x gen_NSEvent) Data2() core.NSInteger {
	ret := C.NSEvent_inst_data2(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// DeltaX The x-coordinate change for mouse-move, mouse-drag, and swipe events.
// https://developer.apple.com/documentation/appkit/nsevent/1534871-deltax?language=objc
func (x gen_NSEvent) DeltaX() core.CGFloat {
	ret := C.NSEvent_inst_deltaX(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// DeltaY The y-coordinate change for mouse-move, mouse-drag, and swipe events.
// https://developer.apple.com/documentation/appkit/nsevent/1534158-deltay?language=objc
func (x gen_NSEvent) DeltaY() core.CGFloat {
	ret := C.NSEvent_inst_deltaY(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// DeltaZ The z-coordinate change for a scroll wheel, mouse-move, or mouse-drag event.
// https://developer.apple.com/documentation/appkit/nsevent/1531528-deltaz?language=objc
func (x gen_NSEvent) DeltaZ() core.CGFloat {
	ret := C.NSEvent_inst_deltaZ(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// Stage A value of 0, 1, or 2, indicating the stage of a gesture event of type NSEventTypePressure.
// https://developer.apple.com/documentation/appkit/nsevent/1527242-stage?language=objc
func (x gen_NSEvent) Stage() core.NSInteger {
	ret := C.NSEvent_inst_stage(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// StageTransition The transition value for the stage of a pressure gesture event of type NSEventTypePressure.
// https://developer.apple.com/documentation/appkit/nsevent/1526739-stagetransition?language=objc
func (x gen_NSEvent) StageTransition() core.CGFloat {
	ret := C.NSEvent_inst_stageTransition(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// CapabilityMask A mask whose set bits indicate the capabilities of the tablet device that generated this event.
// https://developer.apple.com/documentation/appkit/nsevent/1534648-capabilitymask?language=objc
func (x gen_NSEvent) CapabilityMask() core.NSUInteger {
	ret := C.NSEvent_inst_capabilityMask(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// DeviceID A special identifier that is used to match tablet-pointer and tablet-proximity events.
// https://developer.apple.com/documentation/appkit/nsevent/1530014-deviceid?language=objc
func (x gen_NSEvent) DeviceID() core.NSUInteger {
	ret := C.NSEvent_inst_deviceID(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// IsEnteringProximity A Boolean value that indicates whether a pointing device is entering or leaving the proximity of its tablet.
// https://developer.apple.com/documentation/appkit/nsevent/1531702-enteringproximity?language=objc
func (x gen_NSEvent) IsEnteringProximity() bool {
	ret := C.NSEvent_inst_isEnteringProximity(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// PointingDeviceID The index of the pointing device currently in proximity with the tablet.
// https://developer.apple.com/documentation/appkit/nsevent/1528818-pointingdeviceid?language=objc
func (x gen_NSEvent) PointingDeviceID() core.NSUInteger {
	ret := C.NSEvent_inst_pointingDeviceID(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// PointingDeviceSerialNumber The vendor-assigned serial number of a pointing device.
// https://developer.apple.com/documentation/appkit/nsevent/1533420-pointingdeviceserialnumber?language=objc
func (x gen_NSEvent) PointingDeviceSerialNumber() core.NSUInteger {
	ret := C.NSEvent_inst_pointingDeviceSerialNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// SystemTabletID The index of the tablet device connected to the system.
// https://developer.apple.com/documentation/appkit/nsevent/1528299-systemtabletid?language=objc
func (x gen_NSEvent) SystemTabletID() core.NSUInteger {
	ret := C.NSEvent_inst_systemTabletID(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// TabletID The USB model identifier of the tablet device associated with this event.
// https://developer.apple.com/documentation/appkit/nsevent/1527003-tabletid?language=objc
func (x gen_NSEvent) TabletID() core.NSUInteger {
	ret := C.NSEvent_inst_tabletID(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// VendorID The vendor identifier of the tablet associated with the event.
// https://developer.apple.com/documentation/appkit/nsevent/1525177-vendorid?language=objc
func (x gen_NSEvent) VendorID() core.NSUInteger {
	ret := C.NSEvent_inst_vendorID(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// VendorPointingDeviceType A coded bit field whose set bits indicate the type of pointing device (within a vendor selection) associated with the event.
// https://developer.apple.com/documentation/appkit/nsevent/1527736-vendorpointingdevicetype?language=objc
func (x gen_NSEvent) VendorPointingDeviceType() core.NSUInteger {
	ret := C.NSEvent_inst_vendorPointingDeviceType(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// AbsoluteX The absolute x coordinate of a pointing device on its tablet at full tablet resolution.
// https://developer.apple.com/documentation/appkit/nsevent/1530617-absolutex?language=objc
func (x gen_NSEvent) AbsoluteX() core.NSInteger {
	ret := C.NSEvent_inst_absoluteX(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// AbsoluteY The absolute y coordinate of a pointing device on its tablet at full tablet resolution.
// https://developer.apple.com/documentation/appkit/nsevent/1528904-absolutey?language=objc
func (x gen_NSEvent) AbsoluteY() core.NSInteger {
	ret := C.NSEvent_inst_absoluteY(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// AbsoluteZ The absolute z coordinate of pointing device on its tablet at full tablet resolution.
// https://developer.apple.com/documentation/appkit/nsevent/1532154-absolutez?language=objc
func (x gen_NSEvent) AbsoluteZ() core.NSInteger {
	ret := C.NSEvent_inst_absoluteZ(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// Tilt The scaled tilt values of the pointing device that generated this event.
// https://developer.apple.com/documentation/appkit/nsevent/1534226-tilt?language=objc
func (x gen_NSEvent) Tilt() core.NSPoint {
	ret := C.NSEvent_inst_tilt(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// VendorDefined An array of three vendor-defined NSNumber objects associated with a pointing-type event.
// https://developer.apple.com/documentation/appkit/nsevent/1530551-vendordefined?language=objc
func (x gen_NSEvent) VendorDefined() objc.Object {
	ret := C.NSEvent_inst_vendorDefined(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// Magnification The change in magnification.
// https://developer.apple.com/documentation/appkit/nsevent/1531642-magnification?language=objc
func (x gen_NSEvent) Magnification() core.CGFloat {
	ret := C.NSEvent_inst_magnification(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// HasPreciseScrollingDeltas A Boolean value that indicates whether precise scrolling deltas are available.
// https://developer.apple.com/documentation/appkit/nsevent/1525758-hasprecisescrollingdeltas?language=objc
func (x gen_NSEvent) HasPreciseScrollingDeltas() bool {
	ret := C.NSEvent_inst_hasPreciseScrollingDeltas(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// ScrollingDeltaX The scroll wheel’s horizontal delta.
// https://developer.apple.com/documentation/appkit/nsevent/1524505-scrollingdeltax?language=objc
func (x gen_NSEvent) ScrollingDeltaX() core.CGFloat {
	ret := C.NSEvent_inst_scrollingDeltaX(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// ScrollingDeltaY The scroll wheel’s vertical delta.
// https://developer.apple.com/documentation/appkit/nsevent/1535387-scrollingdeltay?language=objc
func (x gen_NSEvent) ScrollingDeltaY() core.CGFloat {
	ret := C.NSEvent_inst_scrollingDeltaY(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// IsDirectionInvertedFromDevice A Boolean value that indicates whether the user has changed the device inversion.
// https://developer.apple.com/documentation/appkit/nsevent/1525151-directioninvertedfromdevice?language=objc
func (x gen_NSEvent) IsDirectionInvertedFromDevice() bool {
	ret := C.NSEvent_inst_isDirectionInvertedFromDevice(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

type NSFontRef interface {
	Pointer() uintptr
	Init_asNSFont() NSFont
}

type gen_NSFont struct {
	objc.Object
}

func NSFont_fromPointer(ptr unsafe.Pointer) NSFont {
	return NSFont{gen_NSFont{
		objc.Object_fromPointer(ptr),
	}}
}

func NSFont_fromRef(ref objc.Ref) NSFont {
	return NSFont_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// FontWithSize
// https://developer.apple.com/documentation/appkit/nsfont/3667454-fontwithsize?language=objc
func (x gen_NSFont) FontWithSize(
	fontSize core.CGFloat,
) NSFont {
	ret := C.NSFont_inst_fontWithSize(
		unsafe.Pointer(x.Pointer()),
		C.double(fontSize),
	)

	return NSFont_fromPointer(ret)

}

// Set Sets this font as the font for the current graphics context.
// https://developer.apple.com/documentation/appkit/nsfont/1531373-set?language=objc
func (x gen_NSFont) Set() {
	C.NSFont_inst_set(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Init_asNSFont
func (x gen_NSFont) Init_asNSFont() NSFont {
	ret := C.NSFont_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSFont_fromPointer(ret)

}

// PointSize The point size of the font.
// https://developer.apple.com/documentation/appkit/nsfont/1524511-pointsize?language=objc
func (x gen_NSFont) PointSize() core.CGFloat {
	ret := C.NSFont_inst_pointSize(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// IsFixedPitch A Boolean value indicating whether all glyphs in the font have the same advancement.
// https://developer.apple.com/documentation/appkit/nsfont/1529210-fixedpitch?language=objc
func (x gen_NSFont) IsFixedPitch() bool {
	ret := C.NSFont_inst_isFixedPitch(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// MostCompatibleStringEncoding The string encoding that works best with the font.
// https://developer.apple.com/documentation/appkit/nsfont/1527635-mostcompatiblestringencoding?language=objc
func (x gen_NSFont) MostCompatibleStringEncoding() core.NSStringEncoding {
	ret := C.NSFont_inst_mostCompatibleStringEncoding(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSStringEncoding(ret)

}

// NumberOfGlyphs The number of glyphs in the font.
// https://developer.apple.com/documentation/appkit/nsfont/1533968-numberofglyphs?language=objc
func (x gen_NSFont) NumberOfGlyphs() core.NSUInteger {
	ret := C.NSFont_inst_numberOfGlyphs(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// DisplayName The name of the font, including family and face names, to use when displaying the font information to the user.
// https://developer.apple.com/documentation/appkit/nsfont/1531660-displayname?language=objc
func (x gen_NSFont) DisplayName() core.NSString {
	ret := C.NSFont_inst_displayName(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// FamilyName The family name of the font—for example, “Times” or “Helvetica.”
// https://developer.apple.com/documentation/appkit/nsfont/1529585-familyname?language=objc
func (x gen_NSFont) FamilyName() core.NSString {
	ret := C.NSFont_inst_familyName(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// FontName The full name of the font, as used in PostScript language code—for example, “Times-Roman” or “Helvetica-Oblique.”
// https://developer.apple.com/documentation/appkit/nsfont/1526183-fontname?language=objc
func (x gen_NSFont) FontName() core.NSString {
	ret := C.NSFont_inst_fontName(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// IsVertical A Boolean value indicating whether the font is a vertical font.
// https://developer.apple.com/documentation/appkit/nsfont/1534644-vertical?language=objc
func (x gen_NSFont) IsVertical() bool {
	ret := C.NSFont_inst_isVertical(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// VerticalFont A vertical version of the font.
// https://developer.apple.com/documentation/appkit/nsfont/1535152-verticalfont?language=objc
func (x gen_NSFont) VerticalFont() NSFont {
	ret := C.NSFont_inst_verticalFont(
		unsafe.Pointer(x.Pointer()),
	)

	return NSFont_fromPointer(ret)

}

type NSImageRef interface {
	Pointer() uintptr
	Init_asNSImage() NSImage
}

type gen_NSImage struct {
	objc.Object
}

func NSImage_fromPointer(ptr unsafe.Pointer) NSImage {
	return NSImage{gen_NSImage{
		objc.Object_fromPointer(ptr),
	}}
}

func NSImage_fromRef(ref objc.Ref) NSImage {
	return NSImage_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// AddRepresentations Adds an array of image representation objects to the image.
// https://developer.apple.com/documentation/appkit/nsimage/1519964-addrepresentations?language=objc
func (x gen_NSImage) AddRepresentations(
	imageReps core.NSArrayRef,
) {
	C.NSImage_inst_addRepresentations(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(imageReps),
	)

	return

}

// CancelIncrementalLoad Cancels the current download operation, if any, for an incrementally loaded image.
// https://developer.apple.com/documentation/appkit/nsimage/1520041-cancelincrementalload?language=objc
func (x gen_NSImage) CancelIncrementalLoad() {
	C.NSImage_inst_cancelIncrementalLoad(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// DrawInRect Draws the image in the specified rectangle.
// https://developer.apple.com/documentation/appkit/nsimage/1519863-drawinrect?language=objc
func (x gen_NSImage) DrawInRect(
	rect core.NSRect,
) {
	C.NSImage_inst_drawInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return

}

// InitByReferencingFile_asNSImage Initializes and returns an image object using the specified file.
// https://developer.apple.com/documentation/appkit/nsimage/1519955-initbyreferencingfile?language=objc
func (x gen_NSImage) InitByReferencingFile_asNSImage(
	fileName core.NSStringRef,
) NSImage {
	ret := C.NSImage_inst_initByReferencingFile(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fileName),
	)

	return NSImage_fromPointer(ret)

}

// InitByReferencingURL_asNSImage Initializes and returns an image object using the specified URL.
// https://developer.apple.com/documentation/appkit/nsimage/1519990-initbyreferencingurl?language=objc
func (x gen_NSImage) InitByReferencingURL_asNSImage(
	url core.NSURLRef,
) NSImage {
	ret := C.NSImage_inst_initByReferencingURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)

	return NSImage_fromPointer(ret)

}

// InitWithContentsOfFile_asNSImage Initializes and returns an image object with the contents of the specified file.
// https://developer.apple.com/documentation/appkit/nsimage/1519918-initwithcontentsoffile?language=objc
func (x gen_NSImage) InitWithContentsOfFile_asNSImage(
	fileName core.NSStringRef,
) NSImage {
	ret := C.NSImage_inst_initWithContentsOfFile(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fileName),
	)

	return NSImage_fromPointer(ret)

}

// InitWithContentsOfURL_asNSImage Initializes and returns an image object with the contents of the specified URL.
// https://developer.apple.com/documentation/appkit/nsimage/1519907-initwithcontentsofurl?language=objc
func (x gen_NSImage) InitWithContentsOfURL_asNSImage(
	url core.NSURLRef,
) NSImage {
	ret := C.NSImage_inst_initWithContentsOfURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)

	return NSImage_fromPointer(ret)

}

// InitWithData_asNSImage Initializes and returns an image object using the provided image data.
// https://developer.apple.com/documentation/appkit/nsimage/1519941-initwithdata?language=objc
func (x gen_NSImage) InitWithData_asNSImage(
	data core.NSDataRef,
) NSImage {
	ret := C.NSImage_inst_initWithData(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
	)

	return NSImage_fromPointer(ret)

}

// InitWithDataIgnoringOrientation_asNSImage Initializes and returns an image object using the provided image data and ignoring the EXIF orientation tags.
// https://developer.apple.com/documentation/appkit/nsimage/1519915-initwithdataignoringorientation?language=objc
func (x gen_NSImage) InitWithDataIgnoringOrientation_asNSImage(
	data core.NSDataRef,
) NSImage {
	ret := C.NSImage_inst_initWithDataIgnoringOrientation(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
	)

	return NSImage_fromPointer(ret)

}

// InitWithPasteboard_asNSImage Initializes and returns an image object with data from the specified pasteboard.
// https://developer.apple.com/documentation/appkit/nsimage/1519952-initwithpasteboard?language=objc
func (x gen_NSImage) InitWithPasteboard_asNSImage(
	pasteboard NSPasteboardRef,
) NSImage {
	ret := C.NSImage_inst_initWithPasteboard(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pasteboard),
	)

	return NSImage_fromPointer(ret)

}

// InitWithSize_asNSImage Initializes and returns an image object with the specified dimensions.
// https://developer.apple.com/documentation/appkit/nsimage/1520033-initwithsize?language=objc
func (x gen_NSImage) InitWithSize_asNSImage(
	size core.NSSize,
) NSImage {
	ret := C.NSImage_inst_initWithSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return NSImage_fromPointer(ret)

}

// IsTemplate Returns a Boolean value that indicates whether the image is a template image.
// https://developer.apple.com/documentation/appkit/nsimage/1807274-istemplate?language=objc
func (x gen_NSImage) IsTemplate() bool {
	ret := C.NSImage_inst_isTemplate(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// LayerContentsForContentsScale Returns an object that may be used as the contents of a layer.
// https://developer.apple.com/documentation/appkit/nsimage/1519851-layercontentsforcontentsscale?language=objc
func (x gen_NSImage) LayerContentsForContentsScale(
	layerContentsScale core.CGFloat,
) objc.Object {
	ret := C.NSImage_inst_layerContentsForContentsScale(
		unsafe.Pointer(x.Pointer()),
		C.double(layerContentsScale),
	)

	return objc.Object_fromPointer(ret)

}

// LockFocus Prepares the image to receive drawing commands.
// https://developer.apple.com/documentation/appkit/nsimage/1519891-lockfocus?language=objc
func (x gen_NSImage) LockFocus() {
	C.NSImage_inst_lockFocus(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// LockFocusFlipped Prepares the image to receive drawing commands using the specified flipped state.
// https://developer.apple.com/documentation/appkit/nsimage/1519914-lockfocusflipped?language=objc
func (x gen_NSImage) LockFocusFlipped(
	flipped bool,
) {
	C.NSImage_inst_lockFocusFlipped(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flipped),
	)

	return

}

// Recache Invalidates and frees offscreen caches of all image representations.
// https://developer.apple.com/documentation/appkit/nsimage/1519890-recache?language=objc
func (x gen_NSImage) Recache() {
	C.NSImage_inst_recache(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// RecommendedLayerContentsScale Returns the recommended layer contents scale for this image.
// https://developer.apple.com/documentation/appkit/nsimage/1519878-recommendedlayercontentsscale?language=objc
func (x gen_NSImage) RecommendedLayerContentsScale(
	preferredContentsScale core.CGFloat,
) core.CGFloat {
	ret := C.NSImage_inst_recommendedLayerContentsScale(
		unsafe.Pointer(x.Pointer()),
		C.double(preferredContentsScale),
	)

	return core.CGFloat(ret)

}

// UnlockFocus Removes the focus from the image.
// https://developer.apple.com/documentation/appkit/nsimage/1519853-unlockfocus?language=objc
func (x gen_NSImage) UnlockFocus() {
	C.NSImage_inst_unlockFocus(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Init_asNSImage
func (x gen_NSImage) Init_asNSImage() NSImage {
	ret := C.NSImage_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_fromPointer(ret)

}

// Delegate The image’s delegate object.
// https://developer.apple.com/documentation/appkit/nsimage/1519926-delegate?language=objc
func (x gen_NSImage) Delegate() objc.Object {
	ret := C.NSImage_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetDelegate The image’s delegate object.
// https://developer.apple.com/documentation/appkit/nsimage/1519926-delegate?language=objc
func (x gen_NSImage) SetDelegate(
	value objc.Ref,
) {
	C.NSImage_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Size The size of the image.
// https://developer.apple.com/documentation/appkit/nsimage/1519987-size?language=objc
func (x gen_NSImage) Size() core.NSSize {
	ret := C.NSImage_inst_size(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetSize The size of the image.
// https://developer.apple.com/documentation/appkit/nsimage/1519987-size?language=objc
func (x gen_NSImage) SetSize(
	value core.NSSize,
) {
	C.NSImage_inst_setSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// SetTemplate A Boolean value that determines whether the image represents a template image.
// https://developer.apple.com/documentation/appkit/nsimage/1520017-template?language=objc
func (x gen_NSImage) SetTemplate(
	value bool,
) {
	C.NSImage_inst_setTemplate(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// Representations An array containing all of the image object’s image representations.
// https://developer.apple.com/documentation/appkit/nsimage/1519858-representations?language=objc
func (x gen_NSImage) Representations() core.NSArray {
	ret := C.NSImage_inst_representations(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// PrefersColorMatch A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
// https://developer.apple.com/documentation/appkit/nsimage/1520010-preferscolormatch?language=objc
func (x gen_NSImage) PrefersColorMatch() bool {
	ret := C.NSImage_inst_prefersColorMatch(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetPrefersColorMatch A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
// https://developer.apple.com/documentation/appkit/nsimage/1520010-preferscolormatch?language=objc
func (x gen_NSImage) SetPrefersColorMatch(
	value bool,
) {
	C.NSImage_inst_setPrefersColorMatch(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// UsesEPSOnResolutionMismatch A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
// https://developer.apple.com/documentation/appkit/nsimage/1519868-usesepsonresolutionmismatch?language=objc
func (x gen_NSImage) UsesEPSOnResolutionMismatch() bool {
	ret := C.NSImage_inst_usesEPSOnResolutionMismatch(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetUsesEPSOnResolutionMismatch A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
// https://developer.apple.com/documentation/appkit/nsimage/1519868-usesepsonresolutionmismatch?language=objc
func (x gen_NSImage) SetUsesEPSOnResolutionMismatch(
	value bool,
) {
	C.NSImage_inst_setUsesEPSOnResolutionMismatch(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// MatchesOnMultipleResolution A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
// https://developer.apple.com/documentation/appkit/nsimage/1519963-matchesonmultipleresolution?language=objc
func (x gen_NSImage) MatchesOnMultipleResolution() bool {
	ret := C.NSImage_inst_matchesOnMultipleResolution(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetMatchesOnMultipleResolution A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
// https://developer.apple.com/documentation/appkit/nsimage/1519963-matchesonmultipleresolution?language=objc
func (x gen_NSImage) SetMatchesOnMultipleResolution(
	value bool,
) {
	C.NSImage_inst_setMatchesOnMultipleResolution(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsValid A Boolean value that indicates whether it is possible to draw an image representation.
// https://developer.apple.com/documentation/appkit/nsimage/1519991-valid?language=objc
func (x gen_NSImage) IsValid() bool {
	ret := C.NSImage_inst_isValid(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// BackgroundColor The background color for the image.
// https://developer.apple.com/documentation/appkit/nsimage/1520059-backgroundcolor?language=objc
func (x gen_NSImage) BackgroundColor() NSColor {
	ret := C.NSImage_inst_backgroundColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_fromPointer(ret)

}

// SetBackgroundColor The background color for the image.
// https://developer.apple.com/documentation/appkit/nsimage/1520059-backgroundcolor?language=objc
func (x gen_NSImage) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSImage_inst_setBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// AlignmentRect A rectangle that you can use to position the image during layout.
// https://developer.apple.com/documentation/appkit/nsimage/1519905-alignmentrect?language=objc
func (x gen_NSImage) AlignmentRect() core.NSRect {
	ret := C.NSImage_inst_alignmentRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// SetAlignmentRect A rectangle that you can use to position the image during layout.
// https://developer.apple.com/documentation/appkit/nsimage/1519905-alignmentrect?language=objc
func (x gen_NSImage) SetAlignmentRect(
	value core.NSRect,
) {
	C.NSImage_inst_setAlignmentRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return

}

// TIFFRepresentation A data object containing TIFF data for all of the image representations in the image.
// https://developer.apple.com/documentation/appkit/nsimage/1519841-tiffrepresentation?language=objc
func (x gen_NSImage) TIFFRepresentation() core.NSData {
	ret := C.NSImage_inst_TIFFRepresentation(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSData_fromPointer(ret)

}

// AccessibilityDescription The image’s accessibility description.
// https://developer.apple.com/documentation/appkit/nsimage/1519943-accessibilitydescription?language=objc
func (x gen_NSImage) AccessibilityDescription() core.NSString {
	ret := C.NSImage_inst_accessibilityDescription(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetAccessibilityDescription The image’s accessibility description.
// https://developer.apple.com/documentation/appkit/nsimage/1519943-accessibilitydescription?language=objc
func (x gen_NSImage) SetAccessibilityDescription(
	value core.NSStringRef,
) {
	C.NSImage_inst_setAccessibilityDescription(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// MatchesOnlyOnBestFittingAxis A Boolean value that indicates whether the image matches only on the best fitting axis.
// https://developer.apple.com/documentation/appkit/nsimage/1519848-matchesonlyonbestfittingaxis?language=objc
func (x gen_NSImage) MatchesOnlyOnBestFittingAxis() bool {
	ret := C.NSImage_inst_matchesOnlyOnBestFittingAxis(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetMatchesOnlyOnBestFittingAxis A Boolean value that indicates whether the image matches only on the best fitting axis.
// https://developer.apple.com/documentation/appkit/nsimage/1519848-matchesonlyonbestfittingaxis?language=objc
func (x gen_NSImage) SetMatchesOnlyOnBestFittingAxis(
	value bool,
) {
	C.NSImage_inst_setMatchesOnlyOnBestFittingAxis(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

type NSImageViewRef interface {
	Pointer() uintptr
	Init_asNSImageView() NSImageView
}

type gen_NSImageView struct {
	NSControl
}

func NSImageView_fromPointer(ptr unsafe.Pointer) NSImageView {
	return NSImageView{gen_NSImageView{
		NSControl_fromPointer(ptr),
	}}
}

func NSImageView_fromRef(ref objc.Ref) NSImageView {
	return NSImageView_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init_asNSImageView
func (x gen_NSImageView) Init_asNSImageView() NSImageView {
	ret := C.NSImageView_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImageView_fromPointer(ret)

}

// Image The image displayed by the image view.
// https://developer.apple.com/documentation/appkit/nsimageview/1404952-image?language=objc
func (x gen_NSImageView) Image() NSImage {
	ret := C.NSImageView_inst_image(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_fromPointer(ret)

}

// SetImage The image displayed by the image view.
// https://developer.apple.com/documentation/appkit/nsimageview/1404952-image?language=objc
func (x gen_NSImageView) SetImage(
	value NSImageRef,
) {
	C.NSImageView_inst_setImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Animates A Boolean value indicating whether the image view automatically plays animated images.
// https://developer.apple.com/documentation/appkit/nsimageview/1404950-animates?language=objc
func (x gen_NSImageView) Animates() bool {
	ret := C.NSImageView_inst_animates(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAnimates A Boolean value indicating whether the image view automatically plays animated images.
// https://developer.apple.com/documentation/appkit/nsimageview/1404950-animates?language=objc
func (x gen_NSImageView) SetAnimates(
	value bool,
) {
	C.NSImageView_inst_setAnimates(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsEditable A Boolean value indicating whether the user can drag a new image into the image view.
// https://developer.apple.com/documentation/appkit/nsimageview/1404954-editable?language=objc
func (x gen_NSImageView) IsEditable() bool {
	ret := C.NSImageView_inst_isEditable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetEditable A Boolean value indicating whether the user can drag a new image into the image view.
// https://developer.apple.com/documentation/appkit/nsimageview/1404954-editable?language=objc
func (x gen_NSImageView) SetEditable(
	value bool,
) {
	C.NSImageView_inst_setEditable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AllowsCutCopyPaste A Boolean value indicating whether the image view lets the user cut, copy, and paste the image contents.
// https://developer.apple.com/documentation/appkit/nsimageview/1404961-allowscutcopypaste?language=objc
func (x gen_NSImageView) AllowsCutCopyPaste() bool {
	ret := C.NSImageView_inst_allowsCutCopyPaste(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsCutCopyPaste A Boolean value indicating whether the image view lets the user cut, copy, and paste the image contents.
// https://developer.apple.com/documentation/appkit/nsimageview/1404961-allowscutcopypaste?language=objc
func (x gen_NSImageView) SetAllowsCutCopyPaste(
	value bool,
) {
	C.NSImageView_inst_setAllowsCutCopyPaste(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// ContentTintColor
// https://developer.apple.com/documentation/appkit/nsimageview/3000783-contenttintcolor?language=objc
func (x gen_NSImageView) ContentTintColor() NSColor {
	ret := C.NSImageView_inst_contentTintColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_fromPointer(ret)

}

// SetContentTintColor
// https://developer.apple.com/documentation/appkit/nsimageview/3000783-contenttintcolor?language=objc
func (x gen_NSImageView) SetContentTintColor(
	value NSColorRef,
) {
	C.NSImageView_inst_setContentTintColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

type NSNibRef interface {
	Pointer() uintptr
	Init_asNSNib() NSNib
}

type gen_NSNib struct {
	objc.Object
}

func NSNib_fromPointer(ptr unsafe.Pointer) NSNib {
	return NSNib{gen_NSNib{
		objc.Object_fromPointer(ptr),
	}}
}

func NSNib_fromRef(ref objc.Ref) NSNib {
	return NSNib_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// InitWithNibData_bundle_asNSNib Initializes an instance with nib data and specified bundle for locating resources.
// https://developer.apple.com/documentation/appkit/nsnib/1535865-initwithnibdata?language=objc
func (x gen_NSNib) InitWithNibData_bundle_asNSNib(
	nibData core.NSDataRef,
	bundle NSBundleRef,
) NSNib {
	ret := C.NSNib_inst_initWithNibData_bundle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(nibData),
		objc.RefPointer(bundle),
	)

	return NSNib_fromPointer(ret)

}

// InstantiateWithOwner_topLevelObjects Instantiates objects in the nib file with the specified owner.
// https://developer.apple.com/documentation/appkit/nsnib/1527173-instantiatewithowner?language=objc
func (x gen_NSNib) InstantiateWithOwner_topLevelObjects(
	owner objc.Ref,
	topLevelObjects core.NSArrayRef,
) bool {
	ret := C.NSNib_inst_instantiateWithOwner_topLevelObjects(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(owner),
		objc.RefPointer(topLevelObjects),
	)

	return convertObjCBoolToGo(ret)

}

// Init_asNSNib
func (x gen_NSNib) Init_asNSNib() NSNib {
	ret := C.NSNib_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSNib_fromPointer(ret)

}

type NSPasteboardRef interface {
	Pointer() uintptr
	Init_asNSPasteboard() NSPasteboard
}

type gen_NSPasteboard struct {
	objc.Object
}

func NSPasteboard_fromPointer(ptr unsafe.Pointer) NSPasteboard {
	return NSPasteboard{gen_NSPasteboard{
		objc.Object_fromPointer(ptr),
	}}
}

func NSPasteboard_fromRef(ref objc.Ref) NSPasteboard {
	return NSPasteboard_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// AddTypes_owner Adds promises for the specified types to the first pasteboard item.
// https://developer.apple.com/documentation/appkit/nspasteboard/1533580-addtypes?language=objc
func (x gen_NSPasteboard) AddTypes_owner(
	newTypes core.NSArrayRef,
	newOwner objc.Ref,
) core.NSInteger {
	ret := C.NSPasteboard_inst_addTypes_owner(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newTypes),
		objc.RefPointer(newOwner),
	)

	return core.NSInteger(ret)

}

// CanReadItemWithDataConformingToTypes Returns a Boolean value that indicates whether the receiver contains any items that conform to the specified UTIs.
// https://developer.apple.com/documentation/appkit/nspasteboard/1533576-canreaditemwithdataconformingtot?language=objc
func (x gen_NSPasteboard) CanReadItemWithDataConformingToTypes(
	types core.NSArrayRef,
) bool {
	ret := C.NSPasteboard_inst_canReadItemWithDataConformingToTypes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(types),
	)

	return convertObjCBoolToGo(ret)

}

// CanReadObjectForClasses_options Returns a Boolean value that indicates whether the receiver contains any items that can be represented as an instance of any class in a given array.
// https://developer.apple.com/documentation/appkit/nspasteboard/1533360-canreadobjectforclasses?language=objc
func (x gen_NSPasteboard) CanReadObjectForClasses_options(
	classArray core.NSArrayRef,
	options core.NSDictionaryRef,
) bool {
	ret := C.NSPasteboard_inst_canReadObjectForClasses_options(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(classArray),
		objc.RefPointer(options),
	)

	return convertObjCBoolToGo(ret)

}

// ClearContents Clears the existing contents of the pasteboard.
// https://developer.apple.com/documentation/appkit/nspasteboard/1533599-clearcontents?language=objc
func (x gen_NSPasteboard) ClearContents() core.NSInteger {
	ret := C.NSPasteboard_inst_clearContents(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// DeclareTypes_owner Prepares the receiver for a change in its contents by declaring the new types of data it will contain and a new owner.
// https://developer.apple.com/documentation/appkit/nspasteboard/1533561-declaretypes?language=objc
func (x gen_NSPasteboard) DeclareTypes_owner(
	newTypes core.NSArrayRef,
	newOwner objc.Ref,
) core.NSInteger {
	ret := C.NSPasteboard_inst_declareTypes_owner(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newTypes),
		objc.RefPointer(newOwner),
	)

	return core.NSInteger(ret)

}

// ReadObjectsForClasses_options Reads from the receiver objects that best match the specified array of classes.
// https://developer.apple.com/documentation/appkit/nspasteboard/1524454-readobjectsforclasses?language=objc
func (x gen_NSPasteboard) ReadObjectsForClasses_options(
	classArray core.NSArrayRef,
	options core.NSDictionaryRef,
) core.NSArray {
	ret := C.NSPasteboard_inst_readObjectsForClasses_options(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(classArray),
		objc.RefPointer(options),
	)

	return core.NSArray_fromPointer(ret)

}

// ReleaseGlobally Releases the receiver’s resources in the pasteboard server.
// https://developer.apple.com/documentation/appkit/nspasteboard/1527044-releaseglobally?language=objc
func (x gen_NSPasteboard) ReleaseGlobally() {
	C.NSPasteboard_inst_releaseGlobally(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// WriteFileContents Writes the contents of the specified file to the pasteboard.
// https://developer.apple.com/documentation/appkit/nspasteboard/1531224-writefilecontents?language=objc
func (x gen_NSPasteboard) WriteFileContents(
	filename core.NSStringRef,
) bool {
	ret := C.NSPasteboard_inst_writeFileContents(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(filename),
	)

	return convertObjCBoolToGo(ret)

}

// WriteObjects Writes an array of objects to the receiver.
// https://developer.apple.com/documentation/appkit/nspasteboard/1525945-writeobjects?language=objc
func (x gen_NSPasteboard) WriteObjects(
	objects core.NSArrayRef,
) bool {
	ret := C.NSPasteboard_inst_writeObjects(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(objects),
	)

	return convertObjCBoolToGo(ret)

}

// Init_asNSPasteboard
func (x gen_NSPasteboard) Init_asNSPasteboard() NSPasteboard {
	ret := C.NSPasteboard_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSPasteboard_fromPointer(ret)

}

// PasteboardItems An array that contains all the items held by the pasteboard.
// https://developer.apple.com/documentation/appkit/nspasteboard/1529995-pasteboarditems?language=objc
func (x gen_NSPasteboard) PasteboardItems() core.NSArray {
	ret := C.NSPasteboard_inst_pasteboardItems(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// Types An array of the receiver’s supported data types.
// https://developer.apple.com/documentation/appkit/nspasteboard/1529599-types?language=objc
func (x gen_NSPasteboard) Types() core.NSArray {
	ret := C.NSPasteboard_inst_types(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// ChangeCount The receiver’s change count.
// https://developer.apple.com/documentation/appkit/nspasteboard/1533544-changecount?language=objc
func (x gen_NSPasteboard) ChangeCount() core.NSInteger {
	ret := C.NSPasteboard_inst_changeCount(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

type NSLayoutManagerRef interface {
	Pointer() uintptr
	Init_asNSLayoutManager() NSLayoutManager
}

type gen_NSLayoutManager struct {
	objc.Object
}

func NSLayoutManager_fromPointer(ptr unsafe.Pointer) NSLayoutManager {
	return NSLayoutManager{gen_NSLayoutManager{
		objc.Object_fromPointer(ptr),
	}}
}

func NSLayoutManager_fromRef(ref objc.Ref) NSLayoutManager {
	return NSLayoutManager_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// AddTextContainer Appends the specified text container to the series of text containers where the layout manager arranges text.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1402946-addtextcontainer?language=objc
func (x gen_NSLayoutManager) AddTextContainer(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_addTextContainer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)

	return

}

// AttachmentSizeForGlyphAtIndex Returns the size of the attachment glyph at the specified index.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403099-attachmentsizeforglyphatindex?language=objc
func (x gen_NSLayoutManager) AttachmentSizeForGlyphAtIndex(
	glyphIndex core.NSUInteger,
) core.NSSize {
	ret := C.NSLayoutManager_inst_attachmentSizeForGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// CharacterIndexForGlyphAtIndex Returns the index in the text storage for the first character of the specified glyph.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1402944-characterindexforglyphatindex?language=objc
func (x gen_NSLayoutManager) CharacterIndexForGlyphAtIndex(
	glyphIndex core.NSUInteger,
) core.NSUInteger {
	ret := C.NSLayoutManager_inst_characterIndexForGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)

	return core.NSUInteger(ret)

}

// DefaultBaselineOffsetForFont Returns the default baseline offset that the layout manager's typesetter uses for the specified font.
// https://developer.apple.com/documentation/appkit/nslayoutmanager/1403058-defaultbaselineoffsetforfont?language=objc
func (x gen_NSLayoutManager) DefaultBaselineOffsetForFont(
	theFont NSFontRef,
) core.CGFloat {
	ret := C.NSLayoutManager_inst_defaultBaselineOffsetForFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(theFont),
	)

	return core.CGFloat(ret)

}

// DefaultLineHeightForFont Returns the default line height for a line of text that uses a specified font.
// https://developer.apple.com/documentation/appkit/nslayoutmanager/1403007-defaultlineheightforfont?language=objc
func (x gen_NSLayoutManager) DefaultLineHeightForFont(
	theFont NSFontRef,
) core.CGFloat {
	ret := C.NSLayoutManager_inst_defaultLineHeightForFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(theFont),
	)

	return core.CGFloat(ret)

}

// DrawsOutsideLineFragmentForGlyphAtIndex Indicates whether the glyph draws outside its line fragment rectangle.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403003-drawsoutsidelinefragmentforglyph?language=objc
func (x gen_NSLayoutManager) DrawsOutsideLineFragmentForGlyphAtIndex(
	glyphIndex core.NSUInteger,
) bool {
	ret := C.NSLayoutManager_inst_drawsOutsideLineFragmentForGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)

	return convertObjCBoolToGo(ret)

}

// EnsureLayoutForBoundingRect_inTextContainer Forces the layout manager to perform layout for the specified area in the specified text container if it hasn’t already.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1402962-ensurelayoutforboundingrect?language=objc
func (x gen_NSLayoutManager) EnsureLayoutForBoundingRect_inTextContainer(
	bounds core.NSRect,
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_ensureLayoutForBoundingRect_inTextContainer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&bounds)),
		objc.RefPointer(container),
	)

	return

}

// EnsureLayoutForTextContainer Forces the layout manager to perform layout for the specified text container if it hasn’t already.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1402967-ensurelayoutfortextcontainer?language=objc
func (x gen_NSLayoutManager) EnsureLayoutForTextContainer(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_ensureLayoutForTextContainer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)

	return

}

// FirstUnlaidCharacterIndex Returns the index for the first character in the layout manager that isn’t in the layout.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403067-firstunlaidcharacterindex?language=objc
func (x gen_NSLayoutManager) FirstUnlaidCharacterIndex() core.NSUInteger {
	ret := C.NSLayoutManager_inst_firstUnlaidCharacterIndex(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// FirstUnlaidGlyphIndex Returns the index for the first glyph in the layout manager that isn’t in the layout.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403245-firstunlaidglyphindex?language=objc
func (x gen_NSLayoutManager) FirstUnlaidGlyphIndex() core.NSUInteger {
	ret := C.NSLayoutManager_inst_firstUnlaidGlyphIndex(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// GlyphIndexForCharacterAtIndex Returns the index of the first glyph of the character at the specified index.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403001-glyphindexforcharacteratindex?language=objc
func (x gen_NSLayoutManager) GlyphIndexForCharacterAtIndex(
	charIndex core.NSUInteger,
) core.NSUInteger {
	ret := C.NSLayoutManager_inst_glyphIndexForCharacterAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(charIndex),
	)

	return core.NSUInteger(ret)

}

// Init_asNSLayoutManager Initializes a newly created layout manager object.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1402975-init?language=objc
func (x gen_NSLayoutManager) Init_asNSLayoutManager() NSLayoutManager {
	ret := C.NSLayoutManager_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSLayoutManager_fromPointer(ret)

}

// InsertTextContainer_atIndex Inserts a text container at the specified index in the list of text containers.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403010-inserttextcontainer?language=objc
func (x gen_NSLayoutManager) InsertTextContainer_atIndex(
	container NSTextContainerRef,
	index core.NSUInteger,
) {
	C.NSLayoutManager_inst_insertTextContainer_atIndex(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
		C.ulong(index),
	)

	return

}

// IsValidGlyphIndex Indicates whether the specified index refers to a valid glyph.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1402950-isvalidglyphindex?language=objc
func (x gen_NSLayoutManager) IsValidGlyphIndex(
	glyphIndex core.NSUInteger,
) bool {
	ret := C.NSLayoutManager_inst_isValidGlyphIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)

	return convertObjCBoolToGo(ret)

}

// LayoutManagerOwnsFirstResponderInWindow Indicates whether the first responder in the specified window is a text view for the layout manager.
// https://developer.apple.com/documentation/appkit/nslayoutmanager/1403026-layoutmanagerownsfirstresponderi?language=objc
func (x gen_NSLayoutManager) LayoutManagerOwnsFirstResponderInWindow(
	window NSWindowRef,
) bool {
	ret := C.NSLayoutManager_inst_layoutManagerOwnsFirstResponderInWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(window),
	)

	return convertObjCBoolToGo(ret)

}

// NotShownAttributeForGlyphAtIndex Indicates whether the glyph at the specified index has a visible representation.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1402931-notshownattributeforglyphatindex?language=objc
func (x gen_NSLayoutManager) NotShownAttributeForGlyphAtIndex(
	glyphIndex core.NSUInteger,
) bool {
	ret := C.NSLayoutManager_inst_notShownAttributeForGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)

	return convertObjCBoolToGo(ret)

}

// RemoveTextContainerAtIndex Removes the text container at the specified index and invalidates the layout as necessary.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403017-removetextcontaineratindex?language=objc
func (x gen_NSLayoutManager) RemoveTextContainerAtIndex(
	index core.NSUInteger,
) {
	C.NSLayoutManager_inst_removeTextContainerAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(index),
	)

	return

}

// SetDrawsOutsideLineFragment_forGlyphAtIndex Indicates whether the specified glyph exceeds the bounds of the line fragment for its layout.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1402964-setdrawsoutsidelinefragment?language=objc
func (x gen_NSLayoutManager) SetDrawsOutsideLineFragment_forGlyphAtIndex(
	flag bool,
	glyphIndex core.NSUInteger,
) {
	C.NSLayoutManager_inst_setDrawsOutsideLineFragment_forGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
		C.ulong(glyphIndex),
	)

	return

}

// SetExtraLineFragmentRect_usedRect_textContainer Sets the bounds and container for the extra line fragment.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403071-setextralinefragmentrect?language=objc
func (x gen_NSLayoutManager) SetExtraLineFragmentRect_usedRect_textContainer(
	fragmentRect core.NSRect,
	usedRect core.NSRect,
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_setExtraLineFragmentRect_usedRect_textContainer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&fragmentRect)),
		*(*C.NSRect)(unsafe.Pointer(&usedRect)),
		objc.RefPointer(container),
	)

	return

}

// SetNotShownAttribute_forGlyphAtIndex Sets the visibility of the glyph at the specified index.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403078-setnotshownattribute?language=objc
func (x gen_NSLayoutManager) SetNotShownAttribute_forGlyphAtIndex(
	flag bool,
	glyphIndex core.NSUInteger,
) {
	C.NSLayoutManager_inst_setNotShownAttribute_forGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
		C.ulong(glyphIndex),
	)

	return

}

// TextContainerChangedGeometry Invalidates the layout information, and possibly glyphs, for the specified text container and all subsequent text container objects.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403091-textcontainerchangedgeometry?language=objc
func (x gen_NSLayoutManager) TextContainerChangedGeometry(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_textContainerChangedGeometry(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)

	return

}

// TextContainerChangedTextView Updates the information necessary to manage text view objects for the specified text container.
// https://developer.apple.com/documentation/appkit/nslayoutmanager/1403229-textcontainerchangedtextview?language=objc
func (x gen_NSLayoutManager) TextContainerChangedTextView(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_textContainerChangedTextView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)

	return

}

// UsedRectForTextContainer Returns the bounding rectangle for the glyphs in the specified text container.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1402980-usedrectfortextcontainer?language=objc
func (x gen_NSLayoutManager) UsedRectForTextContainer(
	container NSTextContainerRef,
) core.NSRect {
	ret := C.NSLayoutManager_inst_usedRectForTextContainer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// Delegate The layout manager’s delegate.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1402920-delegate?language=objc
func (x gen_NSLayoutManager) Delegate() objc.Object {
	ret := C.NSLayoutManager_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetDelegate The layout manager’s delegate.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1402920-delegate?language=objc
func (x gen_NSLayoutManager) SetDelegate(
	value objc.Ref,
) {
	C.NSLayoutManager_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// AllowsNonContiguousLayout A Boolean value that indicates whether the layout manager allows noncontiguous layout.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403197-allowsnoncontiguouslayout?language=objc
func (x gen_NSLayoutManager) AllowsNonContiguousLayout() bool {
	ret := C.NSLayoutManager_inst_allowsNonContiguousLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsNonContiguousLayout A Boolean value that indicates whether the layout manager allows noncontiguous layout.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403197-allowsnoncontiguouslayout?language=objc
func (x gen_NSLayoutManager) SetAllowsNonContiguousLayout(
	value bool,
) {
	C.NSLayoutManager_inst_setAllowsNonContiguousLayout(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// HasNonContiguousLayout A Boolean value that indicates whether the layout manager currently has any areas of noncontiguous layout.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403207-hasnoncontiguouslayout?language=objc
func (x gen_NSLayoutManager) HasNonContiguousLayout() bool {
	ret := C.NSLayoutManager_inst_hasNonContiguousLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// ShowsInvisibleCharacters A Boolean value that indicates whether to substitute visible glyphs for whitespace and other typically invisible characters.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403254-showsinvisiblecharacters?language=objc
func (x gen_NSLayoutManager) ShowsInvisibleCharacters() bool {
	ret := C.NSLayoutManager_inst_showsInvisibleCharacters(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetShowsInvisibleCharacters A Boolean value that indicates whether to substitute visible glyphs for whitespace and other typically invisible characters.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403254-showsinvisiblecharacters?language=objc
func (x gen_NSLayoutManager) SetShowsInvisibleCharacters(
	value bool,
) {
	C.NSLayoutManager_inst_setShowsInvisibleCharacters(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// ShowsControlCharacters A Boolean value that indicates whether the layout manager substitutes visible glyphs for control characters in the layout.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1402912-showscontrolcharacters?language=objc
func (x gen_NSLayoutManager) ShowsControlCharacters() bool {
	ret := C.NSLayoutManager_inst_showsControlCharacters(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetShowsControlCharacters A Boolean value that indicates whether the layout manager substitutes visible glyphs for control characters in the layout.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1402912-showscontrolcharacters?language=objc
func (x gen_NSLayoutManager) SetShowsControlCharacters(
	value bool,
) {
	C.NSLayoutManager_inst_setShowsControlCharacters(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// UsesFontLeading A Boolean value that indicates whether the layout manager uses the leading of the font.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403156-usesfontleading?language=objc
func (x gen_NSLayoutManager) UsesFontLeading() bool {
	ret := C.NSLayoutManager_inst_usesFontLeading(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetUsesFontLeading A Boolean value that indicates whether the layout manager uses the leading of the font.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403156-usesfontleading?language=objc
func (x gen_NSLayoutManager) SetUsesFontLeading(
	value bool,
) {
	C.NSLayoutManager_inst_setUsesFontLeading(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// BackgroundLayoutEnabled A Boolean value that indicates whether the layout manager generates glyphs and lays them out when the app's run loop is idle.
// https://developer.apple.com/documentation/appkit/nslayoutmanager/1402952-backgroundlayoutenabled?language=objc
func (x gen_NSLayoutManager) BackgroundLayoutEnabled() bool {
	ret := C.NSLayoutManager_inst_backgroundLayoutEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetBackgroundLayoutEnabled A Boolean value that indicates whether the layout manager generates glyphs and lays them out when the app's run loop is idle.
// https://developer.apple.com/documentation/appkit/nslayoutmanager/1402952-backgroundlayoutenabled?language=objc
func (x gen_NSLayoutManager) SetBackgroundLayoutEnabled(
	value bool,
) {
	C.NSLayoutManager_inst_setBackgroundLayoutEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// LimitsLayoutForSuspiciousContents A Boolean value that indicates whether the layout manager avoids laying out unusually long or suspicious input.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/3021179-limitslayoutforsuspiciouscontent?language=objc
func (x gen_NSLayoutManager) LimitsLayoutForSuspiciousContents() bool {
	ret := C.NSLayoutManager_inst_limitsLayoutForSuspiciousContents(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetLimitsLayoutForSuspiciousContents A Boolean value that indicates whether the layout manager avoids laying out unusually long or suspicious input.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/3021179-limitslayoutforsuspiciouscontent?language=objc
func (x gen_NSLayoutManager) SetLimitsLayoutForSuspiciousContents(
	value bool,
) {
	C.NSLayoutManager_inst_setLimitsLayoutForSuspiciousContents(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// UsesDefaultHyphenation A Boolean value that indicates whether the layout manager uses the default hyphenation rules to wrap lines.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/3180380-usesdefaulthyphenation?language=objc
func (x gen_NSLayoutManager) UsesDefaultHyphenation() bool {
	ret := C.NSLayoutManager_inst_usesDefaultHyphenation(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetUsesDefaultHyphenation A Boolean value that indicates whether the layout manager uses the default hyphenation rules to wrap lines.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/3180380-usesdefaulthyphenation?language=objc
func (x gen_NSLayoutManager) SetUsesDefaultHyphenation(
	value bool,
) {
	C.NSLayoutManager_inst_setUsesDefaultHyphenation(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// TextContainers The current text containers of the layout manager.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403144-textcontainers?language=objc
func (x gen_NSLayoutManager) TextContainers() core.NSArray {
	ret := C.NSLayoutManager_inst_textContainers(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// NumberOfGlyphs The number of glyphs in the layout manager.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1402937-numberofglyphs?language=objc
func (x gen_NSLayoutManager) NumberOfGlyphs() core.NSUInteger {
	ret := C.NSLayoutManager_inst_numberOfGlyphs(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// ExtraLineFragmentRect The rectangle for the extra line fragment at the end of a document.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403175-extralinefragmentrect?language=objc
func (x gen_NSLayoutManager) ExtraLineFragmentRect() core.NSRect {
	ret := C.NSLayoutManager_inst_extraLineFragmentRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// ExtraLineFragmentTextContainer The text container for the extra line fragment rectangle.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1403165-extralinefragmenttextcontainer?language=objc
func (x gen_NSLayoutManager) ExtraLineFragmentTextContainer() NSTextContainer {
	ret := C.NSLayoutManager_inst_extraLineFragmentTextContainer(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextContainer_fromPointer(ret)

}

// ExtraLineFragmentUsedRect The rectangle that encloses the insertion point in the extra line fragment rectangle.
// https://developer.apple.com/documentation/uikit/nslayoutmanager/1402988-extralinefragmentusedrect?language=objc
func (x gen_NSLayoutManager) ExtraLineFragmentUsedRect() core.NSRect {
	ret := C.NSLayoutManager_inst_extraLineFragmentUsedRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// FirstTextView The first text view in the layout manager’s series of text views.
// https://developer.apple.com/documentation/appkit/nslayoutmanager/1402995-firsttextview?language=objc
func (x gen_NSLayoutManager) FirstTextView() NSTextView {
	ret := C.NSLayoutManager_inst_firstTextView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextView_fromPointer(ret)

}

// TextViewForBeginningOfSelection The text view that contains the first glyph in the selection.
// https://developer.apple.com/documentation/appkit/nslayoutmanager/1403089-textviewforbeginningofselection?language=objc
func (x gen_NSLayoutManager) TextViewForBeginningOfSelection() NSTextView {
	ret := C.NSLayoutManager_inst_textViewForBeginningOfSelection(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextView_fromPointer(ret)

}

type NSMenuRef interface {
	Pointer() uintptr
	Init_asNSMenu() NSMenu
}

type gen_NSMenu struct {
	objc.Object
}

func NSMenu_fromPointer(ptr unsafe.Pointer) NSMenu {
	return NSMenu{gen_NSMenu{
		objc.Object_fromPointer(ptr),
	}}
}

func NSMenu_fromRef(ref objc.Ref) NSMenu {
	return NSMenu_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// AddItem Adds a menu item to the end of the menu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518176-additem?language=objc
func (x gen_NSMenu) AddItem(
	newItem NSMenuItemRef,
) {
	C.NSMenu_inst_addItem(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newItem),
	)

	return

}

// AddItemWithTitle_action_keyEquivalent Creates a new menu item and adds it to the end of the menu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518181-additemwithtitle?language=objc
func (x gen_NSMenu) AddItemWithTitle_action_keyEquivalent(
	string core.NSStringRef,
	selector objc.Selector,
	charCode core.NSStringRef,
) NSMenuItem {
	ret := C.NSMenu_inst_addItemWithTitle_action_keyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		selector.SelectorAddress(),
		objc.RefPointer(charCode),
	)

	return NSMenuItem_fromPointer(ret)

}

// CancelTracking Dismisses the menu and ends all menu tracking.
// https://developer.apple.com/documentation/appkit/nsmenu/1518150-canceltracking?language=objc
func (x gen_NSMenu) CancelTracking() {
	C.NSMenu_inst_cancelTracking(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// CancelTrackingWithoutAnimation Dismisses the menu and ends all menu tracking without displaying the associated animation.
// https://developer.apple.com/documentation/appkit/nsmenu/1518244-canceltrackingwithoutanimation?language=objc
func (x gen_NSMenu) CancelTrackingWithoutAnimation() {
	C.NSMenu_inst_cancelTrackingWithoutAnimation(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// IndexOfItem Returns the index identifying the location of a specified menu item in the menu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518178-indexofitem?language=objc
func (x gen_NSMenu) IndexOfItem(
	item NSMenuItemRef,
) core.NSInteger {
	ret := C.NSMenu_inst_indexOfItem(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
	)

	return core.NSInteger(ret)

}

// IndexOfItemWithRepresentedObject Returns the index of the first menu item in the menu that has a given represented object.
// https://developer.apple.com/documentation/appkit/nsmenu/1518175-indexofitemwithrepresentedobject?language=objc
func (x gen_NSMenu) IndexOfItemWithRepresentedObject(
	object objc.Ref,
) core.NSInteger {
	ret := C.NSMenu_inst_indexOfItemWithRepresentedObject(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)

	return core.NSInteger(ret)

}

// IndexOfItemWithSubmenu Returns the index of the menu item in the menu with the given submenu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518216-indexofitemwithsubmenu?language=objc
func (x gen_NSMenu) IndexOfItemWithSubmenu(
	submenu NSMenuRef,
) core.NSInteger {
	ret := C.NSMenu_inst_indexOfItemWithSubmenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(submenu),
	)

	return core.NSInteger(ret)

}

// IndexOfItemWithTag Returns the index of the first menu item in the menu identified by a tag.
// https://developer.apple.com/documentation/appkit/nsmenu/1518164-indexofitemwithtag?language=objc
func (x gen_NSMenu) IndexOfItemWithTag(
	tag core.NSInteger,
) core.NSInteger {
	ret := C.NSMenu_inst_indexOfItemWithTag(
		unsafe.Pointer(x.Pointer()),
		C.long(tag),
	)

	return core.NSInteger(ret)

}

// IndexOfItemWithTarget_andAction Returns the index of the first menu item in the menu that has a specified action and target.
// https://developer.apple.com/documentation/appkit/nsmenu/1518153-indexofitemwithtarget?language=objc
func (x gen_NSMenu) IndexOfItemWithTarget_andAction(
	target objc.Ref,
	actionSelector objc.Selector,
) core.NSInteger {
	ret := C.NSMenu_inst_indexOfItemWithTarget_andAction(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(target),
		actionSelector.SelectorAddress(),
	)

	return core.NSInteger(ret)

}

// IndexOfItemWithTitle Returns the index of the first menu item in the menu that has a specified title.
// https://developer.apple.com/documentation/appkit/nsmenu/1518237-indexofitemwithtitle?language=objc
func (x gen_NSMenu) IndexOfItemWithTitle(
	title core.NSStringRef,
) core.NSInteger {
	ret := C.NSMenu_inst_indexOfItemWithTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(title),
	)

	return core.NSInteger(ret)

}

// InitWithTitle_asNSMenu Initializes and returns a menu having the specified title and with autoenabling of menu items turned on.
// https://developer.apple.com/documentation/appkit/nsmenu/1518144-initwithtitle?language=objc
func (x gen_NSMenu) InitWithTitle_asNSMenu(
	title core.NSStringRef,
) NSMenu {
	ret := C.NSMenu_inst_initWithTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(title),
	)

	return NSMenu_fromPointer(ret)

}

// InsertItem_atIndex Inserts a menu item into the menu at a specific location.
// https://developer.apple.com/documentation/appkit/nsmenu/1518201-insertitem?language=objc
func (x gen_NSMenu) InsertItem_atIndex(
	newItem NSMenuItemRef,
	index core.NSInteger,
) {
	C.NSMenu_inst_insertItem_atIndex(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newItem),
		C.long(index),
	)

	return

}

// InsertItemWithTitle_action_keyEquivalent_atIndex Creates and adds a menu item at a specified location in the menu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518146-insertitemwithtitle?language=objc
func (x gen_NSMenu) InsertItemWithTitle_action_keyEquivalent_atIndex(
	string core.NSStringRef,
	selector objc.Selector,
	charCode core.NSStringRef,
	index core.NSInteger,
) NSMenuItem {
	ret := C.NSMenu_inst_insertItemWithTitle_action_keyEquivalent_atIndex(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		selector.SelectorAddress(),
		objc.RefPointer(charCode),
		C.long(index),
	)

	return NSMenuItem_fromPointer(ret)

}

// ItemAtIndex Returns the menu item at a specific location of the menu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518218-itematindex?language=objc
func (x gen_NSMenu) ItemAtIndex(
	index core.NSInteger,
) NSMenuItem {
	ret := C.NSMenu_inst_itemAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)

	return NSMenuItem_fromPointer(ret)

}

// ItemChanged Invoked when a menu item is modified visually (for example, its title changes).
// https://developer.apple.com/documentation/appkit/nsmenu/1518154-itemchanged?language=objc
func (x gen_NSMenu) ItemChanged(
	item NSMenuItemRef,
) {
	C.NSMenu_inst_itemChanged(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
	)

	return

}

// ItemWithTag Returns the first menu item in the menu with the specified tag.
// https://developer.apple.com/documentation/appkit/nsmenu/1518223-itemwithtag?language=objc
func (x gen_NSMenu) ItemWithTag(
	tag core.NSInteger,
) NSMenuItem {
	ret := C.NSMenu_inst_itemWithTag(
		unsafe.Pointer(x.Pointer()),
		C.long(tag),
	)

	return NSMenuItem_fromPointer(ret)

}

// ItemWithTitle Returns the first menu item in the menu with a specified title.
// https://developer.apple.com/documentation/appkit/nsmenu/1518248-itemwithtitle?language=objc
func (x gen_NSMenu) ItemWithTitle(
	title core.NSStringRef,
) NSMenuItem {
	ret := C.NSMenu_inst_itemWithTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(title),
	)

	return NSMenuItem_fromPointer(ret)

}

// PerformActionForItemAtIndex Causes the application to send the action message of a specified menu item to its target.
// https://developer.apple.com/documentation/appkit/nsmenu/1518210-performactionforitematindex?language=objc
func (x gen_NSMenu) PerformActionForItemAtIndex(
	index core.NSInteger,
) {
	C.NSMenu_inst_performActionForItemAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)

	return

}

// PerformKeyEquivalent Performs the action for the menu item that corresponds to the given key equivalent.
// https://developer.apple.com/documentation/appkit/nsmenu/1518198-performkeyequivalent?language=objc
func (x gen_NSMenu) PerformKeyEquivalent(
	event NSEventRef,
) bool {
	ret := C.NSMenu_inst_performKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return convertObjCBoolToGo(ret)

}

// PopUpMenuPositioningItem_atLocation_inView Pops up the menu at the specified location.
// https://developer.apple.com/documentation/appkit/nsmenu/1518212-popupmenupositioningitem?language=objc
func (x gen_NSMenu) PopUpMenuPositioningItem_atLocation_inView(
	item NSMenuItemRef,
	location core.NSPoint,
	view NSViewRef,
) bool {
	ret := C.NSMenu_inst_popUpMenuPositioningItem_atLocation_inView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
		*(*C.NSPoint)(unsafe.Pointer(&location)),
		objc.RefPointer(view),
	)

	return convertObjCBoolToGo(ret)

}

// RemoveAllItems Removes all the menu items in the menu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518234-removeallitems?language=objc
func (x gen_NSMenu) RemoveAllItems() {
	C.NSMenu_inst_removeAllItems(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// RemoveItem Removes a menu item from the menu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518257-removeitem?language=objc
func (x gen_NSMenu) RemoveItem(
	item NSMenuItemRef,
) {
	C.NSMenu_inst_removeItem(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
	)

	return

}

// RemoveItemAtIndex Removes the menu item at a specified location in the menu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518207-removeitematindex?language=objc
func (x gen_NSMenu) RemoveItemAtIndex(
	index core.NSInteger,
) {
	C.NSMenu_inst_removeItemAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)

	return

}

// SetSubmenu_forItem Assigns a menu to be a submenu of the menu controlled by a given menu item.
// https://developer.apple.com/documentation/appkit/nsmenu/1518194-setsubmenu?language=objc
func (x gen_NSMenu) SetSubmenu_forItem(
	menu NSMenuRef,
	item NSMenuItemRef,
) {
	C.NSMenu_inst_setSubmenu_forItem(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(menu),
		objc.RefPointer(item),
	)

	return

}

// SubmenuAction The action method assigned to menu items that open submenus.
// https://developer.apple.com/documentation/appkit/nsmenu/1518179-submenuaction?language=objc
func (x gen_NSMenu) SubmenuAction(
	sender objc.Ref,
) {
	C.NSMenu_inst_submenuAction(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// Update Enables or disables the menu items of the menu based on the NSMenuValidation informal protocol and sizes the menu to fit its current menu items if necessary.
// https://developer.apple.com/documentation/appkit/nsmenu/1518249-update?language=objc
func (x gen_NSMenu) Update() {
	C.NSMenu_inst_update(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Init_asNSMenu
func (x gen_NSMenu) Init_asNSMenu() NSMenu {
	ret := C.NSMenu_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenu_fromPointer(ret)

}

// MenuBarHeight The menu bar height for the main menu in pixels.
// https://developer.apple.com/documentation/appkit/nsmenu/1518141-menubarheight?language=objc
func (x gen_NSMenu) MenuBarHeight() core.CGFloat {
	ret := C.NSMenu_inst_menuBarHeight(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// NumberOfItems The number of menu items in the menu, including separator items.
// https://developer.apple.com/documentation/appkit/nsmenu/1518202-numberofitems?language=objc
func (x gen_NSMenu) NumberOfItems() core.NSInteger {
	ret := C.NSMenu_inst_numberOfItems(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// ItemArray An array containing the menu items in the menu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518186-itemarray?language=objc
func (x gen_NSMenu) ItemArray() core.NSArray {
	ret := C.NSMenu_inst_itemArray(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// SetItemArray An array containing the menu items in the menu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518186-itemarray?language=objc
func (x gen_NSMenu) SetItemArray(
	value core.NSArrayRef,
) {
	C.NSMenu_inst_setItemArray(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Supermenu The parent menu that contains the menu as a submenu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518204-supermenu?language=objc
func (x gen_NSMenu) Supermenu() NSMenu {
	ret := C.NSMenu_inst_supermenu(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenu_fromPointer(ret)

}

// SetSupermenu The parent menu that contains the menu as a submenu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518204-supermenu?language=objc
func (x gen_NSMenu) SetSupermenu(
	value NSMenuRef,
) {
	C.NSMenu_inst_setSupermenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// AutoenablesItems Indicates whether the menu automatically enables and disables its menu items.
// https://developer.apple.com/documentation/appkit/nsmenu/1518227-autoenablesitems?language=objc
func (x gen_NSMenu) AutoenablesItems() bool {
	ret := C.NSMenu_inst_autoenablesItems(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAutoenablesItems Indicates whether the menu automatically enables and disables its menu items.
// https://developer.apple.com/documentation/appkit/nsmenu/1518227-autoenablesitems?language=objc
func (x gen_NSMenu) SetAutoenablesItems(
	value bool,
) {
	C.NSMenu_inst_setAutoenablesItems(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// Font The font of the menu and its submenus.
// https://developer.apple.com/documentation/appkit/nsmenu/1518230-font?language=objc
func (x gen_NSMenu) Font() NSFont {
	ret := C.NSMenu_inst_font(
		unsafe.Pointer(x.Pointer()),
	)

	return NSFont_fromPointer(ret)

}

// SetFont The font of the menu and its submenus.
// https://developer.apple.com/documentation/appkit/nsmenu/1518230-font?language=objc
func (x gen_NSMenu) SetFont(
	value NSFontRef,
) {
	C.NSMenu_inst_setFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Title The title of the menu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518192-title?language=objc
func (x gen_NSMenu) Title() core.NSString {
	ret := C.NSMenu_inst_title(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetTitle The title of the menu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518192-title?language=objc
func (x gen_NSMenu) SetTitle(
	value core.NSStringRef,
) {
	C.NSMenu_inst_setTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// MinimumWidth The minimum width of the menu in screen coordinates.
// https://developer.apple.com/documentation/appkit/nsmenu/1518221-minimumwidth?language=objc
func (x gen_NSMenu) MinimumWidth() core.CGFloat {
	ret := C.NSMenu_inst_minimumWidth(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// SetMinimumWidth The minimum width of the menu in screen coordinates.
// https://developer.apple.com/documentation/appkit/nsmenu/1518221-minimumwidth?language=objc
func (x gen_NSMenu) SetMinimumWidth(
	value core.CGFloat,
) {
	C.NSMenu_inst_setMinimumWidth(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// Size The size of the menu in screen coordinates
// https://developer.apple.com/documentation/appkit/nsmenu/1518185-size?language=objc
func (x gen_NSMenu) Size() core.NSSize {
	ret := C.NSMenu_inst_size(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// AllowsContextMenuPlugIns Indicates whether the pop-up menu allows appending of contextual menu plug-in items.
// https://developer.apple.com/documentation/appkit/nsmenu/1518220-allowscontextmenuplugins?language=objc
func (x gen_NSMenu) AllowsContextMenuPlugIns() bool {
	ret := C.NSMenu_inst_allowsContextMenuPlugIns(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsContextMenuPlugIns Indicates whether the pop-up menu allows appending of contextual menu plug-in items.
// https://developer.apple.com/documentation/appkit/nsmenu/1518220-allowscontextmenuplugins?language=objc
func (x gen_NSMenu) SetAllowsContextMenuPlugIns(
	value bool,
) {
	C.NSMenu_inst_setAllowsContextMenuPlugIns(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// ShowsStateColumn Indicates whether the menu displays the state column.
// https://developer.apple.com/documentation/appkit/nsmenu/1518253-showsstatecolumn?language=objc
func (x gen_NSMenu) ShowsStateColumn() bool {
	ret := C.NSMenu_inst_showsStateColumn(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetShowsStateColumn Indicates whether the menu displays the state column.
// https://developer.apple.com/documentation/appkit/nsmenu/1518253-showsstatecolumn?language=objc
func (x gen_NSMenu) SetShowsStateColumn(
	value bool,
) {
	C.NSMenu_inst_setShowsStateColumn(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// HighlightedItem Indicates the currently highlighted item in the menu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518222-highlighteditem?language=objc
func (x gen_NSMenu) HighlightedItem() NSMenuItem {
	ret := C.NSMenu_inst_highlightedItem(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenuItem_fromPointer(ret)

}

// Delegate The delegate of the menu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518169-delegate?language=objc
func (x gen_NSMenu) Delegate() objc.Object {
	ret := C.NSMenu_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetDelegate The delegate of the menu.
// https://developer.apple.com/documentation/appkit/nsmenu/1518169-delegate?language=objc
func (x gen_NSMenu) SetDelegate(
	value objc.Ref,
) {
	C.NSMenu_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

type NSPopoverRef interface {
	Pointer() uintptr
	Init_asNSPopover() NSPopover
}

type gen_NSPopover struct {
	objc.Object
}

func NSPopover_fromPointer(ptr unsafe.Pointer) NSPopover {
	return NSPopover{gen_NSPopover{
		objc.Object_fromPointer(ptr),
	}}
}

func NSPopover_fromRef(ref objc.Ref) NSPopover {
	return NSPopover_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// Close Forces the popover to close without consulting its delegate.
// https://developer.apple.com/documentation/appkit/nspopover/1526823-close?language=objc
func (x gen_NSPopover) Close() {
	C.NSPopover_inst_close(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Init_asNSPopover
// https://developer.apple.com/documentation/appkit/nspopover/1526851-init?language=objc
func (x gen_NSPopover) Init_asNSPopover() NSPopover {
	ret := C.NSPopover_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSPopover_fromPointer(ret)

}

// PerformClose Attempts to close the popover.
// https://developer.apple.com/documentation/appkit/nspopover/1534290-performclose?language=objc
func (x gen_NSPopover) PerformClose(
	sender objc.Ref,
) {
	C.NSPopover_inst_performClose(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// Behavior Specifies the behavior of the popover.
// https://developer.apple.com/documentation/appkit/nspopover/1533539-behavior?language=objc
func (x gen_NSPopover) Behavior() core.NSInteger {
	ret := C.NSPopover_inst_behavior(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// SetBehavior Specifies the behavior of the popover.
// https://developer.apple.com/documentation/appkit/nspopover/1533539-behavior?language=objc
func (x gen_NSPopover) SetBehavior(
	value core.NSInteger,
) {
	C.NSPopover_inst_setBehavior(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return

}

// PositioningRect The rectangle within the positioning view relative to which the popover should be positioned.
// https://developer.apple.com/documentation/appkit/nspopover/1526090-positioningrect?language=objc
func (x gen_NSPopover) PositioningRect() core.NSRect {
	ret := C.NSPopover_inst_positioningRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// SetPositioningRect The rectangle within the positioning view relative to which the popover should be positioned.
// https://developer.apple.com/documentation/appkit/nspopover/1526090-positioningrect?language=objc
func (x gen_NSPopover) SetPositioningRect(
	value core.NSRect,
) {
	C.NSPopover_inst_setPositioningRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return

}

// Animates Specifies if the popover is to be animated.
// https://developer.apple.com/documentation/appkit/nspopover/1526527-animates?language=objc
func (x gen_NSPopover) Animates() bool {
	ret := C.NSPopover_inst_animates(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAnimates Specifies if the popover is to be animated.
// https://developer.apple.com/documentation/appkit/nspopover/1526527-animates?language=objc
func (x gen_NSPopover) SetAnimates(
	value bool,
) {
	C.NSPopover_inst_setAnimates(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// ContentSize The content size of the popover.
// https://developer.apple.com/documentation/appkit/nspopover/1524677-contentsize?language=objc
func (x gen_NSPopover) ContentSize() core.NSSize {
	ret := C.NSPopover_inst_contentSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetContentSize The content size of the popover.
// https://developer.apple.com/documentation/appkit/nspopover/1524677-contentsize?language=objc
func (x gen_NSPopover) SetContentSize(
	value core.NSSize,
) {
	C.NSPopover_inst_setContentSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// IsShown The display state of the popover.
// https://developer.apple.com/documentation/appkit/nspopover/1535120-shown?language=objc
func (x gen_NSPopover) IsShown() bool {
	ret := C.NSPopover_inst_isShown(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsDetached A Boolean value that indicates whether the window created by a popover's detachment is automatically created.
// https://developer.apple.com/documentation/appkit/nspopover/1534278-detached?language=objc
func (x gen_NSPopover) IsDetached() bool {
	ret := C.NSPopover_inst_isDetached(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

type NSMenuItemRef interface {
	Pointer() uintptr
	Init_asNSMenuItem() NSMenuItem
}

type gen_NSMenuItem struct {
	objc.Object
}

func NSMenuItem_fromPointer(ptr unsafe.Pointer) NSMenuItem {
	return NSMenuItem{gen_NSMenuItem{
		objc.Object_fromPointer(ptr),
	}}
}

func NSMenuItem_fromRef(ref objc.Ref) NSMenuItem {
	return NSMenuItem_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// InitWithTitle_action_keyEquivalent_asNSMenuItem Returns an initialized instance of NSMenuItem.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514858-initwithtitle?language=objc
func (x gen_NSMenuItem) InitWithTitle_action_keyEquivalent_asNSMenuItem(
	string core.NSStringRef,
	selector objc.Selector,
	charCode core.NSStringRef,
) NSMenuItem {
	ret := C.NSMenuItem_inst_initWithTitle_action_keyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		selector.SelectorAddress(),
		objc.RefPointer(charCode),
	)

	return NSMenuItem_fromPointer(ret)

}

// Init_asNSMenuItem
func (x gen_NSMenuItem) Init_asNSMenuItem() NSMenuItem {
	ret := C.NSMenuItem_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenuItem_fromPointer(ret)

}

// IsEnabled A Boolean value that indicates whether the menu item is enabled.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514863-enabled?language=objc
func (x gen_NSMenuItem) IsEnabled() bool {
	ret := C.NSMenuItem_inst_isEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetEnabled A Boolean value that indicates whether the menu item is enabled.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514863-enabled?language=objc
func (x gen_NSMenuItem) SetEnabled(
	value bool,
) {
	C.NSMenuItem_inst_setEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsHidden A Boolean value that indicates whether the menu item is hidden.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514846-hidden?language=objc
func (x gen_NSMenuItem) IsHidden() bool {
	ret := C.NSMenuItem_inst_isHidden(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetHidden A Boolean value that indicates whether the menu item is hidden.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514846-hidden?language=objc
func (x gen_NSMenuItem) SetHidden(
	value bool,
) {
	C.NSMenuItem_inst_setHidden(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsHiddenOrHasHiddenAncestor A Boolean value that indicates whether the menu item or any of its superitems is hidden.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514832-hiddenorhashiddenancestor?language=objc
func (x gen_NSMenuItem) IsHiddenOrHasHiddenAncestor() bool {
	ret := C.NSMenuItem_inst_isHiddenOrHasHiddenAncestor(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Target The menu item's target.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514843-target?language=objc
func (x gen_NSMenuItem) Target() objc.Object {
	ret := C.NSMenuItem_inst_target(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetTarget The menu item's target.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514843-target?language=objc
func (x gen_NSMenuItem) SetTarget(
	value objc.Ref,
) {
	C.NSMenuItem_inst_setTarget(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Action The menu item's action-method selector.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514825-action?language=objc
func (x gen_NSMenuItem) Action() objc.Selector {
	ret := C.NSMenuItem_inst_action(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.SelectorAt(ret)

}

// SetAction The menu item's action-method selector.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514825-action?language=objc
func (x gen_NSMenuItem) SetAction(
	value objc.Selector,
) {
	C.NSMenuItem_inst_setAction(
		unsafe.Pointer(x.Pointer()),
		value.SelectorAddress(),
	)

	return

}

// Title The menu item's title.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514805-title?language=objc
func (x gen_NSMenuItem) Title() core.NSString {
	ret := C.NSMenuItem_inst_title(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetTitle The menu item's title.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514805-title?language=objc
func (x gen_NSMenuItem) SetTitle(
	value core.NSStringRef,
) {
	C.NSMenuItem_inst_setTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// AttributedTitle A custom string for a menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514860-attributedtitle?language=objc
func (x gen_NSMenuItem) AttributedTitle() core.NSAttributedString {
	ret := C.NSMenuItem_inst_attributedTitle(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSAttributedString_fromPointer(ret)

}

// SetAttributedTitle A custom string for a menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514860-attributedtitle?language=objc
func (x gen_NSMenuItem) SetAttributedTitle(
	value core.NSAttributedStringRef,
) {
	C.NSMenuItem_inst_setAttributedTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Tag The menu item's tag.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514840-tag?language=objc
func (x gen_NSMenuItem) Tag() core.NSInteger {
	ret := C.NSMenuItem_inst_tag(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// SetTag The menu item's tag.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514840-tag?language=objc
func (x gen_NSMenuItem) SetTag(
	value core.NSInteger,
) {
	C.NSMenuItem_inst_setTag(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return

}

// State The state of the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514804-state?language=objc
func (x gen_NSMenuItem) State() core.NSInteger {
	ret := C.NSMenuItem_inst_state(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// SetState The state of the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514804-state?language=objc
func (x gen_NSMenuItem) SetState(
	value core.NSInteger,
) {
	C.NSMenuItem_inst_setState(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return

}

// Image The menu item’s image.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514819-image?language=objc
func (x gen_NSMenuItem) Image() NSImage {
	ret := C.NSMenuItem_inst_image(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_fromPointer(ret)

}

// SetImage The menu item’s image.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514819-image?language=objc
func (x gen_NSMenuItem) SetImage(
	value NSImageRef,
) {
	C.NSMenuItem_inst_setImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// OnStateImage The image of the menu item that indicates an “on” state.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514861-onstateimage?language=objc
func (x gen_NSMenuItem) OnStateImage() NSImage {
	ret := C.NSMenuItem_inst_onStateImage(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_fromPointer(ret)

}

// SetOnStateImage The image of the menu item that indicates an “on” state.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514861-onstateimage?language=objc
func (x gen_NSMenuItem) SetOnStateImage(
	value NSImageRef,
) {
	C.NSMenuItem_inst_setOnStateImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// OffStateImage The image of the menu item that indicates an “off” state.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514821-offstateimage?language=objc
func (x gen_NSMenuItem) OffStateImage() NSImage {
	ret := C.NSMenuItem_inst_offStateImage(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_fromPointer(ret)

}

// SetOffStateImage The image of the menu item that indicates an “off” state.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514821-offstateimage?language=objc
func (x gen_NSMenuItem) SetOffStateImage(
	value NSImageRef,
) {
	C.NSMenuItem_inst_setOffStateImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// MixedStateImage The image of the menu item that indicates a “mixed” state, that is, a state neither “on” nor “off.”
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514827-mixedstateimage?language=objc
func (x gen_NSMenuItem) MixedStateImage() NSImage {
	ret := C.NSMenuItem_inst_mixedStateImage(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_fromPointer(ret)

}

// SetMixedStateImage The image of the menu item that indicates a “mixed” state, that is, a state neither “on” nor “off.”
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514827-mixedstateimage?language=objc
func (x gen_NSMenuItem) SetMixedStateImage(
	value NSImageRef,
) {
	C.NSMenuItem_inst_setMixedStateImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Submenu The submenu of the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514845-submenu?language=objc
func (x gen_NSMenuItem) Submenu() NSMenu {
	ret := C.NSMenuItem_inst_submenu(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenu_fromPointer(ret)

}

// SetSubmenu The submenu of the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514845-submenu?language=objc
func (x gen_NSMenuItem) SetSubmenu(
	value NSMenuRef,
) {
	C.NSMenuItem_inst_setSubmenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// HasSubmenu A Boolean value that indicates whether the menu item has a submenu.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514817-hassubmenu?language=objc
func (x gen_NSMenuItem) HasSubmenu() bool {
	ret := C.NSMenuItem_inst_hasSubmenu(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// ParentItem The menu item whose submenu contains the receiver.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514813-parentitem?language=objc
func (x gen_NSMenuItem) ParentItem() NSMenuItem {
	ret := C.NSMenuItem_inst_parentItem(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenuItem_fromPointer(ret)

}

// IsSeparatorItem A menu item that is used to separate logical groups of menu commands.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514837-separatoritem?language=objc
func (x gen_NSMenuItem) IsSeparatorItem() bool {
	ret := C.NSMenuItem_inst_isSeparatorItem(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Menu The menu item’s menu.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514830-menu?language=objc
func (x gen_NSMenuItem) Menu() NSMenu {
	ret := C.NSMenuItem_inst_menu(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenu_fromPointer(ret)

}

// SetMenu The menu item’s menu.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514830-menu?language=objc
func (x gen_NSMenuItem) SetMenu(
	value NSMenuRef,
) {
	C.NSMenuItem_inst_setMenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// KeyEquivalent The menu item’s unmodified key equivalent.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514842-keyequivalent?language=objc
func (x gen_NSMenuItem) KeyEquivalent() core.NSString {
	ret := C.NSMenuItem_inst_keyEquivalent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetKeyEquivalent The menu item’s unmodified key equivalent.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514842-keyequivalent?language=objc
func (x gen_NSMenuItem) SetKeyEquivalent(
	value core.NSStringRef,
) {
	C.NSMenuItem_inst_setKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// UserKeyEquivalent The user-assigned key equivalent for the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514850-userkeyequivalent?language=objc
func (x gen_NSMenuItem) UserKeyEquivalent() core.NSString {
	ret := C.NSMenuItem_inst_userKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// IsAlternate A Boolean value that marks the menu item as an alternate to the previous menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514823-alternate?language=objc
func (x gen_NSMenuItem) IsAlternate() bool {
	ret := C.NSMenuItem_inst_isAlternate(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAlternate A Boolean value that marks the menu item as an alternate to the previous menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514823-alternate?language=objc
func (x gen_NSMenuItem) SetAlternate(
	value bool,
) {
	C.NSMenuItem_inst_setAlternate(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IndentationLevel The menu item indentation level for the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514809-indentationlevel?language=objc
func (x gen_NSMenuItem) IndentationLevel() core.NSInteger {
	ret := C.NSMenuItem_inst_indentationLevel(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// SetIndentationLevel The menu item indentation level for the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514809-indentationlevel?language=objc
func (x gen_NSMenuItem) SetIndentationLevel(
	value core.NSInteger,
) {
	C.NSMenuItem_inst_setIndentationLevel(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return

}

// ToolTip A help tag for the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514848-tooltip?language=objc
func (x gen_NSMenuItem) ToolTip() core.NSString {
	ret := C.NSMenuItem_inst_toolTip(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetToolTip A help tag for the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514848-tooltip?language=objc
func (x gen_NSMenuItem) SetToolTip(
	value core.NSStringRef,
) {
	C.NSMenuItem_inst_setToolTip(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// RepresentedObject The object represented by the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514834-representedobject?language=objc
func (x gen_NSMenuItem) RepresentedObject() objc.Object {
	ret := C.NSMenuItem_inst_representedObject(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetRepresentedObject The object represented by the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514834-representedobject?language=objc
func (x gen_NSMenuItem) SetRepresentedObject(
	value objc.Ref,
) {
	C.NSMenuItem_inst_setRepresentedObject(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// View The content view for the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514835-view?language=objc
func (x gen_NSMenuItem) View() NSView {
	ret := C.NSMenuItem_inst_view(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_fromPointer(ret)

}

// SetView The content view for the menu item.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514835-view?language=objc
func (x gen_NSMenuItem) SetView(
	value NSViewRef,
) {
	C.NSMenuItem_inst_setView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// IsHighlighted A Boolean value that indicates whether the menu item should be drawn highlighted.
// https://developer.apple.com/documentation/appkit/nsmenuitem/1514856-highlighted?language=objc
func (x gen_NSMenuItem) IsHighlighted() bool {
	ret := C.NSMenuItem_inst_isHighlighted(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AllowsAutomaticKeyEquivalentLocalization
// https://developer.apple.com/documentation/appkit/nsmenuitem/3787554-allowsautomatickeyequivalentloca?language=objc
func (x gen_NSMenuItem) AllowsAutomaticKeyEquivalentLocalization() bool {
	ret := C.NSMenuItem_inst_allowsAutomaticKeyEquivalentLocalization(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsAutomaticKeyEquivalentLocalization
// https://developer.apple.com/documentation/appkit/nsmenuitem/3787554-allowsautomatickeyequivalentloca?language=objc
func (x gen_NSMenuItem) SetAllowsAutomaticKeyEquivalentLocalization(
	value bool,
) {
	C.NSMenuItem_inst_setAllowsAutomaticKeyEquivalentLocalization(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AllowsAutomaticKeyEquivalentMirroring
// https://developer.apple.com/documentation/appkit/nsmenuitem/3787555-allowsautomatickeyequivalentmirr?language=objc
func (x gen_NSMenuItem) AllowsAutomaticKeyEquivalentMirroring() bool {
	ret := C.NSMenuItem_inst_allowsAutomaticKeyEquivalentMirroring(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsAutomaticKeyEquivalentMirroring
// https://developer.apple.com/documentation/appkit/nsmenuitem/3787555-allowsautomatickeyequivalentmirr?language=objc
func (x gen_NSMenuItem) SetAllowsAutomaticKeyEquivalentMirroring(
	value bool,
) {
	C.NSMenuItem_inst_setAllowsAutomaticKeyEquivalentMirroring(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AllowsKeyEquivalentWhenHidden
// https://developer.apple.com/documentation/appkit/nsmenuitem/2880316-allowskeyequivalentwhenhidden?language=objc
func (x gen_NSMenuItem) AllowsKeyEquivalentWhenHidden() bool {
	ret := C.NSMenuItem_inst_allowsKeyEquivalentWhenHidden(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsKeyEquivalentWhenHidden
// https://developer.apple.com/documentation/appkit/nsmenuitem/2880316-allowskeyequivalentwhenhidden?language=objc
func (x gen_NSMenuItem) SetAllowsKeyEquivalentWhenHidden(
	value bool,
) {
	C.NSMenuItem_inst_setAllowsKeyEquivalentWhenHidden(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

type NSRunningApplicationRef interface {
	Pointer() uintptr
	Init_asNSRunningApplication() NSRunningApplication
}

type gen_NSRunningApplication struct {
	objc.Object
}

func NSRunningApplication_fromPointer(ptr unsafe.Pointer) NSRunningApplication {
	return NSRunningApplication{gen_NSRunningApplication{
		objc.Object_fromPointer(ptr),
	}}
}

func NSRunningApplication_fromRef(ref objc.Ref) NSRunningApplication {
	return NSRunningApplication_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// ForceTerminate Attempts to force the receiver to quit.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1530370-forceterminate?language=objc
func (x gen_NSRunningApplication) ForceTerminate() bool {
	ret := C.NSRunningApplication_inst_forceTerminate(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Hide Attempts to hide or the application.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1526608-hide?language=objc
func (x gen_NSRunningApplication) Hide() bool {
	ret := C.NSRunningApplication_inst_hide(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Terminate Attempts to quit the receiver normally.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1528922-terminate?language=objc
func (x gen_NSRunningApplication) Terminate() bool {
	ret := C.NSRunningApplication_inst_terminate(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Unhide Attempts to unhide or the application.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1534676-unhide?language=objc
func (x gen_NSRunningApplication) Unhide() bool {
	ret := C.NSRunningApplication_inst_unhide(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Init_asNSRunningApplication
func (x gen_NSRunningApplication) Init_asNSRunningApplication() NSRunningApplication {
	ret := C.NSRunningApplication_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSRunningApplication_fromPointer(ret)

}

// IsActive Indicates whether the application is currently frontmost.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1528778-active?language=objc
func (x gen_NSRunningApplication) IsActive() bool {
	ret := C.NSRunningApplication_inst_isActive(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// ActivationPolicy Indicates the activation policy of the application.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1533103-activationpolicy?language=objc
func (x gen_NSRunningApplication) ActivationPolicy() core.NSInteger {
	ret := C.NSRunningApplication_inst_activationPolicy(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// IsHidden Indicates whether the application is currently hidden.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1525949-hidden?language=objc
func (x gen_NSRunningApplication) IsHidden() bool {
	ret := C.NSRunningApplication_inst_isHidden(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// LocalizedName Indicates the localized name of the application.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1526751-localizedname?language=objc
func (x gen_NSRunningApplication) LocalizedName() core.NSString {
	ret := C.NSRunningApplication_inst_localizedName(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// Icon Returns the icon for the receiver’s application.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1529885-icon?language=objc
func (x gen_NSRunningApplication) Icon() NSImage {
	ret := C.NSRunningApplication_inst_icon(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_fromPointer(ret)

}

// BundleIdentifier Indicates the CFBundleIdentifier of the application.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1529140-bundleidentifier?language=objc
func (x gen_NSRunningApplication) BundleIdentifier() core.NSString {
	ret := C.NSRunningApplication_inst_bundleIdentifier(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// BundleURL Indicates the URL to the application's bundle.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1535500-bundleurl?language=objc
func (x gen_NSRunningApplication) BundleURL() core.NSURL {
	ret := C.NSRunningApplication_inst_bundleURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_fromPointer(ret)

}

// ExecutableArchitecture Indicates the executing processor architecture for the application.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1524287-executablearchitecture?language=objc
func (x gen_NSRunningApplication) ExecutableArchitecture() core.NSInteger {
	ret := C.NSRunningApplication_inst_executableArchitecture(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// ExecutableURL Indicates the URL to the application's executable.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1531062-executableurl?language=objc
func (x gen_NSRunningApplication) ExecutableURL() core.NSURL {
	ret := C.NSRunningApplication_inst_executableURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_fromPointer(ret)

}

// IsFinishedLaunching Indicates whether the receiver’s process has finished launching,
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1532002-finishedlaunching?language=objc
func (x gen_NSRunningApplication) IsFinishedLaunching() bool {
	ret := C.NSRunningApplication_inst_isFinishedLaunching(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// OwnsMenuBar Returns whether the application owns the current menu bar.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1525915-ownsmenubar?language=objc
func (x gen_NSRunningApplication) OwnsMenuBar() bool {
	ret := C.NSRunningApplication_inst_ownsMenuBar(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsTerminated Indicates that the receiver’s application has terminated.
// https://developer.apple.com/documentation/appkit/nsrunningapplication/1532239-terminated?language=objc
func (x gen_NSRunningApplication) IsTerminated() bool {
	ret := C.NSRunningApplication_inst_isTerminated(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

type NSScreenRef interface {
	Pointer() uintptr
	Init_asNSScreen() NSScreen
}

type gen_NSScreen struct {
	objc.Object
}

func NSScreen_fromPointer(ptr unsafe.Pointer) NSScreen {
	return NSScreen{gen_NSScreen{
		objc.Object_fromPointer(ptr),
	}}
}

func NSScreen_fromRef(ref objc.Ref) NSScreen {
	return NSScreen_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// ConvertRectFromBacking Converts the rectangle from the device pixel aligned coordinates system of a screen.
// https://developer.apple.com/documentation/appkit/nsscreen/1388364-convertrectfrombacking?language=objc
func (x gen_NSScreen) ConvertRectFromBacking(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSScreen_inst_convertRectFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// ConvertRectToBacking Converts the rectangle to the device pixel aligned coordinates system of a screen.
// https://developer.apple.com/documentation/appkit/nsscreen/1388389-convertrecttobacking?language=objc
func (x gen_NSScreen) ConvertRectToBacking(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSScreen_inst_convertRectToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// Init_asNSScreen
func (x gen_NSScreen) Init_asNSScreen() NSScreen {
	ret := C.NSScreen_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSScreen_fromPointer(ret)

}

// Frame The dimensions and location of the screen.
// https://developer.apple.com/documentation/appkit/nsscreen/1388387-frame?language=objc
func (x gen_NSScreen) Frame() core.NSRect {
	ret := C.NSScreen_inst_frame(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// DeviceDescription The device dictionary for the screen.
// https://developer.apple.com/documentation/appkit/nsscreen/1388360-devicedescription?language=objc
func (x gen_NSScreen) DeviceDescription() core.NSDictionary {
	ret := C.NSScreen_inst_deviceDescription(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_fromPointer(ret)

}

// VisibleFrame The current location and dimensions of the visible screen.
// https://developer.apple.com/documentation/appkit/nsscreen/1388369-visibleframe?language=objc
func (x gen_NSScreen) VisibleFrame() core.NSRect {
	ret := C.NSScreen_inst_visibleFrame(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// BackingScaleFactor The backing store pixel scale factor for the screen.
// https://developer.apple.com/documentation/appkit/nsscreen/1388385-backingscalefactor?language=objc
func (x gen_NSScreen) BackingScaleFactor() core.CGFloat {
	ret := C.NSScreen_inst_backingScaleFactor(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// MaximumPotentialExtendedDynamicRangeColorComponentValue The maximum possible color component value for the screen when it's in extended dynamic range (EDR) mode.
// https://developer.apple.com/documentation/appkit/nsscreen/3180381-maximumpotentialextendeddynamicr?language=objc
func (x gen_NSScreen) MaximumPotentialExtendedDynamicRangeColorComponentValue() core.CGFloat {
	ret := C.NSScreen_inst_maximumPotentialExtendedDynamicRangeColorComponentValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// MaximumExtendedDynamicRangeColorComponentValue The current maximum color component value for the screen.
// https://developer.apple.com/documentation/appkit/nsscreen/1388362-maximumextendeddynamicrangecolor?language=objc
func (x gen_NSScreen) MaximumExtendedDynamicRangeColorComponentValue() core.CGFloat {
	ret := C.NSScreen_inst_maximumExtendedDynamicRangeColorComponentValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// MaximumReferenceExtendedDynamicRangeColorComponentValue The current maximum color component value for reference rendering to the screen.
// https://developer.apple.com/documentation/appkit/nsscreen/3180382-maximumreferenceextendeddynamicr?language=objc
func (x gen_NSScreen) MaximumReferenceExtendedDynamicRangeColorComponentValue() core.CGFloat {
	ret := C.NSScreen_inst_maximumReferenceExtendedDynamicRangeColorComponentValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// LocalizedName
// https://developer.apple.com/documentation/appkit/nsscreen/3228043-localizedname?language=objc
func (x gen_NSScreen) LocalizedName() core.NSString {
	ret := C.NSScreen_inst_localizedName(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// MaximumFramesPerSecond
// https://developer.apple.com/documentation/appkit/nsscreen/3824745-maximumframespersecond?language=objc
func (x gen_NSScreen) MaximumFramesPerSecond() core.NSInteger {
	ret := C.NSScreen_inst_maximumFramesPerSecond(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

type NSStatusBarRef interface {
	Pointer() uintptr
	Init_asNSStatusBar() NSStatusBar
}

type gen_NSStatusBar struct {
	objc.Object
}

func NSStatusBar_fromPointer(ptr unsafe.Pointer) NSStatusBar {
	return NSStatusBar{gen_NSStatusBar{
		objc.Object_fromPointer(ptr),
	}}
}

func NSStatusBar_fromRef(ref objc.Ref) NSStatusBar {
	return NSStatusBar_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// RemoveStatusItem Removes the specified status item from the receiver.
// https://developer.apple.com/documentation/appkit/nsstatusbar/1530377-removestatusitem?language=objc
func (x gen_NSStatusBar) RemoveStatusItem(
	item NSStatusItemRef,
) {
	C.NSStatusBar_inst_removeStatusItem(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
	)

	return

}

// StatusItemWithLength Returns a newly created status item that has been allotted a specified space within the status bar.
// https://developer.apple.com/documentation/appkit/nsstatusbar/1532895-statusitemwithlength?language=objc
func (x gen_NSStatusBar) StatusItemWithLength(
	length core.CGFloat,
) NSStatusItem {
	ret := C.NSStatusBar_inst_statusItemWithLength(
		unsafe.Pointer(x.Pointer()),
		C.double(length),
	)

	return NSStatusItem_fromPointer(ret)

}

// Init_asNSStatusBar
func (x gen_NSStatusBar) Init_asNSStatusBar() NSStatusBar {
	ret := C.NSStatusBar_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSStatusBar_fromPointer(ret)

}

// IsVertical A Boolean value indicating whether the status bar has a vertical orientation.
// https://developer.apple.com/documentation/appkit/nsstatusbar/1530580-vertical?language=objc
func (x gen_NSStatusBar) IsVertical() bool {
	ret := C.NSStatusBar_inst_isVertical(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Thickness The thickness of the status bar, in pixels.
// https://developer.apple.com/documentation/appkit/nsstatusbar/1534591-thickness?language=objc
func (x gen_NSStatusBar) Thickness() core.CGFloat {
	ret := C.NSStatusBar_inst_thickness(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

type NSStatusBarButtonRef interface {
	Pointer() uintptr
	Init_asNSStatusBarButton() NSStatusBarButton
}

type gen_NSStatusBarButton struct {
	NSButton
}

func NSStatusBarButton_fromPointer(ptr unsafe.Pointer) NSStatusBarButton {
	return NSStatusBarButton{gen_NSStatusBarButton{
		NSButton_fromPointer(ptr),
	}}
}

func NSStatusBarButton_fromRef(ref objc.Ref) NSStatusBarButton {
	return NSStatusBarButton_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init_asNSStatusBarButton
func (x gen_NSStatusBarButton) Init_asNSStatusBarButton() NSStatusBarButton {
	ret := C.NSStatusBarButton_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSStatusBarButton_fromPointer(ret)

}

// AppearsDisabled
// https://developer.apple.com/documentation/appkit/nsstatusbarbutton/1409292-appearsdisabled?language=objc
func (x gen_NSStatusBarButton) AppearsDisabled() bool {
	ret := C.NSStatusBarButton_inst_appearsDisabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAppearsDisabled
// https://developer.apple.com/documentation/appkit/nsstatusbarbutton/1409292-appearsdisabled?language=objc
func (x gen_NSStatusBarButton) SetAppearsDisabled(
	value bool,
) {
	C.NSStatusBarButton_inst_setAppearsDisabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

type NSStatusItemRef interface {
	Pointer() uintptr
	Init_asNSStatusItem() NSStatusItem
}

type gen_NSStatusItem struct {
	objc.Object
}

func NSStatusItem_fromPointer(ptr unsafe.Pointer) NSStatusItem {
	return NSStatusItem{gen_NSStatusItem{
		objc.Object_fromPointer(ptr),
	}}
}

func NSStatusItem_fromRef(ref objc.Ref) NSStatusItem {
	return NSStatusItem_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init_asNSStatusItem
func (x gen_NSStatusItem) Init_asNSStatusItem() NSStatusItem {
	ret := C.NSStatusItem_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSStatusItem_fromPointer(ret)

}

// StatusBar The status bar that displays the status item.
// https://developer.apple.com/documentation/appkit/nsstatusitem/1525951-statusbar?language=objc
func (x gen_NSStatusItem) StatusBar() NSStatusBar {
	ret := C.NSStatusItem_inst_statusBar(
		unsafe.Pointer(x.Pointer()),
	)

	return NSStatusBar_fromPointer(ret)

}

// Button The button displayed in the status bar.
// https://developer.apple.com/documentation/appkit/nsstatusitem/1535056-button?language=objc
func (x gen_NSStatusItem) Button() NSStatusBarButton {
	ret := C.NSStatusItem_inst_button(
		unsafe.Pointer(x.Pointer()),
	)

	return NSStatusBarButton_fromPointer(ret)

}

// Menu The pull-down menu displayed when the user clicks the status item.
// https://developer.apple.com/documentation/appkit/nsstatusitem/1535918-menu?language=objc
func (x gen_NSStatusItem) Menu() NSMenu {
	ret := C.NSStatusItem_inst_menu(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenu_fromPointer(ret)

}

// SetMenu The pull-down menu displayed when the user clicks the status item.
// https://developer.apple.com/documentation/appkit/nsstatusitem/1535918-menu?language=objc
func (x gen_NSStatusItem) SetMenu(
	value NSMenuRef,
) {
	C.NSStatusItem_inst_setMenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// IsVisible A Boolean value indicating if the menu bar currently displays the status item.
// https://developer.apple.com/documentation/appkit/nsstatusitem/1644025-visible?language=objc
func (x gen_NSStatusItem) IsVisible() bool {
	ret := C.NSStatusItem_inst_isVisible(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetVisible A Boolean value indicating if the menu bar currently displays the status item.
// https://developer.apple.com/documentation/appkit/nsstatusitem/1644025-visible?language=objc
func (x gen_NSStatusItem) SetVisible(
	value bool,
) {
	C.NSStatusItem_inst_setVisible(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// Length The amount of space in the status bar that should be allocated to the status item.
// https://developer.apple.com/documentation/appkit/nsstatusitem/1529402-length?language=objc
func (x gen_NSStatusItem) Length() core.CGFloat {
	ret := C.NSStatusItem_inst_length(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// SetLength The amount of space in the status bar that should be allocated to the status item.
// https://developer.apple.com/documentation/appkit/nsstatusitem/1529402-length?language=objc
func (x gen_NSStatusItem) SetLength(
	value core.CGFloat,
) {
	C.NSStatusItem_inst_setLength(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

type NSTextRef interface {
	Pointer() uintptr
	Init_asNSText() NSText
}

type gen_NSText struct {
	NSView
}

func NSText_fromPointer(ptr unsafe.Pointer) NSText {
	return NSText{gen_NSText{
		NSView_fromPointer(ptr),
	}}
}

func NSText_fromRef(ref objc.Ref) NSText {
	return NSText_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// AlignCenter This action method applies center alignment to selected paragraphs (or all text if the receiver is a plain text object).
// https://developer.apple.com/documentation/appkit/nstext/1535569-aligncenter?language=objc
func (x gen_NSText) AlignCenter(
	sender objc.Ref,
) {
	C.NSText_inst_alignCenter(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// AlignLeft This action method applies left alignment to selected paragraphs (or all text if the receiver is a plain text object).
// https://developer.apple.com/documentation/appkit/nstext/1535705-alignleft?language=objc
func (x gen_NSText) AlignLeft(
	sender objc.Ref,
) {
	C.NSText_inst_alignLeft(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// AlignRight This action method applies right alignment to selected paragraphs (or all text if the receiver is a plain text object).
// https://developer.apple.com/documentation/appkit/nstext/1526477-alignright?language=objc
func (x gen_NSText) AlignRight(
	sender objc.Ref,
) {
	C.NSText_inst_alignRight(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ChangeFont This action method changes the font of the selection for a rich text object, or of all text for a plain text object.
// https://developer.apple.com/documentation/appkit/nstext/1531459-changefont?language=objc
func (x gen_NSText) ChangeFont(
	sender objc.Ref,
) {
	C.NSText_inst_changeFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// CheckSpelling This action method searches for a misspelled word in the receiver’s text.
// https://developer.apple.com/documentation/appkit/nstext/1534926-checkspelling?language=objc
func (x gen_NSText) CheckSpelling(
	sender objc.Ref,
) {
	C.NSText_inst_checkSpelling(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// Copy This action method copies the selected text onto the general pasteboard, in as many formats as the receiver supports.
// https://developer.apple.com/documentation/appkit/nstext/1525497-copy?language=objc
func (x gen_NSText) Copy(
	sender objc.Ref,
) {
	C.NSText_inst_copy(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// CopyFont This action method copies the font information for the first character of the selection (or for the insertion point) onto the font pasteboard, as NSFontPboardType.
// https://developer.apple.com/documentation/appkit/nstext/1531255-copyfont?language=objc
func (x gen_NSText) CopyFont(
	sender objc.Ref,
) {
	C.NSText_inst_copyFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// CopyRuler This action method copies the paragraph style information for first selected paragraph onto the ruler pasteboard, as NSRulerPboardType, and expands the selection to paragraph boundaries.
// https://developer.apple.com/documentation/appkit/nstext/1533303-copyruler?language=objc
func (x gen_NSText) CopyRuler(
	sender objc.Ref,
) {
	C.NSText_inst_copyRuler(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// Cut This action method deletes the selected text and places it onto the general pasteboard, in as many formats as the receiver supports.
// https://developer.apple.com/documentation/appkit/nstext/1524858-cut?language=objc
func (x gen_NSText) Cut(
	sender objc.Ref,
) {
	C.NSText_inst_cut(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// Delete This action method deletes the selected text.
// https://developer.apple.com/documentation/appkit/nstext/1524660-delete?language=objc
func (x gen_NSText) Delete(
	sender objc.Ref,
) {
	C.NSText_inst_delete(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// InitWithFrame_asNSText
// https://developer.apple.com/documentation/appkit/nstext/1525191-initwithframe?language=objc
func (x gen_NSText) InitWithFrame_asNSText(
	frameRect core.NSRect,
) NSText {
	ret := C.NSText_inst_initWithFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)

	return NSText_fromPointer(ret)

}

// Paste This action method pastes text from the general pasteboard at the insertion point or over the selection.
// https://developer.apple.com/documentation/appkit/nstext/1527209-paste?language=objc
func (x gen_NSText) Paste(
	sender objc.Ref,
) {
	C.NSText_inst_paste(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// PasteFont This action method pastes font information from the font pasteboard onto the selected text or insertion point of a rich text object, or over all text of a plain text object.
// https://developer.apple.com/documentation/appkit/nstext/1531099-pastefont?language=objc
func (x gen_NSText) PasteFont(
	sender objc.Ref,
) {
	C.NSText_inst_pasteFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// PasteRuler This action method pastes paragraph style information from the ruler pasteboard onto the selected paragraphs of a rich text object.
// https://developer.apple.com/documentation/appkit/nstext/1530420-pasteruler?language=objc
func (x gen_NSText) PasteRuler(
	sender objc.Ref,
) {
	C.NSText_inst_pasteRuler(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ReadRTFDFromFile Attempts to read the RTFD file at path, returning YES if successful and NO if not.
// https://developer.apple.com/documentation/appkit/nstext/1532564-readrtfdfromfile?language=objc
func (x gen_NSText) ReadRTFDFromFile(
	path core.NSStringRef,
) bool {
	ret := C.NSText_inst_readRTFDFromFile(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)

	return convertObjCBoolToGo(ret)

}

// SelectAll This action method selects all of the receiver’s text.
// https://developer.apple.com/documentation/appkit/nstext/1527642-selectall?language=objc
func (x gen_NSText) SelectAll(
	sender objc.Ref,
) {
	C.NSText_inst_selectAll(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ShowGuessPanel This action method opens the Spelling panel, allowing the user to make a correction during spell checking.
// https://developer.apple.com/documentation/appkit/nstext/1533456-showguesspanel?language=objc
func (x gen_NSText) ShowGuessPanel(
	sender objc.Ref,
) {
	C.NSText_inst_showGuessPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// SizeToFit Resizes the receiver to fit its text.
// https://developer.apple.com/documentation/appkit/nstext/1533991-sizetofit?language=objc
func (x gen_NSText) SizeToFit() {
	C.NSText_inst_sizeToFit(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Subscript This action method applies a subscript attribute to selected text (or all text if the receiver is a plain text object), lowering its baseline offset by a predefined amount.
// https://developer.apple.com/documentation/appkit/nstext/1535373-subscript?language=objc
func (x gen_NSText) Subscript(
	sender objc.Ref,
) {
	C.NSText_inst_subscript(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// Superscript This action method applies a superscript attribute to selected text (or all text if the receiver is a plain text object), raising its baseline offset by a predefined amount.
// https://developer.apple.com/documentation/appkit/nstext/1525743-superscript?language=objc
func (x gen_NSText) Superscript(
	sender objc.Ref,
) {
	C.NSText_inst_superscript(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleRuler This action method shows or hides the ruler, if the receiver is enclosed in a scroll view.
// https://developer.apple.com/documentation/appkit/nstext/1535773-toggleruler?language=objc
func (x gen_NSText) ToggleRuler(
	sender objc.Ref,
) {
	C.NSText_inst_toggleRuler(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// Underline Adds the underline attribute to the selected text attributes if absent; removes the attribute if present.
// https://developer.apple.com/documentation/appkit/nstext/1534203-underline?language=objc
func (x gen_NSText) Underline(
	sender objc.Ref,
) {
	C.NSText_inst_underline(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// Unscript This action method removes any superscripting or subscripting from selected text (or all text if the receiver is a plain text object).
// https://developer.apple.com/documentation/appkit/nstext/1527542-unscript?language=objc
func (x gen_NSText) Unscript(
	sender objc.Ref,
) {
	C.NSText_inst_unscript(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// WriteRTFDToFile_atomically Writes the receiver’s text as RTF with attachments to a file or directory at path.
// https://developer.apple.com/documentation/appkit/nstext/1527085-writertfdtofile?language=objc
func (x gen_NSText) WriteRTFDToFile_atomically(
	path core.NSStringRef,
	flag bool,
) bool {
	ret := C.NSText_inst_writeRTFDToFile_atomically(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(flag),
	)

	return convertObjCBoolToGo(ret)

}

// Init_asNSText
func (x gen_NSText) Init_asNSText() NSText {
	ret := C.NSText_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSText_fromPointer(ret)

}

// String The characters of the receiver’s text.
// https://developer.apple.com/documentation/appkit/nstext/1528601-string?language=objc
func (x gen_NSText) String() core.NSString {
	ret := C.NSText_inst_string(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetString The characters of the receiver’s text.
// https://developer.apple.com/documentation/appkit/nstext/1528601-string?language=objc
func (x gen_NSText) SetString(
	value core.NSStringRef,
) {
	C.NSText_inst_setString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// BackgroundColor The receiver’s background color to a given color.
// https://developer.apple.com/documentation/appkit/nstext/1535324-backgroundcolor?language=objc
func (x gen_NSText) BackgroundColor() NSColor {
	ret := C.NSText_inst_backgroundColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_fromPointer(ret)

}

// SetBackgroundColor The receiver’s background color to a given color.
// https://developer.apple.com/documentation/appkit/nstext/1535324-backgroundcolor?language=objc
func (x gen_NSText) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSText_inst_setBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// DrawsBackground A Boolean that controls whether the receiver draws its background.
// https://developer.apple.com/documentation/appkit/nstext/1531772-drawsbackground?language=objc
func (x gen_NSText) DrawsBackground() bool {
	ret := C.NSText_inst_drawsBackground(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetDrawsBackground A Boolean that controls whether the receiver draws its background.
// https://developer.apple.com/documentation/appkit/nstext/1531772-drawsbackground?language=objc
func (x gen_NSText) SetDrawsBackground(
	value bool,
) {
	C.NSText_inst_setDrawsBackground(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsEditable A Boolean that controls whether the receiver allows the user to edit its text.
// https://developer.apple.com/documentation/appkit/nstext/1529876-editable?language=objc
func (x gen_NSText) IsEditable() bool {
	ret := C.NSText_inst_isEditable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetEditable A Boolean that controls whether the receiver allows the user to edit its text.
// https://developer.apple.com/documentation/appkit/nstext/1529876-editable?language=objc
func (x gen_NSText) SetEditable(
	value bool,
) {
	C.NSText_inst_setEditable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsSelectable A Boolean that controls whether the receiver allows the user to select its text.
// https://developer.apple.com/documentation/appkit/nstext/1535368-selectable?language=objc
func (x gen_NSText) IsSelectable() bool {
	ret := C.NSText_inst_isSelectable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetSelectable A Boolean that controls whether the receiver allows the user to select its text.
// https://developer.apple.com/documentation/appkit/nstext/1535368-selectable?language=objc
func (x gen_NSText) SetSelectable(
	value bool,
) {
	C.NSText_inst_setSelectable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsFieldEditor A Boolean that controls whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder.
// https://developer.apple.com/documentation/appkit/nstext/1533080-fieldeditor?language=objc
func (x gen_NSText) IsFieldEditor() bool {
	ret := C.NSText_inst_isFieldEditor(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetFieldEditor A Boolean that controls whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder.
// https://developer.apple.com/documentation/appkit/nstext/1533080-fieldeditor?language=objc
func (x gen_NSText) SetFieldEditor(
	value bool,
) {
	C.NSText_inst_setFieldEditor(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsRichText A Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text.
// https://developer.apple.com/documentation/appkit/nstext/1531003-richtext?language=objc
func (x gen_NSText) IsRichText() bool {
	ret := C.NSText_inst_isRichText(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetRichText A Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text.
// https://developer.apple.com/documentation/appkit/nstext/1531003-richtext?language=objc
func (x gen_NSText) SetRichText(
	value bool,
) {
	C.NSText_inst_setRichText(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// ImportsGraphics A Boolean that controls whether the receiver allows the user to import files by dragging.
// https://developer.apple.com/documentation/appkit/nstext/1531887-importsgraphics?language=objc
func (x gen_NSText) ImportsGraphics() bool {
	ret := C.NSText_inst_importsGraphics(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetImportsGraphics A Boolean that controls whether the receiver allows the user to import files by dragging.
// https://developer.apple.com/documentation/appkit/nstext/1531887-importsgraphics?language=objc
func (x gen_NSText) SetImportsGraphics(
	value bool,
) {
	C.NSText_inst_setImportsGraphics(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// UsesFontPanel A Boolean that controls whether the receiver uses the Font panel and Font menu.
// https://developer.apple.com/documentation/appkit/nstext/1527431-usesfontpanel?language=objc
func (x gen_NSText) UsesFontPanel() bool {
	ret := C.NSText_inst_usesFontPanel(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetUsesFontPanel A Boolean that controls whether the receiver uses the Font panel and Font menu.
// https://developer.apple.com/documentation/appkit/nstext/1527431-usesfontpanel?language=objc
func (x gen_NSText) SetUsesFontPanel(
	value bool,
) {
	C.NSText_inst_setUsesFontPanel(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsRulerVisible A Boolean value that indicates whether the receiver’s enclosing scroll view shows its ruler.
// https://developer.apple.com/documentation/appkit/nstext/1533732-rulervisible?language=objc
func (x gen_NSText) IsRulerVisible() bool {
	ret := C.NSText_inst_isRulerVisible(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Font The font of all the receiver’s text.
// https://developer.apple.com/documentation/appkit/nstext/1534646-font?language=objc
func (x gen_NSText) Font() NSFont {
	ret := C.NSText_inst_font(
		unsafe.Pointer(x.Pointer()),
	)

	return NSFont_fromPointer(ret)

}

// SetFont The font of all the receiver’s text.
// https://developer.apple.com/documentation/appkit/nstext/1534646-font?language=objc
func (x gen_NSText) SetFont(
	value NSFontRef,
) {
	C.NSText_inst_setFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// TextColor The text color of all characters in the receiver.
// https://developer.apple.com/documentation/appkit/nstext/1534875-textcolor?language=objc
func (x gen_NSText) TextColor() NSColor {
	ret := C.NSText_inst_textColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_fromPointer(ret)

}

// SetTextColor The text color of all characters in the receiver.
// https://developer.apple.com/documentation/appkit/nstext/1534875-textcolor?language=objc
func (x gen_NSText) SetTextColor(
	value NSColorRef,
) {
	C.NSText_inst_setTextColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// MaxSize The receiver’s maximum size.
// https://developer.apple.com/documentation/appkit/nstext/1535900-maxsize?language=objc
func (x gen_NSText) MaxSize() core.NSSize {
	ret := C.NSText_inst_maxSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetMaxSize The receiver’s maximum size.
// https://developer.apple.com/documentation/appkit/nstext/1535900-maxsize?language=objc
func (x gen_NSText) SetMaxSize(
	value core.NSSize,
) {
	C.NSText_inst_setMaxSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// MinSize The receiver’s minimum size.
// https://developer.apple.com/documentation/appkit/nstext/1526222-minsize?language=objc
func (x gen_NSText) MinSize() core.NSSize {
	ret := C.NSText_inst_minSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetMinSize The receiver’s minimum size.
// https://developer.apple.com/documentation/appkit/nstext/1526222-minsize?language=objc
func (x gen_NSText) SetMinSize(
	value core.NSSize,
) {
	C.NSText_inst_setMinSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// IsVerticallyResizable A Boolean that controls whether the receiver changes its height to fit the height of its text.
// https://developer.apple.com/documentation/appkit/nstext/1535082-verticallyresizable?language=objc
func (x gen_NSText) IsVerticallyResizable() bool {
	ret := C.NSText_inst_isVerticallyResizable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetVerticallyResizable A Boolean that controls whether the receiver changes its height to fit the height of its text.
// https://developer.apple.com/documentation/appkit/nstext/1535082-verticallyresizable?language=objc
func (x gen_NSText) SetVerticallyResizable(
	value bool,
) {
	C.NSText_inst_setVerticallyResizable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsHorizontallyResizable A Boolean that controls whether the receiver changes its width to fit the width of its text.
// https://developer.apple.com/documentation/appkit/nstext/1527489-horizontallyresizable?language=objc
func (x gen_NSText) IsHorizontallyResizable() bool {
	ret := C.NSText_inst_isHorizontallyResizable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetHorizontallyResizable A Boolean that controls whether the receiver changes its width to fit the width of its text.
// https://developer.apple.com/documentation/appkit/nstext/1527489-horizontallyresizable?language=objc
func (x gen_NSText) SetHorizontallyResizable(
	value bool,
) {
	C.NSText_inst_setHorizontallyResizable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// Delegate The receiver’s delegate.
// https://developer.apple.com/documentation/appkit/nstext/1529480-delegate?language=objc
func (x gen_NSText) Delegate() objc.Object {
	ret := C.NSText_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetDelegate The receiver’s delegate.
// https://developer.apple.com/documentation/appkit/nstext/1529480-delegate?language=objc
func (x gen_NSText) SetDelegate(
	value objc.Ref,
) {
	C.NSText_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

type NSTextFieldRef interface {
	Pointer() uintptr
	Init_asNSTextField() NSTextField
}

type gen_NSTextField struct {
	NSControl
}

func NSTextField_fromPointer(ptr unsafe.Pointer) NSTextField {
	return NSTextField{gen_NSTextField{
		NSControl_fromPointer(ptr),
	}}
}

func NSTextField_fromRef(ref objc.Ref) NSTextField {
	return NSTextField_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// SelectText Ends editing in the text field and, if it’s selectable, selects the entire text content.
// https://developer.apple.com/documentation/appkit/nstextfield/1399430-selecttext?language=objc
func (x gen_NSTextField) SelectText(
	sender objc.Ref,
) {
	C.NSTextField_inst_selectText(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// TextShouldBeginEditing Requests permission to begin editing a text object.
// https://developer.apple.com/documentation/appkit/nstextfield/1399399-textshouldbeginediting?language=objc
func (x gen_NSTextField) TextShouldBeginEditing(
	textObject NSTextRef,
) bool {
	ret := C.NSTextField_inst_textShouldBeginEditing(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(textObject),
	)

	return convertObjCBoolToGo(ret)

}

// TextShouldEndEditing Performs validation on the text field’s new value.
// https://developer.apple.com/documentation/appkit/nstextfield/1399434-textshouldendediting?language=objc
func (x gen_NSTextField) TextShouldEndEditing(
	textObject NSTextRef,
) bool {
	ret := C.NSTextField_inst_textShouldEndEditing(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(textObject),
	)

	return convertObjCBoolToGo(ret)

}

// Init_asNSTextField
func (x gen_NSTextField) Init_asNSTextField() NSTextField {
	ret := C.NSTextField_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextField_fromPointer(ret)

}

// IsSelectable A Boolean value that determines whether the user can select the content of the text field.
// https://developer.apple.com/documentation/appkit/nstextfield/1399422-selectable?language=objc
func (x gen_NSTextField) IsSelectable() bool {
	ret := C.NSTextField_inst_isSelectable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetSelectable A Boolean value that determines whether the user can select the content of the text field.
// https://developer.apple.com/documentation/appkit/nstextfield/1399422-selectable?language=objc
func (x gen_NSTextField) SetSelectable(
	value bool,
) {
	C.NSTextField_inst_setSelectable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsEditable A Boolean value that controls whether the user can edit the value in the text field.
// https://developer.apple.com/documentation/appkit/nstextfield/1399407-editable?language=objc
func (x gen_NSTextField) IsEditable() bool {
	ret := C.NSTextField_inst_isEditable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetEditable A Boolean value that controls whether the user can edit the value in the text field.
// https://developer.apple.com/documentation/appkit/nstextfield/1399407-editable?language=objc
func (x gen_NSTextField) SetEditable(
	value bool,
) {
	C.NSTextField_inst_setEditable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AllowsEditingTextAttributes A Boolean value that controls whether the user can change font attributes of the text field’s string.
// https://developer.apple.com/documentation/appkit/nstextfield/1399401-allowseditingtextattributes?language=objc
func (x gen_NSTextField) AllowsEditingTextAttributes() bool {
	ret := C.NSTextField_inst_allowsEditingTextAttributes(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsEditingTextAttributes A Boolean value that controls whether the user can change font attributes of the text field’s string.
// https://developer.apple.com/documentation/appkit/nstextfield/1399401-allowseditingtextattributes?language=objc
func (x gen_NSTextField) SetAllowsEditingTextAttributes(
	value bool,
) {
	C.NSTextField_inst_setAllowsEditingTextAttributes(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// ImportsGraphics A Boolean value that controls whether the user can drag image files into the text field.
// https://developer.apple.com/documentation/appkit/nstextfield/1399428-importsgraphics?language=objc
func (x gen_NSTextField) ImportsGraphics() bool {
	ret := C.NSTextField_inst_importsGraphics(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetImportsGraphics A Boolean value that controls whether the user can drag image files into the text field.
// https://developer.apple.com/documentation/appkit/nstextfield/1399428-importsgraphics?language=objc
func (x gen_NSTextField) SetImportsGraphics(
	value bool,
) {
	C.NSTextField_inst_setImportsGraphics(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// PlaceholderString The string the text field displays when empty to help the user understand the text field’s purpose.
// https://developer.apple.com/documentation/appkit/nstextfield/1399391-placeholderstring?language=objc
func (x gen_NSTextField) PlaceholderString() core.NSString {
	ret := C.NSTextField_inst_placeholderString(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetPlaceholderString The string the text field displays when empty to help the user understand the text field’s purpose.
// https://developer.apple.com/documentation/appkit/nstextfield/1399391-placeholderstring?language=objc
func (x gen_NSTextField) SetPlaceholderString(
	value core.NSStringRef,
) {
	C.NSTextField_inst_setPlaceholderString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// PlaceholderAttributedString The attributed string the text field displays when empty to help the user understand the text field’s purpose.
// https://developer.apple.com/documentation/appkit/nstextfield/1399387-placeholderattributedstring?language=objc
func (x gen_NSTextField) PlaceholderAttributedString() core.NSAttributedString {
	ret := C.NSTextField_inst_placeholderAttributedString(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSAttributedString_fromPointer(ret)

}

// SetPlaceholderAttributedString The attributed string the text field displays when empty to help the user understand the text field’s purpose.
// https://developer.apple.com/documentation/appkit/nstextfield/1399387-placeholderattributedstring?language=objc
func (x gen_NSTextField) SetPlaceholderAttributedString(
	value core.NSAttributedStringRef,
) {
	C.NSTextField_inst_setPlaceholderAttributedString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// AllowsDefaultTighteningForTruncation A Boolean value that controls whether single-line text fields tighten intercharacter spacing before truncating the text.
// https://developer.apple.com/documentation/appkit/nstextfield/1399405-allowsdefaulttighteningfortrunca?language=objc
func (x gen_NSTextField) AllowsDefaultTighteningForTruncation() bool {
	ret := C.NSTextField_inst_allowsDefaultTighteningForTruncation(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsDefaultTighteningForTruncation A Boolean value that controls whether single-line text fields tighten intercharacter spacing before truncating the text.
// https://developer.apple.com/documentation/appkit/nstextfield/1399405-allowsdefaulttighteningfortrunca?language=objc
func (x gen_NSTextField) SetAllowsDefaultTighteningForTruncation(
	value bool,
) {
	C.NSTextField_inst_setAllowsDefaultTighteningForTruncation(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// MaximumNumberOfLines The maximum number of lines a wrapping text field displays before clipping or truncating the text.
// https://developer.apple.com/documentation/appkit/nstextfield/1399424-maximumnumberoflines?language=objc
func (x gen_NSTextField) MaximumNumberOfLines() core.NSInteger {
	ret := C.NSTextField_inst_maximumNumberOfLines(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// SetMaximumNumberOfLines The maximum number of lines a wrapping text field displays before clipping or truncating the text.
// https://developer.apple.com/documentation/appkit/nstextfield/1399424-maximumnumberoflines?language=objc
func (x gen_NSTextField) SetMaximumNumberOfLines(
	value core.NSInteger,
) {
	C.NSTextField_inst_setMaximumNumberOfLines(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return

}

// PreferredMaxLayoutWidth The maximum width of the text field’s intrinsic content size.
// https://developer.apple.com/documentation/appkit/nstextfield/1399395-preferredmaxlayoutwidth?language=objc
func (x gen_NSTextField) PreferredMaxLayoutWidth() core.CGFloat {
	ret := C.NSTextField_inst_preferredMaxLayoutWidth(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// SetPreferredMaxLayoutWidth The maximum width of the text field’s intrinsic content size.
// https://developer.apple.com/documentation/appkit/nstextfield/1399395-preferredmaxlayoutwidth?language=objc
func (x gen_NSTextField) SetPreferredMaxLayoutWidth(
	value core.CGFloat,
) {
	C.NSTextField_inst_setPreferredMaxLayoutWidth(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// TextColor The color of the text field’s content.
// https://developer.apple.com/documentation/appkit/nstextfield/1399409-textcolor?language=objc
func (x gen_NSTextField) TextColor() NSColor {
	ret := C.NSTextField_inst_textColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_fromPointer(ret)

}

// SetTextColor The color of the text field’s content.
// https://developer.apple.com/documentation/appkit/nstextfield/1399409-textcolor?language=objc
func (x gen_NSTextField) SetTextColor(
	value NSColorRef,
) {
	C.NSTextField_inst_setTextColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// BackgroundColor The color of the background the text field’s cell draws behind the text.
// https://developer.apple.com/documentation/appkit/nstextfield/1399389-backgroundcolor?language=objc
func (x gen_NSTextField) BackgroundColor() NSColor {
	ret := C.NSTextField_inst_backgroundColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_fromPointer(ret)

}

// SetBackgroundColor The color of the background the text field’s cell draws behind the text.
// https://developer.apple.com/documentation/appkit/nstextfield/1399389-backgroundcolor?language=objc
func (x gen_NSTextField) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSTextField_inst_setBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// DrawsBackground A Boolean value that controls whether the text field’s cell draws a background color behind the text.
// https://developer.apple.com/documentation/appkit/nstextfield/1399416-drawsbackground?language=objc
func (x gen_NSTextField) DrawsBackground() bool {
	ret := C.NSTextField_inst_drawsBackground(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetDrawsBackground A Boolean value that controls whether the text field’s cell draws a background color behind the text.
// https://developer.apple.com/documentation/appkit/nstextfield/1399416-drawsbackground?language=objc
func (x gen_NSTextField) SetDrawsBackground(
	value bool,
) {
	C.NSTextField_inst_setDrawsBackground(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsBezeled A Boolean value that controls whether the text field draws a bezeled background around its contents.
// https://developer.apple.com/documentation/appkit/nstextfield/1399435-bezeled?language=objc
func (x gen_NSTextField) IsBezeled() bool {
	ret := C.NSTextField_inst_isBezeled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetBezeled A Boolean value that controls whether the text field draws a bezeled background around its contents.
// https://developer.apple.com/documentation/appkit/nstextfield/1399435-bezeled?language=objc
func (x gen_NSTextField) SetBezeled(
	value bool,
) {
	C.NSTextField_inst_setBezeled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsBordered A Boolean value that controls whether the text field draws a solid black border around its contents.
// https://developer.apple.com/documentation/appkit/nstextfield/1399403-bordered?language=objc
func (x gen_NSTextField) IsBordered() bool {
	ret := C.NSTextField_inst_isBordered(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetBordered A Boolean value that controls whether the text field draws a solid black border around its contents.
// https://developer.apple.com/documentation/appkit/nstextfield/1399403-bordered?language=objc
func (x gen_NSTextField) SetBordered(
	value bool,
) {
	C.NSTextField_inst_setBordered(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AcceptsFirstResponder A Boolean value that indicates whether the text field is editable and accepts first responder status.
// https://developer.apple.com/documentation/appkit/nstextfield/1399393-acceptsfirstresponder?language=objc
func (x gen_NSTextField) AcceptsFirstResponder() bool {
	ret := C.NSTextField_inst_acceptsFirstResponder(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AllowsCharacterPickerTouchBarItem A Boolean value that controls whether the Touch Bar displays the character picker item for rich text fields.
// https://developer.apple.com/documentation/appkit/nstextfield/2539553-allowscharacterpickertouchbarite?language=objc
func (x gen_NSTextField) AllowsCharacterPickerTouchBarItem() bool {
	ret := C.NSTextField_inst_allowsCharacterPickerTouchBarItem(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsCharacterPickerTouchBarItem A Boolean value that controls whether the Touch Bar displays the character picker item for rich text fields.
// https://developer.apple.com/documentation/appkit/nstextfield/2539553-allowscharacterpickertouchbarite?language=objc
func (x gen_NSTextField) SetAllowsCharacterPickerTouchBarItem(
	value bool,
) {
	C.NSTextField_inst_setAllowsCharacterPickerTouchBarItem(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsAutomaticTextCompletionEnabled A Boolean value that indicates whether the text field automatically completes text as the user types.
// https://developer.apple.com/documentation/appkit/nstextfield/2539554-automatictextcompletionenabled?language=objc
func (x gen_NSTextField) IsAutomaticTextCompletionEnabled() bool {
	ret := C.NSTextField_inst_isAutomaticTextCompletionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAutomaticTextCompletionEnabled A Boolean value that indicates whether the text field automatically completes text as the user types.
// https://developer.apple.com/documentation/appkit/nstextfield/2539554-automatictextcompletionenabled?language=objc
func (x gen_NSTextField) SetAutomaticTextCompletionEnabled(
	value bool,
) {
	C.NSTextField_inst_setAutomaticTextCompletionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// Delegate The text field’s delegate.
// https://developer.apple.com/documentation/appkit/nstextfield/1399437-delegate?language=objc
func (x gen_NSTextField) Delegate() objc.Object {
	ret := C.NSTextField_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetDelegate The text field’s delegate.
// https://developer.apple.com/documentation/appkit/nstextfield/1399437-delegate?language=objc
func (x gen_NSTextField) SetDelegate(
	value objc.Ref,
) {
	C.NSTextField_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

type NSTextContainerRef interface {
	Pointer() uintptr
	Init_asNSTextContainer() NSTextContainer
}

type gen_NSTextContainer struct {
	objc.Object
}

func NSTextContainer_fromPointer(ptr unsafe.Pointer) NSTextContainer {
	return NSTextContainer{gen_NSTextContainer{
		objc.Object_fromPointer(ptr),
	}}
}

func NSTextContainer_fromRef(ref objc.Ref) NSTextContainer {
	return NSTextContainer_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// InitWithSize_asNSTextContainer Initializes a text container with a specified bounding rectangle.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444529-initwithsize?language=objc
func (x gen_NSTextContainer) InitWithSize_asNSTextContainer(
	size core.NSSize,
) NSTextContainer {
	ret := C.NSTextContainer_inst_initWithSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return NSTextContainer_fromPointer(ret)

}

// ReplaceLayoutManager Replaces the layout manager for the group of text system objects that contains the text container.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444545-replacelayoutmanager?language=objc
func (x gen_NSTextContainer) ReplaceLayoutManager(
	newLayoutManager NSLayoutManagerRef,
) {
	C.NSTextContainer_inst_replaceLayoutManager(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newLayoutManager),
	)

	return

}

// Init_asNSTextContainer
func (x gen_NSTextContainer) Init_asNSTextContainer() NSTextContainer {
	ret := C.NSTextContainer_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextContainer_fromPointer(ret)

}

// LayoutManager The text container’s layout manager.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444517-layoutmanager?language=objc
func (x gen_NSTextContainer) LayoutManager() NSLayoutManager {
	ret := C.NSTextContainer_inst_layoutManager(
		unsafe.Pointer(x.Pointer()),
	)

	return NSLayoutManager_fromPointer(ret)

}

// SetLayoutManager The text container’s layout manager.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444517-layoutmanager?language=objc
func (x gen_NSTextContainer) SetLayoutManager(
	value NSLayoutManagerRef,
) {
	C.NSTextContainer_inst_setLayoutManager(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// TextView The text container’s text view.
// https://developer.apple.com/documentation/appkit/nstextcontainer/1444537-textview?language=objc
func (x gen_NSTextContainer) TextView() NSTextView {
	ret := C.NSTextContainer_inst_textView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextView_fromPointer(ret)

}

// SetTextView The text container’s text view.
// https://developer.apple.com/documentation/appkit/nstextcontainer/1444537-textview?language=objc
func (x gen_NSTextContainer) SetTextView(
	value NSTextViewRef,
) {
	C.NSTextContainer_inst_setTextView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Size The size of the text container’s bounding rectangle.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444553-size?language=objc
func (x gen_NSTextContainer) Size() core.NSSize {
	ret := C.NSTextContainer_inst_size(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetSize The size of the text container’s bounding rectangle.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444553-size?language=objc
func (x gen_NSTextContainer) SetSize(
	value core.NSSize,
) {
	C.NSTextContainer_inst_setSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// ExclusionPaths An array of path objects that represents the regions where text doesn’t display in the text container.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444569-exclusionpaths?language=objc
func (x gen_NSTextContainer) ExclusionPaths() core.NSArray {
	ret := C.NSTextContainer_inst_exclusionPaths(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// SetExclusionPaths An array of path objects that represents the regions where text doesn’t display in the text container.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444569-exclusionpaths?language=objc
func (x gen_NSTextContainer) SetExclusionPaths(
	value core.NSArrayRef,
) {
	C.NSTextContainer_inst_setExclusionPaths(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// WidthTracksTextView A Boolean that controls whether the text container adjusts the width of its bounding rectangle when its text view resizes.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444563-widthtrackstextview?language=objc
func (x gen_NSTextContainer) WidthTracksTextView() bool {
	ret := C.NSTextContainer_inst_widthTracksTextView(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetWidthTracksTextView A Boolean that controls whether the text container adjusts the width of its bounding rectangle when its text view resizes.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444563-widthtrackstextview?language=objc
func (x gen_NSTextContainer) SetWidthTracksTextView(
	value bool,
) {
	C.NSTextContainer_inst_setWidthTracksTextView(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// HeightTracksTextView A Boolean that controls whether the text container adjusts the height of its bounding rectangle when its text view resizes.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444559-heighttrackstextview?language=objc
func (x gen_NSTextContainer) HeightTracksTextView() bool {
	ret := C.NSTextContainer_inst_heightTracksTextView(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetHeightTracksTextView A Boolean that controls whether the text container adjusts the height of its bounding rectangle when its text view resizes.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444559-heighttrackstextview?language=objc
func (x gen_NSTextContainer) SetHeightTracksTextView(
	value bool,
) {
	C.NSTextContainer_inst_setHeightTracksTextView(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// MaximumNumberOfLines The maximum number of lines that the text container can store.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444531-maximumnumberoflines?language=objc
func (x gen_NSTextContainer) MaximumNumberOfLines() core.NSUInteger {
	ret := C.NSTextContainer_inst_maximumNumberOfLines(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// SetMaximumNumberOfLines The maximum number of lines that the text container can store.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444531-maximumnumberoflines?language=objc
func (x gen_NSTextContainer) SetMaximumNumberOfLines(
	value core.NSUInteger,
) {
	C.NSTextContainer_inst_setMaximumNumberOfLines(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)

	return

}

// LineFragmentPadding The value for the text inset within line fragment rectangles.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444527-linefragmentpadding?language=objc
func (x gen_NSTextContainer) LineFragmentPadding() core.CGFloat {
	ret := C.NSTextContainer_inst_lineFragmentPadding(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// SetLineFragmentPadding The value for the text inset within line fragment rectangles.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444527-linefragmentpadding?language=objc
func (x gen_NSTextContainer) SetLineFragmentPadding(
	value core.CGFloat,
) {
	C.NSTextContainer_inst_setLineFragmentPadding(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// IsSimpleRectangularTextContainer A Boolean that indicates whether the text container’s region is a rectangle with no holes or gaps, and whose edges are parallel to the text view's coordinate system axes.
// https://developer.apple.com/documentation/uikit/nstextcontainer/1444525-simplerectangulartextcontainer?language=objc
func (x gen_NSTextContainer) IsSimpleRectangularTextContainer() bool {
	ret := C.NSTextContainer_inst_isSimpleRectangularTextContainer(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

type NSViewControllerRef interface {
	Pointer() uintptr
	Init_asNSViewController() NSViewController
}

type gen_NSViewController struct {
	objc.Object
}

func NSViewController_fromPointer(ptr unsafe.Pointer) NSViewController {
	return NSViewController{gen_NSViewController{
		objc.Object_fromPointer(ptr),
	}}
}

func NSViewController_fromRef(ref objc.Ref) NSViewController {
	return NSViewController_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// AddChildViewController A convenience method for adding a child view controller at the end of the childViewControllers array.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434501-addchildviewcontroller?language=objc
func (x gen_NSViewController) AddChildViewController(
	childViewController NSViewControllerRef,
) {
	C.NSViewController_inst_addChildViewController(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(childViewController),
	)

	return

}

// CommitEditing Returns whether the receiver was able to commit any pending edits.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434485-commitediting?language=objc
func (x gen_NSViewController) CommitEditing() bool {
	ret := C.NSViewController_inst_commitEditing(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// CommitEditingWithDelegate_didCommitSelector_contextInfo Attempt to commit any currently edited results of the receiver.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434464-commiteditingwithdelegate?language=objc
func (x gen_NSViewController) CommitEditingWithDelegate_didCommitSelector_contextInfo(
	delegate objc.Ref,
	didCommitSelector objc.Selector,
	contextInfo unsafe.Pointer,
) {
	C.NSViewController_inst_commitEditingWithDelegate_didCommitSelector_contextInfo(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(delegate),
		didCommitSelector.SelectorAddress(),
		contextInfo,
	)

	return

}

// DiscardEditing Causes the receiver to discard any changes, restoring the previous values.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434487-discardediting?language=objc
func (x gen_NSViewController) DiscardEditing() {
	C.NSViewController_inst_discardEditing(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// DismissController
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434447-dismisscontroller?language=objc
func (x gen_NSViewController) DismissController(
	sender objc.Ref,
) {
	C.NSViewController_inst_dismissController(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// DismissViewController Dismisses a presented view controller, using the same animator that presented it.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434413-dismissviewcontroller?language=objc
func (x gen_NSViewController) DismissViewController(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_dismissViewController(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)

	return

}

// InsertChildViewController_atIndex Inserts a specified child view controller into the childViewControllers array at a specified position.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434437-insertchildviewcontroller?language=objc
func (x gen_NSViewController) InsertChildViewController_atIndex(
	childViewController NSViewControllerRef,
	index core.NSInteger,
) {
	C.NSViewController_inst_insertChildViewController_atIndex(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(childViewController),
		C.long(index),
	)

	return

}

// LoadView Instantiates a view from a nib file and sets the value of the view property.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434405-loadview?language=objc
func (x gen_NSViewController) LoadView() {
	C.NSViewController_inst_loadView(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// PreferredContentSizeDidChangeForViewController Called when there is a change in value of the preferredContentSize property of a child view controller or a presented view controller.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434434-preferredcontentsizedidchangefor?language=objc
func (x gen_NSViewController) PreferredContentSizeDidChangeForViewController(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_preferredContentSizeDidChangeForViewController(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)

	return

}

// PresentViewController_animator Presents another view controller using a specified, custom animator for presentation and dismissal.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434431-presentviewcontroller?language=objc
func (x gen_NSViewController) PresentViewController_animator(
	viewController NSViewControllerRef,
	animator objc.Ref,
) {
	C.NSViewController_inst_presentViewController_animator(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
		objc.RefPointer(animator),
	)

	return

}

// PresentViewControllerAsModalWindow Presents another view controller as a modal window, also known as an alert.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434462-presentviewcontrollerasmodalwind?language=objc
func (x gen_NSViewController) PresentViewControllerAsModalWindow(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_presentViewControllerAsModalWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)

	return

}

// PresentViewControllerAsSheet Presents another view controller as a sheet.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434489-presentviewcontrollerassheet?language=objc
func (x gen_NSViewController) PresentViewControllerAsSheet(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_presentViewControllerAsSheet(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)

	return

}

// RemoveChildViewControllerAtIndex Removes a specified child controller from the view controller.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434404-removechildviewcontrolleratindex?language=objc
func (x gen_NSViewController) RemoveChildViewControllerAtIndex(
	index core.NSInteger,
) {
	C.NSViewController_inst_removeChildViewControllerAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)

	return

}

// RemoveFromParentViewController Removes the called view controller from its parent view controller.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434466-removefromparentviewcontroller?language=objc
func (x gen_NSViewController) RemoveFromParentViewController() {
	C.NSViewController_inst_removeFromParentViewController(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// UpdateViewConstraints Called during Auto Layout constraint updating to enable the view controller to mediate the process.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434400-updateviewconstraints?language=objc
func (x gen_NSViewController) UpdateViewConstraints() {
	C.NSViewController_inst_updateViewConstraints(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewDidAppear Called when the view controller’s view is fully transitioned onto the screen.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434455-viewdidappear?language=objc
func (x gen_NSViewController) ViewDidAppear() {
	C.NSViewController_inst_viewDidAppear(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewDidDisappear Called after the view controller’s view is removed from the view hierarchy in a window.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434416-viewdiddisappear?language=objc
func (x gen_NSViewController) ViewDidDisappear() {
	C.NSViewController_inst_viewDidDisappear(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewDidLayout Called immediately after the layout method of the view controller's view is called.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434451-viewdidlayout?language=objc
func (x gen_NSViewController) ViewDidLayout() {
	C.NSViewController_inst_viewDidLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewDidLoad Called after the view controller’s view has been loaded into memory.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434476-viewdidload?language=objc
func (x gen_NSViewController) ViewDidLoad() {
	C.NSViewController_inst_viewDidLoad(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewWillAppear Called after the view controller’s view has been loaded into memory is about to be added to the view hierarchy in the window.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434415-viewwillappear?language=objc
func (x gen_NSViewController) ViewWillAppear() {
	C.NSViewController_inst_viewWillAppear(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewWillDisappear Called when the view controller’s view is about to be removed from the view hierarchy in the window.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434483-viewwilldisappear?language=objc
func (x gen_NSViewController) ViewWillDisappear() {
	C.NSViewController_inst_viewWillDisappear(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewWillLayout Called just before the layout method of the view controller's view is called.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434495-viewwilllayout?language=objc
func (x gen_NSViewController) ViewWillLayout() {
	C.NSViewController_inst_viewWillLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewWillTransitionToSize For a view controller that is part of an app extension, called when its view is about to be resized.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434443-viewwilltransitiontosize?language=objc
func (x gen_NSViewController) ViewWillTransitionToSize(
	newSize core.NSSize,
) {
	C.NSViewController_inst_viewWillTransitionToSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&newSize)),
	)

	return

}

// Init_asNSViewController
func (x gen_NSViewController) Init_asNSViewController() NSViewController {
	ret := C.NSViewController_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSViewController_fromPointer(ret)

}

// RepresentedObject The object whose value is presented in the receiver’s primary view.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434453-representedobject?language=objc
func (x gen_NSViewController) RepresentedObject() objc.Object {
	ret := C.NSViewController_inst_representedObject(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetRepresentedObject The object whose value is presented in the receiver’s primary view.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434453-representedobject?language=objc
func (x gen_NSViewController) SetRepresentedObject(
	value objc.Ref,
) {
	C.NSViewController_inst_setRepresentedObject(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// NibBundle The nib bundle to be loaded to instantiate the receiver’s primary view.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434433-nibbundle?language=objc
func (x gen_NSViewController) NibBundle() NSBundle {
	ret := C.NSViewController_inst_nibBundle(
		unsafe.Pointer(x.Pointer()),
	)

	return NSBundle_fromPointer(ret)

}

// View The view controller’s primary view.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434401-view?language=objc
func (x gen_NSViewController) View() NSView {
	ret := C.NSViewController_inst_view(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_fromPointer(ret)

}

// SetView The view controller’s primary view.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434401-view?language=objc
func (x gen_NSViewController) SetView(
	value NSViewRef,
) {
	C.NSViewController_inst_setView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Title The localized title of the receiver’s primary view.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434426-title?language=objc
func (x gen_NSViewController) Title() core.NSString {
	ret := C.NSViewController_inst_title(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetTitle The localized title of the receiver’s primary view.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434426-title?language=objc
func (x gen_NSViewController) SetTitle(
	value core.NSStringRef,
) {
	C.NSViewController_inst_setTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// IsViewLoaded A Boolean value indicating whether the view controller’s view is loaded into memory.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434435-viewloaded?language=objc
func (x gen_NSViewController) IsViewLoaded() bool {
	ret := C.NSViewController_inst_isViewLoaded(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// PreferredContentSize The desired size of the view controller’s view, in screen units.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434409-preferredcontentsize?language=objc
func (x gen_NSViewController) PreferredContentSize() core.NSSize {
	ret := C.NSViewController_inst_preferredContentSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetPreferredContentSize The desired size of the view controller’s view, in screen units.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434409-preferredcontentsize?language=objc
func (x gen_NSViewController) SetPreferredContentSize(
	value core.NSSize,
) {
	C.NSViewController_inst_setPreferredContentSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// ChildViewControllers An array of view controllers that are hierarchical children of the view controller.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434432-childviewcontrollers?language=objc
func (x gen_NSViewController) ChildViewControllers() core.NSArray {
	ret := C.NSViewController_inst_childViewControllers(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// SetChildViewControllers An array of view controllers that are hierarchical children of the view controller.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434432-childviewcontrollers?language=objc
func (x gen_NSViewController) SetChildViewControllers(
	value core.NSArrayRef,
) {
	C.NSViewController_inst_setChildViewControllers(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// ParentViewController The immediate ancestor view controller of the view controller.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434491-parentviewcontroller?language=objc
func (x gen_NSViewController) ParentViewController() NSViewController {
	ret := C.NSViewController_inst_parentViewController(
		unsafe.Pointer(x.Pointer()),
	)

	return NSViewController_fromPointer(ret)

}

// PresentedViewControllers The view controllers, if any, that are currently presented by the view controller.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434497-presentedviewcontrollers?language=objc
func (x gen_NSViewController) PresentedViewControllers() core.NSArray {
	ret := C.NSViewController_inst_presentedViewControllers(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// PresentingViewController The view controller that presented the view controller or that presented its farthest ancestor view controller.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434439-presentingviewcontroller?language=objc
func (x gen_NSViewController) PresentingViewController() NSViewController {
	ret := C.NSViewController_inst_presentingViewController(
		unsafe.Pointer(x.Pointer()),
	)

	return NSViewController_fromPointer(ret)

}

// PreferredScreenOrigin For a view controller that is part of an app extension, the preferred screen origin.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434468-preferredscreenorigin?language=objc
func (x gen_NSViewController) PreferredScreenOrigin() core.NSPoint {
	ret := C.NSViewController_inst_preferredScreenOrigin(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// SetPreferredScreenOrigin For a view controller that is part of an app extension, the preferred screen origin.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434468-preferredscreenorigin?language=objc
func (x gen_NSViewController) SetPreferredScreenOrigin(
	value core.NSPoint,
) {
	C.NSViewController_inst_setPreferredScreenOrigin(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&value)),
	)

	return

}

// PreferredMaximumSize For a view controller that is part of an app extension, the largest allowable size for the app extension’s primary view, in screen units.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434403-preferredmaximumsize?language=objc
func (x gen_NSViewController) PreferredMaximumSize() core.NSSize {
	ret := C.NSViewController_inst_preferredMaximumSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// PreferredMinimumSize For a view controller that is part of an app extension, the smallest allowable size for the app extension’s primary view, in screen units.
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434418-preferredminimumsize?language=objc
func (x gen_NSViewController) PreferredMinimumSize() core.NSSize {
	ret := C.NSViewController_inst_preferredMinimumSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SourceItemView
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434479-sourceitemview?language=objc
func (x gen_NSViewController) SourceItemView() NSView {
	ret := C.NSViewController_inst_sourceItemView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_fromPointer(ret)

}

// SetSourceItemView
// https://developer.apple.com/documentation/appkit/nsviewcontroller/1434479-sourceitemview?language=objc
func (x gen_NSViewController) SetSourceItemView(
	value NSViewRef,
) {
	C.NSViewController_inst_setSourceItemView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

type NSVisualEffectViewRef interface {
	Pointer() uintptr
	Init_asNSVisualEffectView() NSVisualEffectView
}

type gen_NSVisualEffectView struct {
	NSView
}

func NSVisualEffectView_fromPointer(ptr unsafe.Pointer) NSVisualEffectView {
	return NSVisualEffectView{gen_NSVisualEffectView{
		NSView_fromPointer(ptr),
	}}
}

func NSVisualEffectView_fromRef(ref objc.Ref) NSVisualEffectView {
	return NSVisualEffectView_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// ViewDidMoveToWindow Notifies the view that it moved to a new window.
// https://developer.apple.com/documentation/appkit/nsvisualeffectview/1534300-viewdidmovetowindow?language=objc
func (x gen_NSVisualEffectView) ViewDidMoveToWindow() {
	C.NSVisualEffectView_inst_viewDidMoveToWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewWillMoveToWindow Notifies the view immediately before it moves to a new window (which may be nil).
// https://developer.apple.com/documentation/appkit/nsvisualeffectview/1534276-viewwillmovetowindow?language=objc
func (x gen_NSVisualEffectView) ViewWillMoveToWindow(
	newWindow NSWindowRef,
) {
	C.NSVisualEffectView_inst_viewWillMoveToWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newWindow),
	)

	return

}

// Init_asNSVisualEffectView
func (x gen_NSVisualEffectView) Init_asNSVisualEffectView() NSVisualEffectView {
	ret := C.NSVisualEffectView_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSVisualEffectView_fromPointer(ret)

}

// IsEmphasized A Boolean value indicating whether to emphasize the look of the material.
// https://developer.apple.com/documentation/appkit/nsvisualeffectview/1644721-emphasized?language=objc
func (x gen_NSVisualEffectView) IsEmphasized() bool {
	ret := C.NSVisualEffectView_inst_isEmphasized(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetEmphasized A Boolean value indicating whether to emphasize the look of the material.
// https://developer.apple.com/documentation/appkit/nsvisualeffectview/1644721-emphasized?language=objc
func (x gen_NSVisualEffectView) SetEmphasized(
	value bool,
) {
	C.NSVisualEffectView_inst_setEmphasized(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// MaskImage An image whose alpha channel masks the visual effect view's material.
// https://developer.apple.com/documentation/appkit/nsvisualeffectview/1535318-maskimage?language=objc
func (x gen_NSVisualEffectView) MaskImage() NSImage {
	ret := C.NSVisualEffectView_inst_maskImage(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_fromPointer(ret)

}

// SetMaskImage An image whose alpha channel masks the visual effect view's material.
// https://developer.apple.com/documentation/appkit/nsvisualeffectview/1535318-maskimage?language=objc
func (x gen_NSVisualEffectView) SetMaskImage(
	value NSImageRef,
) {
	C.NSVisualEffectView_inst_setMaskImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

type NSWindowRef interface {
	Pointer() uintptr
	Init_asNSWindow() NSWindow
}

type gen_NSWindow struct {
	objc.Object
}

func NSWindow_fromPointer(ptr unsafe.Pointer) NSWindow {
	return NSWindow{gen_NSWindow{
		objc.Object_fromPointer(ptr),
	}}
}

func NSWindow_fromRef(ref objc.Ref) NSWindow {
	return NSWindow_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// AddChildWindow_ordered Adds a given window as a child window of the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419152-addchildwindow?language=objc
func (x gen_NSWindow) AddChildWindow_ordered(
	childWin NSWindowRef,
	place core.NSUInteger,
) {
	C.NSWindow_inst_addChildWindow_ordered(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(childWin),
		C.ulong(place),
	)

	return

}

// AddTabbedWindow_ordered Adds the provided window as a new tab in a tabbed window using the specified ordering instruction.
// https://developer.apple.com/documentation/appkit/nswindow/1855947-addtabbedwindow?language=objc
func (x gen_NSWindow) AddTabbedWindow_ordered(
	window NSWindowRef,
	ordered core.NSUInteger,
) {
	C.NSWindow_inst_addTabbedWindow_ordered(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(window),
		C.ulong(ordered),
	)

	return

}

// BecomeKeyWindow Informs the window that it has become the key window.
// https://developer.apple.com/documentation/appkit/nswindow/1419338-becomekeywindow?language=objc
func (x gen_NSWindow) BecomeKeyWindow() {
	C.NSWindow_inst_becomeKeyWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// BecomeMainWindow Informs the window that it has become the main window.
// https://developer.apple.com/documentation/appkit/nswindow/1419084-becomemainwindow?language=objc
func (x gen_NSWindow) BecomeMainWindow() {
	C.NSWindow_inst_becomeMainWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// CascadeTopLeftFromPoint Positions the window’s top-left to a given point.
// https://developer.apple.com/documentation/appkit/nswindow/1419392-cascadetopleftfrompoint?language=objc
func (x gen_NSWindow) CascadeTopLeftFromPoint(
	topLeftPoint core.NSPoint,
) core.NSPoint {
	ret := C.NSWindow_inst_cascadeTopLeftFromPoint(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&topLeftPoint)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// Center Sets the window’s location to the center of the screen.
// https://developer.apple.com/documentation/appkit/nswindow/1419090-center?language=objc
func (x gen_NSWindow) Center() {
	C.NSWindow_inst_center(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Close Removes the window from the screen.
// https://developer.apple.com/documentation/appkit/nswindow/1419662-close?language=objc
func (x gen_NSWindow) Close() {
	C.NSWindow_inst_close(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ConstrainFrameRect_toScreen Modifies and returns a frame rectangle so that its top edge lies on a specific screen.
// https://developer.apple.com/documentation/appkit/nswindow/1419779-constrainframerect?language=objc
func (x gen_NSWindow) ConstrainFrameRect_toScreen(
	frameRect core.NSRect,
	screen NSScreenRef,
) core.NSRect {
	ret := C.NSWindow_inst_constrainFrameRect_toScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
		objc.RefPointer(screen),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// ContentRectForFrameRect Returns the window’s content rectangle with a given frame rectangle.
// https://developer.apple.com/documentation/appkit/nswindow/1419108-contentrectforframerect?language=objc
func (x gen_NSWindow) ContentRectForFrameRect(
	frameRect core.NSRect,
) core.NSRect {
	ret := C.NSWindow_inst_contentRectForFrameRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// ConvertPointFromBacking Converts a point from its pixel-aligned backing store coordinate system to the window’s coordinate system.
// https://developer.apple.com/documentation/appkit/nswindow/2967179-convertpointfrombacking?language=objc
func (x gen_NSWindow) ConvertPointFromBacking(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSWindow_inst_convertPointFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// ConvertPointFromScreen Converts a point from the screen coordinate system to the window’s coordinate system.
// https://developer.apple.com/documentation/appkit/nswindow/2967180-convertpointfromscreen?language=objc
func (x gen_NSWindow) ConvertPointFromScreen(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSWindow_inst_convertPointFromScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// ConvertPointToBacking Converts a point from the window’s coordinate system to its pixel-aligned backing store coordinate system.
// https://developer.apple.com/documentation/appkit/nswindow/2967181-convertpointtobacking?language=objc
func (x gen_NSWindow) ConvertPointToBacking(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSWindow_inst_convertPointToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// ConvertPointToScreen Converts a point to the screen coordinate system from the window’s coordinate system.
// https://developer.apple.com/documentation/appkit/nswindow/2967182-convertpointtoscreen?language=objc
func (x gen_NSWindow) ConvertPointToScreen(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSWindow_inst_convertPointToScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// ConvertRectFromBacking Converts a rectangle from its pixel-aligned backing store coordinate system to the window’s coordinate system.
// https://developer.apple.com/documentation/appkit/nswindow/1419273-convertrectfrombacking?language=objc
func (x gen_NSWindow) ConvertRectFromBacking(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSWindow_inst_convertRectFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// ConvertRectFromScreen Converts a rectangle from the screen coordinate system to the window’s coordinate system.
// https://developer.apple.com/documentation/appkit/nswindow/1419603-convertrectfromscreen?language=objc
func (x gen_NSWindow) ConvertRectFromScreen(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSWindow_inst_convertRectFromScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// ConvertRectToBacking Converts a rectangle from the window’s coordinate system to its pixel-aligned backing store coordinate system.
// https://developer.apple.com/documentation/appkit/nswindow/1419260-convertrecttobacking?language=objc
func (x gen_NSWindow) ConvertRectToBacking(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSWindow_inst_convertRectToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// ConvertRectToScreen Converts a rectangle to the screen coordinate system from the window’s coordinate system.
// https://developer.apple.com/documentation/appkit/nswindow/1419286-convertrecttoscreen?language=objc
func (x gen_NSWindow) ConvertRectToScreen(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSWindow_inst_convertRectToScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// DataWithEPSInsideRect Returns EPS data that draws the region of the window within a given rectangle.
// https://developer.apple.com/documentation/appkit/nswindow/1419128-datawithepsinsiderect?language=objc
func (x gen_NSWindow) DataWithEPSInsideRect(
	rect core.NSRect,
) core.NSData {
	ret := C.NSWindow_inst_dataWithEPSInsideRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return core.NSData_fromPointer(ret)

}

// DataWithPDFInsideRect Returns PDF data that draws the region of the window within a given rectangle.
// https://developer.apple.com/documentation/appkit/nswindow/1419418-datawithpdfinsiderect?language=objc
func (x gen_NSWindow) DataWithPDFInsideRect(
	rect core.NSRect,
) core.NSData {
	ret := C.NSWindow_inst_dataWithPDFInsideRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return core.NSData_fromPointer(ret)

}

// Deminiaturize De-minimizes the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419334-deminiaturize?language=objc
func (x gen_NSWindow) Deminiaturize(
	sender objc.Ref,
) {
	C.NSWindow_inst_deminiaturize(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// DisableCursorRects Disables all cursor rectangle management within the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419639-disablecursorrects?language=objc
func (x gen_NSWindow) DisableCursorRects() {
	C.NSWindow_inst_disableCursorRects(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// DisableKeyEquivalentForDefaultButtonCell Disables the default button cell’s key equivalent, so it doesn’t perform a click when the user presses Return (or Enter).
// https://developer.apple.com/documentation/appkit/nswindow/1419242-disablekeyequivalentfordefaultbu?language=objc
func (x gen_NSWindow) DisableKeyEquivalentForDefaultButtonCell() {
	C.NSWindow_inst_disableKeyEquivalentForDefaultButtonCell(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// DisableScreenUpdatesUntilFlush Disables the window’s screen updates until the window is flushed.
// https://developer.apple.com/documentation/appkit/nswindow/1419483-disablescreenupdatesuntilflush?language=objc
func (x gen_NSWindow) DisableScreenUpdatesUntilFlush() {
	C.NSWindow_inst_disableScreenUpdatesUntilFlush(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// DisableSnapshotRestoration Disables snapshot restoration.
// https://developer.apple.com/documentation/appkit/nswindow/1526239-disablesnapshotrestoration?language=objc
func (x gen_NSWindow) DisableSnapshotRestoration() {
	C.NSWindow_inst_disableSnapshotRestoration(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// DiscardCursorRects Invalidates all cursor rectangles in the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419269-discardcursorrects?language=objc
func (x gen_NSWindow) DiscardCursorRects() {
	C.NSWindow_inst_discardCursorRects(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Display Passes a display message down the window’s view hierarchy, thus redrawing all views within the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419358-display?language=objc
func (x gen_NSWindow) Display() {
	C.NSWindow_inst_display(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// DisplayIfNeeded Passes a display message down the window’s view hierarchy, thus redrawing all views that need displaying.
// https://developer.apple.com/documentation/appkit/nswindow/1419096-displayifneeded?language=objc
func (x gen_NSWindow) DisplayIfNeeded() {
	C.NSWindow_inst_displayIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// DragImage_at_offset_event_pasteboard_source_slideBack Begins a dragging session.
// https://developer.apple.com/documentation/appkit/nswindow/1419224-dragimage?language=objc
func (x gen_NSWindow) DragImage_at_offset_event_pasteboard_source_slideBack(
	image NSImageRef,
	baseLocation core.NSPoint,
	initialOffset core.NSSize,
	event NSEventRef,
	pboard NSPasteboardRef,
	sourceObj objc.Ref,
	slideFlag bool,
) {
	C.NSWindow_inst_dragImage_at_offset_event_pasteboard_source_slideBack(
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

// EnableCursorRects Reenables cursor rectangle management within the window after a disableCursorRects message.
// https://developer.apple.com/documentation/appkit/nswindow/1419202-enablecursorrects?language=objc
func (x gen_NSWindow) EnableCursorRects() {
	C.NSWindow_inst_enableCursorRects(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// EnableKeyEquivalentForDefaultButtonCell Reenables the default button cell’s key equivalent, so it performs a click when the user presses Return (or Enter).
// https://developer.apple.com/documentation/appkit/nswindow/1419276-enablekeyequivalentfordefaultbut?language=objc
func (x gen_NSWindow) EnableKeyEquivalentForDefaultButtonCell() {
	C.NSWindow_inst_enableKeyEquivalentForDefaultButtonCell(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// EnableSnapshotRestoration Enables snapshot restoration.
// https://developer.apple.com/documentation/appkit/nswindow/1525288-enablesnapshotrestoration?language=objc
func (x gen_NSWindow) EnableSnapshotRestoration() {
	C.NSWindow_inst_enableSnapshotRestoration(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// EndEditingFor Forces the field editor to give up its first responder status and prepares it for its next assignment.
// https://developer.apple.com/documentation/appkit/nswindow/1419469-endeditingfor?language=objc
func (x gen_NSWindow) EndEditingFor(
	object objc.Ref,
) {
	C.NSWindow_inst_endEditingFor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)

	return

}

// EndSheet Ends a document-modal session and dismisses the specified sheet.
// https://developer.apple.com/documentation/appkit/nswindow/1419318-endsheet?language=objc
func (x gen_NSWindow) EndSheet(
	sheetWindow NSWindowRef,
) {
	C.NSWindow_inst_endSheet(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sheetWindow),
	)

	return

}

// FieldEditor_forObject Returns the window’s field editor, creating it if requested.
// https://developer.apple.com/documentation/appkit/nswindow/1419647-fieldeditor?language=objc
func (x gen_NSWindow) FieldEditor_forObject(
	createFlag bool,
	object objc.Ref,
) NSText {
	ret := C.NSWindow_inst_fieldEditor_forObject(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(createFlag),
		objc.RefPointer(object),
	)

	return NSText_fromPointer(ret)

}

// FrameRectForContentRect Returns the window’s frame rectangle with a given content rectangle.
// https://developer.apple.com/documentation/appkit/nswindow/1419134-framerectforcontentrect?language=objc
func (x gen_NSWindow) FrameRectForContentRect(
	contentRect core.NSRect,
) core.NSRect {
	ret := C.NSWindow_inst_frameRectForContentRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentRect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// InitWithContentRect_styleMask_backing_defer_asNSWindow Initializes the window with the specified values.
// https://developer.apple.com/documentation/appkit/nswindow/1419477-initwithcontentrect?language=objc
func (x gen_NSWindow) InitWithContentRect_styleMask_backing_defer_asNSWindow(
	contentRect core.NSRect,
	style core.NSUInteger,
	backingStoreType core.NSUInteger,
	flag bool,
) NSWindow {
	ret := C.NSWindow_inst_initWithContentRect_styleMask_backing_defer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentRect)),
		C.ulong(style),
		C.ulong(backingStoreType),
		convertToObjCBool(flag),
	)

	return NSWindow_fromPointer(ret)

}

// InitWithContentRect_styleMask_backing_defer_screen_asNSWindow Initializes an allocated window with the specified values.
// https://developer.apple.com/documentation/appkit/nswindow/1419755-initwithcontentrect?language=objc
func (x gen_NSWindow) InitWithContentRect_styleMask_backing_defer_screen_asNSWindow(
	contentRect core.NSRect,
	style core.NSUInteger,
	backingStoreType core.NSUInteger,
	flag bool,
	screen NSScreenRef,
) NSWindow {
	ret := C.NSWindow_inst_initWithContentRect_styleMask_backing_defer_screen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentRect)),
		C.ulong(style),
		C.ulong(backingStoreType),
		convertToObjCBool(flag),
		objc.RefPointer(screen),
	)

	return NSWindow_fromPointer(ret)

}

// InvalidateCursorRectsForView Marks as invalid the cursor rectangles of a given view object in the window, so they’ll be set up again when the window becomes key.
// https://developer.apple.com/documentation/appkit/nswindow/1419601-invalidatecursorrectsforview?language=objc
func (x gen_NSWindow) InvalidateCursorRectsForView(
	view NSViewRef,
) {
	C.NSWindow_inst_invalidateCursorRectsForView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)

	return

}

// InvalidateShadow Invalidates the window shadow so that it is recomputed based on the current window shape.
// https://developer.apple.com/documentation/appkit/nswindow/1419529-invalidateshadow?language=objc
func (x gen_NSWindow) InvalidateShadow() {
	C.NSWindow_inst_invalidateShadow(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// LayoutIfNeeded Updates the layout of views in the window based on the current views and constraints.
// https://developer.apple.com/documentation/appkit/nswindow/1526910-layoutifneeded?language=objc
func (x gen_NSWindow) LayoutIfNeeded() {
	C.NSWindow_inst_layoutIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// MakeKeyAndOrderFront Moves the window to the front of the screen list, within its level, and makes it the key window; that is, it shows the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419208-makekeyandorderfront?language=objc
func (x gen_NSWindow) MakeKeyAndOrderFront(
	sender objc.Ref,
) {
	C.NSWindow_inst_makeKeyAndOrderFront(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// MakeKeyWindow Makes the window the key window.
// https://developer.apple.com/documentation/appkit/nswindow/1419368-makekeywindow?language=objc
func (x gen_NSWindow) MakeKeyWindow() {
	C.NSWindow_inst_makeKeyWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// MakeMainWindow Makes the window the main window.
// https://developer.apple.com/documentation/appkit/nswindow/1419271-makemainwindow?language=objc
func (x gen_NSWindow) MakeMainWindow() {
	C.NSWindow_inst_makeMainWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// MergeAllWindows Merges all open windows into a single tabbed window.
// https://developer.apple.com/documentation/appkit/nswindow/1644639-mergeallwindows?language=objc
func (x gen_NSWindow) MergeAllWindows(
	sender objc.Ref,
) {
	C.NSWindow_inst_mergeAllWindows(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// Miniaturize Removes the window from the screen list and displays the minimized window in the Dock.
// https://developer.apple.com/documentation/appkit/nswindow/1419426-miniaturize?language=objc
func (x gen_NSWindow) Miniaturize(
	sender objc.Ref,
) {
	C.NSWindow_inst_miniaturize(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// MoveTabToNewWindow Moves the tab to a new containing window.
// https://developer.apple.com/documentation/appkit/nswindow/1644410-movetabtonewwindow?language=objc
func (x gen_NSWindow) MoveTabToNewWindow(
	sender objc.Ref,
) {
	C.NSWindow_inst_moveTabToNewWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// OrderBack Moves the window to the back of its level in the screen list, without changing either the key window or the main window.
// https://developer.apple.com/documentation/appkit/nswindow/1419204-orderback?language=objc
func (x gen_NSWindow) OrderBack(
	sender objc.Ref,
) {
	C.NSWindow_inst_orderBack(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// OrderFront Moves the window to the front of its level in the screen list, without changing either the key window or the main window.
// https://developer.apple.com/documentation/appkit/nswindow/1419495-orderfront?language=objc
func (x gen_NSWindow) OrderFront(
	sender objc.Ref,
) {
	C.NSWindow_inst_orderFront(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// OrderFrontRegardless Moves the window to the front of its level, even if its application isn’t active, without changing either the key window or the main window.
// https://developer.apple.com/documentation/appkit/nswindow/1419444-orderfrontregardless?language=objc
func (x gen_NSWindow) OrderFrontRegardless() {
	C.NSWindow_inst_orderFrontRegardless(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// OrderOut Removes the window from the screen list, which hides the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419660-orderout?language=objc
func (x gen_NSWindow) OrderOut(
	sender objc.Ref,
) {
	C.NSWindow_inst_orderOut(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// OrderWindow_relativeTo Repositions the window’s window device in the window server’s screen list.
// https://developer.apple.com/documentation/appkit/nswindow/1419672-orderwindow?language=objc
func (x gen_NSWindow) OrderWindow_relativeTo(
	place core.NSUInteger,
	otherWin core.NSInteger,
) {
	C.NSWindow_inst_orderWindow_relativeTo(
		unsafe.Pointer(x.Pointer()),
		C.ulong(place),
		C.long(otherWin),
	)

	return

}

// PerformClose Simulates the user clicking the close button by momentarily highlighting the button and then closing the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419288-performclose?language=objc
func (x gen_NSWindow) PerformClose(
	sender objc.Ref,
) {
	C.NSWindow_inst_performClose(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// PerformMiniaturize Simulates the user clicking the minimize button by momentarily highlighting the button, then minimizing the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419749-performminiaturize?language=objc
func (x gen_NSWindow) PerformMiniaturize(
	sender objc.Ref,
) {
	C.NSWindow_inst_performMiniaturize(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// PerformWindowDragWithEvent Starts a window drag based on the specified mouse-down event.
// https://developer.apple.com/documentation/appkit/nswindow/1419386-performwindowdragwithevent?language=objc
func (x gen_NSWindow) PerformWindowDragWithEvent(
	event NSEventRef,
) {
	C.NSWindow_inst_performWindowDragWithEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return

}

// PerformZoom This action method simulates the user clicking the zoom box by momentarily highlighting the button and then zooming the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419450-performzoom?language=objc
func (x gen_NSWindow) PerformZoom(
	sender objc.Ref,
) {
	C.NSWindow_inst_performZoom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// PostEvent_atStart Forwards the message to the global application object.
// https://developer.apple.com/documentation/appkit/nswindow/1419376-postevent?language=objc
func (x gen_NSWindow) PostEvent_atStart(
	event NSEventRef,
	flag bool,
) {
	C.NSWindow_inst_postEvent_atStart(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
		convertToObjCBool(flag),
	)

	return

}

// Print Runs the Print panel, and if the user chooses an option other than canceling, prints the window (its frame view and all subviews).
// https://developer.apple.com/documentation/appkit/nswindow/1419767-print?language=objc
func (x gen_NSWindow) Print(
	sender objc.Ref,
) {
	C.NSWindow_inst_print(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// RecalculateKeyViewLoop Marks the key view loop as “dirty” and in need of recalculation.
// https://developer.apple.com/documentation/appkit/nswindow/1419350-recalculatekeyviewloop?language=objc
func (x gen_NSWindow) RecalculateKeyViewLoop() {
	C.NSWindow_inst_recalculateKeyViewLoop(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// RegisterForDraggedTypes Registers a set of pasteboard types that the window accepts as the destination of an image-dragging session.
// https://developer.apple.com/documentation/appkit/nswindow/1419140-registerfordraggedtypes?language=objc
func (x gen_NSWindow) RegisterForDraggedTypes(
	newTypes core.NSArrayRef,
) {
	C.NSWindow_inst_registerForDraggedTypes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newTypes),
	)

	return

}

// RemoveChildWindow Detaches a given child window from the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419063-removechildwindow?language=objc
func (x gen_NSWindow) RemoveChildWindow(
	childWin NSWindowRef,
) {
	C.NSWindow_inst_removeChildWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(childWin),
	)

	return

}

// RemoveTitlebarAccessoryViewControllerAtIndex Removes the view controller at the specified index from the window’s array of title bar accessory view controllers.
// https://developer.apple.com/documentation/appkit/nswindow/1419643-removetitlebaraccessoryviewcontr?language=objc
func (x gen_NSWindow) RemoveTitlebarAccessoryViewControllerAtIndex(
	index core.NSInteger,
) {
	C.NSWindow_inst_removeTitlebarAccessoryViewControllerAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)

	return

}

// ResetCursorRects Clears the window’s cursor rectangles and the cursor rectangles of the NSView objects in its view hierarchy.
// https://developer.apple.com/documentation/appkit/nswindow/1419464-resetcursorrects?language=objc
func (x gen_NSWindow) ResetCursorRects() {
	C.NSWindow_inst_resetCursorRects(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ResignKeyWindow Resigns the window’s key window status.
// https://developer.apple.com/documentation/appkit/nswindow/1419047-resignkeywindow?language=objc
func (x gen_NSWindow) ResignKeyWindow() {
	C.NSWindow_inst_resignKeyWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ResignMainWindow Resigns the window’s main window status.
// https://developer.apple.com/documentation/appkit/nswindow/1419212-resignmainwindow?language=objc
func (x gen_NSWindow) ResignMainWindow() {
	C.NSWindow_inst_resignMainWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// RunToolbarCustomizationPalette Presents the toolbar customization user interface.
// https://developer.apple.com/documentation/appkit/nswindow/1419284-runtoolbarcustomizationpalette?language=objc
func (x gen_NSWindow) RunToolbarCustomizationPalette(
	sender objc.Ref,
) {
	C.NSWindow_inst_runToolbarCustomizationPalette(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// SelectKeyViewFollowingView Gives key view status to the view that follows the given view.
// https://developer.apple.com/documentation/appkit/nswindow/1419633-selectkeyviewfollowingview?language=objc
func (x gen_NSWindow) SelectKeyViewFollowingView(
	view NSViewRef,
) {
	C.NSWindow_inst_selectKeyViewFollowingView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)

	return

}

// SelectKeyViewPrecedingView Gives key view status to the view that precedes the given view.
// https://developer.apple.com/documentation/appkit/nswindow/1419757-selectkeyviewprecedingview?language=objc
func (x gen_NSWindow) SelectKeyViewPrecedingView(
	view NSViewRef,
) {
	C.NSWindow_inst_selectKeyViewPrecedingView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)

	return

}

// SelectNextKeyView Searches for a candidate next key view and, if it finds one, tries to make it the first responder.
// https://developer.apple.com/documentation/appkit/nswindow/1419715-selectnextkeyview?language=objc
func (x gen_NSWindow) SelectNextKeyView(
	sender objc.Ref,
) {
	C.NSWindow_inst_selectNextKeyView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// SelectNextTab Selects the next tab in the tab group in the trailing direction.
// https://developer.apple.com/documentation/appkit/nswindow/1644693-selectnexttab?language=objc
func (x gen_NSWindow) SelectNextTab(
	sender objc.Ref,
) {
	C.NSWindow_inst_selectNextTab(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// SelectPreviousKeyView Searches for a candidate previous key view and, if it finds one, tries to make it the first responder.
// https://developer.apple.com/documentation/appkit/nswindow/1419110-selectpreviouskeyview?language=objc
func (x gen_NSWindow) SelectPreviousKeyView(
	sender objc.Ref,
) {
	C.NSWindow_inst_selectPreviousKeyView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// SelectPreviousTab Selects the previous tab in the tab group in the leading direction.
// https://developer.apple.com/documentation/appkit/nswindow/1644555-selectprevioustab?language=objc
func (x gen_NSWindow) SelectPreviousTab(
	sender objc.Ref,
) {
	C.NSWindow_inst_selectPreviousTab(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// SendEvent This action method dispatches mouse and keyboard events the global application object sends to the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419228-sendevent?language=objc
func (x gen_NSWindow) SendEvent(
	event NSEventRef,
) {
	C.NSWindow_inst_sendEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return

}

// SetContentSize Sets the size of the window’s content view to a given size, which is expressed in the window’s base coordinate system.
// https://developer.apple.com/documentation/appkit/nswindow/1419100-setcontentsize?language=objc
func (x gen_NSWindow) SetContentSize(
	size core.NSSize,
) {
	C.NSWindow_inst_setContentSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return

}

// SetDynamicDepthLimit Sets a Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on.
// https://developer.apple.com/documentation/appkit/nswindow/1419473-setdynamicdepthlimit?language=objc
func (x gen_NSWindow) SetDynamicDepthLimit(
	flag bool,
) {
	C.NSWindow_inst_setDynamicDepthLimit(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)

	return

}

// SetFrame_display Sets the origin and size of the window’s frame rectangle according to a given frame rectangle, thereby setting its position and size onscreen.
// https://developer.apple.com/documentation/appkit/nswindow/1419753-setframe?language=objc
func (x gen_NSWindow) SetFrame_display(
	frameRect core.NSRect,
	flag bool,
) {
	C.NSWindow_inst_setFrame_display(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
		convertToObjCBool(flag),
	)

	return

}

// SetFrame_display_animate Sets the origin and size of the window’s frame rectangle, with optional animation, according to a given frame rectangle, thereby setting its position and size onscreen.
// https://developer.apple.com/documentation/appkit/nswindow/1419519-setframe?language=objc
func (x gen_NSWindow) SetFrame_display_animate(
	frameRect core.NSRect,
	displayFlag bool,
	animateFlag bool,
) {
	C.NSWindow_inst_setFrame_display_animate(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
		convertToObjCBool(displayFlag),
		convertToObjCBool(animateFlag),
	)

	return

}

// SetFrameOrigin Positions the bottom-left corner of the window’s frame rectangle at a given point in screen coordinates.
// https://developer.apple.com/documentation/appkit/nswindow/1419690-setframeorigin?language=objc
func (x gen_NSWindow) SetFrameOrigin(
	point core.NSPoint,
) {
	C.NSWindow_inst_setFrameOrigin(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return

}

// SetFrameTopLeftPoint Positions the top-left corner of the window’s frame rectangle at a given point in screen coordinates.
// https://developer.apple.com/documentation/appkit/nswindow/1419658-setframetopleftpoint?language=objc
func (x gen_NSWindow) SetFrameTopLeftPoint(
	point core.NSPoint,
) {
	C.NSWindow_inst_setFrameTopLeftPoint(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return

}

// SetIsMiniaturized Sets the window’s miniaturized state to the value you specify.
// https://developer.apple.com/documentation/appkit/nswindow/1449566-setisminiaturized?language=objc
func (x gen_NSWindow) SetIsMiniaturized(
	flag bool,
) {
	C.NSWindow_inst_setIsMiniaturized(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)

	return

}

// SetIsVisible Sets the window’s visible state to the value you specify.
// https://developer.apple.com/documentation/appkit/nswindow/1449570-setisvisible?language=objc
func (x gen_NSWindow) SetIsVisible(
	flag bool,
) {
	C.NSWindow_inst_setIsVisible(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)

	return

}

// SetIsZoomed Sets the window’s zoomed state to the value you specify.
// https://developer.apple.com/documentation/appkit/nswindow/1449589-setiszoomed?language=objc
func (x gen_NSWindow) SetIsZoomed(
	flag bool,
) {
	C.NSWindow_inst_setIsZoomed(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)

	return

}

// SetTitleWithRepresentedFilename Sets a given path as the window’s title, formatting it as a file-system path, and records this path as the window’s associated file.
// https://developer.apple.com/documentation/appkit/nswindow/1419192-settitlewithrepresentedfilename?language=objc
func (x gen_NSWindow) SetTitleWithRepresentedFilename(
	filename core.NSStringRef,
) {
	C.NSWindow_inst_setTitleWithRepresentedFilename(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(filename),
	)

	return

}

// ToggleFullScreen Takes the window into or out of fullscreen mode,
// https://developer.apple.com/documentation/appkit/nswindow/1419527-togglefullscreen?language=objc
func (x gen_NSWindow) ToggleFullScreen(
	sender objc.Ref,
) {
	C.NSWindow_inst_toggleFullScreen(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleTabBar Shows or hides the tab bar.
// https://developer.apple.com/documentation/appkit/nswindow/1644517-toggletabbar?language=objc
func (x gen_NSWindow) ToggleTabBar(
	sender objc.Ref,
) {
	C.NSWindow_inst_toggleTabBar(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleTabOverview Shows or hides the tab overview.
// https://developer.apple.com/documentation/appkit/nswindow/2870175-toggletaboverview?language=objc
func (x gen_NSWindow) ToggleTabOverview(
	sender objc.Ref,
) {
	C.NSWindow_inst_toggleTabOverview(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleToolbarShown Toggles the visibility of the window’s toolbar.
// https://developer.apple.com/documentation/appkit/nswindow/1419554-toggletoolbarshown?language=objc
func (x gen_NSWindow) ToggleToolbarShown(
	sender objc.Ref,
) {
	C.NSWindow_inst_toggleToolbarShown(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// TryToPerform_with Dispatches action messages with a given argument.
// https://developer.apple.com/documentation/appkit/nswindow/1419428-trytoperform?language=objc
func (x gen_NSWindow) TryToPerform_with(
	action objc.Selector,
	object objc.Ref,
) bool {
	ret := C.NSWindow_inst_tryToPerform_with(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(object),
	)

	return convertObjCBoolToGo(ret)

}

// UnregisterDraggedTypes Unregisters the window as a possible destination for dragging operations.
// https://developer.apple.com/documentation/appkit/nswindow/1419456-unregisterdraggedtypes?language=objc
func (x gen_NSWindow) UnregisterDraggedTypes() {
	C.NSWindow_inst_unregisterDraggedTypes(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Update Updates the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419577-update?language=objc
func (x gen_NSWindow) Update() {
	C.NSWindow_inst_update(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// UpdateConstraintsIfNeeded Updates the constraints based on changes to views in the window since the last layout.
// https://developer.apple.com/documentation/appkit/nswindow/1526915-updateconstraintsifneeded?language=objc
func (x gen_NSWindow) UpdateConstraintsIfNeeded() {
	C.NSWindow_inst_updateConstraintsIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// VisualizeConstraints Displays a visual representation of the supplied constraints in the window.
// https://developer.apple.com/documentation/appkit/nswindow/1526997-visualizeconstraints?language=objc
func (x gen_NSWindow) VisualizeConstraints(
	constraints core.NSArrayRef,
) {
	C.NSWindow_inst_visualizeConstraints(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(constraints),
	)

	return

}

// Zoom Toggles the size and location of the window between its standard state (which the application provides as the best size to display the window’s data) and its user state (a new size and location the user may have set by moving or resizing the window).
// https://developer.apple.com/documentation/appkit/nswindow/1419513-zoom?language=objc
func (x gen_NSWindow) Zoom(
	sender objc.Ref,
) {
	C.NSWindow_inst_zoom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// Init_asNSWindow
func (x gen_NSWindow) Init_asNSWindow() NSWindow {
	ret := C.NSWindow_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSWindow_fromPointer(ret)

}

// Delegate The window’s delegate.
// https://developer.apple.com/documentation/appkit/nswindow/1419060-delegate?language=objc
func (x gen_NSWindow) Delegate() objc.Object {
	ret := C.NSWindow_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetDelegate The window’s delegate.
// https://developer.apple.com/documentation/appkit/nswindow/1419060-delegate?language=objc
func (x gen_NSWindow) SetDelegate(
	value objc.Ref,
) {
	C.NSWindow_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// ContentViewController The main content view controller for the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419615-contentviewcontroller?language=objc
func (x gen_NSWindow) ContentViewController() NSViewController {
	ret := C.NSWindow_inst_contentViewController(
		unsafe.Pointer(x.Pointer()),
	)

	return NSViewController_fromPointer(ret)

}

// SetContentViewController The main content view controller for the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419615-contentviewcontroller?language=objc
func (x gen_NSWindow) SetContentViewController(
	value NSViewControllerRef,
) {
	C.NSWindow_inst_setContentViewController(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// ContentView The window’s content view, the highest accessible view object in the window’s view hierarchy.
// https://developer.apple.com/documentation/appkit/nswindow/1419160-contentview?language=objc
func (x gen_NSWindow) ContentView() NSView {
	ret := C.NSWindow_inst_contentView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_fromPointer(ret)

}

// SetContentView The window’s content view, the highest accessible view object in the window’s view hierarchy.
// https://developer.apple.com/documentation/appkit/nswindow/1419160-contentview?language=objc
func (x gen_NSWindow) SetContentView(
	value NSViewRef,
) {
	C.NSWindow_inst_setContentView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// StyleMask Flags that describe the window’s current style, such as if it’s resizable or in full-screen mode.
// https://developer.apple.com/documentation/appkit/nswindow/1419078-stylemask?language=objc
func (x gen_NSWindow) StyleMask() core.NSUInteger {
	ret := C.NSWindow_inst_styleMask(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// SetStyleMask Flags that describe the window’s current style, such as if it’s resizable or in full-screen mode.
// https://developer.apple.com/documentation/appkit/nswindow/1419078-stylemask?language=objc
func (x gen_NSWindow) SetStyleMask(
	value core.NSUInteger,
) {
	C.NSWindow_inst_setStyleMask(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)

	return

}

// WorksWhenModal A Boolean value that indicates whether the window is able to receive keyboard and mouse events even when some other window is being run modally.
// https://developer.apple.com/documentation/appkit/nswindow/1419220-workswhenmodal?language=objc
func (x gen_NSWindow) WorksWhenModal() bool {
	ret := C.NSWindow_inst_worksWhenModal(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AlphaValue The window’s alpha value.
// https://developer.apple.com/documentation/appkit/nswindow/1419186-alphavalue?language=objc
func (x gen_NSWindow) AlphaValue() core.CGFloat {
	ret := C.NSWindow_inst_alphaValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// SetAlphaValue The window’s alpha value.
// https://developer.apple.com/documentation/appkit/nswindow/1419186-alphavalue?language=objc
func (x gen_NSWindow) SetAlphaValue(
	value core.CGFloat,
) {
	C.NSWindow_inst_setAlphaValue(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// BackgroundColor The color of the window’s background.
// https://developer.apple.com/documentation/appkit/nswindow/1419751-backgroundcolor?language=objc
func (x gen_NSWindow) BackgroundColor() NSColor {
	ret := C.NSWindow_inst_backgroundColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_fromPointer(ret)

}

// SetBackgroundColor The color of the window’s background.
// https://developer.apple.com/documentation/appkit/nswindow/1419751-backgroundcolor?language=objc
func (x gen_NSWindow) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSWindow_inst_setBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// CanHide A Boolean value that indicates whether the window can hide when its application becomes hidden.
// https://developer.apple.com/documentation/appkit/nswindow/1419725-canhide?language=objc
func (x gen_NSWindow) CanHide() bool {
	ret := C.NSWindow_inst_canHide(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetCanHide A Boolean value that indicates whether the window can hide when its application becomes hidden.
// https://developer.apple.com/documentation/appkit/nswindow/1419725-canhide?language=objc
func (x gen_NSWindow) SetCanHide(
	value bool,
) {
	C.NSWindow_inst_setCanHide(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsOnActiveSpace A Boolean value that indicates whether the window is on the currently active space.
// https://developer.apple.com/documentation/appkit/nswindow/1419707-onactivespace?language=objc
func (x gen_NSWindow) IsOnActiveSpace() bool {
	ret := C.NSWindow_inst_isOnActiveSpace(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// HidesOnDeactivate A Boolean value that indicates whether the window is removed from the screen when its application becomes inactive.
// https://developer.apple.com/documentation/appkit/nswindow/1419777-hidesondeactivate?language=objc
func (x gen_NSWindow) HidesOnDeactivate() bool {
	ret := C.NSWindow_inst_hidesOnDeactivate(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetHidesOnDeactivate A Boolean value that indicates whether the window is removed from the screen when its application becomes inactive.
// https://developer.apple.com/documentation/appkit/nswindow/1419777-hidesondeactivate?language=objc
func (x gen_NSWindow) SetHidesOnDeactivate(
	value bool,
) {
	C.NSWindow_inst_setHidesOnDeactivate(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// CollectionBehavior A value that identifies the window’s behavior in window collections.
// https://developer.apple.com/documentation/appkit/nswindow/1419471-collectionbehavior?language=objc
func (x gen_NSWindow) CollectionBehavior() core.NSUInteger {
	ret := C.NSWindow_inst_collectionBehavior(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// SetCollectionBehavior A value that identifies the window’s behavior in window collections.
// https://developer.apple.com/documentation/appkit/nswindow/1419471-collectionbehavior?language=objc
func (x gen_NSWindow) SetCollectionBehavior(
	value core.NSUInteger,
) {
	C.NSWindow_inst_setCollectionBehavior(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)

	return

}

// IsOpaque A Boolean value that indicates whether the window is opaque.
// https://developer.apple.com/documentation/appkit/nswindow/1419086-opaque?language=objc
func (x gen_NSWindow) IsOpaque() bool {
	ret := C.NSWindow_inst_isOpaque(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetOpaque A Boolean value that indicates whether the window is opaque.
// https://developer.apple.com/documentation/appkit/nswindow/1419086-opaque?language=objc
func (x gen_NSWindow) SetOpaque(
	value bool,
) {
	C.NSWindow_inst_setOpaque(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// HasShadow A Boolean value that indicates whether the window has a shadow.
// https://developer.apple.com/documentation/appkit/nswindow/1419234-hasshadow?language=objc
func (x gen_NSWindow) HasShadow() bool {
	ret := C.NSWindow_inst_hasShadow(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetHasShadow A Boolean value that indicates whether the window has a shadow.
// https://developer.apple.com/documentation/appkit/nswindow/1419234-hasshadow?language=objc
func (x gen_NSWindow) SetHasShadow(
	value bool,
) {
	C.NSWindow_inst_setHasShadow(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// PreventsApplicationTerminationWhenModal A Boolean value that indicates whether the window prevents application termination when modal.
// https://developer.apple.com/documentation/appkit/nswindow/1419743-preventsapplicationterminationwh?language=objc
func (x gen_NSWindow) PreventsApplicationTerminationWhenModal() bool {
	ret := C.NSWindow_inst_preventsApplicationTerminationWhenModal(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetPreventsApplicationTerminationWhenModal A Boolean value that indicates whether the window prevents application termination when modal.
// https://developer.apple.com/documentation/appkit/nswindow/1419743-preventsapplicationterminationwh?language=objc
func (x gen_NSWindow) SetPreventsApplicationTerminationWhenModal(
	value bool,
) {
	C.NSWindow_inst_setPreventsApplicationTerminationWhenModal(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// HasDynamicDepthLimit A Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on.
// https://developer.apple.com/documentation/appkit/nswindow/1419330-hasdynamicdepthlimit?language=objc
func (x gen_NSWindow) HasDynamicDepthLimit() bool {
	ret := C.NSWindow_inst_hasDynamicDepthLimit(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// WindowNumber The window number of the window’s window device.
// https://developer.apple.com/documentation/appkit/nswindow/1419068-windownumber?language=objc
func (x gen_NSWindow) WindowNumber() core.NSInteger {
	ret := C.NSWindow_inst_windowNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// DeviceDescription A dictionary containing information about the window’s resolution, such as color, depth, and so on.
// https://developer.apple.com/documentation/appkit/nswindow/1419741-devicedescription?language=objc
func (x gen_NSWindow) DeviceDescription() core.NSDictionary {
	ret := C.NSWindow_inst_deviceDescription(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_fromPointer(ret)

}

// CanBecomeVisibleWithoutLogin A Boolean value that indicates whether the window can be displayed at the login window.
// https://developer.apple.com/documentation/appkit/nswindow/1419179-canbecomevisiblewithoutlogin?language=objc
func (x gen_NSWindow) CanBecomeVisibleWithoutLogin() bool {
	ret := C.NSWindow_inst_canBecomeVisibleWithoutLogin(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetCanBecomeVisibleWithoutLogin A Boolean value that indicates whether the window can be displayed at the login window.
// https://developer.apple.com/documentation/appkit/nswindow/1419179-canbecomevisiblewithoutlogin?language=objc
func (x gen_NSWindow) SetCanBecomeVisibleWithoutLogin(
	value bool,
) {
	C.NSWindow_inst_setCanBecomeVisibleWithoutLogin(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// BackingType The window’s backing store type.
// https://developer.apple.com/documentation/appkit/nswindow/1419599-backingtype?language=objc
func (x gen_NSWindow) BackingType() core.NSUInteger {
	ret := C.NSWindow_inst_backingType(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSUInteger(ret)

}

// SetBackingType The window’s backing store type.
// https://developer.apple.com/documentation/appkit/nswindow/1419599-backingtype?language=objc
func (x gen_NSWindow) SetBackingType(
	value core.NSUInteger,
) {
	C.NSWindow_inst_setBackingType(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)

	return

}

// AttachedSheet The sheet attached to the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419467-attachedsheet?language=objc
func (x gen_NSWindow) AttachedSheet() NSWindow {
	ret := C.NSWindow_inst_attachedSheet(
		unsafe.Pointer(x.Pointer()),
	)

	return NSWindow_fromPointer(ret)

}

// IsSheet A Boolean value that indicates whether the window has ever run as a modal sheet.
// https://developer.apple.com/documentation/appkit/nswindow/1419364-sheet?language=objc
func (x gen_NSWindow) IsSheet() bool {
	ret := C.NSWindow_inst_isSheet(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SheetParent The window to which the sheet is attached.
// https://developer.apple.com/documentation/appkit/nswindow/1419052-sheetparent?language=objc
func (x gen_NSWindow) SheetParent() NSWindow {
	ret := C.NSWindow_inst_sheetParent(
		unsafe.Pointer(x.Pointer()),
	)

	return NSWindow_fromPointer(ret)

}

// Sheets An array of the sheets currently attached to the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419765-sheets?language=objc
func (x gen_NSWindow) Sheets() core.NSArray {
	ret := C.NSWindow_inst_sheets(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// Frame The window’s frame rectangle in screen coordinates, including the title bar.
// https://developer.apple.com/documentation/appkit/nswindow/1419697-frame?language=objc
func (x gen_NSWindow) Frame() core.NSRect {
	ret := C.NSWindow_inst_frame(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// AspectRatio The window’s aspect ratio, which constrains the size of its frame rectangle to integral multiples of this ratio when the user resizes it.
// https://developer.apple.com/documentation/appkit/nswindow/1419507-aspectratio?language=objc
func (x gen_NSWindow) AspectRatio() core.NSSize {
	ret := C.NSWindow_inst_aspectRatio(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetAspectRatio The window’s aspect ratio, which constrains the size of its frame rectangle to integral multiples of this ratio when the user resizes it.
// https://developer.apple.com/documentation/appkit/nswindow/1419507-aspectratio?language=objc
func (x gen_NSWindow) SetAspectRatio(
	value core.NSSize,
) {
	C.NSWindow_inst_setAspectRatio(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// MinSize The minimum size to which the window’s frame (including its title bar) can be sized.
// https://developer.apple.com/documentation/appkit/nswindow/1419206-minsize?language=objc
func (x gen_NSWindow) MinSize() core.NSSize {
	ret := C.NSWindow_inst_minSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetMinSize The minimum size to which the window’s frame (including its title bar) can be sized.
// https://developer.apple.com/documentation/appkit/nswindow/1419206-minsize?language=objc
func (x gen_NSWindow) SetMinSize(
	value core.NSSize,
) {
	C.NSWindow_inst_setMinSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// MaxSize The maximum size to which the window’s frame (including its title bar) can be sized.
// https://developer.apple.com/documentation/appkit/nswindow/1419595-maxsize?language=objc
func (x gen_NSWindow) MaxSize() core.NSSize {
	ret := C.NSWindow_inst_maxSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetMaxSize The maximum size to which the window’s frame (including its title bar) can be sized.
// https://developer.apple.com/documentation/appkit/nswindow/1419595-maxsize?language=objc
func (x gen_NSWindow) SetMaxSize(
	value core.NSSize,
) {
	C.NSWindow_inst_setMaxSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// IsZoomed A Boolean value that indicates whether the window is in a zoomed state.
// https://developer.apple.com/documentation/appkit/nswindow/1419398-zoomed?language=objc
func (x gen_NSWindow) IsZoomed() bool {
	ret := C.NSWindow_inst_isZoomed(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// ResizeIncrements The window’s resizing increments.
// https://developer.apple.com/documentation/appkit/nswindow/1419390-resizeincrements?language=objc
func (x gen_NSWindow) ResizeIncrements() core.NSSize {
	ret := C.NSWindow_inst_resizeIncrements(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetResizeIncrements The window’s resizing increments.
// https://developer.apple.com/documentation/appkit/nswindow/1419390-resizeincrements?language=objc
func (x gen_NSWindow) SetResizeIncrements(
	value core.NSSize,
) {
	C.NSWindow_inst_setResizeIncrements(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// PreservesContentDuringLiveResize A Boolean value that indicates whether the window tries to optimize user-initiated resize operations by preserving the content of views that have not changed.
// https://developer.apple.com/documentation/appkit/nswindow/1419588-preservescontentduringliveresize?language=objc
func (x gen_NSWindow) PreservesContentDuringLiveResize() bool {
	ret := C.NSWindow_inst_preservesContentDuringLiveResize(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetPreservesContentDuringLiveResize A Boolean value that indicates whether the window tries to optimize user-initiated resize operations by preserving the content of views that have not changed.
// https://developer.apple.com/documentation/appkit/nswindow/1419588-preservescontentduringliveresize?language=objc
func (x gen_NSWindow) SetPreservesContentDuringLiveResize(
	value bool,
) {
	C.NSWindow_inst_setPreservesContentDuringLiveResize(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// InLiveResize A Boolean value that indicates whether the window is being resized by the user.
// https://developer.apple.com/documentation/appkit/nswindow/1419378-inliveresize?language=objc
func (x gen_NSWindow) InLiveResize() bool {
	ret := C.NSWindow_inst_inLiveResize(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// ContentAspectRatio The window’s content aspect ratio.
// https://developer.apple.com/documentation/appkit/nswindow/1419148-contentaspectratio?language=objc
func (x gen_NSWindow) ContentAspectRatio() core.NSSize {
	ret := C.NSWindow_inst_contentAspectRatio(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetContentAspectRatio The window’s content aspect ratio.
// https://developer.apple.com/documentation/appkit/nswindow/1419148-contentaspectratio?language=objc
func (x gen_NSWindow) SetContentAspectRatio(
	value core.NSSize,
) {
	C.NSWindow_inst_setContentAspectRatio(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// ContentMinSize The minimum size of the window’s content view in the window’s base coordinate system.
// https://developer.apple.com/documentation/appkit/nswindow/1419670-contentminsize?language=objc
func (x gen_NSWindow) ContentMinSize() core.NSSize {
	ret := C.NSWindow_inst_contentMinSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetContentMinSize The minimum size of the window’s content view in the window’s base coordinate system.
// https://developer.apple.com/documentation/appkit/nswindow/1419670-contentminsize?language=objc
func (x gen_NSWindow) SetContentMinSize(
	value core.NSSize,
) {
	C.NSWindow_inst_setContentMinSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// ContentMaxSize The maximum size of the window’s content view in the window’s base coordinate system.
// https://developer.apple.com/documentation/appkit/nswindow/1419154-contentmaxsize?language=objc
func (x gen_NSWindow) ContentMaxSize() core.NSSize {
	ret := C.NSWindow_inst_contentMaxSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetContentMaxSize The maximum size of the window’s content view in the window’s base coordinate system.
// https://developer.apple.com/documentation/appkit/nswindow/1419154-contentmaxsize?language=objc
func (x gen_NSWindow) SetContentMaxSize(
	value core.NSSize,
) {
	C.NSWindow_inst_setContentMaxSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// ContentResizeIncrements The window’s content-view resizing increments.
// https://developer.apple.com/documentation/appkit/nswindow/1419649-contentresizeincrements?language=objc
func (x gen_NSWindow) ContentResizeIncrements() core.NSSize {
	ret := C.NSWindow_inst_contentResizeIncrements(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetContentResizeIncrements The window’s content-view resizing increments.
// https://developer.apple.com/documentation/appkit/nswindow/1419649-contentresizeincrements?language=objc
func (x gen_NSWindow) SetContentResizeIncrements(
	value core.NSSize,
) {
	C.NSWindow_inst_setContentResizeIncrements(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// ContentLayoutGuide A value used by Auto Layout constraints to automatically bind to the value of contentLayoutRect.
// https://developer.apple.com/documentation/appkit/nswindow/1419094-contentlayoutguide?language=objc
func (x gen_NSWindow) ContentLayoutGuide() objc.Object {
	ret := C.NSWindow_inst_contentLayoutGuide(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// ContentLayoutRect The area inside the window that is for non-obscured content, in window coordinates.
// https://developer.apple.com/documentation/appkit/nswindow/1419124-contentlayoutrect?language=objc
func (x gen_NSWindow) ContentLayoutRect() core.NSRect {
	ret := C.NSWindow_inst_contentLayoutRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// MaxFullScreenContentSize A maximum size that is used to determine if a window can fit when it is in full screen in a tile.
// https://developer.apple.com/documentation/appkit/nswindow/1419438-maxfullscreencontentsize?language=objc
func (x gen_NSWindow) MaxFullScreenContentSize() core.NSSize {
	ret := C.NSWindow_inst_maxFullScreenContentSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetMaxFullScreenContentSize A maximum size that is used to determine if a window can fit when it is in full screen in a tile.
// https://developer.apple.com/documentation/appkit/nswindow/1419438-maxfullscreencontentsize?language=objc
func (x gen_NSWindow) SetMaxFullScreenContentSize(
	value core.NSSize,
) {
	C.NSWindow_inst_setMaxFullScreenContentSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// MinFullScreenContentSize A minimum size that is used to determine if a window can fit when it is in full screen in a tile.
// https://developer.apple.com/documentation/appkit/nswindow/1419627-minfullscreencontentsize?language=objc
func (x gen_NSWindow) MinFullScreenContentSize() core.NSSize {
	ret := C.NSWindow_inst_minFullScreenContentSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetMinFullScreenContentSize A minimum size that is used to determine if a window can fit when it is in full screen in a tile.
// https://developer.apple.com/documentation/appkit/nswindow/1419627-minfullscreencontentsize?language=objc
func (x gen_NSWindow) SetMinFullScreenContentSize(
	value core.NSSize,
) {
	C.NSWindow_inst_setMinFullScreenContentSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// Level The window level of the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419511-level?language=objc
func (x gen_NSWindow) Level() core.NSInteger {
	ret := C.NSWindow_inst_level(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// SetLevel The window level of the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419511-level?language=objc
func (x gen_NSWindow) SetLevel(
	value core.NSInteger,
) {
	C.NSWindow_inst_setLevel(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return

}

// IsVisible A Boolean value that indicates whether the window is visible onscreen (even when it’s obscured by other windows).
// https://developer.apple.com/documentation/appkit/nswindow/1419132-visible?language=objc
func (x gen_NSWindow) IsVisible() bool {
	ret := C.NSWindow_inst_isVisible(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsKeyWindow A Boolean value that indicates whether the window is the key window for the application.
// https://developer.apple.com/documentation/appkit/nswindow/1419735-keywindow?language=objc
func (x gen_NSWindow) IsKeyWindow() bool {
	ret := C.NSWindow_inst_isKeyWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// CanBecomeKeyWindow A Boolean value that indicates whether the window can become the key window.
// https://developer.apple.com/documentation/appkit/nswindow/1419543-canbecomekeywindow?language=objc
func (x gen_NSWindow) CanBecomeKeyWindow() bool {
	ret := C.NSWindow_inst_canBecomeKeyWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsMainWindow A Boolean value that indicates whether the window is the application’s main window.
// https://developer.apple.com/documentation/appkit/nswindow/1419130-mainwindow?language=objc
func (x gen_NSWindow) IsMainWindow() bool {
	ret := C.NSWindow_inst_isMainWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// CanBecomeMainWindow A Boolean value that indicates whether the window can become the application’s main window.
// https://developer.apple.com/documentation/appkit/nswindow/1419162-canbecomemainwindow?language=objc
func (x gen_NSWindow) CanBecomeMainWindow() bool {
	ret := C.NSWindow_inst_canBecomeMainWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// ChildWindows An array of the window’s attached child windows.
// https://developer.apple.com/documentation/appkit/nswindow/1419236-childwindows?language=objc
func (x gen_NSWindow) ChildWindows() core.NSArray {
	ret := C.NSWindow_inst_childWindows(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// ParentWindow The parent window to which the window is attached as a child.
// https://developer.apple.com/documentation/appkit/nswindow/1419695-parentwindow?language=objc
func (x gen_NSWindow) ParentWindow() NSWindow {
	ret := C.NSWindow_inst_parentWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return NSWindow_fromPointer(ret)

}

// SetParentWindow The parent window to which the window is attached as a child.
// https://developer.apple.com/documentation/appkit/nswindow/1419695-parentwindow?language=objc
func (x gen_NSWindow) SetParentWindow(
	value NSWindowRef,
) {
	C.NSWindow_inst_setParentWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// IsExcludedFromWindowsMenu A Boolean value that indicates whether the window is excluded from the application’s Windows menu.
// https://developer.apple.com/documentation/appkit/nswindow/1419175-excludedfromwindowsmenu?language=objc
func (x gen_NSWindow) IsExcludedFromWindowsMenu() bool {
	ret := C.NSWindow_inst_isExcludedFromWindowsMenu(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetExcludedFromWindowsMenu A Boolean value that indicates whether the window is excluded from the application’s Windows menu.
// https://developer.apple.com/documentation/appkit/nswindow/1419175-excludedfromwindowsmenu?language=objc
func (x gen_NSWindow) SetExcludedFromWindowsMenu(
	value bool,
) {
	C.NSWindow_inst_setExcludedFromWindowsMenu(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AreCursorRectsEnabled A Boolean value that indicates whether the window’s cursor rectangles are enabled.
// https://developer.apple.com/documentation/appkit/nswindow/1419668-arecursorrectsenabled?language=objc
func (x gen_NSWindow) AreCursorRectsEnabled() bool {
	ret := C.NSWindow_inst_areCursorRectsEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// ShowsToolbarButton A Boolean value that indicates whether the toolbar control button is currently displayed.
// https://developer.apple.com/documentation/appkit/nswindow/1419196-showstoolbarbutton?language=objc
func (x gen_NSWindow) ShowsToolbarButton() bool {
	ret := C.NSWindow_inst_showsToolbarButton(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetShowsToolbarButton A Boolean value that indicates whether the toolbar control button is currently displayed.
// https://developer.apple.com/documentation/appkit/nswindow/1419196-showstoolbarbutton?language=objc
func (x gen_NSWindow) SetShowsToolbarButton(
	value bool,
) {
	C.NSWindow_inst_setShowsToolbarButton(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// TitlebarAppearsTransparent A Boolean value that indicates whether the title bar draws its background.
// https://developer.apple.com/documentation/appkit/nswindow/1419167-titlebarappearstransparent?language=objc
func (x gen_NSWindow) TitlebarAppearsTransparent() bool {
	ret := C.NSWindow_inst_titlebarAppearsTransparent(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetTitlebarAppearsTransparent A Boolean value that indicates whether the title bar draws its background.
// https://developer.apple.com/documentation/appkit/nswindow/1419167-titlebarappearstransparent?language=objc
func (x gen_NSWindow) SetTitlebarAppearsTransparent(
	value bool,
) {
	C.NSWindow_inst_setTitlebarAppearsTransparent(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// TitlebarAccessoryViewControllers An array of title bar accessory view controllers that are currently added to the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419547-titlebaraccessoryviewcontrollers?language=objc
func (x gen_NSWindow) TitlebarAccessoryViewControllers() core.NSArray {
	ret := C.NSWindow_inst_titlebarAccessoryViewControllers(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// SetTitlebarAccessoryViewControllers An array of title bar accessory view controllers that are currently added to the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419547-titlebaraccessoryviewcontrollers?language=objc
func (x gen_NSWindow) SetTitlebarAccessoryViewControllers(
	value core.NSArrayRef,
) {
	C.NSWindow_inst_setTitlebarAccessoryViewControllers(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// TabbedWindows An array of windows that display as tabs.
// https://developer.apple.com/documentation/appkit/nswindow/1792044-tabbedwindows?language=objc
func (x gen_NSWindow) TabbedWindows() core.NSArray {
	ret := C.NSWindow_inst_tabbedWindows(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// AllowsToolTipsWhenApplicationIsInactive A Boolean value that indicates whether the window can display tooltips even when the application is in the background.
// https://developer.apple.com/documentation/appkit/nswindow/1419138-allowstooltipswhenapplicationisi?language=objc
func (x gen_NSWindow) AllowsToolTipsWhenApplicationIsInactive() bool {
	ret := C.NSWindow_inst_allowsToolTipsWhenApplicationIsInactive(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsToolTipsWhenApplicationIsInactive A Boolean value that indicates whether the window can display tooltips even when the application is in the background.
// https://developer.apple.com/documentation/appkit/nswindow/1419138-allowstooltipswhenapplicationisi?language=objc
func (x gen_NSWindow) SetAllowsToolTipsWhenApplicationIsInactive(
	value bool,
) {
	C.NSWindow_inst_setAllowsToolTipsWhenApplicationIsInactive(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// CurrentEvent The event currently being processed by the application.
// https://developer.apple.com/documentation/appkit/nswindow/1419298-currentevent?language=objc
func (x gen_NSWindow) CurrentEvent() NSEvent {
	ret := C.NSWindow_inst_currentEvent(
		unsafe.Pointer(x.Pointer()),
	)

	return NSEvent_fromPointer(ret)

}

// InitialFirstResponder The view that’s made first responder (also called the key view) the first time the window is placed onscreen.
// https://developer.apple.com/documentation/appkit/nswindow/1419479-initialfirstresponder?language=objc
func (x gen_NSWindow) InitialFirstResponder() NSView {
	ret := C.NSWindow_inst_initialFirstResponder(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_fromPointer(ret)

}

// SetInitialFirstResponder The view that’s made first responder (also called the key view) the first time the window is placed onscreen.
// https://developer.apple.com/documentation/appkit/nswindow/1419479-initialfirstresponder?language=objc
func (x gen_NSWindow) SetInitialFirstResponder(
	value NSViewRef,
) {
	C.NSWindow_inst_setInitialFirstResponder(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// AutorecalculatesKeyViewLoop A Boolean value that indicates whether the window automatically recalculates the key view loop when views are added.
// https://developer.apple.com/documentation/appkit/nswindow/1419214-autorecalculateskeyviewloop?language=objc
func (x gen_NSWindow) AutorecalculatesKeyViewLoop() bool {
	ret := C.NSWindow_inst_autorecalculatesKeyViewLoop(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAutorecalculatesKeyViewLoop A Boolean value that indicates whether the window automatically recalculates the key view loop when views are added.
// https://developer.apple.com/documentation/appkit/nswindow/1419214-autorecalculateskeyviewloop?language=objc
func (x gen_NSWindow) SetAutorecalculatesKeyViewLoop(
	value bool,
) {
	C.NSWindow_inst_setAutorecalculatesKeyViewLoop(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AcceptsMouseMovedEvents A Boolean value that indicates whether the window accepts mouse-moved events.
// https://developer.apple.com/documentation/appkit/nswindow/1419340-acceptsmousemovedevents?language=objc
func (x gen_NSWindow) AcceptsMouseMovedEvents() bool {
	ret := C.NSWindow_inst_acceptsMouseMovedEvents(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAcceptsMouseMovedEvents A Boolean value that indicates whether the window accepts mouse-moved events.
// https://developer.apple.com/documentation/appkit/nswindow/1419340-acceptsmousemovedevents?language=objc
func (x gen_NSWindow) SetAcceptsMouseMovedEvents(
	value bool,
) {
	C.NSWindow_inst_setAcceptsMouseMovedEvents(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IgnoresMouseEvents A Boolean value that indicates whether the window is transparent to mouse events.
// https://developer.apple.com/documentation/appkit/nswindow/1419354-ignoresmouseevents?language=objc
func (x gen_NSWindow) IgnoresMouseEvents() bool {
	ret := C.NSWindow_inst_ignoresMouseEvents(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetIgnoresMouseEvents A Boolean value that indicates whether the window is transparent to mouse events.
// https://developer.apple.com/documentation/appkit/nswindow/1419354-ignoresmouseevents?language=objc
func (x gen_NSWindow) SetIgnoresMouseEvents(
	value bool,
) {
	C.NSWindow_inst_setIgnoresMouseEvents(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// MouseLocationOutsideOfEventStream The current location of the pointer reckoned in the window’s base coordinate system, regardless of the current event being handled or of any events pending.
// https://developer.apple.com/documentation/appkit/nswindow/1419280-mouselocationoutsideofeventstrea?language=objc
func (x gen_NSWindow) MouseLocationOutsideOfEventStream() core.NSPoint {
	ret := C.NSWindow_inst_mouseLocationOutsideOfEventStream(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// IsRestorable A Boolean value indicating whether the window configuration is preserved between application launches.
// https://developer.apple.com/documentation/appkit/nswindow/1526255-restorable?language=objc
func (x gen_NSWindow) IsRestorable() bool {
	ret := C.NSWindow_inst_isRestorable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetRestorable A Boolean value indicating whether the window configuration is preserved between application launches.
// https://developer.apple.com/documentation/appkit/nswindow/1526255-restorable?language=objc
func (x gen_NSWindow) SetRestorable(
	value bool,
) {
	C.NSWindow_inst_setRestorable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// ViewsNeedDisplay A Boolean value that indicates whether any of the window’s views need to be displayed.
// https://developer.apple.com/documentation/appkit/nswindow/1419609-viewsneeddisplay?language=objc
func (x gen_NSWindow) ViewsNeedDisplay() bool {
	ret := C.NSWindow_inst_viewsNeedDisplay(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetViewsNeedDisplay A Boolean value that indicates whether any of the window’s views need to be displayed.
// https://developer.apple.com/documentation/appkit/nswindow/1419609-viewsneeddisplay?language=objc
func (x gen_NSWindow) SetViewsNeedDisplay(
	value bool,
) {
	C.NSWindow_inst_setViewsNeedDisplay(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AllowsConcurrentViewDrawing A Boolean value that indicates whether the window allows multithreaded view drawing.
// https://developer.apple.com/documentation/appkit/nswindow/1419300-allowsconcurrentviewdrawing?language=objc
func (x gen_NSWindow) AllowsConcurrentViewDrawing() bool {
	ret := C.NSWindow_inst_allowsConcurrentViewDrawing(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsConcurrentViewDrawing A Boolean value that indicates whether the window allows multithreaded view drawing.
// https://developer.apple.com/documentation/appkit/nswindow/1419300-allowsconcurrentviewdrawing?language=objc
func (x gen_NSWindow) SetAllowsConcurrentViewDrawing(
	value bool,
) {
	C.NSWindow_inst_setAllowsConcurrentViewDrawing(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsDocumentEdited A Boolean value that indicates whether the window’s document has been edited.
// https://developer.apple.com/documentation/appkit/nswindow/1419311-documentedited?language=objc
func (x gen_NSWindow) IsDocumentEdited() bool {
	ret := C.NSWindow_inst_isDocumentEdited(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetDocumentEdited A Boolean value that indicates whether the window’s document has been edited.
// https://developer.apple.com/documentation/appkit/nswindow/1419311-documentedited?language=objc
func (x gen_NSWindow) SetDocumentEdited(
	value bool,
) {
	C.NSWindow_inst_setDocumentEdited(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// BackingScaleFactor The backing scale factor.
// https://developer.apple.com/documentation/appkit/nswindow/1419459-backingscalefactor?language=objc
func (x gen_NSWindow) BackingScaleFactor() core.CGFloat {
	ret := C.NSWindow_inst_backingScaleFactor(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// Title The string that appears in the title bar of the window or the path to the represented file.
// https://developer.apple.com/documentation/appkit/nswindow/1419404-title?language=objc
func (x gen_NSWindow) Title() core.NSString {
	ret := C.NSWindow_inst_title(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetTitle The string that appears in the title bar of the window or the path to the represented file.
// https://developer.apple.com/documentation/appkit/nswindow/1419404-title?language=objc
func (x gen_NSWindow) SetTitle(
	value core.NSStringRef,
) {
	C.NSWindow_inst_setTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Subtitle A secondary line of text that appears in the title bar of the window.
// https://developer.apple.com/documentation/appkit/nswindow/3608198-subtitle?language=objc
func (x gen_NSWindow) Subtitle() core.NSString {
	ret := C.NSWindow_inst_subtitle(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetSubtitle A secondary line of text that appears in the title bar of the window.
// https://developer.apple.com/documentation/appkit/nswindow/3608198-subtitle?language=objc
func (x gen_NSWindow) SetSubtitle(
	value core.NSStringRef,
) {
	C.NSWindow_inst_setSubtitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// TitleVisibility A value that indicates the visibility of the window’s title and title bar buttons.
// https://developer.apple.com/documentation/appkit/nswindow/1419635-titlevisibility?language=objc
func (x gen_NSWindow) TitleVisibility() core.NSInteger {
	ret := C.NSWindow_inst_titleVisibility(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// SetTitleVisibility A value that indicates the visibility of the window’s title and title bar buttons.
// https://developer.apple.com/documentation/appkit/nswindow/1419635-titlevisibility?language=objc
func (x gen_NSWindow) SetTitleVisibility(
	value core.NSInteger,
) {
	C.NSWindow_inst_setTitleVisibility(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return

}

// RepresentedFilename The path to the file of the window’s represented file.
// https://developer.apple.com/documentation/appkit/nswindow/1419631-representedfilename?language=objc
func (x gen_NSWindow) RepresentedFilename() core.NSString {
	ret := C.NSWindow_inst_representedFilename(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetRepresentedFilename The path to the file of the window’s represented file.
// https://developer.apple.com/documentation/appkit/nswindow/1419631-representedfilename?language=objc
func (x gen_NSWindow) SetRepresentedFilename(
	value core.NSStringRef,
) {
	C.NSWindow_inst_setRepresentedFilename(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// RepresentedURL The URL of the file the window represents.
// https://developer.apple.com/documentation/appkit/nswindow/1419066-representedurl?language=objc
func (x gen_NSWindow) RepresentedURL() core.NSURL {
	ret := C.NSWindow_inst_representedURL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_fromPointer(ret)

}

// SetRepresentedURL The URL of the file the window represents.
// https://developer.apple.com/documentation/appkit/nswindow/1419066-representedurl?language=objc
func (x gen_NSWindow) SetRepresentedURL(
	value core.NSURLRef,
) {
	C.NSWindow_inst_setRepresentedURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Screen The screen the window is on.
// https://developer.apple.com/documentation/appkit/nswindow/1419232-screen?language=objc
func (x gen_NSWindow) Screen() NSScreen {
	ret := C.NSWindow_inst_screen(
		unsafe.Pointer(x.Pointer()),
	)

	return NSScreen_fromPointer(ret)

}

// DeepestScreen The deepest screen the window is on (it may be split over several screens).
// https://developer.apple.com/documentation/appkit/nswindow/1419080-deepestscreen?language=objc
func (x gen_NSWindow) DeepestScreen() NSScreen {
	ret := C.NSWindow_inst_deepestScreen(
		unsafe.Pointer(x.Pointer()),
	)

	return NSScreen_fromPointer(ret)

}

// DisplaysWhenScreenProfileChanges A Boolean value that indicates whether the window context should be updated when the screen profile changes or when the window moves to a different screen.
// https://developer.apple.com/documentation/appkit/nswindow/1419430-displayswhenscreenprofilechanges?language=objc
func (x gen_NSWindow) DisplaysWhenScreenProfileChanges() bool {
	ret := C.NSWindow_inst_displaysWhenScreenProfileChanges(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetDisplaysWhenScreenProfileChanges A Boolean value that indicates whether the window context should be updated when the screen profile changes or when the window moves to a different screen.
// https://developer.apple.com/documentation/appkit/nswindow/1419430-displayswhenscreenprofilechanges?language=objc
func (x gen_NSWindow) SetDisplaysWhenScreenProfileChanges(
	value bool,
) {
	C.NSWindow_inst_setDisplaysWhenScreenProfileChanges(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsMovableByWindowBackground A Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background.
// https://developer.apple.com/documentation/appkit/nswindow/1419072-movablebywindowbackground?language=objc
func (x gen_NSWindow) IsMovableByWindowBackground() bool {
	ret := C.NSWindow_inst_isMovableByWindowBackground(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetMovableByWindowBackground A Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background.
// https://developer.apple.com/documentation/appkit/nswindow/1419072-movablebywindowbackground?language=objc
func (x gen_NSWindow) SetMovableByWindowBackground(
	value bool,
) {
	C.NSWindow_inst_setMovableByWindowBackground(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsMovable A Boolean value that indicates whether the window can be dragged by clicking in its title bar or background.
// https://developer.apple.com/documentation/appkit/nswindow/1419579-movable?language=objc
func (x gen_NSWindow) IsMovable() bool {
	ret := C.NSWindow_inst_isMovable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetMovable A Boolean value that indicates whether the window can be dragged by clicking in its title bar or background.
// https://developer.apple.com/documentation/appkit/nswindow/1419579-movable?language=objc
func (x gen_NSWindow) SetMovable(
	value bool,
) {
	C.NSWindow_inst_setMovable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsReleasedWhenClosed A Boolean value that indicates whether the window is released when it receives the close message.
// https://developer.apple.com/documentation/appkit/nswindow/1419062-releasedwhenclosed?language=objc
func (x gen_NSWindow) IsReleasedWhenClosed() bool {
	ret := C.NSWindow_inst_isReleasedWhenClosed(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetReleasedWhenClosed A Boolean value that indicates whether the window is released when it receives the close message.
// https://developer.apple.com/documentation/appkit/nswindow/1419062-releasedwhenclosed?language=objc
func (x gen_NSWindow) SetReleasedWhenClosed(
	value bool,
) {
	C.NSWindow_inst_setReleasedWhenClosed(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsMiniaturized A Boolean value that indicates whether the window is minimized.
// https://developer.apple.com/documentation/appkit/nswindow/1419699-miniaturized?language=objc
func (x gen_NSWindow) IsMiniaturized() bool {
	ret := C.NSWindow_inst_isMiniaturized(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// MiniwindowImage The custom miniaturized window image of the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419185-miniwindowimage?language=objc
func (x gen_NSWindow) MiniwindowImage() NSImage {
	ret := C.NSWindow_inst_miniwindowImage(
		unsafe.Pointer(x.Pointer()),
	)

	return NSImage_fromPointer(ret)

}

// SetMiniwindowImage The custom miniaturized window image of the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419185-miniwindowimage?language=objc
func (x gen_NSWindow) SetMiniwindowImage(
	value NSImageRef,
) {
	C.NSWindow_inst_setMiniwindowImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// MiniwindowTitle The title displayed in the window’s minimized window.
// https://developer.apple.com/documentation/appkit/nswindow/1419571-miniwindowtitle?language=objc
func (x gen_NSWindow) MiniwindowTitle() core.NSString {
	ret := C.NSWindow_inst_miniwindowTitle(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetMiniwindowTitle The title displayed in the window’s minimized window.
// https://developer.apple.com/documentation/appkit/nswindow/1419571-miniwindowtitle?language=objc
func (x gen_NSWindow) SetMiniwindowTitle(
	value core.NSStringRef,
) {
	C.NSWindow_inst_setMiniwindowTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// HasCloseBox A Boolean value that indicates if the window has a close box.
// https://developer.apple.com/documentation/appkit/nswindow/1449574-hasclosebox?language=objc
func (x gen_NSWindow) HasCloseBox() bool {
	ret := C.NSWindow_inst_hasCloseBox(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// HasTitleBar A Boolean value that indicates if the window has a title bar.
// https://developer.apple.com/documentation/appkit/nswindow/1449568-hastitlebar?language=objc
func (x gen_NSWindow) HasTitleBar() bool {
	ret := C.NSWindow_inst_hasTitleBar(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsModalPanel A Boolean value that indicates whether the window is a modal panel.
// https://developer.apple.com/documentation/appkit/nswindow/1449576-modalpanel?language=objc
func (x gen_NSWindow) IsModalPanel() bool {
	ret := C.NSWindow_inst_isModalPanel(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsFloatingPanel A Boolean value that indicates whether the window is a floating panel.
// https://developer.apple.com/documentation/appkit/nswindow/1449579-floatingpanel?language=objc
func (x gen_NSWindow) IsFloatingPanel() bool {
	ret := C.NSWindow_inst_isFloatingPanel(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsZoomable A Boolean value that indicates whether the window allows zooming.
// https://developer.apple.com/documentation/appkit/nswindow/1449587-zoomable?language=objc
func (x gen_NSWindow) IsZoomable() bool {
	ret := C.NSWindow_inst_isZoomable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsResizable A Boolean value that indicates if the user can resize the window.
// https://developer.apple.com/documentation/appkit/nswindow/1449572-resizable?language=objc
func (x gen_NSWindow) IsResizable() bool {
	ret := C.NSWindow_inst_isResizable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsMiniaturizable A Boolean value that indicates whether the window can minimize.
// https://developer.apple.com/documentation/appkit/nswindow/1449583-miniaturizable?language=objc
func (x gen_NSWindow) IsMiniaturizable() bool {
	ret := C.NSWindow_inst_isMiniaturizable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// OrderedIndex The zero-based position of the window, based on its order from front to back among all visible application windows.
// https://developer.apple.com/documentation/appkit/nswindow/1449577-orderedindex?language=objc
func (x gen_NSWindow) OrderedIndex() core.NSInteger {
	ret := C.NSWindow_inst_orderedIndex(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// SetOrderedIndex The zero-based position of the window, based on its order from front to back among all visible application windows.
// https://developer.apple.com/documentation/appkit/nswindow/1449577-orderedindex?language=objc
func (x gen_NSWindow) SetOrderedIndex(
	value core.NSInteger,
) {
	C.NSWindow_inst_setOrderedIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return

}

type NSWorkspaceRef interface {
	Pointer() uintptr
	Init_asNSWorkspace() NSWorkspace
}

type gen_NSWorkspace struct {
	objc.Object
}

func NSWorkspace_fromPointer(ptr unsafe.Pointer) NSWorkspace {
	return NSWorkspace{gen_NSWorkspace{
		objc.Object_fromPointer(ptr),
	}}
}

func NSWorkspace_fromRef(ref objc.Ref) NSWorkspace {
	return NSWorkspace_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// URLForApplicationToOpenURL Returns the URL to the default app that would be opened.
// https://developer.apple.com/documentation/appkit/nsworkspace/1533391-urlforapplicationtoopenurl?language=objc
func (x gen_NSWorkspace) URLForApplicationToOpenURL(
	url core.NSURLRef,
) core.NSURL {
	ret := C.NSWorkspace_inst_URLForApplicationToOpenURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)

	return core.NSURL_fromPointer(ret)

}

// URLForApplicationWithBundleIdentifier Returns the URL for the app with the specified identifier.
// https://developer.apple.com/documentation/appkit/nsworkspace/1534053-urlforapplicationwithbundleident?language=objc
func (x gen_NSWorkspace) URLForApplicationWithBundleIdentifier(
	bundleIdentifier core.NSStringRef,
) core.NSURL {
	ret := C.NSWorkspace_inst_URLForApplicationWithBundleIdentifier(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(bundleIdentifier),
	)

	return core.NSURL_fromPointer(ret)

}

// URLsForApplicationsToOpenURL
// https://developer.apple.com/documentation/appkit/nsworkspace/3753000-urlsforapplicationstoopenurl?language=objc
func (x gen_NSWorkspace) URLsForApplicationsToOpenURL(
	url core.NSURLRef,
) core.NSArray {
	ret := C.NSWorkspace_inst_URLsForApplicationsToOpenURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)

	return core.NSArray_fromPointer(ret)

}

// URLsForApplicationsWithBundleIdentifier
// https://developer.apple.com/documentation/appkit/nsworkspace/3753001-urlsforapplicationswithbundleide?language=objc
func (x gen_NSWorkspace) URLsForApplicationsWithBundleIdentifier(
	bundleIdentifier core.NSStringRef,
) core.NSArray {
	ret := C.NSWorkspace_inst_URLsForApplicationsWithBundleIdentifier(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(bundleIdentifier),
	)

	return core.NSArray_fromPointer(ret)

}

// ActivateFileViewerSelectingURLs Activates the Finder, and opens one or more windows selecting the specified files.
// https://developer.apple.com/documentation/appkit/nsworkspace/1524549-activatefileviewerselectingurls?language=objc
func (x gen_NSWorkspace) ActivateFileViewerSelectingURLs(
	fileURLs core.NSArrayRef,
) {
	C.NSWorkspace_inst_activateFileViewerSelectingURLs(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fileURLs),
	)

	return

}

// DesktopImageOptionsForScreen Returns the desktop image options for the given screen.
// https://developer.apple.com/documentation/appkit/nsworkspace/1530855-desktopimageoptionsforscreen?language=objc
func (x gen_NSWorkspace) DesktopImageOptionsForScreen(
	screen NSScreenRef,
) core.NSDictionary {
	ret := C.NSWorkspace_inst_desktopImageOptionsForScreen(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(screen),
	)

	return core.NSDictionary_fromPointer(ret)

}

// DesktopImageURLForScreen Returns the URL for the desktop image for the given screen.
// https://developer.apple.com/documentation/appkit/nsworkspace/1530635-desktopimageurlforscreen?language=objc
func (x gen_NSWorkspace) DesktopImageURLForScreen(
	screen NSScreenRef,
) core.NSURL {
	ret := C.NSWorkspace_inst_desktopImageURLForScreen(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(screen),
	)

	return core.NSURL_fromPointer(ret)

}

// ExtendPowerOffBy Requests the system wait for the specified amount of time before turning off the power or logging out the user.
// https://developer.apple.com/documentation/appkit/nsworkspace/1533106-extendpoweroffby?language=objc
func (x gen_NSWorkspace) ExtendPowerOffBy(
	requested core.NSInteger,
) core.NSInteger {
	ret := C.NSWorkspace_inst_extendPowerOffBy(
		unsafe.Pointer(x.Pointer()),
		C.long(requested),
	)

	return core.NSInteger(ret)

}

// HideOtherApplications Hides all applications other than the sender.
// https://developer.apple.com/documentation/appkit/nsworkspace/1530417-hideotherapplications?language=objc
func (x gen_NSWorkspace) HideOtherApplications() {
	C.NSWorkspace_inst_hideOtherApplications(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// IconForFile Returns an image containing the icon for the specified file.
// https://developer.apple.com/documentation/appkit/nsworkspace/1528158-iconforfile?language=objc
func (x gen_NSWorkspace) IconForFile(
	fullPath core.NSStringRef,
) NSImage {
	ret := C.NSWorkspace_inst_iconForFile(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fullPath),
	)

	return NSImage_fromPointer(ret)

}

// IconForFiles Returns an image containing the icon for the specified files.
// https://developer.apple.com/documentation/appkit/nsworkspace/1525487-iconforfiles?language=objc
func (x gen_NSWorkspace) IconForFiles(
	fullPaths core.NSArrayRef,
) NSImage {
	ret := C.NSWorkspace_inst_iconForFiles(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fullPaths),
	)

	return NSImage_fromPointer(ret)

}

// IsFilePackageAtPath Determines whether the specified path is a file package.
// https://developer.apple.com/documentation/appkit/nsworkspace/1529991-isfilepackageatpath?language=objc
func (x gen_NSWorkspace) IsFilePackageAtPath(
	fullPath core.NSStringRef,
) bool {
	ret := C.NSWorkspace_inst_isFilePackageAtPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fullPath),
	)

	return convertObjCBoolToGo(ret)

}

// NoteFileSystemChanged Informs the workspace object that the file system changed at the specified path.
// https://developer.apple.com/documentation/appkit/nsworkspace/1525376-notefilesystemchanged?language=objc
func (x gen_NSWorkspace) NoteFileSystemChanged(
	path core.NSStringRef,
) {
	C.NSWorkspace_inst_noteFileSystemChanged(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)

	return

}

// OpenURL Opens the location at the specified URL.
// https://developer.apple.com/documentation/appkit/nsworkspace/1533463-openurl?language=objc
func (x gen_NSWorkspace) OpenURL(
	url core.NSURLRef,
) bool {
	ret := C.NSWorkspace_inst_openURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)

	return convertObjCBoolToGo(ret)

}

// SelectFile_inFileViewerRootedAtPath Selects the file at the specified path.
// https://developer.apple.com/documentation/appkit/nsworkspace/1524399-selectfile?language=objc
func (x gen_NSWorkspace) SelectFile_inFileViewerRootedAtPath(
	fullPath core.NSStringRef,
	rootFullPath core.NSStringRef,
) bool {
	ret := C.NSWorkspace_inst_selectFile_inFileViewerRootedAtPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fullPath),
		objc.RefPointer(rootFullPath),
	)

	return convertObjCBoolToGo(ret)

}

// ShowSearchResultsForQueryString Displays a Spotlight search results window in Finder for the specified query string.
// https://developer.apple.com/documentation/appkit/nsworkspace/1532131-showsearchresultsforquerystring?language=objc
func (x gen_NSWorkspace) ShowSearchResultsForQueryString(
	queryString core.NSStringRef,
) bool {
	ret := C.NSWorkspace_inst_showSearchResultsForQueryString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(queryString),
	)

	return convertObjCBoolToGo(ret)

}

// UnmountAndEjectDeviceAtPath Unmounts and ejects the device at the specified path.
// https://developer.apple.com/documentation/appkit/nsworkspace/1527741-unmountandejectdeviceatpath?language=objc
func (x gen_NSWorkspace) UnmountAndEjectDeviceAtPath(
	path core.NSStringRef,
) bool {
	ret := C.NSWorkspace_inst_unmountAndEjectDeviceAtPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)

	return convertObjCBoolToGo(ret)

}

// Init_asNSWorkspace
func (x gen_NSWorkspace) Init_asNSWorkspace() NSWorkspace {
	ret := C.NSWorkspace_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSWorkspace_fromPointer(ret)

}

// FrontmostApplication Returns the frontmost app, which is the app that receives key events.
// https://developer.apple.com/documentation/appkit/nsworkspace/1532097-frontmostapplication?language=objc
func (x gen_NSWorkspace) FrontmostApplication() NSRunningApplication {
	ret := C.NSWorkspace_inst_frontmostApplication(
		unsafe.Pointer(x.Pointer()),
	)

	return NSRunningApplication_fromPointer(ret)

}

// RunningApplications Returns an array of running apps.
// https://developer.apple.com/documentation/appkit/nsworkspace/1534059-runningapplications?language=objc
func (x gen_NSWorkspace) RunningApplications() core.NSArray {
	ret := C.NSWorkspace_inst_runningApplications(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// MenuBarOwningApplication Returns the app that owns the currently displayed menu bar.
// https://developer.apple.com/documentation/appkit/nsworkspace/1525848-menubarowningapplication?language=objc
func (x gen_NSWorkspace) MenuBarOwningApplication() NSRunningApplication {
	ret := C.NSWorkspace_inst_menuBarOwningApplication(
		unsafe.Pointer(x.Pointer()),
	)

	return NSRunningApplication_fromPointer(ret)

}

// FileLabels The array of file labels, returned as strings.
// https://developer.apple.com/documentation/appkit/nsworkspace/1533953-filelabels?language=objc
func (x gen_NSWorkspace) FileLabels() core.NSArray {
	ret := C.NSWorkspace_inst_fileLabels(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// FileLabelColors The array of colors for the file labels.
// https://developer.apple.com/documentation/appkit/nsworkspace/1527553-filelabelcolors?language=objc
func (x gen_NSWorkspace) FileLabelColors() core.NSArray {
	ret := C.NSWorkspace_inst_fileLabelColors(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// AccessibilityDisplayShouldDifferentiateWithoutColor A Boolean value that indicates whether the app avoids conveying information through color alone.
// https://developer.apple.com/documentation/appkit/nsworkspace/1524656-accessibilitydisplayshoulddiffer?language=objc
func (x gen_NSWorkspace) AccessibilityDisplayShouldDifferentiateWithoutColor() bool {
	ret := C.NSWorkspace_inst_accessibilityDisplayShouldDifferentiateWithoutColor(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AccessibilityDisplayShouldIncreaseContrast A Boolean value that indicates whether the app presents a high-contrast user interface.
// https://developer.apple.com/documentation/appkit/nsworkspace/1526290-accessibilitydisplayshouldincrea?language=objc
func (x gen_NSWorkspace) AccessibilityDisplayShouldIncreaseContrast() bool {
	ret := C.NSWorkspace_inst_accessibilityDisplayShouldIncreaseContrast(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AccessibilityDisplayShouldReduceTransparency A Boolean value that indicates whether the app avoids using semitransparent backgrounds.
// https://developer.apple.com/documentation/appkit/nsworkspace/1533006-accessibilitydisplayshouldreduce?language=objc
func (x gen_NSWorkspace) AccessibilityDisplayShouldReduceTransparency() bool {
	ret := C.NSWorkspace_inst_accessibilityDisplayShouldReduceTransparency(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AccessibilityDisplayShouldInvertColors A Boolean value that indicates whether the accessibility option to invert colors is in an enabled state.
// https://developer.apple.com/documentation/appkit/nsworkspace/1644068-accessibilitydisplayshouldinvert?language=objc
func (x gen_NSWorkspace) AccessibilityDisplayShouldInvertColors() bool {
	ret := C.NSWorkspace_inst_accessibilityDisplayShouldInvertColors(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AccessibilityDisplayShouldReduceMotion A Boolean value that indicates whether the accessibility option to reduce motion is in an enabled state.
// https://developer.apple.com/documentation/appkit/nsworkspace/1644069-accessibilitydisplayshouldreduce?language=objc
func (x gen_NSWorkspace) AccessibilityDisplayShouldReduceMotion() bool {
	ret := C.NSWorkspace_inst_accessibilityDisplayShouldReduceMotion(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsSwitchControlEnabled A Boolean value that indicates whether Switch Control is currently running.
// https://developer.apple.com/documentation/appkit/nsworkspace/2880322-switchcontrolenabled?language=objc
func (x gen_NSWorkspace) IsSwitchControlEnabled() bool {
	ret := C.NSWorkspace_inst_isSwitchControlEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsVoiceOverEnabled A Boolean value that indicates whether VoiceOver is currently running.
// https://developer.apple.com/documentation/appkit/nsworkspace/2880317-voiceoverenabled?language=objc
func (x gen_NSWorkspace) IsVoiceOverEnabled() bool {
	ret := C.NSWorkspace_inst_isVoiceOverEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

type NSColorRef interface {
	Pointer() uintptr
	Init_asNSColor() NSColor
}

type gen_NSColor struct {
	objc.Object
}

func NSColor_fromPointer(ptr unsafe.Pointer) NSColor {
	return NSColor{gen_NSColor{
		objc.Object_fromPointer(ptr),
	}}
}

func NSColor_fromRef(ref objc.Ref) NSColor {
	return NSColor_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// BlendedColorWithFraction_ofColor Creates a new color object whose component values are a weighted sum of the current color object and the specified color object's.
// https://developer.apple.com/documentation/appkit/nscolor/1524689-blendedcolorwithfraction?language=objc
func (x gen_NSColor) BlendedColorWithFraction_ofColor(
	fraction core.CGFloat,
	color NSColorRef,
) NSColor {
	ret := C.NSColor_inst_blendedColorWithFraction_ofColor(
		unsafe.Pointer(x.Pointer()),
		C.double(fraction),
		objc.RefPointer(color),
	)

	return NSColor_fromPointer(ret)

}

// ColorWithAlphaComponent Creates a new color object that has the same color space and component values as the current color object, but the specified alpha component.
// https://developer.apple.com/documentation/appkit/nscolor/1526906-colorwithalphacomponent?language=objc
func (x gen_NSColor) ColorWithAlphaComponent(
	alpha core.CGFloat,
) NSColor {
	ret := C.NSColor_inst_colorWithAlphaComponent(
		unsafe.Pointer(x.Pointer()),
		C.double(alpha),
	)

	return NSColor_fromPointer(ret)

}

// DrawSwatchInRect Draws the current color in the specified rectangle.
// https://developer.apple.com/documentation/appkit/nscolor/1531770-drawswatchinrect?language=objc
func (x gen_NSColor) DrawSwatchInRect(
	rect core.NSRect,
) {
	C.NSColor_inst_drawSwatchInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return

}

// HighlightWithLevel Creates a new color object that represents a blend between the current color and the highlight color.
// https://developer.apple.com/documentation/appkit/nscolor/1533061-highlightwithlevel?language=objc
func (x gen_NSColor) HighlightWithLevel(
	val core.CGFloat,
) NSColor {
	ret := C.NSColor_inst_highlightWithLevel(
		unsafe.Pointer(x.Pointer()),
		C.double(val),
	)

	return NSColor_fromPointer(ret)

}

// Set Sets the color of subsequent drawing to the color that the color object represents.
// https://developer.apple.com/documentation/appkit/nscolor/1527089-set?language=objc
func (x gen_NSColor) Set() {
	C.NSColor_inst_set(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// SetFill Sets the fill color of subsequent drawing to the color object’s color.
// https://developer.apple.com/documentation/appkit/nscolor/1524755-setfill?language=objc
func (x gen_NSColor) SetFill() {
	C.NSColor_inst_setFill(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// SetStroke Sets the stroke color of subsequent drawing to the color object’s color.
// https://developer.apple.com/documentation/appkit/nscolor/1531019-setstroke?language=objc
func (x gen_NSColor) SetStroke() {
	C.NSColor_inst_setStroke(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ShadowWithLevel Creates a new color object that represents a blend between the current color and the shadow color.
// https://developer.apple.com/documentation/appkit/nscolor/1528523-shadowwithlevel?language=objc
func (x gen_NSColor) ShadowWithLevel(
	val core.CGFloat,
) NSColor {
	ret := C.NSColor_inst_shadowWithLevel(
		unsafe.Pointer(x.Pointer()),
		C.double(val),
	)

	return NSColor_fromPointer(ret)

}

// WriteToPasteboard Writes the color object’s data to the specified pasteboard.
// https://developer.apple.com/documentation/appkit/nscolor/1532199-writetopasteboard?language=objc
func (x gen_NSColor) WriteToPasteboard(
	pasteBoard NSPasteboardRef,
) {
	C.NSColor_inst_writeToPasteboard(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pasteBoard),
	)

	return

}

// Init_asNSColor
func (x gen_NSColor) Init_asNSColor() NSColor {
	ret := C.NSColor_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_fromPointer(ret)

}

// NumberOfComponents The number of components in the color.
// https://developer.apple.com/documentation/appkit/nscolor/1531308-numberofcomponents?language=objc
func (x gen_NSColor) NumberOfComponents() core.NSInteger {
	ret := C.NSColor_inst_numberOfComponents(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// AlphaComponent The alpha (opacity) component value of the color.
// https://developer.apple.com/documentation/appkit/nscolor/1532504-alphacomponent?language=objc
func (x gen_NSColor) AlphaComponent() core.CGFloat {
	ret := C.NSColor_inst_alphaComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// WhiteComponent The white component value of the color.
// https://developer.apple.com/documentation/appkit/nscolor/1534051-whitecomponent?language=objc
func (x gen_NSColor) WhiteComponent() core.CGFloat {
	ret := C.NSColor_inst_whiteComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// RedComponent The red component value of the color.
// https://developer.apple.com/documentation/appkit/nscolor/1530483-redcomponent?language=objc
func (x gen_NSColor) RedComponent() core.CGFloat {
	ret := C.NSColor_inst_redComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// GreenComponent The green component value of the color.
// https://developer.apple.com/documentation/appkit/nscolor/1525935-greencomponent?language=objc
func (x gen_NSColor) GreenComponent() core.CGFloat {
	ret := C.NSColor_inst_greenComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// BlueComponent The blue component value of the color.
// https://developer.apple.com/documentation/appkit/nscolor/1534229-bluecomponent?language=objc
func (x gen_NSColor) BlueComponent() core.CGFloat {
	ret := C.NSColor_inst_blueComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// CyanComponent The cyan component value of the color.
// https://developer.apple.com/documentation/appkit/nscolor/1528234-cyancomponent?language=objc
func (x gen_NSColor) CyanComponent() core.CGFloat {
	ret := C.NSColor_inst_cyanComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// MagentaComponent The magenta component value of the color.
// https://developer.apple.com/documentation/appkit/nscolor/1535560-magentacomponent?language=objc
func (x gen_NSColor) MagentaComponent() core.CGFloat {
	ret := C.NSColor_inst_magentaComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// YellowComponent The yellow component value of the color.
// https://developer.apple.com/documentation/appkit/nscolor/1531965-yellowcomponent?language=objc
func (x gen_NSColor) YellowComponent() core.CGFloat {
	ret := C.NSColor_inst_yellowComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// BlackComponent The black component value of the color.
// https://developer.apple.com/documentation/appkit/nscolor/1526883-blackcomponent?language=objc
func (x gen_NSColor) BlackComponent() core.CGFloat {
	ret := C.NSColor_inst_blackComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// HueComponent The hue component value of the color.
// https://developer.apple.com/documentation/appkit/nscolor/1531780-huecomponent?language=objc
func (x gen_NSColor) HueComponent() core.CGFloat {
	ret := C.NSColor_inst_hueComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// SaturationComponent The saturation component value of the color.
// https://developer.apple.com/documentation/appkit/nscolor/1526326-saturationcomponent?language=objc
func (x gen_NSColor) SaturationComponent() core.CGFloat {
	ret := C.NSColor_inst_saturationComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// BrightnessComponent The brightness component value of the color.
// https://developer.apple.com/documentation/appkit/nscolor/1529355-brightnesscomponent?language=objc
func (x gen_NSColor) BrightnessComponent() core.CGFloat {
	ret := C.NSColor_inst_brightnessComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// LocalizedCatalogNameComponent The localized version of the catalog name containing the color.
// https://developer.apple.com/documentation/appkit/nscolor/1535351-localizedcatalognamecomponent?language=objc
func (x gen_NSColor) LocalizedCatalogNameComponent() core.NSString {
	ret := C.NSColor_inst_localizedCatalogNameComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// LocalizedColorNameComponent The localized version of the color name.
// https://developer.apple.com/documentation/appkit/nscolor/1527286-localizedcolornamecomponent?language=objc
func (x gen_NSColor) LocalizedColorNameComponent() core.NSString {
	ret := C.NSColor_inst_localizedColorNameComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

type NSTextViewRef interface {
	Pointer() uintptr
	Init_asNSTextView() NSTextView
}

type gen_NSTextView struct {
	NSText
}

func NSTextView_fromPointer(ptr unsafe.Pointer) NSTextView {
	return NSTextView{gen_NSTextView{
		NSText_fromPointer(ptr),
	}}
}

func NSTextView_fromRef(ref objc.Ref) NSTextView {
	return NSTextView_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// AlignJustified Applies full justification to selected paragraphs (or all text, if the receiver is a plain text object).
// https://developer.apple.com/documentation/appkit/nstextview/1449515-alignjustified?language=objc
func (x gen_NSTextView) AlignJustified(
	sender objc.Ref,
) {
	C.NSTextView_inst_alignJustified(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// BreakUndoCoalescing Informs the receiver that it should begin coalescing successive typing operations in a new undo grouping.
// https://developer.apple.com/documentation/appkit/nstextview/1449384-breakundocoalescing?language=objc
func (x gen_NSTextView) BreakUndoCoalescing() {
	C.NSTextView_inst_breakUndoCoalescing(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ChangeAttributes Changes the attributes of the current selection.
// https://developer.apple.com/documentation/appkit/nstextview/1449216-changeattributes?language=objc
func (x gen_NSTextView) ChangeAttributes(
	sender objc.Ref,
) {
	C.NSTextView_inst_changeAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ChangeColor Sets the color of the selected text.
// https://developer.apple.com/documentation/appkit/nstextview/1449282-changecolor?language=objc
func (x gen_NSTextView) ChangeColor(
	sender objc.Ref,
) {
	C.NSTextView_inst_changeColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ChangeDocumentBackgroundColor An action method used to set the background color.
// https://developer.apple.com/documentation/appkit/nstextview/1449475-changedocumentbackgroundcolor?language=objc
func (x gen_NSTextView) ChangeDocumentBackgroundColor(
	sender objc.Ref,
) {
	C.NSTextView_inst_changeDocumentBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ChangeLayoutOrientation An action method that sets the layout orientation of the text.
// https://developer.apple.com/documentation/appkit/nstextview/1449286-changelayoutorientation?language=objc
func (x gen_NSTextView) ChangeLayoutOrientation(
	sender objc.Ref,
) {
	C.NSTextView_inst_changeLayoutOrientation(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// CharacterIndexForInsertionAtPoint Returns a character index appropriate for placing a zero-length selection for an insertion point associated with the mouse at the given point.
// https://developer.apple.com/documentation/appkit/nstextview/1449505-characterindexforinsertionatpoin?language=objc
func (x gen_NSTextView) CharacterIndexForInsertionAtPoint(
	point core.NSPoint,
) core.NSUInteger {
	ret := C.NSTextView_inst_characterIndexForInsertionAtPoint(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return core.NSUInteger(ret)

}

// CheckTextInDocument Performs the default text checking on the entire document.
// https://developer.apple.com/documentation/appkit/nstextview/1449440-checktextindocument?language=objc
func (x gen_NSTextView) CheckTextInDocument(
	sender objc.Ref,
) {
	C.NSTextView_inst_checkTextInDocument(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// CheckTextInSelection Performs the default text checking on the current selection.
// https://developer.apple.com/documentation/appkit/nstextview/1449382-checktextinselection?language=objc
func (x gen_NSTextView) CheckTextInSelection(
	sender objc.Ref,
) {
	C.NSTextView_inst_checkTextInSelection(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// CleanUpAfterDragOperation Releases the drag information still existing after the dragging session has completed.
// https://developer.apple.com/documentation/appkit/nstextview/1449202-cleanupafterdragoperation?language=objc
func (x gen_NSTextView) CleanUpAfterDragOperation() {
	C.NSTextView_inst_cleanUpAfterDragOperation(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ClickedOnLink_atIndex Causes the text view to act as if the user clicked on some text with the given link as the value of a link attribute associated with the text.
// https://developer.apple.com/documentation/appkit/nstextview/1449497-clickedonlink?language=objc
func (x gen_NSTextView) ClickedOnLink_atIndex(
	link objc.Ref,
	charIndex core.NSUInteger,
) {
	C.NSTextView_inst_clickedOnLink_atIndex(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(link),
		C.ulong(charIndex),
	)

	return

}

// Complete Invokes completion in a text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449359-complete?language=objc
func (x gen_NSTextView) Complete(
	sender objc.Ref,
) {
	C.NSTextView_inst_complete(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// DidChangeText Sends out necessary notifications when a text change completes.
// https://developer.apple.com/documentation/appkit/nstextview/1449296-didchangetext?language=objc
func (x gen_NSTextView) DidChangeText() {
	C.NSTextView_inst_didChangeText(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// DragSelectionWithEvent_offset_slideBack Begins dragging the current selected text range.
// https://developer.apple.com/documentation/appkit/nstextview/1449413-dragselectionwithevent?language=objc
func (x gen_NSTextView) DragSelectionWithEvent_offset_slideBack(
	event NSEventRef,
	mouseOffset core.NSSize,
	slideBack bool,
) bool {
	ret := C.NSTextView_inst_dragSelectionWithEvent_offset_slideBack(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
		*(*C.NSSize)(unsafe.Pointer(&mouseOffset)),
		convertToObjCBool(slideBack),
	)

	return convertObjCBoolToGo(ret)

}

// DrawInsertionPointInRect_color_turnedOn Draws or erases the insertion point.
// https://developer.apple.com/documentation/appkit/nstextview/1449232-drawinsertionpointinrect?language=objc
func (x gen_NSTextView) DrawInsertionPointInRect_color_turnedOn(
	rect core.NSRect,
	color NSColorRef,
	flag bool,
) {
	C.NSTextView_inst_drawInsertionPointInRect_color_turnedOn(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(color),
		convertToObjCBool(flag),
	)

	return

}

// DrawViewBackgroundInRect Draws the background of the text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449135-drawviewbackgroundinrect?language=objc
func (x gen_NSTextView) DrawViewBackgroundInRect(
	rect core.NSRect,
) {
	C.NSTextView_inst_drawViewBackgroundInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return

}

// InitWithFrame_asNSTextView Initializes a text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449262-initwithframe?language=objc
func (x gen_NSTextView) InitWithFrame_asNSTextView(
	frameRect core.NSRect,
) NSTextView {
	ret := C.NSTextView_inst_initWithFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)

	return NSTextView_fromPointer(ret)

}

// InitWithFrame_textContainer_asNSTextView Initializes a text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449347-initwithframe?language=objc
func (x gen_NSTextView) InitWithFrame_textContainer_asNSTextView(
	frameRect core.NSRect,
	container NSTextContainerRef,
) NSTextView {
	ret := C.NSTextView_inst_initWithFrame_textContainer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
		objc.RefPointer(container),
	)

	return NSTextView_fromPointer(ret)

}

// InvalidateTextContainerOrigin Invalidates the calculated origin of the text container.
// https://developer.apple.com/documentation/appkit/nstextview/1449546-invalidatetextcontainerorigin?language=objc
func (x gen_NSTextView) InvalidateTextContainerOrigin() {
	C.NSTextView_inst_invalidateTextContainerOrigin(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// LoosenKerning Increases the space between glyphs in the receiver’s selection, or in all text if the receiver is a plain text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449183-loosenkerning?language=objc
func (x gen_NSTextView) LoosenKerning(
	sender objc.Ref,
) {
	C.NSTextView_inst_loosenKerning(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// LowerBaseline Lowers the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449289-lowerbaseline?language=objc
func (x gen_NSTextView) LowerBaseline(
	sender objc.Ref,
) {
	C.NSTextView_inst_lowerBaseline(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// OrderFrontLinkPanel Brings forward a panel allowing the user to manipulate links in the text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449238-orderfrontlinkpanel?language=objc
func (x gen_NSTextView) OrderFrontLinkPanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontLinkPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// OrderFrontListPanel Brings forward a panel allowing the user to manipulate text lists in the text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449349-orderfrontlistpanel?language=objc
func (x gen_NSTextView) OrderFrontListPanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontListPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// OrderFrontSharingServicePicker Creates and displays a new instance of the sharing service picker.
// https://developer.apple.com/documentation/appkit/nstextview/1449150-orderfrontsharingservicepicker?language=objc
func (x gen_NSTextView) OrderFrontSharingServicePicker(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontSharingServicePicker(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// OrderFrontSpacingPanel Brings forward a panel allowing the user to manipulate text line heights, interline spacing, and paragraph spacing, in the text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449438-orderfrontspacingpanel?language=objc
func (x gen_NSTextView) OrderFrontSpacingPanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontSpacingPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// OrderFrontSubstitutionsPanel Brings forward a panel allowing the user to specify string substitutions in the text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449327-orderfrontsubstitutionspanel?language=objc
func (x gen_NSTextView) OrderFrontSubstitutionsPanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontSubstitutionsPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// OrderFrontTablePanel Brings forward a panel allowing the user to manipulate text tables in the text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449442-orderfronttablepanel?language=objc
func (x gen_NSTextView) OrderFrontTablePanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontTablePanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// Outline Adds the outline attribute to the selected text attributes if absent; removes the attribute if present.
// https://developer.apple.com/documentation/appkit/nstextview/1449386-outline?language=objc
func (x gen_NSTextView) Outline(
	sender objc.Ref,
) {
	C.NSTextView_inst_outline(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// PasteAsPlainText Inserts the contents of the pasteboard into the receiver’s text as plain text.
// https://developer.apple.com/documentation/appkit/nstextview/1449250-pasteasplaintext?language=objc
func (x gen_NSTextView) PasteAsPlainText(
	sender objc.Ref,
) {
	C.NSTextView_inst_pasteAsPlainText(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// PasteAsRichText This action method inserts the contents of the pasteboard into the receiver’s text as rich text, maintaining its attributes.
// https://developer.apple.com/documentation/appkit/nstextview/1449395-pasteasrichtext?language=objc
func (x gen_NSTextView) PasteAsRichText(
	sender objc.Ref,
) {
	C.NSTextView_inst_pasteAsRichText(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// PerformFindPanelAction Performs a find panel action specified by the sender's tag.
// https://developer.apple.com/documentation/appkit/nstextview/1449525-performfindpanelaction?language=objc
func (x gen_NSTextView) PerformFindPanelAction(
	sender objc.Ref,
) {
	C.NSTextView_inst_performFindPanelAction(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// QuickLookPreviewableItemsInRanges Returns an array of URLs for items that can be displayed by QuickLook in the specified ranges.
// https://developer.apple.com/documentation/appkit/nstextview/1449426-quicklookpreviewableitemsinrange?language=objc
func (x gen_NSTextView) QuickLookPreviewableItemsInRanges(
	ranges core.NSArrayRef,
) core.NSArray {
	ret := C.NSTextView_inst_quickLookPreviewableItemsInRanges(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ranges),
	)

	return core.NSArray_fromPointer(ret)

}

// RaiseBaseline Raises the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449198-raisebaseline?language=objc
func (x gen_NSTextView) RaiseBaseline(
	sender objc.Ref,
) {
	C.NSTextView_inst_raiseBaseline(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ReadSelectionFromPasteboard Reads the text view’s preferred type of data from the specified pasteboard.
// https://developer.apple.com/documentation/appkit/nstextview/1449469-readselectionfrompasteboard?language=objc
func (x gen_NSTextView) ReadSelectionFromPasteboard(
	pboard NSPasteboardRef,
) bool {
	ret := C.NSTextView_inst_readSelectionFromPasteboard(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pboard),
	)

	return convertObjCBoolToGo(ret)

}

// ReplaceTextContainer Replaces the text container for the group of text system objects containing the receiver, keeping the association between the receiver and its layout manager intact.
// https://developer.apple.com/documentation/appkit/nstextview/1449479-replacetextcontainer?language=objc
func (x gen_NSTextView) ReplaceTextContainer(
	newContainer NSTextContainerRef,
) {
	C.NSTextView_inst_replaceTextContainer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newContainer),
	)

	return

}

// SetConstrainedFrameSize Attempts to set the frame size as if by user action.
// https://developer.apple.com/documentation/appkit/nstextview/1449230-setconstrainedframesize?language=objc
func (x gen_NSTextView) SetConstrainedFrameSize(
	desiredSize core.NSSize,
) {
	C.NSTextView_inst_setConstrainedFrameSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&desiredSize)),
	)

	return

}

// SetNeedsDisplayInRect_avoidAdditionalLayout Marks the receiver as requiring display.
// https://developer.apple.com/documentation/appkit/nstextview/1449279-setneedsdisplayinrect?language=objc
func (x gen_NSTextView) SetNeedsDisplayInRect_avoidAdditionalLayout(
	rect core.NSRect,
	flag bool,
) {
	C.NSTextView_inst_setNeedsDisplayInRect_avoidAdditionalLayout(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		convertToObjCBool(flag),
	)

	return

}

// ShouldChangeTextInRanges_replacementStrings Initiates a series of delegate messages (and general notifications) to determine whether modifications can be made to the characters and attributes of the receiver’s text.
// https://developer.apple.com/documentation/appkit/nstextview/1449311-shouldchangetextinranges?language=objc
func (x gen_NSTextView) ShouldChangeTextInRanges_replacementStrings(
	affectedRanges core.NSArrayRef,
	replacementStrings core.NSArrayRef,
) bool {
	ret := C.NSTextView_inst_shouldChangeTextInRanges_replacementStrings(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(affectedRanges),
		objc.RefPointer(replacementStrings),
	)

	return convertObjCBoolToGo(ret)

}

// StartSpeaking Speaks the selected text, or all text if no selection.
// https://developer.apple.com/documentation/appkit/nstextview/1449519-startspeaking?language=objc
func (x gen_NSTextView) StartSpeaking(
	sender objc.Ref,
) {
	C.NSTextView_inst_startSpeaking(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// StopSpeaking Stops the speaking of text.
// https://developer.apple.com/documentation/appkit/nstextview/1449172-stopspeaking?language=objc
func (x gen_NSTextView) StopSpeaking(
	sender objc.Ref,
) {
	C.NSTextView_inst_stopSpeaking(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// TightenKerning Decreases the space between glyphs in the receiver’s selection, or for all glyphs if the receiver is a plain text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449137-tightenkerning?language=objc
func (x gen_NSTextView) TightenKerning(
	sender objc.Ref,
) {
	C.NSTextView_inst_tightenKerning(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleAutomaticDashSubstitution Toggles the state of the automatic dash substitution.
// https://developer.apple.com/documentation/appkit/nstextview/1449305-toggleautomaticdashsubstitution?language=objc
func (x gen_NSTextView) ToggleAutomaticDashSubstitution(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticDashSubstitution(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleAutomaticDataDetection Toggles the state of the automatic data detection.
// https://developer.apple.com/documentation/appkit/nstextview/1449499-toggleautomaticdatadetection?language=objc
func (x gen_NSTextView) ToggleAutomaticDataDetection(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticDataDetection(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleAutomaticLinkDetection Changes the state of automatic link detection from enabled to disabled and vice versa.
// https://developer.apple.com/documentation/appkit/nstextview/1449353-toggleautomaticlinkdetection?language=objc
func (x gen_NSTextView) ToggleAutomaticLinkDetection(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticLinkDetection(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleAutomaticQuoteSubstitution Changes the state of automatic quotation mark substitution from enabled to disabled and vice versa.
// https://developer.apple.com/documentation/appkit/nstextview/1449444-toggleautomaticquotesubstitution?language=objc
func (x gen_NSTextView) ToggleAutomaticQuoteSubstitution(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticQuoteSubstitution(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleAutomaticSpellingCorrection Toggles the state of the automatic spelling correction.
// https://developer.apple.com/documentation/appkit/nstextview/1449178-toggleautomaticspellingcorrectio?language=objc
func (x gen_NSTextView) ToggleAutomaticSpellingCorrection(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticSpellingCorrection(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleAutomaticTextCompletion
// https://developer.apple.com/documentation/appkit/nstextview/2544841-toggleautomatictextcompletion?language=objc
func (x gen_NSTextView) ToggleAutomaticTextCompletion(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticTextCompletion(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleAutomaticTextReplacement Toggles the state of the automatic text replacement.
// https://developer.apple.com/documentation/appkit/nstextview/1449200-toggleautomatictextreplacement?language=objc
func (x gen_NSTextView) ToggleAutomaticTextReplacement(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticTextReplacement(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleContinuousSpellChecking Toggles whether continuous spell checking is enabled for the receiver.
// https://developer.apple.com/documentation/appkit/nstextview/1449471-togglecontinuousspellchecking?language=objc
func (x gen_NSTextView) ToggleContinuousSpellChecking(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleContinuousSpellChecking(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleGrammarChecking Changes the state of grammar checking from enabled to disabled and vice versa.
// https://developer.apple.com/documentation/appkit/nstextview/1449393-togglegrammarchecking?language=objc
func (x gen_NSTextView) ToggleGrammarChecking(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleGrammarChecking(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleQuickLookPreviewPanel An action message that toggles the visibility state of the Quick Look preview panel.
// https://developer.apple.com/documentation/appkit/nstextview/1449415-togglequicklookpreviewpanel?language=objc
func (x gen_NSTextView) ToggleQuickLookPreviewPanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleQuickLookPreviewPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ToggleSmartInsertDelete Changes the state of smart insert and delete from enabled to disabled and vice versa.
// https://developer.apple.com/documentation/appkit/nstextview/1449273-togglesmartinsertdelete?language=objc
func (x gen_NSTextView) ToggleSmartInsertDelete(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleSmartInsertDelete(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// TurnOffKerning Sets the receiver to use nominal glyph spacing for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449464-turnoffkerning?language=objc
func (x gen_NSTextView) TurnOffKerning(
	sender objc.Ref,
) {
	C.NSTextView_inst_turnOffKerning(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// TurnOffLigatures Sets the receiver to use only required ligatures when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449436-turnoffligatures?language=objc
func (x gen_NSTextView) TurnOffLigatures(
	sender objc.Ref,
) {
	C.NSTextView_inst_turnOffLigatures(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// UpdateCandidates
// https://developer.apple.com/documentation/appkit/nstextview/2544833-updatecandidates?language=objc
func (x gen_NSTextView) UpdateCandidates() {
	C.NSTextView_inst_updateCandidates(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// UpdateDragTypeRegistration Updates the acceptable drag types of all text views associated with the receiver's layout manager.
// https://developer.apple.com/documentation/appkit/nstextview/1449181-updatedragtyperegistration?language=objc
func (x gen_NSTextView) UpdateDragTypeRegistration() {
	C.NSTextView_inst_updateDragTypeRegistration(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// UpdateFontPanel Updates the Font panel to contain the font attributes of the selection.
// https://developer.apple.com/documentation/appkit/nstextview/1449401-updatefontpanel?language=objc
func (x gen_NSTextView) UpdateFontPanel() {
	C.NSTextView_inst_updateFontPanel(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// UpdateInsertionPointStateAndRestartTimer Updates the insertion point’s location and optionally restarts the blinking cursor timer.
// https://developer.apple.com/documentation/appkit/nstextview/1449268-updateinsertionpointstateandrest?language=objc
func (x gen_NSTextView) UpdateInsertionPointStateAndRestartTimer(
	restartFlag bool,
) {
	C.NSTextView_inst_updateInsertionPointStateAndRestartTimer(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(restartFlag),
	)

	return

}

// UpdateQuickLookPreviewPanel Notifies the QuickLook panel that an update may be required.
// https://developer.apple.com/documentation/appkit/nstextview/1449409-updatequicklookpreviewpanel?language=objc
func (x gen_NSTextView) UpdateQuickLookPreviewPanel() {
	C.NSTextView_inst_updateQuickLookPreviewPanel(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// UpdateRuler Updates the ruler view in the receiver’s enclosing scroll view to reflect the selection’s paragraph and marker attributes.
// https://developer.apple.com/documentation/appkit/nstextview/1449323-updateruler?language=objc
func (x gen_NSTextView) UpdateRuler() {
	C.NSTextView_inst_updateRuler(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// UpdateTextTouchBarItems
// https://developer.apple.com/documentation/appkit/nstextview/2544676-updatetexttouchbaritems?language=objc
func (x gen_NSTextView) UpdateTextTouchBarItems() {
	C.NSTextView_inst_updateTextTouchBarItems(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// UpdateTouchBarItemIdentifiers
// https://developer.apple.com/documentation/appkit/nstextview/2544834-updatetouchbaritemidentifiers?language=objc
func (x gen_NSTextView) UpdateTouchBarItemIdentifiers() {
	C.NSTextView_inst_updateTouchBarItemIdentifiers(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// UseAllLigatures Sets the receiver to use all ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449213-useallligatures?language=objc
func (x gen_NSTextView) UseAllLigatures(
	sender objc.Ref,
) {
	C.NSTextView_inst_useAllLigatures(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// UseStandardKerning Set the receiver to use pair kerning data for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449491-usestandardkerning?language=objc
func (x gen_NSTextView) UseStandardKerning(
	sender objc.Ref,
) {
	C.NSTextView_inst_useStandardKerning(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// UseStandardLigatures Sets the receiver to use the standard ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449144-usestandardligatures?language=objc
func (x gen_NSTextView) UseStandardLigatures(
	sender objc.Ref,
) {
	C.NSTextView_inst_useStandardLigatures(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// WriteSelectionToPasteboard_types Writes the current selection to the specified pasteboard under each given type.
// https://developer.apple.com/documentation/appkit/nstextview/1449277-writeselectiontopasteboard?language=objc
func (x gen_NSTextView) WriteSelectionToPasteboard_types(
	pboard NSPasteboardRef,
	types core.NSArrayRef,
) bool {
	ret := C.NSTextView_inst_writeSelectionToPasteboard_types(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pboard),
		objc.RefPointer(types),
	)

	return convertObjCBoolToGo(ret)

}

// Init_asNSTextView
func (x gen_NSTextView) Init_asNSTextView() NSTextView {
	ret := C.NSTextView_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextView_fromPointer(ret)

}

// Delegate The delegate for all text views sharing the receiver’s layout manager.
// https://developer.apple.com/documentation/appkit/nstextview/1449521-delegate?language=objc
func (x gen_NSTextView) Delegate() objc.Object {
	ret := C.NSTextView_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetDelegate The delegate for all text views sharing the receiver’s layout manager.
// https://developer.apple.com/documentation/appkit/nstextview/1449521-delegate?language=objc
func (x gen_NSTextView) SetDelegate(
	value objc.Ref,
) {
	C.NSTextView_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// TextContainer The receiver’s text container.
// https://developer.apple.com/documentation/appkit/nstextview/1449364-textcontainer?language=objc
func (x gen_NSTextView) TextContainer() NSTextContainer {
	ret := C.NSTextView_inst_textContainer(
		unsafe.Pointer(x.Pointer()),
	)

	return NSTextContainer_fromPointer(ret)

}

// SetTextContainer The receiver’s text container.
// https://developer.apple.com/documentation/appkit/nstextview/1449364-textcontainer?language=objc
func (x gen_NSTextView) SetTextContainer(
	value NSTextContainerRef,
) {
	C.NSTextView_inst_setTextContainer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// TextContainerInset The empty space the receiver leaves around its associated text container.
// https://developer.apple.com/documentation/appkit/nstextview/1449168-textcontainerinset?language=objc
func (x gen_NSTextView) TextContainerInset() core.NSSize {
	ret := C.NSTextView_inst_textContainerInset(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// SetTextContainerInset The empty space the receiver leaves around its associated text container.
// https://developer.apple.com/documentation/appkit/nstextview/1449168-textcontainerinset?language=objc
func (x gen_NSTextView) SetTextContainerInset(
	value core.NSSize,
) {
	C.NSTextView_inst_setTextContainerInset(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// TextContainerOrigin The origin of the receiver’s text container.
// https://developer.apple.com/documentation/appkit/nstextview/1449477-textcontainerorigin?language=objc
func (x gen_NSTextView) TextContainerOrigin() core.NSPoint {
	ret := C.NSTextView_inst_textContainerOrigin(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// LayoutManager The layout manager that lays out text for the receiver’s text container.
// https://developer.apple.com/documentation/appkit/nstextview/1449148-layoutmanager?language=objc
func (x gen_NSTextView) LayoutManager() NSLayoutManager {
	ret := C.NSTextView_inst_layoutManager(
		unsafe.Pointer(x.Pointer()),
	)

	return NSLayoutManager_fromPointer(ret)

}

// BackgroundColor The receiver’s background color.
// https://developer.apple.com/documentation/appkit/nstextview/1449501-backgroundcolor?language=objc
func (x gen_NSTextView) BackgroundColor() NSColor {
	ret := C.NSTextView_inst_backgroundColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_fromPointer(ret)

}

// SetBackgroundColor The receiver’s background color.
// https://developer.apple.com/documentation/appkit/nstextview/1449501-backgroundcolor?language=objc
func (x gen_NSTextView) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSTextView_inst_setBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// DrawsBackground A Boolean value that indicates whether the receiver draws its background.
// https://developer.apple.com/documentation/appkit/nstextview/1449530-drawsbackground?language=objc
func (x gen_NSTextView) DrawsBackground() bool {
	ret := C.NSTextView_inst_drawsBackground(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetDrawsBackground A Boolean value that indicates whether the receiver draws its background.
// https://developer.apple.com/documentation/appkit/nstextview/1449530-drawsbackground?language=objc
func (x gen_NSTextView) SetDrawsBackground(
	value bool,
) {
	C.NSTextView_inst_setDrawsBackground(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AllowsDocumentBackgroundColorChange A Boolean value that indicates whether the receiver allows its background color to change.
// https://developer.apple.com/documentation/appkit/nstextview/1449397-allowsdocumentbackgroundcolorcha?language=objc
func (x gen_NSTextView) AllowsDocumentBackgroundColorChange() bool {
	ret := C.NSTextView_inst_allowsDocumentBackgroundColorChange(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsDocumentBackgroundColorChange A Boolean value that indicates whether the receiver allows its background color to change.
// https://developer.apple.com/documentation/appkit/nstextview/1449397-allowsdocumentbackgroundcolorcha?language=objc
func (x gen_NSTextView) SetAllowsDocumentBackgroundColorChange(
	value bool,
) {
	C.NSTextView_inst_setAllowsDocumentBackgroundColorChange(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// ShouldDrawInsertionPoint A Boolean value that determines whether the receiver should draw its insertion point.
// https://developer.apple.com/documentation/appkit/nstextview/1449152-shoulddrawinsertionpoint?language=objc
func (x gen_NSTextView) ShouldDrawInsertionPoint() bool {
	ret := C.NSTextView_inst_shouldDrawInsertionPoint(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AllowedInputSourceLocales An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
// https://developer.apple.com/documentation/appkit/nstextview/1449370-allowedinputsourcelocales?language=objc
func (x gen_NSTextView) AllowedInputSourceLocales() core.NSArray {
	ret := C.NSTextView_inst_allowedInputSourceLocales(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// SetAllowedInputSourceLocales An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
// https://developer.apple.com/documentation/appkit/nstextview/1449370-allowedinputsourcelocales?language=objc
func (x gen_NSTextView) SetAllowedInputSourceLocales(
	value core.NSArrayRef,
) {
	C.NSTextView_inst_setAllowedInputSourceLocales(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// AllowsUndo A Boolean value that indicates whether the receiver allows undo.
// https://developer.apple.com/documentation/appkit/nstextview/1449450-allowsundo?language=objc
func (x gen_NSTextView) AllowsUndo() bool {
	ret := C.NSTextView_inst_allowsUndo(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsUndo A Boolean value that indicates whether the receiver allows undo.
// https://developer.apple.com/documentation/appkit/nstextview/1449450-allowsundo?language=objc
func (x gen_NSTextView) SetAllowsUndo(
	value bool,
) {
	C.NSTextView_inst_setAllowsUndo(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsEditable A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
// https://developer.apple.com/documentation/appkit/nstextview/1449345-editable?language=objc
func (x gen_NSTextView) IsEditable() bool {
	ret := C.NSTextView_inst_isEditable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetEditable A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
// https://developer.apple.com/documentation/appkit/nstextview/1449345-editable?language=objc
func (x gen_NSTextView) SetEditable(
	value bool,
) {
	C.NSTextView_inst_setEditable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsSelectable A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
// https://developer.apple.com/documentation/appkit/nstextview/1449297-selectable?language=objc
func (x gen_NSTextView) IsSelectable() bool {
	ret := C.NSTextView_inst_isSelectable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetSelectable A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
// https://developer.apple.com/documentation/appkit/nstextview/1449297-selectable?language=objc
func (x gen_NSTextView) SetSelectable(
	value bool,
) {
	C.NSTextView_inst_setSelectable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsFieldEditor A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
// https://developer.apple.com/documentation/appkit/nstextview/1449156-fieldeditor?language=objc
func (x gen_NSTextView) IsFieldEditor() bool {
	ret := C.NSTextView_inst_isFieldEditor(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetFieldEditor A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
// https://developer.apple.com/documentation/appkit/nstextview/1449156-fieldeditor?language=objc
func (x gen_NSTextView) SetFieldEditor(
	value bool,
) {
	C.NSTextView_inst_setFieldEditor(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsRichText A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
// https://developer.apple.com/documentation/appkit/nstextview/1449538-richtext?language=objc
func (x gen_NSTextView) IsRichText() bool {
	ret := C.NSTextView_inst_isRichText(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetRichText A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
// https://developer.apple.com/documentation/appkit/nstextview/1449538-richtext?language=objc
func (x gen_NSTextView) SetRichText(
	value bool,
) {
	C.NSTextView_inst_setRichText(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// ImportsGraphics A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to import files by dragging.
// https://developer.apple.com/documentation/appkit/nstextview/1449266-importsgraphics?language=objc
func (x gen_NSTextView) ImportsGraphics() bool {
	ret := C.NSTextView_inst_importsGraphics(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetImportsGraphics A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to import files by dragging.
// https://developer.apple.com/documentation/appkit/nstextview/1449266-importsgraphics?language=objc
func (x gen_NSTextView) SetImportsGraphics(
	value bool,
) {
	C.NSTextView_inst_setImportsGraphics(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AllowsImageEditing Indicates whether image attachments should permit editing of their images.
// https://developer.apple.com/documentation/appkit/nstextview/1449425-allowsimageediting?language=objc
func (x gen_NSTextView) AllowsImageEditing() bool {
	ret := C.NSTextView_inst_allowsImageEditing(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsImageEditing Indicates whether image attachments should permit editing of their images.
// https://developer.apple.com/documentation/appkit/nstextview/1449425-allowsimageediting?language=objc
func (x gen_NSTextView) SetAllowsImageEditing(
	value bool,
) {
	C.NSTextView_inst_setAllowsImageEditing(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsAutomaticQuoteSubstitutionEnabled A Boolean value that enables and disables automatic quotation mark substitution.
// https://developer.apple.com/documentation/appkit/nstextview/1449258-automaticquotesubstitutionenable?language=objc
func (x gen_NSTextView) IsAutomaticQuoteSubstitutionEnabled() bool {
	ret := C.NSTextView_inst_isAutomaticQuoteSubstitutionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAutomaticQuoteSubstitutionEnabled A Boolean value that enables and disables automatic quotation mark substitution.
// https://developer.apple.com/documentation/appkit/nstextview/1449258-automaticquotesubstitutionenable?language=objc
func (x gen_NSTextView) SetAutomaticQuoteSubstitutionEnabled(
	value bool,
) {
	C.NSTextView_inst_setAutomaticQuoteSubstitutionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsAutomaticLinkDetectionEnabled A Boolean value that enables or disables automatic link detection.
// https://developer.apple.com/documentation/appkit/nstextview/1449170-automaticlinkdetectionenabled?language=objc
func (x gen_NSTextView) IsAutomaticLinkDetectionEnabled() bool {
	ret := C.NSTextView_inst_isAutomaticLinkDetectionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAutomaticLinkDetectionEnabled A Boolean value that enables or disables automatic link detection.
// https://developer.apple.com/documentation/appkit/nstextview/1449170-automaticlinkdetectionenabled?language=objc
func (x gen_NSTextView) SetAutomaticLinkDetectionEnabled(
	value bool,
) {
	C.NSTextView_inst_setAutomaticLinkDetectionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// DisplaysLinkToolTips A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
// https://developer.apple.com/documentation/appkit/nstextview/1449204-displayslinktooltips?language=objc
func (x gen_NSTextView) DisplaysLinkToolTips() bool {
	ret := C.NSTextView_inst_displaysLinkToolTips(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetDisplaysLinkToolTips A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
// https://developer.apple.com/documentation/appkit/nstextview/1449204-displayslinktooltips?language=objc
func (x gen_NSTextView) SetDisplaysLinkToolTips(
	value bool,
) {
	C.NSTextView_inst_setDisplaysLinkToolTips(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsAutomaticTextCompletionEnabled A Boolean vlaue that inidcates whether the text view supplies autocompletion suggestions as the user types.
// https://developer.apple.com/documentation/appkit/nstextview/2544655-automatictextcompletionenabled?language=objc
func (x gen_NSTextView) IsAutomaticTextCompletionEnabled() bool {
	ret := C.NSTextView_inst_isAutomaticTextCompletionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAutomaticTextCompletionEnabled A Boolean vlaue that inidcates whether the text view supplies autocompletion suggestions as the user types.
// https://developer.apple.com/documentation/appkit/nstextview/2544655-automatictextcompletionenabled?language=objc
func (x gen_NSTextView) SetAutomaticTextCompletionEnabled(
	value bool,
) {
	C.NSTextView_inst_setAutomaticTextCompletionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// UsesAdaptiveColorMappingForDarkAppearance A Boolean vlaues that indicates whether the framework should use adaptive color mapping for dark appearance.
// https://developer.apple.com/documentation/appkit/nstextview/3237223-usesadaptivecolormappingfordarka?language=objc
func (x gen_NSTextView) UsesAdaptiveColorMappingForDarkAppearance() bool {
	ret := C.NSTextView_inst_usesAdaptiveColorMappingForDarkAppearance(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetUsesAdaptiveColorMappingForDarkAppearance A Boolean vlaues that indicates whether the framework should use adaptive color mapping for dark appearance.
// https://developer.apple.com/documentation/appkit/nstextview/3237223-usesadaptivecolormappingfordarka?language=objc
func (x gen_NSTextView) SetUsesAdaptiveColorMappingForDarkAppearance(
	value bool,
) {
	C.NSTextView_inst_setUsesAdaptiveColorMappingForDarkAppearance(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// UsesRolloverButtonForSelection
// https://developer.apple.com/documentation/appkit/nstextview/1449357-usesrolloverbuttonforselection?language=objc
func (x gen_NSTextView) UsesRolloverButtonForSelection() bool {
	ret := C.NSTextView_inst_usesRolloverButtonForSelection(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetUsesRolloverButtonForSelection
// https://developer.apple.com/documentation/appkit/nstextview/1449357-usesrolloverbuttonforselection?language=objc
func (x gen_NSTextView) SetUsesRolloverButtonForSelection(
	value bool,
) {
	C.NSTextView_inst_setUsesRolloverButtonForSelection(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// UsesRuler A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
// https://developer.apple.com/documentation/appkit/nstextview/1449218-usesruler?language=objc
func (x gen_NSTextView) UsesRuler() bool {
	ret := C.NSTextView_inst_usesRuler(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetUsesRuler A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
// https://developer.apple.com/documentation/appkit/nstextview/1449218-usesruler?language=objc
func (x gen_NSTextView) SetUsesRuler(
	value bool,
) {
	C.NSTextView_inst_setUsesRuler(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsRulerVisible A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
// https://developer.apple.com/documentation/appkit/nstextview/1449406-rulervisible?language=objc
func (x gen_NSTextView) IsRulerVisible() bool {
	ret := C.NSTextView_inst_isRulerVisible(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetRulerVisible A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
// https://developer.apple.com/documentation/appkit/nstextview/1449406-rulervisible?language=objc
func (x gen_NSTextView) SetRulerVisible(
	value bool,
) {
	C.NSTextView_inst_setRulerVisible(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// UsesInspectorBar A Boolean value that indicates whether this text view uses the inspector bar.
// https://developer.apple.com/documentation/appkit/nstextview/1449407-usesinspectorbar?language=objc
func (x gen_NSTextView) UsesInspectorBar() bool {
	ret := C.NSTextView_inst_usesInspectorBar(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetUsesInspectorBar A Boolean value that indicates whether this text view uses the inspector bar.
// https://developer.apple.com/documentation/appkit/nstextview/1449407-usesinspectorbar?language=objc
func (x gen_NSTextView) SetUsesInspectorBar(
	value bool,
) {
	C.NSTextView_inst_setUsesInspectorBar(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// SelectedRanges An array containing the ranges of characters selected in the receiver’s layout manager.
// https://developer.apple.com/documentation/appkit/nstextview/1449129-selectedranges?language=objc
func (x gen_NSTextView) SelectedRanges() core.NSArray {
	ret := C.NSTextView_inst_selectedRanges(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// SetSelectedRanges An array containing the ranges of characters selected in the receiver’s layout manager.
// https://developer.apple.com/documentation/appkit/nstextview/1449129-selectedranges?language=objc
func (x gen_NSTextView) SetSelectedRanges(
	value core.NSArrayRef,
) {
	C.NSTextView_inst_setSelectedRanges(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// InsertionPointColor The color of the insertion point.
// https://developer.apple.com/documentation/appkit/nstextview/1449309-insertionpointcolor?language=objc
func (x gen_NSTextView) InsertionPointColor() NSColor {
	ret := C.NSTextView_inst_insertionPointColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_fromPointer(ret)

}

// SetInsertionPointColor The color of the insertion point.
// https://developer.apple.com/documentation/appkit/nstextview/1449309-insertionpointcolor?language=objc
func (x gen_NSTextView) SetInsertionPointColor(
	value NSColorRef,
) {
	C.NSTextView_inst_setInsertionPointColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// SelectedTextAttributes The attributes used to indicate the selection.
// https://developer.apple.com/documentation/appkit/nstextview/1449270-selectedtextattributes?language=objc
func (x gen_NSTextView) SelectedTextAttributes() core.NSDictionary {
	ret := C.NSTextView_inst_selectedTextAttributes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_fromPointer(ret)

}

// SetSelectedTextAttributes The attributes used to indicate the selection.
// https://developer.apple.com/documentation/appkit/nstextview/1449270-selectedtextattributes?language=objc
func (x gen_NSTextView) SetSelectedTextAttributes(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_setSelectedTextAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// MarkedTextAttributes The attributes used to draw marked text.
// https://developer.apple.com/documentation/appkit/nstextview/1449179-markedtextattributes?language=objc
func (x gen_NSTextView) MarkedTextAttributes() core.NSDictionary {
	ret := C.NSTextView_inst_markedTextAttributes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_fromPointer(ret)

}

// SetMarkedTextAttributes The attributes used to draw marked text.
// https://developer.apple.com/documentation/appkit/nstextview/1449179-markedtextattributes?language=objc
func (x gen_NSTextView) SetMarkedTextAttributes(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_setMarkedTextAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// LinkTextAttributes The attributes used to draw the onscreen presentation of link text.
// https://developer.apple.com/documentation/appkit/nstextview/1449452-linktextattributes?language=objc
func (x gen_NSTextView) LinkTextAttributes() core.NSDictionary {
	ret := C.NSTextView_inst_linkTextAttributes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_fromPointer(ret)

}

// SetLinkTextAttributes The attributes used to draw the onscreen presentation of link text.
// https://developer.apple.com/documentation/appkit/nstextview/1449452-linktextattributes?language=objc
func (x gen_NSTextView) SetLinkTextAttributes(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_setLinkTextAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// ReadablePasteboardTypes The types this text view can read immediately from the pasteboard.
// https://developer.apple.com/documentation/appkit/nstextview/1449361-readablepasteboardtypes?language=objc
func (x gen_NSTextView) ReadablePasteboardTypes() core.NSArray {
	ret := C.NSTextView_inst_readablePasteboardTypes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// WritablePasteboardTypes The pasteboard types that can be provided from the current selection.
// https://developer.apple.com/documentation/appkit/nstextview/1449222-writablepasteboardtypes?language=objc
func (x gen_NSTextView) WritablePasteboardTypes() core.NSArray {
	ret := C.NSTextView_inst_writablePasteboardTypes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// TypingAttributes The receiver’s typing attributes.
// https://developer.apple.com/documentation/appkit/nstextview/1449487-typingattributes?language=objc
func (x gen_NSTextView) TypingAttributes() core.NSDictionary {
	ret := C.NSTextView_inst_typingAttributes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSDictionary_fromPointer(ret)

}

// SetTypingAttributes The receiver’s typing attributes.
// https://developer.apple.com/documentation/appkit/nstextview/1449487-typingattributes?language=objc
func (x gen_NSTextView) SetTypingAttributes(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_setTypingAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// IsCoalescingUndo A Boolean value that indicates whether undo coalescing is in progress.
// https://developer.apple.com/documentation/appkit/nstextview/1449368-coalescingundo?language=objc
func (x gen_NSTextView) IsCoalescingUndo() bool {
	ret := C.NSTextView_inst_isCoalescingUndo(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AcceptableDragTypes The data types that the receiver accepts as the destination view of a dragging operation.
// https://developer.apple.com/documentation/appkit/nstextview/1449234-acceptabledragtypes?language=objc
func (x gen_NSTextView) AcceptableDragTypes() core.NSArray {
	ret := C.NSTextView_inst_acceptableDragTypes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// RangesForUserCharacterAttributeChange An array containing the ranges of characters affected by an action method that changes character (not paragraph) attributes.
// https://developer.apple.com/documentation/appkit/nstextview/1449503-rangesforusercharacterattributec?language=objc
func (x gen_NSTextView) RangesForUserCharacterAttributeChange() core.NSArray {
	ret := C.NSTextView_inst_rangesForUserCharacterAttributeChange(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// RangesForUserParagraphAttributeChange An array containing the ranges of characters affected by a method that changes paragraph (not character) attributes.
// https://developer.apple.com/documentation/appkit/nstextview/1449161-rangesforuserparagraphattributec?language=objc
func (x gen_NSTextView) RangesForUserParagraphAttributeChange() core.NSArray {
	ret := C.NSTextView_inst_rangesForUserParagraphAttributeChange(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// RangesForUserTextChange An array containing the ranges of characters affected by a method that changes characters (as opposed to attributes).
// https://developer.apple.com/documentation/appkit/nstextview/1449434-rangesforusertextchange?language=objc
func (x gen_NSTextView) RangesForUserTextChange() core.NSArray {
	ret := C.NSTextView_inst_rangesForUserTextChange(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// SmartInsertDeleteEnabled A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
// https://developer.apple.com/documentation/appkit/nstextview/1449236-smartinsertdeleteenabled?language=objc
func (x gen_NSTextView) SmartInsertDeleteEnabled() bool {
	ret := C.NSTextView_inst_smartInsertDeleteEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetSmartInsertDeleteEnabled A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
// https://developer.apple.com/documentation/appkit/nstextview/1449236-smartinsertdeleteenabled?language=objc
func (x gen_NSTextView) SetSmartInsertDeleteEnabled(
	value bool,
) {
	C.NSTextView_inst_setSmartInsertDeleteEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsContinuousSpellCheckingEnabled A Boolean value that indicates whether the receiver has continuous spell checking enabled.
// https://developer.apple.com/documentation/appkit/nstextview/1449430-continuousspellcheckingenabled?language=objc
func (x gen_NSTextView) IsContinuousSpellCheckingEnabled() bool {
	ret := C.NSTextView_inst_isContinuousSpellCheckingEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetContinuousSpellCheckingEnabled A Boolean value that indicates whether the receiver has continuous spell checking enabled.
// https://developer.apple.com/documentation/appkit/nstextview/1449430-continuousspellcheckingenabled?language=objc
func (x gen_NSTextView) SetContinuousSpellCheckingEnabled(
	value bool,
) {
	C.NSTextView_inst_setContinuousSpellCheckingEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// SpellCheckerDocumentTag A tag identifying the text view's text as a document for the spell checker server.
// https://developer.apple.com/documentation/appkit/nstextview/1449513-spellcheckerdocumenttag?language=objc
func (x gen_NSTextView) SpellCheckerDocumentTag() core.NSInteger {
	ret := C.NSTextView_inst_spellCheckerDocumentTag(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// IsGrammarCheckingEnabled Enables and disables grammar checking.
// https://developer.apple.com/documentation/appkit/nstextview/1449166-grammarcheckingenabled?language=objc
func (x gen_NSTextView) IsGrammarCheckingEnabled() bool {
	ret := C.NSTextView_inst_isGrammarCheckingEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetGrammarCheckingEnabled Enables and disables grammar checking.
// https://developer.apple.com/documentation/appkit/nstextview/1449166-grammarcheckingenabled?language=objc
func (x gen_NSTextView) SetGrammarCheckingEnabled(
	value bool,
) {
	C.NSTextView_inst_setGrammarCheckingEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AcceptsGlyphInfo A Boolean value that indicates whether the receiver accepts the glyph info attribute.
// https://developer.apple.com/documentation/appkit/nstextview/1449163-acceptsglyphinfo?language=objc
func (x gen_NSTextView) AcceptsGlyphInfo() bool {
	ret := C.NSTextView_inst_acceptsGlyphInfo(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAcceptsGlyphInfo A Boolean value that indicates whether the receiver accepts the glyph info attribute.
// https://developer.apple.com/documentation/appkit/nstextview/1449163-acceptsglyphinfo?language=objc
func (x gen_NSTextView) SetAcceptsGlyphInfo(
	value bool,
) {
	C.NSTextView_inst_setAcceptsGlyphInfo(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// UsesFontPanel A Boolean value that controls whether the text views sharing the receiver’s layout manager use the Font panel and Font menu.
// https://developer.apple.com/documentation/appkit/nstextview/1449534-usesfontpanel?language=objc
func (x gen_NSTextView) UsesFontPanel() bool {
	ret := C.NSTextView_inst_usesFontPanel(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetUsesFontPanel A Boolean value that controls whether the text views sharing the receiver’s layout manager use the Font panel and Font menu.
// https://developer.apple.com/documentation/appkit/nstextview/1449534-usesfontpanel?language=objc
func (x gen_NSTextView) SetUsesFontPanel(
	value bool,
) {
	C.NSTextView_inst_setUsesFontPanel(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// UsesFindPanel A Boolean value that indicates whether the receiver allows for a find panel.
// https://developer.apple.com/documentation/appkit/nstextview/1449293-usesfindpanel?language=objc
func (x gen_NSTextView) UsesFindPanel() bool {
	ret := C.NSTextView_inst_usesFindPanel(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetUsesFindPanel A Boolean value that indicates whether the receiver allows for a find panel.
// https://developer.apple.com/documentation/appkit/nstextview/1449293-usesfindpanel?language=objc
func (x gen_NSTextView) SetUsesFindPanel(
	value bool,
) {
	C.NSTextView_inst_setUsesFindPanel(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsAutomaticDashSubstitutionEnabled A Boolean value that indicates whether automatic dash substitution is enabled.
// https://developer.apple.com/documentation/appkit/nstextview/1449403-automaticdashsubstitutionenabled?language=objc
func (x gen_NSTextView) IsAutomaticDashSubstitutionEnabled() bool {
	ret := C.NSTextView_inst_isAutomaticDashSubstitutionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAutomaticDashSubstitutionEnabled A Boolean value that indicates whether automatic dash substitution is enabled.
// https://developer.apple.com/documentation/appkit/nstextview/1449403-automaticdashsubstitutionenabled?language=objc
func (x gen_NSTextView) SetAutomaticDashSubstitutionEnabled(
	value bool,
) {
	C.NSTextView_inst_setAutomaticDashSubstitutionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsAutomaticDataDetectionEnabled A Boolean value that indicates whether automatic data detection is enabled.
// https://developer.apple.com/documentation/appkit/nstextview/1449192-automaticdatadetectionenabled?language=objc
func (x gen_NSTextView) IsAutomaticDataDetectionEnabled() bool {
	ret := C.NSTextView_inst_isAutomaticDataDetectionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAutomaticDataDetectionEnabled A Boolean value that indicates whether automatic data detection is enabled.
// https://developer.apple.com/documentation/appkit/nstextview/1449192-automaticdatadetectionenabled?language=objc
func (x gen_NSTextView) SetAutomaticDataDetectionEnabled(
	value bool,
) {
	C.NSTextView_inst_setAutomaticDataDetectionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsAutomaticSpellingCorrectionEnabled A Boolean value that indicates whether automatic spelling correction is enabled.
// https://developer.apple.com/documentation/appkit/nstextview/1449254-automaticspellingcorrectionenabl?language=objc
func (x gen_NSTextView) IsAutomaticSpellingCorrectionEnabled() bool {
	ret := C.NSTextView_inst_isAutomaticSpellingCorrectionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAutomaticSpellingCorrectionEnabled A Boolean value that indicates whether automatic spelling correction is enabled.
// https://developer.apple.com/documentation/appkit/nstextview/1449254-automaticspellingcorrectionenabl?language=objc
func (x gen_NSTextView) SetAutomaticSpellingCorrectionEnabled(
	value bool,
) {
	C.NSTextView_inst_setAutomaticSpellingCorrectionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsAutomaticTextReplacementEnabled A Boolean value that indicates whether automatic text replacement is enabled.
// https://developer.apple.com/documentation/appkit/nstextview/1449210-automatictextreplacementenabled?language=objc
func (x gen_NSTextView) IsAutomaticTextReplacementEnabled() bool {
	ret := C.NSTextView_inst_isAutomaticTextReplacementEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAutomaticTextReplacementEnabled A Boolean value that indicates whether automatic text replacement is enabled.
// https://developer.apple.com/documentation/appkit/nstextview/1449210-automatictextreplacementenabled?language=objc
func (x gen_NSTextView) SetAutomaticTextReplacementEnabled(
	value bool,
) {
	C.NSTextView_inst_setAutomaticTextReplacementEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// UsesFindBar A Boolean value that indicates whether to use the find bar for this text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449456-usesfindbar?language=objc
func (x gen_NSTextView) UsesFindBar() bool {
	ret := C.NSTextView_inst_usesFindBar(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetUsesFindBar A Boolean value that indicates whether to use the find bar for this text view.
// https://developer.apple.com/documentation/appkit/nstextview/1449456-usesfindbar?language=objc
func (x gen_NSTextView) SetUsesFindBar(
	value bool,
) {
	C.NSTextView_inst_setUsesFindBar(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsIncrementalSearchingEnabled A Boolean value that indicates whether incremental searching is enabled.
// https://developer.apple.com/documentation/appkit/nstextview/1449458-incrementalsearchingenabled?language=objc
func (x gen_NSTextView) IsIncrementalSearchingEnabled() bool {
	ret := C.NSTextView_inst_isIncrementalSearchingEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetIncrementalSearchingEnabled A Boolean value that indicates whether incremental searching is enabled.
// https://developer.apple.com/documentation/appkit/nstextview/1449458-incrementalsearchingenabled?language=objc
func (x gen_NSTextView) SetIncrementalSearchingEnabled(
	value bool,
) {
	C.NSTextView_inst_setIncrementalSearchingEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AllowsCharacterPickerTouchBarItem
// https://developer.apple.com/documentation/appkit/nstextview/2544680-allowscharacterpickertouchbarite?language=objc
func (x gen_NSTextView) AllowsCharacterPickerTouchBarItem() bool {
	ret := C.NSTextView_inst_allowsCharacterPickerTouchBarItem(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsCharacterPickerTouchBarItem
// https://developer.apple.com/documentation/appkit/nstextview/2544680-allowscharacterpickertouchbarite?language=objc
func (x gen_NSTextView) SetAllowsCharacterPickerTouchBarItem(
	value bool,
) {
	C.NSTextView_inst_setAllowsCharacterPickerTouchBarItem(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// Font
func (x gen_NSTextView) Font() NSFont {
	ret := C.NSTextView_inst_font(
		unsafe.Pointer(x.Pointer()),
	)

	return NSFont_fromPointer(ret)

}

// SetFont
func (x gen_NSTextView) SetFont(
	value NSFontRef,
) {
	C.NSTextView_inst_setFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

type NSViewRef interface {
	Pointer() uintptr
	Init_asNSView() NSView
}

type gen_NSView struct {
	objc.Object
}

func NSView_fromPointer(ptr unsafe.Pointer) NSView {
	return NSView{gen_NSView{
		objc.Object_fromPointer(ptr),
	}}
}

func NSView_fromRef(ref objc.Ref) NSView {
	return NSView_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// AcceptsFirstMouse Overridden by subclasses to return YES if the view should be sent a mouseDown: message for an initial mouse-down event, NO if not.
// https://developer.apple.com/documentation/appkit/nsview/1483410-acceptsfirstmouse?language=objc
func (x gen_NSView) AcceptsFirstMouse(
	event NSEventRef,
) bool {
	ret := C.NSView_inst_acceptsFirstMouse(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return convertObjCBoolToGo(ret)

}

// AddConstraints Adds multiple constraints on the layout of the receiving view or its subviews.
// https://developer.apple.com/documentation/appkit/nsview/1526931-addconstraints?language=objc
func (x gen_NSView) AddConstraints(
	constraints core.NSArrayRef,
) {
	C.NSView_inst_addConstraints(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(constraints),
	)

	return

}

// AddSubview Adds a view to the view’s subviews so it’s displayed above its siblings.
// https://developer.apple.com/documentation/appkit/nsview/1483783-addsubview?language=objc
func (x gen_NSView) AddSubview(
	view NSViewRef,
) {
	C.NSView_inst_addSubview(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)

	return

}

// AddSubview_positioned_relativeTo Inserts a view among the view’s subviews so it’s displayed immediately above or below another view.
// https://developer.apple.com/documentation/appkit/nsview/1483640-addsubview?language=objc
func (x gen_NSView) AddSubview_positioned_relativeTo(
	view NSViewRef,
	place core.NSUInteger,
	otherView NSViewRef,
) {
	C.NSView_inst_addSubview_positioned_relativeTo(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
		C.ulong(place),
		objc.RefPointer(otherView),
	)

	return

}

// AdjustScroll Overridden by subclasses to modify a given rectangle, returning the altered rectangle.
// https://developer.apple.com/documentation/appkit/nsview/1483616-adjustscroll?language=objc
func (x gen_NSView) AdjustScroll(
	newVisible core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_adjustScroll(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&newVisible)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// AlignmentRectForFrame Returns the view’s alignment rectangle for a given frame.
// https://developer.apple.com/documentation/appkit/nsview/1526905-alignmentrectforframe?language=objc
func (x gen_NSView) AlignmentRectForFrame(
	frame core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_alignmentRectForFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frame)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// AncestorSharedWithView Returns the closest ancestor shared by the view and another specified view.
// https://developer.apple.com/documentation/appkit/nsview/1483353-ancestorsharedwithview?language=objc
func (x gen_NSView) AncestorSharedWithView(
	view NSViewRef,
) NSView {
	ret := C.NSView_inst_ancestorSharedWithView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)

	return NSView_fromPointer(ret)

}

// Autoscroll Scrolls the view’s closest ancestor NSClipView object proportionally to the distance of an event that occurs outside of it.
// https://developer.apple.com/documentation/appkit/nsview/1483471-autoscroll?language=objc
func (x gen_NSView) Autoscroll(
	event NSEventRef,
) bool {
	ret := C.NSView_inst_autoscroll(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return convertObjCBoolToGo(ret)

}

// BeginDocument Invoked at the beginning of the printing session, this method sets up the current graphics context.
// https://developer.apple.com/documentation/appkit/nsview/1483423-begindocument?language=objc
func (x gen_NSView) BeginDocument() {
	C.NSView_inst_beginDocument(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// BeginPageInRect_atPlacement Called at the beginning of each page, this method sets up the coordinate system so that a region inside the view’s bounds is translated to a specified location.
// https://developer.apple.com/documentation/appkit/nsview/1483438-beginpageinrect?language=objc
func (x gen_NSView) BeginPageInRect_atPlacement(
	rect core.NSRect,
	location core.NSPoint,
) {
	C.NSView_inst_beginPageInRect_atPlacement(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		*(*C.NSPoint)(unsafe.Pointer(&location)),
	)

	return

}

// CenterScanRect Converts the corners of a specified rectangle to lie on the center of device pixels, which is useful in compensating for rendering overscanning when the coordinate system has been scaled.
// https://developer.apple.com/documentation/appkit/nsview/1483725-centerscanrect?language=objc
func (x gen_NSView) CenterScanRect(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_centerScanRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// ConvertPoint_fromView Converts a point from the coordinate system of a given view to that of the view.
// https://developer.apple.com/documentation/appkit/nsview/1483269-convertpoint?language=objc
func (x gen_NSView) ConvertPoint_fromView(
	point core.NSPoint,
	view NSViewRef,
) core.NSPoint {
	ret := C.NSView_inst_convertPoint_fromView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
		objc.RefPointer(view),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// ConvertPoint_toView Converts a point from the view’s coordinate system to that of a given view.
// https://developer.apple.com/documentation/appkit/nsview/1483406-convertpoint?language=objc
func (x gen_NSView) ConvertPoint_toView(
	point core.NSPoint,
	view NSViewRef,
) core.NSPoint {
	ret := C.NSView_inst_convertPoint_toView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
		objc.RefPointer(view),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// ConvertPointFromBacking Converts a point from its pixel aligned backing store coordinate system to the view’s interior coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483456-convertpointfrombacking?language=objc
func (x gen_NSView) ConvertPointFromBacking(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSView_inst_convertPointFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// ConvertPointFromLayer Convert the point from the layer's interior coordinate system to the view’s interior coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483554-convertpointfromlayer?language=objc
func (x gen_NSView) ConvertPointFromLayer(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSView_inst_convertPointFromLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// ConvertPointToBacking Converts a point from the view’s interior coordinate system to its pixel aligned backing store coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483803-convertpointtobacking?language=objc
func (x gen_NSView) ConvertPointToBacking(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSView_inst_convertPointToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// ConvertPointToLayer Convert the size from the view’s interior coordinate system to the layer's interior coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483315-convertpointtolayer?language=objc
func (x gen_NSView) ConvertPointToLayer(
	point core.NSPoint,
) core.NSPoint {
	ret := C.NSView_inst_convertPointToLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// ConvertRect_fromView Converts a rectangle from the coordinate system of another view to that of the view.
// https://developer.apple.com/documentation/appkit/nsview/1483785-convertrect?language=objc
func (x gen_NSView) ConvertRect_fromView(
	rect core.NSRect,
	view NSViewRef,
) core.NSRect {
	ret := C.NSView_inst_convertRect_fromView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(view),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// ConvertRect_toView Converts a rectangle from the view’s coordinate system to that of another view.
// https://developer.apple.com/documentation/appkit/nsview/1483217-convertrect?language=objc
func (x gen_NSView) ConvertRect_toView(
	rect core.NSRect,
	view NSViewRef,
) core.NSRect {
	ret := C.NSView_inst_convertRect_toView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(view),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// ConvertRectFromBacking Converts a rectangle from its pixel aligned backing store coordinate system to the view’s interior coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483819-convertrectfrombacking?language=objc
func (x gen_NSView) ConvertRectFromBacking(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_convertRectFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// ConvertRectFromLayer Convert the rectangle from the layer's interior coordinate system to the view’s interior coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483404-convertrectfromlayer?language=objc
func (x gen_NSView) ConvertRectFromLayer(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_convertRectFromLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// ConvertRectToBacking Converts a rectangle from the view’s interior coordinate system to its pixel aligned backing store coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483648-convertrecttobacking?language=objc
func (x gen_NSView) ConvertRectToBacking(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_convertRectToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// ConvertRectToLayer Convert the size from the view’s interior coordinate system to the layer's interior coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483776-convertrecttolayer?language=objc
func (x gen_NSView) ConvertRectToLayer(
	rect core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_convertRectToLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// ConvertSize_fromView Converts a size from another view’s coordinate system to that of the view.
// https://developer.apple.com/documentation/appkit/nsview/1483307-convertsize?language=objc
func (x gen_NSView) ConvertSize_fromView(
	size core.NSSize,
	view NSViewRef,
) core.NSSize {
	ret := C.NSView_inst_convertSize_fromView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
		objc.RefPointer(view),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// ConvertSize_toView Converts a size from the view’s coordinate system to that of another view.
// https://developer.apple.com/documentation/appkit/nsview/1483744-convertsize?language=objc
func (x gen_NSView) ConvertSize_toView(
	size core.NSSize,
	view NSViewRef,
) core.NSSize {
	ret := C.NSView_inst_convertSize_toView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
		objc.RefPointer(view),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// ConvertSizeFromBacking Converts a size from its pixel aligned backing store coordinate system to the view’s interior coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483319-convertsizefrombacking?language=objc
func (x gen_NSView) ConvertSizeFromBacking(
	size core.NSSize,
) core.NSSize {
	ret := C.NSView_inst_convertSizeFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// ConvertSizeFromLayer Convert the size from the layer's interior coordinate system to the view’s interior coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483479-convertsizefromlayer?language=objc
func (x gen_NSView) ConvertSizeFromLayer(
	size core.NSSize,
) core.NSSize {
	ret := C.NSView_inst_convertSizeFromLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// ConvertSizeToBacking Converts a size from the view’s interior coordinate system to its pixel aligned backing store coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483227-convertsizetobacking?language=objc
func (x gen_NSView) ConvertSizeToBacking(
	size core.NSSize,
) core.NSSize {
	ret := C.NSView_inst_convertSizeToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// ConvertSizeToLayer Convert the size from the view’s interior coordinate system to the layer's interior coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483701-convertsizetolayer?language=objc
func (x gen_NSView) ConvertSizeToLayer(
	size core.NSSize,
) core.NSSize {
	ret := C.NSView_inst_convertSizeToLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// DataWithEPSInsideRect Returns EPS data that draws the region of the view within a specified rectangle.
// https://developer.apple.com/documentation/appkit/nsview/1483735-datawithepsinsiderect?language=objc
func (x gen_NSView) DataWithEPSInsideRect(
	rect core.NSRect,
) core.NSData {
	ret := C.NSView_inst_dataWithEPSInsideRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return core.NSData_fromPointer(ret)

}

// DataWithPDFInsideRect Returns PDF data that draws the region of the view within a specified rectangle.
// https://developer.apple.com/documentation/appkit/nsview/1483797-datawithpdfinsiderect?language=objc
func (x gen_NSView) DataWithPDFInsideRect(
	rect core.NSRect,
) core.NSData {
	ret := C.NSView_inst_dataWithPDFInsideRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return core.NSData_fromPointer(ret)

}

// DidAddSubview Overridden by subclasses to perform additional actions when subviews are added to the view.
// https://developer.apple.com/documentation/appkit/nsview/1483454-didaddsubview?language=objc
func (x gen_NSView) DidAddSubview(
	subview NSViewRef,
) {
	C.NSView_inst_didAddSubview(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(subview),
	)

	return

}

// DidCloseMenu_withEvent Called after a contextual menu that was displayed from the receiving view has been closed.
// https://developer.apple.com/documentation/appkit/nsview/1483770-didclosemenu?language=objc
func (x gen_NSView) DidCloseMenu_withEvent(
	menu NSMenuRef,
	event NSEventRef,
) {
	C.NSView_inst_didCloseMenu_withEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(menu),
		objc.RefPointer(event),
	)

	return

}

// DiscardCursorRects Invalidates all cursor rectangles set up using addCursorRect:cursor:.
// https://developer.apple.com/documentation/appkit/nsview/1483733-discardcursorrects?language=objc
func (x gen_NSView) DiscardCursorRects() {
	C.NSView_inst_discardCursorRects(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Display Displays the view and all its subviews if possible, invoking each of the NSView methods lockFocus, drawRect:, and unlockFocus as necessary.
// https://developer.apple.com/documentation/appkit/nsview/1483487-display?language=objc
func (x gen_NSView) Display() {
	C.NSView_inst_display(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// DisplayIfNeeded Displays the view and all its subviews if any part of the view has been marked as needing display.
// https://developer.apple.com/documentation/appkit/nsview/1483566-displayifneeded?language=objc
func (x gen_NSView) DisplayIfNeeded() {
	C.NSView_inst_displayIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// DisplayIfNeededIgnoringOpacity Acts as displayIfNeeded, except that this method doesn’t back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
// https://developer.apple.com/documentation/appkit/nsview/1483526-displayifneededignoringopacity?language=objc
func (x gen_NSView) DisplayIfNeededIgnoringOpacity() {
	C.NSView_inst_displayIfNeededIgnoringOpacity(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// DisplayIfNeededInRect Acts as displayIfNeeded, confining drawing to a specified region of the view.
// https://developer.apple.com/documentation/appkit/nsview/1483813-displayifneededinrect?language=objc
func (x gen_NSView) DisplayIfNeededInRect(
	rect core.NSRect,
) {
	C.NSView_inst_displayIfNeededInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return

}

// DisplayIfNeededInRectIgnoringOpacity Acts as displayIfNeeded, but confining drawing to aRect and not backing up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
// https://developer.apple.com/documentation/appkit/nsview/1483481-displayifneededinrectignoringopa?language=objc
func (x gen_NSView) DisplayIfNeededInRectIgnoringOpacity(
	rect core.NSRect,
) {
	C.NSView_inst_displayIfNeededInRectIgnoringOpacity(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return

}

// DisplayRect Acts as display, but confining drawing to a rectangular region of the view.
// https://developer.apple.com/documentation/appkit/nsview/1483518-displayrect?language=objc
func (x gen_NSView) DisplayRect(
	rect core.NSRect,
) {
	C.NSView_inst_displayRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return

}

// DisplayRectIgnoringOpacity Displays the view but confines drawing to a specified region and does not back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
// https://developer.apple.com/documentation/appkit/nsview/1483699-displayrectignoringopacity?language=objc
func (x gen_NSView) DisplayRectIgnoringOpacity(
	rect core.NSRect,
) {
	C.NSView_inst_displayRectIgnoringOpacity(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return

}

// DrawFocusRingMask Draws the focus ring mask for the view.
// https://developer.apple.com/documentation/appkit/nsview/1483335-drawfocusringmask?language=objc
func (x gen_NSView) DrawFocusRingMask() {
	C.NSView_inst_drawFocusRingMask(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// DrawPageBorderWithSize Allows applications that use the AppKit pagination facility to draw additional marks on each logical page.
// https://developer.apple.com/documentation/appkit/nsview/1483292-drawpageborderwithsize?language=objc
func (x gen_NSView) DrawPageBorderWithSize(
	borderSize core.NSSize,
) {
	C.NSView_inst_drawPageBorderWithSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&borderSize)),
	)

	return

}

// DrawRect Overridden by subclasses to draw the view’s image within the specified rectangle.
// https://developer.apple.com/documentation/appkit/nsview/1483686-drawrect?language=objc
func (x gen_NSView) DrawRect(
	dirtyRect core.NSRect,
) {
	C.NSView_inst_drawRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&dirtyRect)),
	)

	return

}

// EndDocument This method is invoked at the end of the printing session.
// https://developer.apple.com/documentation/appkit/nsview/1483610-enddocument?language=objc
func (x gen_NSView) EndDocument() {
	C.NSView_inst_endDocument(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// EndPage Writes the end of a conforming page.
// https://developer.apple.com/documentation/appkit/nsview/1483549-endpage?language=objc
func (x gen_NSView) EndPage() {
	C.NSView_inst_endPage(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// EnterFullScreenMode_withOptions Sets the view to full screen mode.
// https://developer.apple.com/documentation/appkit/nsview/1483780-enterfullscreenmode?language=objc
func (x gen_NSView) EnterFullScreenMode_withOptions(
	screen NSScreenRef,
	options core.NSDictionaryRef,
) bool {
	ret := C.NSView_inst_enterFullScreenMode_withOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(screen),
		objc.RefPointer(options),
	)

	return convertObjCBoolToGo(ret)

}

// ExerciseAmbiguityInLayout Randomly changes the frame of a view with an ambiguous layout between the different valid values.
// https://developer.apple.com/documentation/appkit/nsview/1526934-exerciseambiguityinlayout?language=objc
func (x gen_NSView) ExerciseAmbiguityInLayout() {
	C.NSView_inst_exerciseAmbiguityInLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ExitFullScreenModeWithOptions Instructs the view to exit full screen mode.
// https://developer.apple.com/documentation/appkit/nsview/1483256-exitfullscreenmodewithoptions?language=objc
func (x gen_NSView) ExitFullScreenModeWithOptions(
	options core.NSDictionaryRef,
) {
	C.NSView_inst_exitFullScreenModeWithOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(options),
	)

	return

}

// FrameForAlignmentRect Returns the view’s frame for a given alignment rectangle.
// https://developer.apple.com/documentation/appkit/nsview/1525584-frameforalignmentrect?language=objc
func (x gen_NSView) FrameForAlignmentRect(
	alignmentRect core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_frameForAlignmentRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&alignmentRect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// HitTest Returns the farthest descendant of the view in the view hierarchy (including itself) that contains a specified point, or nil if that point lies completely outside the view.
// https://developer.apple.com/documentation/appkit/nsview/1483364-hittest?language=objc
func (x gen_NSView) HitTest(
	point core.NSPoint,
) NSView {
	ret := C.NSView_inst_hitTest(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return NSView_fromPointer(ret)

}

// InitWithFrame_asNSView Initializes and returns a newly allocated NSView object with a specified frame rectangle.
// https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func (x gen_NSView) InitWithFrame_asNSView(
	frameRect core.NSRect,
) NSView {
	ret := C.NSView_inst_initWithFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)

	return NSView_fromPointer(ret)

}

// InvalidateIntrinsicContentSize Invalidates the view’s intrinsic content size.
// https://developer.apple.com/documentation/appkit/nsview/1526864-invalidateintrinsiccontentsize?language=objc
func (x gen_NSView) InvalidateIntrinsicContentSize() {
	C.NSView_inst_invalidateIntrinsicContentSize(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// IsDescendantOf Returns YES if the view is a subview of a given view or if it’s identical to that view; otherwise, it returns NO.
// https://developer.apple.com/documentation/appkit/nsview/1483219-isdescendantof?language=objc
func (x gen_NSView) IsDescendantOf(
	view NSViewRef,
) bool {
	ret := C.NSView_inst_isDescendantOf(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)

	return convertObjCBoolToGo(ret)

}

// Layout Perform layout in concert with the constraint-based layout system.
// https://developer.apple.com/documentation/appkit/nsview/1526146-layout?language=objc
func (x gen_NSView) Layout() {
	C.NSView_inst_layout(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// LayoutSubtreeIfNeeded Updates the layout of the receiving view and its subviews based on the current views and constraints.
// https://developer.apple.com/documentation/appkit/nsview/1526871-layoutsubtreeifneeded?language=objc
func (x gen_NSView) LayoutSubtreeIfNeeded() {
	C.NSView_inst_layoutSubtreeIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// LocationOfPrintRect Invoked by print: to determine the location of the region of the view being printed on the physical page.
// https://developer.apple.com/documentation/appkit/nsview/1483223-locationofprintrect?language=objc
func (x gen_NSView) LocationOfPrintRect(
	rect core.NSRect,
) core.NSPoint {
	ret := C.NSView_inst_locationOfPrintRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return *(*core.NSPoint)(unsafe.Pointer(&ret))

}

// MakeBackingLayer Creates the view’s backing layer.
// https://developer.apple.com/documentation/appkit/nsview/1483687-makebackinglayer?language=objc
func (x gen_NSView) MakeBackingLayer() core.CALayer {
	ret := C.NSView_inst_makeBackingLayer(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CALayer_fromPointer(ret)

}

// MenuForEvent Overridden by subclasses to return a context-sensitive pop-up menu for a given mouse-down event.
// https://developer.apple.com/documentation/appkit/nsview/1483231-menuforevent?language=objc
func (x gen_NSView) MenuForEvent(
	event NSEventRef,
) NSMenu {
	ret := C.NSView_inst_menuForEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return NSMenu_fromPointer(ret)

}

// Mouse_inRect Returns whether a region of the view contains a specified point, accounting for whether the view is flipped or not.
// https://developer.apple.com/documentation/appkit/nsview/1483237-mouse?language=objc
func (x gen_NSView) Mouse_inRect(
	point core.NSPoint,
	rect core.NSRect,
) bool {
	ret := C.NSView_inst_mouse_inRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return convertObjCBoolToGo(ret)

}

// NeedsToDrawRect Returns a Boolean value indicating whether the specified rectangle intersects any part of the area that the view is being asked to draw.
// https://developer.apple.com/documentation/appkit/nsview/1483570-needstodrawrect?language=objc
func (x gen_NSView) NeedsToDrawRect(
	rect core.NSRect,
) bool {
	ret := C.NSView_inst_needsToDrawRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return convertObjCBoolToGo(ret)

}

// NoteFocusRingMaskChanged Invoked to notify the view that the focus ring mask requires updating.
// https://developer.apple.com/documentation/appkit/nsview/1483809-notefocusringmaskchanged?language=objc
func (x gen_NSView) NoteFocusRingMaskChanged() {
	C.NSView_inst_noteFocusRingMaskChanged(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// PerformKeyEquivalent Implemented by subclasses to respond to key equivalents (also known as keyboard shortcuts).
// https://developer.apple.com/documentation/appkit/nsview/1483664-performkeyequivalent?language=objc
func (x gen_NSView) PerformKeyEquivalent(
	event NSEventRef,
) bool {
	ret := C.NSView_inst_performKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return convertObjCBoolToGo(ret)

}

// PrepareContentInRect Prepares the overdraw region for drawing.
// https://developer.apple.com/documentation/appkit/nsview/1483427-preparecontentinrect?language=objc
func (x gen_NSView) PrepareContentInRect(
	rect core.NSRect,
) {
	C.NSView_inst_prepareContentInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return

}

// PrepareForReuse Restores the view to an initial state so that it can be reused.
// https://developer.apple.com/documentation/appkit/nsview/1483626-prepareforreuse?language=objc
func (x gen_NSView) PrepareForReuse() {
	C.NSView_inst_prepareForReuse(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Print This action method opens the Print panel, and if the user chooses an option other than canceling, prints the view and all its subviews to the device specified in the Print panel.
// https://developer.apple.com/documentation/appkit/nsview/1483705-print?language=objc
func (x gen_NSView) Print(
	sender objc.Ref,
) {
	C.NSView_inst_print(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// RectForPage Implemented by subclasses to determine the portion of the view to be printed for the specified page number.
// https://developer.apple.com/documentation/appkit/nsview/1483252-rectforpage?language=objc
func (x gen_NSView) RectForPage(
	page core.NSInteger,
) core.NSRect {
	ret := C.NSView_inst_rectForPage(
		unsafe.Pointer(x.Pointer()),
		C.long(page),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// RectForSmartMagnificationAtPoint_inRect Returns the appropriate rectangle to use when magnifying around the specified point.
// https://developer.apple.com/documentation/appkit/nsview/1483305-rectforsmartmagnificationatpoint?language=objc
func (x gen_NSView) RectForSmartMagnificationAtPoint_inRect(
	location core.NSPoint,
	visibleRect core.NSRect,
) core.NSRect {
	ret := C.NSView_inst_rectForSmartMagnificationAtPoint_inRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&location)),
		*(*C.NSRect)(unsafe.Pointer(&visibleRect)),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// RegisterForDraggedTypes Registers the pasteboard types that the view will accept as the destination of an image-dragging session.
// https://developer.apple.com/documentation/appkit/nsview/1483578-registerfordraggedtypes?language=objc
func (x gen_NSView) RegisterForDraggedTypes(
	newTypes core.NSArrayRef,
) {
	C.NSView_inst_registerForDraggedTypes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newTypes),
	)

	return

}

// RemoveAllToolTips Removes all tooltips assigned to the view.
// https://developer.apple.com/documentation/appkit/nsview/1483801-removealltooltips?language=objc
func (x gen_NSView) RemoveAllToolTips() {
	C.NSView_inst_removeAllToolTips(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// RemoveConstraints Removes the specified constraints from the view.
// https://developer.apple.com/documentation/appkit/nsview/1526932-removeconstraints?language=objc
func (x gen_NSView) RemoveConstraints(
	constraints core.NSArrayRef,
) {
	C.NSView_inst_removeConstraints(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(constraints),
	)

	return

}

// RemoveFromSuperview Unlinks the view from its superview and its window, removes it from the responder chain, and invalidates its cursor rectangles.
// https://developer.apple.com/documentation/appkit/nsview/1483265-removefromsuperview?language=objc
func (x gen_NSView) RemoveFromSuperview() {
	C.NSView_inst_removeFromSuperview(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// RemoveFromSuperviewWithoutNeedingDisplay Unlinks the view from its superview and its window and removes it from the responder chain, but does not invalidate its cursor rectangles to cause redrawing.
// https://developer.apple.com/documentation/appkit/nsview/1483644-removefromsuperviewwithoutneedin?language=objc
func (x gen_NSView) RemoveFromSuperviewWithoutNeedingDisplay() {
	C.NSView_inst_removeFromSuperviewWithoutNeedingDisplay(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ReplaceSubview_with Replaces one of the view’s subviews with another view.
// https://developer.apple.com/documentation/appkit/nsview/1483632-replacesubview?language=objc
func (x gen_NSView) ReplaceSubview_with(
	oldView NSViewRef,
	newView NSViewRef,
) {
	C.NSView_inst_replaceSubview_with(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(oldView),
		objc.RefPointer(newView),
	)

	return

}

// ResetCursorRects Overridden by subclasses to define their default cursor rectangles.
// https://developer.apple.com/documentation/appkit/nsview/1483448-resetcursorrects?language=objc
func (x gen_NSView) ResetCursorRects() {
	C.NSView_inst_resetCursorRects(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ResizeSubviewsWithOldSize Informs the view’s subviews that the view’s bounds rectangle size has changed.
// https://developer.apple.com/documentation/appkit/nsview/1483495-resizesubviewswitholdsize?language=objc
func (x gen_NSView) ResizeSubviewsWithOldSize(
	oldSize core.NSSize,
) {
	C.NSView_inst_resizeSubviewsWithOldSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&oldSize)),
	)

	return

}

// ResizeWithOldSuperviewSize Informs the view that the bounds size of its superview has changed.
// https://developer.apple.com/documentation/appkit/nsview/1483697-resizewitholdsuperviewsize?language=objc
func (x gen_NSView) ResizeWithOldSuperviewSize(
	oldSize core.NSSize,
) {
	C.NSView_inst_resizeWithOldSuperviewSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&oldSize)),
	)

	return

}

// RotateByAngle Rotates the view’s bounds rectangle by a specified degree value around the origin of the coordinate system, (0.0, 0.0).
// https://developer.apple.com/documentation/appkit/nsview/1483444-rotatebyangle?language=objc
func (x gen_NSView) RotateByAngle(
	angle core.CGFloat,
) {
	C.NSView_inst_rotateByAngle(
		unsafe.Pointer(x.Pointer()),
		C.double(angle),
	)

	return

}

// ScaleUnitSquareToSize Scales the view’s coordinate system so that the unit square scales to the specified dimensions.
// https://developer.apple.com/documentation/appkit/nsview/1483721-scaleunitsquaretosize?language=objc
func (x gen_NSView) ScaleUnitSquareToSize(
	newUnitSize core.NSSize,
) {
	C.NSView_inst_scaleUnitSquareToSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&newUnitSize)),
	)

	return

}

// ScrollPoint Scrolls the view’s closest ancestor NSClipView object so a point in the view lies at the origin of the clip view's bounds rectangle.
// https://developer.apple.com/documentation/appkit/nsview/1483394-scrollpoint?language=objc
func (x gen_NSView) ScrollPoint(
	point core.NSPoint,
) {
	C.NSView_inst_scrollPoint(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)

	return

}

// ScrollRectToVisible Scrolls the view’s closest ancestor NSClipView object the minimum distance needed so a specified region of the view becomes visible in the clip view.
// https://developer.apple.com/documentation/appkit/nsview/1483811-scrollrecttovisible?language=objc
func (x gen_NSView) ScrollRectToVisible(
	rect core.NSRect,
) bool {
	ret := C.NSView_inst_scrollRectToVisible(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return convertObjCBoolToGo(ret)

}

// SetBoundsOrigin Sets the origin of the view’s bounds rectangle to a specified point.
// https://developer.apple.com/documentation/appkit/nsview/1483345-setboundsorigin?language=objc
func (x gen_NSView) SetBoundsOrigin(
	newOrigin core.NSPoint,
) {
	C.NSView_inst_setBoundsOrigin(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&newOrigin)),
	)

	return

}

// SetBoundsSize Sets the size of the view’s bounds rectangle to specified dimensions, inversely scaling its coordinate system relative to its frame rectangle.
// https://developer.apple.com/documentation/appkit/nsview/1483399-setboundssize?language=objc
func (x gen_NSView) SetBoundsSize(
	newSize core.NSSize,
) {
	C.NSView_inst_setBoundsSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&newSize)),
	)

	return

}

// SetFrameOrigin Sets the origin of the view’s frame rectangle to the specified point, effectively repositioning it within its superview.
// https://developer.apple.com/documentation/appkit/nsview/1483283-setframeorigin?language=objc
func (x gen_NSView) SetFrameOrigin(
	newOrigin core.NSPoint,
) {
	C.NSView_inst_setFrameOrigin(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&newOrigin)),
	)

	return

}

// SetFrameSize Sets the size of the view’s frame rectangle to the specified dimensions, resizing it within its superview without affecting its coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483530-setframesize?language=objc
func (x gen_NSView) SetFrameSize(
	newSize core.NSSize,
) {
	C.NSView_inst_setFrameSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&newSize)),
	)

	return

}

// SetKeyboardFocusRingNeedsDisplayInRect Invalidates the area around the focus ring.
// https://developer.apple.com/documentation/appkit/nsview/1483556-setkeyboardfocusringneedsdisplay?language=objc
func (x gen_NSView) SetKeyboardFocusRingNeedsDisplayInRect(
	rect core.NSRect,
) {
	C.NSView_inst_setKeyboardFocusRingNeedsDisplayInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return

}

// SetNeedsDisplayInRect Marks the region of the view within the specified rectangle as needing display, increasing the view’s existing invalid region to include it.
// https://developer.apple.com/documentation/appkit/nsview/1483475-setneedsdisplayinrect?language=objc
func (x gen_NSView) SetNeedsDisplayInRect(
	invalidRect core.NSRect,
) {
	C.NSView_inst_setNeedsDisplayInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&invalidRect)),
	)

	return

}

// ShouldDelayWindowOrderingForEvent Allows the user to drag objects from the view without activating the app or moving the window of the view forward, possibly obscuring the destination.
// https://developer.apple.com/documentation/appkit/nsview/1483244-shoulddelaywindoworderingforeven?language=objc
func (x gen_NSView) ShouldDelayWindowOrderingForEvent(
	event NSEventRef,
) bool {
	ret := C.NSView_inst_shouldDelayWindowOrderingForEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return convertObjCBoolToGo(ret)

}

// ShowDefinitionForAttributedString_atPoint Shows a window displaying the definition of the attributed string at the specified point.
// https://developer.apple.com/documentation/appkit/nsview/1483747-showdefinitionforattributedstrin?language=objc
func (x gen_NSView) ShowDefinitionForAttributedString_atPoint(
	attrString core.NSAttributedStringRef,
	textBaselineOrigin core.NSPoint,
) {
	C.NSView_inst_showDefinitionForAttributedString_atPoint(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(attrString),
		*(*C.NSPoint)(unsafe.Pointer(&textBaselineOrigin)),
	)

	return

}

// TranslateOriginToPoint Translates the view’s coordinate system so that its origin moves to a new location.
// https://developer.apple.com/documentation/appkit/nsview/1483385-translateorigintopoint?language=objc
func (x gen_NSView) TranslateOriginToPoint(
	translation core.NSPoint,
) {
	C.NSView_inst_translateOriginToPoint(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&translation)),
	)

	return

}

// TranslateRectsNeedingDisplayInRect_by Translates the display rectangles by the specified delta.
// https://developer.apple.com/documentation/appkit/nsview/1483731-translaterectsneedingdisplayinre?language=objc
func (x gen_NSView) TranslateRectsNeedingDisplayInRect_by(
	clipRect core.NSRect,
	delta core.NSSize,
) {
	C.NSView_inst_translateRectsNeedingDisplayInRect_by(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&clipRect)),
		*(*C.NSSize)(unsafe.Pointer(&delta)),
	)

	return

}

// UnregisterDraggedTypes Unregisters the view as a possible destination in a dragging session.
// https://developer.apple.com/documentation/appkit/nsview/1483602-unregisterdraggedtypes?language=objc
func (x gen_NSView) UnregisterDraggedTypes() {
	C.NSView_inst_unregisterDraggedTypes(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// UpdateConstraints Update constraints for the view.
// https://developer.apple.com/documentation/appkit/nsview/1526891-updateconstraints?language=objc
func (x gen_NSView) UpdateConstraints() {
	C.NSView_inst_updateConstraints(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// UpdateConstraintsForSubtreeIfNeeded Updates the constraints for the receiving view and its subviews.
// https://developer.apple.com/documentation/appkit/nsview/1526939-updateconstraintsforsubtreeifnee?language=objc
func (x gen_NSView) UpdateConstraintsForSubtreeIfNeeded() {
	C.NSView_inst_updateConstraintsForSubtreeIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// UpdateLayer Updates the view’s content by modifying its underlying layer.
// https://developer.apple.com/documentation/appkit/nsview/1483580-updatelayer?language=objc
func (x gen_NSView) UpdateLayer() {
	C.NSView_inst_updateLayer(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// UpdateTrackingAreas Invoked automatically when the view’s geometry changes such that its tracking areas need to be recalculated.
// https://developer.apple.com/documentation/appkit/nsview/1483719-updatetrackingareas?language=objc
func (x gen_NSView) UpdateTrackingAreas() {
	C.NSView_inst_updateTrackingAreas(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewDidChangeBackingProperties Responds when the view’s backing store properties change.
// https://developer.apple.com/documentation/appkit/nsview/1483742-viewdidchangebackingproperties?language=objc
func (x gen_NSView) ViewDidChangeBackingProperties() {
	C.NSView_inst_viewDidChangeBackingProperties(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewDidChangeEffectiveAppearance
// https://developer.apple.com/documentation/appkit/nsview/2977088-viewdidchangeeffectiveappearance?language=objc
func (x gen_NSView) ViewDidChangeEffectiveAppearance() {
	C.NSView_inst_viewDidChangeEffectiveAppearance(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewDidEndLiveResize Informs the view of the end of a live resize—the user has finished resizing the view.
// https://developer.apple.com/documentation/appkit/nsview/1483543-viewdidendliveresize?language=objc
func (x gen_NSView) ViewDidEndLiveResize() {
	C.NSView_inst_viewDidEndLiveResize(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewDidHide Invoked when the view is hidden, either directly, or in response to an ancestor being hidden.
// https://developer.apple.com/documentation/appkit/nsview/1483596-viewdidhide?language=objc
func (x gen_NSView) ViewDidHide() {
	C.NSView_inst_viewDidHide(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewDidMoveToSuperview Informs the view that its superview has changed (possibly to nil).
// https://developer.apple.com/documentation/appkit/nsview/1483568-viewdidmovetosuperview?language=objc
func (x gen_NSView) ViewDidMoveToSuperview() {
	C.NSView_inst_viewDidMoveToSuperview(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewDidMoveToWindow Informs the view that it has been added to a new view hierarchy.
// https://developer.apple.com/documentation/appkit/nsview/1483329-viewdidmovetowindow?language=objc
func (x gen_NSView) ViewDidMoveToWindow() {
	C.NSView_inst_viewDidMoveToWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewDidUnhide Invoked when the view is unhidden, either directly, or in response to an ancestor being unhidden
// https://developer.apple.com/documentation/appkit/nsview/1483275-viewdidunhide?language=objc
func (x gen_NSView) ViewDidUnhide() {
	C.NSView_inst_viewDidUnhide(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewWillDraw Informs the view that it’s required to draw content.
// https://developer.apple.com/documentation/appkit/nsview/1483351-viewwilldraw?language=objc
func (x gen_NSView) ViewWillDraw() {
	C.NSView_inst_viewWillDraw(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewWillMoveToSuperview Informs the view that its superview is about to change to the specified superview (which may be nil).
// https://developer.apple.com/documentation/appkit/nsview/1483545-viewwillmovetosuperview?language=objc
func (x gen_NSView) ViewWillMoveToSuperview(
	newSuperview NSViewRef,
) {
	C.NSView_inst_viewWillMoveToSuperview(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newSuperview),
	)

	return

}

// ViewWillMoveToWindow Informs the view that it’s being added to the view hierarchy of the specified window object (which may be nil).
// https://developer.apple.com/documentation/appkit/nsview/1483415-viewwillmovetowindow?language=objc
func (x gen_NSView) ViewWillMoveToWindow(
	newWindow NSWindowRef,
) {
	C.NSView_inst_viewWillMoveToWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newWindow),
	)

	return

}

// ViewWillStartLiveResize Informs the view of the start of a live resize—the user has started resizing the view.
// https://developer.apple.com/documentation/appkit/nsview/1483620-viewwillstartliveresize?language=objc
func (x gen_NSView) ViewWillStartLiveResize() {
	C.NSView_inst_viewWillStartLiveResize(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ViewWithTag Returns the view’s nearest descendant (including itself) with a specific tag, or nil if no subview has that tag.
// https://developer.apple.com/documentation/appkit/nsview/1483294-viewwithtag?language=objc
func (x gen_NSView) ViewWithTag(
	tag core.NSInteger,
) NSView {
	ret := C.NSView_inst_viewWithTag(
		unsafe.Pointer(x.Pointer()),
		C.long(tag),
	)

	return NSView_fromPointer(ret)

}

// WillOpenMenu_withEvent Called just before a contextual menu for a view is opened on screen.
// https://developer.apple.com/documentation/appkit/nsview/1483429-willopenmenu?language=objc
func (x gen_NSView) WillOpenMenu_withEvent(
	menu NSMenuRef,
	event NSEventRef,
) {
	C.NSView_inst_willOpenMenu_withEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(menu),
		objc.RefPointer(event),
	)

	return

}

// WillRemoveSubview Overridden by subclasses to perform additional actions before subviews are removed from the view.
// https://developer.apple.com/documentation/appkit/nsview/1483624-willremovesubview?language=objc
func (x gen_NSView) WillRemoveSubview(
	subview NSViewRef,
) {
	C.NSView_inst_willRemoveSubview(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(subview),
	)

	return

}

// WriteEPSInsideRect_toPasteboard Writes EPS data that draws the region of the view within a specified rectangle onto a pasteboard.
// https://developer.apple.com/documentation/appkit/nsview/1483235-writeepsinsiderect?language=objc
func (x gen_NSView) WriteEPSInsideRect_toPasteboard(
	rect core.NSRect,
	pasteboard NSPasteboardRef,
) {
	C.NSView_inst_writeEPSInsideRect_toPasteboard(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(pasteboard),
	)

	return

}

// WritePDFInsideRect_toPasteboard Writes PDF data that draws the region of the view within a specified rectangle onto a pasteboard.
// https://developer.apple.com/documentation/appkit/nsview/1483499-writepdfinsiderect?language=objc
func (x gen_NSView) WritePDFInsideRect_toPasteboard(
	rect core.NSRect,
	pasteboard NSPasteboardRef,
) {
	C.NSView_inst_writePDFInsideRect_toPasteboard(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(pasteboard),
	)

	return

}

// Init_asNSView
func (x gen_NSView) Init_asNSView() NSView {
	ret := C.NSView_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_fromPointer(ret)

}

// Superview The view that is the parent of the current view.
// https://developer.apple.com/documentation/appkit/nsview/1483737-superview?language=objc
func (x gen_NSView) Superview() NSView {
	ret := C.NSView_inst_superview(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_fromPointer(ret)

}

// Subviews The array of views embedded in the current view.
// https://developer.apple.com/documentation/appkit/nsview/1483539-subviews?language=objc
func (x gen_NSView) Subviews() core.NSArray {
	ret := C.NSView_inst_subviews(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// SetSubviews The array of views embedded in the current view.
// https://developer.apple.com/documentation/appkit/nsview/1483539-subviews?language=objc
func (x gen_NSView) SetSubviews(
	value core.NSArrayRef,
) {
	C.NSView_inst_setSubviews(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Window The view’s window object, if it is installed in a window.
// https://developer.apple.com/documentation/appkit/nsview/1483301-window?language=objc
func (x gen_NSView) Window() NSWindow {
	ret := C.NSView_inst_window(
		unsafe.Pointer(x.Pointer()),
	)

	return NSWindow_fromPointer(ret)

}

// OpaqueAncestor The view’s closest opaque ancestor, which might be the view itself.
// https://developer.apple.com/documentation/appkit/nsview/1483383-opaqueancestor?language=objc
func (x gen_NSView) OpaqueAncestor() NSView {
	ret := C.NSView_inst_opaqueAncestor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_fromPointer(ret)

}

// EnclosingMenuItem The menu item containing the view or any of its superviews in the view hierarchy.
// https://developer.apple.com/documentation/appkit/nsview/1514865-enclosingmenuitem?language=objc
func (x gen_NSView) EnclosingMenuItem() NSMenuItem {
	ret := C.NSView_inst_enclosingMenuItem(
		unsafe.Pointer(x.Pointer()),
	)

	return NSMenuItem_fromPointer(ret)

}

// Frame The view’s frame rectangle, which defines its position and size in its superview’s coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483713-frame?language=objc
func (x gen_NSView) Frame() core.NSRect {
	ret := C.NSView_inst_frame(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// SetFrame The view’s frame rectangle, which defines its position and size in its superview’s coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483713-frame?language=objc
func (x gen_NSView) SetFrame(
	value core.NSRect,
) {
	C.NSView_inst_setFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return

}

// FrameRotation The angle of rotation, measured in degrees, applied to the view’s frame rectangle relative to its superview’s coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483412-framerotation?language=objc
func (x gen_NSView) FrameRotation() core.CGFloat {
	ret := C.NSView_inst_frameRotation(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// SetFrameRotation The angle of rotation, measured in degrees, applied to the view’s frame rectangle relative to its superview’s coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483412-framerotation?language=objc
func (x gen_NSView) SetFrameRotation(
	value core.CGFloat,
) {
	C.NSView_inst_setFrameRotation(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// Bounds The view’s bounds rectangle, which expresses its location and size in its own coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483817-bounds?language=objc
func (x gen_NSView) Bounds() core.NSRect {
	ret := C.NSView_inst_bounds(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// SetBounds The view’s bounds rectangle, which expresses its location and size in its own coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483817-bounds?language=objc
func (x gen_NSView) SetBounds(
	value core.NSRect,
) {
	C.NSView_inst_setBounds(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return

}

// BoundsRotation The angle of rotation, measured in degrees, applied to the view’s bounds rectangle relative to its frame rectangle.
// https://developer.apple.com/documentation/appkit/nsview/1483746-boundsrotation?language=objc
func (x gen_NSView) BoundsRotation() core.CGFloat {
	ret := C.NSView_inst_boundsRotation(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// SetBoundsRotation The angle of rotation, measured in degrees, applied to the view’s bounds rectangle relative to its frame rectangle.
// https://developer.apple.com/documentation/appkit/nsview/1483746-boundsrotation?language=objc
func (x gen_NSView) SetBoundsRotation(
	value core.CGFloat,
) {
	C.NSView_inst_setBoundsRotation(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// WantsLayer A Boolean value indicating whether the view uses a layer as its backing store.
// https://developer.apple.com/documentation/appkit/nsview/1483695-wantslayer?language=objc
func (x gen_NSView) WantsLayer() bool {
	ret := C.NSView_inst_wantsLayer(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetWantsLayer A Boolean value indicating whether the view uses a layer as its backing store.
// https://developer.apple.com/documentation/appkit/nsview/1483695-wantslayer?language=objc
func (x gen_NSView) SetWantsLayer(
	value bool,
) {
	C.NSView_inst_setWantsLayer(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// WantsUpdateLayer A Boolean value indicating which drawing path the view takes when updating its contents.
// https://developer.apple.com/documentation/appkit/nsview/1483461-wantsupdatelayer?language=objc
func (x gen_NSView) WantsUpdateLayer() bool {
	ret := C.NSView_inst_wantsUpdateLayer(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Layer The Core Animation layer that the view uses as its backing store.
// https://developer.apple.com/documentation/appkit/nsview/1483298-layer?language=objc
func (x gen_NSView) Layer() core.CALayer {
	ret := C.NSView_inst_layer(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CALayer_fromPointer(ret)

}

// SetLayer The Core Animation layer that the view uses as its backing store.
// https://developer.apple.com/documentation/appkit/nsview/1483298-layer?language=objc
func (x gen_NSView) SetLayer(
	value core.CALayerRef,
) {
	C.NSView_inst_setLayer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// CanDrawSubviewsIntoLayer A Boolean value indicating whether the view incorporates content from its subviews into its own layer.
// https://developer.apple.com/documentation/appkit/nsview/1483347-candrawsubviewsintolayer?language=objc
func (x gen_NSView) CanDrawSubviewsIntoLayer() bool {
	ret := C.NSView_inst_canDrawSubviewsIntoLayer(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetCanDrawSubviewsIntoLayer A Boolean value indicating whether the view incorporates content from its subviews into its own layer.
// https://developer.apple.com/documentation/appkit/nsview/1483347-candrawsubviewsintolayer?language=objc
func (x gen_NSView) SetCanDrawSubviewsIntoLayer(
	value bool,
) {
	C.NSView_inst_setCanDrawSubviewsIntoLayer(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// LayerUsesCoreImageFilters A Boolean value indicating whether the view’s layer uses Core Image filters and needs in-process rendering.
// https://developer.apple.com/documentation/appkit/nsview/1483576-layerusescoreimagefilters?language=objc
func (x gen_NSView) LayerUsesCoreImageFilters() bool {
	ret := C.NSView_inst_layerUsesCoreImageFilters(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetLayerUsesCoreImageFilters A Boolean value indicating whether the view’s layer uses Core Image filters and needs in-process rendering.
// https://developer.apple.com/documentation/appkit/nsview/1483576-layerusescoreimagefilters?language=objc
func (x gen_NSView) SetLayerUsesCoreImageFilters(
	value bool,
) {
	C.NSView_inst_setLayerUsesCoreImageFilters(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AlphaValue The opacity of the view.
// https://developer.apple.com/documentation/appkit/nsview/1483560-alphavalue?language=objc
func (x gen_NSView) AlphaValue() core.CGFloat {
	ret := C.NSView_inst_alphaValue(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// SetAlphaValue The opacity of the view.
// https://developer.apple.com/documentation/appkit/nsview/1483560-alphavalue?language=objc
func (x gen_NSView) SetAlphaValue(
	value core.CGFloat,
) {
	C.NSView_inst_setAlphaValue(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// FrameCenterRotation The rotation angle of the view around the center of its layer.
// https://developer.apple.com/documentation/appkit/nsview/1483367-framecenterrotation?language=objc
func (x gen_NSView) FrameCenterRotation() core.CGFloat {
	ret := C.NSView_inst_frameCenterRotation(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// SetFrameCenterRotation The rotation angle of the view around the center of its layer.
// https://developer.apple.com/documentation/appkit/nsview/1483367-framecenterrotation?language=objc
func (x gen_NSView) SetFrameCenterRotation(
	value core.CGFloat,
) {
	C.NSView_inst_setFrameCenterRotation(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// BackgroundFilters An array of Core Image filters to apply to the view’s background.
// https://developer.apple.com/documentation/appkit/nsview/1483689-backgroundfilters?language=objc
func (x gen_NSView) BackgroundFilters() core.NSArray {
	ret := C.NSView_inst_backgroundFilters(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// SetBackgroundFilters An array of Core Image filters to apply to the view’s background.
// https://developer.apple.com/documentation/appkit/nsview/1483689-backgroundfilters?language=objc
func (x gen_NSView) SetBackgroundFilters(
	value core.NSArrayRef,
) {
	C.NSView_inst_setBackgroundFilters(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// ContentFilters An array of Core Image filters to apply to the contents of the view and its sublayers.
// https://developer.apple.com/documentation/appkit/nsview/1483703-contentfilters?language=objc
func (x gen_NSView) ContentFilters() core.NSArray {
	ret := C.NSView_inst_contentFilters(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// SetContentFilters An array of Core Image filters to apply to the contents of the view and its sublayers.
// https://developer.apple.com/documentation/appkit/nsview/1483703-contentfilters?language=objc
func (x gen_NSView) SetContentFilters(
	value core.NSArrayRef,
) {
	C.NSView_inst_setContentFilters(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// CanDrawConcurrently A Boolean value indicating whether the view can draw its contents on a background thread.
// https://developer.apple.com/documentation/appkit/nsview/1483425-candrawconcurrently?language=objc
func (x gen_NSView) CanDrawConcurrently() bool {
	ret := C.NSView_inst_canDrawConcurrently(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetCanDrawConcurrently A Boolean value indicating whether the view can draw its contents on a background thread.
// https://developer.apple.com/documentation/appkit/nsview/1483425-candrawconcurrently?language=objc
func (x gen_NSView) SetCanDrawConcurrently(
	value bool,
) {
	C.NSView_inst_setCanDrawConcurrently(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// VisibleRect The portion of the view that is not clipped by its superviews.
// https://developer.apple.com/documentation/appkit/nsview/1483446-visiblerect?language=objc
func (x gen_NSView) VisibleRect() core.NSRect {
	ret := C.NSView_inst_visibleRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// WantsDefaultClipping A Boolean value indicating whether AppKit’s default clipping behavior is in effect.
// https://developer.apple.com/documentation/appkit/nsview/1483365-wantsdefaultclipping?language=objc
func (x gen_NSView) WantsDefaultClipping() bool {
	ret := C.NSView_inst_wantsDefaultClipping(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// PrintJobTitle The view’s print job title.
// https://developer.apple.com/documentation/appkit/nsview/1483753-printjobtitle?language=objc
func (x gen_NSView) PrintJobTitle() core.NSString {
	ret := C.NSView_inst_printJobTitle(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// PageHeader A default header string that includes the print job title and date.
// https://developer.apple.com/documentation/appkit/nsview/1483674-pageheader?language=objc
func (x gen_NSView) PageHeader() core.NSAttributedString {
	ret := C.NSView_inst_pageHeader(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSAttributedString_fromPointer(ret)

}

// PageFooter A default footer string that includes the current page number and page count.
// https://developer.apple.com/documentation/appkit/nsview/1483355-pagefooter?language=objc
func (x gen_NSView) PageFooter() core.NSAttributedString {
	ret := C.NSView_inst_pageFooter(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSAttributedString_fromPointer(ret)

}

// HeightAdjustLimit The fraction of the page that can be pushed onto the next page during automatic pagination to prevent items such as lines of text from being divided across pages.
// https://developer.apple.com/documentation/appkit/nsview/1483691-heightadjustlimit?language=objc
func (x gen_NSView) HeightAdjustLimit() core.CGFloat {
	ret := C.NSView_inst_heightAdjustLimit(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// WidthAdjustLimit The fraction of the page that can be pushed onto the next page during automatic pagination to prevent items such as small images or text columns from being divided across pages.
// https://developer.apple.com/documentation/appkit/nsview/1483392-widthadjustlimit?language=objc
func (x gen_NSView) WidthAdjustLimit() core.CGFloat {
	ret := C.NSView_inst_widthAdjustLimit(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// NeedsDisplay A Boolean value that determines whether the view needs to be redrawn before being displayed.
// https://developer.apple.com/documentation/appkit/nsview/1483360-needsdisplay?language=objc
func (x gen_NSView) NeedsDisplay() bool {
	ret := C.NSView_inst_needsDisplay(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetNeedsDisplay A Boolean value that determines whether the view needs to be redrawn before being displayed.
// https://developer.apple.com/documentation/appkit/nsview/1483360-needsdisplay?language=objc
func (x gen_NSView) SetNeedsDisplay(
	value bool,
) {
	C.NSView_inst_setNeedsDisplay(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsOpaque A Boolean value indicating whether the view fills its frame rectangle with opaque content.
// https://developer.apple.com/documentation/appkit/nsview/1483558-opaque?language=objc
func (x gen_NSView) IsOpaque() bool {
	ret := C.NSView_inst_isOpaque(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsFlipped A Boolean value indicating whether the view uses a flipped coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483532-flipped?language=objc
func (x gen_NSView) IsFlipped() bool {
	ret := C.NSView_inst_isFlipped(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsRotatedFromBase A Boolean value indicating whether the view or any of its ancestors has ever had a rotation factor applied to its frame or bounds.
// https://developer.apple.com/documentation/appkit/nsview/1483709-rotatedfrombase?language=objc
func (x gen_NSView) IsRotatedFromBase() bool {
	ret := C.NSView_inst_isRotatedFromBase(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsRotatedOrScaledFromBase A Boolean value indicating whether the view or any of its ancestors has ever had a rotation factor applied to its frame or bounds, or has been scaled from the window’s base coordinate system.
// https://developer.apple.com/documentation/appkit/nsview/1483390-rotatedorscaledfrombase?language=objc
func (x gen_NSView) IsRotatedOrScaledFromBase() bool {
	ret := C.NSView_inst_isRotatedOrScaledFromBase(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AutoresizesSubviews A Boolean value indicating whether the view applies the autoresizing behavior to its subviews when its frame size changes.
// https://developer.apple.com/documentation/appkit/nsview/1483358-autoresizessubviews?language=objc
func (x gen_NSView) AutoresizesSubviews() bool {
	ret := C.NSView_inst_autoresizesSubviews(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAutoresizesSubviews A Boolean value indicating whether the view applies the autoresizing behavior to its subviews when its frame size changes.
// https://developer.apple.com/documentation/appkit/nsview/1483358-autoresizessubviews?language=objc
func (x gen_NSView) SetAutoresizesSubviews(
	value bool,
) {
	C.NSView_inst_setAutoresizesSubviews(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// Constraints Returns the constraints held by the view.
// https://developer.apple.com/documentation/appkit/nsview/1526917-constraints?language=objc
func (x gen_NSView) Constraints() core.NSArray {
	ret := C.NSView_inst_constraints(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// LayoutGuides The array of layout guide objects owned by this view.
// https://developer.apple.com/documentation/appkit/nsview/1534395-layoutguides?language=objc
func (x gen_NSView) LayoutGuides() core.NSArray {
	ret := C.NSView_inst_layoutGuides(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// FittingSize The minimum size of the view that satisfies the constraints it holds.
// https://developer.apple.com/documentation/appkit/nsview/1526904-fittingsize?language=objc
func (x gen_NSView) FittingSize() core.NSSize {
	ret := C.NSView_inst_fittingSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// IntrinsicContentSize The natural size for the receiving view, considering only properties of the view itself.
// https://developer.apple.com/documentation/appkit/nsview/1526996-intrinsiccontentsize?language=objc
func (x gen_NSView) IntrinsicContentSize() core.NSSize {
	ret := C.NSView_inst_intrinsicContentSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSSize)(unsafe.Pointer(&ret))

}

// BaselineOffsetFromBottom The distance (in points) between the bottom of the view’s alignment rectangle and its baseline.
// https://developer.apple.com/documentation/appkit/nsview/1526949-baselineoffsetfrombottom?language=objc
func (x gen_NSView) BaselineOffsetFromBottom() core.CGFloat {
	ret := C.NSView_inst_baselineOffsetFromBottom(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// FirstBaselineOffsetFromTop The distance (in points) between the top of the view’s alignment rectangle and its topmost baseline.
// https://developer.apple.com/documentation/appkit/nsview/1526963-firstbaselineoffsetfromtop?language=objc
func (x gen_NSView) FirstBaselineOffsetFromTop() core.CGFloat {
	ret := C.NSView_inst_firstBaselineOffsetFromTop(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// LastBaselineOffsetFromBottom The distance (in points) between the bottom of the view’s alignment rectangle and its bottommost baseline.
// https://developer.apple.com/documentation/appkit/nsview/1525942-lastbaselineoffsetfrombottom?language=objc
func (x gen_NSView) LastBaselineOffsetFromBottom() core.CGFloat {
	ret := C.NSView_inst_lastBaselineOffsetFromBottom(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// NeedsLayout A Boolean value indicating whether the view needs a layout pass before it can be drawn.
// https://developer.apple.com/documentation/appkit/nsview/1526912-needslayout?language=objc
func (x gen_NSView) NeedsLayout() bool {
	ret := C.NSView_inst_needsLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetNeedsLayout A Boolean value indicating whether the view needs a layout pass before it can be drawn.
// https://developer.apple.com/documentation/appkit/nsview/1526912-needslayout?language=objc
func (x gen_NSView) SetNeedsLayout(
	value bool,
) {
	C.NSView_inst_setNeedsLayout(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// NeedsUpdateConstraints A Boolean value indicating whether the view’s constraints need to be updated.
// https://developer.apple.com/documentation/appkit/nsview/1526856-needsupdateconstraints?language=objc
func (x gen_NSView) NeedsUpdateConstraints() bool {
	ret := C.NSView_inst_needsUpdateConstraints(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetNeedsUpdateConstraints A Boolean value indicating whether the view’s constraints need to be updated.
// https://developer.apple.com/documentation/appkit/nsview/1526856-needsupdateconstraints?language=objc
func (x gen_NSView) SetNeedsUpdateConstraints(
	value bool,
) {
	C.NSView_inst_setNeedsUpdateConstraints(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// TranslatesAutoresizingMaskIntoConstraints A Boolean value indicating whether the view’s autoresizing mask is translated into constraints for the constraint-based layout system.
// https://developer.apple.com/documentation/appkit/nsview/1526961-translatesautoresizingmaskintoco?language=objc
func (x gen_NSView) TranslatesAutoresizingMaskIntoConstraints() bool {
	ret := C.NSView_inst_translatesAutoresizingMaskIntoConstraints(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetTranslatesAutoresizingMaskIntoConstraints A Boolean value indicating whether the view’s autoresizing mask is translated into constraints for the constraint-based layout system.
// https://developer.apple.com/documentation/appkit/nsview/1526961-translatesautoresizingmaskintoco?language=objc
func (x gen_NSView) SetTranslatesAutoresizingMaskIntoConstraints(
	value bool,
) {
	C.NSView_inst_setTranslatesAutoresizingMaskIntoConstraints(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// HasAmbiguousLayout A Boolean value indicating whether the constraints impacting the layout of the view incompletely specify the location of the view.
// https://developer.apple.com/documentation/appkit/nsview/1526907-hasambiguouslayout?language=objc
func (x gen_NSView) HasAmbiguousLayout() bool {
	ret := C.NSView_inst_hasAmbiguousLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// FocusRingMaskBounds The focus ring mask bounds, specified in the view’s coordinate space.
// https://developer.apple.com/documentation/appkit/nsview/1483287-focusringmaskbounds?language=objc
func (x gen_NSView) FocusRingMaskBounds() core.NSRect {
	ret := C.NSView_inst_focusRingMaskBounds(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// AllowsVibrancy A Boolean value indicating whether the view ensures it is vibrant on top of other content.
// https://developer.apple.com/documentation/appkit/nsview/1483793-allowsvibrancy?language=objc
func (x gen_NSView) AllowsVibrancy() bool {
	ret := C.NSView_inst_allowsVibrancy(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsInFullScreenMode A Boolean value indicating whether the view is in full screen mode.
// https://developer.apple.com/documentation/appkit/nsview/1483337-infullscreenmode?language=objc
func (x gen_NSView) IsInFullScreenMode() bool {
	ret := C.NSView_inst_isInFullScreenMode(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsHidden A Boolean value indicating whether the view is hidden.
// https://developer.apple.com/documentation/appkit/nsview/1483369-hidden?language=objc
func (x gen_NSView) IsHidden() bool {
	ret := C.NSView_inst_isHidden(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetHidden A Boolean value indicating whether the view is hidden.
// https://developer.apple.com/documentation/appkit/nsview/1483369-hidden?language=objc
func (x gen_NSView) SetHidden(
	value bool,
) {
	C.NSView_inst_setHidden(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsHiddenOrHasHiddenAncestor A Boolean value indicating whether the view is hidden from sight because it, or one of its ancestors, is marked as hidden.
// https://developer.apple.com/documentation/appkit/nsview/1483473-hiddenorhashiddenancestor?language=objc
func (x gen_NSView) IsHiddenOrHasHiddenAncestor() bool {
	ret := C.NSView_inst_isHiddenOrHasHiddenAncestor(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// InLiveResize A Boolean value indicating whether the view is being rendered as part of a live resizing operation.
// https://developer.apple.com/documentation/appkit/nsview/1483267-inliveresize?language=objc
func (x gen_NSView) InLiveResize() bool {
	ret := C.NSView_inst_inLiveResize(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// PreservesContentDuringLiveResize A Boolean value indicating whether the view optimizes live-resize operations by preserving content that has not moved.
// https://developer.apple.com/documentation/appkit/nsview/1483795-preservescontentduringliveresize?language=objc
func (x gen_NSView) PreservesContentDuringLiveResize() bool {
	ret := C.NSView_inst_preservesContentDuringLiveResize(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// RectPreservedDuringLiveResize The rectangle identifying the portion of your view that did not change during a live resize operation.
// https://developer.apple.com/documentation/appkit/nsview/1483528-rectpreservedduringliveresize?language=objc
func (x gen_NSView) RectPreservedDuringLiveResize() core.NSRect {
	ret := C.NSView_inst_rectPreservedDuringLiveResize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// GestureRecognizers The gesture recognize objects currently attached to the view.
// https://developer.apple.com/documentation/appkit/nsview/1483658-gesturerecognizers?language=objc
func (x gen_NSView) GestureRecognizers() core.NSArray {
	ret := C.NSView_inst_gestureRecognizers(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// SetGestureRecognizers The gesture recognize objects currently attached to the view.
// https://developer.apple.com/documentation/appkit/nsview/1483658-gesturerecognizers?language=objc
func (x gen_NSView) SetGestureRecognizers(
	value core.NSArrayRef,
) {
	C.NSView_inst_setGestureRecognizers(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// MouseDownCanMoveWindow A Boolean value indicating whether the view can pass mouse down events through to its superviews.
// https://developer.apple.com/documentation/appkit/nsview/1483666-mousedowncanmovewindow?language=objc
func (x gen_NSView) MouseDownCanMoveWindow() bool {
	ret := C.NSView_inst_mouseDownCanMoveWindow(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// WantsRestingTouches A Boolean value indicating whether the view wants resting touches.
// https://developer.apple.com/documentation/appkit/nsview/1483594-wantsrestingtouches?language=objc
func (x gen_NSView) WantsRestingTouches() bool {
	ret := C.NSView_inst_wantsRestingTouches(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetWantsRestingTouches A Boolean value indicating whether the view wants resting touches.
// https://developer.apple.com/documentation/appkit/nsview/1483594-wantsrestingtouches?language=objc
func (x gen_NSView) SetWantsRestingTouches(
	value bool,
) {
	C.NSView_inst_setWantsRestingTouches(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// CanBecomeKeyView A Boolean value indicating whether the view can become key view.
// https://developer.apple.com/documentation/appkit/nsview/1483759-canbecomekeyview?language=objc
func (x gen_NSView) CanBecomeKeyView() bool {
	ret := C.NSView_inst_canBecomeKeyView(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// NeedsPanelToBecomeKey A Boolean value indicating whether the view needs its panel to become the key window before it can handle keyboard input and navigation.
// https://developer.apple.com/documentation/appkit/nsview/1483512-needspaneltobecomekey?language=objc
func (x gen_NSView) NeedsPanelToBecomeKey() bool {
	ret := C.NSView_inst_needsPanelToBecomeKey(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// NextKeyView The view object that follows the current view in the key view loop.
// https://developer.apple.com/documentation/appkit/nsview/1483465-nextkeyview?language=objc
func (x gen_NSView) NextKeyView() NSView {
	ret := C.NSView_inst_nextKeyView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_fromPointer(ret)

}

// SetNextKeyView The view object that follows the current view in the key view loop.
// https://developer.apple.com/documentation/appkit/nsview/1483465-nextkeyview?language=objc
func (x gen_NSView) SetNextKeyView(
	value NSViewRef,
) {
	C.NSView_inst_setNextKeyView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// NextValidKeyView The closest view object in the key view loop that follows the current view in the key view loop and accepts first responder status.
// https://developer.apple.com/documentation/appkit/nsview/1483572-nextvalidkeyview?language=objc
func (x gen_NSView) NextValidKeyView() NSView {
	ret := C.NSView_inst_nextValidKeyView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_fromPointer(ret)

}

// PreviousKeyView The view object preceding the current view in the key view loop.
// https://developer.apple.com/documentation/appkit/nsview/1483646-previouskeyview?language=objc
func (x gen_NSView) PreviousKeyView() NSView {
	ret := C.NSView_inst_previousKeyView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_fromPointer(ret)

}

// PreviousValidKeyView The closest view object in the key view loop that precedes the current view and accepts first responder status.
// https://developer.apple.com/documentation/appkit/nsview/1483371-previousvalidkeyview?language=objc
func (x gen_NSView) PreviousValidKeyView() NSView {
	ret := C.NSView_inst_previousValidKeyView(
		unsafe.Pointer(x.Pointer()),
	)

	return NSView_fromPointer(ret)

}

// PreparedContentRect The portion of the view that has been rendered and is available for responsive scrolling.
// https://developer.apple.com/documentation/appkit/nsview/1483215-preparedcontentrect?language=objc
func (x gen_NSView) PreparedContentRect() core.NSRect {
	ret := C.NSView_inst_preparedContentRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*core.NSRect)(unsafe.Pointer(&ret))

}

// SetPreparedContentRect The portion of the view that has been rendered and is available for responsive scrolling.
// https://developer.apple.com/documentation/appkit/nsview/1483215-preparedcontentrect?language=objc
func (x gen_NSView) SetPreparedContentRect(
	value core.NSRect,
) {
	C.NSView_inst_setPreparedContentRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return

}

// RegisteredDraggedTypes The array of pasteboard drag types that the view can accept.
// https://developer.apple.com/documentation/appkit/nsview/1483564-registereddraggedtypes?language=objc
func (x gen_NSView) RegisteredDraggedTypes() core.NSArray {
	ret := C.NSView_inst_registeredDraggedTypes(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// PostsFrameChangedNotifications A Boolean value indicating whether the view posts notifications when its frame rectangle changes.
// https://developer.apple.com/documentation/appkit/nsview/1483524-postsframechangednotifications?language=objc
func (x gen_NSView) PostsFrameChangedNotifications() bool {
	ret := C.NSView_inst_postsFrameChangedNotifications(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetPostsFrameChangedNotifications A Boolean value indicating whether the view posts notifications when its frame rectangle changes.
// https://developer.apple.com/documentation/appkit/nsview/1483524-postsframechangednotifications?language=objc
func (x gen_NSView) SetPostsFrameChangedNotifications(
	value bool,
) {
	C.NSView_inst_setPostsFrameChangedNotifications(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// PostsBoundsChangedNotifications A Boolean value indicating whether the view posts notifications when its bounds rectangle changes.
// https://developer.apple.com/documentation/appkit/nsview/1483239-postsboundschangednotifications?language=objc
func (x gen_NSView) PostsBoundsChangedNotifications() bool {
	ret := C.NSView_inst_postsBoundsChangedNotifications(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetPostsBoundsChangedNotifications A Boolean value indicating whether the view posts notifications when its bounds rectangle changes.
// https://developer.apple.com/documentation/appkit/nsview/1483239-postsboundschangednotifications?language=objc
func (x gen_NSView) SetPostsBoundsChangedNotifications(
	value bool,
) {
	C.NSView_inst_setPostsBoundsChangedNotifications(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// Tag The view’s tag, which is an integer that you use to identify the view within your app.
// https://developer.apple.com/documentation/appkit/nsview/1483248-tag?language=objc
func (x gen_NSView) Tag() core.NSInteger {
	ret := C.NSView_inst_tag(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSInteger(ret)

}

// ToolTip The text for the view’s tooltip.
// https://developer.apple.com/documentation/appkit/nsview/1483541-tooltip?language=objc
func (x gen_NSView) ToolTip() core.NSString {
	ret := C.NSView_inst_toolTip(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetToolTip The text for the view’s tooltip.
// https://developer.apple.com/documentation/appkit/nsview/1483541-tooltip?language=objc
func (x gen_NSView) SetToolTip(
	value core.NSStringRef,
) {
	C.NSView_inst_setToolTip(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// TrackingAreas An array of the view’s tracking areas.
// https://developer.apple.com/documentation/appkit/nsview/1483333-trackingareas?language=objc
func (x gen_NSView) TrackingAreas() core.NSArray {
	ret := C.NSView_inst_trackingAreas(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_fromPointer(ret)

}

// IsDrawingFindIndicator A Boolean value indicating whether the view or one of its ancestors is being drawn for a find indicator.
// https://developer.apple.com/documentation/appkit/nsview/1483317-drawingfindindicator?language=objc
func (x gen_NSView) IsDrawingFindIndicator() bool {
	ret := C.NSView_inst_isDrawingFindIndicator(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsHorizontalContentSizeConstraintActive
// https://developer.apple.com/documentation/appkit/nsview/3353053-horizontalcontentsizeconstrainta?language=objc
func (x gen_NSView) IsHorizontalContentSizeConstraintActive() bool {
	ret := C.NSView_inst_isHorizontalContentSizeConstraintActive(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetHorizontalContentSizeConstraintActive
// https://developer.apple.com/documentation/appkit/nsview/3353053-horizontalcontentsizeconstrainta?language=objc
func (x gen_NSView) SetHorizontalContentSizeConstraintActive(
	value bool,
) {
	C.NSView_inst_setHorizontalContentSizeConstraintActive(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsVerticalContentSizeConstraintActive
// https://developer.apple.com/documentation/appkit/nsview/3353054-verticalcontentsizeconstraintact?language=objc
func (x gen_NSView) IsVerticalContentSizeConstraintActive() bool {
	ret := C.NSView_inst_isVerticalContentSizeConstraintActive(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetVerticalContentSizeConstraintActive
// https://developer.apple.com/documentation/appkit/nsview/3353054-verticalcontentsizeconstraintact?language=objc
func (x gen_NSView) SetVerticalContentSizeConstraintActive(
	value bool,
) {
	C.NSView_inst_setVerticalContentSizeConstraintActive(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// BackgroundColor
func (x gen_NSView) BackgroundColor() NSColor {
	ret := C.NSView_inst_backgroundColor(
		unsafe.Pointer(x.Pointer()),
	)

	return NSColor_fromPointer(ret)

}

// SetBackgroundColor
func (x gen_NSView) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSView_inst_setBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}
