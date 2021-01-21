package bridge

import (
	"context"
	"io"
	"os"
	"os/exec"

	"github.com/manifold/qtalk/golang/mux"
	"github.com/manifold/qtalk/golang/rpc"
)

type Host struct {
	Cmd  *exec.Cmd
	Pipe io.ReadWriteCloser
	Peer *rpc.Peer
}

func (h *Host) Run() error {
	return h.Cmd.Run()
}

func NewHost(stderr io.Writer) (*Host, error) {
	bridgecmd := os.Getenv("BRIDGECMD")
	if bridgecmd == "" {
		bridgecmd = "macbridge"
	}
	cmd := exec.Command(bridgecmd)
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
	session := mux.NewSession(context.Background(), pipe)
	return &Host{Cmd: cmd, Pipe: pipe, Peer: rpc.NewPeer(session, rpc.JSONCodec{})}, nil
}
