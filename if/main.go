package main

import "fmt"

func main() {

	if age := 19; age > 18 { //作用域
		fmt.Println("澳门")
	} else {
		fmt.Println("否则")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	s := "Hello沙河"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
