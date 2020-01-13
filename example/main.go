package main

import (
	"fmt"
	"github.com/hslam/code"
)

func main() {
	Uint8()
	Uint16()
	Uint32()
	Uint64()
	Varint()
	Float32()
	Float64()
	Bool()
	String()
	Bytes()
	Uint8Slice()
	Uint16Slice()
	Uint32Slice()
	Uint64Slice()
	VarintSlice()
	Float32Slice()
	Float64Slice()
	BoolSlice()
	StringSlice()
	BytesSlice()
}

//Uint8 Example
func Uint8() {
	var i uint8 = 128
	var buf = make([]byte, code.MaxUint8Bytes(i))
	var n uint64
	size := code.SizeofUint8(i)
	fmt.Printf("SizeofUint8:%d sizeof:%d\n", i, size)
	n = code.EncodeUint8(buf, i)
	fmt.Printf("EncodeUint8:%d to []byte:%v\n", i, buf[:n])
	var v uint8
	n = code.DecodeUint8(buf[:n], &v)
	fmt.Printf("DecodeUint8:%d,length:%d\n", v, n)
}

//Uint16 Example
func Uint16() {
	var i uint16 = 128
	var buf = make([]byte, code.MaxUint16Bytes(i))
	var n uint64
	size := code.SizeofUint16(i)
	fmt.Printf("SizeofUint16:%d sizeof:%d\n", i, size)
	n = code.EncodeUint16(buf, i)
	fmt.Printf("EncodeUint16:%d to []byte:%v\n", i, buf[:n])
	var v uint16
	n = code.DecodeUint16(buf[:n], &v)
	fmt.Printf("DecodeUint16:%d,length:%d\n", v, n)
}

//Uint32 Example
func Uint32() {
	var i uint32 = 128
	var buf = make([]byte, code.MaxUint32Bytes(i))
	var n uint64
	size := code.SizeofUint32(i)
	fmt.Printf("SizeofUint32:%d sizeof:%d\n", i, size)
	n = code.EncodeUint32(buf, i)
	fmt.Printf("EncodeUint32:%d to []byte:%v\n", i, buf[:n])
	var v uint32
	n = code.DecodeUint32(buf[:n], &v)
	fmt.Printf("DecodeUint32:%d,length:%d\n", v, n)
}

//Uint64 Example
func Uint64() {
	var i uint64 = 128
	var buf = make([]byte, code.MaxUint64Bytes(i))
	var n uint64
	size := code.SizeofUint64(i)
	fmt.Printf("SizeofUint64:%d sizeof:%d\n", i, size)
	n = code.EncodeUint64(buf, i)
	fmt.Printf("EncodeUint64:%d to []byte:%v\n", i, buf[:n])
	var v uint64
	n = code.DecodeUint64(buf[:n], &v)
	fmt.Printf("DecodeUint64:%d,length:%d\n", v, n)
}

//Varint Example
func Varint() {
	var i uint64 = 128
	var buf = make([]byte, code.MaxVarintBytes(i))
	var n uint64
	size := code.SizeofVarint(i)
	fmt.Printf("SizeofVarint:%d sizeof:%d\n", i, size)
	n = code.EncodeVarint(buf, i)
	fmt.Printf("EncodeVarint:%d to []byte:%v\n", i, buf[:n])
	var v uint64
	n = code.DecodeVarint(buf[:n], &v)
	fmt.Printf("DecodeVarint:%d,length:%d\n", v, n)
}

//Float32 Example
func Float32() {
	var i float32 = 3.14
	var buf = make([]byte, code.MaxFloat32Bytes(i))
	var n uint64
	size := code.SizeofFloat32(i)
	fmt.Printf("SizeofFloat32:%.2f sizeof:%d\n", i, size)
	n = code.EncodeFloat32(buf, i)
	fmt.Printf("EncodeFloat32:%.2f to []byte:%v\n", i, buf[:n])
	var v float32
	n = code.DecodeFloat32(buf[:n], &v)
	fmt.Printf("EncodeFloat32:%.2f,length:%d\n", v, n)
}

//Float64 Example
func Float64() {
	var i float64 = 3.1415926
	var buf = make([]byte, code.MaxFloat64Bytes(i))
	var n uint64
	size := code.SizeofFloat64(i)
	fmt.Printf("SizeofFloat64:%.7f sizeof:%d\n", i, size)
	n = code.EncodeFloat64(buf, i)
	fmt.Printf("EncodeFloat64:%.7f to []byte:%v\n", i, buf[:n])
	var v float64
	n = code.DecodeFloat64(buf[:n], &v)
	fmt.Printf("DecodeFloat64:%.7f,length:%d\n", v, n)
}

