package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type Car interface {
	Run(int) error
	GetInfo() (Info, error)
}
type Info struct {
	Amount int
	Height int
}
type Benz struct {
	Info
}

var _ Car = (*Benz)(nil)

func (b *Benz) GetInfo() (Info, error) {
	return b.Info, nil
}

func (b *Benz) Run(speed int) error {
	fmt.Println("Benz Run at speed:", speed)
	return nil
}

type BMW struct {
	Info
}

func (b BMW) Run(i int) error {
	fmt.Println("BMW Run at speed:", i)
	return nil
}

func (b BMW) GetInfo() (Info, error) {
	return b.Info, nil
}

var _ Car = (*BMW)(nil)

type Combine interface {
	Car
	Trans()
}

func (b BMW) Trans() {
	fmt.Println("BMW TRANS")
}

func Test_Type(t *testing.T) {
	var c = BMW{
		Info: Info{
			Amount: 100,
			Height: 100,
		},
	}
	tye := reflect.TypeOf(c)
	fmt.Println(tye.Kind())
	tye = reflect.TypeOf(&c)
	fmt.Println(tye.Kind())
	pc := &c
	tye = reflect.TypeOf(&pc)
	fmt.Println(tye.Kind())
	//即使是传递指针的指针，得到的Kind也是ptr
	//使用Elem()相当于*ptr
	fmt.Println(tye.Elem().Kind())
	fmt.Println(tye.Elem().Elem().Kind())
	//对值再取Elem()，则panic
	//fmt.Println(tye.Elem().Elem().Elem().Kind())
}

func Test_Value(t *testing.T) {
	var c = BMW{
		Info: Info{
			Amount: 100,
			Height: 100,
		},
	}
	var cs []BMW
	cs = append(cs, c)
	tpy := reflect.TypeOf(cs)
	fmt.Println(tpy.Kind())
	fmt.Println(tpy.Elem().Kind())
	//value := reflect.ValueOf(cs)
	//fmt.Println(value.Elem())
}
func Test_ChangeValue(t *testing.T) {
	var c = BMW{
		Info: Info{
			Amount: 100,
			Height: 100,
		},
	}
	//值，不能set
	value := reflect.ValueOf(c)
	fmt.Println(value.CanSet())
	//指针
	fmt.Printf("Amount:%p\n", &c.Amount)
	value = reflect.ValueOf(&c).Elem()
	fmt.Println(value.CanSet())
	for i := 0; i < value.NumField(); i++ {
		fmt.Println(value.Field(i), " i:", i)
	}
	value = value.Field(0).Field(0)
	value.SetInt(200)
	fmt.Println(c)
	fmt.Printf("Amount:%p\n", &c.Amount)
	b, _ := json.Marshal(&c)
	fmt.Println(string(b))
}

func Test_Assign(t *testing.T) {
	var c *BMW
	fmt.Println("c:", c)
	value := reflect.ValueOf(c)
	//同样不能set
	fmt.Println(value.CanSet())
	value = reflect.ValueOf(&c).Elem()
	fmt.Println(value.CanSet())
	//改变指针指向的值
	fmt.Println("Kind:", value.Kind())
	newC := BMW{
		Info: Info{
			Amount: 1200,
			Height: 1299,
		},
	}
	newValue := reflect.ValueOf(&newC)
	value.Set(newValue)
	fmt.Println(c)
}

var m = make(map[string]Car, 0)

func NewCar(Name string, Amount, Height int) Car {
	c, ok := m[Name]
	if !ok {
		return nil
	}
	//赋值
	t := reflect.TypeOf(c)
	switch t.Kind() {
	case reflect.Struct:
		v := reflect.New(t)
		v.Elem().Field(0).Field(0).SetInt(int64(Amount))
		v.Elem().Field(0).Field(1).SetInt(int64(Height))
		return v.Interface().(Car)
	case reflect.Ptr:
		v := reflect.New(t.Elem())
		v.Elem().Field(0).Field(0).SetInt(int64(Amount))
		v.Elem().Field(0).Field(1).SetInt(int64(Height))
		return v.Interface().(Car)
	default:
		panic("x")
	}
	//return nil
}
func Test_Assign1(t *testing.T) {
	m["BMW"] = BMW{}
	m["BENZ"] = &Benz{}

	bmw := NewCar("BMW", 1, 2)
	benz := NewCar("BENZ", 3, 4)
	bmwinfo, _ := bmw.GetInfo()
	fmt.Println(bmwinfo)
	benzinfo, _ := benz.GetInfo()
	fmt.Println(benzinfo)
}

func Test_Interface(t *testing.T) {
	m["BMW"] = BMW{}
	bmw := m["BMW"]
	rt := reflect.TypeOf(bmw)
	fmt.Println("RT:", rt.Kind())
	var a interface{}
	a = 3
	rt = reflect.TypeOf(a)
	fmt.Println("RT:", rt.Kind())
	//获取slice中元素的类型
	//https://stackoverflow.com/questions/18306151/in-go-which-value-s-kind-is-reflect-interface
	d := make([]Car, 0)
	fmt.Println(reflect.TypeOf(d).Elem().Kind() == reflect.Interface)
}

//性能优化https://zhuanlan.zhihu.com/p/138777955
//方法表https://www.cnblogs.com/ksir16/p/9040656.html
