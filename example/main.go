package main

import (
	"hslam.com/git/x/code"
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