package main

import (
	"fmt"
	"testing"
)

type S int

func (s S) SetValue(a int) {

}

type Stu struct {
	Name string
	Age  int
}

func (s Stu) SetName(name string) {
	fmt.Printf("in s:%p\n", &s)
	s.Name = name
}

func (s *Stu) SetNameP(name string) {
	fmt.Printf("in sp:%p\n", s)
	s.Name = name
}

func Test_Pass(t *testing.T) {
	s := Stu{}
	fmt.Printf("out s:%p\n", &s)
	s.SetName("Wang")
	fmt.Println("s:", s)
	s.SetNameP("Wang")
	fmt.Println("s:", s)
}

func Test_PassP(t *testing.T) {
	s := &Stu{}
	fmt.Printf("out s:%p\n", s)
	s.SetName("Wang")
	fmt.Println("s:", s)
	s.SetNameP("Wang")
	fmt.Println("s:", s)
}
