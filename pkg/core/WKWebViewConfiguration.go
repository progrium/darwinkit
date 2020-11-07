package core

import "github.com/progrium/macdriver/pkg/objc"

type WKWebViewConfiguration struct {
	objc.Object
}

func WKWebViewConfiguration_New() WKWebViewConfiguration {
	return WKWebViewConfiguration{objc.Get("WKWebViewConfiguration").Alloc().Init()}
}
