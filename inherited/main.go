package main

import (
	"encoding/json"
	"fmt"
)

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s 动\n", a.name)
}

type dog struct {
	feet   uint8
	animal //animal 有的方法狗都有
}

func (d dog) wang() {
	fmt.Printf("%s wang\n", d.name)
}

type person struct {
	Name string `json:"name" db:"name"` //可见性
	Age  int    `json:"age"`            //或者用tag
}

func main() {
	d1 := dog{
		animal: animal{name: "zhoulin"},
		feet:   4,
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()

	p1 := person{
		Name: "周玲",
		Age:  9000,
	}

	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Println(string(b))

	//反序列化
	str := `{"name":"理想","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传值针
	fmt.Printf("%#v\n", p2)
}
