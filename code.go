package code

import (
	"unsafe"
)
type Level int

const (
	Speed		Level= 1
	Short		Level= 9
	Default		Level= -1

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
func EncodeUint8(buf []byte,v uint8) []byte {
	var s []byte
	if cap(buf) >= 1 {
		s = buf[:1]
	} else {
		s = make([]byte, 1)
	}
	s[0] = byte(v & mask8)
	return s
}

func DecodeUint8(buf []byte) (v uint8, n int)  {
	v |= uint8(buf[0])
	return v,1
}
func EncodeUint16(buf []byte,v uint16) []byte {
	var s []byte
	if cap(buf) >= 2 {
		s = buf[:2]
	} else {
		s = make([]byte, 2)
	}
	s[0] = byte(v & mask8)
	v >>= 8
	s[1] = byte(v & mask8)
	return s
}

func DecodeUint16(buf []byte) (v uint16, n int)  {
	v |= uint16(buf[0])
	v |= uint16(buf[1]) << 8
	return v,2
}
func EncodeUint32(buf []byte,v uint32) []byte {
	var s []byte
	if cap(buf) >= 4 {
		s = buf[:4]
	} else {
		s = make([]byte, 4)
	}
	s[0] = byte(v & mask8)
	v >>= 8
	s[1] = byte(v & mask8)
	v >>= 8
	s[2] = byte(v & mask8)
	v >>= 8
	s[3] = byte(v & mask8)
	return s
}

func DecodeUint32(buf []byte) (v uint32, n int)  {
	v |= uint32(buf[0])
	v |= uint32(buf[1]) << 8
	v |= uint32(buf[2]) << 16
	v |= uint32(buf[3]) << 24
	return v,4
}
func EncodeUint64(buf []byte,v uint64) []byte {
	var s []byte
	if cap(buf) >= 8 {
		s = buf[:8]
	} else {
		s = make([]byte, 8)
	}
	s[0] = byte(v & mask8)
	v >>= 8
	s[1] = byte(v & mask8)
	v >>= 8
	s[2] = byte(v & mask8)
	v >>= 8
	s[3] = byte(v & mask8)
	v >>= 8
	s[4] = byte(v & mask8)
	v >>= 8
	s[5] = byte(v & mask8)
	v >>= 8
	s[6] = byte(v & mask8)
	v >>= 8
	s[7] = byte(v & mask8)
	return s
}

func DecodeUint64(buf []byte) (v uint64, n int)  {
	v |= uint64(buf[0])
	v |= uint64(buf[1]) << 8
	v |= uint64(buf[2]) << 16
	v |= uint64(buf[3]) << 24
	v |= uint64(buf[4]) << 32
	v |= uint64(buf[5]) << 40
	v |= uint64(buf[6]) << 48
	v |= uint64(buf[7]) << 56
	return v,8
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
		s[i] = byte(v & mask8)
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
	switch {
	case v ==0:
		return 1
	case v < v8:
		return 2
	case v < v16:
		return 3
	case v < v24:
		return 4
	case v < v32:
		return 5
	case v < v40:
		return 6
	case v < v48:
		return 7
	case v < v56:
		return 8
	default:
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
		s[i] = byte(v & mask7 | msb7)
		v >>= 7
	}
	s[size-1] = byte(v)
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
	switch {
	case v < v7:
		return 1
	case v < v14:
		return 2
	case v < v21:
		return 3
	case v < v28:
		return 4
	case v < v35:
		return 5
	case v < v42:
		return 6
	case v < v49:
		return 7
	case v < v56:
		return 8
	case v < v63:
		return 9
	default:
		return 10
	}
}

func  EncodeFloat32(buf []byte,f float32) []byte {
	var s []byte
	if cap(buf) >= 4 {
		s = buf[:4]
	} else {
		s = make([]byte, 4)
	}
	v:=*(*uint32)(unsafe.Pointer(&f))
	s[0]=uint8(v)
	s[1]=uint8(v>>8)
	s[2]=uint8(v>>16)
	s[3]=uint8(v>>24)
	return s
}

func DecodeFloat32(d []byte) (f float32, n int ) {
	var v uint64
	v |= uint64(d[0]) & mask8
	v |= uint64(d[1]) & mask8<<8
	v |= uint64(d[2]) & mask8<<16
	v |= uint64(d[3]) & mask8<<24
	return *(*float32)(unsafe.Pointer(&v)),4
}
func SizeofFloat32() int {
	return 4
}
func EncodeFloat64(buf []byte,f float64) []byte {
	v:=*(*uint64)(unsafe.Pointer(&f))
	var s []byte
	if cap(buf) >= 8 {
		s = buf[:8]
	} else {
		s = make([]byte, 8)
	}
	s[0]=uint8(v)
	s[1]=uint8(v>>8)
	s[2]=uint8(v>>16)
	s[3]=uint8(v>>24)
	s[4]=uint8(v>>32)
	s[5]=uint8(v>>40)
	s[6]=uint8(v>>48)
	s[7]=uint8(v>>56)
	return s
}

func DecodeFloat64(d []byte) (f float64, n int ) {
	var v uint64
	v |= uint64(d[0]) & mask8
	v |= uint64(d[1]) & mask8<<8
	v |= uint64(d[2]) & mask8<<16
	v |= uint64(d[3]) & mask8<<24
	v |= uint64(d[4]) & mask8<<32
	v |= uint64(d[5]) & mask8<<40
	v |= uint64(d[6]) & mask8<<48
	v |= uint64(d[7]) & mask8<<56
	return *(*float64)(unsafe.Pointer(&v)),8
}
func SizeofFloat64() int {
	return 8
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

func EncodeString(buf []byte,v string) []byte {
	var s []byte
	length:=len(v)
	length_size:=SizeofVarint(uint64(length))
	var size int =length_size+length
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	l:=length
	for i:=0; i< length_size-1; i++ {
		s[i] = byte(l & mask7 | msb7)
		l >>= 7
	}
	s[length_size-1] = byte(l)
	copy(s[length_size:],v)
	return s
}

func DecodeString(d []byte) (s string, n int ) {
	v,n:=DecodeVarint(d)
	b:=d[n:v+1]
	return *(*string)(unsafe.Pointer(&b)),n+int(v)
}
func SizeofString(v string) int {
	return SizeofVarint(uint64(len(v)))+len(v)
}
func EncodeBytes(buf []byte,v []byte) []byte {
	var s []byte
	length:=len(v)
	length_size:=SizeofVarint(uint64(length))
	var size int =length_size+length
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	l:=length
	for i:=0; i< length_size-1; i++ {
		s[i] = byte(l & mask7 | msb7)
		l >>= 7
	}
	s[length_size-1] = byte(l)
	copy(s[length_size:],v)
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
	for _,v:=range d{
		length:=len(v)
		length_size:=SizeofVarint(uint64(length))
		l:=length
		for i:=offset; i< offset+length_size-1; i++ {
			s[i] = byte(l & mask7 | msb7)
			l >>= 7
		}
		s[offset+length_size-1] = byte(l)
		copy(s[offset+length_size:],v)
		offset+=length+length_size
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