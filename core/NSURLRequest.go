package core

import "github.com/progrium/macdriver/objc"

type NSURLRequest struct {
	objc.Object
}

var NSURLRequest_ = objc.Get("NSURLRequest")

func NSURLRequest_Init(url NSURL) NSURLRequest {
	return NSURLRequest{NSURLRequest_.Send("requestWithURL:", url)}
}
