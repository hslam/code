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
func EncodeUint8(buf []byte,v *uint8) []byte {
	var s []byte
	if cap(buf) >= 1 {
		s = buf[:1]
	} else {
		s = make([]byte, 1)
	}
	s[0]=uint8(*v)
	return s
}

func DecodeUint8(buf []byte,v *uint8) ( n int)  {
	*v = uint8(buf[0])
	return 1
}
func SizeofUint8() int {
	return 1
}
func EncodeUint16(buf []byte,v *uint16) []byte {
	var s []byte
	if cap(buf) >= 2 {
		s = buf[:2]
	} else {
		s = make([]byte, 2)
	}
	var t = *v
	s[0]=uint8(t)
	s[1]=uint8(t>>8)
	return s
}

func DecodeUint16(buf []byte,v *uint16) ( n int)  {
	var t uint16
	t = uint16(buf[0])
	t |= uint16(buf[1]) << 8
	*v=t
	return 2
}
func SizeofUint16() int {
	return 2
}
func EncodeUint32(buf []byte,v *uint32) []byte {
	var s []byte
	if cap(buf) >= 4 {
		s = buf[:4]
	} else {
		s = make([]byte, 4)
	}
	var t = *v
	s[0]=uint8(t)
	s[1]=uint8(t>>8)
	s[2]=uint8(t>>16)
	s[3]=uint8(t>>24)
	return s
}

func DecodeUint32(buf []byte,v *uint32) ( n int)  {
	var t uint32
	t = uint32(buf[0])
	t |= uint32(buf[1]) << 8
	t |= uint32(buf[2]) << 16
	t |= uint32(buf[3]) << 24
	*v=t
	return 4
}
func SizeofUint32() int {
	return 4
}
func EncodeUint64(buf []byte,v *uint64) []byte {
	var s []byte
	if cap(buf) >= 8 {
		s = buf[:8]
	} else {
		s = make([]byte, 8)
	}
	var t =*v
	s[0]=uint8(t)
	s[1]=uint8(t>>8)
	s[2]=uint8(t>>16)
	s[3]=uint8(t>>24)
	s[4]=uint8(t>>32)
	s[5]=uint8(t>>40)
	s[6]=uint8(t>>48)
	s[7]=uint8(t>>56)
	return s
}

func DecodeUint64(buf []byte,v *uint64) ( n int)  {
	var t uint64
	t = uint64(buf[0])
	t |= uint64(buf[1]) << 8
	t |= uint64(buf[2]) << 16
	t |= uint64(buf[3]) << 24
	t |= uint64(buf[4]) << 32
	t |= uint64(buf[5]) << 40
	t |= uint64(buf[6]) << 48
	t |= uint64(buf[7]) << 56
	*v=t
	return 8
}
func SizeofUint64() int {
	return 8
}
func EncodeInt(buf []byte,v *uint64) []byte {
	var s []byte
	var t = *v

	var size =SizeofInt(v)
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	s[0]=byte(size-1)
	for i:=1;i<size;i++ {
		s[i] = byte(t)
		t>>= 8
	}
	return s
}

func DecodeInt(buf []byte,v *uint64) ( n int)  {
	size:=int(buf[0])
	var t uint64

	//for n = 1; n < size+1; n++ {
	//	t |= uint64(buf[n]) << (uint(n-1)*8)
	//}

	if size==0{
		n=1
		goto done
	}
	t |= uint64(buf[1])
	if size==1{
		n=2
		goto done
	}
	t |= uint64(buf[2])<<8
	if size==2{
		n=3
		goto done
	}
	t |= uint64(buf[3])<<16
	if size==3{
		n=4
		goto done
	}
	t |= uint64(buf[4])<<24
	if size==4{
		n=5
		goto done
	}
	t |= uint64(buf[5])<<32
	if size==5{
		n=6
		goto done
	}
	t |= uint64(buf[6])<<40
	if size==6{
		n=7
		goto done
	}
	t |= uint64(buf[7])<<48
	if size==7{
		n=8
		goto done
	}
	t |= uint64(buf[8])<<56
	n=9
	goto done
done:
	*v=t
	return n
}

