package cocoa

type NSImageView struct {
	gen_NSImageView
}

func NSImageView_New() NSImageView {
	return NSImageView_Alloc().Init()
}
