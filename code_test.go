package code

import (
	"testing"
)

func TestCodeUint8(t *testing.T) {
	var buf =make([]byte,8)
	var v uint8=128
	var n uint64
	n=EncodeUint8(buf,v)
	var d uint8
	n=DecodeUint8(buf[:n],&d)
	if v!=d{
		t.Errorf("error %d != %d",v,d)
	}
	if n!=8{
		t.Errorf("error %d != %d",n,len(buf[:n]))
	}

}
func BenchmarkCodeUint8(b *testing.B) {
	var v uint8=128
	var n uint64
	var buf =make([]byte,8)
	n=EncodeUint8(buf,v)
	var v2 uint8
	DecodeUint8(buf[:n],&v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n=EncodeUint8(buf,v)
		DecodeUint8(buf[:n],&v2)
	}
}
func TestCodeUint16(t *testing.T) {
	var buf =make([]byte,8)
	var v uint16=128
	var n uint64
	n=EncodeUint16(buf,v)
	var d uint16
	n=DecodeUint16(buf[:n],&d)
	if v!=d{
		t.Errorf("error %d != %d",v,d)
	}
	if n!=8{
		t.Errorf("error %d != %d",n,len(buf[:n]))
	}

}

func BenchmarkCodeUint16(b *testing.B) {
	var v uint16=128
	var n uint64
	var buf =make([]byte,8)
	n=EncodeUint16(buf,v)
	var v2 uint16
	DecodeUint16(buf[:n],&v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n=EncodeUint16(buf,v)
		DecodeUint16(buf[:n],&v2)
	}
}

func TestCodeUint32(t *testing.T) {
	var buf =make([]byte,8)
	var v uint32=128
	var n uint64
	n=EncodeUint32(buf,v)
	var d uint32
	n=DecodeUint32(buf[:n],&d)
	if v!=d{
		t.Errorf("error %d != %d",v,d)
	}
	if n!=8{
		t.Errorf("error %d != %d",n,len(buf[:n]))
	}

}

func BenchmarkCodeUint32(b *testing.B) {
	var v uint32=128
	var n uint64
	var buf =make([]byte,8)
	n=EncodeUint32(buf,v)
	var v2 uint32
	DecodeUint32(buf[:n],&v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n=EncodeUint32(buf,v)
		DecodeUint32(buf[:n],&v2)
	}
}


func TestCodeUint64(t *testing.T) {
	var buf =make([]byte,8)
	var v uint64=128
	var n uint64
	n=EncodeUint64(buf,v)
	var d uint64
	n=DecodeUint64(buf[:n],&d)
	if v!=d{
		t.Errorf("error %d != %d",v,d)
	}
	if n!=8{
		t.Errorf("error %d != %d",n,len(buf[:n]))
	}
}

func BenchmarkCodeUint64(b *testing.B) {
	var v uint64=128
	var n uint64
	var buf =make([]byte,8)
	n=EncodeUint64(buf,v)
	var v2 uint64
	DecodeUint64(buf[:n],&v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n=EncodeUint64(buf,v)
		DecodeUint64(buf[:n],&v2)
	}
}


func TestCodeInt(t *testing.T) {
	var buf =make([]byte,9)
	var v uint64=128
	var n uint64
	n=EncodeInt(buf,v)
	var d uint64
	n=DecodeInt(buf[:n],&d)
	if v!=d{
		t.Errorf("error %d != %d",v,d)
	}
	if n!=SizeofInt(v){
		t.Errorf("error %d != %d",n,len(buf[:n]))
	}

}

func BenchmarkCodeInt(b *testing.B) {
	var v uint64=128
	var n uint64
	var buf =make([]byte,9)
	n=EncodeInt(buf,v)
	var v2 uint64
	DecodeInt(buf[:n],&v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n=EncodeInt(buf,v)
		DecodeInt(buf[:n],&v2)
	}
}

