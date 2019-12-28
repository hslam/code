package code

import (
	"unsafe"
)

const (
	v7 =1<<7
	v14=1<<14
	v21=1<<21
	v28=1<<28
	v35=1<<35
	v42=1<<42
	v49=1<<49
	v56=1<<56
	v63=1<<63
	mask7=-1 ^ (-1 << 7)
	msb7= 1<<7
)

func Encode(buf []byte,v interface{}) ([]byte, error) {
	return nil,nil
}

func Decode(data []byte, v interface{}) error {
	return nil
}

func EncodeUint8(buf []byte,v uint8) uint64 {
	if cap(buf) >= 1 {
		buf = buf[:1]
	} else {
		buf = make([]byte, 1)
	}
	buf[0]=uint8(v)
	return 1
}

func DecodeUint8(buf []byte,v *uint8) uint64 {
	*v = uint8(buf[0])
	return 1
}

func SizeofUint8() uint64 {
	return 1
}

func EncodeUint16(buf []byte,v uint16) uint64 {
	if cap(buf) >= 2 {
		buf = buf[:2]
	} else {
		buf = make([]byte, 2)
	}
	var t = v
	buf[0]=uint8(t)
	buf[1]=uint8(t>>8)
	return 2
}

func DecodeUint16(buf []byte,v *uint16) uint64 {
	var t uint16
	t = uint16(buf[0])
	t |= uint16(buf[1]) << 8
	*v=t
	return 2
}

func SizeofUint16() uint64 {
	return 2
}

func EncodeUint32(buf []byte,v uint32) uint64 {
	if cap(buf) >= 4 {
		buf = buf[:4]
	} else {
		buf = make([]byte, 4)
	}
	var t = v
	buf[0]=uint8(t)
	buf[1]=uint8(t>>8)
	buf[2]=uint8(t>>16)
	buf[3]=uint8(t>>24)
	return 4
}

func DecodeUint32(buf []byte,v *uint32) uint64 {
	var t uint32
	t = uint32(buf[0])
	t |= uint32(buf[1]) << 8
	t |= uint32(buf[2]) << 16
	t |= uint32(buf[3]) << 24
	*v=t
	return 4
}

func SizeofUint32() uint64 {
	return 4
}

func EncodeUint64(buf []byte,v uint64) uint64 {
	if cap(buf) >= 8 {
		buf = buf[:8]
	} else {
		buf = make([]byte, 8)
	}
	var t =v
	buf[0]=uint8(t)
	buf[1]=uint8(t>>8)
	buf[2]=uint8(t>>16)
	buf[3]=uint8(t>>24)
	buf[4]=uint8(t>>32)
	buf[5]=uint8(t>>40)
	buf[6]=uint8(t>>48)
	buf[7]=uint8(t>>56)
	return 8
}

