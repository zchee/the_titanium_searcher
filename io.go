// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ti

import (
	"os"
	"path/filepath"

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

// FileDescriptor opens path and returns the file descriptor if path is non-nil.
// If path is empty, returns the stdin file descriptor.
func FileDescriptor(path string) (*os.File, error) {
	if path == "" {
		return os.Stdin, nil
	}
	return os.Open(path)
}
