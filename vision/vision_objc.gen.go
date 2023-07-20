package vision

import (
	core "github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"unsafe"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -lobjc -framework Vision
#define __OBJC2__ 1
#include <objc/message.h>
#include <stdlib.h>

#include <Vision/Vision.h>

bool vision_convertObjCBool(BOOL b) {
	if (b) { return true; }
	return false;
}

// Creates a NSString from a C string
static void *createNSStringFromCString(char *cString) {
    return [NSString stringWithCString: cString encoding: NSUTF8StringEncoding];
}

// Creates a C string from a NSString
static char *createCStringFromNSString(void *objcString)
{
    return [objcString UTF8String];
}


void* VNClassifyImageRequest_type_Alloc() {
	return [VNClassifyImageRequest
		alloc];
}
void* VNGenerateImageFeaturePrintRequest_type_Alloc() {
	return [VNGenerateImageFeaturePrintRequest
		alloc];
}
void* VNImageBasedRequest_type_Alloc() {
	return [VNImageBasedRequest
		alloc];
}
void* VNImageRequestHandler_type_Alloc() {
	return [VNImageRequestHandler
		alloc];
}
void* VNObservation_type_Alloc() {
	return [VNObservation
		alloc];
}
void* VNRecognizedTextObservation_type_Alloc() {
	return [VNRecognizedTextObservation
		alloc];
}
void* VNRecognizeTextRequest_type_Alloc() {
	return [VNRecognizeTextRequest
		alloc];
}
void* VNRequest_type_Alloc() {
	return [VNRequest
		alloc];
}
unsigned long VNRequest_type_CurrentRevision() {
	return [VNRequest
		currentRevision];
}
unsigned long VNRequest_type_DefaultRevision() {
	return [VNRequest
		defaultRevision];
}


void* VNClassifyImageRequest_inst_SupportedIdentifiersAndReturnError(void *id, void* error) {
	return [(VNClassifyImageRequest*)id
		supportedIdentifiersAndReturnError: error];
}

void* VNClassifyImageRequest_inst_Init(void *id) {
	return [(VNClassifyImageRequest*)id
		init];
}

void* VNClassifyImageRequest_inst_Results(void *id) {
	return [(VNClassifyImageRequest*)id
		results];
}

void* VNGenerateImageFeaturePrintRequest_inst_Init(void *id) {
	return [(VNGenerateImageFeaturePrintRequest*)id
		init];
}

void* VNGenerateImageFeaturePrintRequest_inst_Results(void *id) {
	return [(VNGenerateImageFeaturePrintRequest*)id
		results];
}

void* VNImageBasedRequest_inst_Init(void *id) {
	return [(VNImageBasedRequest*)id
		init];
}

// VNImageRequestHandler_inst_Init marked unavailable.

void* VNImageRequestHandler_inst_InitWithDataOptions(void *id, void* imageData, void* options) {
	return [(VNImageRequestHandler*)id
		initWithData: imageData
		options: options];
}

void* VNImageRequestHandler_inst_InitWithURLOptions(void *id, void* imageURL, void* options) {
	return [(VNImageRequestHandler*)id
		initWithURL: imageURL
		options: options];
}

BOOL VNImageRequestHandler_inst_PerformRequestsError(void *id, void* requests, void* error) {
	return [(VNImageRequestHandler*)id
		performRequests: requests
		error: error];
}

void* VNObservation_inst_Init(void *id) {
	return [(VNObservation*)id
		init];
}

void* VNRecognizedTextObservation_inst_TopCandidates(void *id, unsigned long maxCandidateCount) {
	return [(VNRecognizedTextObservation*)id
		topCandidates: maxCandidateCount];
}

void* VNRecognizedTextObservation_inst_Init(void *id) {
	return [(VNRecognizedTextObservation*)id
		init];
}

void* VNRecognizeTextRequest_inst_SupportedRecognitionLanguagesAndReturnError(void *id, void* error) {
	return [(VNRecognizeTextRequest*)id
		supportedRecognitionLanguagesAndReturnError: error];
}

void* VNRecognizeTextRequest_inst_Init(void *id) {
	return [(VNRecognizeTextRequest*)id
		init];
}

void* VNRecognizeTextRequest_inst_Results(void *id) {
	return [(VNRecognizeTextRequest*)id
		results];
}

void VNRequest_inst_Cancel(void *id) {
	[(VNRequest*)id
		cancel];
}

void* VNRequest_inst_Init(void *id) {
	return [(VNRequest*)id
		init];
}

void* VNRequest_inst_SupportedComputeStageDevicesAndReturnError(void *id, void* error) {
	return [(VNRequest*)id
		supportedComputeStageDevicesAndReturnError: error];
}

void* VNRequest_inst_Results(void *id) {
	return [(VNRequest*)id
		results];
}


BOOL vision_objc_bool_true = YES;
BOOL vision_objc_bool_false = NO;

*/
import "C"

func convertObjCBoolToGo(b C.BOOL) bool {
	// NOTE: the prefix here is used to namespace these since the linker will
	// otherwise report a "duplicate symbol" because the C functions have the
	// same name.
	return bool(C.vision_convertObjCBool(b))
}

func convertToObjCBool(b bool) C.BOOL {
	if b {
		return C.vision_objc_bool_true
	}
	return C.vision_objc_bool_false
}

// VNClassifyImageRequest_Alloc is undocumented.
func VNClassifyImageRequest_Alloc() VNClassifyImageRequest {
	ret := C.VNClassifyImageRequest_type_Alloc()

	return VNClassifyImageRequest_FromPointer(ret)
}

// VNGenerateImageFeaturePrintRequest_Alloc is undocumented.
func VNGenerateImageFeaturePrintRequest_Alloc() VNGenerateImageFeaturePrintRequest {
	ret := C.VNGenerateImageFeaturePrintRequest_type_Alloc()

	return VNGenerateImageFeaturePrintRequest_FromPointer(ret)
}

// VNImageBasedRequest_Alloc is undocumented.
func VNImageBasedRequest_Alloc() VNImageBasedRequest {
	ret := C.VNImageBasedRequest_type_Alloc()

	return VNImageBasedRequest_FromPointer(ret)
}

// VNImageRequestHandler_Alloc is undocumented.
func VNImageRequestHandler_Alloc() VNImageRequestHandler {
	ret := C.VNImageRequestHandler_type_Alloc()

	return VNImageRequestHandler_FromPointer(ret)
}

// VNObservation_Alloc is undocumented.
func VNObservation_Alloc() VNObservation {
	ret := C.VNObservation_type_Alloc()

	return VNObservation_FromPointer(ret)
}

// VNRecognizedTextObservation_Alloc is undocumented.
func VNRecognizedTextObservation_Alloc() VNRecognizedTextObservation {
	ret := C.VNRecognizedTextObservation_type_Alloc()

	return VNRecognizedTextObservation_FromPointer(ret)
}

// VNRecognizeTextRequest_Alloc is undocumented.
func VNRecognizeTextRequest_Alloc() VNRecognizeTextRequest {
	ret := C.VNRecognizeTextRequest_type_Alloc()

	return VNRecognizeTextRequest_FromPointer(ret)
}

// VNRequest_Alloc is undocumented.
func VNRequest_Alloc() VNRequest {
	ret := C.VNRequest_type_Alloc()

	return VNRequest_FromPointer(ret)
}

// VNRequest_CurrentRevision returns the current revison supported by the request.
//
// See https://developer.apple.com/documentation/vision/vnrequest/2967108-currentrevision?language=objc for details.
func VNRequest_CurrentRevision() core.NSUInteger {
	ret := C.VNRequest_type_CurrentRevision()

	return core.NSUInteger(ret)
}

// VNRequest_DefaultRevision returns the revision of the latest request for the particular SDK linked with the client application.
//
// See https://developer.apple.com/documentation/vision/vnrequest/2967109-defaultrevision?language=objc for details.
func VNRequest_DefaultRevision() core.NSUInteger {
	ret := C.VNRequest_type_DefaultRevision()

	return core.NSUInteger(ret)
}

type VNClassifyImageRequestRef interface {
	Pointer() uintptr
	Init_AsVNClassifyImageRequest() VNClassifyImageRequest
}

type gen_VNClassifyImageRequest struct {
	VNImageBasedRequest
}

func VNClassifyImageRequest_FromPointer(ptr unsafe.Pointer) VNClassifyImageRequest {
	return VNClassifyImageRequest{gen_VNClassifyImageRequest{
		VNImageBasedRequest_FromPointer(ptr),
	}}
}

func VNClassifyImageRequest_FromRef(ref objc.Ref) VNClassifyImageRequest {
	return VNClassifyImageRequest_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// SupportedIdentifiersAndReturnError returns the classification identifiers that the request supports in its current configuration.
//
// See https://developer.apple.com/documentation/vision/vnclassifyimagerequest/3750957-supportedidentifiersandreturnerr?language=objc for details.
func (x gen_VNClassifyImageRequest) SupportedIdentifiersAndReturnError(
	error core.NSErrorRef,
) core.NSArray {
	ret := C.VNClassifyImageRequest_inst_SupportedIdentifiersAndReturnError(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(error),
	)

	return core.NSArray_FromPointer(ret)
}

// Init initializes a new instance of the VNClassifyImageRequest class.
func (x gen_VNClassifyImageRequest) Init() VNClassifyImageRequest {
	ret := C.VNClassifyImageRequest_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return VNClassifyImageRequest_FromPointer(ret)
}

// Init_AsVNClassifyImageRequest is a typed version of Init.
func (x gen_VNClassifyImageRequest) Init_AsVNClassifyImageRequest() VNClassifyImageRequest {
	ret := C.VNClassifyImageRequest_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return VNClassifyImageRequest_FromPointer(ret)
}

// Results returns the results of the image classification request.
//
// See https://developer.apple.com/documentation/vision/vnclassifyimagerequest/3750956-results?language=objc for details.
func (x gen_VNClassifyImageRequest) Results() core.NSArray {
	ret := C.VNClassifyImageRequest_inst_Results(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

type VNGenerateImageFeaturePrintRequestRef interface {
	Pointer() uintptr
	Init_AsVNGenerateImageFeaturePrintRequest() VNGenerateImageFeaturePrintRequest
}

type gen_VNGenerateImageFeaturePrintRequest struct {
	VNImageBasedRequest
}

func VNGenerateImageFeaturePrintRequest_FromPointer(ptr unsafe.Pointer) VNGenerateImageFeaturePrintRequest {
	return VNGenerateImageFeaturePrintRequest{gen_VNGenerateImageFeaturePrintRequest{
		VNImageBasedRequest_FromPointer(ptr),
	}}
}

func VNGenerateImageFeaturePrintRequest_FromRef(ref objc.Ref) VNGenerateImageFeaturePrintRequest {
	return VNGenerateImageFeaturePrintRequest_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init initializes a new instance of the VNGenerateImageFeaturePrintRequest class.
func (x gen_VNGenerateImageFeaturePrintRequest) Init() VNGenerateImageFeaturePrintRequest {
	ret := C.VNGenerateImageFeaturePrintRequest_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return VNGenerateImageFeaturePrintRequest_FromPointer(ret)
}

// Init_AsVNGenerateImageFeaturePrintRequest is a typed version of Init.
func (x gen_VNGenerateImageFeaturePrintRequest) Init_AsVNGenerateImageFeaturePrintRequest() VNGenerateImageFeaturePrintRequest {
	ret := C.VNGenerateImageFeaturePrintRequest_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return VNGenerateImageFeaturePrintRequest_FromPointer(ret)
}

// Results returns the results of the feature print request.
//
// See https://developer.apple.com/documentation/vision/vngenerateimagefeatureprintrequest/3750984-results?language=objc for details.
func (x gen_VNGenerateImageFeaturePrintRequest) Results() core.NSArray {
	ret := C.VNGenerateImageFeaturePrintRequest_inst_Results(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

type VNImageBasedRequestRef interface {
	Pointer() uintptr
	Init_AsVNImageBasedRequest() VNImageBasedRequest
}

type gen_VNImageBasedRequest struct {
	VNRequest
}

func VNImageBasedRequest_FromPointer(ptr unsafe.Pointer) VNImageBasedRequest {
	return VNImageBasedRequest{gen_VNImageBasedRequest{
		VNRequest_FromPointer(ptr),
	}}
}

func VNImageBasedRequest_FromRef(ref objc.Ref) VNImageBasedRequest {
	return VNImageBasedRequest_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init initializes a new instance of the VNImageBasedRequest class.
func (x gen_VNImageBasedRequest) Init() VNImageBasedRequest {
	ret := C.VNImageBasedRequest_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return VNImageBasedRequest_FromPointer(ret)
}

// Init_AsVNImageBasedRequest is a typed version of Init.
func (x gen_VNImageBasedRequest) Init_AsVNImageBasedRequest() VNImageBasedRequest {
	ret := C.VNImageBasedRequest_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return VNImageBasedRequest_FromPointer(ret)
}

type VNImageRequestHandlerRef interface {
	Pointer() uintptr
	Init_AsVNImageRequestHandler() VNImageRequestHandler
}

type gen_VNImageRequestHandler struct {
	objc.Object
}

func VNImageRequestHandler_FromPointer(ptr unsafe.Pointer) VNImageRequestHandler {
	return VNImageRequestHandler{gen_VNImageRequestHandler{
		objc.Object_FromPointer(ptr),
	}}
}

func VNImageRequestHandler_FromRef(ref objc.Ref) VNImageRequestHandler {
	return VNImageRequestHandler_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init initializes a new instance of the VNImageRequestHandler class.
func (x gen_VNImageRequestHandler) Init() VNImageRequestHandler {
	panic("Init is unavailable")
}

// Init_AsVNImageRequestHandler is a typed version of Init.
func (x gen_VNImageRequestHandler) Init_AsVNImageRequestHandler() VNImageRequestHandler {
	panic("Init_AsVNImageRequestHandler is unavailable")
}

// InitWithDataOptions creates a handler to be used for performing requests on an image contained in an NSData object.
//
// See https://developer.apple.com/documentation/vision/vnimagerequesthandler/2866551-initwithdata?language=objc for details.
func (x gen_VNImageRequestHandler) InitWithDataOptions(
	imageData core.NSDataRef,
	options core.NSDictionaryRef,
) VNImageRequestHandler {
	ret := C.VNImageRequestHandler_inst_InitWithDataOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(imageData),
		objc.RefPointer(options),
	)

	return VNImageRequestHandler_FromPointer(ret)
}

// InitWithURLOptions creates a handler to be used for performing requests on an image at the specified URL.
//
// See https://developer.apple.com/documentation/vision/vnimagerequesthandler/2866553-initwithurl?language=objc for details.
func (x gen_VNImageRequestHandler) InitWithURLOptions(
	imageURL core.NSURLRef,
	options core.NSDictionaryRef,
) VNImageRequestHandler {
	ret := C.VNImageRequestHandler_inst_InitWithURLOptions(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(imageURL),
		objc.RefPointer(options),
	)

	return VNImageRequestHandler_FromPointer(ret)
}

// PerformRequestsError schedules Vision requests to be performed.
//
// See https://developer.apple.com/documentation/vision/vnimagerequesthandler/2880297-performrequests?language=objc for details.
func (x gen_VNImageRequestHandler) PerformRequestsError(
	requests core.NSArrayRef,
	error core.NSErrorRef,
) bool {
	ret := C.VNImageRequestHandler_inst_PerformRequestsError(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(requests),
		objc.RefPointer(error),
	)

	return convertObjCBoolToGo(ret)
}

type VNObservationRef interface {
	Pointer() uintptr
	Init_AsVNObservation() VNObservation
}

type gen_VNObservation struct {
	objc.Object
}

func VNObservation_FromPointer(ptr unsafe.Pointer) VNObservation {
	return VNObservation{gen_VNObservation{
		objc.Object_FromPointer(ptr),
	}}
}

func VNObservation_FromRef(ref objc.Ref) VNObservation {
	return VNObservation_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init initializes a new instance of the VNObservation class.
func (x gen_VNObservation) Init() VNObservation {
	ret := C.VNObservation_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return VNObservation_FromPointer(ret)
}

// Init_AsVNObservation is a typed version of Init.
func (x gen_VNObservation) Init_AsVNObservation() VNObservation {
	ret := C.VNObservation_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return VNObservation_FromPointer(ret)
}

type VNRecognizedTextObservationRef interface {
	Pointer() uintptr
	Init_AsVNRecognizedTextObservation() VNRecognizedTextObservation
}

type gen_VNRecognizedTextObservation struct {
	objc.Object
}

func VNRecognizedTextObservation_FromPointer(ptr unsafe.Pointer) VNRecognizedTextObservation {
	return VNRecognizedTextObservation{gen_VNRecognizedTextObservation{
		objc.Object_FromPointer(ptr),
	}}
}

func VNRecognizedTextObservation_FromRef(ref objc.Ref) VNRecognizedTextObservation {
	return VNRecognizedTextObservation_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// TopCandidates requests the n top candidates for a recognized text string.
//
// See https://developer.apple.com/documentation/vision/vnrecognizedtextobservation/3152637-topcandidates?language=objc for details.
func (x gen_VNRecognizedTextObservation) TopCandidates(
	maxCandidateCount core.NSUInteger,
) core.NSArray {
	ret := C.VNRecognizedTextObservation_inst_TopCandidates(
		unsafe.Pointer(x.Pointer()),
		C.ulong(maxCandidateCount),
	)

	return core.NSArray_FromPointer(ret)
}

// Init initializes a new instance of the VNRecognizedTextObservation class.
func (x gen_VNRecognizedTextObservation) Init() VNRecognizedTextObservation {
	ret := C.VNRecognizedTextObservation_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return VNRecognizedTextObservation_FromPointer(ret)
}

// Init_AsVNRecognizedTextObservation is a typed version of Init.
func (x gen_VNRecognizedTextObservation) Init_AsVNRecognizedTextObservation() VNRecognizedTextObservation {
	ret := C.VNRecognizedTextObservation_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return VNRecognizedTextObservation_FromPointer(ret)
}

type VNRecognizeTextRequestRef interface {
	Pointer() uintptr
	Init_AsVNRecognizeTextRequest() VNRecognizeTextRequest
}

type gen_VNRecognizeTextRequest struct {
	VNImageBasedRequest
}

func VNRecognizeTextRequest_FromPointer(ptr unsafe.Pointer) VNRecognizeTextRequest {
	return VNRecognizeTextRequest{gen_VNRecognizeTextRequest{
		VNImageBasedRequest_FromPointer(ptr),
	}}
}

func VNRecognizeTextRequest_FromRef(ref objc.Ref) VNRecognizeTextRequest {
	return VNRecognizeTextRequest_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// SupportedRecognitionLanguagesAndReturnError returns the identifiers of the languages that the request supports.
//
// See https://developer.apple.com/documentation/vision/vnrecognizetextrequest/3751006-supportedrecognitionlanguagesand?language=objc for details.
func (x gen_VNRecognizeTextRequest) SupportedRecognitionLanguagesAndReturnError(
	error core.NSErrorRef,
) core.NSArray {
	ret := C.VNRecognizeTextRequest_inst_SupportedRecognitionLanguagesAndReturnError(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(error),
	)

	return core.NSArray_FromPointer(ret)
}

// Init initializes a new instance of the VNRecognizeTextRequest class.
func (x gen_VNRecognizeTextRequest) Init() VNRecognizeTextRequest {
	ret := C.VNRecognizeTextRequest_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return VNRecognizeTextRequest_FromPointer(ret)
}

// Init_AsVNRecognizeTextRequest is a typed version of Init.
func (x gen_VNRecognizeTextRequest) Init_AsVNRecognizeTextRequest() VNRecognizeTextRequest {
	ret := C.VNRecognizeTextRequest_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return VNRecognizeTextRequest_FromPointer(ret)
}

// Results returns the results of the text recognition request.
//
// See https://developer.apple.com/documentation/vision/vnrecognizetextrequest/3751005-results?language=objc for details.
func (x gen_VNRecognizeTextRequest) Results() core.NSArray {
	ret := C.VNRecognizeTextRequest_inst_Results(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}

type VNRequestRef interface {
	Pointer() uintptr
	Init_AsVNRequest() VNRequest
}

type gen_VNRequest struct {
	objc.Object
}

func VNRequest_FromPointer(ptr unsafe.Pointer) VNRequest {
	return VNRequest{gen_VNRequest{
		objc.Object_FromPointer(ptr),
	}}
}

func VNRequest_FromRef(ref objc.Ref) VNRequest {
	return VNRequest_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Cancel cancels the request before it can finish executing.
//
// See https://developer.apple.com/documentation/vision/vnrequest/2867234-cancel?language=objc for details.
func (x gen_VNRequest) Cancel() {
	C.VNRequest_inst_Cancel(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// Init creates a new Vision request with no completion handler.
//
// See https://developer.apple.com/documentation/vision/vnrequest/2875423-init?language=objc for details.
func (x gen_VNRequest) Init() VNRequest {
	ret := C.VNRequest_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return VNRequest_FromPointer(ret)
}

// Init_AsVNRequest is a typed version of Init.
//
// See https://developer.apple.com/documentation/vision/vnrequest/2875423-init?language=objc for details.
func (x gen_VNRequest) Init_AsVNRequest() VNRequest {
	ret := C.VNRequest_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return VNRequest_FromPointer(ret)
}

// SupportedComputeStageDevicesAndReturnError returns the collection of compute devices per-stage that a request supports.
//
// See https://developer.apple.com/documentation/vision/vnrequest/4173243-supportedcomputestagedevicesandr?language=objc for details.
func (x gen_VNRequest) SupportedComputeStageDevicesAndReturnError(
	error core.NSErrorRef,
) core.NSDictionary {
	ret := C.VNRequest_inst_SupportedComputeStageDevicesAndReturnError(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(error),
	)

	return core.NSDictionary_FromPointer(ret)
}

// Results returns the collection of VNObservation results generated by request processing.
//
// See https://developer.apple.com/documentation/vision/vnrequest/2867238-results?language=objc for details.
func (x gen_VNRequest) Results() core.NSArray {
	ret := C.VNRequest_inst_Results(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSArray_FromPointer(ret)
}
