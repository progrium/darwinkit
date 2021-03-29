package declparser

import (
	"bytes"
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	validDecls := []string{
		`- (NSRect)convertRectToBacking:(NSRect)rect;`,
		`+ (BOOL)menuBarVisible;`,
		`+ (void)setMenuBarVisible:(BOOL)visible;`,
		`- (NSColor *)blendedColorWithFraction:(CGFloat)fraction ofColor:(NSColor *)color;`,
		`- (void)resetCursorRects;`,
		`- (instancetype)initWithContentRect:(NSRect)contentRect styleMask:(NSWindowStyleMask)style backing:(NSBackingStoreType)backingStoreType defer:(BOOL)flag;`,
		`- (void)removeStatusItem:(NSStatusItem *)item;`,
		`@property CGFloat alphaValue;`,
		`@property(getter=isVisible, readonly) BOOL visible;`,
		`@property(class, readonly, strong) NSStatusBar *systemStatusBar;`,
		`@interface NSMenu : NSObject`,
	}
	for _, m := range validDecls {
		buf := bytes.NewBufferString(m)
		p := NewParser(buf)
		stmt, err := p.Parse()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(stmt.Method, stmt.Property, stmt.Interface)
	}
}
