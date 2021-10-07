package core

type NSURL struct{ gen_NSURL }

func NSURL_Init(url string) NSURL {
	return NSURL_URLWithString_(String(url))
}
