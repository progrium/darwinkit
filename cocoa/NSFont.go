package cocoa

import (
	"github.com/progrium/macdriver/core"
)

type NSFont struct{ gen_NSFont }

func NSFont_Init(fontName string, size float64) NSFont {
	return NSFont_fontWithName_size_(
		core.String(fontName),
		core.CGFloat(size),
	)
}
