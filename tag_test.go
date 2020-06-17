// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package code

import (
	"bytes"
	"testing"
)

func TestTag(t *testing.T) {
	var fieldNumber = 7
	var wireType = Varint
	var n uint64
	b := []byte{7<<3 | 0} // field 7, wire type 0
	tag := MakeTag(fieldNumber, wireType)
	buf := make([]byte, MaxVarintBytes(tag))
	n = EncodeVarint(buf, tag)
	if !bytes.Equal(b, buf[:n]) {
		t.Errorf("Bytes: %v != %v", b, buf[:n])
	}
	f := GetTagFieldNumber(tag)
	if fieldNumber != f {
		t.Errorf("FieldNumber: %d != %d", fieldNumber, f)
	}
	w := GetTagWireType(tag)
	if wireType != w {
		t.Errorf("WireType: %d != %d", wireType, w)
	}
}

func TestGetWireType(t *testing.T) {
	type A struct{}
	var v = A{}
	var a = []interface{}{int32(1024), int64(1024), uint32(1024), uint64(1024), true, float64(3.1415926), "HelloWorld", v, []byte{1, 2, 3}, []A{v}, float32(3.14), int(1024)}
	var b = []WireType{Varint, Varint, Varint, Varint, Varint, Fixed64, LengthDelimited, LengthDelimited, LengthDelimited, LengthDelimited, Fixed32, Invalid}
	for i := 0; i < len(a); i++ {
		if b[i] != GetWireType(a[i]) {
			t.Errorf("%d!=%d", b[i], GetWireType(a[i]))
		}
	}
}
