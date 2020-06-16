## 总结
* 用于快速遍历数组(字符串)、切片、Map和Channel。
* for range a {} 遍历。不关心索引和数据。执行len(a)次。
* for i := range a {} 遍历。只关心索引。
* for i, elem := range a {} 遍历。关心索引和数据。
* for v := range ch {} 阻塞形死循环，其他携程调用close(ch)，循环结束。
* 遍历字符串时，v会返回rune类型并解码。
* v是临时变量。不应该使用a:=&v赋值。
* 遍历Map是无序的。
## 注意
* **_不要在循环中改变slice，map可能会出乎意料的结果_**。
## 遍历字符串
````
func Test_Strings(t *testing.T) {
	s := "ABCDE哈哈一个"
	t.Log("len(s):", len(s))
	for k, v := range s {
		t.Log(k, v)
	}
	r := []rune(s)
	t.Log("len(r):", len(r))
	for k, v := range r {
		t.Log(k, string(v))
	}
}

    Test_Strings: forRange_test.go:9: len(s): 17
    Test_Strings: forRange_test.go:11: 0 65
    Test_Strings: forRange_test.go:11: 1 66
    Test_Strings: forRange_test.go:11: 2 67
    Test_Strings: forRange_test.go:11: 3 68
    Test_Strings: forRange_test.go:11: 4 69
    Test_Strings: forRange_test.go:11: 5 21704
    Test_Strings: forRange_test.go:11: 8 21704
    Test_Strings: forRange_test.go:11: 11 19968
    Test_Strings: forRange_test.go:11: 14 20010
    Test_Strings: forRange_test.go:14: len(r): 9
    Test_Strings: forRange_test.go:16: 0 A
    Test_Strings: forRange_test.go:16: 1 B
    Test_Strings: forRange_test.go:16: 2 C
    Test_Strings: forRange_test.go:16: 3 D
    Test_Strings: forRange_test.go:16: 4 E
    Test_Strings: forRange_test.go:16: 5 哈
    Test_Strings: forRange_test.go:16: 6 哈
    Test_Strings: forRange_test.go:16: 7 一
    Test_Strings: forRange_test.go:16: 8 个
````
   * 可以看出v返回的是rune，并且k也做过相应的改变。utf-8，一个中文占3个字节。
## 遍历Map
````
func Test_Map(t *testing.T) {
	m := map[int]int{}
	for i := 0; i < 10; i++ {
		m[i] = i
	}

	for _, v := range m {
		if v == 8 {
			k := rand.Intn(100) + 10
			m[k] = 8
			fmt.Println("insert:", k)
		}
	}
}
````
   * 多运行几次可以看到输出insert的次数不一样。