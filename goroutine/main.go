package main

import (
	"fmt"
	"runtime"
	"sync"
)

func hello(i int) {
	fmt.Println("hello", i)
	defer wg.Done()
}

var wg sync.WaitGroup

var b chan int

//单项通道
func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 chan int, ch2 chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch2)
}

//main 也是一个goroutin
func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go hello(i)
	}

	fmt.Println("main")
	//main 函数结束 goroutine 都结束
	// time.Sleep(time.Second)
	wg.Wait()
	fmt.Printf("cpu: %d\n", runtime.NumCPU())

	a := make(chan int, 100) //不带缓冲区通道的初始化
	b = make(chan int, 100)  //带缓冲区的通道 塞满堵死
	wg.Add(2)
	go f1(a)
	go f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}

	//多路复用
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select { //随机 or 哪个成功走哪个
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
