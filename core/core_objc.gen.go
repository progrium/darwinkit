package core

import (
	"github.com/progrium/macdriver/objc"
	"unsafe"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -lobjc -framework AppKit -framework QuartzCore -framework Foundation
#define __OBJC2__ 1
#include <objc/message.h>
#include <stdlib.h>

#include <AppKit/AppKit.h>
#include <QuartzCore/QuartzCore.h>
#include <Foundation/Foundation.h>

bool core_convertObjCBool(BOOL b) {
	if (b) { return true; }
	return false;
}


void* NSObject_type_Alloc() {
	return [NSObject
		alloc];
}
void NSObject_type_Initialize() {
	[NSObject
		initialize];
}
void NSObject_type_Load() {
	[NSObject
		load];
}
void* NSObject_type_New() {
	return [NSObject
		new];
}
BOOL NSObject_type_InstancesRespondToSelector(void* aSelector) {
	return [NSObject
		instancesRespondToSelector: aSelector];
}
void* NSObject_type_Description() {
	return [NSObject
		description];
}
void NSObject_type_CancelPreviousPerformRequestsWithTarget(void* aTarget) {
	[NSObject
		cancelPreviousPerformRequestsWithTarget: aTarget];
}
void NSObject_type_CancelPreviousPerformRequestsWithTargetSelectorObject(void* aTarget, void* aSelector, void* anArgument) {
	[NSObject
		cancelPreviousPerformRequestsWithTarget: aTarget
		selector: aSelector
		object: anArgument];
}
BOOL NSObject_type_ResolveClassMethod(void* sel) {
	return [NSObject
		resolveClassMethod: sel];
}
BOOL NSObject_type_ResolveInstanceMethod(void* sel) {
	return [NSObject
		resolveInstanceMethod: sel];
}
void* NSObject_type_ClassFallbacksForKeyedArchiver() {
	return [NSObject
		classFallbacksForKeyedArchiver];
}
void NSObject_type_SetVersion(long aVersion) {
	[NSObject
		setVersion: aVersion];
}
long NSObject_type_Version() {
	return [NSObject
		version];
}
void* NSObject_type_DebugDescription() {
	return [NSObject
		debugDescription];
}
unsigned long NSObject_type_Hash() {
	return [NSObject
		hash];
}
void* CALayer_type_Alloc() {
	return [CALayer
		alloc];
}
void* CALayer_type_Layer() {
	return [CALayer
		layer];
}
BOOL CALayer_type_NeedsDisplayForKey(void* key) {
	return [CALayer
		needsDisplayForKey: key];
}
void* CALayer_type_DefaultActionForKey(void* event) {
	return [CALayer
		defaultActionForKey: event];
}
void* CALayer_type_DefaultValueForKey(void* key) {
	return [CALayer
		defaultValueForKey: key];
}
void* NSArray_type_Alloc() {
	return [NSArray
		alloc];
}
void* NSArray_type_Array() {
	return [NSArray
		array];
}
void* NSArray_type_ArrayWithArray(void* array) {
	return [NSArray
		arrayWithArray: array];
}
void* NSAttributedString_type_Alloc() {
	return [NSAttributedString
		alloc];
}
void* NSAttributedString_type_TextTypes() {
	return [NSAttributedString
		textTypes];
}
void* NSAttributedString_type_TextUnfilteredTypes() {
	return [NSAttributedString
		textUnfilteredTypes];
}
void* NSData_type_Alloc() {
	return [NSData
		alloc];
}
void* NSData_type_Data() {
	return [NSData
		data];
}
void* NSData_type_DataWithBytesLength(void* bytes, unsigned long length) {
	return [NSData
		dataWithBytes: bytes
		length: length];
}
void* NSData_type_DataWithBytesNoCopyLength(void* bytes, unsigned long length) {
	return [NSData
		dataWithBytesNoCopy: bytes
		length: length];
}
void* NSData_type_DataWithBytesNoCopyLengthFreeWhenDone(void* bytes, unsigned long length, BOOL b) {
	return [NSData
		dataWithBytesNoCopy: bytes
		length: length
		freeWhenDone: b];
}
void* NSData_type_DataWithData(void* data) {
	return [NSData
		dataWithData: data];
}
void* NSData_type_DataWithContentsOfFile(void* path) {
	return [NSData
		dataWithContentsOfFile: path];
}
void* NSData_type_DataWithContentsOfURL(void* url) {
	return [NSData
		dataWithContentsOfURL: url];
}
void* NSDictionary_type_Alloc() {
	return [NSDictionary
		alloc];
}
void* NSDictionary_type_Dictionary() {
	return [NSDictionary
		dictionary];
}
void* NSDictionary_type_DictionaryWithObjectsForKeys(void* objects, void* keys) {
	return [NSDictionary
		dictionaryWithObjects: objects
		forKeys: keys];
}
void* NSDictionary_type_DictionaryWithDictionary(void* dict) {
	return [NSDictionary
		dictionaryWithDictionary: dict];
}
void* NSDictionary_type_SharedKeySetForKeys(void* keys) {
	return [NSDictionary
		sharedKeySetForKeys: keys];
}
void* NSNumber_type_Alloc() {
	return [NSNumber
		alloc];
}
void* NSNumber_type_NumberWithBool(BOOL value) {
	return [NSNumber
		numberWithBool: value];
}
void* NSNumber_type_NumberWithInt(int value) {
	return [NSNumber
		numberWithInt: value];
}
void* NSNumber_type_NumberWithInteger(long value) {
	return [NSNumber
		numberWithInteger: value];
}
void* NSNumber_type_NumberWithUnsignedInt(int value) {
	return [NSNumber
		numberWithUnsignedInt: value];
}
void* NSNumber_type_NumberWithUnsignedInteger(unsigned long value) {
	return [NSNumber
		numberWithUnsignedInteger: value];
}
void* NSRunLoop_type_Alloc() {
	return [NSRunLoop
		alloc];
}
void* NSRunLoop_type_CurrentRunLoop() {
	return [NSRunLoop
		currentRunLoop];
}
void* NSRunLoop_type_MainRunLoop() {
	return [NSRunLoop
		mainRunLoop];
}
void* NSString_type_Alloc() {
	return [NSString
		alloc];
}
void* NSString_type_String() {
	return [NSString
		string];
}
void* NSString_type_LocalizedUserNotificationStringForKeyArguments(void* key, void* arguments) {
	return [NSString
		localizedUserNotificationStringForKey: key
		arguments: arguments];
}
void* NSString_type_StringWithString(void* string) {
	return [NSString
		stringWithString: string];
}
void* NSString_type_LocalizedNameOfStringEncoding(unsigned long encoding) {
	return [NSString
		localizedNameOfStringEncoding: encoding];
}
void* NSString_type_PathWithComponents(void* components) {
	return [NSString
		pathWithComponents: components];
}
unsigned long NSString_type_DefaultCStringEncoding() {
	return [NSString
		defaultCStringEncoding];
}
void* NSThread_type_Alloc() {
	return [NSThread
		alloc];
}
void NSThread_type_DetachNewThreadSelectorToTargetWithObject(void* selector, void* target, void* argument) {
	[NSThread
		detachNewThreadSelector: selector
		toTarget: target
		withObject: argument];
}
void NSThread_type_Exit() {
	[NSThread
		exit];
}
BOOL NSThread_type_IsMultiThreaded() {
	return [NSThread
		isMultiThreaded];
}
BOOL NSThread_type_IsMainThread() {
	return [NSThread
		isMainThread];
}
void* NSThread_type_MainThread() {
	return [NSThread
		mainThread];
}
void* NSThread_type_CurrentThread() {
	return [NSThread
		currentThread];
}
void* NSThread_type_CallStackReturnAddresses() {
	return [NSThread
		callStackReturnAddresses];
}
void* NSThread_type_CallStackSymbols() {
	return [NSThread
		callStackSymbols];
}
void* NSURL_type_Alloc() {
	return [NSURL
		alloc];
}
void* NSURL_type_URLWithString(void* URLString) {
	return [NSURL
		URLWithString: URLString];
}
void* NSURL_type_URLWithStringRelativeToURL(void* URLString, void* baseURL) {
	return [NSURL
		URLWithString: URLString
		relativeToURL: baseURL];
}
void* NSURL_type_FileURLWithPathIsDirectory(void* path, BOOL isDir) {
	return [NSURL
		fileURLWithPath: path
		isDirectory: isDir];
}
void* NSURL_type_FileURLWithPathRelativeToURL(void* path, void* baseURL) {
	return [NSURL
		fileURLWithPath: path
		relativeToURL: baseURL];
}
void* NSURL_type_FileURLWithPathIsDirectoryRelativeToURL(void* path, BOOL isDir, void* baseURL) {
	return [NSURL
		fileURLWithPath: path
		isDirectory: isDir
		relativeToURL: baseURL];
}
void* NSURL_type_FileURLWithPath(void* path) {
	return [NSURL
		fileURLWithPath: path];
}
void* NSURL_type_FileURLWithPathComponents(void* components) {
	return [NSURL
		fileURLWithPathComponents: components];
}
void* NSURL_type_AbsoluteURLWithDataRepresentationRelativeToURL(void* data, void* baseURL) {
	return [NSURL
		absoluteURLWithDataRepresentation: data
		relativeToURL: baseURL];
}
void* NSURL_type_URLWithDataRepresentationRelativeToURL(void* data, void* baseURL) {
	return [NSURL
		URLWithDataRepresentation: data
		relativeToURL: baseURL];
}
void* NSURL_type_ResourceValuesForKeysFromBookmarkData(void* keys, void* bookmarkData) {
	return [NSURL
		resourceValuesForKeys: keys
		fromBookmarkData: bookmarkData];
}
void* NSURLRequest_type_Alloc() {
	return [NSURLRequest
		alloc];
}
void* NSURLRequest_type_RequestWithURL(void* URL) {
	return [NSURLRequest
		requestWithURL: URL];
}
BOOL NSURLRequest_type_SupportsSecureCoding() {
	return [NSURLRequest
		supportsSecureCoding];
}
void* NSUserDefaults_type_Alloc() {
	return [NSUserDefaults
		alloc];
}
void NSUserDefaults_type_ResetStandardUserDefaults() {
	[NSUserDefaults
		resetStandardUserDefaults];
}
void* NSUserDefaults_type_StandardUserDefaults() {
	return [NSUserDefaults
		standardUserDefaults];
}


void* NSObject_inst_ActionProperty(void *id) {
	return [(NSObject*)id
		actionProperty];
}

void* NSObject_inst_Candidates(void *id, void* sender) {
	return [(NSObject*)id
		candidates: sender];
}

void NSObject_inst_CommitComposition(void *id, void* sender) {
	[(NSObject*)id
		commitComposition: sender];
}

void* NSObject_inst_ComposedString(void *id, void* sender) {
	return [(NSObject*)id
		composedString: sender];
}

void* NSObject_inst_Copy(void *id) {
	return [(NSObject*)id
		copy];
}

void* NSObject_inst_CopyScriptingValueForKeyWithProperties(void *id, void* value, void* key, void* properties) {
	return [(NSObject*)id
		copyScriptingValue: value
		forKey: key
		withProperties: properties];
}

void NSObject_inst_Dealloc(void *id) {
	[(NSObject*)id
		dealloc];
}

BOOL NSObject_inst_DidCommandBySelectorClient(void *id, void* aSelector, void* sender) {
	return [(NSObject*)id
		didCommandBySelector: aSelector
		client: sender];
}

BOOL NSObject_inst_DoesContain(void *id, void* object) {
	return [(NSObject*)id
		doesContain: object];
}

void NSObject_inst_DoesNotRecognizeSelector(void *id, void* aSelector) {
	[(NSObject*)id
		doesNotRecognizeSelector: aSelector];
}

void* NSObject_inst_ForwardingTargetForSelector(void *id, void* aSelector) {
	return [(NSObject*)id
		forwardingTargetForSelector: aSelector];
}

void* NSObject_inst_ImageRepresentation(void *id) {
	return [(NSObject*)id
		imageRepresentation];
}

void* NSObject_inst_ImageRepresentationType(void *id) {
	return [(NSObject*)id
		imageRepresentationType];
}

void* NSObject_inst_ImageSubtitle(void *id) {
	return [(NSObject*)id
		imageSubtitle];
}

void* NSObject_inst_ImageTitle(void *id) {
	return [(NSObject*)id
		imageTitle];
}

void* NSObject_inst_ImageUID(void *id) {
	return [(NSObject*)id
		imageUID];
}

unsigned long NSObject_inst_ImageVersion(void *id) {
	return [(NSObject*)id
		imageVersion];
}

void* NSObject_inst_Init(void *id) {
	return [(NSObject*)id
		init];
}

BOOL NSObject_inst_InputTextClient(void *id, void* string, void* sender) {
	return [(NSObject*)id
		inputText: string
		client: sender];
}

BOOL NSObject_inst_InputTextKeyModifiersClient(void *id, void* string, long keyCode, unsigned long flags, void* sender) {
	return [(NSObject*)id
		inputText: string
		key: keyCode
		modifiers: flags
		client: sender];
}

void* NSObject_inst_InverseForRelationshipKey(void *id, void* relationshipKey) {
	return [(NSObject*)id
		inverseForRelationshipKey: relationshipKey];
}

BOOL NSObject_inst_IsCaseInsensitiveLike(void *id, void* object) {
	return [(NSObject*)id
		isCaseInsensitiveLike: object];
}

BOOL NSObject_inst_IsEqualTo(void *id, void* object) {
	return [(NSObject*)id
		isEqualTo: object];
}

BOOL NSObject_inst_IsGreaterThan(void *id, void* object) {
	return [(NSObject*)id
		isGreaterThan: object];
}

BOOL NSObject_inst_IsGreaterThanOrEqualTo(void *id, void* object) {
	return [(NSObject*)id
		isGreaterThanOrEqualTo: object];
}

BOOL NSObject_inst_IsLessThan(void *id, void* object) {
	return [(NSObject*)id
		isLessThan: object];
}

BOOL NSObject_inst_IsLessThanOrEqualTo(void *id, void* object) {
	return [(NSObject*)id
		isLessThanOrEqualTo: object];
}

BOOL NSObject_inst_IsLike(void *id, void* object) {
	return [(NSObject*)id
		isLike: object];
}

BOOL NSObject_inst_IsNotEqualTo(void *id, void* object) {
	return [(NSObject*)id
		isNotEqualTo: object];
}

void* NSObject_inst_MutableCopy(void *id) {
	return [(NSObject*)id
		mutableCopy];
}

void* NSObject_inst_OriginalString(void *id, void* sender) {
	return [(NSObject*)id
		originalString: sender];
}

void NSObject_inst_PerformSelectorOnThreadWithObjectWaitUntilDone(void *id, void* aSelector, void* thr, void* arg, BOOL wait) {
	[(NSObject*)id
		performSelector: aSelector
		onThread: thr
		withObject: arg
		waitUntilDone: wait];
}

void NSObject_inst_PerformSelectorOnThreadWithObjectWaitUntilDoneModes(void *id, void* aSelector, void* thr, void* arg, BOOL wait, void* array) {
	[(NSObject*)id
		performSelector: aSelector
		onThread: thr
		withObject: arg
		waitUntilDone: wait
		modes: array];
}

void NSObject_inst_PerformSelectorInBackgroundWithObject(void *id, void* aSelector, void* arg) {
	[(NSObject*)id
		performSelectorInBackground: aSelector
		withObject: arg];
}

void NSObject_inst_PerformSelectorOnMainThreadWithObjectWaitUntilDone(void *id, void* aSelector, void* arg, BOOL wait) {
	[(NSObject*)id
		performSelectorOnMainThread: aSelector
		withObject: arg
		waitUntilDone: wait];
}

void NSObject_inst_PerformSelectorOnMainThreadWithObjectWaitUntilDoneModes(void *id, void* aSelector, void* arg, BOOL wait, void* array) {
	[(NSObject*)id
		performSelectorOnMainThread: aSelector
		withObject: arg
		waitUntilDone: wait
		modes: array];
}

void* NSObject_inst_AutoContentAccessingProxy(void *id) {
	return [(NSObject*)id
		autoContentAccessingProxy];
}

void* NSObject_inst_AttributeKeys(void *id) {
	return [(NSObject*)id
		attributeKeys];
}

void* NSObject_inst_ToManyRelationshipKeys(void *id) {
	return [(NSObject*)id
		toManyRelationshipKeys];
}

void* NSObject_inst_ToOneRelationshipKeys(void *id) {
	return [(NSObject*)id
		toOneRelationshipKeys];
}

void* NSObject_inst_ClassName(void *id) {
	return [(NSObject*)id
		className];
}

void* NSObject_inst_ScriptingProperties(void *id) {
	return [(NSObject*)id
		scriptingProperties];
}

void NSObject_inst_SetScriptingProperties(void *id, void* value) {
	[(NSObject*)id
		setScriptingProperties: value];
}

BOOL NSObject_inst_AccessibilityNotifiesWhenDestroyed(void *id) {
	return [(NSObject*)id
		accessibilityNotifiesWhenDestroyed];
}

BOOL NSObject_inst_IsSelectable(void *id) {
	return [(NSObject*)id
		isSelectable];
}

void* CALayer_inst_ActionForKey(void *id, void* event) {
	return [(CALayer*)id
		actionForKey: event];
}

void CALayer_inst_AddSublayer(void *id, void* layer) {
	[(CALayer*)id
		addSublayer: layer];
}

void* CALayer_inst_AnimationKeys(void *id) {
	return [(CALayer*)id
		animationKeys];
}

BOOL CALayer_inst_ContentsAreFlipped(void *id) {
	return [(CALayer*)id
		contentsAreFlipped];
}

NSRect CALayer_inst_ConvertRectFromLayer(void *id, NSRect r, void* l) {
	return [(CALayer*)id
		convertRect: r
		fromLayer: l];
}

NSRect CALayer_inst_ConvertRectToLayer(void *id, NSRect r, void* l) {
	return [(CALayer*)id
		convertRect: r
		toLayer: l];
}

void CALayer_inst_Display(void *id) {
	[(CALayer*)id
		display];
}

void CALayer_inst_DisplayIfNeeded(void *id) {
	[(CALayer*)id
		displayIfNeeded];
}

void* CALayer_inst_Init(void *id) {
	return [(CALayer*)id
		init];
}

void* CALayer_inst_InitWithLayer(void *id, void* layer) {
	return [(CALayer*)id
		initWithLayer: layer];
}

void CALayer_inst_InsertSublayerAbove(void *id, void* layer, void* sibling) {
	[(CALayer*)id
		insertSublayer: layer
		above: sibling];
}

void CALayer_inst_InsertSublayerAtIndex(void *id, void* layer, int idx) {
	[(CALayer*)id
		insertSublayer: layer
		atIndex: idx];
}

void CALayer_inst_InsertSublayerBelow(void *id, void* layer, void* sibling) {
	[(CALayer*)id
		insertSublayer: layer
		below: sibling];
}

void CALayer_inst_LayoutIfNeeded(void *id) {
	[(CALayer*)id
		layoutIfNeeded];
}

void CALayer_inst_LayoutSublayers(void *id) {
	[(CALayer*)id
		layoutSublayers];
}

void* CALayer_inst_ModelLayer(void *id) {
	return [(CALayer*)id
		modelLayer];
}

BOOL CALayer_inst_NeedsDisplay(void *id) {
	return [(CALayer*)id
		needsDisplay];
}

BOOL CALayer_inst_NeedsLayout(void *id) {
	return [(CALayer*)id
		needsLayout];
}

NSSize CALayer_inst_PreferredFrameSize(void *id) {
	return [(CALayer*)id
		preferredFrameSize];
}

void* CALayer_inst_PresentationLayer(void *id) {
	return [(CALayer*)id
		presentationLayer];
}

void CALayer_inst_RemoveAllAnimations(void *id) {
	[(CALayer*)id
		removeAllAnimations];
}

void CALayer_inst_RemoveAnimationForKey(void *id, void* key) {
	[(CALayer*)id
		removeAnimationForKey: key];
}

void CALayer_inst_RemoveFromSuperlayer(void *id) {
	[(CALayer*)id
		removeFromSuperlayer];
}

void CALayer_inst_ReplaceSublayerWith(void *id, void* oldLayer, void* newLayer) {
	[(CALayer*)id
		replaceSublayer: oldLayer
		with: newLayer];
}

void CALayer_inst_ResizeSublayersWithOldSize(void *id, NSSize size) {
	[(CALayer*)id
		resizeSublayersWithOldSize: size];
}

void CALayer_inst_ResizeWithOldSuperlayerSize(void *id, NSSize size) {
	[(CALayer*)id
		resizeWithOldSuperlayerSize: size];
}

void CALayer_inst_ScrollRectToVisible(void *id, NSRect r) {
	[(CALayer*)id
		scrollRectToVisible: r];
}

void CALayer_inst_SetNeedsDisplay(void *id) {
	[(CALayer*)id
		setNeedsDisplay];
}

void CALayer_inst_SetNeedsDisplayInRect(void *id, NSRect r) {
	[(CALayer*)id
		setNeedsDisplayInRect: r];
}

void CALayer_inst_SetNeedsLayout(void *id) {
	[(CALayer*)id
		setNeedsLayout];
}

BOOL CALayer_inst_ShouldArchiveValueForKey(void *id, void* key) {
	return [(CALayer*)id
		shouldArchiveValueForKey: key];
}

void* CALayer_inst_Delegate(void *id) {
	return [(CALayer*)id
		delegate];
}

void CALayer_inst_SetDelegate(void *id, void* value) {
	[(CALayer*)id
		setDelegate: value];
}

void* CALayer_inst_Contents(void *id) {
	return [(CALayer*)id
		contents];
}

void CALayer_inst_SetContents(void *id, void* value) {
	[(CALayer*)id
		setContents: value];
}

NSRect CALayer_inst_ContentsRect(void *id) {
	return [(CALayer*)id
		contentsRect];
}

void CALayer_inst_SetContentsRect(void *id, NSRect value) {
	[(CALayer*)id
		setContentsRect: value];
}

NSRect CALayer_inst_ContentsCenter(void *id) {
	return [(CALayer*)id
		contentsCenter];
}

void CALayer_inst_SetContentsCenter(void *id, NSRect value) {
	[(CALayer*)id
		setContentsCenter: value];
}

BOOL CALayer_inst_IsHidden(void *id) {
	return [(CALayer*)id
		isHidden];
}

void CALayer_inst_SetHidden(void *id, BOOL value) {
	[(CALayer*)id
		setHidden: value];
}

BOOL CALayer_inst_MasksToBounds(void *id) {
	return [(CALayer*)id
		masksToBounds];
}

void CALayer_inst_SetMasksToBounds(void *id, BOOL value) {
	[(CALayer*)id
		setMasksToBounds: value];
}

void* CALayer_inst_Mask(void *id) {
	return [(CALayer*)id
		mask];
}

void CALayer_inst_SetMask(void *id, void* value) {
	[(CALayer*)id
		setMask: value];
}

BOOL CALayer_inst_IsDoubleSided(void *id) {
	return [(CALayer*)id
		isDoubleSided];
}

void CALayer_inst_SetDoubleSided(void *id, BOOL value) {
	[(CALayer*)id
		setDoubleSided: value];
}

double CALayer_inst_CornerRadius(void *id) {
	return [(CALayer*)id
		cornerRadius];
}

void CALayer_inst_SetCornerRadius(void *id, double value) {
	[(CALayer*)id
		setCornerRadius: value];
}

double CALayer_inst_BorderWidth(void *id) {
	return [(CALayer*)id
		borderWidth];
}

void CALayer_inst_SetBorderWidth(void *id, double value) {
	[(CALayer*)id
		setBorderWidth: value];
}

double CALayer_inst_ShadowRadius(void *id) {
	return [(CALayer*)id
		shadowRadius];
}

void CALayer_inst_SetShadowRadius(void *id, double value) {
	[(CALayer*)id
		setShadowRadius: value];
}

NSSize CALayer_inst_ShadowOffset(void *id) {
	return [(CALayer*)id
		shadowOffset];
}

void CALayer_inst_SetShadowOffset(void *id, NSSize value) {
	[(CALayer*)id
		setShadowOffset: value];
}

void* CALayer_inst_Style(void *id) {
	return [(CALayer*)id
		style];
}

void CALayer_inst_SetStyle(void *id, void* value) {
	[(CALayer*)id
		setStyle: value];
}

BOOL CALayer_inst_AllowsEdgeAntialiasing(void *id) {
	return [(CALayer*)id
		allowsEdgeAntialiasing];
}

void CALayer_inst_SetAllowsEdgeAntialiasing(void *id, BOOL value) {
	[(CALayer*)id
		setAllowsEdgeAntialiasing: value];
}

BOOL CALayer_inst_AllowsGroupOpacity(void *id) {
	return [(CALayer*)id
		allowsGroupOpacity];
}

void CALayer_inst_SetAllowsGroupOpacity(void *id, BOOL value) {
	[(CALayer*)id
		setAllowsGroupOpacity: value];
}

void* CALayer_inst_Filters(void *id) {
	return [(CALayer*)id
		filters];
}

void CALayer_inst_SetFilters(void *id, void* value) {
	[(CALayer*)id
		setFilters: value];
}

void* CALayer_inst_CompositingFilter(void *id) {
	return [(CALayer*)id
		compositingFilter];
}

void CALayer_inst_SetCompositingFilter(void *id, void* value) {
	[(CALayer*)id
		setCompositingFilter: value];
}

void* CALayer_inst_BackgroundFilters(void *id) {
	return [(CALayer*)id
		backgroundFilters];
}

void CALayer_inst_SetBackgroundFilters(void *id, void* value) {
	[(CALayer*)id
		setBackgroundFilters: value];
}

BOOL CALayer_inst_IsOpaque(void *id) {
	return [(CALayer*)id
		isOpaque];
}

void CALayer_inst_SetOpaque(void *id, BOOL value) {
	[(CALayer*)id
		setOpaque: value];
}

BOOL CALayer_inst_IsGeometryFlipped(void *id) {
	return [(CALayer*)id
		isGeometryFlipped];
}

void CALayer_inst_SetGeometryFlipped(void *id, BOOL value) {
	[(CALayer*)id
		setGeometryFlipped: value];
}

BOOL CALayer_inst_DrawsAsynchronously(void *id) {
	return [(CALayer*)id
		drawsAsynchronously];
}

void CALayer_inst_SetDrawsAsynchronously(void *id, BOOL value) {
	[(CALayer*)id
		setDrawsAsynchronously: value];
}

BOOL CALayer_inst_ShouldRasterize(void *id) {
	return [(CALayer*)id
		shouldRasterize];
}

void CALayer_inst_SetShouldRasterize(void *id, BOOL value) {
	[(CALayer*)id
		setShouldRasterize: value];
}

double CALayer_inst_RasterizationScale(void *id) {
	return [(CALayer*)id
		rasterizationScale];
}

void CALayer_inst_SetRasterizationScale(void *id, double value) {
	[(CALayer*)id
		setRasterizationScale: value];
}

NSRect CALayer_inst_Frame(void *id) {
	return [(CALayer*)id
		frame];
}

void CALayer_inst_SetFrame(void *id, NSRect value) {
	[(CALayer*)id
		setFrame: value];
}

NSRect CALayer_inst_Bounds(void *id) {
	return [(CALayer*)id
		bounds];
}

void CALayer_inst_SetBounds(void *id, NSRect value) {
	[(CALayer*)id
		setBounds: value];
}

double CALayer_inst_ZPosition(void *id) {
	return [(CALayer*)id
		zPosition];
}

void CALayer_inst_SetZPosition(void *id, double value) {
	[(CALayer*)id
		setZPosition: value];
}

double CALayer_inst_AnchorPointZ(void *id) {
	return [(CALayer*)id
		anchorPointZ];
}

void CALayer_inst_SetAnchorPointZ(void *id, double value) {
	[(CALayer*)id
		setAnchorPointZ: value];
}

double CALayer_inst_ContentsScale(void *id) {
	return [(CALayer*)id
		contentsScale];
}

void CALayer_inst_SetContentsScale(void *id, double value) {
	[(CALayer*)id
		setContentsScale: value];
}

void* CALayer_inst_Sublayers(void *id) {
	return [(CALayer*)id
		sublayers];
}

void CALayer_inst_SetSublayers(void *id, void* value) {
	[(CALayer*)id
		setSublayers: value];
}

void* CALayer_inst_Superlayer(void *id) {
	return [(CALayer*)id
		superlayer];
}

BOOL CALayer_inst_NeedsDisplayOnBoundsChange(void *id) {
	return [(CALayer*)id
		needsDisplayOnBoundsChange];
}

void CALayer_inst_SetNeedsDisplayOnBoundsChange(void *id, BOOL value) {
	[(CALayer*)id
		setNeedsDisplayOnBoundsChange: value];
}

void* CALayer_inst_LayoutManager(void *id) {
	return [(CALayer*)id
		layoutManager];
}

void CALayer_inst_SetLayoutManager(void *id, void* value) {
	[(CALayer*)id
		setLayoutManager: value];
}

void* CALayer_inst_Constraints(void *id) {
	return [(CALayer*)id
		constraints];
}

void CALayer_inst_SetConstraints(void *id, void* value) {
	[(CALayer*)id
		setConstraints: value];
}

void* CALayer_inst_Actions(void *id) {
	return [(CALayer*)id
		actions];
}

void CALayer_inst_SetActions(void *id, void* value) {
	[(CALayer*)id
		setActions: value];
}

NSRect CALayer_inst_VisibleRect(void *id) {
	return [(CALayer*)id
		visibleRect];
}

void* CALayer_inst_Name(void *id) {
	return [(CALayer*)id
		name];
}

void CALayer_inst_SetName(void *id, void* value) {
	[(CALayer*)id
		setName: value];
}

void* NSArray_inst_ArrayByAddingObjectsFromArray(void *id, void* otherArray) {
	return [(NSArray*)id
		arrayByAddingObjectsFromArray: otherArray];
}

void* NSArray_inst_ComponentsJoinedByString(void *id, void* separator) {
	return [(NSArray*)id
		componentsJoinedByString: separator];
}

void* NSArray_inst_DescriptionWithLocale(void *id, void* locale) {
	return [(NSArray*)id
		descriptionWithLocale: locale];
}

void* NSArray_inst_DescriptionWithLocaleIndent(void *id, void* locale, unsigned long level) {
	return [(NSArray*)id
		descriptionWithLocale: locale
		indent: level];
}

void* NSArray_inst_Init(void *id) {
	return [(NSArray*)id
		init];
}

void* NSArray_inst_InitWithArray(void *id, void* array) {
	return [(NSArray*)id
		initWithArray: array];
}

void* NSArray_inst_InitWithArrayCopyItems(void *id, void* array, BOOL flag) {
	return [(NSArray*)id
		initWithArray: array
		copyItems: flag];
}

BOOL NSArray_inst_IsEqualToArray(void *id, void* otherArray) {
	return [(NSArray*)id
		isEqualToArray: otherArray];
}

void NSArray_inst_MakeObjectsPerformSelector(void *id, void* aSelector) {
	[(NSArray*)id
		makeObjectsPerformSelector: aSelector];
}

void NSArray_inst_MakeObjectsPerformSelectorWithObject(void *id, void* aSelector, void* argument) {
	[(NSArray*)id
		makeObjectsPerformSelector: aSelector
		withObject: argument];
}

void* NSArray_inst_PathsMatchingExtensions(void *id, void* filterTypes) {
	return [(NSArray*)id
		pathsMatchingExtensions: filterTypes];
}

void NSArray_inst_RemoveObserverForKeyPath(void *id, void* observer, void* keyPath) {
	[(NSArray*)id
		removeObserver: observer
		forKeyPath: keyPath];
}

void NSArray_inst_RemoveObserverForKeyPathContext(void *id, void* observer, void* keyPath, void* context) {
	[(NSArray*)id
		removeObserver: observer
		forKeyPath: keyPath
		context: context];
}

void NSArray_inst_SetValueForKey(void *id, void* value, void* key) {
	[(NSArray*)id
		setValue: value
		forKey: key];
}

void* NSArray_inst_ShuffledArray(void *id) {
	return [(NSArray*)id
		shuffledArray];
}

void* NSArray_inst_SortedArrayUsingDescriptors(void *id, void* sortDescriptors) {
	return [(NSArray*)id
		sortedArrayUsingDescriptors: sortDescriptors];
}

void* NSArray_inst_SortedArrayUsingSelector(void *id, void* comparator) {
	return [(NSArray*)id
		sortedArrayUsingSelector: comparator];
}

void* NSArray_inst_ValueForKey(void *id, void* key) {
	return [(NSArray*)id
		valueForKey: key];
}

unsigned long NSArray_inst_Count(void *id) {
	return [(NSArray*)id
		count];
}

void* NSArray_inst_SortedArrayHint(void *id) {
	return [(NSArray*)id
		sortedArrayHint];
}

void* NSArray_inst_Description(void *id) {
	return [(NSArray*)id
		description];
}

void* NSAttributedString_inst_AttributedStringByInflectingString(void *id) {
	return [(NSAttributedString*)id
		attributedStringByInflectingString];
}

void NSAttributedString_inst_DrawInRect(void *id, NSRect rect) {
	[(NSAttributedString*)id
		drawInRect: rect];
}

void* NSAttributedString_inst_InitWithAttributedString(void *id, void* attrStr) {
	return [(NSAttributedString*)id
		initWithAttributedString: attrStr];
}

void* NSAttributedString_inst_InitWithDocFormatDocumentAttributes(void *id, void* data, void* dict) {
	return [(NSAttributedString*)id
		initWithDocFormat: data
		documentAttributes: dict];
}

void* NSAttributedString_inst_InitWithHTMLBaseURLDocumentAttributes(void *id, void* data, void* base, void* dict) {
	return [(NSAttributedString*)id
		initWithHTML: data
		baseURL: base
		documentAttributes: dict];
}

void* NSAttributedString_inst_InitWithHTMLDocumentAttributes(void *id, void* data, void* dict) {
	return [(NSAttributedString*)id
		initWithHTML: data
		documentAttributes: dict];
}

void* NSAttributedString_inst_InitWithHTMLOptionsDocumentAttributes(void *id, void* data, void* options, void* dict) {
	return [(NSAttributedString*)id
		initWithHTML: data
		options: options
		documentAttributes: dict];
}

void* NSAttributedString_inst_InitWithRTFDocumentAttributes(void *id, void* data, void* dict) {
	return [(NSAttributedString*)id
		initWithRTF: data
		documentAttributes: dict];
}

void* NSAttributedString_inst_InitWithRTFDDocumentAttributes(void *id, void* data, void* dict) {
	return [(NSAttributedString*)id
		initWithRTFD: data
		documentAttributes: dict];
}

void* NSAttributedString_inst_InitWithString(void *id, void* str) {
	return [(NSAttributedString*)id
		initWithString: str];
}

void* NSAttributedString_inst_InitWithStringAttributes(void *id, void* str, void* attrs) {
	return [(NSAttributedString*)id
		initWithString: str
		attributes: attrs];
}

BOOL NSAttributedString_inst_IsEqualToAttributedString(void *id, void* other) {
	return [(NSAttributedString*)id
		isEqualToAttributedString: other];
}

unsigned long NSAttributedString_inst_NextWordFromIndexForward(void *id, unsigned long location, BOOL isForward) {
	return [(NSAttributedString*)id
		nextWordFromIndex: location
		forward: isForward];
}

NSSize NSAttributedString_inst_Size(void *id) {
	return [(NSAttributedString*)id
		size];
}

void* NSAttributedString_inst_Init(void *id) {
	return [(NSAttributedString*)id
		init];
}

void* NSAttributedString_inst_String(void *id) {
	return [(NSAttributedString*)id
		string];
}

unsigned long NSAttributedString_inst_Length(void *id) {
	return [(NSAttributedString*)id
		length];
}

void NSData_inst_GetBytesLength(void *id, void* buffer, unsigned long length) {
	[(NSData*)id
		getBytes: buffer
		length: length];
}

void* NSData_inst_InitWithBytesLength(void *id, void* bytes, unsigned long length) {
	return [(NSData*)id
		initWithBytes: bytes
		length: length];
}

void* NSData_inst_InitWithBytesNoCopyLength(void *id, void* bytes, unsigned long length) {
	return [(NSData*)id
		initWithBytesNoCopy: bytes
		length: length];
}

void* NSData_inst_InitWithBytesNoCopyLengthFreeWhenDone(void *id, void* bytes, unsigned long length, BOOL b) {
	return [(NSData*)id
		initWithBytesNoCopy: bytes
		length: length
		freeWhenDone: b];
}

void* NSData_inst_InitWithContentsOfFile(void *id, void* path) {
	return [(NSData*)id
		initWithContentsOfFile: path];
}

void* NSData_inst_InitWithContentsOfURL(void *id, void* url) {
	return [(NSData*)id
		initWithContentsOfURL: url];
}

void* NSData_inst_InitWithData(void *id, void* data) {
	return [(NSData*)id
		initWithData: data];
}

BOOL NSData_inst_IsEqualToData(void *id, void* other) {
	return [(NSData*)id
		isEqualToData: other];
}

BOOL NSData_inst_WriteToFileAtomically(void *id, void* path, BOOL useAuxiliaryFile) {
	return [(NSData*)id
		writeToFile: path
		atomically: useAuxiliaryFile];
}

BOOL NSData_inst_WriteToURLAtomically(void *id, void* url, BOOL atomically) {
	return [(NSData*)id
		writeToURL: url
		atomically: atomically];
}

void* NSData_inst_Init(void *id) {
	return [(NSData*)id
		init];
}

void* NSData_inst_Bytes(void *id) {
	return [(NSData*)id
		bytes];
}

unsigned long NSData_inst_Length(void *id) {
	return [(NSData*)id
		length];
}

void* NSData_inst_Description(void *id) {
	return [(NSData*)id
		description];
}

void* NSDictionary_inst_DescriptionWithLocale(void *id, void* locale) {
	return [(NSDictionary*)id
		descriptionWithLocale: locale];
}

void* NSDictionary_inst_DescriptionWithLocaleIndent(void *id, void* locale, unsigned long level) {
	return [(NSDictionary*)id
		descriptionWithLocale: locale
		indent: level];
}

BOOL NSDictionary_inst_FileExtensionHidden(void *id) {
	return [(NSDictionary*)id
		fileExtensionHidden];
}

void* NSDictionary_inst_FileGroupOwnerAccountID(void *id) {
	return [(NSDictionary*)id
		fileGroupOwnerAccountID];
}

void* NSDictionary_inst_FileGroupOwnerAccountName(void *id) {
	return [(NSDictionary*)id
		fileGroupOwnerAccountName];
}

BOOL NSDictionary_inst_FileIsAppendOnly(void *id) {
	return [(NSDictionary*)id
		fileIsAppendOnly];
}

BOOL NSDictionary_inst_FileIsImmutable(void *id) {
	return [(NSDictionary*)id
		fileIsImmutable];
}

void* NSDictionary_inst_FileOwnerAccountID(void *id) {
	return [(NSDictionary*)id
		fileOwnerAccountID];
}

void* NSDictionary_inst_FileOwnerAccountName(void *id) {
	return [(NSDictionary*)id
		fileOwnerAccountName];
}

unsigned long NSDictionary_inst_FilePosixPermissions(void *id) {
	return [(NSDictionary*)id
		filePosixPermissions];
}

unsigned long NSDictionary_inst_FileSystemFileNumber(void *id) {
	return [(NSDictionary*)id
		fileSystemFileNumber];
}

long NSDictionary_inst_FileSystemNumber(void *id) {
	return [(NSDictionary*)id
		fileSystemNumber];
}

void* NSDictionary_inst_FileType(void *id) {
	return [(NSDictionary*)id
		fileType];
}

void* NSDictionary_inst_Init(void *id) {
	return [(NSDictionary*)id
		init];
}

void* NSDictionary_inst_InitWithDictionary(void *id, void* otherDictionary) {
	return [(NSDictionary*)id
		initWithDictionary: otherDictionary];
}

void* NSDictionary_inst_InitWithDictionaryCopyItems(void *id, void* otherDictionary, BOOL flag) {
	return [(NSDictionary*)id
		initWithDictionary: otherDictionary
		copyItems: flag];
}

void* NSDictionary_inst_InitWithObjectsForKeys(void *id, void* objects, void* keys) {
	return [(NSDictionary*)id
		initWithObjects: objects
		forKeys: keys];
}

BOOL NSDictionary_inst_IsEqualToDictionary(void *id, void* otherDictionary) {
	return [(NSDictionary*)id
		isEqualToDictionary: otherDictionary];
}

void* NSDictionary_inst_KeysSortedByValueUsingSelector(void *id, void* comparator) {
	return [(NSDictionary*)id
		keysSortedByValueUsingSelector: comparator];
}

unsigned long NSDictionary_inst_Count(void *id) {
	return [(NSDictionary*)id
		count];
}

void* NSDictionary_inst_AllKeys(void *id) {
	return [(NSDictionary*)id
		allKeys];
}

void* NSDictionary_inst_AllValues(void *id) {
	return [(NSDictionary*)id
		allValues];
}

void* NSDictionary_inst_Description(void *id) {
	return [(NSDictionary*)id
		description];
}

void* NSDictionary_inst_DescriptionInStringsFileFormat(void *id) {
	return [(NSDictionary*)id
		descriptionInStringsFileFormat];
}

void* NSNumber_inst_DescriptionWithLocale(void *id, void* locale) {
	return [(NSNumber*)id
		descriptionWithLocale: locale];
}

void* NSNumber_inst_InitWithBool(void *id, BOOL value) {
	return [(NSNumber*)id
		initWithBool: value];
}

void* NSNumber_inst_InitWithInt(void *id, int value) {
	return [(NSNumber*)id
		initWithInt: value];
}

void* NSNumber_inst_InitWithInteger(void *id, long value) {
	return [(NSNumber*)id
		initWithInteger: value];
}

void* NSNumber_inst_InitWithUnsignedInt(void *id, int value) {
	return [(NSNumber*)id
		initWithUnsignedInt: value];
}

void* NSNumber_inst_InitWithUnsignedInteger(void *id, unsigned long value) {
	return [(NSNumber*)id
		initWithUnsignedInteger: value];
}

BOOL NSNumber_inst_IsEqualToNumber(void *id, void* number) {
	return [(NSNumber*)id
		isEqualToNumber: number];
}

void* NSNumber_inst_Init(void *id) {
	return [(NSNumber*)id
		init];
}

BOOL NSNumber_inst_BoolValue(void *id) {
	return [(NSNumber*)id
		boolValue];
}

int NSNumber_inst_IntValue(void *id) {
	return [(NSNumber*)id
		intValue];
}

long NSNumber_inst_IntegerValue(void *id) {
	return [(NSNumber*)id
		integerValue];
}

unsigned long NSNumber_inst_UnsignedIntegerValue(void *id) {
	return [(NSNumber*)id
		unsignedIntegerValue];
}

int NSNumber_inst_UnsignedIntValue(void *id) {
	return [(NSNumber*)id
		unsignedIntValue];
}

void* NSNumber_inst_StringValue(void *id) {
	return [(NSNumber*)id
		stringValue];
}

void NSRunLoop_inst_CancelPerformSelectorTargetArgument(void *id, void* aSelector, void* target, void* arg) {
	[(NSRunLoop*)id
		cancelPerformSelector: aSelector
		target: target
		argument: arg];
}

void NSRunLoop_inst_CancelPerformSelectorsWithTarget(void *id, void* target) {
	[(NSRunLoop*)id
		cancelPerformSelectorsWithTarget: target];
}

void NSRunLoop_inst_PerformSelectorTargetArgumentOrderModes(void *id, void* aSelector, void* target, void* arg, unsigned long order, void* modes) {
	[(NSRunLoop*)id
		performSelector: aSelector
		target: target
		argument: arg
		order: order
		modes: modes];
}

void NSRunLoop_inst_Run(void *id) {
	[(NSRunLoop*)id
		run];
}

void* NSRunLoop_inst_Init(void *id) {
	return [(NSRunLoop*)id
		init];
}

BOOL NSString_inst_CanBeConvertedToEncoding(void *id, unsigned long encoding) {
	return [(NSString*)id
		canBeConvertedToEncoding: encoding];
}

unsigned short NSString_inst_CharacterAtIndex(void *id, unsigned long index) {
	return [(NSString*)id
		characterAtIndex: index];
}

unsigned long NSString_inst_CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes(void *id, void* outputName, BOOL flag, void* outputArray, void* filterTypes) {
	return [(NSString*)id
		completePathIntoString: outputName
		caseSensitive: flag
		matchesIntoArray: outputArray
		filterTypes: filterTypes];
}

void* NSString_inst_ComponentsSeparatedByString(void *id, void* separator) {
	return [(NSString*)id
		componentsSeparatedByString: separator];
}

BOOL NSString_inst_ContainsString(void *id, void* str) {
	return [(NSString*)id
		containsString: str];
}

void* NSString_inst_DataUsingEncoding(void *id, unsigned long encoding) {
	return [(NSString*)id
		dataUsingEncoding: encoding];
}

void* NSString_inst_DataUsingEncodingAllowLossyConversion(void *id, unsigned long encoding, BOOL lossy) {
	return [(NSString*)id
		dataUsingEncoding: encoding
		allowLossyConversion: lossy];
}

void NSString_inst_DrawInRectWithAttributes(void *id, NSRect rect, void* attrs) {
	[(NSString*)id
		drawInRect: rect
		withAttributes: attrs];
}

BOOL NSString_inst_HasPrefix(void *id, void* str) {
	return [(NSString*)id
		hasPrefix: str];
}

BOOL NSString_inst_HasSuffix(void *id, void* str) {
	return [(NSString*)id
		hasSuffix: str];
}

void* NSString_inst_Init(void *id) {
	return [(NSString*)id
		init];
}

void* NSString_inst_InitWithBytesLengthEncoding(void *id, void* bytes, unsigned long len, unsigned long encoding) {
	return [(NSString*)id
		initWithBytes: bytes
		length: len
		encoding: encoding];
}

void* NSString_inst_InitWithBytesNoCopyLengthEncodingFreeWhenDone(void *id, void* bytes, unsigned long len, unsigned long encoding, BOOL freeBuffer) {
	return [(NSString*)id
		initWithBytesNoCopy: bytes
		length: len
		encoding: encoding
		freeWhenDone: freeBuffer];
}

void* NSString_inst_InitWithDataEncoding(void *id, void* data, unsigned long encoding) {
	return [(NSString*)id
		initWithData: data
		encoding: encoding];
}

void* NSString_inst_InitWithString(void *id, void* aString) {
	return [(NSString*)id
		initWithString: aString];
}

BOOL NSString_inst_IsEqualToString(void *id, void* aString) {
	return [(NSString*)id
		isEqualToString: aString];
}

unsigned long NSString_inst_LengthOfBytesUsingEncoding(void *id, unsigned long enc) {
	return [(NSString*)id
		lengthOfBytesUsingEncoding: enc];
}

BOOL NSString_inst_LocalizedCaseInsensitiveContainsString(void *id, void* str) {
	return [(NSString*)id
		localizedCaseInsensitiveContainsString: str];
}

BOOL NSString_inst_LocalizedStandardContainsString(void *id, void* str) {
	return [(NSString*)id
		localizedStandardContainsString: str];
}

unsigned long NSString_inst_MaximumLengthOfBytesUsingEncoding(void *id, unsigned long enc) {
	return [(NSString*)id
		maximumLengthOfBytesUsingEncoding: enc];
}

void* NSString_inst_PropertyList(void *id) {
	return [(NSString*)id
		propertyList];
}

void* NSString_inst_PropertyListFromStringsFileFormat(void *id) {
	return [(NSString*)id
		propertyListFromStringsFileFormat];
}

NSSize NSString_inst_SizeWithAttributes(void *id, void* attrs) {
	return [(NSString*)id
		sizeWithAttributes: attrs];
}

void* NSString_inst_StringByAppendingPathComponent(void *id, void* str) {
	return [(NSString*)id
		stringByAppendingPathComponent: str];
}

void* NSString_inst_StringByAppendingPathExtension(void *id, void* str) {
	return [(NSString*)id
		stringByAppendingPathExtension: str];
}

void* NSString_inst_StringByAppendingString(void *id, void* aString) {
	return [(NSString*)id
		stringByAppendingString: aString];
}

void* NSString_inst_StringByPaddingToLengthWithStringStartingAtIndex(void *id, unsigned long newLength, void* padString, unsigned long padIndex) {
	return [(NSString*)id
		stringByPaddingToLength: newLength
		withString: padString
		startingAtIndex: padIndex];
}

void* NSString_inst_StringByReplacingOccurrencesOfStringWithString(void *id, void* target, void* replacement) {
	return [(NSString*)id
		stringByReplacingOccurrencesOfString: target
		withString: replacement];
}

void* NSString_inst_StringsByAppendingPaths(void *id, void* paths) {
	return [(NSString*)id
		stringsByAppendingPaths: paths];
}

void* NSString_inst_SubstringFromIndex(void *id, unsigned long from) {
	return [(NSString*)id
		substringFromIndex: from];
}

void* NSString_inst_SubstringToIndex(void *id, unsigned long to) {
	return [(NSString*)id
		substringToIndex: to];
}

void* NSString_inst_VariantFittingPresentationWidth(void *id, long width) {
	return [(NSString*)id
		variantFittingPresentationWidth: width];
}

unsigned long NSString_inst_Length(void *id) {
	return [(NSString*)id
		length];
}

unsigned long NSString_inst_Hash(void *id) {
	return [(NSString*)id
		hash];
}

void* NSString_inst_LowercaseString(void *id) {
	return [(NSString*)id
		lowercaseString];
}

void* NSString_inst_LocalizedLowercaseString(void *id) {
	return [(NSString*)id
		localizedLowercaseString];
}

void* NSString_inst_UppercaseString(void *id) {
	return [(NSString*)id
		uppercaseString];
}

void* NSString_inst_LocalizedUppercaseString(void *id) {
	return [(NSString*)id
		localizedUppercaseString];
}

void* NSString_inst_CapitalizedString(void *id) {
	return [(NSString*)id
		capitalizedString];
}

void* NSString_inst_LocalizedCapitalizedString(void *id) {
	return [(NSString*)id
		localizedCapitalizedString];
}

void* NSString_inst_DecomposedStringWithCanonicalMapping(void *id) {
	return [(NSString*)id
		decomposedStringWithCanonicalMapping];
}

void* NSString_inst_DecomposedStringWithCompatibilityMapping(void *id) {
	return [(NSString*)id
		decomposedStringWithCompatibilityMapping];
}

void* NSString_inst_PrecomposedStringWithCanonicalMapping(void *id) {
	return [(NSString*)id
		precomposedStringWithCanonicalMapping];
}

void* NSString_inst_PrecomposedStringWithCompatibilityMapping(void *id) {
	return [(NSString*)id
		precomposedStringWithCompatibilityMapping];
}

int NSString_inst_IntValue(void *id) {
	return [(NSString*)id
		intValue];
}

long NSString_inst_IntegerValue(void *id) {
	return [(NSString*)id
		integerValue];
}

BOOL NSString_inst_BoolValue(void *id) {
	return [(NSString*)id
		boolValue];
}

void* NSString_inst_Description(void *id) {
	return [(NSString*)id
		description];
}

unsigned long NSString_inst_FastestEncoding(void *id) {
	return [(NSString*)id
		fastestEncoding];
}

unsigned long NSString_inst_SmallestEncoding(void *id) {
	return [(NSString*)id
		smallestEncoding];
}

void* NSString_inst_PathComponents(void *id) {
	return [(NSString*)id
		pathComponents];
}

BOOL NSString_inst_IsAbsolutePath(void *id) {
	return [(NSString*)id
		isAbsolutePath];
}

void* NSString_inst_LastPathComponent(void *id) {
	return [(NSString*)id
		lastPathComponent];
}

void* NSString_inst_PathExtension(void *id) {
	return [(NSString*)id
		pathExtension];
}

void* NSString_inst_StringByAbbreviatingWithTildeInPath(void *id) {
	return [(NSString*)id
		stringByAbbreviatingWithTildeInPath];
}

void* NSString_inst_StringByDeletingLastPathComponent(void *id) {
	return [(NSString*)id
		stringByDeletingLastPathComponent];
}

void* NSString_inst_StringByDeletingPathExtension(void *id) {
	return [(NSString*)id
		stringByDeletingPathExtension];
}

void* NSString_inst_StringByExpandingTildeInPath(void *id) {
	return [(NSString*)id
		stringByExpandingTildeInPath];
}

void* NSString_inst_StringByResolvingSymlinksInPath(void *id) {
	return [(NSString*)id
		stringByResolvingSymlinksInPath];
}

void* NSString_inst_StringByStandardizingPath(void *id) {
	return [(NSString*)id
		stringByStandardizingPath];
}

void* NSString_inst_StringByRemovingPercentEncoding(void *id) {
	return [(NSString*)id
		stringByRemovingPercentEncoding];
}

void NSThread_inst_Cancel(void *id) {
	[(NSThread*)id
		cancel];
}

void* NSThread_inst_Init(void *id) {
	return [(NSThread*)id
		init];
}

void* NSThread_inst_InitWithTargetSelectorObject(void *id, void* target, void* selector, void* argument) {
	return [(NSThread*)id
		initWithTarget: target
		selector: selector
		object: argument];
}

void NSThread_inst_Main(void *id) {
	[(NSThread*)id
		main];
}

void NSThread_inst_Start(void *id) {
	[(NSThread*)id
		start];
}

BOOL NSThread_inst_IsExecuting(void *id) {
	return [(NSThread*)id
		isExecuting];
}

BOOL NSThread_inst_IsFinished(void *id) {
	return [(NSThread*)id
		isFinished];
}

BOOL NSThread_inst_IsCancelled(void *id) {
	return [(NSThread*)id
		isCancelled];
}

BOOL NSThread_inst_IsMainThread(void *id) {
	return [(NSThread*)id
		isMainThread];
}

void* NSThread_inst_Name(void *id) {
	return [(NSThread*)id
		name];
}

void NSThread_inst_SetName(void *id, void* value) {
	[(NSThread*)id
		setName: value];
}

unsigned long NSThread_inst_StackSize(void *id) {
	return [(NSThread*)id
		stackSize];
}

void NSThread_inst_SetStackSize(void *id, unsigned long value) {
	[(NSThread*)id
		setStackSize: value];
}

void* NSURL_inst_URLByAppendingPathComponent(void *id, void* pathComponent) {
	return [(NSURL*)id
		URLByAppendingPathComponent: pathComponent];
}

void* NSURL_inst_URLByAppendingPathComponentIsDirectory(void *id, void* pathComponent, BOOL isDirectory) {
	return [(NSURL*)id
		URLByAppendingPathComponent: pathComponent
		isDirectory: isDirectory];
}

void* NSURL_inst_URLByAppendingPathExtension(void *id, void* pathExtension) {
	return [(NSURL*)id
		URLByAppendingPathExtension: pathExtension];
}

void* NSURL_inst_FileReferenceURL(void *id) {
	return [(NSURL*)id
		fileReferenceURL];
}

void* NSURL_inst_InitAbsoluteURLWithDataRepresentationRelativeToURL(void *id, void* data, void* baseURL) {
	return [(NSURL*)id
		initAbsoluteURLWithDataRepresentation: data
		relativeToURL: baseURL];
}

void* NSURL_inst_InitFileURLWithPath(void *id, void* path) {
	return [(NSURL*)id
		initFileURLWithPath: path];
}

void* NSURL_inst_InitFileURLWithPathIsDirectory(void *id, void* path, BOOL isDir) {
	return [(NSURL*)id
		initFileURLWithPath: path
		isDirectory: isDir];
}

void* NSURL_inst_InitFileURLWithPathIsDirectoryRelativeToURL(void *id, void* path, BOOL isDir, void* baseURL) {
	return [(NSURL*)id
		initFileURLWithPath: path
		isDirectory: isDir
		relativeToURL: baseURL];
}

void* NSURL_inst_InitFileURLWithPathRelativeToURL(void *id, void* path, void* baseURL) {
	return [(NSURL*)id
		initFileURLWithPath: path
		relativeToURL: baseURL];
}

void* NSURL_inst_InitWithDataRepresentationRelativeToURL(void *id, void* data, void* baseURL) {
	return [(NSURL*)id
		initWithDataRepresentation: data
		relativeToURL: baseURL];
}

void* NSURL_inst_InitWithString(void *id, void* URLString) {
	return [(NSURL*)id
		initWithString: URLString];
}

void* NSURL_inst_InitWithStringRelativeToURL(void *id, void* URLString, void* baseURL) {
	return [(NSURL*)id
		initWithString: URLString
		relativeToURL: baseURL];
}

BOOL NSURL_inst_IsFileReferenceURL(void *id) {
	return [(NSURL*)id
		isFileReferenceURL];
}

void NSURL_inst_RemoveAllCachedResourceValues(void *id) {
	[(NSURL*)id
		removeAllCachedResourceValues];
}

BOOL NSURL_inst_StartAccessingSecurityScopedResource(void *id) {
	return [(NSURL*)id
		startAccessingSecurityScopedResource];
}

void NSURL_inst_StopAccessingSecurityScopedResource(void *id) {
	[(NSURL*)id
		stopAccessingSecurityScopedResource];
}

void* NSURL_inst_Init(void *id) {
	return [(NSURL*)id
		init];
}

void* NSURL_inst_DataRepresentation(void *id) {
	return [(NSURL*)id
		dataRepresentation];
}

BOOL NSURL_inst_IsFileURL(void *id) {
	return [(NSURL*)id
		isFileURL];
}

void* NSURL_inst_AbsoluteString(void *id) {
	return [(NSURL*)id
		absoluteString];
}

void* NSURL_inst_AbsoluteURL(void *id) {
	return [(NSURL*)id
		absoluteURL];
}

void* NSURL_inst_BaseURL(void *id) {
	return [(NSURL*)id
		baseURL];
}

void* NSURL_inst_Fragment(void *id) {
	return [(NSURL*)id
		fragment];
}

void* NSURL_inst_Host(void *id) {
	return [(NSURL*)id
		host];
}

void* NSURL_inst_LastPathComponent(void *id) {
	return [(NSURL*)id
		lastPathComponent];
}

void* NSURL_inst_Password(void *id) {
	return [(NSURL*)id
		password];
}

void* NSURL_inst_Path(void *id) {
	return [(NSURL*)id
		path];
}

void* NSURL_inst_PathComponents(void *id) {
	return [(NSURL*)id
		pathComponents];
}

void* NSURL_inst_PathExtension(void *id) {
	return [(NSURL*)id
		pathExtension];
}

void* NSURL_inst_Port(void *id) {
	return [(NSURL*)id
		port];
}

void* NSURL_inst_Query(void *id) {
	return [(NSURL*)id
		query];
}

void* NSURL_inst_RelativePath(void *id) {
	return [(NSURL*)id
		relativePath];
}

void* NSURL_inst_RelativeString(void *id) {
	return [(NSURL*)id
		relativeString];
}

void* NSURL_inst_ResourceSpecifier(void *id) {
	return [(NSURL*)id
		resourceSpecifier];
}

void* NSURL_inst_Scheme(void *id) {
	return [(NSURL*)id
		scheme];
}

void* NSURL_inst_StandardizedURL(void *id) {
	return [(NSURL*)id
		standardizedURL];
}

void* NSURL_inst_User(void *id) {
	return [(NSURL*)id
		user];
}

void* NSURL_inst_FilePathURL(void *id) {
	return [(NSURL*)id
		filePathURL];
}

void* NSURL_inst_URLByDeletingLastPathComponent(void *id) {
	return [(NSURL*)id
		URLByDeletingLastPathComponent];
}

void* NSURL_inst_URLByDeletingPathExtension(void *id) {
	return [(NSURL*)id
		URLByDeletingPathExtension];
}

void* NSURL_inst_URLByResolvingSymlinksInPath(void *id) {
	return [(NSURL*)id
		URLByResolvingSymlinksInPath];
}

void* NSURL_inst_URLByStandardizingPath(void *id) {
	return [(NSURL*)id
		URLByStandardizingPath];
}

BOOL NSURL_inst_HasDirectoryPath(void *id) {
	return [(NSURL*)id
		hasDirectoryPath];
}

void* NSURLRequest_inst_InitWithURL(void *id, void* URL) {
	return [(NSURLRequest*)id
		initWithURL: URL];
}

void* NSURLRequest_inst_ValueForHTTPHeaderField(void *id, void* field) {
	return [(NSURLRequest*)id
		valueForHTTPHeaderField: field];
}

void* NSURLRequest_inst_Init(void *id) {
	return [(NSURLRequest*)id
		init];
}

void* NSURLRequest_inst_HTTPMethod(void *id) {
	return [(NSURLRequest*)id
		HTTPMethod];
}

void* NSURLRequest_inst_URL(void *id) {
	return [(NSURLRequest*)id
		URL];
}

void* NSURLRequest_inst_HTTPBody(void *id) {
	return [(NSURLRequest*)id
		HTTPBody];
}

void* NSURLRequest_inst_MainDocumentURL(void *id) {
	return [(NSURLRequest*)id
		mainDocumentURL];
}

void* NSURLRequest_inst_AllHTTPHeaderFields(void *id) {
	return [(NSURLRequest*)id
		allHTTPHeaderFields];
}

BOOL NSURLRequest_inst_HTTPShouldHandleCookies(void *id) {
	return [(NSURLRequest*)id
		HTTPShouldHandleCookies];
}

BOOL NSURLRequest_inst_HTTPShouldUsePipelining(void *id) {
	return [(NSURLRequest*)id
		HTTPShouldUsePipelining];
}

BOOL NSURLRequest_inst_AllowsCellularAccess(void *id) {
	return [(NSURLRequest*)id
		allowsCellularAccess];
}

BOOL NSURLRequest_inst_AllowsConstrainedNetworkAccess(void *id) {
	return [(NSURLRequest*)id
		allowsConstrainedNetworkAccess];
}

BOOL NSURLRequest_inst_AllowsExpensiveNetworkAccess(void *id) {
	return [(NSURLRequest*)id
		allowsExpensiveNetworkAccess];
}

BOOL NSURLRequest_inst_AssumesHTTP3Capable(void *id) {
	return [(NSURLRequest*)id
		assumesHTTP3Capable];
}

void* NSUserDefaults_inst_URLForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		URLForKey: defaultName];
}

