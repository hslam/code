# code
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
	"github.com/hslam/code"
	"fmt"
)
func main()  {
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
func Uint8()  {
	var buf =make([]byte,4)
	var i uint8=128
	var n uint64
	size:=code.SizeofUint8()
	fmt.Printf("SizeofUint8:%d sizeof:%d\n",i,size)
	n=code.EncodeUint8(buf,i)
	fmt.Printf("EncodeUint8:%d to []byte:%v\n",i,buf[:n])
	var v uint8
	n=code.DecodeUint8(buf[:n],&v)
	fmt.Printf("DecodeUint8:%d,length:%d\n",v,n)
}
func Uint16()  {
	var buf =make([]byte,4)
	var i uint16=128
	var n uint64
	size:=code.SizeofUint16()
	fmt.Printf("SizeofUint16:%d sizeof:%d\n",i,size)
	n=code.EncodeUint16(buf,i)
	fmt.Printf("EncodeUint16:%d to []byte:%v\n",i,buf[:n])
	var v uint16
	n=code.DecodeUint16(buf[:n],&v)
	fmt.Printf("DecodeUint16:%d,length:%d\n",v,n)
}
func Uint32()  {
	var buf =make([]byte,4)
	var i uint32=128
	var n uint64
	size:=code.SizeofUint32()
	fmt.Printf("SizeofUint32:%d sizeof:%d\n",i,size)
	n=code.EncodeUint32(buf,i)
	fmt.Printf("EncodeUint32:%d to []byte:%v\n",i,buf[:n])
	var v uint32
	n=code.DecodeUint32(buf[:n],&v)
	fmt.Printf("DecodeUint32:%d,length:%d\n",v,n)
}
func Uint64()  {
	var buf =make([]byte,8)
	var i uint64=128
	var n uint64
	size:=code.SizeofUint64()
	fmt.Printf("SizeofUint64:%d sizeof:%d\n",i,size)
	n=code.EncodeUint64(buf,i)
	fmt.Printf("EncodeUint64:%d to []byte:%v\n",i,buf[:n])
	var v uint64
	n=code.DecodeUint64(buf[:n],&v)
	fmt.Printf("DecodeUint64:%d,length:%d\n",v,n)
}

func Varint()  {
	var buf =make([]byte,10)
	var i uint64=128
	var n uint64
	size:=code.SizeofVarint(i)
	fmt.Printf("SizeofVarint:%d sizeof:%d\n",i,size)
	n=code.EncodeVarint(buf,i)
	fmt.Printf("EncodeVarint:%d to []byte:%v\n",i,buf[:n])
	var v uint64
	n=code.DecodeVarint(buf[:n],&v)
	fmt.Printf("DecodeVarint:%d,length:%d\n",v,n)
}
func Float32()  {
	var buf =make([]byte,9)
	var i float32=3.14
	var n uint64
	size:=code.SizeofFloat32()
	fmt.Printf("SizeofFloat32:%.2f sizeof:%d\n",i,size)
	n=code.EncodeFloat32(buf,i)
	fmt.Printf("EncodeFloat32:%.2f to []byte:%v\n",i,buf[:n])
	var v float32
	n=code.DecodeFloat32(buf[:n],&v)
	fmt.Printf("EncodeFloat32:%.2f,length:%d\n",v,n)
}
func Float64()  {
	var buf =make([]byte,9)
	var i float64=3.1415926
	var n uint64
	size:=code.SizeofFloat64()
	fmt.Printf("SizeofFloat64:%.2f sizeof:%d\n",i,size)
	n=code.EncodeFloat64(buf,i)
	fmt.Printf("EncodeFloat64:%.2f to []byte:%v\n",i,buf[:n])
	var v float64
	n=code.DecodeFloat64(buf[:n],&v)
	fmt.Printf("DecodeFloat64:%.2f,length:%d\n",v,n)
}
func Bool()  {
	var buf =make([]byte,16)
	var i bool=true
	var n uint64
	size:=code.SizeofBool()
	fmt.Printf("SizeofBool:%t sizeof:%d\n",i,size)
	n=code.EncodeBool(buf,i)
	fmt.Printf("EncodeBool:%t to []byte:%v\n",i,buf[:n])
	var v bool
	n=code.DecodeBool(buf[:n],&v)
	fmt.Printf("DecodeBool:%t,length:%d\n",v,n)
}
func String()  {
	var buf =make([]byte,16)
	var i string="Hello"
	var n uint64
	size:=code.SizeofString(i)
	fmt.Printf("SizeofString:%s sizeof:%d\n",i,size)
	n=code.EncodeString(buf,i)
	fmt.Printf("EncodeString:%s to []byte:%v\n",i,buf[:n])
	var v string
	n=code.DecodeString(buf[:n],&v)
	fmt.Printf("DecodeString:%s,length:%d\n",v,n)
}
func Bytes()  {
	var buf =make([]byte,16)
	var i []byte=[]byte{1,2}
	var n uint64
	size:=code.SizeofBytes(i)
	fmt.Printf("SizeofBytes:%v sizeof:%d\n",i,size)
	n=code.EncodeBytes(buf,i)
	fmt.Printf("EncodeBytes:%v to []byte:%v\n",i,buf[:n])
	var v =make([]byte,2)
	n=code.DecodeBytes(buf[:n],&v)
	fmt.Printf("DecodeBytes:%v,length:%d\n",v,n)
}


func SliceUint8()  {
	var buf =make([]byte,64)
	var i []uint8=[]uint8{128,255}
	var n uint64
	size:=code.SizeofSliceUint8(i)
	fmt.Printf("SizeofSliceUint8:%v sizeof:%d\n",i,size)
	n=code.EncodeSliceUint8(buf,i)
	fmt.Printf("EncodeSliceUint8:%v to []byte:%v\n",i,buf[:n])
	var v =make([]uint8,2)
	n=code.DecodeSliceUint8(buf[:n],&v)
	fmt.Printf("DecodeSliceUint8:%v,length:%d\n",v,n)
}

func SliceUint16()  {
	var buf =make([]byte,64)
	var i []uint16=[]uint16{128,256}
	var n uint64
	size:=code.SizeofSliceUint16(i)
	fmt.Printf("SizeofSliceUint16:%v sizeof:%d\n",i,size)
	n=code.EncodeSliceUint16(buf,i)
	fmt.Printf("EncodeSliceUint16:%v to []byte:%v\n",i,buf[:n])
	var v =make([]uint16,2)
	n=code.DecodeSliceUint16(buf[:n],&v)
	fmt.Printf("DecodeSliceUint16:%v,length:%d\n",v,n)
}

func SliceUint32()  {
	var buf =make([]byte,64)
	var i []uint32=[]uint32{128,256}
	var n uint64
	size:=code.SizeofSliceUint32(i)
	fmt.Printf("SizeofSliceUint32:%v sizeof:%d\n",i,size)
	n=code.EncodeSliceUint32(buf,i)
	fmt.Printf("EncodeSliceUint32:%v to []byte:%v\n",i,buf[:n])
	var v =make([]uint32,2)
	n=code.DecodeSliceUint32(buf[:n],&v)
	fmt.Printf("DecodeSliceUint32:%v,length:%d\n",v,n)
}

func SliceUint64()  {
	var buf =make([]byte,64)
	var i []uint64=[]uint64{128,256}
	var n uint64
	size:=code.SizeofSliceUint64(i)
	fmt.Printf("SizeofSliceUint64:%v sizeof:%d\n",i,size)
	n=code.EncodeSliceUint64(buf,i)
	fmt.Printf("EncodeSliceUint64:%v to []byte:%v\n",i,buf[:n])
	var v =make([]uint64,2)
	n=code.DecodeSliceUint64(buf[:n],&v)
	fmt.Printf("DecodeSliceUint64:%v,length:%d\n",v,n)
}

func SliceVarint()  {
	var buf =make([]byte,64)
	var i []uint64=[]uint64{128,256}
	var n uint64
	size:=code.SizeofSliceVarint(i)
	fmt.Printf("SizeofSliceVarint:%v sizeof:%d\n",i,size)
	n=code.EncodeSliceVarint(buf,i)
	fmt.Printf("EncodeSliceVarint:%v to []byte:%v\n",i,buf[:n])
	var v =make([]uint64,2)
	n=code.DecodeSliceVarint(buf[:n],&v)
	fmt.Printf("DecodeSliceVarint:%v,length:%d\n",v,n)
}

func SliceFloat32()  {
	var buf =make([]byte,64)
	var i []float32=[]float32{3.14}
	var n uint64
	size:=code.SizeofSliceFloat32(i)
	fmt.Printf("SizeofSliceFloat32:%v sizeof:%d\n",i,size)
	n=code.EncodeSliceFloat32(buf,i)
	fmt.Printf("EncodeSliceFloat32:%v to []byte:%v\n",i,buf[:n])
	var v =make([]float32,2)
	n=code.DecodeSliceFloat32(buf[:n],&v)
	fmt.Printf("DecodeSliceFloat32:%v,length:%d\n",v,n)
}

func SliceFloat64()  {
	var buf =make([]byte,64)
	var i []float64=[]float64{3.1415926}
	var n uint64
	size:=code.SizeofSliceFloat64(i)
	fmt.Printf("SizeofSliceFloat64:%v sizeof:%d\n",i,size)
	n=code.EncodeSliceFloat64(buf,i)
	fmt.Printf("EncodeSliceFloat64:%v to []byte:%v\n",i,buf[:n])
	var v =make([]float64,2)
	n=code.DecodeSliceFloat64(buf[:n],&v)
	fmt.Printf("DecodeSliceFloat64:%v,length:%d\n",v,n)
}

func SliceBool()  {
	var buf =make([]byte,64)
	var i []bool=[]bool{true,false}
	var n uint64
	size:=code.SizeofSliceBool(i)
	fmt.Printf("SizeofSliceBool:%v sizeof:%d\n",i,size)
	n=code.EncodeSliceBool(buf,i)
	fmt.Printf("EncodeSliceBool:%v to []byte:%v\n",i,buf[:n])
	var v =make([]bool,2)
	n=code.DecodeSliceBool(buf[:n],&v)
	fmt.Printf("DecodeSliceBool:%v,length:%d\n",v,n)
}
func SliceString()  {
	var buf =make([]byte,64)
	var i []string=[]string{"Hello","World"}
	var n uint64
	size:=code.SizeofSliceString(i)
	fmt.Printf("SizeofSliceString:%v sizeof:%d\n",i,size)
	n=code.EncodeSliceString(buf,i)
	fmt.Printf("EncodeSliceString:%v to []byte:%v\n",i,buf[:n])
	var v =make([]string,2)
	n=code.DecodeSliceString(buf[:n],&v)
	fmt.Printf("DecodeSliceString:%v,length:%d\n",v,n)
}
func SliceBytes()  {
	var buf =make([]byte,64)
	var i [][]byte=[][]byte{{1,2},{3}}
	var n uint64
	size:=code.SizeofSliceBytes(i)
	fmt.Printf("SizeofSliceBytes:%v sizeof:%d\n",i,size)
	n=code.EncodeSliceBytes(buf,i)
	fmt.Printf("EncodeSliceBytes:%v to []byte:%v\n",i,buf[:n])
	var v =make([][]byte,2)
	n=code.DecodeSliceBytes(buf[:n],&v)
	fmt.Printf("DecodeSliceBytes:%v,length:%d\n",v,n)
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
BenchmarkCodeUint8-4            	1000000000	         0.319 ns/op	3131.82 MB/s
BenchmarkCodeUint16-4           	1000000000	         0.316 ns/op	6338.31 MB/s
BenchmarkCodeUint32-4           	1000000000	         0.314 ns/op	12725.54 MB/s
BenchmarkCodeUint64-4           	1000000000	         3.47 ns/op	2307.95 MB/s
BenchmarkCodeVarint-4           	1000000000	         7.02 ns/op	 285.07 MB/s
BenchmarkCodeFloat32-4          	1000000000	         0.626 ns/op	6387.92 MB/s
BenchmarkCodeFloat64-4          	1000000000	         3.51 ns/op	2276.50 MB/s
BenchmarkCodeBool-4             	1000000000	         0.315 ns/op	3170.87 MB/s
BenchmarkCodeString-4           	1000000000	         9.32 ns/op	 214.54 MB/s
BenchmarkCodeBytes-4            	1000000000	        10.1 ns/op	 198.41 MB/s
BenchmarkCodeSliceUint8-4       	1000000000	         8.11 ns/op	 246.48 MB/s
BenchmarkCodeSliceUint16-4      	1000000000	         9.12 ns/op	 328.77 MB/s
BenchmarkCodeSliceUint32-4      	1000000000	         9.80 ns/op	 510.36 MB/s
BenchmarkCodeSliceUint64-4      	1000000000	        11.6 ns/op	 774.18 MB/s
BenchmarkCodeSliceVarint-4      	1000000000	        10.1 ns/op	 198.59 MB/s
BenchmarkCodeSliceFloat32-4     	1000000000	        10.3 ns/op	 484.21 MB/s
BenchmarkCodeSliceFloat64-4     	1000000000	        13.4 ns/op	 672.42 MB/s
BenchmarkCodeSliceBool-4        	1000000000	         8.45 ns/op	 236.82 MB/s
BenchmarkCodeSliceString-4      	1000000000	        24.6 ns/op	 203.30 MB/s
BenchmarkCodeSliceBytes-4       	1000000000	        26.2 ns/op	 191.20 MB/s
BenchmarkSizeofUint8-4          	1000000000	         0.321 ns/op	3111.28 MB/s
BenchmarkSizeofUint16-4         	1000000000	         0.323 ns/op	6194.01 MB/s
BenchmarkSizeofUint32-4         	1000000000	         0.315 ns/op	12705.34 MB/s
BenchmarkSizeofUint64-4         	1000000000	         0.315 ns/op	25364.37 MB/s
BenchmarkSizeofVarint-4         	1000000000	         0.314 ns/op	6362.95 MB/s
BenchmarkSizeofFloat32-4        	1000000000	         0.313 ns/op	12764.37 MB/s
BenchmarkSizeofFloat64-4        	1000000000	         0.313 ns/op	25520.28 MB/s
BenchmarkSizeofBool-4           	1000000000	         0.315 ns/op	3175.97 MB/s
BenchmarkSizeofString-4         	1000000000	         0.319 ns/op	3132.46 MB/s
BenchmarkSizeofBytes-4          	1000000000	         0.313 ns/op	3196.14 MB/s
BenchmarkSizeofSliceUint8-4     	1000000000	         0.316 ns/op	3165.90 MB/s
BenchmarkSizeofSliceUint16-4    	1000000000	         0.315 ns/op	3171.74 MB/s
BenchmarkSizeofSliceUint32-4    	1000000000	         0.313 ns/op	3191.90 MB/s
BenchmarkSizeofSliceUint64-4    	1000000000	         0.327 ns/op	3061.26 MB/s
BenchmarkSizeofSliceVarint-4    	1000000000	         2.98 ns/op	 335.80 MB/s
BenchmarkSizeofSliceFloat32-4   	1000000000	         0.314 ns/op	12730.18 MB/s
BenchmarkSizeofSliceFloat64-4   	1000000000	         0.317 ns/op	25235.57 MB/s
BenchmarkSizeofSliceBool-4      	1000000000	         0.315 ns/op	3178.47 MB/s
BenchmarkSizeofSliceString-4    	1000000000	         3.27 ns/op	 611.40 MB/s
BenchmarkSizeofSliceBytes-4     	1000000000	         3.12 ns/op	 641.27 MB/s
PASS
ok  	github.com/hslam/code	200.258s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Meng Huang)

### Authors
code was written by Meng Huang.
