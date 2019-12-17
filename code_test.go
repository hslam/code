package code

import (
	"testing"
)

func TestCodeInt(t *testing.T) {
	var buf =make([]byte,9)
	var v uint64=128
	data:=EncodeInt(buf,v)
	d,n:=DecodeInt(data)
	if v!=d{
		t.Errorf("error %d != %d",v,d)
	}
	if n!=SizeofInt(v){
		t.Errorf("error %d != %d",n,len(data))
	}

}
func BenchmarkCodeInt(b *testing.B) {
	var v uint64=128
	data:=EncodeInt(nil,v)
	DecodeInt(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeInt(nil,v)
		DecodeInt(data)
	}
}
func BenchmarkCodeIntWithBuffer(b *testing.B) {
	var v uint64=128
	var buf =make([]byte,9)
	data:=EncodeInt(buf,v)
	DecodeInt(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeInt(buf,v)
		DecodeInt(data)
	}
}

func TestCodeVarint(t *testing.T) {
	var buf =make([]byte,9)
	var v uint64=128
	data:=EncodeVarint(buf,v)
	d,n:=DecodeVarint(data)
	if v!=d{
		t.Errorf("error %d != %d",v,d)
	}
	if n!=SizeofVarint(v){
		t.Errorf("error %d != %d",n,len(data))
	}
}
func BenchmarkCodeVarint(b *testing.B) {
	var v uint64=128
	data:=EncodeVarint(nil,v)
	DecodeVarint(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeVarint(nil,v)
		DecodeVarint(data)
	}
}
func BenchmarkCodeVarintWithBuffer(b *testing.B) {
	var v uint64=128
	var buf =make([]byte,9)
	data:=EncodeVarint(buf,v)
	DecodeVarint(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeVarint(buf,v)
		DecodeVarint(data)
	}
}


func TestCodeFloat32(t *testing.T) {
	var f float32=3.1415926
	data:=EncodeFloat32(nil,f)
	v,n:=DecodeFloat32(data)
	if v!=f{
		t.Errorf("error %.7f != %.7f",v,f)
	}
	if n!=SizeofFloat32(){
		t.Errorf("error %d != %d",n,len(data))
	}
}
func BenchmarkCodeFloat32(b *testing.B) {
	var f float32=3.1415926
	data:=EncodeFloat32(nil,f)
	DecodeFloat32(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeFloat32(nil,f)
		DecodeFloat32(data)
	}
}
func BenchmarkCodeFloat32WithBuffer(b *testing.B) {
	var f float32=3.1415926
	var buf =make([]byte,4)
	data:=EncodeFloat32(buf,f)
	DecodeFloat32(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeFloat32(buf,f)
		DecodeFloat32(data)
	}
}
func TestCodeFloat64(t *testing.T) {
	var f float64=3.1415926
	data:=EncodeFloat64(nil,f)
	v,n:=DecodeFloat64(data)
	if v!=f{
		t.Errorf("error %.7f != %.7f",v,f)
	}
	if n!=SizeofFloat64(){
		t.Errorf("error %d != %d",n,len(data))
	}
}
func BenchmarkCodeFloat64(b *testing.B) {
	var f float64=3.1415926
	data:=EncodeFloat64(nil,f)
	DecodeFloat64(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeFloat64(nil,f)
		DecodeFloat64(data)
	}
}
func BenchmarkCodeFloat64WithBuffer(b *testing.B) {
	var f float64=3.1415926
	var buf =make([]byte,8)
	data:=EncodeFloat64(buf,f)
	DecodeFloat64(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeFloat64(buf,f)
		DecodeFloat64(data)
	}
}


