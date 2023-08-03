// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IPasteboardWriting interface {
	// required
	WritableTypesForPasteboard(pasteboard Pasteboard) []PasteboardType
	ImplementsWritingOptionsForTypePasteboard() bool
	// optional
	WritingOptionsForTypePasteboard(type_ PasteboardType, pasteboard Pasteboard) PasteboardWritingOptions
	// required
	PasteboardPropertyListForType(type_ PasteboardType) objc.IObject
}

type PasteboardWritingWrapper struct {
	objc.Object
}

func (p_ PasteboardWritingWrapper) WritableTypesForPasteboard(pasteboard IPasteboard) []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](p_, objc.GetSelector("writableTypesForPasteboard:"), objc.ExtractPtr(pasteboard))
	return rv
}

func (p_ PasteboardWritingWrapper) ImplementsWritingOptionsForTypePasteboard() bool {
	return p_.RespondsToSelector(objc.GetSelector("writingOptionsForType:pasteboard:"))
}

func (p_ PasteboardWritingWrapper) WritingOptionsForTypePasteboard(type_ PasteboardType, pasteboard IPasteboard) PasteboardWritingOptions {
	rv := objc.CallMethod[PasteboardWritingOptions](p_, objc.GetSelector("writingOptionsForType:pasteboard:"), type_, objc.ExtractPtr(pasteboard))
	return rv
}

func (p_ PasteboardWritingWrapper) PasteboardPropertyListForType(type_ PasteboardType) objc.Object {
	rv := objc.CallMethod[objc.Object](p_, objc.GetSelector("pasteboardPropertyListForType:"), type_)
	return rv
}
