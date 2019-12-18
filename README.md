# code
A code library written in golang for encoding and decoding.

## Feature
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
	fmt.Printf("SizeofUint8:%d sizeof:%d\n",i,1)
	data:=code.EncodeUint8(buf,i)
	fmt.Printf("EncodeUint8:%d to []byte:%v\n",i,data)
	v,n:=code.DecodeUint8(data)
	fmt.Printf("DecodeUint8:%d,length:%d\n",v,n)
}
func Uint16()  {
	var buf =make([]byte,4)
	var i uint16=128
	fmt.Printf("SizeofUint16:%d sizeof:%d\n",i,2)
	data:=code.EncodeUint16(buf,i)
	fmt.Printf("EncodeUint16:%d to []byte:%v\n",i,data)
	v,n:=code.DecodeUint16(data)
	fmt.Printf("DecodeUint16:%d,length:%d\n",v,n)
}
func Uint32()  {
	var buf =make([]byte,4)
	var i uint32=128
	fmt.Printf("SizeofUint32:%d sizeof:%d\n",i,4)
	data:=code.EncodeUint32(buf,i)
	fmt.Printf("EncodeUint32:%d to []byte:%v\n",i,data)
	v,n:=code.DecodeUint32(data)
	fmt.Printf("DecodeUint32:%d,length:%d\n",v,n)
}
func Uint64()  {
	var buf =make([]byte,8)
	var i uint64=128
	fmt.Printf("SizeofUint64:%d sizeof:%d\n",i,8)
	data:=code.EncodeUint64(buf,i)
	fmt.Printf("EncodeUint64:%d to []byte:%v\n",i,data)
	v,n:=code.DecodeUint64(data)
	fmt.Printf("DecodeUint64:%d,length:%d\n",v,n)
}
func Int()  {
	var buf =make([]byte,9)
	var i uint64=128
	size:=code.SizeofInt(i)
	fmt.Printf("SizeofInt:%d sizeof:%d\n",i,size)
	data:=code.EncodeInt(buf,i)
	fmt.Printf("EncodeInt:%d to []byte:%v\n",i,data)
	v,n:=code.DecodeInt(data)
	fmt.Printf("DecodeInt:%d,length:%d\n",v,n)
}
func Varint()  {
	var buf =make([]byte,10)
	var i uint64=128
	size:=code.SizeofVarint(i)
	fmt.Printf("SizeofVarint:%d sizeof:%d\n",i,size)
	data:=code.EncodeVarint(buf,i)
	fmt.Printf("EncodeVarint:%d to []byte:%v\n",i,data)
	v,n:=code.DecodeVarint(data)
	fmt.Printf("DecodeVarint:%d,length:%d\n",v,n)
}
func Float32()  {
	var buf =make([]byte,9)
	var i float32=3.14
	size:=code.SizeofFloat32()
	fmt.Printf("SizeofFloat32:%.2f sizeof:%d\n",i,size)
	data:=code.EncodeFloat32(buf,i)
	fmt.Printf("EncodeFloat32:%.2f to []byte:%v\n",i,data)
	v,n:=code.DecodeFloat32(data)
	fmt.Printf("EncodeFloat32:%.2f,length:%d\n",v,n)
}
func Float64()  {
	var buf =make([]byte,9)
	var i float64=3.14
	size:=code.SizeofFloat64()
	fmt.Printf("SizeofFloat64:%.2f sizeof:%d\n",i,size)
	data:=code.EncodeFloat64(buf,i)
	fmt.Printf("EncodeFloat64:%.2f to []byte:%v\n",i,data)
	v,n:=code.DecodeFloat64(data)
	fmt.Printf("DecodeFloat64:%.2f,length:%d\n",v,n)
}
func Bool()  {
	var buf =make([]byte,16)
	var i bool=true
	size:=code.SizeofBool()
	fmt.Printf("SizeofBool:%t sizeof:%d\n",i,size)
	data:=code.EncodeBool(buf,i)
	fmt.Printf("EncodeBool:%t to []byte:%v\n",i,data)
	v,n:=code.DecodeBool(data)
	fmt.Printf("DecodeBool:%t,length:%d\n",v,n)
}
func String()  {
	var buf =make([]byte,16)
	var i string="Hello"
	size:=code.SizeofString(i)
	fmt.Printf("SizeofString:%s sizeof:%d\n",i,size)
	data:=code.EncodeString(buf,i)
	fmt.Printf("EncodeString:%s to []byte:%v\n",i,data)
	v,n:=code.DecodeString(data)
	fmt.Printf("DecodeString:%s,length:%d\n",v,n)
}
func Bytes()  {
	var buf =make([]byte,16)
	var i []byte=[]byte{1,2}
	size:=code.SizeofBytes(i)
	fmt.Printf("SizeofBytes:%v sizeof:%d\n",i,size)
	data:=code.EncodeBytes(buf,i)
	fmt.Printf("EncodeBytes:%v to []byte:%v\n",i,data)
	v,n:=code.DecodeBytes(data)
	fmt.Printf("DecodeBytes:%v,length:%d\n",v,n)
}
func SliceBytes()  {
	var buf =make([]byte,16)
	var i [][]byte=[][]byte{{1,2},{3}}
	size:=code.SizeofSliceBytes(i)
	fmt.Printf("SizeofSliceBytes:%v sizeof:%d\n",i,size)
	data:=code.EncodeSliceBytes(buf,i)
	fmt.Printf("EncodeSliceBytes:%v to []byte:%v\n",i,data)
	v,n:=code.DecodeSliceBytes(data)
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
go test -v -run="none" -bench=. -benchtime=1s
```
goos: darwin
goarch: amd64
pkg: github.com/hslam/code
BenchmarkCodeUint8-4        	1000000000	         0.324 ns/op	3090.14 MB/s
BenchmarkCodeUint16-4       	1000000000	         0.317 ns/op	6310.61 MB/s
BenchmarkCodeUint32-4       	1000000000	         0.318 ns/op	12568.89 MB/s
BenchmarkCodeUint64-4       	316729856	         3.75 ns/op	2131.79 MB/s
BenchmarkCodeInt-4          	145029232	         8.25 ns/op	 242.31 MB/s
BenchmarkCodeVarint-4       	130683343	         9.24 ns/op	 216.50 MB/s
BenchmarkCodeFloat32-4      	1000000000	         0.321 ns/op	12467.47 MB/s
BenchmarkCodeFloat64-4      	161058280	         7.54 ns/op	1060.87 MB/s
BenchmarkCodeBool-4         	1000000000	         0.323 ns/op	3093.85 MB/s
BenchmarkCodeString-4       	93031238	        12.3 ns/op	 162.07 MB/s
BenchmarkCodeBytes-4        	100000000	        10.6 ns/op	1039.61 MB/s
BenchmarkCodeSliceBytes-4   	11185539	       105 ns/op	  57.40 MB/s
PASS
ok  	github.com/hslam/code	13.009s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Meng Huang)

### Authors
code was written by Meng Huang.
