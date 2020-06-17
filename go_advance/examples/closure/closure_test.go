package main

import (
	"fmt"
	"testing"
)

type Stu struct {
	Name string
	Age  int
}

func Test_ChangeData(t *testing.T) {
	a := 4
	fmt.Printf("a:%p\n", &a)
	b := 2
	s := Stu{
		Name: "Stu",
		Age:  0,
	}
	p := new(int)
	fmt.Printf("p:%p\n", p)
	f := func() {
		fmt.Printf("in a:%p\n", &a)
		fmt.Printf("in p:%p\n", p)
		a++
		b++
		s.Name = "Wang"
		*p++
	}
	f()
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("s:", s)
	fmt.Println("p:", *p)
	f()
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("s:", s)
	fmt.Println("p:", *p)
}

func k(i int) func() int {
	fmt.Printf("k i:%p\n", &i) //在这里new i
	return func() int {
		fmt.Printf("i:%p\n", &i)
		i++
		return i
	}
}

func Test_D(t *testing.T) {
	a := 3
	fmt.Printf("a:%p\n", &a)
	f := k(a)
	d := f()
	fmt.Println("d:", d)
	e := f()
	fmt.Println("e:", e)
}

func Test_C(t *testing.T) {
	f := k(2)
	a := f()
	f2 := k(4)
	b := f2()
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	c := k(1)()
	fmt.Println("c:", c)
}

func B(i *int) func() int {
	return func() int {
		fmt.Printf("in fun i:%p\n", i)
		*i++
		return *i
	}
}
func Test_B(t *testing.T) {
	a := 3
	fmt.Printf("a:%p\n", &a)
	f := B(&a)
	f()
	fmt.Println(a)
}
