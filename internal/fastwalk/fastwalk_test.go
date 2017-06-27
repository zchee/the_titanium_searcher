// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This code copied from: golang.org/x/tools/imports

package fastwalk

import (
	"bytes"
	"flag"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"sync"
	"testing"
)

func formatFileModes(m map[string]os.FileMode) string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, k := range keys {
		fmt.Fprintf(&buf, "%-20s: %v\n", k, m[k])
	}
	return buf.String()
}

func testFastWalk(t *testing.T, files map[string]string, callback func(path string, typ os.FileMode) error, want map[string]os.FileMode) {
	testConfig{
		gopathFiles: files,
	}.test(t, func(t *goimportTest) {
		got := map[string]os.FileMode{}
		var mu sync.Mutex
		if err := Walk(t.gopath, func(path string, typ os.FileMode) error {
			mu.Lock()
			defer mu.Unlock()
			if !strings.HasPrefix(path, t.gopath) {
				t.Fatalf("bogus prefix on %q, expect %q", path, t.gopath)
			}
			key := filepath.ToSlash(strings.TrimPrefix(path, t.gopath))
			if old, dup := got[key]; dup {
				t.Fatalf("callback called twice for key %q: %v -> %v", key, old, typ)
			}
			got[key] = typ
			return callback(path, typ)
		}); err != nil {
			t.Fatalf("callback returned: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("walk mismatch.\n got:\n%v\nwant:\n%v", formatFileModes(got), formatFileModes(want))
		}
	})
}

func TestFastWalk_Basic(t *testing.T) {
	testFastWalk(t, map[string]string{
		"foo/foo.go":   "one",
		"bar/bar.go":   "two",
		"skip/skip.go": "skip",
	},
		func(path string, typ os.FileMode) error {
			return nil
		},
		map[string]os.FileMode{
			"":                  os.ModeDir,
			"/src":              os.ModeDir,
			"/src/bar":          os.ModeDir,
			"/src/bar/bar.go":   0,
			"/src/foo":          os.ModeDir,
			"/src/foo/foo.go":   0,
			"/src/skip":         os.ModeDir,
			"/src/skip/skip.go": 0,
		})
}

func TestFastWalk_Symlink(t *testing.T) {
	switch runtime.GOOS {
	case "windows", "plan9":
		t.Skipf("skipping on %s", runtime.GOOS)
	}
	testFastWalk(t, map[string]string{
		"foo/foo.go": "one",
		"bar/bar.go": "LINK:../foo.go",
		"symdir":     "LINK:foo",
	},
		func(path string, typ os.FileMode) error {
			return nil
		},
		map[string]os.FileMode{
			"":                os.ModeDir,
			"/src":            os.ModeDir,
			"/src/bar":        os.ModeDir,
			"/src/bar/bar.go": os.ModeSymlink,
			"/src/foo":        os.ModeDir,
			"/src/foo/foo.go": 0,
			"/src/symdir":     os.ModeSymlink,
		})
}

func TestFastWalk_SkipDir(t *testing.T) {
	testFastWalk(t, map[string]string{
		"foo/foo.go":   "one",
		"bar/bar.go":   "two",
		"skip/skip.go": "skip",
	},
		func(path string, typ os.FileMode) error {
			if typ == os.ModeDir && strings.HasSuffix(path, "skip") {
				return filepath.SkipDir
			}
			return nil
		},
		map[string]os.FileMode{
			"":                os.ModeDir,
			"/src":            os.ModeDir,
			"/src/bar":        os.ModeDir,
			"/src/bar/bar.go": 0,
			"/src/foo":        os.ModeDir,
			"/src/foo/foo.go": 0,
			"/src/skip":       os.ModeDir,
		})
}

func TestFastWalk_TraverseSymlink(t *testing.T) {
	switch runtime.GOOS {
	case "windows", "plan9":
		t.Skipf("skipping on %s", runtime.GOOS)
	}

	testFastWalk(t, map[string]string{
		"foo/foo.go":   "one",
		"bar/bar.go":   "two",
		"skip/skip.go": "skip",
		"symdir":       "LINK:foo",
	},
		func(path string, typ os.FileMode) error {
			if typ == os.ModeSymlink {
				return traverseLink
			}
			return nil
		},
		map[string]os.FileMode{
			"":                   os.ModeDir,
			"/src":               os.ModeDir,
			"/src/bar":           os.ModeDir,
			"/src/bar/bar.go":    0,
			"/src/foo":           os.ModeDir,
			"/src/foo/foo.go":    0,
			"/src/skip":          os.ModeDir,
			"/src/skip/skip.go":  0,
			"/src/symdir":        os.ModeSymlink,
			"/src/symdir/foo.go": 0,
		})
}

var benchDir = flag.String("benchdir", runtime.GOROOT(), "The directory to scan for BenchmarkFastWalk")

func BenchmarkFastWalk(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := Walk(*benchDir, func(path string, typ os.FileMode) error { return nil })
		if err != nil {
			b.Fatal(err)
		}
	}
}

