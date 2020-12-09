// Usage: largetype [-font=<fontName>] text
// TODO: replace 3 second display with close on any input
// TODO: replace call to center with something to place window at true center
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/progrium/macdriver/pkg/cocoa"
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
)

var fontName = flag.String("font", "Helvetica", "font to use")

func SizeText(text string, fontName string, minSize, maxSize, maxWidth float64) (rect core.NSRect, size float64) {
	t := cocoa.NSTextView_Init(core.Rect(0, 0, 1600, 900))
	defer t.Release()
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
	flag.Parse()

	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		text := fmt.Sprintf(" %s ", strings.Join(flag.Args(), " "))
		tr, fontSize := SizeText(text, *fontName, 70, 550, 1400)

		t := cocoa.NSTextView_Init(core.Rect(0, 0, 1600, tr.Size.Height))
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

		w := cocoa.NSWindow_Init(core.Rect(0, 0, 0, 0), cocoa.NSBorderlessWindowMask, cocoa.NSBackingStoreBuffered, false)
		w.SetTitlebarAppearsTransparent(true)
		w.SetTitleVisibility(cocoa.NSWindowTitleHidden)
		w.SetOpaque(false)
		w.SetBackgroundColor(cocoa.NSColor_Clear())

		tr.Size.Height = tr.Size.Height * 1.05
		w.SetFrameDisplay(tr, true)

		w.Center()
		w.SetContentView(c)
		w.MakeKeyAndOrderFront(nil)
	})
	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyAccessory)
	app.ActivateIgnoringOtherApps(true)
	go func() {
		<-time.After(3 * time.Second)
		os.Exit(0)
	}()
	app.Run()
}
