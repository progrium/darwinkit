package core

import (
	"github.com/progrium/macdriver/objc"
)

import "C"

// NSArray is a static ordered collection of objects.
// https://developer.apple.com/documentation/foundation/nsarray?language=objc
type NSArray struct {
	gen_NSArray
}

// NSArray_WithObjects creates and returns an array containing the objects in the argument list.
// https://developer.apple.com/documentation/foundation/nsarray/1460068-initwithobjects?language=objc
func NSArray_WithObjects(objs ...objc.Object) NSArray {
	objsInterface := make([]interface{}, len(objs))
	for i, obj := range objs {
		objsInterface[i] = obj
	}
	return NSArray_FromRef(objc.Get("NSArray").Send("arrayWithObjects:", objsInterface...))
}

// Count returns the number of objects in the array.
// https://developer.apple.com/documentation/foundation/nsarray/1409982-count?language=objc
func (a NSArray) Count() uint64 {
	return uint64(a.gen_NSArray.Count())
}

// ObjectAtIndex returns the object located at the specified index.
// https://developer.apple.com/documentation/foundation/nsarray/1417555-objectatindex?language=objc
func (a NSArray) ObjectAtIndex(i uint64) objc.Object {
	return a.Send("objectAtIndex:", i)
}

// Strings returns the NSArray as slice of string by calling the String Method on each objc.Object.
func (a NSArray) Strings() []string {
	count := int(a.Count())
	ss := make([]string, count)
	for i := 0; i < count; i++ {
		o := a.ObjectAtIndex(uint64(i))
		ss[i] = o.String()
	}
	return ss
}
