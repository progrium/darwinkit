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

func (x gen_CALayer) ActionForKey(
	event NSStringRef,
) (
	r0 objc.Object,
) {
	ret := C.CALayer_inst_actionForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(event),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_CALayer) AddSublayer(
	layer CALayerRef,
) {
	C.CALayer_inst_addSublayer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(layer),
	)
	return
}

func (x gen_CALayer) AnimationKeys() (
	r0 NSArray,
) {
	ret := C.CALayer_inst_animationKeys(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_CALayer) ContentsAreFlipped() (
	r0 bool,
) {
	ret := C.CALayer_inst_contentsAreFlipped(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CALayer) ConvertRect_fromLayer(
	r NSRect,
	l CALayerRef,
) (
	r0 NSRect,
) {
	ret := C.CALayer_inst_convertRect_fromLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&r)),
		objc.RefPointer(l),
	)
	r0 = *(*NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_CALayer) ConvertRect_toLayer(
	r NSRect,
	l CALayerRef,
) (
	r0 NSRect,
) {
	ret := C.CALayer_inst_convertRect_toLayer(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&r)),
		objc.RefPointer(l),
	)
	r0 = *(*NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_CALayer) Display() {
	C.CALayer_inst_display(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_CALayer) DisplayIfNeeded() {
	C.CALayer_inst_displayIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_CALayer) Init_asCALayer() (
	r0 CALayer,
) {
	ret := C.CALayer_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CALayer_fromPointer(ret)
	return
}

func (x gen_CALayer) InitWithLayer_asCALayer(
	layer objc.Ref,
) (
	r0 CALayer,
) {
	ret := C.CALayer_inst_initWithLayer(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(layer),
	)
	r0 = CALayer_fromPointer(ret)
	return
}

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

func (x gen_CALayer) LayoutIfNeeded() {
	C.CALayer_inst_layoutIfNeeded(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_CALayer) LayoutSublayers() {
	C.CALayer_inst_layoutSublayers(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_CALayer) ModelLayer_asCALayer() (
	r0 CALayer,
) {
	ret := C.CALayer_inst_modelLayer(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CALayer_fromPointer(ret)
	return
}

func (x gen_CALayer) NeedsDisplay() (
	r0 bool,
) {
	ret := C.CALayer_inst_needsDisplay(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CALayer) NeedsLayout() (
	r0 bool,
) {
	ret := C.CALayer_inst_needsLayout(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CALayer) PreferredFrameSize() (
	r0 NSSize,
) {
	ret := C.CALayer_inst_preferredFrameSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_CALayer) PresentationLayer_asCALayer() (
	r0 CALayer,
) {
	ret := C.CALayer_inst_presentationLayer(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CALayer_fromPointer(ret)
	return
}

func (x gen_CALayer) RemoveAllAnimations() {
	C.CALayer_inst_removeAllAnimations(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_CALayer) RemoveAnimationForKey(
	key NSStringRef,
) {
	C.CALayer_inst_removeAnimationForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)
	return
}

func (x gen_CALayer) RemoveFromSuperlayer() {
	C.CALayer_inst_removeFromSuperlayer(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

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

func (x gen_CALayer) ResizeSublayersWithOldSize(
	size NSSize,
) {
	C.CALayer_inst_resizeSublayersWithOldSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	return
}

func (x gen_CALayer) ResizeWithOldSuperlayerSize(
	size NSSize,
) {
	C.CALayer_inst_resizeWithOldSuperlayerSize(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&size)),
	)
	return
}

func (x gen_CALayer) ScrollRectToVisible(
	r NSRect,
) {
	C.CALayer_inst_scrollRectToVisible(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&r)),
	)
	return
}

func (x gen_CALayer) SetNeedsDisplay() {
	C.CALayer_inst_setNeedsDisplay(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_CALayer) SetNeedsDisplayInRect(
	r NSRect,
) {
	C.CALayer_inst_setNeedsDisplayInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&r)),
	)
	return
}

func (x gen_CALayer) SetNeedsLayout() {
	C.CALayer_inst_setNeedsLayout(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_CALayer) ShouldArchiveValueForKey(
	key NSStringRef,
) (
	r0 bool,
) {
	ret := C.CALayer_inst_shouldArchiveValueForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CALayer) Delegate() (
	r0 objc.Object,
) {
	ret := C.CALayer_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_CALayer) SetDelegate(
	value objc.Ref,
) {
	C.CALayer_inst_setDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_CALayer) Contents() (
	r0 objc.Object,
) {
	ret := C.CALayer_inst_contents(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_CALayer) SetContents(
	value objc.Ref,
) {
	C.CALayer_inst_setContents(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_CALayer) ContentsRect() (
	r0 NSRect,
) {
	ret := C.CALayer_inst_contentsRect(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_CALayer) SetContentsRect(
	value NSRect,
) {
	C.CALayer_inst_setContentsRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_CALayer) ContentsCenter() (
	r0 NSRect,
) {
	ret := C.CALayer_inst_contentsCenter(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_CALayer) SetContentsCenter(
	value NSRect,
) {
	C.CALayer_inst_setContentsCenter(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_CALayer) IsHidden() (
	r0 bool,
) {
	ret := C.CALayer_inst_isHidden(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CALayer) SetHidden(
	value bool,
) {
	C.CALayer_inst_setHidden(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_CALayer) MasksToBounds() (
	r0 bool,
) {
	ret := C.CALayer_inst_masksToBounds(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CALayer) SetMasksToBounds(
	value bool,
) {
	C.CALayer_inst_setMasksToBounds(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_CALayer) Mask() (
	r0 CALayer,
) {
	ret := C.CALayer_inst_mask(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CALayer_fromPointer(ret)
	return
}

func (x gen_CALayer) SetMask(
	value CALayerRef,
) {
	C.CALayer_inst_setMask(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_CALayer) IsDoubleSided() (
	r0 bool,
) {
	ret := C.CALayer_inst_isDoubleSided(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CALayer) SetDoubleSided(
	value bool,
) {
	C.CALayer_inst_setDoubleSided(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_CALayer) CornerRadius() (
	r0 CGFloat,
) {
	ret := C.CALayer_inst_cornerRadius(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CGFloat(ret)
	return
}

func (x gen_CALayer) SetCornerRadius(
	value CGFloat,
) {
	C.CALayer_inst_setCornerRadius(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_CALayer) BorderWidth() (
	r0 CGFloat,
) {
	ret := C.CALayer_inst_borderWidth(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CGFloat(ret)
	return
}

func (x gen_CALayer) SetBorderWidth(
	value CGFloat,
) {
	C.CALayer_inst_setBorderWidth(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_CALayer) ShadowRadius() (
	r0 CGFloat,
) {
	ret := C.CALayer_inst_shadowRadius(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CGFloat(ret)
	return
}

func (x gen_CALayer) SetShadowRadius(
	value CGFloat,
) {
	C.CALayer_inst_setShadowRadius(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_CALayer) ShadowOffset() (
	r0 NSSize,
) {
	ret := C.CALayer_inst_shadowOffset(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_CALayer) SetShadowOffset(
	value NSSize,
) {
	C.CALayer_inst_setShadowOffset(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSSize)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_CALayer) Style() (
	r0 NSDictionary,
) {
	ret := C.CALayer_inst_style(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSDictionary_fromPointer(ret)
	return
}

func (x gen_CALayer) SetStyle(
	value NSDictionaryRef,
) {
	C.CALayer_inst_setStyle(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_CALayer) AllowsEdgeAntialiasing() (
	r0 bool,
) {
	ret := C.CALayer_inst_allowsEdgeAntialiasing(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CALayer) SetAllowsEdgeAntialiasing(
	value bool,
) {
	C.CALayer_inst_setAllowsEdgeAntialiasing(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_CALayer) AllowsGroupOpacity() (
	r0 bool,
) {
	ret := C.CALayer_inst_allowsGroupOpacity(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CALayer) SetAllowsGroupOpacity(
	value bool,
) {
	C.CALayer_inst_setAllowsGroupOpacity(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_CALayer) Filters() (
	r0 NSArray,
) {
	ret := C.CALayer_inst_filters(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_CALayer) SetFilters(
	value NSArrayRef,
) {
	C.CALayer_inst_setFilters(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_CALayer) CompositingFilter() (
	r0 objc.Object,
) {
	ret := C.CALayer_inst_compositingFilter(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_CALayer) SetCompositingFilter(
	value objc.Ref,
) {
	C.CALayer_inst_setCompositingFilter(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_CALayer) BackgroundFilters() (
	r0 NSArray,
) {
	ret := C.CALayer_inst_backgroundFilters(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_CALayer) SetBackgroundFilters(
	value NSArrayRef,
) {
	C.CALayer_inst_setBackgroundFilters(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_CALayer) IsOpaque() (
	r0 bool,
) {
	ret := C.CALayer_inst_isOpaque(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CALayer) SetOpaque(
	value bool,
) {
	C.CALayer_inst_setOpaque(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_CALayer) IsGeometryFlipped() (
	r0 bool,
) {
	ret := C.CALayer_inst_isGeometryFlipped(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CALayer) SetGeometryFlipped(
	value bool,
) {
	C.CALayer_inst_setGeometryFlipped(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_CALayer) DrawsAsynchronously() (
	r0 bool,
) {
	ret := C.CALayer_inst_drawsAsynchronously(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CALayer) SetDrawsAsynchronously(
	value bool,
) {
	C.CALayer_inst_setDrawsAsynchronously(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_CALayer) ShouldRasterize() (
	r0 bool,
) {
	ret := C.CALayer_inst_shouldRasterize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CALayer) SetShouldRasterize(
	value bool,
) {
	C.CALayer_inst_setShouldRasterize(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_CALayer) RasterizationScale() (
	r0 CGFloat,
) {
	ret := C.CALayer_inst_rasterizationScale(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CGFloat(ret)
	return
}

func (x gen_CALayer) SetRasterizationScale(
	value CGFloat,
) {
	C.CALayer_inst_setRasterizationScale(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_CALayer) Frame() (
	r0 NSRect,
) {
	ret := C.CALayer_inst_frame(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_CALayer) SetFrame(
	value NSRect,
) {
	C.CALayer_inst_setFrame(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_CALayer) Bounds() (
	r0 NSRect,
) {
	ret := C.CALayer_inst_bounds(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_CALayer) SetBounds(
	value NSRect,
) {
	C.CALayer_inst_setBounds(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&value)),
	)
	return
}

func (x gen_CALayer) ZPosition() (
	r0 CGFloat,
) {
	ret := C.CALayer_inst_zPosition(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CGFloat(ret)
	return
}

func (x gen_CALayer) SetZPosition(
	value CGFloat,
) {
	C.CALayer_inst_setZPosition(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_CALayer) AnchorPointZ() (
	r0 CGFloat,
) {
	ret := C.CALayer_inst_anchorPointZ(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CGFloat(ret)
	return
}

func (x gen_CALayer) SetAnchorPointZ(
	value CGFloat,
) {
	C.CALayer_inst_setAnchorPointZ(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_CALayer) ContentsScale() (
	r0 CGFloat,
) {
	ret := C.CALayer_inst_contentsScale(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CGFloat(ret)
	return
}

func (x gen_CALayer) SetContentsScale(
	value CGFloat,
) {
	C.CALayer_inst_setContentsScale(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_CALayer) Sublayers() (
	r0 NSArray,
) {
	ret := C.CALayer_inst_sublayers(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_CALayer) SetSublayers(
	value NSArrayRef,
) {
	C.CALayer_inst_setSublayers(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_CALayer) Superlayer() (
	r0 CALayer,
) {
	ret := C.CALayer_inst_superlayer(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CALayer_fromPointer(ret)
	return
}

func (x gen_CALayer) NeedsDisplayOnBoundsChange() (
	r0 bool,
) {
	ret := C.CALayer_inst_needsDisplayOnBoundsChange(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CALayer) SetNeedsDisplayOnBoundsChange(
	value bool,
) {
	C.CALayer_inst_setNeedsDisplayOnBoundsChange(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_CALayer) LayoutManager() (
	r0 objc.Object,
) {
	ret := C.CALayer_inst_layoutManager(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_CALayer) SetLayoutManager(
	value objc.Ref,
) {
	C.CALayer_inst_setLayoutManager(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_CALayer) Constraints() (
	r0 NSArray,
) {
	ret := C.CALayer_inst_constraints(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_CALayer) SetConstraints(
	value NSArrayRef,
) {
	C.CALayer_inst_setConstraints(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_CALayer) Actions() (
	r0 NSDictionary,
) {
	ret := C.CALayer_inst_actions(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSDictionary_fromPointer(ret)
	return
}

func (x gen_CALayer) SetActions(
	value NSDictionaryRef,
) {
	C.CALayer_inst_setActions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_CALayer) VisibleRect() (
	r0 NSRect,
) {
	ret := C.CALayer_inst_visibleRect(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*NSRect)(unsafe.Pointer(&ret))
	return
}

func (x gen_CALayer) Name() (
	r0 NSString,
) {
	ret := C.CALayer_inst_name(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

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

func (x gen_NSArray) ArrayByAddingObjectsFromArray(
	otherArray NSArrayRef,
) (
	r0 NSArray,
) {
	ret := C.NSArray_inst_arrayByAddingObjectsFromArray(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(otherArray),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSArray) ComponentsJoinedByString(
	separator NSStringRef,
) (
	r0 NSString,
) {
	ret := C.NSArray_inst_componentsJoinedByString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(separator),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSArray) DescriptionWithLocale(
	locale objc.Ref,
) (
	r0 NSString,
) {
	ret := C.NSArray_inst_descriptionWithLocale(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(locale),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSArray) DescriptionWithLocale_indent(
	locale objc.Ref,
	level NSUInteger,
) (
	r0 NSString,
) {
	ret := C.NSArray_inst_descriptionWithLocale_indent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(locale),
		C.ulong(level),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSArray) Init_asNSArray() (
	r0 NSArray,
) {
	ret := C.NSArray_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSArray) InitWithArray_asNSArray(
	array NSArrayRef,
) (
	r0 NSArray,
) {
	ret := C.NSArray_inst_initWithArray(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(array),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSArray) InitWithArray_copyItems_asNSArray(
	array NSArrayRef,
	flag bool,
) (
	r0 NSArray,
) {
	ret := C.NSArray_inst_initWithArray_copyItems(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(array),
		convertToObjCBool(flag),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSArray) IsEqualToArray(
	otherArray NSArrayRef,
) (
	r0 bool,
) {
	ret := C.NSArray_inst_isEqualToArray(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(otherArray),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSArray) MakeObjectsPerformSelector(
	aSelector objc.Selector,
) {
	C.NSArray_inst_makeObjectsPerformSelector(
		unsafe.Pointer(x.Pointer()),
		aSelector.SelectorAddress(),
	)
	return
}

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

func (x gen_NSArray) PathsMatchingExtensions(
	filterTypes NSArrayRef,
) (
	r0 NSArray,
) {
	ret := C.NSArray_inst_pathsMatchingExtensions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(filterTypes),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

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

func (x gen_NSArray) ShuffledArray() (
	r0 NSArray,
) {
	ret := C.NSArray_inst_shuffledArray(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSArray) SortedArrayUsingDescriptors(
	sortDescriptors NSArrayRef,
) (
	r0 NSArray,
) {
	ret := C.NSArray_inst_sortedArrayUsingDescriptors(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sortDescriptors),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSArray) SortedArrayUsingSelector(
	comparator objc.Selector,
) (
	r0 NSArray,
) {
	ret := C.NSArray_inst_sortedArrayUsingSelector(
		unsafe.Pointer(x.Pointer()),
		comparator.SelectorAddress(),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSArray) ValueForKey(
	key NSStringRef,
) (
	r0 objc.Object,
) {
	ret := C.NSArray_inst_valueForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSArray) Count() (
	r0 NSUInteger,
) {
	ret := C.NSArray_inst_count(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSUInteger(ret)
	return
}

func (x gen_NSArray) SortedArrayHint() (
	r0 NSData,
) {
	ret := C.NSArray_inst_sortedArrayHint(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSData_fromPointer(ret)
	return
}

func (x gen_NSArray) Description() (
	r0 NSString,
) {
	ret := C.NSArray_inst_description(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
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

func (x gen_NSAttributedString) AttributedStringByInflectingString() (
	r0 NSAttributedString,
) {
	ret := C.NSAttributedString_inst_attributedStringByInflectingString(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSAttributedString) DrawInRect(
	rect NSRect,
) {
	C.NSAttributedString_inst_drawInRect(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&rect)),
	)
	return
}

func (x gen_NSAttributedString) InitWithAttributedString_asNSAttributedString(
	attrStr NSAttributedStringRef,
) (
	r0 NSAttributedString,
) {
	ret := C.NSAttributedString_inst_initWithAttributedString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(attrStr),
	)
	r0 = NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSAttributedString) InitWithDocFormat_documentAttributes_asNSAttributedString(
	data NSDataRef,
	dict NSDictionaryRef,
) (
	r0 NSAttributedString,
) {
	ret := C.NSAttributedString_inst_initWithDocFormat_documentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(dict),
	)
	r0 = NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSAttributedString) InitWithHTML_baseURL_documentAttributes_asNSAttributedString(
	data NSDataRef,
	base NSURLRef,
	dict NSDictionaryRef,
) (
	r0 NSAttributedString,
) {
	ret := C.NSAttributedString_inst_initWithHTML_baseURL_documentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(base),
		objc.RefPointer(dict),
	)
	r0 = NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSAttributedString) InitWithHTML_documentAttributes_asNSAttributedString(
	data NSDataRef,
	dict NSDictionaryRef,
) (
	r0 NSAttributedString,
) {
	ret := C.NSAttributedString_inst_initWithHTML_documentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(dict),
	)
	r0 = NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSAttributedString) InitWithHTML_options_documentAttributes_asNSAttributedString(
	data NSDataRef,
	options NSDictionaryRef,
	dict NSDictionaryRef,
) (
	r0 NSAttributedString,
) {
	ret := C.NSAttributedString_inst_initWithHTML_options_documentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(options),
		objc.RefPointer(dict),
	)
	r0 = NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSAttributedString) InitWithRTF_documentAttributes_asNSAttributedString(
	data NSDataRef,
	dict NSDictionaryRef,
) (
	r0 NSAttributedString,
) {
	ret := C.NSAttributedString_inst_initWithRTF_documentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(dict),
	)
	r0 = NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSAttributedString) InitWithRTFD_documentAttributes_asNSAttributedString(
	data NSDataRef,
	dict NSDictionaryRef,
) (
	r0 NSAttributedString,
) {
	ret := C.NSAttributedString_inst_initWithRTFD_documentAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(dict),
	)
	r0 = NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSAttributedString) InitWithString_asNSAttributedString(
	str NSStringRef,
) (
	r0 NSAttributedString,
) {
	ret := C.NSAttributedString_inst_initWithString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)
	r0 = NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSAttributedString) InitWithString_attributes_asNSAttributedString(
	str NSStringRef,
	attrs NSDictionaryRef,
) (
	r0 NSAttributedString,
) {
	ret := C.NSAttributedString_inst_initWithString_attributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
		objc.RefPointer(attrs),
	)
	r0 = NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSAttributedString) IsEqualToAttributedString(
	other NSAttributedStringRef,
) (
	r0 bool,
) {
	ret := C.NSAttributedString_inst_isEqualToAttributedString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(other),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSAttributedString) NextWordFromIndex_forward(
	location NSUInteger,
	isForward bool,
) (
	r0 NSUInteger,
) {
	ret := C.NSAttributedString_inst_nextWordFromIndex_forward(
		unsafe.Pointer(x.Pointer()),
		C.ulong(location),
		convertToObjCBool(isForward),
	)
	r0 = NSUInteger(ret)
	return
}

func (x gen_NSAttributedString) Size() (
	r0 NSSize,
) {
	ret := C.NSAttributedString_inst_size(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = *(*NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSAttributedString) Init_asNSAttributedString() (
	r0 NSAttributedString,
) {
	ret := C.NSAttributedString_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSAttributedString_fromPointer(ret)
	return
}

func (x gen_NSAttributedString) String() (
	r0 NSString,
) {
	ret := C.NSAttributedString_inst_string(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSAttributedString) Length() (
	r0 NSUInteger,
) {
	ret := C.NSAttributedString_inst_length(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSUInteger(ret)
	return
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

func (x gen_NSData) InitWithBytes_length_asNSData(
	bytes unsafe.Pointer,
	length NSUInteger,
) (
	r0 NSData,
) {
	ret := C.NSData_inst_initWithBytes_length(
		unsafe.Pointer(x.Pointer()),
		bytes,
		C.ulong(length),
	)
	r0 = NSData_fromPointer(ret)
	return
}

func (x gen_NSData) InitWithBytesNoCopy_length_asNSData(
	bytes unsafe.Pointer,
	length NSUInteger,
) (
	r0 NSData,
) {
	ret := C.NSData_inst_initWithBytesNoCopy_length(
		unsafe.Pointer(x.Pointer()),
		bytes,
		C.ulong(length),
	)
	r0 = NSData_fromPointer(ret)
	return
}

func (x gen_NSData) InitWithBytesNoCopy_length_freeWhenDone_asNSData(
	bytes unsafe.Pointer,
	length NSUInteger,
	b bool,
) (
	r0 NSData,
) {
	ret := C.NSData_inst_initWithBytesNoCopy_length_freeWhenDone(
		unsafe.Pointer(x.Pointer()),
		bytes,
		C.ulong(length),
		convertToObjCBool(b),
	)
	r0 = NSData_fromPointer(ret)
	return
}

func (x gen_NSData) InitWithContentsOfFile_asNSData(
	path NSStringRef,
) (
	r0 NSData,
) {
	ret := C.NSData_inst_initWithContentsOfFile(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)
	r0 = NSData_fromPointer(ret)
	return
}

func (x gen_NSData) InitWithContentsOfURL_asNSData(
	url NSURLRef,
) (
	r0 NSData,
) {
	ret := C.NSData_inst_initWithContentsOfURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
	)
	r0 = NSData_fromPointer(ret)
	return
}

func (x gen_NSData) InitWithData_asNSData(
	data NSDataRef,
) (
	r0 NSData,
) {
	ret := C.NSData_inst_initWithData(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
	)
	r0 = NSData_fromPointer(ret)
	return
}

func (x gen_NSData) IsEqualToData(
	other NSDataRef,
) (
	r0 bool,
) {
	ret := C.NSData_inst_isEqualToData(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(other),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSData) WriteToFile_atomically(
	path NSStringRef,
	useAuxiliaryFile bool,
) (
	r0 bool,
) {
	ret := C.NSData_inst_writeToFile_atomically(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(useAuxiliaryFile),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSData) WriteToURL_atomically(
	url NSURLRef,
	atomically bool,
) (
	r0 bool,
) {
	ret := C.NSData_inst_writeToURL_atomically(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(url),
		convertToObjCBool(atomically),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSData) Init_asNSData() (
	r0 NSData,
) {
	ret := C.NSData_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSData_fromPointer(ret)
	return
}

func (x gen_NSData) Bytes() (
	r0 unsafe.Pointer,
) {
	ret := C.NSData_inst_bytes(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = ret
	return
}

func (x gen_NSData) Length() (
	r0 NSUInteger,
) {
	ret := C.NSData_inst_length(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSUInteger(ret)
	return
}

func (x gen_NSData) Description() (
	r0 NSString,
) {
	ret := C.NSData_inst_description(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
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

func (x gen_NSDictionary) DescriptionWithLocale(
	locale objc.Ref,
) (
	r0 NSString,
) {
	ret := C.NSDictionary_inst_descriptionWithLocale(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(locale),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSDictionary) DescriptionWithLocale_indent(
	locale objc.Ref,
	level NSUInteger,
) (
	r0 NSString,
) {
	ret := C.NSDictionary_inst_descriptionWithLocale_indent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(locale),
		C.ulong(level),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSDictionary) FileExtensionHidden() (
	r0 bool,
) {
	ret := C.NSDictionary_inst_fileExtensionHidden(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSDictionary) FileGroupOwnerAccountID() (
	r0 NSNumber,
) {
	ret := C.NSDictionary_inst_fileGroupOwnerAccountID(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSNumber_fromPointer(ret)
	return
}

func (x gen_NSDictionary) FileGroupOwnerAccountName() (
	r0 NSString,
) {
	ret := C.NSDictionary_inst_fileGroupOwnerAccountName(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSDictionary) FileIsAppendOnly() (
	r0 bool,
) {
	ret := C.NSDictionary_inst_fileIsAppendOnly(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSDictionary) FileIsImmutable() (
	r0 bool,
) {
	ret := C.NSDictionary_inst_fileIsImmutable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSDictionary) FileOwnerAccountID() (
	r0 NSNumber,
) {
	ret := C.NSDictionary_inst_fileOwnerAccountID(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSNumber_fromPointer(ret)
	return
}

func (x gen_NSDictionary) FileOwnerAccountName() (
	r0 NSString,
) {
	ret := C.NSDictionary_inst_fileOwnerAccountName(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSDictionary) FilePosixPermissions() (
	r0 NSUInteger,
) {
	ret := C.NSDictionary_inst_filePosixPermissions(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSUInteger(ret)
	return
}

func (x gen_NSDictionary) FileSystemFileNumber() (
	r0 NSUInteger,
) {
	ret := C.NSDictionary_inst_fileSystemFileNumber(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSUInteger(ret)
	return
}

func (x gen_NSDictionary) FileSystemNumber() (
	r0 NSInteger,
) {
	ret := C.NSDictionary_inst_fileSystemNumber(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSInteger(ret)
	return
}

func (x gen_NSDictionary) FileType() (
	r0 NSString,
) {
	ret := C.NSDictionary_inst_fileType(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSDictionary) Init_asNSDictionary() (
	r0 NSDictionary,
) {
	ret := C.NSDictionary_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSDictionary) InitWithDictionary_asNSDictionary(
	otherDictionary NSDictionaryRef,
) (
	r0 NSDictionary,
) {
	ret := C.NSDictionary_inst_initWithDictionary(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(otherDictionary),
	)
	r0 = NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSDictionary) InitWithDictionary_copyItems_asNSDictionary(
	otherDictionary NSDictionaryRef,
	flag bool,
) (
	r0 NSDictionary,
) {
	ret := C.NSDictionary_inst_initWithDictionary_copyItems(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(otherDictionary),
		convertToObjCBool(flag),
	)
	r0 = NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSDictionary) InitWithObjects_forKeys_asNSDictionary(
	objects NSArrayRef,
	keys NSArrayRef,
) (
	r0 NSDictionary,
) {
	ret := C.NSDictionary_inst_initWithObjects_forKeys(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(objects),
		objc.RefPointer(keys),
	)
	r0 = NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSDictionary) IsEqualToDictionary(
	otherDictionary NSDictionaryRef,
) (
	r0 bool,
) {
	ret := C.NSDictionary_inst_isEqualToDictionary(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(otherDictionary),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSDictionary) KeysSortedByValueUsingSelector(
	comparator objc.Selector,
) (
	r0 NSArray,
) {
	ret := C.NSDictionary_inst_keysSortedByValueUsingSelector(
		unsafe.Pointer(x.Pointer()),
		comparator.SelectorAddress(),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSDictionary) Count() (
	r0 NSUInteger,
) {
	ret := C.NSDictionary_inst_count(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSUInteger(ret)
	return
}

func (x gen_NSDictionary) AllKeys() (
	r0 NSArray,
) {
	ret := C.NSDictionary_inst_allKeys(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSDictionary) AllValues() (
	r0 NSArray,
) {
	ret := C.NSDictionary_inst_allValues(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSDictionary) Description() (
	r0 NSString,
) {
	ret := C.NSDictionary_inst_description(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSDictionary) DescriptionInStringsFileFormat() (
	r0 NSString,
) {
	ret := C.NSDictionary_inst_descriptionInStringsFileFormat(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
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

func (x gen_NSNumber) DescriptionWithLocale(
	locale objc.Ref,
) (
	r0 NSString,
) {
	ret := C.NSNumber_inst_descriptionWithLocale(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(locale),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSNumber) InitWithBool(
	value bool,
) (
	r0 NSNumber,
) {
	ret := C.NSNumber_inst_initWithBool(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	r0 = NSNumber_fromPointer(ret)
	return
}

func (x gen_NSNumber) InitWithInt(
	value int32,
) (
	r0 NSNumber,
) {
	ret := C.NSNumber_inst_initWithInt(
		unsafe.Pointer(x.Pointer()),
		C.int(value),
	)
	r0 = NSNumber_fromPointer(ret)
	return
}

func (x gen_NSNumber) InitWithInteger(
	value NSInteger,
) (
	r0 NSNumber,
) {
	ret := C.NSNumber_inst_initWithInteger(
		unsafe.Pointer(x.Pointer()),
		C.long(value),
	)
	r0 = NSNumber_fromPointer(ret)
	return
}

func (x gen_NSNumber) InitWithUnsignedInt(
	value int32,
) (
	r0 NSNumber,
) {
	ret := C.NSNumber_inst_initWithUnsignedInt(
		unsafe.Pointer(x.Pointer()),
		C.int(value),
	)
	r0 = NSNumber_fromPointer(ret)
	return
}

func (x gen_NSNumber) InitWithUnsignedInteger(
	value NSUInteger,
) (
	r0 NSNumber,
) {
	ret := C.NSNumber_inst_initWithUnsignedInteger(
		unsafe.Pointer(x.Pointer()),
		C.ulong(value),
	)
	r0 = NSNumber_fromPointer(ret)
	return
}

func (x gen_NSNumber) IsEqualToNumber(
	number NSNumberRef,
) (
	r0 bool,
) {
	ret := C.NSNumber_inst_isEqualToNumber(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(number),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSNumber) Init_asNSNumber() (
	r0 NSNumber,
) {
	ret := C.NSNumber_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSNumber_fromPointer(ret)
	return
}

func (x gen_NSNumber) BoolValue() (
	r0 bool,
) {
	ret := C.NSNumber_inst_boolValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSNumber) IntValue() (
	r0 int32,
) {
	ret := C.NSNumber_inst_intValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = int32(ret)
	return
}

func (x gen_NSNumber) IntegerValue() (
	r0 NSInteger,
) {
	ret := C.NSNumber_inst_integerValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSInteger(ret)
	return
}

func (x gen_NSNumber) UnsignedIntegerValue() (
	r0 NSUInteger,
) {
	ret := C.NSNumber_inst_unsignedIntegerValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSUInteger(ret)
	return
}

func (x gen_NSNumber) UnsignedIntValue() (
	r0 int32,
) {
	ret := C.NSNumber_inst_unsignedIntValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = int32(ret)
	return
}

func (x gen_NSNumber) StringValue() (
	r0 NSString,
) {
	ret := C.NSNumber_inst_stringValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
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

func (x gen_NSRunLoop) CancelPerformSelectorsWithTarget(
	target objc.Ref,
) {
	C.NSRunLoop_inst_cancelPerformSelectorsWithTarget(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(target),
	)
	return
}

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

func (x gen_NSRunLoop) Run() {
	C.NSRunLoop_inst_run(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSRunLoop) Init_asNSRunLoop() (
	r0 NSRunLoop,
) {
	ret := C.NSRunLoop_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSRunLoop_fromPointer(ret)
	return
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

func (x gen_NSString) CanBeConvertedToEncoding(
	encoding NSStringEncoding,
) (
	r0 bool,
) {
	ret := C.NSString_inst_canBeConvertedToEncoding(
		unsafe.Pointer(x.Pointer()),
		C.ulong(encoding),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSString) CharacterAtIndex(
	index NSUInteger,
) (
	r0 Unichar,
) {
	ret := C.NSString_inst_characterAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(index),
	)
	r0 = Unichar(ret)
	return
}

func (x gen_NSString) CompletePathIntoString_caseSensitive_matchesIntoArray_filterTypes(
	outputName NSStringRef,
	flag bool,
	outputArray NSArrayRef,
	filterTypes NSArrayRef,
) (
	r0 NSUInteger,
) {
	ret := C.NSString_inst_completePathIntoString_caseSensitive_matchesIntoArray_filterTypes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(outputName),
		convertToObjCBool(flag),
		objc.RefPointer(outputArray),
		objc.RefPointer(filterTypes),
	)
	r0 = NSUInteger(ret)
	return
}

func (x gen_NSString) ComponentsSeparatedByString(
	separator NSStringRef,
) (
	r0 NSArray,
) {
	ret := C.NSString_inst_componentsSeparatedByString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(separator),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSString) ContainsString(
	str NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSString_inst_containsString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSString) DataUsingEncoding(
	encoding NSStringEncoding,
) (
	r0 NSData,
) {
	ret := C.NSString_inst_dataUsingEncoding(
		unsafe.Pointer(x.Pointer()),
		C.ulong(encoding),
	)
	r0 = NSData_fromPointer(ret)
	return
}

func (x gen_NSString) DataUsingEncoding_allowLossyConversion(
	encoding NSStringEncoding,
	lossy bool,
) (
	r0 NSData,
) {
	ret := C.NSString_inst_dataUsingEncoding_allowLossyConversion(
		unsafe.Pointer(x.Pointer()),
		C.ulong(encoding),
		convertToObjCBool(lossy),
	)
	r0 = NSData_fromPointer(ret)
	return
}

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

func (x gen_NSString) HasPrefix(
	str NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSString_inst_hasPrefix(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSString) HasSuffix(
	str NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSString_inst_hasSuffix(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSString) Init_asNSString() (
	r0 NSString,
) {
	ret := C.NSString_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) InitWithBytes_length_encoding_asNSString(
	bytes unsafe.Pointer,
	len NSUInteger,
	encoding NSStringEncoding,
) (
	r0 NSString,
) {
	ret := C.NSString_inst_initWithBytes_length_encoding(
		unsafe.Pointer(x.Pointer()),
		bytes,
		C.ulong(len),
		C.ulong(encoding),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) InitWithBytesNoCopy_length_encoding_freeWhenDone_asNSString(
	bytes unsafe.Pointer,
	len NSUInteger,
	encoding NSStringEncoding,
	freeBuffer bool,
) (
	r0 NSString,
) {
	ret := C.NSString_inst_initWithBytesNoCopy_length_encoding_freeWhenDone(
		unsafe.Pointer(x.Pointer()),
		bytes,
		C.ulong(len),
		C.ulong(encoding),
		convertToObjCBool(freeBuffer),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) InitWithData_encoding_asNSString(
	data NSDataRef,
	encoding NSStringEncoding,
) (
	r0 NSString,
) {
	ret := C.NSString_inst_initWithData_encoding(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		C.ulong(encoding),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) InitWithString_asNSString(
	aString NSStringRef,
) (
	r0 NSString,
) {
	ret := C.NSString_inst_initWithString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(aString),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) IsEqualToString(
	aString NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSString_inst_isEqualToString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(aString),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSString) LengthOfBytesUsingEncoding(
	enc NSStringEncoding,
) (
	r0 NSUInteger,
) {
	ret := C.NSString_inst_lengthOfBytesUsingEncoding(
		unsafe.Pointer(x.Pointer()),
		C.ulong(enc),
	)
	r0 = NSUInteger(ret)
	return
}

func (x gen_NSString) LocalizedCaseInsensitiveContainsString(
	str NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSString_inst_localizedCaseInsensitiveContainsString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSString) LocalizedStandardContainsString(
	str NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSString_inst_localizedStandardContainsString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSString) MaximumLengthOfBytesUsingEncoding(
	enc NSStringEncoding,
) (
	r0 NSUInteger,
) {
	ret := C.NSString_inst_maximumLengthOfBytesUsingEncoding(
		unsafe.Pointer(x.Pointer()),
		C.ulong(enc),
	)
	r0 = NSUInteger(ret)
	return
}

func (x gen_NSString) PropertyList() (
	r0 objc.Object,
) {
	ret := C.NSString_inst_propertyList(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSString) PropertyListFromStringsFileFormat() (
	r0 NSDictionary,
) {
	ret := C.NSString_inst_propertyListFromStringsFileFormat(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSString) SizeWithAttributes(
	attrs NSDictionaryRef,
) (
	r0 NSSize,
) {
	ret := C.NSString_inst_sizeWithAttributes(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(attrs),
	)
	r0 = *(*NSSize)(unsafe.Pointer(&ret))
	return
}

func (x gen_NSString) StringByAppendingPathComponent(
	str NSStringRef,
) (
	r0 NSString,
) {
	ret := C.NSString_inst_stringByAppendingPathComponent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) StringByAppendingPathExtension(
	str NSStringRef,
) (
	r0 NSString,
) {
	ret := C.NSString_inst_stringByAppendingPathExtension(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(str),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) StringByAppendingString(
	aString NSStringRef,
) (
	r0 NSString,
) {
	ret := C.NSString_inst_stringByAppendingString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(aString),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) StringByPaddingToLength_withString_startingAtIndex(
	newLength NSUInteger,
	padString NSStringRef,
	padIndex NSUInteger,
) (
	r0 NSString,
) {
	ret := C.NSString_inst_stringByPaddingToLength_withString_startingAtIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(newLength),
		objc.RefPointer(padString),
		C.ulong(padIndex),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) StringByReplacingOccurrencesOfString_withString(
	target NSStringRef,
	replacement NSStringRef,
) (
	r0 NSString,
) {
	ret := C.NSString_inst_stringByReplacingOccurrencesOfString_withString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(target),
		objc.RefPointer(replacement),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) StringsByAppendingPaths(
	paths NSArrayRef,
) (
	r0 NSArray,
) {
	ret := C.NSString_inst_stringsByAppendingPaths(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(paths),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSString) SubstringFromIndex(
	from NSUInteger,
) (
	r0 NSString,
) {
	ret := C.NSString_inst_substringFromIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(from),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) SubstringToIndex(
	to NSUInteger,
) (
	r0 NSString,
) {
	ret := C.NSString_inst_substringToIndex(
		unsafe.Pointer(x.Pointer()),
		C.ulong(to),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) VariantFittingPresentationWidth(
	width NSInteger,
) (
	r0 NSString,
) {
	ret := C.NSString_inst_variantFittingPresentationWidth(
		unsafe.Pointer(x.Pointer()),
		C.long(width),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) Length() (
	r0 NSUInteger,
) {
	ret := C.NSString_inst_length(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSUInteger(ret)
	return
}

func (x gen_NSString) Hash() (
	r0 NSUInteger,
) {
	ret := C.NSString_inst_hash(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSUInteger(ret)
	return
}

func (x gen_NSString) LowercaseString() (
	r0 NSString,
) {
	ret := C.NSString_inst_lowercaseString(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) LocalizedLowercaseString() (
	r0 NSString,
) {
	ret := C.NSString_inst_localizedLowercaseString(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) UppercaseString() (
	r0 NSString,
) {
	ret := C.NSString_inst_uppercaseString(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) LocalizedUppercaseString() (
	r0 NSString,
) {
	ret := C.NSString_inst_localizedUppercaseString(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) CapitalizedString() (
	r0 NSString,
) {
	ret := C.NSString_inst_capitalizedString(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) LocalizedCapitalizedString() (
	r0 NSString,
) {
	ret := C.NSString_inst_localizedCapitalizedString(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) DecomposedStringWithCanonicalMapping() (
	r0 NSString,
) {
	ret := C.NSString_inst_decomposedStringWithCanonicalMapping(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) DecomposedStringWithCompatibilityMapping() (
	r0 NSString,
) {
	ret := C.NSString_inst_decomposedStringWithCompatibilityMapping(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) PrecomposedStringWithCanonicalMapping() (
	r0 NSString,
) {
	ret := C.NSString_inst_precomposedStringWithCanonicalMapping(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) PrecomposedStringWithCompatibilityMapping() (
	r0 NSString,
) {
	ret := C.NSString_inst_precomposedStringWithCompatibilityMapping(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) IntValue() (
	r0 int32,
) {
	ret := C.NSString_inst_intValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = int32(ret)
	return
}

func (x gen_NSString) IntegerValue() (
	r0 NSInteger,
) {
	ret := C.NSString_inst_integerValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSInteger(ret)
	return
}

func (x gen_NSString) BoolValue() (
	r0 bool,
) {
	ret := C.NSString_inst_boolValue(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSString) Description() (
	r0 NSString,
) {
	ret := C.NSString_inst_description(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) FastestEncoding() (
	r0 NSStringEncoding,
) {
	ret := C.NSString_inst_fastestEncoding(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSStringEncoding(ret)
	return
}

func (x gen_NSString) SmallestEncoding() (
	r0 NSStringEncoding,
) {
	ret := C.NSString_inst_smallestEncoding(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSStringEncoding(ret)
	return
}

func (x gen_NSString) PathComponents() (
	r0 NSArray,
) {
	ret := C.NSString_inst_pathComponents(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSString) IsAbsolutePath() (
	r0 bool,
) {
	ret := C.NSString_inst_isAbsolutePath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSString) LastPathComponent() (
	r0 NSString,
) {
	ret := C.NSString_inst_lastPathComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) PathExtension() (
	r0 NSString,
) {
	ret := C.NSString_inst_pathExtension(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) StringByAbbreviatingWithTildeInPath() (
	r0 NSString,
) {
	ret := C.NSString_inst_stringByAbbreviatingWithTildeInPath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) StringByDeletingLastPathComponent() (
	r0 NSString,
) {
	ret := C.NSString_inst_stringByDeletingLastPathComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) StringByDeletingPathExtension() (
	r0 NSString,
) {
	ret := C.NSString_inst_stringByDeletingPathExtension(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) StringByExpandingTildeInPath() (
	r0 NSString,
) {
	ret := C.NSString_inst_stringByExpandingTildeInPath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) StringByResolvingSymlinksInPath() (
	r0 NSString,
) {
	ret := C.NSString_inst_stringByResolvingSymlinksInPath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) StringByStandardizingPath() (
	r0 NSString,
) {
	ret := C.NSString_inst_stringByStandardizingPath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSString) StringByRemovingPercentEncoding() (
	r0 NSString,
) {
	ret := C.NSString_inst_stringByRemovingPercentEncoding(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
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

func (x gen_NSThread) Cancel() {
	C.NSThread_inst_cancel(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSThread) Init_asNSThread() (
	r0 NSThread,
) {
	ret := C.NSThread_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSThread_fromPointer(ret)
	return
}

func (x gen_NSThread) InitWithTarget_selector_object_asNSThread(
	target objc.Ref,
	selector objc.Selector,
	argument objc.Ref,
) (
	r0 NSThread,
) {
	ret := C.NSThread_inst_initWithTarget_selector_object(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(target),
		selector.SelectorAddress(),
		objc.RefPointer(argument),
	)
	r0 = NSThread_fromPointer(ret)
	return
}

func (x gen_NSThread) Main() {
	C.NSThread_inst_main(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSThread) Start() {
	C.NSThread_inst_start(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSThread) IsExecuting() (
	r0 bool,
) {
	ret := C.NSThread_inst_isExecuting(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSThread) IsFinished() (
	r0 bool,
) {
	ret := C.NSThread_inst_isFinished(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSThread) IsCancelled() (
	r0 bool,
) {
	ret := C.NSThread_inst_isCancelled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSThread) IsMainThread() (
	r0 bool,
) {
	ret := C.NSThread_inst_isMainThread(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSThread) Name() (
	r0 NSString,
) {
	ret := C.NSThread_inst_name(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSThread) SetName(
	value NSStringRef,
) {
	C.NSThread_inst_setName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_NSThread) StackSize() (
	r0 NSUInteger,
) {
	ret := C.NSThread_inst_stackSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSUInteger(ret)
	return
}

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

func (x gen_NSURL) URLByAppendingPathComponent(
	pathComponent NSStringRef,
) (
	r0 NSURL,
) {
	ret := C.NSURL_inst_URLByAppendingPathComponent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pathComponent),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) URLByAppendingPathComponent_isDirectory(
	pathComponent NSStringRef,
	isDirectory bool,
) (
	r0 NSURL,
) {
	ret := C.NSURL_inst_URLByAppendingPathComponent_isDirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pathComponent),
		convertToObjCBool(isDirectory),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) URLByAppendingPathExtension(
	pathExtension NSStringRef,
) (
	r0 NSURL,
) {
	ret := C.NSURL_inst_URLByAppendingPathExtension(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(pathExtension),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) FileReferenceURL() (
	r0 NSURL,
) {
	ret := C.NSURL_inst_fileReferenceURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) InitAbsoluteURLWithDataRepresentation_relativeToURL_asNSURL(
	data NSDataRef,
	baseURL NSURLRef,
) (
	r0 NSURL,
) {
	ret := C.NSURL_inst_initAbsoluteURLWithDataRepresentation_relativeToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(baseURL),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) InitFileURLWithPath_asNSURL(
	path NSStringRef,
) (
	r0 NSURL,
) {
	ret := C.NSURL_inst_initFileURLWithPath(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) InitFileURLWithPath_isDirectory_asNSURL(
	path NSStringRef,
	isDir bool,
) (
	r0 NSURL,
) {
	ret := C.NSURL_inst_initFileURLWithPath_isDirectory(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(isDir),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) InitFileURLWithPath_isDirectory_relativeToURL_asNSURL(
	path NSStringRef,
	isDir bool,
	baseURL NSURLRef,
) (
	r0 NSURL,
) {
	ret := C.NSURL_inst_initFileURLWithPath_isDirectory_relativeToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		convertToObjCBool(isDir),
		objc.RefPointer(baseURL),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) InitFileURLWithPath_relativeToURL_asNSURL(
	path NSStringRef,
	baseURL NSURLRef,
) (
	r0 NSURL,
) {
	ret := C.NSURL_inst_initFileURLWithPath_relativeToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(path),
		objc.RefPointer(baseURL),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) InitWithDataRepresentation_relativeToURL_asNSURL(
	data NSDataRef,
	baseURL NSURLRef,
) (
	r0 NSURL,
) {
	ret := C.NSURL_inst_initWithDataRepresentation_relativeToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(baseURL),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) InitWithString_asNSURL(
	URLString NSStringRef,
) (
	r0 NSURL,
) {
	ret := C.NSURL_inst_initWithString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(URLString),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) InitWithString_relativeToURL_asNSURL(
	URLString NSStringRef,
	baseURL NSURLRef,
) (
	r0 NSURL,
) {
	ret := C.NSURL_inst_initWithString_relativeToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(URLString),
		objc.RefPointer(baseURL),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) IsFileReferenceURL() (
	r0 bool,
) {
	ret := C.NSURL_inst_isFileReferenceURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSURL) RemoveAllCachedResourceValues() {
	C.NSURL_inst_removeAllCachedResourceValues(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSURL) StartAccessingSecurityScopedResource() (
	r0 bool,
) {
	ret := C.NSURL_inst_startAccessingSecurityScopedResource(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSURL) StopAccessingSecurityScopedResource() {
	C.NSURL_inst_stopAccessingSecurityScopedResource(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_NSURL) Init_asNSURL() (
	r0 NSURL,
) {
	ret := C.NSURL_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) DataRepresentation() (
	r0 NSData,
) {
	ret := C.NSURL_inst_dataRepresentation(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSData_fromPointer(ret)
	return
}

func (x gen_NSURL) IsFileURL() (
	r0 bool,
) {
	ret := C.NSURL_inst_isFileURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSURL) AbsoluteString() (
	r0 NSString,
) {
	ret := C.NSURL_inst_absoluteString(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSURL) AbsoluteURL() (
	r0 NSURL,
) {
	ret := C.NSURL_inst_absoluteURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) BaseURL() (
	r0 NSURL,
) {
	ret := C.NSURL_inst_baseURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) Fragment() (
	r0 NSString,
) {
	ret := C.NSURL_inst_fragment(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSURL) Host() (
	r0 NSString,
) {
	ret := C.NSURL_inst_host(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSURL) LastPathComponent() (
	r0 NSString,
) {
	ret := C.NSURL_inst_lastPathComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSURL) Password() (
	r0 NSString,
) {
	ret := C.NSURL_inst_password(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSURL) Path() (
	r0 NSString,
) {
	ret := C.NSURL_inst_path(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSURL) PathComponents() (
	r0 NSArray,
) {
	ret := C.NSURL_inst_pathComponents(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSURL) PathExtension() (
	r0 NSString,
) {
	ret := C.NSURL_inst_pathExtension(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSURL) Port() (
	r0 NSNumber,
) {
	ret := C.NSURL_inst_port(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSNumber_fromPointer(ret)
	return
}

func (x gen_NSURL) Query() (
	r0 NSString,
) {
	ret := C.NSURL_inst_query(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSURL) RelativePath() (
	r0 NSString,
) {
	ret := C.NSURL_inst_relativePath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSURL) RelativeString() (
	r0 NSString,
) {
	ret := C.NSURL_inst_relativeString(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSURL) ResourceSpecifier() (
	r0 NSString,
) {
	ret := C.NSURL_inst_resourceSpecifier(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSURL) Scheme() (
	r0 NSString,
) {
	ret := C.NSURL_inst_scheme(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSURL) StandardizedURL() (
	r0 NSURL,
) {
	ret := C.NSURL_inst_standardizedURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) User() (
	r0 NSString,
) {
	ret := C.NSURL_inst_user(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSURL) FilePathURL() (
	r0 NSURL,
) {
	ret := C.NSURL_inst_filePathURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) URLByDeletingLastPathComponent() (
	r0 NSURL,
) {
	ret := C.NSURL_inst_URLByDeletingLastPathComponent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) URLByDeletingPathExtension() (
	r0 NSURL,
) {
	ret := C.NSURL_inst_URLByDeletingPathExtension(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) URLByResolvingSymlinksInPath() (
	r0 NSURL,
) {
	ret := C.NSURL_inst_URLByResolvingSymlinksInPath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) URLByStandardizingPath() (
	r0 NSURL,
) {
	ret := C.NSURL_inst_URLByStandardizingPath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURL) HasDirectoryPath() (
	r0 bool,
) {
	ret := C.NSURL_inst_hasDirectoryPath(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
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

func (x gen_NSURLRequest) InitWithURL_asNSURLRequest(
	URL NSURLRef,
) (
	r0 NSURLRequest,
) {
	ret := C.NSURLRequest_inst_initWithURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(URL),
	)
	r0 = NSURLRequest_fromPointer(ret)
	return
}

func (x gen_NSURLRequest) ValueForHTTPHeaderField(
	field NSStringRef,
) (
	r0 NSString,
) {
	ret := C.NSURLRequest_inst_valueForHTTPHeaderField(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(field),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSURLRequest) Init_asNSURLRequest() (
	r0 NSURLRequest,
) {
	ret := C.NSURLRequest_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSURLRequest_fromPointer(ret)
	return
}

func (x gen_NSURLRequest) HTTPMethod() (
	r0 NSString,
) {
	ret := C.NSURLRequest_inst_HTTPMethod(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSURLRequest) URL() (
	r0 NSURL,
) {
	ret := C.NSURLRequest_inst_URL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURLRequest) HTTPBody() (
	r0 NSData,
) {
	ret := C.NSURLRequest_inst_HTTPBody(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSData_fromPointer(ret)
	return
}

func (x gen_NSURLRequest) MainDocumentURL() (
	r0 NSURL,
) {
	ret := C.NSURLRequest_inst_mainDocumentURL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSURLRequest) AllHTTPHeaderFields() (
	r0 NSDictionary,
) {
	ret := C.NSURLRequest_inst_allHTTPHeaderFields(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSURLRequest) HTTPShouldHandleCookies() (
	r0 bool,
) {
	ret := C.NSURLRequest_inst_HTTPShouldHandleCookies(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSURLRequest) HTTPShouldUsePipelining() (
	r0 bool,
) {
	ret := C.NSURLRequest_inst_HTTPShouldUsePipelining(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSURLRequest) AllowsCellularAccess() (
	r0 bool,
) {
	ret := C.NSURLRequest_inst_allowsCellularAccess(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSURLRequest) AllowsConstrainedNetworkAccess() (
	r0 bool,
) {
	ret := C.NSURLRequest_inst_allowsConstrainedNetworkAccess(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSURLRequest) AllowsExpensiveNetworkAccess() (
	r0 bool,
) {
	ret := C.NSURLRequest_inst_allowsExpensiveNetworkAccess(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSURLRequest) AssumesHTTP3Capable() (
	r0 bool,
) {
	ret := C.NSURLRequest_inst_assumesHTTP3Capable(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
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

func (x gen_NSUserDefaults) URLForKey(
	defaultName NSStringRef,
) (
	r0 NSURL,
) {
	ret := C.NSUserDefaults_inst_URLForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)
	r0 = NSURL_fromPointer(ret)
	return
}

func (x gen_NSUserDefaults) AddSuiteNamed(
	suiteName NSStringRef,
) {
	C.NSUserDefaults_inst_addSuiteNamed(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(suiteName),
	)
	return
}

func (x gen_NSUserDefaults) ArrayForKey(
	defaultName NSStringRef,
) (
	r0 NSArray,
) {
	ret := C.NSUserDefaults_inst_arrayForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSUserDefaults) BoolForKey(
	defaultName NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSUserDefaults_inst_boolForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSUserDefaults) DataForKey(
	defaultName NSStringRef,
) (
	r0 NSData,
) {
	ret := C.NSUserDefaults_inst_dataForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)
	r0 = NSData_fromPointer(ret)
	return
}

func (x gen_NSUserDefaults) DictionaryForKey(
	defaultName NSStringRef,
) (
	r0 NSDictionary,
) {
	ret := C.NSUserDefaults_inst_dictionaryForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)
	r0 = NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSUserDefaults) DictionaryRepresentation() (
	r0 NSDictionary,
) {
	ret := C.NSUserDefaults_inst_dictionaryRepresentation(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSUserDefaults) Init_asNSUserDefaults() (
	r0 NSUserDefaults,
) {
	ret := C.NSUserDefaults_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSUserDefaults_fromPointer(ret)
	return
}

func (x gen_NSUserDefaults) InitWithSuiteName_asNSUserDefaults(
	suitename NSStringRef,
) (
	r0 NSUserDefaults,
) {
	ret := C.NSUserDefaults_inst_initWithSuiteName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(suitename),
	)
	r0 = NSUserDefaults_fromPointer(ret)
	return
}

func (x gen_NSUserDefaults) IntegerForKey(
	defaultName NSStringRef,
) (
	r0 NSInteger,
) {
	ret := C.NSUserDefaults_inst_integerForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)
	r0 = NSInteger(ret)
	return
}

func (x gen_NSUserDefaults) ObjectForKey(
	defaultName NSStringRef,
) (
	r0 objc.Object,
) {
	ret := C.NSUserDefaults_inst_objectForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_NSUserDefaults) ObjectIsForcedForKey(
	key NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSUserDefaults_inst_objectIsForcedForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSUserDefaults) ObjectIsForcedForKey_inDomain(
	key NSStringRef,
	domain NSStringRef,
) (
	r0 bool,
) {
	ret := C.NSUserDefaults_inst_objectIsForcedForKey_inDomain(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(key),
		objc.RefPointer(domain),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSUserDefaults) PersistentDomainForName(
	domainName NSStringRef,
) (
	r0 NSDictionary,
) {
	ret := C.NSUserDefaults_inst_persistentDomainForName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domainName),
	)
	r0 = NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSUserDefaults) RegisterDefaults(
	registrationDictionary NSDictionaryRef,
) {
	C.NSUserDefaults_inst_registerDefaults(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(registrationDictionary),
	)
	return
}

func (x gen_NSUserDefaults) RemoveObjectForKey(
	defaultName NSStringRef,
) {
	C.NSUserDefaults_inst_removeObjectForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)
	return
}

func (x gen_NSUserDefaults) RemovePersistentDomainForName(
	domainName NSStringRef,
) {
	C.NSUserDefaults_inst_removePersistentDomainForName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domainName),
	)
	return
}

func (x gen_NSUserDefaults) RemoveSuiteNamed(
	suiteName NSStringRef,
) {
	C.NSUserDefaults_inst_removeSuiteNamed(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(suiteName),
	)
	return
}

func (x gen_NSUserDefaults) RemoveVolatileDomainForName(
	domainName NSStringRef,
) {
	C.NSUserDefaults_inst_removeVolatileDomainForName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domainName),
	)
	return
}

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

func (x gen_NSUserDefaults) StringArrayForKey(
	defaultName NSStringRef,
) (
	r0 NSArray,
) {
	ret := C.NSUserDefaults_inst_stringArrayForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)
	r0 = NSArray_fromPointer(ret)
	return
}

func (x gen_NSUserDefaults) StringForKey(
	defaultName NSStringRef,
) (
	r0 NSString,
) {
	ret := C.NSUserDefaults_inst_stringForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(defaultName),
	)
	r0 = NSString_fromPointer(ret)
	return
}

func (x gen_NSUserDefaults) Synchronize() (
	r0 bool,
) {
	ret := C.NSUserDefaults_inst_synchronize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_NSUserDefaults) VolatileDomainForName(
	domainName NSStringRef,
) (
	r0 NSDictionary,
) {
	ret := C.NSUserDefaults_inst_volatileDomainForName(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(domainName),
	)
	r0 = NSDictionary_fromPointer(ret)
	return
}

func (x gen_NSUserDefaults) VolatileDomainNames() (
	r0 NSArray,
) {
	ret := C.NSUserDefaults_inst_volatileDomainNames(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = NSArray_fromPointer(ret)
	return
}
