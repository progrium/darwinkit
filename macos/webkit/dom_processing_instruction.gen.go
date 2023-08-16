// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMProcessingInstruction] class.
var DOMProcessingInstructionClass = _DOMProcessingInstructionClass{objc.GetClass("DOMProcessingInstruction")}

type _DOMProcessingInstructionClass struct {
	objc.Class
}

// An interface definition for the [DOMProcessingInstruction] class.
type IDOMProcessingInstruction interface {
	IDOMCharacterData
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domprocessinginstruction?language=objc
type DOMProcessingInstruction struct {
	DOMCharacterData
}

func DOMProcessingInstructionFrom(ptr unsafe.Pointer) DOMProcessingInstruction {
	return DOMProcessingInstruction{
		DOMCharacterData: DOMCharacterDataFrom(ptr),
	}
}

func (dc _DOMProcessingInstructionClass) Alloc() DOMProcessingInstruction {
	rv := objc.Call[DOMProcessingInstruction](dc, objc.Sel("alloc"))
	return rv
}

func DOMProcessingInstruction_Alloc() DOMProcessingInstruction {
	return DOMProcessingInstructionClass.Alloc()
}

func (dc _DOMProcessingInstructionClass) New() DOMProcessingInstruction {
	rv := objc.Call[DOMProcessingInstruction](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMProcessingInstruction() DOMProcessingInstruction {
	return DOMProcessingInstructionClass.New()
}

func (d_ DOMProcessingInstruction) Init() DOMProcessingInstruction {
	rv := objc.Call[DOMProcessingInstruction](d_, objc.Sel("init"))
	return rv
}
