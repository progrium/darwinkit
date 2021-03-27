<img src="https://github.com/progrium/macdriver/raw/main/macdriver.gif" alt="MacDriver Logo">

Native Mac APIs for Golang!

[![GoDoc](https://godoc.org/github.com/progrium/macdriver?status.svg)](https://godoc.org/github.com/progrium/macdriver)
<a href="https://github.com/progrium/macdriver/actions?workflow=test"><img alt="Test workflow" src="https://img.shields.io/github/workflow/status/progrium/macdriver/Test?label=test&logo=github&style=flat-square"></a>
[![Go Report Card](https://goreportcard.com/badge/github.com/progrium/macdriver)](https://goreportcard.com/report/github.com/progrium/macdriver)
<a href="https://twitter.com/progriumHQ" title="@progriumHQ on Twitter"><img src="https://img.shields.io/badge/twitter-@progriumHQ-55acee.svg" alt="@progriumHQ on Twitter"></a>
<a href="https://github.com/progrium/macdriver/discussions" title="Project Forum"><img src="https://img.shields.io/badge/community-forum-ff69b4.svg" alt="Project Forum"></a>
<a href="https://github.com/sponsors/progrium" title="Sponsor Project"><img src="https://img.shields.io/static/v1?label=sponsor&message=%E2%9D%A4&logo=GitHub" alt="Sponsor Project" /></a>

------

MacDriver is a toolkit for working with Apple/Mac APIs and frameworks in Go. It currently has 3 "layers":

## 1. Bindings for Objective-C
The `objc` package wraps the [Objective-C runtime](https://developer.apple.com/documentation/objectivec/objective-c_runtime?language=objc) to dynamically interact with Objective-C objects and classes:

```go
cls := objc.NewClass("AppDelegate", "NSObject")
cls.AddMethod("applicationDidFinishLaunching:", func(app objc.Object) {
	fmt.Println("Launched!")
})
objc.RegisterClass(cls)

delegate := objc.Get("AppDelegate").Alloc().Init()
app := objc.Get("NSApplication").Get("sharedApplication")
app.Set("delegate:", delegate)
app.Send("run")
```

* Access any class or method you can access in Objective-C
* Common object convenience methods: Get, Set, Alloc, Init, ...
* Create and extend classes at runtime that can be used by Objective-C code
* Retain and Release methods for working with Objective-C memory management

## 2. Framework Packages
The `cocoa`, `webkit`, and `core` packages wrap `objc` with wrapper types for parts of the Apple/Mac APIs. They're being added to as needed by hand until
we can automate this process with schema data. These packages effectively let you use Apple APIs as if they were native Go libraries, letting
you write Mac applications (potentially also iOS, watchOS, etc) as Go applications:

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

* 1:1 mapping of API names with Objective-C APIs
* Cocoa types: NSApplication, NSImage, NSMenu, NSWindow, more ...
* Webkit types: WKWebView and configuration classes
* Core types: NSData, NSDictionary, NSNumber, NSPoint, NSRect, NSSize, NSString, NSURL, more ...
* Core also allows dispatching Go functions in the Cocoa run loop
* Many constants/enums are defined


### Examples
[examples/largetype](https://github.com/progrium/macdriver/blob/main/examples/largetype/main.go#L1) - A Contacts/Quicksilver-style Large Type utility in under 80 lines:
![largetype screenshot](https://pbs.twimg.com/media/EqaoO2MXIAEJNK2?format=jpg&name=large)

[examples/pomodoro](https://github.com/progrium/macdriver/blob/main/examples/pomodoro/main.go#L1) - A menu bar pomodoro timer in under 80 lines:
![pomodoro gif](https://github.com/progrium/macdriver/blob/main/examples/pomodoro/pomodoro.gif?raw=true)


[examples/topframe](https://github.com/progrium/macdriver/blob/main/examples/_topframe/main.go#L1) - A non-interactive, always-on-top webview with transparent background in 120 lines so you can draw on your
screen with HTML/JS: 
![topframe screenshot](https://pbs.twimg.com/media/EqhYDmlW8AEBC6-?format=jpg&name=large)

**NEW**: See [progrium/topframe](https://github.com/progrium/topframe) for a more fully-featured standalone version!

## 3. Bridge System
Lastly, a common case for this toolkit is not just building full native apps, but integrating Go applications
with Mac systems, like windows, native menus, status icons (systray), etc.
One-off libraries for some of these exist, but besides often limiting what you can do, 
they're also just not composable. They all want to own the main thread!

For this and other reasons, we often run the above kind of code in a separate process altogether from our
Go application. This might seem like a step backwards, but it is safer and more robust in a way. 

The `bridge` package takes advantage of this situation to create a higher-level abstraction more aligned with a potential 
cross-platform toolkit. You can declaratively describe and modify structs that can be copied to the bridge process and applied to the Objective-C
objects in a manner similar to configuration management:

```go
package main 

import (
	"os"

	"github.com/progrium/macdriver/bridge"
)

func main() {
	// start a bridge subprocess
	host := bridge.NewHost(os.Stderr)
	go host.Run()

	// create a window
	window := bridge.Window{
		Title:       "My Title",
		Size:        bridge.Size{W: 480, H: 240},
		Position:    bridge.Point{X: 200, Y: 200},
		Closable:    true,
		Minimizable: false,
		Resizable:   false,
		Borderless:  false,
		AlwaysOnTop: true,
		Background:   &bridge.Color{R: 1, G: 1, B: 1, A: 0.5},
	}
	host.Sync(&window)

	// change its title
	window.Title = "My New Title"
	host.Sync(&window)

	// destroy the window
	host.Release(&window)
}

```
This is the most WIP part of the project, but once developed further we can take this API and build a bridge
system with the same resources for Windows and Linux, making a cross-platform OS "driver". We'll see.

* Current bridge types available:
  * Window
  * StatusItem (systray)
  * Menu

## Development Notes

As far as we know, due to limitations of Go modules, we often need to add `replace` directives to our `go.mod` during development
to work against a local checkout of some dependency (like qtalk). However, these should not be versioned, so for now we encourage
you to use `git update-index --skip-worktree go.mod` on your checkout if you need to add `replace` directives. When updates need to
be checked in, `git update-index --no-skip-worktree go.mod` can be used to reverse this on your local repo to commit changes and then re-enable.

#### Generating wrappers

Eventually we can generate most of the wrapper APIs using bridgesupport and/or doc schemas. However, the number of APIs
is pretty ridiculous so there are lots of edge cases I wouldn't know how to automate yet. We can just continue to create them by hand
as needed until we have enough coverage/confidence to know how we'd generate wrappers.

## Thanks

The original `objc` and `variadic` packages were written by [Mikkel Krautz](https://github.com/mkrautz). The `variadic` package is some assembly magic to make everything possible since libobjc relies heavily on variadic function calls, which aren't possible out of the box in Cgo. 

## License

MIT
