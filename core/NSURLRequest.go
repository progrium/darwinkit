package core

import "github.com/progrium/macdriver/objc"

type NSURLRequest struct {
	objc.Object
}

var nsURLRequest = objc.Get("NSURLRequest")

func NSURLRequest_Init(url NSURL) NSURLRequest {
	return NSURLRequest{nsURLRequest.Send("requestWithURL:", url)}
}
