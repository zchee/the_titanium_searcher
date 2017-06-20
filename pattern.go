// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ti

type pattern struct {
	pattern []byte
	re      *regexp
	opt     Option
}
