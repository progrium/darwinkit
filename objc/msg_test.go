package objc

import (
	"math"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"testing"
	"unsafe"

	"github.com/progrium/macdriver/internal/objctest"
	"github.com/progrium/macdriver/misc/variadic"
)

func TestMain(m *testing.M) {
	pool := GetClass("NSAutoreleasePool").Alloc().Init()
	defer pool.Release()

	os.Exit(m.Run())
}

func BenchmarkSendMsg(b *testing.B) {
	obj := Get("NSObject").Send("new")
	defer obj.Release()

	for i := 0; i < b.N; i++ {
		sendMsg(obj, variadic.F_msgSend, "description")
	}
}

func TestNSStringSendMsg(t *testing.T) {
	cases := []string{
		"",
		"Hello world!",
		"foobar",
	}

	t.Run("C2Go", func(t *testing.T) {
		for _, v := range cases {
			t.Run("", func(t *testing.T) {
				o := objctest.NSStringWith(v)
				u := ObjectPtr(uintptr(o)).Send("UTF8String").CString()

				if v != u {
					t.Fatalf("expected %q, got %q", v, u)
				}
			})
		}
	})

	t.Run("Go2C", func(t *testing.T) {
		t.Skip("Send can't handle strings")
		for _, v := range cases {
			t.Run("", func(t *testing.T) {
				o := Get("NSString").Send("stringWithUTF8String:", v)
				u := objctest.NSStringValue(unsafe.Pointer(o.Pointer()))

				if v != u {
					t.Fatalf("expected %q, got %q", v, u)
				}
			})
		}
	})

	t.Run("Go2Go", func(t *testing.T) {
		t.Skip("Send can't handle strings")
		for _, v := range cases {
			t.Run("", func(t *testing.T) {
				u := Get("NSString").Send("stringWithUTF8String:", v).Send("UTF8String").String()

				if v != u {
					t.Fatalf("expected %q, got %q", v, u)
				}
			})
		}
	})
}

func TestNSNumberSendMsg(t *testing.T) {
	const N = 100

	type Case struct {
		name    string
		value   func() interface{}
		numWith interface{}
		obj2go  interface{}
		numVal  interface{}
	}

	cases := []Case{
		{"Bool", func() interface{} { return rand.Intn(2) != 0 }, objctest.NSNumberWithBool, Object.Bool, objctest.NSNumberBoolValue},
		{"Char", func() interface{} { return int8(rand.Int63()) }, objctest.NSNumberWithInt8, Object.Int, objctest.NSNumberInt8Value},
		{"Short", func() interface{} { return int16(rand.Int63()) }, objctest.NSNumberWithInt16, Object.Int, objctest.NSNumberInt16Value},
		{"Int", func() interface{} { return int32(rand.Int63()) }, objctest.NSNumberWithInt32, Object.Int, objctest.NSNumberInt32Value},
		{"LongLong", func() interface{} { return int64(rand.Int63()) }, objctest.NSNumberWithInt64, Object.Int, objctest.NSNumberInt64Value},
		{"UnsignedChar", func() interface{} { return uint8(rand.Uint64()) }, objctest.NSNumberWithUint8, Object.Uint, objctest.NSNumberUint8Value},
		{"UnsignedShort", func() interface{} { return uint16(rand.Uint64()) }, objctest.NSNumberWithUint16, Object.Uint, objctest.NSNumberUint16Value},
		{"UnsignedInt", func() interface{} { return uint32(rand.Uint64()) }, objctest.NSNumberWithUint32, Object.Uint, objctest.NSNumberUint32Value},
		{"UnsignedLongLong", func() interface{} { return uint64(rand.Uint64()) }, objctest.NSNumberWithUint64, Object.Uint, objctest.NSNumberUint64Value},
		{"Float", func() interface{} { return float32(rand.Float32()) }, objctest.NSNumberWithFloat32, Object.Float, objctest.NSNumberFloat32Value},
		{"Double", func() interface{} { return float64(rand.Float64()) }, objctest.NSNumberWithFloat64, Object.Float, objctest.NSNumberFloat64Value},
	}

	floatCases := map[string]float64{
		"Zero":   0,
		"NaN":    math.NaN(),
		"PosInf": math.Inf(+1),
		"NegInf": math.Inf(-1),
	}

	runFloat := func(run func(Case), v float64) {
		run(Case{"Double", func() interface{} { return v }, objctest.NSNumberWithFloat64, Object.Float, objctest.NSNumberFloat64Value})
	}

	selValue := func(c Case) string { return strings.ToLower(c.name[:1]) + c.name[1:] + "Value" }
	selWith := func(c Case) string { return "numberWith" + c.name + ":" }

	t.Run("C2Go", func(t *testing.T) {
		run := func(c Case) {
			numWith := reflect.ValueOf(c.numWith)
			obj2go := reflect.ValueOf(c.obj2go)

			v := reflect.ValueOf(c.value())
			ptr := numWith.Call([]reflect.Value{v})[0]
			ret := ObjectPtr(ptr.Pointer()).Send(selValue(c))
			u := obj2go.Call([]reflect.Value{reflect.ValueOf(ret)})[0]

			u = u.Convert(v.Type())
			if reflect.DeepEqual(v, u) {
				t.Fatalf("expected %v, got %v", v, u)
			}
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				for i := 0; i < N; i++ {
					run(c)
				}
			})
		}

		for name, v := range floatCases {
			t.Run("Float"+name, func(t *testing.T) {
				runFloat(run, v)
			})
		}
	})

	t.Run("Go2C", func(t *testing.T) {
		run := func(c Case) {
			numVal := reflect.ValueOf(c.numVal)

			v := reflect.ValueOf(c.value())
			o := Get("NSNumber").Send(selWith(c), v.Interface())
			u := numVal.Call([]reflect.Value{reflect.ValueOf(unsafe.Pointer(o.Pointer()))})[0]

			u = u.Convert(v.Type())
			if reflect.DeepEqual(v, u) {
				t.Fatalf("expected %v, got %v", v, u)
			}
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				for i := 0; i < N; i++ {
					run(c)
				}
			})
		}

		for name, v := range floatCases {
			t.Run("Float"+name, func(t *testing.T) {
				runFloat(run, v)
			})
		}
	})

	t.Run("Go2Go", func(t *testing.T) {
		run := func(c Case) {
			obj2go := reflect.ValueOf(c.obj2go)

			v := reflect.ValueOf(c.value())
			ret := Get("NSNumber").Send(selWith(c), v.Interface()).Send(selValue(c))
			u := obj2go.Call([]reflect.Value{reflect.ValueOf(ret)})[0]

			u = u.Convert(v.Type())
			if reflect.DeepEqual(v, u) {
				t.Fatalf("expected %v, got %v", v, u)
			}
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				for i := 0; i < N; i++ {
					run(c)
				}
			})
		}

		for name, v := range floatCases {
			t.Run("Float"+name, func(t *testing.T) {
				runFloat(run, v)
			})
		}
	})
}
