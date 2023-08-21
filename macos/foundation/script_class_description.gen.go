// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScriptClassDescription] class.
var ScriptClassDescriptionClass = _ScriptClassDescriptionClass{objc.GetClass("NSScriptClassDescription")}

type _ScriptClassDescriptionClass struct {
	objc.Class
}

// An interface definition for the [ScriptClassDescription] class.
type IScriptClassDescription interface {
	IClassDescription
	SelectorForCommand(commandDescription IScriptCommandDescription) objc.Selector
	MatchesAppleEventCode(appleEventCode uint) bool
	SupportsCommand(commandDescription IScriptCommandDescription) bool
	HasReadablePropertyForKey(key string) bool
	HasPropertyForKey(key string) bool
	ClassDescriptionForKey(key string) ScriptClassDescription
	HasWritablePropertyForKey(key string) bool
	TypeForKey(key string) string
	KeyWithAppleEventCode(appleEventCode uint) string
	IsLocationRequiredToCreateForKey(toManyRelationshipKey string) bool
	AppleEventCodeForKey(key string) uint
	HasOrderedToManyRelationshipForKey(key string) bool
	AppleEventCode() uint
	SuperclassDescription() ScriptClassDescription
	DefaultSubcontainerAttributeKey() string
	SuiteName() string
	ClassName() string
	ImplementationClassName() string
}

// A scriptable class that a macOS app supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription?language=objc
type ScriptClassDescription struct {
	ClassDescription
}

func ScriptClassDescriptionFrom(ptr unsafe.Pointer) ScriptClassDescription {
	return ScriptClassDescription{
		ClassDescription: ClassDescriptionFrom(ptr),
	}
}

func (s_ ScriptClassDescription) InitWithSuiteNameClassNameDictionary(suiteName string, className string, classDeclaration Dictionary) ScriptClassDescription {
	rv := objc.Call[ScriptClassDescription](s_, objc.Sel("initWithSuiteName:className:dictionary:"), suiteName, className, classDeclaration)
	return rv
}

// Initializes and returns a newly allocated instance of NSScriptClassDescription. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1410370-initwithsuitename?language=objc
func NewScriptClassDescriptionWithSuiteNameClassNameDictionary(suiteName string, className string, classDeclaration Dictionary) ScriptClassDescription {
	instance := ScriptClassDescriptionClass.Alloc().InitWithSuiteNameClassNameDictionary(suiteName, className, classDeclaration)
	instance.Autorelease()
	return instance
}

func (sc _ScriptClassDescriptionClass) Alloc() ScriptClassDescription {
	rv := objc.Call[ScriptClassDescription](sc, objc.Sel("alloc"))
	return rv
}

func ScriptClassDescription_Alloc() ScriptClassDescription {
	return ScriptClassDescriptionClass.Alloc()
}

