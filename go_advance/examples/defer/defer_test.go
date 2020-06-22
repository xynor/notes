package main

import (
	"fmt"
	"testing"
)

func Call() (i int, e error) {
	defer func() {
		i++
		e = fmt.Errorf("Defer error:%v\n", e)
	}()
	return i + 10, fmt.Errorf("call error")
}

func Test_A(t *testing.T) {
	i, e := Call()
	fmt.Println(i, e)
}

//TODO panic recover
