package coregraphics

// #import <CoreGraphics/CGGeometry.h>
import "C"
import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// struct def should be sync with struct in <CoreGraphics/CGGeometry.h> <CoreGraphics/CGAffineTransform.h>

// todo
type PathElement struct{}
type ScreenUpdateMoveDelta struct{}

type ColorConversionInfoRef unsafe.Pointer
type DisplayStreamUpdateRef unsafe.Pointer

type AffineTransform struct {
	A  Float
	B  Float
	C  Float
	D  Float
	TX Float
	TY Float
}

type Point struct {
	X Float
	Y Float
}

type Rect struct {
	Origin Point
	Size   Size
}

type Size struct {
	Width  Float
	Height Float
}

var RectNull = objc.ForceCast[C.CGRect, Rect](C.CGRectNull)
var RectInfinite = objc.ForceCast[C.CGRect, Rect](C.CGRectInfinite)
