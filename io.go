// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ti

import (
	"os"
	"path/filepath"
	"syscall"

	xdgbasedir "github.com/zchee/go-xdgbasedir"
)

var (
	// ConfigHome returns the ti config home directory.
	ConfigHome = filepath.Join(xdgbasedir.ConfigHome(), "ti")
	// ConfigFile returns the ti config file path.
	ConfigFile = filepath.Join(ConfigHome, "config.yml")
	// IgnoreFile returns the ti ignore file path.
	IgnoreFile = filepath.Join(ConfigHome, "ignore")
)

// rlimitNoFileSize returns the RLIMIT_NOFILE size.
func rlimitNoFileSize() uint64 {
	const raceMax = 8192
	size := uint64(256) // 256 is macOS default size
	rlimit := syscall.Rlimit{}
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit); err == nil {
		if raceMax < rlimit.Cur {
			size = raceMax
		}
		size = rlimit.Cur
	}
	return size
}

// FileDescriptor opens path and returns the file descriptor if path is non-nil.
// If path is empty, returns the stdin file descriptor.
func FileDescriptor(path string) (*os.File, error) {
	if path == "" {
		return os.Stdin, nil
	}
	return os.Open(path)
}
