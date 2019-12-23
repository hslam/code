# code
A code library written in golang for encoding and decoding.

## Feature
* Uint8
* Uint16
* Uint32
* Uint64
* Int
* Varint
* Float32
* Float64
* Bool
* String
* Bytes
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
	Int()
	Varint()
	Float32()
	Float64()
	Bool()
	String()
	Bytes()
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
func Int()  {
	var buf =make([]byte,9)
	var i uint64=128
	var n uint64
	size:=code.SizeofInt(i)
	fmt.Printf("SizeofInt:%d sizeof:%d\n",i,size)
	n=code.EncodeInt(buf,i)
	fmt.Printf("EncodeInt:%d to []byte:%v\n",i,buf[:n])
	var v uint64
	n=code.DecodeInt(buf[:n],&v)
	fmt.Printf("DecodeInt:%d,length:%d\n",v,n)
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
	var i float64=3.14
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
func SliceBytes()  {
	var buf =make([]byte,16)
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
SizeofInt:128 sizeof:2
EncodeInt:128 to []byte:[1 128]
DecodeInt:128,length:2
SizeofVarint:128 sizeof:2
EncodeVarint:128 to []byte:[128 1]
DecodeVarint:128,length:2
SizeofFloat32:3.14 sizeof:4
EncodeFloat32:3.14 to []byte:[195 245 72 64]
EncodeFloat32:3.14,length:4
SizeofFloat64:3.14 sizeof:8
EncodeFloat64:3.14 to []byte:[31 133 235 81 184 30 9 64]
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
SizeofSliceBytes:[[1 2] [3]] sizeof:5
EncodeSliceBytes:[[1 2] [3]] to []byte:[2 1 2 1 3]
DecodeSliceBytes:[[1 2] [3]],length:5
```

### Benchmark
go test -v -run="none" -bench=. -benchtime=30s
```
goos: darwin
goarch: amd64
pkg: github.com/hslam/code
BenchmarkCodeUint8-4          	1000000000	         0.625 ns/op	1599.00 MB/s
BenchmarkCodeUint16-4         	1000000000	         0.620 ns/op	3228.22 MB/s
BenchmarkCodeUint32-4         	1000000000	         0.616 ns/op	6496.40 MB/s
BenchmarkCodeUint64-4         	1000000000	         3.41 ns/op	2349.20 MB/s
BenchmarkCodeInt-4            	1000000000	         7.77 ns/op	 257.26 MB/s
BenchmarkCodeVarint-4         	1000000000	         7.44 ns/op	 268.67 MB/s
BenchmarkCodeFloat32-4        	1000000000	         0.615 ns/op	6500.56 MB/s
BenchmarkCodeFloat64-4        	1000000000	         3.39 ns/op	2361.35 MB/s
BenchmarkCodeBool-4           	1000000000	         0.310 ns/op	3226.65 MB/s
BenchmarkCodeString-4         	1000000000	         9.07 ns/op	 220.59 MB/s
BenchmarkCodeBytes-4          	1000000000	         9.70 ns/op	 206.10 MB/s
BenchmarkCodeSliceBytes-4     	1000000000	        23.3 ns/op	 171.96 MB/s
BenchmarkSizeofUint8-4        	1000000000	         0.310 ns/op	3229.34 MB/s
BenchmarkSizeofUint16-4       	1000000000	         0.310 ns/op	6444.85 MB/s
BenchmarkSizeofUint32-4       	1000000000	         0.310 ns/op	12894.80 MB/s
BenchmarkSizeofUint64-4       	1000000000	         0.319 ns/op	25078.29 MB/s
BenchmarkSizeofInt-4          	1000000000	         0.312 ns/op	6411.63 MB/s
BenchmarkSizeofVarint-4       	1000000000	         0.310 ns/op	6442.28 MB/s
BenchmarkSizeofFloat32-4      	1000000000	         0.310 ns/op	12890.61 MB/s
BenchmarkSizeofFloat64-4      	1000000000	         0.309 ns/op	25892.49 MB/s
BenchmarkSizeofBool-4         	1000000000	         0.310 ns/op	3221.76 MB/s
BenchmarkSizeofString-4       	1000000000	         0.312 ns/op	3206.09 MB/s
BenchmarkSizeofBytes-4        	1000000000	         0.310 ns/op	3225.67 MB/s
BenchmarkSizeofSliceBytes-4   	1000000000	         3.11 ns/op	 642.69 MB/s
PASS
ok  	github.com/hslam/code	80.845s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Meng Huang)

### Authors
code was written by Meng Huang.
