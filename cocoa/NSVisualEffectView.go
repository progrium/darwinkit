package cocoa

import "github.com/progrium/macdriver/objc"

const (
	NSVisualEffectBlendingModeBehindWindow = 0
	NSVisualEffectBlendingModeWithinWindow = 1
)

type NSVisualEffectView struct {
	objc.Object
}

func NSVisualEffectView_New() NSVisualEffectView {
	return NSVisualEffectView{objc.Get("NSVisualEffectView").Alloc().Init()}
}

// effect.Set("translatesAutoresizingMaskIntoConstraints:", false)
// effect.Set("state:", 1)
// effect.Set("blendingMode:", cocoa.NSVisualEffectBlendingModeBehindWindow)
// effect.Set("material:", 2)
// effect.Set("wantsLayer:", true)
// effect.Get("layer").Set("masksToBounds:", true)
// effect.Get("layer").Set("cornerRadius:", 16.0)
