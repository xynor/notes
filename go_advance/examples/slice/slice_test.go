package main

import (
	"fmt"
	"reflect"
	"testing"
)

type IntSlice []int

func (sp *IntSlice) appendElemP() {
	*sp = append(*sp, 123, 456, 789)
	fmt.Printf("appendElemP elem:%v,ptr:%p,len:%d,cap:%d\n", *sp, &(*sp)[0], len(*sp), cap(*sp))
}

func (sp IntSlice) appendElem() {
	sp = append(sp, 123, 456, 789)
	fmt.Printf("appendElem elem:%v,ptr:%p,len:%d,cap:%d\n", sp, &sp[0], len(sp), cap(sp))
}

func Test_S(t *testing.T) {
	//创建len为8，cap为10的slice
	a := make([]int, 8, 10)
	for k, _ := range a {
		a[k] = k
	}
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", a, &a[0], len(a), cap(a))
	fmt.Println("a[len(a):]", a[len(a):])
	//取a的第2个到第3个共2个元素，左闭右开
	a = a[2:4]
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", a, &a[0], len(a), cap(a))
	a = append(a, 88, 99)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", a, &a[0], len(a), cap(a))
	a = a[:cap(a)]
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", a, &a[0], len(a), cap(a))
	a = append(a, 44, 55, 66)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", a, &a[0], len(a), cap(a))
	appendElem(a)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", a, &a[0], len(a), cap(a))
	/*
		elem:[0 1 2 3 4 5 6 7],ptr:0xc0000b8000,len:8,cap:10
		elem:[2 3],ptr:0xc0000b8010,len:2,cap:8
		elem:[2 3 88 99],ptr:0xc0000b8010,len:4,cap:8
		elem:[2 3 88 99 6 7 0 0],ptr:0xc0000b8010,len:8,cap:8
		elem:[2 3 88 99 6 7 0 0 44 55 66],ptr:0xc0000c6000,len:11,cap:16
		appendElem elem:[2 3 88 99 6 7 0 0 44 55 66 123 456 789],ptr:0xc0000c6000,len:14,cap:16
		elem:[2 3 88 99 6 7 0 0 44 55 66],ptr:0xc0000c6000,len:11,cap:16
	*/
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := array[:]
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", b, &b[0], len(b), cap(b))
	sliceElem(b)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", b, &b[0], len(b), cap(b))
	deleteElem(b, 4)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", b, &b[0], len(b), cap(b))
	changeElem(b, 4)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", b, &b[0], len(b), cap(b))
	appendElem(b)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", b, &b[0], len(b), cap(b))
	appendElemP(&b)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", b, &b[0], len(b), cap(b))
	/*
		elem:[0 1 2 3 4 5 6 7 8 9],ptr:0xc0000b8050,len:10,cap:10
		sliceElem elem:[2 3 4 5 6],ptr:0xc0000b8060,len:5,cap:8
		elem:[0 1 2 3 4 5 6 7 8 9],ptr:0xc0000b8050,len:10,cap:10
		deleteElem elem:[0 1 2 3 5 6 7 8 9],ptr:0xc0000b8050,len:9,cap:10
		elem:[0 1 2 3 5 6 7 8 9 9],ptr:0xc0000b8050,len:10,cap:10
		changeElem elem:[0 1 2 3 1000 6 7 8 9 9],ptr:0xc0000b8050,len:10,cap:10
		elem:[0 1 2 3 1000 6 7 8 9 9],ptr:0xc0000b8050,len:10,cap:10
		appendElem elem:[0 1 2 3 1000 6 7 8 9 9 123 456 789],ptr:0xc0000c8000,len:13,cap:20
		elem:[0 1 2 3 1000 6 7 8 9 9],ptr:0xc0000b8050,len:10,cap:10
		appendElemP elem:[0 1 2 3 1000 6 7 8 9 9 123 456 789],ptr:0xc0000c80a0,len:13,cap:20
		elem:[0 1 2 3 1000 6 7 8 9 9 123 456 789],ptr:0xc0000c80a0,len:13,cap:20
	*/
	is := make(IntSlice, 10)
	for k, _ := range is {
		is[k] = k
	}
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", is, &is[0], len(is), cap(is))
	is.appendElem()
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", is, &is[0], len(is), cap(is))
	is.appendElemP()
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", is, &is[0], len(is), cap(is))
	/*
		elem:[0 1 2 3 4 5 6 7 8 9],ptr:0xc0000ba0a0,len:10,cap:10
		appendElem elem:[0 1 2 3 4 5 6 7 8 9 123 456 789],ptr:0xc0000ca140,len:13,cap:20
		elem:[0 1 2 3 4 5 6 7 8 9],ptr:0xc0000ba0a0,len:10,cap:10
		appendElemP elem:[0 1 2 3 4 5 6 7 8 9 123 456 789],ptr:0xc0000ca1e0,len:13,cap:20
		elem:[0 1 2 3 4 5 6 7 8 9 123 456 789],ptr:0xc0000ca1e0,len:13,cap:20

	*/
	data := make([]byte, 0, 32)
	data = append(data, []byte("bbbbbbbbbb")...)
	fmt.Printf("%p,%v\n", data, data)
	data = data[:0]
	fmt.Printf("%p,%v\n", data, data)

}