func TestCodeVarint(t *testing.T) {
	var buf =make([]byte,9)
	var v uint64=128
	var n uint64
	n=EncodeVarint(buf,v)
	var d uint64
	n=DecodeVarint(buf[:n],&d)
	if v!=d{
		t.Errorf("error %d != %d",v,d)
	}
	if n!=SizeofVarint(v){
		t.Errorf("error %d != %d",n,len(buf[:n]))
	}
}

func BenchmarkCodeVarint(b *testing.B) {
	var v uint64=128
	var n uint64
	var buf =make([]byte,9)
	n=EncodeVarint(buf,v)
	var v2 uint64
	DecodeVarint(buf[:n],&v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n=EncodeVarint(buf,v)
		DecodeVarint(buf[:n],&v2)
	}
}

func TestCodeFloat32(t *testing.T) {
	var f float32=3.1415926
	var n uint64
	var buf =make([]byte,4)
	n=EncodeFloat32(nil,f)
	var v float32
	n=DecodeFloat32(buf[:n],&v)
	if v!=f{
		t.Errorf("error %.7f != %.7f",v,f)
	}
	if n!=SizeofFloat32(){
		t.Errorf("error %d != %d",n,len(buf[:n]))
	}
}

func BenchmarkCodeFloat32(b *testing.B) {
	var f float32=3.1415926
	var n uint64
	var buf =make([]byte,4)
	n=EncodeFloat32(buf,f)
	var v float32
	DecodeFloat32(buf[:n],&v)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n=EncodeFloat32(buf,f)
		DecodeFloat32(buf[:n],&v)
	}
}
func TestCodeFloat64(t *testing.T) {
	var f float64=3.1415926
	var n uint64
	var buf =make([]byte,8)
	n=EncodeFloat64(buf,f)
	var v float64
	n=DecodeFloat64(buf[:n],&v)
	if v!=f{
		t.Errorf("error %.7f != %.7f",v,f)
	}
	if n!=SizeofFloat64(){
		t.Errorf("error %d != %d",n,len(buf[:n]))
	}
}

func BenchmarkCodeFloat64(b *testing.B) {
	var f float64=3.1415926
	var n uint64
	var buf =make([]byte,8)
	n=EncodeFloat64(buf,f)
	var v float64
	DecodeFloat64(buf[:n],&v)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n=EncodeFloat64(buf,f)
		DecodeFloat64(buf[:n],&v)
	}
}

func TestCodeBool(t *testing.T) {
	var buf =make([]byte,9)
	var v bool=true
	var n uint64
	n=EncodeBool(buf,v)
	var d bool
	n=DecodeBool(buf[:n],&d)
	if v!=d{
		t.Errorf("error %t != %t",v,d)
	}
	if n!=SizeofBool(){
		t.Errorf("error %d != %d",n,len(buf[:n]))
	}
	v=false
	n=EncodeBool(buf,v)
	n=DecodeBool(buf[:n],&d)
	if v!=d{
		t.Errorf("error %t != %t",v,d)
	}
	if n!=SizeofBool(){
		t.Errorf("error %d != %d",n,len(buf[:n]))
	}
}

func BenchmarkCodeBool(b *testing.B) {
	var buf =make([]byte,1)
	var v bool=true
	var n uint64
	n=EncodeBool(buf,v)
	var v2 bool
	DecodeBool(buf[:n],&v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n=EncodeBool(buf,v)
		DecodeBool(buf[:n],&v2)
	}
}

func TestCodeString(t *testing.T) {
	var buf =make([]byte,64)
	var v string="HelloWorld"
	var n uint64
	n=EncodeString(buf,v)
	var d string
	n=DecodeString(buf[:n],&d)
	if v!=d{
		t.Errorf("error %s != %s",v,d)
	}
	if n!=SizeofString(v){
		t.Errorf("error %d != %d",n,len(buf[:n]))
	}
}

func BenchmarkCodeString(b *testing.B) {
	var buf =make([]byte,64)
	var v string="h"
	var n uint64
	n=EncodeString(buf,v)
	var v2 string
	DecodeString(buf[:n],&v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n=EncodeString(buf,v)
		DecodeString(buf[:n],&v2)
	}
}

