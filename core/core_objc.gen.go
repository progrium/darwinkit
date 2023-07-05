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


void* CALayer_type_alloc() {
	return [CALayer
		alloc];
}
void* CALayer_type_layer() {
	return [CALayer
		layer];
}
BOOL CALayer_type_needsDisplayForKey(void* key) {
	return [CALayer
		needsDisplayForKey: key];
}
void* CALayer_type_defaultActionForKey(void* event) {
	return [CALayer
		defaultActionForKey: event];
}
void* CALayer_type_defaultValueForKey(void* key) {
	return [CALayer
		defaultValueForKey: key];
}
void* NSArray_type_alloc() {
	return [NSArray
		alloc];
}
void* NSArray_type_array() {
	return [NSArray
		array];
}
void* NSArray_type_arrayWithArray(void* array) {
	return [NSArray
		arrayWithArray: array];
}
void* NSAttributedString_type_alloc() {
	return [NSAttributedString
		alloc];
}
void* NSAttributedString_type_textTypes() {
	return [NSAttributedString
		textTypes];
}
void* NSAttributedString_type_textUnfilteredTypes() {
	return [NSAttributedString
		textUnfilteredTypes];
}
void* NSData_type_alloc() {
	return [NSData
		alloc];
}
void* NSData_type_data() {
	return [NSData
		data];
}
void* NSData_type_dataWithBytes_length(void* bytes, unsigned long length) {
	return [NSData
		dataWithBytes: bytes
		length: length];
}
void* NSData_type_dataWithBytesNoCopy_length(void* bytes, unsigned long length) {
	return [NSData
		dataWithBytesNoCopy: bytes
		length: length];
}
void* NSData_type_dataWithBytesNoCopy_length_freeWhenDone(void* bytes, unsigned long length, BOOL b) {
	return [NSData
		dataWithBytesNoCopy: bytes
		length: length
		freeWhenDone: b];
}
void* NSData_type_dataWithData(void* data) {
	return [NSData
		dataWithData: data];
}
void* NSData_type_dataWithContentsOfFile(void* path) {
	return [NSData
		dataWithContentsOfFile: path];
}
void* NSData_type_dataWithContentsOfURL(void* url) {
	return [NSData
		dataWithContentsOfURL: url];
}
void* NSDictionary_type_alloc() {
	return [NSDictionary
		alloc];
}
void* NSDictionary_type_dictionary() {
	return [NSDictionary
		dictionary];
}
void* NSDictionary_type_dictionaryWithObjects_forKeys(void* objects, void* keys) {
	return [NSDictionary
		dictionaryWithObjects: objects
		forKeys: keys];
}
void* NSDictionary_type_dictionaryWithDictionary(void* dict) {
	return [NSDictionary
		dictionaryWithDictionary: dict];
}
void* NSDictionary_type_sharedKeySetForKeys(void* keys) {
	return [NSDictionary
		sharedKeySetForKeys: keys];
}
void* NSNumber_type_alloc() {
	return [NSNumber
		alloc];
}
void* NSNumber_type_numberWithBool(BOOL value) {
	return [NSNumber
		numberWithBool: value];
}
void* NSNumber_type_numberWithInt(int value) {
	return [NSNumber
		numberWithInt: value];
}
void* NSNumber_type_numberWithInteger(long value) {
	return [NSNumber
		numberWithInteger: value];
}
void* NSNumber_type_numberWithUnsignedInt(int value) {
	return [NSNumber
		numberWithUnsignedInt: value];
}
void* NSNumber_type_numberWithUnsignedInteger(unsigned long value) {
	return [NSNumber
		numberWithUnsignedInteger: value];
}
void* NSRunLoop_type_alloc() {
	return [NSRunLoop
		alloc];
}
void* NSRunLoop_type_currentRunLoop() {
	return [NSRunLoop
		currentRunLoop];
}
void* NSRunLoop_type_mainRunLoop() {
	return [NSRunLoop
		mainRunLoop];
}
void* NSString_type_alloc() {
	return [NSString
		alloc];
}
void* NSString_type_string() {
	return [NSString
		string];
}
void* NSString_type_localizedUserNotificationStringForKey_arguments(void* key, void* arguments) {
	return [NSString
		localizedUserNotificationStringForKey: key
		arguments: arguments];
}
void* NSString_type_stringWithString(void* string) {
	return [NSString
		stringWithString: string];
}
void* NSString_type_localizedNameOfStringEncoding(unsigned long encoding) {
	return [NSString
		localizedNameOfStringEncoding: encoding];
}
void* NSString_type_pathWithComponents(void* components) {
	return [NSString
		pathWithComponents: components];
}
unsigned long NSString_type_defaultCStringEncoding() {
	return [NSString
		defaultCStringEncoding];
}
void* NSThread_type_alloc() {
	return [NSThread
		alloc];
}
void NSThread_type_detachNewThreadSelector_toTarget_withObject(void* selector, void* target, void* argument) {
	[NSThread
		detachNewThreadSelector: selector
		toTarget: target
		withObject: argument];
}
void NSThread_type_exit() {
	[NSThread
		exit];
}
BOOL NSThread_type_isMultiThreaded() {
	return [NSThread
		isMultiThreaded];
}
BOOL NSThread_type_isMainThread() {
	return [NSThread
		isMainThread];
}
void* NSThread_type_mainThread() {
	return [NSThread
		mainThread];
}
void* NSThread_type_currentThread() {
	return [NSThread
		currentThread];
}
void* NSThread_type_callStackReturnAddresses() {
	return [NSThread
		callStackReturnAddresses];
}
void* NSThread_type_callStackSymbols() {
	return [NSThread
		callStackSymbols];
}
void* NSURL_type_alloc() {
	return [NSURL
		alloc];
}
void* NSURL_type_URLWithString(void* URLString) {
	return [NSURL
		URLWithString: URLString];
}
void* NSURL_type_URLWithString_relativeToURL(void* URLString, void* baseURL) {
	return [NSURL
		URLWithString: URLString
		relativeToURL: baseURL];
}
void* NSURL_type_fileURLWithPath_isDirectory(void* path, BOOL isDir) {
	return [NSURL
		fileURLWithPath: path
		isDirectory: isDir];
}
void* NSURL_type_fileURLWithPath_relativeToURL(void* path, void* baseURL) {
	return [NSURL
		fileURLWithPath: path
		relativeToURL: baseURL];
}
void* NSURL_type_fileURLWithPath_isDirectory_relativeToURL(void* path, BOOL isDir, void* baseURL) {
	return [NSURL
		fileURLWithPath: path
		isDirectory: isDir
		relativeToURL: baseURL];
}
void* NSURL_type_fileURLWithPath(void* path) {
	return [NSURL
		fileURLWithPath: path];
}
void* NSURL_type_fileURLWithPathComponents(void* components) {
	return [NSURL
		fileURLWithPathComponents: components];
}
void* NSURL_type_absoluteURLWithDataRepresentation_relativeToURL(void* data, void* baseURL) {
	return [NSURL
		absoluteURLWithDataRepresentation: data
		relativeToURL: baseURL];
}
void* NSURL_type_URLWithDataRepresentation_relativeToURL(void* data, void* baseURL) {
	return [NSURL
		URLWithDataRepresentation: data
		relativeToURL: baseURL];
}
void* NSURL_type_resourceValuesForKeys_fromBookmarkData(void* keys, void* bookmarkData) {
	return [NSURL
		resourceValuesForKeys: keys
		fromBookmarkData: bookmarkData];
}
void* NSURLRequest_type_alloc() {
	return [NSURLRequest
		alloc];
}
void* NSURLRequest_type_requestWithURL(void* URL) {
	return [NSURLRequest
		requestWithURL: URL];
}
BOOL NSURLRequest_type_supportsSecureCoding() {
	return [NSURLRequest
		supportsSecureCoding];
}
void* NSUserDefaults_type_alloc() {
	return [NSUserDefaults
		alloc];
}
void NSUserDefaults_type_resetStandardUserDefaults() {
	[NSUserDefaults
		resetStandardUserDefaults];
}
void* NSUserDefaults_type_standardUserDefaults() {
	return [NSUserDefaults
		standardUserDefaults];
}


void* CALayer_inst_actionForKey(void *id, void* event) {
	return [(CALayer*)id
		actionForKey: event];
}

void CALayer_inst_addSublayer(void *id, void* layer) {
	[(CALayer*)id
		addSublayer: layer];
}

void* CALayer_inst_animationKeys(void *id) {
	return [(CALayer*)id
		animationKeys];
}

BOOL CALayer_inst_contentsAreFlipped(void *id) {
	return [(CALayer*)id
		contentsAreFlipped];
}

NSRect CALayer_inst_convertRect_fromLayer(void *id, NSRect r, void* l) {
	return [(CALayer*)id
		convertRect: r
		fromLayer: l];
}

NSRect CALayer_inst_convertRect_toLayer(void *id, NSRect r, void* l) {
	return [(CALayer*)id
		convertRect: r
		toLayer: l];
}

void CALayer_inst_display(void *id) {
	[(CALayer*)id
		display];
}

void CALayer_inst_displayIfNeeded(void *id) {
	[(CALayer*)id
		displayIfNeeded];
}

void* CALayer_inst_init(void *id) {
	return [(CALayer*)id
		init];
}

void* CALayer_inst_initWithLayer(void *id, void* layer) {
	return [(CALayer*)id
		initWithLayer: layer];
}

void CALayer_inst_insertSublayer_above(void *id, void* layer, void* sibling) {
	[(CALayer*)id
		insertSublayer: layer
		above: sibling];
}

void CALayer_inst_insertSublayer_atIndex(void *id, void* layer, int idx) {
	[(CALayer*)id
		insertSublayer: layer
		atIndex: idx];
}

void CALayer_inst_insertSublayer_below(void *id, void* layer, void* sibling) {
	[(CALayer*)id
		insertSublayer: layer
		below: sibling];
}

void CALayer_inst_layoutIfNeeded(void *id) {
	[(CALayer*)id
		layoutIfNeeded];
}

void CALayer_inst_layoutSublayers(void *id) {
	[(CALayer*)id
		layoutSublayers];
}

void* CALayer_inst_modelLayer(void *id) {
	return [(CALayer*)id
		modelLayer];
}

BOOL CALayer_inst_needsDisplay(void *id) {
	return [(CALayer*)id
		needsDisplay];
}

BOOL CALayer_inst_needsLayout(void *id) {
	return [(CALayer*)id
		needsLayout];
}

NSSize CALayer_inst_preferredFrameSize(void *id) {
	return [(CALayer*)id
		preferredFrameSize];
}

void* CALayer_inst_presentationLayer(void *id) {
	return [(CALayer*)id
		presentationLayer];
}

void CALayer_inst_removeAllAnimations(void *id) {
	[(CALayer*)id
		removeAllAnimations];
}

void CALayer_inst_removeAnimationForKey(void *id, void* key) {
	[(CALayer*)id
		removeAnimationForKey: key];
}

void CALayer_inst_removeFromSuperlayer(void *id) {
	[(CALayer*)id
		removeFromSuperlayer];
}

void CALayer_inst_replaceSublayer_with(void *id, void* oldLayer, void* newLayer) {
	[(CALayer*)id
		replaceSublayer: oldLayer
		with: newLayer];
}

void CALayer_inst_resizeSublayersWithOldSize(void *id, NSSize size) {
	[(CALayer*)id
		resizeSublayersWithOldSize: size];
}

void CALayer_inst_resizeWithOldSuperlayerSize(void *id, NSSize size) {
	[(CALayer*)id
		resizeWithOldSuperlayerSize: size];
}

void CALayer_inst_scrollRectToVisible(void *id, NSRect r) {
	[(CALayer*)id
		scrollRectToVisible: r];
}

void CALayer_inst_setNeedsDisplay(void *id) {
	[(CALayer*)id
		setNeedsDisplay];
}

void CALayer_inst_setNeedsDisplayInRect(void *id, NSRect r) {
	[(CALayer*)id
		setNeedsDisplayInRect: r];
}

void CALayer_inst_setNeedsLayout(void *id) {
	[(CALayer*)id
		setNeedsLayout];
}

BOOL CALayer_inst_shouldArchiveValueForKey(void *id, void* key) {
	return [(CALayer*)id
		shouldArchiveValueForKey: key];
}

void* CALayer_inst_delegate(void *id) {
	return [(CALayer*)id
		delegate];
}

void CALayer_inst_setDelegate(void *id, void* value) {
	[(CALayer*)id
		setDelegate: value];
}

void* CALayer_inst_contents(void *id) {
	return [(CALayer*)id
		contents];
}

void CALayer_inst_setContents(void *id, void* value) {
	[(CALayer*)id
		setContents: value];
}

NSRect CALayer_inst_contentsRect(void *id) {
	return [(CALayer*)id
		contentsRect];
}

void CALayer_inst_setContentsRect(void *id, NSRect value) {
	[(CALayer*)id
		setContentsRect: value];
}

NSRect CALayer_inst_contentsCenter(void *id) {
	return [(CALayer*)id
		contentsCenter];
}

void CALayer_inst_setContentsCenter(void *id, NSRect value) {
	[(CALayer*)id
		setContentsCenter: value];
}

BOOL CALayer_inst_isHidden(void *id) {
	return [(CALayer*)id
		isHidden];
}

void CALayer_inst_setHidden(void *id, BOOL value) {
	[(CALayer*)id
		setHidden: value];
}

BOOL CALayer_inst_masksToBounds(void *id) {
	return [(CALayer*)id
		masksToBounds];
}

void CALayer_inst_setMasksToBounds(void *id, BOOL value) {
	[(CALayer*)id
		setMasksToBounds: value];
}

void* CALayer_inst_mask(void *id) {
	return [(CALayer*)id
		mask];
}

void CALayer_inst_setMask(void *id, void* value) {
	[(CALayer*)id
		setMask: value];
}

BOOL CALayer_inst_isDoubleSided(void *id) {
	return [(CALayer*)id
		isDoubleSided];
}

void CALayer_inst_setDoubleSided(void *id, BOOL value) {
	[(CALayer*)id
		setDoubleSided: value];
}

double CALayer_inst_cornerRadius(void *id) {
	return [(CALayer*)id
		cornerRadius];
}

void CALayer_inst_setCornerRadius(void *id, double value) {
	[(CALayer*)id
		setCornerRadius: value];
}

double CALayer_inst_borderWidth(void *id) {
	return [(CALayer*)id
		borderWidth];
}

void CALayer_inst_setBorderWidth(void *id, double value) {
	[(CALayer*)id
		setBorderWidth: value];
}

double CALayer_inst_shadowRadius(void *id) {
	return [(CALayer*)id
		shadowRadius];
}

void CALayer_inst_setShadowRadius(void *id, double value) {
	[(CALayer*)id
		setShadowRadius: value];
}

NSSize CALayer_inst_shadowOffset(void *id) {
	return [(CALayer*)id
		shadowOffset];
}

void CALayer_inst_setShadowOffset(void *id, NSSize value) {
	[(CALayer*)id
		setShadowOffset: value];
}

void* CALayer_inst_style(void *id) {
	return [(CALayer*)id
		style];
}

void CALayer_inst_setStyle(void *id, void* value) {
	[(CALayer*)id
		setStyle: value];
}

BOOL CALayer_inst_allowsEdgeAntialiasing(void *id) {
	return [(CALayer*)id
		allowsEdgeAntialiasing];
}

void CALayer_inst_setAllowsEdgeAntialiasing(void *id, BOOL value) {
	[(CALayer*)id
		setAllowsEdgeAntialiasing: value];
}

BOOL CALayer_inst_allowsGroupOpacity(void *id) {
	return [(CALayer*)id
		allowsGroupOpacity];
}

void CALayer_inst_setAllowsGroupOpacity(void *id, BOOL value) {
	[(CALayer*)id
		setAllowsGroupOpacity: value];
}

void* CALayer_inst_filters(void *id) {
	return [(CALayer*)id
		filters];
}

void CALayer_inst_setFilters(void *id, void* value) {
	[(CALayer*)id
		setFilters: value];
}

void* CALayer_inst_compositingFilter(void *id) {
	return [(CALayer*)id
		compositingFilter];
}

void CALayer_inst_setCompositingFilter(void *id, void* value) {
	[(CALayer*)id
		setCompositingFilter: value];
}

void* CALayer_inst_backgroundFilters(void *id) {
	return [(CALayer*)id
		backgroundFilters];
}

void CALayer_inst_setBackgroundFilters(void *id, void* value) {
	[(CALayer*)id
		setBackgroundFilters: value];
}

BOOL CALayer_inst_isOpaque(void *id) {
	return [(CALayer*)id
		isOpaque];
}

void CALayer_inst_setOpaque(void *id, BOOL value) {
	[(CALayer*)id
		setOpaque: value];
}

BOOL CALayer_inst_isGeometryFlipped(void *id) {
	return [(CALayer*)id
		isGeometryFlipped];
}

void CALayer_inst_setGeometryFlipped(void *id, BOOL value) {
	[(CALayer*)id
		setGeometryFlipped: value];
}

BOOL CALayer_inst_drawsAsynchronously(void *id) {
	return [(CALayer*)id
		drawsAsynchronously];
}

void CALayer_inst_setDrawsAsynchronously(void *id, BOOL value) {
	[(CALayer*)id
		setDrawsAsynchronously: value];
}

BOOL CALayer_inst_shouldRasterize(void *id) {
	return [(CALayer*)id
		shouldRasterize];
}

void CALayer_inst_setShouldRasterize(void *id, BOOL value) {
	[(CALayer*)id
		setShouldRasterize: value];
}

double CALayer_inst_rasterizationScale(void *id) {
	return [(CALayer*)id
		rasterizationScale];
}

void CALayer_inst_setRasterizationScale(void *id, double value) {
	[(CALayer*)id
		setRasterizationScale: value];
}

NSRect CALayer_inst_frame(void *id) {
	return [(CALayer*)id
		frame];
}

void CALayer_inst_setFrame(void *id, NSRect value) {
	[(CALayer*)id
		setFrame: value];
}

NSRect CALayer_inst_bounds(void *id) {
	return [(CALayer*)id
		bounds];
}

void CALayer_inst_setBounds(void *id, NSRect value) {
	[(CALayer*)id
		setBounds: value];
}

double CALayer_inst_zPosition(void *id) {
	return [(CALayer*)id
		zPosition];
}

void CALayer_inst_setZPosition(void *id, double value) {
	[(CALayer*)id
		setZPosition: value];
}

double CALayer_inst_anchorPointZ(void *id) {
	return [(CALayer*)id
		anchorPointZ];
}

void CALayer_inst_setAnchorPointZ(void *id, double value) {
	[(CALayer*)id
		setAnchorPointZ: value];
}

double CALayer_inst_contentsScale(void *id) {
	return [(CALayer*)id
		contentsScale];
}

void CALayer_inst_setContentsScale(void *id, double value) {
	[(CALayer*)id
		setContentsScale: value];
}

void* CALayer_inst_sublayers(void *id) {
	return [(CALayer*)id
		sublayers];
}

void CALayer_inst_setSublayers(void *id, void* value) {
	[(CALayer*)id
		setSublayers: value];
}

void* CALayer_inst_superlayer(void *id) {
	return [(CALayer*)id
		superlayer];
}

BOOL CALayer_inst_needsDisplayOnBoundsChange(void *id) {
	return [(CALayer*)id
		needsDisplayOnBoundsChange];
}

void CALayer_inst_setNeedsDisplayOnBoundsChange(void *id, BOOL value) {
	[(CALayer*)id
		setNeedsDisplayOnBoundsChange: value];
}

void* CALayer_inst_layoutManager(void *id) {
	return [(CALayer*)id
		layoutManager];
}

void CALayer_inst_setLayoutManager(void *id, void* value) {
	[(CALayer*)id
		setLayoutManager: value];
}

void* CALayer_inst_constraints(void *id) {
	return [(CALayer*)id
		constraints];
}

void CALayer_inst_setConstraints(void *id, void* value) {
	[(CALayer*)id
		setConstraints: value];
}

void* CALayer_inst_actions(void *id) {
	return [(CALayer*)id
		actions];
}

void CALayer_inst_setActions(void *id, void* value) {
	[(CALayer*)id
		setActions: value];
}

NSRect CALayer_inst_visibleRect(void *id) {
	return [(CALayer*)id
		visibleRect];
}

void* CALayer_inst_name(void *id) {
	return [(CALayer*)id
		name];
}

void CALayer_inst_setName(void *id, void* value) {
	[(CALayer*)id
		setName: value];
}

void* NSArray_inst_arrayByAddingObjectsFromArray(void *id, void* otherArray) {
	return [(NSArray*)id
		arrayByAddingObjectsFromArray: otherArray];
}

void* NSArray_inst_componentsJoinedByString(void *id, void* separator) {
	return [(NSArray*)id
		componentsJoinedByString: separator];
}

void* NSArray_inst_descriptionWithLocale(void *id, void* locale) {
	return [(NSArray*)id
		descriptionWithLocale: locale];
}

void* NSArray_inst_descriptionWithLocale_indent(void *id, void* locale, unsigned long level) {
	return [(NSArray*)id
		descriptionWithLocale: locale
		indent: level];
}

void* NSArray_inst_init(void *id) {
	return [(NSArray*)id
		init];
}

void* NSArray_inst_initWithArray(void *id, void* array) {
	return [(NSArray*)id
		initWithArray: array];
}

void* NSArray_inst_initWithArray_copyItems(void *id, void* array, BOOL flag) {
	return [(NSArray*)id
		initWithArray: array
		copyItems: flag];
}

BOOL NSArray_inst_isEqualToArray(void *id, void* otherArray) {
	return [(NSArray*)id
		isEqualToArray: otherArray];
}

void NSArray_inst_makeObjectsPerformSelector(void *id, void* aSelector) {
	[(NSArray*)id
		makeObjectsPerformSelector: aSelector];
}

void NSArray_inst_makeObjectsPerformSelector_withObject(void *id, void* aSelector, void* argument) {
	[(NSArray*)id
		makeObjectsPerformSelector: aSelector
		withObject: argument];
}

void* NSArray_inst_pathsMatchingExtensions(void *id, void* filterTypes) {
	return [(NSArray*)id
		pathsMatchingExtensions: filterTypes];
}

void NSArray_inst_setValue_forKey(void *id, void* value, void* key) {
	[(NSArray*)id
		setValue: value
		forKey: key];
}

void* NSArray_inst_shuffledArray(void *id) {
	return [(NSArray*)id
		shuffledArray];
}

void* NSArray_inst_sortedArrayUsingDescriptors(void *id, void* sortDescriptors) {
	return [(NSArray*)id
		sortedArrayUsingDescriptors: sortDescriptors];
}

void* NSArray_inst_sortedArrayUsingSelector(void *id, void* comparator) {
	return [(NSArray*)id
		sortedArrayUsingSelector: comparator];
}

void* NSArray_inst_valueForKey(void *id, void* key) {
	return [(NSArray*)id
		valueForKey: key];
}

unsigned long NSArray_inst_count(void *id) {
	return [(NSArray*)id
		count];
}

void* NSArray_inst_sortedArrayHint(void *id) {
	return [(NSArray*)id
		sortedArrayHint];
}

void* NSArray_inst_description(void *id) {
	return [(NSArray*)id
		description];
}

void* NSAttributedString_inst_attributedStringByInflectingString(void *id) {
	return [(NSAttributedString*)id
		attributedStringByInflectingString];
}

void NSAttributedString_inst_drawInRect(void *id, NSRect rect) {
	[(NSAttributedString*)id
		drawInRect: rect];
}

void* NSAttributedString_inst_initWithAttributedString(void *id, void* attrStr) {
	return [(NSAttributedString*)id
		initWithAttributedString: attrStr];
}

void* NSAttributedString_inst_initWithDocFormat_documentAttributes(void *id, void* data, void* dict) {
	return [(NSAttributedString*)id
		initWithDocFormat: data
		documentAttributes: dict];
}

void* NSAttributedString_inst_initWithHTML_baseURL_documentAttributes(void *id, void* data, void* base, void* dict) {
	return [(NSAttributedString*)id
		initWithHTML: data
		baseURL: base
		documentAttributes: dict];
}

void* NSAttributedString_inst_initWithHTML_documentAttributes(void *id, void* data, void* dict) {
	return [(NSAttributedString*)id
		initWithHTML: data
		documentAttributes: dict];
}

void* NSAttributedString_inst_initWithHTML_options_documentAttributes(void *id, void* data, void* options, void* dict) {
	return [(NSAttributedString*)id
		initWithHTML: data
		options: options
		documentAttributes: dict];
}

void* NSAttributedString_inst_initWithRTF_documentAttributes(void *id, void* data, void* dict) {
	return [(NSAttributedString*)id
		initWithRTF: data
		documentAttributes: dict];
}

void* NSAttributedString_inst_initWithRTFD_documentAttributes(void *id, void* data, void* dict) {
	return [(NSAttributedString*)id
		initWithRTFD: data
		documentAttributes: dict];
}

void* NSAttributedString_inst_initWithString(void *id, void* str) {
	return [(NSAttributedString*)id
		initWithString: str];
}

void* NSAttributedString_inst_initWithString_attributes(void *id, void* str, void* attrs) {
	return [(NSAttributedString*)id
		initWithString: str
		attributes: attrs];
}

BOOL NSAttributedString_inst_isEqualToAttributedString(void *id, void* other) {
	return [(NSAttributedString*)id
		isEqualToAttributedString: other];
}

unsigned long NSAttributedString_inst_nextWordFromIndex_forward(void *id, unsigned long location, BOOL isForward) {
	return [(NSAttributedString*)id
		nextWordFromIndex: location
		forward: isForward];
}

NSSize NSAttributedString_inst_size(void *id) {
	return [(NSAttributedString*)id
		size];
}

