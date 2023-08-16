// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The optional methods implemented by the delegate of a keyed unarchiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiverdelegate?language=objc
type PKeyedUnarchiverDelegate interface {
	// optional
	UnarchiverWillFinish(unarchiver KeyedUnarchiver)
	HasUnarchiverWillFinish() bool

	// optional
	UnarchiverDidFinish(unarchiver KeyedUnarchiver)
	HasUnarchiverDidFinish() bool

	// optional
	UnarchiverWillReplaceObjectWithObject(unarchiver KeyedUnarchiver, object objc.Object, newObject objc.Object)
	HasUnarchiverWillReplaceObjectWithObject() bool
}

// A delegate implementation builder for the [PKeyedUnarchiverDelegate] protocol.
type KeyedUnarchiverDelegate struct {
	_UnarchiverWillFinish                  func(unarchiver KeyedUnarchiver)
	_UnarchiverDidFinish                   func(unarchiver KeyedUnarchiver)
	_UnarchiverWillReplaceObjectWithObject func(unarchiver KeyedUnarchiver, object objc.Object, newObject objc.Object)
}

func (di *KeyedUnarchiverDelegate) HasUnarchiverWillFinish() bool {
	return di._UnarchiverWillFinish != nil
}

// Notifies the delegate that decoding is about to finish. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiverdelegate/1415305-unarchiverwillfinish?language=objc
func (di *KeyedUnarchiverDelegate) SetUnarchiverWillFinish(f func(unarchiver KeyedUnarchiver)) {
	di._UnarchiverWillFinish = f
}

// Notifies the delegate that decoding is about to finish. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiverdelegate/1415305-unarchiverwillfinish?language=objc
func (di *KeyedUnarchiverDelegate) UnarchiverWillFinish(unarchiver KeyedUnarchiver) {
	di._UnarchiverWillFinish(unarchiver)
}
func (di *KeyedUnarchiverDelegate) HasUnarchiverDidFinish() bool {
	return di._UnarchiverDidFinish != nil
}

// Notifies the delegate that decoding has finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiverdelegate/1418067-unarchiverdidfinish?language=objc
func (di *KeyedUnarchiverDelegate) SetUnarchiverDidFinish(f func(unarchiver KeyedUnarchiver)) {
	di._UnarchiverDidFinish = f
}

// Notifies the delegate that decoding has finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiverdelegate/1418067-unarchiverdidfinish?language=objc
func (di *KeyedUnarchiverDelegate) UnarchiverDidFinish(unarchiver KeyedUnarchiver) {
	di._UnarchiverDidFinish(unarchiver)
}
func (di *KeyedUnarchiverDelegate) HasUnarchiverWillReplaceObjectWithObject() bool {
	return di._UnarchiverWillReplaceObjectWithObject != nil
}

// Informs the delegate that one object is being substituted for another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiverdelegate/1413012-unarchiver?language=objc
func (di *KeyedUnarchiverDelegate) SetUnarchiverWillReplaceObjectWithObject(f func(unarchiver KeyedUnarchiver, object objc.Object, newObject objc.Object)) {
	di._UnarchiverWillReplaceObjectWithObject = f
}

// Informs the delegate that one object is being substituted for another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiverdelegate/1413012-unarchiver?language=objc
func (di *KeyedUnarchiverDelegate) UnarchiverWillReplaceObjectWithObject(unarchiver KeyedUnarchiver, object objc.Object, newObject objc.Object) {
	di._UnarchiverWillReplaceObjectWithObject(unarchiver, object, newObject)
}

// A concrete type wrapper for the [PKeyedUnarchiverDelegate] protocol.
type KeyedUnarchiverDelegateWrapper struct {
	objc.Object
}

func (k_ KeyedUnarchiverDelegateWrapper) HasUnarchiverWillFinish() bool {
	return k_.RespondsToSelector(objc.Sel("unarchiverWillFinish:"))
}

// Notifies the delegate that decoding is about to finish. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiverdelegate/1415305-unarchiverwillfinish?language=objc
func (k_ KeyedUnarchiverDelegateWrapper) UnarchiverWillFinish(unarchiver IKeyedUnarchiver) {
	objc.Call[objc.Void](k_, objc.Sel("unarchiverWillFinish:"), objc.Ptr(unarchiver))
}

func (k_ KeyedUnarchiverDelegateWrapper) HasUnarchiverDidFinish() bool {
	return k_.RespondsToSelector(objc.Sel("unarchiverDidFinish:"))
}

// Notifies the delegate that decoding has finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiverdelegate/1418067-unarchiverdidfinish?language=objc
func (k_ KeyedUnarchiverDelegateWrapper) UnarchiverDidFinish(unarchiver IKeyedUnarchiver) {
	objc.Call[objc.Void](k_, objc.Sel("unarchiverDidFinish:"), objc.Ptr(unarchiver))
}

func (k_ KeyedUnarchiverDelegateWrapper) HasUnarchiverWillReplaceObjectWithObject() bool {
	return k_.RespondsToSelector(objc.Sel("unarchiver:willReplaceObject:withObject:"))
}

// Informs the delegate that one object is being substituted for another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiverdelegate/1413012-unarchiver?language=objc
func (k_ KeyedUnarchiverDelegateWrapper) UnarchiverWillReplaceObjectWithObject(unarchiver IKeyedUnarchiver, object objc.IObject, newObject objc.IObject) {
	objc.Call[objc.Void](k_, objc.Sel("unarchiver:willReplaceObject:withObject:"), objc.Ptr(unarchiver), object, newObject)
}
