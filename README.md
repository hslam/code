# code
[![GoDoc](https://godoc.org/github.com/hslam/code?status.svg)](https://godoc.org/github.com/hslam/code)
[![Build Status](https://travis-ci.org/hslam/code.svg?branch=master)](https://travis-ci.org/hslam/code)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/code)](https://goreportcard.com/report/github.com/hslam/code)
[![GitHub release](https://img.shields.io/github/release/hslam/code.svg)](https://github.com/hslam/code/releases/latest)
[![LICENSE](https://img.shields.io/github/license/hslam/code.svg?style=flat-square)](https://github.com/hslam/code/blob/master/LICENSE)

A code library written in golang for encoding and decoding.

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
* SliceUint8
* SliceUint16
* SliceUint32
* SliceUint64
* SliceVarint
* SliceFloat32
* SliceFloat64
* SliceBool
* SliceString
* SliceBytes

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
```
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
	SliceUint8()
	SliceUint16()
	SliceUint32()
	SliceUint64()
	SliceVarint()
	SliceFloat32()
	SliceFloat64()
	SliceBool()
	SliceString()
	SliceBytes()
}

//Uint8 example
func Uint8() {
	var buf = make([]byte, 4)
	var i uint8 = 128
	var n uint64
	size := code.SizeofUint8(i)
	fmt.Printf("SizeofUint8:%d sizeof:%d\n", i, size)
	n = code.EncodeUint8(buf, i)
	fmt.Printf("EncodeUint8:%d to []byte:%v\n", i, buf[:n])
	var v uint8
	n = code.DecodeUint8(buf[:n], &v)
	fmt.Printf("DecodeUint8:%d,length:%d\n", v, n)
}

//Uint16 example
func Uint16() {
	var buf = make([]byte, 4)
	var i uint16 = 128
	var n uint64
	size := code.SizeofUint16(i)
	fmt.Printf("SizeofUint16:%d sizeof:%d\n", i, size)
	n = code.EncodeUint16(buf, i)
	fmt.Printf("EncodeUint16:%d to []byte:%v\n", i, buf[:n])
	var v uint16
	n = code.DecodeUint16(buf[:n], &v)
	fmt.Printf("DecodeUint16:%d,length:%d\n", v, n)
}

//Uint32 example
func Uint32() {
	var buf = make([]byte, 4)
	var i uint32 = 128
	var n uint64
	size := code.SizeofUint32(i)
	fmt.Printf("SizeofUint32:%d sizeof:%d\n", i, size)
	n = code.EncodeUint32(buf, i)
	fmt.Printf("EncodeUint32:%d to []byte:%v\n", i, buf[:n])
	var v uint32
	n = code.DecodeUint32(buf[:n], &v)
	fmt.Printf("DecodeUint32:%d,length:%d\n", v, n)
}

//Uint64 example
func Uint64() {
	var buf = make([]byte, 8)
	var i uint64 = 128
	var n uint64
	size := code.SizeofUint64(i)
	fmt.Printf("SizeofUint64:%d sizeof:%d\n", i, size)
	n = code.EncodeUint64(buf, i)
	fmt.Printf("EncodeUint64:%d to []byte:%v\n", i, buf[:n])
	var v uint64
	n = code.DecodeUint64(buf[:n], &v)
	fmt.Printf("DecodeUint64:%d,length:%d\n", v, n)
}

//Varint example
func Varint() {
	var buf = make([]byte, 10)
	var i uint64 = 128
	var n uint64
	size := code.SizeofVarint(i)
	fmt.Printf("SizeofVarint:%d sizeof:%d\n", i, size)
	n = code.EncodeVarint(buf, i)
	fmt.Printf("EncodeVarint:%d to []byte:%v\n", i, buf[:n])
	var v uint64
	n = code.DecodeVarint(buf[:n], &v)
	fmt.Printf("DecodeVarint:%d,length:%d\n", v, n)
}

//Float32 example
func Float32() {
	var buf = make([]byte, 9)
	var i float32 = 3.14
	var n uint64
	size := code.SizeofFloat32(i)
	fmt.Printf("SizeofFloat32:%.2f sizeof:%d\n", i, size)
	n = code.EncodeFloat32(buf, i)
	fmt.Printf("EncodeFloat32:%.2f to []byte:%v\n", i, buf[:n])
	var v float32
	n = code.DecodeFloat32(buf[:n], &v)
	fmt.Printf("EncodeFloat32:%.2f,length:%d\n", v, n)
}

//Float64 example
func Float64() {
	var buf = make([]byte, 9)
	var i float64 = 3.1415926
	var n uint64
	size := code.SizeofFloat64(i)
	fmt.Printf("SizeofFloat64:%.2f sizeof:%d\n", i, size)
	n = code.EncodeFloat64(buf, i)
	fmt.Printf("EncodeFloat64:%.2f to []byte:%v\n", i, buf[:n])
	var v float64
	n = code.DecodeFloat64(buf[:n], &v)
	fmt.Printf("DecodeFloat64:%.2f,length:%d\n", v, n)
}

//Bool example
func Bool() {
	var buf = make([]byte, 16)
	var i bool = true
	var n uint64
	size := code.SizeofBool(i)
	fmt.Printf("SizeofBool:%t sizeof:%d\n", i, size)
	n = code.EncodeBool(buf, i)
	fmt.Printf("EncodeBool:%t to []byte:%v\n", i, buf[:n])
	var v bool
	n = code.DecodeBool(buf[:n], &v)
	fmt.Printf("DecodeBool:%t,length:%d\n", v, n)
}

//String example
func String() {
	var buf = make([]byte, 16)
	var i string = "Hello"
	var n uint64
	size := code.SizeofString(i)
	fmt.Printf("SizeofString:%s sizeof:%d\n", i, size)
	n = code.EncodeString(buf, i)
	fmt.Printf("EncodeString:%s to []byte:%v\n", i, buf[:n])
	var v string
	n = code.DecodeString(buf[:n], &v)
	fmt.Printf("DecodeString:%s,length:%d\n", v, n)
}

//Bytes example
func Bytes() {
	var buf = make([]byte, 16)
	var i []byte = []byte{1, 2}
	var n uint64
	size := code.SizeofBytes(i)
	fmt.Printf("SizeofBytes:%v sizeof:%d\n", i, size)
	n = code.EncodeBytes(buf, i)
	fmt.Printf("EncodeBytes:%v to []byte:%v\n", i, buf[:n])
	var v = make([]byte, 2)
	n = code.DecodeBytes(buf[:n], &v)
	fmt.Printf("DecodeBytes:%v,length:%d\n", v, n)
}

//SliceUint8 example
func SliceUint8() {
	var buf = make([]byte, 64)
	var i []uint8 = []uint8{128, 255}
	var n uint64
	size := code.SizeofSliceUint8(i)
	fmt.Printf("SizeofSliceUint8:%v sizeof:%d\n", i, size)
	n = code.EncodeSliceUint8(buf, i)
	fmt.Printf("EncodeSliceUint8:%v to []byte:%v\n", i, buf[:n])
	var v = make([]uint8, 2)
	n = code.DecodeSliceUint8(buf[:n], &v)
	fmt.Printf("DecodeSliceUint8:%v,length:%d\n", v, n)
}

//SliceUint16 example
func SliceUint16() {
	var buf = make([]byte, 64)
	var i []uint16 = []uint16{128, 256}
	var n uint64
	size := code.SizeofSliceUint16(i)
	fmt.Printf("SizeofSliceUint16:%v sizeof:%d\n", i, size)
	n = code.EncodeSliceUint16(buf, i)
	fmt.Printf("EncodeSliceUint16:%v to []byte:%v\n", i, buf[:n])
	var v = make([]uint16, 2)
	n = code.DecodeSliceUint16(buf[:n], &v)
	fmt.Printf("DecodeSliceUint16:%v,length:%d\n", v, n)
}

//SliceUint32 example
func SliceUint32() {
	var buf = make([]byte, 64)
	var i []uint32 = []uint32{128, 256}
	var n uint64
	size := code.SizeofSliceUint32(i)
	fmt.Printf("SizeofSliceUint32:%v sizeof:%d\n", i, size)
	n = code.EncodeSliceUint32(buf, i)
	fmt.Printf("EncodeSliceUint32:%v to []byte:%v\n", i, buf[:n])
	var v = make([]uint32, 2)
	n = code.DecodeSliceUint32(buf[:n], &v)
	fmt.Printf("DecodeSliceUint32:%v,length:%d\n", v, n)
}

//SliceUint64 example
func SliceUint64() {
	var buf = make([]byte, 64)
	var i []uint64 = []uint64{128, 256}
	var n uint64
	size := code.SizeofSliceUint64(i)
	fmt.Printf("SizeofSliceUint64:%v sizeof:%d\n", i, size)
	n = code.EncodeSliceUint64(buf, i)
	fmt.Printf("EncodeSliceUint64:%v to []byte:%v\n", i, buf[:n])
	var v = make([]uint64, 2)
	n = code.DecodeSliceUint64(buf[:n], &v)
	fmt.Printf("DecodeSliceUint64:%v,length:%d\n", v, n)
}

//SliceVarint example
func SliceVarint() {
	var buf = make([]byte, 64)
	var i []uint64 = []uint64{128, 256}
	var n uint64
	size := code.SizeofSliceVarint(i)
	fmt.Printf("SizeofSliceVarint:%v sizeof:%d\n", i, size)
	n = code.EncodeSliceVarint(buf, i)
	fmt.Printf("EncodeSliceVarint:%v to []byte:%v\n", i, buf[:n])
	var v = make([]uint64, 2)
	n = code.DecodeSliceVarint(buf[:n], &v)
	fmt.Printf("DecodeSliceVarint:%v,length:%d\n", v, n)
}

//SliceFloat32 example
func SliceFloat32() {
	var buf = make([]byte, 64)
	var i []float32 = []float32{3.14}
	var n uint64
	size := code.SizeofSliceFloat32(i)
	fmt.Printf("SizeofSliceFloat32:%v sizeof:%d\n", i, size)
	n = code.EncodeSliceFloat32(buf, i)
	fmt.Printf("EncodeSliceFloat32:%v to []byte:%v\n", i, buf[:n])
	var v = make([]float32, 2)
	n = code.DecodeSliceFloat32(buf[:n], &v)
	fmt.Printf("DecodeSliceFloat32:%v,length:%d\n", v, n)
}

//SliceFloat64 example
func SliceFloat64() {
	var buf = make([]byte, 64)
	var i []float64 = []float64{3.1415926}
	var n uint64
	size := code.SizeofSliceFloat64(i)
	fmt.Printf("SizeofSliceFloat64:%v sizeof:%d\n", i, size)
	n = code.EncodeSliceFloat64(buf, i)
	fmt.Printf("EncodeSliceFloat64:%v to []byte:%v\n", i, buf[:n])
	var v = make([]float64, 2)
	n = code.DecodeSliceFloat64(buf[:n], &v)
	fmt.Printf("DecodeSliceFloat64:%v,length:%d\n", v, n)
}

//SliceBool example
func SliceBool() {
	var buf = make([]byte, 64)
	var i []bool = []bool{true, false}
	var n uint64
	size := code.SizeofSliceBool(i)
	fmt.Printf("SizeofSliceBool:%v sizeof:%d\n", i, size)
	n = code.EncodeSliceBool(buf, i)
	fmt.Printf("EncodeSliceBool:%v to []byte:%v\n", i, buf[:n])
	var v = make([]bool, 2)
	n = code.DecodeSliceBool(buf[:n], &v)
	fmt.Printf("DecodeSliceBool:%v,length:%d\n", v, n)
}

//SliceString example
func SliceString() {
	var buf = make([]byte, 64)
	var i []string = []string{"Hello", "World"}
	var n uint64
	size := code.SizeofSliceString(i)
	fmt.Printf("SizeofSliceString:%v sizeof:%d\n", i, size)
	n = code.EncodeSliceString(buf, i)
	fmt.Printf("EncodeSliceString:%v to []byte:%v\n", i, buf[:n])
	var v = make([]string, 2)
	n = code.DecodeSliceString(buf[:n], &v)
	fmt.Printf("DecodeSliceString:%v,length:%d\n", v, n)
}

//SliceBytes example
func SliceBytes() {
	var buf = make([]byte, 64)
	var i [][]byte = [][]byte{{1, 2}, {3}}
	var n uint64
	size := code.SizeofSliceBytes(i)
	fmt.Printf("SizeofSliceBytes:%v sizeof:%d\n", i, size)
	n = code.EncodeSliceBytes(buf, i)
	fmt.Printf("EncodeSliceBytes:%v to []byte:%v\n", i, buf[:n])
	var v = make([][]byte, 2)
	n = code.DecodeSliceBytes(buf[:n], &v)
	fmt.Printf("DecodeSliceBytes:%v,length:%d\n", v, n)
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
SizeofFloat64:3.14 sizeof:8
EncodeFloat64:3.14 to []byte:[74 216 18 77 251 33 9 64]
DecodeFloat64:3.14,length:8
SizeofBool:true sizeof:1
EncodeBool:true to []byte:[1]
DecodeBool:true,length:1
SizeofString:Hello sizeof:6
EncodeString:Hello to []byte:[5 72 101 108 108 111]
DecodeString:Hello,length:6
SizeofBytes:[1 2] sizeof:3
EncodeBytes:[1 2] to []byte:[2 1 2]
DecodeBytes:[1 2],length:3
SizeofSliceUint8:[128 255] sizeof:3
EncodeSliceUint8:[128 255] to []byte:[2 128 255]
DecodeSliceUint8:[128 255],length:3
SizeofSliceUint16:[128 256] sizeof:5
EncodeSliceUint16:[128 256] to []byte:[2 128 0 0 1]
DecodeSliceUint16:[128 256],length:5
SizeofSliceUint32:[128 256] sizeof:9
EncodeSliceUint32:[128 256] to []byte:[2 128 0 0 0 0 1 0 0]
DecodeSliceUint32:[128 256],length:9
SizeofSliceUint64:[128 256] sizeof:17
EncodeSliceUint64:[128 256] to []byte:[2 128 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0]
DecodeSliceUint64:[128 256],length:17
SizeofSliceVarint:[128 256] sizeof:5
EncodeSliceVarint:[128 256] to []byte:[2 128 1 128 2]
DecodeSliceVarint:[128 256],length:5
SizeofSliceFloat32:[3.14] sizeof:5
EncodeSliceFloat32:[3.14] to []byte:[1 195 245 72 64]
DecodeSliceFloat32:[3.14],length:5
SizeofSliceFloat64:[3.1415926] sizeof:9
EncodeSliceFloat64:[3.1415926] to []byte:[1 74 216 18 77 251 33 9 64]
DecodeSliceFloat64:[3.1415926],length:9
SizeofSliceBool:[true false] sizeof:3
EncodeSliceBool:[true false] to []byte:[2 1 0]
DecodeSliceBool:[true false],length:3
SizeofSliceString:[Hello World] sizeof:13
EncodeSliceString:[Hello World] to []byte:[2 5 72 101 108 108 111 5 87 111 114 108 100]
DecodeSliceString:[Hello World],length:13
SizeofSliceBytes:[[1 2] [3]] sizeof:6
EncodeSliceBytes:[[1 2] [3]] to []byte:[2 2 1 2 1 3]
DecodeSliceBytes:[[1 2] [3]],length:6
```

### Benchmark
go test -v -run="none" -bench=. -benchtime=30s
```
goos: darwin
goarch: amd64
pkg: github.com/hslam/code
BenchmarkCheckBuffer-4          	1000000000	         0.312 ns/op	819937.90 MB/s
BenchmarkCodeUint8-4            	1000000000	         0.312 ns/op	3208.75 MB/s
BenchmarkCodeUint16-4           	1000000000	         0.329 ns/op	6076.44 MB/s
BenchmarkCodeUint32-4           	1000000000	         0.342 ns/op	11688.81 MB/s
BenchmarkCodeUint64-4           	1000000000	         0.311 ns/op	25712.79 MB/s
BenchmarkCodeVarint-4           	1000000000	         7.21 ns/op	 277.23 MB/s
BenchmarkBinaryVarint-4         	1000000000	        10.1 ns/op	 197.22 MB/s
BenchmarkCodeFloat32-4          	1000000000	         0.311 ns/op	12881.38 MB/s
BenchmarkCodeFloat64-4          	1000000000	         0.924 ns/op	8658.87 MB/s
BenchmarkCodeBool-4             	1000000000	         0.330 ns/op	3031.09 MB/s
BenchmarkCodeString-4           	1000000000	        11.0 ns/op	 181.55 MB/s
BenchmarkCodeBytes-4            	1000000000	        10.9 ns/op	 183.31 MB/s
BenchmarkCodeSliceUint8-4       	1000000000	         9.49 ns/op	 210.84 MB/s
BenchmarkCodeSliceUint16-4      	1000000000	         9.95 ns/op	 301.51 MB/s
BenchmarkCodeSliceUint32-4      	1000000000	        10.7 ns/op	 467.08 MB/s
BenchmarkCodeSliceUint64-4      	1000000000	        15.0 ns/op	 601.80 MB/s
BenchmarkCodeSliceVarint-4      	1000000000	        11.4 ns/op	 175.34 MB/s
BenchmarkCodeSliceFloat32-4     	1000000000	        12.5 ns/op	 400.30 MB/s
BenchmarkCodeSliceFloat64-4     	1000000000	        16.1 ns/op	 560.15 MB/s
BenchmarkCodeSliceBool-4        	1000000000	        10.1 ns/op	 198.89 MB/s
BenchmarkCodeSliceString-4      	1000000000	        27.1 ns/op	 184.26 MB/s
BenchmarkCodeSliceBytes-4       	1000000000	        29.9 ns/op	 167.44 MB/s
BenchmarkSizeofUint8-4          	1000000000	         0.311 ns/op	3212.92 MB/s
BenchmarkSizeofUint16-4         	1000000000	         0.311 ns/op	6424.97 MB/s
BenchmarkSizeofUint32-4         	1000000000	         0.309 ns/op	12927.46 MB/s
BenchmarkSizeofUint64-4         	1000000000	         0.310 ns/op	25835.71 MB/s
BenchmarkSizeofVarint-4         	1000000000	         0.311 ns/op	6440.99 MB/s
BenchmarkSizeofFloat32-4        	1000000000	         0.310 ns/op	12904.85 MB/s
BenchmarkSizeofFloat64-4        	1000000000	         0.311 ns/op	25716.63 MB/s
BenchmarkSizeofBool-4           	1000000000	         0.310 ns/op	3228.32 MB/s
BenchmarkSizeofString-4         	1000000000	         0.311 ns/op	3215.73 MB/s
BenchmarkSizeofBytes-4          	1000000000	         0.337 ns/op	2968.08 MB/s
BenchmarkSizeofSliceUint8-4     	1000000000	         0.331 ns/op	3021.86 MB/s
BenchmarkSizeofSliceUint16-4    	1000000000	         0.312 ns/op	3210.18 MB/s
BenchmarkSizeofSliceUint32-4    	1000000000	         0.311 ns/op	3215.84 MB/s
BenchmarkSizeofSliceUint64-4    	1000000000	         0.311 ns/op	3211.50 MB/s
BenchmarkSizeofSliceVarint-4    	1000000000	         3.39 ns/op	 294.86 MB/s
BenchmarkSizeofSliceFloat32-4   	1000000000	         0.311 ns/op	12847.91 MB/s
BenchmarkSizeofSliceFloat64-4   	1000000000	         0.312 ns/op	25663.19 MB/s
BenchmarkSizeofSliceBool-4      	1000000000	         0.311 ns/op	3218.15 MB/s
BenchmarkSizeofSliceString-4    	1000000000	         3.40 ns/op	 588.11 MB/s
BenchmarkSizeofSliceBytes-4     	1000000000	         3.39 ns/op	 589.49 MB/s
PASS
ok  	github.com/hslam/code	232.615s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Meng Huang)

### Authors
code was written by Meng Huang.
