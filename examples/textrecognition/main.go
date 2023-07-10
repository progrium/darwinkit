package main

import (
	"fmt"

	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/vision"
)

func main() {
	u := core.NSURL_URLWithString(core.String("file:///Users/tmc/Desktop/screenshots/ss.png"))

	opts := core.NSDictionary_New()
	rh := vision.VNImageRequestHandler_Alloc().InitWithURLOptions(u, opts)

	req := vision.VNRecognizeTextRequest_Alloc().Init()
	reqs := core.NSArray_Alloc().Init()
	fmt.Println(rh.PerformRequestsError(reqs, nsErr))
	fmt.Println(nsErr)
}
