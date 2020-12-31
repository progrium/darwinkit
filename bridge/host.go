package bridge

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
)

func runHost(peer *rpc.Peer) error {
	data, err := ioutil.ReadFile("/Users/progrium/Source/github.com/manifold/tractor/data/icons/tractor_dark.ico")
	if err != nil {
		return err
	}

	peer.Bind("Invoke", Invoke)
	go peer.Respond()

	window := Window{
		Title:       "Hello",
		Size:        Size{W: 480, H: 240},
		Position:    Point{X: 200, Y: 200},
		Closable:    true,
		Minimizable: false,
		Resizable:   false,
		Borderless:  false,
		// Image:       base64.StdEncoding.EncodeToString(data),
		// Background:   &Color{R: 0, G: 0, B: 1, A: 0.5},
	}
	if err := Sync(peer, &window); err != nil {
		return err
	}

	window2 := Window{
		Title:       "Hello2",
		Size:        Size{W: 480, H: 240},
		Position:    Point{X: 400, Y: 200},
		Closable:    true,
		Minimizable: false,
		Resizable:   false,
		Borderless:  false,
		// Image:       base64.StdEncoding.EncodeToString(data),
		// Background:   &Color{R: 0, G: 0, B: 1, A: 0.5},
	}
	if err := Sync(peer, &window2); err != nil {
		return err
	}

	if err := Release(peer, &window); err != nil {
		return err
	}

	window3 := Window{
		Title:       "Hello3",
		Size:        Size{W: 480, H: 240},
		Position:    Point{X: 500, Y: 200},
		Closable:    true,
		Minimizable: false,
		Resizable:   false,
		Borderless:  false,
		// Image:       base64.StdEncoding.EncodeToString(data),
		// Background:   &Color{R: 0, G: 0, B: 1, A: 0.5},
	}
	if err := Sync(peer, &window3); err != nil {
		return err
	}

	systray := StatusItem{
		Menu: &Menu{
			Items: []MenuItem{
				{Title: "Bar", Enabled: true, OnClick: ExportFunc(func() {
					fmt.Println("Bar clicked")
				})},
				{Title: "Foo", Enabled: true, OnClick: ExportFunc(func() {
					fmt.Println("Foo clicked")
				})},
				{Separator: true},
				{Title: "Quit", Enabled: true},
			},
		},
		Icon: base64.StdEncoding.EncodeToString(data),
	}
	if err := Sync(peer, &systray); err != nil {
		return err
	}

	time.Sleep(4 * time.Second)

	systray.Text = "Hello"
	systray.Menu.Items = []MenuItem{
		{Title: "Zar", Enabled: true, OnClick: ExportFunc(func() {
			fmt.Println("Zar clicked")
		})},
		{Title: "Zoo", Enabled: false},
		{Separator: true},
		{Title: "Shutdown", Enabled: true, OnClick: ExportFunc(func() {
			fmt.Println("shutdown")
		})},
	}
	if err := Sync(peer, &systray); err != nil {
		return err
	}

	select {}
}

type Host struct {
	Cmd  *exec.Cmd
	Pipe io.ReadWriteCloser
}

func (h *Host) Run() error {
	go func() {
		session := mux.NewSession(context.Background(), h.Pipe)
		log.Fatal(runHost(rpc.NewPeer(session, rpc.JSONCodec{})))
	}()
	return h.Cmd.Run()
}

func NewHost(stderr io.Writer) (*Host, error) {
	bridgebin := os.Getenv("BRIDGEBIN")
	if bridgebin == "" {
		bridgebin = "macbridge"
	}
	cmd := exec.Command(bridgebin)
	cmd.Stderr = stderr
	wc, inErr := cmd.StdinPipe()
	if inErr != nil {
		return nil, inErr
	}
	rc, outErr := cmd.StdoutPipe()
	if outErr != nil {
		return nil, outErr
	}
	pipe := struct {
		io.WriteCloser
		io.Reader
	}{wc, rc}
	return &Host{Cmd: cmd, Pipe: pipe}, nil
}
