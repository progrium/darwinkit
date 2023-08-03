package typing

import (
	"strings"

	"github.com/progrium/macdriver/internal/set"
)

// Module Objective-c module
type Module struct {
	Name    string // module name
	DesName string // readable name
	Package string // go package for this module
	Header  string // objc header file
}

func (m *Module) String() string {
	return m.Name
}

var ObjCRuntime = &Module{"objectivec", "Objective-C Runtime", "objc", "objc/runtime.h"}
var Foundation = &Module{"Foundation", "Foundation", "foundation", "Foundation/Foundation.h"}
var AppKit = &Module{"AppKit", "AppKit", "appkit", "Appkit/Appkit.h"}
var UniformTypeIdentifiers = &Module{"UniformTypeIdentifiers", "Uniform Type Identifiers", "uniformtypeidentifiers", "UniformTypeIdentifiers/UniformTypeIdentifiers.h"}
var WebKit = &Module{"WebKit", "WebKit", "webkit", "WebKit/WebKit.h"}
var FileProvider = &Module{"FileProvider", "File Provider", "fileprovider", "FileProvider/FileProvider.h"}
var Quartz = &Module{"Quartz", "Quartz", "quartz", "Quartz/Quartz.h"}
var SecurityInterface = &Module{"SecurityInterface", "Security Interface", "securityinterface", "SecurityInterface/SecurityInterface.h"}
var IOBluetooth = &Module{"IOBluetooth", "IOBluetooth", "iobluetooth", "IOBluetooth/IOBluetooth.h"}
var CoreGraphics = &Module{"CoreGraphics", "Core Graphics", "coregraphics", "CoreGraphics/CoreGraphics.h"}
var CoreFoundation = &Module{"CoreFoundation", "Core Foundation", "corefoundation", "CoreFoundation/CoreFoundation.h"}
var QuartzCore = &Module{"QuartzCore", "QuartzCore", "quartzcore", "QuartzCore/QuartzCore.h"}

var allModules = []*Module{ObjCRuntime, Foundation, AppKit, UniformTypeIdentifiers, WebKit, FileProvider,
	Quartz, SecurityInterface, IOBluetooth, CoreGraphics, QuartzCore}

func FindModule(moduleName string) *Module {
	if moduleName == "UIKit" {
		moduleName = "AppKit"
	}
	if moduleName == "uikit" {
		moduleName = "appkit"
	}
	lcName := strings.ToLower(moduleName)
	for _, module := range allModules {
		if lcName == strings.ToLower(module.Name) || moduleName == module.DesName {
			return module
		}
	}
	return nil
}

// Type interface for all type
type Type interface {
	// GoName Go type name
	GoName(currentModule *Module, receiveFromObjc bool) string
	// ObjcName Objective-c type name
	ObjcName() string

	// GoImports go imports for this type
	GoImports() set.Set[string]

	// DeclareModule the module of this type. return nil if is a built in type
	DeclareModule() *Module
}

func FullGoName(module Module, name string, currentModule Module) string {
	if module == currentModule {
		return name
	}
	return module.Package + "." + name
}

func PrependPackage(module Module, s string, currentModule Module) string {
	if module == currentModule {
		return s
	}
	return module.Package + "." + s
}
