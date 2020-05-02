package main

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strconv"
)

type Cat struct {
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v name:%v kind:%v\n", v, v.Name(), v.Kind())
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本 panic
	}
}

func reflectSetValue2(x interface{}) {
	fmt.Printf("x= %v %T\n", x, x)
	v := reflect.ValueOf(x)
	//要用elem 方法
	//通过反射修改变量值的前提条件之一：这个值必须可以被寻址。
	fmt.Printf("v= %v %T\n", v) //v 的类型不是*int64
	fmt.Printf("v elem= %v %T\n", v.Elem(), v.Elem())
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("runtime caller failed\n")
		return
	}
	fmt.Println(runtime.FuncForPC(pc).Name())
	fmt.Println(file)
	fmt.Println(path.Base(file))
	fmt.Println(line)

	var a float32 = 3.14
	reflectType(a)
	var c = Cat{}
	reflectType(c)
	var b int64 = 100
	reflectType(b)
	//reflectSetValue1(b)
	reflectSetValue2(&b)
	fmt.Println(b)

	str := "10000"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parseint failed, err:", err)
		return
	}
	fmt.Printf("%#v %T", ret1, ret1)
}
