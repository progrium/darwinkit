package macdriver

import (
	"encoding/base64"
	"fmt"
	"os"
	"reflect"

	"github.com/manifold/qtalk/golang/rpc"
	"github.com/progrium/macdriver/pkg/cocoa"
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
	"github.com/rs/xid"
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

type rpc_FuncExport struct {
	Ptr    string `json:"$fnptr" mapstructure:"$fnptr"`
	Caller rpc.Caller
	fn     interface{}
}

func (e *rpc_FuncExport) Call(args, reply interface{}) error {
	_, err := e.Caller.Call("Invoke", e.Ptr, reply)
	return err
}

func (e *rpc_FuncExport) DelegateCallback() (objc.Object, objc.Selector) {
	ee := *e
	return core.Callback(func(o objc.Object) {
		err := ee.Call(nil, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "callback: %v\n", err)
		}
	})
}

var exportedFuncs map[string]rpc_FuncExport

func ExportFunc(fn interface{}) *rpc_FuncExport {
	if exportedFuncs == nil {
		exportedFuncs = make(map[string]rpc_FuncExport)
	}
	id := xid.New().String()
	ef := rpc_FuncExport{
		Ptr: id,
		fn:  fn,
	}
	exportedFuncs[id] = ef
	return &ef
}

type MenuItem struct {
	Title   string
	Icon    string
	Tooltip string
	Enabled bool
	Checked bool
	OnClick *rpc_FuncExport
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
		obj.SetTarget(cocoa.NSApp())
		obj.SetAction(objc.Sel("terminate:"))
	}
	if i.OnClick != nil && i.OnClick.Caller != nil {
		t, sel := i.OnClick.DelegateCallback()
		obj.SetTarget(t)
		obj.SetAction(sel)
	}
	return obj
}
