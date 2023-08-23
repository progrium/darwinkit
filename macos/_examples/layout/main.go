package main

import (
	"fmt"

	"github.com/progrium/macdriver/helper/action"
	"github.com/progrium/macdriver/helper/layout"
	"github.com/progrium/macdriver/helper/widgets"
	"github.com/progrium/macdriver/macos"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

func main() {
	macos.RunApp(launched)
}

func launched(app appkit.Application, delegate *appkit.ApplicationDelegate) {
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)

	w := appkit.NewWindowWithSize(600, 400)
	objc.Retain(&w)
	w.SetTitle("Test Layout")

	label := appkit.NewLabel("label")
	mdButton := appkit.NewButtonWithTitle("modal dialog")
	dButton := appkit.NewButtonWithTitle("dialog")
	textView := appkit.TextViewClass.ScrollableTextView()

	action.Set(mdButton, func(sender objc.Object) {
		d := widgets.NewDialog(400, 300)
		d.SetView(appkit.NewLabel("test modal dialog"))
		// if d.RunModal() == appkit.ModalResponseOK {
		// 	fmt.Println("ok!")
		// }
	})

	action.Set(dButton, func(sender objc.Object) {
		d := widgets.NewDialog(400, 300)
		d.SetView(appkit.NewLabel("test dialog"))
		d.Center()
		d.Show(func() {
			fmt.Println("ok!")
		})
	})

	gridView := appkit.NewGridView()
	for i := 0; i < 3; i++ {
		var views []appkit.IView
		for j := 0; j < 4; j++ {
			label := appkit.NewLabel(fmt.Sprintf("label-%v-%v", i, j))
			views = append(views, label)
		}
		gridView.AddRowWithViews(views)
	}
	gridView.SetContentHuggingPriorityForOrientation(appkit.LayoutPriorityDefaultHigh, appkit.LayoutConstraintOrientationHorizontal)
	gridView.ColumnAtIndex(0).SetXPlacement(appkit.GridCellPlacementTrailing)
	gridView.SetRowAlignment(appkit.GridRowAlignmentLastBaseline)

	stackView := appkit.StackView_StackViewWithViews([]appkit.IView{label, mdButton, dButton, textView, gridView})
	stackView.SetOrientation(appkit.UserInterfaceLayoutOrientationVertical)
	stackView.SetDistribution(appkit.StackViewDistributionFillEqually)
	stackView.SetAlignment(appkit.LayoutAttributeCenterX)
	stackView.SetSpacing(10)

	w.ContentView().AddSubview(stackView)
	layout.PinEdgesToSuperView(stackView, foundation.EdgeInsets{Top: 10, Bottom: 10, Left: 20, Right: 20})

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
}
