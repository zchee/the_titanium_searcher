// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !debug

package ti

func Debug(a ...interface{})                 {}
func Debugf(format string, a ...interface{}) {}
func Debugln(a ...interface{})               {}
func Dump(a ...interface{})                  {}
