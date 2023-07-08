package cocoa

const (
	NSVisualEffectBlendingModeBehindWindow = 0
	NSVisualEffectBlendingModeWithinWindow = 1
)

type NSVisualEffectView struct {
	gen_NSVisualEffectView
}

func NSVisualEffectView_New() NSVisualEffectView {
	return NSVisualEffectView_Alloc().Init()
}

// effect.Set("translatesAutoresizingMaskIntoConstraints:", false)
// effect.Set("state:", 1)
// effect.Set("blendingMode:", cocoa.NSVisualEffectBlendingModeBehindWindow)
// effect.Set("material:", 2)
// effect.Set("wantsLayer:", true)
// effect.Get("layer").Set("masksToBounds:", true)
// effect.Get("layer").Set("cornerRadius:", 16.0)