// The below code is need testing. Copied from golang.org/x/tools/imports/fix.go

type pkg struct {
	dir             string // absolute file path to pkg directory ("/usr/lib/go/src/net/http")
	importPath      string // full pkg import path ("net/http", "foo/bar/vendor/a/b")
	importPathShort string // vendorless import path ("net/http", "a/b")
}

var testHookScanDir = func(dir string) {}

var scanGoRootDone = make(chan struct{}) // closed when scanGoRoot is done

var (
	testMu sync.RWMutex // guards globals reset by tests; used only if inTests
)

// Directory-scanning state.
var (
	// scanGoRootOnce guards calling scanGoRoot (for $GOROOT)
	scanGoRootOnce sync.Once
	// scanGoPathOnce guards calling scanGoPath (for $GOPATH)
	scanGoPathOnce sync.Once

	// populateIgnoreOnce guards calling populateIgnore
	populateIgnoreOnce sync.Once
	ignoredDirs        []os.FileInfo

	dirScanMu sync.RWMutex
	dirScan   map[string]*pkg // abs dir path => *pkg
)

func withEmptyGoPath(fn func()) {
	testMu.Lock()

	dirScanMu.Lock()
	populateIgnoreOnce = sync.Once{}
	scanGoRootOnce = sync.Once{}
	scanGoPathOnce = sync.Once{}
	dirScan = nil
	ignoredDirs = nil
	scanGoRootDone = make(chan struct{})
	dirScanMu.Unlock()

	oldGOPATH := build.Default.GOPATH
	oldGOROOT := build.Default.GOROOT
	build.Default.GOPATH = ""
	testHookScanDir = func(string) {}
	testMu.Unlock()

	defer func() {
		testMu.Lock()
		testHookScanDir = func(string) {}
		build.Default.GOPATH = oldGOPATH
		build.Default.GOROOT = oldGOROOT
		testMu.Unlock()
	}()

	fn()
}

type testConfig struct {
	// goroot and gopath optionally specifies the path on disk
	// to use for the GOROOT and GOPATH. If empty, a temp directory
	// is made if needed.
	goroot, gopath string

	// gorootFiles optionally specifies the complete contents of GOROOT to use,
	// If nil, the normal current $GOROOT is used.
	gorootFiles map[string]string // paths relative to $GOROOT/src to contents

	// gopathFiles is like gorootFiles, but for $GOPATH.
	// If nil, there is no GOPATH, though.
	gopathFiles map[string]string // paths relative to $GOPATH/src to contents
}

func mustTempDir(t *testing.T, prefix string) string {
	dir, err := ioutil.TempDir("", prefix)
	if err != nil {
		t.Fatal(err)
	}
	return dir
}

func mapToDir(destDir string, files map[string]string) error {
	for path, contents := range files {
		file := filepath.Join(destDir, "src", path)
		if err := os.MkdirAll(filepath.Dir(file), 0755); err != nil {
			return err
		}
		var err error
		if strings.HasPrefix(contents, "LINK:") {
			err = os.Symlink(strings.TrimPrefix(contents, "LINK:"), file)
		} else {
			err = ioutil.WriteFile(file, []byte(contents), 0644)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (c testConfig) test(t *testing.T, fn func(*goimportTest)) {
	goroot := c.goroot
	gopath := c.gopath

	if c.gorootFiles != nil && goroot == "" {
		goroot = mustTempDir(t, "goroot-")
		defer os.RemoveAll(goroot)
	}
	if err := mapToDir(goroot, c.gorootFiles); err != nil {
		t.Fatal(err)
	}

	if c.gopathFiles != nil && gopath == "" {
		gopath = mustTempDir(t, "gopath-")
		defer os.RemoveAll(gopath)
	}
	if err := mapToDir(gopath, c.gopathFiles); err != nil {
		t.Fatal(err)
	}

	withEmptyGoPath(func() {
		if goroot != "" {
			build.Default.GOROOT = goroot
		}
		build.Default.GOPATH = gopath

		it := &goimportTest{
			T:      t,
			goroot: build.Default.GOROOT,
			gopath: gopath,
			ctx:    &build.Default,
		}
		fn(it)
	})
}

type goimportTest struct {
	*testing.T
	ctx    *build.Context
	goroot string
	gopath string
}