void NSUserDefaults_inst_AddSuiteNamed(void *id, void* suiteName) {
	[(NSUserDefaults*)id
		addSuiteNamed: suiteName];
}

void* NSUserDefaults_inst_ArrayForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		arrayForKey: defaultName];
}

BOOL NSUserDefaults_inst_BoolForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		boolForKey: defaultName];
}

void* NSUserDefaults_inst_DataForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		dataForKey: defaultName];
}

void* NSUserDefaults_inst_DictionaryForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		dictionaryForKey: defaultName];
}

void* NSUserDefaults_inst_DictionaryRepresentation(void *id) {
	return [(NSUserDefaults*)id
		dictionaryRepresentation];
}

void* NSUserDefaults_inst_Init(void *id) {
	return [(NSUserDefaults*)id
		init];
}

void* NSUserDefaults_inst_InitWithSuiteName(void *id, void* suitename) {
	return [(NSUserDefaults*)id
		initWithSuiteName: suitename];
}

long NSUserDefaults_inst_IntegerForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		integerForKey: defaultName];
}

void* NSUserDefaults_inst_ObjectForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		objectForKey: defaultName];
}

BOOL NSUserDefaults_inst_ObjectIsForcedForKey(void *id, void* key) {
	return [(NSUserDefaults*)id
		objectIsForcedForKey: key];
}

