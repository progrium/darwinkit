package dispatch

import (
	"errors"
	"os"
	"runtime"
	"testing"

	"github.com/progrium/macdriver/cocoa"
)

func init() {
	runtime.LockOSThread()
}

func TestMain(m *testing.M) {
	go func() {
		os.Exit(m.Run())
	}()
	app := cocoa.NSApp()
	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyAccessory)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}

func TestAsync(t *testing.T) {
	ok := make(chan bool)
	Async(MainQueue(), func() {
		ok <- true
	})
	<-ok
}

func TestSync(t *testing.T) {
	var ok bool
	Sync(MainQueue(), func() {
		ok = true
	})
	if !ok {
		t.Fatal("ok not set to true")
	}
}

func TestDo(t *testing.T) {
	err := errors.New("test")
	d := Do(MainQueue(), func() error {
		return err
	})
	ret := d.Wait()
	if ret != err {
		t.Fatal("unexpected return value from Wait:", ret)
	}
}
