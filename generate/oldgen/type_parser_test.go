package oldgen

import (
	"testing"

	"github.com/progrium/macdriver/generate/typing"
	"github.com/progrium/macdriver/internal/assert"
)

func TestParseType(t *testing.T) {
	assert.Equal(t, "NSInteger", ParseType("NSInteger").ObjcName())
	assert.Equal(t, "BOOL", ParseType("BOOL").ObjcName())
	assert.Equal(t, "float", ParseType("float").ObjcName())

	assert.NotNil(t, ParseType("void").(*typing.VoidType))
	assert.NotNil(t, ParseType("void *").(*typing.VoidPointerType))
	assert.NotNil(t, ParseType("void*").(*typing.VoidPointerType))
	assert.NotNil(t, ParseType("SEL").(*typing.SelectorType))
	assert.NotNil(t, ParseType("id").(*typing.IDType))
	assert.NotNil(t, ParseType("instance").(*typing.InstanceType))

	//string
	assert.NotNil(t, ParseType("NSString *").(*typing.StringType))
	assert.NotNil(t, ParseType("NSString*").(*typing.StringType))
	assert.NotNil(t, ParseType("NSData *").(*typing.DataType))

	// class
	ct := ParseType("Foundation.NSBundle*").(*typing.ClassType)
	assert.Equal(t, "NSBundle", ct.Name)
	assert.Equal(t, "Foundation", ct.Module.Name)
	ct = ParseType("objectivec.NSObject *").(*typing.ClassType)
	assert.Equal(t, "NSObject", ct.Name)

	//protocol
	pt := ParseType("id<AppKit.NSDraggingInfo>").(*typing.ProtocolType)
	assert.Equal(t, "NSDraggingInfo", pt.Name)

	// struct
	st := ParseType("CoreGraphics.CGPoint").(*typing.StructType)
	assert.Equal(t, "CGPoint", st.Name)

	// alias
	at := ParseType("Foundation.NSPoint").(*typing.AliasType)
	assert.Equal(t, "NSPoint", at.Name)
	assert.Equal(t, "CGPoint", at.Type.ObjcName())

	// pointer
	pointerType := ParseType("BOOL *").(*typing.PointerType)
	assert.Equal(t, "BOOL", pointerType.Type.ObjcName())

	// array
	arrayType := ParseType("NSArray<AppKit.NSButton *> *").(*typing.ArrayType)
	assert.Equal(t, "NSButton*", arrayType.Type.ObjcName())
	arrayType = ParseType("NSArray*").(*typing.ArrayType)
	assert.Equal(t, "NSObject*", arrayType.Type.ObjcName())

	// dict
	dt := ParseType("NSDictionary<AppKit.NSAboutPanelOptionKey, objectivec.NSObject *> *").(*typing.DictType)
	assert.Equal(t, "NSAboutPanelOptionKey", dt.KeyType.ObjcName())
	assert.Equal(t, "NSObject*", dt.ValueType.ObjcName())
	dt = ParseType("NSDictionary*").(*typing.DictType)
	assert.Equal(t, "NSObject*", dt.KeyType.ObjcName())
	assert.Equal(t, "NSObject*", dt.ValueType.ObjcName())

	// block
	bt := ParseType("void (^) ()").(*typing.BlockType)
	assert.Equal(t, "void", bt.ReturnType.ObjcName())
	assert.Equal(t, 0, len(bt.Params))

	bt = ParseType("BOOL (^) (NSUInteger idx, BOOL * stop)").(*typing.BlockType)
	assert.Equal(t, "BOOL", bt.ReturnType.ObjcName())
	assert.Equal(t, 2, len(bt.Params))
	assert.Equal(t, "idx", bt.Params[0].Name)
	assert.Equal(t, "NSUInteger", bt.Params[0].Type.ObjcName())

	bt = ParseType("objectivec.NSObject * (^) (objectivec.NSObject * param1, NSArray<Foundation.NSExpression *> * param2, Foundation.NSMutableDictionary * param3)").(*typing.BlockType)
	assert.Equal(t, "NSObject*", bt.ReturnType.ObjcName())
	assert.Equal(t, 3, len(bt.Params))
	assert.Equal(t, "param1", bt.Params[0].Name)
	assert.Equal(t, "NSObject*", bt.Params[0].Type.ObjcName())

}