//Bool Example
func Bool() {
	var i bool = true
	var buf = make([]byte, code.MaxBoolBytes(i))
	var n uint64
	size := code.SizeofBool(i)
	fmt.Printf("SizeofBool:%t sizeof:%d\n", i, size)
	n = code.EncodeBool(buf, i)
	fmt.Printf("EncodeBool:%t to []byte:%v\n", i, buf[:n])
	var v bool
	n = code.DecodeBool(buf[:n], &v)
	fmt.Printf("DecodeBool:%t,length:%d\n", v, n)
}

//String Example
func String() {
	var i string = "Hello"
	var buf = make([]byte, code.MaxStringBytes(i))
	var n uint64
	size := code.SizeofString(i)
	fmt.Printf("SizeofString:%s sizeof:%d\n", i, size)
	n = code.EncodeString(buf, i)
	fmt.Printf("EncodeString:%s to []byte:%v\n", i, buf[:n])
	var v string
	n = code.DecodeString(buf[:n], &v)
	fmt.Printf("DecodeString:%s,length:%d\n", v, n)
}

//Bytes Example
func Bytes() {
	var i []byte = []byte{1, 2}
	var buf = make([]byte, code.MaxBytesBytes(i))
	var n uint64
	size := code.SizeofBytes(i)
	fmt.Printf("SizeofBytes:%v sizeof:%d\n", i, size)
	n = code.EncodeBytes(buf, i)
	fmt.Printf("EncodeBytes:%v to []byte:%v\n", i, buf[:n])
	var v = make([]byte, 2)
	n = code.DecodeBytes(buf[:n], &v)
	fmt.Printf("DecodeBytes:%v,length:%d\n", v, n)
}

//Uint8Slice Example
func Uint8Slice() {
	var i []uint8 = []uint8{128, 255}
	var buf = make([]byte, code.MaxUint8SliceBytes(i))
	var n uint64
	size := code.SizeofUint8Slice(i)
	fmt.Printf("SizeofUint8Slice:%v sizeof:%d\n", i, size)
	n = code.EncodeUint8Slice(buf, i)
	fmt.Printf("EncodeUint8Slice:%v to []byte:%v\n", i, buf[:n])
	var v = make([]uint8, 2)
	n = code.DecodeUint8Slice(buf[:n], &v)
	fmt.Printf("DecodeUint8Slice:%v,length:%d\n", v, n)
}

//Uint16Slice Example
func Uint16Slice() {
	var i []uint16 = []uint16{128, 256}
	var buf = make([]byte, code.MaxUint16SliceBytes(i))
	var n uint64
	size := code.SizeofUint16Slice(i)
	fmt.Printf("SizeofUint16Slice:%v sizeof:%d\n", i, size)
	n = code.EncodeUint16Slice(buf, i)
	fmt.Printf("EncodeUint16Slice:%v to []byte:%v\n", i, buf[:n])
	var v = make([]uint16, 2)
	n = code.DecodeUint16Slice(buf[:n], &v)
	fmt.Printf("DecodeUint16Slice:%v,length:%d\n", v, n)
}

//Uint32Slice Example
func Uint32Slice() {
	var i []uint32 = []uint32{128, 256}
	var buf = make([]byte, code.MaxUint32SliceBytes(i))
	var n uint64
	size := code.SizeofUint32Slice(i)
	fmt.Printf("SizeofUint32Slice:%v sizeof:%d\n", i, size)
	n = code.EncodeUint32Slice(buf, i)
	fmt.Printf("EncodeUint32Slice:%v to []byte:%v\n", i, buf[:n])
	var v = make([]uint32, 2)
	n = code.DecodeUint32Slice(buf[:n], &v)
	fmt.Printf("DecodeUint32Slice:%v,length:%d\n", v, n)
}

//Uint64Slice Example
func Uint64Slice() {
	var i []uint64 = []uint64{128, 256}
	var buf = make([]byte, code.MaxUint64SliceBytes(i))
	var n uint64
	size := code.SizeofUint64Slice(i)
	fmt.Printf("SizeofUint64Slice:%v sizeof:%d\n", i, size)
	n = code.EncodeUint64Slice(buf, i)
	fmt.Printf("EncodeUint64Slice:%v to []byte:%v\n", i, buf[:n])
	var v = make([]uint64, 2)
	n = code.DecodeUint64Slice(buf[:n], &v)
	fmt.Printf("DecodeUint64Slice:%v,length:%d\n", v, n)
}

