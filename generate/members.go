package generate

import (
	"fmt"
	"log"
	"strings"

	"github.com/progrium/macdriver/generate/codegen"
	"github.com/progrium/macdriver/generate/declparse"
	"github.com/progrium/macdriver/generate/modules"
)

func (db *Generator) shouldSkipType(ti declparse.TypeInfo) bool {
	for _, name := range ti.TypeNames() {
		ts := db.FindTypeSymbol(name)
		if ts != nil {
			if m := ts.MainModule(); m != nil && modules.CanSkipModuleCoupling(db.Framework, m.Package) {
				log.Printf("skipping use of type '%s' to decouple framework '%s'\n", name, m.Package)
				return true
			}
		}
	}
	for _, n := range []string{
		"EAGLContext", // no OpenGL ES
		"GLKVector2",
		"GLKVector3",
		"GLKVector4",
		"GLKMatrix2",
		"GLKMatrix3",
		"GLKMatrix4",
		"AudioStreamBasicDescription",         // core audio types
		"MPNowPlayingInfoLanguageOptionGroup", // media player
		"MPNowPlayingInfoLanguageOption",      // media player
		"SecTrustRef",                         // security
		"SecIdentityRef",
		"JSContext", // javascriptcore
		"ICScannerDevice",
		"ICCameraDevice",
		"ICDevice",
		"NSDiffableDataSourceSnapshot", // uikit
		"NSLineBreakMode",
		"UNNotificationAction",
		"AppleEvent",
		"AEEventClass",
		"AEEventID",
		"AEKeyword",
		"AEReturnID",
		"AETransactionID",
		"AEDesc",
		"DescType",
		"xpc_type_t",
		"xpc_object_t",
		"MKMapItem",
		"INInteraction",
		"MKCoordinateSpan",
		"tls_protocol_version_t",
		"IOReturn",
		"ByteCount",
		"gid_t",
		"pid_t",
		"uid_t",
		"au_asid_t",
		"va_list",
		"CLBeaconIdentityConstraint",
		"mach_port_t",
		"cpu_type_t",
	} {
		if ti.Name == n {
			return true
		}
		for _, p := range ti.Params {
			if p.Name == n {
				return true
			}
		}
	}
	return false
}

