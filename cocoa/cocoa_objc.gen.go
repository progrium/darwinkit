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
void* NSBundle_type_bundleWithURL_(void* url) {
	return [NSBundle
		bundleWithURL: url];
}
void* NSBundle_type_bundleWithPath_(void* path) {
	return [NSBundle
		bundleWithPath: path];
}
void* NSBundle_type_bundleWithIdentifier_(void* identifier) {
	return [NSBundle
		bundleWithIdentifier: identifier];
}
void* NSBundle_type_URLForResource_withExtension_subdirectory_inBundleWithURL_(void* name, void* ext, void* subpath, void* bundleURL) {
	return [NSBundle
		URLForResource: name
		withExtension: ext
		subdirectory: subpath
		inBundleWithURL: bundleURL];
}
void* NSBundle_type_URLsForResourcesWithExtension_subdirectory_inBundleWithURL_(void* ext, void* subpath, void* bundleURL) {
	return [NSBundle
		URLsForResourcesWithExtension: ext
		subdirectory: subpath
		inBundleWithURL: bundleURL];
}
void* NSBundle_type_pathForResource_ofType_inDirectory_(void* name, void* ext, void* bundlePath) {
	return [NSBundle
		pathForResource: name
		ofType: ext
		inDirectory: bundlePath];
}
void* NSBundle_type_pathsForResourcesOfType_inDirectory_(void* ext, void* bundlePath) {
	return [NSBundle
		pathsForResourcesOfType: ext
		inDirectory: bundlePath];
}
void* NSBundle_type_preferredLocalizationsFromArray_(void* localizationsArray) {
	return [NSBundle
		preferredLocalizationsFromArray: localizationsArray];
}
void* NSBundle_type_preferredLocalizationsFromArray_forPreferences_(void* localizationsArray, void* preferencesArray) {
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
BOOL NSSound_type_canInitWithPasteboard_(void* pasteboard) {
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
void NSApplication_type_detachDrawingThread_toTarget_withObject_(void* selector, void* target, void* argument) {
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
void* NSButton_type_checkboxWithTitle_target_action_(void* title, void* target, void* action) {
	return [NSButton
		checkboxWithTitle: title
		target: target
		action: action];
}
void* NSButton_type_buttonWithImage_target_action_(void* image, void* target, void* action) {
	return [NSButton
		buttonWithImage: image
		target: target
		action: action];
}
void* NSButton_type_radioButtonWithTitle_target_action_(void* title, void* target, void* action) {
	return [NSButton
		radioButtonWithTitle: title
		target: target
		action: action];
}
void* NSButton_type_buttonWithTitle_image_target_action_(void* title, void* image, void* target, void* action) {
	return [NSButton
		buttonWithTitle: title
		image: image
		target: target
		action: action];
}
void* NSButton_type_buttonWithTitle_target_action_(void* title, void* target, void* action) {
	return [NSButton
		buttonWithTitle: title
		target: target
		action: action];
}
void* NSEvent_type_alloc() {
	return [NSEvent
		alloc];
}
void* NSEvent_type_eventWithEventRef_(void* eventRef) {
	return [NSEvent
		eventWithEventRef: eventRef];
}
void NSEvent_type_stopPeriodicEvents() {
	[NSEvent
		stopPeriodicEvents];
}
void NSEvent_type_removeMonitor_(void* eventMonitor) {
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
void NSEvent_type_setMouseCoalescingEnabled_(BOOL value) {
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
void* NSFont_type_fontWithName_size_(void* fontName, double fontSize) {
	return [NSFont
		fontWithName: fontName
		size: fontSize];
}
void* NSFont_type_userFontOfSize_(double fontSize) {
	return [NSFont
		userFontOfSize: fontSize];
}
void* NSFont_type_userFixedPitchFontOfSize_(double fontSize) {
	return [NSFont
		userFixedPitchFontOfSize: fontSize];
}
void* NSFont_type_systemFontOfSize_(double fontSize) {
	return [NSFont
		systemFontOfSize: fontSize];
}
void* NSFont_type_boldSystemFontOfSize_(double fontSize) {
	return [NSFont
		boldSystemFontOfSize: fontSize];
}
void* NSFont_type_labelFontOfSize_(double fontSize) {
	return [NSFont
		labelFontOfSize: fontSize];
}
void* NSFont_type_messageFontOfSize_(double fontSize) {
	return [NSFont
		messageFontOfSize: fontSize];
}
void* NSFont_type_menuBarFontOfSize_(double fontSize) {
	return [NSFont
		menuBarFontOfSize: fontSize];
}
void* NSFont_type_menuFontOfSize_(double fontSize) {
	return [NSFont
		menuFontOfSize: fontSize];
}
void* NSFont_type_controlContentFontOfSize_(double fontSize) {
	return [NSFont
		controlContentFontOfSize: fontSize];
}
void* NSFont_type_titleBarFontOfSize_(double fontSize) {
	return [NSFont
		titleBarFontOfSize: fontSize];
}
void* NSFont_type_paletteFontOfSize_(double fontSize) {
	return [NSFont
		paletteFontOfSize: fontSize];
}
void* NSFont_type_toolTipsFontOfSize_(double fontSize) {
	return [NSFont
		toolTipsFontOfSize: fontSize];
}
void NSFont_type_setUserFont_(void* font) {
	[NSFont
		setUserFont: font];
}
void NSFont_type_setUserFixedPitchFont_(void* font) {
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
void* NSImage_type_imageWithSystemSymbolName_accessibilityDescription_(void* symbolName, void* description) {
	return [NSImage
		imageWithSystemSymbolName: symbolName
		accessibilityDescription: description];
}
BOOL NSImage_type_canInitWithPasteboard_(void* pasteboard) {
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
void* NSImageView_type_imageViewWithImage_(void* image) {
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
void* NSPasteboard_type_pasteboardByFilteringFile_(void* filename) {
	return [NSPasteboard
		pasteboardByFilteringFile: filename];
}
void* NSPasteboard_type_pasteboardByFilteringTypesInPasteboard_(void* pboard) {
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
void NSMenu_type_setMenuBarVisible_(BOOL visible) {
	[NSMenu
		setMenuBarVisible: visible];
}
void NSMenu_type_popUpContextMenu_withEvent_forView_(void* menu, void* event, void* view) {
	[NSMenu
		popUpContextMenu: menu
		withEvent: event
		forView: view];
}
void NSMenu_type_popUpContextMenu_withEvent_forView_withFont_(void* menu, void* event, void* view, void* font) {
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
void NSMenuItem_type_setUsesUserKeyEquivalents_(BOOL value) {
	[NSMenuItem
		setUsesUserKeyEquivalents: value];
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
void* NSWindow_type_windowWithContentViewController_(void* contentViewController) {
	return [NSWindow
		windowWithContentViewController: contentViewController];
}
NSRect NSWindow_type_contentRectForFrameRect_styleMask_(NSRect fRect, unsigned long style) {
	return [NSWindow
		contentRectForFrameRect: fRect
		styleMask: style];
}
NSRect NSWindow_type_frameRectForContentRect_styleMask_(NSRect cRect, unsigned long style) {
	return [NSWindow
		frameRectForContentRect: cRect
		styleMask: style];
}
double NSWindow_type_minFrameWidthWithTitle_styleMask_(void* title, unsigned long style) {
	return [NSWindow
		minFrameWidthWithTitle: title
		styleMask: style];
}
long NSWindow_type_windowNumberAtPoint_belowWindowWithWindowNumber_(NSPoint point, long windowNumber) {
	return [NSWindow
		windowNumberAtPoint: point
		belowWindowWithWindowNumber: windowNumber];
}
BOOL NSWindow_type_allowsAutomaticWindowTabbing() {
	return [NSWindow
		allowsAutomaticWindowTabbing];
}
void NSWindow_type_setAllowsAutomaticWindowTabbing_(BOOL value) {
	[NSWindow
		setAllowsAutomaticWindowTabbing: value];
}
void* NSColor_type_alloc() {
	return [NSColor
		alloc];
}
void* NSColor_type_colorFromPasteboard_(void* pasteBoard) {
	return [NSColor
		colorFromPasteboard: pasteBoard];
}
void* NSColor_type_colorWithRed_green_blue_alpha_(double red, double green, double blue, double alpha) {
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
void NSColor_type_setIgnoresAlpha_(BOOL value) {
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


void* NSBundle_inst_initWithURL_(void *id, void* url) {
	return [(NSBundle*)id
		initWithURL: url];
}

void* NSBundle_inst_initWithPath_(void *id, void* path) {
	return [(NSBundle*)id
		initWithPath: path];
}

void* NSBundle_inst_loadNibNamed_owner_options_(void *id, void* name, void* owner, void* options) {
	return [(NSBundle*)id
		loadNibNamed: name
		owner: owner
		options: options];
}

void* NSBundle_inst_URLForResource_withExtension_subdirectory_(void *id, void* name, void* ext, void* subpath) {
	return [(NSBundle*)id
		URLForResource: name
		withExtension: ext
		subdirectory: subpath];
}

void* NSBundle_inst_URLForResource_withExtension_(void *id, void* name, void* ext) {
	return [(NSBundle*)id
		URLForResource: name
		withExtension: ext];
}

void* NSBundle_inst_URLsForResourcesWithExtension_subdirectory_(void *id, void* ext, void* subpath) {
	return [(NSBundle*)id
		URLsForResourcesWithExtension: ext
		subdirectory: subpath];
}

void* NSBundle_inst_URLForResource_withExtension_subdirectory_localization_(void *id, void* name, void* ext, void* subpath, void* localizationName) {
	return [(NSBundle*)id
		URLForResource: name
		withExtension: ext
		subdirectory: subpath
		localization: localizationName];
}

void* NSBundle_inst_URLsForResourcesWithExtension_subdirectory_localization_(void *id, void* ext, void* subpath, void* localizationName) {
	return [(NSBundle*)id
		URLsForResourcesWithExtension: ext
		subdirectory: subpath
		localization: localizationName];
}

void* NSBundle_inst_pathForResource_ofType_(void *id, void* name, void* ext) {
	return [(NSBundle*)id
		pathForResource: name
		ofType: ext];
}

void* NSBundle_inst_pathForResource_ofType_inDirectory_(void *id, void* name, void* ext, void* subpath) {
	return [(NSBundle*)id
		pathForResource: name
		ofType: ext
		inDirectory: subpath];
}

void* NSBundle_inst_pathForResource_ofType_inDirectory_forLocalization_(void *id, void* name, void* ext, void* subpath, void* localizationName) {
	return [(NSBundle*)id
		pathForResource: name
		ofType: ext
		inDirectory: subpath
		forLocalization: localizationName];
}

void* NSBundle_inst_pathsForResourcesOfType_inDirectory_(void *id, void* ext, void* subpath) {
	return [(NSBundle*)id
		pathsForResourcesOfType: ext
		inDirectory: subpath];
}

void* NSBundle_inst_pathsForResourcesOfType_inDirectory_forLocalization_(void *id, void* ext, void* subpath, void* localizationName) {
	return [(NSBundle*)id
		pathsForResourcesOfType: ext
		inDirectory: subpath
		forLocalization: localizationName];
}

void* NSBundle_inst_localizedStringForKey_value_table_(void *id, void* key, void* value, void* tableName) {
	return [(NSBundle*)id
		localizedStringForKey: key
		value: value
		table: tableName];
}

void* NSBundle_inst_URLForAuxiliaryExecutable_(void *id, void* executableName) {
	return [(NSBundle*)id
		URLForAuxiliaryExecutable: executableName];
}

void* NSBundle_inst_pathForAuxiliaryExecutable_(void *id, void* executableName) {
	return [(NSBundle*)id
		pathForAuxiliaryExecutable: executableName];
}

void* NSBundle_inst_objectForInfoDictionaryKey_(void *id, void* key) {
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

void* NSBundle_inst_localizedAttributedStringForKey_value_table_(void *id, void* key, void* value, void* tableName) {
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

void* NSSound_inst_initWithContentsOfFile_byReference_(void *id, void* path, BOOL byRef) {
	return [(NSSound*)id
		initWithContentsOfFile: path
		byReference: byRef];
}

void* NSSound_inst_initWithContentsOfURL_byReference_(void *id, void* url, BOOL byRef) {
	return [(NSSound*)id
		initWithContentsOfURL: url
		byReference: byRef];
}

void* NSSound_inst_initWithData_(void *id, void* data) {
	return [(NSSound*)id
		initWithData: data];
}

void* NSSound_inst_initWithPasteboard_(void *id, void* pasteboard) {
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

void NSSound_inst_writeToPasteboard_(void *id, void* pasteboard) {
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

void NSSound_inst_setDelegate_(void *id, void* value) {
	[(NSSound*)id
		setDelegate: value];
}

BOOL NSSound_inst_loops(void *id) {
	return [(NSSound*)id
		loops];
}

void NSSound_inst_setLoops_(void *id, BOOL value) {
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

void NSApplication_inst_stop_(void *id, void* sender) {
	[(NSApplication*)id
		stop: sender];
}

void NSApplication_inst_sendEvent_(void *id, void* event) {
	[(NSApplication*)id
		sendEvent: event];
}

void NSApplication_inst_postEvent_atStart_(void *id, void* event, BOOL flag) {
	[(NSApplication*)id
		postEvent: event
		atStart: flag];
}

BOOL NSApplication_inst_tryToPerform_with_(void *id, void* action, void* object) {
	return [(NSApplication*)id
		tryToPerform: action
		with: object];
}

BOOL NSApplication_inst_sendAction_to_from_(void *id, void* action, void* target, void* sender) {
	return [(NSApplication*)id
		sendAction: action
		to: target
		from: sender];
}

void* NSApplication_inst_targetForAction_(void *id, void* action) {
	return [(NSApplication*)id
		targetForAction: action];
}

void* NSApplication_inst_targetForAction_to_from_(void *id, void* action, void* target, void* sender) {
	return [(NSApplication*)id
		targetForAction: action
		to: target
		from: sender];
}

void NSApplication_inst_terminate_(void *id, void* sender) {
	[(NSApplication*)id
		terminate: sender];
}

void NSApplication_inst_replyToApplicationShouldTerminate_(void *id, BOOL shouldTerminate) {
	[(NSApplication*)id
		replyToApplicationShouldTerminate: shouldTerminate];
}

void NSApplication_inst_activateIgnoringOtherApps_(void *id, BOOL flag) {
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

void NSApplication_inst_toggleTouchBarCustomizationPalette_(void *id, void* sender) {
	[(NSApplication*)id
		toggleTouchBarCustomizationPalette: sender];
}

void NSApplication_inst_cancelUserAttentionRequest_(void *id, long request) {
	[(NSApplication*)id
		cancelUserAttentionRequest: request];
}

void NSApplication_inst_registerUserInterfaceItemSearchHandler_(void *id, void* handler) {
	[(NSApplication*)id
		registerUserInterfaceItemSearchHandler: handler];
}

void NSApplication_inst_unregisterUserInterfaceItemSearchHandler_(void *id, void* handler) {
	[(NSApplication*)id
		unregisterUserInterfaceItemSearchHandler: handler];
}

void NSApplication_inst_showHelp_(void *id, void* sender) {
	[(NSApplication*)id
		showHelp: sender];
}

void NSApplication_inst_activateContextHelpMode_(void *id, void* sender) {
	[(NSApplication*)id
		activateContextHelpMode: sender];
}

void NSApplication_inst_hideOtherApplications_(void *id, void* sender) {
	[(NSApplication*)id
		hideOtherApplications: sender];
}

void NSApplication_inst_unhideAllApplications_(void *id, void* sender) {
	[(NSApplication*)id
		unhideAllApplications: sender];
}

long NSApplication_inst_activationPolicy(void *id) {
	return [(NSApplication*)id
		activationPolicy];
}

BOOL NSApplication_inst_setActivationPolicy_(void *id, long activationPolicy) {
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

void NSApplication_inst_setDelegate_(void *id, void* value) {
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

void NSApplication_inst_setApplicationIconImage_(void *id, void* value) {
	[(NSApplication*)id
		setApplicationIconImage: value];
}

void* NSApplication_inst_helpMenu(void *id) {
	return [(NSApplication*)id
		helpMenu];
}

void NSApplication_inst_setHelpMenu_(void *id, void* value) {
	[(NSApplication*)id
		setHelpMenu: value];
}

void* NSApplication_inst_servicesProvider(void *id) {
	return [(NSApplication*)id
		servicesProvider];
}

void NSApplication_inst_setServicesProvider_(void *id, void* value) {
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

void NSApplication_inst_setMainMenu_(void *id, void* value) {
	[(NSApplication*)id
		setMainMenu: value];
}

void* NSControl_inst_initWithFrame_(void *id, NSRect frameRect) {
	return [(NSControl*)id
		initWithFrame: frameRect];
}

void NSControl_inst_takeDoubleValueFrom_(void *id, void* sender) {
	[(NSControl*)id
		takeDoubleValueFrom: sender];
}

void NSControl_inst_takeFloatValueFrom_(void *id, void* sender) {
	[(NSControl*)id
		takeFloatValueFrom: sender];
}

void NSControl_inst_takeIntValueFrom_(void *id, void* sender) {
	[(NSControl*)id
		takeIntValueFrom: sender];
}

void NSControl_inst_takeIntegerValueFrom_(void *id, void* sender) {
	[(NSControl*)id
		takeIntegerValueFrom: sender];
}

void NSControl_inst_takeObjectValueFrom_(void *id, void* sender) {
	[(NSControl*)id
		takeObjectValueFrom: sender];
}

void NSControl_inst_takeStringValueFrom_(void *id, void* sender) {
	[(NSControl*)id
		takeStringValueFrom: sender];
}

void NSControl_inst_drawWithExpansionFrame_inView_(void *id, NSRect contentFrame, void* view) {
	[(NSControl*)id
		drawWithExpansionFrame: contentFrame
		inView: view];
}

NSRect NSControl_inst_expansionFrameWithFrame_(void *id, NSRect contentFrame) {
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

void NSControl_inst_editWithFrame_editor_delegate_event_(void *id, NSRect rect, void* textObj, void* delegate, void* event) {
	[(NSControl*)id
		editWithFrame: rect
		editor: textObj
		delegate: delegate
		event: event];
}

void NSControl_inst_endEditing_(void *id, void* textObj) {
	[(NSControl*)id
		endEditing: textObj];
}

void NSControl_inst_selectWithFrame_editor_delegate_start_length_(void *id, NSRect rect, void* textObj, void* delegate, long selStart, long selLength) {
	[(NSControl*)id
		selectWithFrame: rect
		editor: textObj
		delegate: delegate
		start: selStart
		length: selLength];
}

NSSize NSControl_inst_sizeThatFits_(void *id, NSSize size) {
	return [(NSControl*)id
		sizeThatFits: size];
}

void NSControl_inst_sizeToFit(void *id) {
	[(NSControl*)id
		sizeToFit];
}

BOOL NSControl_inst_sendAction_to_(void *id, void* action, void* target) {
	return [(NSControl*)id
		sendAction: action
		to: target];
}

void NSControl_inst_performClick_(void *id, void* sender) {
	[(NSControl*)id
		performClick: sender];
}

void NSControl_inst_mouseDown_(void *id, void* event) {
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

void NSControl_inst_setEnabled_(void *id, BOOL value) {
	[(NSControl*)id
		setEnabled: value];
}

int NSControl_inst_intValue(void *id) {
	return [(NSControl*)id
		intValue];
}

void NSControl_inst_setIntValue_(void *id, int value) {
	[(NSControl*)id
		setIntValue: value];
}

long NSControl_inst_integerValue(void *id) {
	return [(NSControl*)id
		integerValue];
}

void NSControl_inst_setIntegerValue_(void *id, long value) {
	[(NSControl*)id
		setIntegerValue: value];
}

void* NSControl_inst_objectValue(void *id) {
	return [(NSControl*)id
		objectValue];
}

void NSControl_inst_setObjectValue_(void *id, void* value) {
	[(NSControl*)id
		setObjectValue: value];
}

void* NSControl_inst_stringValue(void *id) {
	return [(NSControl*)id
		stringValue];
}

void NSControl_inst_setStringValue_(void *id, void* value) {
	[(NSControl*)id
		setStringValue: value];
}

void* NSControl_inst_attributedStringValue(void *id) {
	return [(NSControl*)id
		attributedStringValue];
}

void NSControl_inst_setAttributedStringValue_(void *id, void* value) {
	[(NSControl*)id
		setAttributedStringValue: value];
}

void* NSControl_inst_font(void *id) {
	return [(NSControl*)id
		font];
}

void NSControl_inst_setFont_(void *id, void* value) {
	[(NSControl*)id
		setFont: value];
}

BOOL NSControl_inst_usesSingleLineMode(void *id) {
	return [(NSControl*)id
		usesSingleLineMode];
}

void NSControl_inst_setUsesSingleLineMode_(void *id, BOOL value) {
	[(NSControl*)id
		setUsesSingleLineMode: value];
}

BOOL NSControl_inst_allowsExpansionToolTips(void *id) {
	return [(NSControl*)id
		allowsExpansionToolTips];
}

void NSControl_inst_setAllowsExpansionToolTips_(void *id, BOOL value) {
	[(NSControl*)id
		setAllowsExpansionToolTips: value];
}

BOOL NSControl_inst_isHighlighted(void *id) {
	return [(NSControl*)id
		isHighlighted];
}

void NSControl_inst_setHighlighted_(void *id, BOOL value) {
	[(NSControl*)id
		setHighlighted: value];
}

void* NSControl_inst_action(void *id) {
	return [(NSControl*)id
		action];
}

void NSControl_inst_setAction_(void *id, void* value) {
	[(NSControl*)id
		setAction: value];
}

void* NSControl_inst_target(void *id) {
	return [(NSControl*)id
		target];
}

void NSControl_inst_setTarget_(void *id, void* value) {
	[(NSControl*)id
		setTarget: value];
}

BOOL NSControl_inst_isContinuous(void *id) {
	return [(NSControl*)id
		isContinuous];
}

void NSControl_inst_setContinuous_(void *id, BOOL value) {
	[(NSControl*)id
		setContinuous: value];
}

long NSControl_inst_tag(void *id) {
	return [(NSControl*)id
		tag];
}

void NSControl_inst_setTag_(void *id, long value) {
	[(NSControl*)id
		setTag: value];
}

BOOL NSControl_inst_refusesFirstResponder(void *id) {
	return [(NSControl*)id
		refusesFirstResponder];
}

void NSControl_inst_setRefusesFirstResponder_(void *id, BOOL value) {
	[(NSControl*)id
		setRefusesFirstResponder: value];
}

BOOL NSControl_inst_ignoresMultiClick(void *id) {
	return [(NSControl*)id
		ignoresMultiClick];
}

void NSControl_inst_setIgnoresMultiClick_(void *id, BOOL value) {
	[(NSControl*)id
		setIgnoresMultiClick: value];
}

void NSButton_inst_compressWithPrioritizedCompressionOptions_(void *id, void* prioritizedOptions) {
	[(NSButton*)id
		compressWithPrioritizedCompressionOptions: prioritizedOptions];
}

NSSize NSButton_inst_minimumSizeWithPrioritizedCompressionOptions_(void *id, void* prioritizedOptions) {
	return [(NSButton*)id
		minimumSizeWithPrioritizedCompressionOptions: prioritizedOptions];
}

void NSButton_inst_setNextState(void *id) {
	[(NSButton*)id
		setNextState];
}

void NSButton_inst_highlight_(void *id, BOOL flag) {
	[(NSButton*)id
		highlight: flag];
}

BOOL NSButton_inst_performKeyEquivalent_(void *id, void* key) {
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

void NSButton_inst_setContentTintColor_(void *id, void* value) {
	[(NSButton*)id
		setContentTintColor: value];
}

BOOL NSButton_inst_hasDestructiveAction(void *id) {
	return [(NSButton*)id
		hasDestructiveAction];
}

void NSButton_inst_setHasDestructiveAction_(void *id, BOOL value) {
	[(NSButton*)id
		setHasDestructiveAction: value];
}

void* NSButton_inst_alternateTitle(void *id) {
	return [(NSButton*)id
		alternateTitle];
}

void NSButton_inst_setAlternateTitle_(void *id, void* value) {
	[(NSButton*)id
		setAlternateTitle: value];
}

void* NSButton_inst_attributedTitle(void *id) {
	return [(NSButton*)id
		attributedTitle];
}

void NSButton_inst_setAttributedTitle_(void *id, void* value) {
	[(NSButton*)id
		setAttributedTitle: value];
}

void* NSButton_inst_attributedAlternateTitle(void *id) {
	return [(NSButton*)id
		attributedAlternateTitle];
}

void NSButton_inst_setAttributedAlternateTitle_(void *id, void* value) {
	[(NSButton*)id
		setAttributedAlternateTitle: value];
}

void* NSButton_inst_title(void *id) {
	return [(NSButton*)id
		title];
}

void NSButton_inst_setTitle_(void *id, void* value) {
	[(NSButton*)id
		setTitle: value];
}

void* NSButton_inst_sound(void *id) {
	return [(NSButton*)id
		sound];
}

void NSButton_inst_setSound_(void *id, void* value) {
	[(NSButton*)id
		setSound: value];
}

BOOL NSButton_inst_isSpringLoaded(void *id) {
	return [(NSButton*)id
		isSpringLoaded];
}

void NSButton_inst_setSpringLoaded_(void *id, BOOL value) {
	[(NSButton*)id
		setSpringLoaded: value];
}

long NSButton_inst_maxAcceleratorLevel(void *id) {
	return [(NSButton*)id
		maxAcceleratorLevel];
}

void NSButton_inst_setMaxAcceleratorLevel_(void *id, long value) {
	[(NSButton*)id
		setMaxAcceleratorLevel: value];
}

void* NSButton_inst_image(void *id) {
	return [(NSButton*)id
		image];
}

void NSButton_inst_setImage_(void *id, void* value) {
	[(NSButton*)id
		setImage: value];
}

void* NSButton_inst_alternateImage(void *id) {
	return [(NSButton*)id
		alternateImage];
}

void NSButton_inst_setAlternateImage_(void *id, void* value) {
	[(NSButton*)id
		setAlternateImage: value];
}

BOOL NSButton_inst_isBordered(void *id) {
	return [(NSButton*)id
		isBordered];
}

void NSButton_inst_setBordered_(void *id, BOOL value) {
	[(NSButton*)id
		setBordered: value];
}

BOOL NSButton_inst_isTransparent(void *id) {
	return [(NSButton*)id
		isTransparent];
}

void NSButton_inst_setTransparent_(void *id, BOOL value) {
	[(NSButton*)id
		setTransparent: value];
}

void* NSButton_inst_bezelColor(void *id) {
	return [(NSButton*)id
		bezelColor];
}

void NSButton_inst_setBezelColor_(void *id, void* value) {
	[(NSButton*)id
		setBezelColor: value];
}

BOOL NSButton_inst_showsBorderOnlyWhileMouseInside(void *id) {
	return [(NSButton*)id
		showsBorderOnlyWhileMouseInside];
}

void NSButton_inst_setShowsBorderOnlyWhileMouseInside_(void *id, BOOL value) {
	[(NSButton*)id
		setShowsBorderOnlyWhileMouseInside: value];
}

BOOL NSButton_inst_imageHugsTitle(void *id) {
	return [(NSButton*)id
		imageHugsTitle];
}

void NSButton_inst_setImageHugsTitle_(void *id, BOOL value) {
	[(NSButton*)id
		setImageHugsTitle: value];
}

BOOL NSButton_inst_allowsMixedState(void *id) {
	return [(NSButton*)id
		allowsMixedState];
}

void NSButton_inst_setAllowsMixedState_(void *id, BOOL value) {
	[(NSButton*)id
		setAllowsMixedState: value];
}

long NSButton_inst_state(void *id) {
	return [(NSButton*)id
		state];
}

void NSButton_inst_setState_(void *id, long value) {
	[(NSButton*)id
		setState: value];
}

void* NSButton_inst_keyEquivalent(void *id) {
	return [(NSButton*)id
		keyEquivalent];
}

void NSButton_inst_setKeyEquivalent_(void *id, void* value) {
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

void* NSFont_inst_fontWithSize_(void *id, double fontSize) {
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

void* NSImage_inst_initByReferencingFile_(void *id, void* fileName) {
	return [(NSImage*)id
		initByReferencingFile: fileName];
}

void* NSImage_inst_initByReferencingURL_(void *id, void* url) {
	return [(NSImage*)id
		initByReferencingURL: url];
}

void* NSImage_inst_initWithContentsOfFile_(void *id, void* fileName) {
	return [(NSImage*)id
		initWithContentsOfFile: fileName];
}

void* NSImage_inst_initWithContentsOfURL_(void *id, void* url) {
	return [(NSImage*)id
		initWithContentsOfURL: url];
}

void* NSImage_inst_initWithData_(void *id, void* data) {
	return [(NSImage*)id
		initWithData: data];
}

void* NSImage_inst_initWithDataIgnoringOrientation_(void *id, void* data) {
	return [(NSImage*)id
		initWithDataIgnoringOrientation: data];
}

void* NSImage_inst_initWithPasteboard_(void *id, void* pasteboard) {
	return [(NSImage*)id
		initWithPasteboard: pasteboard];
}

void* NSImage_inst_initWithSize_(void *id, NSSize size) {
	return [(NSImage*)id
		initWithSize: size];
}

BOOL NSImage_inst_isTemplate(void *id) {
	return [(NSImage*)id
		isTemplate];
}

void NSImage_inst_addRepresentations_(void *id, void* imageReps) {
	[(NSImage*)id
		addRepresentations: imageReps];
}

void NSImage_inst_drawInRect_(void *id, NSRect rect) {
	[(NSImage*)id
		drawInRect: rect];
}

void NSImage_inst_lockFocus(void *id) {
	[(NSImage*)id
		lockFocus];
}

void NSImage_inst_lockFocusFlipped_(void *id, BOOL flipped) {
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

void* NSImage_inst_layerContentsForContentsScale_(void *id, double layerContentsScale) {
	return [(NSImage*)id
		layerContentsForContentsScale: layerContentsScale];
}

double NSImage_inst_recommendedLayerContentsScale_(void *id, double preferredContentsScale) {
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

void NSImage_inst_setDelegate_(void *id, void* value) {
	[(NSImage*)id
		setDelegate: value];
}

NSSize NSImage_inst_size(void *id) {
	return [(NSImage*)id
		size];
}

void NSImage_inst_setSize_(void *id, NSSize value) {
	[(NSImage*)id
		setSize: value];
}

void NSImage_inst_setTemplate_(void *id, BOOL value) {
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

void NSImage_inst_setPrefersColorMatch_(void *id, BOOL value) {
	[(NSImage*)id
		setPrefersColorMatch: value];
}

BOOL NSImage_inst_usesEPSOnResolutionMismatch(void *id) {
	return [(NSImage*)id
		usesEPSOnResolutionMismatch];
}

void NSImage_inst_setUsesEPSOnResolutionMismatch_(void *id, BOOL value) {
	[(NSImage*)id
		setUsesEPSOnResolutionMismatch: value];
}

BOOL NSImage_inst_matchesOnMultipleResolution(void *id) {
	return [(NSImage*)id
		matchesOnMultipleResolution];
}

void NSImage_inst_setMatchesOnMultipleResolution_(void *id, BOOL value) {
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

void NSImage_inst_setBackgroundColor_(void *id, void* value) {
	[(NSImage*)id
		setBackgroundColor: value];
}

NSRect NSImage_inst_alignmentRect(void *id) {
	return [(NSImage*)id
		alignmentRect];
}

void NSImage_inst_setAlignmentRect_(void *id, NSRect value) {
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

void NSImage_inst_setAccessibilityDescription_(void *id, void* value) {
	[(NSImage*)id
		setAccessibilityDescription: value];
}

BOOL NSImage_inst_matchesOnlyOnBestFittingAxis(void *id) {
	return [(NSImage*)id
		matchesOnlyOnBestFittingAxis];
}

void NSImage_inst_setMatchesOnlyOnBestFittingAxis_(void *id, BOOL value) {
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

void NSImageView_inst_setImage_(void *id, void* value) {
	[(NSImageView*)id
		setImage: value];
}

BOOL NSImageView_inst_animates(void *id) {
	return [(NSImageView*)id
		animates];
}

void NSImageView_inst_setAnimates_(void *id, BOOL value) {
	[(NSImageView*)id
		setAnimates: value];
}

BOOL NSImageView_inst_isEditable(void *id) {
	return [(NSImageView*)id
		isEditable];
}

void NSImageView_inst_setEditable_(void *id, BOOL value) {
	[(NSImageView*)id
		setEditable: value];
}

BOOL NSImageView_inst_allowsCutCopyPaste(void *id) {
	return [(NSImageView*)id
		allowsCutCopyPaste];
}

void NSImageView_inst_setAllowsCutCopyPaste_(void *id, BOOL value) {
	[(NSImageView*)id
		setAllowsCutCopyPaste: value];
}

void* NSImageView_inst_contentTintColor(void *id) {
	return [(NSImageView*)id
		contentTintColor];
}

void NSImageView_inst_setContentTintColor_(void *id, void* value) {
	[(NSImageView*)id
		setContentTintColor: value];
}

void* NSNib_inst_initWithNibData_bundle_(void *id, void* nibData, void* bundle) {
	return [(NSNib*)id
		initWithNibData: nibData
		bundle: bundle];
}

BOOL NSNib_inst_instantiateWithOwner_topLevelObjects_(void *id, void* owner, void* topLevelObjects) {
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

BOOL NSPasteboard_inst_writeObjects_(void *id, void* objects) {
	return [(NSPasteboard*)id
		writeObjects: objects];
}

void* NSPasteboard_inst_readObjectsForClasses_options_(void *id, void* classArray, void* options) {
	return [(NSPasteboard*)id
		readObjectsForClasses: classArray
		options: options];
}

BOOL NSPasteboard_inst_canReadItemWithDataConformingToTypes_(void *id, void* types) {
	return [(NSPasteboard*)id
		canReadItemWithDataConformingToTypes: types];
}

BOOL NSPasteboard_inst_canReadObjectForClasses_options_(void *id, void* classArray, void* options) {
	return [(NSPasteboard*)id
		canReadObjectForClasses: classArray
		options: options];
}

long NSPasteboard_inst_declareTypes_owner_(void *id, void* newTypes, void* newOwner) {
	return [(NSPasteboard*)id
		declareTypes: newTypes
		owner: newOwner];
}

long NSPasteboard_inst_addTypes_owner_(void *id, void* newTypes, void* newOwner) {
	return [(NSPasteboard*)id
		addTypes: newTypes
		owner: newOwner];
}

BOOL NSPasteboard_inst_writeFileContents_(void *id, void* filename) {
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

void NSLayoutManager_inst_addTextContainer_(void *id, void* container) {
	[(NSLayoutManager*)id
		addTextContainer: container];
}

void NSLayoutManager_inst_insertTextContainer_atIndex_(void *id, void* container, unsigned long index) {
	[(NSLayoutManager*)id
		insertTextContainer: container
		atIndex: index];
}

void NSLayoutManager_inst_removeTextContainerAtIndex_(void *id, unsigned long index) {
	[(NSLayoutManager*)id
		removeTextContainerAtIndex: index];
}

void NSLayoutManager_inst_textContainerChangedGeometry_(void *id, void* container) {
	[(NSLayoutManager*)id
		textContainerChangedGeometry: container];
}

void NSLayoutManager_inst_textContainerChangedTextView_(void *id, void* container) {
	[(NSLayoutManager*)id
		textContainerChangedTextView: container];
}

NSRect NSLayoutManager_inst_usedRectForTextContainer_(void *id, void* container) {
	return [(NSLayoutManager*)id
		usedRectForTextContainer: container];
}

void NSLayoutManager_inst_ensureLayoutForBoundingRect_inTextContainer_(void *id, NSRect bounds, void* container) {
	[(NSLayoutManager*)id
		ensureLayoutForBoundingRect: bounds
		inTextContainer: container];
}

void NSLayoutManager_inst_ensureLayoutForTextContainer_(void *id, void* container) {
	[(NSLayoutManager*)id
		ensureLayoutForTextContainer: container];
}

unsigned long NSLayoutManager_inst_characterIndexForGlyphAtIndex_(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		characterIndexForGlyphAtIndex: glyphIndex];
}

unsigned long NSLayoutManager_inst_glyphIndexForCharacterAtIndex_(void *id, unsigned long charIndex) {
	return [(NSLayoutManager*)id
		glyphIndexForCharacterAtIndex: charIndex];
}

BOOL NSLayoutManager_inst_isValidGlyphIndex_(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		isValidGlyphIndex: glyphIndex];
}

void NSLayoutManager_inst_setDrawsOutsideLineFragment_forGlyphAtIndex_(void *id, BOOL flag, unsigned long glyphIndex) {
	[(NSLayoutManager*)id
		setDrawsOutsideLineFragment: flag
		forGlyphAtIndex: glyphIndex];
}

void NSLayoutManager_inst_setExtraLineFragmentRect_usedRect_textContainer_(void *id, NSRect fragmentRect, NSRect usedRect, void* container) {
	[(NSLayoutManager*)id
		setExtraLineFragmentRect: fragmentRect
		usedRect: usedRect
		textContainer: container];
}

void NSLayoutManager_inst_setNotShownAttribute_forGlyphAtIndex_(void *id, BOOL flag, unsigned long glyphIndex) {
	[(NSLayoutManager*)id
		setNotShownAttribute: flag
		forGlyphAtIndex: glyphIndex];
}

NSSize NSLayoutManager_inst_attachmentSizeForGlyphAtIndex_(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		attachmentSizeForGlyphAtIndex: glyphIndex];
}

BOOL NSLayoutManager_inst_drawsOutsideLineFragmentForGlyphAtIndex_(void *id, unsigned long glyphIndex) {
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

BOOL NSLayoutManager_inst_notShownAttributeForGlyphAtIndex_(void *id, unsigned long glyphIndex) {
	return [(NSLayoutManager*)id
		notShownAttributeForGlyphAtIndex: glyphIndex];
}

BOOL NSLayoutManager_inst_layoutManagerOwnsFirstResponderInWindow_(void *id, void* window) {
	return [(NSLayoutManager*)id
		layoutManagerOwnsFirstResponderInWindow: window];
}

double NSLayoutManager_inst_defaultLineHeightForFont_(void *id, void* theFont) {
	return [(NSLayoutManager*)id
		defaultLineHeightForFont: theFont];
}

double NSLayoutManager_inst_defaultBaselineOffsetForFont_(void *id, void* theFont) {
	return [(NSLayoutManager*)id
		defaultBaselineOffsetForFont: theFont];
}

void* NSLayoutManager_inst_delegate(void *id) {
	return [(NSLayoutManager*)id
		delegate];
}

void NSLayoutManager_inst_setDelegate_(void *id, void* value) {
	[(NSLayoutManager*)id
		setDelegate: value];
}

BOOL NSLayoutManager_inst_allowsNonContiguousLayout(void *id) {
	return [(NSLayoutManager*)id
		allowsNonContiguousLayout];
}

void NSLayoutManager_inst_setAllowsNonContiguousLayout_(void *id, BOOL value) {
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

void NSLayoutManager_inst_setShowsInvisibleCharacters_(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setShowsInvisibleCharacters: value];
}

BOOL NSLayoutManager_inst_showsControlCharacters(void *id) {
	return [(NSLayoutManager*)id
		showsControlCharacters];
}

void NSLayoutManager_inst_setShowsControlCharacters_(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setShowsControlCharacters: value];
}

BOOL NSLayoutManager_inst_usesFontLeading(void *id) {
	return [(NSLayoutManager*)id
		usesFontLeading];
}

void NSLayoutManager_inst_setUsesFontLeading_(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setUsesFontLeading: value];
}

BOOL NSLayoutManager_inst_backgroundLayoutEnabled(void *id) {
	return [(NSLayoutManager*)id
		backgroundLayoutEnabled];
}

void NSLayoutManager_inst_setBackgroundLayoutEnabled_(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setBackgroundLayoutEnabled: value];
}

BOOL NSLayoutManager_inst_limitsLayoutForSuspiciousContents(void *id) {
	return [(NSLayoutManager*)id
		limitsLayoutForSuspiciousContents];
}

void NSLayoutManager_inst_setLimitsLayoutForSuspiciousContents_(void *id, BOOL value) {
	[(NSLayoutManager*)id
		setLimitsLayoutForSuspiciousContents: value];
}

BOOL NSLayoutManager_inst_usesDefaultHyphenation(void *id) {
	return [(NSLayoutManager*)id
		usesDefaultHyphenation];
}

void NSLayoutManager_inst_setUsesDefaultHyphenation_(void *id, BOOL value) {
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

void* NSMenu_inst_initWithTitle_(void *id, void* title) {
	return [(NSMenu*)id
		initWithTitle: title];
}

void NSMenu_inst_insertItem_atIndex_(void *id, void* newItem, long index) {
	[(NSMenu*)id
		insertItem: newItem
		atIndex: index];
}

void* NSMenu_inst_insertItemWithTitle_action_keyEquivalent_atIndex_(void *id, void* string, void* selector, void* charCode, long index) {
	return [(NSMenu*)id
		insertItemWithTitle: string
		action: selector
		keyEquivalent: charCode
		atIndex: index];
}

void NSMenu_inst_addItem_(void *id, void* newItem) {
	[(NSMenu*)id
		addItem: newItem];
}

void* NSMenu_inst_addItemWithTitle_action_keyEquivalent_(void *id, void* string, void* selector, void* charCode) {
	return [(NSMenu*)id
		addItemWithTitle: string
		action: selector
		keyEquivalent: charCode];
}

void NSMenu_inst_removeItem_(void *id, void* item) {
	[(NSMenu*)id
		removeItem: item];
}

void NSMenu_inst_removeItemAtIndex_(void *id, long index) {
	[(NSMenu*)id
		removeItemAtIndex: index];
}

void NSMenu_inst_itemChanged_(void *id, void* item) {
	[(NSMenu*)id
		itemChanged: item];
}

void NSMenu_inst_removeAllItems(void *id) {
	[(NSMenu*)id
		removeAllItems];
}

void* NSMenu_inst_itemWithTag_(void *id, long tag) {
	return [(NSMenu*)id
		itemWithTag: tag];
}

void* NSMenu_inst_itemWithTitle_(void *id, void* title) {
	return [(NSMenu*)id
		itemWithTitle: title];
}

void* NSMenu_inst_itemAtIndex_(void *id, long index) {
	return [(NSMenu*)id
		itemAtIndex: index];
}

long NSMenu_inst_indexOfItem_(void *id, void* item) {
	return [(NSMenu*)id
		indexOfItem: item];
}

long NSMenu_inst_indexOfItemWithTitle_(void *id, void* title) {
	return [(NSMenu*)id
		indexOfItemWithTitle: title];
}

long NSMenu_inst_indexOfItemWithTag_(void *id, long tag) {
	return [(NSMenu*)id
		indexOfItemWithTag: tag];
}

long NSMenu_inst_indexOfItemWithTarget_andAction_(void *id, void* target, void* actionSelector) {
	return [(NSMenu*)id
		indexOfItemWithTarget: target
		andAction: actionSelector];
}

long NSMenu_inst_indexOfItemWithRepresentedObject_(void *id, void* object) {
	return [(NSMenu*)id
		indexOfItemWithRepresentedObject: object];
}

long NSMenu_inst_indexOfItemWithSubmenu_(void *id, void* submenu) {
	return [(NSMenu*)id
		indexOfItemWithSubmenu: submenu];
}

void NSMenu_inst_setSubmenu_forItem_(void *id, void* menu, void* item) {
	[(NSMenu*)id
		setSubmenu: menu
		forItem: item];
}

void NSMenu_inst_submenuAction_(void *id, void* sender) {
	[(NSMenu*)id
		submenuAction: sender];
}

void NSMenu_inst_update(void *id) {
	[(NSMenu*)id
		update];
}

BOOL NSMenu_inst_performKeyEquivalent_(void *id, void* event) {
	return [(NSMenu*)id
		performKeyEquivalent: event];
}

void NSMenu_inst_performActionForItemAtIndex_(void *id, long index) {
	[(NSMenu*)id
		performActionForItemAtIndex: index];
}

BOOL NSMenu_inst_popUpMenuPositioningItem_atLocation_inView_(void *id, void* item, NSPoint location, void* view) {
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

void NSMenu_inst_setItemArray_(void *id, void* value) {
	[(NSMenu*)id
		setItemArray: value];
}

void* NSMenu_inst_supermenu(void *id) {
	return [(NSMenu*)id
		supermenu];
}

void NSMenu_inst_setSupermenu_(void *id, void* value) {
	[(NSMenu*)id
		setSupermenu: value];
}

BOOL NSMenu_inst_autoenablesItems(void *id) {
	return [(NSMenu*)id
		autoenablesItems];
}

void NSMenu_inst_setAutoenablesItems_(void *id, BOOL value) {
	[(NSMenu*)id
		setAutoenablesItems: value];
}

void* NSMenu_inst_font(void *id) {
	return [(NSMenu*)id
		font];
}

void NSMenu_inst_setFont_(void *id, void* value) {
	[(NSMenu*)id
		setFont: value];
}

void* NSMenu_inst_title(void *id) {
	return [(NSMenu*)id
		title];
}

void NSMenu_inst_setTitle_(void *id, void* value) {
	[(NSMenu*)id
		setTitle: value];
}

double NSMenu_inst_minimumWidth(void *id) {
	return [(NSMenu*)id
		minimumWidth];
}

void NSMenu_inst_setMinimumWidth_(void *id, double value) {
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

void NSMenu_inst_setAllowsContextMenuPlugIns_(void *id, BOOL value) {
	[(NSMenu*)id
		setAllowsContextMenuPlugIns: value];
}

BOOL NSMenu_inst_showsStateColumn(void *id) {
	return [(NSMenu*)id
		showsStateColumn];
}

void NSMenu_inst_setShowsStateColumn_(void *id, BOOL value) {
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

void NSMenu_inst_setDelegate_(void *id, void* value) {
	[(NSMenu*)id
		setDelegate: value];
}

void NSPopover_inst_performClose_(void *id, void* sender) {
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

void NSPopover_inst_setBehavior_(void *id, long value) {
	[(NSPopover*)id
		setBehavior: value];
}

NSRect NSPopover_inst_positioningRect(void *id) {
	return [(NSPopover*)id
		positioningRect];
}

void NSPopover_inst_setPositioningRect_(void *id, NSRect value) {
	[(NSPopover*)id
		setPositioningRect: value];
}

BOOL NSPopover_inst_animates(void *id) {
	return [(NSPopover*)id
		animates];
}

void NSPopover_inst_setAnimates_(void *id, BOOL value) {
	[(NSPopover*)id
		setAnimates: value];
}

NSSize NSPopover_inst_contentSize(void *id) {
	return [(NSPopover*)id
		contentSize];
}

void NSPopover_inst_setContentSize_(void *id, NSSize value) {
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

void* NSMenuItem_inst_initWithTitle_action_keyEquivalent_(void *id, void* string, void* selector, void* charCode) {
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

void NSMenuItem_inst_setEnabled_(void *id, BOOL value) {
	[(NSMenuItem*)id
		setEnabled: value];
}

BOOL NSMenuItem_inst_isHidden(void *id) {
	return [(NSMenuItem*)id
		isHidden];
}

void NSMenuItem_inst_setHidden_(void *id, BOOL value) {
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

void NSMenuItem_inst_setTarget_(void *id, void* value) {
	[(NSMenuItem*)id
		setTarget: value];
}

void* NSMenuItem_inst_action(void *id) {
	return [(NSMenuItem*)id
		action];
}

void NSMenuItem_inst_setAction_(void *id, void* value) {
	[(NSMenuItem*)id
		setAction: value];
}

void* NSMenuItem_inst_title(void *id) {
	return [(NSMenuItem*)id
		title];
}

void NSMenuItem_inst_setTitle_(void *id, void* value) {
	[(NSMenuItem*)id
		setTitle: value];
}

void* NSMenuItem_inst_attributedTitle(void *id) {
	return [(NSMenuItem*)id
		attributedTitle];
}

void NSMenuItem_inst_setAttributedTitle_(void *id, void* value) {
	[(NSMenuItem*)id
		setAttributedTitle: value];
}

long NSMenuItem_inst_tag(void *id) {
	return [(NSMenuItem*)id
		tag];
}

void NSMenuItem_inst_setTag_(void *id, long value) {
	[(NSMenuItem*)id
		setTag: value];
}

long NSMenuItem_inst_state(void *id) {
	return [(NSMenuItem*)id
		state];
}

void NSMenuItem_inst_setState_(void *id, long value) {
	[(NSMenuItem*)id
		setState: value];
}

void* NSMenuItem_inst_image(void *id) {
	return [(NSMenuItem*)id
		image];
}

void NSMenuItem_inst_setImage_(void *id, void* value) {
	[(NSMenuItem*)id
		setImage: value];
}

void* NSMenuItem_inst_onStateImage(void *id) {
	return [(NSMenuItem*)id
		onStateImage];
}

void NSMenuItem_inst_setOnStateImage_(void *id, void* value) {
	[(NSMenuItem*)id
		setOnStateImage: value];
}

void* NSMenuItem_inst_offStateImage(void *id) {
	return [(NSMenuItem*)id
		offStateImage];
}

void NSMenuItem_inst_setOffStateImage_(void *id, void* value) {
	[(NSMenuItem*)id
		setOffStateImage: value];
}

void* NSMenuItem_inst_mixedStateImage(void *id) {
	return [(NSMenuItem*)id
		mixedStateImage];
}

void NSMenuItem_inst_setMixedStateImage_(void *id, void* value) {
	[(NSMenuItem*)id
		setMixedStateImage: value];
}

void* NSMenuItem_inst_submenu(void *id) {
	return [(NSMenuItem*)id
		submenu];
}

void NSMenuItem_inst_setSubmenu_(void *id, void* value) {
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

void NSMenuItem_inst_setMenu_(void *id, void* value) {
	[(NSMenuItem*)id
		setMenu: value];
}

void* NSMenuItem_inst_keyEquivalent(void *id) {
	return [(NSMenuItem*)id
		keyEquivalent];
}

void NSMenuItem_inst_setKeyEquivalent_(void *id, void* value) {
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

void NSMenuItem_inst_setAlternate_(void *id, BOOL value) {
	[(NSMenuItem*)id
		setAlternate: value];
}

long NSMenuItem_inst_indentationLevel(void *id) {
	return [(NSMenuItem*)id
		indentationLevel];
}

void NSMenuItem_inst_setIndentationLevel_(void *id, long value) {
	[(NSMenuItem*)id
		setIndentationLevel: value];
}

void* NSMenuItem_inst_toolTip(void *id) {
	return [(NSMenuItem*)id
		toolTip];
}

void NSMenuItem_inst_setToolTip_(void *id, void* value) {
	[(NSMenuItem*)id
		setToolTip: value];
}

void* NSMenuItem_inst_representedObject(void *id) {
	return [(NSMenuItem*)id
		representedObject];
}

void NSMenuItem_inst_setRepresentedObject_(void *id, void* value) {
	[(NSMenuItem*)id
		setRepresentedObject: value];
}

void* NSMenuItem_inst_view(void *id) {
	return [(NSMenuItem*)id
		view];
}

void NSMenuItem_inst_setView_(void *id, void* value) {
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

void NSMenuItem_inst_setAllowsAutomaticKeyEquivalentLocalization_(void *id, BOOL value) {
	[(NSMenuItem*)id
		setAllowsAutomaticKeyEquivalentLocalization: value];
}

BOOL NSMenuItem_inst_allowsAutomaticKeyEquivalentMirroring(void *id) {
	return [(NSMenuItem*)id
		allowsAutomaticKeyEquivalentMirroring];
}

void NSMenuItem_inst_setAllowsAutomaticKeyEquivalentMirroring_(void *id, BOOL value) {
	[(NSMenuItem*)id
		setAllowsAutomaticKeyEquivalentMirroring: value];
}

BOOL NSMenuItem_inst_allowsKeyEquivalentWhenHidden(void *id) {
	return [(NSMenuItem*)id
		allowsKeyEquivalentWhenHidden];
}

void NSMenuItem_inst_setAllowsKeyEquivalentWhenHidden_(void *id, BOOL value) {
	[(NSMenuItem*)id
		setAllowsKeyEquivalentWhenHidden: value];
}

NSRect NSScreen_inst_convertRectFromBacking_(void *id, NSRect rect) {
	return [(NSScreen*)id
		convertRectFromBacking: rect];
}

NSRect NSScreen_inst_convertRectToBacking_(void *id, NSRect rect) {
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

void* NSStatusBar_inst_statusItemWithLength_(void *id, double length) {
	return [(NSStatusBar*)id
		statusItemWithLength: length];
}

void NSStatusBar_inst_removeStatusItem_(void *id, void* item) {
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

void NSStatusBarButton_inst_setAppearsDisabled_(void *id, BOOL value) {
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

void NSStatusItem_inst_setMenu_(void *id, void* value) {
	[(NSStatusItem*)id
		setMenu: value];
}

BOOL NSStatusItem_inst_isVisible(void *id) {
	return [(NSStatusItem*)id
		isVisible];
}

void NSStatusItem_inst_setVisible_(void *id, BOOL value) {
	[(NSStatusItem*)id
		setVisible: value];
}

double NSStatusItem_inst_length(void *id) {
	return [(NSStatusItem*)id
		length];
}

void NSStatusItem_inst_setLength_(void *id, double value) {
	[(NSStatusItem*)id
		setLength: value];
}

void* NSText_inst_initWithFrame_(void *id, NSRect frameRect) {
	return [(NSText*)id
		initWithFrame: frameRect];
}

void NSText_inst_toggleRuler_(void *id, void* sender) {
	[(NSText*)id
		toggleRuler: sender];
}

void NSText_inst_selectAll_(void *id, void* sender) {
	[(NSText*)id
		selectAll: sender];
}

void NSText_inst_copy_(void *id, void* sender) {
	[(NSText*)id
		copy: sender];
}

void NSText_inst_cut_(void *id, void* sender) {
	[(NSText*)id
		cut: sender];
}

void NSText_inst_paste_(void *id, void* sender) {
	[(NSText*)id
		paste: sender];
}

void NSText_inst_copyFont_(void *id, void* sender) {
	[(NSText*)id
		copyFont: sender];
}

void NSText_inst_pasteFont_(void *id, void* sender) {
	[(NSText*)id
		pasteFont: sender];
}

void NSText_inst_copyRuler_(void *id, void* sender) {
	[(NSText*)id
		copyRuler: sender];
}

void NSText_inst_pasteRuler_(void *id, void* sender) {
	[(NSText*)id
		pasteRuler: sender];
}

void NSText_inst_delete_(void *id, void* sender) {
	[(NSText*)id
		delete: sender];
}

void NSText_inst_changeFont_(void *id, void* sender) {
	[(NSText*)id
		changeFont: sender];
}

void NSText_inst_alignCenter_(void *id, void* sender) {
	[(NSText*)id
		alignCenter: sender];
}

void NSText_inst_alignLeft_(void *id, void* sender) {
	[(NSText*)id
		alignLeft: sender];
}

void NSText_inst_alignRight_(void *id, void* sender) {
	[(NSText*)id
		alignRight: sender];
}

void NSText_inst_superscript_(void *id, void* sender) {
	[(NSText*)id
		superscript: sender];
}

void NSText_inst_subscript_(void *id, void* sender) {
	[(NSText*)id
		subscript: sender];
}

void NSText_inst_unscript_(void *id, void* sender) {
	[(NSText*)id
		unscript: sender];
}

void NSText_inst_underline_(void *id, void* sender) {
	[(NSText*)id
		underline: sender];
}

BOOL NSText_inst_readRTFDFromFile_(void *id, void* path) {
	return [(NSText*)id
		readRTFDFromFile: path];
}

BOOL NSText_inst_writeRTFDToFile_atomically_(void *id, void* path, BOOL flag) {
	return [(NSText*)id
		writeRTFDToFile: path
		atomically: flag];
}

void NSText_inst_checkSpelling_(void *id, void* sender) {
	[(NSText*)id
		checkSpelling: sender];
}

void NSText_inst_showGuessPanel_(void *id, void* sender) {
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

void NSText_inst_setString_(void *id, void* value) {
	[(NSText*)id
		setString: value];
}

void* NSText_inst_backgroundColor(void *id) {
	return [(NSText*)id
		backgroundColor];
}

void NSText_inst_setBackgroundColor_(void *id, void* value) {
	[(NSText*)id
		setBackgroundColor: value];
}

BOOL NSText_inst_drawsBackground(void *id) {
	return [(NSText*)id
		drawsBackground];
}

void NSText_inst_setDrawsBackground_(void *id, BOOL value) {
	[(NSText*)id
		setDrawsBackground: value];
}

BOOL NSText_inst_isEditable(void *id) {
	return [(NSText*)id
		isEditable];
}

void NSText_inst_setEditable_(void *id, BOOL value) {
	[(NSText*)id
		setEditable: value];
}

BOOL NSText_inst_isSelectable(void *id) {
	return [(NSText*)id
		isSelectable];
}

void NSText_inst_setSelectable_(void *id, BOOL value) {
	[(NSText*)id
		setSelectable: value];
}

BOOL NSText_inst_isFieldEditor(void *id) {
	return [(NSText*)id
		isFieldEditor];
}

void NSText_inst_setFieldEditor_(void *id, BOOL value) {
	[(NSText*)id
		setFieldEditor: value];
}

BOOL NSText_inst_isRichText(void *id) {
	return [(NSText*)id
		isRichText];
}

void NSText_inst_setRichText_(void *id, BOOL value) {
	[(NSText*)id
		setRichText: value];
}

BOOL NSText_inst_importsGraphics(void *id) {
	return [(NSText*)id
		importsGraphics];
}

void NSText_inst_setImportsGraphics_(void *id, BOOL value) {
	[(NSText*)id
		setImportsGraphics: value];
}

BOOL NSText_inst_usesFontPanel(void *id) {
	return [(NSText*)id
		usesFontPanel];
}

void NSText_inst_setUsesFontPanel_(void *id, BOOL value) {
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

void NSText_inst_setFont_(void *id, void* value) {
	[(NSText*)id
		setFont: value];
}

void* NSText_inst_textColor(void *id) {
	return [(NSText*)id
		textColor];
}

void NSText_inst_setTextColor_(void *id, void* value) {
	[(NSText*)id
		setTextColor: value];
}

NSSize NSText_inst_maxSize(void *id) {
	return [(NSText*)id
		maxSize];
}

void NSText_inst_setMaxSize_(void *id, NSSize value) {
	[(NSText*)id
		setMaxSize: value];
}

NSSize NSText_inst_minSize(void *id) {
	return [(NSText*)id
		minSize];
}

void NSText_inst_setMinSize_(void *id, NSSize value) {
	[(NSText*)id
		setMinSize: value];
}

BOOL NSText_inst_isVerticallyResizable(void *id) {
	return [(NSText*)id
		isVerticallyResizable];
}

void NSText_inst_setVerticallyResizable_(void *id, BOOL value) {
	[(NSText*)id
		setVerticallyResizable: value];
}

BOOL NSText_inst_isHorizontallyResizable(void *id) {
	return [(NSText*)id
		isHorizontallyResizable];
}

void NSText_inst_setHorizontallyResizable_(void *id, BOOL value) {
	[(NSText*)id
		setHorizontallyResizable: value];
}

void* NSText_inst_delegate(void *id) {
	return [(NSText*)id
		delegate];
}

void NSText_inst_setDelegate_(void *id, void* value) {
	[(NSText*)id
		setDelegate: value];
}

void* NSTextContainer_inst_initWithSize_(void *id, NSSize size) {
	return [(NSTextContainer*)id
		initWithSize: size];
}

void NSTextContainer_inst_replaceLayoutManager_(void *id, void* newLayoutManager) {
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

void NSTextContainer_inst_setLayoutManager_(void *id, void* value) {
	[(NSTextContainer*)id
		setLayoutManager: value];
}

void* NSTextContainer_inst_textView(void *id) {
	return [(NSTextContainer*)id
		textView];
}

void NSTextContainer_inst_setTextView_(void *id, void* value) {
	[(NSTextContainer*)id
		setTextView: value];
}

NSSize NSTextContainer_inst_size(void *id) {
	return [(NSTextContainer*)id
		size];
}

void NSTextContainer_inst_setSize_(void *id, NSSize value) {
	[(NSTextContainer*)id
		setSize: value];
}

void* NSTextContainer_inst_exclusionPaths(void *id) {
	return [(NSTextContainer*)id
		exclusionPaths];
}

void NSTextContainer_inst_setExclusionPaths_(void *id, void* value) {
	[(NSTextContainer*)id
		setExclusionPaths: value];
}

BOOL NSTextContainer_inst_widthTracksTextView(void *id) {
	return [(NSTextContainer*)id
		widthTracksTextView];
}

void NSTextContainer_inst_setWidthTracksTextView_(void *id, BOOL value) {
	[(NSTextContainer*)id
		setWidthTracksTextView: value];
}

BOOL NSTextContainer_inst_heightTracksTextView(void *id) {
	return [(NSTextContainer*)id
		heightTracksTextView];
}

void NSTextContainer_inst_setHeightTracksTextView_(void *id, BOOL value) {
	[(NSTextContainer*)id
		setHeightTracksTextView: value];
}

unsigned long NSTextContainer_inst_maximumNumberOfLines(void *id) {
	return [(NSTextContainer*)id
		maximumNumberOfLines];
}

void NSTextContainer_inst_setMaximumNumberOfLines_(void *id, unsigned long value) {
	[(NSTextContainer*)id
		setMaximumNumberOfLines: value];
}

double NSTextContainer_inst_lineFragmentPadding(void *id) {
	return [(NSTextContainer*)id
		lineFragmentPadding];
}

void NSTextContainer_inst_setLineFragmentPadding_(void *id, double value) {
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

void NSViewController_inst_commitEditingWithDelegate_didCommitSelector_contextInfo_(void *id, void* delegate, void* didCommitSelector, void* contextInfo) {
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

void NSViewController_inst_dismissController_(void *id, void* sender) {
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

void NSViewController_inst_addChildViewController_(void *id, void* childViewController) {
	[(NSViewController*)id
		addChildViewController: childViewController];
}

void NSViewController_inst_insertChildViewController_atIndex_(void *id, void* childViewController, long index) {
	[(NSViewController*)id
		insertChildViewController: childViewController
		atIndex: index];
}

void NSViewController_inst_removeChildViewControllerAtIndex_(void *id, long index) {
	[(NSViewController*)id
		removeChildViewControllerAtIndex: index];
}

void NSViewController_inst_removeFromParentViewController(void *id) {
	[(NSViewController*)id
		removeFromParentViewController];
}

void NSViewController_inst_preferredContentSizeDidChangeForViewController_(void *id, void* viewController) {
	[(NSViewController*)id
		preferredContentSizeDidChangeForViewController: viewController];
}

void NSViewController_inst_presentViewController_animator_(void *id, void* viewController, void* animator) {
	[(NSViewController*)id
		presentViewController: viewController
		animator: animator];
}

void NSViewController_inst_dismissViewController_(void *id, void* viewController) {
	[(NSViewController*)id
		dismissViewController: viewController];
}

void NSViewController_inst_presentViewControllerAsModalWindow_(void *id, void* viewController) {
	[(NSViewController*)id
		presentViewControllerAsModalWindow: viewController];
}

void NSViewController_inst_presentViewControllerAsSheet_(void *id, void* viewController) {
	[(NSViewController*)id
		presentViewControllerAsSheet: viewController];
}

void NSViewController_inst_viewWillTransitionToSize_(void *id, NSSize newSize) {
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

void NSViewController_inst_setRepresentedObject_(void *id, void* value) {
	[(NSViewController*)id
		setRepresentedObject: value];
}

void* NSViewController_inst_nibBundle(void *id) {
	return [(NSViewController*)id
		nibBundle];
}

void* NSViewController_inst_title(void *id) {
	return [(NSViewController*)id
		title];
}

void NSViewController_inst_setTitle_(void *id, void* value) {
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

void NSViewController_inst_setPreferredContentSize_(void *id, NSSize value) {
	[(NSViewController*)id
		setPreferredContentSize: value];
}

void* NSViewController_inst_childViewControllers(void *id) {
	return [(NSViewController*)id
		childViewControllers];
}

void NSViewController_inst_setChildViewControllers_(void *id, void* value) {
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

void NSViewController_inst_setPreferredScreenOrigin_(void *id, NSPoint value) {
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

void NSVisualEffectView_inst_viewDidMoveToWindow(void *id) {
	[(NSVisualEffectView*)id
		viewDidMoveToWindow];
}

void NSVisualEffectView_inst_viewWillMoveToWindow_(void *id, void* newWindow) {
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

void NSVisualEffectView_inst_setEmphasized_(void *id, BOOL value) {
	[(NSVisualEffectView*)id
		setEmphasized: value];
}

void* NSVisualEffectView_inst_maskImage(void *id) {
	return [(NSVisualEffectView*)id
		maskImage];
}

void NSVisualEffectView_inst_setMaskImage_(void *id, void* value) {
	[(NSVisualEffectView*)id
		setMaskImage: value];
}

void* NSWindow_inst_initWithContentRect_styleMask_backing_defer_(void *id, NSRect contentRect, unsigned long style, unsigned long backingStoreType, BOOL flag) {
	return [(NSWindow*)id
		initWithContentRect: contentRect
		styleMask: style
		backing: backingStoreType
		defer: flag];
}

void* NSWindow_inst_initWithContentRect_styleMask_backing_defer_screen_(void *id, NSRect contentRect, unsigned long style, unsigned long backingStoreType, BOOL flag, void* screen) {
	return [(NSWindow*)id
		initWithContentRect: contentRect
		styleMask: style
		backing: backingStoreType
		defer: flag
		screen: screen];
}

void NSWindow_inst_toggleFullScreen_(void *id, void* sender) {
	[(NSWindow*)id
		toggleFullScreen: sender];
}

void NSWindow_inst_setDynamicDepthLimit_(void *id, BOOL flag) {
	[(NSWindow*)id
		setDynamicDepthLimit: flag];
}

void NSWindow_inst_invalidateShadow(void *id) {
	[(NSWindow*)id
		invalidateShadow];
}

NSRect NSWindow_inst_contentRectForFrameRect_(void *id, NSRect frameRect) {
	return [(NSWindow*)id
		contentRectForFrameRect: frameRect];
}

NSRect NSWindow_inst_frameRectForContentRect_(void *id, NSRect contentRect) {
	return [(NSWindow*)id
		frameRectForContentRect: contentRect];
}

void NSWindow_inst_endSheet_(void *id, void* sheetWindow) {
	[(NSWindow*)id
		endSheet: sheetWindow];
}

void NSWindow_inst_setFrameOrigin_(void *id, NSPoint point) {
	[(NSWindow*)id
		setFrameOrigin: point];
}

void NSWindow_inst_setFrameTopLeftPoint_(void *id, NSPoint point) {
	[(NSWindow*)id
		setFrameTopLeftPoint: point];
}

NSRect NSWindow_inst_constrainFrameRect_toScreen_(void *id, NSRect frameRect, void* screen) {
	return [(NSWindow*)id
		constrainFrameRect: frameRect
		toScreen: screen];
}

NSPoint NSWindow_inst_cascadeTopLeftFromPoint_(void *id, NSPoint topLeftPoint) {
	return [(NSWindow*)id
		cascadeTopLeftFromPoint: topLeftPoint];
}

void NSWindow_inst_setFrame_display_(void *id, NSRect frameRect, BOOL flag) {
	[(NSWindow*)id
		setFrame: frameRect
		display: flag];
}

void NSWindow_inst_setFrame_display_animate_(void *id, NSRect frameRect, BOOL displayFlag, BOOL animateFlag) {
	[(NSWindow*)id
		setFrame: frameRect
		display: displayFlag
		animate: animateFlag];
}

void NSWindow_inst_performZoom_(void *id, void* sender) {
	[(NSWindow*)id
		performZoom: sender];
}

void NSWindow_inst_zoom_(void *id, void* sender) {
	[(NSWindow*)id
		zoom: sender];
}

void NSWindow_inst_setContentSize_(void *id, NSSize size) {
	[(NSWindow*)id
		setContentSize: size];
}

void NSWindow_inst_orderOut_(void *id, void* sender) {
	[(NSWindow*)id
		orderOut: sender];
}

void NSWindow_inst_orderBack_(void *id, void* sender) {
	[(NSWindow*)id
		orderBack: sender];
}

void NSWindow_inst_orderFront_(void *id, void* sender) {
	[(NSWindow*)id
		orderFront: sender];
}

void NSWindow_inst_orderFrontRegardless(void *id) {
	[(NSWindow*)id
		orderFrontRegardless];
}

void NSWindow_inst_orderWindow_relativeTo_(void *id, unsigned long place, long otherWin) {
	[(NSWindow*)id
		orderWindow: place
		relativeTo: otherWin];
}

void NSWindow_inst_makeKeyWindow(void *id) {
	[(NSWindow*)id
		makeKeyWindow];
}

void NSWindow_inst_makeKeyAndOrderFront_(void *id, void* sender) {
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

void NSWindow_inst_toggleToolbarShown_(void *id, void* sender) {
	[(NSWindow*)id
		toggleToolbarShown: sender];
}

void NSWindow_inst_runToolbarCustomizationPalette_(void *id, void* sender) {
	[(NSWindow*)id
		runToolbarCustomizationPalette: sender];
}

void NSWindow_inst_addChildWindow_ordered_(void *id, void* childWin, unsigned long place) {
	[(NSWindow*)id
		addChildWindow: childWin
		ordered: place];
}

void NSWindow_inst_removeChildWindow_(void *id, void* childWin) {
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

void* NSWindow_inst_fieldEditor_forObject_(void *id, BOOL createFlag, void* object) {
	return [(NSWindow*)id
		fieldEditor: createFlag
		forObject: object];
}

void NSWindow_inst_endEditingFor_(void *id, void* object) {
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

void NSWindow_inst_invalidateCursorRectsForView_(void *id, void* view) {
	[(NSWindow*)id
		invalidateCursorRectsForView: view];
}

void NSWindow_inst_resetCursorRects(void *id) {
	[(NSWindow*)id
		resetCursorRects];
}

void NSWindow_inst_removeTitlebarAccessoryViewControllerAtIndex_(void *id, long index) {
	[(NSWindow*)id
		removeTitlebarAccessoryViewControllerAtIndex: index];
}

void NSWindow_inst_addTabbedWindow_ordered_(void *id, void* window, unsigned long ordered) {
	[(NSWindow*)id
		addTabbedWindow: window
		ordered: ordered];
}

void NSWindow_inst_mergeAllWindows_(void *id, void* sender) {
	[(NSWindow*)id
		mergeAllWindows: sender];
}

void NSWindow_inst_selectNextTab_(void *id, void* sender) {
	[(NSWindow*)id
		selectNextTab: sender];
}

void NSWindow_inst_selectPreviousTab_(void *id, void* sender) {
	[(NSWindow*)id
		selectPreviousTab: sender];
}

void NSWindow_inst_moveTabToNewWindow_(void *id, void* sender) {
	[(NSWindow*)id
		moveTabToNewWindow: sender];
}

void NSWindow_inst_toggleTabBar_(void *id, void* sender) {
	[(NSWindow*)id
		toggleTabBar: sender];
}

void NSWindow_inst_toggleTabOverview_(void *id, void* sender) {
	[(NSWindow*)id
		toggleTabOverview: sender];
}

void NSWindow_inst_postEvent_atStart_(void *id, void* event, BOOL flag) {
	[(NSWindow*)id
		postEvent: event
		atStart: flag];
}

void NSWindow_inst_sendEvent_(void *id, void* event) {
	[(NSWindow*)id
		sendEvent: event];
}

BOOL NSWindow_inst_tryToPerform_with_(void *id, void* action, void* object) {
	return [(NSWindow*)id
		tryToPerform: action
		with: object];
}

void NSWindow_inst_selectKeyViewPrecedingView_(void *id, void* view) {
	[(NSWindow*)id
		selectKeyViewPrecedingView: view];
}

void NSWindow_inst_selectKeyViewFollowingView_(void *id, void* view) {
	[(NSWindow*)id
		selectKeyViewFollowingView: view];
}

void NSWindow_inst_selectPreviousKeyView_(void *id, void* sender) {
	[(NSWindow*)id
		selectPreviousKeyView: sender];
}

void NSWindow_inst_selectNextKeyView_(void *id, void* sender) {
	[(NSWindow*)id
		selectNextKeyView: sender];
}

void NSWindow_inst_recalculateKeyViewLoop(void *id) {
	[(NSWindow*)id
		recalculateKeyViewLoop];
}

void NSWindow_inst_performWindowDragWithEvent_(void *id, void* event) {
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

void NSWindow_inst_dragImage_at_offset_event_pasteboard_source_slideBack_(void *id, void* image, NSPoint baseLocation, NSSize initialOffset, void* event, void* pboard, void* sourceObj, BOOL slideFlag) {
	[(NSWindow*)id
		dragImage: image
		at: baseLocation
		offset: initialOffset
		event: event
		pasteboard: pboard
		source: sourceObj
		slideBack: slideFlag];
}

void NSWindow_inst_registerForDraggedTypes_(void *id, void* newTypes) {
	[(NSWindow*)id
		registerForDraggedTypes: newTypes];
}

void NSWindow_inst_unregisterDraggedTypes(void *id) {
	[(NSWindow*)id
		unregisterDraggedTypes];
}

NSRect NSWindow_inst_convertRectFromBacking_(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectFromBacking: rect];
}

NSRect NSWindow_inst_convertRectFromScreen_(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectFromScreen: rect];
}

NSPoint NSWindow_inst_convertPointFromBacking_(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointFromBacking: point];
}

NSPoint NSWindow_inst_convertPointFromScreen_(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointFromScreen: point];
}

NSRect NSWindow_inst_convertRectToBacking_(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectToBacking: rect];
}

NSRect NSWindow_inst_convertRectToScreen_(void *id, NSRect rect) {
	return [(NSWindow*)id
		convertRectToScreen: rect];
}

NSPoint NSWindow_inst_convertPointToBacking_(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointToBacking: point];
}

NSPoint NSWindow_inst_convertPointToScreen_(void *id, NSPoint point) {
	return [(NSWindow*)id
		convertPointToScreen: point];
}

void NSWindow_inst_setTitleWithRepresentedFilename_(void *id, void* filename) {
	[(NSWindow*)id
		setTitleWithRepresentedFilename: filename];
}

void NSWindow_inst_center(void *id) {
	[(NSWindow*)id
		center];
}

void NSWindow_inst_performClose_(void *id, void* sender) {
	[(NSWindow*)id
		performClose: sender];
}

void NSWindow_inst_close(void *id) {
	[(NSWindow*)id
		close];
}

void NSWindow_inst_performMiniaturize_(void *id, void* sender) {
	[(NSWindow*)id
		performMiniaturize: sender];
}

void NSWindow_inst_miniaturize_(void *id, void* sender) {
	[(NSWindow*)id
		miniaturize: sender];
}

void NSWindow_inst_deminiaturize_(void *id, void* sender) {
	[(NSWindow*)id
		deminiaturize: sender];
}

void NSWindow_inst_print_(void *id, void* sender) {
	[(NSWindow*)id
		print: sender];
}

void* NSWindow_inst_dataWithEPSInsideRect_(void *id, NSRect rect) {
	return [(NSWindow*)id
		dataWithEPSInsideRect: rect];
}

void* NSWindow_inst_dataWithPDFInsideRect_(void *id, NSRect rect) {
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

void NSWindow_inst_visualizeConstraints_(void *id, void* constraints) {
	[(NSWindow*)id
		visualizeConstraints: constraints];
}

void NSWindow_inst_setIsMiniaturized_(void *id, BOOL flag) {
	[(NSWindow*)id
		setIsMiniaturized: flag];
}

void NSWindow_inst_setIsVisible_(void *id, BOOL flag) {
	[(NSWindow*)id
		setIsVisible: flag];
}

void NSWindow_inst_setIsZoomed_(void *id, BOOL flag) {
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

void NSWindow_inst_setDelegate_(void *id, void* value) {
	[(NSWindow*)id
		setDelegate: value];
}

void* NSWindow_inst_contentViewController(void *id) {
	return [(NSWindow*)id
		contentViewController];
}

void NSWindow_inst_setContentViewController_(void *id, void* value) {
	[(NSWindow*)id
		setContentViewController: value];
}

void* NSWindow_inst_contentView(void *id) {
	return [(NSWindow*)id
		contentView];
}

void NSWindow_inst_setContentView_(void *id, void* value) {
	[(NSWindow*)id
		setContentView: value];
}

unsigned long NSWindow_inst_styleMask(void *id) {
	return [(NSWindow*)id
		styleMask];
}

void NSWindow_inst_setStyleMask_(void *id, unsigned long value) {
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

void NSWindow_inst_setAlphaValue_(void *id, double value) {
	[(NSWindow*)id
		setAlphaValue: value];
}

void* NSWindow_inst_backgroundColor(void *id) {
	return [(NSWindow*)id
		backgroundColor];
}

void NSWindow_inst_setBackgroundColor_(void *id, void* value) {
	[(NSWindow*)id
		setBackgroundColor: value];
}

BOOL NSWindow_inst_canHide(void *id) {
	return [(NSWindow*)id
		canHide];
}

void NSWindow_inst_setCanHide_(void *id, BOOL value) {
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

void NSWindow_inst_setHidesOnDeactivate_(void *id, BOOL value) {
	[(NSWindow*)id
		setHidesOnDeactivate: value];
}

unsigned long NSWindow_inst_collectionBehavior(void *id) {
	return [(NSWindow*)id
		collectionBehavior];
}

void NSWindow_inst_setCollectionBehavior_(void *id, unsigned long value) {
	[(NSWindow*)id
		setCollectionBehavior: value];
}

BOOL NSWindow_inst_isOpaque(void *id) {
	return [(NSWindow*)id
		isOpaque];
}

void NSWindow_inst_setOpaque_(void *id, BOOL value) {
	[(NSWindow*)id
		setOpaque: value];
}

BOOL NSWindow_inst_hasShadow(void *id) {
	return [(NSWindow*)id
		hasShadow];
}

void NSWindow_inst_setHasShadow_(void *id, BOOL value) {
	[(NSWindow*)id
		setHasShadow: value];
}

BOOL NSWindow_inst_preventsApplicationTerminationWhenModal(void *id) {
	return [(NSWindow*)id
		preventsApplicationTerminationWhenModal];
}

void NSWindow_inst_setPreventsApplicationTerminationWhenModal_(void *id, BOOL value) {
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

void NSWindow_inst_setCanBecomeVisibleWithoutLogin_(void *id, BOOL value) {
	[(NSWindow*)id
		setCanBecomeVisibleWithoutLogin: value];
}

unsigned long NSWindow_inst_backingType(void *id) {
	return [(NSWindow*)id
		backingType];
}

void NSWindow_inst_setBackingType_(void *id, unsigned long value) {
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

void NSWindow_inst_setAspectRatio_(void *id, NSSize value) {
	[(NSWindow*)id
		setAspectRatio: value];
}

NSSize NSWindow_inst_minSize(void *id) {
	return [(NSWindow*)id
		minSize];
}

void NSWindow_inst_setMinSize_(void *id, NSSize value) {
	[(NSWindow*)id
		setMinSize: value];
}

NSSize NSWindow_inst_maxSize(void *id) {
	return [(NSWindow*)id
		maxSize];
}

void NSWindow_inst_setMaxSize_(void *id, NSSize value) {
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

void NSWindow_inst_setResizeIncrements_(void *id, NSSize value) {
	[(NSWindow*)id
		setResizeIncrements: value];
}

BOOL NSWindow_inst_preservesContentDuringLiveResize(void *id) {
	return [(NSWindow*)id
		preservesContentDuringLiveResize];
}

void NSWindow_inst_setPreservesContentDuringLiveResize_(void *id, BOOL value) {
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

void NSWindow_inst_setContentAspectRatio_(void *id, NSSize value) {
	[(NSWindow*)id
		setContentAspectRatio: value];
}

NSSize NSWindow_inst_contentMinSize(void *id) {
	return [(NSWindow*)id
		contentMinSize];
}

void NSWindow_inst_setContentMinSize_(void *id, NSSize value) {
	[(NSWindow*)id
		setContentMinSize: value];
}

NSSize NSWindow_inst_contentMaxSize(void *id) {
	return [(NSWindow*)id
		contentMaxSize];
}

void NSWindow_inst_setContentMaxSize_(void *id, NSSize value) {
	[(NSWindow*)id
		setContentMaxSize: value];
}

NSSize NSWindow_inst_contentResizeIncrements(void *id) {
	return [(NSWindow*)id
		contentResizeIncrements];
}

void NSWindow_inst_setContentResizeIncrements_(void *id, NSSize value) {
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

void NSWindow_inst_setMaxFullScreenContentSize_(void *id, NSSize value) {
	[(NSWindow*)id
		setMaxFullScreenContentSize: value];
}

NSSize NSWindow_inst_minFullScreenContentSize(void *id) {
	return [(NSWindow*)id
		minFullScreenContentSize];
}

void NSWindow_inst_setMinFullScreenContentSize_(void *id, NSSize value) {
	[(NSWindow*)id
		setMinFullScreenContentSize: value];
}

long NSWindow_inst_level(void *id) {
	return [(NSWindow*)id
		level];
}

void NSWindow_inst_setLevel_(void *id, long value) {
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

void NSWindow_inst_setParentWindow_(void *id, void* value) {
	[(NSWindow*)id
		setParentWindow: value];
}

BOOL NSWindow_inst_isExcludedFromWindowsMenu(void *id) {
	return [(NSWindow*)id
		isExcludedFromWindowsMenu];
}

void NSWindow_inst_setExcludedFromWindowsMenu_(void *id, BOOL value) {
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

void NSWindow_inst_setShowsToolbarButton_(void *id, BOOL value) {
	[(NSWindow*)id
		setShowsToolbarButton: value];
}

BOOL NSWindow_inst_titlebarAppearsTransparent(void *id) {
	return [(NSWindow*)id
		titlebarAppearsTransparent];
}

void NSWindow_inst_setTitlebarAppearsTransparent_(void *id, BOOL value) {
	[(NSWindow*)id
		setTitlebarAppearsTransparent: value];
}

void* NSWindow_inst_titlebarAccessoryViewControllers(void *id) {
	return [(NSWindow*)id
		titlebarAccessoryViewControllers];
}

void NSWindow_inst_setTitlebarAccessoryViewControllers_(void *id, void* value) {
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

void NSWindow_inst_setAllowsToolTipsWhenApplicationIsInactive_(void *id, BOOL value) {
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

void NSWindow_inst_setInitialFirstResponder_(void *id, void* value) {
	[(NSWindow*)id
		setInitialFirstResponder: value];
}

BOOL NSWindow_inst_autorecalculatesKeyViewLoop(void *id) {
	return [(NSWindow*)id
		autorecalculatesKeyViewLoop];
}

void NSWindow_inst_setAutorecalculatesKeyViewLoop_(void *id, BOOL value) {
	[(NSWindow*)id
		setAutorecalculatesKeyViewLoop: value];
}

BOOL NSWindow_inst_acceptsMouseMovedEvents(void *id) {
	return [(NSWindow*)id
		acceptsMouseMovedEvents];
}

void NSWindow_inst_setAcceptsMouseMovedEvents_(void *id, BOOL value) {
	[(NSWindow*)id
		setAcceptsMouseMovedEvents: value];
}

BOOL NSWindow_inst_ignoresMouseEvents(void *id) {
	return [(NSWindow*)id
		ignoresMouseEvents];
}

void NSWindow_inst_setIgnoresMouseEvents_(void *id, BOOL value) {
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

void NSWindow_inst_setRestorable_(void *id, BOOL value) {
	[(NSWindow*)id
		setRestorable: value];
}

BOOL NSWindow_inst_viewsNeedDisplay(void *id) {
	return [(NSWindow*)id
		viewsNeedDisplay];
}

void NSWindow_inst_setViewsNeedDisplay_(void *id, BOOL value) {
	[(NSWindow*)id
		setViewsNeedDisplay: value];
}

BOOL NSWindow_inst_allowsConcurrentViewDrawing(void *id) {
	return [(NSWindow*)id
		allowsConcurrentViewDrawing];
}

void NSWindow_inst_setAllowsConcurrentViewDrawing_(void *id, BOOL value) {
	[(NSWindow*)id
		setAllowsConcurrentViewDrawing: value];
}

BOOL NSWindow_inst_isDocumentEdited(void *id) {
	return [(NSWindow*)id
		isDocumentEdited];
}

void NSWindow_inst_setDocumentEdited_(void *id, BOOL value) {
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

void NSWindow_inst_setTitle_(void *id, void* value) {
	[(NSWindow*)id
		setTitle: value];
}

void* NSWindow_inst_subtitle(void *id) {
	return [(NSWindow*)id
		subtitle];
}

void NSWindow_inst_setSubtitle_(void *id, void* value) {
	[(NSWindow*)id
		setSubtitle: value];
}

long NSWindow_inst_titleVisibility(void *id) {
	return [(NSWindow*)id
		titleVisibility];
}

void NSWindow_inst_setTitleVisibility_(void *id, long value) {
	[(NSWindow*)id
		setTitleVisibility: value];
}

void* NSWindow_inst_representedFilename(void *id) {
	return [(NSWindow*)id
		representedFilename];
}

void NSWindow_inst_setRepresentedFilename_(void *id, void* value) {
	[(NSWindow*)id
		setRepresentedFilename: value];
}

void* NSWindow_inst_representedURL(void *id) {
	return [(NSWindow*)id
		representedURL];
}

void NSWindow_inst_setRepresentedURL_(void *id, void* value) {
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

void NSWindow_inst_setDisplaysWhenScreenProfileChanges_(void *id, BOOL value) {
	[(NSWindow*)id
		setDisplaysWhenScreenProfileChanges: value];
}

BOOL NSWindow_inst_isMovableByWindowBackground(void *id) {
	return [(NSWindow*)id
		isMovableByWindowBackground];
}

void NSWindow_inst_setMovableByWindowBackground_(void *id, BOOL value) {
	[(NSWindow*)id
		setMovableByWindowBackground: value];
}

BOOL NSWindow_inst_isMovable(void *id) {
	return [(NSWindow*)id
		isMovable];
}

void NSWindow_inst_setMovable_(void *id, BOOL value) {
	[(NSWindow*)id
		setMovable: value];
}

BOOL NSWindow_inst_isReleasedWhenClosed(void *id) {
	return [(NSWindow*)id
		isReleasedWhenClosed];
}

void NSWindow_inst_setReleasedWhenClosed_(void *id, BOOL value) {
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

void NSWindow_inst_setMiniwindowImage_(void *id, void* value) {
	[(NSWindow*)id
		setMiniwindowImage: value];
}

void* NSWindow_inst_miniwindowTitle(void *id) {
	return [(NSWindow*)id
		miniwindowTitle];
}

void NSWindow_inst_setMiniwindowTitle_(void *id, void* value) {
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

void NSWindow_inst_setOrderedIndex_(void *id, long value) {
	[(NSWindow*)id
		setOrderedIndex: value];
}

void* NSColor_inst_blendedColorWithFraction_ofColor_(void *id, double fraction, void* color) {
	return [(NSColor*)id
		blendedColorWithFraction: fraction
		ofColor: color];
}

void* NSColor_inst_colorWithAlphaComponent_(void *id, double alpha) {
	return [(NSColor*)id
		colorWithAlphaComponent: alpha];
}

void* NSColor_inst_highlightWithLevel_(void *id, double val) {
	return [(NSColor*)id
		highlightWithLevel: val];
}

void* NSColor_inst_shadowWithLevel_(void *id, double val) {
	return [(NSColor*)id
		shadowWithLevel: val];
}

void NSColor_inst_writeToPasteboard_(void *id, void* pasteBoard) {
	[(NSColor*)id
		writeToPasteboard: pasteBoard];
}

void NSColor_inst_drawSwatchInRect_(void *id, NSRect rect) {
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

void* NSTextView_inst_initWithFrame_textContainer_(void *id, NSRect frameRect, void* container) {
	return [(NSTextView*)id
		initWithFrame: frameRect
		textContainer: container];
}

void* NSTextView_inst_initWithFrame_(void *id, NSRect frameRect) {
	return [(NSTextView*)id
		initWithFrame: frameRect];
}

void NSTextView_inst_replaceTextContainer_(void *id, void* newContainer) {
	[(NSTextView*)id
		replaceTextContainer: newContainer];
}

void NSTextView_inst_invalidateTextContainerOrigin(void *id) {
	[(NSTextView*)id
		invalidateTextContainerOrigin];
}

void NSTextView_inst_changeDocumentBackgroundColor_(void *id, void* sender) {
	[(NSTextView*)id
		changeDocumentBackgroundColor: sender];
}

void NSTextView_inst_setNeedsDisplayInRect_avoidAdditionalLayout_(void *id, NSRect rect, BOOL flag) {
	[(NSTextView*)id
		setNeedsDisplayInRect: rect
		avoidAdditionalLayout: flag];
}

void NSTextView_inst_drawInsertionPointInRect_color_turnedOn_(void *id, NSRect rect, void* color, BOOL flag) {
	[(NSTextView*)id
		drawInsertionPointInRect: rect
		color: color
		turnedOn: flag];
}

void NSTextView_inst_drawViewBackgroundInRect_(void *id, NSRect rect) {
	[(NSTextView*)id
		drawViewBackgroundInRect: rect];
}

void NSTextView_inst_setConstrainedFrameSize_(void *id, NSSize desiredSize) {
	[(NSTextView*)id
		setConstrainedFrameSize: desiredSize];
}

void NSTextView_inst_cleanUpAfterDragOperation(void *id) {
	[(NSTextView*)id
		cleanUpAfterDragOperation];
}

void NSTextView_inst_outline_(void *id, void* sender) {
	[(NSTextView*)id
		outline: sender];
}

void NSTextView_inst_toggleAutomaticQuoteSubstitution_(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticQuoteSubstitution: sender];
}

void NSTextView_inst_toggleAutomaticLinkDetection_(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticLinkDetection: sender];
}

void NSTextView_inst_toggleAutomaticTextCompletion_(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticTextCompletion: sender];
}

void NSTextView_inst_updateInsertionPointStateAndRestartTimer_(void *id, BOOL restartFlag) {
	[(NSTextView*)id
		updateInsertionPointStateAndRestartTimer: restartFlag];
}

unsigned long NSTextView_inst_characterIndexForInsertionAtPoint_(void *id, NSPoint point) {
	return [(NSTextView*)id
		characterIndexForInsertionAtPoint: point];
}

void NSTextView_inst_updateCandidates(void *id) {
	[(NSTextView*)id
		updateCandidates];
}

BOOL NSTextView_inst_readSelectionFromPasteboard_(void *id, void* pboard) {
	return [(NSTextView*)id
		readSelectionFromPasteboard: pboard];
}

BOOL NSTextView_inst_writeSelectionToPasteboard_types_(void *id, void* pboard, void* types) {
	return [(NSTextView*)id
		writeSelectionToPasteboard: pboard
		types: types];
}

void NSTextView_inst_alignJustified_(void *id, void* sender) {
	[(NSTextView*)id
		alignJustified: sender];
}

void NSTextView_inst_changeAttributes_(void *id, void* sender) {
	[(NSTextView*)id
		changeAttributes: sender];
}

void NSTextView_inst_changeColor_(void *id, void* sender) {
	[(NSTextView*)id
		changeColor: sender];
}

void NSTextView_inst_useStandardKerning_(void *id, void* sender) {
	[(NSTextView*)id
		useStandardKerning: sender];
}

void NSTextView_inst_lowerBaseline_(void *id, void* sender) {
	[(NSTextView*)id
		lowerBaseline: sender];
}

void NSTextView_inst_raiseBaseline_(void *id, void* sender) {
	[(NSTextView*)id
		raiseBaseline: sender];
}

void NSTextView_inst_turnOffKerning_(void *id, void* sender) {
	[(NSTextView*)id
		turnOffKerning: sender];
}

void NSTextView_inst_loosenKerning_(void *id, void* sender) {
	[(NSTextView*)id
		loosenKerning: sender];
}

void NSTextView_inst_tightenKerning_(void *id, void* sender) {
	[(NSTextView*)id
		tightenKerning: sender];
}

void NSTextView_inst_useStandardLigatures_(void *id, void* sender) {
	[(NSTextView*)id
		useStandardLigatures: sender];
}

void NSTextView_inst_turnOffLigatures_(void *id, void* sender) {
	[(NSTextView*)id
		turnOffLigatures: sender];
}

void NSTextView_inst_useAllLigatures_(void *id, void* sender) {
	[(NSTextView*)id
		useAllLigatures: sender];
}

void NSTextView_inst_clickedOnLink_atIndex_(void *id, void* link, unsigned long charIndex) {
	[(NSTextView*)id
		clickedOnLink: link
		atIndex: charIndex];
}

void NSTextView_inst_pasteAsPlainText_(void *id, void* sender) {
	[(NSTextView*)id
		pasteAsPlainText: sender];
}

void NSTextView_inst_pasteAsRichText_(void *id, void* sender) {
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

BOOL NSTextView_inst_shouldChangeTextInRanges_replacementStrings_(void *id, void* affectedRanges, void* replacementStrings) {
	return [(NSTextView*)id
		shouldChangeTextInRanges: affectedRanges
		replacementStrings: replacementStrings];
}

void NSTextView_inst_didChangeText(void *id) {
	[(NSTextView*)id
		didChangeText];
}

void NSTextView_inst_toggleSmartInsertDelete_(void *id, void* sender) {
	[(NSTextView*)id
		toggleSmartInsertDelete: sender];
}

void NSTextView_inst_toggleContinuousSpellChecking_(void *id, void* sender) {
	[(NSTextView*)id
		toggleContinuousSpellChecking: sender];
}

void NSTextView_inst_toggleGrammarChecking_(void *id, void* sender) {
	[(NSTextView*)id
		toggleGrammarChecking: sender];
}

void NSTextView_inst_orderFrontSharingServicePicker_(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontSharingServicePicker: sender];
}

BOOL NSTextView_inst_dragSelectionWithEvent_offset_slideBack_(void *id, void* event, NSSize mouseOffset, BOOL slideBack) {
	return [(NSTextView*)id
		dragSelectionWithEvent: event
		offset: mouseOffset
		slideBack: slideBack];
}

void NSTextView_inst_startSpeaking_(void *id, void* sender) {
	[(NSTextView*)id
		startSpeaking: sender];
}

void NSTextView_inst_stopSpeaking_(void *id, void* sender) {
	[(NSTextView*)id
		stopSpeaking: sender];
}

void NSTextView_inst_performFindPanelAction_(void *id, void* sender) {
	[(NSTextView*)id
		performFindPanelAction: sender];
}

void NSTextView_inst_orderFrontLinkPanel_(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontLinkPanel: sender];
}

void NSTextView_inst_orderFrontListPanel_(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontListPanel: sender];
}

void NSTextView_inst_orderFrontSpacingPanel_(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontSpacingPanel: sender];
}

void NSTextView_inst_orderFrontTablePanel_(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontTablePanel: sender];
}

void NSTextView_inst_orderFrontSubstitutionsPanel_(void *id, void* sender) {
	[(NSTextView*)id
		orderFrontSubstitutionsPanel: sender];
}

void NSTextView_inst_complete_(void *id, void* sender) {
	[(NSTextView*)id
		complete: sender];
}

void NSTextView_inst_checkTextInDocument_(void *id, void* sender) {
	[(NSTextView*)id
		checkTextInDocument: sender];
}

void NSTextView_inst_checkTextInSelection_(void *id, void* sender) {
	[(NSTextView*)id
		checkTextInSelection: sender];
}

void NSTextView_inst_toggleAutomaticDashSubstitution_(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticDashSubstitution: sender];
}

void NSTextView_inst_toggleAutomaticDataDetection_(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticDataDetection: sender];
}

void NSTextView_inst_toggleAutomaticSpellingCorrection_(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticSpellingCorrection: sender];
}

void NSTextView_inst_toggleAutomaticTextReplacement_(void *id, void* sender) {
	[(NSTextView*)id
		toggleAutomaticTextReplacement: sender];
}

void NSTextView_inst_updateQuickLookPreviewPanel(void *id) {
	[(NSTextView*)id
		updateQuickLookPreviewPanel];
}

void NSTextView_inst_toggleQuickLookPreviewPanel_(void *id, void* sender) {
	[(NSTextView*)id
		toggleQuickLookPreviewPanel: sender];
}

void* NSTextView_inst_quickLookPreviewableItemsInRanges_(void *id, void* ranges) {
	return [(NSTextView*)id
		quickLookPreviewableItemsInRanges: ranges];
}

void NSTextView_inst_changeLayoutOrientation_(void *id, void* sender) {
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

void NSTextView_inst_setDelegate_(void *id, void* value) {
	[(NSTextView*)id
		setDelegate: value];
}

void* NSTextView_inst_textContainer(void *id) {
	return [(NSTextView*)id
		textContainer];
}

void NSTextView_inst_setTextContainer_(void *id, void* value) {
	[(NSTextView*)id
		setTextContainer: value];
}

NSSize NSTextView_inst_textContainerInset(void *id) {
	return [(NSTextView*)id
		textContainerInset];
}

void NSTextView_inst_setTextContainerInset_(void *id, NSSize value) {
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

void NSTextView_inst_setBackgroundColor_(void *id, void* value) {
	[(NSTextView*)id
		setBackgroundColor: value];
}

BOOL NSTextView_inst_drawsBackground(void *id) {
	return [(NSTextView*)id
		drawsBackground];
}

void NSTextView_inst_setDrawsBackground_(void *id, BOOL value) {
	[(NSTextView*)id
		setDrawsBackground: value];
}

BOOL NSTextView_inst_allowsDocumentBackgroundColorChange(void *id) {
	return [(NSTextView*)id
		allowsDocumentBackgroundColorChange];
}

void NSTextView_inst_setAllowsDocumentBackgroundColorChange_(void *id, BOOL value) {
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

void NSTextView_inst_setAllowedInputSourceLocales_(void *id, void* value) {
	[(NSTextView*)id
		setAllowedInputSourceLocales: value];
}

BOOL NSTextView_inst_allowsUndo(void *id) {
	return [(NSTextView*)id
		allowsUndo];
}

void NSTextView_inst_setAllowsUndo_(void *id, BOOL value) {
	[(NSTextView*)id
		setAllowsUndo: value];
}

BOOL NSTextView_inst_isEditable(void *id) {
	return [(NSTextView*)id
		isEditable];
}

void NSTextView_inst_setEditable_(void *id, BOOL value) {
	[(NSTextView*)id
		setEditable: value];
}

BOOL NSTextView_inst_isSelectable(void *id) {
	return [(NSTextView*)id
		isSelectable];
}

void NSTextView_inst_setSelectable_(void *id, BOOL value) {
	[(NSTextView*)id
		setSelectable: value];
}

BOOL NSTextView_inst_isFieldEditor(void *id) {
	return [(NSTextView*)id
		isFieldEditor];
}

void NSTextView_inst_setFieldEditor_(void *id, BOOL value) {
	[(NSTextView*)id
		setFieldEditor: value];
}

BOOL NSTextView_inst_isRichText(void *id) {
	return [(NSTextView*)id
		isRichText];
}

void NSTextView_inst_setRichText_(void *id, BOOL value) {
	[(NSTextView*)id
		setRichText: value];
}

BOOL NSTextView_inst_importsGraphics(void *id) {
	return [(NSTextView*)id
		importsGraphics];
}

void NSTextView_inst_setImportsGraphics_(void *id, BOOL value) {
	[(NSTextView*)id
		setImportsGraphics: value];
}

BOOL NSTextView_inst_allowsImageEditing(void *id) {
	return [(NSTextView*)id
		allowsImageEditing];
}

void NSTextView_inst_setAllowsImageEditing_(void *id, BOOL value) {
	[(NSTextView*)id
		setAllowsImageEditing: value];
}

BOOL NSTextView_inst_isAutomaticQuoteSubstitutionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticQuoteSubstitutionEnabled];
}

void NSTextView_inst_setAutomaticQuoteSubstitutionEnabled_(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticQuoteSubstitutionEnabled: value];
}

BOOL NSTextView_inst_isAutomaticLinkDetectionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticLinkDetectionEnabled];
}

void NSTextView_inst_setAutomaticLinkDetectionEnabled_(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticLinkDetectionEnabled: value];
}

BOOL NSTextView_inst_displaysLinkToolTips(void *id) {
	return [(NSTextView*)id
		displaysLinkToolTips];
}

void NSTextView_inst_setDisplaysLinkToolTips_(void *id, BOOL value) {
	[(NSTextView*)id
		setDisplaysLinkToolTips: value];
}

BOOL NSTextView_inst_isAutomaticTextCompletionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticTextCompletionEnabled];
}

void NSTextView_inst_setAutomaticTextCompletionEnabled_(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticTextCompletionEnabled: value];
}

BOOL NSTextView_inst_usesAdaptiveColorMappingForDarkAppearance(void *id) {
	return [(NSTextView*)id
		usesAdaptiveColorMappingForDarkAppearance];
}

void NSTextView_inst_setUsesAdaptiveColorMappingForDarkAppearance_(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesAdaptiveColorMappingForDarkAppearance: value];
}

BOOL NSTextView_inst_usesRolloverButtonForSelection(void *id) {
	return [(NSTextView*)id
		usesRolloverButtonForSelection];
}

void NSTextView_inst_setUsesRolloverButtonForSelection_(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesRolloverButtonForSelection: value];
}

BOOL NSTextView_inst_usesRuler(void *id) {
	return [(NSTextView*)id
		usesRuler];
}

void NSTextView_inst_setUsesRuler_(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesRuler: value];
}

BOOL NSTextView_inst_isRulerVisible(void *id) {
	return [(NSTextView*)id
		isRulerVisible];
}

void NSTextView_inst_setRulerVisible_(void *id, BOOL value) {
	[(NSTextView*)id
		setRulerVisible: value];
}

BOOL NSTextView_inst_usesInspectorBar(void *id) {
	return [(NSTextView*)id
		usesInspectorBar];
}

void NSTextView_inst_setUsesInspectorBar_(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesInspectorBar: value];
}

void* NSTextView_inst_selectedRanges(void *id) {
	return [(NSTextView*)id
		selectedRanges];
}

void NSTextView_inst_setSelectedRanges_(void *id, void* value) {
	[(NSTextView*)id
		setSelectedRanges: value];
}

void* NSTextView_inst_insertionPointColor(void *id) {
	return [(NSTextView*)id
		insertionPointColor];
}

void NSTextView_inst_setInsertionPointColor_(void *id, void* value) {
	[(NSTextView*)id
		setInsertionPointColor: value];
}

void* NSTextView_inst_selectedTextAttributes(void *id) {
	return [(NSTextView*)id
		selectedTextAttributes];
}

void NSTextView_inst_setSelectedTextAttributes_(void *id, void* value) {
	[(NSTextView*)id
		setSelectedTextAttributes: value];
}

void* NSTextView_inst_markedTextAttributes(void *id) {
	return [(NSTextView*)id
		markedTextAttributes];
}

void NSTextView_inst_setMarkedTextAttributes_(void *id, void* value) {
	[(NSTextView*)id
		setMarkedTextAttributes: value];
}

void* NSTextView_inst_linkTextAttributes(void *id) {
	return [(NSTextView*)id
		linkTextAttributes];
}

void NSTextView_inst_setLinkTextAttributes_(void *id, void* value) {
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

void NSTextView_inst_setTypingAttributes_(void *id, void* value) {
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

void NSTextView_inst_setSmartInsertDeleteEnabled_(void *id, BOOL value) {
	[(NSTextView*)id
		setSmartInsertDeleteEnabled: value];
}

BOOL NSTextView_inst_isContinuousSpellCheckingEnabled(void *id) {
	return [(NSTextView*)id
		isContinuousSpellCheckingEnabled];
}

void NSTextView_inst_setContinuousSpellCheckingEnabled_(void *id, BOOL value) {
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

void NSTextView_inst_setGrammarCheckingEnabled_(void *id, BOOL value) {
	[(NSTextView*)id
		setGrammarCheckingEnabled: value];
}

BOOL NSTextView_inst_acceptsGlyphInfo(void *id) {
	return [(NSTextView*)id
		acceptsGlyphInfo];
}

void NSTextView_inst_setAcceptsGlyphInfo_(void *id, BOOL value) {
	[(NSTextView*)id
		setAcceptsGlyphInfo: value];
}

BOOL NSTextView_inst_usesFontPanel(void *id) {
	return [(NSTextView*)id
		usesFontPanel];
}

void NSTextView_inst_setUsesFontPanel_(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesFontPanel: value];
}

BOOL NSTextView_inst_usesFindPanel(void *id) {
	return [(NSTextView*)id
		usesFindPanel];
}

void NSTextView_inst_setUsesFindPanel_(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesFindPanel: value];
}

BOOL NSTextView_inst_isAutomaticDashSubstitutionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticDashSubstitutionEnabled];
}

void NSTextView_inst_setAutomaticDashSubstitutionEnabled_(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticDashSubstitutionEnabled: value];
}

BOOL NSTextView_inst_isAutomaticDataDetectionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticDataDetectionEnabled];
}

void NSTextView_inst_setAutomaticDataDetectionEnabled_(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticDataDetectionEnabled: value];
}

BOOL NSTextView_inst_isAutomaticSpellingCorrectionEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticSpellingCorrectionEnabled];
}

void NSTextView_inst_setAutomaticSpellingCorrectionEnabled_(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticSpellingCorrectionEnabled: value];
}

BOOL NSTextView_inst_isAutomaticTextReplacementEnabled(void *id) {
	return [(NSTextView*)id
		isAutomaticTextReplacementEnabled];
}

void NSTextView_inst_setAutomaticTextReplacementEnabled_(void *id, BOOL value) {
	[(NSTextView*)id
		setAutomaticTextReplacementEnabled: value];
}

BOOL NSTextView_inst_usesFindBar(void *id) {
	return [(NSTextView*)id
		usesFindBar];
}

void NSTextView_inst_setUsesFindBar_(void *id, BOOL value) {
	[(NSTextView*)id
		setUsesFindBar: value];
}

BOOL NSTextView_inst_isIncrementalSearchingEnabled(void *id) {
	return [(NSTextView*)id
		isIncrementalSearchingEnabled];
}

void NSTextView_inst_setIncrementalSearchingEnabled_(void *id, BOOL value) {
	[(NSTextView*)id
		setIncrementalSearchingEnabled: value];
}

BOOL NSTextView_inst_allowsCharacterPickerTouchBarItem(void *id) {
	return [(NSTextView*)id
		allowsCharacterPickerTouchBarItem];
}

void NSTextView_inst_setAllowsCharacterPickerTouchBarItem_(void *id, BOOL value) {
	[(NSTextView*)id
		setAllowsCharacterPickerTouchBarItem: value];
}

void* NSTextView_inst_font(void *id) {
	return [(NSTextView*)id
		font];
}

void NSTextView_inst_setFont_(void *id, void* value) {
	[(NSTextView*)id
		setFont: value];
}

void* NSView_inst_initWithFrame_(void *id, NSRect frameRect) {
	return [(NSView*)id
		initWithFrame: frameRect];
}

void NSView_inst_prepareForReuse(void *id) {
	[(NSView*)id
		prepareForReuse];
}

void NSView_inst_addSubview_(void *id, void* view) {
	[(NSView*)id
		addSubview: view];
}

void NSView_inst_addSubview_positioned_relativeTo_(void *id, void* view, unsigned long place, void* otherView) {
	[(NSView*)id
		addSubview: view
		positioned: place
		relativeTo: otherView];
}

void NSView_inst_didAddSubview_(void *id, void* subview) {
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

void NSView_inst_replaceSubview_with_(void *id, void* oldView, void* newView) {
	[(NSView*)id
		replaceSubview: oldView
		with: newView];
}

BOOL NSView_inst_isDescendantOf_(void *id, void* view) {
	return [(NSView*)id
		isDescendantOf: view];
}

void* NSView_inst_ancestorSharedWithView_(void *id, void* view) {
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

void NSView_inst_viewWillMoveToSuperview_(void *id, void* newSuperview) {
	[(NSView*)id
		viewWillMoveToSuperview: newSuperview];
}

void NSView_inst_viewWillMoveToWindow_(void *id, void* newWindow) {
	[(NSView*)id
		viewWillMoveToWindow: newWindow];
}

void NSView_inst_willRemoveSubview_(void *id, void* subview) {
	[(NSView*)id
		willRemoveSubview: subview];
}

void NSView_inst_setFrameOrigin_(void *id, NSPoint newOrigin) {
	[(NSView*)id
		setFrameOrigin: newOrigin];
}

void NSView_inst_setFrameSize_(void *id, NSSize newSize) {
	[(NSView*)id
		setFrameSize: newSize];
}

void NSView_inst_setBoundsOrigin_(void *id, NSPoint newOrigin) {
	[(NSView*)id
		setBoundsOrigin: newOrigin];
}

void NSView_inst_setBoundsSize_(void *id, NSSize newSize) {
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

void NSView_inst_drawRect_(void *id, NSRect dirtyRect) {
	[(NSView*)id
		drawRect: dirtyRect];
}

BOOL NSView_inst_needsToDrawRect_(void *id, NSRect rect) {
	return [(NSView*)id
		needsToDrawRect: rect];
}

void NSView_inst_print_(void *id, void* sender) {
	[(NSView*)id
		print: sender];
}

void NSView_inst_beginPageInRect_atPlacement_(void *id, NSRect rect, NSPoint location) {
	[(NSView*)id
		beginPageInRect: rect
		atPlacement: location];
}

void* NSView_inst_dataWithEPSInsideRect_(void *id, NSRect rect) {
	return [(NSView*)id
		dataWithEPSInsideRect: rect];
}

void* NSView_inst_dataWithPDFInsideRect_(void *id, NSRect rect) {
	return [(NSView*)id
		dataWithPDFInsideRect: rect];
}

void NSView_inst_writeEPSInsideRect_toPasteboard_(void *id, NSRect rect, void* pasteboard) {
	[(NSView*)id
		writeEPSInsideRect: rect
		toPasteboard: pasteboard];
}

void NSView_inst_writePDFInsideRect_toPasteboard_(void *id, NSRect rect, void* pasteboard) {
	[(NSView*)id
		writePDFInsideRect: rect
		toPasteboard: pasteboard];
}

void NSView_inst_drawPageBorderWithSize_(void *id, NSSize borderSize) {
	[(NSView*)id
		drawPageBorderWithSize: borderSize];
}

NSRect NSView_inst_rectForPage_(void *id, long page) {
	return [(NSView*)id
		rectForPage: page];
}

NSPoint NSView_inst_locationOfPrintRect_(void *id, NSRect rect) {
	return [(NSView*)id
		locationOfPrintRect: rect];
}

void NSView_inst_setNeedsDisplayInRect_(void *id, NSRect invalidRect) {
	[(NSView*)id
		setNeedsDisplayInRect: invalidRect];
}

void NSView_inst_display(void *id) {
	[(NSView*)id
		display];
}

void NSView_inst_displayRect_(void *id, NSRect rect) {
	[(NSView*)id
		displayRect: rect];
}

void NSView_inst_displayRectIgnoringOpacity_(void *id, NSRect rect) {
	[(NSView*)id
		displayRectIgnoringOpacity: rect];
}

void NSView_inst_displayIfNeeded(void *id) {
	[(NSView*)id
		displayIfNeeded];
}

void NSView_inst_displayIfNeededInRect_(void *id, NSRect rect) {
	[(NSView*)id
		displayIfNeededInRect: rect];
}

void NSView_inst_displayIfNeededIgnoringOpacity(void *id) {
	[(NSView*)id
		displayIfNeededIgnoringOpacity];
}

void NSView_inst_displayIfNeededInRectIgnoringOpacity_(void *id, NSRect rect) {
	[(NSView*)id
		displayIfNeededInRectIgnoringOpacity: rect];
}

void NSView_inst_translateRectsNeedingDisplayInRect_by_(void *id, NSRect clipRect, NSSize delta) {
	[(NSView*)id
		translateRectsNeedingDisplayInRect: clipRect
		by: delta];
}

void NSView_inst_viewWillDraw(void *id) {
	[(NSView*)id
		viewWillDraw];
}

NSPoint NSView_inst_convertPointFromBacking_(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointFromBacking: point];
}

NSPoint NSView_inst_convertPointToBacking_(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointToBacking: point];
}

NSPoint NSView_inst_convertPointFromLayer_(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointFromLayer: point];
}

NSPoint NSView_inst_convertPointToLayer_(void *id, NSPoint point) {
	return [(NSView*)id
		convertPointToLayer: point];
}

NSRect NSView_inst_convertRectFromBacking_(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectFromBacking: rect];
}

NSRect NSView_inst_convertRectToBacking_(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectToBacking: rect];
}

NSRect NSView_inst_convertRectFromLayer_(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectFromLayer: rect];
}

NSRect NSView_inst_convertRectToLayer_(void *id, NSRect rect) {
	return [(NSView*)id
		convertRectToLayer: rect];
}

NSSize NSView_inst_convertSizeFromBacking_(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeFromBacking: size];
}

NSSize NSView_inst_convertSizeToBacking_(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeToBacking: size];
}

NSSize NSView_inst_convertSizeFromLayer_(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeFromLayer: size];
}

NSSize NSView_inst_convertSizeToLayer_(void *id, NSSize size) {
	return [(NSView*)id
		convertSizeToLayer: size];
}

NSPoint NSView_inst_convertPoint_fromView_(void *id, NSPoint point, void* view) {
	return [(NSView*)id
		convertPoint: point
		fromView: view];
}

NSPoint NSView_inst_convertPoint_toView_(void *id, NSPoint point, void* view) {
	return [(NSView*)id
		convertPoint: point
		toView: view];
}

NSSize NSView_inst_convertSize_fromView_(void *id, NSSize size, void* view) {
	return [(NSView*)id
		convertSize: size
		fromView: view];
}

NSSize NSView_inst_convertSize_toView_(void *id, NSSize size, void* view) {
	return [(NSView*)id
		convertSize: size
		toView: view];
}

NSRect NSView_inst_convertRect_fromView_(void *id, NSRect rect, void* view) {
	return [(NSView*)id
		convertRect: rect
		fromView: view];
}

NSRect NSView_inst_convertRect_toView_(void *id, NSRect rect, void* view) {
	return [(NSView*)id
		convertRect: rect
		toView: view];
}

NSRect NSView_inst_centerScanRect_(void *id, NSRect rect) {
	return [(NSView*)id
		centerScanRect: rect];
}

void NSView_inst_translateOriginToPoint_(void *id, NSPoint translation) {
	[(NSView*)id
		translateOriginToPoint: translation];
}

void NSView_inst_scaleUnitSquareToSize_(void *id, NSSize newUnitSize) {
	[(NSView*)id
		scaleUnitSquareToSize: newUnitSize];
}

void NSView_inst_rotateByAngle_(void *id, double angle) {
	[(NSView*)id
		rotateByAngle: angle];
}

void NSView_inst_resizeSubviewsWithOldSize_(void *id, NSSize oldSize) {
	[(NSView*)id
		resizeSubviewsWithOldSize: oldSize];
}

void NSView_inst_resizeWithOldSuperviewSize_(void *id, NSSize oldSize) {
	[(NSView*)id
		resizeWithOldSuperviewSize: oldSize];
}

void NSView_inst_addConstraints_(void *id, void* constraints) {
	[(NSView*)id
		addConstraints: constraints];
}

void NSView_inst_removeConstraints_(void *id, void* constraints) {
	[(NSView*)id
		removeConstraints: constraints];
}

void NSView_inst_invalidateIntrinsicContentSize(void *id) {
	[(NSView*)id
		invalidateIntrinsicContentSize];
}

NSRect NSView_inst_alignmentRectForFrame_(void *id, NSRect frame) {
	return [(NSView*)id
		alignmentRectForFrame: frame];
}

NSRect NSView_inst_frameForAlignmentRect_(void *id, NSRect alignmentRect) {
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

void NSView_inst_setKeyboardFocusRingNeedsDisplayInRect_(void *id, NSRect rect) {
	[(NSView*)id
		setKeyboardFocusRingNeedsDisplayInRect: rect];
}

BOOL NSView_inst_enterFullScreenMode_withOptions_(void *id, void* screen, void* options) {
	return [(NSView*)id
		enterFullScreenMode: screen
		withOptions: options];
}

void NSView_inst_exitFullScreenModeWithOptions_(void *id, void* options) {
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

BOOL NSView_inst_acceptsFirstMouse_(void *id, void* event) {
	return [(NSView*)id
		acceptsFirstMouse: event];
}

void* NSView_inst_hitTest_(void *id, NSPoint point) {
	return [(NSView*)id
		hitTest: point];
}

BOOL NSView_inst_mouse_inRect_(void *id, NSPoint point, NSRect rect) {
	return [(NSView*)id
		mouse: point
		inRect: rect];
}

BOOL NSView_inst_performKeyEquivalent_(void *id, void* event) {
	return [(NSView*)id
		performKeyEquivalent: event];
}

void NSView_inst_prepareContentInRect_(void *id, NSRect rect) {
	[(NSView*)id
		prepareContentInRect: rect];
}

void NSView_inst_scrollPoint_(void *id, NSPoint point) {
	[(NSView*)id
		scrollPoint: point];
}

BOOL NSView_inst_scrollRectToVisible_(void *id, NSRect rect) {
	return [(NSView*)id
		scrollRectToVisible: rect];
}

BOOL NSView_inst_autoscroll_(void *id, void* event) {
	return [(NSView*)id
		autoscroll: event];
}

NSRect NSView_inst_adjustScroll_(void *id, NSRect newVisible) {
	return [(NSView*)id
		adjustScroll: newVisible];
}

void NSView_inst_registerForDraggedTypes_(void *id, void* newTypes) {
	[(NSView*)id
		registerForDraggedTypes: newTypes];
}

void NSView_inst_unregisterDraggedTypes(void *id) {
	[(NSView*)id
		unregisterDraggedTypes];
}

BOOL NSView_inst_shouldDelayWindowOrderingForEvent_(void *id, void* event) {
	return [(NSView*)id
		shouldDelayWindowOrderingForEvent: event];
}

NSRect NSView_inst_rectForSmartMagnificationAtPoint_inRect_(void *id, NSPoint location, NSRect visibleRect) {
	return [(NSView*)id
		rectForSmartMagnificationAtPoint: location
		inRect: visibleRect];
}

void NSView_inst_viewDidChangeBackingProperties(void *id) {
	[(NSView*)id
		viewDidChangeBackingProperties];
}

void* NSView_inst_viewWithTag_(void *id, long tag) {
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

void* NSView_inst_menuForEvent_(void *id, void* event) {
	return [(NSView*)id
		menuForEvent: event];
}

void NSView_inst_willOpenMenu_withEvent_(void *id, void* menu, void* event) {
	[(NSView*)id
		willOpenMenu: menu
		withEvent: event];
}

void NSView_inst_didCloseMenu_withEvent_(void *id, void* menu, void* event) {
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

void NSView_inst_showDefinitionForAttributedString_atPoint_(void *id, void* attrString, NSPoint textBaselineOrigin) {
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

void NSView_inst_setSubviews_(void *id, void* value) {
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

void NSView_inst_setFrame_(void *id, NSRect value) {
	[(NSView*)id
		setFrame: value];
}

double NSView_inst_frameRotation(void *id) {
	return [(NSView*)id
		frameRotation];
}

void NSView_inst_setFrameRotation_(void *id, double value) {
	[(NSView*)id
		setFrameRotation: value];
}

NSRect NSView_inst_bounds(void *id) {
	return [(NSView*)id
		bounds];
}

void NSView_inst_setBounds_(void *id, NSRect value) {
	[(NSView*)id
		setBounds: value];
}

double NSView_inst_boundsRotation(void *id) {
	return [(NSView*)id
		boundsRotation];
}

void NSView_inst_setBoundsRotation_(void *id, double value) {
	[(NSView*)id
		setBoundsRotation: value];
}

BOOL NSView_inst_wantsLayer(void *id) {
	return [(NSView*)id
		wantsLayer];
}

void NSView_inst_setWantsLayer_(void *id, BOOL value) {
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

void NSView_inst_setLayer_(void *id, void* value) {
	[(NSView*)id
		setLayer: value];
}

BOOL NSView_inst_canDrawSubviewsIntoLayer(void *id) {
	return [(NSView*)id
		canDrawSubviewsIntoLayer];
}

void NSView_inst_setCanDrawSubviewsIntoLayer_(void *id, BOOL value) {
	[(NSView*)id
		setCanDrawSubviewsIntoLayer: value];
}

BOOL NSView_inst_layerUsesCoreImageFilters(void *id) {
	return [(NSView*)id
		layerUsesCoreImageFilters];
}

void NSView_inst_setLayerUsesCoreImageFilters_(void *id, BOOL value) {
	[(NSView*)id
		setLayerUsesCoreImageFilters: value];
}

double NSView_inst_alphaValue(void *id) {
	return [(NSView*)id
		alphaValue];
}

void NSView_inst_setAlphaValue_(void *id, double value) {
	[(NSView*)id
		setAlphaValue: value];
}

double NSView_inst_frameCenterRotation(void *id) {
	return [(NSView*)id
		frameCenterRotation];
}

void NSView_inst_setFrameCenterRotation_(void *id, double value) {
	[(NSView*)id
		setFrameCenterRotation: value];
}

void* NSView_inst_backgroundFilters(void *id) {
	return [(NSView*)id
		backgroundFilters];
}

void NSView_inst_setBackgroundFilters_(void *id, void* value) {
	[(NSView*)id
		setBackgroundFilters: value];
}

void* NSView_inst_contentFilters(void *id) {
	return [(NSView*)id
		contentFilters];
}

void NSView_inst_setContentFilters_(void *id, void* value) {
	[(NSView*)id
		setContentFilters: value];
}

BOOL NSView_inst_canDrawConcurrently(void *id) {
	return [(NSView*)id
		canDrawConcurrently];
}

void NSView_inst_setCanDrawConcurrently_(void *id, BOOL value) {
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

void NSView_inst_setNeedsDisplay_(void *id, BOOL value) {
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

void NSView_inst_setAutoresizesSubviews_(void *id, BOOL value) {
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

void NSView_inst_setNeedsLayout_(void *id, BOOL value) {
	[(NSView*)id
		setNeedsLayout: value];
}

BOOL NSView_inst_needsUpdateConstraints(void *id) {
	return [(NSView*)id
		needsUpdateConstraints];
}

void NSView_inst_setNeedsUpdateConstraints_(void *id, BOOL value) {
	[(NSView*)id
		setNeedsUpdateConstraints: value];
}

BOOL NSView_inst_translatesAutoresizingMaskIntoConstraints(void *id) {
	return [(NSView*)id
		translatesAutoresizingMaskIntoConstraints];
}

void NSView_inst_setTranslatesAutoresizingMaskIntoConstraints_(void *id, BOOL value) {
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

void NSView_inst_setHidden_(void *id, BOOL value) {
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

void NSView_inst_setGestureRecognizers_(void *id, void* value) {
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

void NSView_inst_setWantsRestingTouches_(void *id, BOOL value) {
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

void NSView_inst_setNextKeyView_(void *id, void* value) {
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

void NSView_inst_setPreparedContentRect_(void *id, NSRect value) {
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

void NSView_inst_setPostsFrameChangedNotifications_(void *id, BOOL value) {
	[(NSView*)id
		setPostsFrameChangedNotifications: value];
}

BOOL NSView_inst_postsBoundsChangedNotifications(void *id) {
	return [(NSView*)id
		postsBoundsChangedNotifications];
}

void NSView_inst_setPostsBoundsChangedNotifications_(void *id, BOOL value) {
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

void NSView_inst_setToolTip_(void *id, void* value) {
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

void NSView_inst_setHorizontalContentSizeConstraintActive_(void *id, BOOL value) {
	[(NSView*)id
		setHorizontalContentSizeConstraintActive: value];
}

BOOL NSView_inst_isVerticalContentSizeConstraintActive(void *id) {
	return [(NSView*)id
		isVerticalContentSizeConstraintActive];
}

void NSView_inst_setVerticalContentSizeConstraintActive_(void *id, BOOL value) {
	[(NSView*)id
		setVerticalContentSizeConstraintActive: value];
}

void* NSView_inst_backgroundColor(void *id) {
	return [(NSView*)id
		backgroundColor];
}

void NSView_inst_setBackgroundColor_(void *id, void* value) {
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

func NSBundle_bundleWithURL_(
	url core.NSURLRef,
) (
	r0 NSBundle,
) {
	ret := C.NSBundle_type_bundleWithURL_(
		objc.RefPointer(url),
	)
	r0 = NSBundle_fromPointer(ret)
	return
}

func NSBundle_bundleWithPath_(
	path core.NSStringRef,
) (
	r0 NSBundle,
) {
	ret := C.NSBundle_type_bundleWithPath_(
		objc.RefPointer(path),
	)
	r0 = NSBundle_fromPointer(ret)
	return
}

func NSBundle_bundleWithIdentifier_(
	identifier core.NSStringRef,
) (
	r0 NSBundle,
) {
	ret := C.NSBundle_type_bundleWithIdentifier_(
		objc.RefPointer(identifier),
	)
	r0 = NSBundle_fromPointer(ret)
	return
}

func NSBundle_URLForResource_withExtension_subdirectory_inBundleWithURL_(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
	bundleURL core.NSURLRef,
) (
	r0 core.NSURL,
) {
	ret := C.NSBundle_type_URLForResource_withExtension_subdirectory_inBundleWithURL_(
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(bundleURL),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func NSBundle_URLsForResourcesWithExtension_subdirectory_inBundleWithURL_(
	ext core.NSStringRef,
	subpath core.NSStringRef,
	bundleURL core.NSURLRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_type_URLsForResourcesWithExtension_subdirectory_inBundleWithURL_(
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(bundleURL),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func NSBundle_pathForResource_ofType_inDirectory_(
	name core.NSStringRef,
	ext core.NSStringRef,
	bundlePath core.NSStringRef,
) (
	r0 core.NSString,
) {
	ret := C.NSBundle_type_pathForResource_ofType_inDirectory_(
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(bundlePath),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func NSBundle_pathsForResourcesOfType_inDirectory_(
	ext core.NSStringRef,
	bundlePath core.NSStringRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_type_pathsForResourcesOfType_inDirectory_(
		objc.RefPointer(ext),
		objc.RefPointer(bundlePath),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func NSBundle_preferredLocalizationsFromArray_(
	localizationsArray core.NSArrayRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_type_preferredLocalizationsFromArray_(
		objc.RefPointer(localizationsArray),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func NSBundle_preferredLocalizationsFromArray_forPreferences_(
	localizationsArray core.NSArrayRef,
	preferencesArray core.NSArrayRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_type_preferredLocalizationsFromArray_forPreferences_(
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

func NSSound_canInitWithPasteboard_(
	pasteboard NSPasteboardRef,
) (
	r0 bool,
) {
	ret := C.NSSound_type_canInitWithPasteboard_(
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

func NSApplication_detachDrawingThread_toTarget_withObject_(
	selector objc.Selector,
	target objc.Ref,
	argument objc.Ref,
) {
	C.NSApplication_type_detachDrawingThread_toTarget_withObject_(
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

func NSButton_checkboxWithTitle_target_action_(
	title core.NSStringRef,
	target objc.Ref,
	action objc.Selector,
) (
	r0 NSButton,
) {
	ret := C.NSButton_type_checkboxWithTitle_target_action_(
		objc.RefPointer(title),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)
	r0 = NSButton_fromPointer(ret)
	return
}

func NSButton_buttonWithImage_target_action_(
	image NSImageRef,
	target objc.Ref,
	action objc.Selector,
) (
	r0 NSButton,
) {
	ret := C.NSButton_type_buttonWithImage_target_action_(
		objc.RefPointer(image),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)
	r0 = NSButton_fromPointer(ret)
	return
}

func NSButton_radioButtonWithTitle_target_action_(
	title core.NSStringRef,
	target objc.Ref,
	action objc.Selector,
) (
	r0 NSButton,
) {
	ret := C.NSButton_type_radioButtonWithTitle_target_action_(
		objc.RefPointer(title),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)
	r0 = NSButton_fromPointer(ret)
	return
}

func NSButton_buttonWithTitle_image_target_action_(
	title core.NSStringRef,
	image NSImageRef,
	target objc.Ref,
	action objc.Selector,
) (
	r0 NSButton,
) {
	ret := C.NSButton_type_buttonWithTitle_image_target_action_(
		objc.RefPointer(title),
		objc.RefPointer(image),
		objc.RefPointer(target),
		action.SelectorAddress(),
	)
	r0 = NSButton_fromPointer(ret)
	return
}

func NSButton_buttonWithTitle_target_action_(
	title core.NSStringRef,
	target objc.Ref,
	action objc.Selector,
) (
	r0 NSButton,
) {
	ret := C.NSButton_type_buttonWithTitle_target_action_(
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

func NSEvent_eventWithEventRef_(
	eventRef unsafe.Pointer,
) (
	r0 NSEvent,
) {
	ret := C.NSEvent_type_eventWithEventRef_(
		eventRef,
	)
	r0 = NSEvent_fromPointer(ret)
	return
}

func NSEvent_stopPeriodicEvents() {
	C.NSEvent_type_stopPeriodicEvents()
	return
}

func NSEvent_removeMonitor_(
	eventMonitor objc.Ref,
) {
	C.NSEvent_type_removeMonitor_(
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

func NSEvent_setMouseCoalescingEnabled_(
	value bool,
) {
	C.NSEvent_type_setMouseCoalescingEnabled_(
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

func NSFont_fontWithName_size_(
	fontName core.NSStringRef,
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_fontWithName_size_(
		objc.RefPointer(fontName),
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_userFontOfSize_(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_userFontOfSize_(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_userFixedPitchFontOfSize_(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_userFixedPitchFontOfSize_(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_systemFontOfSize_(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_systemFontOfSize_(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_boldSystemFontOfSize_(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_boldSystemFontOfSize_(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_labelFontOfSize_(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_labelFontOfSize_(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_messageFontOfSize_(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_messageFontOfSize_(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_menuBarFontOfSize_(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_menuBarFontOfSize_(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_menuFontOfSize_(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_menuFontOfSize_(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_controlContentFontOfSize_(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_controlContentFontOfSize_(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_titleBarFontOfSize_(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_titleBarFontOfSize_(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_paletteFontOfSize_(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_paletteFontOfSize_(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_toolTipsFontOfSize_(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_type_toolTipsFontOfSize_(
		C.double(fontSize),
	)
	r0 = NSFont_fromPointer(ret)
	return
}

func NSFont_setUserFont_(
	font NSFontRef,
) {
	C.NSFont_type_setUserFont_(
		objc.RefPointer(font),
	)
	return
}

func NSFont_setUserFixedPitchFont_(
	font NSFontRef,
) {
	C.NSFont_type_setUserFixedPitchFont_(
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

func NSImage_imageWithSystemSymbolName_accessibilityDescription_(
	symbolName core.NSStringRef,
	description core.NSStringRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_type_imageWithSystemSymbolName_accessibilityDescription_(
		objc.RefPointer(symbolName),
		objc.RefPointer(description),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func NSImage_canInitWithPasteboard_(
	pasteboard NSPasteboardRef,
) (
	r0 bool,
) {
	ret := C.NSImage_type_canInitWithPasteboard_(
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

func NSImageView_imageViewWithImage_(
	image NSImageRef,
) (
	r0 NSImageView,
) {
	ret := C.NSImageView_type_imageViewWithImage_(
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

func NSPasteboard_pasteboardByFilteringFile_(
	filename core.NSStringRef,
) (
	r0 NSPasteboard,
) {
	ret := C.NSPasteboard_type_pasteboardByFilteringFile_(
		objc.RefPointer(filename),
	)
	r0 = NSPasteboard_fromPointer(ret)
	return
}

func NSPasteboard_pasteboardByFilteringTypesInPasteboard_(
	pboard NSPasteboardRef,
) (
	r0 NSPasteboard,
) {
	ret := C.NSPasteboard_type_pasteboardByFilteringTypesInPasteboard_(
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

func NSMenu_setMenuBarVisible_(
	visible bool,
) {
	C.NSMenu_type_setMenuBarVisible_(
		convertToObjCBool(visible),
	)
	return
}

func NSMenu_popUpContextMenu_withEvent_forView_(
	menu NSMenuRef,
	event NSEventRef,
	view NSViewRef,
) {
	C.NSMenu_type_popUpContextMenu_withEvent_forView_(
		objc.RefPointer(menu),
		objc.RefPointer(event),
		objc.RefPointer(view),
	)
	return
}

func NSMenu_popUpContextMenu_withEvent_forView_withFont_(
	menu NSMenuRef,
	event NSEventRef,
	view NSViewRef,
	font NSFontRef,
) {
	C.NSMenu_type_popUpContextMenu_withEvent_forView_withFont_(
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

func NSMenuItem_setUsesUserKeyEquivalents_(
	value bool,
) {
	C.NSMenuItem_type_setUsesUserKeyEquivalents_(
		convertToObjCBool(value),
	)
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

func NSWindow_windowWithContentViewController_(
	contentViewController NSViewControllerRef,
) (
	r0 NSWindow,
) {
	ret := C.NSWindow_type_windowWithContentViewController_(
		objc.RefPointer(contentViewController),
	)
	r0 = NSWindow_fromPointer(ret)
	return
}

func NSWindow_contentRectForFrameRect_styleMask_(
	fRect core.NSRect,
	style core.NSUInteger,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_type_contentRectForFrameRect_styleMask_(
		*(*C.NSRect)(unsafe.Pointer(&fRect)),
		C.ulong(style),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func NSWindow_frameRectForContentRect_styleMask_(
	cRect core.NSRect,
	style core.NSUInteger,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_type_frameRectForContentRect_styleMask_(
		*(*C.NSRect)(unsafe.Pointer(&cRect)),
		C.ulong(style),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func NSWindow_minFrameWidthWithTitle_styleMask_(
	title core.NSStringRef,
	style core.NSUInteger,
) (
	r0 core.CGFloat,
) {
	ret := C.NSWindow_type_minFrameWidthWithTitle_styleMask_(
		objc.RefPointer(title),
		C.ulong(style),
	)
	r0 = core.CGFloat(ret)
	return
}

func NSWindow_windowNumberAtPoint_belowWindowWithWindowNumber_(
	point core.NSPoint,
	windowNumber core.NSInteger,
) (
	r0 core.NSInteger,
) {
	ret := C.NSWindow_type_windowNumberAtPoint_belowWindowWithWindowNumber_(
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

func NSWindow_setAllowsAutomaticWindowTabbing_(
	value bool,
) {
	C.NSWindow_type_setAllowsAutomaticWindowTabbing_(
		convertToObjCBool(value),
	)
	return
}

func NSColor_alloc() (
	r0 NSColor,
) {
	ret := C.NSColor_type_alloc()
	r0 = NSColor_fromPointer(ret)
	return
}

func NSColor_colorFromPasteboard_(
	pasteBoard NSPasteboardRef,
) (
	r0 NSColor,
) {
	ret := C.NSColor_type_colorFromPasteboard_(
		objc.RefPointer(pasteBoard),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func NSColor_colorWithRed_green_blue_alpha_(
	red core.CGFloat,
	green core.CGFloat,
	blue core.CGFloat,
	alpha core.CGFloat,
) (
	r0 NSColor,
) {
	ret := C.NSColor_type_colorWithRed_green_blue_alpha_(
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

func NSColor_setIgnoresAlpha_(
	value bool,
) {
	C.NSColor_type_setIgnoresAlpha_(
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

func (x gen_NSBundle) InitWithURL__asNSBundle(
	url core.NSURLRef,
) (
	r0 NSBundle,
) {
	ret := C.NSBundle_inst_initWithURL_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)
	r0 = NSBundle_fromPointer(ret)
	return
}

func (x gen_NSBundle) InitWithPath__asNSBundle(
	path core.NSStringRef,
) (
	r0 NSBundle,
) {
	ret := C.NSBundle_inst_initWithPath_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)
	r0 = NSBundle_fromPointer(ret)
	return
}

func (x gen_NSBundle) LoadNibNamed_owner_options_(
	name core.NSStringRef,
	owner objc.Ref,
	options core.NSDictionaryRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_inst_loadNibNamed_owner_options_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(owner),
		objc.RefPointer(options),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSBundle) URLForResource_withExtension_subdirectory_(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
) (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_URLForResource_withExtension_subdirectory_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) URLForResource_withExtension_(
	name core.NSStringRef,
	ext core.NSStringRef,
) (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_URLForResource_withExtension_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) URLsForResourcesWithExtension_subdirectory_(
	ext core.NSStringRef,
	subpath core.NSStringRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_inst_URLsForResourcesWithExtension_subdirectory_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSBundle) URLForResource_withExtension_subdirectory_localization_(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_URLForResource_withExtension_subdirectory_localization_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) URLsForResourcesWithExtension_subdirectory_localization_(
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_inst_URLsForResourcesWithExtension_subdirectory_localization_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSBundle) PathForResource_ofType_(
	name core.NSStringRef,
	ext core.NSStringRef,
) (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_pathForResource_ofType_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) PathForResource_ofType_inDirectory_(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
) (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_pathForResource_ofType_inDirectory_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) PathForResource_ofType_inDirectory_forLocalization_(
	name core.NSStringRef,
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_pathForResource_ofType_inDirectory_forLocalization_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(name),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) PathsForResourcesOfType_inDirectory_(
	ext core.NSStringRef,
	subpath core.NSStringRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_inst_pathsForResourcesOfType_inDirectory_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSBundle) PathsForResourcesOfType_inDirectory_forLocalization_(
	ext core.NSStringRef,
	subpath core.NSStringRef,
	localizationName core.NSStringRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSBundle_inst_pathsForResourcesOfType_inDirectory_forLocalization_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ext),
		objc.RefPointer(subpath),
		objc.RefPointer(localizationName),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSBundle) LocalizedStringForKey_value_table_(
	key core.NSStringRef,
	value core.NSStringRef,
	tableName core.NSStringRef,
) (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_localizedStringForKey_value_table_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
		objc.RefPointer(value),
		objc.RefPointer(tableName),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) URLForAuxiliaryExecutable_(
	executableName core.NSStringRef,
) (
	r0 core.NSURL,
) {
	ret := C.NSBundle_inst_URLForAuxiliaryExecutable_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(executableName),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_NSBundle) PathForAuxiliaryExecutable_(
	executableName core.NSStringRef,
) (
	r0 core.NSString,
) {
	ret := C.NSBundle_inst_pathForAuxiliaryExecutable_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(executableName),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSBundle) ObjectForInfoDictionaryKey_(
	key core.NSStringRef,
) (
	r0 objc.Object,
) {
	ret := C.NSBundle_inst_objectForInfoDictionaryKey_(
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

func (x gen_NSBundle) LocalizedAttributedStringForKey_value_table_(
	key core.NSStringRef,
	value core.NSStringRef,
	tableName core.NSStringRef,
) (
	r0 core.NSAttributedString,
) {
	ret := C.NSBundle_inst_localizedAttributedStringForKey_value_table_(
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

func (x gen_NSSound) InitWithContentsOfFile_byReference__asNSSound(
	path core.NSStringRef,
	byRef bool,
) (
	r0 NSSound,
) {
	ret := C.NSSound_inst_initWithContentsOfFile_byReference_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(byRef),
	)
	r0 = NSSound_fromPointer(ret)
	return
}

func (x gen_NSSound) InitWithContentsOfURL_byReference__asNSSound(
	url core.NSURLRef,
	byRef bool,
) (
	r0 NSSound,
) {
	ret := C.NSSound_inst_initWithContentsOfURL_byReference_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
		convertToObjCBool(byRef),
	)
	r0 = NSSound_fromPointer(ret)
	return
}

func (x gen_NSSound) InitWithData__asNSSound(
	data core.NSDataRef,
) (
	r0 NSSound,
) {
	ret := C.NSSound_inst_initWithData_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
	)
	r0 = NSSound_fromPointer(ret)
	return
}

func (x gen_NSSound) InitWithPasteboard__asNSSound(
	pasteboard NSPasteboardRef,
) (
	r0 NSSound,
) {
	ret := C.NSSound_inst_initWithPasteboard_(
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

func (x gen_NSSound) WriteToPasteboard_(
	pasteboard NSPasteboardRef,
) {
	C.NSSound_inst_writeToPasteboard_(
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

func (x gen_NSSound) SetDelegate_(
	value objc.Ref,
) {
	C.NSSound_inst_setDelegate_(
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

func (x gen_NSSound) SetLoops_(
	value bool,
) {
	C.NSSound_inst_setLoops_(
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

func (x gen_NSApplication) Stop_(
	sender objc.Ref,
) {
	C.NSApplication_inst_stop_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSApplication) SendEvent_(
	event NSEventRef,
) {
	C.NSApplication_inst_sendEvent_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	return
}

func (x gen_NSApplication) PostEvent_atStart_(
	event NSEventRef,
	flag bool,
) {
	C.NSApplication_inst_postEvent_atStart_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
		convertToObjCBool(flag),
	)
	return
}

func (x gen_NSApplication) TryToPerform_with_(
	action objc.Selector,
	object objc.Ref,
) (
	r0 bool,
) {
	ret := C.NSApplication_inst_tryToPerform_with_(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(object),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSApplication) SendAction_to_from_(
	action objc.Selector,
	target objc.Ref,
	sender objc.Ref,
) (
	r0 bool,
) {
	ret := C.NSApplication_inst_sendAction_to_from_(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(sender),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSApplication) TargetForAction_(
	action objc.Selector,
) (
	r0 objc.Object,
) {
	ret := C.NSApplication_inst_targetForAction_(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSApplication) TargetForAction_to_from_(
	action objc.Selector,
	target objc.Ref,
	sender objc.Ref,
) (
	r0 objc.Object,
) {
	ret := C.NSApplication_inst_targetForAction_to_from_(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(sender),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSApplication) Terminate_(
	sender objc.Ref,
) {
	C.NSApplication_inst_terminate_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSApplication) ReplyToApplicationShouldTerminate_(
	shouldTerminate bool,
) {
	C.NSApplication_inst_replyToApplicationShouldTerminate_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(shouldTerminate),
	)
	return
}

func (x gen_NSApplication) ActivateIgnoringOtherApps_(
	flag bool,
) {
	C.NSApplication_inst_activateIgnoringOtherApps_(
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

func (x gen_NSApplication) ToggleTouchBarCustomizationPalette_(
	sender objc.Ref,
) {
	C.NSApplication_inst_toggleTouchBarCustomizationPalette_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSApplication) CancelUserAttentionRequest_(
	request core.NSInteger,
) {
	C.NSApplication_inst_cancelUserAttentionRequest_(
		unsafe.Pointer(x.Pointer()),
		C.long(request),
	)
	return
}

func (x gen_NSApplication) RegisterUserInterfaceItemSearchHandler_(
	handler objc.Ref,
) {
	C.NSApplication_inst_registerUserInterfaceItemSearchHandler_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(handler),
	)
	return
}

func (x gen_NSApplication) UnregisterUserInterfaceItemSearchHandler_(
	handler objc.Ref,
) {
	C.NSApplication_inst_unregisterUserInterfaceItemSearchHandler_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(handler),
	)
	return
}

func (x gen_NSApplication) ShowHelp_(
	sender objc.Ref,
) {
	C.NSApplication_inst_showHelp_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSApplication) ActivateContextHelpMode_(
	sender objc.Ref,
) {
	C.NSApplication_inst_activateContextHelpMode_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSApplication) HideOtherApplications_(
	sender objc.Ref,
) {
	C.NSApplication_inst_hideOtherApplications_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSApplication) UnhideAllApplications_(
	sender objc.Ref,
) {
	C.NSApplication_inst_unhideAllApplications_(
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

func (x gen_NSApplication) SetActivationPolicy_(
	activationPolicy core.NSInteger,
) (
	r0 bool,
) {
	ret := C.NSApplication_inst_setActivationPolicy_(
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

func (x gen_NSApplication) SetDelegate_(
	value objc.Ref,
) {
	C.NSApplication_inst_setDelegate_(
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

func (x gen_NSApplication) SetApplicationIconImage_(
	value NSImageRef,
) {
	C.NSApplication_inst_setApplicationIconImage_(
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

func (x gen_NSApplication) SetHelpMenu_(
	value NSMenuRef,
) {
	C.NSApplication_inst_setHelpMenu_(
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

func (x gen_NSApplication) SetServicesProvider_(
	value objc.Ref,
) {
	C.NSApplication_inst_setServicesProvider_(
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

func (x gen_NSApplication) SetMainMenu_(
	value NSMenuRef,
) {
	C.NSApplication_inst_setMainMenu_(
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

func (x gen_NSControl) InitWithFrame__asNSControl(
	frameRect core.NSRect,
) (
	r0 NSControl,
) {
	ret := C.NSControl_inst_initWithFrame_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)
	r0 = NSControl_fromPointer(ret)
	return
}

func (x gen_NSControl) TakeDoubleValueFrom_(
	sender objc.Ref,
) {
	C.NSControl_inst_takeDoubleValueFrom_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSControl) TakeFloatValueFrom_(
	sender objc.Ref,
) {
	C.NSControl_inst_takeFloatValueFrom_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSControl) TakeIntValueFrom_(
	sender objc.Ref,
) {
	C.NSControl_inst_takeIntValueFrom_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSControl) TakeIntegerValueFrom_(
	sender objc.Ref,
) {
	C.NSControl_inst_takeIntegerValueFrom_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSControl) TakeObjectValueFrom_(
	sender objc.Ref,
) {
	C.NSControl_inst_takeObjectValueFrom_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSControl) TakeStringValueFrom_(
	sender objc.Ref,
) {
	C.NSControl_inst_takeStringValueFrom_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSControl) DrawWithExpansionFrame_inView_(
	contentFrame core.NSRect,
	view NSViewRef,
) {
	C.NSControl_inst_drawWithExpansionFrame_inView_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentFrame)),
		objc.RefPointer(view),
	)
	return
}

func (x gen_NSControl) ExpansionFrameWithFrame_(
	contentFrame core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSControl_inst_expansionFrameWithFrame_(
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

func (x gen_NSControl) EditWithFrame_editor_delegate_event_(
	rect core.NSRect,
	textObj NSTextRef,
	delegate objc.Ref,
	event NSEventRef,
) {
	C.NSControl_inst_editWithFrame_editor_delegate_event_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(textObj),
		objc.RefPointer(delegate),
		objc.RefPointer(event),
	)
	return
}

func (x gen_NSControl) EndEditing_(
	textObj NSTextRef,
) {
	C.NSControl_inst_endEditing_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(textObj),
	)
	return
}

func (x gen_NSControl) SelectWithFrame_editor_delegate_start_length_(
	rect core.NSRect,
	textObj NSTextRef,
	delegate objc.Ref,
	selStart core.NSInteger,
	selLength core.NSInteger,
) {
	C.NSControl_inst_selectWithFrame_editor_delegate_start_length_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(textObj),
		objc.RefPointer(delegate),
		C.long(selStart),
		C.long(selLength),
	)
	return
}

func (x gen_NSControl) SizeThatFits_(
	size core.NSSize,
) (
	r0 core.NSSize,
) {
	ret := C.NSControl_inst_sizeThatFits_(
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

func (x gen_NSControl) SendAction_to_(
	action objc.Selector,
	target objc.Ref,
) (
	r0 bool,
) {
	ret := C.NSControl_inst_sendAction_to_(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(target),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSControl) PerformClick_(
	sender objc.Ref,
) {
	C.NSControl_inst_performClick_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSControl) MouseDown_(
	event NSEventRef,
) {
	C.NSControl_inst_mouseDown_(
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

func (x gen_NSControl) SetEnabled_(
	value bool,
) {
	C.NSControl_inst_setEnabled_(
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

func (x gen_NSControl) SetIntValue_(
	value int32,
) {
	C.NSControl_inst_setIntValue_(
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

func (x gen_NSControl) SetIntegerValue_(
	value core.NSInteger,
) {
	C.NSControl_inst_setIntegerValue_(
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

func (x gen_NSControl) SetObjectValue_(
	value objc.Ref,
) {
	C.NSControl_inst_setObjectValue_(
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

func (x gen_NSControl) SetStringValue_(
	value core.NSStringRef,
) {
	C.NSControl_inst_setStringValue_(
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

func (x gen_NSControl) SetAttributedStringValue_(
	value core.NSAttributedStringRef,
) {
	C.NSControl_inst_setAttributedStringValue_(
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

func (x gen_NSControl) SetFont_(
	value NSFontRef,
) {
	C.NSControl_inst_setFont_(
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

func (x gen_NSControl) SetUsesSingleLineMode_(
	value bool,
) {
	C.NSControl_inst_setUsesSingleLineMode_(
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

func (x gen_NSControl) SetAllowsExpansionToolTips_(
	value bool,
) {
	C.NSControl_inst_setAllowsExpansionToolTips_(
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

func (x gen_NSControl) SetHighlighted_(
	value bool,
) {
	C.NSControl_inst_setHighlighted_(
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

func (x gen_NSControl) SetAction_(
	value objc.Selector,
) {
	C.NSControl_inst_setAction_(
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

func (x gen_NSControl) SetTarget_(
	value objc.Ref,
) {
	C.NSControl_inst_setTarget_(
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

func (x gen_NSControl) SetContinuous_(
	value bool,
) {
	C.NSControl_inst_setContinuous_(
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

func (x gen_NSControl) SetTag_(
	value core.NSInteger,
) {
	C.NSControl_inst_setTag_(
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

func (x gen_NSControl) SetRefusesFirstResponder_(
	value bool,
) {
	C.NSControl_inst_setRefusesFirstResponder_(
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

func (x gen_NSControl) SetIgnoresMultiClick_(
	value bool,
) {
	C.NSControl_inst_setIgnoresMultiClick_(
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

func (x gen_NSButton) CompressWithPrioritizedCompressionOptions_(
	prioritizedOptions core.NSArrayRef,
) {
	C.NSButton_inst_compressWithPrioritizedCompressionOptions_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(prioritizedOptions),
	)
	return
}

func (x gen_NSButton) MinimumSizeWithPrioritizedCompressionOptions_(
	prioritizedOptions core.NSArrayRef,
) (
	r0 core.NSSize,
) {
	ret := C.NSButton_inst_minimumSizeWithPrioritizedCompressionOptions_(
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

func (x gen_NSButton) Highlight_(
	flag bool,
) {
	C.NSButton_inst_highlight_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)
	return
}

func (x gen_NSButton) PerformKeyEquivalent_(
	key NSEventRef,
) (
	r0 bool,
) {
	ret := C.NSButton_inst_performKeyEquivalent_(
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

func (x gen_NSButton) SetContentTintColor_(
	value NSColorRef,
) {
	C.NSButton_inst_setContentTintColor_(
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

func (x gen_NSButton) SetHasDestructiveAction_(
	value bool,
) {
	C.NSButton_inst_setHasDestructiveAction_(
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

func (x gen_NSButton) SetAlternateTitle_(
	value core.NSStringRef,
) {
	C.NSButton_inst_setAlternateTitle_(
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

func (x gen_NSButton) SetAttributedTitle_(
	value core.NSAttributedStringRef,
) {
	C.NSButton_inst_setAttributedTitle_(
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

func (x gen_NSButton) SetAttributedAlternateTitle_(
	value core.NSAttributedStringRef,
) {
	C.NSButton_inst_setAttributedAlternateTitle_(
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

func (x gen_NSButton) SetTitle_(
	value core.NSStringRef,
) {
	C.NSButton_inst_setTitle_(
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

func (x gen_NSButton) SetSound_(
	value NSSoundRef,
) {
	C.NSButton_inst_setSound_(
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

func (x gen_NSButton) SetSpringLoaded_(
	value bool,
) {
	C.NSButton_inst_setSpringLoaded_(
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

func (x gen_NSButton) SetMaxAcceleratorLevel_(
	value core.NSInteger,
) {
	C.NSButton_inst_setMaxAcceleratorLevel_(
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

func (x gen_NSButton) SetImage_(
	value NSImageRef,
) {
	C.NSButton_inst_setImage_(
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

func (x gen_NSButton) SetAlternateImage_(
	value NSImageRef,
) {
	C.NSButton_inst_setAlternateImage_(
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

func (x gen_NSButton) SetBordered_(
	value bool,
) {
	C.NSButton_inst_setBordered_(
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

func (x gen_NSButton) SetTransparent_(
	value bool,
) {
	C.NSButton_inst_setTransparent_(
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

func (x gen_NSButton) SetBezelColor_(
	value NSColorRef,
) {
	C.NSButton_inst_setBezelColor_(
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

func (x gen_NSButton) SetShowsBorderOnlyWhileMouseInside_(
	value bool,
) {
	C.NSButton_inst_setShowsBorderOnlyWhileMouseInside_(
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

func (x gen_NSButton) SetImageHugsTitle_(
	value bool,
) {
	C.NSButton_inst_setImageHugsTitle_(
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

func (x gen_NSButton) SetAllowsMixedState_(
	value bool,
) {
	C.NSButton_inst_setAllowsMixedState_(
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

func (x gen_NSButton) SetState_(
	value core.NSInteger,
) {
	C.NSButton_inst_setState_(
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

func (x gen_NSButton) SetKeyEquivalent_(
	value core.NSStringRef,
) {
	C.NSButton_inst_setKeyEquivalent_(
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

func (x gen_NSFont) FontWithSize_(
	fontSize core.CGFloat,
) (
	r0 NSFont,
) {
	ret := C.NSFont_inst_fontWithSize_(
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

func (x gen_NSImage) InitByReferencingFile__asNSImage(
	fileName core.NSStringRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initByReferencingFile_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fileName),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) InitByReferencingURL__asNSImage(
	url core.NSURLRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initByReferencingURL_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) InitWithContentsOfFile__asNSImage(
	fileName core.NSStringRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initWithContentsOfFile_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(fileName),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) InitWithContentsOfURL__asNSImage(
	url core.NSURLRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initWithContentsOfURL_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) InitWithData__asNSImage(
	data core.NSDataRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initWithData_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) InitWithDataIgnoringOrientation__asNSImage(
	data core.NSDataRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initWithDataIgnoringOrientation_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) InitWithPasteboard__asNSImage(
	pasteboard NSPasteboardRef,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initWithPasteboard_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pasteboard),
	)
	r0 = NSImage_fromPointer(ret)
	return
}

func (x gen_NSImage) InitWithSize__asNSImage(
	size core.NSSize,
) (
	r0 NSImage,
) {
	ret := C.NSImage_inst_initWithSize_(
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

func (x gen_NSImage) AddRepresentations_(
	imageReps core.NSArrayRef,
) {
	C.NSImage_inst_addRepresentations_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(imageReps),
	)
	return
}

func (x gen_NSImage) DrawInRect_(
	rect core.NSRect,
) {
	C.NSImage_inst_drawInRect_(
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

func (x gen_NSImage) LockFocusFlipped_(
	flipped bool,
) {
	C.NSImage_inst_lockFocusFlipped_(
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

func (x gen_NSImage) LayerContentsForContentsScale_(
	layerContentsScale core.CGFloat,
) (
	r0 objc.Object,
) {
	ret := C.NSImage_inst_layerContentsForContentsScale_(
		unsafe.Pointer(x.Pointer()),
		C.double(layerContentsScale),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSImage) RecommendedLayerContentsScale_(
	preferredContentsScale core.CGFloat,
) (
	r0 core.CGFloat,
) {
	ret := C.NSImage_inst_recommendedLayerContentsScale_(
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

func (x gen_NSImage) SetDelegate_(
	value objc.Ref,
) {
	C.NSImage_inst_setDelegate_(
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

func (x gen_NSImage) SetSize_(
	value core.NSSize,
) {
	C.NSImage_inst_setSize_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_NSImage) SetTemplate_(
	value bool,
) {
	C.NSImage_inst_setTemplate_(
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

func (x gen_NSImage) SetPrefersColorMatch_(
	value bool,
) {
	C.NSImage_inst_setPrefersColorMatch_(
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

func (x gen_NSImage) SetUsesEPSOnResolutionMismatch_(
	value bool,
) {
	C.NSImage_inst_setUsesEPSOnResolutionMismatch_(
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

func (x gen_NSImage) SetMatchesOnMultipleResolution_(
	value bool,
) {
	C.NSImage_inst_setMatchesOnMultipleResolution_(
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

func (x gen_NSImage) SetBackgroundColor_(
	value NSColorRef,
) {
	C.NSImage_inst_setBackgroundColor_(
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

func (x gen_NSImage) SetAlignmentRect_(
	value core.NSRect,
) {
	C.NSImage_inst_setAlignmentRect_(
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

func (x gen_NSImage) SetAccessibilityDescription_(
	value core.NSStringRef,
) {
	C.NSImage_inst_setAccessibilityDescription_(
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

func (x gen_NSImage) SetMatchesOnlyOnBestFittingAxis_(
	value bool,
) {
	C.NSImage_inst_setMatchesOnlyOnBestFittingAxis_(
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

func (x gen_NSImageView) SetImage_(
	value NSImageRef,
) {
	C.NSImageView_inst_setImage_(
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

func (x gen_NSImageView) SetAnimates_(
	value bool,
) {
	C.NSImageView_inst_setAnimates_(
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

func (x gen_NSImageView) SetEditable_(
	value bool,
) {
	C.NSImageView_inst_setEditable_(
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

func (x gen_NSImageView) SetAllowsCutCopyPaste_(
	value bool,
) {
	C.NSImageView_inst_setAllowsCutCopyPaste_(
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

func (x gen_NSImageView) SetContentTintColor_(
	value NSColorRef,
) {
	C.NSImageView_inst_setContentTintColor_(
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

func (x gen_NSNib) InitWithNibData_bundle__asNSNib(
	nibData core.NSDataRef,
	bundle NSBundleRef,
) (
	r0 NSNib,
) {
	ret := C.NSNib_inst_initWithNibData_bundle_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(nibData),
		objc.RefPointer(bundle),
	)
	r0 = NSNib_fromPointer(ret)
	return
}

func (x gen_NSNib) InstantiateWithOwner_topLevelObjects_(
	owner objc.Ref,
	topLevelObjects core.NSArrayRef,
) (
	r0 bool,
) {
	ret := C.NSNib_inst_instantiateWithOwner_topLevelObjects_(
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

func (x gen_NSPasteboard) WriteObjects_(
	objects core.NSArrayRef,
) (
	r0 bool,
) {
	ret := C.NSPasteboard_inst_writeObjects_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(objects),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSPasteboard) ReadObjectsForClasses_options_(
	classArray core.NSArrayRef,
	options core.NSDictionaryRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSPasteboard_inst_readObjectsForClasses_options_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(classArray),
		objc.RefPointer(options),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSPasteboard) CanReadItemWithDataConformingToTypes_(
	types core.NSArrayRef,
) (
	r0 bool,
) {
	ret := C.NSPasteboard_inst_canReadItemWithDataConformingToTypes_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(types),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSPasteboard) CanReadObjectForClasses_options_(
	classArray core.NSArrayRef,
	options core.NSDictionaryRef,
) (
	r0 bool,
) {
	ret := C.NSPasteboard_inst_canReadObjectForClasses_options_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(classArray),
		objc.RefPointer(options),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSPasteboard) DeclareTypes_owner_(
	newTypes core.NSArrayRef,
	newOwner objc.Ref,
) (
	r0 core.NSInteger,
) {
	ret := C.NSPasteboard_inst_declareTypes_owner_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newTypes),
		objc.RefPointer(newOwner),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSPasteboard) AddTypes_owner_(
	newTypes core.NSArrayRef,
	newOwner objc.Ref,
) (
	r0 core.NSInteger,
) {
	ret := C.NSPasteboard_inst_addTypes_owner_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newTypes),
		objc.RefPointer(newOwner),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSPasteboard) WriteFileContents_(
	filename core.NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSPasteboard_inst_writeFileContents_(
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

func (x gen_NSLayoutManager) AddTextContainer_(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_addTextContainer_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)
	return
}

func (x gen_NSLayoutManager) InsertTextContainer_atIndex_(
	container NSTextContainerRef,
	index core.NSUInteger,
) {
	C.NSLayoutManager_inst_insertTextContainer_atIndex_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
		C.ulong(index),
	)
	return
}

func (x gen_NSLayoutManager) RemoveTextContainerAtIndex_(
	index core.NSUInteger,
) {
	C.NSLayoutManager_inst_removeTextContainerAtIndex_(
		unsafe.Pointer(x.Pointer()),
		C.ulong(index),
	)
	return
}

func (x gen_NSLayoutManager) TextContainerChangedGeometry_(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_textContainerChangedGeometry_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)
	return
}

func (x gen_NSLayoutManager) TextContainerChangedTextView_(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_textContainerChangedTextView_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)
	return
}

func (x gen_NSLayoutManager) UsedRectForTextContainer_(
	container NSTextContainerRef,
) (
	r0 core.NSRect,
) {
	ret := C.NSLayoutManager_inst_usedRectForTextContainer_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSLayoutManager) EnsureLayoutForBoundingRect_inTextContainer_(
	bounds core.NSRect,
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_ensureLayoutForBoundingRect_inTextContainer_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&bounds)),
		objc.RefPointer(container),
	)
	return
}

func (x gen_NSLayoutManager) EnsureLayoutForTextContainer_(
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_ensureLayoutForTextContainer_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(container),
	)
	return
}

func (x gen_NSLayoutManager) CharacterIndexForGlyphAtIndex_(
	glyphIndex core.NSUInteger,
) (
	r0 core.NSUInteger,
) {
	ret := C.NSLayoutManager_inst_characterIndexForGlyphAtIndex_(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSLayoutManager) GlyphIndexForCharacterAtIndex_(
	charIndex core.NSUInteger,
) (
	r0 core.NSUInteger,
) {
	ret := C.NSLayoutManager_inst_glyphIndexForCharacterAtIndex_(
		unsafe.Pointer(x.Pointer()),
		C.ulong(charIndex),
	)
	r0 = core.NSUInteger(ret)
	return
}

func (x gen_NSLayoutManager) IsValidGlyphIndex_(
	glyphIndex core.NSUInteger,
) (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_isValidGlyphIndex_(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSLayoutManager) SetDrawsOutsideLineFragment_forGlyphAtIndex_(
	flag bool,
	glyphIndex core.NSUInteger,
) {
	C.NSLayoutManager_inst_setDrawsOutsideLineFragment_forGlyphAtIndex_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
		C.ulong(glyphIndex),
	)
	return
}

func (x gen_NSLayoutManager) SetExtraLineFragmentRect_usedRect_textContainer_(
	fragmentRect core.NSRect,
	usedRect core.NSRect,
	container NSTextContainerRef,
) {
	C.NSLayoutManager_inst_setExtraLineFragmentRect_usedRect_textContainer_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&fragmentRect)),
		*(*C.NSRect)(unsafe.Pointer(&usedRect)),
		objc.RefPointer(container),
	)
	return
}

func (x gen_NSLayoutManager) SetNotShownAttribute_forGlyphAtIndex_(
	flag bool,
	glyphIndex core.NSUInteger,
) {
	C.NSLayoutManager_inst_setNotShownAttribute_forGlyphAtIndex_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
		C.ulong(glyphIndex),
	)
	return
}

func (x gen_NSLayoutManager) AttachmentSizeForGlyphAtIndex_(
	glyphIndex core.NSUInteger,
) (
	r0 core.NSSize,
) {
	ret := C.NSLayoutManager_inst_attachmentSizeForGlyphAtIndex_(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSLayoutManager) DrawsOutsideLineFragmentForGlyphAtIndex_(
	glyphIndex core.NSUInteger,
) (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_drawsOutsideLineFragmentForGlyphAtIndex_(
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

func (x gen_NSLayoutManager) NotShownAttributeForGlyphAtIndex_(
	glyphIndex core.NSUInteger,
) (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_notShownAttributeForGlyphAtIndex_(
		unsafe.Pointer(x.Pointer()),
		C.ulong(glyphIndex),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSLayoutManager) LayoutManagerOwnsFirstResponderInWindow_(
	window NSWindowRef,
) (
	r0 bool,
) {
	ret := C.NSLayoutManager_inst_layoutManagerOwnsFirstResponderInWindow_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(window),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSLayoutManager) DefaultLineHeightForFont_(
	theFont NSFontRef,
) (
	r0 core.CGFloat,
) {
	ret := C.NSLayoutManager_inst_defaultLineHeightForFont_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(theFont),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_NSLayoutManager) DefaultBaselineOffsetForFont_(
	theFont NSFontRef,
) (
	r0 core.CGFloat,
) {
	ret := C.NSLayoutManager_inst_defaultBaselineOffsetForFont_(
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

func (x gen_NSLayoutManager) SetDelegate_(
	value objc.Ref,
) {
	C.NSLayoutManager_inst_setDelegate_(
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

func (x gen_NSLayoutManager) SetAllowsNonContiguousLayout_(
	value bool,
) {
	C.NSLayoutManager_inst_setAllowsNonContiguousLayout_(
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

func (x gen_NSLayoutManager) SetShowsInvisibleCharacters_(
	value bool,
) {
	C.NSLayoutManager_inst_setShowsInvisibleCharacters_(
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

func (x gen_NSLayoutManager) SetShowsControlCharacters_(
	value bool,
) {
	C.NSLayoutManager_inst_setShowsControlCharacters_(
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

func (x gen_NSLayoutManager) SetUsesFontLeading_(
	value bool,
) {
	C.NSLayoutManager_inst_setUsesFontLeading_(
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

func (x gen_NSLayoutManager) SetBackgroundLayoutEnabled_(
	value bool,
) {
	C.NSLayoutManager_inst_setBackgroundLayoutEnabled_(
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

func (x gen_NSLayoutManager) SetLimitsLayoutForSuspiciousContents_(
	value bool,
) {
	C.NSLayoutManager_inst_setLimitsLayoutForSuspiciousContents_(
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

func (x gen_NSLayoutManager) SetUsesDefaultHyphenation_(
	value bool,
) {
	C.NSLayoutManager_inst_setUsesDefaultHyphenation_(
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

func (x gen_NSMenu) InitWithTitle__asNSMenu(
	title core.NSStringRef,
) (
	r0 NSMenu,
) {
	ret := C.NSMenu_inst_initWithTitle_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(title),
	)
	r0 = NSMenu_fromPointer(ret)
	return
}

func (x gen_NSMenu) InsertItem_atIndex_(
	newItem NSMenuItemRef,
	index core.NSInteger,
) {
	C.NSMenu_inst_insertItem_atIndex_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newItem),
		C.long(index),
	)
	return
}

func (x gen_NSMenu) InsertItemWithTitle_action_keyEquivalent_atIndex_(
	string core.NSStringRef,
	selector objc.Selector,
	charCode core.NSStringRef,
	index core.NSInteger,
) (
	r0 NSMenuItem,
) {
	ret := C.NSMenu_inst_insertItemWithTitle_action_keyEquivalent_atIndex_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		selector.SelectorAddress(),
		objc.RefPointer(charCode),
		C.long(index),
	)
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func (x gen_NSMenu) AddItem_(
	newItem NSMenuItemRef,
) {
	C.NSMenu_inst_addItem_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newItem),
	)
	return
}

func (x gen_NSMenu) AddItemWithTitle_action_keyEquivalent_(
	string core.NSStringRef,
	selector objc.Selector,
	charCode core.NSStringRef,
) (
	r0 NSMenuItem,
) {
	ret := C.NSMenu_inst_addItemWithTitle_action_keyEquivalent_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		selector.SelectorAddress(),
		objc.RefPointer(charCode),
	)
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func (x gen_NSMenu) RemoveItem_(
	item NSMenuItemRef,
) {
	C.NSMenu_inst_removeItem_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
	)
	return
}

func (x gen_NSMenu) RemoveItemAtIndex_(
	index core.NSInteger,
) {
	C.NSMenu_inst_removeItemAtIndex_(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)
	return
}

func (x gen_NSMenu) ItemChanged_(
	item NSMenuItemRef,
) {
	C.NSMenu_inst_itemChanged_(
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

func (x gen_NSMenu) ItemWithTag_(
	tag core.NSInteger,
) (
	r0 NSMenuItem,
) {
	ret := C.NSMenu_inst_itemWithTag_(
		unsafe.Pointer(x.Pointer()),
		C.long(tag),
	)
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func (x gen_NSMenu) ItemWithTitle_(
	title core.NSStringRef,
) (
	r0 NSMenuItem,
) {
	ret := C.NSMenu_inst_itemWithTitle_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(title),
	)
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func (x gen_NSMenu) ItemAtIndex_(
	index core.NSInteger,
) (
	r0 NSMenuItem,
) {
	ret := C.NSMenu_inst_itemAtIndex_(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)
	r0 = NSMenuItem_fromPointer(ret)
	return
}

func (x gen_NSMenu) IndexOfItem_(
	item NSMenuItemRef,
) (
	r0 core.NSInteger,
) {
	ret := C.NSMenu_inst_indexOfItem_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(item),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSMenu) IndexOfItemWithTitle_(
	title core.NSStringRef,
) (
	r0 core.NSInteger,
) {
	ret := C.NSMenu_inst_indexOfItemWithTitle_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(title),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSMenu) IndexOfItemWithTag_(
	tag core.NSInteger,
) (
	r0 core.NSInteger,
) {
	ret := C.NSMenu_inst_indexOfItemWithTag_(
		unsafe.Pointer(x.Pointer()),
		C.long(tag),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSMenu) IndexOfItemWithTarget_andAction_(
	target objc.Ref,
	actionSelector objc.Selector,
) (
	r0 core.NSInteger,
) {
	ret := C.NSMenu_inst_indexOfItemWithTarget_andAction_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(target),
		actionSelector.SelectorAddress(),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSMenu) IndexOfItemWithRepresentedObject_(
	object objc.Ref,
) (
	r0 core.NSInteger,
) {
	ret := C.NSMenu_inst_indexOfItemWithRepresentedObject_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSMenu) IndexOfItemWithSubmenu_(
	submenu NSMenuRef,
) (
	r0 core.NSInteger,
) {
	ret := C.NSMenu_inst_indexOfItemWithSubmenu_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(submenu),
	)
	r0 = core.NSInteger(ret)
	return
}

func (x gen_NSMenu) SetSubmenu_forItem_(
	menu NSMenuRef,
	item NSMenuItemRef,
) {
	C.NSMenu_inst_setSubmenu_forItem_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(menu),
		objc.RefPointer(item),
	)
	return
}

func (x gen_NSMenu) SubmenuAction_(
	sender objc.Ref,
) {
	C.NSMenu_inst_submenuAction_(
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

func (x gen_NSMenu) PerformKeyEquivalent_(
	event NSEventRef,
) (
	r0 bool,
) {
	ret := C.NSMenu_inst_performKeyEquivalent_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSMenu) PerformActionForItemAtIndex_(
	index core.NSInteger,
) {
	C.NSMenu_inst_performActionForItemAtIndex_(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)
	return
}

func (x gen_NSMenu) PopUpMenuPositioningItem_atLocation_inView_(
	item NSMenuItemRef,
	location core.NSPoint,
	view NSViewRef,
) (
	r0 bool,
) {
	ret := C.NSMenu_inst_popUpMenuPositioningItem_atLocation_inView_(
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

func (x gen_NSMenu) SetItemArray_(
	value core.NSArrayRef,
) {
	C.NSMenu_inst_setItemArray_(
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

func (x gen_NSMenu) SetSupermenu_(
	value NSMenuRef,
) {
	C.NSMenu_inst_setSupermenu_(
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

func (x gen_NSMenu) SetAutoenablesItems_(
	value bool,
) {
	C.NSMenu_inst_setAutoenablesItems_(
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

func (x gen_NSMenu) SetFont_(
	value NSFontRef,
) {
	C.NSMenu_inst_setFont_(
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

func (x gen_NSMenu) SetTitle_(
	value core.NSStringRef,
) {
	C.NSMenu_inst_setTitle_(
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

func (x gen_NSMenu) SetMinimumWidth_(
	value core.CGFloat,
) {
	C.NSMenu_inst_setMinimumWidth_(
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

func (x gen_NSMenu) SetAllowsContextMenuPlugIns_(
	value bool,
) {
	C.NSMenu_inst_setAllowsContextMenuPlugIns_(
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

func (x gen_NSMenu) SetShowsStateColumn_(
	value bool,
) {
	C.NSMenu_inst_setShowsStateColumn_(
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

func (x gen_NSMenu) SetDelegate_(
	value objc.Ref,
) {
	C.NSMenu_inst_setDelegate_(
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

func (x gen_NSPopover) PerformClose_(
	sender objc.Ref,
) {
	C.NSPopover_inst_performClose_(
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

func (x gen_NSPopover) SetBehavior_(
	value core.NSInteger,
) {
	C.NSPopover_inst_setBehavior_(
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

func (x gen_NSPopover) SetPositioningRect_(
	value core.NSRect,
) {
	C.NSPopover_inst_setPositioningRect_(
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

func (x gen_NSPopover) SetAnimates_(
	value bool,
) {
	C.NSPopover_inst_setAnimates_(
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

func (x gen_NSPopover) SetContentSize_(
	value core.NSSize,
) {
	C.NSPopover_inst_setContentSize_(
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

func (x gen_NSMenuItem) InitWithTitle_action_keyEquivalent__asNSMenuItem(
	string core.NSStringRef,
	selector objc.Selector,
	charCode core.NSStringRef,
) (
	r0 NSMenuItem,
) {
	ret := C.NSMenuItem_inst_initWithTitle_action_keyEquivalent_(
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

func (x gen_NSMenuItem) SetEnabled_(
	value bool,
) {
	C.NSMenuItem_inst_setEnabled_(
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

func (x gen_NSMenuItem) SetHidden_(
	value bool,
) {
	C.NSMenuItem_inst_setHidden_(
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

func (x gen_NSMenuItem) SetTarget_(
	value objc.Ref,
) {
	C.NSMenuItem_inst_setTarget_(
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

func (x gen_NSMenuItem) SetAction_(
	value objc.Selector,
) {
	C.NSMenuItem_inst_setAction_(
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

func (x gen_NSMenuItem) SetTitle_(
	value core.NSStringRef,
) {
	C.NSMenuItem_inst_setTitle_(
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

func (x gen_NSMenuItem) SetAttributedTitle_(
	value core.NSAttributedStringRef,
) {
	C.NSMenuItem_inst_setAttributedTitle_(
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

func (x gen_NSMenuItem) SetTag_(
	value core.NSInteger,
) {
	C.NSMenuItem_inst_setTag_(
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

func (x gen_NSMenuItem) SetState_(
	value core.NSInteger,
) {
	C.NSMenuItem_inst_setState_(
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

func (x gen_NSMenuItem) SetImage_(
	value NSImageRef,
) {
	C.NSMenuItem_inst_setImage_(
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

func (x gen_NSMenuItem) SetOnStateImage_(
	value NSImageRef,
) {
	C.NSMenuItem_inst_setOnStateImage_(
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

func (x gen_NSMenuItem) SetOffStateImage_(
	value NSImageRef,
) {
	C.NSMenuItem_inst_setOffStateImage_(
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

func (x gen_NSMenuItem) SetMixedStateImage_(
	value NSImageRef,
) {
	C.NSMenuItem_inst_setMixedStateImage_(
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

func (x gen_NSMenuItem) SetSubmenu_(
	value NSMenuRef,
) {
	C.NSMenuItem_inst_setSubmenu_(
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

func (x gen_NSMenuItem) SetMenu_(
	value NSMenuRef,
) {
	C.NSMenuItem_inst_setMenu_(
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

func (x gen_NSMenuItem) SetKeyEquivalent_(
	value core.NSStringRef,
) {
	C.NSMenuItem_inst_setKeyEquivalent_(
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

func (x gen_NSMenuItem) SetAlternate_(
	value bool,
) {
	C.NSMenuItem_inst_setAlternate_(
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

func (x gen_NSMenuItem) SetIndentationLevel_(
	value core.NSInteger,
) {
	C.NSMenuItem_inst_setIndentationLevel_(
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

func (x gen_NSMenuItem) SetToolTip_(
	value core.NSStringRef,
) {
	C.NSMenuItem_inst_setToolTip_(
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

func (x gen_NSMenuItem) SetRepresentedObject_(
	value objc.Ref,
) {
	C.NSMenuItem_inst_setRepresentedObject_(
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

func (x gen_NSMenuItem) SetView_(
	value NSViewRef,
) {
	C.NSMenuItem_inst_setView_(
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

func (x gen_NSMenuItem) SetAllowsAutomaticKeyEquivalentLocalization_(
	value bool,
) {
	C.NSMenuItem_inst_setAllowsAutomaticKeyEquivalentLocalization_(
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

func (x gen_NSMenuItem) SetAllowsAutomaticKeyEquivalentMirroring_(
	value bool,
) {
	C.NSMenuItem_inst_setAllowsAutomaticKeyEquivalentMirroring_(
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

func (x gen_NSMenuItem) SetAllowsKeyEquivalentWhenHidden_(
	value bool,
) {
	C.NSMenuItem_inst_setAllowsKeyEquivalentWhenHidden_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
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

func (x gen_NSScreen) ConvertRectFromBacking_(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSScreen_inst_convertRectFromBacking_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSScreen) ConvertRectToBacking_(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSScreen_inst_convertRectToBacking_(
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

func (x gen_NSStatusBar) StatusItemWithLength_(
	length core.CGFloat,
) (
	r0 NSStatusItem,
) {
	ret := C.NSStatusBar_inst_statusItemWithLength_(
		unsafe.Pointer(x.Pointer()),
		C.double(length),
	)
	r0 = NSStatusItem_fromPointer(ret)
	return
}

func (x gen_NSStatusBar) RemoveStatusItem_(
	item NSStatusItemRef,
) {
	C.NSStatusBar_inst_removeStatusItem_(
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

func (x gen_NSStatusBarButton) SetAppearsDisabled_(
	value bool,
) {
	C.NSStatusBarButton_inst_setAppearsDisabled_(
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

func (x gen_NSStatusItem) SetMenu_(
	value NSMenuRef,
) {
	C.NSStatusItem_inst_setMenu_(
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

func (x gen_NSStatusItem) SetVisible_(
	value bool,
) {
	C.NSStatusItem_inst_setVisible_(
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

func (x gen_NSStatusItem) SetLength_(
	value core.CGFloat,
) {
	C.NSStatusItem_inst_setLength_(
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

func (x gen_NSText) InitWithFrame__asNSText(
	frameRect core.NSRect,
) (
	r0 NSText,
) {
	ret := C.NSText_inst_initWithFrame_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)
	r0 = NSText_fromPointer(ret)
	return
}

func (x gen_NSText) ToggleRuler_(
	sender objc.Ref,
) {
	C.NSText_inst_toggleRuler_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) SelectAll_(
	sender objc.Ref,
) {
	C.NSText_inst_selectAll_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Copy_(
	sender objc.Ref,
) {
	C.NSText_inst_copy_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Cut_(
	sender objc.Ref,
) {
	C.NSText_inst_cut_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Paste_(
	sender objc.Ref,
) {
	C.NSText_inst_paste_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) CopyFont_(
	sender objc.Ref,
) {
	C.NSText_inst_copyFont_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) PasteFont_(
	sender objc.Ref,
) {
	C.NSText_inst_pasteFont_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) CopyRuler_(
	sender objc.Ref,
) {
	C.NSText_inst_copyRuler_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) PasteRuler_(
	sender objc.Ref,
) {
	C.NSText_inst_pasteRuler_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Delete_(
	sender objc.Ref,
) {
	C.NSText_inst_delete_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) ChangeFont_(
	sender objc.Ref,
) {
	C.NSText_inst_changeFont_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) AlignCenter_(
	sender objc.Ref,
) {
	C.NSText_inst_alignCenter_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) AlignLeft_(
	sender objc.Ref,
) {
	C.NSText_inst_alignLeft_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) AlignRight_(
	sender objc.Ref,
) {
	C.NSText_inst_alignRight_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Superscript_(
	sender objc.Ref,
) {
	C.NSText_inst_superscript_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Subscript_(
	sender objc.Ref,
) {
	C.NSText_inst_subscript_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Unscript_(
	sender objc.Ref,
) {
	C.NSText_inst_unscript_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) Underline_(
	sender objc.Ref,
) {
	C.NSText_inst_underline_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) ReadRTFDFromFile_(
	path core.NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSText_inst_readRTFDFromFile_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSText) WriteRTFDToFile_atomically_(
	path core.NSStringRef,
	flag bool,
) (
	r0 bool,
) {
	ret := C.NSText_inst_writeRTFDToFile_atomically_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(flag),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSText) CheckSpelling_(
	sender objc.Ref,
) {
	C.NSText_inst_checkSpelling_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSText) ShowGuessPanel_(
	sender objc.Ref,
) {
	C.NSText_inst_showGuessPanel_(
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

func (x gen_NSText) SetString_(
	value core.NSStringRef,
) {
	C.NSText_inst_setString_(
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

func (x gen_NSText) SetBackgroundColor_(
	value NSColorRef,
) {
	C.NSText_inst_setBackgroundColor_(
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

func (x gen_NSText) SetDrawsBackground_(
	value bool,
) {
	C.NSText_inst_setDrawsBackground_(
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

func (x gen_NSText) SetEditable_(
	value bool,
) {
	C.NSText_inst_setEditable_(
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

func (x gen_NSText) SetSelectable_(
	value bool,
) {
	C.NSText_inst_setSelectable_(
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

func (x gen_NSText) SetFieldEditor_(
	value bool,
) {
	C.NSText_inst_setFieldEditor_(
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

func (x gen_NSText) SetRichText_(
	value bool,
) {
	C.NSText_inst_setRichText_(
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

func (x gen_NSText) SetImportsGraphics_(
	value bool,
) {
	C.NSText_inst_setImportsGraphics_(
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

func (x gen_NSText) SetUsesFontPanel_(
	value bool,
) {
	C.NSText_inst_setUsesFontPanel_(
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

func (x gen_NSText) SetFont_(
	value NSFontRef,
) {
	C.NSText_inst_setFont_(
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

func (x gen_NSText) SetTextColor_(
	value NSColorRef,
) {
	C.NSText_inst_setTextColor_(
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

func (x gen_NSText) SetMaxSize_(
	value core.NSSize,
) {
	C.NSText_inst_setMaxSize_(
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

func (x gen_NSText) SetMinSize_(
	value core.NSSize,
) {
	C.NSText_inst_setMinSize_(
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

func (x gen_NSText) SetVerticallyResizable_(
	value bool,
) {
	C.NSText_inst_setVerticallyResizable_(
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

func (x gen_NSText) SetHorizontallyResizable_(
	value bool,
) {
	C.NSText_inst_setHorizontallyResizable_(
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

func (x gen_NSText) SetDelegate_(
	value objc.Ref,
) {
	C.NSText_inst_setDelegate_(
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

func (x gen_NSTextContainer) InitWithSize__asNSTextContainer(
	size core.NSSize,
) (
	r0 NSTextContainer,
) {
	ret := C.NSTextContainer_inst_initWithSize_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	r0 = NSTextContainer_fromPointer(ret)
	return
}

func (x gen_NSTextContainer) ReplaceLayoutManager_(
	newLayoutManager NSLayoutManagerRef,
) {
	C.NSTextContainer_inst_replaceLayoutManager_(
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

func (x gen_NSTextContainer) SetLayoutManager_(
	value NSLayoutManagerRef,
) {
	C.NSTextContainer_inst_setLayoutManager_(
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

func (x gen_NSTextContainer) SetTextView_(
	value NSTextViewRef,
) {
	C.NSTextContainer_inst_setTextView_(
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

func (x gen_NSTextContainer) SetSize_(
	value core.NSSize,
) {
	C.NSTextContainer_inst_setSize_(
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

func (x gen_NSTextContainer) SetExclusionPaths_(
	value core.NSArrayRef,
) {
	C.NSTextContainer_inst_setExclusionPaths_(
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

func (x gen_NSTextContainer) SetWidthTracksTextView_(
	value bool,
) {
	C.NSTextContainer_inst_setWidthTracksTextView_(
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

func (x gen_NSTextContainer) SetHeightTracksTextView_(
	value bool,
) {
	C.NSTextContainer_inst_setHeightTracksTextView_(
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

func (x gen_NSTextContainer) SetMaximumNumberOfLines_(
	value core.NSUInteger,
) {
	C.NSTextContainer_inst_setMaximumNumberOfLines_(
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

func (x gen_NSTextContainer) SetLineFragmentPadding_(
	value core.CGFloat,
) {
	C.NSTextContainer_inst_setLineFragmentPadding_(
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

func (x gen_NSViewController) CommitEditingWithDelegate_didCommitSelector_contextInfo_(
	delegate objc.Ref,
	didCommitSelector objc.Selector,
	contextInfo unsafe.Pointer,
) {
	C.NSViewController_inst_commitEditingWithDelegate_didCommitSelector_contextInfo_(
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

func (x gen_NSViewController) DismissController_(
	sender objc.Ref,
) {
	C.NSViewController_inst_dismissController_(
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

func (x gen_NSViewController) AddChildViewController_(
	childViewController NSViewControllerRef,
) {
	C.NSViewController_inst_addChildViewController_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(childViewController),
	)
	return
}

func (x gen_NSViewController) InsertChildViewController_atIndex_(
	childViewController NSViewControllerRef,
	index core.NSInteger,
) {
	C.NSViewController_inst_insertChildViewController_atIndex_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(childViewController),
		C.long(index),
	)
	return
}

func (x gen_NSViewController) RemoveChildViewControllerAtIndex_(
	index core.NSInteger,
) {
	C.NSViewController_inst_removeChildViewControllerAtIndex_(
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

func (x gen_NSViewController) PreferredContentSizeDidChangeForViewController_(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_preferredContentSizeDidChangeForViewController_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)
	return
}

func (x gen_NSViewController) PresentViewController_animator_(
	viewController NSViewControllerRef,
	animator objc.Ref,
) {
	C.NSViewController_inst_presentViewController_animator_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
		objc.RefPointer(animator),
	)
	return
}

func (x gen_NSViewController) DismissViewController_(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_dismissViewController_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)
	return
}

func (x gen_NSViewController) PresentViewControllerAsModalWindow_(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_presentViewControllerAsModalWindow_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)
	return
}

func (x gen_NSViewController) PresentViewControllerAsSheet_(
	viewController NSViewControllerRef,
) {
	C.NSViewController_inst_presentViewControllerAsSheet_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(viewController),
	)
	return
}

func (x gen_NSViewController) ViewWillTransitionToSize_(
	newSize core.NSSize,
) {
	C.NSViewController_inst_viewWillTransitionToSize_(
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

func (x gen_NSViewController) SetRepresentedObject_(
	value objc.Ref,
) {
	C.NSViewController_inst_setRepresentedObject_(
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

func (x gen_NSViewController) Title() (
	r0 core.NSString,
) {
	ret := C.NSViewController_inst_title(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_NSViewController) SetTitle_(
	value core.NSStringRef,
) {
	C.NSViewController_inst_setTitle_(
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

func (x gen_NSViewController) SetPreferredContentSize_(
	value core.NSSize,
) {
	C.NSViewController_inst_setPreferredContentSize_(
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

func (x gen_NSViewController) SetChildViewControllers_(
	value core.NSArrayRef,
) {
	C.NSViewController_inst_setChildViewControllers_(
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

func (x gen_NSViewController) SetPreferredScreenOrigin_(
	value core.NSPoint,
) {
	C.NSViewController_inst_setPreferredScreenOrigin_(
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

func (x gen_NSVisualEffectView) ViewWillMoveToWindow_(
	newWindow NSWindowRef,
) {
	C.NSVisualEffectView_inst_viewWillMoveToWindow_(
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

func (x gen_NSVisualEffectView) SetEmphasized_(
	value bool,
) {
	C.NSVisualEffectView_inst_setEmphasized_(
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

func (x gen_NSVisualEffectView) SetMaskImage_(
	value NSImageRef,
) {
	C.NSVisualEffectView_inst_setMaskImage_(
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

func (x gen_NSWindow) InitWithContentRect_styleMask_backing_defer__asNSWindow(
	contentRect core.NSRect,
	style core.NSUInteger,
	backingStoreType core.NSUInteger,
	flag bool,
) (
	r0 NSWindow,
) {
	ret := C.NSWindow_inst_initWithContentRect_styleMask_backing_defer_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentRect)),
		C.ulong(style),
		C.ulong(backingStoreType),
		convertToObjCBool(flag),
	)
	r0 = NSWindow_fromPointer(ret)
	return
}

func (x gen_NSWindow) InitWithContentRect_styleMask_backing_defer_screen__asNSWindow(
	contentRect core.NSRect,
	style core.NSUInteger,
	backingStoreType core.NSUInteger,
	flag bool,
	screen NSScreenRef,
) (
	r0 NSWindow,
) {
	ret := C.NSWindow_inst_initWithContentRect_styleMask_backing_defer_screen_(
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

func (x gen_NSWindow) ToggleFullScreen_(
	sender objc.Ref,
) {
	C.NSWindow_inst_toggleFullScreen_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) SetDynamicDepthLimit_(
	flag bool,
) {
	C.NSWindow_inst_setDynamicDepthLimit_(
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

func (x gen_NSWindow) ContentRectForFrameRect_(
	frameRect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_contentRectForFrameRect_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) FrameRectForContentRect_(
	contentRect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_frameRectForContentRect_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&contentRect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) EndSheet_(
	sheetWindow NSWindowRef,
) {
	C.NSWindow_inst_endSheet_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sheetWindow),
	)
	return
}

func (x gen_NSWindow) SetFrameOrigin_(
	point core.NSPoint,
) {
	C.NSWindow_inst_setFrameOrigin_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	return
}

func (x gen_NSWindow) SetFrameTopLeftPoint_(
	point core.NSPoint,
) {
	C.NSWindow_inst_setFrameTopLeftPoint_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	return
}

func (x gen_NSWindow) ConstrainFrameRect_toScreen_(
	frameRect core.NSRect,
	screen NSScreenRef,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_constrainFrameRect_toScreen_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
		objc.RefPointer(screen),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) CascadeTopLeftFromPoint_(
	topLeftPoint core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSWindow_inst_cascadeTopLeftFromPoint_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&topLeftPoint)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) SetFrame_display_(
	frameRect core.NSRect,
	flag bool,
) {
	C.NSWindow_inst_setFrame_display_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
		convertToObjCBool(flag),
	)
	return
}

func (x gen_NSWindow) SetFrame_display_animate_(
	frameRect core.NSRect,
	displayFlag bool,
	animateFlag bool,
) {
	C.NSWindow_inst_setFrame_display_animate_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
		convertToObjCBool(displayFlag),
		convertToObjCBool(animateFlag),
	)
	return
}

func (x gen_NSWindow) PerformZoom_(
	sender objc.Ref,
) {
	C.NSWindow_inst_performZoom_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) Zoom_(
	sender objc.Ref,
) {
	C.NSWindow_inst_zoom_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) SetContentSize_(
	size core.NSSize,
) {
	C.NSWindow_inst_setContentSize_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	return
}

func (x gen_NSWindow) OrderOut_(
	sender objc.Ref,
) {
	C.NSWindow_inst_orderOut_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) OrderBack_(
	sender objc.Ref,
) {
	C.NSWindow_inst_orderBack_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) OrderFront_(
	sender objc.Ref,
) {
	C.NSWindow_inst_orderFront_(
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

func (x gen_NSWindow) OrderWindow_relativeTo_(
	place core.NSUInteger,
	otherWin core.NSInteger,
) {
	C.NSWindow_inst_orderWindow_relativeTo_(
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

func (x gen_NSWindow) MakeKeyAndOrderFront_(
	sender objc.Ref,
) {
	C.NSWindow_inst_makeKeyAndOrderFront_(
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

func (x gen_NSWindow) ToggleToolbarShown_(
	sender objc.Ref,
) {
	C.NSWindow_inst_toggleToolbarShown_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) RunToolbarCustomizationPalette_(
	sender objc.Ref,
) {
	C.NSWindow_inst_runToolbarCustomizationPalette_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) AddChildWindow_ordered_(
	childWin NSWindowRef,
	place core.NSUInteger,
) {
	C.NSWindow_inst_addChildWindow_ordered_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(childWin),
		C.ulong(place),
	)
	return
}

func (x gen_NSWindow) RemoveChildWindow_(
	childWin NSWindowRef,
) {
	C.NSWindow_inst_removeChildWindow_(
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

func (x gen_NSWindow) FieldEditor_forObject_(
	createFlag bool,
	object objc.Ref,
) (
	r0 NSText,
) {
	ret := C.NSWindow_inst_fieldEditor_forObject_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(createFlag),
		objc.RefPointer(object),
	)
	r0 = NSText_fromPointer(ret)
	return
}

func (x gen_NSWindow) EndEditingFor_(
	object objc.Ref,
) {
	C.NSWindow_inst_endEditingFor_(
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

func (x gen_NSWindow) InvalidateCursorRectsForView_(
	view NSViewRef,
) {
	C.NSWindow_inst_invalidateCursorRectsForView_(
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

func (x gen_NSWindow) RemoveTitlebarAccessoryViewControllerAtIndex_(
	index core.NSInteger,
) {
	C.NSWindow_inst_removeTitlebarAccessoryViewControllerAtIndex_(
		unsafe.Pointer(x.Pointer()),
		C.long(index),
	)
	return
}

func (x gen_NSWindow) AddTabbedWindow_ordered_(
	window NSWindowRef,
	ordered core.NSUInteger,
) {
	C.NSWindow_inst_addTabbedWindow_ordered_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(window),
		C.ulong(ordered),
	)
	return
}

func (x gen_NSWindow) MergeAllWindows_(
	sender objc.Ref,
) {
	C.NSWindow_inst_mergeAllWindows_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) SelectNextTab_(
	sender objc.Ref,
) {
	C.NSWindow_inst_selectNextTab_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) SelectPreviousTab_(
	sender objc.Ref,
) {
	C.NSWindow_inst_selectPreviousTab_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) MoveTabToNewWindow_(
	sender objc.Ref,
) {
	C.NSWindow_inst_moveTabToNewWindow_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) ToggleTabBar_(
	sender objc.Ref,
) {
	C.NSWindow_inst_toggleTabBar_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) ToggleTabOverview_(
	sender objc.Ref,
) {
	C.NSWindow_inst_toggleTabOverview_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) PostEvent_atStart_(
	event NSEventRef,
	flag bool,
) {
	C.NSWindow_inst_postEvent_atStart_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
		convertToObjCBool(flag),
	)
	return
}

func (x gen_NSWindow) SendEvent_(
	event NSEventRef,
) {
	C.NSWindow_inst_sendEvent_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	return
}

func (x gen_NSWindow) TryToPerform_with_(
	action objc.Selector,
	object objc.Ref,
) (
	r0 bool,
) {
	ret := C.NSWindow_inst_tryToPerform_with_(
		unsafe.Pointer(x.Pointer()),
		action.SelectorAddress(),
		objc.RefPointer(object),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSWindow) SelectKeyViewPrecedingView_(
	view NSViewRef,
) {
	C.NSWindow_inst_selectKeyViewPrecedingView_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)
	return
}

func (x gen_NSWindow) SelectKeyViewFollowingView_(
	view NSViewRef,
) {
	C.NSWindow_inst_selectKeyViewFollowingView_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)
	return
}

func (x gen_NSWindow) SelectPreviousKeyView_(
	sender objc.Ref,
) {
	C.NSWindow_inst_selectPreviousKeyView_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) SelectNextKeyView_(
	sender objc.Ref,
) {
	C.NSWindow_inst_selectNextKeyView_(
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

func (x gen_NSWindow) PerformWindowDragWithEvent_(
	event NSEventRef,
) {
	C.NSWindow_inst_performWindowDragWithEvent_(
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

func (x gen_NSWindow) DragImage_at_offset_event_pasteboard_source_slideBack_(
	image NSImageRef,
	baseLocation core.NSPoint,
	initialOffset core.NSSize,
	event NSEventRef,
	pboard NSPasteboardRef,
	sourceObj objc.Ref,
	slideFlag bool,
) {
	C.NSWindow_inst_dragImage_at_offset_event_pasteboard_source_slideBack_(
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

func (x gen_NSWindow) RegisterForDraggedTypes_(
	newTypes core.NSArrayRef,
) {
	C.NSWindow_inst_registerForDraggedTypes_(
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

func (x gen_NSWindow) ConvertRectFromBacking_(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_convertRectFromBacking_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) ConvertRectFromScreen_(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_convertRectFromScreen_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) ConvertPointFromBacking_(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSWindow_inst_convertPointFromBacking_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) ConvertPointFromScreen_(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSWindow_inst_convertPointFromScreen_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) ConvertRectToBacking_(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_convertRectToBacking_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) ConvertRectToScreen_(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSWindow_inst_convertRectToScreen_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) ConvertPointToBacking_(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSWindow_inst_convertPointToBacking_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) ConvertPointToScreen_(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSWindow_inst_convertPointToScreen_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSWindow) SetTitleWithRepresentedFilename_(
	filename core.NSStringRef,
) {
	C.NSWindow_inst_setTitleWithRepresentedFilename_(
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

func (x gen_NSWindow) PerformClose_(
	sender objc.Ref,
) {
	C.NSWindow_inst_performClose_(
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

func (x gen_NSWindow) PerformMiniaturize_(
	sender objc.Ref,
) {
	C.NSWindow_inst_performMiniaturize_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) Miniaturize_(
	sender objc.Ref,
) {
	C.NSWindow_inst_miniaturize_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) Deminiaturize_(
	sender objc.Ref,
) {
	C.NSWindow_inst_deminiaturize_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) Print_(
	sender objc.Ref,
) {
	C.NSWindow_inst_print_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSWindow) DataWithEPSInsideRect_(
	rect core.NSRect,
) (
	r0 core.NSData,
) {
	ret := C.NSWindow_inst_dataWithEPSInsideRect_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = core.NSData_fromPointer(ret)
	return
}

func (x gen_NSWindow) DataWithPDFInsideRect_(
	rect core.NSRect,
) (
	r0 core.NSData,
) {
	ret := C.NSWindow_inst_dataWithPDFInsideRect_(
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

func (x gen_NSWindow) VisualizeConstraints_(
	constraints core.NSArrayRef,
) {
	C.NSWindow_inst_visualizeConstraints_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(constraints),
	)
	return
}

func (x gen_NSWindow) SetIsMiniaturized_(
	flag bool,
) {
	C.NSWindow_inst_setIsMiniaturized_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)
	return
}

func (x gen_NSWindow) SetIsVisible_(
	flag bool,
) {
	C.NSWindow_inst_setIsVisible_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(flag),
	)
	return
}

func (x gen_NSWindow) SetIsZoomed_(
	flag bool,
) {
	C.NSWindow_inst_setIsZoomed_(
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

func (x gen_NSWindow) SetDelegate_(
	value objc.Ref,
) {
	C.NSWindow_inst_setDelegate_(
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

func (x gen_NSWindow) SetContentViewController_(
	value NSViewControllerRef,
) {
	C.NSWindow_inst_setContentViewController_(
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

func (x gen_NSWindow) SetContentView_(
	value NSViewRef,
) {
	C.NSWindow_inst_setContentView_(
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

func (x gen_NSWindow) SetStyleMask_(
	value core.NSUInteger,
) {
	C.NSWindow_inst_setStyleMask_(
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

func (x gen_NSWindow) SetAlphaValue_(
	value core.CGFloat,
) {
	C.NSWindow_inst_setAlphaValue_(
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

func (x gen_NSWindow) SetBackgroundColor_(
	value NSColorRef,
) {
	C.NSWindow_inst_setBackgroundColor_(
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

func (x gen_NSWindow) SetCanHide_(
	value bool,
) {
	C.NSWindow_inst_setCanHide_(
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

func (x gen_NSWindow) SetHidesOnDeactivate_(
	value bool,
) {
	C.NSWindow_inst_setHidesOnDeactivate_(
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

func (x gen_NSWindow) SetCollectionBehavior_(
	value core.NSUInteger,
) {
	C.NSWindow_inst_setCollectionBehavior_(
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

func (x gen_NSWindow) SetOpaque_(
	value bool,
) {
	C.NSWindow_inst_setOpaque_(
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

func (x gen_NSWindow) SetHasShadow_(
	value bool,
) {
	C.NSWindow_inst_setHasShadow_(
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

func (x gen_NSWindow) SetPreventsApplicationTerminationWhenModal_(
	value bool,
) {
	C.NSWindow_inst_setPreventsApplicationTerminationWhenModal_(
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

func (x gen_NSWindow) SetCanBecomeVisibleWithoutLogin_(
	value bool,
) {
	C.NSWindow_inst_setCanBecomeVisibleWithoutLogin_(
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

func (x gen_NSWindow) SetBackingType_(
	value core.NSUInteger,
) {
	C.NSWindow_inst_setBackingType_(
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

func (x gen_NSWindow) SetAspectRatio_(
	value core.NSSize,
) {
	C.NSWindow_inst_setAspectRatio_(
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

func (x gen_NSWindow) SetMinSize_(
	value core.NSSize,
) {
	C.NSWindow_inst_setMinSize_(
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

func (x gen_NSWindow) SetMaxSize_(
	value core.NSSize,
) {
	C.NSWindow_inst_setMaxSize_(
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

func (x gen_NSWindow) SetResizeIncrements_(
	value core.NSSize,
) {
	C.NSWindow_inst_setResizeIncrements_(
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

func (x gen_NSWindow) SetPreservesContentDuringLiveResize_(
	value bool,
) {
	C.NSWindow_inst_setPreservesContentDuringLiveResize_(
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

func (x gen_NSWindow) SetContentAspectRatio_(
	value core.NSSize,
) {
	C.NSWindow_inst_setContentAspectRatio_(
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

func (x gen_NSWindow) SetContentMinSize_(
	value core.NSSize,
) {
	C.NSWindow_inst_setContentMinSize_(
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

func (x gen_NSWindow) SetContentMaxSize_(
	value core.NSSize,
) {
	C.NSWindow_inst_setContentMaxSize_(
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

func (x gen_NSWindow) SetContentResizeIncrements_(
	value core.NSSize,
) {
	C.NSWindow_inst_setContentResizeIncrements_(
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

func (x gen_NSWindow) SetMaxFullScreenContentSize_(
	value core.NSSize,
) {
	C.NSWindow_inst_setMaxFullScreenContentSize_(
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

func (x gen_NSWindow) SetMinFullScreenContentSize_(
	value core.NSSize,
) {
	C.NSWindow_inst_setMinFullScreenContentSize_(
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

func (x gen_NSWindow) SetLevel_(
	value core.NSInteger,
) {
	C.NSWindow_inst_setLevel_(
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

func (x gen_NSWindow) SetParentWindow_(
	value NSWindowRef,
) {
	C.NSWindow_inst_setParentWindow_(
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

func (x gen_NSWindow) SetExcludedFromWindowsMenu_(
	value bool,
) {
	C.NSWindow_inst_setExcludedFromWindowsMenu_(
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

func (x gen_NSWindow) SetShowsToolbarButton_(
	value bool,
) {
	C.NSWindow_inst_setShowsToolbarButton_(
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

func (x gen_NSWindow) SetTitlebarAppearsTransparent_(
	value bool,
) {
	C.NSWindow_inst_setTitlebarAppearsTransparent_(
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

func (x gen_NSWindow) SetTitlebarAccessoryViewControllers_(
	value core.NSArrayRef,
) {
	C.NSWindow_inst_setTitlebarAccessoryViewControllers_(
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

func (x gen_NSWindow) SetAllowsToolTipsWhenApplicationIsInactive_(
	value bool,
) {
	C.NSWindow_inst_setAllowsToolTipsWhenApplicationIsInactive_(
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

func (x gen_NSWindow) SetInitialFirstResponder_(
	value NSViewRef,
) {
	C.NSWindow_inst_setInitialFirstResponder_(
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

func (x gen_NSWindow) SetAutorecalculatesKeyViewLoop_(
	value bool,
) {
	C.NSWindow_inst_setAutorecalculatesKeyViewLoop_(
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

func (x gen_NSWindow) SetAcceptsMouseMovedEvents_(
	value bool,
) {
	C.NSWindow_inst_setAcceptsMouseMovedEvents_(
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

func (x gen_NSWindow) SetIgnoresMouseEvents_(
	value bool,
) {
	C.NSWindow_inst_setIgnoresMouseEvents_(
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

func (x gen_NSWindow) SetRestorable_(
	value bool,
) {
	C.NSWindow_inst_setRestorable_(
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

func (x gen_NSWindow) SetViewsNeedDisplay_(
	value bool,
) {
	C.NSWindow_inst_setViewsNeedDisplay_(
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

func (x gen_NSWindow) SetAllowsConcurrentViewDrawing_(
	value bool,
) {
	C.NSWindow_inst_setAllowsConcurrentViewDrawing_(
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

func (x gen_NSWindow) SetDocumentEdited_(
	value bool,
) {
	C.NSWindow_inst_setDocumentEdited_(
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

func (x gen_NSWindow) SetTitle_(
	value core.NSStringRef,
) {
	C.NSWindow_inst_setTitle_(
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

func (x gen_NSWindow) SetSubtitle_(
	value core.NSStringRef,
) {
	C.NSWindow_inst_setSubtitle_(
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

func (x gen_NSWindow) SetTitleVisibility_(
	value core.NSInteger,
) {
	C.NSWindow_inst_setTitleVisibility_(
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

func (x gen_NSWindow) SetRepresentedFilename_(
	value core.NSStringRef,
) {
	C.NSWindow_inst_setRepresentedFilename_(
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

func (x gen_NSWindow) SetRepresentedURL_(
	value core.NSURLRef,
) {
	C.NSWindow_inst_setRepresentedURL_(
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

func (x gen_NSWindow) SetDisplaysWhenScreenProfileChanges_(
	value bool,
) {
	C.NSWindow_inst_setDisplaysWhenScreenProfileChanges_(
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

func (x gen_NSWindow) SetMovableByWindowBackground_(
	value bool,
) {
	C.NSWindow_inst_setMovableByWindowBackground_(
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

func (x gen_NSWindow) SetMovable_(
	value bool,
) {
	C.NSWindow_inst_setMovable_(
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

func (x gen_NSWindow) SetReleasedWhenClosed_(
	value bool,
) {
	C.NSWindow_inst_setReleasedWhenClosed_(
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

func (x gen_NSWindow) SetMiniwindowImage_(
	value NSImageRef,
) {
	C.NSWindow_inst_setMiniwindowImage_(
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

func (x gen_NSWindow) SetMiniwindowTitle_(
	value core.NSStringRef,
) {
	C.NSWindow_inst_setMiniwindowTitle_(
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

func (x gen_NSWindow) SetOrderedIndex_(
	value core.NSInteger,
) {
	C.NSWindow_inst_setOrderedIndex_(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)
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

func (x gen_NSColor) BlendedColorWithFraction_ofColor_(
	fraction core.CGFloat,
	color NSColorRef,
) (
	r0 NSColor,
) {
	ret := C.NSColor_inst_blendedColorWithFraction_ofColor_(
		unsafe.Pointer(x.Pointer()),
		C.double(fraction),
		objc.RefPointer(color),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSColor) ColorWithAlphaComponent_(
	alpha core.CGFloat,
) (
	r0 NSColor,
) {
	ret := C.NSColor_inst_colorWithAlphaComponent_(
		unsafe.Pointer(x.Pointer()),
		C.double(alpha),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSColor) HighlightWithLevel_(
	val core.CGFloat,
) (
	r0 NSColor,
) {
	ret := C.NSColor_inst_highlightWithLevel_(
		unsafe.Pointer(x.Pointer()),
		C.double(val),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSColor) ShadowWithLevel_(
	val core.CGFloat,
) (
	r0 NSColor,
) {
	ret := C.NSColor_inst_shadowWithLevel_(
		unsafe.Pointer(x.Pointer()),
		C.double(val),
	)
	r0 = NSColor_fromPointer(ret)
	return
}

func (x gen_NSColor) WriteToPasteboard_(
	pasteBoard NSPasteboardRef,
) {
	C.NSColor_inst_writeToPasteboard_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pasteBoard),
	)
	return
}

func (x gen_NSColor) DrawSwatchInRect_(
	rect core.NSRect,
) {
	C.NSColor_inst_drawSwatchInRect_(
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

func (x gen_NSTextView) InitWithFrame_textContainer__asNSTextView(
	frameRect core.NSRect,
	container NSTextContainerRef,
) (
	r0 NSTextView,
) {
	ret := C.NSTextView_inst_initWithFrame_textContainer_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
		objc.RefPointer(container),
	)
	r0 = NSTextView_fromPointer(ret)
	return
}

func (x gen_NSTextView) InitWithFrame__asNSTextView(
	frameRect core.NSRect,
) (
	r0 NSTextView,
) {
	ret := C.NSTextView_inst_initWithFrame_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frameRect)),
	)
	r0 = NSTextView_fromPointer(ret)
	return
}

func (x gen_NSTextView) ReplaceTextContainer_(
	newContainer NSTextContainerRef,
) {
	C.NSTextView_inst_replaceTextContainer_(
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

func (x gen_NSTextView) ChangeDocumentBackgroundColor_(
	sender objc.Ref,
) {
	C.NSTextView_inst_changeDocumentBackgroundColor_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) SetNeedsDisplayInRect_avoidAdditionalLayout_(
	rect core.NSRect,
	flag bool,
) {
	C.NSTextView_inst_setNeedsDisplayInRect_avoidAdditionalLayout_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		convertToObjCBool(flag),
	)
	return
}

func (x gen_NSTextView) DrawInsertionPointInRect_color_turnedOn_(
	rect core.NSRect,
	color NSColorRef,
	flag bool,
) {
	C.NSTextView_inst_drawInsertionPointInRect_color_turnedOn_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(color),
		convertToObjCBool(flag),
	)
	return
}

func (x gen_NSTextView) DrawViewBackgroundInRect_(
	rect core.NSRect,
) {
	C.NSTextView_inst_drawViewBackgroundInRect_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	return
}

func (x gen_NSTextView) SetConstrainedFrameSize_(
	desiredSize core.NSSize,
) {
	C.NSTextView_inst_setConstrainedFrameSize_(
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

func (x gen_NSTextView) Outline_(
	sender objc.Ref,
) {
	C.NSTextView_inst_outline_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleAutomaticQuoteSubstitution_(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticQuoteSubstitution_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleAutomaticLinkDetection_(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticLinkDetection_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleAutomaticTextCompletion_(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticTextCompletion_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) UpdateInsertionPointStateAndRestartTimer_(
	restartFlag bool,
) {
	C.NSTextView_inst_updateInsertionPointStateAndRestartTimer_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(restartFlag),
	)
	return
}

func (x gen_NSTextView) CharacterIndexForInsertionAtPoint_(
	point core.NSPoint,
) (
	r0 core.NSUInteger,
) {
	ret := C.NSTextView_inst_characterIndexForInsertionAtPoint_(
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

func (x gen_NSTextView) ReadSelectionFromPasteboard_(
	pboard NSPasteboardRef,
) (
	r0 bool,
) {
	ret := C.NSTextView_inst_readSelectionFromPasteboard_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pboard),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) WriteSelectionToPasteboard_types_(
	pboard NSPasteboardRef,
	types core.NSArrayRef,
) (
	r0 bool,
) {
	ret := C.NSTextView_inst_writeSelectionToPasteboard_types_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pboard),
		objc.RefPointer(types),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) AlignJustified_(
	sender objc.Ref,
) {
	C.NSTextView_inst_alignJustified_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ChangeAttributes_(
	sender objc.Ref,
) {
	C.NSTextView_inst_changeAttributes_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ChangeColor_(
	sender objc.Ref,
) {
	C.NSTextView_inst_changeColor_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) UseStandardKerning_(
	sender objc.Ref,
) {
	C.NSTextView_inst_useStandardKerning_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) LowerBaseline_(
	sender objc.Ref,
) {
	C.NSTextView_inst_lowerBaseline_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) RaiseBaseline_(
	sender objc.Ref,
) {
	C.NSTextView_inst_raiseBaseline_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) TurnOffKerning_(
	sender objc.Ref,
) {
	C.NSTextView_inst_turnOffKerning_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) LoosenKerning_(
	sender objc.Ref,
) {
	C.NSTextView_inst_loosenKerning_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) TightenKerning_(
	sender objc.Ref,
) {
	C.NSTextView_inst_tightenKerning_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) UseStandardLigatures_(
	sender objc.Ref,
) {
	C.NSTextView_inst_useStandardLigatures_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) TurnOffLigatures_(
	sender objc.Ref,
) {
	C.NSTextView_inst_turnOffLigatures_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) UseAllLigatures_(
	sender objc.Ref,
) {
	C.NSTextView_inst_useAllLigatures_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ClickedOnLink_atIndex_(
	link objc.Ref,
	charIndex core.NSUInteger,
) {
	C.NSTextView_inst_clickedOnLink_atIndex_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(link),
		C.ulong(charIndex),
	)
	return
}

func (x gen_NSTextView) PasteAsPlainText_(
	sender objc.Ref,
) {
	C.NSTextView_inst_pasteAsPlainText_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) PasteAsRichText_(
	sender objc.Ref,
) {
	C.NSTextView_inst_pasteAsRichText_(
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

func (x gen_NSTextView) ShouldChangeTextInRanges_replacementStrings_(
	affectedRanges core.NSArrayRef,
	replacementStrings core.NSArrayRef,
) (
	r0 bool,
) {
	ret := C.NSTextView_inst_shouldChangeTextInRanges_replacementStrings_(
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

func (x gen_NSTextView) ToggleSmartInsertDelete_(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleSmartInsertDelete_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleContinuousSpellChecking_(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleContinuousSpellChecking_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleGrammarChecking_(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleGrammarChecking_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) OrderFrontSharingServicePicker_(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontSharingServicePicker_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) DragSelectionWithEvent_offset_slideBack_(
	event NSEventRef,
	mouseOffset core.NSSize,
	slideBack bool,
) (
	r0 bool,
) {
	ret := C.NSTextView_inst_dragSelectionWithEvent_offset_slideBack_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
		*(*C.NSSize)(unsafe.Pointer(&mouseOffset)),
		convertToObjCBool(slideBack),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSTextView) StartSpeaking_(
	sender objc.Ref,
) {
	C.NSTextView_inst_startSpeaking_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) StopSpeaking_(
	sender objc.Ref,
) {
	C.NSTextView_inst_stopSpeaking_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) PerformFindPanelAction_(
	sender objc.Ref,
) {
	C.NSTextView_inst_performFindPanelAction_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) OrderFrontLinkPanel_(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontLinkPanel_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) OrderFrontListPanel_(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontListPanel_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) OrderFrontSpacingPanel_(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontSpacingPanel_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) OrderFrontTablePanel_(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontTablePanel_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) OrderFrontSubstitutionsPanel_(
	sender objc.Ref,
) {
	C.NSTextView_inst_orderFrontSubstitutionsPanel_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) Complete_(
	sender objc.Ref,
) {
	C.NSTextView_inst_complete_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) CheckTextInDocument_(
	sender objc.Ref,
) {
	C.NSTextView_inst_checkTextInDocument_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) CheckTextInSelection_(
	sender objc.Ref,
) {
	C.NSTextView_inst_checkTextInSelection_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleAutomaticDashSubstitution_(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticDashSubstitution_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleAutomaticDataDetection_(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticDataDetection_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleAutomaticSpellingCorrection_(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticSpellingCorrection_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) ToggleAutomaticTextReplacement_(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleAutomaticTextReplacement_(
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

func (x gen_NSTextView) ToggleQuickLookPreviewPanel_(
	sender objc.Ref,
) {
	C.NSTextView_inst_toggleQuickLookPreviewPanel_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSTextView) QuickLookPreviewableItemsInRanges_(
	ranges core.NSArrayRef,
) (
	r0 core.NSArray,
) {
	ret := C.NSTextView_inst_quickLookPreviewableItemsInRanges_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(ranges),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_NSTextView) ChangeLayoutOrientation_(
	sender objc.Ref,
) {
	C.NSTextView_inst_changeLayoutOrientation_(
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

func (x gen_NSTextView) SetDelegate_(
	value objc.Ref,
) {
	C.NSTextView_inst_setDelegate_(
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

func (x gen_NSTextView) SetTextContainer_(
	value NSTextContainerRef,
) {
	C.NSTextView_inst_setTextContainer_(
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

func (x gen_NSTextView) SetTextContainerInset_(
	value core.NSSize,
) {
	C.NSTextView_inst_setTextContainerInset_(
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

func (x gen_NSTextView) SetBackgroundColor_(
	value NSColorRef,
) {
	C.NSTextView_inst_setBackgroundColor_(
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

func (x gen_NSTextView) SetDrawsBackground_(
	value bool,
) {
	C.NSTextView_inst_setDrawsBackground_(
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

func (x gen_NSTextView) SetAllowsDocumentBackgroundColorChange_(
	value bool,
) {
	C.NSTextView_inst_setAllowsDocumentBackgroundColorChange_(
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

func (x gen_NSTextView) SetAllowedInputSourceLocales_(
	value core.NSArrayRef,
) {
	C.NSTextView_inst_setAllowedInputSourceLocales_(
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

func (x gen_NSTextView) SetAllowsUndo_(
	value bool,
) {
	C.NSTextView_inst_setAllowsUndo_(
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

func (x gen_NSTextView) SetEditable_(
	value bool,
) {
	C.NSTextView_inst_setEditable_(
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

func (x gen_NSTextView) SetSelectable_(
	value bool,
) {
	C.NSTextView_inst_setSelectable_(
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

func (x gen_NSTextView) SetFieldEditor_(
	value bool,
) {
	C.NSTextView_inst_setFieldEditor_(
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

func (x gen_NSTextView) SetRichText_(
	value bool,
) {
	C.NSTextView_inst_setRichText_(
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

func (x gen_NSTextView) SetImportsGraphics_(
	value bool,
) {
	C.NSTextView_inst_setImportsGraphics_(
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

func (x gen_NSTextView) SetAllowsImageEditing_(
	value bool,
) {
	C.NSTextView_inst_setAllowsImageEditing_(
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

func (x gen_NSTextView) SetAutomaticQuoteSubstitutionEnabled_(
	value bool,
) {
	C.NSTextView_inst_setAutomaticQuoteSubstitutionEnabled_(
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

func (x gen_NSTextView) SetAutomaticLinkDetectionEnabled_(
	value bool,
) {
	C.NSTextView_inst_setAutomaticLinkDetectionEnabled_(
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

func (x gen_NSTextView) SetDisplaysLinkToolTips_(
	value bool,
) {
	C.NSTextView_inst_setDisplaysLinkToolTips_(
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

func (x gen_NSTextView) SetAutomaticTextCompletionEnabled_(
	value bool,
) {
	C.NSTextView_inst_setAutomaticTextCompletionEnabled_(
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

func (x gen_NSTextView) SetUsesAdaptiveColorMappingForDarkAppearance_(
	value bool,
) {
	C.NSTextView_inst_setUsesAdaptiveColorMappingForDarkAppearance_(
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

func (x gen_NSTextView) SetUsesRolloverButtonForSelection_(
	value bool,
) {
	C.NSTextView_inst_setUsesRolloverButtonForSelection_(
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

func (x gen_NSTextView) SetUsesRuler_(
	value bool,
) {
	C.NSTextView_inst_setUsesRuler_(
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

func (x gen_NSTextView) SetRulerVisible_(
	value bool,
) {
	C.NSTextView_inst_setRulerVisible_(
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

func (x gen_NSTextView) SetUsesInspectorBar_(
	value bool,
) {
	C.NSTextView_inst_setUsesInspectorBar_(
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

func (x gen_NSTextView) SetSelectedRanges_(
	value core.NSArrayRef,
) {
	C.NSTextView_inst_setSelectedRanges_(
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

func (x gen_NSTextView) SetInsertionPointColor_(
	value NSColorRef,
) {
	C.NSTextView_inst_setInsertionPointColor_(
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

func (x gen_NSTextView) SetSelectedTextAttributes_(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_setSelectedTextAttributes_(
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

func (x gen_NSTextView) SetMarkedTextAttributes_(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_setMarkedTextAttributes_(
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

func (x gen_NSTextView) SetLinkTextAttributes_(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_setLinkTextAttributes_(
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

func (x gen_NSTextView) SetTypingAttributes_(
	value core.NSDictionaryRef,
) {
	C.NSTextView_inst_setTypingAttributes_(
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

func (x gen_NSTextView) SetSmartInsertDeleteEnabled_(
	value bool,
) {
	C.NSTextView_inst_setSmartInsertDeleteEnabled_(
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

func (x gen_NSTextView) SetContinuousSpellCheckingEnabled_(
	value bool,
) {
	C.NSTextView_inst_setContinuousSpellCheckingEnabled_(
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

func (x gen_NSTextView) SetGrammarCheckingEnabled_(
	value bool,
) {
	C.NSTextView_inst_setGrammarCheckingEnabled_(
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

func (x gen_NSTextView) SetAcceptsGlyphInfo_(
	value bool,
) {
	C.NSTextView_inst_setAcceptsGlyphInfo_(
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

func (x gen_NSTextView) SetUsesFontPanel_(
	value bool,
) {
	C.NSTextView_inst_setUsesFontPanel_(
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

func (x gen_NSTextView) SetUsesFindPanel_(
	value bool,
) {
	C.NSTextView_inst_setUsesFindPanel_(
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

func (x gen_NSTextView) SetAutomaticDashSubstitutionEnabled_(
	value bool,
) {
	C.NSTextView_inst_setAutomaticDashSubstitutionEnabled_(
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

func (x gen_NSTextView) SetAutomaticDataDetectionEnabled_(
	value bool,
) {
	C.NSTextView_inst_setAutomaticDataDetectionEnabled_(
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

func (x gen_NSTextView) SetAutomaticSpellingCorrectionEnabled_(
	value bool,
) {
	C.NSTextView_inst_setAutomaticSpellingCorrectionEnabled_(
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

func (x gen_NSTextView) SetAutomaticTextReplacementEnabled_(
	value bool,
) {
	C.NSTextView_inst_setAutomaticTextReplacementEnabled_(
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

func (x gen_NSTextView) SetUsesFindBar_(
	value bool,
) {
	C.NSTextView_inst_setUsesFindBar_(
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

func (x gen_NSTextView) SetIncrementalSearchingEnabled_(
	value bool,
) {
	C.NSTextView_inst_setIncrementalSearchingEnabled_(
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

func (x gen_NSTextView) SetAllowsCharacterPickerTouchBarItem_(
	value bool,
) {
	C.NSTextView_inst_setAllowsCharacterPickerTouchBarItem_(
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

func (x gen_NSTextView) SetFont_(
	value NSFontRef,
) {
	C.NSTextView_inst_setFont_(
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

func (x gen_NSView) InitWithFrame__asNSView(
	frameRect core.NSRect,
) (
	r0 NSView,
) {
	ret := C.NSView_inst_initWithFrame_(
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

func (x gen_NSView) AddSubview_(
	view NSViewRef,
) {
	C.NSView_inst_addSubview_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)
	return
}

func (x gen_NSView) AddSubview_positioned_relativeTo_(
	view NSViewRef,
	place core.NSUInteger,
	otherView NSViewRef,
) {
	C.NSView_inst_addSubview_positioned_relativeTo_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
		C.ulong(place),
		objc.RefPointer(otherView),
	)
	return
}

func (x gen_NSView) DidAddSubview_(
	subview NSViewRef,
) {
	C.NSView_inst_didAddSubview_(
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

func (x gen_NSView) ReplaceSubview_with_(
	oldView NSViewRef,
	newView NSViewRef,
) {
	C.NSView_inst_replaceSubview_with_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(oldView),
		objc.RefPointer(newView),
	)
	return
}

func (x gen_NSView) IsDescendantOf_(
	view NSViewRef,
) (
	r0 bool,
) {
	ret := C.NSView_inst_isDescendantOf_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(view),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) AncestorSharedWithView_(
	view NSViewRef,
) (
	r0 NSView,
) {
	ret := C.NSView_inst_ancestorSharedWithView_(
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

func (x gen_NSView) ViewWillMoveToSuperview_(
	newSuperview NSViewRef,
) {
	C.NSView_inst_viewWillMoveToSuperview_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newSuperview),
	)
	return
}

func (x gen_NSView) ViewWillMoveToWindow_(
	newWindow NSWindowRef,
) {
	C.NSView_inst_viewWillMoveToWindow_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(newWindow),
	)
	return
}

func (x gen_NSView) WillRemoveSubview_(
	subview NSViewRef,
) {
	C.NSView_inst_willRemoveSubview_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(subview),
	)
	return
}

func (x gen_NSView) SetFrameOrigin_(
	newOrigin core.NSPoint,
) {
	C.NSView_inst_setFrameOrigin_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&newOrigin)),
	)
	return
}

func (x gen_NSView) SetFrameSize_(
	newSize core.NSSize,
) {
	C.NSView_inst_setFrameSize_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&newSize)),
	)
	return
}

func (x gen_NSView) SetBoundsOrigin_(
	newOrigin core.NSPoint,
) {
	C.NSView_inst_setBoundsOrigin_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&newOrigin)),
	)
	return
}

func (x gen_NSView) SetBoundsSize_(
	newSize core.NSSize,
) {
	C.NSView_inst_setBoundsSize_(
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

func (x gen_NSView) DrawRect_(
	dirtyRect core.NSRect,
) {
	C.NSView_inst_drawRect_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&dirtyRect)),
	)
	return
}

func (x gen_NSView) NeedsToDrawRect_(
	rect core.NSRect,
) (
	r0 bool,
) {
	ret := C.NSView_inst_needsToDrawRect_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) Print_(
	sender objc.Ref,
) {
	C.NSView_inst_print_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_NSView) BeginPageInRect_atPlacement_(
	rect core.NSRect,
	location core.NSPoint,
) {
	C.NSView_inst_beginPageInRect_atPlacement_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		*(*C.NSPoint)(unsafe.Pointer(&location)),
	)
	return
}

func (x gen_NSView) DataWithEPSInsideRect_(
	rect core.NSRect,
) (
	r0 core.NSData,
) {
	ret := C.NSView_inst_dataWithEPSInsideRect_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = core.NSData_fromPointer(ret)
	return
}

func (x gen_NSView) DataWithPDFInsideRect_(
	rect core.NSRect,
) (
	r0 core.NSData,
) {
	ret := C.NSView_inst_dataWithPDFInsideRect_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = core.NSData_fromPointer(ret)
	return
}

func (x gen_NSView) WriteEPSInsideRect_toPasteboard_(
	rect core.NSRect,
	pasteboard NSPasteboardRef,
) {
	C.NSView_inst_writeEPSInsideRect_toPasteboard_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(pasteboard),
	)
	return
}

func (x gen_NSView) WritePDFInsideRect_toPasteboard_(
	rect core.NSRect,
	pasteboard NSPasteboardRef,
) {
	C.NSView_inst_writePDFInsideRect_toPasteboard_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(pasteboard),
	)
	return
}

func (x gen_NSView) DrawPageBorderWithSize_(
	borderSize core.NSSize,
) {
	C.NSView_inst_drawPageBorderWithSize_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&borderSize)),
	)
	return
}

func (x gen_NSView) RectForPage_(
	page core.NSInteger,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_rectForPage_(
		unsafe.Pointer(x.Pointer()),
		C.long(page),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) LocationOfPrintRect_(
	rect core.NSRect,
) (
	r0 core.NSPoint,
) {
	ret := C.NSView_inst_locationOfPrintRect_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) SetNeedsDisplayInRect_(
	invalidRect core.NSRect,
) {
	C.NSView_inst_setNeedsDisplayInRect_(
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

func (x gen_NSView) DisplayRect_(
	rect core.NSRect,
) {
	C.NSView_inst_displayRect_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	return
}

func (x gen_NSView) DisplayRectIgnoringOpacity_(
	rect core.NSRect,
) {
	C.NSView_inst_displayRectIgnoringOpacity_(
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

func (x gen_NSView) DisplayIfNeededInRect_(
	rect core.NSRect,
) {
	C.NSView_inst_displayIfNeededInRect_(
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

func (x gen_NSView) DisplayIfNeededInRectIgnoringOpacity_(
	rect core.NSRect,
) {
	C.NSView_inst_displayIfNeededInRectIgnoringOpacity_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	return
}

func (x gen_NSView) TranslateRectsNeedingDisplayInRect_by_(
	clipRect core.NSRect,
	delta core.NSSize,
) {
	C.NSView_inst_translateRectsNeedingDisplayInRect_by_(
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

func (x gen_NSView) ConvertPointFromBacking_(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSView_inst_convertPointFromBacking_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertPointToBacking_(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSView_inst_convertPointToBacking_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertPointFromLayer_(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSView_inst_convertPointFromLayer_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertPointToLayer_(
	point core.NSPoint,
) (
	r0 core.NSPoint,
) {
	ret := C.NSView_inst_convertPointToLayer_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertRectFromBacking_(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_convertRectFromBacking_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertRectToBacking_(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_convertRectToBacking_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertRectFromLayer_(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_convertRectFromLayer_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertRectToLayer_(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_convertRectToLayer_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertSizeFromBacking_(
	size core.NSSize,
) (
	r0 core.NSSize,
) {
	ret := C.NSView_inst_convertSizeFromBacking_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertSizeToBacking_(
	size core.NSSize,
) (
	r0 core.NSSize,
) {
	ret := C.NSView_inst_convertSizeToBacking_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertSizeFromLayer_(
	size core.NSSize,
) (
	r0 core.NSSize,
) {
	ret := C.NSView_inst_convertSizeFromLayer_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertSizeToLayer_(
	size core.NSSize,
) (
	r0 core.NSSize,
) {
	ret := C.NSView_inst_convertSizeToLayer_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertPoint_fromView_(
	point core.NSPoint,
	view NSViewRef,
) (
	r0 core.NSPoint,
) {
	ret := C.NSView_inst_convertPoint_fromView_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
		objc.RefPointer(view),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertPoint_toView_(
	point core.NSPoint,
	view NSViewRef,
) (
	r0 core.NSPoint,
) {
	ret := C.NSView_inst_convertPoint_toView_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
		objc.RefPointer(view),
	)
	r0 = *(*core.NSPoint)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertSize_fromView_(
	size core.NSSize,
	view NSViewRef,
) (
	r0 core.NSSize,
) {
	ret := C.NSView_inst_convertSize_fromView_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
		objc.RefPointer(view),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertSize_toView_(
	size core.NSSize,
	view NSViewRef,
) (
	r0 core.NSSize,
) {
	ret := C.NSView_inst_convertSize_toView_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
		objc.RefPointer(view),
	)
	r0 = *(*core.NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertRect_fromView_(
	rect core.NSRect,
	view NSViewRef,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_convertRect_fromView_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(view),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) ConvertRect_toView_(
	rect core.NSRect,
	view NSViewRef,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_convertRect_toView_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(view),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) CenterScanRect_(
	rect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_centerScanRect_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) TranslateOriginToPoint_(
	translation core.NSPoint,
) {
	C.NSView_inst_translateOriginToPoint_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&translation)),
	)
	return
}

func (x gen_NSView) ScaleUnitSquareToSize_(
	newUnitSize core.NSSize,
) {
	C.NSView_inst_scaleUnitSquareToSize_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&newUnitSize)),
	)
	return
}

func (x gen_NSView) RotateByAngle_(
	angle core.CGFloat,
) {
	C.NSView_inst_rotateByAngle_(
		unsafe.Pointer(x.Pointer()),
		C.double(angle),
	)
	return
}

func (x gen_NSView) ResizeSubviewsWithOldSize_(
	oldSize core.NSSize,
) {
	C.NSView_inst_resizeSubviewsWithOldSize_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&oldSize)),
	)
	return
}

func (x gen_NSView) ResizeWithOldSuperviewSize_(
	oldSize core.NSSize,
) {
	C.NSView_inst_resizeWithOldSuperviewSize_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&oldSize)),
	)
	return
}

func (x gen_NSView) AddConstraints_(
	constraints core.NSArrayRef,
) {
	C.NSView_inst_addConstraints_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(constraints),
	)
	return
}

func (x gen_NSView) RemoveConstraints_(
	constraints core.NSArrayRef,
) {
	C.NSView_inst_removeConstraints_(
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

func (x gen_NSView) AlignmentRectForFrame_(
	frame core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_alignmentRectForFrame_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frame)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) FrameForAlignmentRect_(
	alignmentRect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_frameForAlignmentRect_(
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

func (x gen_NSView) SetKeyboardFocusRingNeedsDisplayInRect_(
	rect core.NSRect,
) {
	C.NSView_inst_setKeyboardFocusRingNeedsDisplayInRect_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	return
}

func (x gen_NSView) EnterFullScreenMode_withOptions_(
	screen NSScreenRef,
	options core.NSDictionaryRef,
) (
	r0 bool,
) {
	ret := C.NSView_inst_enterFullScreenMode_withOptions_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(screen),
		objc.RefPointer(options),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) ExitFullScreenModeWithOptions_(
	options core.NSDictionaryRef,
) {
	C.NSView_inst_exitFullScreenModeWithOptions_(
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

func (x gen_NSView) AcceptsFirstMouse_(
	event NSEventRef,
) (
	r0 bool,
) {
	ret := C.NSView_inst_acceptsFirstMouse_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) HitTest_(
	point core.NSPoint,
) (
	r0 NSView,
) {
	ret := C.NSView_inst_hitTest_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	r0 = NSView_fromPointer(ret)
	return
}

func (x gen_NSView) Mouse_inRect_(
	point core.NSPoint,
	rect core.NSRect,
) (
	r0 bool,
) {
	ret := C.NSView_inst_mouse_inRect_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) PerformKeyEquivalent_(
	event NSEventRef,
) (
	r0 bool,
) {
	ret := C.NSView_inst_performKeyEquivalent_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) PrepareContentInRect_(
	rect core.NSRect,
) {
	C.NSView_inst_prepareContentInRect_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	return
}

func (x gen_NSView) ScrollPoint_(
	point core.NSPoint,
) {
	C.NSView_inst_scrollPoint_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSPoint)(unsafe.Pointer(&point)),
	)
	return
}

func (x gen_NSView) ScrollRectToVisible_(
	rect core.NSRect,
) (
	r0 bool,
) {
	ret := C.NSView_inst_scrollRectToVisible_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) Autoscroll_(
	event NSEventRef,
) (
	r0 bool,
) {
	ret := C.NSView_inst_autoscroll_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) AdjustScroll_(
	newVisible core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_adjustScroll_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&newVisible)),
	)
	r0 = *(*core.NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSView) RegisterForDraggedTypes_(
	newTypes core.NSArrayRef,
) {
	C.NSView_inst_registerForDraggedTypes_(
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

func (x gen_NSView) ShouldDelayWindowOrderingForEvent_(
	event NSEventRef,
) (
	r0 bool,
) {
	ret := C.NSView_inst_shouldDelayWindowOrderingForEvent_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSView) RectForSmartMagnificationAtPoint_inRect_(
	location core.NSPoint,
	visibleRect core.NSRect,
) (
	r0 core.NSRect,
) {
	ret := C.NSView_inst_rectForSmartMagnificationAtPoint_inRect_(
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

func (x gen_NSView) ViewWithTag_(
	tag core.NSInteger,
) (
	r0 NSView,
) {
	ret := C.NSView_inst_viewWithTag_(
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

func (x gen_NSView) MenuForEvent_(
	event NSEventRef,
) (
	r0 NSMenu,
) {
	ret := C.NSView_inst_menuForEvent_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	r0 = NSMenu_fromPointer(ret)
	return
}

func (x gen_NSView) WillOpenMenu_withEvent_(
	menu NSMenuRef,
	event NSEventRef,
) {
	C.NSView_inst_willOpenMenu_withEvent_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(menu),
		objc.RefPointer(event),
	)
	return
}

func (x gen_NSView) DidCloseMenu_withEvent_(
	menu NSMenuRef,
	event NSEventRef,
) {
	C.NSView_inst_didCloseMenu_withEvent_(
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

func (x gen_NSView) ShowDefinitionForAttributedString_atPoint_(
	attrString core.NSAttributedStringRef,
	textBaselineOrigin core.NSPoint,
) {
	C.NSView_inst_showDefinitionForAttributedString_atPoint_(
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

func (x gen_NSView) SetSubviews_(
	value core.NSArrayRef,
) {
	C.NSView_inst_setSubviews_(
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

func (x gen_NSView) SetFrame_(
	value core.NSRect,
) {
	C.NSView_inst_setFrame_(
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

func (x gen_NSView) SetFrameRotation_(
	value core.CGFloat,
) {
	C.NSView_inst_setFrameRotation_(
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

func (x gen_NSView) SetBounds_(
	value core.NSRect,
) {
	C.NSView_inst_setBounds_(
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

func (x gen_NSView) SetBoundsRotation_(
	value core.CGFloat,
) {
	C.NSView_inst_setBoundsRotation_(
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

func (x gen_NSView) SetWantsLayer_(
	value bool,
) {
	C.NSView_inst_setWantsLayer_(
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

func (x gen_NSView) SetLayer_(
	value core.CALayerRef,
) {
	C.NSView_inst_setLayer_(
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

func (x gen_NSView) SetCanDrawSubviewsIntoLayer_(
	value bool,
) {
	C.NSView_inst_setCanDrawSubviewsIntoLayer_(
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

func (x gen_NSView) SetLayerUsesCoreImageFilters_(
	value bool,
) {
	C.NSView_inst_setLayerUsesCoreImageFilters_(
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

func (x gen_NSView) SetAlphaValue_(
	value core.CGFloat,
) {
	C.NSView_inst_setAlphaValue_(
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

func (x gen_NSView) SetFrameCenterRotation_(
	value core.CGFloat,
) {
	C.NSView_inst_setFrameCenterRotation_(
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

func (x gen_NSView) SetBackgroundFilters_(
	value core.NSArrayRef,
) {
	C.NSView_inst_setBackgroundFilters_(
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

func (x gen_NSView) SetContentFilters_(
	value core.NSArrayRef,
) {
	C.NSView_inst_setContentFilters_(
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

func (x gen_NSView) SetCanDrawConcurrently_(
	value bool,
) {
	C.NSView_inst_setCanDrawConcurrently_(
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

func (x gen_NSView) SetNeedsDisplay_(
	value bool,
) {
	C.NSView_inst_setNeedsDisplay_(
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

func (x gen_NSView) SetAutoresizesSubviews_(
	value bool,
) {
	C.NSView_inst_setAutoresizesSubviews_(
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

func (x gen_NSView) SetNeedsLayout_(
	value bool,
) {
	C.NSView_inst_setNeedsLayout_(
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

func (x gen_NSView) SetNeedsUpdateConstraints_(
	value bool,
) {
	C.NSView_inst_setNeedsUpdateConstraints_(
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

func (x gen_NSView) SetTranslatesAutoresizingMaskIntoConstraints_(
	value bool,
) {
	C.NSView_inst_setTranslatesAutoresizingMaskIntoConstraints_(
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

func (x gen_NSView) SetHidden_(
	value bool,
) {
	C.NSView_inst_setHidden_(
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

func (x gen_NSView) SetGestureRecognizers_(
	value core.NSArrayRef,
) {
	C.NSView_inst_setGestureRecognizers_(
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

func (x gen_NSView) SetWantsRestingTouches_(
	value bool,
) {
	C.NSView_inst_setWantsRestingTouches_(
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

func (x gen_NSView) SetNextKeyView_(
	value NSViewRef,
) {
	C.NSView_inst_setNextKeyView_(
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

func (x gen_NSView) SetPreparedContentRect_(
	value core.NSRect,
) {
	C.NSView_inst_setPreparedContentRect_(
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

func (x gen_NSView) SetPostsFrameChangedNotifications_(
	value bool,
) {
	C.NSView_inst_setPostsFrameChangedNotifications_(
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

func (x gen_NSView) SetPostsBoundsChangedNotifications_(
	value bool,
) {
	C.NSView_inst_setPostsBoundsChangedNotifications_(
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

func (x gen_NSView) SetToolTip_(
	value core.NSStringRef,
) {
	C.NSView_inst_setToolTip_(
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

func (x gen_NSView) SetHorizontalContentSizeConstraintActive_(
	value bool,
) {
	C.NSView_inst_setHorizontalContentSizeConstraintActive_(
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

func (x gen_NSView) SetVerticalContentSizeConstraintActive_(
	value bool,
) {
	C.NSView_inst_setVerticalContentSizeConstraintActive_(
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

func (x gen_NSView) SetBackgroundColor_(
	value NSColorRef,
) {
	C.NSView_inst_setBackgroundColor_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}
