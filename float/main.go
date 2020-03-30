package main

import (
	"fmt"
	"strings"
)

func main() {
	// math.MaxFloat32 //浮点32最大值
	f1 := 1.23456
	fmt.Printf("%T\n", f1) //默认float64
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2)
	// f1 = f2
	//float32 不能直接赋值 float64
	b1 := true
	var b2 bool
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T value %v\n", b2, b2)
	var s = "Hello 沙河！"
	fmt.Printf("%#v\n", s)

	s2 := `
	多行字符串
	多行字符串`

	fmt.Println(s2)

	s3 := "a\\b"
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)
	fmt.Println(strings.Contains(s3, "a"))
	//拼接
	fmt.Println(strings.Join(ret, "+"))
}
