package macdriver

import (
	"encoding/base64"
	"reflect"

	"github.com/progrium/macdriver/pkg/cocoa"
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
)

type Menu struct {
	resource

	Icon    string
	Title   string
	Tooltip string
	Items   []MenuItem
}

func (m *Menu) Apply(old, new reflect.Value, target objc.Object) (objc.Object, error) {
	if target == nil {
		menu := cocoa.NSMenu_New()
		menu.SetAutoenablesItems(false)
		for _, i := range m.Items {
			menu.AddItem(i.NSMenuItem())
		}
		target = menu.Object
	}
	return target, nil
}

type MenuItem struct {
	Title   string
	Icon    string
	Tooltip string
	Enabled bool
	Checked bool
	// TODO: submenus
	// TODO: onclick
}

func (i *MenuItem) NSMenuItem() cocoa.NSMenuItem {
	obj := cocoa.NSMenuItem_New()
	obj.SetTitle(i.Title)
	obj.SetEnabled(i.Enabled)
	obj.SetToolTip(i.Tooltip)
	if i.Checked {
		obj.SetState(cocoa.NSControlStateValueOn)
	}
	if i.Icon != "" {
		b, err := base64.StdEncoding.DecodeString(i.Icon)
		if err == nil {
			data := core.NSData_WithBytes(b, uint64(len(b)))
			obj.SetImage(cocoa.NSImage_InitWithData(data))
		}
	}
	if i.Title == "Quit" {
		obj.SetTarget(cocoa.NSApp().Delegate())
		obj.SetAction(objc.Sel("terminate:"))
	}
	return obj
}
