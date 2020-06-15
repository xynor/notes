## 原子操作
* 原子操作只有两种状态，没做和全部做完。不应该被看到中间状态。
* 现代计算机保证了基本操作是原子的，比如获取/存储16位，32位，64位操作数。  
* 原子操作本身没有锁的语义，其锁的特点是在汇编里的LOCK指令。  
* 并发问题的根源在于写不可见。读出得数已经不是最新值。  
[博文](https://zhuanlan.zhihu.com/p/26159285)
## 例子分析
* atomic.AddInt64
````
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
````
   * 反编译
   ````
        0x0027 00039 (cas_test.go:15)   PCDATA  $0, $0
        0x0027 00039 (cas_test.go:15)   LOCK
        0x0028 00040 (cas_test.go:15)   XADDQ   AX, (CX)
        0x002c 00044 (<unknown line number>)    NOP
   ````
   * 可以看到atomic.AddInt64(&k, 1) 翻译成汇编是 XADDQ   AX, (CX)，虽然是一条指令，但是也包含了多步。
   读内存，将值加一，写回内存。这在其他cpu上看来并不是一个原子的操作。
   * 可以看到，在调用XADDQ指令的时候，加上了**LOCK**指令来锁总线，使得其他cpu不能读取和修改内存，进行内存屏障，从而达到原子操作。
* CAS
````
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
````
   * 反编译
   ````
        0x0022 00034 (cas_test.go:30)   MOVQ    (CX), DX
        0x0025 00037 (cas_test.go:31)   LEAQ    1(DX), BX
        0x0029 00041 (cas_test.go:31)   MOVQ    DX, AX
        0x002c 00044 (cas_test.go:31)   LOCK
        0x002d 00045 (cas_test.go:31)   CMPXCHGQ        BX, (CX)
        0x0031 00049 (cas_test.go:31)   SETEQ   DL
        0x0034 00052 (cas_test.go:31)   TESTB   DL, DL
   ````
   * 同样可以看到使用**LOCK**进行了锁总线
   * 在Load操作时，并没有LOCK操作，cpu对于基本字节对齐读是原子操作的。cpu的[缓存一致性](https://blog.csdn.net/mingwulipo/article/details/89845156) 可以保证读出的数是内存中最新的。
   * 在CompareAndSwap操作时则需要对上一步取出的值与内存中的值进行比较，不一致说明有其他线程修改过。一致则替换为新值。
   * 使用CAS操作会产生ABA问题，即线程一将值修改为A，线程二将值修改为B，线程三又将值修改为A，这时线程四看到的值仍然是A认为没有修改过。解决办法是加入版本号控制(比如64位，高位值，低位版本)。