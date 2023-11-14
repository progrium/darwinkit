package modules

import (
	"log"
	"strings"
)

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
			if strings.HasPrefix(symbolName, prefix) {
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
	} {
		if strings.EqualFold(m, mod) {
			return true
		}
	}
	return false
}

var All = []Module{
	{"objectivec", "Objective-C Runtime", "objc", "objc/runtime.h", []string{}},
	{"dispatch", "Dispatch", "dispatch", "dispatch/dispatch.h", []string{}},
	{"Kernel", "Kernel", "kernel", "Kernel/Kernel.h", []string{}},

	{"Foundation", "Foundation", "foundation", "Foundation/Foundation.h", []string{"NS"}},
	{"AppKit", "AppKit", "appkit", "Appkit/Appkit.h", []string{"NS"}},
	{"UIKit", "UIKit", "uikit", "UIKit/UIKit.h", []string{"NS"}},
	{"UniformTypeIdentifiers", "Uniform Type Identifiers", "uti", "UniformTypeIdentifiers/UniformTypeIdentifiers.h", []string{"UT"}},
	{"WebKit", "WebKit", "webkit", "WebKit/WebKit.h", []string{"WK"}},
	{"FileProvider", "File Provider", "fileprovider", "FileProvider/FileProvider.h", []string{"NS"}},
	{"Quartz", "Quartz", "quartz", "Quartz/Quartz.h", []string{"IK", "kQC", "kQuartz", "QC", "IK_"}},
	{"SecurityInterface", "Security Interface", "securityinterface", "SecurityInterface/SecurityInterface.h", []string{"SF"}},
	{"IOBluetooth", "IOBluetooth", "iobluetooth", "IOBluetooth/IOBluetooth.h", []string{"kIOBluetooth", "kBluetooth", "IOBluetooth", "Bluetooth", "kFTS", "kOBEX"}},
	{"CoreGraphics", "Core Graphics", "coregraphics", "CoreGraphics/CoreGraphics.h", []string{"CG", "kCG"}},
	{"CoreFoundation", "Core Foundation", "corefoundation", "CoreFoundation/CoreFoundation.h", []string{"CF", "kCF"}},
	{"QuartzCore", "Core Animation", "quartzcore", "QuartzCore/QuartzCore.h", []string{"kCA", "CA"}},
	{"Vision", "Vision", "vision", "Vision/Vision.h", []string{"VN"}},
	{"CoreSpotlight", "Core Spotlight", "corespotlight", "CoreSpotlight/CoreSpotlight.h", []string{"CS"}},
	{"CoreAudioKit", "Core Audio Kit", "coreaudiokit", "CoreAudioKit/CoreAudioKit.h", []string{"CA", "AU"}},
	{"CoreML", "Core ML", "coreml", "CoreML/CoreML.h", []string{"ML"}},
	{"CoreData", "Core Data", "coredata", "CoreData/CoreData.h", []string{"NS"}},
	{"CoreMediaIO", "Core Media I/O", "coremediaio", "CoreMediaIO/CMIOHardware.h", []string{"CMIO"}},
	{"CoreMedia", "Core Media", "coremedia", "CoreMedia/CoreMedia.h", []string{"CM"}},
	{"CoreImage", "Core Image", "coreimage", "CoreImage/CoreImage.h", []string{"CI"}},
	{"CoreMIDI", "Core MIDI", "coremidi", "CoreMIDI/CoreMIDI.h", []string{"MIDI", "kMIDI"}},
	{"CoreVideo", "Core Video", "corevideo", "CoreVideo/CoreVideo.h", []string{"CV", "kCV"}},
	{"CloudKit", "Cloud Kit", "cloudkit", "CloudKit/CloudKit.h", []string{"CK"}},
	{"AudioToolbox", "Audio Toolbox", "audiotoolbox", "AudioToolbox/AudioToolbox.h", []string{"AU"}},
	{"CoreAudio", "Core Audio", "coreaudio", "CoreAudio/CoreAudio.h", []string{"Audio", "kAudio"}},
	{"CoreAudioTypes", "Core Audio Types", "coreaudiotypes", "CoreAudio/CoreAudioTypes.h", []string{"AV"}},
	{"CoreLocation", "Core Location", "corelocation", "CoreLocation/CoreLocation.h", []string{"CL"}},
	{"Contacts", "Contacts", "contacts", "Contacts/Contacts.h", []string{"CN"}},
	{"ContactsUI", "Contacts UI", "contactsui", "ContactsUI/ContactsUI.h", []string{"CN"}},
	{"ImageIO", "Image I/O", "imageio", "ImageIO/ImageIO.h", []string{"CG", "kCG", "kCF"}},
	{"AVFAudio", "AVFAudio", "avfaudio", "AVFAudio/AVFAudio.h", []string{"AVAudio"}},
	{"AVFoundation", "AVFoundation", "avfoundation", "AVFoundation/AVFoundation.h", []string{"AV"}},
	{"AVKit", "AVKit", "avkit", "AVKit/AVKit.h", []string{"AV"}},
	{"GameplayKit", "GameplayKit", "gameplaykit", "GameplayKit/GameplayKit.h", []string{"GK"}},
	{"SystemConfiguration", "System Configuration", "sysconfig", "SystemConfiguration/SystemConfiguration.h", []string{"SC", "kSC"}},
	{"SceneKit", "SceneKit", "scenekit", "SceneKit/SceneKit.h", []string{"SCN"}},
	{"SpriteKit", "SpriteKit", "spritekit", "SpriteKit/SpriteKit.h", []string{"SK"}},
	{"ModelIO", "Model I/O", "modelio", "ModelIO/ModelIO.h", []string{"MDL"}},
	{"IOSurface", "IOSurface", "iosurface", "IOSurface/IOSurface.h", []string{"IOSurface", "kIOSurface"}},
	{"Metal", "Metal", "metal", "Metal/Metal.h", []string{"MTL"}},
	{"MetalKit", "Metal Kit", "metalkit", "MetalKit/MetalKit.h", []string{"MTK"}},
	{"MetalPerformanceShadersGraph", "Metal Performance Shaders Graph", "mpsgraph", "MetalPerformanceShadersGraph/MetalPerformanceShadersGraph.h", []string{"MPSGraph"}},
	{"MetalPerformanceShaders", "Metal Performance Shaders", "mps", "MetalPerformanceShaders/MetalPerformanceShaders.h", []string{"MPS"}},
	{"MediaPlayer", "Media Player", "mediaplayer", "MediaPlayer/MediaPlayer.h", []string{"MP"}},
}
