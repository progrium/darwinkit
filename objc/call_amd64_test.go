// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"testing"
	"unsafe"
)

// todo(mkrautz): move to call_amd64.go?
const (
	nint     = 6
	nfloat   = 8
	nregular = nint + nfloat
	nstack   = 5
)

func TestFrameFetcherNoStack(t *testing.T) {
	frame := &amd64frame{
		rdi:  0,
		rsi:  1,
		rdx:  2,
		rcx:  3,
		r8:   4,
		r9:   5,
		xmm0: 0,
		xmm1: 1,
		xmm2: 2,
		xmm3: 3,
		xmm4: 4,
		xmm5: 5,
		xmm6: 6,
		xmm7: 7,
	}
	fetcher := frameFetcher(frame)
	for i := 0; i < 6; i++ {
		if val := fetcher.Int(); val != uintptr(i) {
			t.Errorf("fetched int: got %v, expected %v", val, i)
		}
	}
	for i := 0; i < 8; i++ {
		if val := fetcher.Float(); val != uintptr(i) {
			t.Errorf("fetched float: got %v, expected %v", val, i)
		}
	}
}

func TestFrameFetcherStack(t *testing.T) {
	var buf [nregular + nstack]uintptr
	for i := 0; i < nstack; i++ {
		buf[nregular+i] = uintptr(i)
	}
	frame := (*amd64frame)(unsafe.Pointer(&buf))
	fetcher := frameFetcher(frame)
	for i := 0; i < nstack; i++ {
		if val := fetcher.Stack(); val != uintptr(i) {
			t.Errorf("fetched stack: got %v, expected %v", val, i)
		}
	}
}

func TestFrameFetcherStackIntFallback(t *testing.T) {
	var buf [nregular + nstack]uintptr
	for i := 0; i < nstack; i++ {
		buf[nregular+i] = uintptr(i)
	}
	frame := (*amd64frame)(unsafe.Pointer(&buf))
	fetcher := frameFetcher(frame)
	for i := 0; i < nint; i++ {
		_ = fetcher.Int()
	}
	for i := 0; i < nstack; i++ {
		if val := fetcher.Int(); val != uintptr(i) {
			t.Errorf("int stack fallback: got %v, expected %v", val, i)
		}
	}
}

func TestFrameFetcherStackFloatFallback(t *testing.T) {
	var buf [nregular + nstack]uintptr
	for i := 0; i < nstack; i++ {
		buf[nregular+i] = uintptr(i)
	}
	frame := (*amd64frame)(unsafe.Pointer(&buf))
	fetcher := frameFetcher(frame)
	for i := 0; i < nfloat; i++ {
		_ = fetcher.Float()
	}
	for i := 0; i < nstack; i++ {
		if val := fetcher.Float(); val != uintptr(i) {
			t.Errorf("float stack fallback: got %v, expected %v", val, i)
		}
	}
}

func TestFrameFetcherStackMixedFallback(t *testing.T) {
	var buf [nregular + nstack]uintptr
	for i := 0; i < nstack; i++ {
		buf[nregular+i] = uintptr(i)
	}
	frame := (*amd64frame)(unsafe.Pointer(&buf))
	fetcher := frameFetcher(frame)
	for i := 0; i < nint; i++ {
		_ = fetcher.Int()
	}
	for i := 0; i < nfloat; i++ {
		_ = fetcher.Float()
	}
	for i := 0; i < nstack; i++ {
		if i%2 == 0 {
			if val := fetcher.Float(); val != uintptr(i) {
				t.Errorf("mixed (float) stack fallback: got %v, expected %v", val, i)
			}
		} else {
			if val := fetcher.Int(); val != uintptr(i) {
				t.Errorf("mixed (int) stack fallback: got %v, expected %v", val, i)
			}
		}
	}
}
