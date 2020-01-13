// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package code

//WireType defines the type of wire.
type WireType uint8

const (
	//Varint is used for int32, int64, uint32, uint64, sint32, sint64, bool, enum.
	Varint WireType = 0
	//Fixed64 is used for fixed64, sfixed64, double.
	Fixed64 WireType = 1
	//LengthDelimited is used for string, bytes, embedded messages, packed repeated fields.
	LengthDelimited WireType = 2
	//StartGroup is used for groups (deprecated).
	StartGroup WireType = 3
	//EndGroup is used for groups (deprecated).
	EndGroup WireType = 4
	//Fixed32 is used for fixed32, sfixed32, float.
	Fixed32 WireType = 5
)

//GetTagFieldNumber given a tag value, returns the field number (the upper 29 bits).
func GetTagFieldNumber(tag uint64) int {
	return int(tag >> 3)
}

//GetTagWireType given a tag value, returns the wire type (lower 3 bits).
func GetTagWireType(tag uint64) WireType {
	return WireType(tag & 0x7)
}

//MakeTag makes a tag value given a field number and wire type.
func MakeTag(fieldNumber int, wireType WireType) uint64 {
	return uint64(fieldNumber<<3) | uint64(wireType)
}
