package webkit

type WKWebViewConfiguration struct {
	gen_WKWebViewConfiguration
}

func WKWebViewConfiguration_New() WKWebViewConfiguration {
	return WKWebViewConfiguration_alloc().Init_asWKWebViewConfiguration()
}
