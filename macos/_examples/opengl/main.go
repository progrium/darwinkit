package main

import (
	_ "embed"
	"log"
	"sync"
	"time"
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
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

var app appkit.Application
var view appkit.View

func main() {
	app = appkit.Application_SharedApplication()
	ad := &appkit.ApplicationDelegate{}
	ad.SetApplicationDidFinishLaunching(didLaunch)
	app.SetDelegate(ad)
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.Run()
}

func didLaunch(foundation.Notification) {
	loadNib(app, nibData, func(o objc.Object) {
		switch o.GetClass().GetName() {
		case "NSKVONotifying_NSWindow":
			view = appkit.MakeWindow(o.Ptr()).ContentView()
		default:
			//fmt.Println(o.GetClass().GetName())
		}
	})

	// Did we find a window?
	if uintptr(view.Ptr()) == 0 {
		log.Fatal("Could not locate window")
	}

	// Is the content view a MyOpenGLView?
	if view.GetClass().GetName() != "MyOpenGLView" {
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
			dispatch.GetMainQueue().DispatchAsync(func() { objc.CallMethod[objc.Void](view, objc.GetSelector("setNeedsDisplay")) })
		}
	}()

	app.ActivateIgnoringOtherApps(true)
}

func loadNib(owner appkit.Application, data []byte, fn func(objc.Object)) bool {
	bundle := foundation.Bundle_MainBundle()
	nib := appkit.Nib_InitWithNibDataBundle(data, bundle)

	if fn == nil {
		return nib.InstantiateWithOwnerTopLevelObjects(owner, nil)
	}

	tlo := foundation.NewMutableArray()
	ok := nib.InstantiateWithOwnerTopLevelObjects(owner, &tlo.Array)
	if !ok {
		return false
	}

	for i, n := uint(0), tlo.Count(); i < n; i++ {
		fn(tlo.ObjectAtIndex(i))
	}

	return true
}

type glRenderFunc func(appkit.View) bool

func setGLRenderFunc(view appkit.View, fn glRenderFunc) {
	if view.GetClass().GetName() != "MyOpenGLView" {
		panic("view is not a MyOpenGLView")
	}
	renderFuncs.Store(view.Ptr(), fn)
}

//export callGLRenderFunc
func callGLRenderFunc(id unsafe.Pointer) C.bool {
	fn, ok := renderFuncs.Load(uintptr(id))
	if !ok {
		return false
	}

	return (C.bool)(fn.(glRenderFunc)(appkit.MakeView(id)))
}

//export deleteGLRenderFunc
func deleteGLRenderFunc(id unsafe.Pointer) {
	renderFuncs.Delete(uintptr(id))
}
