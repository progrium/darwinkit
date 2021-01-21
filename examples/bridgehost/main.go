package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/progrium/macdriver/bridge"
)

func main() {
	h, err := bridge.NewHost(os.Stderr)
	if err != nil {
		log.Fatal(err)
	}
	go h.Run()

	data, err := ioutil.ReadFile("/Users/progrium/Source/github.com/manifold/tractor/data/icons/tractor_dark.ico")
	if err != nil {
		log.Fatal(err)
	}

	h.Peer.Bind("Invoke", bridge.Invoke)
	go h.Peer.Respond()

	window := bridge.Window{
		Title:       "Hello 1",
		Size:        bridge.Size{W: 480, H: 240},
		Position:    bridge.Point{X: 200, Y: 200},
		Closable:    true,
		Minimizable: false,
		Resizable:   false,
		Borderless:  false,
		// Image:       base64.StdEncoding.EncodeToString(data),
		// Background:   &Color{R: 0, G: 0, B: 1, A: 0.5},
	}
	if err := bridge.Sync(h.Peer, &window); err != nil {
		log.Fatal(err)
	}

	window2 := bridge.Window{
		Title:       "Hello 2",
		Size:        bridge.Size{W: 480, H: 240},
		Position:    bridge.Point{X: 400, Y: 200},
		Closable:    true,
		Minimizable: false,
		Resizable:   false,
		Borderless:  false,
		// Image:       base64.StdEncoding.EncodeToString(data),
		// Background:   &Color{R: 0, G: 0, B: 1, A: 0.5},
	}
	if err := bridge.Sync(h.Peer, &window2); err != nil {
		log.Fatal(err)
	}

	if err := bridge.Release(h.Peer, &window); err != nil {
		log.Fatal(err)
	}

	window3 := bridge.Window{
		Title:       "Hello 3",
		Size:        bridge.Size{W: 480, H: 240},
		Position:    bridge.Point{X: 500, Y: 200},
		Closable:    true,
		Minimizable: false,
		Resizable:   false,
		Borderless:  false,
		// Image:       base64.StdEncoding.EncodeToString(data),
		// Background:   &Color{R: 0, G: 0, B: 1, A: 0.5},
	}
	if err := bridge.Sync(h.Peer, &window3); err != nil {
		log.Fatal(err)
	}

	systray := bridge.Indicator{
		Menu: &bridge.Menu{
			Items: []bridge.MenuItem{
				{Title: "Bar", Enabled: true, OnClick: bridge.ExportFunc(func() {
					fmt.Println("Bar clicked")
				})},
				{Title: "Foo", Enabled: true, OnClick: bridge.ExportFunc(func() {
					fmt.Println("Foo clicked")
				})},
				{Separator: true},
				{Title: "Quit", Enabled: true},
			},
		},
		Icon: base64.StdEncoding.EncodeToString(data),
	}
	if err := bridge.Sync(h.Peer, &systray); err != nil {
		log.Fatal(err)
	}

	time.Sleep(4 * time.Second)

	systray.Text = "Hello"
	systray.Menu.Items = []bridge.MenuItem{
		{Title: "Zar", Enabled: true, OnClick: bridge.ExportFunc(func() {
			fmt.Println("Zar clicked")
		})},
		{Title: "Zoo", Enabled: false},
		{Separator: true},
		{Title: "Shutdown", Enabled: true, OnClick: bridge.ExportFunc(func() {
			fmt.Println("shutdown")
		})},
	}
	if err := bridge.Sync(h.Peer, &systray); err != nil {
		log.Fatal(err)
	}

	select {}

}
