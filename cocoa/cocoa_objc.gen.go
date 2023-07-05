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


void* NSBundle_inst_initWithURL(void *id, void* url) {
	return [(NSBundle*)id
		initWithURL: url];
}

void* NSBundle_inst_initWithPath(void *id, void* path) {
	return [(NSBundle*)id
		initWithPath: path];
}

void* NSBundle_inst_loadNibNamed_owner_options(void *id, void* name, void* owner, void* options) {
	return [(NSBundle*)id
		loadNibNamed: name
		owner: owner
		options: options];
}

void* NSBundle_inst_URLForResource_withExtension_subdirectory(void *id, void* name, void* ext, void* subpath) {
	return [(NSBundle*)id
		URLForResource: name
		withExtension: ext
		subdirectory: subpath];
}

void* NSBundle_inst_URLForResource_withExtension(void *id, void* name, void* ext) {
	return [(NSBundle*)id
		URLForResource: name
		withExtension: ext];
}

void* NSBundle_inst_URLsForResourcesWithExtension_subdirectory(void *id, void* ext, void* subpath) {
	return [(NSBundle*)id
		URLsForResourcesWithExtension: ext
		subdirectory: subpath];
}

void* NSBundle_inst_URLForResource_withExtension_subdirectory_localization(void *id, void* name, void* ext, void* subpath, void* localizationName) {
	return [(NSBundle*)id
		URLForResource: name
		withExtension: ext
		subdirectory: subpath
		localization: localizationName];
}

void* NSBundle_inst_URLsForResourcesWithExtension_subdirectory_localization(void *id, void* ext, void* subpath, void* localizationName) {
	return [(NSBundle*)id
		URLsForResourcesWithExtension: ext
		subdirectory: subpath
		localization: localizationName];
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

void* NSBundle_inst_localizedStringForKey_value_table(void *id, void* key, void* value, void* tableName) {
	return [(NSBundle*)id
		localizedStringForKey: key
		value: value
		table: tableName];
}

void* NSBundle_inst_URLForAuxiliaryExecutable(void *id, void* executableName) {
	return [(NSBundle*)id
		URLForAuxiliaryExecutable: executableName];
}

void* NSBundle_inst_pathForAuxiliaryExecutable(void *id, void* executableName) {
	return [(NSBundle*)id
		pathForAuxiliaryExecutable: executableName];
}

void* NSBundle_inst_objectForInfoDictionaryKey(void *id, void* key) {
	return [(NSBundle*)id
		objectForInfoDictionaryKey: key];
}

BOOL NSBundle_inst_load(void *id) {
	return [(NSBundle*)id
		load];
}

BOOL NSBundle_inst_unload(void *id) {
	return [(NSBundle*)id
		unload];
}

void* NSBundle_inst_localizedAttributedStringForKey_value_table(void *id, void* key, void* value, void* tableName) {
	return [(NSBundle*)id
		localizedAttributedStringForKey: key
		value: value
		table: tableName];
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

void NSApplication_inst_run(void *id) {
	[(NSApplication*)id
		run];
}

void NSApplication_inst_finishLaunching(void *id) {
	[(NSApplication*)id
		finishLaunching];
}

void NSApplication_inst_stop(void *id, void* sender) {
	[(NSApplication*)id
		stop: sender];
}

void NSApplication_inst_sendEvent(void *id, void* event) {
	[(NSApplication*)id
		sendEvent: event];
}

void NSApplication_inst_postEvent_atStart(void *id, void* event, BOOL flag) {
	[(NSApplication*)id
		postEvent: event
		atStart: flag];
}

BOOL NSApplication_inst_tryToPerform_with(void *id, void* action, void* object) {
	return [(NSApplication*)id
		tryToPerform: action
		with: object];
}

BOOL NSApplication_inst_sendAction_to_from(void *id, void* action, void* target, void* sender) {
	return [(NSApplication*)id
		sendAction: action
		to: target
		from: sender];
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

void NSApplication_inst_replyToApplicationShouldTerminate(void *id, BOOL shouldTerminate) {
	[(NSApplication*)id
		replyToApplicationShouldTerminate: shouldTerminate];
}

void NSApplication_inst_activateIgnoringOtherApps(void *id, BOOL flag) {
	[(NSApplication*)id
		activateIgnoringOtherApps: flag];
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

void NSApplication_inst_registerForRemoteNotifications(void *id) {
	[(NSApplication*)id
		registerForRemoteNotifications];
}

void NSApplication_inst_unregisterForRemoteNotifications(void *id) {
	[(NSApplication*)id
		unregisterForRemoteNotifications];
}

void NSApplication_inst_toggleTouchBarCustomizationPalette(void *id, void* sender) {
	[(NSApplication*)id
		toggleTouchBarCustomizationPalette: sender];
}

void NSApplication_inst_cancelUserAttentionRequest(void *id, long request) {
	[(NSApplication*)id
		cancelUserAttentionRequest: request];
}

void NSApplication_inst_registerUserInterfaceItemSearchHandler(void *id, void* handler) {
	[(NSApplication*)id
		registerUserInterfaceItemSearchHandler: handler];
}

void NSApplication_inst_unregisterUserInterfaceItemSearchHandler(void *id, void* handler) {
	[(NSApplication*)id
		unregisterUserInterfaceItemSearchHandler: handler];
}

void NSApplication_inst_showHelp(void *id, void* sender) {
	[(NSApplication*)id
		showHelp: sender];
}

void NSApplication_inst_activateContextHelpMode(void *id, void* sender) {
	[(NSApplication*)id
		activateContextHelpMode: sender];
}

void NSApplication_inst_hideOtherApplications(void *id, void* sender) {
	[(NSApplication*)id
		hideOtherApplications: sender];
}

void NSApplication_inst_unhideAllApplications(void *id, void* sender) {
	[(NSApplication*)id
		unhideAllApplications: sender];
}

long NSApplication_inst_activationPolicy(void *id) {
	return [(NSApplication*)id
		activationPolicy];
}

BOOL NSApplication_inst_setActivationPolicy(void *id, long activationPolicy) {
	return [(NSApplication*)id
		setActivationPolicy: activationPolicy];
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

void* NSControl_inst_initWithFrame(void *id, NSRect frameRect) {
	return [(NSControl*)id
		initWithFrame: frameRect];
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

void NSControl_inst_drawWithExpansionFrame_inView(void *id, NSRect contentFrame, void* view) {
	[(NSControl*)id
		drawWithExpansionFrame: contentFrame
		inView: view];
}

NSRect NSControl_inst_expansionFrameWithFrame(void *id, NSRect contentFrame) {
	return [(NSControl*)id
		expansionFrameWithFrame: contentFrame];
}

BOOL NSControl_inst_abortEditing(void *id) {
	return [(NSControl*)id
		abortEditing];
}

void* NSControl_inst_currentEditor(void *id) {
	return [(NSControl*)id
		currentEditor];
}

void NSControl_inst_validateEditing(void *id) {
	[(NSControl*)id
		validateEditing];
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

void NSControl_inst_selectWithFrame_editor_delegate_start_length(void *id, NSRect rect, void* textObj, void* delegate, long selStart, long selLength) {
	[(NSControl*)id
		selectWithFrame: rect
		editor: textObj
		delegate: delegate
		start: selStart
		length: selLength];
}

NSSize NSControl_inst_sizeThatFits(void *id, NSSize size) {
	return [(NSControl*)id
		sizeThatFits: size];
}

void NSControl_inst_sizeToFit(void *id) {
	[(NSControl*)id
		sizeToFit];
}

BOOL NSControl_inst_sendAction_to(void *id, void* action, void* target) {
	return [(NSControl*)id
		sendAction: action
		to: target];
}

void NSControl_inst_performClick(void *id, void* sender) {
	[(NSControl*)id
		performClick: sender];
}

void NSControl_inst_mouseDown(void *id, void* event) {
	[(NSControl*)id
		mouseDown: event];
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

NSSize NSButton_inst_minimumSizeWithPrioritizedCompressionOptions(void *id, void* prioritizedOptions) {
	return [(NSButton*)id
		minimumSizeWithPrioritizedCompressionOptions: prioritizedOptions];
}

void NSButton_inst_setNextState(void *id) {
	[(NSButton*)id
		setNextState];
}

void NSButton_inst_highlight(void *id, BOOL flag) {
	[(NSButton*)id
		highlight: flag];
}

BOOL NSButton_inst_performKeyEquivalent(void *id, void* key) {
	return [(NSButton*)id
		performKeyEquivalent: key];
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

void NSFont_inst_set(void *id) {
	[(NSFont*)id
		set];
}

void* NSFont_inst_fontWithSize(void *id, double fontSize) {
	return [(NSFont*)id
		fontWithSize: fontSize];
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

void NSImage_inst_addRepresentations(void *id, void* imageReps) {
	[(NSImage*)id
		addRepresentations: imageReps];
}

void NSImage_inst_drawInRect(void *id, NSRect rect) {
	[(NSImage*)id
		drawInRect: rect];
}

void NSImage_inst_lockFocus(void *id) {
	[(NSImage*)id
		lockFocus];
}

void NSImage_inst_lockFocusFlipped(void *id, BOOL flipped) {
	[(NSImage*)id
		lockFocusFlipped: flipped];
}

void NSImage_inst_unlockFocus(void *id) {
	[(NSImage*)id
		unlockFocus];
}

void NSImage_inst_recache(void *id) {
	[(NSImage*)id
		recache];
}

void NSImage_inst_cancelIncrementalLoad(void *id) {
	[(NSImage*)id
		cancelIncrementalLoad];
}

void* NSImage_inst_layerContentsForContentsScale(void *id, double layerContentsScale) {
	return [(NSImage*)id
		layerContentsForContentsScale: layerContentsScale];
}

double NSImage_inst_recommendedLayerContentsScale(void *id, double preferredContentsScale) {
	return [(NSImage*)id
		recommendedLayerContentsScale: preferredContentsScale];
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

void NSPasteboard_inst_releaseGlobally(void *id) {
	[(NSPasteboard*)id
		releaseGlobally];
}

long NSPasteboard_inst_clearContents(void *id) {
	return [(NSPasteboard*)id
		clearContents];
}

BOOL NSPasteboard_inst_writeObjects(void *id, void* objects) {
	return [(NSPasteboard*)id
		writeObjects: objects];
}

void* NSPasteboard_inst_readObjectsForClasses_options(void *id, void* classArray, void* options) {
	return [(NSPasteboard*)id
		readObjectsForClasses: classArray
		options: options];
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

long NSPasteboard_inst_declareTypes_owner(void *id, void* newTypes, void* newOwner) {
	return [(NSPasteboard*)id
		declareTypes: newTypes
		owner: newOwner];
}

long NSPasteboard_inst_addTypes_owner(void *id, void* newTypes, void* newOwner) {
	return [(NSPasteboard*)id
		addTypes: newTypes
		owner: newOwner];
}

BOOL NSPasteboard_inst_writeFileContents(void *id, void* filename) {
	return [(NSPasteboard*)id
		writeFileContents: filename];
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

void* NSLayoutManager_inst_init(void *id) {
	return [(NSLayoutManager*)id
		init];
}

void NSLayoutManager_inst_addTextContainer(void *id, void* container) {
	[(NSLayoutManager*)id
		addTextContainer: container];
}

void NSLayoutManager_inst_insertTextContainer_atIndex(void *id, void* container, unsigned long index) {
	[(NSLayoutManager*)id
		insertTextContainer: container
		atIndex: index];
}

void NSLayoutManager_inst_removeTextContainerAtIndex(void *id, unsigned long index) {
	[(NSLayoutManager*)id
		removeTextContainerAtIndex: index];
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

void NSLayoutManager_inst_ensureLayoutForBoundingRect_inTextContainer(void *id, NSRect bounds, void* container) {
	[(NSLayoutManager*)id
		ensureLayoutForBoundingRect: bounds
		inTextContainer: container];
}

void NSLayoutManager_inst_ensureLayoutForTextContainer(void *id, void* container) {
	[(NSLayoutManager*)id
		ensureLayoutForTextContainer: container];
}

unsigned long NSLayoutManager_inst_characterIndexForGlyphAtIndex(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		characterIndexForGlyphAtIndex: glyphIndex];
}

unsigned long NSLayoutManager_inst_glyphIndexForCharacterAtIndex(void *id, unsigned long charIndex) {
	return [(NSLayoutManager*)id
		glyphIndexForCharacterAtIndex: charIndex];
}

BOOL NSLayoutManager_inst_isValidGlyphIndex(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		isValidGlyphIndex: glyphIndex];
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

NSSize NSLayoutManager_inst_attachmentSizeForGlyphAtIndex(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		attachmentSizeForGlyphAtIndex: glyphIndex];
}

BOOL NSLayoutManager_inst_drawsOutsideLineFragmentForGlyphAtIndex(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		drawsOutsideLineFragmentForGlyphAtIndex: glyphIndex];
}

unsigned long NSLayoutManager_inst_firstUnlaidCharacterIndex(void *id) {
	return [(NSLayoutManager*)id
		firstUnlaidCharacterIndex];
}

unsigned long NSLayoutManager_inst_firstUnlaidGlyphIndex(void *id) {
	return [(NSLayoutManager*)id
		firstUnlaidGlyphIndex];
}

BOOL NSLayoutManager_inst_notShownAttributeForGlyphAtIndex(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		notShownAttributeForGlyphAtIndex: glyphIndex];
}

BOOL NSLayoutManager_inst_layoutManagerOwnsFirstResponderInWindow(void *id, void* window) {
	return [(NSLayoutManager*)id
		layoutManagerOwnsFirstResponderInWindow: window];
}

double NSLayoutManager_inst_defaultLineHeightForFont(void *id, void* theFont) {
	return [(NSLayoutManager*)id
		defaultLineHeightForFont: theFont];
}

double NSLayoutManager_inst_defaultBaselineOffsetForFont(void *id, void* theFont) {
	return [(NSLayoutManager*)id
		defaultBaselineOffsetForFont: theFont];
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

void NSMenu_inst_removeItem(void *id, void* item) {
	[(NSMenu*)id
		removeItem: item];
}

void NSMenu_inst_removeItemAtIndex(void *id, long index) {
	[(NSMenu*)id
		removeItemAtIndex: index];
}

void NSMenu_inst_itemChanged(void *id, void* item) {
	[(NSMenu*)id
		itemChanged: item];
}

void NSMenu_inst_removeAllItems(void *id) {
	[(NSMenu*)id
		removeAllItems];
}

void* NSMenu_inst_itemWithTag(void *id, long tag) {
	return [(NSMenu*)id
		itemWithTag: tag];
}

void* NSMenu_inst_itemWithTitle(void *id, void* title) {
	return [(NSMenu*)id
		itemWithTitle: title];
}

void* NSMenu_inst_itemAtIndex(void *id, long index) {
	return [(NSMenu*)id
		itemAtIndex: index];
}

long NSMenu_inst_indexOfItem(void *id, void* item) {
	return [(NSMenu*)id
		indexOfItem: item];
}

long NSMenu_inst_indexOfItemWithTitle(void *id, void* title) {
	return [(NSMenu*)id
		indexOfItemWithTitle: title];
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

long NSMenu_inst_indexOfItemWithRepresentedObject(void *id, void* object) {
	return [(NSMenu*)id
		indexOfItemWithRepresentedObject: object];
}

long NSMenu_inst_indexOfItemWithSubmenu(void *id, void* submenu) {
	return [(NSMenu*)id
		indexOfItemWithSubmenu: submenu];
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

BOOL NSMenu_inst_performKeyEquivalent(void *id, void* event) {
	return [(NSMenu*)id
		performKeyEquivalent: event];
}

void NSMenu_inst_performActionForItemAtIndex(void *id, long index) {
	[(NSMenu*)id
		performActionForItemAtIndex: index];
}

BOOL NSMenu_inst_popUpMenuPositioningItem_atLocation_inView(void *id, void* item, NSPoint location, void* view) {
	return [(NSMenu*)id
		popUpMenuPositioningItem: item
		atLocation: location
		inView: view];
}

void NSMenu_inst_cancelTracking(void *id) {
	[(NSMenu*)id
		cancelTracking];
}

void NSMenu_inst_cancelTrackingWithoutAnimation(void *id) {
	[(NSMenu*)id
		cancelTrackingWithoutAnimation];
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

void NSPopover_inst_performClose(void *id, void* sender) {
	[(NSPopover*)id
		performClose: sender];
}

void NSPopover_inst_close(void *id) {
	[(NSPopover*)id
		close];
}

void* NSPopover_inst_init(void *id) {
	return [(NSPopover*)id
		init];
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

BOOL NSRunningApplication_inst_hide(void *id) {
	return [(NSRunningApplication*)id
		hide];
}

BOOL NSRunningApplication_inst_unhide(void *id) {
	return [(NSRunningApplication*)id
		unhide];
}

BOOL NSRunningApplication_inst_forceTerminate(void *id) {
	return [(NSRunningApplication*)id
		forceTerminate];
}

BOOL NSRunningApplication_inst_terminate(void *id) {
	return [(NSRunningApplication*)id
		terminate];
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

void* NSStatusBar_inst_statusItemWithLength(void *id, double length) {
	return [(NSStatusBar*)id
		statusItemWithLength: length];
}

void NSStatusBar_inst_removeStatusItem(void *id, void* item) {
	[(NSStatusBar*)id
		removeStatusItem: item];
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

void* NSText_inst_initWithFrame(void *id, NSRect frameRect) {
	return [(NSText*)id
		initWithFrame: frameRect];
}

void NSText_inst_toggleRuler(void *id, void* sender) {
	[(NSText*)id
		toggleRuler: sender];
}

void NSText_inst_selectAll(void *id, void* sender) {
	[(NSText*)id
		selectAll: sender];
}

void NSText_inst_copy(void *id, void* sender) {
	[(NSText*)id
		copy: sender];
}

void NSText_inst_cut(void *id, void* sender) {
	[(NSText*)id
		cut: sender];
}

void NSText_inst_paste(void *id, void* sender) {
	[(NSText*)id
		paste: sender];
}

void NSText_inst_copyFont(void *id, void* sender) {
	[(NSText*)id
		copyFont: sender];
}

void NSText_inst_pasteFont(void *id, void* sender) {
	[(NSText*)id
		pasteFont: sender];
}

void NSText_inst_copyRuler(void *id, void* sender) {
	[(NSText*)id
		copyRuler: sender];
}

void NSText_inst_pasteRuler(void *id, void* sender) {
	[(NSText*)id
		pasteRuler: sender];
}

void NSText_inst_delete(void *id, void* sender) {
	[(NSText*)id
		delete: sender];
}

void NSText_inst_changeFont(void *id, void* sender) {
	[(NSText*)id
		changeFont: sender];
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

void NSText_inst_superscript(void *id, void* sender) {
	[(NSText*)id
		superscript: sender];
}

void NSText_inst_subscript(void *id, void* sender) {
	[(NSText*)id
		subscript: sender];
}

void NSText_inst_unscript(void *id, void* sender) {
	[(NSText*)id
		unscript: sender];
}

void NSText_inst_underline(void *id, void* sender) {
	[(NSText*)id
		underline: sender];
}

BOOL NSText_inst_readRTFDFromFile(void *id, void* path) {
	return [(NSText*)id
		readRTFDFromFile: path];
}

BOOL NSText_inst_writeRTFDToFile_atomically(void *id, void* path, BOOL flag) {
	return [(NSText*)id
		writeRTFDToFile: path
		atomically: flag];
}

void NSText_inst_checkSpelling(void *id, void* sender) {
	[(NSText*)id
		checkSpelling: sender];
}

void NSText_inst_showGuessPanel(void *id, void* sender) {
	[(NSText*)id
		showGuessPanel: sender];
}

void NSText_inst_sizeToFit(void *id) {
	[(NSText*)id
		sizeToFit];
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

void NSViewController_inst_loadView(void *id) {
	[(NSViewController*)id
		loadView];
}

void NSViewController_inst_commitEditingWithDelegate_didCommitSelector_contextInfo(void *id, void* delegate, void* didCommitSelector, void* contextInfo) {
	[(NSViewController*)id
		commitEditingWithDelegate: delegate
		didCommitSelector: didCommitSelector
		contextInfo: contextInfo];
}

BOOL NSViewController_inst_commitEditing(void *id) {
	return [(NSViewController*)id
		commitEditing];
}

void NSViewController_inst_discardEditing(void *id) {
	[(NSViewController*)id
		discardEditing];
}

void NSViewController_inst_dismissController(void *id, void* sender) {
	[(NSViewController*)id
		dismissController: sender];
}

void NSViewController_inst_viewDidLoad(void *id) {
	[(NSViewController*)id
		viewDidLoad];
}

void NSViewController_inst_viewWillAppear(void *id) {
	[(NSViewController*)id
		viewWillAppear];
}

void NSViewController_inst_viewDidAppear(void *id) {
	[(NSViewController*)id
		viewDidAppear];
}

void NSViewController_inst_viewWillDisappear(void *id) {
	[(NSViewController*)id
		viewWillDisappear];
}

void NSViewController_inst_viewDidDisappear(void *id) {
	[(NSViewController*)id
		viewDidDisappear];
}

void NSViewController_inst_updateViewConstraints(void *id) {
	[(NSViewController*)id
		updateViewConstraints];
}

void NSViewController_inst_viewWillLayout(void *id) {
	[(NSViewController*)id
		viewWillLayout];
}

void NSViewController_inst_viewDidLayout(void *id) {
	[(NSViewController*)id
		viewDidLayout];
}

void NSViewController_inst_addChildViewController(void *id, void* childViewController) {
	[(NSViewController*)id
		addChildViewController: childViewController];
}

void NSViewController_inst_insertChildViewController_atIndex(void *id, void* childViewController, long index) {
	[(NSViewController*)id
		insertChildViewController: childViewController
		atIndex: index];
}

void NSViewController_inst_removeChildViewControllerAtIndex(void *id, long index) {
	[(NSViewController*)id
		removeChildViewControllerAtIndex: index];
}

void NSViewController_inst_removeFromParentViewController(void *id) {
	[(NSViewController*)id
		removeFromParentViewController];
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

void NSViewController_inst_dismissViewController(void *id, void* viewController) {
	[(NSViewController*)id
		dismissViewController: viewController];
}

void NSViewController_inst_presentViewControllerAsModalWindow(void *id, void* viewController) {
	[(NSViewController*)id
		presentViewControllerAsModalWindow: viewController];
}

void NSViewController_inst_presentViewControllerAsSheet(void *id, void* viewController) {
	[(NSViewController*)id
		presentViewControllerAsSheet: viewController];
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

void NSWindow_inst_toggleFullScreen(void *id, void* sender) {
	[(NSWindow*)id
		toggleFullScreen: sender];
}

void NSWindow_inst_setDynamicDepthLimit(void *id, BOOL flag) {
	[(NSWindow*)id
		setDynamicDepthLimit: flag];
}

void NSWindow_inst_invalidateShadow(void *id) {
	[(NSWindow*)id
		invalidateShadow];
}

NSRect NSWindow_inst_contentRectForFrameRect(void *id, NSRect frameRect) {
	return [(NSWindow*)id
		contentRectForFrameRect: frameRect];
}

NSRect NSWindow_inst_frameRectForContentRect(void *id, NSRect contentRect) {
	return [(NSWindow*)id
		frameRectForContentRect: contentRect];
}

void NSWindow_inst_endSheet(void *id, void* sheetWindow) {
	[(NSWindow*)id
		endSheet: sheetWindow];
}

void NSWindow_inst_setFrameOrigin(void *id, NSPoint point) {
	[(NSWindow*)id
		setFrameOrigin: point];
}

void NSWindow_inst_setFrameTopLeftPoint(void *id, NSPoint point) {
	[(NSWindow*)id
		setFrameTopLeftPoint: point];
}

NSRect NSWindow_inst_constrainFrameRect_toScreen(void *id, NSRect frameRect, void* screen) {
	return [(NSWindow*)id
		constrainFrameRect: frameRect
		toScreen: screen];
}

NSPoint NSWindow_inst_cascadeTopLeftFromPoint(void *id, NSPoint topLeftPoint) {
	return [(NSWindow*)id
		cascadeTopLeftFromPoint: topLeftPoint];
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

void NSWindow_inst_performZoom(void *id, void* sender) {
	[(NSWindow*)id
		performZoom: sender];
}

void NSWindow_inst_zoom(void *id, void* sender) {
	[(NSWindow*)id
		zoom: sender];
}

void NSWindow_inst_setContentSize(void *id, NSSize size) {
	[(NSWindow*)id
		setContentSize: size];
}

void NSWindow_inst_orderOut(void *id, void* sender) {
	[(NSWindow*)id
		orderOut: sender];
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

void NSWindow_inst_orderWindow_relativeTo(void *id, unsigned long place, long otherWin) {
	[(NSWindow*)id
		orderWindow: place
		relativeTo: otherWin];
}

void NSWindow_inst_makeKeyWindow(void *id) {
	[(NSWindow*)id
		makeKeyWindow];
}

void NSWindow_inst_makeKeyAndOrderFront(void *id, void* sender) {
	[(NSWindow*)id
		makeKeyAndOrderFront: sender];
}

void NSWindow_inst_becomeKeyWindow(void *id) {
	[(NSWindow*)id
		becomeKeyWindow];
}

void NSWindow_inst_resignKeyWindow(void *id) {
	[(NSWindow*)id
		resignKeyWindow];
}

void NSWindow_inst_makeMainWindow(void *id) {
	[(NSWindow*)id
		makeMainWindow];
}

void NSWindow_inst_becomeMainWindow(void *id) {
	[(NSWindow*)id
		becomeMainWindow];
}

void NSWindow_inst_resignMainWindow(void *id) {
	[(NSWindow*)id
		resignMainWindow];
}

void NSWindow_inst_toggleToolbarShown(void *id, void* sender) {
	[(NSWindow*)id
		toggleToolbarShown: sender];
}

void NSWindow_inst_runToolbarCustomizationPalette(void *id, void* sender) {
	[(NSWindow*)id
		runToolbarCustomizationPalette: sender];
}

void NSWindow_inst_addChildWindow_ordered(void *id, void* childWin, unsigned long place) {
	[(NSWindow*)id
		addChildWindow: childWin
		ordered: place];
}

void NSWindow_inst_removeChildWindow(void *id, void* childWin) {
	[(NSWindow*)id
		removeChildWindow: childWin];
}

void NSWindow_inst_enableKeyEquivalentForDefaultButtonCell(void *id) {
	[(NSWindow*)id
		enableKeyEquivalentForDefaultButtonCell];
}

void NSWindow_inst_disableKeyEquivalentForDefaultButtonCell(void *id) {
	[(NSWindow*)id
		disableKeyEquivalentForDefaultButtonCell];
}

void* NSWindow_inst_fieldEditor_forObject(void *id, BOOL createFlag, void* object) {
	return [(NSWindow*)id
		fieldEditor: createFlag
		forObject: object];
}

void NSWindow_inst_endEditingFor(void *id, void* object) {
	[(NSWindow*)id
		endEditingFor: object];
}

void NSWindow_inst_enableCursorRects(void *id) {
	[(NSWindow*)id
		enableCursorRects];
}

void NSWindow_inst_disableCursorRects(void *id) {
	[(NSWindow*)id
		disableCursorRects];
}

void NSWindow_inst_discardCursorRects(void *id) {
	[(NSWindow*)id
		discardCursorRects];
}

void NSWindow_inst_invalidateCursorRectsForView(void *id, void* view) {
	[(NSWindow*)id
		invalidateCursorRectsForView: view];
}

void NSWindow_inst_resetCursorRects(void *id) {
	[(NSWindow*)id
		resetCursorRects];
}

void NSWindow_inst_removeTitlebarAccessoryViewControllerAtIndex(void *id, long index) {
	[(NSWindow*)id
		removeTitlebarAccessoryViewControllerAtIndex: index];
}

void NSWindow_inst_addTabbedWindow_ordered(void *id, void* window, unsigned long ordered) {
	[(NSWindow*)id
		addTabbedWindow: window
		ordered: ordered];
}

void NSWindow_inst_mergeAllWindows(void *id, void* sender) {
	[(NSWindow*)id
		mergeAllWindows: sender];
}

void NSWindow_inst_selectNextTab(void *id, void* sender) {
	[(NSWindow*)id
		selectNextTab: sender];
}

void NSWindow_inst_selectPreviousTab(void *id, void* sender) {
	[(NSWindow*)id
		selectPreviousTab: sender];
}

void NSWindow_inst_moveTabToNewWindow(void *id, void* sender) {
	[(NSWindow*)id
		moveTabToNewWindow: sender];
}

void NSWindow_inst_toggleTabBar(void *id, void* sender) {
	[(NSWindow*)id
		toggleTabBar: sender];
}

void NSWindow_inst_toggleTabOverview(void *id, void* sender) {
	[(NSWindow*)id
		toggleTabOverview: sender];
}

void NSWindow_inst_postEvent_atStart(void *id, void* event, BOOL flag) {
	[(NSWindow*)id
		postEvent: event
		atStart: flag];
}

void NSWindow_inst_sendEvent(void *id, void* event) {
	[(NSWindow*)id
		sendEvent: event];
}

BOOL NSWindow_inst_tryToPerform_with(void *id, void* action, void* object) {
	return [(NSWindow*)id
		tryToPerform: action
		with: object];
}

void NSWindow_inst_selectKeyViewPrecedingView(void *id, void* view) {
	[(NSWindow*)id
		selectKeyViewPrecedingView: view];
}

void NSWindow_inst_selectKeyViewFollowingView(void *id, void* view) {
	[(NSWindow*)id
		selectKeyViewFollowingView: view];
}

void NSWindow_inst_selectPreviousKeyView(void *id, void* sender) {
	[(NSWindow*)id
		selectPreviousKeyView: sender];
}

void NSWindow_inst_selectNextKeyView(void *id, void* sender) {
	[(NSWindow*)id
		selectNextKeyView: sender];
}

void NSWindow_inst_recalculateKeyViewLoop(void *id) {
	[(NSWindow*)id
		recalculateKeyViewLoop];
}

void NSWindow_inst_performWindowDragWithEvent(void *id, void* event) {
	[(NSWindow*)id
		performWindowDragWithEvent: event];
}

void NSWindow_inst_disableSnapshotRestoration(void *id) {
	[(NSWindow*)id
		disableSnapshotRestoration];
}

void NSWindow_inst_enableSnapshotRestoration(void *id) {
	[(NSWindow*)id
		enableSnapshotRestoration];
}

void NSWindow_inst_display(void *id) {
	[(NSWindow*)id
		display];
}

void NSWindow_inst_displayIfNeeded(void *id) {
	[(NSWindow*)id
		displayIfNeeded];
}

void NSWindow_inst_disableScreenUpdatesUntilFlush(void *id) {
	[(NSWindow*)id
		disableScreenUpdatesUntilFlush];
}

void NSWindow_inst_update(void *id) {
	[(NSWindow*)id
		update];
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

void NSWindow_inst_registerForDraggedTypes(void *id, void* newTypes) {
	[(NSWindow*)id
		registerForDraggedTypes: newTypes];
}

void NSWindow_inst_unregisterDraggedTypes(void *id) {
	[(NSWindow*)id
		unregisterDraggedTypes];
}

NSRect NSWindow_inst_convertRectFromBacking(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectFromBacking: rect];
}

NSRect NSWindow_inst_convertRectFromScreen(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectFromScreen: rect];
}

NSPoint NSWindow_inst_convertPointFromBacking(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointFromBacking: point];
}

NSPoint NSWindow_inst_convertPointFromScreen(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointFromScreen: point];
}

NSRect NSWindow_inst_convertRectToBacking(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectToBacking: rect];
}

NSRect NSWindow_inst_convertRectToScreen(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectToScreen: rect];
}

NSPoint NSWindow_inst_convertPointToBacking(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointToBacking: point];
}

NSPoint NSWindow_inst_convertPointToScreen(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointToScreen: point];
}

void NSWindow_inst_setTitleWithRepresentedFilename(void *id, void* filename) {
	[(NSWindow*)id
		setTitleWithRepresentedFilename: filename];
}

void NSWindow_inst_center(void *id) {
	[(NSWindow*)id
		center];
}

void NSWindow_inst_performClose(void *id, void* sender) {
	[(NSWindow*)id
		performClose: sender];
}

void NSWindow_inst_close(void *id) {
	[(NSWindow*)id
		close];
}

void NSWindow_inst_performMiniaturize(void *id, void* sender) {
	[(NSWindow*)id
		performMiniaturize: sender];
}

void NSWindow_inst_miniaturize(void *id, void* sender) {
	[(NSWindow*)id
		miniaturize: sender];
}

void NSWindow_inst_deminiaturize(void *id, void* sender) {
	[(NSWindow*)id
		deminiaturize: sender];
}

void NSWindow_inst_print(void *id, void* sender) {
	[(NSWindow*)id
		print: sender];
}

void* NSWindow_inst_dataWithEPSInsideRect(void *id, NSRect rect) {
	return [(NSWindow*)id
		dataWithEPSInsideRect: rect];
}

void* NSWindow_inst_dataWithPDFInsideRect(void *id, NSRect rect) {
	return [(NSWindow*)id
		dataWithPDFInsideRect: rect];
}

void NSWindow_inst_updateConstraintsIfNeeded(void *id) {
	[(NSWindow*)id
		updateConstraintsIfNeeded];
}

void NSWindow_inst_layoutIfNeeded(void *id) {
	[(NSWindow*)id
		layoutIfNeeded];
}

void NSWindow_inst_visualizeConstraints(void *id, void* constraints) {
	[(NSWindow*)id
		visualizeConstraints: constraints];
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

BOOL NSWorkspace_inst_openURL(void *id, void* url) {
	return [(NSWorkspace*)id
		openURL: url];
}

void NSWorkspace_inst_hideOtherApplications(void *id) {
	[(NSWorkspace*)id
		hideOtherApplications];
}

void NSWorkspace_inst_activateFileViewerSelectingURLs(void *id, void* fileURLs) {
	[(NSWorkspace*)id
		activateFileViewerSelectingURLs: fileURLs];
}

BOOL NSWorkspace_inst_selectFile_inFileViewerRootedAtPath(void *id, void* fullPath, void* rootFullPath) {
	return [(NSWorkspace*)id
		selectFile: fullPath
		inFileViewerRootedAtPath: rootFullPath];
}

void* NSWorkspace_inst_URLForApplicationWithBundleIdentifier(void *id, void* bundleIdentifier) {
	return [(NSWorkspace*)id
		URLForApplicationWithBundleIdentifier: bundleIdentifier];
}

void* NSWorkspace_inst_URLForApplicationToOpenURL(void *id, void* url) {
	return [(NSWorkspace*)id
		URLForApplicationToOpenURL: url];
}

BOOL NSWorkspace_inst_isFilePackageAtPath(void *id, void* fullPath) {
	return [(NSWorkspace*)id
		isFilePackageAtPath: fullPath];
}

void* NSWorkspace_inst_iconForFile(void *id, void* fullPath) {
	return [(NSWorkspace*)id
		iconForFile: fullPath];
}

void* NSWorkspace_inst_iconForFiles(void *id, void* fullPaths) {
	return [(NSWorkspace*)id
		iconForFiles: fullPaths];
}

BOOL NSWorkspace_inst_unmountAndEjectDeviceAtPath(void *id, void* path) {
	return [(NSWorkspace*)id
		unmountAndEjectDeviceAtPath: path];
}

void* NSWorkspace_inst_desktopImageURLForScreen(void *id, void* screen) {
	return [(NSWorkspace*)id
		desktopImageURLForScreen: screen];
}

void* NSWorkspace_inst_desktopImageOptionsForScreen(void *id, void* screen) {
	return [(NSWorkspace*)id
		desktopImageOptionsForScreen: screen];
}

BOOL NSWorkspace_inst_showSearchResultsForQueryString(void *id, void* queryString) {
	return [(NSWorkspace*)id
		showSearchResultsForQueryString: queryString];
}

void NSWorkspace_inst_noteFileSystemChanged(void *id, void* path) {
	[(NSWorkspace*)id
		noteFileSystemChanged: path];
}

long NSWorkspace_inst_extendPowerOffBy(void *id, long requested) {
	return [(NSWorkspace*)id
		extendPowerOffBy: requested];
}

void* NSWorkspace_inst_URLsForApplicationsToOpenURL(void *id, void* url) {
	return [(NSWorkspace*)id
		URLsForApplicationsToOpenURL: url];
}

void* NSWorkspace_inst_URLsForApplicationsWithBundleIdentifier(void *id, void* bundleIdentifier) {
	return [(NSWorkspace*)id
		URLsForApplicationsWithBundleIdentifier: bundleIdentifier];
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

void* NSColor_inst_highlightWithLevel(void *id, double val) {
	return [(NSColor*)id
		highlightWithLevel: val];
}

void* NSColor_inst_shadowWithLevel(void *id, double val) {
	return [(NSColor*)id
		shadowWithLevel: val];
}

void NSColor_inst_writeToPasteboard(void *id, void* pasteBoard) {
	[(NSColor*)id
		writeToPasteboard: pasteBoard];
}

void NSColor_inst_drawSwatchInRect(void *id, NSRect rect) {
	[(NSColor*)id
		drawSwatchInRect: rect];
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

void* NSTextView_inst_initWithFrame_textContainer(void *id, NSRect frameRect, void* container) {
	return [(NSTextView*)id
		initWithFrame: frameRect
		textContainer: container];
}

void* NSTextView_inst_initWithFrame(void *id, NSRect frameRect) {
	return [(NSTextView*)id
		initWithFrame: frameRect];
}

void NSTextView_inst_replaceTextContainer(void *id, void* newContainer) {
	[(NSTextView*)id
		replaceTextContainer: newContainer];
}

void NSTextView_inst_invalidateTextContainerOrigin(void *id) {
	[(NSTextView*)id
		invalidateTextContainerOrigin];
}

void NSTextView_inst_changeDocumentBackgroundColor(void *id, void* sender) {
	[(NSTextView*)id
		changeDocumentBackgroundColor: sender];
}

void NSTextView_inst_setNeedsDisplayInRect_avoidAdditionalLayout(void *id, NSRect rect, BOOL flag) {
	[(NSTextView*)id
		setNeedsDisplayInRect: rect
		avoidAdditionalLayout: flag];
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

void NSTextView_inst_setConstrainedFrameSize(void *id, NSSize desiredSize) {
	[(NSTextView*)id
		setConstrainedFrameSize: desiredSize];
}

void NSTextView_inst_cleanUpAfterDragOperation(void *id) {
	[(NSTextView*)id
		cleanUpAfterDragOperation];
}

void NSTextView_inst_outline(void *id, void* sender) {
	[(NSTextView*)id
		outline: sender];
}

void NSTextView_inst_toggleAutomaticQuoteSubstitution(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticQuoteSubstitution: sender];
}

void NSTextView_inst_toggleAutomaticLinkDetection(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticLinkDetection: sender];
}

void NSTextView_inst_toggleAutomaticTextCompletion(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticTextCompletion: sender];
}

void NSTextView_inst_updateInsertionPointStateAndRestartTimer(void *id, BOOL restartFlag) {
	[(NSTextView*)id
		updateInsertionPointStateAndRestartTimer: restartFlag];
}

unsigned long NSTextView_inst_characterIndexForInsertionAtPoint(void *id, NSPoint point) {
	return [(NSTextView*)id
		characterIndexForInsertionAtPoint: point];
}

void NSTextView_inst_updateCandidates(void *id) {
	[(NSTextView*)id
		updateCandidates];
}

BOOL NSTextView_inst_readSelectionFromPasteboard(void *id, void* pboard) {
	return [(NSTextView*)id
		readSelectionFromPasteboard: pboard];
}

BOOL NSTextView_inst_writeSelectionToPasteboard_types(void *id, void* pboard, void* types) {
	return [(NSTextView*)id
		writeSelectionToPasteboard: pboard
		types: types];
}

void NSTextView_inst_alignJustified(void *id, void* sender) {
	[(NSTextView*)id
		alignJustified: sender];
}

void NSTextView_inst_changeAttributes(void *id, void* sender) {
	[(NSTextView*)id
		changeAttributes: sender];
}

void NSTextView_inst_changeColor(void *id, void* sender) {
	[(NSTextView*)id
		changeColor: sender];
}

void NSTextView_inst_useStandardKerning(void *id, void* sender) {
	[(NSTextView*)id
		useStandardKerning: sender];
}

void NSTextView_inst_lowerBaseline(void *id, void* sender) {
	[(NSTextView*)id
		lowerBaseline: sender];
}

void NSTextView_inst_raiseBaseline(void *id, void* sender) {
	[(NSTextView*)id
		raiseBaseline: sender];
}

void NSTextView_inst_turnOffKerning(void *id, void* sender) {
	[(NSTextView*)id
		turnOffKerning: sender];
}

void NSTextView_inst_loosenKerning(void *id, void* sender) {
	[(NSTextView*)id
		loosenKerning: sender];
}

void NSTextView_inst_tightenKerning(void *id, void* sender) {
	[(NSTextView*)id
		tightenKerning: sender];
}

void NSTextView_inst_useStandardLigatures(void *id, void* sender) {
	[(NSTextView*)id
		useStandardLigatures: sender];
}

void NSTextView_inst_turnOffLigatures(void *id, void* sender) {
	[(NSTextView*)id
		turnOffLigatures: sender];
}

void NSTextView_inst_useAllLigatures(void *id, void* sender) {
	[(NSTextView*)id
		useAllLigatures: sender];
}

void NSTextView_inst_clickedOnLink_atIndex(void *id, void* link, unsigned long charIndex) {
	[(NSTextView*)id
		clickedOnLink: link
		atIndex: charIndex];
}

void NSTextView_inst_pasteAsPlainText(void *id, void* sender) {
	[(NSTextView*)id
		pasteAsPlainText: sender];
}

void NSTextView_inst_pasteAsRichText(void *id, void* sender) {
	[(NSTextView*)id
		pasteAsRichText: sender];
}

void NSTextView_inst_breakUndoCoalescing(void *id) {
	[(NSTextView*)id
		breakUndoCoalescing];
}

void NSTextView_inst_updateFontPanel(void *id) {
	[(NSTextView*)id
		updateFontPanel];
}

void NSTextView_inst_updateRuler(void *id) {
	[(NSTextView*)id
		updateRuler];
}

void NSTextView_inst_updateDragTypeRegistration(void *id) {
	[(NSTextView*)id
		updateDragTypeRegistration];
}

BOOL NSTextView_inst_shouldChangeTextInRanges_replacementStrings(void *id, void* affectedRanges, void* replacementStrings) {
	return [(NSTextView*)id
		shouldChangeTextInRanges: affectedRanges
		replacementStrings: replacementStrings];
}

void NSTextView_inst_didChangeText(void *id) {
	[(NSTextView*)id
		didChangeText];
}

void NSTextView_inst_toggleSmartInsertDelete(void *id, void* sender) {
	[(NSTextView*)id
		toggleSmartInsertDelete: sender];
}

void NSTextView_inst_toggleContinuousSpellChecking(void *id, void* sender) {
	[(NSTextView*)id
		toggleContinuousSpellChecking: sender];
}

void NSTextView_inst_toggleGrammarChecking(void *id, void* sender) {
	[(NSTextView*)id
		toggleGrammarChecking: sender];
}

void NSTextView_inst_orderFrontSharingServicePicker(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontSharingServicePicker: sender];
}

BOOL NSTextView_inst_dragSelectionWithEvent_offset_slideBack(void *id, void* event, NSSize mouseOffset, BOOL slideBack) {
	return [(NSTextView*)id
		dragSelectionWithEvent: event
		offset: mouseOffset
		slideBack: slideBack];
}

void NSTextView_inst_startSpeaking(void *id, void* sender) {
	[(NSTextView*)id
		startSpeaking: sender];
}

void NSTextView_inst_stopSpeaking(void *id, void* sender) {
	[(NSTextView*)id
		stopSpeaking: sender];
}

void NSTextView_inst_performFindPanelAction(void *id, void* sender) {
	[(NSTextView*)id
		performFindPanelAction: sender];
}

void NSTextView_inst_orderFrontLinkPanel(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontLinkPanel: sender];
}

void NSTextView_inst_orderFrontListPanel(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontListPanel: sender];
}

void NSTextView_inst_orderFrontSpacingPanel(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontSpacingPanel: sender];
}

void NSTextView_inst_orderFrontTablePanel(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontTablePanel: sender];
}

void NSTextView_inst_orderFrontSubstitutionsPanel(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontSubstitutionsPanel: sender];
}

void NSTextView_inst_complete(void *id, void* sender) {
	[(NSTextView*)id
		complete: sender];
}

void NSTextView_inst_checkTextInDocument(void *id, void* sender) {
	[(NSTextView*)id
		checkTextInDocument: sender];
}

void NSTextView_inst_checkTextInSelection(void *id, void* sender) {
	[(NSTextView*)id
		checkTextInSelection: sender];
}

void NSTextView_inst_toggleAutomaticDashSubstitution(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticDashSubstitution: sender];
}

void NSTextView_inst_toggleAutomaticDataDetection(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticDataDetection: sender];
}

void NSTextView_inst_toggleAutomaticSpellingCorrection(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticSpellingCorrection: sender];
}

void NSTextView_inst_toggleAutomaticTextReplacement(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticTextReplacement: sender];
}

void NSTextView_inst_updateQuickLookPreviewPanel(void *id) {
	[(NSTextView*)id
		updateQuickLookPreviewPanel];
}

void NSTextView_inst_toggleQuickLookPreviewPanel(void *id, void* sender) {
	[(NSTextView*)id
		toggleQuickLookPreviewPanel: sender];
}

void* NSTextView_inst_quickLookPreviewableItemsInRanges(void *id, void* ranges) {
	return [(NSTextView*)id
		quickLookPreviewableItemsInRanges: ranges];
}

void NSTextView_inst_changeLayoutOrientation(void *id, void* sender) {
	[(NSTextView*)id
		changeLayoutOrientation: sender];
}

void NSTextView_inst_updateTextTouchBarItems(void *id) {
	[(NSTextView*)id
		updateTextTouchBarItems];
}

void NSTextView_inst_updateTouchBarItemIdentifiers(void *id) {
	[(NSTextView*)id
		updateTouchBarItemIdentifiers];
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

void* NSView_inst_initWithFrame(void *id, NSRect frameRect) {
	return [(NSView*)id
		initWithFrame: frameRect];
}

void NSView_inst_prepareForReuse(void *id) {
	[(NSView*)id
		prepareForReuse];
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

void NSView_inst_didAddSubview(void *id, void* subview) {
	[(NSView*)id
		didAddSubview: subview];
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

BOOL NSView_inst_isDescendantOf(void *id, void* view) {
	return [(NSView*)id
		isDescendantOf: view];
}

void* NSView_inst_ancestorSharedWithView(void *id, void* view) {
	return [(NSView*)id
		ancestorSharedWithView: view];
}

void NSView_inst_viewDidMoveToSuperview(void *id) {
	[(NSView*)id
		viewDidMoveToSuperview];
}

void NSView_inst_viewDidMoveToWindow(void *id) {
	[(NSView*)id
		viewDidMoveToWindow];
}

void NSView_inst_viewWillMoveToSuperview(void *id, void* newSuperview) {
	[(NSView*)id
		viewWillMoveToSuperview: newSuperview];
}

void NSView_inst_viewWillMoveToWindow(void *id, void* newWindow) {
	[(NSView*)id
		viewWillMoveToWindow: newWindow];
}

void NSView_inst_willRemoveSubview(void *id, void* subview) {
	[(NSView*)id
		willRemoveSubview: subview];
}

void NSView_inst_setFrameOrigin(void *id, NSPoint newOrigin) {
	[(NSView*)id
		setFrameOrigin: newOrigin];
}

void NSView_inst_setFrameSize(void *id, NSSize newSize) {
	[(NSView*)id
		setFrameSize: newSize];
}

void NSView_inst_setBoundsOrigin(void *id, NSPoint newOrigin) {
	[(NSView*)id
		setBoundsOrigin: newOrigin];
}

void NSView_inst_setBoundsSize(void *id, NSSize newSize) {
	[(NSView*)id
		setBoundsSize: newSize];
}

void* NSView_inst_makeBackingLayer(void *id) {
	return [(NSView*)id
		makeBackingLayer];
}

void NSView_inst_updateLayer(void *id) {
	[(NSView*)id
		updateLayer];
}

void NSView_inst_drawRect(void *id, NSRect dirtyRect) {
	[(NSView*)id
		drawRect: dirtyRect];
}

BOOL NSView_inst_needsToDrawRect(void *id, NSRect rect) {
	return [(NSView*)id
		needsToDrawRect: rect];
}

void NSView_inst_print(void *id, void* sender) {
	[(NSView*)id
		print: sender];
}

void NSView_inst_beginPageInRect_atPlacement(void *id, NSRect rect, NSPoint location) {
	[(NSView*)id
		beginPageInRect: rect
		atPlacement: location];
}

void* NSView_inst_dataWithEPSInsideRect(void *id, NSRect rect) {
	return [(NSView*)id
		dataWithEPSInsideRect: rect];
}

void* NSView_inst_dataWithPDFInsideRect(void *id, NSRect rect) {
	return [(NSView*)id
		dataWithPDFInsideRect: rect];
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

void NSView_inst_drawPageBorderWithSize(void *id, NSSize borderSize) {
	[(NSView*)id
		drawPageBorderWithSize: borderSize];
}

NSRect NSView_inst_rectForPage(void *id, long page) {
	return [(NSView*)id
		rectForPage: page];
}

NSPoint NSView_inst_locationOfPrintRect(void *id, NSRect rect) {
	return [(NSView*)id
		locationOfPrintRect: rect];
}

void NSView_inst_setNeedsDisplayInRect(void *id, NSRect invalidRect) {
	[(NSView*)id
		setNeedsDisplayInRect: invalidRect];
}

void NSView_inst_display(void *id) {
	[(NSView*)id
		display];
}

void NSView_inst_displayRect(void *id, NSRect rect) {
	[(NSView*)id
		displayRect: rect];
}

void NSView_inst_displayRectIgnoringOpacity(void *id, NSRect rect) {
	[(NSView*)id
		displayRectIgnoringOpacity: rect];
}

void NSView_inst_displayIfNeeded(void *id) {
	[(NSView*)id
		displayIfNeeded];
}

void NSView_inst_displayIfNeededInRect(void *id, NSRect rect) {
	[(NSView*)id
		displayIfNeededInRect: rect];
}

void NSView_inst_displayIfNeededIgnoringOpacity(void *id) {
	[(NSView*)id
		displayIfNeededIgnoringOpacity];
}

void NSView_inst_displayIfNeededInRectIgnoringOpacity(void *id, NSRect rect) {
	[(NSView*)id
		displayIfNeededInRectIgnoringOpacity: rect];
}

void NSView_inst_translateRectsNeedingDisplayInRect_by(void *id, NSRect clipRect, NSSize delta) {
	[(NSView*)id
		translateRectsNeedingDisplayInRect: clipRect
		by: delta];
}

void NSView_inst_viewWillDraw(void *id) {
	[(NSView*)id
		viewWillDraw];
}

NSPoint NSView_inst_convertPointFromBacking(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointFromBacking: point];
}

NSPoint NSView_inst_convertPointToBacking(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointToBacking: point];
}

NSPoint NSView_inst_convertPointFromLayer(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointFromLayer: point];
}

NSPoint NSView_inst_convertPointToLayer(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointToLayer: point];
}

NSRect NSView_inst_convertRectFromBacking(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectFromBacking: rect];
}

NSRect NSView_inst_convertRectToBacking(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectToBacking: rect];
}

NSRect NSView_inst_convertRectFromLayer(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectFromLayer: rect];
}

NSRect NSView_inst_convertRectToLayer(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectToLayer: rect];
}

NSSize NSView_inst_convertSizeFromBacking(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeFromBacking: size];
}

NSSize NSView_inst_convertSizeToBacking(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeToBacking: size];
}

NSSize NSView_inst_convertSizeFromLayer(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeFromLayer: size];
}

