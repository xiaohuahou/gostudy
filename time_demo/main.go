package main

import (
	"fmt"
	"time"
)

//时间

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	//time.Unix()
	fmt.Println(time.Unix(1564803667, 0))

	//定时器
	timer := time.Tick(time.Second)
	for t := range timer {
		fmt.Println(t) //1秒钟执行一次 本质是一个channel
		break
	}

	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM"))

	//还有一个parseinlocation
	timeObj, err := time.Parse("2006-01-02", "2000-08-03")
	if err != nil {
		fmt.Println("parse time fialed")
		return
	}
	fmt.Println(timeObj)
}
