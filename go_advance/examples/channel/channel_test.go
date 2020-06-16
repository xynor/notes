package main

import (
	"testing"
	"time"
)

func Test_Pass(t *testing.T) {
	ch := make(chan int, 10)
	stop := make(chan struct{})
	t.Logf("len:%v,cap:%v\n", len(ch), cap(ch))
	go func() {
		a := <-ch
		t.Logf("Read:%p\n", &a)
		stop <- struct{}{}
	}()
	aa := 20
	t.Logf("aa:%p\n", &aa)
	ch <- aa
	<-stop
}

func Test_SelectWrite(t *testing.T) {
	ch := make(chan int)
	go func() {
		for {
			a := <-ch
			t.Logf("Read:%p\n", &a)
			time.Sleep(3 * time.Second)
		}
	}()
	go func() {
		for {
			select {
			case ch <- 20:
				t.Logf("Write:\n")
			}
		}
	}()
	select {}
}
