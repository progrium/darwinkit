// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PositionalSpecifier] class.
var PositionalSpecifierClass = _PositionalSpecifierClass{objc.GetClass("NSPositionalSpecifier")}

type _PositionalSpecifierClass struct {
	objc.Class
}

// An interface definition for the [PositionalSpecifier] class.
type IPositionalSpecifier interface {
	objc.IObject
	SetInsertionClassDescription(classDescription IScriptClassDescription)
	Evaluate()
	ObjectSpecifier() ScriptObjectSpecifier
	InsertionContainer() objc.Object
	InsertionIndex() int
	Position() InsertionPosition
	InsertionReplaces() bool
	InsertionKey() string
}

// A specifier for an insertion point in a container relative to another object in the container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier?language=objc
type PositionalSpecifier struct {
	objc.Object
}

func PositionalSpecifierFrom(ptr unsafe.Pointer) PositionalSpecifier {
	return PositionalSpecifier{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PositionalSpecifier) InitWithPositionObjectSpecifier(position InsertionPosition, specifier IScriptObjectSpecifier) PositionalSpecifier {
	rv := objc.Call[PositionalSpecifier](p_, objc.Sel("initWithPosition:objectSpecifier:"), position, objc.Ptr(specifier))
	return rv
}

// Initializes a positional specifier with a given position relative to another given specifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/1416546-initwithposition?language=objc
func PositionalSpecifier_InitWithPositionObjectSpecifier(position InsertionPosition, specifier IScriptObjectSpecifier) PositionalSpecifier {
	return PositionalSpecifierClass.Alloc().InitWithPositionObjectSpecifier(position, specifier)
}

func (pc _PositionalSpecifierClass) Alloc() PositionalSpecifier {
	rv := objc.Call[PositionalSpecifier](pc, objc.Sel("alloc"))
	return rv
}

func PositionalSpecifier_Alloc() PositionalSpecifier {
	return PositionalSpecifierClass.Alloc()
}

func (pc _PositionalSpecifierClass) New() PositionalSpecifier {
	rv := objc.Call[PositionalSpecifier](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPositionalSpecifier() PositionalSpecifier {
	return PositionalSpecifierClass.New()
}

func (p_ PositionalSpecifier) Init() PositionalSpecifier {
	rv := objc.Call[PositionalSpecifier](p_, objc.Sel("init"))
	return rv
}

// Sets the class description for the object or objects to be inserted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/1414707-setinsertionclassdescription?language=objc
func (p_ PositionalSpecifier) SetInsertionClassDescription(classDescription IScriptClassDescription) {
	objc.Call[objc.Void](p_, objc.Sel("setInsertionClassDescription:"), objc.Ptr(classDescription))
}

// Causes the receiver to evaluate its position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/1417035-evaluate?language=objc
func (p_ PositionalSpecifier) Evaluate() {
	objc.Call[objc.Void](p_, objc.Sel("evaluate"))
}

// Returns the object specifier specified at initialization time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/1412839-objectspecifier?language=objc
func (p_ PositionalSpecifier) ObjectSpecifier() ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](p_, objc.Sel("objectSpecifier"))
	return rv
}

// Returns the container in which the new or copied object or objects should be placed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/1414957-insertioncontainer?language=objc
func (p_ PositionalSpecifier) InsertionContainer() objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("insertionContainer"))
	return rv
}

// Returns an insertion index that indicates where the new or copied object or objects should be placed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/1414703-insertionindex?language=objc
func (p_ PositionalSpecifier) InsertionIndex() int {
	rv := objc.Call[int](p_, objc.Sel("insertionIndex"))
	return rv
}

// Returns the insertion position specified at initialization time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/1413865-position?language=objc
func (p_ PositionalSpecifier) Position() InsertionPosition {
	rv := objc.Call[InsertionPosition](p_, objc.Sel("position"))
	return rv
}

// Returns a Boolean value that indicates whether evaluation has been successful and the object to be inserted should actually replace the keyed, indexed object in the insertion container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/1411646-insertionreplaces?language=objc
func (p_ PositionalSpecifier) InsertionReplaces() bool {
	rv := objc.Call[bool](p_, objc.Sel("insertionReplaces"))
	return rv
}

// Returns the key that identifies the relationship into which the new or copied object or objects should be inserted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/1414059-insertionkey?language=objc
func (p_ PositionalSpecifier) InsertionKey() string {
	rv := objc.Call[string](p_, objc.Sel("insertionKey"))
	return rv
}