func DecodeUint64(buf []byte,v *uint64) uint64 {
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

func SizeofUint64() uint64 {
	return 8
}

func EncodeVarint(buf []byte,v uint64) uint64{
	var t = v
	var size =SizeofVarint(v)
	if uint64(cap(buf)) >= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	for i := uint64(0);i<size-1;i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[size-1] = byte(t)
	return size
}

func DecodeVarint(d []byte,v *uint64)uint64 {
	var t uint64
	var n uint64

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


func SizeofVarint(v uint64) uint64 {
	var t = v
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

func  EncodeFloat32(buf []byte,f float32) uint64 {
	if cap(buf) >= 4 {
		buf = buf[:4]
	} else {
		buf = make([]byte, 4)
	}
	v:=*(*uint32)(unsafe.Pointer(&f))
	buf[0]=uint8(v)
	buf[1]=uint8(v>>8)
	buf[2]=uint8(v>>16)
	buf[3]=uint8(v>>24)
	return 4
}

func DecodeFloat32(d []byte,f *float32) uint64 {
	var v uint64
	v |= uint64(d[0])
	v |= uint64(d[1]) <<8
	v |= uint64(d[2]) <<16
	v |= uint64(d[3]) <<24
	*f=*(*float32)(unsafe.Pointer(&v))
	return 4
}

func SizeofFloat32() uint64 {
	return 4
}

func EncodeFloat64(buf []byte,f float64) uint64 {
	v:=*(*uint64)(unsafe.Pointer(&f))
	if cap(buf) >= 8 {
		buf = buf[:8]
	} else {
		buf = make([]byte, 8)
	}
	buf[0]=uint8(v)
	buf[1]=uint8(v>>8)
	buf[2]=uint8(v>>16)
	buf[3]=uint8(v>>24)
	buf[4]=uint8(v>>32)
	buf[5]=uint8(v>>40)
	buf[6]=uint8(v>>48)
	buf[7]=uint8(v>>56)
	return 8
}

func DecodeFloat64(d []byte,f *float64)uint64{
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

func SizeofFloat64() uint64 {
	return 8
}

func EncodeBool(buf []byte,v bool) uint64 {
	size:=1
	if cap(buf) >= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	if !v{
		buf[0]=0
	}else{
		buf[0]=1
	}
	return 1
}

func DecodeBool(d []byte,s *bool)uint64 {
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

func SizeofBool() uint64 {
	return 1
}


func EncodeString(buf []byte,v string) uint64{
	length:=uint64(len(v))
	var length_size uint64
	var size uint64
	length_size=SizeofVarint(length)
	size =length_size+length
	if uint64(cap(buf) )>= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	t:=length
	for i := uint64(0);i<length_size-1;i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[length_size-1] = byte(t)
	copy(buf[length_size:],v)
	return size
}

func DecodeString(d []byte,s *string)uint64{
	var t uint64
	var n uint64
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
	b:=d[n:n+t]
	*s=*(*string)(unsafe.Pointer(&b))
	return n+t
}

func SizeofString(v string) uint64 {
	length:=uint64(len(v))
	return SizeofVarint(length)+length
}

func EncodeBytes(buf []byte,v []byte) uint64 {
	length:=uint64(len(v))
	length_size:=SizeofVarint(length)
	var size uint64 =length_size+length
	if uint64(cap(buf) )>= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	t:=length
	for i := uint64(0);i<length_size-1;i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[length_size-1] = byte(t)
	copy(buf[length_size:],v)
	return size
}

func DecodeBytes(d []byte,s *[]byte) uint64 {
	var t uint64
	var n uint64
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
	*s=d[n:n+t]
	return n+t
}

func SizeofBytes(v []byte) uint64 {
	length:=uint64(len(v))
	return SizeofVarint(length)+length
}

func EncodeSliceUint8(buf []byte,d []uint8)uint64 {
	var offset uint64
	var size uint64
	length:=uint64(len(d))
	length_size:=SizeofVarint(length)
	size =length_size+length
	if uint64(cap(buf) )>= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	t:=length
	for i := uint64(0);i<length_size-1;i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[length_size-1] = byte(t)
	offset=length_size
	for _,v:=range d{
		t:=v
		buf[offset+0]=uint8(t)
		offset+=1
	}
	return size
}

func DecodeSliceUint8(d []byte,s *[]uint8) uint64{
	var length uint64
	var offset uint64
	var n uint64
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
	length=t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s=make([]uint8,length)
	}
	offset=n
	for i:=uint64(0);i<length;i++ {
		var v uint8
		v |= uint8(d[offset+0])
		(*s)[i]=v
		offset+=1
	}
	return offset
}

func SizeofSliceUint8(d []uint8) uint64 {
	var length uint64 =uint64(len(d))
	return SizeofVarint(length)+length
}

func EncodeSliceUint16(buf []byte,d []uint16)uint64 {
	var offset uint64
	var size uint64
	length:=uint64(len(d))
	length_size:=SizeofVarint(length)
	size =length_size+length*2
	if uint64(cap(buf) )>= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	t:=length
	for i := uint64(0);i<length_size-1;i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[length_size-1] = byte(t)
	offset=length_size
	for _,v:=range d{
		t:=v
		buf[offset+0]=uint8(t)
		buf[offset+1]=uint8(t>>8)
		offset+=2
	}
	return size
}

func DecodeSliceUint16(d []byte,s *[]uint16) uint64{
	var length uint64
	var offset uint64
	var n uint64
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
	length=t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s=make([]uint16,length)
	}
	offset=n
	for i:=uint64(0);i<length;i++ {
		var v uint16
		v |= uint16(d[offset+0])
		v |= uint16(d[offset+1]) <<8
		(*s)[i]=v
		offset+=2
	}
	return offset
}

func SizeofSliceUint16(d []uint16) uint64 {
	var length uint64 =uint64(len(d))
	return SizeofVarint(length)+length*2
}

func EncodeSliceUint32(buf []byte,d []uint32)uint64 {
	var offset uint64
	var size uint64
	length:=uint64(len(d))
	length_size:=SizeofVarint(length)
	size =length_size+length*4
	if uint64(cap(buf) )>= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	t:=length
	for i := uint64(0);i<length_size-1;i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[length_size-1] = byte(t)
	offset=length_size
	for _,v:=range d{
		t:=v
		buf[offset+0]=uint8(t)
		buf[offset+1]=uint8(t>>8)
		buf[offset+2]=uint8(t>>16)
		buf[offset+3]=uint8(t>>24)
		offset+=4
	}
	return size
}

func DecodeSliceUint32(d []byte,s *[]uint32) uint64{
	var length uint64
	var offset uint64
	var n uint64
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
	length=t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s=make([]uint32,length)
	}
	offset=n
	for i:=uint64(0);i<length;i++ {
		var v uint32
		v |= uint32(d[offset+0])
		v |= uint32(d[offset+1]) <<8
		v |= uint32(d[offset+2]) <<16
		v |= uint32(d[offset+3]) <<24
		(*s)[i]=v
		offset+=4
	}
	return offset
}

func SizeofSliceUint32(d []uint32) uint64 {
	var length uint64 =uint64(len(d))
	return SizeofVarint(length)+length*4
}

func EncodeSliceUint64(buf []byte,d []uint64)uint64 {
	var offset uint64
	var size uint64
	length:=uint64(len(d))
	length_size:=SizeofVarint(length)
	size =length_size+length*8
	if uint64(cap(buf) )>= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	t:=length
	for i := uint64(0);i<length_size-1;i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[length_size-1] = byte(t)
	offset=length_size
	for _,v:=range d{
		t:=v
		buf[offset+0]=uint8(t)
		buf[offset+1]=uint8(t>>8)
		buf[offset+2]=uint8(t>>16)
		buf[offset+3]=uint8(t>>24)
		buf[offset+4]=uint8(t>>32)
		buf[offset+5]=uint8(t>>40)
		buf[offset+6]=uint8(t>>48)
		buf[offset+7]=uint8(t>>56)
		offset+=8
	}
	return size
}

func DecodeSliceUint64(d []byte,s *[]uint64) uint64{
	var length uint64
	var offset uint64
	var n uint64
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
	length=t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s=make([]uint64,length)
	}
	offset=n
	for i:=uint64(0);i<length;i++ {
		var v uint64
		v |= uint64(d[offset+0])
		v |= uint64(d[offset+1]) <<8
		v |= uint64(d[offset+2]) <<16
		v |= uint64(d[offset+3]) <<24
		v |= uint64(d[offset+4]) <<32
		v |= uint64(d[offset+5]) <<40
		v |= uint64(d[offset+6]) <<48
		v |= uint64(d[offset+7]) <<56
		(*s)[i]=v
		offset+=8
	}
	return offset
}

func SizeofSliceUint64(d []uint64) uint64 {
	var length uint64 =uint64(len(d))
	return SizeofVarint(length)+length*8
}

func EncodeSliceVarint(buf []byte,d []uint64)uint64 {
	var offset uint64
	var size uint64
	length:=uint64(len(d))
	length_size:=SizeofVarint(length)
	size =length_size
	if uint64(cap(buf) )>= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	t:=length
	for i := uint64(0);i<length_size-1;i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[length_size-1] = byte(t)
	offset=length_size
	for _,v:=range d{
		s:=SizeofVarint(v)
		size+=s
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			tmp:=make([]byte,size)
			copy(tmp,buf)
			buf=tmp
		}
		t:=v
		for i := uint64(0);i<s-1;i++ {
			buf[offset+i] = byte(t) | msb7
			t >>= 7
		}
		buf[offset+s-1] = byte(t)
		offset+=s
	}
	return size
}

