package cocoa

import foundation "github.com/progrium/macdriver/pkg/cocoa/Foundation"

func String(str string) foundation.NSString {
	return foundation.NSString_FromString(str)
}

func Point(x float32, y float32) foundation.NSPoint {
	return foundation.NSPoint{X: x, Y: y}
}

func Size(width float32, height float32) foundation.NSSize {
	return foundation.NSSize{Width: width, Height: height}
}

func Rect(x, y, w, h float32) foundation.NSRect {
	return foundation.NSMakeRect(x, y, w, h)
}