func SizeofInt(v *uint64) int {
	var t = *v
	switch {
	case t ==0:
		return 1
	case t < v8:
		return 2
	case t < v16:
		return 3
	case t < v24:
		return 4
	case t < v32:
		return 5
	case t < v40:
		return 6
	case t < v48:
		return 7
	case t < v56:
		return 8
	default:
		return 9
	}
}

func EncodeVarint(buf []byte,v *uint64) []byte {
	var s []byte
	var t = *v
	var size =SizeofVarint(v)
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	for i := 0;i<size-1;i++ {
		s[i] = byte(t) | msb7
		t >>= 7
	}
	s[size-1] = byte(t)
	return s
}

func DecodeVarint(d []byte,v *uint64) (n int) {
	var t uint64

	//shift := uint8(7)
	//t = uint64(d[n]&mask7)
	//for d[n]&msb7 == msb7 {
	//	n++
	//	t |= uint64(d[n]&mask7) << shift
	//	shift += 7
	//}

	t = uint64(d[n]&mask7)
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 7
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 14
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 21
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 28
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 35
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 42
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 49
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 56
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 63
	goto done
done:
	n++
	*v=t
	return n
}


func SizeofVarint(v *uint64) int {
	var t = *v
	switch {
	case t < v7:
		return 1
	case t < v14:
		return 2
	case t < v21:
		return 3
	case t < v28:
		return 4
	case t < v35:
		return 5
	case t < v42:
		return 6
	case t < v49:
		return 7
	case t < v56:
		return 8
	case t < v63:
		return 9
	default:
		return 10
	}
}

func  EncodeFloat32(buf []byte,f *float32) []byte {
	var s []byte
	if cap(buf) >= 4 {
		s = buf[:4]
	} else {
		s = make([]byte, 4)
	}
	v:=*(*uint32)(unsafe.Pointer(f))
	s[0]=uint8(v)
	s[1]=uint8(v>>8)
	s[2]=uint8(v>>16)
	s[3]=uint8(v>>24)
	return s
}

