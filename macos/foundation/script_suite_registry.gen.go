// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScriptSuiteRegistry] class.
var ScriptSuiteRegistryClass = _ScriptSuiteRegistryClass{objc.GetClass("NSScriptSuiteRegistry")}

type _ScriptSuiteRegistryClass struct {
	objc.Class
}

// An interface definition for the [ScriptSuiteRegistry] class.
type IScriptSuiteRegistry interface {
	objc.IObject
	RegisterClassDescription(classDescription IScriptClassDescription)
	RegisterCommandDescription(commandDescription IScriptCommandDescription)
	SuiteForAppleEventCode(appleEventCode uint) string
	LoadSuitesFromBundle(bundle IBundle)
	CommandDescriptionsInSuite(suiteName string) map[string]ScriptCommandDescription
	ClassDescriptionsInSuite(suiteName string) map[string]ScriptClassDescription
	ClassDescriptionWithAppleEventCode(appleEventCode uint) ScriptClassDescription
	LoadSuiteWithDictionaryFromBundle(suiteDeclaration Dictionary, bundle IBundle)
	AppleEventCodeForSuite(suiteName string) uint
	BundleForSuite(suiteName string) Bundle
	AeteResource(languageName string) []byte
	CommandDescriptionWithAppleEventClassAndAppleEventCode(appleEventClassCode uint, appleEventIDCode uint) ScriptCommandDescription
	SuiteNames() []string
}

// The top-level repository of scriptability information for an app at runtime. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry?language=objc
type ScriptSuiteRegistry struct {
	objc.Object
}

