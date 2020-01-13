// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package code implements encoding and decoding.
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

// DecodeUint8 decodes a uint8 from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeUint8, DecodeUint8 will panic.
func DecodeUint8(buf []byte, v *uint8) uint64 {
	*v = uint8(buf[0])
	return 1
}

// SizeofUint8 takes a uint8 and returns the number of bytes.
func SizeofUint8(v uint8) uint64 {
	return 1
}

// MaxUint8Bytes returns maximum length of a uint8.
func MaxUint8Bytes(v uint8) uint64 {
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

// DecodeUint16 decodes a uint16 from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
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

// MaxUint16Bytes returns maximum length of a uint16.
func MaxUint16Bytes(v uint16) uint64 {
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

// DecodeUint32 decodes a uint32 from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
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

// MaxUint32Bytes returns maximum length of a uint32.
func MaxUint32Bytes(v uint32) uint64 {
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

// DecodeUint64 decodes a uint64 from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
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

// MaxUint64Bytes returns maximum length of a uint64.
func MaxUint64Bytes(v uint64) uint64 {
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

// DecodeVarint decodes a uint64 from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeVarint, DecodeVarint will panic.
func DecodeVarint(buf []byte, v *uint64) uint64 {
	var t uint64
	var n uint64
	t = uint64(buf[n] & mask7)
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 7
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 14
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 21
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 28
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 35
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 42
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 49
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 56
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 63
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

// MaxVarintBytes returns maximum length of a varint.
func MaxVarintBytes(v uint64) uint64 {
	return 10
}

// EncodeFloat32 encodes a float32 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeFloat32 will panic.
func EncodeFloat32(buf []byte, v float32) uint64 {
	t := *(*uint32)(unsafe.Pointer(&v))
	buf[0] = uint8(t)
	buf[1] = uint8(t >> 8)
	buf[2] = uint8(t >> 16)
	buf[3] = uint8(t >> 24)
	return 4
}

// DecodeFloat32 decodes a float32 from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeFloat32, DecodeFloat32 will panic.
func DecodeFloat32(buf []byte, v *float32) uint64 {
	var t uint64
	t |= uint64(buf[0])
	t |= uint64(buf[1]) << 8
	t |= uint64(buf[2]) << 16
	t |= uint64(buf[3]) << 24
	*v = *(*float32)(unsafe.Pointer(&t))
	return 4
}

// SizeofFloat32 takes a float32 and returns the number of bytes.
func SizeofFloat32(v float32) uint64 {
	return 4
}

// MaxFloat32Bytes returns maximum length of a float32.
func MaxFloat32Bytes(v float32) uint64 {
	return 4
}

// EncodeFloat64 encodes a float64 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeFloat64 will panic.
func EncodeFloat64(buf []byte, v float64) uint64 {
	t := *(*uint64)(unsafe.Pointer(&v))
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

// DecodeFloat64 decodes a float64 from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeFloat64, DecodeFloat64 will panic.
func DecodeFloat64(buf []byte, v *float64) uint64 {
	var t uint64
	t |= uint64(buf[0])
	t |= uint64(buf[1]) << 8
	t |= uint64(buf[2]) << 16
	t |= uint64(buf[3]) << 24
	t |= uint64(buf[4]) << 32
	t |= uint64(buf[5]) << 40
	t |= uint64(buf[6]) << 48
	t |= uint64(buf[7]) << 56
	*v = *(*float64)(unsafe.Pointer(&t))
	return 8
}

// SizeofFloat64 takes a float64 and returns the number of bytes.
func SizeofFloat64(v float64) uint64 {
	return 8
}

// MaxFloat64Bytes returns maximum length of a float64.
func MaxFloat64Bytes(v float64) uint64 {
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

// DecodeBool decodes a bool from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeBool, DecodeBool will panic.
func DecodeBool(buf []byte, v *bool) uint64 {
	if len(buf) == 0 {
		*v = false
		return 0
	}
	if buf[0] == 0 {
		*v = false
		return 1
	}
	*v = true
	return 1
}

// SizeofBool takes a bool and returns the number of bytes.
func SizeofBool(v bool) uint64 {
	return 1
}

// MaxBoolBytes returns maximum length of a bool.
func MaxBoolBytes(v bool) uint64 {
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

// DecodeString decodes a string from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeString, DecodeString will panic.
func DecodeString(buf []byte, v *string) uint64 {
	var t uint64
	var n uint64
	t = uint64(buf[n] & mask7)
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 7
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 14
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 21
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 28
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 35
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 42
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 49
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 56
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 63
	goto done
done:
	n++
	b := buf[n : n+t]
	*v = *(*string)(unsafe.Pointer(&b))
	return n + t
}

// SizeofString takes a string and returns the number of bytes.
func SizeofString(v string) uint64 {
	length := uint64(len(v))
	return SizeofVarint(length) + length
}

// MaxStringBytes returns maximum length of a string.
func MaxStringBytes(v string) uint64 {
	return 10 + uint64(len(v))
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

// DecodeBytes decodes a []byte from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeBytes, DecodeBytes will panic.
func DecodeBytes(buf []byte, v *[]byte) uint64 {
	var t uint64
	var n uint64
	t = uint64(buf[n] & mask7)
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 7
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 14
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 21
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 28
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 35
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 42
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 49
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 56
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 63
	goto done
done:
	n++
	*v = buf[n : n+t]
	return n + t
}

// SizeofBytes takes a []byte and returns the number of bytes.
func SizeofBytes(v []byte) uint64 {
	length := uint64(len(v))
	return SizeofVarint(length) + length
}

// MaxBytesBytes returns maximum length of a []byte.
func MaxBytesBytes(v []byte) uint64 {
	return 10 + uint64(len(v))
}

// EncodeUint8Slice encodes a []uint8 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeUint8Slice will panic.
func EncodeUint8Slice(buf []byte, v []uint8) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(v))
	lengthSize := SizeofVarint(length)
	size = lengthSize + length
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, s := range v {
		t := s
		buf[offset+0] = uint8(t)
		offset++
	}
	return size
}

// DecodeUint8Slice decodes a []uint8 from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeUint8Slice, DecodeUint8Slice will panic.
func DecodeUint8Slice(buf []byte, v *[]uint8) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(buf[n] & mask7)
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 7
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 14
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 21
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 28
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 35
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 42
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 49
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 56
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 63
	goto done
done:
	n++
	length = t
	if uint64(cap(*v)) >= length {
		*v = (*v)[:length]
	} else {
		*v = make([]uint8, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		var s uint8
		s |= uint8(buf[offset+0])
		(*v)[i] = s
		offset++
	}
	return offset
}

// SizeofUint8Slice takes a []uint8 and returns the number of bytes.
func SizeofUint8Slice(v []uint8) uint64 {
	var length uint64 = uint64(len(v))
	return SizeofVarint(length) + length
}

// MaxUint8SliceBytes returns maximum length of a []uint8.
func MaxUint8SliceBytes(v []uint8) uint64 {
	var length uint64 = uint64(len(v))
	return 10 + length
}

// EncodeUint16Slice encodes a []uint16 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeUint16Slice will panic.
func EncodeUint16Slice(buf []byte, v []uint16) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(v))
	lengthSize := SizeofVarint(length)
	size = lengthSize + length*2
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, s := range v {
		t := s
		buf[offset+0] = uint8(t)
		buf[offset+1] = uint8(t >> 8)
		offset += 2
	}
	return size
}

// DecodeUint16Slice decodes a uint64 from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeUint16Slice, DecodeUint16Slice will panic.
func DecodeUint16Slice(buf []byte, v *[]uint16) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(buf[n] & mask7)
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 7
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 14
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 21
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 28
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 35
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 42
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 49
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 56
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 63
	goto done
done:
	n++
	length = t
	if uint64(cap(*v)) >= length {
		*v = (*v)[:length]
	} else {
		*v = make([]uint16, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		var s uint16
		s |= uint16(buf[offset+0])
		s |= uint16(buf[offset+1]) << 8
		(*v)[i] = s
		offset += 2
	}
	return offset
}

// SizeofUint16Slice takes a []uint16 and returns the number of bytes.
func SizeofUint16Slice(v []uint16) uint64 {
	var length uint64 = uint64(len(v))
	return SizeofVarint(length) + length*2
}

// MaxUint16SliceBytes returns maximum length of a []uint16.
func MaxUint16SliceBytes(v []uint16) uint64 {
	var length uint64 = uint64(len(v))
	return 10 + length*2
}

// EncodeUint32Slice encodes a []uint32 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeUint32Slice will panic.
func EncodeUint32Slice(buf []byte, v []uint32) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(v))
	lengthSize := SizeofVarint(length)
	size = lengthSize + length*4
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, v := range v {
		t := v
		buf[offset+0] = uint8(t)
		buf[offset+1] = uint8(t >> 8)
		buf[offset+2] = uint8(t >> 16)
		buf[offset+3] = uint8(t >> 24)
		offset += 4
	}
	return size
}

// DecodeUint32Slice decodes a []uint32 from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeUint32Slice, DecodeUint32Slice will panic.
func DecodeUint32Slice(buf []byte, v *[]uint32) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(buf[n] & mask7)
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 7
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 14
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 21
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 28
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 35
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 42
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 49
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 56
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 63
	goto done
done:
	n++
	length = t
	if uint64(cap(*v)) >= length {
		*v = (*v)[:length]
	} else {
		*v = make([]uint32, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		var s uint32
		s |= uint32(buf[offset+0])
		s |= uint32(buf[offset+1]) << 8
		s |= uint32(buf[offset+2]) << 16
		s |= uint32(buf[offset+3]) << 24
		(*v)[i] = s
		offset += 4
	}
	return offset
}

// SizeofUint32Slice takes a []uint32 and returns the number of bytes.
func SizeofUint32Slice(v []uint32) uint64 {
	var length uint64 = uint64(len(v))
	return SizeofVarint(length) + length*4
}

// MaxUint32SliceBytes returns maximum length of a []uint32.
func MaxUint32SliceBytes(v []uint32) uint64 {
	var length uint64 = uint64(len(v))
	return 10 + length*4
}

// EncodeUint64Slice encodes a []uint64 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeUint64Slice will panic.
func EncodeUint64Slice(buf []byte, v []uint64) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(v))
	lengthSize := SizeofVarint(length)
	size = lengthSize + length*8
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, v := range v {
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

// DecodeUint64Slice decodes a []uint64 from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeUint64Slice, DecodeUint64Slice will panic.
func DecodeUint64Slice(buf []byte, v *[]uint64) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(buf[n] & mask7)
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 7
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 14
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 21
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 28
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 35
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 42
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 49
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 56
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 63
	goto done
done:
	n++
	length = t
	if uint64(cap(*v)) >= length {
		*v = (*v)[:length]
	} else {
		*v = make([]uint64, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		var s uint64
		s |= uint64(buf[offset+0])
		s |= uint64(buf[offset+1]) << 8
		s |= uint64(buf[offset+2]) << 16
		s |= uint64(buf[offset+3]) << 24
		s |= uint64(buf[offset+4]) << 32
		s |= uint64(buf[offset+5]) << 40
		s |= uint64(buf[offset+6]) << 48
		s |= uint64(buf[offset+7]) << 56
		(*v)[i] = s
		offset += 8
	}
	return offset
}

// SizeofUint64Slice takes a []uint64 and returns the number of bytes.
func SizeofUint64Slice(v []uint64) uint64 {
	var length uint64 = uint64(len(v))
	return SizeofVarint(length) + length*8
}

// MaxUint64SliceBytes returns maximum length of a []uint64.
func MaxUint64SliceBytes(v []uint64) uint64 {
	var length uint64 = uint64(len(v))
	return 10 + length*8
}

// EncodeVarintSlice encodes a buf []uint64 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeVarintSlice will panic.
func EncodeVarintSlice(buf []byte, v []uint64) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(v))
	lengthSize := SizeofVarint(length)
	size = lengthSize
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, s := range v {
		sizeof := SizeofVarint(s)
		size += sizeof
		t := s
		for i := uint64(0); i < sizeof-1; i++ {
			buf[offset+i] = byte(t) | msb7
			t >>= 7
		}
		buf[offset+sizeof-1] = byte(t)
		offset += sizeof
	}
	return size
}

// DecodeVarintSlice decodes a []uint64 from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeVarintSlice, DecodeVarintSlice will panic.
func DecodeVarintSlice(buf []byte, v *[]uint64) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(buf[n] & mask7)
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 7
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 14
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 21
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 28
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 35
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 42
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 49
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 56
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 63
	goto done
done:
	n++
	length = t
	if uint64(cap(*v)) >= length {
		*v = (*v)[:length]
	} else {
		*v = make([]uint64, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		b := buf[offset:]
		j := uint64(0)
		var t uint64
		t = uint64(b[j] & mask7)
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 7
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 14
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 21
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 28
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 35
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 42
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 49
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 56
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 63
		goto fordone
	fordone:
		j++
		(*v)[i] = t
		offset += j
	}
	return offset
}

// SizeofVarintSlice takes a []uint64 and returns the number of bytes.
func SizeofVarintSlice(v []uint64) uint64 {
	var size uint64
	size = SizeofVarint(uint64(len(v)))
	for _, s := range v {
		size += SizeofVarint(s)
	}
	return size
}

// MaxVarintSliceBytes returns maximum length of a []varint.
func MaxVarintSliceBytes(v []uint64) uint64 {
	var length uint64 = uint64(len(v))
	return 10 + length*10
}

// EncodeFloat32Slice encodes a []float32 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeFloat32Slice will panic.
func EncodeFloat32Slice(buf []byte, v []float32) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(v))
	lengthSize := SizeofVarint(length)
	size = lengthSize + length*4
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, s := range v {
		f := *(*uint32)(unsafe.Pointer(&s))
		buf[offset+0] = uint8(f)
		buf[offset+1] = uint8(f >> 8)
		buf[offset+2] = uint8(f >> 16)
		buf[offset+3] = uint8(f >> 24)
		offset += 4
	}
	return size
}

