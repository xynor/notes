package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func Benchmark_Add(b *testing.B) {
	var k int64
	var lock sync.WaitGroup
	for i := 0; i < b.N; i++ {
		lock.Add(1)
		go func() {
			atomic.AddInt64(&k, 1)
			lock.Done()
		}()
	}
	lock.Wait()
	b.Log(k, b.N)
}

func Benchmark_Cas(b *testing.B) {
	var k int64
	var lock sync.WaitGroup
	for i := 0; i < b.N; i++ {
		lock.Add(1)
		go func() {
			for {
				current := atomic.LoadInt64(&k)
				if atomic.CompareAndSwapInt64(&k, current, current+1) {
					break
				}
			}
			lock.Done()
		}()
	}
	lock.Wait()
	b.Log(k, b.N)
}

func Benchmark_Lock(b *testing.B) {
	var k int64
	var lock sync.WaitGroup
	var syncLock sync.Mutex
	for i := 0; i < b.N; i++ {
		lock.Add(1)
		go func() {
			syncLock.Lock()
			k++
			syncLock.Unlock()
			lock.Done()
		}()
	}
	lock.Wait()
	b.Log(k, b.N)
}

func Benchmark_NoLock(b *testing.B) {
	var k int64
	var lock sync.WaitGroup
	for i := 0; i < b.N; i++ {
		lock.Add(1)
		go func() {
			k++
			lock.Done()
		}()
	}
	lock.Wait()
	b.Log(k, b.N)
}

type AValue struct {
	a int64
	b string
}

func Benchmark_Value(b *testing.B) {
	var v atomic.Value
	v.Store(33)
	var lock sync.WaitGroup
	for i := 0; i < b.N; i++ {
		lock.Add(1)
		go func() {
			a, ok := v.Load().(int)
			if !ok {
				b.Error("Not ok")
			}
			v.Store(a)
			lock.Done()
		}()
	}
	lock.Wait()
	b.Log(v.Load(), b.N)
}
