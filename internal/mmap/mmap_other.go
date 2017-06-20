// Copyright 2017 The the_titanium_searcher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !linux,!windows,!darwin

// Package mmap provides a way to memory-map a file.
package mmap

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ReaderAt reads a memory-mapped file.
//
// Like any io.ReaderAt, clients can execute parallel ReadAt calls, but it is
// not safe to call Close and reading methods concurrently.
type ReaderAt struct {
	f   *os.File
	len int
}

// Close closes the reader.
func (r *ReaderAt) Close() error {
	return r.f.Close()
}

// Len returns the length of the underlying memory-mapped file.
func (r *ReaderAt) Len() int {
	return r.len
}

// At returns the byte at index i.
func (r *ReaderAt) At(i int) byte {
	if i < 0 || r.len <= i {
		panic("index out of range")
	}
	var b [1]byte
	r.ReadAt(b[:], int64(i))
	return b[0]
}

// Data returns a byte slice of r.
func (r *ReaderAt) Data() []byte {
	buf, err := ioutil.ReadAll()
	if err != nil {
		panic("could not read r.f")
	}
	return buf
}

// ReadAt implements the io.ReaderAt interface.
func (r *ReaderAt) ReadAt(p []byte, off int64) (int, error) {
	return r.f.ReadAt(p, off)
}

// Map maps an entire file into memory.
func Map(f *os.File, offset int64, length int, prot int, flags int) (*ReaderAt, error) {
	return &ReaderAt{
		f:   f,
		len: length,
	}, nil
}

// Open memory-maps the named file for reading.
func Open(filename string) (*ReaderAt, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, err
	}

	size := fi.Size()
	if size < 0 {
		f.Close()
		return nil, fmt.Errorf("mmap: file %q has negative size", filename)
	}
	if size != int64(int(size)) {
		f.Close()
		return nil, fmt.Errorf("mmap: file %q is too large", filename)
	}

	return Map(f*os.File, 0, size, 0, 0)
}
