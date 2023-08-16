// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RunningApplication] class.
var RunningApplicationClass = _RunningApplicationClass{objc.GetClass("NSRunningApplication")}

type _RunningApplicationClass struct {
	objc.Class
}

// An interface definition for the [RunningApplication] class.
type IRunningApplication interface {
	objc.IObject
	Hide() bool
	ForceTerminate() bool
	ActivateWithOptions(options ApplicationActivationOptions) bool
	Unhide() bool
	Terminate() bool
	BundleIdentifier() string
	ExecutableURL() foundation.URL
	IsHidden() bool
	OwnsMenuBar() bool
	IsTerminated() bool
	LaunchDate() foundation.Date
	IsActive() bool
	ExecutableArchitecture() int
	ActivationPolicy() ApplicationActivationPolicy
	LocalizedName() string
	Icon() Image
	BundleURL() foundation.URL
	IsFinishedLaunching() bool
}

// An object that can manipulate and provide information for a single instance of an app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication?language=objc
type RunningApplication struct {
	objc.Object
}

func RunningApplicationFrom(ptr unsafe.Pointer) RunningApplication {
	return RunningApplication{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RunningApplicationClass) Alloc() RunningApplication {
	rv := objc.Call[RunningApplication](rc, objc.Sel("alloc"))
	return rv
}

func RunningApplication_Alloc() RunningApplication {
	return RunningApplicationClass.Alloc()
}

func (rc _RunningApplicationClass) New() RunningApplication {
	rv := objc.Call[RunningApplication](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRunningApplication() RunningApplication {
	return RunningApplicationClass.New()
}

func (r_ RunningApplication) Init() RunningApplication {
	rv := objc.Call[RunningApplication](r_, objc.Sel("init"))
	return rv
}

// Attempts to hide or the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1526608-hide?language=objc
func (r_ RunningApplication) Hide() bool {
	rv := objc.Call[bool](r_, objc.Sel("hide"))
	return rv
}

// Returns an array of currently running applications with the specified bundle identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1530798-runningapplicationswithbundleide?language=objc
func (rc _RunningApplicationClass) RunningApplicationsWithBundleIdentifier(bundleIdentifier string) []RunningApplication {
	rv := objc.Call[[]RunningApplication](rc, objc.Sel("runningApplicationsWithBundleIdentifier:"), bundleIdentifier)
	return rv
}

// Returns an array of currently running applications with the specified bundle identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1530798-runningapplicationswithbundleide?language=objc
func RunningApplication_RunningApplicationsWithBundleIdentifier(bundleIdentifier string) []RunningApplication {
	return RunningApplicationClass.RunningApplicationsWithBundleIdentifier(bundleIdentifier)
}

// Terminates invisibly running applications as if triggered by system memory pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1529538-terminateautomaticallyterminable?language=objc
func (rc _RunningApplicationClass) TerminateAutomaticallyTerminableApplications() {
	objc.Call[objc.Void](rc, objc.Sel("terminateAutomaticallyTerminableApplications"))
}

// Terminates invisibly running applications as if triggered by system memory pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1529538-terminateautomaticallyterminable?language=objc
func RunningApplication_TerminateAutomaticallyTerminableApplications() {
	RunningApplicationClass.TerminateAutomaticallyTerminableApplications()
}

// Attempts to force the receiver to quit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1530370-forceterminate?language=objc
func (r_ RunningApplication) ForceTerminate() bool {
	rv := objc.Call[bool](r_, objc.Sel("forceTerminate"))
	return rv
}

// Attempts to activate the application using the specified options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1528725-activatewithoptions?language=objc
func (r_ RunningApplication) ActivateWithOptions(options ApplicationActivationOptions) bool {
	rv := objc.Call[bool](r_, objc.Sel("activateWithOptions:"), options)
	return rv
}

// Attempts to unhide or the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1534676-unhide?language=objc
func (r_ RunningApplication) Unhide() bool {
	rv := objc.Call[bool](r_, objc.Sel("unhide"))
	return rv
}

// Attempts to quit the receiver normally. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1528922-terminate?language=objc
func (r_ RunningApplication) Terminate() bool {
	rv := objc.Call[bool](r_, objc.Sel("terminate"))
	return rv
}

// Indicates the CFBundleIdentifier of the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1529140-bundleidentifier?language=objc
func (r_ RunningApplication) BundleIdentifier() string {
	rv := objc.Call[string](r_, objc.Sel("bundleIdentifier"))
	return rv
}

// Indicates the URL to the application's executable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1531062-executableurl?language=objc
func (r_ RunningApplication) ExecutableURL() foundation.URL {
	rv := objc.Call[foundation.URL](r_, objc.Sel("executableURL"))
	return rv
}

// Indicates whether the application is currently hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1525949-hidden?language=objc
func (r_ RunningApplication) IsHidden() bool {
	rv := objc.Call[bool](r_, objc.Sel("isHidden"))
	return rv
}

// Returns whether the application owns the current menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1525915-ownsmenubar?language=objc
func (r_ RunningApplication) OwnsMenuBar() bool {
	rv := objc.Call[bool](r_, objc.Sel("ownsMenuBar"))
	return rv
}

// Indicates that the receiver’s application has terminated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1532239-terminated?language=objc
func (r_ RunningApplication) IsTerminated() bool {
	rv := objc.Call[bool](r_, objc.Sel("isTerminated"))
	return rv
}

// Indicates the date when the application was launched. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1532595-launchdate?language=objc
func (r_ RunningApplication) LaunchDate() foundation.Date {
	rv := objc.Call[foundation.Date](r_, objc.Sel("launchDate"))
	return rv
}

// Returns an NSRunningApplication representing this application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1533604-currentapplication?language=objc
func (rc _RunningApplicationClass) CurrentApplication() RunningApplication {
	rv := objc.Call[RunningApplication](rc, objc.Sel("currentApplication"))
	return rv
}

// Returns an NSRunningApplication representing this application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1533604-currentapplication?language=objc
func RunningApplication_CurrentApplication() RunningApplication {
	return RunningApplicationClass.CurrentApplication()
}

// Indicates whether the application is currently frontmost. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1528778-active?language=objc
func (r_ RunningApplication) IsActive() bool {
	rv := objc.Call[bool](r_, objc.Sel("isActive"))
	return rv
}

// Indicates the executing processor architecture for the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1524287-executablearchitecture?language=objc
func (r_ RunningApplication) ExecutableArchitecture() int {
	rv := objc.Call[int](r_, objc.Sel("executableArchitecture"))
	return rv
}

// Indicates the activation policy of the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1533103-activationpolicy?language=objc
func (r_ RunningApplication) ActivationPolicy() ApplicationActivationPolicy {
	rv := objc.Call[ApplicationActivationPolicy](r_, objc.Sel("activationPolicy"))
	return rv
}

// Indicates the localized name of the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1526751-localizedname?language=objc
func (r_ RunningApplication) LocalizedName() string {
	rv := objc.Call[string](r_, objc.Sel("localizedName"))
	return rv
}

// Returns the icon for the receiver’s application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1529885-icon?language=objc
func (r_ RunningApplication) Icon() Image {
	rv := objc.Call[Image](r_, objc.Sel("icon"))
	return rv
}

// Indicates the URL to the application's bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1535500-bundleurl?language=objc
func (r_ RunningApplication) BundleURL() foundation.URL {
	rv := objc.Call[foundation.URL](r_, objc.Sel("bundleURL"))
	return rv
}

// Indicates whether the receiver’s process has finished launching, [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/1532002-finishedlaunching?language=objc
func (r_ RunningApplication) IsFinishedLaunching() bool {
	rv := objc.Call[bool](r_, objc.Sel("isFinishedLaunching"))
	return rv
}
