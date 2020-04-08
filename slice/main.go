package main

import "fmt"

func main() {
	//切片的定义s
	var s1 []int //定义一个存放int的切片 底层是一个数组
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) //没有开辟空间
	fmt.Println(s2 == nil)
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张江", "平山村"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4]
	s6 := a1[3:]
	fmt.Println(s3)
	//切片的容量是底层数组的容量
	fmt.Printf("len(s3):%d cap(s3):%d\n", len(s3), cap(s3))
	//底层数组从切片的第一个元素刀最后一个元素
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6))
	//切片再切片
	s8 := s6[3:]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8))

	//make
	s9 := make([]int, 5, 10)
	fmt.Printf("s9=%v len(s9)=%d cap(s9)=%d\n", s9, len(s9), cap(s9))

	//切片的数值拷贝
	c1 := []int{1, 2, 3}
	c2 := c1
	c1[0] = 1000
	fmt.Println(c1, c2)

	//append
	c3 := []string{"bj", "sh", "sz"}
	// c3[3] = "gz" 越界
	c3 = append(c3, "gz")
	fmt.Println(c3)
	cc := []string{"wh", "xa", "sz"}
	c3 = append(c3, cc...) //...表示拆开
	fmt.Println(c3)
}
