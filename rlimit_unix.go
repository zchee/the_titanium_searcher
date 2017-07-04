// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

package ti

import (
	"syscall"
)

// rlimitNoFileSize returns the RLIMIT_NOFILE size.
func rlimitNoFileSize() uint64 {
	size := uint64(256) // 256 is macOS default size
	rlimit := syscall.Rlimit{}
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit); err == nil {
		size = rlimit.Cur
	}
	return size
}
