package main

import "fmt"

//自定义类型
type myInt int

//类型别名
type yourInt = int

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func f(x person) {
	x.gender = "女" //修改的是副本的gender
}

func f2(x *person) {
	(*x).gender = "女"
	// x.gender = "女"//和上面一样
}

//go 中如果标识符首字母大写就表示对外部可见 共有的 Dog
//Go要求加注释
type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

//方法是作用于特定类型的函数
//接收者 多用类型首字母小写
//不能给别的包类型添加方法
//值接收者
func (d dog) wang() {
	fmt.Println("汪汪汪")
}

//指针接收者
//如果一个方法是指针 保持一致性
func (d *dog) pwang() {
	fmt.Println("p汪汪汪")
}

//匿名字段不常用
//匿名嵌套可以直接访问下一级
type person1 struct {
	string
	int
}

func main() {
	var n myInt
	var m yourInt
	var c rune //类型别名
	n = 100
	fmt.Printf("%T\n", n)
	m = 100
	fmt.Printf("%T\n", m)
	c = '中'
	fmt.Printf("%T\n", c)

	var p person
	p.name = "周琳"
	p.age = 9000
	p.gender = "男"
	p.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(p)
	f(p)
	fmt.Println(p)
	f2(&p)
	fmt.Println(p)

	//匿名
	var s struct {
		x string
		y int
	}
	s.x = "heiheihei"
	s.y = 100
	fmt.Printf("%T %#v\n", s, s)

	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)
	fmt.Printf("%p\n", &p2)

	var p3 = &person{
		name:   "元帅",
		gender: "男",
	}
	fmt.Printf("%T\n", p3) //不是指针
	fmt.Printf("%#v\n", p3)

	//go语言用new开头 构造函数
	//结构体是值类型
	//结构体大返回指针

	d1 := newDog("zhoulin")
	d1.wang()

	//匿名字段
	p1 := person1{"周六", 9000}
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
