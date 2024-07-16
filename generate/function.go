package generate

import (
	"fmt"
	"log"
	"strings"

	"github.com/progrium/darwinkit/generate/codegen"
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/generate/typing"
)

// this keeps track of types that we know we can't handle.
var unhandledStructTypes = map[string]bool{
	"CFAllocatorContext":                true,
	"CFArrayCallBacks":                  true,
	"CFBagCallBacks":                    true,
	"CFBinaryHeapCallBacks":             true,
	"CFBinaryHeapCompareContext":        true,
	"CFCalendarIdentifier":              true,
	"CFDateFormatterKey":                true,
	"CFDictionaryKeyCallBacks":          true,
	"CFDictionaryValueCallBacks":        true,
	"CFErrorDomain":                     true,
	"CFFileDescriptorContext":           true,
	"CFLocaleIdentifier":                true,
	"CFLocaleKey":                       true,
	"CFMachPortContext":                 true,
	"CFMessagePortContext":              true,
	"CFMessagePortInvalidationCallBack": true,
	"CFNumberFormatterKey":              true,
	"CFRange":                           true,
	"CFRunLoopMode":                     true,
	"CFRunLoopObserverContext":          true,
	"CFRunLoopSourceContext":            true,
	"CFRunLoopTimerContext":             true,
	"CFSetCallBacks":                    true,
	"CFSocketContext":                   true,
	"CFSocketSignature":                 true,
	"CFStreamPropertyKey":               true,
	"CFStringInlineBuffer":              true,
	"CFSwappedFloat32":                  true,
	"CFSwappedFloat64":                  true,
	"CFTreeContext":                     true,
	"CFUUIDBytes":                       true,
	"dispatch_queue_t":                  true, // for return values, not parameters
	"va_list":                           true,

	"MTLIndirectCommandBufferExecutionRange": true,
	"MTLPackedFloat3":                        true,
}

func (db *Generator) ToFunction(fw string, sym Symbol) *codegen.Function {
	// these functions have known declparse failures
	knownIssues := map[string]bool{

		"CFCharacterSetIsLongCharacterMember":      true, // "UTF32Char theChar"
		"CFStringGetCharactersPtr":                 true, // return handling
		"CFStringCreateWithFormat":                 true, // format param
		"CFStringAppendFormat":                     true, // format param
		"CFFileSecurityCopyAccessControlList":      true, // "acl_t accessControlList"
		"CFFileSecuritySetAccessControlList":       true, // "acl_t accessControlList"
		"CFMachPortGetInvalidationCallBack":        true, // converting return value to function pointer
		"CFMachPortGetPort":                        true, // handling return of mach_port_t
		"CFStringAppendPascalString":               true, // "ConstStr255Param pStr"
		"CFStringCreateWithPascalString":           true, // StringPtr
		"CFStringCreateWithPascalStringNoCopy":     true, // StringPtr
		"CFStringGetLongCharacterForSurrogatePair": true, // "UTF16Char surrogateHigh, UTF16Char surrogateLow"
		"CFStringGetPascalString":                  true, // StringPtr
		"CFStringGetPascalStringPtr":               true, // StringPtr
		"CFStringGetSurrogatePairForLongCharacter": true, // "UTF32Char character"
		"CGColorSpaceCreateIndexed":                true, // "const unsigned char *"
		"CGPDFArrayGetName":                        true, // "const char * _Nullable *"
		"CGPDFDictionaryGetName":                   true, // "const char *key, const char * _Nullable *"
		"CGPDFScannerPopName":                      true, // "const char * _Nullable *"

		"MTLSizeMake": true, // duplicate symbol issue
	}
	if knownIssues[sym.Name] {
		_, err := sym.Parse(db.Platform)
		log.Printf("skipping function %s %s because of known issue: decl='%s' err='%v'\n", fw, sym.Name, sym.Declaration, err)
		return nil
	}
	typ := db.TypeFromSymbol(sym)
	fntyp, ok := typ.(*typing.FunctionType)
	if !ok {
		return nil
	}
	fn := &codegen.Function{
		Name:        sym.Name,
		Deprecated:  sym.Deprecated,
		GoName:      modules.TrimPrefix(sym.Name),
		Description: sym.Description,
		DocURL:      sym.DocURL(),
		Type:        fntyp,
	}
	// temporary skip for things deprecated in 14.0
	// check if macOS platform is DeprecatedAt 14.0
	for _, p := range sym.Platforms {
		if p.Name == "macOS" && p.Deprecated {
			fn.Deprecated = true
		}
	}

	// populate params:
	for _, p := range fntyp.Parameters {
		if p.Type == nil {
			fmt.Printf("skipping %s: %s because of nil type\n", sym.Name, p.Name)
			return nil
		}
		// skip pointers to ref types (for now)
		if pt, ok := p.Type.(*typing.PointerType); ok {
			if _, ok := pt.Type.(*typing.RefType); ok {
				fmt.Printf("skipping %s: %s because of pointer to ref type\n", sym.Name, p.Name)
				return nil
			}
		}
		// skip DispatchType (for now)
		if _, ok := p.Type.(*typing.DispatchType); ok {
			fmt.Printf("skipping %s: %s because of DispatchType\n", sym.Name, p.Name)
			return nil
		}

		// detect if we have a CFRange, CFTypeRef, or CFStringInlineBuffer as an argument type
		if unhandledStructType(p.Type.ObjcName()) {
			fmt.Printf("skipping %s: %s because of unhandled struct type %s\n", sym.Name, p.Name, p.Type.ObjcName())
			return nil
		}

		fn.Parameters = append(fn.Parameters, &codegen.Param{
			Name: p.Name,
			Type: p.Type,
		})
	}
	if unhandledStructType(fntyp.ReturnType.ObjcName()) {
		fmt.Printf("skipping %s because of unhandled struct type %s\n", sym.Name, fntyp.ReturnType.ObjcName())
		return nil
	}

	// we (unfortuantely) don't handle array returns cleanly yet:
	if _, ok := fntyp.ReturnType.(*typing.ArrayType); ok {
		fmt.Printf("skipping %s because of array return type\n", sym.Name)
		return nil
	}
	// populate return type
	if fntyp.ReturnType != nil {
		fn.ReturnType = fntyp.ReturnType
	}

	return fn

}

func unhandledStructType(t string) bool {
	return unhandledStructTypes[t] || unhandledStructTypes[strings.TrimSuffix(t, "*")]
}

/*
 */
