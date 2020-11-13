package webkit

import "github.com/progrium/macdriver/pkg/objc"

type WKPreferences struct {
	objc.Object
}

func (p WKPreferences) SetValueForKey(value, key objc.Object) {
	p.Send("setValue:forKey:", value, key)
}
