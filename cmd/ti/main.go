// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/spf13/pflag"
	ti "github.com/zchee/the_titanium_searcher"
)

func Parse(o *ti.Option) {
	// TODO(zchee): implements self flag parser
	// TODO(zchee): uppercase of lowercase for description
	pflag.BoolVarP(&o.Version, "version", "v", false, "show ti version")
	pflag.BoolVar(&o.Color.Enable, "color", true, "print color output")
	pflag.StringVar(&o.Color.LineNumber, "color-line-number", "yellow", "color code for line number. color name or escape sequence code")
	pflag.StringVar(&o.Color.Path, "color-path", "green", "color code for line number. color name or escape sequence code")
	pflag.StringVar(&o.Color.Match, "color-match", "bgyellow", "color code for line number. color name or escape sequence code")
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

	pflag.BoolVar(&o.Regexp, "e", false, "Parse PATTERN as a regular expression")
	pflag.BoolVarP(&o.IgnoreCase, "ignore-case", "i", false, "Match case insensitively")
	pflag.BoolVarP(&o.SmartCase, "smart-case", "S", false, "Match case insensitively unless PATTERN contains uppercase characters")
	pflag.BoolVarP(&o.WordRegexp, "word-regexp", "w", false, "Only match whole words")
	pflag.StringSliceVar(&o.Ignore, "ignore", nil, "Ignore files/directories matching pattern")
	pflag.StringSliceVar(&o.VCSIgnore, "vcs-ignore", []string{".gitignore"}, "VCS ignore files")
	pflag.BoolVar(&o.GlobalGitIgnore, "global-ignore", false, "Use git's global gitignore file for ignore patterns")
	pflag.BoolVar(&o.TIIgnore, "ti-ignore", false, "Use .tiignore config file for ignore patterns")
	pflag.BoolVar(&o.SkipVCSIgnores, "skip-vcs-ignores", false, "Don't use VCS ignore file for ignore patterns")
	pflag.BoolVar(&o.FilesWithRegexp, "g", false, "Print filenames matching PATTERN")
	pflag.StringVarP(&o.FileSearchRegexp, "file-search-regexp", "G", "", "PATTERN Limit search to filenames matching PATTERN")
	pflag.IntVar(&o.Depth, "depth", 25, "Search up to NUM directories deep")
	pflag.BoolVarP(&o.Follow, "follow", "f", false, "Follow symlinks")
	pflag.BoolVarP(&o.Hidden, "hidden", "a", false, "search hidden files and directories")

	pflag.Parse()
}