func TestCodeBytes(t *testing.T) {
	var buf =make([]byte,64)
	var v []byte=[]byte{1}
	var n uint64
	n=EncodeBytes(buf,v)
	var d =make([]byte,2)
	n=DecodeBytes(buf[:n],&d)
	if v[0]!=d[0]{
		t.Errorf("error %d != %d",v[0],d[0])
	}
	if n!=SizeofBytes(v){
		t.Errorf("error %d != %d",n,len(buf[:n]))
	}
}

func BenchmarkCodeBytes(b *testing.B) {
	var buf =make([]byte,64)
	var v []byte=[]byte{1}
	var n uint64
	n=EncodeBytes(buf,v)
	var d =make([]byte,1)
	DecodeBytes(buf[:n],&d)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n=EncodeBytes(buf,v)
		DecodeBytes(buf[:n],&d)
	}
}

func TestCodeSliceBytes(t *testing.T) {
	var buf =make([]byte,64)
	var v [][]byte=[][]byte{{1},{2}}
	var n uint64
	n=EncodeSliceBytes(buf,v)
	var v2 [][]byte=make([][]byte,2)
	n=DecodeSliceBytes(buf[:n],&v2)
	if v[0][0]!=v2[0][0]{
		t.Errorf("error %d != %d",v[0][0],v2[0][0])
	}
	if n!=SizeofSliceBytes(v){
		t.Errorf("error %d != %d",n,len(buf[:n]))
	}
}

func BenchmarkCodeSliceBytes(b *testing.B) {
	var buf =make([]byte,64)
	var v [][]byte=[][]byte{{1},{2}}
	var n uint64
	n=EncodeSliceBytes(buf,v)
	var v2 [][]byte=make([][]byte,2)
	DecodeSliceBytes(buf[:n],&v2)
	b.SetBytes(int64(len(buf[:n])))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n=EncodeSliceBytes(buf,v)
		DecodeSliceBytes(buf[:n],&v2)
	}
}


//Sizeof
func BenchmarkSizeofUint8(b *testing.B) {
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofUint8()
	}
}
func BenchmarkSizeofUint16(b *testing.B) {
	b.SetBytes(int64(2))
	for i := 0; i < b.N; i++ {
		SizeofUint16()
	}
}
func BenchmarkSizeofUint32(b *testing.B) {
	b.SetBytes(int64(4))
	for i := 0; i < b.N; i++ {
		SizeofUint32()
	}
}

func BenchmarkSizeofUint64(b *testing.B) {
	b.SetBytes(int64(8))
	for i := 0; i < b.N; i++ {
		SizeofUint64()
	}
}


func BenchmarkSizeofInt(b *testing.B) {
	var v uint64 =128
	b.SetBytes(int64(2))
	for i := 0; i < b.N; i++ {
		SizeofInt(v)
	}
}

func BenchmarkSizeofVarint(b *testing.B) {
	b.SetBytes(int64(2))
	var v uint64 =128
	for i := 0; i < b.N; i++ {
		SizeofVarint(v)
	}
}

func BenchmarkSizeofFloat32(b *testing.B) {
	b.SetBytes(int64(4))
	for i := 0; i < b.N; i++ {
		SizeofFloat32()
	}
}
func BenchmarkSizeofFloat64(b *testing.B) {
	b.SetBytes(int64(8))
	for i := 0; i < b.N; i++ {
		SizeofFloat64()
	}
}

func BenchmarkSizeofBool(b *testing.B) {
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofBool()
	}
}

func BenchmarkSizeofString(b *testing.B) {
	var v string="h"
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofString(v)
	}
}

func BenchmarkSizeofBytes(b *testing.B) {
	var v []byte=[]byte{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofBytes(v)
	}
}

func BenchmarkSizeofSliceBytes(b *testing.B) {
	var v [][]byte=[][]byte{{1},{2}}
	b.SetBytes(int64(2))
	for i := 0; i < b.N; i++ {
		SizeofSliceBytes(v)
	}
}