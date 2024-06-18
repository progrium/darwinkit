# DarwinKit

Native Apple APIs for Golang!

[![GoDoc](https://godoc.org/github.com/progrium/darwinkit?status.svg)](https://pkg.go.dev/github.com/progrium/darwinkit@main)
[![Go Report Card](https://goreportcard.com/badge/github.com/progrium/darwinkit)](https://goreportcard.com/report/github.com/progrium/darwinkit)
<a href="https://twitter.com/progrium" title="@progrium on Twitter"><img src="https://img.shields.io/badge/twitter-@progrium-55acee.svg" alt="@progrium on Twitter"></a>
<a href="https://github.com/progrium/darwinkit/discussions" title="Project Forum"><img src="https://img.shields.io/badge/community-forum-ff69b4.svg" alt="Project Forum"></a>
<a href="https://github.com/sponsors/progrium" title="Sponsor Project"><img src="https://img.shields.io/static/v1?label=sponsor&message=%E2%9D%A4&logo=GitHub" alt="Sponsor Project" /></a>

> [!IMPORTANT]
> June 13, 2024: **MacDriver is now DarwinKit and we're about to release 0.5.0!** The [legacy branch](https://github.com/progrium/darwinkit/tree/legacy) and [previous releases](https://github.com/progrium/darwinkit/releases) are still available for existing code to work against. Use `main` until 0.5.0 is released.

------


<img alt="SnowScape Demo" src="https://github-production-user-asset-6210df.s3.amazonaws.com/1813419/297168913-71fea878-878b-4fa8-ad19-c19458b0594e.gif?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAVCODYLSA53PQK4ZA%2F20240618%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240618T153557Z&X-Amz-Expires=300&X-Amz-Signature=b1dfe34dd04f608524b6b75370084a6d1bb8a87600f44b51d2617cfdfa08866b&X-Amz-SignedHeaders=host&actor_id=647&key_id=0&repo_id=681734532" style="height:125px;" /><img alt="ClipTrail Demo" src="https://private-user-images.githubusercontent.com/1813419/297170797-29cc3907-9fc3-44c6-bd15-91ae0a8a74ba.gif?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTg3MjUzMzgsIm5iZiI6MTcxODcyNTAzOCwicGF0aCI6Ii8xODEzNDE5LzI5NzE3MDc5Ny0yOWNjMzkwNy05ZmMzLTQ0YzYtYmQxNS05MWFlMGE4YTc0YmEuZ2lmP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9QUtJQVZDT0RZTFNBNTNQUUs0WkElMkYyMDI0MDYxOCUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNDA2MThUMTUzNzE4WiZYLUFtei1FeHBpcmVzPTMwMCZYLUFtei1TaWduYXR1cmU9ODgwOWIyZTQ2YjA2NmZlYTY2MWE1MWM2ODc3ZjZiNDZjMTY5YjcxOTExNDVlZmQ5NGI4OGRhZThmZTY3NzJmYiZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmYWN0b3JfaWQ9MCZrZXlfaWQ9MCZyZXBvX2lkPTAifQ.pxMRXGio1SIRDpL9sz_G8g-e2Nm08shb_b6oqrXFS3k" style="height:125px;" /><img alt="ScanDrop Demo" src="https://github-production-user-asset-6210df.s3.amazonaws.com/1813419/297171190-339168f9-b768-461d-a90e-559689bd3a82.gif?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAVCODYLSA53PQK4ZA%2F20240618%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240618T154008Z&X-Amz-Expires=300&X-Amz-Signature=36f5a91041225dd6c9ce54e0b89552f45fd8777f174e8c7db3c01530f6c0e8ea&X-Amz-SignedHeaders=host&actor_id=647&key_id=0&repo_id=681734532" style="height:125px;" /><img alt="MenuSpacer Demo" src="https://github-production-user-asset-6210df.s3.amazonaws.com/1813419/297170958-04da4dd8-c329-4e80-b8c1-72683712f0fc.gif?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAVCODYLSA53PQK4ZA%2F20240618%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240618T154137Z&X-Amz-Expires=300&X-Amz-Signature=3f5795df17e2ab8ba13d10905001ea960d1a885a91619d027bdac4863bf2e2c2&X-Amz-SignedHeaders=host&actor_id=647&key_id=0&repo_id=681734532" style="height:125px;" />





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

### [largetype](./macos/_examples/largetype/main.go)
A Contacts/Quicksilver-style Large Type utility in under 80 lines:

![largetype screenshot](./macos/_examples/largetype/largetype.jpeg?raw=true)

### [pomodoro](./macos/_examples/pomodoro/main.go)
A menu bar pomodoro timer in under 80 lines:

![pomodoro gif](./macos/_examples/pomodoro/pomodoro.gif?raw=true)

### [webshot](./macos/_examples/webshot/main.go)
A webview PNG capture example in under 100 lines:

![webshot screenshot](./macos/_examples/webshot/webshot.png?raw=true)

### [See all examples](./macos/_examples)

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
