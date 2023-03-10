package cocoa

import (
	"log"
	"testing"

	"github.com/progrium/macdriver/objc"
)

type TestView struct {
	objc.Object `objc:"TestView : NSView"`
}

func (v *TestView) NewMethod() int64 {
	return 32
}

func (v *TestView) KeyDown(event objc.Object) {
	log.Println("Keydown!!")
}

func TestSubclass(t *testing.T) {
	c := objc.NewClassFromStruct(TestView{})
	c.AddMethod("newMethod", (*TestView).NewMethod)
	c.AddMethod("keyDown:", (*TestView).KeyDown)
	objc.RegisterClass(c)

	o := objc.Get("TestView").Send("alloc").Send("init")
	i := o.Send("newMethod").Int()
	if i != 32 {
		t.Fatal("unexpected value")
	}

	o.Send("keyDown:", nil)
}
