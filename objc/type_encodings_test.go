// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"reflect"
	"testing"

	"github.com/progrium/darwinkit/internal/assert"
)

type myInnerStruct struct {
	a int
	x uint8
	b float32
}

type myStruct struct {
	ms      myInnerStruct
	f       float64
	remains [8]uint16
}

func Test_getTypeEncoding(t *testing.T) {
	tp := reflect.TypeOf(myStruct{})
	assert.Equal(t, "{?={?=qCf}d[8S]}", getTypeEncoding(tp))
}