BOOL NSUserDefaults_inst_ObjectIsForcedForKeyInDomain(void *id, void* key, void* domain) {
	return [(NSUserDefaults*)id
		objectIsForcedForKey: key
		inDomain: domain];
}

void* NSUserDefaults_inst_PersistentDomainForName(void *id, void* domainName) {
	return [(NSUserDefaults*)id
		persistentDomainForName: domainName];
}

void NSUserDefaults_inst_RegisterDefaults(void *id, void* registrationDictionary) {
	[(NSUserDefaults*)id
		registerDefaults: registrationDictionary];
}

void NSUserDefaults_inst_RemoveObjectForKey(void *id, void* defaultName) {
	[(NSUserDefaults*)id
		removeObjectForKey: defaultName];
}

void NSUserDefaults_inst_RemovePersistentDomainForName(void *id, void* domainName) {
	[(NSUserDefaults*)id
		removePersistentDomainForName: domainName];
}

void NSUserDefaults_inst_RemoveSuiteNamed(void *id, void* suiteName) {
	[(NSUserDefaults*)id
		removeSuiteNamed: suiteName];
}

void NSUserDefaults_inst_RemoveVolatileDomainForName(void *id, void* domainName) {
	[(NSUserDefaults*)id
		removeVolatileDomainForName: domainName];
}

void NSUserDefaults_inst_SetBoolForKey(void *id, BOOL value, void* defaultName) {
	[(NSUserDefaults*)id
		setBool: value
		forKey: defaultName];
}

void NSUserDefaults_inst_SetIntegerForKey(void *id, long value, void* defaultName) {
	[(NSUserDefaults*)id
		setInteger: value
		forKey: defaultName];
}

void NSUserDefaults_inst_SetObjectForKey(void *id, void* value, void* defaultName) {
	[(NSUserDefaults*)id
		setObject: value
		forKey: defaultName];
}

void NSUserDefaults_inst_SetPersistentDomainForName(void *id, void* domain, void* domainName) {
	[(NSUserDefaults*)id
		setPersistentDomain: domain
		forName: domainName];
}

void NSUserDefaults_inst_SetURLForKey(void *id, void* url, void* defaultName) {
	[(NSUserDefaults*)id
		setURL: url
		forKey: defaultName];
}

void NSUserDefaults_inst_SetVolatileDomainForName(void *id, void* domain, void* domainName) {
	[(NSUserDefaults*)id
		setVolatileDomain: domain
		forName: domainName];
}

void* NSUserDefaults_inst_StringArrayForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		stringArrayForKey: defaultName];
}

void* NSUserDefaults_inst_StringForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		stringForKey: defaultName];
}

BOOL NSUserDefaults_inst_Synchronize(void *id) {
	return [(NSUserDefaults*)id
		synchronize];
}

void* NSUserDefaults_inst_VolatileDomainForName(void *id, void* domainName) {
	return [(NSUserDefaults*)id
		volatileDomainForName: domainName];
}

void* NSUserDefaults_inst_VolatileDomainNames(void *id) {
	return [(NSUserDefaults*)id
		volatileDomainNames];
}


BOOL core_objc_bool_true = YES;
BOOL core_objc_bool_false = NO;

*/
import "C"

func convertObjCBoolToGo(b C.BOOL) bool {
	// NOTE: the prefix here is used to namespace these since the linker will
	// otherwise report a "duplicate symbol" because the C functions have the
	// same name.
	return bool(C.core_convertObjCBool(b))
}

func convertToObjCBool(b bool) C.BOOL {
	if b {
		return C.core_objc_bool_true
	}
	return C.core_objc_bool_false
}

// NSObject_Alloc is undocumented.
func NSObject_Alloc() NSObject {
	ret := C.NSObject_type_Alloc()

	return NSObject_FromPointer(ret)
}

// NSObject_Initialize initializes the class before it receives its first message.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1418639-initialize?language=objc for details.
func NSObject_Initialize() {
	C.NSObject_type_Initialize()

	return
}

// NSObject_Load invoked whenever a class or category is added to the Objective-C runtime; implement this method to perform class-specific behavior upon loading.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1418815-load?language=objc for details.
func NSObject_Load() {
	C.NSObject_type_Load()

	return
}

// NSObject_New allocates a new instance of the receiving class, sends it an init message, and returns the initialized object.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1571948-new?language=objc for details.
func NSObject_New() NSObject {
	ret := C.NSObject_type_New()

	return NSObject_FromPointer(ret)
}

// NSObject_InstancesRespondToSelector returns a Boolean value that indicates whether instances of the receiver are capable of responding to a given selector.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1418555-instancesrespondtoselector?language=objc for details.
func NSObject_InstancesRespondToSelector(aSelector objc.Selector) bool {
	ret := C.NSObject_type_InstancesRespondToSelector(
		aSelector.SelectorAddress(),
	)

	return convertObjCBoolToGo(ret)
}

// NSObject_Description returns a string that represents the contents of the receiving class.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1418799-description?language=objc for details.
func NSObject_Description() NSString {
	ret := C.NSObject_type_Description()

	return NSString_FromPointer(ret)
}

// NSObject_CancelPreviousPerformRequestsWithTarget cancels perform requests previously registered with the performSelector:withObject:afterDelay: instance method.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1417611-cancelpreviousperformrequestswit?language=objc for details.
func NSObject_CancelPreviousPerformRequestsWithTarget(aTarget objc.Ref) {
	C.NSObject_type_CancelPreviousPerformRequestsWithTarget(
		objc.RefPointer(aTarget),
	)

	return
}

// NSObject_CancelPreviousPerformRequestsWithTargetSelectorObject cancels perform requests previously registered with performSelector:withObject:afterDelay:.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1410849-cancelpreviousperformrequestswit?language=objc for details.
func NSObject_CancelPreviousPerformRequestsWithTargetSelectorObject(aTarget objc.Ref, aSelector objc.Selector, anArgument objc.Ref) {
	C.NSObject_type_CancelPreviousPerformRequestsWithTargetSelectorObject(
		objc.RefPointer(aTarget),
		aSelector.SelectorAddress(),
		objc.RefPointer(anArgument),
	)

	return
}

// NSObject_ResolveClassMethod dynamically provides an implementation for a given selector for a class method.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1418889-resolveclassmethod?language=objc for details.
func NSObject_ResolveClassMethod(sel objc.Selector) bool {
	ret := C.NSObject_type_ResolveClassMethod(
		sel.SelectorAddress(),
	)

	return convertObjCBoolToGo(ret)
}

// NSObject_ResolveInstanceMethod dynamically provides an implementation for a given selector for an instance method.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1418500-resolveinstancemethod?language=objc for details.
func NSObject_ResolveInstanceMethod(sel objc.Selector) bool {
	ret := C.NSObject_type_ResolveInstanceMethod(
		sel.SelectorAddress(),
	)

	return convertObjCBoolToGo(ret)
}

// NSObject_ClassFallbacksForKeyedArchiver overridden to return the names of classes that can be used to decode objects if their class is unavailable.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1411048-classfallbacksforkeyedarchiver?language=objc for details.
func NSObject_ClassFallbacksForKeyedArchiver() NSArray {
	ret := C.NSObject_type_ClassFallbacksForKeyedArchiver()

	return NSArray_FromPointer(ret)
}

// NSObject_SetVersion sets the receiver's version number.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1416538-setversion?language=objc for details.
func NSObject_SetVersion(aVersion NSInteger) {
	C.NSObject_type_SetVersion(
		C.long(aVersion),
	)

	return
}

// NSObject_Version returns the version number assigned to the class.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1415151-version?language=objc for details.
func NSObject_Version() NSInteger {
	ret := C.NSObject_type_Version()

	return NSInteger(ret)
}

// NSObject_DebugDescription is undocumented.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1418711-debugdescription?language=objc for details.
func NSObject_DebugDescription() NSString {
	ret := C.NSObject_type_DebugDescription()

	return NSString_FromPointer(ret)
}

// NSObject_Hash is undocumented.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1418561-hash?language=objc for details.
func NSObject_Hash() NSUInteger {
	ret := C.NSObject_type_Hash()

	return NSUInteger(ret)
}

// CALayer_Alloc is undocumented.
func CALayer_Alloc() CALayer {
	ret := C.CALayer_type_Alloc()

	return CALayer_FromPointer(ret)
}

// CALayer_Layer creates and returns an instance of the layer object.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc for details.
func CALayer_Layer() CALayer {
	ret := C.CALayer_type_Layer()

	return CALayer_FromPointer(ret)
}

// CALayer_NeedsDisplayForKey returns a Boolean indicating whether changes to the specified key require the layer to be redisplayed.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410769-needsdisplayforkey?language=objc for details.
func CALayer_NeedsDisplayForKey(key NSStringRef) bool {
	ret := C.CALayer_type_NeedsDisplayForKey(
		objc.RefPointer(key),
	)

	return convertObjCBoolToGo(ret)
}

// CALayer_DefaultActionForKey returns the default action for the current class.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410954-defaultactionforkey?language=objc for details.
func CALayer_DefaultActionForKey(event NSStringRef) objc.Object {
	ret := C.CALayer_type_DefaultActionForKey(
		objc.RefPointer(event),
	)

	return objc.Object_FromPointer(ret)
}

// CALayer_DefaultValueForKey specifies the default value associated with the specified key.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410886-defaultvalueforkey?language=objc for details.
func CALayer_DefaultValueForKey(key NSStringRef) objc.Object {
	ret := C.CALayer_type_DefaultValueForKey(
		objc.RefPointer(key),
	)

	return objc.Object_FromPointer(ret)
}

// NSArray_Alloc is undocumented.
func NSArray_Alloc() NSArray {
	ret := C.NSArray_type_Alloc()

	return NSArray_FromPointer(ret)
}

// NSArray_Array creates and returns an empty array.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1460120-array?language=objc for details.
func NSArray_Array() NSArray {
	ret := C.NSArray_type_Array()

	return NSArray_FromPointer(ret)
}

// NSArray_ArrayWithArray creates and returns an array containing the objects in another given array.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1460122-arraywitharray?language=objc for details.
func NSArray_ArrayWithArray(array NSArrayRef) NSArray {
	ret := C.NSArray_type_ArrayWithArray(
		objc.RefPointer(array),
	)

	return NSArray_FromPointer(ret)
}

// NSAttributedString_Alloc is undocumented.
func NSAttributedString_Alloc() NSAttributedString {
	ret := C.NSAttributedString_type_Alloc()

	return NSAttributedString_FromPointer(ret)
}

// NSAttributedString_TextTypes an array of UTI strings that identify the file types that attributed strings support, either directly or through a user-installed filter service.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1535409-texttypes?language=objc for details.
func NSAttributedString_TextTypes() NSArray {
	ret := C.NSAttributedString_type_TextTypes()

	return NSArray_FromPointer(ret)
}

// NSAttributedString_TextUnfilteredTypes an array of UTI strings that identify the file types that attributed strings support directly.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1528269-textunfilteredtypes?language=objc for details.
func NSAttributedString_TextUnfilteredTypes() NSArray {
	ret := C.NSAttributedString_type_TextUnfilteredTypes()

	return NSArray_FromPointer(ret)
}

// NSData_Alloc is undocumented.
func NSData_Alloc() NSData {
	ret := C.NSData_type_Alloc()

	return NSData_FromPointer(ret)
}

// NSData_Data creates an empty data object.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1547234-data?language=objc for details.
func NSData_Data() NSData {
	ret := C.NSData_type_Data()

	return NSData_FromPointer(ret)
}

// NSData_DataWithBytesLength creates a data object containing a given number of bytes copied from a given buffer.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1547231-datawithbytes?language=objc for details.
func NSData_DataWithBytesLength(bytes unsafe.Pointer, length NSUInteger) NSData {
	ret := C.NSData_type_DataWithBytesLength(
		bytes,
		C.ulong(length),
	)

	return NSData_FromPointer(ret)
}

// NSData_DataWithBytesNoCopyLength creates a data object that holds a given number of bytes from a given buffer.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1547229-datawithbytesnocopy?language=objc for details.
func NSData_DataWithBytesNoCopyLength(bytes unsafe.Pointer, length NSUInteger) NSData {
	ret := C.NSData_type_DataWithBytesNoCopyLength(
		bytes,
		C.ulong(length),
	)

	return NSData_FromPointer(ret)
}

// NSData_DataWithBytesNoCopyLengthFreeWhenDone creates a data object that holds a given number of bytes from a given buffer.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1547240-datawithbytesnocopy?language=objc for details.
func NSData_DataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length NSUInteger, b bool) NSData {
	ret := C.NSData_type_DataWithBytesNoCopyLengthFreeWhenDone(
		bytes,
		C.ulong(length),
		convertToObjCBool(b),
	)

	return NSData_FromPointer(ret)
}

// NSData_DataWithData creates a data object containing the contents of another data object.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1547230-datawithdata?language=objc for details.
func NSData_DataWithData(data NSDataRef) NSData {
	ret := C.NSData_type_DataWithData(
		objc.RefPointer(data),
	)

	return NSData_FromPointer(ret)
}

// NSData_DataWithContentsOfFile creates a data object by reading every byte from the file at a given path.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1547226-datawithcontentsoffile?language=objc for details.
func NSData_DataWithContentsOfFile(path NSStringRef) NSData {
	ret := C.NSData_type_DataWithContentsOfFile(
		objc.RefPointer(path),
	)

	return NSData_FromPointer(ret)
}

// NSData_DataWithContentsOfURL creates a data object containing the data from the location specified by a given URL.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1547245-datawithcontentsofurl?language=objc for details.
func NSData_DataWithContentsOfURL(url NSURLRef) NSData {
	ret := C.NSData_type_DataWithContentsOfURL(
		objc.RefPointer(url),
	)

	return NSData_FromPointer(ret)
}

// NSDictionary_Alloc is undocumented.
func NSDictionary_Alloc() NSDictionary {
	ret := C.NSDictionary_type_Alloc()

	return NSDictionary_FromPointer(ret)
}

// NSDictionary_Dictionary creates an empty dictionary.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1574180-dictionary?language=objc for details.
func NSDictionary_Dictionary() NSDictionary {
	ret := C.NSDictionary_type_Dictionary()

	return NSDictionary_FromPointer(ret)
}

// NSDictionary_DictionaryWithObjectsForKeys creates a dictionary containing entries constructed from the contents of an array of keys and an array of values.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1574183-dictionarywithobjects?language=objc for details.
func NSDictionary_DictionaryWithObjectsForKeys(objects NSArrayRef, keys NSArrayRef) NSDictionary {
	ret := C.NSDictionary_type_DictionaryWithObjectsForKeys(
		objc.RefPointer(objects),
		objc.RefPointer(keys),
	)

	return NSDictionary_FromPointer(ret)
}

// NSDictionary_DictionaryWithDictionary creates a dictionary containing the keys and values from another given dictionary.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1574191-dictionarywithdictionary?language=objc for details.
func NSDictionary_DictionaryWithDictionary(dict NSDictionaryRef) NSDictionary {
	ret := C.NSDictionary_type_DictionaryWithDictionary(
		objc.RefPointer(dict),
	)

	return NSDictionary_FromPointer(ret)
}

// NSDictionary_SharedKeySetForKeys creates a shared key set object for the specified keys.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1408190-sharedkeysetforkeys?language=objc for details.
func NSDictionary_SharedKeySetForKeys(keys NSArrayRef) objc.Object {
	ret := C.NSDictionary_type_SharedKeySetForKeys(
		objc.RefPointer(keys),
	)

	return objc.Object_FromPointer(ret)
}

// NSNumber_Alloc is undocumented.
func NSNumber_Alloc() NSNumber {
	ret := C.NSNumber_type_Alloc()

	return NSNumber_FromPointer(ret)
}

// NSNumber_NumberWithBool creates and returns an NSNumber object containing a given value, treating it as a BOOL.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1551475-numberwithbool?language=objc for details.
func NSNumber_NumberWithBool(value bool) NSNumber {
	ret := C.NSNumber_type_NumberWithBool(
		convertToObjCBool(value),
	)

	return NSNumber_FromPointer(ret)
}

// NSNumber_NumberWithInt creates and returns an NSNumber object containing a given value, treating it as a signed int.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1551470-numberwithint?language=objc for details.
func NSNumber_NumberWithInt(value int32) NSNumber {
	ret := C.NSNumber_type_NumberWithInt(
		C.int(value),
	)

	return NSNumber_FromPointer(ret)
}

// NSNumber_NumberWithInteger creates and returns an NSNumber object containing a given value, treating it as an NSInteger.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1551473-numberwithinteger?language=objc for details.
func NSNumber_NumberWithInteger(value NSInteger) NSNumber {
	ret := C.NSNumber_type_NumberWithInteger(
		C.long(value),
	)

	return NSNumber_FromPointer(ret)
}

// NSNumber_NumberWithUnsignedInt creates and returns an NSNumber object containing a given value, treating it as an unsigned int.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1551472-numberwithunsignedint?language=objc for details.
func NSNumber_NumberWithUnsignedInt(value int32) NSNumber {
	ret := C.NSNumber_type_NumberWithUnsignedInt(
		C.int(value),
	)

	return NSNumber_FromPointer(ret)
}

// NSNumber_NumberWithUnsignedInteger creates and returns an NSNumber object containing a given value, treating it as an NSUInteger.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1551469-numberwithunsignedinteger?language=objc for details.
func NSNumber_NumberWithUnsignedInteger(value NSUInteger) NSNumber {
	ret := C.NSNumber_type_NumberWithUnsignedInteger(
		C.ulong(value),
	)

	return NSNumber_FromPointer(ret)
}

// NSRunLoop_Alloc is undocumented.
func NSRunLoop_Alloc() NSRunLoop {
	ret := C.NSRunLoop_type_Alloc()

	return NSRunLoop_FromPointer(ret)
}

// NSRunLoop_CurrentRunLoop returns the run loop for the current thread.
//
// See https://developer.apple.com/documentation/foundation/nsrunloop/1412291-currentrunloop?language=objc for details.
func NSRunLoop_CurrentRunLoop() NSRunLoop {
	ret := C.NSRunLoop_type_CurrentRunLoop()

	return NSRunLoop_FromPointer(ret)
}

// NSRunLoop_MainRunLoop returns the run loop of the main thread.
//
// See https://developer.apple.com/documentation/foundation/nsrunloop/1418388-mainrunloop?language=objc for details.
func NSRunLoop_MainRunLoop() NSRunLoop {
	ret := C.NSRunLoop_type_MainRunLoop()

	return NSRunLoop_FromPointer(ret)
}

// NSString_Alloc is undocumented.
func NSString_Alloc() NSString {
	ret := C.NSString_type_Alloc()

	return NSString_FromPointer(ret)
}

// NSString_String returns an empty string.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1497312-string?language=objc for details.
func NSString_String() NSString {
	ret := C.NSString_type_String()

	return NSString_FromPointer(ret)
}

// NSString_LocalizedUserNotificationStringForKeyArguments returns a localized string intended for display in a notification alert.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1649585-localizedusernotificationstringf?language=objc for details.
func NSString_LocalizedUserNotificationStringForKeyArguments(key NSStringRef, arguments NSArrayRef) NSString {
	ret := C.NSString_type_LocalizedUserNotificationStringForKeyArguments(
		objc.RefPointer(key),
		objc.RefPointer(arguments),
	)

	return NSString_FromPointer(ret)
}

// NSString_StringWithString returns a string created by copying the characters from another given string.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1497372-stringwithstring?language=objc for details.
func NSString_StringWithString(string NSStringRef) NSString {
	ret := C.NSString_type_StringWithString(
		objc.RefPointer(string),
	)

	return NSString_FromPointer(ret)
}

// NSString_LocalizedNameOfStringEncoding returns a human-readable string giving the name of a given encoding.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1408318-localizednameofstringencoding?language=objc for details.
func NSString_LocalizedNameOfStringEncoding(encoding NSStringEncoding) NSString {
	ret := C.NSString_type_LocalizedNameOfStringEncoding(
		C.ulong(encoding),
	)

	return NSString_FromPointer(ret)
}

// NSString_PathWithComponents returns a string built from the strings in a given array by concatenating them with a path separator between each pair.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1417198-pathwithcomponents?language=objc for details.
func NSString_PathWithComponents(components NSArrayRef) NSString {
	ret := C.NSString_type_PathWithComponents(
		objc.RefPointer(components),
	)

	return NSString_FromPointer(ret)
}

// NSString_DefaultCStringEncoding returns the C-string encoding assumed for any method accepting a C string as an argument.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1410091-defaultcstringencoding?language=objc for details.
func NSString_DefaultCStringEncoding() NSStringEncoding {
	ret := C.NSString_type_DefaultCStringEncoding()

	return NSStringEncoding(ret)
}

// NSThread_Alloc is undocumented.
func NSThread_Alloc() NSThread {
	ret := C.NSThread_type_Alloc()

	return NSThread_FromPointer(ret)
}

// NSThread_DetachNewThreadSelectorToTargetWithObject detaches a new thread and uses the specified selector as the thread entry point.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1415633-detachnewthreadselector?language=objc for details.
func NSThread_DetachNewThreadSelectorToTargetWithObject(selector objc.Selector, target objc.Ref, argument objc.Ref) {
	C.NSThread_type_DetachNewThreadSelectorToTargetWithObject(
		selector.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(argument),
	)

	return
}

