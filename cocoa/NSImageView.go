package cocoa

type NSImageView struct {
	gen_NSImageView
}

func NSImageView_New() NSImageView {
	return NSImageView_alloc().Init_asNSImageView()
}

func (imgView NSImageView) SetImage(img NSImageRef) {
	imgView.gen_NSImageView.SetImage_(img)
}
