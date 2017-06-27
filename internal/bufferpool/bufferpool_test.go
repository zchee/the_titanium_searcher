// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufferpool

import (
	"bytes"
	"reflect"
	"sync"
	"testing"
)

func Test_alloc(t *testing.T) {
	tests := []struct {
		name string
		want interface{}
	}{
		{
			name: "new",
			want: new(bytes.Buffer),
		},
		{
			name: "reference",
			want: &bytes.Buffer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alloc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("alloc() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO(zchee): testcase
func TestNew(t *testing.T) {}

func TestBufferPool_Get(t *testing.T) {
	type fields struct {
		pool *sync.Pool
	}
	tests := []struct {
		name   string
		fields fields
		want   *bytes.Buffer
	}{
		{
			name:   "new",
			fields: fields{pool: &sync.Pool{New: alloc}},
			want:   new(bytes.Buffer),
		},
		{
			name:   "reference",
			fields: fields{pool: &sync.Pool{New: alloc}},
			want:   &bytes.Buffer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bp := &BufferPool{
				pool: tt.fields.pool,
			}
			if got := bp.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BufferPool.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBufferPool_Put(t *testing.T) {
	type fields struct {
		pool *sync.Pool
	}
	type args struct {
		buf *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test",
			fields: fields{
				pool: &sync.Pool{New: alloc},
			},
			args: args{buf: new(bytes.Buffer)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bp := &BufferPool{
				pool: tt.fields.pool,
			}
			bp.Put(tt.args.buf)
		})
	}
}

// TODO(zchee): naming
func TestBufferPool_UnitTest(t *testing.T) {
	type fields struct {
		pool *sync.Pool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "test",
			fields: fields{
				pool: &sync.Pool{New: alloc},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bp := &BufferPool{
				pool: tt.fields.pool,
			}
			got := bp.Get()
			got.Write([]byte("test"))
			bp.Put(got)
			if got2 := bp.Get(); !((got2.Len() == 0) == tt.want) {
				t.Errorf("BufferPool unittest = %v, want %v", got2, tt.want)
			}
		})
	}
}