func DecodeSliceVarint(d []byte,s *[]uint64) uint64{
	var length uint64
	var offset uint64
	var n uint64
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
	length=t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s=make([]uint64,length)
	}
	offset=n
	for i:=uint64(0);i<length;i++ {
		buf:=d[offset:]
		j:=uint64(0)
		var t uint64
		t = uint64(buf[j]&mask7)
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 7
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 14
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 21
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 28
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 35
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 42
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 49
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 56
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 63
		goto fordone
	fordone:
		j++
		(*s)[i]=t
		offset+=j
	}
	return offset
}

func SizeofSliceVarint(d []uint64) uint64 {
	var size uint64
	size=SizeofVarint(uint64(len(d)))
	for _,v:=range d{
		size+=SizeofVarint(v)
	}
	return size
}

func EncodeSliceFloat32 (buf []byte,d []float32)uint64 {
	var offset uint64
	var size uint64
	length:=uint64(len(d))
	length_size:=SizeofVarint(length)
	size =length_size+length*4
	if uint64(cap(buf) )>= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	t:=length
	for i := uint64(0);i<length_size-1;i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[length_size-1] = byte(t)
	offset=length_size
	for _,f:=range d{
		v:=*(*uint32)(unsafe.Pointer(&f))
		buf[offset+0]=uint8(v)
		buf[offset+1]=uint8(v>>8)
		buf[offset+2]=uint8(v>>16)
		buf[offset+3]=uint8(v>>24)
		offset+=4
	}
	return size
}