func (db *Generator) Members(fw string, sym Symbol, covariantTypes []string) (properties []*codegen.Property, methods []*codegen.Method) {
	var selectorToSuffix = map[string]string{}
	var selectorSeen = map[string]string{}

	for _, s := range db.ModuleSymbols(fw) {
		if !s.HasPlatform(db.Platform, db.Version, false) {
			continue
		}
		if s.Parent == sym.Name {
			// replace covariant type names with NSObject
			// for now until we come up with something better
			for _, ct := range covariantTypes {
				s.Declaration = strings.ReplaceAll(s.Declaration, fmt.Sprintf(" %s", ct), " NSObject")
				s.Declaration = strings.ReplaceAll(s.Declaration, fmt.Sprintf("<%s>", ct), "<NSObject>")
			}

			st, err := s.Parse()
			if err != nil {
				log.Println("Members:", sym.Name, "::", s.Declaration)
				panic(err)
			}
			switch s.Kind {
			case "Property":
				if st.Property == nil {
					// per platform declarations arent supported
					// in symbols.zip but we might be able to infer
					// when we get a method declaration instead
					if st.Method != nil {
						st.Property = &declparse.PropertyDecl{
							Name: st.Method.Name(),
							Type: st.Method.ReturnType,
						}
					} else {
						log.Printf("bad declaration for prop: %s.%s\n", sym.Name, s.Name)
						continue
					}
				}
				// added for AVAudioUnitComponent.icon
				if st.Property.Type.Name == "UIImage" {
					st.Property.Type.Name = "NSImage"
				}
				// added for SKEmiterNode.particleColor
				if st.Property.Type.Name == "UIColor" {
					st.Property.Type.Name = "NSColor"
				}
				if st.Property.Type.Name == "UIBezierPath" {
					st.Property.Type.Name = "NSBezierPath"
				}
				if st.Property.Type.Name == "UIFont" {
					st.Property.Type.Name = "NSFont"
				}
				// fmt.Println("prop:", sym.Name, s.Name)
				if db.shouldSkipType(st.Property.Type) {
					log.Printf("skipping prop for blacklist type '%s': owner=%s decl=%s\n", st.Property.Type, sym.Name, s.Declaration)
					continue
				}
				typ := db.ParseType(st.Property.Type)
				if typ == nil {
					log.Fatalf("Property type failure: owner=%s propdecl=%s", sym.Name, s.Declaration)
				}
				properties = append(properties, &codegen.Property{
					Name:          st.Property.Name,
					GoName:        st.Property.Name,
					Type:          typ,
					GetterName:    st.Property.Attrs[declparse.PropAttrGetter],
					SetterName:    st.Property.Attrs[declparse.PropAttrSetter],
					ReadOnly:      st.Property.HasAttr(declparse.PropAttrReadonly),
					ClassProperty: st.Property.HasAttr(declparse.PropAttrClass),
					Weak:          st.Property.HasAttr(declparse.PropAttrWeak),
					Description:   s.Description,
					DocURL:        s.DocURL(),
					//Deprecated:    TODO,
					//Required:    ??,
				})
			case "Method":
				if s.Deprecated || !s.HasPlatform(db.Platform, db.Version, false) {
					continue
				}
				if st.Method == nil {
					log.Printf("bad declaration for method: %s.%s :: %s\n", sym.Name, s.Name, s.Declaration)
					continue
				}

				var params []*codegen.Param
				var skipType *declparse.TypeInfo
				for idx, arg := range st.Method.Args {
					// added for SKAction.colorizeWithColor:colorBlendFactor:duration:
					if arg.Type.Name == "UIColor" {
						arg.Type.Name = "NSColor"
					}
					if arg.Type.Name == "UIImage" {
						arg.Type.Name = "NSImage"
					}
					if arg.Type.Name == "UIBezierPath" {
						arg.Type.Name = "NSBezierPath"
					}
					if db.shouldSkipType(arg.Type) {
						skipType = &arg.Type
						break
					}
					ptyp := db.ParseType(arg.Type)
					if ptyp == nil {
						log.Fatalf("Method param type failure: owner=%s arg=%s type=%s methoddecl=%s", sym.Name, arg.Name, arg.Type.Name, s.Declaration)
					}
					param := &codegen.Param{
						Name:      arg.Name,
						Type:      ptyp,
						FieldName: st.Method.NameParts[idx],
					}
					params = append(params, param)
				}
				if skipType != nil {
					log.Printf("skipping method for blacklist type '%s': owner=%s decl=%s\n", skipType.Name, sym.Name, s.Declaration)
					continue
				}

				if st.Method.ReturnType.Name == "UIImage" {
					st.Method.ReturnType.Name = "NSImage"
				}
				if db.shouldSkipType(st.Method.ReturnType) {
					log.Printf("skipping method for blacklist type '%s': owner=%s decl=%s\n", st.Method.ReturnType.Name, sym.Name, s.Declaration)
					continue
				}
				rettyp := db.ParseType(st.Method.ReturnType)
				if rettyp == nil {
					log.Fatalf("Method return type failure: owner=%s decl=%s", sym.Name, s.Declaration)
				}

				gm := &codegen.Method{
					Name:        st.Method.NameParts[0],
					GoName:      st.Method.NameParts[0],
					Params:      params,
					ReturnType:  rettyp,
					ClassMethod: st.Method.TypeMethod,
					Description: s.Description,
					DocURL:      s.DocURL(),
					Deprecated:  false, // todo: support deprecated
					//Required:    ??,
				}

				// skip if defined in custom
				qualifiedName := fmt.Sprintf("%s#%s", modules.TrimPrefix(sym.Name), gm.GoFuncName())
				if db.customMethods.Contains(qualifiedName) {
					log.Printf("skipping custom defined method '%s'\n", qualifiedName)
					continue
				}

				// handle name conflicts
				sel := gm.Selector()
				goSel := gm.ProtocolGoFuncName()
				if selectorSeen[goSel] != "" {
					// if conflict, mark the longer selector to add suffix
					if len(selectorSeen[goSel]) > len(sel) {
						selectorToSuffix[goSel] = selectorSeen[goSel]
					} else {
						selectorToSuffix[goSel] = sel
					}
				} else {
					selectorSeen[goSel] = sel
				}

				methods = append(methods, gm)
			}
		}
	}

	// mark suffix
	for idx := range methods {
		if selectorToSuffix[methods[idx].ProtocolGoFuncName()] == methods[idx].Selector() {
			methods[idx].Suffix = true
		}
	}

	return
}
