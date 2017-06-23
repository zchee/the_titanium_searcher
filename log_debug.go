// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build debug

package ti

import (
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	colorpkg "github.com/fatih/color"
)

func init() {
	log.SetFlags(log.Lshortfile)
	log.SetPrefix(colorpkg.HiBlackString("[DEBUG]: "))
}

func Debug(a ...interface{}) {
	log.Output(2, fmt.Sprint(a...))
}

func Debugf(format string, a ...interface{}) {
	log.Output(2, fmt.Sprintf(format, a...))
}

func Debugln(a ...interface{}) {
	log.Output(2, fmt.Sprintln(a...))
}

func Dump(a ...interface{}) {
	log.Output(2, spew.Sdump(a...))
}

func Profile(name string, now time.Time) {
	log.Output(2, fmt.Sprintf("%s: %fsec", name, time.Since(now).Seconds()))
}
