package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/manifold/qtalk/golang/mux"
	"github.com/manifold/qtalk/golang/rpc"
	"github.com/progrium/macdriver"
)

func run(peer *rpc.Peer) {
	// menu := macdriver.Menu{
	// 	Title: "MyMenu",
	// }
	window := macdriver.Window{
		Title: "MyWindow",
	}

	// var ret bool
	// _, err := peer.Call("create", []interface{}{"Menu", menu}, &ret)
	// fatal(err)
	// mustTrue(ret)

	var handle string
	_, err := peer.Call("create", []interface{}{"Window", window}, &handle)
	fatal(err)
	// mustTrue(ret)

	window.Title = "UPdatedWindow"
	window.Position = macdriver.Point{X: 123, Y: 456}

	var ok bool
	_, err = peer.Call("update", []interface{}{handle, window}, &ok)
	fatal(err)

	var sret string
	_, err = peer.Call("debug", handle, &sret)
	fatal(err)
	fmt.Println(ok, sret)
}

func main() {
	cmd := exec.Command("/usr/local/bin/go", "run", "./cmd/macdriver/main.go")
	cmd.Stderr = os.Stdout
	wc, inErr := cmd.StdinPipe()
	fatal(inErr)
	rc, outErr := cmd.StdoutPipe()
	fatal(outErr)
	go func(rwc io.ReadWriteCloser) {
		session := mux.NewSession(context.Background(), rwc)
		run(rpc.NewPeer(session, rpc.JSONCodec{}))
		wc.Close()
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

func mustTrue(v bool) {
	if v == false {
		log.Fatal("remote failure")
	}
}
