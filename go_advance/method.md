## 总结
* T必须是用type重新定义过的，不限于struct,也可以是基础类型。
* T不能是type定义的指针。
* T和Method的定义要在一个包里面。
* T不能是接口类型。
* T和*T作为Receiver都是值传递，会copy。
* T为Receiver，改变不会传出来。
* *T为Receiver，改变会传出来。
* T调用`*T`的Receiver，会将&T传入(t `*T`)。
* `*T`调用 `*T`的Receiver，则直接传入(t `*T`)。
* `*T`调用 `T`的Receiver，会将`*T`传入(t `T`)

## 例子
````
type Stu struct {
	Name string
	Age  int
}

func (s Stu) SetName(name string) {
	fmt.Printf("in s:%p\n", &s)
	s.Name = name
}

func (s *Stu) SetNameP(name string) {
	fmt.Printf("in sp:%p\n", s)
	s.Name = name
}

func Test_Pass(t *testing.T) {
	s := Stu{}
	fmt.Printf("out s:%p\n", &s)
	s.SetName("Wang")
	fmt.Println("s:", s)
	s.SetNameP("Wang")
	fmt.Println("s:", s)
}
/*
out s:0xc0000be080
in s:0xc0000be0a0
s: { 0}
in sp:0xc0000be080
s: {Wang 0}
*/
````
## 例子
````
func Test_PassP(t *testing.T) {
	s := &Stu{}
	fmt.Printf("out s:%p\n", s)
	s.SetName("Wang")
	fmt.Println("s:", s)
	s.SetNameP("Wang")
	fmt.Println("s:", s)
}
/*
out s:0xc00000c0e0
in s:0xc00000c100
s: &{ 0}
in sp:0xc00000c0e0
s: &{Wang 0}
*/
````