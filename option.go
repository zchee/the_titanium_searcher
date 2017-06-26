// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ti

// Option represents the_titanium_searcher option.
type Option struct {
	Version bool
	Profile string

	Output OutputOption
	Search SearchOption
}

// OutputOption represents the_titanium_searcher output option.
type OutputOption struct {
	EnableColor bool
	Color       struct {
		Match  string
		Number string
		Path   string
	}
	Group   bool
	Null    bool
	Column  bool
	Number  bool
	Context struct {
		After   int
		Before  int
		Context int
	}
	FilesWithMatches bool
	Count            bool
	Encode           string
}

// SearchOption represents the_titanium_searcher search option.
type SearchOption struct {
	Regexp           bool
	IgnoreCase       bool
	SmartCase        bool
	WordRegexp       bool
	Ignore           []string
	VCSIgnore        []string
	GlobalGitIgnore  bool
	TIIgnore         bool
	SkipVCSIgnores   bool
	FilesWithRegexp  bool
	FileSearchRegexp string
	Depth            int
	Follow           bool
	Hidden           bool
}

// NewOption returns the Option with pointer.
func NewOption() *Option {
	return &Option{
		Output: OutputOption{},
		Search: SearchOption{},
	}
}
