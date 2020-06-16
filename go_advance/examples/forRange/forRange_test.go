package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_Strings(t *testing.T) {
	s := "ABCDE哈哈一个"
	t.Log("len(s):", len(s))
	for k, v := range s {
		t.Log(k, v)
	}
	r := []rune(s)
	t.Log("len(r):", len(r))
	for k, v := range r {
		t.Log(k, string(v))
	}
}

func Test_Map(t *testing.T) {
	m := map[int]int{}
	for i := 0; i < 10; i++ {
		m[i] = i
	}

	for _, v := range m {
		if v == 8 {
			k := rand.Intn(100) + 10
			m[k] = 8
			fmt.Println("insert:", k)
		}
	}
}
