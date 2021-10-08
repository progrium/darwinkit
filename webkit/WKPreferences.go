package webkit

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type WKPreferences struct {
	gen_WKPreferences
}

func (p WKPreferences) SetValueForKey(value objc.Object, key core.NSStringRef) {
	p.SetValue_forKey_(value, key)
}
