package code

import (
	"unsafe"
)

const (
	v7    = 1 << 7
	v14   = 1 << 14
	v21   = 1 << 21
	v28   = 1 << 28
	v35   = 1 << 35
	v42   = 1 << 42
	v49   = 1 << 49
	v56   = 1 << 56
	v63   = 1 << 63
	mask7 = -1 ^ (-1 << 7)
	msb7  = 1 << 7
)

// CheckBuffer checks the capacity of buf and returns the buffer of sufficient size.
// If the buf is too small, CheckBuffer will make a new buffer.
func CheckBuffer(buf []byte, n uint64) []byte {
	if cap(buf) >= int(n) {
		buf = buf[:n]
	} else {
		buf = make([]byte, n)
	}
	return buf
}

// EncodeUint8 encodes a uint8 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeUint8 will panic.
func EncodeUint8(buf []byte, v uint8) uint64 {
	buf[0] = uint8(v)
	return 1
}

// DecodeUint8 decodes a uint8 from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeUint8, DecodeUint8 will panic.
func DecodeUint8(buf []byte, v *uint8) uint64 {
	*v = uint8(buf[0])
	return 1
}

// SizeofUint8 takes a uint8 and returns the number of bytes.
func SizeofUint8(v uint8) uint64 {
	return 1
}

// EncodeUint16 encodes a uint16 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeUint16 will panic.
func EncodeUint16(buf []byte, v uint16) uint64 {
	var t = v
	buf[0] = uint8(t)
	buf[1] = uint8(t >> 8)
	return 2
}

// DecodeUint16 decodes a uint16 from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeUint16, DecodeUint16 will panic.
func DecodeUint16(buf []byte, v *uint16) uint64 {
	var t uint16
	t = uint16(buf[0])
	t |= uint16(buf[1]) << 8
	*v = t
	return 2
}

// SizeofUint16 takes a uint16 and returns the number of bytes.
func SizeofUint16(v uint16) uint64 {
	return 2
}

// EncodeUint32 encodes a uint32 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeUint32 will panic.
func EncodeUint32(buf []byte, v uint32) uint64 {
	var t = v
	buf[0] = uint8(t)
	buf[1] = uint8(t >> 8)
	buf[2] = uint8(t >> 16)
	buf[3] = uint8(t >> 24)
	return 4
}

// DecodeUint32 decodes a uint32 from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeUint32, DecodeUint32 will panic.
func DecodeUint32(buf []byte, v *uint32) uint64 {
	var t uint32
	t = uint32(buf[0])
	t |= uint32(buf[1]) << 8
	t |= uint32(buf[2]) << 16
	t |= uint32(buf[3]) << 24
	*v = t
	return 4
}

// SizeofUint32 takes a uint32 and returns the number of bytes.
func SizeofUint32(v uint32) uint64 {
	return 4
}

// EncodeUint64 encodes a uint64 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeUint64 will panic.
func EncodeUint64(buf []byte, v uint64) uint64 {
	var t = v
	buf[0] = uint8(t)
	buf[1] = uint8(t >> 8)
	buf[2] = uint8(t >> 16)
	buf[3] = uint8(t >> 24)
	buf[4] = uint8(t >> 32)
	buf[5] = uint8(t >> 40)
	buf[6] = uint8(t >> 48)
	buf[7] = uint8(t >> 56)
	return 8
}

// DecodeUint64 decodes a uint64 from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeUint64, DecodeUint64 will panic.
func DecodeUint64(buf []byte, v *uint64) uint64 {
	var t uint64
	t = uint64(buf[0])
	t |= uint64(buf[1]) << 8
	t |= uint64(buf[2]) << 16
	t |= uint64(buf[3]) << 24
	t |= uint64(buf[4]) << 32
	t |= uint64(buf[5]) << 40
	t |= uint64(buf[6]) << 48
	t |= uint64(buf[7]) << 56
	*v = t
	return 8
}

// SizeofUint64 takes a uint64 and returns the number of bytes.
func SizeofUint64(v uint64) uint64 {
	return 8
}

// EncodeVarint encodes a uint64 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeVarint will panic.
func EncodeVarint(buf []byte, v uint64) uint64 {
	var t = v
	var size = SizeofVarint(v)
	for i := uint64(0); i < size-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[size-1] = byte(t)
	return size
}

// DecodeVarint decodes a uint64 from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeVarint, DecodeVarint will panic.
func DecodeVarint(d []byte, v *uint64) uint64 {
	var t uint64
	var n uint64
	t = uint64(d[n] & mask7)
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
	*v = t
	return n
}

// SizeofVarint takes a uint64 and returns the number of bytes.
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

