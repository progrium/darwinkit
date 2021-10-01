package cocoa

type NSViewController struct{ gen_NSViewController }

func NSViewController_New() NSViewController {
	return NSViewController_alloc().Init_asNSViewController()
}

func (c NSViewController) SetView(v NSView) {
	c.Set("view:", v)
}
