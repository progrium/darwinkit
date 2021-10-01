package cocoa

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

	NSNormalWindowLevel      = 4
	NSFloatingWindowLevel    = 5
	NSSubmenuWindowLevel     = 6
	NSTornOffMenuWindowLevel = 6
	NSMainMenuWindowLevel    = 8
	NSStatusWindowLevel      = 9
	NSModalPanelWindowLevel  = 10
	NSPopUpMenuWindowLevel   = 11
	NSScreenSaverWindowLevel = 13

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
