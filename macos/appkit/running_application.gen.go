// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var RunningApplicationClass = _RunningApplicationClass{objc.GetClass("NSRunningApplication")}

type _RunningApplicationClass struct {
	objc.Class
}

type IRunningApplication interface {
	objc.IObject
	ActivateWithOptions(options ApplicationActivationOptions) bool
	Hide() bool
	Unhide() bool
	ForceTerminate() bool
	Terminate() bool
	IsActive() bool
	ActivationPolicy() ApplicationActivationPolicy
	IsHidden() bool
	LocalizedName() string
	Icon() Image
	BundleIdentifier() string
	BundleURL() foundation.URL
	ExecutableArchitecture() int
	ExecutableURL() foundation.URL
	LaunchDate() foundation.Date
	IsFinishedLaunching() bool
	OwnsMenuBar() bool
	IsTerminated() bool
}

type RunningApplication struct {
	objc.Object
}

func MakeRunningApplication(ptr unsafe.Pointer) RunningApplication {
	return RunningApplication{
		Object: objc.MakeObject(ptr),
	}
}

func (rc _RunningApplicationClass) Alloc() RunningApplication {
	rv := objc.CallMethod[RunningApplication](rc, objc.GetSelector("alloc"))
	return rv
}

func RunningApplication_Alloc() RunningApplication {
	return RunningApplicationClass.Alloc()
}

func (rc _RunningApplicationClass) New() RunningApplication {
	rv := objc.CallMethod[RunningApplication](rc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewRunningApplication() RunningApplication {
	return RunningApplicationClass.New()
}

func RunningApplication_New() RunningApplication {
	return RunningApplicationClass.New()
}

func (r_ RunningApplication) Init() RunningApplication {
	rv := objc.CallMethod[RunningApplication](r_, objc.GetSelector("init"))
	return rv
}

func RunningApplication_Init() RunningApplication {
	return RunningApplicationClass.Alloc().Init()
}

func (rc _RunningApplicationClass) RunningApplicationsWithBundleIdentifier(bundleIdentifier string) []RunningApplication {
	rv := objc.CallMethod[[]RunningApplication](rc, objc.GetSelector("runningApplicationsWithBundleIdentifier:"), bundleIdentifier)
	return rv
}

func RunningApplication_RunningApplicationsWithBundleIdentifier(bundleIdentifier string) []RunningApplication {
	return RunningApplicationClass.RunningApplicationsWithBundleIdentifier(bundleIdentifier)
}

func (r_ RunningApplication) ActivateWithOptions(options ApplicationActivationOptions) bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("activateWithOptions:"), options)
	return rv
}

func (r_ RunningApplication) Hide() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("hide"))
	return rv
}

func (r_ RunningApplication) Unhide() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("unhide"))
	return rv
}

func (r_ RunningApplication) ForceTerminate() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("forceTerminate"))
	return rv
}

func (r_ RunningApplication) Terminate() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("terminate"))
	return rv
}

func (rc _RunningApplicationClass) TerminateAutomaticallyTerminableApplications() {
	objc.CallMethod[objc.Void](rc, objc.GetSelector("terminateAutomaticallyTerminableApplications"))
}

func RunningApplication_TerminateAutomaticallyTerminableApplications() {
	RunningApplicationClass.TerminateAutomaticallyTerminableApplications()
}

func (rc _RunningApplicationClass) CurrentApplication() RunningApplication {
	rv := objc.CallMethod[RunningApplication](rc, objc.GetSelector("currentApplication"))
	return rv
}

func RunningApplication_CurrentApplication() RunningApplication {
	return RunningApplicationClass.CurrentApplication()
}

func (r_ RunningApplication) IsActive() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("isActive"))
	return rv
}

func (r_ RunningApplication) ActivationPolicy() ApplicationActivationPolicy {
	rv := objc.CallMethod[ApplicationActivationPolicy](r_, objc.GetSelector("activationPolicy"))
	return rv
}

func (r_ RunningApplication) IsHidden() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("isHidden"))
	return rv
}

func (r_ RunningApplication) LocalizedName() string {
	rv := objc.CallMethod[string](r_, objc.GetSelector("localizedName"))
	return rv
}

func (r_ RunningApplication) Icon() Image {
	rv := objc.CallMethod[Image](r_, objc.GetSelector("icon"))
	return rv
}

func (r_ RunningApplication) BundleIdentifier() string {
	rv := objc.CallMethod[string](r_, objc.GetSelector("bundleIdentifier"))
	return rv
}

func (r_ RunningApplication) BundleURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](r_, objc.GetSelector("bundleURL"))
	return rv
}

func (r_ RunningApplication) ExecutableArchitecture() int {
	rv := objc.CallMethod[int](r_, objc.GetSelector("executableArchitecture"))
	return rv
}

func (r_ RunningApplication) ExecutableURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](r_, objc.GetSelector("executableURL"))
	return rv
}

func (r_ RunningApplication) LaunchDate() foundation.Date {
	rv := objc.CallMethod[foundation.Date](r_, objc.GetSelector("launchDate"))
	return rv
}

func (r_ RunningApplication) IsFinishedLaunching() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("isFinishedLaunching"))
	return rv
}

func (r_ RunningApplication) OwnsMenuBar() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("ownsMenuBar"))
	return rv
}

func (r_ RunningApplication) IsTerminated() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("isTerminated"))
	return rv
}
