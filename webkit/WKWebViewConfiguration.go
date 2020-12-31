package webkit

import "github.com/progrium/macdriver/objc"

type WKWebViewConfiguration struct {
	objc.Object
}

func WKWebViewConfiguration_New() WKWebViewConfiguration {
	return WKWebViewConfiguration{objc.Get("WKWebViewConfiguration").Alloc().Init()}
}

func (wv WKWebViewConfiguration) Preferences() WKPreferences {
	return WKPreferences{wv.Get("preferences")}
}
