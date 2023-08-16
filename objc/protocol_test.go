// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_stringConvert(t *testing.T) {
	var s = "test"
	ptr := ToNSString(s)
	ns := ToGoString(ptr)
	if ns != s {
		t.Fail()
	}
}

func Test_bytesConvert(t *testing.T) {
	var bs = []byte("test")
	ptr := ToNSData(bs)
	nbs := ToGoBytes(ptr)
	if bytes.Compare(bs, nbs) != 0 {
		t.Fail()
	}
}

func Test_sliceConvert(t *testing.T) {
	s := []string{"1", "2"}
	ptr := ToNSArray(reflect.ValueOf(s))
	ns := ToGoSlice(ptr, reflect.TypeOf(s)).Interface().([]string)
	if len(s) != len(ns) || s[0] != ns[0] {
		t.Fail()
	}

	bs := [][]byte{[]byte("test")}
	ptr = ToNSArray(reflect.ValueOf(bs))
	nbs := ToGoSlice(ptr, reflect.TypeOf(bs)).Interface().([][]byte)
	if len(bs) != len(nbs) || bytes.Compare(bs[0], nbs[0]) != 0 {
		t.Fail()
	}
}

func Test_mapConvert(t *testing.T) {
	o1 := NewObject()
	o2 := NewObject()
	m := map[string]Object{
		"1": o1,
		"2": o2,
	}
	ptr := ToNSDict(reflect.ValueOf(m))
	nm := ToGoMap(ptr, reflect.TypeOf(map[string]Object{})).Interface().(map[string]Object)
	if len(m) != len(nm) || m["1"] != nm["1"] {
		t.Fail()
	}
}

func Test_selectorToGoName(t *testing.T) {
	type args struct {
		sel string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{sel: "menuWillOpen:"},
			want: "MenuWillOpen",
		},
		{
			args: args{sel: "menu:updateItem:atIndex:shouldCancel:"},
			want: "MenuUpdateItemAtIndexShouldCancel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectorToGoName(tt.args.sel); got != tt.want {
				t.Errorf("selectorToGoName() = %v, want %v", got, tt.want)
			}
		})
	}
}
