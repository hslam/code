// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package code implements encoding and decoding.
package code

import (
	"encoding/binary"
	"testing"
)

func TestCheckBuffer(t *testing.T) {
	var buf = make([]byte, 512)
	n := uint64(256)
	b := CheckBuffer(buf, n)
	if uint64(len(b)) != n {
		t.Errorf("error %d != %d", len(b), n)
	}
}

func BenchmarkCheckBuffer(b *testing.B) {
	var buf = make([]byte, 512)
	n := uint64(256)
	b.SetBytes(int64(n))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckBuffer(buf, n)
	}
}
func TestCodeUint8(t *testing.T) {
	var buf = make([]byte, 1)
	var v uint8 = 128
	var n uint64
	n = EncodeUint8(buf, v)
	var d uint8
	n = DecodeUint8(buf[:n], &d)
	if v != d {
		t.Errorf("error %d != %d", v, d)
	}
	if n != 1 {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeUint8(b *testing.B) {
	var v uint8 = 128
	var n uint64
	var buf = make([]byte, 1)
	n = EncodeUint8(buf, v)
	var v2 uint8
	DecodeUint8(buf[:n], &v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeUint8(buf, v)
		DecodeUint8(buf[:n], &v2)
	}
}
func TestCodeUint16(t *testing.T) {
	var buf = make([]byte, 2)
	var v uint16 = 128
	var n uint64
	n = EncodeUint16(buf, v)
	var d uint16
	n = DecodeUint16(buf[:n], &d)
	if v != d {
		t.Errorf("error %d != %d", v, d)
	}
	if n != 2 {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}

}

func BenchmarkCodeUint16(b *testing.B) {
	var v uint16 = 128
	var n uint64
	var buf = make([]byte, 2)
	n = EncodeUint16(buf, v)
	var v2 uint16
	DecodeUint16(buf[:n], &v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeUint16(buf, v)
		DecodeUint16(buf[:n], &v2)
	}
}

func TestCodeUint32(t *testing.T) {
	var buf = make([]byte, 4)
	var v uint32 = 128
	var n uint64
	n = EncodeUint32(buf, v)
	var d uint32
	n = DecodeUint32(buf[:n], &d)
	if v != d {
		t.Errorf("error %d != %d", v, d)
	}
	if n != 4 {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}

}

func BenchmarkCodeUint32(b *testing.B) {
	var v uint32 = 128
	var n uint64
	var buf = make([]byte, 4)
	n = EncodeUint32(buf, v)
	var v2 uint32
	DecodeUint32(buf[:n], &v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeUint32(buf, v)
		DecodeUint32(buf[:n], &v2)
	}
}

func TestCodeUint64(t *testing.T) {
	var buf = make([]byte, 8)
	var v uint64 = 128
	var n uint64
	n = EncodeUint64(buf, v)
	var d uint64
	n = DecodeUint64(buf[:n], &d)
	if v != d {
		t.Errorf("error %d != %d", v, d)
	}
	if n != 8 {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeUint64(b *testing.B) {
	var v uint64 = 128
	var n uint64
	var buf = make([]byte, 8)
	n = EncodeUint64(buf, v)
	var v2 uint64
	DecodeUint64(buf[:n], &v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeUint64(buf, v)
		DecodeUint64(buf[:n], &v2)
	}
}

func TestCodeVarint(t *testing.T) {
	var buf = make([]byte, 10)
	var v uint64 = 128
	var n uint64
	n = EncodeVarint(buf, v)
	var d uint64
	n = DecodeVarint(buf[:n], &d)
	if v != d {
		t.Errorf("error %d != %d", v, d)
	}
	if n != SizeofVarint(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeVarint(b *testing.B) {
	var v uint64 = 128
	var n uint64
	var buf = make([]byte, 10)
	n = EncodeVarint(buf, v)
	var v2 uint64
	DecodeVarint(buf[:n], &v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeVarint(buf, v)
		DecodeVarint(buf[:n], &v2)
	}
}
func TestBinaryVarint(t *testing.T) {
	var buf = make([]byte, 10)
	var v uint64 = 128
	var n int
	n = binary.PutUvarint(buf, v)
	var d uint64
	var n2 int
	d, n2 = binary.Uvarint(buf[:n])
	if v != d {
		t.Errorf("error %d != %d", v, d)
	}
	if n != n2 {
		t.Errorf("error %d != %d", n, n2)
	}
}
func BenchmarkBinaryVarint(b *testing.B) {
	var v uint64 = 128
	var n int
	var buf = make([]byte, 10)
	n = binary.PutUvarint(buf, v)
	binary.Uvarint(buf[:n])
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = binary.PutUvarint(buf, v)
		binary.Uvarint(buf[:n])
	}
}
func TestCodeFloat32(t *testing.T) {
	var f float32 = 3.14
	var n uint64
	var buf = make([]byte, 4)
	n = EncodeFloat32(buf, f)
	var v float32
	n = DecodeFloat32(buf[:n], &v)
	if v != f {
		t.Errorf("error %.7f != %.7f", v, f)
	}
	if n != SizeofFloat32(f) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeFloat32(b *testing.B) {
	var f float32 = 3.1415926
	var n uint64
	var buf = make([]byte, 4)
	n = EncodeFloat32(buf, f)
	var v float32
	DecodeFloat32(buf[:n], &v)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeFloat32(buf, f)
		DecodeFloat32(buf[:n], &v)
	}
}
func TestCodeFloat64(t *testing.T) {
	var f float64 = 3.1415926
	var n uint64
	var buf = make([]byte, 8)
	n = EncodeFloat64(buf, f)
	var v float64
	n = DecodeFloat64(buf[:n], &v)
	if v != f {
		t.Errorf("error %.7f != %.7f", v, f)
	}
	if n != SizeofFloat64(f) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeFloat64(b *testing.B) {
	var f float64 = 3.1415926
	var n uint64
	var buf = make([]byte, 8)
	n = EncodeFloat64(buf, f)
	var v float64
	DecodeFloat64(buf[:n], &v)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeFloat64(buf, f)
		DecodeFloat64(buf[:n], &v)
	}
}

func TestCodeBool(t *testing.T) {
	var buf = make([]byte, 9)
	var v bool = true
	var n uint64
	n = EncodeBool(buf, v)
	var d bool
	n = DecodeBool(buf[:n], &d)
	if v != d {
		t.Errorf("error %t != %t", v, d)
	}
	if n != SizeofBool(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
	v = false
	n = EncodeBool(buf, v)
	n = DecodeBool(buf[:n], &d)
	if v != d {
		t.Errorf("error %t != %t", v, d)
	}
	if n != SizeofBool(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeBool(b *testing.B) {
	var buf = make([]byte, 1)
	var v bool = true
	var n uint64
	n = EncodeBool(buf, v)
	var v2 bool
	DecodeBool(buf[:n], &v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeBool(buf, v)
		DecodeBool(buf[:n], &v2)
	}
}

func TestCodeString(t *testing.T) {
	var buf = make([]byte, 64)
	var v string = "HelloWorld"
	var n uint64
	n = EncodeString(buf, v)
	var d string
	n = DecodeString(buf[:n], &d)
	if v != d {
		t.Errorf("error %s != %s", v, d)
	}
	if n != SizeofString(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeString(b *testing.B) {
	var buf = make([]byte, 64)
	var v string = "h"
	var n uint64
	n = EncodeString(buf, v)
	var v2 string
	DecodeString(buf[:n], &v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeString(buf, v)
		DecodeString(buf[:n], &v2)
	}
}

func TestCodeBytes(t *testing.T) {
	var buf = make([]byte, 64)
	var v []byte = []byte{1}
	var n uint64
	n = EncodeBytes(buf, v)
	var d = make([]byte, 2)
	n = DecodeBytes(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %d != %d", v[0], d[0])
	}
	if n != SizeofBytes(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeBytes(b *testing.B) {
	var buf = make([]byte, 64)
	var v []byte = []byte{1}
	var n uint64
	n = EncodeBytes(buf, v)
	var d = make([]byte, 1)
	DecodeBytes(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeBytes(buf, v)
		DecodeBytes(buf[:n], &d)
	}
}

func TestCodeUint8Slice(t *testing.T) {
	var buf = make([]byte, 64)
	var v []uint8 = []uint8{1}
	var n uint64
	n = EncodeUint8Slice(buf, v)
	var d = make([]uint8, 2)
	n = DecodeUint8Slice(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %d != %d", v[0], d[0])
	}
	if n != SizeofUint8Slice(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeUint8Slice(b *testing.B) {
	var buf = make([]byte, 64)
	var v []uint8 = []uint8{1}
	var n uint64
	n = EncodeUint8Slice(buf, v)
	var d = make([]uint8, 1)
	DecodeUint8Slice(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeUint8Slice(buf, v)
		DecodeUint8Slice(buf[:n], &d)
	}
}

func TestCodeUint16Slice(t *testing.T) {
	var buf = make([]byte, 64)
	var v []uint16 = []uint16{1}
	var n uint64
	n = EncodeUint16Slice(buf, v)
	var d = make([]uint16, 2)
	n = DecodeUint16Slice(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %d != %d", v[0], d[0])
	}
	if n != SizeofUint16Slice(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeUint16Slice(b *testing.B) {
	var buf = make([]byte, 64)
	var v []uint16 = []uint16{1}
	var n uint64
	n = EncodeUint16Slice(buf, v)
	var d = make([]uint16, 1)
	DecodeUint16Slice(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeUint16Slice(buf, v)
		DecodeUint16Slice(buf[:n], &d)
	}
}

func TestCodeUint32Slice(t *testing.T) {
	var buf = make([]byte, 64)
	var v []uint32 = []uint32{1}
	var n uint64
	n = EncodeUint32Slice(buf, v)
	var d = make([]uint32, 2)
	n = DecodeUint32Slice(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %d != %d", v[0], d[0])
	}
	if n != SizeofUint32Slice(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeUint32Slice(b *testing.B) {
	var buf = make([]byte, 64)
	var v []uint32 = []uint32{1}
	var n uint64
	n = EncodeUint32Slice(buf, v)
	var d = make([]uint32, 1)
	DecodeUint32Slice(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeUint32Slice(buf, v)
		DecodeUint32Slice(buf[:n], &d)
	}
}

func TestCodeUint64Slice(t *testing.T) {
	var buf = make([]byte, 64)
	var v []uint64 = []uint64{1}
	var n uint64
	n = EncodeUint64Slice(buf, v)
	var d = make([]uint64, 2)
	n = DecodeUint64Slice(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %d != %d", v[0], d[0])
	}
	if n != SizeofUint64Slice(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeUint64Slice(b *testing.B) {
	var buf = make([]byte, 64)
	var v []uint64 = []uint64{1}
	var n uint64
	n = EncodeUint64Slice(buf, v)
	var d = make([]uint64, 1)
	DecodeUint64Slice(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeUint64Slice(buf, v)
		DecodeUint64Slice(buf[:n], &d)
	}
}

func TestCodeVarintSlice(t *testing.T) {
	var buf = make([]byte, 64)
	var v []uint64 = []uint64{1}
	var n uint64
	n = EncodeVarintSlice(buf, v)
	var d = make([]uint64, 2)
	n = DecodeVarintSlice(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %d != %d", v[0], d[0])
	}
	if n != SizeofVarintSlice(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeVarintSlice(b *testing.B) {
	var buf = make([]byte, 64)
	var v []uint64 = []uint64{1}
	var n uint64
	n = EncodeVarintSlice(buf, v)
	var d = make([]uint64, 1)
	DecodeVarintSlice(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeVarintSlice(buf, v)
		DecodeVarintSlice(buf[:n], &d)
	}
}

func TestCodeFloat32Slice(t *testing.T) {
	var buf = make([]byte, 64)
	var v []float32 = []float32{3.14}
	var n uint64
	n = EncodeFloat32Slice(buf, v)
	var d = make([]float32, 2)
	n = DecodeFloat32Slice(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %.2f != %.2f", v[0], d[0])
	}
	if n != SizeofFloat32Slice(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeFloat32Slice(b *testing.B) {
	var buf = make([]byte, 64)
	var v []float32 = []float32{3.14}
	var n uint64
	n = EncodeFloat32Slice(buf, v)
	var d = make([]float32, 1)
	DecodeFloat32Slice(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeFloat32Slice(buf, v)
		DecodeFloat32Slice(buf[:n], &d)
	}
}

func TestCodeFloat64Slice(t *testing.T) {
	var buf = make([]byte, 64)
	var v []float64 = []float64{3.14}
	var n uint64
	n = EncodeFloat64Slice(buf, v)
	var d = make([]float64, 2)
	n = DecodeFloat64Slice(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %.2f != %.2f", v[0], d[0])
	}
	if n != SizeofFloat64Slice(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeFloat64Slice(b *testing.B) {
	var buf = make([]byte, 64)
	var v []float64 = []float64{3.14}
	var n uint64
	n = EncodeFloat64Slice(buf, v)
	var d = make([]float64, 1)
	DecodeFloat64Slice(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeFloat64Slice(buf, v)
		DecodeFloat64Slice(buf[:n], &d)
	}
}

func TestCodeBoolSlice(t *testing.T) {
	var buf = make([]byte, 64)
	var v []bool = []bool{true, false}
	var n uint64
	n = EncodeBoolSlice(buf, v)
	var d = make([]bool, 2)
	n = DecodeBoolSlice(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %t != %t", v[0], d[0])
	}
	if n != SizeofBoolSlice(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeBoolSlice(b *testing.B) {
	var buf = make([]byte, 64)
	var v []bool = []bool{true}
	var n uint64
	n = EncodeBoolSlice(buf, v)
	var d = make([]bool, 1)
	DecodeBoolSlice(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeBoolSlice(buf, v)
		DecodeBoolSlice(buf[:n], &d)
	}
}

func TestCodeStringSlice(t *testing.T) {
	var buf = make([]byte, 64)
	var v []string = []string{"h", "w"}
	var n uint64
	n = EncodeStringSlice(buf, v)
	var v2 []string = make([]string, 2)
	n = DecodeStringSlice(buf[:n], &v2)
	if v[0][0] != v2[0][0] {
		t.Errorf("error %d != %d", v[0][0], v2[0][0])
	}
	if n != SizeofStringSlice(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeStringSlice(b *testing.B) {
	var buf = make([]byte, 64)
	var v []string = []string{"h", "w"}
	var n uint64
	n = EncodeStringSlice(buf, v)
	var v2 []string = make([]string, 2)
	DecodeStringSlice(buf[:n], &v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeStringSlice(buf, v)
		DecodeStringSlice(buf[:n], &v2)
	}
}

func TestCodeBytesSlice(t *testing.T) {
	var buf = make([]byte, 64)
	var v [][]byte = [][]byte{{1}, {2}}
	var n uint64
	n = EncodeBytesSlice(buf, v)
	var v2 [][]byte = make([][]byte, 2)
	n = DecodeBytesSlice(buf[:n], &v2)
	if v[0][0] != v2[0][0] {
		t.Errorf("error %d != %d", v[0][0], v2[0][0])
	}
	if n != SizeofBytesSlice(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeBytesSlice(b *testing.B) {
	var buf = make([]byte, 64)
	var v [][]byte = [][]byte{{1}, {2}}
	var n uint64
	n = EncodeBytesSlice(buf, v)
	var v2 [][]byte = make([][]byte, 2)
	DecodeBytesSlice(buf[:n], &v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeBytesSlice(buf, v)
		DecodeBytesSlice(buf[:n], &v2)
	}
}

func BenchmarkSizeofUint8(b *testing.B) {
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofUint8(1)
	}
}

func BenchmarkSizeofUint16(b *testing.B) {
	b.SetBytes(int64(2))
	for i := 0; i < b.N; i++ {
		SizeofUint16(1)
	}
}

func BenchmarkSizeofUint32(b *testing.B) {
	b.SetBytes(int64(4))
	for i := 0; i < b.N; i++ {
		SizeofUint32(1)
	}
}

func BenchmarkSizeofUint64(b *testing.B) {
	b.SetBytes(int64(8))
	for i := 0; i < b.N; i++ {
		SizeofUint64(1)
	}
}

func BenchmarkSizeofVarint(b *testing.B) {
	b.SetBytes(int64(2))
	var v uint64 = 128
	for i := 0; i < b.N; i++ {
		SizeofVarint(v)
	}
}

func BenchmarkSizeofFloat32(b *testing.B) {
	b.SetBytes(int64(4))
	for i := 0; i < b.N; i++ {
		SizeofFloat32(3.14)
	}
}

func BenchmarkSizeofFloat64(b *testing.B) {
	b.SetBytes(int64(8))
	for i := 0; i < b.N; i++ {
		SizeofFloat64(3.14)
	}
}

func BenchmarkSizeofBool(b *testing.B) {
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofBool(true)
	}
}

func BenchmarkSizeofString(b *testing.B) {
	var v string = "h"
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofString(v)
	}
}

func BenchmarkSizeofBytes(b *testing.B) {
	var v []byte = []byte{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofBytes(v)
	}
}

func BenchmarkSizeofUint8Slice(b *testing.B) {
	var v []uint8 = []uint8{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofUint8Slice(v)
	}
}

func BenchmarkSizeofUint16Slice(b *testing.B) {
	var v []uint16 = []uint16{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofUint16Slice(v)
	}
}

func BenchmarkSizeofUint32Slice(b *testing.B) {
	var v []uint32 = []uint32{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofUint32Slice(v)
	}
}

func BenchmarkSizeofUint64Slice(b *testing.B) {
	var v []uint64 = []uint64{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofUint64Slice(v)
	}
}

func BenchmarkSizeofVarintSlice(b *testing.B) {
	var v []uint64 = []uint64{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofVarintSlice(v)
	}
}

func BenchmarkSizeofFloat32Slice(b *testing.B) {
	var v []float32 = []float32{3.14}
	b.SetBytes(int64(4))
	for i := 0; i < b.N; i++ {
		SizeofFloat32Slice(v)
	}
}

func BenchmarkSizeofFloat64Slice(b *testing.B) {
	var v []float64 = []float64{3.14}
	b.SetBytes(int64(8))
	for i := 0; i < b.N; i++ {
		SizeofFloat64Slice(v)
	}
}

func BenchmarkSizeofBoolSlice(b *testing.B) {
	var v []bool = []bool{true}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofBoolSlice(v)
	}
}

func BenchmarkSizeofStringSlice(b *testing.B) {
	var v []string = []string{"h", "w"}
	b.SetBytes(int64(2))
	for i := 0; i < b.N; i++ {
		SizeofStringSlice(v)
	}
}

func BenchmarkSizeofBytesSlice(b *testing.B) {
	var v [][]byte = [][]byte{{1}, {2}}
	b.SetBytes(int64(2))
	for i := 0; i < b.N; i++ {
		SizeofBytesSlice(v)
	}
}

func BenchmarkMaxUint8Bytes(b *testing.B) {
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		MaxUint8Bytes(1)
	}
}

func BenchmarkMaxUint16Bytes(b *testing.B) {
	b.SetBytes(int64(2))
	for i := 0; i < b.N; i++ {
		MaxUint16Bytes(1)
	}
}

func BenchmarkMaxUint32Bytes(b *testing.B) {
	b.SetBytes(int64(4))
	for i := 0; i < b.N; i++ {
		MaxUint32Bytes(1)
	}
}

func BenchmarkMaxUint64Bytes(b *testing.B) {
	b.SetBytes(int64(8))
	for i := 0; i < b.N; i++ {
		MaxUint64Bytes(1)
	}
}

func BenchmarkMaxVarintBytes(b *testing.B) {
	b.SetBytes(int64(2))
	var v uint64 = 128
	for i := 0; i < b.N; i++ {
		MaxVarintBytes(v)
	}
}

func BenchmarkMaxFloat32Bytes(b *testing.B) {
	b.SetBytes(int64(4))
	for i := 0; i < b.N; i++ {
		MaxFloat32Bytes(3.14)
	}
}

func BenchmarkMaxFloat64Bytes(b *testing.B) {
	b.SetBytes(int64(8))
	for i := 0; i < b.N; i++ {
		MaxFloat64Bytes(3.14)
	}
}

func BenchmarkMaxBoolBytes(b *testing.B) {
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		MaxBoolBytes(true)
	}
}

func BenchmarkMaxStringBytes(b *testing.B) {
	var v string = "h"
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		MaxStringBytes(v)
	}
}

func BenchmarkMaxBytesBytes(b *testing.B) {
	var v []byte = []byte{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		MaxBytesBytes(v)
	}
}

func BenchmarkMaxUint8SliceBytes(b *testing.B) {
	var v []uint8 = []uint8{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		MaxUint8SliceBytes(v)
	}
}

func BenchmarkMaxUint16SliceBytes(b *testing.B) {
	var v []uint16 = []uint16{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		MaxUint16SliceBytes(v)
	}
}

func BenchmarkMaxUint32SliceBytes(b *testing.B) {
	var v []uint32 = []uint32{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		MaxUint32SliceBytes(v)
	}
}

func BenchmarkMaxUint64SliceBytes(b *testing.B) {
	var v []uint64 = []uint64{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		MaxUint64SliceBytes(v)
	}
}

func BenchmarkMaxVarintSliceBytes(b *testing.B) {
	var v []uint64 = []uint64{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		MaxVarintSliceBytes(v)
	}
}

func BenchmarkMaxFloat32SliceBytes(b *testing.B) {
	var v []float32 = []float32{3.14}
	b.SetBytes(int64(4))
	for i := 0; i < b.N; i++ {
		MaxFloat32SliceBytes(v)
	}
}

func BenchmarkMaxFloat64SliceBytes(b *testing.B) {
	var v []float64 = []float64{3.14}
	b.SetBytes(int64(8))
	for i := 0; i < b.N; i++ {
		MaxFloat64SliceBytes(v)
	}
}

func BenchmarkMaxBoolSliceBytes(b *testing.B) {
	var v []bool = []bool{true}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		MaxBoolSliceBytes(v)
	}
}

func BenchmarkMaxStringSliceBytes(b *testing.B) {
	var v []string = []string{"h", "w"}
	b.SetBytes(int64(2))
	for i := 0; i < b.N; i++ {
		MaxStringSliceBytes(v)
	}
}

func BenchmarkMaxBytesSliceBytes(b *testing.B) {
	var v [][]byte = [][]byte{{1}, {2}}
	b.SetBytes(int64(2))
	for i := 0; i < b.N; i++ {
		MaxBytesSliceBytes(v)
	}
}