// DecodeFloat32Slice decodes a []float32 from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeFloat32Slice, DecodeFloat32Slice will panic.
func DecodeFloat32Slice(buf []byte, v *[]float32) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(buf[n] & mask7)
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 7
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 14
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 21
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 28
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 35
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 42
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 49
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 56
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 63
	goto done
done:
	n++
	length = t
	if uint64(cap(*v)) >= length {
		*v = (*v)[:length]
	} else {
		*v = make([]float32, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		var f uint32
		f |= uint32(buf[offset+0])
		f |= uint32(buf[offset+1]) << 8
		f |= uint32(buf[offset+2]) << 16
		f |= uint32(buf[offset+3]) << 24
		(*v)[i] = *(*float32)(unsafe.Pointer(&f))
		offset += 4
	}
	return offset
}

// SizeofFloat32Slice takes a []float32 and returns the number of bytes.
func SizeofFloat32Slice(v []float32) uint64 {
	var length uint64 = uint64(len(v))
	return SizeofVarint(length) + length*4
}

// MaxFloat32SliceBytes returns maximum length of a []float32.
func MaxFloat32SliceBytes(v []float32) uint64 {
	var length uint64 = uint64(len(v))
	return 10 + length*4
}

// EncodeFloat64Slice encodes a []float64 into buf and returns the number of bytes written.
// If the buffer is too small, EncodeFloat64Slice will panic.
func EncodeFloat64Slice(buf []byte, v []float64) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(v))
	lengthSize := SizeofVarint(length)
	size = lengthSize + length*8
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, s := range v {
		f := *(*uint64)(unsafe.Pointer(&s))
		buf[offset+0] = uint8(f)
		buf[offset+1] = uint8(f >> 8)
		buf[offset+2] = uint8(f >> 16)
		buf[offset+3] = uint8(f >> 24)
		buf[offset+4] = uint8(f >> 32)
		buf[offset+5] = uint8(f >> 40)
		buf[offset+6] = uint8(f >> 48)
		buf[offset+7] = uint8(f >> 56)
		offset += 8
	}
	return size
}

