package core

import "github.com/progrium/macdriver/objc"

type NSNumber struct {
	objc.Object
}

func NSNumber_WithBool(b bool) NSNumber {
	return NSNumber{objc.Get("NSNumber").Send("numberWithBool:", b)}
}
