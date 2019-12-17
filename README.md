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
	Int()
	Varint()
	Float32()
	Float64()
	Bool()
	String()
	Bytes()
	SliceBytes()
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
BenchmarkCodeInt-4                    	57345968	        20.4 ns/op	  98.23 MB/s
BenchmarkCodeIntWithBuffer-4          	151837402	         7.85 ns/op	 254.90 MB/s
BenchmarkCodeVarint-4                 	56025999	        20.9 ns/op	  95.81 MB/s
BenchmarkCodeVarintWithBuffer-4       	137225653	         8.53 ns/op	 234.48 MB/s
BenchmarkCodeFloat32-4                	44772171	        26.3 ns/op	 152.14 MB/s
BenchmarkCodeFloat32WithBuffer-4      	87854490	        13.7 ns/op	 291.17 MB/s
BenchmarkCodeFloat64-4                	29906814	        38.3 ns/op	 208.67 MB/s
BenchmarkCodeFloat64WithBuffer-4      	51124149	        23.8 ns/op	 336.18 MB/s
BenchmarkCodeBool-4                   	98387308	        12.0 ns/op	  83.58 MB/s
BenchmarkCodeBoolWithBuffer-4         	1000000000	         0.328 ns/op	3049.39 MB/s
BenchmarkCodeString-4                 	22178115	        53.2 ns/op	 206.68 MB/s
BenchmarkCodeStringWithBuffer-4       	36173110	        33.1 ns/op	 331.89 MB/s
BenchmarkCodeBytes-4                  	24055592	        50.3 ns/op	 218.52 MB/s
BenchmarkCodeBytesWithBuffer-4        	44093928	        27.2 ns/op	 403.97 MB/s
BenchmarkCodeSliceBytes-4             	 8689598	       132 ns/op	  45.49 MB/s
BenchmarkCodeSliceBytesWithBuffer-4   	 9467461	       114 ns/op	  52.58 MB/s
PASS
ok  	github.com/hslam/code	21.209s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Meng Huang)

### Authors
code was written by Meng Huang.
