package core

type NSURLRequest struct {
	gen_NSURLRequest
}

func NSURLRequest_Init(url NSURL) NSURLRequest {
	return NSURLRequest_RequestWithURL(url)
}
