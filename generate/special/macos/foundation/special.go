// Description: where methods that can't work when generated are manually implemented.

package foundation

import (
	"reflect"
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>

// makes a NSMutableDictionary from two NSArrays
 void *MakeMutableDictionaryInternal(void *valuesPtr, void *keysPtr) {
	 NSArray *keysArray = (NSArray *)keysPtr;
	 NSArray *valuesArray = (NSArray *)valuesPtr;
	 NSMutableDictionary *mdict = [[NSMutableDictionary alloc] init];
	 for(int i = 0; i < [keysArray count]; i++) {
		  [mdict setObject:[valuesArray objectAtIndex: i] forKey:[keysArray objectAtIndex: i]];
	 }
	 return mdict;
 }

// makes a NSMutableDictionary from two NSArrays
void *MakeDictionaryInternal(void *valuesPtr, void *keysPtr) {
	void *mdict = MakeMutableDictionaryInternal(valuesPtr, keysPtr);
	return [NSDictionary dictionaryWithDictionary: mdict];
}
*/
import "C"

// tell the generation system not to generate these methods
/*
Note: fields: class selector note
begin-skip
"NSDictionary", "initWithObjectsAndKeys:", "using custom implementation"
"NSDictionary", "dictionaryWithObjectsAndKeys:", "using custom implementation"
"NSMutableDictionary", "initWithObjectsAndKeys:", "using custom implementation"
"NSMutableDictionary", "dictionaryWithObjectsAndKeys:", "using custom implementation"
"NSArray", "arrayWithObjects:", "using custom implementation"
"NSArray", "initWithObjects:", "using custom implementation"
"NSMutableArray", "arrayWithObjects:", "using custom implementation"
"NSMutableArray", "initWithObjects:", "using custom implementation"
end-skip
*/

// create an array for values and a array for keys using one array
func createKeyValueArrays(args []objc.IObject) (unsafe.Pointer, unsafe.Pointer) {
	keys := make([]objc.IObject, 0)
	values := make([]objc.IObject, 0)
	for i := 0; i < len(args); i++ {
		if i%2 == 0 {
			// if last item is nil
			if i == len(args)-1 && args[i] == nil {
				continue // remove terminating nil
			}
			values = append(values, args[i])
		} else {
			keys = append(keys, args[i])
		}
	}
	valuesArray := objc.ToNSArray(reflect.ValueOf(values))
	keysArray := objc.ToNSArray(reflect.ValueOf(keys))
	return keysArray, valuesArray
}

// Uses key-value pairs to make a dictionary
func makeDictionary(args []objc.IObject) Dictionary {
	keysArray, valuesArray := createKeyValueArrays(args)
	pointer := C.MakeDictionaryInternal(valuesArray, keysArray)
	dict := DictionaryFrom(pointer)
	return dict
}

// Uses key-value pairs to make a mutable dictionary
func makeMutableDictionary(args []objc.IObject) MutableDictionary {
	keysArray, valuesArray := createKeyValueArrays(args)
	pointer := C.MakeMutableDictionaryInternal(valuesArray, keysArray)
	mdict := MutableDictionaryFrom(pointer)
	return mdict
}

/** Dictionary class methods **/

// implemented here because Go's nil causes an exception to be thrown
func (a_ Dictionary) InitWithObjectsAndKeys(firstObject objc.IObject, args ...objc.IObject) Dictionary {
	args = append([]objc.IObject{firstObject}, args...)
	dict := makeDictionary(args)
	return dict
}

// implemented here because Go's nil causes an exception to be thrown
func Dictionary_DictionaryWithObjectsAndKeys(firstObject objc.IObject, args ...objc.IObject) Dictionary {
	args = append([]objc.IObject{firstObject}, args...)
	dict := makeDictionary(args)
	return dict
}

func NewDictionaryWithObjectsAndKeys(firstObject objc.IObject, args ...objc.IObject) Dictionary {
	args = append([]objc.IObject{firstObject}, args...)
	dict := makeDictionary(args)
	return dict
}

/** MutableDictionary class methods **/

// implemented here because Go's nil causes an exception to be thrown
func (a_ MutableDictionary) InitWithObjectsAndKeys(firstObject objc.IObject, args ...objc.IObject) MutableDictionary {
	args = append([]objc.IObject{firstObject}, args...)
	mdict := makeMutableDictionary(args)
	return mdict
}

// implemented here because Go's nil causes an exception to be thrown
func MutableDictionary_DictionaryWithObjectsAndKeys(firstObject objc.IObject, args ...objc.IObject) MutableDictionary {
	args = append([]objc.IObject{firstObject}, args...)
	mdict := makeMutableDictionary(args)
	return mdict
}

func NewMutableDictionaryWithObjectsAndKeys(firstObject objc.IObject, args ...objc.IObject) MutableDictionary {
	args = append([]objc.IObject{firstObject}, args...)
	mdict := makeMutableDictionary(args)
	return mdict
}

/** Array class methods **/

// implemented here because only one object appears in the array otherwise
func Array_ArrayWithObjects(firstObject objc.IObject, args ...objc.IObject) Array {
	args = append([]objc.IObject{firstObject}, args...)
	arrayPtr := objc.ToNSArray(reflect.ValueOf(args))
	newArray := ArrayFrom(arrayPtr)
	return newArray
}

// implemented here because only one object appears in the array otherwise
func (a_ Array) InitWithObjects(firstObj objc.IObject, args ...objc.IObject) Array {
	args = append([]objc.IObject{firstObj}, args...)
	arrayPtr := objc.ToNSArray(reflect.ValueOf(args))
	newArray := ArrayFrom(arrayPtr)
	return newArray
}

// implemented here because only one object appears in the array otherwise
func NewArrayWithObjects(firstObj objc.IObject, args ...objc.IObject) Array {
	args = append([]objc.IObject{firstObj}, args...)
	arrayPtr := objc.ToNSArray(reflect.ValueOf(args))
	newArray := ArrayFrom(arrayPtr)
	return newArray
}

/** MutableArray class methods **/

// implemented here because only one object appears in the array otherwise
func MutableArray_ArrayWithObjects(firstObj objc.IObject, args ...objc.IObject) MutableArray {
	args = append([]objc.IObject{firstObj}, args...)
	mArray := MutableArray_ArrayWithArray(args)
	return mArray
}

// implemented here because only one object appears in the array otherwise
func (a_ MutableArray) InitWithObjects(firstObj objc.IObject, args ...objc.IObject) MutableArray {
	args = append([]objc.IObject{firstObj}, args...)
	mArray := MutableArray_ArrayWithArray(args)
	return mArray
}

// implemented here because only one object appears in the array otherwise
func NewMutableArrayWithObjects(firstObj objc.IObject, args ...objc.IObject) MutableArray {
	args = append([]objc.IObject{firstObj}, args...)
	mArray := MutableArray_ArrayWithArray(args)
	return mArray
}