// DecodeFloat64Slice decodes a uint64 from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeFloat64Slice, DecodeFloat64Slice will panic.
func DecodeFloat64Slice(buf []byte, v *[]float64) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(buf[n] & mask7)
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 7
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 14
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 21
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 28
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 35
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 42
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 49
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 56
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 63
	goto done
done:
	n++
	length = t
	if uint64(cap(*v)) >= length {
		*v = (*v)[:length]
	} else {
		*v = make([]float64, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		var f uint64
		f |= uint64(buf[offset+0])
		f |= uint64(buf[offset+1]) << 8
		f |= uint64(buf[offset+2]) << 16
		f |= uint64(buf[offset+3]) << 24
		f |= uint64(buf[offset+4]) << 32
		f |= uint64(buf[offset+5]) << 40
		f |= uint64(buf[offset+6]) << 48
		f |= uint64(buf[offset+7]) << 56
		(*v)[i] = *(*float64)(unsafe.Pointer(&f))
		offset += 8
	}
	return offset
}

// SizeofFloat64Slice takes a []float64 and returns the number of bytes.
func SizeofFloat64Slice(v []float64) uint64 {
	var length uint64 = uint64(len(v))
	return SizeofVarint(length) + length*8
}

// MaxFloat64SliceBytes returns maximum length of a []float64.
func MaxFloat64SliceBytes(v []float64) uint64 {
	var length uint64 = uint64(len(v))
	return 10 + length*8
}

