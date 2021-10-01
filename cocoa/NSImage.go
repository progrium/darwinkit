package cocoa

import (
	"github.com/progrium/macdriver/core"
)

type NSImage struct {
	gen_NSImage
}

func NSImage_InitWithData(data core.NSDataRef) NSImage {
	return NSImage_alloc().InitWithData__asNSImage(data)
}

func NSImage_InitWithURL(url core.NSURL) NSImage {
	return NSImage_alloc().InitWithContentsOfURL__asNSImage(url)
}

func NSImage_ImageNamed(name string) NSImage {
	return NSImage_fromRef(NSImage_alloc().Send("imageNamed:", core.String(name)))
}

func (i NSImage) SetSize(size core.NSSize) {
	i.SetSize_(size)
}

func (i NSImage) SetTemplate(b bool) {
	i.SetTemplate_(b)
}

func (i NSImage) SetValueForKey(value, key interface{}) {
	i.Send("setValue:forKey:", value, key)
}
