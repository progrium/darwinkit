// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ProcessInfo] class.
var ProcessInfoClass = _ProcessInfoClass{objc.GetClass("NSProcessInfo")}

type _ProcessInfoClass struct {
	objc.Class
}

// An interface definition for the [ProcessInfo] class.
type IProcessInfo interface {
	objc.IObject
	EndActivity(activity objc.IObject)
	BeginActivityWithOptionsReason(options ActivityOptions, reason string) objc.Object
	DisableAutomaticTermination(reason string)
	IsOperatingSystemAtLeastVersion(version OperatingSystemVersion) bool
	EnableAutomaticTermination(reason string)
	PerformActivityWithOptionsReasonUsingBlock(options ActivityOptions, reason string, block func())
	EnableSuddenTermination()
	DisableSuddenTermination()
	AutomaticTerminationSupportEnabled() bool
	SetAutomaticTerminationSupportEnabled(value bool)
	Arguments() []string
	OperatingSystemVersionString() string
	ThermalState() ProcessInfoThermalState
	IsMacCatalystApp() bool
	SystemUptime() TimeInterval
	IsiOSAppOnMac() bool
	ActiveProcessorCount() uint
	ProcessName() string
	SetProcessName(value string)
	OperatingSystemVersion() OperatingSystemVersion
	UserName() string
	IsLowPowerModeEnabled() bool
	PhysicalMemory() int64
	FullUserName() string
	GloballyUniqueString() string
	HostName() string
	ProcessorCount() uint
	ProcessIdentifier() int
	Environment() map[string]string
}

// A collection of information about the current process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo?language=objc
type ProcessInfo struct {
	objc.Object
}

func ProcessInfoFrom(ptr unsafe.Pointer) ProcessInfo {
	return ProcessInfo{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _ProcessInfoClass) Alloc() ProcessInfo {
	rv := objc.Call[ProcessInfo](pc, objc.Sel("alloc"))
	return rv
}

func ProcessInfo_Alloc() ProcessInfo {
	return ProcessInfoClass.Alloc()
}

func (pc _ProcessInfoClass) New() ProcessInfo {
	rv := objc.Call[ProcessInfo](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewProcessInfo() ProcessInfo {
	return ProcessInfoClass.New()
}

func (p_ ProcessInfo) Init() ProcessInfo {
	rv := objc.Call[ProcessInfo](p_, objc.Sel("init"))
	return rv
}

// Ends the given activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1411321-endactivity?language=objc
func (p_ ProcessInfo) EndActivity(activity objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("endActivity:"), activity)
}

// Begin an activity using the given options and reason. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1415995-beginactivitywithoptions?language=objc
func (p_ ProcessInfo) BeginActivityWithOptionsReason(options ActivityOptions, reason string) objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("beginActivityWithOptions:reason:"), options, reason)
	return rv
}

// Disables automatic termination for the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1417402-disableautomatictermination?language=objc
func (p_ ProcessInfo) DisableAutomaticTermination(reason string) {
	objc.Call[objc.Void](p_, objc.Sel("disableAutomaticTermination:"), reason)
}

// Returns a Boolean value indicating whether the version of the operating system on which the process is executing is the same or later than the given version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1414876-isoperatingsystematleastversion?language=objc
func (p_ ProcessInfo) IsOperatingSystemAtLeastVersion(version OperatingSystemVersion) bool {
	rv := objc.Call[bool](p_, objc.Sel("isOperatingSystemAtLeastVersion:"), version)
	return rv
}

// Enables automatic termination for the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1409422-enableautomatictermination?language=objc
func (p_ ProcessInfo) EnableAutomaticTermination(reason string) {
	objc.Call[objc.Void](p_, objc.Sel("enableAutomaticTermination:"), reason)
}

// Synchronously perform an activity defined by a given block using the given options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1418048-performactivitywithoptions?language=objc
func (p_ ProcessInfo) PerformActivityWithOptionsReasonUsingBlock(options ActivityOptions, reason string, block func()) {
	objc.Call[objc.Void](p_, objc.Sel("performActivityWithOptions:reason:usingBlock:"), options, reason, block)
}

// Enables the application for quick killing using sudden termination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1409836-enablesuddentermination?language=objc
func (p_ ProcessInfo) EnableSuddenTermination() {
	objc.Call[objc.Void](p_, objc.Sel("enableSuddenTermination"))
}

// Disables the application for quickly killing using sudden termination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1412841-disablesuddentermination?language=objc
func (p_ ProcessInfo) DisableSuddenTermination() {
	objc.Call[objc.Void](p_, objc.Sel("disableSuddenTermination"))
}

// A Boolean value indicating whether the app supports automatic termination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1407578-automaticterminationsupportenabl?language=objc
func (p_ ProcessInfo) AutomaticTerminationSupportEnabled() bool {
	rv := objc.Call[bool](p_, objc.Sel("automaticTerminationSupportEnabled"))
	return rv
}

// A Boolean value indicating whether the app supports automatic termination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1407578-automaticterminationsupportenabl?language=objc
func (p_ ProcessInfo) SetAutomaticTerminationSupportEnabled(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setAutomaticTerminationSupportEnabled:"), value)
}

