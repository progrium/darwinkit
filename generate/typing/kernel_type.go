package typing

import (
	"unicode"

	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/internal/set"
)

func GetKernelType(typeName string) (Type, bool) {
	for _, name := range []string{
		"vector_int2",
		"vector_int4",
		"vector_float2",
		"vector_float3",
		"vector_float4",
		"vector_double2",
		"vector_double3",
		"vector_double4",
		"matrix_float2x2",
		"matrix_float3x3",
		"matrix_float4x4",
		"matrix_float4x3",
		"matrix_double4x4",
	} {
		if typeName == name {
			return &KernelType{ObjcName_: typeName}, true
		}
	}
	return nil, false
}

type KernelType struct {
	ObjcName_ string // objc type name
}

func (k *KernelType) GoImports() set.Set[string] {
	return set.New("github.com/progrium/macdriver/kernel")
}

func (k *KernelType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	r := []rune(k.ObjcName_)
	r[0] = unicode.ToUpper(r[0])
	return FullGoName(*k.DeclareModule(), string(r), *currentModule)
}

func (k *KernelType) ObjcName() string {
	return k.ObjcName_
}

func (k *KernelType) DeclareModule() *modules.Module {
	return modules.Get("kernel")
}
