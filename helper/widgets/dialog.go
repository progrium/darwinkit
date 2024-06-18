package widgets

import (
	"github.com/progrium/darwinkit/helper/action"
	"github.com/progrium/darwinkit/helper/layout"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// Dialog is appkit.Panel with optional OK|CANCEL buttons
type Dialog struct {
	appkit.Panel
	content appkit.View
	ok      appkit.Button
	cancel  appkit.Button
}

// NewDialog create new Dialog
func NewDialog(width, height float64) *Dialog {
	return NewDialogWithEdgeInsets(width, height, foundation.EdgeInsets{
		Top:    10,
		Left:   10,
		Bottom: 10,
		Right:  10,
	})
}

// NewDialogWithEdgeInsets create new Dialog width edgeInsets
func NewDialogWithEdgeInsets(width, height float64, edgeInsets foundation.EdgeInsets) *Dialog {
	panel := appkit.NewPanel()
	panel.SetFrameDisplay(foundation.Rect{Size: foundation.Size{Width: width, Height: height}}, true)

	contentView := panel.ContentView()
	view := appkit.NewView()
	view.SetTranslatesAutoresizingMaskIntoConstraints(false)
	ok := appkit.NewButtonWithTitle("OK")
	ok.SetTranslatesAutoresizingMaskIntoConstraints(false)
	cancel := appkit.NewButtonWithTitle("Cancel")
	cancel.SetTranslatesAutoresizingMaskIntoConstraints(false)
	contentView.AddSubview(view)
	contentView.AddSubview(ok)
	contentView.AddSubview(cancel)

	layout.PinAnchorTo(ok.BottomAnchor(), contentView.BottomAnchor(), -edgeInsets.Bottom)
	layout.PinAnchorTo(ok.RightAnchor(), contentView.RightAnchor(), -edgeInsets.Right)
	layout.PinAnchorTo(cancel.BottomAnchor(), contentView.BottomAnchor(), -edgeInsets.Bottom)
	layout.PinAnchorTo(cancel.RightAnchor(), ok.LeftAnchor(), -10)
	layout.AlignWidth(ok, cancel)
	layout.PinAnchorTo(view.BottomAnchor(), ok.TopAnchor(), -10)
	layout.PinAnchorTo(view.LeftAnchor(), contentView.LeftAnchor(), edgeInsets.Left)
	layout.PinAnchorTo(view.TopAnchor(), contentView.TopAnchor(), edgeInsets.Top)
	layout.PinAnchorTo(view.RightAnchor(), contentView.RightAnchor(), -edgeInsets.Right)

	return &Dialog{
		Panel:   panel,
		content: view,
		ok:      ok,
		cancel:  cancel,
	}
}

// SetView set inner content view for Dialog
func (d *Dialog) SetView(view appkit.IView) {
	view.SetTranslatesAutoresizingMaskIntoConstraints(false)
	d.content.AddSubview(view)
	view.LeftAnchor().ConstraintEqualToAnchor(d.content.LeftAnchor()).SetActive(true)
	view.TopAnchor().ConstraintEqualToAnchor(d.content.TopAnchor()).SetActive(true)
	view.RightAnchor().ConstraintEqualToAnchor(d.content.RightAnchor()).SetActive(true)
	view.BottomAnchor().ConstraintEqualToAnchor(d.content.BottomAnchor()).SetActive(true)
}

// Show display dialog in non-modal mode
func (d *Dialog) Show(handle func()) {
	action.Set(d.ok, func(sender objc.Object) {
		handle()
		d.Close()
	})

	action.Set(d.cancel, func(sender objc.Object) {
		d.Close()
	})

	d.MakeKeyAndOrderFront(d.Panel)
}

// commented out due to lack of RunModalForWindow method atm

// RunModal display dialog in modal mode
// func (d *Dialog) RunModal() appkit.ModalResponse {
// 	app := appkit.Application_SharedApplication()

// 	action.Set(d.ok, func(sender objc.Object) {
// 		app.StopModalWithCode(appkit.ModalResponseOK)
// 		d.Close()
// 	})

// 	action.Set(d.cancel, func(sender objc.Object) {
// 		app.StopModalWithCode(appkit.ModalResponseCancel)
// 		d.Close()
// 	})

// 	dialogDelegate := &appkit.WindowDelegate{}
// 	dialogDelegate.SetWindowShouldClose(func(sender appkit.Window) bool {
// 		app.StopModalWithCode(appkit.ModalResponseCancel)
// 		return true
// 	})

// 	d.SetDelegate(dialogDelegate)

// 	return app.RunModalForWindow(d)
// }
