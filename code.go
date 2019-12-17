package code

import (
	"unsafe"
)


const (
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

	mask7=-1 ^ (-1 << 7)
	mask8=-1 ^ (-1 << 8)

	msb7= 1<<7
)

func Encode(buf []byte,v interface{}) ([]byte, error) {
	return nil,nil
}

func Decode(data []byte, v interface{}) error {
	return nil
}
func EncodeInt(buf []byte,v uint64) []byte {
	var s []byte
	size:=SizeofInt(v)
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	s[0]=byte(size-1)
	for i:=1; i< size; i++ {
		s[i] = uint8(v & mask8)
		v >>= 8
	}
	return s
}

func DecodeInt(buf []byte) (v uint64, n int)  {
	size:=int(buf[0])
	for n = 1; n < size+1; n++ {
		b := buf[n]
		v |= uint64(b) << (uint(n-1)*8)
	}
	return v,n
}
func SizeofInt(v uint64) int {
	if v==0{
		return 1
	}else if v<v8{
		return 2
	}else if v < v16{
		return 3
	}else if v < v24{
		return 4
	}else if v < v32{
		return 5
	}else if v < v40{
		return 6
	}else if v < v48{
		return 7
	}else if v < v56{
		return 8
	}else{
		return 9
	}
}

func EncodeVarint(buf []byte,v uint64) []byte {
	var s []byte
	size:=SizeofVarint(v)
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

func DecodeVarint(d []byte) (v uint64, n int) {
	for i := 0; i < 10; i++ {
		if i >= len(d) {
			return 0, 0
		}
		b := d[i]
		v |= uint64(b) & mask7 << (uint(i)*7)
		if b & msb7 == 0 {
			return v, i+1
		}
	}
	return 0, 0
}

func SizeofVarint(v uint64) int {
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

func  EncodeFloat32(buf []byte,f float32) []byte {
	var s []byte
	size:=4
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	v:=*(*uint32)(unsafe.Pointer(&f))
	for i:=0;i<size ;i++  {
		s[i]=uint8(v>>uint8(i*8))
	}
	return s
}

func DecodeFloat32(d []byte) (f float32, n int ) {
	var v uint64
	size:=4
	for i := 0; i < size; i++ {
		if i >= len(d) {
			return 0,4
		}
		b := d[i]
		v |= uint64(b) & mask8 << (uint(i)*8)
	}
	return *(*float32)(unsafe.Pointer(&v)),4
}
func SizeofFloat32() int {
	return 4
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

func DecodeFloat64(d []byte) (f float64, n int ) {
	var v uint64
	for i := 0; i < 8; i++ {
		if i >= len(d) {
			return 0,8
		}
		b := d[i]
		v |= uint64(b) & mask8 << (uint(i)*8)
	}
	return *(*float64)(unsafe.Pointer(&v)),8
}
func SizeofFloat64() int {
	return 8
}
func EncodeString(buf []byte,v string) []byte {
	var s []byte
	length:=len(v)
	b:=EncodeVarint(nil,uint64(length))
	size:=len(b)+length
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	copy(s[:len(b)],b)
	copy(s[len(b):],v)
	return s
}

func DecodeString(d []byte) (s string, n int ) {
	v,n:=DecodeVarint(d)
	s=string(d[n:v+1])
	return s,n+int(v)
}
func SizeofString(v string) int {
	return SizeofVarint(uint64(len(v)))+len(v)
}
func EncodeBool(buf []byte,v bool) []byte {
	var s []byte
	size:=1
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	if !v{
		s[0]=0
	}else{
		s[0]=1
	}
	return s
}

func DecodeBool(d []byte) (s bool, n int) {
	if len(d)==0{
		return false,0
	}
	if d[0]==0{
		return false,1
	}
	return true,1
}
func SizeofBool() int {
	return 1
}
func EncodeBytes(buf []byte,v []byte) []byte {
	var s []byte
	length:=len(v)
	b:=EncodeVarint(nil,uint64(length))
	size:=len(b)+length
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	copy(s[:len(b)],b)
	copy(s[len(b):],v)
	return s
}

func DecodeBytes(d []byte) (s []byte, n int) {
	v,n:=DecodeVarint(d)
	s=d[n:v+1]
	return s ,n+int(v)
}
func SizeofBytes(v []byte) int {
	return SizeofVarint(uint64(len(v)))+len(v)
}
func EncodeSliceBytes(buf []byte,d [][]byte) []byte {
	var s []byte
	var size int
	for _,v:=range d{
		l:=len(v)
		s:=SizeofVarint(uint64(l))
		size+=s+l
	}
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	var offset int
	var tmpbuf =make([]byte,10)
	for _,v:=range d{
		l:=len(v)
		b:=EncodeVarint(tmpbuf,uint64(l))
		copy(s[offset:offset+len(b)],b)
		copy(s[offset+len(b):],v)
		offset+=len(b)+l
	}
	return s
}

func DecodeSliceBytes(d []byte) (s [][]byte, n int) {
	var l int
	var offset int
	for offset<len(d){
		v,n:=DecodeVarint(d[offset:])
		offset+=n+int(v)
		l++
	}
	s=make([][]byte,l)
	offset=0
	var i int
	for offset<len(d){
		v,n:=DecodeVarint(d[offset:])
		s[i]=d[offset+n:offset+n+int(v)]
		i++
		offset+=n+int(v)
	}
	return s,offset
}

func SizeofSliceBytes(d [][]byte) int {
	var size int
	for _,v:=range d{
		l:=len(v)
		s:=SizeofVarint(uint64(l))
		size+=s+l
	}
	return size
}