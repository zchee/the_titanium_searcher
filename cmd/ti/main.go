// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/pkg/profile"
	"github.com/spf13/pflag"
	ti "github.com/zchee/the_titanium_searcher"
)

// parse parses command line flags and pass values to o.
//
// TODO(zchee): implements self flag parser
// TODO(zchee): uppercase of lowercase for description
func parse(o *ti.Option) {
	pflag.BoolVarP(&o.Version, "version", "v", false, "show ti version")
	pflag.StringVar(&o.Profile, "profile", "", "profiling")

	// Output
	pflag.BoolVar(&o.Output.EnableColor, "color", true, "print color output")
	pflag.StringVar(&o.Output.Color.Number, "color-line-number", "yellow", "color code for line number. color name or escape sequence code")
	pflag.StringVar(&o.Output.Color.Path, "color-path", "green", "color code for line number. color name or escape sequence code")
	pflag.StringVar(&o.Output.Color.Match, "color-match", "bgyellow", "color code for line number. color name or escape sequence code")
	pflag.BoolVar(&o.Output.Group, "group", true, "print file name at header")
	pflag.BoolVarP(&o.Output.Null, "null", "0", false, "separate filenames with null for 'xargs -0'")
	pflag.BoolVar(&o.Output.Column, "column", false, "print column")
	pflag.BoolVar(&o.Output.Number, "numbers", true, "print line number")
	pflag.IntVar(&o.Output.Context.After, "after", 0, "print lines after match")
	pflag.IntVar(&o.Output.Context.Before, "before", 0, "print lines before match")
	pflag.IntVar(&o.Output.Context.Context, "context", 0, "print lines before and after match")
	pflag.BoolVarP(&o.Output.FilesWithMatches, "files-with-matches", "l", false, "only print filenames that contain matches")
	pflag.BoolVarP(&o.Output.Count, "count", "c", false, "only print number of matching lines for each input file")
	pflag.StringVarP(&o.Output.Encode, "output-encode", "o", "", "specify output encodeing (none, jis, sjis, euc)")

	// Search
	pflag.BoolVarP(&o.Search.Regexp, "regexp", "e", false, "Parse PATTERN as a regular expression")
	pflag.BoolVarP(&o.Search.IgnoreCase, "ignore-case", "i", false, "Match case insensitively")
	pflag.BoolVarP(&o.Search.SmartCase, "smart-case", "S", false, "Match case insensitively unless PATTERN contains uppercase characters")
	pflag.BoolVarP(&o.Search.WordRegexp, "word-regexp", "w", false, "Only match whole words")
	pflag.StringSliceVar(&o.Search.Ignore, "ignore", nil, "Ignore files/directories matching pattern")
	pflag.StringSliceVar(&o.Search.VCSIgnore, "vcs-ignore", []string{".gitignore"}, "VCS ignore files")
	pflag.BoolVar(&o.Search.GlobalGitIgnore, "global-ignore", false, "Use git's global gitignore file for ignore patterns")
	pflag.BoolVar(&o.Search.TIIgnore, "ti-ignore", false, fmt.Sprintf("Use %s config file for ignore patterns", ti.IgnoreFile))
	pflag.BoolVar(&o.Search.SkipVCSIgnores, "skip-vcs-ignores", false, "Don't use VCS ignore file for ignore patterns")
	pflag.BoolVarP(&o.Search.FilesWithRegexp, "file-name", "g", false, "Print filenames matching PATTERN")
	pflag.StringVarP(&o.Search.FileSearchRegexp, "file-search-regexp", "G", "", "PATTERN Limit search to filenames matching PATTERN")
	pflag.IntVar(&o.Search.Depth, "depth", 25, "Search up to NUM directories deep")
	pflag.BoolVarP(&o.Search.Follow, "follow", "f", false, "Follow symlinks")
	pflag.BoolVarP(&o.Search.Hidden, "hidden", "a", false, "search hidden files and directories")

	pflag.Parse()
}
