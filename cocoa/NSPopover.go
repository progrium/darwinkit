package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSPopover struct{ gen_NSPopover }

const (
	NSPopoverBehaviorApplicationDefined = 0
	NSPopoverBehaviorTransient          = 1
	NSPopoverBehaviorSemitransient      = 2
)

func NSPopover_Init() NSPopover {
	return NSPopover_Alloc().Init()
}

func (p NSPopover) SetContentSize(s core.NSSize) {
	p.Send("setContentSize:", s)
}

func (p NSPopover) SetBehavior(b int) {
	p.gen_NSPopover.SetBehavior(core.NSInteger(b))
}

func (p NSPopover) SetDelegate(delegate objc.Object) {
	p.Send("setDelegate:", delegate)
}

func (p NSPopover) SetViewController(controller objc.Object) {
	p.Send("setContentViewController:", controller)
}