func DecodeSliceFloat32(d []byte,s *[]float32) uint64{
	var length uint64
	var offset uint64
	var n uint64
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
	length=t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s=make([]float32,length)
	}
	offset=n
	for i:=uint64(0);i<length;i++ {
		var v uint32
		v |= uint32(d[offset+0])
		v |= uint32(d[offset+1]) <<8
		v |= uint32(d[offset+2]) <<16
		v |= uint32(d[offset+3]) <<24
		(*s)[i]=*(*float32)(unsafe.Pointer(&v))
		offset+=4
	}
	return offset
}

func SizeofSliceFloat32(d []float32) uint64 {
	var length uint64 =uint64(len(d))
	return SizeofVarint(length)+length*4
}

func EncodeSliceFloat64 (buf []byte,d []float64)uint64 {
	var offset uint64
	var size uint64
	length:=uint64(len(d))
	length_size:=SizeofVarint(length)
	size =length_size+length*8
	if uint64(cap(buf) )>= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	t:=length
	for i := uint64(0);i<length_size-1;i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[length_size-1] = byte(t)
	offset=length_size
	for _,f:=range d{
		v:=*(*uint64)(unsafe.Pointer(&f))
		buf[offset+0]=uint8(v)
		buf[offset+1]=uint8(v>>8)
		buf[offset+2]=uint8(v>>16)
		buf[offset+3]=uint8(v>>24)
		buf[offset+4]=uint8(v>>32)
		buf[offset+5]=uint8(v>>40)
		buf[offset+6]=uint8(v>>48)
		buf[offset+7]=uint8(v>>56)
		offset+=8
	}
	return size
}

func DecodeSliceFloat64(d []byte,s *[]float64) uint64{
	var length uint64
	var offset uint64
	var n uint64
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
	length=t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s=make([]float64,length)
	}
	offset=n
	for i:=uint64(0);i<length;i++ {
		var v uint64
		v |= uint64(d[offset+0])
		v |= uint64(d[offset+1]) <<8
		v |= uint64(d[offset+2]) <<16
		v |= uint64(d[offset+3]) <<24
		v |= uint64(d[offset+4]) <<32
		v |= uint64(d[offset+5]) <<40
		v |= uint64(d[offset+6]) <<48
		v |= uint64(d[offset+7]) <<56
		(*s)[i]=*(*float64)(unsafe.Pointer(&v))
		offset+=8
	}
	return offset
}

func SizeofSliceFloat64(d []float64) uint64 {
	var length uint64 =uint64(len(d))
	return SizeofVarint(length)+length*8
}

func EncodeSliceBool(buf []byte,d []bool)uint64 {
	var offset uint64
	var size uint64
	length:=uint64(len(d))
	length_size:=SizeofVarint(length)
	size =length_size+length
	if uint64(cap(buf) )>= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	t:=length
	for i := uint64(0);i<length_size-1;i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[length_size-1] = byte(t)
	offset=length_size
	for _,v:=range d{
		if !v{
			buf[offset]=0
		}else{
			buf[offset]=1
		}
		offset+=1
	}
	return size
}

