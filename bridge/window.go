package bridge

import (
	"encoding/base64"

	"github.com/progrium/macdriver/bridge/resource"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"github.com/progrium/macdriver/webkit"
)

type Window struct {
	*resource.Handle // win:

	Title        string
	Position     Point
	Size         Size
	Closable     bool
	Minimizable  bool
	Resizable    bool
	Background   *Color
	Borderless   bool
	CornerRadius float64
	AlwaysOnTop  bool
	IgnoreMouse  bool
	Center       bool
	URL          string
	Image        string

	webview *webkit.WKWebView
	image   *cocoa.NSImage
}

func (w *Window) Discard(target objc.Object) error {
	obj := cocoa.NSWindow{Object: target}
	obj.Close()
	return nil
}

func (w *Window) Apply(target objc.Object) (objc.Object, error) {
	frame := core.Rect(w.Position.X, w.Position.Y, w.Size.W, w.Size.H)
	if target == nil {
		win := cocoa.NSWindow_Init(core.Rect(0, 0, 0, 0), cocoa.NSTitledWindowMask, cocoa.NSBackingStoreBuffered, false)
		win.Retain()
		win.MakeKeyAndOrderFront(nil)
		target = win.Object
		if w.Center {
			screenRect := cocoa.NSScreen_Main().Frame()
			frame = core.Rect(
				(screenRect.Size.Width/2)-(w.Size.W/2),
				(screenRect.Size.Height/2)-(w.Size.H/2),
				w.Size.W,
				w.Size.H,
			)
		}
	}
	obj := cocoa.NSWindow{Object: target}

	if w.URL != "" && w.webview == nil {
		config := webkit.WKWebViewConfiguration_New()
		config.Preferences().SetValueForKey(core.True, core.String("developerExtrasEnabled"))

		wv := webkit.WKWebView_Init(core.Rect(0, 0, w.Size.W, w.Size.H), config)
		w.webview = &wv

		req := core.NSURLRequest_Init(core.URL(w.URL))
		wv.LoadRequest(req)
	}

	if w.Image != "" && w.image == nil {
		b, err := base64.StdEncoding.DecodeString(w.Image)
		if err != nil {
			return nil, err
		}
		data := core.NSData_WithBytes(b, uint64(len(b)))
		image := cocoa.NSImage_InitWithData(data)
		w.image = &image
	}

	mask := cocoa.NSTitledWindowMask
	needsTitleBar := w.Closable || w.Minimizable
	if w.Borderless {
		if !needsTitleBar {
			mask = cocoa.NSBorderlessWindowMask
		}
		mask = mask | cocoa.NSFullSizeContentViewWindowMask
	}
	if w.Closable {
		mask = mask | cocoa.NSClosableWindowMask
	}
	if w.Minimizable {
		mask = mask | cocoa.NSMiniaturizableWindowMask
	}
	if w.Resizable {
		mask = mask | cocoa.NSResizableWindowMask
	}
	obj.SetStyleMask(mask)

	if w.Title != "" {
		obj.SetTitle(w.Title)
	} else {
		obj.SetMovableByWindowBackground(true)
		obj.SetTitlebarAppearsTransparent(true)
	}

	if w.Borderless && w.CornerRadius > 0 {
		obj.SetBackgroundColor(cocoa.NSColor_Clear())
		obj.SetOpaque(false)
		v := cocoa.NSView_Init(core.Rect(0, 0, 0, 0))
		if w.Background != nil {
			v.SetBackgroundColor(w.Background.NSColor())
		}
		v.SetWantsLayer(true)
		v.Layer().SetCornerRadius(w.CornerRadius)

		if w.webview != nil {
			v.AddSubviewPositionedRelativeTo(*w.webview, cocoa.NSWindowAbove, nil)
		}
		obj.SetContentView(v)
	} else {
		if w.Background != nil {
			obj.SetBackgroundColor(w.Background.NSColor())
			obj.SetOpaque(w.Background.A == 1)
		}
		if w.webview != nil {
			obj.SetContentView(*w.webview)
		}
	}

	if w.webview != nil && w.Background != nil && w.Background.A == 0 {
		w.webview.SetOpaque(false)
		w.webview.SetBackgroundColor(cocoa.NSColor_Clear())
		w.webview.SetValueForKey(core.False, core.String("drawsBackground"))
	}

	if w.image != nil {
		obj.ContentView().SetWantsLayer(true)
		obj.ContentView().Layer().SetContents(w.image)
	}

	if w.AlwaysOnTop {
		obj.SetLevel(cocoa.NSMainMenuWindowLevel)
	}

	if w.IgnoreMouse {
		obj.SetIgnoresMouseEvents(true)
	}

	obj.SetFrameDisplay(frame, true)
	return target, nil
}
