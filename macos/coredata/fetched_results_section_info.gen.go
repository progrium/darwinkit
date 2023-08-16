// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines the interface for section objects vended by a fetched results controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultssectioninfo?language=objc
type PFetchedResultsSectionInfo interface {
	// optional
	IndexTitle() string
	HasIndexTitle() bool

	// optional
	Name() string
	HasName() bool

	// optional
	NumberOfObjects() uint
	HasNumberOfObjects() bool

	// optional
	Objects() []objc.IObject
	HasObjects() bool
}

// A concrete type wrapper for the [PFetchedResultsSectionInfo] protocol.
type FetchedResultsSectionInfoWrapper struct {
	objc.Object
}

func (f_ FetchedResultsSectionInfoWrapper) HasIndexTitle() bool {
	return f_.RespondsToSelector(objc.Sel("indexTitle"))
}

// The index title of the section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultssectioninfo/1622300-indextitle?language=objc
func (f_ FetchedResultsSectionInfoWrapper) IndexTitle() string {
	rv := objc.Call[string](f_, objc.Sel("indexTitle"))
	return rv
}

func (f_ FetchedResultsSectionInfoWrapper) HasName() bool {
	return f_.RespondsToSelector(objc.Sel("name"))
}

// The name of the section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultssectioninfo/1622302-name?language=objc
func (f_ FetchedResultsSectionInfoWrapper) Name() string {
	rv := objc.Call[string](f_, objc.Sel("name"))
	return rv
}

func (f_ FetchedResultsSectionInfoWrapper) HasNumberOfObjects() bool {
	return f_.RespondsToSelector(objc.Sel("numberOfObjects"))
}

// The number of objects (rows) in the section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultssectioninfo/1622289-numberofobjects?language=objc
func (f_ FetchedResultsSectionInfoWrapper) NumberOfObjects() uint {
	rv := objc.Call[uint](f_, objc.Sel("numberOfObjects"))
	return rv
}

func (f_ FetchedResultsSectionInfoWrapper) HasObjects() bool {
	return f_.RespondsToSelector(objc.Sel("objects"))
}

// The array of objects in the section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultssectioninfo/1622293-objects?language=objc
func (f_ FetchedResultsSectionInfoWrapper) Objects() []objc.Object {
	rv := objc.Call[[]objc.Object](f_, objc.Sel("objects"))
	return rv
}
