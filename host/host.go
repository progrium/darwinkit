package host

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/manifold/qtalk/golang/mux"
	"github.com/manifold/qtalk/golang/rpc"
	"github.com/progrium/macdriver"
)

func run(peer *rpc.Peer) {
	data, err := ioutil.ReadFile("/Users/progrium/Source/github.com/manifold/tractor/data/icons/tractor_dark.ico")
	fatal(err)

	peer.Bind("Invoke", macdriver.Invoke)
	go peer.Respond()

	// window := macdriver.Window{
	// 	Title:       "Hello",
	// 	Size:        macdriver.Size{W: 480, H: 240},
	// 	Position:    macdriver.Point{X: 200, Y: 200},
	// 	Closable:    true,
	// 	Minimizable: false,
	// 	Resizable:   false,
	// 	Borderless:  false,
	// 	// Image:       base64.StdEncoding.EncodeToString(data),
	// 	// Background:   &macdriver.Color{R: 0, G: 0, B: 1, A: 0.5},
	// }
	// fatal(macdriver.Sync(peer, &window))

	systray := macdriver.StatusItem{
		Menu: &macdriver.Menu{
			Items: []macdriver.MenuItem{
				{Title: "Bar", Enabled: true, OnClick: macdriver.ExportFunc(func() {
					fmt.Println("Bar clicked")
				})},
				{Title: "Foo", Enabled: true, OnClick: macdriver.ExportFunc(func() {
					fmt.Println("Foo clicked")
				})},
				{Separator: true},
				{Title: "Quit", Enabled: true},
			},
		},
		Icon: base64.StdEncoding.EncodeToString(data),
	}
	fatal(macdriver.Sync(peer, &systray))

	time.Sleep(4 * time.Second)

	systray.Text = "Hello"
	systray.Menu.Items = []macdriver.MenuItem{
		{Title: "Zar", Enabled: true, OnClick: macdriver.ExportFunc(func() {
			fmt.Println("Zar clicked")
		})},
		{Title: "Zoo", Enabled: false},
		{Separator: true},
		{Title: "Shutdown", Enabled: true, OnClick: macdriver.ExportFunc(func() {
			fmt.Println("shutdown")
		})},
	}
	fatal(macdriver.Sync(peer, &systray))

	// macdriver.Release(peer, &window)

	// for t := 0; t < 240; t++ {
	// 	pt := macdriver.Point{X: 200 + float64(t*8), Y: 400}
	// 	//window.Title = fmt.Sprintf("%fx%f", pt.X, pt.Y)
	// 	window.Position = pt
	// 	fatal(window.Sync(peer))
	// 	time.Sleep(2 * time.Millisecond)
	// }

	// var sret string
	// _, err = peer.Call("debug", []interface{}{window.Handle().String(), systray}, &sret)
	// fmt.Println("RET: ", sret)
	// fatal(err)

	select {}

}

func Run() {
	cmd := exec.Command("/usr/local/bin/go", "run", "./cmd/macdriver/main.go")
	cmd.Stderr = os.Stdout
	wc, inErr := cmd.StdinPipe()
	fatal(inErr)
	rc, outErr := cmd.StdoutPipe()
	fatal(outErr)
	go func(rwc io.ReadWriteCloser) {
		session := mux.NewSession(context.Background(), rwc)
		run(rpc.NewPeer(session, rpc.JSONCodec{}))
	}(struct {
		io.WriteCloser
		io.Reader
	}{wc, rc})
	fatal(cmd.Run())
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
