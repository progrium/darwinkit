package cocoa

const (
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
