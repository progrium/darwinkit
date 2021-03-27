package cocoa

import (
	"github.com/progrium/macdriver/objc"
)

type NSImageView struct {
	objc.Object
}

func NSImageView_New() NSImageView {
	return NSImageView{objc.Get("NSImageView").Alloc().Init()}
}

func (imgView NSImageView) SetImage(img objc.Object) {
	imgView.Send("setImage:", img)
}

func (imgView NSImageView) Image() NSImage {
	return NSImage{imgView.Send("image")}
}
