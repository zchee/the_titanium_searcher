// Copyright 2017 The the_titanium_searcher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmap

import (
	"bufio"
	"os"
	"testing"
)

const (
	smallFilename = "mmap_test.go"
	// medFilename   = "testdata/Lint.cpp" // TODO(zchee): fetch medium size(about 50000) file
	// hugeFilename  = "testdata/huge.txt" // TODO(zchee): create huge.txt use `dd`
)

var (
	smallFilesize int64
	// medFilesize   int64
	// hugeFilesize  int64
)

func TestMain(m *testing.M) {
	// for Benchmark SetBytes.
	smallFilesize = readSize(smallFilename)
	// medFilesize = readSize(medFilename)
	// hugeFilesize = readSize(hugeFilename)

	os.Exit(m.Run())
}

// NOTE(zchee): If size is about 50000 or higher, mmap.Open is faster than read(2)
// We will use `1024<<6`(65536) for the switching to mmap(2) or read(2).
func benchmarkOpen(b *testing.B, filename string, size int64) {
	b.SetBytes(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Open(filename); err != nil {
			b.Fatal(err)
		}
	}
}

func benchmarkRead(b *testing.B, filename string, size int64) {
	b.SetBytes(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f, err := os.Open(filename)
		if err != nil {
			b.Fatal(err)
		}
		fi, err := f.Stat()
		if err != nil {
			b.Fatal(err)
		}
		buf := make([]byte, fi.Size())
		_, err = f.Read(buf)
		if err != nil {
			b.Fatal(err)
		}
		f.Close()
	}
}

func benchmarkBufioRead(b *testing.B, filename string, size int64) {
	b.SetBytes(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f, err := os.Open(filename)
		if err != nil {
			b.Fatal(err)
		}
		fi, err := f.Stat()
		if err != nil {
			b.Fatal(err)
		}
		reader := bufio.NewReaderSize(f, int(fi.Size()))
		buf := make([]byte, fi.Size())
		_, err = reader.Read(buf)
		if err != nil {
			b.Fatal(err)
		}
		f.Close()
	}
}

func BenchmarkSmallOpen(b *testing.B)       { benchmarkOpen(b, smallFilename, smallFilesize) }
func BenchmarkSmallRead(b *testing.B)       { benchmarkRead(b, smallFilename, smallFilesize) }
func BenchmarkSmallBufioRead(b *testing.B)  { benchmarkBufioRead(b, smallFilename, smallFilesize) }
// func BenchmarkMediumOpen(b *testing.B)      { benchmarkOpen(b, medFilename, medFilesize) }
// func BenchmarkMediumRead(b *testing.B)      { benchmarkRead(b, medFilename, medFilesize) }
// func BenchmarkMediumBufioRead(b *testing.B) { benchmarkBufioRead(b, medFilename, medFilesize) }
// func BenchmarkHugeOpen(b *testing.B)        { benchmarkOpen(b, hugeFilename, hugeFilesize) }
// func BenchmarkHugeRead(b *testing.B)        { benchmarkRead(b, hugeFilename, hugeFilesize) }
// func BenchmarkHugeBufioRead(b *testing.B)   { benchmarkBufioRead(b, hugeFilename, hugeFilesize) }