// EncodeFloat32 encodes a float32 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeFloat32 will panic.
func EncodeFloat32(buf []byte, f float32) uint64 {
	v := *(*uint32)(unsafe.Pointer(&f))
	buf[0] = uint8(v)
	buf[1] = uint8(v >> 8)
	buf[2] = uint8(v >> 16)
	buf[3] = uint8(v >> 24)
	return 4
}

// DecodeFloat32 decodes a float32 from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeFloat32, DecodeFloat32 will panic.
func DecodeFloat32(d []byte, f *float32) uint64 {
	var v uint64
	v |= uint64(d[0])
	v |= uint64(d[1]) << 8
	v |= uint64(d[2]) << 16
	v |= uint64(d[3]) << 24
	*f = *(*float32)(unsafe.Pointer(&v))
	return 4
}

// SizeofFloat32 takes a float32 and returns the number of bytes.
func SizeofFloat32(f float32) uint64 {
	return 4
}

// EncodeFloat64 encodes a float64 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeFloat64 will panic.
func EncodeFloat64(buf []byte, f float64) uint64 {
	v := *(*uint64)(unsafe.Pointer(&f))
	buf[0] = uint8(v)
	buf[1] = uint8(v >> 8)
	buf[2] = uint8(v >> 16)
	buf[3] = uint8(v >> 24)
	buf[4] = uint8(v >> 32)
	buf[5] = uint8(v >> 40)
	buf[6] = uint8(v >> 48)
	buf[7] = uint8(v >> 56)
	return 8
}

// DecodeFloat64 decodes a float64 from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeFloat64, DecodeFloat64 will panic.
func DecodeFloat64(d []byte, f *float64) uint64 {
	var v uint64
	v |= uint64(d[0])
	v |= uint64(d[1]) << 8
	v |= uint64(d[2]) << 16
	v |= uint64(d[3]) << 24
	v |= uint64(d[4]) << 32
	v |= uint64(d[5]) << 40
	v |= uint64(d[6]) << 48
	v |= uint64(d[7]) << 56
	*f = *(*float64)(unsafe.Pointer(&v))
	return 8
}

// SizeofFloat64 takes a float64 and returns the number of bytes.
func SizeofFloat64(f float64) uint64 {
	return 8
}

// EncodeBool encodes a bool into buf and returns the number of bytes written.
// If the buffer is too small, EncodeBool will panic.
func EncodeBool(buf []byte, v bool) uint64 {
	if !v {
		buf[0] = 0
	} else {
		buf[0] = 1
	}
	return 1
}

// DecodeBool decodes a bool from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeBool, DecodeBool will panic.
func DecodeBool(d []byte, s *bool) uint64 {
	if len(d) == 0 {
		*s = false
		return 0
	}
	if d[0] == 0 {
		*s = false
		return 1
	}
	*s = true
	return 1
}

// SizeofBool takes a bool and returns the number of bytes.
func SizeofBool(v bool) uint64 {
	return 1
}

// EncodeString encodes a string into buf and returns the number of bytes written.
// If the buffer is too small, EncodeString will panic.
func EncodeString(buf []byte, v string) uint64 {
	length := uint64(len(v))
	var lengthSize uint64
	var size uint64
	lengthSize = SizeofVarint(length)
	size = lengthSize + length
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	copy(buf[lengthSize:], v)
	return size
}

// DecodeString decodes a string from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeString, DecodeString will panic.
func DecodeString(d []byte, s *string) uint64 {
	var t uint64
	var n uint64
	t = uint64(d[n] & mask7)
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
	b := d[n : n+t]
	*s = *(*string)(unsafe.Pointer(&b))
	return n + t
}

// SizeofString takes a string and returns the number of bytes.
func SizeofString(v string) uint64 {
	length := uint64(len(v))
	return SizeofVarint(length) + length
}

// EncodeBytes encodes a []byte into buf and returns the number of bytes written.
// If the buffer is too small, EncodeBytes will panic.
func EncodeBytes(buf []byte, v []byte) uint64 {
	length := uint64(len(v))
	lengthSize := SizeofVarint(length)
	var size uint64 = lengthSize + length
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	copy(buf[lengthSize:], v)
	return size
}

// DecodeBytes decodes a []byte from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeBytes, DecodeBytes will panic.
func DecodeBytes(d []byte, s *[]byte) uint64 {
	var t uint64
	var n uint64
	t = uint64(d[n] & mask7)
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
	*s = d[n : n+t]
	return n + t
}

// SizeofBytes takes a []byte and returns the number of bytes.
func SizeofBytes(v []byte) uint64 {
	length := uint64(len(v))
	return SizeofVarint(length) + length
}

