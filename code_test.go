package code

import (
	"testing"
)

func TestCodeUint8(t *testing.T) {
	var buf =make([]byte,8)
	var v uint8=128
	data:=EncodeUint8(buf,&v)
	var d uint8
	n:=DecodeUint8(data,&d)
	if v!=d{
		t.Errorf("error %d != %d",v,d)
	}
	if n!=8{
		t.Errorf("error %d != %d",n,len(data))
	}

}
func BenchmarkCodeUint8(b *testing.B) {
	var v uint8=128
	var buf =make([]byte,8)
	data:=EncodeUint8(buf,&v)
	var v2 uint8
	DecodeUint8(data,&v2)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeUint8(buf,&v)
		DecodeUint8(data,&v2)
	}
}
func TestCodeUint16(t *testing.T) {
	var buf =make([]byte,8)
	var v uint16=128
	data:=EncodeUint16(buf,&v)
	var d uint16
	n:=DecodeUint16(data,&d)
	if v!=d{
		t.Errorf("error %d != %d",v,d)
	}
	if n!=8{
		t.Errorf("error %d != %d",n,len(data))
	}

}

func BenchmarkCodeUint16(b *testing.B) {
	var v uint16=128
	var buf =make([]byte,8)
	data:=EncodeUint16(buf,&v)
	var v2 uint16
	DecodeUint16(data,&v2)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeUint16(buf,&v)
		DecodeUint16(data,&v2)
	}
}

func TestCodeUint32(t *testing.T) {
	var buf =make([]byte,8)
	var v uint32=128
	data:=EncodeUint32(buf,&v)
	var d uint32
	n:=DecodeUint32(data,&d)
	if v!=d{
		t.Errorf("error %d != %d",v,d)
	}
	if n!=8{
		t.Errorf("error %d != %d",n,len(data))
	}

}

func BenchmarkCodeUint32(b *testing.B) {
	var v uint32=128
	var buf =make([]byte,8)
	data:=EncodeUint32(buf,&v)
	var v2 uint32
	DecodeUint32(data,&v2)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeUint32(buf,&v)
		DecodeUint32(data,&v2)
	}
}


func TestCodeUint64(t *testing.T) {
	var buf =make([]byte,8)
	var v uint64=128
	data:=EncodeUint64(buf,&v)
	var d uint64
	n:=DecodeUint64(data,&d)
	if v!=d{
		t.Errorf("error %d != %d",v,d)
	}
	if n!=8{
		t.Errorf("error %d != %d",n,len(data))
	}

}

func BenchmarkCodeUint64(b *testing.B) {
	var v uint64=128
	var buf =make([]byte,8)
	data:=EncodeUint64(buf,&v)
	var v2 uint64
	DecodeUint64(data,&v2)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeUint64(buf,&v)
		DecodeUint64(data,&v2)
	}
}


func TestCodeInt(t *testing.T) {
	var buf =make([]byte,9)
	var v uint64=128
	data:=EncodeInt(buf,&v)
	var d uint64
	n:=DecodeInt(data,&d)
	if v!=d{
		t.Errorf("error %d != %d",v,d)
	}
	if n!=SizeofInt(&v){
		t.Errorf("error %d != %d",n,len(data))
	}

}

func BenchmarkCodeInt(b *testing.B) {
	var v uint64=128
	var buf =make([]byte,9)
	data:=EncodeInt(buf,&v)
	var v2 uint64
	DecodeInt(data,&v2)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeInt(buf,&v)
		DecodeInt(data,&v2)
	}
}

func TestCodeVarint(t *testing.T) {
	var buf =make([]byte,9)
	var v uint64=128
	data:=EncodeVarint(buf,&v)
	var d uint64
	n:=DecodeVarint(data,&d)
	if v!=d{
		t.Errorf("error %d != %d",v,d)
	}
	if n!=SizeofVarint(&v){
		t.Errorf("error %d != %d",n,len(data))
	}
}

func BenchmarkCodeVarint(b *testing.B) {
	var v uint64=128
	var buf =make([]byte,9)
	data:=EncodeVarint(buf,&v)
	var v2 uint64
	DecodeVarint(data,&v2)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeVarint(buf,&v)
		DecodeVarint(data,&v2)
	}
}

func TestCodeFloat32(t *testing.T) {
	var f float32=3.1415926
	data:=EncodeFloat32(nil,&f)
	var v float32
	n:=DecodeFloat32(data,&v)
	if v!=f{
		t.Errorf("error %.7f != %.7f",v,f)
	}
	if n!=SizeofFloat32(){
		t.Errorf("error %d != %d",n,len(data))
	}
}

