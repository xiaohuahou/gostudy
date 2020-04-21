package main

import "fmt"

//defer 多用于资源释放
func deferdemo() {
	fmt.Println("start")
	//defer 把语句延迟刀返回的时候再执行
	defer fmt.Println("嘿嘿嘿")
	defer fmt.Println("哈哈哈") //defer 先进后出
	fmt.Println("end")
}

//return x
//1. x = value
//2. defer
//3. ret
func f1() int {
	x := 5
	defer func() {
		x++ //修改的是x 不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回值赋值给x
}

func f4() (x int) {
	defer func(x int) {
		x++ //copy
	}(x)
	return 5 //返回值赋值给x
}

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func funcA() {
	fmt.Println("a")
}
func funcB() {
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接。。。")
	}()
	panic("出现了严重的错误")
	fmt.Println("b")
}
func funcC() {
	fmt.Println("c")
}

func main() {
	deferdemo()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f4())
	//函数内部只能用匿名函数

	ret := adder(100)
	ret2 := ret(200)
	fmt.Println(ret2)
	ret3 := ret(200)
	fmt.Println(ret3)

	// a := 1
	// b := 2
	//只defer 一层
	//a为参数入栈是的值 不是最后defer执行时的值
	// defer calc("1", a, calc("10", a, b))
	// a = 0
	// defer calc("2", a, calc("20", a, b))
	// b = 1

	//panic and recover
	funcA()
	funcB()
	funcC()
}
