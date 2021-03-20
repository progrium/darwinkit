package core

import (
	"github.com/progrium/macdriver/objc"
)

import "C"

// NSArray is a static ordered collection of objects.
// https://developer.apple.com/documentation/foundation/nsarray?language=objc
type NSArray struct {
	objc.Object
}

// NSArray_WithObjects creates and returns an array containing the objects in the argument list.
// https://developer.apple.com/documentation/foundation/nsarray/1460068-initwithobjects?language=objc
func NSArray_WithObjects(objs ...objc.Object) NSArray {
	objsInterface := make([]interface{}, len(objs))
	for i, obj := range objs {
		objsInterface[i] = obj
	}
	return NSArray{objc.Get("NSArray").Send("arrayWithObjects:", objsInterface...)}
}

// Count returns the number of objects in the array.
// https://developer.apple.com/documentation/foundation/nsarray/1409982-count?language=objc
func (a NSArray) Count() uint64 {
	return a.Get("count").Uint()
}

// ObjectAtIndex returns the object located at the specified index.
// https://developer.apple.com/documentation/foundation/nsarray/1417555-objectatindex?language=objc
func (a NSArray) ObjectAtIndex(i uint64) objc.Object {
	return a.Send("objectAtIndex:", i)
}

