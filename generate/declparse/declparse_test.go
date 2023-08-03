package declparse

var tests = []struct {
	s         string
	n         Node
	ParseOnly bool
	Hint      Hint
}{

	{
		s: `@interface NSMenu : NSObject`,
		n: &InterfaceDecl{
			Name:      "NSMenu",
			SuperName: "NSObject",
		},
	},

	{
		s: `+ (BOOL)menuBarVisible`,
		n: &MethodDecl{
			TypeMethod: true,
			NameParts:  []string{"menuBarVisible"},
			ReturnType: TypeInfo{
				Name: "BOOL",
			},
		},
	},

	{
		s: `+ (instancetype)arrayWithObjects:(ObjectType)firstObj, ...;`,
		n: &Statement{
			Method: &MethodDecl{
				TypeMethod: true,
				NameParts:  []string{"arrayWithObjects"},
				ReturnType: TypeInfo{
					Name: "instancetype",
				},
				Args: []ArgInfo{
					{
						Name: "firstObj",
						Type: TypeInfo{
							Name: "ObjectType",
						},
					},
				},
				Variadic: true,
			},
		},
	},

	{
		s: `+ (void)doSomething:(NSArray * _Nullable *)withData error:(NSError **)error;`,
		n: &Statement{
			Method: &MethodDecl{
				TypeMethod: true,
				ReturnType: TypeInfo{
					Name: "void",
				},
				NameParts: []string{"doSomething", "error"},
				Args: []ArgInfo{
					{
						Name: "withData",
						Type: TypeInfo{
							Name:     "NSArray",
							IsPtr:    true,
							IsPtrPtr: true,
							Annots: map[TypeAnnotation]bool{
								TypeAnnotNullable: true,
							},
						},
					},
					{
						Name: "error",
						Type: TypeInfo{
							Name:     "NSError",
							IsPtr:    true,
							IsPtrPtr: true,
						},
					},
				},
			},
		},
	},

	{
		s: `+ (NSArray<NSView *> * _Nullable)interestingObjectsForKey:(NSString * _Nonnull)key;`,
		n: &Statement{
			Method: &MethodDecl{
				TypeMethod: true,
				ReturnType: TypeInfo{
					Name: "NSArray",
					Params: []TypeInfo{
						{
							Name:  "NSView",
							IsPtr: true,
						},
					},
					IsPtr: true,
					Annots: map[TypeAnnotation]bool{
						TypeAnnotNullable: true,
					},
				},
				NameParts: []string{"interestingObjectsForKey"},
				Args: []ArgInfo{
					{
						Name: "key",
						Type: TypeInfo{
							Name:  "NSString",
							IsPtr: true,
							Annots: map[TypeAnnotation]bool{
								TypeAnnotNonnull: true,
							},
						},
					},
				},
			},
		},
	},

	{
		s: `+ (void)setMenuBarVisible:(BOOL)visible;`,
		n: &Statement{
			Method: &MethodDecl{
				TypeMethod: true,
				ReturnType: TypeInfo{
					Name: "void",
				},
				NameParts: []string{"setMenuBarVisible"},
				Args: []ArgInfo{
					{
						Name: "visible",
						Type: TypeInfo{
							Name: "BOOL",
						},
					},
				},
			},
		},
	},

	{
		s: `- (void)numberWithLongLong:(unsigned long long)value atIndex:(unsigned int)idx keyCode:(unsigned short)code;`,
		n: &Statement{
			Method: &MethodDecl{
				ReturnType: TypeInfo{
					Name: "void",
				},
				NameParts: []string{"numberWithLongLong", "atIndex", "keyCode"},
				Args: []ArgInfo{
					{
						Name: "value",
						Type: TypeInfo{
							Name: "long long",
							Annots: map[TypeAnnotation]bool{
								TypeAnnotUnsigned: true,
							},
						},
					},
					{
						Name: "idx",
						Type: TypeInfo{
							Name: "int",
							Annots: map[TypeAnnotation]bool{
								TypeAnnotUnsigned: true,
							},
						},
					},
					{
						Name: "code",
						Type: TypeInfo{
							Name: "short",
							Annots: map[TypeAnnotation]bool{
								TypeAnnotUnsigned: true,
							},
						},
					},
				},
			},
		},
	},

	{
		s: `- (void)beginSheet:(NSWindow *)sheetWindow completionHandler:(void (^)(NSModalResponse returnCode))handler;`,
		n: &Statement{
			Method: &MethodDecl{
				ReturnType: TypeInfo{
					Name: "void",
				},
				NameParts: []string{"beginSheet", "completionHandler"},
				Args: []ArgInfo{
					{
						Name: "sheetWindow",
						Type: TypeInfo{
							Name:  "NSWindow",
							IsPtr: true,
						},
					},
					{
						Name: "handler",
						Type: TypeInfo{
							Func: &FunctionDecl{
								IsBlock: true,
								ReturnType: TypeInfo{
									Name: "void",
								},
								Args: FuncArgs{
									{
										Name: "returnCode",
										Type: TypeInfo{
											Name: "NSModalResponse",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},

	{
		s: `- (void)sortSubviewsUsingFunction:(NSComparisonResult (*)(__kindof NSView *, __kindof NSView *, void *))compare;`,
		n: &Statement{
			Method: &MethodDecl{
				ReturnType: TypeInfo{
					Name: "void",
				},
				NameParts: []string{"sortSubviewsUsingFunction"},
				Args: []ArgInfo{
					{
						Name: "compare",
						Type: TypeInfo{
							Func: &FunctionDecl{
								IsPtr: true,
								ReturnType: TypeInfo{
									Name: "NSComparisonResult",
								},
								Args: FuncArgs{
									{
										Type: TypeInfo{
											Name:  "NSView",
											IsPtr: true,
											Annots: map[TypeAnnotation]bool{
												TypeAnnotKindOf: true,
											},
										},
									},
									{
										Type: TypeInfo{
											Name:  "NSView",
											IsPtr: true,
											Annots: map[TypeAnnotation]bool{
												TypeAnnotKindOf: true,
											},
										},
									},
									{
										Type: TypeInfo{
											Name:  "void",
											IsPtr: true,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},

	{
		s: `- (oneway void)releaseGlobally;`,
		n: &Statement{
			Method: &MethodDecl{
				ReturnType: TypeInfo{
					Name: "void",
					Annots: map[TypeAnnotation]bool{
						TypeAnnotOneway: true,
					},
				},
				NameParts: []string{"releaseGlobally"},
			},
		},
	},

	{
		s: `- (NSRect)convertRectToBacking:(NSRect)rect;`,
		n: &Statement{
			Method: &MethodDecl{
				ReturnType: TypeInfo{
					Name: "NSRect",
				},
				NameParts: []string{"convertRectToBacking"},
				Args: []ArgInfo{
					{
						Name: "rect",
						Type: TypeInfo{
							Name: "NSRect",
						},
					},
				},
			},
		},
	},

	{
		s: `- (BOOL)instantiateNibWithOwner:(id)owner topLevelObjects:(NSArray * _Null_unspecified *)topLevelObjects;`,
		n: &Statement{
			Method: &MethodDecl{
				ReturnType: TypeInfo{
					Name: "BOOL",
				},
				NameParts: []string{"instantiateNibWithOwner", "topLevelObjects"},
				Args: []ArgInfo{
					{
						Name: "owner",
						Type: TypeInfo{
							Name: "id",
						},
					},
					{
						Name: "topLevelObjects",
						Type: TypeInfo{
							Name:     "NSArray",
							IsPtr:    true,
							IsPtrPtr: true,
							Annots: map[TypeAnnotation]bool{
								TypeAnnotNullUnspecified: true,
							},
						},
					},
				},
			},
		},
	},

	{
		s: `- (NSColor *)blendedColorWithFraction:(CGFloat)fraction ofColor:(NSColor *)color;`,
		n: &Statement{
			Method: &MethodDecl{
				ReturnType: TypeInfo{
					Name:  "NSColor",
					IsPtr: true,
				},
				NameParts: []string{"blendedColorWithFraction", "ofColor"},
				Args: []ArgInfo{
					{
						Name: "fraction",
						Type: TypeInfo{
							Name: "CGFloat",
						},
					},
					{
						Name: "color",
						Type: TypeInfo{
							Name:  "NSColor",
							IsPtr: true,
						},
					},
				},
			},
		},
	},

	{
		s: `- (void)resetCursorRects;`,
		n: &Statement{
			Method: &MethodDecl{
				ReturnType: TypeInfo{
					Name: "void",
				},
				NameParts: []string{"resetCursorRects"},
			},
		},
	},

	{
		s: `- (instancetype)initWithContentRect:(NSRect)contentRect styleMask:(NSWindowStyleMask)style backing:(NSBackingStoreType)backingStoreType defer:(BOOL)flag;`,
		n: &Statement{
			Method: &MethodDecl{
				ReturnType: TypeInfo{
					Name: "instancetype",
				},
				NameParts: []string{"initWithContentRect", "styleMask", "backing", "defer"},
				Args: []ArgInfo{
					{
						Name: "contentRect",
						Type: TypeInfo{
							Name: "NSRect",
						},
					}, {
						Name: "style",
						Type: TypeInfo{
							Name: "NSWindowStyleMask",
						},
					}, {
						Name: "backingStoreType",
						Type: TypeInfo{
							Name: "NSBackingStoreType",
						},
					}, {
						Name: "flag",
						Type: TypeInfo{
							Name: "BOOL",
						},
					},
				},
			},
		},
	},

	{
		s: `- (void)removeStatusItem:(NSStatusItem *)item;`,
		n: &Statement{
			Method: &MethodDecl{
				NameParts: []string{"removeStatusItem"},
				ReturnType: TypeInfo{
					Name: "void",
				},
				Args: []ArgInfo{
					{
						Name: "item",
						Type: TypeInfo{
							Name:  "NSStatusItem",
							IsPtr: true,
						},
					},
				},
			},
		},
	},

	{
		s: `@property CGFloat alphaValue;`,
		n: &Statement{
			Property: &PropertyDecl{
				Name: "alphaValue",
				Type: TypeInfo{
					Name: "CGFloat",
				},
			},
		},
	},

	{
		s: `@property(readonly, getter=isVisible) BOOL visible;`,
		n: &Statement{
			Property: &PropertyDecl{
				Name: "visible",
				Attrs: map[PropAttr]string{
					PropAttrGetter:   "isVisible",
					PropAttrReadonly: "",
				},
				Type: TypeInfo{
					Name: "BOOL",
				},
			},
		},
	},

	{
		s: `@property(class, strong, readonly) NSStatusBar *systemStatusBar;`,
		n: &Statement{
			Property: &PropertyDecl{
				Name: "systemStatusBar",
				Attrs: map[PropAttr]string{
					PropAttrReadonly: "",
					PropAttrClass:    "",
					PropAttrStrong:   "",
				},
				Type: TypeInfo{
					Name:  "NSStatusBar",
					IsPtr: true,
				},
			},
		},
	},

	{
		s: `@property(readonly) const void *eventRef;`,
		n: &Statement{
			Property: &PropertyDecl{
				Name: "eventRef",
				Attrs: map[PropAttr]string{
					PropAttrReadonly: "",
				},
				Type: TypeInfo{
					Name:  "void",
					IsPtr: true,
					Annots: map[TypeAnnotation]bool{
						TypeAnnotConst: true,
					},
				},
			},
		},
	},

	{
		s: `@property(weak) id<NSWindowDelegate> delegate;`,
		n: &Statement{
			Property: &PropertyDecl{
				Name: "delegate",
				Attrs: map[PropAttr]string{
					PropAttrWeak: "",
				},
				Type: TypeInfo{
					Name: "id",
					Params: []TypeInfo{
						{
							Name: "NSWindowDelegate",
						},
					},
				},
			},
		},
	},

	{
		s: `@property(readonly, copy) NSArray<__kindof NSWindow *> *sheets;`,
		n: &Statement{
			Property: &PropertyDecl{
				Name: "sheets",
				Attrs: map[PropAttr]string{
					PropAttrCopy:     "",
					PropAttrReadonly: "",
				},
				Type: TypeInfo{
					Name:  "NSArray",
					IsPtr: true,
					Params: []TypeInfo{
						{
							Name:  "NSWindow",
							IsPtr: true,
							Annots: map[TypeAnnotation]bool{
								TypeAnnotKindOf: true,
							},
						},
					},
				},
			},
		},
	},

	{
		s: `@property(weak) __kindof NSWindowController *windowController;`,
		n: &Statement{
			Property: &PropertyDecl{
				Name: "windowController",
				Attrs: map[PropAttr]string{
					PropAttrWeak: "",
				},
				Type: TypeInfo{
					Name:  "NSWindowController",
					IsPtr: true,
					Annots: map[TypeAnnotation]bool{
						TypeAnnotKindOf: true,
					},
				},
			},
		},
	},

	{
		s: `@property(strong, readonly, nullable) NSScreen *screen;`,
		n: &Statement{
			Property: &PropertyDecl{
				Name: "screen",
				Attrs: map[PropAttr]string{
					PropAttrStrong:   "",
					PropAttrReadonly: "",
					PropAttrNullable: "",
				},
				Type: TypeInfo{
					Name:  "NSScreen",
					IsPtr: true,
				},
			},
		},
	},

	{
		s: `@property(strong) IBOutlet NSView *view;`,
		n: &Statement{
			Property: &PropertyDecl{
				Name: "view",
				Attrs: map[PropAttr]string{
					PropAttrStrong: "",
				},
				Type: TypeInfo{
					Name:  "NSView",
					IsPtr: true,
				},
				IsOutlet: true,
			},
		},
	},

	{
		s: `- (BOOL)writeObjects:(NSArray<id<NSPasteboardWriting>> *)objects;`,
		n: &Statement{
			Method: &MethodDecl{
				ReturnType: TypeInfo{
					Name: "BOOL",
				},
				NameParts: []string{"writeObjects"},
				Args: []ArgInfo{
					{
						Name: "objects",
						Type: TypeInfo{
							Name:  "NSArray",
							IsPtr: true,
							Params: []TypeInfo{
								{
									Name: "id",
									Params: []TypeInfo{
										{
											Name: "NSPasteboardWriting",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},

	{
		s: `+ (id)addLocalMonitorForEventsMatchingMask:(NSEventMask)mask handler:(NSEvent * _Nullable (^)(NSEvent * event))block;`,
		n: &Statement{
			Method: &MethodDecl{
				TypeMethod: true,
				ReturnType: TypeInfo{
					Name: "id",
				},
				NameParts: []string{"addLocalMonitorForEventsMatchingMask", "handler"},
				Args: []ArgInfo{
					{
						Name: "mask",
						Type: TypeInfo{
							Name: "NSEventMask",
						},
					},
					{
						Name: "block",
						Type: TypeInfo{
							Func: &FunctionDecl{
								IsBlock: true,
								ReturnType: TypeInfo{
									Name:  "NSEvent",
									IsPtr: true,
									Annots: map[TypeAnnotation]bool{
										TypeAnnotNullable: true,
									},
								},
								Args: FuncArgs{
									{
										Name: "event",
										Type: TypeInfo{
											Name:  "NSEvent",
											IsPtr: true,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},

	{
		s: `enum {
			NSScaleProportionally = 0,
			NSScaleToFit,
			NSScaleNone
		 };`,
		n: &Statement{
			Enum: &EnumDecl{
				Cases: []VariableDecl{
					{
						Name:  "NSScaleProportionally",
						Value: "0",
					},
					{
						Name: "NSScaleToFit",
					},
					{
						Name: "NSScaleNone",
					},
				},
			},
		},
	},

	{
		s: `const NSURLResourceKey NSURLVolumeURLForRemountingKey;`,
		n: &Statement{
			Variable: &VariableDecl{
				Name: "NSURLVolumeURLForRemountingKey",
				Type: TypeInfo{
					Name: "NSURLResourceKey",
					Annots: map[TypeAnnotation]bool{
						TypeAnnotConst: true,
					},
				},
			},
		},
	},

	{
		ParseOnly: true,
		Hint:      HintVariable,
		s:         `NSString *const NSAssertionHandlerKey;`,
		n: &Statement{
			Variable: &VariableDecl{
				Name: "NSAssertionHandlerKey",
				Type: TypeInfo{
					Name:  "NSString",
					IsPtr: true,
					Annots: map[TypeAnnotation]bool{
						TypeAnnotConst: true,
					},
				},
			},
		},
	},

	{
		ParseOnly: true,
		Hint:      HintEnumCase,
		s:         `kCALayerLeftEdge = 1U << 0`,
		n: &Statement{
			Variable: &VariableDecl{
				Name:  "kCALayerLeftEdge",
				Value: "1U<<0",
			},
		},
	},

	{
		ParseOnly: true,
		s: `typedef enum WKUserScriptInjectionTime : NSInteger {
			...
		} WKUserScriptInjectionTime;`,
		n: &Statement{
			Enum: &EnumDecl{
				Name: "WKUserScriptInjectionTime",
				Type: TypeInfo{
					Name: "NSInteger",
				},
			},
			Typedef: "WKUserScriptInjectionTime",
		},
	},

	{
		ParseOnly: true,
		s:         `typedef NSString *NSDeviceDescriptionKey;`,
		n: &Statement{
			TypeAlias: &TypeInfo{
				Name:  "NSString",
				IsPtr: true,
			},
			Typedef: "NSDeviceDescriptionKey",
		},
	},

	{
		ParseOnly: true,
		Hint:      HintVariable,
		s:         `CGPoint origin;`,
		n: &Statement{
			Variable: &VariableDecl{
				Name: "origin",
				Type: TypeInfo{
					Name: "CGPoint",
				},
			},
		},
	},
	{
		s:    `CFTypeID CGEventGetTypeID(void);`,
		Hint: HintFunction,
		n: &Statement{
			Function: &FunctionDecl{
				ReturnType: TypeInfo{
					Name: "CFTypeID",
				},
				Name: "CGEventGetTypeID",
			},
		},
	},
	{
		s:    `CGEventRef CGEventCreateScrollWheelEvent(CGEventSourceRef source, CGScrollEventUnit units, uint32_t wheelCount, int32_t wheel1, ...);`,
		Hint: HintFunction,
		n: &Statement{
			Function: &FunctionDecl{
				ReturnType: TypeInfo{
					Name: "CGEventRef",
				},
				Name: "CGEventCreateScrollWheelEvent",
				Args: FuncArgs{
					{
						Name: "source",
						Type: TypeInfo{
							Name: "CGEventSourceRef",
						},
					},
					{
						Name: "units",
						Type: TypeInfo{
							Name: "CGScrollEventUnit",
						},
					},
					{
						Name: "wheelCount",
						Type: TypeInfo{
							Name: "uint32_t",
						},
					},
					{
						Name: "wheel1",
						Type: TypeInfo{
							Name: "int32_t",
						},
					},
				},
				Variadic: true,
			},
		},
	},
	{
		s: `+ (instancetype)allocWithZone:(struct _NSZone *)zone;`,
		n: &Statement{
			Method: &MethodDecl{
				TypeMethod: true,
				ReturnType: TypeInfo{
					Name: "instancetype",
				},
				NameParts: []string{"allocWithZone"},
				Args: FuncArgs{
					{
						Name: "zone",
						Type: TypeInfo{
							Name:  "_NSZone",
							IsPtr: true,
							Annots: map[TypeAnnotation]bool{
								TypeAnnotStruct: true,
							},
						},
					},
				},
			},
		},
	},
	{
		s: `- (instancetype)init NS_UNAVAILABLE;`,
		n: &Statement{
			Method: &MethodDecl{
				ReturnType: TypeInfo{
					Name: "instancetype",
				},
				NameParts:   []string{"init"},
				Unavailable: true,
			},
		},
	},
}