// EncodeBoolSlice encodes a []bool into buf and returns the number of bytes written.
// If the buffer is too small, EncodeBoolSlice will panic.
func EncodeBoolSlice(buf []byte, v []bool) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(v))
	lengthSize := SizeofVarint(length)
	size = lengthSize + length
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, s := range v {
		if !s {
			buf[offset] = 0
		} else {
			buf[offset] = 1
		}
		offset++
	}
	return size
}

// DecodeBoolSlice decodes a []bool from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeBoolSlice, DecodeBoolSlice will panic.
func DecodeBoolSlice(buf []byte, v *[]bool) uint64 {
	var length uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(buf[n] & mask7)
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 7
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 14
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 21
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 28
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 35
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 42
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 49
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 56
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 63
	goto done
done:
	n++
	length = t
	if uint64(cap(*v)) >= length {
		*v = (*v)[:length]
	} else {
		*v = make([]bool, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		if buf[offset] == 0 {
			(*v)[i] = false
		} else {
			(*v)[i] = true
		}
		offset++
	}
	return offset
}

// SizeofBoolSlice takes a []bool and returns the number of bytes.
func SizeofBoolSlice(v []bool) uint64 {
	length := uint64(len(v))
	return SizeofVarint(length) + length
}

// MaxBoolSliceBytes returns maximum length of a []bool.
func MaxBoolSliceBytes(v []bool) uint64 {
	var length uint64 = uint64(len(v))
	return 10 + length
}

// EncodeStringSlice encodes a []string into buf and returns the number of bytes written.
// If the buffer is too small, EncodeStringSlice will panic.
func EncodeStringSlice(buf []byte, v []string) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(v))
	lengthSize := SizeofVarint(length)
	size = lengthSize
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, s := range v {
		length := uint64(len(s))
		lengthSize := SizeofVarint(length)
		size += lengthSize + length
		t := length
		for i := uint64(0); i < lengthSize-1; i++ {
			buf[offset+i] = byte(t) | msb7
			t >>= 7
		}
		buf[offset+lengthSize-1] = byte(t)
		copy(buf[offset+lengthSize:], s)
		offset += length + lengthSize
	}
	return size
}

