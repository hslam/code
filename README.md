# code
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hslam/code)](https://pkg.go.dev/github.com/hslam/code)
[![Build Status](https://travis-ci.org/hslam/code.svg?branch=master)](https://travis-ci.org/hslam/code)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/code)](https://goreportcard.com/report/github.com/hslam/code)
[![GitHub release](https://img.shields.io/github/release/hslam/code.svg)](https://github.com/hslam/code/releases/latest)
[![LICENSE](https://img.shields.io/github/license/hslam/code.svg?style=flat-square)](https://github.com/hslam/code/blob/master/LICENSE)

Package code implements encoding and decoding in golang.

## Feature
* Uint8
* Uint16
* Uint32
* Uint64
* Varint
* Float32
* Float64
* Bool
* String
* Bytes
* Uint8Slice
* Uint16Slice
* Uint32Slice
* Uint64Slice
* VarintSlice
* Float32Slice
* Float64Slice
* BoolSlice
* StringSlice
* BytesSlice

## Get started

### Install
```
go get github.com/hslam/code
```
### Import
```
import "github.com/hslam/code"
```
### Usage
#### Example
```go
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
```

### Output
```
SizeofUint8:128 sizeof:1
EncodeUint8:128 to []byte:[128]
DecodeUint8:128,length:1
SizeofUint16:128 sizeof:2
EncodeUint16:128 to []byte:[128 0]
DecodeUint16:128,length:2
SizeofUint32:128 sizeof:4
EncodeUint32:128 to []byte:[128 0 0 0]
DecodeUint32:128,length:4
SizeofUint64:128 sizeof:8
EncodeUint64:128 to []byte:[128 0 0 0 0 0 0 0]
DecodeUint64:128,length:8
SizeofVarint:128 sizeof:2
EncodeVarint:128 to []byte:[128 1]
DecodeVarint:128,length:2
SizeofFloat32:3.14 sizeof:4
EncodeFloat32:3.14 to []byte:[195 245 72 64]
EncodeFloat32:3.14,length:4
SizeofFloat64:3.1415926 sizeof:8
EncodeFloat64:3.1415926 to []byte:[74 216 18 77 251 33 9 64]
DecodeFloat64:3.1415926,length:8
SizeofBool:true sizeof:1
EncodeBool:true to []byte:[1]
DecodeBool:true,length:1
SizeofString:Hello sizeof:6
EncodeString:Hello to []byte:[5 72 101 108 108 111]
DecodeString:Hello,length:6
SizeofBytes:[1 2] sizeof:3
EncodeBytes:[1 2] to []byte:[2 1 2]
DecodeBytes:[1 2],length:3
SizeofUint8Slice:[128 255] sizeof:3
EncodeUint8Slice:[128 255] to []byte:[2 128 255]
DecodeUint8Slice:[128 255],length:3
SizeofUint16Slice:[128 256] sizeof:5
EncodeUint16Slice:[128 256] to []byte:[2 128 0 0 1]
DecodeUint16Slice:[128 256],length:5
SizeofUint32Slice:[128 256] sizeof:9
EncodeUint32Slice:[128 256] to []byte:[2 128 0 0 0 0 1 0 0]
DecodeUint32Slice:[128 256],length:9
SizeofUint64Slice:[128 256] sizeof:17
EncodeUint64Slice:[128 256] to []byte:[2 128 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0]
DecodeUint64Slice:[128 256],length:17
SizeofVarintSlice:[128 256] sizeof:5
EncodeVarintSlice:[128 256] to []byte:[2 128 1 128 2]
DecodeVarintSlice:[128 256],length:5
SizeofFloat32Slice:[3.14] sizeof:5
EncodeFloat32Slice:[3.14] to []byte:[1 195 245 72 64]
DecodeFloat32Slice:[3.14],length:5
SizeofFloat64Slice:[3.1415926] sizeof:9
EncodeFloat64Slice:[3.1415926] to []byte:[1 74 216 18 77 251 33 9 64]
DecodeFloat64Slice:[3.1415926],length:9
SizeofBoolSlice:[true false] sizeof:3
EncodeBoolSlice:[true false] to []byte:[2 1 0]
DecodeBoolSlice:[true false],length:3
SizeofStringSlice:[Hello World] sizeof:13
EncodeStringSlice:[Hello World] to []byte:[2 5 72 101 108 108 111 5 87 111 114 108 100]
DecodeStringSlice:[Hello World],length:13
SizeofBytesSlice:[[1 2] [3]] sizeof:6
EncodeBytesSlice:[[1 2] [3]] to []byte:[2 2 1 2 1 3]
DecodeBytesSlice:[[1 2] [3]],length:6
```

### Benchmark
go test -v -run="none" -bench=. -benchtime=30s
```
goos: darwin
goarch: amd64
pkg: github.com/hslam/code
BenchmarkCheckBuffer-4            	1000000000	         0.617 ns/op	414612.88 MB/s
BenchmarkCodeUint8-4              	1000000000	         0.310 ns/op	3229.52 MB/s
BenchmarkCodeUint16-4             	1000000000	         0.310 ns/op	6453.14 MB/s
BenchmarkCodeUint32-4             	1000000000	         0.309 ns/op	12937.68 MB/s
BenchmarkCodeUint64-4             	1000000000	         0.310 ns/op	25797.99 MB/s
BenchmarkCodeVarint-4             	1000000000	         7.42 ns/op	 269.70 MB/s
BenchmarkBinaryVarint-4           	1000000000	         8.87 ns/op	 225.47 MB/s
BenchmarkCodeFloat32-4            	1000000000	         0.336 ns/op	11888.17 MB/s
BenchmarkCodeFloat64-4            	1000000000	         0.926 ns/op	8643.01 MB/s
BenchmarkCodeBool-4               	1000000000	         0.314 ns/op	3185.92 MB/s
BenchmarkCodeString-4             	1000000000	        10.9 ns/op	 183.75 MB/s
BenchmarkCodeBytes-4              	1000000000	        10.8 ns/op	 184.66 MB/s
BenchmarkCodeUint8Slice-4         	1000000000	        10.7 ns/op	 186.36 MB/s
BenchmarkCodeUint16Slice-4        	1000000000	         8.97 ns/op	 334.55 MB/s
BenchmarkCodeUint32Slice-4        	1000000000	        10.5 ns/op	 474.01 MB/s
BenchmarkCodeUint64Slice-4        	1000000000	        14.2 ns/op	 634.65 MB/s
BenchmarkCodeVarintSlice-4        	1000000000	        11.1 ns/op	 179.58 MB/s
BenchmarkCodeFloat32Slice-4       	1000000000	        12.8 ns/op	 390.38 MB/s
BenchmarkCodeFloat64Slice-4       	1000000000	        16.2 ns/op	 555.53 MB/s
BenchmarkCodeBoolSlice-4          	1000000000	         9.63 ns/op	 207.60 MB/s
BenchmarkCodeStringSlice-4        	1000000000	        30.3 ns/op	 164.94 MB/s
BenchmarkCodeBytesSlice-4         	1000000000	        30.0 ns/op	 166.70 MB/s
BenchmarkSizeofUint8-4            	1000000000	         0.310 ns/op	3224.62 MB/s
BenchmarkSizeofUint16-4           	1000000000	         0.321 ns/op	6221.12 MB/s
BenchmarkSizeofUint32-4           	1000000000	         0.310 ns/op	12904.22 MB/s
BenchmarkSizeofUint64-4           	1000000000	         0.309 ns/op	25885.08 MB/s
BenchmarkSizeofVarint-4           	1000000000	         0.313 ns/op	6393.58 MB/s
BenchmarkSizeofFloat32-4          	1000000000	         0.312 ns/op	12834.91 MB/s
BenchmarkSizeofFloat64-4          	1000000000	         0.310 ns/op	25811.82 MB/s
BenchmarkSizeofBool-4             	1000000000	         0.310 ns/op	3227.18 MB/s
BenchmarkSizeofString-4           	1000000000	         0.314 ns/op	3184.25 MB/s
BenchmarkSizeofBytes-4            	1000000000	         0.312 ns/op	3207.41 MB/s
BenchmarkSizeofUint8Slice-4       	1000000000	         0.310 ns/op	3220.83 MB/s
BenchmarkSizeofUint16Slice-4      	1000000000	         0.311 ns/op	3210.29 MB/s
BenchmarkSizeofUint32Slice-4      	1000000000	         0.312 ns/op	3205.69 MB/s
BenchmarkSizeofUint64Slice-4      	1000000000	         0.309 ns/op	3232.42 MB/s
BenchmarkSizeofVarintSlice-4      	1000000000	         3.07 ns/op	 325.79 MB/s
BenchmarkSizeofFloat32Slice-4     	1000000000	         0.310 ns/op	12901.08 MB/s
BenchmarkSizeofFloat64Slice-4     	1000000000	         0.311 ns/op	25684.42 MB/s
BenchmarkSizeofBoolSlice-4        	1000000000	         0.312 ns/op	3205.72 MB/s
BenchmarkSizeofStringSlice-4      	1000000000	         3.16 ns/op	 632.50 MB/s
BenchmarkSizeofBytesSlice-4       	1000000000	         4.01 ns/op	 498.30 MB/s
BenchmarkMaxUint8Bytes-4          	1000000000	         0.313 ns/op	3194.00 MB/s
BenchmarkMaxUint16Bytes-4         	1000000000	         0.310 ns/op	6460.12 MB/s
BenchmarkMaxUint32Bytes-4         	1000000000	         0.311 ns/op	12854.70 MB/s
BenchmarkMaxUint64Bytes-4         	1000000000	         0.310 ns/op	25837.95 MB/s
BenchmarkMaxVarintBytes-4         	1000000000	         0.312 ns/op	6415.42 MB/s
BenchmarkMaxFloat32Bytes-4        	1000000000	         0.309 ns/op	12946.37 MB/s
BenchmarkMaxFloat64Bytes-4        	1000000000	         0.310 ns/op	25809.76 MB/s
BenchmarkMaxBoolBytes-4           	1000000000	         0.313 ns/op	3193.68 MB/s
BenchmarkMaxStringBytes-4         	1000000000	         0.311 ns/op	3214.47 MB/s
BenchmarkMaxBytesBytes-4          	1000000000	         0.310 ns/op	3226.55 MB/s
BenchmarkMaxUint8SliceBytes-4     	1000000000	         0.311 ns/op	3210.37 MB/s
BenchmarkMaxUint16SliceBytes-4    	1000000000	         0.312 ns/op	3210.11 MB/s
BenchmarkMaxUint32SliceBytes-4    	1000000000	         0.310 ns/op	3226.48 MB/s
BenchmarkMaxUint64SliceBytes-4    	1000000000	         0.309 ns/op	3237.72 MB/s
BenchmarkMaxVarintSliceBytes-4    	1000000000	         0.313 ns/op	3193.07 MB/s
BenchmarkMaxFloat32SliceBytes-4   	1000000000	         0.310 ns/op	12913.98 MB/s
BenchmarkMaxFloat64SliceBytes-4   	1000000000	         0.309 ns/op	25866.04 MB/s
BenchmarkMaxBoolSliceBytes-4      	1000000000	         0.312 ns/op	3208.88 MB/s
BenchmarkMaxStringSliceBytes-4    	1000000000	         3.11 ns/op	 642.38 MB/s
BenchmarkMaxBytesSliceBytes-4     	1000000000	         2.68 ns/op	 746.79 MB/s
PASS
ok  	github.com/hslam/code	246.123s
```

### License
This package is licensed under a MIT license (Copyright (c) 2019 Meng Huang)

### Author
code was written by Meng Huang.
