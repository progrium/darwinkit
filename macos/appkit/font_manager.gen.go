// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var FontManagerClass = _FontManagerClass{objc.GetClass("NSFontManager")}

type _FontManagerClass struct {
	objc.Class
}

type IFontManager interface {
	objc.IObject
	AvailableFontNamesWithTraits(someTraits FontTraitMask) []string
	AvailableMembersOfFontFamily(fam string) [][]objc.Object
	SetSelectedFontIsMultiple(fontObj IFont, flag bool)
	SendAction() bool
	LocalizedNameForFamilyFace(family string, faceKey string) string
	AddFontTrait(sender objc.IObject)
	RemoveFontTrait(sender objc.IObject)
	ModifyFont(sender objc.IObject)
	ModifyFontViaPanel(sender objc.IObject)
	OrderFrontStylesPanel(sender objc.IObject)
	OrderFrontFontPanel(sender objc.IObject)
	ConvertFont(fontObj IFont) Font
	ConvertFontToFace(fontObj IFont, typeface string) Font
	ConvertFontToFamily(fontObj IFont, family string) Font
	ConvertFontToHaveTrait(fontObj IFont, trait FontTraitMask) Font
	ConvertFontToNotHaveTrait(fontObj IFont, trait FontTraitMask) Font
	ConvertFontToSize(fontObj IFont, size float64) Font
	ConvertWeightOfFont(upFlag bool, fontObj IFont) Font
	ConvertFontTraits(traits FontTraitMask) FontTraitMask
	FontWithFamilyTraitsWeightSize(family string, traits FontTraitMask, weight int, size float64) Font
	TraitsOfFont(fontObj IFont) FontTraitMask
	FontNamedHasTraits(fName string, someTraits FontTraitMask) bool
	WeightOfFont(fontObj IFont) int
	FontPanel(create bool) FontPanel
	SetFontMenu(newMenu IMenu)
	FontMenu(create bool) Menu
	SetSelectedAttributesIsMultiple(attributes map[string]objc.IObject, flag bool)
	ConvertAttributes(attributes map[string]objc.IObject) map[string]objc.Object
	AvailableFonts() []string
	AvailableFontFamilies() []string
	SelectedFont() Font
	IsMultiple() bool
	CurrentFontAction() FontAction
	IsEnabled() bool
	SetEnabled(value bool)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Target() objc.Object
	SetTarget(value objc.IObject)
}

type FontManager struct {
	objc.Object
}

func MakeFontManager(ptr unsafe.Pointer) FontManager {
	return FontManager{
		Object: objc.MakeObject(ptr),
	}
}

func (fc _FontManagerClass) Alloc() FontManager {
	rv := objc.CallMethod[FontManager](fc, objc.GetSelector("alloc"))
	return rv
}

func FontManager_Alloc() FontManager {
	return FontManagerClass.Alloc()
}

func (fc _FontManagerClass) New() FontManager {
	rv := objc.CallMethod[FontManager](fc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewFontManager() FontManager {
	return FontManagerClass.New()
}

func FontManager_New() FontManager {
	return FontManagerClass.New()
}

func (f_ FontManager) Init() FontManager {
	rv := objc.CallMethod[FontManager](f_, objc.GetSelector("init"))
	return rv
}

func FontManager_Init() FontManager {
	return FontManagerClass.Alloc().Init()
}

func (fc _FontManagerClass) SetFontManagerFactory(factoryId objc.IClass) {
	objc.CallMethod[objc.Void](fc, objc.GetSelector("setFontManagerFactory:"), objc.ExtractPtr(factoryId))
}

func FontManager_SetFontManagerFactory(factoryId objc.IClass) {
	FontManagerClass.SetFontManagerFactory(factoryId)
}

func (fc _FontManagerClass) SetFontPanelFactory(factoryId objc.IClass) {
	objc.CallMethod[objc.Void](fc, objc.GetSelector("setFontPanelFactory:"), objc.ExtractPtr(factoryId))
}

func FontManager_SetFontPanelFactory(factoryId objc.IClass) {
	FontManagerClass.SetFontPanelFactory(factoryId)
}

func (f_ FontManager) AvailableFontNamesWithTraits(someTraits FontTraitMask) []string {
	rv := objc.CallMethod[[]string](f_, objc.GetSelector("availableFontNamesWithTraits:"), someTraits)
	return rv
}

func (f_ FontManager) AvailableMembersOfFontFamily(fam string) [][]objc.Object {
	rv := objc.CallMethod[[][]objc.Object](f_, objc.GetSelector("availableMembersOfFontFamily:"), fam)
	return rv
}

func (f_ FontManager) SetSelectedFontIsMultiple(fontObj IFont, flag bool) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setSelectedFont:isMultiple:"), objc.ExtractPtr(fontObj), flag)
}

func (f_ FontManager) SendAction() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("sendAction"))
	return rv
}

func (f_ FontManager) LocalizedNameForFamilyFace(family string, faceKey string) string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("localizedNameForFamily:face:"), family, faceKey)
	return rv
}

func (f_ FontManager) AddFontTrait(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("addFontTrait:"), objc.ExtractPtr(sender))
}

func (f_ FontManager) RemoveFontTrait(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("removeFontTrait:"), objc.ExtractPtr(sender))
}

