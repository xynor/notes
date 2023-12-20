package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
	"unsafe"
)

func Test_Complete(t *testing.T) {
	//a := int32(10)
	a := int32(-10)
	ap := &a
	apointer := unsafe.Pointer(ap)
	apointer0 := (*byte)(apointer)
	fmt.Printf("apointer0:%x\n", *apointer0)
	fmt.Printf("apointer1:%x\n", *(*byte)(unsafe.Pointer(uintptr(apointer) + 1)))
	fmt.Printf("apointer2:%x\n", *(*byte)(unsafe.Pointer(uintptr(apointer) + 2)))
	fmt.Printf("apointer3:%x\n", *(*byte)(unsafe.Pointer(uintptr(apointer) + 3)))
	return
}

func Test_Complete2(t *testing.T) {
	//a := int32(10)
	a := int16(-10)
	ap := &a
	apointer := unsafe.Pointer(ap)
	apointer0 := (*byte)(apointer)
	fmt.Printf("apointer0:%x\n", *apointer0)
	fmt.Printf("apointer1:%x\n", *(*byte)(unsafe.Pointer(uintptr(apointer) + 1)))
	return
}

func Test_Complete3(t *testing.T) {
	//a := int32(10)
	a := int16(-11)
	b := uint16(10)
	c := uint16(a) + b
	fmt.Println(c)
	return
}

func Test_Rshift(t *testing.T) {
	a := uint8(10)          //0000_1010
	fmt.Println("a:", a>>1) //0000_0101

	b := int8(-4) //1111_1100
	fmt.Printf("b:%x\n", *(*byte)(unsafe.Pointer(&b)))
	rb := b >> 1
	fmt.Println("shift:", rb) //1111_1110
	fmt.Printf("rB:%x\n", *(*byte)(unsafe.Pointer(&rb)))
	return
}

func Test_Lshift(t *testing.T) {
	a := uint8(10)          //0000_1010
	fmt.Println("a:", a<<1) //000_01010

	b := int8(-4) //1111_1100
	fmt.Printf("b:%x\n", *(*byte)(unsafe.Pointer(&b)))
	lb := b << 1
	fmt.Println("shift:", lb) //1111_1000
	fmt.Printf("lB:%x\n", *(*byte)(unsafe.Pointer(&lb)))

	c := int8(-65) //1011_1111
	fmt.Printf("c:%x\n", *(*byte)(unsafe.Pointer(&c)))
	lc := c << 1 //0111_1110
	fmt.Printf("lc:%d, lc:%x\n", lc, *(*byte)(unsafe.Pointer(&lc)))

	d := int8(64) //0100_0000
	fmt.Printf("d:%x\n", *(*byte)(unsafe.Pointer(&d)))
	ld := d << 1 //1000_0000
	fmt.Printf("ld:%d, ld:%x\n", ld, *(*byte)(unsafe.Pointer(&ld)))
	return
}

func Test_Float(t *testing.T) {
	a := decimal.New(1, -18)
	b := decimal.NewFromFloat32(3.10)
	fmt.Println("b:", b)
	c := a.Mul(b)
	fmt.Println(c)
	aa := 0.1
	fmt.Println(aa + 0.2)
	d := decimal.NewFromFloat(0.1)
	fmt.Println(d.Add(decimal.NewFromFloat(0.2)))
}
