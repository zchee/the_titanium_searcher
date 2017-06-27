// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package ti

// rlimitNoFileSize returns the RLIMIT_NOFILE size.
func rlimitNoFileSize() uint64 {
	return 1024 // TODO(zchee): hardcoded. check Windows default rlimit size
}
