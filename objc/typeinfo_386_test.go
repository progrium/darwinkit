// Copyright (c) 2013 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var prog = `
#include <objc/runtime.h>
#include <stdio.h>
#import <Foundation/Foundation.h>

@interface TestObject : NSObject {
}
- (void) testMethod:(__TEST_TYPE__)arg;
@end

@implementation TestObject 
- (void) testMethod:(__TEST_TYPE__)arg {
}
@end

int main(int argc, char *argv[]) {
	Method m = class_getInstanceMethod([TestObject class], @selector(testMethod:));;
	printf("%s\n", method_getTypeEncoding(m));
}`

// typeInfoForCType runs clang on the `prog' test
// program above to determine the correct typeinfo
// kinds for 386.
func typeInfoForCType(ctype string) (string, error) {
	tempDir, err := ioutil.TempDir("", "typeinfo")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tempDir)

	progName := filepath.Join(tempDir, "prog")
	args := []string{
		"-x", "objective-c",
		"-arch", "i386",
		"-lobjc",
		"-D__TEST_TYPE__=" + ctype,

		"-", // stdout
		"-o", progName,
	}
	cmd := exec.Command("clang", args...)
	cmd.Stdin = bytes.NewBufferString(prog)
	buf, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(string(buf))
	}

	buf, err = exec.Command(progName).Output()
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

// expectTypeInfoMatch calls t.Fatlaf if the the ctype does
// not match the expected encoding, 'enc'.
func expectTypeInfoMatch(t *testing.T, ctype string, enc string) {
	typeInfo, err := typeInfoForCType(ctype)
	if err != nil {
		t.Fatalf("unable to exec test program: %v", err)
		return
	}

	sti := simplifyTypeInfo(typeInfo)

	prefix := sti[0:3]
	if prefix != "v@:" {
		t.Fatalf("bad prefix: %v", prefix)
	}

	actualEnc := string(sti[3])
	if enc != actualEnc {
		t.Fatalf("bad typeinfo for %v; got %v, expected %v", ctype, actualEnc, enc)
	}
}

// TestTypeInfo386 tests that our expectations for
// type encodings are actually correct.
func TestTypeInfo386(t *testing.T) {
	expectTypeInfoMatch(t, "uint64_t", encULongLong)
	expectTypeInfoMatch(t, "int64_t", encLongLong)
}