void* NSAttributedString_inst_init(void *id) {
	return [(NSAttributedString*)id
		init];
}

void* NSAttributedString_inst_string(void *id) {
	return [(NSAttributedString*)id
		string];
}

unsigned long NSAttributedString_inst_length(void *id) {
	return [(NSAttributedString*)id
		length];
}

void NSData_inst_getBytes_length(void *id, void* buffer, unsigned long length) {
	[(NSData*)id
		getBytes: buffer
		length: length];
}

void* NSData_inst_initWithBytes_length(void *id, void* bytes, unsigned long length) {
	return [(NSData*)id
		initWithBytes: bytes
		length: length];
}

void* NSData_inst_initWithBytesNoCopy_length(void *id, void* bytes, unsigned long length) {
	return [(NSData*)id
		initWithBytesNoCopy: bytes
		length: length];
}

void* NSData_inst_initWithBytesNoCopy_length_freeWhenDone(void *id, void* bytes, unsigned long length, BOOL b) {
	return [(NSData*)id
		initWithBytesNoCopy: bytes
		length: length
		freeWhenDone: b];
}

void* NSData_inst_initWithContentsOfFile(void *id, void* path) {
	return [(NSData*)id
		initWithContentsOfFile: path];
}

void* NSData_inst_initWithContentsOfURL(void *id, void* url) {
	return [(NSData*)id
		initWithContentsOfURL: url];
}

void* NSData_inst_initWithData(void *id, void* data) {
	return [(NSData*)id
		initWithData: data];
}

BOOL NSData_inst_isEqualToData(void *id, void* other) {
	return [(NSData*)id
		isEqualToData: other];
}

BOOL NSData_inst_writeToFile_atomically(void *id, void* path, BOOL useAuxiliaryFile) {
	return [(NSData*)id
		writeToFile: path
		atomically: useAuxiliaryFile];
}

BOOL NSData_inst_writeToURL_atomically(void *id, void* url, BOOL atomically) {
	return [(NSData*)id
		writeToURL: url
		atomically: atomically];
}

void* NSData_inst_init(void *id) {
	return [(NSData*)id
		init];
}

void* NSData_inst_bytes(void *id) {
	return [(NSData*)id
		bytes];
}

unsigned long NSData_inst_length(void *id) {
	return [(NSData*)id
		length];
}

void* NSData_inst_description(void *id) {
	return [(NSData*)id
		description];
}

void* NSDictionary_inst_descriptionWithLocale(void *id, void* locale) {
	return [(NSDictionary*)id
		descriptionWithLocale: locale];
}

void* NSDictionary_inst_descriptionWithLocale_indent(void *id, void* locale, unsigned long level) {
	return [(NSDictionary*)id
		descriptionWithLocale: locale
		indent: level];
}

BOOL NSDictionary_inst_fileExtensionHidden(void *id) {
	return [(NSDictionary*)id
		fileExtensionHidden];
}

void* NSDictionary_inst_fileGroupOwnerAccountID(void *id) {
	return [(NSDictionary*)id
		fileGroupOwnerAccountID];
}

void* NSDictionary_inst_fileGroupOwnerAccountName(void *id) {
	return [(NSDictionary*)id
		fileGroupOwnerAccountName];
}

BOOL NSDictionary_inst_fileIsAppendOnly(void *id) {
	return [(NSDictionary*)id
		fileIsAppendOnly];
}

BOOL NSDictionary_inst_fileIsImmutable(void *id) {
	return [(NSDictionary*)id
		fileIsImmutable];
}

void* NSDictionary_inst_fileOwnerAccountID(void *id) {
	return [(NSDictionary*)id
		fileOwnerAccountID];
}

void* NSDictionary_inst_fileOwnerAccountName(void *id) {
	return [(NSDictionary*)id
		fileOwnerAccountName];
}

unsigned long NSDictionary_inst_filePosixPermissions(void *id) {
	return [(NSDictionary*)id
		filePosixPermissions];
}

unsigned long NSDictionary_inst_fileSystemFileNumber(void *id) {
	return [(NSDictionary*)id
		fileSystemFileNumber];
}

long NSDictionary_inst_fileSystemNumber(void *id) {
	return [(NSDictionary*)id
		fileSystemNumber];
}

void* NSDictionary_inst_fileType(void *id) {
	return [(NSDictionary*)id
		fileType];
}

void* NSDictionary_inst_init(void *id) {
	return [(NSDictionary*)id
		init];
}

void* NSDictionary_inst_initWithDictionary(void *id, void* otherDictionary) {
	return [(NSDictionary*)id
		initWithDictionary: otherDictionary];
}

void* NSDictionary_inst_initWithDictionary_copyItems(void *id, void* otherDictionary, BOOL flag) {
	return [(NSDictionary*)id
		initWithDictionary: otherDictionary
		copyItems: flag];
}

void* NSDictionary_inst_initWithObjects_forKeys(void *id, void* objects, void* keys) {
	return [(NSDictionary*)id
		initWithObjects: objects
		forKeys: keys];
}

BOOL NSDictionary_inst_isEqualToDictionary(void *id, void* otherDictionary) {
	return [(NSDictionary*)id
		isEqualToDictionary: otherDictionary];
}

void* NSDictionary_inst_keysSortedByValueUsingSelector(void *id, void* comparator) {
	return [(NSDictionary*)id
		keysSortedByValueUsingSelector: comparator];
}

unsigned long NSDictionary_inst_count(void *id) {
	return [(NSDictionary*)id
		count];
}

void* NSDictionary_inst_allKeys(void *id) {
	return [(NSDictionary*)id
		allKeys];
}

void* NSDictionary_inst_allValues(void *id) {
	return [(NSDictionary*)id
		allValues];
}

void* NSDictionary_inst_description(void *id) {
	return [(NSDictionary*)id
		description];
}

void* NSDictionary_inst_descriptionInStringsFileFormat(void *id) {
	return [(NSDictionary*)id
		descriptionInStringsFileFormat];
}

void* NSNumber_inst_descriptionWithLocale(void *id, void* locale) {
	return [(NSNumber*)id
		descriptionWithLocale: locale];
}

void* NSNumber_inst_initWithBool(void *id, BOOL value) {
	return [(NSNumber*)id
		initWithBool: value];
}

void* NSNumber_inst_initWithInt(void *id, int value) {
	return [(NSNumber*)id
		initWithInt: value];
}

void* NSNumber_inst_initWithInteger(void *id, long value) {
	return [(NSNumber*)id
		initWithInteger: value];
}

void* NSNumber_inst_initWithUnsignedInt(void *id, int value) {
	return [(NSNumber*)id
		initWithUnsignedInt: value];
}

void* NSNumber_inst_initWithUnsignedInteger(void *id, unsigned long value) {
	return [(NSNumber*)id
		initWithUnsignedInteger: value];
}

BOOL NSNumber_inst_isEqualToNumber(void *id, void* number) {
	return [(NSNumber*)id
		isEqualToNumber: number];
}

void* NSNumber_inst_init(void *id) {
	return [(NSNumber*)id
		init];
}

BOOL NSNumber_inst_boolValue(void *id) {
	return [(NSNumber*)id
		boolValue];
}

int NSNumber_inst_intValue(void *id) {
	return [(NSNumber*)id
		intValue];
}

long NSNumber_inst_integerValue(void *id) {
	return [(NSNumber*)id
		integerValue];
}

unsigned long NSNumber_inst_unsignedIntegerValue(void *id) {
	return [(NSNumber*)id
		unsignedIntegerValue];
}

int NSNumber_inst_unsignedIntValue(void *id) {
	return [(NSNumber*)id
		unsignedIntValue];
}

void* NSNumber_inst_stringValue(void *id) {
	return [(NSNumber*)id
		stringValue];
}

void NSRunLoop_inst_cancelPerformSelector_target_argument(void *id, void* aSelector, void* target, void* arg) {
	[(NSRunLoop*)id
		cancelPerformSelector: aSelector
		target: target
		argument: arg];
}

void NSRunLoop_inst_cancelPerformSelectorsWithTarget(void *id, void* target) {
	[(NSRunLoop*)id
		cancelPerformSelectorsWithTarget: target];
}

void NSRunLoop_inst_performSelector_target_argument_order_modes(void *id, void* aSelector, void* target, void* arg, unsigned long order, void* modes) {
	[(NSRunLoop*)id
		performSelector: aSelector
		target: target
		argument: arg
		order: order
		modes: modes];
}

void NSRunLoop_inst_run(void *id) {
	[(NSRunLoop*)id
		run];
}

void* NSRunLoop_inst_init(void *id) {
	return [(NSRunLoop*)id
		init];
}

BOOL NSString_inst_canBeConvertedToEncoding(void *id, unsigned long encoding) {
	return [(NSString*)id
		canBeConvertedToEncoding: encoding];
}

unsigned short NSString_inst_characterAtIndex(void *id, unsigned long index) {
	return [(NSString*)id
		characterAtIndex: index];
}

unsigned long NSString_inst_completePathIntoString_caseSensitive_matchesIntoArray_filterTypes(void *id, void* outputName, BOOL flag, void* outputArray, void* filterTypes) {
	return [(NSString*)id
		completePathIntoString: outputName
		caseSensitive: flag
		matchesIntoArray: outputArray
		filterTypes: filterTypes];
}

void* NSString_inst_componentsSeparatedByString(void *id, void* separator) {
	return [(NSString*)id
		componentsSeparatedByString: separator];
}

BOOL NSString_inst_containsString(void *id, void* str) {
	return [(NSString*)id
		containsString: str];
}

void* NSString_inst_dataUsingEncoding(void *id, unsigned long encoding) {
	return [(NSString*)id
		dataUsingEncoding: encoding];
}

void* NSString_inst_dataUsingEncoding_allowLossyConversion(void *id, unsigned long encoding, BOOL lossy) {
	return [(NSString*)id
		dataUsingEncoding: encoding
		allowLossyConversion: lossy];
}

void NSString_inst_drawInRect_withAttributes(void *id, NSRect rect, void* attrs) {
	[(NSString*)id
		drawInRect: rect
		withAttributes: attrs];
}

BOOL NSString_inst_hasPrefix(void *id, void* str) {
	return [(NSString*)id
		hasPrefix: str];
}

BOOL NSString_inst_hasSuffix(void *id, void* str) {
	return [(NSString*)id
		hasSuffix: str];
}

void* NSString_inst_init(void *id) {
	return [(NSString*)id
		init];
}

void* NSString_inst_initWithBytes_length_encoding(void *id, void* bytes, unsigned long len, unsigned long encoding) {
	return [(NSString*)id
		initWithBytes: bytes
		length: len
		encoding: encoding];
}

void* NSString_inst_initWithBytesNoCopy_length_encoding_freeWhenDone(void *id, void* bytes, unsigned long len, unsigned long encoding, BOOL freeBuffer) {
	return [(NSString*)id
		initWithBytesNoCopy: bytes
		length: len
		encoding: encoding
		freeWhenDone: freeBuffer];
}

void* NSString_inst_initWithData_encoding(void *id, void* data, unsigned long encoding) {
	return [(NSString*)id
		initWithData: data
		encoding: encoding];
}

void* NSString_inst_initWithString(void *id, void* aString) {
	return [(NSString*)id
		initWithString: aString];
}

BOOL NSString_inst_isEqualToString(void *id, void* aString) {
	return [(NSString*)id
		isEqualToString: aString];
}

unsigned long NSString_inst_lengthOfBytesUsingEncoding(void *id, unsigned long enc) {
	return [(NSString*)id
		lengthOfBytesUsingEncoding: enc];
}

BOOL NSString_inst_localizedCaseInsensitiveContainsString(void *id, void* str) {
	return [(NSString*)id
		localizedCaseInsensitiveContainsString: str];
}

BOOL NSString_inst_localizedStandardContainsString(void *id, void* str) {
	return [(NSString*)id
		localizedStandardContainsString: str];
}

unsigned long NSString_inst_maximumLengthOfBytesUsingEncoding(void *id, unsigned long enc) {
	return [(NSString*)id
		maximumLengthOfBytesUsingEncoding: enc];
}

void* NSString_inst_propertyList(void *id) {
	return [(NSString*)id
		propertyList];
}

void* NSString_inst_propertyListFromStringsFileFormat(void *id) {
	return [(NSString*)id
		propertyListFromStringsFileFormat];
}

NSSize NSString_inst_sizeWithAttributes(void *id, void* attrs) {
	return [(NSString*)id
		sizeWithAttributes: attrs];
}

void* NSString_inst_stringByAppendingPathComponent(void *id, void* str) {
	return [(NSString*)id
		stringByAppendingPathComponent: str];
}

void* NSString_inst_stringByAppendingPathExtension(void *id, void* str) {
	return [(NSString*)id
		stringByAppendingPathExtension: str];
}

void* NSString_inst_stringByAppendingString(void *id, void* aString) {
	return [(NSString*)id
		stringByAppendingString: aString];
}

void* NSString_inst_stringByPaddingToLength_withString_startingAtIndex(void *id, unsigned long newLength, void* padString, unsigned long padIndex) {
	return [(NSString*)id
		stringByPaddingToLength: newLength
		withString: padString
		startingAtIndex: padIndex];
}

void* NSString_inst_stringByReplacingOccurrencesOfString_withString(void *id, void* target, void* replacement) {
	return [(NSString*)id
		stringByReplacingOccurrencesOfString: target
		withString: replacement];
}

void* NSString_inst_stringsByAppendingPaths(void *id, void* paths) {
	return [(NSString*)id
		stringsByAppendingPaths: paths];
}

void* NSString_inst_substringFromIndex(void *id, unsigned long from) {
	return [(NSString*)id
		substringFromIndex: from];
}

void* NSString_inst_substringToIndex(void *id, unsigned long to) {
	return [(NSString*)id
		substringToIndex: to];
}

void* NSString_inst_variantFittingPresentationWidth(void *id, long width) {
	return [(NSString*)id
		variantFittingPresentationWidth: width];
}

unsigned long NSString_inst_length(void *id) {
	return [(NSString*)id
		length];
}

unsigned long NSString_inst_hash(void *id) {
	return [(NSString*)id
		hash];
}

void* NSString_inst_lowercaseString(void *id) {
	return [(NSString*)id
		lowercaseString];
}

void* NSString_inst_localizedLowercaseString(void *id) {
	return [(NSString*)id
		localizedLowercaseString];
}

void* NSString_inst_uppercaseString(void *id) {
	return [(NSString*)id
		uppercaseString];
}

void* NSString_inst_localizedUppercaseString(void *id) {
	return [(NSString*)id
		localizedUppercaseString];
}

void* NSString_inst_capitalizedString(void *id) {
	return [(NSString*)id
		capitalizedString];
}

void* NSString_inst_localizedCapitalizedString(void *id) {
	return [(NSString*)id
		localizedCapitalizedString];
}

void* NSString_inst_decomposedStringWithCanonicalMapping(void *id) {
	return [(NSString*)id
		decomposedStringWithCanonicalMapping];
}

void* NSString_inst_decomposedStringWithCompatibilityMapping(void *id) {
	return [(NSString*)id
		decomposedStringWithCompatibilityMapping];
}

void* NSString_inst_precomposedStringWithCanonicalMapping(void *id) {
	return [(NSString*)id
		precomposedStringWithCanonicalMapping];
}

void* NSString_inst_precomposedStringWithCompatibilityMapping(void *id) {
	return [(NSString*)id
		precomposedStringWithCompatibilityMapping];
}

int NSString_inst_intValue(void *id) {
	return [(NSString*)id
		intValue];
}

long NSString_inst_integerValue(void *id) {
	return [(NSString*)id
		integerValue];
}

BOOL NSString_inst_boolValue(void *id) {
	return [(NSString*)id
		boolValue];
}

void* NSString_inst_description(void *id) {
	return [(NSString*)id
		description];
}

unsigned long NSString_inst_fastestEncoding(void *id) {
	return [(NSString*)id
		fastestEncoding];
}

unsigned long NSString_inst_smallestEncoding(void *id) {
	return [(NSString*)id
		smallestEncoding];
}

void* NSString_inst_pathComponents(void *id) {
	return [(NSString*)id
		pathComponents];
}

BOOL NSString_inst_isAbsolutePath(void *id) {
	return [(NSString*)id
		isAbsolutePath];
}

void* NSString_inst_lastPathComponent(void *id) {
	return [(NSString*)id
		lastPathComponent];
}

void* NSString_inst_pathExtension(void *id) {
	return [(NSString*)id
		pathExtension];
}

void* NSString_inst_stringByAbbreviatingWithTildeInPath(void *id) {
	return [(NSString*)id
		stringByAbbreviatingWithTildeInPath];
}

void* NSString_inst_stringByDeletingLastPathComponent(void *id) {
	return [(NSString*)id
		stringByDeletingLastPathComponent];
}

void* NSString_inst_stringByDeletingPathExtension(void *id) {
	return [(NSString*)id
		stringByDeletingPathExtension];
}

void* NSString_inst_stringByExpandingTildeInPath(void *id) {
	return [(NSString*)id
		stringByExpandingTildeInPath];
}

void* NSString_inst_stringByResolvingSymlinksInPath(void *id) {
	return [(NSString*)id
		stringByResolvingSymlinksInPath];
}

void* NSString_inst_stringByStandardizingPath(void *id) {
	return [(NSString*)id
		stringByStandardizingPath];
}

void* NSString_inst_stringByRemovingPercentEncoding(void *id) {
	return [(NSString*)id
		stringByRemovingPercentEncoding];
}

void NSThread_inst_cancel(void *id) {
	[(NSThread*)id
		cancel];
}

void* NSThread_inst_init(void *id) {
	return [(NSThread*)id
		init];
}

void* NSThread_inst_initWithTarget_selector_object(void *id, void* target, void* selector, void* argument) {
	return [(NSThread*)id
		initWithTarget: target
		selector: selector
		object: argument];
}

void NSThread_inst_main(void *id) {
	[(NSThread*)id
		main];
}

void NSThread_inst_start(void *id) {
	[(NSThread*)id
		start];
}

BOOL NSThread_inst_isExecuting(void *id) {
	return [(NSThread*)id
		isExecuting];
}

BOOL NSThread_inst_isFinished(void *id) {
	return [(NSThread*)id
		isFinished];
}

BOOL NSThread_inst_isCancelled(void *id) {
	return [(NSThread*)id
		isCancelled];
}

BOOL NSThread_inst_isMainThread(void *id) {
	return [(NSThread*)id
		isMainThread];
}

void* NSThread_inst_name(void *id) {
	return [(NSThread*)id
		name];
}

void NSThread_inst_setName(void *id, void* value) {
	[(NSThread*)id
		setName: value];
}

unsigned long NSThread_inst_stackSize(void *id) {
	return [(NSThread*)id
		stackSize];
}

void NSThread_inst_setStackSize(void *id, unsigned long value) {
	[(NSThread*)id
		setStackSize: value];
}

void* NSURL_inst_URLByAppendingPathComponent(void *id, void* pathComponent) {
	return [(NSURL*)id
		URLByAppendingPathComponent: pathComponent];
}

void* NSURL_inst_URLByAppendingPathComponent_isDirectory(void *id, void* pathComponent, BOOL isDirectory) {
	return [(NSURL*)id
		URLByAppendingPathComponent: pathComponent
		isDirectory: isDirectory];
}

void* NSURL_inst_URLByAppendingPathExtension(void *id, void* pathExtension) {
	return [(NSURL*)id
		URLByAppendingPathExtension: pathExtension];
}

void* NSURL_inst_fileReferenceURL(void *id) {
	return [(NSURL*)id
		fileReferenceURL];
}

void* NSURL_inst_initAbsoluteURLWithDataRepresentation_relativeToURL(void *id, void* data, void* baseURL) {
	return [(NSURL*)id
		initAbsoluteURLWithDataRepresentation: data
		relativeToURL: baseURL];
}

void* NSURL_inst_initFileURLWithPath(void *id, void* path) {
	return [(NSURL*)id
		initFileURLWithPath: path];
}

void* NSURL_inst_initFileURLWithPath_isDirectory(void *id, void* path, BOOL isDir) {
	return [(NSURL*)id
		initFileURLWithPath: path
		isDirectory: isDir];
}

void* NSURL_inst_initFileURLWithPath_isDirectory_relativeToURL(void *id, void* path, BOOL isDir, void* baseURL) {
	return [(NSURL*)id
		initFileURLWithPath: path
		isDirectory: isDir
		relativeToURL: baseURL];
}

void* NSURL_inst_initFileURLWithPath_relativeToURL(void *id, void* path, void* baseURL) {
	return [(NSURL*)id
		initFileURLWithPath: path
		relativeToURL: baseURL];
}

void* NSURL_inst_initWithDataRepresentation_relativeToURL(void *id, void* data, void* baseURL) {
	return [(NSURL*)id
		initWithDataRepresentation: data
		relativeToURL: baseURL];
}

void* NSURL_inst_initWithString(void *id, void* URLString) {
	return [(NSURL*)id
		initWithString: URLString];
}

void* NSURL_inst_initWithString_relativeToURL(void *id, void* URLString, void* baseURL) {
	return [(NSURL*)id
		initWithString: URLString
		relativeToURL: baseURL];
}

BOOL NSURL_inst_isFileReferenceURL(void *id) {
	return [(NSURL*)id
		isFileReferenceURL];
}

void NSURL_inst_removeAllCachedResourceValues(void *id) {
	[(NSURL*)id
		removeAllCachedResourceValues];
}

BOOL NSURL_inst_startAccessingSecurityScopedResource(void *id) {
	return [(NSURL*)id
		startAccessingSecurityScopedResource];
}

void NSURL_inst_stopAccessingSecurityScopedResource(void *id) {
	[(NSURL*)id
		stopAccessingSecurityScopedResource];
}

void* NSURL_inst_init(void *id) {
	return [(NSURL*)id
		init];
}

void* NSURL_inst_dataRepresentation(void *id) {
	return [(NSURL*)id
		dataRepresentation];
}

BOOL NSURL_inst_isFileURL(void *id) {
	return [(NSURL*)id
		isFileURL];
}

void* NSURL_inst_absoluteString(void *id) {
	return [(NSURL*)id
		absoluteString];
}

void* NSURL_inst_absoluteURL(void *id) {
	return [(NSURL*)id
		absoluteURL];
}

void* NSURL_inst_baseURL(void *id) {
	return [(NSURL*)id
		baseURL];
}

void* NSURL_inst_fragment(void *id) {
	return [(NSURL*)id
		fragment];
}

void* NSURL_inst_host(void *id) {
	return [(NSURL*)id
		host];
}

void* NSURL_inst_lastPathComponent(void *id) {
	return [(NSURL*)id
		lastPathComponent];
}

void* NSURL_inst_password(void *id) {
	return [(NSURL*)id
		password];
}

void* NSURL_inst_path(void *id) {
	return [(NSURL*)id
		path];
}

void* NSURL_inst_pathComponents(void *id) {
	return [(NSURL*)id
		pathComponents];
}

void* NSURL_inst_pathExtension(void *id) {
	return [(NSURL*)id
		pathExtension];
}

void* NSURL_inst_port(void *id) {
	return [(NSURL*)id
		port];
}

void* NSURL_inst_query(void *id) {
	return [(NSURL*)id
		query];
}

void* NSURL_inst_relativePath(void *id) {
	return [(NSURL*)id
		relativePath];
}

void* NSURL_inst_relativeString(void *id) {
	return [(NSURL*)id
		relativeString];
}

void* NSURL_inst_resourceSpecifier(void *id) {
	return [(NSURL*)id
		resourceSpecifier];
}

void* NSURL_inst_scheme(void *id) {
	return [(NSURL*)id
		scheme];
}

void* NSURL_inst_standardizedURL(void *id) {
	return [(NSURL*)id
		standardizedURL];
}

void* NSURL_inst_user(void *id) {
	return [(NSURL*)id
		user];
}

void* NSURL_inst_filePathURL(void *id) {
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

BOOL NSURL_inst_hasDirectoryPath(void *id) {
	return [(NSURL*)id
		hasDirectoryPath];
}

void* NSURLRequest_inst_initWithURL(void *id, void* URL) {
	return [(NSURLRequest*)id
		initWithURL: URL];
}

void* NSURLRequest_inst_valueForHTTPHeaderField(void *id, void* field) {
	return [(NSURLRequest*)id
		valueForHTTPHeaderField: field];
}

void* NSURLRequest_inst_init(void *id) {
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

void* NSURLRequest_inst_mainDocumentURL(void *id) {
	return [(NSURLRequest*)id
		mainDocumentURL];
}

void* NSURLRequest_inst_allHTTPHeaderFields(void *id) {
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

BOOL NSURLRequest_inst_allowsCellularAccess(void *id) {
	return [(NSURLRequest*)id
		allowsCellularAccess];
}

BOOL NSURLRequest_inst_allowsConstrainedNetworkAccess(void *id) {
	return [(NSURLRequest*)id
		allowsConstrainedNetworkAccess];
}

BOOL NSURLRequest_inst_allowsExpensiveNetworkAccess(void *id) {
	return [(NSURLRequest*)id
		allowsExpensiveNetworkAccess];
}

BOOL NSURLRequest_inst_assumesHTTP3Capable(void *id) {
	return [(NSURLRequest*)id
		assumesHTTP3Capable];
}

void* NSUserDefaults_inst_URLForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		URLForKey: defaultName];
}

void NSUserDefaults_inst_addSuiteNamed(void *id, void* suiteName) {
	[(NSUserDefaults*)id
		addSuiteNamed: suiteName];
}

void* NSUserDefaults_inst_arrayForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		arrayForKey: defaultName];
}

BOOL NSUserDefaults_inst_boolForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		boolForKey: defaultName];
}

void* NSUserDefaults_inst_dataForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		dataForKey: defaultName];
}

void* NSUserDefaults_inst_dictionaryForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		dictionaryForKey: defaultName];
}

