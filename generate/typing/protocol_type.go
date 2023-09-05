package typing

import (
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/internal/set"
)

var _ Type = (*ProtocolType)(nil)

// ProtocolType objective-c protocol type
type ProtocolType struct {
	Name   string          // the objc type name
	GName  string          // Go name, usually is objc type name without prefix 'NS'
	Module *modules.Module // object-c module name
}

func (p *ProtocolType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	if receiveFromObjc {
		return FullGoName(*p.Module, p.GoWrapperName(), *currentModule)
	} else {
		return FullGoName(*p.Module, p.GoInterfaceName(), *currentModule)
	}
}

func (p *ProtocolType) ObjcName() string {
	return "id<" + p.Name + ">"
}

func (p *ProtocolType) CName() string {
	return "void *"
}

func (p *ProtocolType) CSignature() string {
	return "void *"
}

func (p *ProtocolType) GoImports() set.Set[string] {
	return set.New("github.com/progrium/darwinkit/objc", "github.com/progrium/darwinkit/macos/"+p.Module.Package)
}

func (p *ProtocolType) DeclareModule() *modules.Module {
	return p.Module
}

func (p *ProtocolType) GoInterfaceName() string {
	return "P" + p.GName
}

func (p *ProtocolType) GoStructName() string {
	return p.GName
}

func (p *ProtocolType) GoWrapperName() string {
	if p.GName == "Object" {
		return p.GName
	}
	return p.GName + "Object"
}