func BenchmarkCodeFloat32(b *testing.B) {
	var f float32=3.1415926
	var buf =make([]byte,4)
	data:=EncodeFloat32(buf,&f)
	var v float32
	DecodeFloat32(data,&v)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeFloat32(buf,&f)
		DecodeFloat32(data,&v)
	}
}
func TestCodeFloat64(t *testing.T) {
	var f float64=3.1415926
	data:=EncodeFloat64(nil,&f)
	var v float64
	n:=DecodeFloat64(data,&v)
	if v!=f{
		t.Errorf("error %.7f != %.7f",v,f)
	}
	if n!=SizeofFloat64(){
		t.Errorf("error %d != %d",n,len(data))
	}
}

func BenchmarkCodeFloat64(b *testing.B) {
	var f float64=3.1415926
	var buf =make([]byte,8)
	data:=EncodeFloat64(buf,&f)
	var v float64
	DecodeFloat64(data,&v)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeFloat64(buf,&f)
		DecodeFloat64(data,&v)
	}
}

func TestCodeBool(t *testing.T) {
	var buf =make([]byte,9)
	var v bool=true
	data:=EncodeBool(buf,&v)
	var d bool
	n:=DecodeBool(data,&d)
	if v!=d{
		t.Errorf("error %t != %t",v,d)
	}
	if n!=SizeofBool(){
		t.Errorf("error %d != %d",n,len(data))
	}
	v=false
	data=EncodeBool(buf,&v)
	n=DecodeBool(data,&d)
	if v!=d{
		t.Errorf("error %t != %t",v,d)
	}
	if n!=SizeofBool(){
		t.Errorf("error %d != %d",n,len(data))
	}
}

func BenchmarkCodeBool(b *testing.B) {
	var buf =make([]byte,1)
	var v bool=true
	data:=EncodeBool(buf,&v)
	var v2 bool
	DecodeBool(data,&v2)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeBool(buf,&v)
		DecodeBool(data,&v2)
	}
}

func TestCodeString(t *testing.T) {
	var buf =make([]byte,64)
	var v string="HelloWorld"
	data:=EncodeString(buf,&v)
	var d string
	n:=DecodeString(data,&d)
	if v!=d{
		t.Errorf("error %s != %s",v,d)
	}
	if n!=SizeofString(&v){
		t.Errorf("error %d != %d",n,len(data))
	}
}

func BenchmarkCodeString(b *testing.B) {
	var buf =make([]byte,64)
	var v string="h"
	data:=EncodeString(buf,&v)
	var v2 string
	DecodeString(data,&v2)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeString(buf,&v)
		DecodeString(data,&v2)
	}
}

func TestCodeBytes(t *testing.T) {
	var buf =make([]byte,64)
	var v []byte=[]byte{1}
	data:=EncodeBytes(buf,&v)
	var d =make([]byte,2)
	n:=DecodeBytes(data,&d)
	if v[0]!=d[0]{
		t.Errorf("error %d != %d",v[0],d[0])
	}
	if n!=SizeofBytes(&v){
		t.Errorf("error %d != %d",n,len(data))
	}
}

func BenchmarkCodeBytes(b *testing.B) {
	var buf =make([]byte,64)
	var v []byte=[]byte{1}
	data:=EncodeBytes(buf,&v)
	var d =make([]byte,1)
	DecodeBytes(data,&d)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeBytes(buf,&v)
		DecodeBytes(data,&d)
	}
}

func TestCodeSliceBytes(t *testing.T) {
	var buf =make([]byte,64)
	var v [][]byte=[][]byte{{1},{2}}
	data:=EncodeSliceBytes(buf,&v)
	var v2 [][]byte=make([][]byte,2)
	n:=DecodeSliceBytes(data,&v2)
	if v[0][0]!=v2[0][0]{
		t.Errorf("error %d != %d",v[0][0],v2[0][0])
	}
	if n!=SizeofSliceBytes(&v){
		t.Errorf("error %d != %d",n,len(data))
	}
}

func BenchmarkCodeSliceBytes(b *testing.B) {
	var buf =make([]byte,64)
	var v [][]byte=[][]byte{{1},{2}}
	data:=EncodeSliceBytes(buf,&v)
	var v2 [][]byte=make([][]byte,2)
	DecodeSliceBytes(data,&v2)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeSliceBytes(buf,&v)
		DecodeSliceBytes(data,&v2)
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
		SizeofInt(&v)
	}
}

func BenchmarkSizeofVarint(b *testing.B) {
	b.SetBytes(int64(2))
	var v uint64 =128
	for i := 0; i < b.N; i++ {
		SizeofVarint(&v)
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
		SizeofString(&v)
	}
}

func BenchmarkSizeofBytes(b *testing.B) {
	var v []byte=[]byte{1}
	b.SetBytes(int64(1))
	for i := 0; i < b.N; i++ {
		SizeofBytes(&v)
	}
}

func BenchmarkSizeofSliceBytes(b *testing.B) {
	var v [][]byte=[][]byte{{1},{2}}
	b.SetBytes(int64(2))
	for i := 0; i < b.N; i++ {
		SizeofSliceBytes(&v)
	}
}