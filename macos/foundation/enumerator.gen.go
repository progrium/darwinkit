// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Enumerator] class.
var EnumeratorClass = _EnumeratorClass{objc.GetClass("NSEnumerator")}

type _EnumeratorClass struct {
	objc.Class
}

// An interface definition for the [Enumerator] class.
type IEnumerator interface {
	objc.IObject
	NextObject() objc.Object
	AllObjects() []objc.Object
}

// An abstract class whose subclasses enumerate collections of objects, such as arrays and dictionaries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenumerator?language=objc
type Enumerator struct {
	objc.Object
}

func EnumeratorFrom(ptr unsafe.Pointer) Enumerator {
	return Enumerator{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _EnumeratorClass) Alloc() Enumerator {
	rv := objc.Call[Enumerator](ec, objc.Sel("alloc"))
	return rv
}

func Enumerator_Alloc() Enumerator {
	return EnumeratorClass.Alloc()
}

func (ec _EnumeratorClass) New() Enumerator {
	rv := objc.Call[Enumerator](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewEnumerator() Enumerator {
	return EnumeratorClass.New()
}

func (e_ Enumerator) Init() Enumerator {
	rv := objc.Call[Enumerator](e_, objc.Sel("init"))
	return rv
}

// Returns the next object from the collection being enumerated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenumerator/1409616-nextobject?language=objc
func (e_ Enumerator) NextObject() objc.Object {
	rv := objc.Call[objc.Object](e_, objc.Sel("nextObject"))
	return rv
}

// The array of unenumerated objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenumerator/1417755-allobjects?language=objc
func (e_ Enumerator) AllObjects() []objc.Object {
	rv := objc.Call[[]objc.Object](e_, objc.Sel("allObjects"))
	return rv
}