void* NSUserDefaults_inst_dictionaryRepresentation(void *id) {
	return [(NSUserDefaults*)id
		dictionaryRepresentation];
}

void* NSUserDefaults_inst_init(void *id) {
	return [(NSUserDefaults*)id
		init];
}

void* NSUserDefaults_inst_initWithSuiteName(void *id, void* suitename) {
	return [(NSUserDefaults*)id
		initWithSuiteName: suitename];
}

long NSUserDefaults_inst_integerForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		integerForKey: defaultName];
}

void* NSUserDefaults_inst_objectForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		objectForKey: defaultName];
}

BOOL NSUserDefaults_inst_objectIsForcedForKey(void *id, void* key) {
	return [(NSUserDefaults*)id
		objectIsForcedForKey: key];
}

BOOL NSUserDefaults_inst_objectIsForcedForKey_inDomain(void *id, void* key, void* domain) {
	return [(NSUserDefaults*)id
		objectIsForcedForKey: key
		inDomain: domain];
}

void* NSUserDefaults_inst_persistentDomainForName(void *id, void* domainName) {
	return [(NSUserDefaults*)id
		persistentDomainForName: domainName];
}

void NSUserDefaults_inst_registerDefaults(void *id, void* registrationDictionary) {
	[(NSUserDefaults*)id
		registerDefaults: registrationDictionary];
}

void NSUserDefaults_inst_removeObjectForKey(void *id, void* defaultName) {
	[(NSUserDefaults*)id
		removeObjectForKey: defaultName];
}

void NSUserDefaults_inst_removePersistentDomainForName(void *id, void* domainName) {
	[(NSUserDefaults*)id
		removePersistentDomainForName: domainName];
}

void NSUserDefaults_inst_removeSuiteNamed(void *id, void* suiteName) {
	[(NSUserDefaults*)id
		removeSuiteNamed: suiteName];
}

void NSUserDefaults_inst_removeVolatileDomainForName(void *id, void* domainName) {
	[(NSUserDefaults*)id
		removeVolatileDomainForName: domainName];
}

void NSUserDefaults_inst_setBool_forKey(void *id, BOOL value, void* defaultName) {
	[(NSUserDefaults*)id
		setBool: value
		forKey: defaultName];
}

void NSUserDefaults_inst_setInteger_forKey(void *id, long value, void* defaultName) {
	[(NSUserDefaults*)id
		setInteger: value
		forKey: defaultName];
}

void NSUserDefaults_inst_setObject_forKey(void *id, void* value, void* defaultName) {
	[(NSUserDefaults*)id
		setObject: value
		forKey: defaultName];
}

void NSUserDefaults_inst_setPersistentDomain_forName(void *id, void* domain, void* domainName) {
	[(NSUserDefaults*)id
		setPersistentDomain: domain
		forName: domainName];
}

void NSUserDefaults_inst_setURL_forKey(void *id, void* url, void* defaultName) {
	[(NSUserDefaults*)id
		setURL: url
		forKey: defaultName];
}

void NSUserDefaults_inst_setVolatileDomain_forName(void *id, void* domain, void* domainName) {
	[(NSUserDefaults*)id
		setVolatileDomain: domain
		forName: domainName];
}

void* NSUserDefaults_inst_stringArrayForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		stringArrayForKey: defaultName];
}

void* NSUserDefaults_inst_stringForKey(void *id, void* defaultName) {
	return [(NSUserDefaults*)id
		stringForKey: defaultName];
}

BOOL NSUserDefaults_inst_synchronize(void *id) {
	return [(NSUserDefaults*)id
		synchronize];
}

void* NSUserDefaults_inst_volatileDomainForName(void *id, void* domainName) {
	return [(NSUserDefaults*)id
		volatileDomainForName: domainName];
}

void* NSUserDefaults_inst_volatileDomainNames(void *id) {
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

// CALayer_alloc
func CALayer_alloc() CALayer {
	ret := C.CALayer_type_alloc()

	return CALayer_fromPointer(ret)

}

// CALayer_layer Creates and returns an instance of the layer object.
// https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func CALayer_layer() CALayer {
	ret := C.CALayer_type_layer()

	return CALayer_fromPointer(ret)

}

// CALayer_needsDisplayForKey Returns a Boolean indicating whether changes to the specified key require the layer to be redisplayed.
// https://developer.apple.com/documentation/quartzcore/calayer/1410769-needsdisplayforkey?language=objc
func CALayer_needsDisplayForKey(key NSStringRef) bool {
	ret := C.CALayer_type_needsDisplayForKey(
		objc.RefPointer(key),
	)

	return convertObjCBoolToGo(ret)

}

// CALayer_defaultActionForKey Returns the default action for the current class.
// https://developer.apple.com/documentation/quartzcore/calayer/1410954-defaultactionforkey?language=objc
func CALayer_defaultActionForKey(event NSStringRef) objc.Object {
	ret := C.CALayer_type_defaultActionForKey(
		objc.RefPointer(event),
	)

	return objc.Object_fromPointer(ret)

}

// CALayer_defaultValueForKey Specifies the default value associated with the specified key.
// https://developer.apple.com/documentation/quartzcore/calayer/1410886-defaultvalueforkey?language=objc
func CALayer_defaultValueForKey(key NSStringRef) objc.Object {
	ret := C.CALayer_type_defaultValueForKey(
		objc.RefPointer(key),
	)

	return objc.Object_fromPointer(ret)

}

// NSArray_alloc
func NSArray_alloc() NSArray {
	ret := C.NSArray_type_alloc()

	return NSArray_fromPointer(ret)

}

// NSArray_array Creates and returns an empty array.
// https://developer.apple.com/documentation/foundation/nsarray/1460120-array?language=objc
func NSArray_array() NSArray {
	ret := C.NSArray_type_array()

	return NSArray_fromPointer(ret)

}

// NSArray_arrayWithArray Creates and returns an array containing the objects in another given array.
// https://developer.apple.com/documentation/foundation/nsarray/1460122-arraywitharray?language=objc
func NSArray_arrayWithArray(array NSArrayRef) NSArray {
	ret := C.NSArray_type_arrayWithArray(
		objc.RefPointer(array),
	)

	return NSArray_fromPointer(ret)

}

// NSAttributedString_alloc
func NSAttributedString_alloc() NSAttributedString {
	ret := C.NSAttributedString_type_alloc()

	return NSAttributedString_fromPointer(ret)

}

// NSAttributedString_textTypes An array of UTI strings that identify the file types that attributed strings support, either directly or through a user-installed filter service.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1535409-texttypes?language=objc
func NSAttributedString_textTypes() NSArray {
	ret := C.NSAttributedString_type_textTypes()

	return NSArray_fromPointer(ret)

}

// NSAttributedString_textUnfilteredTypes An array of UTI strings that identify the file types that attributed strings support directly.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1528269-textunfilteredtypes?language=objc
func NSAttributedString_textUnfilteredTypes() NSArray {
	ret := C.NSAttributedString_type_textUnfilteredTypes()

	return NSArray_fromPointer(ret)

}

// NSData_alloc
func NSData_alloc() NSData {
	ret := C.NSData_type_alloc()

	return NSData_fromPointer(ret)

}

// NSData_data Creates an empty data object.
// https://developer.apple.com/documentation/foundation/nsdata/1547234-data?language=objc
func NSData_data() NSData {
	ret := C.NSData_type_data()

	return NSData_fromPointer(ret)

}

// NSData_dataWithBytes_length Creates a data object containing a given number of bytes copied from a given buffer.
// https://developer.apple.com/documentation/foundation/nsdata/1547231-datawithbytes?language=objc
func NSData_dataWithBytes_length(bytes unsafe.Pointer, length NSUInteger) NSData {
	ret := C.NSData_type_dataWithBytes_length(
		bytes,
		C.ulong(length),
	)

	return NSData_fromPointer(ret)

}

// NSData_dataWithBytesNoCopy_length Creates a data object that holds a given number of bytes from a given buffer.
// https://developer.apple.com/documentation/foundation/nsdata/1547229-datawithbytesnocopy?language=objc
func NSData_dataWithBytesNoCopy_length(bytes unsafe.Pointer, length NSUInteger) NSData {
	ret := C.NSData_type_dataWithBytesNoCopy_length(
		bytes,
		C.ulong(length),
	)

	return NSData_fromPointer(ret)

}

// NSData_dataWithBytesNoCopy_length_freeWhenDone Creates a data object that holds a given number of bytes from a given buffer.
// https://developer.apple.com/documentation/foundation/nsdata/1547240-datawithbytesnocopy?language=objc
func NSData_dataWithBytesNoCopy_length_freeWhenDone(bytes unsafe.Pointer, length NSUInteger, b bool) NSData {
	ret := C.NSData_type_dataWithBytesNoCopy_length_freeWhenDone(
		bytes,
		C.ulong(length),
		convertToObjCBool(b),
	)

	return NSData_fromPointer(ret)

}

// NSData_dataWithData Creates a data object containing the contents of another data object.
// https://developer.apple.com/documentation/foundation/nsdata/1547230-datawithdata?language=objc
func NSData_dataWithData(data NSDataRef) NSData {
	ret := C.NSData_type_dataWithData(
		objc.RefPointer(data),
	)

	return NSData_fromPointer(ret)

}

// NSData_dataWithContentsOfFile Creates a data object by reading every byte from the file at a given path.
// https://developer.apple.com/documentation/foundation/nsdata/1547226-datawithcontentsoffile?language=objc
func NSData_dataWithContentsOfFile(path NSStringRef) NSData {
	ret := C.NSData_type_dataWithContentsOfFile(
		objc.RefPointer(path),
	)

	return NSData_fromPointer(ret)

}

// NSData_dataWithContentsOfURL Creates a data object containing the data from the location specified by a given URL.
// https://developer.apple.com/documentation/foundation/nsdata/1547245-datawithcontentsofurl?language=objc
func NSData_dataWithContentsOfURL(url NSURLRef) NSData {
	ret := C.NSData_type_dataWithContentsOfURL(
		objc.RefPointer(url),
	)

	return NSData_fromPointer(ret)

}

// NSDictionary_alloc
func NSDictionary_alloc() NSDictionary {
	ret := C.NSDictionary_type_alloc()

	return NSDictionary_fromPointer(ret)

}

// NSDictionary_dictionary Creates an empty dictionary.
// https://developer.apple.com/documentation/foundation/nsdictionary/1574180-dictionary?language=objc
func NSDictionary_dictionary() NSDictionary {
	ret := C.NSDictionary_type_dictionary()

	return NSDictionary_fromPointer(ret)

}

// NSDictionary_dictionaryWithObjects_forKeys Creates a dictionary containing entries constructed from the contents of an array of keys and an array of values.
// https://developer.apple.com/documentation/foundation/nsdictionary/1574183-dictionarywithobjects?language=objc
func NSDictionary_dictionaryWithObjects_forKeys(objects NSArrayRef, keys NSArrayRef) NSDictionary {
	ret := C.NSDictionary_type_dictionaryWithObjects_forKeys(
		objc.RefPointer(objects),
		objc.RefPointer(keys),
	)

	return NSDictionary_fromPointer(ret)

}

// NSDictionary_dictionaryWithDictionary Creates a dictionary containing the keys and values from another given dictionary.
// https://developer.apple.com/documentation/foundation/nsdictionary/1574191-dictionarywithdictionary?language=objc
func NSDictionary_dictionaryWithDictionary(dict NSDictionaryRef) NSDictionary {
	ret := C.NSDictionary_type_dictionaryWithDictionary(
		objc.RefPointer(dict),
	)

	return NSDictionary_fromPointer(ret)

}

// NSDictionary_sharedKeySetForKeys Creates a shared key set object for the specified keys.
// https://developer.apple.com/documentation/foundation/nsdictionary/1408190-sharedkeysetforkeys?language=objc
func NSDictionary_sharedKeySetForKeys(keys NSArrayRef) objc.Object {
	ret := C.NSDictionary_type_sharedKeySetForKeys(
		objc.RefPointer(keys),
	)

	return objc.Object_fromPointer(ret)

}

// NSNumber_alloc
func NSNumber_alloc() NSNumber {
	ret := C.NSNumber_type_alloc()

	return NSNumber_fromPointer(ret)

}

// NSNumber_numberWithBool Creates and returns an NSNumber object containing a given value, treating it as a BOOL.
// https://developer.apple.com/documentation/foundation/nsnumber/1551475-numberwithbool?language=objc
func NSNumber_numberWithBool(value bool) NSNumber {
	ret := C.NSNumber_type_numberWithBool(
		convertToObjCBool(value),
	)

	return NSNumber_fromPointer(ret)

}

// NSNumber_numberWithInt Creates and returns an NSNumber object containing a given value, treating it as a signed int.
// https://developer.apple.com/documentation/foundation/nsnumber/1551470-numberwithint?language=objc
func NSNumber_numberWithInt(value int32) NSNumber {
	ret := C.NSNumber_type_numberWithInt(
		C.int(value),
	)

	return NSNumber_fromPointer(ret)

}

// NSNumber_numberWithInteger Creates and returns an NSNumber object containing a given value, treating it as an NSInteger.
// https://developer.apple.com/documentation/foundation/nsnumber/1551473-numberwithinteger?language=objc
func NSNumber_numberWithInteger(value NSInteger) NSNumber {
	ret := C.NSNumber_type_numberWithInteger(
		C.long(value),
	)

	return NSNumber_fromPointer(ret)

}

// NSNumber_numberWithUnsignedInt Creates and returns an NSNumber object containing a given value, treating it as an unsigned int.
// https://developer.apple.com/documentation/foundation/nsnumber/1551472-numberwithunsignedint?language=objc
func NSNumber_numberWithUnsignedInt(value int32) NSNumber {
	ret := C.NSNumber_type_numberWithUnsignedInt(
		C.int(value),
	)

	return NSNumber_fromPointer(ret)

}

// NSNumber_numberWithUnsignedInteger Creates and returns an NSNumber object containing a given value, treating it as an NSUInteger.
// https://developer.apple.com/documentation/foundation/nsnumber/1551469-numberwithunsignedinteger?language=objc
func NSNumber_numberWithUnsignedInteger(value NSUInteger) NSNumber {
	ret := C.NSNumber_type_numberWithUnsignedInteger(
		C.ulong(value),
	)

	return NSNumber_fromPointer(ret)

}

// NSRunLoop_alloc
func NSRunLoop_alloc() NSRunLoop {
	ret := C.NSRunLoop_type_alloc()

	return NSRunLoop_fromPointer(ret)

}

// NSRunLoop_currentRunLoop Returns the run loop for the current thread.
// https://developer.apple.com/documentation/foundation/nsrunloop/1412291-currentrunloop?language=objc
func NSRunLoop_currentRunLoop() NSRunLoop {
	ret := C.NSRunLoop_type_currentRunLoop()

	return NSRunLoop_fromPointer(ret)

}

// NSRunLoop_mainRunLoop Returns the run loop of the main thread.
// https://developer.apple.com/documentation/foundation/nsrunloop/1418388-mainrunloop?language=objc
func NSRunLoop_mainRunLoop() NSRunLoop {
	ret := C.NSRunLoop_type_mainRunLoop()

	return NSRunLoop_fromPointer(ret)

}

// NSString_alloc
func NSString_alloc() NSString {
	ret := C.NSString_type_alloc()

	return NSString_fromPointer(ret)

}

// NSString_string Returns an empty string.
// https://developer.apple.com/documentation/foundation/nsstring/1497312-string?language=objc
func NSString_string() NSString {
	ret := C.NSString_type_string()

	return NSString_fromPointer(ret)

}

// NSString_localizedUserNotificationStringForKey_arguments Returns a localized string intended for display in a notification alert.
// https://developer.apple.com/documentation/foundation/nsstring/1649585-localizedusernotificationstringf?language=objc
func NSString_localizedUserNotificationStringForKey_arguments(key NSStringRef, arguments NSArrayRef) NSString {
	ret := C.NSString_type_localizedUserNotificationStringForKey_arguments(
		objc.RefPointer(key),
		objc.RefPointer(arguments),
	)

	return NSString_fromPointer(ret)

}

// NSString_stringWithString Returns a string created by copying the characters from another given string.
// https://developer.apple.com/documentation/foundation/nsstring/1497372-stringwithstring?language=objc
func NSString_stringWithString(string NSStringRef) NSString {
	ret := C.NSString_type_stringWithString(
		objc.RefPointer(string),
	)

	return NSString_fromPointer(ret)

}

// NSString_localizedNameOfStringEncoding Returns a human-readable string giving the name of a given encoding.
// https://developer.apple.com/documentation/foundation/nsstring/1408318-localizednameofstringencoding?language=objc
func NSString_localizedNameOfStringEncoding(encoding NSStringEncoding) NSString {
	ret := C.NSString_type_localizedNameOfStringEncoding(
		C.ulong(encoding),
	)

	return NSString_fromPointer(ret)

}

// NSString_pathWithComponents Returns a string built from the strings in a given array by concatenating them with a path separator between each pair.
// https://developer.apple.com/documentation/foundation/nsstring/1417198-pathwithcomponents?language=objc
func NSString_pathWithComponents(components NSArrayRef) NSString {
	ret := C.NSString_type_pathWithComponents(
		objc.RefPointer(components),
	)

	return NSString_fromPointer(ret)

}

// NSString_defaultCStringEncoding Returns the C-string encoding assumed for any method accepting a C string as an argument.
// https://developer.apple.com/documentation/foundation/nsstring/1410091-defaultcstringencoding?language=objc
func NSString_defaultCStringEncoding() NSStringEncoding {
	ret := C.NSString_type_defaultCStringEncoding()

	return NSStringEncoding(ret)

}

// NSThread_alloc
func NSThread_alloc() NSThread {
	ret := C.NSThread_type_alloc()

	return NSThread_fromPointer(ret)

}

// NSThread_detachNewThreadSelector_toTarget_withObject Detaches a new thread and uses the specified selector as the thread entry point.
// https://developer.apple.com/documentation/foundation/nsthread/1415633-detachnewthreadselector?language=objc
func NSThread_detachNewThreadSelector_toTarget_withObject(selector objc.Selector, target objc.Ref, argument objc.Ref) {
	C.NSThread_type_detachNewThreadSelector_toTarget_withObject(
		selector.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(argument),
	)

	return

}

// NSThread_exit Terminates the current thread.
// https://developer.apple.com/documentation/foundation/nsthread/1409404-exit?language=objc
func NSThread_exit() {
	C.NSThread_type_exit()

	return

}

// NSThread_isMultiThreaded Returns whether the application is multithreaded.
// https://developer.apple.com/documentation/foundation/nsthread/1410702-ismultithreaded?language=objc
func NSThread_isMultiThreaded() bool {
	ret := C.NSThread_type_isMultiThreaded()

	return convertObjCBoolToGo(ret)

}

// NSThread_isMainThread Returns a Boolean value that indicates whether the current thread is the main thread.
// https://developer.apple.com/documentation/foundation/nsthread/1412704-ismainthread?language=objc
func NSThread_isMainThread() bool {
	ret := C.NSThread_type_isMainThread()

	return convertObjCBoolToGo(ret)

}

// NSThread_mainThread Returns the NSThread object representing the main thread.
// https://developer.apple.com/documentation/foundation/nsthread/1414782-mainthread?language=objc
func NSThread_mainThread() NSThread {
	ret := C.NSThread_type_mainThread()

	return NSThread_fromPointer(ret)

}

// NSThread_currentThread Returns the thread object representing the current thread of execution.
// https://developer.apple.com/documentation/foundation/nsthread/1410679-currentthread?language=objc
func NSThread_currentThread() NSThread {
	ret := C.NSThread_type_currentThread()

	return NSThread_fromPointer(ret)

}

// NSThread_callStackReturnAddresses Returns an array containing the call stack return addresses.
// https://developer.apple.com/documentation/foundation/nsthread/1409565-callstackreturnaddresses?language=objc
func NSThread_callStackReturnAddresses() NSArray {
	ret := C.NSThread_type_callStackReturnAddresses()

	return NSArray_fromPointer(ret)

}

// NSThread_callStackSymbols Returns an array containing the call stack symbols.
// https://developer.apple.com/documentation/foundation/nsthread/1414836-callstacksymbols?language=objc
func NSThread_callStackSymbols() NSArray {
	ret := C.NSThread_type_callStackSymbols()

	return NSArray_fromPointer(ret)

}

// NSURL_alloc
func NSURL_alloc() NSURL {
	ret := C.NSURL_type_alloc()

	return NSURL_fromPointer(ret)

}

// NSURL_URLWithString Creates and returns an NSURL object initialized with a provided URL string.
// https://developer.apple.com/documentation/foundation/nsurl/1572047-urlwithstring?language=objc
func NSURL_URLWithString(URLString NSStringRef) NSURL {
	ret := C.NSURL_type_URLWithString(
		objc.RefPointer(URLString),
	)

	return NSURL_fromPointer(ret)

}

// NSURL_URLWithString_relativeToURL Creates and returns an NSURL object initialized with a base URL and a relative string.
// https://developer.apple.com/documentation/foundation/nsurl/1572049-urlwithstring?language=objc
func NSURL_URLWithString_relativeToURL(URLString NSStringRef, baseURL NSURLRef) NSURL {
	ret := C.NSURL_type_URLWithString_relativeToURL(
		objc.RefPointer(URLString),
		objc.RefPointer(baseURL),
	)

	return NSURL_fromPointer(ret)

}

// NSURL_fileURLWithPath_isDirectory Initializes and returns a newly created NSURL object as a file URL with a specified path.
// https://developer.apple.com/documentation/foundation/nsurl/1414650-fileurlwithpath?language=objc
func NSURL_fileURLWithPath_isDirectory(path NSStringRef, isDir bool) NSURL {
	ret := C.NSURL_type_fileURLWithPath_isDirectory(
		objc.RefPointer(path),
		convertToObjCBool(isDir),
	)

	return NSURL_fromPointer(ret)

}

// NSURL_fileURLWithPath_relativeToURL
// https://developer.apple.com/documentation/foundation/nsurl/1413201-fileurlwithpath?language=objc
func NSURL_fileURLWithPath_relativeToURL(path NSStringRef, baseURL NSURLRef) NSURL {
	ret := C.NSURL_type_fileURLWithPath_relativeToURL(
		objc.RefPointer(path),
		objc.RefPointer(baseURL),
	)

	return NSURL_fromPointer(ret)

}

// NSURL_fileURLWithPath_isDirectory_relativeToURL
// https://developer.apple.com/documentation/foundation/nsurl/1413020-fileurlwithpath?language=objc
func NSURL_fileURLWithPath_isDirectory_relativeToURL(path NSStringRef, isDir bool, baseURL NSURLRef) NSURL {
	ret := C.NSURL_type_fileURLWithPath_isDirectory_relativeToURL(
		objc.RefPointer(path),
		convertToObjCBool(isDir),
		objc.RefPointer(baseURL),
	)

	return NSURL_fromPointer(ret)

}

// NSURL_fileURLWithPath Initializes and returns a newly created NSURL object as a file URL with a specified path.
// https://developer.apple.com/documentation/foundation/nsurl/1410828-fileurlwithpath?language=objc
func NSURL_fileURLWithPath(path NSStringRef) NSURL {
	ret := C.NSURL_type_fileURLWithPath(
		objc.RefPointer(path),
	)

	return NSURL_fromPointer(ret)

}

// NSURL_fileURLWithPathComponents Initializes and returns a newly created NSURL object as a file URL with specified path components.
// https://developer.apple.com/documentation/foundation/nsurl/1414206-fileurlwithpathcomponents?language=objc
func NSURL_fileURLWithPathComponents(components NSArrayRef) NSURL {
	ret := C.NSURL_type_fileURLWithPathComponents(
		objc.RefPointer(components),
	)

	return NSURL_fromPointer(ret)

}

// NSURL_absoluteURLWithDataRepresentation_relativeToURL
// https://developer.apple.com/documentation/foundation/nsurl/1412404-absoluteurlwithdatarepresentatio?language=objc
func NSURL_absoluteURLWithDataRepresentation_relativeToURL(data NSDataRef, baseURL NSURLRef) NSURL {
	ret := C.NSURL_type_absoluteURLWithDataRepresentation_relativeToURL(
		objc.RefPointer(data),
		objc.RefPointer(baseURL),
	)

	return NSURL_fromPointer(ret)

}

// NSURL_URLWithDataRepresentation_relativeToURL
// https://developer.apple.com/documentation/foundation/nsurl/1572042-urlwithdatarepresentation?language=objc
func NSURL_URLWithDataRepresentation_relativeToURL(data NSDataRef, baseURL NSURLRef) NSURL {
	ret := C.NSURL_type_URLWithDataRepresentation_relativeToURL(
		objc.RefPointer(data),
		objc.RefPointer(baseURL),
	)

	return NSURL_fromPointer(ret)

}

// NSURL_resourceValuesForKeys_fromBookmarkData Returns the resource values for properties identified by a specified array of keys contained in specified bookmark data.
// https://developer.apple.com/documentation/foundation/nsurl/1418097-resourcevaluesforkeys?language=objc
func NSURL_resourceValuesForKeys_fromBookmarkData(keys NSArrayRef, bookmarkData NSDataRef) NSDictionary {
	ret := C.NSURL_type_resourceValuesForKeys_fromBookmarkData(
		objc.RefPointer(keys),
		objc.RefPointer(bookmarkData),
	)

	return NSDictionary_fromPointer(ret)

}

// NSURLRequest_alloc
func NSURLRequest_alloc() NSURLRequest {
	ret := C.NSURLRequest_type_alloc()

	return NSURLRequest_fromPointer(ret)

}