// NSThread_Exit terminates the current thread.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1409404-exit?language=objc for details.
func NSThread_Exit() {
	C.NSThread_type_Exit()

	return
}

// NSThread_IsMultiThreaded returns whether the application is multithreaded.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1410702-ismultithreaded?language=objc for details.
func NSThread_IsMultiThreaded() bool {
	ret := C.NSThread_type_IsMultiThreaded()

	return convertObjCBoolToGo(ret)
}

// NSThread_IsMainThread returns a Boolean value that indicates whether the current thread is the main thread.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1412704-ismainthread?language=objc for details.
func NSThread_IsMainThread() bool {
	ret := C.NSThread_type_IsMainThread()

	return convertObjCBoolToGo(ret)
}

// NSThread_MainThread returns the NSThread object representing the main thread.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1414782-mainthread?language=objc for details.
func NSThread_MainThread() NSThread {
	ret := C.NSThread_type_MainThread()

	return NSThread_FromPointer(ret)
}

// NSThread_CurrentThread returns the thread object representing the current thread of execution.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1410679-currentthread?language=objc for details.
func NSThread_CurrentThread() NSThread {
	ret := C.NSThread_type_CurrentThread()

	return NSThread_FromPointer(ret)
}

// NSThread_CallStackReturnAddresses returns an array containing the call stack return addresses.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1409565-callstackreturnaddresses?language=objc for details.
func NSThread_CallStackReturnAddresses() NSArray {
	ret := C.NSThread_type_CallStackReturnAddresses()

	return NSArray_FromPointer(ret)
}

// NSThread_CallStackSymbols returns an array containing the call stack symbols.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1414836-callstacksymbols?language=objc for details.
func NSThread_CallStackSymbols() NSArray {
	ret := C.NSThread_type_CallStackSymbols()

	return NSArray_FromPointer(ret)
}

// NSURL_Alloc is undocumented.
func NSURL_Alloc() NSURL {
	ret := C.NSURL_type_Alloc()

	return NSURL_FromPointer(ret)
}

// NSURL_URLWithString creates and returns an NSURL object initialized with a provided URL string.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1572047-urlwithstring?language=objc for details.
func NSURL_URLWithString(URLString NSStringRef) NSURL {
	ret := C.NSURL_type_URLWithString(
		objc.RefPointer(URLString),
	)

	return NSURL_FromPointer(ret)
}

// NSURL_URLWithStringRelativeToURL creates and returns an NSURL object initialized with a base URL and a relative string.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1572049-urlwithstring?language=objc for details.
func NSURL_URLWithStringRelativeToURL(URLString NSStringRef, baseURL NSURLRef) NSURL {
	ret := C.NSURL_type_URLWithStringRelativeToURL(
		objc.RefPointer(URLString),
		objc.RefPointer(baseURL),
	)

	return NSURL_FromPointer(ret)
}

// NSURL_FileURLWithPathIsDirectory initializes and returns a newly created NSURL object as a file URL with a specified path.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1414650-fileurlwithpath?language=objc for details.
func NSURL_FileURLWithPathIsDirectory(path NSStringRef, isDir bool) NSURL {
	ret := C.NSURL_type_FileURLWithPathIsDirectory(
		objc.RefPointer(path),
		convertToObjCBool(isDir),
	)

	return NSURL_FromPointer(ret)
}

// NSURL_FileURLWithPathRelativeToURL is undocumented.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1413201-fileurlwithpath?language=objc for details.
func NSURL_FileURLWithPathRelativeToURL(path NSStringRef, baseURL NSURLRef) NSURL {
	ret := C.NSURL_type_FileURLWithPathRelativeToURL(
		objc.RefPointer(path),
		objc.RefPointer(baseURL),
	)

	return NSURL_FromPointer(ret)
}

// NSURL_FileURLWithPathIsDirectoryRelativeToURL is undocumented.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1413020-fileurlwithpath?language=objc for details.
func NSURL_FileURLWithPathIsDirectoryRelativeToURL(path NSStringRef, isDir bool, baseURL NSURLRef) NSURL {
	ret := C.NSURL_type_FileURLWithPathIsDirectoryRelativeToURL(
		objc.RefPointer(path),
		convertToObjCBool(isDir),
		objc.RefPointer(baseURL),
	)

	return NSURL_FromPointer(ret)
}

// NSURL_FileURLWithPath initializes and returns a newly created NSURL object as a file URL with a specified path.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1410828-fileurlwithpath?language=objc for details.
func NSURL_FileURLWithPath(path NSStringRef) NSURL {
	ret := C.NSURL_type_FileURLWithPath(
		objc.RefPointer(path),
	)

	return NSURL_FromPointer(ret)
}

// NSURL_FileURLWithPathComponents initializes and returns a newly created NSURL object as a file URL with specified path components.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1414206-fileurlwithpathcomponents?language=objc for details.
func NSURL_FileURLWithPathComponents(components NSArrayRef) NSURL {
	ret := C.NSURL_type_FileURLWithPathComponents(
		objc.RefPointer(components),
	)

	return NSURL_FromPointer(ret)
}

// NSURL_AbsoluteURLWithDataRepresentationRelativeToURL is undocumented.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1412404-absoluteurlwithdatarepresentatio?language=objc for details.
func NSURL_AbsoluteURLWithDataRepresentationRelativeToURL(data NSDataRef, baseURL NSURLRef) NSURL {
	ret := C.NSURL_type_AbsoluteURLWithDataRepresentationRelativeToURL(
		objc.RefPointer(data),
		objc.RefPointer(baseURL),
	)

	return NSURL_FromPointer(ret)
}

// NSURL_URLWithDataRepresentationRelativeToURL is undocumented.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1572042-urlwithdatarepresentation?language=objc for details.
func NSURL_URLWithDataRepresentationRelativeToURL(data NSDataRef, baseURL NSURLRef) NSURL {
	ret := C.NSURL_type_URLWithDataRepresentationRelativeToURL(
		objc.RefPointer(data),
		objc.RefPointer(baseURL),
	)

	return NSURL_FromPointer(ret)
}

// NSURL_ResourceValuesForKeysFromBookmarkData returns the resource values for properties identified by a specified array of keys contained in specified bookmark data.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1418097-resourcevaluesforkeys?language=objc for details.
func NSURL_ResourceValuesForKeysFromBookmarkData(keys NSArrayRef, bookmarkData NSDataRef) NSDictionary {
	ret := C.NSURL_type_ResourceValuesForKeysFromBookmarkData(
		objc.RefPointer(keys),
		objc.RefPointer(bookmarkData),
	)

	return NSDictionary_FromPointer(ret)
}

// NSURLRequest_Alloc is undocumented.
func NSURLRequest_Alloc() NSURLRequest {
	ret := C.NSURLRequest_type_Alloc()

	return NSURLRequest_FromPointer(ret)
}

// NSURLRequest_RequestWithURL creates and returns a URL request for a specified URL.
//
// See https://developer.apple.com/documentation/foundation/nsurlrequest/1528603-requestwithurl?language=objc for details.
func NSURLRequest_RequestWithURL(URL NSURLRef) NSURLRequest {
	ret := C.NSURLRequest_type_RequestWithURL(
		objc.RefPointer(URL),
	)

	return NSURLRequest_FromPointer(ret)
}

// NSURLRequest_SupportsSecureCoding returns a Boolean value indicating whether the NSURLRequest implements the NSSecureCoding protocol.
//
// See https://developer.apple.com/documentation/foundation/nsurlrequest/1416510-supportssecurecoding?language=objc for details.
func NSURLRequest_SupportsSecureCoding() bool {
	ret := C.NSURLRequest_type_SupportsSecureCoding()

	return convertObjCBoolToGo(ret)
}

// NSUserDefaults_Alloc is undocumented.
func NSUserDefaults_Alloc() NSUserDefaults {
	ret := C.NSUserDefaults_type_Alloc()

	return NSUserDefaults_FromPointer(ret)
}

// NSUserDefaults_ResetStandardUserDefaults this method has no effect and shouldn't be used.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1407708-resetstandarduserdefaults?language=objc for details.
func NSUserDefaults_ResetStandardUserDefaults() {
	C.NSUserDefaults_type_ResetStandardUserDefaults()

	return
}

// NSUserDefaults_StandardUserDefaults returns the shared defaults object.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1416603-standarduserdefaults?language=objc for details.
func NSUserDefaults_StandardUserDefaults() NSUserDefaults {
	ret := C.NSUserDefaults_type_StandardUserDefaults()

	return NSUserDefaults_FromPointer(ret)
}

type NSObjectRef interface {
	Pointer() uintptr
	Init_AsNSObject() NSObject
}

type gen_NSObject struct {
	objc.Object
}

func NSObject_FromPointer(ptr unsafe.Pointer) NSObject {
	return NSObject{gen_NSObject{
		objc.Object_FromPointer(ptr),
	}}
}