func (sc _ScriptClassDescriptionClass) New() ScriptClassDescription {
	rv := objc.Call[ScriptClassDescription](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScriptClassDescription() ScriptClassDescription {
	return ScriptClassDescriptionClass.New()
}

func (s_ ScriptClassDescription) Init() ScriptClassDescription {
	rv := objc.Call[ScriptClassDescription](s_, objc.Sel("init"))
	return rv
}

// Returns the selector associated with the receiver for the specified command description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1409327-selectorforcommand?language=objc
func (s_ ScriptClassDescription) SelectorForCommand(commandDescription IScriptCommandDescription) objc.Selector {
	rv := objc.Call[objc.Selector](s_, objc.Sel("selectorForCommand:"), objc.Ptr(commandDescription))
	return rv
}

// Returns a Boolean value indicating whether a primary or secondary Apple event code in the receiver matches the passed code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1411166-matchesappleeventcode?language=objc
func (s_ ScriptClassDescription) MatchesAppleEventCode(appleEventCode uint) bool {
	rv := objc.Call[bool](s_, objc.Sel("matchesAppleEventCode:"), appleEventCode)
	return rv
}

// Returns a Boolean value indicating whether the receiver or any superclass supports the specified command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1411902-supportscommand?language=objc
func (s_ ScriptClassDescription) SupportsCommand(commandDescription IScriptCommandDescription) bool {
	rv := objc.Call[bool](s_, objc.Sel("supportsCommand:"), objc.Ptr(commandDescription))
	return rv
}

// Returns a Boolean value indicating whether the described class has a readable property identified by the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1411467-hasreadablepropertyforkey?language=objc
func (s_ ScriptClassDescription) HasReadablePropertyForKey(key string) bool {
	rv := objc.Call[bool](s_, objc.Sel("hasReadablePropertyForKey:"), key)
	return rv
}

// Returns a Boolean value indicating whether the described class has a property identified by the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1407574-haspropertyforkey?language=objc
func (s_ ScriptClassDescription) HasPropertyForKey(key string) bool {
	rv := objc.Call[bool](s_, objc.Sel("hasPropertyForKey:"), key)
	return rv
}

// Returns the class description instance for the class type of the specified attribute or relationship. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1407343-classdescriptionforkey?language=objc
func (s_ ScriptClassDescription) ClassDescriptionForKey(key string) ScriptClassDescription {
	rv := objc.Call[ScriptClassDescription](s_, objc.Sel("classDescriptionForKey:"), key)
	return rv
}

// Returns a Boolean value indicating whether the described class has a writable property identified by the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1411935-haswritablepropertyforkey?language=objc
func (s_ ScriptClassDescription) HasWritablePropertyForKey(key string) bool {
	rv := objc.Call[bool](s_, objc.Sel("hasWritablePropertyForKey:"), key)
	return rv
}

// Returns the name of the declared type of the attribute or relationship identified by the passed key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1417186-typeforkey?language=objc
func (s_ ScriptClassDescription) TypeForKey(key string) string {
	rv := objc.Call[string](s_, objc.Sel("typeForKey:"), key)
	return rv
}

// Given an Apple event code that identifies a property or element class, returns the key for the corresponding attribute, one-to-one relationship, or one-to-many relationship. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1415315-keywithappleeventcode?language=objc
func (s_ ScriptClassDescription) KeyWithAppleEventCode(appleEventCode uint) string {
	rv := objc.Call[string](s_, objc.Sel("keyWithAppleEventCode:"), appleEventCode)
	return rv
}

// Returns a Boolean value indicating whether an insertion location must be specified when creating a new object in the specified to-many relationship of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1416531-islocationrequiredtocreateforkey?language=objc
func (s_ ScriptClassDescription) IsLocationRequiredToCreateForKey(toManyRelationshipKey string) bool {
	rv := objc.Call[bool](s_, objc.Sel("isLocationRequiredToCreateForKey:"), toManyRelationshipKey)
	return rv
}

// Returns the Apple event code for the specified attribute or relationship in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1414657-appleeventcodeforkey?language=objc
func (s_ ScriptClassDescription) AppleEventCodeForKey(key string) uint {
	rv := objc.Call[uint](s_, objc.Sel("appleEventCodeForKey:"), key)
	return rv
}

// Returns a Boolean value indicating whether the described class has an ordered to-many relationship identified by the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1413542-hasorderedtomanyrelationshipfork?language=objc
func (s_ ScriptClassDescription) HasOrderedToManyRelationshipForKey(key string) bool {
	rv := objc.Call[bool](s_, objc.Sel("hasOrderedToManyRelationshipForKey:"), key)
	return rv
}

// Returns the Apple event code associated with the receiver’s class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1414920-appleeventcode?language=objc
func (s_ ScriptClassDescription) AppleEventCode() uint {
	rv := objc.Call[uint](s_, objc.Sel("appleEventCode"))
	return rv
}

// Returns the class description instance for the superclass of the receiver’s class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1416243-superclassdescription?language=objc
func (s_ ScriptClassDescription) SuperclassDescription() ScriptClassDescription {
	rv := objc.Call[ScriptClassDescription](s_, objc.Sel("superclassDescription"))
	return rv
}

// Returns the value of the DefaultSubcontainerAttribute entry of the class dictionary from which the receiver was instantiated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1410261-defaultsubcontainerattributekey?language=objc
func (s_ ScriptClassDescription) DefaultSubcontainerAttributeKey() string {
	rv := objc.Call[string](s_, objc.Sel("defaultSubcontainerAttributeKey"))
	return rv
}

// Returns the name of the receiver’s suite. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1410782-suitename?language=objc
func (s_ ScriptClassDescription) SuiteName() string {
	rv := objc.Call[string](s_, objc.Sel("suiteName"))
	return rv
}

// Returns the name of the class the receiver describes, as provided at initialization time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1418029-classname?language=objc
func (s_ ScriptClassDescription) ClassName() string {
	rv := objc.Call[string](s_, objc.Sel("className"))
	return rv
}

// Returns the name of the Objective-C class instantiated to implement the scripting class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/1409575-implementationclassname?language=objc
func (s_ ScriptClassDescription) ImplementationClassName() string {
	rv := objc.Call[string](s_, objc.Sel("implementationClassName"))
	return rv
}
