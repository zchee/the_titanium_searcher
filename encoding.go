// Copyright 2017 The the_titanium_searcher Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This code based by:
//  https://github.com/monochromegane/the_platinum_searcher/blob/c7d8eec66dca50773e6b4ee7dfdad2174860b9b1/encoding.go
// The detect binary algorithm based by:
//  https://github.com/ggreer/the_silver_searcher/blob/99cf1834ce2c038c184d64793aaa6686381c49c5/src/util.c#L324-L377

package ti

type encodingType int

const (
	UNKNOWN encodingType = iota
	ERROR
	BINARY
	ASCII
	UTF8
	EUCJP
	SHIFTJIS
)

func encoding(bs []byte) encodingType {
	var (
		suspiciousBytes = 0
		likelyUtf8      = 0
		likelyEucjp     = 0
		likelyShiftjis  = 0
	)

	length := len(bs)

	if length == 0 {
		return ASCII
	}

	if length >= 3 && bs[0] == 0xEF && bs[1] == 0xBB && bs[2] == 0xBF {
		// UTF-8 BOM. This isn't binary.
		return UTF8
	}

	if length >= 5 && bs[0] == 0x25 && bs[1] == 0x50 && bs[2] == 0x44 && bs[3] == 0x46 && bs[4] == 0x2D {
		/*  %PDF-. This is binary. */
		return BINARY
	}

	for i := 0; i < length; i++ {
		if bs[i] == 0x00 {
			/* NULL char. It's binary */
			return BINARY
		} else if (bs[i] < 7 || bs[i] > 14) && (bs[i] < 32 || bs[i] > 127) {
			/* UTF-8 detection */
			if bs[i] > 193 && bs[i] < 224 && i+1 < length {
				i++
				if bs[i] > 127 && bs[i] < 192 {
					likelyUtf8++
					continue
				}

			} else if bs[i] > 223 && bs[i] < 240 && i+2 < length {
				i++
				if bs[i] > 127 && bs[i] < 192 && bs[i+1] > 127 && bs[i+1] < 192 {
					i++
					likelyUtf8++
					continue
				}
			}

			/* EUC-JP detection */
			if bs[i] == 142 && i+1 < length {
				i++
				if bs[i] > 160 && bs[i] < 224 {
					likelyEucjp++
					continue
				}
			} else if bs[i] > 160 && bs[i] < 255 && i+1 < length {
				i++
				if bs[i] > 160 && bs[i] < 255 {
					likelyEucjp++
					continue
				}
			}

			/* Shift-JIS detection */
			if bs[i] > 160 && bs[i] < 224 {
				likelyShiftjis++
				continue
			} else if ((bs[i] > 128 && bs[i] < 160) || (bs[i] > 223 && bs[i] < 240)) && i+1 < length {
				i++
				if (bs[i] > 63 && bs[i] < 127) || (bs[i] > 127 && bs[i] < 253) {
					likelyShiftjis++
					continue
				}
			}

			suspiciousBytes++
			if i >= 32 && (suspiciousBytes*100)/length > 10 {
				return BINARY
			}

		}
	}

	if (suspiciousBytes*100)/length > 10 {
		return BINARY
	}

	// Debugf("Detected points[utf8/eucjp/shiftjis] is %d/%d/%d.\n", likelyUtf8, likelyEucjp, likelyShiftjis)
	if likelyUtf8 == 0 && likelyEucjp == 0 && likelyShiftjis == 0 {
		return ASCII
	} else if likelyUtf8 >= likelyEucjp && likelyUtf8 >= likelyShiftjis {
		return UTF8
	} else if likelyEucjp >= likelyUtf8 && likelyEucjp >= likelyShiftjis {
		return EUCJP
	} else if likelyShiftjis >= likelyUtf8 && likelyShiftjis >= likelyEucjp {
		return SHIFTJIS
	}

	return ASCII

}
