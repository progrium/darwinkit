package bridge

import (
	"encoding/base64"
	"reflect"

	"github.com/progrium/macdriver/pkg/cocoa"
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
)

type StatusItem struct {
	resource

	Icon string
	Text string
	Menu *Menu
}

func (s *StatusItem) Apply(old, new reflect.Value, target objc.Object) (objc.Object, error) {
	obj := cocoa.NSStatusItem{Object: target}
	if target == nil {
		obj = cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		obj.Retain()
		target = obj.Object
	}
	obj.Button().SetTitle(s.Text)
	if s.Icon != "" {
		b, err := base64.StdEncoding.DecodeString(s.Icon)
		if err != nil {
			return nil, err
		}
		data := core.NSData_WithBytes(b, uint64(len(b)))
		image := cocoa.NSImage_InitWithData(data)
		image.SetSize(core.Size(16.0, 16.0))
		image.SetTemplate(true)
		obj.Button().SetImage(image)
		if s.Text != "" {
			obj.Button().SetImagePosition(cocoa.NSImageLeft)
		} else {
			obj.Button().SetImagePosition(cocoa.NSImageOnly)
		}

	}
	if s.Menu != nil {
		menu, err := s.Menu.Apply(reflect.Value{}, reflect.Value{}, nil)
		if err != nil {
			return nil, err
		}
		obj.SetMenu(cocoa.NSMenu{Object: menu})
	}
	return target, nil
}
