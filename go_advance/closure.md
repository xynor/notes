## 总结
* 闭包是有数据的方法，结构体是有方法的数据。
* 闭包函数体中包含的体外的数据，在闭包中是指针，改变其值会影响外边。
* 闭包最大的使用场景可能是非侵入式编程了。
* 闭包会翻译成如下
* [使用场景](https://www.calhoun.io/5-useful-ways-to-use-closures-in-go/)
```
type Closure struct {
    F func()() 
    i *int
    s *struct
}
```
