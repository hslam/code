package code

import (
	"testing"
)
func TestInterfaceTypeString0(t *testing.T) {
	typ,n,_:=InterfaceType("")
	if typ!=String||n!=0{
		t.Errorf("error %d %d",typ,n)
	}
}
func BenchmarkInterfaceTypeString0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InterfaceType("")
	}
}
func TestInterfaceTypeString(t *testing.T) {
	typ,n,_:=InterfaceType("123456789")
	if typ!=String||n!=1{
		t.Errorf("error %d %d",typ,n)
	}
}
func BenchmarkInterfaceTypeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InterfaceType("123456789")
	}
}
func TestStringType(t *testing.T) {
	n,b:=StringType("123456789")
	if n!=1{
		t.Errorf("error %d",n)
	}
	var c CodeType
	c,n=Type(b)
	if n!=1{
		t.Errorf("error %d",n)
	}
	if c!=String{
		t.Errorf("error %d",c)
	}
}
func BenchmarkStringType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringType("123456789")
	}
}
func TestInterfaceTypeInt0(t *testing.T) {
	typ,n,_:=InterfaceType(uint64(0))
	if typ!=Int||n!=0{
		t.Errorf("error %d %d",typ,n)
	}
}
func BenchmarkInterfaceTypeInt0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InterfaceType(uint64(0))
	}
}
func TestInterfaceTypeInt(t *testing.T) {
	typ,n,_:=InterfaceType(uint64(256))
	if typ!=Int||n!=2{
		t.Errorf("error %d %d",typ,n)
	}
}
func BenchmarkInterfaceTypeInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InterfaceType(uint64(256))
	}
}
func TestIntType(t *testing.T) {
	n,b:=IntType(uint64(256))
	if n!=2{
		t.Errorf("error %d",n)
	}
	var c CodeType
	c,n=Type(b)
	if n!=2{
		t.Errorf("error %d",n)
	}
	if c!=Int{
		t.Errorf("error %d",c)
	}
}
func BenchmarkIntType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntType(uint64(256))
	}
}
func TestInterfaceTypeTrue(t *testing.T) {
	typ,n,_:=InterfaceType(true)
	if typ!=True||n!=0{
		t.Errorf("error %d %d",typ,n)
	}
}
func BenchmarkInterfaceTypeTrue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InterfaceType(true)
	}
}
func TestTrueType(t *testing.T) {
	n,b:=TrueType()
	if n!=0{
		t.Errorf("error %d",n)
	}
	var c CodeType
	c,n=Type(b)
	if n!=0{
		t.Errorf("error %d",n)
	}
	if c!=True{
		t.Errorf("error %d",c)
	}
}
func BenchmarkTrueType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrueType()
	}
}

func TestInterfaceTypeFalse(t *testing.T) {
	typ,n,_:=InterfaceType(false)
	if typ!=False||n!=0{
		t.Errorf("error %d %d",typ,n)
	}
}
func BenchmarkInterfaceTypeFalse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InterfaceType(false)
	}
}
func TestFalseType(t *testing.T) {
	n,b:=FalseType()
	if n!=0{
		t.Errorf("error %d",n)
	}
	var c CodeType
	c,n=Type(b)
	if n!=0{
		t.Errorf("error %d",n)
	}
	if c!=False{
		t.Errorf("error %d",c)
	}
}
func BenchmarkFalseType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FalseType()
	}
}
func TestSizeInt(t *testing.T) {
	n:=SizeInt(uint64(256))
	if n!=2{
		t.Errorf("error %d != 2",n)
	}
}
func BenchmarkSizeInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var j = uint(i%64)
		SizeInt(uint64(1<<j))
	}
}
func TestCodeInt(t *testing.T) {
	var i uint64=256
	data:=EncodeInt(nil,i)
	v,_:=DecodeInt(data)
	if v!=i{
		t.Errorf("error %d != %d",v,i)
	}
}
func BenchmarkCodeInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data:=EncodeInt(nil,uint64(i))
		DecodeInt(data)
	}
}
func BenchmarkCodeIntWithBuffer(b *testing.B) {
	var buf =make([]byte,9)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeInt(buf,uint64(i))
		DecodeInt(data)
	}
}

func TestSizeVarint(t *testing.T) {
	n:=SizeVarint(uint64(128))
	if n!=2{
		t.Errorf("error %d != 2",n)
	}
}
func BenchmarkSizeVarint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var j = uint(i%64)
		SizeVarint(uint64(1<<j))
	}
}
func TestCodeVarint(t *testing.T) {
	data:=EncodeVarint(nil,10000)
	v,_:=DecodeVarint(data)
	if v!=10000{
		t.Errorf("error %d != 10000",v)
	}
}
func BenchmarkCodeVarint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data:=EncodeVarint(nil,uint64(i))
		DecodeVarint(data)
	}
}
func BenchmarkCodeVarintWithBuffer(b *testing.B) {
	var buf =make([]byte,10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeVarint(buf,uint64(i))
		DecodeVarint(data)
	}
}
func TestFloatType(t *testing.T) {
	n,b:=FloatType(4)
	if n!=4{
		t.Errorf("error %d",n)
	}
	var c CodeType
	c,n=Type(b)
	if n!=4{
		t.Errorf("error %d",n)
	}
	if c!=Float{
		t.Errorf("error %d",c)
	}
	n,b=FloatType(8)
	if n!=8{
		t.Errorf("error %d",n)
	}
	c,n=Type(b)
	if n!=8{
		t.Errorf("error %d",n)
	}
	if c!=Float{
		t.Errorf("error %d",c)
	}
}
func TestCodeFloat32(t *testing.T) {
	var f float32=3.1415926
	data:=EncodeFloat32(nil,f)
	v:=DecodeFloat32(data)
	if v!=f{
		t.Errorf("error %.7f != %.7f",v,f)
	}
}
func BenchmarkCodeFloat32(b *testing.B) {
	var f float32=3.1415926
	for i := 0; i < b.N; i++ {
		data:=EncodeFloat32(nil,f)
		DecodeFloat32(data)
	}
}
func BenchmarkCodeFloat32WithBuffer(b *testing.B) {
	var f float32=3.1415926
	var buf =make([]byte,5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeFloat32(buf,f)
		DecodeFloat32(data)
	}
}
func TestCodeFloat64(t *testing.T) {
	var f float64=3.1415926
	data:=EncodeFloat64(nil,f)
	v:=DecodeFloat64(data)
	if v!=f{
		t.Errorf("error %.7f != %.7f",v,f)
	}
}
func BenchmarkCodeFloat64(b *testing.B) {
	var f float64=3.1415926
	for i := 0; i < b.N; i++ {
		data:=EncodeFloat64(nil,f)
		DecodeFloat64(data)
	}
}
func BenchmarkCodeFloat64WithBuffer(b *testing.B) {
	var f float64=3.1415926
	var buf =make([]byte,9)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data:=EncodeFloat64(buf,f)
		DecodeFloat64(data)
	}
}
