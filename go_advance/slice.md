## 总结
* slice类似于变长数组，在长度小于1024时按照两倍容量增长，大于1024时按照1/4增长。容量增长后会发生内存分配和拷贝。
* slice的传递是值传递。但是操作元素仍然有效，因为使用的是指向底层数组的指针。当然也可以使用指向slice的指针。
* slice的切割本质上是指针的移动以及长度和容量的改变，所以需要注意GC。
* 使用append函数会进行内存拷贝。
* 使用make创建的slice返回的也是值。
* 左闭右开[,)比较奇怪的是a[len(a)]会越界，但是a[len(a):]却可以运行，返回空slice。
* 数组array可以直接通过array[:]变成切片，但是切片变成数组需要copy函数。数组是值传递，当然也可以用指针*array。
## 使用建议
* `cars := make([]Car, 2) cars = append(cars, b, bmw) `这种代码要注意cars的长度为4，前两个接口为空。
* 将slice当作入参进行长度，容量修改时最好使用函数返回值接受改变后的slice。
* 如果从文件读取了大的slice而只需要其中很小的片段，可以在考虑GC的时候将小片段装入新的slice，以保证大的slice能被GC。
* 可能的话尽量指定slice的cap,以减少内存分配和拷贝。
## Slice 结构
恰如 [slice-intro](https://blog.golang.org/slices-intro) 所示，slice实际上是数组的抽象。定义在runtime/slice.go中
````
    type slice struct {
        array unsafe.Pointer    //指向底层数组的第一个元素的地址
        len   int               //slice的长度
        cap   int               //slice的容量
    }
````
![slice-1](images/slice-1.png)
## 思考如下问题
* 以下代码的输出结果
````
	//创建len为8，cap为10的slice
	a := make([]int, 8, 10)
	for k, _ := range a {
		a[k] = k
	}
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", a, &a[0], len(a), cap(a))
	//取a的第2个到第3个共2个元素，左闭右开
	a = a[2:4]
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", a, &a[0], len(a), cap(a))
	a = append(a, 88, 99)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", a, &a[0], len(a), cap(a))
	a = a[:cap(a)]
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", a, &a[0], len(a), cap(a))
	a = append(a, 44, 55, 66)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", a, &a[0], len(a), cap(a))
	appendElem(a)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", a, &a[0], len(a), cap(a))
	/*
		elem:[0 1 2 3 4 5 6 7],ptr:0xc0000b8000,len:8,cap:10
		elem:[2 3],ptr:0xc0000b8010,len:2,cap:8
		elem:[2 3 88 99],ptr:0xc0000b8010,len:4,cap:8
		elem:[2 3 88 99 6 7 0 0],ptr:0xc0000b8010,len:8,cap:8
		elem:[2 3 88 99 6 7 0 0 44 55 66],ptr:0xc0000c6000,len:11,cap:16
		appendElem elem:[2 3 88 99 6 7 0 0 44 55 66 123 456 789],ptr:0xc0000c6000,len:14,cap:16
		elem:[2 3 88 99 6 7 0 0 44 55 66],ptr:0xc0000c6000,len:11,cap:16
	*/
````
从上面的例子可以看出
    1. cap代表了slice的容量，在进行内存分配的时候就确定了。
    2. 在进行slice的切割后，cap值变为cap-slice.array指针移动量。
    3. 切割后的slice，仍然可以grow。
    4. 使用append可以改变slice大小，在超过原来的cap时会重新分配并拷贝内存。
---
* 以下代码的输出结果
````
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := array[:]
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", b, &b[0], len(b), cap(b))
	sliceElem(b)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", b, &b[0], len(b), cap(b))
	deleteElem(b, 4)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", b, &b[0], len(b), cap(b))
	changeElem(b, 4)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", b, &b[0], len(b), cap(b))
	appendElem(b)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", b, &b[0], len(b), cap(b))
	appendElemP(&b)
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", b, &b[0], len(b), cap(b))
	/*
		elem:[0 1 2 3 4 5 6 7 8 9],ptr:0xc0000b8050,len:10,cap:10
		sliceElem elem:[2 3 4 5 6],ptr:0xc0000b8060,len:5,cap:8
		elem:[0 1 2 3 4 5 6 7 8 9],ptr:0xc0000b8050,len:10,cap:10
		deleteElem elem:[0 1 2 3 5 6 7 8 9],ptr:0xc0000b8050,len:9,cap:10
		elem:[0 1 2 3 5 6 7 8 9 9],ptr:0xc0000b8050,len:10,cap:10
		changeElem elem:[0 1 2 3 1000 6 7 8 9 9],ptr:0xc0000b8050,len:10,cap:10
		elem:[0 1 2 3 1000 6 7 8 9 9],ptr:0xc0000b8050,len:10,cap:10
		appendElem elem:[0 1 2 3 1000 6 7 8 9 9 123 456 789],ptr:0xc0000c8000,len:13,cap:20
		elem:[0 1 2 3 1000 6 7 8 9 9],ptr:0xc0000b8050,len:10,cap:10
		appendElemP elem:[0 1 2 3 1000 6 7 8 9 9 123 456 789],ptr:0xc0000c80a0,len:13,cap:20
		elem:[0 1 2 3 1000 6 7 8 9 9 123 456 789],ptr:0xc0000c80a0,len:13,cap:20
	*/
func sliceElem(s []int) {
	s = s[2:7]
	fmt.Printf("sliceElem elem:%v,ptr:%p,len:%d,cap:%d\n", s, &s[0], len(s), cap(s))
}
func deleteElem(s []int, k int) {
	s = append(s[:k], s[k+1:]...)
	fmt.Printf("deleteElem elem:%v,ptr:%p,len:%d,cap:%d\n", s, &s[0], len(s), cap(s))
}
func changeElem(s []int, k int) {
	s[k] = 1000
	fmt.Printf("changeElem elem:%v,ptr:%p,len:%d,cap:%d\n", s, &s[0], len(s), cap(s))
}
func appendElem(s []int) {
	s = append(s, 123, 456, 789)
	fmt.Printf("appendElem elem:%v,ptr:%p,len:%d,cap:%d\n", s, &s[0], len(s), cap(s))
}
func appendElemP(s *[]int) {
	*s = append(*s, 123, 456, 789)
	fmt.Printf("appendElemP elem:%v,ptr:%p,len:%d,cap:%d\n", *s, &(*s)[0], len(*s), cap(*s))
}
````
从上面的例子可以看出
    1. slice的传递是值传递。即函数调用后slice的array，len，cap不会改变。
    2. 使用append来删除元素，是通过slice的array指针修改底层数组(内存copy)。
    3. 可以使用指针来改变实参。
---
* 以下代码的输出结果
````
type IntSlice []int
func (sp *IntSlice) appendElemP() {
	*sp = append(*sp, 123, 456, 789)
	fmt.Printf("appendElemP elem:%v,ptr:%p,len:%d,cap:%d\n", *sp, &(*sp)[0], len(*sp), cap(*sp))
}
func (sp IntSlice) appendElem() {
	sp = append(sp, 123, 456, 789)
	fmt.Printf("appendElem elem:%v,ptr:%p,len:%d,cap:%d\n", sp, &sp[0], len(sp), cap(sp))
}
````
````
	is := make(IntSlice, 10)
	for k, _ := range is {
		is[k] = k
	}
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", is, &is[0], len(is), cap(is))
	is.appendElem()
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", is, &is[0], len(is), cap(is))
	is.appendElemP()
	fmt.Printf("elem:%v,ptr:%p,len:%d,cap:%d\n", is, &is[0], len(is), cap(is))
	/*
		elem:[0 1 2 3 4 5 6 7 8 9],ptr:0xc0000ba0a0,len:10,cap:10
		appendElem elem:[0 1 2 3 4 5 6 7 8 9 123 456 789],ptr:0xc0000ca140,len:13,cap:20
		elem:[0 1 2 3 4 5 6 7 8 9],ptr:0xc0000ba0a0,len:10,cap:10
		appendElemP elem:[0 1 2 3 4 5 6 7 8 9 123 456 789],ptr:0xc0000ca1e0,len:13,cap:20
		elem:[0 1 2 3 4 5 6 7 8 9 123 456 789],ptr:0xc0000ca1e0,len:13,cap:20

	*/
````
从上面的例子也可以看出指针和值的区别  
[完整代码](https://github.com/xinxuwang/notes/blob/master/go_advance/examples/slice/main.go)