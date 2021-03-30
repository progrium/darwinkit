package declparser

import (
	"bytes"
	"reflect"
	"testing"
)


func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *Statement
		wantErr bool
	}{
		{
			name:    "empty input errors",
			input:   "",
			want:    nil,
			wantErr: true,
		},
		{
			name: "interface with super class",
			input: `@interface NSMenu : NSObject`,
			want: &Statement{
				Interface: &InterfaceDecl{
					Name:      "NSMenu",
					SuperName: "NSObject",
				},
			},
		},
		{
			name: "type method",
			input: `+ (BOOL)menuBarVisible;`,
			want: &Statement{
				Method: &MethodDecl{
					TypeMethod: true,
					NameParts: []string{"menuBarVisible"},
					ReturnType: TypeInfo{
						Name: "BOOL",
					},
				},
			},
		},
		{
			name: "type method with one arg",
			input: `+ (void)setMenuBarVisible:(BOOL)visible;`,
			want: &Statement{
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
			name: "instance method",
			input: `- (NSRect)convertRectToBacking:(NSRect)rect;`,
			want: &Statement{
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
			name: "instance method with pointer return and arg",
			input: `- (NSColor *)blendedColorWithFraction:(CGFloat)fraction 
                              ofColor:(NSColor *)color;`,
			want: &Statement{
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
			name: "instance method with no arg and void return",
			input: `- (void)resetCursorRects;`,
			want: &Statement{
				Method: &MethodDecl{
					ReturnType: TypeInfo{
						Name: "void",
					},
					NameParts: []string{"resetCursorRects"},
				},
			},
		},
		{
			name: "instance method with more than one arg",
			input: `- (instancetype)initWithContentRect:(NSRect)contentRect 
                          styleMask:(NSWindowStyleMask)style 
                            backing:(NSBackingStoreType)backingStoreType 
                              defer:(BOOL)flag;`,
			want: &Statement{
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
			name: "instance method  with pointer arg and void return",
			input: `- (void)removeStatusItem:(NSStatusItem *)item;`,
			want: &Statement{
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
			name: "property",
			input: `@property CGFloat alphaValue;`,
			want: &Statement{
				Property: &PropertyDecl{
					Name: "alphaValue",
					Type: TypeInfo{
						Name: "CGFloat",
					},
				},
			},
		},
		{
			name: "property with getter and read only",
			input: `@property(getter=isVisible, readonly) BOOL visible;`,
			want: &Statement{
				Property: &PropertyDecl{
					Name:     "visible",
					Readonly: true,
					Getter:   "isVisible",
					Type: TypeInfo{
						Name: "BOOL",
					},
				},
			},
		},
		{
			name: "property on class with strong, read only and pointer type",
			input: `@property(class, readonly, strong) NSStatusBar *systemStatusBar;`,
			want: &Statement{
				Property: &PropertyDecl{
					Name:     "systemStatusBar",
					TypeProperty: true,
					Readonly: true,
					Type: TypeInfo{
						Name: "NSStatusBar",
						IsPtr: true,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBufferString(tt.input)
			p := NewParser(buf)
			got, err := p.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %s, stmt %s", got, tt.want)
			}
		})
	}
}

func TestStatement_String(t *testing.T) {
	 tests := []struct {
		name  string
		input *Statement
		want  string
	}{
		 {
			 name: "interface with super class",
			 want: `@interface NSMenu : NSObject`,
			 input: &Statement{
				 Interface: &InterfaceDecl{
					 Name:      "NSMenu",
					 SuperName: "NSObject",
				 },
			 },
		 },
		 {
			 name: "type method",
			 want: `+ (BOOL)menuBarVisible;`,
			 input: &Statement{
				 Method: &MethodDecl{
					 TypeMethod: true,
					 NameParts: []string{"menuBarVisible"},
					 ReturnType: TypeInfo{
						 Name: "BOOL",
					 },
				 },
			 },
		 },
		 {
			 name: "type method with one arg",
			 want: `+ (void)setMenuBarVisible:(BOOL)visible;`,
			 input: &Statement{
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
			 name: "instance method",
			 want: `- (NSRect)convertRectToBacking:(NSRect)rect;`,
			 input: &Statement{
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
			 name: "instance method with pointer return and arg",
			 want: `- (NSColor *)blendedColorWithFraction:(CGFloat)fraction 
ofColor:(NSColor *)color;`,
			 input: &Statement{
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
			 name: "instance method with no arg and void return",
			 want: `- (void)resetCursorRects;`,
			 input: &Statement{
				 Method: &MethodDecl{
					 ReturnType: TypeInfo{
						 Name: "void",
					 },
					 NameParts: []string{"resetCursorRects"},
				 },
			 },
		 },
		 {
			 name: "instance method with more than one arg",
			 want: `- (instancetype)initWithContentRect:(NSRect)contentRect 
styleMask:(NSWindowStyleMask)style 
backing:(NSBackingStoreType)backingStoreType 
defer:(BOOL)flag;`,
			 input: &Statement{
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
			 name: "instance method  with pointer arg and void return",
			 want: `- (void)removeStatusItem:(NSStatusItem *)item;`,
			 input: &Statement{
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
			 name: "property",
			 want: `@property CGFloat alphaValue;`,
			 input: &Statement{
				 Property: &PropertyDecl{
					 Name: "alphaValue",
					 Type: TypeInfo{
						 Name: "CGFloat",
					 },
				 },
			 },
		 },
		 {
			 name: "property with getter and read only",
			 want: `@property(getter=isVisible, readonly) BOOL visible;`,
			 input: &Statement{
				 Property: &PropertyDecl{
					 Name:     "visible",
					 Readonly: true,
					 Getter:   "isVisible",
					 Type: TypeInfo{
						 Name: "BOOL",
					 },
				 },
			 },
		 },
		 {
			 name: "property on class with strong, read only and pointer type",
			 want: `@property(class, readonly) NSStatusBar *systemStatusBar;`,
			 input: &Statement{
				 Property: &PropertyDecl{
					 Name:     "systemStatusBar",
					 TypeProperty: true,
					 Readonly: true,
					 Type: TypeInfo{
						 Name: "NSStatusBar",
						 IsPtr: true,
					 },
				 },
			 },
		 },
	 }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := (*tt.input).String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %q, want = %q", got, tt.want)
			}
		})
	}
}
