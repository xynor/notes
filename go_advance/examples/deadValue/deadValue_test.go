package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func Test_A(t *testing.T) {
	running := true
	go func() {
		println("start thread1")
		count := 1
		for running {
			count++
		}
		println("end thread1: count =", count) // 这句代码永远执行不到为什么？
	}()
	go func() {
		println("start thread2")
		for {
			running = false
			//break
		}
	}()
	time.Sleep(time.Hour)
}

func Test_B(t *testing.T) {
	running := atomic.Bool{}
	running.Store(true)
	go func() {
		println("start thread1")
		count := 1
		for running.Load() {
			count++
		}
		println("end thread1: count =", count)
	}()
	go func() {
		println("start thread2")
		for {
			running.Store(false)
			//break
		}
	}()
	time.Sleep(time.Hour)
}
