// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import "reflect"

type UInteger = uint64
type Integer = int64
type Void struct{}

var voidType = reflect.TypeOf((*Void)(nil)).Elem()
