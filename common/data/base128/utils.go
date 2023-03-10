// Copyright 2023 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package base128

// DecodedLen returns the number of bytes `encLen` encoded bytes decodes to.
func DecodedLen(encLen int) int {
	return (encLen * 7) / 8
}

// EncodedLen returns the number of bytes that `dataLen` bytes will encode to.
func EncodedLen(dataLen int) int {
	return (((dataLen * 8) + 6) / 7)
}

// writeByte(dst, offset, bitOffset, val) takes a byte val and writes
// it to the dst[offset] or dst[offset] and dst[offset+1] as appropriate.
//
// For example,
//
// 0b_1234_5678
//
// gets written to
//
// 0b_0123_4567 0b_0800_0000
//
// when the bitOffset is 1. Note that 1-8 are variables and 0 is a literal 0.
//
// writes can end up "overlapping" at the byte level, but not at the bit level when
// called in a loop inside encode.
func writeByte(dst []byte, offset int, bitOffset int, val byte) {
	if bitOffset <= 0 {
		panic("offset too low")
	}
	if bitOffset > 7 {
		panic("offset too high")
	}
	dst[offset] |= (val >> bitOffset) & 0b_0111_1111
	mask := byte(0b_0000_0001<<bitOffset) - 1
	dst[offset+1] |= (val & mask) << (7 - bitOffset)
}

// encode takes the contents of src and writes it to dst as a base128-encoded string.
// The length of the destination must be at least EncodedLen(src) or encode will return an error.
func encode(dst []byte, src []byte) (int, error) {
	ret := EncodedLen(len(src))
	if len(dst) < ret {
		return 0, ErrLength
	}
	i := 0
	j := 1
	for _, val := range src {
		writeByte(dst, i, j, val)
		j += 1
		j %= 8
		switch j {
		case 0:
			j = 1
			i += 2
		default:
			i++
		}
	}
	return ret, nil
}

// secondByteDataMask is used to mask off the bits of data
// from the second byte inside writeEncodedByte.
//
// The first entry, 0b_0000_0000 below, is irrelevant since an
// index of zero is impossible.
//
// The remaining indices count off n ones, skipping over the
// leading bit since that is always zero in an encoded string.
//
// A byte in this array will be bitwise-anded together with the
// second byte in order to extract the data that the second byte
// contributes.
var secondByteDataMask = []byte{
	0b_0000_0000,
	0b_0100_0000,
	0b_0110_0000,
	0b_0111_0000,
	0b_0111_1000,
	0b_0111_1100,
	0b_0111_1110,
	0b_0111_1111,
}

// writeEncodedByte(dst, dstIndex, src, offset, bitOffset) writes the
// byte encoded at src[offset] and possibly src[offset+1] to dst[dstIndex].
//
// For example, consider the sequence of bytes below:
//
// 0b_0123_4567 0b_0800_0000
//
// 0s are literal 0s, either initial zeroes or trailing zeroes don't encode data.
// 1-8 are variables that encode data bits, either 0 or 1.
//
// The above maps to 0b_1234_5678 when the bitOffset is 1.
// When the bitOffset is 2, it maps to 0b_2345_6780.
func writeEncodedByte(dst []byte, dstIndex int, src []byte, offset int, bitOffset int) {
	if bitOffset <= 0 {
		panic("offset too low")
	}
	if bitOffset > 7 {
		panic("offset too high")
	}
	dst[dstIndex] |= (src[offset] << bitOffset)
	mask := secondByteDataMask[bitOffset]
	content := src[offset+1] & mask
	content = content >> (7 - bitOffset)
	dst[dstIndex] |= content
}