func DecodeSliceBool(d []byte,s *[]bool) uint64{
	var length uint64
	var offset uint64
	var n uint64
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
	length=t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s=make([]bool,length)
	}
	offset=n
	for i:=uint64(0);i<length;i++ {
		if d[offset]==0{
			(*s)[i]=false
		}else {
			(*s)[i]=true
		}
		offset+=1
	}
	return offset
}

func SizeofSliceBool(d []bool) uint64 {
	length:=uint64(len(d))
	return SizeofVarint(length)+length
}

func EncodeSliceString(buf []byte,d []string)uint64 {
	var offset uint64
	var size uint64
	length:=uint64(len(d))
	length_size:=SizeofVarint(length)
	size =length_size
	if uint64(cap(buf) )>= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	t:=length
	for i := uint64(0);i<length_size-1;i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[length_size-1] = byte(t)
	offset=length_size
	for _,v:=range d{
		length:=uint64(len(v))
		length_size:=SizeofVarint(length)
		size+=length_size+length
		t:=length
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			tmp:=make([]byte,size)
			copy(tmp,buf)
			buf=tmp
		}
		for i := uint64(0);i<length_size-1;i++ {
			buf[offset+i] = byte(t) | msb7
			t >>= 7
		}
		buf[offset+length_size-1] = byte(t)
		copy(buf[offset+length_size:],v)
		offset+=length+length_size
	}
	return size
}

func DecodeSliceString(d []byte,s *[]string) uint64{
	var length uint64
	var l uint64
	var offset uint64
	var n uint64
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
	length=t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s=make([]string,length)
	}
	offset=n
	for i:=uint64(0);i<length;i++ {
		buf:=d[offset:]
		j:=uint64(0)
		var t uint64
		t = uint64(buf[j]&mask7)
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 7
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 14
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 21
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 28
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 35
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 42
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 49
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 56
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 63
		goto fordone
	fordone:
		j++
		l++
		b:=d[offset+j:offset+j+t]
		(*s)[i]=*(*string)(unsafe.Pointer(&b))
		offset+=j+t
	}
	return offset
}

func SizeofSliceString(d []string) uint64 {
	var size uint64
	length:=uint64(len(d))
	s:=SizeofVarint(length)
	size=s
	for _,v:=range d{
		length:=uint64(len(v))
		s:=SizeofVarint(length)
		size+=s+length
	}
	return size
}

func EncodeSliceBytes(buf []byte,d [][]byte)uint64 {
	var offset uint64
	var size uint64
	length:=uint64(len(d))
	length_size:=SizeofVarint(length)
	size =length_size
	if uint64(cap(buf) )>= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	t:=length
	for i := uint64(0);i<length_size-1;i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[length_size-1] = byte(t)
	offset=length_size
	for _,v:=range d{
		length:=uint64(len(v))
		length_size:=SizeofVarint(length)
		size+=length_size+length
		t:=length
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			tmp:=make([]byte,size)
			copy(tmp,buf)
			buf=tmp
		}
		for i := uint64(0);i<length_size-1;i++ {
			buf[offset+i] = byte(t) | msb7
			t >>= 7
		}
		buf[offset+length_size-1] = byte(t)
		copy(buf[offset+length_size:],v)
		offset+=length+length_size
	}
	return size
}

func DecodeSliceBytes(d []byte,s *[][]byte) uint64{
	var length uint64
	var l uint64
	var offset uint64
	var n uint64
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
	length=t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s=make([][]byte,length)
	}
	offset=n
	for i:=uint64(0);i<length;i++ {
		buf:=d[offset:]
		j:=uint64(0)
		var t uint64
		t = uint64(buf[j]&mask7)
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 7
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 14
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 21
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 28
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 35
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 42
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 49
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 56
		if buf[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(buf[j]&mask7) << 63
		goto fordone
	fordone:
		j++
		l++
		b:=d[offset+j:offset+j+t]
		(*s)[i]=b
		offset+=j+t
	}
	return offset
}

func SizeofSliceBytes(d [][]byte) uint64 {
	var size uint64
	length:=uint64(len(d))
	s:=SizeofVarint(length)
	size=s
	for _,v:=range d{
		length:=uint64(len(v))
		s:=SizeofVarint(length)
		size+=s+length
	}
	return size
}

