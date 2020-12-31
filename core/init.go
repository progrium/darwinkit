package core

/*
#cgo CFLAGS: -x objective-c

void dispatch(void *f);

*/
import "C"

import (
	"github.com/progrium/macdriver/objc"
)

var (
	True  objc.Object
	False objc.Object

	dispatched chan func()
)

func init() {
	True = NSNumber_WithBool(true)
	False = NSNumber_WithBool(false)

	dispatched = make(chan func(), 1)
}

//export dispatchCb
func dispatchCb() {
	(<-dispatched)()
}

func Dispatch(fn func()) {
	dispatched <- fn
	C.dispatch(nil)
}

func String(str string) NSString {
	return NSString_FromString(str)
}

func Point(x float64, y float64) NSPoint {
	return NSPoint{X: x, Y: y}
}

func Size(width float64, height float64) NSSize {
	return NSSize{Width: width, Height: height}
}

func Rect(x, y, w, h float64) NSRect {
	return NSMakeRect(x, y, w, h)
}

func URL(url string) NSURL {
	return NSURL_Init(url)
}
