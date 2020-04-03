package main

import "fmt"

func main() {
	var (
		a = 5
		b = 2
	)
	a++ //单独的语句，不能放在=右边
	b-- //单独的语句，不能放在=右边
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(a == b) //Go 强类型 只有相同类型才能比较

	//数组长度是类型的一部分
	var a1 [3]bool
	fmt.Printf("a1:%T\n", a1)

	//数组的初始化
	fmt.Println(a1)
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//根据初始值自动推断
	a10 := [...]int{0, 1, 2, 3}
	fmt.Println(a10)
	//
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	citys := [...]string{"北京", "上海", "深圳"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	for i, v := range citys {
		fmt.Println(i, v)
	}
	//多维数组
	//[[1 2] [3 4] [5 6]]
	var all [3][2]int
	all = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(all)

	for i, v := range all {
		for n, m := range v {
			fmt.Println(i, n, m)
		}
	}

	//数组是值类型 注意和c不一样
	b1 := [3]int{1, 2, 3}
	b2 := b1 //copy
	b2[0] = 100
	fmt.Println(b1, b2)
}
