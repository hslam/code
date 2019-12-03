package code

import (
	"unsafe"
)

type CodeType uint8

const (
	Null 		CodeType = 0
	Int 		CodeType = 1
	Float 		CodeType = 2
	False		CodeType = 3
	True		CodeType = 4
	String 		CodeType = 5
	Bytes 		CodeType = 6

	v7 =1<<7
	v8 =1<<8
	v14=1<<14
	v16=1<<16
	v21=1<<21
	v24=1<<24
	v28=1<<28
	v32=1<<32
	v35=1<<35
	v40=1<<40
	v42=1<<42
	v48=1<<48
	v49=1<<49
	v56=1<<56
	v63=1<<63
	v64=1<<64

	mask4=-1 ^ (-1 << 4)
	mask7=-1 ^ (-1 << 7)
	mask8=-1 ^ (-1 << 8)

	msb7= 1<<7



)

func Encode(v interface{}) ([]byte, error) {
	//c,n,b:=InterfaceType(v)
	return nil,nil
}

func Decode(data []byte, v interface{}) error {
	return nil
}

func InterfaceType(i interface{}) (c CodeType,n int,b uint8){
	switch i.(type) {
	case string:
		n,b=StringType(i.(string))
		return String,n,b
	case uint,uint8,uint16,uint32,uint64,int,int8,int16,int32,int64:
		n,b=IntType(i.(uint64))
		return Int,n,b
	case float32:
		n,b=FloatType(4)
		return Float,n,b
	case float64:
		n,b=FloatType(8)
		return Float,n,b
	case bool:
		if i.(bool)==true{
			n,b=TrueType()
			return True,n,b
		}else {
			n,b=FalseType()
			return False,n,b
		}
	default:
		return Null,0,0
	}
}
func Type(b byte) (c CodeType,n int){
	n= int(b & mask4)
	c=CodeType(b>>4)
	return c,n
}
func StringType(i string) (n int,b uint8){
	length:=len(i)
	n=SizeInt(uint64(length))
	b=uint8(String)<<4 |uint8(n)
	return n,b
}
func IntType(i uint64) (n int,b uint8){
	n=SizeInt(i)
	b=uint8(Int)<<4 |uint8(n)
	return n,b
}
func FloatType(l int) (n int,b uint8){
	n=l
	b=uint8(Float)<<4 |uint8(n)
	return n,b
}
func TrueType() (n int,b uint8){
	b=uint8(True)<<4
	return n,b
}
func FalseType() (n int,b uint8){
	b=uint8(False)<<4
	return n,b
}
func EncodeInt(buf []byte,v uint64) []byte {
	var s []byte
	size,b:=IntType(v)
	size+=1
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	s[0]=b
	for i:=0; i< size; i++ {
		s[i] = uint8(v & mask8)
		v >>= 8
	}
	return s
}

func DecodeInt(buf []byte) (v uint64, n int) {
	for n = 0; n < 8; n++ {
		if n >= len(buf) {
			return v,n
		}
		b := buf[n]
		v |= uint64(b) << (uint(n)*8)
	}
	return v, n
}
func SizeInt(v uint64) int {
	if v==0{
		return 0
	}else if v<v8{
		return 1
	}else if v < v16{
		return 2
	}else if v < v24{
		return 3
	}else if v < v32{
		return 4
	}else if v < v40{
		return 5
	}else if v < v48{
		return 6
	}else if v < v56{
		return 7
	}else{
		return 8
	}
}

func EncodeVarint(buf []byte,v uint64) []byte {
	var s []byte
	size:=SizeVarint(v)
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	for i:=0; i< size-1; i++ {
		s[i] = uint8(v & mask7 | msb7)
		v >>= 7
	}
	s[size-1] = uint8(v)
	return s
}

func DecodeVarint(buf []byte) (v uint64, n int) {
	for i := 0; i < 10; i++ {
		if i >= len(buf) {
			return 0, 0
		}
		b := buf[i]
		v |= uint64(b) & mask7 << (uint(i)*7)
		if b & msb7 == 0 {
			return v, i
		}
	}
	return 0, 0
}

func SizeVarint(v uint64) int {
	if v<v7{
		return 1
	}else if v < v14{
		return 2
	}else if v < v21{
		return 3
	}else if v < v28{
		return 4
	}else if v < v35{
		return 5
	}else if v < v42{
		return 6
	}else if v < v49{
		return 7
	}else if v < v56{
		return 8
	}else if v < v63 {
		return 9
	}else {
		return 10
	}
}

func EncodeFloat64(buf []byte,f float64) []byte {
	v:=*(*uint64)(unsafe.Pointer(&f))
	var s []byte
	size:=8
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	for i:=0;i<8 ;i++  {
		s[i]=uint8(v>>uint8(i*8))
	}
	return s
}

func DecodeFloat64(buf []byte) (f float64) {
	var v uint64
	for i := 0; i < 8; i++ {
		if i >= len(buf) {
			return 0
		}
		b := buf[i]
		v |= uint64(b) & mask8 << (uint(i)*8)
	}
	return *(*float64)(unsafe.Pointer(&v))
}
func  EncodeFloat32(buf []byte,f float32) []byte {
	v:=*(*uint32)(unsafe.Pointer(&f))
	var s []byte
	size:=4
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	for i:=0;i<4 ;i++  {
		s[i]=uint8(v>>uint8(i*8))
	}
	return s
}

func DecodeFloat32(buf []byte) (f float32) {
	var v uint64
	for i := 0; i < 4; i++ {
		if i >= len(buf) {
			return 0
		}
		b := buf[i]
		v |= uint64(b) & mask8 << (uint(i)*8)
	}
	return *(*float32)(unsafe.Pointer(&v))
}