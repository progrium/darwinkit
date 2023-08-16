// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The optional methods implemented by the delegate of a spell server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate?language=objc
type PSpellServerDelegate interface {
	// optional
	SpellServerDidLearnWordInLanguage(sender SpellServer, word string, language string)
	HasSpellServerDidLearnWordInLanguage() bool
}

// A delegate implementation builder for the [PSpellServerDelegate] protocol.
type SpellServerDelegate struct {
	_SpellServerDidLearnWordInLanguage func(sender SpellServer, word string, language string)
}

func (di *SpellServerDelegate) HasSpellServerDidLearnWordInLanguage() bool {
	return di._SpellServerDidLearnWordInLanguage != nil
}

// Notifies the delegate that the sender has added the specified word to the user’s list of acceptable words in the specified language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1407851-spellserver?language=objc
func (di *SpellServerDelegate) SetSpellServerDidLearnWordInLanguage(f func(sender SpellServer, word string, language string)) {
	di._SpellServerDidLearnWordInLanguage = f
}

// Notifies the delegate that the sender has added the specified word to the user’s list of acceptable words in the specified language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1407851-spellserver?language=objc
func (di *SpellServerDelegate) SpellServerDidLearnWordInLanguage(sender SpellServer, word string, language string) {
	di._SpellServerDidLearnWordInLanguage(sender, word, language)
}

// A concrete type wrapper for the [PSpellServerDelegate] protocol.
type SpellServerDelegateWrapper struct {
	objc.Object
}

func (s_ SpellServerDelegateWrapper) HasSpellServerDidLearnWordInLanguage() bool {
	return s_.RespondsToSelector(objc.Sel("spellServer:didLearnWord:inLanguage:"))
}

// Notifies the delegate that the sender has added the specified word to the user’s list of acceptable words in the specified language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1407851-spellserver?language=objc
func (s_ SpellServerDelegateWrapper) SpellServerDidLearnWordInLanguage(sender ISpellServer, word string, language string) {
	objc.Call[objc.Void](s_, objc.Sel("spellServer:didLearnWord:inLanguage:"), objc.Ptr(sender), word, language)
}
