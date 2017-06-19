// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ti

type Option struct {
	Version bool

	Color            *color
	Group            bool
	Null             bool
	Column           bool
	LineNumber       bool
	Context          *lineContext
	FilesWithMatches bool
	Count            bool
	Encode           string

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

type color struct {
	Enable     bool
	LineNumber string
	Path       string
	Match      string
}

type lineContext struct {
	After   int
	Before  int
	Context int
}

func NewOption() *Option {
	return &Option{
		Color:   new(color),
		Context: new(lineContext),
	}
}
