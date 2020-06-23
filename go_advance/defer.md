## 总结
[博文](https://www.cnblogs.com/makelu/p/11226974.html)
* defer类型函数调用压栈。后压先执行。
* defer在压栈时会被求值。
* defer的函数如果是个闭包或者有指针，虽然压栈时被取值，但是仍然会被影响。
* 在循环中谨慎使用defer，除非清楚知道如何执行。
* defer可以修改函数error值，这在返回统一格式的error时非常有用。