func DecodeFloat32(d []byte,f *float32) ( n int ) {
	var v uint64
	v |= uint64(d[0])
	v |= uint64(d[1]) <<8
	v |= uint64(d[2]) <<16
	v |= uint64(d[3]) <<24
	*f=*(*float32)(unsafe.Pointer(&v))
	return 4
}
func SizeofFloat32() int {
	return 4
}
func EncodeFloat64(buf []byte,f *float64) []byte {
	v:=*(*uint64)(unsafe.Pointer(f))
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

func DecodeFloat64(d []byte,f *float64) (n int ) {
	var v uint64
	v |= uint64(d[0])
	v |= uint64(d[1]) <<8
	v |= uint64(d[2]) <<16
	v |= uint64(d[3]) <<24
	v |= uint64(d[4]) <<32
	v |= uint64(d[5]) <<40
	v |= uint64(d[6]) <<48
	v |= uint64(d[7]) <<56
	*f=*(*float64)(unsafe.Pointer(&v))
	return 8
}
func SizeofFloat64() int {
	return 8
}

func EncodeBool(buf []byte,v *bool) []byte {
	var s []byte
	size:=1
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	if !*v{
		s[0]=0
	}else{
		s[0]=1
	}
	return s
}

func DecodeBool(d []byte,s *bool) ( n int) {
	if len(d)==0{
		*s=false
		return 0
	}
	if d[0]==0{
		*s=false
		return 1
	}
	*s=true
	return 1
}
func SizeofBool() int {
	return 1
}

func EncodeString(buf []byte,v *string) []byte {
	var s []byte
	length:=uint64(len(*v))
	length_size:=SizeofVarint(&length)
	var size int =length_size+int(length)
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	t:=length
	for i := 0;i<length_size-1;i++ {
		s[i] = byte(t) | msb7
		t >>= 7
	}
	s[length_size-1] = byte(t)
	copy(s[length_size:],*v)
	return s
}

func DecodeString(d []byte,s *string) ( n int ) {
	var t uint64
	t = uint64(d[n]&mask7)
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 7
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 14
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 21
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 28
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 35
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 42
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 49
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 56
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 63
	goto done
done:
	n++

	b:=d[n:n+int(t)]
	*s=*(*string)(unsafe.Pointer(&b))
	return n+int(t)
}
func SizeofString(v *string) int {
	length:=uint64(len(*v))
	return SizeofVarint(&length)+len(*v)
}
func EncodeBytes(buf []byte,v *[]byte) []byte {
	var s []byte
	length:=uint64(len(*v))
	length_size:=SizeofVarint(&length)
	var size int =length_size+int(length)
	if cap(buf) >= size {
		s = buf[:size]
	} else {
		s = make([]byte, size)
	}
	t:=length
	for i := 0;i<length_size-1;i++ {
		s[i] = byte(t) | msb7
		t >>= 7
	}
	s[length_size-1] = byte(t)
	copy(s[length_size:],*v)
	return s
}
func DecodeBytes(d []byte,s *[]byte) ( n int) {
	var t uint64
	t = uint64(d[n]&mask7)
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 7
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 14
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 21
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 28
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 35
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 42
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 49
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 56
	if d[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(d[n]&mask7) << 63
	goto done
done:
	n++
	*s=d[n:n+int(t)]
	return n+int(t)
}

func SizeofBytes(v *[]byte) int {
	length:=uint64(len(*v))
	return SizeofVarint(&length)+len(*v)
}
func EncodeSliceBytes(buf []byte,d *[][]byte) []byte {
	var s []byte
	var offset int
	var size int
	for _,v:=range *d{
		length:=uint64(len(v))
		length_size:=SizeofVarint(&length)
		size+=length_size+int(length)
		t:=length
		for i := 0;i<length_size-1;i++ {
			s[offset+i] = byte(t) | msb7
			t >>= 7
		}
		if cap(buf) >= size {
			s = buf[:size]
			s[offset+length_size-1] = byte(t)
			copy(s[offset+length_size:],v)
		} else {
			s=append(s,byte(t))
			s=append(s,v...)
		}
		offset+=int(length)+length_size
	}
	return s
}

func DecodeSliceBytes(d []byte,s *[][]byte) (n int) {
	var l int
	var offset int
	var i int
	for offset<len(d){
		buf:=d[offset:]
		j:=0
		var t uint64
		t = uint64(buf[j]&mask7)
		if buf[j]&msb7 == 0 {
			goto done
		}
		j++
		t |= uint64(buf[j]&mask7) << 7
		if buf[j]&msb7 == 0 {
			goto done
		}
		j++
		t |= uint64(buf[j]&mask7) << 14
		if buf[j]&msb7 == 0 {
			goto done
		}
		j++
		t |= uint64(buf[j]&mask7) << 21
		if buf[j]&msb7 == 0 {
			goto done
		}
		j++
		t |= uint64(buf[j]&mask7) << 28
		if buf[j]&msb7 == 0 {
			goto done
		}
		j++
		t |= uint64(buf[j]&mask7) << 35
		if buf[j]&msb7 == 0 {
			goto done
		}
		j++
		t |= uint64(buf[j]&mask7) << 42
		if buf[j]&msb7 == 0 {
			goto done
		}
		j++
		t |= uint64(buf[j]&mask7) << 49
		if buf[j]&msb7 == 0 {
			goto done
		}
		j++
		t |= uint64(buf[j]&mask7) << 56
		if buf[j]&msb7 == 0 {
			goto done
		}
		j++
		t |= uint64(buf[j]&mask7) << 63
		goto done
	done:
		j++

		n=j
		l++

		b:=d[offset+n:offset+n+int(t)]
		if cap(*s) >= l {
			*s = (*s)[:l]
			(*s)[i]=b
		} else {
			*s=append(*s,b)
		}
		i++
		offset+=n+int(t)
	}
	return offset
}

func SizeofSliceBytes(d *[][]byte) int {
	var size int
	for _,v:=range *d{
		length:=uint64(len(v))
		s:=SizeofVarint(&length)
		size+=s+int(length)
	}
	return size
}