NSSize NSView_inst_convertSizeToLayer(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeToLayer: size];
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

NSRect NSView_inst_centerScanRect(void *id, NSRect rect) {
	return [(NSView*)id
		centerScanRect: rect];
}

void NSView_inst_translateOriginToPoint(void *id, NSPoint translation) {
	[(NSView*)id
		translateOriginToPoint: translation];
}

void NSView_inst_scaleUnitSquareToSize(void *id, NSSize newUnitSize) {
	[(NSView*)id
		scaleUnitSquareToSize: newUnitSize];
}

void NSView_inst_rotateByAngle(void *id, double angle) {
	[(NSView*)id
		rotateByAngle: angle];
}

void NSView_inst_resizeSubviewsWithOldSize(void *id, NSSize oldSize) {
	[(NSView*)id
		resizeSubviewsWithOldSize: oldSize];
}

void NSView_inst_resizeWithOldSuperviewSize(void *id, NSSize oldSize) {
	[(NSView*)id
		resizeWithOldSuperviewSize: oldSize];
}

void NSView_inst_addConstraints(void *id, void* constraints) {
	[(NSView*)id
		addConstraints: constraints];
}

void NSView_inst_removeConstraints(void *id, void* constraints) {
	[(NSView*)id
		removeConstraints: constraints];
}

void NSView_inst_invalidateIntrinsicContentSize(void *id) {
	[(NSView*)id
		invalidateIntrinsicContentSize];
}

NSRect NSView_inst_alignmentRectForFrame(void *id, NSRect frame) {
	return [(NSView*)id
		alignmentRectForFrame: frame];
}

NSRect NSView_inst_frameForAlignmentRect(void *id, NSRect alignmentRect) {
	return [(NSView*)id
		frameForAlignmentRect: alignmentRect];
}

void NSView_inst_layout(void *id) {
	[(NSView*)id
		layout];
}

void NSView_inst_layoutSubtreeIfNeeded(void *id) {
	[(NSView*)id
		layoutSubtreeIfNeeded];
}

void NSView_inst_updateConstraints(void *id) {
	[(NSView*)id
		updateConstraints];
}

void NSView_inst_updateConstraintsForSubtreeIfNeeded(void *id) {
	[(NSView*)id
		updateConstraintsForSubtreeIfNeeded];
}

void NSView_inst_exerciseAmbiguityInLayout(void *id) {
	[(NSView*)id
		exerciseAmbiguityInLayout];
}

void NSView_inst_drawFocusRingMask(void *id) {
	[(NSView*)id
		drawFocusRingMask];
}

void NSView_inst_noteFocusRingMaskChanged(void *id) {
	[(NSView*)id
		noteFocusRingMaskChanged];
}

void NSView_inst_setKeyboardFocusRingNeedsDisplayInRect(void *id, NSRect rect) {
	[(NSView*)id
		setKeyboardFocusRingNeedsDisplayInRect: rect];
}

BOOL NSView_inst_enterFullScreenMode_withOptions(void *id, void* screen, void* options) {
	return [(NSView*)id
		enterFullScreenMode: screen
		withOptions: options];
}

void NSView_inst_exitFullScreenModeWithOptions(void *id, void* options) {
	[(NSView*)id
		exitFullScreenModeWithOptions: options];
}

void NSView_inst_viewDidHide(void *id) {
	[(NSView*)id
		viewDidHide];
}

void NSView_inst_viewDidUnhide(void *id) {
	[(NSView*)id
		viewDidUnhide];
}

void NSView_inst_viewWillStartLiveResize(void *id) {
	[(NSView*)id
		viewWillStartLiveResize];
}

void NSView_inst_viewDidEndLiveResize(void *id) {
	[(NSView*)id
		viewDidEndLiveResize];
}

BOOL NSView_inst_acceptsFirstMouse(void *id, void* event) {
	return [(NSView*)id
		acceptsFirstMouse: event];
}

void* NSView_inst_hitTest(void *id, NSPoint point) {
	return [(NSView*)id
		hitTest: point];
}

BOOL NSView_inst_mouse_inRect(void *id, NSPoint point, NSRect rect) {
	return [(NSView*)id
		mouse: point
		inRect: rect];
}

BOOL NSView_inst_performKeyEquivalent(void *id, void* event) {
	return [(NSView*)id
		performKeyEquivalent: event];
}

void NSView_inst_prepareContentInRect(void *id, NSRect rect) {
	[(NSView*)id
		prepareContentInRect: rect];
}

void NSView_inst_scrollPoint(void *id, NSPoint point) {
	[(NSView*)id
		scrollPoint: point];
}

BOOL NSView_inst_scrollRectToVisible(void *id, NSRect rect) {
	return [(NSView*)id
		scrollRectToVisible: rect];
}

BOOL NSView_inst_autoscroll(void *id, void* event) {
	return [(NSView*)id
		autoscroll: event];
}

NSRect NSView_inst_adjustScroll(void *id, NSRect newVisible) {
	return [(NSView*)id
		adjustScroll: newVisible];
}

void NSView_inst_registerForDraggedTypes(void *id, void* newTypes) {
	[(NSView*)id
		registerForDraggedTypes: newTypes];
}

void NSView_inst_unregisterDraggedTypes(void *id) {
	[(NSView*)id
		unregisterDraggedTypes];
}

BOOL NSView_inst_shouldDelayWindowOrderingForEvent(void *id, void* event) {
	return [(NSView*)id
		shouldDelayWindowOrderingForEvent: event];
}

NSRect NSView_inst_rectForSmartMagnificationAtPoint_inRect(void *id, NSPoint location, NSRect visibleRect) {
	return [(NSView*)id
		rectForSmartMagnificationAtPoint: location
		inRect: visibleRect];
}

void NSView_inst_viewDidChangeBackingProperties(void *id) {
	[(NSView*)id
		viewDidChangeBackingProperties];
}

void* NSView_inst_viewWithTag(void *id, long tag) {
	return [(NSView*)id
		viewWithTag: tag];
}

void NSView_inst_removeAllToolTips(void *id) {
	[(NSView*)id
		removeAllToolTips];
}

void NSView_inst_updateTrackingAreas(void *id) {
	[(NSView*)id
		updateTrackingAreas];
}

void NSView_inst_discardCursorRects(void *id) {
	[(NSView*)id
		discardCursorRects];
}

void NSView_inst_resetCursorRects(void *id) {
	[(NSView*)id
		resetCursorRects];
}

void* NSView_inst_menuForEvent(void *id, void* event) {
	return [(NSView*)id
		menuForEvent: event];
}

void NSView_inst_willOpenMenu_withEvent(void *id, void* menu, void* event) {
	[(NSView*)id
		willOpenMenu: menu
		withEvent: event];
}

void NSView_inst_didCloseMenu_withEvent(void *id, void* menu, void* event) {
	[(NSView*)id
		didCloseMenu: menu
		withEvent: event];
}

void NSView_inst_beginDocument(void *id) {
	[(NSView*)id
		beginDocument];
}

void NSView_inst_endDocument(void *id) {
	[(NSView*)id
		endDocument];
}

void NSView_inst_endPage(void *id) {
	[(NSView*)id
		endPage];
}

void NSView_inst_showDefinitionForAttributedString_atPoint(void *id, void* attrString, NSPoint textBaselineOrigin) {
	[(NSView*)id
		showDefinitionForAttributedString: attrString
		atPoint: textBaselineOrigin];
}

