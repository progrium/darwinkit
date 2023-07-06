package cocoa

import "github.com/progrium/macdriver/core"

type NSColor struct{ gen_NSColor }

func NSColor_Init(r, g, b, a float64) NSColor {
	return NSColor_ColorWithRedGreenBlueAlpha(core.CGFloat(r), core.CGFloat(g), core.CGFloat(b), core.CGFloat(a))
}

func NSColor_Clear() NSColor {
	return NSColor_ClearColor()
}
