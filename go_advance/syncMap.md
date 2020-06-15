## 并发安全的Map
* 由于map本身不是并发安全的。所以官方提供了sync.Map。  
* map本身配合读写锁，互斥锁能达到并发安全。但是效率低。[Benchmark](https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c)
* sync.Map使用了读写分离的缓存设计。读的map使用了lock-free的设计。  
[博文](https://zhuanlan.zhihu.com/p/44585993)  
## 数据结构
````
    type Map struct {
        mu sync.Mutex
        read atomic.Value // readOnly
        dirty map[interface{}]*entry
        misses int
    }
    
    type readOnly struct {
        m       map[interface{}]*entry
        amended bool // true if the dirty map contains some key not in m.
    }

    type entry struct {
        p unsafe.Pointer // *interface{}
    }
````
   * read的map使用atomic.Value来保证读写的[原子操作](./cas.md)。  
   * dirty则是built-in的Map，使用独占锁保护读写的一致性。
   * 两个map中的value都是*entry，这是因为删除map不是线程安全的。通过判断指针atomic.CompareAndSwapPointer来标识可以达到lock-free。其有两个
   状态nil，和expunged。  
   * read Map命中失败次数多后，将dirty Map提升为read Map。所以整个sync.Map有机会缩小。  
## 思考，学习
   * 读写分离提高效率。引入了读写分离缓存也就引入了以下问题  
    1. 一致性:有写操作时，使用互斥锁保证读写操作可见。  
    2. 原子操作做lock-free。  