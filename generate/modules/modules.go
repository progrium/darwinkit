package modules

import (
	"log"
	"strings"
	"unicode"
)

// Module Objective-c module
type Module struct {
	Name     string   // module name
	Title    string   // readable name
	Package  string   // go package for this module
	Header   string   // objc header file
	Prefixes []string // symbol prefixes

	Flags ModuleCodeGenFlags // controls code generation behavior
}

type ModuleCodeGenFlags struct {
	GenerateStructs   bool
	GenerateFunctions bool
}

func (m *Module) String() string {
	return m.Name
}

func Get(moduleName string) *Module {
	for _, module := range All {
		if strings.EqualFold(moduleName, module.Name) ||
			strings.EqualFold(moduleName, module.Title) ||
			moduleName == module.Package {
			return &module
		}
	}
	log.Panicf("module not found: %s", moduleName)
	return nil
}

func TrimPrefix(symbolName string) string {
	for _, m := range All {
		for _, prefix := range m.Prefixes {
			if len(prefix) == len(symbolName) {
				continue
			}
			if strings.HasPrefix(symbolName, prefix) && unicode.IsUpper(rune(symbolName[len(prefix)])) {
				name := strings.TrimPrefix(symbolName, prefix)
				if strings.HasPrefix(symbolName, "k") {
					return "K" + name
				}
				return name
			}
		}
	}
	return symbolName
}

// modules that will cause types to become IObject/unsafe.Pointer or more primitive type
func CanAbstractModuleCoupling(in string, mod string) bool {
	mods, ok := map[string][]string{
		"foundation": []string{"appkit", "coreimage", "corespotlight", "fileprovider", "gameplaykit", "iobluetooth", "uti"},
		"appkit":     []string{"spritekit", "cloudkit"},
		"coreimage":  []string{"appkit"},
		"coredata":   []string{"corespotlight"},
		"cloudkit":   []string{"corelocation"},
	}[in]
	if !ok {
		return false
	}
	for _, m := range mods {
		if m == mod {
			return true
		}
	}
	return false
}

// modules that will cause methods/props to be skipped
func CanSkipModuleCoupling(in string, mod string) bool {
	mods, ok := map[string][]string{
		"foundation": []string{"webkit", "scenekit", "quartzcore", "corelocation", "cloudkit"},
		"appkit":     []string{},
		"coreimage":  []string{"avfoundation", "quartz"},
		"quartzcore": []string{"scenekit"},
	}[in]
	if !ok {
		return false
	}
	for _, m := range mods {
		if m == mod {
			return true
		}
	}
	return false
}

func CanIgnoreNotFound(p any) bool {
	mod := strings.TrimPrefix(p.(string), "module not found: ")
	for _, m := range []string{
		"Security",
		"JavaScriptCore",
		"ImageCaptureCore",
		"User Notifications",
		"Core Services",
		"XPC",
		"MapKit",
		"Intents",
		"QuickLook",
		"force feedback",
		"opengl es",
		"ColorSync",
	} {
		if strings.EqualFold(m, mod) {
			return true
		}
	}
	return false
}

var DefaultFlags = ModuleCodeGenFlags{
	GenerateStructs:   false,
	GenerateFunctions: false,
}

