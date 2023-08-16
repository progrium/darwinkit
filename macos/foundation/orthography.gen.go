// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Orthography] class.
var OrthographyClass = _OrthographyClass{objc.GetClass("NSOrthography")}

type _OrthographyClass struct {
	objc.Class
}

// An interface definition for the [Orthography] class.
type IOrthography interface {
	objc.IObject
	LanguagesForScript(script string) []string
	DominantLanguageForScript(script string) string
	AllLanguages() []string
	DominantScript() string
	AllScripts() []string
	DominantLanguage() string
	LanguageMap() map[string][]string
}

// A description of the linguistic content of natural language text, typically used for spelling and grammar checking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorthography?language=objc
type Orthography struct {
	objc.Object
}

func OrthographyFrom(ptr unsafe.Pointer) Orthography {
	return Orthography{
		Object: objc.ObjectFrom(ptr),
	}
}

func (oc _OrthographyClass) OrthographyWithDominantScriptLanguageMap(script string, map_ map[string][]string) Orthography {
	rv := objc.Call[Orthography](oc, objc.Sel("orthographyWithDominantScript:languageMap:"), script, map_)
	return rv
}

// Creates and returns an orthography object with the specified dominant script and language map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorthography/1585529-orthographywithdominantscript?language=objc
func Orthography_OrthographyWithDominantScriptLanguageMap(script string, map_ map[string][]string) Orthography {
	return OrthographyClass.OrthographyWithDominantScriptLanguageMap(script, map_)
}

func (oc _OrthographyClass) DefaultOrthographyForLanguage(language string) Orthography {
	rv := objc.Call[Orthography](oc, objc.Sel("defaultOrthographyForLanguage:"), language)
	return rv
}

// Creates and returns an orthography object with the default language map for the specified language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorthography/2875843-defaultorthographyforlanguage?language=objc
func Orthography_DefaultOrthographyForLanguage(language string) Orthography {
	return OrthographyClass.DefaultOrthographyForLanguage(language)
}

func (o_ Orthography) InitWithDominantScriptLanguageMap(script string, map_ map[string][]string) Orthography {
	rv := objc.Call[Orthography](o_, objc.Sel("initWithDominantScript:languageMap:"), script, map_)
	return rv
}

// Creates an orthography object with the specified dominant script and language map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorthography/1408708-initwithdominantscript?language=objc
func Orthography_InitWithDominantScriptLanguageMap(script string, map_ map[string][]string) Orthography {
	return OrthographyClass.Alloc().InitWithDominantScriptLanguageMap(script, map_)
}

func (oc _OrthographyClass) Alloc() Orthography {
	rv := objc.Call[Orthography](oc, objc.Sel("alloc"))
	return rv
}

func Orthography_Alloc() Orthography {
	return OrthographyClass.Alloc()
}

func (oc _OrthographyClass) New() Orthography {
	rv := objc.Call[Orthography](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOrthography() Orthography {
	return OrthographyClass.New()
}

func (o_ Orthography) Init() Orthography {
	rv := objc.Call[Orthography](o_, objc.Sel("init"))
	return rv
}

// Returns the list of languages for the specified script. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorthography/1412606-languagesforscript?language=objc
func (o_ Orthography) LanguagesForScript(script string) []string {
	rv := objc.Call[[]string](o_, objc.Sel("languagesForScript:"), script)
	return rv
}

// Returns the dominant language for the specified script. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorthography/1407326-dominantlanguageforscript?language=objc
func (o_ Orthography) DominantLanguageForScript(script string) string {
	rv := objc.Call[string](o_, objc.Sel("dominantLanguageForScript:"), script)
	return rv
}

// The languages appearing in values of the language map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorthography/1416205-alllanguages?language=objc
func (o_ Orthography) AllLanguages() []string {
	rv := objc.Call[[]string](o_, objc.Sel("allLanguages"))
	return rv
}

// The dominant script for the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorthography/1407965-dominantscript?language=objc
func (o_ Orthography) DominantScript() string {
	rv := objc.Call[string](o_, objc.Sel("dominantScript"))
	return rv
}

// The scripts appearing as keys in the language map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorthography/1410722-allscripts?language=objc
func (o_ Orthography) AllScripts() []string {
	rv := objc.Call[[]string](o_, objc.Sel("allScripts"))
	return rv
}

// The first language in the list of languages for the dominant script. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorthography/1415229-dominantlanguage?language=objc
func (o_ Orthography) DominantLanguage() string {
	rv := objc.Call[string](o_, objc.Sel("dominantLanguage"))
	return rv
}

// A dictionary that maps script tags to arrays of language tags. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorthography/1409533-languagemap?language=objc
func (o_ Orthography) LanguageMap() map[string][]string {
	rv := objc.Call[map[string][]string](o_, objc.Sel("languageMap"))
	return rv
}
