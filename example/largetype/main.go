package main

import (
	"flag"
	"fmt"
	"runtime"
	"strings"

	"github.com/progrium/macdriver/pkg/cocoa"
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
)

func SizeText(text string, fontName string, minSize, maxSize, maxWidth float64) (rect core.NSRect, size float64) {
	t := cocoa.NSTextView_Init(core.Rect(0, 0, 0, 0))
	t.SetString(text)
	for s := minSize; s <= maxSize; s += 12 {
		t.SetFont(cocoa.Font(fontName, s))
		t.LayoutManager().EnsureLayoutForTextContainer(t.TextContainer())
		rect = t.LayoutManager().UsedRectForTextContainer(t.TextContainer())
		size = s
		if rect.Size.Width >= maxWidth {
			break
		}
	}
	return rect, size
}

func main() {
	runtime.LockOSThread()

	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		fontName := flag.String("font", "Helvetica", "font to use")
		flag.Parse()

		screen := cocoa.NSScreen_Main().Frame().Size
		text := fmt.Sprintf(" %s ", strings.Join(flag.Args(), " "))
		tr, fontSize := SizeText(text, *fontName, 70, 550, screen.Width*0.8)

		height := tr.Size.Height * 1.5
		tr.Origin.Y = (height / 2) - (tr.Size.Height / 2)
		t := cocoa.NSTextView_Init(tr)
		t.SetString(text)
		t.SetFont(cocoa.Font(*fontName, fontSize))
		t.SetEditable(false)
		t.SetImportsGraphics(false)
		t.SetDrawsBackground(false)

		c := cocoa.NSView_Init(core.Rect(0, 0, 0, 0))
		c.SetBackgroundColor(cocoa.Color(0, 0, 0, 0.75))
		c.SetWantsLayer(true)
		c.Layer().SetCornerRadius(32.0)
		c.AddSubviewPositionedRelativeTo(t, cocoa.NSWindowAbove, nil)

		tr.Size.Height = height
		tr.Origin.X = (screen.Width / 2) - (tr.Size.Width / 2)
		tr.Origin.Y = (screen.Height / 2) - (tr.Size.Height / 2)

		w := cocoa.NSWindow_Init(core.Rect(0, 0, 0, 0), cocoa.NSBorderlessWindowMask, cocoa.NSBackingStoreBuffered, false)
		w.SetContentView(c)
		w.SetTitlebarAppearsTransparent(true)
		w.SetTitleVisibility(cocoa.NSWindowTitleHidden)
		w.SetOpaque(false)
		w.SetBackgroundColor(cocoa.NSColor_Clear())
		w.SetLevel(cocoa.NSMainMenuWindowLevel + 2)
		w.SetFrameDisplay(tr, true)
		w.MakeKeyAndOrderFront(nil)

		events := make(chan cocoa.NSEvent, 64)
		go func() {
			<-events
			cocoa.NSApp().Terminate()
		}()
		cocoa.NSEvent_GlobalMonitorForEventsMatchingMask(cocoa.NSEventMaskAny, events)
	})
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}
