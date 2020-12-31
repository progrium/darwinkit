// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSTextAlignment uint

const (
	NSTextAlignmentLeft      NSTextAlignment = 0
	NSTextAlignmentRight     NSTextAlignment = 1
	NSTextAlignmentCenter    NSTextAlignment = 2
	NSTextAlignmentJustified NSTextAlignment = 3
	NSTextAlignmentNatural   NSTextAlignment = 4
)

type NSTextView struct {
	NSView
}

func NSTextView_Init(frame core.NSRect) NSTextView {
	return NSTextView{NSView{objc.Get("NSTextView").Alloc().Send("initWithFrame:", frame)}}
}

func (v NSTextView) String() string {
	return v.Get("string").String()
}

func (v NSTextView) SetString(s string) {
	v.Set("string:", core.String(s))
}

func (v NSTextView) Selectable() bool {
	return v.Get("selectable").Bool()
}

func (v NSTextView) SetSelectable(b bool) {
	v.Set("selectable:", b)
}

func (v NSTextView) RichText() bool {
	return v.Get("richText").Bool()
}

func (v NSTextView) SetRichText(b bool) {
	v.Set("richText:", b)
}

func (v NSTextView) Editable() bool {
	return v.Get("editable").Bool()
}

func (v NSTextView) SetEditable(b bool) {
	v.Set("editable:", b)
}

func (v NSTextView) FieldEditor() bool {
	return v.Get("fieldEditor").Bool()
}

func (v NSTextView) SetFieldEditor(b bool) {
	v.Set("fieldEditor:", b)
}

func (v NSTextView) ImportsGraphics() bool {
	return v.Get("importsGraphics").Bool()
}

func (v NSTextView) SetImportsGraphics(b bool) {
	v.Set("importsGraphics:", b)
}

func (v NSTextView) DrawsBackground() bool {
	return v.Get("drawsBackground").Bool()
}

func (v NSTextView) SetDrawsBackground(b bool) {
	v.Set("drawsBackground:", b)
}

func (v NSTextView) Font() NSFont {
	return NSFont{v.Get("font")}
}

func (v NSTextView) SetFont(f NSFont) {
	v.Set("font:", f)
}

func (v NSTextView) Alignment() NSTextAlignment {
	return NSTextAlignment(v.Get("alignment").Int())
}

func (v NSTextView) SetAlignment(d NSTextAlignment) {
	v.Set("alignment:", d)
}

func (v NSTextView) TextContainer() NSTextContainer {
	return NSTextContainer{v.Get("textContainer")}
}

func (v NSTextView) SetTextContainer(tc NSTextContainer) {
	v.Set("textContainer:", tc)
}

func (v NSTextView) LayoutManager() NSLayoutManager {
	return NSLayoutManager{v.Get("layoutManager")}
}

func (v NSTextView) SetLayoutManager(lm NSLayoutManager) {
	v.Set("layoutManager:", lm)
}
