package cocoa

import (
	"github.com/progrium/macdriver/core"
)

type NSMenu struct {
	gen_NSMenu
}

func NSMenu_New() NSMenu {
	return NSMenu_alloc().Init_asNSMenu()
}

func NSMenu_Init(title string) NSMenu {
	return NSMenu_alloc().InitWithTitle__asNSMenu(core.String(title))
}

func (menu NSMenu) SetTitle(title string) {
	menu.SetTitle_(core.String(title))
}

func (menu NSMenu) Title() string {
	return menu.gen_NSMenu.Title().String()
}

func (menu NSMenu) AddItem(item NSMenuItemRef) {
	menu.AddItem_(item)
}

func (menu NSMenu) RemoveItem(item NSMenuItemRef) {
	menu.RemoveItem_(item)
}

func (menu NSMenu) SetAutoenablesItems(b bool) {
	menu.SetAutoenablesItems_(b)
}