// EncodeSliceUint8 encodes a []uint8 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeSliceUint8 will panic.
func EncodeSliceUint8(buf []byte, d []uint8) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(d))
	lengthSize := SizeofVarint(length)
	size = lengthSize + length
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, v := range d {
		t := v
		buf[offset+0] = uint8(t)
		offset++
	}
	return size
}

// DecodeSliceUint8 decodes a []uint8 from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeSliceUint8, DecodeSliceUint8 will panic.
func DecodeSliceUint8(d []byte, s *[]uint8) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(d[n] & mask7)
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
	length = t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s = make([]uint8, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		var v uint8
		v |= uint8(d[offset+0])
		(*s)[i] = v
		offset++
	}
	return offset
}

// SizeofSliceUint8 takes a []uint8 and returns the number of bytes.
func SizeofSliceUint8(d []uint8) uint64 {
	var length uint64 = uint64(len(d))
	return SizeofVarint(length) + length
}

// EncodeSliceUint16 encodes a []uint16 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeSliceUint16 will panic.
func EncodeSliceUint16(buf []byte, d []uint16) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(d))
	lengthSize := SizeofVarint(length)
	size = lengthSize + length*2
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, v := range d {
		t := v
		buf[offset+0] = uint8(t)
		buf[offset+1] = uint8(t >> 8)
		offset += 2
	}
	return size
}

// DecodeSliceUint16 decodes a uint64 from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeSliceUint16, DecodeSliceUint16 will panic.
func DecodeSliceUint16(d []byte, s *[]uint16) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(d[n] & mask7)
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
	length = t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s = make([]uint16, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		var v uint16
		v |= uint16(d[offset+0])
		v |= uint16(d[offset+1]) << 8
		(*s)[i] = v
		offset += 2
	}
	return offset
}

// SizeofSliceUint16 takes a []uint16 and returns the number of bytes.
func SizeofSliceUint16(d []uint16) uint64 {
	var length uint64 = uint64(len(d))
	return SizeofVarint(length) + length*2
}

// EncodeSliceUint32 encodes a []uint32 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeSliceUint32 will panic.
func EncodeSliceUint32(buf []byte, d []uint32) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(d))
	lengthSize := SizeofVarint(length)
	size = lengthSize + length*4
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, v := range d {
		t := v
		buf[offset+0] = uint8(t)
		buf[offset+1] = uint8(t >> 8)
		buf[offset+2] = uint8(t >> 16)
		buf[offset+3] = uint8(t >> 24)
		offset += 4
	}
	return size
}

// DecodeSliceUint32 decodes a []uint32 from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeSliceUint32, DecodeSliceUint32 will panic.
func DecodeSliceUint32(d []byte, s *[]uint32) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(d[n] & mask7)
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
	length = t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s = make([]uint32, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		var v uint32
		v |= uint32(d[offset+0])
		v |= uint32(d[offset+1]) << 8
		v |= uint32(d[offset+2]) << 16
		v |= uint32(d[offset+3]) << 24
		(*s)[i] = v
		offset += 4
	}
	return offset
}

// SizeofSliceUint32 takes a []uint32 and returns the number of bytes.
func SizeofSliceUint32(d []uint32) uint64 {
	var length uint64 = uint64(len(d))
	return SizeofVarint(length) + length*4
}

// EncodeSliceUint64 encodes a []uint64 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeSliceUint64 will panic.
func EncodeSliceUint64(buf []byte, d []uint64) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(d))
	lengthSize := SizeofVarint(length)
	size = lengthSize + length*8
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, v := range d {
		t := v
		buf[offset+0] = uint8(t)
		buf[offset+1] = uint8(t >> 8)
		buf[offset+2] = uint8(t >> 16)
		buf[offset+3] = uint8(t >> 24)
		buf[offset+4] = uint8(t >> 32)
		buf[offset+5] = uint8(t >> 40)
		buf[offset+6] = uint8(t >> 48)
		buf[offset+7] = uint8(t >> 56)
		offset += 8
	}
	return size
}

// DecodeSliceUint64 decodes a []uint64 from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeSliceUint64, DecodeSliceUint64 will panic.
func DecodeSliceUint64(d []byte, s *[]uint64) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(d[n] & mask7)
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
	length = t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s = make([]uint64, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		var v uint64
		v |= uint64(d[offset+0])
		v |= uint64(d[offset+1]) << 8
		v |= uint64(d[offset+2]) << 16
		v |= uint64(d[offset+3]) << 24
		v |= uint64(d[offset+4]) << 32
		v |= uint64(d[offset+5]) << 40
		v |= uint64(d[offset+6]) << 48
		v |= uint64(d[offset+7]) << 56
		(*s)[i] = v
		offset += 8
	}
	return offset
}