func sliceElem(s []int) {
	s = s[2:7]
	fmt.Printf("sliceElem elem:%v,ptr:%p,len:%d,cap:%d\n", s, &s[0], len(s), cap(s))
}

func deleteElem(s []int, k int) {
	s = append(s[:k], s[k+1:]...)
	fmt.Printf("deleteElem elem:%v,ptr:%p,len:%d,cap:%d\n", s, &s[0], len(s), cap(s))
}

func changeElem(s []int, k int) {
	s[k] = 1000
	fmt.Printf("changeElem elem:%v,ptr:%p,len:%d,cap:%d\n", s, &s[0], len(s), cap(s))
}

func appendElem(s []int) {
	s = append(s, 123, 456, 789)
	fmt.Printf("appendElem elem:%v,ptr:%p,len:%d,cap:%d\n", s, &s[0], len(s), cap(s))
}

func appendElemP(s *[]int) {
	*s = append(*s, 123, 456, 789)
	fmt.Printf("appendElemP elem:%v,ptr:%p,len:%d,cap:%d\n", *s, &(*s)[0], len(*s), cap(*s))
}

func Test_Slice(t *testing.T) {
	arr := make([]int, 5, 10)
	for k, _ := range arr {
		arr[k] = k
	}
	fmt.Printf("arr:%p\n", arr)
	s := arr[3:10]
	fmt.Printf("s:%p,s:%v\n", s, s)
	s[2] = 99
	fmt.Printf("b:%v\n", arr)
}

func Test_COPY(t *testing.T) {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[:]
	fmt.Printf("array:%p,slice:%p\n", &array[0], slice)
	///
	s := make([]int, 8, 10)
	for k, _ := range s {
		s[k] = k
	}
	arr := [5]int{}
	fmt.Printf("arr[0:3],%v", arr[0:3])
	copy(arr[:], s)
	fmt.Printf("arr:%p,slice:%p\n", &arr[0], s)
}

func Test_Array(t *testing.T) {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("type:%v\n", reflect.TypeOf(array))
	fmt.Printf("%v,type:%v\n", array[0:2], reflect.TypeOf(array[0:2]))
}

func passArray(a [5]int) {
	fmt.Printf("a:%p,a:%v\n", &a[0], a)
	a[0] = 100
}

func passArrayP(a *[5]int) {
	fmt.Printf("a:%p,a:%v\n", &a[0], a)
	a[0] = 100
}

func Test_Pass(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("slice:%p\n", slice)
	b := [5]int{}
	copy(b[:], slice)
	fmt.Printf("b:%p\n", &b[0])
	passArray(b)
	fmt.Printf("b:%v\n", b)
}

func Test_PassP(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("slice:%p\n", slice)
	b := [5]int{}
	copy(b[:], slice)
	fmt.Printf("b:%p\n", &b[0])
	passArrayP(&b)
	fmt.Printf("b:%v\n", b)
}
