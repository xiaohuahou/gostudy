package main

import "fmt"

//接口是一种类型
type speaker interface {
	speak()
}

type cat struct {
}

type dog struct {
}

type person struct {
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func (p person) speak() {
	fmt.Println("啊啊啊")
}
func da(x speaker) {
	//接受一个参数，传进来什么就打什么
	x.speak()
}

//空接口
//所有类型都实现了空接口 可以存储任何类型
//interface 关键字
//interface{} 类型

//类型断言
func assign(a interface{}) {
	fmt.Println("%T\n", a)
	switch t:=a.(type) {
	case string:
		fmt.Println("是一个字符串"，t)
	case string:
		fmt.Println("是一个int"，t)
	}

	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println(str)
	}
}

//全局声明->init->main
func init(){
	fmt.Println("自动执行，不能主动调用")
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person
	da(c1)
	da(d1)
	da(p1)

	//接口保存两部分 动态类型和动态值
	var ss speaker
	ss = c1
	ss = d1
	ss = p1
	fmt.Println(ss)

	//使用值接收者和指针接收者的区别
	//值接收者 值和指针变量都能存
	//指针接收者 只能存结构体指针类型变量

	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "周六"
	m1["age"] = 9000
	m1["merried"] = true
	m1["hobby1"] = [...]string{"唱", "跳", "rap"}

	//类型断言
	assign(100)
}
