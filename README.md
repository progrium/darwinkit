<svg 
 xmlns="http://www.w3.org/2000/svg"
 xmlns:xlink="http://www.w3.org/1999/xlink"
 width="516px" height="108px">
<image  x="0px" y="0px" width="516px" height="108px"  xlink:href="data:img/png;base64,iVBORw0KGgoAAAANSUhEUgAAAgQAAABsCAQAAAC/M2IDAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAHdElNRQfkDB0KHTDyoEgRAAADmklEQVR42u3dwW3bMBQG4OfC6EI99JIBskfRCTpHJ8gpW2SAXtpTlwiQAXrKJT1UggtbtkiJFKXo+4DARiLZktD+IJ8o8vAjgK24e+veHP69lPr/+6H1iQHtCQJAEAARx9YHAIw7rw2UpkUACAJAEAChRgBJHt/mf8YUD93rU/+Ls+P4U+h7tAgAQQBEHD+2PoLGXlsfAKyAGgGs0EVtoDJdA0AQAIIACEEAhCAAQhAAIQiAMI4AVmXp8QO93Y8sjDC6EHQNAEEAqBFAVX2f/2vl77nvXqfWFrQIAEEARBy/DE3BlDZl8vQ9V+QlIp6799+zp6P6trLzHTqDtR1jaxb5G6JGACvQavxAT9cAEASAIABCEABxrVh4a8HFRgs9LOFUXx++f9D//fTX/7dTnWe7tAgAQQAYRwBVjT1jUGr8wNz9tQiAkRbBOy4MzjFcVDy9b1c2VLBkGi0CQBAAggCIZe4apFYa8vu3+TWMRn3o1AFKw8Z6/rf3v7b39EeWb59N+pHln23d67hnWgSAIAAEARCCAIh1DTE+lXoOidvV+45FpJbS+u22UuyaUiIc2jv1fPOu48/4vJHruKThIOgv1O3pScvcDSgzenGJ75hh7AHn3O1K7NXuGlxKPe6613FJrecoPKdrAAgCQBAAIQiAGLtrUKa62r5cs/Iq8VbuBsxTv2y3j+tYhxYBIAiAdQ0ogt25z9y+1rgDLQKgYosgddH09qXEVXp/ha82Z3T5rZZFH6JFAKgRQAtj6x2ce8jcPpcWAaBFsF9DT+jdfgR4+iyHrJ0WASAIAEEAhCAAYi/FwtTBTYnWOv1VScueY/qSKdSxjyBYxHb/2aauFPQezrW2/n5/7jiBsc/LfcbgPnM/XQNAEACCAAg1Aiiqrw08XPl9LefzGuTWFA6fplfUyzxofHvBlNILl1x83ktEPHfv8yvlpdfimVaEq/HpZYYTL7Mset5nTHsM+THzTEoFQWqxcG4Q6BoAggBQI4AqzrsCc+cTGJvbcO5chloEQBx+tz6Cxl67H7glt1h4rvYMQ1oEwGxqBLCA2rcP59IiAAQBIAiAUCOAVetrC7nzC+TSIgAEASAIgBAEQEQc9z68du/nDxFaBEAIAiAEARCCAAhBAIQgACLi8Kv1ETTm9iFbcNfPkNRN1D5tUvbrtAgAQQAIAiAMMYZt6Bdx62oFd4VrBVoEgCAABAEQggAIQQCEIABCEAAR8Rdok6bPB+U8FQAAAABJRU5ErkJggg==" />
</svg>

Native Mac APIs for Go!

[![GoDoc](https://godoc.org/github.com/progrium/macdriver?status.svg)](https://godoc.org/github.com/progrium/macdriver)
[![Go Report Card](https://goreportcard.com/badge/github.com/progrium/macdriver)](https://goreportcard.com/report/github.com/progrium/macdriver)
<a href="https://twitter.com/progriumHQ" title="@progriumHQ on Twitter"><img src="https://img.shields.io/badge/twitter-@progriumHQ-55acee.svg" alt="@progriumHQ on Twitter"></a>
<a href="https://github.com/progrium/macdriver/discussions" title="Project Forum"><img src="https://img.shields.io/badge/community-forum-ff69b4.svg" alt="Project Forum"></a>

------


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

The high-level API lets you declaratively describe a handful of common system resources such as windows, menus, systray items,
etc in your application and then "sync" them to a managed subprocess that reconciles state with Apple APIs using the
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