// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"runtime"

	. "github.com/progrium/macdriver/pkg/ns/AppKit"
	"github.com/progrium/macdriver/pkg/objc"
)

func init() {
	defer runtime.LockOSThread()

	c := objc.NewClass(AppDelegate{})
	c.AddMethod("applicationDidFinishLaunching:", (*AppDelegate).ApplicationDidFinishLaunching)
	objc.RegisterClass(c)
}

type AppDelegate struct {
	objc.Object `objc:"GOAppDelegate : NSObject"`
	Window      objc.Object `objc:"IBOutlet"`
}

func (delegate *AppDelegate) ApplicationDidFinishLaunching(notification objc.Object) {
	log.Printf("applicationDidFinishLaunching! %v", notification)

	log.Printf("Object: %v", delegate.Object)
	log.Printf("Window: %v", delegate.Window)

	mainMenu := NSSharedApplication().MainMenu()
	log.Printf("%v", mainMenu)
}

func main() {
	NSApplicationMain()
}
