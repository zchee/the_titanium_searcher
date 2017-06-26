// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufferpool

import (
	"bytes"
	"sync"
)

var g = New()

// BufferPool represents a bytes.Buffer pooling using sync.Pool.
type BufferPool struct {
	pool sync.Pool
}

func alloc() interface{} {
	return new(bytes.Buffer)
}

// New returns the new BufferPool.
func New() *BufferPool {
	var b BufferPool
	b.pool.New = alloc
	return &b
}

// Get returns the get bytes.Buffer pointer from sync.Pool.
func (bp *BufferPool) Get() *bytes.Buffer {
	return g.pool.Get().(*bytes.Buffer)
}

// Put puts the bytes.Buffer pointer to sync.Pool.
func (bp *BufferPool) Put(buf *bytes.Buffer) {
	buf.Reset()
	g.pool.Put(buf)
}
