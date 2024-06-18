package action

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #import <stdint.h>
// void* C_NewAction(uintptr_t hp);
import "C"
import (
	"runtime/cgo"
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// Handler is a callback function for an ActionTarget.
type Handler func(sender objc.Object)

// CanSet is an interface for objc instance which can set a target and action
type CanSet interface {
	objc.IObject
	SetTarget(target objc.IObject)
	SetAction(sel objc.Selector)
}

// Target hold an objc Target target instance
type Target struct {
	objc.Object
}

// Wrap create a new objc Target instance, from Handler
func Wrap(handler Handler) (target Target, selector objc.Selector) {
	if handler == nil {
		panic("handler is nil")
	}
	h := cgo.NewHandle(handler)
	return Target{
		Object: objc.ObjectFrom(C.C_NewAction(C.uintptr_t(h))),
	}, objc.RegisterSelectorName("onAction:")
}

// Set set action for an ojbc instance, if it has target and setAction method.
func Set(instance CanSet, handler Handler) {
	target, selector := Wrap(handler)
	defer target.Release()
	objc.SetAssociatedObject(instance, objc.AssociationKey("target"), target, objc.ASSOCIATION_RETAIN)
	instance.SetTarget(target)
	instance.SetAction(selector)
}

//export callAction
func callAction(hp C.uintptr_t, senderPtr unsafe.Pointer) {
	h := cgo.Handle(hp)
	handler := h.Value().(Handler)
	handler(objc.ObjectFrom(senderPtr))
}

//export deleteActionHandle
func deleteActionHandle(hp C.uintptr_t) {
	h := cgo.Handle(hp)
	h.Delete()
}
