package bridge

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) NSPoint() core.NSPoint {
	return core.NSPoint{X: p.X, Y: p.Y}
}

type Size struct {
	W float64
	H float64
}

func (s *Size) NSSize() core.NSSize {
	return core.NSSize{Width: s.W, Height: s.H}
}

type Color struct {
	R float64
	G float64
	B float64
	A float64
}

func (c *Color) NSColor() cocoa.NSColor {
	return cocoa.NSColor_Init(c.R, c.G, c.B, c.A)
}

//////// unused?

func walk(v reflect.Value, path []string, visitor func(v reflect.Value, parent reflect.Value, path []string) error) error {
	for _, k := range keys(v) {
		subpath := append(path, k)
		vv := prop(v, k)
		if !vv.IsValid() {
			continue
		}
		if err := visitor(vv, v, subpath); err != nil {
			return err
		}
		if err := walk(vv, subpath, visitor); err != nil {
			return err
		}
	}
	return nil
}

func Walk(v interface{}, visitor func(v reflect.Value, parent reflect.Value, path []string) error) error {
	return walk(reflect.ValueOf(v), []string{}, visitor)
}

func prop(robj reflect.Value, key string) reflect.Value {
	rtyp := robj.Type()
	switch rtyp.Kind() {
	case reflect.Slice, reflect.Array:
		idx, err := strconv.Atoi(key)
		if err != nil {
			panic("non-numeric index given for slice")
		}
		rval := robj.Index(idx)
		if rval.IsValid() {
			return reflect.ValueOf(rval.Interface())
		}
	case reflect.Ptr:
		return prop(robj.Elem(), key)
	case reflect.Map:
		rval := robj.MapIndex(reflect.ValueOf(key))
		if rval.IsValid() {
			return reflect.ValueOf(rval.Interface())
		}
	case reflect.Struct:
		rval := robj.FieldByName(key)
		if rval.IsValid() {
			return rval
		}
		for i := 0; i < rtyp.NumField(); i++ {
			field := rtyp.Field(i)
			tag := strings.Split(field.Tag.Get("json"), ",")
			if tag[0] == key || field.Name == key {
				return robj.FieldByName(field.Name)
			}
		}
		panic("struct field not found: " + key)
	}
	//spew.Dump(robj, key)
	panic("unexpected kind: " + rtyp.Kind().String())
}

func keys(v reflect.Value) []string {
	switch v.Type().Kind() {
	case reflect.Map:
		var keys []string
		for _, key := range v.MapKeys() {
			k, ok := key.Interface().(string)
			if !ok {
				continue
			}
			keys = append(keys, k)
		}
		sort.Sort(sort.StringSlice(keys))
		return keys
	case reflect.Struct:
		t := v.Type()
		var f []string
		for i := 0; i < t.NumField(); i++ {
			name := t.Field(i).Name
			// first letter capitalized means exported
			if name[0] == strings.ToUpper(name)[0] {
				f = append(f, name)
			}
		}
		return f
	case reflect.Slice, reflect.Array:
		var k []string
		for n := 0; n < v.Len(); n++ {
			k = append(k, strconv.Itoa(n))
		}
		return k
	case reflect.Ptr:
		if !v.IsNil() {
			return keys(v.Elem())
		}
		return []string{}
	case reflect.String, reflect.Bool, reflect.Float64, reflect.Float32, reflect.Interface:
		return []string{}
	default:
		fmt.Fprintf(os.Stderr, "unexpected type: %s\n", v.Type().Kind())
		return []string{}
	}
}