// SizeofSliceUint64 takes a []uint64 and returns the number of bytes.
func SizeofSliceUint64(d []uint64) uint64 {
	var length uint64 = uint64(len(d))
	return SizeofVarint(length) + length*8
}

// EncodeSliceVarint encodes a d []uint64 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeSliceVarint will panic.
func EncodeSliceVarint(buf []byte, d []uint64) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(d))
	lengthSize := SizeofVarint(length)
	size = lengthSize
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, v := range d {
		s := SizeofVarint(v)
		size += s
		t := v
		for i := uint64(0); i < s-1; i++ {
			buf[offset+i] = byte(t) | msb7
			t >>= 7
		}
		buf[offset+s-1] = byte(t)
		offset += s
	}
	return size
}

// DecodeSliceVarint decodes a []uint64 from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeSliceVarint, DecodeSliceVarint will panic.
func DecodeSliceVarint(d []byte, s *[]uint64) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(d[n] & mask7)
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
	length = t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s = make([]uint64, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		buf := d[offset:]
		j := uint64(0)
		var t uint64
		t = uint64(buf[j] & mask7)
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
		(*s)[i] = t
		offset += j
	}
	return offset
}

// SizeofSliceVarint takes a []uint64 and returns the number of bytes.
func SizeofSliceVarint(d []uint64) uint64 {
	var size uint64
	size = SizeofVarint(uint64(len(d)))
	for _, v := range d {
		size += SizeofVarint(v)
	}
	return size
}

// EncodeSliceFloat32 encodes a []float32 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeSliceFloat32 will panic.
func EncodeSliceFloat32(buf []byte, d []float32) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(d))
	lengthSize := SizeofVarint(length)
	size = lengthSize + length*4
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, f := range d {
		v := *(*uint32)(unsafe.Pointer(&f))
		buf[offset+0] = uint8(v)
		buf[offset+1] = uint8(v >> 8)
		buf[offset+2] = uint8(v >> 16)
		buf[offset+3] = uint8(v >> 24)
		offset += 4
	}
	return size
}

// DecodeSliceFloat32 decodes a []float32 from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeSliceFloat32, DecodeSliceFloat32 will panic.
func DecodeSliceFloat32(d []byte, s *[]float32) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(d[n] & mask7)
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
	length = t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s = make([]float32, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		var v uint32
		v |= uint32(d[offset+0])
		v |= uint32(d[offset+1]) << 8
		v |= uint32(d[offset+2]) << 16
		v |= uint32(d[offset+3]) << 24
		(*s)[i] = *(*float32)(unsafe.Pointer(&v))
		offset += 4
	}
	return offset
}

// SizeofSliceFloat32 takes a []float32 and returns the number of bytes.
func SizeofSliceFloat32(d []float32) uint64 {
	var length uint64 = uint64(len(d))
	return SizeofVarint(length) + length*4
}

// EncodeSliceFloat64 encodes a []float64 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeSliceFloat64 will panic.
func EncodeSliceFloat64(buf []byte, d []float64) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(d))
	lengthSize := SizeofVarint(length)
	size = lengthSize + length*8
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, f := range d {
		v := *(*uint64)(unsafe.Pointer(&f))
		buf[offset+0] = uint8(v)
		buf[offset+1] = uint8(v >> 8)
		buf[offset+2] = uint8(v >> 16)
		buf[offset+3] = uint8(v >> 24)
		buf[offset+4] = uint8(v >> 32)
		buf[offset+5] = uint8(v >> 40)
		buf[offset+6] = uint8(v >> 48)
		buf[offset+7] = uint8(v >> 56)
		offset += 8
	}
	return size
}

// DecodeSliceFloat64 decodes a uint64 from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeSliceFloat64, DecodeSliceFloat64 will panic.
func DecodeSliceFloat64(d []byte, s *[]float64) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(d[n] & mask7)
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
	length = t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s = make([]float64, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		var v uint64
		v |= uint64(d[offset+0])
		v |= uint64(d[offset+1]) << 8
		v |= uint64(d[offset+2]) << 16
		v |= uint64(d[offset+3]) << 24
		v |= uint64(d[offset+4]) << 32
		v |= uint64(d[offset+5]) << 40
		v |= uint64(d[offset+6]) << 48
		v |= uint64(d[offset+7]) << 56
		(*s)[i] = *(*float64)(unsafe.Pointer(&v))
		offset += 8
	}
	return offset
}

