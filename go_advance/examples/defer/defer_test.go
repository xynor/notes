package main

import (
	"fmt"
	"testing"
	"time"
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

func Test_PanicRecover(t *testing.T) {
	fmt.Println("c")
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
		fmt.Println("e")
	}()
	f()              //开始调用f
	fmt.Println("f") //这里开始下面代码不会再执行
}

func f() {
	fmt.Println("a")
	panic("异常信息")
	fmt.Println("b") //这里开始下面代码不会再执行
}

func Test_DeferS(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Timer Call")
			defer func() {
				fmt.Println("Deferrrr")
			}()
		}
	}
}
