// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var OrthographyClass = _OrthographyClass{objc.GetClass("NSOrthography")}

type _OrthographyClass struct {
	objc.Class
}

type IOrthography interface {
	objc.IObject
	DominantLanguageForScript(script string) string
	LanguagesForScript(script string) []string
	LanguageMap() map[string][]string
	DominantLanguage() string
	DominantScript() string
	AllScripts() []string
	AllLanguages() []string
}

type Orthography struct {
	objc.Object
}

func MakeOrthography(ptr unsafe.Pointer) Orthography {
	return Orthography{
		Object: objc.MakeObject(ptr),
	}
}

func (oc _OrthographyClass) DefaultOrthographyForLanguage(language string) Orthography {
	rv := objc.CallMethod[Orthography](oc, objc.GetSelector("defaultOrthographyForLanguage:"), language)
	return rv
}

func Orthography_DefaultOrthographyForLanguage(language string) Orthography {
	return OrthographyClass.DefaultOrthographyForLanguage(language)
}

func (o_ Orthography) InitWithDominantScriptLanguageMap(script string, map_ map[string][]string) Orthography {
	rv := objc.CallMethod[Orthography](o_, objc.GetSelector("initWithDominantScript:languageMap:"), script, map_)
	return rv
}

func Orthography_InitWithDominantScriptLanguageMap(script string, map_ map[string][]string) Orthography {
	return OrthographyClass.Alloc().InitWithDominantScriptLanguageMap(script, map_)
}

func (oc _OrthographyClass) OrthographyWithDominantScriptLanguageMap(script string, map_ map[string][]string) Orthography {
	rv := objc.CallMethod[Orthography](oc, objc.GetSelector("orthographyWithDominantScript:languageMap:"), script, map_)
	return rv
}

func Orthography_OrthographyWithDominantScriptLanguageMap(script string, map_ map[string][]string) Orthography {
	return OrthographyClass.OrthographyWithDominantScriptLanguageMap(script, map_)
}

func (oc _OrthographyClass) Alloc() Orthography {
	rv := objc.CallMethod[Orthography](oc, objc.GetSelector("alloc"))
	return rv
}

func Orthography_Alloc() Orthography {
	return OrthographyClass.Alloc()
}

func (oc _OrthographyClass) New() Orthography {
	rv := objc.CallMethod[Orthography](oc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewOrthography() Orthography {
	return OrthographyClass.New()
}

func Orthography_New() Orthography {
	return OrthographyClass.New()
}

func (o_ Orthography) Init() Orthography {
	rv := objc.CallMethod[Orthography](o_, objc.GetSelector("init"))
	return rv
}

func Orthography_Init() Orthography {
	return OrthographyClass.Alloc().Init()
}

func (o_ Orthography) DominantLanguageForScript(script string) string {
	rv := objc.CallMethod[string](o_, objc.GetSelector("dominantLanguageForScript:"), script)
	return rv
}

func (o_ Orthography) LanguagesForScript(script string) []string {
	rv := objc.CallMethod[[]string](o_, objc.GetSelector("languagesForScript:"), script)
	return rv
}

func (o_ Orthography) LanguageMap() map[string][]string {
	rv := objc.CallMethod[map[string][]string](o_, objc.GetSelector("languageMap"))
	return rv
}

func (o_ Orthography) DominantLanguage() string {
	rv := objc.CallMethod[string](o_, objc.GetSelector("dominantLanguage"))
	return rv
}

func (o_ Orthography) DominantScript() string {
	rv := objc.CallMethod[string](o_, objc.GetSelector("dominantScript"))
	return rv
}

func (o_ Orthography) AllScripts() []string {
	rv := objc.CallMethod[[]string](o_, objc.GetSelector("allScripts"))
	return rv
}

func (o_ Orthography) AllLanguages() []string {
	rv := objc.CallMethod[[]string](o_, objc.GetSelector("allLanguages"))
	return rv
}
