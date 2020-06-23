## 总结  
[博文1](https://www.jianshu.com/p/9cfcac638147)  [博文2](https://www.jianshu.com/p/040d5f1698ec)
* Go中是支持指针的。包括多级指针。使用中不要超过2级指针。
* 指针不支持指针运算。虽然可以通过uintptr来hack，但是因为Go在GC的时候可能将内存搬运。从而造成不可遇见错误。
* 指针也是有类型的。
* unsafe.Pointer类似c中的void *p。
* 接口 `*interface{}` 或 `*Car` 是不能取地址的。所以不能出现在入参和出参中。