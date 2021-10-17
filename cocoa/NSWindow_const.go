package cocoa

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -lobjc -framework AppKit
#include <AppKit/AppKit.h>
*/
import "C"

type NSBackingStoreType uintptr

const (
	NSBorderlessWindowMask         = 0
	NSTitledWindowMask             = 1 << 0
	NSClosableWindowMask           = 1 << 1
	NSMiniaturizableWindowMask     = 1 << 2
	NSResizableWindowMask          = 1 << 3
	NSTexturedBackgroundWindowMask = 1 << 8
	NSWindowStyleMaskFullScreen    = 1 << 14

	NSFullSizeContentViewWindowMask = 32768

	NSWindowTitleVisible = 0
	NSWindowTitleHidden  = 1

	NSWindowAbove = 1
	NSWindowBelow = -1
	NSWindowOut   = 0

	NSBackingStoreRetained    NSBackingStoreType = 0
	NSBackingStoreNonretained NSBackingStoreType = 1
	NSBackingStoreBuffered    NSBackingStoreType = 2

	NSNormalWindowLevel      = C.NSNormalWindowLevel
	NSFloatingWindowLevel    = C.NSFloatingWindowLevel
	NSSubmenuWindowLevel     = C.NSSubmenuWindowLevel
	NSTornOffMenuWindowLevel = C.NSTornOffMenuWindowLevel
	NSMainMenuWindowLevel    = C.NSMainMenuWindowLevel
	NSStatusWindowLevel      = C.NSStatusWindowLevel
	NSModalPanelWindowLevel  = C.NSModalPanelWindowLevel
	NSPopUpMenuWindowLevel   = C.NSPopUpMenuWindowLevel
	NSScreenSaverWindowLevel = C.NSScreenSaverWindowLevel

	NSWindowCollectionBehaviorDefault                   = 0
	NSWindowCollectionBehaviorCanJoinAllSpaces          = 1 << 0
	NSWindowCollectionBehaviorMoveToActiveSpace         = 1 << 1
	NSWindowCollectionBehaviorManaged                   = 1 << 2
	NSWindowCollectionBehaviorTransient                 = 1 << 3
	NSWindowCollectionBehaviorStationary                = 1 << 4
	NSWindowCollectionBehaviorParticipatesInCycle       = 1 << 5
	NSWindowCollectionBehaviorIgnoresCycle              = 1 << 6
	NSWindowCollectionBehaviorFullScreenPrimary         = 1 << 7
	NSWindowCollectionBehaviorFullScreenAuxiliary       = 1 << 8
	NSWindowCollectionBehaviorFullScreenNone            = 1 << 9
	NSWindowCollectionBehaviorFullScreenAllowsTiling    = 1 << 11
	NSWindowCollectionBehaviorFullScreenDisallowsTiling = 1 << 12
)
