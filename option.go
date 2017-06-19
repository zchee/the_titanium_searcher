// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ti

import (
	"github.com/spf13/pflag"
)

type Option struct {
	Version          bool
	Color            *color
	Group            bool
	Null             bool
	Column           bool
	LineNumber       bool
	Context          *lineContext
	FilesWithMatches bool
	Count            bool
	Encode           string
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

func (o *Option) Parse() {
	// TODO(zchee): implements self flag parser
	// TODO(zchee): uppercase of lowercase for description
	pflag.BoolVarP(&o.Version, "version", "v", false, "show ti version")
	pflag.BoolVar(&o.Color.Enable, "color", true, "print color output")
	pflag.StringVar(&o.Color.LineNumber, "color-line-number", "1;33", "color code for line number. escape sequence color code or color name")
	pflag.BoolVar(&o.Group, "group", true, "print file name at header")
	pflag.BoolVarP(&o.Null, "null", "0", false, "separate filenames with null for 'xargs -0'")
	pflag.BoolVar(&o.Column, "column", false, "print column")
	pflag.BoolVar(&o.LineNumber, "numbers", true, "print line number")
	pflag.IntVar(&o.Context.After, "after", 0, "print lines after match")
	pflag.IntVar(&o.Context.Before, "before", 0, "print lines before match")
	pflag.IntVar(&o.Context.Context, "context", 0, "print lines before and after match")
	pflag.BoolVarP(&o.FilesWithMatches, "files-with-matches", "l", false, "only print filenames that contain matches")
	pflag.BoolVarP(&o.Count, "count", "c", false, "only print number of matching lines for each input file")
	pflag.StringVarP(&o.Encode, "output-encode", "o", "", "specify output encodeing (none, jis, sjis, euc)")

	pflag.Parse()
}
