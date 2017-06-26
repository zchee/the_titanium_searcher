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
			name: "value",
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
