package main

import "fmt"

//常量
const pi = 3.1415926

const (
	statusOK = 200
	notFound = 404
)

const (
	n1 = 100
	n2
	n3
)

const (
	a1 = iota //0
	a2
	a3
)

const b1 = iota
const b2 = iota //重置为0

const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4
)

const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	// pi = 123
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)
	// iota 常量计数器
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)

	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)

	//整形
	var i1 = 101 //类型推导
	fmt.Printf("%d\n", i1)
	fmt.Printf("%o\n", i1)
	fmt.Printf("%x\n", i1)
	fmt.Printf("%b\n", i1)
	i2 := 077
	fmt.Printf("%d\n", i2)
	i3 := 0xff
	fmt.Printf("%d\n", i3)

	//查看变量类型
	fmt.Printf("%T\n", i3)

	i4 := int8(9)
	fmt.Printf("%T\n", i4)
}