func (f_ FontManager) ModifyFont(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("modifyFont:"), objc.ExtractPtr(sender))
}

func (f_ FontManager) ModifyFontViaPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("modifyFontViaPanel:"), objc.ExtractPtr(sender))
}

func (f_ FontManager) OrderFrontStylesPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("orderFrontStylesPanel:"), objc.ExtractPtr(sender))
}

func (f_ FontManager) OrderFrontFontPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("orderFrontFontPanel:"), objc.ExtractPtr(sender))
}

func (f_ FontManager) ConvertFont(fontObj IFont) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("convertFont:"), objc.ExtractPtr(fontObj))
	return rv
}

func (f_ FontManager) ConvertFontToFace(fontObj IFont, typeface string) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("convertFont:toFace:"), objc.ExtractPtr(fontObj), typeface)
	return rv
}

func (f_ FontManager) ConvertFontToFamily(fontObj IFont, family string) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("convertFont:toFamily:"), objc.ExtractPtr(fontObj), family)
	return rv
}

func (f_ FontManager) ConvertFontToHaveTrait(fontObj IFont, trait FontTraitMask) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("convertFont:toHaveTrait:"), objc.ExtractPtr(fontObj), trait)
	return rv
}

func (f_ FontManager) ConvertFontToNotHaveTrait(fontObj IFont, trait FontTraitMask) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("convertFont:toNotHaveTrait:"), objc.ExtractPtr(fontObj), trait)
	return rv
}

func (f_ FontManager) ConvertFontToSize(fontObj IFont, size float64) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("convertFont:toSize:"), objc.ExtractPtr(fontObj), size)
	return rv
}

func (f_ FontManager) ConvertWeightOfFont(upFlag bool, fontObj IFont) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("convertWeight:ofFont:"), upFlag, objc.ExtractPtr(fontObj))
	return rv
}

func (f_ FontManager) ConvertFontTraits(traits FontTraitMask) FontTraitMask {
	rv := objc.CallMethod[FontTraitMask](f_, objc.GetSelector("convertFontTraits:"), traits)
	return rv
}

func (f_ FontManager) FontWithFamilyTraitsWeightSize(family string, traits FontTraitMask, weight int, size float64) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("fontWithFamily:traits:weight:size:"), family, traits, weight, size)
	return rv
}

func (f_ FontManager) TraitsOfFont(fontObj IFont) FontTraitMask {
	rv := objc.CallMethod[FontTraitMask](f_, objc.GetSelector("traitsOfFont:"), objc.ExtractPtr(fontObj))
	return rv
}

func (f_ FontManager) FontNamedHasTraits(fName string, someTraits FontTraitMask) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("fontNamed:hasTraits:"), fName, someTraits)
	return rv
}

func (f_ FontManager) WeightOfFont(fontObj IFont) int {
	rv := objc.CallMethod[int](f_, objc.GetSelector("weightOfFont:"), objc.ExtractPtr(fontObj))
	return rv
}

func (f_ FontManager) FontPanel(create bool) FontPanel {
	rv := objc.CallMethod[FontPanel](f_, objc.GetSelector("fontPanel:"), create)
	return rv
}

func (f_ FontManager) SetFontMenu(newMenu IMenu) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setFontMenu:"), objc.ExtractPtr(newMenu))
}

func (f_ FontManager) FontMenu(create bool) Menu {
	rv := objc.CallMethod[Menu](f_, objc.GetSelector("fontMenu:"), create)
	return rv
}

func (f_ FontManager) SetSelectedAttributesIsMultiple(attributes map[string]objc.IObject, flag bool) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setSelectedAttributes:isMultiple:"), attributes, flag)
}

func (f_ FontManager) ConvertAttributes(attributes map[string]objc.IObject) map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](f_, objc.GetSelector("convertAttributes:"), attributes)
	return rv
}

func (fc _FontManagerClass) SharedFontManager() FontManager {
	rv := objc.CallMethod[FontManager](fc, objc.GetSelector("sharedFontManager"))
	return rv
}

func FontManager_SharedFontManager() FontManager {
	return FontManagerClass.SharedFontManager()
}

func (f_ FontManager) AvailableFonts() []string {
	rv := objc.CallMethod[[]string](f_, objc.GetSelector("availableFonts"))
	return rv
}

func (f_ FontManager) AvailableFontFamilies() []string {
	rv := objc.CallMethod[[]string](f_, objc.GetSelector("availableFontFamilies"))
	return rv
}

func (f_ FontManager) SelectedFont() Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("selectedFont"))
	return rv
}

func (f_ FontManager) IsMultiple() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isMultiple"))
	return rv
}

func (f_ FontManager) CurrentFontAction() FontAction {
	rv := objc.CallMethod[FontAction](f_, objc.GetSelector("currentFontAction"))
	return rv
}

func (f_ FontManager) IsEnabled() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isEnabled"))
	return rv
}

func (f_ FontManager) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setEnabled:"), value)
}

func (f_ FontManager) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](f_, objc.GetSelector("action"))
	return rv
}

func (f_ FontManager) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setAction:"), value)
}

func (f_ FontManager) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](f_, objc.GetSelector("target"))
	return rv
}

func (f_ FontManager) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setTarget:"), objc.ExtractPtr(value))
}
