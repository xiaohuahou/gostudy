package main

import "fmt"

func main() {
	s2 := "白萝卜"
	s3 := []rune(s2) //把字符串强制转换成rune切片, rune(int32)
	//byte(uint8)
	s3[0] = '红'
	fmt.Println(string(s3))

	//类型转换
	n := 10
	var f float64
	f = float64(n)
	fmt.Printf("%T\n", f)
}