func ScriptSuiteRegistryFrom(ptr unsafe.Pointer) ScriptSuiteRegistry {
	return ScriptSuiteRegistry{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _ScriptSuiteRegistryClass) Alloc() ScriptSuiteRegistry {
	rv := objc.Call[ScriptSuiteRegistry](sc, objc.Sel("alloc"))
	return rv
}

func ScriptSuiteRegistry_Alloc() ScriptSuiteRegistry {
	return ScriptSuiteRegistryClass.Alloc()
}

func (sc _ScriptSuiteRegistryClass) New() ScriptSuiteRegistry {
	rv := objc.Call[ScriptSuiteRegistry](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScriptSuiteRegistry() ScriptSuiteRegistry {
	return ScriptSuiteRegistryClass.New()
}

func (s_ ScriptSuiteRegistry) Init() ScriptSuiteRegistry {
	rv := objc.Call[ScriptSuiteRegistry](s_, objc.Sel("init"))
	return rv
}

// Returns the single, shared instance of NSScriptSuiteRegistry, creating it first if it doesn’t exist. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1417166-sharedscriptsuiteregistry?language=objc
func (sc _ScriptSuiteRegistryClass) SharedScriptSuiteRegistry() ScriptSuiteRegistry {
	rv := objc.Call[ScriptSuiteRegistry](sc, objc.Sel("sharedScriptSuiteRegistry"))
	return rv
}

// Returns the single, shared instance of NSScriptSuiteRegistry, creating it first if it doesn’t exist. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1417166-sharedscriptsuiteregistry?language=objc
func ScriptSuiteRegistry_SharedScriptSuiteRegistry() ScriptSuiteRegistry {
	return ScriptSuiteRegistryClass.SharedScriptSuiteRegistry()
}

// Registers class description classDescription for use by Cocoa’s built-in scripting support by storing it in a per-suite internal dictionary under the class name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1412869-registerclassdescription?language=objc
func (s_ ScriptSuiteRegistry) RegisterClassDescription(classDescription IScriptClassDescription) {
	objc.Call[objc.Void](s_, objc.Sel("registerClassDescription:"), objc.Ptr(classDescription))
}

// Registers command description commandDesc for use by Cocoa’s built-in scripting support by storing it in a per-suite internal dictionary under the command name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1408858-registercommanddescription?language=objc
func (s_ ScriptSuiteRegistry) RegisterCommandDescription(commandDescription IScriptCommandDescription) {
	objc.Call[objc.Void](s_, objc.Sel("registerCommandDescription:"), objc.Ptr(commandDescription))
}

// Returns the name of the suite definition associated with the given four-character Apple event code, code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1415178-suiteforappleeventcode?language=objc
func (s_ ScriptSuiteRegistry) SuiteForAppleEventCode(appleEventCode uint) string {
	rv := objc.Call[string](s_, objc.Sel("suiteForAppleEventCode:"), appleEventCode)
	return rv
}

// Loads the suite definitions in bundle aBundle, invoking loadSuiteWithDictionary:fromBundle: for each suite found. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1410575-loadsuitesfrombundle?language=objc
func (s_ ScriptSuiteRegistry) LoadSuitesFromBundle(bundle IBundle) {
	objc.Call[objc.Void](s_, objc.Sel("loadSuitesFromBundle:"), objc.Ptr(bundle))
}

// Returns the command descriptions contained in the suite identified by suiteName. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1416396-commanddescriptionsinsuite?language=objc
func (s_ ScriptSuiteRegistry) CommandDescriptionsInSuite(suiteName string) map[string]ScriptCommandDescription {
	rv := objc.Call[map[string]ScriptCommandDescription](s_, objc.Sel("commandDescriptionsInSuite:"), suiteName)
	return rv
}

// Returns the class descriptions contained in the suite identified by suiteName. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1414328-classdescriptionsinsuite?language=objc
func (s_ ScriptSuiteRegistry) ClassDescriptionsInSuite(suiteName string) map[string]ScriptClassDescription {
	rv := objc.Call[map[string]ScriptClassDescription](s_, objc.Sel("classDescriptionsInSuite:"), suiteName)
	return rv
}

// Returns the class description associated with the given four-character Apple event code, code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1411184-classdescriptionwithappleeventco?language=objc
func (s_ ScriptSuiteRegistry) ClassDescriptionWithAppleEventCode(appleEventCode uint) ScriptClassDescription {
	rv := objc.Call[ScriptClassDescription](s_, objc.Sel("classDescriptionWithAppleEventCode:"), appleEventCode)
	return rv
}

// Loads the suite definition encapsulated in dictionary; previously, this suite definition was parsed from a .scriptSuite property list contained in a framework or in bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1413397-loadsuitewithdictionary?language=objc
func (s_ ScriptSuiteRegistry) LoadSuiteWithDictionaryFromBundle(suiteDeclaration Dictionary, bundle IBundle) {
	objc.Call[objc.Void](s_, objc.Sel("loadSuiteWithDictionary:fromBundle:"), suiteDeclaration, objc.Ptr(bundle))
}

// Returns the Apple event code associated with the suite named suiteName, such as ‘core’ for the Core suite. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1412492-appleeventcodeforsuite?language=objc
func (s_ ScriptSuiteRegistry) AppleEventCodeForSuite(suiteName string) uint {
	rv := objc.Call[uint](s_, objc.Sel("appleEventCodeForSuite:"), suiteName)
	return rv
}

// Returns the bundle containing the suite-definition property list (extension .scriptSuite) identified by suiteName. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1414868-bundleforsuite?language=objc
func (s_ ScriptSuiteRegistry) BundleForSuite(suiteName string) Bundle {
	rv := objc.Call[Bundle](s_, objc.Sel("bundleForSuite:"), suiteName)
	return rv
}

// Returns an NSData object that contains data in 'aete' resource format describing the scriptability information currently known to the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1409186-aeteresource?language=objc
func (s_ ScriptSuiteRegistry) AeteResource(languageName string) []byte {
	rv := objc.Call[[]byte](s_, objc.Sel("aeteResource:"), languageName)
	return rv
}

// Returns the command description identified by a suite’s four-character Apple event code of the class (eventClass) and the four-character Apple event code of the command (commandCode). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1416734-commanddescriptionwithappleevent?language=objc
func (s_ ScriptSuiteRegistry) CommandDescriptionWithAppleEventClassAndAppleEventCode(appleEventClassCode uint, appleEventIDCode uint) ScriptCommandDescription {
	rv := objc.Call[ScriptCommandDescription](s_, objc.Sel("commandDescriptionWithAppleEventClass:andAppleEventCode:"), appleEventClassCode, appleEventIDCode)
	return rv
}

// Sets the single, shared instance of NSScriptSuiteRegistry to registry. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1409955-setsharedscriptsuiteregistry?language=objc
func (sc _ScriptSuiteRegistryClass) SetSharedScriptSuiteRegistry(registry IScriptSuiteRegistry) {
	objc.Call[objc.Void](sc, objc.Sel("setSharedScriptSuiteRegistry:"), objc.Ptr(registry))
}

// Sets the single, shared instance of NSScriptSuiteRegistry to registry. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1409955-setsharedscriptsuiteregistry?language=objc
func ScriptSuiteRegistry_SetSharedScriptSuiteRegistry(registry IScriptSuiteRegistry) {
	ScriptSuiteRegistryClass.SetSharedScriptSuiteRegistry(registry)
}

// Returns the names of the suite definitions currently loaded by the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/1414452-suitenames?language=objc
func (s_ ScriptSuiteRegistry) SuiteNames() []string {
	rv := objc.Call[[]string](s_, objc.Sel("suiteNames"))
	return rv
}