// SizeofSliceFloat64 takes a []float64 and returns the number of bytes.
func SizeofSliceFloat64(d []float64) uint64 {
	var length uint64 = uint64(len(d))
	return SizeofVarint(length) + length*8
}

// EncodeSliceBool encodes a []bool into buf and returns the number of bytes written.
// If the buffer is too small, EncodeSliceBool will panic.
func EncodeSliceBool(buf []byte, d []bool) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(d))
	lengthSize := SizeofVarint(length)
	size = lengthSize + length
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, v := range d {
		if !v {
			buf[offset] = 0
		} else {
			buf[offset] = 1
		}
		offset++
	}
	return size
}

// DecodeSliceBool decodes a []bool from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeSliceBool, DecodeSliceBool will panic.
func DecodeSliceBool(d []byte, s *[]bool) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(d[n] & mask7)
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
	length = t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s = make([]bool, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		if d[offset] == 0 {
			(*s)[i] = false
		} else {
			(*s)[i] = true
		}
		offset++
	}
	return offset
}

// SizeofSliceBool takes a []bool and returns the number of bytes.
func SizeofSliceBool(d []bool) uint64 {
	length := uint64(len(d))
	return SizeofVarint(length) + length
}

// EncodeSliceString encodes a []string into buf and returns the number of bytes written.
// If the buffer is too small, EncodeSliceString will panic.
func EncodeSliceString(buf []byte, d []string) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(d))
	lengthSize := SizeofVarint(length)
	size = lengthSize
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, v := range d {
		length := uint64(len(v))
		lengthSize := SizeofVarint(length)
		size += lengthSize + length
		t := length
		for i := uint64(0); i < lengthSize-1; i++ {
			buf[offset+i] = byte(t) | msb7
			t >>= 7
		}
		buf[offset+lengthSize-1] = byte(t)
		copy(buf[offset+lengthSize:], v)
		offset += length + lengthSize
	}
	return size
}

// DecodeSliceString decodes a []string from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeSliceString, DecodeSliceString will panic.
func DecodeSliceString(d []byte, s *[]string) uint64 {
	var length uint64
	var l uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(d[n] & mask7)
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
	length = t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s = make([]string, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		buf := d[offset:]
		j := uint64(0)
		var t uint64
		t = uint64(buf[j] & mask7)
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
		b := d[offset+j : offset+j+t]
		(*s)[i] = *(*string)(unsafe.Pointer(&b))
		offset += j + t
	}
	return offset
}

// SizeofSliceString takes a []string and returns the number of bytes.
func SizeofSliceString(d []string) uint64 {
	var size uint64
	length := uint64(len(d))
	s := SizeofVarint(length)
	size = s
	for _, v := range d {
		length := uint64(len(v))
		s := SizeofVarint(length)
		size += s + length
	}
	return size
}

// EncodeSliceBytes encodes a d [][]byte into buf and returns the number of bytes written.
// If the buffer is too small, EncodeSliceBytes will panic.
func EncodeSliceBytes(buf []byte, d [][]byte) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(d))
	lengthSize := SizeofVarint(length)
	size = lengthSize
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, v := range d {
		length := uint64(len(v))
		lengthSize := SizeofVarint(length)
		size += lengthSize + length
		t := length
		for i := uint64(0); i < lengthSize-1; i++ {
			buf[offset+i] = byte(t) | msb7
			t >>= 7
		}
		buf[offset+lengthSize-1] = byte(t)
		copy(buf[offset+lengthSize:], v)
		offset += length + lengthSize
	}
	return size
}

// DecodeSliceBytes decodes a [][]byte from buf and returns the number of bytes read (> 0).
// If the buffer is not from EncodeSliceBytes, DecodeSliceBytes will panic.
func DecodeSliceBytes(d []byte, s *[][]byte) uint64 {
	var length uint64
	var l uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(d[n] & mask7)
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
	length = t
	if uint64(cap(*s)) >= length {
		*s = (*s)[:length]
	} else {
		*s = make([][]byte, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		buf := d[offset:]
		j := uint64(0)
		var t uint64
		t = uint64(buf[j] & mask7)
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
		b := d[offset+j : offset+j+t]
		(*s)[i] = b
		offset += j + t
	}
	return offset
}

// SizeofSliceBytes takes a [][]byte and returns the number of bytes.
func SizeofSliceBytes(d [][]byte) uint64 {
	var size uint64
	length := uint64(len(d))
	s := SizeofVarint(length)
	size = s
	for _, v := range d {
		length := uint64(len(v))
		s := SizeofVarint(length)
		size += s + length
	}
	return size
}
