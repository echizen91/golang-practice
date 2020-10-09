package pointers

import "fmt"

// PlayWithPointer : Assign different kind of pointers
func PlayWithPointer() {
	var s1 string = `some strings`
	var i1 int32 = 12
	var i2 int64 = 24
	var f1 float32 = 4.123
	var f2 float64 = 233.235874742
	var b1 bool = true

	fmt.Printf("s1 = %v; i1 = %v; i2 = %v; f1 = %v; f2 = %v; b1 = %v;\n", s1, i1, i2, f1, f2, b1)

	var s1p *string = &s1
	var i1p *int32 = &i1
	var i2p *int64 = &i2
	var f1p *float32 = &f1
	var f2p *float64 = &f2
	var b1p *bool = &b1

	fmt.Printf("Addresses s1 = %v; i1 = %v; i2 = %v; f1 = %v; f2 = %v; b1 = %v;\n", s1p, i1p, i2p, f1p, f2p, b1p)
	fmt.Printf("Dereferencing Addresses s1 = %v; i1 = %v; i2 = %v; f1 = %v; f2 = %v; b1 = %v;\n", *s1p, *i1p, *i2p, *f1p, *f2p, *b1p)

	// Change values in my pointer
	*s1p = `Hello World, I've changed`
	fmt.Printf("Changed s1p using pointer: s1p = %v with s1p address = %v\n", *s1p, s1p)
}
