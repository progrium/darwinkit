// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMNode] class.
var DOMNodeClass = _DOMNodeClass{objc.GetClass("DOMNode")}

type _DOMNodeClass struct {
	objc.Class
}

// An interface definition for the [DOMNode] class.
type IDOMNode interface {
	IDOMObject
	LineBoxRects() []objc.Object
	BoundingBox() foundation.Rect
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domnode?language=objc
type DOMNode struct {
	DOMObject
}

func DOMNodeFrom(ptr unsafe.Pointer) DOMNode {
	return DOMNode{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMNodeClass) Alloc() DOMNode {
	rv := objc.Call[DOMNode](dc, objc.Sel("alloc"))
	return rv
}

func DOMNode_Alloc() DOMNode {
	return DOMNodeClass.Alloc()
}

func (dc _DOMNodeClass) New() DOMNode {
	rv := objc.Call[DOMNode](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMNode() DOMNode {
	return DOMNodeClass.New()
}

func (d_ DOMNode) Init() DOMNode {
	rv := objc.Call[DOMNode](d_, objc.Sel("init"))
	return rv
}

// Returns the rectangles that bound each line of text in the node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domnode/1536778-lineboxrects?language=objc
func (d_ DOMNode) LineBoxRects() []objc.Object {
	rv := objc.Call[[]objc.Object](d_, objc.Sel("lineBoxRects"))
	return rv
}

// Returns a rectangle that bounds the onscreen rendering of the node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domnode/1536392-boundingbox?language=objc
func (d_ DOMNode) BoundingBox() foundation.Rect {
	rv := objc.Call[foundation.Rect](d_, objc.Sel("boundingBox"))
	return rv
}
