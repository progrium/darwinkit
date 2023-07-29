package cocoa

type NSMenu struct {
	gen_NSMenu
}

func NSMenu_New() NSMenu {
	return NSMenu_Alloc().Init()
}

func NSMenu_Init(title string) NSMenu {
	return NSMenu_Alloc().InitWithTitle(title)
}

func (menu NSMenu) Title() string {
	return menu.gen_NSMenu.Title()
}