func NSObject_FromRef(ref objc.Ref) NSObject {
	return NSObject_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// ActionProperty sent to the delegate to request the property the action applies to.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1411302-actionproperty?language=objc for details.
func (x gen_NSObject) ActionProperty() NSString {
	ret := C.NSObject_inst_ActionProperty(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// Candidates returns an array of candidates.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1385360-candidates?language=objc for details.
func (x gen_NSObject) Candidates(
	sender objc.Ref,
) NSArray {
	ret := C.NSObject_inst_Candidates(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return NSArray_FromPointer(ret)
}

// CommitComposition informs the controller that the composition should be committed.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1385539-commitcomposition?language=objc for details.
func (x gen_NSObject) CommitComposition(
	sender objc.Ref,
) {
	C.NSObject_inst_CommitComposition(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ComposedString return the current composed string.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1385416-composedstring?language=objc for details.
func (x gen_NSObject) ComposedString(
	sender objc.Ref,
) objc.Object {
	ret := C.NSObject_inst_ComposedString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return objc.Object_FromPointer(ret)
}

// Copy returns the object returned by copyWithZone:.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1418807-copy?language=objc for details.
func (x gen_NSObject) Copy() objc.Object {
	ret := C.NSObject_inst_Copy(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// CopyScriptingValueForKeyWithProperties creates and returns one or more scripting objects to be inserted into the specified relationship by copying the passed-in value and setting the properties in the copied object or objects.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1410291-copyscriptingvalue?language=objc for details.
func (x gen_NSObject) CopyScriptingValueForKeyWithProperties(
	value objc.Ref,
	key NSStringRef,
	properties NSDictionaryRef,
) objc.Object {
	ret := C.NSObject_inst_CopyScriptingValueForKeyWithProperties(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
		objc.RefPointer(key),
		objc.RefPointer(properties),
	)

	return objc.Object_FromPointer(ret)
}

// Dealloc deallocates the memory occupied by the receiver.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1571947-dealloc?language=objc for details.
func (x gen_NSObject) Dealloc() {
	C.NSObject_inst_Dealloc(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DidCommandBySelectorClient processes a command generated by user action such as typing certain keys or pressing the mouse button.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1385394-didcommandbyselector?language=objc for details.
func (x gen_NSObject) DidCommandBySelectorClient(
	aSelector objc.Selector,
	sender objc.Ref,
) bool {
	ret := C.NSObject_inst_DidCommandBySelectorClient(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
		objc.RefPointer(sender),
	)

	return convertObjCBoolToGo(ret)
}

// DoesContain returns a Boolean value that indicates whether the receiver contains a given object.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1393848-doescontain?language=objc for details.
func (x gen_NSObject) DoesContain(
	object objc.Ref,
) bool {
	ret := C.NSObject_inst_DoesContain(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)

	return convertObjCBoolToGo(ret)
}

// DoesNotRecognizeSelector handles messages the receiver doesn’t recognize.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1418637-doesnotrecognizeselector?language=objc for details.
func (x gen_NSObject) DoesNotRecognizeSelector(
	aSelector objc.Selector,
) {
	C.NSObject_inst_DoesNotRecognizeSelector(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
	)

	return
}

// ForwardingTargetForSelector returns the object to which unrecognized messages should first be directed.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1418855-forwardingtargetforselector?language=objc for details.
func (x gen_NSObject) ForwardingTargetForSelector(
	aSelector objc.Selector,
) objc.Object {
	ret := C.NSObject_inst_ForwardingTargetForSelector(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
	)

	return objc.Object_FromPointer(ret)
}

// ImageRepresentation returns the image to display.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1504801-imagerepresentation?language=objc for details.
func (x gen_NSObject) ImageRepresentation() objc.Object {
	ret := C.NSObject_inst_ImageRepresentation(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// ImageRepresentationType returns the representation type of the image to display.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1503547-imagerepresentationtype?language=objc for details.
func (x gen_NSObject) ImageRepresentationType() NSString {
	ret := C.NSObject_inst_ImageRepresentationType(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// ImageSubtitle returns the display subtitle of the image.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1503725-imagesubtitle?language=objc for details.
func (x gen_NSObject) ImageSubtitle() NSString {
	ret := C.NSObject_inst_ImageSubtitle(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// ImageTitle returns the display title of the image.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1504080-imagetitle?language=objc for details.
func (x gen_NSObject) ImageTitle() NSString {
	ret := C.NSObject_inst_ImageTitle(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// ImageUID returns a unique string that identifies the data source item.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1503516-imageuid?language=objc for details.
func (x gen_NSObject) ImageUID() NSString {
	ret := C.NSObject_inst_ImageUID(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// ImageVersion returns the version of the item.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1504444-imageversion?language=objc for details.
func (x gen_NSObject) ImageVersion() NSUInteger {
	ret := C.NSObject_inst_ImageVersion(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)
}

// Init implemented by subclasses to initialize a new object (the receiver) immediately after memory for it has been allocated.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1418641-init?language=objc for details.
func (x gen_NSObject) Init_AsNSObject() NSObject {
	ret := C.NSObject_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSObject_FromPointer(ret)
}

// InputTextClient handles key down events that do not map to an action method.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1385446-inputtext?language=objc for details.
func (x gen_NSObject) InputTextClient(
	string NSStringRef,
	sender objc.Ref,
) bool {
	ret := C.NSObject_inst_InputTextClient(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		objc.RefPointer(sender),
	)

	return convertObjCBoolToGo(ret)
}

// InputTextKeyModifiersClient receives Unicode, the key code that generated it, and any modifier flags.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1385436-inputtext?language=objc for details.
func (x gen_NSObject) InputTextKeyModifiersClient(
	string NSStringRef,
	keyCode NSInteger,
	flags NSUInteger,
	sender objc.Ref,
) bool {
	ret := C.NSObject_inst_InputTextKeyModifiersClient(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		C.long(keyCode),
		C.ulong(flags),
		objc.RefPointer(sender),
	)

	return convertObjCBoolToGo(ret)
}

// InverseForRelationshipKey for a given key that defines the name of the relationship from the receiver’s class to another class, returns the name of the relationship from the other class to the receiver’s class.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1411046-inverseforrelationshipkey?language=objc for details.
func (x gen_NSObject) InverseForRelationshipKey(
	relationshipKey NSStringRef,
) NSString {
	ret := C.NSObject_inst_InverseForRelationshipKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(relationshipKey),
	)

	return NSString_FromPointer(ret)
}

// IsCaseInsensitiveLike returns a Boolean value that indicates whether receiver is considered to be “like” a given string when the case of characters in the receiver is ignored.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1393837-iscaseinsensitivelike?language=objc for details.
func (x gen_NSObject) IsCaseInsensitiveLike(
	object NSStringRef,
) bool {
	ret := C.NSObject_inst_IsCaseInsensitiveLike(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)

	return convertObjCBoolToGo(ret)
}

// IsEqualTo returns a Boolean value that indicates whether the receiver is equal to another given object.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1393823-isequalto?language=objc for details.
func (x gen_NSObject) IsEqualTo(
	object objc.Ref,
) bool {
	ret := C.NSObject_inst_IsEqualTo(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)

	return convertObjCBoolToGo(ret)
}

// IsGreaterThan returns a Boolean value that indicates whether the receiver is greater than another given object.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1393885-isgreaterthan?language=objc for details.
func (x gen_NSObject) IsGreaterThan(
	object objc.Ref,
) bool {
	ret := C.NSObject_inst_IsGreaterThan(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)

	return convertObjCBoolToGo(ret)
}

// IsGreaterThanOrEqualTo returns a Boolean value that indicates whether the receiver is greater than or equal to another given object.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1393862-isgreaterthanorequalto?language=objc for details.
func (x gen_NSObject) IsGreaterThanOrEqualTo(
	object objc.Ref,
) bool {
	ret := C.NSObject_inst_IsGreaterThanOrEqualTo(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)

	return convertObjCBoolToGo(ret)
}

// IsLessThan returns a Boolean value that indicates whether the receiver is less than another given object.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1393841-islessthan?language=objc for details.
func (x gen_NSObject) IsLessThan(
	object objc.Ref,
) bool {
	ret := C.NSObject_inst_IsLessThan(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)

	return convertObjCBoolToGo(ret)
}

// IsLessThanOrEqualTo returns a Boolean value that indicates whether the receiver is less than or equal to another given object.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1393827-islessthanorequalto?language=objc for details.
func (x gen_NSObject) IsLessThanOrEqualTo(
	object objc.Ref,
) bool {
	ret := C.NSObject_inst_IsLessThanOrEqualTo(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)

	return convertObjCBoolToGo(ret)
}

// IsLike returns a Boolean value that indicates whether the receiver is "like" another given object.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1393866-islike?language=objc for details.
func (x gen_NSObject) IsLike(
	object NSStringRef,
) bool {
	ret := C.NSObject_inst_IsLike(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)

	return convertObjCBoolToGo(ret)
}

// IsNotEqualTo returns a Boolean value that indicates whether the receiver is not equal to another given object.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1393843-isnotequalto?language=objc for details.
func (x gen_NSObject) IsNotEqualTo(
	object objc.Ref,
) bool {
	ret := C.NSObject_inst_IsNotEqualTo(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(object),
	)

	return convertObjCBoolToGo(ret)
}

// MutableCopy returns the object returned by mutableCopyWithZone: where the zone is nil.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1418978-mutablecopy?language=objc for details.
func (x gen_NSObject) MutableCopy() objc.Object {
	ret := C.NSObject_inst_MutableCopy(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// OriginalString return the string that consists of the precomposed Unicode characters.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1385400-originalstring?language=objc for details.
func (x gen_NSObject) OriginalString(
	sender objc.Ref,
) NSAttributedString {
	ret := C.NSObject_inst_OriginalString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return NSAttributedString_FromPointer(ret)
}

// PerformSelectorOnThreadWithObjectWaitUntilDone invokes a method of the receiver on the specified thread using the default mode.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1414476-performselector?language=objc for details.
func (x gen_NSObject) PerformSelectorOnThreadWithObjectWaitUntilDone(
	aSelector objc.Selector,
	thr NSThreadRef,
	arg objc.Ref,
	wait bool,
) {
	C.NSObject_inst_PerformSelectorOnThreadWithObjectWaitUntilDone(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
		objc.RefPointer(thr),
		objc.RefPointer(arg),
		convertToObjCBool(wait),
	)

	return
}

// PerformSelectorOnThreadWithObjectWaitUntilDoneModes invokes a method of the receiver on the specified thread using the specified modes.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1417922-performselector?language=objc for details.
func (x gen_NSObject) PerformSelectorOnThreadWithObjectWaitUntilDoneModes(
	aSelector objc.Selector,
	thr NSThreadRef,
	arg objc.Ref,
	wait bool,
	array NSArrayRef,
) {
	C.NSObject_inst_PerformSelectorOnThreadWithObjectWaitUntilDoneModes(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
		objc.RefPointer(thr),
		objc.RefPointer(arg),
		convertToObjCBool(wait),
		objc.RefPointer(array),
	)

	return
}

// PerformSelectorInBackgroundWithObject invokes a method of the receiver on a new background thread.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1412390-performselectorinbackground?language=objc for details.
func (x gen_NSObject) PerformSelectorInBackgroundWithObject(
	aSelector objc.Selector,
	arg objc.Ref,
) {
	C.NSObject_inst_PerformSelectorInBackgroundWithObject(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
		objc.RefPointer(arg),
	)

	return
}

// PerformSelectorOnMainThreadWithObjectWaitUntilDone invokes a method of the receiver on the main thread using the default mode.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1414900-performselectoronmainthread?language=objc for details.
func (x gen_NSObject) PerformSelectorOnMainThreadWithObjectWaitUntilDone(
	aSelector objc.Selector,
	arg objc.Ref,
	wait bool,
) {
	C.NSObject_inst_PerformSelectorOnMainThreadWithObjectWaitUntilDone(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
		objc.RefPointer(arg),
		convertToObjCBool(wait),
	)

	return
}

// PerformSelectorOnMainThreadWithObjectWaitUntilDoneModes invokes a method of the receiver on the main thread using the specified modes.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1411637-performselectoronmainthread?language=objc for details.
func (x gen_NSObject) PerformSelectorOnMainThreadWithObjectWaitUntilDoneModes(
	aSelector objc.Selector,
	arg objc.Ref,
	wait bool,
	array NSArrayRef,
) {
	C.NSObject_inst_PerformSelectorOnMainThreadWithObjectWaitUntilDoneModes(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
		objc.RefPointer(arg),
		convertToObjCBool(wait),
		objc.RefPointer(array),
	)

	return
}

// AutoContentAccessingProxy returns a proxy for the receiving object
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1409224-autocontentaccessingproxy?language=objc for details.
func (x gen_NSObject) AutoContentAccessingProxy() objc.Object {
	ret := C.NSObject_inst_AutoContentAccessingProxy(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// AttributeKeys an array of NSString objects containing the names of immutable values that instances of the receiver's class contain.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1415656-attributekeys?language=objc for details.
func (x gen_NSObject) AttributeKeys() NSArray {
	ret := C.NSObject_inst_AttributeKeys(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}

// ToManyRelationshipKeys an array containing the keys for the to-many relationship properties of the receiver.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1415662-tomanyrelationshipkeys?language=objc for details.
func (x gen_NSObject) ToManyRelationshipKeys() NSArray {
	ret := C.NSObject_inst_ToManyRelationshipKeys(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}

// ToOneRelationshipKeys returns the keys for the to-one relationship properties of the receiver, if any.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1414814-toonerelationshipkeys?language=objc for details.
func (x gen_NSObject) ToOneRelationshipKeys() NSArray {
	ret := C.NSObject_inst_ToOneRelationshipKeys(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}

// ClassName returns a string containing the name of the class.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1411337-classname?language=objc for details.
func (x gen_NSObject) ClassName() NSString {
	ret := C.NSObject_inst_ClassName(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// ScriptingProperties an NSString-keyed dictionary of the receiver's scriptable properties.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1417254-scriptingproperties?language=objc for details.
func (x gen_NSObject) ScriptingProperties() NSDictionary {
	ret := C.NSObject_inst_ScriptingProperties(
		unsafe.Pointer(x.Pointer()),
	)

	return NSDictionary_FromPointer(ret)
}

// SetScriptingProperties an NSString-keyed dictionary of the receiver's scriptable properties.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1417254-scriptingproperties?language=objc for details.
func (x gen_NSObject) SetScriptingProperties(
	value NSDictionaryRef,
) {
	C.NSObject_inst_SetScriptingProperties(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// AccessibilityNotifiesWhenDestroyed returns a Boolean value that indicates whether a custom accessibility object sends a notification when its corresponding UI element is destroyed.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/1534050-accessibilitynotifieswhendestroy?language=objc for details.
func (x gen_NSObject) AccessibilityNotifiesWhenDestroyed() bool {
	ret := C.NSObject_inst_AccessibilityNotifiesWhenDestroyed(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsSelectable is undocumented.
//
// See https://developer.apple.com/documentation/objectivec/nsobject/2369549-selectable?language=objc for details.
func (x gen_NSObject) IsSelectable() bool {
	ret := C.NSObject_inst_IsSelectable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

type CALayerRef interface {
	Pointer() uintptr
	Init_AsCALayer() CALayer
}

type gen_CALayer struct {
	objc.Object
}

func CALayer_FromPointer(ptr unsafe.Pointer) CALayer {
	return CALayer{gen_CALayer{
		objc.Object_FromPointer(ptr),
	}}
}

func CALayer_FromRef(ref objc.Ref) CALayer {
	return CALayer_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// ActionForKey returns the action object assigned to the specified key.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410844-actionforkey?language=objc for details.
func (x gen_CALayer) ActionForKey(
	event NSStringRef,
) objc.Object {
	ret := C.CALayer_inst_ActionForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return objc.Object_FromPointer(ret)
}

// AddSublayer appends the layer to the layer’s list of sublayers.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410833-addsublayer?language=objc for details.
func (x gen_CALayer) AddSublayer(
	layer CALayerRef,
) {
	C.CALayer_inst_AddSublayer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(layer),
	)

	return
}

// AnimationKeys returns an array of strings that identify the animations currently attached to the layer.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410937-animationkeys?language=objc for details.
func (x gen_CALayer) AnimationKeys() NSArray {
	ret := C.CALayer_inst_AnimationKeys(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}

// ContentsAreFlipped returns a Boolean indicating whether the layer content is implicitly flipped when rendered.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410777-contentsareflipped?language=objc for details.
func (x gen_CALayer) ContentsAreFlipped() bool {
	ret := C.CALayer_inst_ContentsAreFlipped(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// ConvertRectFromLayer converts the rectangle from the specified layer’s coordinate system to the receiver’s coordinate system.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410948-convertrect?language=objc for details.
func (x gen_CALayer) ConvertRectFromLayer(
	r NSRect,
	l CALayerRef,
) NSRect {
	ret := C.CALayer_inst_ConvertRectFromLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&r)),
		objc.RefPointer(l),
	)

	return *(*NSRect)(unsafe.Pointer(&ret))
}

// ConvertRectToLayer converts the rectangle from the receiver’s coordinate system to the specified layer’s coordinate system.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410742-convertrect?language=objc for details.
func (x gen_CALayer) ConvertRectToLayer(
	r NSRect,
	l CALayerRef,
) NSRect {
	ret := C.CALayer_inst_ConvertRectToLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&r)),
		objc.RefPointer(l),
	)

	return *(*NSRect)(unsafe.Pointer(&ret))
}

// Display reloads the content of this layer.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410926-display?language=objc for details.
func (x gen_CALayer) Display() {
	C.CALayer_inst_Display(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// DisplayIfNeeded initiates the update process for a layer if it is currently marked as needing an update.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410813-displayifneeded?language=objc for details.
func (x gen_CALayer) DisplayIfNeeded() {
	C.CALayer_inst_DisplayIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Init returns an initialized CALayer object.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410835-init?language=objc for details.
func (x gen_CALayer) Init() CALayer {
	ret := C.CALayer_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return CALayer_FromPointer(ret)
}

// Init_AsCALayer is a typed version of Init.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410835-init?language=objc for details.
func (x gen_CALayer) Init_AsCALayer() CALayer {
	ret := C.CALayer_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return CALayer_FromPointer(ret)
}

// InitWithLayer override to copy or initialize custom fields of the specified layer.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc for details.
func (x gen_CALayer) InitWithLayer(
	layer objc.Ref,
) CALayer {
	ret := C.CALayer_inst_InitWithLayer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(layer),
	)

	return CALayer_FromPointer(ret)
}

// InsertSublayerAbove inserts the specified sublayer above a different sublayer that already belongs to the receiver.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410798-insertsublayer?language=objc for details.
func (x gen_CALayer) InsertSublayerAbove(
	layer CALayerRef,
	sibling CALayerRef,
) {
	C.CALayer_inst_InsertSublayerAbove(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(layer),
		objc.RefPointer(sibling),
	)

	return
}

// InsertSublayerAtIndex inserts the specified layer into the receiver’s list of sublayers at the specified index.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410944-insertsublayer?language=objc for details.
func (x gen_CALayer) InsertSublayerAtIndex(
	layer CALayerRef,
	idx int32,
) {
	C.CALayer_inst_InsertSublayerAtIndex(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(layer),
		C.int(idx),
	)

	return
}

// InsertSublayerBelow inserts the specified sublayer below a different sublayer that already belongs to the receiver.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410840-insertsublayer?language=objc for details.
func (x gen_CALayer) InsertSublayerBelow(
	layer CALayerRef,
	sibling CALayerRef,
) {
	C.CALayer_inst_InsertSublayerBelow(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(layer),
		objc.RefPointer(sibling),
	)

	return
}

// LayoutIfNeeded recalculate the receiver’s layout, if required.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410873-layoutifneeded?language=objc for details.
func (x gen_CALayer) LayoutIfNeeded() {
	C.CALayer_inst_LayoutIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// LayoutSublayers tells the layer to update its layout.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410935-layoutsublayers?language=objc for details.
func (x gen_CALayer) LayoutSublayers() {
	C.CALayer_inst_LayoutSublayers(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ModelLayer returns the model layer object associated with the receiver, if any.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc for details.
func (x gen_CALayer) ModelLayer() CALayer {
	ret := C.CALayer_inst_ModelLayer(
		unsafe.Pointer(x.Pointer()),
	)

	return CALayer_FromPointer(ret)
}

// NeedsDisplay returns a Boolean indicating whether the layer has been marked as needing an update.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410958-needsdisplay?language=objc for details.
func (x gen_CALayer) NeedsDisplay() bool {
	ret := C.CALayer_inst_NeedsDisplay(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// NeedsLayout returns a Boolean indicating whether the layer has been marked as needing a layout update.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410956-needslayout?language=objc for details.
func (x gen_CALayer) NeedsLayout() bool {
	ret := C.CALayer_inst_NeedsLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// PreferredFrameSize returns the preferred size of the layer in the coordinate space of its superlayer.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410980-preferredframesize?language=objc for details.
func (x gen_CALayer) PreferredFrameSize() NSSize {
	ret := C.CALayer_inst_PreferredFrameSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSSize)(unsafe.Pointer(&ret))
}

// PresentationLayer returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc for details.
func (x gen_CALayer) PresentationLayer() CALayer {
	ret := C.CALayer_inst_PresentationLayer(
		unsafe.Pointer(x.Pointer()),
	)

	return CALayer_FromPointer(ret)
}

// RemoveAllAnimations remove all animations attached to the layer.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410810-removeallanimations?language=objc for details.
func (x gen_CALayer) RemoveAllAnimations() {
	C.CALayer_inst_RemoveAllAnimations(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// RemoveAnimationForKey remove the animation object with the specified key.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410939-removeanimationforkey?language=objc for details.
func (x gen_CALayer) RemoveAnimationForKey(
	key NSStringRef,
) {
	C.CALayer_inst_RemoveAnimationForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)

	return
}

// RemoveFromSuperlayer detaches the layer from its parent layer.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410767-removefromsuperlayer?language=objc for details.
func (x gen_CALayer) RemoveFromSuperlayer() {
	C.CALayer_inst_RemoveFromSuperlayer(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ReplaceSublayerWith replaces the specified sublayer with a different layer object.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410820-replacesublayer?language=objc for details.
func (x gen_CALayer) ReplaceSublayerWith(
	oldLayer CALayerRef,
	newLayer CALayerRef,
) {
	C.CALayer_inst_ReplaceSublayerWith(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(oldLayer),
		objc.RefPointer(newLayer),
	)

	return
}

// ResizeSublayersWithOldSize informs the receiver’s sublayers that the receiver’s size has changed.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410929-resizesublayerswitholdsize?language=objc for details.
func (x gen_CALayer) ResizeSublayersWithOldSize(
	size NSSize,
) {
	C.CALayer_inst_ResizeSublayersWithOldSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return
}

// ResizeWithOldSuperlayerSize informs the receiver that the size of its superlayer changed.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410894-resizewitholdsuperlayersize?language=objc for details.
func (x gen_CALayer) ResizeWithOldSuperlayerSize(
	size NSSize,
) {
	C.CALayer_inst_ResizeWithOldSuperlayerSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return
}

// ScrollRectToVisible initiates a scroll in the layer’s closest ancestor scroll layer so that the specified rectangle becomes visible.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1522139-scrollrecttovisible?language=objc for details.
func (x gen_CALayer) ScrollRectToVisible(
	r NSRect,
) {
	C.CALayer_inst_ScrollRectToVisible(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&r)),
	)

	return
}

// SetNeedsDisplay marks the layer’s contents as needing to be updated.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410855-setneedsdisplay?language=objc for details.
func (x gen_CALayer) SetNeedsDisplay() {
	C.CALayer_inst_SetNeedsDisplay(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// SetNeedsDisplayInRect marks the region within the specified rectangle as needing to be updated.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410800-setneedsdisplayinrect?language=objc for details.
func (x gen_CALayer) SetNeedsDisplayInRect(
	r NSRect,
) {
	C.CALayer_inst_SetNeedsDisplayInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&r)),
	)

	return
}

// SetNeedsLayout invalidates the layer’s layout and marks it as needing an update.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410946-setneedslayout?language=objc for details.
func (x gen_CALayer) SetNeedsLayout() {
	C.CALayer_inst_SetNeedsLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// ShouldArchiveValueForKey returns a Boolean indicating whether the value of the specified key should be archived.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410753-shouldarchivevalueforkey?language=objc for details.
func (x gen_CALayer) ShouldArchiveValueForKey(
	key NSStringRef,
) bool {
	ret := C.CALayer_inst_ShouldArchiveValueForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)

	return convertObjCBoolToGo(ret)
}

// Delegate returns the layer’s delegate object.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410984-delegate?language=objc for details.
func (x gen_CALayer) Delegate() objc.Object {
	ret := C.CALayer_inst_Delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetDelegate returns the layer’s delegate object.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410984-delegate?language=objc for details.
func (x gen_CALayer) SetDelegate(
	value objc.Ref,
) {
	C.CALayer_inst_SetDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Contents an object that provides the contents of the layer. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410773-contents?language=objc for details.
func (x gen_CALayer) Contents() objc.Object {
	ret := C.CALayer_inst_Contents(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetContents an object that provides the contents of the layer. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410773-contents?language=objc for details.
func (x gen_CALayer) SetContents(
	value objc.Ref,
) {
	C.CALayer_inst_SetContents(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// ContentsRect returns the rectangle, in the unit coordinate space, that defines the portion of the layer’s contents that should be used. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410866-contentsrect?language=objc for details.
func (x gen_CALayer) ContentsRect() NSRect {
	ret := C.CALayer_inst_ContentsRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSRect)(unsafe.Pointer(&ret))
}

// SetContentsRect returns the rectangle, in the unit coordinate space, that defines the portion of the layer’s contents that should be used. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410866-contentsrect?language=objc for details.
func (x gen_CALayer) SetContentsRect(
	value NSRect,
) {
	C.CALayer_inst_SetContentsRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return
}

// ContentsCenter returns the rectangle that defines how the layer contents are scaled if the layer’s contents are resized. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410740-contentscenter?language=objc for details.
func (x gen_CALayer) ContentsCenter() NSRect {
	ret := C.CALayer_inst_ContentsCenter(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSRect)(unsafe.Pointer(&ret))
}

// SetContentsCenter returns the rectangle that defines how the layer contents are scaled if the layer’s contents are resized. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410740-contentscenter?language=objc for details.
func (x gen_CALayer) SetContentsCenter(
	value NSRect,
) {
	C.CALayer_inst_SetContentsCenter(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return
}

// IsHidden returns a Boolean indicating whether the layer is displayed. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410838-hidden?language=objc for details.
func (x gen_CALayer) IsHidden() bool {
	ret := C.CALayer_inst_IsHidden(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetHidden returns a Boolean indicating whether the layer is displayed. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410838-hidden?language=objc for details.
func (x gen_CALayer) SetHidden(
	value bool,
) {
	C.CALayer_inst_SetHidden(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// MasksToBounds returns a Boolean indicating whether sublayers are clipped to the layer’s bounds. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410896-maskstobounds?language=objc for details.
func (x gen_CALayer) MasksToBounds() bool {
	ret := C.CALayer_inst_MasksToBounds(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetMasksToBounds returns a Boolean indicating whether sublayers are clipped to the layer’s bounds. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410896-maskstobounds?language=objc for details.
func (x gen_CALayer) SetMasksToBounds(
	value bool,
) {
	C.CALayer_inst_SetMasksToBounds(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// Mask an optional layer whose alpha channel is used to mask the layer’s content.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410861-mask?language=objc for details.
func (x gen_CALayer) Mask() CALayer {
	ret := C.CALayer_inst_Mask(
		unsafe.Pointer(x.Pointer()),
	)

	return CALayer_FromPointer(ret)
}

// SetMask an optional layer whose alpha channel is used to mask the layer’s content.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410861-mask?language=objc for details.
func (x gen_CALayer) SetMask(
	value CALayerRef,
) {
	C.CALayer_inst_SetMask(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// IsDoubleSided returns a Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410924-doublesided?language=objc for details.
func (x gen_CALayer) IsDoubleSided() bool {
	ret := C.CALayer_inst_IsDoubleSided(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetDoubleSided returns a Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410924-doublesided?language=objc for details.
func (x gen_CALayer) SetDoubleSided(
	value bool,
) {
	C.CALayer_inst_SetDoubleSided(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// CornerRadius returns the radius to use when drawing rounded corners for the layer’s background. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410818-cornerradius?language=objc for details.
func (x gen_CALayer) CornerRadius() CGFloat {
	ret := C.CALayer_inst_CornerRadius(
		unsafe.Pointer(x.Pointer()),
	)

	return CGFloat(ret)
}

// SetCornerRadius returns the radius to use when drawing rounded corners for the layer’s background. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410818-cornerradius?language=objc for details.
func (x gen_CALayer) SetCornerRadius(
	value CGFloat,
) {
	C.CALayer_inst_SetCornerRadius(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// BorderWidth returns the width of the layer’s border. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410917-borderwidth?language=objc for details.
func (x gen_CALayer) BorderWidth() CGFloat {
	ret := C.CALayer_inst_BorderWidth(
		unsafe.Pointer(x.Pointer()),
	)

	return CGFloat(ret)
}

// SetBorderWidth returns the width of the layer’s border. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410917-borderwidth?language=objc for details.
func (x gen_CALayer) SetBorderWidth(
	value CGFloat,
) {
	C.CALayer_inst_SetBorderWidth(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// ShadowRadius returns the blur radius (in points) used to render the layer’s shadow. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410819-shadowradius?language=objc for details.
func (x gen_CALayer) ShadowRadius() CGFloat {
	ret := C.CALayer_inst_ShadowRadius(
		unsafe.Pointer(x.Pointer()),
	)

	return CGFloat(ret)
}

// SetShadowRadius returns the blur radius (in points) used to render the layer’s shadow. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410819-shadowradius?language=objc for details.
func (x gen_CALayer) SetShadowRadius(
	value CGFloat,
) {
	C.CALayer_inst_SetShadowRadius(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// ShadowOffset returns the offset (in points) of the layer’s shadow. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410970-shadowoffset?language=objc for details.
func (x gen_CALayer) ShadowOffset() NSSize {
	ret := C.CALayer_inst_ShadowOffset(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSSize)(unsafe.Pointer(&ret))
}

// SetShadowOffset returns the offset (in points) of the layer’s shadow. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410970-shadowoffset?language=objc for details.
func (x gen_CALayer) SetShadowOffset(
	value NSSize,
) {
	C.CALayer_inst_SetShadowOffset(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return
}

// Style an optional dictionary used to store property values that aren't explicitly defined by the layer.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410875-style?language=objc for details.
func (x gen_CALayer) Style() NSDictionary {
	ret := C.CALayer_inst_Style(
		unsafe.Pointer(x.Pointer()),
	)

	return NSDictionary_FromPointer(ret)
}

// SetStyle an optional dictionary used to store property values that aren't explicitly defined by the layer.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410875-style?language=objc for details.
func (x gen_CALayer) SetStyle(
	value NSDictionaryRef,
) {
	C.CALayer_inst_SetStyle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// AllowsEdgeAntialiasing returns a Boolean indicating whether the layer is allowed to perform edge antialiasing.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1621285-allowsedgeantialiasing?language=objc for details.
func (x gen_CALayer) AllowsEdgeAntialiasing() bool {
	ret := C.CALayer_inst_AllowsEdgeAntialiasing(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsEdgeAntialiasing returns a Boolean indicating whether the layer is allowed to perform edge antialiasing.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1621285-allowsedgeantialiasing?language=objc for details.
func (x gen_CALayer) SetAllowsEdgeAntialiasing(
	value bool,
) {
	C.CALayer_inst_SetAllowsEdgeAntialiasing(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AllowsGroupOpacity returns a Boolean indicating whether the layer is allowed to composite itself as a group separate from its parent.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1621277-allowsgroupopacity?language=objc for details.
func (x gen_CALayer) AllowsGroupOpacity() bool {
	ret := C.CALayer_inst_AllowsGroupOpacity(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsGroupOpacity returns a Boolean indicating whether the layer is allowed to composite itself as a group separate from its parent.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1621277-allowsgroupopacity?language=objc for details.
func (x gen_CALayer) SetAllowsGroupOpacity(
	value bool,
) {
	C.CALayer_inst_SetAllowsGroupOpacity(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// Filters an array of Core Image filters to apply to the contents of the layer and its sublayers. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410901-filters?language=objc for details.
func (x gen_CALayer) Filters() NSArray {
	ret := C.CALayer_inst_Filters(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}

// SetFilters an array of Core Image filters to apply to the contents of the layer and its sublayers. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410901-filters?language=objc for details.
func (x gen_CALayer) SetFilters(
	value NSArrayRef,
) {
	C.CALayer_inst_SetFilters(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// CompositingFilter returns a CoreImage filter used to composite the layer and the content behind it. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410748-compositingfilter?language=objc for details.
func (x gen_CALayer) CompositingFilter() objc.Object {
	ret := C.CALayer_inst_CompositingFilter(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetCompositingFilter returns a CoreImage filter used to composite the layer and the content behind it. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410748-compositingfilter?language=objc for details.
func (x gen_CALayer) SetCompositingFilter(
	value objc.Ref,
) {
	C.CALayer_inst_SetCompositingFilter(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// BackgroundFilters an array of Core Image filters to apply to the content immediately behind the layer. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410827-backgroundfilters?language=objc for details.
func (x gen_CALayer) BackgroundFilters() NSArray {
	ret := C.CALayer_inst_BackgroundFilters(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}

// SetBackgroundFilters an array of Core Image filters to apply to the content immediately behind the layer. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410827-backgroundfilters?language=objc for details.
func (x gen_CALayer) SetBackgroundFilters(
	value NSArrayRef,
) {
	C.CALayer_inst_SetBackgroundFilters(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// IsOpaque returns a Boolean value indicating whether the layer contains completely opaque content.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410763-opaque?language=objc for details.
func (x gen_CALayer) IsOpaque() bool {
	ret := C.CALayer_inst_IsOpaque(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetOpaque returns a Boolean value indicating whether the layer contains completely opaque content.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410763-opaque?language=objc for details.
func (x gen_CALayer) SetOpaque(
	value bool,
) {
	C.CALayer_inst_SetOpaque(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsGeometryFlipped returns a Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410960-geometryflipped?language=objc for details.
func (x gen_CALayer) IsGeometryFlipped() bool {
	ret := C.CALayer_inst_IsGeometryFlipped(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetGeometryFlipped returns a Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410960-geometryflipped?language=objc for details.
func (x gen_CALayer) SetGeometryFlipped(
	value bool,
) {
	C.CALayer_inst_SetGeometryFlipped(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// DrawsAsynchronously returns a Boolean indicating whether drawing commands are deferred and processed asynchronously in a background thread.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410974-drawsasynchronously?language=objc for details.
func (x gen_CALayer) DrawsAsynchronously() bool {
	ret := C.CALayer_inst_DrawsAsynchronously(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetDrawsAsynchronously returns a Boolean indicating whether drawing commands are deferred and processed asynchronously in a background thread.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410974-drawsasynchronously?language=objc for details.
func (x gen_CALayer) SetDrawsAsynchronously(
	value bool,
) {
	C.CALayer_inst_SetDrawsAsynchronously(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// ShouldRasterize returns a Boolean that indicates whether the layer is rendered as a bitmap before compositing. Animatable
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410905-shouldrasterize?language=objc for details.
func (x gen_CALayer) ShouldRasterize() bool {
	ret := C.CALayer_inst_ShouldRasterize(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetShouldRasterize returns a Boolean that indicates whether the layer is rendered as a bitmap before compositing. Animatable
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410905-shouldrasterize?language=objc for details.
func (x gen_CALayer) SetShouldRasterize(
	value bool,
) {
	C.CALayer_inst_SetShouldRasterize(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// RasterizationScale returns the scale at which to rasterize content, relative to the coordinate space of the layer. Animatable
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410801-rasterizationscale?language=objc for details.
func (x gen_CALayer) RasterizationScale() CGFloat {
	ret := C.CALayer_inst_RasterizationScale(
		unsafe.Pointer(x.Pointer()),
	)

	return CGFloat(ret)
}

// SetRasterizationScale returns the scale at which to rasterize content, relative to the coordinate space of the layer. Animatable
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410801-rasterizationscale?language=objc for details.
func (x gen_CALayer) SetRasterizationScale(
	value CGFloat,
) {
	C.CALayer_inst_SetRasterizationScale(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// Frame returns the layer’s frame rectangle.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410779-frame?language=objc for details.
func (x gen_CALayer) Frame() NSRect {
	ret := C.CALayer_inst_Frame(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSRect)(unsafe.Pointer(&ret))
}

// SetFrame returns the layer’s frame rectangle.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410779-frame?language=objc for details.
func (x gen_CALayer) SetFrame(
	value NSRect,
) {
	C.CALayer_inst_SetFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return
}

// Bounds returns the layer’s bounds rectangle. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410915-bounds?language=objc for details.
func (x gen_CALayer) Bounds() NSRect {
	ret := C.CALayer_inst_Bounds(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSRect)(unsafe.Pointer(&ret))
}

// SetBounds returns the layer’s bounds rectangle. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410915-bounds?language=objc for details.
func (x gen_CALayer) SetBounds(
	value NSRect,
) {
	C.CALayer_inst_SetBounds(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return
}

// ZPosition returns the layer’s position on the z axis. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410884-zposition?language=objc for details.
func (x gen_CALayer) ZPosition() CGFloat {
	ret := C.CALayer_inst_ZPosition(
		unsafe.Pointer(x.Pointer()),
	)

	return CGFloat(ret)
}

// SetZPosition returns the layer’s position on the z axis. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410884-zposition?language=objc for details.
func (x gen_CALayer) SetZPosition(
	value CGFloat,
) {
	C.CALayer_inst_SetZPosition(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// AnchorPointZ returns the anchor point for the layer’s position along the z axis. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410796-anchorpointz?language=objc for details.
func (x gen_CALayer) AnchorPointZ() CGFloat {
	ret := C.CALayer_inst_AnchorPointZ(
		unsafe.Pointer(x.Pointer()),
	)

	return CGFloat(ret)
}

// SetAnchorPointZ returns the anchor point for the layer’s position along the z axis. Animatable.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410796-anchorpointz?language=objc for details.
func (x gen_CALayer) SetAnchorPointZ(
	value CGFloat,
) {
	C.CALayer_inst_SetAnchorPointZ(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// ContentsScale returns the scale factor applied to the layer.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410746-contentsscale?language=objc for details.
func (x gen_CALayer) ContentsScale() CGFloat {
	ret := C.CALayer_inst_ContentsScale(
		unsafe.Pointer(x.Pointer()),
	)

	return CGFloat(ret)
}

// SetContentsScale returns the scale factor applied to the layer.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410746-contentsscale?language=objc for details.
func (x gen_CALayer) SetContentsScale(
	value CGFloat,
) {
	C.CALayer_inst_SetContentsScale(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// Sublayers an array containing the layer’s sublayers.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410802-sublayers?language=objc for details.
func (x gen_CALayer) Sublayers() NSArray {
	ret := C.CALayer_inst_Sublayers(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}

// SetSublayers an array containing the layer’s sublayers.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410802-sublayers?language=objc for details.
func (x gen_CALayer) SetSublayers(
	value NSArrayRef,
) {
	C.CALayer_inst_SetSublayers(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Superlayer returns the superlayer of the layer.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410761-superlayer?language=objc for details.
func (x gen_CALayer) Superlayer() CALayer {
	ret := C.CALayer_inst_Superlayer(
		unsafe.Pointer(x.Pointer()),
	)

	return CALayer_FromPointer(ret)
}

// NeedsDisplayOnBoundsChange returns a Boolean indicating whether the layer contents must be updated when its bounds rectangle changes.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410923-needsdisplayonboundschange?language=objc for details.
func (x gen_CALayer) NeedsDisplayOnBoundsChange() bool {
	ret := C.CALayer_inst_NeedsDisplayOnBoundsChange(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetNeedsDisplayOnBoundsChange returns a Boolean indicating whether the layer contents must be updated when its bounds rectangle changes.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410923-needsdisplayonboundschange?language=objc for details.
func (x gen_CALayer) SetNeedsDisplayOnBoundsChange(
	value bool,
) {
	C.CALayer_inst_SetNeedsDisplayOnBoundsChange(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// LayoutManager returns the object responsible for laying out the layer’s sublayers.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410749-layoutmanager?language=objc for details.
func (x gen_CALayer) LayoutManager() objc.Object {
	ret := C.CALayer_inst_LayoutManager(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetLayoutManager returns the object responsible for laying out the layer’s sublayers.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410749-layoutmanager?language=objc for details.
func (x gen_CALayer) SetLayoutManager(
	value objc.Ref,
) {
	C.CALayer_inst_SetLayoutManager(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Constraints returns the constraints used to position current layer’s sublayers.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1521906-constraints?language=objc for details.
func (x gen_CALayer) Constraints() NSArray {
	ret := C.CALayer_inst_Constraints(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}

// SetConstraints returns the constraints used to position current layer’s sublayers.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1521906-constraints?language=objc for details.
func (x gen_CALayer) SetConstraints(
	value NSArrayRef,
) {
	C.CALayer_inst_SetConstraints(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// Actions returns a dictionary containing layer actions.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410789-actions?language=objc for details.
func (x gen_CALayer) Actions() NSDictionary {
	ret := C.CALayer_inst_Actions(
		unsafe.Pointer(x.Pointer()),
	)

	return NSDictionary_FromPointer(ret)
}

// SetActions returns a dictionary containing layer actions.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410789-actions?language=objc for details.
func (x gen_CALayer) SetActions(
	value NSDictionaryRef,
) {
	C.CALayer_inst_SetActions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// VisibleRect returns the visible region of the layer in its own coordinate space.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1521892-visiblerect?language=objc for details.
func (x gen_CALayer) VisibleRect() NSRect {
	ret := C.CALayer_inst_VisibleRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSRect)(unsafe.Pointer(&ret))
}

// Name returns the name of the receiver.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410879-name?language=objc for details.
func (x gen_CALayer) Name() NSString {
	ret := C.CALayer_inst_Name(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// SetName returns the name of the receiver.
//
// See https://developer.apple.com/documentation/quartzcore/calayer/1410879-name?language=objc for details.
func (x gen_CALayer) SetName(
	value NSStringRef,
) {
	C.CALayer_inst_SetName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

type NSArrayRef interface {
	Pointer() uintptr
	Init_AsNSArray() NSArray
}

type gen_NSArray struct {
	objc.Object
}

func NSArray_FromPointer(ptr unsafe.Pointer) NSArray {
	return NSArray{gen_NSArray{
		objc.Object_FromPointer(ptr),
	}}
}

func NSArray_FromRef(ref objc.Ref) NSArray {
	return NSArray_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// ArrayByAddingObjectsFromArray returns a new array that is a copy of the receiving array with the objects contained in another array added to the end.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1412087-arraybyaddingobjectsfromarray?language=objc for details.
func (x gen_NSArray) ArrayByAddingObjectsFromArray(
	otherArray NSArrayRef,
) NSArray {
	ret := C.NSArray_inst_ArrayByAddingObjectsFromArray(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(otherArray),
	)

	return NSArray_FromPointer(ret)
}

// ComponentsJoinedByString constructs and returns an NSString object that is the result of interposing a given separator between the elements of the array.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1412075-componentsjoinedbystring?language=objc for details.
func (x gen_NSArray) ComponentsJoinedByString(
	separator NSStringRef,
) NSString {
	ret := C.NSArray_inst_ComponentsJoinedByString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(separator),
	)

	return NSString_FromPointer(ret)
}

// DescriptionWithLocale returns a string that represents the contents of the array, formatted as a property list.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1412374-descriptionwithlocale?language=objc for details.
func (x gen_NSArray) DescriptionWithLocale(
	locale objc.Ref,
) NSString {
	ret := C.NSArray_inst_DescriptionWithLocale(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(locale),
	)

	return NSString_FromPointer(ret)
}

// DescriptionWithLocaleIndent returns a string that represents the contents of the array, formatted as a property list.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1416257-descriptionwithlocale?language=objc for details.
func (x gen_NSArray) DescriptionWithLocaleIndent(
	locale objc.Ref,
	level NSUInteger,
) NSString {
	ret := C.NSArray_inst_DescriptionWithLocaleIndent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(locale),
		C.ulong(level),
	)

	return NSString_FromPointer(ret)
}

// Init initializes a newly allocated array.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1414315-init?language=objc for details.
func (x gen_NSArray) Init() NSArray {
	ret := C.NSArray_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}

// Init_AsNSArray is a typed version of Init.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1414315-init?language=objc for details.
func (x gen_NSArray) Init_AsNSArray() NSArray {
	ret := C.NSArray_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}

// InitWithArray initializes a newly allocated array by placing in it the objects contained in a given array.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1412169-initwitharray?language=objc for details.
func (x gen_NSArray) InitWithArray(
	array NSArrayRef,
) NSArray {
	ret := C.NSArray_inst_InitWithArray(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(array),
	)

	return NSArray_FromPointer(ret)
}

// InitWithArrayCopyItems initializes a newly allocated array using anArray as the source of data objects for the array.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1408557-initwitharray?language=objc for details.
func (x gen_NSArray) InitWithArrayCopyItems(
	array NSArrayRef,
	flag bool,
) NSArray {
	ret := C.NSArray_inst_InitWithArrayCopyItems(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(array),
		convertToObjCBool(flag),
	)

	return NSArray_FromPointer(ret)
}

// IsEqualToArray compares the receiving array to another array.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1411770-isequaltoarray?language=objc for details.
func (x gen_NSArray) IsEqualToArray(
	otherArray NSArrayRef,
) bool {
	ret := C.NSArray_inst_IsEqualToArray(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(otherArray),
	)

	return convertObjCBoolToGo(ret)
}

// MakeObjectsPerformSelector sends to each object in the array the message identified by a given selector, starting with the first object and continuing through the array to the last object.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1460115-makeobjectsperformselector?language=objc for details.
func (x gen_NSArray) MakeObjectsPerformSelector(
	aSelector objc.Selector,
) {
	C.NSArray_inst_MakeObjectsPerformSelector(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
	)

	return
}

// MakeObjectsPerformSelectorWithObject sends the aSelector message to each object in the array, starting with the first object and continuing through the array to the last object.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1460107-makeobjectsperformselector?language=objc for details.
func (x gen_NSArray) MakeObjectsPerformSelectorWithObject(
	aSelector objc.Selector,
	argument objc.Ref,
) {
	C.NSArray_inst_MakeObjectsPerformSelectorWithObject(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
		objc.RefPointer(argument),
	)

	return
}

// PathsMatchingExtensions returns an array containing all the pathname elements in the receiving array that have filename extensions from a given array.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1418275-pathsmatchingextensions?language=objc for details.
func (x gen_NSArray) PathsMatchingExtensions(
	filterTypes NSArrayRef,
) NSArray {
	ret := C.NSArray_inst_PathsMatchingExtensions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(filterTypes),
	)

	return NSArray_FromPointer(ret)
}

// RemoveObserverForKeyPath raises an exception.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1414976-removeobserver?language=objc for details.
func (x gen_NSArray) RemoveObserverForKeyPath(
	observer NSObjectRef,
	keyPath NSStringRef,
) {
	C.NSArray_inst_RemoveObserverForKeyPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(observer),
		objc.RefPointer(keyPath),
	)

	return
}

// RemoveObserverForKeyPathContext raises an exception.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1418441-removeobserver?language=objc for details.
func (x gen_NSArray) RemoveObserverForKeyPathContext(
	observer NSObjectRef,
	keyPath NSStringRef,
	context unsafe.Pointer,
) {
	C.NSArray_inst_RemoveObserverForKeyPathContext(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(observer),
		objc.RefPointer(keyPath),
		context,
	)

	return
}

// SetValueForKey invokes setValue:forKey: on each of the array's items using the specified value and key.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1408301-setvalue?language=objc for details.
func (x gen_NSArray) SetValueForKey(
	value objc.Ref,
	key NSStringRef,
) {
	C.NSArray_inst_SetValueForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
		objc.RefPointer(key),
	)

	return
}

// ShuffledArray returns a new array that lists this array’s elements in a random order.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1640855-shuffledarray?language=objc for details.
func (x gen_NSArray) ShuffledArray() NSArray {
	ret := C.NSArray_inst_ShuffledArray(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}

// SortedArrayUsingDescriptors returns a copy of the receiving array sorted as specified by a given array of sort descriptors.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1415069-sortedarrayusingdescriptors?language=objc for details.
func (x gen_NSArray) SortedArrayUsingDescriptors(
	sortDescriptors NSArrayRef,
) NSArray {
	ret := C.NSArray_inst_SortedArrayUsingDescriptors(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sortDescriptors),
	)

	return NSArray_FromPointer(ret)
}

// SortedArrayUsingSelector returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given selector.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1410025-sortedarrayusingselector?language=objc for details.
func (x gen_NSArray) SortedArrayUsingSelector(
	comparator objc.Selector,
) NSArray {
	ret := C.NSArray_inst_SortedArrayUsingSelector(
		unsafe.Pointer(x.Pointer()),
		comparator.SelectorAddress(),
	)

	return NSArray_FromPointer(ret)
}

// ValueForKey returns an array containing the results of invoking valueForKey: using key on each of the array's objects.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1412219-valueforkey?language=objc for details.
func (x gen_NSArray) ValueForKey(
	key NSStringRef,
) objc.Object {
	ret := C.NSArray_inst_ValueForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)

	return objc.Object_FromPointer(ret)
}

// Count returns the number of objects in the array.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1409982-count?language=objc for details.
func (x gen_NSArray) Count() NSUInteger {
	ret := C.NSArray_inst_Count(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)
}

// SortedArrayHint analyzes the array and returns a “hint” that speeds the sorting of the array when the hint is supplied to sortedArrayUsingFunction:context:hint:.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1413063-sortedarrayhint?language=objc for details.
func (x gen_NSArray) SortedArrayHint() NSData {
	ret := C.NSArray_inst_SortedArrayHint(
		unsafe.Pointer(x.Pointer()),
	)

	return NSData_FromPointer(ret)
}

// Description returns a string that represents the contents of the array, formatted as a property list.
//
// See https://developer.apple.com/documentation/foundation/nsarray/1413042-description?language=objc for details.
func (x gen_NSArray) Description() NSString {
	ret := C.NSArray_inst_Description(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

type NSAttributedStringRef interface {
	Pointer() uintptr
	Init_AsNSAttributedString() NSAttributedString
}

type gen_NSAttributedString struct {
	objc.Object
}

func NSAttributedString_FromPointer(ptr unsafe.Pointer) NSAttributedString {
	return NSAttributedString{gen_NSAttributedString{
		objc.Object_FromPointer(ptr),
	}}
}

func NSAttributedString_FromRef(ref objc.Ref) NSAttributedString {
	return NSAttributedString_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// AttributedStringByInflectingString is undocumented.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/3746871-attributedstringbyinflectingstri?language=objc for details.
func (x gen_NSAttributedString) AttributedStringByInflectingString() NSAttributedString {
	ret := C.NSAttributedString_inst_AttributedStringByInflectingString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSAttributedString_FromPointer(ret)
}

// DrawInRect draws the attributed string inside the specified bounding rectangle in the current graphics context.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1531631-drawinrect?language=objc for details.
func (x gen_NSAttributedString) DrawInRect(
	rect NSRect,
) {
	C.NSAttributedString_inst_DrawInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return
}

// InitWithAttributedString creates an attributed string with the characters and attributes of the specified attributed string.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1415342-initwithattributedstring?language=objc for details.
func (x gen_NSAttributedString) InitWithAttributedString(
	attrStr NSAttributedStringRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_InitWithAttributedString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(attrStr),
	)

	return NSAttributedString_FromPointer(ret)
}

// InitWithDocFormatDocumentAttributes creates an attributed string from Microsoft Word format data in the specified data object.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1534329-initwithdocformat?language=objc for details.
func (x gen_NSAttributedString) InitWithDocFormatDocumentAttributes(
	data NSDataRef,
	dict NSDictionaryRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_InitWithDocFormatDocumentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(dict),
	)

	return NSAttributedString_FromPointer(ret)
}

// InitWithHTMLBaseURLDocumentAttributes creates an attributed string from the HTML in the specified data object and base URL.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1524624-initwithhtml?language=objc for details.
func (x gen_NSAttributedString) InitWithHTMLBaseURLDocumentAttributes(
	data NSDataRef,
	base NSURLRef,
	dict NSDictionaryRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_InitWithHTMLBaseURLDocumentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(base),
		objc.RefPointer(dict),
	)

	return NSAttributedString_FromPointer(ret)
}

// InitWithHTMLDocumentAttributes creates an attributed string from the HTML in the specified data object.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1525953-initwithhtml?language=objc for details.
func (x gen_NSAttributedString) InitWithHTMLDocumentAttributes(
	data NSDataRef,
	dict NSDictionaryRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_InitWithHTMLDocumentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(dict),
	)

	return NSAttributedString_FromPointer(ret)
}

// InitWithHTMLOptionsDocumentAttributes creates an attributed string from the HTML in the specified data object.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1535412-initwithhtml?language=objc for details.
func (x gen_NSAttributedString) InitWithHTMLOptionsDocumentAttributes(
	data NSDataRef,
	options NSDictionaryRef,
	dict NSDictionaryRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_InitWithHTMLOptionsDocumentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(options),
		objc.RefPointer(dict),
	)

	return NSAttributedString_FromPointer(ret)
}

// InitWithRTFDocumentAttributes creates an attributed string by decoding the stream of RTF commands and data in the specified data object.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1532912-initwithrtf?language=objc for details.
func (x gen_NSAttributedString) InitWithRTFDocumentAttributes(
	data NSDataRef,
	dict NSDictionaryRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_InitWithRTFDocumentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(dict),
	)

	return NSAttributedString_FromPointer(ret)
}

// InitWithRTFDDocumentAttributes creates an attributed string by decoding the stream of RTFD commands and data in the specified data object.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1530987-initwithrtfd?language=objc for details.
func (x gen_NSAttributedString) InitWithRTFDDocumentAttributes(
	data NSDataRef,
	dict NSDictionaryRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_InitWithRTFDDocumentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(dict),
	)

	return NSAttributedString_FromPointer(ret)
}

// InitWithString creates an attributed string with the characters of the specified string and no attribute information.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1407481-initwithstring?language=objc for details.
func (x gen_NSAttributedString) InitWithString(
	str NSStringRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_InitWithString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return NSAttributedString_FromPointer(ret)
}

// InitWithStringAttributes creates an attributed string with the specified string and attributes.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1408136-initwithstring?language=objc for details.
func (x gen_NSAttributedString) InitWithStringAttributes(
	str NSStringRef,
	attrs NSDictionaryRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_InitWithStringAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
		objc.RefPointer(attrs),
	)

	return NSAttributedString_FromPointer(ret)
}

// IsEqualToAttributedString returns a Boolean value that indicates whether the attributed string is equal to another attributed string.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1414808-isequaltoattributedstring?language=objc for details.
func (x gen_NSAttributedString) IsEqualToAttributedString(
	other NSAttributedStringRef,
) bool {
	ret := C.NSAttributedString_inst_IsEqualToAttributedString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(other),
	)

	return convertObjCBoolToGo(ret)
}

// NextWordFromIndexForward returns the index of the first character of the word after or before the specified index.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1535305-nextwordfromindex?language=objc for details.
func (x gen_NSAttributedString) NextWordFromIndexForward(
	location NSUInteger,
	isForward bool,
) NSUInteger {
	ret := C.NSAttributedString_inst_NextWordFromIndexForward(
		unsafe.Pointer(x.Pointer()),
		C.ulong(location),
		convertToObjCBool(isForward),
	)

	return NSUInteger(ret)
}

// Size returns the size necessary to draw the string.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1528362-size?language=objc for details.
func (x gen_NSAttributedString) Size() NSSize {
	ret := C.NSAttributedString_inst_Size(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSSize)(unsafe.Pointer(&ret))
}

// Init initializes a new instance of the NSAttributedString class.
func (x gen_NSAttributedString) Init() NSAttributedString {
	ret := C.NSAttributedString_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSAttributedString_FromPointer(ret)
}

// Init_AsNSAttributedString is a typed version of Init.
func (x gen_NSAttributedString) Init_AsNSAttributedString() NSAttributedString {
	ret := C.NSAttributedString_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSAttributedString_FromPointer(ret)
}

// String returns the character contents of the attributed string as a string.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1412616-string?language=objc for details.
func (x gen_NSAttributedString) String() NSString {
	ret := C.NSAttributedString_inst_String(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// Length returns the length of the attributed string.
//
// See https://developer.apple.com/documentation/foundation/nsattributedstring/1418432-length?language=objc for details.
func (x gen_NSAttributedString) Length() NSUInteger {
	ret := C.NSAttributedString_inst_Length(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)
}

type NSDataRef interface {
	Pointer() uintptr
	Init_AsNSData() NSData
}

type gen_NSData struct {
	objc.Object
}

func NSData_FromPointer(ptr unsafe.Pointer) NSData {
	return NSData{gen_NSData{
		objc.Object_FromPointer(ptr),
	}}
}

func NSData_FromRef(ref objc.Ref) NSData {
	return NSData_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// GetBytesLength copies a number of bytes from the start of the data object into a given buffer.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1411450-getbytes?language=objc for details.
func (x gen_NSData) GetBytesLength(
	buffer unsafe.Pointer,
	length NSUInteger,
) {
	C.NSData_inst_GetBytesLength(
		unsafe.Pointer(x.Pointer()),
		buffer,
		C.ulong(length),
	)

	return
}

// InitWithBytesLength initializes a data object filled with a given number of bytes copied from a given buffer.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1412793-initwithbytes?language=objc for details.
func (x gen_NSData) InitWithBytesLength(
	bytes unsafe.Pointer,
	length NSUInteger,
) NSData {
	ret := C.NSData_inst_InitWithBytesLength(
		unsafe.Pointer(x.Pointer()),
		bytes,
		C.ulong(length),
	)

	return NSData_FromPointer(ret)
}

// InitWithBytesNoCopyLength initializes a data object filled with a given number of bytes of data from a given buffer.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1409454-initwithbytesnocopy?language=objc for details.
func (x gen_NSData) InitWithBytesNoCopyLength(
	bytes unsafe.Pointer,
	length NSUInteger,
) NSData {
	ret := C.NSData_inst_InitWithBytesNoCopyLength(
		unsafe.Pointer(x.Pointer()),
		bytes,
		C.ulong(length),
	)

	return NSData_FromPointer(ret)
}

// InitWithBytesNoCopyLengthFreeWhenDone initializes a newly allocated data object by adding the given number of bytes from the given buffer.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1416020-initwithbytesnocopy?language=objc for details.
func (x gen_NSData) InitWithBytesNoCopyLengthFreeWhenDone(
	bytes unsafe.Pointer,
	length NSUInteger,
	b bool,
) NSData {
	ret := C.NSData_inst_InitWithBytesNoCopyLengthFreeWhenDone(
		unsafe.Pointer(x.Pointer()),
		bytes,
		C.ulong(length),
		convertToObjCBool(b),
	)

	return NSData_FromPointer(ret)
}

// InitWithContentsOfFile initializes a data object with the content of the file at a given path.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1408672-initwithcontentsoffile?language=objc for details.
func (x gen_NSData) InitWithContentsOfFile(
	path NSStringRef,
) NSData {
	ret := C.NSData_inst_InitWithContentsOfFile(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)

	return NSData_FromPointer(ret)
}

// InitWithContentsOfURL initializes a data object with the data from the location specified by a given URL.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1413892-initwithcontentsofurl?language=objc for details.
func (x gen_NSData) InitWithContentsOfURL(
	url NSURLRef,
) NSData {
	ret := C.NSData_inst_InitWithContentsOfURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)

	return NSData_FromPointer(ret)
}

// InitWithData initializes a data object with the contents of another data object.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1417055-initwithdata?language=objc for details.
func (x gen_NSData) InitWithData(
	data NSDataRef,
) NSData {
	ret := C.NSData_inst_InitWithData(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
	)

	return NSData_FromPointer(ret)
}

// IsEqualToData returns a Boolean value indicating whether this data object is the same as another.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1409330-isequaltodata?language=objc for details.
func (x gen_NSData) IsEqualToData(
	other NSDataRef,
) bool {
	ret := C.NSData_inst_IsEqualToData(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(other),
	)

	return convertObjCBoolToGo(ret)
}

// WriteToFileAtomically writes the data object's bytes to the file specified by a given path.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1408033-writetofile?language=objc for details.
func (x gen_NSData) WriteToFileAtomically(
	path NSStringRef,
	useAuxiliaryFile bool,
) bool {
	ret := C.NSData_inst_WriteToFileAtomically(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(useAuxiliaryFile),
	)

	return convertObjCBoolToGo(ret)
}

// WriteToURLAtomically writes the data object's bytes to the location specified by a given URL.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1415134-writetourl?language=objc for details.
func (x gen_NSData) WriteToURLAtomically(
	url NSURLRef,
	atomically bool,
) bool {
	ret := C.NSData_inst_WriteToURLAtomically(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
		convertToObjCBool(atomically),
	)

	return convertObjCBoolToGo(ret)
}

// Init initializes a new instance of the NSData class.
func (x gen_NSData) Init() NSData {
	ret := C.NSData_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSData_FromPointer(ret)
}

// Init_AsNSData is a typed version of Init.
func (x gen_NSData) Init_AsNSData() NSData {
	ret := C.NSData_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSData_FromPointer(ret)
}

// Bytes returns a pointer to the data object's contents.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1410616-bytes?language=objc for details.
func (x gen_NSData) Bytes() unsafe.Pointer {
	ret := C.NSData_inst_Bytes(
		unsafe.Pointer(x.Pointer()),
	)

	return ret
}

// Length returns the number of bytes contained by the data object.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1416769-length?language=objc for details.
func (x gen_NSData) Length() NSUInteger {
	ret := C.NSData_inst_Length(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)
}

// Description returns a string that contains a hexadecimal representation of the data object’s contents in a property list format.
//
// See https://developer.apple.com/documentation/foundation/nsdata/1412579-description?language=objc for details.
func (x gen_NSData) Description() NSString {
	ret := C.NSData_inst_Description(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

type NSDictionaryRef interface {
	Pointer() uintptr
	Init_AsNSDictionary() NSDictionary
}

type gen_NSDictionary struct {
	objc.Object
}

func NSDictionary_FromPointer(ptr unsafe.Pointer) NSDictionary {
	return NSDictionary{gen_NSDictionary{
		objc.Object_FromPointer(ptr),
	}}
}

func NSDictionary_FromRef(ref objc.Ref) NSDictionary {
	return NSDictionary_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// DescriptionWithLocale returns a string object that represents the contents of the dictionary, formatted as a property list.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1417665-descriptionwithlocale?language=objc for details.
func (x gen_NSDictionary) DescriptionWithLocale(
	locale objc.Ref,
) NSString {
	ret := C.NSDictionary_inst_DescriptionWithLocale(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(locale),
	)

	return NSString_FromPointer(ret)
}

// DescriptionWithLocaleIndent returns a string object that represents the contents of the dictionary, formatted as a property list.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1412690-descriptionwithlocale?language=objc for details.
func (x gen_NSDictionary) DescriptionWithLocaleIndent(
	locale objc.Ref,
	level NSUInteger,
) NSString {
	ret := C.NSDictionary_inst_DescriptionWithLocaleIndent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(locale),
		C.ulong(level),
	)

	return NSString_FromPointer(ret)
}

// FileExtensionHidden returns a Boolean value indicating whether the file hides its extension.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1413177-fileextensionhidden?language=objc for details.
func (x gen_NSDictionary) FileExtensionHidden() bool {
	ret := C.NSDictionary_inst_FileExtensionHidden(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// FileGroupOwnerAccountID returns file’s group owner account ID.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1413626-filegroupowneraccountid?language=objc for details.
func (x gen_NSDictionary) FileGroupOwnerAccountID() NSNumber {
	ret := C.NSDictionary_inst_FileGroupOwnerAccountID(
		unsafe.Pointer(x.Pointer()),
	)

	return NSNumber_FromPointer(ret)
}

// FileGroupOwnerAccountName returns the file’s group owner account name.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1416788-filegroupowneraccountname?language=objc for details.
func (x gen_NSDictionary) FileGroupOwnerAccountName() NSString {
	ret := C.NSDictionary_inst_FileGroupOwnerAccountName(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// FileIsAppendOnly returns a Boolean value indicating whether the file is append only.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1416083-fileisappendonly?language=objc for details.
func (x gen_NSDictionary) FileIsAppendOnly() bool {
	ret := C.NSDictionary_inst_FileIsAppendOnly(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// FileIsImmutable returns a Boolean value indicating whether the file is immutable.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1416500-fileisimmutable?language=objc for details.
func (x gen_NSDictionary) FileIsImmutable() bool {
	ret := C.NSDictionary_inst_FileIsImmutable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// FileOwnerAccountID returns the file’s owner account ID.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1412281-fileowneraccountid?language=objc for details.
func (x gen_NSDictionary) FileOwnerAccountID() NSNumber {
	ret := C.NSDictionary_inst_FileOwnerAccountID(
		unsafe.Pointer(x.Pointer()),
	)

	return NSNumber_FromPointer(ret)
}

// FileOwnerAccountName returns the file’s owner account name.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1417533-fileowneraccountname?language=objc for details.
func (x gen_NSDictionary) FileOwnerAccountName() NSString {
	ret := C.NSDictionary_inst_FileOwnerAccountName(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// FilePosixPermissions returns the file’s POSIX permissions.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1409446-fileposixpermissions?language=objc for details.
func (x gen_NSDictionary) FilePosixPermissions() NSUInteger {
	ret := C.NSDictionary_inst_FilePosixPermissions(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)
}

// FileSystemFileNumber returns the filesystem file number.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1408396-filesystemfilenumber?language=objc for details.
func (x gen_NSDictionary) FileSystemFileNumber() NSUInteger {
	ret := C.NSDictionary_inst_FileSystemFileNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)
}

// FileSystemNumber returns the filesystem number.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1415329-filesystemnumber?language=objc for details.
func (x gen_NSDictionary) FileSystemNumber() NSInteger {
	ret := C.NSDictionary_inst_FileSystemNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return NSInteger(ret)
}

// FileType returns the file type.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1416809-filetype?language=objc for details.
func (x gen_NSDictionary) FileType() NSString {
	ret := C.NSDictionary_inst_FileType(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// Init initializes a newly allocated dictionary.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1418147-init?language=objc for details.
func (x gen_NSDictionary) Init() NSDictionary {
	ret := C.NSDictionary_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSDictionary_FromPointer(ret)
}

// Init_AsNSDictionary is a typed version of Init.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1418147-init?language=objc for details.
func (x gen_NSDictionary) Init_AsNSDictionary() NSDictionary {
	ret := C.NSDictionary_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSDictionary_FromPointer(ret)
}

// InitWithDictionary initializes a newly allocated dictionary by placing in it the keys and values contained in another given dictionary.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1418434-initwithdictionary?language=objc for details.
func (x gen_NSDictionary) InitWithDictionary(
	otherDictionary NSDictionaryRef,
) NSDictionary {
	ret := C.NSDictionary_inst_InitWithDictionary(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(otherDictionary),
	)

	return NSDictionary_FromPointer(ret)
}

// InitWithDictionaryCopyItems initializes a newly allocated dictionary using the objects contained in another given dictionary.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1410124-initwithdictionary?language=objc for details.
func (x gen_NSDictionary) InitWithDictionaryCopyItems(
	otherDictionary NSDictionaryRef,
	flag bool,
) NSDictionary {
	ret := C.NSDictionary_inst_InitWithDictionaryCopyItems(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(otherDictionary),
		convertToObjCBool(flag),
	)

	return NSDictionary_FromPointer(ret)
}

// InitWithObjectsForKeys initializes a newly allocated dictionary with key-value pairs constructed from the provided arrays of keys and objects.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1410010-initwithobjects?language=objc for details.
func (x gen_NSDictionary) InitWithObjectsForKeys(
	objects NSArrayRef,
	keys NSArrayRef,
) NSDictionary {
	ret := C.NSDictionary_inst_InitWithObjectsForKeys(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(objects),
		objc.RefPointer(keys),
	)

	return NSDictionary_FromPointer(ret)
}

// IsEqualToDictionary returns a Boolean value that indicates whether the contents of the receiving dictionary are equal to the contents of another given dictionary.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1415445-isequaltodictionary?language=objc for details.
func (x gen_NSDictionary) IsEqualToDictionary(
	otherDictionary NSDictionaryRef,
) bool {
	ret := C.NSDictionary_inst_IsEqualToDictionary(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(otherDictionary),
	)

	return convertObjCBoolToGo(ret)
}

// KeysSortedByValueUsingSelector returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1412484-keyssortedbyvalueusingselector?language=objc for details.
func (x gen_NSDictionary) KeysSortedByValueUsingSelector(
	comparator objc.Selector,
) NSArray {
	ret := C.NSDictionary_inst_KeysSortedByValueUsingSelector(
		unsafe.Pointer(x.Pointer()),
		comparator.SelectorAddress(),
	)

	return NSArray_FromPointer(ret)
}

// Count returns the number of entries in the dictionary.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1409628-count?language=objc for details.
func (x gen_NSDictionary) Count() NSUInteger {
	ret := C.NSDictionary_inst_Count(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)
}

// AllKeys returns a new array containing the dictionary’s keys, or an empty array if the dictionary has no entries.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1409150-allkeys?language=objc for details.
func (x gen_NSDictionary) AllKeys() NSArray {
	ret := C.NSDictionary_inst_AllKeys(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}

// AllValues returns a new array containing the dictionary’s values, or an empty array if the dictionary has no entries.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1408915-allvalues?language=objc for details.
func (x gen_NSDictionary) AllValues() NSArray {
	ret := C.NSDictionary_inst_AllValues(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}

// Description returns a string that represents the contents of the dictionary, formatted as a property list.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1410799-description?language=objc for details.
func (x gen_NSDictionary) Description() NSString {
	ret := C.NSDictionary_inst_Description(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// DescriptionInStringsFileFormat returns a string that represents the contents of the dictionary, formatted in .strings file format.
//
// See https://developer.apple.com/documentation/foundation/nsdictionary/1413282-descriptioninstringsfileformat?language=objc for details.
func (x gen_NSDictionary) DescriptionInStringsFileFormat() NSString {
	ret := C.NSDictionary_inst_DescriptionInStringsFileFormat(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

type NSNumberRef interface {
	Pointer() uintptr
	Init_AsNSNumber() NSNumber
}

type gen_NSNumber struct {
	objc.Object
}

func NSNumber_FromPointer(ptr unsafe.Pointer) NSNumber {
	return NSNumber{gen_NSNumber{
		objc.Object_FromPointer(ptr),
	}}
}

func NSNumber_FromRef(ref objc.Ref) NSNumber {
	return NSNumber_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// DescriptionWithLocale returns a string that represents the contents of the number object for a given locale.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1409984-descriptionwithlocale?language=objc for details.
func (x gen_NSNumber) DescriptionWithLocale(
	locale objc.Ref,
) NSString {
	ret := C.NSNumber_inst_DescriptionWithLocale(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(locale),
	)

	return NSString_FromPointer(ret)
}

// InitWithBool returns an NSNumber object initialized to contain a given value, treated as a BOOL.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1415728-initwithbool?language=objc for details.
func (x gen_NSNumber) InitWithBool(
	value bool,
) NSNumber {
	ret := C.NSNumber_inst_InitWithBool(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return NSNumber_FromPointer(ret)
}

// InitWithInt returns an NSNumber object initialized to contain a given value, treated as a signed int.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1407580-initwithint?language=objc for details.
func (x gen_NSNumber) InitWithInt(
	value int32,
) NSNumber {
	ret := C.NSNumber_inst_InitWithInt(
		unsafe.Pointer(x.Pointer()),
		C.int(value),
	)

	return NSNumber_FromPointer(ret)
}

// InitWithInteger returns an NSNumber object initialized to contain a given value, treated as an NSInteger.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1409397-initwithinteger?language=objc for details.
func (x gen_NSNumber) InitWithInteger(
	value NSInteger,
) NSNumber {
	ret := C.NSNumber_inst_InitWithInteger(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return NSNumber_FromPointer(ret)
}

// InitWithUnsignedInt returns an NSNumber object initialized to contain a given value, treated as an unsigned int.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1414598-initwithunsignedint?language=objc for details.
func (x gen_NSNumber) InitWithUnsignedInt(
	value int32,
) NSNumber {
	ret := C.NSNumber_inst_InitWithUnsignedInt(
		unsafe.Pointer(x.Pointer()),
		C.int(value),
	)

	return NSNumber_FromPointer(ret)
}

// InitWithUnsignedInteger returns an NSNumber object initialized to contain a given value, treated as an NSUInteger.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1412531-initwithunsignedinteger?language=objc for details.
func (x gen_NSNumber) InitWithUnsignedInteger(
	value NSUInteger,
) NSNumber {
	ret := C.NSNumber_inst_InitWithUnsignedInteger(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)

	return NSNumber_FromPointer(ret)
}

// IsEqualToNumber returns a Boolean value that indicates whether the number object’s value and a given number are equal.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1411315-isequaltonumber?language=objc for details.
func (x gen_NSNumber) IsEqualToNumber(
	number NSNumberRef,
) bool {
	ret := C.NSNumber_inst_IsEqualToNumber(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(number),
	)

	return convertObjCBoolToGo(ret)
}

// Init initializes a new instance of the NSNumber class.
func (x gen_NSNumber) Init() NSNumber {
	ret := C.NSNumber_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSNumber_FromPointer(ret)
}

// Init_AsNSNumber is a typed version of Init.
func (x gen_NSNumber) Init_AsNSNumber() NSNumber {
	ret := C.NSNumber_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSNumber_FromPointer(ret)
}

// BoolValue returns the number object's value expressed as a Boolean value.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1410865-boolvalue?language=objc for details.
func (x gen_NSNumber) BoolValue() bool {
	ret := C.NSNumber_inst_BoolValue(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IntValue returns the number object's value expressed as an int, converted as necessary.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1407153-intvalue?language=objc for details.
func (x gen_NSNumber) IntValue() int32 {
	ret := C.NSNumber_inst_IntValue(
		unsafe.Pointer(x.Pointer()),
	)

	return int32(ret)
}

// IntegerValue returns the number object's value expressed as an NSInteger object, converted as necessary.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1412554-integervalue?language=objc for details.
func (x gen_NSNumber) IntegerValue() NSInteger {
	ret := C.NSNumber_inst_IntegerValue(
		unsafe.Pointer(x.Pointer()),
	)

	return NSInteger(ret)
}

// UnsignedIntegerValue returns the number object's value expressed as an NSUInteger object, converted as necessary.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1413324-unsignedintegervalue?language=objc for details.
func (x gen_NSNumber) UnsignedIntegerValue() NSUInteger {
	ret := C.NSNumber_inst_UnsignedIntegerValue(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)
}

// UnsignedIntValue returns the number object's value expressed as an unsigned int, converted as necessary.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1417875-unsignedintvalue?language=objc for details.
func (x gen_NSNumber) UnsignedIntValue() int32 {
	ret := C.NSNumber_inst_UnsignedIntValue(
		unsafe.Pointer(x.Pointer()),
	)

	return int32(ret)
}

// StringValue returns the number object's value expressed as a human-readable string.
//
// See https://developer.apple.com/documentation/foundation/nsnumber/1415802-stringvalue?language=objc for details.
func (x gen_NSNumber) StringValue() NSString {
	ret := C.NSNumber_inst_StringValue(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

type NSRunLoopRef interface {
	Pointer() uintptr
	Init_AsNSRunLoop() NSRunLoop
}

type gen_NSRunLoop struct {
	objc.Object
}

func NSRunLoop_FromPointer(ptr unsafe.Pointer) NSRunLoop {
	return NSRunLoop{gen_NSRunLoop{
		objc.Object_FromPointer(ptr),
	}}
}

func NSRunLoop_FromRef(ref objc.Ref) NSRunLoop {
	return NSRunLoop_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// CancelPerformSelectorTargetArgument cancels the sending of a previously scheduled message.
//
// See https://developer.apple.com/documentation/foundation/nsrunloop/1418077-cancelperformselector?language=objc for details.
func (x gen_NSRunLoop) CancelPerformSelectorTargetArgument(
	aSelector objc.Selector,
	target objc.Ref,
	arg objc.Ref,
) {
	C.NSRunLoop_inst_CancelPerformSelectorTargetArgument(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(arg),
	)

	return
}

// CancelPerformSelectorsWithTarget cancels all outstanding ordered performs scheduled with a given target.
//
// See https://developer.apple.com/documentation/foundation/nsrunloop/1414208-cancelperformselectorswithtarget?language=objc for details.
func (x gen_NSRunLoop) CancelPerformSelectorsWithTarget(
	target objc.Ref,
) {
	C.NSRunLoop_inst_CancelPerformSelectorsWithTarget(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(target),
	)

	return
}

// PerformSelectorTargetArgumentOrderModes schedules the sending of a message on the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsrunloop/1409310-performselector?language=objc for details.
func (x gen_NSRunLoop) PerformSelectorTargetArgumentOrderModes(
	aSelector objc.Selector,
	target objc.Ref,
	arg objc.Ref,
	order NSUInteger,
	modes NSArrayRef,
) {
	C.NSRunLoop_inst_PerformSelectorTargetArgumentOrderModes(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(arg),
		C.ulong(order),
		objc.RefPointer(modes),
	)

	return
}

// Run puts the receiver into a permanent loop, during which time it processes data from all attached input sources.
//
// See https://developer.apple.com/documentation/foundation/nsrunloop/1412430-run?language=objc for details.
func (x gen_NSRunLoop) Run() {
	C.NSRunLoop_inst_Run(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Init initializes a new instance of the NSRunLoop class.
func (x gen_NSRunLoop) Init() NSRunLoop {
	ret := C.NSRunLoop_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSRunLoop_FromPointer(ret)
}

// Init_AsNSRunLoop is a typed version of Init.
func (x gen_NSRunLoop) Init_AsNSRunLoop() NSRunLoop {
	ret := C.NSRunLoop_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSRunLoop_FromPointer(ret)
}

type NSStringRef interface {
	Pointer() uintptr
	Init_AsNSString() NSString
}

type gen_NSString struct {
	objc.Object
}

func NSString_FromPointer(ptr unsafe.Pointer) NSString {
	return NSString{gen_NSString{
		objc.Object_FromPointer(ptr),
	}}
}

func NSString_FromRef(ref objc.Ref) NSString {
	return NSString_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// CanBeConvertedToEncoding returns a Boolean value that indicates whether the receiver can be converted to a given encoding without loss of information.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1409496-canbeconvertedtoencoding?language=objc for details.
func (x gen_NSString) CanBeConvertedToEncoding(
	encoding NSStringEncoding,
) bool {
	ret := C.NSString_inst_CanBeConvertedToEncoding(
		unsafe.Pointer(x.Pointer()),
		C.ulong(encoding),
	)

	return convertObjCBoolToGo(ret)
}

// CharacterAtIndex returns the character at a given UTF-16 code unit index.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1414645-characteratindex?language=objc for details.
func (x gen_NSString) CharacterAtIndex(
	index NSUInteger,
) Unichar {
	ret := C.NSString_inst_CharacterAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(index),
	)

	return Unichar(ret)
}

// CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes interprets the receiver as a path in the file system and attempts to perform filename completion, returning a numeric value that indicates whether a match was possible, and by reference the longest path that matches the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1411841-completepathintostring?language=objc for details.
func (x gen_NSString) CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes(
	outputName NSStringRef,
	flag bool,
	outputArray NSArrayRef,
	filterTypes NSArrayRef,
) NSUInteger {
	ret := C.NSString_inst_CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(outputName),
		convertToObjCBool(flag),
		objc.RefPointer(outputArray),
		objc.RefPointer(filterTypes),
	)

	return NSUInteger(ret)
}

// ComponentsSeparatedByString returns an array containing substrings from the receiver that have been divided by a given separator.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1413214-componentsseparatedbystring?language=objc for details.
func (x gen_NSString) ComponentsSeparatedByString(
	separator NSStringRef,
) NSArray {
	ret := C.NSString_inst_ComponentsSeparatedByString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(separator),
	)

	return NSArray_FromPointer(ret)
}

// ContainsString returns a Boolean value indicating whether the string contains a given string by performing a case-sensitive, locale-unaware search.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1414563-containsstring?language=objc for details.
func (x gen_NSString) ContainsString(
	str NSStringRef,
) bool {
	ret := C.NSString_inst_ContainsString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return convertObjCBoolToGo(ret)
}

// DataUsingEncoding returns an NSData object containing a representation of the receiver encoded using a given encoding.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1416696-datausingencoding?language=objc for details.
func (x gen_NSString) DataUsingEncoding(
	encoding NSStringEncoding,
) NSData {
	ret := C.NSString_inst_DataUsingEncoding(
		unsafe.Pointer(x.Pointer()),
		C.ulong(encoding),
	)

	return NSData_FromPointer(ret)
}

// DataUsingEncodingAllowLossyConversion returns an NSData object containing a representation of the receiver encoded using a given encoding.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1413692-datausingencoding?language=objc for details.
func (x gen_NSString) DataUsingEncodingAllowLossyConversion(
	encoding NSStringEncoding,
	lossy bool,
) NSData {
	ret := C.NSString_inst_DataUsingEncodingAllowLossyConversion(
		unsafe.Pointer(x.Pointer()),
		C.ulong(encoding),
		convertToObjCBool(lossy),
	)

	return NSData_FromPointer(ret)
}

// DrawInRectWithAttributes draws the attributed string inside the specified bounding rectangle.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1529855-drawinrect?language=objc for details.
func (x gen_NSString) DrawInRectWithAttributes(
	rect NSRect,
	attrs NSDictionaryRef,
) {
	C.NSString_inst_DrawInRectWithAttributes(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(attrs),
	)

	return
}

// HasPrefix returns a Boolean value that indicates whether a given string matches the beginning characters of the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1410309-hasprefix?language=objc for details.
func (x gen_NSString) HasPrefix(
	str NSStringRef,
) bool {
	ret := C.NSString_inst_HasPrefix(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return convertObjCBoolToGo(ret)
}

// HasSuffix returns a Boolean value that indicates whether a given string matches the ending characters of the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1416529-hassuffix?language=objc for details.
func (x gen_NSString) HasSuffix(
	str NSStringRef,
) bool {
	ret := C.NSString_inst_HasSuffix(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return convertObjCBoolToGo(ret)
}

// Init returns an initialized NSString object that contains no characters.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1409306-init?language=objc for details.
func (x gen_NSString) Init() NSString {
	ret := C.NSString_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// Init_AsNSString is a typed version of Init.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1409306-init?language=objc for details.
func (x gen_NSString) Init_AsNSString() NSString {
	ret := C.NSString_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// InitWithBytesLengthEncoding returns an initialized NSString object containing a given number of bytes from a given buffer of bytes interpreted in a given encoding.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1407339-initwithbytes?language=objc for details.
func (x gen_NSString) InitWithBytesLengthEncoding(
	bytes unsafe.Pointer,
	len NSUInteger,
	encoding NSStringEncoding,
) NSString {
	ret := C.NSString_inst_InitWithBytesLengthEncoding(
		unsafe.Pointer(x.Pointer()),
		bytes,
		C.ulong(len),
		C.ulong(encoding),
	)

	return NSString_FromPointer(ret)
}

// InitWithBytesNoCopyLengthEncodingFreeWhenDone returns an initialized NSString object that contains a given number of bytes from a given buffer of bytes interpreted in a given encoding, and optionally frees the buffer.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1413830-initwithbytesnocopy?language=objc for details.
func (x gen_NSString) InitWithBytesNoCopyLengthEncodingFreeWhenDone(
	bytes unsafe.Pointer,
	len NSUInteger,
	encoding NSStringEncoding,
	freeBuffer bool,
) NSString {
	ret := C.NSString_inst_InitWithBytesNoCopyLengthEncodingFreeWhenDone(
		unsafe.Pointer(x.Pointer()),
		bytes,
		C.ulong(len),
		C.ulong(encoding),
		convertToObjCBool(freeBuffer),
	)

	return NSString_FromPointer(ret)
}

// InitWithDataEncoding returns an NSString object initialized by converting given data into UTF-16 code units using a given encoding.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1416374-initwithdata?language=objc for details.
func (x gen_NSString) InitWithDataEncoding(
	data NSDataRef,
	encoding NSStringEncoding,
) NSString {
	ret := C.NSString_inst_InitWithDataEncoding(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		C.ulong(encoding),
	)

	return NSString_FromPointer(ret)
}

// InitWithString returns an NSString object initialized by copying the characters from another given string.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1411293-initwithstring?language=objc for details.
func (x gen_NSString) InitWithString(
	aString NSStringRef,
) NSString {
	ret := C.NSString_inst_InitWithString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(aString),
	)

	return NSString_FromPointer(ret)
}

// IsEqualToString returns a Boolean value that indicates whether a given string is equal to the receiver using a literal Unicode-based comparison.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1407803-isequaltostring?language=objc for details.
func (x gen_NSString) IsEqualToString(
	aString NSStringRef,
) bool {
	ret := C.NSString_inst_IsEqualToString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(aString),
	)

	return convertObjCBoolToGo(ret)
}

// LengthOfBytesUsingEncoding returns the number of bytes required to store the receiver in a given encoding.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1410710-lengthofbytesusingencoding?language=objc for details.
func (x gen_NSString) LengthOfBytesUsingEncoding(
	enc NSStringEncoding,
) NSUInteger {
	ret := C.NSString_inst_LengthOfBytesUsingEncoding(
		unsafe.Pointer(x.Pointer()),
		C.ulong(enc),
	)

	return NSUInteger(ret)
}

// LocalizedCaseInsensitiveContainsString returns a Boolean value indicating whether the string contains a given string by performing a case-insensitive, locale-aware search.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1412098-localizedcaseinsensitivecontains?language=objc for details.
func (x gen_NSString) LocalizedCaseInsensitiveContainsString(
	str NSStringRef,
) bool {
	ret := C.NSString_inst_LocalizedCaseInsensitiveContainsString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return convertObjCBoolToGo(ret)
}

// LocalizedStandardContainsString returns a Boolean value indicating whether the string contains a given string by performing a case and diacritic insensitive, locale-aware search.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1416328-localizedstandardcontainsstring?language=objc for details.
func (x gen_NSString) LocalizedStandardContainsString(
	str NSStringRef,
) bool {
	ret := C.NSString_inst_LocalizedStandardContainsString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return convertObjCBoolToGo(ret)
}

// MaximumLengthOfBytesUsingEncoding returns the maximum number of bytes needed to store the receiver in a given encoding.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1411611-maximumlengthofbytesusingencodin?language=objc for details.
func (x gen_NSString) MaximumLengthOfBytesUsingEncoding(
	enc NSStringEncoding,
) NSUInteger {
	ret := C.NSString_inst_MaximumLengthOfBytesUsingEncoding(
		unsafe.Pointer(x.Pointer()),
		C.ulong(enc),
	)

	return NSUInteger(ret)
}

// PropertyList parses the receiver as a text representation of a property list, returning an NSString, NSData, NSArray, or NSDictionary object, according to the topmost element.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1413115-propertylist?language=objc for details.
func (x gen_NSString) PropertyList() objc.Object {
	ret := C.NSString_inst_PropertyList(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// PropertyListFromStringsFileFormat returns a dictionary object initialized with the keys and values found in the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1407697-propertylistfromstringsfileforma?language=objc for details.
func (x gen_NSString) PropertyListFromStringsFileFormat() NSDictionary {
	ret := C.NSString_inst_PropertyListFromStringsFileFormat(
		unsafe.Pointer(x.Pointer()),
	)

	return NSDictionary_FromPointer(ret)
}

// SizeWithAttributes returns the bounding box size the receiver occupies when drawn with the given attributes.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1531844-sizewithattributes?language=objc for details.
func (x gen_NSString) SizeWithAttributes(
	attrs NSDictionaryRef,
) NSSize {
	ret := C.NSString_inst_SizeWithAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(attrs),
	)

	return *(*NSSize)(unsafe.Pointer(&ret))
}

// StringByAppendingPathComponent returns a new string made by appending to the receiver a given string.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1417069-stringbyappendingpathcomponent?language=objc for details.
func (x gen_NSString) StringByAppendingPathComponent(
	str NSStringRef,
) NSString {
	ret := C.NSString_inst_StringByAppendingPathComponent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return NSString_FromPointer(ret)
}

// StringByAppendingPathExtension returns a new string made by appending to the receiver an extension separator followed by a given extension.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1412501-stringbyappendingpathextension?language=objc for details.
func (x gen_NSString) StringByAppendingPathExtension(
	str NSStringRef,
) NSString {
	ret := C.NSString_inst_StringByAppendingPathExtension(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return NSString_FromPointer(ret)
}

// StringByAppendingString returns a new string made by appending a given string to the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1412307-stringbyappendingstring?language=objc for details.
func (x gen_NSString) StringByAppendingString(
	aString NSStringRef,
) NSString {
	ret := C.NSString_inst_StringByAppendingString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(aString),
	)

	return NSString_FromPointer(ret)
}

// StringByPaddingToLengthWithStringStartingAtIndex returns a new string formed from the receiver by either removing characters from the end, or by appending as many occurrences as necessary of a given pad string.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1416395-stringbypaddingtolength?language=objc for details.
func (x gen_NSString) StringByPaddingToLengthWithStringStartingAtIndex(
	newLength NSUInteger,
	padString NSStringRef,
	padIndex NSUInteger,
) NSString {
	ret := C.NSString_inst_StringByPaddingToLengthWithStringStartingAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(newLength),
		objc.RefPointer(padString),
		C.ulong(padIndex),
	)

	return NSString_FromPointer(ret)
}

// StringByReplacingOccurrencesOfStringWithString returns a new string in which all occurrences of a target string in the receiver are replaced by another given string.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1412937-stringbyreplacingoccurrencesofst?language=objc for details.
func (x gen_NSString) StringByReplacingOccurrencesOfStringWithString(
	target NSStringRef,
	replacement NSStringRef,
) NSString {
	ret := C.NSString_inst_StringByReplacingOccurrencesOfStringWithString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(target),
		objc.RefPointer(replacement),
	)

	return NSString_FromPointer(ret)
}

// StringsByAppendingPaths returns an array of strings made by separately appending to the receiver each string in a given array.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1415100-stringsbyappendingpaths?language=objc for details.
func (x gen_NSString) StringsByAppendingPaths(
	paths NSArrayRef,
) NSArray {
	ret := C.NSString_inst_StringsByAppendingPaths(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(paths),
	)

	return NSArray_FromPointer(ret)
}

// SubstringFromIndex returns a new string containing the characters of the receiver from the one at a given index to the end.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1414368-substringfromindex?language=objc for details.
func (x gen_NSString) SubstringFromIndex(
	from NSUInteger,
) NSString {
	ret := C.NSString_inst_SubstringFromIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(from),
	)

	return NSString_FromPointer(ret)
}

// SubstringToIndex returns a new string containing the characters of the receiver up to, but not including, the one at a given index.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1408017-substringtoindex?language=objc for details.
func (x gen_NSString) SubstringToIndex(
	to NSUInteger,
) NSString {
	ret := C.NSString_inst_SubstringToIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(to),
	)

	return NSString_FromPointer(ret)
}

// VariantFittingPresentationWidth returns a string variation suitable for the specified presentation width.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1413104-variantfittingpresentationwidth?language=objc for details.
func (x gen_NSString) VariantFittingPresentationWidth(
	width NSInteger,
) NSString {
	ret := C.NSString_inst_VariantFittingPresentationWidth(
		unsafe.Pointer(x.Pointer()),
		C.long(width),
	)

	return NSString_FromPointer(ret)
}

// Length returns the number of UTF-16 code units in the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1414212-length?language=objc for details.
func (x gen_NSString) Length() NSUInteger {
	ret := C.NSString_inst_Length(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)
}

// Hash an unsigned integer that can be used as a hash table address.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1417245-hash?language=objc for details.
func (x gen_NSString) Hash() NSUInteger {
	ret := C.NSString_inst_Hash(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)
}

// LowercaseString returns a lowercase representation of the string.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1408467-lowercasestring?language=objc for details.
func (x gen_NSString) LowercaseString() NSString {
	ret := C.NSString_inst_LowercaseString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// LocalizedLowercaseString returns a version of the string with all letters converted to lowercase, taking into account the current locale.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1414125-localizedlowercasestring?language=objc for details.
func (x gen_NSString) LocalizedLowercaseString() NSString {
	ret := C.NSString_inst_LocalizedLowercaseString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// UppercaseString an uppercase representation of the string.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1409855-uppercasestring?language=objc for details.
func (x gen_NSString) UppercaseString() NSString {
	ret := C.NSString_inst_UppercaseString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// LocalizedUppercaseString returns a version of the string with all letters converted to uppercase, taking into account the current locale.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1413331-localizeduppercasestring?language=objc for details.
func (x gen_NSString) LocalizedUppercaseString() NSString {
	ret := C.NSString_inst_LocalizedUppercaseString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// CapitalizedString returns a capitalized representation of the string.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1416784-capitalizedstring?language=objc for details.
func (x gen_NSString) CapitalizedString() NSString {
	ret := C.NSString_inst_CapitalizedString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// LocalizedCapitalizedString returns a capitalized representation of the receiver using the current locale.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1414885-localizedcapitalizedstring?language=objc for details.
func (x gen_NSString) LocalizedCapitalizedString() NSString {
	ret := C.NSString_inst_LocalizedCapitalizedString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// DecomposedStringWithCanonicalMapping returns a string made by normalizing the string’s contents using the Unicode Normalization Form D.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1409474-decomposedstringwithcanonicalmap?language=objc for details.
func (x gen_NSString) DecomposedStringWithCanonicalMapping() NSString {
	ret := C.NSString_inst_DecomposedStringWithCanonicalMapping(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// DecomposedStringWithCompatibilityMapping returns a string made by normalizing the receiver’s contents using the Unicode Normalization Form KD.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1415417-decomposedstringwithcompatibilit?language=objc for details.
func (x gen_NSString) DecomposedStringWithCompatibilityMapping() NSString {
	ret := C.NSString_inst_DecomposedStringWithCompatibilityMapping(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// PrecomposedStringWithCanonicalMapping returns a string made by normalizing the string’s contents using the Unicode Normalization Form C.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1412645-precomposedstringwithcanonicalma?language=objc for details.
func (x gen_NSString) PrecomposedStringWithCanonicalMapping() NSString {
	ret := C.NSString_inst_PrecomposedStringWithCanonicalMapping(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// PrecomposedStringWithCompatibilityMapping returns a string made by normalizing the receiver’s contents using the Unicode Normalization Form KC.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1412625-precomposedstringwithcompatibili?language=objc for details.
func (x gen_NSString) PrecomposedStringWithCompatibilityMapping() NSString {
	ret := C.NSString_inst_PrecomposedStringWithCompatibilityMapping(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// IntValue returns the integer value of the string.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1414988-intvalue?language=objc for details.
func (x gen_NSString) IntValue() int32 {
	ret := C.NSString_inst_IntValue(
		unsafe.Pointer(x.Pointer()),
	)

	return int32(ret)
}

// IntegerValue returns the NSInteger value of the string.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1410267-integervalue?language=objc for details.
func (x gen_NSString) IntegerValue() NSInteger {
	ret := C.NSString_inst_IntegerValue(
		unsafe.Pointer(x.Pointer()),
	)

	return NSInteger(ret)
}

// BoolValue returns the Boolean value of the string.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1409420-boolvalue?language=objc for details.
func (x gen_NSString) BoolValue() bool {
	ret := C.NSString_inst_BoolValue(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Description this NSString object.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1410889-description?language=objc for details.
func (x gen_NSString) Description() NSString {
	ret := C.NSString_inst_Description(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// FastestEncoding returns the fastest encoding to which the receiver may be converted without loss of information.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1409567-fastestencoding?language=objc for details.
func (x gen_NSString) FastestEncoding() NSStringEncoding {
	ret := C.NSString_inst_FastestEncoding(
		unsafe.Pointer(x.Pointer()),
	)

	return NSStringEncoding(ret)
}

// SmallestEncoding returns the smallest encoding to which the receiver can be converted without loss of information.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1418037-smallestencoding?language=objc for details.
func (x gen_NSString) SmallestEncoding() NSStringEncoding {
	ret := C.NSString_inst_SmallestEncoding(
		unsafe.Pointer(x.Pointer()),
	)

	return NSStringEncoding(ret)
}

// PathComponents returns the file-system path components of the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1414489-pathcomponents?language=objc for details.
func (x gen_NSString) PathComponents() NSArray {
	ret := C.NSString_inst_PathComponents(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}

// IsAbsolutePath returns a Boolean value that indicates whether the receiver represents an absolute path.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1409068-absolutepath?language=objc for details.
func (x gen_NSString) IsAbsolutePath() bool {
	ret := C.NSString_inst_IsAbsolutePath(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// LastPathComponent returns the last path component of the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1416528-lastpathcomponent?language=objc for details.
func (x gen_NSString) LastPathComponent() NSString {
	ret := C.NSString_inst_LastPathComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// PathExtension returns the path extension, if any, of the string as interpreted as a path.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1407801-pathextension?language=objc for details.
func (x gen_NSString) PathExtension() NSString {
	ret := C.NSString_inst_PathExtension(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// StringByAbbreviatingWithTildeInPath returns a new string that replaces the current home directory portion of the current path with a tilde (~) character.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1407943-stringbyabbreviatingwithtildeinp?language=objc for details.
func (x gen_NSString) StringByAbbreviatingWithTildeInPath() NSString {
	ret := C.NSString_inst_StringByAbbreviatingWithTildeInPath(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// StringByDeletingLastPathComponent returns a new string made by deleting the last path component from the receiver, along with any final path separator.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1411141-stringbydeletinglastpathcomponen?language=objc for details.
func (x gen_NSString) StringByDeletingLastPathComponent() NSString {
	ret := C.NSString_inst_StringByDeletingLastPathComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// StringByDeletingPathExtension returns a new string made by deleting the extension (if any, and only the last) from the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1418214-stringbydeletingpathextension?language=objc for details.
func (x gen_NSString) StringByDeletingPathExtension() NSString {
	ret := C.NSString_inst_StringByDeletingPathExtension(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// StringByExpandingTildeInPath returns a new string made by expanding the initial component of the receiver to its full path value.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1407716-stringbyexpandingtildeinpath?language=objc for details.
func (x gen_NSString) StringByExpandingTildeInPath() NSString {
	ret := C.NSString_inst_StringByExpandingTildeInPath(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// StringByResolvingSymlinksInPath returns a new string made from the receiver by resolving all symbolic links and standardizing path.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1417783-stringbyresolvingsymlinksinpath?language=objc for details.
func (x gen_NSString) StringByResolvingSymlinksInPath() NSString {
	ret := C.NSString_inst_StringByResolvingSymlinksInPath(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// StringByStandardizingPath returns a new string made by removing extraneous path components from the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1407194-stringbystandardizingpath?language=objc for details.
func (x gen_NSString) StringByStandardizingPath() NSString {
	ret := C.NSString_inst_StringByStandardizingPath(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// StringByRemovingPercentEncoding returns a new string made from the receiver by replacing all percent encoded sequences with the matching UTF-8 characters.
//
// See https://developer.apple.com/documentation/foundation/nsstring/1409569-stringbyremovingpercentencoding?language=objc for details.
func (x gen_NSString) StringByRemovingPercentEncoding() NSString {
	ret := C.NSString_inst_StringByRemovingPercentEncoding(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

type NSThreadRef interface {
	Pointer() uintptr
	Init_AsNSThread() NSThread
}

type gen_NSThread struct {
	objc.Object
}

func NSThread_FromPointer(ptr unsafe.Pointer) NSThread {
	return NSThread{gen_NSThread{
		objc.Object_FromPointer(ptr),
	}}
}

func NSThread_FromRef(ref objc.Ref) NSThread {
	return NSThread_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Cancel changes the cancelled state of the receiver to indicate that it should exit.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1411303-cancel?language=objc for details.
func (x gen_NSThread) Cancel() {
	C.NSThread_inst_Cancel(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Init returns an initialized NSThread object.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1416464-init?language=objc for details.
func (x gen_NSThread) Init() NSThread {
	ret := C.NSThread_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSThread_FromPointer(ret)
}

// Init_AsNSThread is a typed version of Init.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1416464-init?language=objc for details.
func (x gen_NSThread) Init_AsNSThread() NSThread {
	ret := C.NSThread_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSThread_FromPointer(ret)
}

// InitWithTargetSelectorObject returns an NSThread object initialized with the given arguments.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1414773-initwithtarget?language=objc for details.
func (x gen_NSThread) InitWithTargetSelectorObject(
	target objc.Ref,
	selector objc.Selector,
	argument objc.Ref,
) NSThread {
	ret := C.NSThread_inst_InitWithTargetSelectorObject(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(target),
		selector.SelectorAddress(),
		objc.RefPointer(argument),
	)

	return NSThread_FromPointer(ret)
}

// Main returns the main entry point routine for the thread.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1418421-main?language=objc for details.
func (x gen_NSThread) Main() {
	C.NSThread_inst_Main(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Start starts the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1418166-start?language=objc for details.
func (x gen_NSThread) Start() {
	C.NSThread_inst_Start(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// IsExecuting returns a Boolean value that indicates whether the receiver is executing.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1411240-executing?language=objc for details.
func (x gen_NSThread) IsExecuting() bool {
	ret := C.NSThread_inst_IsExecuting(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsFinished returns a Boolean value that indicates whether the receiver has finished execution.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1409297-finished?language=objc for details.
func (x gen_NSThread) IsFinished() bool {
	ret := C.NSThread_inst_IsFinished(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsCancelled returns a Boolean value that indicates whether the receiver is cancelled.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1417366-cancelled?language=objc for details.
func (x gen_NSThread) IsCancelled() bool {
	ret := C.NSThread_inst_IsCancelled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// IsMainThread returns a Boolean value that indicates whether the receiver is the main thread.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1408455-ismainthread?language=objc for details.
func (x gen_NSThread) IsMainThread() bool {
	ret := C.NSThread_inst_IsMainThread(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Name returns the name of the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1414122-name?language=objc for details.
func (x gen_NSThread) Name() NSString {
	ret := C.NSThread_inst_Name(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// SetName returns the name of the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1414122-name?language=objc for details.
func (x gen_NSThread) SetName(
	value NSStringRef,
) {
	C.NSThread_inst_SetName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// StackSize returns the stack size of the receiver, in bytes.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1415190-stacksize?language=objc for details.
func (x gen_NSThread) StackSize() NSUInteger {
	ret := C.NSThread_inst_StackSize(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)
}

// SetStackSize returns the stack size of the receiver, in bytes.
//
// See https://developer.apple.com/documentation/foundation/nsthread/1415190-stacksize?language=objc for details.
func (x gen_NSThread) SetStackSize(
	value NSUInteger,
) {
	C.NSThread_inst_SetStackSize(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)

	return
}

type NSURLRef interface {
	Pointer() uintptr
	Init_AsNSURL() NSURL
}

type gen_NSURL struct {
	objc.Object
}

func NSURL_FromPointer(ptr unsafe.Pointer) NSURL {
	return NSURL{gen_NSURL{
		objc.Object_FromPointer(ptr),
	}}
}

func NSURL_FromRef(ref objc.Ref) NSURL {
	return NSURL_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// URLByAppendingPathComponent returns a new URL made by appending a path component to the original URL.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1410614-urlbyappendingpathcomponent?language=objc for details.
func (x gen_NSURL) URLByAppendingPathComponent(
	pathComponent NSStringRef,
) NSURL {
	ret := C.NSURL_inst_URLByAppendingPathComponent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pathComponent),
	)

	return NSURL_FromPointer(ret)
}

// URLByAppendingPathComponentIsDirectory returns a new URL made by appending a path component to the original URL, along with a trailing slash if the component is designated a directory.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1413953-urlbyappendingpathcomponent?language=objc for details.
func (x gen_NSURL) URLByAppendingPathComponentIsDirectory(
	pathComponent NSStringRef,
	isDirectory bool,
) NSURL {
	ret := C.NSURL_inst_URLByAppendingPathComponentIsDirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pathComponent),
		convertToObjCBool(isDirectory),
	)

	return NSURL_FromPointer(ret)
}

// URLByAppendingPathExtension returns a new URL made by appending a path extension to the original URL.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1417082-urlbyappendingpathextension?language=objc for details.
func (x gen_NSURL) URLByAppendingPathExtension(
	pathExtension NSStringRef,
) NSURL {
	ret := C.NSURL_inst_URLByAppendingPathExtension(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pathExtension),
	)

	return NSURL_FromPointer(ret)
}

// FileReferenceURL returns a new file reference URL that points to the same resource as the receiver.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1408631-filereferenceurl?language=objc for details.
func (x gen_NSURL) FileReferenceURL() NSURL {
	ret := C.NSURL_inst_FileReferenceURL(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_FromPointer(ret)
}

// InitAbsoluteURLWithDataRepresentationRelativeToURL is undocumented.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1410750-initabsoluteurlwithdatarepresent?language=objc for details.
func (x gen_NSURL) InitAbsoluteURLWithDataRepresentationRelativeToURL(
	data NSDataRef,
	baseURL NSURLRef,
) NSURL {
	ret := C.NSURL_inst_InitAbsoluteURLWithDataRepresentationRelativeToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(baseURL),
	)

	return NSURL_FromPointer(ret)
}

// InitFileURLWithPath initializes a newly created NSURL referencing the local file or directory at path.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1410301-initfileurlwithpath?language=objc for details.
func (x gen_NSURL) InitFileURLWithPath(
	path NSStringRef,
) NSURL {
	ret := C.NSURL_inst_InitFileURLWithPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)

	return NSURL_FromPointer(ret)
}

// InitFileURLWithPathIsDirectory initializes a newly created NSURL referencing the local file or directory at path.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1417505-initfileurlwithpath?language=objc for details.
func (x gen_NSURL) InitFileURLWithPathIsDirectory(
	path NSStringRef,
	isDir bool,
) NSURL {
	ret := C.NSURL_inst_InitFileURLWithPathIsDirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(isDir),
	)

	return NSURL_FromPointer(ret)
}

// InitFileURLWithPathIsDirectoryRelativeToURL is undocumented.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1417932-initfileurlwithpath?language=objc for details.
func (x gen_NSURL) InitFileURLWithPathIsDirectoryRelativeToURL(
	path NSStringRef,
	isDir bool,
	baseURL NSURLRef,
) NSURL {
	ret := C.NSURL_inst_InitFileURLWithPathIsDirectoryRelativeToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(isDir),
		objc.RefPointer(baseURL),
	)

	return NSURL_FromPointer(ret)
}

// InitFileURLWithPathRelativeToURL is undocumented.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1415077-initfileurlwithpath?language=objc for details.
func (x gen_NSURL) InitFileURLWithPathRelativeToURL(
	path NSStringRef,
	baseURL NSURLRef,
) NSURL {
	ret := C.NSURL_inst_InitFileURLWithPathRelativeToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		objc.RefPointer(baseURL),
	)

	return NSURL_FromPointer(ret)
}

// InitWithDataRepresentationRelativeToURL is undocumented.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1416851-initwithdatarepresentation?language=objc for details.
func (x gen_NSURL) InitWithDataRepresentationRelativeToURL(
	data NSDataRef,
	baseURL NSURLRef,
) NSURL {
	ret := C.NSURL_inst_InitWithDataRepresentationRelativeToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(baseURL),
	)

	return NSURL_FromPointer(ret)
}

// InitWithString initializes an NSURL object with a provided URL string.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1413146-initwithstring?language=objc for details.
func (x gen_NSURL) InitWithString(
	URLString NSStringRef,
) NSURL {
	ret := C.NSURL_inst_InitWithString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(URLString),
	)

	return NSURL_FromPointer(ret)
}

// InitWithStringRelativeToURL initializes an NSURL object with a base URL and a relative string.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1417949-initwithstring?language=objc for details.
func (x gen_NSURL) InitWithStringRelativeToURL(
	URLString NSStringRef,
	baseURL NSURLRef,
) NSURL {
	ret := C.NSURL_inst_InitWithStringRelativeToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(URLString),
		objc.RefPointer(baseURL),
	)

	return NSURL_FromPointer(ret)
}

// IsFileReferenceURL returns whether the URL is a file reference URL.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1408507-isfilereferenceurl?language=objc for details.
func (x gen_NSURL) IsFileReferenceURL() bool {
	ret := C.NSURL_inst_IsFileReferenceURL(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// RemoveAllCachedResourceValues removes all cached resource values and temporary resource values from the URL object.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1417078-removeallcachedresourcevalues?language=objc for details.
func (x gen_NSURL) RemoveAllCachedResourceValues() {
	C.NSURL_inst_RemoveAllCachedResourceValues(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// StartAccessingSecurityScopedResource in an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1417051-startaccessingsecurityscopedreso?language=objc for details.
func (x gen_NSURL) StartAccessingSecurityScopedResource() bool {
	ret := C.NSURL_inst_StartAccessingSecurityScopedResource(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// StopAccessingSecurityScopedResource in an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1413736-stopaccessingsecurityscopedresou?language=objc for details.
func (x gen_NSURL) StopAccessingSecurityScopedResource() {
	C.NSURL_inst_StopAccessingSecurityScopedResource(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Init initializes a new instance of the NSURL class.
func (x gen_NSURL) Init() NSURL {
	ret := C.NSURL_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_FromPointer(ret)
}

// Init_AsNSURL is a typed version of Init.
func (x gen_NSURL) Init_AsNSURL() NSURL {
	ret := C.NSURL_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_FromPointer(ret)
}

// DataRepresentation is undocumented.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1407656-datarepresentation?language=objc for details.
func (x gen_NSURL) DataRepresentation() NSData {
	ret := C.NSURL_inst_DataRepresentation(
		unsafe.Pointer(x.Pointer()),
	)

	return NSData_FromPointer(ret)
}

// IsFileURL returns a boolean value that determines whether the receiver is a file URL.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1408782-fileurl?language=objc for details.
func (x gen_NSURL) IsFileURL() bool {
	ret := C.NSURL_inst_IsFileURL(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AbsoluteString returns the URL string for the receiver as an absolute URL. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1409868-absolutestring?language=objc for details.
func (x gen_NSURL) AbsoluteString() NSString {
	ret := C.NSURL_inst_AbsoluteString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// AbsoluteURL an absolute URL that refers to the same resource as the receiver. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1414266-absoluteurl?language=objc for details.
func (x gen_NSURL) AbsoluteURL() NSURL {
	ret := C.NSURL_inst_AbsoluteURL(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_FromPointer(ret)
}

// BaseURL returns the base URL. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1412311-baseurl?language=objc for details.
func (x gen_NSURL) BaseURL() NSURL {
	ret := C.NSURL_inst_BaseURL(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_FromPointer(ret)
}

// Fragment returns the fragment identifier, conforming to RFC 1808. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1413775-fragment?language=objc for details.
func (x gen_NSURL) Fragment() NSString {
	ret := C.NSURL_inst_Fragment(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// Host returns the host, conforming to RFC 1808. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1413640-host?language=objc for details.
func (x gen_NSURL) Host() NSString {
	ret := C.NSURL_inst_Host(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// LastPathComponent returns the last path component. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1417444-lastpathcomponent?language=objc for details.
func (x gen_NSURL) LastPathComponent() NSString {
	ret := C.NSURL_inst_LastPathComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// Password returns the password conforming to RFC 1808. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1412096-password?language=objc for details.
func (x gen_NSURL) Password() NSString {
	ret := C.NSURL_inst_Password(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// Path returns the path, conforming to RFC 1808. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1408809-path?language=objc for details.
func (x gen_NSURL) Path() NSString {
	ret := C.NSURL_inst_Path(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// PathComponents an array containing the path components. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1407365-pathcomponents?language=objc for details.
func (x gen_NSURL) PathComponents() NSArray {
	ret := C.NSURL_inst_PathComponents(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}

// PathExtension returns the path extension. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1410208-pathextension?language=objc for details.
func (x gen_NSURL) PathExtension() NSString {
	ret := C.NSURL_inst_PathExtension(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// Port returns the port, conforming to RFC 1808.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1413455-port?language=objc for details.
func (x gen_NSURL) Port() NSNumber {
	ret := C.NSURL_inst_Port(
		unsafe.Pointer(x.Pointer()),
	)

	return NSNumber_FromPointer(ret)
}

// Query returns the query string, conforming to RFC 1808.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1407543-query?language=objc for details.
func (x gen_NSURL) Query() NSString {
	ret := C.NSURL_inst_Query(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// RelativePath returns the relative path, conforming to RFC 1808. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1410263-relativepath?language=objc for details.
func (x gen_NSURL) RelativePath() NSString {
	ret := C.NSURL_inst_RelativePath(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// RelativeString returns a string representation of the relative portion of the URL. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1411417-relativestring?language=objc for details.
func (x gen_NSURL) RelativeString() NSString {
	ret := C.NSURL_inst_RelativeString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// ResourceSpecifier returns the resource specifier. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1415309-resourcespecifier?language=objc for details.
func (x gen_NSURL) ResourceSpecifier() NSString {
	ret := C.NSURL_inst_ResourceSpecifier(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// Scheme returns the scheme. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1413437-scheme?language=objc for details.
func (x gen_NSURL) Scheme() NSString {
	ret := C.NSURL_inst_Scheme(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// StandardizedURL returns a copy of the URL with any instances of ".." or "." removed from its path. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1411073-standardizedurl?language=objc for details.
func (x gen_NSURL) StandardizedURL() NSURL {
	ret := C.NSURL_inst_StandardizedURL(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_FromPointer(ret)
}

// User returns the user name, conforming to RFC 1808.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1418335-user?language=objc for details.
func (x gen_NSURL) User() NSString {
	ret := C.NSURL_inst_User(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// FilePathURL returns a file path URL that points to the same resource as the URL object. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1408442-filepathurl?language=objc for details.
func (x gen_NSURL) FilePathURL() NSURL {
	ret := C.NSURL_inst_FilePathURL(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_FromPointer(ret)
}

// URLByDeletingLastPathComponent returns a URL created by taking the receiver and removing the last path component. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1411592-urlbydeletinglastpathcomponent?language=objc for details.
func (x gen_NSURL) URLByDeletingLastPathComponent() NSURL {
	ret := C.NSURL_inst_URLByDeletingLastPathComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_FromPointer(ret)
}

// URLByDeletingPathExtension returns a URL created by taking the receiver and removing the path extension, if any. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1412357-urlbydeletingpathextension?language=objc for details.
func (x gen_NSURL) URLByDeletingPathExtension() NSURL {
	ret := C.NSURL_inst_URLByDeletingPathExtension(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_FromPointer(ret)
}

// URLByResolvingSymlinksInPath returns a URL that points to the same resource as the receiver and includes no symbolic links. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1415965-urlbyresolvingsymlinksinpath?language=objc for details.
func (x gen_NSURL) URLByResolvingSymlinksInPath() NSURL {
	ret := C.NSURL_inst_URLByResolvingSymlinksInPath(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_FromPointer(ret)
}

// URLByStandardizingPath returns a URL that points to the same resource as the original URL using an absolute path. (read-only)
//
// See https://developer.apple.com/documentation/foundation/nsurl/1414302-urlbystandardizingpath?language=objc for details.
func (x gen_NSURL) URLByStandardizingPath() NSURL {
	ret := C.NSURL_inst_URLByStandardizingPath(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_FromPointer(ret)
}

// HasDirectoryPath is undocumented.
//
// See https://developer.apple.com/documentation/foundation/nsurl/1411475-hasdirectorypath?language=objc for details.
func (x gen_NSURL) HasDirectoryPath() bool {
	ret := C.NSURL_inst_HasDirectoryPath(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

type NSURLRequestRef interface {
	Pointer() uintptr
	Init_AsNSURLRequest() NSURLRequest
}

type gen_NSURLRequest struct {
	objc.Object
}

func NSURLRequest_FromPointer(ptr unsafe.Pointer) NSURLRequest {
	return NSURLRequest{gen_NSURLRequest{
		objc.Object_FromPointer(ptr),
	}}
}

func NSURLRequest_FromRef(ref objc.Ref) NSURLRequest {
	return NSURLRequest_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// InitWithURL creates a URL request for a specified URL.
//
// See https://developer.apple.com/documentation/foundation/nsurlrequest/1410303-initwithurl?language=objc for details.
func (x gen_NSURLRequest) InitWithURL(
	URL NSURLRef,
) NSURLRequest {
	ret := C.NSURLRequest_inst_InitWithURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(URL),
	)

	return NSURLRequest_FromPointer(ret)
}

// ValueForHTTPHeaderField returns the value of the specified HTTP header field.
//
// See https://developer.apple.com/documentation/foundation/nsurlrequest/1409376-valueforhttpheaderfield?language=objc for details.
func (x gen_NSURLRequest) ValueForHTTPHeaderField(
	field NSStringRef,
) NSString {
	ret := C.NSURLRequest_inst_ValueForHTTPHeaderField(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(field),
	)

	return NSString_FromPointer(ret)
}

// Init initializes a new instance of the NSURLRequest class.
func (x gen_NSURLRequest) Init() NSURLRequest {
	ret := C.NSURLRequest_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURLRequest_FromPointer(ret)
}

// Init_AsNSURLRequest is a typed version of Init.
func (x gen_NSURLRequest) Init_AsNSURLRequest() NSURLRequest {
	ret := C.NSURLRequest_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURLRequest_FromPointer(ret)
}

// HTTPMethod returns the HTTP request method.
//
// See https://developer.apple.com/documentation/foundation/nsurlrequest/1413030-httpmethod?language=objc for details.
func (x gen_NSURLRequest) HTTPMethod() NSString {
	ret := C.NSURLRequest_inst_HTTPMethod(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_FromPointer(ret)
}

// URL returns the URL being requested.
//
// See https://developer.apple.com/documentation/foundation/nsurlrequest/1408996-url?language=objc for details.
func (x gen_NSURLRequest) URL() NSURL {
	ret := C.NSURLRequest_inst_URL(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_FromPointer(ret)
}

// HTTPBody returns the request body.
//
// See https://developer.apple.com/documentation/foundation/nsurlrequest/1411317-httpbody?language=objc for details.
func (x gen_NSURLRequest) HTTPBody() NSData {
	ret := C.NSURLRequest_inst_HTTPBody(
		unsafe.Pointer(x.Pointer()),
	)

	return NSData_FromPointer(ret)
}

// MainDocumentURL returns the main document URL associated with the request.
//
// See https://developer.apple.com/documentation/foundation/nsurlrequest/1414134-maindocumenturl?language=objc for details.
func (x gen_NSURLRequest) MainDocumentURL() NSURL {
	ret := C.NSURLRequest_inst_MainDocumentURL(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_FromPointer(ret)
}

// AllHTTPHeaderFields returns a dictionary containing all of the HTTP header fields for a request.
//
// See https://developer.apple.com/documentation/foundation/nsurlrequest/1418477-allhttpheaderfields?language=objc for details.
func (x gen_NSURLRequest) AllHTTPHeaderFields() NSDictionary {
	ret := C.NSURLRequest_inst_AllHTTPHeaderFields(
		unsafe.Pointer(x.Pointer()),
	)

	return NSDictionary_FromPointer(ret)
}

// HTTPShouldHandleCookies returns a Boolean value that indicates whether the default cookie handling will be used for this request.
//
// See https://developer.apple.com/documentation/foundation/nsurlrequest/1418369-httpshouldhandlecookies?language=objc for details.
func (x gen_NSURLRequest) HTTPShouldHandleCookies() bool {
	ret := C.NSURLRequest_inst_HTTPShouldHandleCookies(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// HTTPShouldUsePipelining returns a Boolean value that indicates whether the request should continue transmitting data before receiving a response from an earlier transmission.
//
// See https://developer.apple.com/documentation/foundation/nsurlrequest/1409170-httpshouldusepipelining?language=objc for details.
func (x gen_NSURLRequest) HTTPShouldUsePipelining() bool {
	ret := C.NSURLRequest_inst_HTTPShouldUsePipelining(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AllowsCellularAccess returns a Boolean value that indicates whether the request is allowed to use the cellular radio (if present).
//
// See https://developer.apple.com/documentation/foundation/nsurlrequest/1412032-allowscellularaccess?language=objc for details.
func (x gen_NSURLRequest) AllowsCellularAccess() bool {
	ret := C.NSURLRequest_inst_AllowsCellularAccess(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AllowsConstrainedNetworkAccess returns a Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// See https://developer.apple.com/documentation/foundation/nsurlrequest/3325678-allowsconstrainednetworkaccess?language=objc for details.
func (x gen_NSURLRequest) AllowsConstrainedNetworkAccess() bool {
	ret := C.NSURLRequest_inst_AllowsConstrainedNetworkAccess(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AllowsExpensiveNetworkAccess returns a Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// See https://developer.apple.com/documentation/foundation/nsurlrequest/3325679-allowsexpensivenetworkaccess?language=objc for details.
func (x gen_NSURLRequest) AllowsExpensiveNetworkAccess() bool {
	ret := C.NSURLRequest_inst_AllowsExpensiveNetworkAccess(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AssumesHTTP3Capable is undocumented.
//
// See https://developer.apple.com/documentation/foundation/nsurlrequest/3735880-assumeshttp3capable?language=objc for details.
func (x gen_NSURLRequest) AssumesHTTP3Capable() bool {
	ret := C.NSURLRequest_inst_AssumesHTTP3Capable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

type NSUserDefaultsRef interface {
	Pointer() uintptr
	Init_AsNSUserDefaults() NSUserDefaults
}

type gen_NSUserDefaults struct {
	objc.Object
}

func NSUserDefaults_FromPointer(ptr unsafe.Pointer) NSUserDefaults {
	return NSUserDefaults{gen_NSUserDefaults{
		objc.Object_FromPointer(ptr),
	}}
}

func NSUserDefaults_FromRef(ref objc.Ref) NSUserDefaults {
	return NSUserDefaults_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// URLForKey returns the URL associated with the specified key.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1408648-urlforkey?language=objc for details.
func (x gen_NSUserDefaults) URLForKey(
	defaultName NSStringRef,
) NSURL {
	ret := C.NSUserDefaults_inst_URLForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return NSURL_FromPointer(ret)
}

// AddSuiteNamed inserts the specified domain name into the receiver’s search list.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1410294-addsuitenamed?language=objc for details.
func (x gen_NSUserDefaults) AddSuiteNamed(
	suiteName NSStringRef,
) {
	C.NSUserDefaults_inst_AddSuiteNamed(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(suiteName),
	)

	return
}

// ArrayForKey returns the array associated with the specified key.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1414792-arrayforkey?language=objc for details.
func (x gen_NSUserDefaults) ArrayForKey(
	defaultName NSStringRef,
) NSArray {
	ret := C.NSUserDefaults_inst_ArrayForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return NSArray_FromPointer(ret)
}

// BoolForKey returns the Boolean value associated with the specified key.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1416388-boolforkey?language=objc for details.
func (x gen_NSUserDefaults) BoolForKey(
	defaultName NSStringRef,
) bool {
	ret := C.NSUserDefaults_inst_BoolForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return convertObjCBoolToGo(ret)
}

// DataForKey returns the data object associated with the specified key.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1409590-dataforkey?language=objc for details.
func (x gen_NSUserDefaults) DataForKey(
	defaultName NSStringRef,
) NSData {
	ret := C.NSUserDefaults_inst_DataForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return NSData_FromPointer(ret)
}

// DictionaryForKey returns the dictionary object associated with the specified key.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1408563-dictionaryforkey?language=objc for details.
func (x gen_NSUserDefaults) DictionaryForKey(
	defaultName NSStringRef,
) NSDictionary {
	ret := C.NSUserDefaults_inst_DictionaryForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return NSDictionary_FromPointer(ret)
}

// DictionaryRepresentation returns a dictionary that contains a union of all key-value pairs in the domains in the search list.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1415919-dictionaryrepresentation?language=objc for details.
func (x gen_NSUserDefaults) DictionaryRepresentation() NSDictionary {
	ret := C.NSUserDefaults_inst_DictionaryRepresentation(
		unsafe.Pointer(x.Pointer()),
	)

	return NSDictionary_FromPointer(ret)
}

// Init creates a user defaults object initialized with the defaults for the app and current user.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1414356-init?language=objc for details.
func (x gen_NSUserDefaults) Init() NSUserDefaults {
	ret := C.NSUserDefaults_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUserDefaults_FromPointer(ret)
}

// Init_AsNSUserDefaults is a typed version of Init.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1414356-init?language=objc for details.
func (x gen_NSUserDefaults) Init_AsNSUserDefaults() NSUserDefaults {
	ret := C.NSUserDefaults_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUserDefaults_FromPointer(ret)
}

// InitWithSuiteName creates a user defaults object initialized with the defaults for the specified database name.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1409957-initwithsuitename?language=objc for details.
func (x gen_NSUserDefaults) InitWithSuiteName(
	suitename NSStringRef,
) NSUserDefaults {
	ret := C.NSUserDefaults_inst_InitWithSuiteName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(suitename),
	)

	return NSUserDefaults_FromPointer(ret)
}

// IntegerForKey returns the integer value associated with the specified key.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1407405-integerforkey?language=objc for details.
func (x gen_NSUserDefaults) IntegerForKey(
	defaultName NSStringRef,
) NSInteger {
	ret := C.NSUserDefaults_inst_IntegerForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return NSInteger(ret)
}

// ObjectForKey returns the object associated with the specified key.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1410095-objectforkey?language=objc for details.
func (x gen_NSUserDefaults) ObjectForKey(
	defaultName NSStringRef,
) objc.Object {
	ret := C.NSUserDefaults_inst_ObjectForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return objc.Object_FromPointer(ret)
}

// ObjectIsForcedForKey returns a Boolean value indicating whether the specified key is managed by an administrator.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1408635-objectisforcedforkey?language=objc for details.
func (x gen_NSUserDefaults) ObjectIsForcedForKey(
	key NSStringRef,
) bool {
	ret := C.NSUserDefaults_inst_ObjectIsForcedForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)

	return convertObjCBoolToGo(ret)
}

// ObjectIsForcedForKeyInDomain returns a Boolean value indicating whether the key in the specified domain is managed by an administrator.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1416306-objectisforcedforkey?language=objc for details.
func (x gen_NSUserDefaults) ObjectIsForcedForKeyInDomain(
	key NSStringRef,
	domain NSStringRef,
) bool {
	ret := C.NSUserDefaults_inst_ObjectIsForcedForKeyInDomain(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
		objc.RefPointer(domain),
	)

	return convertObjCBoolToGo(ret)
}

// PersistentDomainForName returns a dictionary representation of the defaults for the specified domain.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1412197-persistentdomainforname?language=objc for details.
func (x gen_NSUserDefaults) PersistentDomainForName(
	domainName NSStringRef,
) NSDictionary {
	ret := C.NSUserDefaults_inst_PersistentDomainForName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domainName),
	)

	return NSDictionary_FromPointer(ret)
}

// RegisterDefaults adds the contents of the specified dictionary to the registration domain.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1417065-registerdefaults?language=objc for details.
func (x gen_NSUserDefaults) RegisterDefaults(
	registrationDictionary NSDictionaryRef,
) {
	C.NSUserDefaults_inst_RegisterDefaults(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(registrationDictionary),
	)

	return
}

// RemoveObjectForKey removes the value of the specified default key.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1411182-removeobjectforkey?language=objc for details.
func (x gen_NSUserDefaults) RemoveObjectForKey(
	defaultName NSStringRef,
) {
	C.NSUserDefaults_inst_RemoveObjectForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return
}

// RemovePersistentDomainForName removes the contents of the specified persistent domain from the user’s defaults.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1417339-removepersistentdomainforname?language=objc for details.
func (x gen_NSUserDefaults) RemovePersistentDomainForName(
	domainName NSStringRef,
) {
	C.NSUserDefaults_inst_RemovePersistentDomainForName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domainName),
	)

	return
}

// RemoveSuiteNamed removes the specified domain name from the receiver’s search list.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1408047-removesuitenamed?language=objc for details.
func (x gen_NSUserDefaults) RemoveSuiteNamed(
	suiteName NSStringRef,
) {
	C.NSUserDefaults_inst_RemoveSuiteNamed(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(suiteName),
	)

	return
}

// RemoveVolatileDomainForName removes the specified volatile domain from the user’s defaults.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1415955-removevolatiledomainforname?language=objc for details.
func (x gen_NSUserDefaults) RemoveVolatileDomainForName(
	domainName NSStringRef,
) {
	C.NSUserDefaults_inst_RemoveVolatileDomainForName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domainName),
	)

	return
}

// SetBoolForKey sets the value of the specified default key to the specified Boolean value.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1408905-setbool?language=objc for details.
func (x gen_NSUserDefaults) SetBoolForKey(
	value bool,
	defaultName NSStringRef,
) {
	C.NSUserDefaults_inst_SetBoolForKey(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
		objc.RefPointer(defaultName),
	)

	return
}

// SetIntegerForKey sets the value of the specified default key to the specified integer value.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1413614-setinteger?language=objc for details.
func (x gen_NSUserDefaults) SetIntegerForKey(
	value NSInteger,
	defaultName NSStringRef,
) {
	C.NSUserDefaults_inst_SetIntegerForKey(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
		objc.RefPointer(defaultName),
	)

	return
}

// SetObjectForKey sets the value of the specified default key.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1414067-setobject?language=objc for details.
func (x gen_NSUserDefaults) SetObjectForKey(
	value objc.Ref,
	defaultName NSStringRef,
) {
	C.NSUserDefaults_inst_SetObjectForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
		objc.RefPointer(defaultName),
	)

	return
}

// SetPersistentDomainForName sets a dictionary for the specified persistent domain.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1408187-setpersistentdomain?language=objc for details.
func (x gen_NSUserDefaults) SetPersistentDomainForName(
	domain NSDictionaryRef,
	domainName NSStringRef,
) {
	C.NSUserDefaults_inst_SetPersistentDomainForName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domain),
		objc.RefPointer(domainName),
	)

	return
}

// SetURLForKey sets the value of the specified default key to the specified URL.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1414194-seturl?language=objc for details.
func (x gen_NSUserDefaults) SetURLForKey(
	url NSURLRef,
	defaultName NSStringRef,
) {
	C.NSUserDefaults_inst_SetURLForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
		objc.RefPointer(defaultName),
	)

	return
}

// SetVolatileDomainForName sets the dictionary for the specified volatile domain.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1413720-setvolatiledomain?language=objc for details.
func (x gen_NSUserDefaults) SetVolatileDomainForName(
	domain NSDictionaryRef,
	domainName NSStringRef,
) {
	C.NSUserDefaults_inst_SetVolatileDomainForName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domain),
		objc.RefPointer(domainName),
	)

	return
}

// StringArrayForKey returns the array of strings associated with the specified key.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1416414-stringarrayforkey?language=objc for details.
func (x gen_NSUserDefaults) StringArrayForKey(
	defaultName NSStringRef,
) NSArray {
	ret := C.NSUserDefaults_inst_StringArrayForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return NSArray_FromPointer(ret)
}

// StringForKey returns the string associated with the specified key.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1416700-stringforkey?language=objc for details.
func (x gen_NSUserDefaults) StringForKey(
	defaultName NSStringRef,
) NSString {
	ret := C.NSUserDefaults_inst_StringForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return NSString_FromPointer(ret)
}

// Synchronize waits for any pending asynchronous updates to the defaults database and returns; this method is unnecessary and shouldn't be used.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1414005-synchronize?language=objc for details.
func (x gen_NSUserDefaults) Synchronize() bool {
	ret := C.NSUserDefaults_inst_Synchronize(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// VolatileDomainForName returns the dictionary for the specified volatile domain.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1409592-volatiledomainforname?language=objc for details.
func (x gen_NSUserDefaults) VolatileDomainForName(
	domainName NSStringRef,
) NSDictionary {
	ret := C.NSUserDefaults_inst_VolatileDomainForName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domainName),
	)

	return NSDictionary_FromPointer(ret)
}

// VolatileDomainNames returns the current volatile domain names.
//
// See https://developer.apple.com/documentation/foundation/nsuserdefaults/1414231-volatiledomainnames?language=objc for details.
func (x gen_NSUserDefaults) VolatileDomainNames() NSArray {
	ret := C.NSUserDefaults_inst_VolatileDomainNames(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_FromPointer(ret)
}
