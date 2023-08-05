package typing

import (
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/internal/set"
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

func (p *ProtocolType) GoImports() set.Set[string] {
	return set.New("github.com/progrium/macdriver/objc", "github.com/progrium/macdriver/macos/"+p.Module.Package)
}

func (p *ProtocolType) DeclareModule() *modules.Module {
	return p.Module
}

// GoInterfaceName return the go wrapper interface name
func (p *ProtocolType) GoInterfaceName() string {
	return "I" + p.GName
}

// GoStructName return the go wrapper struct name
func (p *ProtocolType) GoStructName() string {
	return p.GName
}

func (p *ProtocolType) GoWrapperName() string {
	if p.GName == "Object" {
		return p.GName
	}
	return p.GName + "Wrapper"
}
