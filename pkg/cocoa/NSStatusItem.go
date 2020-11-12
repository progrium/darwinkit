package cocoa

import "github.com/progrium/macdriver/pkg/objc"

type NSStatusItem struct {
	objc.Object
}

// statusBarItem.Send("button").Send("setTitle:", core.String("Hello world"))
// 	statusBarItem.Send("setTarget:", delegate)
// 	statusBarItem.Send("setAction:", objc.Sel("foobar:"))
