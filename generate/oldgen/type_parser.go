package oldgen

import (
	"strings"

	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/generate/oldgen/data"
	"github.com/progrium/macdriver/generate/typing"
)

func ParseType(typeStr string) typing.Type {
	t := &Tokenizer{Text: typeStr}
	token := t.MustNextToken()
	var currentType typing.Type

	switch token {
	case "SEL":
		currentType = &typing.SelectorType{}
	case "void":
		currentType = &typing.VoidType{}
		token, hasToken := t.NextToken()
		if hasToken {
			if token == "*" {
				currentType = &typing.VoidPointerType{}
			} else {
				t.PutBackToken(token)
			}
		}
	case "instance":
		currentType = &typing.InstanceType{}
	case "id":
		currentType = &typing.IDType{}
		token, hasToken := t.NextToken()
		if hasToken {
			if token == "<" {
				// prototcol.
				// TODO: handle multi protocol
				protocolName := t.MustNextToken()
				currentType = loadTypeCached2(protocolName).(*typing.ProtocolType)
				t.MustNextTokenIs(">")
			} else {
				t.PutBackToken(token)
			}
		}
	case "NSString":
		currentType = &typing.StringType{}
		t.MustNextTokenIs("*")
	case "NSStringNeedNil":
		currentType = &typing.StringType{NeedNil: true}
		t.MustNextTokenIs("*")
	case "NSData":
		currentType = &typing.DataType{}
		t.MustNextTokenIs("*")
	case "NSArray":
		at := &typing.ArrayType{}
		currentType = at
		token = t.MustNextToken()
		if token == "<" {
			token = t.MustNextTillClosingAngle()
			at.Type = ParseType(token)
			t.MustNextTokenIs(">")
		} else {
			t.PutBackToken(token)
			at.Type = typing.Object
		}
		t.MustNextTokenIs("*")
	case "NSDictionary":
		dt := &typing.DictType{}
		currentType = dt
		token = t.MustNextToken()
		if token == "<" {
			token = t.MustNextTill(',')
			dt.KeyType = ParseType(token)
			t.MustNextTokenIs(",")
			token = t.MustNextTillClosingAngle()
			dt.ValueType = ParseType(token)
			t.MustNextTokenIs(">")
		} else {
			t.PutBackToken(token)
			dt.KeyType = typing.Object
			dt.ValueType = typing.Object
		}
		t.MustNextTokenIs("*")
	default:
		pt, ok := typing.GetPrimitiveType(token)
		if ok {
			currentType = pt
		} else {
			currentType = loadTypeCached2(token)
			switch currentType.(type) {
			case *typing.ClassType:
				t.MustNextTokenIs("*")
			case *typing.ProtocolType:
				panic("standalone proto type")
			case *typing.AliasType:
			case *typing.StructType:
			}
		}
	}

	// pointers
	// TODO: pointer to pointer, etc.
	token, hasToken := t.NextToken()
	if hasToken {
		if token == "*" {
			currentType = &typing.PointerType{
				Type: currentType,
			}
		} else {
			t.PutBackToken(token)
		}
	}

	// blocks
	// TODO: block type that return a block type
	token, hasToken = t.NextToken()
	if hasToken {
		if token == "(" {
			t.MustNextTokenIs("^")
			t.MustNextTokenIs(")")
			t.MustNextTokenIs("(")
			blockParams := parseBlockParams(t)
			t.MustNextTokenIs(")")
			currentType = &typing.BlockType{
				ReturnType: currentType,
				Params:     blockParams,
			}
		} else {
			t.PutBackToken(token)
		}
	}

	token, hasToken = t.NextToken()
	if hasToken {
		panic("expect end")
	}

	return currentType
}

func parseBlockParams(t *Tokenizer) []typing.BlockParam {

	var params []typing.BlockParam
	openParenthesisCount := 1
	offset := t.index
	for {
		c, ok := t.nextChar()
		if !ok {
			panic("unexpected end")
		}

		switch c {
		case '(':
			openParenthesisCount++
		case ')':
			openParenthesisCount--
			if openParenthesisCount == 0 {
				t.putBackChar()
				text := t.Text[offset:t.index]
				if len(text) > 0 {
					param := parseOneBlockParam(text)
					params = append(params, param)
				}
				return params
			}
		case ',':
			if openParenthesisCount == 1 {
				param := parseOneBlockParam(t.Text[offset : t.index-1])
				params = append(params, param)
				offset = t.index
			}
		default:
		}
	}
}

func parseOneBlockParam(str string) typing.BlockParam {
	str = strings.TrimSpace(str)
	idx := strings.LastIndexByte(str, ' ')
	if idx < 0 {
		panic("expect param name, but not found. str: " + str)
	}
	name := str[idx+1:]
	pt := ParseType(strings.TrimSpace(str[:idx]))
	return typing.BlockParam{
		Name: name,
		Type: pt,
	}
}

func loadTypeCached2(fullName string) typing.Type {
	module, name, found := strings.Cut(fullName, ".")
	if !found {
		panic("unknown type name: " + fullName)
	}

	return loadTypeCached(module, name)
}

var typeCache = map[string]typing.Type{}

func loadTypeCached(module, name string) typing.Type {
	key := module + "." + name
	t, ok := typeCache[key]
	if ok {
		return t
	}
	t = loadType(module, name)
	typeCache[key] = t
	return t
}

func loadType(module, name string) typing.Type {
	ti := loadOneTypeInfo(module, name)

	switch tti := ti.(type) {
	case *data.Class:
		ct := &typing.ClassType{
			Name:   tti.Name,
			GName:  toGoSymbolName(tti.Name),
			Module: modules.Get(tti.Module),
		}

		return ct
	case *data.Protocol:
		pt := &typing.ProtocolType{
			Name:   tti.Name,
			GName:  toGoSymbolName(tti.Name),
			Module: modules.Get(tti.Module),
		}

		return pt
	case *data.Alias:
		at := &typing.AliasType{
			Name:   tti.Name,
			GName:  toGoSymbolName(tti.Name),
			Module: modules.Get(tti.Module),
			Type:   ParseType(tti.TargetType),
		}
		return at
	case *data.Struct:
		st := &typing.StructType{
			Name:   tti.Name,
			GName:  toGoSymbolName(tti.Name),
			Module: modules.Get(tti.Module),
		}

		return st
	default:
		panic("unknown type")
	}
}

var prefixes = []string{"NS", "CG", "WK", "CA", "kCA", "UT"}

func toGoSymbolName(name string) string {
	for _, prefix := range prefixes {
		if strings.HasPrefix(name, prefix) {
			return name[len(prefix):]
		}
	}
	return name
}
