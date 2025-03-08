package coreservices

import (
	"github.com/progrium/darwinkit/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// LaunchServices provides interfaces for launching applications, opening files, etc.
type LaunchServices struct {
	objc.Object
}

// LSSharedFileListRef is a reference to a shared file list
type LSSharedFileListRef objc.Object

// LSSharedFileListItemRef is a reference to an item in a shared file list
type LSSharedFileListItemRef objc.Object

// LSOpenApplication opens an application with the specified options
func LSOpenApplication(appParams map[string]interface{}, outPSN *ProcessSerialNumber) OSStatus {
	// Create an application parameters dictionary
	dict := foundation.Dictionary_Dictionary()
	for key, value := range appParams {
		if strValue, ok := value.(string); ok {
			dict.SetObjectForKey(foundation.String_StringWithString(strValue), foundation.String_StringWithString(key))
		}
	}
	return OSStatus(int(objc.Call[int](objc.GetClass("LSOpenApplication"), objc.Sel("LSOpenApplication:outPSN:"), dict, outPSN)))
}

// LSOpenFromURLSpec opens a URL with the specified options
func LSOpenFromURLSpec(inURLSpec map[string]interface{}, outLaunchedURL *foundation.URL) OSStatus {
	// Create a URL specification dictionary
	dict := foundation.Dictionary_Dictionary()
	for key, value := range inURLSpec {
		if strValue, ok := value.(string); ok {
			dict.SetObjectForKey(foundation.String_StringWithString(strValue), foundation.String_StringWithString(key))
		} else if urlValue, ok := value.(foundation.URL); ok {
			dict.SetObjectForKey(urlValue, foundation.String_StringWithString(key))
		}
	}
	return OSStatus(int(objc.Call[int](objc.GetClass("LSOpenFromURLSpec"), objc.Sel("LSOpenFromURLSpec:outLaunchedURL:"), dict, outLaunchedURL)))
}

// LSOpenFromURL opens a URL
func LSOpenFromURL(inURL foundation.URL, outURL *foundation.URL) OSStatus {
	return OSStatus(int(objc.Call[int](objc.GetClass("LSOpenFromURL"), objc.Sel("LSOpenFromURL:outURL:"), inURL, outURL)))
}

// LSCopyApplicationURLsForURL gets application URLs for a given document URL
func LSCopyApplicationURLsForURL(inURL foundation.URL, inRole foundation.String) foundation.Array {
	return objc.Call[foundation.Array](objc.GetClass("LSCopyApplicationURLsForURL"), objc.Sel("LSCopyApplicationURLsForURL:inRole:"), inURL, inRole)
}

// FSEvents provides interfaces for monitoring file system changes
type FSEvents struct {
	objc.Object
}

// FSEventStreamRef is a reference to a file system event stream
type FSEventStreamRef objc.Object

// FSEventStreamCreate creates a file system event stream
func FSEventStreamCreate(
	allocator foundation.Object,
	callback interface{},
	context *FSEventStreamContext,
	pathsToWatch foundation.Array,
	sinceWhen FSEventStreamEventId,
	latency float64,
	flags FSEventStreamCreateFlags) FSEventStreamRef {
	// This is a complex function that requires a callback conversion
	// For now, it's a simplified version
	return FSEventStreamRef(objc.Call[objc.Object](objc.GetClass("FSEventStreamCreate"), objc.Sel("FSEventStreamCreate:callback:context:pathsToWatch:sinceWhen:latency:flags:"), 
		allocator, callback, context, pathsToWatch, sinceWhen, latency, flags))
}

// FSEventStreamScheduleWithRunLoop schedules a file system event stream with a run loop
func FSEventStreamScheduleWithRunLoop(streamRef FSEventStreamRef, runLoop foundation.Object, runLoopMode foundation.String) {
	objc.Call[objc.Void](objc.GetClass("FSEventStreamScheduleWithRunLoop"), objc.Sel("FSEventStreamScheduleWithRunLoop:runLoop:runLoopMode:"), 
		streamRef, runLoop, runLoopMode)
}

// FSEventStreamStart starts a file system event stream
func FSEventStreamStart(streamRef FSEventStreamRef) bool {
	return objc.Call[bool](objc.GetClass("FSEventStreamStart"), objc.Sel("FSEventStreamStart:"), streamRef)
}

// FSEventStreamStop stops a file system event stream
func FSEventStreamStop(streamRef FSEventStreamRef) {
	objc.Call[objc.Void](objc.GetClass("FSEventStreamStop"), objc.Sel("FSEventStreamStop:"), streamRef)
}

// FSEventStreamInvalidate invalidates a file system event stream
func FSEventStreamInvalidate(streamRef FSEventStreamRef) {
	objc.Call[objc.Void](objc.GetClass("FSEventStreamInvalidate"), objc.Sel("FSEventStreamInvalidate:"), streamRef)
}

// FSEventStreamRelease releases a file system event stream
func FSEventStreamRelease(streamRef FSEventStreamRef) {
	objc.Call[objc.Void](objc.GetClass("FSEventStreamRelease"), objc.Sel("FSEventStreamRelease:"), streamRef)
}

// FSEventStreamContext represents the context for a file system event stream
type FSEventStreamContext struct {
	Version         int
	Info            objc.Object
	Retain          objc.Object
	Release         objc.Object
	CopyDescription objc.Object
}

// OSStatus represents a Core Services result code
type OSStatus int32

// ProcessSerialNumber identifies a process
type ProcessSerialNumber struct {
	HighLongOfPSN uint32
	LowLongOfPSN  uint32
}

// FSEventStreamEventId represents an event ID for a file system event
type FSEventStreamEventId uint64

// FSEventStreamCreateFlags represent flags for creating a file system event stream
type FSEventStreamCreateFlags uint32

// Various LaunchServices constants
const (
	LSLaunchDefaults             = 0x00000001
	LSLaunchAndPrint             = 0x00000002
	LSLaunchReserved2            = 0x00000004
	LSLaunchReserved3            = 0x00000008
	LSLaunchReserved4            = 0x00000010
	LSLaunchReserved5            = 0x00000020
	LSLaunchAndDisplayErrors     = 0x00000040
	LSLaunchInhibitBGOnly        = 0x00000080
	LSLaunchDontAddToRecents     = 0x00000100
	LSLaunchDontSwitch           = 0x00000200
	LSLaunchNoParams             = 0x00000800
	LSLaunchAsync                = 0x00010000
	LSLaunchStartClassic         = 0x00020000
	LSLaunchInClassic            = 0x00040000
	LSLaunchNewInstance          = 0x00080000
	LSLaunchAndHide              = 0x00100000
	LSLaunchAndHideOthers        = 0x00200000
	LSLaunchHasUntrustedContents = 0x00400000
)

// Various FSEvents constants
const (
	FSEventStreamCreateFlagNone              FSEventStreamCreateFlags = 0x00000000
	FSEventStreamCreateFlagUseCFTypes        FSEventStreamCreateFlags = 0x00000001
	FSEventStreamCreateFlagNoDefer           FSEventStreamCreateFlags = 0x00000002
	FSEventStreamCreateFlagWatchRoot         FSEventStreamCreateFlags = 0x00000004
	FSEventStreamCreateFlagIgnoreSelf        FSEventStreamCreateFlags = 0x00000008
	FSEventStreamCreateFlagFileEvents        FSEventStreamCreateFlags = 0x00000010
	FSEventStreamCreateFlagMarkSelf          FSEventStreamCreateFlags = 0x00000020
	FSEventStreamCreateFlagUseExtendedData   FSEventStreamCreateFlags = 0x00000040
	FSEventStreamCreateFlagFullHistory       FSEventStreamCreateFlags = 0x00000080
)

// FSEventStreamEventFlags represent flags for file system events
type FSEventStreamEventFlags uint32

const (
	FSEventStreamEventFlagNone              FSEventStreamEventFlags = 0x00000000
	FSEventStreamEventFlagMustScanSubDirs   FSEventStreamEventFlags = 0x00000001
	FSEventStreamEventFlagUserDropped       FSEventStreamEventFlags = 0x00000002
	FSEventStreamEventFlagKernelDropped     FSEventStreamEventFlags = 0x00000004
	FSEventStreamEventFlagEventIdsWrapped   FSEventStreamEventFlags = 0x00000008
	FSEventStreamEventFlagHistoryDone       FSEventStreamEventFlags = 0x00000010
	FSEventStreamEventFlagRootChanged       FSEventStreamEventFlags = 0x00000020
	FSEventStreamEventFlagMount             FSEventStreamEventFlags = 0x00000040
	FSEventStreamEventFlagUnmount           FSEventStreamEventFlags = 0x00000080
	FSEventStreamEventFlagItemCreated       FSEventStreamEventFlags = 0x00000100
	FSEventStreamEventFlagItemRemoved       FSEventStreamEventFlags = 0x00000200
	FSEventStreamEventFlagItemInodeMetaMod  FSEventStreamEventFlags = 0x00000400
	FSEventStreamEventFlagItemRenamed       FSEventStreamEventFlags = 0x00000800
	FSEventStreamEventFlagItemModified      FSEventStreamEventFlags = 0x00001000
	FSEventStreamEventFlagItemFinderInfoMod FSEventStreamEventFlags = 0x00002000
	FSEventStreamEventFlagItemChangeOwner   FSEventStreamEventFlags = 0x00004000
	FSEventStreamEventFlagItemXattrMod      FSEventStreamEventFlags = 0x00008000
	FSEventStreamEventFlagItemIsFile        FSEventStreamEventFlags = 0x00010000
	FSEventStreamEventFlagItemIsDir         FSEventStreamEventFlags = 0x00020000
	FSEventStreamEventFlagItemIsSymlink     FSEventStreamEventFlags = 0x00040000
	FSEventStreamEventFlagOwnEvent          FSEventStreamEventFlags = 0x00080000
	FSEventStreamEventFlagItemIsHardlink    FSEventStreamEventFlags = 0x00100000
	FSEventStreamEventFlagItemIsLastHardlink FSEventStreamEventFlags = 0x00200000
	FSEventStreamEventFlagItemCloned        FSEventStreamEventFlags = 0x00400000
)

// Special constants
const (
	KFSEventStreamEventIdSinceNow = FSEventStreamEventId(0xFFFFFFFFFFFFFFFF)
)

// NSURLBookmarkCreationOptions for bookmark creation
type NSURLBookmarkCreationOptions uint

const (
	NSURLBookmarkCreationSuitableForBookmarkFile      NSURLBookmarkCreationOptions = 1 << 10
	NSURLBookmarkCreationWithSecurityScope            NSURLBookmarkCreationOptions = 1 << 11
	NSURLBookmarkCreationSecurityScopeAllowOnlyReadAccess NSURLBookmarkCreationOptions = 1 << 12
)