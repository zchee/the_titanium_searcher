// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ti

import (
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	colorpkg "github.com/fatih/color"
)

const (
	debug = false
)

func init() {
	log.SetFlags(log.Lshortfile)
	log.SetPrefix(colorpkg.HiBlackString("[DEBUG]: "))
}

// Debug prints debug log.
// Arguments are handled in the manner of fmt.Print.
func Debug(a ...interface{}) {
	if debug {
		log.Output(2, fmt.Sprint(a...))
	}
}

// Debugf prints debug log.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, a ...interface{}) {
	if debug {
		log.Output(2, fmt.Sprintf(format, a...))
	}
}

// Debugln prints debug log.
// Arguments are handled in the manner of fmt.Println.
func Debugln(a ...interface{}) {
	if debug {
		log.Output(2, fmt.Sprintln(a...))
	}
}

// Dump dumps a and prints debug log.
// Arguments are handled in the manner of fmt.Print.
func Dump(a ...interface{}) {
	if debug {
		log.Output(2, spew.Sdump(a...))
	}
}

// Profile the end time of the function.
func Profile(name string, now time.Time) {
	if debug {
		log.Output(2, fmt.Sprintf("%s: %fsec", name, time.Since(now).Seconds()))
	}
}
