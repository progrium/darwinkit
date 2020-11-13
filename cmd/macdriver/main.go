package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/manifold/qtalk/golang/mux"
	"github.com/manifold/qtalk/golang/rpc"
	"github.com/progrium/macdriver"
	"github.com/progrium/macdriver/pkg/cocoa"
	"github.com/progrium/macdriver/pkg/objc"
)

// TODO: lock
var (
	state *macdriver.State
)

// func init() {
// 	runtime.LockOSThread()
// }

func main() {
	c := objc.NewClass(macdriver.AppDelegate{})
	c.AddMethod("applicationDidFinishLaunching:", (*macdriver.AppDelegate).ApplicationDidFinishLaunching)
	objc.RegisterClass(c)

	state = macdriver.NewState()
	go server()

	app := cocoa.NSApp()
	app.SetDelegate(objc.Get("AppDelegate").Alloc().Init())
	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyAccessory)
	go func() {
		<-time.After(3 * time.Second)
		app.Terminate()
	}()
	app.Run()
}

func server() {
	session := mux.NewSession(context.Background(), struct {
		io.ReadCloser
		io.Writer
	}{os.Stdin, os.Stdout})
	peer := rpc.NewPeer(session, rpc.JSONCodec{})
	peer.Bind("create", create)
	peer.Bind("update", update)
	peer.Bind("delete", delete)
	peer.Bind("debug", debug)
	peer.Respond()
}

func create(type_ string, data interface{}) string {
	defer reconcile(state)
	handle, err := state.Init(type_, data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}
	return handle
}

func update(handle string, data interface{}) bool {
	defer reconcile(state)
	err := state.Update(macdriver.Handle(handle), data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}
	return true
}

func delete(handle string) bool {
	defer reconcile(state)
	return true
}

func debug(handle string) string {
	r, err := state.Lookup(macdriver.Handle(handle))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return fmt.Sprintf("%#v %#v", state, r)
}

func reconcile(state *macdriver.State) {
	if err := state.Reconcile(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
