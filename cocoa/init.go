package cocoa

/*
#cgo LDFLAGS: -framework AppKit
*/
import "C"

func Color(r, g, b, a float64) NSColor {
	return NSColor_Init(r, g, b, a)
}

func Font(fontName string, size float64) NSFont {
	return NSFont_Init(fontName, size)
}
