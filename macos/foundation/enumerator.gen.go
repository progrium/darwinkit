// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var EnumeratorClass = _EnumeratorClass{objc.GetClass("NSEnumerator")}

type _EnumeratorClass struct {
	objc.Class
}

type IEnumerator interface {
	objc.IObject
	NextObject() objc.Object
	AllObjects() []objc.Object
}

type Enumerator struct {
	objc.Object
}

func MakeEnumerator(ptr unsafe.Pointer) Enumerator {
	return Enumerator{
		Object: objc.MakeObject(ptr),
	}
}

func (ec _EnumeratorClass) Alloc() Enumerator {
	rv := objc.CallMethod[Enumerator](ec, objc.GetSelector("alloc"))
	return rv
}

func Enumerator_Alloc() Enumerator {
	return EnumeratorClass.Alloc()
}

func (ec _EnumeratorClass) New() Enumerator {
	rv := objc.CallMethod[Enumerator](ec, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewEnumerator() Enumerator {
	return EnumeratorClass.New()
}

func Enumerator_New() Enumerator {
	return EnumeratorClass.New()
}

func (e_ Enumerator) Init() Enumerator {
	rv := objc.CallMethod[Enumerator](e_, objc.GetSelector("init"))
	return rv
}

func Enumerator_Init() Enumerator {
	return EnumeratorClass.Alloc().Init()
}

func (e_ Enumerator) NextObject() objc.Object {
	rv := objc.CallMethod[objc.Object](e_, objc.GetSelector("nextObject"))
	return rv
}

func (e_ Enumerator) AllObjects() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](e_, objc.GetSelector("allObjects"))
	// TODO: convert slice items...
	return rv
}
