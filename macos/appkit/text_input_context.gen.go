// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var TextInputContextClass = _TextInputContextClass{objc.GetClass("NSTextInputContext")}

type _TextInputContextClass struct {
	objc.Class
}

type ITextInputContext interface {
	objc.IObject
	Activate()
	Deactivate()
	HandleEvent(event IEvent) bool
	DiscardMarkedText()
	InvalidateCharacterCoordinates()
	AcceptsGlyphInfo() bool
	SetAcceptsGlyphInfo(value bool)
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
	KeyboardInputSources() []TextInputSourceIdentifier
	SelectedKeyboardInputSource() TextInputSourceIdentifier
	SetSelectedKeyboardInputSource(value TextInputSourceIdentifier)
}

type TextInputContext struct {
	objc.Object
}

func MakeTextInputContext(ptr unsafe.Pointer) TextInputContext {
	return TextInputContext{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TextInputContextClass) Alloc() TextInputContext {
	rv := objc.CallMethod[TextInputContext](tc, objc.GetSelector("alloc"))
	return rv
}

func TextInputContext_Alloc() TextInputContext {
	return TextInputContextClass.Alloc()
}

func (tc _TextInputContextClass) New() TextInputContext {
	rv := objc.CallMethod[TextInputContext](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextInputContext() TextInputContext {
	return TextInputContextClass.New()
}

func TextInputContext_New() TextInputContext {
	return TextInputContextClass.New()
}

func (t_ TextInputContext) Init() TextInputContext {
	rv := objc.CallMethod[TextInputContext](t_, objc.GetSelector("init"))
	return rv
}

func TextInputContext_Init() TextInputContext {
	return TextInputContextClass.Alloc().Init()
}

func (t_ TextInputContext) Activate() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("activate"))
}

func (t_ TextInputContext) Deactivate() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("deactivate"))
}

func (t_ TextInputContext) HandleEvent(event IEvent) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("handleEvent:"), objc.ExtractPtr(event))
	return rv
}

func (t_ TextInputContext) DiscardMarkedText() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("discardMarkedText"))
}

func (t_ TextInputContext) InvalidateCharacterCoordinates() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("invalidateCharacterCoordinates"))
}

func (tc _TextInputContextClass) LocalizedNameForInputSource(inputSourceIdentifier TextInputSourceIdentifier) string {
	rv := objc.CallMethod[string](tc, objc.GetSelector("localizedNameForInputSource:"), inputSourceIdentifier)
	return rv
}

func TextInputContext_LocalizedNameForInputSource(inputSourceIdentifier TextInputSourceIdentifier) string {
	return TextInputContextClass.LocalizedNameForInputSource(inputSourceIdentifier)
}

func (tc _TextInputContextClass) CurrentInputContext() TextInputContext {
	rv := objc.CallMethod[TextInputContext](tc, objc.GetSelector("currentInputContext"))
	return rv
}

func TextInputContext_CurrentInputContext() TextInputContext {
	return TextInputContextClass.CurrentInputContext()
}

func (t_ TextInputContext) AcceptsGlyphInfo() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("acceptsGlyphInfo"))
	return rv
}

func (t_ TextInputContext) SetAcceptsGlyphInfo(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAcceptsGlyphInfo:"), value)
}

func (t_ TextInputContext) AllowedInputSourceLocales() []string {
	rv := objc.CallMethod[[]string](t_, objc.GetSelector("allowedInputSourceLocales"))
	return rv
}

func (t_ TextInputContext) SetAllowedInputSourceLocales(value []string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowedInputSourceLocales:"), value)
}

func (t_ TextInputContext) KeyboardInputSources() []TextInputSourceIdentifier {
	rv := objc.CallMethod[[]TextInputSourceIdentifier](t_, objc.GetSelector("keyboardInputSources"))
	return rv
}

func (t_ TextInputContext) SelectedKeyboardInputSource() TextInputSourceIdentifier {
	rv := objc.CallMethod[TextInputSourceIdentifier](t_, objc.GetSelector("selectedKeyboardInputSource"))
	return rv
}

func (t_ TextInputContext) SetSelectedKeyboardInputSource(value TextInputSourceIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectedKeyboardInputSource:"), value)
}