// DecodeStringSlice decodes a []string from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeStringSlice, DecodeStringSlice will panic.
func DecodeStringSlice(buf []byte, v *[]string) uint64 {
	var length uint64
	var l uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(buf[n] & mask7)
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 7
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 14
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 21
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 28
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 35
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 42
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 49
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 56
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 63
	goto done
done:
	n++
	length = t
	if uint64(cap(*v)) >= length {
		*v = (*v)[:length]
	} else {
		*v = make([]string, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		b := buf[offset:]
		j := uint64(0)
		var t uint64
		t = uint64(b[j] & mask7)
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 7
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 14
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 21
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 28
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 35
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 42
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 49
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 56
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 63
		goto fordone
	fordone:
		j++
		l++
		s := buf[offset+j : offset+j+t]
		(*v)[i] = *(*string)(unsafe.Pointer(&s))
		offset += j + t
	}
	return offset
}

// SizeofStringSlice takes a []string and returns the number of bytes.
func SizeofStringSlice(v []string) uint64 {
	var size uint64
	length := uint64(len(v))
	sizeof := SizeofVarint(length)
	size = sizeof
	for _, s := range v {
		length := uint64(len(s))
		sizeof := SizeofVarint(length)
		size += sizeof + length
	}
	return size
}

// MaxStringSliceBytes returns maximum length of a []string.
func MaxStringSliceBytes(v []string) uint64 {
	var size uint64
	size = uint64(10)
	for _, s := range v {
		length := uint64(len(s))
		size += 10 + length
	}
	return size
}

// EncodeBytesSlice encodes a buf [][]byte into buf and returns the number of bytes written.
// If the buffer is too small, EncodeBytesSlice will panic.
func EncodeBytesSlice(buf []byte, v [][]byte) uint64 {
	var offset uint64
	var size uint64
	length := uint64(len(v))
	lengthSize := SizeofVarint(length)
	size = lengthSize
	t := length
	for i := uint64(0); i < lengthSize-1; i++ {
		buf[i] = byte(t) | msb7
		t >>= 7
	}
	buf[lengthSize-1] = byte(t)
	offset = lengthSize
	for _, s := range v {
		length := uint64(len(s))
		lengthSize := SizeofVarint(length)
		size += lengthSize + length
		t := length
		for i := uint64(0); i < lengthSize-1; i++ {
			buf[offset+i] = byte(t) | msb7
			t >>= 7
		}
		buf[offset+lengthSize-1] = byte(t)
		copy(buf[offset+lengthSize:], s)
		offset += length + lengthSize
	}
	return size
}

// DecodeBytesSlice decodes a [][]byte from buf, stores the result in the value pointed to by v
// and returns the number of bytes read (> 0).
// If the buffer is not from EncodeBytesSlice, DecodeBytesSlice will panic.
func DecodeBytesSlice(buf []byte, v *[][]byte) uint64 {
	var length uint64
	var l uint64
	var offset uint64
	var n uint64
	var t uint64
	t = uint64(buf[n] & mask7)
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 7
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 14
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 21
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 28
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 35
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 42
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 49
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 56
	if buf[n]&msb7 == 0 {
		goto done
	}
	n++
	t |= uint64(buf[n]&mask7) << 63
	goto done
done:
	n++
	length = t
	if uint64(cap(*v)) >= length {
		*v = (*v)[:length]
	} else {
		*v = make([][]byte, length)
	}
	offset = n
	for i := uint64(0); i < length; i++ {
		b := buf[offset:]
		j := uint64(0)
		var t uint64
		t = uint64(b[j] & mask7)
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 7
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 14
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 21
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 28
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 35
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 42
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 49
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 56
		if b[j]&msb7 == 0 {
			goto fordone
		}
		j++
		t |= uint64(b[j]&mask7) << 63
		goto fordone
	fordone:
		j++
		l++
		s := buf[offset+j : offset+j+t]
		(*v)[i] = s
		offset += j + t
	}
	return offset
}

// SizeofBytesSlice takes a [][]byte and returns the number of bytes.
func SizeofBytesSlice(v [][]byte) uint64 {
	var size uint64
	length := uint64(len(v))
	sizeof := SizeofVarint(length)
	size = sizeof
	for _, s := range v {
		length := uint64(len(s))
		sizeof := SizeofVarint(length)
		size += sizeof + length
	}
	return size
}

// MaxBytesSliceBytes returns maximum length of a [][]byte.
func MaxBytesSliceBytes(v [][]byte) uint64 {
	var size uint64
	size = uint64(10)
	for _, s := range v {
		length := uint64(len(s))
		size += 10 + length
	}
	return size
}
