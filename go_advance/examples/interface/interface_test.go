package main

import (
	"fmt"
	"testing"
)

func Print(s interface{}) {
	fmt.Printf("s=(%T,%v)\n", s, s)
	fmt.Println("s == nil ", s == nil)
}

func PrintP(s *int) {
	fmt.Printf("s=(%T,%v)\n", s, s)
	fmt.Println("s == nil ", s == nil)
}

type F func(int) int

func FA(a int) int {
	return a
}
func Test_InterfaceEmpty(t *testing.T) {
	var a int
	Print(a)
	var b interface{}
	Print(b)
	p := &a
	Print(p)
	var d *int
	PrintP(d)
	var f F
	Print(f)
	f = FA
	Print(f)
}

func Assert(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("int v=(%T,%v)\n", v, v)
	case *int:
		fmt.Printf("*int v=(%T,%v)\n", v, v)
	}
}

func Test_Assert(t *testing.T) {
	var a int
	Assert(a)
	fmt.Println("xxx")
	Assert(&a)
	var v interface{}
	v = a
	if b, ok := v.(int); ok {
		fmt.Println("convert:", b)
	}
}

//
//
//

type Car interface {
	Run(int) error
	GetInfo() (Info, error)
	ChangeAmount(int)
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

func (b *Benz) ChangeAmount(amount int) {
	b.Amount = amount
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

func (b BMW) ChangeAmount(amount int) {
	b.Amount = amount
}

var _ Car = (*BMW)(nil)

func Test_Benz(t *testing.T) {
	b := &Benz{
		Info: Info{
			Amount: 100,
			Height: 4,
		},
	}
	var c Car = b
	info, _ := c.GetInfo()
	fmt.Println("Car INfo:", info)
	_ = c.Run(3)
}

func Test_BMW(t *testing.T) {
	b := BMW{
		Info: Info{
			Amount: 100,
			Height: 4,
		},
	}
	bp := &BMW{
		Info: Info{
			Amount: 1000,
			Height: 0,
		},
	}
	var c Car = b
	info, _ := c.GetInfo()
	fmt.Println("Car INfo:", info)
	_ = c.Run(3)
	var cp Car = bp
	info, _ = cp.GetInfo()
	fmt.Println("Car INfo:", info)
	_ = cp.Run(3)
}

func Test_Cars(t *testing.T) {
	b := &Benz{
		Info: Info{
			Amount: 100,
			Height: 4,
		},
	}
	bmw := BMW{
		Info: Info{
			Amount: 1,
			Height: 1,
		},
	}
	var cars []Car
	cars = append(cars, b, bmw)
	for _, v := range cars {
		info, _ := v.GetInfo()
		fmt.Println("Car INfo:", info)
		v.Run(2)
	}
}

type Combine interface {
	Car
	Trans()
}

func (b BMW) Trans() {
	fmt.Println("BMW TRANS")
}

func Test_Conbine(t *testing.T) {
	bmw := BMW{}
	//benz := Benz{}
	var c Combine = bmw
	fmt.Printf("v=(%T,%v)\n", c, c)
	if cc, ok := c.(Car); ok {
		fmt.Println("Can cast:")
		cc.Run(2)
	}
}

func Test_Change(t *testing.T) {
	bmw := BMW{}
	var c Car = bmw
	c.ChangeAmount(19)
	info, _ := c.GetInfo()
	fmt.Println(info)
	//值不能改变
	var cp Car = &bmw
	//虽然是指针，传递的时候被取值
	cp.ChangeAmount(200)
	info, _ = cp.GetInfo()
	fmt.Println(info)

	//指针ok
	var bCar = &Benz{}
	bCar.ChangeAmount(1222)
	info, _ = bCar.GetInfo()
	fmt.Println(info)
}

func Test_CarP(t *testing.T) {
	//报错
	//var CarP *Car = &BMW{}
	//报错
	//var CarP *Car = &Benz{}
	// *interface{} != interface{}
	//bmw := BMW{}
	//PassCarP(&bmw)
	//benz := &Benz{}
	//PassCarP(benz)
}

func PassCarP(carP *Car) {
	fmt.Println("Call")
}

func PassInterfaceP(ip *interface{}) {
	fmt.Println("Call")
}
func TestPassInter(t *testing.T) {
	//bmw := BMW{}
	//PassInterfaceP(&bmw)
	//benz := &Benz{}
	//PassInterfaceP(benz)
}

func TestPrintInter(t *testing.T) {
	bmw := BMW{}
	Print(bmw)
	benz := &Benz{}
	Print(benz)
	Print(&bmw)
	Print(&benz)
}

func work(arr []interface{}) {
	fmt.Println(arr)
}

func TestSliceInterface(t *testing.T) {
	//ss := []string{"hello", "golang"}
	//work(ss) //不能协变
	st := []interface{}{"hello", "golang"}
	work(st)
}

type Robot struct {
	Head int
	Car
}

func TestStructInterface(t *testing.T) {
	var robot = Robot{
		Head: 12,
		Car: &BMW{
			Info: Info{
				Amount: 100,
				Height: 1,
			},
		},
	}
	t.Log(robot.Run(12))
}
