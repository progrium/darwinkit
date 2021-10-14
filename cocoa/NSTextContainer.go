package cocoa

type NSTextContainer struct {
	gen_NSTextContainer
}

func (tc NSTextContainer) SetHeightTracksTextView(b bool) {
	tc.SetHeightTracksTextView_(b)
}
