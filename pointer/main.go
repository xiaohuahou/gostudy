package main

import "fmt"

//命名返回值相当于声明了变量
func sum(x int, y int) (ret int) {
	ret = x + y
	return ret
}

func f1(x int, y int) {
	fmt.Println(x + y)
}

func f3() int {
	return 3
}

func f5() (int, string) {
	return 1, "沙河"
}

//可变参数，必须放在最后
//go语言没有默认参数
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) //y 的类型是切片
}

//类型简写
func f6(x, y int) int {
	return x + y
}

//vs 不支持go mudule
func main() {
	//1. &
	//2. *
	//go 指针只能读不能修改
	n := 18
	fmt.Println(&n)

	p := &n
	fmt.Println(*p)

	var a = new(int)
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b = make(map[string]int, 10) //避免动态扩容
	b["沙河娜扎"] = 100
	fmt.Println(b)
	//如果key 不存在这个key拿到对应类型的零值
	value, ok := b["娜扎"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}

	delete(b, "娜扎") //不报错

	//类型为map的切片
	var s1 = make([]map[int]string, 10, 10)
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "沙河"
	// s1[1][0] = "shahe" //没有开空间
	fmt.Println(s1)

	//值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30}
	fmt.Println(m1)

	r := sum(1, 2)
	fmt.Println(r)
}
