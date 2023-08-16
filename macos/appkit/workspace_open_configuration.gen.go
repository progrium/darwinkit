// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WorkspaceOpenConfiguration] class.
var WorkspaceOpenConfigurationClass = _WorkspaceOpenConfigurationClass{objc.GetClass("NSWorkspaceOpenConfiguration")}

type _WorkspaceOpenConfigurationClass struct {
	objc.Class
}

// An interface definition for the [WorkspaceOpenConfiguration] class.
type IWorkspaceOpenConfiguration interface {
	objc.IObject
	Arguments() []string
	SetArguments(value []string)
	Hides() bool
	SetHides(value bool)
	AllowsRunningApplicationSubstitution() bool
	SetAllowsRunningApplicationSubstitution(value bool)
	IsForPrinting() bool
	SetForPrinting(value bool)
	HidesOthers() bool
	SetHidesOthers(value bool)
	RequiresUniversalLinks() bool
	SetRequiresUniversalLinks(value bool)
	Activates() bool
	SetActivates(value bool)
	PromptsUserIfNeeded() bool
	SetPromptsUserIfNeeded(value bool)
	CreatesNewApplicationInstance() bool
	SetCreatesNewApplicationInstance(value bool)
	AppleEvent() foundation.AppleEventDescriptor
	SetAppleEvent(value foundation.IAppleEventDescriptor)
	Environment() map[string]string
	SetEnvironment(value map[string]string)
	AddsToRecentItems() bool
	SetAddsToRecentItems(value bool)
}

// The configuration options for opening URLs or launching apps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration?language=objc
type WorkspaceOpenConfiguration struct {
	objc.Object
}

func WorkspaceOpenConfigurationFrom(ptr unsafe.Pointer) WorkspaceOpenConfiguration {
	return WorkspaceOpenConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WorkspaceOpenConfigurationClass) Configuration() WorkspaceOpenConfiguration {
	rv := objc.Call[WorkspaceOpenConfiguration](wc, objc.Sel("configuration"))
	return rv
}

// Creates and returns a new workspace configuration object containing default values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172709-configuration?language=objc
func WorkspaceOpenConfiguration_Configuration() WorkspaceOpenConfiguration {
	return WorkspaceOpenConfigurationClass.Configuration()
}

func (wc _WorkspaceOpenConfigurationClass) Alloc() WorkspaceOpenConfiguration {
	rv := objc.Call[WorkspaceOpenConfiguration](wc, objc.Sel("alloc"))
	return rv
}

func WorkspaceOpenConfiguration_Alloc() WorkspaceOpenConfiguration {
	return WorkspaceOpenConfigurationClass.Alloc()
}

