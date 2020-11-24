# macdriver
Native Mac APIs in Go

```go
func main() {
	app := cocoa.NSApp_WithDidLaunch(func(notification objc.Object) {
		config := webkit.WKWebViewConfiguration_New()
		wv := webkit.WKWebView_Init(core.Rect(0, 0, 1440, 900), config)
		url := core.URL("http://progrium.com")
		req := core.NSURLRequest_Init(url)
		wv.LoadRequest(req)

		w := cocoa.NSWindow_Init(core.Rect(0, 0, 1440, 900),
			cocoa.NSClosableWindowMask|
				cocoa.NSTitledWindowMask,
			cocoa.NSBackingStoreBuffered, false)
		w.SetContentView(wv)
		w.MakeKeyAndOrderFront(w)
		w.Center()
	})
	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}
```

## Low-level Mac APIs

The low-level API lets you work directly with Apple APIs from Go. It's based on an Objective-C bridge using libobjc
for working with objects and classes, and then there are Go versions of common classes that wrap the libobjc bridge.

Since an NSApplication has its own run loop and modifications to the UI need to happen in the main thread, this basically
takes over your process. This is partly why the high-level API talks to a subprocess.

```go
w := cocoa.NSWindow_Init(cocoa.NSScreen_Main().Frame(),
    cocoa.NSBorderlessWindowMask, cocoa.NSBackingStoreBuffered, false)
w.SetContentView(wv)
w.SetBackgroundColor(cocoa.NSColor_Clear())
w.SetOpaque(false)
w.SetTitleVisibility(cocoa.NSWindowTitleHidden)
w.SetTitlebarAppearsTransparent(true)
w.SetIgnoresMouseEvents(true)
w.SetLevel(cocoa.NSMainMenuWindowLevel + 2)
w.MakeKeyAndOrderFront(w)
```

Not all APIs have wrappers and certain methods (such as those taking/returning structs) may not work. Deficiencies in the 
Objective-C bridge can be made up with CGO, letting you use C or even Objective-C to expose and call methods.

#### Generating wrappers

Eventually we can generate most of the wrapper APIs using bridgesupport and/or doc schemas. However, the number of APIs
is pretty ridiculous so there are lots of edge cases I wouldn't know how to automate yet. We can just continue to create them by hand
as needed until we have enough coverage/confidence to know how we'd generate wrappers.

## High-level Macdriver API

The high-level API lets you declaratively describe a handful of common system resources such as windows, menus, status items,
etc in your application and easily "sync" to a managed subprocess that handles direct communication with Apple APIs using the
low-level API.

```go
// create a window
window := macdriver.Window{
	Title:       "My Title",
	Size:        macdriver.Size{W: 480, H: 240},
	Position:    macdriver.Point{X: 200, Y: 200},
	Closable:    true,
	Minimizable: false,
	Resizable:   false,
    Borderless:  false,
    AlwaysOnTop: true,
	Background:   &macdriver.Color{R: 1, G: 1, B: 1, A: 0.5},
}
macdriver.Sync(peer, &window)

// change its title
window.Title = "My New Title"
macdriver.Sync(peer, &window)

// destroy the window
macdriver.Release(peer, &window)
```

The high-level API is meant to be platform agnostic, so you can imagine windriver and linuxdriver equivalent projects. It's also
based on a non-Go specific communication protocol, so this API could be exposed to other languages.

## Thanks

The original `objc` and `variadic` packages were written by Mikkel Krautz. The `variadic` package makes everything possible since
libobjc relies heavily on variadic function calls, which weren't possible in Cgo. 

## License

MIT