void NSView_inst_viewDidChangeEffectiveAppearance(void *id) {
	[(NSView*)id
		viewDidChangeEffectiveAppearance];
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

func NSBundle_alloc() (
	r0 NSBundle,
) {
	ret := C.NSBundle_type_alloc()
	r0 = NSBundle_fromPointer(ret)
	return
}

func NSBundle_bundleWithURL(
	url core.NSURLRef,
) (
	r0 NSBundle,
) {
	ret := C.NSBundle_type_bundleWithURL(
		objc.RefPointer(url),
	)
	r0 = NSBundle_fromPointer(ret)
	return
}

func NSBundle_bundleWithPath(
	path core.NSStringRef,
) (
	r0 NSBundle,
) {
	ret := C.NSBundle_type_bundleWithPath(
		objc.RefPointer(path),
	)
	r0 = NSBundle_fromPointer(ret)
	return
}

func NSBundle_bundleWithIdentifier(
	identifier core.NSStringRef,
) (
	r0 NSBundle,
) {
	ret := C.NSBundle_type_bundleWithIdentifier(
		objc.RefPointer(identifier),
	)
	r0 = NSBundle_fromPointer(ret)
	return
}

func NSBundle_URLForResource_withExtension_subdirectory_inBundleWithURL(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
	bundleURL core.NSURLRef,
) (
	r0 core.NSURL,
) {
	ret := C.NSBundle_type_URLForResource_withExtension_subdirectory_inBundleWithURL(
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(bundleURL),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func NSBundle_URLsForResourcesWithExtension_subdirectory_inBundleWithURL(
	ext core.NSStringRef,
	subpath core.NSStringRef,
	bundleURL core.NSURLRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_type_URLsForResourcesWithExtension_subdirectory_inBundleWithURL(
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(bundleURL),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func NSBundle_pathForResource_ofType_inDirectory(
	name core.NSStringRef,
	ext core.NSStringRef,
	bundlePath core.NSStringRef,
) (
	r0 core.NSString,
) {
	ret := C.NSBundle_type_pathForResource_ofType_inDirectory(
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(bundlePath),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func NSBundle_pathsForResourcesOfType_inDirectory(
	ext core.NSStringRef,
	bundlePath core.NSStringRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_type_pathsForResourcesOfType_inDirectory(
		objc.RefPointer(ext),
		objc.RefPointer(bundlePath),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func NSBundle_preferredLocalizationsFromArray(
	localizationsArray core.NSArrayRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_type_preferredLocalizationsFromArray(
		objc.RefPointer(localizationsArray),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func NSBundle_preferredLocalizationsFromArray_forPreferences(
	localizationsArray core.NSArrayRef,
	preferencesArray core.NSArrayRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_type_preferredLocalizationsFromArray_forPreferences(
		objc.RefPointer(localizationsArray),
		objc.RefPointer(preferencesArray),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func NSBundle_mainBundle() (
	r0 NSBundle,
) {
	ret := C.NSBundle_type_mainBundle()
	r0 = NSBundle_fromPointer(ret)
	return
}

func NSBundle_allFrameworks() (
	r0 core.NSArray,
) {
	ret := C.NSBundle_type_allFrameworks()
	r0 = core.NSArray_fromPointer(ret)
	return
}

func NSBundle_allBundles() (
	r0 core.NSArray,
) {
	ret := C.NSBundle_type_allBundles()
	r0 = core.NSArray_fromPointer(ret)
	return
}

func NSSound_alloc() (
	r0 NSSound,
) {
	ret := C.NSSound_type_alloc()
	r0 = NSSound_fromPointer(ret)
	return
}

func NSSound_canInitWithPasteboard(
	pasteboard NSPasteboardRef,
) (
	r0 bool,
) {
	ret := C.NSSound_type_canInitWithPasteboard(
		objc.RefPointer(pasteboard),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func NSSound_soundUnfilteredTypes() (
	r0 core.NSArray,
) {
	ret := C.NSSound_type_soundUnfilteredTypes()
	r0 = core.NSArray_fromPointer(ret)
	return
}

func NSApplication_alloc() (
	r0 NSApplication,
) {
	ret := C.NSApplication_type_alloc()
	r0 = NSApplication_fromPointer(ret)
	return
}

func NSApplication_detachDrawingThread_toTarget_withObject(
	selector objc.Selector,
	target objc.Ref,
	argument objc.Ref,
) {
	C.NSApplication_type_detachDrawingThread_toTarget_withObject(
		selector.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(argument),
	)
	return
}

func NSApplication_sharedApplication() (
	r0 NSApplication,
) {
	ret := C.NSApplication_type_sharedApplication()
	r0 = NSApplication_fromPointer(ret)
	return
}

func NSControl_alloc() (
	r0 NSControl,
) {
	ret := C.NSControl_type_alloc()
	r0 = NSControl_fromPointer(ret)
	return
}

func NSButton_alloc() (
	r0 NSButton,
) {
	ret := C.NSButton_type_alloc()
	r0 = NSButton_fromPointer(ret)
	return
}

func NSButton_checkboxWithTitle_target_action(
	title core.NSStringRef,
	target objc.Ref,
	action objc.Selector,
) (
	r0 NSButton,
) {
	ret := C.NSButton_type_checkboxWithTitle_target_action(
		objc.RefPointer(title),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)
	r0 = NSButton_fromPointer(ret)
	return
}

func NSButton_buttonWithImage_target_action(
	image NSImageRef,
	target objc.Ref,
	action objc.Selector,
) (
	r0 NSButton,
) {
	ret := C.NSButton_type_buttonWithImage_target_action(
		objc.RefPointer(image),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)
	r0 = NSButton_fromPointer(ret)
	return
}

func NSButton_radioButtonWithTitle_target_action(
	title core.NSStringRef,
	target objc.Ref,
	action objc.Selector,
) (
	r0 NSButton,
) {
	ret := C.NSButton_type_radioButtonWithTitle_target_action(
		objc.RefPointer(title),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)
	r0 = NSButton_fromPointer(ret)
	return
}

func NSButton_buttonWithTitle_image_target_action(
	title core.NSStringRef,
	image NSImageRef,
	target objc.Ref,
	action objc.Selector,
) (
	r0 NSButton,
) {
	ret := C.NSButton_type_buttonWithTitle_image_target_action(
		objc.RefPointer(title),
		objc.RefPointer(image),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)
	r0 = NSButton_fromPointer(ret)
	return
}

func NSButton_buttonWithTitle_target_action(
	title core.NSStringRef,
	target objc.Ref,
	action objc.Selector,
) (
	r0 NSButton,
) {
	ret := C.NSButton_type_buttonWithTitle_target_action(
		objc.RefPointer(title),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)
	r0 = NSButton_fromPointer(ret)
	return
}

func NSEvent_alloc() (
	r0 NSEvent,
) {
	ret := C.NSEvent_type_alloc()
	r0 = NSEvent_fromPointer(ret)
	return
}

func NSEvent_eventWithEventRef(
	eventRef unsafe.Pointer,
) (
	r0 NSEvent,
) {
	ret := C.NSEvent_type_eventWithEventRef(
		eventRef,
	)
	r0 = NSEvent_fromPointer(ret)
	return
}

func NSEvent_stopPeriodicEvents() {
	C.NSEvent_type_stopPeriodicEvents()
	return
}

func NSEvent_removeMonitor(
	eventMonitor objc.Ref,
) {
	C.NSEvent_type_removeMonitor(
		objc.RefPointer(eventMonitor),
	)
	return
}

func NSEvent_pressedMouseButtons() (
	r0 core.NSUInteger,
) {
	ret := C.NSEvent_type_pressedMouseButtons()
	r0 = core.NSUInteger(ret)
	return
}

func NSEvent_mouseLocation() (
	r0 core.NSPoint,
) {
	ret := C.NSEvent_type_mouseLocation()
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func NSEvent_mouseCoalescingEnabled() (
	r0 bool,
) {
	ret := C.NSEvent_type_mouseCoalescingEnabled()
	r0 = convertObjCBoolToGo(ret)
	return
}

func NSEvent_setMouseCoalescingEnabled(
	value bool,
) {
	C.NSEvent_type_setMouseCoalescingEnabled(
		convertToObjCBool(value),
	)
	return
}

func NSEvent_swipeTrackingFromScrollEventsEnabled() (
	r0 bool,
) {
	ret := C.NSEvent_type_swipeTrackingFromScrollEventsEnabled()
	r0 = convertObjCBoolToGo(ret)
	return
}

func NSFont_alloc() (
	r0 NSFont,
) {
	ret := C.NSFont_type_alloc()
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_fontWithName_size(
	fontName core.NSStringRef,
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_fontWithName_size(
		objc.RefPointer(fontName),
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_userFontOfSize(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_userFontOfSize(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_userFixedPitchFontOfSize(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_userFixedPitchFontOfSize(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_systemFontOfSize(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_systemFontOfSize(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_boldSystemFontOfSize(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_boldSystemFontOfSize(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_labelFontOfSize(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_labelFontOfSize(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_messageFontOfSize(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_messageFontOfSize(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_menuBarFontOfSize(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_menuBarFontOfSize(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_menuFontOfSize(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_menuFontOfSize(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_controlContentFontOfSize(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_controlContentFontOfSize(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_titleBarFontOfSize(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_titleBarFontOfSize(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_paletteFontOfSize(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_paletteFontOfSize(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_toolTipsFontOfSize(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_toolTipsFontOfSize(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_setUserFont(
	font NSFontRef,
) {
	C.NSFont_type_setUserFont(
		objc.RefPointer(font),
	)
	return
}

func NSFont_setUserFixedPitchFont(
	font NSFontRef,
) {
	C.NSFont_type_setUserFixedPitchFont(
		objc.RefPointer(font),
	)
	return
}

func NSFont_systemFontSize() (
	r0 core.CGFloat,
) {
	ret := C.NSFont_type_systemFontSize()
	r0 = core.CGFloat(ret)
	return
}

func NSFont_smallSystemFontSize() (
	r0 core.CGFloat,
) {
	ret := C.NSFont_type_smallSystemFontSize()
	r0 = core.CGFloat(ret)
	return
}

func NSFont_labelFontSize() (
	r0 core.CGFloat,
) {
	ret := C.NSFont_type_labelFontSize()
	r0 = core.CGFloat(ret)
	return
}

func NSImage_alloc() (
	r0 NSImage,
) {
	ret := C.NSImage_type_alloc()
	r0 = NSImage_fromPointer(ret)
	return
}

func NSImage_imageWithSystemSymbolName_accessibilityDescription(
	symbolName core.NSStringRef,
	description core.NSStringRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_type_imageWithSystemSymbolName_accessibilityDescription(
		objc.RefPointer(symbolName),
		objc.RefPointer(description),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func NSImage_canInitWithPasteboard(
	pasteboard NSPasteboardRef,
) (
	r0 bool,
) {
	ret := C.NSImage_type_canInitWithPasteboard(
		objc.RefPointer(pasteboard),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func NSImage_imageTypes() (
	r0 core.NSArray,
) {
	ret := C.NSImage_type_imageTypes()
	r0 = core.NSArray_fromPointer(ret)
	return
}

func NSImage_imageUnfilteredTypes() (
	r0 core.NSArray,
) {
	ret := C.NSImage_type_imageUnfilteredTypes()
	r0 = core.NSArray_fromPointer(ret)
	return
}

func NSImageView_alloc() (
	r0 NSImageView,
) {
	ret := C.NSImageView_type_alloc()
	r0 = NSImageView_fromPointer(ret)
	return
}

func NSImageView_imageViewWithImage(
	image NSImageRef,
) (
	r0 NSImageView,
) {
	ret := C.NSImageView_type_imageViewWithImage(
		objc.RefPointer(image),
	)
	r0 = NSImageView_fromPointer(ret)
	return
}

func NSNib_alloc() (
	r0 NSNib,
) {
	ret := C.NSNib_type_alloc()
	r0 = NSNib_fromPointer(ret)
	return
}

func NSPasteboard_alloc() (
	r0 NSPasteboard,
) {
	ret := C.NSPasteboard_type_alloc()
	r0 = NSPasteboard_fromPointer(ret)
	return
}

func NSPasteboard_pasteboardByFilteringFile(
	filename core.NSStringRef,
) (
	r0 NSPasteboard,
) {
	ret := C.NSPasteboard_type_pasteboardByFilteringFile(
		objc.RefPointer(filename),
	)
	r0 = NSPasteboard_fromPointer(ret)
	return
}

func NSPasteboard_pasteboardByFilteringTypesInPasteboard(
	pboard NSPasteboardRef,
) (
	r0 NSPasteboard,
) {
	ret := C.NSPasteboard_type_pasteboardByFilteringTypesInPasteboard(
		objc.RefPointer(pboard),
	)
	r0 = NSPasteboard_fromPointer(ret)
	return
}

func NSPasteboard_pasteboardWithUniqueName() (
	r0 NSPasteboard,
) {
	ret := C.NSPasteboard_type_pasteboardWithUniqueName()
	r0 = NSPasteboard_fromPointer(ret)
	return
}

func NSPasteboard_generalPasteboard() (
	r0 NSPasteboard,
) {
	ret := C.NSPasteboard_type_generalPasteboard()
	r0 = NSPasteboard_fromPointer(ret)
	return
}

func NSLayoutManager_alloc() (
	r0 NSLayoutManager,
) {
	ret := C.NSLayoutManager_type_alloc()
	r0 = NSLayoutManager_fromPointer(ret)
	return
}

func NSMenu_alloc() (
	r0 NSMenu,
) {
	ret := C.NSMenu_type_alloc()
	r0 = NSMenu_fromPointer(ret)
	return
}

func NSMenu_menuBarVisible() (
	r0 bool,
) {
	ret := C.NSMenu_type_menuBarVisible()
	r0 = convertObjCBoolToGo(ret)
	return
}

func NSMenu_setMenuBarVisible(
	visible bool,
) {
	C.NSMenu_type_setMenuBarVisible(
		convertToObjCBool(visible),
	)
	return
}

func NSMenu_popUpContextMenu_withEvent_forView(
	menu NSMenuRef,
	event NSEventRef,
	view NSViewRef,
) {
	C.NSMenu_type_popUpContextMenu_withEvent_forView(
		objc.RefPointer(menu),
		objc.RefPointer(event),
		objc.RefPointer(view),
	)
	return
}

func NSMenu_popUpContextMenu_withEvent_forView_withFont(
	menu NSMenuRef,
	event NSEventRef,
	view NSViewRef,
	font NSFontRef,
) {
	C.NSMenu_type_popUpContextMenu_withEvent_forView_withFont(
		objc.RefPointer(menu),
		objc.RefPointer(event),
		objc.RefPointer(view),
		objc.RefPointer(font),
	)
	return
}

func NSPopover_alloc() (
	r0 NSPopover,
) {
	ret := C.NSPopover_type_alloc()
	r0 = NSPopover_fromPointer(ret)
	return
}

func NSMenuItem_alloc() (
	r0 NSMenuItem,
) {
	ret := C.NSMenuItem_type_alloc()
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func NSMenuItem_separatorItem() (
	r0 NSMenuItem,
) {
	ret := C.NSMenuItem_type_separatorItem()
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func NSMenuItem_usesUserKeyEquivalents() (
	r0 bool,
) {
	ret := C.NSMenuItem_type_usesUserKeyEquivalents()
	r0 = convertObjCBoolToGo(ret)
	return
}

func NSMenuItem_setUsesUserKeyEquivalents(
	value bool,
) {
	C.NSMenuItem_type_setUsesUserKeyEquivalents(
		convertToObjCBool(value),
	)
	return
}

func NSRunningApplication_alloc() (
	r0 NSRunningApplication,
) {
	ret := C.NSRunningApplication_type_alloc()
	r0 = NSRunningApplication_fromPointer(ret)
	return
}

func NSRunningApplication_runningApplicationsWithBundleIdentifier(
	bundleIdentifier core.NSStringRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSRunningApplication_type_runningApplicationsWithBundleIdentifier(
		objc.RefPointer(bundleIdentifier),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func NSRunningApplication_terminateAutomaticallyTerminableApplications() {
	C.NSRunningApplication_type_terminateAutomaticallyTerminableApplications()
	return
}

func NSRunningApplication_currentApplication() (
	r0 NSRunningApplication,
) {
	ret := C.NSRunningApplication_type_currentApplication()
	r0 = NSRunningApplication_fromPointer(ret)
	return
}

func NSScreen_alloc() (
	r0 NSScreen,
) {
	ret := C.NSScreen_type_alloc()
	r0 = NSScreen_fromPointer(ret)
	return
}

func NSScreen_mainScreen() (
	r0 NSScreen,
) {
	ret := C.NSScreen_type_mainScreen()
	r0 = NSScreen_fromPointer(ret)
	return
}

func NSScreen_deepestScreen() (
	r0 NSScreen,
) {
	ret := C.NSScreen_type_deepestScreen()
	r0 = NSScreen_fromPointer(ret)
	return
}

func NSScreen_screens() (
	r0 core.NSArray,
) {
	ret := C.NSScreen_type_screens()
	r0 = core.NSArray_fromPointer(ret)
	return
}

func NSScreen_screensHaveSeparateSpaces() (
	r0 bool,
) {
	ret := C.NSScreen_type_screensHaveSeparateSpaces()
	r0 = convertObjCBoolToGo(ret)
	return
}

func NSStatusBar_alloc() (
	r0 NSStatusBar,
) {
	ret := C.NSStatusBar_type_alloc()
	r0 = NSStatusBar_fromPointer(ret)
	return
}

func NSStatusBar_systemStatusBar() (
	r0 NSStatusBar,
) {
	ret := C.NSStatusBar_type_systemStatusBar()
	r0 = NSStatusBar_fromPointer(ret)
	return
}

func NSStatusBarButton_alloc() (
	r0 NSStatusBarButton,
) {
	ret := C.NSStatusBarButton_type_alloc()
	r0 = NSStatusBarButton_fromPointer(ret)
	return
}

func NSStatusItem_alloc() (
	r0 NSStatusItem,
) {
	ret := C.NSStatusItem_type_alloc()
	r0 = NSStatusItem_fromPointer(ret)
	return
}

func NSText_alloc() (
	r0 NSText,
) {
	ret := C.NSText_type_alloc()
	r0 = NSText_fromPointer(ret)
	return
}

func NSTextField_alloc() (
	r0 NSTextField,
) {
	ret := C.NSTextField_type_alloc()
	r0 = NSTextField_fromPointer(ret)
	return
}

func NSTextField_labelWithAttributedString(
	attributedStringValue core.NSAttributedStringRef,
) (
	r0 NSTextField,
) {
	ret := C.NSTextField_type_labelWithAttributedString(
		objc.RefPointer(attributedStringValue),
	)
	r0 = NSTextField_fromPointer(ret)
	return
}

func NSTextField_labelWithString(
	stringValue core.NSStringRef,
) (
	r0 NSTextField,
) {
	ret := C.NSTextField_type_labelWithString(
		objc.RefPointer(stringValue),
	)
	r0 = NSTextField_fromPointer(ret)
	return
}

func NSTextField_textFieldWithString(
	stringValue core.NSStringRef,
) (
	r0 NSTextField,
) {
	ret := C.NSTextField_type_textFieldWithString(
		objc.RefPointer(stringValue),
	)
	r0 = NSTextField_fromPointer(ret)
	return
}

func NSTextField_wrappingLabelWithString(
	stringValue core.NSStringRef,
) (
	r0 NSTextField,
) {
	ret := C.NSTextField_type_wrappingLabelWithString(
		objc.RefPointer(stringValue),
	)
	r0 = NSTextField_fromPointer(ret)
	return
}

func NSTextContainer_alloc() (
	r0 NSTextContainer,
) {
	ret := C.NSTextContainer_type_alloc()
	r0 = NSTextContainer_fromPointer(ret)
	return
}

func NSViewController_alloc() (
	r0 NSViewController,
) {
	ret := C.NSViewController_type_alloc()
	r0 = NSViewController_fromPointer(ret)
	return
}

func NSVisualEffectView_alloc() (
	r0 NSVisualEffectView,
) {
	ret := C.NSVisualEffectView_type_alloc()
	r0 = NSVisualEffectView_fromPointer(ret)
	return
}

func NSWindow_alloc() (
	r0 NSWindow,
) {
	ret := C.NSWindow_type_alloc()
	r0 = NSWindow_fromPointer(ret)
	return
}

func NSWindow_windowWithContentViewController(
	contentViewController NSViewControllerRef,
) (
	r0 NSWindow,
) {
	ret := C.NSWindow_type_windowWithContentViewController(
		objc.RefPointer(contentViewController),
	)
	r0 = NSWindow_fromPointer(ret)
	return
}

func NSWindow_contentRectForFrameRect_styleMask(
	fRect core.NSRect,
	style core.NSUInteger,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_type_contentRectForFrameRect_styleMask(
		*(*C.NSRect)(unsafe.Pointer(&fRect)),
		C.ulong(style),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func NSWindow_frameRectForContentRect_styleMask(
	cRect core.NSRect,
	style core.NSUInteger,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_type_frameRectForContentRect_styleMask(
		*(*C.NSRect)(unsafe.Pointer(&cRect)),
		C.ulong(style),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func NSWindow_minFrameWidthWithTitle_styleMask(
	title core.NSStringRef,
	style core.NSUInteger,
) (
	r0 core.CGFloat,
) {
	ret := C.NSWindow_type_minFrameWidthWithTitle_styleMask(
		objc.RefPointer(title),
		C.ulong(style),
	)
	r0 = core.CGFloat(ret)
	return
}

func NSWindow_windowNumberAtPoint_belowWindowWithWindowNumber(
	point core.NSPoint,
	windowNumber core.NSInteger,
) (
	r0 core.NSInteger,
) {
	ret := C.NSWindow_type_windowNumberAtPoint_belowWindowWithWindowNumber(
		*(*C.NSPoint)(unsafe.Pointer(&point)),
		C.long(windowNumber),
	)
	r0 = core.NSInteger(ret)
	return
}

func NSWindow_allowsAutomaticWindowTabbing() (
	r0 bool,
) {
	ret := C.NSWindow_type_allowsAutomaticWindowTabbing()
	r0 = convertObjCBoolToGo(ret)
	return
}

func NSWindow_setAllowsAutomaticWindowTabbing(
	value bool,
) {
	C.NSWindow_type_setAllowsAutomaticWindowTabbing(
		convertToObjCBool(value),
	)
	return
}

func NSWorkspace_alloc() (
	r0 NSWorkspace,
) {
	ret := C.NSWorkspace_type_alloc()
	r0 = NSWorkspace_fromPointer(ret)
	return
}

func NSWorkspace_sharedWorkspace() (
	r0 NSWorkspace,
) {
	ret := C.NSWorkspace_type_sharedWorkspace()
	r0 = NSWorkspace_fromPointer(ret)
	return
}

func NSColor_alloc() (
	r0 NSColor,
) {
	ret := C.NSColor_type_alloc()
	r0 = NSColor_fromPointer(ret)
	return
}

func NSColor_colorFromPasteboard(
	pasteBoard NSPasteboardRef,
) (
	r0 NSColor,
) {
	ret := C.NSColor_type_colorFromPasteboard(
		objc.RefPointer(pasteBoard),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func NSColor_colorWithRed_green_blue_alpha(
	red core.CGFloat,
	green core.CGFloat,
	blue core.CGFloat,
	alpha core.CGFloat,
) (
	r0 NSColor,
) {
	ret := C.NSColor_type_colorWithRed_green_blue_alpha(
		C.double(red),
		C.double(green),
		C.double(blue),
		C.double(alpha),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func NSColor_ignoresAlpha() (
	r0 bool,
) {
	ret := C.NSColor_type_ignoresAlpha()
	r0 = convertObjCBoolToGo(ret)
	return
}

func NSColor_setIgnoresAlpha(
	value bool,
) {
	C.NSColor_type_setIgnoresAlpha(
		convertToObjCBool(value),
	)
	return
}

func NSColor_systemCyanColor() (
	r0 NSColor,
) {
	ret := C.NSColor_type_systemCyanColor()
	r0 = NSColor_fromPointer(ret)
	return
}

func NSColor_systemMintColor() (
	r0 NSColor,
) {
	ret := C.NSColor_type_systemMintColor()
	r0 = NSColor_fromPointer(ret)
	return
}

func NSColor_clearColor() (
	r0 NSColor,
) {
	ret := C.NSColor_type_clearColor()
	r0 = NSColor_fromPointer(ret)
	return
}

func NSTextView_alloc() (
	r0 NSTextView,
) {
	ret := C.NSTextView_type_alloc()
	r0 = NSTextView_fromPointer(ret)
	return
}

func NSTextView_registerForServices() {
	C.NSTextView_type_registerForServices()
	return
}

func NSTextView_fieldEditor() (
	r0 NSTextView,
) {
	ret := C.NSTextView_type_fieldEditor()
	r0 = NSTextView_fromPointer(ret)
	return
}

func NSTextView_stronglyReferencesTextStorage() (
	r0 bool,
) {
	ret := C.NSTextView_type_stronglyReferencesTextStorage()
	r0 = convertObjCBoolToGo(ret)
	return
}

func NSView_alloc() (
	r0 NSView,
) {
	ret := C.NSView_type_alloc()
	r0 = NSView_fromPointer(ret)
	return
}

func NSView_requiresConstraintBasedLayout() (
	r0 bool,
) {
	ret := C.NSView_type_requiresConstraintBasedLayout()
	r0 = convertObjCBoolToGo(ret)
	return
}

func NSView_focusView() (
	r0 NSView,
) {
	ret := C.NSView_type_focusView()
	r0 = NSView_fromPointer(ret)
	return
}

func NSView_defaultMenu() (
	r0 NSMenu,
) {
	ret := C.NSView_type_defaultMenu()
	r0 = NSMenu_fromPointer(ret)
	return
}

func NSView_compatibleWithResponsiveScrolling() (
	r0 bool,
) {
	ret := C.NSView_type_compatibleWithResponsiveScrolling()
	r0 = convertObjCBoolToGo(ret)
	return
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

func (x gen_NSBundle) InitWithURL_asNSBundle(
	url core.NSURLRef,
) (
	r0 NSBundle,
) {
	ret := C.NSBundle_inst_initWithURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)
	r0 = NSBundle_fromPointer(ret)
	return
}

func (x gen_NSBundle) InitWithPath_asNSBundle(
	path core.NSStringRef,
) (
	r0 NSBundle,
) {
	ret := C.NSBundle_inst_initWithPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)
	r0 = NSBundle_fromPointer(ret)
	return
}

func (x gen_NSBundle) LoadNibNamed_owner_options(
	name core.NSStringRef,
	owner objc.Ref,
	options core.NSDictionaryRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_inst_loadNibNamed_owner_options(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(owner),
		objc.RefPointer(options),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSBundle) URLForResource_withExtension_subdirectory(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
) (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_URLForResource_withExtension_subdirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) URLForResource_withExtension(
	name core.NSStringRef,
	ext core.NSStringRef,
) (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_URLForResource_withExtension(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) URLsForResourcesWithExtension_subdirectory(
	ext core.NSStringRef,
	subpath core.NSStringRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_inst_URLsForResourcesWithExtension_subdirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSBundle) URLForResource_withExtension_subdirectory_localization(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_URLForResource_withExtension_subdirectory_localization(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) URLsForResourcesWithExtension_subdirectory_localization(
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_inst_URLsForResourcesWithExtension_subdirectory_localization(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSBundle) PathForResource_ofType(
	name core.NSStringRef,
	ext core.NSStringRef,
) (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_pathForResource_ofType(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) PathForResource_ofType_inDirectory(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
) (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_pathForResource_ofType_inDirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) PathForResource_ofType_inDirectory_forLocalization(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_pathForResource_ofType_inDirectory_forLocalization(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) PathsForResourcesOfType_inDirectory(
	ext core.NSStringRef,
	subpath core.NSStringRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_inst_pathsForResourcesOfType_inDirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSBundle) PathsForResourcesOfType_inDirectory_forLocalization(
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_inst_pathsForResourcesOfType_inDirectory_forLocalization(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSBundle) LocalizedStringForKey_value_table(
	key core.NSStringRef,
	value core.NSStringRef,
	tableName core.NSStringRef,
) (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_localizedStringForKey_value_table(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
		objc.RefPointer(value),
		objc.RefPointer(tableName),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) URLForAuxiliaryExecutable(
	executableName core.NSStringRef,
) (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_URLForAuxiliaryExecutable(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(executableName),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) PathForAuxiliaryExecutable(
	executableName core.NSStringRef,
) (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_pathForAuxiliaryExecutable(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(executableName),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) ObjectForInfoDictionaryKey(
	key core.NSStringRef,
) (
	r0 objc.Object,
) {
	ret := C.NSBundle_inst_objectForInfoDictionaryKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSBundle) Load() (
	r0 bool,
) {
	ret := C.NSBundle_inst_load(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSBundle) Unload() (
	r0 bool,
) {
	ret := C.NSBundle_inst_unload(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSBundle) LocalizedAttributedStringForKey_value_table(
	key core.NSStringRef,
	value core.NSStringRef,
	tableName core.NSStringRef,
) (
	r0 core.NSAttributedString,
) {
	ret := C.NSBundle_inst_localizedAttributedStringForKey_value_table(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
		objc.RefPointer(value),
		objc.RefPointer(tableName),
	)
	r0 = core.NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSBundle) Init_asNSBundle() (
	r0 NSBundle,
) {
	ret := C.NSBundle_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSBundle_fromPointer(ret)
	return
}

func (x gen_NSBundle) ResourceURL() (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_resourceURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) ExecutableURL() (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_executableURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) PrivateFrameworksURL() (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_privateFrameworksURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) SharedFrameworksURL() (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_sharedFrameworksURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) BuiltInPlugInsURL() (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_builtInPlugInsURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) SharedSupportURL() (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_sharedSupportURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) AppStoreReceiptURL() (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_appStoreReceiptURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) ResourcePath() (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_resourcePath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) ExecutablePath() (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_executablePath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) PrivateFrameworksPath() (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_privateFrameworksPath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) SharedFrameworksPath() (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_sharedFrameworksPath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) BuiltInPlugInsPath() (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_builtInPlugInsPath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) SharedSupportPath() (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_sharedSupportPath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) BundleURL() (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_bundleURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) BundlePath() (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_bundlePath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) BundleIdentifier() (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_bundleIdentifier(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) InfoDictionary() (
	r0 core.NSDictionary,
) {
	ret := C.NSBundle_inst_infoDictionary(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSBundle) Localizations() (
	r0 core.NSArray,
) {
	ret := C.NSBundle_inst_localizations(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSBundle) PreferredLocalizations() (
	r0 core.NSArray,
) {
	ret := C.NSBundle_inst_preferredLocalizations(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSBundle) DevelopmentLocalization() (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_developmentLocalization(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) LocalizedInfoDictionary() (
	r0 core.NSDictionary,
) {
	ret := C.NSBundle_inst_localizedInfoDictionary(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSBundle) ExecutableArchitectures() (
	r0 core.NSArray,
) {
	ret := C.NSBundle_inst_executableArchitectures(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSBundle) IsLoaded() (
	r0 bool,
) {
	ret := C.NSBundle_inst_isLoaded(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
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

func (x gen_NSSound) InitWithContentsOfFile_byReference_asNSSound(
	path core.NSStringRef,
	byRef bool,
) (
	r0 NSSound,
) {
	ret := C.NSSound_inst_initWithContentsOfFile_byReference(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(byRef),
	)
	r0 = NSSound_fromPointer(ret)
	return
}

func (x gen_NSSound) InitWithContentsOfURL_byReference_asNSSound(
	url core.NSURLRef,
	byRef bool,
) (
	r0 NSSound,
) {
	ret := C.NSSound_inst_initWithContentsOfURL_byReference(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
		convertToObjCBool(byRef),
	)
	r0 = NSSound_fromPointer(ret)
	return
}

func (x gen_NSSound) InitWithData_asNSSound(
	data core.NSDataRef,
) (
	r0 NSSound,
) {
	ret := C.NSSound_inst_initWithData(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
	)
	r0 = NSSound_fromPointer(ret)
	return
}

func (x gen_NSSound) InitWithPasteboard_asNSSound(
	pasteboard NSPasteboardRef,
) (
	r0 NSSound,
) {
	ret := C.NSSound_inst_initWithPasteboard(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pasteboard),
	)
	r0 = NSSound_fromPointer(ret)
	return
}

func (x gen_NSSound) Pause() (
	r0 bool,
) {
	ret := C.NSSound_inst_pause(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSSound) Play() (
	r0 bool,
) {
	ret := C.NSSound_inst_play(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSSound) Resume() (
	r0 bool,
) {
	ret := C.NSSound_inst_resume(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSSound) Stop() (
	r0 bool,
) {
	ret := C.NSSound_inst_stop(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSSound) WriteToPasteboard(
	pasteboard NSPasteboardRef,
) {
	C.NSSound_inst_writeToPasteboard(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pasteboard),
	)
	return
}

func (x gen_NSSound) Init_asNSSound() (
	r0 NSSound,
) {
	ret := C.NSSound_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSSound_fromPointer(ret)
	return
}

func (x gen_NSSound) Delegate() (
	r0 objc.Object,
) {
	ret := C.NSSound_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSSound) SetDelegate(
	value objc.Ref,
) {
	C.NSSound_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSSound) Loops() (
	r0 bool,
) {
	ret := C.NSSound_inst_loops(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSSound) SetLoops(
	value bool,
) {
	C.NSSound_inst_setLoops(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSSound) IsPlaying() (
	r0 bool,
) {
	ret := C.NSSound_inst_isPlaying(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
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

func (x gen_NSApplication) Run() {
	C.NSApplication_inst_run(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSApplication) FinishLaunching() {
	C.NSApplication_inst_finishLaunching(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSApplication) Stop(
	sender objc.Ref,
) {
	C.NSApplication_inst_stop(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSApplication) SendEvent(
	event NSEventRef,
) {
	C.NSApplication_inst_sendEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	return
}

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

func (x gen_NSApplication) TryToPerform_with(
	action objc.Selector,
	object objc.Ref,
) (
	r0 bool,
) {
	ret := C.NSApplication_inst_tryToPerform_with(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(object),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSApplication) SendAction_to_from(
	action objc.Selector,
	target objc.Ref,
	sender objc.Ref,
) (
	r0 bool,
) {
	ret := C.NSApplication_inst_sendAction_to_from(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(sender),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSApplication) TargetForAction(
	action objc.Selector,
) (
	r0 objc.Object,
) {
	ret := C.NSApplication_inst_targetForAction(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSApplication) TargetForAction_to_from(
	action objc.Selector,
	target objc.Ref,
	sender objc.Ref,
) (
	r0 objc.Object,
) {
	ret := C.NSApplication_inst_targetForAction_to_from(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(sender),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSApplication) Terminate(
	sender objc.Ref,
) {
	C.NSApplication_inst_terminate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSApplication) ReplyToApplicationShouldTerminate(
	shouldTerminate bool,
) {
	C.NSApplication_inst_replyToApplicationShouldTerminate(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(shouldTerminate),
	)
	return
}

func (x gen_NSApplication) ActivateIgnoringOtherApps(
	flag bool,
) {
	C.NSApplication_inst_activateIgnoringOtherApps(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)
	return
}

func (x gen_NSApplication) Deactivate() {
	C.NSApplication_inst_deactivate(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSApplication) DisableRelaunchOnLogin() {
	C.NSApplication_inst_disableRelaunchOnLogin(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSApplication) EnableRelaunchOnLogin() {
	C.NSApplication_inst_enableRelaunchOnLogin(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSApplication) RegisterForRemoteNotifications() {
	C.NSApplication_inst_registerForRemoteNotifications(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSApplication) UnregisterForRemoteNotifications() {
	C.NSApplication_inst_unregisterForRemoteNotifications(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSApplication) ToggleTouchBarCustomizationPalette(
	sender objc.Ref,
) {
	C.NSApplication_inst_toggleTouchBarCustomizationPalette(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSApplication) CancelUserAttentionRequest(
	request core.NSInteger,
) {
	C.NSApplication_inst_cancelUserAttentionRequest(
		unsafe.Pointer(x.Pointer()),
		C.long(request),
	)
	return
}

func (x gen_NSApplication) RegisterUserInterfaceItemSearchHandler(
	handler objc.Ref,
) {
	C.NSApplication_inst_registerUserInterfaceItemSearchHandler(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(handler),
	)
	return
}

func (x gen_NSApplication) UnregisterUserInterfaceItemSearchHandler(
	handler objc.Ref,
) {
	C.NSApplication_inst_unregisterUserInterfaceItemSearchHandler(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(handler),
	)
	return
}

func (x gen_NSApplication) ShowHelp(
	sender objc.Ref,
) {
	C.NSApplication_inst_showHelp(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSApplication) ActivateContextHelpMode(
	sender objc.Ref,
) {
	C.NSApplication_inst_activateContextHelpMode(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSApplication) HideOtherApplications(
	sender objc.Ref,
) {
	C.NSApplication_inst_hideOtherApplications(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSApplication) UnhideAllApplications(
	sender objc.Ref,
) {
	C.NSApplication_inst_unhideAllApplications(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSApplication) ActivationPolicy() (
	r0 core.NSInteger,
) {
	ret := C.NSApplication_inst_activationPolicy(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSApplication) SetActivationPolicy(
	activationPolicy core.NSInteger,
) (
	r0 bool,
) {
	ret := C.NSApplication_inst_setActivationPolicy(
		unsafe.Pointer(x.Pointer()),
		C.long(activationPolicy),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSApplication) Init_asNSApplication() (
	r0 NSApplication,
) {
	ret := C.NSApplication_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSApplication_fromPointer(ret)
	return
}

func (x gen_NSApplication) Delegate() (
	r0 objc.Object,
) {
	ret := C.NSApplication_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSApplication) SetDelegate(
	value objc.Ref,
) {
	C.NSApplication_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSApplication) CurrentEvent() (
	r0 NSEvent,
) {
	ret := C.NSApplication_inst_currentEvent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSEvent_fromPointer(ret)
	return
}

func (x gen_NSApplication) IsRunning() (
	r0 bool,
) {
	ret := C.NSApplication_inst_isRunning(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSApplication) IsActive() (
	r0 bool,
) {
	ret := C.NSApplication_inst_isActive(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSApplication) IsRegisteredForRemoteNotifications() (
	r0 bool,
) {
	ret := C.NSApplication_inst_isRegisteredForRemoteNotifications(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSApplication) ApplicationIconImage() (
	r0 NSImage,
) {
	ret := C.NSApplication_inst_applicationIconImage(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSApplication) SetApplicationIconImage(
	value NSImageRef,
) {
	C.NSApplication_inst_setApplicationIconImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSApplication) HelpMenu() (
	r0 NSMenu,
) {
	ret := C.NSApplication_inst_helpMenu(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSMenu_fromPointer(ret)
	return
}

func (x gen_NSApplication) SetHelpMenu(
	value NSMenuRef,
) {
	C.NSApplication_inst_setHelpMenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSApplication) ServicesProvider() (
	r0 objc.Object,
) {
	ret := C.NSApplication_inst_servicesProvider(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSApplication) SetServicesProvider(
	value objc.Ref,
) {
	C.NSApplication_inst_setServicesProvider(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSApplication) IsFullKeyboardAccessEnabled() (
	r0 bool,
) {
	ret := C.NSApplication_inst_isFullKeyboardAccessEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSApplication) OrderedDocuments() (
	r0 core.NSArray,
) {
	ret := C.NSApplication_inst_orderedDocuments(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSApplication) OrderedWindows() (
	r0 core.NSArray,
) {
	ret := C.NSApplication_inst_orderedWindows(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSApplication) MainMenu() (
	r0 NSMenu,
) {
	ret := C.NSApplication_inst_mainMenu(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSMenu_fromPointer(ret)
	return
}

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

func (x gen_NSControl) InitWithFrame_asNSControl(
	frameRect core.NSRect,
) (
	r0 NSControl,
) {
	ret := C.NSControl_inst_initWithFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)
	r0 = NSControl_fromPointer(ret)
	return
}

func (x gen_NSControl) TakeDoubleValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_takeDoubleValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSControl) TakeFloatValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_takeFloatValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSControl) TakeIntValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_takeIntValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSControl) TakeIntegerValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_takeIntegerValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSControl) TakeObjectValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_takeObjectValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSControl) TakeStringValueFrom(
	sender objc.Ref,
) {
	C.NSControl_inst_takeStringValueFrom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

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

func (x gen_NSControl) ExpansionFrameWithFrame(
	contentFrame core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSControl_inst_expansionFrameWithFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentFrame)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSControl) AbortEditing() (
	r0 bool,
) {
	ret := C.NSControl_inst_abortEditing(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSControl) CurrentEditor() (
	r0 NSText,
) {
	ret := C.NSControl_inst_currentEditor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSText_fromPointer(ret)
	return
}

func (x gen_NSControl) ValidateEditing() {
	C.NSControl_inst_validateEditing(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

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

func (x gen_NSControl) EndEditing(
	textObj NSTextRef,
) {
	C.NSControl_inst_endEditing(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(textObj),
	)
	return
}

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

func (x gen_NSControl) SizeThatFits(
	size core.NSSize,
) (
	r0 core.NSSize,
) {
	ret := C.NSControl_inst_sizeThatFits(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSControl) SizeToFit() {
	C.NSControl_inst_sizeToFit(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSControl) SendAction_to(
	action objc.Selector,
	target objc.Ref,
) (
	r0 bool,
) {
	ret := C.NSControl_inst_sendAction_to(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(target),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSControl) PerformClick(
	sender objc.Ref,
) {
	C.NSControl_inst_performClick(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSControl) MouseDown(
	event NSEventRef,
) {
	C.NSControl_inst_mouseDown(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	return
}

func (x gen_NSControl) Init_asNSControl() (
	r0 NSControl,
) {
	ret := C.NSControl_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSControl_fromPointer(ret)
	return
}

func (x gen_NSControl) IsEnabled() (
	r0 bool,
) {
	ret := C.NSControl_inst_isEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSControl) SetEnabled(
	value bool,
) {
	C.NSControl_inst_setEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSControl) IntValue() (
	r0 int32,
) {
	ret := C.NSControl_inst_intValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = int32(ret)
	return
}

func (x gen_NSControl) SetIntValue(
	value int32,
) {
	C.NSControl_inst_setIntValue(
		unsafe.Pointer(x.Pointer()),
		C.int(value),
	)
	return
}

func (x gen_NSControl) IntegerValue() (
	r0 core.NSInteger,
) {
	ret := C.NSControl_inst_integerValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSControl) SetIntegerValue(
	value core.NSInteger,
) {
	C.NSControl_inst_setIntegerValue(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)
	return
}

func (x gen_NSControl) ObjectValue() (
	r0 objc.Object,
) {
	ret := C.NSControl_inst_objectValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSControl) SetObjectValue(
	value objc.Ref,
) {
	C.NSControl_inst_setObjectValue(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSControl) StringValue() (
	r0 core.NSString,
) {
	ret := C.NSControl_inst_stringValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSControl) SetStringValue(
	value core.NSStringRef,
) {
	C.NSControl_inst_setStringValue(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSControl) AttributedStringValue() (
	r0 core.NSAttributedString,
) {
	ret := C.NSControl_inst_attributedStringValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSControl) SetAttributedStringValue(
	value core.NSAttributedStringRef,
) {
	C.NSControl_inst_setAttributedStringValue(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSControl) Font() (
	r0 NSFont,
) {
	ret := C.NSControl_inst_font(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func (x gen_NSControl) SetFont(
	value NSFontRef,
) {
	C.NSControl_inst_setFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSControl) UsesSingleLineMode() (
	r0 bool,
) {
	ret := C.NSControl_inst_usesSingleLineMode(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSControl) SetUsesSingleLineMode(
	value bool,
) {
	C.NSControl_inst_setUsesSingleLineMode(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSControl) AllowsExpansionToolTips() (
	r0 bool,
) {
	ret := C.NSControl_inst_allowsExpansionToolTips(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSControl) SetAllowsExpansionToolTips(
	value bool,
) {
	C.NSControl_inst_setAllowsExpansionToolTips(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSControl) IsHighlighted() (
	r0 bool,
) {
	ret := C.NSControl_inst_isHighlighted(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSControl) SetHighlighted(
	value bool,
) {
	C.NSControl_inst_setHighlighted(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSControl) Action() (
	r0 objc.Selector,
) {
	ret := C.NSControl_inst_action(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.SelectorAt(ret)
	return
}

func (x gen_NSControl) SetAction(
	value objc.Selector,
) {
	C.NSControl_inst_setAction(
		unsafe.Pointer(x.Pointer()),
		value.SelectorAddress(),
	)
	return
}

func (x gen_NSControl) Target() (
	r0 objc.Object,
) {
	ret := C.NSControl_inst_target(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSControl) SetTarget(
	value objc.Ref,
) {
	C.NSControl_inst_setTarget(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSControl) IsContinuous() (
	r0 bool,
) {
	ret := C.NSControl_inst_isContinuous(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSControl) SetContinuous(
	value bool,
) {
	C.NSControl_inst_setContinuous(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSControl) Tag() (
	r0 core.NSInteger,
) {
	ret := C.NSControl_inst_tag(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSControl) SetTag(
	value core.NSInteger,
) {
	C.NSControl_inst_setTag(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)
	return
}

func (x gen_NSControl) RefusesFirstResponder() (
	r0 bool,
) {
	ret := C.NSControl_inst_refusesFirstResponder(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSControl) SetRefusesFirstResponder(
	value bool,
) {
	C.NSControl_inst_setRefusesFirstResponder(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSControl) IgnoresMultiClick() (
	r0 bool,
) {
	ret := C.NSControl_inst_ignoresMultiClick(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

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

func (x gen_NSButton) CompressWithPrioritizedCompressionOptions(
	prioritizedOptions core.NSArrayRef,
) {
	C.NSButton_inst_compressWithPrioritizedCompressionOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(prioritizedOptions),
	)
	return
}

func (x gen_NSButton) MinimumSizeWithPrioritizedCompressionOptions(
	prioritizedOptions core.NSArrayRef,
) (
	r0 core.NSSize,
) {
	ret := C.NSButton_inst_minimumSizeWithPrioritizedCompressionOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(prioritizedOptions),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSButton) SetNextState() {
	C.NSButton_inst_setNextState(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSButton) Highlight(
	flag bool,
) {
	C.NSButton_inst_highlight(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)
	return
}

func (x gen_NSButton) PerformKeyEquivalent(
	key NSEventRef,
) (
	r0 bool,
) {
	ret := C.NSButton_inst_performKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSButton) Init_asNSButton() (
	r0 NSButton,
) {
	ret := C.NSButton_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSButton_fromPointer(ret)
	return
}

func (x gen_NSButton) ContentTintColor() (
	r0 NSColor,
) {
	ret := C.NSButton_inst_contentTintColor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSButton) SetContentTintColor(
	value NSColorRef,
) {
	C.NSButton_inst_setContentTintColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSButton) HasDestructiveAction() (
	r0 bool,
) {
	ret := C.NSButton_inst_hasDestructiveAction(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSButton) SetHasDestructiveAction(
	value bool,
) {
	C.NSButton_inst_setHasDestructiveAction(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSButton) AlternateTitle() (
	r0 core.NSString,
) {
	ret := C.NSButton_inst_alternateTitle(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSButton) SetAlternateTitle(
	value core.NSStringRef,
) {
	C.NSButton_inst_setAlternateTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSButton) AttributedTitle() (
	r0 core.NSAttributedString,
) {
	ret := C.NSButton_inst_attributedTitle(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSButton) SetAttributedTitle(
	value core.NSAttributedStringRef,
) {
	C.NSButton_inst_setAttributedTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSButton) AttributedAlternateTitle() (
	r0 core.NSAttributedString,
) {
	ret := C.NSButton_inst_attributedAlternateTitle(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSButton) SetAttributedAlternateTitle(
	value core.NSAttributedStringRef,
) {
	C.NSButton_inst_setAttributedAlternateTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSButton) Title() (
	r0 core.NSString,
) {
	ret := C.NSButton_inst_title(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSButton) SetTitle(
	value core.NSStringRef,
) {
	C.NSButton_inst_setTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSButton) Sound() (
	r0 NSSound,
) {
	ret := C.NSButton_inst_sound(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSSound_fromPointer(ret)
	return
}

func (x gen_NSButton) SetSound(
	value NSSoundRef,
) {
	C.NSButton_inst_setSound(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSButton) IsSpringLoaded() (
	r0 bool,
) {
	ret := C.NSButton_inst_isSpringLoaded(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSButton) SetSpringLoaded(
	value bool,
) {
	C.NSButton_inst_setSpringLoaded(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSButton) MaxAcceleratorLevel() (
	r0 core.NSInteger,
) {
	ret := C.NSButton_inst_maxAcceleratorLevel(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSButton) SetMaxAcceleratorLevel(
	value core.NSInteger,
) {
	C.NSButton_inst_setMaxAcceleratorLevel(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)
	return
}

func (x gen_NSButton) Image() (
	r0 NSImage,
) {
	ret := C.NSButton_inst_image(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSButton) SetImage(
	value NSImageRef,
) {
	C.NSButton_inst_setImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSButton) AlternateImage() (
	r0 NSImage,
) {
	ret := C.NSButton_inst_alternateImage(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSButton) SetAlternateImage(
	value NSImageRef,
) {
	C.NSButton_inst_setAlternateImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSButton) IsBordered() (
	r0 bool,
) {
	ret := C.NSButton_inst_isBordered(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSButton) SetBordered(
	value bool,
) {
	C.NSButton_inst_setBordered(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSButton) IsTransparent() (
	r0 bool,
) {
	ret := C.NSButton_inst_isTransparent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSButton) SetTransparent(
	value bool,
) {
	C.NSButton_inst_setTransparent(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSButton) BezelColor() (
	r0 NSColor,
) {
	ret := C.NSButton_inst_bezelColor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSButton) SetBezelColor(
	value NSColorRef,
) {
	C.NSButton_inst_setBezelColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSButton) ShowsBorderOnlyWhileMouseInside() (
	r0 bool,
) {
	ret := C.NSButton_inst_showsBorderOnlyWhileMouseInside(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSButton) SetShowsBorderOnlyWhileMouseInside(
	value bool,
) {
	C.NSButton_inst_setShowsBorderOnlyWhileMouseInside(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSButton) ImageHugsTitle() (
	r0 bool,
) {
	ret := C.NSButton_inst_imageHugsTitle(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSButton) SetImageHugsTitle(
	value bool,
) {
	C.NSButton_inst_setImageHugsTitle(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSButton) AllowsMixedState() (
	r0 bool,
) {
	ret := C.NSButton_inst_allowsMixedState(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSButton) SetAllowsMixedState(
	value bool,
) {
	C.NSButton_inst_setAllowsMixedState(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSButton) State() (
	r0 core.NSInteger,
) {
	ret := C.NSButton_inst_state(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSButton) SetState(
	value core.NSInteger,
) {
	C.NSButton_inst_setState(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)
	return
}

func (x gen_NSButton) KeyEquivalent() (
	r0 core.NSString,
) {
	ret := C.NSButton_inst_keyEquivalent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

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

func (x gen_NSEvent) Init_asNSEvent() (
	r0 NSEvent,
) {
	ret := C.NSEvent_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSEvent_fromPointer(ret)
	return
}

func (x gen_NSEvent) LocationInWindow() (
	r0 core.NSPoint,
) {
	ret := C.NSEvent_inst_locationInWindow(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSEvent) Window() (
	r0 NSWindow,
) {
	ret := C.NSEvent_inst_window(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSWindow_fromPointer(ret)
	return
}

func (x gen_NSEvent) WindowNumber() (
	r0 core.NSInteger,
) {
	ret := C.NSEvent_inst_windowNumber(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSEvent) EventRef() (
	r0 unsafe.Pointer,
) {
	ret := C.NSEvent_inst_eventRef(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = ret
	return
}

func (x gen_NSEvent) Characters() (
	r0 core.NSString,
) {
	ret := C.NSEvent_inst_characters(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSEvent) CharactersIgnoringModifiers() (
	r0 core.NSString,
) {
	ret := C.NSEvent_inst_charactersIgnoringModifiers(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSEvent) IsARepeat() (
	r0 bool,
) {
	ret := C.NSEvent_inst_isARepeat(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSEvent) ButtonNumber() (
	r0 core.NSInteger,
) {
	ret := C.NSEvent_inst_buttonNumber(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSEvent) ClickCount() (
	r0 core.NSInteger,
) {
	ret := C.NSEvent_inst_clickCount(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSEvent) EventNumber() (
	r0 core.NSInteger,
) {
	ret := C.NSEvent_inst_eventNumber(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSEvent) TrackingNumber() (
	r0 core.NSInteger,
) {
	ret := C.NSEvent_inst_trackingNumber(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSEvent) UserData() (
	r0 unsafe.Pointer,
) {
	ret := C.NSEvent_inst_userData(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = ret
	return
}

func (x gen_NSEvent) Data1() (
	r0 core.NSInteger,
) {
	ret := C.NSEvent_inst_data1(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSEvent) Data2() (
	r0 core.NSInteger,
) {
	ret := C.NSEvent_inst_data2(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSEvent) DeltaX() (
	r0 core.CGFloat,
) {
	ret := C.NSEvent_inst_deltaX(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSEvent) DeltaY() (
	r0 core.CGFloat,
) {
	ret := C.NSEvent_inst_deltaY(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSEvent) DeltaZ() (
	r0 core.CGFloat,
) {
	ret := C.NSEvent_inst_deltaZ(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSEvent) Stage() (
	r0 core.NSInteger,
) {
	ret := C.NSEvent_inst_stage(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSEvent) StageTransition() (
	r0 core.CGFloat,
) {
	ret := C.NSEvent_inst_stageTransition(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSEvent) CapabilityMask() (
	r0 core.NSUInteger,
) {
	ret := C.NSEvent_inst_capabilityMask(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSEvent) DeviceID() (
	r0 core.NSUInteger,
) {
	ret := C.NSEvent_inst_deviceID(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSEvent) IsEnteringProximity() (
	r0 bool,
) {
	ret := C.NSEvent_inst_isEnteringProximity(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSEvent) PointingDeviceID() (
	r0 core.NSUInteger,
) {
	ret := C.NSEvent_inst_pointingDeviceID(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSEvent) PointingDeviceSerialNumber() (
	r0 core.NSUInteger,
) {
	ret := C.NSEvent_inst_pointingDeviceSerialNumber(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSEvent) SystemTabletID() (
	r0 core.NSUInteger,
) {
	ret := C.NSEvent_inst_systemTabletID(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSEvent) TabletID() (
	r0 core.NSUInteger,
) {
	ret := C.NSEvent_inst_tabletID(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSEvent) VendorID() (
	r0 core.NSUInteger,
) {
	ret := C.NSEvent_inst_vendorID(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSEvent) VendorPointingDeviceType() (
	r0 core.NSUInteger,
) {
	ret := C.NSEvent_inst_vendorPointingDeviceType(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSEvent) AbsoluteX() (
	r0 core.NSInteger,
) {
	ret := C.NSEvent_inst_absoluteX(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSEvent) AbsoluteY() (
	r0 core.NSInteger,
) {
	ret := C.NSEvent_inst_absoluteY(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSEvent) AbsoluteZ() (
	r0 core.NSInteger,
) {
	ret := C.NSEvent_inst_absoluteZ(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSEvent) Tilt() (
	r0 core.NSPoint,
) {
	ret := C.NSEvent_inst_tilt(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSEvent) VendorDefined() (
	r0 objc.Object,
) {
	ret := C.NSEvent_inst_vendorDefined(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSEvent) Magnification() (
	r0 core.CGFloat,
) {
	ret := C.NSEvent_inst_magnification(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSEvent) HasPreciseScrollingDeltas() (
	r0 bool,
) {
	ret := C.NSEvent_inst_hasPreciseScrollingDeltas(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSEvent) ScrollingDeltaX() (
	r0 core.CGFloat,
) {
	ret := C.NSEvent_inst_scrollingDeltaX(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSEvent) ScrollingDeltaY() (
	r0 core.CGFloat,
) {
	ret := C.NSEvent_inst_scrollingDeltaY(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSEvent) IsDirectionInvertedFromDevice() (
	r0 bool,
) {
	ret := C.NSEvent_inst_isDirectionInvertedFromDevice(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
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

func (x gen_NSFont) Set() {
	C.NSFont_inst_set(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSFont) FontWithSize(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_inst_fontWithSize(
		unsafe.Pointer(x.Pointer()),
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func (x gen_NSFont) Init_asNSFont() (
	r0 NSFont,
) {
	ret := C.NSFont_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func (x gen_NSFont) PointSize() (
	r0 core.CGFloat,
) {
	ret := C.NSFont_inst_pointSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSFont) IsFixedPitch() (
	r0 bool,
) {
	ret := C.NSFont_inst_isFixedPitch(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSFont) MostCompatibleStringEncoding() (
	r0 core.NSStringEncoding,
) {
	ret := C.NSFont_inst_mostCompatibleStringEncoding(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSStringEncoding(ret)
	return
}

func (x gen_NSFont) NumberOfGlyphs() (
	r0 core.NSUInteger,
) {
	ret := C.NSFont_inst_numberOfGlyphs(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSFont) DisplayName() (
	r0 core.NSString,
) {
	ret := C.NSFont_inst_displayName(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSFont) FamilyName() (
	r0 core.NSString,
) {
	ret := C.NSFont_inst_familyName(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSFont) FontName() (
	r0 core.NSString,
) {
	ret := C.NSFont_inst_fontName(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSFont) IsVertical() (
	r0 bool,
) {
	ret := C.NSFont_inst_isVertical(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSFont) VerticalFont() (
	r0 NSFont,
) {
	ret := C.NSFont_inst_verticalFont(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSFont_fromPointer(ret)
	return
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

func (x gen_NSImage) InitByReferencingFile_asNSImage(
	fileName core.NSStringRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initByReferencingFile(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fileName),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) InitByReferencingURL_asNSImage(
	url core.NSURLRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initByReferencingURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) InitWithContentsOfFile_asNSImage(
	fileName core.NSStringRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initWithContentsOfFile(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fileName),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) InitWithContentsOfURL_asNSImage(
	url core.NSURLRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initWithContentsOfURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) InitWithData_asNSImage(
	data core.NSDataRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initWithData(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) InitWithDataIgnoringOrientation_asNSImage(
	data core.NSDataRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initWithDataIgnoringOrientation(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) InitWithPasteboard_asNSImage(
	pasteboard NSPasteboardRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initWithPasteboard(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pasteboard),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) InitWithSize_asNSImage(
	size core.NSSize,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initWithSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) IsTemplate() (
	r0 bool,
) {
	ret := C.NSImage_inst_isTemplate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSImage) AddRepresentations(
	imageReps core.NSArrayRef,
) {
	C.NSImage_inst_addRepresentations(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(imageReps),
	)
	return
}

func (x gen_NSImage) DrawInRect(
	rect core.NSRect,
) {
	C.NSImage_inst_drawInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	return
}

func (x gen_NSImage) LockFocus() {
	C.NSImage_inst_lockFocus(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSImage) LockFocusFlipped(
	flipped bool,
) {
	C.NSImage_inst_lockFocusFlipped(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flipped),
	)
	return
}

func (x gen_NSImage) UnlockFocus() {
	C.NSImage_inst_unlockFocus(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSImage) Recache() {
	C.NSImage_inst_recache(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSImage) CancelIncrementalLoad() {
	C.NSImage_inst_cancelIncrementalLoad(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSImage) LayerContentsForContentsScale(
	layerContentsScale core.CGFloat,
) (
	r0 objc.Object,
) {
	ret := C.NSImage_inst_layerContentsForContentsScale(
		unsafe.Pointer(x.Pointer()),
		C.double(layerContentsScale),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSImage) RecommendedLayerContentsScale(
	preferredContentsScale core.CGFloat,
) (
	r0 core.CGFloat,
) {
	ret := C.NSImage_inst_recommendedLayerContentsScale(
		unsafe.Pointer(x.Pointer()),
		C.double(preferredContentsScale),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSImage) Init_asNSImage() (
	r0 NSImage,
) {
	ret := C.NSImage_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) Delegate() (
	r0 objc.Object,
) {
	ret := C.NSImage_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSImage) SetDelegate(
	value objc.Ref,
) {
	C.NSImage_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSImage) Size() (
	r0 core.NSSize,
) {
	ret := C.NSImage_inst_size(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSImage) SetSize(
	value core.NSSize,
) {
	C.NSImage_inst_setSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSImage) SetTemplate(
	value bool,
) {
	C.NSImage_inst_setTemplate(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSImage) Representations() (
	r0 core.NSArray,
) {
	ret := C.NSImage_inst_representations(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSImage) PrefersColorMatch() (
	r0 bool,
) {
	ret := C.NSImage_inst_prefersColorMatch(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSImage) SetPrefersColorMatch(
	value bool,
) {
	C.NSImage_inst_setPrefersColorMatch(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSImage) UsesEPSOnResolutionMismatch() (
	r0 bool,
) {
	ret := C.NSImage_inst_usesEPSOnResolutionMismatch(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSImage) SetUsesEPSOnResolutionMismatch(
	value bool,
) {
	C.NSImage_inst_setUsesEPSOnResolutionMismatch(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSImage) MatchesOnMultipleResolution() (
	r0 bool,
) {
	ret := C.NSImage_inst_matchesOnMultipleResolution(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSImage) SetMatchesOnMultipleResolution(
	value bool,
) {
	C.NSImage_inst_setMatchesOnMultipleResolution(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSImage) IsValid() (
	r0 bool,
) {
	ret := C.NSImage_inst_isValid(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSImage) BackgroundColor() (
	r0 NSColor,
) {
	ret := C.NSImage_inst_backgroundColor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSImage) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSImage_inst_setBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSImage) AlignmentRect() (
	r0 core.NSRect,
) {
	ret := C.NSImage_inst_alignmentRect(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSImage) SetAlignmentRect(
	value core.NSRect,
) {
	C.NSImage_inst_setAlignmentRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSImage) TIFFRepresentation() (
	r0 core.NSData,
) {
	ret := C.NSImage_inst_TIFFRepresentation(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSData_fromPointer(ret)
	return
}

func (x gen_NSImage) AccessibilityDescription() (
	r0 core.NSString,
) {
	ret := C.NSImage_inst_accessibilityDescription(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSImage) SetAccessibilityDescription(
	value core.NSStringRef,
) {
	C.NSImage_inst_setAccessibilityDescription(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSImage) MatchesOnlyOnBestFittingAxis() (
	r0 bool,
) {
	ret := C.NSImage_inst_matchesOnlyOnBestFittingAxis(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

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

func (x gen_NSImageView) Init_asNSImageView() (
	r0 NSImageView,
) {
	ret := C.NSImageView_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSImageView_fromPointer(ret)
	return
}

func (x gen_NSImageView) Image() (
	r0 NSImage,
) {
	ret := C.NSImageView_inst_image(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImageView) SetImage(
	value NSImageRef,
) {
	C.NSImageView_inst_setImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSImageView) Animates() (
	r0 bool,
) {
	ret := C.NSImageView_inst_animates(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSImageView) SetAnimates(
	value bool,
) {
	C.NSImageView_inst_setAnimates(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSImageView) IsEditable() (
	r0 bool,
) {
	ret := C.NSImageView_inst_isEditable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSImageView) SetEditable(
	value bool,
) {
	C.NSImageView_inst_setEditable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSImageView) AllowsCutCopyPaste() (
	r0 bool,
) {
	ret := C.NSImageView_inst_allowsCutCopyPaste(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSImageView) SetAllowsCutCopyPaste(
	value bool,
) {
	C.NSImageView_inst_setAllowsCutCopyPaste(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSImageView) ContentTintColor() (
	r0 NSColor,
) {
	ret := C.NSImageView_inst_contentTintColor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

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

func (x gen_NSNib) InitWithNibData_bundle_asNSNib(
	nibData core.NSDataRef,
	bundle NSBundleRef,
) (
	r0 NSNib,
) {
	ret := C.NSNib_inst_initWithNibData_bundle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(nibData),
		objc.RefPointer(bundle),
	)
	r0 = NSNib_fromPointer(ret)
	return
}

func (x gen_NSNib) InstantiateWithOwner_topLevelObjects(
	owner objc.Ref,
	topLevelObjects core.NSArrayRef,
) (
	r0 bool,
) {
	ret := C.NSNib_inst_instantiateWithOwner_topLevelObjects(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(owner),
		objc.RefPointer(topLevelObjects),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSNib) Init_asNSNib() (
	r0 NSNib,
) {
	ret := C.NSNib_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSNib_fromPointer(ret)
	return
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

func (x gen_NSPasteboard) ReleaseGlobally() {
	C.NSPasteboard_inst_releaseGlobally(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSPasteboard) ClearContents() (
	r0 core.NSInteger,
) {
	ret := C.NSPasteboard_inst_clearContents(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSPasteboard) WriteObjects(
	objects core.NSArrayRef,
) (
	r0 bool,
) {
	ret := C.NSPasteboard_inst_writeObjects(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(objects),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSPasteboard) ReadObjectsForClasses_options(
	classArray core.NSArrayRef,
	options core.NSDictionaryRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSPasteboard_inst_readObjectsForClasses_options(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(classArray),
		objc.RefPointer(options),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSPasteboard) CanReadItemWithDataConformingToTypes(
	types core.NSArrayRef,
) (
	r0 bool,
) {
	ret := C.NSPasteboard_inst_canReadItemWithDataConformingToTypes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(types),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSPasteboard) CanReadObjectForClasses_options(
	classArray core.NSArrayRef,
	options core.NSDictionaryRef,
) (
	r0 bool,
) {
	ret := C.NSPasteboard_inst_canReadObjectForClasses_options(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(classArray),
		objc.RefPointer(options),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSPasteboard) DeclareTypes_owner(
	newTypes core.NSArrayRef,
	newOwner objc.Ref,
) (
	r0 core.NSInteger,
) {
	ret := C.NSPasteboard_inst_declareTypes_owner(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newTypes),
		objc.RefPointer(newOwner),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSPasteboard) AddTypes_owner(
	newTypes core.NSArrayRef,
	newOwner objc.Ref,
) (
	r0 core.NSInteger,
) {
	ret := C.NSPasteboard_inst_addTypes_owner(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newTypes),
		objc.RefPointer(newOwner),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSPasteboard) WriteFileContents(
	filename core.NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSPasteboard_inst_writeFileContents(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(filename),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSPasteboard) Init_asNSPasteboard() (
	r0 NSPasteboard,
) {
	ret := C.NSPasteboard_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSPasteboard_fromPointer(ret)
	return
}

func (x gen_NSPasteboard) PasteboardItems() (
	r0 core.NSArray,
) {
	ret := C.NSPasteboard_inst_pasteboardItems(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSPasteboard) Types() (
	r0 core.NSArray,
) {
	ret := C.NSPasteboard_inst_types(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSPasteboard) ChangeCount() (
	r0 core.NSInteger,
) {
	ret := C.NSPasteboard_inst_changeCount(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
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

func (x gen_NSLayoutManager) Init_asNSLayoutManager() (
	r0 NSLayoutManager,
) {
	ret := C.NSLayoutManager_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSLayoutManager_fromPointer(ret)
	return
}

func (x gen_NSLayoutManager) AddTextContainer(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_addTextContainer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)
	return
}

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

func (x gen_NSLayoutManager) RemoveTextContainerAtIndex(
	index core.NSUInteger,
) {
	C.NSLayoutManager_inst_removeTextContainerAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(index),
	)
	return
}

func (x gen_NSLayoutManager) TextContainerChangedGeometry(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_textContainerChangedGeometry(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)
	return
}

func (x gen_NSLayoutManager) TextContainerChangedTextView(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_textContainerChangedTextView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)
	return
}

func (x gen_NSLayoutManager) UsedRectForTextContainer(
	container NSTextContainerRef,
) (
	r0 core.NSRect,
) {
	ret := C.NSLayoutManager_inst_usedRectForTextContainer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

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

func (x gen_NSLayoutManager) EnsureLayoutForTextContainer(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_ensureLayoutForTextContainer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)
	return
}

func (x gen_NSLayoutManager) CharacterIndexForGlyphAtIndex(
	glyphIndex core.NSUInteger,
) (
	r0 core.NSUInteger,
) {
	ret := C.NSLayoutManager_inst_characterIndexForGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSLayoutManager) GlyphIndexForCharacterAtIndex(
	charIndex core.NSUInteger,
) (
	r0 core.NSUInteger,
) {
	ret := C.NSLayoutManager_inst_glyphIndexForCharacterAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(charIndex),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSLayoutManager) IsValidGlyphIndex(
	glyphIndex core.NSUInteger,
) (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_isValidGlyphIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

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

func (x gen_NSLayoutManager) AttachmentSizeForGlyphAtIndex(
	glyphIndex core.NSUInteger,
) (
	r0 core.NSSize,
) {
	ret := C.NSLayoutManager_inst_attachmentSizeForGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSLayoutManager) DrawsOutsideLineFragmentForGlyphAtIndex(
	glyphIndex core.NSUInteger,
) (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_drawsOutsideLineFragmentForGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSLayoutManager) FirstUnlaidCharacterIndex() (
	r0 core.NSUInteger,
) {
	ret := C.NSLayoutManager_inst_firstUnlaidCharacterIndex(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSLayoutManager) FirstUnlaidGlyphIndex() (
	r0 core.NSUInteger,
) {
	ret := C.NSLayoutManager_inst_firstUnlaidGlyphIndex(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSLayoutManager) NotShownAttributeForGlyphAtIndex(
	glyphIndex core.NSUInteger,
) (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_notShownAttributeForGlyphAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSLayoutManager) LayoutManagerOwnsFirstResponderInWindow(
	window NSWindowRef,
) (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_layoutManagerOwnsFirstResponderInWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(window),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSLayoutManager) DefaultLineHeightForFont(
	theFont NSFontRef,
) (
	r0 core.CGFloat,
) {
	ret := C.NSLayoutManager_inst_defaultLineHeightForFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(theFont),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSLayoutManager) DefaultBaselineOffsetForFont(
	theFont NSFontRef,
) (
	r0 core.CGFloat,
) {
	ret := C.NSLayoutManager_inst_defaultBaselineOffsetForFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(theFont),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSLayoutManager) Delegate() (
	r0 objc.Object,
) {
	ret := C.NSLayoutManager_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSLayoutManager) SetDelegate(
	value objc.Ref,
) {
	C.NSLayoutManager_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSLayoutManager) AllowsNonContiguousLayout() (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_allowsNonContiguousLayout(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSLayoutManager) SetAllowsNonContiguousLayout(
	value bool,
) {
	C.NSLayoutManager_inst_setAllowsNonContiguousLayout(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSLayoutManager) HasNonContiguousLayout() (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_hasNonContiguousLayout(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSLayoutManager) ShowsInvisibleCharacters() (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_showsInvisibleCharacters(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSLayoutManager) SetShowsInvisibleCharacters(
	value bool,
) {
	C.NSLayoutManager_inst_setShowsInvisibleCharacters(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSLayoutManager) ShowsControlCharacters() (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_showsControlCharacters(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSLayoutManager) SetShowsControlCharacters(
	value bool,
) {
	C.NSLayoutManager_inst_setShowsControlCharacters(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSLayoutManager) UsesFontLeading() (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_usesFontLeading(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSLayoutManager) SetUsesFontLeading(
	value bool,
) {
	C.NSLayoutManager_inst_setUsesFontLeading(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSLayoutManager) BackgroundLayoutEnabled() (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_backgroundLayoutEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSLayoutManager) SetBackgroundLayoutEnabled(
	value bool,
) {
	C.NSLayoutManager_inst_setBackgroundLayoutEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSLayoutManager) LimitsLayoutForSuspiciousContents() (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_limitsLayoutForSuspiciousContents(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSLayoutManager) SetLimitsLayoutForSuspiciousContents(
	value bool,
) {
	C.NSLayoutManager_inst_setLimitsLayoutForSuspiciousContents(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSLayoutManager) UsesDefaultHyphenation() (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_usesDefaultHyphenation(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSLayoutManager) SetUsesDefaultHyphenation(
	value bool,
) {
	C.NSLayoutManager_inst_setUsesDefaultHyphenation(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSLayoutManager) TextContainers() (
	r0 core.NSArray,
) {
	ret := C.NSLayoutManager_inst_textContainers(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSLayoutManager) NumberOfGlyphs() (
	r0 core.NSUInteger,
) {
	ret := C.NSLayoutManager_inst_numberOfGlyphs(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSLayoutManager) ExtraLineFragmentRect() (
	r0 core.NSRect,
) {
	ret := C.NSLayoutManager_inst_extraLineFragmentRect(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSLayoutManager) ExtraLineFragmentTextContainer() (
	r0 NSTextContainer,
) {
	ret := C.NSLayoutManager_inst_extraLineFragmentTextContainer(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSTextContainer_fromPointer(ret)
	return
}

func (x gen_NSLayoutManager) ExtraLineFragmentUsedRect() (
	r0 core.NSRect,
) {
	ret := C.NSLayoutManager_inst_extraLineFragmentUsedRect(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSLayoutManager) FirstTextView() (
	r0 NSTextView,
) {
	ret := C.NSLayoutManager_inst_firstTextView(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSTextView_fromPointer(ret)
	return
}

func (x gen_NSLayoutManager) TextViewForBeginningOfSelection() (
	r0 NSTextView,
) {
	ret := C.NSLayoutManager_inst_textViewForBeginningOfSelection(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSTextView_fromPointer(ret)
	return
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

func (x gen_NSMenu) InitWithTitle_asNSMenu(
	title core.NSStringRef,
) (
	r0 NSMenu,
) {
	ret := C.NSMenu_inst_initWithTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(title),
	)
	r0 = NSMenu_fromPointer(ret)
	return
}

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

func (x gen_NSMenu) InsertItemWithTitle_action_keyEquivalent_atIndex(
	string core.NSStringRef,
	selector objc.Selector,
	charCode core.NSStringRef,
	index core.NSInteger,
) (
	r0 NSMenuItem,
) {
	ret := C.NSMenu_inst_insertItemWithTitle_action_keyEquivalent_atIndex(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		selector.SelectorAddress(),
		objc.RefPointer(charCode),
		C.long(index),
	)
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func (x gen_NSMenu) AddItem(
	newItem NSMenuItemRef,
) {
	C.NSMenu_inst_addItem(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newItem),
	)
	return
}

func (x gen_NSMenu) AddItemWithTitle_action_keyEquivalent(
	string core.NSStringRef,
	selector objc.Selector,
	charCode core.NSStringRef,
) (
	r0 NSMenuItem,
) {
	ret := C.NSMenu_inst_addItemWithTitle_action_keyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		selector.SelectorAddress(),
		objc.RefPointer(charCode),
	)
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func (x gen_NSMenu) RemoveItem(
	item NSMenuItemRef,
) {
	C.NSMenu_inst_removeItem(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
	)
	return
}

func (x gen_NSMenu) RemoveItemAtIndex(
	index core.NSInteger,
) {
	C.NSMenu_inst_removeItemAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)
	return
}

func (x gen_NSMenu) ItemChanged(
	item NSMenuItemRef,
) {
	C.NSMenu_inst_itemChanged(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
	)
	return
}

func (x gen_NSMenu) RemoveAllItems() {
	C.NSMenu_inst_removeAllItems(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSMenu) ItemWithTag(
	tag core.NSInteger,
) (
	r0 NSMenuItem,
) {
	ret := C.NSMenu_inst_itemWithTag(
		unsafe.Pointer(x.Pointer()),
		C.long(tag),
	)
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func (x gen_NSMenu) ItemWithTitle(
	title core.NSStringRef,
) (
	r0 NSMenuItem,
) {
	ret := C.NSMenu_inst_itemWithTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(title),
	)
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func (x gen_NSMenu) ItemAtIndex(
	index core.NSInteger,
) (
	r0 NSMenuItem,
) {
	ret := C.NSMenu_inst_itemAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func (x gen_NSMenu) IndexOfItem(
	item NSMenuItemRef,
) (
	r0 core.NSInteger,
) {
	ret := C.NSMenu_inst_indexOfItem(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSMenu) IndexOfItemWithTitle(
	title core.NSStringRef,
) (
	r0 core.NSInteger,
) {
	ret := C.NSMenu_inst_indexOfItemWithTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(title),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSMenu) IndexOfItemWithTag(
	tag core.NSInteger,
) (
	r0 core.NSInteger,
) {
	ret := C.NSMenu_inst_indexOfItemWithTag(
		unsafe.Pointer(x.Pointer()),
		C.long(tag),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSMenu) IndexOfItemWithTarget_andAction(
	target objc.Ref,
	actionSelector objc.Selector,
) (
	r0 core.NSInteger,
) {
	ret := C.NSMenu_inst_indexOfItemWithTarget_andAction(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(target),
		actionSelector.SelectorAddress(),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSMenu) IndexOfItemWithRepresentedObject(
	object objc.Ref,
) (
	r0 core.NSInteger,
) {
	ret := C.NSMenu_inst_indexOfItemWithRepresentedObject(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSMenu) IndexOfItemWithSubmenu(
	submenu NSMenuRef,
) (
	r0 core.NSInteger,
) {
	ret := C.NSMenu_inst_indexOfItemWithSubmenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(submenu),
	)
	r0 = core.NSInteger(ret)
	return
}

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

func (x gen_NSMenu) SubmenuAction(
	sender objc.Ref,
) {
	C.NSMenu_inst_submenuAction(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSMenu) Update() {
	C.NSMenu_inst_update(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSMenu) PerformKeyEquivalent(
	event NSEventRef,
) (
	r0 bool,
) {
	ret := C.NSMenu_inst_performKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSMenu) PerformActionForItemAtIndex(
	index core.NSInteger,
) {
	C.NSMenu_inst_performActionForItemAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)
	return
}

func (x gen_NSMenu) PopUpMenuPositioningItem_atLocation_inView(
	item NSMenuItemRef,
	location core.NSPoint,
	view NSViewRef,
) (
	r0 bool,
) {
	ret := C.NSMenu_inst_popUpMenuPositioningItem_atLocation_inView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
		*(*C.NSPoint)(unsafe.Pointer(&location)),
		objc.RefPointer(view),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSMenu) CancelTracking() {
	C.NSMenu_inst_cancelTracking(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSMenu) CancelTrackingWithoutAnimation() {
	C.NSMenu_inst_cancelTrackingWithoutAnimation(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSMenu) Init_asNSMenu() (
	r0 NSMenu,
) {
	ret := C.NSMenu_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSMenu_fromPointer(ret)
	return
}

func (x gen_NSMenu) MenuBarHeight() (
	r0 core.CGFloat,
) {
	ret := C.NSMenu_inst_menuBarHeight(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSMenu) NumberOfItems() (
	r0 core.NSInteger,
) {
	ret := C.NSMenu_inst_numberOfItems(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSMenu) ItemArray() (
	r0 core.NSArray,
) {
	ret := C.NSMenu_inst_itemArray(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSMenu) SetItemArray(
	value core.NSArrayRef,
) {
	C.NSMenu_inst_setItemArray(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenu) Supermenu() (
	r0 NSMenu,
) {
	ret := C.NSMenu_inst_supermenu(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSMenu_fromPointer(ret)
	return
}

func (x gen_NSMenu) SetSupermenu(
	value NSMenuRef,
) {
	C.NSMenu_inst_setSupermenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenu) AutoenablesItems() (
	r0 bool,
) {
	ret := C.NSMenu_inst_autoenablesItems(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSMenu) SetAutoenablesItems(
	value bool,
) {
	C.NSMenu_inst_setAutoenablesItems(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSMenu) Font() (
	r0 NSFont,
) {
	ret := C.NSMenu_inst_font(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func (x gen_NSMenu) SetFont(
	value NSFontRef,
) {
	C.NSMenu_inst_setFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenu) Title() (
	r0 core.NSString,
) {
	ret := C.NSMenu_inst_title(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSMenu) SetTitle(
	value core.NSStringRef,
) {
	C.NSMenu_inst_setTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenu) MinimumWidth() (
	r0 core.CGFloat,
) {
	ret := C.NSMenu_inst_minimumWidth(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSMenu) SetMinimumWidth(
	value core.CGFloat,
) {
	C.NSMenu_inst_setMinimumWidth(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_NSMenu) Size() (
	r0 core.NSSize,
) {
	ret := C.NSMenu_inst_size(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSMenu) AllowsContextMenuPlugIns() (
	r0 bool,
) {
	ret := C.NSMenu_inst_allowsContextMenuPlugIns(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSMenu) SetAllowsContextMenuPlugIns(
	value bool,
) {
	C.NSMenu_inst_setAllowsContextMenuPlugIns(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSMenu) ShowsStateColumn() (
	r0 bool,
) {
	ret := C.NSMenu_inst_showsStateColumn(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSMenu) SetShowsStateColumn(
	value bool,
) {
	C.NSMenu_inst_setShowsStateColumn(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSMenu) HighlightedItem() (
	r0 NSMenuItem,
) {
	ret := C.NSMenu_inst_highlightedItem(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func (x gen_NSMenu) Delegate() (
	r0 objc.Object,
) {
	ret := C.NSMenu_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

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

func (x gen_NSPopover) PerformClose(
	sender objc.Ref,
) {
	C.NSPopover_inst_performClose(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSPopover) Close() {
	C.NSPopover_inst_close(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSPopover) Init_asNSPopover() (
	r0 NSPopover,
) {
	ret := C.NSPopover_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSPopover_fromPointer(ret)
	return
}

func (x gen_NSPopover) Behavior() (
	r0 core.NSInteger,
) {
	ret := C.NSPopover_inst_behavior(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSPopover) SetBehavior(
	value core.NSInteger,
) {
	C.NSPopover_inst_setBehavior(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)
	return
}

func (x gen_NSPopover) PositioningRect() (
	r0 core.NSRect,
) {
	ret := C.NSPopover_inst_positioningRect(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSPopover) SetPositioningRect(
	value core.NSRect,
) {
	C.NSPopover_inst_setPositioningRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSPopover) Animates() (
	r0 bool,
) {
	ret := C.NSPopover_inst_animates(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSPopover) SetAnimates(
	value bool,
) {
	C.NSPopover_inst_setAnimates(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSPopover) ContentSize() (
	r0 core.NSSize,
) {
	ret := C.NSPopover_inst_contentSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSPopover) SetContentSize(
	value core.NSSize,
) {
	C.NSPopover_inst_setContentSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSPopover) IsShown() (
	r0 bool,
) {
	ret := C.NSPopover_inst_isShown(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSPopover) IsDetached() (
	r0 bool,
) {
	ret := C.NSPopover_inst_isDetached(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
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

func (x gen_NSMenuItem) InitWithTitle_action_keyEquivalent_asNSMenuItem(
	string core.NSStringRef,
	selector objc.Selector,
	charCode core.NSStringRef,
) (
	r0 NSMenuItem,
) {
	ret := C.NSMenuItem_inst_initWithTitle_action_keyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		selector.SelectorAddress(),
		objc.RefPointer(charCode),
	)
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) Init_asNSMenuItem() (
	r0 NSMenuItem,
) {
	ret := C.NSMenuItem_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) IsEnabled() (
	r0 bool,
) {
	ret := C.NSMenuItem_inst_isEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSMenuItem) SetEnabled(
	value bool,
) {
	C.NSMenuItem_inst_setEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSMenuItem) IsHidden() (
	r0 bool,
) {
	ret := C.NSMenuItem_inst_isHidden(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSMenuItem) SetHidden(
	value bool,
) {
	C.NSMenuItem_inst_setHidden(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSMenuItem) IsHiddenOrHasHiddenAncestor() (
	r0 bool,
) {
	ret := C.NSMenuItem_inst_isHiddenOrHasHiddenAncestor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSMenuItem) Target() (
	r0 objc.Object,
) {
	ret := C.NSMenuItem_inst_target(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) SetTarget(
	value objc.Ref,
) {
	C.NSMenuItem_inst_setTarget(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenuItem) Action() (
	r0 objc.Selector,
) {
	ret := C.NSMenuItem_inst_action(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.SelectorAt(ret)
	return
}

func (x gen_NSMenuItem) SetAction(
	value objc.Selector,
) {
	C.NSMenuItem_inst_setAction(
		unsafe.Pointer(x.Pointer()),
		value.SelectorAddress(),
	)
	return
}

func (x gen_NSMenuItem) Title() (
	r0 core.NSString,
) {
	ret := C.NSMenuItem_inst_title(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) SetTitle(
	value core.NSStringRef,
) {
	C.NSMenuItem_inst_setTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenuItem) AttributedTitle() (
	r0 core.NSAttributedString,
) {
	ret := C.NSMenuItem_inst_attributedTitle(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) SetAttributedTitle(
	value core.NSAttributedStringRef,
) {
	C.NSMenuItem_inst_setAttributedTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenuItem) Tag() (
	r0 core.NSInteger,
) {
	ret := C.NSMenuItem_inst_tag(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSMenuItem) SetTag(
	value core.NSInteger,
) {
	C.NSMenuItem_inst_setTag(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)
	return
}

func (x gen_NSMenuItem) State() (
	r0 core.NSInteger,
) {
	ret := C.NSMenuItem_inst_state(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSMenuItem) SetState(
	value core.NSInteger,
) {
	C.NSMenuItem_inst_setState(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)
	return
}

func (x gen_NSMenuItem) Image() (
	r0 NSImage,
) {
	ret := C.NSMenuItem_inst_image(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) SetImage(
	value NSImageRef,
) {
	C.NSMenuItem_inst_setImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenuItem) OnStateImage() (
	r0 NSImage,
) {
	ret := C.NSMenuItem_inst_onStateImage(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) SetOnStateImage(
	value NSImageRef,
) {
	C.NSMenuItem_inst_setOnStateImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenuItem) OffStateImage() (
	r0 NSImage,
) {
	ret := C.NSMenuItem_inst_offStateImage(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) SetOffStateImage(
	value NSImageRef,
) {
	C.NSMenuItem_inst_setOffStateImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenuItem) MixedStateImage() (
	r0 NSImage,
) {
	ret := C.NSMenuItem_inst_mixedStateImage(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) SetMixedStateImage(
	value NSImageRef,
) {
	C.NSMenuItem_inst_setMixedStateImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenuItem) Submenu() (
	r0 NSMenu,
) {
	ret := C.NSMenuItem_inst_submenu(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSMenu_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) SetSubmenu(
	value NSMenuRef,
) {
	C.NSMenuItem_inst_setSubmenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenuItem) HasSubmenu() (
	r0 bool,
) {
	ret := C.NSMenuItem_inst_hasSubmenu(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSMenuItem) ParentItem() (
	r0 NSMenuItem,
) {
	ret := C.NSMenuItem_inst_parentItem(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) IsSeparatorItem() (
	r0 bool,
) {
	ret := C.NSMenuItem_inst_isSeparatorItem(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSMenuItem) Menu() (
	r0 NSMenu,
) {
	ret := C.NSMenuItem_inst_menu(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSMenu_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) SetMenu(
	value NSMenuRef,
) {
	C.NSMenuItem_inst_setMenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenuItem) KeyEquivalent() (
	r0 core.NSString,
) {
	ret := C.NSMenuItem_inst_keyEquivalent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) SetKeyEquivalent(
	value core.NSStringRef,
) {
	C.NSMenuItem_inst_setKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenuItem) UserKeyEquivalent() (
	r0 core.NSString,
) {
	ret := C.NSMenuItem_inst_userKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) IsAlternate() (
	r0 bool,
) {
	ret := C.NSMenuItem_inst_isAlternate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSMenuItem) SetAlternate(
	value bool,
) {
	C.NSMenuItem_inst_setAlternate(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSMenuItem) IndentationLevel() (
	r0 core.NSInteger,
) {
	ret := C.NSMenuItem_inst_indentationLevel(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSMenuItem) SetIndentationLevel(
	value core.NSInteger,
) {
	C.NSMenuItem_inst_setIndentationLevel(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)
	return
}

func (x gen_NSMenuItem) ToolTip() (
	r0 core.NSString,
) {
	ret := C.NSMenuItem_inst_toolTip(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) SetToolTip(
	value core.NSStringRef,
) {
	C.NSMenuItem_inst_setToolTip(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenuItem) RepresentedObject() (
	r0 objc.Object,
) {
	ret := C.NSMenuItem_inst_representedObject(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) SetRepresentedObject(
	value objc.Ref,
) {
	C.NSMenuItem_inst_setRepresentedObject(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenuItem) View() (
	r0 NSView,
) {
	ret := C.NSMenuItem_inst_view(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSMenuItem) SetView(
	value NSViewRef,
) {
	C.NSMenuItem_inst_setView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSMenuItem) IsHighlighted() (
	r0 bool,
) {
	ret := C.NSMenuItem_inst_isHighlighted(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSMenuItem) AllowsAutomaticKeyEquivalentLocalization() (
	r0 bool,
) {
	ret := C.NSMenuItem_inst_allowsAutomaticKeyEquivalentLocalization(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSMenuItem) SetAllowsAutomaticKeyEquivalentLocalization(
	value bool,
) {
	C.NSMenuItem_inst_setAllowsAutomaticKeyEquivalentLocalization(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSMenuItem) AllowsAutomaticKeyEquivalentMirroring() (
	r0 bool,
) {
	ret := C.NSMenuItem_inst_allowsAutomaticKeyEquivalentMirroring(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSMenuItem) SetAllowsAutomaticKeyEquivalentMirroring(
	value bool,
) {
	C.NSMenuItem_inst_setAllowsAutomaticKeyEquivalentMirroring(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSMenuItem) AllowsKeyEquivalentWhenHidden() (
	r0 bool,
) {
	ret := C.NSMenuItem_inst_allowsKeyEquivalentWhenHidden(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

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

func (x gen_NSRunningApplication) Hide() (
	r0 bool,
) {
	ret := C.NSRunningApplication_inst_hide(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSRunningApplication) Unhide() (
	r0 bool,
) {
	ret := C.NSRunningApplication_inst_unhide(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSRunningApplication) ForceTerminate() (
	r0 bool,
) {
	ret := C.NSRunningApplication_inst_forceTerminate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSRunningApplication) Terminate() (
	r0 bool,
) {
	ret := C.NSRunningApplication_inst_terminate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSRunningApplication) Init_asNSRunningApplication() (
	r0 NSRunningApplication,
) {
	ret := C.NSRunningApplication_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSRunningApplication_fromPointer(ret)
	return
}

func (x gen_NSRunningApplication) IsActive() (
	r0 bool,
) {
	ret := C.NSRunningApplication_inst_isActive(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSRunningApplication) ActivationPolicy() (
	r0 core.NSInteger,
) {
	ret := C.NSRunningApplication_inst_activationPolicy(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSRunningApplication) IsHidden() (
	r0 bool,
) {
	ret := C.NSRunningApplication_inst_isHidden(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSRunningApplication) LocalizedName() (
	r0 core.NSString,
) {
	ret := C.NSRunningApplication_inst_localizedName(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSRunningApplication) Icon() (
	r0 NSImage,
) {
	ret := C.NSRunningApplication_inst_icon(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSRunningApplication) BundleIdentifier() (
	r0 core.NSString,
) {
	ret := C.NSRunningApplication_inst_bundleIdentifier(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSRunningApplication) BundleURL() (
	r0 core.NSURL,
) {
	ret := C.NSRunningApplication_inst_bundleURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSRunningApplication) ExecutableArchitecture() (
	r0 core.NSInteger,
) {
	ret := C.NSRunningApplication_inst_executableArchitecture(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSRunningApplication) ExecutableURL() (
	r0 core.NSURL,
) {
	ret := C.NSRunningApplication_inst_executableURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSRunningApplication) IsFinishedLaunching() (
	r0 bool,
) {
	ret := C.NSRunningApplication_inst_isFinishedLaunching(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSRunningApplication) OwnsMenuBar() (
	r0 bool,
) {
	ret := C.NSRunningApplication_inst_ownsMenuBar(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSRunningApplication) IsTerminated() (
	r0 bool,
) {
	ret := C.NSRunningApplication_inst_isTerminated(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
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

func (x gen_NSScreen) ConvertRectFromBacking(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSScreen_inst_convertRectFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSScreen) ConvertRectToBacking(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSScreen_inst_convertRectToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSScreen) Init_asNSScreen() (
	r0 NSScreen,
) {
	ret := C.NSScreen_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSScreen_fromPointer(ret)
	return
}

func (x gen_NSScreen) Frame() (
	r0 core.NSRect,
) {
	ret := C.NSScreen_inst_frame(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSScreen) DeviceDescription() (
	r0 core.NSDictionary,
) {
	ret := C.NSScreen_inst_deviceDescription(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSScreen) VisibleFrame() (
	r0 core.NSRect,
) {
	ret := C.NSScreen_inst_visibleFrame(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSScreen) BackingScaleFactor() (
	r0 core.CGFloat,
) {
	ret := C.NSScreen_inst_backingScaleFactor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSScreen) MaximumPotentialExtendedDynamicRangeColorComponentValue() (
	r0 core.CGFloat,
) {
	ret := C.NSScreen_inst_maximumPotentialExtendedDynamicRangeColorComponentValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSScreen) MaximumExtendedDynamicRangeColorComponentValue() (
	r0 core.CGFloat,
) {
	ret := C.NSScreen_inst_maximumExtendedDynamicRangeColorComponentValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSScreen) MaximumReferenceExtendedDynamicRangeColorComponentValue() (
	r0 core.CGFloat,
) {
	ret := C.NSScreen_inst_maximumReferenceExtendedDynamicRangeColorComponentValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSScreen) LocalizedName() (
	r0 core.NSString,
) {
	ret := C.NSScreen_inst_localizedName(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSScreen) MaximumFramesPerSecond() (
	r0 core.NSInteger,
) {
	ret := C.NSScreen_inst_maximumFramesPerSecond(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
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

func (x gen_NSStatusBar) StatusItemWithLength(
	length core.CGFloat,
) (
	r0 NSStatusItem,
) {
	ret := C.NSStatusBar_inst_statusItemWithLength(
		unsafe.Pointer(x.Pointer()),
		C.double(length),
	)
	r0 = NSStatusItem_fromPointer(ret)
	return
}

func (x gen_NSStatusBar) RemoveStatusItem(
	item NSStatusItemRef,
) {
	C.NSStatusBar_inst_removeStatusItem(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
	)
	return
}

func (x gen_NSStatusBar) Init_asNSStatusBar() (
	r0 NSStatusBar,
) {
	ret := C.NSStatusBar_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSStatusBar_fromPointer(ret)
	return
}

func (x gen_NSStatusBar) IsVertical() (
	r0 bool,
) {
	ret := C.NSStatusBar_inst_isVertical(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSStatusBar) Thickness() (
	r0 core.CGFloat,
) {
	ret := C.NSStatusBar_inst_thickness(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
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

func (x gen_NSStatusBarButton) Init_asNSStatusBarButton() (
	r0 NSStatusBarButton,
) {
	ret := C.NSStatusBarButton_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSStatusBarButton_fromPointer(ret)
	return
}

func (x gen_NSStatusBarButton) AppearsDisabled() (
	r0 bool,
) {
	ret := C.NSStatusBarButton_inst_appearsDisabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

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

func (x gen_NSStatusItem) Init_asNSStatusItem() (
	r0 NSStatusItem,
) {
	ret := C.NSStatusItem_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSStatusItem_fromPointer(ret)
	return
}

func (x gen_NSStatusItem) StatusBar() (
	r0 NSStatusBar,
) {
	ret := C.NSStatusItem_inst_statusBar(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSStatusBar_fromPointer(ret)
	return
}

func (x gen_NSStatusItem) Button() (
	r0 NSStatusBarButton,
) {
	ret := C.NSStatusItem_inst_button(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSStatusBarButton_fromPointer(ret)
	return
}

func (x gen_NSStatusItem) Menu() (
	r0 NSMenu,
) {
	ret := C.NSStatusItem_inst_menu(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSMenu_fromPointer(ret)
	return
}

func (x gen_NSStatusItem) SetMenu(
	value NSMenuRef,
) {
	C.NSStatusItem_inst_setMenu(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSStatusItem) IsVisible() (
	r0 bool,
) {
	ret := C.NSStatusItem_inst_isVisible(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSStatusItem) SetVisible(
	value bool,
) {
	C.NSStatusItem_inst_setVisible(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSStatusItem) Length() (
	r0 core.CGFloat,
) {
	ret := C.NSStatusItem_inst_length(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

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

func (x gen_NSText) InitWithFrame_asNSText(
	frameRect core.NSRect,
) (
	r0 NSText,
) {
	ret := C.NSText_inst_initWithFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)
	r0 = NSText_fromPointer(ret)
	return
}

func (x gen_NSText) ToggleRuler(
	sender objc.Ref,
) {
	C.NSText_inst_toggleRuler(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) SelectAll(
	sender objc.Ref,
) {
	C.NSText_inst_selectAll(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Copy(
	sender objc.Ref,
) {
	C.NSText_inst_copy(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Cut(
	sender objc.Ref,
) {
	C.NSText_inst_cut(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Paste(
	sender objc.Ref,
) {
	C.NSText_inst_paste(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) CopyFont(
	sender objc.Ref,
) {
	C.NSText_inst_copyFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) PasteFont(
	sender objc.Ref,
) {
	C.NSText_inst_pasteFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) CopyRuler(
	sender objc.Ref,
) {
	C.NSText_inst_copyRuler(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) PasteRuler(
	sender objc.Ref,
) {
	C.NSText_inst_pasteRuler(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Delete(
	sender objc.Ref,
) {
	C.NSText_inst_delete(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) ChangeFont(
	sender objc.Ref,
) {
	C.NSText_inst_changeFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) AlignCenter(
	sender objc.Ref,
) {
	C.NSText_inst_alignCenter(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) AlignLeft(
	sender objc.Ref,
) {
	C.NSText_inst_alignLeft(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) AlignRight(
	sender objc.Ref,
) {
	C.NSText_inst_alignRight(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Superscript(
	sender objc.Ref,
) {
	C.NSText_inst_superscript(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Subscript(
	sender objc.Ref,
) {
	C.NSText_inst_subscript(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Unscript(
	sender objc.Ref,
) {
	C.NSText_inst_unscript(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Underline(
	sender objc.Ref,
) {
	C.NSText_inst_underline(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) ReadRTFDFromFile(
	path core.NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSText_inst_readRTFDFromFile(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSText) WriteRTFDToFile_atomically(
	path core.NSStringRef,
	flag bool,
) (
	r0 bool,
) {
	ret := C.NSText_inst_writeRTFDToFile_atomically(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(flag),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSText) CheckSpelling(
	sender objc.Ref,
) {
	C.NSText_inst_checkSpelling(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) ShowGuessPanel(
	sender objc.Ref,
) {
	C.NSText_inst_showGuessPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) SizeToFit() {
	C.NSText_inst_sizeToFit(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSText) Init_asNSText() (
	r0 NSText,
) {
	ret := C.NSText_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSText_fromPointer(ret)
	return
}

func (x gen_NSText) String() (
	r0 core.NSString,
) {
	ret := C.NSText_inst_string(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSText) SetString(
	value core.NSStringRef,
) {
	C.NSText_inst_setString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSText) BackgroundColor() (
	r0 NSColor,
) {
	ret := C.NSText_inst_backgroundColor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSText) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSText_inst_setBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSText) DrawsBackground() (
	r0 bool,
) {
	ret := C.NSText_inst_drawsBackground(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSText) SetDrawsBackground(
	value bool,
) {
	C.NSText_inst_setDrawsBackground(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSText) IsEditable() (
	r0 bool,
) {
	ret := C.NSText_inst_isEditable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSText) SetEditable(
	value bool,
) {
	C.NSText_inst_setEditable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSText) IsSelectable() (
	r0 bool,
) {
	ret := C.NSText_inst_isSelectable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSText) SetSelectable(
	value bool,
) {
	C.NSText_inst_setSelectable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSText) IsFieldEditor() (
	r0 bool,
) {
	ret := C.NSText_inst_isFieldEditor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSText) SetFieldEditor(
	value bool,
) {
	C.NSText_inst_setFieldEditor(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSText) IsRichText() (
	r0 bool,
) {
	ret := C.NSText_inst_isRichText(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSText) SetRichText(
	value bool,
) {
	C.NSText_inst_setRichText(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSText) ImportsGraphics() (
	r0 bool,
) {
	ret := C.NSText_inst_importsGraphics(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSText) SetImportsGraphics(
	value bool,
) {
	C.NSText_inst_setImportsGraphics(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSText) UsesFontPanel() (
	r0 bool,
) {
	ret := C.NSText_inst_usesFontPanel(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSText) SetUsesFontPanel(
	value bool,
) {
	C.NSText_inst_setUsesFontPanel(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSText) IsRulerVisible() (
	r0 bool,
) {
	ret := C.NSText_inst_isRulerVisible(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSText) Font() (
	r0 NSFont,
) {
	ret := C.NSText_inst_font(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func (x gen_NSText) SetFont(
	value NSFontRef,
) {
	C.NSText_inst_setFont(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSText) TextColor() (
	r0 NSColor,
) {
	ret := C.NSText_inst_textColor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSText) SetTextColor(
	value NSColorRef,
) {
	C.NSText_inst_setTextColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSText) MaxSize() (
	r0 core.NSSize,
) {
	ret := C.NSText_inst_maxSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSText) SetMaxSize(
	value core.NSSize,
) {
	C.NSText_inst_setMaxSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSText) MinSize() (
	r0 core.NSSize,
) {
	ret := C.NSText_inst_minSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSText) SetMinSize(
	value core.NSSize,
) {
	C.NSText_inst_setMinSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSText) IsVerticallyResizable() (
	r0 bool,
) {
	ret := C.NSText_inst_isVerticallyResizable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSText) SetVerticallyResizable(
	value bool,
) {
	C.NSText_inst_setVerticallyResizable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSText) IsHorizontallyResizable() (
	r0 bool,
) {
	ret := C.NSText_inst_isHorizontallyResizable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSText) SetHorizontallyResizable(
	value bool,
) {
	C.NSText_inst_setHorizontallyResizable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSText) Delegate() (
	r0 objc.Object,
) {
	ret := C.NSText_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

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

func (x gen_NSTextField) SelectText(
	sender objc.Ref,
) {
	C.NSTextField_inst_selectText(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextField) TextShouldBeginEditing(
	textObject NSTextRef,
) (
	r0 bool,
) {
	ret := C.NSTextField_inst_textShouldBeginEditing(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(textObject),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextField) TextShouldEndEditing(
	textObject NSTextRef,
) (
	r0 bool,
) {
	ret := C.NSTextField_inst_textShouldEndEditing(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(textObject),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextField) Init_asNSTextField() (
	r0 NSTextField,
) {
	ret := C.NSTextField_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSTextField_fromPointer(ret)
	return
}

func (x gen_NSTextField) IsSelectable() (
	r0 bool,
) {
	ret := C.NSTextField_inst_isSelectable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextField) SetSelectable(
	value bool,
) {
	C.NSTextField_inst_setSelectable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextField) IsEditable() (
	r0 bool,
) {
	ret := C.NSTextField_inst_isEditable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextField) SetEditable(
	value bool,
) {
	C.NSTextField_inst_setEditable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextField) AllowsEditingTextAttributes() (
	r0 bool,
) {
	ret := C.NSTextField_inst_allowsEditingTextAttributes(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextField) SetAllowsEditingTextAttributes(
	value bool,
) {
	C.NSTextField_inst_setAllowsEditingTextAttributes(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextField) ImportsGraphics() (
	r0 bool,
) {
	ret := C.NSTextField_inst_importsGraphics(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextField) SetImportsGraphics(
	value bool,
) {
	C.NSTextField_inst_setImportsGraphics(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextField) PlaceholderString() (
	r0 core.NSString,
) {
	ret := C.NSTextField_inst_placeholderString(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSTextField) SetPlaceholderString(
	value core.NSStringRef,
) {
	C.NSTextField_inst_setPlaceholderString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextField) PlaceholderAttributedString() (
	r0 core.NSAttributedString,
) {
	ret := C.NSTextField_inst_placeholderAttributedString(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSTextField) SetPlaceholderAttributedString(
	value core.NSAttributedStringRef,
) {
	C.NSTextField_inst_setPlaceholderAttributedString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextField) AllowsDefaultTighteningForTruncation() (
	r0 bool,
) {
	ret := C.NSTextField_inst_allowsDefaultTighteningForTruncation(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextField) SetAllowsDefaultTighteningForTruncation(
	value bool,
) {
	C.NSTextField_inst_setAllowsDefaultTighteningForTruncation(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextField) MaximumNumberOfLines() (
	r0 core.NSInteger,
) {
	ret := C.NSTextField_inst_maximumNumberOfLines(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSTextField) SetMaximumNumberOfLines(
	value core.NSInteger,
) {
	C.NSTextField_inst_setMaximumNumberOfLines(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)
	return
}

func (x gen_NSTextField) PreferredMaxLayoutWidth() (
	r0 core.CGFloat,
) {
	ret := C.NSTextField_inst_preferredMaxLayoutWidth(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSTextField) SetPreferredMaxLayoutWidth(
	value core.CGFloat,
) {
	C.NSTextField_inst_setPreferredMaxLayoutWidth(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_NSTextField) TextColor() (
	r0 NSColor,
) {
	ret := C.NSTextField_inst_textColor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSTextField) SetTextColor(
	value NSColorRef,
) {
	C.NSTextField_inst_setTextColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextField) BackgroundColor() (
	r0 NSColor,
) {
	ret := C.NSTextField_inst_backgroundColor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSTextField) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSTextField_inst_setBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextField) DrawsBackground() (
	r0 bool,
) {
	ret := C.NSTextField_inst_drawsBackground(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextField) SetDrawsBackground(
	value bool,
) {
	C.NSTextField_inst_setDrawsBackground(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextField) IsBezeled() (
	r0 bool,
) {
	ret := C.NSTextField_inst_isBezeled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextField) SetBezeled(
	value bool,
) {
	C.NSTextField_inst_setBezeled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextField) IsBordered() (
	r0 bool,
) {
	ret := C.NSTextField_inst_isBordered(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextField) SetBordered(
	value bool,
) {
	C.NSTextField_inst_setBordered(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextField) AcceptsFirstResponder() (
	r0 bool,
) {
	ret := C.NSTextField_inst_acceptsFirstResponder(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextField) AllowsCharacterPickerTouchBarItem() (
	r0 bool,
) {
	ret := C.NSTextField_inst_allowsCharacterPickerTouchBarItem(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextField) SetAllowsCharacterPickerTouchBarItem(
	value bool,
) {
	C.NSTextField_inst_setAllowsCharacterPickerTouchBarItem(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextField) IsAutomaticTextCompletionEnabled() (
	r0 bool,
) {
	ret := C.NSTextField_inst_isAutomaticTextCompletionEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextField) SetAutomaticTextCompletionEnabled(
	value bool,
) {
	C.NSTextField_inst_setAutomaticTextCompletionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextField) Delegate() (
	r0 objc.Object,
) {
	ret := C.NSTextField_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

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

func (x gen_NSTextContainer) InitWithSize_asNSTextContainer(
	size core.NSSize,
) (
	r0 NSTextContainer,
) {
	ret := C.NSTextContainer_inst_initWithSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	r0 = NSTextContainer_fromPointer(ret)
	return
}

func (x gen_NSTextContainer) ReplaceLayoutManager(
	newLayoutManager NSLayoutManagerRef,
) {
	C.NSTextContainer_inst_replaceLayoutManager(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newLayoutManager),
	)
	return
}

func (x gen_NSTextContainer) Init_asNSTextContainer() (
	r0 NSTextContainer,
) {
	ret := C.NSTextContainer_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSTextContainer_fromPointer(ret)
	return
}

func (x gen_NSTextContainer) LayoutManager() (
	r0 NSLayoutManager,
) {
	ret := C.NSTextContainer_inst_layoutManager(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSLayoutManager_fromPointer(ret)
	return
}

func (x gen_NSTextContainer) SetLayoutManager(
	value NSLayoutManagerRef,
) {
	C.NSTextContainer_inst_setLayoutManager(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextContainer) TextView() (
	r0 NSTextView,
) {
	ret := C.NSTextContainer_inst_textView(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSTextView_fromPointer(ret)
	return
}

func (x gen_NSTextContainer) SetTextView(
	value NSTextViewRef,
) {
	C.NSTextContainer_inst_setTextView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextContainer) Size() (
	r0 core.NSSize,
) {
	ret := C.NSTextContainer_inst_size(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSTextContainer) SetSize(
	value core.NSSize,
) {
	C.NSTextContainer_inst_setSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSTextContainer) ExclusionPaths() (
	r0 core.NSArray,
) {
	ret := C.NSTextContainer_inst_exclusionPaths(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSTextContainer) SetExclusionPaths(
	value core.NSArrayRef,
) {
	C.NSTextContainer_inst_setExclusionPaths(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextContainer) WidthTracksTextView() (
	r0 bool,
) {
	ret := C.NSTextContainer_inst_widthTracksTextView(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextContainer) SetWidthTracksTextView(
	value bool,
) {
	C.NSTextContainer_inst_setWidthTracksTextView(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextContainer) HeightTracksTextView() (
	r0 bool,
) {
	ret := C.NSTextContainer_inst_heightTracksTextView(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextContainer) SetHeightTracksTextView(
	value bool,
) {
	C.NSTextContainer_inst_setHeightTracksTextView(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextContainer) MaximumNumberOfLines() (
	r0 core.NSUInteger,
) {
	ret := C.NSTextContainer_inst_maximumNumberOfLines(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSTextContainer) SetMaximumNumberOfLines(
	value core.NSUInteger,
) {
	C.NSTextContainer_inst_setMaximumNumberOfLines(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)
	return
}

func (x gen_NSTextContainer) LineFragmentPadding() (
	r0 core.CGFloat,
) {
	ret := C.NSTextContainer_inst_lineFragmentPadding(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSTextContainer) SetLineFragmentPadding(
	value core.CGFloat,
) {
	C.NSTextContainer_inst_setLineFragmentPadding(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_NSTextContainer) IsSimpleRectangularTextContainer() (
	r0 bool,
) {
	ret := C.NSTextContainer_inst_isSimpleRectangularTextContainer(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
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

func (x gen_NSViewController) LoadView() {
	C.NSViewController_inst_loadView(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

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

func (x gen_NSViewController) CommitEditing() (
	r0 bool,
) {
	ret := C.NSViewController_inst_commitEditing(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSViewController) DiscardEditing() {
	C.NSViewController_inst_discardEditing(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSViewController) DismissController(
	sender objc.Ref,
) {
	C.NSViewController_inst_dismissController(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSViewController) ViewDidLoad() {
	C.NSViewController_inst_viewDidLoad(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSViewController) ViewWillAppear() {
	C.NSViewController_inst_viewWillAppear(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSViewController) ViewDidAppear() {
	C.NSViewController_inst_viewDidAppear(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSViewController) ViewWillDisappear() {
	C.NSViewController_inst_viewWillDisappear(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSViewController) ViewDidDisappear() {
	C.NSViewController_inst_viewDidDisappear(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSViewController) UpdateViewConstraints() {
	C.NSViewController_inst_updateViewConstraints(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSViewController) ViewWillLayout() {
	C.NSViewController_inst_viewWillLayout(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSViewController) ViewDidLayout() {
	C.NSViewController_inst_viewDidLayout(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSViewController) AddChildViewController(
	childViewController NSViewControllerRef,
) {
	C.NSViewController_inst_addChildViewController(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(childViewController),
	)
	return
}

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

func (x gen_NSViewController) RemoveChildViewControllerAtIndex(
	index core.NSInteger,
) {
	C.NSViewController_inst_removeChildViewControllerAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)
	return
}

func (x gen_NSViewController) RemoveFromParentViewController() {
	C.NSViewController_inst_removeFromParentViewController(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSViewController) PreferredContentSizeDidChangeForViewController(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_preferredContentSizeDidChangeForViewController(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)
	return
}

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

func (x gen_NSViewController) DismissViewController(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_dismissViewController(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)
	return
}

func (x gen_NSViewController) PresentViewControllerAsModalWindow(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_presentViewControllerAsModalWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)
	return
}

func (x gen_NSViewController) PresentViewControllerAsSheet(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_presentViewControllerAsSheet(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)
	return
}

func (x gen_NSViewController) ViewWillTransitionToSize(
	newSize core.NSSize,
) {
	C.NSViewController_inst_viewWillTransitionToSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&newSize)),
	)
	return
}

func (x gen_NSViewController) Init_asNSViewController() (
	r0 NSViewController,
) {
	ret := C.NSViewController_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSViewController_fromPointer(ret)
	return
}

func (x gen_NSViewController) RepresentedObject() (
	r0 objc.Object,
) {
	ret := C.NSViewController_inst_representedObject(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSViewController) SetRepresentedObject(
	value objc.Ref,
) {
	C.NSViewController_inst_setRepresentedObject(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSViewController) NibBundle() (
	r0 NSBundle,
) {
	ret := C.NSViewController_inst_nibBundle(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSBundle_fromPointer(ret)
	return
}

func (x gen_NSViewController) View() (
	r0 NSView,
) {
	ret := C.NSViewController_inst_view(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSViewController) SetView(
	value NSViewRef,
) {
	C.NSViewController_inst_setView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSViewController) Title() (
	r0 core.NSString,
) {
	ret := C.NSViewController_inst_title(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSViewController) SetTitle(
	value core.NSStringRef,
) {
	C.NSViewController_inst_setTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSViewController) IsViewLoaded() (
	r0 bool,
) {
	ret := C.NSViewController_inst_isViewLoaded(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSViewController) PreferredContentSize() (
	r0 core.NSSize,
) {
	ret := C.NSViewController_inst_preferredContentSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSViewController) SetPreferredContentSize(
	value core.NSSize,
) {
	C.NSViewController_inst_setPreferredContentSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSViewController) ChildViewControllers() (
	r0 core.NSArray,
) {
	ret := C.NSViewController_inst_childViewControllers(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSViewController) SetChildViewControllers(
	value core.NSArrayRef,
) {
	C.NSViewController_inst_setChildViewControllers(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSViewController) ParentViewController() (
	r0 NSViewController,
) {
	ret := C.NSViewController_inst_parentViewController(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSViewController_fromPointer(ret)
	return
}

func (x gen_NSViewController) PresentedViewControllers() (
	r0 core.NSArray,
) {
	ret := C.NSViewController_inst_presentedViewControllers(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSViewController) PresentingViewController() (
	r0 NSViewController,
) {
	ret := C.NSViewController_inst_presentingViewController(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSViewController_fromPointer(ret)
	return
}

func (x gen_NSViewController) PreferredScreenOrigin() (
	r0 core.NSPoint,
) {
	ret := C.NSViewController_inst_preferredScreenOrigin(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSViewController) SetPreferredScreenOrigin(
	value core.NSPoint,
) {
	C.NSViewController_inst_setPreferredScreenOrigin(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSViewController) PreferredMaximumSize() (
	r0 core.NSSize,
) {
	ret := C.NSViewController_inst_preferredMaximumSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSViewController) PreferredMinimumSize() (
	r0 core.NSSize,
) {
	ret := C.NSViewController_inst_preferredMinimumSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSViewController) SourceItemView() (
	r0 NSView,
) {
	ret := C.NSViewController_inst_sourceItemView(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSView_fromPointer(ret)
	return
}

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

func (x gen_NSVisualEffectView) ViewDidMoveToWindow() {
	C.NSVisualEffectView_inst_viewDidMoveToWindow(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSVisualEffectView) ViewWillMoveToWindow(
	newWindow NSWindowRef,
) {
	C.NSVisualEffectView_inst_viewWillMoveToWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newWindow),
	)
	return
}

func (x gen_NSVisualEffectView) Init_asNSVisualEffectView() (
	r0 NSVisualEffectView,
) {
	ret := C.NSVisualEffectView_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSVisualEffectView_fromPointer(ret)
	return
}

func (x gen_NSVisualEffectView) IsEmphasized() (
	r0 bool,
) {
	ret := C.NSVisualEffectView_inst_isEmphasized(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSVisualEffectView) SetEmphasized(
	value bool,
) {
	C.NSVisualEffectView_inst_setEmphasized(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSVisualEffectView) MaskImage() (
	r0 NSImage,
) {
	ret := C.NSVisualEffectView_inst_maskImage(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

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

func (x gen_NSWindow) InitWithContentRect_styleMask_backing_defer_asNSWindow(
	contentRect core.NSRect,
	style core.NSUInteger,
	backingStoreType core.NSUInteger,
	flag bool,
) (
	r0 NSWindow,
) {
	ret := C.NSWindow_inst_initWithContentRect_styleMask_backing_defer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentRect)),
		C.ulong(style),
		C.ulong(backingStoreType),
		convertToObjCBool(flag),
	)
	r0 = NSWindow_fromPointer(ret)
	return
}

func (x gen_NSWindow) InitWithContentRect_styleMask_backing_defer_screen_asNSWindow(
	contentRect core.NSRect,
	style core.NSUInteger,
	backingStoreType core.NSUInteger,
	flag bool,
	screen NSScreenRef,
) (
	r0 NSWindow,
) {
	ret := C.NSWindow_inst_initWithContentRect_styleMask_backing_defer_screen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentRect)),
		C.ulong(style),
		C.ulong(backingStoreType),
		convertToObjCBool(flag),
		objc.RefPointer(screen),
	)
	r0 = NSWindow_fromPointer(ret)
	return
}

func (x gen_NSWindow) ToggleFullScreen(
	sender objc.Ref,
) {
	C.NSWindow_inst_toggleFullScreen(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) SetDynamicDepthLimit(
	flag bool,
) {
	C.NSWindow_inst_setDynamicDepthLimit(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)
	return
}

func (x gen_NSWindow) InvalidateShadow() {
	C.NSWindow_inst_invalidateShadow(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) ContentRectForFrameRect(
	frameRect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_contentRectForFrameRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) FrameRectForContentRect(
	contentRect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_frameRectForContentRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentRect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) EndSheet(
	sheetWindow NSWindowRef,
) {
	C.NSWindow_inst_endSheet(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sheetWindow),
	)
	return
}

func (x gen_NSWindow) SetFrameOrigin(
	point core.NSPoint,
) {
	C.NSWindow_inst_setFrameOrigin(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	return
}

func (x gen_NSWindow) SetFrameTopLeftPoint(
	point core.NSPoint,
) {
	C.NSWindow_inst_setFrameTopLeftPoint(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	return
}

func (x gen_NSWindow) ConstrainFrameRect_toScreen(
	frameRect core.NSRect,
	screen NSScreenRef,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_constrainFrameRect_toScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
		objc.RefPointer(screen),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) CascadeTopLeftFromPoint(
	topLeftPoint core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSWindow_inst_cascadeTopLeftFromPoint(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&topLeftPoint)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

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

func (x gen_NSWindow) PerformZoom(
	sender objc.Ref,
) {
	C.NSWindow_inst_performZoom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) Zoom(
	sender objc.Ref,
) {
	C.NSWindow_inst_zoom(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) SetContentSize(
	size core.NSSize,
) {
	C.NSWindow_inst_setContentSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	return
}

func (x gen_NSWindow) OrderOut(
	sender objc.Ref,
) {
	C.NSWindow_inst_orderOut(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) OrderBack(
	sender objc.Ref,
) {
	C.NSWindow_inst_orderBack(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) OrderFront(
	sender objc.Ref,
) {
	C.NSWindow_inst_orderFront(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) OrderFrontRegardless() {
	C.NSWindow_inst_orderFrontRegardless(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

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

func (x gen_NSWindow) MakeKeyWindow() {
	C.NSWindow_inst_makeKeyWindow(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) MakeKeyAndOrderFront(
	sender objc.Ref,
) {
	C.NSWindow_inst_makeKeyAndOrderFront(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) BecomeKeyWindow() {
	C.NSWindow_inst_becomeKeyWindow(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) ResignKeyWindow() {
	C.NSWindow_inst_resignKeyWindow(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) MakeMainWindow() {
	C.NSWindow_inst_makeMainWindow(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) BecomeMainWindow() {
	C.NSWindow_inst_becomeMainWindow(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) ResignMainWindow() {
	C.NSWindow_inst_resignMainWindow(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) ToggleToolbarShown(
	sender objc.Ref,
) {
	C.NSWindow_inst_toggleToolbarShown(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) RunToolbarCustomizationPalette(
	sender objc.Ref,
) {
	C.NSWindow_inst_runToolbarCustomizationPalette(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

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

func (x gen_NSWindow) RemoveChildWindow(
	childWin NSWindowRef,
) {
	C.NSWindow_inst_removeChildWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(childWin),
	)
	return
}

func (x gen_NSWindow) EnableKeyEquivalentForDefaultButtonCell() {
	C.NSWindow_inst_enableKeyEquivalentForDefaultButtonCell(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) DisableKeyEquivalentForDefaultButtonCell() {
	C.NSWindow_inst_disableKeyEquivalentForDefaultButtonCell(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) FieldEditor_forObject(
	createFlag bool,
	object objc.Ref,
) (
	r0 NSText,
) {
	ret := C.NSWindow_inst_fieldEditor_forObject(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(createFlag),
		objc.RefPointer(object),
	)
	r0 = NSText_fromPointer(ret)
	return
}

func (x gen_NSWindow) EndEditingFor(
	object objc.Ref,
) {
	C.NSWindow_inst_endEditingFor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)
	return
}

func (x gen_NSWindow) EnableCursorRects() {
	C.NSWindow_inst_enableCursorRects(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) DisableCursorRects() {
	C.NSWindow_inst_disableCursorRects(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) DiscardCursorRects() {
	C.NSWindow_inst_discardCursorRects(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) InvalidateCursorRectsForView(
	view NSViewRef,
) {
	C.NSWindow_inst_invalidateCursorRectsForView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)
	return
}

func (x gen_NSWindow) ResetCursorRects() {
	C.NSWindow_inst_resetCursorRects(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) RemoveTitlebarAccessoryViewControllerAtIndex(
	index core.NSInteger,
) {
	C.NSWindow_inst_removeTitlebarAccessoryViewControllerAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)
	return
}

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

func (x gen_NSWindow) MergeAllWindows(
	sender objc.Ref,
) {
	C.NSWindow_inst_mergeAllWindows(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) SelectNextTab(
	sender objc.Ref,
) {
	C.NSWindow_inst_selectNextTab(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) SelectPreviousTab(
	sender objc.Ref,
) {
	C.NSWindow_inst_selectPreviousTab(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) MoveTabToNewWindow(
	sender objc.Ref,
) {
	C.NSWindow_inst_moveTabToNewWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) ToggleTabBar(
	sender objc.Ref,
) {
	C.NSWindow_inst_toggleTabBar(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) ToggleTabOverview(
	sender objc.Ref,
) {
	C.NSWindow_inst_toggleTabOverview(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

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

func (x gen_NSWindow) SendEvent(
	event NSEventRef,
) {
	C.NSWindow_inst_sendEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	return
}

func (x gen_NSWindow) TryToPerform_with(
	action objc.Selector,
	object objc.Ref,
) (
	r0 bool,
) {
	ret := C.NSWindow_inst_tryToPerform_with(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(object),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SelectKeyViewPrecedingView(
	view NSViewRef,
) {
	C.NSWindow_inst_selectKeyViewPrecedingView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)
	return
}

func (x gen_NSWindow) SelectKeyViewFollowingView(
	view NSViewRef,
) {
	C.NSWindow_inst_selectKeyViewFollowingView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)
	return
}

func (x gen_NSWindow) SelectPreviousKeyView(
	sender objc.Ref,
) {
	C.NSWindow_inst_selectPreviousKeyView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) SelectNextKeyView(
	sender objc.Ref,
) {
	C.NSWindow_inst_selectNextKeyView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) RecalculateKeyViewLoop() {
	C.NSWindow_inst_recalculateKeyViewLoop(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) PerformWindowDragWithEvent(
	event NSEventRef,
) {
	C.NSWindow_inst_performWindowDragWithEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	return
}

func (x gen_NSWindow) DisableSnapshotRestoration() {
	C.NSWindow_inst_disableSnapshotRestoration(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) EnableSnapshotRestoration() {
	C.NSWindow_inst_enableSnapshotRestoration(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) Display() {
	C.NSWindow_inst_display(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) DisplayIfNeeded() {
	C.NSWindow_inst_displayIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) DisableScreenUpdatesUntilFlush() {
	C.NSWindow_inst_disableScreenUpdatesUntilFlush(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) Update() {
	C.NSWindow_inst_update(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

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

func (x gen_NSWindow) RegisterForDraggedTypes(
	newTypes core.NSArrayRef,
) {
	C.NSWindow_inst_registerForDraggedTypes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newTypes),
	)
	return
}

func (x gen_NSWindow) UnregisterDraggedTypes() {
	C.NSWindow_inst_unregisterDraggedTypes(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) ConvertRectFromBacking(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_convertRectFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) ConvertRectFromScreen(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_convertRectFromScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) ConvertPointFromBacking(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSWindow_inst_convertPointFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) ConvertPointFromScreen(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSWindow_inst_convertPointFromScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) ConvertRectToBacking(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_convertRectToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) ConvertRectToScreen(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_convertRectToScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) ConvertPointToBacking(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSWindow_inst_convertPointToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) ConvertPointToScreen(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSWindow_inst_convertPointToScreen(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) SetTitleWithRepresentedFilename(
	filename core.NSStringRef,
) {
	C.NSWindow_inst_setTitleWithRepresentedFilename(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(filename),
	)
	return
}

func (x gen_NSWindow) Center() {
	C.NSWindow_inst_center(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) PerformClose(
	sender objc.Ref,
) {
	C.NSWindow_inst_performClose(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) Close() {
	C.NSWindow_inst_close(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) PerformMiniaturize(
	sender objc.Ref,
) {
	C.NSWindow_inst_performMiniaturize(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) Miniaturize(
	sender objc.Ref,
) {
	C.NSWindow_inst_miniaturize(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) Deminiaturize(
	sender objc.Ref,
) {
	C.NSWindow_inst_deminiaturize(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) Print(
	sender objc.Ref,
) {
	C.NSWindow_inst_print(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) DataWithEPSInsideRect(
	rect core.NSRect,
) (
	r0 core.NSData,
) {
	ret := C.NSWindow_inst_dataWithEPSInsideRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = core.NSData_fromPointer(ret)
	return
}

func (x gen_NSWindow) DataWithPDFInsideRect(
	rect core.NSRect,
) (
	r0 core.NSData,
) {
	ret := C.NSWindow_inst_dataWithPDFInsideRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = core.NSData_fromPointer(ret)
	return
}

func (x gen_NSWindow) UpdateConstraintsIfNeeded() {
	C.NSWindow_inst_updateConstraintsIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) LayoutIfNeeded() {
	C.NSWindow_inst_layoutIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWindow) VisualizeConstraints(
	constraints core.NSArrayRef,
) {
	C.NSWindow_inst_visualizeConstraints(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(constraints),
	)
	return
}

func (x gen_NSWindow) SetIsMiniaturized(
	flag bool,
) {
	C.NSWindow_inst_setIsMiniaturized(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)
	return
}

func (x gen_NSWindow) SetIsVisible(
	flag bool,
) {
	C.NSWindow_inst_setIsVisible(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)
	return
}

func (x gen_NSWindow) SetIsZoomed(
	flag bool,
) {
	C.NSWindow_inst_setIsZoomed(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)
	return
}

func (x gen_NSWindow) Init_asNSWindow() (
	r0 NSWindow,
) {
	ret := C.NSWindow_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSWindow_fromPointer(ret)
	return
}

func (x gen_NSWindow) Delegate() (
	r0 objc.Object,
) {
	ret := C.NSWindow_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSWindow) SetDelegate(
	value objc.Ref,
) {
	C.NSWindow_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSWindow) ContentViewController() (
	r0 NSViewController,
) {
	ret := C.NSWindow_inst_contentViewController(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSViewController_fromPointer(ret)
	return
}

func (x gen_NSWindow) SetContentViewController(
	value NSViewControllerRef,
) {
	C.NSWindow_inst_setContentViewController(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSWindow) ContentView() (
	r0 NSView,
) {
	ret := C.NSWindow_inst_contentView(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSWindow) SetContentView(
	value NSViewRef,
) {
	C.NSWindow_inst_setContentView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSWindow) StyleMask() (
	r0 core.NSUInteger,
) {
	ret := C.NSWindow_inst_styleMask(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSWindow) SetStyleMask(
	value core.NSUInteger,
) {
	C.NSWindow_inst_setStyleMask(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)
	return
}

func (x gen_NSWindow) WorksWhenModal() (
	r0 bool,
) {
	ret := C.NSWindow_inst_worksWhenModal(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) AlphaValue() (
	r0 core.CGFloat,
) {
	ret := C.NSWindow_inst_alphaValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSWindow) SetAlphaValue(
	value core.CGFloat,
) {
	C.NSWindow_inst_setAlphaValue(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_NSWindow) BackgroundColor() (
	r0 NSColor,
) {
	ret := C.NSWindow_inst_backgroundColor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSWindow) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSWindow_inst_setBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSWindow) CanHide() (
	r0 bool,
) {
	ret := C.NSWindow_inst_canHide(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetCanHide(
	value bool,
) {
	C.NSWindow_inst_setCanHide(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) IsOnActiveSpace() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isOnActiveSpace(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) HidesOnDeactivate() (
	r0 bool,
) {
	ret := C.NSWindow_inst_hidesOnDeactivate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetHidesOnDeactivate(
	value bool,
) {
	C.NSWindow_inst_setHidesOnDeactivate(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) CollectionBehavior() (
	r0 core.NSUInteger,
) {
	ret := C.NSWindow_inst_collectionBehavior(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSWindow) SetCollectionBehavior(
	value core.NSUInteger,
) {
	C.NSWindow_inst_setCollectionBehavior(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)
	return
}

func (x gen_NSWindow) IsOpaque() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isOpaque(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetOpaque(
	value bool,
) {
	C.NSWindow_inst_setOpaque(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) HasShadow() (
	r0 bool,
) {
	ret := C.NSWindow_inst_hasShadow(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetHasShadow(
	value bool,
) {
	C.NSWindow_inst_setHasShadow(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) PreventsApplicationTerminationWhenModal() (
	r0 bool,
) {
	ret := C.NSWindow_inst_preventsApplicationTerminationWhenModal(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetPreventsApplicationTerminationWhenModal(
	value bool,
) {
	C.NSWindow_inst_setPreventsApplicationTerminationWhenModal(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) HasDynamicDepthLimit() (
	r0 bool,
) {
	ret := C.NSWindow_inst_hasDynamicDepthLimit(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) WindowNumber() (
	r0 core.NSInteger,
) {
	ret := C.NSWindow_inst_windowNumber(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSWindow) DeviceDescription() (
	r0 core.NSDictionary,
) {
	ret := C.NSWindow_inst_deviceDescription(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSWindow) CanBecomeVisibleWithoutLogin() (
	r0 bool,
) {
	ret := C.NSWindow_inst_canBecomeVisibleWithoutLogin(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetCanBecomeVisibleWithoutLogin(
	value bool,
) {
	C.NSWindow_inst_setCanBecomeVisibleWithoutLogin(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) BackingType() (
	r0 core.NSUInteger,
) {
	ret := C.NSWindow_inst_backingType(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSWindow) SetBackingType(
	value core.NSUInteger,
) {
	C.NSWindow_inst_setBackingType(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)
	return
}

func (x gen_NSWindow) AttachedSheet() (
	r0 NSWindow,
) {
	ret := C.NSWindow_inst_attachedSheet(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSWindow_fromPointer(ret)
	return
}

func (x gen_NSWindow) IsSheet() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isSheet(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SheetParent() (
	r0 NSWindow,
) {
	ret := C.NSWindow_inst_sheetParent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSWindow_fromPointer(ret)
	return
}

func (x gen_NSWindow) Sheets() (
	r0 core.NSArray,
) {
	ret := C.NSWindow_inst_sheets(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSWindow) Frame() (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_frame(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) AspectRatio() (
	r0 core.NSSize,
) {
	ret := C.NSWindow_inst_aspectRatio(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) SetAspectRatio(
	value core.NSSize,
) {
	C.NSWindow_inst_setAspectRatio(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSWindow) MinSize() (
	r0 core.NSSize,
) {
	ret := C.NSWindow_inst_minSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) SetMinSize(
	value core.NSSize,
) {
	C.NSWindow_inst_setMinSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSWindow) MaxSize() (
	r0 core.NSSize,
) {
	ret := C.NSWindow_inst_maxSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) SetMaxSize(
	value core.NSSize,
) {
	C.NSWindow_inst_setMaxSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSWindow) IsZoomed() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isZoomed(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) ResizeIncrements() (
	r0 core.NSSize,
) {
	ret := C.NSWindow_inst_resizeIncrements(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) SetResizeIncrements(
	value core.NSSize,
) {
	C.NSWindow_inst_setResizeIncrements(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSWindow) PreservesContentDuringLiveResize() (
	r0 bool,
) {
	ret := C.NSWindow_inst_preservesContentDuringLiveResize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetPreservesContentDuringLiveResize(
	value bool,
) {
	C.NSWindow_inst_setPreservesContentDuringLiveResize(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) InLiveResize() (
	r0 bool,
) {
	ret := C.NSWindow_inst_inLiveResize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) ContentAspectRatio() (
	r0 core.NSSize,
) {
	ret := C.NSWindow_inst_contentAspectRatio(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) SetContentAspectRatio(
	value core.NSSize,
) {
	C.NSWindow_inst_setContentAspectRatio(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSWindow) ContentMinSize() (
	r0 core.NSSize,
) {
	ret := C.NSWindow_inst_contentMinSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) SetContentMinSize(
	value core.NSSize,
) {
	C.NSWindow_inst_setContentMinSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSWindow) ContentMaxSize() (
	r0 core.NSSize,
) {
	ret := C.NSWindow_inst_contentMaxSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) SetContentMaxSize(
	value core.NSSize,
) {
	C.NSWindow_inst_setContentMaxSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSWindow) ContentResizeIncrements() (
	r0 core.NSSize,
) {
	ret := C.NSWindow_inst_contentResizeIncrements(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) SetContentResizeIncrements(
	value core.NSSize,
) {
	C.NSWindow_inst_setContentResizeIncrements(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSWindow) ContentLayoutGuide() (
	r0 objc.Object,
) {
	ret := C.NSWindow_inst_contentLayoutGuide(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSWindow) ContentLayoutRect() (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_contentLayoutRect(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) MaxFullScreenContentSize() (
	r0 core.NSSize,
) {
	ret := C.NSWindow_inst_maxFullScreenContentSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) SetMaxFullScreenContentSize(
	value core.NSSize,
) {
	C.NSWindow_inst_setMaxFullScreenContentSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSWindow) MinFullScreenContentSize() (
	r0 core.NSSize,
) {
	ret := C.NSWindow_inst_minFullScreenContentSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) SetMinFullScreenContentSize(
	value core.NSSize,
) {
	C.NSWindow_inst_setMinFullScreenContentSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSWindow) Level() (
	r0 core.NSInteger,
) {
	ret := C.NSWindow_inst_level(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSWindow) SetLevel(
	value core.NSInteger,
) {
	C.NSWindow_inst_setLevel(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)
	return
}

func (x gen_NSWindow) IsVisible() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isVisible(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) IsKeyWindow() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isKeyWindow(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) CanBecomeKeyWindow() (
	r0 bool,
) {
	ret := C.NSWindow_inst_canBecomeKeyWindow(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) IsMainWindow() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isMainWindow(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) CanBecomeMainWindow() (
	r0 bool,
) {
	ret := C.NSWindow_inst_canBecomeMainWindow(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) ChildWindows() (
	r0 core.NSArray,
) {
	ret := C.NSWindow_inst_childWindows(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSWindow) ParentWindow() (
	r0 NSWindow,
) {
	ret := C.NSWindow_inst_parentWindow(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSWindow_fromPointer(ret)
	return
}

func (x gen_NSWindow) SetParentWindow(
	value NSWindowRef,
) {
	C.NSWindow_inst_setParentWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSWindow) IsExcludedFromWindowsMenu() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isExcludedFromWindowsMenu(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetExcludedFromWindowsMenu(
	value bool,
) {
	C.NSWindow_inst_setExcludedFromWindowsMenu(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) AreCursorRectsEnabled() (
	r0 bool,
) {
	ret := C.NSWindow_inst_areCursorRectsEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) ShowsToolbarButton() (
	r0 bool,
) {
	ret := C.NSWindow_inst_showsToolbarButton(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetShowsToolbarButton(
	value bool,
) {
	C.NSWindow_inst_setShowsToolbarButton(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) TitlebarAppearsTransparent() (
	r0 bool,
) {
	ret := C.NSWindow_inst_titlebarAppearsTransparent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetTitlebarAppearsTransparent(
	value bool,
) {
	C.NSWindow_inst_setTitlebarAppearsTransparent(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) TitlebarAccessoryViewControllers() (
	r0 core.NSArray,
) {
	ret := C.NSWindow_inst_titlebarAccessoryViewControllers(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSWindow) SetTitlebarAccessoryViewControllers(
	value core.NSArrayRef,
) {
	C.NSWindow_inst_setTitlebarAccessoryViewControllers(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSWindow) TabbedWindows() (
	r0 core.NSArray,
) {
	ret := C.NSWindow_inst_tabbedWindows(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSWindow) AllowsToolTipsWhenApplicationIsInactive() (
	r0 bool,
) {
	ret := C.NSWindow_inst_allowsToolTipsWhenApplicationIsInactive(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetAllowsToolTipsWhenApplicationIsInactive(
	value bool,
) {
	C.NSWindow_inst_setAllowsToolTipsWhenApplicationIsInactive(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) CurrentEvent() (
	r0 NSEvent,
) {
	ret := C.NSWindow_inst_currentEvent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSEvent_fromPointer(ret)
	return
}

func (x gen_NSWindow) InitialFirstResponder() (
	r0 NSView,
) {
	ret := C.NSWindow_inst_initialFirstResponder(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSWindow) SetInitialFirstResponder(
	value NSViewRef,
) {
	C.NSWindow_inst_setInitialFirstResponder(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSWindow) AutorecalculatesKeyViewLoop() (
	r0 bool,
) {
	ret := C.NSWindow_inst_autorecalculatesKeyViewLoop(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetAutorecalculatesKeyViewLoop(
	value bool,
) {
	C.NSWindow_inst_setAutorecalculatesKeyViewLoop(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) AcceptsMouseMovedEvents() (
	r0 bool,
) {
	ret := C.NSWindow_inst_acceptsMouseMovedEvents(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetAcceptsMouseMovedEvents(
	value bool,
) {
	C.NSWindow_inst_setAcceptsMouseMovedEvents(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) IgnoresMouseEvents() (
	r0 bool,
) {
	ret := C.NSWindow_inst_ignoresMouseEvents(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetIgnoresMouseEvents(
	value bool,
) {
	C.NSWindow_inst_setIgnoresMouseEvents(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) MouseLocationOutsideOfEventStream() (
	r0 core.NSPoint,
) {
	ret := C.NSWindow_inst_mouseLocationOutsideOfEventStream(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) IsRestorable() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isRestorable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetRestorable(
	value bool,
) {
	C.NSWindow_inst_setRestorable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) ViewsNeedDisplay() (
	r0 bool,
) {
	ret := C.NSWindow_inst_viewsNeedDisplay(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetViewsNeedDisplay(
	value bool,
) {
	C.NSWindow_inst_setViewsNeedDisplay(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) AllowsConcurrentViewDrawing() (
	r0 bool,
) {
	ret := C.NSWindow_inst_allowsConcurrentViewDrawing(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetAllowsConcurrentViewDrawing(
	value bool,
) {
	C.NSWindow_inst_setAllowsConcurrentViewDrawing(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) IsDocumentEdited() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isDocumentEdited(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetDocumentEdited(
	value bool,
) {
	C.NSWindow_inst_setDocumentEdited(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) BackingScaleFactor() (
	r0 core.CGFloat,
) {
	ret := C.NSWindow_inst_backingScaleFactor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSWindow) Title() (
	r0 core.NSString,
) {
	ret := C.NSWindow_inst_title(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSWindow) SetTitle(
	value core.NSStringRef,
) {
	C.NSWindow_inst_setTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSWindow) Subtitle() (
	r0 core.NSString,
) {
	ret := C.NSWindow_inst_subtitle(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSWindow) SetSubtitle(
	value core.NSStringRef,
) {
	C.NSWindow_inst_setSubtitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSWindow) TitleVisibility() (
	r0 core.NSInteger,
) {
	ret := C.NSWindow_inst_titleVisibility(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSWindow) SetTitleVisibility(
	value core.NSInteger,
) {
	C.NSWindow_inst_setTitleVisibility(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)
	return
}

func (x gen_NSWindow) RepresentedFilename() (
	r0 core.NSString,
) {
	ret := C.NSWindow_inst_representedFilename(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSWindow) SetRepresentedFilename(
	value core.NSStringRef,
) {
	C.NSWindow_inst_setRepresentedFilename(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSWindow) RepresentedURL() (
	r0 core.NSURL,
) {
	ret := C.NSWindow_inst_representedURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSWindow) SetRepresentedURL(
	value core.NSURLRef,
) {
	C.NSWindow_inst_setRepresentedURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSWindow) Screen() (
	r0 NSScreen,
) {
	ret := C.NSWindow_inst_screen(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSScreen_fromPointer(ret)
	return
}

func (x gen_NSWindow) DeepestScreen() (
	r0 NSScreen,
) {
	ret := C.NSWindow_inst_deepestScreen(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSScreen_fromPointer(ret)
	return
}

func (x gen_NSWindow) DisplaysWhenScreenProfileChanges() (
	r0 bool,
) {
	ret := C.NSWindow_inst_displaysWhenScreenProfileChanges(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetDisplaysWhenScreenProfileChanges(
	value bool,
) {
	C.NSWindow_inst_setDisplaysWhenScreenProfileChanges(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) IsMovableByWindowBackground() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isMovableByWindowBackground(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetMovableByWindowBackground(
	value bool,
) {
	C.NSWindow_inst_setMovableByWindowBackground(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) IsMovable() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isMovable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetMovable(
	value bool,
) {
	C.NSWindow_inst_setMovable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) IsReleasedWhenClosed() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isReleasedWhenClosed(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SetReleasedWhenClosed(
	value bool,
) {
	C.NSWindow_inst_setReleasedWhenClosed(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSWindow) IsMiniaturized() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isMiniaturized(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) MiniwindowImage() (
	r0 NSImage,
) {
	ret := C.NSWindow_inst_miniwindowImage(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSWindow) SetMiniwindowImage(
	value NSImageRef,
) {
	C.NSWindow_inst_setMiniwindowImage(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSWindow) MiniwindowTitle() (
	r0 core.NSString,
) {
	ret := C.NSWindow_inst_miniwindowTitle(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSWindow) SetMiniwindowTitle(
	value core.NSStringRef,
) {
	C.NSWindow_inst_setMiniwindowTitle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSWindow) HasCloseBox() (
	r0 bool,
) {
	ret := C.NSWindow_inst_hasCloseBox(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) HasTitleBar() (
	r0 bool,
) {
	ret := C.NSWindow_inst_hasTitleBar(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) IsModalPanel() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isModalPanel(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) IsFloatingPanel() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isFloatingPanel(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) IsZoomable() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isZoomable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) IsResizable() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isResizable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) IsMiniaturizable() (
	r0 bool,
) {
	ret := C.NSWindow_inst_isMiniaturizable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) OrderedIndex() (
	r0 core.NSInteger,
) {
	ret := C.NSWindow_inst_orderedIndex(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

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

func (x gen_NSWorkspace) OpenURL(
	url core.NSURLRef,
) (
	r0 bool,
) {
	ret := C.NSWorkspace_inst_openURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWorkspace) HideOtherApplications() {
	C.NSWorkspace_inst_hideOtherApplications(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSWorkspace) ActivateFileViewerSelectingURLs(
	fileURLs core.NSArrayRef,
) {
	C.NSWorkspace_inst_activateFileViewerSelectingURLs(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fileURLs),
	)
	return
}

func (x gen_NSWorkspace) SelectFile_inFileViewerRootedAtPath(
	fullPath core.NSStringRef,
	rootFullPath core.NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSWorkspace_inst_selectFile_inFileViewerRootedAtPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fullPath),
		objc.RefPointer(rootFullPath),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWorkspace) URLForApplicationWithBundleIdentifier(
	bundleIdentifier core.NSStringRef,
) (
	r0 core.NSURL,
) {
	ret := C.NSWorkspace_inst_URLForApplicationWithBundleIdentifier(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(bundleIdentifier),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSWorkspace) URLForApplicationToOpenURL(
	url core.NSURLRef,
) (
	r0 core.NSURL,
) {
	ret := C.NSWorkspace_inst_URLForApplicationToOpenURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSWorkspace) IsFilePackageAtPath(
	fullPath core.NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSWorkspace_inst_isFilePackageAtPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fullPath),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWorkspace) IconForFile(
	fullPath core.NSStringRef,
) (
	r0 NSImage,
) {
	ret := C.NSWorkspace_inst_iconForFile(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fullPath),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSWorkspace) IconForFiles(
	fullPaths core.NSArrayRef,
) (
	r0 NSImage,
) {
	ret := C.NSWorkspace_inst_iconForFiles(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fullPaths),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSWorkspace) UnmountAndEjectDeviceAtPath(
	path core.NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSWorkspace_inst_unmountAndEjectDeviceAtPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWorkspace) DesktopImageURLForScreen(
	screen NSScreenRef,
) (
	r0 core.NSURL,
) {
	ret := C.NSWorkspace_inst_desktopImageURLForScreen(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(screen),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSWorkspace) DesktopImageOptionsForScreen(
	screen NSScreenRef,
) (
	r0 core.NSDictionary,
) {
	ret := C.NSWorkspace_inst_desktopImageOptionsForScreen(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(screen),
	)
	r0 = core.NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSWorkspace) ShowSearchResultsForQueryString(
	queryString core.NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSWorkspace_inst_showSearchResultsForQueryString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(queryString),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWorkspace) NoteFileSystemChanged(
	path core.NSStringRef,
) {
	C.NSWorkspace_inst_noteFileSystemChanged(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)
	return
}

func (x gen_NSWorkspace) ExtendPowerOffBy(
	requested core.NSInteger,
) (
	r0 core.NSInteger,
) {
	ret := C.NSWorkspace_inst_extendPowerOffBy(
		unsafe.Pointer(x.Pointer()),
		C.long(requested),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSWorkspace) URLsForApplicationsToOpenURL(
	url core.NSURLRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSWorkspace_inst_URLsForApplicationsToOpenURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSWorkspace) URLsForApplicationsWithBundleIdentifier(
	bundleIdentifier core.NSStringRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSWorkspace_inst_URLsForApplicationsWithBundleIdentifier(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(bundleIdentifier),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSWorkspace) Init_asNSWorkspace() (
	r0 NSWorkspace,
) {
	ret := C.NSWorkspace_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSWorkspace_fromPointer(ret)
	return
}

func (x gen_NSWorkspace) FrontmostApplication() (
	r0 NSRunningApplication,
) {
	ret := C.NSWorkspace_inst_frontmostApplication(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSRunningApplication_fromPointer(ret)
	return
}

func (x gen_NSWorkspace) RunningApplications() (
	r0 core.NSArray,
) {
	ret := C.NSWorkspace_inst_runningApplications(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSWorkspace) MenuBarOwningApplication() (
	r0 NSRunningApplication,
) {
	ret := C.NSWorkspace_inst_menuBarOwningApplication(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSRunningApplication_fromPointer(ret)
	return
}

func (x gen_NSWorkspace) FileLabels() (
	r0 core.NSArray,
) {
	ret := C.NSWorkspace_inst_fileLabels(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSWorkspace) FileLabelColors() (
	r0 core.NSArray,
) {
	ret := C.NSWorkspace_inst_fileLabelColors(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSWorkspace) AccessibilityDisplayShouldDifferentiateWithoutColor() (
	r0 bool,
) {
	ret := C.NSWorkspace_inst_accessibilityDisplayShouldDifferentiateWithoutColor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWorkspace) AccessibilityDisplayShouldIncreaseContrast() (
	r0 bool,
) {
	ret := C.NSWorkspace_inst_accessibilityDisplayShouldIncreaseContrast(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWorkspace) AccessibilityDisplayShouldReduceTransparency() (
	r0 bool,
) {
	ret := C.NSWorkspace_inst_accessibilityDisplayShouldReduceTransparency(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWorkspace) AccessibilityDisplayShouldInvertColors() (
	r0 bool,
) {
	ret := C.NSWorkspace_inst_accessibilityDisplayShouldInvertColors(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWorkspace) AccessibilityDisplayShouldReduceMotion() (
	r0 bool,
) {
	ret := C.NSWorkspace_inst_accessibilityDisplayShouldReduceMotion(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWorkspace) IsSwitchControlEnabled() (
	r0 bool,
) {
	ret := C.NSWorkspace_inst_isSwitchControlEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWorkspace) IsVoiceOverEnabled() (
	r0 bool,
) {
	ret := C.NSWorkspace_inst_isVoiceOverEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
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

func (x gen_NSColor) BlendedColorWithFraction_ofColor(
	fraction core.CGFloat,
	color NSColorRef,
) (
	r0 NSColor,
) {
	ret := C.NSColor_inst_blendedColorWithFraction_ofColor(
		unsafe.Pointer(x.Pointer()),
		C.double(fraction),
		objc.RefPointer(color),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSColor) ColorWithAlphaComponent(
	alpha core.CGFloat,
) (
	r0 NSColor,
) {
	ret := C.NSColor_inst_colorWithAlphaComponent(
		unsafe.Pointer(x.Pointer()),
		C.double(alpha),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSColor) HighlightWithLevel(
	val core.CGFloat,
) (
	r0 NSColor,
) {
	ret := C.NSColor_inst_highlightWithLevel(
		unsafe.Pointer(x.Pointer()),
		C.double(val),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSColor) ShadowWithLevel(
	val core.CGFloat,
) (
	r0 NSColor,
) {
	ret := C.NSColor_inst_shadowWithLevel(
		unsafe.Pointer(x.Pointer()),
		C.double(val),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSColor) WriteToPasteboard(
	pasteBoard NSPasteboardRef,
) {
	C.NSColor_inst_writeToPasteboard(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pasteBoard),
	)
	return
}

func (x gen_NSColor) DrawSwatchInRect(
	rect core.NSRect,
) {
	C.NSColor_inst_drawSwatchInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	return
}

func (x gen_NSColor) Set() {
	C.NSColor_inst_set(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSColor) SetFill() {
	C.NSColor_inst_setFill(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSColor) SetStroke() {
	C.NSColor_inst_setStroke(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSColor) Init_asNSColor() (
	r0 NSColor,
) {
	ret := C.NSColor_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSColor) NumberOfComponents() (
	r0 core.NSInteger,
) {
	ret := C.NSColor_inst_numberOfComponents(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSColor) AlphaComponent() (
	r0 core.CGFloat,
) {
	ret := C.NSColor_inst_alphaComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSColor) WhiteComponent() (
	r0 core.CGFloat,
) {
	ret := C.NSColor_inst_whiteComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSColor) RedComponent() (
	r0 core.CGFloat,
) {
	ret := C.NSColor_inst_redComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSColor) GreenComponent() (
	r0 core.CGFloat,
) {
	ret := C.NSColor_inst_greenComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSColor) BlueComponent() (
	r0 core.CGFloat,
) {
	ret := C.NSColor_inst_blueComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSColor) CyanComponent() (
	r0 core.CGFloat,
) {
	ret := C.NSColor_inst_cyanComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSColor) MagentaComponent() (
	r0 core.CGFloat,
) {
	ret := C.NSColor_inst_magentaComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSColor) YellowComponent() (
	r0 core.CGFloat,
) {
	ret := C.NSColor_inst_yellowComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSColor) BlackComponent() (
	r0 core.CGFloat,
) {
	ret := C.NSColor_inst_blackComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSColor) HueComponent() (
	r0 core.CGFloat,
) {
	ret := C.NSColor_inst_hueComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSColor) SaturationComponent() (
	r0 core.CGFloat,
) {
	ret := C.NSColor_inst_saturationComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSColor) BrightnessComponent() (
	r0 core.CGFloat,
) {
	ret := C.NSColor_inst_brightnessComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSColor) LocalizedCatalogNameComponent() (
	r0 core.NSString,
) {
	ret := C.NSColor_inst_localizedCatalogNameComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSColor) LocalizedColorNameComponent() (
	r0 core.NSString,
) {
	ret := C.NSColor_inst_localizedColorNameComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
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

func (x gen_NSTextView) InitWithFrame_textContainer_asNSTextView(
	frameRect core.NSRect,
	container NSTextContainerRef,
) (
	r0 NSTextView,
) {
	ret := C.NSTextView_inst_initWithFrame_textContainer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
		objc.RefPointer(container),
	)
	r0 = NSTextView_fromPointer(ret)
	return
}

func (x gen_NSTextView) InitWithFrame_asNSTextView(
	frameRect core.NSRect,
) (
	r0 NSTextView,
) {
	ret := C.NSTextView_inst_initWithFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)
	r0 = NSTextView_fromPointer(ret)
	return
}

func (x gen_NSTextView) ReplaceTextContainer(
	newContainer NSTextContainerRef,
) {
	C.NSTextView_inst_replaceTextContainer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newContainer),
	)
	return
}

func (x gen_NSTextView) InvalidateTextContainerOrigin() {
	C.NSTextView_inst_invalidateTextContainerOrigin(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSTextView) ChangeDocumentBackgroundColor(
	sender objc.Ref,
) {
	C.NSTextView_inst_changeDocumentBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

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

func (x gen_NSTextView) DrawViewBackgroundInRect(
	rect core.NSRect,
) {
	C.NSTextView_inst_drawViewBackgroundInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	return
}

func (x gen_NSTextView) SetConstrainedFrameSize(
	desiredSize core.NSSize,
) {
	C.NSTextView_inst_setConstrainedFrameSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&desiredSize)),
	)
	return
}

func (x gen_NSTextView) CleanUpAfterDragOperation() {
	C.NSTextView_inst_cleanUpAfterDragOperation(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSTextView) Outline(
	sender objc.Ref,
) {
	C.NSTextView_inst_outline(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleAutomaticQuoteSubstitution(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticQuoteSubstitution(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleAutomaticLinkDetection(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticLinkDetection(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleAutomaticTextCompletion(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticTextCompletion(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) UpdateInsertionPointStateAndRestartTimer(
	restartFlag bool,
) {
	C.NSTextView_inst_updateInsertionPointStateAndRestartTimer(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(restartFlag),
	)
	return
}

func (x gen_NSTextView) CharacterIndexForInsertionAtPoint(
	point core.NSPoint,
) (
	r0 core.NSUInteger,
) {
	ret := C.NSTextView_inst_characterIndexForInsertionAtPoint(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSTextView) UpdateCandidates() {
	C.NSTextView_inst_updateCandidates(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSTextView) ReadSelectionFromPasteboard(
	pboard NSPasteboardRef,
) (
	r0 bool,
) {
	ret := C.NSTextView_inst_readSelectionFromPasteboard(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pboard),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) WriteSelectionToPasteboard_types(
	pboard NSPasteboardRef,
	types core.NSArrayRef,
) (
	r0 bool,
) {
	ret := C.NSTextView_inst_writeSelectionToPasteboard_types(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pboard),
		objc.RefPointer(types),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) AlignJustified(
	sender objc.Ref,
) {
	C.NSTextView_inst_alignJustified(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ChangeAttributes(
	sender objc.Ref,
) {
	C.NSTextView_inst_changeAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ChangeColor(
	sender objc.Ref,
) {
	C.NSTextView_inst_changeColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) UseStandardKerning(
	sender objc.Ref,
) {
	C.NSTextView_inst_useStandardKerning(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) LowerBaseline(
	sender objc.Ref,
) {
	C.NSTextView_inst_lowerBaseline(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) RaiseBaseline(
	sender objc.Ref,
) {
	C.NSTextView_inst_raiseBaseline(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) TurnOffKerning(
	sender objc.Ref,
) {
	C.NSTextView_inst_turnOffKerning(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) LoosenKerning(
	sender objc.Ref,
) {
	C.NSTextView_inst_loosenKerning(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) TightenKerning(
	sender objc.Ref,
) {
	C.NSTextView_inst_tightenKerning(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) UseStandardLigatures(
	sender objc.Ref,
) {
	C.NSTextView_inst_useStandardLigatures(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) TurnOffLigatures(
	sender objc.Ref,
) {
	C.NSTextView_inst_turnOffLigatures(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) UseAllLigatures(
	sender objc.Ref,
) {
	C.NSTextView_inst_useAllLigatures(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

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

func (x gen_NSTextView) PasteAsPlainText(
	sender objc.Ref,
) {
	C.NSTextView_inst_pasteAsPlainText(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) PasteAsRichText(
	sender objc.Ref,
) {
	C.NSTextView_inst_pasteAsRichText(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) BreakUndoCoalescing() {
	C.NSTextView_inst_breakUndoCoalescing(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSTextView) UpdateFontPanel() {
	C.NSTextView_inst_updateFontPanel(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSTextView) UpdateRuler() {
	C.NSTextView_inst_updateRuler(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSTextView) UpdateDragTypeRegistration() {
	C.NSTextView_inst_updateDragTypeRegistration(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSTextView) ShouldChangeTextInRanges_replacementStrings(
	affectedRanges core.NSArrayRef,
	replacementStrings core.NSArrayRef,
) (
	r0 bool,
) {
	ret := C.NSTextView_inst_shouldChangeTextInRanges_replacementStrings(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(affectedRanges),
		objc.RefPointer(replacementStrings),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) DidChangeText() {
	C.NSTextView_inst_didChangeText(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSTextView) ToggleSmartInsertDelete(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleSmartInsertDelete(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleContinuousSpellChecking(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleContinuousSpellChecking(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleGrammarChecking(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleGrammarChecking(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) OrderFrontSharingServicePicker(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontSharingServicePicker(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) DragSelectionWithEvent_offset_slideBack(
	event NSEventRef,
	mouseOffset core.NSSize,
	slideBack bool,
) (
	r0 bool,
) {
	ret := C.NSTextView_inst_dragSelectionWithEvent_offset_slideBack(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
		*(*C.NSSize)(unsafe.Pointer(&mouseOffset)),
		convertToObjCBool(slideBack),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) StartSpeaking(
	sender objc.Ref,
) {
	C.NSTextView_inst_startSpeaking(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) StopSpeaking(
	sender objc.Ref,
) {
	C.NSTextView_inst_stopSpeaking(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) PerformFindPanelAction(
	sender objc.Ref,
) {
	C.NSTextView_inst_performFindPanelAction(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) OrderFrontLinkPanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontLinkPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) OrderFrontListPanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontListPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) OrderFrontSpacingPanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontSpacingPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) OrderFrontTablePanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontTablePanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) OrderFrontSubstitutionsPanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontSubstitutionsPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) Complete(
	sender objc.Ref,
) {
	C.NSTextView_inst_complete(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) CheckTextInDocument(
	sender objc.Ref,
) {
	C.NSTextView_inst_checkTextInDocument(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) CheckTextInSelection(
	sender objc.Ref,
) {
	C.NSTextView_inst_checkTextInSelection(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleAutomaticDashSubstitution(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticDashSubstitution(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleAutomaticDataDetection(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticDataDetection(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleAutomaticSpellingCorrection(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticSpellingCorrection(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleAutomaticTextReplacement(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticTextReplacement(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) UpdateQuickLookPreviewPanel() {
	C.NSTextView_inst_updateQuickLookPreviewPanel(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSTextView) ToggleQuickLookPreviewPanel(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleQuickLookPreviewPanel(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) QuickLookPreviewableItemsInRanges(
	ranges core.NSArrayRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSTextView_inst_quickLookPreviewableItemsInRanges(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ranges),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSTextView) ChangeLayoutOrientation(
	sender objc.Ref,
) {
	C.NSTextView_inst_changeLayoutOrientation(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) UpdateTextTouchBarItems() {
	C.NSTextView_inst_updateTextTouchBarItems(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSTextView) UpdateTouchBarItemIdentifiers() {
	C.NSTextView_inst_updateTouchBarItemIdentifiers(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSTextView) Init_asNSTextView() (
	r0 NSTextView,
) {
	ret := C.NSTextView_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSTextView_fromPointer(ret)
	return
}

func (x gen_NSTextView) Delegate() (
	r0 objc.Object,
) {
	ret := C.NSTextView_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSTextView) SetDelegate(
	value objc.Ref,
) {
	C.NSTextView_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextView) TextContainer() (
	r0 NSTextContainer,
) {
	ret := C.NSTextView_inst_textContainer(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSTextContainer_fromPointer(ret)
	return
}

func (x gen_NSTextView) SetTextContainer(
	value NSTextContainerRef,
) {
	C.NSTextView_inst_setTextContainer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextView) TextContainerInset() (
	r0 core.NSSize,
) {
	ret := C.NSTextView_inst_textContainerInset(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSTextView) SetTextContainerInset(
	value core.NSSize,
) {
	C.NSTextView_inst_setTextContainerInset(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSTextView) TextContainerOrigin() (
	r0 core.NSPoint,
) {
	ret := C.NSTextView_inst_textContainerOrigin(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSTextView) LayoutManager() (
	r0 NSLayoutManager,
) {
	ret := C.NSTextView_inst_layoutManager(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSLayoutManager_fromPointer(ret)
	return
}

func (x gen_NSTextView) BackgroundColor() (
	r0 NSColor,
) {
	ret := C.NSTextView_inst_backgroundColor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSTextView) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSTextView_inst_setBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextView) DrawsBackground() (
	r0 bool,
) {
	ret := C.NSTextView_inst_drawsBackground(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetDrawsBackground(
	value bool,
) {
	C.NSTextView_inst_setDrawsBackground(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) AllowsDocumentBackgroundColorChange() (
	r0 bool,
) {
	ret := C.NSTextView_inst_allowsDocumentBackgroundColorChange(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetAllowsDocumentBackgroundColorChange(
	value bool,
) {
	C.NSTextView_inst_setAllowsDocumentBackgroundColorChange(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) ShouldDrawInsertionPoint() (
	r0 bool,
) {
	ret := C.NSTextView_inst_shouldDrawInsertionPoint(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) AllowedInputSourceLocales() (
	r0 core.NSArray,
) {
	ret := C.NSTextView_inst_allowedInputSourceLocales(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSTextView) SetAllowedInputSourceLocales(
	value core.NSArrayRef,
) {
	C.NSTextView_inst_setAllowedInputSourceLocales(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextView) AllowsUndo() (
	r0 bool,
) {
	ret := C.NSTextView_inst_allowsUndo(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetAllowsUndo(
	value bool,
) {
	C.NSTextView_inst_setAllowsUndo(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) IsEditable() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isEditable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetEditable(
	value bool,
) {
	C.NSTextView_inst_setEditable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) IsSelectable() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isSelectable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetSelectable(
	value bool,
) {
	C.NSTextView_inst_setSelectable(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) IsFieldEditor() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isFieldEditor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetFieldEditor(
	value bool,
) {
	C.NSTextView_inst_setFieldEditor(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) IsRichText() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isRichText(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetRichText(
	value bool,
) {
	C.NSTextView_inst_setRichText(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) ImportsGraphics() (
	r0 bool,
) {
	ret := C.NSTextView_inst_importsGraphics(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetImportsGraphics(
	value bool,
) {
	C.NSTextView_inst_setImportsGraphics(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) AllowsImageEditing() (
	r0 bool,
) {
	ret := C.NSTextView_inst_allowsImageEditing(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetAllowsImageEditing(
	value bool,
) {
	C.NSTextView_inst_setAllowsImageEditing(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) IsAutomaticQuoteSubstitutionEnabled() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isAutomaticQuoteSubstitutionEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetAutomaticQuoteSubstitutionEnabled(
	value bool,
) {
	C.NSTextView_inst_setAutomaticQuoteSubstitutionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) IsAutomaticLinkDetectionEnabled() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isAutomaticLinkDetectionEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetAutomaticLinkDetectionEnabled(
	value bool,
) {
	C.NSTextView_inst_setAutomaticLinkDetectionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) DisplaysLinkToolTips() (
	r0 bool,
) {
	ret := C.NSTextView_inst_displaysLinkToolTips(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetDisplaysLinkToolTips(
	value bool,
) {
	C.NSTextView_inst_setDisplaysLinkToolTips(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) IsAutomaticTextCompletionEnabled() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isAutomaticTextCompletionEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetAutomaticTextCompletionEnabled(
	value bool,
) {
	C.NSTextView_inst_setAutomaticTextCompletionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) UsesAdaptiveColorMappingForDarkAppearance() (
	r0 bool,
) {
	ret := C.NSTextView_inst_usesAdaptiveColorMappingForDarkAppearance(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetUsesAdaptiveColorMappingForDarkAppearance(
	value bool,
) {
	C.NSTextView_inst_setUsesAdaptiveColorMappingForDarkAppearance(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) UsesRolloverButtonForSelection() (
	r0 bool,
) {
	ret := C.NSTextView_inst_usesRolloverButtonForSelection(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetUsesRolloverButtonForSelection(
	value bool,
) {
	C.NSTextView_inst_setUsesRolloverButtonForSelection(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) UsesRuler() (
	r0 bool,
) {
	ret := C.NSTextView_inst_usesRuler(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetUsesRuler(
	value bool,
) {
	C.NSTextView_inst_setUsesRuler(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) IsRulerVisible() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isRulerVisible(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetRulerVisible(
	value bool,
) {
	C.NSTextView_inst_setRulerVisible(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) UsesInspectorBar() (
	r0 bool,
) {
	ret := C.NSTextView_inst_usesInspectorBar(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetUsesInspectorBar(
	value bool,
) {
	C.NSTextView_inst_setUsesInspectorBar(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) SelectedRanges() (
	r0 core.NSArray,
) {
	ret := C.NSTextView_inst_selectedRanges(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSTextView) SetSelectedRanges(
	value core.NSArrayRef,
) {
	C.NSTextView_inst_setSelectedRanges(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextView) InsertionPointColor() (
	r0 NSColor,
) {
	ret := C.NSTextView_inst_insertionPointColor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSTextView) SetInsertionPointColor(
	value NSColorRef,
) {
	C.NSTextView_inst_setInsertionPointColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextView) SelectedTextAttributes() (
	r0 core.NSDictionary,
) {
	ret := C.NSTextView_inst_selectedTextAttributes(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSTextView) SetSelectedTextAttributes(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_setSelectedTextAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextView) MarkedTextAttributes() (
	r0 core.NSDictionary,
) {
	ret := C.NSTextView_inst_markedTextAttributes(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSTextView) SetMarkedTextAttributes(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_setMarkedTextAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextView) LinkTextAttributes() (
	r0 core.NSDictionary,
) {
	ret := C.NSTextView_inst_linkTextAttributes(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSTextView) SetLinkTextAttributes(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_setLinkTextAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextView) ReadablePasteboardTypes() (
	r0 core.NSArray,
) {
	ret := C.NSTextView_inst_readablePasteboardTypes(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSTextView) WritablePasteboardTypes() (
	r0 core.NSArray,
) {
	ret := C.NSTextView_inst_writablePasteboardTypes(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSTextView) TypingAttributes() (
	r0 core.NSDictionary,
) {
	ret := C.NSTextView_inst_typingAttributes(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSTextView) SetTypingAttributes(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_setTypingAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSTextView) IsCoalescingUndo() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isCoalescingUndo(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) AcceptableDragTypes() (
	r0 core.NSArray,
) {
	ret := C.NSTextView_inst_acceptableDragTypes(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSTextView) RangesForUserCharacterAttributeChange() (
	r0 core.NSArray,
) {
	ret := C.NSTextView_inst_rangesForUserCharacterAttributeChange(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSTextView) RangesForUserParagraphAttributeChange() (
	r0 core.NSArray,
) {
	ret := C.NSTextView_inst_rangesForUserParagraphAttributeChange(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSTextView) RangesForUserTextChange() (
	r0 core.NSArray,
) {
	ret := C.NSTextView_inst_rangesForUserTextChange(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSTextView) SmartInsertDeleteEnabled() (
	r0 bool,
) {
	ret := C.NSTextView_inst_smartInsertDeleteEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetSmartInsertDeleteEnabled(
	value bool,
) {
	C.NSTextView_inst_setSmartInsertDeleteEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) IsContinuousSpellCheckingEnabled() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isContinuousSpellCheckingEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetContinuousSpellCheckingEnabled(
	value bool,
) {
	C.NSTextView_inst_setContinuousSpellCheckingEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) SpellCheckerDocumentTag() (
	r0 core.NSInteger,
) {
	ret := C.NSTextView_inst_spellCheckerDocumentTag(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSTextView) IsGrammarCheckingEnabled() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isGrammarCheckingEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetGrammarCheckingEnabled(
	value bool,
) {
	C.NSTextView_inst_setGrammarCheckingEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) AcceptsGlyphInfo() (
	r0 bool,
) {
	ret := C.NSTextView_inst_acceptsGlyphInfo(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetAcceptsGlyphInfo(
	value bool,
) {
	C.NSTextView_inst_setAcceptsGlyphInfo(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) UsesFontPanel() (
	r0 bool,
) {
	ret := C.NSTextView_inst_usesFontPanel(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetUsesFontPanel(
	value bool,
) {
	C.NSTextView_inst_setUsesFontPanel(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) UsesFindPanel() (
	r0 bool,
) {
	ret := C.NSTextView_inst_usesFindPanel(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetUsesFindPanel(
	value bool,
) {
	C.NSTextView_inst_setUsesFindPanel(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) IsAutomaticDashSubstitutionEnabled() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isAutomaticDashSubstitutionEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetAutomaticDashSubstitutionEnabled(
	value bool,
) {
	C.NSTextView_inst_setAutomaticDashSubstitutionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) IsAutomaticDataDetectionEnabled() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isAutomaticDataDetectionEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetAutomaticDataDetectionEnabled(
	value bool,
) {
	C.NSTextView_inst_setAutomaticDataDetectionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) IsAutomaticSpellingCorrectionEnabled() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isAutomaticSpellingCorrectionEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetAutomaticSpellingCorrectionEnabled(
	value bool,
) {
	C.NSTextView_inst_setAutomaticSpellingCorrectionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) IsAutomaticTextReplacementEnabled() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isAutomaticTextReplacementEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetAutomaticTextReplacementEnabled(
	value bool,
) {
	C.NSTextView_inst_setAutomaticTextReplacementEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) UsesFindBar() (
	r0 bool,
) {
	ret := C.NSTextView_inst_usesFindBar(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetUsesFindBar(
	value bool,
) {
	C.NSTextView_inst_setUsesFindBar(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) IsIncrementalSearchingEnabled() (
	r0 bool,
) {
	ret := C.NSTextView_inst_isIncrementalSearchingEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetIncrementalSearchingEnabled(
	value bool,
) {
	C.NSTextView_inst_setIncrementalSearchingEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) AllowsCharacterPickerTouchBarItem() (
	r0 bool,
) {
	ret := C.NSTextView_inst_allowsCharacterPickerTouchBarItem(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) SetAllowsCharacterPickerTouchBarItem(
	value bool,
) {
	C.NSTextView_inst_setAllowsCharacterPickerTouchBarItem(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSTextView) Font() (
	r0 NSFont,
) {
	ret := C.NSTextView_inst_font(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

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

func (x gen_NSView) InitWithFrame_asNSView(
	frameRect core.NSRect,
) (
	r0 NSView,
) {
	ret := C.NSView_inst_initWithFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSView) PrepareForReuse() {
	C.NSView_inst_prepareForReuse(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) AddSubview(
	view NSViewRef,
) {
	C.NSView_inst_addSubview(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)
	return
}

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

func (x gen_NSView) DidAddSubview(
	subview NSViewRef,
) {
	C.NSView_inst_didAddSubview(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(subview),
	)
	return
}

func (x gen_NSView) RemoveFromSuperview() {
	C.NSView_inst_removeFromSuperview(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) RemoveFromSuperviewWithoutNeedingDisplay() {
	C.NSView_inst_removeFromSuperviewWithoutNeedingDisplay(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

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

func (x gen_NSView) IsDescendantOf(
	view NSViewRef,
) (
	r0 bool,
) {
	ret := C.NSView_inst_isDescendantOf(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) AncestorSharedWithView(
	view NSViewRef,
) (
	r0 NSView,
) {
	ret := C.NSView_inst_ancestorSharedWithView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSView) ViewDidMoveToSuperview() {
	C.NSView_inst_viewDidMoveToSuperview(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) ViewDidMoveToWindow() {
	C.NSView_inst_viewDidMoveToWindow(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) ViewWillMoveToSuperview(
	newSuperview NSViewRef,
) {
	C.NSView_inst_viewWillMoveToSuperview(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newSuperview),
	)
	return
}

func (x gen_NSView) ViewWillMoveToWindow(
	newWindow NSWindowRef,
) {
	C.NSView_inst_viewWillMoveToWindow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newWindow),
	)
	return
}

func (x gen_NSView) WillRemoveSubview(
	subview NSViewRef,
) {
	C.NSView_inst_willRemoveSubview(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(subview),
	)
	return
}

func (x gen_NSView) SetFrameOrigin(
	newOrigin core.NSPoint,
) {
	C.NSView_inst_setFrameOrigin(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&newOrigin)),
	)
	return
}

func (x gen_NSView) SetFrameSize(
	newSize core.NSSize,
) {
	C.NSView_inst_setFrameSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&newSize)),
	)
	return
}

func (x gen_NSView) SetBoundsOrigin(
	newOrigin core.NSPoint,
) {
	C.NSView_inst_setBoundsOrigin(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&newOrigin)),
	)
	return
}

func (x gen_NSView) SetBoundsSize(
	newSize core.NSSize,
) {
	C.NSView_inst_setBoundsSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&newSize)),
	)
	return
}

func (x gen_NSView) MakeBackingLayer() (
	r0 core.CALayer,
) {
	ret := C.NSView_inst_makeBackingLayer(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CALayer_fromPointer(ret)
	return
}

func (x gen_NSView) UpdateLayer() {
	C.NSView_inst_updateLayer(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) DrawRect(
	dirtyRect core.NSRect,
) {
	C.NSView_inst_drawRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&dirtyRect)),
	)
	return
}

func (x gen_NSView) NeedsToDrawRect(
	rect core.NSRect,
) (
	r0 bool,
) {
	ret := C.NSView_inst_needsToDrawRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) Print(
	sender objc.Ref,
) {
	C.NSView_inst_print(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

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

func (x gen_NSView) DataWithEPSInsideRect(
	rect core.NSRect,
) (
	r0 core.NSData,
) {
	ret := C.NSView_inst_dataWithEPSInsideRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = core.NSData_fromPointer(ret)
	return
}

func (x gen_NSView) DataWithPDFInsideRect(
	rect core.NSRect,
) (
	r0 core.NSData,
) {
	ret := C.NSView_inst_dataWithPDFInsideRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = core.NSData_fromPointer(ret)
	return
}

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

func (x gen_NSView) DrawPageBorderWithSize(
	borderSize core.NSSize,
) {
	C.NSView_inst_drawPageBorderWithSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&borderSize)),
	)
	return
}

func (x gen_NSView) RectForPage(
	page core.NSInteger,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_rectForPage(
		unsafe.Pointer(x.Pointer()),
		C.long(page),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) LocationOfPrintRect(
	rect core.NSRect,
) (
	r0 core.NSPoint,
) {
	ret := C.NSView_inst_locationOfPrintRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) SetNeedsDisplayInRect(
	invalidRect core.NSRect,
) {
	C.NSView_inst_setNeedsDisplayInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&invalidRect)),
	)
	return
}

func (x gen_NSView) Display() {
	C.NSView_inst_display(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) DisplayRect(
	rect core.NSRect,
) {
	C.NSView_inst_displayRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	return
}

func (x gen_NSView) DisplayRectIgnoringOpacity(
	rect core.NSRect,
) {
	C.NSView_inst_displayRectIgnoringOpacity(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	return
}

func (x gen_NSView) DisplayIfNeeded() {
	C.NSView_inst_displayIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) DisplayIfNeededInRect(
	rect core.NSRect,
) {
	C.NSView_inst_displayIfNeededInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	return
}

func (x gen_NSView) DisplayIfNeededIgnoringOpacity() {
	C.NSView_inst_displayIfNeededIgnoringOpacity(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) DisplayIfNeededInRectIgnoringOpacity(
	rect core.NSRect,
) {
	C.NSView_inst_displayIfNeededInRectIgnoringOpacity(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	return
}

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

func (x gen_NSView) ViewWillDraw() {
	C.NSView_inst_viewWillDraw(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) ConvertPointFromBacking(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSView_inst_convertPointFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertPointToBacking(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSView_inst_convertPointToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertPointFromLayer(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSView_inst_convertPointFromLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertPointToLayer(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSView_inst_convertPointToLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertRectFromBacking(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_convertRectFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertRectToBacking(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_convertRectToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertRectFromLayer(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_convertRectFromLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertRectToLayer(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_convertRectToLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertSizeFromBacking(
	size core.NSSize,
) (
	r0 core.NSSize,
) {
	ret := C.NSView_inst_convertSizeFromBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertSizeToBacking(
	size core.NSSize,
) (
	r0 core.NSSize,
) {
	ret := C.NSView_inst_convertSizeToBacking(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertSizeFromLayer(
	size core.NSSize,
) (
	r0 core.NSSize,
) {
	ret := C.NSView_inst_convertSizeFromLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertSizeToLayer(
	size core.NSSize,
) (
	r0 core.NSSize,
) {
	ret := C.NSView_inst_convertSizeToLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertPoint_fromView(
	point core.NSPoint,
	view NSViewRef,
) (
	r0 core.NSPoint,
) {
	ret := C.NSView_inst_convertPoint_fromView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
		objc.RefPointer(view),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertPoint_toView(
	point core.NSPoint,
	view NSViewRef,
) (
	r0 core.NSPoint,
) {
	ret := C.NSView_inst_convertPoint_toView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
		objc.RefPointer(view),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertSize_fromView(
	size core.NSSize,
	view NSViewRef,
) (
	r0 core.NSSize,
) {
	ret := C.NSView_inst_convertSize_fromView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
		objc.RefPointer(view),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertSize_toView(
	size core.NSSize,
	view NSViewRef,
) (
	r0 core.NSSize,
) {
	ret := C.NSView_inst_convertSize_toView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
		objc.RefPointer(view),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertRect_fromView(
	rect core.NSRect,
	view NSViewRef,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_convertRect_fromView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(view),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertRect_toView(
	rect core.NSRect,
	view NSViewRef,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_convertRect_toView(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(view),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) CenterScanRect(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_centerScanRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) TranslateOriginToPoint(
	translation core.NSPoint,
) {
	C.NSView_inst_translateOriginToPoint(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&translation)),
	)
	return
}

func (x gen_NSView) ScaleUnitSquareToSize(
	newUnitSize core.NSSize,
) {
	C.NSView_inst_scaleUnitSquareToSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&newUnitSize)),
	)
	return
}

func (x gen_NSView) RotateByAngle(
	angle core.CGFloat,
) {
	C.NSView_inst_rotateByAngle(
		unsafe.Pointer(x.Pointer()),
		C.double(angle),
	)
	return
}

func (x gen_NSView) ResizeSubviewsWithOldSize(
	oldSize core.NSSize,
) {
	C.NSView_inst_resizeSubviewsWithOldSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&oldSize)),
	)
	return
}

func (x gen_NSView) ResizeWithOldSuperviewSize(
	oldSize core.NSSize,
) {
	C.NSView_inst_resizeWithOldSuperviewSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&oldSize)),
	)
	return
}

func (x gen_NSView) AddConstraints(
	constraints core.NSArrayRef,
) {
	C.NSView_inst_addConstraints(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(constraints),
	)
	return
}

func (x gen_NSView) RemoveConstraints(
	constraints core.NSArrayRef,
) {
	C.NSView_inst_removeConstraints(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(constraints),
	)
	return
}

func (x gen_NSView) InvalidateIntrinsicContentSize() {
	C.NSView_inst_invalidateIntrinsicContentSize(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) AlignmentRectForFrame(
	frame core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_alignmentRectForFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frame)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) FrameForAlignmentRect(
	alignmentRect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_frameForAlignmentRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&alignmentRect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) Layout() {
	C.NSView_inst_layout(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) LayoutSubtreeIfNeeded() {
	C.NSView_inst_layoutSubtreeIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) UpdateConstraints() {
	C.NSView_inst_updateConstraints(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) UpdateConstraintsForSubtreeIfNeeded() {
	C.NSView_inst_updateConstraintsForSubtreeIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) ExerciseAmbiguityInLayout() {
	C.NSView_inst_exerciseAmbiguityInLayout(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) DrawFocusRingMask() {
	C.NSView_inst_drawFocusRingMask(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) NoteFocusRingMaskChanged() {
	C.NSView_inst_noteFocusRingMaskChanged(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) SetKeyboardFocusRingNeedsDisplayInRect(
	rect core.NSRect,
) {
	C.NSView_inst_setKeyboardFocusRingNeedsDisplayInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	return
}

func (x gen_NSView) EnterFullScreenMode_withOptions(
	screen NSScreenRef,
	options core.NSDictionaryRef,
) (
	r0 bool,
) {
	ret := C.NSView_inst_enterFullScreenMode_withOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(screen),
		objc.RefPointer(options),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) ExitFullScreenModeWithOptions(
	options core.NSDictionaryRef,
) {
	C.NSView_inst_exitFullScreenModeWithOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(options),
	)
	return
}

func (x gen_NSView) ViewDidHide() {
	C.NSView_inst_viewDidHide(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) ViewDidUnhide() {
	C.NSView_inst_viewDidUnhide(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) ViewWillStartLiveResize() {
	C.NSView_inst_viewWillStartLiveResize(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) ViewDidEndLiveResize() {
	C.NSView_inst_viewDidEndLiveResize(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) AcceptsFirstMouse(
	event NSEventRef,
) (
	r0 bool,
) {
	ret := C.NSView_inst_acceptsFirstMouse(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) HitTest(
	point core.NSPoint,
) (
	r0 NSView,
) {
	ret := C.NSView_inst_hitTest(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSView) Mouse_inRect(
	point core.NSPoint,
	rect core.NSRect,
) (
	r0 bool,
) {
	ret := C.NSView_inst_mouse_inRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) PerformKeyEquivalent(
	event NSEventRef,
) (
	r0 bool,
) {
	ret := C.NSView_inst_performKeyEquivalent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) PrepareContentInRect(
	rect core.NSRect,
) {
	C.NSView_inst_prepareContentInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	return
}

func (x gen_NSView) ScrollPoint(
	point core.NSPoint,
) {
	C.NSView_inst_scrollPoint(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	return
}

func (x gen_NSView) ScrollRectToVisible(
	rect core.NSRect,
) (
	r0 bool,
) {
	ret := C.NSView_inst_scrollRectToVisible(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) Autoscroll(
	event NSEventRef,
) (
	r0 bool,
) {
	ret := C.NSView_inst_autoscroll(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) AdjustScroll(
	newVisible core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_adjustScroll(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&newVisible)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) RegisterForDraggedTypes(
	newTypes core.NSArrayRef,
) {
	C.NSView_inst_registerForDraggedTypes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newTypes),
	)
	return
}

func (x gen_NSView) UnregisterDraggedTypes() {
	C.NSView_inst_unregisterDraggedTypes(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) ShouldDelayWindowOrderingForEvent(
	event NSEventRef,
) (
	r0 bool,
) {
	ret := C.NSView_inst_shouldDelayWindowOrderingForEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) RectForSmartMagnificationAtPoint_inRect(
	location core.NSPoint,
	visibleRect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_rectForSmartMagnificationAtPoint_inRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&location)),
		*(*C.NSRect)(unsafe.Pointer(&visibleRect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ViewDidChangeBackingProperties() {
	C.NSView_inst_viewDidChangeBackingProperties(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) ViewWithTag(
	tag core.NSInteger,
) (
	r0 NSView,
) {
	ret := C.NSView_inst_viewWithTag(
		unsafe.Pointer(x.Pointer()),
		C.long(tag),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSView) RemoveAllToolTips() {
	C.NSView_inst_removeAllToolTips(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) UpdateTrackingAreas() {
	C.NSView_inst_updateTrackingAreas(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) DiscardCursorRects() {
	C.NSView_inst_discardCursorRects(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) ResetCursorRects() {
	C.NSView_inst_resetCursorRects(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) MenuForEvent(
	event NSEventRef,
) (
	r0 NSMenu,
) {
	ret := C.NSView_inst_menuForEvent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	r0 = NSMenu_fromPointer(ret)
	return
}

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

func (x gen_NSView) BeginDocument() {
	C.NSView_inst_beginDocument(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) EndDocument() {
	C.NSView_inst_endDocument(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) EndPage() {
	C.NSView_inst_endPage(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

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

func (x gen_NSView) ViewDidChangeEffectiveAppearance() {
	C.NSView_inst_viewDidChangeEffectiveAppearance(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSView) Init_asNSView() (
	r0 NSView,
) {
	ret := C.NSView_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSView) Superview() (
	r0 NSView,
) {
	ret := C.NSView_inst_superview(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSView) Subviews() (
	r0 core.NSArray,
) {
	ret := C.NSView_inst_subviews(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSView) SetSubviews(
	value core.NSArrayRef,
) {
	C.NSView_inst_setSubviews(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSView) Window() (
	r0 NSWindow,
) {
	ret := C.NSView_inst_window(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSWindow_fromPointer(ret)
	return
}

func (x gen_NSView) OpaqueAncestor() (
	r0 NSView,
) {
	ret := C.NSView_inst_opaqueAncestor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSView) EnclosingMenuItem() (
	r0 NSMenuItem,
) {
	ret := C.NSView_inst_enclosingMenuItem(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func (x gen_NSView) Frame() (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_frame(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) SetFrame(
	value core.NSRect,
) {
	C.NSView_inst_setFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSView) FrameRotation() (
	r0 core.CGFloat,
) {
	ret := C.NSView_inst_frameRotation(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSView) SetFrameRotation(
	value core.CGFloat,
) {
	C.NSView_inst_setFrameRotation(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_NSView) Bounds() (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_bounds(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) SetBounds(
	value core.NSRect,
) {
	C.NSView_inst_setBounds(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSView) BoundsRotation() (
	r0 core.CGFloat,
) {
	ret := C.NSView_inst_boundsRotation(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSView) SetBoundsRotation(
	value core.CGFloat,
) {
	C.NSView_inst_setBoundsRotation(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_NSView) WantsLayer() (
	r0 bool,
) {
	ret := C.NSView_inst_wantsLayer(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) SetWantsLayer(
	value bool,
) {
	C.NSView_inst_setWantsLayer(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSView) WantsUpdateLayer() (
	r0 bool,
) {
	ret := C.NSView_inst_wantsUpdateLayer(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) Layer() (
	r0 core.CALayer,
) {
	ret := C.NSView_inst_layer(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CALayer_fromPointer(ret)
	return
}

func (x gen_NSView) SetLayer(
	value core.CALayerRef,
) {
	C.NSView_inst_setLayer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSView) CanDrawSubviewsIntoLayer() (
	r0 bool,
) {
	ret := C.NSView_inst_canDrawSubviewsIntoLayer(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) SetCanDrawSubviewsIntoLayer(
	value bool,
) {
	C.NSView_inst_setCanDrawSubviewsIntoLayer(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSView) LayerUsesCoreImageFilters() (
	r0 bool,
) {
	ret := C.NSView_inst_layerUsesCoreImageFilters(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) SetLayerUsesCoreImageFilters(
	value bool,
) {
	C.NSView_inst_setLayerUsesCoreImageFilters(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSView) AlphaValue() (
	r0 core.CGFloat,
) {
	ret := C.NSView_inst_alphaValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSView) SetAlphaValue(
	value core.CGFloat,
) {
	C.NSView_inst_setAlphaValue(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_NSView) FrameCenterRotation() (
	r0 core.CGFloat,
) {
	ret := C.NSView_inst_frameCenterRotation(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSView) SetFrameCenterRotation(
	value core.CGFloat,
) {
	C.NSView_inst_setFrameCenterRotation(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_NSView) BackgroundFilters() (
	r0 core.NSArray,
) {
	ret := C.NSView_inst_backgroundFilters(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSView) SetBackgroundFilters(
	value core.NSArrayRef,
) {
	C.NSView_inst_setBackgroundFilters(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSView) ContentFilters() (
	r0 core.NSArray,
) {
	ret := C.NSView_inst_contentFilters(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSView) SetContentFilters(
	value core.NSArrayRef,
) {
	C.NSView_inst_setContentFilters(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSView) CanDrawConcurrently() (
	r0 bool,
) {
	ret := C.NSView_inst_canDrawConcurrently(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) SetCanDrawConcurrently(
	value bool,
) {
	C.NSView_inst_setCanDrawConcurrently(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSView) VisibleRect() (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_visibleRect(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) WantsDefaultClipping() (
	r0 bool,
) {
	ret := C.NSView_inst_wantsDefaultClipping(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) PrintJobTitle() (
	r0 core.NSString,
) {
	ret := C.NSView_inst_printJobTitle(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSView) PageHeader() (
	r0 core.NSAttributedString,
) {
	ret := C.NSView_inst_pageHeader(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSView) PageFooter() (
	r0 core.NSAttributedString,
) {
	ret := C.NSView_inst_pageFooter(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSView) HeightAdjustLimit() (
	r0 core.CGFloat,
) {
	ret := C.NSView_inst_heightAdjustLimit(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSView) WidthAdjustLimit() (
	r0 core.CGFloat,
) {
	ret := C.NSView_inst_widthAdjustLimit(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSView) NeedsDisplay() (
	r0 bool,
) {
	ret := C.NSView_inst_needsDisplay(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) SetNeedsDisplay(
	value bool,
) {
	C.NSView_inst_setNeedsDisplay(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSView) IsOpaque() (
	r0 bool,
) {
	ret := C.NSView_inst_isOpaque(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) IsFlipped() (
	r0 bool,
) {
	ret := C.NSView_inst_isFlipped(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) IsRotatedFromBase() (
	r0 bool,
) {
	ret := C.NSView_inst_isRotatedFromBase(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) IsRotatedOrScaledFromBase() (
	r0 bool,
) {
	ret := C.NSView_inst_isRotatedOrScaledFromBase(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) AutoresizesSubviews() (
	r0 bool,
) {
	ret := C.NSView_inst_autoresizesSubviews(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) SetAutoresizesSubviews(
	value bool,
) {
	C.NSView_inst_setAutoresizesSubviews(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSView) Constraints() (
	r0 core.NSArray,
) {
	ret := C.NSView_inst_constraints(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSView) LayoutGuides() (
	r0 core.NSArray,
) {
	ret := C.NSView_inst_layoutGuides(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSView) FittingSize() (
	r0 core.NSSize,
) {
	ret := C.NSView_inst_fittingSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) IntrinsicContentSize() (
	r0 core.NSSize,
) {
	ret := C.NSView_inst_intrinsicContentSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) BaselineOffsetFromBottom() (
	r0 core.CGFloat,
) {
	ret := C.NSView_inst_baselineOffsetFromBottom(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSView) FirstBaselineOffsetFromTop() (
	r0 core.CGFloat,
) {
	ret := C.NSView_inst_firstBaselineOffsetFromTop(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSView) LastBaselineOffsetFromBottom() (
	r0 core.CGFloat,
) {
	ret := C.NSView_inst_lastBaselineOffsetFromBottom(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSView) NeedsLayout() (
	r0 bool,
) {
	ret := C.NSView_inst_needsLayout(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) SetNeedsLayout(
	value bool,
) {
	C.NSView_inst_setNeedsLayout(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSView) NeedsUpdateConstraints() (
	r0 bool,
) {
	ret := C.NSView_inst_needsUpdateConstraints(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) SetNeedsUpdateConstraints(
	value bool,
) {
	C.NSView_inst_setNeedsUpdateConstraints(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSView) TranslatesAutoresizingMaskIntoConstraints() (
	r0 bool,
) {
	ret := C.NSView_inst_translatesAutoresizingMaskIntoConstraints(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) SetTranslatesAutoresizingMaskIntoConstraints(
	value bool,
) {
	C.NSView_inst_setTranslatesAutoresizingMaskIntoConstraints(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSView) HasAmbiguousLayout() (
	r0 bool,
) {
	ret := C.NSView_inst_hasAmbiguousLayout(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) FocusRingMaskBounds() (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_focusRingMaskBounds(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) AllowsVibrancy() (
	r0 bool,
) {
	ret := C.NSView_inst_allowsVibrancy(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) IsInFullScreenMode() (
	r0 bool,
) {
	ret := C.NSView_inst_isInFullScreenMode(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) IsHidden() (
	r0 bool,
) {
	ret := C.NSView_inst_isHidden(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) SetHidden(
	value bool,
) {
	C.NSView_inst_setHidden(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSView) IsHiddenOrHasHiddenAncestor() (
	r0 bool,
) {
	ret := C.NSView_inst_isHiddenOrHasHiddenAncestor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) InLiveResize() (
	r0 bool,
) {
	ret := C.NSView_inst_inLiveResize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) PreservesContentDuringLiveResize() (
	r0 bool,
) {
	ret := C.NSView_inst_preservesContentDuringLiveResize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) RectPreservedDuringLiveResize() (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_rectPreservedDuringLiveResize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) GestureRecognizers() (
	r0 core.NSArray,
) {
	ret := C.NSView_inst_gestureRecognizers(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSView) SetGestureRecognizers(
	value core.NSArrayRef,
) {
	C.NSView_inst_setGestureRecognizers(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSView) MouseDownCanMoveWindow() (
	r0 bool,
) {
	ret := C.NSView_inst_mouseDownCanMoveWindow(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) WantsRestingTouches() (
	r0 bool,
) {
	ret := C.NSView_inst_wantsRestingTouches(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) SetWantsRestingTouches(
	value bool,
) {
	C.NSView_inst_setWantsRestingTouches(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSView) CanBecomeKeyView() (
	r0 bool,
) {
	ret := C.NSView_inst_canBecomeKeyView(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) NeedsPanelToBecomeKey() (
	r0 bool,
) {
	ret := C.NSView_inst_needsPanelToBecomeKey(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) NextKeyView() (
	r0 NSView,
) {
	ret := C.NSView_inst_nextKeyView(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSView) SetNextKeyView(
	value NSViewRef,
) {
	C.NSView_inst_setNextKeyView(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSView) NextValidKeyView() (
	r0 NSView,
) {
	ret := C.NSView_inst_nextValidKeyView(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSView) PreviousKeyView() (
	r0 NSView,
) {
	ret := C.NSView_inst_previousKeyView(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSView) PreviousValidKeyView() (
	r0 NSView,
) {
	ret := C.NSView_inst_previousValidKeyView(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSView) PreparedContentRect() (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_preparedContentRect(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) SetPreparedContentRect(
	value core.NSRect,
) {
	C.NSView_inst_setPreparedContentRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSView) RegisteredDraggedTypes() (
	r0 core.NSArray,
) {
	ret := C.NSView_inst_registeredDraggedTypes(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSView) PostsFrameChangedNotifications() (
	r0 bool,
) {
	ret := C.NSView_inst_postsFrameChangedNotifications(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) SetPostsFrameChangedNotifications(
	value bool,
) {
	C.NSView_inst_setPostsFrameChangedNotifications(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSView) PostsBoundsChangedNotifications() (
	r0 bool,
) {
	ret := C.NSView_inst_postsBoundsChangedNotifications(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) SetPostsBoundsChangedNotifications(
	value bool,
) {
	C.NSView_inst_setPostsBoundsChangedNotifications(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSView) Tag() (
	r0 core.NSInteger,
) {
	ret := C.NSView_inst_tag(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSView) ToolTip() (
	r0 core.NSString,
) {
	ret := C.NSView_inst_toolTip(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSView) SetToolTip(
	value core.NSStringRef,
) {
	C.NSView_inst_setToolTip(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSView) TrackingAreas() (
	r0 core.NSArray,
) {
	ret := C.NSView_inst_trackingAreas(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSView) IsDrawingFindIndicator() (
	r0 bool,
) {
	ret := C.NSView_inst_isDrawingFindIndicator(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) IsHorizontalContentSizeConstraintActive() (
	r0 bool,
) {
	ret := C.NSView_inst_isHorizontalContentSizeConstraintActive(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) SetHorizontalContentSizeConstraintActive(
	value bool,
) {
	C.NSView_inst_setHorizontalContentSizeConstraintActive(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSView) IsVerticalContentSizeConstraintActive() (
	r0 bool,
) {
	ret := C.NSView_inst_isVerticalContentSizeConstraintActive(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) SetVerticalContentSizeConstraintActive(
	value bool,
) {
	C.NSView_inst_setVerticalContentSizeConstraintActive(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_NSView) BackgroundColor() (
	r0 NSColor,
) {
	ret := C.NSView_inst_backgroundColor(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSView) SetBackgroundColor(
	value NSColorRef,
) {
	C.NSView_inst_setBackgroundColor(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}
