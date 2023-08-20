package foundation

import (
	"reflect"
	"testing"
)

func TestFoundationValid(t *testing.T) {}

func TestFoundationDictionary(t *testing.T) {
	s := String_StringWithString("bar")
	d1 := Dictionary_DictionaryWithObjectsAndKeys(s, "foo", "value", "key", nil)
	m1 := DictToMap[string, string](d1)
	d2 := DictOf(m1)
	m2 := DictToMap[string, string](d2)
	if !reflect.DeepEqual(m2, map[string]string{
		"foo": "bar",
		"key": "value",
	}) {
		t.Fatal("unexpected final map from dictionary")
	}
}

func TestFoundationArray(t *testing.T) {
	s := String_StringWithString("one")
	a1 := Array_ArrayWithObjects(s, "two", "three", nil)
	s1 := ArrayToSlice[string](a1)
	a2 := ArrayOf(s1)
	s2 := ArrayToSlice[string](a2)
	if !reflect.DeepEqual(s2, []string{
		"one",
		"two",
		"three",
	}) {
		t.Fatal("unexpected final slice from array")
	}
}
