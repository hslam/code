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
	data:=code.EncodeInt(buf,i)
	fmt.Printf("EncodeInt:%d to []byte:%v\n",i,data)
	v,n:=code.DecodeInt(data)
	fmt.Printf("DecodeInt:%d,length:%d\n",v,n)
}

func Varint()  {
	var buf =make([]byte,10)
	var i uint64=128
	data:=code.EncodeVarint(buf,i)
	fmt.Printf("EncodeVarint:%d to []byte:%v\n",i,data)
	v,n:=code.DecodeVarint(data)
	fmt.Printf("DecodeVarint:%d,length:%d\n",v,n)
}

func Float32()  {
	var buf =make([]byte,9)
	var i float32=3.14
	data:=code.EncodeFloat32(buf,i)
	fmt.Printf("EncodeFloat32:%.2f to []byte:%v\n",i,data)
	v,n:=code.DecodeFloat32(data)
	fmt.Printf("EncodeFloat32:%.2f,length:%d\n",v,n)
}

func Float64()  {
	var buf =make([]byte,9)
	var i float64=3.14
	data:=code.EncodeFloat64(buf,i)
	fmt.Printf("EncodeFloat64:%.2f to []byte:%v\n",i,data)
	v,n:=code.DecodeFloat64(data)
	fmt.Printf("DecodeFloat64:%.2f,length:%d\n",v,n)
}

func Bool()  {
	var buf =make([]byte,16)
	var i bool=true
	data:=code.EncodeBool(buf,i)
	fmt.Printf("EncodeBool:%t to []byte:%v\n",i,data)
	v,n:=code.DecodeBool(data)
	fmt.Printf("DecodeBool:%t,length:%d\n",v,n)
}

func String()  {
	var buf =make([]byte,16)
	var i string="Hello"
	data:=code.EncodeString(buf,i)
	fmt.Printf("EncodeString:%s to []byte:%v\n",i,data)
	v,n:=code.DecodeString(data)
	fmt.Printf("DecodeString:%s,length:%d\n",v,n)
}

func Bytes()  {
	var buf =make([]byte,16)
	var i []byte=[]byte{1,2}
	data:=code.EncodeBytes(buf,i)
	fmt.Printf("EncodeBytes:%v to []byte:%v\n",i,data)
	v,n:=code.DecodeBytes(data)
	fmt.Printf("DecodeBytes:%v,length:%d\n",v,n)
}

func SliceBytes()  {
	var buf =make([]byte,16)
	var i [][]byte=[][]byte{{1,2},{3}}
	data:=code.EncodeSliceBytes(buf,i)
	fmt.Printf("EncodeSliceBytes:%v to []byte:%v\n",i,data)
	v,n:=code.DecodeSliceBytes(data)
	fmt.Printf("DecodeSliceBytes:%v,length:%d\n",v,n)
}