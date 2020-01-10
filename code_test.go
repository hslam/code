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

func TestCodeSliceUint8(t *testing.T) {
	var buf = make([]byte, 64)
	var v []uint8 = []uint8{1}
	var n uint64
	n = EncodeSliceUint8(buf, v)
	var d = make([]uint8, 2)
	n = DecodeSliceUint8(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %d != %d", v[0], d[0])
	}
	if n != SizeofSliceUint8(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeSliceUint8(b *testing.B) {
	var buf = make([]byte, 64)
	var v []uint8 = []uint8{1}
	var n uint64
	n = EncodeSliceUint8(buf, v)
	var d = make([]uint8, 1)
	DecodeSliceUint8(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeSliceUint8(buf, v)
		DecodeSliceUint8(buf[:n], &d)
	}
}

func TestCodeSliceUint16(t *testing.T) {
	var buf = make([]byte, 64)
	var v []uint16 = []uint16{1}
	var n uint64
	n = EncodeSliceUint16(buf, v)
	var d = make([]uint16, 2)
	n = DecodeSliceUint16(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %d != %d", v[0], d[0])
	}
	if n != SizeofSliceUint16(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeSliceUint16(b *testing.B) {
	var buf = make([]byte, 64)
	var v []uint16 = []uint16{1}
	var n uint64
	n = EncodeSliceUint16(buf, v)
	var d = make([]uint16, 1)
	DecodeSliceUint16(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeSliceUint16(buf, v)
		DecodeSliceUint16(buf[:n], &d)
	}
}

func TestCodeSliceUint32(t *testing.T) {
	var buf = make([]byte, 64)
	var v []uint32 = []uint32{1}
	var n uint64
	n = EncodeSliceUint32(buf, v)
	var d = make([]uint32, 2)
	n = DecodeSliceUint32(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %d != %d", v[0], d[0])
	}
	if n != SizeofSliceUint32(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeSliceUint32(b *testing.B) {
	var buf = make([]byte, 64)
	var v []uint32 = []uint32{1}
	var n uint64
	n = EncodeSliceUint32(buf, v)
	var d = make([]uint32, 1)
	DecodeSliceUint32(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeSliceUint32(buf, v)
		DecodeSliceUint32(buf[:n], &d)
	}
}

func TestCodeSliceUint64(t *testing.T) {
	var buf = make([]byte, 64)
	var v []uint64 = []uint64{1}
	var n uint64
	n = EncodeSliceUint64(buf, v)
	var d = make([]uint64, 2)
	n = DecodeSliceUint64(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %d != %d", v[0], d[0])
	}
	if n != SizeofSliceUint64(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeSliceUint64(b *testing.B) {
	var buf = make([]byte, 64)
	var v []uint64 = []uint64{1}
	var n uint64
	n = EncodeSliceUint64(buf, v)
	var d = make([]uint64, 1)
	DecodeSliceUint64(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeSliceUint64(buf, v)
		DecodeSliceUint64(buf[:n], &d)
	}
}

func TestCodeSliceVarint(t *testing.T) {
	var buf = make([]byte, 64)
	var v []uint64 = []uint64{1}
	var n uint64
	n = EncodeSliceVarint(buf, v)
	var d = make([]uint64, 2)
	n = DecodeSliceVarint(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %d != %d", v[0], d[0])
	}
	if n != SizeofSliceVarint(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeSliceVarint(b *testing.B) {
	var buf = make([]byte, 64)
	var v []uint64 = []uint64{1}
	var n uint64
	n = EncodeSliceVarint(buf, v)
	var d = make([]uint64, 1)
	DecodeSliceVarint(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeSliceVarint(buf, v)
		DecodeSliceVarint(buf[:n], &d)
	}
}

func TestCodeSliceFloat32(t *testing.T) {
	var buf = make([]byte, 64)
	var v []float32 = []float32{3.14}
	var n uint64
	n = EncodeSliceFloat32(buf, v)
	var d = make([]float32, 2)
	n = DecodeSliceFloat32(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %.2f != %.2f", v[0], d[0])
	}
	if n != SizeofSliceFloat32(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeSliceFloat32(b *testing.B) {
	var buf = make([]byte, 64)
	var v []float32 = []float32{3.14}
	var n uint64
	n = EncodeSliceFloat32(buf, v)
	var d = make([]float32, 1)
	DecodeSliceFloat32(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeSliceFloat32(buf, v)
		DecodeSliceFloat32(buf[:n], &d)
	}
}

func TestCodeSliceFloat64(t *testing.T) {
	var buf = make([]byte, 64)
	var v []float64 = []float64{3.14}
	var n uint64
	n = EncodeSliceFloat64(buf, v)
	var d = make([]float64, 2)
	n = DecodeSliceFloat64(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %.2f != %.2f", v[0], d[0])
	}
	if n != SizeofSliceFloat64(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeSliceFloat64(b *testing.B) {
	var buf = make([]byte, 64)
	var v []float64 = []float64{3.14}
	var n uint64
	n = EncodeSliceFloat64(buf, v)
	var d = make([]float64, 1)
	DecodeSliceFloat64(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeSliceFloat64(buf, v)
		DecodeSliceFloat64(buf[:n], &d)
	}
}

func TestCodeSliceBool(t *testing.T) {
	var buf = make([]byte, 64)
	var v []bool = []bool{true, false}
	var n uint64
	n = EncodeSliceBool(buf, v)
	var d = make([]bool, 2)
	n = DecodeSliceBool(buf[:n], &d)
	if v[0] != d[0] {
		t.Errorf("error %t != %t", v[0], d[0])
	}
	if n != SizeofSliceBool(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeSliceBool(b *testing.B) {
	var buf = make([]byte, 64)
	var v []bool = []bool{true}
	var n uint64
	n = EncodeSliceBool(buf, v)
	var d = make([]bool, 1)
	DecodeSliceBool(buf[:n], &d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeSliceBool(buf, v)
		DecodeSliceBool(buf[:n], &d)
	}
}

func TestCodeSliceString(t *testing.T) {
	var buf = make([]byte, 64)
	var v []string = []string{"h", "w"}
	var n uint64
	n = EncodeSliceString(buf, v)
	var v2 []string = make([]string, 2)
	n = DecodeSliceString(buf[:n], &v2)
	if v[0][0] != v2[0][0] {
		t.Errorf("error %d != %d", v[0][0], v2[0][0])
	}
	if n != SizeofSliceString(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeSliceString(b *testing.B) {
	var buf = make([]byte, 64)
	var v []string = []string{"h", "w"}
	var n uint64
	n = EncodeSliceString(buf, v)
	var v2 []string = make([]string, 2)
	DecodeSliceString(buf[:n], &v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeSliceString(buf, v)
		DecodeSliceString(buf[:n], &v2)
	}
}

func TestCodeSliceBytes(t *testing.T) {
	var buf = make([]byte, 64)
	var v [][]byte = [][]byte{{1}, {2}}
	var n uint64
	n = EncodeSliceBytes(buf, v)
	var v2 [][]byte = make([][]byte, 2)
	n = DecodeSliceBytes(buf[:n], &v2)
	if v[0][0] != v2[0][0] {
		t.Errorf("error %d != %d", v[0][0], v2[0][0])
	}
	if n != SizeofSliceBytes(v) {
		t.Errorf("error %d != %d", n, len(buf[:n]))
	}
}

func BenchmarkCodeSliceBytes(b *testing.B) {
	var buf = make([]byte, 64)
	var v [][]byte = [][]byte{{1}, {2}}
	var n uint64
	n = EncodeSliceBytes(buf, v)
	var v2 [][]byte = make([][]byte, 2)
	DecodeSliceBytes(buf[:n], &v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = EncodeSliceBytes(buf, v)
		DecodeSliceBytes(buf[:n], &v2)
	}
}

//Sizeof
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

func BenchmarkSizeofSliceUint8(b *testing.B) {
	var v []uint8 = []uint8{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofSliceUint8(v)
	}
}

func BenchmarkSizeofSliceUint16(b *testing.B) {
	var v []uint16 = []uint16{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofSliceUint16(v)
	}
}

func BenchmarkSizeofSliceUint32(b *testing.B) {
	var v []uint32 = []uint32{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofSliceUint32(v)
	}
}

func BenchmarkSizeofSliceUint64(b *testing.B) {
	var v []uint64 = []uint64{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofSliceUint64(v)
	}
}

func BenchmarkSizeofSliceVarint(b *testing.B) {
	var v []uint64 = []uint64{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofSliceVarint(v)
	}
}

func BenchmarkSizeofSliceFloat32(b *testing.B) {
	var v []float32 = []float32{3.14}
	b.SetBytes(int64(4))
	for i := 0; i < b.N; i++ {
		SizeofSliceFloat32(v)
	}
}

func BenchmarkSizeofSliceFloat64(b *testing.B) {
	var v []float64 = []float64{3.14}
	b.SetBytes(int64(8))
	for i := 0; i < b.N; i++ {
		SizeofSliceFloat64(v)
	}
}

func BenchmarkSizeofSliceBool(b *testing.B) {
	var v []bool = []bool{true}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofSliceBool(v)
	}
}

func BenchmarkSizeofSliceString(b *testing.B) {
	var v []string = []string{"h", "w"}
	b.SetBytes(int64(2))
	for i := 0; i < b.N; i++ {
		SizeofSliceString(v)
	}
}

func BenchmarkSizeofSliceBytes(b *testing.B) {
	var v [][]byte = [][]byte{{1}, {2}}
	b.SetBytes(int64(2))
	for i := 0; i < b.N; i++ {
		SizeofSliceBytes(v)
	}
}
