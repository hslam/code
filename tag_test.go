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
	b := []byte{7<<3 | 0} // field 7, wire type 6
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