// Array of strings with the command-line arguments for the process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1415596-arguments?language=objc
func (p_ ProcessInfo) Arguments() []string {
	rv := objc.Call[[]string](p_, objc.Sel("arguments"))
	return rv
}

// A string containing the version of the operating system on which the process is executing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1408730-operatingsystemversionstring?language=objc
func (p_ ProcessInfo) OperatingSystemVersionString() string {
	rv := objc.Call[string](p_, objc.Sel("operatingSystemVersionString"))
	return rv
}

// The current thermal state of the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1417480-thermalstate?language=objc
func (p_ ProcessInfo) ThermalState() ProcessInfoThermalState {
	rv := objc.Call[ProcessInfoThermalState](p_, objc.Sel("thermalState"))
	return rv
}

// A Boolean value that indicates whether the process originated as an iOS app and runs on macOS. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/3362531-maccatalystapp?language=objc
func (p_ ProcessInfo) IsMacCatalystApp() bool {
	rv := objc.Call[bool](p_, objc.Sel("isMacCatalystApp"))
	return rv
}

// The amount of time the system has been awake since the last time it was restarted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1414553-systemuptime?language=objc
func (p_ ProcessInfo) SystemUptime() TimeInterval {
	rv := objc.Call[TimeInterval](p_, objc.Sel("systemUptime"))
	return rv
}

// A Boolean value that indicates whether the process is an iPhone or iPad app running on a Mac. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/3608556-iosapponmac?language=objc
func (p_ ProcessInfo) IsiOSAppOnMac() bool {
	rv := objc.Call[bool](p_, objc.Sel("isiOSAppOnMac"))
	return rv
}

// The number of active processing cores available on the computer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1408184-activeprocessorcount?language=objc
func (p_ ProcessInfo) ActiveProcessorCount() uint {
	rv := objc.Call[uint](p_, objc.Sel("activeProcessorCount"))
	return rv
}

// The name of the process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1416428-processname?language=objc
func (p_ ProcessInfo) ProcessName() string {
	rv := objc.Call[string](p_, objc.Sel("processName"))
	return rv
}

// The name of the process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1416428-processname?language=objc
func (p_ ProcessInfo) SetProcessName(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setProcessName:"), value)
}

// The version of the operating system on which the process is executing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1410906-operatingsystemversion?language=objc
func (p_ ProcessInfo) OperatingSystemVersion() OperatingSystemVersion {
	rv := objc.Call[OperatingSystemVersion](p_, objc.Sel("operatingSystemVersion"))
	return rv
}

// Returns the process information agent for the process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1408734-processinfo?language=objc
func (pc _ProcessInfoClass) ProcessInfo() ProcessInfo {
	rv := objc.Call[ProcessInfo](pc, objc.Sel("processInfo"))
	return rv
}

// Returns the process information agent for the process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1408734-processinfo?language=objc
func ProcessInfo_ProcessInfo() ProcessInfo {
	return ProcessInfoClass.ProcessInfo()
}

// Returns the account name of the current user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1643193-username?language=objc
func (p_ ProcessInfo) UserName() string {
	rv := objc.Call[string](p_, objc.Sel("userName"))
	return rv
}

// A Boolean value that indicates the current state of Low Power Mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1617047-lowpowermodeenabled?language=objc
func (p_ ProcessInfo) IsLowPowerModeEnabled() bool {
	rv := objc.Call[bool](p_, objc.Sel("isLowPowerModeEnabled"))
	return rv
}

// The amount of physical memory on the computer in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1408211-physicalmemory?language=objc
func (p_ ProcessInfo) PhysicalMemory() int64 {
	rv := objc.Call[int64](p_, objc.Sel("physicalMemory"))
	return rv
}

// Returns the full name of the current user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1643199-fullusername?language=objc
func (p_ ProcessInfo) FullUserName() string {
	rv := objc.Call[string](p_, objc.Sel("fullUserName"))
	return rv
}

// Global unique identifier for the process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1416432-globallyuniquestring?language=objc
func (p_ ProcessInfo) GloballyUniqueString() string {
	rv := objc.Call[string](p_, objc.Sel("globallyUniqueString"))
	return rv
}

// The name of the host computer on which the process is executing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1417236-hostname?language=objc
func (p_ ProcessInfo) HostName() string {
	rv := objc.Call[string](p_, objc.Sel("hostName"))
	return rv
}

// The number of processing cores available on the computer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1415622-processorcount?language=objc
func (p_ ProcessInfo) ProcessorCount() uint {
	rv := objc.Call[uint](p_, objc.Sel("processorCount"))
	return rv
}

// The identifier of the process (often called process ID). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1415929-processidentifier?language=objc
func (p_ ProcessInfo) ProcessIdentifier() int {
	rv := objc.Call[int](p_, objc.Sel("processIdentifier"))
	return rv
}

// The variable names (keys) and their values in the environment from which the process was launched. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprocessinfo/1417911-environment?language=objc
func (p_ ProcessInfo) Environment() map[string]string {
	rv := objc.Call[map[string]string](p_, objc.Sel("environment"))
	return rv
}