//VarintSlice Example
func VarintSlice() {
	var i []uint64 = []uint64{128, 256}
	var buf = make([]byte, code.MaxVarintSliceBytes(i))
	var n uint64
	size := code.SizeofVarintSlice(i)
	fmt.Printf("SizeofVarintSlice:%v sizeof:%d\n", i, size)
	n = code.EncodeVarintSlice(buf, i)
	fmt.Printf("EncodeVarintSlice:%v to []byte:%v\n", i, buf[:n])
	var v = make([]uint64, 2)
	n = code.DecodeVarintSlice(buf[:n], &v)
	fmt.Printf("DecodeVarintSlice:%v,length:%d\n", v, n)
}

//Float32Slice Example
func Float32Slice() {
	var i []float32 = []float32{3.14}
	var buf = make([]byte, code.MaxFloat32SliceBytes(i))
	var n uint64
	size := code.SizeofFloat32Slice(i)
	fmt.Printf("SizeofFloat32Slice:%v sizeof:%d\n", i, size)
	n = code.EncodeFloat32Slice(buf, i)
	fmt.Printf("EncodeFloat32Slice:%v to []byte:%v\n", i, buf[:n])
	var v = make([]float32, 2)
	n = code.DecodeFloat32Slice(buf[:n], &v)
	fmt.Printf("DecodeFloat32Slice:%v,length:%d\n", v, n)
}

//Float64Slice Example
func Float64Slice() {
	var i []float64 = []float64{3.1415926}
	var buf = make([]byte, code.MaxFloat64SliceBytes(i))
	var n uint64
	size := code.SizeofFloat64Slice(i)
	fmt.Printf("SizeofFloat64Slice:%v sizeof:%d\n", i, size)
	n = code.EncodeFloat64Slice(buf, i)
	fmt.Printf("EncodeFloat64Slice:%v to []byte:%v\n", i, buf[:n])
	var v = make([]float64, 2)
	n = code.DecodeFloat64Slice(buf[:n], &v)
	fmt.Printf("DecodeFloat64Slice:%v,length:%d\n", v, n)
}

//BoolSlice Example
func BoolSlice() {
	var i []bool = []bool{true, false}
	var buf = make([]byte, code.MaxBoolSliceBytes(i))
	var n uint64
	size := code.SizeofBoolSlice(i)
	fmt.Printf("SizeofBoolSlice:%v sizeof:%d\n", i, size)
	n = code.EncodeBoolSlice(buf, i)
	fmt.Printf("EncodeBoolSlice:%v to []byte:%v\n", i, buf[:n])
	var v = make([]bool, 2)
	n = code.DecodeBoolSlice(buf[:n], &v)
	fmt.Printf("DecodeBoolSlice:%v,length:%d\n", v, n)
}

//StringSlice Example
func StringSlice() {
	var i []string = []string{"Hello", "World"}
	var buf = make([]byte, code.MaxStringSliceBytes(i))
	var n uint64
	size := code.SizeofStringSlice(i)
	fmt.Printf("SizeofStringSlice:%v sizeof:%d\n", i, size)
	n = code.EncodeStringSlice(buf, i)
	fmt.Printf("EncodeStringSlice:%v to []byte:%v\n", i, buf[:n])
	var v = make([]string, 2)
	n = code.DecodeStringSlice(buf[:n], &v)
	fmt.Printf("DecodeStringSlice:%v,length:%d\n", v, n)
}

//BytesSlice Example
func BytesSlice() {
	var i [][]byte = [][]byte{{1, 2}, {3}}
	var buf = make([]byte, code.MaxBytesSliceBytes(i))
	var n uint64
	size := code.SizeofBytesSlice(i)
	fmt.Printf("SizeofBytesSlice:%v sizeof:%d\n", i, size)
	n = code.EncodeBytesSlice(buf, i)
	fmt.Printf("EncodeBytesSlice:%v to []byte:%v\n", i, buf[:n])
	var v = make([][]byte, 2)
	n = code.DecodeBytesSlice(buf[:n], &v)
	fmt.Printf("DecodeBytesSlice:%v,length:%d\n", v, n)
}