// NSURLRequest_requestWithURL Creates and returns a URL request for a specified URL.
// https://developer.apple.com/documentation/foundation/nsurlrequest/1528603-requestwithurl?language=objc
func NSURLRequest_requestWithURL(URL NSURLRef) NSURLRequest {
	ret := C.NSURLRequest_type_requestWithURL(
		objc.RefPointer(URL),
	)

	return NSURLRequest_fromPointer(ret)

}

// NSURLRequest_supportsSecureCoding A Boolean value indicating whether the NSURLRequest implements the NSSecureCoding protocol.
// https://developer.apple.com/documentation/foundation/nsurlrequest/1416510-supportssecurecoding?language=objc
func NSURLRequest_supportsSecureCoding() bool {
	ret := C.NSURLRequest_type_supportsSecureCoding()

	return convertObjCBoolToGo(ret)

}

// NSUserDefaults_alloc
func NSUserDefaults_alloc() NSUserDefaults {
	ret := C.NSUserDefaults_type_alloc()

	return NSUserDefaults_fromPointer(ret)

}

// NSUserDefaults_resetStandardUserDefaults This method has no effect and shouldn't be used.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1407708-resetstandarduserdefaults?language=objc
func NSUserDefaults_resetStandardUserDefaults() {
	C.NSUserDefaults_type_resetStandardUserDefaults()

	return

}

// NSUserDefaults_standardUserDefaults Returns the shared defaults object.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1416603-standarduserdefaults?language=objc
func NSUserDefaults_standardUserDefaults() NSUserDefaults {
	ret := C.NSUserDefaults_type_standardUserDefaults()

	return NSUserDefaults_fromPointer(ret)

}

type CALayerRef interface {
	Pointer() uintptr
	Init_asCALayer() CALayer
}

type gen_CALayer struct {
	objc.Object
}

func CALayer_fromPointer(ptr unsafe.Pointer) CALayer {
	return CALayer{gen_CALayer{
		objc.Object_fromPointer(ptr),
	}}
}

func CALayer_fromRef(ref objc.Ref) CALayer {
	return CALayer_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// ActionForKey Returns the action object assigned to the specified key.
// https://developer.apple.com/documentation/quartzcore/calayer/1410844-actionforkey?language=objc
func (x gen_CALayer) ActionForKey(
	event NSStringRef,
) objc.Object {
	ret := C.CALayer_inst_actionForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)

	return objc.Object_fromPointer(ret)

}

// AddSublayer Appends the layer to the layer’s list of sublayers.
// https://developer.apple.com/documentation/quartzcore/calayer/1410833-addsublayer?language=objc
func (x gen_CALayer) AddSublayer(
	layer CALayerRef,
) {
	C.CALayer_inst_addSublayer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(layer),
	)

	return

}

// AnimationKeys Returns an array of strings that identify the animations currently attached to the layer.
// https://developer.apple.com/documentation/quartzcore/calayer/1410937-animationkeys?language=objc
func (x gen_CALayer) AnimationKeys() NSArray {
	ret := C.CALayer_inst_animationKeys(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_fromPointer(ret)

}

// ContentsAreFlipped Returns a Boolean indicating whether the layer content is implicitly flipped when rendered.
// https://developer.apple.com/documentation/quartzcore/calayer/1410777-contentsareflipped?language=objc
func (x gen_CALayer) ContentsAreFlipped() bool {
	ret := C.CALayer_inst_contentsAreFlipped(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// ConvertRect_fromLayer Converts the rectangle from the specified layer’s coordinate system to the receiver’s coordinate system.
// https://developer.apple.com/documentation/quartzcore/calayer/1410948-convertrect?language=objc
func (x gen_CALayer) ConvertRect_fromLayer(
	r NSRect,
	l CALayerRef,
) NSRect {
	ret := C.CALayer_inst_convertRect_fromLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&r)),
		objc.RefPointer(l),
	)

	return *(*NSRect)(unsafe.Pointer(&ret))

}

// ConvertRect_toLayer Converts the rectangle from the receiver’s coordinate system to the specified layer’s coordinate system.
// https://developer.apple.com/documentation/quartzcore/calayer/1410742-convertrect?language=objc
func (x gen_CALayer) ConvertRect_toLayer(
	r NSRect,
	l CALayerRef,
) NSRect {
	ret := C.CALayer_inst_convertRect_toLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&r)),
		objc.RefPointer(l),
	)

	return *(*NSRect)(unsafe.Pointer(&ret))

}

