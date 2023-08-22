// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetWriterInputGroup] class.
var AssetWriterInputGroupClass = _AssetWriterInputGroupClass{objc.GetClass("AVAssetWriterInputGroup")}

type _AssetWriterInputGroupClass struct {
	objc.Class
}

// An interface definition for the [AssetWriterInputGroup] class.
type IAssetWriterInputGroup interface {
	IMediaSelectionGroup
	DefaultInput() AssetWriterInput
	Inputs() []AssetWriterInput
}

// A group of inputs with tracks that are mutually exclusive to each other for playback or processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputgroup?language=objc
type AssetWriterInputGroup struct {
	MediaSelectionGroup
}

func AssetWriterInputGroupFrom(ptr unsafe.Pointer) AssetWriterInputGroup {
	return AssetWriterInputGroup{
		MediaSelectionGroup: MediaSelectionGroupFrom(ptr),
	}
}

func (a_ AssetWriterInputGroup) InitWithInputsDefaultInput(inputs []IAssetWriterInput, defaultInput IAssetWriterInput) AssetWriterInputGroup {
	rv := objc.Call[AssetWriterInputGroup](a_, objc.Sel("initWithInputs:defaultInput:"), inputs, objc.Ptr(defaultInput))
	return rv
}

// Creates a group for the asset writer inputs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputgroup/1389502-initwithinputs?language=objc
func NewAssetWriterInputGroupWithInputsDefaultInput(inputs []IAssetWriterInput, defaultInput IAssetWriterInput) AssetWriterInputGroup {
	instance := AssetWriterInputGroupClass.Alloc().InitWithInputsDefaultInput(inputs, defaultInput)
	instance.Autorelease()
	return instance
}

func (ac _AssetWriterInputGroupClass) AssetWriterInputGroupWithInputsDefaultInput(inputs []IAssetWriterInput, defaultInput IAssetWriterInput) AssetWriterInputGroup {
	rv := objc.Call[AssetWriterInputGroup](ac, objc.Sel("assetWriterInputGroupWithInputs:defaultInput:"), inputs, objc.Ptr(defaultInput))
	return rv
}

// Returns a new group for the asset writer inputs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputgroup/1426655-assetwriterinputgroupwithinputs?language=objc
func AssetWriterInputGroup_AssetWriterInputGroupWithInputsDefaultInput(inputs []IAssetWriterInput, defaultInput IAssetWriterInput) AssetWriterInputGroup {
	return AssetWriterInputGroupClass.AssetWriterInputGroupWithInputsDefaultInput(inputs, defaultInput)
}

func (ac _AssetWriterInputGroupClass) Alloc() AssetWriterInputGroup {
	rv := objc.Call[AssetWriterInputGroup](ac, objc.Sel("alloc"))
	return rv
}

func AssetWriterInputGroup_Alloc() AssetWriterInputGroup {
	return AssetWriterInputGroupClass.Alloc()
}

func (ac _AssetWriterInputGroupClass) New() AssetWriterInputGroup {
	rv := objc.Call[AssetWriterInputGroup](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetWriterInputGroup() AssetWriterInputGroup {
	return AssetWriterInputGroupClass.New()
}

func (a_ AssetWriterInputGroup) Init() AssetWriterInputGroup {
	rv := objc.Call[AssetWriterInputGroup](a_, objc.Sel("init"))
	return rv
}

// The default input for the group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputgroup/1389698-defaultinput?language=objc
func (a_ AssetWriterInputGroup) DefaultInput() AssetWriterInput {
	rv := objc.Call[AssetWriterInput](a_, objc.Sel("defaultInput"))
	return rv
}

// The inputs with tracks that are mutually exclusive to each other for playback or processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputgroup/1388226-inputs?language=objc
func (a_ AssetWriterInputGroup) Inputs() []AssetWriterInput {
	rv := objc.Call[[]AssetWriterInput](a_, objc.Sel("inputs"))
	return rv
}
