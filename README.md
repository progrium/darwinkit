<img alt="DarwinKit Logo" src="https://github.com/progrium/darwinkit/raw/main/darwinkit.png" style="height:125px;" />

Native Apple APIs for Golang!




[![GoDoc](https://godoc.org/github.com/progrium/darwinkit?status.svg)](https://pkg.go.dev/github.com/progrium/darwinkit@main)
[![Go Report Card](https://goreportcard.com/badge/github.com/progrium/darwinkit)](https://goreportcard.com/report/github.com/progrium/darwinkit)
<a href="https://twitter.com/progrium" title="@progrium on Twitter"><img src="https://img.shields.io/badge/twitter-@progrium-55acee.svg" alt="@progrium on Twitter"></a>
<a href="https://github.com/progrium/darwinkit/discussions" title="Project Forum"><img src="https://img.shields.io/badge/community-forum-ff69b4.svg" alt="Project Forum"></a>
<a href="https://github.com/sponsors/darwinkitdev" title="Sponsor Project"><img src="https://img.shields.io/static/v1?label=sponsor&message=%E2%9D%A4&logo=GitHub" alt="Sponsor Project" /></a>

> [!IMPORTANT]
> July 11, 2024: **MacDriver is now DarwinKit!** You can also now get high-quality example starter apps when you [support the project as a sponsor](https://github.com/sponsors/darwinkitdev).

------

![ScanDrop Demo](https://github.com/progrium/darwinkit/assets/1813419/e66e4ca1-af6b-4952-9e53-f10ab436993f)
![SnowScape Demo](https://github.com/progrium/darwinkit/assets/1813419/3b04b849-4292-41c5-8563-7735ff46c745)
![MenuSpacer Demo](https://github.com/progrium/darwinkit/assets/1813419/ed86724d-f1fc-42e4-b9c9-0112fffd86fc)
![ClipTrail Demo](https://github.com/progrium/darwinkit/assets/1813419/c018851d-5af7-4cf6-8e21-213c28752438)


DarwinKit lets you work with [supported Apple frameworks](https://pkg.go.dev/github.com/progrium/darwinkit/macos@main#section-directories) and build native applications using Go. With XCode and Go 1.18+ installed, you can write this program in a `main.go` file:

```go
package main

import (
	"github.com/progrium/darwinkit/objc"
	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/webkit"
)

func main() {
	// runs macOS application event loop with a callback on success
	macos.RunApp(func(app appkit.Application, delegate *appkit.ApplicationDelegate) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		url := foundation.URL_URLWithString("http://progrium.com")
		req := foundation.NewURLRequestWithURL(url)
		frame := foundation.Rect{Size: foundation.Size{1440, 900}}

		config := webkit.NewWebViewConfiguration()
		wv := webkit.NewWebViewWithFrameConfiguration(frame, config)
		wv.LoadRequest(req)

		w := appkit.NewWindowWithContentRectStyleMaskBackingDefer(frame,
			appkit.ClosableWindowMask|appkit.TitledWindowMask,
			appkit.BackingStoreBuffered, false)
		objc.Retain(&w)
		w.SetContentView(wv)
		w.MakeKeyAndOrderFront(w)
		w.Center()

		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			return true
		})
	})
}

```

Then in this directory run:

```
go mod init helloworld
go get github.com/progrium/darwinkit@main
go run main.go
```

This may take a moment the first time, but once the window pops up you just made a macOS program without using XCode or Objective-C. Run `go build` to get an executable. 

Although currently outside the scope of this project, if you wanted you could put this executable [into an Application bundle](https://stackoverflow.com/a/3251285). You could even add [entitlements](https://developer.apple.com/documentation/bundleresources/entitlements?language=objc), then [sign and notarize](https://developer.apple.com/support/code-signing/) this bundle or executable to let others run it. It could theoretically even be put on the App Store. It could *theoretically* be put on an iOS, tvOS, or watchOS device, though you would have to use different platform specific frameworks.

### Caveats

* You still need to know or learn how Apple frameworks work, so you'll have to use Apple documentation and understand how to translate Objective-C example code to the equivalent Go with DarwinKit.
* Your programs link against the actual Apple frameworks using [cgo](https://pkg.go.dev/cmd/cgo), so XCode needs to be installed for the framework headers.
* You will be using two memory management systems. Framework objects are managed by [Objective-C memory management](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/MemoryMgmt/Articles/MemoryMgmt.html#//apple_ref/doc/uid/10000011-SW1), so be sure to read our docs on [memory management](docs/memorymanagement.md) with DarwinKit.
* Exceptions in frameworks will segfault, giving you both an Objective-C stacktrace and a Go panic stacktrace. You will be debugging a hybrid Go and Objective-C program.
* Goroutines that interact with GUI objects need to [dispatch](https://pkg.go.dev/github.com/progrium/darwinkit@main/dispatch) operations on the main thread otherwise it will segfault.

This is all tenable for simple programs, but these are the reasons we don't *recommend* large/complex programs using DarwinKit.

## Examples

There are basic usage examples [here in this repo](./macos/_examples), but there are higher-quality full utilities you can use as examples or starter templates [when you sponsor](https://github.com/sponsors/darwinkitdev).

![LargeType Demo](https://github.com/progrium/darwinkit/assets/1813419/b54a8015-dfef-40ce-bc56-f75cb8455306)
![ScanDrop Demo](https://github.com/progrium/darwinkit/assets/1813419/e66e4ca1-af6b-4952-9e53-f10ab436993f)
![SnowScape Demo](https://github.com/progrium/darwinkit/assets/1813419/3b04b849-4292-41c5-8563-7735ff46c745)
![LiveTemp Demo](https://github.com/progrium/darwinkit/assets/1813419/c57cd196-3ff6-4068-bd0a-05b9283d199d)
![MenuSpacer Demo](https://github.com/progrium/darwinkit/assets/1813419/ed86724d-f1fc-42e4-b9c9-0112fffd86fc)
![ClipTrail Demo](https://github.com/progrium/darwinkit/assets/1813419/c018851d-5af7-4cf6-8e21-213c28752438)
![FocusTimer Demo](https://github.com/progrium/darwinkit/assets/1813419/9e14fe2f-88a9-4b89-9f9c-a305fb103d4c)


## How it works

<details>
<summary>Brief background on Objective-C</summary>
	
Ever since acquiring NeXT Computer in the 90s, Apple has used [NeXTSTEP](https://en.wikipedia.org/wiki/NeXTSTEP) as the basis of their software stack, which is written in Objective-C. Unlike most systems languages with object orientation, Objective-C implements OOP as a runtime library. In fact, Objective-C is just C with the weird OOP specific syntax rewritten into C calls to [libobjc](https://developer.apple.com/documentation/objectivec/objective-c_runtime?language=objc), which is a normal C library implementing an object runtime. This runtime could be used to bring OOP to any language that can make calls to C code. It also lets you interact with objects and classes registered by other libraries, such as the Apple frameworks. 

</details>

At the heart of DarwinKit is a package wrapping the Objective-C runtime using [cgo](https://pkg.go.dev/cmd/cgo) and [libffi](https://github.com/libffi/libffi). This is actually all you need to interact with Objective-C objects and classes, it'll just look like this:

```go
app := objc.Call[objc.Object](objc.GetClass("NSApplication"), objc.Sel("sharedApplication"))
objc.Call[objc.Void](app, objc.Sel("run"))
```

So we wrap these calls in a [Go API](docs/bindings.md) that lets us write code like this:

```go
app := appkit.Application_SharedApplication()
app.Run()
```

These bindings are great, but we need to define them for every API we want to use. Presently,
Apple has around 200 frameworks of nearly 5000 classes with 77k combined methods and properties. Not to 
mention all the constants, functions, structs, unions, and enums we need to work with those objects.

So DarwinKit generates its bindings. This is the hard part. Making sure the generation pipeline accurately produces usable bindings for all possible symbols is quite an arduous, iterative, manual process. Then since we're moving symbols that lived in a single namespace into Go packages, we have to manually decouple dependencies between them enough to avoid circular imports. If you want to help add frameworks, read our documentation on [generation](docs/generation.md).

Objects are passed around as typed pointer values in Objective-C. When we receive an object from a method
call in Go, the `objc` package receives it as a raw pointer, which it first puts into an `unsafe.Pointer`. The
bindings for a class define a struct type that embeds an `objc.Object` struct, which contains a single
field to hold the `unsafe.Pointer`. So unless working with a primitive type, you're working with an `unsafe.Pointer` wrapped in an `objc.Object` wrapped in a struct type that has the methods for the class of the object of the pointer. Be sure to read our documentation on [memory management](docs/memorymanagement.md).

If you have questions, feel free to ask in the [discussion forums](https://github.com/progrium/darwinkit/discussions).

## Thanks

This project was inspired by and originally based on packages written by [Mikkel Krautz](https://github.com/mkrautz). The latest version is based on packages written by [Dong Liu](https://github.com/hsiafan).

## Notice

This project is not affiliated or supported by Apple.

## License

MIT
