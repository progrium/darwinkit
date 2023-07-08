package core

import (
	"github.com/progrium/macdriver/objc"
)

type NSDictionary struct {
	gen_NSDictionary
}

func NSDictionary_New() NSDictionary {
	return NSDictionary_Alloc().Init()
}

func NSDictionary_Init(valueKeys ...interface{}) NSDictionary {
	return NSDictionary_FromRef(
		NSDictionary_Alloc().Send("initWithObjectsAndKeys:", valueKeys...))
}

func (d NSDictionary) ObjectForKey(key objc.Object) objc.Object {
	return d.Send("objectForKey:", key)
}

func (d NSDictionary) Count() int {
	return int(d.gen_NSDictionary.Count())
}
