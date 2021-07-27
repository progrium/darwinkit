package main

import (
	_ "embed"
	"log"
	"sync"
	"time"
	"unsafe"

	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

/*
#cgo CFLAGS: -x objective-c -DGL_SILENCE_DEPRECATION
#cgo LDFLAGS: -framework OpenGL -framework AppKit
#import "MyOpenGLView.h"
*/
import "C"

//go:generate ibtool --output-format human-readable-text --compile app.nib app.xib

//go:embed app.nib
var nibData []byte

var renderFuncs sync.Map

var app cocoa.NSApplication
var view cocoa.NSView

func main() {
	app = cocoa.NSApp_WithDidLaunch(didLaunch)
	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyRegular)
	app.Run()
}

func didLaunch(objc.Object) {
	loadNib(app, nibData, func(o objc.Object) {
		switch o.Class().String() {
		case "NSWindow":
			view = cocoa.NSWindow{Object: o}.ContentView()
		}
	})

	// Did we find a window?
	if view.Pointer() == 0 {
		log.Fatal("Could not locate window")
	}

	// Is the content view a MyOpenGLView?
	if view.Class().String() != "MyOpenGLView" {
		log.Fatal("Content view is the wrong type")
	}

	// Initialize OpenGL
	err := setupGL()
	if err != nil {
		log.Fatal(err)
	}

	// Request updates at 60Hz
	go func() {
		for range time.Tick(time.Second / 60) {
			core.Dispatch(func() { view.Send("setNeedsDisplay:", true) })
		}
	}()

	app.ActivateIgnoringOtherApps(true)
}

func loadNib(owner objc.Object, data []byte, fn func(objc.Object)) bool {
	nsdata := core.NSData_WithBytes(data, uint64(len(data)))
	bundle := cocoa.NSBundle_Main()
	nib := objc.Get("NSNib").Alloc().Send("initWithNibData:bundle:", nsdata, bundle)

	if fn == nil {
		return nib.Send("instantiateWithOwner:topLevelObjects:", owner, nil).Bool()
	}

	var ptr uintptr
	ok := nib.Send("instantiateWithOwner:topLevelObjects:", owner, &ptr).Bool()
	if !ok {
		return false
	}

	tlo := core.NSArray{Object: objc.ObjectPtr(ptr)}
	for i, n := uint64(0), tlo.Count(); i < n; i++ {
		fn(tlo.ObjectAtIndex(i))
	}

	return true
}

type glRenderFunc func(cocoa.NSView) bool

func setGLRenderFunc(view objc.Object, fn glRenderFunc) {
	if view.Class().String() != "MyOpenGLView" {
		panic("view is not a MyOpenGLView")
	}
	renderFuncs.Store(view.Pointer(), fn)
}

//export callGLRenderFunc
func callGLRenderFunc(id unsafe.Pointer) C.bool {
	fn, ok := renderFuncs.Load(uintptr(id))
	if !ok {
		return false
	}

	o := objc.ObjectPtr(uintptr(id))
	return (C.bool)(fn.(glRenderFunc)(cocoa.NSView{Object: o}))
}

//export deleteGLRenderFunc
func deleteGLRenderFunc(id unsafe.Pointer) {
	renderFuncs.Delete(uintptr(id))
}
