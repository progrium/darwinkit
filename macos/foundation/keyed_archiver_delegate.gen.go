// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The optional methods implemented by the delegate of a keyed archiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiverdelegate?language=objc
type PKeyedArchiverDelegate interface {
	// optional
	ArchiverWillReplaceObjectWithObject(archiver KeyedArchiver, object objc.Object, newObject objc.Object)
	HasArchiverWillReplaceObjectWithObject() bool

	// optional
	ArchiverDidFinish(archiver KeyedArchiver)
	HasArchiverDidFinish() bool

	// optional
	ArchiverWillFinish(archiver KeyedArchiver)
	HasArchiverWillFinish() bool
}

// A delegate implementation builder for the [PKeyedArchiverDelegate] protocol.
type KeyedArchiverDelegate struct {
	_ArchiverWillReplaceObjectWithObject func(archiver KeyedArchiver, object objc.Object, newObject objc.Object)
	_ArchiverDidFinish                   func(archiver KeyedArchiver)
	_ArchiverWillFinish                  func(archiver KeyedArchiver)
}

func (di *KeyedArchiverDelegate) HasArchiverWillReplaceObjectWithObject() bool {
	return di._ArchiverWillReplaceObjectWithObject != nil
}

// Informs the delegate that one given object is being substituted for another given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiverdelegate/1409389-archiver?language=objc
func (di *KeyedArchiverDelegate) SetArchiverWillReplaceObjectWithObject(f func(archiver KeyedArchiver, object objc.Object, newObject objc.Object)) {
	di._ArchiverWillReplaceObjectWithObject = f
}

// Informs the delegate that one given object is being substituted for another given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiverdelegate/1409389-archiver?language=objc
func (di *KeyedArchiverDelegate) ArchiverWillReplaceObjectWithObject(archiver KeyedArchiver, object objc.Object, newObject objc.Object) {
	di._ArchiverWillReplaceObjectWithObject(archiver, object, newObject)
}
func (di *KeyedArchiverDelegate) HasArchiverDidFinish() bool {
	return di._ArchiverDidFinish != nil
}

// Notifies the delegate that encoding has finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiverdelegate/1412480-archiverdidfinish?language=objc
func (di *KeyedArchiverDelegate) SetArchiverDidFinish(f func(archiver KeyedArchiver)) {
	di._ArchiverDidFinish = f
}

// Notifies the delegate that encoding has finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiverdelegate/1412480-archiverdidfinish?language=objc
func (di *KeyedArchiverDelegate) ArchiverDidFinish(archiver KeyedArchiver) {
	di._ArchiverDidFinish(archiver)
}
func (di *KeyedArchiverDelegate) HasArchiverWillFinish() bool {
	return di._ArchiverWillFinish != nil
}

// Notifies the delegate that encoding is about to finish. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiverdelegate/1411119-archiverwillfinish?language=objc
func (di *KeyedArchiverDelegate) SetArchiverWillFinish(f func(archiver KeyedArchiver)) {
	di._ArchiverWillFinish = f
}

// Notifies the delegate that encoding is about to finish. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiverdelegate/1411119-archiverwillfinish?language=objc
func (di *KeyedArchiverDelegate) ArchiverWillFinish(archiver KeyedArchiver) {
	di._ArchiverWillFinish(archiver)
}

// A concrete type wrapper for the [PKeyedArchiverDelegate] protocol.
type KeyedArchiverDelegateWrapper struct {
	objc.Object
}

func (k_ KeyedArchiverDelegateWrapper) HasArchiverWillReplaceObjectWithObject() bool {
	return k_.RespondsToSelector(objc.Sel("archiver:willReplaceObject:withObject:"))
}

// Informs the delegate that one given object is being substituted for another given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiverdelegate/1409389-archiver?language=objc
func (k_ KeyedArchiverDelegateWrapper) ArchiverWillReplaceObjectWithObject(archiver IKeyedArchiver, object objc.IObject, newObject objc.IObject) {
	objc.Call[objc.Void](k_, objc.Sel("archiver:willReplaceObject:withObject:"), objc.Ptr(archiver), object, newObject)
}

func (k_ KeyedArchiverDelegateWrapper) HasArchiverDidFinish() bool {
	return k_.RespondsToSelector(objc.Sel("archiverDidFinish:"))
}

// Notifies the delegate that encoding has finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiverdelegate/1412480-archiverdidfinish?language=objc
func (k_ KeyedArchiverDelegateWrapper) ArchiverDidFinish(archiver IKeyedArchiver) {
	objc.Call[objc.Void](k_, objc.Sel("archiverDidFinish:"), objc.Ptr(archiver))
}

func (k_ KeyedArchiverDelegateWrapper) HasArchiverWillFinish() bool {
	return k_.RespondsToSelector(objc.Sel("archiverWillFinish:"))
}

// Notifies the delegate that encoding is about to finish. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiverdelegate/1411119-archiverwillfinish?language=objc
func (k_ KeyedArchiverDelegateWrapper) ArchiverWillFinish(archiver IKeyedArchiver) {
	objc.Call[objc.Void](k_, objc.Sel("archiverWillFinish:"), objc.Ptr(archiver))
}
