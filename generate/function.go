package generate

import (
	"fmt"

	"github.com/progrium/darwinkit/generate/codegen"
	"github.com/progrium/darwinkit/generate/typing"
)

func (db *Generator) ToFunction(fw string, sym Symbol) *codegen.Function {
	// these functions have known declparse failures
	knownIssues := map[string]bool{
		"AUListenerAddParameter":                   true,
		"MIDIClientCreate":                         true,
		"AudioDeviceGetProperty":                   true,
		"AudioFileComponentOpenWithCallbacks":      true,
		"AudioObjectSetPropertyData":               true,
		"AudioUnitRemoveRenderNotify":              true,
		"CFBagRemoveValue":                         true,
		"CFBinaryHeapGetMinimumIfPresent":          true,
		"CFURLCopyResourcePropertyForKey":          true,
		"CGBitmapContextCreate":                    true,
		"CGBitmapContextCreateWithData":            true,
		"CGColorSpaceCreateWithPlatformColorSpace": true,
		"CGConvertColorDataWithFormat":             true,
		"CGDataConsumerCreate":                     true,
		"CGDataProviderCreateDirect":               true,
		"CGDataProviderCreateSequential":           true,
		"CGDataProviderCreateWithData":             true,
		"CGDisplayRegisterReconfigurationCallback": true,
		"CGDisplayRemoveReconfigurationCallback":   true,
		"CGEventPostToPSN":                         true,
		"CGEventTapCreate":                         true,
		"CGEventTapCreateForPSN":                   true,
		"CGEventTapCreateForPid":                   true,
		"CGFontCreateWithPlatformFont":             true,
		"CGFunctionCreate":                         true,
		"CGPDFArrayApplyBlock":                     true,
		"CGPDFDictionaryApplyBlock":                true,
		"CGPDFDictionaryApplyFunction":             true,
		"CGPDFObjectGetValue":                      true,
		"CGPDFScannerCreate":                       true,
		"CGPSConverterCreate":                      true,
		"CGPathApply":                              true,
		"CGPatternCreate":                          true,
		"CGRegisterScreenRefreshCallback":          true,
		"CGScreenRegisterMoveCallback":             true,
		"CGScreenUnregisterMoveCallback":           true,
		"CGUnregisterScreenRefreshCallback":        true,
		"CMBlockBufferAppendMemoryBlock":           true,
		"CMIOStreamClockCreate":                    true,
		"MIDIEndpointSetRefCons":                   true,
		"NSBeginAlertSheet":                        true,
		"NSBeginCriticalAlertSheet":                true,
		"NSShowAnimationEffect":                    true,
		"NSZoneFromPointer":                        true,
		"OBEXAddHTTPHeader":                        true,
		"OBEXAddUserDefinedHeader":                 true,
	}
	if knownIssues[sym.Name] {
		return nil
	}
	typ := db.TypeFromSymbol(sym)
	fntyp, ok := typ.(*typing.FunctionType)
	if !ok {
		return nil
	}
	fn := &codegen.Function{
		Name:        sym.Name,
		GoName:      sym.Name,
		Description: sym.Description,
		DocURL:      sym.DocURL(),
		Type:        fntyp,
	}
	// populate params:
	for _, p := range fntyp.Parameters {
		if p.Type == nil {
			fmt.Printf("skipping %s: %s because of nil type\n", sym.Name, p.Name)
			continue
		}
		fn.Parameters = append(fn.Parameters, &codegen.Param{
			Name: p.Name,
			Type: p.Type,
		})
	}
	// populate return type
	if fntyp.ReturnType != nil {
		fn.ReturnType = fntyp.ReturnType
	}

	return fn

}
