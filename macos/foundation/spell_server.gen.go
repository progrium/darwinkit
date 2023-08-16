// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SpellServer] class.
var SpellServerClass = _SpellServerClass{objc.GetClass("NSSpellServer")}

type _SpellServerClass struct {
	objc.Class
}

// An interface definition for the [SpellServer] class.
type ISpellServer interface {
	objc.IObject
	RegisterLanguageByVendor(language string, vendor string) bool
	IsWordInUserDictionariesCaseSensitive(word string, flag bool) bool
	Run()
	Delegate() SpellServerDelegateWrapper
	SetDelegate(value PSpellServerDelegate)
	SetDelegateObject(valueObject objc.IObject)
}

// A server that your app uses to provide a spell checker service to other apps running in the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserver?language=objc
type SpellServer struct {
	objc.Object
}

func SpellServerFrom(ptr unsafe.Pointer) SpellServer {
	return SpellServer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SpellServerClass) Alloc() SpellServer {
	rv := objc.Call[SpellServer](sc, objc.Sel("alloc"))
	return rv
}

func SpellServer_Alloc() SpellServer {
	return SpellServerClass.Alloc()
}

func (sc _SpellServerClass) New() SpellServer {
	rv := objc.Call[SpellServer](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSpellServer() SpellServer {
	return SpellServerClass.New()
}

func (s_ SpellServer) Init() SpellServer {
	rv := objc.Call[SpellServer](s_, objc.Sel("init"))
	return rv
}

// Notifies the receiver of a language your spelling checker can check. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserver/1411187-registerlanguage?language=objc
func (s_ SpellServer) RegisterLanguageByVendor(language string, vendor string) bool {
	rv := objc.Call[bool](s_, objc.Sel("registerLanguage:byVendor:"), language, vendor)
	return rv
}

// Indicates whether a given word is in the user’s list of learned words or the document’s list of words to ignore. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserver/1412254-iswordinuserdictionaries?language=objc
func (s_ SpellServer) IsWordInUserDictionariesCaseSensitive(word string, flag bool) bool {
	rv := objc.Call[bool](s_, objc.Sel("isWordInUserDictionaries:caseSensitive:"), word, flag)
	return rv
}

// Causes the receiver to start listening for spell-checking requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserver/1414343-run?language=objc
func (s_ SpellServer) Run() {
	objc.Call[objc.Void](s_, objc.Sel("run"))
}

// Returns the receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserver/1414240-delegate?language=objc
func (s_ SpellServer) Delegate() SpellServerDelegateWrapper {
	rv := objc.Call[SpellServerDelegateWrapper](s_, objc.Sel("delegate"))
	return rv
}

// Returns the receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserver/1414240-delegate?language=objc
func (s_ SpellServer) SetDelegate(value PSpellServerDelegate) {
	po0 := objc.WrapAsProtocol("NSSpellServerDelegate", value)
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), po0)
}

// Returns the receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserver/1414240-delegate?language=objc
func (s_ SpellServer) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}
