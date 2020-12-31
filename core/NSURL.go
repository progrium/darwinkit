package core

import (
	"github.com/progrium/macdriver/objc"
)

type NSURL struct {
	objc.Object
}

var NSURL_ = objc.Get("NSURL")

func NSURL_Init(url string) NSURL {
	return NSURL{NSURL_.Send("URLWithString:", NSString_FromString(url))}
}
