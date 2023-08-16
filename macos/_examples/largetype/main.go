package main

import (
	"flag"
	"fmt"
	"runtime"
	"strings"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	fontName := flag.String("font", "Helvetica", "font to use")
	flag.Parse()
	text := strings.Join(flag.Args(), " ")
	if text == "" {
		text = "Hello world"
	}

	app := appkit.Application_SharedApplication()
	screen := appkit.Screen_MainScreen().Frame().Size
	text = fmt.Sprintf(" %s ", text)
	fmt.Println(text)

	tr, fontSize := func() (rect foundation.Rect, size float64) {
		t := appkit.TextView_InitWithFrame(rectOf(0, 0, 0, 0))
		t.SetString(text)
		for s := 70.0; s <= 550; s += 12 {
			t.SetFont(appkit.Font_FontWithNameSize(*fontName, s))
			t.LayoutManager().EnsureLayoutForTextContainer(t.TextContainer())
			rect = t.LayoutManager().UsedRectForTextContainer(t.TextContainer())
			size = s
			if rect.Size.Width >= screen.Width*0.8 {
				break
			}
		}
		return rect, size
	}()

	height := tr.Size.Height * 1.5
	tr.Origin.Y = (height / 2) - (tr.Size.Height / 2)
	t := appkit.TextView_InitWithFrame(tr)
	t.SetString(text)
	t.SetFont(appkit.Font_FontWithNameSize(*fontName, fontSize))
	t.SetEditable(false)
	t.SetImportsGraphics(false)
	t.SetDrawsBackground(false)

	c := appkit.View_InitWithFrame(rectOf(0, 0, 0, 0))
	// deprecated...
	// c.SetBackgroundColor(appkit.Color_ColorWithRedGreenBlueAlpha(0, 0, 0, 0.75))
	c.SetWantsLayer(true)
	c.Layer().SetCornerRadius(32.0)
	c.AddSubviewPositionedRelativeTo(t, appkit.WindowAbove, nil)

	tr.Size.Height = height
	tr.Origin.X = (screen.Width / 2) - (tr.Size.Width / 2)
	tr.Origin.Y = (screen.Height / 2) - (tr.Size.Height / 2)

	w := appkit.Window_InitWithContentRectStyleMaskBackingDefer(rectOf(0, 0, 0, 0),
		appkit.WindowStyleMaskBorderless, appkit.BackingStoreBuffered, false)
	w.SetContentView(c)
	w.SetTitlebarAppearsTransparent(true)
	w.SetTitleVisibility(appkit.WindowTitleHidden)
	w.SetOpaque(false)
	w.SetBackgroundColor(appkit.Color_ClearColor())
	w.SetLevel(appkit.MainMenuWindowLevel + 2)
	w.SetFrameDisplay(tr, true)
	w.MakeKeyAndOrderFront(nil)

	appkit.Event_AddGlobalMonitorForEventsMatchingMaskHandler(appkit.EventMaskAny, func(event appkit.Event) {
		appkit.Application_SharedApplication().Terminate(nil)
	})

	ad := &appkit.ApplicationDelegate{}
	ad.SetApplicationDidFinishLaunching(func(foundation.Notification) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)
	})
	app.SetDelegate(ad)

	app.Run()
}

func rectOf(x, y, width, height float64) foundation.Rect {
	return foundation.Rect{Origin: foundation.Point{X: x, Y: y}, Size: foundation.Size{Width: width, Height: height}}
}
