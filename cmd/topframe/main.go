package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	"github.com/progrium/macdriver/pkg/cocoa"
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
	"github.com/progrium/macdriver/pkg/webkit"
	"github.com/progrium/watcher"
)

// console.log
//

var (
	listener  net.Listener
	serveRoot string
	reloadCh  chan struct{}
)

type AppDelegate struct {
	objc.Object `objc:"AppDelegate : NSObject"`
}

func (delegate *AppDelegate) ApplicationDidFinishLaunching(notification objc.Object) {
	config := webkit.WKWebViewConfiguration_New()
	config.Preferences().SetValueForKey(core.True, core.String("developerExtrasEnabled"))

	wv := webkit.WKWebView_Init(cocoa.NSScreen_Main().Frame(), config)
	wv.SetOpaque(false)
	wv.SetBackgroundColor(cocoa.NSColor_Clear())
	wv.SetValueForKey(core.False, core.String("drawsBackground"))

	req := core.NSURLRequest_Init(core.URL(fmt.Sprintf("http://%s", listener.Addr().String())))
	wv.LoadRequest(req)

	w := cocoa.NSWindow_Init(cocoa.NSScreen_Main().Frame(),
		cocoa.NSBorderlessWindowMask, cocoa.NSBackingStoreBuffered, false)
	w.SetTitlebarAppearsTransparent(true)
	w.SetTitleVisibility(cocoa.NSWindowTitleHidden)
	w.SetMovableByWindowBackground(true)
	w.SetOpaque(false)
	w.SetBackgroundColor(cocoa.NSColor_Clear())
	w.SetIgnoresMouseEvents(true)
	w.SetLevel(cocoa.NSMainMenuWindowLevel + 2)
	w.SetContentView(wv)
	w.MakeKeyAndOrderFront(w)

	go func() {
		for range reloadCh {
			wv.Reload(nil)
		}
	}()
}

func main() {
	runtime.LockOSThread()

	var err error
	listener, err = net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}

	reloadCh = make(chan struct{})

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	serveRoot = filepath.Join(usr.HomeDir, ".topframe")
	os.MkdirAll(serveRoot, 0755)

	c := objc.NewClass(AppDelegate{})
	c.AddMethod("applicationDidFinishLaunching:", (*AppDelegate).ApplicationDidFinishLaunching)
	objc.RegisterClass(c)

	go server()
	go watch()

	delegate := objc.Get("AppDelegate").Alloc().Init()
	app := cocoa.NSApp()
	app.SetDelegate(delegate)
	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyAccessory)
	app.ActivateIgnoringOtherApps(true)
	fmt.Println("running...")
	app.Run()

}

func watch() {
	w := watcher.New()
	if err := w.AddRecursive(serveRoot); err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case event := <-w.Event:
				if event.IsDir() {
					continue
				}
				reloadCh <- struct{}{}
			case err := <-w.Error:
				if err != watcher.ErrWatchedFileDeleted {
					log.Fatal(err)
				}
			case <-w.Closed:
				return
			}
		}
	}()
	log.Fatal(w.Start(500 * time.Millisecond))
}

func server() {
	srv := http.Server{
		Handler: http.FileServer(http.Dir(serveRoot)),
	}
	log.Fatal(srv.Serve(listener))
}