// Display Reloads the content of this layer.
// https://developer.apple.com/documentation/quartzcore/calayer/1410926-display?language=objc
func (x gen_CALayer) Display() {
	C.CALayer_inst_display(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// DisplayIfNeeded Initiates the update process for a layer if it is currently marked as needing an update.
// https://developer.apple.com/documentation/quartzcore/calayer/1410813-displayifneeded?language=objc
func (x gen_CALayer) DisplayIfNeeded() {
	C.CALayer_inst_displayIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Init_asCALayer Returns an initialized CALayer object.
// https://developer.apple.com/documentation/quartzcore/calayer/1410835-init?language=objc
func (x gen_CALayer) Init_asCALayer() CALayer {
	ret := C.CALayer_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return CALayer_fromPointer(ret)

}

// InitWithLayer_asCALayer Override to copy or initialize custom fields of the specified layer.
// https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func (x gen_CALayer) InitWithLayer_asCALayer(
	layer objc.Ref,
) CALayer {
	ret := C.CALayer_inst_initWithLayer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(layer),
	)

	return CALayer_fromPointer(ret)

}

// InsertSublayer_above Inserts the specified sublayer above a different sublayer that already belongs to the receiver.
// https://developer.apple.com/documentation/quartzcore/calayer/1410798-insertsublayer?language=objc
func (x gen_CALayer) InsertSublayer_above(
	layer CALayerRef,
	sibling CALayerRef,
) {
	C.CALayer_inst_insertSublayer_above(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(layer),
		objc.RefPointer(sibling),
	)

	return

}

// InsertSublayer_atIndex Inserts the specified layer into the receiver’s list of sublayers at the specified index.
// https://developer.apple.com/documentation/quartzcore/calayer/1410944-insertsublayer?language=objc
func (x gen_CALayer) InsertSublayer_atIndex(
	layer CALayerRef,
	idx int32,
) {
	C.CALayer_inst_insertSublayer_atIndex(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(layer),
		C.int(idx),
	)

	return

}

// InsertSublayer_below Inserts the specified sublayer below a different sublayer that already belongs to the receiver.
// https://developer.apple.com/documentation/quartzcore/calayer/1410840-insertsublayer?language=objc
func (x gen_CALayer) InsertSublayer_below(
	layer CALayerRef,
	sibling CALayerRef,
) {
	C.CALayer_inst_insertSublayer_below(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(layer),
		objc.RefPointer(sibling),
	)

	return

}

// LayoutIfNeeded Recalculate the receiver’s layout, if required.
// https://developer.apple.com/documentation/quartzcore/calayer/1410873-layoutifneeded?language=objc
func (x gen_CALayer) LayoutIfNeeded() {
	C.CALayer_inst_layoutIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// LayoutSublayers Tells the layer to update its layout.
// https://developer.apple.com/documentation/quartzcore/calayer/1410935-layoutsublayers?language=objc
func (x gen_CALayer) LayoutSublayers() {
	C.CALayer_inst_layoutSublayers(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ModelLayer_asCALayer Returns the model layer object associated with the receiver, if any.
// https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func (x gen_CALayer) ModelLayer_asCALayer() CALayer {
	ret := C.CALayer_inst_modelLayer(
		unsafe.Pointer(x.Pointer()),
	)

	return CALayer_fromPointer(ret)

}

// NeedsDisplay Returns a Boolean indicating whether the layer has been marked as needing an update.
// https://developer.apple.com/documentation/quartzcore/calayer/1410958-needsdisplay?language=objc
func (x gen_CALayer) NeedsDisplay() bool {
	ret := C.CALayer_inst_needsDisplay(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// NeedsLayout Returns a Boolean indicating whether the layer has been marked as needing a layout update.
// https://developer.apple.com/documentation/quartzcore/calayer/1410956-needslayout?language=objc
func (x gen_CALayer) NeedsLayout() bool {
	ret := C.CALayer_inst_needsLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// PreferredFrameSize Returns the preferred size of the layer in the coordinate space of its superlayer.
// https://developer.apple.com/documentation/quartzcore/calayer/1410980-preferredframesize?language=objc
func (x gen_CALayer) PreferredFrameSize() NSSize {
	ret := C.CALayer_inst_preferredFrameSize(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSSize)(unsafe.Pointer(&ret))

}

// PresentationLayer_asCALayer Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen.
// https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func (x gen_CALayer) PresentationLayer_asCALayer() CALayer {
	ret := C.CALayer_inst_presentationLayer(
		unsafe.Pointer(x.Pointer()),
	)

	return CALayer_fromPointer(ret)

}

// RemoveAllAnimations Remove all animations attached to the layer.
// https://developer.apple.com/documentation/quartzcore/calayer/1410810-removeallanimations?language=objc
func (x gen_CALayer) RemoveAllAnimations() {
	C.CALayer_inst_removeAllAnimations(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// RemoveAnimationForKey Remove the animation object with the specified key.
// https://developer.apple.com/documentation/quartzcore/calayer/1410939-removeanimationforkey?language=objc
func (x gen_CALayer) RemoveAnimationForKey(
	key NSStringRef,
) {
	C.CALayer_inst_removeAnimationForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)

	return

}

// RemoveFromSuperlayer Detaches the layer from its parent layer.
// https://developer.apple.com/documentation/quartzcore/calayer/1410767-removefromsuperlayer?language=objc
func (x gen_CALayer) RemoveFromSuperlayer() {
	C.CALayer_inst_removeFromSuperlayer(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ReplaceSublayer_with Replaces the specified sublayer with a different layer object.
// https://developer.apple.com/documentation/quartzcore/calayer/1410820-replacesublayer?language=objc
func (x gen_CALayer) ReplaceSublayer_with(
	oldLayer CALayerRef,
	newLayer CALayerRef,
) {
	C.CALayer_inst_replaceSublayer_with(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(oldLayer),
		objc.RefPointer(newLayer),
	)

	return

}

// ResizeSublayersWithOldSize Informs the receiver’s sublayers that the receiver’s size has changed.
// https://developer.apple.com/documentation/quartzcore/calayer/1410929-resizesublayerswitholdsize?language=objc
func (x gen_CALayer) ResizeSublayersWithOldSize(
	size NSSize,
) {
	C.CALayer_inst_resizeSublayersWithOldSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return

}

// ResizeWithOldSuperlayerSize Informs the receiver that the size of its superlayer changed.
// https://developer.apple.com/documentation/quartzcore/calayer/1410894-resizewitholdsuperlayersize?language=objc
func (x gen_CALayer) ResizeWithOldSuperlayerSize(
	size NSSize,
) {
	C.CALayer_inst_resizeWithOldSuperlayerSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)

	return

}

// ScrollRectToVisible Initiates a scroll in the layer’s closest ancestor scroll layer so that the specified rectangle becomes visible.
// https://developer.apple.com/documentation/quartzcore/calayer/1522139-scrollrecttovisible?language=objc
func (x gen_CALayer) ScrollRectToVisible(
	r NSRect,
) {
	C.CALayer_inst_scrollRectToVisible(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&r)),
	)

	return

}

// SetNeedsDisplay Marks the layer’s contents as needing to be updated.
// https://developer.apple.com/documentation/quartzcore/calayer/1410855-setneedsdisplay?language=objc
func (x gen_CALayer) SetNeedsDisplay() {
	C.CALayer_inst_setNeedsDisplay(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// SetNeedsDisplayInRect Marks the region within the specified rectangle as needing to be updated.
// https://developer.apple.com/documentation/quartzcore/calayer/1410800-setneedsdisplayinrect?language=objc
func (x gen_CALayer) SetNeedsDisplayInRect(
	r NSRect,
) {
	C.CALayer_inst_setNeedsDisplayInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&r)),
	)

	return

}

// SetNeedsLayout Invalidates the layer’s layout and marks it as needing an update.
// https://developer.apple.com/documentation/quartzcore/calayer/1410946-setneedslayout?language=objc
func (x gen_CALayer) SetNeedsLayout() {
	C.CALayer_inst_setNeedsLayout(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// ShouldArchiveValueForKey Returns a Boolean indicating whether the value of the specified key should be archived.
// https://developer.apple.com/documentation/quartzcore/calayer/1410753-shouldarchivevalueforkey?language=objc
func (x gen_CALayer) ShouldArchiveValueForKey(
	key NSStringRef,
) bool {
	ret := C.CALayer_inst_shouldArchiveValueForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)

	return convertObjCBoolToGo(ret)

}

// Delegate The layer’s delegate object.
// https://developer.apple.com/documentation/quartzcore/calayer/1410984-delegate?language=objc
func (x gen_CALayer) Delegate() objc.Object {
	ret := C.CALayer_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetDelegate The layer’s delegate object.
// https://developer.apple.com/documentation/quartzcore/calayer/1410984-delegate?language=objc
func (x gen_CALayer) SetDelegate(
	value objc.Ref,
) {
	C.CALayer_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Contents An object that provides the contents of the layer. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410773-contents?language=objc
func (x gen_CALayer) Contents() objc.Object {
	ret := C.CALayer_inst_contents(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetContents An object that provides the contents of the layer. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410773-contents?language=objc
func (x gen_CALayer) SetContents(
	value objc.Ref,
) {
	C.CALayer_inst_setContents(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// ContentsRect The rectangle, in the unit coordinate space, that defines the portion of the layer’s contents that should be used. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410866-contentsrect?language=objc
func (x gen_CALayer) ContentsRect() NSRect {
	ret := C.CALayer_inst_contentsRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSRect)(unsafe.Pointer(&ret))

}

// SetContentsRect The rectangle, in the unit coordinate space, that defines the portion of the layer’s contents that should be used. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410866-contentsrect?language=objc
func (x gen_CALayer) SetContentsRect(
	value NSRect,
) {
	C.CALayer_inst_setContentsRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return

}

// ContentsCenter The rectangle that defines how the layer contents are scaled if the layer’s contents are resized. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410740-contentscenter?language=objc
func (x gen_CALayer) ContentsCenter() NSRect {
	ret := C.CALayer_inst_contentsCenter(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSRect)(unsafe.Pointer(&ret))

}

// SetContentsCenter The rectangle that defines how the layer contents are scaled if the layer’s contents are resized. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410740-contentscenter?language=objc
func (x gen_CALayer) SetContentsCenter(
	value NSRect,
) {
	C.CALayer_inst_setContentsCenter(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return

}

// IsHidden A Boolean indicating whether the layer is displayed. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410838-hidden?language=objc
func (x gen_CALayer) IsHidden() bool {
	ret := C.CALayer_inst_isHidden(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetHidden A Boolean indicating whether the layer is displayed. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410838-hidden?language=objc
func (x gen_CALayer) SetHidden(
	value bool,
) {
	C.CALayer_inst_setHidden(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// MasksToBounds A Boolean indicating whether sublayers are clipped to the layer’s bounds. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410896-maskstobounds?language=objc
func (x gen_CALayer) MasksToBounds() bool {
	ret := C.CALayer_inst_masksToBounds(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetMasksToBounds A Boolean indicating whether sublayers are clipped to the layer’s bounds. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410896-maskstobounds?language=objc
func (x gen_CALayer) SetMasksToBounds(
	value bool,
) {
	C.CALayer_inst_setMasksToBounds(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// Mask An optional layer whose alpha channel is used to mask the layer’s content.
// https://developer.apple.com/documentation/quartzcore/calayer/1410861-mask?language=objc
func (x gen_CALayer) Mask() CALayer {
	ret := C.CALayer_inst_mask(
		unsafe.Pointer(x.Pointer()),
	)

	return CALayer_fromPointer(ret)

}

// SetMask An optional layer whose alpha channel is used to mask the layer’s content.
// https://developer.apple.com/documentation/quartzcore/calayer/1410861-mask?language=objc
func (x gen_CALayer) SetMask(
	value CALayerRef,
) {
	C.CALayer_inst_setMask(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// IsDoubleSided A Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410924-doublesided?language=objc
func (x gen_CALayer) IsDoubleSided() bool {
	ret := C.CALayer_inst_isDoubleSided(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetDoubleSided A Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410924-doublesided?language=objc
func (x gen_CALayer) SetDoubleSided(
	value bool,
) {
	C.CALayer_inst_setDoubleSided(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// CornerRadius The radius to use when drawing rounded corners for the layer’s background. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410818-cornerradius?language=objc
func (x gen_CALayer) CornerRadius() CGFloat {
	ret := C.CALayer_inst_cornerRadius(
		unsafe.Pointer(x.Pointer()),
	)

	return CGFloat(ret)

}

// SetCornerRadius The radius to use when drawing rounded corners for the layer’s background. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410818-cornerradius?language=objc
func (x gen_CALayer) SetCornerRadius(
	value CGFloat,
) {
	C.CALayer_inst_setCornerRadius(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// BorderWidth The width of the layer’s border. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410917-borderwidth?language=objc
func (x gen_CALayer) BorderWidth() CGFloat {
	ret := C.CALayer_inst_borderWidth(
		unsafe.Pointer(x.Pointer()),
	)

	return CGFloat(ret)

}

// SetBorderWidth The width of the layer’s border. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410917-borderwidth?language=objc
func (x gen_CALayer) SetBorderWidth(
	value CGFloat,
) {
	C.CALayer_inst_setBorderWidth(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// ShadowRadius The blur radius (in points) used to render the layer’s shadow. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410819-shadowradius?language=objc
func (x gen_CALayer) ShadowRadius() CGFloat {
	ret := C.CALayer_inst_shadowRadius(
		unsafe.Pointer(x.Pointer()),
	)

	return CGFloat(ret)

}

// SetShadowRadius The blur radius (in points) used to render the layer’s shadow. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410819-shadowradius?language=objc
func (x gen_CALayer) SetShadowRadius(
	value CGFloat,
) {
	C.CALayer_inst_setShadowRadius(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// ShadowOffset The offset (in points) of the layer’s shadow. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410970-shadowoffset?language=objc
func (x gen_CALayer) ShadowOffset() NSSize {
	ret := C.CALayer_inst_shadowOffset(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSSize)(unsafe.Pointer(&ret))

}

// SetShadowOffset The offset (in points) of the layer’s shadow. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410970-shadowoffset?language=objc
func (x gen_CALayer) SetShadowOffset(
	value NSSize,
) {
	C.CALayer_inst_setShadowOffset(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)

	return

}

// Style An optional dictionary used to store property values that aren't explicitly defined by the layer.
// https://developer.apple.com/documentation/quartzcore/calayer/1410875-style?language=objc
func (x gen_CALayer) Style() NSDictionary {
	ret := C.CALayer_inst_style(
		unsafe.Pointer(x.Pointer()),
	)

	return NSDictionary_fromPointer(ret)

}

// SetStyle An optional dictionary used to store property values that aren't explicitly defined by the layer.
// https://developer.apple.com/documentation/quartzcore/calayer/1410875-style?language=objc
func (x gen_CALayer) SetStyle(
	value NSDictionaryRef,
) {
	C.CALayer_inst_setStyle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// AllowsEdgeAntialiasing A Boolean indicating whether the layer is allowed to perform edge antialiasing.
// https://developer.apple.com/documentation/quartzcore/calayer/1621285-allowsedgeantialiasing?language=objc
func (x gen_CALayer) AllowsEdgeAntialiasing() bool {
	ret := C.CALayer_inst_allowsEdgeAntialiasing(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsEdgeAntialiasing A Boolean indicating whether the layer is allowed to perform edge antialiasing.
// https://developer.apple.com/documentation/quartzcore/calayer/1621285-allowsedgeantialiasing?language=objc
func (x gen_CALayer) SetAllowsEdgeAntialiasing(
	value bool,
) {
	C.CALayer_inst_setAllowsEdgeAntialiasing(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AllowsGroupOpacity A Boolean indicating whether the layer is allowed to composite itself as a group separate from its parent.
// https://developer.apple.com/documentation/quartzcore/calayer/1621277-allowsgroupopacity?language=objc
func (x gen_CALayer) AllowsGroupOpacity() bool {
	ret := C.CALayer_inst_allowsGroupOpacity(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsGroupOpacity A Boolean indicating whether the layer is allowed to composite itself as a group separate from its parent.
// https://developer.apple.com/documentation/quartzcore/calayer/1621277-allowsgroupopacity?language=objc
func (x gen_CALayer) SetAllowsGroupOpacity(
	value bool,
) {
	C.CALayer_inst_setAllowsGroupOpacity(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// Filters An array of Core Image filters to apply to the contents of the layer and its sublayers. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410901-filters?language=objc
func (x gen_CALayer) Filters() NSArray {
	ret := C.CALayer_inst_filters(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_fromPointer(ret)

}

// SetFilters An array of Core Image filters to apply to the contents of the layer and its sublayers. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410901-filters?language=objc
func (x gen_CALayer) SetFilters(
	value NSArrayRef,
) {
	C.CALayer_inst_setFilters(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// CompositingFilter A CoreImage filter used to composite the layer and the content behind it. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410748-compositingfilter?language=objc
func (x gen_CALayer) CompositingFilter() objc.Object {
	ret := C.CALayer_inst_compositingFilter(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetCompositingFilter A CoreImage filter used to composite the layer and the content behind it. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410748-compositingfilter?language=objc
func (x gen_CALayer) SetCompositingFilter(
	value objc.Ref,
) {
	C.CALayer_inst_setCompositingFilter(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// BackgroundFilters An array of Core Image filters to apply to the content immediately behind the layer. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410827-backgroundfilters?language=objc
func (x gen_CALayer) BackgroundFilters() NSArray {
	ret := C.CALayer_inst_backgroundFilters(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_fromPointer(ret)

}

// SetBackgroundFilters An array of Core Image filters to apply to the content immediately behind the layer. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410827-backgroundfilters?language=objc
func (x gen_CALayer) SetBackgroundFilters(
	value NSArrayRef,
) {
	C.CALayer_inst_setBackgroundFilters(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// IsOpaque A Boolean value indicating whether the layer contains completely opaque content.
// https://developer.apple.com/documentation/quartzcore/calayer/1410763-opaque?language=objc
func (x gen_CALayer) IsOpaque() bool {
	ret := C.CALayer_inst_isOpaque(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetOpaque A Boolean value indicating whether the layer contains completely opaque content.
// https://developer.apple.com/documentation/quartzcore/calayer/1410763-opaque?language=objc
func (x gen_CALayer) SetOpaque(
	value bool,
) {
	C.CALayer_inst_setOpaque(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsGeometryFlipped A Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically.
// https://developer.apple.com/documentation/quartzcore/calayer/1410960-geometryflipped?language=objc
func (x gen_CALayer) IsGeometryFlipped() bool {
	ret := C.CALayer_inst_isGeometryFlipped(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetGeometryFlipped A Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically.
// https://developer.apple.com/documentation/quartzcore/calayer/1410960-geometryflipped?language=objc
func (x gen_CALayer) SetGeometryFlipped(
	value bool,
) {
	C.CALayer_inst_setGeometryFlipped(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// DrawsAsynchronously A Boolean indicating whether drawing commands are deferred and processed asynchronously in a background thread.
// https://developer.apple.com/documentation/quartzcore/calayer/1410974-drawsasynchronously?language=objc
func (x gen_CALayer) DrawsAsynchronously() bool {
	ret := C.CALayer_inst_drawsAsynchronously(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetDrawsAsynchronously A Boolean indicating whether drawing commands are deferred and processed asynchronously in a background thread.
// https://developer.apple.com/documentation/quartzcore/calayer/1410974-drawsasynchronously?language=objc
func (x gen_CALayer) SetDrawsAsynchronously(
	value bool,
) {
	C.CALayer_inst_setDrawsAsynchronously(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// ShouldRasterize A Boolean that indicates whether the layer is rendered as a bitmap before compositing. Animatable
// https://developer.apple.com/documentation/quartzcore/calayer/1410905-shouldrasterize?language=objc
func (x gen_CALayer) ShouldRasterize() bool {
	ret := C.CALayer_inst_shouldRasterize(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetShouldRasterize A Boolean that indicates whether the layer is rendered as a bitmap before compositing. Animatable
// https://developer.apple.com/documentation/quartzcore/calayer/1410905-shouldrasterize?language=objc
func (x gen_CALayer) SetShouldRasterize(
	value bool,
) {
	C.CALayer_inst_setShouldRasterize(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// RasterizationScale The scale at which to rasterize content, relative to the coordinate space of the layer. Animatable
// https://developer.apple.com/documentation/quartzcore/calayer/1410801-rasterizationscale?language=objc
func (x gen_CALayer) RasterizationScale() CGFloat {
	ret := C.CALayer_inst_rasterizationScale(
		unsafe.Pointer(x.Pointer()),
	)

	return CGFloat(ret)

}

// SetRasterizationScale The scale at which to rasterize content, relative to the coordinate space of the layer. Animatable
// https://developer.apple.com/documentation/quartzcore/calayer/1410801-rasterizationscale?language=objc
func (x gen_CALayer) SetRasterizationScale(
	value CGFloat,
) {
	C.CALayer_inst_setRasterizationScale(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// Frame The layer’s frame rectangle.
// https://developer.apple.com/documentation/quartzcore/calayer/1410779-frame?language=objc
func (x gen_CALayer) Frame() NSRect {
	ret := C.CALayer_inst_frame(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSRect)(unsafe.Pointer(&ret))

}

// SetFrame The layer’s frame rectangle.
// https://developer.apple.com/documentation/quartzcore/calayer/1410779-frame?language=objc
func (x gen_CALayer) SetFrame(
	value NSRect,
) {
	C.CALayer_inst_setFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return

}

// Bounds The layer’s bounds rectangle. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410915-bounds?language=objc
func (x gen_CALayer) Bounds() NSRect {
	ret := C.CALayer_inst_bounds(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSRect)(unsafe.Pointer(&ret))

}

// SetBounds The layer’s bounds rectangle. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410915-bounds?language=objc
func (x gen_CALayer) SetBounds(
	value NSRect,
) {
	C.CALayer_inst_setBounds(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)

	return

}

// ZPosition The layer’s position on the z axis. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410884-zposition?language=objc
func (x gen_CALayer) ZPosition() CGFloat {
	ret := C.CALayer_inst_zPosition(
		unsafe.Pointer(x.Pointer()),
	)

	return CGFloat(ret)

}

// SetZPosition The layer’s position on the z axis. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410884-zposition?language=objc
func (x gen_CALayer) SetZPosition(
	value CGFloat,
) {
	C.CALayer_inst_setZPosition(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// AnchorPointZ The anchor point for the layer’s position along the z axis. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410796-anchorpointz?language=objc
func (x gen_CALayer) AnchorPointZ() CGFloat {
	ret := C.CALayer_inst_anchorPointZ(
		unsafe.Pointer(x.Pointer()),
	)

	return CGFloat(ret)

}

// SetAnchorPointZ The anchor point for the layer’s position along the z axis. Animatable.
// https://developer.apple.com/documentation/quartzcore/calayer/1410796-anchorpointz?language=objc
func (x gen_CALayer) SetAnchorPointZ(
	value CGFloat,
) {
	C.CALayer_inst_setAnchorPointZ(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// ContentsScale The scale factor applied to the layer.
// https://developer.apple.com/documentation/quartzcore/calayer/1410746-contentsscale?language=objc
func (x gen_CALayer) ContentsScale() CGFloat {
	ret := C.CALayer_inst_contentsScale(
		unsafe.Pointer(x.Pointer()),
	)

	return CGFloat(ret)

}

// SetContentsScale The scale factor applied to the layer.
// https://developer.apple.com/documentation/quartzcore/calayer/1410746-contentsscale?language=objc
func (x gen_CALayer) SetContentsScale(
	value CGFloat,
) {
	C.CALayer_inst_setContentsScale(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// Sublayers An array containing the layer’s sublayers.
// https://developer.apple.com/documentation/quartzcore/calayer/1410802-sublayers?language=objc
func (x gen_CALayer) Sublayers() NSArray {
	ret := C.CALayer_inst_sublayers(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_fromPointer(ret)

}

// SetSublayers An array containing the layer’s sublayers.
// https://developer.apple.com/documentation/quartzcore/calayer/1410802-sublayers?language=objc
func (x gen_CALayer) SetSublayers(
	value NSArrayRef,
) {
	C.CALayer_inst_setSublayers(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Superlayer The superlayer of the layer.
// https://developer.apple.com/documentation/quartzcore/calayer/1410761-superlayer?language=objc
func (x gen_CALayer) Superlayer() CALayer {
	ret := C.CALayer_inst_superlayer(
		unsafe.Pointer(x.Pointer()),
	)

	return CALayer_fromPointer(ret)

}

// NeedsDisplayOnBoundsChange A Boolean indicating whether the layer contents must be updated when its bounds rectangle changes.
// https://developer.apple.com/documentation/quartzcore/calayer/1410923-needsdisplayonboundschange?language=objc
func (x gen_CALayer) NeedsDisplayOnBoundsChange() bool {
	ret := C.CALayer_inst_needsDisplayOnBoundsChange(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetNeedsDisplayOnBoundsChange A Boolean indicating whether the layer contents must be updated when its bounds rectangle changes.
// https://developer.apple.com/documentation/quartzcore/calayer/1410923-needsdisplayonboundschange?language=objc
func (x gen_CALayer) SetNeedsDisplayOnBoundsChange(
	value bool,
) {
	C.CALayer_inst_setNeedsDisplayOnBoundsChange(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// LayoutManager The object responsible for laying out the layer’s sublayers.
// https://developer.apple.com/documentation/quartzcore/calayer/1410749-layoutmanager?language=objc
func (x gen_CALayer) LayoutManager() objc.Object {
	ret := C.CALayer_inst_layoutManager(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetLayoutManager The object responsible for laying out the layer’s sublayers.
// https://developer.apple.com/documentation/quartzcore/calayer/1410749-layoutmanager?language=objc
func (x gen_CALayer) SetLayoutManager(
	value objc.Ref,
) {
	C.CALayer_inst_setLayoutManager(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Constraints The constraints used to position current layer’s sublayers.
// https://developer.apple.com/documentation/quartzcore/calayer/1521906-constraints?language=objc
func (x gen_CALayer) Constraints() NSArray {
	ret := C.CALayer_inst_constraints(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_fromPointer(ret)

}

// SetConstraints The constraints used to position current layer’s sublayers.
// https://developer.apple.com/documentation/quartzcore/calayer/1521906-constraints?language=objc
func (x gen_CALayer) SetConstraints(
	value NSArrayRef,
) {
	C.CALayer_inst_setConstraints(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// Actions A dictionary containing layer actions.
// https://developer.apple.com/documentation/quartzcore/calayer/1410789-actions?language=objc
func (x gen_CALayer) Actions() NSDictionary {
	ret := C.CALayer_inst_actions(
		unsafe.Pointer(x.Pointer()),
	)

	return NSDictionary_fromPointer(ret)

}

// SetActions A dictionary containing layer actions.
// https://developer.apple.com/documentation/quartzcore/calayer/1410789-actions?language=objc
func (x gen_CALayer) SetActions(
	value NSDictionaryRef,
) {
	C.CALayer_inst_setActions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// VisibleRect The visible region of the layer in its own coordinate space.
// https://developer.apple.com/documentation/quartzcore/calayer/1521892-visiblerect?language=objc
func (x gen_CALayer) VisibleRect() NSRect {
	ret := C.CALayer_inst_visibleRect(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSRect)(unsafe.Pointer(&ret))

}

// Name The name of the receiver.
// https://developer.apple.com/documentation/quartzcore/calayer/1410879-name?language=objc
func (x gen_CALayer) Name() NSString {
	ret := C.CALayer_inst_name(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// SetName The name of the receiver.
// https://developer.apple.com/documentation/quartzcore/calayer/1410879-name?language=objc
func (x gen_CALayer) SetName(
	value NSStringRef,
) {
	C.CALayer_inst_setName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

type NSArrayRef interface {
	Pointer() uintptr
	Init_asNSArray() NSArray
}

type gen_NSArray struct {
	objc.Object
}

func NSArray_fromPointer(ptr unsafe.Pointer) NSArray {
	return NSArray{gen_NSArray{
		objc.Object_fromPointer(ptr),
	}}
}

func NSArray_fromRef(ref objc.Ref) NSArray {
	return NSArray_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// ArrayByAddingObjectsFromArray Returns a new array that is a copy of the receiving array with the objects contained in another array added to the end.
// https://developer.apple.com/documentation/foundation/nsarray/1412087-arraybyaddingobjectsfromarray?language=objc
func (x gen_NSArray) ArrayByAddingObjectsFromArray(
	otherArray NSArrayRef,
) NSArray {
	ret := C.NSArray_inst_arrayByAddingObjectsFromArray(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(otherArray),
	)

	return NSArray_fromPointer(ret)

}

// ComponentsJoinedByString Constructs and returns an NSString object that is the result of interposing a given separator between the elements of the array.
// https://developer.apple.com/documentation/foundation/nsarray/1412075-componentsjoinedbystring?language=objc
func (x gen_NSArray) ComponentsJoinedByString(
	separator NSStringRef,
) NSString {
	ret := C.NSArray_inst_componentsJoinedByString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(separator),
	)

	return NSString_fromPointer(ret)

}

// DescriptionWithLocale Returns a string that represents the contents of the array, formatted as a property list.
// https://developer.apple.com/documentation/foundation/nsarray/1412374-descriptionwithlocale?language=objc
func (x gen_NSArray) DescriptionWithLocale(
	locale objc.Ref,
) NSString {
	ret := C.NSArray_inst_descriptionWithLocale(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(locale),
	)

	return NSString_fromPointer(ret)

}

// DescriptionWithLocale_indent Returns a string that represents the contents of the array, formatted as a property list.
// https://developer.apple.com/documentation/foundation/nsarray/1416257-descriptionwithlocale?language=objc
func (x gen_NSArray) DescriptionWithLocale_indent(
	locale objc.Ref,
	level NSUInteger,
) NSString {
	ret := C.NSArray_inst_descriptionWithLocale_indent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(locale),
		C.ulong(level),
	)

	return NSString_fromPointer(ret)

}

// Init_asNSArray Initializes a newly allocated array.
// https://developer.apple.com/documentation/foundation/nsarray/1414315-init?language=objc
func (x gen_NSArray) Init_asNSArray() NSArray {
	ret := C.NSArray_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_fromPointer(ret)

}

// InitWithArray_asNSArray Initializes a newly allocated array by placing in it the objects contained in a given array.
// https://developer.apple.com/documentation/foundation/nsarray/1412169-initwitharray?language=objc
func (x gen_NSArray) InitWithArray_asNSArray(
	array NSArrayRef,
) NSArray {
	ret := C.NSArray_inst_initWithArray(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(array),
	)

	return NSArray_fromPointer(ret)

}

// InitWithArray_copyItems_asNSArray Initializes a newly allocated array using anArray as the source of data objects for the array.
// https://developer.apple.com/documentation/foundation/nsarray/1408557-initwitharray?language=objc
func (x gen_NSArray) InitWithArray_copyItems_asNSArray(
	array NSArrayRef,
	flag bool,
) NSArray {
	ret := C.NSArray_inst_initWithArray_copyItems(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(array),
		convertToObjCBool(flag),
	)

	return NSArray_fromPointer(ret)

}

// IsEqualToArray Compares the receiving array to another array.
// https://developer.apple.com/documentation/foundation/nsarray/1411770-isequaltoarray?language=objc
func (x gen_NSArray) IsEqualToArray(
	otherArray NSArrayRef,
) bool {
	ret := C.NSArray_inst_isEqualToArray(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(otherArray),
	)

	return convertObjCBoolToGo(ret)

}

// MakeObjectsPerformSelector Sends to each object in the array the message identified by a given selector, starting with the first object and continuing through the array to the last object.
// https://developer.apple.com/documentation/foundation/nsarray/1460115-makeobjectsperformselector?language=objc
func (x gen_NSArray) MakeObjectsPerformSelector(
	aSelector objc.Selector,
) {
	C.NSArray_inst_makeObjectsPerformSelector(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
	)

	return

}

// MakeObjectsPerformSelector_withObject Sends the aSelector message to each object in the array, starting with the first object and continuing through the array to the last object.
// https://developer.apple.com/documentation/foundation/nsarray/1460107-makeobjectsperformselector?language=objc
func (x gen_NSArray) MakeObjectsPerformSelector_withObject(
	aSelector objc.Selector,
	argument objc.Ref,
) {
	C.NSArray_inst_makeObjectsPerformSelector_withObject(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
		objc.RefPointer(argument),
	)

	return

}

// PathsMatchingExtensions Returns an array containing all the pathname elements in the receiving array that have filename extensions from a given array.
// https://developer.apple.com/documentation/foundation/nsarray/1418275-pathsmatchingextensions?language=objc
func (x gen_NSArray) PathsMatchingExtensions(
	filterTypes NSArrayRef,
) NSArray {
	ret := C.NSArray_inst_pathsMatchingExtensions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(filterTypes),
	)

	return NSArray_fromPointer(ret)

}

// SetValue_forKey Invokes setValue:forKey: on each of the array's items using the specified value and key.
// https://developer.apple.com/documentation/foundation/nsarray/1408301-setvalue?language=objc
func (x gen_NSArray) SetValue_forKey(
	value objc.Ref,
	key NSStringRef,
) {
	C.NSArray_inst_setValue_forKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
		objc.RefPointer(key),
	)

	return

}

// ShuffledArray Returns a new array that lists this array’s elements in a random order.
// https://developer.apple.com/documentation/foundation/nsarray/1640855-shuffledarray?language=objc
func (x gen_NSArray) ShuffledArray() NSArray {
	ret := C.NSArray_inst_shuffledArray(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_fromPointer(ret)

}

// SortedArrayUsingDescriptors Returns a copy of the receiving array sorted as specified by a given array of sort descriptors.
// https://developer.apple.com/documentation/foundation/nsarray/1415069-sortedarrayusingdescriptors?language=objc
func (x gen_NSArray) SortedArrayUsingDescriptors(
	sortDescriptors NSArrayRef,
) NSArray {
	ret := C.NSArray_inst_sortedArrayUsingDescriptors(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sortDescriptors),
	)

	return NSArray_fromPointer(ret)

}

// SortedArrayUsingSelector Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given selector.
// https://developer.apple.com/documentation/foundation/nsarray/1410025-sortedarrayusingselector?language=objc
func (x gen_NSArray) SortedArrayUsingSelector(
	comparator objc.Selector,
) NSArray {
	ret := C.NSArray_inst_sortedArrayUsingSelector(
		unsafe.Pointer(x.Pointer()),
		comparator.SelectorAddress(),
	)

	return NSArray_fromPointer(ret)

}

// ValueForKey Returns an array containing the results of invoking valueForKey: using key on each of the array's objects.
// https://developer.apple.com/documentation/foundation/nsarray/1412219-valueforkey?language=objc
func (x gen_NSArray) ValueForKey(
	key NSStringRef,
) objc.Object {
	ret := C.NSArray_inst_valueForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)

	return objc.Object_fromPointer(ret)

}

// Count The number of objects in the array.
// https://developer.apple.com/documentation/foundation/nsarray/1409982-count?language=objc
func (x gen_NSArray) Count() NSUInteger {
	ret := C.NSArray_inst_count(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)

}

// SortedArrayHint Analyzes the array and returns a “hint” that speeds the sorting of the array when the hint is supplied to sortedArrayUsingFunction:context:hint:.
// https://developer.apple.com/documentation/foundation/nsarray/1413063-sortedarrayhint?language=objc
func (x gen_NSArray) SortedArrayHint() NSData {
	ret := C.NSArray_inst_sortedArrayHint(
		unsafe.Pointer(x.Pointer()),
	)

	return NSData_fromPointer(ret)

}

// Description A string that represents the contents of the array, formatted as a property list.
// https://developer.apple.com/documentation/foundation/nsarray/1413042-description?language=objc
func (x gen_NSArray) Description() NSString {
	ret := C.NSArray_inst_description(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

type NSAttributedStringRef interface {
	Pointer() uintptr
	Init_asNSAttributedString() NSAttributedString
}

type gen_NSAttributedString struct {
	objc.Object
}

func NSAttributedString_fromPointer(ptr unsafe.Pointer) NSAttributedString {
	return NSAttributedString{gen_NSAttributedString{
		objc.Object_fromPointer(ptr),
	}}
}

func NSAttributedString_fromRef(ref objc.Ref) NSAttributedString {
	return NSAttributedString_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// AttributedStringByInflectingString
// https://developer.apple.com/documentation/foundation/nsattributedstring/3746871-attributedstringbyinflectingstri?language=objc
func (x gen_NSAttributedString) AttributedStringByInflectingString() NSAttributedString {
	ret := C.NSAttributedString_inst_attributedStringByInflectingString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSAttributedString_fromPointer(ret)

}

// DrawInRect Draws the attributed string inside the specified bounding rectangle in the current graphics context.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1531631-drawinrect?language=objc
func (x gen_NSAttributedString) DrawInRect(
	rect NSRect,
) {
	C.NSAttributedString_inst_drawInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)

	return

}

// InitWithAttributedString_asNSAttributedString Creates an attributed string with the characters and attributes of the specified attributed string.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1415342-initwithattributedstring?language=objc
func (x gen_NSAttributedString) InitWithAttributedString_asNSAttributedString(
	attrStr NSAttributedStringRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_initWithAttributedString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(attrStr),
	)

	return NSAttributedString_fromPointer(ret)

}

// InitWithDocFormat_documentAttributes_asNSAttributedString Creates an attributed string from Microsoft Word format data in the specified data object.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1534329-initwithdocformat?language=objc
func (x gen_NSAttributedString) InitWithDocFormat_documentAttributes_asNSAttributedString(
	data NSDataRef,
	dict NSDictionaryRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_initWithDocFormat_documentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(dict),
	)

	return NSAttributedString_fromPointer(ret)

}

// InitWithHTML_baseURL_documentAttributes_asNSAttributedString Creates an attributed string from the HTML in the specified data object and base URL.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1524624-initwithhtml?language=objc
func (x gen_NSAttributedString) InitWithHTML_baseURL_documentAttributes_asNSAttributedString(
	data NSDataRef,
	base NSURLRef,
	dict NSDictionaryRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_initWithHTML_baseURL_documentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(base),
		objc.RefPointer(dict),
	)

	return NSAttributedString_fromPointer(ret)

}

// InitWithHTML_documentAttributes_asNSAttributedString Creates an attributed string from the HTML in the specified data object.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1525953-initwithhtml?language=objc
func (x gen_NSAttributedString) InitWithHTML_documentAttributes_asNSAttributedString(
	data NSDataRef,
	dict NSDictionaryRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_initWithHTML_documentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(dict),
	)

	return NSAttributedString_fromPointer(ret)

}

// InitWithHTML_options_documentAttributes_asNSAttributedString Creates an attributed string from the HTML in the specified data object.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1535412-initwithhtml?language=objc
func (x gen_NSAttributedString) InitWithHTML_options_documentAttributes_asNSAttributedString(
	data NSDataRef,
	options NSDictionaryRef,
	dict NSDictionaryRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_initWithHTML_options_documentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(options),
		objc.RefPointer(dict),
	)

	return NSAttributedString_fromPointer(ret)

}

// InitWithRTF_documentAttributes_asNSAttributedString Creates an attributed string by decoding the stream of RTF commands and data in the specified data object.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1532912-initwithrtf?language=objc
func (x gen_NSAttributedString) InitWithRTF_documentAttributes_asNSAttributedString(
	data NSDataRef,
	dict NSDictionaryRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_initWithRTF_documentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(dict),
	)

	return NSAttributedString_fromPointer(ret)

}

// InitWithRTFD_documentAttributes_asNSAttributedString Creates an attributed string by decoding the stream of RTFD commands and data in the specified data object.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1530987-initwithrtfd?language=objc
func (x gen_NSAttributedString) InitWithRTFD_documentAttributes_asNSAttributedString(
	data NSDataRef,
	dict NSDictionaryRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_initWithRTFD_documentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(dict),
	)

	return NSAttributedString_fromPointer(ret)

}

// InitWithString_asNSAttributedString Creates an attributed string with the characters of the specified string and no attribute information.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1407481-initwithstring?language=objc
func (x gen_NSAttributedString) InitWithString_asNSAttributedString(
	str NSStringRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_initWithString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return NSAttributedString_fromPointer(ret)

}

// InitWithString_attributes_asNSAttributedString Creates an attributed string with the specified string and attributes.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1408136-initwithstring?language=objc
func (x gen_NSAttributedString) InitWithString_attributes_asNSAttributedString(
	str NSStringRef,
	attrs NSDictionaryRef,
) NSAttributedString {
	ret := C.NSAttributedString_inst_initWithString_attributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
		objc.RefPointer(attrs),
	)

	return NSAttributedString_fromPointer(ret)

}

// IsEqualToAttributedString Returns a Boolean value that indicates whether the attributed string is equal to another attributed string.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1414808-isequaltoattributedstring?language=objc
func (x gen_NSAttributedString) IsEqualToAttributedString(
	other NSAttributedStringRef,
) bool {
	ret := C.NSAttributedString_inst_isEqualToAttributedString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(other),
	)

	return convertObjCBoolToGo(ret)

}

// NextWordFromIndex_forward Returns the index of the first character of the word after or before the specified index.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1535305-nextwordfromindex?language=objc
func (x gen_NSAttributedString) NextWordFromIndex_forward(
	location NSUInteger,
	isForward bool,
) NSUInteger {
	ret := C.NSAttributedString_inst_nextWordFromIndex_forward(
		unsafe.Pointer(x.Pointer()),
		C.ulong(location),
		convertToObjCBool(isForward),
	)

	return NSUInteger(ret)

}

// Size Returns the size necessary to draw the string.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1528362-size?language=objc
func (x gen_NSAttributedString) Size() NSSize {
	ret := C.NSAttributedString_inst_size(
		unsafe.Pointer(x.Pointer()),
	)

	return *(*NSSize)(unsafe.Pointer(&ret))

}

// Init_asNSAttributedString
func (x gen_NSAttributedString) Init_asNSAttributedString() NSAttributedString {
	ret := C.NSAttributedString_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSAttributedString_fromPointer(ret)

}

// String The character contents of the attributed string as a string.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1412616-string?language=objc
func (x gen_NSAttributedString) String() NSString {
	ret := C.NSAttributedString_inst_string(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// Length The length of the attributed string.
// https://developer.apple.com/documentation/foundation/nsattributedstring/1418432-length?language=objc
func (x gen_NSAttributedString) Length() NSUInteger {
	ret := C.NSAttributedString_inst_length(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)

}

type NSDataRef interface {
	Pointer() uintptr
	Init_asNSData() NSData
}

type gen_NSData struct {
	objc.Object
}

func NSData_fromPointer(ptr unsafe.Pointer) NSData {
	return NSData{gen_NSData{
		objc.Object_fromPointer(ptr),
	}}
}

func NSData_fromRef(ref objc.Ref) NSData {
	return NSData_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// GetBytes_length Copies a number of bytes from the start of the data object into a given buffer.
// https://developer.apple.com/documentation/foundation/nsdata/1411450-getbytes?language=objc
func (x gen_NSData) GetBytes_length(
	buffer unsafe.Pointer,
	length NSUInteger,
) {
	C.NSData_inst_getBytes_length(
		unsafe.Pointer(x.Pointer()),
		buffer,
		C.ulong(length),
	)

	return

}

// InitWithBytes_length_asNSData Initializes a data object filled with a given number of bytes copied from a given buffer.
// https://developer.apple.com/documentation/foundation/nsdata/1412793-initwithbytes?language=objc
func (x gen_NSData) InitWithBytes_length_asNSData(
	bytes unsafe.Pointer,
	length NSUInteger,
) NSData {
	ret := C.NSData_inst_initWithBytes_length(
		unsafe.Pointer(x.Pointer()),
		bytes,
		C.ulong(length),
	)

	return NSData_fromPointer(ret)

}

// InitWithBytesNoCopy_length_asNSData Initializes a data object filled with a given number of bytes of data from a given buffer.
// https://developer.apple.com/documentation/foundation/nsdata/1409454-initwithbytesnocopy?language=objc
func (x gen_NSData) InitWithBytesNoCopy_length_asNSData(
	bytes unsafe.Pointer,
	length NSUInteger,
) NSData {
	ret := C.NSData_inst_initWithBytesNoCopy_length(
		unsafe.Pointer(x.Pointer()),
		bytes,
		C.ulong(length),
	)

	return NSData_fromPointer(ret)

}

// InitWithBytesNoCopy_length_freeWhenDone_asNSData Initializes a newly allocated data object by adding the given number of bytes from the given buffer.
// https://developer.apple.com/documentation/foundation/nsdata/1416020-initwithbytesnocopy?language=objc
func (x gen_NSData) InitWithBytesNoCopy_length_freeWhenDone_asNSData(
	bytes unsafe.Pointer,
	length NSUInteger,
	b bool,
) NSData {
	ret := C.NSData_inst_initWithBytesNoCopy_length_freeWhenDone(
		unsafe.Pointer(x.Pointer()),
		bytes,
		C.ulong(length),
		convertToObjCBool(b),
	)

	return NSData_fromPointer(ret)

}

// InitWithContentsOfFile_asNSData Initializes a data object with the content of the file at a given path.
// https://developer.apple.com/documentation/foundation/nsdata/1408672-initwithcontentsoffile?language=objc
func (x gen_NSData) InitWithContentsOfFile_asNSData(
	path NSStringRef,
) NSData {
	ret := C.NSData_inst_initWithContentsOfFile(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)

	return NSData_fromPointer(ret)

}

// InitWithContentsOfURL_asNSData Initializes a data object with the data from the location specified by a given URL.
// https://developer.apple.com/documentation/foundation/nsdata/1413892-initwithcontentsofurl?language=objc
func (x gen_NSData) InitWithContentsOfURL_asNSData(
	url NSURLRef,
) NSData {
	ret := C.NSData_inst_initWithContentsOfURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)

	return NSData_fromPointer(ret)

}

// InitWithData_asNSData Initializes a data object with the contents of another data object.
// https://developer.apple.com/documentation/foundation/nsdata/1417055-initwithdata?language=objc
func (x gen_NSData) InitWithData_asNSData(
	data NSDataRef,
) NSData {
	ret := C.NSData_inst_initWithData(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
	)

	return NSData_fromPointer(ret)

}

// IsEqualToData Returns a Boolean value indicating whether this data object is the same as another.
// https://developer.apple.com/documentation/foundation/nsdata/1409330-isequaltodata?language=objc
func (x gen_NSData) IsEqualToData(
	other NSDataRef,
) bool {
	ret := C.NSData_inst_isEqualToData(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(other),
	)

	return convertObjCBoolToGo(ret)

}

// WriteToFile_atomically Writes the data object's bytes to the file specified by a given path.
// https://developer.apple.com/documentation/foundation/nsdata/1408033-writetofile?language=objc
func (x gen_NSData) WriteToFile_atomically(
	path NSStringRef,
	useAuxiliaryFile bool,
) bool {
	ret := C.NSData_inst_writeToFile_atomically(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(useAuxiliaryFile),
	)

	return convertObjCBoolToGo(ret)

}

// WriteToURL_atomically Writes the data object's bytes to the location specified by a given URL.
// https://developer.apple.com/documentation/foundation/nsdata/1415134-writetourl?language=objc
func (x gen_NSData) WriteToURL_atomically(
	url NSURLRef,
	atomically bool,
) bool {
	ret := C.NSData_inst_writeToURL_atomically(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
		convertToObjCBool(atomically),
	)

	return convertObjCBoolToGo(ret)

}

// Init_asNSData
func (x gen_NSData) Init_asNSData() NSData {
	ret := C.NSData_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSData_fromPointer(ret)

}

// Bytes A pointer to the data object's contents.
// https://developer.apple.com/documentation/foundation/nsdata/1410616-bytes?language=objc
func (x gen_NSData) Bytes() unsafe.Pointer {
	ret := C.NSData_inst_bytes(
		unsafe.Pointer(x.Pointer()),
	)

	return ret

}

// Length The number of bytes contained by the data object.
// https://developer.apple.com/documentation/foundation/nsdata/1416769-length?language=objc
func (x gen_NSData) Length() NSUInteger {
	ret := C.NSData_inst_length(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)

}

// Description A string that contains a hexadecimal representation of the data object’s contents in a property list format.
// https://developer.apple.com/documentation/foundation/nsdata/1412579-description?language=objc
func (x gen_NSData) Description() NSString {
	ret := C.NSData_inst_description(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

type NSDictionaryRef interface {
	Pointer() uintptr
	Init_asNSDictionary() NSDictionary
}

type gen_NSDictionary struct {
	objc.Object
}

func NSDictionary_fromPointer(ptr unsafe.Pointer) NSDictionary {
	return NSDictionary{gen_NSDictionary{
		objc.Object_fromPointer(ptr),
	}}
}

func NSDictionary_fromRef(ref objc.Ref) NSDictionary {
	return NSDictionary_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// DescriptionWithLocale Returns a string object that represents the contents of the dictionary, formatted as a property list.
// https://developer.apple.com/documentation/foundation/nsdictionary/1417665-descriptionwithlocale?language=objc
func (x gen_NSDictionary) DescriptionWithLocale(
	locale objc.Ref,
) NSString {
	ret := C.NSDictionary_inst_descriptionWithLocale(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(locale),
	)

	return NSString_fromPointer(ret)

}

// DescriptionWithLocale_indent Returns a string object that represents the contents of the dictionary, formatted as a property list.
// https://developer.apple.com/documentation/foundation/nsdictionary/1412690-descriptionwithlocale?language=objc
func (x gen_NSDictionary) DescriptionWithLocale_indent(
	locale objc.Ref,
	level NSUInteger,
) NSString {
	ret := C.NSDictionary_inst_descriptionWithLocale_indent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(locale),
		C.ulong(level),
	)

	return NSString_fromPointer(ret)

}

// FileExtensionHidden Returns a Boolean value indicating whether the file hides its extension.
// https://developer.apple.com/documentation/foundation/nsdictionary/1413177-fileextensionhidden?language=objc
func (x gen_NSDictionary) FileExtensionHidden() bool {
	ret := C.NSDictionary_inst_fileExtensionHidden(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// FileGroupOwnerAccountID Returns file’s group owner account ID.
// https://developer.apple.com/documentation/foundation/nsdictionary/1413626-filegroupowneraccountid?language=objc
func (x gen_NSDictionary) FileGroupOwnerAccountID() NSNumber {
	ret := C.NSDictionary_inst_fileGroupOwnerAccountID(
		unsafe.Pointer(x.Pointer()),
	)

	return NSNumber_fromPointer(ret)

}

// FileGroupOwnerAccountName Returns the file’s group owner account name.
// https://developer.apple.com/documentation/foundation/nsdictionary/1416788-filegroupowneraccountname?language=objc
func (x gen_NSDictionary) FileGroupOwnerAccountName() NSString {
	ret := C.NSDictionary_inst_fileGroupOwnerAccountName(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// FileIsAppendOnly Returns a Boolean value indicating whether the file is append only.
// https://developer.apple.com/documentation/foundation/nsdictionary/1416083-fileisappendonly?language=objc
func (x gen_NSDictionary) FileIsAppendOnly() bool {
	ret := C.NSDictionary_inst_fileIsAppendOnly(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// FileIsImmutable Returns a Boolean value indicating whether the file is immutable.
// https://developer.apple.com/documentation/foundation/nsdictionary/1416500-fileisimmutable?language=objc
func (x gen_NSDictionary) FileIsImmutable() bool {
	ret := C.NSDictionary_inst_fileIsImmutable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// FileOwnerAccountID Returns the file’s owner account ID.
// https://developer.apple.com/documentation/foundation/nsdictionary/1412281-fileowneraccountid?language=objc
func (x gen_NSDictionary) FileOwnerAccountID() NSNumber {
	ret := C.NSDictionary_inst_fileOwnerAccountID(
		unsafe.Pointer(x.Pointer()),
	)

	return NSNumber_fromPointer(ret)

}

// FileOwnerAccountName Returns the file’s owner account name.
// https://developer.apple.com/documentation/foundation/nsdictionary/1417533-fileowneraccountname?language=objc
func (x gen_NSDictionary) FileOwnerAccountName() NSString {
	ret := C.NSDictionary_inst_fileOwnerAccountName(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// FilePosixPermissions Returns the file’s POSIX permissions.
// https://developer.apple.com/documentation/foundation/nsdictionary/1409446-fileposixpermissions?language=objc
func (x gen_NSDictionary) FilePosixPermissions() NSUInteger {
	ret := C.NSDictionary_inst_filePosixPermissions(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)

}

// FileSystemFileNumber Returns the filesystem file number.
// https://developer.apple.com/documentation/foundation/nsdictionary/1408396-filesystemfilenumber?language=objc
func (x gen_NSDictionary) FileSystemFileNumber() NSUInteger {
	ret := C.NSDictionary_inst_fileSystemFileNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)

}

// FileSystemNumber Returns the filesystem number.
// https://developer.apple.com/documentation/foundation/nsdictionary/1415329-filesystemnumber?language=objc
func (x gen_NSDictionary) FileSystemNumber() NSInteger {
	ret := C.NSDictionary_inst_fileSystemNumber(
		unsafe.Pointer(x.Pointer()),
	)

	return NSInteger(ret)

}

// FileType Returns the file type.
// https://developer.apple.com/documentation/foundation/nsdictionary/1416809-filetype?language=objc
func (x gen_NSDictionary) FileType() NSString {
	ret := C.NSDictionary_inst_fileType(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// Init_asNSDictionary Initializes a newly allocated dictionary.
// https://developer.apple.com/documentation/foundation/nsdictionary/1418147-init?language=objc
func (x gen_NSDictionary) Init_asNSDictionary() NSDictionary {
	ret := C.NSDictionary_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSDictionary_fromPointer(ret)

}

// InitWithDictionary_asNSDictionary Initializes a newly allocated dictionary by placing in it the keys and values contained in another given dictionary.
// https://developer.apple.com/documentation/foundation/nsdictionary/1418434-initwithdictionary?language=objc
func (x gen_NSDictionary) InitWithDictionary_asNSDictionary(
	otherDictionary NSDictionaryRef,
) NSDictionary {
	ret := C.NSDictionary_inst_initWithDictionary(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(otherDictionary),
	)

	return NSDictionary_fromPointer(ret)

}

// InitWithDictionary_copyItems_asNSDictionary Initializes a newly allocated dictionary using the objects contained in another given dictionary.
// https://developer.apple.com/documentation/foundation/nsdictionary/1410124-initwithdictionary?language=objc
func (x gen_NSDictionary) InitWithDictionary_copyItems_asNSDictionary(
	otherDictionary NSDictionaryRef,
	flag bool,
) NSDictionary {
	ret := C.NSDictionary_inst_initWithDictionary_copyItems(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(otherDictionary),
		convertToObjCBool(flag),
	)

	return NSDictionary_fromPointer(ret)

}

// InitWithObjects_forKeys_asNSDictionary Initializes a newly allocated dictionary with key-value pairs constructed from the provided arrays of keys and objects.
// https://developer.apple.com/documentation/foundation/nsdictionary/1410010-initwithobjects?language=objc
func (x gen_NSDictionary) InitWithObjects_forKeys_asNSDictionary(
	objects NSArrayRef,
	keys NSArrayRef,
) NSDictionary {
	ret := C.NSDictionary_inst_initWithObjects_forKeys(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(objects),
		objc.RefPointer(keys),
	)

	return NSDictionary_fromPointer(ret)

}

// IsEqualToDictionary Returns a Boolean value that indicates whether the contents of the receiving dictionary are equal to the contents of another given dictionary.
// https://developer.apple.com/documentation/foundation/nsdictionary/1415445-isequaltodictionary?language=objc
func (x gen_NSDictionary) IsEqualToDictionary(
	otherDictionary NSDictionaryRef,
) bool {
	ret := C.NSDictionary_inst_isEqualToDictionary(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(otherDictionary),
	)

	return convertObjCBoolToGo(ret)

}

// KeysSortedByValueUsingSelector Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values.
// https://developer.apple.com/documentation/foundation/nsdictionary/1412484-keyssortedbyvalueusingselector?language=objc
func (x gen_NSDictionary) KeysSortedByValueUsingSelector(
	comparator objc.Selector,
) NSArray {
	ret := C.NSDictionary_inst_keysSortedByValueUsingSelector(
		unsafe.Pointer(x.Pointer()),
		comparator.SelectorAddress(),
	)

	return NSArray_fromPointer(ret)

}

// Count The number of entries in the dictionary.
// https://developer.apple.com/documentation/foundation/nsdictionary/1409628-count?language=objc
func (x gen_NSDictionary) Count() NSUInteger {
	ret := C.NSDictionary_inst_count(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)

}

// AllKeys A new array containing the dictionary’s keys, or an empty array if the dictionary has no entries.
// https://developer.apple.com/documentation/foundation/nsdictionary/1409150-allkeys?language=objc
func (x gen_NSDictionary) AllKeys() NSArray {
	ret := C.NSDictionary_inst_allKeys(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_fromPointer(ret)

}

// AllValues A new array containing the dictionary’s values, or an empty array if the dictionary has no entries.
// https://developer.apple.com/documentation/foundation/nsdictionary/1408915-allvalues?language=objc
func (x gen_NSDictionary) AllValues() NSArray {
	ret := C.NSDictionary_inst_allValues(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_fromPointer(ret)

}

// Description A string that represents the contents of the dictionary, formatted as a property list.
// https://developer.apple.com/documentation/foundation/nsdictionary/1410799-description?language=objc
func (x gen_NSDictionary) Description() NSString {
	ret := C.NSDictionary_inst_description(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// DescriptionInStringsFileFormat A string that represents the contents of the dictionary, formatted in .strings file format.
// https://developer.apple.com/documentation/foundation/nsdictionary/1413282-descriptioninstringsfileformat?language=objc
func (x gen_NSDictionary) DescriptionInStringsFileFormat() NSString {
	ret := C.NSDictionary_inst_descriptionInStringsFileFormat(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

type NSNumberRef interface {
	Pointer() uintptr
	Init_asNSNumber() NSNumber
}

type gen_NSNumber struct {
	objc.Object
}

func NSNumber_fromPointer(ptr unsafe.Pointer) NSNumber {
	return NSNumber{gen_NSNumber{
		objc.Object_fromPointer(ptr),
	}}
}

func NSNumber_fromRef(ref objc.Ref) NSNumber {
	return NSNumber_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// DescriptionWithLocale Returns a string that represents the contents of the number object for a given locale.
// https://developer.apple.com/documentation/foundation/nsnumber/1409984-descriptionwithlocale?language=objc
func (x gen_NSNumber) DescriptionWithLocale(
	locale objc.Ref,
) NSString {
	ret := C.NSNumber_inst_descriptionWithLocale(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(locale),
	)

	return NSString_fromPointer(ret)

}

// InitWithBool Returns an NSNumber object initialized to contain a given value, treated as a BOOL.
// https://developer.apple.com/documentation/foundation/nsnumber/1415728-initwithbool?language=objc
func (x gen_NSNumber) InitWithBool(
	value bool,
) NSNumber {
	ret := C.NSNumber_inst_initWithBool(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return NSNumber_fromPointer(ret)

}

// InitWithInt Returns an NSNumber object initialized to contain a given value, treated as a signed int.
// https://developer.apple.com/documentation/foundation/nsnumber/1407580-initwithint?language=objc
func (x gen_NSNumber) InitWithInt(
	value int32,
) NSNumber {
	ret := C.NSNumber_inst_initWithInt(
		unsafe.Pointer(x.Pointer()),
		C.int(value),
	)

	return NSNumber_fromPointer(ret)

}

// InitWithInteger Returns an NSNumber object initialized to contain a given value, treated as an NSInteger.
// https://developer.apple.com/documentation/foundation/nsnumber/1409397-initwithinteger?language=objc
func (x gen_NSNumber) InitWithInteger(
	value NSInteger,
) NSNumber {
	ret := C.NSNumber_inst_initWithInteger(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)

	return NSNumber_fromPointer(ret)

}

// InitWithUnsignedInt Returns an NSNumber object initialized to contain a given value, treated as an unsigned int.
// https://developer.apple.com/documentation/foundation/nsnumber/1414598-initwithunsignedint?language=objc
func (x gen_NSNumber) InitWithUnsignedInt(
	value int32,
) NSNumber {
	ret := C.NSNumber_inst_initWithUnsignedInt(
		unsafe.Pointer(x.Pointer()),
		C.int(value),
	)

	return NSNumber_fromPointer(ret)

}

// InitWithUnsignedInteger Returns an NSNumber object initialized to contain a given value, treated as an NSUInteger.
// https://developer.apple.com/documentation/foundation/nsnumber/1412531-initwithunsignedinteger?language=objc
func (x gen_NSNumber) InitWithUnsignedInteger(
	value NSUInteger,
) NSNumber {
	ret := C.NSNumber_inst_initWithUnsignedInteger(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)

	return NSNumber_fromPointer(ret)

}

// IsEqualToNumber Returns a Boolean value that indicates whether the number object’s value and a given number are equal.
// https://developer.apple.com/documentation/foundation/nsnumber/1411315-isequaltonumber?language=objc
func (x gen_NSNumber) IsEqualToNumber(
	number NSNumberRef,
) bool {
	ret := C.NSNumber_inst_isEqualToNumber(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(number),
	)

	return convertObjCBoolToGo(ret)

}

// Init_asNSNumber
func (x gen_NSNumber) Init_asNSNumber() NSNumber {
	ret := C.NSNumber_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSNumber_fromPointer(ret)

}

// BoolValue The number object's value expressed as a Boolean value.
// https://developer.apple.com/documentation/foundation/nsnumber/1410865-boolvalue?language=objc
func (x gen_NSNumber) BoolValue() bool {
	ret := C.NSNumber_inst_boolValue(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IntValue The number object's value expressed as an int, converted as necessary.
// https://developer.apple.com/documentation/foundation/nsnumber/1407153-intvalue?language=objc
func (x gen_NSNumber) IntValue() int32 {
	ret := C.NSNumber_inst_intValue(
		unsafe.Pointer(x.Pointer()),
	)

	return int32(ret)

}

// IntegerValue The number object's value expressed as an NSInteger object, converted as necessary.
// https://developer.apple.com/documentation/foundation/nsnumber/1412554-integervalue?language=objc
func (x gen_NSNumber) IntegerValue() NSInteger {
	ret := C.NSNumber_inst_integerValue(
		unsafe.Pointer(x.Pointer()),
	)

	return NSInteger(ret)

}

// UnsignedIntegerValue The number object's value expressed as an NSUInteger object, converted as necessary.
// https://developer.apple.com/documentation/foundation/nsnumber/1413324-unsignedintegervalue?language=objc
func (x gen_NSNumber) UnsignedIntegerValue() NSUInteger {
	ret := C.NSNumber_inst_unsignedIntegerValue(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)

}

// UnsignedIntValue The number object's value expressed as an unsigned int, converted as necessary.
// https://developer.apple.com/documentation/foundation/nsnumber/1417875-unsignedintvalue?language=objc
func (x gen_NSNumber) UnsignedIntValue() int32 {
	ret := C.NSNumber_inst_unsignedIntValue(
		unsafe.Pointer(x.Pointer()),
	)

	return int32(ret)

}

// StringValue The number object's value expressed as a human-readable string.
// https://developer.apple.com/documentation/foundation/nsnumber/1415802-stringvalue?language=objc
func (x gen_NSNumber) StringValue() NSString {
	ret := C.NSNumber_inst_stringValue(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

type NSRunLoopRef interface {
	Pointer() uintptr
	Init_asNSRunLoop() NSRunLoop
}

type gen_NSRunLoop struct {
	objc.Object
}

func NSRunLoop_fromPointer(ptr unsafe.Pointer) NSRunLoop {
	return NSRunLoop{gen_NSRunLoop{
		objc.Object_fromPointer(ptr),
	}}
}

func NSRunLoop_fromRef(ref objc.Ref) NSRunLoop {
	return NSRunLoop_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// CancelPerformSelector_target_argument Cancels the sending of a previously scheduled message.
// https://developer.apple.com/documentation/foundation/nsrunloop/1418077-cancelperformselector?language=objc
func (x gen_NSRunLoop) CancelPerformSelector_target_argument(
	aSelector objc.Selector,
	target objc.Ref,
	arg objc.Ref,
) {
	C.NSRunLoop_inst_cancelPerformSelector_target_argument(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(arg),
	)

	return

}

// CancelPerformSelectorsWithTarget Cancels all outstanding ordered performs scheduled with a given target.
// https://developer.apple.com/documentation/foundation/nsrunloop/1414208-cancelperformselectorswithtarget?language=objc
func (x gen_NSRunLoop) CancelPerformSelectorsWithTarget(
	target objc.Ref,
) {
	C.NSRunLoop_inst_cancelPerformSelectorsWithTarget(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(target),
	)

	return

}

// PerformSelector_target_argument_order_modes Schedules the sending of a message on the receiver.
// https://developer.apple.com/documentation/foundation/nsrunloop/1409310-performselector?language=objc
func (x gen_NSRunLoop) PerformSelector_target_argument_order_modes(
	aSelector objc.Selector,
	target objc.Ref,
	arg objc.Ref,
	order NSUInteger,
	modes NSArrayRef,
) {
	C.NSRunLoop_inst_performSelector_target_argument_order_modes(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
		objc.RefPointer(target),
		objc.RefPointer(arg),
		C.ulong(order),
		objc.RefPointer(modes),
	)

	return

}

// Run Puts the receiver into a permanent loop, during which time it processes data from all attached input sources.
// https://developer.apple.com/documentation/foundation/nsrunloop/1412430-run?language=objc
func (x gen_NSRunLoop) Run() {
	C.NSRunLoop_inst_run(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Init_asNSRunLoop
func (x gen_NSRunLoop) Init_asNSRunLoop() NSRunLoop {
	ret := C.NSRunLoop_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSRunLoop_fromPointer(ret)

}

type NSStringRef interface {
	Pointer() uintptr
	Init_asNSString() NSString
}

type gen_NSString struct {
	objc.Object
}

func NSString_fromPointer(ptr unsafe.Pointer) NSString {
	return NSString{gen_NSString{
		objc.Object_fromPointer(ptr),
	}}
}

func NSString_fromRef(ref objc.Ref) NSString {
	return NSString_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// CanBeConvertedToEncoding Returns a Boolean value that indicates whether the receiver can be converted to a given encoding without loss of information.
// https://developer.apple.com/documentation/foundation/nsstring/1409496-canbeconvertedtoencoding?language=objc
func (x gen_NSString) CanBeConvertedToEncoding(
	encoding NSStringEncoding,
) bool {
	ret := C.NSString_inst_canBeConvertedToEncoding(
		unsafe.Pointer(x.Pointer()),
		C.ulong(encoding),
	)

	return convertObjCBoolToGo(ret)

}

// CharacterAtIndex Returns the character at a given UTF-16 code unit index.
// https://developer.apple.com/documentation/foundation/nsstring/1414645-characteratindex?language=objc
func (x gen_NSString) CharacterAtIndex(
	index NSUInteger,
) Unichar {
	ret := C.NSString_inst_characterAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(index),
	)

	return Unichar(ret)

}

// CompletePathIntoString_caseSensitive_matchesIntoArray_filterTypes Interprets the receiver as a path in the file system and attempts to perform filename completion, returning a numeric value that indicates whether a match was possible, and by reference the longest path that matches the receiver.
// https://developer.apple.com/documentation/foundation/nsstring/1411841-completepathintostring?language=objc
func (x gen_NSString) CompletePathIntoString_caseSensitive_matchesIntoArray_filterTypes(
	outputName NSStringRef,
	flag bool,
	outputArray NSArrayRef,
	filterTypes NSArrayRef,
) NSUInteger {
	ret := C.NSString_inst_completePathIntoString_caseSensitive_matchesIntoArray_filterTypes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(outputName),
		convertToObjCBool(flag),
		objc.RefPointer(outputArray),
		objc.RefPointer(filterTypes),
	)

	return NSUInteger(ret)

}

// ComponentsSeparatedByString Returns an array containing substrings from the receiver that have been divided by a given separator.
// https://developer.apple.com/documentation/foundation/nsstring/1413214-componentsseparatedbystring?language=objc
func (x gen_NSString) ComponentsSeparatedByString(
	separator NSStringRef,
) NSArray {
	ret := C.NSString_inst_componentsSeparatedByString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(separator),
	)

	return NSArray_fromPointer(ret)

}

// ContainsString Returns a Boolean value indicating whether the string contains a given string by performing a case-sensitive, locale-unaware search.
// https://developer.apple.com/documentation/foundation/nsstring/1414563-containsstring?language=objc
func (x gen_NSString) ContainsString(
	str NSStringRef,
) bool {
	ret := C.NSString_inst_containsString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return convertObjCBoolToGo(ret)

}

// DataUsingEncoding Returns an NSData object containing a representation of the receiver encoded using a given encoding.
// https://developer.apple.com/documentation/foundation/nsstring/1416696-datausingencoding?language=objc
func (x gen_NSString) DataUsingEncoding(
	encoding NSStringEncoding,
) NSData {
	ret := C.NSString_inst_dataUsingEncoding(
		unsafe.Pointer(x.Pointer()),
		C.ulong(encoding),
	)

	return NSData_fromPointer(ret)

}

// DataUsingEncoding_allowLossyConversion Returns an NSData object containing a representation of the receiver encoded using a given encoding.
// https://developer.apple.com/documentation/foundation/nsstring/1413692-datausingencoding?language=objc
func (x gen_NSString) DataUsingEncoding_allowLossyConversion(
	encoding NSStringEncoding,
	lossy bool,
) NSData {
	ret := C.NSString_inst_dataUsingEncoding_allowLossyConversion(
		unsafe.Pointer(x.Pointer()),
		C.ulong(encoding),
		convertToObjCBool(lossy),
	)

	return NSData_fromPointer(ret)

}

// DrawInRect_withAttributes Draws the attributed string inside the specified bounding rectangle.
// https://developer.apple.com/documentation/foundation/nsstring/1529855-drawinrect?language=objc
func (x gen_NSString) DrawInRect_withAttributes(
	rect NSRect,
	attrs NSDictionaryRef,
) {
	C.NSString_inst_drawInRect_withAttributes(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
		objc.RefPointer(attrs),
	)

	return

}

// HasPrefix Returns a Boolean value that indicates whether a given string matches the beginning characters of the receiver.
// https://developer.apple.com/documentation/foundation/nsstring/1410309-hasprefix?language=objc
func (x gen_NSString) HasPrefix(
	str NSStringRef,
) bool {
	ret := C.NSString_inst_hasPrefix(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return convertObjCBoolToGo(ret)

}

// HasSuffix Returns a Boolean value that indicates whether a given string matches the ending characters of the receiver.
// https://developer.apple.com/documentation/foundation/nsstring/1416529-hassuffix?language=objc
func (x gen_NSString) HasSuffix(
	str NSStringRef,
) bool {
	ret := C.NSString_inst_hasSuffix(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return convertObjCBoolToGo(ret)

}

// Init_asNSString Returns an initialized NSString object that contains no characters.
// https://developer.apple.com/documentation/foundation/nsstring/1409306-init?language=objc
func (x gen_NSString) Init_asNSString() NSString {
	ret := C.NSString_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// InitWithBytes_length_encoding_asNSString Returns an initialized NSString object containing a given number of bytes from a given buffer of bytes interpreted in a given encoding.
// https://developer.apple.com/documentation/foundation/nsstring/1407339-initwithbytes?language=objc
func (x gen_NSString) InitWithBytes_length_encoding_asNSString(
	bytes unsafe.Pointer,
	len NSUInteger,
	encoding NSStringEncoding,
) NSString {
	ret := C.NSString_inst_initWithBytes_length_encoding(
		unsafe.Pointer(x.Pointer()),
		bytes,
		C.ulong(len),
		C.ulong(encoding),
	)

	return NSString_fromPointer(ret)

}

// InitWithBytesNoCopy_length_encoding_freeWhenDone_asNSString Returns an initialized NSString object that contains a given number of bytes from a given buffer of bytes interpreted in a given encoding, and optionally frees the buffer.
// https://developer.apple.com/documentation/foundation/nsstring/1413830-initwithbytesnocopy?language=objc
func (x gen_NSString) InitWithBytesNoCopy_length_encoding_freeWhenDone_asNSString(
	bytes unsafe.Pointer,
	len NSUInteger,
	encoding NSStringEncoding,
	freeBuffer bool,
) NSString {
	ret := C.NSString_inst_initWithBytesNoCopy_length_encoding_freeWhenDone(
		unsafe.Pointer(x.Pointer()),
		bytes,
		C.ulong(len),
		C.ulong(encoding),
		convertToObjCBool(freeBuffer),
	)

	return NSString_fromPointer(ret)

}

// InitWithData_encoding_asNSString Returns an NSString object initialized by converting given data into UTF-16 code units using a given encoding.
// https://developer.apple.com/documentation/foundation/nsstring/1416374-initwithdata?language=objc
func (x gen_NSString) InitWithData_encoding_asNSString(
	data NSDataRef,
	encoding NSStringEncoding,
) NSString {
	ret := C.NSString_inst_initWithData_encoding(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		C.ulong(encoding),
	)

	return NSString_fromPointer(ret)

}

// InitWithString_asNSString Returns an NSString object initialized by copying the characters from another given string.
// https://developer.apple.com/documentation/foundation/nsstring/1411293-initwithstring?language=objc
func (x gen_NSString) InitWithString_asNSString(
	aString NSStringRef,
) NSString {
	ret := C.NSString_inst_initWithString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(aString),
	)

	return NSString_fromPointer(ret)

}

// IsEqualToString Returns a Boolean value that indicates whether a given string is equal to the receiver using a literal Unicode-based comparison.
// https://developer.apple.com/documentation/foundation/nsstring/1407803-isequaltostring?language=objc
func (x gen_NSString) IsEqualToString(
	aString NSStringRef,
) bool {
	ret := C.NSString_inst_isEqualToString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(aString),
	)

	return convertObjCBoolToGo(ret)

}

// LengthOfBytesUsingEncoding Returns the number of bytes required to store the receiver in a given encoding.
// https://developer.apple.com/documentation/foundation/nsstring/1410710-lengthofbytesusingencoding?language=objc
func (x gen_NSString) LengthOfBytesUsingEncoding(
	enc NSStringEncoding,
) NSUInteger {
	ret := C.NSString_inst_lengthOfBytesUsingEncoding(
		unsafe.Pointer(x.Pointer()),
		C.ulong(enc),
	)

	return NSUInteger(ret)

}

// LocalizedCaseInsensitiveContainsString Returns a Boolean value indicating whether the string contains a given string by performing a case-insensitive, locale-aware search.
// https://developer.apple.com/documentation/foundation/nsstring/1412098-localizedcaseinsensitivecontains?language=objc
func (x gen_NSString) LocalizedCaseInsensitiveContainsString(
	str NSStringRef,
) bool {
	ret := C.NSString_inst_localizedCaseInsensitiveContainsString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return convertObjCBoolToGo(ret)

}

// LocalizedStandardContainsString Returns a Boolean value indicating whether the string contains a given string by performing a case and diacritic insensitive, locale-aware search.
// https://developer.apple.com/documentation/foundation/nsstring/1416328-localizedstandardcontainsstring?language=objc
func (x gen_NSString) LocalizedStandardContainsString(
	str NSStringRef,
) bool {
	ret := C.NSString_inst_localizedStandardContainsString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return convertObjCBoolToGo(ret)

}

// MaximumLengthOfBytesUsingEncoding Returns the maximum number of bytes needed to store the receiver in a given encoding.
// https://developer.apple.com/documentation/foundation/nsstring/1411611-maximumlengthofbytesusingencodin?language=objc
func (x gen_NSString) MaximumLengthOfBytesUsingEncoding(
	enc NSStringEncoding,
) NSUInteger {
	ret := C.NSString_inst_maximumLengthOfBytesUsingEncoding(
		unsafe.Pointer(x.Pointer()),
		C.ulong(enc),
	)

	return NSUInteger(ret)

}

// PropertyList Parses the receiver as a text representation of a property list, returning an NSString, NSData, NSArray, or NSDictionary object, according to the topmost element.
// https://developer.apple.com/documentation/foundation/nsstring/1413115-propertylist?language=objc
func (x gen_NSString) PropertyList() objc.Object {
	ret := C.NSString_inst_propertyList(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// PropertyListFromStringsFileFormat Returns a dictionary object initialized with the keys and values found in the receiver.
// https://developer.apple.com/documentation/foundation/nsstring/1407697-propertylistfromstringsfileforma?language=objc
func (x gen_NSString) PropertyListFromStringsFileFormat() NSDictionary {
	ret := C.NSString_inst_propertyListFromStringsFileFormat(
		unsafe.Pointer(x.Pointer()),
	)

	return NSDictionary_fromPointer(ret)

}

// SizeWithAttributes Returns the bounding box size the receiver occupies when drawn with the given attributes.
// https://developer.apple.com/documentation/foundation/nsstring/1531844-sizewithattributes?language=objc
func (x gen_NSString) SizeWithAttributes(
	attrs NSDictionaryRef,
) NSSize {
	ret := C.NSString_inst_sizeWithAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(attrs),
	)

	return *(*NSSize)(unsafe.Pointer(&ret))

}

// StringByAppendingPathComponent Returns a new string made by appending to the receiver a given string.
// https://developer.apple.com/documentation/foundation/nsstring/1417069-stringbyappendingpathcomponent?language=objc
func (x gen_NSString) StringByAppendingPathComponent(
	str NSStringRef,
) NSString {
	ret := C.NSString_inst_stringByAppendingPathComponent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return NSString_fromPointer(ret)

}

// StringByAppendingPathExtension Returns a new string made by appending to the receiver an extension separator followed by a given extension.
// https://developer.apple.com/documentation/foundation/nsstring/1412501-stringbyappendingpathextension?language=objc
func (x gen_NSString) StringByAppendingPathExtension(
	str NSStringRef,
) NSString {
	ret := C.NSString_inst_stringByAppendingPathExtension(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)

	return NSString_fromPointer(ret)

}

// StringByAppendingString Returns a new string made by appending a given string to the receiver.
// https://developer.apple.com/documentation/foundation/nsstring/1412307-stringbyappendingstring?language=objc
func (x gen_NSString) StringByAppendingString(
	aString NSStringRef,
) NSString {
	ret := C.NSString_inst_stringByAppendingString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(aString),
	)

	return NSString_fromPointer(ret)

}

// StringByPaddingToLength_withString_startingAtIndex Returns a new string formed from the receiver by either removing characters from the end, or by appending as many occurrences as necessary of a given pad string.
// https://developer.apple.com/documentation/foundation/nsstring/1416395-stringbypaddingtolength?language=objc
func (x gen_NSString) StringByPaddingToLength_withString_startingAtIndex(
	newLength NSUInteger,
	padString NSStringRef,
	padIndex NSUInteger,
) NSString {
	ret := C.NSString_inst_stringByPaddingToLength_withString_startingAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(newLength),
		objc.RefPointer(padString),
		C.ulong(padIndex),
	)

	return NSString_fromPointer(ret)

}

// StringByReplacingOccurrencesOfString_withString Returns a new string in which all occurrences of a target string in the receiver are replaced by another given string.
// https://developer.apple.com/documentation/foundation/nsstring/1412937-stringbyreplacingoccurrencesofst?language=objc
func (x gen_NSString) StringByReplacingOccurrencesOfString_withString(
	target NSStringRef,
	replacement NSStringRef,
) NSString {
	ret := C.NSString_inst_stringByReplacingOccurrencesOfString_withString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(target),
		objc.RefPointer(replacement),
	)

	return NSString_fromPointer(ret)

}

// StringsByAppendingPaths Returns an array of strings made by separately appending to the receiver each string in a given array.
// https://developer.apple.com/documentation/foundation/nsstring/1415100-stringsbyappendingpaths?language=objc
func (x gen_NSString) StringsByAppendingPaths(
	paths NSArrayRef,
) NSArray {
	ret := C.NSString_inst_stringsByAppendingPaths(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(paths),
	)

	return NSArray_fromPointer(ret)

}

// SubstringFromIndex Returns a new string containing the characters of the receiver from the one at a given index to the end.
// https://developer.apple.com/documentation/foundation/nsstring/1414368-substringfromindex?language=objc
func (x gen_NSString) SubstringFromIndex(
	from NSUInteger,
) NSString {
	ret := C.NSString_inst_substringFromIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(from),
	)

	return NSString_fromPointer(ret)

}

// SubstringToIndex Returns a new string containing the characters of the receiver up to, but not including, the one at a given index.
// https://developer.apple.com/documentation/foundation/nsstring/1408017-substringtoindex?language=objc
func (x gen_NSString) SubstringToIndex(
	to NSUInteger,
) NSString {
	ret := C.NSString_inst_substringToIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(to),
	)

	return NSString_fromPointer(ret)

}

// VariantFittingPresentationWidth Returns a string variation suitable for the specified presentation width.
// https://developer.apple.com/documentation/foundation/nsstring/1413104-variantfittingpresentationwidth?language=objc
func (x gen_NSString) VariantFittingPresentationWidth(
	width NSInteger,
) NSString {
	ret := C.NSString_inst_variantFittingPresentationWidth(
		unsafe.Pointer(x.Pointer()),
		C.long(width),
	)

	return NSString_fromPointer(ret)

}

// Length The number of UTF-16 code units in the receiver.
// https://developer.apple.com/documentation/foundation/nsstring/1414212-length?language=objc
func (x gen_NSString) Length() NSUInteger {
	ret := C.NSString_inst_length(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)

}

// Hash An unsigned integer that can be used as a hash table address.
// https://developer.apple.com/documentation/foundation/nsstring/1417245-hash?language=objc
func (x gen_NSString) Hash() NSUInteger {
	ret := C.NSString_inst_hash(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)

}

// LowercaseString A lowercase representation of the string.
// https://developer.apple.com/documentation/foundation/nsstring/1408467-lowercasestring?language=objc
func (x gen_NSString) LowercaseString() NSString {
	ret := C.NSString_inst_lowercaseString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// LocalizedLowercaseString Returns a version of the string with all letters converted to lowercase, taking into account the current locale.
// https://developer.apple.com/documentation/foundation/nsstring/1414125-localizedlowercasestring?language=objc
func (x gen_NSString) LocalizedLowercaseString() NSString {
	ret := C.NSString_inst_localizedLowercaseString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// UppercaseString An uppercase representation of the string.
// https://developer.apple.com/documentation/foundation/nsstring/1409855-uppercasestring?language=objc
func (x gen_NSString) UppercaseString() NSString {
	ret := C.NSString_inst_uppercaseString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// LocalizedUppercaseString Returns a version of the string with all letters converted to uppercase, taking into account the current locale.
// https://developer.apple.com/documentation/foundation/nsstring/1413331-localizeduppercasestring?language=objc
func (x gen_NSString) LocalizedUppercaseString() NSString {
	ret := C.NSString_inst_localizedUppercaseString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// CapitalizedString A capitalized representation of the string.
// https://developer.apple.com/documentation/foundation/nsstring/1416784-capitalizedstring?language=objc
func (x gen_NSString) CapitalizedString() NSString {
	ret := C.NSString_inst_capitalizedString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// LocalizedCapitalizedString Returns a capitalized representation of the receiver using the current locale.
// https://developer.apple.com/documentation/foundation/nsstring/1414885-localizedcapitalizedstring?language=objc
func (x gen_NSString) LocalizedCapitalizedString() NSString {
	ret := C.NSString_inst_localizedCapitalizedString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// DecomposedStringWithCanonicalMapping A string made by normalizing the string’s contents using the Unicode Normalization Form D.
// https://developer.apple.com/documentation/foundation/nsstring/1409474-decomposedstringwithcanonicalmap?language=objc
func (x gen_NSString) DecomposedStringWithCanonicalMapping() NSString {
	ret := C.NSString_inst_decomposedStringWithCanonicalMapping(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// DecomposedStringWithCompatibilityMapping A string made by normalizing the receiver’s contents using the Unicode Normalization Form KD.
// https://developer.apple.com/documentation/foundation/nsstring/1415417-decomposedstringwithcompatibilit?language=objc
func (x gen_NSString) DecomposedStringWithCompatibilityMapping() NSString {
	ret := C.NSString_inst_decomposedStringWithCompatibilityMapping(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// PrecomposedStringWithCanonicalMapping A string made by normalizing the string’s contents using the Unicode Normalization Form C.
// https://developer.apple.com/documentation/foundation/nsstring/1412645-precomposedstringwithcanonicalma?language=objc
func (x gen_NSString) PrecomposedStringWithCanonicalMapping() NSString {
	ret := C.NSString_inst_precomposedStringWithCanonicalMapping(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// PrecomposedStringWithCompatibilityMapping A string made by normalizing the receiver’s contents using the Unicode Normalization Form KC.
// https://developer.apple.com/documentation/foundation/nsstring/1412625-precomposedstringwithcompatibili?language=objc
func (x gen_NSString) PrecomposedStringWithCompatibilityMapping() NSString {
	ret := C.NSString_inst_precomposedStringWithCompatibilityMapping(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// IntValue The integer value of the string.
// https://developer.apple.com/documentation/foundation/nsstring/1414988-intvalue?language=objc
func (x gen_NSString) IntValue() int32 {
	ret := C.NSString_inst_intValue(
		unsafe.Pointer(x.Pointer()),
	)

	return int32(ret)

}

// IntegerValue The NSInteger value of the string.
// https://developer.apple.com/documentation/foundation/nsstring/1410267-integervalue?language=objc
func (x gen_NSString) IntegerValue() NSInteger {
	ret := C.NSString_inst_integerValue(
		unsafe.Pointer(x.Pointer()),
	)

	return NSInteger(ret)

}

// BoolValue The Boolean value of the string.
// https://developer.apple.com/documentation/foundation/nsstring/1409420-boolvalue?language=objc
func (x gen_NSString) BoolValue() bool {
	ret := C.NSString_inst_boolValue(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Description This NSString object.
// https://developer.apple.com/documentation/foundation/nsstring/1410889-description?language=objc
func (x gen_NSString) Description() NSString {
	ret := C.NSString_inst_description(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// FastestEncoding The fastest encoding to which the receiver may be converted without loss of information.
// https://developer.apple.com/documentation/foundation/nsstring/1409567-fastestencoding?language=objc
func (x gen_NSString) FastestEncoding() NSStringEncoding {
	ret := C.NSString_inst_fastestEncoding(
		unsafe.Pointer(x.Pointer()),
	)

	return NSStringEncoding(ret)

}

// SmallestEncoding The smallest encoding to which the receiver can be converted without loss of information.
// https://developer.apple.com/documentation/foundation/nsstring/1418037-smallestencoding?language=objc
func (x gen_NSString) SmallestEncoding() NSStringEncoding {
	ret := C.NSString_inst_smallestEncoding(
		unsafe.Pointer(x.Pointer()),
	)

	return NSStringEncoding(ret)

}

// PathComponents The file-system path components of the receiver.
// https://developer.apple.com/documentation/foundation/nsstring/1414489-pathcomponents?language=objc
func (x gen_NSString) PathComponents() NSArray {
	ret := C.NSString_inst_pathComponents(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_fromPointer(ret)

}

// IsAbsolutePath A Boolean value that indicates whether the receiver represents an absolute path.
// https://developer.apple.com/documentation/foundation/nsstring/1409068-absolutepath?language=objc
func (x gen_NSString) IsAbsolutePath() bool {
	ret := C.NSString_inst_isAbsolutePath(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// LastPathComponent The last path component of the receiver.
// https://developer.apple.com/documentation/foundation/nsstring/1416528-lastpathcomponent?language=objc
func (x gen_NSString) LastPathComponent() NSString {
	ret := C.NSString_inst_lastPathComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// PathExtension The path extension, if any, of the string as interpreted as a path.
// https://developer.apple.com/documentation/foundation/nsstring/1407801-pathextension?language=objc
func (x gen_NSString) PathExtension() NSString {
	ret := C.NSString_inst_pathExtension(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// StringByAbbreviatingWithTildeInPath A new string that replaces the current home directory portion of the current path with a tilde (~) character.
// https://developer.apple.com/documentation/foundation/nsstring/1407943-stringbyabbreviatingwithtildeinp?language=objc
func (x gen_NSString) StringByAbbreviatingWithTildeInPath() NSString {
	ret := C.NSString_inst_stringByAbbreviatingWithTildeInPath(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// StringByDeletingLastPathComponent A new string made by deleting the last path component from the receiver, along with any final path separator.
// https://developer.apple.com/documentation/foundation/nsstring/1411141-stringbydeletinglastpathcomponen?language=objc
func (x gen_NSString) StringByDeletingLastPathComponent() NSString {
	ret := C.NSString_inst_stringByDeletingLastPathComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// StringByDeletingPathExtension A new string made by deleting the extension (if any, and only the last) from the receiver.
// https://developer.apple.com/documentation/foundation/nsstring/1418214-stringbydeletingpathextension?language=objc
func (x gen_NSString) StringByDeletingPathExtension() NSString {
	ret := C.NSString_inst_stringByDeletingPathExtension(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// StringByExpandingTildeInPath A new string made by expanding the initial component of the receiver to its full path value.
// https://developer.apple.com/documentation/foundation/nsstring/1407716-stringbyexpandingtildeinpath?language=objc
func (x gen_NSString) StringByExpandingTildeInPath() NSString {
	ret := C.NSString_inst_stringByExpandingTildeInPath(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// StringByResolvingSymlinksInPath A new string made from the receiver by resolving all symbolic links and standardizing path.
// https://developer.apple.com/documentation/foundation/nsstring/1417783-stringbyresolvingsymlinksinpath?language=objc
func (x gen_NSString) StringByResolvingSymlinksInPath() NSString {
	ret := C.NSString_inst_stringByResolvingSymlinksInPath(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// StringByStandardizingPath A new string made by removing extraneous path components from the receiver.
// https://developer.apple.com/documentation/foundation/nsstring/1407194-stringbystandardizingpath?language=objc
func (x gen_NSString) StringByStandardizingPath() NSString {
	ret := C.NSString_inst_stringByStandardizingPath(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// StringByRemovingPercentEncoding Returns a new string made from the receiver by replacing all percent encoded sequences with the matching UTF-8 characters.
// https://developer.apple.com/documentation/foundation/nsstring/1409569-stringbyremovingpercentencoding?language=objc
func (x gen_NSString) StringByRemovingPercentEncoding() NSString {
	ret := C.NSString_inst_stringByRemovingPercentEncoding(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

type NSThreadRef interface {
	Pointer() uintptr
	Init_asNSThread() NSThread
}

type gen_NSThread struct {
	objc.Object
}

func NSThread_fromPointer(ptr unsafe.Pointer) NSThread {
	return NSThread{gen_NSThread{
		objc.Object_fromPointer(ptr),
	}}
}

func NSThread_fromRef(ref objc.Ref) NSThread {
	return NSThread_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// Cancel Changes the cancelled state of the receiver to indicate that it should exit.
// https://developer.apple.com/documentation/foundation/nsthread/1411303-cancel?language=objc
func (x gen_NSThread) Cancel() {
	C.NSThread_inst_cancel(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Init_asNSThread Returns an initialized NSThread object.
// https://developer.apple.com/documentation/foundation/nsthread/1416464-init?language=objc
func (x gen_NSThread) Init_asNSThread() NSThread {
	ret := C.NSThread_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSThread_fromPointer(ret)

}

// InitWithTarget_selector_object_asNSThread Returns an NSThread object initialized with the given arguments.
// https://developer.apple.com/documentation/foundation/nsthread/1414773-initwithtarget?language=objc
func (x gen_NSThread) InitWithTarget_selector_object_asNSThread(
	target objc.Ref,
	selector objc.Selector,
	argument objc.Ref,
) NSThread {
	ret := C.NSThread_inst_initWithTarget_selector_object(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(target),
		selector.SelectorAddress(),
		objc.RefPointer(argument),
	)

	return NSThread_fromPointer(ret)

}

// Main The main entry point routine for the thread.
// https://developer.apple.com/documentation/foundation/nsthread/1418421-main?language=objc
func (x gen_NSThread) Main() {
	C.NSThread_inst_main(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Start Starts the receiver.
// https://developer.apple.com/documentation/foundation/nsthread/1418166-start?language=objc
func (x gen_NSThread) Start() {
	C.NSThread_inst_start(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// IsExecuting A Boolean value that indicates whether the receiver is executing.
// https://developer.apple.com/documentation/foundation/nsthread/1411240-executing?language=objc
func (x gen_NSThread) IsExecuting() bool {
	ret := C.NSThread_inst_isExecuting(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsFinished A Boolean value that indicates whether the receiver has finished execution.
// https://developer.apple.com/documentation/foundation/nsthread/1409297-finished?language=objc
func (x gen_NSThread) IsFinished() bool {
	ret := C.NSThread_inst_isFinished(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsCancelled A Boolean value that indicates whether the receiver is cancelled.
// https://developer.apple.com/documentation/foundation/nsthread/1417366-cancelled?language=objc
func (x gen_NSThread) IsCancelled() bool {
	ret := C.NSThread_inst_isCancelled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// IsMainThread A Boolean value that indicates whether the receiver is the main thread.
// https://developer.apple.com/documentation/foundation/nsthread/1408455-ismainthread?language=objc
func (x gen_NSThread) IsMainThread() bool {
	ret := C.NSThread_inst_isMainThread(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Name The name of the receiver.
// https://developer.apple.com/documentation/foundation/nsthread/1414122-name?language=objc
func (x gen_NSThread) Name() NSString {
	ret := C.NSThread_inst_name(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// SetName The name of the receiver.
// https://developer.apple.com/documentation/foundation/nsthread/1414122-name?language=objc
func (x gen_NSThread) SetName(
	value NSStringRef,
) {
	C.NSThread_inst_setName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// StackSize The stack size of the receiver, in bytes.
// https://developer.apple.com/documentation/foundation/nsthread/1415190-stacksize?language=objc
func (x gen_NSThread) StackSize() NSUInteger {
	ret := C.NSThread_inst_stackSize(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUInteger(ret)

}

// SetStackSize The stack size of the receiver, in bytes.
// https://developer.apple.com/documentation/foundation/nsthread/1415190-stacksize?language=objc
func (x gen_NSThread) SetStackSize(
	value NSUInteger,
) {
	C.NSThread_inst_setStackSize(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)

	return

}

type NSURLRef interface {
	Pointer() uintptr
	Init_asNSURL() NSURL
}

type gen_NSURL struct {
	objc.Object
}

func NSURL_fromPointer(ptr unsafe.Pointer) NSURL {
	return NSURL{gen_NSURL{
		objc.Object_fromPointer(ptr),
	}}
}

func NSURL_fromRef(ref objc.Ref) NSURL {
	return NSURL_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// URLByAppendingPathComponent Returns a new URL made by appending a path component to the original URL.
// https://developer.apple.com/documentation/foundation/nsurl/1410614-urlbyappendingpathcomponent?language=objc
func (x gen_NSURL) URLByAppendingPathComponent(
	pathComponent NSStringRef,
) NSURL {
	ret := C.NSURL_inst_URLByAppendingPathComponent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pathComponent),
	)

	return NSURL_fromPointer(ret)

}

// URLByAppendingPathComponent_isDirectory Returns a new URL made by appending a path component to the original URL, along with a trailing slash if the component is designated a directory.
// https://developer.apple.com/documentation/foundation/nsurl/1413953-urlbyappendingpathcomponent?language=objc
func (x gen_NSURL) URLByAppendingPathComponent_isDirectory(
	pathComponent NSStringRef,
	isDirectory bool,
) NSURL {
	ret := C.NSURL_inst_URLByAppendingPathComponent_isDirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pathComponent),
		convertToObjCBool(isDirectory),
	)

	return NSURL_fromPointer(ret)

}

// URLByAppendingPathExtension Returns a new URL made by appending a path extension to the original URL.
// https://developer.apple.com/documentation/foundation/nsurl/1417082-urlbyappendingpathextension?language=objc
func (x gen_NSURL) URLByAppendingPathExtension(
	pathExtension NSStringRef,
) NSURL {
	ret := C.NSURL_inst_URLByAppendingPathExtension(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pathExtension),
	)

	return NSURL_fromPointer(ret)

}

// FileReferenceURL Returns a new file reference URL that points to the same resource as the receiver.
// https://developer.apple.com/documentation/foundation/nsurl/1408631-filereferenceurl?language=objc
func (x gen_NSURL) FileReferenceURL() NSURL {
	ret := C.NSURL_inst_fileReferenceURL(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_fromPointer(ret)

}

// InitAbsoluteURLWithDataRepresentation_relativeToURL_asNSURL
// https://developer.apple.com/documentation/foundation/nsurl/1410750-initabsoluteurlwithdatarepresent?language=objc
func (x gen_NSURL) InitAbsoluteURLWithDataRepresentation_relativeToURL_asNSURL(
	data NSDataRef,
	baseURL NSURLRef,
) NSURL {
	ret := C.NSURL_inst_initAbsoluteURLWithDataRepresentation_relativeToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(baseURL),
	)

	return NSURL_fromPointer(ret)

}

// InitFileURLWithPath_asNSURL Initializes a newly created NSURL referencing the local file or directory at path.
// https://developer.apple.com/documentation/foundation/nsurl/1410301-initfileurlwithpath?language=objc
func (x gen_NSURL) InitFileURLWithPath_asNSURL(
	path NSStringRef,
) NSURL {
	ret := C.NSURL_inst_initFileURLWithPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)

	return NSURL_fromPointer(ret)

}

// InitFileURLWithPath_isDirectory_asNSURL Initializes a newly created NSURL referencing the local file or directory at path.
// https://developer.apple.com/documentation/foundation/nsurl/1417505-initfileurlwithpath?language=objc
func (x gen_NSURL) InitFileURLWithPath_isDirectory_asNSURL(
	path NSStringRef,
	isDir bool,
) NSURL {
	ret := C.NSURL_inst_initFileURLWithPath_isDirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(isDir),
	)

	return NSURL_fromPointer(ret)

}

// InitFileURLWithPath_isDirectory_relativeToURL_asNSURL
// https://developer.apple.com/documentation/foundation/nsurl/1417932-initfileurlwithpath?language=objc
func (x gen_NSURL) InitFileURLWithPath_isDirectory_relativeToURL_asNSURL(
	path NSStringRef,
	isDir bool,
	baseURL NSURLRef,
) NSURL {
	ret := C.NSURL_inst_initFileURLWithPath_isDirectory_relativeToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(isDir),
		objc.RefPointer(baseURL),
	)

	return NSURL_fromPointer(ret)

}

// InitFileURLWithPath_relativeToURL_asNSURL
// https://developer.apple.com/documentation/foundation/nsurl/1415077-initfileurlwithpath?language=objc
func (x gen_NSURL) InitFileURLWithPath_relativeToURL_asNSURL(
	path NSStringRef,
	baseURL NSURLRef,
) NSURL {
	ret := C.NSURL_inst_initFileURLWithPath_relativeToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		objc.RefPointer(baseURL),
	)

	return NSURL_fromPointer(ret)

}

// InitWithDataRepresentation_relativeToURL_asNSURL
// https://developer.apple.com/documentation/foundation/nsurl/1416851-initwithdatarepresentation?language=objc
func (x gen_NSURL) InitWithDataRepresentation_relativeToURL_asNSURL(
	data NSDataRef,
	baseURL NSURLRef,
) NSURL {
	ret := C.NSURL_inst_initWithDataRepresentation_relativeToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(baseURL),
	)

	return NSURL_fromPointer(ret)

}

// InitWithString_asNSURL Initializes an NSURL object with a provided URL string.
// https://developer.apple.com/documentation/foundation/nsurl/1413146-initwithstring?language=objc
func (x gen_NSURL) InitWithString_asNSURL(
	URLString NSStringRef,
) NSURL {
	ret := C.NSURL_inst_initWithString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(URLString),
	)

	return NSURL_fromPointer(ret)

}

// InitWithString_relativeToURL_asNSURL Initializes an NSURL object with a base URL and a relative string.
// https://developer.apple.com/documentation/foundation/nsurl/1417949-initwithstring?language=objc
func (x gen_NSURL) InitWithString_relativeToURL_asNSURL(
	URLString NSStringRef,
	baseURL NSURLRef,
) NSURL {
	ret := C.NSURL_inst_initWithString_relativeToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(URLString),
		objc.RefPointer(baseURL),
	)

	return NSURL_fromPointer(ret)

}

// IsFileReferenceURL Returns whether the URL is a file reference URL.
// https://developer.apple.com/documentation/foundation/nsurl/1408507-isfilereferenceurl?language=objc
func (x gen_NSURL) IsFileReferenceURL() bool {
	ret := C.NSURL_inst_isFileReferenceURL(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// RemoveAllCachedResourceValues Removes all cached resource values and temporary resource values from the URL object.
// https://developer.apple.com/documentation/foundation/nsurl/1417078-removeallcachedresourcevalues?language=objc
func (x gen_NSURL) RemoveAllCachedResourceValues() {
	C.NSURL_inst_removeAllCachedResourceValues(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// StartAccessingSecurityScopedResource In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
// https://developer.apple.com/documentation/foundation/nsurl/1417051-startaccessingsecurityscopedreso?language=objc
func (x gen_NSURL) StartAccessingSecurityScopedResource() bool {
	ret := C.NSURL_inst_startAccessingSecurityScopedResource(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// StopAccessingSecurityScopedResource In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL.
// https://developer.apple.com/documentation/foundation/nsurl/1413736-stopaccessingsecurityscopedresou?language=objc
func (x gen_NSURL) StopAccessingSecurityScopedResource() {
	C.NSURL_inst_stopAccessingSecurityScopedResource(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// Init_asNSURL
func (x gen_NSURL) Init_asNSURL() NSURL {
	ret := C.NSURL_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_fromPointer(ret)

}

// DataRepresentation
// https://developer.apple.com/documentation/foundation/nsurl/1407656-datarepresentation?language=objc
func (x gen_NSURL) DataRepresentation() NSData {
	ret := C.NSURL_inst_dataRepresentation(
		unsafe.Pointer(x.Pointer()),
	)

	return NSData_fromPointer(ret)

}

// IsFileURL A boolean value that determines whether the receiver is a file URL.
// https://developer.apple.com/documentation/foundation/nsurl/1408782-fileurl?language=objc
func (x gen_NSURL) IsFileURL() bool {
	ret := C.NSURL_inst_isFileURL(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AbsoluteString The URL string for the receiver as an absolute URL. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1409868-absolutestring?language=objc
func (x gen_NSURL) AbsoluteString() NSString {
	ret := C.NSURL_inst_absoluteString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// AbsoluteURL An absolute URL that refers to the same resource as the receiver. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1414266-absoluteurl?language=objc
func (x gen_NSURL) AbsoluteURL() NSURL {
	ret := C.NSURL_inst_absoluteURL(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_fromPointer(ret)

}

// BaseURL The base URL. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1412311-baseurl?language=objc
func (x gen_NSURL) BaseURL() NSURL {
	ret := C.NSURL_inst_baseURL(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_fromPointer(ret)

}

// Fragment The fragment identifier, conforming to RFC 1808. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1413775-fragment?language=objc
func (x gen_NSURL) Fragment() NSString {
	ret := C.NSURL_inst_fragment(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// Host The host, conforming to RFC 1808. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1413640-host?language=objc
func (x gen_NSURL) Host() NSString {
	ret := C.NSURL_inst_host(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// LastPathComponent The last path component. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1417444-lastpathcomponent?language=objc
func (x gen_NSURL) LastPathComponent() NSString {
	ret := C.NSURL_inst_lastPathComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// Password The password conforming to RFC 1808. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1412096-password?language=objc
func (x gen_NSURL) Password() NSString {
	ret := C.NSURL_inst_password(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// Path The path, conforming to RFC 1808. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1408809-path?language=objc
func (x gen_NSURL) Path() NSString {
	ret := C.NSURL_inst_path(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// PathComponents An array containing the path components. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1407365-pathcomponents?language=objc
func (x gen_NSURL) PathComponents() NSArray {
	ret := C.NSURL_inst_pathComponents(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_fromPointer(ret)

}

// PathExtension The path extension. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1410208-pathextension?language=objc
func (x gen_NSURL) PathExtension() NSString {
	ret := C.NSURL_inst_pathExtension(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// Port The port, conforming to RFC 1808.
// https://developer.apple.com/documentation/foundation/nsurl/1413455-port?language=objc
func (x gen_NSURL) Port() NSNumber {
	ret := C.NSURL_inst_port(
		unsafe.Pointer(x.Pointer()),
	)

	return NSNumber_fromPointer(ret)

}

// Query The query string, conforming to RFC 1808.
// https://developer.apple.com/documentation/foundation/nsurl/1407543-query?language=objc
func (x gen_NSURL) Query() NSString {
	ret := C.NSURL_inst_query(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// RelativePath The relative path, conforming to RFC 1808. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1410263-relativepath?language=objc
func (x gen_NSURL) RelativePath() NSString {
	ret := C.NSURL_inst_relativePath(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// RelativeString A string representation of the relative portion of the URL. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1411417-relativestring?language=objc
func (x gen_NSURL) RelativeString() NSString {
	ret := C.NSURL_inst_relativeString(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// ResourceSpecifier The resource specifier. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1415309-resourcespecifier?language=objc
func (x gen_NSURL) ResourceSpecifier() NSString {
	ret := C.NSURL_inst_resourceSpecifier(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// Scheme The scheme. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1413437-scheme?language=objc
func (x gen_NSURL) Scheme() NSString {
	ret := C.NSURL_inst_scheme(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// StandardizedURL A copy of the URL with any instances of ".." or "." removed from its path. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1411073-standardizedurl?language=objc
func (x gen_NSURL) StandardizedURL() NSURL {
	ret := C.NSURL_inst_standardizedURL(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_fromPointer(ret)

}

// User The user name, conforming to RFC 1808.
// https://developer.apple.com/documentation/foundation/nsurl/1418335-user?language=objc
func (x gen_NSURL) User() NSString {
	ret := C.NSURL_inst_user(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// FilePathURL A file path URL that points to the same resource as the URL object. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1408442-filepathurl?language=objc
func (x gen_NSURL) FilePathURL() NSURL {
	ret := C.NSURL_inst_filePathURL(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_fromPointer(ret)

}

// URLByDeletingLastPathComponent A URL created by taking the receiver and removing the last path component. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1411592-urlbydeletinglastpathcomponent?language=objc
func (x gen_NSURL) URLByDeletingLastPathComponent() NSURL {
	ret := C.NSURL_inst_URLByDeletingLastPathComponent(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_fromPointer(ret)

}

// URLByDeletingPathExtension A URL created by taking the receiver and removing the path extension, if any. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1412357-urlbydeletingpathextension?language=objc
func (x gen_NSURL) URLByDeletingPathExtension() NSURL {
	ret := C.NSURL_inst_URLByDeletingPathExtension(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_fromPointer(ret)

}

// URLByResolvingSymlinksInPath A URL that points to the same resource as the receiver and includes no symbolic links. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1415965-urlbyresolvingsymlinksinpath?language=objc
func (x gen_NSURL) URLByResolvingSymlinksInPath() NSURL {
	ret := C.NSURL_inst_URLByResolvingSymlinksInPath(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_fromPointer(ret)

}

// URLByStandardizingPath A URL that points to the same resource as the original URL using an absolute path. (read-only)
// https://developer.apple.com/documentation/foundation/nsurl/1414302-urlbystandardizingpath?language=objc
func (x gen_NSURL) URLByStandardizingPath() NSURL {
	ret := C.NSURL_inst_URLByStandardizingPath(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_fromPointer(ret)

}

// HasDirectoryPath
// https://developer.apple.com/documentation/foundation/nsurl/1411475-hasdirectorypath?language=objc
func (x gen_NSURL) HasDirectoryPath() bool {
	ret := C.NSURL_inst_hasDirectoryPath(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

type NSURLRequestRef interface {
	Pointer() uintptr
	Init_asNSURLRequest() NSURLRequest
}

type gen_NSURLRequest struct {
	objc.Object
}

func NSURLRequest_fromPointer(ptr unsafe.Pointer) NSURLRequest {
	return NSURLRequest{gen_NSURLRequest{
		objc.Object_fromPointer(ptr),
	}}
}

func NSURLRequest_fromRef(ref objc.Ref) NSURLRequest {
	return NSURLRequest_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// InitWithURL_asNSURLRequest Creates a URL request for a specified URL.
// https://developer.apple.com/documentation/foundation/nsurlrequest/1410303-initwithurl?language=objc
func (x gen_NSURLRequest) InitWithURL_asNSURLRequest(
	URL NSURLRef,
) NSURLRequest {
	ret := C.NSURLRequest_inst_initWithURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(URL),
	)

	return NSURLRequest_fromPointer(ret)

}

// ValueForHTTPHeaderField Returns the value of the specified HTTP header field.
// https://developer.apple.com/documentation/foundation/nsurlrequest/1409376-valueforhttpheaderfield?language=objc
func (x gen_NSURLRequest) ValueForHTTPHeaderField(
	field NSStringRef,
) NSString {
	ret := C.NSURLRequest_inst_valueForHTTPHeaderField(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(field),
	)

	return NSString_fromPointer(ret)

}

// Init_asNSURLRequest
func (x gen_NSURLRequest) Init_asNSURLRequest() NSURLRequest {
	ret := C.NSURLRequest_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURLRequest_fromPointer(ret)

}

// HTTPMethod The HTTP request method.
// https://developer.apple.com/documentation/foundation/nsurlrequest/1413030-httpmethod?language=objc
func (x gen_NSURLRequest) HTTPMethod() NSString {
	ret := C.NSURLRequest_inst_HTTPMethod(
		unsafe.Pointer(x.Pointer()),
	)

	return NSString_fromPointer(ret)

}

// URL The URL being requested.
// https://developer.apple.com/documentation/foundation/nsurlrequest/1408996-url?language=objc
func (x gen_NSURLRequest) URL() NSURL {
	ret := C.NSURLRequest_inst_URL(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_fromPointer(ret)

}

// HTTPBody The request body.
// https://developer.apple.com/documentation/foundation/nsurlrequest/1411317-httpbody?language=objc
func (x gen_NSURLRequest) HTTPBody() NSData {
	ret := C.NSURLRequest_inst_HTTPBody(
		unsafe.Pointer(x.Pointer()),
	)

	return NSData_fromPointer(ret)

}

// MainDocumentURL The main document URL associated with the request.
// https://developer.apple.com/documentation/foundation/nsurlrequest/1414134-maindocumenturl?language=objc
func (x gen_NSURLRequest) MainDocumentURL() NSURL {
	ret := C.NSURLRequest_inst_mainDocumentURL(
		unsafe.Pointer(x.Pointer()),
	)

	return NSURL_fromPointer(ret)

}

// AllHTTPHeaderFields A dictionary containing all of the HTTP header fields for a request.
// https://developer.apple.com/documentation/foundation/nsurlrequest/1418477-allhttpheaderfields?language=objc
func (x gen_NSURLRequest) AllHTTPHeaderFields() NSDictionary {
	ret := C.NSURLRequest_inst_allHTTPHeaderFields(
		unsafe.Pointer(x.Pointer()),
	)

	return NSDictionary_fromPointer(ret)

}

// HTTPShouldHandleCookies A Boolean value that indicates whether the default cookie handling will be used for this request.
// https://developer.apple.com/documentation/foundation/nsurlrequest/1418369-httpshouldhandlecookies?language=objc
func (x gen_NSURLRequest) HTTPShouldHandleCookies() bool {
	ret := C.NSURLRequest_inst_HTTPShouldHandleCookies(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// HTTPShouldUsePipelining A Boolean value that indicates whether the request should continue transmitting data before receiving a response from an earlier transmission.
// https://developer.apple.com/documentation/foundation/nsurlrequest/1409170-httpshouldusepipelining?language=objc
func (x gen_NSURLRequest) HTTPShouldUsePipelining() bool {
	ret := C.NSURLRequest_inst_HTTPShouldUsePipelining(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AllowsCellularAccess A Boolean value that indicates whether the request is allowed to use the cellular radio (if present).
// https://developer.apple.com/documentation/foundation/nsurlrequest/1412032-allowscellularaccess?language=objc
func (x gen_NSURLRequest) AllowsCellularAccess() bool {
	ret := C.NSURLRequest_inst_allowsCellularAccess(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AllowsConstrainedNetworkAccess A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
// https://developer.apple.com/documentation/foundation/nsurlrequest/3325678-allowsconstrainednetworkaccess?language=objc
func (x gen_NSURLRequest) AllowsConstrainedNetworkAccess() bool {
	ret := C.NSURLRequest_inst_allowsConstrainedNetworkAccess(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AllowsExpensiveNetworkAccess A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
// https://developer.apple.com/documentation/foundation/nsurlrequest/3325679-allowsexpensivenetworkaccess?language=objc
func (x gen_NSURLRequest) AllowsExpensiveNetworkAccess() bool {
	ret := C.NSURLRequest_inst_allowsExpensiveNetworkAccess(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AssumesHTTP3Capable
// https://developer.apple.com/documentation/foundation/nsurlrequest/3735880-assumeshttp3capable?language=objc
func (x gen_NSURLRequest) AssumesHTTP3Capable() bool {
	ret := C.NSURLRequest_inst_assumesHTTP3Capable(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

type NSUserDefaultsRef interface {
	Pointer() uintptr
	Init_asNSUserDefaults() NSUserDefaults
}

type gen_NSUserDefaults struct {
	objc.Object
}

func NSUserDefaults_fromPointer(ptr unsafe.Pointer) NSUserDefaults {
	return NSUserDefaults{gen_NSUserDefaults{
		objc.Object_fromPointer(ptr),
	}}
}

func NSUserDefaults_fromRef(ref objc.Ref) NSUserDefaults {
	return NSUserDefaults_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// URLForKey Returns the URL associated with the specified key.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1408648-urlforkey?language=objc
func (x gen_NSUserDefaults) URLForKey(
	defaultName NSStringRef,
) NSURL {
	ret := C.NSUserDefaults_inst_URLForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return NSURL_fromPointer(ret)

}

// AddSuiteNamed Inserts the specified domain name into the receiver’s search list.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1410294-addsuitenamed?language=objc
func (x gen_NSUserDefaults) AddSuiteNamed(
	suiteName NSStringRef,
) {
	C.NSUserDefaults_inst_addSuiteNamed(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(suiteName),
	)

	return

}

// ArrayForKey Returns the array associated with the specified key.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1414792-arrayforkey?language=objc
func (x gen_NSUserDefaults) ArrayForKey(
	defaultName NSStringRef,
) NSArray {
	ret := C.NSUserDefaults_inst_arrayForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return NSArray_fromPointer(ret)

}

// BoolForKey Returns the Boolean value associated with the specified key.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1416388-boolforkey?language=objc
func (x gen_NSUserDefaults) BoolForKey(
	defaultName NSStringRef,
) bool {
	ret := C.NSUserDefaults_inst_boolForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return convertObjCBoolToGo(ret)

}

// DataForKey Returns the data object associated with the specified key.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1409590-dataforkey?language=objc
func (x gen_NSUserDefaults) DataForKey(
	defaultName NSStringRef,
) NSData {
	ret := C.NSUserDefaults_inst_dataForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return NSData_fromPointer(ret)

}

// DictionaryForKey Returns the dictionary object associated with the specified key.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1408563-dictionaryforkey?language=objc
func (x gen_NSUserDefaults) DictionaryForKey(
	defaultName NSStringRef,
) NSDictionary {
	ret := C.NSUserDefaults_inst_dictionaryForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return NSDictionary_fromPointer(ret)

}

// DictionaryRepresentation Returns a dictionary that contains a union of all key-value pairs in the domains in the search list.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1415919-dictionaryrepresentation?language=objc
func (x gen_NSUserDefaults) DictionaryRepresentation() NSDictionary {
	ret := C.NSUserDefaults_inst_dictionaryRepresentation(
		unsafe.Pointer(x.Pointer()),
	)

	return NSDictionary_fromPointer(ret)

}

// Init_asNSUserDefaults Creates a user defaults object initialized with the defaults for the app and current user.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1414356-init?language=objc
func (x gen_NSUserDefaults) Init_asNSUserDefaults() NSUserDefaults {
	ret := C.NSUserDefaults_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return NSUserDefaults_fromPointer(ret)

}

// InitWithSuiteName_asNSUserDefaults Creates a user defaults object initialized with the defaults for the specified database name.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1409957-initwithsuitename?language=objc
func (x gen_NSUserDefaults) InitWithSuiteName_asNSUserDefaults(
	suitename NSStringRef,
) NSUserDefaults {
	ret := C.NSUserDefaults_inst_initWithSuiteName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(suitename),
	)

	return NSUserDefaults_fromPointer(ret)

}

// IntegerForKey Returns the integer value associated with the specified key.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1407405-integerforkey?language=objc
func (x gen_NSUserDefaults) IntegerForKey(
	defaultName NSStringRef,
) NSInteger {
	ret := C.NSUserDefaults_inst_integerForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return NSInteger(ret)

}

// ObjectForKey Returns the object associated with the specified key.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1410095-objectforkey?language=objc
func (x gen_NSUserDefaults) ObjectForKey(
	defaultName NSStringRef,
) objc.Object {
	ret := C.NSUserDefaults_inst_objectForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return objc.Object_fromPointer(ret)

}

// ObjectIsForcedForKey Returns a Boolean value indicating whether the specified key is managed by an administrator.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1408635-objectisforcedforkey?language=objc
func (x gen_NSUserDefaults) ObjectIsForcedForKey(
	key NSStringRef,
) bool {
	ret := C.NSUserDefaults_inst_objectIsForcedForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)

	return convertObjCBoolToGo(ret)

}

// ObjectIsForcedForKey_inDomain Returns a Boolean value indicating whether the key in the specified domain is managed by an administrator.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1416306-objectisforcedforkey?language=objc
func (x gen_NSUserDefaults) ObjectIsForcedForKey_inDomain(
	key NSStringRef,
	domain NSStringRef,
) bool {
	ret := C.NSUserDefaults_inst_objectIsForcedForKey_inDomain(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
		objc.RefPointer(domain),
	)

	return convertObjCBoolToGo(ret)

}

// PersistentDomainForName Returns a dictionary representation of the defaults for the specified domain.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1412197-persistentdomainforname?language=objc
func (x gen_NSUserDefaults) PersistentDomainForName(
	domainName NSStringRef,
) NSDictionary {
	ret := C.NSUserDefaults_inst_persistentDomainForName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domainName),
	)

	return NSDictionary_fromPointer(ret)

}

// RegisterDefaults Adds the contents of the specified dictionary to the registration domain.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1417065-registerdefaults?language=objc
func (x gen_NSUserDefaults) RegisterDefaults(
	registrationDictionary NSDictionaryRef,
) {
	C.NSUserDefaults_inst_registerDefaults(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(registrationDictionary),
	)

	return

}

// RemoveObjectForKey Removes the value of the specified default key.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1411182-removeobjectforkey?language=objc
func (x gen_NSUserDefaults) RemoveObjectForKey(
	defaultName NSStringRef,
) {
	C.NSUserDefaults_inst_removeObjectForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return

}

// RemovePersistentDomainForName Removes the contents of the specified persistent domain from the user’s defaults.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1417339-removepersistentdomainforname?language=objc
func (x gen_NSUserDefaults) RemovePersistentDomainForName(
	domainName NSStringRef,
) {
	C.NSUserDefaults_inst_removePersistentDomainForName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domainName),
	)

	return

}

// RemoveSuiteNamed Removes the specified domain name from the receiver’s search list.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1408047-removesuitenamed?language=objc
func (x gen_NSUserDefaults) RemoveSuiteNamed(
	suiteName NSStringRef,
) {
	C.NSUserDefaults_inst_removeSuiteNamed(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(suiteName),
	)

	return

}

// RemoveVolatileDomainForName Removes the specified volatile domain from the user’s defaults.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1415955-removevolatiledomainforname?language=objc
func (x gen_NSUserDefaults) RemoveVolatileDomainForName(
	domainName NSStringRef,
) {
	C.NSUserDefaults_inst_removeVolatileDomainForName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domainName),
	)

	return

}

// SetBool_forKey Sets the value of the specified default key to the specified Boolean value.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1408905-setbool?language=objc
func (x gen_NSUserDefaults) SetBool_forKey(
	value bool,
	defaultName NSStringRef,
) {
	C.NSUserDefaults_inst_setBool_forKey(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
		objc.RefPointer(defaultName),
	)

	return

}

// SetInteger_forKey Sets the value of the specified default key to the specified integer value.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1413614-setinteger?language=objc
func (x gen_NSUserDefaults) SetInteger_forKey(
	value NSInteger,
	defaultName NSStringRef,
) {
	C.NSUserDefaults_inst_setInteger_forKey(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
		objc.RefPointer(defaultName),
	)

	return

}

// SetObject_forKey Sets the value of the specified default key.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1414067-setobject?language=objc
func (x gen_NSUserDefaults) SetObject_forKey(
	value objc.Ref,
	defaultName NSStringRef,
) {
	C.NSUserDefaults_inst_setObject_forKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
		objc.RefPointer(defaultName),
	)

	return

}

// SetPersistentDomain_forName Sets a dictionary for the specified persistent domain.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1408187-setpersistentdomain?language=objc
func (x gen_NSUserDefaults) SetPersistentDomain_forName(
	domain NSDictionaryRef,
	domainName NSStringRef,
) {
	C.NSUserDefaults_inst_setPersistentDomain_forName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domain),
		objc.RefPointer(domainName),
	)

	return

}

// SetURL_forKey Sets the value of the specified default key to the specified URL.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1414194-seturl?language=objc
func (x gen_NSUserDefaults) SetURL_forKey(
	url NSURLRef,
	defaultName NSStringRef,
) {
	C.NSUserDefaults_inst_setURL_forKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
		objc.RefPointer(defaultName),
	)

	return

}

// SetVolatileDomain_forName Sets the dictionary for the specified volatile domain.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1413720-setvolatiledomain?language=objc
func (x gen_NSUserDefaults) SetVolatileDomain_forName(
	domain NSDictionaryRef,
	domainName NSStringRef,
) {
	C.NSUserDefaults_inst_setVolatileDomain_forName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domain),
		objc.RefPointer(domainName),
	)

	return

}

// StringArrayForKey Returns the array of strings associated with the specified key.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1416414-stringarrayforkey?language=objc
func (x gen_NSUserDefaults) StringArrayForKey(
	defaultName NSStringRef,
) NSArray {
	ret := C.NSUserDefaults_inst_stringArrayForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return NSArray_fromPointer(ret)

}

// StringForKey Returns the string associated with the specified key.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1416700-stringforkey?language=objc
func (x gen_NSUserDefaults) StringForKey(
	defaultName NSStringRef,
) NSString {
	ret := C.NSUserDefaults_inst_stringForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)

	return NSString_fromPointer(ret)

}

// Synchronize Waits for any pending asynchronous updates to the defaults database and returns; this method is unnecessary and shouldn't be used.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1414005-synchronize?language=objc
func (x gen_NSUserDefaults) Synchronize() bool {
	ret := C.NSUserDefaults_inst_synchronize(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// VolatileDomainForName Returns the dictionary for the specified volatile domain.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1409592-volatiledomainforname?language=objc
func (x gen_NSUserDefaults) VolatileDomainForName(
	domainName NSStringRef,
) NSDictionary {
	ret := C.NSUserDefaults_inst_volatileDomainForName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domainName),
	)

	return NSDictionary_fromPointer(ret)

}

// VolatileDomainNames The current volatile domain names.
// https://developer.apple.com/documentation/foundation/nsuserdefaults/1414231-volatiledomainnames?language=objc
func (x gen_NSUserDefaults) VolatileDomainNames() NSArray {
	ret := C.NSUserDefaults_inst_volatileDomainNames(
		unsafe.Pointer(x.Pointer()),
	)

	return NSArray_fromPointer(ret)

}
