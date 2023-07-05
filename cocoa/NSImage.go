package cocoa

import (
	"github.com/progrium/macdriver/core"
)

type NSImage struct {
	gen_NSImage
}

func NSImage_InitWithData(data core.NSDataRef) NSImage {
	return NSImage_alloc().InitWithData_asNSImage(data)
}

func NSImage_InitWithURL(url core.NSURL) NSImage {
	return NSImage_alloc().InitWithContentsOfURL_asNSImage(url)
}

func NSImage_ImageNamed(name string) NSImage {
	return NSImage_fromRef(NSImage_alloc().Send("imageNamed:", core.String(name)))
}

func (i NSImage) SetValueForKey(value, key interface{}) {
	i.Send("setValue:forKey:", value, key)
}
