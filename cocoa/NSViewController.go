package cocoa

type NSViewController struct{ gen_NSViewController }

func NSViewController_New() NSViewController {
	return NSViewController_Alloc().Init()
}

func (c NSViewController) SetView(v NSView) {
	c.Set("view:", v)
}
