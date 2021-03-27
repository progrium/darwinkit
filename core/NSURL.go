package core

import (
	"github.com/progrium/macdriver/objc"
)

type NSURL struct {
	objc.Object
}

var nsURL = objc.Get("NSURL")

func NSURL_Init(url string) NSURL {
	return NSURL{nsURL.Send("URLWithString:", NSString_FromString(url))}
}