func (wc _WorkspaceOpenConfigurationClass) New() WorkspaceOpenConfiguration {
	rv := objc.Call[WorkspaceOpenConfiguration](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWorkspaceOpenConfiguration() WorkspaceOpenConfiguration {
	return WorkspaceOpenConfigurationClass.New()
}

func (w_ WorkspaceOpenConfiguration) Init() WorkspaceOpenConfiguration {
	rv := objc.Call[WorkspaceOpenConfiguration](w_, objc.Sel("init"))
	return rv
}

// The set of command-line arguments to pass to a new app instance at launch time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172708-arguments?language=objc
func (w_ WorkspaceOpenConfiguration) Arguments() []string {
	rv := objc.Call[[]string](w_, objc.Sel("arguments"))
	return rv
}

// The set of command-line arguments to pass to a new app instance at launch time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172708-arguments?language=objc
func (w_ WorkspaceOpenConfiguration) SetArguments(value []string) {
	objc.Call[objc.Void](w_, objc.Sel("setArguments:"), value)
}

// A Boolean value indicating whether you want the app to hide itself after it launches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172713-hides?language=objc
func (w_ WorkspaceOpenConfiguration) Hides() bool {
	rv := objc.Call[bool](w_, objc.Sel("hides"))
	return rv
}

// A Boolean value indicating whether you want the app to hide itself after it launches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172713-hides?language=objc
func (w_ WorkspaceOpenConfiguration) SetHides(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setHides:"), value)
}

// A Boolean value that indicates whether to use a running instance of an application even if it’s at a different URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3375728-allowsrunningapplicationsubstitu?language=objc
func (w_ WorkspaceOpenConfiguration) AllowsRunningApplicationSubstitution() bool {
	rv := objc.Call[bool](w_, objc.Sel("allowsRunningApplicationSubstitution"))
	return rv
}

// A Boolean value that indicates whether to use a running instance of an application even if it’s at a different URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3375728-allowsrunningapplicationsubstitu?language=objc
func (w_ WorkspaceOpenConfiguration) SetAllowsRunningApplicationSubstitution(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setAllowsRunningApplicationSubstitution:"), value)
}

// A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172712-forprinting?language=objc
func (w_ WorkspaceOpenConfiguration) IsForPrinting() bool {
	rv := objc.Call[bool](w_, objc.Sel("isForPrinting"))
	return rv
}

// A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172712-forprinting?language=objc
func (w_ WorkspaceOpenConfiguration) SetForPrinting(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setForPrinting:"), value)
}

// A Boolean value indicating whether you want to hide all apps except the one that launched. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172714-hidesothers?language=objc
func (w_ WorkspaceOpenConfiguration) HidesOthers() bool {
	rv := objc.Call[bool](w_, objc.Sel("hidesOthers"))
	return rv
}

// A Boolean value indicating whether you want to hide all apps except the one that launched. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172714-hidesothers?language=objc
func (w_ WorkspaceOpenConfiguration) SetHidesOthers(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setHidesOthers:"), value)
}

// A Boolean value indicating whether you require the URL to have an associated universal link. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172716-requiresuniversallinks?language=objc
func (w_ WorkspaceOpenConfiguration) RequiresUniversalLinks() bool {
	rv := objc.Call[bool](w_, objc.Sel("requiresUniversalLinks"))
	return rv
}

// A Boolean value indicating whether you require the URL to have an associated universal link. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172716-requiresuniversallinks?language=objc
func (w_ WorkspaceOpenConfiguration) SetRequiresUniversalLinks(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setRequiresUniversalLinks:"), value)
}

// A Boolean value indicating whether the system activates the app and brings it to the foreground. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172704-activates?language=objc
func (w_ WorkspaceOpenConfiguration) Activates() bool {
	rv := objc.Call[bool](w_, objc.Sel("activates"))
	return rv
}

// A Boolean value indicating whether the system activates the app and brings it to the foreground. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172704-activates?language=objc
func (w_ WorkspaceOpenConfiguration) SetActivates(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setActivates:"), value)
}

// A Boolean value indicating whether to display errors, authentication requests, or other UI elements to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172715-promptsuserifneeded?language=objc
func (w_ WorkspaceOpenConfiguration) PromptsUserIfNeeded() bool {
	rv := objc.Call[bool](w_, objc.Sel("promptsUserIfNeeded"))
	return rv
}

// A Boolean value indicating whether to display errors, authentication requests, or other UI elements to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172715-promptsuserifneeded?language=objc
func (w_ WorkspaceOpenConfiguration) SetPromptsUserIfNeeded(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setPromptsUserIfNeeded:"), value)
}

// A Boolean value indicating whether you want the system to launch a new instance of the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172710-createsnewapplicationinstance?language=objc
func (w_ WorkspaceOpenConfiguration) CreatesNewApplicationInstance() bool {
	rv := objc.Call[bool](w_, objc.Sel("createsNewApplicationInstance"))
	return rv
}

// A Boolean value indicating whether you want the system to launch a new instance of the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172710-createsnewapplicationinstance?language=objc
func (w_ WorkspaceOpenConfiguration) SetCreatesNewApplicationInstance(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setCreatesNewApplicationInstance:"), value)
}

// The first Apple event to send to the new app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172706-appleevent?language=objc
func (w_ WorkspaceOpenConfiguration) AppleEvent() foundation.AppleEventDescriptor {
	rv := objc.Call[foundation.AppleEventDescriptor](w_, objc.Sel("appleEvent"))
	return rv
}

// The first Apple event to send to the new app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172706-appleevent?language=objc
func (w_ WorkspaceOpenConfiguration) SetAppleEvent(value foundation.IAppleEventDescriptor) {
	objc.Call[objc.Void](w_, objc.Sel("setAppleEvent:"), objc.Ptr(value))
}

// The set of environment variables to set in a new app instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172711-environment?language=objc
func (w_ WorkspaceOpenConfiguration) Environment() map[string]string {
	rv := objc.Call[map[string]string](w_, objc.Sel("environment"))
	return rv
}

// The set of environment variables to set in a new app instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172711-environment?language=objc
func (w_ WorkspaceOpenConfiguration) SetEnvironment(value map[string]string) {
	objc.Call[objc.Void](w_, objc.Sel("setEnvironment:"), value)
}

// A Boolean value indicating whether to add the app or documents to the Recent Items menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172705-addstorecentitems?language=objc
func (w_ WorkspaceOpenConfiguration) AddsToRecentItems() bool {
	rv := objc.Call[bool](w_, objc.Sel("addsToRecentItems"))
	return rv
}

// A Boolean value indicating whether to add the app or documents to the Recent Items menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspaceopenconfiguration/3172705-addstorecentitems?language=objc
func (w_ WorkspaceOpenConfiguration) SetAddsToRecentItems(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setAddsToRecentItems:"), value)
}
