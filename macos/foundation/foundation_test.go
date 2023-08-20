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
