// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetVariantQualifier] class.
var AssetVariantQualifierClass = _AssetVariantQualifierClass{objc.GetClass("AVAssetVariantQualifier")}

type _AssetVariantQualifierClass struct {
	objc.Class
}

// An interface definition for the [AssetVariantQualifier] class.
type IAssetVariantQualifier interface {
	objc.IObject
}

// An object that represents an HTTP Live Streaming asset variant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantqualifier?language=objc
type AssetVariantQualifier struct {
	objc.Object
}

func AssetVariantQualifierFrom(ptr unsafe.Pointer) AssetVariantQualifier {
	return AssetVariantQualifier{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetVariantQualifierClass) AssetVariantQualifierWithVariant(variant IAssetVariant) AssetVariantQualifier {
	rv := objc.Call[AssetVariantQualifier](ac, objc.Sel("assetVariantQualifierWithVariant:"), objc.Ptr(variant))
	return rv
}

// Creates a variant qualifier with an asset variant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantqualifier/3750228-assetvariantqualifierwithvariant?language=objc
func AssetVariantQualifier_AssetVariantQualifierWithVariant(variant IAssetVariant) AssetVariantQualifier {
	return AssetVariantQualifierClass.AssetVariantQualifierWithVariant(variant)
}

func (ac _AssetVariantQualifierClass) AssetVariantQualifierWithPredicate(predicate foundation.IPredicate) AssetVariantQualifier {
	rv := objc.Call[AssetVariantQualifier](ac, objc.Sel("assetVariantQualifierWithPredicate:"), objc.Ptr(predicate))
	return rv
}

// Creates a variant qualifier with a predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantqualifier/3750227-assetvariantqualifierwithpredica?language=objc
func AssetVariantQualifier_AssetVariantQualifierWithPredicate(predicate foundation.IPredicate) AssetVariantQualifier {
	return AssetVariantQualifierClass.AssetVariantQualifierWithPredicate(predicate)
}

func (ac _AssetVariantQualifierClass) Alloc() AssetVariantQualifier {
	rv := objc.Call[AssetVariantQualifier](ac, objc.Sel("alloc"))
	return rv
}

func AssetVariantQualifier_Alloc() AssetVariantQualifier {
	return AssetVariantQualifierClass.Alloc()
}

func (ac _AssetVariantQualifierClass) New() AssetVariantQualifier {
	rv := objc.Call[AssetVariantQualifier](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetVariantQualifier() AssetVariantQualifier {
	return AssetVariantQualifierClass.New()
}

func (a_ AssetVariantQualifier) Init() AssetVariantQualifier {
	rv := objc.Call[AssetVariantQualifier](a_, objc.Sel("init"))
	return rv
}

// Creates a predicate with a width and operator type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantqualifier/3857562-predicateforpresentationwidth?language=objc
func (ac _AssetVariantQualifierClass) PredicateForPresentationWidthOperatorType(width float64, operatorType foundation.PredicateOperatorType) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](ac, objc.Sel("predicateForPresentationWidth:operatorType:"), width, operatorType)
	return rv
}

// Creates a predicate with a width and operator type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantqualifier/3857562-predicateforpresentationwidth?language=objc
func AssetVariantQualifier_PredicateForPresentationWidthOperatorType(width float64, operatorType foundation.PredicateOperatorType) foundation.Predicate {
	return AssetVariantQualifierClass.PredicateForPresentationWidthOperatorType(width, operatorType)
}

// Creates a predicate with a channel count, media selection option, and operator type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantqualifier/3750230-predicateforchannelcount?language=objc
func (ac _AssetVariantQualifierClass) PredicateForChannelCountMediaSelectionOptionOperatorType(channelCount int, mediaSelectionOption IMediaSelectionOption, operatorType foundation.PredicateOperatorType) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](ac, objc.Sel("predicateForChannelCount:mediaSelectionOption:operatorType:"), channelCount, objc.Ptr(mediaSelectionOption), operatorType)
	return rv
}

// Creates a predicate with a channel count, media selection option, and operator type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantqualifier/3750230-predicateforchannelcount?language=objc
func AssetVariantQualifier_PredicateForChannelCountMediaSelectionOptionOperatorType(channelCount int, mediaSelectionOption IMediaSelectionOption, operatorType foundation.PredicateOperatorType) foundation.Predicate {
	return AssetVariantQualifierClass.PredicateForChannelCountMediaSelectionOptionOperatorType(channelCount, mediaSelectionOption, operatorType)
}

// Creates a predicate with a height and operator type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantqualifier/3857561-predicateforpresentationheight?language=objc
func (ac _AssetVariantQualifierClass) PredicateForPresentationHeightOperatorType(height float64, operatorType foundation.PredicateOperatorType) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](ac, objc.Sel("predicateForPresentationHeight:operatorType:"), height, operatorType)
	return rv
}

// Creates a predicate with a height and operator type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantqualifier/3857561-predicateforpresentationheight?language=objc
func AssetVariantQualifier_PredicateForPresentationHeightOperatorType(height float64, operatorType foundation.PredicateOperatorType) foundation.Predicate {
	return AssetVariantQualifierClass.PredicateForPresentationHeightOperatorType(height, operatorType)
}
