// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"testing"

	"github.com/progrium/macdriver/internal/assert"
)

func TestMsgSend(t *testing.T) {
	o := NewObject()
	o.RetainCount()
	count := MsgSend[uint](o, GetSelector("retainCount"))
	assert.Equal(t, uint(1), count)
}
