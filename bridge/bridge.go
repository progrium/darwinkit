package bridge

import (
	"context"
	"fmt"
	"io"
	"os"
	"reflect"
	"sync"

	"github.com/manifold/qtalk/golang/mux"
	"github.com/manifold/qtalk/golang/rpc"
	"github.com/mitchellh/mapstructure"
	"github.com/progrium/macdriver/bridge/resource"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

var bridge *Bridge

func init() {
	fmt.Fprintln(os.Stderr, "registering...")
	resource.RegisterType("win", reflect.TypeOf(Window{}))
	resource.RegisterType("ind", reflect.TypeOf(Indicator{}))
	resource.RegisterType("mnu", reflect.TypeOf(Menu{}))
}

func Sync(p *rpc.Peer, v interface{}) error {
	if !resource.HasHandle(v) {
		return fmt.Errorf("not a resource")
	}
	handle := resource.GetHandle(v)
	fmt.Println("get handle:", handle.Handle())
	if handle == nil {
		// TODO: use proper registered prefix
		handle = resource.NewHandle(reflect.ValueOf(v).Type().Elem().Name())
		fmt.Println("new handle:", handle.Handle())
	}
	var h string
	fmt.Println("sync:", handle.Handle())
	_, err := p.Call("Apply", []interface{}{*handle, v}, &h)
	resource.SetHandle(v, h)
	return err
}

func Release(p *rpc.Peer, v interface{}) error {
	if !resource.HasHandle(v) {
		return fmt.Errorf("not a resource")
	}
	handle := resource.GetHandle(v)
	fmt.Println("release:", handle.Handle())
	if handle == nil {
		return fmt.Errorf("unable to release an uninitialized resource")
	}
	_, err := p.Call("Release", *handle, nil)
	if err == nil {
		resource.SetHandle(v, "")
	}
	return err
}

func Run() {
	bridge = NewBridge()
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		session := mux.NewSession(context.Background(), struct {
			io.ReadCloser
			io.Writer
		}{os.Stdin, os.Stdout})
		peer := rpc.NewPeer(session, rpc.JSONCodec{})
		// peer.Bind("debug", debug)
		peer.Bind("Apply", bridge.Apply)
		peer.Bind("Release", bridge.Release)
		//peer.Bind("Invoke", state.Invoke)
		go peer.Respond()
	})
	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}

// func debug(handle string, data interface{}) string {
// 	return fmt.Sprintf("%#v", state.StatusItems[0].Menu)
// }

// func newResource(h resource.Handle) (reflect.Value, error) {
// 	t, found := initTypes[h.Type()]
// 	if !found {
// 		return reflect.Value{}, fmt.Errorf("type not found")
// 	}
// 	return reflect.New(t), nil
// }

func Invoke(ptr string) error {
	// todo args
	ef, ok := exportedFuncs[ptr]
	// reflect call
	if ok {
		ef.fn.(func())()
	}
	return nil
}

type Bridge struct {
	Resources []interface{}

	objects  map[resource.Handle]objc.Object
	released map[resource.Handle]bool

	caller rpc.Caller

	sync.Mutex
}

func NewBridge() *Bridge {
	return &Bridge{
		objects:  make(map[resource.Handle]objc.Object),
		released: make(map[resource.Handle]bool),
	}
}

func (s *Bridge) Release(h resource.Handle) (err error) {
	s.Lock()
	s.released[h] = true
	core.Dispatch(func() {
		//fmt.Fprintf(os.Stderr, "RECONCILING %v\n", handle)
		if err := s.Reconcile(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		s.Unlock()
	})
	// TODO: remove from State slices
	return nil
}

func (s *Bridge) Apply(h string, patch map[string]interface{}, call *rpc.Call) (resource.Handle, error) {
	s.Lock()
	handle := resource.Handle(h)
	fmt.Fprintln(os.Stderr, "handle:", handle, handle.Prefix(), handle.ID())

	Walk(patch, func(v, p reflect.Value, path []string) error {
		if path[len(path)-1] == "$fnptr" {
			p.SetMapIndex(reflect.ValueOf("Caller"), reflect.ValueOf(call.Caller))
		}
		return nil
	})

	v, err := s.Lookup(handle)
	if err != nil {
		return handle, err
	}
	if !v.IsValid() {
		// var err error
		v = reflect.ValueOf(resource.New(handle.Prefix()))
		// if err != nil {
		// 	return handle, err
		// }
		resource.SetHandle(v.Interface(), handle.Handle())
		delete(patch, "Handle")
	}
	if err := mapstructure.Decode(patch, v.Interface()); err != nil {
		return handle, err
	}
	s.Resources = append(s.Resources, v.Interface())

	core.Dispatch(func() {
		//fmt.Fprintf(os.Stderr, "RECONCILING %v\n", handle)
		if err := s.Reconcile(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		s.Unlock()
	})
	return handle, err
}

func (s *Bridge) Lookup(handle resource.Handle) (found reflect.Value, err error) {
	for _, r := range s.Resources {
		if !resource.HasHandle(r) {
			continue
		}
		h := resource.GetHandle(r)
		if h != nil && *h == handle {
			found = reflect.ValueOf(r)
			return found, err
		}
	}
	return found, err
}

func (s *Bridge) Reconcile() error {
	for _, r := range s.Resources {
		if !resource.HasHandle(r) {
			continue
		}
		h := resource.GetHandle(r)
		if h != nil {
			target, exists := s.objects[*h]
			if s.released[*h] {
				// if in released but not in objects,
				// its stale state that should have been cleaned up
				// so we will ignore it here
				if !exists {
					continue
				}
				rd, ok := r.(resource.Discarder)
				if ok && target != nil {
					if err := rd.Discard(target); err != nil {
						//delete(s.objects, *h)
						return err
					}
				}
				delete(s.objects, *h)
				continue
			}
			ra, ok := r.(resource.Applier)
			if ok {
				var err error
				target, err = ra.Apply(target)
				if err != nil {
					return err
				}
				s.objects[*h] = target
			}
		}
	}
	return nil
}
