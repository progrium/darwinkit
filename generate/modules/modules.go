package modules

import "strings"

// Module Objective-c module
type Module struct {
	Name     string   // module name
	Title    string   // readable name
	Package  string   // go package for this module
	Header   string   // objc header file
	Prefixes []string // symbol prefixes
}

func (m *Module) String() string {
	return m.Name
}

func Get(moduleName string) *Module {
	for _, module := range All {
		if strings.EqualFold(moduleName, module.Name) ||
			moduleName == module.Title ||
			moduleName == module.Package {
			return &module
		}
	}
	return nil
}

var All = []Module{
	//{"objectivec", "Objective-C Runtime", "objc", "objc/runtime.h", []string{}},
	{"Foundation", "Foundation", "foundation", "Foundation/Foundation.h", []string{"NS"}},
	{"AppKit", "AppKit", "appkit", "Appkit/Appkit.h", []string{"NS"}},
	{"UIKit", "UIKit", "uikit", "UIKit/UIKit.h", []string{}},
	{"UniformTypeIdentifiers", "Uniform Type Identifiers", "uniformtypeidentifiers", "UniformTypeIdentifiers/UniformTypeIdentifiers.h", []string{"UT"}},
	{"WebKit", "WebKit", "webkit", "WebKit/WebKit.h", []string{"WK"}},
	{"FileProvider", "File Provider", "fileprovider", "FileProvider/FileProvider.h", []string{"NS"}},
	{"Quartz", "Quartz", "quartz", "Quartz/Quartz.h", []string{"IK", "kQC", "kQuartz", "QC", "IK_"}},
	{"SecurityInterface", "Security Interface", "securityinterface", "SecurityInterface/SecurityInterface.h", []string{"SF"}},
	{"IOBluetooth", "IOBluetooth", "iobluetooth", "IOBluetooth/IOBluetooth.h", []string{"kIOBluetooth", "kBluetooth", "IOBluetooth", "Bluetooth"}},
	{"CoreGraphics", "Core Graphics", "coregraphics", "CoreGraphics/CoreGraphics.h", []string{"CG", "kCG"}},
	{"CoreFoundation", "Core Foundation", "corefoundation", "CoreFoundation/CoreFoundation.h", []string{"CF", "kCF"}},
	{"QuartzCore", "QuartzCore", "quartzcore", "QuartzCore/QuartzCore.h", []string{"kCA"}},
}
