package main

import "fmt"

//批量声明
var (
	name string //""
	age  int    //0
	isOk bool   //false
)

func main() {
	fmt.Println("Hello world!")
	name = "理想"
	age = 16
	isOk = true //Go语言推荐使用驼峰命名
	//Go语言中声明 非全局 变量必须使用，不使用不能编译通过
	//而全局变量可以
	fmt.Print(isOk)
	fmt.Println()
	fmt.Printf("name:%s\n", name)
	fmt.Println(age)

	//类型推导
	var s2 = "20"
	fmt.Println(s2)
	//短变量声明, 只能再函数里面用
	s3 := "哈哈哈"
	fmt.Println(s3)
	//匿名变量_
}