func TestCodeBool(t *testing.T) {
	var buf =make([]byte,9)
	var v bool=true
	data:=EncodeBool(buf,v)
	d,n:=DecodeBool(data)
	if v!=d{
		t.Errorf("error %t != %t",v,d)
	}
	if n!=SizeofBool(){
		t.Errorf("error %d != %d",n,len(data))
	}
	v=false
	data=EncodeBool(buf,v)
	d,n=DecodeBool(data)
	if v!=d{
		t.Errorf("error %t != %t",v,d)
	}
	if n!=SizeofBool(){
		t.Errorf("error %d != %d",n,len(data))
	}
}
func BenchmarkCodeBool(b *testing.B) {
	var v bool=true
	data:=EncodeBool(nil,v)
	DecodeBool(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeBool(nil,v)
		DecodeBool(data)
	}
}
func BenchmarkCodeBoolWithBuffer(b *testing.B) {
	var buf =make([]byte,1)
	var v bool=true
	data:=EncodeBool(buf,v)
	DecodeBool(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeBool(buf,v)
		DecodeBool(data)
	}
}

func TestCodeString(t *testing.T) {
	var buf =make([]byte,64)
	var v string="HelloWorld"
	data:=EncodeString(buf,v)
	d,n:=DecodeString(data)
	if v!=d{
		t.Errorf("error %s != %s",v,d)
	}
	if n!=SizeofString(v){
		t.Errorf("error %d != %d",n,len(data))
	}
}
func BenchmarkCodeString(b *testing.B) {
	var v string="HelloWorld"
	data:=EncodeString(nil,v)
	DecodeString(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeString(nil,v)
		DecodeString(data)
	}
}
func BenchmarkCodeStringWithBuffer(b *testing.B) {
	var buf =make([]byte,64)
	var v string="HelloWorld"
	data:=EncodeString(buf,v)
	DecodeString(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeString(buf,v)
		DecodeString(data)
	}
}

func TestCodeBytes(t *testing.T) {
	var buf =make([]byte,64)
	var v []byte=[]byte{1}
	data:=EncodeBytes(buf,v)
	d,n:=DecodeBytes(data)
	if v[0]!=d[0]{
		t.Errorf("error %d != %d",v[0],d[0])
	}
	if n!=SizeofBytes(v){
		t.Errorf("error %d != %d",n,len(data))
	}
}
func BenchmarkCodeBytes(b *testing.B) {
	var v []byte=[]byte{1,2,3,4,5,6,7,8,9,0}
	data:=EncodeBytes(nil,v)
	DecodeBytes(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeBytes(nil,v)
		DecodeBytes(data)
	}
}
func BenchmarkCodeBytesWithBuffer(b *testing.B) {
	var buf =make([]byte,64)
	var v []byte=[]byte{1,2,3,4,5,6,7,8,9,0}
	data:=EncodeBytes(buf,v)
	DecodeBytes(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeBytes(buf,v)
		DecodeBytes(data)
	}
}

func TestCodeSliceBytes(t *testing.T) {
	var buf =make([]byte,64)
	var v [][]byte=[][]byte{{1},{2}}
	data:=EncodeSliceBytes(buf,v)
	d,n:=DecodeSliceBytes(data)
	if v[0][0]!=d[0][0]{
		t.Errorf("error %d != %d",v[0][0],d[0][0])
	}
	if n!=SizeofSliceBytes(v){
		t.Errorf("error %d != %d",n,len(data))
	}
}
func BenchmarkCodeSliceBytes(b *testing.B) {
	var v [][]byte=[][]byte{{1,2},{3,4}}
	data:=EncodeSliceBytes(nil,v)
	DecodeSliceBytes(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeSliceBytes(nil,v)
		DecodeSliceBytes(data)
	}
}
func BenchmarkCodeSliceBytesWithBuffer(b *testing.B) {
	var buf =make([]byte,64)
	var v [][]byte=[][]byte{{1,2},{3,4}}
	data:=EncodeSliceBytes(buf,v)
	DecodeSliceBytes(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeSliceBytes(buf,v)
		DecodeSliceBytes(data)
	}
}
