package foundation

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/progrium/darwinkit/objc"
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

// Test the (NS)Error type
func TestFoundationError(t *testing.T) {
	t.Run("Empty Map", func(t *testing.T) {

		// declare the variables
		var domain ErrorDomain = "My Domain"
		code := -99
		userInfo := make(map[ErrorUserInfoKey]objc.IObject)

		// create the (NS)Error object
		myError := Error_ErrorWithDomainCodeUserInfo(domain, code, userInfo)

		// check that length of map is zero
		length := len(myError.UserInfo())
		if length != 0 {
			t.Log("Error: Empty map length is not zero")
			t.Fail()
		}
	})

	t.Run("Map with items in it", func(t *testing.T) {

		// declare the variables
		var domain ErrorDomain = "My Domain"
		code := -99
		userInfo := make(map[ErrorUserInfoKey]objc.IObject)
		userInfo["one"] = String_StringWithString("ONE")
		userInfo["two"] = String_StringWithString("TWO")
		userInfo["three"] = String_StringWithString("THREE")
		userInfo["four"] = String_StringWithString("FOUR")
		userInfo["five"] = String_StringWithString("FIVE")

		// create the (NS)Error object
		myError := Error_ErrorWithDomainCodeUserInfo(domain, code, userInfo)

		// check each key with its value to see if they go together
		for key, value := range myError.UserInfo() {
			goKey := fmt.Sprintf("%s", reflect.ValueOf(key))
			goValue := string(value.Description())
			if strings.ToUpper(goKey) != goValue {
				message := fmt.Sprintf("Failure detected: %s != %s", strings.ToUpper(goKey), goValue)
				t.Log(message)
				t.Fail()
			}
		}
	})

	t.Run("Check domain and code parameters", func(t *testing.T) {

		// declare the variables
		var domain ErrorDomain = "My Domain"
		code := -99
		userInfo := make(map[ErrorUserInfoKey]objc.IObject)

		// create the (NS)Error object
		myError := Error_ErrorWithDomainCodeUserInfo(domain, code, userInfo)

		// check domain
		checkDomain := myError.Domain()
		if domain != checkDomain {
			message := fmt.Sprintf("Expected domain value %s - actual %s", domain, checkDomain)
			t.Log(message)
			t.Fail()
		}

		// check code
		checkCode := myError.Code()
		if code != checkCode {
			message := fmt.Sprintf("Expected code value %d - actual %d", code, checkCode)
			t.Log(message)
			t.Fail()
		}
	})
}
