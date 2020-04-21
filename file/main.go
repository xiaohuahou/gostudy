package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file fialed, %v\n", err)
		return
	}
	defer f.Close()

	var tmp [128]byte
	for {
		n, err := f.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read file failed, %v\n", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}
