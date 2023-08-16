// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import "sync"

type SyncCache[K comparable, V any] struct {
	m           sync.Map
	computeLock sync.Mutex
}

func (c *SyncCache[K, V]) Load(key K, Compute func(key K) V) V {
	v, ok := c.m.Load(key)
	if ok {
		return v.(V)
	}
	c.computeLock.Lock()
	defer c.computeLock.Unlock()
	v, ok = c.m.Load(key)
	if ok {
		return v.(V)
	}

	v = Compute(key)
	c.m.Store(key, v)
	return v.(V)
}