var All = []Module{
	{"objectivec", "Objective-C Runtime", "objc", "objc/runtime.h", []string{}, DefaultFlags},
	{"dispatch", "Dispatch", "dispatch", "dispatch/dispatch.h", []string{}, DefaultFlags},
	{"Kernel", "Kernel", "kernel", "Kernel/Kernel.h", []string{}, DefaultFlags},

	{"Foundation", "Foundation", "foundation", "Foundation/Foundation.h", []string{"NS"}, DefaultFlags},
	{"AppKit", "AppKit", "appkit", "AppKit/AppKit.h", []string{"NS"}, DefaultFlags},
	{"UIKit", "UIKit", "uikit", "UIKit/UIKit.h", []string{"NS"}, DefaultFlags},
	{"UniformTypeIdentifiers", "Uniform Type Identifiers", "uti", "UniformTypeIdentifiers/UniformTypeIdentifiers.h", []string{"UT"}, DefaultFlags},
	{"WebKit", "WebKit", "webkit", "WebKit/WebKit.h", []string{"WK"}, DefaultFlags},
	{"FileProvider", "File Provider", "fileprovider", "FileProvider/FileProvider.h", []string{"NS"}, DefaultFlags},
	{"Quartz", "Quartz", "quartz", "Quartz/Quartz.h", []string{"IK", "kQC", "kQuartz", "QC", "IK_"}, DefaultFlags},
	{"SecurityInterface", "Security Interface", "securityinterface", "SecurityInterface/SecurityInterface.h", []string{"SF"}, DefaultFlags},
	{"IOBluetooth", "IOBluetooth", "iobluetooth", "IOBluetooth/IOBluetooth.h", []string{"kIOBluetooth", "kBluetooth", "IOBluetooth", "Bluetooth", "kFTS", "kOBEX"}, DefaultFlags},
	{"CoreGraphics", "Core Graphics", "coregraphics", "CoreGraphics/CoreGraphics.h", []string{"CG", "kCG"}, DefaultFlags},
	{"CoreFoundation", "Core Foundation", "corefoundation", "CoreFoundation/CoreFoundation.h", []string{"CF", "kCF"}, DefaultFlags},
	{"QuartzCore", "Core Animation", "quartzcore", "QuartzCore/QuartzCore.h", []string{"kCA", "CA"}, DefaultFlags},
	{"Vision", "Vision", "vision", "Vision/Vision.h", []string{"VN"}, DefaultFlags},
	{"CoreSpotlight", "Core Spotlight", "corespotlight", "CoreSpotlight/CoreSpotlight.h", []string{"CS"}, DefaultFlags},
	{"CoreAudioKit", "Core Audio Kit", "coreaudiokit", "CoreAudioKit/CoreAudioKit.h", []string{"CA", "AU"}, DefaultFlags},
	{"CoreML", "Core ML", "coreml", "CoreML/CoreML.h", []string{"ML"}, DefaultFlags},
	{"CoreData", "Core Data", "coredata", "CoreData/CoreData.h", []string{"NS"}, DefaultFlags},
	{"CoreMediaIO", "Core Media I/O", "coremediaio", "CoreMediaIO/CMIOHardware.h", []string{"CMIO"}, DefaultFlags},
	{"CoreMedia", "Core Media", "coremedia", "CoreMedia/CoreMedia.h", []string{"CM"}, DefaultFlags},
	{"CoreImage", "Core Image", "coreimage", "CoreImage/CoreImage.h", []string{"CI"}, DefaultFlags},
	{"CoreMIDI", "Core MIDI", "coremidi", "CoreMIDI/CoreMIDI.h", []string{"MIDI", "kMIDI"}, DefaultFlags},
	{"CoreVideo", "Core Video", "corevideo", "CoreVideo/CoreVideo.h", []string{"CV", "kCV"}, DefaultFlags},
	{"CloudKit", "Cloud Kit", "cloudkit", "CloudKit/CloudKit.h", []string{"CK"}, DefaultFlags},
	{"AudioToolbox", "Audio Toolbox", "audiotoolbox", "AudioToolbox/AudioToolbox.h", []string{"AU"}, DefaultFlags},
	{"CoreAudio", "Core Audio", "coreaudio", "CoreAudio/CoreAudio.h", []string{"Audio", "kAudio"}, DefaultFlags},
	{"CoreAudioTypes", "Core Audio Types", "coreaudiotypes", "CoreAudio/CoreAudioTypes.h", []string{"AV"}, DefaultFlags},
	{"CoreLocation", "Core Location", "corelocation", "CoreLocation/CoreLocation.h", []string{"CL"}, DefaultFlags},
	{"Contacts", "Contacts", "contacts", "Contacts/Contacts.h", []string{"CN"}, DefaultFlags},
	{"ContactsUI", "Contacts UI", "contactsui", "ContactsUI/ContactsUI.h", []string{"CN"}, DefaultFlags},
	{"ImageIO", "Image I/O", "imageio", "ImageIO/ImageIO.h", []string{"CG", "kCG", "kCF"}, DefaultFlags},
	{"AVFAudio", "AVFAudio", "avfaudio", "AVFAudio/AVFAudio.h", []string{"AVAudio"}, DefaultFlags},
	{"AVFoundation", "AVFoundation", "avfoundation", "AVFoundation/AVFoundation.h", []string{"AV"}, DefaultFlags},
	{"AVKit", "AVKit", "avkit", "AVKit/AVKit.h", []string{"AV"}, DefaultFlags},
	{"GameplayKit", "GameplayKit", "gameplaykit", "GameplayKit/GameplayKit.h", []string{"GK"}, DefaultFlags},
	{"SystemConfiguration", "System Configuration", "sysconfig", "SystemConfiguration/SystemConfiguration.h", []string{"SC", "kSC"}, DefaultFlags},
	{"SceneKit", "SceneKit", "scenekit", "SceneKit/SceneKit.h", []string{"SCN"}, DefaultFlags},
	{"SpriteKit", "SpriteKit", "spritekit", "SpriteKit/SpriteKit.h", []string{"SK"}, DefaultFlags},
	{"ModelIO", "Model I/O", "modelio", "ModelIO/ModelIO.h", []string{"MDL"}, DefaultFlags},
	{"IOSurface", "IOSurface", "iosurface", "IOSurface/IOSurface.h", []string{"IOSurface", "kIOSurface"}, DefaultFlags},
	{"Metal", "Metal", "metal", "Metal/Metal.h", []string{"MTL"}, DefaultFlags},
	{"MetalKit", "Metal Kit", "metalkit", "MetalKit/MetalKit.h", []string{"MTK"}, DefaultFlags},
	{"MetalPerformanceShadersGraph", "Metal Performance Shaders Graph", "mpsgraph", "MetalPerformanceShadersGraph/MetalPerformanceShadersGraph.h", []string{"MPSGraph"}, DefaultFlags},
	{"MetalPerformanceShaders", "Metal Performance Shaders", "mps", "MetalPerformanceShaders/MetalPerformanceShaders.h", []string{"MPS"}, DefaultFlags},
	{"MediaPlayer", "Media Player", "mediaplayer", "MediaPlayer/MediaPlayer.h", []string{"MP"}, DefaultFlags},
}
